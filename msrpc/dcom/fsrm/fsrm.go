// The fsrm package implements the FSRM client protocol.
//
// # Introduction
//
// The File Server Resource Manager (FSRM) Protocol is a set of DCOM interfaces for
// managing the configuration of directory quotas, file screens, classification properties,
// classification rules, file management jobs, report jobs, classifier modules, and
// storage modules on a machine. The File Server Resource Manager Protocol deals with
// operating system, file system, and storage concepts. Although the basic concepts
// are outlined in this specification, the specification assumes that the reader has
// familiarity with these technologies. For background information about storage, disk,
// and volume concepts, see [MSDN-STC], [MSDN-DISKMAN], and [MSDN-PARTITIONINFO].
//
// This protocol is used to programmatically enumerate and configure directory quotas,
// file screens, report jobs, classifier modules, and storage modules on local and remote
// machines.<1>
//
// # Overview
//
// Using the File Server Resource Manager (FSRM) Protocol, a client can perform the
// following operations:
//
// * Limit the size of a given directory through directory quotas ( 809180e0-571e-40a0-8067-ffe89cbdcc40#gt_4b48a792-493d-4fa8-9ba8-175d43fb471d
// ).
//
// * Restrict the type of data that can be stored under a given directory through file
// screens ( 809180e0-571e-40a0-8067-ffe89cbdcc40#gt_4029e82b-3938-41a4-9a52-dff21ec37dcb
// ).
//
// * Define a property schema ( 809180e0-571e-40a0-8067-ffe89cbdcc40#gt_b2a5fac3-064b-4f84-9679-74dfb2ce5d2e
// ) that can be used to label files stored on the server.
//
// * Retrieve and modify the values assigned to classification properties for files
// stored on the server.
//
// * Configure automatic mechanisms to assign values to classification properties.
//
// * Register classification modules ( 809180e0-571e-40a0-8067-ffe89cbdcc40#gt_19e3eb54-a998-44af-b5a8-d0b02c57448c
// ) to alter the behavior of how files are classified and properties stored.
//
// * Register storage modules ( 809180e0-571e-40a0-8067-ffe89cbdcc40#gt_de2820bc-dc8b-4abe-8ea4-6ad86a6dc34f
// ) to alter the behavior of how the properties of a file are stored.
//
// * Apply policy to subsets of files.
//
// * Analyze storage utilization through report jobs ( 809180e0-571e-40a0-8067-ffe89cbdcc40#gt_15957e50-08ab-48d2-8824-57a71b23f833
// ).
//
// The FSRM protocol is expressed as a set of DCOM interfaces. The FSRM server implements
// support for the DCOM interface to manage FSRM objects. An FSRM client invokes method
// calls on the interface to perform various FSRM object configuration tasks on the
// server. More specifically, this protocol can be used for the following purposes:
//
// * Creating, enumerating, modifying, and deleting directory quotas and related objects
// ( auto apply quotas ( 809180e0-571e-40a0-8067-ffe89cbdcc40#gt_81039d49-4277-4f11-8251-7f1ab55eef77
// ) and quota templates ( 809180e0-571e-40a0-8067-ffe89cbdcc40#gt_1860c2ff-a8be-4088-988c-d03b43e9e006
// ) ) on the FSRM server.
//
// * Creating, enumerating, modifying, and deleting file screens and related objects
// ( file screen exceptions ( 809180e0-571e-40a0-8067-ffe89cbdcc40#gt_285a588e-ca67-4d67-90ec-c182c60c9d66
// ) and file groups ( 809180e0-571e-40a0-8067-ffe89cbdcc40#gt_1f828be2-7a1b-4671-b82d-76b974874edc
// ) ) on the FSRM server.
//
// * Creating, enumerating, modifying, and deleting classification properties on the
// FSRM server.
//
// * Setting, enumerating, modifying, and deleting properties values ( 809180e0-571e-40a0-8067-ffe89cbdcc40#gt_13091c0e-746f-478c-b54c-ad0294a94ea8
// ) for specific files on the FSRM server.
//
// * Creating, enumerating, modifying, and deleting classification rules ( 809180e0-571e-40a0-8067-ffe89cbdcc40#gt_4bd4bf29-9046-4cb8-b637-fd107cd756aa
// ) on the FSRM server.
//
// * Creating, enumerating, modifying, and deleting classification modules on the FSRM
// server.
//
// * Creating, enumerating, modifying, and deleting storage modules on the FSRM server.
//
// * Creating, enumerating, modifying, and deleting file management jobs ( 809180e0-571e-40a0-8067-ffe89cbdcc40#gt_d97f21a8-2a20-45fd-8e7e-ec697810ac45
// ) on the FSRM server.
//
// * Creating, enumerating, modifying, and deleting report jobs on the FSRM server.
//
// * Querying and setting FSRM server general settings; for example, the Simple Mail
// Transfer Protocol [RFC821] ( https://go.microsoft.com/fwlink/?LinkId=90496 ) server
// name and report default parameters.
//
// A typical FSRM session involves a client connecting to the server and requesting
// an interface that allows performing high-level operations, such as enumeration and
// creation for a class of FSRM objects. If the server accepts the request, it responds
// with the requested interface. The client can then use the interface to request that
// the server enumerate the objects of the desired class. If the server accepts the
// request, it responds with a collection of interfaces that allow access to the requested
// type of FSRM object. The client uses the interfaces returned by the server to send
// additional requests to the server specifying the type of operation to perform and
// any operation-specific parameters. If the server accepts the operation request, it
// attempts to query or change the state of the corresponding FSRM object based on the
// request parameters and returns to the client the result of the operation. To persist
// changes to the manipulated FSRM objects, the client can explicitly request that the
// server commit any outstanding changes.
//
// The following are FSRM objects:
package fsrm

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
	GoPackage = "dcom/fsrm"
)

// DispatchIDFeatureMask represents the FSRM_DISPID_FEATURE_MASK RPC constant
var DispatchIDFeatureMask = 251658240

// DispatchIDInterfaceAMask represents the FSRM_DISPID_INTERFACE_A_MASK RPC constant
var DispatchIDInterfaceAMask = 15728640

// DispatchIDInterfaceBMask represents the FSRM_DISPID_INTERFACE_B_MASK RPC constant
var DispatchIDInterfaceBMask = 983040

// DispatchIDInterfaceCMask represents the FSRM_DISPID_INTERFACE_C_MASK RPC constant
var DispatchIDInterfaceCMask = 61440

// DispatchIDInterfaceDMask represents the FSRM_DISPID_INTERFACE_D_MASK RPC constant
var DispatchIDInterfaceDMask = 3840

// DispatchIDInterfaceMask represents the FSRM_DISPID_INTERFACE_MASK RPC constant
var DispatchIDInterfaceMask = 16776960

// DispatchIDIsProperty represents the FSRM_DISPID_IS_PROPERTY RPC constant
var DispatchIDIsProperty = 128

// DispatchIDMethodNumMask represents the FSRM_DISPID_METHOD_NUM_MASK RPC constant
var DispatchIDMethodNumMask = 127

// DispatchIDMethodMask represents the FSRM_DISPID_METHOD_MASK RPC constant
var DispatchIDMethodMask = 255

// DispatchIDFeatureGeneral represents the FSRM_DISPID_FEATURE_GENERAL RPC constant
var DispatchIDFeatureGeneral = 16777216

// DispatchIDFeatureQuota represents the FSRM_DISPID_FEATURE_QUOTA RPC constant
var DispatchIDFeatureQuota = 33554432

// DispatchIDFeatureFilescreen represents the FSRM_DISPID_FEATURE_FILESCREEN RPC constant
var DispatchIDFeatureFilescreen = 50331648

// DispatchIDFeatureReports represents the FSRM_DISPID_FEATURE_REPORTS RPC constant
var DispatchIDFeatureReports = 67108864

// DispatchIDFeatureClassification represents the FSRM_DISPID_FEATURE_CLASSIFICATION RPC constant
var DispatchIDFeatureClassification = 83886080

// DispatchIDFeaturePipeline represents the FSRM_DISPID_FEATURE_PIPELINE RPC constant
var DispatchIDFeaturePipeline = 100663296

// DispatchIDObject represents the FSRM_DISPID_OBJECT RPC constant
var DispatchIDObject = 17825792

// DispatchIDCollection represents the FSRM_DISPID_COLLECTION RPC constant
var DispatchIDCollection = 18874368

// DispatchIDCollectionMutable represents the FSRM_DISPID_COLLECTION_MUTABLE RPC constant
var DispatchIDCollectionMutable = 18939904

// DispatchIDCollectionCommittable represents the FSRM_DISPID_COLLECTION_COMMITTABLE RPC constant
var DispatchIDCollectionCommittable = 18944000

// DispatchIDAction represents the FSRM_DISPID_ACTION RPC constant
var DispatchIDAction = 19922944

// DispatchIDActionEmail represents the FSRM_DISPID_ACTION_EMAIL RPC constant
var DispatchIDActionEmail = 19988480

// DispatchIDActionReport represents the FSRM_DISPID_ACTION_REPORT RPC constant
var DispatchIDActionReport = 20054016

// DispatchIDActionEventlog represents the FSRM_DISPID_ACTION_EVENTLOG RPC constant
var DispatchIDActionEventlog = 20119552

// DispatchIDActionCommand represents the FSRM_DISPID_ACTION_COMMAND RPC constant
var DispatchIDActionCommand = 20185088

// DispatchIDActionEmail2 represents the FSRM_DISPID_ACTION_EMAIL2 RPC constant
var DispatchIDActionEmail2 = 20250624

// DispatchIDSetting represents the FSRM_DISPID_SETTING RPC constant
var DispatchIDSetting = 20971520

// DispatchIDPathmapper represents the FSRM_DISPID_PATHMAPPER RPC constant
var DispatchIDPathmapper = 22020096

// DispatchIDDerivedobjectsresult represents the FSRM_DISPID_DERIVEDOBJECTSRESULT RPC constant
var DispatchIDDerivedobjectsresult = 24117248

// DispatchIDPropertyDefinition represents the FSRM_DISPID_PROPERTY_DEFINITION RPC constant
var DispatchIDPropertyDefinition = 84934656

// DispatchIDPropertyDefinition2 represents the FSRM_DISPID_PROPERTY_DEFINITION2 RPC constant
var DispatchIDPropertyDefinition2 = 85000192

// DispatchIDProperty represents the FSRM_DISPID_PROPERTY RPC constant
var DispatchIDProperty = 85983232

// DispatchIDRule represents the FSRM_DISPID_RULE RPC constant
var DispatchIDRule = 87031808

// DispatchIDClassificationRule represents the FSRM_DISPID_CLASSIFICATION_RULE RPC constant
var DispatchIDClassificationRule = 87097344

// DispatchIDExpirationRule represents the FSRM_DISPID_EXPIRATION_RULE RPC constant
var DispatchIDExpirationRule = 87162880

// DispatchIDPipelineModuleDefinition represents the FSRM_DISPID_PIPELINE_MODULE_DEFINITION RPC constant
var DispatchIDPipelineModuleDefinition = 88080384

// DispatchIDClassifierModuleDefinition represents the FSRM_DISPID_CLASSIFIER_MODULE_DEFINITION RPC constant
var DispatchIDClassifierModuleDefinition = 88145920

// DispatchIDStorageModuleDefinition represents the FSRM_DISPID_STORAGE_MODULE_DEFINITION RPC constant
var DispatchIDStorageModuleDefinition = 88211456

// DispatchIDClassificationManager represents the FSRM_DISPID_CLASSIFICATION_MANAGER RPC constant
var DispatchIDClassificationManager = 89128960

// DispatchIDClassificationManager2 represents the FSRM_DISPID_CLASSIFICATION_MANAGER2 RPC constant
var DispatchIDClassificationManager2 = 89194496

