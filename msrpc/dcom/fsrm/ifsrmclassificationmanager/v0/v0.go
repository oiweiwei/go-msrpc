package ifsrmclassificationmanager

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	dcom "github.com/oiweiwei/go-msrpc/msrpc/dcom"
	fsrm "github.com/oiweiwei/go-msrpc/msrpc/dcom/fsrm"
	oaut "github.com/oiweiwei/go-msrpc/msrpc/dcom/oaut"
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
	_ = dcom.GoPackage
	_ = idispatch.GoPackage
	_ = oaut.GoPackage
	_ = fsrm.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/fsrm"
)

var (
	// IFsrmClassificationManager interface identifier d2dc89da-ee91-48a0-85d8-cc72a56f7d04
	ClassificationManagerIID = &dcom.IID{Data1: 0xd2dc89da, Data2: 0xee91, Data3: 0x48a0, Data4: []byte{0x85, 0xd8, 0xcc, 0x72, 0xa5, 0x6f, 0x7d, 0x04}}
	// Syntax UUID
	ClassificationManagerSyntaxUUID = &uuid.UUID{TimeLow: 0xd2dc89da, TimeMid: 0xee91, TimeHiAndVersion: 0x48a0, ClockSeqHiAndReserved: 0x85, ClockSeqLow: 0xd8, Node: [6]uint8{0xcc, 0x72, 0xa5, 0x6f, 0x7d, 0x4}}
	// Syntax ID
	ClassificationManagerSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: ClassificationManagerSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IFsrmClassificationManager interface.
type ClassificationManagerClient interface {

	// IDispatch retrieval method.
	Dispatch() idispatch.DispatchClient

	// ClassificationReportFormats operation.
	GetClassificationReportFormats(context.Context, *GetClassificationReportFormatsRequest, ...dcerpc.CallOption) (*GetClassificationReportFormatsResponse, error)

	// ClassificationReportFormats operation.
	SetClassificationReportFormats(context.Context, *SetClassificationReportFormatsRequest, ...dcerpc.CallOption) (*SetClassificationReportFormatsResponse, error)

	// Logging operation.
	GetLogging(context.Context, *GetLoggingRequest, ...dcerpc.CallOption) (*GetLoggingResponse, error)

	// Logging operation.
	SetLogging(context.Context, *SetLoggingRequest, ...dcerpc.CallOption) (*SetLoggingResponse, error)

	// ClassificationReportMailTo operation.
	GetClassificationReportMailTo(context.Context, *GetClassificationReportMailToRequest, ...dcerpc.CallOption) (*GetClassificationReportMailToResponse, error)

	// ClassificationReportMailTo operation.
	SetClassificationReportMailTo(context.Context, *SetClassificationReportMailToRequest, ...dcerpc.CallOption) (*SetClassificationReportMailToResponse, error)

	// ClassificationReportEnabled operation.
	GetClassificationReportEnabled(context.Context, *GetClassificationReportEnabledRequest, ...dcerpc.CallOption) (*GetClassificationReportEnabledResponse, error)

	// ClassificationReportEnabled operation.
	SetClassificationReportEnabled(context.Context, *SetClassificationReportEnabledRequest, ...dcerpc.CallOption) (*SetClassificationReportEnabledResponse, error)

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
	GetClassificationLastReportPathWithoutExtension(context.Context, *GetClassificationLastReportPathWithoutExtensionRequest, ...dcerpc.CallOption) (*GetClassificationLastReportPathWithoutExtensionResponse, error)

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
	GetClassificationLastError(context.Context, *GetClassificationLastErrorRequest, ...dcerpc.CallOption) (*GetClassificationLastErrorResponse, error)

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
	GetClassificationRunningStatus(context.Context, *GetClassificationRunningStatusRequest, ...dcerpc.CallOption) (*GetClassificationRunningStatusResponse, error)

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
	EnumPropertyDefinitions(context.Context, *EnumPropertyDefinitionsRequest, ...dcerpc.CallOption) (*EnumPropertyDefinitionsResponse, error)

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
	CreatePropertyDefinition(context.Context, *CreatePropertyDefinitionRequest, ...dcerpc.CallOption) (*CreatePropertyDefinitionResponse, error)

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
	GetPropertyDefinition(context.Context, *GetPropertyDefinitionRequest, ...dcerpc.CallOption) (*GetPropertyDefinitionResponse, error)

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
	EnumRules(context.Context, *EnumRulesRequest, ...dcerpc.CallOption) (*EnumRulesResponse, error)

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
	CreateRule(context.Context, *CreateRuleRequest, ...dcerpc.CallOption) (*CreateRuleResponse, error)

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
	GetRule(context.Context, *GetRuleRequest, ...dcerpc.CallOption) (*GetRuleResponse, error)

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
	EnumModuleDefinitions(context.Context, *EnumModuleDefinitionsRequest, ...dcerpc.CallOption) (*EnumModuleDefinitionsResponse, error)

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
	CreateModuleDefinition(context.Context, *CreateModuleDefinitionRequest, ...dcerpc.CallOption) (*CreateModuleDefinitionResponse, error)

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
	GetModuleDefinition(context.Context, *GetModuleDefinitionRequest, ...dcerpc.CallOption) (*GetModuleDefinitionResponse, error)

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
	RunClassification(context.Context, *RunClassificationRequest, ...dcerpc.CallOption) (*RunClassificationResponse, error)

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
	WaitForClassificationCompletion(context.Context, *WaitForClassificationCompletionRequest, ...dcerpc.CallOption) (*WaitForClassificationCompletionResponse, error)

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
	CancelClassification(context.Context, *CancelClassificationRequest, ...dcerpc.CallOption) (*CancelClassificationResponse, error)

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
	EnumFileProperties(context.Context, *EnumFilePropertiesRequest, ...dcerpc.CallOption) (*EnumFilePropertiesResponse, error)

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
	GetFileProperty(context.Context, *GetFilePropertyRequest, ...dcerpc.CallOption) (*GetFilePropertyResponse, error)

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
	SetFileProperty(context.Context, *SetFilePropertyRequest, ...dcerpc.CallOption) (*SetFilePropertyResponse, error)

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
	ClearFileProperty(context.Context, *ClearFilePropertyRequest, ...dcerpc.CallOption) (*ClearFilePropertyResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) ClassificationManagerClient
}

type xxx_DefaultClassificationManagerClient struct {
	idispatch.DispatchClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultClassificationManagerClient) Dispatch() idispatch.DispatchClient {
	return o.DispatchClient
}

func (o *xxx_DefaultClassificationManagerClient) GetClassificationReportFormats(ctx context.Context, in *GetClassificationReportFormatsRequest, opts ...dcerpc.CallOption) (*GetClassificationReportFormatsResponse, error) {
	op := in.xxx_ToOp(ctx)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetClassificationReportFormatsResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClassificationManagerClient) SetClassificationReportFormats(ctx context.Context, in *SetClassificationReportFormatsRequest, opts ...dcerpc.CallOption) (*SetClassificationReportFormatsResponse, error) {
	op := in.xxx_ToOp(ctx)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &SetClassificationReportFormatsResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClassificationManagerClient) GetLogging(ctx context.Context, in *GetLoggingRequest, opts ...dcerpc.CallOption) (*GetLoggingResponse, error) {
	op := in.xxx_ToOp(ctx)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetLoggingResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClassificationManagerClient) SetLogging(ctx context.Context, in *SetLoggingRequest, opts ...dcerpc.CallOption) (*SetLoggingResponse, error) {
	op := in.xxx_ToOp(ctx)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &SetLoggingResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClassificationManagerClient) GetClassificationReportMailTo(ctx context.Context, in *GetClassificationReportMailToRequest, opts ...dcerpc.CallOption) (*GetClassificationReportMailToResponse, error) {
	op := in.xxx_ToOp(ctx)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetClassificationReportMailToResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClassificationManagerClient) SetClassificationReportMailTo(ctx context.Context, in *SetClassificationReportMailToRequest, opts ...dcerpc.CallOption) (*SetClassificationReportMailToResponse, error) {
	op := in.xxx_ToOp(ctx)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &SetClassificationReportMailToResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClassificationManagerClient) GetClassificationReportEnabled(ctx context.Context, in *GetClassificationReportEnabledRequest, opts ...dcerpc.CallOption) (*GetClassificationReportEnabledResponse, error) {
	op := in.xxx_ToOp(ctx)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetClassificationReportEnabledResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClassificationManagerClient) SetClassificationReportEnabled(ctx context.Context, in *SetClassificationReportEnabledRequest, opts ...dcerpc.CallOption) (*SetClassificationReportEnabledResponse, error) {
	op := in.xxx_ToOp(ctx)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &SetClassificationReportEnabledResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClassificationManagerClient) GetClassificationLastReportPathWithoutExtension(ctx context.Context, in *GetClassificationLastReportPathWithoutExtensionRequest, opts ...dcerpc.CallOption) (*GetClassificationLastReportPathWithoutExtensionResponse, error) {
	op := in.xxx_ToOp(ctx)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetClassificationLastReportPathWithoutExtensionResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClassificationManagerClient) GetClassificationLastError(ctx context.Context, in *GetClassificationLastErrorRequest, opts ...dcerpc.CallOption) (*GetClassificationLastErrorResponse, error) {
	op := in.xxx_ToOp(ctx)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetClassificationLastErrorResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClassificationManagerClient) GetClassificationRunningStatus(ctx context.Context, in *GetClassificationRunningStatusRequest, opts ...dcerpc.CallOption) (*GetClassificationRunningStatusResponse, error) {
	op := in.xxx_ToOp(ctx)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetClassificationRunningStatusResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClassificationManagerClient) EnumPropertyDefinitions(ctx context.Context, in *EnumPropertyDefinitionsRequest, opts ...dcerpc.CallOption) (*EnumPropertyDefinitionsResponse, error) {
	op := in.xxx_ToOp(ctx)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &EnumPropertyDefinitionsResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClassificationManagerClient) CreatePropertyDefinition(ctx context.Context, in *CreatePropertyDefinitionRequest, opts ...dcerpc.CallOption) (*CreatePropertyDefinitionResponse, error) {
	op := in.xxx_ToOp(ctx)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &CreatePropertyDefinitionResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClassificationManagerClient) GetPropertyDefinition(ctx context.Context, in *GetPropertyDefinitionRequest, opts ...dcerpc.CallOption) (*GetPropertyDefinitionResponse, error) {
	op := in.xxx_ToOp(ctx)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetPropertyDefinitionResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClassificationManagerClient) EnumRules(ctx context.Context, in *EnumRulesRequest, opts ...dcerpc.CallOption) (*EnumRulesResponse, error) {
	op := in.xxx_ToOp(ctx)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &EnumRulesResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClassificationManagerClient) CreateRule(ctx context.Context, in *CreateRuleRequest, opts ...dcerpc.CallOption) (*CreateRuleResponse, error) {
	op := in.xxx_ToOp(ctx)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &CreateRuleResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClassificationManagerClient) GetRule(ctx context.Context, in *GetRuleRequest, opts ...dcerpc.CallOption) (*GetRuleResponse, error) {
	op := in.xxx_ToOp(ctx)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetRuleResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClassificationManagerClient) EnumModuleDefinitions(ctx context.Context, in *EnumModuleDefinitionsRequest, opts ...dcerpc.CallOption) (*EnumModuleDefinitionsResponse, error) {
	op := in.xxx_ToOp(ctx)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &EnumModuleDefinitionsResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClassificationManagerClient) CreateModuleDefinition(ctx context.Context, in *CreateModuleDefinitionRequest, opts ...dcerpc.CallOption) (*CreateModuleDefinitionResponse, error) {
	op := in.xxx_ToOp(ctx)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &CreateModuleDefinitionResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClassificationManagerClient) GetModuleDefinition(ctx context.Context, in *GetModuleDefinitionRequest, opts ...dcerpc.CallOption) (*GetModuleDefinitionResponse, error) {
	op := in.xxx_ToOp(ctx)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetModuleDefinitionResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClassificationManagerClient) RunClassification(ctx context.Context, in *RunClassificationRequest, opts ...dcerpc.CallOption) (*RunClassificationResponse, error) {
	op := in.xxx_ToOp(ctx)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &RunClassificationResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClassificationManagerClient) WaitForClassificationCompletion(ctx context.Context, in *WaitForClassificationCompletionRequest, opts ...dcerpc.CallOption) (*WaitForClassificationCompletionResponse, error) {
	op := in.xxx_ToOp(ctx)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &WaitForClassificationCompletionResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClassificationManagerClient) CancelClassification(ctx context.Context, in *CancelClassificationRequest, opts ...dcerpc.CallOption) (*CancelClassificationResponse, error) {
	op := in.xxx_ToOp(ctx)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &CancelClassificationResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClassificationManagerClient) EnumFileProperties(ctx context.Context, in *EnumFilePropertiesRequest, opts ...dcerpc.CallOption) (*EnumFilePropertiesResponse, error) {
	op := in.xxx_ToOp(ctx)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &EnumFilePropertiesResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClassificationManagerClient) GetFileProperty(ctx context.Context, in *GetFilePropertyRequest, opts ...dcerpc.CallOption) (*GetFilePropertyResponse, error) {
	op := in.xxx_ToOp(ctx)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetFilePropertyResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClassificationManagerClient) SetFileProperty(ctx context.Context, in *SetFilePropertyRequest, opts ...dcerpc.CallOption) (*SetFilePropertyResponse, error) {
	op := in.xxx_ToOp(ctx)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &SetFilePropertyResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClassificationManagerClient) ClearFileProperty(ctx context.Context, in *ClearFilePropertyRequest, opts ...dcerpc.CallOption) (*ClearFilePropertyResponse, error) {
	op := in.xxx_ToOp(ctx)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &ClearFilePropertyResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClassificationManagerClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultClassificationManagerClient) IPID(ctx context.Context, ipid *dcom.IPID) ClassificationManagerClient {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultClassificationManagerClient{
		DispatchClient: o.DispatchClient.IPID(ctx, ipid),
		cc:             o.cc,
		ipid:           ipid,
	}
}
func NewClassificationManagerClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (ClassificationManagerClient, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(ClassificationManagerSyntaxV0_0))...)
		if err != nil {
			return nil, err
		}
	}
	base, err := idispatch.NewDispatchClient(ctx, cc, append(opts, dcom.Superclass(cc))...)
	if err != nil {
		return nil, err
	}
	ipid, ok := dcom.HasIPID(opts)
	if ok {
		base = base.IPID(ctx, ipid)
	}
	return &xxx_DefaultClassificationManagerClient{
		DispatchClient: base,
		cc:             cc,
		ipid:           ipid,
	}, nil
}

