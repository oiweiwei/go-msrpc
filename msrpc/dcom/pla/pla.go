// The pla package implements the PLA client protocol.
//
// # Introduction
//
// The Performance Logs and Alerts (PLA) protocol is a set of Distributed Component
// Object Model (DCOM) interfaces (as specified in [MS-DCOM]) for logging diagnosis
// data on a remote computer.
//
// The PLA Protocol allows users to:
//
// * Specify the diagnostic data to be collected and logged.
//
// * Specify the data retention and reporting policies for the logged data.
//
// * Specify alert conditions.
//
// * Start, stop, or schedule the collection.
//
// # Overview
//
// Software components can be designed to assist in serviceability, manageability, supportability,
// and diagnostic ability. For instance, performance counters are a simple way of exposing
// state information that can be sampled or polled. Event-based instrumentation typically
// generates a state change notification. Alerts are a simple way of turning a sampled
// counter into an event notification, based on a threshold value.
//
// System administrators often want to collect diagnosis data on a remote system in
// a periodic or ongoing basis to better support and diagnose problems on the systems.
// Furthermore, the collected data can be processed by tools for in-depth problem analysis.
//
// The Performance Logs and Alerts Protocol provides a set of DCOM interfaces to control
// data collection on a remote system. The control includes creating, starting, stopping,
// scheduling, and configuring data collector objects and the creation of alerts.
//
// The capabilities of the Performance Logs and Alerts Protocol are summarized as follows:
//
// * Performance Counter Logging (section 3.2.4.6 ( 489d285e-94eb-46b2-8a53-b3501a61b572
// ) ): The Performance Logs and Alerts Protocol allows users to log performance counters'
// data of resources on a remote system. A resource can be hardware (for example, CPU,
// memory) or software (for example, application, process). The logged performance counter
// data is often useful for the analysis of performance trends and bottlenecks. The
// PLA Protocol also supports logging performance counter data in a SQL database format
// (section 3.2.4.6). This option defines the name of an existing SQL database and log
// set within the database where the performance counter data will be read or written.
// This file format is useful when collecting and analyzing performance counter data
// at an enterprise level rather than on a per-computer basis.
//
// * Event Trace Logging (section 3.2.4.9 ( 94e049b9-532a-4c3a-ada2-43c75802e825 ) ):
// The Performance Logs and Alerts Protocol allows users to log event tracing data of
// resources on a remote system. The event provider is software that can create event
// notifications and generate events when certain activities, such as a disk ( f43f48aa-80a5-4b39-971e-7b3ac0bd9d0d#gt_c4133b2a-a990-4042-ba44-7fda3090f118
// ) I/O operation or a page fault, occur. The application that uses the Performance
// Logs and Alerts Protocol can enable or disable event providers and selectively log
// the events of interest into a file.
//
// * API Trace Logging (section 3.2.4.10 ( fed87908-2a85-4cf6-b4ac-723b81b00e1b ) ):
// The Performance Logs and Alerts Protocol allows users to log the API call activity
// of an executable on a remote system. Observing API call activity is useful for the
// diagnosis of various executable issues (for example, detecting unnecessary API calls.)
//
// * Configuration Data Logging (section 3.2.4.7 ( 0f979443-9db2-4805-b2ea-e1540d8b0533
// ) ): The Performance Logs and Alerts Protocol allows users to log the computer configuration
// information on a remote system. Readjustment of an incorrect setting is one of the
// common diagnosis root causes.
//
// * Alerts (section 3.2.4.8 ( 07d907e9-10ed-42f6-9135-9131657a240f ) ): The Performance
// Logs and Alerts Protocol allows users to create alerts based on performance counter
// values on a remote system. An alert can trigger running a program, logging the alert
// as an event, or starting another data collection.
//
// * Data Collector Set (section 3.2.4.1 ( 1809d280-56e0-4c78-9546-ad1869c3a16a ) ):
// The Performance Logs and Alerts Protocol allows users to group multiple logging entities'
// data collectors and apply operations to them at once. The operations include start
// (section 3.2.4.1.56 ( d2ebf91b-9e67-440d-90e9-3134ee1613a0 ) ), stop (section 3.2.4.1.57
// ( 2617595b-63fc-4c23-b674-4cb3e062eb6f ) ), schedule (section 3.2.4.1.20 ( 790ac479-7c9a-4ce2-ac82-108ac3e5121d
// ) ), and configure (section 3.2.4.1).
//
// * Data Management (section 3.2.4.2 ( 16c70f7d-0f0e-4ae6-9785-be0032013c9f ) ): The
// Performance Logs and Alerts Protocol allows users to set a data retention policy
// against logged data and define post-actions of the collection. The post-actions,
// such as delete largest log file and compress log file, can be defined with the Performance
// Logs and Alerts Protocol interfaces.
package pla

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	dcom "github.com/oiweiwei/go-msrpc/msrpc/dcom"
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
	_ = dcom.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/pla"
)

// DataCollectorType type represents DataCollectorType RPC enumeration.
//
// The DataCollectorType enumeration defines the data collector types.
//
// The DataCollectorType (Get) method retrieves the DataCollectorType property.
type DataCollectorType uint16

var (
	// plaPerformanceCounter:  Collects performance counter data. The IPerformanceCounterDataCollector
	// interface represents this data collector.
	DataCollectorTypePerformanceCounter DataCollectorType = 0
	// plaTrace:  Collects events from an event trace session. The ITraceDataCollector
	// interface represents this data collector.
	DataCollectorTypeTrace DataCollectorType = 1
	// plaConfiguration:  Collects computer configuration information. The IConfigurationDataCollector
	// interface represents this data collector.
	DataCollectorTypeConfiguration DataCollectorType = 2
	// plaAlert:  Monitors performance counters and performs actions if the counter value
	// crosses the given threshold. The IAlertDataCollector interface represents this data
	// collector.
	DataCollectorTypeAlert DataCollectorType = 3
	// plaApiTrace:  Logs API calls made by the process. The IApiTracingDataCollector interface
	// represents this data collector.
	DataCollectorTypeAPITrace DataCollectorType = 4
)

func (o DataCollectorType) String() string {
	switch o {
	case DataCollectorTypePerformanceCounter:
		return "DataCollectorTypePerformanceCounter"
	case DataCollectorTypeTrace:
		return "DataCollectorTypeTrace"
	case DataCollectorTypeConfiguration:
		return "DataCollectorTypeConfiguration"
	case DataCollectorTypeAlert:
		return "DataCollectorTypeAlert"
	case DataCollectorTypeAPITrace:
		return "DataCollectorTypeAPITrace"
	}
	return "Invalid"
}

// FileFormat type represents FileFormat RPC enumeration.
//
// The FileFormat enumeration defines the format of the data in the log file.
type FileFormat uint16

var (
	// plaCommaSeparated:  Comma-separated log file. The first line in the text file contains
	// column headers followed by comma-separated data in the remaining lines of the log
	// file.
	FileFormatCommaSeparated FileFormat = 0
	// plaTabSeparated:  Tab-separated log file. The first line in the text file contains
	// column headers followed by tab-separated data in the remaining lines of the log file.
	FileFormatTabSeparated FileFormat = 1
	// plaSql:  The data is saved into a SQL database, instead of to a file. The SQL database
	// contains three tables: CounterData, CounterDetails, and DisplayToId. All three tables
	// are specified below.
	//
	// The CounterData table contains a row for each counter that is collected at a particular
	// time. There will be a large number of these rows. The GUID, CounterID, and RecordIndex
	// fields make up the primary key for this table.
	//
	// * GUID(uniqueidentifier, NOT NULL): GUID, as specified in [MS-DTYP] ( ../ms-dtyp/cca27429-5689-4a16-b2b4-9325d93e4ba2
	// ) section 2.3.4 ( ../ms-dtyp/4926e530-816e-41c2-b251-ec5c7aca018a ) , for this data
	// set. Use this key to join with the DisplayToID table.
	//
	// * CounterID(int, NOT NULL): Identifies the counter. Use this key to join with the
	// CounterDetails table.
	//
	// * RecordIndex(int, NOT NULL): The sample index for a specific counter identifier
	// and collection PLA-UID ( f43f48aa-80a5-4b39-971e-7b3ac0bd9d0d#gt_1a61d947-4d31-4365-b2b4-5249de682b56
	// ). The value increases for each successive sample in this log file.
	//
	// * CounterDateTime(char(24), NOT NULL): The time the collection was started, in UTC
	// time.
	//
	// * CounterValue(float, NOT NULL): The formatted value of the counter. This value can
	// be zero for the first record if the counter requires two samples to compute a displayable
	// value.
	//
	// * FirstValueA(int): Combine this 32-bit value with the value of FirstValueB to create
	// the FirstValue member of PDH_RAW_COUNTER. FirstValueA contains the low-order bits.
	//
	// * FirstValueB(int): Combine this 32-bit value with the value of FirstValueA to create
	// the FirstValue member of PDH_RAW_COUNTER. FirstValueB contains the high-order bits.
	//
	// * SecondValueA(int): Combine this 32-bit value with the value of SecondValueB to
	// create the SecondValue member of PDH_RAW_COUNTER. SecondValueA contains the low-order
	// bits.
	//
	// * SecondValueB(int): Combine this 32-bit value with the value of SecondValueA to
	// create the SecondValue member of PDH_RAW_COUNTER. SecondValueB contains the high
	// order bits.
	//
	// The CounterDetails table describes a specific counter on a particular computer. The
	// CounterDetails table defines the following fields:
	//
	// * CounterID(int, IDENTITY PRIMARY KEY): A unique identifier in the database that
	// maps to a specific counter name text string. This field is the primary key of this
	// table.
	//
	// * MachineName(varchar(1024), NOT NULL): The name of the computer that logged this
	// data set.
	//
	// * ObjectName(varchar(1024), NOT NULL): The name of the performance object ( f43f48aa-80a5-4b39-971e-7b3ac0bd9d0d#gt_8bb43a65-7a8c-4585-a7ed-23044772f8ca
	// ).
	//
	// * CounterName(varchar(1024), NOT NULL): The name of the counter.
	//
	// * CounterType(int, NOT NULL): The counter type.
	//
	// * DefaultScale(int, NOT NULL): The default scaling to be applied to the raw performance
	// counter ( f43f48aa-80a5-4b39-971e-7b3ac0bd9d0d#gt_35645a67-9e0b-4c05-b2d9-3b2b25f2beac
	// ) data.
	//
	// * InstanceName(varchar(1024)): The name of the counter instance.
	//
	// * InstanceIndex(int): The index number of the counter instance.
	//
	// * ParentName(varchar(1024)): Some counters are logically associated with others,
	// and are referred to as parents. For example, the parent of a thread is a process
	// and the parent of a logical disk ( f43f48aa-80a5-4b39-971e-7b3ac0bd9d0d#gt_c4133b2a-a990-4042-ba44-7fda3090f118
	// ) driver is a physical drive. This field contains the name of the parent. Either
	// the value in this field or the ParentObjectID field identifies a specific parent
	// instance. If the value in this field is NULL, the value in the ParentObjectID field
	// will be checked to identify the parent. If the values in both fields are NULL, the
	// counter does not have a parent.
	//
	// * ParentObjectID(int): The unique identifier of the parent. The value in either this
	// field or the ParentName field identifies a specific parent instance. If the value
	// in this field is NULL, the value in the ParentName field will be checked to identify
	// the parent.
	//
	// The DisplayToID table relates the user-friendly string displayed by the System Monitor
	// to the PLA-UID stored in the other tables. The DisplayToID table defines the following
	// fields:
	//
	// * GUID(uniqueidentifier, NOT NULL PRIMARY KEY): A PLA-UID generated for a log. This
	// field is the primary key of this table. Note that these do not correspond to the
	// values in: HKEY_LOCAL_MACHINE \SYSTEM \CurrentControlSet \Services \SysmonLog \Log
	// Queries\
	//
	// * RunID(int): Reserved for internal use.
	//
	// * DisplayString(varchar(1024), NOT NULL UNIQUE): Name of the log file as displayed
	// in the System Monitor.
	//
	// * LogStartTime(char(24)): Time the logging process started in yyyy-mm-dd hh:mm:ss:nnn
	// format.
	//
	// * LogStopTime(char(24)): Time the logging process stopped in yyyy-mm-dd hh:mm:ss:nnn
	// format. Multiple log files with the same DisplayString value can be differentiated
	// by using the value in this and the LogStartTime fields. The values in the LogStartTime
	// and LogStopTime fields also allow the total collection time to be accessed quickly.
	//
	// * NumberOfRecords(int): Number of samples stored in the table for each log collection.
	//
	// * MinutesToUTC(int): Value used to convert the row data stored in UTC time to local
	// time.
	//
	// * TimeZoneName(char(32)): Name of the time zone where the data was collected. If
	// collecting or analyzing relogged data from a file collected on systems in the user's
	// time zone, this field will state the location.
	FileFormatSQL FileFormat = 2
	// plaBinary:  Binary log file.
	FileFormatBinary FileFormat = 3
)

func (o FileFormat) String() string {
	switch o {
	case FileFormatCommaSeparated:
		return "FileFormatCommaSeparated"
	case FileFormatTabSeparated:
		return "FileFormatTabSeparated"
	case FileFormatSQL:
		return "FileFormatSQL"
	case FileFormatBinary:
		return "FileFormatBinary"
	}
	return "Invalid"
}

// AutoPathFormat type represents AutoPathFormat RPC enumeration.
//
// The AutoPathFormat enumeration defines the information to be appended to the file
// name or subdirectory name. Any combination of the bits is allowed; multiple bits
// specify strings to be appended to the file name. When a combination with than one
// of these bits is specified, the strings are appended in the following order: plaComputer,
// plaPattern, plaMonthDayHour, plaSerialNumber, plaYearDayOfYear, plaYearMonth, plaYearMonthDay,
// plaYearMonthDayHour, plaMonthDayHourMinute. Consequently, if all bits are set, the
// name is represented as follows: [plaComputer]base_name[plaPattern][plaMonthDayHour][plaSerialNumber][plaYearDayOfYear][plaYearMonth][plaYearMonthDay][plaYearMonthDayHour][plaMonthDayHourMinute].
type AutoPathFormat uint16

var (
	// plaNone:   Does not append any information to the name.
	AutoPathFormatNone AutoPathFormat = 0
	// plaPattern:  Adds a pattern specified in IDataCollectorSet::SubdirectoryFormatPattern
	// 3.2.4.1.32 or IDataCollector::FileNameFormatPattern 3.2.4.5.7 to the name.
	AutoPathFormatPattern AutoPathFormat = 1
	// plaComputer:  Prefixes the name with the computer name.
	AutoPathFormatComputer AutoPathFormat = 2
	// plaMonthDayHour:  Appends the month, day, and hour to the name in the form, MMddHH.
	AutoPathFormatMonthDayHour AutoPathFormat = 256
	// plaSerialNumber:  Appends the serial number specified in IDataCollectorSet::SerialNumber
	// to the subdirectory name in the form, NNNNNN.
	AutoPathFormatSerialNumber AutoPathFormat = 512
	// plaYearDayOfYear:  Appends the year and day of the year to the name in the form,
	// yyyyDDD.
	AutoPathFormatYearDayOfYear AutoPathFormat = 1024
	// plaYearMonth:  Appends the year and month to the name in the form, yyyyMM.
	AutoPathFormatYearMonth AutoPathFormat = 2048
	// plaYearMonthDay:  Appends the year, month, and day to the name in the form, yyyyMMdd.
	AutoPathFormatYearMonthDay AutoPathFormat = 4096
	// plaYearMonthDayHour:  Appends the year, month, day, and hour to the name in the
	// form, yyyyMMddHH.
	AutoPathFormatYearMonthDayHour AutoPathFormat = 8192
	// plaMonthDayHourMinute:  Appends the month, day, hour, and minute to the name in
	// the form, MMddHHmm.
	AutoPathFormatMonthDayHourMinute AutoPathFormat = 16384
)

func (o AutoPathFormat) String() string {
	switch o {
	case AutoPathFormatNone:
		return "AutoPathFormatNone"
	case AutoPathFormatPattern:
		return "AutoPathFormatPattern"
	case AutoPathFormatComputer:
		return "AutoPathFormatComputer"
	case AutoPathFormatMonthDayHour:
		return "AutoPathFormatMonthDayHour"
	case AutoPathFormatSerialNumber:
		return "AutoPathFormatSerialNumber"
	case AutoPathFormatYearDayOfYear:
		return "AutoPathFormatYearDayOfYear"
	case AutoPathFormatYearMonth:
		return "AutoPathFormatYearMonth"
	case AutoPathFormatYearMonthDay:
		return "AutoPathFormatYearMonthDay"
	case AutoPathFormatYearMonthDayHour:
		return "AutoPathFormatYearMonthDayHour"
	case AutoPathFormatMonthDayHourMinute:
		return "AutoPathFormatMonthDayHourMinute"
	}
	return "Invalid"
}

// DataCollectorSetStatus type represents DataCollectorSetStatus RPC enumeration.
//
// The DataCollectorSetStatus enumeration defines the running status of the data collector
// set.
type DataCollectorSetStatus uint16

var (
	// plaStopped:  The data collector set is stopped.
	DataCollectorSetStatusStopped DataCollectorSetStatus = 0
	// plaRunning:  The data collector set is running.
	DataCollectorSetStatusRunning DataCollectorSetStatus = 1
	// plaCompiling:  The data collector set is performing data management (see section
	// 3.2.4.2). A running data collector set transitions from running to compiling if the
	// data manager is enabled.
	DataCollectorSetStatusCompiling DataCollectorSetStatus = 2
	// plaPending:  Not used.
	DataCollectorSetStatusPending DataCollectorSetStatus = 3
	// plaUndefined:  Cannot determine the status but no error has occurred. Typically,
	// this status is set for boot trace sessions.
	DataCollectorSetStatusUndefined DataCollectorSetStatus = 4
)

func (o DataCollectorSetStatus) String() string {
	switch o {
	case DataCollectorSetStatusStopped:
		return "DataCollectorSetStatusStopped"
	case DataCollectorSetStatusRunning:
		return "DataCollectorSetStatusRunning"
	case DataCollectorSetStatusCompiling:
		return "DataCollectorSetStatusCompiling"
	case DataCollectorSetStatusPending:
		return "DataCollectorSetStatusPending"
	case DataCollectorSetStatusUndefined:
		return "DataCollectorSetStatusUndefined"
	}
	return "Invalid"
}

// ClockType type represents ClockType RPC enumeration.
//
// The ClockType enumeration defines the clock resolution to use when tracing events.
type ClockType uint16

var (
	// plaTimeStamp:  Use the raw (unconverted) time stamp.
	ClockTypeTimestamp ClockType = 0
	// plaPerformance:  Query performance counter (QPC). Provides a high-resolution (100
	// nanoseconds) time stamp that is more expensive to retrieve.
	ClockTypePerformance ClockType = 1
	// plaSystem:  System time. Provides a low-resolution (10 milliseconds) time stamp
	// that is less expensive to retrieve.
	ClockTypeSystem ClockType = 2
	// plaCycle:  CPU cycle counter. MAY provide the highest resolution time stamp and
	// is the least expensive to retrieve. However, the CPU counter is unreliable and its
	// use is not recommended.
	ClockTypeCycle ClockType = 3
)

func (o ClockType) String() string {
	switch o {
	case ClockTypeTimestamp:
		return "ClockTypeTimestamp"
	case ClockTypePerformance:
		return "ClockTypePerformance"
	case ClockTypeSystem:
		return "ClockTypeSystem"
	case ClockTypeCycle:
		return "ClockTypeCycle"
	}
	return "Invalid"
}

// StreamMode type represents StreamMode RPC enumeration.
//
// The StreamMode enumeration defines where the trace events are delivered.
type StreamMode uint16

var (
	// plaFile:  Writes the trace events to a log file.
	StreamModeFile StreamMode = 1
	// plaRealTime:  Delivers the trace events to a real time consumer.
	StreamModeRealTime StreamMode = 2
	// plaBoth:  Writes the trace events to a log file and delivers them to a real-time
	// consumer.
	StreamModeBoth StreamMode = 3
	// plaBuffering:   Keeps events in a circular buffer in memory only. For more information,
	// see the EVENT_TRACE_BUFFERING_MODE logging mode in [MSDN-LMC].
	StreamModeBuffering StreamMode = 4
)

func (o StreamMode) String() string {
	switch o {
	case StreamModeFile:
		return "StreamModeFile"
	case StreamModeRealTime:
		return "StreamModeRealTime"
	case StreamModeBoth:
		return "StreamModeBoth"
	case StreamModeBuffering:
		return "StreamModeBuffering"
	}
	return "Invalid"
}

// CommitMode type represents CommitMode RPC enumeration.
//
// The CommitMode enumeration defines the type of actions to be performed when the changes
// are committed to the data collector set. Any combination of bits MUST be allowed.
type CommitMode uint16

var (
	// plaCreateNew:  For a persistent data collector set, save it to storage. The set
	// MUST not have existed previously on storage.
	CommitModeCreateNew CommitMode = 1
	// plaModify:  Update a previously committed data collector set.
	CommitModeModify CommitMode = 2
	// plaCreateOrModify:  For a persistent data collector set, save it to storage. If
	// the set already exists, the PLA Protocol will update it.
	CommitModeCreateOrModify CommitMode = 3
	// plaUpdateRunningInstance:  If the data collector set is running, apply the updated
	// property values to it.
	CommitModeUpdateRunningInstance CommitMode = 16
	// plaFlushTrace:  If multiple data collector sets are running, flush the event trace
	// data collectors memory buffers to storage or real-time consumers.
	CommitModeFlushTrace CommitMode = 32
	// plaValidateOnly:  Perform validation only on the data collector set.
	CommitModeValidateOnly CommitMode = 4096
)

func (o CommitMode) String() string {
	switch o {
	case CommitModeCreateNew:
		return "CommitModeCreateNew"
	case CommitModeModify:
		return "CommitModeModify"
	case CommitModeCreateOrModify:
		return "CommitModeCreateOrModify"
	case CommitModeUpdateRunningInstance:
		return "CommitModeUpdateRunningInstance"
	case CommitModeFlushTrace:
		return "CommitModeFlushTrace"
	case CommitModeValidateOnly:
		return "CommitModeValidateOnly"
	}
	return "Invalid"
}

// ValueMapType type represents ValueMapType RPC enumeration.
//
// The ValueMapType enumeration defines a value map type. A value map defines a named-value
// pair. A value map can be used in different ways. A value map type defines which way
// the value map is to be used; each type has a different way of evaluating the "value"
// of the "value map" based on the "values" of each individual "value map item".
type ValueMapType uint16

