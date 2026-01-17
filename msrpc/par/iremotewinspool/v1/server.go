package iremotewinspool

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
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
)

// IRemoteWinspool server interface.
type RemoteWinspoolServer interface {

	// RpcAsyncOpenPrinter retrieves a handle to a specified printer, port, print job or
	// print server. A client uses this method to obtain a print handle to an existing printer
	// on a remote computer.
	//
	// The counterpart of this method in the Print System Remote Protocol is RpcOpenPrinterEx.
	// All parameters not defined below are specified in [MS-RPRN] section 3.1.4.2.14.
	//
	// Return Values: This method returns zero to indicate successful completion or a nonzero
	// Win32 error code ([MS-ERREF] section 2.2) to indicate failure. The client MUST treat
	// any nonzero return value as a fatal error.
	//
	// Exceptions Thrown: This method MUST NOT throw any exceptions other than those that
	// are thrown by the underlying RPC protocol [MS-RPCE].
	//
	// This method MUST adhere to the parameter validation, processing, and response requirements
	// that are specified in [MS-RPRN] section 3.1.4.2.14.
	OpenPrinter(context.Context, *OpenPrinterRequest) (*OpenPrinterResponse, error)

	// RpcAsyncAddPrinter installs a printer on the print server.
	//
	// The counterpart of this method in the Print System Remote Protocol is RpcAddPrinterEx.
	// All parameters not defined below are specified in [MS-RPRN] section 3.1.4.2.15.
	//
	// Return Values: This method returns zero to indicate successful completion or a nonzero
	// Win32 error code ([MS-ERREF] section 2.2) to indicate failure. The client MUST treat
	// any nonzero return value as a fatal error.
	//
	// Exceptions Thrown: This method MUST NOT throw any exceptions other than those that
	// are thrown by the underlying RPC protocol [MS-RPCE].
	//
	// This method MUST adhere to the parameter validation, processing, and response requirements
	// that are specified in [MS-RPRN] section 3.1.4.2.15.
	AddPrinter(context.Context, *AddPrinterRequest) (*AddPrinterResponse, error)

	// RpcAsyncSetJob pauses, resumes, cancels, or restarts a print job on a specified printer.
	// This method can also set print job parameters, including the job priority and document
	// name.
	//
	// The counterpart of this method in the Print System Remote Protocol is RpcSetJob.
	// All parameters not defined below are specified in [MS-RPRN] section 3.1.4.3.1.
	//
	// Return Values: This method returns zero to indicate successful completion or a nonzero
	// Win32 error code ([MS-ERREF] section 2.2) to indicate failure. The client MUST treat
	// any nonzero return value as a fatal error.
	//
	// Exceptions Thrown: This method MUST NOT throw any exceptions other than those that
	// are thrown by the underlying RPC protocol [MS-RPCE].
	//
	// This method MUST adhere to the parameter validation, processing, and response requirements
	// that are specified in [MS-RPRN] section 3.1.4.3.1.
	SetJob(context.Context, *SetJobRequest) (*SetJobResponse, error)

	// RpcAsyncGetJob retrieves information about a specified print job on a specified printer.
	//
	// The counterpart of this method in the Print System Remote Protocol is RpcGetJob.
	// All parameters not defined below are specified in [MS-RPRN] section 3.1.4.3.2.
	//
	// Return Values: This method returns zero to indicate successful completion or a nonzero
	// Win32 error code ([MS-ERREF] section 2.2) to indicate failure. Aside from the specific
	// nonzero return values documented in section 3.1.4, the client MUST treat any nonzero
	// return value as a fatal error.
	//
	// Exceptions Thrown: This method MUST NOT throw any exceptions other than those that
	// are thrown by the underlying RPC protocol [MS-RPCE].
	//
	// This method MUST adhere to the parameter validation, processing, and response requirements
	// that are specified in [MS-RPRN] section 3.1.4.3.2.
	GetJob(context.Context, *GetJobRequest) (*GetJobResponse, error)

	// RpcAsyncEnumJobs retrieves information about a specified set of print jobs on a specified
	// printer.
	//
	// The counterpart of this method in the Print System Remote Protocol is RpcEnumJobs.
	// All parameters not defined below are specified in [MS-RPRN] section 3.1.4.3.3.
	//
	// Return Values: This method returns zero to indicate successful completion or a nonzero
	// Win32 error code ([MS-ERREF] section 2.2) to indicate failure. Aside from the specific
	// nonzero return values documented in section 3.1.4, the client MUST treat any nonzero
	// return value as a fatal error.
	//
	// Exceptions Thrown: This method MUST NOT throw any exceptions other than those that
	// are thrown by the underlying RPC protocol [MS-RPCE].
	//
	// This method MUST adhere to the parameter validation, processing, and response requirements
	// that are specified in [MS-RPRN] section 3.1.4.3.3.
	EnumJobs(context.Context, *EnumJobsRequest) (*EnumJobsResponse, error)

	// RpcAsyncAddJob does not perform any function, but returns ERROR_INVALID_PARAMETER.
	//
	// The counterpart of this method in the Print System Remote Protocol is RpcAddJob.
	// All parameters not defined below are specified in [MS-RPRN] section 3.1.4.3.4.
	//
	// Return Values: This method MUST return ERROR_INVALID_PARAMETER ([MS-ERREF] section
	// 2.2).
	//
	// This method MUST be implemented to ensure compatibility with protocol clients.
	AddJob(context.Context, *AddJobRequest) (*AddJobResponse, error)

	// RpcAsyncScheduleJob does not perform any function, but returns ERROR_SPL_NO_ADDJOB.
	//
	// The counterpart of this method in the Print System Remote Protocol is RpcScheduleJob.
	// All parameters not defined below are specified in [MS-RPRN] section 3.1.4.3.5.
	//
	// Return Values: This method MUST return ERROR_SPL_NO_ADDJOB ([MS-ERREF] section 2.2).
	//
	// This method MUST be implemented to ensure compatibility with protocol clients.
	ScheduleJob(context.Context, *ScheduleJobRequest) (*ScheduleJobResponse, error)

	// RpcAsyncDeletePrinter deletes the specified printer object.
	//
	// The counterpart of this method in the Print System Remote Protocol is RpcDeletePrinter.
	// All parameters not defined below are specified in [MS-RPRN] section 3.1.4.2.4.
	//
	// The client MUST call RpcAsyncClosePrinter (section 3.1.4.1.10) with the PRINTER_HANDLE
	// ([MS-RPRN] section 2.2.1.1.4) represented by the hPrinter parameter after calling
	// RpcAsyncDeletePrinter.
	//
	// Return Values: This method returns zero to indicate successful completion or a nonzero
	// Win32 error code ([MS-ERREF] section 2.2) to indicate failure. The client MUST treat
	// any nonzero return value as a fatal error.
	//
	// Exceptions Thrown: This method MUST NOT throw any exceptions other than those that
	// are thrown by the underlying RPC protocol [MS-RPCE].
	//
	// This method MUST adhere to the parameter validation, processing, and response requirements
	// that are specified in [MS-RPRN] section 3.1.4.2.4.
	DeletePrinter(context.Context, *DeletePrinterRequest) (*DeletePrinterResponse, error)

	// RpcAsyncSetPrinter sets configuration information, initialization data, and security
	// information of the specified printer to the values contained in the method parameters.
	// It can also perform an action to change the active status of the printer.
	//
	// The counterpart of this method in the Print System Remote Protocol is RpcSetPrinter.
	// All parameters not defined below are specified in [MS-RPRN] section 3.1.4.2.5.
	//
	// Return Values: This method returns zero to indicate successful completion or a nonzero
	// Win32 error code ([MS-ERREF] section 2.2) to indicate failure. The client MUST treat
	// any nonzero return value as a fatal error.
	//
	// Exceptions Thrown: This method MUST NOT throw any exceptions other than those that
	// are thrown by the underlying RPC protocol [MS-RPCE].
	//
	// This method MUST adhere to the parameter validation, processing, and response requirements
	// that are specified in [MS-RPRN] section 3.1.4.2.5.
	SetPrinter(context.Context, *SetPrinterRequest) (*SetPrinterResponse, error)

	// RpcAsyncGetPrinter retrieves information about a specified printer.
	//
	// The counterpart of this method in the Print System Remote Protocol is RpcGetPrinter.
	// All parameters not defined below are specified in [MS-RPRN] section 3.1.4.2.6.
	//
	// Return Values: This method returns zero to indicate successful completion or a nonzero
	// Win32 error code ([MS-ERREF] section 2.2) to indicate failure. Aside from the specific
	// nonzero return values documented in section 3.1.4, the client MUST treat any nonzero
	// return value as a fatal error.
	//
	// Exceptions Thrown: This method MUST NOT throw any exceptions other than those that
	// are thrown by the underlying RPC protocol [MS-RPCE].
	//
	// This method MUST adhere to the parameter validation, processing, and response requirements
	// that are specified in [MS-RPRN] section 3.1.4.2.6.
	GetPrinter(context.Context, *GetPrinterRequest) (*GetPrinterResponse, error)

	// RpcStartDocPrinter notifies a specified printer that a document is being spooled
	// for printing.
	//
	// The counterpart of this method in the Print System Remote Protocol is RpcStartDocPrinter.
	// All parameters not defined below are specified in [MS-RPRN] section 3.1.4.9.1.
	//
	// Return Values: This method returns zero to indicate successful completion or a nonzero
	// Win32 error code ([MS-ERREF] section 2.2) to indicate failure. The client MUST treat
	// any nonzero return value as a fatal error.
	//
	// Exceptions Thrown: This method MUST NOT throw any exceptions other than those that
	// are thrown by the underlying RPC protocol [MS-RPCE].
	//
	// This method MUST adhere to the parameter validation, processing, and response requirements
	// that are specified in [MS-RPRN] section 3.1.4.9.1.
	StartDocPrinter(context.Context, *StartDocPrinterRequest) (*StartDocPrinterResponse, error)

	// RpcAsyncStartPagePrinter notifies a specified printer that a page is about to be
	// printed.
	//
	// The counterpart of this method in the Print System Remote Protocol is RpcStartPagePrinter.
	// All parameters not defined below are specified in [MS-RPRN] section 3.1.4.9.2.
	//
	// Return Values: This method returns zero to indicate successful completion or a nonzero
	// Win32 error code ([MS-ERREF] section 2.2) to indicate failure. The client MUST treat
	// any nonzero return value as a fatal error.
	//
	// Exceptions Thrown: This method MUST NOT throw any exceptions other than those that
	// are thrown by the underlying RPC protocol [MS-RPCE].
	//
	// This method MUST adhere to the parameter validation, processing, and response requirements
	// that are specified in [MS-RPRN] section 3.1.4.9.2.
	StartPagePrinter(context.Context, *StartPagePrinterRequest) (*StartPagePrinterResponse, error)

	// RpcAsyncWritePrinter adds data to the file representing the spool file for a specified
	// printer, if the spooling option is turned on; or it sends data to the device directly,
	// if the printer is configured for direct printing.
	//
	// The counterpart of this method in the Print System Remote Protocol is RpcWritePrinter.
	// All parameters not defined below are specified in [MS-RPRN] section 3.1.4.9.3.
	//
	// Return Values: This method returns zero to indicate successful completion or a nonzero
	// Win32 error code ([MS-ERREF] section 2.2) to indicate failure. The client MUST treat
	// any nonzero return value as a fatal error.
	//
	// Exceptions Thrown: This method MUST NOT throw any exceptions other than those that
	// are thrown by the underlying RPC protocol [MS-RPCE].
	//
	// This method MUST adhere to the parameter validation, processing, and response requirements
	// that are specified in [MS-RPRN] section 3.1.4.9.3.
	WritePrinter(context.Context, *WritePrinterRequest) (*WritePrinterResponse, error)

	// RpcAsyncEndPagePrinter notifies a specified printer that the application is at the
	// end of a page in a print job.
	//
	// The counterpart of this method in the Print System Remote Protocol is RpcEndPagePrinter.
	// All parameters not defined below are specified in [MS-RPRN] section 3.1.4.9.4.
	//
	// Return Values: This method returns zero to indicate successful completion or a nonzero
	// Win32 error code ([MS-ERREF] section 2.2) to indicate failure. The client MUST treat
	// any nonzero return value as a fatal error.
	//
	// Exceptions Thrown: This method MUST NOT throw any exceptions other than those that
	// are thrown by the underlying RPC protocol [MS-RPCE].
	//
	// This method MUST adhere to the parameter validation, processing, and response requirements
	// that are specified in [MS-RPRN] section 3.1.4.9.4.
	EndPagePrinter(context.Context, *EndPagePrinterRequest) (*EndPagePrinterResponse, error)

	// RpcAsyncEndDocPrinter signals the completion of the current print job on a specified
	// printer.
	//
	// The counterpart of this method in the Print System Remote Protocol is RpcEndDocPrinter.
	// All parameters not defined below are specified in [MS-RPRN] section 3.1.4.9.7.
	//
	// Return Values: This method returns zero to indicate successful completion or a nonzero
	// Win32 error code ([MS-ERREF] section 2.2) to indicate failure. The client MUST treat
	// any nonzero return value as a fatal error.
	//
	// Exceptions Thrown: This method MUST NOT throw any exceptions other than those that
	// are thrown by the underlying RPC protocol [MS-RPCE].
	//
	// This method MUST adhere to the parameter validation, processing, and response requirements
	// that are specified in [MS-RPRN] section 3.1.4.9.7.
	EndDocPrinter(context.Context, *EndDocPrinterRequest) (*EndDocPrinterResponse, error)

	// RpcAsyncAbortPrinter aborts the current document on a specified printer.
	//
	// The counterpart of this method in the Print System Remote Protocol is RpcAbortPrinter.
	// All parameters not defined below are specified in [MS-RPRN] section 3.1.4.9.5.
	//
	// Return Values: This method returns zero to indicate successful completion or a nonzero
	// Win32 error code ([MS-ERREF] section 2.2) to indicate failure. The client MUST treat
	// any nonzero return value as a fatal error.
	//
	// Exceptions Thrown: This method MUST NOT throw any exceptions other than those that
	// are thrown by the underlying RPC protocol [MS-RPCE].
	//
	// This method MUST adhere to the parameter validation, processing, and response requirements
	// that are specified in [MS-RPRN] section 3.1.4.9.5.
	AbortPrinter(context.Context, *AbortPrinterRequest) (*AbortPrinterResponse, error)

	// RpcAsyncGetPrinterData retrieves configuration data for the specified printer or
	// print server.
	//
	// The counterpart of this method in the Print System Remote Protocol is RpcGetPrinterData.
	// All parameters not defined below are specified in [MS-RPRN] section 3.1.4.2.7.
	//
	// Return Values: This method returns zero to indicate successful completion or a nonzero
	// Win32 error code ([MS-ERREF] section 2.2) to indicate failure. Aside from the specific
	// nonzero return values documented in section 3.1.4, the client MUST treat any nonzero
	// return value as a fatal error.
	//
	// Exceptions Thrown: This method MUST NOT throw any exceptions other than those that
	// are thrown by the underlying RPC protocol [MS-RPCE].
	//
	// This method MUST adhere to the parameter validation, processing, and response requirements
	// that are specified in [MS-RPRN] section 3.1.4.2.7.
	GetPrinterData(context.Context, *GetPrinterDataRequest) (*GetPrinterDataResponse, error)

	// RpcAsyncGetPrinterDataEx retrieves configuration data for the specified printer or
	// print server. This method extends RpcAsyncGetPrinterData (section 3.1.4.1.6) and
	// can retrieve values stored under the specified key by RpcAsyncSetPrinterDataEx (section
	// 3.1.4.1.9).
	//
	// The counterpart of this method in the Print System Remote Protocol is RpcGetPrinterDataEx.
	// All parameters not defined below are specified in [MS-RPRN] section 3.1.4.2.19.
	//
	// Return Values: This method returns zero to indicate successful completion or a nonzero
	// Win32 error code ([MS-ERREF] section 2.2) to indicate failure. Aside from the specific
	// nonzero return values documented in section 3.1.4, the client MUST treat any nonzero
	// return value as a fatal error.
	//
	// Exceptions Thrown: This method MUST NOT throw any exceptions other than those that
	// are thrown by the underlying RPC protocol [MS-RPCE].
	//
	// This method MUST adhere to the parameter validation, processing, and response requirements
	// that are specified in [MS-RPRN] section 3.1.4.2.19.
	GetPrinterDataEx(context.Context, *GetPrinterDataExRequest) (*GetPrinterDataExResponse, error)

	// RpcAsyncSetPrinterData sets configuration data for the specified printer or print
	// server.
	//
	// The counterpart of this method in the Print System Remote Protocol is RpcSetPrinterData.
	// All parameters not defined below are specified in [MS-RPRN] section 3.1.4.2.8.
	//
	// Return Values: This method returns zero to indicate successful completion or a nonzero
	// Win32 error code ([MS-ERREF] section 2.2) to indicate failure. The client MUST treat
	// any nonzero return value as a fatal error.
	//
	// Exceptions Thrown: This method MUST NOT throw any exceptions other than those that
	// are thrown by the underlying RPC protocol [MS-RPCE].
	//
	// This method MUST adhere to the parameter validation, processing, and response requirements
	// that are specified in [MS-RPRN] section 3.1.4.2.8.
	SetPrinterData(context.Context, *SetPrinterDataRequest) (*SetPrinterDataResponse, error)

	// RpcAsyncSetPrinterDataEx sets configuration data for the specified printer or print
	// server. This method is similar to RpcAsyncSetPrinterData (section 3.1.4.1.8) but
	// also allows the caller to specify the registry key under which to store the data.
	//
	// The counterpart of this method in the Print System Remote Protocol is RpcSetPrinterDataEx.
	// All parameters not defined below are specified in [MS-RPRN] section 3.1.4.2.18.
	//
	// Return Values: This method returns zero to indicate successful completion or a nonzero
	// Win32 error code ([MS-ERREF] section 2.2) to indicate failure. The client MUST treat
	// any nonzero return value as a fatal error.
	//
	// Exceptions Thrown: This method MUST NOT throw any exceptions other than those that
	// are thrown by the underlying RPC protocol [MS-RPCE].
	//
	// This method MUST adhere to the parameter validation, processing, and response requirements
	// specified in [MS-RPRN] section 3.1.4.2.18.
	SetPrinterDataEx(context.Context, *SetPrinterDataExRequest) (*SetPrinterDataExResponse, error)

	// RpcAsyncClosePrinter closes a handle to a printer, server, job, or port object that
	// was previously opened by either RpcAsyncOpenPrinter (section 3.1.4.1.1) or RpcAsyncAddPrinter
	// (section 3.1.4.1.2).
	//
	// The counterpart of this method in the Print System Remote Protocol is RpcClosePrinter.
	// All parameters not defined below are specified in [MS-RPRN] section 3.1.4.2.9.
	//
	// Return Values: This method returns zero to indicate successful completion or a nonzero
	// Win32 error code ([MS-ERREF] section 2.2) to indicate failure. The client MUST treat
	// any nonzero return value as a fatal error.
	//
	// Exceptions Thrown: This method MUST NOT throw any exceptions other than those that
	// are thrown by the underlying RPC protocol [MS-RPCE].
	//
	// This method MUST adhere to the parameter validation, processing, and response requirements
	// that are specified in [MS-RPRN] section 3.1.4.2.9.
	ClosePrinter(context.Context, *ClosePrinterRequest) (*ClosePrinterResponse, error)

	// RpcAsyncAddForm adds a form name to the list of supported printer forms.
	//
	// The counterpart of this method in the Print System Remote Protocol is RpcAddForm.
	// All parameters not defined below are specified in [MS-RPRN] section 3.1.4.5.1.
	//
	// Return Values: This method returns zero to indicate successful completion or a nonzero
	// Win32 error code ([MS-ERREF] section 2.2) to indicate failure. The client MUST treat
	// any nonzero return value as a fatal error.
	//
	// Exceptions Thrown: This method MUST NOT throw any exceptions other than those that
	// are thrown by the underlying RPC protocol [MS-RPCE].
	//
	// This method MUST adhere to the parameter validation, processing, and response requirements
	// that are specified in [MS-RPRN] section 3.1.4.5.1.
	AddForm(context.Context, *AddFormRequest) (*AddFormResponse, error)

	// RpcAsyncDeleteForm removes a form name from the list of supported printer forms.
	//
	// The counterpart of this method in the Print System Remote Protocol is RpcDeleteForm.
	// All parameters not defined below are specified in [MS-RPRN] section 3.1.4.5.2.
	//
	// Return Values: This method returns zero to indicate successful completion or a nonzero
	// Win32 error code ([MS-ERREF] section 2.2) to indicate failure. The client MUST treat
	// any nonzero return value as a fatal error.
	//
	// Exceptions Thrown: This method MUST NOT throw any exceptions other than those that
	// are thrown by the underlying RPC protocol [MS-RPCE].
	//
	// This method MUST adhere to the parameter validation, processing, and response requirements
	// that are specified in [MS-RPRN] section 3.1.4.5.2.
	DeleteForm(context.Context, *DeleteFormRequest) (*DeleteFormResponse, error)

	// RpcAsyncGetForm retrieves information about a specified printer form.
	//
	// The counterpart of this method in the Print System Remote Protocol is RpcGetForm.
	// All parameters not defined below are specified in [MS-RPRN] section 3.1.4.5.3.
	//
	// Return Values: This method returns zero to indicate successful completion or a nonzero
	// Win32 error code ([MS-ERREF] section 2.2) to indicate failure. Aside from the specific
	// nonzero return values documented in section 3.1.4, the client MUST treat any nonzero
	// return value as a fatal error.
	//
	// Exceptions Thrown: This method MUST NOT throw any exceptions other than those that
	// are thrown by the underlying RPC protocol [MS-RPCE].
	//
	// This method MUST adhere to the parameter validation, processing, and response requirements
	// that are specified in [MS-RPRN] section 3.1.4.5.3.
	GetForm(context.Context, *GetFormRequest) (*GetFormResponse, error)

	// RpcAsyncSetForm sets the printer form information for the specified printer.
	//
	// The counterpart of this method in the Print System Remote Protocol is RpcSetForm.
	// All parameters not defined below are specified in [MS-RPRN] section 3.1.4.5.4.
	//
	// Return Values: This method returns zero to indicate successful completion or a nonzero
	// Win32 error code ([MS-ERREF] section 2.2) to indicate failure. The client MUST treat
	// any nonzero return value as a fatal error.
	//
	// Exceptions Thrown: This method MUST NOT throw any exceptions other than those that
	// are thrown by the underlying RPC protocol [MS-RPCE].
	//
	// This method MUST adhere to the parameter validation, processing, and response requirements
	// that are specified in [MS-RPRN] section 3.1.4.5.4.
	SetForm(context.Context, *SetFormRequest) (*SetFormResponse, error)

	// RpcAsyncEnumForms enumerates the printer forms that the specified printer supports.
	//
	// The counterpart of this method in the Print System Remote Protocol is RpcEnumForms.
	// All parameters not defined below are specified in [MS-RPRN] section 3.1.4.5.5.
	//
	// Return Values: This method returns zero to indicate successful completion or a nonzero
	// Win32 error code ([MS-ERREF] section 2.2) to indicate failure. Aside from the specific
	// nonzero return values documented in section 3.1.4, the client MUST treat any nonzero
	// return value as a fatal error.
	//
	// Exceptions Thrown: This method MUST NOT throw any exceptions other than those that
	// are thrown by the underlying RPC protocol [MS-RPCE].
	//
	// This method MUST adhere to the parameter validation, processing, and response requirements
	// that are specified in [MS-RPRN] section 3.1.4.5.5.
	EnumForms(context.Context, *EnumFormsRequest) (*EnumFormsResponse, error)

	// RpcAsyncGetPrinterDriver retrieves data about a specified printer driver on a specified
	// printer.
	//
	// The counterpart of this method in the Print System Remote Protocol is RpcGetPrinterDriver2.
	// All parameters not defined below are specified in [MS-RPRN] section 3.1.4.4.6.
	//
	// Return Values: This method returns zero to indicate successful completion or a nonzero
	// Win32 error code ([MS-ERREF] section 2.2) to indicate failure. Aside from the specific
	// nonzero return values documented in section 3.1.4, the client MUST treat any nonzero
	// return value as a fatal error.
	//
	// Exceptions Thrown: This method MUST NOT throw any exceptions other than those that
	// are thrown by the underlying RPC protocol [MS-RPCE].
	//
	// This method MUST adhere to the parameter validation, processing, and response requirements
	// that are specified in [MS-RPRN] section 3.1.4.4.6.
	GetPrinterDriver(context.Context, *GetPrinterDriverRequest) (*GetPrinterDriverResponse, error)

	// RpcAsyncEnumPrinterData enumerates configuration data for a specified printer.
	//
	// The counterpart of this method in the Print System Remote Protocol is RpcEnumPrinterData.
	// All parameters not defined below are specified in [MS-RPRN] section 3.1.4.2.16.
	//
	// Return Values: This method returns zero to indicate successful completion or a nonzero
	// Win32 error code ([MS-ERREF] section 2.2) to indicate failure. Aside from the specific
	// nonzero return values documented in section 3.1.4, the client MUST treat any nonzero
	// return value as a fatal error.
	//
	// Exceptions Thrown: This method MUST NOT throw any exceptions other than those that
	// are thrown by the underlying RPC protocol [MS-RPCE].
	//
	// This method MUST adhere to the parameter validation, processing, and response requirements
	// that are specified in [MS-RPRN] section 3.1.4.2.16.
	EnumPrinterData(context.Context, *EnumPrinterDataRequest) (*EnumPrinterDataResponse, error)

	// RpcAsyncEnumPrinterDataEx enumerates all registry value names and data under the
	// key for the specified printer.
	//
	// The counterpart of this method in the Print System Remote Protocol is RpcEnumPrinterDataEx.
	// All parameters not defined below are specified in [MS-RPRN] section 3.1.4.2.20.
	//
	// Return Values: This method returns zero to indicate successful completion or a nonzero
	// Win32 error code ([MS-ERREF] section 2.2) to indicate failure. Aside from the specific
	// nonzero return values documented in section 3.1.4, the client MUST treat any nonzero
	// return value as a fatal error.
	//
	// Exceptions Thrown: This method MUST NOT throw any exceptions other than those that
	// are thrown by the underlying RPC protocol [MS-RPCE].
	//
	// This method MUST adhere to the parameter validation, processing, and response requirements
	// that are specified in [MS-RPRN] section 3.1.4.2.20.
	EnumPrinterDataEx(context.Context, *EnumPrinterDataExRequest) (*EnumPrinterDataExResponse, error)

	// RpcAsyncEnumPrinterKey enumerates the subkeys of a specified key for a specified
	// printer.
	//
	// The counterpart of this method in the Print System Remote Protocol is RpcEnumPrinterKey.
	// All parameters not defined below are specified in [MS-RPRN] section 3.1.4.2.21.
	//
	// Exceptions Thrown: This method MUST NOT throw any exceptions other than those that
	// are thrown by the underlying RPC protocol [MS-RPCE].
	//
	// This method MUST adhere to the parameter validation, processing, and response requirements
	// that are specified in [MS-RPRN] section 3.1.4.2.21.
	EnumPrinterKey(context.Context, *EnumPrinterKeyRequest) (*EnumPrinterKeyResponse, error)

	// RpcAsyncDeletePrinterData deletes a specified value from the configuration of a specified
	// printer.
	//
	// The counterpart of this method in the Print System Remote Protocol is RpcDeletePrinterData.
	// All parameters not defined below are specified in [MS-RPRN] section 3.1.4.2.17.
	//
	// Return Values: This method returns zero to indicate successful completion or a nonzero
	// Win32 error code ([MS-ERREF] section 2.2) to indicate failure. The client MUST treat
	// any nonzero return value as a fatal error.
	//
	// Exceptions Thrown: This method MUST NOT throw any exceptions other than those that
	// are thrown by the underlying RPC protocol [MS-RPCE].
	//
	// This method MUST adhere to the parameter validation, processing, and response requirements
	// that are specified in [MS-RPRN] section 3.1.4.2.17.
	DeletePrinterData(context.Context, *DeletePrinterDataRequest) (*DeletePrinterDataResponse, error)

	// RpcAsyncDeletePrinterDataEx deletes a specified value from the configuration data
	// of a specified printer, which consists of a set of named and typed values stored
	// in a hierarchy of registry keys.
	//
	// The counterpart of this method in the Print System Remote Protocol is RpcDeletePrinterDataEx.
	// All parameters not defined below are specified in [MS-RPRN] section 3.1.4.2.22.
	//
	// Return Values: This method returns zero to indicate successful completion or a nonzero
	// Win32 error code ([MS-ERREF] section 2.2) to indicate failure. The client MUST treat
	// any nonzero return value as a fatal error.
	//
	// Exceptions Thrown: This method MUST NOT throw any exceptions other than those that
	// are thrown by the underlying RPC protocol [MS-RPCE].
	//
	// This method MUST adhere to the parameter validation, processing, and response requirements
	// that are specified in [MS-RPRN] section 3.1.4.2.22.
	DeletePrinterDataEx(context.Context, *DeletePrinterDataExRequest) (*DeletePrinterDataExResponse, error)

	// RpcAsyncDeletePrinterKey deletes a specified key and all of its subkeys from the
	// configuration of a specified printer.
	//
	// The counterpart of this method in the Print System Remote Protocol is RpcDeletePrinterKey.
	// All parameters not defined below are specified in [MS-RPRN] section 3.1.4.2.23.
	//
	// Return Values: This method returns zero to indicate successful completion or a nonzero
	// Win32 error code ([MS-ERREF] section 2.2) to indicate failure. The client MUST treat
	// any nonzero return value as a fatal error.
	//
	// Exceptions Thrown: This method MUST NOT throw any exceptions other than those that
	// are thrown by the underlying RPC protocol [MS-RPCE].
	//
	// This method MUST adhere to the parameter validation, processing, and response requirements
	// that are specified in [MS-RPRN] section 3.1.4.2.23.
	DeletePrinterKey(context.Context, *DeletePrinterKeyRequest) (*DeletePrinterKeyResponse, error)

	// RpcAsyncXcvData provides the means by which a port monitor client component can communicate
	// with its server-side counterpart, the actual port monitor hosted by the server.
	//
	// The counterpart of this method in the Print System Remote Protocol is RpcXcvData.
	// All parameters not defined below are specified in [MS-RPRN] section 3.1.4.6.5.
	//
	// Return Values: This method returns zero to indicate successful completion or a nonzero
	// Win32 error code ([MS-ERREF] section 2.2) to indicate failure. Aside from the specific
	// nonzero return values documented in section 3.1.4, the client MUST treat any nonzero
	// return value as a fatal error.
	//
	// Exceptions Thrown: This method MUST NOT throw any exceptions other than those that
	// are thrown by the underlying RPC protocol [MS-RPCE].
	//
	// This method MUST adhere to the parameter validation, processing, and response requirements
	// that are specified in [MS-RPRN] section 3.1.4.6.5.
	XcvData(context.Context, *XcvDataRequest) (*XcvDataResponse, error)

	// RpcAsyncSendRecvBidiData sends and receives bidirectional data. This method is used
	// to communicate with print monitors that support such data.
	//
	// The counterpart of this method in the Print System Remote Protocol is RpcSendRecvBidiData.
	// All parameters not defined below are specified in [MS-RPRN] section 3.1.4.2.27.
	//
	// Return Values: This method returns zero to indicate successful completion or a nonzero
	// Win32 error code ([MS-ERREF] section 2.2) to indicate failure. The client MUST treat
	// any nonzero return value as a fatal error.
	//
	// Exceptions Thrown: This method MUST NOT throw any exceptions other than those that
	// are thrown by the underlying RPC protocol [MS-RPCE].
	//
	// This method MUST adhere to the parameter validation, processing, and response requirements
	// that are specified in [MS-RPRN] section 3.1.4.2.27.
	SendRecvBIDIData(context.Context, *SendRecvBIDIDataRequest) (*SendRecvBIDIDataResponse, error)

	// RpcAsyncCreatePrinterIC creates an information context for a specified printer.
	//
	// The counterpart of this method in the Print System Remote Protocol is RpcCreatePrinterIC.
	// All parameters not defined below are specified in [MS-RPRN] section 3.1.4.2.10.
	//
	// Return Values: This method returns zero to indicate successful completion or a nonzero
	// Win32 error code ([MS-ERREF] section 2.2) to indicate failure. The client MUST treat
	// any nonzero return value as a fatal error.
	//
	// Exceptions Thrown: This method MUST NOT throw any exceptions other than those that
	// are thrown by the underlying RPC protocol [MS-RPCE].
	//
	// This method MUST adhere to the parameter validation, processing, and response requirements
	// that are specified in [MS-RPRN] section 3.1.4.2.10.
	CreatePrinterIC(context.Context, *CreatePrinterICRequest) (*CreatePrinterICResponse, error)

	// RpcAsyncPlayGdiScriptOnPrinterIC returns font information for a printer connection.
	//
	// The counterpart of this method in the Print System Remote Protocol is RpcPlayGdiScriptOnPrinterIC.
	// All parameters not defined below are specified in [MS-RPRN] section 3.1.4.2.11.
	//
	// Return Values: This method returns zero to indicate successful completion or a nonzero
	// Win32 error code ([MS-ERREF] section 2.2) to indicate failure. The client MUST treat
	// any nonzero return value as a fatal error.
	//
	// Exceptions Thrown: This method MUST NOT throw any exceptions other than those that
	// are thrown by the underlying RPC protocol [MS-RPCE].
	//
	// This method MUST adhere to the parameter validation, processing, and response requirements
	// that are specified in [MS-RPRN] section 3.1.4.2.11.
	PlayGDIScriptOnPrinterIC(context.Context, *PlayGDIScriptOnPrinterICRequest) (*PlayGDIScriptOnPrinterICResponse, error)

	// RpcAsyncDeletePrinterIC deletes a printer information context.
	//
	// The counterpart of this method in the Print System Remote Protocol is RpcDeletePrinterIC.
	// All parameters not defined below are specified in [MS-RPRN] section 3.1.4.2.12.
	//
	// Return Values: This method returns zero to indicate successful completion or a nonzero
	// Win32 error code ([MS-ERREF] section 2.2) to indicate failure. The client MUST treat
	// any nonzero return value as a fatal error.
	//
	// Exceptions Thrown: This method MUST NOT throw any exceptions other than those that
	// are thrown by the underlying RPC protocol [MS-RPCE].
	//
	// This method MUST adhere to the parameter validation, processing, and response requirements
	// that are specified in [MS-RPRN] section 3.1.4.2.12.
	DeletePrinterIC(context.Context, *DeletePrinterICRequest) (*DeletePrinterICResponse, error)

	// RpcAsyncEnumPrinters enumerates available local printers, printers on a specified
	// print server, printers in a specified domain, or print providers.
	//
	// The counterpart of this method in the Print System Remote Protocol is RpcEnumPrinters.
	// All parameters not defined below are specified in [MS-RPRN] section 3.1.4.2.1.
	//
	// Return Values: This method returns zero to indicate successful completion or a nonzero
	// Win32 error code ([MS-ERREF] section 2.2) to indicate failure. Aside from the specific
	// nonzero return values documented in section 3.1.4, the client MUST treat any nonzero
	// return value as a fatal error.
	//
	// Exceptions Thrown: This method MUST NOT throw any exceptions other than those that
	// are thrown by the underlying RPC protocol [MS-RPCE].
	//
	// This method MUST adhere to the parameter validation, processing, and response requirements
	// that are specified in [MS-RPRN] section 3.1.4.2.1.
	EnumPrinters(context.Context, *EnumPrintersRequest) (*EnumPrintersResponse, error)

	// RpcAsyncAddPrinterDriver installs a specified local or a remote printer driver on
	// a specified print server, and it links the configuration, data, and driver files.
	//
	// The counterpart of this method in the Print System Remote Protocol is RpcAddPrinterDriverEx.
	// All parameters not defined below are specified in [MS-RPRN] section 3.1.4.4.8.
	//
	// Return Values: This method returns zero to indicate successful completion or a nonzero
	// Win32 error code ([MS-ERREF] section 2.2) to indicate failure. The client MUST treat
	// any nonzero return value as a fatal error.
	//
	// Exceptions Thrown: This method MUST NOT throw any exceptions other than those that
	// are thrown by the underlying RPC protocol [MS-RPCE].
	//
	// This method MUST adhere to the parameter validation, processing, and response requirements
	// that are specified in [MS-RPRN] section 3.1.4.4.8.
	AddPrinterDriver(context.Context, *AddPrinterDriverRequest) (*AddPrinterDriverResponse, error)

	// RpcAsyncEnumPrinterDrivers enumerates the printer drivers installed on a specified
	// print server.
	//
	// The counterpart of this method in the Print System Remote Protocol is RpcEnumPrinterDrivers.
	// All parameters not defined below are specified in [MS-RPRN] section 3.1.4.4.2.
	//
	// Return Values: This method returns zero to indicate successful completion or a nonzero
	// Win32 error code ([MS-ERREF] section 2.2) to indicate failure. Aside from the specific
	// nonzero return values documented in section 3.1.4, the client MUST treat any nonzero
	// return value as a fatal error.
	//
	// Exceptions Thrown: This method MUST NOT throw any exceptions other than those that
	// are thrown by the underlying RPC protocol [MS-RPCE].
	//
	// This method MUST adhere to the parameter validation, processing, and response requirements
	// that are specified in [MS-RPRN] section 3.1.4.4.2.
	EnumPrinterDrivers(context.Context, *EnumPrinterDriversRequest) (*EnumPrinterDriversResponse, error)

	// RpcAsyncGetPrinterDriverDirectory retrieves the path of the printer driver directory
	// on a specified print server.
	//
	// The counterpart of this method in the Print System Remote Protocol is RpcGetPrinterDriverDirectory.
	// All parameters not defined below are specified in [MS-RPRN] section 3.1.4.4.4.
	//
	// Return Values: This method returns zero to indicate successful completion or a nonzero
	// Win32 error code ([MS-ERREF] section 2.2) to indicate failure. Aside from the specific
	// nonzero return values documented in section 3.1.4, the client MUST treat any nonzero
	// return value as a fatal error.
	//
	// Exceptions Thrown: This method MUST NOT throw any exceptions other than those that
	// are thrown by the underlying RPC protocol [MS-RPCE].
	//
	// This method MUST adhere to the parameter validation, processing, and response requirements
	// that are specified in [MS-RPRN] section 3.1.4.4.4.
	GetPrinterDriverDirectory(context.Context, *GetPrinterDriverDirectoryRequest) (*GetPrinterDriverDirectoryResponse, error)

	// RpcAsyncDeletePrinterDriver removes the specified printer driver from the list of
	// supported drivers for a specified print server.
	//
	// The counterpart of this method in the Print System Remote Protocol is RpcDeletePrinterDriver.
	// All parameters not defined below are specified in [MS-RPRN] section 3.1.4.4.5.
	//
	// Return Values: This method returns zero to indicate successful completion or a nonzero
	// Win32 error code ([MS-ERREF] section 2.2) to indicate failure. The client MUST treat
	// any nonzero return value as a fatal error.
	//
	// Exceptions Thrown: This method MUST NOT throw any exceptions other than those that
	// are thrown by the underlying RPC protocol [MS-RPCE].
	//
	// This method MUST adhere to the parameter validation, processing, and response requirements
	// that are specified in [MS-RPRN] section 3.1.4.4.5.
	DeletePrinterDriver(context.Context, *DeletePrinterDriverRequest) (*DeletePrinterDriverResponse, error)

	// RpcAsyncDeletePrinterDriverEx removes the specified printer driver from the list
	// of supported drivers on a specified print server, and deletes the files associated
	// with the driver. This method also can delete specific versions of the driver.
	//
	// The counterpart of this method in the Print System Remote Protocol is RpcDeletePrinterDriverEx.
	// All parameters not defined below are specified in [MS-RPRN] section 3.1.4.4.7.
	//
	// Return Values: This method returns zero to indicate successful completion or a nonzero
	// Win32 error code ([MS-ERREF] section 2.2) to indicate failure. The client MUST treat
	// any nonzero return value as a fatal error.
	//
	// Exceptions Thrown: This method MUST NOT throw any exceptions other than those that
	// are thrown by the underlying RPC protocol [MS-RPCE].
	//
	// This method MUST adhere to the parameter validation, processing, and response requirements
	// that are specified in [MS-RPRN] section 3.1.4.4.7.
	DeletePrinterDriverEx(context.Context, *DeletePrinterDriverExRequest) (*DeletePrinterDriverExResponse, error)

	// RpcAsyncAddPrintProcessor installs a specified print processor on the specified server
	// and adds its name to an internal list of supported print processors.
	//
	// The counterpart of this method in the Print System Remote Protocol is RpcAddPrintProcessor.
	// All parameters not defined below are specified in [MS-RPRN] section 3.1.4.8.1.
	//
	// Return Values: This method returns zero to indicate successful completion or a nonzero
	// Win32 error code ([MS-ERREF] section 2.2) to indicate failure. The client MUST treat
	// any nonzero return value as a fatal error.
	//
	// Exceptions Thrown: This method MUST NOT throw any exceptions other than those that
	// are thrown by the underlying RPC protocol [MS-RPCE].
	//
	// This method MUST adhere to the parameter validation, processing, and response requirements
	// that are specified in [MS-RPRN] section 3.1.4.8.1.
	AddPrintProcessor(context.Context, *AddPrintProcessorRequest) (*AddPrintProcessorResponse, error)

	// RpcAsyncEnumPrintProcessors enumerates the print processors installed on a specified
	// server.
	//
	// The counterpart of this method in the Print System Remote Protocol is RpcEnumPrintProcessors.
	// All parameters not defined below are specified in [MS-RPRN] section 3.1.4.8.2.
	//
	// Return Values: This method returns zero to indicate successful completion or a nonzero
	// Win32 error code ([MS-ERREF] section 2.2) to indicate failure. Aside from the specific
	// nonzero return values documented in section 3.1.4, the client MUST treat any nonzero
	// return value as a fatal error.
	//
	// Exceptions Thrown: This method MUST NOT throw any exceptions other than those that
	// are thrown by the underlying RPC protocol [MS-RPCE].
	//
	// This method MUST adhere to the parameter validation, processing, and response requirements
	// that are specified in [MS-RPRN] section 3.1.4.8.2.
	EnumPrintProcessors(context.Context, *EnumPrintProcessorsRequest) (*EnumPrintProcessorsResponse, error)

	// RpcAsyncGetPrintProcessorDirectory retrieves the path for the print processor on
	// the specified server.
	//
	// The counterpart of this method in the Print System Remote Protocol is RpcGetPrintProcessorDirectory.
	// All parameters not defined below are specified in [MS-RPRN] section 3.1.4.8.3.
	//
	// Return Values: This method returns zero to indicate successful completion or a nonzero
	// Win32 error code ([MS-ERREF] section 2.2) to indicate failure. Aside from the specific
	// nonzero return values documented in section 3.1.4, the client MUST treat any nonzero
	// return value as a fatal error.
	//
	// Exceptions Thrown: This method MUST NOT throw any exceptions other than those that
	// are thrown by the underlying RPC protocol [MS-RPCE].
	//
	// This method MUST adhere to the parameter validation, processing, and response requirements
	// that are specified in [MS-RPRN] section 3.1.4.8.3.
	GetPrintProcessorDirectory(context.Context, *GetPrintProcessorDirectoryRequest) (*GetPrintProcessorDirectoryResponse, error)

	// RpcAsyncEnumPorts enumerates the ports that are available for printing on a specified
	// server.
	//
	// The counterpart of this method in the Print System Remote Protocol is RpcEnumPorts.
	// All parameters not defined below are specified in [MS-RPRN] section 3.1.4.6.1.
	//
	// Return Values: This method returns zero to indicate successful completion or a nonzero
	// Win32 error code ([MS-ERREF] section 2.2) to indicate failure. Aside from the specific
	// nonzero return values documented in section 3.1.4, the client MUST treat any nonzero
	// return value as a fatal error.
	//
	// Exceptions Thrown: This method MUST NOT throw any exceptions other than those that
	// are thrown by the underlying RPC protocol [MS-RPCE].
	//
	// This method MUST adhere to the parameter validation, processing, and response requirements
	// that are specified in [MS-RPRN] section 3.1.4.6.1.
	EnumPorts(context.Context, *EnumPortsRequest) (*EnumPortsResponse, error)

	// RpcAsyncEnumMonitors retrieves information about the port monitors installed on a
	// specified server.
	//
	// The counterpart of this method in the Print System Remote Protocol is RpcEnumMonitors.
	// All parameters not defined below are specified in [MS-RPRN] section 3.1.4.7.1.
	//
	// Return Values: This method returns zero to indicate successful completion or a nonzero
	// Win32 error code ([MS-ERREF] section 2.2) to indicate failure. Aside from the specific
	// nonzero return values documented in section 3.1.4, the client MUST treat any nonzero
	// return value as a fatal error.
	//
	// Exceptions Thrown: This method MUST NOT throw any exceptions other than those that
	// are thrown by the underlying RPC protocol [MS-RPCE].
	//
	// This method MUST adhere to the parameter validation, processing, and response requirements
	// that are specified in [MS-RPRN] section 3.1.4.7.1.
	EnumMonitors(context.Context, *EnumMonitorsRequest) (*EnumMonitorsResponse, error)

	// RpcAsyncAddPort adds a specified port name to the list of supported ports on a specified
	// print server.
	//
	// The counterpart of this method in the Print System Remote Protocol is RpcAddPortEx.
	// All parameters not defined below are specified in [MS-RPRN] section 3.1.4.6.3.
	//
	// Return Values: This method returns zero to indicate successful completion or a nonzero
	// Win32 error code ([MS-ERREF] section 2.2) to indicate failure. The client MUST treat
	// any nonzero return value as a fatal error.
	//
	// Exceptions Thrown: This method MUST NOT throw any exceptions other than those that
	// are thrown by the underlying RPC protocol [MS-RPCE].
	//
	// This method MUST adhere to the parameter validation, processing, and response requirements
	// that are specified in [MS-RPRN] section 3.1.4.6.3.
	AddPort(context.Context, *AddPortRequest) (*AddPortResponse, error)

	// RpcAsyncSetPort sets the status associated with a specified port on a specified print
	// server.
	//
	// The counterpart of this method in the Print System Remote Protocol is RpcSetPort.
	// All parameters not defined below are specified in [MS-RPRN] section 3.1.4.6.4.
	//
	// Return Values: This method returns zero to indicate successful completion or a nonzero
	// Win32 error code ([MS-ERREF] section 2.2) to indicate failure. The client MUST treat
	// any nonzero return value as a fatal error.
	//
	// Exceptions Thrown: This method MUST NOT throw any exceptions other than those that
	// are thrown by the underlying RPC protocol [MS-RPCE].
	//
	// This method MUST adhere to the parameter validation, processing, and response requirements
	// that are specified in [MS-RPRN] section 3.1.4.6.4.
	SetPort(context.Context, *SetPortRequest) (*SetPortResponse, error)

	// RpcAsyncAddMonitor installs a specified local port monitor, and links the configuration,
	// data, and monitor files on a specified print server.
	//
	// The counterpart of this method in the Print System Remote Protocol is RpcAddMonitor.
	// All parameters not defined below are specified in [MS-RPRN] section 3.1.4.7.2.
	//
	// Return Values: This method returns zero to indicate successful completion or a nonzero
	// Win32 error code ([MS-ERREF] section 2.2) to indicate failure. The client MUST treat
	// any nonzero return value as a fatal error.
	//
	// Exceptions Thrown: This method MUST NOT throw any exceptions other than those that
	// are thrown by the underlying RPC protocol [MS-RPCE].
	//
	// This method MUST adhere to the parameter validation, processing, and response requirements
	// that are specified in [MS-RPRN] section 3.1.4.7.2.
	AddMonitor(context.Context, *AddMonitorRequest) (*AddMonitorResponse, error)

	// RpcAsyncDeleteMonitor removes a specified port monitor from a specified print server.
	//
	// The counterpart of this method in the Print System Remote Protocol is RpcDeleteMonitor.
	// All parameters not defined below are specified in [MS-RPRN] section 3.1.4.7.3.
	//
	// Return Values: This method returns zero to indicate successful completion or a nonzero
	// Win32 error code ([MS-ERREF] section 2.2) to indicate failure. The client MUST treat
	// any nonzero return value as a fatal error.
	//
	// Exceptions Thrown: This method MUST NOT throw any exceptions other than those that
	// are thrown by the underlying RPC protocol [MS-RPCE].
	//
	// This method MUST adhere to the parameter validation, processing, and response requirements
	// that are specified in [MS-RPRN] section 3.1.4.7.3.
	DeleteMonitor(context.Context, *DeleteMonitorRequest) (*DeleteMonitorResponse, error)

	// RpcAsyncDeletePrintProcessor removes a specified print processor from a specified
	// server.
	//
	// The counterpart of this method in the Print System Remote Protocol is RpcDeletePrintProcessor.
	// All parameters not defined below are specified in [MS-RPRN] section 3.1.4.8.4.
	//
	// Return Values: This method returns zero to indicate successful completion or a nonzero
	// Win32 error code ([MS-ERREF] section 2.2) to indicate failure. The client MUST treat
	// any nonzero return value as a fatal error.
	//
	// Exceptions Thrown: This method MUST NOT throw any exceptions other than those that
	// are thrown by the underlying RPC protocol [MS-RPCE].
	//
	// This method MUST adhere to the parameter validation, processing, and response requirements
	// that are specified in [MS-RPRN] section 3.1.4.8.4.
	DeletePrintProcessor(context.Context, *DeletePrintProcessorRequest) (*DeletePrintProcessorResponse, error)

	// RpcAsyncEnumPrintProcessorDatatypes enumerates the data types that a specified print
	// processor supports.
	//
	// The counterpart of this method in the Print System Remote Protocol is RpcEnumPrintProcessorDatatypes.
	// All parameters not defined below are specified in [MS-RPRN] section 3.1.4.8.5.
	//
	// Return Values: This method returns zero to indicate successful completion or a nonzero
	// Win32 error code ([MS-ERREF] section 2.2) to indicate failure. Aside from the specific
	// nonzero return values documented in section 3.1.4, the client MUST treat any nonzero
	// return value as a fatal error.
	//
	// Exceptions Thrown: This method MUST NOT throw any exceptions other than those that
	// are thrown by the underlying RPC protocol [MS-RPCE].
	//
	// This method MUST adhere to the parameter validation, processing, and response requirements
	// that are specified in [MS-RPRN] section 3.1.4.8.5.
	EnumPrintProcessorDataTypes(context.Context, *EnumPrintProcessorDataTypesRequest) (*EnumPrintProcessorDataTypesResponse, error)

	// RpcAsyncAddPerMachineConnection persistently saves the configuration information
	// for a connection, including the print server name and the name of the print providers
	// for the specified connection.
	//
	// The counterpart of this method in the Print System Remote Protocol is RpcAddPerMachineConnection.
	// All parameters not defined below are specified in [MS-RPRN] section 3.1.4.2.24.
	//
	// Return Values: This method returns zero to indicate successful completion or a nonzero
	// Win32 error code ([MS-ERREF] section 2.2) to indicate failure. The client MUST treat
	// any nonzero return value as a fatal error.
	//
	// Exceptions Thrown: This method MUST NOT throw any exceptions other than those that
	// are thrown by the underlying RPC protocol [MS-RPCE].
	//
	// This method MUST adhere to the parameter validation, processing, and response requirements
	// that are specified in [MS-RPRN] section 3.1.4.2.24.
	AddPerMachineConnection(context.Context, *AddPerMachineConnectionRequest) (*AddPerMachineConnectionResponse, error)

	// RpcAsyncDeletePerMachineConnection deletes the stored connection configuration information
	// that corresponds to the pPrinterName parameter value.
	//
	// The counterpart of this method in the Print System Remote Protocol is RpcDeletePerMachineConnection.
	// All parameters not defined below are specified in [MS-RPRN] section 3.1.4.2.25.
	//
	// Return Values: This method returns zero to indicate successful completion or a nonzero
	// Win32 error code ([MS-ERREF] section 2.2) to indicate failure. The client MUST treat
	// any nonzero return value as a fatal error.
	//
	// Exceptions Thrown: This method MUST NOT throw any exceptions other than those that
	// are thrown by the underlying RPC protocol [MS-RPCE].
	//
	// This method MUST adhere to the parameter validation, processing, and response requirements
	// that are specified in [MS-RPRN] section 3.1.4.2.25.
	DeletePerMachineConnection(context.Context, *DeletePerMachineConnectionRequest) (*DeletePerMachineConnectionResponse, error)

	// RpcAsyncEnumPerMachineConnections enumerates each of the per-machine connections
	// into a specified buffer.
	//
	// The counterpart of this method in the Print System Remote Protocol is RpcEnumPerMachineConnections.
	// All parameters not defined below are specified in [MS-RPRN] section 3.1.4.2.26.
	//
	// Return Values: This method returns zero to indicate successful completion or a nonzero
	// Win32 error code ([MS-ERREF] section 2.2) to indicate failure. Aside from the specific
	// nonzero return values documented in section 3.1.4, the client MUST treat any nonzero
	// return value as a fatal error.
	//
	// Exceptions Thrown: This method MUST NOT throw any exceptions other than those that
	// are thrown by the underlying RPC protocol [MS-RPCE].
	//
	// This method MUST adhere to the parameter validation, processing, and response requirements
	// that are specified in [MS-RPRN] section 3.1.4.2.26.
	EnumPerMachineConnections(context.Context, *EnumPerMachineConnectionsRequest) (*EnumPerMachineConnectionsResponse, error)

	// RpcSyncRegisterForRemoteNotifications opens a notification handle by using a printer
	// handle or print server handle, to listen for remote printer change notifications.
	//
	// Return Values: This method returns either an HRESULT success value ([MS-ERREF] section
	// 2.1) to indicate successful completion, or an HRESULT error value to indicate failure.
	//
	// Exceptions Thrown: This method MUST NOT throw any exceptions other than those that
	// are thrown by the underlying RPC protocol [MS-RPCE].
	SyncRegisterForRemoteNotifications(context.Context, *SyncRegisterForRemoteNotificationsRequest) (*SyncRegisterForRemoteNotificationsResponse, error)

	// RpcSyncUnRegisterForRemoteNotifications closes a notification handle opened by calling
	// RpcSyncRegisterForRemoteNotifications (section 3.1.4.9.1).
	//
	// Return Values: This method returns either an HRESULT success value ([MS-ERREF] section
	// 2.1) to indicate successful completion or an HRESULT error value to indicate failure.
	//
	// Exceptions Thrown: This method MUST NOT throw any exceptions other than those that
	// are thrown by the underlying RPC protocol [MS-RPCE].
	SyncUnregisterForRemoteNotifications(context.Context, *SyncUnregisterForRemoteNotificationsRequest) (*SyncUnregisterForRemoteNotificationsResponse, error)

	// RpcSyncRefreshRemoteNotifications gets notification information for all requested
	// members. This SHOULD be called by a client if the "RemoteNotifyData Flags" key in
	// the RpcPrintPropertiesCollection instance (section 2.2.4), which was returned as
	// part of the notification from calling RpcAsyncGetRemoteNotifications (section 3.1.4.9.4),
	// has the PRINTER_NOTIFY_INFO_DISCARDED bit set ([MS-RPRN] section 2.2.3.2).
	//
	// Return Values: This method returns either an HRESULT success value ([MS-ERREF] section
	// 2.1) to indicate successful completion or an HRESULT error value to indicate failure.
	//
	// Exceptions Thrown: This method MUST NOT throw any exceptions other than those that
	// are thrown by the underlying RPC protocol [MS-RPCE].
	SyncRefreshRemoteNotifications(context.Context, *SyncRefreshRemoteNotificationsRequest) (*SyncRefreshRemoteNotificationsResponse, error)

	// A print client uses RpcAsyncGetRemoteNotifications to poll the print server for specified
	// change notifications. A call to this method is suspended until the server has a new
	// change notification for the client. Subsequently, the client calls this method again
	// to poll for additional notifications from the server.
	//
	// Return Values: This method returns either an HRESULT success value ([MS-ERREF] section
	// 2.1) to indicate successful completion or an HRESULT error value to indicate failure.
	//
	// Exceptions Thrown: This method MUST NOT throw any exceptions other than those that
	// are thrown by the underlying RPC protocol [MS-RPCE].
	GetRemoteNotifications(context.Context, *GetRemoteNotificationsRequest) (*GetRemoteNotificationsResponse, error)

	// RpcAsyncInstallPrinterDriverFromPackage installs a printer driver from a driver package.
	//
	// Return Values: This method returns either an HRESULT success value ([MS-ERREF] section
	// 2.1) or an HRESULT error value to indicate failure.
	//
	// Exceptions Thrown: This method MUST NOT throw any exceptions other than those that
	// are thrown by the underlying RPC protocol [MS-RPCE].
	InstallPrinterDriverFromPackage(context.Context, *InstallPrinterDriverFromPackageRequest) (*InstallPrinterDriverFromPackageResponse, error)

	// RpcAsyncUploadPrinterDriverPackage uploads a driver package so it can be installed
	// with RpcAsyncInstallPrinterDriverFromPackage (section 3.1.4.2.7).
	//
	// Return Values: This method returns either an HRESULT success value ([MS-ERREF] section
	// 2.1) or an HRESULT error value to indicate failure.
	//
	// Exceptions Thrown: This method MUST NOT throw any exceptions other than those that
	// are thrown by the underlying RPC protocol [MS-RPCE].
	UploadPrinterDriverPackage(context.Context, *UploadPrinterDriverPackageRequest) (*UploadPrinterDriverPackageResponse, error)

	// RpcAsyncGetCorePrinterDrivers gets the GUID, versions, and publish dates of the specified
	// core printer drivers, and the paths to their packages.
	//
	// The counterpart of this method in the Print System Remote Protocol is RpcGetCorePrinterDrivers
	// ([MS-RPRN] section 3.1.4.4.9).
	//
	// Return Values: This method returns either an HRESULT success value ([MS-ERREF] section
	// 2.1) or an HRESULT error value to indicate failure.
	//
	// Exceptions Thrown: This method MUST NOT throw any exceptions other than those that
	// are thrown by the underlying RPC protocol [MS-RPCE].
	//
	// This method MUST adhere to the parameter validation, processing, and response requirements
	// that are specified in [MS-RPRN] section 3.1.4.4.9.
	GetCorePrinterDrivers(context.Context, *GetCorePrinterDriversRequest) (*GetCorePrinterDriversResponse, error)

	// RpcAsyncCorePrinterDriverInstalled determines if a specific core printer driver is
	// installed.
	//
	// Return Values: This method returns either an HRESULT success value ([MS-ERREF] section
	// 2.1) or an HRESULT error value to indicate failure.
	//
	// Exceptions Thrown: This method MUST NOT throw any exceptions other than those that
	// are thrown by the underlying RPC protocol [MS-RPCE].
	CorePrinterDriverInstalled(context.Context, *CorePrinterDriverInstalledRequest) (*CorePrinterDriverInstalledResponse, error)

	// RpcAsyncGetPrinterDriverPackagePath gets the path to the specified printer driver
	// package.
	//
	// The counterpart of this method in the Print System Remote Protocol is RpcGetPrinterDriverPackagePath,
	// [MS-RPRN] section 3.1.4.4.10.
	//
	// Return Values: This method returns either an HRESULT success value ([MS-ERREF] section
	// 2.1) or an HRESULT error value to indicate failure.
	//
	// Exceptions Thrown: This method MUST NOT throw any exceptions other than those that
	// are thrown by the underlying RPC protocol [MS-RPCE].
	GetPrinterDriverPackagePath(context.Context, *GetPrinterDriverPackagePathRequest) (*GetPrinterDriverPackagePathResponse, error)

	// RpcAsyncDeletePrinterDriverPackage deletes a specified printer driver package.
	//
	// Return Values: This method returns either an HRESULT success value ([MS-ERREF] section
	// 2.1) or an HRESULT error value to indicate failure.
	//
	// Exceptions Thrown: This method MUST NOT throw any exceptions other than those that
	// are thrown by the underlying RPC protocol [MS-RPCE].
	DeletePrinterDriverPackage(context.Context, *DeletePrinterDriverPackageRequest) (*DeletePrinterDriverPackageResponse, error)

	// RpcAsyncReadPrinter retrieves data from the specified job object.
	//
	// The counterpart of this method in the Print System Remote Protocol is RpcReadPrinter.
	// All parameters not defined below are specified in [MS-RPRN] section 3.1.4.9.6.
	//
	// Return Values: This method returns zero to indicate successful completion or a nonzero
	// Win32 error code ([MS-ERREF] section 2.2) to indicate failure. The client MUST treat
	// any nonzero return value as a fatal error.
	//
	// Exceptions Thrown: This method MUST NOT throw any exceptions other than those that
	// are thrown by the underlying RPC protocol [MS-RPCE].
	//
	// This method MUST adhere to the parameter validation, processing, and response requirements
	// that are specified in [MS-RPRN] section 3.1.4.9.6.
	ReadPrinter(context.Context, *ReadPrinterRequest) (*ReadPrinterResponse, error)

	// RpcAsyncResetPrinter resets the data type and device mode (For more information,
	// see [DEVMODE]) values to use for printing documents submitted by the RpcAsyncStartDocPrinter
	// method (section 3.1.4.8.1).
	//
	// The counterpart of this method in the Print System Remote Protocol is RpcResetPrinter.
	// All parameters not defined below are specified in [MS-RPRN] section 3.1.4.2.13.
	//
	// Return Values: This method returns zero to indicate successful completion or a nonzero
	// Win32 error code ([MS-ERREF] section 2.2) to indicate failure. The client MUST treat
	// any nonzero return value as a fatal error.
	//
	// Exceptions Thrown: This method MUST NOT throw any exceptions other than those that
	// are thrown by the underlying RPC protocol [MS-RPCE].
	//
	// This method MUST adhere to the parameter validation, processing, and response requirements
	// that are specified in [MS-RPRN] section 3.1.4.2.13.
	ResetPrinter(context.Context, *ResetPrinterRequest) (*ResetPrinterResponse, error)

	// RpcAsyncGetJobNamedPropertyValue retrieves the current value of the specified Job
	// Named Property (section 3.1.1).<31>
	//
	// The counterpart of this method in the Print System Remote Protocol ([MS-RPRN]) is
	// RpcGetJobNamedPropertyValue. All parameters not defined below are specified in [MS-RPRN]
	// section 3.1.4.12.1.
	//
	// Return Values: This method returns zero to indicate successful completion or a nonzero
	// Win32 error code ([MS-ERREF] section 2.2) to indicate failure. The client MUST treat
	// any nonzero return value as a fatal error.
	//
	// Exceptions Thrown: This method MUST NOT throw any exceptions other than those that
	// are thrown by the underlying RPC protocol specified in [MS-RPCE].
	//
	// This method MUST adhere to the parameter validation, processing, and response requirements
	// that are specified in [MS-RPRN] section 3.1.4.12.1.
	GetJobNamedPropertyValue(context.Context, *GetJobNamedPropertyValueRequest) (*GetJobNamedPropertyValueResponse, error)

	// RpcAsyncSetJobNamedProperty creates a new Job Named Property (section 3.1.1), or
	// changes the value of an existing Job Named Property for the specified print job.<32>
	//
	// The counterpart of this method in the Print System Remote Protocol is RpcSetJobNamedProperty.
	// All parameters not defined below are specified in [MS-RPRN] section 3.1.4.12.2.
	//
	// Return Values: This method returns zero to indicate successful completion or a nonzero
	// Win32 error code ([MS-ERREF] section 2.2) to indicate failure. The client MUST treat
	// any nonzero return value as a fatal error.
	//
	// Exceptions Thrown: This method MUST NOT throw any exceptions other than those that
	// are thrown by the underlying RPC protocol specified in [MS-RPCE].
	//
	// This method MUST adhere to the parameter validation, processing, and response requirements
	// that are specified in [MS-RPRN] section 3.1.4.12.2.
	SetJobNamedProperty(context.Context, *SetJobNamedPropertyRequest) (*SetJobNamedPropertyResponse, error)

	// RpcAsyncDeleteJobNamedProperty deletes an existing Job Named Property (section 3.1.1)
	// for the specified print job.<33>
	//
	// The counterpart of this method in the Print System Remote Protocol is RpcDeleteJobNamedProperty.
	// All parameters not defined below are specified in [MS-RPRN] section 3.1.4.12.3.
	//
	// Return Values: This method returns zero to indicate successful completion or a nonzero
	// Win32 error code ([MS-ERREF] section 2.2) to indicate failure. The client MUST treat
	// any nonzero return value as a fatal error.
	//
	// Exceptions Thrown: This method MUST NOT throw any exceptions other than those that
	// are thrown by the underlying RPC protocol specified in [MS-RPCE].
	//
	// This method MUST adhere to the parameter validation, processing, and response requirements
	// that are specified in [MS-RPRN] section 3.1.4.12.3.
	DeleteJobNamedProperty(context.Context, *DeleteJobNamedPropertyRequest) (*DeleteJobNamedPropertyResponse, error)

	// RpcAsyncEnumJobNamedProperties enumerates the Job Named Property (section 3.1.1)
	// for the specified print job.<34>
	//
	// The counterpart of this method in the Print System Remote Protocol is RpcEnumJobNamedProperties
	// (section 3.1.4.12.4). All parameters not defined below are specified in [MS-RPRN]
	// section 3.1.4.12.4.
	//
	// Return Values: This method returns zero to indicate successful completion or a nonzero
	// Win32 error code ([MS-ERREF] section 2.2) to indicate failure. The client MUST treat
	// any nonzero return value as a fatal error.
	//
	// Exceptions Thrown: This method MUST NOT throw any exceptions other than those that
	// are thrown by the underlying RPC protocol [MS-RPCE].
	//
	// This method MUST adhere to the parameter validation, processing, and response requirements
	// that are specified in [MS-RPRN] section 3.1.4.12.4.
	EnumJobNamedProperties(context.Context, *EnumJobNamedPropertiesRequest) (*EnumJobNamedPropertiesResponse, error)

	// RpcAsyncLogJobInfoForBranchOffice processes one or more Branch Office Print Remote
	// Log Entries (section 3.1.1).<35>
	//
	// The counterpart of this method in the Print System Remote Protocol is RpcLogJobInfoForBranchOffice.
	// All parameters not defined below are specified in [MS-RPRN] section 3.1.4.13.1.
	//
	// Return Values: This method returns zero to indicate successful completion or a nonzero
	// Win32 error code ([MS-ERREF] section 2.2) to indicate failure. The client MUST treat
	// any nonzero return value as a fatal error.
	//
	// Exceptions Thrown: This method MUST NOT throw any exceptions other than those that
	// are thrown by the underlying RPC protocol specified in [MS-RPCE].
	//
	// This method MUST adhere to the parameter validation, processing, and response requirements
	// that are specified in [MS-RPRN] section 3.1.4.13.1.
	LogJobInfoForBranchOffice(context.Context, *LogJobInfoForBranchOfficeRequest) (*LogJobInfoForBranchOfficeResponse, error)
}