// DispatchIDClassificationEvents represents the FSRM_DISPID_CLASSIFICATION_EVENTS RPC constant
var DispatchIDClassificationEvents = 90177536

// DispatchIDPropertyDefinitionValue represents the FSRM_DISPID_PROPERTY_DEFINITION_VALUE RPC constant
var DispatchIDPropertyDefinitionValue = 91226112

// DispatchIDPropertyBag represents the FSRM_DISPID_PROPERTY_BAG RPC constant
var DispatchIDPropertyBag = 106954752

// DispatchIDPipelineModuleImplementation represents the FSRM_DISPID_PIPELINE_MODULE_IMPLEMENTATION RPC constant
var DispatchIDPipelineModuleImplementation = 108003328

// DispatchIDPipelineModuleConnector represents the FSRM_DISPID_PIPELINE_MODULE_CONNECTOR RPC constant
var DispatchIDPipelineModuleConnector = 109051904

// DispatchIDPipelineModuleHost represents the FSRM_DISPID_PIPELINE_MODULE_HOST RPC constant
var DispatchIDPipelineModuleHost = 110100480

// MaxNumberPropertyDefinitions represents the FsrmMaxNumberPropertyDefinitions RPC constant
const MaxNumberPropertyDefinitions = 0x000000C8

// DispatchIDQuotaBase represents the FSRM_DISPID_QUOTA_BASE RPC constant
var DispatchIDQuotaBase = 34603008

// DispatchIDQuotaObject represents the FSRM_DISPID_QUOTA_OBJECT RPC constant
var DispatchIDQuotaObject = 34668544

// DispatchIDQuota represents the FSRM_DISPID_QUOTA RPC constant
var DispatchIDQuota = 34672640

// DispatchIDAutoapplyquota represents the FSRM_DISPID_AUTOAPPLYQUOTA RPC constant
var DispatchIDAutoapplyquota = 34676736

// DispatchIDQuotaTemplate represents the FSRM_DISPID_QUOTA_TEMPLATE RPC constant
var DispatchIDQuotaTemplate = 34734080

// DispatchIDQuotaTemplateImported represents the FSRM_DISPID_QUOTA_TEMPLATE_IMPORTED RPC constant
var DispatchIDQuotaTemplateImported = 34738176

// DispatchIDQuotaManager represents the FSRM_DISPID_QUOTA_MANAGER RPC constant
var DispatchIDQuotaManager = 35651584

// DispatchIDQuotaTemplateManager represents the FSRM_DISPID_QUOTA_TEMPLATE_MANAGER RPC constant
var DispatchIDQuotaTemplateManager = 36700160

// DispatchIDQuotaManagerEx represents the FSRM_DISPID_QUOTA_MANAGER_EX RPC constant
var DispatchIDQuotaManagerEx = 37748736

// MaxNumberThresholds represents the FsrmMaxNumberThresholds RPC constant
const MaxNumberThresholds = 0x00000010

// MinThresholdValue represents the FsrmMinThresholdValue RPC constant
const MinThresholdValue = 0x00000001

// MaxThresholdValue represents the FsrmMaxThresholdValue RPC constant
const MaxThresholdValue = 0x000000FA

// MinQuotaLimit represents the FsrmMinQuotaLimit RPC constant
const MinQuotaLimit = 0x00000400

// MaxExcludeFolders represents the FsrmMaxExcludeFolders RPC constant
const MaxExcludeFolders = 0x00000020

// DispatchIDReportManager represents the FSRM_DISPID_REPORT_MANAGER RPC constant
var DispatchIDReportManager = 68157440

// DispatchIDReportJob represents the FSRM_DISPID_REPORT_JOB RPC constant
var DispatchIDReportJob = 69206016

// DispatchIDReport represents the FSRM_DISPID_REPORT RPC constant
var DispatchIDReport = 70254592

// DispatchIDReportScheduler represents the FSRM_DISPID_REPORT_SCHEDULER RPC constant
var DispatchIDReportScheduler = 71303168

// DispatchIDFileManagementJobManager represents the FSRM_DISPID_FILE_MANAGEMENT_JOB_MANAGER RPC constant
var DispatchIDFileManagementJobManager = 72351744

// DispatchIDFileManagementJob represents the FSRM_DISPID_FILE_MANAGEMENT_JOB RPC constant
var DispatchIDFileManagementJob = 73400320

// DispatchIDFileManagementNotification represents the FSRM_DISPID_FILE_MANAGEMENT_NOTIFICATION RPC constant
var DispatchIDFileManagementNotification = 74448896

// DispatchIDPropertyCondition represents the FSRM_DISPID_PROPERTY_CONDITION RPC constant
var DispatchIDPropertyCondition = 75497472

// DispatchIDFilegroup represents the FSRM_DISPID_FILEGROUP RPC constant
var DispatchIDFilegroup = 51380224

// DispatchIDFilegroupImported represents the FSRM_DISPID_FILEGROUP_IMPORTED RPC constant
var DispatchIDFilegroupImported = 51445760

// DispatchIDFilegroupManager represents the FSRM_DISPID_FILEGROUP_MANAGER RPC constant
var DispatchIDFilegroupManager = 52428800

// DispatchIDFilescreenBase represents the FSRM_DISPID_FILESCREEN_BASE RPC constant
var DispatchIDFilescreenBase = 53477376

// DispatchIDFilescreen represents the FSRM_DISPID_FILESCREEN RPC constant
var DispatchIDFilescreen = 53542912

// DispatchIDFilescreenTemplate represents the FSRM_DISPID_FILESCREEN_TEMPLATE RPC constant
var DispatchIDFilescreenTemplate = 53608448

// DispatchIDFilescreenTemplateImported represents the FSRM_DISPID_FILESCREEN_TEMPLATE_IMPORTED RPC constant
var DispatchIDFilescreenTemplateImported = 53612544

// DispatchIDFilescreenException represents the FSRM_DISPID_FILESCREEN_EXCEPTION RPC constant
var DispatchIDFilescreenException = 54525952

// DispatchIDFilescreenManager represents the FSRM_DISPID_FILESCREEN_MANAGER RPC constant
var DispatchIDFilescreenManager = 55574528

// DispatchIDFilescreenTemplateManager represents the FSRM_DISPID_FILESCREEN_TEMPLATE_MANAGER RPC constant
var DispatchIDFilescreenTemplateManager = 56623104

// ObjectID structure represents FSRM_OBJECT_ID RPC structure.
type ObjectID dtyp.GUID

func (o *ObjectID) GUID() *dtyp.GUID { return (*dtyp.GUID)(o) }

func (o *ObjectID) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *ObjectID) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.Data1); err != nil {
		return err
	}
	if err := w.WriteData(o.Data2); err != nil {
		return err
	}
	if err := w.WriteData(o.Data3); err != nil {
		return err
	}
	for i1 := range o.Data4 {
		i1 := i1
		if uint64(i1) >= 8 {
			break
		}
		if err := w.WriteData(o.Data4[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.Data4); uint64(i1) < 8; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	if err := w.WriteTrailingGap(4); err != nil {
		return err
	}
	return nil
}
func (o *ObjectID) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.Data1); err != nil {
		return err
	}
	if err := w.ReadData(&o.Data2); err != nil {
		return err
	}
	if err := w.ReadData(&o.Data3); err != nil {
		return err
	}
	o.Data4 = make([]byte, 8)
	for i1 := range o.Data4 {
		i1 := i1
		if err := w.ReadData(&o.Data4[i1]); err != nil {
			return err
		}
	}
	if err := w.ReadTrailingGap(4); err != nil {
		return err
	}
	return nil
}

// QuotaFlags type represents FsrmQuotaFlags RPC enumeration.
//
// The FsrmQuotaFlags enumeration defines bitmasks for the possible states of the quota
// objects in the File Server Resource Manager Protocol.
type QuotaFlags uint32

var (
	// FsrmQuotaFlags_Enforce:  If the FsrmQuotaFlags_Enforce bitmask is set as part of
	// the QuotaFlags property of an IFsrmQuotaBase object, the server fails an I/O operation
	// that causes the disk space usage to exceed the quota limit. If not set, the server
	// does not fail operations that violate the I/O limit, but still run actions associated
	// with the quota limit.
	QuotaFlagsEnforce QuotaFlags = 256
	// FsrmQuotaFlags_Disable:  If the FsrmQuotaFlags_Disable bitmask is set as part of
	// the QuotaFlags property of an IFsrmQuotaBase object, the server does not track quota
	// data for the quota and does not run any actions associated with quota thresholds.
	QuotaFlagsDisable QuotaFlags = 512
	// FsrmQuotaFlags_StatusIncomplete:  If the FsrmQuotaFlags_StatusIncomplete bitmask
	// is set as part of the QuotaFlags property of an IFsrmQuotaBase object, a quota is
	// defined on the server but the rebuilding procedure has not yet started.
	QuotaFlagsStatusIncomplete QuotaFlags = 65536
	// FsrmQuotaFlags_StatusRebuilding:  If the FsrmQuotaFlags_StatusRebuilding bitmask
	// is set as part of the QuotaFlags property of an IFsrmQuotaBase object, a quota is
	// rebuilding its data from the disk.
	QuotaFlagsStatusRebuilding QuotaFlags = 131072
)

func (o QuotaFlags) String() string {
	switch o {
	case QuotaFlagsEnforce:
		return "QuotaFlagsEnforce"
	case QuotaFlagsDisable:
		return "QuotaFlagsDisable"
	case QuotaFlagsStatusIncomplete:
		return "QuotaFlagsStatusIncomplete"
	case QuotaFlagsStatusRebuilding:
		return "QuotaFlagsStatusRebuilding"
	}
	return "Invalid"
}

// FileScreenFlags type represents FsrmFileScreenFlags RPC enumeration.
//
// The FsrmFileScreenFlags enumeration defines bitmasks for possible states of the file
// screen objects in the File Server Resource Manager Protocol.
type FileScreenFlags uint16

var (
	// FsrmFileScreenFlags_Enforce:  If this bitmask is set as part of the fileScreenFlags
	// member of a IFsrmFileScreenBase object, the server fails any I/O operation that violates
	// the file screen. If not set, the server does not fail operations that violate the
	// file screen but still run actions associated with the file screen.
	FileScreenFlagsEnforce FileScreenFlags = 1
)

func (o FileScreenFlags) String() string {
	switch o {
	case FileScreenFlagsEnforce:
		return "FileScreenFlagsEnforce"
	}
	return "Invalid"
}

// CollectionState type represents FsrmCollectionState RPC enumeration.
//
// The FsrmCollectionState enumeration defines the possible states of collection objects
// in the File Server Resource Manager Protocol.
type CollectionState uint16

var (
	// FsrmCollectionState_Fetching:  The collection object is currently fetching data.
	CollectionStateFetching CollectionState = 1
	// FsrmCollectionState_Committing:  The collection object is currently committing its
	// data.
	CollectionStateCommitting CollectionState = 2
	// FsrmCollectionState_Complete:  The collection object is complete and has stopped
	// fetching or committing data.
	CollectionStateComplete  CollectionState = 3
	CollectionStateCancelled CollectionState = 4
)

func (o CollectionState) String() string {
	switch o {
	case CollectionStateFetching:
		return "CollectionStateFetching"
	case CollectionStateCommitting:
		return "CollectionStateCommitting"
	case CollectionStateComplete:
		return "CollectionStateComplete"
	case CollectionStateCancelled:
		return "CollectionStateCancelled"
	}
	return "Invalid"
}

// EnumOptions type represents FsrmEnumOptions RPC enumeration.
//
// The FsrmEnumOptions enumeration defines the different options that can be used when
// enumerating collections of File Server Resource Manager Protocol objects.
type EnumOptions uint16

