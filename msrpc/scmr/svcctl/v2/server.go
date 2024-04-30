package svcctl

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

// svcctl server interface.
type SvcctlServer interface {

	// The RCloseServiceHandle method is called by the client. In response, the server releases
	// the handle to the specified service or the SCM database.
	//
	// Return Values: The method returns 0x00000000 (ERROR_SUCCESS) on success; otherwise,
	// it returns the following error code.
	//
	//	+------------------------+----------------------------------------------------------------------------------+
	//	|         RETURN         |                                                                                  |
	//	|       VALUE/CODE       |                                   DESCRIPTION                                    |
	//	|                        |                                                                                  |
	//	+------------------------+----------------------------------------------------------------------------------+
	//	+------------------------+----------------------------------------------------------------------------------+
	//	| 6 ERROR_INVALID_HANDLE | The handle is no longer valid.                                                   |
	//	+------------------------+----------------------------------------------------------------------------------+
	//	| 0xFFFF75FD             | The operation completed successfully. Additionally, the passed handle was the    |
	//	|                        | last one created for the associated service record that was previously used in a |
	//	|                        | successful call to the RNotifyServiceStatusChange (section 3.1.4.43) method.     |
	//	+------------------------+----------------------------------------------------------------------------------+
	//	| 0xFFFF75FE             | The operation completed successfully. Additionally, the passed handle was        |
	//	|                        | previously used in a successful call to the RNotifyServiceStatusChange method.   |
	//	+------------------------+----------------------------------------------------------------------------------+
	CloseService(context.Context, *CloseServiceRequest) (*CloseServiceResponse, error)

	// The RControlService method receives a control code for a specific service handle,
	// as specified by the client.
	//
	// Return Values: The method returns 0x00000000 (ERROR_SUCCESS) on success; otherwise,
	// it returns one of the following error codes.
	//
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	|                RETURN                 |                                                                                  |
	//	|              VALUE/CODE               |                                   DESCRIPTION                                    |
	//	|                                       |                                                                                  |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| 5 ERROR_ACCESS_DENIED                 | The required access right had not been granted to the caller when the RPC        |
	//	|                                       | context handle to the service record was created.                                |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| 1051 ERROR_DEPENDENT_SERVICES_RUNNING | The service cannot be stopped because other running services are dependent on    |
	//	|                                       | it.                                                                              |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| 6 ERROR_INVALID_HANDLE                | The handle is no longer valid.                                                   |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| 87 ERROR_INVALID_PARAMETER            | The requested control code is undefined                                          |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| 1052 ERROR_INVALID_SERVICE_CONTROL    | The requested control code is not valid, or it is unacceptable to the service.   |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| 1053 ERROR_SERVICE_REQUEST_TIMEOUT    | The process for the service was started, but it did not respond within an        |
	//	|                                       | implementation-specific time-out.<35>                                            |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| 1061 ERROR_SERVICE_CANNOT_ACCEPT_CTRL | The requested control code cannot be sent to the service because the             |
	//	|                                       | ServiceStatus.dwCurrentState in the service record is SERVICE_START_PENDING or   |
	//	|                                       | SERVICE_STOP_PENDING.                                                            |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| 1062 ERROR_SERVICE_NOT_ACTIVE         | The service has not been started, or the ServiceStatus.dwCurrentState in the     |
	//	|                                       | service record is SERVICE_STOPPED.                                               |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| 1115 ERROR_SHUTDOWN_IN_PROGRESS       | The system is shutting down.                                                     |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	ControlService(context.Context, *ControlServiceRequest) (*ControlServiceResponse, error)

	// The RDeleteService method marks the specified service for deletion from the SCM database.
	//
	// Return Values: The method returns 0x00000000 (ERROR_SUCCESS) on success; otherwise,
	// it returns one of the following error codes.
	//
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	|                RETURN                |                                                                                  |
	//	|              VALUE/CODE              |                                   DESCRIPTION                                    |
	//	|                                      |                                                                                  |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| 5 ERROR_ACCESS_DENIED                | The DELETE access right had not been granted to the caller when the RPC context  |
	//	|                                      | handle to the service record was created.                                        |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| 6 ERROR_INVALID_HANDLE               | The handle is no longer valid.                                                   |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| 1072 ERROR_SERVICE_MARKED_FOR_DELETE | The RDeleteService has already been called for the service record identified by  |
	//	|                                      | the hService parameter.                                                          |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| 1115 ERROR_SHUTDOWN_IN_PROGRESS      | The system is shutting down.                                                     |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	DeleteService(context.Context, *DeleteServiceRequest) (*DeleteServiceResponse, error)

	// The RLockServiceDatabase method acquires a lock on an SCM database.
	//
	// Return Values: The method returns 0x00000000 (ERROR_SUCCESS) on success; otherwise,
	// it returns one of the following error codes.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN               |                                                                                  |
	//	|             VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 5 ERROR_ACCESS_DENIED              | The SC_MANAGER_LOCK access rights had not been granted to the caller when the    |
	//	|                                    | RPC context handle was created.                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 6 ERROR_INVALID_HANDLE             | The handle is no longer valid.                                                   |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 1055 ERROR_SERVICE_DATABASE_LOCKED | The service database is locked.                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	LockServiceDatabase(context.Context, *LockServiceDatabaseRequest) (*LockServiceDatabaseResponse, error)

	// The RQueryServiceObjectSecurity method returns a copy of the SECURITY_DESCRIPTOR
	// structure associated with a service object.
	//
	// Return Values: The method returns 0x00000000 (ERROR_SUCCESS) on success; otherwise,
	// it returns one of the following error codes.
	//
	//	+-------------------------------+----------------------------------------------------------------------------------+
	//	|            RETURN             |                                                                                  |
	//	|          VALUE/CODE           |                                   DESCRIPTION                                    |
	//	|                               |                                                                                  |
	//	+-------------------------------+----------------------------------------------------------------------------------+
	//	+-------------------------------+----------------------------------------------------------------------------------+
	//	| 5 ERROR_ACCESS_DENIED         | The required access rights had not been granted to the caller when the RPC       |
	//	|                               | context handle was created.                                                      |
	//	+-------------------------------+----------------------------------------------------------------------------------+
	//	| 6 ERROR_INVALID_HANDLE        | The handle is no longer valid.                                                   |
	//	+-------------------------------+----------------------------------------------------------------------------------+
	//	| 87 ERROR_INVALID_PARAMETER    | A parameter that was specified is invalid.                                       |
	//	+-------------------------------+----------------------------------------------------------------------------------+
	//	| 122 ERROR_INSUFFICIENT_BUFFER | The data area passed to a system call is too small.                              |
	//	+-------------------------------+----------------------------------------------------------------------------------+
	//
	// The client MAY provide a combination of one or more SECURITY_INFORMATION bit flags
	// for dwSecurityInformation.
	//
	// If SACL_SECURITY_INFORMATION is specified for the dwSecurityInformation parameter,
	// then an ACCESS_SYSTEM_SECURITY right MUST have been granted to the caller when hService
	// was created. (See AS in ACCESS_MASK in [MS-DTYP] 2.4.3.)
	//
	// If DACL_SECURITY_INFORMATION, LABEL_SECURITY_INFORMATION, OWNER_SECURITY_INFORMATION,
	// or GROUP_SECURITY_INFORMATION is specified for the dwSecurityInformation parameter,
	// then a READ_CONTROL right MUST have been granted to the caller when hService was
	// created. (See RC in ACCESS_MASK in [MS-DTYP] 2.4.3.)
	QueryServiceObjectSecurity(context.Context, *QueryServiceObjectSecurityRequest) (*QueryServiceObjectSecurityResponse, error)

	// The RSetServiceObjectSecurity method sets the SECURITY_DESCRIPTOR structure associated
	// with a service object.
	//
	// Return Values: The method returns 0x00000000 (ERROR_SUCCESS) on success; otherwise,
	// it returns one of the following error codes.
	//
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	|                RETURN                |                                                                                  |
	//	|              VALUE/CODE              |                                   DESCRIPTION                                    |
	//	|                                      |                                                                                  |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| 5 ERROR_ACCESS_DENIED                | The required access rights had not been granted to the caller when the RPC       |
	//	|                                      | context handle was created.                                                      |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| 6 ERROR_INVALID_HANDLE               | The handle is no longer valid.                                                   |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| 87 ERROR_INVALID_PARAMETER           | A parameter that was specified is invalid.                                       |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| 1072 ERROR_SERVICE_MARKED_FOR_DELETE | The RDeleteService method has been called with an RPC context handle identifying |
	//	|                                      | the same service record as the hService parameter for this call.                 |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//
	// The client MAY provide a combination of one or more SECURITY_INFORMATION bit flags
	// for dwSecurityInformation.
	//
	// If SACL_SECURITY_INFORMATION is specified via dwSecurityInformation, then an ACCESS_SYSTEM_SECURITY
	// right MUSThave been granted to the caller when hService was created. (See WD in ACCESS_MASK
	// in [MS-DTYP] 2.4.3.
	//
	// If LABEL_SECURITY_INFORMATION or OWNER_SECURITY_INFORMATION or GROUP_SECURITY_INFORMATION
	// is specified via dwSecurityInformation, then a WRITE_OWNER right MUST have been granted
	// to the caller when hService was created. (See WO in ACCESS_MASK in [MS-DTYP] 2.4.3.)
	//
	// If DACL_SECURITY_INFORMATION is specified via dwSecurityInformation, then a WRITE_DAC
	// right MUST have been granted to the caller when hService was created. (See WD in
	// ACCESS_MASK in [MS-DTYP] 2.4.3.)
	SetServiceObjectSecurity(context.Context, *SetServiceObjectSecurityRequest) (*SetServiceObjectSecurityResponse, error)

	// The RQueryServiceStatus method returns the current status of the specified service.
	//
	// Return Values: The method returns 0x00000000 (ERROR_SUCCESS) on success; otherwise,
	// it returns one of the following error codes.
	//
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	|             RETURN              |                                                                                  |
	//	|           VALUE/CODE            |                                   DESCRIPTION                                    |
	//	|                                 |                                                                                  |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	| 5 ERROR_ACCESS_DENIED           | The SERVICE_QUERY_STATUS access right had not been granted to the caller when    |
	//	|                                 | the RPC context handle was created.                                              |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	| 6 ERROR_INVALID_HANDLE          | The handle is no longer valid.                                                   |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	| 3 ERROR_PATH_NOT_FOUND          | The ImagePath of the service record identified by the hService parameter does    |
	//	|                                 | not exist.                                                                       |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	| 1115 ERROR_SHUTDOWN_IN_PROGRESS | The system is shutting down.                                                     |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//
	// If the type of the service record is SERVICE_KERNEL_DRIVER or SERVICE_FILESYSTEM_DRIVER,
	// the SCM queries the current status of the driver from the operating system and set
	// the ServiceStatus.dwCurrentState of the service record to SERVICE_RUNNING if driver
	// is loaded and to SERVICE_STOPPED if it is not.
	QueryServiceStatus(context.Context, *QueryServiceStatusRequest) (*QueryServiceStatusResponse, error)

	// The RSetServiceStatus method updates the SCM status information for the calling service.
	//
	// Return Values: The method returns 0x00000000 (ERROR_SUCCESS) on success; otherwise,
	// it returns one of the following error codes.
	//
	//	+------------------------+----------------------------------------------------------------------------------+
	//	|         RETURN         |                                                                                  |
	//	|       VALUE/CODE       |                                   DESCRIPTION                                    |
	//	|                        |                                                                                  |
	//	+------------------------+----------------------------------------------------------------------------------+
	//	+------------------------+----------------------------------------------------------------------------------+
	//	| 6 ERROR_INVALID_HANDLE | Either the handle is no longer valid or the SERVICE_SET_STATUS access rights had |
	//	|                        | not been granted to the caller when the RPC context handle was created.          |
	//	+------------------------+----------------------------------------------------------------------------------+
	//	| 13 ERROR_INVALID_DATA  | The data provided in the lpServiceStatus parameter is invalid.                   |
	//	+------------------------+----------------------------------------------------------------------------------+
	SetServiceStatus(context.Context, *SetServiceStatusRequest) (*SetServiceStatusResponse, error)

	// The RUnlockServiceDatabase method releases a lock on a service database.
	//
	// Return Values: The method returns 0x00000000 (ERROR_SUCCESS) on success; otherwise,
	// it returns the following error code.
	//
	//	+---------------------------------+----------------------------------------------+
	//	|             RETURN              |                                              |
	//	|           VALUE/CODE            |                 DESCRIPTION                  |
	//	|                                 |                                              |
	//	+---------------------------------+----------------------------------------------+
	//	+---------------------------------+----------------------------------------------+
	//	| 1071 ERROR_INVALID_SERVICE_LOCK | The specified RPC context handle is invalid. |
	//	+---------------------------------+----------------------------------------------+
	UnlockServiceDatabase(context.Context, *UnlockServiceDatabaseRequest) (*UnlockServiceDatabaseResponse, error)

	// The RNotifyBootConfigStatus method reports the boot status to the SCM.
	//
	// Return Values: The method returns ERROR_SUCCESS (0x00000000) on success; otherwise,
	// it returns one of the following error codes.
	//
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	|              RETURN              |                                                                                  |
	//	|            VALUE/CODE            |                                   DESCRIPTION                                    |
	//	|                                  |                                                                                  |
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	| 5 ERROR_ACCESS_DENIED            | The caller does not have the SC_MANAGER_MODIFY_BOOT_CONFIG access rights granted |
	//	|                                  | in the SCM Security Descriptor.                                                  |
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	| 1074 ERROR_ALREADY_RUNNING_LKG   | The system is currently running with the last-known-good configuration.          |
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	| 1076 ERROR_BOOT_ALREADY_ACCEPTED | The BootAccepted field of the SCM on the target machine indicated that a         |
	//	|                                  | successful call to RNotifyBootConfigStatus has already been made.                |
	//	+----------------------------------+----------------------------------------------------------------------------------+
	NotifyBootConfigStatus(context.Context, *NotifyBootConfigStatusRequest) (*NotifyBootConfigStatusResponse, error)

	// Opnum10NotUsedOnWire operation.
	// Opnum10NotUsedOnWire

	// The RChangeServiceConfigW method changes a service's configuration parameters in
	// the SCM database.
	//
	// Return Values: The method returns 0x00000000 (ERROR_SUCCESS) on success; otherwise
	// it returns one of the following error codes.
	//
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	|                RETURN                |                                                                                  |
	//	|              VALUE/CODE              |                                   DESCRIPTION                                    |
	//	|                                      |                                                                                  |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| 5 ERROR_ACCESS_DENIED                | The SERVICE_CHANGE_CONFIG access right had not been granted to the caller when   |
	//	|                                      | the RPC context handle to the service record was created.                        |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| 6 ERROR_INVALID_HANDLE               | The handle specified is invalid.                                                 |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| 87 ERROR_INVALID_PARAMETER           | A parameter that was specified is invalid.                                       |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| 1057 ERROR_INVALID_SERVICE_ACCOUNT   | The user account name specified in the lpServiceStartName parameter does not     |
	//	|                                      | exist.                                                                           |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| 1059 ERROR_CIRCULAR_DEPENDENCY       | A circular service dependency was specified.                                     |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| 1078 ERROR_DUPLICATE_SERVICE_NAME    | The lpDisplayName matches either the ServiceName or the DisplayName of another   |
	//	|                                      | service record in the service control manager database.                          |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| 1072 ERROR_SERVICE_MARKED_FOR_DELETE | The RDeleteService has been called for the service record identified by the      |
	//	|                                      | hService parameter.                                                              |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| 1115 ERROR_SHUTDOWN_IN_PROGRESS      | The system is shutting down.                                                     |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	ChangeServiceConfigW(context.Context, *ChangeServiceConfigWRequest) (*ChangeServiceConfigWResponse, error)

	// The RCreateServiceW method creates the service record in the SCM database.
	//
	// Return Values: The method returns 0x00000000 (ERROR_SUCCESS) on success; otherwise,
	// it returns one of the following error codes.
	//
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	|                RETURN                |                                                                                  |
	//	|              VALUE/CODE              |                                   DESCRIPTION                                    |
	//	|                                      |                                                                                  |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| 5 ERROR_ACCESS_DENIED                | The SC_MANAGER_CREATE_SERVICE access right had not been granted to the caller    |
	//	|                                      | when the RPC context handle was created.                                         |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| 6 ERROR_INVALID_HANDLE               | The handle specified is invalid.                                                 |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| 13 ERROR_INVALID_DATA                | The data is invalid.                                                             |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| 87 ERROR_INVALID_PARAMETER           | A parameter that was specified is invalid.                                       |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| 123 ERROR_INVALID_NAME               | The specified service name is invalid.                                           |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| 1057 ERROR_INVALID_SERVICE_ACCOUNT   | The user account name specified in the lpServiceStartName parameter does not     |
	//	|                                      | exist.                                                                           |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| 1059 ERROR_CIRCULAR_DEPENDENCY       | A circular service dependency was specified.                                     |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| 1072 ERROR_SERVICE_MARKED_FOR_DELETE | The service record with a specified name already exists and RDeleteService has   |
	//	|                                      | been called for it.                                                              |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| 1073 ERROR_SERVICE_EXISTS            | The service record with the ServiceName matching the specified lpServiceName     |
	//	|                                      | already exists.                                                                  |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| 1078 ERROR_DUPLICATE_SERVICE_NAME    | The service record with the same DisplayName or the same ServiceName as the      |
	//	|                                      | passed in lpDisplayName already exists in the service control manager database.  |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| 1115 ERROR_SHUTDOWN_IN_PROGRESS      | The system is shutting down.                                                     |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	CreateServiceW(context.Context, *CreateServiceWRequest) (*CreateServiceWResponse, error)

	// The REnumDependentServicesW method returns the ServiceName, DisplayName, and ServiceStatus
	// values of service records that are listed as dependents of a specified service.
	//
	// Return Values: The method returns 0x00000000 (ERROR_SUCCESS) on success; otherwise,
	// it returns one of the following error codes.
	//
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	|             RETURN              |                                                                                  |
	//	|           VALUE/CODE            |                                   DESCRIPTION                                    |
	//	|                                 |                                                                                  |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	| 5 ERROR_ACCESS_DENIED           | The SERVICE_ENUMERATE_DEPENDENT access right had not been granted to the caller  |
	//	|                                 | when the RPC context handle to the service record was created.                   |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	| 6 ERROR_INVALID_HANDLE          | The handle is no longer valid.                                                   |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	| 87 ERROR_INVALID_PARAMETER      | A parameter that was specified is invalid.                                       |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	| 234 ERROR_MORE_DATA             | More data is available.                                                          |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	| 1115 ERROR_SHUTDOWN_IN_PROGRESS | The system is shutting down.                                                     |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	EnumDependentServicesW(context.Context, *EnumDependentServicesWRequest) (*EnumDependentServicesWResponse, error)

	// The REnumServicesStatusW method enumerates service records in the specified SCM database.
	//
	// Return Values: The method returns 0x00000000 (ERROR_SUCCESS) on success; otherwise,
	// it returns one of the following error codes.
	//
	//	+----------------------------+----------------------------------------------------------------------------------+
	//	|           RETURN           |                                                                                  |
	//	|         VALUE/CODE         |                                   DESCRIPTION                                    |
	//	|                            |                                                                                  |
	//	+----------------------------+----------------------------------------------------------------------------------+
	//	+----------------------------+----------------------------------------------------------------------------------+
	//	| 5 ERROR_ACCESS_DENIED      | The SM_MANAGER_ENUMERATE_SERVICE access right had not been granted to the caller |
	//	|                            | when the RPC context handle to the service record was created.                   |
	//	+----------------------------+----------------------------------------------------------------------------------+
	//	| 6 ERROR_INVALID_HANDLE     | The handle is no longer valid.                                                   |
	//	+----------------------------+----------------------------------------------------------------------------------+
	//	| 87 ERROR_INVALID_PARAMETER | A parameter that was specified is invalid.                                       |
	//	+----------------------------+----------------------------------------------------------------------------------+
	//	| 234 ERROR_MORE_DATA        | More data is available.                                                          |
	//	+----------------------------+----------------------------------------------------------------------------------+
	EnumServicesStatusW(context.Context, *EnumServicesStatusWRequest) (*EnumServicesStatusWResponse, error)

	// The ROpenSCManagerW method establishes a connection to server and opens the SCM database
	// on the specified server.
	//
	// Return Values: The method returns 0x00000000 (ERROR_SUCCESS) on success; otherwise,
	// it returns one of the following error codes.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN               |                                                                                  |
	//	|             VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 5 ERROR_ACCESS_DENIED              | The client does not have the required access rights to open the SCM              |
	//	|                                    | database on the server or the desired access is not granted to it in the SCM     |
	//	|                                    | SecurityDescriptor.                                                              |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 123 ERROR_INVALID_NAME             | The specified service name is invalid.                                           |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 1065 ERROR_DATABASE_DOES_NOT_EXIST | The database specified does not exist.                                           |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 1115 ERROR_SHUTDOWN_IN_PROGRESS    | The system is shutting down.                                                     |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	OpenSCMW(context.Context, *OpenSCMWRequest) (*OpenSCMWResponse, error)

	// The ROpenServiceW method creates an RPC context handle to an existing service record.
	//
	// Return Values: The method returns 0x00000000 (ERROR_SUCCESS) on success; otherwise,
	// it returns one of the following error codes.
	//
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	|              RETURN               |                                                                                  |
	//	|            VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                   |                                                                                  |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	| 5 ERROR_ACCESS_DENIED             | The access specified by the dwDesiredAccess parameter cannot be granted to the   |
	//	|                                   | caller.                                                                          |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	| 6 ERROR_INVALID_HANDLE            | The handle is no longer valid.                                                   |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	| 123 ERROR_INVALID_NAME            | The specified service name is invalid.                                           |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	| 1060 ERROR_SERVICE_DOES_NOT_EXIST | The service record with a specified DisplayName does not exist in the SCM        |
	//	|                                   | database.                                                                        |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	| 1115 ERROR_SHUTDOWN_IN_PROGRESS   | The system is shutting down.                                                     |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	OpenServiceW(context.Context, *OpenServiceWRequest) (*OpenServiceWResponse, error)

	// The RQueryServiceConfigW method returns the configuration parameters of the specified
	// service.
	//
	// Return Values: The method returns 0x00000000 (ERROR_SUCCESS) on success; otherwise,
	// it returns one of the following error codes.
	//
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	|             RETURN              |                                                                                  |
	//	|           VALUE/CODE            |                                   DESCRIPTION                                    |
	//	|                                 |                                                                                  |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	| 5 ERROR_ACCESS_DENIED           | The SERVICE_QUERY_CONFIG access right had not been granted to the caller when    |
	//	|                                 | the RPC context handle was created.                                              |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	| 6 ERROR_INVALID_HANDLE          | The handle is no longer valid.                                                   |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	| 122 ERROR_INSUFFICIENT_BUFFER   | The data area passed to a system call is too small.                              |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	| 1115 ERROR_SHUTDOWN_IN_PROGRESS | The system is shutting down.                                                     |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	QueryServiceConfigW(context.Context, *QueryServiceConfigWRequest) (*QueryServiceConfigWResponse, error)

	// The RQueryServiceLockStatusW method returns the lock status of the specified SCM
	// database.
	//
	// Return Values: The method returns 0x00000000 (ERROR_SUCCESS) on success; otherwise,
	// it returns one of the following error codes.
	//
	//	+-------------------------------+----------------------------------------------------------------------------------+
	//	|            RETURN             |                                                                                  |
	//	|          VALUE/CODE           |                                   DESCRIPTION                                    |
	//	|                               |                                                                                  |
	//	+-------------------------------+----------------------------------------------------------------------------------+
	//	+-------------------------------+----------------------------------------------------------------------------------+
	//	| 5 ERROR_ACCESS_DENIED         | The SC_MANAGER_QUERY_LOCK_STATUS access right had not been granted to the caller |
	//	|                               | when the RPC context handle was created.                                         |
	//	+-------------------------------+----------------------------------------------------------------------------------+
	//	| 6 ERROR_INVALID_HANDLE        | The handle is no longer valid.                                                   |
	//	+-------------------------------+----------------------------------------------------------------------------------+
	//	| 122 ERROR_INSUFFICIENT_BUFFER | The data area passed to a system call is too small.                              |
	//	+-------------------------------+----------------------------------------------------------------------------------+
	QueryServiceLockStatusW(context.Context, *QueryServiceLockStatusWRequest) (*QueryServiceLockStatusWResponse, error)

	// The RStartServiceW method starts a specified service.
	//
	// Return Values: The method returns 0x00000000 (ERROR_SUCCESS) on success; otherwise,
	// it returns one of the following error codes.<41>
	//
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	|                RETURN                 |                                                                                  |
	//	|              VALUE/CODE               |                                   DESCRIPTION                                    |
	//	|                                       |                                                                                  |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| 2 ERROR_FILE_NOT_FOUND                | The system cannot find the file specified.                                       |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| 3 ERROR_PATH_NOT_FOUND                | The system cannot find the path specified.                                       |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| 5 ERROR_ACCESS_DENIED                 | The SERVICE_START access right had not been granted to the caller when the RPC   |
	//	|                                       | context handle to the service record was created.                                |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| 6 ERROR_INVALID_HANDLE                | The handle is no longer valid.                                                   |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| 87 ERROR_INVALID_PARAMETER            | A parameter that was specified is invalid.                                       |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| 1053 ERROR_SERVICE_REQUEST_TIMEOUT    | The process for the service was started, but it did not respond within an        |
	//	|                                       | implementation-specific time-out.<42>                                            |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| 1054 ERROR_SERVICE_NO_THREAD          | A thread could not be created for the service.                                   |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| 1055 ERROR_SERVICE_DATABASE_LOCKED    | The service database is locked by the call to the BlockServiceDatabase           |
	//	|                                       | method.<43>                                                                      |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| 1056 ERROR_SERVICE_ALREADY_RUNNING    | The ServiceStatus.dwCurrentState in the service record is not set to             |
	//	|                                       | SERVICE_STOPPED.                                                                 |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| 1058 ERROR_SERVICE_DISABLED           | The service cannot be started because the Start field in the service record is   |
	//	|                                       | set to SERVICE_DISABLED.                                                         |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| 1068 ERROR_SERVICE_DEPENDENCY_FAIL    | The specified service depends on another service that has failed to start.       |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| 1069 ERROR_SERVICE_LOGON_FAILED       | The service did not start due to a logon failure.                                |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| 1072 ERROR_SERVICE_MARKED_FOR_DELETE  | The RDeleteService method has been called for the service record identified by   |
	//	|                                       | the hService parameter.                                                          |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| 1075 ERROR_SERVICE_DEPENDENCY_DELETED | The specified service depends on a service that does not exist or has been       |
	//	|                                       | marked for deletion.                                                             |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| 1115 ERROR_SHUTDOWN_IN_PROGRESS       | The system is shutting down.                                                     |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	StartServiceW(context.Context, *StartServiceWRequest) (*StartServiceWResponse, error)

	// The RGetServiceDisplayNameW method returns the display name of the specified service.
	//
	// Return Values: The method returns 0x00000000 (ERROR_SUCCESS) on success; otherwise,
	// it returns one of the following error codes.
	//
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	|              RETURN               |                                                                                  |
	//	|            VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                   |                                                                                  |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	| 122 ERROR_INSUFFICIENT_BUFFER     | The display name does not fit in the buffer.                                     |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	| 123 ERROR_INVALID_NAME            | The specified service name is invalid.                                           |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	| 1060 ERROR_SERVICE_DOES_NOT_EXIST | The service record with the specified ServiceName does not exist in the SCM      |
	//	|                                   | database identified by the hSCManager parameter.                                 |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	GetServiceDisplayNameW(context.Context, *GetServiceDisplayNameWRequest) (*GetServiceDisplayNameWResponse, error)

	// The RGetServiceKeyNameW method returns the ServiceName of the service record with
	// the specified DisplayName.
	//
	// Return Values: The method returns 0x00000000 (ERROR_SUCCESS) on success; otherwise,
	// it returns one of the following error codes.
	//
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	|              RETURN               |                                                                                  |
	//	|            VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                   |                                                                                  |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	| 123 ERROR_INVALID_NAME            | The name specified in the lpDisplayName parameter is invalid or set to NULL.     |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	| 1060 ERROR_SERVICE_DOES_NOT_EXIST | The service record with the DisplayName matching the value specified in the      |
	//	|                                   | lpDisplayName parameter does not exist in the SCM database identified by the     |
	//	|                                   | hSCManager parameter.                                                            |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	GetServiceKeyNameW(context.Context, *GetServiceKeyNameWRequest) (*GetServiceKeyNameWResponse, error)

	// Opnum22NotUsedOnWire operation.
	// Opnum22NotUsedOnWire

	// The RChangeServiceConfigA method changes a service's configuration parameters in
	// the SCM database.
	//
	// Return Values: The method returns 0x00000000 (ERROR_SUCCESS) on success; otherwise,
	// it returns one of the following error codes.
	//
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	|                RETURN                |                                                                                  |
	//	|              VALUE/CODE              |                                   DESCRIPTION                                    |
	//	|                                      |                                                                                  |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| 5 ERROR_ACCESS_DENIED                | The SERVICE_CHANGE_CONFIG access right had not been granted to the caller when   |
	//	|                                      | the RPC context handle to the service record was created.                        |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| 6 ERROR_INVALID_HANDLE               | The handle specified is invalid.                                                 |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| 87 ERROR_INVALID_PARAMETER           | A parameter that was specified is invalid.                                       |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| 1057 ERROR_INVALID_SERVICE_ACCOUNT   | The user account name specified in the lpServiceStartName parameter does not     |
	//	|                                      | exist.                                                                           |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| 1059 ERROR_CIRCULAR_DEPENDENCY       | A circular service dependency was specified.                                     |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| 1078 ERROR_DUPLICATE_SERVICE_NAME    | The lpDisplayName matches either the ServiceName or the DisplayName of another   |
	//	|                                      | service record in the service control manager database.                          |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| 1072 ERROR_SERVICE_MARKED_FOR_DELETE | The RDeleteService has been called for the service record identified by the      |
	//	|                                      | hService parameter.                                                              |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| 1115 ERROR_SHUTDOWN_IN_PROGRESS      | The system is shutting down.                                                     |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	ChangeServiceConfigA(context.Context, *ChangeServiceConfigARequest) (*ChangeServiceConfigAResponse, error)

	// The RCreateServiceA method creates the service record in the SCM database.
	//
	// Return Values: The method returns 0x00000000 (ERROR_SUCCESS) on success; otherwise,
	// it returns one of the following error codes.
	//
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	|                RETURN                |                                                                                  |
	//	|              VALUE/CODE              |                                   DESCRIPTION                                    |
	//	|                                      |                                                                                  |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| 5 ERROR_ACCESS_DENIED                | The SC_MANAGER_CREATE_SERVICE access right had not been granted to the caller    |
	//	|                                      | when the RPC context handle was created.                                         |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| 6 ERROR_INVALID_HANDLE               | The handle specified is invalid.                                                 |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| 13 ERROR_INVALID_DATA                | The data is invalid.                                                             |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| 87 ERROR_INVALID_PARAMETER           | A parameter that was specified is invalid.                                       |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| 123 ERROR_INVALID_NAME               | The specified service name is invalid.                                           |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| 1057 ERROR_INVALID_SERVICE_ACCOUNT   | The user account name specified in the lpServiceStartName parameter does not     |
	//	|                                      | exist.                                                                           |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| 1059 ERROR_CIRCULAR_DEPENDENCY       | A circular service dependency was specified.                                     |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| 1072 ERROR_SERVICE_MARKED_FOR_DELETE | The service record with a specified name already exists, and RDeleteService has  |
	//	|                                      | been called for it.                                                              |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| 1073 ERROR_SERVICE_EXISTS            | The service record with the ServiceName matching the specified lpServiceName     |
	//	|                                      | already exists.                                                                  |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| 1078 ERROR_DUPLICATE_SERVICE_NAME    | The service record with the same DisplayName or the same ServiceName as the      |
	//	|                                      | passed-in lpDisplayName already exists in the service control manager database.  |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| 1115 ERROR_SHUTDOWN_IN_PROGRESS      | The system is shutting down.                                                     |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	CreateServiceA(context.Context, *CreateServiceARequest) (*CreateServiceAResponse, error)

	// The REnumDependentServicesA method returns the ServiceName, DisplayName, and ServiceStatus
	// of each service record that depends on the specified service.
	//
	// Return Values: The method returns 0x00000000 (ERROR_SUCCESS) on success; otherwise,
	// it returns one of the following error codes.
	//
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	|             RETURN              |                                                                                  |
	//	|           VALUE/CODE            |                                   DESCRIPTION                                    |
	//	|                                 |                                                                                  |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	| 5 ERROR_ACCESS_DENIED           | The SERVICE_ENUMERATE_DEPENDENT access right had not been granted to the caller  |
	//	|                                 | when the RPC context handle to the service record was created.                   |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	| 6 ERROR_INVALID_HANDLE          | The handle is no longer valid.                                                   |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	| 87 ERROR_INVALID_PARAMETER      | A parameter that was specified is invalid.                                       |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	| 234 ERROR_MORE_DATA             | More data is available.                                                          |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	| 1115 ERROR_SHUTDOWN_IN_PROGRESS | The system is shutting down.                                                     |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	EnumDependentServicesA(context.Context, *EnumDependentServicesARequest) (*EnumDependentServicesAResponse, error)

	// The REnumServicesStatusA method enumerates service records in the specified SCM database.
	//
	// Return Values: The method returns 0x00000000 (ERROR_SUCCESS) on success; otherwise,
	// it returns one of the following error codes.
	//
	//	+----------------------------+----------------------------------------------------------------------------------+
	//	|           RETURN           |                                                                                  |
	//	|         VALUE/CODE         |                                   DESCRIPTION                                    |
	//	|                            |                                                                                  |
	//	+----------------------------+----------------------------------------------------------------------------------+
	//	+----------------------------+----------------------------------------------------------------------------------+
	//	| 5 ERROR_ACCESS_DENIED      | The SC_MANAGER_ENUMERATE_SERVICE access right had not been granted to the caller |
	//	|                            | when the RPC context handle to the service record was created.                   |
	//	+----------------------------+----------------------------------------------------------------------------------+
	//	| 6 ERROR_INVALID_HANDLE     | The handle is no longer valid.                                                   |
	//	+----------------------------+----------------------------------------------------------------------------------+
	//	| 87 ERROR_INVALID_PARAMETER | A parameter that was specified is invalid.                                       |
	//	+----------------------------+----------------------------------------------------------------------------------+
	//	| 234 ERROR_MORE_DATA        | More data is available.                                                          |
	//	+----------------------------+----------------------------------------------------------------------------------+
	EnumServicesStatusA(context.Context, *EnumServicesStatusARequest) (*EnumServicesStatusAResponse, error)

	// The ROpenSCManagerA method opens a connection to the SCM from the client and then
	// opens the specified SCM database.
	//
	// Return Values: The method returns 0x00000000 (ERROR_SUCCESS) on success; otherwise,
	// it returns one of the following error codes.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN               |                                                                                  |
	//	|             VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 5 ERROR_ACCESS_DENIED              | The SC_MANAGER_CONNECT access right or the desired access is not granted to the  |
	//	|                                    | caller in the SCM SecurityDescriptor.                                            |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 123 ERROR_INVALID_NAME             | The specified service name is invalid.                                           |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 1065 ERROR_DATABASE_DOES_NOT_EXIST | The database specified does not exist.                                           |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 1115 ERROR_SHUTDOWN_IN_PROGRESS    | The system is shutting down.                                                     |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	OpenSCMA(context.Context, *OpenSCMARequest) (*OpenSCMAResponse, error)

	// The ROpenServiceA method creates an RPC context handle to an existing service record.
	//
	// Return Values: The method returns 0x00000000 (ERROR_SUCCESS) on success; otherwise,
	// it returns one of the following error codes.
	//
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	|              RETURN               |                                                                                  |
	//	|            VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                   |                                                                                  |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	| 5 ERROR_ACCESS_DENIED             | The access specified by the dwDesiredAccess parameter cannot be granted to the   |
	//	|                                   | caller.                                                                          |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	| 6 ERROR_INVALID_HANDLE            | The handle is no longer valid.                                                   |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	| 123 ERROR_INVALID_NAME            | The specified service name is invalid.                                           |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	| 1060 ERROR_SERVICE_DOES_NOT_EXIST | The service record with a specified DisplayName does not exist in the SCM        |
	//	|                                   | database.                                                                        |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	| 1115 ERROR_SHUTDOWN_IN_PROGRESS   | The system is shutting down.                                                     |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	OpenServiceA(context.Context, *OpenServiceARequest) (*OpenServiceAResponse, error)

	// The RQueryServiceConfigA method returns the configuration parameters of the specified
	// service.
	//
	// Return Values: The method returns 0x00000000 (ERROR_SUCCESS) on success; otherwise,
	// it returns one of the following error codes.
	//
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	|             RETURN              |                                                                                  |
	//	|           VALUE/CODE            |                                   DESCRIPTION                                    |
	//	|                                 |                                                                                  |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	| 5 ERROR_ACCESS_DENIED           | The SERVICE_QUERY_CONFIG access right had not been granted to the caller when    |
	//	|                                 | the RPC context handle was created.                                              |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	| 6 ERROR_INVALID_HANDLE          | The handle is no longer valid.                                                   |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	| 122 ERROR_INSUFFICIENT_BUFFER   | The data area passed to a system call is too small.                              |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	| 1115 ERROR_SHUTDOWN_IN_PROGRESS | The system is shutting down.                                                     |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	QueryServiceConfigA(context.Context, *QueryServiceConfigARequest) (*QueryServiceConfigAResponse, error)

	// The RQueryServiceLockStatusA method returns the lock status of the specified SCM
	// database.
	//
	// Return Values: The method returns 0x00000000 (ERROR_SUCCESS) on success; otherwise,
	// it returns one of the following error codes.
	//
	//	+-------------------------------+----------------------------------------------------------------------------------+
	//	|            RETURN             |                                                                                  |
	//	|          VALUE/CODE           |                                   DESCRIPTION                                    |
	//	|                               |                                                                                  |
	//	+-------------------------------+----------------------------------------------------------------------------------+
	//	+-------------------------------+----------------------------------------------------------------------------------+
	//	| 5 ERROR_ACCESS_DENIED         | The SC_MANAGER_QUERY_LOCK_STATUS access right had not been granted to the caller |
	//	|                               | when the RPC context handle was created.                                         |
	//	+-------------------------------+----------------------------------------------------------------------------------+
	//	| 6 ERROR_INVALID_HANDLE        | The handle is no longer valid.                                                   |
	//	+-------------------------------+----------------------------------------------------------------------------------+
	//	| 122 ERROR_INSUFFICIENT_BUFFER | The data area passed to a system call is too small.                              |
	//	+-------------------------------+----------------------------------------------------------------------------------+
	QueryServiceLockStatusA(context.Context, *QueryServiceLockStatusARequest) (*QueryServiceLockStatusAResponse, error)

	// The RStartServiceA method starts a specified service.
	//
	// Return Values: The method returns 0x00000000 (ERROR_SUCCESS) on success; otherwise,
	// it returns one of the following error codes.<46>
	//
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	|                RETURN                 |                                                                                  |
	//	|              VALUE/CODE               |                                   DESCRIPTION                                    |
	//	|                                       |                                                                                  |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| 2 ERROR_FILE_NOT_FOUND                | The system cannot find the file specified.                                       |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| 3 ERROR_PATH_NOT_FOUND                | The system cannot find the path specified.                                       |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| 5 ERROR_ACCESS_DENIED                 | The SERVICE_START access right had not been granted to the caller when the RPC   |
	//	|                                       | context handle to the service was created.                                       |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| 6 ERROR_INVALID_HANDLE                | The handle is no longer valid.                                                   |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| 87 ERROR_INVALID_PARAMETER            | A parameter that was specified is invalid.                                       |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| 1053 ERROR_SERVICE_REQUEST_TIMEOUT    | The process for the service was started, but it did not respond within an        |
	//	|                                       | implementation-specific time-out.<47>                                            |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| 1054 ERROR_SERVICE_NO_THREAD          | A thread could not be created for the service.                                   |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| 1055 ERROR_SERVICE_DATABASE_LOCKED    | The service database is locked by the call to the RLockServiceDatabase (section  |
	//	|                                       | 3.1.4.4) method.<48>                                                             |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| 1056 ERROR_SERVICE_ALREADY_RUNNING    | The ServiceStatus.dwCurrentState in the service record is not set to             |
	//	|                                       | SERVICE_STOPPED.                                                                 |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| 1058 ERROR_SERVICE_DISABLED           | The service cannot be started because the Start field in the service record is   |
	//	|                                       | set to SERVICE_DISABLED.                                                         |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| 1068 ERROR_SERVICE_DEPENDENCY_FAIL    | The specified service depends on another service that has failed to start.       |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| 1069 ERROR_SERVICE_LOGON_FAILED       | The service did not start due to a logon failure.                                |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| 1072 ERROR_SERVICE_MARKED_FOR_DELETE  | The RDeleteService method has been called for the service record identified by   |
	//	|                                       | the hService parameter.                                                          |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| 1075 ERROR_SERVICE_DEPENDENCY_DELETED | The specified service depends on a service that does not exist or has been       |
	//	|                                       | marked for deletion.                                                             |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| 1115 ERROR_SHUTDOWN_IN_PROGRESS       | The system is shutting down.                                                     |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	StartServiceA(context.Context, *StartServiceARequest) (*StartServiceAResponse, error)

	// The RGetServiceDisplayNameA method returns the display name of the specified service.
	//
	// Return Values: The method returns 0x00000000 (ERROR_SUCCESS) on success; otherwise,
	// it returns one of the following error codes.
	//
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	|              RETURN               |                                                                                  |
	//	|            VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                   |                                                                                  |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	| 122 ERROR_INSUFFICIENT_BUFFER     | The display name does not fit in the buffer.                                     |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	| 123 ERROR_INVALID_NAME            | The specified service name is invalid.                                           |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	| 1060 ERROR_SERVICE_DOES_NOT_EXIST | The service record with the specified ServiceName does not exist in the SCM      |
	//	|                                   | database identified by the hSCManager parameter.                                 |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	GetServiceDisplayNameA(context.Context, *GetServiceDisplayNameARequest) (*GetServiceDisplayNameAResponse, error)

	// The RGetServiceKeyNameA method returns the ServiceName of the service record with
	// the specified DisplayName.
	//
	// Return Values: The method returns 0x00000000 (ERROR_SUCCESS) on success; otherwise,
	// it returns one of the following error codes.
	//
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	|              RETURN               |                                                                                  |
	//	|            VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                   |                                                                                  |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	| 122 ERROR_INSUFFICIENT_BUFFER     | The data area passed to a system call is too small.                              |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	| 123 ERROR_INVALID_NAME            | The name specified in lpDisplayName is invalid or set to NULL.                   |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	| 1060 ERROR_SERVICE_DOES_NOT_EXIST | The service record with the DisplayName matching the specified lpDisplayName     |
	//	|                                   | does not exist in the SCM database identified by the hSCManager parameter.       |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	GetServiceKeyNameA(context.Context, *GetServiceKeyNameARequest) (*GetServiceKeyNameAResponse, error)

	// Opnum34NotUsedOnWire operation.
	// Opnum34NotUsedOnWire

	// The REnumServiceGroupW method returns the members of a service group.
	//
	// Return Values: The method returns 0x00000000 (ERROR_SUCCESS) on success; otherwise,
	// it returns one of the following error codes.
	//
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	|              RETURN               |                                                                                  |
	//	|            VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                   |                                                                                  |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	| 5 ERROR_ACCESS_DENIED             | The SC_MANAGER_ENUMERATE_SERVICE access right had not been granted to the caller |
	//	|                                   | when the RPC context handle was created.                                         |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	| 6 ERROR_INVALID_HANDLE            | The handle is no longer valid.                                                   |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	| 87 ERROR_INVALID_PARAMETER        | A parameter that was specified is invalid.                                       |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	| 234 ERROR_MORE_DATA               | More data is available.                                                          |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	| 1060 ERROR_SERVICE_DOES_NOT_EXIST | The group specified by pszGroupName does not exist in the SCM GroupList.         |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	| 1115 ERROR_SHUTDOWN_IN_PROGRESS   | The system is shutting down.                                                     |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	EnumServiceGroupW(context.Context, *EnumServiceGroupWRequest) (*EnumServiceGroupWResponse, error)

	// The RChangeServiceConfig2A method SHOULD<51> change the optional configuration parameters
	// of a service.
	//
	// Return Values: The method returns 0x00000000 (ERROR_SUCCESS) on success; otherwise
	// it returns one of the following error codes.<52>
	//
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	|                 RETURN                  |                                                                                  |
	//	|               VALUE/CODE                |                                   DESCRIPTION                                    |
	//	|                                         |                                                                                  |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| 5 ERROR_ACCESS_DENIED                   | The SERVICE_CHANGE_CONFIG access right had not been granted to the caller when   |
	//	|                                         | the RPC context handle to the service record was created.                        |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| 6 ERROR_INVALID_HANDLE                  | The handle is no longer valid.                                                   |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| 87 ERROR_INVALID_PARAMETER              | A parameter that was specified is invalid.                                       |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| 1072 ERROR_SERVICE_MARKED_FOR_DELETE    | The RDeleteService has been called for the service record identified by the      |
	//	|                                         | hService parameter.                                                              |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| 1080 ERROR_CANNOT_DETECT_DRIVER_FAILURE | SERVICE_CONFIG_FAILURE_ACTIONS cannot be used as a dwInfoLevel in the Info       |
	//	|                                         | parameter for service records with a Type value defined for drivers.             |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| 1115 ERROR_SHUTDOWN_IN_PROGRESS         | The system is shutting down.                                                     |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	ChangeServiceConfig2A(context.Context, *ChangeServiceConfig2ARequest) (*ChangeServiceConfig2AResponse, error)

	// The RChangeServiceConfig2W <53> method changes the optional configuration parameters
	// of a service.
	//
	// Return Values: The method returns 0x00000000 (ERROR_SUCCESS) on success; otherwise
	// it returns one of the following error codes.<54>
	//
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	|                 RETURN                  |                                                                                  |
	//	|               VALUE/CODE                |                                   DESCRIPTION                                    |
	//	|                                         |                                                                                  |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| 5 ERROR_ACCESS_DENIED                   | The SERVICE_CHANGE_CONFIG access right had not been granted to the caller when   |
	//	|                                         | the RPC context handle to the service record was created.                        |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| 6 ERROR_INVALID_HANDLE                  | The handle is no longer valid.                                                   |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| 87 ERROR_INVALID_PARAMETER              | A parameter that was specified is invalid.                                       |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| 1072 ERROR_SERVICE_MARKED_FOR_DELETE    | The RDeleteService has been called for the service record identified by the      |
	//	|                                         | hService parameter.                                                              |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| 1080 ERROR_CANNOT_DETECT_DRIVER_FAILURE | SERVICE_CONFIG_FAILURE_ACTIONS cannot be used as a dwInfoLevel in the Info       |
	//	|                                         | parameter for service records with a Type value defined for drivers.             |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| 1115 ERROR_SHUTDOWN_IN_PROGRESS         | The system is shutting down.                                                     |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	ChangeServiceConfig2W(context.Context, *ChangeServiceConfig2WRequest) (*ChangeServiceConfig2WResponse, error)

	// The RQueryServiceConfig2A <55> method returns the optional configuration parameters
	// of the specified service based on the specified information level.
	//
	// Return Values: The method returns 0x00000000 (ERROR_SUCCESS) on success; otherwise,
	// it returns one of the following error codes.
	//
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	|             RETURN              |                                                                                  |
	//	|           VALUE/CODE            |                                   DESCRIPTION                                    |
	//	|                                 |                                                                                  |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	| 5 ERROR_ACCESS_DENIED           | The SERVICE_QUERY_CONFIG access right had not been granted to the caller when    |
	//	|                                 | the RPC context handle to the service record was created.                        |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	| 6 ERROR_INVALID_HANDLE          | The handle is no longer valid.                                                   |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	| 87 ERROR_INVALID_PARAMETER      | A parameter that was specified is invalid.                                       |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	| 122 ERROR_INSUFFICIENT_BUFFER   | The data area passed to a system call is too small.                              |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	| 124 ERROR_INVALID_LEVEL         | The dwInfoLevel parameter contains an unsupported value.                         |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	| 1115 ERROR_SHUTDOWN_IN_PROGRESS | The system is shutting down.                                                     |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	QueryServiceConfig2A(context.Context, *QueryServiceConfig2ARequest) (*QueryServiceConfig2AResponse, error)

	// The RQueryServiceConfig2W <64> method returns the optional configuration parameters
	// of the specified service based on the specified information level.
	//
	// Return Values: The method returns 0x00000000 (ERROR_SUCCESS) on success; otherwise,
	// it returns one of the following error codes.
	//
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	|                RETURN                 |                                                                                  |
	//	|              VALUE/CODE               |                                   DESCRIPTION                                    |
	//	|                                       |                                                                                  |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000005 ERROR_ACCESS_DENIED        | The SERVICE_QUERY_CONFIG access right had not been granted to the caller when    |
	//	|                                       | the RPC context handle to the service record was created.                        |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000006 ERROR_INVALID_HANDLE       | The handle is no longer valid.                                                   |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000087 ERROR_INVALID_PARAMETER    | A parameter that was specified is invalid.                                       |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000122 ERROR_INSUFFICIENT_BUFFER  | The data area passed to a system call is too small.                              |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000124 ERROR_INVALID_LEVEL        | The dwInfoLevel parameter contains an unsupported value.                         |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00001115 ERROR_SHUTDOWN_IN_PROGRESS | The system is shutting down.                                                     |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	QueryServiceConfig2W(context.Context, *QueryServiceConfig2WRequest) (*QueryServiceConfig2WResponse, error)

	// The RQueryServiceStatusEx method returns the current status of the specified service,
	// based on the specified information level.
	//
	// Return Values: The method returns 0x00000000 (ERROR_SUCCESS) on success; otherwise,
	// it returns one of the following error codes.
	//
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	|             RETURN              |                                                                                  |
	//	|           VALUE/CODE            |                                   DESCRIPTION                                    |
	//	|                                 |                                                                                  |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	| 5 ERROR_ACCESS_DENIED           | The SERVICE_QUERY_STATUS access right had not been granted to the caller when    |
	//	|                                 | the RPC context handle to the service record was created.                        |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	| 6 ERROR_INVALID_HANDLE          | The handle is no longer valid.                                                   |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	| 87 ERROR_INVALID_PARAMETER      | A parameter that was specified is invalid.                                       |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	| 122 ERROR_INSUFFICIENT_BUFFER   | The data area passed to a system call is too small.                              |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	| 124 ERROR_INVALID_LEVEL         | The InfoLevel parameter contains an unsupported value.                           |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	| 1115 ERROR_SHUTDOWN_IN_PROGRESS | The system is shutting down.                                                     |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	QueryServiceStatusEx(context.Context, *QueryServiceStatusExRequest) (*QueryServiceStatusExResponse, error)

	// The REnumServicesStatusExA method enumerates services in the specified SCM database,
	// based on the specified information level.
	//
	// Return Values: The method returns 0x00000000 (ERROR_SUCCESS) on success; otherwise,
	// it returns one of the following error codes.
	//
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	|              RETURN               |                                                                                  |
	//	|            VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                   |                                                                                  |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	| 5 ERROR_ACCESS_DENIED             | The SC_MANAGER_ENUMERATE_SERVICE access right had not been granted to the caller |
	//	|                                   | when the RPC context handle to the SCM was created.                              |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	| 6 ERROR_INVALID_HANDLE            | The handle is no longer valid.                                                   |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	| 87 ERROR_INVALID_PARAMETER        | A parameter that was specified is invalid.                                       |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	| 124 ERROR_INVALID_LEVEL           | The InfoLevel parameter contains an unsupported value.                           |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	| 234 ERROR_MORE_DATA               | More data is available.                                                          |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	| 1060 ERROR_SERVICE_DOES_NOT_EXIST | The group specified by the pszGroupName parameter does not exist in the SCM      |
	//	|                                   | GroupList.                                                                       |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	| 1115 ERROR_SHUTDOWN_IN_PROGRESS   | The system is shutting down.                                                     |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	EnumServicesStatusExA(context.Context, *EnumServicesStatusExARequest) (*EnumServicesStatusExAResponse, error)

	// The REnumServicesStatusExW method enumerates services in the specified SCM database,
	// based on the specified information level.
	//
	// Return Values: The method returns 0x00000000 (ERROR_SUCCESS) (ERROR_SUCCESS) on success;
	// otherwise, it returns one of the following error codes.
	//
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	|              RETURN               |                                                                                  |
	//	|            VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                   |                                                                                  |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	| 5 ERROR_ACCESS_DENIED             | The SC_MANAGER_ENUMERATE_SERVICE access right had not been granted to the caller |
	//	|                                   | when the RPC context handle to the SCM was created.                              |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	| 6 ERROR_INVALID_HANDLE            | The handle is no longer valid.                                                   |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	| 87 ERROR_INVALID_PARAMETER        | A parameter that was specified is invalid.                                       |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	| 124 ERROR_INVALID_LEVEL           | The InfoLevel parameter contains an unsupported value.                           |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	| 234 ERROR_MORE_DATA               | More data is available.                                                          |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	| 1060 ERROR_SERVICE_DOES_NOT_EXIST | The group specified by the pszGroupName parameter does not exist in the SCM      |
	//	|                                   | GroupList.                                                                       |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	| 1115 ERROR_SHUTDOWN_IN_PROGRESS   | The system is shutting down.                                                     |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	EnumServicesStatusExW(context.Context, *EnumServicesStatusExWRequest) (*EnumServicesStatusExWResponse, error)

	// Opnum43NotUsedOnWire operation.
	// Opnum43NotUsedOnWire

	// The RCreateServiceWOW64A method creates the service record for a 32-bit service on
	// a 64-bit system with the path to the file image automatically adjusted to point to
	// a 32-bit file location on the system.
	//
	// Return Values: The method returns 0x00000000 (ERROR_SUCCESS) on success; otherwise,
	// one of the following error codes.
	//
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	|                RETURN                |                                                                                  |
	//	|              VALUE/CODE              |                                   DESCRIPTION                                    |
	//	|                                      |                                                                                  |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| 5 ERROR_ACCESS_DENIED                | The SC_MANAGER_CREATE_SERVICE access right had not been granted to the caller    |
	//	|                                      | when the RPC context handle to the SCM was created.                              |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| 6 ERROR_INVALID_HANDLE               | The handle specified is invalid.                                                 |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| 13 ERROR_INVALID_DATA                | The data is invalid.                                                             |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| 87 ERROR_INVALID_PARAMETER           | A parameter that was specified is invalid.                                       |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| 123 ERROR_INVALID_NAME               | The specified service name is invalid.                                           |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| 1057 ERROR_INVALID_SERVICE_ACCOUNT   | The user account name specified in the lpServiceStartName parameter does not     |
	//	|                                      | exist.                                                                           |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| 1059 ERROR_CIRCULAR_DEPENDENCY       | A circular service dependency was specified.                                     |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| 1072 ERROR_SERVICE_MARKED_FOR_DELETE | The service record with a specified name already exists and RDeleteService has   |
	//	|                                      | been called for it.                                                              |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| 1073 ERROR_SERVICE_EXISTS            | The service record with the ServiceName matching the specified lpServiceName     |
	//	|                                      | already exists.                                                                  |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| 1078 ERROR_DUPLICATE_SERVICE_NAME    | The service record with the same DisplayName or the same ServiceName as the      |
	//	|                                      | passed-in lpDisplayName already exists in the SCM database.                      |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| 1115 ERROR_SHUTDOWN_IN_PROGRESS      | The system is shutting down.                                                     |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	CreateServiceWOW64A(context.Context, *CreateServiceWOW64ARequest) (*CreateServiceWOW64AResponse, error)

	// The RCreateServiceWOW64W method creates the service record for a 32-bit service on
	// a 64-bit system with the path to the file image automatically adjusted to point to
	// a 32-bit file location on the system.
	//
	// Return Values: The method returns 0x00000000 (ERROR_SUCCESS) on success; otherwise,
	// it returns one of the following error codes.
	//
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	|                RETURN                |                                                                                  |
	//	|              VALUE/CODE              |                                   DESCRIPTION                                    |
	//	|                                      |                                                                                  |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| 5 ERROR_ACCESS_DENIED                | The SC_MANAGER_CREATE_SERVICE access right had not been granted to the caller    |
	//	|                                      | when the RPC context handle to the SCM was created.                              |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| 6 ERROR_INVALID_HANDLE               | The handle specified is invalid.                                                 |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| 13 ERROR_INVALID_DATA                | The data is invalid.                                                             |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| 87 ERROR_INVALID_PARAMETER           | A parameter that was specified is invalid.                                       |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| 123 ERROR_INVALID_NAME               | The specified service name is invalid.                                           |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| 1057 ERROR_INVALID_SERVICE_ACCOUNT   | The user account name specified in the lpServiceStartName parameter does not     |
	//	|                                      | exist.                                                                           |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| 1059 ERROR_CIRCULAR_DEPENDENCY       | A circular service dependency was specified.                                     |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| 1072 ERROR_SERVICE_MARKED_FOR_DELETE | The service record with a specified name already exists, and RDeleteService has  |
	//	|                                      | been called for it.                                                              |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| 1073 ERROR_SERVICE_EXISTS            | The service record with the ServiceName matching the specified lpServiceName     |
	//	|                                      | already exists.                                                                  |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| 1078 ERROR_DUPLICATE_SERVICE_NAME    | The service record with the same DisplayName or the same ServiceName as the      |
	//	|                                      | passed-in lpDisplayName already exists in the service control manager database.  |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| 1115 ERROR_SHUTDOWN_IN_PROGRESS      | The system is shutting down.                                                     |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	CreateServiceWOW64W(context.Context, *CreateServiceWOW64WRequest) (*CreateServiceWOW64WResponse, error)

	// Opnum46NotUsedOnWire operation.
	// Opnum46NotUsedOnWire

	// The RNotifyServiceStatusChange method<74> allows the client to register for notifications
	// and check, via RGetNotifyResults (section 3.1.4.44), when the specified service of
	// type SERVICE_WIN32_OWN_PROCESS or SERVICE_WIN32_SHARE_PROCESS is created or deleted
	// or when its status changes.
	//
	// Return Values: The method returns 0x00000000 (ERROR_SUCCESS) on success; otherwise,
	// it returns one of the following error codes.
	//
	//	+------------------------------------------+----------------------------------------------------------------------------------+
	//	|                  RETURN                  |                                                                                  |
	//	|                VALUE/CODE                |                                   DESCRIPTION                                    |
	//	|                                          |                                                                                  |
	//	+------------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------------+----------------------------------------------------------------------------------+
	//	| 5 ERROR_ACCESS_DENIED                    | The SC_MANAGER_ENUMERATE_SERVICE access right had not been granted to the caller |
	//	|                                          | when the RPC context handle to the SCM was created, or the SERVICE_QUERY_STATUS  |
	//	|                                          | access right had not been granted to the caller when the RPC context handle to   |
	//	|                                          | the service record was created.                                                  |
	//	+------------------------------------------+----------------------------------------------------------------------------------+
	//	| 6 ERROR_INVALID_HANDLE                   | The handle is no longer valid or is not supported for the specified              |
	//	|                                          | notification.                                                                    |
	//	+------------------------------------------+----------------------------------------------------------------------------------+
	//	| 50 ERROR_NOT_SUPPORTED                   | The request is not supported.                                                    |
	//	+------------------------------------------+----------------------------------------------------------------------------------+
	//	| 87 ERROR_INVALID_PARAMETER               | A parameter that was specified is invalid.                                       |
	//	+------------------------------------------+----------------------------------------------------------------------------------+
	//	| 124 ERROR_INVALID_LEVEL                  | The system call level is not correct.                                            |
	//	+------------------------------------------+----------------------------------------------------------------------------------+
	//	| 1072 ERROR_SERVICE_MARKED_FOR_DELETE     | The RDeleteService has been called for the service record identified by the      |
	//	|                                          | hService parameter.                                                              |
	//	+------------------------------------------+----------------------------------------------------------------------------------+
	//	| 1115 ERROR_SHUTDOWN_IN_PROGRESS          | The system is shutting down.                                                     |
	//	+------------------------------------------+----------------------------------------------------------------------------------+
	//	| 1242 ERROR_ALREADY_REGISTERED            | A notification status handle has already been created for the service handle     |
	//	|                                          | passed in the hService parameter.                                                |
	//	+------------------------------------------+----------------------------------------------------------------------------------+
	//	| 1294 ERROR_SERVICE_NOTIFY_CLIENT_LAGGING | The service notification client is lagging too far behind the current state of   |
	//	|                                          | services in the machine.                                                         |
	//	+------------------------------------------+----------------------------------------------------------------------------------+
	NotifyServiceStatusChange(context.Context, *NotifyServiceStatusChangeRequest) (*NotifyServiceStatusChangeResponse, error)

	// The RGetNotifyResults method<75> returns notification information when the specified
	// status change that was previously requested by the client via RNotifyServiceStatusChange
	// (section 3.1.4.43) occurs on a specified service.
	//
	// The client MUST make one call to RGetNotifyResults for each call to RNotifyServiceStatusChange.
	//
	// Return Values: The method returns 0x00000000 (ERROR_SUCCESS) on success; otherwise,
	// it returns one of the following error codes.
	//
	//	+---------------------------------+--------------------------------+
	//	|             RETURN              |                                |
	//	|           VALUE/CODE            |          DESCRIPTION           |
	//	|                                 |                                |
	//	+---------------------------------+--------------------------------+
	//	+---------------------------------+--------------------------------+
	//	| 6 ERROR_INVALID_HANDLE          | The handle is no longer valid. |
	//	+---------------------------------+--------------------------------+
	//	| 1115 ERROR_SHUTDOWN_IN_PROGRESS | The system is shutting down.   |
	//	+---------------------------------+--------------------------------+
	//	| 1235 ERROR_REQUEST_ABORTED      | The request was aborted.       |
	//	+---------------------------------+--------------------------------+
	GetNotifyResults(context.Context, *GetNotifyResultsRequest) (*GetNotifyResultsResponse, error)

	// The RCloseNotifyHandle method<77> unregisters the client from receiving future notifications
	// via the RGetNotifyResults (section 3.1.4.44) method from the server for specified
	// status changes on a specified service.
	//
	// Return Values: The method returns 0x00000000 (ERROR_SUCCESS) on success; otherwise,
	// it returns the following error code.
	//
	//	+------------------------+--------------------------------+
	//	|         RETURN         |                                |
	//	|       VALUE/CODE       |          DESCRIPTION           |
	//	|                        |                                |
	//	+------------------------+--------------------------------+
	//	+------------------------+--------------------------------+
	//	| 6 ERROR_INVALID_HANDLE | The handle is no longer valid. |
	//	+------------------------+--------------------------------+
	CloseNotify(context.Context, *CloseNotifyRequest) (*CloseNotifyResponse, error)

	// The RControlServiceExA method<78> receives a control code for a specific service.
	//
	// Return Values: The method returns 0x00000000 (ERROR_SUCCESS) on success; otherwise,
	// it returns one of the following error codes.
	//
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	|                RETURN                 |                                                                                  |
	//	|              VALUE/CODE               |                                   DESCRIPTION                                    |
	//	|                                       |                                                                                  |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| 5 ERROR_ACCESS_DENIED                 | The required access right had not been granted to the caller when the RPC        |
	//	|                                       | context handle to the service record was created.                                |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| 6 ERROR_INVALID_HANDLE                | The handle is no longer valid.                                                   |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| 87 ERROR_INVALID_PARAMETER            | The requested control code is undefined.                                         |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| 124 ERROR_INVALID_LEVEL               | The dwInfoLevel parameter contains an unsupported value.                         |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| 1051 ERROR_DEPENDENT_SERVICES_RUNNING | The service cannot be stopped because other running services are dependent on    |
	//	|                                       | it.                                                                              |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| 1052 ERROR_INVALID_SERVICE_CONTROL    | The requested control code is not valid, or it is unacceptable to the service.   |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| 1053 ERROR_SERVICE_REQUEST_TIMEOUT    | The process for the service was started, but it did not respond within an        |
	//	|                                       | implementation-specific time-out.<79>                                            |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| 1061 ERROR_SERVICE_CANNOT_ACCEPT_CTRL | The requested control code cannot be sent to the service because the state of    |
	//	|                                       | the service is SERVICE_START_PENDING or SERVICE_STOP_PENDING.                    |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| 1062 ERROR_SERVICE_NOT_ACTIVE         | The service has not been started, or the ServiceStatus.dwCurrentState in the     |
	//	|                                       | service record is SERVICE_STOPPED.                                               |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| 1115 ERROR_SHUTDOWN_IN_PROGRESS       | The system is shutting down.                                                     |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	ControlServiceExA(context.Context, *ControlServiceExARequest) (*ControlServiceExAResponse, error)

	// The RControlServiceExW method<80> receives a control code for a specific service.
	//
	// Return Values: The method returns 0x00000000 (ERROR_SUCCESS) on success; otherwise,
	// it returns one of the following error codes.
	//
	//	+---------------------------------------------+----------------------------------------------------------------------------------+
	//	|                   RETURN                    |                                                                                  |
	//	|                 VALUE/CODE                  |                                   DESCRIPTION                                    |
	//	|                                             |                                                                                  |
	//	+---------------------------------------------+----------------------------------------------------------------------------------+
	//	+---------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000005 ERROR_ACCESS_DENIED              | The required access right had not been granted to the caller when the RPC        |
	//	|                                             | context handle to the service record was created.                                |
	//	+---------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000006 ERROR_INVALID_HANDLE             | The handle is no longer valid.                                                   |
	//	+---------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000087 ERROR_INVALID_PARAMETER          | The requested control code is undefined.                                         |
	//	+---------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000124 ERROR_INVALID_LEVEL              | The dwInfoLevel parameter contains an unsupported level.                         |
	//	+---------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00001051 ERROR_DEPENDENT_SERVICES_RUNNING | The service cannot be stopped because other running services are dependent on    |
	//	|                                             | it.                                                                              |
	//	+---------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00001052 ERROR_INVALID_SERVICE_CONTROL    | The requested control code is not valid, or it is unacceptable to the service.   |
	//	+---------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00001053 ERROR_SERVICE_REQUEST_TIMEOUT    | The process for the service was started, but it did not respond within an        |
	//	|                                             | implementation-specific timeout.<81>                                             |
	//	+---------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00001061 ERROR_SERVICE_CANNOT_ACCEPT_CTRL | The requested control code cannot be sent to the service because the state of    |
	//	|                                             | the service is SERVICE_START_PENDING or SERVICE_STOP_PENDING.                    |
	//	+---------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00001062 ERROR_SERVICE_NOT_ACTIVE         | The service has not been started, or the ServiceStatus.dwCurrentState in the     |
	//	|                                             | service record is SERVICE_STOPPED.                                               |
	//	+---------------------------------------------+----------------------------------------------------------------------------------+
	//	| 1115 ERROR_SHUTDOWN_IN_PROGRESS             | The system is shutting down.                                                     |
	//	+---------------------------------------------+----------------------------------------------------------------------------------+
	ControlServiceExW(context.Context, *ControlServiceExWRequest) (*ControlServiceExWResponse, error)

	// Opnum52NotUsedOnWire operation.
	// Opnum52NotUsedOnWire

	// Opnum53NotUsedOnWire operation.
	// Opnum53NotUsedOnWire

	// Opnum54NotUsedOnWire operation.
	// Opnum54NotUsedOnWire

	// Opnum55NotUsedOnWire operation.
	// Opnum55NotUsedOnWire

	// The RQueryServiceConfigEx  method SHOULD<82> query the optional configuration parameters
	// of a service.
	//
	// Return Values: The method returns 0x00000000 (ERROR_SUCCESS) on success; otherwise,
	// it returns one of the following error codes.
	//
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	|             RETURN              |                                                                                  |
	//	|           VALUE/CODE            |                                   DESCRIPTION                                    |
	//	|                                 |                                                                                  |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	| 5 ERROR_ACCESS_DENIED           | The SERVICE_QUERY_CONFIG access right had not been granted to the caller when    |
	//	|                                 | the RPC context handle was created.                                              |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	| 6 ERROR_INVALID_HANDLE          | The handle is no longer valid.                                                   |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	| 124 ERROR_INVALID_LEVEL         | The dwInfoLevel parameter contains an unsupported value.                         |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	| 1115 ERROR_SHUTDOWN_IN_PROGRESS | The system is shutting down.                                                     |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	QueryServiceConfigEx(context.Context, *QueryServiceConfigExRequest) (*QueryServiceConfigExResponse, error)

	// Opnum57NotUsedOnWire operation.
	// Opnum57NotUsedOnWire

	// Opnum58NotUsedOnWire operation.
	// Opnum58NotUsedOnWire

	// Opnum59NotUsedOnWire operation.
	// Opnum59NotUsedOnWire

	// The RCreateWowService method creates a service whose binary is compiled for a specified
	// computer architecture.<83> The path to the file image is automatically adjusted to
	// point to the correct WoW-redirected location.
	//
	// Return Values: The method returns 0x00000000 (ERROR_SUCCESS) on success; otherwise,
	// it returns one of the following error codes.
	//
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	|                RETURN                |                                                                                  |
	//	|              VALUE/CODE              |                                   DESCRIPTION                                    |
	//	|                                      |                                                                                  |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| 5 ERROR_ACCESS_DENIED                | The SC_MANAGER_CREATE_SERVICE access right had not been granted to the caller    |
	//	|                                      | when the RPC context handle to the SCM was created.                              |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| 6 ERROR_INVALID_HANDLE               | The handle specified is invalid.                                                 |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| 13 ERROR_INVALID_DATA                | The data is invalid.                                                             |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| 50 ERROR_NOT_SUPPORTED               | dwServiceWowType was an architecture that is not supported.                      |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| 87 ERROR_INVALID_PARAMETER           | A parameter that was specified is invalid.                                       |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| 123 ERROR_INVALID_NAME               | The specified service name is invalid.                                           |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| 1057 ERROR_INVALID_SERVICE_ACCOUNT   | The user account name specified in the lpServiceStartName parameter does not     |
	//	|                                      | exist.                                                                           |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| 1059 ERROR_CIRCULAR_DEPENDENCY       | A circular service dependency was specified.                                     |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| 1072 ERROR_SERVICE_MARKED_FOR_DELETE | The service record with a specified name already exists, and RDeleteService has  |
	//	|                                      | been called for it.                                                              |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| 1073 ERROR_SERVICE_EXISTS            | The service record with the ServiceName matching the specified lpServiceName     |
	//	|                                      | already exists.                                                                  |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| 1078 ERROR_DUPLICATE_SERVICE_NAME    | The service record with the same DisplayName or the same ServiceName as the      |
	//	|                                      | passed-in lpDisplayName already exists in the service control manager database.  |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| 1115 ERROR_SHUTDOWN_IN_PROGRESS      | The system is shutting down.                                                     |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	CreateWOWService(context.Context, *CreateWOWServiceRequest) (*CreateWOWServiceResponse, error)

	// The ROpenSCManager2 method establishes a connection to server and opens the SCM database
	// on the specified server.<87>
	//
	// Return Values: The method returns 0x00000000 (ERROR_SUCCESS) on success; otherwise,
	// it returns one of the following error codes.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN               |                                                                                  |
	//	|             VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 5 ERROR_ACCESS_DENIED              | The client does not have the required access rights to open the SCM              |
	//	|                                    | database on the server or the desired access is not granted to it in the SCM     |
	//	|                                    | SecurityDescriptor.                                                              |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 123 ERROR_INVALID_NAME             | The specified service name is invalid.                                           |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 1065 ERROR_DATABASE_DOES_NOT_EXIST | The database specified does not exist.                                           |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 1115 ERROR_SHUTDOWN_IN_PROGRESS    | The system is shutting down.                                                     |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	OpenSCM2(context.Context, *OpenSCM2Request) (*OpenSCM2Response, error)
}

