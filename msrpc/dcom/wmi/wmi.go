// The wmi package implements the WMI client protocol.
//
// # Introduction
//
// Windows Management Instrumentation (WMI) Remote Protocol is a Distributed Component
// Object Model (DCOM), as specified in [MS-DCOM], a client/server–based framework
// that provides an open and automated means of systems management. WMI leverages the
// Common Information Model (CIM), as specified in [DMTF-DSP0004], to represent various
// components of the operating system. CIM is the conceptual model for storing enterprise
// management information. The information available from CIM is specified by a series
// of classes and associations, and the elements contained in them (methods, properties,
// and references). These constructs describe the data available to WMI clients.
//
// # Overview
//
// The Windows Management Instrumentation (WMI) Remote Protocol is used to communicate
// management data conforming to Common Information Model (CIM), as specified in [DMTF-DSP0004].
// The Windows Management Instrumentation Remote Protocol uses CIM as the conceptual
// model for representing enterprise management information that can be managed by an
// administrator. However WMI is not fully compliant with [DMTF-DSP0004]. The exceptions
// are documented where applicable in the WMI Remote Protocol.
//
// The Windows Management Instrumentation Remote Protocol is implemented as a three-tier
// architecture, as illustrated in the following figure.
package wmi

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	dcom "github.com/oiweiwei/go-msrpc/msrpc/dcom"
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
	_ = dcom.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/wmi"
)

// QueryFlagType type represents WBEM_QUERY_FLAG_TYPE RPC enumeration.
//
// The WBEM_QUERY_FLAG_TYPE enumeration gives information about the type of the flag.
type QueryFlagType uint32

var (
	// WBEM_FLAG_DEEP:  If used in IWbemServices::CreateClassEnum or IWbemServices::CreateClassEnumAsync,
	// the WBEM_FLAG_DEEP constant causes the enumeration to return all the subclasses in
	// the hierarchy of a specified class but to not return the class itself.
	QueryFlagTypeDeep QueryFlagType = 0
	// WBEM_FLAG_SHALLOW:  If used in IWbemServices::CreateClassEnum or IWbemServices::CreateClassEnumAsync,
	// the WBEM_FLAG_SHALLOW constant causes the enumeration to return the immediate subclasses
	// of a specified class.
	QueryFlagTypeShallow QueryFlagType = 1
	// WBEM_FLAG_PROTOTYPE:  This flag is used for prototyping. It does not run the query;
	// instead, it returns the Prototype Result Object as specified in section 2.2.4.1.
	QueryFlagTypePrototype QueryFlagType = 2
)

func (o QueryFlagType) String() string {
	switch o {
	case QueryFlagTypeDeep:
		return "QueryFlagTypeDeep"
	case QueryFlagTypeShallow:
		return "QueryFlagTypeShallow"
	case QueryFlagTypePrototype:
		return "QueryFlagTypePrototype"
	}
	return "Invalid"
}

// ChangeFlagType type represents WBEM_CHANGE_FLAG_TYPE RPC enumeration.
//
// The WBEM_CHANGE_FLAG_TYPE enumeration is used to indicate and update the type of
// the flag.
type ChangeFlagType uint32

var (
	// WBEM_FLAG_CREATE_OR_UPDATE:  This flag causes the put operation to update the class
	// or instance if it does not exist, or to overwrite the class or instance if it exists
	// already.
	ChangeFlagTypeCreateOrUpdate ChangeFlagType = 0
	// WBEM_FLAG_UPDATE_ONLY:  This flag causes the put operation to update the class or
	// instance. The class or instance MUST exist for the call to be successful.
	ChangeFlagTypeUpdateOnly ChangeFlagType = 1
	// WBEM_FLAG_CREATE_ONLY:  This flag causes the put operation to create the class or
	// instance. For the call to be successful, the class or instance MUST NOT exist.
	ChangeFlagTypeCreateOnly ChangeFlagType = 2
	// WBEM_FLAG_UPDATE_SAFE_MODE:  This flag allows updates of classes even if there are
	// child classes, as long as the change does not cause any conflicts with child classes.
	// An example of an update that this flag allows is the adding of a new property to
	// the base class that was not previously mentioned in any of the child classes. If
	// the class has instances, the update fails.
	ChangeFlagTypeUpdateSafeMode ChangeFlagType = 32
	// WBEM_FLAG_UPDATE_FORCE_MODE:  This flag forces updates of classes when conflicting
	// child classes exist. An example of an update that this flag forces is when a class
	// qualifier is defined in a child class and the base class tries to add the same qualifier
	// that conflicted with the existing one. In force mode, this conflict is resolved by
	// deleting the conflicting qualifier in the child class.
	ChangeFlagTypeUpdateForceMode ChangeFlagType = 64
)

func (o ChangeFlagType) String() string {
	switch o {
	case ChangeFlagTypeCreateOrUpdate:
		return "ChangeFlagTypeCreateOrUpdate"
	case ChangeFlagTypeUpdateOnly:
		return "ChangeFlagTypeUpdateOnly"
	case ChangeFlagTypeCreateOnly:
		return "ChangeFlagTypeCreateOnly"
	case ChangeFlagTypeUpdateSafeMode:
		return "ChangeFlagTypeUpdateSafeMode"
	case ChangeFlagTypeUpdateForceMode:
		return "ChangeFlagTypeUpdateForceMode"
	}
	return "Invalid"
}

// ConnectOptions type represents WBEM_CONNECT_OPTIONS RPC enumeration.
//
// The WBEM_CONNECT_OPTIONS enumeration gives information about the type of options
// of the connection.
type ConnectOptions uint32

var (
	// WBEM_FLAG_CONNECT_REPOSITORY_ONLY:  Reserved for local use.
	ConnectOptionsFlagConnectRepositoryOnly ConnectOptions = 64
	// WBEM_FLAG_CONNECT_PROVIDERS:  Reserved for local use.<2>
	ConnectOptionsFlagConnectProviders ConnectOptions = 256
)

func (o ConnectOptions) String() string {
	switch o {
	case ConnectOptionsFlagConnectRepositoryOnly:
		return "ConnectOptionsFlagConnectRepositoryOnly"
	case ConnectOptionsFlagConnectProviders:
		return "ConnectOptionsFlagConnectProviders"
	}
	return "Invalid"
}

// GenericFlagType type represents WBEM_GENERIC_FLAG_TYPE RPC enumeration.
//
// The WBEM_GENERIC_FLAG_TYPE enumeration is used to indicate and update the type of
// the flag.
type GenericFlagType uint32

var (
	// WBEM_FLAG_RETURN_WBEM_COMPLETE:  This flag makes the operation synchronous. This
	// is the default behavior and so this flag need not be explicitly specified.
	GenericFlagTypeReturnWBEMComplete GenericFlagType = 0
	// WBEM_FLAG_RETURN_IMMEDIATELY:  This flag causes the call to return without waiting
	// for the operation to complete. The call result parameter contains the IWbemCallResult
	// object by using the status of the operation that can be retrieved.
	GenericFlagTypeReturnImmediately GenericFlagType = 16
	// WBEM_FLAG_FORWARD_ONLY:  This flag causes a forward-only enumerator, IEnumWbemClassObject,
	// (section 3.1.4.4), to be returned. Forward-only enumerators are typically much faster
	// and use less memory than conventional enumerators; however, they do not allow calls
	// to IEnumWbemClassObject::Clone or IEnumWbemClassObject::Reset.
	GenericFlagTypeForwardOnly GenericFlagType = 32
	// WBEM_FLAG_NO_ERROR_OBJECT:  This flag MUST NOT be set, and MUST be ignored on receipt.
	GenericFlagTypeNoErrorObject GenericFlagType = 64
	// WBEM_FLAG_SEND_STATUS:  This flag registers a request with WMI to receive intermediate
	// status reports through the client implementation of IWbemObjectSink::SetStatus, if
	// supported by the server implementation.
	GenericFlagTypeSendStatus GenericFlagType = 128
	// WBEM_FLAG_ENSURE_LOCATABLE:  This flag ensures that any returned objects have enough
	// information in them so that system properties, such as __PATH, __RELPATH, and __SERVER,<1>
	// are non-NULL.
	GenericFlagTypeEnsureLocatable GenericFlagType = 256
	// WBEM_FLAG_DIRECT_READ:  This flag causes direct access to the specified class without
	// regard to its parent class or subclasses.
	GenericFlagTypeDirectRead GenericFlagType = 512
	// WBEM_MASK_RESERVED_FLAGS:  This flag MUST NOT be set, and MUST be ignored on receipt.
	GenericFlagTypeMaskReservedFlags GenericFlagType = 126976
	// WBEM_FLAG_USE_AMENDED_QUALIFIERS:  If this flag is set, the server retrieves any
	// qualifiers in the CIM object that can be localized in the current connection's locale.
	// The set of localized qualifiers and the list of locales for which the qualifier is
	// localized are implementation dependent. When the localized information is available,
	// the server retrieves the localized values using the client-preferred locale. If the
	// localized values are not available, the server returns values using the default locale.
	//
	// The localized qualifiers or amended qualifiers are identified by the qualifier flavor
	// as defined in [MS-WMIO] section 2.2.62.
	GenericFlagTypeUseAmendedQualifiers GenericFlagType = 131072
	// WBEM_FLAG_STRONG_VALIDATION:  This flag MUST NOT be set, and MUST be ignored on
	// receipt.
	GenericFlagTypeStrongValidation GenericFlagType = 1048576
)