var (
	// FsrmEnumOptions_None:  Use no options and enumerate objects synchronously.
	EnumOptionsNone EnumOptions = 0
	// FsrmEnumOptions_Asynchronous:  Enumerate the objects asynchronously.
	EnumOptionsAsynchronous EnumOptions = 1
	// FsrmEnumOptions_CheckRecycleBin:  Include items that are in the Recycle Bin when
	// enumerating. This will include files that are located in a folder that has "$RECYCLE.BIN"
	// in its path regardless of capitalization. Without this option, those files will be
	// excluded.
	EnumOptionsCheckRecycleBin EnumOptions = 2
	// FsrmEnumOptions_IncludeClusterNodes:  If the system is configured to be part of
	// a cluster, include all objects even if they are not currently available on the system
	// (identified by the machine name). Without this option, only objects available on
	// the current system will be included.
	EnumOptionsIncludeClusterNodes EnumOptions = 4
	// FsrmEnumOptions_IncludeDeprecatedObjects:  If any objects were marked Deprecated,
	// they will appear only when enumerated with this option.
	EnumOptionsIncludeDeprecatedObjects EnumOptions = 8
)

func (o EnumOptions) String() string {
	switch o {
	case EnumOptionsNone:
		return "EnumOptionsNone"
	case EnumOptionsAsynchronous:
		return "EnumOptionsAsynchronous"
	case EnumOptionsCheckRecycleBin:
		return "EnumOptionsCheckRecycleBin"
	case EnumOptionsIncludeClusterNodes:
		return "EnumOptionsIncludeClusterNodes"
	case EnumOptionsIncludeDeprecatedObjects:
		return "EnumOptionsIncludeDeprecatedObjects"
	}
	return "Invalid"
}

// CommitOptions type represents FsrmCommitOptions RPC enumeration.
//
// The FsrmCommitOptions enumeration defines the different options that can be used
// when committing a collection of File Server Resource Manager Protocol objects.
type CommitOptions uint16

var (
	// FsrmCommitOptions_None:  Use no options and commit the collection of objects synchronously.
	CommitOptionsNone CommitOptions = 0
	// FsrmCommitOptions_Asynchronous:  Commit the collection of objects asynchronously.
	CommitOptionsAsynchronous CommitOptions = 1
)

func (o CommitOptions) String() string {
	switch o {
	case CommitOptionsNone:
		return "CommitOptionsNone"
	case CommitOptionsAsynchronous:
		return "CommitOptionsAsynchronous"
	}
	return "Invalid"
}

// TemplateApplyOptions type represents FsrmTemplateApplyOptions RPC enumeration.
//
// The FsrmTemplateApplyOptions enumeration defines the different options that are available
// when applying changes that have been made to a template to the objects derived from
// that template.
type TemplateApplyOptions uint16

var (
	// FsrmTemplateApplyOptions_ApplyToDerivedMatching:  Apply template changes only to
	// derived objects whose properties match the template.
	TemplateApplyOptionsToDerivedMatching TemplateApplyOptions = 1
	// FsrmTemplateApplyOptions_ApplyToDerivedAll:  Apply template changes to all derived
	// objects, whether their properties match the template's or not.
	TemplateApplyOptionsToDerivedAll TemplateApplyOptions = 2
)

func (o TemplateApplyOptions) String() string {
	switch o {
	case TemplateApplyOptionsToDerivedMatching:
		return "TemplateApplyOptionsToDerivedMatching"
	case TemplateApplyOptionsToDerivedAll:
		return "TemplateApplyOptionsToDerivedAll"
	}
	return "Invalid"
}

// ActionType type represents FsrmActionType RPC enumeration.
//
// The FsrmActionType enumeration defines the set of the action types that can be triggered
// in response to a quota or file screen event.
type ActionType uint16

var (
	// FsrmActionType_Unknown:  This enumeration value is not used by FSRM and MUST NOT
	// be referenced. If the server receives this enumeration value, it MUST consider the
	// value invalid and not apply any changes.
	ActionTypeUnknown ActionType = 0
	// FsrmActionType_EventLog:  The action will log an event to the application event
	// log.
	ActionTypeEventLog ActionType = 1
	// FsrmActionType_Email:  The action will send an email.
	ActionTypeEmail ActionType = 2
	// FsrmActionType_Command:  The action will execute a command or script.
	ActionTypeCommand ActionType = 3
	// FsrmActionType_Report:  The action will generate a report.
	ActionTypeReport ActionType = 4
)

func (o ActionType) String() string {
	switch o {
	case ActionTypeUnknown:
		return "ActionTypeUnknown"
	case ActionTypeEventLog:
		return "ActionTypeEventLog"
	case ActionTypeEmail:
		return "ActionTypeEmail"
	case ActionTypeCommand:
		return "ActionTypeCommand"
	case ActionTypeReport:
		return "ActionTypeReport"
	}
	return "Invalid"
}

// EventType type represents FsrmEventType RPC enumeration.
//
// The FsrmEventType enumeration defines the set of event types that can be logged as
// part of an FsrmActionType_EventLog action.
type EventType uint16

var (
	// FsrmEventType_Unknown:  This enumeration value is not used by FSRM and MUST NOT
	// be referenced. If the server receives this enumeration value, it MUST consider the
	// value invalid and not apply any changes.
	EventTypeUnknown EventType = 0
	// FsrmEventType_Information:  The event is an information event.
	EventTypeInformation EventType = 1
	// FsrmEventType_Warning:  The event is a warning event.
	EventTypeWarning EventType = 2
	// FsrmEventType_Error:  The event is an error event.
	EventTypeError EventType = 3
)

func (o EventType) String() string {
	switch o {
	case EventTypeUnknown:
		return "EventTypeUnknown"
	case EventTypeInformation:
		return "EventTypeInformation"
	case EventTypeWarning:
		return "EventTypeWarning"
	case EventTypeError:
		return "EventTypeError"
	}
	return "Invalid"
}

// AccountType type represents FsrmAccountType RPC enumeration.
//
// The FsrmAccountType enumeration defines the set of machine account types under which
// an FsrmActionType_Command action or a module definition can be run.
type AccountType uint16

var (
	// FsrmAccountType_Unknown:  This enumeration value is not used by FSRM and MUST NOT
	// be referenced. If the server receives this enumeration value, it MUST consider the
	// value invalid and not apply any changes.
	AccountTypeUnknown AccountType = 0
	// FsrmAccountType_NetworkService:  Run the command or module definition under a restricted
	// account with network access (see [MSDN-NetworkSvcAcct] for more information).<10>
	AccountTypeNetworkService AccountType = 1
	// FsrmAccountType_LocalService:  Run the command or module definition under a restricted
	// account without network access (see [MSDN-LocSvcAcct] for more information).<11>
	AccountTypeLocalService AccountType = 2
	// FsrmAccountType_LocalSystem:  Run the command or module definition under an administrative
	// account with network access. See [MSDN-LocSysAcct] for more information.<12>
	AccountTypeLocalSystem AccountType = 3
	// FsrmAccountType_InProc:  Run the module definition in an administrative account
	// in the same process used for pipeline processing.<13>
	AccountTypeInProc AccountType = 4
	// FsrmAccountType_External:  Run the module definition in its own process.<14>
	AccountTypeExternal AccountType = 5
	// FsrmAccountType_Automatic:  Run the module definition in a process determined by
	// the server.<15>
	AccountTypeAutomatic AccountType = 500
)

func (o AccountType) String() string {
	switch o {
	case AccountTypeUnknown:
		return "AccountTypeUnknown"
	case AccountTypeNetworkService:
		return "AccountTypeNetworkService"
	case AccountTypeLocalService:
		return "AccountTypeLocalService"
	case AccountTypeLocalSystem:
		return "AccountTypeLocalSystem"
	case AccountTypeInProc:
		return "AccountTypeInProc"
	case AccountTypeExternal:
		return "AccountTypeExternal"
	case AccountTypeAutomatic:
		return "AccountTypeAutomatic"
	}
	return "Invalid"
}

// ReportType type represents FsrmReportType RPC enumeration.
//
// The FsrmReportType enumeration defines the set of report types that can be generated
// by the File Server Resource Manager Protocol.
type ReportType uint16

var (
	// FsrmReportType_Unknown:  This enumeration value is not used by FSRM and MUST NOT
	// be referenced. If the server receives this enumeration value, it MUST consider the
	// value invalid and not apply any changes.
	ReportTypeUnknown ReportType = 0
	// FsrmReportType_LargeFiles:  This report type lists files over a given size.
	ReportTypeLargeFiles ReportType = 1
	// FsrmReportType_FilesByType:  This report type lists files grouped by type.
	ReportTypeFilesByType ReportType = 2
	// FsrmReportType_LeastRecentlyAccessed:  This report type lists files that have not
	// been accessed recently.
	ReportTypeLeastRecentlyAccessed ReportType = 3
	// FsrmReportType_MostRecentlyAccessed:  This report type lists files that have been
	// accessed most recently.
	ReportTypeMostRecentlyAccessed ReportType = 4
	// FsrmReportType_QuotaUsage:  This report type lists quotas that exceed a certain
	// threshold.
	ReportTypeQuotaUsage ReportType = 5
	// FsrmReportType_FilesByOwner:  This report lists files grouped by their owner.
	ReportTypeFilesByOwner ReportType = 6
	// FsrmReportType_ExportReport:  This report lists files without any grouping or limiting.
	ReportTypeExportReport ReportType = 7
	// FsrmReportType_DuplicateFiles:  This report lists duplicate files.<16>
	ReportTypeDuplicateFiles ReportType = 8
	// FsrmReportType_FileScreenAudit:  This report lists file screening events that have
	// occurred.
	ReportTypeFileScreenAudit ReportType = 9
	// FsrmReportType_FilesByProperty:  This report lists files grouped by classification
	// property.
	ReportTypeFilesByProperty ReportType = 10
	// FsrmReportType_AutomaticClassification:  This report lists files that have been
	// classified during an automatic classification run.
	ReportTypeAutomaticClassification ReportType = 11
	// FsrmReportType_Expiration:  This report lists files that have been expired during
	// a file management operation.
	ReportTypeExpiration ReportType = 12
	// FsrmReportType_FoldersByProperty:  This report lists folders grouped by classification
	// property.
	ReportTypeFoldersByProperty ReportType = 13
)

func (o ReportType) String() string {
	switch o {
	case ReportTypeUnknown:
		return "ReportTypeUnknown"
	case ReportTypeLargeFiles:
		return "ReportTypeLargeFiles"
	case ReportTypeFilesByType:
		return "ReportTypeFilesByType"
	case ReportTypeLeastRecentlyAccessed:
		return "ReportTypeLeastRecentlyAccessed"
	case ReportTypeMostRecentlyAccessed:
		return "ReportTypeMostRecentlyAccessed"
	case ReportTypeQuotaUsage:
		return "ReportTypeQuotaUsage"
	case ReportTypeFilesByOwner:
		return "ReportTypeFilesByOwner"
	case ReportTypeExportReport:
		return "ReportTypeExportReport"
	case ReportTypeDuplicateFiles:
		return "ReportTypeDuplicateFiles"
	case ReportTypeFileScreenAudit:
		return "ReportTypeFileScreenAudit"
	case ReportTypeFilesByProperty:
		return "ReportTypeFilesByProperty"
	case ReportTypeAutomaticClassification:
		return "ReportTypeAutomaticClassification"
	case ReportTypeExpiration:
		return "ReportTypeExpiration"
	case ReportTypeFoldersByProperty:
		return "ReportTypeFoldersByProperty"
	}
	return "Invalid"
}

// ReportFormat type represents FsrmReportFormat RPC enumeration.
//
// The FsrmReportFormat enumeration defines the set of formats that the File Server
// Resource Manager Protocol can use when generating reports.
type ReportFormat uint16

