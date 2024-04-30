package ifsrmclassificationmanager

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	idispatch "github.com/oiweiwei/go-msrpc/msrpc/dcom/oaut/idispatch/v0"
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
	_ = idispatch.GoPackage
)

// IFsrmClassificationManager server interface.
type ClassificationManagerServer interface {

	// IDispatch base class.
	idispatch.DispatchServer

	// ClassificationReportFormats operation.
	GetClassificationReportFormats(context.Context, *GetClassificationReportFormatsRequest) (*GetClassificationReportFormatsResponse, error)

	// ClassificationReportFormats operation.
	SetClassificationReportFormats(context.Context, *SetClassificationReportFormatsRequest) (*SetClassificationReportFormatsResponse, error)

	// Logging operation.
	GetLogging(context.Context, *GetLoggingRequest) (*GetLoggingResponse, error)

	// Logging operation.
	SetLogging(context.Context, *SetLoggingRequest) (*SetLoggingResponse, error)

	// ClassificationReportMailTo operation.
	GetClassificationReportMailTo(context.Context, *GetClassificationReportMailToRequest) (*GetClassificationReportMailToResponse, error)

	// ClassificationReportMailTo operation.
	SetClassificationReportMailTo(context.Context, *SetClassificationReportMailToRequest) (*SetClassificationReportMailToResponse, error)

	// ClassificationReportEnabled operation.
	GetClassificationReportEnabled(context.Context, *GetClassificationReportEnabledRequest) (*GetClassificationReportEnabledResponse, error)

	// ClassificationReportEnabled operation.
	SetClassificationReportEnabled(context.Context, *SetClassificationReportEnabledRequest) (*SetClassificationReportEnabledResponse, error)

	// The ClassificationLastReportPathWithoutExtension (get) method retrieves the local
	// directory path and file name where the generated report(s) was (were) stored when
	// classification was previously run and returns S_OK upon successful completion.
	//
	// Return Values: The method MUST return zero on success, or a nonzero error code on
	// failure.
	//
	//	+-------------------------+---------------------------------------+
	//	|         RETURN          |                                       |
	//	|       VALUE/CODE        |              DESCRIPTION              |
	//	|                         |                                       |
	//	+-------------------------+---------------------------------------+
	//	+-------------------------+---------------------------------------+
	//	| 0x80070057 E_INVALIDARG | The lastReportPath parameter is NULL. |
	//	+-------------------------+---------------------------------------+
	GetClassificationLastReportPathWithoutExtension(context.Context, *GetClassificationLastReportPathWithoutExtensionRequest) (*GetClassificationLastReportPathWithoutExtensionResponse, error)

	// The ClassificationLastError (get) method retrieves the last error, if any, from when
	// classification was previously run and returns S_OK upon successful completion. If
	// no error occurred on the previous classification run, the returned string will be
	// empty and ClassificationLastError (get) MUST return S_FALSE.
	//
	// Return Values: The method MUST return zero on success, or a nonzero error code on
	// failure.
	//
	//	+-------------------------+----------------------------------+
	//	|         RETURN          |                                  |
	//	|       VALUE/CODE        |           DESCRIPTION            |
	//	|                         |                                  |
	//	+-------------------------+----------------------------------+
	//	+-------------------------+----------------------------------+
	//	| 0x80070057 E_INVALIDARG | The lastError parameter is NULL. |
	//	+-------------------------+----------------------------------+
	GetClassificationLastError(context.Context, *GetClassificationLastErrorRequest) (*GetClassificationLastErrorResponse, error)

	// The ClassificationRunningStatus (get) method retrieves the current running status
	// of the running classification task, if present, as defined in the FsrmReportRunningStatus
	// (section 2.2.1.2.13) enumeration and returns S_OK upon successful completion.
	//
	// Return Values: The method MUST return zero on success, or a nonzero error code on
	// failure.
	//
	//	+-------------------------+--------------------------------------+
	//	|         RETURN          |                                      |
	//	|       VALUE/CODE        |             DESCRIPTION              |
	//	|                         |                                      |
	//	+-------------------------+--------------------------------------+
	//	+-------------------------+--------------------------------------+
	//	| 0x80070057 E_INVALIDARG | The runningStatus parameter is NULL. |
	//	+-------------------------+--------------------------------------+
	GetClassificationRunningStatus(context.Context, *GetClassificationRunningStatusRequest) (*GetClassificationRunningStatusResponse, error)

	// The EnumPropertyDefinitions method returns all the property definitions from the
	// List of Persisted Property Definitions (section 3.2.1.6) on the server.
	//
	// Return Values: The method MUST return zero on success, or a nonzero error code on
	// failure.
	//
	//	+---------------------------------+---------------------------------------------------------------------------------+
	//	|             RETURN              |                                                                                 |
	//	|           VALUE/CODE            |                                   DESCRIPTION                                   |
	//	|                                 |                                                                                 |
	//	+---------------------------------+---------------------------------------------------------------------------------+
	//	+---------------------------------+---------------------------------------------------------------------------------+
	//	| 0x80045311 FSRM_E_NOT_SUPPORTED | The options parameter is not a valid FsrmEnumOptions (section 2.2.1.2.5) value. |
	//	+---------------------------------+---------------------------------------------------------------------------------+
	//	| 0x80070057 E_INVALIDARG         | The propertyDefinitions parameter is NULL.                                      |
	//	+---------------------------------+---------------------------------------------------------------------------------+
	EnumPropertyDefinitions(context.Context, *EnumPropertyDefinitionsRequest) (*EnumPropertyDefinitionsResponse, error)

	// The CreatePropertyDefinition method creates a blank Non-Persisted Property Definition
	// Instance (section 3.2.1.6.1.2) and returns S_OK upon successful completion.
	//
	// Return Values: The method MUST return zero on success, or a nonzero error code on
	// failure.
	//
	//	+-------------------------+-------------------------------------------+
	//	|         RETURN          |                                           |
	//	|       VALUE/CODE        |                DESCRIPTION                |
	//	|                         |                                           |
	//	+-------------------------+-------------------------------------------+
	//	+-------------------------+-------------------------------------------+
	//	| 0x80070057 E_INVALIDARG | The propertyDefinition parameter is NULL. |
	//	+-------------------------+-------------------------------------------+
	CreatePropertyDefinition(context.Context, *CreatePropertyDefinitionRequest) (*CreatePropertyDefinitionResponse, error)

	// The GetPropertyDefinition method returns a pointer to the property definition from
	// the List of Persisted Property Definitions (section 3.2.1.6) with the specified name.
	//
	// Return Values: The method MUST return zero on success, or a nonzero error code on
	// failure.
	//
	//	+--------------------------------+-------------------------------------------------------+
	//	|             RETURN             |                                                       |
	//	|           VALUE/CODE           |                      DESCRIPTION                      |
	//	|                                |                                                       |
	//	+--------------------------------+-------------------------------------------------------+
	//	+--------------------------------+-------------------------------------------------------+
	//	| 0x80045301 FSRM_E_NOT_FOUND    | The specified property definition could not be found. |
	//	+--------------------------------+-------------------------------------------------------+
	//	| 0x80045308 FSRM_E_INVALID_NAME | The propertyName parameter is empty or NULL.          |
	//	+--------------------------------+-------------------------------------------------------+
	//	| 0x80070057 E_INVALIDARG        | The propertyDefinition parameter is NULL.             |
	//	+--------------------------------+-------------------------------------------------------+
	GetPropertyDefinition(context.Context, *GetPropertyDefinitionRequest) (*GetPropertyDefinitionResponse, error)

	// The EnumRules method returns all the Rules from the List of Persisted Rules (section
	// 3.2.1.6) on the server that have the specified ruleType.
	//
	// Return Values: The method MUST return zero on success, or a nonzero error code on
	// failure.
	//
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	|             RETURN              |                                                                                  |
	//	|           VALUE/CODE            |                                   DESCRIPTION                                    |
	//	|                                 |                                                                                  |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80045311 FSRM_E_NOT_SUPPORTED | The options parameter is not a valid FsrmEnumOptions (section 2.2.1.2.5) value.  |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070057 E_INVALIDARG         | This code is returned for the following reasons: The ruleType parameter is not a |
	//	|                                 | valid FsrmRuleType (section 2.2.1.2.11) value. The Rules parameter is NULL.      |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	EnumRules(context.Context, *EnumRulesRequest) (*EnumRulesResponse, error)

	// The CreateRule method creates a blank Non-Persisted Rule Instance (section 3.2.1.6.3.2)
	// with the specified classification rule type and returns S_OK upon successful completion.
	//
	// Return Values: The method MUST return zero on success, or a nonzero error code on
	// failure.
	//
	//	+-------------------------+----------------------------------------------------------------------------------+
	//	|         RETURN          |                                                                                  |
	//	|       VALUE/CODE        |                                   DESCRIPTION                                    |
	//	|                         |                                                                                  |
	//	+-------------------------+----------------------------------------------------------------------------------+
	//	+-------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070057 E_INVALIDARG | This code is returned for the following reasons: The ruleType parameter is not   |
	//	|                         | FsrmRuleType_Classification. The rule parameter is NULL.                         |
	//	+-------------------------+----------------------------------------------------------------------------------+
	CreateRule(context.Context, *CreateRuleRequest) (*CreateRuleResponse, error)

	// The GetRule method returns a pointer to the classification rule with the specified
	// Name and Rule type from the List of Persisted Rules (section 3.2.1.6).
	//
	// Return Values: The method MUST return zero on success, or a nonzero error code on
	// failure.
	//
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	|             RETURN             |                                                                                  |
	//	|           VALUE/CODE           |                                   DESCRIPTION                                    |
	//	|                                |                                                                                  |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80045301 FSRM_E_NOT_FOUND    | The specified rule could not be found.                                           |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80045308 FSRM_E_INVALID_NAME | The ruleName parameter is empty or NULL.                                         |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070057 E_INVALIDARG        | This code is returned for the following reasons: The ruleType parameter is a     |
	//	|                                | not a valid value. If ruleType  is FsrmRuleType_Generic, the parameter MUST be   |
	//	|                                | considered an valid value. The specified name exceeds the maximum length of 1000 |
	//	|                                | characters.                                                                      |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	GetRule(context.Context, *GetRuleRequest) (*GetRuleResponse, error)

	// The EnumModuleDefinitions method returns all the moduleDefinitions from the List
	// of Persisted Module Definitions (section 3.2.1.6) on the server that have the specified
	// moduleType.
	//
	// Return Values: The method MUST return zero on success, or a nonzero error code on
	// failure.
	//
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	|             RETURN              |                                                                                  |
	//	|           VALUE/CODE            |                                   DESCRIPTION                                    |
	//	|                                 |                                                                                  |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80045311 FSRM_E_NOT_SUPPORTED | The options parameter does not contain a valid FsrmEnumOptions (section          |
	//	|                                 | 2.2.1.2.5) value.                                                                |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070057 E_INVALIDARG         | This code is returned for the following reasons: The moduleType parameter is     |
	//	|                                 | not a valid value. If the moduleType is FsrmPipelineModuleType_Unknown, the      |
	//	|                                 | parameter MUST be considered an invalid value. The moduleDefinitions parameter   |
	//	|                                 | is NULL.                                                                         |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	EnumModuleDefinitions(context.Context, *EnumModuleDefinitionsRequest) (*EnumModuleDefinitionsResponse, error)

	// The CreateModuleDefinition method is used to create a new Non-Persisted Module Definition
	// Instance (section 3.2.1.6.2.2) of a specified module type and returns S_OK upon successful
	// completion.
	//
	// Return Values: The method MUST return zero on success, or an error code on failure.
	//
	//	+-------------------------+----------------------------------------------------------------------------------+
	//	|         RETURN          |                                                                                  |
	//	|       VALUE/CODE        |                                   DESCRIPTION                                    |
	//	|                         |                                                                                  |
	//	+-------------------------+----------------------------------------------------------------------------------+
	//	+-------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070057 E_INVALIDARG | The moduleType parameter is not a valid FsrmPipelineModuleType (section          |
	//	|                         | 2.2.1.2.12) value.                                                               |
	//	+-------------------------+----------------------------------------------------------------------------------+
	CreateModuleDefinition(context.Context, *CreateModuleDefinitionRequest) (*CreateModuleDefinitionResponse, error)

	// The GetModuleDefinition method returns a pointer to the module definition with the
	// specified Name and Module type from the List of Persisted Module Definitions (section
	// 3.2.1.6).
	//
	// Return Values: The method MUST return zero on success, or a nonzero error code on
	// failure.
	//
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	|             RETURN             |                                                                                  |
	//	|           VALUE/CODE           |                                   DESCRIPTION                                    |
	//	|                                |                                                                                  |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80045301 FSRM_E_NOT_FOUND    | The specified module definition could not be found.                              |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80045308 FSRM_E_INVALID_NAME | The specified name is empty or NULL.                                             |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070057 E_INVALIDARG        | This code is returned for the following reasons: The moduleType parameter is not |
	//	|                                | a valid FsrmPipelineModuleType (section 2.2.1.2.12) value. The specified name    |
	//	|                                | exceeds the maximum length of 100 characters. The moduleDefinitions parameter is |
	//	|                                | NULL.                                                                            |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	GetModuleDefinition(context.Context, *GetModuleDefinitionRequest) (*GetModuleDefinitionResponse, error)

	// The RunClassification method queues a Running Job to the Classification Job Queue.
	//
	// Return Values: The method MUST return zero on success, or a nonzero error code on
	// failure.
	//
	//	+--------------------------------------------------+----------------------------------------------------------------------------------+
	//	|                      RETURN                      |                                                                                  |
	//	|                    VALUE/CODE                    |                                   DESCRIPTION                                    |
	//	|                                                  |                                                                                  |
	//	+--------------------------------------------------+----------------------------------------------------------------------------------+
	//	+--------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x8004533D FSRM_E_CLASSIFICATION_ALREADY_RUNNING | Classification is already running.                                               |
	//	+--------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070057 E_INVALIDARG                          | The context parameter is not a valid FsrmReportGenerationContext (section        |
	//	|                                                  | 2.2.1.2.15) value.                                                               |
	//	+--------------------------------------------------+----------------------------------------------------------------------------------+
	RunClassification(context.Context, *RunClassificationRequest) (*RunClassificationResponse, error)

	// The WaitForClassificationCompletion method blocks the caller for the specified time
	// period or until the running job associated with the single class (section 3.2.1.12.2),
	// if present, completes, whichever occurs first.
	//
	// Return Values: The method MUST return zero on success, or a nonzero error code on
	// failure.
	//
	//	+-------------------------+----------------------------------------------------------------------------------+
	//	|         RETURN          |                                                                                  |
	//	|       VALUE/CODE        |                                   DESCRIPTION                                    |
	//	|                         |                                                                                  |
	//	+-------------------------+----------------------------------------------------------------------------------+
	//	+-------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070057 E_INVALIDARG | The waitSeconds parameter is not a valid value; the number of seconds to wait    |
	//	|                         | must be in the range of -1 through 2,147,483.                                    |
	//	+-------------------------+----------------------------------------------------------------------------------+
	WaitForClassificationCompletion(context.Context, *WaitForClassificationCompletionRequest) (*WaitForClassificationCompletionResponse, error)

	// The CancelClassification method stops the running job associated with the Default
	// Classification job object, if present, from continuing and returns S_OK upon successful
	// completion.
	//
	// This method has no parameters.
	//
	// Return Values: The method MUST return zero on success, or a nonzero error code on
	// failure.
	//
	//	+----------------------------------------------+----------------------------------------------+
	//	|                    RETURN                    |                                              |
	//	|                  VALUE/CODE                  |                 DESCRIPTION                  |
	//	|                                              |                                              |
	//	+----------------------------------------------+----------------------------------------------+
	//	+----------------------------------------------+----------------------------------------------+
	//	| 0x8004533E FSRM_E_CLASSIFICATION_NOT_RUNNING | The classification is not currently running. |
	//	+----------------------------------------------+----------------------------------------------+
	CancelClassification(context.Context, *CancelClassificationRequest) (*CancelClassificationResponse, error)

	// The EnumFileProperties method enumerates the Property Definition Instances of the
	// specified file or folder.
	//
	// Return Values: The method MUST return zero on success or a nonzero error code on
	// failure.
	//
	//	+------------------------------------------+----------------------------------------------------------------------------------+
	//	|                  RETURN                  |                                                                                  |
	//	|                VALUE/CODE                |                                   DESCRIPTION                                    |
	//	|                                          |                                                                                  |
	//	+------------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80045305 FSRM_S_PARTIAL_CLASSIFICATION | The enumerated properties could not be completely classified because a failure   |
	//	|                                          | occurred while loading or classifying the file properties.                       |
	//	+------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80045353 FSRM_E_ENUM_PROPERTIES_FAILED | The properties could not be enumerated because a failure occurred while loading  |
	//	|                                          | or classifying the file properties.                                              |
	//	+------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070057 E_INVALIDARG                  | This code is returned for the following reasons: The options parameter is not a  |
	//	|                                          | valid FsrmGetFilePropertyOptions (section 2.2.2.5.1.2) value. The fileProperties |
	//	|                                          | parameter is NULL. The filePath parameter is NULL.                               |
	//	+------------------------------------------+----------------------------------------------------------------------------------+
	EnumFileProperties(context.Context, *EnumFilePropertiesRequest) (*EnumFilePropertiesResponse, error)

	// The GetFileProperty method is used to get a specific Property Definition Instance
	// from a file or folder.
	//
	// Return Values: The method MUST return zero on success, or a nonzero error code on
	// failure.
	//
	//	+-----------------------------+-----------------------------------------------------------+
	//	|           RETURN            |                                                           |
	//	|         VALUE/CODE          |                        DESCRIPTION                        |
	//	|                             |                                                           |
	//	+-----------------------------+-----------------------------------------------------------+
	//	+-----------------------------+-----------------------------------------------------------+
	//	| 0x80045301 FSRM_E_NOT_FOUND | An object with the specified property name was not found. |
	//	+-----------------------------+-----------------------------------------------------------+
	//	| 0x80070057 E_INVALIDARG     | The property parameter is NULL.                           |
	//	+-----------------------------+-----------------------------------------------------------+
	GetFileProperty(context.Context, *GetFilePropertyRequest) (*GetFilePropertyResponse, error)

	// The SetFileProperty method is used to set a Property Definition Instance on a file
	// or folder.
	//
	// Return Values: The method MUST return zero on success, or a nonzero error code on
	// failure.
	//
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	|                RETURN                 |                                                                                  |
	//	|              VALUE/CODE               |                                   DESCRIPTION                                    |
	//	|                                       |                                                                                  |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80045304 FSRM_E_PATH_NOT_FOUND      | The specified file pointed to by the filePath parameter is not found.            |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80045306 FSRM_E_INVALID_PATH        | The specified path is not valid. A path cannot be a relative path; it must be a  |
	//	|                                       | full, absolute path to a file. A file share path cannot be specified.            |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80045354 FSRM_E_SET_PROPERTY_FAILED | The property could not be set.                                                   |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070057 E_INVALIDARG               | The propertyValue parameter is not a valid for the type of property definition   |
	//	|                                       | specified.                                                                       |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	SetFileProperty(context.Context, *SetFilePropertyRequest) (*SetFilePropertyResponse, error)

	// The ClearFileProperty method is used to clear the specified Property Definition Instance
	// from a file or folder.
	//
	// Return Values: The method MUST return zero on success, or a nonzero error code on
	// failure.
	//
	//	+-------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	|                           RETURN                            |                                                                                  |
	//	|                         VALUE/CODE                          |                                   DESCRIPTION                                    |
	//	|                                                             |                                                                                  |
	//	+-------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	+-------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80045301 FSRM_E_NOT_FOUND                                 | The specified property could not be found.                                       |
	//	+-------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80045354 FSRM_E_SET_PROPERTY_FAILED                       | The property could not be set.                                                   |
	//	+-------------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80045357 FSRM_E_PARTIAL_CLASSIFICATION_PROPERTY_NOT_FOUND | The requested property was not found, but the file might only have partial       |
	//	|                                                             | classification because a failure occurred while loading or classifying the file  |
	//	|                                                             | properties.                                                                      |
	//	+-------------------------------------------------------------+----------------------------------------------------------------------------------+
	ClearFileProperty(context.Context, *ClearFilePropertyRequest) (*ClearFilePropertyResponse, error)
}