func (o GenericFlagType) String() string {
	switch o {
	case GenericFlagTypeReturnWBEMComplete:
		return "GenericFlagTypeReturnWBEMComplete"
	case GenericFlagTypeReturnImmediately:
		return "GenericFlagTypeReturnImmediately"
	case GenericFlagTypeForwardOnly:
		return "GenericFlagTypeForwardOnly"
	case GenericFlagTypeNoErrorObject:
		return "GenericFlagTypeNoErrorObject"
	case GenericFlagTypeSendStatus:
		return "GenericFlagTypeSendStatus"
	case GenericFlagTypeEnsureLocatable:
		return "GenericFlagTypeEnsureLocatable"
	case GenericFlagTypeDirectRead:
		return "GenericFlagTypeDirectRead"
	case GenericFlagTypeMaskReservedFlags:
		return "GenericFlagTypeMaskReservedFlags"
	case GenericFlagTypeUseAmendedQualifiers:
		return "GenericFlagTypeUseAmendedQualifiers"
	case GenericFlagTypeStrongValidation:
		return "GenericFlagTypeStrongValidation"
	}
	return "Invalid"
}

// StatusType type represents WBEM_STATUS_TYPE RPC enumeration.
//
// The WBEM_STATUS_TYPE enumeration gives information about the status of the operation.
type StatusType uint16

var (
	// WBEM_STATUS_COMPLETE:  When the WMI operation is completed, WMI calls IWbemObjectSink::SetStatus
	// with WBEM_STATUS_COMPLETE.
	StatusTypeComplete StatusType = 0
	// WBEM_STATUS_REQUIREMENTS:  This flag MUST NOT be set, and MUST be ignored on receipt.
	StatusTypeRequirements StatusType = 1
	// WBEM_STATUS_PROGRESS:  WMI reports the progress of the operation to IWbemObjectSink::SetStatus
	// with flag WBEM_STATUS_PROGRESS.
	StatusTypeProgress StatusType = 2
)

func (o StatusType) String() string {
	switch o {
	case StatusTypeComplete:
		return "StatusTypeComplete"
	case StatusTypeRequirements:
		return "StatusTypeRequirements"
	case StatusTypeProgress:
		return "StatusTypeProgress"
	}
	return "Invalid"
}

// TimeoutType type represents WBEM_TIMEOUT_TYPE RPC enumeration.
//
// The WBEM_TIMEOUT_TYPE enumeration gives information about the type of time-out for
// the process.
type TimeoutType uint32

var (
	// WBEM_NO_WAIT:  If passed as a time-out parameter to the IEnumWbemClassObject::Next
	// method, the call returns the available objects, if any, at the time of the call;
	// it does not wait for any more objects.
	TimeoutTypeNoWait TimeoutType = 0
	// WBEM_INFINITE:  If passed as a time-out parameter to IEnumWbemClassObject::Next,
	// the call blocks until objects are available.
	TimeoutTypeInfinite TimeoutType = 4294967295
)

func (o TimeoutType) String() string {
	switch o {
	case TimeoutTypeNoWait:
		return "TimeoutTypeNoWait"
	case TimeoutTypeInfinite:
		return "TimeoutTypeInfinite"
	}
	return "Invalid"
}

// BackupRestoreFlags type represents WBEM_BACKUP_RESTORE_FLAGS RPC enumeration.
//
// The WBEM_BACKUP_RESTORE_FLAGS enumeration gives information about the backup and
// restore state of the process.
type BackupRestoreFlags uint32

var (
	// WBEM_FLAG_BACKUP_RESTORE_FORCE_SHUTDOWN:  While the CIM database is being restored,
	// any clients connected to WMI are forcibly disconnected.
	BackupRestoreFlagsFlagBackupRestoreForceShutdown BackupRestoreFlags = 1
)

func (o BackupRestoreFlags) String() string {
	switch o {
	case BackupRestoreFlagsFlagBackupRestoreForceShutdown:
		return "BackupRestoreFlagsFlagBackupRestoreForceShutdown"
	}
	return "Invalid"
}

// Status type represents WBEMSTATUS RPC enumeration.
//
// The WBEMSTATUS enumeration gives information about the status of an operation. If
// the server encounters an error condition for which this protocol does not explicitly
// state an error value, the server can return any HRESULT to indicate failure by setting
// the Severity (S bit) of the HRESULT, as defined in [MS-ERREF] section 2.1.
//
// The statuses of operations that are not explicitly called out in this document but
// are part of the associated IDL are deemed to be local-only and are implementation-specific.
type Status uint32