// xxx_GetClassificationReportFormatsOperation structure represents the ClassificationReportFormats operation
type xxx_GetClassificationReportFormatsOperation struct {
	This    *dcom.ORPCThis  `idl:"name:This" json:"this"`
	That    *dcom.ORPCThat  `idl:"name:That" json:"that"`
	Formats *oaut.SafeArray `idl:"name:formats" json:"formats"`
	Return  int32           `idl:"name:Return" json:"return"`
}

func (o *xxx_GetClassificationReportFormatsOperation) OpNum() int { return 7 }

func (o *xxx_GetClassificationReportFormatsOperation) OpName() string {
	return "/IFsrmClassificationManager/v0/ClassificationReportFormats"
}

func (o *xxx_GetClassificationReportFormatsOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetClassificationReportFormatsOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetClassificationReportFormatsOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetClassificationReportFormatsOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetClassificationReportFormatsOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// formats {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=SAFEARRAY}*(1))(3:{pointer=unique, alias=_SAFEARRAY}(struct))
	{
		if o.Formats != nil {
			_ptr_formats := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Formats != nil {
					if err := o.Formats.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.SafeArray{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Formats, _ptr_formats); err != nil {
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
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetClassificationReportFormatsOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// formats {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=SAFEARRAY}*(1))(3:{pointer=unique, alias=_SAFEARRAY}(struct))
	{
		_ptr_formats := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Formats == nil {
				o.Formats = &oaut.SafeArray{}
			}
			if err := o.Formats.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_formats := func(ptr interface{}) { o.Formats = *ptr.(**oaut.SafeArray) }
		if err := w.ReadPointer(&o.Formats, _s_formats, _ptr_formats); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// GetClassificationReportFormatsRequest structure represents the ClassificationReportFormats operation request
type GetClassificationReportFormatsRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetClassificationReportFormatsRequest) xxx_ToOp(ctx context.Context) *xxx_GetClassificationReportFormatsOperation {
	if o == nil {
		return &xxx_GetClassificationReportFormatsOperation{}
	}
	return &xxx_GetClassificationReportFormatsOperation{
		This: o.This,
	}
}

func (o *GetClassificationReportFormatsRequest) xxx_FromOp(ctx context.Context, op *xxx_GetClassificationReportFormatsOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetClassificationReportFormatsRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetClassificationReportFormatsRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetClassificationReportFormatsOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetClassificationReportFormatsResponse structure represents the ClassificationReportFormats operation response
type GetClassificationReportFormatsResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That    *dcom.ORPCThat  `idl:"name:That" json:"that"`
	Formats *oaut.SafeArray `idl:"name:formats" json:"formats"`
	// Return: The ClassificationReportFormats return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetClassificationReportFormatsResponse) xxx_ToOp(ctx context.Context) *xxx_GetClassificationReportFormatsOperation {
	if o == nil {
		return &xxx_GetClassificationReportFormatsOperation{}
	}
	return &xxx_GetClassificationReportFormatsOperation{
		That:    o.That,
		Formats: o.Formats,
		Return:  o.Return,
	}
}

func (o *GetClassificationReportFormatsResponse) xxx_FromOp(ctx context.Context, op *xxx_GetClassificationReportFormatsOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Formats = op.Formats
	o.Return = op.Return
}
func (o *GetClassificationReportFormatsResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetClassificationReportFormatsResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetClassificationReportFormatsOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetClassificationReportFormatsOperation structure represents the ClassificationReportFormats operation
type xxx_SetClassificationReportFormatsOperation struct {
	This    *dcom.ORPCThis  `idl:"name:This" json:"this"`
	That    *dcom.ORPCThat  `idl:"name:That" json:"that"`
	Formats *oaut.SafeArray `idl:"name:formats" json:"formats"`
	Return  int32           `idl:"name:Return" json:"return"`
}

func (o *xxx_SetClassificationReportFormatsOperation) OpNum() int { return 8 }

func (o *xxx_SetClassificationReportFormatsOperation) OpName() string {
	return "/IFsrmClassificationManager/v0/ClassificationReportFormats"
}

func (o *xxx_SetClassificationReportFormatsOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetClassificationReportFormatsOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// formats {in} (1:{pointer=unique, alias=SAFEARRAY}*(1))(2:{pointer=unique, alias=_SAFEARRAY}(struct))
	{
		if o.Formats != nil {
			_ptr_formats := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Formats != nil {
					if err := o.Formats.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.SafeArray{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Formats, _ptr_formats); err != nil {
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

func (o *xxx_SetClassificationReportFormatsOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// formats {in} (1:{pointer=unique, alias=SAFEARRAY}*(1))(2:{pointer=unique, alias=_SAFEARRAY}(struct))
	{
		_ptr_formats := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Formats == nil {
				o.Formats = &oaut.SafeArray{}
			}
			if err := o.Formats.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_formats := func(ptr interface{}) { o.Formats = *ptr.(**oaut.SafeArray) }
		if err := w.ReadPointer(&o.Formats, _s_formats, _ptr_formats); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetClassificationReportFormatsOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetClassificationReportFormatsOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetClassificationReportFormatsOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// SetClassificationReportFormatsRequest structure represents the ClassificationReportFormats operation request
type SetClassificationReportFormatsRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This    *dcom.ORPCThis  `idl:"name:This" json:"this"`
	Formats *oaut.SafeArray `idl:"name:formats" json:"formats"`
}

func (o *SetClassificationReportFormatsRequest) xxx_ToOp(ctx context.Context) *xxx_SetClassificationReportFormatsOperation {
	if o == nil {
		return &xxx_SetClassificationReportFormatsOperation{}
	}
	return &xxx_SetClassificationReportFormatsOperation{
		This:    o.This,
		Formats: o.Formats,
	}
}

func (o *SetClassificationReportFormatsRequest) xxx_FromOp(ctx context.Context, op *xxx_SetClassificationReportFormatsOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Formats = op.Formats
}
func (o *SetClassificationReportFormatsRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *SetClassificationReportFormatsRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetClassificationReportFormatsOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetClassificationReportFormatsResponse structure represents the ClassificationReportFormats operation response
type SetClassificationReportFormatsResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The ClassificationReportFormats return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetClassificationReportFormatsResponse) xxx_ToOp(ctx context.Context) *xxx_SetClassificationReportFormatsOperation {
	if o == nil {
		return &xxx_SetClassificationReportFormatsOperation{}
	}
	return &xxx_SetClassificationReportFormatsOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *SetClassificationReportFormatsResponse) xxx_FromOp(ctx context.Context, op *xxx_SetClassificationReportFormatsOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetClassificationReportFormatsResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *SetClassificationReportFormatsResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetClassificationReportFormatsOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetLoggingOperation structure represents the Logging operation
type xxx_GetLoggingOperation struct {
	This    *dcom.ORPCThis `idl:"name:This" json:"this"`
	That    *dcom.ORPCThat `idl:"name:That" json:"that"`
	Logging int32          `idl:"name:logging" json:"logging"`
	Return  int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetLoggingOperation) OpNum() int { return 9 }

func (o *xxx_GetLoggingOperation) OpName() string { return "/IFsrmClassificationManager/v0/Logging" }

func (o *xxx_GetLoggingOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetLoggingOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetLoggingOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetLoggingOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetLoggingOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// logging {out, retval} (1:{pointer=ref}*(1)(int32))
	{
		if err := w.WriteData(o.Logging); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetLoggingOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// logging {out, retval} (1:{pointer=ref}*(1)(int32))
	{
		if err := w.ReadData(&o.Logging); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// GetLoggingRequest structure represents the Logging operation request
type GetLoggingRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetLoggingRequest) xxx_ToOp(ctx context.Context) *xxx_GetLoggingOperation {
	if o == nil {
		return &xxx_GetLoggingOperation{}
	}
	return &xxx_GetLoggingOperation{
		This: o.This,
	}
}

func (o *GetLoggingRequest) xxx_FromOp(ctx context.Context, op *xxx_GetLoggingOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetLoggingRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetLoggingRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetLoggingOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetLoggingResponse structure represents the Logging operation response
type GetLoggingResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That    *dcom.ORPCThat `idl:"name:That" json:"that"`
	Logging int32          `idl:"name:logging" json:"logging"`
	// Return: The Logging return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetLoggingResponse) xxx_ToOp(ctx context.Context) *xxx_GetLoggingOperation {
	if o == nil {
		return &xxx_GetLoggingOperation{}
	}
	return &xxx_GetLoggingOperation{
		That:    o.That,
		Logging: o.Logging,
		Return:  o.Return,
	}
}

func (o *GetLoggingResponse) xxx_FromOp(ctx context.Context, op *xxx_GetLoggingOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Logging = op.Logging
	o.Return = op.Return
}
func (o *GetLoggingResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetLoggingResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetLoggingOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetLoggingOperation structure represents the Logging operation
type xxx_SetLoggingOperation struct {
	This    *dcom.ORPCThis `idl:"name:This" json:"this"`
	That    *dcom.ORPCThat `idl:"name:That" json:"that"`
	Logging int32          `idl:"name:logging" json:"logging"`
	Return  int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetLoggingOperation) OpNum() int { return 10 }

func (o *xxx_SetLoggingOperation) OpName() string { return "/IFsrmClassificationManager/v0/Logging" }

func (o *xxx_SetLoggingOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetLoggingOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// logging {in} (1:(int32))
	{
		if err := w.WriteData(o.Logging); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetLoggingOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// logging {in} (1:(int32))
	{
		if err := w.ReadData(&o.Logging); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetLoggingOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetLoggingOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetLoggingOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// SetLoggingRequest structure represents the Logging operation request
type SetLoggingRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This    *dcom.ORPCThis `idl:"name:This" json:"this"`
	Logging int32          `idl:"name:logging" json:"logging"`
}

func (o *SetLoggingRequest) xxx_ToOp(ctx context.Context) *xxx_SetLoggingOperation {
	if o == nil {
		return &xxx_SetLoggingOperation{}
	}
	return &xxx_SetLoggingOperation{
		This:    o.This,
		Logging: o.Logging,
	}
}

func (o *SetLoggingRequest) xxx_FromOp(ctx context.Context, op *xxx_SetLoggingOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Logging = op.Logging
}
func (o *SetLoggingRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *SetLoggingRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetLoggingOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetLoggingResponse structure represents the Logging operation response
type SetLoggingResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The Logging return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetLoggingResponse) xxx_ToOp(ctx context.Context) *xxx_SetLoggingOperation {
	if o == nil {
		return &xxx_SetLoggingOperation{}
	}
	return &xxx_SetLoggingOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *SetLoggingResponse) xxx_FromOp(ctx context.Context, op *xxx_SetLoggingOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetLoggingResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *SetLoggingResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetLoggingOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetClassificationReportMailToOperation structure represents the ClassificationReportMailTo operation
type xxx_GetClassificationReportMailToOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	MailTo *oaut.String   `idl:"name:mailTo" json:"mail_to"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetClassificationReportMailToOperation) OpNum() int { return 11 }

func (o *xxx_GetClassificationReportMailToOperation) OpName() string {
	return "/IFsrmClassificationManager/v0/ClassificationReportMailTo"
}

func (o *xxx_GetClassificationReportMailToOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetClassificationReportMailToOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetClassificationReportMailToOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetClassificationReportMailToOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetClassificationReportMailToOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// mailTo {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.MailTo != nil {
			_ptr_mailTo := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.MailTo != nil {
					if err := o.MailTo.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.MailTo, _ptr_mailTo); err != nil {
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
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetClassificationReportMailToOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// mailTo {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_mailTo := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.MailTo == nil {
				o.MailTo = &oaut.String{}
			}
			if err := o.MailTo.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_mailTo := func(ptr interface{}) { o.MailTo = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.MailTo, _s_mailTo, _ptr_mailTo); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// GetClassificationReportMailToRequest structure represents the ClassificationReportMailTo operation request
type GetClassificationReportMailToRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetClassificationReportMailToRequest) xxx_ToOp(ctx context.Context) *xxx_GetClassificationReportMailToOperation {
	if o == nil {
		return &xxx_GetClassificationReportMailToOperation{}
	}
	return &xxx_GetClassificationReportMailToOperation{
		This: o.This,
	}
}

func (o *GetClassificationReportMailToRequest) xxx_FromOp(ctx context.Context, op *xxx_GetClassificationReportMailToOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetClassificationReportMailToRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetClassificationReportMailToRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetClassificationReportMailToOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetClassificationReportMailToResponse structure represents the ClassificationReportMailTo operation response
type GetClassificationReportMailToResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	MailTo *oaut.String   `idl:"name:mailTo" json:"mail_to"`
	// Return: The ClassificationReportMailTo return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetClassificationReportMailToResponse) xxx_ToOp(ctx context.Context) *xxx_GetClassificationReportMailToOperation {
	if o == nil {
		return &xxx_GetClassificationReportMailToOperation{}
	}
	return &xxx_GetClassificationReportMailToOperation{
		That:   o.That,
		MailTo: o.MailTo,
		Return: o.Return,
	}
}

func (o *GetClassificationReportMailToResponse) xxx_FromOp(ctx context.Context, op *xxx_GetClassificationReportMailToOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.MailTo = op.MailTo
	o.Return = op.Return
}
func (o *GetClassificationReportMailToResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetClassificationReportMailToResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetClassificationReportMailToOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetClassificationReportMailToOperation structure represents the ClassificationReportMailTo operation
type xxx_SetClassificationReportMailToOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	MailTo *oaut.String   `idl:"name:mailTo" json:"mail_to"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetClassificationReportMailToOperation) OpNum() int { return 12 }

func (o *xxx_SetClassificationReportMailToOperation) OpName() string {
	return "/IFsrmClassificationManager/v0/ClassificationReportMailTo"
}

func (o *xxx_SetClassificationReportMailToOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetClassificationReportMailToOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// mailTo {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.MailTo != nil {
			_ptr_mailTo := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.MailTo != nil {
					if err := o.MailTo.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.MailTo, _ptr_mailTo); err != nil {
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

func (o *xxx_SetClassificationReportMailToOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// mailTo {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_mailTo := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.MailTo == nil {
				o.MailTo = &oaut.String{}
			}
			if err := o.MailTo.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_mailTo := func(ptr interface{}) { o.MailTo = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.MailTo, _s_mailTo, _ptr_mailTo); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetClassificationReportMailToOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetClassificationReportMailToOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetClassificationReportMailToOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// SetClassificationReportMailToRequest structure represents the ClassificationReportMailTo operation request
type SetClassificationReportMailToRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	MailTo *oaut.String   `idl:"name:mailTo" json:"mail_to"`
}

func (o *SetClassificationReportMailToRequest) xxx_ToOp(ctx context.Context) *xxx_SetClassificationReportMailToOperation {
	if o == nil {
		return &xxx_SetClassificationReportMailToOperation{}
	}
	return &xxx_SetClassificationReportMailToOperation{
		This:   o.This,
		MailTo: o.MailTo,
	}
}

func (o *SetClassificationReportMailToRequest) xxx_FromOp(ctx context.Context, op *xxx_SetClassificationReportMailToOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.MailTo = op.MailTo
}
func (o *SetClassificationReportMailToRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *SetClassificationReportMailToRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetClassificationReportMailToOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetClassificationReportMailToResponse structure represents the ClassificationReportMailTo operation response
type SetClassificationReportMailToResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The ClassificationReportMailTo return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetClassificationReportMailToResponse) xxx_ToOp(ctx context.Context) *xxx_SetClassificationReportMailToOperation {
	if o == nil {
		return &xxx_SetClassificationReportMailToOperation{}
	}
	return &xxx_SetClassificationReportMailToOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *SetClassificationReportMailToResponse) xxx_FromOp(ctx context.Context, op *xxx_SetClassificationReportMailToOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetClassificationReportMailToResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *SetClassificationReportMailToResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetClassificationReportMailToOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetClassificationReportEnabledOperation structure represents the ClassificationReportEnabled operation
type xxx_GetClassificationReportEnabledOperation struct {
	This          *dcom.ORPCThis `idl:"name:This" json:"this"`
	That          *dcom.ORPCThat `idl:"name:That" json:"that"`
	ReportEnabled int16          `idl:"name:reportEnabled" json:"report_enabled"`
	Return        int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetClassificationReportEnabledOperation) OpNum() int { return 13 }

func (o *xxx_GetClassificationReportEnabledOperation) OpName() string {
	return "/IFsrmClassificationManager/v0/ClassificationReportEnabled"
}

func (o *xxx_GetClassificationReportEnabledOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetClassificationReportEnabledOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetClassificationReportEnabledOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetClassificationReportEnabledOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetClassificationReportEnabledOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// reportEnabled {out, retval} (1:{pointer=ref}*(1))(2:{alias=VARIANT_BOOL}(int16))
	{
		if err := w.WriteData(o.ReportEnabled); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetClassificationReportEnabledOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// reportEnabled {out, retval} (1:{pointer=ref}*(1))(2:{alias=VARIANT_BOOL}(int16))
	{
		if err := w.ReadData(&o.ReportEnabled); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// GetClassificationReportEnabledRequest structure represents the ClassificationReportEnabled operation request
type GetClassificationReportEnabledRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetClassificationReportEnabledRequest) xxx_ToOp(ctx context.Context) *xxx_GetClassificationReportEnabledOperation {
	if o == nil {
		return &xxx_GetClassificationReportEnabledOperation{}
	}
	return &xxx_GetClassificationReportEnabledOperation{
		This: o.This,
	}
}

func (o *GetClassificationReportEnabledRequest) xxx_FromOp(ctx context.Context, op *xxx_GetClassificationReportEnabledOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetClassificationReportEnabledRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetClassificationReportEnabledRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetClassificationReportEnabledOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetClassificationReportEnabledResponse structure represents the ClassificationReportEnabled operation response
type GetClassificationReportEnabledResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That          *dcom.ORPCThat `idl:"name:That" json:"that"`
	ReportEnabled int16          `idl:"name:reportEnabled" json:"report_enabled"`
	// Return: The ClassificationReportEnabled return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetClassificationReportEnabledResponse) xxx_ToOp(ctx context.Context) *xxx_GetClassificationReportEnabledOperation {
	if o == nil {
		return &xxx_GetClassificationReportEnabledOperation{}
	}
	return &xxx_GetClassificationReportEnabledOperation{
		That:          o.That,
		ReportEnabled: o.ReportEnabled,
		Return:        o.Return,
	}
}

func (o *GetClassificationReportEnabledResponse) xxx_FromOp(ctx context.Context, op *xxx_GetClassificationReportEnabledOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.ReportEnabled = op.ReportEnabled
	o.Return = op.Return
}
func (o *GetClassificationReportEnabledResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetClassificationReportEnabledResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetClassificationReportEnabledOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetClassificationReportEnabledOperation structure represents the ClassificationReportEnabled operation
type xxx_SetClassificationReportEnabledOperation struct {
	This          *dcom.ORPCThis `idl:"name:This" json:"this"`
	That          *dcom.ORPCThat `idl:"name:That" json:"that"`
	ReportEnabled int16          `idl:"name:reportEnabled" json:"report_enabled"`
	Return        int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetClassificationReportEnabledOperation) OpNum() int { return 14 }

func (o *xxx_SetClassificationReportEnabledOperation) OpName() string {
	return "/IFsrmClassificationManager/v0/ClassificationReportEnabled"
}

func (o *xxx_SetClassificationReportEnabledOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetClassificationReportEnabledOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// reportEnabled {in} (1:{alias=VARIANT_BOOL}(int16))
	{
		if err := w.WriteData(o.ReportEnabled); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetClassificationReportEnabledOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// reportEnabled {in} (1:{alias=VARIANT_BOOL}(int16))
	{
		if err := w.ReadData(&o.ReportEnabled); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetClassificationReportEnabledOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetClassificationReportEnabledOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetClassificationReportEnabledOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// SetClassificationReportEnabledRequest structure represents the ClassificationReportEnabled operation request
type SetClassificationReportEnabledRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This          *dcom.ORPCThis `idl:"name:This" json:"this"`
	ReportEnabled int16          `idl:"name:reportEnabled" json:"report_enabled"`
}

func (o *SetClassificationReportEnabledRequest) xxx_ToOp(ctx context.Context) *xxx_SetClassificationReportEnabledOperation {
	if o == nil {
		return &xxx_SetClassificationReportEnabledOperation{}
	}
	return &xxx_SetClassificationReportEnabledOperation{
		This:          o.This,
		ReportEnabled: o.ReportEnabled,
	}
}

func (o *SetClassificationReportEnabledRequest) xxx_FromOp(ctx context.Context, op *xxx_SetClassificationReportEnabledOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.ReportEnabled = op.ReportEnabled
}
func (o *SetClassificationReportEnabledRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *SetClassificationReportEnabledRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetClassificationReportEnabledOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetClassificationReportEnabledResponse structure represents the ClassificationReportEnabled operation response
type SetClassificationReportEnabledResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The ClassificationReportEnabled return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetClassificationReportEnabledResponse) xxx_ToOp(ctx context.Context) *xxx_SetClassificationReportEnabledOperation {
	if o == nil {
		return &xxx_SetClassificationReportEnabledOperation{}
	}
	return &xxx_SetClassificationReportEnabledOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *SetClassificationReportEnabledResponse) xxx_FromOp(ctx context.Context, op *xxx_SetClassificationReportEnabledOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetClassificationReportEnabledResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *SetClassificationReportEnabledResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetClassificationReportEnabledOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetClassificationLastReportPathWithoutExtensionOperation structure represents the ClassificationLastReportPathWithoutExtension operation
type xxx_GetClassificationLastReportPathWithoutExtensionOperation struct {
	This           *dcom.ORPCThis `idl:"name:This" json:"this"`
	That           *dcom.ORPCThat `idl:"name:That" json:"that"`
	LastReportPath *oaut.String   `idl:"name:lastReportPath" json:"last_report_path"`
	Return         int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetClassificationLastReportPathWithoutExtensionOperation) OpNum() int { return 15 }

func (o *xxx_GetClassificationLastReportPathWithoutExtensionOperation) OpName() string {
	return "/IFsrmClassificationManager/v0/ClassificationLastReportPathWithoutExtension"
}

func (o *xxx_GetClassificationLastReportPathWithoutExtensionOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetClassificationLastReportPathWithoutExtensionOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetClassificationLastReportPathWithoutExtensionOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetClassificationLastReportPathWithoutExtensionOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetClassificationLastReportPathWithoutExtensionOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// lastReportPath {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.LastReportPath != nil {
			_ptr_lastReportPath := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.LastReportPath != nil {
					if err := o.LastReportPath.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.LastReportPath, _ptr_lastReportPath); err != nil {
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
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetClassificationLastReportPathWithoutExtensionOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// lastReportPath {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_lastReportPath := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.LastReportPath == nil {
				o.LastReportPath = &oaut.String{}
			}
			if err := o.LastReportPath.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_lastReportPath := func(ptr interface{}) { o.LastReportPath = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.LastReportPath, _s_lastReportPath, _ptr_lastReportPath); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// GetClassificationLastReportPathWithoutExtensionRequest structure represents the ClassificationLastReportPathWithoutExtension operation request
type GetClassificationLastReportPathWithoutExtensionRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetClassificationLastReportPathWithoutExtensionRequest) xxx_ToOp(ctx context.Context) *xxx_GetClassificationLastReportPathWithoutExtensionOperation {
	if o == nil {
		return &xxx_GetClassificationLastReportPathWithoutExtensionOperation{}
	}
	return &xxx_GetClassificationLastReportPathWithoutExtensionOperation{
		This: o.This,
	}
}

func (o *GetClassificationLastReportPathWithoutExtensionRequest) xxx_FromOp(ctx context.Context, op *xxx_GetClassificationLastReportPathWithoutExtensionOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetClassificationLastReportPathWithoutExtensionRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetClassificationLastReportPathWithoutExtensionRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetClassificationLastReportPathWithoutExtensionOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetClassificationLastReportPathWithoutExtensionResponse structure represents the ClassificationLastReportPathWithoutExtension operation response
type GetClassificationLastReportPathWithoutExtensionResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// lastReportPath: Pointer to a variable that upon completion contains the path and
	// file name where the generated report(s) was (were) stored when classification was
	// previously run.
	LastReportPath *oaut.String `idl:"name:lastReportPath" json:"last_report_path"`
	// Return: The ClassificationLastReportPathWithoutExtension return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetClassificationLastReportPathWithoutExtensionResponse) xxx_ToOp(ctx context.Context) *xxx_GetClassificationLastReportPathWithoutExtensionOperation {
	if o == nil {
		return &xxx_GetClassificationLastReportPathWithoutExtensionOperation{}
	}
	return &xxx_GetClassificationLastReportPathWithoutExtensionOperation{
		That:           o.That,
		LastReportPath: o.LastReportPath,
		Return:         o.Return,
	}
}

func (o *GetClassificationLastReportPathWithoutExtensionResponse) xxx_FromOp(ctx context.Context, op *xxx_GetClassificationLastReportPathWithoutExtensionOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.LastReportPath = op.LastReportPath
	o.Return = op.Return
}
func (o *GetClassificationLastReportPathWithoutExtensionResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetClassificationLastReportPathWithoutExtensionResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetClassificationLastReportPathWithoutExtensionOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetClassificationLastErrorOperation structure represents the ClassificationLastError operation
type xxx_GetClassificationLastErrorOperation struct {
	This      *dcom.ORPCThis `idl:"name:This" json:"this"`
	That      *dcom.ORPCThat `idl:"name:That" json:"that"`
	LastError *oaut.String   `idl:"name:lastError" json:"last_error"`
	Return    int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetClassificationLastErrorOperation) OpNum() int { return 16 }

func (o *xxx_GetClassificationLastErrorOperation) OpName() string {
	return "/IFsrmClassificationManager/v0/ClassificationLastError"
}

func (o *xxx_GetClassificationLastErrorOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetClassificationLastErrorOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetClassificationLastErrorOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetClassificationLastErrorOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetClassificationLastErrorOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// lastError {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.LastError != nil {
			_ptr_lastError := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.LastError != nil {
					if err := o.LastError.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.LastError, _ptr_lastError); err != nil {
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
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetClassificationLastErrorOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// lastError {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_lastError := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.LastError == nil {
				o.LastError = &oaut.String{}
			}
			if err := o.LastError.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_lastError := func(ptr interface{}) { o.LastError = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.LastError, _s_lastError, _ptr_lastError); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// GetClassificationLastErrorRequest structure represents the ClassificationLastError operation request
type GetClassificationLastErrorRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetClassificationLastErrorRequest) xxx_ToOp(ctx context.Context) *xxx_GetClassificationLastErrorOperation {
	if o == nil {
		return &xxx_GetClassificationLastErrorOperation{}
	}
	return &xxx_GetClassificationLastErrorOperation{
		This: o.This,
	}
}

func (o *GetClassificationLastErrorRequest) xxx_FromOp(ctx context.Context, op *xxx_GetClassificationLastErrorOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetClassificationLastErrorRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetClassificationLastErrorRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetClassificationLastErrorOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetClassificationLastErrorResponse structure represents the ClassificationLastError operation response
type GetClassificationLastErrorResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// lastError: Pointer to a variable that upon completion contains the last error, if
	// any, from when classification was previously run.
	LastError *oaut.String `idl:"name:lastError" json:"last_error"`
	// Return: The ClassificationLastError return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetClassificationLastErrorResponse) xxx_ToOp(ctx context.Context) *xxx_GetClassificationLastErrorOperation {
	if o == nil {
		return &xxx_GetClassificationLastErrorOperation{}
	}
	return &xxx_GetClassificationLastErrorOperation{
		That:      o.That,
		LastError: o.LastError,
		Return:    o.Return,
	}
}

func (o *GetClassificationLastErrorResponse) xxx_FromOp(ctx context.Context, op *xxx_GetClassificationLastErrorOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.LastError = op.LastError
	o.Return = op.Return
}
func (o *GetClassificationLastErrorResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetClassificationLastErrorResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetClassificationLastErrorOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetClassificationRunningStatusOperation structure represents the ClassificationRunningStatus operation
type xxx_GetClassificationRunningStatusOperation struct {
	This          *dcom.ORPCThis           `idl:"name:This" json:"this"`
	That          *dcom.ORPCThat           `idl:"name:That" json:"that"`
	RunningStatus fsrm.ReportRunningStatus `idl:"name:runningStatus" json:"running_status"`
	Return        int32                    `idl:"name:Return" json:"return"`
}

func (o *xxx_GetClassificationRunningStatusOperation) OpNum() int { return 17 }

func (o *xxx_GetClassificationRunningStatusOperation) OpName() string {
	return "/IFsrmClassificationManager/v0/ClassificationRunningStatus"
}

func (o *xxx_GetClassificationRunningStatusOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetClassificationRunningStatusOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetClassificationRunningStatusOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetClassificationRunningStatusOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetClassificationRunningStatusOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// runningStatus {out, retval} (1:{pointer=ref}*(1))(2:{alias=FsrmReportRunningStatus}(enum))
	{
		if err := w.WriteData(uint16(o.RunningStatus)); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetClassificationRunningStatusOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// runningStatus {out, retval} (1:{pointer=ref}*(1))(2:{alias=FsrmReportRunningStatus}(enum))
	{
		if err := w.ReadData((*uint16)(&o.RunningStatus)); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// GetClassificationRunningStatusRequest structure represents the ClassificationRunningStatus operation request
type GetClassificationRunningStatusRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetClassificationRunningStatusRequest) xxx_ToOp(ctx context.Context) *xxx_GetClassificationRunningStatusOperation {
	if o == nil {
		return &xxx_GetClassificationRunningStatusOperation{}
	}
	return &xxx_GetClassificationRunningStatusOperation{
		This: o.This,
	}
}

func (o *GetClassificationRunningStatusRequest) xxx_FromOp(ctx context.Context, op *xxx_GetClassificationRunningStatusOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetClassificationRunningStatusRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetClassificationRunningStatusRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetClassificationRunningStatusOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetClassificationRunningStatusResponse structure represents the ClassificationRunningStatus operation response
type GetClassificationRunningStatusResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// runningStatus:  A pointer to a variable that upon completion contains the current
	// running status of the running classification task.
	RunningStatus fsrm.ReportRunningStatus `idl:"name:runningStatus" json:"running_status"`
	// Return: The ClassificationRunningStatus return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetClassificationRunningStatusResponse) xxx_ToOp(ctx context.Context) *xxx_GetClassificationRunningStatusOperation {
	if o == nil {
		return &xxx_GetClassificationRunningStatusOperation{}
	}
	return &xxx_GetClassificationRunningStatusOperation{
		That:          o.That,
		RunningStatus: o.RunningStatus,
		Return:        o.Return,
	}
}

func (o *GetClassificationRunningStatusResponse) xxx_FromOp(ctx context.Context, op *xxx_GetClassificationRunningStatusOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.RunningStatus = op.RunningStatus
	o.Return = op.Return
}
func (o *GetClassificationRunningStatusResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetClassificationRunningStatusResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetClassificationRunningStatusOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_EnumPropertyDefinitionsOperation structure represents the EnumPropertyDefinitions operation
type xxx_EnumPropertyDefinitionsOperation struct {
	This                *dcom.ORPCThis   `idl:"name:This" json:"this"`
	That                *dcom.ORPCThat   `idl:"name:That" json:"that"`
	Options             fsrm.EnumOptions `idl:"name:options" json:"options"`
	PropertyDefinitions *fsrm.Collection `idl:"name:propertyDefinitions" json:"property_definitions"`
	Return              int32            `idl:"name:Return" json:"return"`
}

func (o *xxx_EnumPropertyDefinitionsOperation) OpNum() int { return 18 }

func (o *xxx_EnumPropertyDefinitionsOperation) OpName() string {
	return "/IFsrmClassificationManager/v0/EnumPropertyDefinitions"
}

func (o *xxx_EnumPropertyDefinitionsOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumPropertyDefinitionsOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// options {in, default_value={0}} (1:{alias=FsrmEnumOptions}(enum))
	{
		if err := w.WriteData(uint16(o.Options)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumPropertyDefinitionsOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// options {in, default_value={0}} (1:{alias=FsrmEnumOptions}(enum))
	{
		if err := w.ReadData((*uint16)(&o.Options)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumPropertyDefinitionsOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumPropertyDefinitionsOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// propertyDefinitions {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IFsrmCollection}(interface))
	{
		if o.PropertyDefinitions != nil {
			_ptr_propertyDefinitions := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.PropertyDefinitions != nil {
					if err := o.PropertyDefinitions.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&fsrm.Collection{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.PropertyDefinitions, _ptr_propertyDefinitions); err != nil {
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
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumPropertyDefinitionsOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// propertyDefinitions {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IFsrmCollection}(interface))
	{
		_ptr_propertyDefinitions := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.PropertyDefinitions == nil {
				o.PropertyDefinitions = &fsrm.Collection{}
			}
			if err := o.PropertyDefinitions.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_propertyDefinitions := func(ptr interface{}) { o.PropertyDefinitions = *ptr.(**fsrm.Collection) }
		if err := w.ReadPointer(&o.PropertyDefinitions, _s_propertyDefinitions, _ptr_propertyDefinitions); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// EnumPropertyDefinitionsRequest structure represents the EnumPropertyDefinitions operation request
type EnumPropertyDefinitionsRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// options: Contains the FsrmEnumOptions (section 2.2.1.2.5) to use when enumerating
	// the property definitions.
	Options fsrm.EnumOptions `idl:"name:options" json:"options"`
}

func (o *EnumPropertyDefinitionsRequest) xxx_ToOp(ctx context.Context) *xxx_EnumPropertyDefinitionsOperation {
	if o == nil {
		return &xxx_EnumPropertyDefinitionsOperation{}
	}
	return &xxx_EnumPropertyDefinitionsOperation{
		This:    o.This,
		Options: o.Options,
	}
}

func (o *EnumPropertyDefinitionsRequest) xxx_FromOp(ctx context.Context, op *xxx_EnumPropertyDefinitionsOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Options = op.Options
}
func (o *EnumPropertyDefinitionsRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *EnumPropertyDefinitionsRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EnumPropertyDefinitionsOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// EnumPropertyDefinitionsResponse structure represents the EnumPropertyDefinitions operation response
type EnumPropertyDefinitionsResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// propertyDefinitions: Pointer to an IFsrmCollection interface pointer (section 3.2.4.2.1)
	// that upon completion contains a pointer to every property definition on the server.
	// The caller MUST release the collection when the caller is done with it.
	PropertyDefinitions *fsrm.Collection `idl:"name:propertyDefinitions" json:"property_definitions"`
	// Return: The EnumPropertyDefinitions return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *EnumPropertyDefinitionsResponse) xxx_ToOp(ctx context.Context) *xxx_EnumPropertyDefinitionsOperation {
	if o == nil {
		return &xxx_EnumPropertyDefinitionsOperation{}
	}
	return &xxx_EnumPropertyDefinitionsOperation{
		That:                o.That,
		PropertyDefinitions: o.PropertyDefinitions,
		Return:              o.Return,
	}
}

func (o *EnumPropertyDefinitionsResponse) xxx_FromOp(ctx context.Context, op *xxx_EnumPropertyDefinitionsOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.PropertyDefinitions = op.PropertyDefinitions
	o.Return = op.Return
}
func (o *EnumPropertyDefinitionsResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *EnumPropertyDefinitionsResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EnumPropertyDefinitionsOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_CreatePropertyDefinitionOperation structure represents the CreatePropertyDefinition operation
type xxx_CreatePropertyDefinitionOperation struct {
	This               *dcom.ORPCThis           `idl:"name:This" json:"this"`
	That               *dcom.ORPCThat           `idl:"name:That" json:"that"`
	PropertyDefinition *fsrm.PropertyDefinition `idl:"name:propertyDefinition" json:"property_definition"`
	Return             int32                    `idl:"name:Return" json:"return"`
}

func (o *xxx_CreatePropertyDefinitionOperation) OpNum() int { return 19 }

func (o *xxx_CreatePropertyDefinitionOperation) OpName() string {
	return "/IFsrmClassificationManager/v0/CreatePropertyDefinition"
}

func (o *xxx_CreatePropertyDefinitionOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreatePropertyDefinitionOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreatePropertyDefinitionOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreatePropertyDefinitionOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreatePropertyDefinitionOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// propertyDefinition {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IFsrmPropertyDefinition}(interface))
	{
		if o.PropertyDefinition != nil {
			_ptr_propertyDefinition := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.PropertyDefinition != nil {
					if err := o.PropertyDefinition.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&fsrm.PropertyDefinition{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.PropertyDefinition, _ptr_propertyDefinition); err != nil {
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
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreatePropertyDefinitionOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// propertyDefinition {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IFsrmPropertyDefinition}(interface))
	{
		_ptr_propertyDefinition := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.PropertyDefinition == nil {
				o.PropertyDefinition = &fsrm.PropertyDefinition{}
			}
			if err := o.PropertyDefinition.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_propertyDefinition := func(ptr interface{}) { o.PropertyDefinition = *ptr.(**fsrm.PropertyDefinition) }
		if err := w.ReadPointer(&o.PropertyDefinition, _s_propertyDefinition, _ptr_propertyDefinition); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// CreatePropertyDefinitionRequest structure represents the CreatePropertyDefinition operation request
type CreatePropertyDefinitionRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *CreatePropertyDefinitionRequest) xxx_ToOp(ctx context.Context) *xxx_CreatePropertyDefinitionOperation {
	if o == nil {
		return &xxx_CreatePropertyDefinitionOperation{}
	}
	return &xxx_CreatePropertyDefinitionOperation{
		This: o.This,
	}
}

func (o *CreatePropertyDefinitionRequest) xxx_FromOp(ctx context.Context, op *xxx_CreatePropertyDefinitionOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *CreatePropertyDefinitionRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *CreatePropertyDefinitionRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreatePropertyDefinitionOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// CreatePropertyDefinitionResponse structure represents the CreatePropertyDefinition operation response
type CreatePropertyDefinitionResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// propertyDefinition: Pointer to an IFsrmPropertyDefinition interface pointer (section
	// 3.2.4.2.37) that upon completion points to a blank property definition. To have the
	// property definition added to the server's List of Persisted Property Definitions
	// (section 3.2.1.6), the client MUST call Commit (section 3.2.4.2.37.1). The caller
	// MUST release the property definition when the caller is done with it.
	PropertyDefinition *fsrm.PropertyDefinition `idl:"name:propertyDefinition" json:"property_definition"`
	// Return: The CreatePropertyDefinition return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *CreatePropertyDefinitionResponse) xxx_ToOp(ctx context.Context) *xxx_CreatePropertyDefinitionOperation {
	if o == nil {
		return &xxx_CreatePropertyDefinitionOperation{}
	}
	return &xxx_CreatePropertyDefinitionOperation{
		That:               o.That,
		PropertyDefinition: o.PropertyDefinition,
		Return:             o.Return,
	}
}

func (o *CreatePropertyDefinitionResponse) xxx_FromOp(ctx context.Context, op *xxx_CreatePropertyDefinitionOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.PropertyDefinition = op.PropertyDefinition
	o.Return = op.Return
}
func (o *CreatePropertyDefinitionResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *CreatePropertyDefinitionResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreatePropertyDefinitionOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetPropertyDefinitionOperation structure represents the GetPropertyDefinition operation
type xxx_GetPropertyDefinitionOperation struct {
	This               *dcom.ORPCThis           `idl:"name:This" json:"this"`
	That               *dcom.ORPCThat           `idl:"name:That" json:"that"`
	PropertyName       *oaut.String             `idl:"name:propertyName" json:"property_name"`
	PropertyDefinition *fsrm.PropertyDefinition `idl:"name:propertyDefinition" json:"property_definition"`
	Return             int32                    `idl:"name:Return" json:"return"`
}

func (o *xxx_GetPropertyDefinitionOperation) OpNum() int { return 20 }

func (o *xxx_GetPropertyDefinitionOperation) OpName() string {
	return "/IFsrmClassificationManager/v0/GetPropertyDefinition"
}

func (o *xxx_GetPropertyDefinitionOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetPropertyDefinitionOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// propertyName {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.PropertyName != nil {
			_ptr_propertyName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.PropertyName != nil {
					if err := o.PropertyName.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.PropertyName, _ptr_propertyName); err != nil {
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

func (o *xxx_GetPropertyDefinitionOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// propertyName {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_propertyName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.PropertyName == nil {
				o.PropertyName = &oaut.String{}
			}
			if err := o.PropertyName.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_propertyName := func(ptr interface{}) { o.PropertyName = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.PropertyName, _s_propertyName, _ptr_propertyName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetPropertyDefinitionOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetPropertyDefinitionOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// propertyDefinition {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IFsrmPropertyDefinition}(interface))
	{
		if o.PropertyDefinition != nil {
			_ptr_propertyDefinition := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.PropertyDefinition != nil {
					if err := o.PropertyDefinition.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&fsrm.PropertyDefinition{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.PropertyDefinition, _ptr_propertyDefinition); err != nil {
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
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetPropertyDefinitionOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// propertyDefinition {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IFsrmPropertyDefinition}(interface))
	{
		_ptr_propertyDefinition := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.PropertyDefinition == nil {
				o.PropertyDefinition = &fsrm.PropertyDefinition{}
			}
			if err := o.PropertyDefinition.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_propertyDefinition := func(ptr interface{}) { o.PropertyDefinition = *ptr.(**fsrm.PropertyDefinition) }
		if err := w.ReadPointer(&o.PropertyDefinition, _s_propertyDefinition, _ptr_propertyDefinition); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// GetPropertyDefinitionRequest structure represents the GetPropertyDefinition operation request
type GetPropertyDefinitionRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// propertyName: Contains the name of the property definition to return.
	PropertyName *oaut.String `idl:"name:propertyName" json:"property_name"`
}

func (o *GetPropertyDefinitionRequest) xxx_ToOp(ctx context.Context) *xxx_GetPropertyDefinitionOperation {
	if o == nil {
		return &xxx_GetPropertyDefinitionOperation{}
	}
	return &xxx_GetPropertyDefinitionOperation{
		This:         o.This,
		PropertyName: o.PropertyName,
	}
}

func (o *GetPropertyDefinitionRequest) xxx_FromOp(ctx context.Context, op *xxx_GetPropertyDefinitionOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.PropertyName = op.PropertyName
}
func (o *GetPropertyDefinitionRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetPropertyDefinitionRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetPropertyDefinitionOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetPropertyDefinitionResponse structure represents the GetPropertyDefinition operation response
type GetPropertyDefinitionResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// propertyDefinition: Pointer to an IFsrmPropertyDefinition interface pointer (section
	// 3.2.4.2.37) that upon completion points to the IFsrmPropertyDefinition with the specified
	// name. Additionally, the server SHOULD implement IFsrmPropertyDefinition2 for the
	// same interface pointer.<68> The caller MUST release the property definition when
	// the caller is done with it.
	PropertyDefinition *fsrm.PropertyDefinition `idl:"name:propertyDefinition" json:"property_definition"`
	// Return: The GetPropertyDefinition return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetPropertyDefinitionResponse) xxx_ToOp(ctx context.Context) *xxx_GetPropertyDefinitionOperation {
	if o == nil {
		return &xxx_GetPropertyDefinitionOperation{}
	}
	return &xxx_GetPropertyDefinitionOperation{
		That:               o.That,
		PropertyDefinition: o.PropertyDefinition,
		Return:             o.Return,
	}
}

func (o *GetPropertyDefinitionResponse) xxx_FromOp(ctx context.Context, op *xxx_GetPropertyDefinitionOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.PropertyDefinition = op.PropertyDefinition
	o.Return = op.Return
}
func (o *GetPropertyDefinitionResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetPropertyDefinitionResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetPropertyDefinitionOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_EnumRulesOperation structure represents the EnumRules operation
type xxx_EnumRulesOperation struct {
	This     *dcom.ORPCThis   `idl:"name:This" json:"this"`
	That     *dcom.ORPCThat   `idl:"name:That" json:"that"`
	RuleType fsrm.RuleType    `idl:"name:ruleType" json:"rule_type"`
	Options  fsrm.EnumOptions `idl:"name:options" json:"options"`
	Rules    *fsrm.Collection `idl:"name:Rules" json:"rules"`
	Return   int32            `idl:"name:Return" json:"return"`
}

func (o *xxx_EnumRulesOperation) OpNum() int { return 21 }

func (o *xxx_EnumRulesOperation) OpName() string { return "/IFsrmClassificationManager/v0/EnumRules" }

func (o *xxx_EnumRulesOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumRulesOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// ruleType {in} (1:{alias=FsrmRuleType}(enum))
	{
		if err := w.WriteData(uint16(o.RuleType)); err != nil {
			return err
		}
	}
	// options {in, default_value={0}} (1:{alias=FsrmEnumOptions}(enum))
	{
		if err := w.WriteData(uint16(o.Options)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumRulesOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// ruleType {in} (1:{alias=FsrmRuleType}(enum))
	{
		if err := w.ReadData((*uint16)(&o.RuleType)); err != nil {
			return err
		}
	}
	// options {in, default_value={0}} (1:{alias=FsrmEnumOptions}(enum))
	{
		if err := w.ReadData((*uint16)(&o.Options)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumRulesOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumRulesOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Rules {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IFsrmCollection}(interface))
	{
		if o.Rules != nil {
			_ptr_Rules := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Rules != nil {
					if err := o.Rules.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&fsrm.Collection{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Rules, _ptr_Rules); err != nil {
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
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumRulesOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Rules {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IFsrmCollection}(interface))
	{
		_ptr_Rules := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Rules == nil {
				o.Rules = &fsrm.Collection{}
			}
			if err := o.Rules.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_Rules := func(ptr interface{}) { o.Rules = *ptr.(**fsrm.Collection) }
		if err := w.ReadPointer(&o.Rules, _s_Rules, _ptr_Rules); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// EnumRulesRequest structure represents the EnumRules operation request
type EnumRulesRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// ruleType: Contains the FsrmRuleType (section 2.2.1.2.11) to which to limit the returned
	// collection of Rules.
	RuleType fsrm.RuleType `idl:"name:ruleType" json:"rule_type"`
	// options: Contains the FsrmEnumOptions (section 2.2.1.2.5) to use when enumerating
	// the Rules.
	Options fsrm.EnumOptions `idl:"name:options" json:"options"`
}

func (o *EnumRulesRequest) xxx_ToOp(ctx context.Context) *xxx_EnumRulesOperation {
	if o == nil {
		return &xxx_EnumRulesOperation{}
	}
	return &xxx_EnumRulesOperation{
		This:     o.This,
		RuleType: o.RuleType,
		Options:  o.Options,
	}
}

func (o *EnumRulesRequest) xxx_FromOp(ctx context.Context, op *xxx_EnumRulesOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.RuleType = op.RuleType
	o.Options = op.Options
}
func (o *EnumRulesRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *EnumRulesRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EnumRulesOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// EnumRulesResponse structure represents the EnumRules operation response
type EnumRulesResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Rules: Pointer to an IFsrmCollection interface pointer (section 3.2.4.2.1) that upon
	// completion contains pointers to every rule on the server that has the rule type specified
	// by ruleType. A caller MUST release the collection received when the caller is done
	// with it.
	Rules *fsrm.Collection `idl:"name:Rules" json:"rules"`
	// Return: The EnumRules return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *EnumRulesResponse) xxx_ToOp(ctx context.Context) *xxx_EnumRulesOperation {
	if o == nil {
		return &xxx_EnumRulesOperation{}
	}
	return &xxx_EnumRulesOperation{
		That:   o.That,
		Rules:  o.Rules,
		Return: o.Return,
	}
}

func (o *EnumRulesResponse) xxx_FromOp(ctx context.Context, op *xxx_EnumRulesOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Rules = op.Rules
	o.Return = op.Return
}
func (o *EnumRulesResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *EnumRulesResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EnumRulesOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_CreateRuleOperation structure represents the CreateRule operation
type xxx_CreateRuleOperation struct {
	This     *dcom.ORPCThis `idl:"name:This" json:"this"`
	That     *dcom.ORPCThat `idl:"name:That" json:"that"`
	RuleType fsrm.RuleType  `idl:"name:ruleType" json:"rule_type"`
	Rule     *fsrm.Rule     `idl:"name:Rule" json:"rule"`
	Return   int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_CreateRuleOperation) OpNum() int { return 22 }

func (o *xxx_CreateRuleOperation) OpName() string { return "/IFsrmClassificationManager/v0/CreateRule" }

func (o *xxx_CreateRuleOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateRuleOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// ruleType {in} (1:{alias=FsrmRuleType}(enum))
	{
		if err := w.WriteData(uint16(o.RuleType)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateRuleOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// ruleType {in} (1:{alias=FsrmRuleType}(enum))
	{
		if err := w.ReadData((*uint16)(&o.RuleType)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateRuleOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateRuleOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Rule {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IFsrmRule}(interface))
	{
		if o.Rule != nil {
			_ptr_Rule := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Rule != nil {
					if err := o.Rule.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&fsrm.Rule{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Rule, _ptr_Rule); err != nil {
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
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateRuleOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Rule {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IFsrmRule}(interface))
	{
		_ptr_Rule := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Rule == nil {
				o.Rule = &fsrm.Rule{}
			}
			if err := o.Rule.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_Rule := func(ptr interface{}) { o.Rule = *ptr.(**fsrm.Rule) }
		if err := w.ReadPointer(&o.Rule, _s_Rule, _ptr_Rule); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// CreateRuleRequest structure represents the CreateRule operation request
type CreateRuleRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// ruleType: Contains the FsrmRuleType (section 2.2.1.2.11) of the classification rule
	// to create.
	RuleType fsrm.RuleType `idl:"name:ruleType" json:"rule_type"`
}

func (o *CreateRuleRequest) xxx_ToOp(ctx context.Context) *xxx_CreateRuleOperation {
	if o == nil {
		return &xxx_CreateRuleOperation{}
	}
	return &xxx_CreateRuleOperation{
		This:     o.This,
		RuleType: o.RuleType,
	}
}

func (o *CreateRuleRequest) xxx_FromOp(ctx context.Context, op *xxx_CreateRuleOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.RuleType = op.RuleType
}
func (o *CreateRuleRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *CreateRuleRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreateRuleOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// CreateRuleResponse structure represents the CreateRule operation response
type CreateRuleResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Rule: Pointer to an IFsrmRule interface pointer (section 3.2.4.2.41) that upon completion
	// points to a blank classification rule. To have the rule added to the server's List
	// of Persisted Rules (section 3.2.1.6) the client MUST call Commit (section 3.2.4.2.10.5)
	// on the rule. The caller MUST release the rule when the caller is done with it.
	Rule *fsrm.Rule `idl:"name:Rule" json:"rule"`
	// Return: The CreateRule return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *CreateRuleResponse) xxx_ToOp(ctx context.Context) *xxx_CreateRuleOperation {
	if o == nil {
		return &xxx_CreateRuleOperation{}
	}
	return &xxx_CreateRuleOperation{
		That:   o.That,
		Rule:   o.Rule,
		Return: o.Return,
	}
}

func (o *CreateRuleResponse) xxx_FromOp(ctx context.Context, op *xxx_CreateRuleOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Rule = op.Rule
	o.Return = op.Return
}
func (o *CreateRuleResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *CreateRuleResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreateRuleOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetRuleOperation structure represents the GetRule operation
type xxx_GetRuleOperation struct {
	This     *dcom.ORPCThis `idl:"name:This" json:"this"`
	That     *dcom.ORPCThat `idl:"name:That" json:"that"`
	RuleName *oaut.String   `idl:"name:ruleName" json:"rule_name"`
	RuleType fsrm.RuleType  `idl:"name:ruleType" json:"rule_type"`
	Rule     *fsrm.Rule     `idl:"name:Rule" json:"rule"`
	Return   int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetRuleOperation) OpNum() int { return 23 }

func (o *xxx_GetRuleOperation) OpName() string { return "/IFsrmClassificationManager/v0/GetRule" }

func (o *xxx_GetRuleOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetRuleOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// ruleName {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.RuleName != nil {
			_ptr_ruleName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.RuleName != nil {
					if err := o.RuleName.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.RuleName, _ptr_ruleName); err != nil {
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
	// ruleType {in} (1:{alias=FsrmRuleType}(enum))
	{
		if err := w.WriteData(uint16(o.RuleType)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetRuleOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// ruleName {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_ruleName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.RuleName == nil {
				o.RuleName = &oaut.String{}
			}
			if err := o.RuleName.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ruleName := func(ptr interface{}) { o.RuleName = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.RuleName, _s_ruleName, _ptr_ruleName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// ruleType {in} (1:{alias=FsrmRuleType}(enum))
	{
		if err := w.ReadData((*uint16)(&o.RuleType)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetRuleOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetRuleOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Rule {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IFsrmRule}(interface))
	{
		if o.Rule != nil {
			_ptr_Rule := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Rule != nil {
					if err := o.Rule.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&fsrm.Rule{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Rule, _ptr_Rule); err != nil {
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
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetRuleOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Rule {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IFsrmRule}(interface))
	{
		_ptr_Rule := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Rule == nil {
				o.Rule = &fsrm.Rule{}
			}
			if err := o.Rule.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_Rule := func(ptr interface{}) { o.Rule = *ptr.(**fsrm.Rule) }
		if err := w.ReadPointer(&o.Rule, _s_Rule, _ptr_Rule); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// GetRuleRequest structure represents the GetRule operation request
type GetRuleRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// ruleName: Contains the Name of the classification rule to return.
	RuleName *oaut.String `idl:"name:ruleName" json:"rule_name"`
	// ruleType: Contains the FsrmRuleType (section 2.2.1.2.11) of the classification rule
	// to return.
	RuleType fsrm.RuleType `idl:"name:ruleType" json:"rule_type"`
}

func (o *GetRuleRequest) xxx_ToOp(ctx context.Context) *xxx_GetRuleOperation {
	if o == nil {
		return &xxx_GetRuleOperation{}
	}
	return &xxx_GetRuleOperation{
		This:     o.This,
		RuleName: o.RuleName,
		RuleType: o.RuleType,
	}
}

func (o *GetRuleRequest) xxx_FromOp(ctx context.Context, op *xxx_GetRuleOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.RuleName = op.RuleName
	o.RuleType = op.RuleType
}
func (o *GetRuleRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetRuleRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetRuleOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetRuleResponse structure represents the GetRule operation response
type GetRuleResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Rule: Pointer to an IFsrmRule interface pointer (section 3.2.4.2.41) that upon completion
	// points to the classification rule with the specified Name and Rule type. The caller
	// MUST release the rule when the caller is done with it.
	Rule *fsrm.Rule `idl:"name:Rule" json:"rule"`
	// Return: The GetRule return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetRuleResponse) xxx_ToOp(ctx context.Context) *xxx_GetRuleOperation {
	if o == nil {
		return &xxx_GetRuleOperation{}
	}
	return &xxx_GetRuleOperation{
		That:   o.That,
		Rule:   o.Rule,
		Return: o.Return,
	}
}

func (o *GetRuleResponse) xxx_FromOp(ctx context.Context, op *xxx_GetRuleOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Rule = op.Rule
	o.Return = op.Return
}
func (o *GetRuleResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetRuleResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetRuleOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_EnumModuleDefinitionsOperation structure represents the EnumModuleDefinitions operation
type xxx_EnumModuleDefinitionsOperation struct {
	This              *dcom.ORPCThis          `idl:"name:This" json:"this"`
	That              *dcom.ORPCThat          `idl:"name:That" json:"that"`
	ModuleType        fsrm.PipelineModuleType `idl:"name:moduleType" json:"module_type"`
	Options           fsrm.EnumOptions        `idl:"name:options" json:"options"`
	ModuleDefinitions *fsrm.Collection        `idl:"name:moduleDefinitions" json:"module_definitions"`
	Return            int32                   `idl:"name:Return" json:"return"`
}

func (o *xxx_EnumModuleDefinitionsOperation) OpNum() int { return 24 }

func (o *xxx_EnumModuleDefinitionsOperation) OpName() string {
	return "/IFsrmClassificationManager/v0/EnumModuleDefinitions"
}

func (o *xxx_EnumModuleDefinitionsOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumModuleDefinitionsOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// moduleType {in} (1:{alias=FsrmPipelineModuleType}(enum))
	{
		if err := w.WriteData(uint16(o.ModuleType)); err != nil {
			return err
		}
	}
	// options {in, default_value={0}} (1:{alias=FsrmEnumOptions}(enum))
	{
		if err := w.WriteData(uint16(o.Options)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumModuleDefinitionsOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// moduleType {in} (1:{alias=FsrmPipelineModuleType}(enum))
	{
		if err := w.ReadData((*uint16)(&o.ModuleType)); err != nil {
			return err
		}
	}
	// options {in, default_value={0}} (1:{alias=FsrmEnumOptions}(enum))
	{
		if err := w.ReadData((*uint16)(&o.Options)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumModuleDefinitionsOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumModuleDefinitionsOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// moduleDefinitions {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IFsrmCollection}(interface))
	{
		if o.ModuleDefinitions != nil {
			_ptr_moduleDefinitions := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ModuleDefinitions != nil {
					if err := o.ModuleDefinitions.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&fsrm.Collection{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ModuleDefinitions, _ptr_moduleDefinitions); err != nil {
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
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumModuleDefinitionsOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// moduleDefinitions {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IFsrmCollection}(interface))
	{
		_ptr_moduleDefinitions := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ModuleDefinitions == nil {
				o.ModuleDefinitions = &fsrm.Collection{}
			}
			if err := o.ModuleDefinitions.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_moduleDefinitions := func(ptr interface{}) { o.ModuleDefinitions = *ptr.(**fsrm.Collection) }
		if err := w.ReadPointer(&o.ModuleDefinitions, _s_moduleDefinitions, _ptr_moduleDefinitions); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// EnumModuleDefinitionsRequest structure represents the EnumModuleDefinitions operation request
type EnumModuleDefinitionsRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// moduleType: Contains the FsrmPipelineModuleType (section 2.2.1.2.12) of the modules
	// to get.
	ModuleType fsrm.PipelineModuleType `idl:"name:moduleType" json:"module_type"`
	// options: Contains the FsrmEnumOptions (section 2.2.1.2.5) to use when enumerating
	// the moduleDefinitions.
	Options fsrm.EnumOptions `idl:"name:options" json:"options"`
}

func (o *EnumModuleDefinitionsRequest) xxx_ToOp(ctx context.Context) *xxx_EnumModuleDefinitionsOperation {
	if o == nil {
		return &xxx_EnumModuleDefinitionsOperation{}
	}
	return &xxx_EnumModuleDefinitionsOperation{
		This:       o.This,
		ModuleType: o.ModuleType,
		Options:    o.Options,
	}
}

func (o *EnumModuleDefinitionsRequest) xxx_FromOp(ctx context.Context, op *xxx_EnumModuleDefinitionsOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.ModuleType = op.ModuleType
	o.Options = op.Options
}
func (o *EnumModuleDefinitionsRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *EnumModuleDefinitionsRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EnumModuleDefinitionsOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// EnumModuleDefinitionsResponse structure represents the EnumModuleDefinitions operation response
type EnumModuleDefinitionsResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// moduleDefinitions: Pointer to an IFsrmCollection interface pointer (section 3.2.4.2.1)
	// that upon completion contains pointers to every module definition on the server that
	// has the specified moduleType. A caller MUST release the collection received when
	// the caller is done with it.
	ModuleDefinitions *fsrm.Collection `idl:"name:moduleDefinitions" json:"module_definitions"`
	// Return: The EnumModuleDefinitions return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *EnumModuleDefinitionsResponse) xxx_ToOp(ctx context.Context) *xxx_EnumModuleDefinitionsOperation {
	if o == nil {
		return &xxx_EnumModuleDefinitionsOperation{}
	}
	return &xxx_EnumModuleDefinitionsOperation{
		That:              o.That,
		ModuleDefinitions: o.ModuleDefinitions,
		Return:            o.Return,
	}
}

func (o *EnumModuleDefinitionsResponse) xxx_FromOp(ctx context.Context, op *xxx_EnumModuleDefinitionsOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.ModuleDefinitions = op.ModuleDefinitions
	o.Return = op.Return
}
func (o *EnumModuleDefinitionsResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *EnumModuleDefinitionsResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EnumModuleDefinitionsOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_CreateModuleDefinitionOperation structure represents the CreateModuleDefinition operation
type xxx_CreateModuleDefinitionOperation struct {
	This             *dcom.ORPCThis                 `idl:"name:This" json:"this"`
	That             *dcom.ORPCThat                 `idl:"name:That" json:"that"`
	ModuleType       fsrm.PipelineModuleType        `idl:"name:moduleType" json:"module_type"`
	ModuleDefinition *fsrm.PipelineModuleDefinition `idl:"name:moduleDefinition" json:"module_definition"`
	Return           int32                          `idl:"name:Return" json:"return"`
}

func (o *xxx_CreateModuleDefinitionOperation) OpNum() int { return 25 }

func (o *xxx_CreateModuleDefinitionOperation) OpName() string {
	return "/IFsrmClassificationManager/v0/CreateModuleDefinition"
}

func (o *xxx_CreateModuleDefinitionOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateModuleDefinitionOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// moduleType {in} (1:{alias=FsrmPipelineModuleType}(enum))
	{
		if err := w.WriteData(uint16(o.ModuleType)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateModuleDefinitionOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// moduleType {in} (1:{alias=FsrmPipelineModuleType}(enum))
	{
		if err := w.ReadData((*uint16)(&o.ModuleType)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateModuleDefinitionOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateModuleDefinitionOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// moduleDefinition {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IFsrmPipelineModuleDefinition}(interface))
	{
		if o.ModuleDefinition != nil {
			_ptr_moduleDefinition := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ModuleDefinition != nil {
					if err := o.ModuleDefinition.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&fsrm.PipelineModuleDefinition{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ModuleDefinition, _ptr_moduleDefinition); err != nil {
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
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateModuleDefinitionOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// moduleDefinition {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IFsrmPipelineModuleDefinition}(interface))
	{
		_ptr_moduleDefinition := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ModuleDefinition == nil {
				o.ModuleDefinition = &fsrm.PipelineModuleDefinition{}
			}
			if err := o.ModuleDefinition.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_moduleDefinition := func(ptr interface{}) { o.ModuleDefinition = *ptr.(**fsrm.PipelineModuleDefinition) }
		if err := w.ReadPointer(&o.ModuleDefinition, _s_moduleDefinition, _ptr_moduleDefinition); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// CreateModuleDefinitionRequest structure represents the CreateModuleDefinition operation request
type CreateModuleDefinitionRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// moduleType: Contains the type of module to create (for example, a classifier or storage
	// module). For possible types, see the FsrmPipelineModuleType (section 2.2.1.2.12)
	// enumeration.
	ModuleType fsrm.PipelineModuleType `idl:"name:moduleType" json:"module_type"`
}

func (o *CreateModuleDefinitionRequest) xxx_ToOp(ctx context.Context) *xxx_CreateModuleDefinitionOperation {
	if o == nil {
		return &xxx_CreateModuleDefinitionOperation{}
	}
	return &xxx_CreateModuleDefinitionOperation{
		This:       o.This,
		ModuleType: o.ModuleType,
	}
}

func (o *CreateModuleDefinitionRequest) xxx_FromOp(ctx context.Context, op *xxx_CreateModuleDefinitionOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.ModuleType = op.ModuleType
}
func (o *CreateModuleDefinitionRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *CreateModuleDefinitionRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreateModuleDefinitionOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// CreateModuleDefinitionResponse structure represents the CreateModuleDefinition operation response
type CreateModuleDefinitionResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// moduleDefinition: An IFsrmPipelineModuleDefinition interface pointer (section 3.2.4.2.43)
	// to the new module definition. Query the IFsrmPipelineModuleDefinition interface to
	// get the interface for the specified module. To add the module definition to the server's
	// List of Persisted Module Definitions (section 3.2.1.6), the client MUST call Commit
	// (section 3.2.4.2.10.5).
	ModuleDefinition *fsrm.PipelineModuleDefinition `idl:"name:moduleDefinition" json:"module_definition"`
	// Return: The CreateModuleDefinition return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *CreateModuleDefinitionResponse) xxx_ToOp(ctx context.Context) *xxx_CreateModuleDefinitionOperation {
	if o == nil {
		return &xxx_CreateModuleDefinitionOperation{}
	}
	return &xxx_CreateModuleDefinitionOperation{
		That:             o.That,
		ModuleDefinition: o.ModuleDefinition,
		Return:           o.Return,
	}
}

func (o *CreateModuleDefinitionResponse) xxx_FromOp(ctx context.Context, op *xxx_CreateModuleDefinitionOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.ModuleDefinition = op.ModuleDefinition
	o.Return = op.Return
}
func (o *CreateModuleDefinitionResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *CreateModuleDefinitionResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreateModuleDefinitionOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetModuleDefinitionOperation structure represents the GetModuleDefinition operation
type xxx_GetModuleDefinitionOperation struct {
	This             *dcom.ORPCThis                 `idl:"name:This" json:"this"`
	That             *dcom.ORPCThat                 `idl:"name:That" json:"that"`
	ModuleName       *oaut.String                   `idl:"name:moduleName" json:"module_name"`
	ModuleType       fsrm.PipelineModuleType        `idl:"name:moduleType" json:"module_type"`
	ModuleDefinition *fsrm.PipelineModuleDefinition `idl:"name:moduleDefinition" json:"module_definition"`
	Return           int32                          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetModuleDefinitionOperation) OpNum() int { return 26 }

func (o *xxx_GetModuleDefinitionOperation) OpName() string {
	return "/IFsrmClassificationManager/v0/GetModuleDefinition"
}

func (o *xxx_GetModuleDefinitionOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetModuleDefinitionOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// moduleName {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.ModuleName != nil {
			_ptr_moduleName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ModuleName != nil {
					if err := o.ModuleName.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ModuleName, _ptr_moduleName); err != nil {
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
	// moduleType {in} (1:{alias=FsrmPipelineModuleType}(enum))
	{
		if err := w.WriteData(uint16(o.ModuleType)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetModuleDefinitionOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// moduleName {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_moduleName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ModuleName == nil {
				o.ModuleName = &oaut.String{}
			}
			if err := o.ModuleName.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_moduleName := func(ptr interface{}) { o.ModuleName = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.ModuleName, _s_moduleName, _ptr_moduleName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// moduleType {in} (1:{alias=FsrmPipelineModuleType}(enum))
	{
		if err := w.ReadData((*uint16)(&o.ModuleType)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetModuleDefinitionOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetModuleDefinitionOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// moduleDefinition {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IFsrmPipelineModuleDefinition}(interface))
	{
		if o.ModuleDefinition != nil {
			_ptr_moduleDefinition := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ModuleDefinition != nil {
					if err := o.ModuleDefinition.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&fsrm.PipelineModuleDefinition{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ModuleDefinition, _ptr_moduleDefinition); err != nil {
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
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetModuleDefinitionOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// moduleDefinition {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IFsrmPipelineModuleDefinition}(interface))
	{
		_ptr_moduleDefinition := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ModuleDefinition == nil {
				o.ModuleDefinition = &fsrm.PipelineModuleDefinition{}
			}
			if err := o.ModuleDefinition.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_moduleDefinition := func(ptr interface{}) { o.ModuleDefinition = *ptr.(**fsrm.PipelineModuleDefinition) }
		if err := w.ReadPointer(&o.ModuleDefinition, _s_moduleDefinition, _ptr_moduleDefinition); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// GetModuleDefinitionRequest structure represents the GetModuleDefinition operation request
type GetModuleDefinitionRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// moduleName: Contains the name of the module definition to retrieve.
	ModuleName *oaut.String `idl:"name:moduleName" json:"module_name"`
	// moduleType: Contains the type of the module definition to retrieve. For possible
	// types, see the FsrmPipelineModuleType (section 2.2.1.2.12) enumeration.
	ModuleType fsrm.PipelineModuleType `idl:"name:moduleType" json:"module_type"`
}

func (o *GetModuleDefinitionRequest) xxx_ToOp(ctx context.Context) *xxx_GetModuleDefinitionOperation {
	if o == nil {
		return &xxx_GetModuleDefinitionOperation{}
	}
	return &xxx_GetModuleDefinitionOperation{
		This:       o.This,
		ModuleName: o.ModuleName,
		ModuleType: o.ModuleType,
	}
}

func (o *GetModuleDefinitionRequest) xxx_FromOp(ctx context.Context, op *xxx_GetModuleDefinitionOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.ModuleName = op.ModuleName
	o.ModuleType = op.ModuleType
}
func (o *GetModuleDefinitionRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetModuleDefinitionRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetModuleDefinitionOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetModuleDefinitionResponse structure represents the GetModuleDefinition operation response
type GetModuleDefinitionResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// moduleDefinition: An IFsrmPipelineModuleDefinition interface pointer (section 3.2.4.2.43)
	// to the module definition. Query the IFsrmPipelineModuleDefinition interface to get
	// the interface pointer for the specified module.
	ModuleDefinition *fsrm.PipelineModuleDefinition `idl:"name:moduleDefinition" json:"module_definition"`
	// Return: The GetModuleDefinition return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetModuleDefinitionResponse) xxx_ToOp(ctx context.Context) *xxx_GetModuleDefinitionOperation {
	if o == nil {
		return &xxx_GetModuleDefinitionOperation{}
	}
	return &xxx_GetModuleDefinitionOperation{
		That:             o.That,
		ModuleDefinition: o.ModuleDefinition,
		Return:           o.Return,
	}
}

func (o *GetModuleDefinitionResponse) xxx_FromOp(ctx context.Context, op *xxx_GetModuleDefinitionOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.ModuleDefinition = op.ModuleDefinition
	o.Return = op.Return
}
func (o *GetModuleDefinitionResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetModuleDefinitionResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetModuleDefinitionOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_RunClassificationOperation structure represents the RunClassification operation
type xxx_RunClassificationOperation struct {
	This    *dcom.ORPCThis               `idl:"name:This" json:"this"`
	That    *dcom.ORPCThat               `idl:"name:That" json:"that"`
	Context fsrm.ReportGenerationContext `idl:"name:context" json:"context"`
	_       *oaut.String                 `idl:"name:reserved"`
	Return  int32                        `idl:"name:Return" json:"return"`
}

func (o *xxx_RunClassificationOperation) OpNum() int { return 27 }

func (o *xxx_RunClassificationOperation) OpName() string {
	return "/IFsrmClassificationManager/v0/RunClassification"
}

func (o *xxx_RunClassificationOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RunClassificationOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// context {in} (1:{alias=FsrmReportGenerationContext}(enum))
	{
		if err := w.WriteData(uint16(o.Context)); err != nil {
			return err
		}
	}
	// reserved {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		// reserved reserved
		if err := w.WritePointer(nil); err != nil {
			return err
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RunClassificationOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// context {in} (1:{alias=FsrmReportGenerationContext}(enum))
	{
		if err := w.ReadData((*uint16)(&o.Context)); err != nil {
			return err
		}
	}
	// reserved {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		// reserved reserved
		var _reserved *oaut.String
		_ptr_reserved := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if _reserved == nil {
				_reserved = &oaut.String{}
			}
			if err := _reserved.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_reserved := func(ptr interface{}) { _reserved = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&_reserved, _s_reserved, _ptr_reserved); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RunClassificationOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RunClassificationOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RunClassificationOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// RunClassificationRequest structure represents the RunClassification operation request
type RunClassificationRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// context: Contains the value of the report generation context in which the running
	// classification task will run.
	Context fsrm.ReportGenerationContext `idl:"name:context" json:"context"`
}

func (o *RunClassificationRequest) xxx_ToOp(ctx context.Context) *xxx_RunClassificationOperation {
	if o == nil {
		return &xxx_RunClassificationOperation{}
	}
	return &xxx_RunClassificationOperation{
		This:    o.This,
		Context: o.Context,
	}
}

func (o *RunClassificationRequest) xxx_FromOp(ctx context.Context, op *xxx_RunClassificationOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Context = op.Context
}
func (o *RunClassificationRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *RunClassificationRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RunClassificationOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RunClassificationResponse structure represents the RunClassification operation response
type RunClassificationResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The RunClassification return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *RunClassificationResponse) xxx_ToOp(ctx context.Context) *xxx_RunClassificationOperation {
	if o == nil {
		return &xxx_RunClassificationOperation{}
	}
	return &xxx_RunClassificationOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *RunClassificationResponse) xxx_FromOp(ctx context.Context, op *xxx_RunClassificationOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *RunClassificationResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *RunClassificationResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RunClassificationOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_WaitForClassificationCompletionOperation structure represents the WaitForClassificationCompletion operation
type xxx_WaitForClassificationCompletionOperation struct {
	This        *dcom.ORPCThis `idl:"name:This" json:"this"`
	That        *dcom.ORPCThat `idl:"name:That" json:"that"`
	WaitSeconds int32          `idl:"name:waitSeconds" json:"wait_seconds"`
	Completed   int16          `idl:"name:completed" json:"completed"`
	Return      int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_WaitForClassificationCompletionOperation) OpNum() int { return 28 }

func (o *xxx_WaitForClassificationCompletionOperation) OpName() string {
	return "/IFsrmClassificationManager/v0/WaitForClassificationCompletion"
}

func (o *xxx_WaitForClassificationCompletionOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WaitForClassificationCompletionOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// waitSeconds {in} (1:(int32))
	{
		if err := w.WriteData(o.WaitSeconds); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WaitForClassificationCompletionOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// waitSeconds {in} (1:(int32))
	{
		if err := w.ReadData(&o.WaitSeconds); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WaitForClassificationCompletionOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WaitForClassificationCompletionOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// completed {out, retval} (1:{pointer=ref}*(1))(2:{alias=VARIANT_BOOL}(int16))
	{
		if err := w.WriteData(o.Completed); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WaitForClassificationCompletionOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// completed {out, retval} (1:{pointer=ref}*(1))(2:{alias=VARIANT_BOOL}(int16))
	{
		if err := w.ReadData(&o.Completed); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// WaitForClassificationCompletionRequest structure represents the WaitForClassificationCompletion operation request
type WaitForClassificationCompletionRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// waitSeconds: Contains the maximum number of seconds the call will block before returning.
	WaitSeconds int32 `idl:"name:waitSeconds" json:"wait_seconds"`
}

func (o *WaitForClassificationCompletionRequest) xxx_ToOp(ctx context.Context) *xxx_WaitForClassificationCompletionOperation {
	if o == nil {
		return &xxx_WaitForClassificationCompletionOperation{}
	}
	return &xxx_WaitForClassificationCompletionOperation{
		This:        o.This,
		WaitSeconds: o.WaitSeconds,
	}
}

func (o *WaitForClassificationCompletionRequest) xxx_FromOp(ctx context.Context, op *xxx_WaitForClassificationCompletionOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.WaitSeconds = op.WaitSeconds
}
func (o *WaitForClassificationCompletionRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *WaitForClassificationCompletionRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_WaitForClassificationCompletionOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// WaitForClassificationCompletionResponse structure represents the WaitForClassificationCompletion operation response
type WaitForClassificationCompletionResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// completed: Pointer to a VARIANT_BOOL variable that upon completion contains an indication
	// of whether the running classification task has completed.
	Completed int16 `idl:"name:completed" json:"completed"`
	// Return: The WaitForClassificationCompletion return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *WaitForClassificationCompletionResponse) xxx_ToOp(ctx context.Context) *xxx_WaitForClassificationCompletionOperation {
	if o == nil {
		return &xxx_WaitForClassificationCompletionOperation{}
	}
	return &xxx_WaitForClassificationCompletionOperation{
		That:      o.That,
		Completed: o.Completed,
		Return:    o.Return,
	}
}

func (o *WaitForClassificationCompletionResponse) xxx_FromOp(ctx context.Context, op *xxx_WaitForClassificationCompletionOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Completed = op.Completed
	o.Return = op.Return
}
func (o *WaitForClassificationCompletionResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *WaitForClassificationCompletionResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_WaitForClassificationCompletionOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_CancelClassificationOperation structure represents the CancelClassification operation
type xxx_CancelClassificationOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_CancelClassificationOperation) OpNum() int { return 29 }

func (o *xxx_CancelClassificationOperation) OpName() string {
	return "/IFsrmClassificationManager/v0/CancelClassification"
}

func (o *xxx_CancelClassificationOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CancelClassificationOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CancelClassificationOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CancelClassificationOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CancelClassificationOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CancelClassificationOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// CancelClassificationRequest structure represents the CancelClassification operation request
type CancelClassificationRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *CancelClassificationRequest) xxx_ToOp(ctx context.Context) *xxx_CancelClassificationOperation {
	if o == nil {
		return &xxx_CancelClassificationOperation{}
	}
	return &xxx_CancelClassificationOperation{
		This: o.This,
	}
}

func (o *CancelClassificationRequest) xxx_FromOp(ctx context.Context, op *xxx_CancelClassificationOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *CancelClassificationRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *CancelClassificationRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CancelClassificationOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// CancelClassificationResponse structure represents the CancelClassification operation response
type CancelClassificationResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The CancelClassification return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *CancelClassificationResponse) xxx_ToOp(ctx context.Context) *xxx_CancelClassificationOperation {
	if o == nil {
		return &xxx_CancelClassificationOperation{}
	}
	return &xxx_CancelClassificationOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *CancelClassificationResponse) xxx_FromOp(ctx context.Context, op *xxx_CancelClassificationOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *CancelClassificationResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *CancelClassificationResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CancelClassificationOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_EnumFilePropertiesOperation structure represents the EnumFileProperties operation
type xxx_EnumFilePropertiesOperation struct {
	This           *dcom.ORPCThis              `idl:"name:This" json:"this"`
	That           *dcom.ORPCThat              `idl:"name:That" json:"that"`
	FilePath       *oaut.String                `idl:"name:filePath" json:"file_path"`
	Options        fsrm.GetFilePropertyOptions `idl:"name:options" json:"options"`
	FileProperties *fsrm.Collection            `idl:"name:fileProperties" json:"file_properties"`
	Return         int32                       `idl:"name:Return" json:"return"`
}

func (o *xxx_EnumFilePropertiesOperation) OpNum() int { return 30 }

func (o *xxx_EnumFilePropertiesOperation) OpName() string {
	return "/IFsrmClassificationManager/v0/EnumFileProperties"
}

func (o *xxx_EnumFilePropertiesOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumFilePropertiesOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// filePath {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.FilePath != nil {
			_ptr_filePath := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.FilePath != nil {
					if err := o.FilePath.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.FilePath, _ptr_filePath); err != nil {
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
	// options {in, default_value={0}} (1:{alias=FsrmGetFilePropertyOptions}(enum))
	{
		if err := w.WriteData(uint16(o.Options)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumFilePropertiesOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// filePath {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_filePath := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.FilePath == nil {
				o.FilePath = &oaut.String{}
			}
			if err := o.FilePath.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_filePath := func(ptr interface{}) { o.FilePath = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.FilePath, _s_filePath, _ptr_filePath); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// options {in, default_value={0}} (1:{alias=FsrmGetFilePropertyOptions}(enum))
	{
		if err := w.ReadData((*uint16)(&o.Options)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumFilePropertiesOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumFilePropertiesOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// fileProperties {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IFsrmCollection}(interface))
	{
		if o.FileProperties != nil {
			_ptr_fileProperties := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.FileProperties != nil {
					if err := o.FileProperties.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&fsrm.Collection{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.FileProperties, _ptr_fileProperties); err != nil {
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
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumFilePropertiesOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// fileProperties {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IFsrmCollection}(interface))
	{
		_ptr_fileProperties := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.FileProperties == nil {
				o.FileProperties = &fsrm.Collection{}
			}
			if err := o.FileProperties.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_fileProperties := func(ptr interface{}) { o.FileProperties = *ptr.(**fsrm.Collection) }
		if err := w.ReadPointer(&o.FileProperties, _s_fileProperties, _ptr_fileProperties); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// EnumFilePropertiesRequest structure represents the EnumFileProperties operation request
type EnumFilePropertiesRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// filePath: The file or folder that contains the Property Definition Instances that
	// you want to enumerate. You can specify an absolute or relative path to the file or
	// folder. You cannot specify a file share.
	FilePath *oaut.String `idl:"name:filePath" json:"file_path"`
	// options: Contains the options to use for enumerating the file's Property Definition
	// Instances. For possible values, see the FsrmGetFilePropertyOptions (section 2.2.2.5.1.2)
	// enumeration.
	Options fsrm.GetFilePropertyOptions `idl:"name:options" json:"options"`
}

func (o *EnumFilePropertiesRequest) xxx_ToOp(ctx context.Context) *xxx_EnumFilePropertiesOperation {
	if o == nil {
		return &xxx_EnumFilePropertiesOperation{}
	}
	return &xxx_EnumFilePropertiesOperation{
		This:     o.This,
		FilePath: o.FilePath,
		Options:  o.Options,
	}
}

func (o *EnumFilePropertiesRequest) xxx_FromOp(ctx context.Context, op *xxx_EnumFilePropertiesOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.FilePath = op.FilePath
	o.Options = op.Options
}
func (o *EnumFilePropertiesRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *EnumFilePropertiesRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EnumFilePropertiesOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// EnumFilePropertiesResponse structure represents the EnumFileProperties operation response
type EnumFilePropertiesResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// fileProperties: Pointer to IFsrmCollection interface pointer (section 3.2.4.2.1)
	// that contains a collection of file Property Definition Instances. Each item in the
	// collection is a VARIANT of type VT_DISPATCH. Query the pdispVal member of the variant
	// for the IFsrmProperty interface (section 3.2.4.2.40).
	FileProperties *fsrm.Collection `idl:"name:fileProperties" json:"file_properties"`
	// Return: The EnumFileProperties return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *EnumFilePropertiesResponse) xxx_ToOp(ctx context.Context) *xxx_EnumFilePropertiesOperation {
	if o == nil {
		return &xxx_EnumFilePropertiesOperation{}
	}
	return &xxx_EnumFilePropertiesOperation{
		That:           o.That,
		FileProperties: o.FileProperties,
		Return:         o.Return,
	}
}

func (o *EnumFilePropertiesResponse) xxx_FromOp(ctx context.Context, op *xxx_EnumFilePropertiesOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.FileProperties = op.FileProperties
	o.Return = op.Return
}
func (o *EnumFilePropertiesResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *EnumFilePropertiesResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EnumFilePropertiesOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetFilePropertyOperation structure represents the GetFileProperty operation
type xxx_GetFilePropertyOperation struct {
	This         *dcom.ORPCThis              `idl:"name:This" json:"this"`
	That         *dcom.ORPCThat              `idl:"name:That" json:"that"`
	FilePath     *oaut.String                `idl:"name:filePath" json:"file_path"`
	PropertyName *oaut.String                `idl:"name:propertyName" json:"property_name"`
	Options      fsrm.GetFilePropertyOptions `idl:"name:options" json:"options"`
	Property     *fsrm.Property              `idl:"name:property" json:"property"`
	Return       int32                       `idl:"name:Return" json:"return"`
}

func (o *xxx_GetFilePropertyOperation) OpNum() int { return 31 }

func (o *xxx_GetFilePropertyOperation) OpName() string {
	return "/IFsrmClassificationManager/v0/GetFileProperty"
}

func (o *xxx_GetFilePropertyOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetFilePropertyOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// filePath {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.FilePath != nil {
			_ptr_filePath := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.FilePath != nil {
					if err := o.FilePath.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.FilePath, _ptr_filePath); err != nil {
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
	// propertyName {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.PropertyName != nil {
			_ptr_propertyName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.PropertyName != nil {
					if err := o.PropertyName.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.PropertyName, _ptr_propertyName); err != nil {
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
	// options {in, default_value={0}} (1:{alias=FsrmGetFilePropertyOptions}(enum))
	{
		if err := w.WriteData(uint16(o.Options)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetFilePropertyOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// filePath {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_filePath := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.FilePath == nil {
				o.FilePath = &oaut.String{}
			}
			if err := o.FilePath.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_filePath := func(ptr interface{}) { o.FilePath = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.FilePath, _s_filePath, _ptr_filePath); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// propertyName {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_propertyName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.PropertyName == nil {
				o.PropertyName = &oaut.String{}
			}
			if err := o.PropertyName.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_propertyName := func(ptr interface{}) { o.PropertyName = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.PropertyName, _s_propertyName, _ptr_propertyName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// options {in, default_value={0}} (1:{alias=FsrmGetFilePropertyOptions}(enum))
	{
		if err := w.ReadData((*uint16)(&o.Options)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetFilePropertyOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetFilePropertyOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// property {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IFsrmProperty}(interface))
	{
		if o.Property != nil {
			_ptr_property := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Property != nil {
					if err := o.Property.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&fsrm.Property{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Property, _ptr_property); err != nil {
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
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetFilePropertyOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// property {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IFsrmProperty}(interface))
	{
		_ptr_property := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Property == nil {
				o.Property = &fsrm.Property{}
			}
			if err := o.Property.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_property := func(ptr interface{}) { o.Property = *ptr.(**fsrm.Property) }
		if err := w.ReadPointer(&o.Property, _s_property, _ptr_property); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// GetFilePropertyRequest structure represents the GetFileProperty operation request
type GetFilePropertyRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// filePath: The file or folder that contains the Property Definition Instance that
	// you want to retrieve. You can specify an absolute or relative path to the file or
	// folder. You cannot specify a file share.
	FilePath *oaut.String `idl:"name:filePath" json:"file_path"`
	// propertyName: Contains the name of the Property Definition Instance to retrieve.
	PropertyName *oaut.String `idl:"name:propertyName" json:"property_name"`
	// options: Contains the option to use for retrieving the file's Property Definition
	// Instance. For possible values, see the FsrmGetFilePropertyOptions (section 2.2.2.5.1.2)
	// enumeration.
	Options fsrm.GetFilePropertyOptions `idl:"name:options" json:"options"`
}

func (o *GetFilePropertyRequest) xxx_ToOp(ctx context.Context) *xxx_GetFilePropertyOperation {
	if o == nil {
		return &xxx_GetFilePropertyOperation{}
	}
	return &xxx_GetFilePropertyOperation{
		This:         o.This,
		FilePath:     o.FilePath,
		PropertyName: o.PropertyName,
		Options:      o.Options,
	}
}

func (o *GetFilePropertyRequest) xxx_FromOp(ctx context.Context, op *xxx_GetFilePropertyOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.FilePath = op.FilePath
	o.PropertyName = op.PropertyName
	o.Options = op.Options
}
func (o *GetFilePropertyRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetFilePropertyRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetFilePropertyOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetFilePropertyResponse structure represents the GetFileProperty operation response
type GetFilePropertyResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// property: Pointer to an IFsrmProperty interface pointer (section 3.2.4.2.40) that
	// contains the retrieved Property Definition Instance.
	Property *fsrm.Property `idl:"name:property" json:"property"`
	// Return: The GetFileProperty return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetFilePropertyResponse) xxx_ToOp(ctx context.Context) *xxx_GetFilePropertyOperation {
	if o == nil {
		return &xxx_GetFilePropertyOperation{}
	}
	return &xxx_GetFilePropertyOperation{
		That:     o.That,
		Property: o.Property,
		Return:   o.Return,
	}
}

func (o *GetFilePropertyResponse) xxx_FromOp(ctx context.Context, op *xxx_GetFilePropertyOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Property = op.Property
	o.Return = op.Return
}
func (o *GetFilePropertyResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetFilePropertyResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetFilePropertyOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetFilePropertyOperation structure represents the SetFileProperty operation
type xxx_SetFilePropertyOperation struct {
	This          *dcom.ORPCThis `idl:"name:This" json:"this"`
	That          *dcom.ORPCThat `idl:"name:That" json:"that"`
	FilePath      *oaut.String   `idl:"name:filePath" json:"file_path"`
	PropertyName  *oaut.String   `idl:"name:propertyName" json:"property_name"`
	PropertyValue *oaut.String   `idl:"name:propertyValue" json:"property_value"`
	Return        int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetFilePropertyOperation) OpNum() int { return 32 }

func (o *xxx_SetFilePropertyOperation) OpName() string {
	return "/IFsrmClassificationManager/v0/SetFileProperty"
}

func (o *xxx_SetFilePropertyOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetFilePropertyOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// filePath {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.FilePath != nil {
			_ptr_filePath := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.FilePath != nil {
					if err := o.FilePath.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.FilePath, _ptr_filePath); err != nil {
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
	// propertyName {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.PropertyName != nil {
			_ptr_propertyName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.PropertyName != nil {
					if err := o.PropertyName.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.PropertyName, _ptr_propertyName); err != nil {
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
	// propertyValue {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.PropertyValue != nil {
			_ptr_propertyValue := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.PropertyValue != nil {
					if err := o.PropertyValue.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.PropertyValue, _ptr_propertyValue); err != nil {
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

func (o *xxx_SetFilePropertyOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// filePath {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_filePath := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.FilePath == nil {
				o.FilePath = &oaut.String{}
			}
			if err := o.FilePath.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_filePath := func(ptr interface{}) { o.FilePath = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.FilePath, _s_filePath, _ptr_filePath); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// propertyName {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_propertyName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.PropertyName == nil {
				o.PropertyName = &oaut.String{}
			}
			if err := o.PropertyName.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_propertyName := func(ptr interface{}) { o.PropertyName = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.PropertyName, _s_propertyName, _ptr_propertyName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// propertyValue {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_propertyValue := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.PropertyValue == nil {
				o.PropertyValue = &oaut.String{}
			}
			if err := o.PropertyValue.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_propertyValue := func(ptr interface{}) { o.PropertyValue = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.PropertyValue, _s_propertyValue, _ptr_propertyValue); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetFilePropertyOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetFilePropertyOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetFilePropertyOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// SetFilePropertyRequest structure represents the SetFileProperty operation request
type SetFilePropertyRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// filePath: The file or folder that contains the Property Definition Instance that
	// you want to set. You can specify an absolute or relative path to the file or folder.
	// You cannot specify a file share.
	FilePath *oaut.String `idl:"name:filePath" json:"file_path"`
	// propertyName: Contains the name of the Property Definition Instance whose value you
	// want to set.
	PropertyName *oaut.String `idl:"name:propertyName" json:"property_name"`
	// propertyValue: Contains the value to set the Property Definition Instance to.
	PropertyValue *oaut.String `idl:"name:propertyValue" json:"property_value"`
}

func (o *SetFilePropertyRequest) xxx_ToOp(ctx context.Context) *xxx_SetFilePropertyOperation {
	if o == nil {
		return &xxx_SetFilePropertyOperation{}
	}
	return &xxx_SetFilePropertyOperation{
		This:          o.This,
		FilePath:      o.FilePath,
		PropertyName:  o.PropertyName,
		PropertyValue: o.PropertyValue,
	}
}

func (o *SetFilePropertyRequest) xxx_FromOp(ctx context.Context, op *xxx_SetFilePropertyOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.FilePath = op.FilePath
	o.PropertyName = op.PropertyName
	o.PropertyValue = op.PropertyValue
}
func (o *SetFilePropertyRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *SetFilePropertyRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetFilePropertyOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetFilePropertyResponse structure represents the SetFileProperty operation response
type SetFilePropertyResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The SetFileProperty return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetFilePropertyResponse) xxx_ToOp(ctx context.Context) *xxx_SetFilePropertyOperation {
	if o == nil {
		return &xxx_SetFilePropertyOperation{}
	}
	return &xxx_SetFilePropertyOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *SetFilePropertyResponse) xxx_FromOp(ctx context.Context, op *xxx_SetFilePropertyOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetFilePropertyResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *SetFilePropertyResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetFilePropertyOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_ClearFilePropertyOperation structure represents the ClearFileProperty operation
type xxx_ClearFilePropertyOperation struct {
	This     *dcom.ORPCThis `idl:"name:This" json:"this"`
	That     *dcom.ORPCThat `idl:"name:That" json:"that"`
	FilePath *oaut.String   `idl:"name:filePath" json:"file_path"`
	Property *oaut.String   `idl:"name:property" json:"property"`
	Return   int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_ClearFilePropertyOperation) OpNum() int { return 33 }

func (o *xxx_ClearFilePropertyOperation) OpName() string {
	return "/IFsrmClassificationManager/v0/ClearFileProperty"
}

func (o *xxx_ClearFilePropertyOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ClearFilePropertyOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// filePath {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.FilePath != nil {
			_ptr_filePath := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.FilePath != nil {
					if err := o.FilePath.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.FilePath, _ptr_filePath); err != nil {
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
	// property {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.Property != nil {
			_ptr_property := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Property != nil {
					if err := o.Property.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Property, _ptr_property); err != nil {
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

func (o *xxx_ClearFilePropertyOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// filePath {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_filePath := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.FilePath == nil {
				o.FilePath = &oaut.String{}
			}
			if err := o.FilePath.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_filePath := func(ptr interface{}) { o.FilePath = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.FilePath, _s_filePath, _ptr_filePath); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// property {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_property := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Property == nil {
				o.Property = &oaut.String{}
			}
			if err := o.Property.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_property := func(ptr interface{}) { o.Property = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.Property, _s_property, _ptr_property); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ClearFilePropertyOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ClearFilePropertyOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ClearFilePropertyOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// ClearFilePropertyRequest structure represents the ClearFileProperty operation request
type ClearFilePropertyRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// filePath: The file or folder that contains the Property Definition Instance that
	// you want to remove. You can specify an absolute or relative path to the file or folder.
	// You cannot specify a file share.
	FilePath *oaut.String `idl:"name:filePath" json:"file_path"`
	// property: Contains the name of the Property Definition Instance to remove from the
	// file.
	Property *oaut.String `idl:"name:property" json:"property"`
}

func (o *ClearFilePropertyRequest) xxx_ToOp(ctx context.Context) *xxx_ClearFilePropertyOperation {
	if o == nil {
		return &xxx_ClearFilePropertyOperation{}
	}
	return &xxx_ClearFilePropertyOperation{
		This:     o.This,
		FilePath: o.FilePath,
		Property: o.Property,
	}
}

func (o *ClearFilePropertyRequest) xxx_FromOp(ctx context.Context, op *xxx_ClearFilePropertyOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.FilePath = op.FilePath
	o.Property = op.Property
}
func (o *ClearFilePropertyRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *ClearFilePropertyRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ClearFilePropertyOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ClearFilePropertyResponse structure represents the ClearFileProperty operation response
type ClearFilePropertyResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The ClearFileProperty return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *ClearFilePropertyResponse) xxx_ToOp(ctx context.Context) *xxx_ClearFilePropertyOperation {
	if o == nil {
		return &xxx_ClearFilePropertyOperation{}
	}
	return &xxx_ClearFilePropertyOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *ClearFilePropertyResponse) xxx_FromOp(ctx context.Context, op *xxx_ClearFilePropertyOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *ClearFilePropertyResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *ClearFilePropertyResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ClearFilePropertyOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