var (
	// plaIndex:  Only one item in the collection can be enabled. The enabled item is the
	// value of IValueMap::Value. If more than one is enabled, the first enabled item MUST
	// be used as the value.
	ValueMapTypeIndex ValueMapType = 1
	// plaFlag:  One or more items in the collection can be enabled. The enabled items
	// in the collection are combined together by using the bitwise OR operation to become
	// the value of IValueMap::Value.
	ValueMapTypeFlag ValueMapType = 2
	// plaFlagArray:  One or more items in the collection can be enabled. An item in the
	// collection represents a 32-bit unsigned value (ULONG). The enabled items are not
	// combined together as they are for the plaFlag type, but rather each item can be retrieved
	// separately.<2>
	ValueMapTypeFlagArray ValueMapType = 3
	// plaValidation:  The collection contains a list of HRESULT values that are returned
	// in an IValueMap by the validation process. The validation process occurs when IDataCollectorSet::Commit
	// is called. In the validation process, the PLA Protocol analyzes the values of all
	// the properties in the IDataCollectorSet, including the values of the IDataCollectors
	// contained in the IDataCollectorSet and inserts a ValueMapItem into the ValueMap for
	// any property that is problematic. The ValueMapItem holds the name of the property
	// and the HRESULT describing why it is problematic. The following codes can be set
	// in a validation ValueMap:
	//
	//	+------------------------------------------+----------------------------------------------------------------------------------+
	//	|                                          |                                                                                  |
	//	|                NAME/VALUE                |                                   DESCRIPTION                                    |
	//	|                                          |                                                                                  |
	//	+------------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------------+----------------------------------------------------------------------------------+
	//	| PLA_S_PROPERTY_IGNORED/(0x00300100)      | This value can be returned anytime the value of a property is being ignored      |
	//	|                                          | by this implementation of the protocol. The code is intended to inform the       |
	//	|                                          | client when a property is not needed or supported by an implementation. The      |
	//	|                                          | following is a list of properties for the different types of data collectors     |
	//	|                                          | (that are encapsulated within a data collector set) that MUST be ignored         |
	//	|                                          | by the server when the client calls the Commit method on the data collector      |
	//	|                                          | set; the server MUST return the property name and PLA_S_PROPERTY_IGNORED         |
	//	|                                          | in the IValueMapItem for each property that it ignored. Note that certain        |
	//	|                                          | properties can pertain to the base DataCollector interface. If there is no       |
	//	|                                          | task specified, the TaskArguments property of the DataCollectorSet MUST be       |
	//	|                                          | ignored. If the SubdirectoryFormat property is not set to plaPattern, the        |
	//	|                                          | SubdirectoryFormatPattern property is ignored. For the base DataCollector,       |
	//	|                                          | if the SegmentMaxSize property is zero and LogCircular is false, LogCircular     |
	//	|                                          | is ignored. If the LogOverwrite property is true or the LogCircular is           |
	//	|                                          | true, and the LogAppend property is false, LogAppend is ignored. For the         |
	//	|                                          | AlertDataCollector data collector, the following properties MUST be ignored:     |
	//	|                                          | FileName, FileNameFormat, FileNameFormatPattern, LogAppend, LogCircular, and     |
	//	|                                          | LogOverwrite. For the ApiTracingDataCollector data collector, the following      |
	//	|                                          | properties MUST be ignored: FileNameFormat, FileNameFormatPattern, LogAppend,    |
	//	|                                          | and LogOverwrite. For the ApiTracingDataCollector data collector, the            |
	//	|                                          | following properties MUST be ignored: FileName and LogCircular. <3> For the      |
	//	|                                          | ConfigurationDataCollector data collector, the following properties MUST be      |
	//	|                                          | ignored: LogCircular and LogAppend. For the PerformanceCounterDataCollector      |
	//	|                                          | data collector, the following properties MUST be ignored if the LogFileFormat    |
	//	|                                          | property is set to plaSql: LogCircular, LogOverwrite, and LogAppend. LogAppend   |
	//	|                                          | is also returned if the LogFileFormat property is set to plaTabSeparated         |
	//	|                                          | or plaCommaSeparated. For the TraceDataCollector data collector, the             |
	//	|                                          | following properties MUST be ignored if the StreamMode is not plaFile:           |
	//	|                                          | FileName, LogAppend, LogCircular, LogOverwrite, FileNameFormat, and              |
	//	|                                          | FileNameFormatPattern. For TraceSession, the following properties MUST be        |
	//	|                                          | ignored: RootPath, Duration, Description, Keywords, Segment, SegmentMaxDuration, |
	//	|                                          | SerialNumber, Subdirectory, SubdirectoryFormat, SubdirectoryFormatPattern,       |
	//	|                                          | Task, Schedules, TraceDataCollector[1]/FileNameFormat,                           |
	//	|                                          | TraceDataCollector[1]/FileNameFormatPattern, and                                 |
	//	|                                          | TraceDataCollector[1]/LogOverwrite. If IDataCollectorSet::Commit() with the      |
	//	|                                          | flag plaUpdateRunningInstance set is called on an IDataCollectorSet of type      |
	//	|                                          | TraceSession, as specified in section 3.2.1, the following properties MUST be    |
	//	|                                          | ignored: TraceDataCollector[1]/BufferSize, TraceDataCollector[1]/MinimumBuffers, |
	//	|                                          | TraceDataCollector[1]/NumberOfBuffers, TraceDataCollector[1]/ClockType,          |
	//	|                                          | TraceDataCollector[1]/ProcessMode, TraceDataCollector[1]/PreallocateFile, and    |
	//	|                                          | SegmentMaxSize.                                                                  |
	//	+------------------------------------------+----------------------------------------------------------------------------------+
	//	| PLA_E_PROPERTY_CONFLICT/(0x80300101)     | This value can be returned anytime two properties are in conflict. This          |
	//	|                                          | code can be returned for the following properties under the following            |
	//	|                                          | conditions: IApiTracingDataCollector::ExePath: Returned when ExePath is equal    |
	//	|                                          | to the empty string. IDataCollector::FileNameFormatPattern: Returned when        |
	//	|                                          | IDataCollector::FileNameFormat is equal to plaPattern and FileNameFormatPattern  |
	//	|                                          | is equal to the empty string. IDataCollector::LogCircular: Returned              |
	//	|                                          | when IDataCollectorSet::SegmentMaxSize is equal to 0 and LogCircular             |
	//	|                                          | is equal to true. IDataCollector::LogAppend: Returned when either                |
	//	|                                          | IDataCollector::LogCircular is true or IDataCollector::LogOverwrite is true      |
	//	|                                          | and LogAppend is true. IPerformanceCounterDataCollector::DataSourceName:         |
	//	|                                          | Returned when DataSourceName is equal to the empty string and                    |
	//	|                                          | IPerformanceCounterDataCollector::LogFileFormat is set to plaSql.                |
	//	|                                          | ITraceDataCollector::MaximumBuffers: Returned when MaximumBuffers is less than   |
	//	|                                          | ITraceDataCollector::MinimumBuffers.<4> ITraceDataCollector::TraceDataProviders: |
	//	|                                          | Returned if ITraceDataProviderCollection::Count is greater than 1 and            |
	//	|                                          | isKernelTrace is TRUE. ITraceDataCollector::Guid: Returned if isKernelTrace is   |
	//	|                                          | true and the specific PLA-UID does not match the kernel PLA-UID.<5>              |
	//	+------------------------------------------+----------------------------------------------------------------------------------+
	//	| PLA_E_EXE_FULL_PATH_REQUIRED/0x8030010E) | This value can be returned anytime a relative path, with respect to the current  |
	//	|                                          | working directory, to a file is provided when a full path is required. This code |
	//	|                                          | can be returned for the following properties under the following conditions:     |
	//	|                                          | IApiTracingDataCollector::ExePath: Returned when the provided path is relative   |
	//	|                                          | to the current working directory instead of absolute.                            |
	//	+------------------------------------------+----------------------------------------------------------------------------------+
	//	| PLA_E_EXE_PATH_NOT_VALID/(0x80300108)    | This value can be returned when the executable referenced by the ExePath         |
	//	|                                          | property for an IApiTracingDataCollector does not exist. This code can           |
	//	|                                          | be returned for the following properties under the following conditions:         |
	//	|                                          | IApiTracingDataCollector::ExePath: Returned when the executable referenced by    |
	//	|                                          | the ExePath property does not exist.                                             |
	//	+------------------------------------------+----------------------------------------------------------------------------------+
	//	| PLA_E_NETWORK_EXE_NOT_VALID/(0x80300106  | This value can be returned when the executable referenced by the ExePath is on   |
	//	|                                          | a remote machine. This code can be returned for the following properties under   |
	//	|                                          | the following conditions: IApiTracingDataCollector::ExePath: Returned when the   |
	//	|                                          | executable referenced by the ExePath is on a remote machine.                     |
	//	+------------------------------------------+----------------------------------------------------------------------------------+
	ValueMapTypeValidation ValueMapType = 4
)

func (o ValueMapType) String() string {
	switch o {
	case ValueMapTypeIndex:
		return "ValueMapTypeIndex"
	case ValueMapTypeFlag:
		return "ValueMapTypeFlag"
	case ValueMapTypeFlagArray:
		return "ValueMapTypeFlagArray"
	case ValueMapTypeValidation:
		return "ValueMapTypeValidation"
	}
	return "Invalid"
}

// WeekDays type represents WeekDays RPC enumeration.
//
// The WeekDays enumeration defines the days of the week on which to run the data collector
// set. Any combination of the bits MUST be allowed.
type WeekDays uint16

var (
	// plaRunOnce:  Run only once on the given start date and time.
	WeekDaysRunOnce WeekDays = 0
	// plaSunday:  Run on Sunday.
	WeekDaysSunday WeekDays = 1
	// plaMonday:   Run on Monday.
	WeekDaysMonday WeekDays = 2
	// plaTuesday:  Run on Tuesday.
	WeekDaysTuesday WeekDays = 4
	// plaWednesday:  Run on Wednesday.
	WeekDaysWednesday WeekDays = 8
	// plaThursday:  Run on Thursday.
	WeekDaysThursday WeekDays = 16
	// plaFriday:  Run on Friday.
	WeekDaysFriday WeekDays = 32
	// plaSaturday:  Run on Saturday.
	WeekDaysSaturday WeekDays = 64
	// plaEveryday:  Run every day of the week.
	WeekDaysEveryday WeekDays = 127
)

func (o WeekDays) String() string {
	switch o {
	case WeekDaysRunOnce:
		return "WeekDaysRunOnce"
	case WeekDaysSunday:
		return "WeekDaysSunday"
	case WeekDaysMonday:
		return "WeekDaysMonday"
	case WeekDaysTuesday:
		return "WeekDaysTuesday"
	case WeekDaysWednesday:
		return "WeekDaysWednesday"
	case WeekDaysThursday:
		return "WeekDaysThursday"
	case WeekDaysFriday:
		return "WeekDaysFriday"
	case WeekDaysSaturday:
		return "WeekDaysSaturday"
	case WeekDaysEveryday:
		return "WeekDaysEveryday"
	}
	return "Invalid"
}

// ResourcePolicy type represents ResourcePolicy RPC enumeration.
//
// The ResourcePolicy enumeration defines the order in which folders are deleted when
// one of the disk resource limits is exceeded.
type ResourcePolicy uint16

var (
	// plaDeleteLargest:  Deletes the largest folders first.
	ResourcePolicyDeleteLargest ResourcePolicy = 0
	// plaDeleteOldest:  Deletes the oldest folders first.
	ResourcePolicyDeleteOldest ResourcePolicy = 1
)

func (o ResourcePolicy) String() string {
	switch o {
	case ResourcePolicyDeleteLargest:
		return "ResourcePolicyDeleteLargest"
	case ResourcePolicyDeleteOldest:
		return "ResourcePolicyDeleteOldest"
	}
	return "Invalid"
}

// DataManagerSteps type represents DataManagerSteps RPC enumeration.
//
// The DataManagerSteps enumeration defines the actions that the data manager takes
// when it runs. Any combination of the bits are allowed.
type DataManagerSteps uint16

var (
	// plaCreateReport:  Creates a report if data is available. The file name MUST be IDataManager::RuleTargetFileName,
	// and IDataManager::ReportSchema can be used to customize the way the report is created.
	// This value indicates the run of TraceRpt.exe by using as input all of the binary
	// performance files (.blg) and event trace files (.etl) in the collection.
	DataManagerStepsCreateReport DataManagerSteps = 1
	// plaRunRules:  If a report exists, the PLA Protocol MUST apply the rules specified
	// in IDataManager::Rules to the report. The IDataManager::RuleTargetFileName(Get) returns
	// the name of the file to which the rules are applied.
	DataManagerStepsRunRules DataManagerSteps = 2
	// plaCreateHtml:  Converts the XML file obtained by IDataManager::RuleTargetFileName(Get)
	// to HTML format. The HTML format is written to the file specified in IDataManager::ReportFileName.
	DataManagerStepsCreateHTML DataManagerSteps = 4
	// plaFolderActions:  Apply the folder actions obtained by IDataManager::FolderActions(Get)
	// to all folders defined in the collection.
	DataManagerStepsFolderActions DataManagerSteps = 8
	// plaResourceFreeing:  If IDataManager::MaxFolderCount, IDataManager::MaxSize, or
	// MinFreeDisk exceeds its limit, the PLA Protocol MUST apply the resource policy specified
	// in IDataManager::ResourcePolicy.
	DataManagerStepsResourceFreeing DataManagerSteps = 16
)

func (o DataManagerSteps) String() string {
	switch o {
	case DataManagerStepsCreateReport:
		return "DataManagerStepsCreateReport"
	case DataManagerStepsRunRules:
		return "DataManagerStepsRunRules"
	case DataManagerStepsCreateHTML:
		return "DataManagerStepsCreateHTML"
	case DataManagerStepsFolderActions:
		return "DataManagerStepsFolderActions"
	case DataManagerStepsResourceFreeing:
		return "DataManagerStepsResourceFreeing"
	}
	return "Invalid"
}

// FolderActionSteps type represents FolderActionSteps RPC enumeration.
//
// The FolderActionSteps enumeration defines the action that the data manager takes
// when both the age and size limits are met. Any combination of the bits MUST be allowed.
type FolderActionSteps uint16

var (
	// plaCreateCab:  Creates a cabinet file. The name of the cabinet file is <name of
	// the subfolder>.cab. For example, if the name of the subfolder was "MyFolder", the
	// cab file would be named "MyFolder.cab". The name of the subfolder is specified by
	// the combination of the Subdirectory, SubdirectoryFormat, and SubdirectoryFormatPattern
	// properties of the IDataCollectorSet. The Subdirectory property provides the base
	// name for the Subfolder, the SubdirectoryFormat property specifies the suffix and
	// prefix that will be appended and prepended to the base name, and the SubdirectoryFormatPattern
	// specifies the pattern that will be used in the suffix. The SubdirectoryFormat is
	// specified in section 2.2.2.1. The SubdirectoryFormatPattern is specified in section
	// 2.2.3.1.
	FolderActionStepsCreateCab FolderActionSteps = 1
	// plaDeleteData:  Deletes all files in the folder, except the report and cabinet file.
	FolderActionStepsDeleteData FolderActionSteps = 2
	// plaSendCab:  Sends the cabinet file to the location specified in the IFolderAction::SendCabTo
	// property.
	FolderActionStepsSendCab FolderActionSteps = 4
	// plaDeleteCab:  Deletes the cabinet file.
	FolderActionStepsDeleteCab FolderActionSteps = 8
	// plaDeleteReport:  Deletes the report file.
	FolderActionStepsDeleteReport FolderActionSteps = 16
)

func (o FolderActionSteps) String() string {
	switch o {
	case FolderActionStepsCreateCab:
		return "FolderActionStepsCreateCab"
	case FolderActionStepsDeleteData:
		return "FolderActionStepsDeleteData"
	case FolderActionStepsSendCab:
		return "FolderActionStepsSendCab"
	case FolderActionStepsDeleteCab:
		return "FolderActionStepsDeleteCab"
	case FolderActionStepsDeleteReport:
		return "FolderActionStepsDeleteReport"
	}
	return "Invalid"
}

// PerformanceCounterDataCollector structure represents IPerformanceCounterDataCollector RPC structure.
//
// The IPerformanceCounterDataCollector interface is used to specify the performance
// counters to query and the log file to which the counter data is written.
//
// The following properties MUST be implemented by the objects that implement the IPerformanceCounterDataCollector
// interface.
//
//	+---------------------+------------+----------------------------------------------------------------------------------+
//	|                     |            |                                                                                  |
//	|      PROPERTY       | READ/WRITE |                                   DESCRIPTION                                    |
//	|                     |            |                                                                                  |
//	+---------------------+------------+----------------------------------------------------------------------------------+
//	+---------------------+------------+----------------------------------------------------------------------------------+
//	| DataSourceName      | RW         | The data source name if the log file is an SQL log file.                         |
//	+---------------------+------------+----------------------------------------------------------------------------------+
//	| LogFileFormat       | RW         | The format in which data MUST be stored. The format is specified by the          |
//	|                     |            | FileFormat enumeration.                                                          |
//	+---------------------+------------+----------------------------------------------------------------------------------+
//	| PerformanceCounters | RW         | List of performance counters to be collected.                                    |
//	+---------------------+------------+----------------------------------------------------------------------------------+
//	| SampleInterval      | RW         | The time, in seconds, between two consecutive samples. The default is 15         |
//	|                     |            | seconds. The minimum interval is 1 second. There is no maximum interval.         |
//	+---------------------+------------+----------------------------------------------------------------------------------+
//	| SegmentMaxRecords   | RW         | Maximum number of samples to log in a segment. If set to 0, there is no segment  |
//	|                     |            | record limit. Any unsigned long is a valid value for this property.              |
//	+---------------------+------------+----------------------------------------------------------------------------------+
//
// A data collector can be represented as an XML file, which can be used to serialize
// (using Xml (Get) 3.2.4.5.21) and deserialize (using SetXml 3.2.4.5.22) it. The format
// of the XML that defines a performance counter data collector is as follows(the full
// XML specification of the data collector set XML is in section 3.2.4.19):
//
// The XML given above does not show the property elements inherited from IDataCollector
// that also need to be specified.
//
// Methods in RPC Opnum Order
//
//	+---------------------------+-------------------------------------------------------+
//	|                           |                                                       |
//	|          METHOD           |                      DESCRIPTION                      |
//	|                           |                                                       |
//	+---------------------------+-------------------------------------------------------+
//	+---------------------------+-------------------------------------------------------+
//	| DataSourceName (Get)      | Retrieves the DataSourceName property. Opnum: 32      |
//	+---------------------------+-------------------------------------------------------+
//	| DataSourceName (Put)      | Sets the DataSourceName property. Opnum: 33           |
//	+---------------------------+-------------------------------------------------------+
//	| PerformanceCounters (Get) | Retrieves the PerformanceCounters property. Opnum: 34 |
//	+---------------------------+-------------------------------------------------------+
//	| PerformanceCounters (Put) | Sets the PerformanceCounters property. Opnum: 35      |
//	+---------------------------+-------------------------------------------------------+
//	| LogFileFormat (Get)       | Retrieves the LogFileFormat property. Opnum: 36       |
//	+---------------------------+-------------------------------------------------------+
//	| LogFileFormat (Put)       | Sets the LogFileFormat property. Opnum: 37            |
//	+---------------------------+-------------------------------------------------------+
//	| SampleInterval (Get)      | Retrieves the SampleInterval property. Opnum: 38      |
//	+---------------------------+-------------------------------------------------------+
//	| SampleInterval (Put)      | Sets the SampleInterval property. Opnum: 39           |
//	+---------------------------+-------------------------------------------------------+
//	| SegmentMaxRecords (Get)   | Retrieves the SegmentMaxRecords property. Opnum: 40   |
//	+---------------------------+-------------------------------------------------------+
//	| SegmentMaxRecords (Put)   | Sets the SegmentMaxRecords property. Opnum: 41        |
//	+---------------------------+-------------------------------------------------------+
//
// Opnums 0, 1, and 2 are reserved for the IUnknown interface. Opnums 3, 4, 5, and 6
// are reserved for the IDispatch interface. Opnums 7–31 are used by IDataCollector.
type PerformanceCounterDataCollector dcom.InterfacePointer

func (o *PerformanceCounterDataCollector) InterfacePointer() *dcom.InterfacePointer {
	return (*dcom.InterfacePointer)(o)
}

func (o *PerformanceCounterDataCollector) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if o.Data != nil && o.DataCount == 0 {
		o.DataCount = uint32(len(o.Data))
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}

func (o *PerformanceCounterDataCollector) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *PerformanceCounterDataCollector) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *PerformanceCounterDataCollector) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// FolderActionCollection structure represents IFolderActionCollection RPC structure.
//
// The IFolderActionCollection interface is used to manage a collection of FolderAction
// objects.
//
// The following properties MUST be implemented by the objects that implement the IFolderActionCollection
// interface.
//
//	+----------+------------+----------------------------------------------------------------------------------+
//	|          |            |                                                                                  |
//	| PROPERTY | READ/WRITE |                                   DESCRIPTION                                    |
//	|          |            |                                                                                  |
//	+----------+------------+----------------------------------------------------------------------------------+
//	+----------+------------+----------------------------------------------------------------------------------+
//	| _NewEnum | R          | An enumeration object of type IEnumVariant containing a snapshot of the          |
//	|          |            | IFolderAction in this collection. The enumeration object is specified in         |
//	|          |            | [MS-OAUT] section 3.3                                                            |
//	+----------+------------+----------------------------------------------------------------------------------+
//	| Count    | R          | Number of folder actions in this collection.                                     |
//	+----------+------------+----------------------------------------------------------------------------------+
//	| Item     | R          | Retrieves the requested folder action from the collection.                       |
//	+----------+------------+----------------------------------------------------------------------------------+
//
// Methods in RPC Opnum Order
//
//	+--------------------+--------------------------------------------------------------+
//	|                    |                                                              |
//	|       METHOD       |                         DESCRIPTION                          |
//	|                    |                                                              |
//	+--------------------+--------------------------------------------------------------+
//	+--------------------+--------------------------------------------------------------+
//	| Count (Get)        | Retrieves the Count property. Opnum: 7                       |
//	+--------------------+--------------------------------------------------------------+
//	| Item (Get)         | Retrieves the Item property. Opnum: 8                        |
//	+--------------------+--------------------------------------------------------------+
//	| _NewEnum (Get)     | Retrieves the NewEnum property. Opnum: 9                     |
//	+--------------------+--------------------------------------------------------------+
//	| Add                | Adds a folder action to the collection. Opnum: 10            |
//	+--------------------+--------------------------------------------------------------+
//	| Remove             | Removes a folder action from the collection. Opnum: 11       |
//	+--------------------+--------------------------------------------------------------+
//	| Clear              | Removes all folder actions from the collection. Opnum: 12    |
//	+--------------------+--------------------------------------------------------------+
//	| AddRange           | Adds one or more folder actions to the collection. Opnum: 13 |
//	+--------------------+--------------------------------------------------------------+
//	| CreateFolderAction | Creates a folder action object. Opnum: 14                    |
//	+--------------------+--------------------------------------------------------------+
//
// Opnums 0, 1, and 2 are reserved for the IUnknown interface. Opnums 3, 4, 5, and 6
// are reserved for the IDispatch interface.
type FolderActionCollection dcom.InterfacePointer