var (
	// WBEM_S_NO_ERROR:  The operation completed successfully.
	StatusNoError Status = 0
	// WBEM_S_FALSE:  Either no more CIM objects are available, the number of returned
	// CIM objects is less than the number requested, or this is the end of an enumeration.
	// This error code is returned from the IEnumWbemClassObject and IWbemWCOSmartEnum interface
	// methods.
	StatusFalse Status = 1
	// WBEM_S_TIMEDOUT:  The attempt to establish the connection has expired.
	StatusTimedout Status = 262148
	// WBEM_S_NEW_STYLE:  The server supports ObjectArray encoding; see section 3.1.4.2.1
	// for details.
	StatusNewStyle Status = 262399
	// WBEM_S_PARTIAL_RESULTS:  The server could not return all the objects and/or properties
	// requested.
	StatusPartialResults Status = 262160
	// WBEM_E_FAILED:  The server has encountered an unknown error while processing the
	// client's request.
	StatusFailed Status = 2147749889
	// WBEM_E_NOT_FOUND:  The object specified in the path does not exist.
	StatusNotFound Status = 2147749890
	// WBEM_E_ACCESS_DENIED:  The permission required to perform the operation is not helped
	// by the security principal performing the operation.
	StatusAccessDenied Status = 2147749891
	// WBEM_E_PROVIDER_FAILURE:  The server has encountered an unknown error while processing
	// the client's request.
	StatusProviderFailure Status = 2147749892
	// WBEM_E_TYPE_MISMATCH:  The server has found an incorrect data type associated with
	// property or input parameter in client's request.
	StatusTypeMismatch Status = 2147749893
	// WBEM_E_OUT_OF_MEMORY:  The server ran out of memory before completing the operation.
	StatusOutOfMemory Status = 2147749894
	// WBEM_E_INVALID_CONTEXT:  The IWbemContext object sent as part of client's request
	// does not contain the required properties.
	StatusInvalidContext Status = 2147749895
	// WBEM_E_INVALID_PARAMETER:  One or more of the parameters passed to the method is
	// not valid. Methods return this error in any of the following circumstances: (1) a
	// parameter is NULL where a non-NULL value is required, (2) the flags specified in
	// the lFlags parameter are not allowed in this method.
	StatusInvalidParameter Status = 2147749896
	// WBEM_E_NOT_AVAILABLE:  The resource is unavailable.
	StatusNotAvailable Status = 2147749897
	// WBEM_E_CRITICAL_ERROR :  The server has encountered a catastrophic failure and cannot
	// process any client's request.
	StatusCriticalError Status = 2147749898
	// WBEM_E_NOT_SUPPORTED:  The attempted operation is not supported.
	StatusNotSupported Status = 2147749900
	// WBEM_E_PROVIDER_NOT_FOUND:  The server has encountered an implementation-specific
	// error.
	StatusProviderNotFound Status = 2147749905
	// WBEM_E_INVALID_PROVIDER_REGISTRATION:  The server has encountered an implementation-specific
	// error.
	StatusInvalidProviderRegistration Status = 2147749906
	// WBEM_E_PROVIDER_LOAD_FAILURE:  The server has encountered an implementation-specific
	// error.
	StatusProviderLoadFailure Status = 2147749907
	// WBEM_E_INITIALIZATION_FAILURE:  The server has encountered failure during its initialization.
	StatusInitializationFailure Status = 2147749908
	// WBEM_E_TRANSPORT_FAILURE:  There is a network problem detected in reaching the server.
	StatusTransportFailure Status = 2147749909
	// WBEM_E_INVALID_OPERATION:  The operation performed is not valid.
	StatusInvalidOperation Status = 2147749910
	// WBEM_E_ALREADY_EXISTS:  When a Put method is called for a CIM object with the flag
	// WBEM_FLAG_CREATE_ONLY and the object already exists, WBEM_E_ALREADY_EXISTS is returned.
	StatusAlreadyExists Status = 2147749913
	// WBEM_E_UNEXPECTED:  An unspecified error has occurred.
	StatusUnexpected Status = 2147749917
	// WBEM_E_INCOMPLETE_CLASS:  The object passed doesn't correspond to any of classes
	// registered with WMI.
	StatusIncompleteClass Status = 2147749920
	// WBEM_E_SHUTTING_DOWN:  The server cannot process the requested operation as it is
	// shutting down.
	StatusShuttingDown Status = 2147749939
	// E_NOTIMPL:  The attempted operation is not implemented. The value of this element
	// is as specified in [MS-ERREF] section 2.1.
	StatusNotimpl Status = 2147500033
	// WBEM_E_INVALID_SUPERCLASS:  When putting a class, the server did not find the parent
	// class specified for the new class to be added.
	StatusInvalidSuperclass Status = 2147749901
	// WBEM_E_INVALID_NAMESPACE:  When connecting to WMI, the namespace specified is not
	// found.
	StatusInvalidNamespace Status = 2147749902
	// WBEM_E_INVALID_OBJECT:  The CIM instance passed to the server doesn't have required
	// information.
	StatusInvalidObject Status = 2147749903
	// WBEM_E_INVALID_CLASS:  The class name is invalid.
	StatusInvalidClass Status = 2147749904
	// WBEM_E_INVALID_QUERY:  The query sent to the server doesn't semantically conform
	// to the rules specified in section 2.2.1.
	StatusInvalidQuery Status = 2147749911
	// WBEM_E_INVALID_QUERY_TYPE:  The query language specified is invalid.
	StatusInvalidQueryType Status = 2147749912
	// WBEM_E_PROVIDER_NOT_CAPABLE:  The server does not support the requested operation
	// on the given CIM class.
	StatusProviderNotCapable Status = 2147749924
	// WBEM_E_CLASS_HAS_CHILDREN:  The class cannot be updated because it has derived classes.
	StatusClassHasChildren Status = 2147749925
	// WBEM_E_CLASS_HAS_INSTANCES:  The class cannot be updated because it has instances.
	StatusClassHasInstances Status = 2147749926
	// WBEM_E_ILLEGAL_NULL:  The server identifies that one of the non-nullable NULL properties
	// was set to NULL in the Put operation.
	StatusIllegalNull Status = 2147749928
	// WBEM_E_INVALID_CIM_TYPE:  The CIM type specified is not valid.
	StatusInvalidCimType Status = 2147749933
	// WBEM_E_INVALID_METHOD:  The CIM object does not implement the specified method.
	StatusInvalidMethod Status = 2147749934
	// WBEM_E_INVALID_METHOD_PARAMETERS:  One or more of the parameters passed to the CIM
	// method are not valid.
	StatusInvalidMethodParameters Status = 2147749935
	// WBEM_E_INVALID_PROPERTY:  The property for which the operation is made is no longer
	// present in the CIM database.
	StatusInvalidProperty Status = 2147749937
	// WBEM_E_CALL_CANCELLED:  The server canceled the execution of the request due to
	// resource constraints. The client can try the call again.
	StatusCallCancelled Status = 2147749938
	// WBEM_E_INVALID_OBJECT_PATH:  The object path is not syntactically valid.
	StatusInvalidObjectPath Status = 2147749946
	// WBEM_E_OUT_OF_DISK_SPACE:  Insufficient resources on the server to satisfy the client's
	// request.
	StatusOutOfDiskSpace Status = 2147749947
	// WBEM_E_UNSUPPORTED_PUT_EXTENSION:  The server has encountered an implementation-specific
	// error.
	StatusUnsupportedPutExtension Status = 2147749949
	// WBEM_E_QUOTA_VIOLATION:  Quota violation.
	StatusQuotaViolation Status = 2147749996
	// WBEM_E_SERVER_TOO_BUSY:  The server cannot complete the operation at this point.
	StatusServerTooBusy Status = 2147749957
	// WBEM_E_METHOD_NOT_IMPLEMENTED:  An attempt was made to execute a method not marked
	// with "implemented" in this class or any of its derived classes.
	StatusMethodNotImplemented Status = 2147749973
	// WBEM_E_METHOD_DISABLED:  An attempt was made to execute a method marked with "disabled"
	// qualifier in MOF.
	StatusMethodDisabled Status = 2147749974
	// WBEM_E_UNPARSABLE_QUERY:  The query sent to the server doesn't syntactically conform
	// to the rules specified in section 2.2.1.
	StatusUnparsableQuery Status = 2147749976
	// WBEM_E_NOT_EVENT_CLASS:  The FROM clause of WQL Event Query (section 2.2.1.2) represents
	// a class that is not derived from Event.
	StatusNotEventClass Status = 2147749977
	// WBEM_E_MISSING_GROUP_WITHIN:  The GROUP BY clause of WQL query does not have WITHIN
	// specified.
	StatusMissingGroupWithin Status = 2147749978
	// WBEM_E_MISSING_AGGREGATION_LIST:  The GROUP BY clause was used with aggregation,
	// which is not supported.
	StatusMissingAggregationList Status = 2147749979
	// WBEM_E_PROPERTY_NOT_AN_OBJECT:  The GROUP BY clause references an object that is
	// an embedded object without using Dot notation.
	StatusPropertyNotAnObject Status = 2147749980
	// WBEM_E_AGGREGATING_BY_OBJECT:  The GROUP BY clause references an object that is
	// an embedded object without using Dot notation.
	StatusAggregatingByObject Status = 2147749981
	// WBEM_E_BACKUP_RESTORE_WINMGMT_RUNNING:  A request for backing up or restoring the
	// CIM database was sent while the server was using it.
	StatusBackupRestoreWinManagementRunning Status = 2147749984
	// WBEM_E_QUEUE_OVERFLOW:  The EventQueue on the server has more events than can be
	// consumed by the client.
	StatusQueueOverflow Status = 2147749985
	// WBEM_E_PRIVILEGE_NOT_HELD:  The server could not find the required privilege for
	// performing operations on CIM classes or CIM instances.
	StatusPrivilegeNotHeld Status = 2147749986
	// WBEM_E_INVALID_OPERATOR:  An operator in the WQL query is invalid for this property
	// type.
	StatusInvalidOperator Status = 2147749987
	// WBEM_E_CANNOT_BE_ABSTRACT:  The CIM class on the server had the abstract qualifier
	// set to true, while its parent class does not have the abstract qualifier set to false.
	StatusCannotBeAbstract Status = 2147749989
	// WBEM_E_AMENDED_OBJECT:  A CIM instance with amended qualifier set to true is being
	// updated without WBEM_FLAG_USE_AMENDED_QUALIFIERS flag.
	StatusAmendedObject Status = 2147749990
	// WBEM_E_VETO_PUT:  The server cannot perform a PUT operation because it is not supported
	// for the given CIM class.
	StatusVetoPut Status = 2147750010
	// WBEM_E_PROVIDER_SUSPENDED:  The server has encountered an implementation-specific
	// error.
	StatusProviderSuspended Status = 2147750017
	// WBEM_E_ENCRYPTED_CONNECTION_REQUIRED:  The server has encountered an implementation-specific
	// error.
	StatusEncryptedConnectionRequired Status = 2147750023
	// WBEM_E_PROVIDER_TIMED_OUT:
	StatusProviderTimedOut Status = 2147750024
	// WBEM_E_NO_KEY:  The IWbemServices::PuInstance or IWbemServices::PutInstanceAsync
	// operation was attempted with no value set for the key properties.
	StatusNoKey Status = 2147750025
	// WBEM_E_PROVIDER_DISABLED:  The server has encountered an implementation-specific
	// error.
	StatusProviderDisabled Status = 2147750026
	// WBEM_E_REGISTRATION_TOO_BROAD:  The server has encountered an implementation-specific
	// error.
	StatusRegistrationTooBroad Status = 2147753985
	// WBEM_E_REGISTRATION_TOO_PRECISE:  The WQL query for intrinsic events for a class
	// issued without a WITHIN clause.
	StatusRegistrationTooPrecise Status = 2147753986
)