var (
	// FsrmReportFormat_Unknown:  This enumeration value is not used by FSRM and MUST NOT
	// be referenced. If the server receives this enumeration value, it MUST consider the
	// value invalid and not apply any changes.
	ReportFormatUnknown ReportFormat = 0
	// FsrmReportFormat_DHtml:  The report is rendered in Dynamic Hypertext Markup Language
	// (DHTML).
	ReportFormatDHTML ReportFormat = 1
	// FsrmReportFormat_Html:  The report is rendered in HTML.
	ReportFormatHTML ReportFormat = 2
	// FsrmReportFormat_Txt:  The report is rendered as a text file.
	ReportFormatText ReportFormat = 3
	// FsrmReportFormat_Csv:  The report is rendered as a comma-separated value file.
	ReportFormatCSV ReportFormat = 4
	// FsrmReportFormat_Xml:  The report is rendered in XML.
	ReportFormatXML ReportFormat = 5
)

func (o ReportFormat) String() string {
	switch o {
	case ReportFormatUnknown:
		return "ReportFormatUnknown"
	case ReportFormatDHTML:
		return "ReportFormatDHTML"
	case ReportFormatHTML:
		return "ReportFormatHTML"
	case ReportFormatText:
		return "ReportFormatText"
	case ReportFormatCSV:
		return "ReportFormatCSV"
	case ReportFormatXML:
		return "ReportFormatXML"
	}
	return "Invalid"
}

// ReportRunningStatus type represents FsrmReportRunningStatus RPC enumeration.
//
// The FsrmReportRunningStatus enumeration defines the set of running states for a report,
// classification, or file management job.
type ReportRunningStatus uint16

var (
	// FsrmReportRunningStatus_Unknown:  This enumeration value is not used by FSRM and
	// MUST NOT be referenced. If the server receives this enumeration value, it MUST consider
	// the value invalid and not apply any changes.
	ReportRunningStatusUnknown ReportRunningStatus = 0
	// FsrmReportRunningStatus_NotRunning:  The report, classification, or file management
	// job is not running.
	ReportRunningStatusNotRunning ReportRunningStatus = 1
	// FsrmReportRunningStatus_Queued:  The request to run the Report job, Classification
	// job, or File Management Job has been made and an associated Running Job (section
	// 3.2.1.5.1.3), Running Classification job has been added to the Running Report Job
	// Queue, Running Classification Job Queue, or Running File Management Job Queue respectively,
	// but the task is not running at the moment.
	ReportRunningStatusQueued ReportRunningStatus = 2
	// FsrmReportRunningStatus_Running:  The Report job, Classification job, or File Management
	// Job is running.
	ReportRunningStatusRunning ReportRunningStatus = 3
)

func (o ReportRunningStatus) String() string {
	switch o {
	case ReportRunningStatusUnknown:
		return "ReportRunningStatusUnknown"
	case ReportRunningStatusNotRunning:
		return "ReportRunningStatusNotRunning"
	case ReportRunningStatusQueued:
		return "ReportRunningStatusQueued"
	case ReportRunningStatusRunning:
		return "ReportRunningStatusRunning"
	}
	return "Invalid"
}

// ReportGenerationContext type represents FsrmReportGenerationContext RPC enumeration.
//
// The FsrmReportGenerationContext enumeration defines the set of contexts under which
// a report is run.
type ReportGenerationContext uint16

var (
	// FsrmReportGenerationContext_Undefined:  This enumeration value is not used by FSRM
	// and MUST NOT be referenced. If the server receives this enumeration value, it MUST
	// consider the value invalid and not apply any changes.
	ReportGenerationContextUndefined ReportGenerationContext = 1
	// FsrmReportGenerationContext_ScheduledReport:  The report will run as a scheduled
	// report.
	ReportGenerationContextScheduledReport ReportGenerationContext = 2
	// FsrmReportGenerationContext_InteractiveReport:  The report will run on demand.
	ReportGenerationContextInteractiveReport ReportGenerationContext = 3
	// FsrmReportGenerationContext_IncidentReport:  The report will run in response to
	// a quota or file screen event.
	ReportGenerationContextIncidentReport ReportGenerationContext = 4
)

func (o ReportGenerationContext) String() string {
	switch o {
	case ReportGenerationContextUndefined:
		return "ReportGenerationContextUndefined"
	case ReportGenerationContextScheduledReport:
		return "ReportGenerationContextScheduledReport"
	case ReportGenerationContextInteractiveReport:
		return "ReportGenerationContextInteractiveReport"
	case ReportGenerationContextIncidentReport:
		return "ReportGenerationContextIncidentReport"
	}
	return "Invalid"
}

// ReportFilter type represents FsrmReportFilter RPC enumeration.
//
// The FsrmReportFilter enumeration defines the set of filters that can be used to limit
// the files listed in a report.
type ReportFilter uint16

var (
	// FsrmReportFilter_MinSize:  The report will only show files that meet a minimum size.
	ReportFilterMinSize ReportFilter = 1
	// FsrmReportFilter_MinAgeDays:  The report will only show files that were accessed
	// more than a minimum number of days ago.
	ReportFilterMinAgeDays ReportFilter = 2
	// FsrmReportFilter_MaxAgeDays:  The report will only show files that were accessed
	// prior to a maximum number of days ago.
	ReportFilterMaxAgeDays ReportFilter = 3
	// FsrmReportFilter_MinQuotaUsage:  The report will only show quotas that meet a certain
	// disk space usage level.
	ReportFilterMinQuotaUsage ReportFilter = 4
	// FsrmReportFilter_FileGroups:  The report will only show files from a given set of
	// groups.
	ReportFilterFileGroups ReportFilter = 5
	// FsrmReportFilter_Owners:  The report will only show files that belong to a certain
	// set of owners.
	ReportFilterOwners ReportFilter = 6
	// FsrmReportFilter_NamePattern:  The report will only show files whose name matches
	// the given pattern.
	ReportFilterNamePattern ReportFilter = 7
	// FsrmReportFilter_Property:  The report will show only files whose property matches
	// the given property name.
	ReportFilterProperty ReportFilter = 8
)

func (o ReportFilter) String() string {
	switch o {
	case ReportFilterMinSize:
		return "ReportFilterMinSize"
	case ReportFilterMinAgeDays:
		return "ReportFilterMinAgeDays"
	case ReportFilterMaxAgeDays:
		return "ReportFilterMaxAgeDays"
	case ReportFilterMinQuotaUsage:
		return "ReportFilterMinQuotaUsage"
	case ReportFilterFileGroups:
		return "ReportFilterFileGroups"
	case ReportFilterOwners:
		return "ReportFilterOwners"
	case ReportFilterNamePattern:
		return "ReportFilterNamePattern"
	case ReportFilterProperty:
		return "ReportFilterProperty"
	}
	return "Invalid"
}

// ReportLimit type represents FsrmReportLimit RPC enumeration.
//
// The FsrmReportLimit enumeration defines the set of maxima that can be used to limit
// the files listed in a report.
type ReportLimit uint16

var (
	// FsrmReportLimit_MaxFiles:  The report will list a maximum number of files.
	ReportLimitMaxFiles ReportLimit = 1
	// FsrmReportLimit_MaxFileGroups:  The report will list a maximum number of file groups.
	ReportLimitMaxFileGroups ReportLimit = 2
	// FsrmReportLimit_MaxOwners:  The report will list a maximum number of owners.
	ReportLimitMaxOwners ReportLimit = 3
	// FsrmReportLimit_MaxFilesPerFileGroup:  The report will list a maximum number of
	// files per file group.
	ReportLimitMaxFilesPerFileGroup ReportLimit = 4
	// FsrmReportLimit_MaxFilesPerOwner:  The report will be limited to a maximum number
	// of files per owner.
	ReportLimitMaxFilesPerOwner ReportLimit = 5
	// FsrmReportLimit_MaxFilesPerDuplGroup:  The report will list a maximum number of
	// file entries per duplicated file.
	ReportLimitMaxFilesPerDuplGroup ReportLimit = 6
	// FsrmReportLimit_MaxDuplicateGroups:  The report will list a maximum number of groups
	// for duplicated files (each set of duplicate files is one group).
	ReportLimitMaxDuplicateGroups ReportLimit = 7
	// FsrmReportLimit_MaxQuotas:  The report will list a maximum number of quotas.
	ReportLimitMaxQuotas ReportLimit = 8
	// FsrmReportLimit_MaxFileScreenEvents:  The report will list a maximum number of file
	// screen events.
	ReportLimitMaxFileScreenEvents ReportLimit = 9
	// FsrmReportLimit_MaxPropertyValues:  The report will list a maximum number of property
	// values per property.
	ReportLimitMaxPropertyValues ReportLimit = 10
	// FsrmReportLimit_MaxFilesPerPropertyValue:  The report will list a maximum number
	// of files per property value.
	ReportLimitMaxFilesPerPropertyValue ReportLimit = 11
	// FsrmReportLimit_MaxFolders:  The report will list a maximum number of folders.
	ReportLimitMaxFolders ReportLimit = 12
)

func (o ReportLimit) String() string {
	switch o {
	case ReportLimitMaxFiles:
		return "ReportLimitMaxFiles"
	case ReportLimitMaxFileGroups:
		return "ReportLimitMaxFileGroups"
	case ReportLimitMaxOwners:
		return "ReportLimitMaxOwners"
	case ReportLimitMaxFilesPerFileGroup:
		return "ReportLimitMaxFilesPerFileGroup"
	case ReportLimitMaxFilesPerOwner:
		return "ReportLimitMaxFilesPerOwner"
	case ReportLimitMaxFilesPerDuplGroup:
		return "ReportLimitMaxFilesPerDuplGroup"
	case ReportLimitMaxDuplicateGroups:
		return "ReportLimitMaxDuplicateGroups"
	case ReportLimitMaxQuotas:
		return "ReportLimitMaxQuotas"
	case ReportLimitMaxFileScreenEvents:
		return "ReportLimitMaxFileScreenEvents"
	case ReportLimitMaxPropertyValues:
		return "ReportLimitMaxPropertyValues"
	case ReportLimitMaxFilesPerPropertyValue:
		return "ReportLimitMaxFilesPerPropertyValue"
	case ReportLimitMaxFolders:
		return "ReportLimitMaxFolders"
	}
	return "Invalid"
}

// PropertyDefinitionType type represents FsrmPropertyDefinitionType RPC enumeration.
//
// The FsrmPropertyDefinitionType enumeration defines the set of property definition
// types that can be used to define file classification properties.
type PropertyDefinitionType uint16

var (
	// FsrmPropertyDefinitionType_Unknown:  The property definition type is unknown.
	PropertyDefinitionTypeUnknown PropertyDefinitionType = 0
	// FsrmPropertyDefinitionType_OrderedList:  The property definition defines a list
	// of possible values, one of which can be assigned to the property.
	PropertyDefinitionTypeOrderedList PropertyDefinitionType = 1
	// FsrmPropertyDefinitionType_MultiChoiceList: The property definition defines a list of possible values, one or more of which can be assigned to the property. When a property value of this type is set for a file, the individual choices are separated with the "|" character.
	PropertyDefinitionTypeMultiChoiceList PropertyDefinitionType = 2
	// FsrmPropertyDefinitionType_SingleChoiceList: The property definition defines a list
	// of possible values, one of which can be assigned to the property.<18>
	PropertyDefinitionTypeSingleChoiceList PropertyDefinitionType = 3
	// FsrmPropertyDefinitionType_String:  The property definition type indicates that
	// an arbitrary string value can be assigned to the property.
	PropertyDefinitionTypeString PropertyDefinitionType = 4
	// FsrmPropertyDefinitionType_MultiString:  The property definition indicates that one or more arbitrary string values can be assigned to the property. When a property value of this type is set for a file, the individual strings are separated with the "|" character.
	PropertyDefinitionTypeMultiString PropertyDefinitionType = 5
	// FsrmPropertyDefinitionType_Int:  The property definition indicates that an integer
	// value can be assigned to the property.
	PropertyDefinitionTypeInt PropertyDefinitionType = 6
	// FsrmPropertyDefinitionType_Bool:  The property definition indicates that a Boolean
	// value can be assigned to the property.
	PropertyDefinitionTypeBool PropertyDefinitionType = 7
	// FsrmPropertyDefinitionType_Date:  The property definition indicates that a date
	// value can be assigned to the property.
	PropertyDefinitionTypeDate PropertyDefinitionType = 8
)