func (o *FolderActionCollection) InterfacePointer() *dcom.InterfacePointer {
	return (*dcom.InterfacePointer)(o)
}

func (o *FolderActionCollection) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if o.Data != nil && o.DataCount == 0 {
		o.DataCount = uint32(len(o.Data))
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}

func (o *FolderActionCollection) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *FolderActionCollection) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *FolderActionCollection) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// ValueMapItem structure represents IValueMapItem RPC structure.
//
// The IValueMapItem interface is used to define a named-value pair.
//
// The following properties MUST be implemented by the objects that implement the IValueMapItem
// interface.
//
// The following is an XML representation of a ValueMapItem.
//
//	+--------------+------------+----------------------------------------------------------------------------------+
//	|              |            |                                                                                  |
//	|   PROPERTY   | READ/WRITE |                                   DESCRIPTION                                    |
//	|              |            |                                                                                  |
//	+--------------+------------+----------------------------------------------------------------------------------+
//	+--------------+------------+----------------------------------------------------------------------------------+
//	| Description  | RW         | Specifies the description of the item.                                           |
//	+--------------+------------+----------------------------------------------------------------------------------+
//	| Enabled      | RW         | Specifies whether or not the item is enabled. If an item is not enabled, its     |
//	|              |            | Value property will be ignored. In typical usage, a ValueMap will contain a      |
//	|              |            | ValueMapItem for each of the multiple possible settings of the property that the |
//	|              |            | ValueMap is passed to (such as an entry for every Keyword or every Level that    |
//	|              |            | can be used by an ITraceDataProvider). This Enabled property indicates whether   |
//	|              |            | or not the ValueMapItem will be used. Essentially, the ValueMap contains a       |
//	|              |            | ValueMapItem for each of the multiple possible options, and the Enabled property |
//	|              |            | indicates which options are actually selected.                                   |
//	+--------------+------------+----------------------------------------------------------------------------------+
//	| Key          | RW         | Specifies the name of the item. The name of the item can be any non-empty        |
//	|              |            | BSTR. The semantics of the key depend on the ValueMapType property, specified    |
//	|              |            | in section 3.2.4.18, of the IValueMap to which this IValueMapItem belongs. The   |
//	|              |            | over-the-wire transmission of a BSTR is specified in [MS-OAUT] section 2.2.23.   |
//	|              |            | The Performance Logs and Alerts Protocol does not have any predefined Key values |
//	|              |            | or semantic definitions. The only condition Performance Logs and Alerts Protocol |
//	|              |            | places on the Key is that it cannot be an empty BSTR.<36>                        |
//	+--------------+------------+----------------------------------------------------------------------------------+
//	| Value        | RW         | Specifies the value of the item. The value is stored in a VARIANT. The format    |
//	|              |            | and over the wire transmission of a VARIANT is specified in [MS-OAUT] section    |
//	|              |            | 2.2.29. Any VARIANT is a legal value for this property. The ValueMapItem is a    |
//	|              |            | generic container and the semantics of the Value property depend on what it is   |
//	|              |            | being used to contain. The ValueMapType property contains information regarding  |
//	|              |            | how the Value property will be interpreted. <37>                                 |
//	+--------------+------------+----------------------------------------------------------------------------------+
//	| ValueMapType | RW         | Specifies the type of ValueMap in which the ValueMapItem will be inserted.       |
//	|              |            | Information on the different types of ValueMaps are specified in section         |
//	|              |            | 2.2.2.11.                                                                        |
//	+--------------+------------+----------------------------------------------------------------------------------+
//
// The following is an XML representation of a ValueMapItem; please see section 3.2.4.19
// the XML layout of entire data collector set element.
//
// Methods in RPC Opnum Order
//
//	+--------------------+------------------------------------------------+
//	|                    |                                                |
//	|       METHOD       |                  DESCRIPTION                   |
//	|                    |                                                |
//	+--------------------+------------------------------------------------+
//	+--------------------+------------------------------------------------+
//	| Description (Get)  | Retrieves the Description property. Opnum: 7   |
//	+--------------------+------------------------------------------------+
//	| Description (Put)  | Sets the Description property. Opnum: 8        |
//	+--------------------+------------------------------------------------+
//	| Enabled (Get)      | Retrieves the Enabled property. Opnum: 9       |
//	+--------------------+------------------------------------------------+
//	| Enabled (Put)      | Sets the Enabled property. Opnum: 10           |
//	+--------------------+------------------------------------------------+
//	| Key (Get)          | Retrieves the Key property. Opnum: 11          |
//	+--------------------+------------------------------------------------+
//	| Key (Put)          | Sets the Key property. Opnum: 12               |
//	+--------------------+------------------------------------------------+
//	| Value (Get)        | Retrieves the Value property. Opnum: 13        |
//	+--------------------+------------------------------------------------+
//	| Value (Put)        | Sets the Value property. Opnum: 14             |
//	+--------------------+------------------------------------------------+
//	| ValueMapType (Get) | Retrieves the ValueMapType property. Opnum: 15 |
//	+--------------------+------------------------------------------------+
//	| ValueMapType (Put) | Sets the ValueMapType property. Opnum: 16      |
//	+--------------------+------------------------------------------------+
//
// Opnums 0, 1, and 2 are reserved for the IUnknown interface. Opnums 3, 4, 5, and 6
// are reserved for the IDispatch interface.
type ValueMapItem dcom.InterfacePointer

func (o *ValueMapItem) InterfacePointer() *dcom.InterfacePointer { return (*dcom.InterfacePointer)(o) }

func (o *ValueMapItem) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if o.Data != nil && o.DataCount == 0 {
		o.DataCount = uint32(len(o.Data))
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}

func (o *ValueMapItem) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *ValueMapItem) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *ValueMapItem) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// ConfigurationDataCollector structure represents IConfigurationDataCollector RPC structure.
//
// The IConfigurationDataCollector is used to collect computer settings.
//
// The following properties MUST be implemented by the objects that implement the IConfigurationDataCollector
// interface.
//
//	+---------------------------+------------+----------------------------------------------------------------------------------+
//	|                           |            |                                                                                  |
//	|         PROPERTY          | READ/WRITE |                                   DESCRIPTION                                    |
//	|                           |            |                                                                                  |
//	+---------------------------+------------+----------------------------------------------------------------------------------+
//	+---------------------------+------------+----------------------------------------------------------------------------------+
//	| FileMaxCount              | RW         | Specifies the maximum number of files to collect. If set to zero or not set,     |
//	|                           |            | there is no maximum. Any unsigned long is a valid value for this property.       |
//	+---------------------------+------------+----------------------------------------------------------------------------------+
//	| FileMaxRecursiveDepth     | RW         | Specifies the maximum depth in a file system hierarchy that a recursive file     |
//	|                           |            | collection MUST attempt. If set to zero, the maximum depth is 30. Any unsigned   |
//	|                           |            | long is a valid value for this property.                                         |
//	+---------------------------+------------+----------------------------------------------------------------------------------+
//	| FileMaxTotalSize          | RW         | Specifies the maximum size, in megabytes, of all files to collect. If set to     |
//	|                           |            | zero or not set, there is no maximum size. Any unsigned long is a valid value    |
//	|                           |            | for this property.                                                               |
//	+---------------------------+------------+----------------------------------------------------------------------------------+
//	| Files                     | RW         | List of paths to files which will be copied to the output directory. Any         |
//	|                           |            | arbitrary files can be specified. Absolute, relative, and UncPaths are           |
//	|                           |            | supported. The '*' and '?' wildcards can be used, and collection can be made     |
//	|                           |            | recursive by using two backslashes ("\\") for the last folder delimiter. If a    |
//	|                           |            | specified file is not found, an error is added to the output file but collection |
//	|                           |            | continues.                                                                       |
//	+---------------------------+------------+----------------------------------------------------------------------------------+
//	| ManagementQueries         | RW         | List of Windows Management Instrumentation (WMI) queries whose results MUST      |
//	|                           |            | be collected. The syntax for specifying the queries is "namespace:WQL select     |
//	|                           |            | statement". If a specified query cannot be executed, an error is added to        |
//	|                           |            | the output file but collection continues. The format of ManagementQueries is     |
//	|                           |            | specified in [MS-WMI] section 2.                                                 |
//	+---------------------------+------------+----------------------------------------------------------------------------------+
//	| QueryNetworkAdapters      | RW         | Specifies whether network adapter information MUST be queried. If set to TRUE,   |
//	|                           |            | the installed network adapters are enumerated along with their IP addresses and  |
//	|                           |            | offload capabilities. When the client sets this property to VARIANT_TRUE, the    |
//	|                           |            | server SHOULD retrieve the network adapter information and store it locally on   |
//	|                           |            | the server. The PLA Protocol has no knowledge of what information is captured by |
//	|                           |            | the server and written to an XML file, and neither the contents of the XML file  |
//	|                           |            | nor whether the server was successful in writing the XML file can be retrieved   |
//	|                           |            | by the client using the PLA Protocol. Only the VARIANT_BOOL, which indicates     |
//	|                           |            | whether the server MUST query for network adapter information, is transferred    |
//	|                           |            | across the wire. If the client wants to read the network adapter information     |
//	|                           |            | from the server, it needs to use other means or protocols. Whether the server    |
//	|                           |            | queries for network adapter information, and what information it queries, has    |
//	|                           |            | no impact on the behavior of the PLA Protocol. For more information about how    |
//	|                           |            | VARIANT_BOOL types are transferred over the wire, please see [MS-OAUT].<18>      |
//	+---------------------------+------------+----------------------------------------------------------------------------------+
//	| RegistryKeys              | RW         | List of registry keys to be collected. If a specified registry key cannot be     |
//	|                           |            | queried, an error is added to the output file and collection continues. The PLA  |
//	|                           |            | Protocol allows users to log registry keys to understand the configured status   |
//	|                           |            | of a remote system. Registry keys are used to refer to state information that is |
//	|                           |            | stored on the system about an application, driver, or the system. For example,   |
//	|                           |            | what default settings the user has saved for an application might be associated  |
//	|                           |            | with a particular registry key; to retrieve that information, the registry key   |
//	|                           |            | is specified. The format used for the registry keys is specified in [MS-RRP]     |
//	|                           |            | section 3.1.1.1.                                                                 |
//	+---------------------------+------------+----------------------------------------------------------------------------------+
//	| RegistryMaxRecursiveDepth | RW         | Specifies the maximum depth in the registry hierarchy that a recursive registry  |
//	|                           |            | key collection MUST attempt. The maximum depth is relative to the depth of       |
//	|                           |            | the starting key, not absolute. If this value is set to 0, or is not set, then   |
//	|                           |            | registry keys at any depth will be collected. Any unsigned long is a valid value |
//	|                           |            | for this property.                                                               |
//	+---------------------------+------------+----------------------------------------------------------------------------------+
//	| SystemStateFile           | RW         | Specifies the name of the file where the system state will be saved. The system  |
//	|                           |            | state is a set of kernel events generated by taking a snapshot of the Circular   |
//	|                           |            | Kernel Context Logger. The events of the Circular Kernel Context Logger include  |
//	|                           |            | process events, thread events, disk operations, and other kernel information     |
//	|                           |            | that provide an indication of what action the operating system was performing    |
//	|                           |            | when the event was raised. Events for the Circular Kernel Context Logger remain  |
//	|                           |            | in the operating system memory and are only written to file when a snapshot is   |
//	|                           |            | taken of the Circular Kernel Context Logger. This property indicates the name    |
//	|                           |            | of the file to which the contents of the Circular Kernel Context Logger will be  |
//	|                           |            | written; the file will reside on the local system. The file name needs to be a   |
//	|                           |            | file name only and cannot include the path to the file.                          |
//	+---------------------------+------------+----------------------------------------------------------------------------------+
//
// A data collector can be represented as an XML file, which can be used to serialize
// (using Xml (Get) 3.2.4.5.21) and deserialize (using SetXml 3.2.4.5.22)it (the full
// XML specification is available in section 3.2.4.19). The format of the XML that defines
// a configuration data collector is as follows:
//
// Note that the example does not show the property elements inherited from IDataCollector
// that the caller also needs to specify.
//
// Methods in RPC Opnum Order
//
//	+---------------------------------+-------------------------------------------------------------+
//	|                                 |                                                             |
//	|             METHOD              |                         DESCRIPTION                         |
//	|                                 |                                                             |
//	+---------------------------------+-------------------------------------------------------------+
//	+---------------------------------+-------------------------------------------------------------+
//	| FileMaxCount (Get)              | Retrieves the FileMaxCount property. Opnum: 32              |
//	+---------------------------------+-------------------------------------------------------------+
//	| FileMaxCount (Put)              | Sets the FileMaxCount property. Opnum: 33                   |
//	+---------------------------------+-------------------------------------------------------------+
//	| FileMaxRecursiveDepth (Get)     | Retrieves the FileMaxRecursiveDepth property. Opnum: 34     |
//	+---------------------------------+-------------------------------------------------------------+
//	| FileMaxRecursiveDepth (Put)     | Sets the FileMaxRecursiveDepth property. Opnum: 35          |
//	+---------------------------------+-------------------------------------------------------------+
//	| FileMaxTotalSize (Get)          | Retrieves the FileMaxTotalSize property. Opnum: 36          |
//	+---------------------------------+-------------------------------------------------------------+
//	| FileMaxTotalSize (Put)          | Sets the FileMaxTotalSize property. Opnum: 37               |
//	+---------------------------------+-------------------------------------------------------------+
//	| Files (Get)                     | Retrieves the Files property. Opnum: 38                     |
//	+---------------------------------+-------------------------------------------------------------+
//	| Files (Put)                     | Sets the Files property. Opnum: 39                          |
//	+---------------------------------+-------------------------------------------------------------+
//	| ManagementQueries (Get)         | Sets the ManagementQueries property. Opnum: 40              |
//	+---------------------------------+-------------------------------------------------------------+
//	| ManagementQueries (Put)         | Retrieves the ManagementQueries property. Opnum: 41         |
//	+---------------------------------+-------------------------------------------------------------+
//	| QueryNetworkAdapters (Get)      | Retrieves the QueryNetworkAdapters property. Opnum: 42      |
//	+---------------------------------+-------------------------------------------------------------+
//	| QueryNetworkAdapters (Put)      | Sets the QueryNetworkAdapters property. Opnum: 43           |
//	+---------------------------------+-------------------------------------------------------------+
//	| RegistryKeys (Get)              | Retrieves the RegistryKeys property. Opnum: 44              |
//	+---------------------------------+-------------------------------------------------------------+
//	| RegistryKeys (Put)              | Sets the RegistryKeys property. Opnum: 45                   |
//	+---------------------------------+-------------------------------------------------------------+
//	| RegistryMaxRecursiveDepth (Get) | Retrieves the RegistryMaxRecursiveDepth property. Opnum: 46 |
//	+---------------------------------+-------------------------------------------------------------+
//	| RegistryMaxRecursiveDepth (Put) | Sets the RegistryMaxRecursiveDepth property. Opnum: 47      |
//	+---------------------------------+-------------------------------------------------------------+
//	| SystemStateFile (Get)           | Retrieves the SystemStateFile property. Opnum: 48           |
//	+---------------------------------+-------------------------------------------------------------+
//	| SystemStateFile (Put)           | Sets the SystemStateFile property. Opnum: 49                |
//	+---------------------------------+-------------------------------------------------------------+
//
// Opnums 0, 1, and 2 are reserved for the IUnknown interface. Opnums 3, 4, 5, and 6
// are reserved for the IDispatch interface. Opnums 7–31 are used by IDataCollector.
type ConfigurationDataCollector dcom.InterfacePointer

func (o *ConfigurationDataCollector) InterfacePointer() *dcom.InterfacePointer {
	return (*dcom.InterfacePointer)(o)
}

func (o *ConfigurationDataCollector) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if o.Data != nil && o.DataCount == 0 {
		o.DataCount = uint32(len(o.Data))
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}

func (o *ConfigurationDataCollector) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *ConfigurationDataCollector) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *ConfigurationDataCollector) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// Schedule structure represents ISchedule RPC structure.
//
// The ISchedule interface is used to specify when the data collector set runs.
//
// The following properties MUST be implemented by the objects that implement the ISchedule
// interface.
//
//	+-----------+------------+----------------------------------------------------------------------------------+
//	|           |            |                                                                                  |
//	| PROPERTY  | READ/WRITE |                                   DESCRIPTION                                    |
//	|           |            |                                                                                  |
//	+-----------+------------+----------------------------------------------------------------------------------+
//	+-----------+------------+----------------------------------------------------------------------------------+
//	| Days      | RW         | Specifies the days of the week on which the data collector set runs.             |
//	+-----------+------------+----------------------------------------------------------------------------------+
//	| EndDate   | RW         | Specifies the ending date for when the schedule is valid. The value is stored in |
//	|           |            | a VARIANT. Any VARIANT of type Date is a legal value for this type. The format   |
//	|           |            | and over-the-wire transmission of a VARIANT is specified in [MS-OAUT] section    |
//	|           |            | 2.2.29. The time portion of the VARIANT is ignored; only the date portion is     |
//	|           |            | used.                                                                            |
//	+-----------+------------+----------------------------------------------------------------------------------+
//	| StartDate | RW         | Specifies the date that the schedule becomes valid. The value is stored in a     |
//	|           |            | VARIANT. Any VARIANT of type Date is a legal value for this type. The format     |
//	|           |            | and over-the-wire transmission of a VARIANT is specified in [MS-OAUT] section    |
//	|           |            | 2.2.29. The time portion of the VARIANT is ignored; only the date portion is     |
//	|           |            | used.                                                                            |
//	+-----------+------------+----------------------------------------------------------------------------------+
//	| StartTime | RW         | Specifies the time of day when the data collector set starts. The value is       |
//	|           |            | stored in a VARIANT. Any VARIANT of type Date is a legal value for this type.    |
//	|           |            | The format and over-the-wire transmission of a VARIANT is specified in [MS-OAUT] |
//	|           |            | section 2.2.29. The date portion of the VARIANT is ignored; only the time        |
//	|           |            | portion is used.                                                                 |
//	+-----------+------------+----------------------------------------------------------------------------------+
//
// Methods in RPC Opnum Order
//
//	+-----------------+--------------------------------------------+
//	|                 |                                            |
//	|     METHOD      |                DESCRIPTION                 |
//	|                 |                                            |
//	+-----------------+--------------------------------------------+
//	+-----------------+--------------------------------------------+
//	| StartDate (Get) | Retrieves the StartDate property. Opnum: 7 |
//	+-----------------+--------------------------------------------+
//	| StartDate (Put) | Sets the StartDate property. Opnum: 8      |
//	+-----------------+--------------------------------------------+
//	| EndDate (Get)   | Retrieves the EndDate property. Opnum: 9   |
//	+-----------------+--------------------------------------------+
//	| EndDate (Put)   | Sets the EndDate property. Opnum: 10       |
//	+-----------------+--------------------------------------------+
//	| StartTime (Get) | Retrieves the StartTime property Opnum: 11 |
//	+-----------------+--------------------------------------------+
//	| StartTime (Put) | Sets the StartTime property Opnum: 12      |
//	+-----------------+--------------------------------------------+
//	| Days (Get)      | Retrieves the Days property. Opnum: 13     |
//	+-----------------+--------------------------------------------+
//	| Days (Put)      | Sets the Days property. Opnum: 14          |
//	+-----------------+--------------------------------------------+
//
// Opnums 0, 1, and 2 are reserved for the IUnknown interface. Opnums 3, 4, 5, and 6
// are reserved for the IDispatch interface.
type Schedule dcom.InterfacePointer

func (o *Schedule) InterfacePointer() *dcom.InterfacePointer { return (*dcom.InterfacePointer)(o) }

func (o *Schedule) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if o.Data != nil && o.DataCount == 0 {
		o.DataCount = uint32(len(o.Data))
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}