func (o Status) String() string {
	switch o {
	case StatusNoError:
		return "StatusNoError"
	case StatusFalse:
		return "StatusFalse"
	case StatusTimedout:
		return "StatusTimedout"
	case StatusNewStyle:
		return "StatusNewStyle"
	case StatusPartialResults:
		return "StatusPartialResults"
	case StatusFailed:
		return "StatusFailed"
	case StatusNotFound:
		return "StatusNotFound"
	case StatusAccessDenied:
		return "StatusAccessDenied"
	case StatusProviderFailure:
		return "StatusProviderFailure"
	case StatusTypeMismatch:
		return "StatusTypeMismatch"
	case StatusOutOfMemory:
		return "StatusOutOfMemory"
	case StatusInvalidContext:
		return "StatusInvalidContext"
	case StatusInvalidParameter:
		return "StatusInvalidParameter"
	case StatusNotAvailable:
		return "StatusNotAvailable"
	case StatusCriticalError:
		return "StatusCriticalError"
	case StatusNotSupported:
		return "StatusNotSupported"
	case StatusProviderNotFound:
		return "StatusProviderNotFound"
	case StatusInvalidProviderRegistration:
		return "StatusInvalidProviderRegistration"
	case StatusProviderLoadFailure:
		return "StatusProviderLoadFailure"
	case StatusInitializationFailure:
		return "StatusInitializationFailure"
	case StatusTransportFailure:
		return "StatusTransportFailure"
	case StatusInvalidOperation:
		return "StatusInvalidOperation"
	case StatusAlreadyExists:
		return "StatusAlreadyExists"
	case StatusUnexpected:
		return "StatusUnexpected"
	case StatusIncompleteClass:
		return "StatusIncompleteClass"
	case StatusShuttingDown:
		return "StatusShuttingDown"
	case StatusNotimpl:
		return "StatusNotimpl"
	case StatusInvalidSuperclass:
		return "StatusInvalidSuperclass"
	case StatusInvalidNamespace:
		return "StatusInvalidNamespace"
	case StatusInvalidObject:
		return "StatusInvalidObject"
	case StatusInvalidClass:
		return "StatusInvalidClass"
	case StatusInvalidQuery:
		return "StatusInvalidQuery"
	case StatusInvalidQueryType:
		return "StatusInvalidQueryType"
	case StatusProviderNotCapable:
		return "StatusProviderNotCapable"
	case StatusClassHasChildren:
		return "StatusClassHasChildren"
	case StatusClassHasInstances:
		return "StatusClassHasInstances"
	case StatusIllegalNull:
		return "StatusIllegalNull"
	case StatusInvalidCimType:
		return "StatusInvalidCimType"
	case StatusInvalidMethod:
		return "StatusInvalidMethod"
	case StatusInvalidMethodParameters:
		return "StatusInvalidMethodParameters"
	case StatusInvalidProperty:
		return "StatusInvalidProperty"
	case StatusCallCancelled:
		return "StatusCallCancelled"
	case StatusInvalidObjectPath:
		return "StatusInvalidObjectPath"
	case StatusOutOfDiskSpace:
		return "StatusOutOfDiskSpace"
	case StatusUnsupportedPutExtension:
		return "StatusUnsupportedPutExtension"
	case StatusQuotaViolation:
		return "StatusQuotaViolation"
	case StatusServerTooBusy:
		return "StatusServerTooBusy"
	case StatusMethodNotImplemented:
		return "StatusMethodNotImplemented"
	case StatusMethodDisabled:
		return "StatusMethodDisabled"
	case StatusUnparsableQuery:
		return "StatusUnparsableQuery"
	case StatusNotEventClass:
		return "StatusNotEventClass"
	case StatusMissingGroupWithin:
		return "StatusMissingGroupWithin"
	case StatusMissingAggregationList:
		return "StatusMissingAggregationList"
	case StatusPropertyNotAnObject:
		return "StatusPropertyNotAnObject"
	case StatusAggregatingByObject:
		return "StatusAggregatingByObject"
	case StatusBackupRestoreWinManagementRunning:
		return "StatusBackupRestoreWinManagementRunning"
	case StatusQueueOverflow:
		return "StatusQueueOverflow"
	case StatusPrivilegeNotHeld:
		return "StatusPrivilegeNotHeld"
	case StatusInvalidOperator:
		return "StatusInvalidOperator"
	case StatusCannotBeAbstract:
		return "StatusCannotBeAbstract"
	case StatusAmendedObject:
		return "StatusAmendedObject"
	case StatusVetoPut:
		return "StatusVetoPut"
	case StatusProviderSuspended:
		return "StatusProviderSuspended"
	case StatusEncryptedConnectionRequired:
		return "StatusEncryptedConnectionRequired"
	case StatusProviderTimedOut:
		return "StatusProviderTimedOut"
	case StatusNoKey:
		return "StatusNoKey"
	case StatusProviderDisabled:
		return "StatusProviderDisabled"
	case StatusRegistrationTooBroad:
		return "StatusRegistrationTooBroad"
	case StatusRegistrationTooPrecise:
		return "StatusRegistrationTooPrecise"
	}
	return "Invalid"
}

// RefresherVersionNumber type represents WBEM_REFR_VERSION_NUMBER RPC enumeration.
type RefresherVersionNumber uint16

var (
	RefresherVersionNumberVersion RefresherVersionNumber = 2
)

func (o RefresherVersionNumber) String() string {
	switch o {
	case RefresherVersionNumberVersion:
		return "RefresherVersionNumberVersion"
	}
	return "Invalid"
}

// InstanceBlobType type represents WBEM_INSTANCE_BLOB_TYPE RPC enumeration.
//
// The WBEM_INSTANCE_BLOB_TYPE enumeration is used to indicate the type of a CIM object.
type InstanceBlobType uint32

var (
	// WBEM_BLOB_TYPE_ALL:  The object is a single CIM object.
	InstanceBlobTypeAll InstanceBlobType = 2
	// WBEM_BLOB_TYPE_ERROR:  Represents an error condition. In this case the object is
	// NULL.
	InstanceBlobTypeError InstanceBlobType = 3
	// WBEM_BLOB_TYPE_ENUM:  The object is an enumeration of objects of a specific CIM
	// type.
	InstanceBlobTypeEnum InstanceBlobType = 4
)

func (o InstanceBlobType) String() string {
	switch o {
	case InstanceBlobTypeAll:
		return "InstanceBlobTypeAll"
	case InstanceBlobTypeError:
		return "InstanceBlobTypeError"
	case InstanceBlobTypeEnum:
		return "InstanceBlobTypeEnum"
	}
	return "Invalid"
}

// RefreshedObject structure represents WBEM_REFRESHED_OBJECT RPC structure.
//
// The WBEM_REFRESHED_OBJECT structure MUST be used to encode the results of the remote
// refreshing service that is returned by the IWbemRemoteRefresher::RemoteRefresh (section
// 3.1.4.13.1) interface method.
type RefreshedObject struct {
	// m_lRequestId:  MUST contain the request ID.
	RequestID int32 `idl:"name:m_lRequestId" json:"request_id"`
	// m_lBlobType:  MUST represent the type of the CIM object that is encoded in m_pbBlob
	// as specified in 2.2.17.
	BlobType InstanceBlobType `idl:"name:m_lBlobType" json:"blob_type"`
	// m_lBlobLength:  MUST represent the length of the m_pbBlob array.
	BlobLength int32  `idl:"name:m_lBlobLength" json:"blob_length"`
	Blob       []byte `idl:"name:m_pbBlob;size_is:(m_lBlobLength)" json:"blob"`
}

func (o *RefreshedObject) xxx_PreparePayload(ctx context.Context) error {
	if o.Blob != nil && o.BlobLength == 0 {
		o.BlobLength = int32(len(o.Blob))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *RefreshedObject) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.RequestID); err != nil {
		return err
	}
	if err := w.WriteEnum(uint32(o.BlobType)); err != nil {
		return err
	}
	if err := w.WriteData(o.BlobLength); err != nil {
		return err
	}
	if o.Blob != nil || o.BlobLength > 0 {
		_ptr_m_pbBlob := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.BlobLength)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.Blob {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if err := w.WriteData(o.Blob[i1]); err != nil {
					return err
				}
			}
			for i1 := len(o.Blob); uint64(i1) < sizeInfo[0]; i1++ {
				if err := w.WriteData(uint8(0)); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Blob, _ptr_m_pbBlob); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *RefreshedObject) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.RequestID); err != nil {
		return err
	}
	if err := w.ReadEnum((*uint32)(&o.BlobType)); err != nil {
		return err
	}
	if err := w.ReadData(&o.BlobLength); err != nil {
		return err
	}
	_ptr_m_pbBlob := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.BlobLength > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.BlobLength)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Blob", sizeInfo[0])
		}
		o.Blob = make([]byte, sizeInfo[0])
		for i1 := range o.Blob {
			i1 := i1
			if err := w.ReadData(&o.Blob[i1]); err != nil {
				return err
			}
		}
		return nil
	})
	_s_m_pbBlob := func(ptr interface{}) { o.Blob = *ptr.(*[]byte) }
	if err := w.ReadPointer(&o.Blob, _s_m_pbBlob, _ptr_m_pbBlob); err != nil {
		return err
	}
	return nil
}

// RefreshInfoRemote structure represents _WBEM_REFRESH_INFO_REMOTE RPC structure.
type RefreshInfoRemote struct {
	Refresher *RemoteRefresher `idl:"name:m_pRefresher" json:"refresher"`
	Template  *ClassObject     `idl:"name:m_pTemplate" json:"template"`
	GUID      *dtyp.GUID       `idl:"name:m_Guid" json:"guid"`
}

