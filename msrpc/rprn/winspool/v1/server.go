package winspool

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

// winspool server interface.
type WinspoolServer interface {

	// RpcEnumPrinters enumerates available printers, print servers, domains, or print providers.
	//
	// Return Values: This method MUST return zero (ERROR_SUCCESS) to indicate successful
	// completion and SHOULD<264> return a nonzero Windows error code ([MS-ERREF] section
	// 2.2) to indicate failure.
	EnumPrinters(context.Context, *EnumPrintersRequest) (*EnumPrintersResponse, error)

	// RpcOpenPrinter retrieves a handle for a printer, port, port monitor, print job, or
	// print server.
	//
	// Return Values: This method MUST return zero (ERROR_SUCCESS) to indicate successful
	// completion or a nonzero Windows error code to indicate failure [MS-ERREF].
	OpenPrinter(context.Context, *OpenPrinterRequest) (*OpenPrinterResponse, error)

	// RpcSetJob pauses, resumes, cancels, or restarts a print job. It also sets print job
	// parameters, such as the job priority and the document name.
	//
	// Return Values: This method MUST return zero (ERROR_SUCCESS) to indicate successful
	// completion or a nonzero Windows error code to indicate failure [MS-ERREF].
	SetJob(context.Context, *SetJobRequest) (*SetJobResponse, error)

	// RpcGetJob retrieves information about a specified print job.
	//
	// Return Values: This method MUST return zero (ERROR_SUCCESS) to indicate successful
	// completion or a nonzero Windows error code to indicate failure [MS-ERREF].
	GetJob(context.Context, *GetJobRequest) (*GetJobResponse, error)

	// RpcEnumJobs retrieves information about a specified set of print jobs for a specified
	// printer or port.
	//
	// Return Values: This method MUST return zero (ERROR_SUCCESS) to indicate successful
	// completion or a nonzero Windows error code to indicate failure [MS-ERREF].
	EnumJobs(context.Context, *EnumJobsRequest) (*EnumJobsResponse, error)

	// RpcAddPrinter adds a printer to the list of supported printers for a specified server.
	//
	// Return Values: This method MUST return zero (ERROR_SUCCESS) to indicate successful
	// completion or a nonzero Windows error code to indicate failure [MS-ERREF].
	AddPrinter(context.Context, *AddPrinterRequest) (*AddPrinterResponse, error)

	// RpcDeletePrinter is a method that deletes the specified printer object.
	//
	// Return Values: This method MUST return zero (ERROR_SUCCESS) to indicate successful
	// completion or a nonzero Windows error code to indicate failure [MS-ERREF].
	DeletePrinter(context.Context, *DeletePrinterRequest) (*DeletePrinterResponse, error)

	// RpcSetPrinter sets the data or state of a specified printer by pausing or resuming
	// printing or by clearing all print jobs.
	//
	// Return Values: This method MUST return zero (ERROR_SUCCESS) to indicate successful
	// completion or a nonzero Windows error code to indicate failure [MS-ERREF].
	SetPrinter(context.Context, *SetPrinterRequest) (*SetPrinterResponse, error)

	// RpcGetPrinter retrieves information about a specified printer.
	//
	// Return Values: This method MUST return zero (ERROR_SUCCESS) to indicate successful
	// completion or a nonzero Windows error code to indicate failure [MS-ERREF].
	GetPrinter(context.Context, *GetPrinterRequest) (*GetPrinterResponse, error)

	// RpcAddPrinterDriver installs a printer driver on the print server and links the configuration,
	// data, and printer driver files.
	//
	// Return Values: This method MUST return zero (ERROR_SUCCESS) to indicate successful
	// completion or a nonzero Windows error code to indicate failure [MS-ERREF].
	AddPrinterDriver(context.Context, *AddPrinterDriverRequest) (*AddPrinterDriverResponse, error)

	// RpcEnumPrinterDrivers enumerates the printer drivers installed on a specified print
	// server.
	//
	// Return Values: This method MUST return zero (ERROR_SUCCESS) to indicate successful
	// completion or a nonzero Windows error code to indicate failure [MS-ERREF].
	EnumPrinterDrivers(context.Context, *EnumPrinterDriversRequest) (*EnumPrinterDriversResponse, error)

	// RpcGetPrinterDriver retrieves printer driver data for the specified printer.
	//
	// Return Values: This method MUST return zero (ERROR_SUCCESS) to indicate successful
	// completion or a nonzero Windows error code to indicate failure [MS-ERREF].
	GetPrinterDriver(context.Context, *GetPrinterDriverRequest) (*GetPrinterDriverResponse, error)

	// RpcGetPrinterDriverDirectory retrieves the path of the printer driver directory.
	//
	// Return Values: This method MUST return zero (ERROR_SUCCESS) to indicate successful
	// completion or a nonzero Windows error code to indicate failure [MS-ERREF].
	GetPrinterDriverDirectory(context.Context, *GetPrinterDriverDirectoryRequest) (*GetPrinterDriverDirectoryResponse, error)

	// RpcDeletePrinterDriver removes the specified printer driver from the list of supported
	// drivers for a server.
	//
	// Return Values: This method MUST return zero (ERROR_SUCCESS) to indicate successful
	// completion or a nonzero Windows error code to indicate failure [MS-ERREF].
	DeletePrinterDriver(context.Context, *DeletePrinterDriverRequest) (*DeletePrinterDriverResponse, error)

	// RpcAddPrintProcessor installs a print processor on the specified server and adds
	// its name to an internal list of supported print processors.
	//
	// Return Values: This method MUST return zero (ERROR_SUCCESS) to indicate successful
	// completion or a nonzero Windows error code to indicate failure [MS-ERREF].
	AddPrintProcessor(context.Context, *AddPrintProcessorRequest) (*AddPrintProcessorResponse, error)

	// RpcEnumPrintProcessors enumerates the print processors installed on a specified server.
	//
	// Return Values: This method MUST return zero (ERROR_SUCCESS) to indicate successful
	// completion or a nonzero Windows error code to indicate failure [MS-ERREF].
	EnumPrintProcessors(context.Context, *EnumPrintProcessorsRequest) (*EnumPrintProcessorsResponse, error)

	// RpcGetPrintProcessorDirectory retrieves the path for the print processor on the specified
	// server.
	//
	// Return Values: This method MUST return zero (ERROR_SUCCESS) to indicate successful
	// completion or a nonzero Windows error code to indicate failure [MS-ERREF].
	GetPrintProcessorDirectory(context.Context, *GetPrintProcessorDirectoryRequest) (*GetPrintProcessorDirectoryResponse, error)

	// RpcStartDocPrinter notifies the print server that a document is being spooled for
	// printing.
	//
	// Return Values: This method MUST return zero (ERROR_SUCCESS) to indicate successful
	// completion or a nonzero Windows error code to indicate failure [MS-ERREF].
	StartDocPrinter(context.Context, *StartDocPrinterRequest) (*StartDocPrinterResponse, error)

	// RpcStartPagePrinter notifies the spooler that a page is about to be printed on the
	// specified printer.
	//
	// Return Values: This method MUST return zero (ERROR_SUCCESS) to indicate successful
	// completion or a nonzero Windows error code to indicate failure [MS-ERREF].
	StartPagePrinter(context.Context, *StartPagePrinterRequest) (*StartPagePrinterResponse, error)

	// RpcWritePrinter sends data to the print server.
	//
	// Return Values: This method MUST return zero (ERROR_SUCCESS) to indicate successful
	// completion or a nonzero Windows error code to indicate failure [MS-ERREF].
	WritePrinter(context.Context, *WritePrinterRequest) (*WritePrinterResponse, error)

	// RpcEndPagePrinter notifies the print server that the application is at the end of
	// a page in a print job.
	//
	// Return Values: This method MUST return zero (ERROR_SUCCESS) to indicate successful
	// completion or a nonzero Windows error code to indicate failure [MS-ERREF].
	EndPagePrinter(context.Context, *EndPagePrinterRequest) (*EndPagePrinterResponse, error)

	// RpcAbortPrinter aborts the currently spooling print document.
	//
	// Return Values: This method MUST return zero (ERROR_SUCCESS) to indicate successful
	// completion or a nonzero Windows error code to indicate failure [MS-ERREF].
	AbortPrinter(context.Context, *AbortPrinterRequest) (*AbortPrinterResponse, error)

	// RpcReadPrinter retrieves data from the specified job or port.
	//
	// Return Values: This method MUST return zero (ERROR_SUCCESS) to indicate successful
	// completion or a nonzero Windows error code to indicate failure [MS-ERREF].
	ReadPrinter(context.Context, *ReadPrinterRequest) (*ReadPrinterResponse, error)

	// RpcEndDocPrinter notifies the print server that the application is at the end of
	// the current print job.
	//
	// Return Values: This method MUST return zero (ERROR_SUCCESS) to indicate successful
	// completion or a nonzero Windows error code to indicate failure [MS-ERREF].
	EndDocPrinter(context.Context, *EndDocPrinterRequest) (*EndDocPrinterResponse, error)

	// RpcAddJob does not perform any function but returns a nonzero Windows error code
	// to indicate failure.
	//
	// Return Values: This method MUST return a nonzero Windows error code to indicate failure
	// [MS-ERREF].
	AddJob(context.Context, *AddJobRequest) (*AddJobResponse, error)

	// RpcScheduleJob does not perform any function, but returns a nonzero Windows error
	// code to indicate failure.
	//
	// Return Values: This method MUST return a nonzero Windows error code to indicate failure
	// [MS-ERREF].
	ScheduleJob(context.Context, *ScheduleJobRequest) (*ScheduleJobResponse, error)

	// RpcGetPrinterData retrieves printer configuration data for a printer or print server.
	//
	// Return Values: This method MUST return zero (ERROR_SUCCESS) to indicate successful
	// completion or a nonzero Windows error code to indicate failure [MS-ERREF].
	GetPrinterData(context.Context, *GetPrinterDataRequest) (*GetPrinterDataResponse, error)

	// RpcSetPrinterData sets the configuration data for a printer or print server.
	//
	// Return Values: This method MUST return zero (ERROR_SUCCESS) to indicate successful
	// completion or a nonzero Windows error code to indicate failure [MS-ERREF].
	SetPrinterData(context.Context, *SetPrinterDataRequest) (*SetPrinterDataResponse, error)

	// RpcWaitForPrinterChange retrieves information about the most recent change notification
	// that is associated with a printer or print server.
	//
	// Return Values: This method MUST return zero (ERROR_SUCCESS) to indicate successful
	// completion or a nonzero Windows error code to indicate failure [MS-ERREF].
	WaitForPrinterChange(context.Context, *WaitForPrinterChangeRequest) (*WaitForPrinterChangeResponse, error)

	// RpcClosePrinter closes a handle to a printer object, server object, job object, or
	// port object.
	//
	// Return Values: This method MUST return zero (ERROR_SUCCESS) to indicate successful
	// completion or a nonzero Windows error code to indicate failure [MS-ERREF].
	ClosePrinter(context.Context, *ClosePrinterRequest) (*ClosePrinterResponse, error)

	// RpcAddForm adds a form name to the list of supported forms.
	//
	// Return Values: This method MUST return zero (ERROR_SUCCESS) to indicate successful
	// completion or a nonzero Windows error code to indicate failure [MS-ERREF].
	AddForm(context.Context, *AddFormRequest) (*AddFormResponse, error)

	// RpcDeleteForm removes a form name from the list of supported forms.
	//
	// Return Values: This method MUST return zero (ERROR_SUCCESS) to indicate successful
	// completion or a nonzero Windows error code to indicate failure [MS-ERREF].
	DeleteForm(context.Context, *DeleteFormRequest) (*DeleteFormResponse, error)

	// RpcGetForm retrieves information about a specified form.
	//
	// Return Values: This method MUST return zero (ERROR_SUCCESS) to indicate successful
	// completion or a nonzero Windows error code to indicate failure [MS-ERREF].
	GetForm(context.Context, *GetFormRequest) (*GetFormResponse, error)

	// RpcSetForm replaces the form information for the specified form.
	//
	// Return Values: This method MUST return zero (ERROR_SUCCESS) to indicate successful
	// completion or a nonzero Windows error code to indicate failure [MS-ERREF].
	SetForm(context.Context, *SetFormRequest) (*SetFormResponse, error)

	// The RpcEnumForms method enumerates the forms that the specified printer supports.
	//
	// Return Values: This method MUST return zero (ERROR_SUCCESS) to indicate successful
	// completion or a nonzero Windows error code to indicate failure [MS-ERREF].
	EnumForms(context.Context, *EnumFormsRequest) (*EnumFormsResponse, error)

	// RpcEnumPorts enumerates the ports that are available for printing on a specified
	// server.
	//
	// Return Values: This method MUST return zero (ERROR_SUCCESS) to indicate successful
	// completion or a nonzero Windows error code to indicate failure [MS-ERREF].
	EnumPorts(context.Context, *EnumPortsRequest) (*EnumPortsResponse, error)

	// The RpcEnumMonitors method retrieves information about the port monitors installed
	// on the specified server.
	//
	// Return Values: This method MUST return zero (ERROR_SUCCESS) to indicate successful
	// completion or a nonzero Windows error code to indicate failure [MS-ERREF].
	EnumMonitors(context.Context, *EnumMonitorsRequest) (*EnumMonitorsResponse, error)

	// Opnum37NotUsedOnWire operation.
	// Opnum37NotUsedOnWire

	// Opnum38NotUsedOnWire operation.
	// Opnum38NotUsedOnWire

	// Removes a port.<357>
	//
	// Return Values: This method MUST return zero (ERROR_SUCCESS) to indicate successful
	// completion or a nonzero Windows error code to indicate failure [MS-ERREF].
	DeletePort(context.Context, *DeletePortRequest) (*DeletePortResponse, error)

	// RpcCreatePrinterIC is called by the Graphics Device Interface (GDI) to create an
	// information context for a specified printer.
	//
	// Return Values: This method MUST return zero (ERROR_SUCCESS) to indicate successful
	// completion or a nonzero Windows error code to indicate failure [MS-ERREF].
	CreatePrinterIC(context.Context, *CreatePrinterICRequest) (*CreatePrinterICResponse, error)

	// RpcPlayGdiScriptOnPrinterIC returns font information for a printer connection. UNIVERSAL_FONT_ID
	// (section 2.2.2.12) structures are used to identify the fonts.
	//
	// Return Values: This method MUST return zero (ERROR_SUCCESS) to indicate successful
	// completion or a nonzero Windows error code to indicate failure [MS-ERREF].
	PlayGDIScriptOnPrinterIC(context.Context, *PlayGDIScriptOnPrinterICRequest) (*PlayGDIScriptOnPrinterICResponse, error)

	// RpcDeletePrinterIC deletes a printer information context.
	//
	// Return Values: This method MUST return zero (ERROR_SUCCESS) to indicate successful
	// completion or a nonzero Windows error code to indicate failure [MS-ERREF].
	DeletePrinterIC(context.Context, *DeletePrinterICRequest) (*DeletePrinterICResponse, error)

	// Opnum43NotUsedOnWire operation.
	// Opnum43NotUsedOnWire

	// Opnum44NotUsedOnWire operation.
	// Opnum44NotUsedOnWire

	// Opnum45NotUsedOnWire operation.
	// Opnum45NotUsedOnWire

	// RpcAddMonitor installs a local port monitor and links the configuration, data, and
	// monitor files.
	//
	// Return Values: This method MUST return zero (ERROR_SUCCESS) to indicate successful
	// completion or a nonzero Windows error code to indicate failure [MS-ERREF].
	AddMonitor(context.Context, *AddMonitorRequest) (*AddMonitorResponse, error)

	// RpcDeleteMonitor removes a port monitor.
	//
	// Return Values: This method MUST return zero (ERROR_SUCCESS) to indicate successful
	// completion or a nonzero Windows error code to indicate failure [MS-ERREF].
	DeleteMonitor(context.Context, *DeleteMonitorRequest) (*DeleteMonitorResponse, error)

	// RpcDeletePrintProcessor removes a print processor.
	//
	// Return Values: This method MUST return zero (ERROR_SUCCESS) to indicate successful
	// completion or a nonzero Windows error code to indicate failure [MS-ERREF].
	DeletePrintProcessor(context.Context, *DeletePrintProcessorRequest) (*DeletePrintProcessorResponse, error)

	// Opnum49NotUsedOnWire operation.
	// Opnum49NotUsedOnWire

	// Opnum50NotUsedOnWire operation.
	// Opnum50NotUsedOnWire

	// RpcEnumPrintProcessorDatatypes enumerates the data types that a specified print processor
	// supports.
	//
	// Return Values: This method MUST return zero (ERROR_SUCCESS) to indicate successful
	// completion or a nonzero Windows error code to indicate failure [MS-ERREF].
	EnumPrintProcessorDataTypes(context.Context, *EnumPrintProcessorDataTypesRequest) (*EnumPrintProcessorDataTypesResponse, error)

	// RpcResetPrinter resets the data type and device mode (For more information, see [DEVMODE])
	// values to use for printing documents submitted by the RpcStartDocPrinter (section
	// 3.1.4.9.1) method.
	//
	// Return Values: This method MUST return zero (ERROR_SUCCESS) to indicate successful
	// completion or a nonzero Windows error code to indicate failure [MS-ERREF].
	ResetPrinter(context.Context, *ResetPrinterRequest) (*ResetPrinterResponse, error)

	// RpcGetPrinterDriver2 retrieves printer driver data for the specified printer.<334>
	//
	// Return Values: This method MUST return zero (ERROR_SUCCESS) to indicate successful
	// completion or a nonzero Windows error code to indicate failure [MS-ERREF].
	GetPrinterDriver2(context.Context, *GetPrinterDriver2Request) (*GetPrinterDriver2Response, error)

	// Opnum54NotUsedOnWire operation.
	// Opnum54NotUsedOnWire

	// Opnum55NotUsedOnWire operation.
	// Opnum55NotUsedOnWire

	// The RpcFindClosePrinterChangeNotification method closes a change notification object
	// created by RpcRemoteFindFirstPrinterChangeNotification (section 3.1.4.10.3) or RpcRemoteFindFirstPrinterChangeNotificationEx
	// (section 3.1.4.10.4).<380> The printer or print server associated with the change
	// notification object is no longer monitored by that object.
	//
	// Return Values: This method MUST return zero (ERROR_SUCCESS) to indicate successful
	// completion or a nonzero Windows error code to indicate failure [MS-ERREF].
	FindClosePrinterChangeNotification(context.Context, *FindClosePrinterChangeNotificationRequest) (*FindClosePrinterChangeNotificationResponse, error)

	// Opnum57NotUsedOnWire operation.
	// Opnum57NotUsedOnWire

	// RpcReplyOpenPrinter establishes a context handle from a print server to a print client.<412>
	// The server uses the RPC context handle returned by this method to send notification
	// data to the client machine.
	//
	// Return Values: This method MUST return zero (ERROR_SUCCESS) to indicate successful
	// completion or a nonzero Windows error code to indicate failure [MS-ERREF].
	ReplyOpenPrinter(context.Context, *ReplyOpenPrinterRequest) (*ReplyOpenPrinterResponse, error)

	// RpcRouterReplyPrinter handles a notification from a print server.<414>
	//
	// Return Values: This method MUST return zero (ERROR_SUCCESS) to indicate successful
	// completion or a nonzero Windows error code to indicate failure [MS-ERREF].
	RouterReplyPrinter(context.Context, *RouterReplyPrinterRequest) (*RouterReplyPrinterResponse, error)

	// RpcReplyClosePrinter closes the notification channel between a print server and a
	// print client.<415>
	//
	// Return Values: This method MUST return zero (ERROR_SUCCESS) to indicate successful
	// completion or a nonzero Windows error code to indicate failure [MS-ERREF].
	ReplyClosePrinter(context.Context, *ReplyClosePrinterRequest) (*ReplyClosePrinterResponse, error)

	// RpcAddPortEx adds a port name to the list of supported ports.<359>
	//
	// Return Values: This method MUST return zero (ERROR_SUCCESS) to indicate successful
	// completion or a nonzero Windows error code to indicate failure [MS-ERREF].
	AddPortEx(context.Context, *AddPortExRequest) (*AddPortExResponse, error)

	// RpcRemoteFindFirstPrinterChangeNotification creates a remote change notification
	// object that monitors changes to printer objects and sends change notifications to
	// a print client using either RpcRouterReplyPrinter (section 3.2.4.1.2) or RpcRouterReplyPrinterEx
	// (section 3.2.4.1.4).
	//
	// Return Values: This method MUST return zero (ERROR_SUCCESS) to indicate successful
	// completion or a nonzero Windows error code to indicate failure [MS-ERREF].
	RemoteFindFirstPrinterChangeNotification(context.Context, *RemoteFindFirstPrinterChangeNotificationRequest) (*RemoteFindFirstPrinterChangeNotificationResponse, error)

	// Opnum63NotUsedOnWire operation.
	// Opnum63NotUsedOnWire

	// Opnum64NotUsedOnWire operation.
	// Opnum64NotUsedOnWire

	// RpcRemoteFindFirstPrinterChangeNotificationEx creates a remote change notification
	// object that monitors changes to printer objects and sends change notifications to
	// a print client using either RpcRouterReplyPrinter (section 3.2.4.1.2) or RpcRouterReplyPrinterEx
	// (section 3.2.4.1.4).
	//
	// Return Values: This method MUST return zero (ERROR_SUCCESS) to indicate successful
	// completion or a nonzero Windows error code to indicate failure [MS-ERREF].
	RemoteFindFirstPrinterChangeNotificationEx(context.Context, *RemoteFindFirstPrinterChangeNotificationExRequest) (*RemoteFindFirstPrinterChangeNotificationExResponse, error)

	// RpcRouterReplyPrinterEx handles a notification from a print server.<416>
	//
	// Return Values: This method MUST return zero (ERROR_SUCCESS) to indicate successful
	// completion or a nonzero Windows error code to indicate failure [MS-ERREF].
	RouterReplyPrinterEx(context.Context, *RouterReplyPrinterExRequest) (*RouterReplyPrinterExResponse, error)

	// RpcRouterRefreshPrinterChangeNotification returns change notification information.<383>
	//
	// Return Values: This method MUST return zero (ERROR_SUCCESS) to indicate successful
	// completion or a nonzero Windows error code to indicate failure [MS-ERREF].
	RouterRefreshPrinterChangeNotification(context.Context, *RouterRefreshPrinterChangeNotificationRequest) (*RouterRefreshPrinterChangeNotificationResponse, error)

	// Opnum68NotUsedOnWire operation.
	// Opnum68NotUsedOnWire

	// RpcOpenPrinterEx retrieves a handle for a printer, port, port monitor, print job,
	// or print server.<287>
	//
	// Return Values: This method MUST return zero (ERROR_SUCCESS) to indicate successful
	// completion or a nonzero Windows error code to indicate failure [MS-ERREF].
	OpenPrinterEx(context.Context, *OpenPrinterExRequest) (*OpenPrinterExResponse, error)

	// RpcAddPrinterEx installs a printer on the print server.<288>
	//
	// Return Values: This method MUST return zero (ERROR_SUCCESS) to indicate successful
	// completion or a nonzero Windows error code to indicate failure [MS-ERREF].
	AddPrinterEx(context.Context, *AddPrinterExRequest) (*AddPrinterExResponse, error)

	// RpcSetPort sets the status associated with a printer port.<361>
	//
	// Return Values: This method MUST return zero (ERROR_SUCCESS) to indicate successful
	// completion or a nonzero Windows error code to indicate failure [MS-ERREF].
	SetPort(context.Context, *SetPortRequest) (*SetPortResponse, error)

	// RpcEnumPrinterData enumerates configuration data for a specified printer.<292>
	//
	// Return Values: This method MUST return zero (ERROR_SUCCESS) to indicate successful
	// completion or a nonzero Windows error code to indicate failure [MS-ERREF].
	EnumPrinterData(context.Context, *EnumPrinterDataRequest) (*EnumPrinterDataResponse, error)

	// RpcDeletePrinterData deletes specified configuration data for a printer.<294>
	//
	// Return Values: This method MUST return zero (ERROR_SUCCESS) to indicate successful
	// completion or a nonzero Windows error code to indicate failure [MS-ERREF].
	DeletePrinterData(context.Context, *DeletePrinterDataRequest) (*DeletePrinterDataResponse, error)

	// Opnum74NotUsedOnWire operation.
	// Opnum74NotUsedOnWire

	// Opnum75NotUsedOnWire operation.
	// Opnum75NotUsedOnWire

	// Opnum76NotUsedOnWire operation.
	// Opnum76NotUsedOnWire

	// RpcSetPrinterDataEx sets the configuration data for a printer or print server.<298>
	// This method is similar to RpcSetPrinterData (section 3.1.4.2.8) but additionally
	// allows the caller to specify the registry key under which to store the data.
	//
	// Return Values: This method MUST return zero (ERROR_SUCCESS) to indicate successful
	// completion or a nonzero Windows error code to indicate failure [MS-ERREF].
	SetPrinterDataEx(context.Context, *SetPrinterDataExRequest) (*SetPrinterDataExResponse, error)

	// RpcGetPrinterDataEx retrieves configuration data for the specified printer or print
	// server.<301> This method is similar to RpcGetPrinterData (section 3.1.4.2.7), but
	// it also allows the caller to specify the registry key from which to retrieve the
	// data.
	//
	// Return Values: This method MUST return zero (ERROR_SUCCESS) to indicate successful
	// completion or a nonzero Windows error code to indicate failure [MS-ERREF].
	GetPrinterDataEx(context.Context, *GetPrinterDataExRequest) (*GetPrinterDataExResponse, error)

	// RpcEnumPrinterDataEx enumerates all value names and data for a specified printer
	// and key.<304> This method is similar to RpcEnumPrinterData (section 3.1.4.2.16) but
	// also allows the caller to specify the registry key from which to enumerate the data,
	// and allows retrieving several values in a single call.
	//
	// Return Values: This method MUST return zero (ERROR_SUCCESS) to indicate successful
	// completion or a nonzero Windows error code to indicate failure [MS-ERREF].
	EnumPrinterDataEx(context.Context, *EnumPrinterDataExRequest) (*EnumPrinterDataExResponse, error)

	// RpcEnumPrinterKey enumerates the subkeys of a specified key for a specified printer.<305>
	//
	// Return Values: This method MUST return zero (ERROR_SUCCESS) to indicate successful
	// completion or a nonzero Windows error code to indicate failure [MS-ERREF].
	EnumPrinterKey(context.Context, *EnumPrinterKeyRequest) (*EnumPrinterKeyResponse, error)

	// RpcDeletePrinterDataEx deletes a specified value from a printer's configuration data,
	// which consists of a set of named and typed values stored in a hierarchy of registry
	// keys.<306>
	//
	// Return Values: This method MUST return zero (ERROR_SUCCESS) to indicate successful
	// completion or a nonzero Windows error code to indicate failure [MS-ERREF].
	DeletePrinterDataEx(context.Context, *DeletePrinterDataExRequest) (*DeletePrinterDataExResponse, error)

	// RpcDeletePrinterKey deletes a specified key and all of its subkeys for a specified
	// printer.<309>
	//
	// Return Values: This method MUST return zero (ERROR_SUCCESS) to indicate successful
	// completion or a nonzero Windows error code to indicate failure [MS-ERREF].
	DeletePrinterKey(context.Context, *DeletePrinterKeyRequest) (*DeletePrinterKeyResponse, error)

	// Opnum83NotUsedOnWire operation.
	// Opnum83NotUsedOnWire

	// RpcDeletePrinterDriverEx removes the specified printer driver from the list of supported
	// drivers for a server and deletes the files associated with it.<339> This method can
	// also be used to delete specific versions of a driver.
	//
	// Return Values: This method MUST return zero (ERROR_SUCCESS) to indicate successful
	// completion or a nonzero Windows error code to indicate failure [MS-ERREF].
	DeletePrinterDriverEx(context.Context, *DeletePrinterDriverExRequest) (*DeletePrinterDriverExResponse, error)

	// RpcAddPerMachineConnection adds a remote printer name to the list of supported printer
	// connections for every user who locally logs onto the computer running the print server.<311>
	//
	// This method is used for remote administration of client computers running the print
	// system.
	//
	// Return Values: This method MUST return zero (ERROR_SUCCESS) to indicate successful
	// completion or a nonzero Windows error code to indicate failure [MS-ERREF].
	AddPerMachineConnection(context.Context, *AddPerMachineConnectionRequest) (*AddPerMachineConnectionResponse, error)

	// RpcDeletePerMachineConnection deletes information about a printer connection.<314>
	//
	// This method is used for remote administration of client computers running the print
	// system.
	//
	// Return Values: This method MUST return zero (ERROR_SUCCESS) to indicate successful
	// completion or a nonzero Windows error code to indicate failure [MS-ERREF].
	DeletePerMachineConnection(context.Context, *DeletePerMachineConnectionRequest) (*DeletePerMachineConnectionResponse, error)

	// Enumerates each of the connections and copies PRINTER_INFO_4 (section 2.2.1.10.5)
	// structures for all the per-machine connections into the buffer pPrinterEnum.<316>
	//
	// This method is used for remote administration of client computers running the print
	// system.
	//
	// Return Values: This method MUST return zero (ERROR_SUCCESS) to indicate successful
	// completion or a nonzero Windows error code to indicate failure [MS-ERREF].
	EnumPerMachineConnections(context.Context, *EnumPerMachineConnectionsRequest) (*EnumPerMachineConnectionsResponse, error)

	// RpcXcvData provides an extensible mechanism by which a client can control ports on
	// the server and exchange port specific commands and data with the server.<363>
	//
	// Return Values: This method MUST return zero (ERROR_SUCCESS) to indicate that the
	// print server successfully called the port monitor's XcvData method, or a nonzero
	// Windows error code to indicate failure [MS-ERREF].
	XcvData(context.Context, *XcvDataRequest) (*XcvDataResponse, error)

	// RpcAddPrinterDriverEx installs a printer driver on the print server.<342> This method
	// performs a function similar to RpcAddPrinterDriver (section 3.1.4.4.1) and is also
	// used to specify options that permit printer driver upgrade, printer driver downgrade,
	// copying of newer files only, and copying of all files regardless of their time stamps.
	//
	// Return Values: This method MUST return zero (ERROR_SUCCESS) to indicate successful
	// completion or a nonzero Windows error code to indicate failure [MS-ERREF].
	AddPrinterDriverEx(context.Context, *AddPrinterDriverExRequest) (*AddPrinterDriverExResponse, error)

	// Opnum90NotUsedOnWire operation.
	// Opnum90NotUsedOnWire

	// Opnum91NotUsedOnWire operation.
	// Opnum91NotUsedOnWire

	// Opnum92NotUsedOnWire operation.
	// Opnum92NotUsedOnWire

	// Opnum93NotUsedOnWire operation.
	// Opnum93NotUsedOnWire

	// Opnum94NotUsedOnWire operation.
	// Opnum94NotUsedOnWire

	// Opnum95NotUsedOnWire operation.
	// Opnum95NotUsedOnWire

	// RpcFlushPrinter is used by printer drivers to send a buffer of bytes to a specified
	// port to cleanly abort a print job.<378> It also allows delaying the I/O line to the
	// printer.
	//
	// Return Values: This method MUST return zero (ERROR_SUCCESS) to indicate successful
	// completion or a nonzero Windows error code to indicate failure [MS-ERREF].
	FlushPrinter(context.Context, *FlushPrinterRequest) (*FlushPrinterResponse, error)

	// The RpcSendRecvBidiData method sends and receives bidirectional data. This method
	// is used to communicate with port monitors that support such data.<317>
	//
	// Return Values: This method MUST return zero (ERROR_SUCCESS) to indicate successful
	// completion or a nonzero Windows error code to indicate failure [MS-ERREF].
	SendRecvBIDIData(context.Context, *SendRecvBIDIDataRequest) (*SendRecvBIDIDataResponse, error)

	// Opnum98NotUsedOnWire operation.
	// Opnum98NotUsedOnWire

	// Opnum99NotUsedOnWire operation.
	// Opnum99NotUsedOnWire

	// Opnum100NotUsedOnWire operation.
	// Opnum100NotUsedOnWire

	// Opnum101NotUsedOnWire operation.
	// Opnum101NotUsedOnWire

	// RpcGetCorePrinterDrivers gets the GUIDs, versions, and publish dates of the specified
	// core printer drivers, and the paths to their packages.<347>
	//
	// Return Values: This method MUST return zero or an HRESULT success value ([MS-ERREF]
	// section 2.1) to indicate successful completion or an HRESULT error value to indicate
	// failure.
	//
	// Exceptions Thrown: This method MUST NOT throw any exceptions other than those that
	// are thrown by the underlying RPC protocol [MS-RPCE].
	GetCorePrinterDrivers(context.Context, *GetCorePrinterDriversRequest) (*GetCorePrinterDriversResponse, error)

	// Opnum103NotUsedOnWire operation.
	// Opnum103NotUsedOnWire

	// RpcGetPrinterDriverPackagePath gets the path to the specified printer driver package.<349>
	//
	// Return Values: This method MUST return zero or an HRESULT success value ([MS-ERREF]
	// section 2.1) to indicate successful completion or an HRESULT error value to indicate
	// failure.
	//
	// Exceptions Thrown: This method MUST NOT throw any exceptions other than those that
	// are thrown by the underlying RPC protocol [MS-RPCE].
	GetPrinterDriverPackagePath(context.Context, *GetPrinterDriverPackagePathRequest) (*GetPrinterDriverPackagePathResponse, error)

	// Opnum105NotUsedOnWire operation.
	// Opnum105NotUsedOnWire

	// Opnum106NotUsedOnWire operation.
	// Opnum106NotUsedOnWire

	// Opnum107NotUsedOnWire operation.
	// Opnum107NotUsedOnWire

	// Opnum108NotUsedOnWire operation.
	// Opnum108NotUsedOnWire

	// Opnum109NotUsedOnWire operation.
	// Opnum109NotUsedOnWire

	// RpcGetJobNamedPropertyValue retrieves the current value of the specified Job Named
	// Property (section 3.1.1).<397>
	//
	// Return Values: This method MUST return zero (ERROR_SUCCESS) to indicate successful
	// completion or a nonzero Windows error code to indicate failure [MS-ERREF].
	GetJobNamedPropertyValue(context.Context, *GetJobNamedPropertyValueRequest) (*GetJobNamedPropertyValueResponse, error)

	// RpcSetJobNamedProperty creates a new Job Named Property (section 3.1.1), or changes
	// the value of an existing Job Named Property for the specified print job.<398>
	//
	// Return Values: This method MUST return zero (ERROR_SUCCESS) to indicate successful
	// completion or a nonzero Windows error code to indicate failure [MS-ERREF].
	SetJobNamedProperty(context.Context, *SetJobNamedPropertyRequest) (*SetJobNamedPropertyResponse, error)

	// RpcDeleteJobNamedProperty deletes an existing Job Named Property (section 3.1.1)
	// for the specified print job.<399>
	//
	// Return Values: This method MUST return zero (ERROR_SUCCESS) to indicate successful
	// completion or a nonzero Windows error code to indicate failure [MS-ERREF].
	DeleteJobNamedProperty(context.Context, *DeleteJobNamedPropertyRequest) (*DeleteJobNamedPropertyResponse, error)

	// RpcEnumJobNamedProperties enumerates the Job Named Properties (section 3.1.1) for
	// the specified print job.<400>
	//
	// Return Values: This method MUST return zero (ERROR_SUCCESS) to indicate successful
	// completion or a nonzero Windows error code to indicate failure [MS-ERREF].
	EnumJobNamedProperties(context.Context, *EnumJobNamedPropertiesRequest) (*EnumJobNamedPropertiesResponse, error)

	// Opnum114NotUsedOnWire operation.
	// Opnum114NotUsedOnWire

	// Opnum115NotUsedOnWire operation.
	// Opnum115NotUsedOnWire

	// RpcLogJobInfoForBranchOffice operation.
	LogJobInfoForBranchOffice(context.Context, *LogJobInfoForBranchOfficeRequest) (*LogJobInfoForBranchOfficeResponse, error)

	// RpcRegeneratePrintDeviceCapabilities operation.
	RegeneratePrintDeviceCapabilities(context.Context, *RegeneratePrintDeviceCapabilitiesRequest) (*RegeneratePrintDeviceCapabilitiesResponse, error)

	// Opnum118NotUsedOnWire operation.
	// Opnum118NotUsedOnWire

	// RpcIppCreateJobOnPrinter operation.
	CreateJobOnPrinter(context.Context, *CreateJobOnPrinterRequest) (*CreateJobOnPrinterResponse, error)

	// RpcIppGetJobAttributes operation.
	GetJobAttributes(context.Context, *GetJobAttributesRequest) (*GetJobAttributesResponse, error)

	// RpcIppSetJobAttributes operation.
	SetJobAttributes(context.Context, *SetJobAttributesRequest) (*SetJobAttributesResponse, error)

	// RpcIppGetPrinterAttributes operation.
	GetPrinterAttributes(context.Context, *GetPrinterAttributesRequest) (*GetPrinterAttributesResponse, error)

	// RpcIppSetPrinterAttributes operation.
	SetPrinterAttributes(context.Context, *SetPrinterAttributesRequest) (*SetPrinterAttributesResponse, error)
}