func (o *Schedule) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *Schedule) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *Schedule) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// DataCollector structure represents IDataCollector RPC structure.
//
// The following are the properties that MUST be implemented by the objects that implement
// the IDataCollector interface.
//
//	+-----------------------+------------+----------------------------------------------------------------------------------+
//	|                       |            |                                                                                  |
//	|       PROPERTY        | READ/WRITE |                                   DESCRIPTION                                    |
//	|                       |            |                                                                                  |
//	+-----------------------+------------+----------------------------------------------------------------------------------+
//	+-----------------------+------------+----------------------------------------------------------------------------------+
//	| DataCollectorSet      | R          | The data collector set to which this data collector belongs.                     |
//	+-----------------------+------------+----------------------------------------------------------------------------------+
//	| DataCollectorType     | R          | The type of data collector.                                                      |
//	+-----------------------+------------+----------------------------------------------------------------------------------+
//	| FileName              | RW         | The base name of the file containing the output of the data collector.           |
//	+-----------------------+------------+----------------------------------------------------------------------------------+
//	| FileNameFormat        | RW         | Determines how the name of the file storing the output of the data collector     |
//	|                       |            | will be formatted. The FileName property itself (described above) is always      |
//	|                       |            | included. The filename can optionally be decorated with other information.       |
//	|                       |            | These possible decorations are specified by the AutoPathFormat enumeration. This |
//	|                       |            | property stores an AutoPathFormat value. If the AutoPathFormat value specified   |
//	|                       |            | by this property includes the 'plaPattern' bit, the FileNameFormatPattern        |
//	|                       |            | (specified below) contains the pattern that will be appended to the filename.    |
//	|                       |            | The format of the pattern is specified in section 2.2.3.1.                       |
//	+-----------------------+------------+----------------------------------------------------------------------------------+
//	| FileNameFormatPattern | RW         | If patterns are to be included in the decoration of file names, determines the   |
//	|                       |            | pattern to use. Patterns are included in the decoration if the value of the      |
//	|                       |            | FileNameFormat property (specified above) includes the 'plaPattern' bit. The     |
//	|                       |            | format of the pattern is specified in section 2.2.3.1.                           |
//	+-----------------------+------------+----------------------------------------------------------------------------------+
//	| Index                 | R          | The index value of the data collector.                                           |
//	+-----------------------+------------+----------------------------------------------------------------------------------+
//	| LatestOutputLocation  | RW         | Full path of the file where data was stored the last time the data collector     |
//	|                       |            | ran.                                                                             |
//	+-----------------------+------------+----------------------------------------------------------------------------------+
//	| LogAppend             | RW         | Specifies whether existing files MUST be appended.                               |
//	+-----------------------+------------+----------------------------------------------------------------------------------+
//	| LogCircular           | RW         | Specifies whether files MUST be circular.                                        |
//	+-----------------------+------------+----------------------------------------------------------------------------------+
//	| LogOverwrite          | RW         | Specifies whether existing files MUST be overwritten.                            |
//	+-----------------------+------------+----------------------------------------------------------------------------------+
//	| Name                  | RW         | Name of the data collector.                                                      |
//	+-----------------------+------------+----------------------------------------------------------------------------------+
//	| OutputLocation        | R          | Full path of the file where data would be stored if the data collector were to   |
//	|                       |            | start now.                                                                       |
//	+-----------------------+------------+----------------------------------------------------------------------------------+
//	| Xml                   | R          | The XML representation of the data collector set.                                |
//	+-----------------------+------------+----------------------------------------------------------------------------------+
//
// A data collector can be represented as an XML file, which can be used to serialize
// (using Xml (Get) 3.2.4.5.21) and deserialize (using SetXml 3.2.4.5.22) it. The full
// XML is specified in section 3.2.4.19. The format of the XML that defines a data collector,
// and is common to all types of data collectors, is as follows:
//
// Opnums 8, 28 and 31 are not used across the network. These opnums are reserved and
// MUST NOT be reused by non-Microsoft implementations.<17>
//
// Methods in RPC Opnum Order
//
//	+-----------------------------+----------------------------------------------------------------------------------+
//	|                             |                                                                                  |
//	|           METHOD            |                                   DESCRIPTION                                    |
//	|                             |                                                                                  |
//	+-----------------------------+----------------------------------------------------------------------------------+
//	+-----------------------------+----------------------------------------------------------------------------------+
//	| DataCollectorSet (Get)      | Retrieves the DataCollectorSet property. Opnum: 7                                |
//	+-----------------------------+----------------------------------------------------------------------------------+
//	| Opnum8NotUsedOnWire         | Reserved for local use. Opnum: 8                                                 |
//	+-----------------------------+----------------------------------------------------------------------------------+
//	| DataCollectorType (Get)     | Retrieves the DataCollectorType property. Opnum: 9                               |
//	+-----------------------------+----------------------------------------------------------------------------------+
//	| FileName (Get)              | Retrieves the FileName property. Opnum: 10                                       |
//	+-----------------------------+----------------------------------------------------------------------------------+
//	| FileName (Put)              | Sets the FileName property. Opnum: 11                                            |
//	+-----------------------------+----------------------------------------------------------------------------------+
//	| FileNameFormat (Get)        | Retrieves the FileNameFormat property. Opnum: 12                                 |
//	+-----------------------------+----------------------------------------------------------------------------------+
//	| FileNameFormat (Put)        | Sets the FileNameFormat property. Opnum: 13                                      |
//	+-----------------------------+----------------------------------------------------------------------------------+
//	| FileNameFormatPattern (Get) | Retrieves the FileNameFormatPattern property. Opnum: 14                          |
//	+-----------------------------+----------------------------------------------------------------------------------+
//	| FileNameFormatPattern (Put) | Sets the FileNameFormatPattern property. Opnum: 15                               |
//	+-----------------------------+----------------------------------------------------------------------------------+
//	| LatestOutputLocation (Get)  | Retrieves the LatestOutputLocation property. Opnum: 16                           |
//	+-----------------------------+----------------------------------------------------------------------------------+
//	| LatestOutputLocation (Put)  | Sets the LatestOutputLocation property. Opnum: 17                                |
//	+-----------------------------+----------------------------------------------------------------------------------+
//	| LogAppend (Get)             | Retrieves the LogAppend property. Opnum: 18                                      |
//	+-----------------------------+----------------------------------------------------------------------------------+
//	| LogAppend (Put)             | Sets the LogAppend property. Opnum: 19                                           |
//	+-----------------------------+----------------------------------------------------------------------------------+
//	| LogCircular (Get)           | Retrieves the LogCircular property Opnum: 20                                     |
//	+-----------------------------+----------------------------------------------------------------------------------+
//	| LogCircular (Put)           | Sets the LogCircular property. Opnum: 21                                         |
//	+-----------------------------+----------------------------------------------------------------------------------+
//	| LogOverwrite (Get)          | Retrieves the LogOverwrite property. Opnum: 22                                   |
//	+-----------------------------+----------------------------------------------------------------------------------+
//	| LogOverwrite (Put)          | Sets the LogOverwrite property. Opnum: 23                                        |
//	+-----------------------------+----------------------------------------------------------------------------------+
//	| Name (Get)                  | Retrieves the Name property. Opnum: 24                                           |
//	+-----------------------------+----------------------------------------------------------------------------------+
//	| Name (Put)                  | Sets the Name property. Opnum: 25                                                |
//	+-----------------------------+----------------------------------------------------------------------------------+
//	| OutputLocation (Get)        | Retrieves the OutputLocation property. Opnum: 26                                 |
//	+-----------------------------+----------------------------------------------------------------------------------+
//	| Index (Get)                 | Retrieves the Index property. Opnum: 27                                          |
//	+-----------------------------+----------------------------------------------------------------------------------+
//	| Opnum28NotUsedOnWire        | Reserved for local use. Opnum: 28                                                |
//	+-----------------------------+----------------------------------------------------------------------------------+
//	| Xml (Get)                   | Retrieves the XML property. Opnum: 29                                            |
//	+-----------------------------+----------------------------------------------------------------------------------+
//	| SetXml                      | Sets the properties of the data collector using the values in the XML file.      |
//	|                             | Opnum: 30                                                                        |
//	+-----------------------------+----------------------------------------------------------------------------------+
//	| Opnum31NotUsedOnWire        | Reserved for local use. Opnum: 31                                                |
//	+-----------------------------+----------------------------------------------------------------------------------+
//
// Opnums 0, 1, and 2 are reserved for the IUnknown interface. Opnums 3, 4, 5, and 6
// are reserved for the IDispatch interface.
type DataCollector dcom.InterfacePointer

func (o *DataCollector) InterfacePointer() *dcom.InterfacePointer { return (*dcom.InterfacePointer)(o) }

func (o *DataCollector) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if o.Data != nil && o.DataCount == 0 {
		o.DataCount = uint32(len(o.Data))
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}

func (o *DataCollector) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *DataCollector) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *DataCollector) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// TraceDataProvider structure represents ITraceDataProvider RPC structure.
//
// The ITraceDataProvider interface is used to specify the details on each provider
// that is to be enabled to an event trace session.
//
// The following properties MUST be implemented by objects that implement the ITraceDataProvider
// interface.
//
//	+---------------+------------+----------------------------------------------------------------------------------+
//	|   PROPERTY    |            |                                                                                  |
//	|     NAME      | READ/WRITE |                                   DESCRIPTION                                    |
//	|               |            |                                                                                  |
//	+---------------+------------+----------------------------------------------------------------------------------+
//	+---------------+------------+----------------------------------------------------------------------------------+
//	| DisplayName   | RW         | The name of the provider. The name is provided by the user and can be read by    |
//	|               |            | the user but is otherwise ignored by the PLA protocol. The field exists so that  |
//	|               |            | the user can attach a semantically meaningful name to the ITraceDataProvider     |
//	|               |            | if he or she so chooses rather than having to differentiate providers based on   |
//	|               |            | the Guid property, defined below. The Guid in this context refers to the COM     |
//	|               |            | interface property, which is used as a PLA-UID and does not comply with the      |
//	|               |            | semantics of GUID specified in [C706].                                           |
//	+---------------+------------+----------------------------------------------------------------------------------+
//	| Guid          | RW         | The PLA-UID of the provider. On collection, this PLA-UID uniquely identifies     |
//	|               |            | the provider to be enabled. The value 0000-0000-0000-0000-0000-0000-0000-0000    |
//	|               |            | is never a valid provider PLA-UID, and therefore the Guid property is set to     |
//	|               |            | 0000-0000-0000-0000-0000-0000-0000-0000 when no provider is specified. The       |
//	|               |            | Guid in this context refers to the COM interface property, which is used as a    |
//	|               |            | PLA-UID and does not comply with the semantics of GUID specified in [C706]. The  |
//	|               |            | valid range for this property is from 0000-0000-0000-0000-0000-0000-0000-0001 to |
//	|               |            | FFFF-FFFF-FFFF-FFFF-FFFF-FFFF-FFFF-FFFF.                                         |
//	+---------------+------------+----------------------------------------------------------------------------------+
//	| Level         | R          | The list of levels for the provider. On collection, events from this provider    |
//	|               |            | are collected only if their levels are less than or equal to the value of the    |
//	|               |            | enabled level; this property is used in conjunction with the KeywordsAny and     |
//	|               |            | KeywordsAll properties to control which events are collected from a provider.    |
//	|               |            | The enabled level is stored in the Value property of the Level property. The     |
//	|               |            | level denotes the severity of event (as defined by the event provider). There    |
//	|               |            | are predefined trace levels that can be used to control tracing; there can be    |
//	|               |            | more trace levels that are defined in addition to these predefined levels. The   |
//	|               |            | maximum value of a level is 0x000000FF. The possible pre-defined levels are:     |
//	|               |            | Log Always Critical Error Warning Informational If the enabled level has the     |
//	|               |            | value of 3, all events with level 3, 2, or 1 will be collected. These levels     |
//	|               |            | semantically represent Warning, Error, and Critical events. If the enabled level |
//	|               |            | has the value of 0, then events with any level will be collected. Setting the    |
//	|               |            | enabled level to 0 is equivalent to setting the enabled level to 0x000000FF.     |
//	+---------------+------------+----------------------------------------------------------------------------------+
//	| KeywordsAny   | R          | The list of keywords of the provider. The keywords determine the category of     |
//	|               |            | events for the provider to write; this property is used in conjunction with      |
//	|               |            | the Level and KeywordsAll properties. On collection, events from this provider   |
//	|               |            | are collected only if their keywords include at least one of the keywords in     |
//	|               |            | KeywordsAny and all of the keywords in KeywordsAll. If KeywordsAny is zero,      |
//	|               |            | then the provider will successfully write all events assuming that the event's   |
//	|               |            | keywords pass the KeywordsAll check. The Value property of the KeywordsAny       |
//	|               |            | property stores the bitwise-or of the keywords in the KeywordsAny property.      |
//	+---------------+------------+----------------------------------------------------------------------------------+
//	| KeywordsAll   | R          | The list of keywords of the provider. The keywords determine the category        |
//	|               |            | of events for the provider to write; this is used in conjunction with the        |
//	|               |            | KeywordsAny and Level properties. On collection, events from this provider are   |
//	|               |            | collected only if their keywords include all of the keywords in KeywordsAll.     |
//	|               |            | The Value property of the KeywordsAll property stores the bitwise-or of the      |
//	|               |            | keywords in the KeywordsAll property. If KeywordsAll is zero, then the provider  |
//	|               |            | will successfully write all events assuming that the event's keywords pass the   |
//	|               |            | KeywordsAny check.                                                               |
//	+---------------+------------+----------------------------------------------------------------------------------+
//	| Properties    | R          | The list of extra information that can be collected when events from this        |
//	|               |            | provider are collected. The possible properties are the user's security          |
//	|               |            | identifier, as specified in [MS-DTYP] section 2.4.2.3, (value 1), or the session |
//	|               |            | identifier that is assigned by either the Remote Desktop Session Host server,    |
//	|               |            | Remote Desktop Virtualization Host server, or virtual machine (value 2).         |
//	+---------------+------------+----------------------------------------------------------------------------------+
//	| FilterEnabled | RW         | Determines whether provider-side filtering MUST be enabled. If the FilterEnabled |
//	|               |            | property is set to VARIANT_TRUE, the filter stored in the FilterData property    |
//	|               |            | (as specified in section 3.2.4.11.10) will be used to filter the provider.       |
//	|               |            | Otherwise, the FilterData property will be ignored.                              |
//	+---------------+------------+----------------------------------------------------------------------------------+
//	| FilterType    | RW         | Not used. Because the value is currently not used, any ULONG is a valid value,   |
//	|               |            | so validation will always succeed for the property.                              |
//	+---------------+------------+----------------------------------------------------------------------------------+
//	| FilterData    | RW         | When the client enables tracing for a provider on the server, it has the option  |
//	|               |            | of passing back data to that provider. This data is provider-specific, and the   |
//	|               |            | client MUST know how the provider expects the FilterData to be formatted. The    |
//	|               |            | FilterData property can contain any arbitrary type that is understood by the     |
//	|               |            | trace provider but MUST NOT exceed 1 KB in total size. The PLA protocol has      |
//	|               |            | no knowledge of how this FilterData is constructed and what are its possible     |
//	|               |            | values. This data is opaque to the protocol. It serves as the transport for this |
//	|               |            | data between the client and the server, and the PLA protocol only restriction    |
//	|               |            | on this property is that this data MUST NOT exceed 1 KB in size. The FilterData  |
//	|               |            | that is specified by the client will be sent back to the server trace provider   |
//	|               |            | being enabled if the FilterEnabled property is set to VARIANT_TRUE. Upon         |
//	|               |            | receiving this FilterData, the provider MUST use it to control which events are  |
//	|               |            | logged; this property serves as a filter on the events that are logged by the    |
//	|               |            | provider. The PLA protocol has no knowledge of whether or not the provider did   |
//	|               |            | use the FilterData to control which events it logs. For example, the client can  |
//	|               |            | specify an IP address as the value of the FilterData. When the trace provider    |
//	|               |            | receives this FilterData, it can only log events that have a matching IP         |
//	|               |            | address.                                                                         |
//	+---------------+------------+----------------------------------------------------------------------------------+
//
// Methods in RPC Opnum Order
//
//	+------------------------+----------------------------------------------------------------------------------+
//	|                        |                                                                                  |
//	|         METHOD         |                                   DESCRIPTION                                    |
//	|                        |                                                                                  |
//	+------------------------+----------------------------------------------------------------------------------+
//	+------------------------+----------------------------------------------------------------------------------+
//	| DisplayName (Get)      | Retrieves the DisplayName property. Opnum: 7                                     |
//	+------------------------+----------------------------------------------------------------------------------+
//	| DisplayName (Put)      | Sets the DisplayName property. Opnum: 8                                          |
//	+------------------------+----------------------------------------------------------------------------------+
//	| Guid (Get)             | Retrieves the Guid property. Opnum: 9                                            |
//	+------------------------+----------------------------------------------------------------------------------+
//	| Guid (Put)             | Sets the Guid property. Opnum: 10                                                |
//	+------------------------+----------------------------------------------------------------------------------+
//	| Level (Get)            | Retrieves the Level property. Opnum: 11                                          |
//	+------------------------+----------------------------------------------------------------------------------+
//	| KeywordsAny (Get)      | Retrieves the KeywordsAny property. Opnum: 12                                    |
//	+------------------------+----------------------------------------------------------------------------------+
//	| KeywordsAll (Get)      | Retrieves the KeywordsAll property. Opnum: 13                                    |
//	+------------------------+----------------------------------------------------------------------------------+
//	| Properties (Get)       | Retrieves the Properties property. Opnum: 14                                     |
//	+------------------------+----------------------------------------------------------------------------------+
//	| FilterEnabled (Get)    | Retrieves the FilterEnabled property. Opnum: 15                                  |
//	+------------------------+----------------------------------------------------------------------------------+
//	| FilterEnabled (Put)    | Sets the FilterEnabled property. Opnum: 16                                       |
//	+------------------------+----------------------------------------------------------------------------------+
//	| FilterType (Get)       | Retrieves the FilterType property. Opnum: 17                                     |
//	+------------------------+----------------------------------------------------------------------------------+
//	| FilterType (Put)       | Sets the FilterType property. Opnum: 18                                          |
//	+------------------------+----------------------------------------------------------------------------------+
//	| FilterData (Get)       | Retrieves the FilterData property. Opnum: 19                                     |
//	+------------------------+----------------------------------------------------------------------------------+
//	| FilterData (Put)       | Sets the FilterData property. Opnum: 20                                          |
//	+------------------------+----------------------------------------------------------------------------------+
//	| Query                  | Populates the other properties based on the local repository of providers.       |
//	|                        | Opnum: 21                                                                        |
//	+------------------------+----------------------------------------------------------------------------------+
//	| Resolve                | Used to resolve the properties. Opnum: 22                                        |
//	+------------------------+----------------------------------------------------------------------------------+
//	| SetSecurity            | Updates the system-wide security descriptor of the provider. Opnum: 23           |
//	+------------------------+----------------------------------------------------------------------------------+
//	| GetSecurity            | Retrieves the system-wide security descriptor of the provider. Opnum: 24         |
//	+------------------------+----------------------------------------------------------------------------------+
//	| GetRegisteredProcesses | Retrieves a list of processes that have registered as an event trace data        |
//	|                        | provider. Opnum: 25                                                              |
//	+------------------------+----------------------------------------------------------------------------------+
//
// Opnums 0, 1, and 2 are reserved for the IUnknown interface. Opnums 3, 4, 5, and 6
// are reserved for the IDispatch interface.
type TraceDataProvider dcom.InterfacePointer

func (o *TraceDataProvider) InterfacePointer() *dcom.InterfacePointer {
	return (*dcom.InterfacePointer)(o)
}

func (o *TraceDataProvider) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if o.Data != nil && o.DataCount == 0 {
		o.DataCount = uint32(len(o.Data))
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}

func (o *TraceDataProvider) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *TraceDataProvider) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *TraceDataProvider) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// FolderAction structure represents IFolderAction RPC structure.
//
// The IFolderAction interface is used to specify the actions that the data manager
// is to take on each folder under the data collector set root path if both age and
// size conditions are met.
//
// The following properties MUST be implemented by the objects that implement the IFolderAction
// interface.
//
//	+-----------+------------+----------------------------------------------------------------------------------+
//	|           |            |                                                                                  |
//	| PROPERTY  | READ/WRITE |                                   DESCRIPTION                                    |
//	|           |            |                                                                                  |
//	+-----------+------------+----------------------------------------------------------------------------------+
//	+-----------+------------+----------------------------------------------------------------------------------+
//	| Actions   | RW         | Specifies the actions that the data manager is to take if both conditions (age   |
//	|           |            | and size) are met.                                                               |
//	+-----------+------------+----------------------------------------------------------------------------------+
//	| Age       | RW         | The minimum age of a folder, in days, before it can be considered for these      |
//	|           |            | actions. The age of the folder is the number of days since the folder was        |
//	|           |            | created. If set to zero, no folders will be excluded because of age. The default |
//	|           |            | value is zero. Any unsigned long is a valid value for this property.             |
//	+-----------+------------+----------------------------------------------------------------------------------+
//	| SendCabTo | RW         | Specifies the path for sending the CAB file, if the action includes sending the  |
//	|           |            | CAB file. The path needs to be formatted as a UncPath. If the cab cannot be sent |
//	|           |            | (because the destination does not exist or the DataCollectorSet does not have    |
//	|           |            | write privileges to the destination), the DataManager does not fail and data     |
//	|           |            | management continues.                                                            |
//	+-----------+------------+----------------------------------------------------------------------------------+
//	| Size      | RW         | Specifies the minimum size, in megabytes (MB), of any folder against which the   |
//	|           |            | actions specified in the Actions property will be executed. If set to zero,      |
//	|           |            | no folders will be excluded because of size. The default value is zero. Any      |
//	|           |            | unsigned long is a valid value for this property.                                |
//	+-----------+------------+----------------------------------------------------------------------------------+
//
// Methods in RPC Opnum Order
//
//	+-----------------+---------------------------------------------+
//	|                 |                                             |
//	|     METHOD      |                 DESCRIPTION                 |
//	|                 |                                             |
//	+-----------------+---------------------------------------------+
//	+-----------------+---------------------------------------------+
//	| Age (Get)       | Retrieves the Age property. Opnum: 7        |
//	+-----------------+---------------------------------------------+
//	| Age (Put)       | Sets the Age property. Opnum: 8             |
//	+-----------------+---------------------------------------------+
//	| Size (Get)      | Retrieves the Size property. Opnum: 9       |
//	+-----------------+---------------------------------------------+
//	| Size (Put)      | Sets the Size property. Opnum: 10           |
//	+-----------------+---------------------------------------------+
//	| Actions (Get)   | Retrieves the Actions property. Opnum: 11   |
//	+-----------------+---------------------------------------------+
//	| Actions (Put)   | Sets the Actions property. Opnum: 12        |
//	+-----------------+---------------------------------------------+
//	| SendCabTo (Get) | Retrieves the SendCabTo property. Opnum: 13 |
//	+-----------------+---------------------------------------------+
//	| SendCabTo (Put) | Sets the SendCabTo property. Opnum: 14      |
//	+-----------------+---------------------------------------------+
//
// Opnums 0, 1, and 2 are reserved for the IUnknown interface. Opnums 3, 4, 5, and 6
// are reserved for the IDispatch interface.
type FolderAction dcom.InterfacePointer

func (o *FolderAction) InterfacePointer() *dcom.InterfacePointer { return (*dcom.InterfacePointer)(o) }

func (o *FolderAction) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if o.Data != nil && o.DataCount == 0 {
		o.DataCount = uint32(len(o.Data))
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}