func (o *RefreshInfoRemote) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *RefreshInfoRemote) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if o.Refresher != nil {
		_ptr_m_pRefresher := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.Refresher != nil {
				if err := o.Refresher.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&RemoteRefresher{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Refresher, _ptr_m_pRefresher); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.Template != nil {
		_ptr_m_pTemplate := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.Template != nil {
				if err := o.Template.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&ClassObject{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Template, _ptr_m_pTemplate); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.GUID != nil {
		if err := o.GUID.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *RefreshInfoRemote) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	_ptr_m_pRefresher := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.Refresher == nil {
			o.Refresher = &RemoteRefresher{}
		}
		if err := o.Refresher.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_m_pRefresher := func(ptr interface{}) { o.Refresher = *ptr.(**RemoteRefresher) }
	if err := w.ReadPointer(&o.Refresher, _s_m_pRefresher, _ptr_m_pRefresher); err != nil {
		return err
	}
	_ptr_m_pTemplate := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.Template == nil {
			o.Template = &ClassObject{}
		}
		if err := o.Template.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_m_pTemplate := func(ptr interface{}) { o.Template = *ptr.(**ClassObject) }
	if err := w.ReadPointer(&o.Template, _s_m_pTemplate, _ptr_m_pTemplate); err != nil {
		return err
	}
	if o.GUID == nil {
		o.GUID = &dtyp.GUID{}
	}
	if err := o.GUID.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// RefreshInfoNonHiPerf structure represents _WBEM_REFRESH_INFO_NON_HIPERF RPC structure.
type RefreshInfoNonHiPerf struct {
	Namespace string       `idl:"name:m_wszNamespace;string" json:"namespace"`
	Template  *ClassObject `idl:"name:m_pTemplate" json:"template"`
}

func (o *RefreshInfoNonHiPerf) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *RefreshInfoNonHiPerf) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(6); err != nil {
		return err
	}
	if o.Namespace != "" {
		_ptr_m_wszNamespace := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.Namespace); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.Namespace, _ptr_m_wszNamespace); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.Template != nil {
		_ptr_m_pTemplate := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.Template != nil {
				if err := o.Template.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&ClassObject{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Template, _ptr_m_pTemplate); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *RefreshInfoNonHiPerf) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(6); err != nil {
		return err
	}
	_ptr_m_wszNamespace := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.Namespace); err != nil {
			return err
		}
		return nil
	})
	_s_m_wszNamespace := func(ptr interface{}) { o.Namespace = *ptr.(*string) }
	if err := w.ReadPointer(&o.Namespace, _s_m_wszNamespace, _ptr_m_wszNamespace); err != nil {
		return err
	}
	_ptr_m_pTemplate := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.Template == nil {
			o.Template = &ClassObject{}
		}
		if err := o.Template.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_m_pTemplate := func(ptr interface{}) { o.Template = *ptr.(**ClassObject) }
	if err := w.ReadPointer(&o.Template, _s_m_pTemplate, _ptr_m_pTemplate); err != nil {
		return err
	}
	return nil
}

// RefreshType type represents WBEM_REFRESH_TYPE RPC enumeration.
//
// The WBEM_REFRESH_TYPE enumeration defines refresh types for the _WBEM_REFRESH_INFO
// structure.
type RefreshType uint16

var (
	// WBEM_REFRESH_TYPE_INVALID:  The server uses this value internally. The server MUST
	// NOT return this value.
	RefreshTypeInvalid RefreshType = 0
	// WBEM_REFRESH_TYPE_REMOTE:  The m_Info member of the _WBEM_REFRESH_INFO structure
	// contains the _WBEM_REFRESH_INFO_REMOTE structure.
	RefreshTypeRemote RefreshType = 3
	// WBEM_REFRESH_TYPE_NON_HIPERF:  The m_Info member of the _WBEM_REFRESH_INFO structure
	// contains the _WBEM_REFRESH_INFO_NON_HIPERF structure.
	RefreshTypeNonHiPerf RefreshType = 6
)

func (o RefreshType) String() string {
	switch o {
	case RefreshTypeInvalid:
		return "RefreshTypeInvalid"
	case RefreshTypeRemote:
		return "RefreshTypeRemote"
	case RefreshTypeNonHiPerf:
		return "RefreshTypeNonHiPerf"
	}
	return "Invalid"
}

// RefreshInfoUnion structure represents WBEM_REFRESH_INFO_UNION RPC union.
//
// The _WBEM_REFRESH_INFO_UNION union defines a union of one of the following types:
// m_Remote, m_NonHiPerf, or m_hres.
type RefreshInfoUnion struct {
	// Types that are assignable to Value
	//
	// *RefreshInfoUnion_Remote
	// *RefreshInfoUnion_NonHiPerf
	// *RefreshInfoUnion_Hres
	Value is_RefreshInfoUnion `json:"value"`
}

func (o *RefreshInfoUnion) GetValue() any {
	if o == nil {
		return nil
	}
	switch value := (interface{})(o.Value).(type) {
	case *RefreshInfoUnion_Remote:
		if value != nil {
			return value.Remote
		}
	case *RefreshInfoUnion_NonHiPerf:
		if value != nil {
			return value.NonHiPerf
		}
	case *RefreshInfoUnion_Hres:
		if value != nil {
			return value.Hres
		}
	}
	return nil
}

type is_RefreshInfoUnion interface {
	ndr.Marshaler
	ndr.Unmarshaler
	is_RefreshInfoUnion()
}

func (o *RefreshInfoUnion) NDRSwitchValue(sw int32) int32 {
	if o == nil {
		return int32(0)
	}
	switch (interface{})(o.Value).(type) {
	case *RefreshInfoUnion_Remote:
		return int32(3)
	case *RefreshInfoUnion_NonHiPerf:
		return int32(6)
	case *RefreshInfoUnion_Hres:
		return int32(0)
	}
	return int32(0)
}

func (o *RefreshInfoUnion) MarshalUnionNDR(ctx context.Context, w ndr.Writer, sw int32) error {
	if err := w.WriteSwitch(int32(sw)); err != nil {
		return err
	}
	if err := w.WriteUnionAlign(9); err != nil {
		return err
	}
	switch sw {
	case int32(3):
		_o, _ := o.Value.(*RefreshInfoUnion_Remote)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&RefreshInfoUnion_Remote{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case int32(6):
		_o, _ := o.Value.(*RefreshInfoUnion_NonHiPerf)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&RefreshInfoUnion_NonHiPerf{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case int32(0):
		_o, _ := o.Value.(*RefreshInfoUnion_Hres)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&RefreshInfoUnion_Hres{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	default:
		return fmt.Errorf("unsupported switch case value %v", sw)
	}
	return nil
}

func (o *RefreshInfoUnion) UnmarshalUnionNDR(ctx context.Context, w ndr.Reader, sw int32) error {
	if err := w.ReadSwitch((*int32)(&sw)); err != nil {
		return err
	}
	if err := w.ReadUnionAlign(9); err != nil {
		return err
	}
	switch sw {
	case int32(3):
		o.Value = &RefreshInfoUnion_Remote{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case int32(6):
		o.Value = &RefreshInfoUnion_NonHiPerf{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case int32(0):
		o.Value = &RefreshInfoUnion_Hres{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	default:
		return fmt.Errorf("unsupported switch case value %v", sw)
	}
	return nil
}

// RefreshInfoUnion_Remote structure represents WBEM_REFRESH_INFO_UNION RPC union arm.
//
// It has following labels: 3
type RefreshInfoUnion_Remote struct {
	// m_Remote:  An m_Remote _WBEM_REFRESH_INFO_REMOTE type.
	Remote *RefreshInfoRemote `idl:"name:m_Remote" json:"remote"`
}

func (*RefreshInfoUnion_Remote) is_RefreshInfoUnion() {}

func (o *RefreshInfoUnion_Remote) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.Remote != nil {
		if err := o.Remote.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&RefreshInfoRemote{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *RefreshInfoUnion_Remote) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.Remote == nil {
		o.Remote = &RefreshInfoRemote{}
	}
	if err := o.Remote.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// RefreshInfoUnion_NonHiPerf structure represents WBEM_REFRESH_INFO_UNION RPC union arm.
//
// It has following labels: 6
type RefreshInfoUnion_NonHiPerf struct {
	// m_NonHiPerf:  An m_NonHiPerf _WBEM_REFRESH_INFO_NON_HIPERF type.
	NonHiPerf *RefreshInfoNonHiPerf `idl:"name:m_NonHiPerf" json:"non_hi_perf"`
}

func (*RefreshInfoUnion_NonHiPerf) is_RefreshInfoUnion() {}

func (o *RefreshInfoUnion_NonHiPerf) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.NonHiPerf != nil {
		if err := o.NonHiPerf.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&RefreshInfoNonHiPerf{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *RefreshInfoUnion_NonHiPerf) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.NonHiPerf == nil {
		o.NonHiPerf = &RefreshInfoNonHiPerf{}
	}
	if err := o.NonHiPerf.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// RefreshInfoUnion_Hres structure represents WBEM_REFRESH_INFO_UNION RPC union arm.
//
// It has following labels: 0
type RefreshInfoUnion_Hres struct {
	// m_hres:  An m_hres HRESULT type.
	Hres int32 `idl:"name:m_hres" json:"hres"`
}

func (*RefreshInfoUnion_Hres) is_RefreshInfoUnion() {}

func (o *RefreshInfoUnion_Hres) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := w.WriteData(o.Hres); err != nil {
		return err
	}
	return nil
}
func (o *RefreshInfoUnion_Hres) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadData(&o.Hres); err != nil {
		return err
	}
	return nil
}

// RefreshInfo structure represents _WBEM_REFRESH_INFO RPC structure.
type RefreshInfo struct {
	Type     int32             `idl:"name:m_lType" json:"type"`
	Info     *RefreshInfoUnion `idl:"name:m_Info;switch_is:m_lType" json:"info"`
	CancelID int32             `idl:"name:m_lCancelId" json:"cancel_id"`
}

func (o *RefreshInfo) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *RefreshInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.Type); err != nil {
		return err
	}
	_swInfo := int32(o.Type)
	if o.Info != nil {
		if err := o.Info.MarshalUnionNDR(ctx, w, _swInfo); err != nil {
			return err
		}
	} else {
		if err := (&RefreshInfoUnion{}).MarshalUnionNDR(ctx, w, _swInfo); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.CancelID); err != nil {
		return err
	}
	return nil
}
func (o *RefreshInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.Type); err != nil {
		return err
	}
	if o.Info == nil {
		o.Info = &RefreshInfoUnion{}
	}
	_swInfo := int32(o.Type)
	if err := o.Info.UnmarshalUnionNDR(ctx, w, _swInfo); err != nil {
		return err
	}
	if err := w.ReadData(&o.CancelID); err != nil {
		return err
	}
	return nil
}

// RefresherID structure represents _WBEM_REFRESHER_ID RPC structure.
type RefresherID struct {
	MachineName string     `idl:"name:m_szMachineName;string" json:"machine_name"`
	ProcessID   uint32     `idl:"name:m_dwProcessId" json:"process_id"`
	RefresherID *dtyp.GUID `idl:"name:m_guidRefresherId" json:"refresher_id"`
}

func (o *RefresherID) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *RefresherID) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if o.MachineName != "" {
		_ptr_m_szMachineName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteCharNString(ctx, w, o.MachineName); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.MachineName, _ptr_m_szMachineName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.ProcessID); err != nil {
		return err
	}
	if o.RefresherID != nil {
		if err := o.RefresherID.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *RefresherID) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	_ptr_m_szMachineName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadCharNString(ctx, w, &o.MachineName); err != nil {
			return err
		}
		return nil
	})
	_s_m_szMachineName := func(ptr interface{}) { o.MachineName = *ptr.(*string) }
	if err := w.ReadPointer(&o.MachineName, _s_m_szMachineName, _ptr_m_szMachineName); err != nil {
		return err
	}
	if err := w.ReadData(&o.ProcessID); err != nil {
		return err
	}
	if o.RefresherID == nil {
		o.RefresherID = &dtyp.GUID{}
	}
	if err := o.RefresherID.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// ReconnectType type represents WBEM_RECONNECT_TYPE RPC enumeration.