func RegisterWinspoolServer(conn dcerpc.Conn, o WinspoolServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewWinspoolServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(WinspoolSyntaxV1_0))...)
}

func NewWinspoolServerHandle(o WinspoolServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return WinspoolServerHandle(ctx, o, opNum, r)
	}
}

func WinspoolServerHandle(ctx context.Context, o WinspoolServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	switch opNum {
	case 0: // RpcEnumPrinters
		in := &EnumPrintersRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.EnumPrinters(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 1: // RpcOpenPrinter
		in := &OpenPrinterRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.OpenPrinter(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 2: // RpcSetJob
		in := &SetJobRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetJob(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 3: // RpcGetJob
		in := &GetJobRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetJob(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 4: // RpcEnumJobs
		in := &EnumJobsRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.EnumJobs(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 5: // RpcAddPrinter
		in := &AddPrinterRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.AddPrinter(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 6: // RpcDeletePrinter
		in := &DeletePrinterRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.DeletePrinter(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 7: // RpcSetPrinter
		in := &SetPrinterRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetPrinter(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 8: // RpcGetPrinter
		in := &GetPrinterRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetPrinter(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 9: // RpcAddPrinterDriver
		in := &AddPrinterDriverRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.AddPrinterDriver(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 10: // RpcEnumPrinterDrivers
		in := &EnumPrinterDriversRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.EnumPrinterDrivers(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 11: // RpcGetPrinterDriver
		in := &GetPrinterDriverRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetPrinterDriver(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 12: // RpcGetPrinterDriverDirectory
		in := &GetPrinterDriverDirectoryRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetPrinterDriverDirectory(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 13: // RpcDeletePrinterDriver
		in := &DeletePrinterDriverRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.DeletePrinterDriver(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 14: // RpcAddPrintProcessor
		in := &AddPrintProcessorRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.AddPrintProcessor(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 15: // RpcEnumPrintProcessors
		in := &EnumPrintProcessorsRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.EnumPrintProcessors(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 16: // RpcGetPrintProcessorDirectory
		in := &GetPrintProcessorDirectoryRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetPrintProcessorDirectory(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 17: // RpcStartDocPrinter
		in := &StartDocPrinterRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.StartDocPrinter(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 18: // RpcStartPagePrinter
		in := &StartPagePrinterRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.StartPagePrinter(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 19: // RpcWritePrinter
		in := &WritePrinterRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.WritePrinter(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 20: // RpcEndPagePrinter
		in := &EndPagePrinterRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.EndPagePrinter(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 21: // RpcAbortPrinter
		in := &AbortPrinterRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.AbortPrinter(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 22: // RpcReadPrinter
		in := &ReadPrinterRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.ReadPrinter(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 23: // RpcEndDocPrinter
		in := &EndDocPrinterRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.EndDocPrinter(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 24: // RpcAddJob
		in := &AddJobRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.AddJob(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 25: // RpcScheduleJob
		in := &ScheduleJobRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.ScheduleJob(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 26: // RpcGetPrinterData
		in := &GetPrinterDataRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetPrinterData(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 27: // RpcSetPrinterData
		in := &SetPrinterDataRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetPrinterData(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 28: // RpcWaitForPrinterChange
		in := &WaitForPrinterChangeRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.WaitForPrinterChange(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 29: // RpcClosePrinter
		in := &ClosePrinterRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.ClosePrinter(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 30: // RpcAddForm
		in := &AddFormRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.AddForm(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 31: // RpcDeleteForm
		in := &DeleteFormRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.DeleteForm(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 32: // RpcGetForm
		in := &GetFormRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetForm(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 33: // RpcSetForm
		in := &SetFormRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetForm(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 34: // RpcEnumForms
		in := &EnumFormsRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.EnumForms(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 35: // RpcEnumPorts
		in := &EnumPortsRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.EnumPorts(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 36: // RpcEnumMonitors
		in := &EnumMonitorsRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.EnumMonitors(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 37: // Opnum37NotUsedOnWire
		// Opnum37NotUsedOnWire
		return nil, nil
	case 38: // Opnum38NotUsedOnWire
		// Opnum38NotUsedOnWire
		return nil, nil
	case 39: // RpcDeletePort
		in := &DeletePortRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.DeletePort(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 40: // RpcCreatePrinterIC
		in := &CreatePrinterICRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.CreatePrinterIC(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 41: // RpcPlayGdiScriptOnPrinterIC
		in := &PlayGDIScriptOnPrinterICRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.PlayGDIScriptOnPrinterIC(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 42: // RpcDeletePrinterIC
		in := &DeletePrinterICRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.DeletePrinterIC(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 43: // Opnum43NotUsedOnWire
		// Opnum43NotUsedOnWire
		return nil, nil
	case 44: // Opnum44NotUsedOnWire
		// Opnum44NotUsedOnWire
		return nil, nil
	case 45: // Opnum45NotUsedOnWire
		// Opnum45NotUsedOnWire
		return nil, nil
	case 46: // RpcAddMonitor
		in := &AddMonitorRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.AddMonitor(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 47: // RpcDeleteMonitor
		in := &DeleteMonitorRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.DeleteMonitor(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 48: // RpcDeletePrintProcessor
		in := &DeletePrintProcessorRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.DeletePrintProcessor(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 49: // Opnum49NotUsedOnWire
		// Opnum49NotUsedOnWire
		return nil, nil
	case 50: // Opnum50NotUsedOnWire
		// Opnum50NotUsedOnWire
		return nil, nil
	case 51: // RpcEnumPrintProcessorDatatypes
		in := &EnumPrintProcessorDataTypesRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.EnumPrintProcessorDataTypes(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 52: // RpcResetPrinter
		in := &ResetPrinterRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.ResetPrinter(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 53: // RpcGetPrinterDriver2
		in := &GetPrinterDriver2Request{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetPrinterDriver2(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 54: // Opnum54NotUsedOnWire
		// Opnum54NotUsedOnWire
		return nil, nil
	case 55: // Opnum55NotUsedOnWire
		// Opnum55NotUsedOnWire
		return nil, nil
	case 56: // RpcFindClosePrinterChangeNotification
		in := &FindClosePrinterChangeNotificationRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.FindClosePrinterChangeNotification(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 57: // Opnum57NotUsedOnWire
		// Opnum57NotUsedOnWire
		return nil, nil
	case 58: // RpcReplyOpenPrinter
		in := &ReplyOpenPrinterRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.ReplyOpenPrinter(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 59: // RpcRouterReplyPrinter
		in := &RouterReplyPrinterRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.RouterReplyPrinter(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 60: // RpcReplyClosePrinter
		in := &ReplyClosePrinterRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.ReplyClosePrinter(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 61: // RpcAddPortEx
		in := &AddPortExRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.AddPortEx(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 62: // RpcRemoteFindFirstPrinterChangeNotification
		in := &RemoteFindFirstPrinterChangeNotificationRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.RemoteFindFirstPrinterChangeNotification(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 63: // Opnum63NotUsedOnWire
		// Opnum63NotUsedOnWire
		return nil, nil
	case 64: // Opnum64NotUsedOnWire
		// Opnum64NotUsedOnWire
		return nil, nil
	case 65: // RpcRemoteFindFirstPrinterChangeNotificationEx
		in := &RemoteFindFirstPrinterChangeNotificationExRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.RemoteFindFirstPrinterChangeNotificationEx(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 66: // RpcRouterReplyPrinterEx
		in := &RouterReplyPrinterExRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.RouterReplyPrinterEx(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 67: // RpcRouterRefreshPrinterChangeNotification
		in := &RouterRefreshPrinterChangeNotificationRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.RouterRefreshPrinterChangeNotification(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 68: // Opnum68NotUsedOnWire
		// Opnum68NotUsedOnWire
		return nil, nil
	case 69: // RpcOpenPrinterEx
		in := &OpenPrinterExRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.OpenPrinterEx(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 70: // RpcAddPrinterEx
		in := &AddPrinterExRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.AddPrinterEx(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 71: // RpcSetPort
		in := &SetPortRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetPort(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 72: // RpcEnumPrinterData
		in := &EnumPrinterDataRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.EnumPrinterData(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 73: // RpcDeletePrinterData
		in := &DeletePrinterDataRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.DeletePrinterData(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 74: // Opnum74NotUsedOnWire
		// Opnum74NotUsedOnWire
		return nil, nil
	case 75: // Opnum75NotUsedOnWire
		// Opnum75NotUsedOnWire
		return nil, nil
	case 76: // Opnum76NotUsedOnWire
		// Opnum76NotUsedOnWire
		return nil, nil
	case 77: // RpcSetPrinterDataEx
		in := &SetPrinterDataExRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetPrinterDataEx(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 78: // RpcGetPrinterDataEx
		in := &GetPrinterDataExRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetPrinterDataEx(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 79: // RpcEnumPrinterDataEx
		in := &EnumPrinterDataExRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.EnumPrinterDataEx(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 80: // RpcEnumPrinterKey
		in := &EnumPrinterKeyRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.EnumPrinterKey(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 81: // RpcDeletePrinterDataEx
		in := &DeletePrinterDataExRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.DeletePrinterDataEx(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 82: // RpcDeletePrinterKey
		in := &DeletePrinterKeyRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.DeletePrinterKey(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 83: // Opnum83NotUsedOnWire
		// Opnum83NotUsedOnWire
		return nil, nil
	case 84: // RpcDeletePrinterDriverEx
		in := &DeletePrinterDriverExRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.DeletePrinterDriverEx(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 85: // RpcAddPerMachineConnection
		in := &AddPerMachineConnectionRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.AddPerMachineConnection(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 86: // RpcDeletePerMachineConnection
		in := &DeletePerMachineConnectionRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.DeletePerMachineConnection(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 87: // RpcEnumPerMachineConnections
		in := &EnumPerMachineConnectionsRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.EnumPerMachineConnections(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 88: // RpcXcvData
		in := &XcvDataRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.XcvData(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 89: // RpcAddPrinterDriverEx
		in := &AddPrinterDriverExRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.AddPrinterDriverEx(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 90: // Opnum90NotUsedOnWire
		// Opnum90NotUsedOnWire
		return nil, nil
	case 91: // Opnum91NotUsedOnWire
		// Opnum91NotUsedOnWire
		return nil, nil
	case 92: // Opnum92NotUsedOnWire
		// Opnum92NotUsedOnWire
		return nil, nil
	case 93: // Opnum93NotUsedOnWire
		// Opnum93NotUsedOnWire
		return nil, nil
	case 94: // Opnum94NotUsedOnWire
		// Opnum94NotUsedOnWire
		return nil, nil
	case 95: // Opnum95NotUsedOnWire
		// Opnum95NotUsedOnWire
		return nil, nil
	case 96: // RpcFlushPrinter
		in := &FlushPrinterRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.FlushPrinter(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 97: // RpcSendRecvBidiData
		in := &SendRecvBIDIDataRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SendRecvBIDIData(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 98: // Opnum98NotUsedOnWire
		// Opnum98NotUsedOnWire
		return nil, nil
	case 99: // Opnum99NotUsedOnWire
		// Opnum99NotUsedOnWire
		return nil, nil
	case 100: // Opnum100NotUsedOnWire
		// Opnum100NotUsedOnWire
		return nil, nil
	case 101: // Opnum101NotUsedOnWire
		// Opnum101NotUsedOnWire
		return nil, nil
	case 102: // RpcGetCorePrinterDrivers
		in := &GetCorePrinterDriversRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetCorePrinterDrivers(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 103: // Opnum103NotUsedOnWire
		// Opnum103NotUsedOnWire
		return nil, nil
	case 104: // RpcGetPrinterDriverPackagePath
		in := &GetPrinterDriverPackagePathRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetPrinterDriverPackagePath(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 105: // Opnum105NotUsedOnWire
		// Opnum105NotUsedOnWire
		return nil, nil
	case 106: // Opnum106NotUsedOnWire
		// Opnum106NotUsedOnWire
		return nil, nil
	case 107: // Opnum107NotUsedOnWire
		// Opnum107NotUsedOnWire
		return nil, nil
	case 108: // Opnum108NotUsedOnWire
		// Opnum108NotUsedOnWire
		return nil, nil
	case 109: // Opnum109NotUsedOnWire
		// Opnum109NotUsedOnWire
		return nil, nil
	case 110: // RpcGetJobNamedPropertyValue
		in := &GetJobNamedPropertyValueRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetJobNamedPropertyValue(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 111: // RpcSetJobNamedProperty
		in := &SetJobNamedPropertyRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetJobNamedProperty(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 112: // RpcDeleteJobNamedProperty
		in := &DeleteJobNamedPropertyRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.DeleteJobNamedProperty(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 113: // RpcEnumJobNamedProperties
		in := &EnumJobNamedPropertiesRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.EnumJobNamedProperties(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 114: // Opnum114NotUsedOnWire
		// Opnum114NotUsedOnWire
		return nil, nil
	case 115: // Opnum115NotUsedOnWire
		// Opnum115NotUsedOnWire
		return nil, nil
	case 116: // RpcLogJobInfoForBranchOffice
		in := &LogJobInfoForBranchOfficeRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.LogJobInfoForBranchOffice(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 117: // RpcRegeneratePrintDeviceCapabilities
		in := &RegeneratePrintDeviceCapabilitiesRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.RegeneratePrintDeviceCapabilities(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 118: // Opnum118NotUsedOnWire
		// Opnum118NotUsedOnWire
		return nil, nil
	case 119: // RpcIppCreateJobOnPrinter
		in := &CreateJobOnPrinterRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.CreateJobOnPrinter(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 120: // RpcIppGetJobAttributes
		in := &GetJobAttributesRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetJobAttributes(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 121: // RpcIppSetJobAttributes
		in := &SetJobAttributesRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetJobAttributes(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 122: // RpcIppGetPrinterAttributes
		in := &GetPrinterAttributesRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetPrinterAttributes(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 123: // RpcIppSetPrinterAttributes
		in := &SetPrinterAttributesRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetPrinterAttributes(ctx, in)
		return resp.xxx_ToOp(ctx), err
	}
	return nil, nil
}