func (o *FolderAction) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *FolderAction) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *FolderAction) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// APITracingDataCollector structure represents IApiTracingDataCollector RPC structure.
//
// The IApiTraceDataCollector interface is used to specify the executables whose API
// calls are to be logged.
//
// The following properties MUST be implemented by the objects that implement the IApiTracing
// DataCollector interface.
//
//	+--------------------+------------+----------------------------------------------------------------------------------+
//	|                    |            |                                                                                  |
//	|      PROPERTY      | READ/WRITE |                                   DESCRIPTION                                    |
//	|                    |            |                                                                                  |
//	+--------------------+------------+----------------------------------------------------------------------------------+
//	+--------------------+------------+----------------------------------------------------------------------------------+
//	| ExcludeApis        | RW         | List of APIs to exclude from the log.                                            |
//	+--------------------+------------+----------------------------------------------------------------------------------+
//	| ExePath            | RW         | Specifies the full path to the executable whose API calls are to be logged.      |
//	+--------------------+------------+----------------------------------------------------------------------------------+
//	| IncludeApis        | RW         | List of APIs to include in the trace. All calls to these APIs that are made in   |
//	|                    |            | the specified executable are logged, even if the modules in which the APIs are   |
//	|                    |            | defined are not included.                                                        |
//	+--------------------+------------+----------------------------------------------------------------------------------+
//	| IncludeModules     | RW         | List of modules to include in the trace. All calls to functions defined in these |
//	|                    |            | modules that are made in the specified executable are logged.                    |
//	+--------------------+------------+----------------------------------------------------------------------------------+
//	| LogApiNamesOnly    | RW         | Specifies whether PLA MUST log only the API name or the arguments and return a   |
//	|                    |            | value as well.                                                                   |
//	+--------------------+------------+----------------------------------------------------------------------------------+
//	| LogApisRecursively | RW         | Specifies whether PLA MUST log only calls that are imported directly by the      |
//	|                    |            | application or all calls to the specified APIs.                                  |
//	+--------------------+------------+----------------------------------------------------------------------------------+
//	| LogFilePath        | RW         | Specifies the name of the file where data MUST be logged to.                     |
//	+--------------------+------------+----------------------------------------------------------------------------------+
//
// A data collector can be represented as an XML file, which can be used to serialize
// (using Xml (Get) 3.2.4.5.21) and deserialize (using SetXml 3.2.4.5.22) it. The format
// of the XML that defines a data collector is as follows (note that the full XML specification
// of the data collector set is in section 3.2.4.19):
//
// This does not show the property elements inherited from IDataCollector that also
// need to be specified.
//
// Methods in RPC Opnum Order
//
//	+--------------------------+------------------------------------------------------+
//	|                          |                                                      |
//	|          METHOD          |                     DESCRIPTION                      |
//	|                          |                                                      |
//	+--------------------------+------------------------------------------------------+
//	+--------------------------+------------------------------------------------------+
//	| LogApiNamesOnly (Get)    | Retrieves the LogApiNamesOnly property. Opnum: 32    |
//	+--------------------------+------------------------------------------------------+
//	| LogApiNamesOnly (Put)    | Sets the LogApiNamesOnly property. Opnum: 33         |
//	+--------------------------+------------------------------------------------------+
//	| LogApisRecursively (Get) | Retrieves the LogApisRecursively property. Opnum: 34 |
//	+--------------------------+------------------------------------------------------+
//	| LogApisRecursively (Put) | Sets the LogApisRecursively property. Opnum: 35      |
//	+--------------------------+------------------------------------------------------+
//	| ExePath (Get)            | Retrieves the ExePath property. Opnum: 36            |
//	+--------------------------+------------------------------------------------------+
//	| ExePath (Put)            | Sets the ExePath property. Opnum: 37                 |
//	+--------------------------+------------------------------------------------------+
//	| LogFilePath (Get)        | Retrieves the LogFilePath property. Opnum: 38        |
//	+--------------------------+------------------------------------------------------+
//	| LogFilePath (Put)        | Sets the LogFilePath property. Opnum: 39             |
//	+--------------------------+------------------------------------------------------+
//	| IncludeModules (Get)     | Retrieves the IncludeModules property. Opnum: 40     |
//	+--------------------------+------------------------------------------------------+
//	| IncludeModules (Put)     | Sets the IncludeModules property. Opnum: 41          |
//	+--------------------------+------------------------------------------------------+
//	| IncludeApis (Get)        | Retrieves the IncludeApis property. Opnum: 42        |
//	+--------------------------+------------------------------------------------------+
//	| IncludeApis (Put)        | Sets the IncludeApis property. Opnum: 43             |
//	+--------------------------+------------------------------------------------------+
//	| ExcludeApis (Get)        | Retrieves the ExcludeApis property. Opnum: 44        |
//	+--------------------------+------------------------------------------------------+
//	| ExcludeApis (Put)        | Sets the ExcludeApis property. Opnum: 45             |
//	+--------------------------+------------------------------------------------------+
//
// Opnums 0, 1, and 2 are reserved for the IUnknown interface. Opnums 3, 4, 5, and 6
// are reserved for the IDispatch interface. Opnums 7–31 are used by IDataCollector.
type APITracingDataCollector dcom.InterfacePointer

func (o *APITracingDataCollector) InterfacePointer() *dcom.InterfacePointer {
	return (*dcom.InterfacePointer)(o)
}

func (o *APITracingDataCollector) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if o.Data != nil && o.DataCount == 0 {
		o.DataCount = uint32(len(o.Data))
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}

func (o *APITracingDataCollector) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *APITracingDataCollector) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *APITracingDataCollector) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// DataCollectorCollection structure represents IDataCollectorCollection RPC structure.
//
// The IDataCollectorCollection interface is used to manage a collection of DataCollector
// objects.
//
// The following properties MUST be implemented by the objects that implement the IDataCollectorCollection
// interface.
//
//	+----------+------------+----------------------------------------------------------------------------------+
//	|          |            |                                                                                  |
//	| PROPERTY | READ/WRITE |                                   DESCRIPTION                                    |
//	|          |            |                                                                                  |
//	+----------+------------+----------------------------------------------------------------------------------+
//	+----------+------------+----------------------------------------------------------------------------------+
//	| _NewEnum | R          | An enumeration object of type IEnumVariant containing a snapshot of the          |
//	|          |            | IDataCollector objects in this collection. The enumeration object is specified   |
//	|          |            | in [MS-OAUT] section 3.3.                                                        |
//	+----------+------------+----------------------------------------------------------------------------------+
//	| Count    | R          | Number of data collectors in this collection.                                    |
//	+----------+------------+----------------------------------------------------------------------------------+
//	| Item     | R          | Retrieves the requested data collector from the collection.                      |
//	+----------+------------+----------------------------------------------------------------------------------+
//
// Methods in RPC Opnum Order
//
//	+----------------------------+---------------------------------------------------------------+
//	|                            |                                                               |
//	|           METHOD           |                          DESCRIPTION                          |
//	|                            |                                                               |
//	+----------------------------+---------------------------------------------------------------+
//	+----------------------------+---------------------------------------------------------------+
//	| Count (Get)                | Retrieves the Count property. Opnum: 7                        |
//	+----------------------------+---------------------------------------------------------------+
//	| Item (Get)                 | Retrieves the Item property. Opnum: 8                         |
//	+----------------------------+---------------------------------------------------------------+
//	| _NewEnum (Get)             | Retrieves the NewEnum property. Opnum: 9                      |
//	+----------------------------+---------------------------------------------------------------+
//	| Add                        | Adds a data collector to the collection. Opnum: 10            |
//	+----------------------------+---------------------------------------------------------------+
//	| Remove                     | Removes a data collector from the collection. Opnum: 11       |
//	+----------------------------+---------------------------------------------------------------+
//	| Clear                      | Removes all data collectors from the collection. Opnum: 12    |
//	+----------------------------+---------------------------------------------------------------+
//	| AddRange                   | Adds one or more data collectors to the collection. Opnum: 13 |
//	+----------------------------+---------------------------------------------------------------+
//	| CreateDataCollectorFromXml | Creates a data collector using XML. Opnum: 14                 |
//	+----------------------------+---------------------------------------------------------------+
//	| CreateDataCollector        | Creates a data collector of the specified type. Opnum: 15     |
//	+----------------------------+---------------------------------------------------------------+
//
// Opnums 0, 1, and 2 are reserved for the IUnknown interface. Opnums 3, 4, 5, and 6
// are reserved for the IDispatch interface.
type DataCollectorCollection dcom.InterfacePointer

func (o *DataCollectorCollection) InterfacePointer() *dcom.InterfacePointer {
	return (*dcom.InterfacePointer)(o)
}

func (o *DataCollectorCollection) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if o.Data != nil && o.DataCount == 0 {
		o.DataCount = uint32(len(o.Data))
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}

func (o *DataCollectorCollection) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *DataCollectorCollection) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *DataCollectorCollection) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// TraceDataCollector structure represents ITraceDataCollector RPC structure.
//
// The ITraceDataCollector interface is used to collect trace events from trace data
// providers.
//
// The following properties MUST be implemented by the objects that implement the ITraceDataCollector
// interface.
//
//	+---------------------+---------------------------------+----------------------------------------------------------------------------------+
//	|                     |                                 |                                                                                  |
//	|      PROPERTY       |           READ/WRITE            |                                   DESCRIPTION                                    |
//	|                     |                                 |                                                                                  |
//	+---------------------+---------------------------------+----------------------------------------------------------------------------------+
//	+---------------------+---------------------------------+----------------------------------------------------------------------------------+
//	| BufferSize          | RW                              | Specifies the suggested buffer size, in kilobytes (KB), for each event tracing   |
//	|                     |                                 | session buffer. The minimum value of the BufferSize property is 1 kilobyte. The  |
//	|                     |                                 | maximum value is 1,024 KB. The default value is 8 KB.                            |
//	+---------------------+---------------------------------+----------------------------------------------------------------------------------+
//	| BuffersLost         | R                               | Specifies the number of buffers that could not be written to the log file. Any   |
//	|                     |                                 | unsigned long is a valid value for this property. This property is only updated  |
//	|                     |                                 | when the IDataCollectorSet::Start or IDataCollectorSet::Query methods are        |
//	|                     |                                 | called.<22>                                                                      |
//	+---------------------+---------------------------------+----------------------------------------------------------------------------------+
//	| BuffersWritten      | R                               | If running, specifies the number of buffers written to the log file. Any         |
//	|                     |                                 | unsigned long is a valid value for this property.                                |
//	+---------------------+---------------------------------+----------------------------------------------------------------------------------+
//	| ClockType           | RW                              | Retrieves or sets the clock resolution to use when logging the time stamp for    |
//	|                     |                                 | each event.                                                                      |
//	+---------------------+---------------------------------+----------------------------------------------------------------------------------+
//	| EventsLost          | R                               | If running, specifies the number of events that were lost due to the             |
//	|                     |                                 | lack of buffers to write to. Any unsigned long is a valid value for this         |
//	|                     |                                 | property. This property is only updated when the IDataCollectorSet::Start or     |
//	|                     |                                 | IDataCollectorSet::Query methods are called.<23>                                 |
//	+---------------------+---------------------------------+----------------------------------------------------------------------------------+
//	| ExtendedModes       | RW                              | Retrieves or sets the log file modes that are not already set by this or other   |
//	|                     |                                 | objects. The valid values for ExtendedModes are specified in section 2.2.9.      |
//	+---------------------+---------------------------------+----------------------------------------------------------------------------------+
//	| FlushTimer          | RW                              | Specifies the time, in seconds, to wait before flushing buffers either to a      |
//	|                     |                                 | log file or by delivering the buffers to an event consuming the application      |
//	|                     |                                 | in real-time. If zero, the buffers are flushed as soon as they are filled. If    |
//	|                     |                                 | nonzero, all buffers containing at least one event are flushed every time the    |
//	|                     |                                 | number of seconds specified in this property elapse. Any unsigned long is a      |
//	|                     |                                 | valid value for this property.                                                   |
//	+---------------------+---------------------------------+----------------------------------------------------------------------------------+
//	| FreeBuffers         | R                               | If running, specifies the number of buffers that are allocated but unused        |
//	|                     |                                 | in the event tracing session's buffer pool. Any unsigned long is a valid         |
//	|                     |                                 | value for the FreeBuffers property. This property is only updated when the       |
//	|                     |                                 | IDataCollectorSet::Start or IDataCollectorSet::Query methods are called.<24>     |
//	+---------------------+---------------------------------+----------------------------------------------------------------------------------+
//	| Guid                | RW                              | Specifies the PLA-UID of the session. If the supplied PLA-UID maps to a          |
//	|                     |                                 | Security Descriptor, the session will run using that Security Descriptor. If no  |
//	|                     |                                 | PLA-UID is supplied, a PLA-UID is generated. The PLA-UID is used for internal    |
//	|                     |                                 | state tracking and does not correspond to any of the DCOM GUID subtypes. The     |
//	|                     |                                 | valid range for this property is from 0000-0000-0000-0000-0000-0000-0000-0001    |
//	|                     |                                 | to FFFF-FFFF-FFFF-FFFF-FFFF-FFFF-FFFF-FFFF. Any PLA-UID other than               |
//	|                     |                                 | 0000-0000-0000-0000-0000-0000-0000-0000 is a valid value for this property. The  |
//	|                     |                                 | Guid in this context refers to the COM interface property which is used as a     |
//	|                     |                                 | PLA-UID and does not comply with the semantics of GUID specified in [C706].      |
//	+---------------------+---------------------------------+----------------------------------------------------------------------------------+
//	| IsKernelTrace       | R                               | Specifies whether this trace data collector includes kernel event trace data     |
//	|                     |                                 | providers.                                                                       |
//	+---------------------+---------------------------------+----------------------------------------------------------------------------------+
//	| MaximumBuffers      | RW                              | Specifies the suggested maximum number of buffers to allocate for the event      |
//	|                     |                                 | tracing session's buffer pool. The value of the MaximumBuffers property has      |
//	|                     |                                 | to be greater than or equal to the value of the MinimumBuffers property. Any     |
//	|                     |                                 | unsigned long is a valid number of maximum buffers to suggest, but there is no   |
//	|                     |                                 | guarantee that the suggestion will be followed. If the MaximumBuffers property   |
//	|                     |                                 | is set to zero, it implies that the user is requesting that no more than         |
//	|                     |                                 | zero buffers be used. As with any other possible value for the MaximumBuffers    |
//	|                     |                                 | property, the suggestion might or might not be followed depending on whether the |
//	|                     |                                 | protocol implementation supports the requested value. No error will occur if the |
//	|                     |                                 | suggestion is not followed, unless MaximumBuffers was set to a lower value than  |
//	|                     |                                 | MinimumBuffers, and it is not possible to detect if the suggestion is followed.  |
//	|                     |                                 | The protocol does not provide any mechanism to discover the actual number of     |
//	|                     |                                 | buffers being used.                                                              |
//	+---------------------+---------------------------------+----------------------------------------------------------------------------------+
//	| MinimumBuffers      | RW                              | Specifies the suggested minimum number of buffers to allocate for the event      |
//	|                     |                                 | tracing session's buffer pool. The value of the MinimumBuffers property has to   |
//	|                     |                                 | be less than or equal to the value of the MaximumBuffers property. Any unsigned  |
//	|                     |                                 | long is a valid number of minimum buffers to suggest, but there is no guarantee  |
//	|                     |                                 | that the suggestion will be followed. If the MinimumBuffers property is set      |
//	|                     |                                 | to zero, it implies that the user is requesting that as few as zero buffers      |
//	|                     |                                 | be used. As with any other possible value for the MinimumBuffers property, the   |
//	|                     |                                 | suggestion might or might not be followed depending on whether the protocol      |
//	|                     |                                 | implementation supports the requested value. No error will occur if the          |
//	|                     |                                 | suggestion is not followed and it is not possible to detect if the suggestion    |
//	|                     |                                 | is followed. The protocol does not provide any mechanism to discover the actual  |
//	|                     |                                 | number of buffers being used.                                                    |
//	+---------------------+---------------------------------+----------------------------------------------------------------------------------+
//	| NumberOfBuffers     | RW                              | Specifies the suggested number of buffers to use for logging. Any unsigned long  |
//	|                     |                                 | is a valid number of buffers to suggest, though there is no guarantee that the   |
//	|                     |                                 | suggestion will be followed. If the NumberOfBuffers property is set to zero, it  |
//	|                     |                                 | implies that the user is requesting that zero buffers be used. As with any other |
//	|                     |                                 | possible value for the NumberOfBuffers property, the suggestion might or might   |
//	|                     |                                 | not be followed depending on whether the protocol implementation supports the    |
//	|                     |                                 | requested value. No error will occur if the suggestion is not followed and it    |
//	|                     |                                 | is not possible to detect if the suggestion is followed. The protocol does not   |
//	|                     |                                 | provide any mechanism to discover the actual number of buffers being used.       |
//	+---------------------+---------------------------------+----------------------------------------------------------------------------------+
//	| PreallocateFile     | RW                              | Specifies whether or not PLA MUST allocate the IDataCollectorSet::SegmentMaxSize |
//	|                     |                                 | on disk prior to the start of the trace data collector. However, if              |
//	|                     |                                 | IDataCollectorSet::SegmentMaxSize is set to zero, then this property is ignored. |
//	+---------------------+---------------------------------+----------------------------------------------------------------------------------+
//	| ProcessMode         | RW                              | Specifies whether or not a process-private logger MUST be used when the          |
//	|                     |                                 | ITraceDataCollector is executing on the server. When events are logged using     |
//	|                     |                                 | ETW, they are temporarily directed to buffers before they are written to a file  |
//	|                     |                                 | or delivered to a real-time monitoring application. If this property is set      |
//	|                     |                                 | to TRUE, then when the ITraceDataCollector begins executing on the server, the   |
//	|                     |                                 | buffers will not be allocated from kernel memory, but from process memory. If    |
//	|                     |                                 | this property is set to FALSE, then the buffers will be allocated from kernel    |
//	|                     |                                 | memory. If this property is set to TRUE, and the ITraceDataProviderCollection is |
//	|                     |                                 | empty or if any one of the specified trace providers is a kernel-mode provider,  |
//	|                     |                                 | then the IDataCollectorSet::Start method will fail and an error code will be     |
//	|                     |                                 | returned. The ITraceDataCollector specifies which providers to enable to this    |
//	|                     |                                 | process-private logger. The buffers will be allocated in the process-space       |
//	|                     |                                 | when the ITraceDataCollector begins executing. However, nothing will be          |
//	|                     |                                 | written to these buffers unless the provider, which is also specified in the     |
//	|                     |                                 | ITraceDataCollector, registers with ETW on the server and begins using the ETW   |
//	|                     |                                 | API to log events. If this property is set to TRUE, and the provider registers   |
//	|                     |                                 | multiple times with ETW from different processes, then there will be different   |
//	|                     |                                 | trace files that are generated, one for each process. The file name is specified |
//	|                     |                                 | in the ITraceDataCollector, and to this file name is appended the extension,     |
//	|                     |                                 | etl, and the process ID. If the process-private logger were running in a process |
//	|                     |                                 | with a process ID of 4, and the file name specified in the ITraceDataCollector   |
//	|                     |                                 | is MyFile, then the file name that the events would be written to would be       |
//	|                     |                                 | MyFile.etl_4. This is not the case if this property is set to FALSE, since the   |
//	|                     |                                 | different provider instances from the different processes will all log to the    |
//	|                     |                                 | buffers that are allocated in kernel memory, and these buffers are associated    |
//	|                     |                                 | with a single trace file. As a result, no process ID is appended after the file  |
//	|                     |                                 | extension etl. If a process-private logger is used, the process in which the     |
//	|                     |                                 | buffers are allocated has the ability to modify the contents of the buffers.     |
//	|                     |                                 | However, because the buffers are in the process, they are only visible to that   |
//	|                     |                                 | process. While a process hosting a private session can edit the contents of a    |
//	|                     |                                 | buffer after an event is written to it, only that process can view those buffers |
//	|                     |                                 | and consequently see those edits. If a process-private logger is not used, the   |
//	|                     |                                 | buffers are allocated in the kernel. In this case, all processes can potentially |
//	|                     |                                 | view the contents of buffers, but no process has the ability to edit them.       |
//	+---------------------+---------------------------------+----------------------------------------------------------------------------------+
//	| RealTimeBuffersLost | R                               | If running, specifies the number of buffers that could not be delivered in real  |
//	|                     |                                 | time to the consumer. RealTimeBuffers are lost when the backup file for storing  |
//	|                     |                                 | events cannot be read or written to by the Event Tracing for infrastructure. In  |
//	|                     |                                 | these situations, the buffers are not recoverable. It is not the case that the   |
//	|                     |                                 | buffers are arriving late; instead, they are not arriving at all. Any unsigned   |
//	|                     |                                 | long is a valid value for RealTimeBuffersLost. This property is only updated     |
//	|                     |                                 | when the IDataCollectorSet::Start or IDataCollectorSet::Query methods are        |
//	|                     |                                 | called.<25>                                                                      |
//	+---------------------+---------------------------------+----------------------------------------------------------------------------------+
//	| SessionId           | R                               | This property refers to the session identifier of the ETW trace session. When    |
//	|                     |                                 | an ITraceDataCollector executes, it starts an ETW trace session; this session    |
//	|                     |                                 | is marked with a numeric identifier. There can be up to 64 different sessions.   |
//	|                     |                                 | Each of these sessions is marked by a different numeric SessionId, which is      |
//	|                     |                                 | specified in this property. Therefore, this property can have the value of       |
//	|                     |                                 | 0x0000000000000000 to 0x000000000000003F. The lower 2 bytes have the session     |
//	|                     |                                 | IDs that are possible (0x0000 to 0x003F) while the upper 6 bytes MUST be         |
//	|                     |                                 | ignored. This property is only updated when the IDataCollectorSet::Start or      |
//	|                     |                                 | IDataCollectorSet::Query methods are called.<26>                                 |
//	+---------------------+---------------------------------+----------------------------------------------------------------------------------+
//	| SessionName         | RW                              | Specifies the name of the session to be created to collect event trace data.     |
//	+---------------------+---------------------------------+----------------------------------------------------------------------------------+
//	| SessionThreadId     | R                               | If running, specifies the ID of the thread performing the logging of the         |
//	|                     |                                 | session. This property is only updated when the IDataCollectorSet::Start or      |
//	|                     |                                 | IDataCollectorSet::Query methods are called.<27>                                 |
//	+---------------------+---------------------------------+----------------------------------------------------------------------------------+
//	| StreamMode          | RW                              | Specifies the logging mode of the trace session.                                 |
//	+---------------------+---------------------------------+----------------------------------------------------------------------------------+
//	| TraceDataProviders  | R (returned object is writable) | List of providers to be enabled for this trace session.                          |
//	+---------------------+---------------------------------+----------------------------------------------------------------------------------+
//
// A data collector can be represented as an XML file, which can be used to serialize
// (using Xml (Get) 3.2.4.5.21) and deserialize (using SetXml 3.2.4.5.22) it. The format
// of the XML that defines a trace data collector is as follows (note that the full
// specification of the data collector set XML is in section 3.2.4.19):
//
// The XML given above does not show the property elements inherited from IDataCollector
// that also need to be specified.
//
// Opnums 35, 37, 41, 47, 62, 64, and 68 are not used across the network. These opnums
// are reserved and MUST NOT be reused by non-Microsoft implementations.
//
// Methods in RPC Opnum Order
//
//	+---------------------------+-------------------------------------------------------+
//	|                           |                                                       |
//	|          METHOD           |                      DESCRIPTION                      |
//	|                           |                                                       |
//	+---------------------------+-------------------------------------------------------+
//	+---------------------------+-------------------------------------------------------+
//	| BufferSize (Get)          | Retrieves the BufferSize property. Opnum: 32          |
//	+---------------------------+-------------------------------------------------------+
//	| BufferSize (Put)          | Sets the BufferSize property. Opnum: 33               |
//	+---------------------------+-------------------------------------------------------+
//	| BuffersLost (Get)         | Retrieves the BufferLost property. Opnum: 34          |
//	+---------------------------+-------------------------------------------------------+
//	| Opnum35NotUsedOnWire      | Reserved for local use. Opnum: 35                     |
//	+---------------------------+-------------------------------------------------------+
//	| BuffersWritten (Get)      | Retrieves the BuffersWritten property. Opnum: 36      |
//	+---------------------------+-------------------------------------------------------+
//	| Opnum37NotUsedOnWire      | Reserved for local use. Opnum: 37                     |
//	+---------------------------+-------------------------------------------------------+
//	| ClockType (Get)           | Retrieves the ClockType property. Opnum: 38           |
//	+---------------------------+-------------------------------------------------------+
//	| ClockType (Put)           | Sets the ClockType property. Opnum: 39                |
//	+---------------------------+-------------------------------------------------------+
//	| EventsLost (Get)          | Retrieves the EventsLost property. Opnum: 40          |
//	+---------------------------+-------------------------------------------------------+
//	| Opnum41NotUsedOnWire      | Reserved for local use. Opnum: 41                     |
//	+---------------------------+-------------------------------------------------------+
//	| ExtendedModes (Get)       | Retrieves the ExtendedModes property. Opnum: 42       |
//	+---------------------------+-------------------------------------------------------+
//	| ExtendedModes (Put)       | Sets the ExtendedModes property. Opnum: 43            |
//	+---------------------------+-------------------------------------------------------+
//	| FlushTimer (Get)          | Retrieves the FlushTimer property. Opnum: 44          |
//	+---------------------------+-------------------------------------------------------+
//	| FlushTimer (Put)          | Sets the FlushTimer property. Opnum: 45               |
//	+---------------------------+-------------------------------------------------------+
//	| FreeBuffers (Get)         | Retrieves the FreeBuffers property. Opnum: 46         |
//	+---------------------------+-------------------------------------------------------+
//	| Opnum47NotUsedOnWire      | Reserved for local use. Opnum: 47                     |
//	+---------------------------+-------------------------------------------------------+
//	| Guid (Get)                | Retrieves the Guid property. Opnum: 48                |
//	+---------------------------+-------------------------------------------------------+
//	| Guid (Put)                | Sets the Guid property. Opnum: 49                     |
//	+---------------------------+-------------------------------------------------------+
//	| IsKernelTrace (Get)       | Retrieves the IsKernelTrace property. Opnum: 50       |
//	+---------------------------+-------------------------------------------------------+
//	| MaximumBuffers (Get)      | Retrieves the MaximumBuffers property. Opnum: 51      |
//	+---------------------------+-------------------------------------------------------+
//	| MaximumBuffers (Put)      | Sets the MaximumBuffers property. Opnum: 52           |
//	+---------------------------+-------------------------------------------------------+
//	| MinimumBuffers (Get)      | Retrieves the MinimumBuffers property. Opnum: 53      |
//	+---------------------------+-------------------------------------------------------+
//	| MinimumBuffers (Put)      | Sets the MinimumBuffers property. Opnum: 54           |
//	+---------------------------+-------------------------------------------------------+
//	| NumberOfBuffers (Get)     | Retrieves the NumberOfBuffers property. Opnum: 55     |
//	+---------------------------+-------------------------------------------------------+
//	| NumberOfBuffers (Put)     | Sets the NumberOfBuffers property. Opnum: 56          |
//	+---------------------------+-------------------------------------------------------+
//	| PreallocateFile (Get)     | Retrieves the PreallocateFile property. Opnum: 57     |
//	+---------------------------+-------------------------------------------------------+
//	| PreallocateFile (Put)     | Sets the PreallocateFile property. Opnum: 58          |
//	+---------------------------+-------------------------------------------------------+
//	| ProcessMode (Get)         | Retrieves the ProcessMode property. Opnum: 59         |
//	+---------------------------+-------------------------------------------------------+
//	| ProcessMode (Put)         | Sets the ProcessMode property. Opnum: 60              |
//	+---------------------------+-------------------------------------------------------+
//	| RealTimeBuffersLost (Get) | Retrieves the RealTimeBuffersLost property. Opnum: 61 |
//	+---------------------------+-------------------------------------------------------+
//	| Opnum62NotUsedOnWire      | Reserved for local use. Opnum: 62                     |
//	+---------------------------+-------------------------------------------------------+
//	| SessionId (Get)           | Retrieves the SessionId property. Opnum: 63           |
//	+---------------------------+-------------------------------------------------------+
//	| Opnum64NotUsedOnWire      | Reserved for local use. Opnum: 64                     |
//	+---------------------------+-------------------------------------------------------+
//	| SessionName (Get)         | Retrieves the SessionName. Opnum: 65                  |
//	+---------------------------+-------------------------------------------------------+
//	| SessionName (Put)         | Sets the SessionName Opnum: 66                        |
//	+---------------------------+-------------------------------------------------------+
//	| SessionThreadId (Get)     | Retrieves the SessionThreadId property. Opnum: 67     |
//	+---------------------------+-------------------------------------------------------+
//	| Opnum68NotUsedOnWire      | Reserved for local use. Opnum: 68                     |
//	+---------------------------+-------------------------------------------------------+
//	| StreamMode (Get)          | Retrieves the StreamMode property. Opnum: 69          |
//	+---------------------------+-------------------------------------------------------+
//	| StreamMode (Put)          | Retrieves the StreamMode property. Opnum: 70          |
//	+---------------------------+-------------------------------------------------------+
//	| TraceDataProviders (Get)  | Retrieves the TraceDataProviders property. Opnum: 71  |
//	+---------------------------+-------------------------------------------------------+
//
// Opnums 0, 1, and 2 are reserved for the IUnknown interface. Opnums 3, 4, 5, and 6
// are reserved for the IDispatch interface. Opnums 7–31 are used by IDataCollector.
type TraceDataCollector dcom.InterfacePointer

