package iremotewinspool

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
		in := &OpenPrinterRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.OpenPrinter(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 1: // RpcAsyncAddPrinter
		in := &AddPrinterRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.AddPrinter(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 2: // RpcAsyncSetJob
		in := &SetJobRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetJob(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 3: // RpcAsyncGetJob
		in := &GetJobRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetJob(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 4: // RpcAsyncEnumJobs
		in := &EnumJobsRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.EnumJobs(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 5: // RpcAsyncAddJob
		in := &AddJobRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.AddJob(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 6: // RpcAsyncScheduleJob
		in := &ScheduleJobRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.ScheduleJob(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 7: // RpcAsyncDeletePrinter
		in := &DeletePrinterRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.DeletePrinter(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 8: // RpcAsyncSetPrinter
		in := &SetPrinterRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetPrinter(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 9: // RpcAsyncGetPrinter
		in := &GetPrinterRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetPrinter(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 10: // RpcAsyncStartDocPrinter
		in := &StartDocPrinterRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.StartDocPrinter(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 11: // RpcAsyncStartPagePrinter
		in := &StartPagePrinterRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.StartPagePrinter(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 12: // RpcAsyncWritePrinter
		in := &WritePrinterRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.WritePrinter(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 13: // RpcAsyncEndPagePrinter
		in := &EndPagePrinterRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.EndPagePrinter(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 14: // RpcAsyncEndDocPrinter
		in := &EndDocPrinterRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.EndDocPrinter(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 15: // RpcAsyncAbortPrinter
		in := &AbortPrinterRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.AbortPrinter(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 16: // RpcAsyncGetPrinterData
		in := &GetPrinterDataRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetPrinterData(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 17: // RpcAsyncGetPrinterDataEx
		in := &GetPrinterDataExRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetPrinterDataEx(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 18: // RpcAsyncSetPrinterData
		in := &SetPrinterDataRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetPrinterData(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 19: // RpcAsyncSetPrinterDataEx
		in := &SetPrinterDataExRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetPrinterDataEx(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 20: // RpcAsyncClosePrinter
		in := &ClosePrinterRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.ClosePrinter(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 21: // RpcAsyncAddForm
		in := &AddFormRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.AddForm(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 22: // RpcAsyncDeleteForm
		in := &DeleteFormRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.DeleteForm(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 23: // RpcAsyncGetForm
		in := &GetFormRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetForm(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 24: // RpcAsyncSetForm
		in := &SetFormRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetForm(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 25: // RpcAsyncEnumForms
		in := &EnumFormsRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.EnumForms(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 26: // RpcAsyncGetPrinterDriver
		in := &GetPrinterDriverRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetPrinterDriver(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 27: // RpcAsyncEnumPrinterData
		in := &EnumPrinterDataRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.EnumPrinterData(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 28: // RpcAsyncEnumPrinterDataEx
		in := &EnumPrinterDataExRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.EnumPrinterDataEx(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 29: // RpcAsyncEnumPrinterKey
		in := &EnumPrinterKeyRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.EnumPrinterKey(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 30: // RpcAsyncDeletePrinterData
		in := &DeletePrinterDataRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.DeletePrinterData(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 31: // RpcAsyncDeletePrinterDataEx
		in := &DeletePrinterDataExRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.DeletePrinterDataEx(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 32: // RpcAsyncDeletePrinterKey
		in := &DeletePrinterKeyRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.DeletePrinterKey(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 33: // RpcAsyncXcvData
		in := &XcvDataRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.XcvData(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 34: // RpcAsyncSendRecvBidiData
		in := &SendRecvBIDIDataRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SendRecvBIDIData(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 35: // RpcAsyncCreatePrinterIC
		in := &CreatePrinterICRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.CreatePrinterIC(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 36: // RpcAsyncPlayGdiScriptOnPrinterIC
		in := &PlayGDIScriptOnPrinterICRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.PlayGDIScriptOnPrinterIC(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 37: // RpcAsyncDeletePrinterIC
		in := &DeletePrinterICRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.DeletePrinterIC(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 38: // RpcAsyncEnumPrinters
		in := &EnumPrintersRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.EnumPrinters(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 39: // RpcAsyncAddPrinterDriver
		in := &AddPrinterDriverRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.AddPrinterDriver(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 40: // RpcAsyncEnumPrinterDrivers
		in := &EnumPrinterDriversRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.EnumPrinterDrivers(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 41: // RpcAsyncGetPrinterDriverDirectory
		in := &GetPrinterDriverDirectoryRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetPrinterDriverDirectory(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 42: // RpcAsyncDeletePrinterDriver
		in := &DeletePrinterDriverRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.DeletePrinterDriver(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 43: // RpcAsyncDeletePrinterDriverEx
		in := &DeletePrinterDriverExRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.DeletePrinterDriverEx(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 44: // RpcAsyncAddPrintProcessor
		in := &AddPrintProcessorRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.AddPrintProcessor(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 45: // RpcAsyncEnumPrintProcessors
		in := &EnumPrintProcessorsRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.EnumPrintProcessors(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 46: // RpcAsyncGetPrintProcessorDirectory
		in := &GetPrintProcessorDirectoryRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetPrintProcessorDirectory(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 47: // RpcAsyncEnumPorts
		in := &EnumPortsRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.EnumPorts(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 48: // RpcAsyncEnumMonitors
		in := &EnumMonitorsRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.EnumMonitors(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 49: // RpcAsyncAddPort
		in := &AddPortRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.AddPort(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 50: // RpcAsyncSetPort
		in := &SetPortRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetPort(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 51: // RpcAsyncAddMonitor
		in := &AddMonitorRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.AddMonitor(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 52: // RpcAsyncDeleteMonitor
		in := &DeleteMonitorRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.DeleteMonitor(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 53: // RpcAsyncDeletePrintProcessor
		in := &DeletePrintProcessorRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.DeletePrintProcessor(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 54: // RpcAsyncEnumPrintProcessorDatatypes
		in := &EnumPrintProcessorDataTypesRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.EnumPrintProcessorDataTypes(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 55: // RpcAsyncAddPerMachineConnection
		in := &AddPerMachineConnectionRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.AddPerMachineConnection(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 56: // RpcAsyncDeletePerMachineConnection
		in := &DeletePerMachineConnectionRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.DeletePerMachineConnection(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 57: // RpcAsyncEnumPerMachineConnections
		in := &EnumPerMachineConnectionsRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.EnumPerMachineConnections(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 58: // RpcSyncRegisterForRemoteNotifications
		in := &SyncRegisterForRemoteNotificationsRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SyncRegisterForRemoteNotifications(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 59: // RpcSyncUnRegisterForRemoteNotifications
		in := &SyncUnregisterForRemoteNotificationsRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SyncUnregisterForRemoteNotifications(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 60: // RpcSyncRefreshRemoteNotifications
		in := &SyncRefreshRemoteNotificationsRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SyncRefreshRemoteNotifications(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 61: // RpcAsyncGetRemoteNotifications
		in := &GetRemoteNotificationsRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetRemoteNotifications(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 62: // RpcAsyncInstallPrinterDriverFromPackage
		in := &InstallPrinterDriverFromPackageRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.InstallPrinterDriverFromPackage(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 63: // RpcAsyncUploadPrinterDriverPackage
		in := &UploadPrinterDriverPackageRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.UploadPrinterDriverPackage(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 64: // RpcAsyncGetCorePrinterDrivers
		in := &GetCorePrinterDriversRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetCorePrinterDrivers(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 65: // RpcAsyncCorePrinterDriverInstalled
		in := &CorePrinterDriverInstalledRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.CorePrinterDriverInstalled(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 66: // RpcAsyncGetPrinterDriverPackagePath
		in := &GetPrinterDriverPackagePathRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetPrinterDriverPackagePath(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 67: // RpcAsyncDeletePrinterDriverPackage
		in := &DeletePrinterDriverPackageRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.DeletePrinterDriverPackage(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 68: // RpcAsyncReadPrinter
		in := &ReadPrinterRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.ReadPrinter(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 69: // RpcAsyncResetPrinter
		in := &ResetPrinterRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.ResetPrinter(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 70: // RpcAsyncGetJobNamedPropertyValue
		in := &GetJobNamedPropertyValueRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetJobNamedPropertyValue(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 71: // RpcAsyncSetJobNamedProperty
		in := &SetJobNamedPropertyRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetJobNamedProperty(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 72: // RpcAsyncDeleteJobNamedProperty
		in := &DeleteJobNamedPropertyRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.DeleteJobNamedProperty(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 73: // RpcAsyncEnumJobNamedProperties
		in := &EnumJobNamedPropertiesRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.EnumJobNamedProperties(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 74: // RpcAsyncLogJobInfoForBranchOffice
		in := &LogJobInfoForBranchOfficeRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.LogJobInfoForBranchOffice(ctx, in)
		return resp.xxx_ToOp(ctx), err
	}
	return nil, nil
}