//
// The _WBEM_RECONNECT_TYPE enumeration defines possible types of remote CIM instances.
// The structure MUST be used to return to information from IWbemRefreshingServices
// (section 3.1.4.12) interface methods.
type ReconnectType uint16

var (
	// WBEM_RECONNECT_TYPE_OBJECT:  The refresher MUST connect to refresh an object.
	ReconnectTypeObject ReconnectType = 0
	// WBEM_RECONNECT_TYPE_ENUM:  The refresher MUST connect to refresh an enumeration.
	ReconnectTypeEnum ReconnectType = 1
	// WBEM_RECONNECT_TYPE_LAST:  This member is used only by the server to track the range
	// of values for this enumeration. It MUST NOT be used by the client.
	ReconnectTypeLast ReconnectType = 2
)

func (o ReconnectType) String() string {
	switch o {
	case ReconnectTypeObject:
		return "ReconnectTypeObject"
	case ReconnectTypeEnum:
		return "ReconnectTypeEnum"
	case ReconnectTypeLast:
		return "ReconnectTypeLast"
	}
	return "Invalid"
}

// ReconnectInfo structure represents _WBEM_RECONNECT_INFO RPC structure.
type ReconnectInfo struct {
	Type int32  `idl:"name:m_lType" json:"type"`
	Path string `idl:"name:m_pwcsPath;string" json:"path"`
}

func (o *ReconnectInfo) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *ReconnectInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.Type); err != nil {
		return err
	}
	if o.Path != "" {
		_ptr_m_pwcsPath := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.Path); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.Path, _ptr_m_pwcsPath); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *ReconnectInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.Type); err != nil {
		return err
	}
	_ptr_m_pwcsPath := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.Path); err != nil {
			return err
		}
		return nil
	})
	_s_m_pwcsPath := func(ptr interface{}) { o.Path = *ptr.(*string) }
	if err := w.ReadPointer(&o.Path, _s_m_pwcsPath, _ptr_m_pwcsPath); err != nil {
		return err
	}
	return nil
}

// ReconnectResults structure represents _WBEM_RECONNECT_RESULTS RPC structure.
type ReconnectResults struct {
	ID      int32 `idl:"name:m_lId" json:"id"`
	HResult int32 `idl:"name:m_hr" json:"hresult"`
}

func (o *ReconnectResults) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *ReconnectResults) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.ID); err != nil {
		return err
	}
	if err := w.WriteData(o.HResult); err != nil {
		return err
	}
	return nil
}
func (o *ReconnectResults) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.ID); err != nil {
		return err
	}
	if err := w.ReadData(&o.HResult); err != nil {
		return err
	}
	return nil
}

// LoginClientID structure represents IWbemLoginClientID RPC structure.
type LoginClientID dcom.InterfacePointer

func (o *LoginClientID) InterfacePointer() *dcom.InterfacePointer { return (*dcom.InterfacePointer)(o) }

func (o *LoginClientID) xxx_PreparePayload(ctx context.Context) error {
	if o.Data != nil && o.DataCount == 0 {
		o.DataCount = uint32(len(o.Data))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *LoginClientID) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *LoginClientID) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.DataCount); err != nil {
		return err
	}
	for i1 := range o.Data {
		i1 := i1
		if uint64(i1) >= sizeInfo[0] {
			break
		}
		if err := w.WriteData(o.Data[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.Data); uint64(i1) < sizeInfo[0]; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	return nil
}
func (o *LoginClientID) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.DataCount); err != nil {
		return err
	}
	// XXX: for opaque unmarshaling
	if o.DataCount > 0 && sizeInfo[0] == 0 {
		sizeInfo[0] = uint64(o.DataCount)
	}
	if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
		return fmt.Errorf("buffer overflow for size %d of array o.Data", sizeInfo[0])
	}
	o.Data = make([]byte, sizeInfo[0])
	for i1 := range o.Data {
		i1 := i1
		if err := w.ReadData(&o.Data[i1]); err != nil {
			return err
		}
	}
	return nil
}

// Level1Login structure represents IWbemLevel1Login RPC structure.
type Level1Login dcom.InterfacePointer

func (o *Level1Login) InterfacePointer() *dcom.InterfacePointer { return (*dcom.InterfacePointer)(o) }

func (o *Level1Login) xxx_PreparePayload(ctx context.Context) error {
	if o.Data != nil && o.DataCount == 0 {
		o.DataCount = uint32(len(o.Data))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *Level1Login) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *Level1Login) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.DataCount); err != nil {
		return err
	}
	for i1 := range o.Data {
		i1 := i1
		if uint64(i1) >= sizeInfo[0] {
			break
		}
		if err := w.WriteData(o.Data[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.Data); uint64(i1) < sizeInfo[0]; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	return nil
}
func (o *Level1Login) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.DataCount); err != nil {
		return err
	}
	// XXX: for opaque unmarshaling
	if o.DataCount > 0 && sizeInfo[0] == 0 {
		sizeInfo[0] = uint64(o.DataCount)
	}
	if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
		return fmt.Errorf("buffer overflow for size %d of array o.Data", sizeInfo[0])
	}
	o.Data = make([]byte, sizeInfo[0])
	for i1 := range o.Data {
		i1 := i1
		if err := w.ReadData(&o.Data[i1]); err != nil {
			return err
		}
	}
	return nil
}

// Services structure represents IWbemServices RPC structure.
type Services dcom.InterfacePointer

func (o *Services) InterfacePointer() *dcom.InterfacePointer { return (*dcom.InterfacePointer)(o) }

func (o *Services) xxx_PreparePayload(ctx context.Context) error {
	if o.Data != nil && o.DataCount == 0 {
		o.DataCount = uint32(len(o.Data))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *Services) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *Services) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.DataCount); err != nil {
		return err
	}
	for i1 := range o.Data {
		i1 := i1
		if uint64(i1) >= sizeInfo[0] {
			break
		}
		if err := w.WriteData(o.Data[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.Data); uint64(i1) < sizeInfo[0]; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	return nil
}
func (o *Services) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.DataCount); err != nil {
		return err
	}
	// XXX: for opaque unmarshaling
	if o.DataCount > 0 && sizeInfo[0] == 0 {
		sizeInfo[0] = uint64(o.DataCount)
	}
	if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
		return fmt.Errorf("buffer overflow for size %d of array o.Data", sizeInfo[0])
	}
	o.Data = make([]byte, sizeInfo[0])
	for i1 := range o.Data {
		i1 := i1
		if err := w.ReadData(&o.Data[i1]); err != nil {
			return err
		}
	}
	return nil
}