func (o *TraceDataCollector) InterfacePointer() *dcom.InterfacePointer {
	return (*dcom.InterfacePointer)(o)
}

func (o *TraceDataCollector) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if o.Data != nil && o.DataCount == 0 {
		o.DataCount = uint32(len(o.Data))
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}

func (o *TraceDataCollector) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *TraceDataCollector) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *TraceDataCollector) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// ValueMap structure represents IValueMap RPC structure.
//
// The IValueMap interface is used to manage a collection of named-value pairs.
//
// Objects that implement this interface represent value maps. The following properties
// MUST be implemented by the objects that implement the IValueMap interface.
//
//	+--------------+------------+----------------------------------------------------------------------------------+
//	|              |            |                                                                                  |
//	|   PROPERTY   | READ/WRITE |                                   DESCRIPTION                                    |
//	|              |            |                                                                                  |
//	+--------------+------------+----------------------------------------------------------------------------------+
//	+--------------+------------+----------------------------------------------------------------------------------+
//	| _NewEnum     | R          | An enumeration object of type IEnumVariant containing a snapshot of the          |
//	|              |            | IValueMapItem objects in this collection. The enumeration object is specified in |
//	|              |            | [MS-OAUT] section 3.3.                                                           |
//	+--------------+------------+----------------------------------------------------------------------------------+
//	| Count        | R          | The number of value map items in the value map.                                  |
//	+--------------+------------+----------------------------------------------------------------------------------+
//	| Description  | RW         | Specifies the description of the value map.                                      |
//	+--------------+------------+----------------------------------------------------------------------------------+
//	| Item         | R          | Retrieves the requested value map item from the value map.                       |
//	+--------------+------------+----------------------------------------------------------------------------------+
//	| Value        | RW         | Specifies the value of the value map. The value is stored in a VARIANT. Any      |
//	|              |            | VARIANT is a legal value for this type. The Value property can be used for       |
//	|              |            | several purposes. Most commonly, the APIs described in this document use the     |
//	|              |            | Value property to indicate which ValueMapItem is considered to be currently      |
//	|              |            | selected. In these cases, the ValueMapItems each represent a possible value for  |
//	|              |            | a ValueMap, and the Value property stores the key to the ValueMapItem that is    |
//	|              |            | currently selected.                                                              |
//	+--------------+------------+----------------------------------------------------------------------------------+
//	| ValueMapType | RW         | Specifies the type of the value map. The possible types of the value map are     |
//	|              |            | specified in section 2.2.2.11.                                                   |
//	+--------------+------------+----------------------------------------------------------------------------------+
//
// Methods in RPC Opnum Order
//
//	+--------------------+-----------------------------------------------------+
//	|                    |                                                     |
//	|       METHOD       |                     DESCRIPTION                     |
//	|                    |                                                     |
//	+--------------------+-----------------------------------------------------+
//	+--------------------+-----------------------------------------------------+
//	| Count (Get)        | Retrieves the Count property. Opnum: 7              |
//	+--------------------+-----------------------------------------------------+
//	| Item (Get)         | Retrieves the Item property. Opnum: 8               |
//	+--------------------+-----------------------------------------------------+
//	| _NewEnum (Get)     | Retrieves the NewEnum property. Opnum: 9            |
//	+--------------------+-----------------------------------------------------+
//	| Description (Get)  | Retrieves the Description property. Opnum: 10       |
//	+--------------------+-----------------------------------------------------+
//	| Description (Put)  | Sets the Description property. Opnum: 11            |
//	+--------------------+-----------------------------------------------------+
//	| Value (Get)        | Retrieves the Value property. Opnum: 12             |
//	+--------------------+-----------------------------------------------------+
//	| Value (Put)        | Sets the Value property. Opnum: 13                  |
//	+--------------------+-----------------------------------------------------+
//	| ValueMapType (Get) | Retrieves the ValueMapType property. Opnum: 14      |
//	+--------------------+-----------------------------------------------------+
//	| ValueMapType (Put) | Sets the ValueMapType property. Opnum: 15           |
//	+--------------------+-----------------------------------------------------+
//	| Add                | Adds an item to the collection. Opnum: 16           |
//	+--------------------+-----------------------------------------------------+
//	| Remove             | Removes an item from the collection. Opnum: 17      |
//	+--------------------+-----------------------------------------------------+
//	| Clear              | Removes all items from the collection Opnum: 18     |
//	+--------------------+-----------------------------------------------------+
//	| AddRange           | Adds one or more items to the collection. Opnum: 19 |
//	+--------------------+-----------------------------------------------------+
//	| CreateValueMapItem | Creates a value map item object. Opnum: 20          |
//	+--------------------+-----------------------------------------------------+
//
// Opnums 0, 1, and 2 are reserved for the IUnknown interface. Opnums 3, 4, 5, and 6
// are reserved for the IDispatch interface.
//
// The IValueMap is used by the Performance Logs and Alerts Protocol for primarily two
// purposes. The first use of the IValueMap is to return properties of the data collector
// set (or its encapsulated objects) that could not be set by the server. Sections 3.2.4.1.54,
// 3.2.4.1.58, 3.2.4.2.24, 3.2.4.5.22, and 3.2.4.15.8 provide more information on how
// the IValueMap is used for this purpose. The second primary use of IValueMap in the
// Performance Logs and Alerts Protocol is to return information about a trace provider
// in the ITraceDataProvider interface. Sections 3.2.4.11.5, 3.2.4.11.6, and 3.2.4.11.7
// detail how information about the trace provider is returned using the IValueMap.
type ValueMap dcom.InterfacePointer

func (o *ValueMap) InterfacePointer() *dcom.InterfacePointer { return (*dcom.InterfacePointer)(o) }

func (o *ValueMap) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if o.Data != nil && o.DataCount == 0 {
		o.DataCount = uint32(len(o.Data))
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}

func (o *ValueMap) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *ValueMap) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *ValueMap) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// DataCollectorSetCollection structure represents IDataCollectorSetCollection RPC structure.
//
// The IDataCollectorSetCollection interface is used to manage a collection of DataCollectorSet
// objects.
//
// The following properties MUST be implemented by the objects that implement the IDataCollectorSetCollection
// interface.
//
//	+----------+------------+----------------------------------------------------------------------------------+
//	|          |            |                                                                                  |
//	| PROPERTY | READ/WRITE |                                   DESCRIPTION                                    |
//	|          |            |                                                                                  |
//	+----------+------------+----------------------------------------------------------------------------------+
//	+----------+------------+----------------------------------------------------------------------------------+
//	| _NewEnum | R          | An enumeration object of type IEnumVariant containing a snapshot of the          |
//	|          |            | IDataCollectorSet objects in this collection. The enumeration object is          |
//	|          |            | specified in [MS-OAUT] section 3.3.                                              |
//	+----------+------------+----------------------------------------------------------------------------------+
//	| Count    | R          | Number of data collector sets in this collection.                                |
//	+----------+------------+----------------------------------------------------------------------------------+
//	| Item     | R          | Retrieves the requested data collector set from the collection.                  |
//	+----------+------------+----------------------------------------------------------------------------------+
//
// Methods in RPC Opnum Order
//
//	+----------------------+----------------------------------------------------------------------------------+
//	|                      |                                                                                  |
//	|        METHOD        |                                   DESCRIPTION                                    |
//	|                      |                                                                                  |
//	+----------------------+----------------------------------------------------------------------------------+
//	+----------------------+----------------------------------------------------------------------------------+
//	| Count (Get)          | Retrieves the Count property. Opnum: 7                                           |
//	+----------------------+----------------------------------------------------------------------------------+
//	| Item (Get)           | Retrieves the Item property. Opnum: 8                                            |
//	+----------------------+----------------------------------------------------------------------------------+
//	| _NewEnum (Get)       | Retrieves the NewEnum property. Opnum: 9                                         |
//	+----------------------+----------------------------------------------------------------------------------+
//	| Add                  | Adds a data collector set to the collection. Opnum: 10                           |
//	+----------------------+----------------------------------------------------------------------------------+
//	| Remove               | Removes a data collector set from the collection. Opnum: 11                      |
//	+----------------------+----------------------------------------------------------------------------------+
//	| Clear                | Removes all data collector sets from the collection. Opnum: 12                   |
//	+----------------------+----------------------------------------------------------------------------------+
//	| AddRange             | Adds one or more data collector set to the collection. Opnum: 13                 |
//	+----------------------+----------------------------------------------------------------------------------+
//	| GetDataCollectorSets | Populates data collector set collection with the persisted data collector sets.  |
//	|                      | Opnum: 14                                                                        |
//	+----------------------+----------------------------------------------------------------------------------+
//
// Opnums 0, 1, and 2 are reserved for the IUnknown interface. Opnums 3, 4, 5, and 6
// are reserved for the IDispatch interface.
type DataCollectorSetCollection dcom.InterfacePointer

func (o *DataCollectorSetCollection) InterfacePointer() *dcom.InterfacePointer {
	return (*dcom.InterfacePointer)(o)
}

func (o *DataCollectorSetCollection) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if o.Data != nil && o.DataCount == 0 {
		o.DataCount = uint32(len(o.Data))
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}

func (o *DataCollectorSetCollection) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *DataCollectorSetCollection) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *DataCollectorSetCollection) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// AlertDataCollector structure represents IAlertDataCollector RPC structure.
//
// The IAlertDataCollector is used to monitor performance counters and to perform actions
// each time a counter value crosses the given threshold.
//
// The following properties MUST be implemented by the objects that implement the IAlertDataCollector
// interface.
//
//	+-------------------------+------------+----------------------------------------------------------------------------------+
//	|                         |            |                                                                                  |
//	|        PROPERTY         | READ/WRITE |                                   DESCRIPTION                                    |
//	|                         |            |                                                                                  |
//	+-------------------------+------------+----------------------------------------------------------------------------------+
//	+-------------------------+------------+----------------------------------------------------------------------------------+
//	| AlertThresholds         | RW         | List of Performance Counters to monitor, along with the threshold values which   |
//	|                         |            | are used to generate alerts. The threshold is specified by appending either      |
//	|                         |            | a '>' or '<' sign along with a value to the Performance Counter path. This       |
//	|                         |            | defines the threshold to be either greater than or less than the provided value, |
//	|                         |            | respectively.                                                                    |
//	+-------------------------+------------+----------------------------------------------------------------------------------+
//	| EventLog                | RW         | Specifies whether or not an event MUST be written to the Eventlog each time the  |
//	|                         |            | value of any counter specified in the AlertThresholds property. If set to true   |
//	|                         |            | and the threshold is crossed, then the event will be logged; otherwise, if set   |
//	|                         |            | to false, and even if the threshold is crossed, the event is not logged. <20>    |
//	+-------------------------+------------+----------------------------------------------------------------------------------+
//	| SampleInterval          | RW         | The time, in seconds, between two consecutive samples. The default value is 15   |
//	|                         |            | seconds. The minimum sample interval is 1 second, and there is no maximum sample |
//	|                         |            | interval. However, if the sample interval is set to 0xFFFFFFFF, only one sample  |
//	|                         |            | will ever be collected.                                                          |
//	+-------------------------+------------+----------------------------------------------------------------------------------+
//	| Task                    | RW         | Name of the Task Scheduler job to be executed each time a Performance Counter    |
//	|                         |            | value crosses the specified threshold.                                           |
//	+-------------------------+------------+----------------------------------------------------------------------------------+
//	| TaskArguments           | RW         | If a task is to run, this specifies the arguments that are passed to it. The     |
//	|                         |            | arguments need to be formatted as command-line arguments. See section 3.2.4.8.11 |
//	|                         |            | for a table of the command line arguments.                                       |
//	+-------------------------+------------+----------------------------------------------------------------------------------+
//	| TaskRunAsSelf           | RW         | When a Task Scheduler job is executed by this AlertDataCollector, this property  |
//	|                         |            | determines which user it runs as. If the property is set to true, the Task       |
//	|                         |            | Scheduler job runs with the same user credentials as the DataCollectorSet.       |
//	|                         |            | By default, this means the Task Scheduler job runs with System credentials.      |
//	|                         |            | Consequently, it is inadvisable to set this property to true when the task       |
//	|                         |            | to be run is not fully trusted unless the UserAccount property for the           |
//	|                         |            | DataCollectorSet has been carefully configured. When the property is set to      |
//	|                         |            | false, the Task Scheduler job runs with the credentials it was created with. The |
//	|                         |            | mechanism in use here is delegation. When the creator of a data collector set    |
//	|                         |            | sets this property to true, he or she is granting this task the same rights that |
//	|                         |            | the data collector set is running with. When the RunAsSelf property is set to    |
//	|                         |            | false, no delegation occurs. The task will run only with the permissions it was  |
//	|                         |            | created with. The credentials that the task runs with are initially created with |
//	|                         |            | SchRpcRegisterTask specified in [MS-TSCH] section 3.2.5.4.2 and can be updated   |
//	|                         |            | by SASetAccountInformation specified in [MS-TSCH] section 3.2.5.3.4.             |
//	+-------------------------+------------+----------------------------------------------------------------------------------+
//	| TaskUserTextArguments   | RW         | If a task is to run and the arguments include the {usertext} variable, this      |
//	|                         |            | property determines the value of this variable. Any BSTR is potentially a valid  |
//	|                         |            | value for this property. For example a random string such as "ch&(26D@!k" is     |
//	|                         |            | a valid value, as are the strings which would normally reference other task      |
//	|                         |            | arguments, such as "{name}". In no case will only substring contained in the     |
//	|                         |            | TaskUserTextArguments be expanded (so if the string includes {name}, the string  |
//	|                         |            | will be passed to the Task with the {name}, not with the value of {name}. The    |
//	|                         |            | put method for this property will never fail. The actual semantic validity of    |
//	|                         |            | any particular BSTR depends on the task specified by the Task property.          |
//	+-------------------------+------------+----------------------------------------------------------------------------------+
//	| TriggerDataCollectorSet | RW         | Name of the data collector set to be started each time a counter value crosses   |
//	|                         |            | the threshold.                                                                   |
//	+-------------------------+------------+----------------------------------------------------------------------------------+
//
// A data collector can be represented as an XML file, which can be used to serialize
// (using Xml (Get) 3.2.4.5.21) and deserialize (using SetXml 3.2.4.5.22) it. The format
// of the XML that defines an alert data collector is as follows (the full XML specification
// of the data collector set is in section 3.2.4.19):
//
//	The format of <Alert> is [performance counter] [<|>] [number]. For example: \Processor(_Total)\% Processor Time>1. See [MSDN-COUNT] for the performance counter path representation.
//
// The example does not show the property elements inherited from IDataCollector that
// also need to be specified.
//
// Methods in RPC Opnum Order
//
//	+-------------------------------+-----------------------------------------------------------+
//	|                               |                                                           |
//	|            METHOD             |                        DESCRIPTION                        |
//	|                               |                                                           |
//	+-------------------------------+-----------------------------------------------------------+
//	+-------------------------------+-----------------------------------------------------------+
//	| AlertThresholds (Get)         | Retrieves the AlertThresholds property. Opnum: 32         |
//	+-------------------------------+-----------------------------------------------------------+
//	| AlertThresholds (Put)         | Sets the AlertThresholds property. Opnum: 33              |
//	+-------------------------------+-----------------------------------------------------------+
//	| EventLog (Get)                | Retrieves the EventLog property. Opnum: 34                |
//	+-------------------------------+-----------------------------------------------------------+
//	| EventLog (Put)                | Sets the EventLog property. Opnum: 35                     |
//	+-------------------------------+-----------------------------------------------------------+
//	| SampleInterval (Get)          | Retrieves the SampleInterval property. Opnum: 36          |
//	+-------------------------------+-----------------------------------------------------------+
//	| SampleInterval (Put)          | Sets the SampleInterval property. Opnum: 37               |
//	+-------------------------------+-----------------------------------------------------------+
//	| Task (Get)                    | Retrieves the Task property. Opnum: 38                    |
//	+-------------------------------+-----------------------------------------------------------+
//	| Task (Put)                    | Sets the Task property. Opnum: 39                         |
//	+-------------------------------+-----------------------------------------------------------+
//	| TaskRunAsSelf (Get)           | Retrieves the TaskRunAsSelf property. Opnum: 40           |
//	+-------------------------------+-----------------------------------------------------------+
//	| TaskRunAsSelf (Put)           | Sets the TaskRunAsSelf property. Opnum: 41                |
//	+-------------------------------+-----------------------------------------------------------+
//	| TaskArguments (Get)           | Retrieves the TaskArguments property. Opnum: 42           |
//	+-------------------------------+-----------------------------------------------------------+
//	| TaskArguments (Put)           | Sets the TaskArguments property. Opnum: 43                |
//	+-------------------------------+-----------------------------------------------------------+
//	| TaskUserTextArguments (Get)   | Retrieves the TaskUserTextArguments property. Opnum: 44   |
//	+-------------------------------+-----------------------------------------------------------+
//	| TaskUserTextArguments (Put)   | Retrieves the TaskUserTextArguments property. Opnum: 45   |
//	+-------------------------------+-----------------------------------------------------------+
//	| TriggerDataCollectorSet (Get) | Retrieves the TriggerDataCollectorSet property. Opnum: 46 |
//	+-------------------------------+-----------------------------------------------------------+
//	| TriggerDataCollectorSet (Put) | Sets the TriggerDataCollectorSet property. Opnum: 47      |
//	+-------------------------------+-----------------------------------------------------------+
//
// Opnums 0, 1, and 2 are reserved for the IUnknown interface. Opnums 3, 4, 5, and 6
// are reserved for the IDispatch interface. Opnums 7–31 are used by IDataCollector.
type AlertDataCollector dcom.InterfacePointer