func (o PropertyDefinitionType) String() string {
	switch o {
	case PropertyDefinitionTypeUnknown:
		return "PropertyDefinitionTypeUnknown"
	case PropertyDefinitionTypeOrderedList:
		return "PropertyDefinitionTypeOrderedList"
	case PropertyDefinitionTypeMultiChoiceList:
		return "PropertyDefinitionTypeMultiChoiceList"
	case PropertyDefinitionTypeSingleChoiceList:
		return "PropertyDefinitionTypeSingleChoiceList"
	case PropertyDefinitionTypeString:
		return "PropertyDefinitionTypeString"
	case PropertyDefinitionTypeMultiString:
		return "PropertyDefinitionTypeMultiString"
	case PropertyDefinitionTypeInt:
		return "PropertyDefinitionTypeInt"
	case PropertyDefinitionTypeBool:
		return "PropertyDefinitionTypeBool"
	case PropertyDefinitionTypeDate:
		return "PropertyDefinitionTypeDate"
	}
	return "Invalid"
}

// RuleType type represents FsrmRuleType RPC enumeration.
//
// The FsrmRuleType enumeration defines the set of rule types that can be defined for
// automatic file classification.
type RuleType uint16

var (
	// FsrmRuleType_Unknown:  The rule is of an unknown type.
	RuleTypeUnknown RuleType = 0
	// FsrmRuleType_Classification:  The rule defines parameters for how a classification
	// module will operate on a file.
	RuleTypeClassification RuleType = 1
	// FsrmRuleType_Generic:  The rule defines parameters for how modules that are not
	// classification modules will operate on a file.
	RuleTypeGeneric RuleType = 2
)

func (o RuleType) String() string {
	switch o {
	case RuleTypeUnknown:
		return "RuleTypeUnknown"
	case RuleTypeClassification:
		return "RuleTypeClassification"
	case RuleTypeGeneric:
		return "RuleTypeGeneric"
	}
	return "Invalid"
}

// RuleFlags type represents FsrmRuleFlags RPC enumeration.
//
// The FsrmRuleFlags enumeration defines the possible states of the rule objects in
// the File Server Resource Manager Protocol.
type RuleFlags uint16

var (
	// FsrmRuleFlags_Disabled:  If set, the server does not run the rule when classifying
	// a file.
	RuleFlagsDisabled RuleFlags = 256
	// FsrmRuleFlags_Invalid:  If the FsrmRuleFlags_Invalid flag is set, the rule defines
	// an invalid set of parameters and will not be run when classifying a file.
	RuleFlagsInvalid RuleFlags = 4096
)

func (o RuleFlags) String() string {
	switch o {
	case RuleFlagsDisabled:
		return "RuleFlagsDisabled"
	case RuleFlagsInvalid:
		return "RuleFlagsInvalid"
	}
	return "Invalid"
}

// ClassificationLoggingFlags type represents FsrmClassificationLoggingFlags RPC enumeration.
//
// The FsrmClassificationLoggingFlags enumeration defines the different options for
// logging during automatic classification.
type ClassificationLoggingFlags uint16

var (
	// FsrmClassificationLoggingFlags_None: Indicates that no flags are set.
	ClassificationLoggingFlagsNone ClassificationLoggingFlags = 0
	// FsrmClassificationLoggingFlags_ClassificationsInLogFile: If the FsrmClassificationLoggingFlags_ClassificationsInLogFile
	// flag is set, File Server Resource Manager will log how files are classified during
	// automatic classification in a log file.
	ClassificationLoggingFlagsClassificationsInLogFile ClassificationLoggingFlags = 1
	// FsrmClassificationLoggingFlags_ErrorsInLogFile: If the FsrmClassificationLoggingFlags_ErrorsInLogFile
	// flag is set, File Server Resource Manager will log errors that occur during automatic
	// classification in a log file.
	ClassificationLoggingFlagsErrorsInLogFile ClassificationLoggingFlags = 2
	// FsrmClassificationLoggingFlags_ClassificationsInSystemLog: If the FsrmClassificationLoggingFlags_ClassificationsInSystemLog
	// flag is set, File Server Resource Manager will log how files are classified during
	// automatic classification in the System event log.
	ClassificationLoggingFlagsClassificationsInSystemLog ClassificationLoggingFlags = 4
	// FsrmClassificationLoggingFlags_ErrorsInSystemLog: If the FsrmClassificationLoggingFlags_ErrorsInSystemLog
	// flag is set, File Server Resource Manager will log errors that occur during automatic
	// classification in the System event log.
	ClassificationLoggingFlagsErrorsInSystemLog ClassificationLoggingFlags = 8
)

func (o ClassificationLoggingFlags) String() string {
	switch o {
	case ClassificationLoggingFlagsNone:
		return "ClassificationLoggingFlagsNone"
	case ClassificationLoggingFlagsClassificationsInLogFile:
		return "ClassificationLoggingFlagsClassificationsInLogFile"
	case ClassificationLoggingFlagsErrorsInLogFile:
		return "ClassificationLoggingFlagsErrorsInLogFile"
	case ClassificationLoggingFlagsClassificationsInSystemLog:
		return "ClassificationLoggingFlagsClassificationsInSystemLog"
	case ClassificationLoggingFlagsErrorsInSystemLog:
		return "ClassificationLoggingFlagsErrorsInSystemLog"
	}
	return "Invalid"
}

// ExecutionOption type represents FsrmExecutionOption RPC enumeration.
//
// The FsrmExecutionOption enumeration defines the set of execution options that can
// be used to specify when a classification rule will be evaluated.
type ExecutionOption uint16

var (
	// FsrmExecutionOption_Unknown:  The execution option is unknown.
	ExecutionOptionUnknown ExecutionOption = 0
	// FsrmExecutionOption_EvaluateUnset:  The classification rule will be evaluated only
	// if the property it sets is not already set on the file.
	ExecutionOptionEvaluateUnset ExecutionOption = 1
	// FsrmExecutionOption_ReEvaluate_ConsiderExistingValue: The classification rule will
	// always be evaluated and the property value it tries to set will be aggregated with
	// the current value of the property in the file, if any.
	ExecutionOptionReEvaluateConsiderExistingValue ExecutionOption = 2
	// FsrmExecutionOption_ReEvaluate_IgnoreExistingValue: The classification rule will
	// always be evaluated and the property value it tries to set will not be aggregated
	// with the current value of the property in the file, if any.
	ExecutionOptionReEvaluateIgnoreExistingValue ExecutionOption = 3
)

func (o ExecutionOption) String() string {
	switch o {
	case ExecutionOptionUnknown:
		return "ExecutionOptionUnknown"
	case ExecutionOptionEvaluateUnset:
		return "ExecutionOptionEvaluateUnset"
	case ExecutionOptionReEvaluateConsiderExistingValue:
		return "ExecutionOptionReEvaluateConsiderExistingValue"
	case ExecutionOptionReEvaluateIgnoreExistingValue:
		return "ExecutionOptionReEvaluateIgnoreExistingValue"
	}
	return "Invalid"
}

// StorageModuleCaps type represents FsrmStorageModuleCaps RPC enumeration.
//
// The FsrmStorageModuleCaps enumeration defines the capabilities of the storage module.
type StorageModuleCaps uint16

var (
	// FsrmStorageModuleCaps_Unknown:  This enumeration value is not used by FSRM and MUST
	// NOT be referenced. If the server receives this enumeration value, it MUST consider
	// the value invalid and not apply any changes.
	StorageModuleCapsUnknown StorageModuleCaps = 0
	// FsrmStorageModuleCaps_CanGet:  If the FsrmStorageModuleCaps_CanGet flag is set,
	// the storage module is allowed to retrieve classification properties.
	StorageModuleCapsCanGet StorageModuleCaps = 1
	// FsrmStorageModuleCaps_CanSet:  If the FsrmStorageModuleCaps_CanSet flag is set,
	// the storage module is allowed to store classification properties.
	StorageModuleCapsCanSet StorageModuleCaps = 2
	// FsrmStorageModuleCaps_CanHandleDirectories:  If the FsrmStorageModuleCaps_CanHandleDirectories
	// flag is set, the storage module can process folders.
	StorageModuleCapsCanHandleDirectories StorageModuleCaps = 4
	// FsrmStorageModuleCaps_CanHandleFiles:  If the FsrmStorageModuleCaps_CanHandleFiles
	// flag is set, the storage module can process files.
	StorageModuleCapsCanHandleFiles StorageModuleCaps = 8
)

func (o StorageModuleCaps) String() string {
	switch o {
	case StorageModuleCapsUnknown:
		return "StorageModuleCapsUnknown"
	case StorageModuleCapsCanGet:
		return "StorageModuleCapsCanGet"
	case StorageModuleCapsCanSet:
		return "StorageModuleCapsCanSet"
	case StorageModuleCapsCanHandleDirectories:
		return "StorageModuleCapsCanHandleDirectories"
	case StorageModuleCapsCanHandleFiles:
		return "StorageModuleCapsCanHandleFiles"
	}
	return "Invalid"
}

// StorageModuleType type represents FsrmStorageModuleType RPC enumeration.
//
// The FsrmStorageModuleType enumeration defines the possible storage module types.
type StorageModuleType uint16

var (
	// FsrmStorageModuleType_Unknown:  The module type is unknown. Do not use this value.
	StorageModuleTypeUnknown StorageModuleType = 0
	// FsrmStorageModuleType_Cache:  If the FsrmStorageModuleType_Cache flag is set, the
	// classification properties are cached for quick access by storage module.
	StorageModuleTypeCache StorageModuleType = 1
	// FsrmStorageModuleType_InFile:  If the FsrmStorageModuleType_InFile flag is set,
	// the classification properties are cached within the file itself by storage.
	StorageModuleTypeInFile StorageModuleType = 2
	// FsrmStorageModuleType_Database:  If the FsrmStorageModuleType_Database flag is set,
	// the classification properties are cached outside the file (such as using a local
	// database) by storage module.
	StorageModuleTypeDatabase StorageModuleType = 3
	// FsrmStorageModuleType_System:  If the FsrmStorageModuleType_ System flag is set,
	// the classification properties are cached in a system-specific storage.
	StorageModuleTypeSystem StorageModuleType = 100
)

func (o StorageModuleType) String() string {
	switch o {
	case StorageModuleTypeUnknown:
		return "StorageModuleTypeUnknown"
	case StorageModuleTypeCache:
		return "StorageModuleTypeCache"
	case StorageModuleTypeInFile:
		return "StorageModuleTypeInFile"
	case StorageModuleTypeDatabase:
		return "StorageModuleTypeDatabase"
	case StorageModuleTypeSystem:
		return "StorageModuleTypeSystem"
	}
	return "Invalid"
}

// PropertyFlags type represents FsrmPropertyFlags RPC enumeration.
//
// The FsrmPropertyFlags enumeration defines the set of possible states of classification
// properties.
type PropertyFlags uint16

