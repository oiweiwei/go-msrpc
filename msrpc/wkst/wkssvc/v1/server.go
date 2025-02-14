package wkssvc

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

// wkssvc server interface.
type WkssvcServer interface {

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
	GetInfo(context.Context, *GetInfoRequest) (*GetInfoResponse, error)

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
	SetInfo(context.Context, *SetInfoRequest) (*SetInfoResponse, error)

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
	UserEnum(context.Context, *UserEnumRequest) (*UserEnumResponse, error)

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
	TransportEnum(context.Context, *TransportEnumRequest) (*TransportEnumResponse, error)

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
	TransportAdd(context.Context, *TransportAddRequest) (*TransportAddResponse, error)

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
	TransportDelete(context.Context, *TransportDeleteRequest) (*TransportDeleteResponse, error)

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
	UseAdd(context.Context, *UseAddRequest) (*UseAddResponse, error)

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
	UseGetInfo(context.Context, *UseGetInfoRequest) (*UseGetInfoResponse, error)

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
	UseDelete(context.Context, *UseDeleteRequest) (*UseDeleteResponse, error)

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
	UseEnum(context.Context, *UseEnumRequest) (*UseEnumResponse, error)

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
	WorkstationStatisticsGet(context.Context, *WorkstationStatisticsGetRequest) (*WorkstationStatisticsGetResponse, error)

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
	GetJoinInformation(context.Context, *GetJoinInformationRequest) (*GetJoinInformationResponse, error)

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
	JoinDomain2(context.Context, *JoinDomain2Request) (*JoinDomain2Response, error)

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
	UnjoinDomain2(context.Context, *UnjoinDomain2Request) (*UnjoinDomain2Response, error)

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
	RenameMachineInDomain2(context.Context, *RenameMachineInDomain2Request) (*RenameMachineInDomain2Response, error)

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
	ValidateName2(context.Context, *ValidateName2Request) (*ValidateName2Response, error)

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
	GetJoinableOUs2(context.Context, *GetJoinableOUs2Request) (*GetJoinableOUs2Response, error)

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
	AddAlternateComputerName(context.Context, *AddAlternateComputerNameRequest) (*AddAlternateComputerNameResponse, error)

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
	RemoveAlternateComputerName(context.Context, *RemoveAlternateComputerNameRequest) (*RemoveAlternateComputerNameResponse, error)

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
	SetPrimaryComputerName(context.Context, *SetPrimaryComputerNameRequest) (*SetPrimaryComputerNameResponse, error)

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
	EnumerateComputerNames(context.Context, *EnumerateComputerNamesRequest) (*EnumerateComputerNamesResponse, error)
}

func RegisterWkssvcServer(conn dcerpc.Conn, o WkssvcServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewWkssvcServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(WkssvcSyntaxV1_0))...)
}

func NewWkssvcServerHandle(o WkssvcServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return WkssvcServerHandle(ctx, o, opNum, r)
	}
}