func RegisterClassificationManagerServer(conn dcerpc.Conn, o ClassificationManagerServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewClassificationManagerServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(ClassificationManagerSyntaxV0_0))...)
}

func NewClassificationManagerServerHandle(o ClassificationManagerServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return ClassificationManagerServerHandle(ctx, o, opNum, r)
	}
}

func ClassificationManagerServerHandle(ctx context.Context, o ClassificationManagerServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 7 {
		// IDispatch base method.
		return idispatch.DispatchServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 7: // ClassificationReportFormats
		in := &GetClassificationReportFormatsRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetClassificationReportFormats(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 8: // ClassificationReportFormats
		in := &SetClassificationReportFormatsRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetClassificationReportFormats(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 9: // Logging
		in := &GetLoggingRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetLogging(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 10: // Logging
		in := &SetLoggingRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetLogging(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 11: // ClassificationReportMailTo
		in := &GetClassificationReportMailToRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetClassificationReportMailTo(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 12: // ClassificationReportMailTo
		in := &SetClassificationReportMailToRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetClassificationReportMailTo(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 13: // ClassificationReportEnabled
		in := &GetClassificationReportEnabledRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetClassificationReportEnabled(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 14: // ClassificationReportEnabled
		in := &SetClassificationReportEnabledRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetClassificationReportEnabled(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 15: // ClassificationLastReportPathWithoutExtension
		in := &GetClassificationLastReportPathWithoutExtensionRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetClassificationLastReportPathWithoutExtension(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 16: // ClassificationLastError
		in := &GetClassificationLastErrorRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetClassificationLastError(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 17: // ClassificationRunningStatus
		in := &GetClassificationRunningStatusRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetClassificationRunningStatus(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 18: // EnumPropertyDefinitions
		in := &EnumPropertyDefinitionsRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.EnumPropertyDefinitions(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 19: // CreatePropertyDefinition
		in := &CreatePropertyDefinitionRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.CreatePropertyDefinition(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 20: // GetPropertyDefinition
		in := &GetPropertyDefinitionRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetPropertyDefinition(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 21: // EnumRules
		in := &EnumRulesRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.EnumRules(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 22: // CreateRule
		in := &CreateRuleRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.CreateRule(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 23: // GetRule
		in := &GetRuleRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetRule(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 24: // EnumModuleDefinitions
		in := &EnumModuleDefinitionsRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.EnumModuleDefinitions(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 25: // CreateModuleDefinition
		in := &CreateModuleDefinitionRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.CreateModuleDefinition(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 26: // GetModuleDefinition
		in := &GetModuleDefinitionRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetModuleDefinition(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 27: // RunClassification
		in := &RunClassificationRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.RunClassification(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 28: // WaitForClassificationCompletion
		in := &WaitForClassificationCompletionRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.WaitForClassificationCompletion(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 29: // CancelClassification
		in := &CancelClassificationRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.CancelClassification(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 30: // EnumFileProperties
		in := &EnumFilePropertiesRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.EnumFileProperties(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 31: // GetFileProperty
		in := &GetFilePropertyRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetFileProperty(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 32: // SetFileProperty
		in := &SetFilePropertyRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetFileProperty(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 33: // ClearFileProperty
		in := &ClearFilePropertyRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.ClearFileProperty(ctx, in)
		return resp.xxx_ToOp(ctx), err
	}
	return nil, nil
}