func (o *AlertDataCollector) InterfacePointer() *dcom.InterfacePointer {
	return (*dcom.InterfacePointer)(o)
}

func (o *AlertDataCollector) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if o.Data != nil && o.DataCount == 0 {
		o.DataCount = uint32(len(o.Data))
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}

func (o *AlertDataCollector) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *AlertDataCollector) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *AlertDataCollector) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// DataManager structure represents IDataManager RPC structure.
//
// The IDataManager interface is used to manage data generated by the data collectors,
// including report generation, data retention policy, and data transfer.
//
// The following properties MUST be implemented by the objects that implement the IDataManager
// interface.
//
//	+--------------------+------------+----------------------------------------------------------------------------------+
//	|                    |            |                                                                                  |
//	|      PROPERTY      | READ/WRITE |                                   DESCRIPTION                                    |
//	|                    |            |                                                                                  |
//	+--------------------+------------+----------------------------------------------------------------------------------+
//	+--------------------+------------+----------------------------------------------------------------------------------+
//	| CheckBeforeRunning | RW         | Indicates whether maximum folder count and minimum free disk space thresholds    |
//	|                    |            | MUST be checked before running the data collector set. If set to VARIANT_TRUE    |
//	|                    |            | and either one of the conditions is not met, the data collector set MUST         |
//	|                    |            | fail the start. If set to VARIANT_FALSE, the thresholds will still be used       |
//	|                    |            | after collection to determine if the resource policy will be invoked. The        |
//	|                    |            | ResourcePolicy property is defined below. The MaxSize threshold is never checked |
//	|                    |            | before running an IDataManager, irrespective of the value of this property.      |
//	+--------------------+------------+----------------------------------------------------------------------------------+
//	| Enabled            | RW         | Determines whether data management is enabled. If set to VARIANT_FALSE, all      |
//	|                    |            | other settings of this object are ignored and no data management operations      |
//	|                    |            | (such as creating a report) execute; the data manager is not run. If set         |
//	|                    |            | to VARIANT_TRUE, all other properties of this object are used and the data       |
//	|                    |            | manager is run. The properties that are used when the data manager runs are as   |
//	|                    |            | follows: CheckBeforeRunning EventsFileName FolderAction MaxFolderCount MaxSize   |
//	|                    |            | MinFreeDisk ReportFileName ReportSchema ResourcePolicy Rules RulesTargetFileName |
//	+--------------------+------------+----------------------------------------------------------------------------------+
//	| EventsFileName     | RW         | Specifies the name of the file that MUST be created by the Performance Logs and  |
//	|                    |            | Alerts Protocol, during report generation. This file contains all the events     |
//	|                    |            | that were collected, serialized to XML. This file differs from the report file   |
//	|                    |            | because the events are not normalized, performance data is not included, and     |
//	|                    |            | rules are not run against this file.                                             |
//	+--------------------+------------+----------------------------------------------------------------------------------+
//	| FolderActions      | R          | List of actions to be performed on the subfolders that match the criteria. The   |
//	|                    |            | IFolderAction interface is specified in section 3.2.4.3.                         |
//	+--------------------+------------+----------------------------------------------------------------------------------+
//	| MaxFolderCount     | RW         | Specifies the threshold for the maximum number of subfolders under the data      |
//	|                    |            | collector sets root path. If this threshold is violated, the action specified by |
//	|                    |            | the ResourcePolicy property is invoked. If this property is set to zero, there   |
//	|                    |            | is no maximum. The default value is zero. Any unsigned long is a valid value for |
//	|                    |            | this property.                                                                   |
//	+--------------------+------------+----------------------------------------------------------------------------------+
//	| MaxSize            | RW         | Specifies the maximum size, in megabytes, of all files under the data collector  |
//	|                    |            | set root path. If this threshold is violated, the action specified by the        |
//	|                    |            | ResourcePolicy property is invoked. If this property is set to zero, there is no |
//	|                    |            | maximum. The default value is zero. Any unsigned long is a valid value for this  |
//	|                    |            | property.                                                                        |
//	+--------------------+------------+----------------------------------------------------------------------------------+
//	| MinFreeDisk        | RW         | Specifies the minimum free space, in MB, that MUST remain available in the       |
//	|                    |            | data collector set root path volume. If this threshold is violated, the data     |
//	|                    |            | collector set will not start collecting data. If this property is set to zero,   |
//	|                    |            | there is no minimum. The default value is zero. Any unsigned long is a valid     |
//	|                    |            | value for this property.                                                         |
//	+--------------------+------------+----------------------------------------------------------------------------------+
//	| ReportFileName     | RW         | Specifies the name of the HTML file that results from converting the file in     |
//	|                    |            | RuleTargetFileName from XML to HTML.                                             |
//	+--------------------+------------+----------------------------------------------------------------------------------+
//	| ReportSchema       | RW         | Specifies the XML used to customize the report generated from the data. The XSD  |
//	|                    |            | that defines the format of the XML is specified in section 2.2.3.3 .             |
//	+--------------------+------------+----------------------------------------------------------------------------------+
//	| ResourcePolicy     | RW         | Specifies the action to be performed if one of the disk resource limits          |
//	|                    |            | is violated. The limits are the maximum folder count, the maximum size           |
//	|                    |            | and the minimum free disk space. The possible actions are described in the       |
//	|                    |            | ResourcePolicy enumeration in section 2.2.2.9.                                   |
//	+--------------------+------------+----------------------------------------------------------------------------------+
//	| Rules              | RW         | The rules to be applied to the report file. The rules are specified in an XML    |
//	|                    |            | format, which is specified in section 2.2.3.4.                                   |
//	+--------------------+------------+----------------------------------------------------------------------------------+
//	| RuleTargetFileName | RW         | Specifies the name of the file where the report data is stored before rules are  |
//	|                    |            | run against it and it is converted to HTML.                                      |
//	+--------------------+------------+----------------------------------------------------------------------------------+
//
// Methods in RPC Opnum Order
//
//	+--------------------------+------------------------------------------------------+
//	|                          |                                                      |
//	|          METHOD          |                     DESCRIPTION                      |
//	|                          |                                                      |
//	+--------------------------+------------------------------------------------------+
//	+--------------------------+------------------------------------------------------+
//	| Enabled (Get)            | Retrieves the Enabled property. Opnum: 7             |
//	+--------------------------+------------------------------------------------------+
//	| Enabled (Put)            | Sets the Enabled property. Opnum: 8                  |
//	+--------------------------+------------------------------------------------------+
//	| CheckBeforeRunning (Get) | Retrieves the CheckBeforeRunning property. Opnum: 9  |
//	+--------------------------+------------------------------------------------------+
//	| CheckBeforeRunning (Put) | Sets the CheckBeforeRunning property. Opnum: 10      |
//	+--------------------------+------------------------------------------------------+
//	| MinFreeDisk (Get)        | Retrieves the MinFreeDisk property. Opnum: 11        |
//	+--------------------------+------------------------------------------------------+
//	| MinFreeDisk (Put)        | Sets the MinFreeDisk property. Opnum: 12             |
//	+--------------------------+------------------------------------------------------+
//	| MaxSize (Get)            | Retrieves the MaxSize property. Opnum: 13            |
//	+--------------------------+------------------------------------------------------+
//	| MaxSize (Put)            | Sets the MaxSize property. Opnum: 14                 |
//	+--------------------------+------------------------------------------------------+
//	| MaxFolderCount (Get)     | Retrieves the MaxFolderCount property. Opnum: 15     |
//	+--------------------------+------------------------------------------------------+
//	| MaxFolderCount (Put)     | Sets the MaxFolderCount property. Opnum: 16          |
//	+--------------------------+------------------------------------------------------+
//	| ResourcePolicy (Get)     | Retrieves the ResourcePolicy property. Opnum: 17     |
//	+--------------------------+------------------------------------------------------+
//	| ResourcePolicy (Put)     | Sets the ResourcePolicy property. Opnum: 18          |
//	+--------------------------+------------------------------------------------------+
//	| FolderActions (Get)      | Retrieves the FolderAction property. Opnum: 19       |
//	+--------------------------+------------------------------------------------------+
//	| ReportSchema (Get)       | Retrieves the ReportSchema property. Opnum: 20       |
//	+--------------------------+------------------------------------------------------+
//	| ReportSchema (Put)       | Sets the ReportSchema property. Opnum: 21            |
//	+--------------------------+------------------------------------------------------+
//	| ReportFileName (Get)     | Retrieves the ReportFileName property. Opnum: 22     |
//	+--------------------------+------------------------------------------------------+
//	| ReportFileName (Put)     | Sets the ReportFileName property. Opnum: 23          |
//	+--------------------------+------------------------------------------------------+
//	| RuleTargetFileName (Get) | Retrieves the RuleTargetFileName property. Opnum: 24 |
//	+--------------------------+------------------------------------------------------+
//	| RuleTargetFileName (Put) | Sets the RuleTargetFileName property. Opnum: 25      |
//	+--------------------------+------------------------------------------------------+
//	| EventsFileName (Get)     | Retrieves the EventsFileName property. Opnum: 26     |
//	+--------------------------+------------------------------------------------------+
//	| EventsFileName (Put)     | Sets the EventsFileName property. Opnum: 27          |
//	+--------------------------+------------------------------------------------------+
//	| Rules (Get)              | Retrieves the Rules property. Opnum: 28              |
//	+--------------------------+------------------------------------------------------+
//	| Rules (Put)              | Sets the Rules property. Opnum: 29                   |
//	+--------------------------+------------------------------------------------------+
//	| Run                      | Manually runs the data manager. Opnum: 30            |
//	+--------------------------+------------------------------------------------------+
//	| Extract                  | Extracts the specified CAB file. Opnum: 31           |
//	+--------------------------+------------------------------------------------------+
//
// Opnums 0, 1, and 2 are reserved for the IUnknown interface. Opnums 3, 4, 5, and 6
// are reserved for the IDispatch interface.
type DataManager dcom.InterfacePointer

func (o *DataManager) InterfacePointer() *dcom.InterfacePointer { return (*dcom.InterfacePointer)(o) }

func (o *DataManager) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if o.Data != nil && o.DataCount == 0 {
		o.DataCount = uint32(len(o.Data))
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}

func (o *DataManager) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *DataManager) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *DataManager) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// ScheduleCollection structure represents IScheduleCollection RPC structure.
//
// # The IScheduleCollection interface is used to manage a collection of Schedule objects
//
// The following properties MUST be implemented by the objects that implement the IScheduleCollection
// interface.
//
//	+----------+------------+----------------------------------------------------------------------------------+
//	|          |            |                                                                                  |
//	| PROPERTY | READ/WRITE |                                   DESCRIPTION                                    |
//	|          |            |                                                                                  |
//	+----------+------------+----------------------------------------------------------------------------------+
//	+----------+------------+----------------------------------------------------------------------------------+
//	| _NewEnum | R          | An enumeration object of type IEnumVariant containing a snapshot of the          |
//	|          |            | ISchedule objects in this collection. The enumeration object is specified in     |
//	|          |            | [MS-OAUT] section 3.3.                                                           |
//	+----------+------------+----------------------------------------------------------------------------------+
//	| Count    | R          | Number of schedules in this collection.                                          |
//	+----------+------------+----------------------------------------------------------------------------------+
//	| Item     | R          | Retrieves the requested schedule from the collection.                            |
//	+----------+------------+----------------------------------------------------------------------------------+
//
// Methods in RPC Opnum Order
//
//	+----------------+---------------------------------------------------------+
//	|                |                                                         |
//	|     METHOD     |                       DESCRIPTION                       |
//	|                |                                                         |
//	+----------------+---------------------------------------------------------+
//	+----------------+---------------------------------------------------------+
//	| Count (Get)    | Retrieves the Count property. Opnum: 7                  |
//	+----------------+---------------------------------------------------------+
//	| Item (Get)     | Retrieves the Item property. Opnum: 8                   |
//	+----------------+---------------------------------------------------------+
//	| _NewEnum (Get) | Retrieves the NewEnum property. Opnum: 9                |
//	+----------------+---------------------------------------------------------+
//	| Add            | Adds a schedule to the collection. Opnum: 10            |
//	+----------------+---------------------------------------------------------+
//	| Remove         | Removes a schedule from the collection. Opnum: 11       |
//	+----------------+---------------------------------------------------------+
//	| Clear          | Removes all schedules from the collection. Opnum: 12    |
//	+----------------+---------------------------------------------------------+
//	| AddRange       | Adds one or more schedules to the collection. Opnum: 13 |
//	+----------------+---------------------------------------------------------+
//	| CreateSchedule | Creates a schedule object. Opnum: 14                    |
//	+----------------+---------------------------------------------------------+
//
// Opnums 0, 1, and 2 are reserved for the IUnknown interface. Opnums 3, 4, 5, and 6
// are reserved for the IDispatch interface.
type ScheduleCollection dcom.InterfacePointer

func (o *ScheduleCollection) InterfacePointer() *dcom.InterfacePointer {
	return (*dcom.InterfacePointer)(o)
}

func (o *ScheduleCollection) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if o.Data != nil && o.DataCount == 0 {
		o.DataCount = uint32(len(o.Data))
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}