// EnumClassObject structure represents IEnumWbemClassObject RPC structure.
type EnumClassObject dcom.InterfacePointer

func (o *EnumClassObject) InterfacePointer() *dcom.InterfacePointer {
	return (*dcom.InterfacePointer)(o)
}

func (o *EnumClassObject) xxx_PreparePayload(ctx context.Context) error {
	if o.Data != nil && o.DataCount == 0 {
		o.DataCount = uint32(len(o.Data))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *EnumClassObject) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *EnumClassObject) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.DataCount); err != nil {
		return err
	}
	for i1 := range o.Data {
		i1 := i1
		if uint64(i1) >= sizeInfo[0] {
			break
		}
		if err := w.WriteData(o.Data[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.Data); uint64(i1) < sizeInfo[0]; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	return nil
}
func (o *EnumClassObject) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.DataCount); err != nil {
		return err
	}
	// XXX: for opaque unmarshaling
	if o.DataCount > 0 && sizeInfo[0] == 0 {
		sizeInfo[0] = uint64(o.DataCount)
	}
	if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
		return fmt.Errorf("buffer overflow for size %d of array o.Data", sizeInfo[0])
	}
	o.Data = make([]byte, sizeInfo[0])
	for i1 := range o.Data {
		i1 := i1
		if err := w.ReadData(&o.Data[i1]); err != nil {
			return err
		}
	}
	return nil
}

// FetchSmartEnum structure represents IWbemFetchSmartEnum RPC structure.
type FetchSmartEnum dcom.InterfacePointer

func (o *FetchSmartEnum) InterfacePointer() *dcom.InterfacePointer {
	return (*dcom.InterfacePointer)(o)
}

func (o *FetchSmartEnum) xxx_PreparePayload(ctx context.Context) error {
	if o.Data != nil && o.DataCount == 0 {
		o.DataCount = uint32(len(o.Data))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *FetchSmartEnum) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *FetchSmartEnum) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.DataCount); err != nil {
		return err
	}
	for i1 := range o.Data {
		i1 := i1
		if uint64(i1) >= sizeInfo[0] {
			break
		}
		if err := w.WriteData(o.Data[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.Data); uint64(i1) < sizeInfo[0]; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	return nil
}
func (o *FetchSmartEnum) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.DataCount); err != nil {
		return err
	}
	// XXX: for opaque unmarshaling
	if o.DataCount > 0 && sizeInfo[0] == 0 {
		sizeInfo[0] = uint64(o.DataCount)
	}
	if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
		return fmt.Errorf("buffer overflow for size %d of array o.Data", sizeInfo[0])
	}
	o.Data = make([]byte, sizeInfo[0])
	for i1 := range o.Data {
		i1 := i1
		if err := w.ReadData(&o.Data[i1]); err != nil {
			return err
		}
	}
	return nil
}

// ClassObject structure represents IWbemClassObject RPC structure.
type ClassObject dcom.InterfacePointer

func (o *ClassObject) InterfacePointer() *dcom.InterfacePointer { return (*dcom.InterfacePointer)(o) }

func (o *ClassObject) xxx_PreparePayload(ctx context.Context) error {
	if o.Data != nil && o.DataCount == 0 {
		o.DataCount = uint32(len(o.Data))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *ClassObject) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *ClassObject) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.DataCount); err != nil {
		return err
	}
	for i1 := range o.Data {
		i1 := i1
		if uint64(i1) >= sizeInfo[0] {
			break
		}
		if err := w.WriteData(o.Data[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.Data); uint64(i1) < sizeInfo[0]; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	return nil
}
func (o *ClassObject) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.DataCount); err != nil {
		return err
	}
	// XXX: for opaque unmarshaling
	if o.DataCount > 0 && sizeInfo[0] == 0 {
		sizeInfo[0] = uint64(o.DataCount)
	}
	if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
		return fmt.Errorf("buffer overflow for size %d of array o.Data", sizeInfo[0])
	}
	o.Data = make([]byte, sizeInfo[0])
	for i1 := range o.Data {
		i1 := i1
		if err := w.ReadData(&o.Data[i1]); err != nil {
			return err
		}
	}
	return nil
}

// RemoteRefresher structure represents IWbemRemoteRefresher RPC structure.
type RemoteRefresher dcom.InterfacePointer

func (o *RemoteRefresher) InterfacePointer() *dcom.InterfacePointer {
	return (*dcom.InterfacePointer)(o)
}

func (o *RemoteRefresher) xxx_PreparePayload(ctx context.Context) error {
	if o.Data != nil && o.DataCount == 0 {
		o.DataCount = uint32(len(o.Data))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *RemoteRefresher) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *RemoteRefresher) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.DataCount); err != nil {
		return err
	}
	for i1 := range o.Data {
		i1 := i1
		if uint64(i1) >= sizeInfo[0] {
			break
		}
		if err := w.WriteData(o.Data[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.Data); uint64(i1) < sizeInfo[0]; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	return nil
}
func (o *RemoteRefresher) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.DataCount); err != nil {
		return err
	}
	// XXX: for opaque unmarshaling
	if o.DataCount > 0 && sizeInfo[0] == 0 {
		sizeInfo[0] = uint64(o.DataCount)
	}
	if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
		return fmt.Errorf("buffer overflow for size %d of array o.Data", sizeInfo[0])
	}
	o.Data = make([]byte, sizeInfo[0])
	for i1 := range o.Data {
		i1 := i1
		if err := w.ReadData(&o.Data[i1]); err != nil {
			return err
		}
	}
	return nil
}

// RefreshingServices structure represents IWbemRefreshingServices RPC structure.
type RefreshingServices dcom.InterfacePointer

func (o *RefreshingServices) InterfacePointer() *dcom.InterfacePointer {
	return (*dcom.InterfacePointer)(o)
}

func (o *RefreshingServices) xxx_PreparePayload(ctx context.Context) error {
	if o.Data != nil && o.DataCount == 0 {
		o.DataCount = uint32(len(o.Data))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *RefreshingServices) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *RefreshingServices) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.DataCount); err != nil {
		return err
	}
	for i1 := range o.Data {
		i1 := i1
		if uint64(i1) >= sizeInfo[0] {
			break
		}
		if err := w.WriteData(o.Data[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.Data); uint64(i1) < sizeInfo[0]; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	return nil
}
func (o *RefreshingServices) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.DataCount); err != nil {
		return err
	}
	// XXX: for opaque unmarshaling
	if o.DataCount > 0 && sizeInfo[0] == 0 {
		sizeInfo[0] = uint64(o.DataCount)
	}
	if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
		return fmt.Errorf("buffer overflow for size %d of array o.Data", sizeInfo[0])
	}
	o.Data = make([]byte, sizeInfo[0])
	for i1 := range o.Data {
		i1 := i1
		if err := w.ReadData(&o.Data[i1]); err != nil {
			return err
		}
	}
	return nil
}

// CallResult structure represents IWbemCallResult RPC structure.
type CallResult dcom.InterfacePointer

func (o *CallResult) InterfacePointer() *dcom.InterfacePointer { return (*dcom.InterfacePointer)(o) }

func (o *CallResult) xxx_PreparePayload(ctx context.Context) error {
	if o.Data != nil && o.DataCount == 0 {
		o.DataCount = uint32(len(o.Data))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *CallResult) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *CallResult) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.DataCount); err != nil {
		return err
	}
	for i1 := range o.Data {
		i1 := i1
		if uint64(i1) >= sizeInfo[0] {
			break
		}
		if err := w.WriteData(o.Data[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.Data); uint64(i1) < sizeInfo[0]; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	return nil
}
func (o *CallResult) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.DataCount); err != nil {
		return err
	}
	// XXX: for opaque unmarshaling
	if o.DataCount > 0 && sizeInfo[0] == 0 {
		sizeInfo[0] = uint64(o.DataCount)
	}
	if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
		return fmt.Errorf("buffer overflow for size %d of array o.Data", sizeInfo[0])
	}
	o.Data = make([]byte, sizeInfo[0])
	for i1 := range o.Data {
		i1 := i1
		if err := w.ReadData(&o.Data[i1]); err != nil {
			return err
		}
	}
	return nil
}

// Context structure represents IWbemContext RPC structure.
type Context dcom.InterfacePointer

func (o *Context) InterfacePointer() *dcom.InterfacePointer { return (*dcom.InterfacePointer)(o) }