var (
	// FsrmPropertyFlags_Orphaned:  If set, the classification property does not have a
	// corresponding property definition defined in the File Server Resource Manager.
	PropertyFlagsOrphaned PropertyFlags = 1
	// FsrmPropertyFlags_RetrievedFromCache:  If set, the value of the classification property
	// was retrieved from a cache storage module.
	PropertyFlagsRetrievedFromCache PropertyFlags = 2
	// FsrmPropertyFlags_RetrievedFromStorage:  If set, the value of the classification
	// property was retrieved from the file content.
	PropertyFlagsRetrievedFromStorage PropertyFlags = 4
	// FsrmPropertyFlags_SetByClassifier:  If set, the value of the classification property
	// was set by a classification rule.
	PropertyFlagsSetByClassifier PropertyFlags = 8
	// FsrmPropertyFlags_Deleted:  If set, indicates that the classification property has
	// been deleted.
	PropertyFlagsDeleted PropertyFlags = 16
	// FsrmPropertyFlags_Reclassified:  If set, the value was loaded by a storage module
	// but changed by a classification module.
	PropertyFlagsReclassified PropertyFlags = 32
	// FsrmPropertyFlags_AggregationFailed:  If set, the server could not properly aggregate
	// different values of the property supplied by different pipeline modules.
	PropertyFlagsAggregationFailed PropertyFlags = 64
	// FsrmPropertyFlags_Existing:  If set, the property was initially retrieved from a
	// storage module.
	PropertyFlagsExisting PropertyFlags = 128
	// FsrmPropertyFlags_FailedLoadingProperties:  If set, the classification property
	// might only be partially classified because a failure occurred while loading properties
	// from storage.
	PropertyFlagsFailedLoadingProperties PropertyFlags = 256
	// FsrmPropertyFlags_FailedClassifyingProperties:  If set, the classification property
	// might only be partially classified because a failure occurred while classifying properties.
	PropertyFlagsFailedClassifyingProperties PropertyFlags = 512
	// FsrmPropertyFlags_FailedSavingProperties:  If set, the classification property failed
	// to be saved by a storage module.
	PropertyFlagsFailedSavingProperties PropertyFlags = 1024
	// FsrmPropertyFlags_Secure:  If set, the classification property is defined to be
	// a secure property.
	PropertyFlagsSecure PropertyFlags = 2048
	// FsrmPropertyFlags_PolicyDerived:  If set, the classification property was applied
	// as a result of a classification rule.
	PropertyFlagsPolicyDerived PropertyFlags = 4096
	// FsrmPropertyFlags_Inherited:  If set, the classification property value was inherited
	// from the property value of the file's parent folder.
	PropertyFlagsInherited PropertyFlags = 8192
	// FsrmPropertyFlags_Manual:  If set, the classification property value was set manually.
	PropertyFlagsManual PropertyFlags = 16384
	// FsrmPropertyFlags_PropertySourceMask:  This is the bitwise-OR'd combination of FsrmPropertyFlags_RetrievedFromCache,
	// FsrmPropertyFlags_RetrievedFromStorage, and FsrmPropertyFlags_SetByClassifier, which
	// reference to the source of the property.
	PropertyFlagsSourceMask PropertyFlags = 14
)

func (o PropertyFlags) String() string {
	switch o {
	case PropertyFlagsOrphaned:
		return "PropertyFlagsOrphaned"
	case PropertyFlagsRetrievedFromCache:
		return "PropertyFlagsRetrievedFromCache"
	case PropertyFlagsRetrievedFromStorage:
		return "PropertyFlagsRetrievedFromStorage"
	case PropertyFlagsSetByClassifier:
		return "PropertyFlagsSetByClassifier"
	case PropertyFlagsDeleted:
		return "PropertyFlagsDeleted"
	case PropertyFlagsReclassified:
		return "PropertyFlagsReclassified"
	case PropertyFlagsAggregationFailed:
		return "PropertyFlagsAggregationFailed"
	case PropertyFlagsExisting:
		return "PropertyFlagsExisting"
	case PropertyFlagsFailedLoadingProperties:
		return "PropertyFlagsFailedLoadingProperties"
	case PropertyFlagsFailedClassifyingProperties:
		return "PropertyFlagsFailedClassifyingProperties"
	case PropertyFlagsFailedSavingProperties:
		return "PropertyFlagsFailedSavingProperties"
	case PropertyFlagsSecure:
		return "PropertyFlagsSecure"
	case PropertyFlagsPolicyDerived:
		return "PropertyFlagsPolicyDerived"
	case PropertyFlagsInherited:
		return "PropertyFlagsInherited"
	case PropertyFlagsManual:
		return "PropertyFlagsManual"
	case PropertyFlagsSourceMask:
		return "PropertyFlagsSourceMask"
	}
	return "Invalid"
}

// PipelineModuleType type represents FsrmPipelineModuleType RPC enumeration.
//
// The FsrmPipelineModuleType enumeration defines the set of types of modules used in
// the File Server Resource Manager classification pipeline.
type PipelineModuleType uint16

var (
	// FsrmPipelineModuleType_Unknown:  This enumeration value is not used by FSRM and
	// MUST NOT be referenced. If the server receives this enumeration value, it MUST consider
	// the value invalid and not apply any changes.
	PipelineModuleTypeUnknown PipelineModuleType = 0
	// FsrmPipelineModuleType_Storage:  The module is a storage module, which can persist
	// or retrieve property values for files that it processes.
	PipelineModuleTypeStorage PipelineModuleType = 1
	// FsrmPipelineModuleType_Classifier:  The module is a classifier, which can assign
	// property values to files that it processes based on classification rules.
	PipelineModuleTypeClassifier PipelineModuleType = 2
)

func (o PipelineModuleType) String() string {
	switch o {
	case PipelineModuleTypeUnknown:
		return "PipelineModuleTypeUnknown"
	case PipelineModuleTypeStorage:
		return "PipelineModuleTypeStorage"
	case PipelineModuleTypeClassifier:
		return "PipelineModuleTypeClassifier"
	}
	return "Invalid"
}

// GetFilePropertyOptions type represents FsrmGetFilePropertyOptions RPC enumeration.
//
// The FsrmGetFilePropertyOptions enumeration defines how classification properties
// associated with a file are retrieved.
type GetFilePropertyOptions uint16

var (
	// FsrmGetFilePropertyOptions_None:  If the FsrmGetFilePropertyOptions_None flag is
	// set, File Server Resource Manager retrieves classification properties for the given
	// file.
	GetFilePropertyOptionsNone GetFilePropertyOptions = 0
	// FsrmGetFilePropertyOptions_NoRuleEvaluation:  If the FsrmGetFilePropertyOptions_NoRuleEvaluation
	// flag is set, File Server Resource Manager retrieves only classification properties
	// that are not assigned by evaluating the current set of classification rules.
	GetFilePropertyOptionsNoRuleEvaluation GetFilePropertyOptions = 1
	// FsrmGetFilePropertyOptions_Persistent:  If the FsrmGetFilePropertyOptions_Persistent
	// flag is set, File Server Resource Manager retrieves classification properties and
	// saves them.
	GetFilePropertyOptionsPersistent GetFilePropertyOptions = 2
	// FsrmGetFilePropertyOptions_FailOnPersistErrors:  If the FsrmGetFilePropertyOptions_FailOnPersistErrors
	// flag is set, File Server Resource Manager retrieves classification properties and
	// fails the call if there are any errors while saving them.
	GetFilePropertyOptionsFailOnPersistErrors GetFilePropertyOptions = 4
	// FsrmGetFilePropertyOptions_SkipOrphaned:  If the FsrmGetFilePropertyOptions_SkipOrphaned
	// flag is set, File Server Resource Manager only retrieves classification properties
	// for which a Persisted Property Definition exists.
	GetFilePropertyOptionsSkipOrphaned GetFilePropertyOptions = 8
)

func (o GetFilePropertyOptions) String() string {
	switch o {
	case GetFilePropertyOptionsNone:
		return "GetFilePropertyOptionsNone"
	case GetFilePropertyOptionsNoRuleEvaluation:
		return "GetFilePropertyOptionsNoRuleEvaluation"
	case GetFilePropertyOptionsPersistent:
		return "GetFilePropertyOptionsPersistent"
	case GetFilePropertyOptionsFailOnPersistErrors:
		return "GetFilePropertyOptionsFailOnPersistErrors"
	case GetFilePropertyOptionsSkipOrphaned:
		return "GetFilePropertyOptionsSkipOrphaned"
	}
	return "Invalid"
}

// FileManagementType type represents FsrmFileManagementType RPC enumeration.
//
// The FsrmFileManagementType enumeration defines the set of file management job types
// that are available in the File Server Resource Manager.
type FileManagementType uint16

var (
	// FsrmFileManagementType_Unknown:  The file management job type is unknown.
	FileManagementTypeUnknown FileManagementType = 0
	// FsrmFileManagementType_Expiration:  This file management job performs an expiration
	// policy on files meeting a certain criteria.
	FileManagementTypeExpiration FileManagementType = 1
	// FsrmFileManagementType_Custom:  This file management job performs a custom policy
	// on files meeting a certain criteria.
	FileManagementTypeCustom FileManagementType = 2
	// FsrmFileManagementType_Rms:  This file management job performs an Active Directory
	// Rights Management Services policy on files meeting certain criteria.
	FileManagementTypeRMS FileManagementType = 3
)

func (o FileManagementType) String() string {
	switch o {
	case FileManagementTypeUnknown:
		return "FileManagementTypeUnknown"
	case FileManagementTypeExpiration:
		return "FileManagementTypeExpiration"
	case FileManagementTypeCustom:
		return "FileManagementTypeCustom"
	case FileManagementTypeRMS:
		return "FileManagementTypeRMS"
	}
	return "Invalid"
}

// FileManagementLoggingFlags type represents FsrmFileManagementLoggingFlags RPC enumeration.
//
// The FsrmFileManagementLoggingFlags enumeration defines the different options for
// logging when running a file management job.
type FileManagementLoggingFlags uint16

var (
	// FsrmFileManagementLoggingFlags_None: Indicates that no flags are set.
	FileManagementLoggingFlagsNone FileManagementLoggingFlags = 0
	// FsrmFileManagementLoggingFlags_Error:  If the FsrmFileManagementLoggingFlags_Error
	// flag is set, File Server Resource Manager logs errors that occur when running the
	// file management job to the error log.
	FileManagementLoggingFlagsError FileManagementLoggingFlags = 1
	// FsrmFileManagementLoggingFlags_Information:  If the FsrmFileManagementLoggingFlags_Information
	// flag is set, File Server Resource Manager logs information status messages that occur
	// when running the file management job to the information log.
	FileManagementLoggingFlagsInformation FileManagementLoggingFlags = 2
	// FsrmFileManagementLoggingFlags_Audit:  If the FsrmFileManagementLoggingFlags_Audit
	// flag is set, File Server Resource Manager logs information about files that are processed
	// when the server is running the file management job to the server's security audit
	// log.
	FileManagementLoggingFlagsAudit FileManagementLoggingFlags = 4
)

func (o FileManagementLoggingFlags) String() string {
	switch o {
	case FileManagementLoggingFlagsNone:
		return "FileManagementLoggingFlagsNone"
	case FileManagementLoggingFlagsError:
		return "FileManagementLoggingFlagsError"
	case FileManagementLoggingFlagsInformation:
		return "FileManagementLoggingFlagsInformation"
	case FileManagementLoggingFlagsAudit:
		return "FileManagementLoggingFlagsAudit"
	}
	return "Invalid"
}

// PropertyConditionType type represents FsrmPropertyConditionType RPC enumeration.
//
// The FsrmPropertyConditionType enumeration defines the set of comparison operations
// that can be used to determine whether a property value of a file meets a particular
// condition.
type PropertyConditionType uint16

