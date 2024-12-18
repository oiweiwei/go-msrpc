package wkssvc

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
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
)

var (
	// import guard
	GoPackage = "wkst"
)

var (
	// Syntax UUID
	WkssvcSyntaxUUID = &uuid.UUID{TimeLow: 0x6bffd098, TimeMid: 0xa112, TimeHiAndVersion: 0x3610, ClockSeqHiAndReserved: 0x98, ClockSeqLow: 0x33, Node: [6]uint8{0x46, 0xc3, 0xf8, 0x7e, 0x34, 0x5a}}
	// Syntax ID
	WkssvcSyntaxV1_0 = &dcerpc.SyntaxID{IfUUID: WkssvcSyntaxUUID, IfVersionMajor: 1, IfVersionMinor: 0}
)

// wkssvc interface.
type WkssvcClient interface {

	// The NetrWkstaGetInfo method returns details about the configuration of a remote computer,
	// including the computer name and major and minor version numbers of the operating
	// system.
	//
	// Return Values: When the message processing result meets the description in column
	// two of the following table, this method MUST return one of the following values ([MS-ERREF]
	// section 2.2).
	//
	//	+--------------------------------+--------------------------------------------------------------------+
	//	|                                |                                                                    |
	//	|           VALUE/CODE           |                              MEANING                               |
	//	|                                |                                                                    |
	//	+--------------------------------+--------------------------------------------------------------------+
	//	+--------------------------------+--------------------------------------------------------------------+
	//	| NERR_Success 0x00000000        | The operation completed successfully.                              |
	//	+--------------------------------+--------------------------------------------------------------------+
	//	| ERROR_ACCESS_DENIED 0x00000005 | The caller does not have the permissions to perform the operation. |
	//	+--------------------------------+--------------------------------------------------------------------+
	//	| ERROR_INVALID_LEVEL 0x0000007C | The information level is invalid.                                  |
	//	+--------------------------------+--------------------------------------------------------------------+
	GetInfo(context.Context, *GetInfoRequest, ...dcerpc.CallOption) (*GetInfoResponse, error)

	// The NetrWkstaSetInfo method configures a remote computer according to the information
	// structure passed in the call.
	//
	// Return Values: When the message processing result meets the description in column
	// two of the following table, this method MUST return one of the following values ([MS-ERREF]
	// section 2.2).
	//
	//	+------------------------------------+----------------------------------------------+
	//	|                                    |                                              |
	//	|             VALUE/CODE             |                   MEANING                    |
	//	|                                    |                                              |
	//	+------------------------------------+----------------------------------------------+
	//	+------------------------------------+----------------------------------------------+
	//	| NERR_Success 0x00000000            | The operation completed successfully.        |
	//	+------------------------------------+----------------------------------------------+
	//	| ERROR_ACCESS_DENIED 0x00000005     | Access is denied.                            |
	//	+------------------------------------+----------------------------------------------+
	//	| ERROR_INVALID_PARAMETER 0x00000057 | One of the function parameters is not valid. |
	//	+------------------------------------+----------------------------------------------+
	//
	// Any other return value MUST conform to the error code requirements in Protocol Details
	// (section 3).
	//
	// If the Level parameter value is not valid, the server MUST fail the call as follows.
	//
	// Note  All value ranges are inclusive.
	//
	//	+----------------------------------------------------------------------------------+---------------------------------------------------+
	//	|                                  INVALID LEVEL                                   |                      FAILURE                      |
	//	|                                      VALUE                                       |                    PROCESSING                     |
	//	+----------------------------------------------------------------------------------+---------------------------------------------------+
	//	+----------------------------------------------------------------------------------+---------------------------------------------------+
	//	| 0x00000000--0x000001F5, 0x000001F7--0x000003F4, 0x000003F6--0x000003F9,          | The server SHOULD return ERROR_INVALID_LEVEL.<31> |
	//	| 0x000003FB--0x00000415 0x00000417--0xFFFFFFFF                                    |                                                   |
	//	+----------------------------------------------------------------------------------+---------------------------------------------------+
	//
	// Otherwise, if the Level parameter equals 0x000001F6, the server MUST store values
	// from the WKSTA_INFO_502 structure WkstaInfo502 member (section 2.2.4.1) specified
	// by the WkstaInfo parameter into elements of the ADM (section 3.2.1) as follows:
	//
	// * *wki502_keep_conn* stored in *Keep_Connection*
	//
	// * *wki502_max_cmds* stored in *Max_Commands*
	//
	// * *wki502_sess_timeout* stored in *Session_TimeOut*
	//
	// * *wki502_dormant_file_limit* stored in *DormantFileLimit*
	//
	// If the Level parameter equals 0x000003F5, the server MUST store values from the WKSTA_INFO_1013
	// structure WkstaInfo1013 member (section 2.2.4.1) of the WkstaInfo parameter into
	// elements of the ADM, as follows:
	//
	// * *wki1013_keep_conn* stored in *Keep_Connection*.
	//
	// If the Level parameter equals 0x000003FA, the server MUST store values from the WKSTA_INFO_1018
	// structure WkstaInfo1018 member (section 2.2.4.1) of the WkstaInfo parameter into
	// elements of the ADM, as follows:
	//
	// * *wki1018_sess_timeout* stored in *Session_TimeOut*.
	//
	// If the Level parameter equals 0x00000416, the server MUST store values from the WKSTA_INFO_1046
	// structure WkstaInfo1046 member (section 2.2.4.1) of the WkstaInfo parameter into
	// elements of the ADM, as follows:
	//
	// * *wki1046_dormant_file_limit* stored in *DormantFileLimit*.
	SetInfo(context.Context, *SetInfoRequest, ...dcerpc.CallOption) (*SetInfoResponse, error)

	// The NetrWkstaUserEnum method returns details about users who are currently active
	// on a remote computer.
	//
	// Return Values: When the message processing result matches the description in column
	// two of the following table, this method MUST return one of the following values ([MS-ERREF]
	// section 2.2). The most common error codes are listed in the following table.
	//
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	|                                |                                                                                  |
	//	|           VALUE/CODE           |                                     MEANING                                      |
	//	|                                |                                                                                  |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	| NERR_Success 0x00000000        | The operation completed successfully.                                            |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_ACCESS_DENIED 0x00000005 | Access is denied.                                                                |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_INVALID_LEVEL 0x0000007C | The information level is invalid.                                                |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_MORE_DATA 0x000000EA     | More entries are available. The UserInfo buffer was not large enough to contain  |
	//	|                                | all the entries.                                                                 |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//
	// Any other return value MUST conform to the error code requirements in Protocol Details
	// (section 3).
	//
	// The server (2) SHOULD<35> enforce security measures to verify that the caller has
	// the required permissions to execute this routine. If the server enforces security
	// measures, and the caller does not have the required credentials, the server MUST
	// fail the call with ERROR_ACCESS_DENIED. Specifications for determining the identity
	// of the caller for the purpose of performing an access check are in [MS-RPCE] section
	// 3.3.3.1.3.
	//
	// If the Level member of the WKSTA_USER_ENUM_STRUCT structure passed in the UserInfo
	// parameter does not equal 0x00000000 or 0x00000001, the server MUST fail the call.
	//
	// If the Level member equals 0x00000000, the server MUST return an array of the names
	// of users currently logged on the computer by filling the WKSTA_USER_INFO_0_CONTAINER
	// structure (section 2.2.5.14) in the WkstaUserInfo field of the UserInfo parameter.
	//
	// If the Level member equals 0x00000001, the server MUST return an array of the names
	// and domain details of each user currently logged on the computer, and a list of OtherDomains
	// (section 3.2.1.3) in the computer.
	//
	// If the PreferredMaximumLength parameter equals MAX_PREFERRED_LENGTH (section 2.2.1.3),
	// the server MUST return all the requested data. Otherwise, if the PreferredMaximumLength
	// is insufficient to hold all the entries, then the server MUST return the maximum
	// number of entries that fit in the UserInfo buffer and return ERROR_MORE_DATA.
	//
	// The following rules specify processing of the ResumeHandle parameter:
	//
	// * If the ResumeHandle parameter is either NULL or points to 0x00000000, the enumeration
	// MUST start from the beginning of the list of the currently logged on users. <36>
	// ( def80006-2495-4571-8a93-1668e0f8af31#Appendix_A_36 )
	//
	// * If the ResumeHandle parameter points to a non-zero value, the server MUST continue
	// enumeration based on the value of ResumeHandle. The server is not required to maintain
	// any state between calls to the NetrWkstaUserEnum method.
	//
	// * If the client specifies a ResumeHandle , and if the server returns ERROR_MORE_DATA,
	// the server MUST set the value to which ResumeHandle points to an implementation-specific
	// value that allow the server to continue with this enumeration on a subsequent call
	// to this method, with the same value for ResumeHandle.
	//
	// The server is not required to maintain any state between calls to the NetrWkstaUserEnum
	// method. If the server returns NERR_Success or ERROR_MORE_DATA, it MUST set the TotalEntries
	// parameter to equal the total number of entries that could have been enumerated from
	// the current resume position.
	UserEnum(context.Context, *UserEnumRequest, ...dcerpc.CallOption) (*UserEnumResponse, error)

	// Opnum3NotUsedOnWire operation.
	// Opnum3NotUsedOnWire

	// Opnum4NotUsedOnWire operation.
	// Opnum4NotUsedOnWire

	// The NetrWkstaTransportEnum method provides details about the transport protocols
	// currently enabled for use by the SMB network redirector on a remote computer.
	//
	// Return Values: When the message processing result matches the description in column
	// two of the following table, this method MUST return one of the following values ([MS-ERREF]
	// section 2.2).
	//
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	|                                |                                                                                  |
	//	|           VALUE/CODE           |                                     MEANING                                      |
	//	|                                |                                                                                  |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	| NERR_Success 0x00000000        | The operation completed successfully.                                            |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_ACCESS_DENIED 0x00000005 | Access is denied.                                                                |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_INVALID_LEVEL 0x0000007C | The information level is invalid.                                                |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	| NERR_BufTooSmall 0x0000084B    | More entries are available. The TransportInfo buffer was not large enough to     |
	//	|                                | contain all the entries.                                                         |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//
	// Any other return value MUST conform to the error code requirements in Protocol Details
	// (section 3).
	//
	// The server SHOULD<38> enforce security measures to verify that the caller has the
	// required permissions to execute this routine. If the server enforces security measures,
	// and the caller does not have the required credentials, the server MUST fail the call
	// with ERROR_ACCESS_DENIED. Specifications for determining the identity of the caller
	// for the purpose of performing an access check are in [MS-RPCE] section 3.3.3.1.3.
	//
	// For any other conditions, this method MUST return any other value, and the client
	// MUST treat all other values the same.
	//
	// If the Level member in the WKSTA_TRANSPORT_ENUM_STRUCT structure passed in the TransportInfo
	// parameter does not equal 0x00000000, the server MUST fail the call.
	//
	// If the Level member is 0x00000000, the server MUST return an array of details about
	// the transport protocols currently enabled for use by the SMB network redirector by
	// filling the WKSTA_TRANSPORT_INFO_0_CONTAINER structure WkstaTransportInfo member
	// (section 2.2.5.15) of the TransportInfo parameter for each transport in TransportList
	// (section 3.2.1.4), as follows:
	//
	// * *wkti0_transport_address* set to *Transport.Address*
	//
	// * *wkti0_transport_name* set to *Transport.Name*
	//
	// * *wkti0_number_of_vcs* set to *Transport.VC_Count*
	//
	// * *wkti0_wan_ish* set to *Transport.Wannish*
	//
	// If the PreferredMaximumLength parameter equals MAX_PREFERRED_LENGTH (section 2.2.1.3),
	// the server MUST return all the requested data. If the PreferredMaximumLength is insufficient
	// to hold all the entries, the server MUST return the maximum number of entries that
	// fit in the TransportInfo buffer and return NERR_BufTooSmall.
	//
	// The following rules specify processing of the ResumeHandle parameter:
	//
	// * If the ResumeHandle parameter is either NULL or points to 0x00000000, the enumeration
	// MUST start from the beginning of the list of the currently enabled transport protocols.
	// <39> ( def80006-2495-4571-8a93-1668e0f8af31#Appendix_A_39 )
	//
	// * If the ResumeHandle parameter is nonzero, the server MUST begin enumeration based
	// on the value of ResumeHandle. The server is not required to maintain any state between
	// calls invoking the *NetrWkstaTransportEnum* method.
	//
	// * If the client specified a ResumeHandle , and if the server returns NERR_BufTooSmall,
	// the server MUST set ResumeHandle to an implementation-specific value that allows
	// the server to continue with this enumeration on a subsequent call to this method,
	// using the same value for ResumeHandle.
	//
	// The server is not required to maintain any state between calls to the NetrWkstaTransportEnum
	// method. If the server returns NERR_Success, it MUST set the TotalEntries parameter
	// to equal the total number of entries that could have been enumerated from the current
	// resume position. If the server returns NERR_BufTooSmall, it SHOULD set the TotalEntries
	// value to the total number of entries that could have been enumerated from the current
	// resume position.<40>
	TransportEnum(context.Context, *TransportEnumRequest, ...dcerpc.CallOption) (*TransportEnumResponse, error)

	// The NetrWkstaTransportAdd method enables the SMB network redirector to use a transport
	// protocol on a remote computer.
	//
	// Return Values: When the message processing result meets the description in column
	// two of the following table, this method MUST return one of the following values ([MS-ERREF]
	// section 2.2).
	//
	//	+------------------------------------+----------------------------------------------+
	//	|                                    |                                              |
	//	|             VALUE/CODE             |                   MEANING                    |
	//	|                                    |                                              |
	//	+------------------------------------+----------------------------------------------+
	//	+------------------------------------+----------------------------------------------+
	//	| NERR_Success 0x00000000            | The operation completed successfully.        |
	//	+------------------------------------+----------------------------------------------+
	//	| ERROR_ACCESS_DENIED 0x00000005     | Access is denied.                            |
	//	+------------------------------------+----------------------------------------------+
	//	| ERROR_INVALID_PARAMETER 0x00000057 | One of the function parameters is not valid. |
	//	+------------------------------------+----------------------------------------------+
	//	| ERROR_INVALID_LEVEL 0x0000007C     | The information level is invalid.            |
	//	+------------------------------------+----------------------------------------------+
	//
	// Any other return value MUST conform to the error code requirements in Protocol Details
	// (section 3).
	//
	// If the Level parameter is not equal to zero, the server MUST fail the call and return
	// ERROR_INVALID_LEVEL.
	//
	// If the server does not support this method then it SHOULD be processed as follows.
	//
	// If any of the input parameters are invalid, the server SHOULD return ERROR_INVALID_PARAMETER.
	// Otherwise, it SHOULD<41> return NERR_Success.
	//
	// The server SHOULD<42> enforce security measures to verify that the caller has authorization
	// to execute this routine. If the server enforces security measures and the caller
	// does not have the required credentials, the server SHOULD fail the call with ERROR_ACCESS_DENIED.
	// Specifications for determining the identity of the caller for the purpose of performing
	// an access check are in [MS-RPCE] section 3.3.3.1.3.
	//
	// The TransportInfo parameter specifies details about the transport protocol that is
	// to be enabled. If any of the input parameters are invalid, the server MUST return
	// ERROR_INVALID_PARAMETER. If the caller has passed the ErrorParameter parameter, the
	// server MUST return the zero-based index of the first member of the structure the
	// TransportInfo parameter points to that was invalid.
	//
	// If this method call is successful, the server MUST store values from members of the
	// WKSTA_TRANSPORT_INFO_0 structure passed in the TransportInfo parameter into the ADM
	// elements for each transport in TransportList (section 3.2.1.4) as follows:
	//
	// * *wkti0_transport_address* stored in *Transport.Address*
	//
	// * *wkti0_transport_name* stored in *Transport.Name*
	//
	// * *wkti0_number_of_vcs* stored in *Transport.VC_Count*
	//
	// * *wkti0_wan_ish* stored in *Transport.Wannish*
	TransportAdd(context.Context, *TransportAddRequest, ...dcerpc.CallOption) (*TransportAddResponse, error)

	// The NetrWkstaTransportDel method disables the use of a transport protocol by the
	// SMB network redirector on a remote computer. The transport can be re-enabled by calling
	// the NetrWkstaTransportAdd method (section 3.2.4.5).
	//
	// Return Values: When the message processing result meets the description in column
	// two of the following table, this method MUST return one of the following values ([MS-ERREF]
	// section 2.2).
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|                                    |                                                                                  |
	//	|             VALUE/CODE             |                                     MEANING                                      |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| NERR_Success 0x00000000            | The operation completed successfully.                                            |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_ACCESS_DENIED 0x00000005     | Access is denied.                                                                |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_INVALID_PARAMETER 0x00000057 | One of the function parameters is invalid.                                       |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_OPEN_FILES 0x00002401        | There are open files, or printer handles are using the transport protocol        |
	//	|                                    | pending on this connection.                                                      |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_DEVICE_IN_USE 0x00002404     | The device or open directory handle is using the transport protocol and cannot   |
	//	|                                    | be disconnected.                                                                 |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//
	// If the ForceLevel parameter does not equal 0x00000000, 0x00000001, or 0x00000002,
	// the server MUST fail the call with ERROR_INVALID_PARAMETER. If the ForceLevel parameter
	// is 0x00000000 or 0x00000001 and any open directory handle is using the transport
	// protocol provided in the TransportName field, the server MUST fail the call with
	// ERROR_DEVICE_IN_USE. If the ForceLevel parameter is 0x00000000 or 0x00000001 and
	// any open files or printer handles are using the transport protocol provided in the
	// TransportName field, fail the call with ERROR_OPEN_FILES.
	//
	// If the server does not support this method, it SHOULD<43> return NERR_Success if
	// the ForceLevel parameter is valid. If the server does support this method, it MUST
	// be processed as follows.
	//
	// The server SHOULD<44> enforce security measures to verify that the caller has the
	// required permissions to execute this routine. If the server enforces security measures
	// and the caller does not have the required credentials, the server MUST fail the call
	// with ERROR_ACCESS_DENIED. Specifications for determining the identity of the caller
	// for the purpose of performing an access check are in ([MS-RPCE] section 3.3.3.1.3.
	//
	// If any open file or printer handles are using the transport protocol that this call
	// is trying to disable, the server behavior MUST depend on the value of the ForceLevel
	// parameter. If the ForceLevel parameter is 0x00000000 or 0x00000001, the server MUST
	// fail the call. If the ForceLevel parameter is 0x00000002, the server MUST forcefully
	// close all open handles and disable the transport protocol.
	//
	// If this method call is successful, the server MUST remove this protocol from its
	// list of currently enabled transport protocols.
	TransportDelete(context.Context, *TransportDeleteRequest, ...dcerpc.CallOption) (*TransportDeleteResponse, error)

	// The NetrUseAdd method establishes a connection between the workstation server and
	// an SMB server. Workstation servers SHOULD NOT allow this method to be invoked remotely<45>
	// and SHOULD return ERROR_CALL_NOT_IMPLEMENTED.
	//
	// Return Values: When the message processing result meets the description in the right-hand
	// column of the following table, this method MUST return one of the following values
	// ([MS-ERREF] section 2.2).
	//
	//	+---------------------------------------+------------------------------------------------+
	//	|                                       |                                                |
	//	|              VALUE/CODE               |                    MEANING                     |
	//	|                                       |                                                |
	//	+---------------------------------------+------------------------------------------------+
	//	+---------------------------------------+------------------------------------------------+
	//	| NERR_Success 0x00000000               | The operation completed successfully.          |
	//	+---------------------------------------+------------------------------------------------+
	//	| ERROR_ACCESS_DENIED 0x00000005        | Access is denied.                              |
	//	+---------------------------------------+------------------------------------------------+
	//	| ERROR_INVALID_PARAMETER 0x00000057    | One of the function parameters is not valid.   |
	//	+---------------------------------------+------------------------------------------------+
	//	| ERROR_INVALID_LEVEL 0x0000007C        | The information level is invalid.              |
	//	+---------------------------------------+------------------------------------------------+
	//	| ERROR_CALL_NOT_IMPLEMENTED 0x00000078 | This function is not supported on this system. |
	//	+---------------------------------------+------------------------------------------------+
	//
	// The server SHOULD<46> enforce security measures to verify that the caller has the
	// required permissions to execute this routine. If the server enforces security measures
	// and the caller does not have the required credentials, the server SHOULD fail the
	// call with ERROR_ACCESS_DENIED. Specifications for determining the identity of the
	// caller for the purpose of performing an access check are in [MS-RPCE] section 3.3.3.1.3.
	//
	// The Level parameter determines the type of structure that the client has used to
	// specify details about the new connection. The value MUST be 0, 1, 2, or 3. If the
	// Level parameter is not equal to one of the valid values, the server MUST fail the
	// call with an ERROR_INVALID_LEVEL error code.
	//
	// * If the Level parameter is 0x00000000, the Buffer parameter points to a *USE_INFO_0*
	// structure.
	//
	// * If the Level parameter is 0x00000001, the Buffer parameter points to a *USE_INFO_1*
	// structure.
	//
	// * If the Level parameter is 0x00000002, the Buffer parameter points to a *USE_INFO_2*
	// structure.
	//
	// * If the Level parameter is 0x00000003, the Buffer parameter points to a *USE_INFO_3*
	// structure.
	UseAdd(context.Context, *UseAddRequest, ...dcerpc.CallOption) (*UseAddResponse, error)

	// The NetrUseGetInfo method retrieves details from a remote workstation about a connection
	// to a shared resource on an SMB server. The server SHOULD NOT allow this method to
	// be invoked remotely<47> and SHOULD return ERROR_CALL_NOT_IMPLEMENTED.
	//
	// Return Values: When the message processing result meets the description in the right-hand
	// column of the following table, this method MUST return one of the following values
	// ([MS-ERREF] section 2.2).
	//
	//	+------------------------------------+----------------------------------------------+
	//	|                                    |                                              |
	//	|             VALUE/CODE             |                   MEANING                    |
	//	|                                    |                                              |
	//	+------------------------------------+----------------------------------------------+
	//	+------------------------------------+----------------------------------------------+
	//	| NERR_Success 0x00000000            | The operation completed successfully.        |
	//	+------------------------------------+----------------------------------------------+
	//	| ERROR_ACCESS_DENIED 0x00000005     | Access is denied.                            |
	//	+------------------------------------+----------------------------------------------+
	//	| ERROR_INVALID_PARAMETER 0x00000057 | One of the function parameters is not valid. |
	//	+------------------------------------+----------------------------------------------+
	//	| ERROR_INVALID_LEVEL 0x0000007C     | The information level is invalid.            |
	//	+------------------------------------+----------------------------------------------+
	//	| NERR_UseNotFound 0x000008CA        | The network connection could not be found.   |
	//	+------------------------------------+----------------------------------------------+
	//
	// The server SHOULD<48> enforce security measures to verify that the caller has the
	// required permissions to execute this routine. If the server enforces security measures
	// and the caller does not have the required credentials, the server MUST fail the call
	// with ERROR_ACCESS_DENIED. Specifications for determining the identity of the caller
	// for the purpose of performing an access check are in [MS-RPCE] section 3.3.3.1.3.
	//
	// The UseName parameter specifies the local device name or shared resource name for
	// which to return information. The server MUST canonicalize UseName ([MS-SRVS] section
	// 3.1.4.33). This MUST be a nonempty, null-terminated UTF-16 string; otherwise, the
	// server MUST fail the call with an ERROR_INVALID_PARAMETER error code.
	UseGetInfo(context.Context, *UseGetInfoRequest, ...dcerpc.CallOption) (*UseGetInfoResponse, error)

	// The NetrUseDel function terminates a connection from the workstation server to a
	// shared resource on an SMB server. The server SHOULD NOT<49> allow this method to
	// be invoked remotely and SHOULD return ERROR_CALL_NOT_IMPLEMENTED.
	//
	// Return Values: When the message processing result meets the description in the right-hand
	// column of the following table, this method MUST return one of the following values
	// ([MS-ERREF] section 2.2).
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|                                    |                                                                                  |
	//	|             VALUE/CODE             |                                     MEANING                                      |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| NERR_Success 0x00000000            | The operation completed successfully.                                            |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_ACCESS_DENIED 0x00000005     | Access is denied.                                                                |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_INVALID_PARAMETER 0x00000057 | One of the function parameters is not valid.                                     |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_INVALID_LEVEL 0x0000007C     | The force level is invalid.                                                      |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_DEVICE_IN_USE 0x00002404     | The connection handle is in use and cannot be disconnected.                      |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_REDIR_PAUSED 0x00000048      | Remote access to the specified printer or serial communications device has been  |
	//	|                                    | paused.                                                                          |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//
	// The server SHOULD<50> enforce security measures to verify that the caller has the
	// required permissions to execute this routine. If the server enforces security measures
	// and the caller does not have the required credentials, the server MUST fail the call
	// with ERROR_ACCESS_DENIED. Specifications for determining the identity of the caller
	// for the purpose of performing an access check are in [MS-RPCE] section 3.3.3.1.3.
	//
	// The UseName parameter specifies the local device name or shared resource name for
	// which to delete a tree connection. The server MUST canonicalize UseName ([MS-SRVS]
	// section 3.1.4.33). This MUST be a nonempty, null-terminated UTF-16 string; otherwise,
	// the server MUST fail the call with an ERROR_INVALID_PARAMETER error code.
	UseDelete(context.Context, *UseDeleteRequest, ...dcerpc.CallOption) (*UseDeleteResponse, error)

	// The NetrUseEnum method lists open connections between a workstation server and a
	// remote SMB server. The server SHOULD NOT allow this method to be invoked remotely<51>
	// and SHOULD return ERROR_CALL_NOT_IMPLEMENTED.
	//
	// Return Values: The method returns 0x00000000 (NERR_Success) to indicate success;
	// otherwise, it returns a nonzero error code. The method can take any specific error
	// code value ([MS-ERREF] section 2.2). The most common error codes are listed in the
	// following table.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|                                    |                                                                                  |
	//	|             VALUE/CODE             |                                     MEANING                                      |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| NERR_Success 0x00000000            | The client request succeeded.                                                    |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_INVALID_LEVEL 0x0000007C     | The system call level is not correct.                                            |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_MORE_DATA 0x000000EA         | The client request succeeded. More entries are available. Not all entries could  |
	//	|                                    | be returned in the buffer size that is specified by PreferredMaximumLength.      |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_NOT_ENOUGH_MEMORY 0x00000008 | Not enough storage is available to process this command.                         |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| NERR_BufTooSmall 0x0000084B        | The client request succeeded. More entries are available. The buffer size that   |
	//	|                                    | is specified by PreferredMaximumLength was too small to fit even a single entry. |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//
	// The server SHOULD<52> enforce security measures to verify that the caller has the
	// required permissions to execute this routine. If the server enforces security measures
	// and the caller does not have the required credentials, the server MUST fail the call
	// with ERROR_ACCESS_DENIED. Specifications for determining the identity of the caller
	// for the purpose of performing an access check are in [MS-RPCE] section 3.3.3.1.3.
	//
	// The InfoStruct parameter has a Level member. The value of Level MUST be 0, 1, or
	// 2. If the Level member is not equal to one of the valid values, the server MUST fail
	// the call with an ERROR_INVALID_LEVEL error code.
	UseEnum(context.Context, *UseEnumRequest, ...dcerpc.CallOption) (*UseEnumResponse, error)

	// Opnum12NotUsedOnWire operation.
	// Opnum12NotUsedOnWire

	// The NetrWorkstationStatisticsGet method returns various statistics about the SMB
	// network redirector on a remote computer.
	//
	// Return Values: When the message processing result matches the description in column
	// two of the following table, this method MUST return one of the following values ([MS-ERREF]
	// section 2.2).
	//
	//	+------------------------------------+--------------------------------------------+
	//	|                                    |                                            |
	//	|             VALUE/CODE             |                  MEANING                   |
	//	|                                    |                                            |
	//	+------------------------------------+--------------------------------------------+
	//	+------------------------------------+--------------------------------------------+
	//	| NERR_Success 0x00000000            | The operation completed successfully.      |
	//	+------------------------------------+--------------------------------------------+
	//	| ERROR_ACCESS_DENIED 0x00000005     | Access is denied.                          |
	//	+------------------------------------+--------------------------------------------+
	//	| ERROR_INVALID_LEVEL 0x0000007C     | The information level is invalid.          |
	//	+------------------------------------+--------------------------------------------+
	//	| ERROR_INVALID_PARAMETER 0x00000057 | One of the function parameters is invalid. |
	//	+------------------------------------+--------------------------------------------+
	//
	// If the Level parameter does not equal 0x00000000, the server MUST fail the call and
	// return ERROR_INVALID_LEVEL.
	//
	// If the Options parameter does not equal 0x00000000, the server MUST fail the call
	// and return ERROR_INVALID_PARAMETER. The server SHOULD<54> enforce security measures
	// to verify that the caller has the required permissions to execute this routine. If
	// the server enforces security measures and the caller does not have the required credentials,
	// the server MUST fail the call and return ERROR_ACCESS_DENIED.
	WorkstationStatisticsGet(context.Context, *WorkstationStatisticsGetRequest, ...dcerpc.CallOption) (*WorkstationStatisticsGetResponse, error)

	// Opnum14NotUsedOnWire operation.
	// Opnum14NotUsedOnWire

	// Opnum15NotUsedOnWire operation.
	// Opnum15NotUsedOnWire

	// Opnum16NotUsedOnWire operation.
	// Opnum16NotUsedOnWire

	// Opnum17NotUsedOnWire operation.
	// Opnum17NotUsedOnWire

	// Opnum18NotUsedOnWire operation.
	// Opnum18NotUsedOnWire

	// Opnum19NotUsedOnWire operation.
	// Opnum19NotUsedOnWire

	// The NetrGetJoinInformation method retrieves details about the workgroup or domain
	// to which the specified computer is joined.
	//
	// Return Values: When the message processing result meets the description in column
	// two of the following table, this method MUST return one of the following values ([MS-ERREF]
	// section 2.2).
	//
	//	+----------------------------------------+---------------------------------------------+
	//	|                                        |                                             |
	//	|               VALUE/CODE               |                   MEANING                   |
	//	|                                        |                                             |
	//	+----------------------------------------+---------------------------------------------+
	//	+----------------------------------------+---------------------------------------------+
	//	| NERR_Success 0x00000000                | The operation completed successfully.       |
	//	+----------------------------------------+---------------------------------------------+
	//	| ERROR_ACCESS_DENIED 0x00000005         | Access is denied.                           |
	//	+----------------------------------------+---------------------------------------------+
	//	| RPC_S_PROTSEQ_NOT_SUPPORTED 0x000006A7 | The RPC protocol sequence is not supported. |
	//	+----------------------------------------+---------------------------------------------+
	//
	// Any other return value MUST conform to the error code requirements in Protocol Details
	// (section 3).
	GetJoinInformation(context.Context, *GetJoinInformationRequest, ...dcerpc.CallOption) (*GetJoinInformationResponse, error)

	// Opnum21NotUsedOnWire operation.
	// Opnum21NotUsedOnWire

	// The NetrJoinDomain2 method uses encrypted credentials to join a computer to a domain
	// or a workgroup.<57>
	//
	// For high-level, informative discussions about domain controller location and domain
	// join and unjoin, see [MS-ADOD] sections 2.7.7 and 3.1. For more information, see
	// the example in section 4.3.
	//
	// Return Values: When the message processing result meets the description in column
	// two of the following table, this method MUST return one of the following values ([MS-ERREF]
	// section 2.2).
	//
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	//	|                                        |                                                                                  |
	//	|               VALUE/CODE               |                                     MEANING                                      |
	//	|                                        |                                                                                  |
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	//	| NERR_Success 0x00000000                | The operation completed successfully.                                            |
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_FILE_NOT_FOUND 0x00000002        | The object was not found.                                                        |
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_ACCESS_DENIED 0x00000005         | Access is denied.                                                                |
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_NOT_SUPPORTED 0x00000032         | The request is not supported.                                                    |
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_INVALID_PASSWORD 0x00000056      | The specified network password is not correct.                                   |
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_INVALID_PARAMETER 0x00000057     | The parameter is incorrect.                                                      |
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_PASSWORD_RESTRICTION 0x0000052D  | Unable to update the password. The value provided for the new password does not  |
	//	|                                        | meet the length, complexity, or history requirements of the domain.              |
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_LOGON_FAILURE 0x0000052E         | Logon failure: unknown user name or bad password.                                |
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_NONE_MAPPED 0x00000534           | The account was not found.                                                       |
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_INVALID_DOMAIN_ROLE 0x0000054A   | The name of a domain controller was provided in the DomainNameParam parameter,   |
	//	|                                        | and validation of that domain controller failed. Validation is specified in the  |
	//	|                                        | message-processing steps for the section "Domain Join" later.                    |
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_NO_SUCH_DOMAIN 0x0000054B        | The specified domain either does not exist or could not be contacted.            |
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	//	| RPC_S_PROTSEQ_NOT_SUPPORTED 0x000006A7 | The RPC protocol sequence is not supported.                                      |
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	//	| RPC_S_CALL_IN_PROGRESS 0x000006FF      | A remote procedure call is already in progress.<66>                              |
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	//	| NERR_UserExists 0x000008B0             | The user account already exists.                                                 |
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	//	| NERR_SetupAlreadyJoined 0x00000A83     | This computer is already joined to a domain.                                     |
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	//	| NERR_SetupDomainController 0x00000A85  | This computer is a domain controller and cannot be unjoined from a domain.       |
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	//	| NERR_InvalidWorkgroupName 0x00000A87   | The specified workgroup name is invalid.                                         |
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	//
	// Any other return value MUST conform to the error code requirements in Protocol Details
	// (section 3).
	JoinDomain2(context.Context, *JoinDomain2Request, ...dcerpc.CallOption) (*JoinDomain2Response, error)

	// The NetrUnjoinDomain2 method uses encrypted credentials to unjoin a computer from
	// a workgroup or domain.<76>
	//
	// Return Values: When the message processing result meets the description in column
	// two of the following table, this method MUST return one of the following values ([MS-ERREF]
	// section 2.2).
	//
	//	+----------------------------------------+----------------------------------------------------------------------------+
	//	|                                        |                                                                            |
	//	|               VALUE/CODE               |                                  MEANING                                   |
	//	|                                        |                                                                            |
	//	+----------------------------------------+----------------------------------------------------------------------------+
	//	+----------------------------------------+----------------------------------------------------------------------------+
	//	| NERR_Success 0x00000000                | The operation completed successfully.                                      |
	//	+----------------------------------------+----------------------------------------------------------------------------+
	//	| ERROR_ACCESS_DENIED 0x00000005         | Access is denied.                                                          |
	//	+----------------------------------------+----------------------------------------------------------------------------+
	//	| ERROR_INVALID_PASSWORD 0x00000056      | The specified network password is not correct.                             |
	//	+----------------------------------------+----------------------------------------------------------------------------+
	//	| ERROR_INVALID_PARAMETER 0x00000057     | One of the function parameters is not valid.                               |
	//	+----------------------------------------+----------------------------------------------------------------------------+
	//	| ERROR_INVALID_FLAGS 0x000003EC         | Invalid option flags are specified.                                        |
	//	+----------------------------------------+----------------------------------------------------------------------------+
	//	| RPC_S_PROTSEQ_NOT_SUPPORTED 0x000006A7 | The RPC protocol sequence is not supported.                                |
	//	+----------------------------------------+----------------------------------------------------------------------------+
	//	| NERR_SetupNotJoined 0x00000A84         | This computer is not currently joined to a domain.                         |
	//	+----------------------------------------+----------------------------------------------------------------------------+
	//	| NERR_SetupDomainController 0x00000A85  | This computer is a domain controller and cannot be unjoined from a domain. |
	//	+----------------------------------------+----------------------------------------------------------------------------+
	//
	// Unless otherwise noted, if the server encounters an error during message processing,
	// the server SHOULD revert any state changes made, MUST stop message processing, and
	// MUST return the error to the caller.<78>
	UnjoinDomain2(context.Context, *UnjoinDomain2Request, ...dcerpc.CallOption) (*UnjoinDomain2Response, error)

	// The NetrRenameMachineInDomain2 method uses encrypted credentials to change the locally
	// persisted variable ComputerNameNetBIOS (section 3.2.1.5) and to optionally rename
	// the computer account for a server currently in a domain, without first removing the
	// computer from the domain and then adding it back.<81>
	//
	// Return Values: When the message processing result meets the description in column
	// two of the following table, this method MUST return one of the following values ([MS-ERREF]
	// section 2.2).
	//
	//	+----------------------------------------+-----------------------------------------------------------------+
	//	|                                        |                                                                 |
	//	|               VALUE/CODE               |                             MEANING                             |
	//	|                                        |                                                                 |
	//	+----------------------------------------+-----------------------------------------------------------------+
	//	+----------------------------------------+-----------------------------------------------------------------+
	//	| NERR_Success 0x00000000                | The operation completed successfully.                           |
	//	+----------------------------------------+-----------------------------------------------------------------+
	//	| ERROR_ACCESS_DENIED 0x00000005         | Access is denied.                                               |
	//	+----------------------------------------+-----------------------------------------------------------------+
	//	| ERROR_NOT_SUPPORTED 0x00000032         | The request is not supported.                                   |
	//	+----------------------------------------+-----------------------------------------------------------------+
	//	| ERROR_INVALID_PASSWORD 0x00000056      | The specified network password is not correct.                  |
	//	+----------------------------------------+-----------------------------------------------------------------+
	//	| ERROR_INVALID_PARAMETER 0x00000057     | The parameter is incorrect.                                     |
	//	+----------------------------------------+-----------------------------------------------------------------+
	//	| RPC_S_PROTSEQ_NOT_SUPPORTED 0x000006A7 | The RPC protocol sequence is not supported.                     |
	//	+----------------------------------------+-----------------------------------------------------------------+
	//	| NERR_SetupNotJoined 0x00000A84         | This computer is not currently joined to a domain.              |
	//	+----------------------------------------+-----------------------------------------------------------------+
	//	| NERR_SetupDomainController 0x00000A85  | This computer is a domain controller and cannot be renamed.<82> |
	//	+----------------------------------------+-----------------------------------------------------------------+
	//
	// Any other return value MUST conform to the error code requirements in Protocol Details
	// (section 3).
	//
	// Unless otherwise noted, if the server encounters an error during message processing,
	// the server SHOULD revert any state changes made, MUST stop message processing, and
	// MUST return the error to the caller.<83>
	RenameMachineInDomain2(context.Context, *RenameMachineInDomain2Request, ...dcerpc.CallOption) (*RenameMachineInDomain2Response, error)

	// The NetrValidateName2 method verifies the validity of a computer, workgroup, or domain
	// name (2).<87>
	//
	// Return Values: When the message processing result matches the description in column
	// 2 of the following table, this method MUST return one of the following values ([MS-ERREF]
	// section 2.2).
	//
	//	+----------------------------------------+---------------------------------------------------------------------------+
	//	|                                        |                                                                           |
	//	|               VALUE/CODE               |                                  MEANING                                  |
	//	|                                        |                                                                           |
	//	+----------------------------------------+---------------------------------------------------------------------------+
	//	+----------------------------------------+---------------------------------------------------------------------------+
	//	| NERR_Success 0x00000000                | The operation completed successfully.                                     |
	//	+----------------------------------------+---------------------------------------------------------------------------+
	//	| ERROR_ACCESS_DENIED 0x00000005         | Access is denied.                                                         |
	//	+----------------------------------------+---------------------------------------------------------------------------+
	//	| ERROR_DUP_NAME 0x00000034              | The connection was denied because a duplicate name exists on the network. |
	//	+----------------------------------------+---------------------------------------------------------------------------+
	//	| ERROR_INVALID_PASSWORD 0x00000056      | The specified network password is incorrect.                              |
	//	+----------------------------------------+---------------------------------------------------------------------------+
	//	| ERROR_INVALID_PARAMETER 0x00000057     | The parameter is incorrect.                                               |
	//	+----------------------------------------+---------------------------------------------------------------------------+
	//	| ERROR_INVALID_NAME 0x0000007B          | The file name, directory name, or volume label syntax is incorrect.       |
	//	+----------------------------------------+---------------------------------------------------------------------------+
	//	| ERROR_INVALID_DOMAINNAME 0x000004BC    | The format of the specified domain name (2) is invalid.                   |
	//	+----------------------------------------+---------------------------------------------------------------------------+
	//	| ERROR_NO_SUCH_DOMAIN 0x0000054B        | The specified domain either does not exist or could not be contacted.     |
	//	+----------------------------------------+---------------------------------------------------------------------------+
	//	| RPC_S_PROTSEQ_NOT_SUPPORTED 0x000006A7 | The RPC protocol sequence is not supported.                               |
	//	+----------------------------------------+---------------------------------------------------------------------------+
	//	| NERR_InvalidComputer 0x0000092F        | This computer name is invalid.                                            |
	//	+----------------------------------------+---------------------------------------------------------------------------+
	//	| NERR_InvalidWorkgroupName 0x00000A87   | The specified workgroup name is invalid.                                  |
	//	+----------------------------------------+---------------------------------------------------------------------------+
	//	| DNS_ERROR_NON_RFC_NAME 0x00002554      | The Internet host name does not comply with RFC specifications.           |
	//	+----------------------------------------+---------------------------------------------------------------------------+
	//	| DNS_ERROR_INVALID_NAME_CHAR 0x00002558 | The Internet host name contains an invalid character.                     |
	//	+----------------------------------------+---------------------------------------------------------------------------+
	//	| RPC_E_REMOTE_DISABLED 0x8001011C       | Remote calls are not allowed for this process.                            |
	//	+----------------------------------------+---------------------------------------------------------------------------+
	//
	// The following definition is used in the specification of message processing that
	// follows.
	//
	// * PasswordString : A Unicode ( 3acf0e02-9bbd-4ce0-a7a0-586bc72d3ef4#gt_c305d0ab-8b94-461a-bd76-13b40cb8c4d8
	// ) UTF-8 ( 3acf0e02-9bbd-4ce0-a7a0-586bc72d3ef4#gt_409411c4-b4ed-4ab6-b0ee-6d7815f85a35
	// ) string containing a password in cleartext ( 3acf0e02-9bbd-4ce0-a7a0-586bc72d3ef4#gt_f6e0fdd0-cbc1-4c9d-93b8-f25125f9c5ef
	// ).
	ValidateName2(context.Context, *ValidateName2Request, ...dcerpc.CallOption) (*ValidateName2Response, error)

	// The NetrGetJoinableOUs2 method returns a list of organizational units (OUs) in which
	// the user can create an object.<94>
	//
	// Return Values: When the message processing result matches the description in column
	// two of the following table, this method MUST return one of the following values ([MS-ERREF]
	// section 2.2).
	//
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	//	|                                        |                                                                                  |
	//	|               VALUE/CODE               |                                     MEANING                                      |
	//	|                                        |                                                                                  |
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	//	| NERR_Success 0x00000000                | The operation completed successfully.                                            |
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_ACCESS_DENIED 0x00000005         | Access is denied.                                                                |
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_NOT_ENOUGH_MEMORY 0x00000008     | Not enough storage is available to process this command.                         |
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_INVALID_PARAMETER 0x00000057     | A parameter is incorrect.<95>                                                    |
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	//	| RPC_S_PROTSEQ_NOT_SUPPORTED 0x000006A7 | The RPC protocol sequence is not supported.                                      |
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	//	| NERR_InvalidAPI 0x0000085E             | The requested API is not supported on domain controllers.                        |
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	//	| NERR_DefaultJoinRequired 0x00000A86    | The destination domain controller does not support creating machine accounts in  |
	//	|                                        | OUs.                                                                             |
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	GetJoinableOUs2(context.Context, *GetJoinableOUs2Request, ...dcerpc.CallOption) (*GetJoinableOUs2Response, error)

	// The NetrAddAlternateComputerName method adds an alternate name for a specified server
	// (2).<99>
	//
	// Return Values: When the message processing result matches the description in column
	// two of the following table, this method MUST return one of the following values ([MS-ERREF]
	// section 2.2).
	//
	//	+----------------------------------------+---------------------------------------------------------------------+
	//	|                                        |                                                                     |
	//	|               VALUE/CODE               |                               MEANING                               |
	//	|                                        |                                                                     |
	//	+----------------------------------------+---------------------------------------------------------------------+
	//	+----------------------------------------+---------------------------------------------------------------------+
	//	| NERR_Success 0x00000000                | The operation completed successfully.                               |
	//	+----------------------------------------+---------------------------------------------------------------------+
	//	| ERROR_ACCESS_DENIED 0x00000005         | Access is denied.                                                   |
	//	+----------------------------------------+---------------------------------------------------------------------+
	//	| ERROR_NOT_SUPPORTED 0x00000032         | This method is not supported by this server.                        |
	//	+----------------------------------------+---------------------------------------------------------------------+
	//	| ERROR_INVALID_PASSWORD 0x00000056      | The specified network password is incorrect.                        |
	//	+----------------------------------------+---------------------------------------------------------------------+
	//	| ERROR_INVALID_PARAMETER 0x00000057     | One of the function parameters is not valid.                        |
	//	+----------------------------------------+---------------------------------------------------------------------+
	//	| ERROR_INVALID_NAME 0x0000007B          | The file name, directory name, or volume label syntax is incorrect. |
	//	+----------------------------------------+---------------------------------------------------------------------+
	//	| ERROR_INVALID_FLAGS 0x000003EC         | Reserved contains an invalid value.                                 |
	//	+----------------------------------------+---------------------------------------------------------------------+
	//	| RPC_S_PROTSEQ_NOT_SUPPORTED 0x000006A7 | The RPC protocol sequence is not supported.                         |
	//	+----------------------------------------+---------------------------------------------------------------------+
	//	| RPC_S_CALL_IN_PROGRESS 0x000006FF      | A remote procedure call is already in progress.<101>                |
	//	+----------------------------------------+---------------------------------------------------------------------+
	//	| DNS_ERROR_INVALID_NAME_CHAR 0x00002558 | The Internet host name contains an invalid character.               |
	//	+----------------------------------------+---------------------------------------------------------------------+
	//
	// Unless otherwise noted, if the server encounters an error during message processing,
	// it SHOULD revert any state changes made, MUST stop message processing, and MUST return
	// the error to the caller.<102>
	AddAlternateComputerName(context.Context, *AddAlternateComputerNameRequest, ...dcerpc.CallOption) (*AddAlternateComputerNameResponse, error)

	// The NetrRemoveAlternateComputerName method removes an alternate name for a specified
	// server (2).<106>
	//
	// Return Values: When the message processing result matches the description in column
	// two of the following table, this method MUST return one of the following values ([MS-ERREF]
	// section 2.2).
	//
	//	+----------------------------------------+---------------------------------------------------------------------+
	//	|                                        |                                                                     |
	//	|               VALUE/CODE               |                               MEANING                               |
	//	|                                        |                                                                     |
	//	+----------------------------------------+---------------------------------------------------------------------+
	//	+----------------------------------------+---------------------------------------------------------------------+
	//	| NERR_Success 0x00000000                | The operation completed successfully.                               |
	//	+----------------------------------------+---------------------------------------------------------------------+
	//	| ERROR_ACCESS_DENIED 0x00000005         | Access is denied.                                                   |
	//	+----------------------------------------+---------------------------------------------------------------------+
	//	| ERROR_NOT_SUPPORTED 0x00000032         | This method is not supported by this server.                        |
	//	+----------------------------------------+---------------------------------------------------------------------+
	//	| ERROR_INVALID_PASSWORD 0x00000056      | The specified network password is not correct.                      |
	//	+----------------------------------------+---------------------------------------------------------------------+
	//	| ERROR_INVALID_PARAMETER 0x00000057     | One of the function parameters is not valid.                        |
	//	+----------------------------------------+---------------------------------------------------------------------+
	//	| ERROR_INVALID_NAME 0x0000007B          | An invalid name parameter is specified.                             |
	//	+----------------------------------------+---------------------------------------------------------------------+
	//	| ERROR_INVALID_FLAGS 0x000003EC         | The Reserved parameter contains an invalid value.                   |
	//	+----------------------------------------+---------------------------------------------------------------------+
	//	| ERROR_NOT_FOUND 0x00000490             | AlternateName was not found in the current list of alternate names. |
	//	+----------------------------------------+---------------------------------------------------------------------+
	//	| RPC_S_PROTSEQ_NOT_SUPPORTED 0x000006A7 | The RPC protocol sequence is not supported.                         |
	//	+----------------------------------------+---------------------------------------------------------------------+
	//	| RPC_S_CALL_IN_PROGRESS 0x000006FF      | A remote procedure call is already in progress.<108>                |
	//	+----------------------------------------+---------------------------------------------------------------------+
	//	| DNS_ERROR_INVALID_NAME_CHAR 0x00002558 | The Internet host name contains an invalid character.               |
	//	+----------------------------------------+---------------------------------------------------------------------+
	//
	// Any other return value MUST conform to the error code requirements in Protocol Details
	// (section 3).
	//
	// Unless otherwise noted, if the server encounters an error during message processing,
	// the server SHOULD revert any state changes made, MUST stop message processing, and
	// MUST return the error to the caller.<109>
	RemoveAlternateComputerName(context.Context, *RemoveAlternateComputerNameRequest, ...dcerpc.CallOption) (*RemoveAlternateComputerNameResponse, error)

	// The NetrSetPrimaryComputerName method sets the primary computer name for a specified
	// server (2).<113>
	//
	// Return Values: When the message processing result matches the description in column
	// two of the following table, this method MUST return one of the following values ([MS-ERREF]
	// section 2.2).
	//
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	//	|                                        |                                                                                  |
	//	|               VALUE/CODE               |                                     MEANING                                      |
	//	|                                        |                                                                                  |
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	//	| NERR_Success 0x00000000                | The operation completed successfully.                                            |
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_ACCESS_DENIED 0x00000005         | Access is denied.                                                                |
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_NOT_SUPPORTED 0x00000032         | This method is not supported by this server.                                     |
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_INVALID_PASSWORD 0x00000056      | The specified network password is incorrect.                                     |
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_INVALID_PARAMETER 0x00000057     | The parameter is incorrect.                                                      |
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x0000007B ERROR_INVALID_NAME          | An invalid name parameter is specified.                                          |
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_INVALID_FLAGS 0x000003EC         | Reserved contains an invalid value.                                              |
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	//	| RPC_S_PROTSEQ_NOT_SUPPORTED 0x000006A7 | The RPC protocol sequence is not supported.                                      |
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	//	| RPC_S_CALL_IN_PROGRESS 0x000006FF      | A remote procedure call is already in progress.<115>                             |
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	//	| NERR_DefaultJoinRequired 0x00000A86    | The destination domain controller does not support creating machine accounts in  |
	//	|                                        | OUs.                                                                             |
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	//	| DNS_ERROR_INVALID_NAME_CHAR 0x00002558 | The Internet host name contains an invalid character.                            |
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	//
	// Any other return value MUST conform to the error code requirements in Protocol Details
	// (section 3).
	//
	// Unless otherwise noted, if the server encounters an error during message processing,
	// the server SHOULD revert any state changes made, MUST stop message processing, and
	// MUST return the error to the caller.<116>
	SetPrimaryComputerName(context.Context, *SetPrimaryComputerNameRequest, ...dcerpc.CallOption) (*SetPrimaryComputerNameResponse, error)

	// The NetrEnumerateComputerNames method returns a list of computer names for a specified
	// server (2). The results of the query are determined by the type of the name. <120>
	//
	// Return Values: When the message processing result meets the description in column
	// two of the following table, this method MUST return one of the following values ([MS-ERREF]
	// section 2.2).
	//
	//	+----------------------------------------+----------------------------------------------------------+
	//	|                                        |                                                          |
	//	|               VALUE/CODE               |                         MEANING                          |
	//	|                                        |                                                          |
	//	+----------------------------------------+----------------------------------------------------------+
	//	+----------------------------------------+----------------------------------------------------------+
	//	| NERR_Success 0x00000000                | The operation completed successfully.                    |
	//	+----------------------------------------+----------------------------------------------------------+
	//	| ERROR_ACCESS_DENIED 0x00000005         | Access is denied.                                        |
	//	+----------------------------------------+----------------------------------------------------------+
	//	| ERROR_NOT_ENOUGH_MEMORY 0x00000008     | Not enough storage is available to process this command. |
	//	+----------------------------------------+----------------------------------------------------------+
	//	| ERROR_INVALID_PARAMETER 0x00000057     | The parameter is incorrect.                              |
	//	+----------------------------------------+----------------------------------------------------------+
	//	| ERROR_NOT_SUPPORTED 0x00000032         | This method is not supported by this server.             |
	//	+----------------------------------------+----------------------------------------------------------+
	//	| ERROR_INVALID_FLAGS 0x000003EC         | Reserved contains an invalid value.                      |
	//	+----------------------------------------+----------------------------------------------------------+
	//	| RPC_S_PROTSEQ_NOT_SUPPORTED 0x000006A7 | The RPC protocol sequence is not supported.              |
	//	+----------------------------------------+----------------------------------------------------------+
	//	| RPC_S_CALL_IN_PROGRESS 0x000006FF      | A remote procedure call is already in progress.          |
	//	+----------------------------------------+----------------------------------------------------------+
	//
	// Any other return value MUST conform to the error code requirements in Protocol Details
	// (section 3).
	EnumerateComputerNames(context.Context, *EnumerateComputerNamesRequest, ...dcerpc.CallOption) (*EnumerateComputerNamesResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn
}

// JoinObfuscatorLength represents the JOIN_OBFUSCATOR_LENGTH RPC constant
var JoinObfuscatorLength = 8

// JoinMaxPasswordLength represents the JOIN_MAX_PASSWORD_LENGTH RPC constant
var JoinMaxPasswordLength = 256

// NetsetupJoinStatus type represents NETSETUP_JOIN_STATUS RPC enumeration.
//
// The NETSETUP_JOIN_STATUS enumeration specifies details about the domain join status
// of a machine.
type NetsetupJoinStatus uint16

var (
	// NetSetupUnknownStatus: Domain join status of the machine is unknown.
	NetsetupJoinStatusSetupUnknownStatus NetsetupJoinStatus = 0
	// NetSetupUnjoined: Machine is not joined to a domain or to a workgroup.
	NetsetupJoinStatusSetupUnjoined NetsetupJoinStatus = 1
	// NetSetupWorkgroupName: Machine is joined to a workgroup.
	NetsetupJoinStatusSetupWorkgroupName NetsetupJoinStatus = 2
	// NetSetupDomainName: Machine is joined to a domain.
	NetsetupJoinStatusSetupDomainName NetsetupJoinStatus = 3
)

func (o NetsetupJoinStatus) String() string {
	switch o {
	case NetsetupJoinStatusSetupUnknownStatus:
		return "NetsetupJoinStatusSetupUnknownStatus"
	case NetsetupJoinStatusSetupUnjoined:
		return "NetsetupJoinStatusSetupUnjoined"
	case NetsetupJoinStatusSetupWorkgroupName:
		return "NetsetupJoinStatusSetupWorkgroupName"
	case NetsetupJoinStatusSetupDomainName:
		return "NetsetupJoinStatusSetupDomainName"
	}
	return "Invalid"
}

// NetsetupNameType type represents NETSETUP_NAME_TYPE RPC enumeration.
//
// The NETSETUP_NAME_TYPE enumeration specifies the types of validation that can be
// performed for a computer name, workgroup name, or domain name (2).
type NetsetupNameType uint16

var (
	// NetSetupUnknown: Reserved.
	NetsetupNameTypeSetupUnknown NetsetupNameType = 0
	// NetSetupMachine: Verify that the name is valid as a NetBIOS computer name and that
	// it is not in use.
	NetsetupNameTypeSetupMachine NetsetupNameType = 1
	// NetSetupWorkgroup: Verify that the name is valid as a workgroup name.
	NetsetupNameTypeSetupWorkgroup NetsetupNameType = 2
	// NetSetupDomain: Verify that the name is valid as a NetBIOS domain name and that a
	// domain with that name exists.
	NetsetupNameTypeSetupDomain NetsetupNameType = 3
	// NetSetupNonExistentDomain:  Verify that the name is valid as a NetBIOS domain name
	// and that a domain with that name does not exist.
	NetsetupNameTypeSetupNonExistentDomain NetsetupNameType = 4
	// NetSetupDnsMachine:  Verify that the name is valid as a DNS computer name.
	NetsetupNameTypeSetupDNSMachine NetsetupNameType = 5
)

func (o NetsetupNameType) String() string {
	switch o {
	case NetsetupNameTypeSetupUnknown:
		return "NetsetupNameTypeSetupUnknown"
	case NetsetupNameTypeSetupMachine:
		return "NetsetupNameTypeSetupMachine"
	case NetsetupNameTypeSetupWorkgroup:
		return "NetsetupNameTypeSetupWorkgroup"
	case NetsetupNameTypeSetupDomain:
		return "NetsetupNameTypeSetupDomain"
	case NetsetupNameTypeSetupNonExistentDomain:
		return "NetsetupNameTypeSetupNonExistentDomain"
	case NetsetupNameTypeSetupDNSMachine:
		return "NetsetupNameTypeSetupDNSMachine"
	}
	return "Invalid"
}

// NetComputerNameType type represents NET_COMPUTER_NAME_TYPE RPC enumeration.
//
// The NET_COMPUTER_NAME_TYPE enumeration specifies the types of names that can be enumerated
// for a computer using the NetrEnumerateComputerNames method (section 3.2.4.21).
type NetComputerNameType uint16

var (
	// NetPrimaryComputerName: Query the primary name of a computer.
	NetComputerNameTypePrimaryComputerName NetComputerNameType = 0
	// NetAlternateComputerNames: Query the alternate names of a computer.
	NetComputerNameTypeAlternateComputerNames NetComputerNameType = 1
	// NetAllComputerNames: Query all names of a computer.
	NetComputerNameTypeAllComputerNames NetComputerNameType = 2
	// NetComputerNameTypeMax: Maximum number of name types.
	NetComputerNameTypeMax NetComputerNameType = 3
)

func (o NetComputerNameType) String() string {
	switch o {
	case NetComputerNameTypePrimaryComputerName:
		return "NetComputerNameTypePrimaryComputerName"
	case NetComputerNameTypeAlternateComputerNames:
		return "NetComputerNameTypeAlternateComputerNames"
	case NetComputerNameTypeAllComputerNames:
		return "NetComputerNameTypeAllComputerNames"
	case NetComputerNameTypeMax:
		return "NetComputerNameTypeMax"
	}
	return "Invalid"
}

// StatWorkstation0 structure represents STAT_WORKSTATION_0 RPC structure.
//
// The STAT_WORKSTATION_0 structure contains statistical details about the SMB network
// redirector.
type StatWorkstation0 struct {
	// StatisticsStartTime: The time that statistics collection started. The value MUST
	// be stored as the number of seconds elapsed since 00:00:00, January 1, 1970 GMT.
	StatisticsStartTime *dtyp.LargeInteger `idl:"name:StatisticsStartTime" json:"statistics_start_time"`
	// BytesReceived: The total number of bytes the SMB network redirector has received.
	BytesReceived *dtyp.LargeInteger `idl:"name:BytesReceived" json:"bytes_received"`
	// SmbsReceived: The total number of SMB messages that the SMB network redirector has
	// received.
	SMBsReceived *dtyp.LargeInteger `idl:"name:SmbsReceived" json:"smbs_received"`
	// PagingReadBytesRequested: If applicable to the server (2), an implementation-specific
	// value; otherwise, it MUST be set to zero.
	PagingReadBytesRequested *dtyp.LargeInteger `idl:"name:PagingReadBytesRequested" json:"paging_read_bytes_requested"`
	// NonPagingReadBytesRequested: If applicable to the server, an implementation-specific
	// value; otherwise, it MUST be set to zero.
	NonPagingReadBytesRequested *dtyp.LargeInteger `idl:"name:NonPagingReadBytesRequested" json:"non_paging_read_bytes_requested"`
	// CacheReadBytesRequested: If applicable to the server, an implementation-specific
	// value; otherwise, it MUST be set to zero.
	CacheReadBytesRequested *dtyp.LargeInteger `idl:"name:CacheReadBytesRequested" json:"cache_read_bytes_requested"`
	// NetworkReadBytesRequested: If applicable to the server, an implementation-specific
	// value; otherwise, it MUST be set to zero.
	NetworkReadBytesRequested *dtyp.LargeInteger `idl:"name:NetworkReadBytesRequested" json:"network_read_bytes_requested"`
	// BytesTransmitted: The total number of bytes that the SMB network redirector has transmitted.
	BytesTransmitted *dtyp.LargeInteger `idl:"name:BytesTransmitted" json:"bytes_transmitted"`
	// SmbsTransmitted: The total number of SMB messages that the SMB network redirector
	// has transmitted.
	SMBsTransmitted *dtyp.LargeInteger `idl:"name:SmbsTransmitted" json:"smbs_transmitted"`
	// PagingWriteBytesRequested: If applicable to the server, an implementation-specific
	// value; otherwise, it MUST be set to zero.
	PagingWriteBytesRequested *dtyp.LargeInteger `idl:"name:PagingWriteBytesRequested" json:"paging_write_bytes_requested"`
	// NonPagingWriteBytesRequested: If applicable to the server, an implementation-specific
	// value; otherwise, it MUST be set to zero.
	NonPagingWriteBytesRequested *dtyp.LargeInteger `idl:"name:NonPagingWriteBytesRequested" json:"non_paging_write_bytes_requested"`
	// CacheWriteBytesRequested: If applicable to the server, an implementation-specific
	// value; otherwise, it MUST be set to zero.
	CacheWriteBytesRequested *dtyp.LargeInteger `idl:"name:CacheWriteBytesRequested" json:"cache_write_bytes_requested"`
	// NetworkWriteBytesRequested: If applicable to the server, an implementation-specific
	// value; otherwise, it MUST be set to zero.
	NetworkWriteBytesRequested *dtyp.LargeInteger `idl:"name:NetworkWriteBytesRequested" json:"network_write_bytes_requested"`
	// InitiallyFailedOperations: The total number of network operations that have failed
	// to start.
	InitiallyFailedOperations uint32 `idl:"name:InitiallyFailedOperations" json:"initially_failed_operations"`
	// FailedCompletionOperations: The total number of network operations that have failed
	// to complete.
	FailedCompletionOperations uint32 `idl:"name:FailedCompletionOperations" json:"failed_completion_operations"`
	// ReadOperations: The total number of read operations that the SMB network redirector
	// has initiated.
	ReadOperations uint32 `idl:"name:ReadOperations" json:"read_operations"`
	// RandomReadOperations: If applicable to the server, an implementation-specific value;
	// otherwise, it MUST be set to zero.
	RandomReadOperations uint32 `idl:"name:RandomReadOperations" json:"random_read_operations"`
	// ReadSmbs: The total number of read requests that the SMB network redirector has sent
	// to remote computers.
	ReadSMBs uint32 `idl:"name:ReadSmbs" json:"read_smbs"`
	// LargeReadSmbs: The total number of read requests greater than twice the size of the
	// remote computer’s negotiated buffer size that the SMB network redirector has sent
	// to remote computers.
	LargeReadSMBs uint32 `idl:"name:LargeReadSmbs" json:"large_read_smbs"`
	// SmallReadSmbs: The total number of read requests that are less than one-quarter the
	// size of the remote computer’s negotiated buffer size that the SMB network redirector
	// has sent to remote computers.
	SmallReadSMBs uint32 `idl:"name:SmallReadSmbs" json:"small_read_smbs"`
	// WriteOperations: The total number of write operations that the SMB network redirector
	// has initiated.
	WriteOperations uint32 `idl:"name:WriteOperations" json:"write_operations"`
	// RandomWriteOperations: If applicable to the server, an implementation-specific value;
	// otherwise, it MUST be set to zero.
	RandomWriteOperations uint32 `idl:"name:RandomWriteOperations" json:"random_write_operations"`
	// WriteSmbs: The total number of write requests that the SMB network redirector has
	// sent to remote computers.
	WriteSMBs uint32 `idl:"name:WriteSmbs" json:"write_smbs"`
	// LargeWriteSmbs: The total number of write requests that are greater than twice the
	// size of the remote computer’s negotiated buffer size and that the SMB network redirector
	// has sent to remote computers.
	LargeWriteSMBs uint32 `idl:"name:LargeWriteSmbs" json:"large_write_smbs"`
	// SmallWriteSmbs: The total number of write requests that are less than one-quarter
	// the size of the remote computer’s negotiated buffer size and that the SMB network
	// redirector has sent to remote computers ([MS-CIFS] section 3.2.4.15).
	SmallWriteSMBs uint32 `idl:"name:SmallWriteSmbs" json:"small_write_smbs"`
	// RawReadsDenied: The total number of raw read requests made by the SMB network redirector
	// that have been denied by the remote computer. This field MAY<8> be ignored.
	RawReadsDenied uint32 `idl:"name:RawReadsDenied" json:"raw_reads_denied"`
	// RawWritesDenied: The total number of raw write requests made by the SMB network redirector
	// that have been denied by the remote computer. This field MAY<9> be ignored.
	RawWritesDenied uint32 `idl:"name:RawWritesDenied" json:"raw_writes_denied"`
	// NetworkErrors: The total number of network errors that the SMB network redirector
	// has received.
	NetworkErrors uint32 `idl:"name:NetworkErrors" json:"network_errors"`
	// Sessions: The total number of remote SMB sessions that the SMB network redirector
	// has established.
	Sessions uint32 `idl:"name:Sessions" json:"sessions"`
	// FailedSessions: The number of times that the SMB network redirector has attempted
	// to create an SMB session but failed.
	FailedSessions uint32 `idl:"name:FailedSessions" json:"failed_sessions"`
	// Reconnects: The total number of SMB connections that have failed.
	Reconnects uint32 `idl:"name:Reconnects" json:"reconnects"`
	// CoreConnects: The total number of SMB connections to remote computers supporting
	// the PCNET1 dialect that have succeeded ([MS-CIFS] section 3.2.4.2.2).
	CoreConnects uint32 `idl:"name:CoreConnects" json:"core_connects"`
	// Lanman20Connects: The total number of SMB connections that have succeeded to remote
	// computers supporting the LM1.2X002 dialect.
	LANMAN20Connects uint32 `idl:"name:Lanman20Connects" json:"lanman20_connects"`
	// Lanman21Connects: The total number of SMB connections that have succeeded to remote
	// computers supporting the LANMAN2.1 dialect.
	LANMAN21Connects uint32 `idl:"name:Lanman21Connects" json:"lanman21_connects"`
	// LanmanNtConnects: The total number of SMB connections that have succeeded to remote
	// computers supporting the NTLANMAN dialect.
	LANMANNTConnects uint32 `idl:"name:LanmanNtConnects" json:"lanman_nt_connects"`
	// ServerDisconnects: The number of times that a remote computer has disconnected the
	// SMB network redirector.
	ServerDisconnects uint32 `idl:"name:ServerDisconnects" json:"server_disconnects"`
	// HungSessions: The total number of SMB sessions that have timed out due to lack of
	// response from the remote computer.
	HungSessions uint32 `idl:"name:HungSessions" json:"hung_sessions"`
	// UseCount: The total number of SMB connections that the SMB network redirector has
	// established.
	UseCount uint32 `idl:"name:UseCount" json:"use_count"`
	// FailedUseCount: The total number of failed SMB connections for the SMB network redirector.
	FailedUseCount uint32 `idl:"name:FailedUseCount" json:"failed_use_count"`
	// CurrentCommands: The number of current requests that the SMB network redirector has
	// completed.
	CurrentCommands uint32 `idl:"name:CurrentCommands" json:"current_commands"`
}

func (o *StatWorkstation0) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *StatWorkstation0) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(8); err != nil {
		return err
	}
	if o.StatisticsStartTime != nil {
		if err := o.StatisticsStartTime.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.LargeInteger{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.BytesReceived != nil {
		if err := o.BytesReceived.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.LargeInteger{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.SMBsReceived != nil {
		if err := o.SMBsReceived.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.LargeInteger{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.PagingReadBytesRequested != nil {
		if err := o.PagingReadBytesRequested.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.LargeInteger{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.NonPagingReadBytesRequested != nil {
		if err := o.NonPagingReadBytesRequested.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.LargeInteger{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.CacheReadBytesRequested != nil {
		if err := o.CacheReadBytesRequested.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.LargeInteger{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.NetworkReadBytesRequested != nil {
		if err := o.NetworkReadBytesRequested.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.LargeInteger{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.BytesTransmitted != nil {
		if err := o.BytesTransmitted.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.LargeInteger{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.SMBsTransmitted != nil {
		if err := o.SMBsTransmitted.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.LargeInteger{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.PagingWriteBytesRequested != nil {
		if err := o.PagingWriteBytesRequested.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.LargeInteger{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.NonPagingWriteBytesRequested != nil {
		if err := o.NonPagingWriteBytesRequested.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.LargeInteger{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.CacheWriteBytesRequested != nil {
		if err := o.CacheWriteBytesRequested.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.LargeInteger{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.NetworkWriteBytesRequested != nil {
		if err := o.NetworkWriteBytesRequested.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.LargeInteger{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.InitiallyFailedOperations); err != nil {
		return err
	}
	if err := w.WriteData(o.FailedCompletionOperations); err != nil {
		return err
	}
	if err := w.WriteData(o.ReadOperations); err != nil {
		return err
	}
	if err := w.WriteData(o.RandomReadOperations); err != nil {
		return err
	}
	if err := w.WriteData(o.ReadSMBs); err != nil {
		return err
	}
	if err := w.WriteData(o.LargeReadSMBs); err != nil {
		return err
	}
	if err := w.WriteData(o.SmallReadSMBs); err != nil {
		return err
	}
	if err := w.WriteData(o.WriteOperations); err != nil {
		return err
	}
	if err := w.WriteData(o.RandomWriteOperations); err != nil {
		return err
	}
	if err := w.WriteData(o.WriteSMBs); err != nil {
		return err
	}
	if err := w.WriteData(o.LargeWriteSMBs); err != nil {
		return err
	}
	if err := w.WriteData(o.SmallWriteSMBs); err != nil {
		return err
	}
	if err := w.WriteData(o.RawReadsDenied); err != nil {
		return err
	}
	if err := w.WriteData(o.RawWritesDenied); err != nil {
		return err
	}
	if err := w.WriteData(o.NetworkErrors); err != nil {
		return err
	}
	if err := w.WriteData(o.Sessions); err != nil {
		return err
	}
	if err := w.WriteData(o.FailedSessions); err != nil {
		return err
	}
	if err := w.WriteData(o.Reconnects); err != nil {
		return err
	}
	if err := w.WriteData(o.CoreConnects); err != nil {
		return err
	}
	if err := w.WriteData(o.LANMAN20Connects); err != nil {
		return err
	}
	if err := w.WriteData(o.LANMAN21Connects); err != nil {
		return err
	}
	if err := w.WriteData(o.LANMANNTConnects); err != nil {
		return err
	}
	if err := w.WriteData(o.ServerDisconnects); err != nil {
		return err
	}
	if err := w.WriteData(o.HungSessions); err != nil {
		return err
	}
	if err := w.WriteData(o.UseCount); err != nil {
		return err
	}
	if err := w.WriteData(o.FailedUseCount); err != nil {
		return err
	}
	if err := w.WriteData(o.CurrentCommands); err != nil {
		return err
	}
	return nil
}
func (o *StatWorkstation0) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(8); err != nil {
		return err
	}
	if o.StatisticsStartTime == nil {
		o.StatisticsStartTime = &dtyp.LargeInteger{}
	}
	if err := o.StatisticsStartTime.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.BytesReceived == nil {
		o.BytesReceived = &dtyp.LargeInteger{}
	}
	if err := o.BytesReceived.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.SMBsReceived == nil {
		o.SMBsReceived = &dtyp.LargeInteger{}
	}
	if err := o.SMBsReceived.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.PagingReadBytesRequested == nil {
		o.PagingReadBytesRequested = &dtyp.LargeInteger{}
	}
	if err := o.PagingReadBytesRequested.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.NonPagingReadBytesRequested == nil {
		o.NonPagingReadBytesRequested = &dtyp.LargeInteger{}
	}
	if err := o.NonPagingReadBytesRequested.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.CacheReadBytesRequested == nil {
		o.CacheReadBytesRequested = &dtyp.LargeInteger{}
	}
	if err := o.CacheReadBytesRequested.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.NetworkReadBytesRequested == nil {
		o.NetworkReadBytesRequested = &dtyp.LargeInteger{}
	}
	if err := o.NetworkReadBytesRequested.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.BytesTransmitted == nil {
		o.BytesTransmitted = &dtyp.LargeInteger{}
	}
	if err := o.BytesTransmitted.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.SMBsTransmitted == nil {
		o.SMBsTransmitted = &dtyp.LargeInteger{}
	}
	if err := o.SMBsTransmitted.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.PagingWriteBytesRequested == nil {
		o.PagingWriteBytesRequested = &dtyp.LargeInteger{}
	}
	if err := o.PagingWriteBytesRequested.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.NonPagingWriteBytesRequested == nil {
		o.NonPagingWriteBytesRequested = &dtyp.LargeInteger{}
	}
	if err := o.NonPagingWriteBytesRequested.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.CacheWriteBytesRequested == nil {
		o.CacheWriteBytesRequested = &dtyp.LargeInteger{}
	}
	if err := o.CacheWriteBytesRequested.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.NetworkWriteBytesRequested == nil {
		o.NetworkWriteBytesRequested = &dtyp.LargeInteger{}
	}
	if err := o.NetworkWriteBytesRequested.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.InitiallyFailedOperations); err != nil {
		return err
	}
	if err := w.ReadData(&o.FailedCompletionOperations); err != nil {
		return err
	}
	if err := w.ReadData(&o.ReadOperations); err != nil {
		return err
	}
	if err := w.ReadData(&o.RandomReadOperations); err != nil {
		return err
	}
	if err := w.ReadData(&o.ReadSMBs); err != nil {
		return err
	}
	if err := w.ReadData(&o.LargeReadSMBs); err != nil {
		return err
	}
	if err := w.ReadData(&o.SmallReadSMBs); err != nil {
		return err
	}
	if err := w.ReadData(&o.WriteOperations); err != nil {
		return err
	}
	if err := w.ReadData(&o.RandomWriteOperations); err != nil {
		return err
	}
	if err := w.ReadData(&o.WriteSMBs); err != nil {
		return err
	}
	if err := w.ReadData(&o.LargeWriteSMBs); err != nil {
		return err
	}
	if err := w.ReadData(&o.SmallWriteSMBs); err != nil {
		return err
	}
	if err := w.ReadData(&o.RawReadsDenied); err != nil {
		return err
	}
	if err := w.ReadData(&o.RawWritesDenied); err != nil {
		return err
	}
	if err := w.ReadData(&o.NetworkErrors); err != nil {
		return err
	}
	if err := w.ReadData(&o.Sessions); err != nil {
		return err
	}
	if err := w.ReadData(&o.FailedSessions); err != nil {
		return err
	}
	if err := w.ReadData(&o.Reconnects); err != nil {
		return err
	}
	if err := w.ReadData(&o.CoreConnects); err != nil {
		return err
	}
	if err := w.ReadData(&o.LANMAN20Connects); err != nil {
		return err
	}
	if err := w.ReadData(&o.LANMAN21Connects); err != nil {
		return err
	}
	if err := w.ReadData(&o.LANMANNTConnects); err != nil {
		return err
	}
	if err := w.ReadData(&o.ServerDisconnects); err != nil {
		return err
	}
	if err := w.ReadData(&o.HungSessions); err != nil {
		return err
	}
	if err := w.ReadData(&o.UseCount); err != nil {
		return err
	}
	if err := w.ReadData(&o.FailedUseCount); err != nil {
		return err
	}
	if err := w.ReadData(&o.CurrentCommands); err != nil {
		return err
	}
	return nil
}

// WorkstationInfo100 structure represents WKSTA_INFO_100 RPC structure.
//
// The WKSTA_INFO_100 structure specifies details about a computer environment, including
// platform-specific information, the names of the domain and local computer, and the
// operating system.
type WorkstationInfo100 struct {
	// wki100_platform_id: The type of operating system. This MUST be one of the following
	// values.
	//
	//	+------------+------------------------+
	//	|            |                        |
	//	|   VALUE    |        MEANING         |
	//	|            |                        |
	//	+------------+------------------------+
	//	+------------+------------------------+
	//	| 0x0000012C | DOS (300 decimal)      |
	//	+------------+------------------------+
	//	| 0x00000190 | OS2 (400 decimal)      |
	//	+------------+------------------------+
	//	| 0x000001F4 | Windows: (500 decimal) |
	//	+------------+------------------------+
	//	| 0x00000258 | OSF: (600 decimal)     |
	//	+------------+------------------------+
	//	| 0x000002BC | VMS: (700 decimal)     |
	//	+------------+------------------------+
	PlatformID uint32 `idl:"name:wki100_platform_id" json:"platform_id"`
	// wki100_computername: A null-terminated, Internet host name or NetBIOS name [RFC1001]
	// of the local computer.
	ComputerName string `idl:"name:wki100_computername;string" json:"computer_name"`
	// wki100_langroup: A null-terminated, fully qualified domain name (FQDN) (2) of the
	// domain to which the computer belongs.
	LANGroup string `idl:"name:wki100_langroup;string" json:"lan_group"`
	// wki100_ver_major: The major version number of the operating system running on the
	// computer.
	VerMajor uint32 `idl:"name:wki100_ver_major" json:"ver_major"`
	// wki100_ver_minor: The minor version number of the operating system running on the
	// computer.
	VerMinor uint32 `idl:"name:wki100_ver_minor" json:"ver_minor"`
}

func (o *WorkstationInfo100) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *WorkstationInfo100) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.PlatformID); err != nil {
		return err
	}
	if o.ComputerName != "" {
		_ptr_wki100_computername := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.ComputerName); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.ComputerName, _ptr_wki100_computername); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.LANGroup != "" {
		_ptr_wki100_langroup := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.LANGroup); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.LANGroup, _ptr_wki100_langroup); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.VerMajor); err != nil {
		return err
	}
	if err := w.WriteData(o.VerMinor); err != nil {
		return err
	}
	return nil
}
func (o *WorkstationInfo100) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.PlatformID); err != nil {
		return err
	}
	_ptr_wki100_computername := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.ComputerName); err != nil {
			return err
		}
		return nil
	})
	_s_wki100_computername := func(ptr interface{}) { o.ComputerName = *ptr.(*string) }
	if err := w.ReadPointer(&o.ComputerName, _s_wki100_computername, _ptr_wki100_computername); err != nil {
		return err
	}
	_ptr_wki100_langroup := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.LANGroup); err != nil {
			return err
		}
		return nil
	})
	_s_wki100_langroup := func(ptr interface{}) { o.LANGroup = *ptr.(*string) }
	if err := w.ReadPointer(&o.LANGroup, _s_wki100_langroup, _ptr_wki100_langroup); err != nil {
		return err
	}
	if err := w.ReadData(&o.VerMajor); err != nil {
		return err
	}
	if err := w.ReadData(&o.VerMinor); err != nil {
		return err
	}
	return nil
}

// WorkstationInfo101 structure represents WKSTA_INFO_101 RPC structure.
//
// The WKSTA_INFO_101 structure specifies details about a computer environment, including
// platform-specific information, the names of the domain and local computer, and the
// operating system.
type WorkstationInfo101 struct {
	// wki101_platform_id: The type of operating system (section 2.2.5.1).
	PlatformID uint32 `idl:"name:wki101_platform_id" json:"platform_id"`
	// wki101_computername: A null-terminated, Internet host name or NetBIOS name [RFC1001]
	// of the local computer.
	ComputerName string `idl:"name:wki101_computername;string" json:"computer_name"`
	// wki101_langroup: A null-terminated, fully qualified domain name (FQDN) (2) of the
	// domain to which the computer belongs.
	LANGroup string `idl:"name:wki101_langroup;string" json:"lan_group"`
	// wki101_ver_major: The major version number of the operating system running on the
	// computer.
	VerMajor uint32 `idl:"name:wki101_ver_major" json:"ver_major"`
	// wki101_ver_minor: The minor version number of the operating system running on the
	// computer.
	VerMinor uint32 `idl:"name:wki101_ver_minor" json:"ver_minor"`
	// wki101_lanroot: A value that is not used and MUST be returned as NULL by the server
	// (2).
	LANRoot string `idl:"name:wki101_lanroot;string" json:"lan_root"`
}

func (o *WorkstationInfo101) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *WorkstationInfo101) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.PlatformID); err != nil {
		return err
	}
	if o.ComputerName != "" {
		_ptr_wki101_computername := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.ComputerName); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.ComputerName, _ptr_wki101_computername); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.LANGroup != "" {
		_ptr_wki101_langroup := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.LANGroup); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.LANGroup, _ptr_wki101_langroup); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.VerMajor); err != nil {
		return err
	}
	if err := w.WriteData(o.VerMinor); err != nil {
		return err
	}
	if o.LANRoot != "" {
		_ptr_wki101_lanroot := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.LANRoot); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.LANRoot, _ptr_wki101_lanroot); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *WorkstationInfo101) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.PlatformID); err != nil {
		return err
	}
	_ptr_wki101_computername := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.ComputerName); err != nil {
			return err
		}
		return nil
	})
	_s_wki101_computername := func(ptr interface{}) { o.ComputerName = *ptr.(*string) }
	if err := w.ReadPointer(&o.ComputerName, _s_wki101_computername, _ptr_wki101_computername); err != nil {
		return err
	}
	_ptr_wki101_langroup := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.LANGroup); err != nil {
			return err
		}
		return nil
	})
	_s_wki101_langroup := func(ptr interface{}) { o.LANGroup = *ptr.(*string) }
	if err := w.ReadPointer(&o.LANGroup, _s_wki101_langroup, _ptr_wki101_langroup); err != nil {
		return err
	}
	if err := w.ReadData(&o.VerMajor); err != nil {
		return err
	}
	if err := w.ReadData(&o.VerMinor); err != nil {
		return err
	}
	_ptr_wki101_lanroot := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.LANRoot); err != nil {
			return err
		}
		return nil
	})
	_s_wki101_lanroot := func(ptr interface{}) { o.LANRoot = *ptr.(*string) }
	if err := w.ReadPointer(&o.LANRoot, _s_wki101_lanroot, _ptr_wki101_lanroot); err != nil {
		return err
	}
	return nil
}

// WorkstationInfo102 structure represents WKSTA_INFO_102 RPC structure.
//
// The WKSTA_INFO_102 structure specifies details about a computer environment, including
// platform-specific information, the names of the domain and local computer, the operating
// system, and the logged-on users.
type WorkstationInfo102 struct {
	// wki102_platform_id: The type of operating system (section 2.2.5.1).
	PlatformID uint32 `idl:"name:wki102_platform_id" json:"platform_id"`
	// wki102_computername: A null-terminated, Internet host name or NetBIOS name [RFC1001]
	// of the local computer.
	ComputerName string `idl:"name:wki102_computername;string" json:"computer_name"`
	// wki102_langroup: A null-terminated, fully qualified domain name (FQDN) (2) of the
	// domain to which the computer belongs.
	LANGroup string `idl:"name:wki102_langroup;string" json:"lan_group"`
	// wki102_ver_major: The major version number of the operating system running on the
	// computer.
	VerMajor uint32 `idl:"name:wki102_ver_major" json:"ver_major"`
	// wki102_ver_minor: The minor version number of the operating system running on the
	// computer.
	VerMinor uint32 `idl:"name:wki102_ver_minor" json:"ver_minor"`
	// wki102_lanroot: A value that is not used and MUST be returned as NULL by the server
	// (2).
	LANRoot string `idl:"name:wki102_lanroot;string" json:"lan_root"`
	// wki102_logged_on_users: The number of users who are currently active on the computer.
	LoggedOnUsers uint32 `idl:"name:wki102_logged_on_users" json:"logged_on_users"`
}

func (o *WorkstationInfo102) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *WorkstationInfo102) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.PlatformID); err != nil {
		return err
	}
	if o.ComputerName != "" {
		_ptr_wki102_computername := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.ComputerName); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.ComputerName, _ptr_wki102_computername); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.LANGroup != "" {
		_ptr_wki102_langroup := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.LANGroup); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.LANGroup, _ptr_wki102_langroup); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.VerMajor); err != nil {
		return err
	}
	if err := w.WriteData(o.VerMinor); err != nil {
		return err
	}
	if o.LANRoot != "" {
		_ptr_wki102_lanroot := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.LANRoot); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.LANRoot, _ptr_wki102_lanroot); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.LoggedOnUsers); err != nil {
		return err
	}
	return nil
}
func (o *WorkstationInfo102) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.PlatformID); err != nil {
		return err
	}
	_ptr_wki102_computername := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.ComputerName); err != nil {
			return err
		}
		return nil
	})
	_s_wki102_computername := func(ptr interface{}) { o.ComputerName = *ptr.(*string) }
	if err := w.ReadPointer(&o.ComputerName, _s_wki102_computername, _ptr_wki102_computername); err != nil {
		return err
	}
	_ptr_wki102_langroup := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.LANGroup); err != nil {
			return err
		}
		return nil
	})
	_s_wki102_langroup := func(ptr interface{}) { o.LANGroup = *ptr.(*string) }
	if err := w.ReadPointer(&o.LANGroup, _s_wki102_langroup, _ptr_wki102_langroup); err != nil {
		return err
	}
	if err := w.ReadData(&o.VerMajor); err != nil {
		return err
	}
	if err := w.ReadData(&o.VerMinor); err != nil {
		return err
	}
	_ptr_wki102_lanroot := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.LANRoot); err != nil {
			return err
		}
		return nil
	})
	_s_wki102_lanroot := func(ptr interface{}) { o.LANRoot = *ptr.(*string) }
	if err := w.ReadPointer(&o.LANRoot, _s_wki102_lanroot, _ptr_wki102_lanroot); err != nil {
		return err
	}
	if err := w.ReadData(&o.LoggedOnUsers); err != nil {
		return err
	}
	return nil
}

// WorkstationInfo502 structure represents WKSTA_INFO_502 RPC structure.
//
// The WKSTA_INFO_502 structure specifies details about a computer environment.
type WorkstationInfo502 struct {
	// wki502_char_wait: Any value when sent, and MUST be ignored on receipt.
	CharWait uint32 `idl:"name:wki502_char_wait" json:"char_wait"`
	// wki502_collection_time: Any value when sent, and MUST be ignored on receipt.
	CollectionTime uint32 `idl:"name:wki502_collection_time" json:"collection_time"`
	// wki502_maximum_collection_count: Any value when sent, and MUST be ignored on receipt.
	MaximumCollectionCount uint32 `idl:"name:wki502_maximum_collection_count" json:"maximum_collection_count"`
	// wki502_keep_conn: The number of seconds the SMB network redirector maintains an inactive
	// SMB connection to a remote computer’s resource before closing it.
	KeepConn uint32 `idl:"name:wki502_keep_conn" json:"keep_conn"`
	// wki502_max_cmds: The number of simultaneous network commands that can be sent to
	// the SMB network redirector.
	MaxCmds uint32 `idl:"name:wki502_max_cmds" json:"max_cmds"`
	// wki502_sess_timeout: The number of seconds the server (2) waits before disconnecting
	// an inactive session.
	SessTimeout uint32 `idl:"name:wki502_sess_timeout" json:"sess_timeout"`
	// wki502_siz_char_buf: Any value when sent, and MUST be ignored on receipt.
	SizCharBuffer uint32 `idl:"name:wki502_siz_char_buf" json:"siz_char_buffer"`
	// wki502_max_threads: Any value when sent, and MUST be ignored on receipt.
	MaxThreads uint32 `idl:"name:wki502_max_threads" json:"max_threads"`
	// wki502_lock_quota: Any value when sent, and MUST be ignored on receipt.
	LockQuota uint32 `idl:"name:wki502_lock_quota" json:"lock_quota"`
	// wki502_lock_increment: Any value when sent, and MUST be ignored on receipt.
	LockIncrement uint32 `idl:"name:wki502_lock_increment" json:"lock_increment"`
	// wki502_lock_maximum: Any value when sent, and MUST be ignored on receipt.
	LockMaximum uint32 `idl:"name:wki502_lock_maximum" json:"lock_maximum"`
	// wki502_pipe_increment: Any value when sent, and MUST be ignored on receipt.
	PipeIncrement uint32 `idl:"name:wki502_pipe_increment" json:"pipe_increment"`
	// wki502_pipe_maximum: Any value when sent, and MUST be ignored on receipt.
	PipeMaximum uint32 `idl:"name:wki502_pipe_maximum" json:"pipe_maximum"`
	// wki502_cache_file_timeout:  Any value when sent, and MUST be ignored on receipt.
	CacheFileTimeout uint32 `idl:"name:wki502_cache_file_timeout" json:"cache_file_timeout"`
	// wki502_dormant_file_limit: The maximum number of file or printer handles the SMB
	// network redirector can continue to keep open, even after the application has closed
	// the corresponding handle.
	DormantFileLimit uint32 `idl:"name:wki502_dormant_file_limit" json:"dormant_file_limit"`
	// wki502_read_ahead_throughput: Any value when sent, and MUST be ignored on receipt.
	ReadAheadThroughput uint32 `idl:"name:wki502_read_ahead_throughput" json:"read_ahead_throughput"`
	// wki502_num_mailslot_buffers: Any value when sent, and MUST be ignored on receipt.
	NumMailslotBuffers uint32 `idl:"name:wki502_num_mailslot_buffers" json:"num_mailslot_buffers"`
	// wki502_num_srv_announce_buffers: Any value when sent, and MUST be ignored on receipt.
	NumServerAnnounceBuffers uint32 `idl:"name:wki502_num_srv_announce_buffers" json:"num_server_announce_buffers"`
	// wki502_max_illegal_datagram_events: Any value when sent, and MUST be ignored on receipt.
	MaxIllegalDatagramEvents uint32 `idl:"name:wki502_max_illegal_datagram_events" json:"max_illegal_datagram_events"`
	// wki502_illegal_datagram_event_reset_frequency: Any value when sent, and MUST be ignored
	// on receipt.
	IllegalDatagramEventResetFrequency uint32 `idl:"name:wki502_illegal_datagram_event_reset_frequency" json:"illegal_datagram_event_reset_frequency"`
	// wki502_log_election_packets: Any value when sent, and MUST be ignored on receipt.
	LogElectionPackets int32 `idl:"name:wki502_log_election_packets" json:"log_election_packets"`
	// wki502_use_opportunistic_locking: Any value when sent, and MUST be ignored on receipt.
	UseOpportunisticLocking int32 `idl:"name:wki502_use_opportunistic_locking" json:"use_opportunistic_locking"`
	// wki502_use_unlock_behind: Any value when sent, and MUST be ignored on receipt.
	UseUnlockBehind int32 `idl:"name:wki502_use_unlock_behind" json:"use_unlock_behind"`
	// wki502_use_close_behind: Any value when sent, and MUST be ignored on receipt.
	UseCloseBehind int32 `idl:"name:wki502_use_close_behind" json:"use_close_behind"`
	// wki502_buf_named_pipes: Any value when sent, and MUST be ignored on receipt.
	BufferNamedPipes int32 `idl:"name:wki502_buf_named_pipes" json:"buffer_named_pipes"`
	// wki502_use_lock_read_unlock: Any value when sent, and MUST be ignored on receipt.
	UseLockReadUnlock int32 `idl:"name:wki502_use_lock_read_unlock" json:"use_lock_read_unlock"`
	// wki502_utilize_nt_caching: Any value when sent, and MUST be ignored on receipt.
	UtilizeNTCaching int32 `idl:"name:wki502_utilize_nt_caching" json:"utilize_nt_caching"`
	// wki502_use_raw_read: Any value when sent, and MUST be ignored on receipt.
	UseRawRead int32 `idl:"name:wki502_use_raw_read" json:"use_raw_read"`
	// wki502_use_raw_write:  Any value when sent, and MUST be ignored on receipt.
	UseRawWrite int32 `idl:"name:wki502_use_raw_write" json:"use_raw_write"`
	// wki502_use_write_raw_data: Any value when sent, and MUST be ignored on receipt.
	UseWriteRawData int32 `idl:"name:wki502_use_write_raw_data" json:"use_write_raw_data"`
	// wki502_use_encryption: Any value when sent, and MUST be ignored on receipt.
	UseEncryption int32 `idl:"name:wki502_use_encryption" json:"use_encryption"`
	// wki502_buf_files_deny_write: Any value when sent, and MUST be ignored on receipt.
	BufferFilesDenyWrite int32 `idl:"name:wki502_buf_files_deny_write" json:"buffer_files_deny_write"`
	// wki502_buf_read_only_files: Any value when sent, and MUST be ignored on receipt.
	BufferReadOnlyFiles int32 `idl:"name:wki502_buf_read_only_files" json:"buffer_read_only_files"`
	// wki502_force_core_create_mode: Any value when sent, and MUST be ignored on receipt.
	ForceCoreCreateMode int32 `idl:"name:wki502_force_core_create_mode" json:"force_core_create_mode"`
	// wki502_use_512_byte_max_transfer: Any value when sent, and MUST be ignored on receipt.
	//
	// The wki502_keep_conn, wki502_max_cmds, wki502_sess_timeout, and wki502_dormant_file_limit
	// are the only values the server (2) can use to configure the redirector. The server
	// MUST store all the values and return the existing values upon a client’s request.
	Use512ByteMaxTransfer int32 `idl:"name:wki502_use_512_byte_max_transfer" json:"use_512_byte_max_transfer"`
}

func (o *WorkstationInfo502) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *WorkstationInfo502) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.CharWait); err != nil {
		return err
	}
	if err := w.WriteData(o.CollectionTime); err != nil {
		return err
	}
	if err := w.WriteData(o.MaximumCollectionCount); err != nil {
		return err
	}
	if err := w.WriteData(o.KeepConn); err != nil {
		return err
	}
	if err := w.WriteData(o.MaxCmds); err != nil {
		return err
	}
	if err := w.WriteData(o.SessTimeout); err != nil {
		return err
	}
	if err := w.WriteData(o.SizCharBuffer); err != nil {
		return err
	}
	if err := w.WriteData(o.MaxThreads); err != nil {
		return err
	}
	if err := w.WriteData(o.LockQuota); err != nil {
		return err
	}
	if err := w.WriteData(o.LockIncrement); err != nil {
		return err
	}
	if err := w.WriteData(o.LockMaximum); err != nil {
		return err
	}
	if err := w.WriteData(o.PipeIncrement); err != nil {
		return err
	}
	if err := w.WriteData(o.PipeMaximum); err != nil {
		return err
	}
	if err := w.WriteData(o.CacheFileTimeout); err != nil {
		return err
	}
	if err := w.WriteData(o.DormantFileLimit); err != nil {
		return err
	}
	if err := w.WriteData(o.ReadAheadThroughput); err != nil {
		return err
	}
	if err := w.WriteData(o.NumMailslotBuffers); err != nil {
		return err
	}
	if err := w.WriteData(o.NumServerAnnounceBuffers); err != nil {
		return err
	}
	if err := w.WriteData(o.MaxIllegalDatagramEvents); err != nil {
		return err
	}
	if err := w.WriteData(o.IllegalDatagramEventResetFrequency); err != nil {
		return err
	}
	if err := w.WriteData(o.LogElectionPackets); err != nil {
		return err
	}
	if err := w.WriteData(o.UseOpportunisticLocking); err != nil {
		return err
	}
	if err := w.WriteData(o.UseUnlockBehind); err != nil {
		return err
	}
	if err := w.WriteData(o.UseCloseBehind); err != nil {
		return err
	}
	if err := w.WriteData(o.BufferNamedPipes); err != nil {
		return err
	}
	if err := w.WriteData(o.UseLockReadUnlock); err != nil {
		return err
	}
	if err := w.WriteData(o.UtilizeNTCaching); err != nil {
		return err
	}
	if err := w.WriteData(o.UseRawRead); err != nil {
		return err
	}
	if err := w.WriteData(o.UseRawWrite); err != nil {
		return err
	}
	if err := w.WriteData(o.UseWriteRawData); err != nil {
		return err
	}
	if err := w.WriteData(o.UseEncryption); err != nil {
		return err
	}
	if err := w.WriteData(o.BufferFilesDenyWrite); err != nil {
		return err
	}
	if err := w.WriteData(o.BufferReadOnlyFiles); err != nil {
		return err
	}
	if err := w.WriteData(o.ForceCoreCreateMode); err != nil {
		return err
	}
	if err := w.WriteData(o.Use512ByteMaxTransfer); err != nil {
		return err
	}
	return nil
}
func (o *WorkstationInfo502) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.CharWait); err != nil {
		return err
	}
	if err := w.ReadData(&o.CollectionTime); err != nil {
		return err
	}
	if err := w.ReadData(&o.MaximumCollectionCount); err != nil {
		return err
	}
	if err := w.ReadData(&o.KeepConn); err != nil {
		return err
	}
	if err := w.ReadData(&o.MaxCmds); err != nil {
		return err
	}
	if err := w.ReadData(&o.SessTimeout); err != nil {
		return err
	}
	if err := w.ReadData(&o.SizCharBuffer); err != nil {
		return err
	}
	if err := w.ReadData(&o.MaxThreads); err != nil {
		return err
	}
	if err := w.ReadData(&o.LockQuota); err != nil {
		return err
	}
	if err := w.ReadData(&o.LockIncrement); err != nil {
		return err
	}
	if err := w.ReadData(&o.LockMaximum); err != nil {
		return err
	}
	if err := w.ReadData(&o.PipeIncrement); err != nil {
		return err
	}
	if err := w.ReadData(&o.PipeMaximum); err != nil {
		return err
	}
	if err := w.ReadData(&o.CacheFileTimeout); err != nil {
		return err
	}
	if err := w.ReadData(&o.DormantFileLimit); err != nil {
		return err
	}
	if err := w.ReadData(&o.ReadAheadThroughput); err != nil {
		return err
	}
	if err := w.ReadData(&o.NumMailslotBuffers); err != nil {
		return err
	}
	if err := w.ReadData(&o.NumServerAnnounceBuffers); err != nil {
		return err
	}
	if err := w.ReadData(&o.MaxIllegalDatagramEvents); err != nil {
		return err
	}
	if err := w.ReadData(&o.IllegalDatagramEventResetFrequency); err != nil {
		return err
	}
	if err := w.ReadData(&o.LogElectionPackets); err != nil {
		return err
	}
	if err := w.ReadData(&o.UseOpportunisticLocking); err != nil {
		return err
	}
	if err := w.ReadData(&o.UseUnlockBehind); err != nil {
		return err
	}
	if err := w.ReadData(&o.UseCloseBehind); err != nil {
		return err
	}
	if err := w.ReadData(&o.BufferNamedPipes); err != nil {
		return err
	}
	if err := w.ReadData(&o.UseLockReadUnlock); err != nil {
		return err
	}
	if err := w.ReadData(&o.UtilizeNTCaching); err != nil {
		return err
	}
	if err := w.ReadData(&o.UseRawRead); err != nil {
		return err
	}
	if err := w.ReadData(&o.UseRawWrite); err != nil {
		return err
	}
	if err := w.ReadData(&o.UseWriteRawData); err != nil {
		return err
	}
	if err := w.ReadData(&o.UseEncryption); err != nil {
		return err
	}
	if err := w.ReadData(&o.BufferFilesDenyWrite); err != nil {
		return err
	}
	if err := w.ReadData(&o.BufferReadOnlyFiles); err != nil {
		return err
	}
	if err := w.ReadData(&o.ForceCoreCreateMode); err != nil {
		return err
	}
	if err := w.ReadData(&o.Use512ByteMaxTransfer); err != nil {
		return err
	}
	return nil
}

// WorkstationInfo1013 structure represents WKSTA_INFO_1013 RPC structure.
//
// The WKSTA_INFO_1013 structure specifies details about the state of the SMB network
// redirector.
type WorkstationInfo1013 struct {
	// wki1013_keep_conn: The number of seconds the SMB network redirector maintains an
	// inactive SMB connection to a remote computer’s resource before closing it.
	KeepConn uint32 `idl:"name:wki1013_keep_conn" json:"keep_conn"`
}

func (o *WorkstationInfo1013) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *WorkstationInfo1013) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.KeepConn); err != nil {
		return err
	}
	return nil
}
func (o *WorkstationInfo1013) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.KeepConn); err != nil {
		return err
	}
	return nil
}

// WorkstationInfo1018 structure represents WKSTA_INFO_1018 RPC structure.
//
// The WKSTA_INFO_1018 structure specifies details about the state of the SMB network
// redirector.
type WorkstationInfo1018 struct {
	// wki1018_sess_timeout: The number of seconds the server (2) MUST wait before disconnecting
	// an inactive session.
	SessTimeout uint32 `idl:"name:wki1018_sess_timeout" json:"sess_timeout"`
}

func (o *WorkstationInfo1018) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *WorkstationInfo1018) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.SessTimeout); err != nil {
		return err
	}
	return nil
}
func (o *WorkstationInfo1018) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.SessTimeout); err != nil {
		return err
	}
	return nil
}

// WorkstationInfo1046 structure represents WKSTA_INFO_1046 RPC structure.
//
// The WKSTA_INFO_1046 structure specifies details about the state of the SMB network
// redirector.
type WorkstationInfo1046 struct {
	// wki1046_dormant_file_limit: The maximum number of file or printer handles the SMB
	// network redirector can continue to keep open, even after the application has closed
	// the corresponding handle.
	DormantFileLimit uint32 `idl:"name:wki1046_dormant_file_limit" json:"dormant_file_limit"`
}

func (o *WorkstationInfo1046) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *WorkstationInfo1046) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.DormantFileLimit); err != nil {
		return err
	}
	return nil
}
func (o *WorkstationInfo1046) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.DormantFileLimit); err != nil {
		return err
	}
	return nil
}

// WorkstationUserInfo0 structure represents WKSTA_USER_INFO_0 RPC structure.
//
// The WKSTA_USER_INFO_0 structure contains the name of a user who is currently active
// on the computer.
type WorkstationUserInfo0 struct {
	// wkui0_username: Null-terminated name of a user<7> who is currently active on the
	// computer. Multiple users can be currently active on a computer; this is the name
	// of any such user.
	UserName string `idl:"name:wkui0_username;string" json:"user_name"`
}

func (o *WorkstationUserInfo0) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *WorkstationUserInfo0) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(6); err != nil {
		return err
	}
	if o.UserName != "" {
		_ptr_wkui0_username := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.UserName); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.UserName, _ptr_wkui0_username); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *WorkstationUserInfo0) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(6); err != nil {
		return err
	}
	_ptr_wkui0_username := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.UserName); err != nil {
			return err
		}
		return nil
	})
	_s_wkui0_username := func(ptr interface{}) { o.UserName = *ptr.(*string) }
	if err := w.ReadPointer(&o.UserName, _s_wkui0_username, _ptr_wkui0_username); err != nil {
		return err
	}
	return nil
}

// WorkstationUserInfo1 structure represents WKSTA_USER_INFO_1 RPC structure.
//
// The WKSTA_USER_INFO_1 structure contains user information as it pertains to a specific
// computer.
type WorkstationUserInfo1 struct {
	// wkui1_username: A null-terminated name of a user who is currently active on the computer.
	UserName string `idl:"name:wkui1_username;string" json:"user_name"`
	// wkui1_logon_domain: A null-terminated name of the domain to which the user belongs.
	LogonDomain string `idl:"name:wkui1_logon_domain;string" json:"logon_domain"`
	// wkui1_oth_domains: Null-terminated, NetBIOS names of other domains browsed by the
	// computer, according to the OtherDomains Name Abstract Data Model (section 3.2.1.3).
	OtherDomains string `idl:"name:wkui1_oth_domains;string" json:"other_domains"`
	// wkui1_logon_server: A null-terminated, NetBIOS name of the server (2) that authenticated
	// (2) the user.
	LogonServer string `idl:"name:wkui1_logon_server;string" json:"logon_server"`
}

func (o *WorkstationUserInfo1) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *WorkstationUserInfo1) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(6); err != nil {
		return err
	}
	if o.UserName != "" {
		_ptr_wkui1_username := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.UserName); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.UserName, _ptr_wkui1_username); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.LogonDomain != "" {
		_ptr_wkui1_logon_domain := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.LogonDomain); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.LogonDomain, _ptr_wkui1_logon_domain); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.OtherDomains != "" {
		_ptr_wkui1_oth_domains := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.OtherDomains); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.OtherDomains, _ptr_wkui1_oth_domains); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.LogonServer != "" {
		_ptr_wkui1_logon_server := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.LogonServer); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.LogonServer, _ptr_wkui1_logon_server); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *WorkstationUserInfo1) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(6); err != nil {
		return err
	}
	_ptr_wkui1_username := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.UserName); err != nil {
			return err
		}
		return nil
	})
	_s_wkui1_username := func(ptr interface{}) { o.UserName = *ptr.(*string) }
	if err := w.ReadPointer(&o.UserName, _s_wkui1_username, _ptr_wkui1_username); err != nil {
		return err
	}
	_ptr_wkui1_logon_domain := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.LogonDomain); err != nil {
			return err
		}
		return nil
	})
	_s_wkui1_logon_domain := func(ptr interface{}) { o.LogonDomain = *ptr.(*string) }
	if err := w.ReadPointer(&o.LogonDomain, _s_wkui1_logon_domain, _ptr_wkui1_logon_domain); err != nil {
		return err
	}
	_ptr_wkui1_oth_domains := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.OtherDomains); err != nil {
			return err
		}
		return nil
	})
	_s_wkui1_oth_domains := func(ptr interface{}) { o.OtherDomains = *ptr.(*string) }
	if err := w.ReadPointer(&o.OtherDomains, _s_wkui1_oth_domains, _ptr_wkui1_oth_domains); err != nil {
		return err
	}
	_ptr_wkui1_logon_server := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.LogonServer); err != nil {
			return err
		}
		return nil
	})
	_s_wkui1_logon_server := func(ptr interface{}) { o.LogonServer = *ptr.(*string) }
	if err := w.ReadPointer(&o.LogonServer, _s_wkui1_logon_server, _ptr_wkui1_logon_server); err != nil {
		return err
	}
	return nil
}

// WorkstationTransportInfo0 structure represents WKSTA_TRANSPORT_INFO_0 RPC structure.
//
// The WKSTA_TRANSPORT_INFO_0 structure specifies details about the network transport
// protocol that the SMB network redirector uses.
type WorkstationTransportInfo0 struct {
	// wkti0_quality_of_service: Unused. Any value when sent, and MUST be ignored on receipt.
	QualityOfService uint32 `idl:"name:wkti0_quality_of_service" json:"quality_of_service"`
	// wkti0_number_of_vcs: The current number of remote connections using this transport
	// protocol.
	NumberOfVCS uint32 `idl:"name:wkti0_number_of_vcs" json:"number_of_vcs"`
	// wkti0_transport_name: The null-terminated, implementation-specific<5> name of the
	// device that implements the transport protocol.
	TransportName string `idl:"name:wkti0_transport_name;string" json:"transport_name"`
	// wkti0_transport_address: The null-terminated, implementation-specific<6> string that
	// represents the address of the transport protocol.
	TransportAddress string `idl:"name:wkti0_transport_address;string" json:"transport_address"`
	// wkti0_wan_ish: Whether the transport protocol is a routable protocol. If set to TRUE,
	// this is a routable protocol. If set to FALSE, this is not a routable protocol.
	WANIsh uint32 `idl:"name:wkti0_wan_ish" json:"wan_ish"`
}

func (o *WorkstationTransportInfo0) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *WorkstationTransportInfo0) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.QualityOfService); err != nil {
		return err
	}
	if err := w.WriteData(o.NumberOfVCS); err != nil {
		return err
	}
	if o.TransportName != "" {
		_ptr_wkti0_transport_name := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.TransportName); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.TransportName, _ptr_wkti0_transport_name); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.TransportAddress != "" {
		_ptr_wkti0_transport_address := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.TransportAddress); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.TransportAddress, _ptr_wkti0_transport_address); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.WANIsh); err != nil {
		return err
	}
	return nil
}
func (o *WorkstationTransportInfo0) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.QualityOfService); err != nil {
		return err
	}
	if err := w.ReadData(&o.NumberOfVCS); err != nil {
		return err
	}
	_ptr_wkti0_transport_name := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.TransportName); err != nil {
			return err
		}
		return nil
	})
	_s_wkti0_transport_name := func(ptr interface{}) { o.TransportName = *ptr.(*string) }
	if err := w.ReadPointer(&o.TransportName, _s_wkti0_transport_name, _ptr_wkti0_transport_name); err != nil {
		return err
	}
	_ptr_wkti0_transport_address := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.TransportAddress); err != nil {
			return err
		}
		return nil
	})
	_s_wkti0_transport_address := func(ptr interface{}) { o.TransportAddress = *ptr.(*string) }
	if err := w.ReadPointer(&o.TransportAddress, _s_wkti0_transport_address, _ptr_wkti0_transport_address); err != nil {
		return err
	}
	if err := w.ReadData(&o.WANIsh); err != nil {
		return err
	}
	return nil
}

// WorkstationInfo structure represents WKSTA_INFO RPC union.
//
// The WKSTA_INFO union specifies details about a computer. This union is used by the
// methods NetrWkstaGetInfo (section 3.2.4.1) and NetrWkstaSetInfo (section 3.2.4.2).
type WorkstationInfo struct {
	// Types that are assignable to Value
	//
	// *WorkstationInfo_100
	// *WorkstationInfo_101
	// *WorkstationInfo_102
	// *WorkstationInfo_502
	// *WorkstationInfo_1013
	// *WorkstationInfo_1018
	// *WorkstationInfo_1046
	Value is_WorkstationInfo `json:"value"`
}

func (o *WorkstationInfo) GetValue() any {
	if o == nil {
		return nil
	}
	switch value := (interface{})(o.Value).(type) {
	case *WorkstationInfo_100:
		if value != nil {
			return value.WorkstationInfo100
		}
	case *WorkstationInfo_101:
		if value != nil {
			return value.WorkstationInfo101
		}
	case *WorkstationInfo_102:
		if value != nil {
			return value.WorkstationInfo102
		}
	case *WorkstationInfo_502:
		if value != nil {
			return value.WorkstationInfo502
		}
	case *WorkstationInfo_1013:
		if value != nil {
			return value.WorkstationInfo1013
		}
	case *WorkstationInfo_1018:
		if value != nil {
			return value.WorkstationInfo1018
		}
	case *WorkstationInfo_1046:
		if value != nil {
			return value.WorkstationInfo1046
		}
	}
	return nil
}

type is_WorkstationInfo interface {
	ndr.Marshaler
	ndr.Unmarshaler
	is_WorkstationInfo()
}

func (o *WorkstationInfo) NDRSwitchValue(sw uint32) uint32 {
	if o == nil {
		return uint32(0)
	}
	switch (interface{})(o.Value).(type) {
	case *WorkstationInfo_100:
		return uint32(100)
	case *WorkstationInfo_101:
		return uint32(101)
	case *WorkstationInfo_102:
		return uint32(102)
	case *WorkstationInfo_502:
		return uint32(502)
	case *WorkstationInfo_1013:
		return uint32(1013)
	case *WorkstationInfo_1018:
		return uint32(1018)
	case *WorkstationInfo_1046:
		return uint32(1046)
	}
	return uint32(0)
}

func (o *WorkstationInfo) MarshalUnionNDR(ctx context.Context, w ndr.Writer, sw uint32) error {
	if err := w.WriteSwitch(uint32(sw)); err != nil {
		return err
	}
	switch sw {
	case uint32(100):
		_o, _ := o.Value.(*WorkstationInfo_100)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&WorkstationInfo_100{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(101):
		_o, _ := o.Value.(*WorkstationInfo_101)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&WorkstationInfo_101{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(102):
		_o, _ := o.Value.(*WorkstationInfo_102)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&WorkstationInfo_102{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(502):
		_o, _ := o.Value.(*WorkstationInfo_502)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&WorkstationInfo_502{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(1013):
		_o, _ := o.Value.(*WorkstationInfo_1013)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&WorkstationInfo_1013{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(1018):
		_o, _ := o.Value.(*WorkstationInfo_1018)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&WorkstationInfo_1018{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(1046):
		_o, _ := o.Value.(*WorkstationInfo_1046)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&WorkstationInfo_1046{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	default:
	}
	return nil
}

func (o *WorkstationInfo) UnmarshalUnionNDR(ctx context.Context, w ndr.Reader, sw uint32) error {
	if err := w.ReadSwitch((*uint32)(&sw)); err != nil {
		return err
	}
	switch sw {
	case uint32(100):
		o.Value = &WorkstationInfo_100{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(101):
		o.Value = &WorkstationInfo_101{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(102):
		o.Value = &WorkstationInfo_102{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(502):
		o.Value = &WorkstationInfo_502{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(1013):
		o.Value = &WorkstationInfo_1013{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(1018):
		o.Value = &WorkstationInfo_1018{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(1046):
		o.Value = &WorkstationInfo_1046{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	default:
	}
	return nil
}

// WorkstationInfo_100 structure represents WKSTA_INFO RPC union arm.
//
// It has following labels: 100
type WorkstationInfo_100 struct {
	// WkstaInfo100: Details about the computer environment (section 2.2.5.1)
	WorkstationInfo100 *WorkstationInfo100 `idl:"name:WkstaInfo100" json:"workstation_info100"`
}

func (*WorkstationInfo_100) is_WorkstationInfo() {}

func (o *WorkstationInfo_100) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.WorkstationInfo100 != nil {
		_ptr_WkstaInfo100 := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.WorkstationInfo100 != nil {
				if err := o.WorkstationInfo100.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&WorkstationInfo100{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.WorkstationInfo100, _ptr_WkstaInfo100); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *WorkstationInfo_100) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_WkstaInfo100 := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.WorkstationInfo100 == nil {
			o.WorkstationInfo100 = &WorkstationInfo100{}
		}
		if err := o.WorkstationInfo100.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_WkstaInfo100 := func(ptr interface{}) { o.WorkstationInfo100 = *ptr.(**WorkstationInfo100) }
	if err := w.ReadPointer(&o.WorkstationInfo100, _s_WkstaInfo100, _ptr_WkstaInfo100); err != nil {
		return err
	}
	return nil
}

// WorkstationInfo_101 structure represents WKSTA_INFO RPC union arm.
//
// It has following labels: 101
type WorkstationInfo_101 struct {
	// WkstaInfo101: Details about the computer environment (section 2.2.5.2).
	WorkstationInfo101 *WorkstationInfo101 `idl:"name:WkstaInfo101" json:"workstation_info101"`
}

func (*WorkstationInfo_101) is_WorkstationInfo() {}

func (o *WorkstationInfo_101) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.WorkstationInfo101 != nil {
		_ptr_WkstaInfo101 := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.WorkstationInfo101 != nil {
				if err := o.WorkstationInfo101.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&WorkstationInfo101{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.WorkstationInfo101, _ptr_WkstaInfo101); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *WorkstationInfo_101) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_WkstaInfo101 := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.WorkstationInfo101 == nil {
			o.WorkstationInfo101 = &WorkstationInfo101{}
		}
		if err := o.WorkstationInfo101.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_WkstaInfo101 := func(ptr interface{}) { o.WorkstationInfo101 = *ptr.(**WorkstationInfo101) }
	if err := w.ReadPointer(&o.WorkstationInfo101, _s_WkstaInfo101, _ptr_WkstaInfo101); err != nil {
		return err
	}
	return nil
}

// WorkstationInfo_102 structure represents WKSTA_INFO RPC union arm.
//
// It has following labels: 102
type WorkstationInfo_102 struct {
	// WkstaInfo102: Details about the computer environment (section 2.2.5.3).
	WorkstationInfo102 *WorkstationInfo102 `idl:"name:WkstaInfo102" json:"workstation_info102"`
}

func (*WorkstationInfo_102) is_WorkstationInfo() {}

func (o *WorkstationInfo_102) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.WorkstationInfo102 != nil {
		_ptr_WkstaInfo102 := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.WorkstationInfo102 != nil {
				if err := o.WorkstationInfo102.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&WorkstationInfo102{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.WorkstationInfo102, _ptr_WkstaInfo102); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *WorkstationInfo_102) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_WkstaInfo102 := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.WorkstationInfo102 == nil {
			o.WorkstationInfo102 = &WorkstationInfo102{}
		}
		if err := o.WorkstationInfo102.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_WkstaInfo102 := func(ptr interface{}) { o.WorkstationInfo102 = *ptr.(**WorkstationInfo102) }
	if err := w.ReadPointer(&o.WorkstationInfo102, _s_WkstaInfo102, _ptr_WkstaInfo102); err != nil {
		return err
	}
	return nil
}

// WorkstationInfo_502 structure represents WKSTA_INFO RPC union arm.
//
// It has following labels: 502
type WorkstationInfo_502 struct {
	// WkstaInfo502: Details about the computer environment (section 2.2.5.4).
	WorkstationInfo502 *WorkstationInfo502 `idl:"name:WkstaInfo502" json:"workstation_info502"`
}

func (*WorkstationInfo_502) is_WorkstationInfo() {}

func (o *WorkstationInfo_502) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.WorkstationInfo502 != nil {
		_ptr_WkstaInfo502 := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.WorkstationInfo502 != nil {
				if err := o.WorkstationInfo502.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&WorkstationInfo502{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.WorkstationInfo502, _ptr_WkstaInfo502); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *WorkstationInfo_502) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_WkstaInfo502 := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.WorkstationInfo502 == nil {
			o.WorkstationInfo502 = &WorkstationInfo502{}
		}
		if err := o.WorkstationInfo502.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_WkstaInfo502 := func(ptr interface{}) { o.WorkstationInfo502 = *ptr.(**WorkstationInfo502) }
	if err := w.ReadPointer(&o.WorkstationInfo502, _s_WkstaInfo502, _ptr_WkstaInfo502); err != nil {
		return err
	}
	return nil
}

// WorkstationInfo_1013 structure represents WKSTA_INFO RPC union arm.
//
// It has following labels: 1013
type WorkstationInfo_1013 struct {
	// WkstaInfo1013: Details about the state of the SMB network redirector (section 2.2.5.5).
	WorkstationInfo1013 *WorkstationInfo1013 `idl:"name:WkstaInfo1013" json:"workstation_info1013"`
}

func (*WorkstationInfo_1013) is_WorkstationInfo() {}

func (o *WorkstationInfo_1013) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.WorkstationInfo1013 != nil {
		_ptr_WkstaInfo1013 := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.WorkstationInfo1013 != nil {
				if err := o.WorkstationInfo1013.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&WorkstationInfo1013{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.WorkstationInfo1013, _ptr_WkstaInfo1013); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *WorkstationInfo_1013) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_WkstaInfo1013 := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.WorkstationInfo1013 == nil {
			o.WorkstationInfo1013 = &WorkstationInfo1013{}
		}
		if err := o.WorkstationInfo1013.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_WkstaInfo1013 := func(ptr interface{}) { o.WorkstationInfo1013 = *ptr.(**WorkstationInfo1013) }
	if err := w.ReadPointer(&o.WorkstationInfo1013, _s_WkstaInfo1013, _ptr_WkstaInfo1013); err != nil {
		return err
	}
	return nil
}

// WorkstationInfo_1018 structure represents WKSTA_INFO RPC union arm.
//
// It has following labels: 1018
type WorkstationInfo_1018 struct {
	// WkstaInfo1018: Details about the state of the SMB network redirector (section 2.2.5.6).
	WorkstationInfo1018 *WorkstationInfo1018 `idl:"name:WkstaInfo1018" json:"workstation_info1018"`
}

func (*WorkstationInfo_1018) is_WorkstationInfo() {}

func (o *WorkstationInfo_1018) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.WorkstationInfo1018 != nil {
		_ptr_WkstaInfo1018 := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.WorkstationInfo1018 != nil {
				if err := o.WorkstationInfo1018.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&WorkstationInfo1018{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.WorkstationInfo1018, _ptr_WkstaInfo1018); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *WorkstationInfo_1018) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_WkstaInfo1018 := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.WorkstationInfo1018 == nil {
			o.WorkstationInfo1018 = &WorkstationInfo1018{}
		}
		if err := o.WorkstationInfo1018.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_WkstaInfo1018 := func(ptr interface{}) { o.WorkstationInfo1018 = *ptr.(**WorkstationInfo1018) }
	if err := w.ReadPointer(&o.WorkstationInfo1018, _s_WkstaInfo1018, _ptr_WkstaInfo1018); err != nil {
		return err
	}
	return nil
}

// WorkstationInfo_1046 structure represents WKSTA_INFO RPC union arm.
//
// It has following labels: 1046
type WorkstationInfo_1046 struct {
	// WkstaInfo1046: Details about the state of the SMB network redirector (section 2.2.5.7).
	WorkstationInfo1046 *WorkstationInfo1046 `idl:"name:WkstaInfo1046" json:"workstation_info1046"`
}

func (*WorkstationInfo_1046) is_WorkstationInfo() {}

func (o *WorkstationInfo_1046) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.WorkstationInfo1046 != nil {
		_ptr_WkstaInfo1046 := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.WorkstationInfo1046 != nil {
				if err := o.WorkstationInfo1046.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&WorkstationInfo1046{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.WorkstationInfo1046, _ptr_WkstaInfo1046); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *WorkstationInfo_1046) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_WkstaInfo1046 := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.WorkstationInfo1046 == nil {
			o.WorkstationInfo1046 = &WorkstationInfo1046{}
		}
		if err := o.WorkstationInfo1046.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_WkstaInfo1046 := func(ptr interface{}) { o.WorkstationInfo1046 = *ptr.(**WorkstationInfo1046) }
	if err := w.ReadPointer(&o.WorkstationInfo1046, _s_WkstaInfo1046, _ptr_WkstaInfo1046); err != nil {
		return err
	}
	return nil
}

// UseInfo0 structure represents USE_INFO_0 RPC structure.
//
// The USE_INFO_0 structure specifies details about the connection between a machine
// on which the workstation service is running and a shared resource.
type UseInfo0 struct {
	// ui0_local: A pointer to a string that contains the device name (for example, drive
	// E or LPT1) being redirected to the shared resource.
	Local string `idl:"name:ui0_local;string" json:"local"`
	// ui0_remote: A pointer to a string that contains the share name of the remote resource
	// being accessed. The string MUST be in the following form: "\\servername\sharename".
	Remote string `idl:"name:ui0_remote;string" json:"remote"`
}

func (o *UseInfo0) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *UseInfo0) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(6); err != nil {
		return err
	}
	if o.Local != "" {
		_ptr_ui0_local := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.Local); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.Local, _ptr_ui0_local); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.Remote != "" {
		_ptr_ui0_remote := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.Remote); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.Remote, _ptr_ui0_remote); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *UseInfo0) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(6); err != nil {
		return err
	}
	_ptr_ui0_local := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.Local); err != nil {
			return err
		}
		return nil
	})
	_s_ui0_local := func(ptr interface{}) { o.Local = *ptr.(*string) }
	if err := w.ReadPointer(&o.Local, _s_ui0_local, _ptr_ui0_local); err != nil {
		return err
	}
	_ptr_ui0_remote := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.Remote); err != nil {
			return err
		}
		return nil
	})
	_s_ui0_remote := func(ptr interface{}) { o.Remote = *ptr.(*string) }
	if err := w.ReadPointer(&o.Remote, _s_ui0_remote, _ptr_ui0_remote); err != nil {
		return err
	}
	return nil
}

// UseInfo1 structure represents USE_INFO_1 RPC structure.
//
// The USE_INFO_1 structure specifies details about the connection between a machine
// on which the workstation service is running and a shared resource, including connection
// status and connection type.
type UseInfo1 struct {
	// ui1_local: A pointer to a string that contains the device name (for example, drive
	// E or LPT1) being redirected to the shared resource.
	Local string `idl:"name:ui1_local;string" json:"local"`
	// ui1_remote: A pointer to a string that contains the share name of the remote resource
	// being accessed. The string MUST be in the following form: "\\servername\sharename".
	Remote string `idl:"name:ui1_remote;string" json:"remote"`
	// ui1_password: A pointer to a string that contains the password needed to establish
	// a session between a machine on which the workstation service is running and a server.
	Password string `idl:"name:ui1_password;string" json:"password"`
	// ui1_status: The current status of the connection, which MUST contain one of the following
	// values:
	//
	//	+-------------------------+-------------------------------+
	//	|                         |                               |
	//	|       VALUE/CODE        |            MEANING            |
	//	|                         |                               |
	//	+-------------------------+-------------------------------+
	//	+-------------------------+-------------------------------+
	//	| USE_OK 0x00000000       | The connection is valid.      |
	//	+-------------------------+-------------------------------+
	//	| USE_PAUSED 0x00000001   | Paused by local workstation.  |
	//	+-------------------------+-------------------------------+
	//	| USE_SESSLOST 0x00000002 | Disconnected.                 |
	//	+-------------------------+-------------------------------+
	//	| USE_NETERR 0x00000003   | A network error occurred.     |
	//	+-------------------------+-------------------------------+
	//	| USE_CONN 0x00000004     | The connection is being made. |
	//	+-------------------------+-------------------------------+
	//	| USE_RECONN 0x00000005   | Reconnecting.                 |
	//	+-------------------------+-------------------------------+
	Status uint32 `idl:"name:ui1_status" json:"status"`
	// ui1_asg_type: The type of remote resource being accessed, which MUST contain one
	// of the following values:
	//
	//	+-------------------------+----------------------------------------------------------------------------------+
	//	|                         |                                                                                  |
	//	|       VALUE/CODE        |                                     MEANING                                      |
	//	|                         |                                                                                  |
	//	+-------------------------+----------------------------------------------------------------------------------+
	//	+-------------------------+----------------------------------------------------------------------------------+
	//	| USE_WILDCARD 0xFFFFFFFF | Matches the type of the server’s shared resources. Wildcards can be used only    |
	//	|                         | with the NetrUseAdd function (section 3.2.4.7), and only when the ui1_local      |
	//	|                         | member is NULL.                                                                  |
	//	+-------------------------+----------------------------------------------------------------------------------+
	//	| USE_DISKDEV 0x00000000  | Disk device.                                                                     |
	//	+-------------------------+----------------------------------------------------------------------------------+
	//	| USE_SPOOLDEV 0x00000001 | Spooled printer.                                                                 |
	//	+-------------------------+----------------------------------------------------------------------------------+
	//	| USE_CHARDEV 0x00000002  | Serial device.                                                                   |
	//	+-------------------------+----------------------------------------------------------------------------------+
	//	| USE_IPC 0x00000003      | Inter process communication (IPC).                                               |
	//	+-------------------------+----------------------------------------------------------------------------------+
	AsgType uint32 `idl:"name:ui1_asg_type" json:"asg_type"`
	// ui1_refcount: The number of files, directories, and other processes that can be opened
	// on the remote resource.
	ReferenceCount uint32 `idl:"name:ui1_refcount" json:"reference_count"`
	// ui1_usecount: The number of explicit connections (with a device name) or implicit
	// UNC connections (without the device name) that are established with the resource.
	UseCount uint32 `idl:"name:ui1_usecount" json:"use_count"`
}

func (o *UseInfo1) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *UseInfo1) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if o.Local != "" {
		_ptr_ui1_local := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.Local); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.Local, _ptr_ui1_local); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.Remote != "" {
		_ptr_ui1_remote := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.Remote); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.Remote, _ptr_ui1_remote); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.Password != "" {
		_ptr_ui1_password := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.Password); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.Password, _ptr_ui1_password); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.Status); err != nil {
		return err
	}
	if err := w.WriteData(o.AsgType); err != nil {
		return err
	}
	if err := w.WriteData(o.ReferenceCount); err != nil {
		return err
	}
	if err := w.WriteData(o.UseCount); err != nil {
		return err
	}
	return nil
}
func (o *UseInfo1) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	_ptr_ui1_local := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.Local); err != nil {
			return err
		}
		return nil
	})
	_s_ui1_local := func(ptr interface{}) { o.Local = *ptr.(*string) }
	if err := w.ReadPointer(&o.Local, _s_ui1_local, _ptr_ui1_local); err != nil {
		return err
	}
	_ptr_ui1_remote := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.Remote); err != nil {
			return err
		}
		return nil
	})
	_s_ui1_remote := func(ptr interface{}) { o.Remote = *ptr.(*string) }
	if err := w.ReadPointer(&o.Remote, _s_ui1_remote, _ptr_ui1_remote); err != nil {
		return err
	}
	_ptr_ui1_password := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.Password); err != nil {
			return err
		}
		return nil
	})
	_s_ui1_password := func(ptr interface{}) { o.Password = *ptr.(*string) }
	if err := w.ReadPointer(&o.Password, _s_ui1_password, _ptr_ui1_password); err != nil {
		return err
	}
	if err := w.ReadData(&o.Status); err != nil {
		return err
	}
	if err := w.ReadData(&o.AsgType); err != nil {
		return err
	}
	if err := w.ReadData(&o.ReferenceCount); err != nil {
		return err
	}
	if err := w.ReadData(&o.UseCount); err != nil {
		return err
	}
	return nil
}

// UseInfo2 structure represents USE_INFO_2 RPC structure.
//
// The USE_INFO_2 structure specifies details about the connection between a machine
// on which the workstation service is running and a shared resource, including user
// name and domain name.
type UseInfo2 struct {
	// ui2_useinfo: A pointer to the USE_INFO_1 structure (section 2.2.5.23) that the method
	// returns.
	Useinfo *UseInfo1 `idl:"name:ui2_useinfo" json:"useinfo"`
	// ui2_username: A pointer to a string that contains the name of the user who initiated
	// the connection.
	UserName string `idl:"name:ui2_username;string" json:"user_name"`
	// ui2_domainname: A pointer to a string that contains the domain name of the remote
	// resource.
	DomainName string `idl:"name:ui2_domainname;string" json:"domain_name"`
}

func (o *UseInfo2) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *UseInfo2) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if o.Useinfo != nil {
		if err := o.Useinfo.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&UseInfo1{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.UserName != "" {
		_ptr_ui2_username := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.UserName); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.UserName, _ptr_ui2_username); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.DomainName != "" {
		_ptr_ui2_domainname := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.DomainName); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.DomainName, _ptr_ui2_domainname); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *UseInfo2) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if o.Useinfo == nil {
		o.Useinfo = &UseInfo1{}
	}
	if err := o.Useinfo.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	_ptr_ui2_username := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.UserName); err != nil {
			return err
		}
		return nil
	})
	_s_ui2_username := func(ptr interface{}) { o.UserName = *ptr.(*string) }
	if err := w.ReadPointer(&o.UserName, _s_ui2_username, _ptr_ui2_username); err != nil {
		return err
	}
	_ptr_ui2_domainname := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.DomainName); err != nil {
			return err
		}
		return nil
	})
	_s_ui2_domainname := func(ptr interface{}) { o.DomainName = *ptr.(*string) }
	if err := w.ReadPointer(&o.DomainName, _s_ui2_domainname, _ptr_ui2_domainname); err != nil {
		return err
	}
	return nil
}

// UseInfo3 structure represents USE_INFO_3 RPC structure.
//
// The USE_INFO_3 structure specifies details about the connection between a machine
// on which the workstation service is running and a shared resource, including user
// name and domain name.
type UseInfo3 struct {
	// ui3_ui2: A pointer to the USE_INFO_2 structure (section 2.2.5.24) that the method
	// returns.
	UI2 *UseInfo2 `idl:"name:ui3_ui2" json:"ui2"`
	// ui3_flags: A reserved field that the client MUST set to zero, and the server MUST
	// ignore on receipt.
	Flags uint32 `idl:"name:ui3_flags" json:"flags"`
}

func (o *UseInfo3) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *UseInfo3) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if o.UI2 != nil {
		if err := o.UI2.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&UseInfo2{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.Flags); err != nil {
		return err
	}
	return nil
}
func (o *UseInfo3) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if o.UI2 == nil {
		o.UI2 = &UseInfo2{}
	}
	if err := o.UI2.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.Flags); err != nil {
		return err
	}
	return nil
}

// UseInfo structure represents USE_INFO RPC union.
//
// The USE_INFO union specifies details about the connection between a machine on which
// the workstation service is running and a shared resource. This union is used by the
// methods NetrUseAdd (section 3.2.4.7) and NetrUseGetInfo (section 3.2.4.8).
type UseInfo struct {
	// Types that are assignable to Value
	//
	// *UseInfo_0
	// *UseInfo_1
	// *UseInfo_2
	// *UseInfo_3
	Value is_UseInfo `json:"value"`
}

func (o *UseInfo) GetValue() any {
	if o == nil {
		return nil
	}
	switch value := (interface{})(o.Value).(type) {
	case *UseInfo_0:
		if value != nil {
			return value.UseInfo0
		}
	case *UseInfo_1:
		if value != nil {
			return value.UseInfo1
		}
	case *UseInfo_2:
		if value != nil {
			return value.UseInfo2
		}
	case *UseInfo_3:
		if value != nil {
			return value.UseInfo3
		}
	}
	return nil
}

type is_UseInfo interface {
	ndr.Marshaler
	ndr.Unmarshaler
	is_UseInfo()
}

func (o *UseInfo) NDRSwitchValue(sw uint32) uint32 {
	if o == nil {
		return uint32(0)
	}
	switch (interface{})(o.Value).(type) {
	case *UseInfo_0:
		return uint32(0)
	case *UseInfo_1:
		return uint32(1)
	case *UseInfo_2:
		return uint32(2)
	case *UseInfo_3:
		return uint32(3)
	}
	return uint32(0)
}

func (o *UseInfo) MarshalUnionNDR(ctx context.Context, w ndr.Writer, sw uint32) error {
	if err := w.WriteSwitch(uint32(sw)); err != nil {
		return err
	}
	switch sw {
	case uint32(0):
		_o, _ := o.Value.(*UseInfo_0)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&UseInfo_0{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(1):
		_o, _ := o.Value.(*UseInfo_1)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&UseInfo_1{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(2):
		_o, _ := o.Value.(*UseInfo_2)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&UseInfo_2{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(3):
		_o, _ := o.Value.(*UseInfo_3)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&UseInfo_3{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	default:
	}
	return nil
}

func (o *UseInfo) UnmarshalUnionNDR(ctx context.Context, w ndr.Reader, sw uint32) error {
	if err := w.ReadSwitch((*uint32)(&sw)); err != nil {
		return err
	}
	switch sw {
	case uint32(0):
		o.Value = &UseInfo_0{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(1):
		o.Value = &UseInfo_1{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(2):
		o.Value = &UseInfo_2{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(3):
		o.Value = &UseInfo_3{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	default:
	}
	return nil
}

// UseInfo_0 structure represents USE_INFO RPC union arm.
//
// It has following labels: 0
type UseInfo_0 struct {
	// UseInfo0: Details about a connection (section 2.2.5.22).
	UseInfo0 *UseInfo0 `idl:"name:UseInfo0" json:"use_info0"`
}

func (*UseInfo_0) is_UseInfo() {}

func (o *UseInfo_0) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.UseInfo0 != nil {
		_ptr_UseInfo0 := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.UseInfo0 != nil {
				if err := o.UseInfo0.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&UseInfo0{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.UseInfo0, _ptr_UseInfo0); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *UseInfo_0) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_UseInfo0 := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.UseInfo0 == nil {
			o.UseInfo0 = &UseInfo0{}
		}
		if err := o.UseInfo0.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_UseInfo0 := func(ptr interface{}) { o.UseInfo0 = *ptr.(**UseInfo0) }
	if err := w.ReadPointer(&o.UseInfo0, _s_UseInfo0, _ptr_UseInfo0); err != nil {
		return err
	}
	return nil
}

// UseInfo_1 structure represents USE_INFO RPC union arm.
//
// It has following labels: 1
type UseInfo_1 struct {
	// UseInfo1: Details about a connection (section 2.2.5.23).
	UseInfo1 *UseInfo1 `idl:"name:UseInfo1" json:"use_info1"`
}

func (*UseInfo_1) is_UseInfo() {}

func (o *UseInfo_1) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.UseInfo1 != nil {
		_ptr_UseInfo1 := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.UseInfo1 != nil {
				if err := o.UseInfo1.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&UseInfo1{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.UseInfo1, _ptr_UseInfo1); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *UseInfo_1) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_UseInfo1 := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.UseInfo1 == nil {
			o.UseInfo1 = &UseInfo1{}
		}
		if err := o.UseInfo1.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_UseInfo1 := func(ptr interface{}) { o.UseInfo1 = *ptr.(**UseInfo1) }
	if err := w.ReadPointer(&o.UseInfo1, _s_UseInfo1, _ptr_UseInfo1); err != nil {
		return err
	}
	return nil
}

// UseInfo_2 structure represents USE_INFO RPC union arm.
//
// It has following labels: 2
type UseInfo_2 struct {
	// UseInfo2: Details about a connection (section 2.2.5.24).
	UseInfo2 *UseInfo2 `idl:"name:UseInfo2" json:"use_info2"`
}

func (*UseInfo_2) is_UseInfo() {}

func (o *UseInfo_2) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.UseInfo2 != nil {
		_ptr_UseInfo2 := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.UseInfo2 != nil {
				if err := o.UseInfo2.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&UseInfo2{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.UseInfo2, _ptr_UseInfo2); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *UseInfo_2) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_UseInfo2 := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.UseInfo2 == nil {
			o.UseInfo2 = &UseInfo2{}
		}
		if err := o.UseInfo2.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_UseInfo2 := func(ptr interface{}) { o.UseInfo2 = *ptr.(**UseInfo2) }
	if err := w.ReadPointer(&o.UseInfo2, _s_UseInfo2, _ptr_UseInfo2); err != nil {
		return err
	}
	return nil
}

// UseInfo_3 structure represents USE_INFO RPC union arm.
//
// It has following labels: 3
type UseInfo_3 struct {
	// UseInfo3: Details about a connection (section 2.2.5.25).
	UseInfo3 *UseInfo3 `idl:"name:UseInfo3" json:"use_info3"`
}

func (*UseInfo_3) is_UseInfo() {}

func (o *UseInfo_3) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.UseInfo3 != nil {
		_ptr_UseInfo3 := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.UseInfo3 != nil {
				if err := o.UseInfo3.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&UseInfo3{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.UseInfo3, _ptr_UseInfo3); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *UseInfo_3) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_UseInfo3 := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.UseInfo3 == nil {
			o.UseInfo3 = &UseInfo3{}
		}
		if err := o.UseInfo3.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_UseInfo3 := func(ptr interface{}) { o.UseInfo3 = *ptr.(**UseInfo3) }
	if err := w.ReadPointer(&o.UseInfo3, _s_UseInfo3, _ptr_UseInfo3); err != nil {
		return err
	}
	return nil
}

// UseInfo0Container structure represents USE_INFO_0_CONTAINER RPC structure.
//
// The USE_INFO_0_CONTAINER structure contains a value that indicates the number of
// entries that the NetrUseEnum method (section 3.2.4.10) returns, as well as a pointer
// to the buffer.
type UseInfo0Container struct {
	// EntriesRead: The number of entries that the method returns.
	EntriesRead uint32 `idl:"name:EntriesRead" json:"entries_read"`
	// Buffer: Details about the connection between a device and a shared resource.
	Buffer *UseInfo0 `idl:"name:Buffer" json:"buffer"`
}

func (o *UseInfo0Container) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *UseInfo0Container) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.EntriesRead); err != nil {
		return err
	}
	if o.Buffer != nil {
		_ptr_Buffer := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.Buffer != nil {
				if err := o.Buffer.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&UseInfo0{}).MarshalNDR(ctx, w); err != nil {
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
func (o *UseInfo0Container) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.EntriesRead); err != nil {
		return err
	}
	_ptr_Buffer := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.Buffer == nil {
			o.Buffer = &UseInfo0{}
		}
		if err := o.Buffer.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_Buffer := func(ptr interface{}) { o.Buffer = *ptr.(**UseInfo0) }
	if err := w.ReadPointer(&o.Buffer, _s_Buffer, _ptr_Buffer); err != nil {
		return err
	}
	return nil
}

// UseInfo1Container structure represents USE_INFO_1_CONTAINER RPC structure.
//
// The USE_INFO_1_CONTAINER structure contains a value that indicates the number of
// entries that the NetrUseEnum method (section 3.2.4.10) returns, as well as a pointer
// to the buffer.
type UseInfo1Container struct {
	// EntriesRead: The number of entries that the method returns.
	EntriesRead uint32 `idl:"name:EntriesRead" json:"entries_read"`
	// Buffer: Details about the connection between a machine on which the workstation service
	// is running and a shared resource.
	Buffer *UseInfo1 `idl:"name:Buffer" json:"buffer"`
}

func (o *UseInfo1Container) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *UseInfo1Container) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.EntriesRead); err != nil {
		return err
	}
	if o.Buffer != nil {
		_ptr_Buffer := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.Buffer != nil {
				if err := o.Buffer.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&UseInfo1{}).MarshalNDR(ctx, w); err != nil {
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
func (o *UseInfo1Container) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.EntriesRead); err != nil {
		return err
	}
	_ptr_Buffer := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.Buffer == nil {
			o.Buffer = &UseInfo1{}
		}
		if err := o.Buffer.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_Buffer := func(ptr interface{}) { o.Buffer = *ptr.(**UseInfo1) }
	if err := w.ReadPointer(&o.Buffer, _s_Buffer, _ptr_Buffer); err != nil {
		return err
	}
	return nil
}

// UseInfo2Container structure represents USE_INFO_2_CONTAINER RPC structure.
//
// The USE_INFO_2_CONTAINER structure contains a value that indicates the number of
// entries that the NetrUseEnum method (section 3.2.4.10) returns, as well as a pointer
// to the buffer.
type UseInfo2Container struct {
	// EntriesRead: The number of entries that the method returns.
	EntriesRead uint32 `idl:"name:EntriesRead" json:"entries_read"`
	// Buffer: Details about the connection between a machine on which the workstation service
	// is running and a shared resource.
	Buffer *UseInfo2 `idl:"name:Buffer" json:"buffer"`
}

func (o *UseInfo2Container) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *UseInfo2Container) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.EntriesRead); err != nil {
		return err
	}
	if o.Buffer != nil {
		_ptr_Buffer := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.Buffer != nil {
				if err := o.Buffer.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&UseInfo2{}).MarshalNDR(ctx, w); err != nil {
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
func (o *UseInfo2Container) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.EntriesRead); err != nil {
		return err
	}
	_ptr_Buffer := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.Buffer == nil {
			o.Buffer = &UseInfo2{}
		}
		if err := o.Buffer.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_Buffer := func(ptr interface{}) { o.Buffer = *ptr.(**UseInfo2) }
	if err := w.ReadPointer(&o.Buffer, _s_Buffer, _ptr_Buffer); err != nil {
		return err
	}
	return nil
}

// UseEnum structure represents USE_ENUM_STRUCT RPC structure.
//
// The USE_ENUM_STRUCT structure is used by the NetrUseEnum method (section 3.2.4.10)
// to encapsulate the _USE_ENUM_UNION union.
type UseEnum struct {
	// Level: A value that specifies the information level of the data. This parameter MUST
	// be one of the following values.
	//
	//	+------------+------------------------------------------------------------------+
	//	|            |                                                                  |
	//	|   VALUE    |                             MEANING                              |
	//	|            |                                                                  |
	//	+------------+------------------------------------------------------------------+
	//	+------------+------------------------------------------------------------------+
	//	| 0x00000000 | The UseInfo buffer is a USE_INFO_0 structure (section 2.2.5.22). |
	//	+------------+------------------------------------------------------------------+
	//	| 0x00000001 | The UseInfo buffer is a USE_INFO_1 structure (section 2.2.5.23). |
	//	+------------+------------------------------------------------------------------+
	//	| 0x00000002 | The UseInfo buffer is a USE_INFO_2 structure (section 2.2.5.24). |
	//	+------------+------------------------------------------------------------------+
	Level uint32 `idl:"name:Level" json:"level"`
	// UseInfo: A buffer containing any one of the USE_INFO_0, USE_INFO_1, or USE_INFO_2
	// structures.
	UseInfo *UseEnum_UseInfo `idl:"name:UseInfo;switch_is:Level" json:"use_info"`
}

func (o *UseEnum) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *UseEnum) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.Level); err != nil {
		return err
	}
	_swUseInfo := uint32(o.Level)
	if o.UseInfo != nil {
		if err := o.UseInfo.MarshalUnionNDR(ctx, w, _swUseInfo); err != nil {
			return err
		}
	} else {
		if err := (&UseEnum_UseInfo{}).MarshalUnionNDR(ctx, w, _swUseInfo); err != nil {
			return err
		}
	}
	return nil
}
func (o *UseEnum) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.Level); err != nil {
		return err
	}
	if o.UseInfo == nil {
		o.UseInfo = &UseEnum_UseInfo{}
	}
	_swUseInfo := uint32(o.Level)
	if err := o.UseInfo.UnmarshalUnionNDR(ctx, w, _swUseInfo); err != nil {
		return err
	}
	return nil
}

// UseEnum_UseInfo structure represents USE_ENUM_STRUCT union anonymous member.
//
// The USE_ENUM_STRUCT structure is used by the NetrUseEnum method (section 3.2.4.10)
// to encapsulate the _USE_ENUM_UNION union.
type UseEnum_UseInfo struct {
	// Types that are assignable to Value
	//
	// *UseInfo_Level0
	// *UseInfo_Level1
	// *UseInfo_Level2
	Value is_UseEnum_UseInfo `json:"value"`
}

func (o *UseEnum_UseInfo) GetValue() any {
	if o == nil {
		return nil
	}
	switch value := (interface{})(o.Value).(type) {
	case *UseInfo_Level0:
		if value != nil {
			return value.Level0
		}
	case *UseInfo_Level1:
		if value != nil {
			return value.Level1
		}
	case *UseInfo_Level2:
		if value != nil {
			return value.Level2
		}
	}
	return nil
}

type is_UseEnum_UseInfo interface {
	ndr.Marshaler
	ndr.Unmarshaler
	is_UseEnum_UseInfo()
}

func (o *UseEnum_UseInfo) NDRSwitchValue(sw uint32) uint32 {
	if o == nil {
		return uint32(0)
	}
	switch (interface{})(o.Value).(type) {
	case *UseInfo_Level0:
		return uint32(0)
	case *UseInfo_Level1:
		return uint32(1)
	case *UseInfo_Level2:
		return uint32(2)
	}
	return uint32(0)
}

func (o *UseEnum_UseInfo) MarshalUnionNDR(ctx context.Context, w ndr.Writer, sw uint32) error {
	if err := w.WriteSwitch(uint32(sw)); err != nil {
		return err
	}
	switch sw {
	case uint32(0):
		_o, _ := o.Value.(*UseInfo_Level0)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&UseInfo_Level0{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(1):
		_o, _ := o.Value.(*UseInfo_Level1)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&UseInfo_Level1{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(2):
		_o, _ := o.Value.(*UseInfo_Level2)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&UseInfo_Level2{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	default:
	}
	return nil
}

func (o *UseEnum_UseInfo) UnmarshalUnionNDR(ctx context.Context, w ndr.Reader, sw uint32) error {
	if err := w.ReadSwitch((*uint32)(&sw)); err != nil {
		return err
	}
	switch sw {
	case uint32(0):
		o.Value = &UseInfo_Level0{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(1):
		o.Value = &UseInfo_Level1{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(2):
		o.Value = &UseInfo_Level2{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	default:
	}
	return nil
}

// UseInfo_Level0 structure represents UseEnum_UseInfo RPC union arm.
//
// It has following labels: 0
type UseInfo_Level0 struct {
	Level0 *UseInfo0Container `idl:"name:Level0" json:"level0"`
}

func (*UseInfo_Level0) is_UseEnum_UseInfo() {}

func (o *UseInfo_Level0) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.Level0 != nil {
		_ptr_Level0 := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.Level0 != nil {
				if err := o.Level0.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&UseInfo0Container{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Level0, _ptr_Level0); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *UseInfo_Level0) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_Level0 := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.Level0 == nil {
			o.Level0 = &UseInfo0Container{}
		}
		if err := o.Level0.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_Level0 := func(ptr interface{}) { o.Level0 = *ptr.(**UseInfo0Container) }
	if err := w.ReadPointer(&o.Level0, _s_Level0, _ptr_Level0); err != nil {
		return err
	}
	return nil
}

// UseInfo_Level1 structure represents UseEnum_UseInfo RPC union arm.
//
// It has following labels: 1
type UseInfo_Level1 struct {
	Level1 *UseInfo1Container `idl:"name:Level1" json:"level1"`
}

func (*UseInfo_Level1) is_UseEnum_UseInfo() {}

func (o *UseInfo_Level1) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.Level1 != nil {
		_ptr_Level1 := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.Level1 != nil {
				if err := o.Level1.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&UseInfo1Container{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Level1, _ptr_Level1); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *UseInfo_Level1) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_Level1 := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.Level1 == nil {
			o.Level1 = &UseInfo1Container{}
		}
		if err := o.Level1.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_Level1 := func(ptr interface{}) { o.Level1 = *ptr.(**UseInfo1Container) }
	if err := w.ReadPointer(&o.Level1, _s_Level1, _ptr_Level1); err != nil {
		return err
	}
	return nil
}

// UseInfo_Level2 structure represents UseEnum_UseInfo RPC union arm.
//
// It has following labels: 2
type UseInfo_Level2 struct {
	Level2 *UseInfo2Container `idl:"name:Level2" json:"level2"`
}

func (*UseInfo_Level2) is_UseEnum_UseInfo() {}

func (o *UseInfo_Level2) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.Level2 != nil {
		_ptr_Level2 := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.Level2 != nil {
				if err := o.Level2.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&UseInfo2Container{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Level2, _ptr_Level2); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *UseInfo_Level2) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_Level2 := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.Level2 == nil {
			o.Level2 = &UseInfo2Container{}
		}
		if err := o.Level2.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_Level2 := func(ptr interface{}) { o.Level2 = *ptr.(**UseInfo2Container) }
	if err := w.ReadPointer(&o.Level2, _s_Level2, _ptr_Level2); err != nil {
		return err
	}
	return nil
}

// WorkstationUserInfo0Container structure represents WKSTA_USER_INFO_0_CONTAINER RPC structure.
//
// The WKSTA_USER_INFO_0_CONTAINER structure contains a value that indicates the number
// of entries that the NetrWkstaUserEnum method (section 3.2.4.3) returns, as well as
// a pointer to the buffer.
type WorkstationUserInfo0Container struct {
	// EntriesRead: The number of entries that the method returns.
	EntriesRead uint32 `idl:"name:EntriesRead" json:"entries_read"`
	// Buffer: The names of the user accounts logged onto the remote computer.
	Buffer []*WorkstationUserInfo0 `idl:"name:Buffer;size_is:(EntriesRead)" json:"buffer"`
}

func (o *WorkstationUserInfo0Container) xxx_PreparePayload(ctx context.Context) error {
	if o.Buffer != nil && o.EntriesRead == 0 {
		o.EntriesRead = uint32(len(o.Buffer))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *WorkstationUserInfo0Container) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.EntriesRead); err != nil {
		return err
	}
	if o.Buffer != nil || o.EntriesRead > 0 {
		_ptr_Buffer := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.EntriesRead)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.Buffer {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if o.Buffer[i1] != nil {
					if err := o.Buffer[i1].MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&WorkstationUserInfo0{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.Buffer); uint64(i1) < sizeInfo[0]; i1++ {
				if err := (&WorkstationUserInfo0{}).MarshalNDR(ctx, w); err != nil {
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
func (o *WorkstationUserInfo0Container) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.EntriesRead); err != nil {
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
		// XXX: for opaque unmarshaling
		if o.EntriesRead > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.EntriesRead)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Buffer", sizeInfo[0])
		}
		o.Buffer = make([]*WorkstationUserInfo0, sizeInfo[0])
		for i1 := range o.Buffer {
			i1 := i1
			if o.Buffer[i1] == nil {
				o.Buffer[i1] = &WorkstationUserInfo0{}
			}
			if err := o.Buffer[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		return nil
	})
	_s_Buffer := func(ptr interface{}) { o.Buffer = *ptr.(*[]*WorkstationUserInfo0) }
	if err := w.ReadPointer(&o.Buffer, _s_Buffer, _ptr_Buffer); err != nil {
		return err
	}
	return nil
}

// WorkstationUserInfo1Container structure represents WKSTA_USER_INFO_1_CONTAINER RPC structure.
//
// The WKSTA_USER_INFO_1_CONTAINER structure contains a value that indicates the number
// of entries that the NetrWkstaUserEnum method (section 3.2.4.3) returns, as well as
// a pointer to the buffer.
type WorkstationUserInfo1Container struct {
	// EntriesRead: The number of entries that the method returns.
	EntriesRead uint32 `idl:"name:EntriesRead" json:"entries_read"`
	// Buffer: Details about the user accounts logged onto the remote computer.
	Buffer []*WorkstationUserInfo1 `idl:"name:Buffer;size_is:(EntriesRead)" json:"buffer"`
}

func (o *WorkstationUserInfo1Container) xxx_PreparePayload(ctx context.Context) error {
	if o.Buffer != nil && o.EntriesRead == 0 {
		o.EntriesRead = uint32(len(o.Buffer))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *WorkstationUserInfo1Container) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.EntriesRead); err != nil {
		return err
	}
	if o.Buffer != nil || o.EntriesRead > 0 {
		_ptr_Buffer := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.EntriesRead)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.Buffer {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if o.Buffer[i1] != nil {
					if err := o.Buffer[i1].MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&WorkstationUserInfo1{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.Buffer); uint64(i1) < sizeInfo[0]; i1++ {
				if err := (&WorkstationUserInfo1{}).MarshalNDR(ctx, w); err != nil {
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
func (o *WorkstationUserInfo1Container) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.EntriesRead); err != nil {
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
		// XXX: for opaque unmarshaling
		if o.EntriesRead > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.EntriesRead)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Buffer", sizeInfo[0])
		}
		o.Buffer = make([]*WorkstationUserInfo1, sizeInfo[0])
		for i1 := range o.Buffer {
			i1 := i1
			if o.Buffer[i1] == nil {
				o.Buffer[i1] = &WorkstationUserInfo1{}
			}
			if err := o.Buffer[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		return nil
	})
	_s_Buffer := func(ptr interface{}) { o.Buffer = *ptr.(*[]*WorkstationUserInfo1) }
	if err := w.ReadPointer(&o.Buffer, _s_Buffer, _ptr_Buffer); err != nil {
		return err
	}
	return nil
}

// WorkstationUserEnum structure represents WKSTA_USER_ENUM_STRUCT RPC structure.
//
// The WKSTA_USER_ENUM_STRUCT structure is used by the NetrWkstaUserEnum method (section
// 3.2.4.3) to encapsulate the _WKSTA_USER_ENUM_UNION union.
type WorkstationUserEnum struct {
	// Level: The information level of the data, which determines the type of structure
	// that the method returns. It MUST be one of the following values.
	//
	//	+------------+-------------------------------------------------------------------------+
	//	|            |                                                                         |
	//	|   VALUE    |                                 MEANING                                 |
	//	|            |                                                                         |
	//	+------------+-------------------------------------------------------------------------+
	//	+------------+-------------------------------------------------------------------------+
	//	| 0x00000000 | Specifies the WKSTA_USER_INFO_0_CONTAINER structure (section 2.2.5.12). |
	//	+------------+-------------------------------------------------------------------------+
	//	| 0x00000001 | Specifies the WKSTA_USER_INFO_1_CONTAINER structure (section 2.2.5.13). |
	//	+------------+-------------------------------------------------------------------------+
	Level uint32 `idl:"name:Level" json:"level"`
	// WkstaUserInfo: A WKSTA_USER_INFO_0_CONTAINER structure or a WKSTA_USER_INFO_1_CONTAINER
	// structure.
	WorkstationUserInfo *WorkstationUserEnum_WorkstationUserInfo `idl:"name:WkstaUserInfo;switch_is:Level" json:"workstation_user_info"`
}

func (o *WorkstationUserEnum) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *WorkstationUserEnum) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.Level); err != nil {
		return err
	}
	_swWorkstationUserInfo := uint32(o.Level)
	if o.WorkstationUserInfo != nil {
		if err := o.WorkstationUserInfo.MarshalUnionNDR(ctx, w, _swWorkstationUserInfo); err != nil {
			return err
		}
	} else {
		if err := (&WorkstationUserEnum_WorkstationUserInfo{}).MarshalUnionNDR(ctx, w, _swWorkstationUserInfo); err != nil {
			return err
		}
	}
	return nil
}
func (o *WorkstationUserEnum) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.Level); err != nil {
		return err
	}
	if o.WorkstationUserInfo == nil {
		o.WorkstationUserInfo = &WorkstationUserEnum_WorkstationUserInfo{}
	}
	_swWorkstationUserInfo := uint32(o.Level)
	if err := o.WorkstationUserInfo.UnmarshalUnionNDR(ctx, w, _swWorkstationUserInfo); err != nil {
		return err
	}
	return nil
}

// WorkstationUserEnum_WorkstationUserInfo structure represents WKSTA_USER_ENUM_STRUCT union anonymous member.
//
// The WKSTA_USER_ENUM_STRUCT structure is used by the NetrWkstaUserEnum method (section
// 3.2.4.3) to encapsulate the _WKSTA_USER_ENUM_UNION union.
type WorkstationUserEnum_WorkstationUserInfo struct {
	// Types that are assignable to Value
	//
	// *WorkstationUserInfo_Level0
	// *WorkstationUserInfo_Level1
	Value is_WorkstationUserEnum_WorkstationUserInfo `json:"value"`
}

func (o *WorkstationUserEnum_WorkstationUserInfo) GetValue() any {
	if o == nil {
		return nil
	}
	switch value := (interface{})(o.Value).(type) {
	case *WorkstationUserInfo_Level0:
		if value != nil {
			return value.Level0
		}
	case *WorkstationUserInfo_Level1:
		if value != nil {
			return value.Level1
		}
	}
	return nil
}

type is_WorkstationUserEnum_WorkstationUserInfo interface {
	ndr.Marshaler
	ndr.Unmarshaler
	is_WorkstationUserEnum_WorkstationUserInfo()
}

func (o *WorkstationUserEnum_WorkstationUserInfo) NDRSwitchValue(sw uint32) uint32 {
	if o == nil {
		return uint32(0)
	}
	switch (interface{})(o.Value).(type) {
	case *WorkstationUserInfo_Level0:
		return uint32(0)
	case *WorkstationUserInfo_Level1:
		return uint32(1)
	}
	return uint32(0)
}

func (o *WorkstationUserEnum_WorkstationUserInfo) MarshalUnionNDR(ctx context.Context, w ndr.Writer, sw uint32) error {
	if err := w.WriteSwitch(uint32(sw)); err != nil {
		return err
	}
	switch sw {
	case uint32(0):
		_o, _ := o.Value.(*WorkstationUserInfo_Level0)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&WorkstationUserInfo_Level0{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(1):
		_o, _ := o.Value.(*WorkstationUserInfo_Level1)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&WorkstationUserInfo_Level1{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	default:
	}
	return nil
}

func (o *WorkstationUserEnum_WorkstationUserInfo) UnmarshalUnionNDR(ctx context.Context, w ndr.Reader, sw uint32) error {
	if err := w.ReadSwitch((*uint32)(&sw)); err != nil {
		return err
	}
	switch sw {
	case uint32(0):
		o.Value = &WorkstationUserInfo_Level0{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(1):
		o.Value = &WorkstationUserInfo_Level1{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	default:
	}
	return nil
}

// WorkstationUserInfo_Level0 structure represents WorkstationUserEnum_WorkstationUserInfo RPC union arm.
//
// It has following labels: 0
type WorkstationUserInfo_Level0 struct {
	Level0 *WorkstationUserInfo0Container `idl:"name:Level0" json:"level0"`
}

func (*WorkstationUserInfo_Level0) is_WorkstationUserEnum_WorkstationUserInfo() {}

func (o *WorkstationUserInfo_Level0) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.Level0 != nil {
		_ptr_Level0 := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.Level0 != nil {
				if err := o.Level0.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&WorkstationUserInfo0Container{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Level0, _ptr_Level0); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *WorkstationUserInfo_Level0) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_Level0 := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.Level0 == nil {
			o.Level0 = &WorkstationUserInfo0Container{}
		}
		if err := o.Level0.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_Level0 := func(ptr interface{}) { o.Level0 = *ptr.(**WorkstationUserInfo0Container) }
	if err := w.ReadPointer(&o.Level0, _s_Level0, _ptr_Level0); err != nil {
		return err
	}
	return nil
}

// WorkstationUserInfo_Level1 structure represents WorkstationUserEnum_WorkstationUserInfo RPC union arm.
//
// It has following labels: 1
type WorkstationUserInfo_Level1 struct {
	Level1 *WorkstationUserInfo1Container `idl:"name:Level1" json:"level1"`
}

func (*WorkstationUserInfo_Level1) is_WorkstationUserEnum_WorkstationUserInfo() {}

func (o *WorkstationUserInfo_Level1) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.Level1 != nil {
		_ptr_Level1 := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.Level1 != nil {
				if err := o.Level1.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&WorkstationUserInfo1Container{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Level1, _ptr_Level1); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *WorkstationUserInfo_Level1) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_Level1 := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.Level1 == nil {
			o.Level1 = &WorkstationUserInfo1Container{}
		}
		if err := o.Level1.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_Level1 := func(ptr interface{}) { o.Level1 = *ptr.(**WorkstationUserInfo1Container) }
	if err := w.ReadPointer(&o.Level1, _s_Level1, _ptr_Level1); err != nil {
		return err
	}
	return nil
}

// WorkstationTransportInfo0Container structure represents WKSTA_TRANSPORT_INFO_0_CONTAINER RPC structure.
//
// The WKSTA_TRANSPORT_INFO_0_CONTAINER structure is used by the NetrWkstaTransportEnum
// method (section 3.2.4.4) to specify the number of entries and a pointer to the base
// WKSTA_TRANSPORT_INFO_0 structure (section 2.2.5.8) that the method returns.
type WorkstationTransportInfo0Container struct {
	// EntriesRead: The number of entries that the method returns.
	EntriesRead uint32 `idl:"name:EntriesRead" json:"entries_read"`
	// Buffer: A pointer to the array of WKSTA_TRANSPORT_INFO_0 structures that specify
	// details about transport protocols.
	Buffer []*WorkstationTransportInfo0 `idl:"name:Buffer;size_is:(EntriesRead)" json:"buffer"`
}

func (o *WorkstationTransportInfo0Container) xxx_PreparePayload(ctx context.Context) error {
	if o.Buffer != nil && o.EntriesRead == 0 {
		o.EntriesRead = uint32(len(o.Buffer))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *WorkstationTransportInfo0Container) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.EntriesRead); err != nil {
		return err
	}
	if o.Buffer != nil || o.EntriesRead > 0 {
		_ptr_Buffer := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.EntriesRead)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.Buffer {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if o.Buffer[i1] != nil {
					if err := o.Buffer[i1].MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&WorkstationTransportInfo0{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.Buffer); uint64(i1) < sizeInfo[0]; i1++ {
				if err := (&WorkstationTransportInfo0{}).MarshalNDR(ctx, w); err != nil {
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
func (o *WorkstationTransportInfo0Container) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.EntriesRead); err != nil {
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
		// XXX: for opaque unmarshaling
		if o.EntriesRead > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.EntriesRead)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Buffer", sizeInfo[0])
		}
		o.Buffer = make([]*WorkstationTransportInfo0, sizeInfo[0])
		for i1 := range o.Buffer {
			i1 := i1
			if o.Buffer[i1] == nil {
				o.Buffer[i1] = &WorkstationTransportInfo0{}
			}
			if err := o.Buffer[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		return nil
	})
	_s_Buffer := func(ptr interface{}) { o.Buffer = *ptr.(*[]*WorkstationTransportInfo0) }
	if err := w.ReadPointer(&o.Buffer, _s_Buffer, _ptr_Buffer); err != nil {
		return err
	}
	return nil
}

// WorkstationTransportEnum structure represents WKSTA_TRANSPORT_ENUM_STRUCT RPC structure.
//
// The WKSTA_TRANSPORT_ENUM_STRUCT structure is used by the NetrWkstaTransportEnum method
// (section 3.2.4.4). The Level parameter in the submitted structure determines the
// information level of the data that the method returns.
type WorkstationTransportEnum struct {
	// Level: The data’s information level, which MUST be set to zero.
	Level uint32 `idl:"name:Level" json:"level"`
	// WkstaTransportInfo: A pointer to a WKSTA_TRANSPORT_INFO_0_CONTAINER structure (section
	// 2.2.5.15).
	WorkstationTransportInfo *WorkstationTransportEnum_WorkstationTransportInfo `idl:"name:WkstaTransportInfo;switch_is:Level" json:"workstation_transport_info"`
}

func (o *WorkstationTransportEnum) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *WorkstationTransportEnum) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.Level); err != nil {
		return err
	}
	_swWorkstationTransportInfo := uint32(o.Level)
	if o.WorkstationTransportInfo != nil {
		if err := o.WorkstationTransportInfo.MarshalUnionNDR(ctx, w, _swWorkstationTransportInfo); err != nil {
			return err
		}
	} else {
		if err := (&WorkstationTransportEnum_WorkstationTransportInfo{}).MarshalUnionNDR(ctx, w, _swWorkstationTransportInfo); err != nil {
			return err
		}
	}
	return nil
}
func (o *WorkstationTransportEnum) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.Level); err != nil {
		return err
	}
	if o.WorkstationTransportInfo == nil {
		o.WorkstationTransportInfo = &WorkstationTransportEnum_WorkstationTransportInfo{}
	}
	_swWorkstationTransportInfo := uint32(o.Level)
	if err := o.WorkstationTransportInfo.UnmarshalUnionNDR(ctx, w, _swWorkstationTransportInfo); err != nil {
		return err
	}
	return nil
}

// WorkstationTransportEnum_WorkstationTransportInfo structure represents WKSTA_TRANSPORT_ENUM_STRUCT union anonymous member.
//
// The WKSTA_TRANSPORT_ENUM_STRUCT structure is used by the NetrWkstaTransportEnum method
// (section 3.2.4.4). The Level parameter in the submitted structure determines the
// information level of the data that the method returns.
type WorkstationTransportEnum_WorkstationTransportInfo struct {
	// Types that are assignable to Value
	//
	// *WorkstationTransportInfo_Level0
	Value is_WorkstationTransportEnum_WorkstationTransportInfo `json:"value"`
}

func (o *WorkstationTransportEnum_WorkstationTransportInfo) GetValue() any {
	if o == nil {
		return nil
	}
	switch value := (interface{})(o.Value).(type) {
	case *WorkstationTransportInfo_Level0:
		if value != nil {
			return value.Level0
		}
	}
	return nil
}

type is_WorkstationTransportEnum_WorkstationTransportInfo interface {
	ndr.Marshaler
	ndr.Unmarshaler
	is_WorkstationTransportEnum_WorkstationTransportInfo()
}

func (o *WorkstationTransportEnum_WorkstationTransportInfo) NDRSwitchValue(sw uint32) uint32 {
	if o == nil {
		return uint32(0)
	}
	switch (interface{})(o.Value).(type) {
	case *WorkstationTransportInfo_Level0:
		return uint32(0)
	}
	return uint32(0)
}

func (o *WorkstationTransportEnum_WorkstationTransportInfo) MarshalUnionNDR(ctx context.Context, w ndr.Writer, sw uint32) error {
	if err := w.WriteSwitch(uint32(sw)); err != nil {
		return err
	}
	switch sw {
	case uint32(0):
		_o, _ := o.Value.(*WorkstationTransportInfo_Level0)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&WorkstationTransportInfo_Level0{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	default:
	}
	return nil
}

func (o *WorkstationTransportEnum_WorkstationTransportInfo) UnmarshalUnionNDR(ctx context.Context, w ndr.Reader, sw uint32) error {
	if err := w.ReadSwitch((*uint32)(&sw)); err != nil {
		return err
	}
	switch sw {
	case uint32(0):
		o.Value = &WorkstationTransportInfo_Level0{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	default:
	}
	return nil
}

// WorkstationTransportInfo_Level0 structure represents WorkstationTransportEnum_WorkstationTransportInfo RPC union arm.
//
// It has following labels: 0
type WorkstationTransportInfo_Level0 struct {
	Level0 *WorkstationTransportInfo0Container `idl:"name:Level0" json:"level0"`
}

func (*WorkstationTransportInfo_Level0) is_WorkstationTransportEnum_WorkstationTransportInfo() {}

func (o *WorkstationTransportInfo_Level0) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.Level0 != nil {
		_ptr_Level0 := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.Level0 != nil {
				if err := o.Level0.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&WorkstationTransportInfo0Container{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Level0, _ptr_Level0); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *WorkstationTransportInfo_Level0) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_Level0 := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.Level0 == nil {
			o.Level0 = &WorkstationTransportInfo0Container{}
		}
		if err := o.Level0.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_Level0 := func(ptr interface{}) { o.Level0 = *ptr.(**WorkstationTransportInfo0Container) }
	if err := w.ReadPointer(&o.Level0, _s_Level0, _ptr_Level0); err != nil {
		return err
	}
	return nil
}

// UserPassword structure represents JOINPR_USER_PASSWORD RPC structure.
//
// The JOINPR_USER_PASSWORD structure represents a decrypted password in the Buffer
// member of a JOINPR_ENCRYPTED_USER_PASSWORD structure (section 2.2.5.18).
type UserPassword struct {
	// Obfuscator: An array of unsigned characters that contains a salt, which is filled
	// with random bytes by the caller.
	Obfuscator []byte `idl:"name:Obfuscator" json:"obfuscator"`
	// Buffer: A cleartext string of no more than JOIN_MAX_PASSWORD_LENGTH (section 2.2.1.1)
	// UTF-16 characters in little-endian order. The start of the string MUST be Length
	// number of bytes from the end of the buffer. The unused portion of the buffer contains
	// indeterminate values.
	Buffer []uint16 `idl:"name:Buffer" json:"buffer"`
	// Length: An unsigned integer, in little-endian order, that specifies the length in
	// bytes of the cleartext string in the Buffer member.
	Length uint32 `idl:"name:Length" json:"length"`
}

func (o *UserPassword) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *UserPassword) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	for i1 := range o.Obfuscator {
		i1 := i1
		if uint64(i1) >= 8 {
			break
		}
		if err := w.WriteData(o.Obfuscator[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.Obfuscator); uint64(i1) < 8; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	for i1 := range o.Buffer {
		i1 := i1
		if uint64(i1) >= 256 {
			break
		}
		if err := w.WriteData(o.Buffer[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.Buffer); uint64(i1) < 256; i1++ {
		if err := w.WriteData(uint16(0)); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.Length); err != nil {
		return err
	}
	return nil
}
func (o *UserPassword) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	o.Obfuscator = make([]byte, 8)
	for i1 := range o.Obfuscator {
		i1 := i1
		if err := w.ReadData(&o.Obfuscator[i1]); err != nil {
			return err
		}
	}
	o.Buffer = make([]uint16, 256)
	for i1 := range o.Buffer {
		i1 := i1
		if err := w.ReadData(&o.Buffer[i1]); err != nil {
			return err
		}
	}
	if err := w.ReadData(&o.Length); err != nil {
		return err
	}
	return nil
}

// EncryptedUserPassword structure represents JOINPR_ENCRYPTED_USER_PASSWORD RPC structure.
//
// The JOINPR_ENCRYPTED_USER_PASSWORD structure is the container for a password during
// the encoding, encryption, decryption and decoding process.
type EncryptedUserPassword struct {
	// Buffer: An array of bytes that contains a JOINPR_USER_PASSWORD structure (section
	// 2.2.5.17).
	//
	// The sections that follow specify the encoding, encryption, decryption, and decoding
	// of a password. (Encoding and encryption are performed by the client, but their explanations
	// are included for completeness and to facilitate the reader's understanding of server
	// (2) message processing.) The server decrypts and decodes a Buffer structure to extract
	// the cleartext password.
	//
	// The encoding, encryption, decryption, and decoding of a password requires the following
	// steps:
	//
	// *
	//
	// Encoding the cleartext password (section 2.2.5.18.1 ( 6613c3dc-744f-424c-b652-7f8804370c85
	// ) ).
	//
	// *
	//
	// Initializing *JOINPR_USER_PASSWORD* with the result of step 1 (section 2.2.5.18.2
	// ( fabd6340-5c6d-437d-b50a-0bc94340bcff ) ).
	//
	// *
	//
	// Initializing *JOINPR_ENCRYPTED_USER_PASSWORD.Buffer* with the encrypted result of
	// step 2, and subsequently decrypting *JOINPR_ENCRYPTED_USER_PASSWORD.Buffer* (section
	// 2.2.5.18.3 ( bead3d50-0440-448a-90e3-e478d74c7555 ) ).
	//
	// *
	//
	// Decoding the result of step 3, as a *JOINPR_USER_PASSWORD* structure, to recover
	// the cleartext password (section 2.2.5.18.4 ( b8d68e41-76a1-4895-97e5-b1f20a86fb04
	// ) ).
	Buffer []byte `idl:"name:Buffer" json:"buffer"`
}

func (o *EncryptedUserPassword) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *EncryptedUserPassword) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	for i1 := range o.Buffer {
		i1 := i1
		if uint64(i1) >= 524 {
			break
		}
		if err := w.WriteData(o.Buffer[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.Buffer); uint64(i1) < 524; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	return nil
}
func (o *EncryptedUserPassword) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	o.Buffer = make([]byte, 524)
	for i1 := range o.Buffer {
		i1 := i1
		if err := w.ReadData(&o.Buffer[i1]); err != nil {
			return err
		}
	}
	return nil
}

// UnicodeString structure represents UNICODE_STRING RPC structure.
//
// The UNICODE_STRING structure specifies a Unicode string.
type UnicodeString struct {
	// Length: The length, in bytes, of the string pointed to by the Buffer member, not
	// including the terminating null character, if any. This value MUST be a multiple of
	// 2.
	Length uint16 `idl:"name:Length" json:"length"`
	// MaximumLength: The total size, in bytes, of the Buffer. If this value is not a multiple
	// of 2, the server (2) MUST decrement this value by 1. This value MUST NOT be less
	// than Length.
	MaximumLength uint16 `idl:"name:MaximumLength" json:"maximum_length"`
	// Buffer: The Unicode UTF-8 string. If the MaximumLength value is greater than zero,
	// this field MUST contain a non-null character. Buffer can contain a terminating null
	// character.
	Buffer []uint16 `idl:"name:Buffer;size_is:((MaximumLength/2));length_is:((Length/2))" json:"buffer"`
}

func (o *UnicodeString) xxx_PreparePayload(ctx context.Context) error {
	if o.Buffer != nil && o.MaximumLength == 0 {
		o.MaximumLength = uint16((len(o.Buffer) * 2))
	}
	if o.Buffer != nil && o.Length == 0 {
		o.Length = uint16((len(o.Buffer) * 2))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *UnicodeString) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
	if o.Buffer != nil || (o.MaximumLength/2) > 0 {
		_ptr_Buffer := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64((o.MaximumLength / 2))
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			dimLength1 := uint64((o.Length / 2))
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
				if err := w.WriteData(uint16(0)); err != nil {
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
func (o *UnicodeString) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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
		o.Buffer = make([]uint16, sizeInfo[0])
		for i1 := range o.Buffer {
			i1 := i1
			if err := w.ReadData(&o.Buffer[i1]); err != nil {
				return err
			}
		}
		return nil
	})
	_s_Buffer := func(ptr interface{}) { o.Buffer = *ptr.(*[]uint16) }
	if err := w.ReadPointer(&o.Buffer, _s_Buffer, _ptr_Buffer); err != nil {
		return err
	}
	return nil
}

// NetComputerNameArray structure represents NET_COMPUTER_NAME_ARRAY RPC structure.
//
// The NET_COMPUTER_NAME_ARRAY structure specifies the number of names associated with
// a computer and a buffer containing the names.
type NetComputerNameArray struct {
	// EntryCount: The number of entries that the method call returns.
	EntryCount uint32 `idl:"name:EntryCount" json:"entry_count"`
	// ComputerNames: The names as an array of UNICODE_STRING structures (section 2.2.5.20)
	// that are associated with a machine.
	ComputerNames []*UnicodeString `idl:"name:ComputerNames;size_is:(EntryCount)" json:"computer_names"`
}

func (o *NetComputerNameArray) xxx_PreparePayload(ctx context.Context) error {
	if o.ComputerNames != nil && o.EntryCount == 0 {
		o.EntryCount = uint32(len(o.ComputerNames))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *NetComputerNameArray) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.EntryCount); err != nil {
		return err
	}
	if o.ComputerNames != nil || o.EntryCount > 0 {
		_ptr_ComputerNames := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.EntryCount)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.ComputerNames {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if o.ComputerNames[i1] != nil {
					if err := o.ComputerNames[i1].MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&UnicodeString{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.ComputerNames); uint64(i1) < sizeInfo[0]; i1++ {
				if err := (&UnicodeString{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.ComputerNames, _ptr_ComputerNames); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *NetComputerNameArray) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.EntryCount); err != nil {
		return err
	}
	_ptr_ComputerNames := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.EntryCount > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.EntryCount)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.ComputerNames", sizeInfo[0])
		}
		o.ComputerNames = make([]*UnicodeString, sizeInfo[0])
		for i1 := range o.ComputerNames {
			i1 := i1
			if o.ComputerNames[i1] == nil {
				o.ComputerNames[i1] = &UnicodeString{}
			}
			if err := o.ComputerNames[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		return nil
	})
	_s_ComputerNames := func(ptr interface{}) { o.ComputerNames = *ptr.(*[]*UnicodeString) }
	if err := w.ReadPointer(&o.ComputerNames, _s_ComputerNames, _ptr_ComputerNames); err != nil {
		return err
	}
	return nil
}

type xxx_DefaultWkssvcClient struct {
	cc dcerpc.Conn
}

func (o *xxx_DefaultWkssvcClient) GetInfo(ctx context.Context, in *GetInfoRequest, opts ...dcerpc.CallOption) (*GetInfoResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetInfoResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultWkssvcClient) SetInfo(ctx context.Context, in *SetInfoRequest, opts ...dcerpc.CallOption) (*SetInfoResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &SetInfoResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultWkssvcClient) UserEnum(ctx context.Context, in *UserEnumRequest, opts ...dcerpc.CallOption) (*UserEnumResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &UserEnumResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultWkssvcClient) TransportEnum(ctx context.Context, in *TransportEnumRequest, opts ...dcerpc.CallOption) (*TransportEnumResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &TransportEnumResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultWkssvcClient) TransportAdd(ctx context.Context, in *TransportAddRequest, opts ...dcerpc.CallOption) (*TransportAddResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &TransportAddResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultWkssvcClient) TransportDelete(ctx context.Context, in *TransportDeleteRequest, opts ...dcerpc.CallOption) (*TransportDeleteResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &TransportDeleteResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultWkssvcClient) UseAdd(ctx context.Context, in *UseAddRequest, opts ...dcerpc.CallOption) (*UseAddResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &UseAddResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultWkssvcClient) UseGetInfo(ctx context.Context, in *UseGetInfoRequest, opts ...dcerpc.CallOption) (*UseGetInfoResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &UseGetInfoResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultWkssvcClient) UseDelete(ctx context.Context, in *UseDeleteRequest, opts ...dcerpc.CallOption) (*UseDeleteResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &UseDeleteResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultWkssvcClient) UseEnum(ctx context.Context, in *UseEnumRequest, opts ...dcerpc.CallOption) (*UseEnumResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &UseEnumResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultWkssvcClient) WorkstationStatisticsGet(ctx context.Context, in *WorkstationStatisticsGetRequest, opts ...dcerpc.CallOption) (*WorkstationStatisticsGetResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &WorkstationStatisticsGetResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultWkssvcClient) GetJoinInformation(ctx context.Context, in *GetJoinInformationRequest, opts ...dcerpc.CallOption) (*GetJoinInformationResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetJoinInformationResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultWkssvcClient) JoinDomain2(ctx context.Context, in *JoinDomain2Request, opts ...dcerpc.CallOption) (*JoinDomain2Response, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &JoinDomain2Response{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultWkssvcClient) UnjoinDomain2(ctx context.Context, in *UnjoinDomain2Request, opts ...dcerpc.CallOption) (*UnjoinDomain2Response, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &UnjoinDomain2Response{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultWkssvcClient) RenameMachineInDomain2(ctx context.Context, in *RenameMachineInDomain2Request, opts ...dcerpc.CallOption) (*RenameMachineInDomain2Response, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &RenameMachineInDomain2Response{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultWkssvcClient) ValidateName2(ctx context.Context, in *ValidateName2Request, opts ...dcerpc.CallOption) (*ValidateName2Response, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &ValidateName2Response{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultWkssvcClient) GetJoinableOUs2(ctx context.Context, in *GetJoinableOUs2Request, opts ...dcerpc.CallOption) (*GetJoinableOUs2Response, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetJoinableOUs2Response{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultWkssvcClient) AddAlternateComputerName(ctx context.Context, in *AddAlternateComputerNameRequest, opts ...dcerpc.CallOption) (*AddAlternateComputerNameResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &AddAlternateComputerNameResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultWkssvcClient) RemoveAlternateComputerName(ctx context.Context, in *RemoveAlternateComputerNameRequest, opts ...dcerpc.CallOption) (*RemoveAlternateComputerNameResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &RemoveAlternateComputerNameResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultWkssvcClient) SetPrimaryComputerName(ctx context.Context, in *SetPrimaryComputerNameRequest, opts ...dcerpc.CallOption) (*SetPrimaryComputerNameResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &SetPrimaryComputerNameResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultWkssvcClient) EnumerateComputerNames(ctx context.Context, in *EnumerateComputerNamesRequest, opts ...dcerpc.CallOption) (*EnumerateComputerNamesResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &EnumerateComputerNamesResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultWkssvcClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultWkssvcClient) Conn() dcerpc.Conn {
	return o.cc
}

func NewWkssvcClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (WkssvcClient, error) {
	cc, err := cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(WkssvcSyntaxV1_0))...)
	if err != nil {
		return nil, err
	}
	return &xxx_DefaultWkssvcClient{cc: cc}, nil
}

// xxx_GetInfoOperation structure represents the NetrWkstaGetInfo operation
type xxx_GetInfoOperation struct {
	ServerName      string           `idl:"name:ServerName;string;pointer:unique" json:"server_name"`
	Level           uint32           `idl:"name:Level" json:"level"`
	WorkstationInfo *WorkstationInfo `idl:"name:WkstaInfo;switch_is:Level" json:"workstation_info"`
	Return          uint32           `idl:"name:Return" json:"return"`
}

func (o *xxx_GetInfoOperation) OpNum() int { return 0 }

func (o *xxx_GetInfoOperation) OpName() string { return "/wkssvc/v1/NetrWkstaGetInfo" }

func (o *xxx_GetInfoOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetInfoOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// ServerName {in} (1:{handle, string, pointer=unique, alias=WKSSVC_IDENTIFY_HANDLE}*(1)[dim:0,string,null](wchar))
	{
		if o.ServerName != "" {
			_ptr_ServerName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.ServerName); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.ServerName, _ptr_ServerName); err != nil {
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
	// Level {in} (1:(uint32))
	{
		if err := w.WriteData(o.Level); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetInfoOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// ServerName {in} (1:{handle, string, pointer=unique, alias=WKSSVC_IDENTIFY_HANDLE}*(1)[dim:0,string,null](wchar))
	{
		_ptr_ServerName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.ServerName); err != nil {
				return err
			}
			return nil
		})
		_s_ServerName := func(ptr interface{}) { o.ServerName = *ptr.(*string) }
		if err := w.ReadPointer(&o.ServerName, _s_ServerName, _ptr_ServerName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Level {in} (1:(uint32))
	{
		if err := w.ReadData(&o.Level); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetInfoOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetInfoOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// WkstaInfo {out} (1:{switch_type={}(uint32), alias=LPWKSTA_INFO}*(1))(2:{switch_type={}(uint32), alias=WKSTA_INFO}(union))
	{
		_swWorkstationInfo := uint32(o.Level)
		if o.WorkstationInfo != nil {
			if err := o.WorkstationInfo.MarshalUnionNDR(ctx, w, _swWorkstationInfo); err != nil {
				return err
			}
		} else {
			if err := (&WorkstationInfo{}).MarshalUnionNDR(ctx, w, _swWorkstationInfo); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:(uint32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetInfoOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// WkstaInfo {out} (1:{switch_type={}(uint32), alias=LPWKSTA_INFO,pointer=ref}*(1))(2:{switch_type={}(uint32), alias=WKSTA_INFO}(union))
	{
		if o.WorkstationInfo == nil {
			o.WorkstationInfo = &WorkstationInfo{}
		}
		_swWorkstationInfo := uint32(o.Level)
		if err := o.WorkstationInfo.UnmarshalUnionNDR(ctx, w, _swWorkstationInfo); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// GetInfoRequest structure represents the NetrWkstaGetInfo operation request
type GetInfoRequest struct {
	// ServerName: A WKSSVC_IDENTIFY_HANDLE structure (section 2.2.2.1) that identifies
	// the server (2). The client MUST map this structure to an RPC binding handle ([C706]
	// sections 4.3.5 and 5.1.5.2). The server MUST ignore this parameter.
	ServerName string `idl:"name:ServerName;string;pointer:unique" json:"server_name"`
	// Level: The information level of the data. This parameter MUST be one of the following
	// values.
	//
	//	+------------+-----------------------------------------------------------------------+
	//	|            |                                                                       |
	//	|   VALUE    |                                MEANING                                |
	//	|            |                                                                       |
	//	+------------+-----------------------------------------------------------------------+
	//	+------------+-----------------------------------------------------------------------+
	//	| 0x00000064 | Details are returned in a WKSTA_INFO_100 structure (section 2.2.5.1). |
	//	+------------+-----------------------------------------------------------------------+
	//	| 0x00000065 | Details are returned in a WKSTA_INFO_101 structure (section 2.2.5.2). |
	//	+------------+-----------------------------------------------------------------------+
	//	| 0x00000066 | Details are returned in a WKSTA_INFO_102 structure (section 2.2.5.3). |
	//	+------------+-----------------------------------------------------------------------+
	//	| 0x000001F6 | Details are returned in a WKSTA_INFO_502 structure (section 2.2.5.4). |
	//	+------------+-----------------------------------------------------------------------+
	Level uint32 `idl:"name:Level" json:"level"`
}

func (o *GetInfoRequest) xxx_ToOp(ctx context.Context) *xxx_GetInfoOperation {
	if o == nil {
		return &xxx_GetInfoOperation{}
	}
	return &xxx_GetInfoOperation{
		ServerName: o.ServerName,
		Level:      o.Level,
	}
}

func (o *GetInfoRequest) xxx_FromOp(ctx context.Context, op *xxx_GetInfoOperation) {
	if o == nil {
		return
	}
	o.ServerName = op.ServerName
	o.Level = op.Level
}
func (o *GetInfoRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetInfoRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetInfoOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetInfoResponse structure represents the NetrWkstaGetInfo operation response
type GetInfoResponse struct {
	// WkstaInfo: A pointer to the buffer that receives the data. The format of this data
	// depends on the value of the level parameter.
	WorkstationInfo *WorkstationInfo `idl:"name:WkstaInfo;switch_is:Level" json:"workstation_info"`
	// Return: The NetrWkstaGetInfo return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *GetInfoResponse) xxx_ToOp(ctx context.Context) *xxx_GetInfoOperation {
	if o == nil {
		return &xxx_GetInfoOperation{}
	}
	return &xxx_GetInfoOperation{
		WorkstationInfo: o.WorkstationInfo,
		Return:          o.Return,
	}
}

func (o *GetInfoResponse) xxx_FromOp(ctx context.Context, op *xxx_GetInfoOperation) {
	if o == nil {
		return
	}
	o.WorkstationInfo = op.WorkstationInfo
	o.Return = op.Return
}
func (o *GetInfoResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetInfoResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetInfoOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetInfoOperation structure represents the NetrWkstaSetInfo operation
type xxx_SetInfoOperation struct {
	ServerName      string           `idl:"name:ServerName;string;pointer:unique" json:"server_name"`
	Level           uint32           `idl:"name:Level" json:"level"`
	WorkstationInfo *WorkstationInfo `idl:"name:WkstaInfo;switch_is:Level" json:"workstation_info"`
	ErrorParameter  uint32           `idl:"name:ErrorParameter;pointer:unique" json:"error_parameter"`
	Return          uint32           `idl:"name:Return" json:"return"`
}

func (o *xxx_SetInfoOperation) OpNum() int { return 1 }

func (o *xxx_SetInfoOperation) OpName() string { return "/wkssvc/v1/NetrWkstaSetInfo" }

func (o *xxx_SetInfoOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetInfoOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// ServerName {in} (1:{handle, string, pointer=unique, alias=WKSSVC_IDENTIFY_HANDLE}*(1)[dim:0,string,null](wchar))
	{
		if o.ServerName != "" {
			_ptr_ServerName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.ServerName); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.ServerName, _ptr_ServerName); err != nil {
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
	// Level {in} (1:(uint32))
	{
		if err := w.WriteData(o.Level); err != nil {
			return err
		}
	}
	// WkstaInfo {in} (1:{switch_type={}(uint32), alias=LPWKSTA_INFO}*(1))(2:{switch_type={}(uint32), alias=WKSTA_INFO}(union))
	{
		_swWorkstationInfo := uint32(o.Level)
		if o.WorkstationInfo != nil {
			if err := o.WorkstationInfo.MarshalUnionNDR(ctx, w, _swWorkstationInfo); err != nil {
				return err
			}
		} else {
			if err := (&WorkstationInfo{}).MarshalUnionNDR(ctx, w, _swWorkstationInfo); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// ErrorParameter {in, out} (1:{pointer=unique}*(1)(uint32))
	{
		_ptr_ErrorParameter := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := w.WriteData(o.ErrorParameter); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.ErrorParameter, _ptr_ErrorParameter); err != nil {
			return err
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetInfoOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// ServerName {in} (1:{handle, string, pointer=unique, alias=WKSSVC_IDENTIFY_HANDLE}*(1)[dim:0,string,null](wchar))
	{
		_ptr_ServerName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.ServerName); err != nil {
				return err
			}
			return nil
		})
		_s_ServerName := func(ptr interface{}) { o.ServerName = *ptr.(*string) }
		if err := w.ReadPointer(&o.ServerName, _s_ServerName, _ptr_ServerName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Level {in} (1:(uint32))
	{
		if err := w.ReadData(&o.Level); err != nil {
			return err
		}
	}
	// WkstaInfo {in} (1:{switch_type={}(uint32), alias=LPWKSTA_INFO,pointer=ref}*(1))(2:{switch_type={}(uint32), alias=WKSTA_INFO}(union))
	{
		if o.WorkstationInfo == nil {
			o.WorkstationInfo = &WorkstationInfo{}
		}
		_swWorkstationInfo := uint32(o.Level)
		if err := o.WorkstationInfo.UnmarshalUnionNDR(ctx, w, _swWorkstationInfo); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// ErrorParameter {in, out} (1:{pointer=unique}*(1)(uint32))
	{
		_ptr_ErrorParameter := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := w.ReadData(&o.ErrorParameter); err != nil {
				return err
			}
			return nil
		})
		_s_ErrorParameter := func(ptr interface{}) { o.ErrorParameter = *ptr.(*uint32) }
		if err := w.ReadPointer(&o.ErrorParameter, _s_ErrorParameter, _ptr_ErrorParameter); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetInfoOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetInfoOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// ErrorParameter {in, out} (1:{pointer=unique}*(1)(uint32))
	{
		_ptr_ErrorParameter := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := w.WriteData(o.ErrorParameter); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.ErrorParameter, _ptr_ErrorParameter); err != nil {
			return err
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:(uint32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetInfoOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// ErrorParameter {in, out} (1:{pointer=unique}*(1)(uint32))
	{
		_ptr_ErrorParameter := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := w.ReadData(&o.ErrorParameter); err != nil {
				return err
			}
			return nil
		})
		_s_ErrorParameter := func(ptr interface{}) { o.ErrorParameter = *ptr.(*uint32) }
		if err := w.ReadPointer(&o.ErrorParameter, _s_ErrorParameter, _ptr_ErrorParameter); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// SetInfoRequest structure represents the NetrWkstaSetInfo operation request
type SetInfoRequest struct {
	// ServerName: A WKSSVC_IDENTIFY_HANDLE structure (section 2.2.2.1) that identifies
	// the server (2). The client MUST map this structure to an RPC binding handle ([C706]
	// sections 4.3.5 and 5.1.5.2). The server MUST ignore this parameter.
	ServerName string `idl:"name:ServerName;string;pointer:unique" json:"server_name"`
	// Level: The information level of the data. This parameter SHOULD be one of the following
	// values.
	//
	//	+------------+----------------------------------------------------------------------------------+
	//	|            |                                                                                  |
	//	|   VALUE    |                                     MEANING                                      |
	//	|            |                                                                                  |
	//	+------------+----------------------------------------------------------------------------------+
	//	+------------+----------------------------------------------------------------------------------+
	//	| 0x000001F6 | The WkstaInfo parameter points to a WKSTA_INFO_502 structure (section 2.2.5.4).  |
	//	+------------+----------------------------------------------------------------------------------+
	//	| 0x000003F5 | The WkstaInfo parameter points to a WKSTA_INFO_1013 structure (section 2.2.5.5). |
	//	+------------+----------------------------------------------------------------------------------+
	//	| 0x000003FA | The WkstaInfo parameter points to a WKSTA_INFO_1018 structure (section 2.2.5.6). |
	//	+------------+----------------------------------------------------------------------------------+
	//	| 0x00000416 | The WkstaInfo parameter points to a WKSTA_INFO_1046 structure (section 2.2.5.7). |
	//	+------------+----------------------------------------------------------------------------------+
	Level uint32 `idl:"name:Level" json:"level"`
	// WkstaInfo: A pointer to a buffer that specifies the data. The format of this data
	// depends on the value of the Level parameter.
	WorkstationInfo *WorkstationInfo `idl:"name:WkstaInfo;switch_is:Level" json:"workstation_info"`
	// ErrorParameter: A pointer to a value that receives an unsigned 32-bit integer. This
	// parameter is meaningful only if the method returns ERROR_INVALID_PARAMETER and Level
	// is equal to one of the values in the preceding table.
	ErrorParameter uint32 `idl:"name:ErrorParameter;pointer:unique" json:"error_parameter"`
}

func (o *SetInfoRequest) xxx_ToOp(ctx context.Context) *xxx_SetInfoOperation {
	if o == nil {
		return &xxx_SetInfoOperation{}
	}
	return &xxx_SetInfoOperation{
		ServerName:      o.ServerName,
		Level:           o.Level,
		WorkstationInfo: o.WorkstationInfo,
		ErrorParameter:  o.ErrorParameter,
	}
}

func (o *SetInfoRequest) xxx_FromOp(ctx context.Context, op *xxx_SetInfoOperation) {
	if o == nil {
		return
	}
	o.ServerName = op.ServerName
	o.Level = op.Level
	o.WorkstationInfo = op.WorkstationInfo
	o.ErrorParameter = op.ErrorParameter
}
func (o *SetInfoRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *SetInfoRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetInfoOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetInfoResponse structure represents the NetrWkstaSetInfo operation response
type SetInfoResponse struct {
	// ErrorParameter: A pointer to a value that receives an unsigned 32-bit integer. This
	// parameter is meaningful only if the method returns ERROR_INVALID_PARAMETER and Level
	// is equal to one of the values in the preceding table.
	ErrorParameter uint32 `idl:"name:ErrorParameter;pointer:unique" json:"error_parameter"`
	// Return: The NetrWkstaSetInfo return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *SetInfoResponse) xxx_ToOp(ctx context.Context) *xxx_SetInfoOperation {
	if o == nil {
		return &xxx_SetInfoOperation{}
	}
	return &xxx_SetInfoOperation{
		ErrorParameter: o.ErrorParameter,
		Return:         o.Return,
	}
}

func (o *SetInfoResponse) xxx_FromOp(ctx context.Context, op *xxx_SetInfoOperation) {
	if o == nil {
		return
	}
	o.ErrorParameter = op.ErrorParameter
	o.Return = op.Return
}
func (o *SetInfoResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *SetInfoResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetInfoOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_UserEnumOperation structure represents the NetrWkstaUserEnum operation
type xxx_UserEnumOperation struct {
	ServerName             string               `idl:"name:ServerName;string;pointer:unique" json:"server_name"`
	UserInfo               *WorkstationUserEnum `idl:"name:UserInfo" json:"user_info"`
	PreferredMaximumLength uint32               `idl:"name:PreferredMaximumLength" json:"preferred_maximum_length"`
	TotalEntries           uint32               `idl:"name:TotalEntries" json:"total_entries"`
	Resume                 uint32               `idl:"name:ResumeHandle;pointer:unique" json:"resume"`
	Return                 uint32               `idl:"name:Return" json:"return"`
}

func (o *xxx_UserEnumOperation) OpNum() int { return 2 }

func (o *xxx_UserEnumOperation) OpName() string { return "/wkssvc/v1/NetrWkstaUserEnum" }

func (o *xxx_UserEnumOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_UserEnumOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// ServerName {in} (1:{handle, string, pointer=unique, alias=WKSSVC_IDENTIFY_HANDLE}*(1)[dim:0,string,null](wchar))
	{
		if o.ServerName != "" {
			_ptr_ServerName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.ServerName); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.ServerName, _ptr_ServerName); err != nil {
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
	// UserInfo {in, out} (1:{alias=LPWKSTA_USER_ENUM_STRUCT}*(1))(2:{alias=WKSTA_USER_ENUM_STRUCT}(struct))
	{
		if o.UserInfo != nil {
			if err := o.UserInfo.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&WorkstationUserEnum{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// PreferredMaximumLength {in} (1:(uint32))
	{
		if err := w.WriteData(o.PreferredMaximumLength); err != nil {
			return err
		}
	}
	// ResumeHandle {in, out} (1:{pointer=unique}*(1)(uint32))
	{
		_ptr_ResumeHandle := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := w.WriteData(o.Resume); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.Resume, _ptr_ResumeHandle); err != nil {
			return err
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_UserEnumOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// ServerName {in} (1:{handle, string, pointer=unique, alias=WKSSVC_IDENTIFY_HANDLE}*(1)[dim:0,string,null](wchar))
	{
		_ptr_ServerName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.ServerName); err != nil {
				return err
			}
			return nil
		})
		_s_ServerName := func(ptr interface{}) { o.ServerName = *ptr.(*string) }
		if err := w.ReadPointer(&o.ServerName, _s_ServerName, _ptr_ServerName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// UserInfo {in, out} (1:{alias=LPWKSTA_USER_ENUM_STRUCT,pointer=ref}*(1))(2:{alias=WKSTA_USER_ENUM_STRUCT}(struct))
	{
		if o.UserInfo == nil {
			o.UserInfo = &WorkstationUserEnum{}
		}
		if err := o.UserInfo.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// PreferredMaximumLength {in} (1:(uint32))
	{
		if err := w.ReadData(&o.PreferredMaximumLength); err != nil {
			return err
		}
	}
	// ResumeHandle {in, out} (1:{pointer=unique}*(1)(uint32))
	{
		_ptr_ResumeHandle := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := w.ReadData(&o.Resume); err != nil {
				return err
			}
			return nil
		})
		_s_ResumeHandle := func(ptr interface{}) { o.Resume = *ptr.(*uint32) }
		if err := w.ReadPointer(&o.Resume, _s_ResumeHandle, _ptr_ResumeHandle); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_UserEnumOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_UserEnumOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// UserInfo {in, out} (1:{alias=LPWKSTA_USER_ENUM_STRUCT}*(1))(2:{alias=WKSTA_USER_ENUM_STRUCT}(struct))
	{
		if o.UserInfo != nil {
			if err := o.UserInfo.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&WorkstationUserEnum{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// TotalEntries {out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.WriteData(o.TotalEntries); err != nil {
			return err
		}
	}
	// ResumeHandle {in, out} (1:{pointer=unique}*(1)(uint32))
	{
		_ptr_ResumeHandle := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := w.WriteData(o.Resume); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.Resume, _ptr_ResumeHandle); err != nil {
			return err
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:(uint32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_UserEnumOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// UserInfo {in, out} (1:{alias=LPWKSTA_USER_ENUM_STRUCT,pointer=ref}*(1))(2:{alias=WKSTA_USER_ENUM_STRUCT}(struct))
	{
		if o.UserInfo == nil {
			o.UserInfo = &WorkstationUserEnum{}
		}
		if err := o.UserInfo.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// TotalEntries {out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.ReadData(&o.TotalEntries); err != nil {
			return err
		}
	}
	// ResumeHandle {in, out} (1:{pointer=unique}*(1)(uint32))
	{
		_ptr_ResumeHandle := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := w.ReadData(&o.Resume); err != nil {
				return err
			}
			return nil
		})
		_s_ResumeHandle := func(ptr interface{}) { o.Resume = *ptr.(*uint32) }
		if err := w.ReadPointer(&o.Resume, _s_ResumeHandle, _ptr_ResumeHandle); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// UserEnumRequest structure represents the NetrWkstaUserEnum operation request
type UserEnumRequest struct {
	// ServerName: A WKSSVC_IDENTIFY_HANDLE structure (section 2.2.2.1) that identifies
	// the server (2). The client MUST map this structure to an RPC binding handle ([C706]
	// sections 4.3.5 and 5.1.5.2). The server MUST ignore this parameter.
	ServerName string `idl:"name:ServerName;string;pointer:unique" json:"server_name"`
	// UserInfo: A pointer to the buffer to receive the data. The data MUST be returned
	// as a WKSTA_USER_ENUM_STRUCT structure (section 2.2.5.14) that contains a Level member
	// that specifies the type of structure to return.
	UserInfo *WorkstationUserEnum `idl:"name:UserInfo" json:"user_info"`
	// PreferredMaximumLength: The number of bytes to allocate for the return data.
	PreferredMaximumLength uint32 `idl:"name:PreferredMaximumLength" json:"preferred_maximum_length"`
	// ResumeHandle: A pointer that, if specified, and if this method returns ERROR_MORE_DATA,
	// MUST receive an implementation-specific value<34> that can be passed in subsequent
	// calls to this method, to continue with the enumeration of currently logged-on users.
	Resume uint32 `idl:"name:ResumeHandle;pointer:unique" json:"resume"`
}

func (o *UserEnumRequest) xxx_ToOp(ctx context.Context) *xxx_UserEnumOperation {
	if o == nil {
		return &xxx_UserEnumOperation{}
	}
	return &xxx_UserEnumOperation{
		ServerName:             o.ServerName,
		UserInfo:               o.UserInfo,
		PreferredMaximumLength: o.PreferredMaximumLength,
		Resume:                 o.Resume,
	}
}

func (o *UserEnumRequest) xxx_FromOp(ctx context.Context, op *xxx_UserEnumOperation) {
	if o == nil {
		return
	}
	o.ServerName = op.ServerName
	o.UserInfo = op.UserInfo
	o.PreferredMaximumLength = op.PreferredMaximumLength
	o.Resume = op.Resume
}
func (o *UserEnumRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *UserEnumRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_UserEnumOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// UserEnumResponse structure represents the NetrWkstaUserEnum operation response
type UserEnumResponse struct {
	// UserInfo: A pointer to the buffer to receive the data. The data MUST be returned
	// as a WKSTA_USER_ENUM_STRUCT structure (section 2.2.5.14) that contains a Level member
	// that specifies the type of structure to return.
	UserInfo *WorkstationUserEnum `idl:"name:UserInfo" json:"user_info"`
	// TotalEntries: The total number of entries that could have been enumerated if the
	// buffer were big enough to hold all the entries.
	TotalEntries uint32 `idl:"name:TotalEntries" json:"total_entries"`
	// ResumeHandle: A pointer that, if specified, and if this method returns ERROR_MORE_DATA,
	// MUST receive an implementation-specific value<34> that can be passed in subsequent
	// calls to this method, to continue with the enumeration of currently logged-on users.
	Resume uint32 `idl:"name:ResumeHandle;pointer:unique" json:"resume"`
	// Return: The NetrWkstaUserEnum return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *UserEnumResponse) xxx_ToOp(ctx context.Context) *xxx_UserEnumOperation {
	if o == nil {
		return &xxx_UserEnumOperation{}
	}
	return &xxx_UserEnumOperation{
		UserInfo:     o.UserInfo,
		TotalEntries: o.TotalEntries,
		Resume:       o.Resume,
		Return:       o.Return,
	}
}

func (o *UserEnumResponse) xxx_FromOp(ctx context.Context, op *xxx_UserEnumOperation) {
	if o == nil {
		return
	}
	o.UserInfo = op.UserInfo
	o.TotalEntries = op.TotalEntries
	o.Resume = op.Resume
	o.Return = op.Return
}
func (o *UserEnumResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *UserEnumResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_UserEnumOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_TransportEnumOperation structure represents the NetrWkstaTransportEnum operation
type xxx_TransportEnumOperation struct {
	ServerName             string                    `idl:"name:ServerName;string;pointer:unique" json:"server_name"`
	TransportInfo          *WorkstationTransportEnum `idl:"name:TransportInfo" json:"transport_info"`
	PreferredMaximumLength uint32                    `idl:"name:PreferredMaximumLength" json:"preferred_maximum_length"`
	TotalEntries           uint32                    `idl:"name:TotalEntries" json:"total_entries"`
	Resume                 uint32                    `idl:"name:ResumeHandle;pointer:unique" json:"resume"`
	Return                 uint32                    `idl:"name:Return" json:"return"`
}

func (o *xxx_TransportEnumOperation) OpNum() int { return 5 }

func (o *xxx_TransportEnumOperation) OpName() string { return "/wkssvc/v1/NetrWkstaTransportEnum" }

func (o *xxx_TransportEnumOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_TransportEnumOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// ServerName {in} (1:{handle, string, pointer=unique, alias=WKSSVC_IDENTIFY_HANDLE}*(1)[dim:0,string,null](wchar))
	{
		if o.ServerName != "" {
			_ptr_ServerName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.ServerName); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.ServerName, _ptr_ServerName); err != nil {
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
	// TransportInfo {in, out} (1:{alias=LPWKSTA_TRANSPORT_ENUM_STRUCT}*(1))(2:{alias=WKSTA_TRANSPORT_ENUM_STRUCT}(struct))
	{
		if o.TransportInfo != nil {
			if err := o.TransportInfo.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&WorkstationTransportEnum{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// PreferredMaximumLength {in} (1:(uint32))
	{
		if err := w.WriteData(o.PreferredMaximumLength); err != nil {
			return err
		}
	}
	// ResumeHandle {in, out} (1:{pointer=unique}*(1)(uint32))
	{
		_ptr_ResumeHandle := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := w.WriteData(o.Resume); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.Resume, _ptr_ResumeHandle); err != nil {
			return err
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_TransportEnumOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// ServerName {in} (1:{handle, string, pointer=unique, alias=WKSSVC_IDENTIFY_HANDLE}*(1)[dim:0,string,null](wchar))
	{
		_ptr_ServerName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.ServerName); err != nil {
				return err
			}
			return nil
		})
		_s_ServerName := func(ptr interface{}) { o.ServerName = *ptr.(*string) }
		if err := w.ReadPointer(&o.ServerName, _s_ServerName, _ptr_ServerName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// TransportInfo {in, out} (1:{alias=LPWKSTA_TRANSPORT_ENUM_STRUCT,pointer=ref}*(1))(2:{alias=WKSTA_TRANSPORT_ENUM_STRUCT}(struct))
	{
		if o.TransportInfo == nil {
			o.TransportInfo = &WorkstationTransportEnum{}
		}
		if err := o.TransportInfo.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// PreferredMaximumLength {in} (1:(uint32))
	{
		if err := w.ReadData(&o.PreferredMaximumLength); err != nil {
			return err
		}
	}
	// ResumeHandle {in, out} (1:{pointer=unique}*(1)(uint32))
	{
		_ptr_ResumeHandle := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := w.ReadData(&o.Resume); err != nil {
				return err
			}
			return nil
		})
		_s_ResumeHandle := func(ptr interface{}) { o.Resume = *ptr.(*uint32) }
		if err := w.ReadPointer(&o.Resume, _s_ResumeHandle, _ptr_ResumeHandle); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_TransportEnumOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_TransportEnumOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// TransportInfo {in, out} (1:{alias=LPWKSTA_TRANSPORT_ENUM_STRUCT}*(1))(2:{alias=WKSTA_TRANSPORT_ENUM_STRUCT}(struct))
	{
		if o.TransportInfo != nil {
			if err := o.TransportInfo.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&WorkstationTransportEnum{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// TotalEntries {out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.WriteData(o.TotalEntries); err != nil {
			return err
		}
	}
	// ResumeHandle {in, out} (1:{pointer=unique}*(1)(uint32))
	{
		_ptr_ResumeHandle := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := w.WriteData(o.Resume); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.Resume, _ptr_ResumeHandle); err != nil {
			return err
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:(uint32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_TransportEnumOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// TransportInfo {in, out} (1:{alias=LPWKSTA_TRANSPORT_ENUM_STRUCT,pointer=ref}*(1))(2:{alias=WKSTA_TRANSPORT_ENUM_STRUCT}(struct))
	{
		if o.TransportInfo == nil {
			o.TransportInfo = &WorkstationTransportEnum{}
		}
		if err := o.TransportInfo.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// TotalEntries {out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.ReadData(&o.TotalEntries); err != nil {
			return err
		}
	}
	// ResumeHandle {in, out} (1:{pointer=unique}*(1)(uint32))
	{
		_ptr_ResumeHandle := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := w.ReadData(&o.Resume); err != nil {
				return err
			}
			return nil
		})
		_s_ResumeHandle := func(ptr interface{}) { o.Resume = *ptr.(*uint32) }
		if err := w.ReadPointer(&o.Resume, _s_ResumeHandle, _ptr_ResumeHandle); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// TransportEnumRequest structure represents the NetrWkstaTransportEnum operation request
type TransportEnumRequest struct {
	// ServerName: A WKSSVC_IDENTIFY_HANDLE structure (section 2.2.2.1) that identifies
	// the server (2). The client MUST map this structure to an RPC binding handle ([C706]
	// sections 4.3.5 and 5.1.5.2). The server MUST ignore this parameter.
	ServerName string `idl:"name:ServerName;string;pointer:unique" json:"server_name"`
	// TransportInfo: A pointer to a buffer that receives a WKSTA_TRANSPORT_ENUM_STRUCT
	// structure (section 2.2.5.16), which contains a Level member that MUST be set to zero.
	TransportInfo *WorkstationTransportEnum `idl:"name:TransportInfo" json:"transport_info"`
	// PreferredMaximumLength: The number of bytes to allocate for the return data.
	PreferredMaximumLength uint32 `idl:"name:PreferredMaximumLength" json:"preferred_maximum_length"`
	// ResumeHandle: A pointer that, if specified, and if this method returns NERR_BufTooSmall,
	// MUST receive an implementation-specific value<37> that can be passed in subsequent
	// calls to this method, to continue with the enumeration of currently enabled transport
	// protocols.
	Resume uint32 `idl:"name:ResumeHandle;pointer:unique" json:"resume"`
}

func (o *TransportEnumRequest) xxx_ToOp(ctx context.Context) *xxx_TransportEnumOperation {
	if o == nil {
		return &xxx_TransportEnumOperation{}
	}
	return &xxx_TransportEnumOperation{
		ServerName:             o.ServerName,
		TransportInfo:          o.TransportInfo,
		PreferredMaximumLength: o.PreferredMaximumLength,
		Resume:                 o.Resume,
	}
}

func (o *TransportEnumRequest) xxx_FromOp(ctx context.Context, op *xxx_TransportEnumOperation) {
	if o == nil {
		return
	}
	o.ServerName = op.ServerName
	o.TransportInfo = op.TransportInfo
	o.PreferredMaximumLength = op.PreferredMaximumLength
	o.Resume = op.Resume
}
func (o *TransportEnumRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *TransportEnumRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_TransportEnumOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// TransportEnumResponse structure represents the NetrWkstaTransportEnum operation response
type TransportEnumResponse struct {
	// TransportInfo: A pointer to a buffer that receives a WKSTA_TRANSPORT_ENUM_STRUCT
	// structure (section 2.2.5.16), which contains a Level member that MUST be set to zero.
	TransportInfo *WorkstationTransportEnum `idl:"name:TransportInfo" json:"transport_info"`
	// TotalEntries: The total number of entries that could have been enumerated from the
	// current resume position. This field can be set to any value when sent and MUST be
	// ignored on receipt.
	TotalEntries uint32 `idl:"name:TotalEntries" json:"total_entries"`
	// ResumeHandle: A pointer that, if specified, and if this method returns NERR_BufTooSmall,
	// MUST receive an implementation-specific value<37> that can be passed in subsequent
	// calls to this method, to continue with the enumeration of currently enabled transport
	// protocols.
	Resume uint32 `idl:"name:ResumeHandle;pointer:unique" json:"resume"`
	// Return: The NetrWkstaTransportEnum return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *TransportEnumResponse) xxx_ToOp(ctx context.Context) *xxx_TransportEnumOperation {
	if o == nil {
		return &xxx_TransportEnumOperation{}
	}
	return &xxx_TransportEnumOperation{
		TransportInfo: o.TransportInfo,
		TotalEntries:  o.TotalEntries,
		Resume:        o.Resume,
		Return:        o.Return,
	}
}

func (o *TransportEnumResponse) xxx_FromOp(ctx context.Context, op *xxx_TransportEnumOperation) {
	if o == nil {
		return
	}
	o.TransportInfo = op.TransportInfo
	o.TotalEntries = op.TotalEntries
	o.Resume = op.Resume
	o.Return = op.Return
}
func (o *TransportEnumResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *TransportEnumResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_TransportEnumOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_TransportAddOperation structure represents the NetrWkstaTransportAdd operation
type xxx_TransportAddOperation struct {
	ServerName     string                     `idl:"name:ServerName;string;pointer:unique" json:"server_name"`
	Level          uint32                     `idl:"name:Level" json:"level"`
	TransportInfo  *WorkstationTransportInfo0 `idl:"name:TransportInfo" json:"transport_info"`
	ErrorParameter uint32                     `idl:"name:ErrorParameter;pointer:unique" json:"error_parameter"`
	Return         uint32                     `idl:"name:Return" json:"return"`
}

func (o *xxx_TransportAddOperation) OpNum() int { return 6 }

func (o *xxx_TransportAddOperation) OpName() string { return "/wkssvc/v1/NetrWkstaTransportAdd" }

func (o *xxx_TransportAddOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_TransportAddOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// ServerName {in} (1:{handle, string, pointer=unique, alias=WKSSVC_IDENTIFY_HANDLE}*(1)[dim:0,string,null](wchar))
	{
		if o.ServerName != "" {
			_ptr_ServerName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.ServerName); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.ServerName, _ptr_ServerName); err != nil {
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
	// Level {in} (1:(uint32))
	{
		if err := w.WriteData(o.Level); err != nil {
			return err
		}
	}
	// TransportInfo {in} (1:{alias=LPWKSTA_TRANSPORT_INFO_0}*(1))(2:{alias=WKSTA_TRANSPORT_INFO_0}(struct))
	{
		if o.TransportInfo != nil {
			if err := o.TransportInfo.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&WorkstationTransportInfo0{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// ErrorParameter {in, out} (1:{pointer=unique}*(1)(uint32))
	{
		_ptr_ErrorParameter := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := w.WriteData(o.ErrorParameter); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.ErrorParameter, _ptr_ErrorParameter); err != nil {
			return err
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_TransportAddOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// ServerName {in} (1:{handle, string, pointer=unique, alias=WKSSVC_IDENTIFY_HANDLE}*(1)[dim:0,string,null](wchar))
	{
		_ptr_ServerName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.ServerName); err != nil {
				return err
			}
			return nil
		})
		_s_ServerName := func(ptr interface{}) { o.ServerName = *ptr.(*string) }
		if err := w.ReadPointer(&o.ServerName, _s_ServerName, _ptr_ServerName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Level {in} (1:(uint32))
	{
		if err := w.ReadData(&o.Level); err != nil {
			return err
		}
	}
	// TransportInfo {in} (1:{alias=LPWKSTA_TRANSPORT_INFO_0,pointer=ref}*(1))(2:{alias=WKSTA_TRANSPORT_INFO_0}(struct))
	{
		if o.TransportInfo == nil {
			o.TransportInfo = &WorkstationTransportInfo0{}
		}
		if err := o.TransportInfo.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// ErrorParameter {in, out} (1:{pointer=unique}*(1)(uint32))
	{
		_ptr_ErrorParameter := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := w.ReadData(&o.ErrorParameter); err != nil {
				return err
			}
			return nil
		})
		_s_ErrorParameter := func(ptr interface{}) { o.ErrorParameter = *ptr.(*uint32) }
		if err := w.ReadPointer(&o.ErrorParameter, _s_ErrorParameter, _ptr_ErrorParameter); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_TransportAddOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_TransportAddOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// ErrorParameter {in, out} (1:{pointer=unique}*(1)(uint32))
	{
		_ptr_ErrorParameter := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := w.WriteData(o.ErrorParameter); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.ErrorParameter, _ptr_ErrorParameter); err != nil {
			return err
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:(uint32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_TransportAddOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// ErrorParameter {in, out} (1:{pointer=unique}*(1)(uint32))
	{
		_ptr_ErrorParameter := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := w.ReadData(&o.ErrorParameter); err != nil {
				return err
			}
			return nil
		})
		_s_ErrorParameter := func(ptr interface{}) { o.ErrorParameter = *ptr.(*uint32) }
		if err := w.ReadPointer(&o.ErrorParameter, _s_ErrorParameter, _ptr_ErrorParameter); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// TransportAddRequest structure represents the NetrWkstaTransportAdd operation request
type TransportAddRequest struct {
	// ServerName: A WKSSVC_IDENTIFY_HANDLE structure (section 2.2.2.1) that identifies
	// the server (2). The client MUST map this structure to an RPC binding handle ([C706]
	// sections 4.3.5 and 5.1.5.2). The server MUST ignore this parameter.
	ServerName string `idl:"name:ServerName;string;pointer:unique" json:"server_name"`
	// Level: The information level of the data. Level is set to zero, meaning the TransportInfo
	// parameter points to a WKSTA_TRANSPORT_INFO_0 structure (section 2.2.5.8).
	Level uint32 `idl:"name:Level" json:"level"`
	// TransportInfo: A pointer to a WKSTA_TRANSPORT_INFO_0 structure.
	TransportInfo *WorkstationTransportInfo0 `idl:"name:TransportInfo" json:"transport_info"`
	// ErrorParameter: A pointer to a value that receives the index, starting at 0, of the
	// first member of the WKSTA_TRANSPORT_INFO_0 structure that causes the function to
	// return ERROR_INVALID_PARAMETER. If this parameter is NULL, the index is not returned
	// on error.
	ErrorParameter uint32 `idl:"name:ErrorParameter;pointer:unique" json:"error_parameter"`
}

func (o *TransportAddRequest) xxx_ToOp(ctx context.Context) *xxx_TransportAddOperation {
	if o == nil {
		return &xxx_TransportAddOperation{}
	}
	return &xxx_TransportAddOperation{
		ServerName:     o.ServerName,
		Level:          o.Level,
		TransportInfo:  o.TransportInfo,
		ErrorParameter: o.ErrorParameter,
	}
}

func (o *TransportAddRequest) xxx_FromOp(ctx context.Context, op *xxx_TransportAddOperation) {
	if o == nil {
		return
	}
	o.ServerName = op.ServerName
	o.Level = op.Level
	o.TransportInfo = op.TransportInfo
	o.ErrorParameter = op.ErrorParameter
}
func (o *TransportAddRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *TransportAddRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_TransportAddOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// TransportAddResponse structure represents the NetrWkstaTransportAdd operation response
type TransportAddResponse struct {
	// ErrorParameter: A pointer to a value that receives the index, starting at 0, of the
	// first member of the WKSTA_TRANSPORT_INFO_0 structure that causes the function to
	// return ERROR_INVALID_PARAMETER. If this parameter is NULL, the index is not returned
	// on error.
	ErrorParameter uint32 `idl:"name:ErrorParameter;pointer:unique" json:"error_parameter"`
	// Return: The NetrWkstaTransportAdd return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *TransportAddResponse) xxx_ToOp(ctx context.Context) *xxx_TransportAddOperation {
	if o == nil {
		return &xxx_TransportAddOperation{}
	}
	return &xxx_TransportAddOperation{
		ErrorParameter: o.ErrorParameter,
		Return:         o.Return,
	}
}

func (o *TransportAddResponse) xxx_FromOp(ctx context.Context, op *xxx_TransportAddOperation) {
	if o == nil {
		return
	}
	o.ErrorParameter = op.ErrorParameter
	o.Return = op.Return
}
func (o *TransportAddResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *TransportAddResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_TransportAddOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_TransportDeleteOperation structure represents the NetrWkstaTransportDel operation
type xxx_TransportDeleteOperation struct {
	ServerName    string `idl:"name:ServerName;string;pointer:unique" json:"server_name"`
	TransportName string `idl:"name:TransportName;string;pointer:unique" json:"transport_name"`
	ForceLevel    uint32 `idl:"name:ForceLevel" json:"force_level"`
	Return        uint32 `idl:"name:Return" json:"return"`
}

func (o *xxx_TransportDeleteOperation) OpNum() int { return 7 }

func (o *xxx_TransportDeleteOperation) OpName() string { return "/wkssvc/v1/NetrWkstaTransportDel" }

func (o *xxx_TransportDeleteOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_TransportDeleteOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// ServerName {in} (1:{handle, string, pointer=unique, alias=WKSSVC_IDENTIFY_HANDLE}*(1)[dim:0,string,null](wchar))
	{
		if o.ServerName != "" {
			_ptr_ServerName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.ServerName); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.ServerName, _ptr_ServerName); err != nil {
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
	// TransportName {in} (1:{string, pointer=unique}*(1)[dim:0,string,null](wchar))
	{
		if o.TransportName != "" {
			_ptr_TransportName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.TransportName); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.TransportName, _ptr_TransportName); err != nil {
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
	// ForceLevel {in} (1:(uint32))
	{
		if err := w.WriteData(o.ForceLevel); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_TransportDeleteOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// ServerName {in} (1:{handle, string, pointer=unique, alias=WKSSVC_IDENTIFY_HANDLE}*(1)[dim:0,string,null](wchar))
	{
		_ptr_ServerName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.ServerName); err != nil {
				return err
			}
			return nil
		})
		_s_ServerName := func(ptr interface{}) { o.ServerName = *ptr.(*string) }
		if err := w.ReadPointer(&o.ServerName, _s_ServerName, _ptr_ServerName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// TransportName {in} (1:{string, pointer=unique}*(1)[dim:0,string,null](wchar))
	{
		_ptr_TransportName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.TransportName); err != nil {
				return err
			}
			return nil
		})
		_s_TransportName := func(ptr interface{}) { o.TransportName = *ptr.(*string) }
		if err := w.ReadPointer(&o.TransportName, _s_TransportName, _ptr_TransportName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// ForceLevel {in} (1:(uint32))
	{
		if err := w.ReadData(&o.ForceLevel); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_TransportDeleteOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_TransportDeleteOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:(uint32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_TransportDeleteOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// TransportDeleteRequest structure represents the NetrWkstaTransportDel operation request
type TransportDeleteRequest struct {
	// ServerName: A WKSSVC_IDENTIFY_HANDLE structure (section 2.2.2.1) that identifies
	// the server. The client MUST map this structure to an RPC binding handle ([C706] sections
	// 4.3.5 and 5.1.5.2). The server MUST ignore this parameter.
	ServerName string `idl:"name:ServerName;string;pointer:unique" json:"server_name"`
	// TransportName: A pointer to a string that specifies the name of the transport protocol
	// to disconnect from the SMB network redirector.
	TransportName string `idl:"name:TransportName;string;pointer:unique" json:"transport_name"`
	// ForceLevel: The action to take if there are handles open to files or printers using
	// the transport protocol. This parameter MUST be one of the following values:
	//
	//	+------------------------------+----------------------------------------------------------------------------------+
	//	|                              |                                                                                  |
	//	|          VALUE/CODE          |                                     MEANING                                      |
	//	|                              |                                                                                  |
	//	+------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------+----------------------------------------------------------------------------------+
	//	| USE_NOFORCE 0x00000000       | Do not disconnect or close the open handles if open handles are using the        |
	//	|                              | transport protocol.                                                              |
	//	+------------------------------+----------------------------------------------------------------------------------+
	//	| USE_FORCE 0x00000001         | Same as 0x00000000 (USE_NOFORCE); do not disconnect or close the open handles if |
	//	|                              | open handles are using the transport protocol.                                   |
	//	+------------------------------+----------------------------------------------------------------------------------+
	//	| USE_LOTS_OF_FORCE 0x00000002 | Forcefully close any open handles and disable the specified transport protocol   |
	//	|                              | from the SMB network redirector.                                                 |
	//	+------------------------------+----------------------------------------------------------------------------------+
	ForceLevel uint32 `idl:"name:ForceLevel" json:"force_level"`
}

func (o *TransportDeleteRequest) xxx_ToOp(ctx context.Context) *xxx_TransportDeleteOperation {
	if o == nil {
		return &xxx_TransportDeleteOperation{}
	}
	return &xxx_TransportDeleteOperation{
		ServerName:    o.ServerName,
		TransportName: o.TransportName,
		ForceLevel:    o.ForceLevel,
	}
}

func (o *TransportDeleteRequest) xxx_FromOp(ctx context.Context, op *xxx_TransportDeleteOperation) {
	if o == nil {
		return
	}
	o.ServerName = op.ServerName
	o.TransportName = op.TransportName
	o.ForceLevel = op.ForceLevel
}
func (o *TransportDeleteRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *TransportDeleteRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_TransportDeleteOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// TransportDeleteResponse structure represents the NetrWkstaTransportDel operation response
type TransportDeleteResponse struct {
	// Return: The NetrWkstaTransportDel return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *TransportDeleteResponse) xxx_ToOp(ctx context.Context) *xxx_TransportDeleteOperation {
	if o == nil {
		return &xxx_TransportDeleteOperation{}
	}
	return &xxx_TransportDeleteOperation{
		Return: o.Return,
	}
}

func (o *TransportDeleteResponse) xxx_FromOp(ctx context.Context, op *xxx_TransportDeleteOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *TransportDeleteResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *TransportDeleteResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_TransportDeleteOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_UseAddOperation structure represents the NetrUseAdd operation
type xxx_UseAddOperation struct {
	ServerName     string   `idl:"name:ServerName;string;pointer:unique" json:"server_name"`
	Level          uint32   `idl:"name:Level" json:"level"`
	Info           *UseInfo `idl:"name:InfoStruct;switch_is:Level" json:"info"`
	ErrorParameter uint32   `idl:"name:ErrorParameter;pointer:unique" json:"error_parameter"`
	Return         uint32   `idl:"name:Return" json:"return"`
}

func (o *xxx_UseAddOperation) OpNum() int { return 8 }

func (o *xxx_UseAddOperation) OpName() string { return "/wkssvc/v1/NetrUseAdd" }

func (o *xxx_UseAddOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_UseAddOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// ServerName {in} (1:{handle, string, pointer=unique, alias=WKSSVC_IMPERSONATE_HANDLE}*(1)[dim:0,string,null](wchar))
	{
		if o.ServerName != "" {
			_ptr_ServerName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.ServerName); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.ServerName, _ptr_ServerName); err != nil {
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
	// Level {in} (1:(uint32))
	{
		if err := w.WriteData(o.Level); err != nil {
			return err
		}
	}
	// InfoStruct {in} (1:{switch_type={}(uint32), alias=LPUSE_INFO}*(1))(2:{switch_type={}(uint32), alias=USE_INFO}(union))
	{
		_swInfo := uint32(o.Level)
		if o.Info != nil {
			if err := o.Info.MarshalUnionNDR(ctx, w, _swInfo); err != nil {
				return err
			}
		} else {
			if err := (&UseInfo{}).MarshalUnionNDR(ctx, w, _swInfo); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// ErrorParameter {in, out} (1:{pointer=unique}*(1)(uint32))
	{
		_ptr_ErrorParameter := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := w.WriteData(o.ErrorParameter); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.ErrorParameter, _ptr_ErrorParameter); err != nil {
			return err
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_UseAddOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// ServerName {in} (1:{handle, string, pointer=unique, alias=WKSSVC_IMPERSONATE_HANDLE}*(1)[dim:0,string,null](wchar))
	{
		_ptr_ServerName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.ServerName); err != nil {
				return err
			}
			return nil
		})
		_s_ServerName := func(ptr interface{}) { o.ServerName = *ptr.(*string) }
		if err := w.ReadPointer(&o.ServerName, _s_ServerName, _ptr_ServerName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Level {in} (1:(uint32))
	{
		if err := w.ReadData(&o.Level); err != nil {
			return err
		}
	}
	// InfoStruct {in} (1:{switch_type={}(uint32), alias=LPUSE_INFO,pointer=ref}*(1))(2:{switch_type={}(uint32), alias=USE_INFO}(union))
	{
		if o.Info == nil {
			o.Info = &UseInfo{}
		}
		_swInfo := uint32(o.Level)
		if err := o.Info.UnmarshalUnionNDR(ctx, w, _swInfo); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// ErrorParameter {in, out} (1:{pointer=unique}*(1)(uint32))
	{
		_ptr_ErrorParameter := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := w.ReadData(&o.ErrorParameter); err != nil {
				return err
			}
			return nil
		})
		_s_ErrorParameter := func(ptr interface{}) { o.ErrorParameter = *ptr.(*uint32) }
		if err := w.ReadPointer(&o.ErrorParameter, _s_ErrorParameter, _ptr_ErrorParameter); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_UseAddOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_UseAddOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// ErrorParameter {in, out} (1:{pointer=unique}*(1)(uint32))
	{
		_ptr_ErrorParameter := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := w.WriteData(o.ErrorParameter); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.ErrorParameter, _ptr_ErrorParameter); err != nil {
			return err
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:(uint32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_UseAddOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// ErrorParameter {in, out} (1:{pointer=unique}*(1)(uint32))
	{
		_ptr_ErrorParameter := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := w.ReadData(&o.ErrorParameter); err != nil {
				return err
			}
			return nil
		})
		_s_ErrorParameter := func(ptr interface{}) { o.ErrorParameter = *ptr.(*uint32) }
		if err := w.ReadPointer(&o.ErrorParameter, _s_ErrorParameter, _ptr_ErrorParameter); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// UseAddRequest structure represents the NetrUseAdd operation request
type UseAddRequest struct {
	// ServerName: A WKSSVC_IMPERSONATE_HANDLE structure (section 2.2.2.2) that identifies
	// the server (2). The client MUST map this structure to an RPC binding handle ([C706]
	// sections 4.3.5 and 5.1.5.2). The server MUST ignore this parameter.
	ServerName string `idl:"name:ServerName;string;pointer:unique" json:"server_name"`
	// Level: A value that specifies the information level of the data. This parameter MUST
	// be one of the following values; otherwise, the server MUST fail the call with an
	// ERROR_INVALID_LEVEL code.
	//
	//	+------------+----------------------------------------------------------+
	//	|            |                                                          |
	//	|   VALUE    |                         MEANING                          |
	//	|            |                                                          |
	//	+------------+----------------------------------------------------------+
	//	+------------+----------------------------------------------------------+
	//	| 0x00000000 | The buffer is a USE_INFO_0 structure (section 2.2.5.22). |
	//	+------------+----------------------------------------------------------+
	//	| 0x00000001 | The buffer is a USE_INFO_1 structure (section 2.2.5.23). |
	//	+------------+----------------------------------------------------------+
	//	| 0x00000002 | The buffer is a USE_INFO_2 structure (section 2.2.5.24). |
	//	+------------+----------------------------------------------------------+
	//	| 0x00000003 | The buffer is a USE_INFO_3 structure (section 2.2.5.25). |
	//	+------------+----------------------------------------------------------+
	Level uint32 `idl:"name:Level" json:"level"`
	// InfoStruct: A pointer to the buffer that specifies the data. The format of this data
	// depends on the value of the Level parameter.
	Info *UseInfo `idl:"name:InfoStruct;switch_is:Level" json:"info"`
	// ErrorParameter: A pointer to a value that receives an unsigned 32-bit integer. This
	// parameter is meaningful only if the method returns ERROR_INVALID_PARAMETER.
	ErrorParameter uint32 `idl:"name:ErrorParameter;pointer:unique" json:"error_parameter"`
}

func (o *UseAddRequest) xxx_ToOp(ctx context.Context) *xxx_UseAddOperation {
	if o == nil {
		return &xxx_UseAddOperation{}
	}
	return &xxx_UseAddOperation{
		ServerName:     o.ServerName,
		Level:          o.Level,
		Info:           o.Info,
		ErrorParameter: o.ErrorParameter,
	}
}

func (o *UseAddRequest) xxx_FromOp(ctx context.Context, op *xxx_UseAddOperation) {
	if o == nil {
		return
	}
	o.ServerName = op.ServerName
	o.Level = op.Level
	o.Info = op.Info
	o.ErrorParameter = op.ErrorParameter
}
func (o *UseAddRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *UseAddRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_UseAddOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// UseAddResponse structure represents the NetrUseAdd operation response
type UseAddResponse struct {
	// ErrorParameter: A pointer to a value that receives an unsigned 32-bit integer. This
	// parameter is meaningful only if the method returns ERROR_INVALID_PARAMETER.
	ErrorParameter uint32 `idl:"name:ErrorParameter;pointer:unique" json:"error_parameter"`
	// Return: The NetrUseAdd return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *UseAddResponse) xxx_ToOp(ctx context.Context) *xxx_UseAddOperation {
	if o == nil {
		return &xxx_UseAddOperation{}
	}
	return &xxx_UseAddOperation{
		ErrorParameter: o.ErrorParameter,
		Return:         o.Return,
	}
}

func (o *UseAddResponse) xxx_FromOp(ctx context.Context, op *xxx_UseAddOperation) {
	if o == nil {
		return
	}
	o.ErrorParameter = op.ErrorParameter
	o.Return = op.Return
}
func (o *UseAddResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *UseAddResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_UseAddOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_UseGetInfoOperation structure represents the NetrUseGetInfo operation
type xxx_UseGetInfoOperation struct {
	ServerName string   `idl:"name:ServerName;string;pointer:unique" json:"server_name"`
	UseName    string   `idl:"name:UseName;string" json:"use_name"`
	Level      uint32   `idl:"name:Level" json:"level"`
	Info       *UseInfo `idl:"name:InfoStruct;switch_is:Level" json:"info"`
	Return     uint32   `idl:"name:Return" json:"return"`
}

func (o *xxx_UseGetInfoOperation) OpNum() int { return 9 }

func (o *xxx_UseGetInfoOperation) OpName() string { return "/wkssvc/v1/NetrUseGetInfo" }

func (o *xxx_UseGetInfoOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_UseGetInfoOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// ServerName {in} (1:{handle, string, pointer=unique, alias=WKSSVC_IMPERSONATE_HANDLE}*(1)[dim:0,string,null](wchar))
	{
		if o.ServerName != "" {
			_ptr_ServerName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.ServerName); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.ServerName, _ptr_ServerName); err != nil {
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
	// UseName {in} (1:{string, pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.UseName); err != nil {
			return err
		}
	}
	// Level {in} (1:(uint32))
	{
		if err := w.WriteData(o.Level); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_UseGetInfoOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// ServerName {in} (1:{handle, string, pointer=unique, alias=WKSSVC_IMPERSONATE_HANDLE}*(1)[dim:0,string,null](wchar))
	{
		_ptr_ServerName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.ServerName); err != nil {
				return err
			}
			return nil
		})
		_s_ServerName := func(ptr interface{}) { o.ServerName = *ptr.(*string) }
		if err := w.ReadPointer(&o.ServerName, _s_ServerName, _ptr_ServerName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// UseName {in} (1:{string, pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.UseName); err != nil {
			return err
		}
	}
	// Level {in} (1:(uint32))
	{
		if err := w.ReadData(&o.Level); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_UseGetInfoOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_UseGetInfoOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// InfoStruct {out} (1:{switch_type={}(uint32), alias=LPUSE_INFO}*(1))(2:{switch_type={}(uint32), alias=USE_INFO}(union))
	{
		_swInfo := uint32(o.Level)
		if o.Info != nil {
			if err := o.Info.MarshalUnionNDR(ctx, w, _swInfo); err != nil {
				return err
			}
		} else {
			if err := (&UseInfo{}).MarshalUnionNDR(ctx, w, _swInfo); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:(uint32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_UseGetInfoOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// InfoStruct {out} (1:{switch_type={}(uint32), alias=LPUSE_INFO,pointer=ref}*(1))(2:{switch_type={}(uint32), alias=USE_INFO}(union))
	{
		if o.Info == nil {
			o.Info = &UseInfo{}
		}
		_swInfo := uint32(o.Level)
		if err := o.Info.UnmarshalUnionNDR(ctx, w, _swInfo); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// UseGetInfoRequest structure represents the NetrUseGetInfo operation request
type UseGetInfoRequest struct {
	// ServerName: A WKSSVC_IMPERSONATE_HANDLE structure (section 2.2.2.2) that identifies
	// the server (2). The client MUST map this structure to an RPC binding handle ([C706]
	// sections 4.3.5 and 5.1.5.2). The server MUST ignore this parameter.
	ServerName string `idl:"name:ServerName;string;pointer:unique" json:"server_name"`
	// UseName: A pointer to a string that specifies the local device name or shared resource
	// name for which to return information.
	UseName string `idl:"name:UseName;string" json:"use_name"`
	// Level: A value that specifies the information level of the data. This parameter MUST
	// be one of the following values; otherwise, the server MUST fail the call with an
	// ERROR_INVALID_LEVEL code.
	//
	//	+------------+----------------------------------------------------------+
	//	|            |                                                          |
	//	|   VALUE    |                         MEANING                          |
	//	|            |                                                          |
	//	+------------+----------------------------------------------------------+
	//	+------------+----------------------------------------------------------+
	//	| 0x00000000 | The buffer is a USE_INFO_0 structure (section 2.2.5.22). |
	//	+------------+----------------------------------------------------------+
	//	| 0x00000001 | The buffer is a USE_INFO_1 structure (section 2.2.5.23). |
	//	+------------+----------------------------------------------------------+
	//	| 0x00000002 | The buffer is a USE_INFO_2 structure (section 2.2.5.24). |
	//	+------------+----------------------------------------------------------+
	//	| 0x00000003 | The buffer is a USE_INFO_3 structure (section 2.2.5.25). |
	//	+------------+----------------------------------------------------------+
	Level uint32 `idl:"name:Level" json:"level"`
}

func (o *UseGetInfoRequest) xxx_ToOp(ctx context.Context) *xxx_UseGetInfoOperation {
	if o == nil {
		return &xxx_UseGetInfoOperation{}
	}
	return &xxx_UseGetInfoOperation{
		ServerName: o.ServerName,
		UseName:    o.UseName,
		Level:      o.Level,
	}
}

func (o *UseGetInfoRequest) xxx_FromOp(ctx context.Context, op *xxx_UseGetInfoOperation) {
	if o == nil {
		return
	}
	o.ServerName = op.ServerName
	o.UseName = op.UseName
	o.Level = op.Level
}
func (o *UseGetInfoRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *UseGetInfoRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_UseGetInfoOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// UseGetInfoResponse structure represents the NetrUseGetInfo operation response
type UseGetInfoResponse struct {
	// InfoStruct: A pointer to the buffer that specifies the data. The format of this data
	// depends on the value of the Level parameter.
	Info *UseInfo `idl:"name:InfoStruct;switch_is:Level" json:"info"`
	// Return: The NetrUseGetInfo return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *UseGetInfoResponse) xxx_ToOp(ctx context.Context) *xxx_UseGetInfoOperation {
	if o == nil {
		return &xxx_UseGetInfoOperation{}
	}
	return &xxx_UseGetInfoOperation{
		Info:   o.Info,
		Return: o.Return,
	}
}

func (o *UseGetInfoResponse) xxx_FromOp(ctx context.Context, op *xxx_UseGetInfoOperation) {
	if o == nil {
		return
	}
	o.Info = op.Info
	o.Return = op.Return
}
func (o *UseGetInfoResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *UseGetInfoResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_UseGetInfoOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_UseDeleteOperation structure represents the NetrUseDel operation
type xxx_UseDeleteOperation struct {
	ServerName string `idl:"name:ServerName;string;pointer:unique" json:"server_name"`
	UseName    string `idl:"name:UseName;string" json:"use_name"`
	ForceLevel uint32 `idl:"name:ForceLevel" json:"force_level"`
	Return     uint32 `idl:"name:Return" json:"return"`
}

func (o *xxx_UseDeleteOperation) OpNum() int { return 10 }

func (o *xxx_UseDeleteOperation) OpName() string { return "/wkssvc/v1/NetrUseDel" }

func (o *xxx_UseDeleteOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_UseDeleteOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// ServerName {in} (1:{handle, string, pointer=unique, alias=WKSSVC_IMPERSONATE_HANDLE}*(1)[dim:0,string,null](wchar))
	{
		if o.ServerName != "" {
			_ptr_ServerName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.ServerName); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.ServerName, _ptr_ServerName); err != nil {
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
	// UseName {in} (1:{string, pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.UseName); err != nil {
			return err
		}
	}
	// ForceLevel {in} (1:(uint32))
	{
		if err := w.WriteData(o.ForceLevel); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_UseDeleteOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// ServerName {in} (1:{handle, string, pointer=unique, alias=WKSSVC_IMPERSONATE_HANDLE}*(1)[dim:0,string,null](wchar))
	{
		_ptr_ServerName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.ServerName); err != nil {
				return err
			}
			return nil
		})
		_s_ServerName := func(ptr interface{}) { o.ServerName = *ptr.(*string) }
		if err := w.ReadPointer(&o.ServerName, _s_ServerName, _ptr_ServerName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// UseName {in} (1:{string, pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.UseName); err != nil {
			return err
		}
	}
	// ForceLevel {in} (1:(uint32))
	{
		if err := w.ReadData(&o.ForceLevel); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_UseDeleteOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_UseDeleteOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:(uint32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_UseDeleteOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// UseDeleteRequest structure represents the NetrUseDel operation request
type UseDeleteRequest struct {
	// ServerName: A WKSSVC_IMPERSONATE_HANDLE structure (section 2.2.2.2) that identifies
	// the server (2). The client MUST map this structure to an RPC binding handle ([C706]
	// sections 4.3.5 and 5.1.5.2). The server MUST ignore this parameter.
	ServerName string `idl:"name:ServerName;string;pointer:unique" json:"server_name"`
	// UseName: A pointer to a string that specifies the local device name or shared resource
	// name for which to return information.
	UseName string `idl:"name:UseName;string" json:"use_name"`
	// ForceLevel: The level of force to use in deleting the connection. This parameter
	// MUST be one of the following values; otherwise, the server MUST fail the call with
	// an ERROR_INVALID_LEVEL error code.
	//
	//	+------------------------------+----------------------------------------------------------------------------------+
	//	|                              |                                                                                  |
	//	|          VALUE/CODE          |                                     MEANING                                      |
	//	|                              |                                                                                  |
	//	+------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------+----------------------------------------------------------------------------------+
	//	| USE_NOFORCE 0x00000000       | Do not disconnect the connection if open files exist on the connection.          |
	//	+------------------------------+----------------------------------------------------------------------------------+
	//	| USE_FORCE 0x00000001         | Same as 0x00000000 (USE_NOFORCE); do not disconnect the connection if open files |
	//	|                              | exist on the connection.                                                         |
	//	+------------------------------+----------------------------------------------------------------------------------+
	//	| USE_LOTS_OF_FORCE 0x00000002 | Close any open files and disconnect the connection.                              |
	//	+------------------------------+----------------------------------------------------------------------------------+
	ForceLevel uint32 `idl:"name:ForceLevel" json:"force_level"`
}

func (o *UseDeleteRequest) xxx_ToOp(ctx context.Context) *xxx_UseDeleteOperation {
	if o == nil {
		return &xxx_UseDeleteOperation{}
	}
	return &xxx_UseDeleteOperation{
		ServerName: o.ServerName,
		UseName:    o.UseName,
		ForceLevel: o.ForceLevel,
	}
}

func (o *UseDeleteRequest) xxx_FromOp(ctx context.Context, op *xxx_UseDeleteOperation) {
	if o == nil {
		return
	}
	o.ServerName = op.ServerName
	o.UseName = op.UseName
	o.ForceLevel = op.ForceLevel
}
func (o *UseDeleteRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *UseDeleteRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_UseDeleteOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// UseDeleteResponse structure represents the NetrUseDel operation response
type UseDeleteResponse struct {
	// Return: The NetrUseDel return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *UseDeleteResponse) xxx_ToOp(ctx context.Context) *xxx_UseDeleteOperation {
	if o == nil {
		return &xxx_UseDeleteOperation{}
	}
	return &xxx_UseDeleteOperation{
		Return: o.Return,
	}
}

func (o *UseDeleteResponse) xxx_FromOp(ctx context.Context, op *xxx_UseDeleteOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *UseDeleteResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *UseDeleteResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_UseDeleteOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_UseEnumOperation structure represents the NetrUseEnum operation
type xxx_UseEnumOperation struct {
	ServerName             string   `idl:"name:ServerName;string;pointer:unique" json:"server_name"`
	Info                   *UseEnum `idl:"name:InfoStruct" json:"info"`
	PreferredMaximumLength uint32   `idl:"name:PreferredMaximumLength" json:"preferred_maximum_length"`
	TotalEntries           uint32   `idl:"name:TotalEntries" json:"total_entries"`
	Resume                 uint32   `idl:"name:ResumeHandle;pointer:unique" json:"resume"`
	Return                 uint32   `idl:"name:Return" json:"return"`
}

func (o *xxx_UseEnumOperation) OpNum() int { return 11 }

func (o *xxx_UseEnumOperation) OpName() string { return "/wkssvc/v1/NetrUseEnum" }

func (o *xxx_UseEnumOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_UseEnumOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// ServerName {in} (1:{handle, string, pointer=unique, alias=WKSSVC_IDENTIFY_HANDLE}*(1)[dim:0,string,null](wchar))
	{
		if o.ServerName != "" {
			_ptr_ServerName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.ServerName); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.ServerName, _ptr_ServerName); err != nil {
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
	// InfoStruct {in, out} (1:{alias=LPUSE_ENUM_STRUCT}*(1))(2:{alias=USE_ENUM_STRUCT}(struct))
	{
		if o.Info != nil {
			if err := o.Info.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&UseEnum{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// PreferredMaximumLength {in} (1:(uint32))
	{
		if err := w.WriteData(o.PreferredMaximumLength); err != nil {
			return err
		}
	}
	// ResumeHandle {in, out} (1:{pointer=unique}*(1)(uint32))
	{
		_ptr_ResumeHandle := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := w.WriteData(o.Resume); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.Resume, _ptr_ResumeHandle); err != nil {
			return err
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_UseEnumOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// ServerName {in} (1:{handle, string, pointer=unique, alias=WKSSVC_IDENTIFY_HANDLE}*(1)[dim:0,string,null](wchar))
	{
		_ptr_ServerName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.ServerName); err != nil {
				return err
			}
			return nil
		})
		_s_ServerName := func(ptr interface{}) { o.ServerName = *ptr.(*string) }
		if err := w.ReadPointer(&o.ServerName, _s_ServerName, _ptr_ServerName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// InfoStruct {in, out} (1:{alias=LPUSE_ENUM_STRUCT,pointer=ref}*(1))(2:{alias=USE_ENUM_STRUCT}(struct))
	{
		if o.Info == nil {
			o.Info = &UseEnum{}
		}
		if err := o.Info.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// PreferredMaximumLength {in} (1:(uint32))
	{
		if err := w.ReadData(&o.PreferredMaximumLength); err != nil {
			return err
		}
	}
	// ResumeHandle {in, out} (1:{pointer=unique}*(1)(uint32))
	{
		_ptr_ResumeHandle := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := w.ReadData(&o.Resume); err != nil {
				return err
			}
			return nil
		})
		_s_ResumeHandle := func(ptr interface{}) { o.Resume = *ptr.(*uint32) }
		if err := w.ReadPointer(&o.Resume, _s_ResumeHandle, _ptr_ResumeHandle); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_UseEnumOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_UseEnumOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// InfoStruct {in, out} (1:{alias=LPUSE_ENUM_STRUCT}*(1))(2:{alias=USE_ENUM_STRUCT}(struct))
	{
		if o.Info != nil {
			if err := o.Info.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&UseEnum{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// TotalEntries {out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.WriteData(o.TotalEntries); err != nil {
			return err
		}
	}
	// ResumeHandle {in, out} (1:{pointer=unique}*(1)(uint32))
	{
		_ptr_ResumeHandle := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := w.WriteData(o.Resume); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.Resume, _ptr_ResumeHandle); err != nil {
			return err
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:(uint32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_UseEnumOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// InfoStruct {in, out} (1:{alias=LPUSE_ENUM_STRUCT,pointer=ref}*(1))(2:{alias=USE_ENUM_STRUCT}(struct))
	{
		if o.Info == nil {
			o.Info = &UseEnum{}
		}
		if err := o.Info.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// TotalEntries {out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.ReadData(&o.TotalEntries); err != nil {
			return err
		}
	}
	// ResumeHandle {in, out} (1:{pointer=unique}*(1)(uint32))
	{
		_ptr_ResumeHandle := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := w.ReadData(&o.Resume); err != nil {
				return err
			}
			return nil
		})
		_s_ResumeHandle := func(ptr interface{}) { o.Resume = *ptr.(*uint32) }
		if err := w.ReadPointer(&o.Resume, _s_ResumeHandle, _ptr_ResumeHandle); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// UseEnumRequest structure represents the NetrUseEnum operation request
type UseEnumRequest struct {
	// ServerName: A WKSSVC_IDENTIFY_HANDLE structure (section 2.2.2.1) that identifies
	// the server (2). The client MUST map this structure to an RPC binding handle ([C706]
	// sections 4.3.5 and 5.1.5.2). The server MUST ignore this parameter.
	ServerName string `idl:"name:ServerName;string;pointer:unique" json:"server_name"`
	// InfoStruct: The USE_ENUM_STRUCT structure (section 2.2.5.29) contains a Level parameter
	// that indicates the type of structure to return.
	//
	//	+------------+----------------------------------------------------------------------------------+
	//	|            |                                                                                  |
	//	|   VALUE    |                                     MEANING                                      |
	//	|            |                                                                                  |
	//	+------------+----------------------------------------------------------------------------------+
	//	+------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 | Specifies a local device name and the share name of a remote resource.           |
	//	+------------+----------------------------------------------------------------------------------+
	//	| 0x00000001 | Specifies details about the connection between a local device and a shared       |
	//	|            | resource, including connection status and connection type.                       |
	//	+------------+----------------------------------------------------------------------------------+
	//	| 0x00000002 | Specifies details about the connection between a local device and a shared       |
	//	|            | resource, including connection status, connection type, user name, and domain    |
	//	|            | name.                                                                            |
	//	+------------+----------------------------------------------------------------------------------+
	Info *UseEnum `idl:"name:InfoStruct" json:"info"`
	// PreferredMaximumLength: The number of bytes to allocate for the return data.
	PreferredMaximumLength uint32 `idl:"name:PreferredMaximumLength" json:"preferred_maximum_length"`
	// ResumeHandle: A pointer that, if specified and if this method returns ERROR_MORE_DATA,
	// MUST receive an implementation-specific value that can be passed in subsequent calls
	// to this method in order to continue with the enumeration of currently logged-on users.
	Resume uint32 `idl:"name:ResumeHandle;pointer:unique" json:"resume"`
}

func (o *UseEnumRequest) xxx_ToOp(ctx context.Context) *xxx_UseEnumOperation {
	if o == nil {
		return &xxx_UseEnumOperation{}
	}
	return &xxx_UseEnumOperation{
		ServerName:             o.ServerName,
		Info:                   o.Info,
		PreferredMaximumLength: o.PreferredMaximumLength,
		Resume:                 o.Resume,
	}
}

func (o *UseEnumRequest) xxx_FromOp(ctx context.Context, op *xxx_UseEnumOperation) {
	if o == nil {
		return
	}
	o.ServerName = op.ServerName
	o.Info = op.Info
	o.PreferredMaximumLength = op.PreferredMaximumLength
	o.Resume = op.Resume
}
func (o *UseEnumRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *UseEnumRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_UseEnumOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// UseEnumResponse structure represents the NetrUseEnum operation response
type UseEnumResponse struct {
	// InfoStruct: The USE_ENUM_STRUCT structure (section 2.2.5.29) contains a Level parameter
	// that indicates the type of structure to return.
	//
	//	+------------+----------------------------------------------------------------------------------+
	//	|            |                                                                                  |
	//	|   VALUE    |                                     MEANING                                      |
	//	|            |                                                                                  |
	//	+------------+----------------------------------------------------------------------------------+
	//	+------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 | Specifies a local device name and the share name of a remote resource.           |
	//	+------------+----------------------------------------------------------------------------------+
	//	| 0x00000001 | Specifies details about the connection between a local device and a shared       |
	//	|            | resource, including connection status and connection type.                       |
	//	+------------+----------------------------------------------------------------------------------+
	//	| 0x00000002 | Specifies details about the connection between a local device and a shared       |
	//	|            | resource, including connection status, connection type, user name, and domain    |
	//	|            | name.                                                                            |
	//	+------------+----------------------------------------------------------------------------------+
	Info *UseEnum `idl:"name:InfoStruct" json:"info"`
	// TotalEntries: The total number of entries that could have been enumerated if the
	// buffer were big enough to hold all the entries.
	TotalEntries uint32 `idl:"name:TotalEntries" json:"total_entries"`
	// ResumeHandle: A pointer that, if specified and if this method returns ERROR_MORE_DATA,
	// MUST receive an implementation-specific value that can be passed in subsequent calls
	// to this method in order to continue with the enumeration of currently logged-on users.
	Resume uint32 `idl:"name:ResumeHandle;pointer:unique" json:"resume"`
	// Return: The NetrUseEnum return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *UseEnumResponse) xxx_ToOp(ctx context.Context) *xxx_UseEnumOperation {
	if o == nil {
		return &xxx_UseEnumOperation{}
	}
	return &xxx_UseEnumOperation{
		Info:         o.Info,
		TotalEntries: o.TotalEntries,
		Resume:       o.Resume,
		Return:       o.Return,
	}
}

func (o *UseEnumResponse) xxx_FromOp(ctx context.Context, op *xxx_UseEnumOperation) {
	if o == nil {
		return
	}
	o.Info = op.Info
	o.TotalEntries = op.TotalEntries
	o.Resume = op.Resume
	o.Return = op.Return
}
func (o *UseEnumResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *UseEnumResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_UseEnumOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_WorkstationStatisticsGetOperation structure represents the NetrWorkstationStatisticsGet operation
type xxx_WorkstationStatisticsGetOperation struct {
	ServerName  string            `idl:"name:ServerName;string;pointer:unique" json:"server_name"`
	ServiceName string            `idl:"name:ServiceName;string;pointer:unique" json:"service_name"`
	Level       uint32            `idl:"name:Level" json:"level"`
	Options     uint32            `idl:"name:Options" json:"options"`
	Buffer      *StatWorkstation0 `idl:"name:Buffer" json:"buffer"`
	Return      uint32            `idl:"name:Return" json:"return"`
}

func (o *xxx_WorkstationStatisticsGetOperation) OpNum() int { return 13 }

func (o *xxx_WorkstationStatisticsGetOperation) OpName() string {
	return "/wkssvc/v1/NetrWorkstationStatisticsGet"
}

func (o *xxx_WorkstationStatisticsGetOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WorkstationStatisticsGetOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// ServerName {in} (1:{handle, string, pointer=unique, alias=WKSSVC_IDENTIFY_HANDLE}*(1)[dim:0,string,null](wchar))
	{
		if o.ServerName != "" {
			_ptr_ServerName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.ServerName); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.ServerName, _ptr_ServerName); err != nil {
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
	// ServiceName {in} (1:{string, pointer=unique}*(1)[dim:0,string,null](wchar))
	{
		if o.ServiceName != "" {
			_ptr_ServiceName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.ServiceName); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.ServiceName, _ptr_ServiceName); err != nil {
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
	// Level {in} (1:(uint32))
	{
		if err := w.WriteData(o.Level); err != nil {
			return err
		}
	}
	// Options {in} (1:(uint32))
	{
		if err := w.WriteData(o.Options); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WorkstationStatisticsGetOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// ServerName {in} (1:{handle, string, pointer=unique, alias=WKSSVC_IDENTIFY_HANDLE}*(1)[dim:0,string,null](wchar))
	{
		_ptr_ServerName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.ServerName); err != nil {
				return err
			}
			return nil
		})
		_s_ServerName := func(ptr interface{}) { o.ServerName = *ptr.(*string) }
		if err := w.ReadPointer(&o.ServerName, _s_ServerName, _ptr_ServerName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// ServiceName {in} (1:{string, pointer=unique}*(1)[dim:0,string,null](wchar))
	{
		_ptr_ServiceName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.ServiceName); err != nil {
				return err
			}
			return nil
		})
		_s_ServiceName := func(ptr interface{}) { o.ServiceName = *ptr.(*string) }
		if err := w.ReadPointer(&o.ServiceName, _s_ServiceName, _ptr_ServiceName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Level {in} (1:(uint32))
	{
		if err := w.ReadData(&o.Level); err != nil {
			return err
		}
	}
	// Options {in} (1:(uint32))
	{
		if err := w.ReadData(&o.Options); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WorkstationStatisticsGetOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WorkstationStatisticsGetOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Buffer {out} (1:{pointer=ref}*(2))(2:{alias=LPSTAT_WORKSTATION_0}*(1))(3:{alias=STAT_WORKSTATION_0}(struct))
	{
		if o.Buffer != nil {
			_ptr_Buffer := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Buffer != nil {
					if err := o.Buffer.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&StatWorkstation0{}).MarshalNDR(ctx, w); err != nil {
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
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:(uint32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WorkstationStatisticsGetOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Buffer {out} (1:{pointer=ref}*(2))(2:{alias=LPSTAT_WORKSTATION_0,pointer=ref}*(1))(3:{alias=STAT_WORKSTATION_0}(struct))
	{
		_ptr_Buffer := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Buffer == nil {
				o.Buffer = &StatWorkstation0{}
			}
			if err := o.Buffer.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_Buffer := func(ptr interface{}) { o.Buffer = *ptr.(**StatWorkstation0) }
		if err := w.ReadPointer(&o.Buffer, _s_Buffer, _ptr_Buffer); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// WorkstationStatisticsGetRequest structure represents the NetrWorkstationStatisticsGet operation request
type WorkstationStatisticsGetRequest struct {
	// ServerName: A WKSSVC_IDENTIFY_HANDLE structure (section 2.2.2.1) that identifies
	// the server (2). The client MUST map this structure to an RPC binding handle ([C706]
	// sections 4.3.5 and 5.1.5.2). The server MUST ignore this parameter.
	ServerName string `idl:"name:ServerName;string;pointer:unique" json:"server_name"`
	// ServiceName: A pointer to a string specifying the name of the workstation service.
	// This value MUST be ignored on receipt.
	ServiceName string `idl:"name:ServiceName;string;pointer:unique" json:"service_name"`
	// Level: The information level of the data. This value MUST be zero.
	Level uint32 `idl:"name:Level" json:"level"`
	// Options: This value MUST be zero.
	Options uint32 `idl:"name:Options" json:"options"`
}

func (o *WorkstationStatisticsGetRequest) xxx_ToOp(ctx context.Context) *xxx_WorkstationStatisticsGetOperation {
	if o == nil {
		return &xxx_WorkstationStatisticsGetOperation{}
	}
	return &xxx_WorkstationStatisticsGetOperation{
		ServerName:  o.ServerName,
		ServiceName: o.ServiceName,
		Level:       o.Level,
		Options:     o.Options,
	}
}

func (o *WorkstationStatisticsGetRequest) xxx_FromOp(ctx context.Context, op *xxx_WorkstationStatisticsGetOperation) {
	if o == nil {
		return
	}
	o.ServerName = op.ServerName
	o.ServiceName = op.ServiceName
	o.Level = op.Level
	o.Options = op.Options
}
func (o *WorkstationStatisticsGetRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *WorkstationStatisticsGetRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_WorkstationStatisticsGetOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// WorkstationStatisticsGetResponse structure represents the NetrWorkstationStatisticsGet operation response
type WorkstationStatisticsGetResponse struct {
	// Buffer: A pointer to a STAT_WORKSTATION_0 structure (section 2.2.5.11) that contains
	// the statistical information.
	Buffer *StatWorkstation0 `idl:"name:Buffer" json:"buffer"`
	// Return: The NetrWorkstationStatisticsGet return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *WorkstationStatisticsGetResponse) xxx_ToOp(ctx context.Context) *xxx_WorkstationStatisticsGetOperation {
	if o == nil {
		return &xxx_WorkstationStatisticsGetOperation{}
	}
	return &xxx_WorkstationStatisticsGetOperation{
		Buffer: o.Buffer,
		Return: o.Return,
	}
}

func (o *WorkstationStatisticsGetResponse) xxx_FromOp(ctx context.Context, op *xxx_WorkstationStatisticsGetOperation) {
	if o == nil {
		return
	}
	o.Buffer = op.Buffer
	o.Return = op.Return
}
func (o *WorkstationStatisticsGetResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *WorkstationStatisticsGetResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_WorkstationStatisticsGetOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetJoinInformationOperation structure represents the NetrGetJoinInformation operation
type xxx_GetJoinInformationOperation struct {
	ServerName string             `idl:"name:ServerName;string;pointer:unique" json:"server_name"`
	NameBuffer string             `idl:"name:NameBuffer;string" json:"name_buffer"`
	BufferType NetsetupJoinStatus `idl:"name:BufferType" json:"buffer_type"`
	Return     uint32             `idl:"name:Return" json:"return"`
}

func (o *xxx_GetJoinInformationOperation) OpNum() int { return 20 }

func (o *xxx_GetJoinInformationOperation) OpName() string { return "/wkssvc/v1/NetrGetJoinInformation" }

func (o *xxx_GetJoinInformationOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetJoinInformationOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// ServerName {in} (1:{handle, string, pointer=unique, alias=WKSSVC_IMPERSONATE_HANDLE}*(1)[dim:0,string,null](wchar))
	{
		if o.ServerName != "" {
			_ptr_ServerName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.ServerName); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.ServerName, _ptr_ServerName); err != nil {
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
	// NameBuffer {in, out} (1:{string, pointer=ref}*(2)*(1)[dim:0,string,null](wchar))
	{
		if o.NameBuffer != "" {
			_ptr_NameBuffer := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.NameBuffer); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.NameBuffer, _ptr_NameBuffer); err != nil {
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

func (o *xxx_GetJoinInformationOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// ServerName {in} (1:{handle, string, pointer=unique, alias=WKSSVC_IMPERSONATE_HANDLE}*(1)[dim:0,string,null](wchar))
	{
		_ptr_ServerName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.ServerName); err != nil {
				return err
			}
			return nil
		})
		_s_ServerName := func(ptr interface{}) { o.ServerName = *ptr.(*string) }
		if err := w.ReadPointer(&o.ServerName, _s_ServerName, _ptr_ServerName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// NameBuffer {in, out} (1:{string, pointer=ref}*(2)*(1)[dim:0,string,null](wchar))
	{
		_ptr_NameBuffer := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.NameBuffer); err != nil {
				return err
			}
			return nil
		})
		_s_NameBuffer := func(ptr interface{}) { o.NameBuffer = *ptr.(*string) }
		if err := w.ReadPointer(&o.NameBuffer, _s_NameBuffer, _ptr_NameBuffer); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetJoinInformationOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetJoinInformationOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// NameBuffer {in, out} (1:{string, pointer=ref}*(2)*(1)[dim:0,string,null](wchar))
	{
		if o.NameBuffer != "" {
			_ptr_NameBuffer := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.NameBuffer); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.NameBuffer, _ptr_NameBuffer); err != nil {
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
	// BufferType {out} (1:{alias=PNETSETUP_JOIN_STATUS}*(1))(2:{alias=NETSETUP_JOIN_STATUS}(enum))
	{
		if err := w.WriteData(uint16(o.BufferType)); err != nil {
			return err
		}
	}
	// Return {out} (1:(uint32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetJoinInformationOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// NameBuffer {in, out} (1:{string, pointer=ref}*(2)*(1)[dim:0,string,null](wchar))
	{
		_ptr_NameBuffer := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.NameBuffer); err != nil {
				return err
			}
			return nil
		})
		_s_NameBuffer := func(ptr interface{}) { o.NameBuffer = *ptr.(*string) }
		if err := w.ReadPointer(&o.NameBuffer, _s_NameBuffer, _ptr_NameBuffer); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// BufferType {out} (1:{alias=PNETSETUP_JOIN_STATUS,pointer=ref}*(1))(2:{alias=NETSETUP_JOIN_STATUS}(enum))
	{
		if err := w.ReadData((*uint16)(&o.BufferType)); err != nil {
			return err
		}
	}
	// Return {out} (1:(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// GetJoinInformationRequest structure represents the NetrGetJoinInformation operation request
type GetJoinInformationRequest struct {
	// ServerName: A WKSSVC_IMPERSONATE_HANDLE structure (section 2.2.2.2) that specifies
	// the server (2). The client MUST map this structure to an RPC binding handle ([C706]
	// sections 4.3.5 and 5.1.5.2). The server MUST ignore this parameter.
	ServerName string `idl:"name:ServerName;string;pointer:unique" json:"server_name"`
	// NameBuffer: A pointer to the address of the buffer that receives the name of the
	// domain or workgroup to which the computer is joined, and that also specifies the
	// computer name as input. The server MUST ignore this parameter on input.
	NameBuffer string `idl:"name:NameBuffer;string" json:"name_buffer"`
}

func (o *GetJoinInformationRequest) xxx_ToOp(ctx context.Context) *xxx_GetJoinInformationOperation {
	if o == nil {
		return &xxx_GetJoinInformationOperation{}
	}
	return &xxx_GetJoinInformationOperation{
		ServerName: o.ServerName,
		NameBuffer: o.NameBuffer,
	}
}

func (o *GetJoinInformationRequest) xxx_FromOp(ctx context.Context, op *xxx_GetJoinInformationOperation) {
	if o == nil {
		return
	}
	o.ServerName = op.ServerName
	o.NameBuffer = op.NameBuffer
}
func (o *GetJoinInformationRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetJoinInformationRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetJoinInformationOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetJoinInformationResponse structure represents the NetrGetJoinInformation operation response
type GetJoinInformationResponse struct {
	// NameBuffer: A pointer to the address of the buffer that receives the name of the
	// domain or workgroup to which the computer is joined, and that also specifies the
	// computer name as input. The server MUST ignore this parameter on input.
	NameBuffer string `idl:"name:NameBuffer;string" json:"name_buffer"`
	// BufferType: A pointer to a value from the NETSETUP_JOIN_STATUS enumeration (section
	// 2.2.3.1) that specifies the status of a workstation.
	BufferType NetsetupJoinStatus `idl:"name:BufferType" json:"buffer_type"`
	// Return: The NetrGetJoinInformation return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *GetJoinInformationResponse) xxx_ToOp(ctx context.Context) *xxx_GetJoinInformationOperation {
	if o == nil {
		return &xxx_GetJoinInformationOperation{}
	}
	return &xxx_GetJoinInformationOperation{
		NameBuffer: o.NameBuffer,
		BufferType: o.BufferType,
		Return:     o.Return,
	}
}

func (o *GetJoinInformationResponse) xxx_FromOp(ctx context.Context, op *xxx_GetJoinInformationOperation) {
	if o == nil {
		return
	}
	o.NameBuffer = op.NameBuffer
	o.BufferType = op.BufferType
	o.Return = op.Return
}
func (o *GetJoinInformationResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetJoinInformationResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetJoinInformationOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_JoinDomain2Operation structure represents the NetrJoinDomain2 operation
type xxx_JoinDomain2Operation struct {
	ServerName       string                 `idl:"name:ServerName;string;pointer:unique" json:"server_name"`
	DomainNameParam  string                 `idl:"name:DomainNameParam;string" json:"domain_name_param"`
	MachineAccountOU string                 `idl:"name:MachineAccountOU;string;pointer:unique" json:"machine_account_ou"`
	AccountName      string                 `idl:"name:AccountName;string;pointer:unique" json:"account_name"`
	Password         *EncryptedUserPassword `idl:"name:Password;pointer:unique" json:"password"`
	Options          uint32                 `idl:"name:Options" json:"options"`
	Return           uint32                 `idl:"name:Return" json:"return"`
}

func (o *xxx_JoinDomain2Operation) OpNum() int { return 22 }

func (o *xxx_JoinDomain2Operation) OpName() string { return "/wkssvc/v1/NetrJoinDomain2" }

func (o *xxx_JoinDomain2Operation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_JoinDomain2Operation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// ServerName {in} (1:{string, pointer=unique}*(1)[dim:0,string,null](wchar))
	{
		if o.ServerName != "" {
			_ptr_ServerName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.ServerName); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.ServerName, _ptr_ServerName); err != nil {
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
	// DomainNameParam {in} (1:{string, pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.DomainNameParam); err != nil {
			return err
		}
	}
	// MachineAccountOU {in} (1:{string, pointer=unique}*(1)[dim:0,string,null](wchar))
	{
		if o.MachineAccountOU != "" {
			_ptr_MachineAccountOU := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.MachineAccountOU); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.MachineAccountOU, _ptr_MachineAccountOU); err != nil {
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
	// AccountName {in} (1:{string, pointer=unique}*(1)[dim:0,string,null](wchar))
	{
		if o.AccountName != "" {
			_ptr_AccountName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.AccountName); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.AccountName, _ptr_AccountName); err != nil {
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
	// Password {in} (1:{pointer=unique, alias=PJOINPR_ENCRYPTED_USER_PASSWORD}*(1))(2:{alias=JOINPR_ENCRYPTED_USER_PASSWORD}(struct))
	{
		if o.Password != nil {
			_ptr_Password := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Password != nil {
					if err := o.Password.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&EncryptedUserPassword{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Password, _ptr_Password); err != nil {
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
	// Options {in} (1:(uint32))
	{
		if err := w.WriteData(o.Options); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_JoinDomain2Operation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// ServerName {in} (1:{string, pointer=unique}*(1)[dim:0,string,null](wchar))
	{
		_ptr_ServerName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.ServerName); err != nil {
				return err
			}
			return nil
		})
		_s_ServerName := func(ptr interface{}) { o.ServerName = *ptr.(*string) }
		if err := w.ReadPointer(&o.ServerName, _s_ServerName, _ptr_ServerName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// DomainNameParam {in} (1:{string, pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.DomainNameParam); err != nil {
			return err
		}
	}
	// MachineAccountOU {in} (1:{string, pointer=unique}*(1)[dim:0,string,null](wchar))
	{
		_ptr_MachineAccountOU := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.MachineAccountOU); err != nil {
				return err
			}
			return nil
		})
		_s_MachineAccountOU := func(ptr interface{}) { o.MachineAccountOU = *ptr.(*string) }
		if err := w.ReadPointer(&o.MachineAccountOU, _s_MachineAccountOU, _ptr_MachineAccountOU); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// AccountName {in} (1:{string, pointer=unique}*(1)[dim:0,string,null](wchar))
	{
		_ptr_AccountName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.AccountName); err != nil {
				return err
			}
			return nil
		})
		_s_AccountName := func(ptr interface{}) { o.AccountName = *ptr.(*string) }
		if err := w.ReadPointer(&o.AccountName, _s_AccountName, _ptr_AccountName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Password {in} (1:{pointer=unique, alias=PJOINPR_ENCRYPTED_USER_PASSWORD}*(1))(2:{alias=JOINPR_ENCRYPTED_USER_PASSWORD}(struct))
	{
		_ptr_Password := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Password == nil {
				o.Password = &EncryptedUserPassword{}
			}
			if err := o.Password.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_Password := func(ptr interface{}) { o.Password = *ptr.(**EncryptedUserPassword) }
		if err := w.ReadPointer(&o.Password, _s_Password, _ptr_Password); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Options {in} (1:(uint32))
	{
		if err := w.ReadData(&o.Options); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_JoinDomain2Operation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_JoinDomain2Operation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:(uint32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_JoinDomain2Operation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// JoinDomain2Request structure represents the NetrJoinDomain2 operation request
type JoinDomain2Request struct {
	// ServerName: This parameter has no effect on message processing in any environment.
	// The client MUST set this parameter to a value that resolves to the IP protocol layer
	// destination address of the RPC packets it transmits ([MS-RPCE] section 2.1.1.2).
	// The server (2) MUST ignore this parameter.
	ServerName string `idl:"name:ServerName;string;pointer:unique" json:"server_name"`
	// DomainNameParam: A pointer to a string that specifies the domain name (2) or workgroup
	// name to join, and optionally the domain controller machine name within the domain.
	// This parameter MUST NOT be NULL.
	//
	// * NetBIOS name\NetBIOS name
	//
	// * NetBIOS name\Internet host name
	//
	// * FQDN (2)\NetBIOS name
	DomainNameParam string `idl:"name:DomainNameParam;string" json:"domain_name_param"`
	// MachineAccountOU: A pointer to a string that contains [RFC1777] the format name of
	// the organizational unit (OU) directory object under which the machine account directory
	// object is created. This parameter is optional. If specified, this string MUST contain
	// the full path; for example, OU=testOU,DC=domain,DC=Domain,DC=com.
	MachineAccountOU string `idl:"name:MachineAccountOU;string;pointer:unique" json:"machine_account_ou"`
	// AccountName: A pointer to a string that specifies an account name in the domain DomainNameParam
	// to use when connecting to a domain controller. This parameter is optional. If this
	// parameter is NULL, the caller's account name MUST be used. If this parameter is specified,
	// the format MUST be one of the following:
	//
	// * <NetBIOSDomainName>\<UserName>
	//
	// * <FullyQualifiedDNSDomainName>\<UserName>
	//
	// * <UserName>@<FullyQualifiedDNSDomainName>
	AccountName string `idl:"name:AccountName;string;pointer:unique" json:"account_name"`
	// Password: A pointer to a JOINPR_ENCRYPTED_USER_PASSWORD structure (section 2.2.5.18)
	// that specifies the encrypted password to use with the AccountName parameter. Sections
	// 3.2.4.13.1 and 3.2.4.13.3 specify the processing of this parameter.
	Password *EncryptedUserPassword `idl:"name:Password;pointer:unique" json:"password"`
	// Options: A 32-bit bitfield that specifies modifications to default server (2) behavior
	// in message processing.<58>
	//
	//	+-------------------------------------------+----------------------------------------------------------------------------------+
	//	|                                           |                                                                                  |
	//	|                VALUE/CODE                 |                                     MEANING                                      |
	//	|                                           |                                                                                  |
	//	+-------------------------------------------+----------------------------------------------------------------------------------+
	//	+-------------------------------------------+----------------------------------------------------------------------------------+
	//	| NETSETUP_JOIN_DOMAIN 0x00000001           | Joins the computer to a domain. The default action is to join the computer to a  |
	//	|                                           | workgroup.                                                                       |
	//	+-------------------------------------------+----------------------------------------------------------------------------------+
	//	| NETSETUP_ACCT_CREATE 0x00000002           | Creates the account on the domain. The name is the persisted abstract state      |
	//	|                                           | ComputerNameNetBIOS (section 3.2.1.5) unless this behavior is altered by another |
	//	|                                           | option such as NETSETUP_JOIN_WITH_NEW_NAME.                                      |
	//	+-------------------------------------------+----------------------------------------------------------------------------------+
	//	| NETSETUP_ACCT_DELETE 0x00000004           | Disables the old account when the join operation occurs on a computer that is    |
	//	|                                           | already joined to a domain. Important  This flag is neither supported nor tested |
	//	|                                           | for use with NetrJoinDomain2; therefore, its use is not specified in any message |
	//	|                                           | processing.                                                                      |
	//	+-------------------------------------------+----------------------------------------------------------------------------------+
	//	| NETSETUP_DOMAIN_JOIN_IF_JOINED 0x00000020 | Allows a join to a new domain even if the computer is already joined to a        |
	//	|                                           | domain.                                                                          |
	//	+-------------------------------------------+----------------------------------------------------------------------------------+
	//	| NETSETUP_JOIN_UNSECURE 0x00000040         | Performs an unsecured join. It MUST be used only in conjunction with the         |
	//	|                                           | NETSETUP_MACHINE_PWD_PASSED flag.                                                |
	//	+-------------------------------------------+----------------------------------------------------------------------------------+
	//	| NETSETUP_MACHINE_PWD_PASSED 0x00000080    | Indicates that the Password parameter SHOULD<59> specify the password for the    |
	//	|                                           | machine joining the domain. This flag is valid for unsecured joins, which are    |
	//	|                                           | indicated by setting the NETSETUP_JOIN_UNSECURE flag, or for read-only joins,    |
	//	|                                           | which are indicated by setting the NETSETUP_JOIN_READONLY flag. If this flag is  |
	//	|                                           | set, the value of Password determines the value stored for the computer password |
	//	|                                           | during the join process.                                                         |
	//	+-------------------------------------------+----------------------------------------------------------------------------------+
	//	| NETSETUP_DEFER_SPN_SET 0x00000100         | Indicates that the service principal name (SPN) and the DnsHostName properties   |
	//	|                                           | on the computer SHOULD NOT<60> be updated at this time, but instead SHOULD<61>   |
	//	|                                           | be updated during a subsequent call to NetrRenameMachineInDomain2 (section       |
	//	|                                           | 3.2.4.15).                                                                       |
	//	+-------------------------------------------+----------------------------------------------------------------------------------+
	//	| NETSETUP_JOIN_DC_ACCOUNT 0x00000200       | Indicates that the join SHOULD<62> be allowed if an existing account exists and  |
	//	|                                           | it is a domain controller account.<63>                                           |
	//	+-------------------------------------------+----------------------------------------------------------------------------------+
	//	| NETSETUP_JOIN_WITH_NEW_NAME 0x00000400    | Indicates that the join SHOULD<64>occur using the new computer name.             |
	//	+-------------------------------------------+----------------------------------------------------------------------------------+
	//	| NETSETUP_JOIN_READONLY 0x00000800         | Specifies that the join SHOULD<65> be performed in a read-only manner against an |
	//	|                                           | existing account object. This option is intended to enable the server to join a  |
	//	|                                           | domain using a read-only domain controller.                                      |
	//	+-------------------------------------------+----------------------------------------------------------------------------------+
	//	| NETSETUP_INSTALL_INVOCATION 0x00040000    | Indicates that the protocol method was invoked during installation.              |
	//	+-------------------------------------------+----------------------------------------------------------------------------------+
	Options uint32 `idl:"name:Options" json:"options"`
}

func (o *JoinDomain2Request) xxx_ToOp(ctx context.Context) *xxx_JoinDomain2Operation {
	if o == nil {
		return &xxx_JoinDomain2Operation{}
	}
	return &xxx_JoinDomain2Operation{
		ServerName:       o.ServerName,
		DomainNameParam:  o.DomainNameParam,
		MachineAccountOU: o.MachineAccountOU,
		AccountName:      o.AccountName,
		Password:         o.Password,
		Options:          o.Options,
	}
}

func (o *JoinDomain2Request) xxx_FromOp(ctx context.Context, op *xxx_JoinDomain2Operation) {
	if o == nil {
		return
	}
	o.ServerName = op.ServerName
	o.DomainNameParam = op.DomainNameParam
	o.MachineAccountOU = op.MachineAccountOU
	o.AccountName = op.AccountName
	o.Password = op.Password
	o.Options = op.Options
}
func (o *JoinDomain2Request) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *JoinDomain2Request) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_JoinDomain2Operation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// JoinDomain2Response structure represents the NetrJoinDomain2 operation response
type JoinDomain2Response struct {
	// Return: The NetrJoinDomain2 return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *JoinDomain2Response) xxx_ToOp(ctx context.Context) *xxx_JoinDomain2Operation {
	if o == nil {
		return &xxx_JoinDomain2Operation{}
	}
	return &xxx_JoinDomain2Operation{
		Return: o.Return,
	}
}

func (o *JoinDomain2Response) xxx_FromOp(ctx context.Context, op *xxx_JoinDomain2Operation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *JoinDomain2Response) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *JoinDomain2Response) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_JoinDomain2Operation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_UnjoinDomain2Operation structure represents the NetrUnjoinDomain2 operation
type xxx_UnjoinDomain2Operation struct {
	ServerName  string                 `idl:"name:ServerName;string;pointer:unique" json:"server_name"`
	AccountName string                 `idl:"name:AccountName;string;pointer:unique" json:"account_name"`
	Password    *EncryptedUserPassword `idl:"name:Password;pointer:unique" json:"password"`
	Options     uint32                 `idl:"name:Options" json:"options"`
	Return      uint32                 `idl:"name:Return" json:"return"`
}

func (o *xxx_UnjoinDomain2Operation) OpNum() int { return 23 }

func (o *xxx_UnjoinDomain2Operation) OpName() string { return "/wkssvc/v1/NetrUnjoinDomain2" }

func (o *xxx_UnjoinDomain2Operation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_UnjoinDomain2Operation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// ServerName {in} (1:{string, pointer=unique}*(1)[dim:0,string,null](wchar))
	{
		if o.ServerName != "" {
			_ptr_ServerName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.ServerName); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.ServerName, _ptr_ServerName); err != nil {
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
	// AccountName {in} (1:{string, pointer=unique}*(1)[dim:0,string,null](wchar))
	{
		if o.AccountName != "" {
			_ptr_AccountName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.AccountName); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.AccountName, _ptr_AccountName); err != nil {
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
	// Password {in} (1:{pointer=unique, alias=PJOINPR_ENCRYPTED_USER_PASSWORD}*(1))(2:{alias=JOINPR_ENCRYPTED_USER_PASSWORD}(struct))
	{
		if o.Password != nil {
			_ptr_Password := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Password != nil {
					if err := o.Password.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&EncryptedUserPassword{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Password, _ptr_Password); err != nil {
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
	// Options {in} (1:(uint32))
	{
		if err := w.WriteData(o.Options); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_UnjoinDomain2Operation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// ServerName {in} (1:{string, pointer=unique}*(1)[dim:0,string,null](wchar))
	{
		_ptr_ServerName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.ServerName); err != nil {
				return err
			}
			return nil
		})
		_s_ServerName := func(ptr interface{}) { o.ServerName = *ptr.(*string) }
		if err := w.ReadPointer(&o.ServerName, _s_ServerName, _ptr_ServerName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// AccountName {in} (1:{string, pointer=unique}*(1)[dim:0,string,null](wchar))
	{
		_ptr_AccountName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.AccountName); err != nil {
				return err
			}
			return nil
		})
		_s_AccountName := func(ptr interface{}) { o.AccountName = *ptr.(*string) }
		if err := w.ReadPointer(&o.AccountName, _s_AccountName, _ptr_AccountName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Password {in} (1:{pointer=unique, alias=PJOINPR_ENCRYPTED_USER_PASSWORD}*(1))(2:{alias=JOINPR_ENCRYPTED_USER_PASSWORD}(struct))
	{
		_ptr_Password := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Password == nil {
				o.Password = &EncryptedUserPassword{}
			}
			if err := o.Password.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_Password := func(ptr interface{}) { o.Password = *ptr.(**EncryptedUserPassword) }
		if err := w.ReadPointer(&o.Password, _s_Password, _ptr_Password); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Options {in} (1:(uint32))
	{
		if err := w.ReadData(&o.Options); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_UnjoinDomain2Operation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_UnjoinDomain2Operation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:(uint32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_UnjoinDomain2Operation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// UnjoinDomain2Request structure represents the NetrUnjoinDomain2 operation request
type UnjoinDomain2Request struct {
	// ServerName: This parameter has no effect on message processing in any environment.
	// The client MUST set this parameter to a value that resolves to the IP destination
	// address of the RPC packets it transmits ([MS-RPCE] section 2.1.1.2). The server (2)
	// MUST ignore this parameter.
	ServerName string `idl:"name:ServerName;string;pointer:unique" json:"server_name"`
	// AccountName: A pointer to a string that specifies the account name in the joined
	// domain to use when connecting to a domain controller. This parameter is optional.
	// If this parameter is NULL, the caller's account name MUST be used.
	AccountName string `idl:"name:AccountName;string;pointer:unique" json:"account_name"`
	// Password: An optional pointer to a JOINPR_ENCRYPTED_USER_PASSWORD structure (section
	// 2.2.5.18) that specifies the encrypted password to use with the AccountName parameter.
	// If this parameter is NULL, the caller's security context MUST be used.
	Password *EncryptedUserPassword `idl:"name:Password;pointer:unique" json:"password"`
	// Options: A 32-bit bitfield specifying modifications to default message processing
	// behavior.
	//
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	|                                              |                                                                                  |
	//	|                  VALUE/CODE                  |                                     MEANING                                      |
	//	|                                              |                                                                                  |
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	| NETSETUP_ACCT_DELETE 0x00000004              | Disables the account when the unjoin operation occurs.                           |
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	| NETSETUP_IGNORE_UNSUPPORTED_FLAGS 0x10000000 | The server ignores undefined flags when this bit is set.<77> This option is      |
	//	|                                              | present to allow for the addition of new optional values in the future.          |
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	Options uint32 `idl:"name:Options" json:"options"`
}

func (o *UnjoinDomain2Request) xxx_ToOp(ctx context.Context) *xxx_UnjoinDomain2Operation {
	if o == nil {
		return &xxx_UnjoinDomain2Operation{}
	}
	return &xxx_UnjoinDomain2Operation{
		ServerName:  o.ServerName,
		AccountName: o.AccountName,
		Password:    o.Password,
		Options:     o.Options,
	}
}

func (o *UnjoinDomain2Request) xxx_FromOp(ctx context.Context, op *xxx_UnjoinDomain2Operation) {
	if o == nil {
		return
	}
	o.ServerName = op.ServerName
	o.AccountName = op.AccountName
	o.Password = op.Password
	o.Options = op.Options
}
func (o *UnjoinDomain2Request) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *UnjoinDomain2Request) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_UnjoinDomain2Operation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// UnjoinDomain2Response structure represents the NetrUnjoinDomain2 operation response
type UnjoinDomain2Response struct {
	// Return: The NetrUnjoinDomain2 return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *UnjoinDomain2Response) xxx_ToOp(ctx context.Context) *xxx_UnjoinDomain2Operation {
	if o == nil {
		return &xxx_UnjoinDomain2Operation{}
	}
	return &xxx_UnjoinDomain2Operation{
		Return: o.Return,
	}
}

func (o *UnjoinDomain2Response) xxx_FromOp(ctx context.Context, op *xxx_UnjoinDomain2Operation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *UnjoinDomain2Response) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *UnjoinDomain2Response) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_UnjoinDomain2Operation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_RenameMachineInDomain2Operation structure represents the NetrRenameMachineInDomain2 operation
type xxx_RenameMachineInDomain2Operation struct {
	ServerName  string                 `idl:"name:ServerName;string;pointer:unique" json:"server_name"`
	MachineName string                 `idl:"name:MachineName;string;pointer:unique" json:"machine_name"`
	AccountName string                 `idl:"name:AccountName;string;pointer:unique" json:"account_name"`
	Password    *EncryptedUserPassword `idl:"name:Password;pointer:unique" json:"password"`
	Options     uint32                 `idl:"name:Options" json:"options"`
	Return      uint32                 `idl:"name:Return" json:"return"`
}

func (o *xxx_RenameMachineInDomain2Operation) OpNum() int { return 24 }

func (o *xxx_RenameMachineInDomain2Operation) OpName() string {
	return "/wkssvc/v1/NetrRenameMachineInDomain2"
}

func (o *xxx_RenameMachineInDomain2Operation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RenameMachineInDomain2Operation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// ServerName {in} (1:{string, pointer=unique}*(1)[dim:0,string,null](wchar))
	{
		if o.ServerName != "" {
			_ptr_ServerName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.ServerName); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.ServerName, _ptr_ServerName); err != nil {
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
	// MachineName {in} (1:{string, pointer=unique}*(1)[dim:0,string,null](wchar))
	{
		if o.MachineName != "" {
			_ptr_MachineName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.MachineName); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.MachineName, _ptr_MachineName); err != nil {
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
	// AccountName {in} (1:{string, pointer=unique}*(1)[dim:0,string,null](wchar))
	{
		if o.AccountName != "" {
			_ptr_AccountName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.AccountName); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.AccountName, _ptr_AccountName); err != nil {
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
	// Password {in} (1:{pointer=unique, alias=PJOINPR_ENCRYPTED_USER_PASSWORD}*(1))(2:{alias=JOINPR_ENCRYPTED_USER_PASSWORD}(struct))
	{
		if o.Password != nil {
			_ptr_Password := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Password != nil {
					if err := o.Password.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&EncryptedUserPassword{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Password, _ptr_Password); err != nil {
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
	// Options {in} (1:(uint32))
	{
		if err := w.WriteData(o.Options); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RenameMachineInDomain2Operation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// ServerName {in} (1:{string, pointer=unique}*(1)[dim:0,string,null](wchar))
	{
		_ptr_ServerName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.ServerName); err != nil {
				return err
			}
			return nil
		})
		_s_ServerName := func(ptr interface{}) { o.ServerName = *ptr.(*string) }
		if err := w.ReadPointer(&o.ServerName, _s_ServerName, _ptr_ServerName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// MachineName {in} (1:{string, pointer=unique}*(1)[dim:0,string,null](wchar))
	{
		_ptr_MachineName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.MachineName); err != nil {
				return err
			}
			return nil
		})
		_s_MachineName := func(ptr interface{}) { o.MachineName = *ptr.(*string) }
		if err := w.ReadPointer(&o.MachineName, _s_MachineName, _ptr_MachineName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// AccountName {in} (1:{string, pointer=unique}*(1)[dim:0,string,null](wchar))
	{
		_ptr_AccountName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.AccountName); err != nil {
				return err
			}
			return nil
		})
		_s_AccountName := func(ptr interface{}) { o.AccountName = *ptr.(*string) }
		if err := w.ReadPointer(&o.AccountName, _s_AccountName, _ptr_AccountName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Password {in} (1:{pointer=unique, alias=PJOINPR_ENCRYPTED_USER_PASSWORD}*(1))(2:{alias=JOINPR_ENCRYPTED_USER_PASSWORD}(struct))
	{
		_ptr_Password := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Password == nil {
				o.Password = &EncryptedUserPassword{}
			}
			if err := o.Password.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_Password := func(ptr interface{}) { o.Password = *ptr.(**EncryptedUserPassword) }
		if err := w.ReadPointer(&o.Password, _s_Password, _ptr_Password); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Options {in} (1:(uint32))
	{
		if err := w.ReadData(&o.Options); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RenameMachineInDomain2Operation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RenameMachineInDomain2Operation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:(uint32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RenameMachineInDomain2Operation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// RenameMachineInDomain2Request structure represents the NetrRenameMachineInDomain2 operation request
type RenameMachineInDomain2Request struct {
	// ServerName: This parameter has no effect on message processing in any environment.
	// The client MUST set this parameter to a value that resolves to the IP protocol layer
	// destination address of the RPC packets it transmits ([MS-RPCE] section 2.1.1.2).
	// The server (2) MUST ignore this parameter.
	ServerName string `idl:"name:ServerName;string;pointer:unique" json:"server_name"`
	// MachineName: A pointer to a string that specifies the new computer name. This parameter
	// is optional. If this parameter is NULL, the current machine name is used.
	MachineName string `idl:"name:MachineName;string;pointer:unique" json:"machine_name"`
	// AccountName: A pointer to a string that specifies an account name in the joined domain
	// to use when connecting to a domain controller. This parameter is optional. If this
	// parameter is NULL, the caller's account name is used.
	AccountName string `idl:"name:AccountName;string;pointer:unique" json:"account_name"`
	// Password: An optional pointer to a JOINPR_ENCRYPTED_USER_PASSWORD structure (section
	// 2.2.5.18) that specifies the encrypted password to use with the AccountName parameter.
	// If this parameter is NULL, the caller's security context MUST be used.
	Password *EncryptedUserPassword `idl:"name:Password;pointer:unique" json:"password"`
	// Options: A 32-bit bitfield that specifies modifications to default server behavior
	// in message processing.
	//
	//	+-------------------------------------------+----------------------------------------------------------------------------------+
	//	|                                           |                                                                                  |
	//	|                VALUE/CODE                 |                                     MEANING                                      |
	//	|                                           |                                                                                  |
	//	+-------------------------------------------+----------------------------------------------------------------------------------+
	//	+-------------------------------------------+----------------------------------------------------------------------------------+
	//	| NETSETUP_ACCT_CREATE 0x00000002           | Renames the computer account in the domain. If this flag is not set, the         |
	//	|                                           | computer name is changed locally but no changes are made to the computer account |
	//	|                                           | in the domain.                                                                   |
	//	+-------------------------------------------+----------------------------------------------------------------------------------+
	//	| NETSETUP_DNS_NAME_CHANGES_ONLY 0x00001000 | Limits any updates to DNS-based names only.                                      |
	//	+-------------------------------------------+----------------------------------------------------------------------------------+
	Options uint32 `idl:"name:Options" json:"options"`
}

func (o *RenameMachineInDomain2Request) xxx_ToOp(ctx context.Context) *xxx_RenameMachineInDomain2Operation {
	if o == nil {
		return &xxx_RenameMachineInDomain2Operation{}
	}
	return &xxx_RenameMachineInDomain2Operation{
		ServerName:  o.ServerName,
		MachineName: o.MachineName,
		AccountName: o.AccountName,
		Password:    o.Password,
		Options:     o.Options,
	}
}

func (o *RenameMachineInDomain2Request) xxx_FromOp(ctx context.Context, op *xxx_RenameMachineInDomain2Operation) {
	if o == nil {
		return
	}
	o.ServerName = op.ServerName
	o.MachineName = op.MachineName
	o.AccountName = op.AccountName
	o.Password = op.Password
	o.Options = op.Options
}
func (o *RenameMachineInDomain2Request) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *RenameMachineInDomain2Request) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RenameMachineInDomain2Operation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RenameMachineInDomain2Response structure represents the NetrRenameMachineInDomain2 operation response
type RenameMachineInDomain2Response struct {
	// Return: The NetrRenameMachineInDomain2 return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *RenameMachineInDomain2Response) xxx_ToOp(ctx context.Context) *xxx_RenameMachineInDomain2Operation {
	if o == nil {
		return &xxx_RenameMachineInDomain2Operation{}
	}
	return &xxx_RenameMachineInDomain2Operation{
		Return: o.Return,
	}
}

func (o *RenameMachineInDomain2Response) xxx_FromOp(ctx context.Context, op *xxx_RenameMachineInDomain2Operation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *RenameMachineInDomain2Response) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *RenameMachineInDomain2Response) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RenameMachineInDomain2Operation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_ValidateName2Operation structure represents the NetrValidateName2 operation
type xxx_ValidateName2Operation struct {
	ServerName     string                 `idl:"name:ServerName;string;pointer:unique" json:"server_name"`
	NameToValidate string                 `idl:"name:NameToValidate;string" json:"name_to_validate"`
	AccountName    string                 `idl:"name:AccountName;string;pointer:unique" json:"account_name"`
	Password       *EncryptedUserPassword `idl:"name:Password;pointer:unique" json:"password"`
	NameType       NetsetupNameType       `idl:"name:NameType" json:"name_type"`
	Return         uint32                 `idl:"name:Return" json:"return"`
}

func (o *xxx_ValidateName2Operation) OpNum() int { return 25 }

func (o *xxx_ValidateName2Operation) OpName() string { return "/wkssvc/v1/NetrValidateName2" }

func (o *xxx_ValidateName2Operation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ValidateName2Operation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// ServerName {in} (1:{string, pointer=unique}*(1)[dim:0,string,null](wchar))
	{
		if o.ServerName != "" {
			_ptr_ServerName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.ServerName); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.ServerName, _ptr_ServerName); err != nil {
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
	// NameToValidate {in} (1:{string, pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.NameToValidate); err != nil {
			return err
		}
	}
	// AccountName {in} (1:{string, pointer=unique}*(1)[dim:0,string,null](wchar))
	{
		if o.AccountName != "" {
			_ptr_AccountName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.AccountName); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.AccountName, _ptr_AccountName); err != nil {
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
	// Password {in} (1:{pointer=unique, alias=PJOINPR_ENCRYPTED_USER_PASSWORD}*(1))(2:{alias=JOINPR_ENCRYPTED_USER_PASSWORD}(struct))
	{
		if o.Password != nil {
			_ptr_Password := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Password != nil {
					if err := o.Password.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&EncryptedUserPassword{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Password, _ptr_Password); err != nil {
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
	// NameType {in} (1:{alias=NETSETUP_NAME_TYPE}(enum))
	{
		if err := w.WriteData(uint16(o.NameType)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ValidateName2Operation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// ServerName {in} (1:{string, pointer=unique}*(1)[dim:0,string,null](wchar))
	{
		_ptr_ServerName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.ServerName); err != nil {
				return err
			}
			return nil
		})
		_s_ServerName := func(ptr interface{}) { o.ServerName = *ptr.(*string) }
		if err := w.ReadPointer(&o.ServerName, _s_ServerName, _ptr_ServerName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// NameToValidate {in} (1:{string, pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.NameToValidate); err != nil {
			return err
		}
	}
	// AccountName {in} (1:{string, pointer=unique}*(1)[dim:0,string,null](wchar))
	{
		_ptr_AccountName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.AccountName); err != nil {
				return err
			}
			return nil
		})
		_s_AccountName := func(ptr interface{}) { o.AccountName = *ptr.(*string) }
		if err := w.ReadPointer(&o.AccountName, _s_AccountName, _ptr_AccountName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Password {in} (1:{pointer=unique, alias=PJOINPR_ENCRYPTED_USER_PASSWORD}*(1))(2:{alias=JOINPR_ENCRYPTED_USER_PASSWORD}(struct))
	{
		_ptr_Password := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Password == nil {
				o.Password = &EncryptedUserPassword{}
			}
			if err := o.Password.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_Password := func(ptr interface{}) { o.Password = *ptr.(**EncryptedUserPassword) }
		if err := w.ReadPointer(&o.Password, _s_Password, _ptr_Password); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// NameType {in} (1:{alias=NETSETUP_NAME_TYPE}(enum))
	{
		if err := w.ReadData((*uint16)(&o.NameType)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ValidateName2Operation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ValidateName2Operation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:(uint32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ValidateName2Operation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// ValidateName2Request structure represents the NetrValidateName2 operation request
type ValidateName2Request struct {
	// ServerName: This parameter has no effect on message processing in any environment.
	// The client MUST set this parameter to a value that resolves to the IP protocol layer
	// destination address of the RPC packets it transmits ([MS-RPCE] section 2.1.1.2).
	// The server (2) MUST ignore this parameter.
	ServerName string `idl:"name:ServerName;string;pointer:unique" json:"server_name"`
	// NameToValidate: A pointer to a string that specifies the name to validate, according
	// to its type.
	NameToValidate string `idl:"name:NameToValidate;string" json:"name_to_validate"`
	// AccountName: The server SHOULD ignore this parameter.
	AccountName string `idl:"name:AccountName;string;pointer:unique" json:"account_name"`
	// Password: The server SHOULD ignore this parameter.
	Password *EncryptedUserPassword `idl:"name:Password;pointer:unique" json:"password"`
	// NameType: Specifies the type of validation to perform (section 2.2.3.2).
	NameType NetsetupNameType `idl:"name:NameType" json:"name_type"`
}

func (o *ValidateName2Request) xxx_ToOp(ctx context.Context) *xxx_ValidateName2Operation {
	if o == nil {
		return &xxx_ValidateName2Operation{}
	}
	return &xxx_ValidateName2Operation{
		ServerName:     o.ServerName,
		NameToValidate: o.NameToValidate,
		AccountName:    o.AccountName,
		Password:       o.Password,
		NameType:       o.NameType,
	}
}

func (o *ValidateName2Request) xxx_FromOp(ctx context.Context, op *xxx_ValidateName2Operation) {
	if o == nil {
		return
	}
	o.ServerName = op.ServerName
	o.NameToValidate = op.NameToValidate
	o.AccountName = op.AccountName
	o.Password = op.Password
	o.NameType = op.NameType
}
func (o *ValidateName2Request) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *ValidateName2Request) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ValidateName2Operation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ValidateName2Response structure represents the NetrValidateName2 operation response
type ValidateName2Response struct {
	// Return: The NetrValidateName2 return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *ValidateName2Response) xxx_ToOp(ctx context.Context) *xxx_ValidateName2Operation {
	if o == nil {
		return &xxx_ValidateName2Operation{}
	}
	return &xxx_ValidateName2Operation{
		Return: o.Return,
	}
}

func (o *ValidateName2Response) xxx_FromOp(ctx context.Context, op *xxx_ValidateName2Operation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *ValidateName2Response) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *ValidateName2Response) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ValidateName2Operation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetJoinableOUs2Operation structure represents the NetrGetJoinableOUs2 operation
type xxx_GetJoinableOUs2Operation struct {
	ServerName      string                 `idl:"name:ServerName;string;pointer:unique" json:"server_name"`
	DomainNameParam string                 `idl:"name:DomainNameParam;string" json:"domain_name_param"`
	AccountName     string                 `idl:"name:AccountName;string;pointer:unique" json:"account_name"`
	Password        *EncryptedUserPassword `idl:"name:Password;pointer:unique" json:"password"`
	OUCount         uint32                 `idl:"name:OUCount" json:"ou_count"`
	OUs             []string               `idl:"name:OUs;size_is:(, OUCount);string" json:"ous"`
	Return          uint32                 `idl:"name:Return" json:"return"`
}

func (o *xxx_GetJoinableOUs2Operation) OpNum() int { return 26 }

func (o *xxx_GetJoinableOUs2Operation) OpName() string { return "/wkssvc/v1/NetrGetJoinableOUs2" }

func (o *xxx_GetJoinableOUs2Operation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetJoinableOUs2Operation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// ServerName {in} (1:{string, pointer=unique}*(1)[dim:0,string,null](wchar))
	{
		if o.ServerName != "" {
			_ptr_ServerName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.ServerName); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.ServerName, _ptr_ServerName); err != nil {
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
	// DomainNameParam {in} (1:{string, pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.DomainNameParam); err != nil {
			return err
		}
	}
	// AccountName {in} (1:{string, pointer=unique}*(1)[dim:0,string,null](wchar))
	{
		if o.AccountName != "" {
			_ptr_AccountName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.AccountName); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.AccountName, _ptr_AccountName); err != nil {
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
	// Password {in} (1:{pointer=unique, alias=PJOINPR_ENCRYPTED_USER_PASSWORD}*(1))(2:{alias=JOINPR_ENCRYPTED_USER_PASSWORD}(struct))
	{
		if o.Password != nil {
			_ptr_Password := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Password != nil {
					if err := o.Password.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&EncryptedUserPassword{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Password, _ptr_Password); err != nil {
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
	// OUCount {in, out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.WriteData(o.OUCount); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetJoinableOUs2Operation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// ServerName {in} (1:{string, pointer=unique}*(1)[dim:0,string,null](wchar))
	{
		_ptr_ServerName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.ServerName); err != nil {
				return err
			}
			return nil
		})
		_s_ServerName := func(ptr interface{}) { o.ServerName = *ptr.(*string) }
		if err := w.ReadPointer(&o.ServerName, _s_ServerName, _ptr_ServerName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// DomainNameParam {in} (1:{string, pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.DomainNameParam); err != nil {
			return err
		}
	}
	// AccountName {in} (1:{string, pointer=unique}*(1)[dim:0,string,null](wchar))
	{
		_ptr_AccountName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.AccountName); err != nil {
				return err
			}
			return nil
		})
		_s_AccountName := func(ptr interface{}) { o.AccountName = *ptr.(*string) }
		if err := w.ReadPointer(&o.AccountName, _s_AccountName, _ptr_AccountName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Password {in} (1:{pointer=unique, alias=PJOINPR_ENCRYPTED_USER_PASSWORD}*(1))(2:{alias=JOINPR_ENCRYPTED_USER_PASSWORD}(struct))
	{
		_ptr_Password := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Password == nil {
				o.Password = &EncryptedUserPassword{}
			}
			if err := o.Password.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_Password := func(ptr interface{}) { o.Password = *ptr.(**EncryptedUserPassword) }
		if err := w.ReadPointer(&o.Password, _s_Password, _ptr_Password); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// OUCount {in, out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.ReadData(&o.OUCount); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetJoinableOUs2Operation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if o.OUs != nil && o.OUCount == 0 {
		o.OUCount = uint32(len(o.OUs))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetJoinableOUs2Operation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// OUCount {in, out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.WriteData(o.OUCount); err != nil {
			return err
		}
	}
	// OUs {out} (1:{string, pointer=ref}*(2)*(1)[dim:0,size_is=OUCount]*(1)[dim:0,string,null](wchar))
	{
		if o.OUs != nil || o.OUCount > 0 {
			_ptr_OUs := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.OUCount)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				for i1 := range o.OUs {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if o.OUs[i1] != "" {
						_ptr_OUs := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
							if err := ndr.WriteUTF16NString(ctx, w, o.OUs[i1]); err != nil {
								return err
							}
							return nil
						})
						if err := w.WritePointer(&o.OUs[i1], _ptr_OUs); err != nil {
							return err
						}
					} else {
						if err := w.WritePointer(nil); err != nil {
							return err
						}
					}
				}
				for i1 := len(o.OUs); uint64(i1) < sizeInfo[0]; i1++ {
					if err := w.WritePointer(nil); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.OUs, _ptr_OUs); err != nil {
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
	// Return {out} (1:(uint32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetJoinableOUs2Operation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// OUCount {in, out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.ReadData(&o.OUCount); err != nil {
			return err
		}
	}
	// OUs {out} (1:{string, pointer=ref}*(2)*(1)[dim:0,size_is=OUCount]*(1)[dim:0,string,null](wchar))
	{
		_ptr_OUs := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array o.OUs", sizeInfo[0])
			}
			o.OUs = make([]string, sizeInfo[0])
			for i1 := range o.OUs {
				i1 := i1
				_ptr_OUs := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
					if err := ndr.ReadUTF16NString(ctx, w, &o.OUs[i1]); err != nil {
						return err
					}
					return nil
				})
				_s_OUs := func(ptr interface{}) { o.OUs[i1] = *ptr.(*string) }
				if err := w.ReadPointer(&o.OUs[i1], _s_OUs, _ptr_OUs); err != nil {
					return err
				}
			}
			return nil
		})
		_s_OUs := func(ptr interface{}) { o.OUs = *ptr.(*[]string) }
		if err := w.ReadPointer(&o.OUs, _s_OUs, _ptr_OUs); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// GetJoinableOUs2Request structure represents the NetrGetJoinableOUs2 operation request
type GetJoinableOUs2Request struct {
	// ServerName: This parameter has no effect on message processing in any environment.
	// The client MUST set this parameter to a value that resolves to the IP protocol layer
	// destination address of the RPC packets it transmits ([MS-RPCE] section 2.1.1.2).
	// The server (2) MUST ignore this parameter.
	ServerName string `idl:"name:ServerName;string;pointer:unique" json:"server_name"`
	// DomainNameParam: A pointer to a string that specifies the root domain under which
	// the method searches for OUs. This parameter is also the domain of the account that
	// the AccountName parameter is in.
	DomainNameParam string `idl:"name:DomainNameParam;string" json:"domain_name_param"`
	// AccountName: A pointer to a string that specifies the account name to use when connecting
	// to a domain controller. This parameter is optional. If this parameter is NULL, the
	// caller's account name MUST be used.
	AccountName string `idl:"name:AccountName;string;pointer:unique" json:"account_name"`
	// Password: An optional pointer to a JOINPR_ENCRYPTED_USER_PASSWORD structure (section
	// 2.2.5.18) that specifies the encrypted password to use with the AccountName parameter.
	// If the AccountName parameter is NULL, the caller's security context MUST be used,
	// and this parameter MUST be ignored.
	Password *EncryptedUserPassword `idl:"name:Password;pointer:unique" json:"password"`
	// OUCount: A pointer to the count of OUs that the method returns. The server MUST ignore
	// this parameter on input.
	OUCount uint32 `idl:"name:OUCount" json:"ou_count"`
}

func (o *GetJoinableOUs2Request) xxx_ToOp(ctx context.Context) *xxx_GetJoinableOUs2Operation {
	if o == nil {
		return &xxx_GetJoinableOUs2Operation{}
	}
	return &xxx_GetJoinableOUs2Operation{
		ServerName:      o.ServerName,
		DomainNameParam: o.DomainNameParam,
		AccountName:     o.AccountName,
		Password:        o.Password,
		OUCount:         o.OUCount,
	}
}

func (o *GetJoinableOUs2Request) xxx_FromOp(ctx context.Context, op *xxx_GetJoinableOUs2Operation) {
	if o == nil {
		return
	}
	o.ServerName = op.ServerName
	o.DomainNameParam = op.DomainNameParam
	o.AccountName = op.AccountName
	o.Password = op.Password
	o.OUCount = op.OUCount
}
func (o *GetJoinableOUs2Request) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetJoinableOUs2Request) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetJoinableOUs2Operation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetJoinableOUs2Response structure represents the NetrGetJoinableOUs2 operation response
type GetJoinableOUs2Response struct {
	// OUCount: A pointer to the count of OUs that the method returns. The server MUST ignore
	// this parameter on input.
	OUCount uint32 `idl:"name:OUCount" json:"ou_count"`
	// OUs: A pointer to a pointer of size OUCount to a block of strings that are the joinable
	// OUs that the method returns.
	OUs []string `idl:"name:OUs;size_is:(, OUCount);string" json:"ous"`
	// Return: The NetrGetJoinableOUs2 return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *GetJoinableOUs2Response) xxx_ToOp(ctx context.Context) *xxx_GetJoinableOUs2Operation {
	if o == nil {
		return &xxx_GetJoinableOUs2Operation{}
	}
	return &xxx_GetJoinableOUs2Operation{
		OUCount: o.OUCount,
		OUs:     o.OUs,
		Return:  o.Return,
	}
}

func (o *GetJoinableOUs2Response) xxx_FromOp(ctx context.Context, op *xxx_GetJoinableOUs2Operation) {
	if o == nil {
		return
	}
	o.OUCount = op.OUCount
	o.OUs = op.OUs
	o.Return = op.Return
}
func (o *GetJoinableOUs2Response) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetJoinableOUs2Response) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetJoinableOUs2Operation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_AddAlternateComputerNameOperation structure represents the NetrAddAlternateComputerName operation
type xxx_AddAlternateComputerNameOperation struct {
	ServerName        string                 `idl:"name:ServerName;string;pointer:unique" json:"server_name"`
	AlternateName     string                 `idl:"name:AlternateName;string;pointer:unique" json:"alternate_name"`
	DomainAccount     string                 `idl:"name:DomainAccount;string;pointer:unique" json:"domain_account"`
	EncryptedPassword *EncryptedUserPassword `idl:"name:EncryptedPassword;pointer:unique" json:"encrypted_password"`
	_                 uint32                 `idl:"name:Reserved"`
	Return            uint32                 `idl:"name:Return" json:"return"`
}

func (o *xxx_AddAlternateComputerNameOperation) OpNum() int { return 27 }

func (o *xxx_AddAlternateComputerNameOperation) OpName() string {
	return "/wkssvc/v1/NetrAddAlternateComputerName"
}

func (o *xxx_AddAlternateComputerNameOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AddAlternateComputerNameOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// ServerName {in} (1:{string, pointer=unique}*(1)[dim:0,string,null](wchar))
	{
		if o.ServerName != "" {
			_ptr_ServerName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.ServerName); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.ServerName, _ptr_ServerName); err != nil {
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
	// AlternateName {in} (1:{string, pointer=unique}*(1)[dim:0,string,null](wchar))
	{
		if o.AlternateName != "" {
			_ptr_AlternateName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.AlternateName); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.AlternateName, _ptr_AlternateName); err != nil {
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
	// DomainAccount {in} (1:{string, pointer=unique}*(1)[dim:0,string,null](wchar))
	{
		if o.DomainAccount != "" {
			_ptr_DomainAccount := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.DomainAccount); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.DomainAccount, _ptr_DomainAccount); err != nil {
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
	// EncryptedPassword {in} (1:{pointer=unique, alias=PJOINPR_ENCRYPTED_USER_PASSWORD}*(1))(2:{alias=JOINPR_ENCRYPTED_USER_PASSWORD}(struct))
	{
		if o.EncryptedPassword != nil {
			_ptr_EncryptedPassword := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.EncryptedPassword != nil {
					if err := o.EncryptedPassword.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&EncryptedUserPassword{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.EncryptedPassword, _ptr_EncryptedPassword); err != nil {
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
	// Reserved {in} (1:(uint32))
	{
		// reserved Reserved
		if err := w.WriteData(uint32(0)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AddAlternateComputerNameOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// ServerName {in} (1:{string, pointer=unique}*(1)[dim:0,string,null](wchar))
	{
		_ptr_ServerName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.ServerName); err != nil {
				return err
			}
			return nil
		})
		_s_ServerName := func(ptr interface{}) { o.ServerName = *ptr.(*string) }
		if err := w.ReadPointer(&o.ServerName, _s_ServerName, _ptr_ServerName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// AlternateName {in} (1:{string, pointer=unique}*(1)[dim:0,string,null](wchar))
	{
		_ptr_AlternateName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.AlternateName); err != nil {
				return err
			}
			return nil
		})
		_s_AlternateName := func(ptr interface{}) { o.AlternateName = *ptr.(*string) }
		if err := w.ReadPointer(&o.AlternateName, _s_AlternateName, _ptr_AlternateName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// DomainAccount {in} (1:{string, pointer=unique}*(1)[dim:0,string,null](wchar))
	{
		_ptr_DomainAccount := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.DomainAccount); err != nil {
				return err
			}
			return nil
		})
		_s_DomainAccount := func(ptr interface{}) { o.DomainAccount = *ptr.(*string) }
		if err := w.ReadPointer(&o.DomainAccount, _s_DomainAccount, _ptr_DomainAccount); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// EncryptedPassword {in} (1:{pointer=unique, alias=PJOINPR_ENCRYPTED_USER_PASSWORD}*(1))(2:{alias=JOINPR_ENCRYPTED_USER_PASSWORD}(struct))
	{
		_ptr_EncryptedPassword := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.EncryptedPassword == nil {
				o.EncryptedPassword = &EncryptedUserPassword{}
			}
			if err := o.EncryptedPassword.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_EncryptedPassword := func(ptr interface{}) { o.EncryptedPassword = *ptr.(**EncryptedUserPassword) }
		if err := w.ReadPointer(&o.EncryptedPassword, _s_EncryptedPassword, _ptr_EncryptedPassword); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Reserved {in} (1:(uint32))
	{
		// reserved Reserved
		var _Reserved uint32
		if err := w.ReadData(&_Reserved); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AddAlternateComputerNameOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AddAlternateComputerNameOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:(uint32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AddAlternateComputerNameOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// AddAlternateComputerNameRequest structure represents the NetrAddAlternateComputerName operation request
type AddAlternateComputerNameRequest struct {
	// ServerName: This parameter has no effect on message processing in any environment.
	// The client MUST set this parameter to a value that resolves to the IP protocol layer
	// destination address of the RPC packets it transmits ([MS-RPCE] section 2.1.1.2).
	// The server MUST ignore this parameter.
	ServerName string `idl:"name:ServerName;string;pointer:unique" json:"server_name"`
	// AlternateName: A pointer to a string that specifies the new alternate name to add.
	// The name MUST be a valid DNS host name [RFC1035].
	AlternateName string `idl:"name:AlternateName;string;pointer:unique" json:"alternate_name"`
	// DomainAccount: A pointer to a string that specifies the account name in the domain
	// to use when connecting to a domain controller. This parameter is optional. If this
	// parameter is NULL, the caller's account name MUST be used. If this parameter is specified,
	// the format MUST be one of the following:
	//
	// * <NetBIOSDomainName>\<UserName>
	//
	// * <FullyQualifiedDNSDomainName>\<UserName>
	//
	// * <UserName>@<FullyQualifiedDNSDomainName>
	DomainAccount string `idl:"name:DomainAccount;string;pointer:unique" json:"domain_account"`
	// EncryptedPassword: An optional pointer to a JOINPR_ENCRYPTED_USER_PASSWORD structure
	// (section 2.2.5.18) that specifies the encrypted password to use with the DomainAccount
	// parameter. If the DomainAccount parameter is NULL, the caller's security context
	// MUST be used, and this parameter MUST be ignored.
	EncryptedPassword *EncryptedUserPassword `idl:"name:EncryptedPassword;pointer:unique" json:"encrypted_password"`
}

func (o *AddAlternateComputerNameRequest) xxx_ToOp(ctx context.Context) *xxx_AddAlternateComputerNameOperation {
	if o == nil {
		return &xxx_AddAlternateComputerNameOperation{}
	}
	return &xxx_AddAlternateComputerNameOperation{
		ServerName:        o.ServerName,
		AlternateName:     o.AlternateName,
		DomainAccount:     o.DomainAccount,
		EncryptedPassword: o.EncryptedPassword,
	}
}

func (o *AddAlternateComputerNameRequest) xxx_FromOp(ctx context.Context, op *xxx_AddAlternateComputerNameOperation) {
	if o == nil {
		return
	}
	o.ServerName = op.ServerName
	o.AlternateName = op.AlternateName
	o.DomainAccount = op.DomainAccount
	o.EncryptedPassword = op.EncryptedPassword
}
func (o *AddAlternateComputerNameRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *AddAlternateComputerNameRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_AddAlternateComputerNameOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// AddAlternateComputerNameResponse structure represents the NetrAddAlternateComputerName operation response
type AddAlternateComputerNameResponse struct {
	// Return: The NetrAddAlternateComputerName return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *AddAlternateComputerNameResponse) xxx_ToOp(ctx context.Context) *xxx_AddAlternateComputerNameOperation {
	if o == nil {
		return &xxx_AddAlternateComputerNameOperation{}
	}
	return &xxx_AddAlternateComputerNameOperation{
		Return: o.Return,
	}
}

func (o *AddAlternateComputerNameResponse) xxx_FromOp(ctx context.Context, op *xxx_AddAlternateComputerNameOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *AddAlternateComputerNameResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *AddAlternateComputerNameResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_AddAlternateComputerNameOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_RemoveAlternateComputerNameOperation structure represents the NetrRemoveAlternateComputerName operation
type xxx_RemoveAlternateComputerNameOperation struct {
	ServerName        string                 `idl:"name:ServerName;string;pointer:unique" json:"server_name"`
	AlternateName     string                 `idl:"name:AlternateName;string;pointer:unique" json:"alternate_name"`
	DomainAccount     string                 `idl:"name:DomainAccount;string;pointer:unique" json:"domain_account"`
	EncryptedPassword *EncryptedUserPassword `idl:"name:EncryptedPassword;pointer:unique" json:"encrypted_password"`
	_                 uint32                 `idl:"name:Reserved"`
	Return            uint32                 `idl:"name:Return" json:"return"`
}

func (o *xxx_RemoveAlternateComputerNameOperation) OpNum() int { return 28 }

func (o *xxx_RemoveAlternateComputerNameOperation) OpName() string {
	return "/wkssvc/v1/NetrRemoveAlternateComputerName"
}

func (o *xxx_RemoveAlternateComputerNameOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RemoveAlternateComputerNameOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// ServerName {in} (1:{string, pointer=unique}*(1)[dim:0,string,null](wchar))
	{
		if o.ServerName != "" {
			_ptr_ServerName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.ServerName); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.ServerName, _ptr_ServerName); err != nil {
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
	// AlternateName {in} (1:{string, pointer=unique}*(1)[dim:0,string,null](wchar))
	{
		if o.AlternateName != "" {
			_ptr_AlternateName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.AlternateName); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.AlternateName, _ptr_AlternateName); err != nil {
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
	// DomainAccount {in} (1:{string, pointer=unique}*(1)[dim:0,string,null](wchar))
	{
		if o.DomainAccount != "" {
			_ptr_DomainAccount := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.DomainAccount); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.DomainAccount, _ptr_DomainAccount); err != nil {
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
	// EncryptedPassword {in} (1:{pointer=unique, alias=PJOINPR_ENCRYPTED_USER_PASSWORD}*(1))(2:{alias=JOINPR_ENCRYPTED_USER_PASSWORD}(struct))
	{
		if o.EncryptedPassword != nil {
			_ptr_EncryptedPassword := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.EncryptedPassword != nil {
					if err := o.EncryptedPassword.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&EncryptedUserPassword{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.EncryptedPassword, _ptr_EncryptedPassword); err != nil {
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
	// Reserved {in} (1:(uint32))
	{
		// reserved Reserved
		if err := w.WriteData(uint32(0)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RemoveAlternateComputerNameOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// ServerName {in} (1:{string, pointer=unique}*(1)[dim:0,string,null](wchar))
	{
		_ptr_ServerName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.ServerName); err != nil {
				return err
			}
			return nil
		})
		_s_ServerName := func(ptr interface{}) { o.ServerName = *ptr.(*string) }
		if err := w.ReadPointer(&o.ServerName, _s_ServerName, _ptr_ServerName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// AlternateName {in} (1:{string, pointer=unique}*(1)[dim:0,string,null](wchar))
	{
		_ptr_AlternateName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.AlternateName); err != nil {
				return err
			}
			return nil
		})
		_s_AlternateName := func(ptr interface{}) { o.AlternateName = *ptr.(*string) }
		if err := w.ReadPointer(&o.AlternateName, _s_AlternateName, _ptr_AlternateName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// DomainAccount {in} (1:{string, pointer=unique}*(1)[dim:0,string,null](wchar))
	{
		_ptr_DomainAccount := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.DomainAccount); err != nil {
				return err
			}
			return nil
		})
		_s_DomainAccount := func(ptr interface{}) { o.DomainAccount = *ptr.(*string) }
		if err := w.ReadPointer(&o.DomainAccount, _s_DomainAccount, _ptr_DomainAccount); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// EncryptedPassword {in} (1:{pointer=unique, alias=PJOINPR_ENCRYPTED_USER_PASSWORD}*(1))(2:{alias=JOINPR_ENCRYPTED_USER_PASSWORD}(struct))
	{
		_ptr_EncryptedPassword := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.EncryptedPassword == nil {
				o.EncryptedPassword = &EncryptedUserPassword{}
			}
			if err := o.EncryptedPassword.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_EncryptedPassword := func(ptr interface{}) { o.EncryptedPassword = *ptr.(**EncryptedUserPassword) }
		if err := w.ReadPointer(&o.EncryptedPassword, _s_EncryptedPassword, _ptr_EncryptedPassword); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Reserved {in} (1:(uint32))
	{
		// reserved Reserved
		var _Reserved uint32
		if err := w.ReadData(&_Reserved); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RemoveAlternateComputerNameOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RemoveAlternateComputerNameOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:(uint32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RemoveAlternateComputerNameOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// RemoveAlternateComputerNameRequest structure represents the NetrRemoveAlternateComputerName operation request
type RemoveAlternateComputerNameRequest struct {
	// ServerName: This parameter has no effect on message processing in any environment.
	// The client MUST set this parameter to a value that resolves to the IP protocol layer
	// destination address of the RPC packets it transmits ([MS-RPCE] section 2.1.1.2).
	// The server MUST ignore this parameter.
	ServerName string `idl:"name:ServerName;string;pointer:unique" json:"server_name"`
	// AlternateName: A pointer to a string that specifies the alternate name to remove.
	// The name MUST be a valid DNS host name [RFC1035].
	AlternateName string `idl:"name:AlternateName;string;pointer:unique" json:"alternate_name"`
	// DomainAccount: A pointer to a string that specifies the account name in the domain
	// to use when connecting to a domain controller. This parameter is optional. If this
	// parameter is NULL, the caller's account name MUST be used. If this parameter is specified,
	// the format MUST be one of the following:
	//
	// * <NetBIOSDomainName>\<UserName>
	//
	// * <FullyQualifiedDNSDomainName>\<UserName>
	//
	// * <UserName>@<FullyQualifiedDNSDomainName>
	DomainAccount string `idl:"name:DomainAccount;string;pointer:unique" json:"domain_account"`
	// EncryptedPassword: An optional pointer to a JOINPR_ENCRYPTED_USER_PASSWORD structure
	// (section 2.2.5.18) that specifies the encrypted password to use with the DomainAccount
	// parameter. If the DomainAccount parameter is NULL, the caller's security context
	// MUST be used, and this parameter MUST be ignored.
	EncryptedPassword *EncryptedUserPassword `idl:"name:EncryptedPassword;pointer:unique" json:"encrypted_password"`
}

func (o *RemoveAlternateComputerNameRequest) xxx_ToOp(ctx context.Context) *xxx_RemoveAlternateComputerNameOperation {
	if o == nil {
		return &xxx_RemoveAlternateComputerNameOperation{}
	}
	return &xxx_RemoveAlternateComputerNameOperation{
		ServerName:        o.ServerName,
		AlternateName:     o.AlternateName,
		DomainAccount:     o.DomainAccount,
		EncryptedPassword: o.EncryptedPassword,
	}
}

func (o *RemoveAlternateComputerNameRequest) xxx_FromOp(ctx context.Context, op *xxx_RemoveAlternateComputerNameOperation) {
	if o == nil {
		return
	}
	o.ServerName = op.ServerName
	o.AlternateName = op.AlternateName
	o.DomainAccount = op.DomainAccount
	o.EncryptedPassword = op.EncryptedPassword
}
func (o *RemoveAlternateComputerNameRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *RemoveAlternateComputerNameRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RemoveAlternateComputerNameOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RemoveAlternateComputerNameResponse structure represents the NetrRemoveAlternateComputerName operation response
type RemoveAlternateComputerNameResponse struct {
	// Return: The NetrRemoveAlternateComputerName return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *RemoveAlternateComputerNameResponse) xxx_ToOp(ctx context.Context) *xxx_RemoveAlternateComputerNameOperation {
	if o == nil {
		return &xxx_RemoveAlternateComputerNameOperation{}
	}
	return &xxx_RemoveAlternateComputerNameOperation{
		Return: o.Return,
	}
}

func (o *RemoveAlternateComputerNameResponse) xxx_FromOp(ctx context.Context, op *xxx_RemoveAlternateComputerNameOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *RemoveAlternateComputerNameResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *RemoveAlternateComputerNameResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RemoveAlternateComputerNameOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetPrimaryComputerNameOperation structure represents the NetrSetPrimaryComputerName operation
type xxx_SetPrimaryComputerNameOperation struct {
	ServerName        string                 `idl:"name:ServerName;string;pointer:unique" json:"server_name"`
	PrimaryName       string                 `idl:"name:PrimaryName;string;pointer:unique" json:"primary_name"`
	DomainAccount     string                 `idl:"name:DomainAccount;string;pointer:unique" json:"domain_account"`
	EncryptedPassword *EncryptedUserPassword `idl:"name:EncryptedPassword;pointer:unique" json:"encrypted_password"`
	_                 uint32                 `idl:"name:Reserved"`
	Return            uint32                 `idl:"name:Return" json:"return"`
}

func (o *xxx_SetPrimaryComputerNameOperation) OpNum() int { return 29 }

func (o *xxx_SetPrimaryComputerNameOperation) OpName() string {
	return "/wkssvc/v1/NetrSetPrimaryComputerName"
}

func (o *xxx_SetPrimaryComputerNameOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetPrimaryComputerNameOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// ServerName {in} (1:{string, pointer=unique}*(1)[dim:0,string,null](wchar))
	{
		if o.ServerName != "" {
			_ptr_ServerName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.ServerName); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.ServerName, _ptr_ServerName); err != nil {
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
	// PrimaryName {in} (1:{string, pointer=unique}*(1)[dim:0,string,null](wchar))
	{
		if o.PrimaryName != "" {
			_ptr_PrimaryName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.PrimaryName); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.PrimaryName, _ptr_PrimaryName); err != nil {
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
	// DomainAccount {in} (1:{string, pointer=unique}*(1)[dim:0,string,null](wchar))
	{
		if o.DomainAccount != "" {
			_ptr_DomainAccount := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.DomainAccount); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.DomainAccount, _ptr_DomainAccount); err != nil {
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
	// EncryptedPassword {in} (1:{pointer=unique, alias=PJOINPR_ENCRYPTED_USER_PASSWORD}*(1))(2:{alias=JOINPR_ENCRYPTED_USER_PASSWORD}(struct))
	{
		if o.EncryptedPassword != nil {
			_ptr_EncryptedPassword := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.EncryptedPassword != nil {
					if err := o.EncryptedPassword.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&EncryptedUserPassword{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.EncryptedPassword, _ptr_EncryptedPassword); err != nil {
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
	// Reserved {in} (1:(uint32))
	{
		// reserved Reserved
		if err := w.WriteData(uint32(0)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetPrimaryComputerNameOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// ServerName {in} (1:{string, pointer=unique}*(1)[dim:0,string,null](wchar))
	{
		_ptr_ServerName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.ServerName); err != nil {
				return err
			}
			return nil
		})
		_s_ServerName := func(ptr interface{}) { o.ServerName = *ptr.(*string) }
		if err := w.ReadPointer(&o.ServerName, _s_ServerName, _ptr_ServerName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// PrimaryName {in} (1:{string, pointer=unique}*(1)[dim:0,string,null](wchar))
	{
		_ptr_PrimaryName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.PrimaryName); err != nil {
				return err
			}
			return nil
		})
		_s_PrimaryName := func(ptr interface{}) { o.PrimaryName = *ptr.(*string) }
		if err := w.ReadPointer(&o.PrimaryName, _s_PrimaryName, _ptr_PrimaryName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// DomainAccount {in} (1:{string, pointer=unique}*(1)[dim:0,string,null](wchar))
	{
		_ptr_DomainAccount := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.DomainAccount); err != nil {
				return err
			}
			return nil
		})
		_s_DomainAccount := func(ptr interface{}) { o.DomainAccount = *ptr.(*string) }
		if err := w.ReadPointer(&o.DomainAccount, _s_DomainAccount, _ptr_DomainAccount); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// EncryptedPassword {in} (1:{pointer=unique, alias=PJOINPR_ENCRYPTED_USER_PASSWORD}*(1))(2:{alias=JOINPR_ENCRYPTED_USER_PASSWORD}(struct))
	{
		_ptr_EncryptedPassword := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.EncryptedPassword == nil {
				o.EncryptedPassword = &EncryptedUserPassword{}
			}
			if err := o.EncryptedPassword.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_EncryptedPassword := func(ptr interface{}) { o.EncryptedPassword = *ptr.(**EncryptedUserPassword) }
		if err := w.ReadPointer(&o.EncryptedPassword, _s_EncryptedPassword, _ptr_EncryptedPassword); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Reserved {in} (1:(uint32))
	{
		// reserved Reserved
		var _Reserved uint32
		if err := w.ReadData(&_Reserved); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetPrimaryComputerNameOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetPrimaryComputerNameOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:(uint32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetPrimaryComputerNameOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// SetPrimaryComputerNameRequest structure represents the NetrSetPrimaryComputerName operation request
type SetPrimaryComputerNameRequest struct {
	// ServerName: This parameter has no effect on message processing in any environment.
	// The client MUST set this parameter to a value that resolves to the IP protocol layer
	// destination address of the RPC packets it transmits ([MS-RPCE] section 2.1.1.2).
	// The server MUST ignore this parameter.
	ServerName string `idl:"name:ServerName;string;pointer:unique" json:"server_name"`
	// PrimaryName: A pointer to a string that specifies the primary computer name to set.
	// The name MUST be a valid DNS host name [RFC1035].
	PrimaryName string `idl:"name:PrimaryName;string;pointer:unique" json:"primary_name"`
	// DomainAccount: A pointer to a string that specifies the account name in the joined
	// domain to use when connecting to a domain controller. This parameter is optional.
	// If this parameter is NULL, the caller's account name MUST be used. This parameter
	// is not used if the server is not joined to a domain.
	//
	// * <NetBIOSDomainName>\<UserName>
	//
	// * <FullyQualifiedDNSDomainName>\<UserName>
	//
	// * <UserName>@<FullyQualifiedDNSDomainName>
	DomainAccount string `idl:"name:DomainAccount;string;pointer:unique" json:"domain_account"`
	// EncryptedPassword: An optional pointer to a JOINPR_ENCRYPTED_USER_PASSWORD structure
	// (section 2.2.5.18) that specifies the encrypted password to use with the DomainAccount
	// parameter. If the DomainAccount parameter is NULL, the caller's security context
	// MUST be used, and this parameter MUST be ignored.
	EncryptedPassword *EncryptedUserPassword `idl:"name:EncryptedPassword;pointer:unique" json:"encrypted_password"`
}

func (o *SetPrimaryComputerNameRequest) xxx_ToOp(ctx context.Context) *xxx_SetPrimaryComputerNameOperation {
	if o == nil {
		return &xxx_SetPrimaryComputerNameOperation{}
	}
	return &xxx_SetPrimaryComputerNameOperation{
		ServerName:        o.ServerName,
		PrimaryName:       o.PrimaryName,
		DomainAccount:     o.DomainAccount,
		EncryptedPassword: o.EncryptedPassword,
	}
}

func (o *SetPrimaryComputerNameRequest) xxx_FromOp(ctx context.Context, op *xxx_SetPrimaryComputerNameOperation) {
	if o == nil {
		return
	}
	o.ServerName = op.ServerName
	o.PrimaryName = op.PrimaryName
	o.DomainAccount = op.DomainAccount
	o.EncryptedPassword = op.EncryptedPassword
}
func (o *SetPrimaryComputerNameRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *SetPrimaryComputerNameRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetPrimaryComputerNameOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetPrimaryComputerNameResponse structure represents the NetrSetPrimaryComputerName operation response
type SetPrimaryComputerNameResponse struct {
	// Return: The NetrSetPrimaryComputerName return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *SetPrimaryComputerNameResponse) xxx_ToOp(ctx context.Context) *xxx_SetPrimaryComputerNameOperation {
	if o == nil {
		return &xxx_SetPrimaryComputerNameOperation{}
	}
	return &xxx_SetPrimaryComputerNameOperation{
		Return: o.Return,
	}
}

func (o *SetPrimaryComputerNameResponse) xxx_FromOp(ctx context.Context, op *xxx_SetPrimaryComputerNameOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *SetPrimaryComputerNameResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *SetPrimaryComputerNameResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetPrimaryComputerNameOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_EnumerateComputerNamesOperation structure represents the NetrEnumerateComputerNames operation
type xxx_EnumerateComputerNamesOperation struct {
	ServerName    string                `idl:"name:ServerName;string;pointer:unique" json:"server_name"`
	NameType      NetComputerNameType   `idl:"name:NameType" json:"name_type"`
	_             uint32                `idl:"name:Reserved"`
	ComputerNames *NetComputerNameArray `idl:"name:ComputerNames" json:"computer_names"`
	Return        uint32                `idl:"name:Return" json:"return"`
}

func (o *xxx_EnumerateComputerNamesOperation) OpNum() int { return 30 }

func (o *xxx_EnumerateComputerNamesOperation) OpName() string {
	return "/wkssvc/v1/NetrEnumerateComputerNames"
}

func (o *xxx_EnumerateComputerNamesOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumerateComputerNamesOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// ServerName {in} (1:{handle, string, pointer=unique, alias=WKSSVC_IMPERSONATE_HANDLE}*(1)[dim:0,string,null](wchar))
	{
		if o.ServerName != "" {
			_ptr_ServerName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.ServerName); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.ServerName, _ptr_ServerName); err != nil {
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
	// NameType {in} (1:{alias=NET_COMPUTER_NAME_TYPE}(enum))
	{
		if err := w.WriteData(uint16(o.NameType)); err != nil {
			return err
		}
	}
	// Reserved {in} (1:(uint32))
	{
		// reserved Reserved
		if err := w.WriteData(uint32(0)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumerateComputerNamesOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// ServerName {in} (1:{handle, string, pointer=unique, alias=WKSSVC_IMPERSONATE_HANDLE}*(1)[dim:0,string,null](wchar))
	{
		_ptr_ServerName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.ServerName); err != nil {
				return err
			}
			return nil
		})
		_s_ServerName := func(ptr interface{}) { o.ServerName = *ptr.(*string) }
		if err := w.ReadPointer(&o.ServerName, _s_ServerName, _ptr_ServerName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// NameType {in} (1:{alias=NET_COMPUTER_NAME_TYPE}(enum))
	{
		if err := w.ReadData((*uint16)(&o.NameType)); err != nil {
			return err
		}
	}
	// Reserved {in} (1:(uint32))
	{
		// reserved Reserved
		var _Reserved uint32
		if err := w.ReadData(&_Reserved); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumerateComputerNamesOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumerateComputerNamesOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// ComputerNames {out} (1:{pointer=ref}*(2))(2:{alias=PNET_COMPUTER_NAME_ARRAY}*(1))(3:{alias=NET_COMPUTER_NAME_ARRAY}(struct))
	{
		if o.ComputerNames != nil {
			_ptr_ComputerNames := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ComputerNames != nil {
					if err := o.ComputerNames.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&NetComputerNameArray{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ComputerNames, _ptr_ComputerNames); err != nil {
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
	// Return {out} (1:(uint32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumerateComputerNamesOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// ComputerNames {out} (1:{pointer=ref}*(2))(2:{alias=PNET_COMPUTER_NAME_ARRAY,pointer=ref}*(1))(3:{alias=NET_COMPUTER_NAME_ARRAY}(struct))
	{
		_ptr_ComputerNames := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ComputerNames == nil {
				o.ComputerNames = &NetComputerNameArray{}
			}
			if err := o.ComputerNames.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ComputerNames := func(ptr interface{}) { o.ComputerNames = *ptr.(**NetComputerNameArray) }
		if err := w.ReadPointer(&o.ComputerNames, _s_ComputerNames, _ptr_ComputerNames); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// EnumerateComputerNamesRequest structure represents the NetrEnumerateComputerNames operation request
type EnumerateComputerNamesRequest struct {
	// ServerName: A WKSSVC_IMPERSONATE_HANDLE structure (section 2.2.2.2) that specifies
	// the server. The client MUST map this structure to an RPC binding handle ([C706] sections
	// 4.3.5 and 5.1.5.2). The server MUST ignore this parameter.
	ServerName string `idl:"name:ServerName;string;pointer:unique" json:"server_name"`
	// NameType: The type of query issued. See NET_COMPUTER_NAME_TYPE (section 2.2.3.3).
	NameType NetComputerNameType `idl:"name:NameType" json:"name_type"`
}

func (o *EnumerateComputerNamesRequest) xxx_ToOp(ctx context.Context) *xxx_EnumerateComputerNamesOperation {
	if o == nil {
		return &xxx_EnumerateComputerNamesOperation{}
	}
	return &xxx_EnumerateComputerNamesOperation{
		ServerName: o.ServerName,
		NameType:   o.NameType,
	}
}

func (o *EnumerateComputerNamesRequest) xxx_FromOp(ctx context.Context, op *xxx_EnumerateComputerNamesOperation) {
	if o == nil {
		return
	}
	o.ServerName = op.ServerName
	o.NameType = op.NameType
}
func (o *EnumerateComputerNamesRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *EnumerateComputerNamesRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EnumerateComputerNamesOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// EnumerateComputerNamesResponse structure represents the NetrEnumerateComputerNames operation response
type EnumerateComputerNamesResponse struct {
	// ComputerNames: A pointer to structure containing a list of computer name strings.
	// See NET_COMPUTER_NAME_ARRAY (section 2.2.5.21).
	ComputerNames *NetComputerNameArray `idl:"name:ComputerNames" json:"computer_names"`
	// Return: The NetrEnumerateComputerNames return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *EnumerateComputerNamesResponse) xxx_ToOp(ctx context.Context) *xxx_EnumerateComputerNamesOperation {
	if o == nil {
		return &xxx_EnumerateComputerNamesOperation{}
	}
	return &xxx_EnumerateComputerNamesOperation{
		ComputerNames: o.ComputerNames,
		Return:        o.Return,
	}
}

func (o *EnumerateComputerNamesResponse) xxx_FromOp(ctx context.Context, op *xxx_EnumerateComputerNamesOperation) {
	if o == nil {
		return
	}
	o.ComputerNames = op.ComputerNames
	o.Return = op.Return
}
func (o *EnumerateComputerNamesResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *EnumerateComputerNamesResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EnumerateComputerNamesOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