func RegisterSvcctlServer(conn dcerpc.Conn, o SvcctlServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewSvcctlServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(SvcctlSyntaxV2_0))...)
}

func NewSvcctlServerHandle(o SvcctlServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return SvcctlServerHandle(ctx, o, opNum, r)
	}
}

func SvcctlServerHandle(ctx context.Context, o SvcctlServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	switch opNum {
	case 0: // RCloseServiceHandle
		in := &CloseServiceRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.CloseService(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 1: // RControlService
		in := &ControlServiceRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.ControlService(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 2: // RDeleteService
		in := &DeleteServiceRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.DeleteService(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 3: // RLockServiceDatabase
		in := &LockServiceDatabaseRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.LockServiceDatabase(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 4: // RQueryServiceObjectSecurity
		in := &QueryServiceObjectSecurityRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.QueryServiceObjectSecurity(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 5: // RSetServiceObjectSecurity
		in := &SetServiceObjectSecurityRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetServiceObjectSecurity(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 6: // RQueryServiceStatus
		in := &QueryServiceStatusRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.QueryServiceStatus(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 7: // RSetServiceStatus
		in := &SetServiceStatusRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetServiceStatus(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 8: // RUnlockServiceDatabase
		in := &UnlockServiceDatabaseRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.UnlockServiceDatabase(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 9: // RNotifyBootConfigStatus
		in := &NotifyBootConfigStatusRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.NotifyBootConfigStatus(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 10: // Opnum10NotUsedOnWire
		// Opnum10NotUsedOnWire
		return nil, nil
	case 11: // RChangeServiceConfigW
		in := &ChangeServiceConfigWRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.ChangeServiceConfigW(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 12: // RCreateServiceW
		in := &CreateServiceWRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.CreateServiceW(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 13: // REnumDependentServicesW
		in := &EnumDependentServicesWRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.EnumDependentServicesW(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 14: // REnumServicesStatusW
		in := &EnumServicesStatusWRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.EnumServicesStatusW(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 15: // ROpenSCManagerW
		in := &OpenSCMWRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.OpenSCMW(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 16: // ROpenServiceW
		in := &OpenServiceWRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.OpenServiceW(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 17: // RQueryServiceConfigW
		in := &QueryServiceConfigWRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.QueryServiceConfigW(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 18: // RQueryServiceLockStatusW
		in := &QueryServiceLockStatusWRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.QueryServiceLockStatusW(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 19: // RStartServiceW
		in := &StartServiceWRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.StartServiceW(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 20: // RGetServiceDisplayNameW
		in := &GetServiceDisplayNameWRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetServiceDisplayNameW(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 21: // RGetServiceKeyNameW
		in := &GetServiceKeyNameWRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetServiceKeyNameW(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 22: // Opnum22NotUsedOnWire
		// Opnum22NotUsedOnWire
		return nil, nil
	case 23: // RChangeServiceConfigA
		in := &ChangeServiceConfigARequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.ChangeServiceConfigA(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 24: // RCreateServiceA
		in := &CreateServiceARequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.CreateServiceA(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 25: // REnumDependentServicesA
		in := &EnumDependentServicesARequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.EnumDependentServicesA(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 26: // REnumServicesStatusA
		in := &EnumServicesStatusARequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.EnumServicesStatusA(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 27: // ROpenSCManagerA
		in := &OpenSCMARequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.OpenSCMA(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 28: // ROpenServiceA
		in := &OpenServiceARequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.OpenServiceA(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 29: // RQueryServiceConfigA
		in := &QueryServiceConfigARequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.QueryServiceConfigA(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 30: // RQueryServiceLockStatusA
		in := &QueryServiceLockStatusARequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.QueryServiceLockStatusA(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 31: // RStartServiceA
		in := &StartServiceARequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.StartServiceA(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 32: // RGetServiceDisplayNameA
		in := &GetServiceDisplayNameARequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetServiceDisplayNameA(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 33: // RGetServiceKeyNameA
		in := &GetServiceKeyNameARequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetServiceKeyNameA(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 34: // Opnum34NotUsedOnWire
		// Opnum34NotUsedOnWire
		return nil, nil
	case 35: // REnumServiceGroupW
		in := &EnumServiceGroupWRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.EnumServiceGroupW(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 36: // RChangeServiceConfig2A
		in := &ChangeServiceConfig2ARequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.ChangeServiceConfig2A(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 37: // RChangeServiceConfig2W
		in := &ChangeServiceConfig2WRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.ChangeServiceConfig2W(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 38: // RQueryServiceConfig2A
		in := &QueryServiceConfig2ARequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.QueryServiceConfig2A(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 39: // RQueryServiceConfig2W
		in := &QueryServiceConfig2WRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.QueryServiceConfig2W(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 40: // RQueryServiceStatusEx
		in := &QueryServiceStatusExRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.QueryServiceStatusEx(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 41: // REnumServicesStatusExA
		in := &EnumServicesStatusExARequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.EnumServicesStatusExA(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 42: // REnumServicesStatusExW
		in := &EnumServicesStatusExWRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.EnumServicesStatusExW(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 43: // Opnum43NotUsedOnWire
		// Opnum43NotUsedOnWire
		return nil, nil
	case 44: // RCreateServiceWOW64A
		in := &CreateServiceWOW64ARequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.CreateServiceWOW64A(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 45: // RCreateServiceWOW64W
		in := &CreateServiceWOW64WRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.CreateServiceWOW64W(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 46: // Opnum46NotUsedOnWire
		// Opnum46NotUsedOnWire
		return nil, nil
	case 47: // RNotifyServiceStatusChange
		in := &NotifyServiceStatusChangeRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.NotifyServiceStatusChange(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 48: // RGetNotifyResults
		in := &GetNotifyResultsRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetNotifyResults(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 49: // RCloseNotifyHandle
		in := &CloseNotifyRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.CloseNotify(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 50: // RControlServiceExA
		in := &ControlServiceExARequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.ControlServiceExA(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 51: // RControlServiceExW
		in := &ControlServiceExWRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.ControlServiceExW(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 52: // Opnum52NotUsedOnWire
		// Opnum52NotUsedOnWire
		return nil, nil
	case 53: // Opnum53NotUsedOnWire
		// Opnum53NotUsedOnWire
		return nil, nil
	case 54: // Opnum54NotUsedOnWire
		// Opnum54NotUsedOnWire
		return nil, nil
	case 55: // Opnum55NotUsedOnWire
		// Opnum55NotUsedOnWire
		return nil, nil
	case 56: // RQueryServiceConfigEx
		in := &QueryServiceConfigExRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.QueryServiceConfigEx(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 57: // Opnum57NotUsedOnWire
		// Opnum57NotUsedOnWire
		return nil, nil
	case 58: // Opnum58NotUsedOnWire
		// Opnum58NotUsedOnWire
		return nil, nil
	case 59: // Opnum59NotUsedOnWire
		// Opnum59NotUsedOnWire
		return nil, nil
	case 60: // RCreateWowService
		in := &CreateWOWServiceRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.CreateWOWService(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 61: // ROpenSCManager2
		in := &OpenSCM2Request{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.OpenSCM2(ctx, in)
		return resp.xxx_ToOp(ctx), err
	}
	return nil, nil
}