func (o *Context) xxx_PreparePayload(ctx context.Context) error {
	if o.Data != nil && o.DataCount == 0 {
		o.DataCount = uint32(len(o.Data))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *Context) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *Context) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.DataCount); err != nil {
		return err
	}
	for i1 := range o.Data {
		i1 := i1
		if uint64(i1) >= sizeInfo[0] {
			break
		}
		if err := w.WriteData(o.Data[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.Data); uint64(i1) < sizeInfo[0]; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	return nil
}
func (o *Context) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.DataCount); err != nil {
		return err
	}
	// XXX: for opaque unmarshaling
	if o.DataCount > 0 && sizeInfo[0] == 0 {
		sizeInfo[0] = uint64(o.DataCount)
	}
	if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
		return fmt.Errorf("buffer overflow for size %d of array o.Data", sizeInfo[0])
	}
	o.Data = make([]byte, sizeInfo[0])
	for i1 := range o.Data {
		i1 := i1
		if err := w.ReadData(&o.Data[i1]); err != nil {
			return err
		}
	}
	return nil
}

// WCOSmartEnum structure represents IWbemWCOSmartEnum RPC structure.
type WCOSmartEnum dcom.InterfacePointer

func (o *WCOSmartEnum) InterfacePointer() *dcom.InterfacePointer { return (*dcom.InterfacePointer)(o) }

func (o *WCOSmartEnum) xxx_PreparePayload(ctx context.Context) error {
	if o.Data != nil && o.DataCount == 0 {
		o.DataCount = uint32(len(o.Data))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *WCOSmartEnum) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *WCOSmartEnum) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.DataCount); err != nil {
		return err
	}
	for i1 := range o.Data {
		i1 := i1
		if uint64(i1) >= sizeInfo[0] {
			break
		}
		if err := w.WriteData(o.Data[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.Data); uint64(i1) < sizeInfo[0]; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	return nil
}
func (o *WCOSmartEnum) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.DataCount); err != nil {
		return err
	}
	// XXX: for opaque unmarshaling
	if o.DataCount > 0 && sizeInfo[0] == 0 {
		sizeInfo[0] = uint64(o.DataCount)
	}
	if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
		return fmt.Errorf("buffer overflow for size %d of array o.Data", sizeInfo[0])
	}
	o.Data = make([]byte, sizeInfo[0])
	for i1 := range o.Data {
		i1 := i1
		if err := w.ReadData(&o.Data[i1]); err != nil {
			return err
		}
	}
	return nil
}

// BackupRestoreEx structure represents IWbemBackupRestoreEx RPC structure.
type BackupRestoreEx dcom.InterfacePointer

func (o *BackupRestoreEx) InterfacePointer() *dcom.InterfacePointer {
	return (*dcom.InterfacePointer)(o)
}

func (o *BackupRestoreEx) xxx_PreparePayload(ctx context.Context) error {
	if o.Data != nil && o.DataCount == 0 {
		o.DataCount = uint32(len(o.Data))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *BackupRestoreEx) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *BackupRestoreEx) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.DataCount); err != nil {
		return err
	}
	for i1 := range o.Data {
		i1 := i1
		if uint64(i1) >= sizeInfo[0] {
			break
		}
		if err := w.WriteData(o.Data[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.Data); uint64(i1) < sizeInfo[0]; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	return nil
}
func (o *BackupRestoreEx) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.DataCount); err != nil {
		return err
	}
	// XXX: for opaque unmarshaling
	if o.DataCount > 0 && sizeInfo[0] == 0 {
		sizeInfo[0] = uint64(o.DataCount)
	}
	if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
		return fmt.Errorf("buffer overflow for size %d of array o.Data", sizeInfo[0])
	}
	o.Data = make([]byte, sizeInfo[0])
	for i1 := range o.Data {
		i1 := i1
		if err := w.ReadData(&o.Data[i1]); err != nil {
			return err
		}
	}
	return nil
}

// LoginHelper structure represents IWbemLoginHelper RPC structure.
type LoginHelper dcom.InterfacePointer

func (o *LoginHelper) InterfacePointer() *dcom.InterfacePointer { return (*dcom.InterfacePointer)(o) }

func (o *LoginHelper) xxx_PreparePayload(ctx context.Context) error {
	if o.Data != nil && o.DataCount == 0 {
		o.DataCount = uint32(len(o.Data))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *LoginHelper) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *LoginHelper) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.DataCount); err != nil {
		return err
	}
	for i1 := range o.Data {
		i1 := i1
		if uint64(i1) >= sizeInfo[0] {
			break
		}
		if err := w.WriteData(o.Data[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.Data); uint64(i1) < sizeInfo[0]; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	return nil
}
func (o *LoginHelper) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.DataCount); err != nil {
		return err
	}
	// XXX: for opaque unmarshaling
	if o.DataCount > 0 && sizeInfo[0] == 0 {
		sizeInfo[0] = uint64(o.DataCount)
	}
	if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
		return fmt.Errorf("buffer overflow for size %d of array o.Data", sizeInfo[0])
	}
	o.Data = make([]byte, sizeInfo[0])
	for i1 := range o.Data {
		i1 := i1
		if err := w.ReadData(&o.Data[i1]); err != nil {
			return err
		}
	}
	return nil
}

// ObjectSink structure represents IWbemObjectSink RPC structure.
type ObjectSink dcom.InterfacePointer

func (o *ObjectSink) InterfacePointer() *dcom.InterfacePointer { return (*dcom.InterfacePointer)(o) }

func (o *ObjectSink) xxx_PreparePayload(ctx context.Context) error {
	if o.Data != nil && o.DataCount == 0 {
		o.DataCount = uint32(len(o.Data))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *ObjectSink) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *ObjectSink) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.DataCount); err != nil {
		return err
	}
	for i1 := range o.Data {
		i1 := i1
		if uint64(i1) >= sizeInfo[0] {
			break
		}
		if err := w.WriteData(o.Data[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.Data); uint64(i1) < sizeInfo[0]; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	return nil
}
func (o *ObjectSink) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.DataCount); err != nil {
		return err
	}
	// XXX: for opaque unmarshaling
	if o.DataCount > 0 && sizeInfo[0] == 0 {
		sizeInfo[0] = uint64(o.DataCount)
	}
	if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
		return fmt.Errorf("buffer overflow for size %d of array o.Data", sizeInfo[0])
	}
	o.Data = make([]byte, sizeInfo[0])
	for i1 := range o.Data {
		i1 := i1
		if err := w.ReadData(&o.Data[i1]); err != nil {
			return err
		}
	}
	return nil
}

// BackupRestore structure represents IWbemBackupRestore RPC structure.
type BackupRestore dcom.InterfacePointer

func (o *BackupRestore) InterfacePointer() *dcom.InterfacePointer { return (*dcom.InterfacePointer)(o) }

func (o *BackupRestore) xxx_PreparePayload(ctx context.Context) error {
	if o.Data != nil && o.DataCount == 0 {
		o.DataCount = uint32(len(o.Data))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *BackupRestore) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *BackupRestore) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.DataCount); err != nil {
		return err
	}
	for i1 := range o.Data {
		i1 := i1
		if uint64(i1) >= sizeInfo[0] {
			break
		}
		if err := w.WriteData(o.Data[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.Data); uint64(i1) < sizeInfo[0]; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	return nil
}
func (o *BackupRestore) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.DataCount); err != nil {
		return err
	}
	// XXX: for opaque unmarshaling
	if o.DataCount > 0 && sizeInfo[0] == 0 {
		sizeInfo[0] = uint64(o.DataCount)
	}
	if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
		return fmt.Errorf("buffer overflow for size %d of array o.Data", sizeInfo[0])
	}
	o.Data = make([]byte, sizeInfo[0])
	for i1 := range o.Data {
		i1 := i1
		if err := w.ReadData(&o.Data[i1]); err != nil {
			return err
		}
	}
	return nil
}

// WbemLevel1Login class identifier 8bc3f05e-d86b-11d0-a075-00c04fb68820
var Level1LoginClassID = &dcom.ClassID{Data1: 0x8bc3f05e, Data2: 0xd86b, Data3: 0x11d0, Data4: []byte{0xa0, 0x75, 0x00, 0xc0, 0x4f, 0xb6, 0x88, 0x20}}

// WbemContext class identifier 674b6698-ee92-11d0-ad71-00c04fd8fdff
var ContextClassID = &dcom.ClassID{Data1: 0x674b6698, Data2: 0xee92, Data3: 0x11d0, Data4: []byte{0xad, 0x71, 0x00, 0xc0, 0x4f, 0xd8, 0xfd, 0xff}}

// WbemClassObject class identifier 9a653086-174f-11d2-b5f9-00104b703efd
var ClassObjectClassID = &dcom.ClassID{Data1: 0x9a653086, Data2: 0x174f, Data3: 0x11d2, Data4: []byte{0xb5, 0xf9, 0x00, 0x10, 0x4b, 0x70, 0x3e, 0xfd}}

// WbemBackupRestore class identifier c49e32c6-bc8b-11d2-85d4-00105a1f8304
var BackupRestoreClassID = &dcom.ClassID{Data1: 0xc49e32c6, Data2: 0xbc8b, Data3: 0x11d2, Data4: []byte{0x85, 0xd4, 0x00, 0x10, 0x5a, 0x1f, 0x83, 0x04}}