var (
	// FsrmPropertyConditionType_Unknown:  The property condition type is unknown.
	PropertyConditionTypeUnknown PropertyConditionType = 0
	// FsrmPropertyConditionType_Equal:  This property condition is met if the property
	// value is equal to a specified value.
	PropertyConditionTypeEqual PropertyConditionType = 1
	// FsrmPropertyConditionType_NotEqual:  This property condition is met if the property
	// value is not equal to a specified value.
	PropertyConditionTypeNotEqual PropertyConditionType = 2
	// FsrmPropertyConditionType_GreaterThan:  This property condition is met if the property
	// value is greater than a specified value.
	PropertyConditionTypeGreaterThan PropertyConditionType = 3
	// FsrmPropertyConditionType_LessThan:  This property condition is met if the property
	// value is less than a specified value.
	PropertyConditionTypeLessThan PropertyConditionType = 4
	// FsrmPropertyConditionType_Contain:  This property condition is met if the property
	// value is contains a specified value.
	PropertyConditionTypeContain PropertyConditionType = 5
	// FsrmPropertyConditionType_Exist:  This property condition is met if the property
	// value exists.
	PropertyConditionTypeExist PropertyConditionType = 6
	// FsrmPropertyConditionType_NotExist:  This property condition is met if the property
	// value does not exist.
	PropertyConditionTypeNotExist PropertyConditionType = 7
	// FsrmPropertyConditionType_StartWith:  This property condition is met if the property
	// value starts with a specified value.
	PropertyConditionTypeStartWith PropertyConditionType = 8
	// FsrmPropertyConditionType_EndWith:  This property condition is met if the property
	// value ends with a specified value.
	PropertyConditionTypeEndWith PropertyConditionType = 9
	// FsrmPropertyConditionType_ContainedIn:  This property condition is met if the property
	// value is one of a specified value.
	PropertyConditionTypeContainedIn PropertyConditionType = 10
	// FsrmPropertyConditionType_PrefixOf:  This property condition is met if the property
	// value is the prefix of a specified value.
	PropertyConditionTypePrefixOf PropertyConditionType = 11
	// FsrmPropertyConditionType_SuffixOf:  This property condition is met if the property
	// value is the suffix of a specified value.
	PropertyConditionTypeSuffixOf PropertyConditionType = 12
)

func (o PropertyConditionType) String() string {
	switch o {
	case PropertyConditionTypeUnknown:
		return "PropertyConditionTypeUnknown"
	case PropertyConditionTypeEqual:
		return "PropertyConditionTypeEqual"
	case PropertyConditionTypeNotEqual:
		return "PropertyConditionTypeNotEqual"
	case PropertyConditionTypeGreaterThan:
		return "PropertyConditionTypeGreaterThan"
	case PropertyConditionTypeLessThan:
		return "PropertyConditionTypeLessThan"
	case PropertyConditionTypeContain:
		return "PropertyConditionTypeContain"
	case PropertyConditionTypeExist:
		return "PropertyConditionTypeExist"
	case PropertyConditionTypeNotExist:
		return "PropertyConditionTypeNotExist"
	case PropertyConditionTypeStartWith:
		return "PropertyConditionTypeStartWith"
	case PropertyConditionTypeEndWith:
		return "PropertyConditionTypeEndWith"
	case PropertyConditionTypeContainedIn:
		return "PropertyConditionTypeContainedIn"
	case PropertyConditionTypePrefixOf:
		return "PropertyConditionTypePrefixOf"
	case PropertyConditionTypeSuffixOf:
		return "PropertyConditionTypeSuffixOf"
	}
	return "Invalid"
}

// QuotaBase structure represents IFsrmQuotaBase RPC structure.
type QuotaBase dcom.InterfacePointer

func (o *QuotaBase) InterfacePointer() *dcom.InterfacePointer { return (*dcom.InterfacePointer)(o) }

func (o *QuotaBase) xxx_PreparePayload(ctx context.Context) error {
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

func (o *QuotaBase) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *QuotaBase) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *QuotaBase) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// Report structure represents IFsrmReport RPC structure.
type Report dcom.InterfacePointer

func (o *Report) InterfacePointer() *dcom.InterfacePointer { return (*dcom.InterfacePointer)(o) }

func (o *Report) xxx_PreparePayload(ctx context.Context) error {
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

func (o *Report) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *Report) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *Report) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// QuotaObject structure represents IFsrmQuotaObject RPC structure.
type QuotaObject dcom.InterfacePointer

func (o *QuotaObject) InterfacePointer() *dcom.InterfacePointer { return (*dcom.InterfacePointer)(o) }

func (o *QuotaObject) xxx_PreparePayload(ctx context.Context) error {
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

func (o *QuotaObject) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *QuotaObject) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *QuotaObject) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// ReportScheduler structure represents IFsrmReportScheduler RPC structure.
type ReportScheduler dcom.InterfacePointer

func (o *ReportScheduler) InterfacePointer() *dcom.InterfacePointer {
	return (*dcom.InterfacePointer)(o)
}

func (o *ReportScheduler) xxx_PreparePayload(ctx context.Context) error {
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

func (o *ReportScheduler) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *ReportScheduler) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *ReportScheduler) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// PropertyDefinition structure represents IFsrmPropertyDefinition RPC structure.
type PropertyDefinition dcom.InterfacePointer

func (o *PropertyDefinition) InterfacePointer() *dcom.InterfacePointer {
	return (*dcom.InterfacePointer)(o)
}

func (o *PropertyDefinition) xxx_PreparePayload(ctx context.Context) error {
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

func (o *PropertyDefinition) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *PropertyDefinition) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *PropertyDefinition) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// Setting structure represents IFsrmSetting RPC structure.
type Setting dcom.InterfacePointer

func (o *Setting) InterfacePointer() *dcom.InterfacePointer { return (*dcom.InterfacePointer)(o) }

func (o *Setting) xxx_PreparePayload(ctx context.Context) error {
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

func (o *Setting) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *Setting) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *Setting) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// QuotaTemplate structure represents IFsrmQuotaTemplate RPC structure.
type QuotaTemplate dcom.InterfacePointer

func (o *QuotaTemplate) InterfacePointer() *dcom.InterfacePointer { return (*dcom.InterfacePointer)(o) }

func (o *QuotaTemplate) xxx_PreparePayload(ctx context.Context) error {
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

func (o *QuotaTemplate) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *QuotaTemplate) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *QuotaTemplate) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// ActionReport structure represents IFsrmActionReport RPC structure.
type ActionReport dcom.InterfacePointer

func (o *ActionReport) InterfacePointer() *dcom.InterfacePointer { return (*dcom.InterfacePointer)(o) }

func (o *ActionReport) xxx_PreparePayload(ctx context.Context) error {
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

func (o *ActionReport) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *ActionReport) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *ActionReport) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// FileScreenException structure represents IFsrmFileScreenException RPC structure.
type FileScreenException dcom.InterfacePointer

func (o *FileScreenException) InterfacePointer() *dcom.InterfacePointer {
	return (*dcom.InterfacePointer)(o)
}

func (o *FileScreenException) xxx_PreparePayload(ctx context.Context) error {
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

func (o *FileScreenException) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *FileScreenException) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *FileScreenException) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// ClassificationManager structure represents IFsrmClassificationManager RPC structure.
type ClassificationManager dcom.InterfacePointer

func (o *ClassificationManager) InterfacePointer() *dcom.InterfacePointer {
	return (*dcom.InterfacePointer)(o)
}

func (o *ClassificationManager) xxx_PreparePayload(ctx context.Context) error {
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

func (o *ClassificationManager) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *ClassificationManager) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *ClassificationManager) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// ActionEmail2 structure represents IFsrmActionEmail2 RPC structure.
type ActionEmail2 dcom.InterfacePointer

func (o *ActionEmail2) InterfacePointer() *dcom.InterfacePointer { return (*dcom.InterfacePointer)(o) }

func (o *ActionEmail2) xxx_PreparePayload(ctx context.Context) error {
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

func (o *ActionEmail2) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *ActionEmail2) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *ActionEmail2) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// MutableCollection structure represents IFsrmMutableCollection RPC structure.
type MutableCollection dcom.InterfacePointer

func (o *MutableCollection) InterfacePointer() *dcom.InterfacePointer {
	return (*dcom.InterfacePointer)(o)
}

func (o *MutableCollection) xxx_PreparePayload(ctx context.Context) error {
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

func (o *MutableCollection) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *MutableCollection) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *MutableCollection) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// ActionEmail structure represents IFsrmActionEmail RPC structure.
type ActionEmail dcom.InterfacePointer

func (o *ActionEmail) InterfacePointer() *dcom.InterfacePointer { return (*dcom.InterfacePointer)(o) }

func (o *ActionEmail) xxx_PreparePayload(ctx context.Context) error {
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

func (o *ActionEmail) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *ActionEmail) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *ActionEmail) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// CommittableCollection structure represents IFsrmCommittableCollection RPC structure.
type CommittableCollection dcom.InterfacePointer

func (o *CommittableCollection) InterfacePointer() *dcom.InterfacePointer {
	return (*dcom.InterfacePointer)(o)
}

func (o *CommittableCollection) xxx_PreparePayload(ctx context.Context) error {
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

func (o *CommittableCollection) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *CommittableCollection) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *CommittableCollection) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// AutoApplyQuota structure represents IFsrmAutoApplyQuota RPC structure.
type AutoApplyQuota dcom.InterfacePointer

func (o *AutoApplyQuota) InterfacePointer() *dcom.InterfacePointer {
	return (*dcom.InterfacePointer)(o)
}

func (o *AutoApplyQuota) xxx_PreparePayload(ctx context.Context) error {
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

func (o *AutoApplyQuota) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *AutoApplyQuota) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *AutoApplyQuota) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// QuotaTemplateManager structure represents IFsrmQuotaTemplateManager RPC structure.
type QuotaTemplateManager dcom.InterfacePointer

func (o *QuotaTemplateManager) InterfacePointer() *dcom.InterfacePointer {
	return (*dcom.InterfacePointer)(o)
}

func (o *QuotaTemplateManager) xxx_PreparePayload(ctx context.Context) error {
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

func (o *QuotaTemplateManager) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *QuotaTemplateManager) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *QuotaTemplateManager) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// FileScreenTemplateManager structure represents IFsrmFileScreenTemplateManager RPC structure.
type FileScreenTemplateManager dcom.InterfacePointer

func (o *FileScreenTemplateManager) InterfacePointer() *dcom.InterfacePointer {
	return (*dcom.InterfacePointer)(o)
}

func (o *FileScreenTemplateManager) xxx_PreparePayload(ctx context.Context) error {
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

func (o *FileScreenTemplateManager) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *FileScreenTemplateManager) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *FileScreenTemplateManager) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// PropertyDefinition2 structure represents IFsrmPropertyDefinition2 RPC structure.
type PropertyDefinition2 dcom.InterfacePointer

func (o *PropertyDefinition2) InterfacePointer() *dcom.InterfacePointer {
	return (*dcom.InterfacePointer)(o)
}

func (o *PropertyDefinition2) xxx_PreparePayload(ctx context.Context) error {
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

func (o *PropertyDefinition2) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *PropertyDefinition2) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *PropertyDefinition2) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// ReportJob structure represents IFsrmReportJob RPC structure.
type ReportJob dcom.InterfacePointer

func (o *ReportJob) InterfacePointer() *dcom.InterfacePointer { return (*dcom.InterfacePointer)(o) }

func (o *ReportJob) xxx_PreparePayload(ctx context.Context) error {
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

func (o *ReportJob) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *ReportJob) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *ReportJob) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// StorageModuleDefinition structure represents IFsrmStorageModuleDefinition RPC structure.
type StorageModuleDefinition dcom.InterfacePointer

func (o *StorageModuleDefinition) InterfacePointer() *dcom.InterfacePointer {
	return (*dcom.InterfacePointer)(o)
}

func (o *StorageModuleDefinition) xxx_PreparePayload(ctx context.Context) error {
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

func (o *StorageModuleDefinition) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *StorageModuleDefinition) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *StorageModuleDefinition) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// PropertyCondition structure represents IFsrmPropertyCondition RPC structure.
type PropertyCondition dcom.InterfacePointer

func (o *PropertyCondition) InterfacePointer() *dcom.InterfacePointer {
	return (*dcom.InterfacePointer)(o)
}

func (o *PropertyCondition) xxx_PreparePayload(ctx context.Context) error {
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

func (o *PropertyCondition) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *PropertyCondition) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *PropertyCondition) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// Collection structure represents IFsrmCollection RPC structure.
type Collection dcom.InterfacePointer