func WkssvcServerHandle(ctx context.Context, o WkssvcServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	switch opNum {
	case 0: // NetrWkstaGetInfo
		op := &xxx_GetInfoOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetInfoRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetInfo(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 1: // NetrWkstaSetInfo
		op := &xxx_SetInfoOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetInfoRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetInfo(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 2: // NetrWkstaUserEnum
		op := &xxx_UserEnumOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &UserEnumRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.UserEnum(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 3: // Opnum3NotUsedOnWire
		// Opnum3NotUsedOnWire
		return nil, nil
	case 4: // Opnum4NotUsedOnWire
		// Opnum4NotUsedOnWire
		return nil, nil
	case 5: // NetrWkstaTransportEnum
		op := &xxx_TransportEnumOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &TransportEnumRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.TransportEnum(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 6: // NetrWkstaTransportAdd
		op := &xxx_TransportAddOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &TransportAddRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.TransportAdd(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 7: // NetrWkstaTransportDel
		op := &xxx_TransportDeleteOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &TransportDeleteRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.TransportDelete(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 8: // NetrUseAdd
		op := &xxx_UseAddOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &UseAddRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.UseAdd(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 9: // NetrUseGetInfo
		op := &xxx_UseGetInfoOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &UseGetInfoRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.UseGetInfo(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 10: // NetrUseDel
		op := &xxx_UseDeleteOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &UseDeleteRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.UseDelete(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 11: // NetrUseEnum
		op := &xxx_UseEnumOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &UseEnumRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.UseEnum(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 12: // Opnum12NotUsedOnWire
		// Opnum12NotUsedOnWire
		return nil, nil
	case 13: // NetrWorkstationStatisticsGet
		op := &xxx_WorkstationStatisticsGetOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &WorkstationStatisticsGetRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.WorkstationStatisticsGet(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 14: // Opnum14NotUsedOnWire
		// Opnum14NotUsedOnWire
		return nil, nil
	case 15: // Opnum15NotUsedOnWire
		// Opnum15NotUsedOnWire
		return nil, nil
	case 16: // Opnum16NotUsedOnWire
		// Opnum16NotUsedOnWire
		return nil, nil
	case 17: // Opnum17NotUsedOnWire
		// Opnum17NotUsedOnWire
		return nil, nil
	case 18: // Opnum18NotUsedOnWire
		// Opnum18NotUsedOnWire
		return nil, nil
	case 19: // Opnum19NotUsedOnWire
		// Opnum19NotUsedOnWire
		return nil, nil
	case 20: // NetrGetJoinInformation
		op := &xxx_GetJoinInformationOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetJoinInformationRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetJoinInformation(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 21: // Opnum21NotUsedOnWire
		// Opnum21NotUsedOnWire
		return nil, nil
	case 22: // NetrJoinDomain2
		op := &xxx_JoinDomain2Operation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &JoinDomain2Request{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.JoinDomain2(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 23: // NetrUnjoinDomain2
		op := &xxx_UnjoinDomain2Operation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &UnjoinDomain2Request{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.UnjoinDomain2(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 24: // NetrRenameMachineInDomain2
		op := &xxx_RenameMachineInDomain2Operation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RenameMachineInDomain2Request{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RenameMachineInDomain2(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 25: // NetrValidateName2
		op := &xxx_ValidateName2Operation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ValidateName2Request{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.ValidateName2(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 26: // NetrGetJoinableOUs2
		op := &xxx_GetJoinableOUs2Operation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetJoinableOUs2Request{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetJoinableOUs2(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 27: // NetrAddAlternateComputerName
		op := &xxx_AddAlternateComputerNameOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &AddAlternateComputerNameRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.AddAlternateComputerName(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 28: // NetrRemoveAlternateComputerName
		op := &xxx_RemoveAlternateComputerNameOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RemoveAlternateComputerNameRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RemoveAlternateComputerName(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 29: // NetrSetPrimaryComputerName
		op := &xxx_SetPrimaryComputerNameOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetPrimaryComputerNameRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetPrimaryComputerName(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 30: // NetrEnumerateComputerNames
		op := &xxx_EnumerateComputerNamesOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &EnumerateComputerNamesRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.EnumerateComputerNames(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented wkssvc
type UnimplementedWkssvcServer struct {
}

func (UnimplementedWkssvcServer) GetInfo(context.Context, *GetInfoRequest) (*GetInfoResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedWkssvcServer) SetInfo(context.Context, *SetInfoRequest) (*SetInfoResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedWkssvcServer) UserEnum(context.Context, *UserEnumRequest) (*UserEnumResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedWkssvcServer) TransportEnum(context.Context, *TransportEnumRequest) (*TransportEnumResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedWkssvcServer) TransportAdd(context.Context, *TransportAddRequest) (*TransportAddResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedWkssvcServer) TransportDelete(context.Context, *TransportDeleteRequest) (*TransportDeleteResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedWkssvcServer) UseAdd(context.Context, *UseAddRequest) (*UseAddResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedWkssvcServer) UseGetInfo(context.Context, *UseGetInfoRequest) (*UseGetInfoResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedWkssvcServer) UseDelete(context.Context, *UseDeleteRequest) (*UseDeleteResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedWkssvcServer) UseEnum(context.Context, *UseEnumRequest) (*UseEnumResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedWkssvcServer) WorkstationStatisticsGet(context.Context, *WorkstationStatisticsGetRequest) (*WorkstationStatisticsGetResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedWkssvcServer) GetJoinInformation(context.Context, *GetJoinInformationRequest) (*GetJoinInformationResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedWkssvcServer) JoinDomain2(context.Context, *JoinDomain2Request) (*JoinDomain2Response, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedWkssvcServer) UnjoinDomain2(context.Context, *UnjoinDomain2Request) (*UnjoinDomain2Response, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedWkssvcServer) RenameMachineInDomain2(context.Context, *RenameMachineInDomain2Request) (*RenameMachineInDomain2Response, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedWkssvcServer) ValidateName2(context.Context, *ValidateName2Request) (*ValidateName2Response, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedWkssvcServer) GetJoinableOUs2(context.Context, *GetJoinableOUs2Request) (*GetJoinableOUs2Response, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedWkssvcServer) AddAlternateComputerName(context.Context, *AddAlternateComputerNameRequest) (*AddAlternateComputerNameResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedWkssvcServer) RemoveAlternateComputerName(context.Context, *RemoveAlternateComputerNameRequest) (*RemoveAlternateComputerNameResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedWkssvcServer) SetPrimaryComputerName(context.Context, *SetPrimaryComputerNameRequest) (*SetPrimaryComputerNameResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedWkssvcServer) EnumerateComputerNames(context.Context, *EnumerateComputerNamesRequest) (*EnumerateComputerNamesResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ WkssvcServer = (*UnimplementedWkssvcServer)(nil)