func (o *ScheduleCollection) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *ScheduleCollection) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *ScheduleCollection) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// DataCollectorSet structure represents IDataCollectorSet RPC structure.
//
// The IDataCollectorSet interface is used to query, update, and control a data collector
// set.
//
// Objects that implement this interface represent a data collector set. The following
// are the properties that MUST be implemented by objects that implement the IDataCollectorSet
// interface.
//
//	+---------------------------+------------+----------------------------------------------------------------------------------+
//	|                           |            |                                                                                  |
//	|         PROPERTY          | READ/WRITE |                                   DESCRIPTION                                    |
//	|                           |            |                                                                                  |
//	+---------------------------+------------+----------------------------------------------------------------------------------+
//	+---------------------------+------------+----------------------------------------------------------------------------------+
//	| DataCollectors            | R          | List of data collectors in this set.                                             |
//	+---------------------------+------------+----------------------------------------------------------------------------------+
//	| DataManager               | R          | Object that defines the policies for data retention and report generation.       |
//	|                           |            | An example of data retention policies include when and if to compress data to    |
//	|                           |            | a cabinet file or delete it. Example report generation policies include the      |
//	|                           |            | steps to take to generate the report. The IDataManager interface, which the      |
//	|                           |            | DataManager object implements, is specified in section 3.2.4.2.                  |
//	+---------------------------+------------+----------------------------------------------------------------------------------+
//	| Description               | RW         | The localizable description of the data collector set.                           |
//	+---------------------------+------------+----------------------------------------------------------------------------------+
//	| DescriptionUnresolved     | R          | The description of the data collector set in its raw form (prior to              |
//	|                           |            | localization).                                                                   |
//	+---------------------------+------------+----------------------------------------------------------------------------------+
//	| DisplayName               | RW         | The localizable display name of the data collector set.                          |
//	+---------------------------+------------+----------------------------------------------------------------------------------+
//	| DisplayNameUnresolved     | R          | The display name of the data collector set in its raw form (prior to             |
//	|                           |            | localization).                                                                   |
//	+---------------------------+------------+----------------------------------------------------------------------------------+
//	| Duration                  | RW         | Determines the maximum amount of time that a data collector set MUST run.        |
//	|                           |            | A value of 0 indicates that the data collector set MUST NOT not be stopped       |
//	|                           |            | automatically.                                                                   |
//	+---------------------------+------------+----------------------------------------------------------------------------------+
//	| Keywords                  | RW         | List of keywords associated with the data collector sets. Keywords are metadata  |
//	|                           |            | for describing a data collector set; they are not parsed by the data collector   |
//	|                           |            | set. They are intended to help the end user understand the contents of the data  |
//	|                           |            | collector set, and serve no functional purpose as to how the data collector      |
//	|                           |            | set is executed on the server. There MUST be at most 256 keywords that are       |
//	|                           |            | associated with a data collector set; each keyword MUST NOT be more than 1024    |
//	|                           |            | characters in length. The keyword string cannot be the empty string, nor can the |
//	|                           |            | keyword string contain the semicolon (";") character.                            |
//	+---------------------------+------------+----------------------------------------------------------------------------------+
//	| LatestOutputLocation      | RW         | The full path of the directory where data was stored the last time the data      |
//	|                           |            | collector set ran.                                                               |
//	+---------------------------+------------+----------------------------------------------------------------------------------+
//	| Name                      | R          | Name of the data collector set.                                                  |
//	+---------------------------+------------+----------------------------------------------------------------------------------+
//	| OutputLocation            | R          | The full path of the directory where data would be stored if the data collector  |
//	|                           |            | set were to start now.                                                           |
//	+---------------------------+------------+----------------------------------------------------------------------------------+
//	| RootPath                  | RW         | The full path of the directory under which the data collector set stores its     |
//	|                           |            | files. When subdirectories are used, they are created under this root directory. |
//	+---------------------------+------------+----------------------------------------------------------------------------------+
//	| Schedules                 | R          | List of schedules that determines when the data collector set runs               |
//	|                           |            | automatically. Each schedule specifies a time when the data collector will be    |
//	|                           |            | started, the first day it will be started at that time, the last day it will be  |
//	|                           |            | started at that time, and the days of the week it will be started. Each schedule |
//	|                           |            | is specified by an object implementing the ISchedule interface, specified in     |
//	|                           |            | section 3.2.4.12.                                                                |
//	+---------------------------+------------+----------------------------------------------------------------------------------+
//	| SchedulesEnabled          | RW         | Determines if the automatic start of the data collector set based on its         |
//	|                           |            | schedules MUST be enabled or disabled. If enabled, the data collector set is     |
//	|                           |            | automatically started when the conditions for one of the schedules (stored in    |
//	|                           |            | the Schedules property) is met. If the data collector set is already running     |
//	|                           |            | when a schedule condition is met, it is not restarted, and instead continues to  |
//	|                           |            | run. If disabled, the data collector set can only be started manually. A data    |
//	|                           |            | collector set is manually started by a call to Start, as specified in section    |
//	|                           |            | 3.2.4.1.56.                                                                      |
//	+---------------------------+------------+----------------------------------------------------------------------------------+
//	| Security                  | RW         | The security descriptor of the data collector set that determines the access     |
//	|                           |            | rights of groups or users. The security descriptor is expressed using the        |
//	|                           |            | Security Descriptor Description Language (SDDL), as specified in [MS-DTYP]       |
//	|                           |            | section 2.5.1. Changing the security descriptor can impact the ability of both   |
//	|                           |            | local and remote users to read, modify, run, and delete the data collector set.  |
//	+---------------------------+------------+----------------------------------------------------------------------------------+
//	| Segment                   | RW         | If this property is set to TRUE, the server MUST stop all data collectors        |
//	|                           |            | when a segmentation condition occurs, increment the SerialNumber                 |
//	|                           |            | property, update LatestOutputLocation property, and restart all the data         |
//	|                           |            | collectors, which begins logging to the new LatestOutputLocation. The            |
//	|                           |            | FileNameFormat and SubdirectoryNameFormat properties are used to determine       |
//	|                           |            | the new value of LatestOutputLocation. If both FileNameFormatPattern and         |
//	|                           |            | SubdirectoryFormatPattern are set to plaNone, or both are set to plaPattern but  |
//	|                           |            | no value is specified for FileNameFormatPattern and SubdirectoryFormatPattern,   |
//	|                           |            | then the value of the LatestOutputLocation is not changed. In this case, the     |
//	|                           |            | retention of the data that was obtained during the previous execution of the     |
//	|                           |            | data collector depends on the respective properties of the data collector.       |
//	|                           |            | For example, if LogAppend is specified for a data collector, then the events     |
//	|                           |            | generated in the new segment are appended to the events that were added to the   |
//	|                           |            | file during the previous segment execution. If this property is set to FALSE,    |
//	|                           |            | the server MUST stop all the data collectors when a segmentation condition       |
//	|                           |            | occurs. The segmentation condition is specified either by SegmentMaxDuration     |
//	|                           |            | or SegmentMaxSize. PLA SHOULD NOT<12> stop the data collector set when the size  |
//	|                           |            | of the active log file reaches SegmentMaxSize; rather, it creates a new file to  |
//	|                           |            | enable the data collector set to continue logging information.                   |
//	+---------------------------+------------+----------------------------------------------------------------------------------+
//	| SegmentMaxDuration        | RW         | Determines for how long a data collector set MUST run, in seconds, before a new  |
//	|                           |            | segment MUST be created. A value of 0 indicates that there is no segment time    |
//	|                           |            | limit. The default value is zero. Any unsigned long is a valid value for this    |
//	|                           |            | property.                                                                        |
//	+---------------------------+------------+----------------------------------------------------------------------------------+
//	| SegmentMaxSize            | RW         | Determines the maximum size, in megabytes, of a log file. When the size is       |
//	|                           |            | reached, a new log file MUST be created. A value of 0 indicates that there is    |
//	|                           |            | no segment size limit. The segment size is unlimited. The default value is zero. |
//	|                           |            | Any unsigned long is a valid value for this property.                            |
//	+---------------------------+------------+----------------------------------------------------------------------------------+
//	| SerialNumber              | RW         | The serial number of the data collector set. Each time the data collector set    |
//	|                           |            | runs it is assigned a serial number. The serial number for each data collector   |
//	|                           |            | set starts at 1 and is incremented each time the data collector set runs. Each   |
//	|                           |            | run of the data collector set has a serial number that is unique across all      |
//	|                           |            | runs of the data collector set. Accordingly, each run of the data collector      |
//	|                           |            | set has its own serial number and no two runs of the same data collector set     |
//	|                           |            | have the same serial number. However, it is possible that two different runs of  |
//	|                           |            | two different data collector sets will have the same serial number (the serial   |
//	|                           |            | number is unique among all runs of one data collector set; it is not unique      |
//	|                           |            | among all runs of all data collector sets). The serial number can optionally be  |
//	|                           |            | used by an AutoPathFormat, which will cause the serial number to be appended to  |
//	|                           |            | the name of the directory or files pertaining to each run of the data collector  |
//	|                           |            | set. Using the serial number as an AutoPathFormat prevents possible collisions   |
//	|                           |            | in directory or file naming. This property serves as a serial number across all  |
//	|                           |            | runs of a particular data collector set on a particular machine, not across all  |
//	|                           |            | data collector sets or all machines. Any unsigned long is a valid value for this |
//	|                           |            | property.                                                                        |
//	+---------------------------+------------+----------------------------------------------------------------------------------+
//	| Server                    | R          | Name of the server this data collector set belongs to.                           |
//	+---------------------------+------------+----------------------------------------------------------------------------------+
//	| Status                    | R          | Status of the data collector set.                                                |
//	+---------------------------+------------+----------------------------------------------------------------------------------+
//	| StopOnCompletion          | RW         | Determines whether a data collector set MUST stop when all data collectors       |
//	|                           |            | complete. A data collector completes when the first segment is collected.        |
//	|                           |            | The definition of completion depends on the data collector type, and is not      |
//	|                           |            | generally defined as the point at which the data collector has collected all     |
//	|                           |            | possible data. For an IConfigurationDataCollector, completion occurs when all    |
//	|                           |            | data has been collected. For an IAlertDataCollector or IApiTracingDataCollector, |
//	|                           |            | completion occurs immediately (that is, no data will be collected if             |
//	|                           |            | this property is set to true). For an IPerformanceCounterDataCollector or        |
//	|                           |            | ITraceDataCollector, completion occurs immediately if no limit is set on the     |
//	|                           |            | output file size or time spent writing to the output file. If there is a maximum |
//	|                           |            | file size per output file, or a maximum duration per output file, completion     |
//	|                           |            | occurs when the data collector has finished writing to the current file.         |
//	+---------------------------+------------+----------------------------------------------------------------------------------+
//	| Subdirectory              | RW         | Retrieves or sets a base subdirectory of the root path where the next instance   |
//	|                           |            | of the data collector set will write its logs.                                   |
//	+---------------------------+------------+----------------------------------------------------------------------------------+
//	| SubdirectoryFormat        | RW         | Determines what to include in the decoration the subdirectory name.              |
//	+---------------------------+------------+----------------------------------------------------------------------------------+
//	| SubdirectoryFormatPattern | RW         | If patterns are to be included in the decoration of subdirectory names,          |
//	|                           |            | determines the pattern to use.                                                   |
//	+---------------------------+------------+----------------------------------------------------------------------------------+
//	| Task                      | RW         | Name of the task that is executed when the data collector set stops or prior     |
//	|                           |            | to switching to a new segment. The name of the task needs to correspond to the   |
//	|                           |            | name of a job stored in the Task Scheduler. When a task is created, the creator  |
//	|                           |            | of a task specifies its name as a BSTR. More information on the names of Task    |
//	|                           |            | Scheduler jobs (referred to as paths in the Task Scheduler documentation)        |
//	|                           |            | is specified in [MS-TSCH] section 2.3.11. This documentation describes the       |
//	|                           |            | semantics of the task name and explains the restrictions on task names.          |
//	+---------------------------+------------+----------------------------------------------------------------------------------+
//	| TaskArguments             | RW         | If a task is to run, this specifies what arguments MUST be passed to it.         |
//	+---------------------------+------------+----------------------------------------------------------------------------------+
//	| TaskRunAsSelf             | RW         | When a Task Scheduler job is executed by the DataCollectorSet, this property     |
//	|                           |            | determines which user it will run as. If the property is set to true, the Task   |
//	|                           |            | Scheduler job runs as the same user that the DataCollectorSet is running as.     |
//	|                           |            | By default, this means the Task Scheduler job will run with System credentials.  |
//	|                           |            | Consequently, it is not advisable to set this property to true when the          |
//	|                           |            | task to be run is not fully trusted, unless the UserAccount property for the     |
//	|                           |            | DataCollectorSet has been carefully configured. When the property is set to      |
//	|                           |            | false, the Task Scheduler job will run with the credentials it was created       |
//	|                           |            | with. The mechanism in use here is delegation. When the creator of a data        |
//	|                           |            | collector set sets this property to true, this task is thereby granted the same  |
//	|                           |            | rights that the data collector set has. When the RunAsSelf property is set to    |
//	|                           |            | false, no delegation occurs. The task will run only with the permissions it was  |
//	|                           |            | created with. The credentials that the task runs with are initially created with |
//	|                           |            | SchRpcRegisterTask specified in [MS-TSCH] section 3.2.5.4.2 and can be updated   |
//	|                           |            | by SASetAccountInformation specified in [MS-TSCH] section 3.2.5.3.4.             |
//	+---------------------------+------------+----------------------------------------------------------------------------------+
//	| TaskUserTextArguments     | RW         | If a task is to run and the arguments, as specified in section 3.2.4.1.38,       |
//	|                           |            | include the {usertext} variable, this property determines the value of this      |
//	|                           |            | variable.                                                                        |
//	+---------------------------+------------+----------------------------------------------------------------------------------+
//	| UserAccount               | R          | Determines the user account under which the data collector set will run.         |
//	+---------------------------+------------+----------------------------------------------------------------------------------+
//	| Xml                       | R          | The XML representation of the data collector set.                                |
//	+---------------------------+------------+----------------------------------------------------------------------------------+
//
// A data collector set can be represented as an XML file, which can be used to serialize
// (using IDataCollectorSet::Xml (Get)3.2.4.1.46) and deserialize (using IDataCollectorSet::SetXml
// 3.2.4.1.58) it. The format of the XML that defines a data collector set is as follows.
// All the elements of the data collector set, as well as all child elements within
// the data collector set element, are specified in section 3.2.4.19:
//
// The nodes "Keyword", "Schedule", and "FolderAction" can have multiple copies—one
// for each keyword, schedule, or folder action, respectively. For each data collector,
// one node under the "DataCollectorSet" node is also added; the name of the node depends
// on the type of data collector, and is documented in the data collector section.
//
// The Keywords property is a safearray, which is translated into XML as a series of
// Keyword nodes. For example, if Keywords is set to {"A", "B", "C"}, there are three
// Keyword nodes, one for each keyword.
//
// Similarly, Schedules is a collection of Schedule objects, which means that if the
// Schedules property contains three schedules, three nodes called "Schedule" are created
// for each schedule.
//
// DataCollectors is not in XML for the same reason as Schedules. However, because data
// collectors have types instead of having a number of "DataCollector" nodes, there
// are a number of typed data collector nodes. For example, "AlertDataCollector" or
// "PerformanceCounterDataCollector".
//
// Methods in RPC Opnum order:
//
//	+---------------------------------+----------------------------------------------------------------------------------+
//	|                                 |                                                                                  |
//	|             METHOD              |                                   DESCRIPTION                                    |
//	|                                 |                                                                                  |
//	+---------------------------------+----------------------------------------------------------------------------------+
//	+---------------------------------+----------------------------------------------------------------------------------+
//	| DataCollectors (Get)            | List of data collectors in this data collector set. Opnum: 7                     |
//	+---------------------------------+----------------------------------------------------------------------------------+
//	| Duration (Get)                  | Retrieves the Duration property. Opnum: 8                                        |
//	+---------------------------------+----------------------------------------------------------------------------------+
//	| Duration (Put)                  | Sets the Duration property. Opnum: 9                                             |
//	+---------------------------------+----------------------------------------------------------------------------------+
//	| Description (Get)               | Retrieves the Description property. Opnum: 10                                    |
//	+---------------------------------+----------------------------------------------------------------------------------+
//	| Description (Put)               | Sets the Description property. Opnum: 11                                         |
//	+---------------------------------+----------------------------------------------------------------------------------+
//	| DescriptionUnresolved (Get)     | Retrieves the DescriptionUnresolved property. Opnum: 12                          |
//	+---------------------------------+----------------------------------------------------------------------------------+
//	| DisplayName (Get)               | Retrieves the display name of the data collector set . Opnum: 13                 |
//	+---------------------------------+----------------------------------------------------------------------------------+
//	| DisplayName (Put)               | Sets the DisplayName property. Opnum: 14                                         |
//	+---------------------------------+----------------------------------------------------------------------------------+
//	| DisplayNameUnresolved (Get)     | Receives the display name of the data collector set in its original form. Opnum: |
//	|                                 |                                                                               15 |
//	+---------------------------------+----------------------------------------------------------------------------------+
//	| Keywords (Get)                  | Retrieves the Keywords property. Opnum: 16                                       |
//	+---------------------------------+----------------------------------------------------------------------------------+
//	| Keywords (Put)                  | Sets the Keywords property. Opnum: 17                                            |
//	+---------------------------------+----------------------------------------------------------------------------------+
//	| LatestOutputLocation (Get)      | Retrieves the LatestOutputLocation property. Opnum: 18                           |
//	+---------------------------------+----------------------------------------------------------------------------------+
//	| LatestOutputLocation (Put)      | Sets the LatestOutputLocation property. Opnum: 19                                |
//	+---------------------------------+----------------------------------------------------------------------------------+
//	| Name (Get)                      | Retrieves the Name property. Opnum: 20                                           |
//	+---------------------------------+----------------------------------------------------------------------------------+
//	| OutputLocation (Get)            | Retrieves the OutputLocation property. Opnum: 21                                 |
//	+---------------------------------+----------------------------------------------------------------------------------+
//	| RootPath (Get)                  | Retrieves the RootPath property. Opnum: 22                                       |
//	+---------------------------------+----------------------------------------------------------------------------------+
//	| RootPath (Put)                  | Sets the RootPath property. Opnum: 23                                            |
//	+---------------------------------+----------------------------------------------------------------------------------+
//	| Segment (Get)                   | Retrieves the Segment property. Opnum: 24                                        |
//	+---------------------------------+----------------------------------------------------------------------------------+
//	| Segment (Put)                   | Sets the Segment property. Opnum: 25                                             |
//	+---------------------------------+----------------------------------------------------------------------------------+
//	| SegmentMaxDuration (Get)        | Retrieves the SegmentMaxDuration property. Opnum: 26                             |
//	+---------------------------------+----------------------------------------------------------------------------------+
//	| SegmentMaxDuration (Put)        | Sets the SegmentMaxDuration property. Opnum: 27                                  |
//	+---------------------------------+----------------------------------------------------------------------------------+
//	| SegmentMaxSize (Get)            | Retrieves the SegmentMaxSize property. Opnum: 28                                 |
//	+---------------------------------+----------------------------------------------------------------------------------+
//	| SegmentMaxSize (Put)            | Sets the SegmentMaxSize property. Opnum: 29                                      |
//	+---------------------------------+----------------------------------------------------------------------------------+
//	| SerialNumber (Get)              | Retrieves the SerialNumber property. Opnum: 30                                   |
//	+---------------------------------+----------------------------------------------------------------------------------+
//	| SerialNumber (Put)              | Sets the SerialNumber property. Opnum: 31                                        |
//	+---------------------------------+----------------------------------------------------------------------------------+
//	| Server (Get)                    | Retrieves the Server property. Opnum: 32                                         |
//	+---------------------------------+----------------------------------------------------------------------------------+
//	| Status (Get)                    | Retrieves the Status property. Opnum: 33                                         |
//	+---------------------------------+----------------------------------------------------------------------------------+
//	| Subdirectory (Get)              | Retrieves the Subdirectory property. Opnum: 34                                   |
//	+---------------------------------+----------------------------------------------------------------------------------+
//	| Subdirectory (Put)              | Sets the Subdirectory property. Opnum: 35                                        |
//	+---------------------------------+----------------------------------------------------------------------------------+
//	| SubdirectoryFormat (Get)        | Retrieves the SubdirectoryFormat property. Opnum: 36                             |
//	+---------------------------------+----------------------------------------------------------------------------------+
//	| SubdirectoryFormat (Put)        | Sets the SubdirectoryFormat property. Opnum: 37                                  |
//	+---------------------------------+----------------------------------------------------------------------------------+
//	| SubdirectoryFormatPattern (Get) | Retrieves the SubdirectoryFormatPattern property. Opnum: 38                      |
//	+---------------------------------+----------------------------------------------------------------------------------+
//	| SubdirectoryFormatPattern (Put) | Sets the SubdirectoryFormatPattern property. Opnum: 39                           |
//	+---------------------------------+----------------------------------------------------------------------------------+
//	| Task (Get)                      | Retrieves the Task property. Opnum: 40                                           |
//	+---------------------------------+----------------------------------------------------------------------------------+
//	| Task (Put)                      | Sets the Task property. Opnum: 41                                                |
//	+---------------------------------+----------------------------------------------------------------------------------+
//	| TaskRunAsSelf (Get)             | Retrieves the TaskRunAsSelf property. Opnum: 42                                  |
//	+---------------------------------+----------------------------------------------------------------------------------+
//	| TaskRunAsSelf (Put)             | Sets the TaskRunAsSelf property. Opnum: 43                                       |
//	+---------------------------------+----------------------------------------------------------------------------------+
//	| TaskArguments (Get)             | Retrieves the TaskArguments property. Opnum: 44                                  |
//	+---------------------------------+----------------------------------------------------------------------------------+
//	| TaskArguments (Put)             | Sets the TaskArguments property. Opnum: 45                                       |
//	+---------------------------------+----------------------------------------------------------------------------------+
//	| TaskUserTextArguments (Get)     | Retrieves the TaskUserTextArguments property. Opnum: 46                          |
//	+---------------------------------+----------------------------------------------------------------------------------+
//	| TaskUserTextArguments (Put)     | Sets the TaskUserTextArguments property. Opnum: 47                               |
//	+---------------------------------+----------------------------------------------------------------------------------+
//	| Schedules (Get)                 | Retrieves the Schedules property. Opnum: 48                                      |
//	+---------------------------------+----------------------------------------------------------------------------------+
//	| SchedulesEnabled (Get)          | Retrieves the SchedulesEnabled property. Opnum: 49                               |
//	+---------------------------------+----------------------------------------------------------------------------------+
//	| SchedulesEnabled (Put)          | Sets the SchedulesEnabled property. Opnum: 50                                    |
//	+---------------------------------+----------------------------------------------------------------------------------+
//	| UserAccount (Get)               | Retrieves the UserAccount property. Opnum: 51                                    |
//	+---------------------------------+----------------------------------------------------------------------------------+
//	| Xml (Get)                       | Retrieves the Xml property. Opnum: 52                                            |
//	+---------------------------------+----------------------------------------------------------------------------------+
//	| Security (Get)                  | Retrieves the Security property. Opnum: 53                                       |
//	+---------------------------------+----------------------------------------------------------------------------------+
//	| Security (Put)                  | Sets the Security property. Opnum: 54                                            |
//	+---------------------------------+----------------------------------------------------------------------------------+
//	| StopOnCompletion (Get)          | Retrieves the StopOnCompletion property. Opnum: 55                               |
//	+---------------------------------+----------------------------------------------------------------------------------+
//	| StopOnCompletion (Put)          | Sets the StopOnCompletion property. Opnum: 56                                    |
//	+---------------------------------+----------------------------------------------------------------------------------+
//	| DataManager (Get)               | Retrieves the DataManager property. Opnum: 57                                    |
//	+---------------------------------+----------------------------------------------------------------------------------+
//	| SetCredentials                  | Specifies the user account under which the data collector set runs. Opnum: 58    |
//	+---------------------------------+----------------------------------------------------------------------------------+
//	| Query                           | Loads the properties of a data collector set from storage to memory. Opnum: 59   |
//	+---------------------------------+----------------------------------------------------------------------------------+
//	| Commit                          | Updates, validates, or saves a data collector set or flushes the event trace     |
//	|                                 | data collectors of a data collector set. Opnum: 60                               |
//	+---------------------------------+----------------------------------------------------------------------------------+
//	| Delete                          | Removes the data collector set from storage if it is not running. Opnum: 61      |
//	+---------------------------------+----------------------------------------------------------------------------------+
//	| Start                           | Manually starts the data collector set. Opnum: 62                                |
//	+---------------------------------+----------------------------------------------------------------------------------+
//	| Stop                            | Manually stops the data collector set. Opnum: 63                                 |
//	+---------------------------------+----------------------------------------------------------------------------------+
//	| SetXml                          | Sets the property values of the data collector set based on an XML file. Opnum:  |
//	|                                 |                                                                               64 |
//	+---------------------------------+----------------------------------------------------------------------------------+
//	| SetValue                        | Sets a user-defined value. Opnum: 65                                             |
//	+---------------------------------+----------------------------------------------------------------------------------+
//	| GetValue                        | Retrieves a user-defined value. Opnum: 66                                        |
//	+---------------------------------+----------------------------------------------------------------------------------+
//
// Opnums 0, 1, and 2 are reserved for the IUnknown interface. Opnums 3, 4, 5, and 6
// are reserved for the IDispatch interface.
type DataCollectorSet dcom.InterfacePointer

func (o *DataCollectorSet) InterfacePointer() *dcom.InterfacePointer {
	return (*dcom.InterfacePointer)(o)
}

func (o *DataCollectorSet) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if o.Data != nil && o.DataCount == 0 {
		o.DataCount = uint32(len(o.Data))
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}

func (o *DataCollectorSet) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *DataCollectorSet) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *DataCollectorSet) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// TraceDataProviderCollection structure represents ITraceDataProviderCollection RPC structure.
//
// The ITraceDataProviderCollection interface is used to manage a collection of TraceDataProvider
// objects.
//
// The following properties MUST be implemented by the objects that implement the ITraceDataProviderCollection
// interface.
//
//	+----------+------------+----------------------------------------------------------------------------------+
//	|          |            |                                                                                  |
//	| PROPERTY | READ/WRITE |                                   DESCRIPTION                                    |
//	|          |            |                                                                                  |
//	+----------+------------+----------------------------------------------------------------------------------+
//	+----------+------------+----------------------------------------------------------------------------------+
//	| _NewEnum | R          | An enumeration object of type IEnumVariant containing a snapshot of the          |
//	|          |            | ITraceDataProvider objects in this collection. The enumeration object is         |
//	|          |            | specified in [MS-OAUT] section 3.3.                                              |
//	+----------+------------+----------------------------------------------------------------------------------+
//	| Count    | R          | Number of trace data providers in this collection.                               |
//	+----------+------------+----------------------------------------------------------------------------------+
//	| Item     | R          | Retrieves the requested trace data provider from the collection.                 |
//	+----------+------------+----------------------------------------------------------------------------------+
//
// Methods in RPC Opnum Order
//
//	+--------------------------------+----------------------------------------------------------------------------------+
//	|                                |                                                                                  |
//	|             METHOD             |                                   DESCRIPTION                                    |
//	|                                |                                                                                  |
//	+--------------------------------+----------------------------------------------------------------------------------+
//	+--------------------------------+----------------------------------------------------------------------------------+
//	| Count (Get)                    | Retrieves the Count property. Opnum: 7                                           |
//	+--------------------------------+----------------------------------------------------------------------------------+
//	| Item (Get)                     | Retrieves the Item property. Opnum: 8                                            |
//	+--------------------------------+----------------------------------------------------------------------------------+
//	| _NewEnum (Get)                 | Retrieves the NewEnum property. Opnum: 9                                         |
//	+--------------------------------+----------------------------------------------------------------------------------+
//	| Add                            | Adds a trace provider to the collection. Opnum: 10                               |
//	+--------------------------------+----------------------------------------------------------------------------------+
//	| Remove                         | Removes a trace provider from the collection. Opnum: 11                          |
//	+--------------------------------+----------------------------------------------------------------------------------+
//	| Clear                          | Removes all trace providers from the collection. Opnum: 12                       |
//	+--------------------------------+----------------------------------------------------------------------------------+
//	| AddRange                       | Adds one or more trace providers to the collection. Opnum: 13                    |
//	+--------------------------------+----------------------------------------------------------------------------------+
//	| CreateTraceDataProvider        | Creates a trace data provider object. Opnum: 14                                  |
//	+--------------------------------+----------------------------------------------------------------------------------+
//	| GetTraceDataProviders          | Populates the collection with the registered trace providers. Opnum: 15          |
//	+--------------------------------+----------------------------------------------------------------------------------+
//	| GetTraceDataProvidersByProcess | Populates the collection with the list of providers that were registered by the  |
//	|                                | given process. Opnum: 16                                                         |
//	+--------------------------------+----------------------------------------------------------------------------------+
//
// Opnums 0, 1, and 2 are reserved for the IUnknown interface. Opnums 3, 4, 5, and 6
// are reserved for the IDispatch interface.
type TraceDataProviderCollection dcom.InterfacePointer

func (o *TraceDataProviderCollection) InterfacePointer() *dcom.InterfacePointer {
	return (*dcom.InterfacePointer)(o)
}

func (o *TraceDataProviderCollection) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if o.Data != nil && o.DataCount == 0 {
		o.DataCount = uint32(len(o.Data))
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}

func (o *TraceDataProviderCollection) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *TraceDataProviderCollection) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *TraceDataProviderCollection) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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