func (o *Collection) InterfacePointer() *dcom.InterfacePointer { return (*dcom.InterfacePointer)(o) }

func (o *Collection) xxx_PreparePayload(ctx context.Context) error {
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

func (o *Collection) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *Collection) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *Collection) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// PropertyDefinitionValue structure represents IFsrmPropertyDefinitionValue RPC structure.
type PropertyDefinitionValue dcom.InterfacePointer

func (o *PropertyDefinitionValue) InterfacePointer() *dcom.InterfacePointer {
	return (*dcom.InterfacePointer)(o)
}

func (o *PropertyDefinitionValue) xxx_PreparePayload(ctx context.Context) error {
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

func (o *PropertyDefinitionValue) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *PropertyDefinitionValue) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *PropertyDefinitionValue) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// ReportManager structure represents IFsrmReportManager RPC structure.
type ReportManager dcom.InterfacePointer

func (o *ReportManager) InterfacePointer() *dcom.InterfacePointer { return (*dcom.InterfacePointer)(o) }

func (o *ReportManager) xxx_PreparePayload(ctx context.Context) error {
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

func (o *ReportManager) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *ReportManager) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *ReportManager) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// ActionEventLog structure represents IFsrmActionEventLog RPC structure.
type ActionEventLog dcom.InterfacePointer

func (o *ActionEventLog) InterfacePointer() *dcom.InterfacePointer {
	return (*dcom.InterfacePointer)(o)
}

func (o *ActionEventLog) xxx_PreparePayload(ctx context.Context) error {
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

func (o *ActionEventLog) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *ActionEventLog) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *ActionEventLog) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// FileGroupManager structure represents IFsrmFileGroupManager RPC structure.
type FileGroupManager dcom.InterfacePointer

func (o *FileGroupManager) InterfacePointer() *dcom.InterfacePointer {
	return (*dcom.InterfacePointer)(o)
}

func (o *FileGroupManager) xxx_PreparePayload(ctx context.Context) error {
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

func (o *FileGroupManager) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *FileGroupManager) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *FileGroupManager) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// QuotaTemplateImported structure represents IFsrmQuotaTemplateImported RPC structure.
type QuotaTemplateImported dcom.InterfacePointer

func (o *QuotaTemplateImported) InterfacePointer() *dcom.InterfacePointer {
	return (*dcom.InterfacePointer)(o)
}

func (o *QuotaTemplateImported) xxx_PreparePayload(ctx context.Context) error {
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

func (o *QuotaTemplateImported) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *QuotaTemplateImported) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *QuotaTemplateImported) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// ClassifierModuleDefinition structure represents IFsrmClassifierModuleDefinition RPC structure.
type ClassifierModuleDefinition dcom.InterfacePointer

func (o *ClassifierModuleDefinition) InterfacePointer() *dcom.InterfacePointer {
	return (*dcom.InterfacePointer)(o)
}

func (o *ClassifierModuleDefinition) xxx_PreparePayload(ctx context.Context) error {
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

func (o *ClassifierModuleDefinition) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *ClassifierModuleDefinition) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *ClassifierModuleDefinition) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// FileScreenManager structure represents IFsrmFileScreenManager RPC structure.
type FileScreenManager dcom.InterfacePointer

func (o *FileScreenManager) InterfacePointer() *dcom.InterfacePointer {
	return (*dcom.InterfacePointer)(o)
}

func (o *FileScreenManager) xxx_PreparePayload(ctx context.Context) error {
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

func (o *FileScreenManager) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *FileScreenManager) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *FileScreenManager) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// FileGroupImported structure represents IFsrmFileGroupImported RPC structure.
type FileGroupImported dcom.InterfacePointer

func (o *FileGroupImported) InterfacePointer() *dcom.InterfacePointer {
	return (*dcom.InterfacePointer)(o)
}

func (o *FileGroupImported) xxx_PreparePayload(ctx context.Context) error {
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

func (o *FileGroupImported) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *FileGroupImported) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *FileGroupImported) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// FileScreenBase structure represents IFsrmFileScreenBase RPC structure.
type FileScreenBase dcom.InterfacePointer

func (o *FileScreenBase) InterfacePointer() *dcom.InterfacePointer {
	return (*dcom.InterfacePointer)(o)
}

func (o *FileScreenBase) xxx_PreparePayload(ctx context.Context) error {
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

func (o *FileScreenBase) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *FileScreenBase) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *FileScreenBase) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// Object structure represents IFsrmObject RPC structure.
type Object dcom.InterfacePointer

func (o *Object) InterfacePointer() *dcom.InterfacePointer { return (*dcom.InterfacePointer)(o) }

func (o *Object) xxx_PreparePayload(ctx context.Context) error {
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

func (o *Object) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *Object) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *Object) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// FileManagementJobManager structure represents IFsrmFileManagementJobManager RPC structure.
type FileManagementJobManager dcom.InterfacePointer

func (o *FileManagementJobManager) InterfacePointer() *dcom.InterfacePointer {
	return (*dcom.InterfacePointer)(o)
}

func (o *FileManagementJobManager) xxx_PreparePayload(ctx context.Context) error {
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

func (o *FileManagementJobManager) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *FileManagementJobManager) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *FileManagementJobManager) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// Property structure represents IFsrmProperty RPC structure.
type Property dcom.InterfacePointer

func (o *Property) InterfacePointer() *dcom.InterfacePointer { return (*dcom.InterfacePointer)(o) }

func (o *Property) xxx_PreparePayload(ctx context.Context) error {
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

func (o *Property) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *Property) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *Property) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// PipelineModuleDefinition structure represents IFsrmPipelineModuleDefinition RPC structure.
type PipelineModuleDefinition dcom.InterfacePointer

func (o *PipelineModuleDefinition) InterfacePointer() *dcom.InterfacePointer {
	return (*dcom.InterfacePointer)(o)
}

func (o *PipelineModuleDefinition) xxx_PreparePayload(ctx context.Context) error {
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

func (o *PipelineModuleDefinition) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *PipelineModuleDefinition) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *PipelineModuleDefinition) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// FileScreenTemplate structure represents IFsrmFileScreenTemplate RPC structure.
type FileScreenTemplate dcom.InterfacePointer

func (o *FileScreenTemplate) InterfacePointer() *dcom.InterfacePointer {
	return (*dcom.InterfacePointer)(o)
}

func (o *FileScreenTemplate) xxx_PreparePayload(ctx context.Context) error {
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

func (o *FileScreenTemplate) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *FileScreenTemplate) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *FileScreenTemplate) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// Quota structure represents IFsrmQuota RPC structure.
type Quota dcom.InterfacePointer

func (o *Quota) InterfacePointer() *dcom.InterfacePointer { return (*dcom.InterfacePointer)(o) }

func (o *Quota) xxx_PreparePayload(ctx context.Context) error {
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

func (o *Quota) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *Quota) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *Quota) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// FileGroup structure represents IFsrmFileGroup RPC structure.
type FileGroup dcom.InterfacePointer

func (o *FileGroup) InterfacePointer() *dcom.InterfacePointer { return (*dcom.InterfacePointer)(o) }

func (o *FileGroup) xxx_PreparePayload(ctx context.Context) error {
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

func (o *FileGroup) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *FileGroup) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *FileGroup) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// QuotaManager structure represents IFsrmQuotaManager RPC structure.
type QuotaManager dcom.InterfacePointer

func (o *QuotaManager) InterfacePointer() *dcom.InterfacePointer { return (*dcom.InterfacePointer)(o) }

func (o *QuotaManager) xxx_PreparePayload(ctx context.Context) error {
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

func (o *QuotaManager) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *QuotaManager) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *QuotaManager) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// FileManagementJob structure represents IFsrmFileManagementJob RPC structure.
type FileManagementJob dcom.InterfacePointer

func (o *FileManagementJob) InterfacePointer() *dcom.InterfacePointer {
	return (*dcom.InterfacePointer)(o)
}

func (o *FileManagementJob) xxx_PreparePayload(ctx context.Context) error {
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

func (o *FileManagementJob) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *FileManagementJob) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *FileManagementJob) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// Action structure represents IFsrmAction RPC structure.
type Action dcom.InterfacePointer

func (o *Action) InterfacePointer() *dcom.InterfacePointer { return (*dcom.InterfacePointer)(o) }

func (o *Action) xxx_PreparePayload(ctx context.Context) error {
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

func (o *Action) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *Action) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *Action) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// PathMapper structure represents IFsrmPathMapper RPC structure.
type PathMapper dcom.InterfacePointer

func (o *PathMapper) InterfacePointer() *dcom.InterfacePointer { return (*dcom.InterfacePointer)(o) }

func (o *PathMapper) xxx_PreparePayload(ctx context.Context) error {
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

func (o *PathMapper) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *PathMapper) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *PathMapper) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// QuotaManagerEx structure represents IFsrmQuotaManagerEx RPC structure.
type QuotaManagerEx dcom.InterfacePointer

func (o *QuotaManagerEx) InterfacePointer() *dcom.InterfacePointer {
	return (*dcom.InterfacePointer)(o)
}

func (o *QuotaManagerEx) xxx_PreparePayload(ctx context.Context) error {
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

func (o *QuotaManagerEx) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *QuotaManagerEx) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *QuotaManagerEx) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// FileScreen structure represents IFsrmFileScreen RPC structure.
type FileScreen dcom.InterfacePointer

func (o *FileScreen) InterfacePointer() *dcom.InterfacePointer { return (*dcom.InterfacePointer)(o) }

func (o *FileScreen) xxx_PreparePayload(ctx context.Context) error {
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

func (o *FileScreen) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *FileScreen) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *FileScreen) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// ActionCommand structure represents IFsrmActionCommand RPC structure.
type ActionCommand dcom.InterfacePointer

func (o *ActionCommand) InterfacePointer() *dcom.InterfacePointer { return (*dcom.InterfacePointer)(o) }

func (o *ActionCommand) xxx_PreparePayload(ctx context.Context) error {
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

func (o *ActionCommand) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *ActionCommand) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *ActionCommand) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// DerivedObjectsResult structure represents IFsrmDerivedObjectsResult RPC structure.
type DerivedObjectsResult dcom.InterfacePointer

func (o *DerivedObjectsResult) InterfacePointer() *dcom.InterfacePointer {
	return (*dcom.InterfacePointer)(o)
}

func (o *DerivedObjectsResult) xxx_PreparePayload(ctx context.Context) error {
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

func (o *DerivedObjectsResult) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *DerivedObjectsResult) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *DerivedObjectsResult) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// FileScreenTemplateImported structure represents IFsrmFileScreenTemplateImported RPC structure.
type FileScreenTemplateImported dcom.InterfacePointer

func (o *FileScreenTemplateImported) InterfacePointer() *dcom.InterfacePointer {
	return (*dcom.InterfacePointer)(o)
}

func (o *FileScreenTemplateImported) xxx_PreparePayload(ctx context.Context) error {
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

func (o *FileScreenTemplateImported) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *FileScreenTemplateImported) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *FileScreenTemplateImported) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// Rule structure represents IFsrmRule RPC structure.
type Rule dcom.InterfacePointer

func (o *Rule) InterfacePointer() *dcom.InterfacePointer { return (*dcom.InterfacePointer)(o) }

func (o *Rule) xxx_PreparePayload(ctx context.Context) error {
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

func (o *Rule) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *Rule) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *Rule) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// ClassificationRule structure represents IFsrmClassificationRule RPC structure.
type ClassificationRule dcom.InterfacePointer

func (o *ClassificationRule) InterfacePointer() *dcom.InterfacePointer {
	return (*dcom.InterfacePointer)(o)
}

func (o *ClassificationRule) xxx_PreparePayload(ctx context.Context) error {
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

func (o *ClassificationRule) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *ClassificationRule) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *ClassificationRule) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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