func RegisterRemoteWinspoolServer(conn dcerpc.Conn, o RemoteWinspoolServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewRemoteWinspoolServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(RemoteWinspoolSyntaxV1_0))...)
}

func NewRemoteWinspoolServerHandle(o RemoteWinspoolServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return RemoteWinspoolServerHandle(ctx, o, opNum, r)
	}
}

func RemoteWinspoolServerHandle(ctx context.Context, o RemoteWinspoolServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	switch opNum {
	case 0: // RpcAsyncOpenPrinter
		op := &xxx_OpenPrinterOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &OpenPrinterRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.OpenPrinter(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 1: // RpcAsyncAddPrinter
		op := &xxx_AddPrinterOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &AddPrinterRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.AddPrinter(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 2: // RpcAsyncSetJob
		op := &xxx_SetJobOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetJobRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetJob(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 3: // RpcAsyncGetJob
		op := &xxx_GetJobOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetJobRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetJob(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 4: // RpcAsyncEnumJobs
		op := &xxx_EnumJobsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &EnumJobsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.EnumJobs(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 5: // RpcAsyncAddJob
		op := &xxx_AddJobOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &AddJobRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.AddJob(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 6: // RpcAsyncScheduleJob
		op := &xxx_ScheduleJobOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ScheduleJobRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.ScheduleJob(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 7: // RpcAsyncDeletePrinter
		op := &xxx_DeletePrinterOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &DeletePrinterRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.DeletePrinter(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 8: // RpcAsyncSetPrinter
		op := &xxx_SetPrinterOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetPrinterRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetPrinter(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 9: // RpcAsyncGetPrinter
		op := &xxx_GetPrinterOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetPrinterRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetPrinter(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 10: // RpcAsyncStartDocPrinter
		op := &xxx_StartDocPrinterOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &StartDocPrinterRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.StartDocPrinter(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 11: // RpcAsyncStartPagePrinter
		op := &xxx_StartPagePrinterOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &StartPagePrinterRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.StartPagePrinter(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 12: // RpcAsyncWritePrinter
		op := &xxx_WritePrinterOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &WritePrinterRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.WritePrinter(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 13: // RpcAsyncEndPagePrinter
		op := &xxx_EndPagePrinterOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &EndPagePrinterRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.EndPagePrinter(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 14: // RpcAsyncEndDocPrinter
		op := &xxx_EndDocPrinterOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &EndDocPrinterRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.EndDocPrinter(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 15: // RpcAsyncAbortPrinter
		op := &xxx_AbortPrinterOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &AbortPrinterRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.AbortPrinter(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 16: // RpcAsyncGetPrinterData
		op := &xxx_GetPrinterDataOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetPrinterDataRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetPrinterData(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 17: // RpcAsyncGetPrinterDataEx
		op := &xxx_GetPrinterDataExOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetPrinterDataExRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetPrinterDataEx(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 18: // RpcAsyncSetPrinterData
		op := &xxx_SetPrinterDataOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetPrinterDataRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetPrinterData(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 19: // RpcAsyncSetPrinterDataEx
		op := &xxx_SetPrinterDataExOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetPrinterDataExRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetPrinterDataEx(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 20: // RpcAsyncClosePrinter
		op := &xxx_ClosePrinterOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ClosePrinterRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.ClosePrinter(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 21: // RpcAsyncAddForm
		op := &xxx_AddFormOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &AddFormRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.AddForm(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 22: // RpcAsyncDeleteForm
		op := &xxx_DeleteFormOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &DeleteFormRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.DeleteForm(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 23: // RpcAsyncGetForm
		op := &xxx_GetFormOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetFormRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetForm(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 24: // RpcAsyncSetForm
		op := &xxx_SetFormOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetFormRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetForm(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 25: // RpcAsyncEnumForms
		op := &xxx_EnumFormsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &EnumFormsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.EnumForms(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 26: // RpcAsyncGetPrinterDriver
		op := &xxx_GetPrinterDriverOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetPrinterDriverRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetPrinterDriver(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 27: // RpcAsyncEnumPrinterData
		op := &xxx_EnumPrinterDataOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &EnumPrinterDataRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.EnumPrinterData(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 28: // RpcAsyncEnumPrinterDataEx
		op := &xxx_EnumPrinterDataExOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &EnumPrinterDataExRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.EnumPrinterDataEx(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 29: // RpcAsyncEnumPrinterKey
		op := &xxx_EnumPrinterKeyOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &EnumPrinterKeyRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.EnumPrinterKey(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 30: // RpcAsyncDeletePrinterData
		op := &xxx_DeletePrinterDataOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &DeletePrinterDataRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.DeletePrinterData(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 31: // RpcAsyncDeletePrinterDataEx
		op := &xxx_DeletePrinterDataExOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &DeletePrinterDataExRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.DeletePrinterDataEx(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 32: // RpcAsyncDeletePrinterKey
		op := &xxx_DeletePrinterKeyOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &DeletePrinterKeyRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.DeletePrinterKey(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 33: // RpcAsyncXcvData
		op := &xxx_XcvDataOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &XcvDataRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.XcvData(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 34: // RpcAsyncSendRecvBidiData
		op := &xxx_SendRecvBIDIDataOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SendRecvBIDIDataRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SendRecvBIDIData(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 35: // RpcAsyncCreatePrinterIC
		op := &xxx_CreatePrinterICOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &CreatePrinterICRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.CreatePrinterIC(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 36: // RpcAsyncPlayGdiScriptOnPrinterIC
		op := &xxx_PlayGDIScriptOnPrinterICOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &PlayGDIScriptOnPrinterICRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.PlayGDIScriptOnPrinterIC(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 37: // RpcAsyncDeletePrinterIC
		op := &xxx_DeletePrinterICOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &DeletePrinterICRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.DeletePrinterIC(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 38: // RpcAsyncEnumPrinters
		op := &xxx_EnumPrintersOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &EnumPrintersRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.EnumPrinters(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 39: // RpcAsyncAddPrinterDriver
		op := &xxx_AddPrinterDriverOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &AddPrinterDriverRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.AddPrinterDriver(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 40: // RpcAsyncEnumPrinterDrivers
		op := &xxx_EnumPrinterDriversOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &EnumPrinterDriversRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.EnumPrinterDrivers(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 41: // RpcAsyncGetPrinterDriverDirectory
		op := &xxx_GetPrinterDriverDirectoryOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetPrinterDriverDirectoryRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetPrinterDriverDirectory(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 42: // RpcAsyncDeletePrinterDriver
		op := &xxx_DeletePrinterDriverOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &DeletePrinterDriverRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.DeletePrinterDriver(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 43: // RpcAsyncDeletePrinterDriverEx
		op := &xxx_DeletePrinterDriverExOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &DeletePrinterDriverExRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.DeletePrinterDriverEx(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 44: // RpcAsyncAddPrintProcessor
		op := &xxx_AddPrintProcessorOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &AddPrintProcessorRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.AddPrintProcessor(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 45: // RpcAsyncEnumPrintProcessors
		op := &xxx_EnumPrintProcessorsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &EnumPrintProcessorsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.EnumPrintProcessors(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 46: // RpcAsyncGetPrintProcessorDirectory
		op := &xxx_GetPrintProcessorDirectoryOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetPrintProcessorDirectoryRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetPrintProcessorDirectory(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 47: // RpcAsyncEnumPorts
		op := &xxx_EnumPortsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &EnumPortsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.EnumPorts(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 48: // RpcAsyncEnumMonitors
		op := &xxx_EnumMonitorsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &EnumMonitorsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.EnumMonitors(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 49: // RpcAsyncAddPort
		op := &xxx_AddPortOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &AddPortRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.AddPort(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 50: // RpcAsyncSetPort
		op := &xxx_SetPortOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetPortRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetPort(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 51: // RpcAsyncAddMonitor
		op := &xxx_AddMonitorOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &AddMonitorRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.AddMonitor(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 52: // RpcAsyncDeleteMonitor
		op := &xxx_DeleteMonitorOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &DeleteMonitorRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.DeleteMonitor(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 53: // RpcAsyncDeletePrintProcessor
		op := &xxx_DeletePrintProcessorOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &DeletePrintProcessorRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.DeletePrintProcessor(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 54: // RpcAsyncEnumPrintProcessorDatatypes
		op := &xxx_EnumPrintProcessorDataTypesOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &EnumPrintProcessorDataTypesRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.EnumPrintProcessorDataTypes(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 55: // RpcAsyncAddPerMachineConnection
		op := &xxx_AddPerMachineConnectionOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &AddPerMachineConnectionRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.AddPerMachineConnection(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 56: // RpcAsyncDeletePerMachineConnection
		op := &xxx_DeletePerMachineConnectionOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &DeletePerMachineConnectionRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.DeletePerMachineConnection(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 57: // RpcAsyncEnumPerMachineConnections
		op := &xxx_EnumPerMachineConnectionsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &EnumPerMachineConnectionsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.EnumPerMachineConnections(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 58: // RpcSyncRegisterForRemoteNotifications
		op := &xxx_SyncRegisterForRemoteNotificationsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SyncRegisterForRemoteNotificationsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SyncRegisterForRemoteNotifications(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 59: // RpcSyncUnRegisterForRemoteNotifications
		op := &xxx_SyncUnregisterForRemoteNotificationsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SyncUnregisterForRemoteNotificationsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SyncUnregisterForRemoteNotifications(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 60: // RpcSyncRefreshRemoteNotifications
		op := &xxx_SyncRefreshRemoteNotificationsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SyncRefreshRemoteNotificationsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SyncRefreshRemoteNotifications(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 61: // RpcAsyncGetRemoteNotifications
		op := &xxx_GetRemoteNotificationsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetRemoteNotificationsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetRemoteNotifications(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 62: // RpcAsyncInstallPrinterDriverFromPackage
		op := &xxx_InstallPrinterDriverFromPackageOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &InstallPrinterDriverFromPackageRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.InstallPrinterDriverFromPackage(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 63: // RpcAsyncUploadPrinterDriverPackage
		op := &xxx_UploadPrinterDriverPackageOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &UploadPrinterDriverPackageRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.UploadPrinterDriverPackage(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 64: // RpcAsyncGetCorePrinterDrivers
		op := &xxx_GetCorePrinterDriversOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetCorePrinterDriversRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetCorePrinterDrivers(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 65: // RpcAsyncCorePrinterDriverInstalled
		op := &xxx_CorePrinterDriverInstalledOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &CorePrinterDriverInstalledRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.CorePrinterDriverInstalled(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 66: // RpcAsyncGetPrinterDriverPackagePath
		op := &xxx_GetPrinterDriverPackagePathOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetPrinterDriverPackagePathRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetPrinterDriverPackagePath(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 67: // RpcAsyncDeletePrinterDriverPackage
		op := &xxx_DeletePrinterDriverPackageOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &DeletePrinterDriverPackageRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.DeletePrinterDriverPackage(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 68: // RpcAsyncReadPrinter
		op := &xxx_ReadPrinterOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ReadPrinterRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.ReadPrinter(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 69: // RpcAsyncResetPrinter
		op := &xxx_ResetPrinterOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ResetPrinterRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.ResetPrinter(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 70: // RpcAsyncGetJobNamedPropertyValue
		op := &xxx_GetJobNamedPropertyValueOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetJobNamedPropertyValueRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetJobNamedPropertyValue(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 71: // RpcAsyncSetJobNamedProperty
		op := &xxx_SetJobNamedPropertyOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetJobNamedPropertyRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetJobNamedProperty(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 72: // RpcAsyncDeleteJobNamedProperty
		op := &xxx_DeleteJobNamedPropertyOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &DeleteJobNamedPropertyRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.DeleteJobNamedProperty(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 73: // RpcAsyncEnumJobNamedProperties
		op := &xxx_EnumJobNamedPropertiesOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &EnumJobNamedPropertiesRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.EnumJobNamedProperties(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 74: // RpcAsyncLogJobInfoForBranchOffice
		op := &xxx_LogJobInfoForBranchOfficeOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &LogJobInfoForBranchOfficeRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.LogJobInfoForBranchOffice(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IRemoteWinspool
type UnimplementedRemoteWinspoolServer struct {
}

func (UnimplementedRemoteWinspoolServer) OpenPrinter(context.Context, *OpenPrinterRequest) (*OpenPrinterResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRemoteWinspoolServer) AddPrinter(context.Context, *AddPrinterRequest) (*AddPrinterResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRemoteWinspoolServer) SetJob(context.Context, *SetJobRequest) (*SetJobResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRemoteWinspoolServer) GetJob(context.Context, *GetJobRequest) (*GetJobResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRemoteWinspoolServer) EnumJobs(context.Context, *EnumJobsRequest) (*EnumJobsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRemoteWinspoolServer) AddJob(context.Context, *AddJobRequest) (*AddJobResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRemoteWinspoolServer) ScheduleJob(context.Context, *ScheduleJobRequest) (*ScheduleJobResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRemoteWinspoolServer) DeletePrinter(context.Context, *DeletePrinterRequest) (*DeletePrinterResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRemoteWinspoolServer) SetPrinter(context.Context, *SetPrinterRequest) (*SetPrinterResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRemoteWinspoolServer) GetPrinter(context.Context, *GetPrinterRequest) (*GetPrinterResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRemoteWinspoolServer) StartDocPrinter(context.Context, *StartDocPrinterRequest) (*StartDocPrinterResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRemoteWinspoolServer) StartPagePrinter(context.Context, *StartPagePrinterRequest) (*StartPagePrinterResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRemoteWinspoolServer) WritePrinter(context.Context, *WritePrinterRequest) (*WritePrinterResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRemoteWinspoolServer) EndPagePrinter(context.Context, *EndPagePrinterRequest) (*EndPagePrinterResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRemoteWinspoolServer) EndDocPrinter(context.Context, *EndDocPrinterRequest) (*EndDocPrinterResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRemoteWinspoolServer) AbortPrinter(context.Context, *AbortPrinterRequest) (*AbortPrinterResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRemoteWinspoolServer) GetPrinterData(context.Context, *GetPrinterDataRequest) (*GetPrinterDataResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRemoteWinspoolServer) GetPrinterDataEx(context.Context, *GetPrinterDataExRequest) (*GetPrinterDataExResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRemoteWinspoolServer) SetPrinterData(context.Context, *SetPrinterDataRequest) (*SetPrinterDataResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRemoteWinspoolServer) SetPrinterDataEx(context.Context, *SetPrinterDataExRequest) (*SetPrinterDataExResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRemoteWinspoolServer) ClosePrinter(context.Context, *ClosePrinterRequest) (*ClosePrinterResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRemoteWinspoolServer) AddForm(context.Context, *AddFormRequest) (*AddFormResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRemoteWinspoolServer) DeleteForm(context.Context, *DeleteFormRequest) (*DeleteFormResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRemoteWinspoolServer) GetForm(context.Context, *GetFormRequest) (*GetFormResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRemoteWinspoolServer) SetForm(context.Context, *SetFormRequest) (*SetFormResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRemoteWinspoolServer) EnumForms(context.Context, *EnumFormsRequest) (*EnumFormsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRemoteWinspoolServer) GetPrinterDriver(context.Context, *GetPrinterDriverRequest) (*GetPrinterDriverResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRemoteWinspoolServer) EnumPrinterData(context.Context, *EnumPrinterDataRequest) (*EnumPrinterDataResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRemoteWinspoolServer) EnumPrinterDataEx(context.Context, *EnumPrinterDataExRequest) (*EnumPrinterDataExResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRemoteWinspoolServer) EnumPrinterKey(context.Context, *EnumPrinterKeyRequest) (*EnumPrinterKeyResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRemoteWinspoolServer) DeletePrinterData(context.Context, *DeletePrinterDataRequest) (*DeletePrinterDataResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRemoteWinspoolServer) DeletePrinterDataEx(context.Context, *DeletePrinterDataExRequest) (*DeletePrinterDataExResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRemoteWinspoolServer) DeletePrinterKey(context.Context, *DeletePrinterKeyRequest) (*DeletePrinterKeyResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRemoteWinspoolServer) XcvData(context.Context, *XcvDataRequest) (*XcvDataResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRemoteWinspoolServer) SendRecvBIDIData(context.Context, *SendRecvBIDIDataRequest) (*SendRecvBIDIDataResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRemoteWinspoolServer) CreatePrinterIC(context.Context, *CreatePrinterICRequest) (*CreatePrinterICResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRemoteWinspoolServer) PlayGDIScriptOnPrinterIC(context.Context, *PlayGDIScriptOnPrinterICRequest) (*PlayGDIScriptOnPrinterICResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRemoteWinspoolServer) DeletePrinterIC(context.Context, *DeletePrinterICRequest) (*DeletePrinterICResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRemoteWinspoolServer) EnumPrinters(context.Context, *EnumPrintersRequest) (*EnumPrintersResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRemoteWinspoolServer) AddPrinterDriver(context.Context, *AddPrinterDriverRequest) (*AddPrinterDriverResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRemoteWinspoolServer) EnumPrinterDrivers(context.Context, *EnumPrinterDriversRequest) (*EnumPrinterDriversResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRemoteWinspoolServer) GetPrinterDriverDirectory(context.Context, *GetPrinterDriverDirectoryRequest) (*GetPrinterDriverDirectoryResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRemoteWinspoolServer) DeletePrinterDriver(context.Context, *DeletePrinterDriverRequest) (*DeletePrinterDriverResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRemoteWinspoolServer) DeletePrinterDriverEx(context.Context, *DeletePrinterDriverExRequest) (*DeletePrinterDriverExResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRemoteWinspoolServer) AddPrintProcessor(context.Context, *AddPrintProcessorRequest) (*AddPrintProcessorResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRemoteWinspoolServer) EnumPrintProcessors(context.Context, *EnumPrintProcessorsRequest) (*EnumPrintProcessorsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRemoteWinspoolServer) GetPrintProcessorDirectory(context.Context, *GetPrintProcessorDirectoryRequest) (*GetPrintProcessorDirectoryResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRemoteWinspoolServer) EnumPorts(context.Context, *EnumPortsRequest) (*EnumPortsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRemoteWinspoolServer) EnumMonitors(context.Context, *EnumMonitorsRequest) (*EnumMonitorsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRemoteWinspoolServer) AddPort(context.Context, *AddPortRequest) (*AddPortResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRemoteWinspoolServer) SetPort(context.Context, *SetPortRequest) (*SetPortResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRemoteWinspoolServer) AddMonitor(context.Context, *AddMonitorRequest) (*AddMonitorResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRemoteWinspoolServer) DeleteMonitor(context.Context, *DeleteMonitorRequest) (*DeleteMonitorResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRemoteWinspoolServer) DeletePrintProcessor(context.Context, *DeletePrintProcessorRequest) (*DeletePrintProcessorResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRemoteWinspoolServer) EnumPrintProcessorDataTypes(context.Context, *EnumPrintProcessorDataTypesRequest) (*EnumPrintProcessorDataTypesResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRemoteWinspoolServer) AddPerMachineConnection(context.Context, *AddPerMachineConnectionRequest) (*AddPerMachineConnectionResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRemoteWinspoolServer) DeletePerMachineConnection(context.Context, *DeletePerMachineConnectionRequest) (*DeletePerMachineConnectionResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRemoteWinspoolServer) EnumPerMachineConnections(context.Context, *EnumPerMachineConnectionsRequest) (*EnumPerMachineConnectionsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRemoteWinspoolServer) SyncRegisterForRemoteNotifications(context.Context, *SyncRegisterForRemoteNotificationsRequest) (*SyncRegisterForRemoteNotificationsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRemoteWinspoolServer) SyncUnregisterForRemoteNotifications(context.Context, *SyncUnregisterForRemoteNotificationsRequest) (*SyncUnregisterForRemoteNotificationsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRemoteWinspoolServer) SyncRefreshRemoteNotifications(context.Context, *SyncRefreshRemoteNotificationsRequest) (*SyncRefreshRemoteNotificationsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRemoteWinspoolServer) GetRemoteNotifications(context.Context, *GetRemoteNotificationsRequest) (*GetRemoteNotificationsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRemoteWinspoolServer) InstallPrinterDriverFromPackage(context.Context, *InstallPrinterDriverFromPackageRequest) (*InstallPrinterDriverFromPackageResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRemoteWinspoolServer) UploadPrinterDriverPackage(context.Context, *UploadPrinterDriverPackageRequest) (*UploadPrinterDriverPackageResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRemoteWinspoolServer) GetCorePrinterDrivers(context.Context, *GetCorePrinterDriversRequest) (*GetCorePrinterDriversResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRemoteWinspoolServer) CorePrinterDriverInstalled(context.Context, *CorePrinterDriverInstalledRequest) (*CorePrinterDriverInstalledResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRemoteWinspoolServer) GetPrinterDriverPackagePath(context.Context, *GetPrinterDriverPackagePathRequest) (*GetPrinterDriverPackagePathResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRemoteWinspoolServer) DeletePrinterDriverPackage(context.Context, *DeletePrinterDriverPackageRequest) (*DeletePrinterDriverPackageResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRemoteWinspoolServer) ReadPrinter(context.Context, *ReadPrinterRequest) (*ReadPrinterResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRemoteWinspoolServer) ResetPrinter(context.Context, *ResetPrinterRequest) (*ResetPrinterResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRemoteWinspoolServer) GetJobNamedPropertyValue(context.Context, *GetJobNamedPropertyValueRequest) (*GetJobNamedPropertyValueResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRemoteWinspoolServer) SetJobNamedProperty(context.Context, *SetJobNamedPropertyRequest) (*SetJobNamedPropertyResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRemoteWinspoolServer) DeleteJobNamedProperty(context.Context, *DeleteJobNamedPropertyRequest) (*DeleteJobNamedPropertyResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRemoteWinspoolServer) EnumJobNamedProperties(context.Context, *EnumJobNamedPropertiesRequest) (*EnumJobNamedPropertiesResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRemoteWinspoolServer) LogJobInfoForBranchOffice(context.Context, *LogJobInfoForBranchOfficeRequest) (*LogJobInfoForBranchOfficeResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ RemoteWinspoolServer = (*UnimplementedRemoteWinspoolServer)(nil)
