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
	dcom_client "github.com/oiweiwei/go-msrpc/msrpc/dcom/client"
	ifsrmaction "github.com/oiweiwei/go-msrpc/msrpc/dcom/fsrm/ifsrmaction/v0"
	ifsrmactioncommand "github.com/oiweiwei/go-msrpc/msrpc/dcom/fsrm/ifsrmactioncommand/v0"
	ifsrmactionemail "github.com/oiweiwei/go-msrpc/msrpc/dcom/fsrm/ifsrmactionemail/v0"
	ifsrmactionemail2 "github.com/oiweiwei/go-msrpc/msrpc/dcom/fsrm/ifsrmactionemail2/v0"
	ifsrmactioneventlog "github.com/oiweiwei/go-msrpc/msrpc/dcom/fsrm/ifsrmactioneventlog/v0"
	ifsrmactionreport "github.com/oiweiwei/go-msrpc/msrpc/dcom/fsrm/ifsrmactionreport/v0"
	ifsrmautoapplyquota "github.com/oiweiwei/go-msrpc/msrpc/dcom/fsrm/ifsrmautoapplyquota/v0"
	ifsrmclassificationmanager "github.com/oiweiwei/go-msrpc/msrpc/dcom/fsrm/ifsrmclassificationmanager/v0"
	ifsrmclassificationrule "github.com/oiweiwei/go-msrpc/msrpc/dcom/fsrm/ifsrmclassificationrule/v0"
	ifsrmclassifiermoduledefinition "github.com/oiweiwei/go-msrpc/msrpc/dcom/fsrm/ifsrmclassifiermoduledefinition/v0"
	ifsrmcollection "github.com/oiweiwei/go-msrpc/msrpc/dcom/fsrm/ifsrmcollection/v0"
	ifsrmcommittablecollection "github.com/oiweiwei/go-msrpc/msrpc/dcom/fsrm/ifsrmcommittablecollection/v0"
	ifsrmderivedobjectsresult "github.com/oiweiwei/go-msrpc/msrpc/dcom/fsrm/ifsrmderivedobjectsresult/v0"
	ifsrmfilegroup "github.com/oiweiwei/go-msrpc/msrpc/dcom/fsrm/ifsrmfilegroup/v0"
	ifsrmfilegroupimported "github.com/oiweiwei/go-msrpc/msrpc/dcom/fsrm/ifsrmfilegroupimported/v0"
	ifsrmfilegroupmanager "github.com/oiweiwei/go-msrpc/msrpc/dcom/fsrm/ifsrmfilegroupmanager/v0"
	ifsrmfilemanagementjob "github.com/oiweiwei/go-msrpc/msrpc/dcom/fsrm/ifsrmfilemanagementjob/v1"
	ifsrmfilemanagementjobmanager "github.com/oiweiwei/go-msrpc/msrpc/dcom/fsrm/ifsrmfilemanagementjobmanager/v1"
	ifsrmfilescreen "github.com/oiweiwei/go-msrpc/msrpc/dcom/fsrm/ifsrmfilescreen/v0"
	ifsrmfilescreenbase "github.com/oiweiwei/go-msrpc/msrpc/dcom/fsrm/ifsrmfilescreenbase/v0"
	ifsrmfilescreenexception "github.com/oiweiwei/go-msrpc/msrpc/dcom/fsrm/ifsrmfilescreenexception/v0"
	ifsrmfilescreenmanager "github.com/oiweiwei/go-msrpc/msrpc/dcom/fsrm/ifsrmfilescreenmanager/v0"
	ifsrmfilescreentemplate "github.com/oiweiwei/go-msrpc/msrpc/dcom/fsrm/ifsrmfilescreentemplate/v0"
	ifsrmfilescreentemplateimported "github.com/oiweiwei/go-msrpc/msrpc/dcom/fsrm/ifsrmfilescreentemplateimported/v0"
	ifsrmfilescreentemplatemanager "github.com/oiweiwei/go-msrpc/msrpc/dcom/fsrm/ifsrmfilescreentemplatemanager/v0"
	ifsrmmutablecollection "github.com/oiweiwei/go-msrpc/msrpc/dcom/fsrm/ifsrmmutablecollection/v0"
	ifsrmobject "github.com/oiweiwei/go-msrpc/msrpc/dcom/fsrm/ifsrmobject/v0"
	ifsrmpathmapper "github.com/oiweiwei/go-msrpc/msrpc/dcom/fsrm/ifsrmpathmapper/v0"
	ifsrmpipelinemoduledefinition "github.com/oiweiwei/go-msrpc/msrpc/dcom/fsrm/ifsrmpipelinemoduledefinition/v0"
	ifsrmproperty "github.com/oiweiwei/go-msrpc/msrpc/dcom/fsrm/ifsrmproperty/v0"
	ifsrmpropertycondition "github.com/oiweiwei/go-msrpc/msrpc/dcom/fsrm/ifsrmpropertycondition/v1"
	ifsrmpropertydefinition "github.com/oiweiwei/go-msrpc/msrpc/dcom/fsrm/ifsrmpropertydefinition/v0"
	ifsrmpropertydefinition2 "github.com/oiweiwei/go-msrpc/msrpc/dcom/fsrm/ifsrmpropertydefinition2/v0"
	ifsrmpropertydefinitionvalue "github.com/oiweiwei/go-msrpc/msrpc/dcom/fsrm/ifsrmpropertydefinitionvalue/v0"
	ifsrmquota "github.com/oiweiwei/go-msrpc/msrpc/dcom/fsrm/ifsrmquota/v0"
	ifsrmquotabase "github.com/oiweiwei/go-msrpc/msrpc/dcom/fsrm/ifsrmquotabase/v0"
	ifsrmquotamanager "github.com/oiweiwei/go-msrpc/msrpc/dcom/fsrm/ifsrmquotamanager/v0"
	ifsrmquotamanagerex "github.com/oiweiwei/go-msrpc/msrpc/dcom/fsrm/ifsrmquotamanagerex/v0"
	ifsrmquotaobject "github.com/oiweiwei/go-msrpc/msrpc/dcom/fsrm/ifsrmquotaobject/v0"
	ifsrmquotatemplate "github.com/oiweiwei/go-msrpc/msrpc/dcom/fsrm/ifsrmquotatemplate/v0"
	ifsrmquotatemplateimported "github.com/oiweiwei/go-msrpc/msrpc/dcom/fsrm/ifsrmquotatemplateimported/v0"
	ifsrmquotatemplatemanager "github.com/oiweiwei/go-msrpc/msrpc/dcom/fsrm/ifsrmquotatemplatemanager/v0"
	ifsrmreport "github.com/oiweiwei/go-msrpc/msrpc/dcom/fsrm/ifsrmreport/v1"
	ifsrmreportjob "github.com/oiweiwei/go-msrpc/msrpc/dcom/fsrm/ifsrmreportjob/v1"
	ifsrmreportmanager "github.com/oiweiwei/go-msrpc/msrpc/dcom/fsrm/ifsrmreportmanager/v1"
	ifsrmreportscheduler "github.com/oiweiwei/go-msrpc/msrpc/dcom/fsrm/ifsrmreportscheduler/v1"
	ifsrmrule "github.com/oiweiwei/go-msrpc/msrpc/dcom/fsrm/ifsrmrule/v0"
	ifsrmsetting "github.com/oiweiwei/go-msrpc/msrpc/dcom/fsrm/ifsrmsetting/v0"
	ifsrmstoragemoduledefinition "github.com/oiweiwei/go-msrpc/msrpc/dcom/fsrm/ifsrmstoragemoduledefinition/v0"
	iremunknown "github.com/oiweiwei/go-msrpc/msrpc/dcom/iremunknown/v0"
	iremunknown2 "github.com/oiweiwei/go-msrpc/msrpc/dcom/iremunknown2/v0"
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
	_ = dcom_client.GoPackage
	_ = dcom.GoPackage
	_ = iremunknown.GoPackage
	_ = iremunknown2.GoPackage
	_ = ifsrmobject.GoPackage
	_ = ifsrmcollection.GoPackage
	_ = ifsrmmutablecollection.GoPackage
	_ = ifsrmcommittablecollection.GoPackage
	_ = ifsrmaction.GoPackage
	_ = ifsrmactionemail.GoPackage
	_ = ifsrmactionemail2.GoPackage
	_ = ifsrmactionreport.GoPackage
	_ = ifsrmactioneventlog.GoPackage
	_ = ifsrmactioncommand.GoPackage
	_ = ifsrmsetting.GoPackage
	_ = ifsrmpathmapper.GoPackage
	_ = ifsrmderivedobjectsresult.GoPackage
	_ = ifsrmpropertydefinition.GoPackage
	_ = ifsrmpropertydefinition2.GoPackage
	_ = ifsrmpropertydefinitionvalue.GoPackage
	_ = ifsrmproperty.GoPackage
	_ = ifsrmrule.GoPackage
	_ = ifsrmclassificationrule.GoPackage
	_ = ifsrmpipelinemoduledefinition.GoPackage
	_ = ifsrmclassifiermoduledefinition.GoPackage
	_ = ifsrmclassificationmanager.GoPackage
	_ = ifsrmstoragemoduledefinition.GoPackage
	_ = ifsrmquotabase.GoPackage
	_ = ifsrmquotaobject.GoPackage
	_ = ifsrmquota.GoPackage
	_ = ifsrmautoapplyquota.GoPackage
	_ = ifsrmquotamanager.GoPackage
	_ = ifsrmquotamanagerex.GoPackage
	_ = ifsrmquotatemplate.GoPackage
	_ = ifsrmquotatemplateimported.GoPackage
	_ = ifsrmquotatemplatemanager.GoPackage
	_ = ifsrmreportmanager.GoPackage
	_ = ifsrmreportjob.GoPackage
	_ = ifsrmreport.GoPackage
	_ = ifsrmreportscheduler.GoPackage
	_ = ifsrmfilemanagementjobmanager.GoPackage
	_ = ifsrmfilemanagementjob.GoPackage
	_ = ifsrmpropertycondition.GoPackage
	_ = ifsrmfilegroup.GoPackage
	_ = ifsrmfilegroupimported.GoPackage
	_ = ifsrmfilegroupmanager.GoPackage
	_ = ifsrmfilescreenbase.GoPackage
	_ = ifsrmfilescreen.GoPackage
	_ = ifsrmfilescreenexception.GoPackage
	_ = ifsrmfilescreenmanager.GoPackage
	_ = ifsrmfilescreentemplate.GoPackage
	_ = ifsrmfilescreentemplateimported.GoPackage
	_ = ifsrmfilescreentemplatemanager.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/fsrm"
)

// dcom/fsrm client set.
type Client interface {

	// DCOM common interfaces
	RemoteUnknown() iremunknown.RemoteUnknownClient
	RemoteUnknown2() iremunknown2.RemoteUnknown2Client

	// Package specific interfaces
	Object() ifsrmobject.ObjectClient
	Collection() ifsrmcollection.CollectionClient
	MutableCollection() ifsrmmutablecollection.MutableCollectionClient
	CommittableCollection() ifsrmcommittablecollection.CommittableCollectionClient
	Action() ifsrmaction.ActionClient
	ActionEmail() ifsrmactionemail.ActionEmailClient
	ActionEmail2() ifsrmactionemail2.ActionEmail2Client
	ActionReport() ifsrmactionreport.ActionReportClient
	ActionEventLog() ifsrmactioneventlog.ActionEventLogClient
	ActionCommand() ifsrmactioncommand.ActionCommandClient
	Setting() ifsrmsetting.SettingClient
	PathMapper() ifsrmpathmapper.PathMapperClient
	DerivedObjectsResult() ifsrmderivedobjectsresult.DerivedObjectsResultClient
	PropertyDefinition() ifsrmpropertydefinition.PropertyDefinitionClient
	PropertyDefinition2() ifsrmpropertydefinition2.PropertyDefinition2Client
	PropertyDefinitionValue() ifsrmpropertydefinitionvalue.PropertyDefinitionValueClient
	Property() ifsrmproperty.PropertyClient
	Rule() ifsrmrule.RuleClient
	ClassificationRule() ifsrmclassificationrule.ClassificationRuleClient
	PipelineModuleDefinition() ifsrmpipelinemoduledefinition.PipelineModuleDefinitionClient
	ClassifierModuleDefinition() ifsrmclassifiermoduledefinition.ClassifierModuleDefinitionClient
	ClassificationManager() ifsrmclassificationmanager.ClassificationManagerClient
	StorageModuleDefinition() ifsrmstoragemoduledefinition.StorageModuleDefinitionClient
	QuotaBase() ifsrmquotabase.QuotaBaseClient
	QuotaObject() ifsrmquotaobject.QuotaObjectClient
	Quota() ifsrmquota.QuotaClient
	AutoApplyQuota() ifsrmautoapplyquota.AutoApplyQuotaClient
	QuotaManager() ifsrmquotamanager.QuotaManagerClient
	QuotaManagerEx() ifsrmquotamanagerex.QuotaManagerExClient
	QuotaTemplate() ifsrmquotatemplate.QuotaTemplateClient
	QuotaTemplateImported() ifsrmquotatemplateimported.QuotaTemplateImportedClient
	QuotaTemplateManager() ifsrmquotatemplatemanager.QuotaTemplateManagerClient
	ReportManager() ifsrmreportmanager.ReportManagerClient
	ReportJob() ifsrmreportjob.ReportJobClient
	Report() ifsrmreport.ReportClient
	ReportScheduler() ifsrmreportscheduler.ReportSchedulerClient
	FileManagementJobManager() ifsrmfilemanagementjobmanager.FileManagementJobManagerClient
	FileManagementJob() ifsrmfilemanagementjob.FileManagementJobClient
	PropertyCondition() ifsrmpropertycondition.PropertyConditionClient
	FileGroup() ifsrmfilegroup.FileGroupClient
	FileGroupImported() ifsrmfilegroupimported.FileGroupImportedClient
	FileGroupManager() ifsrmfilegroupmanager.FileGroupManagerClient
	FileScreenBase() ifsrmfilescreenbase.FileScreenBaseClient
	FileScreen() ifsrmfilescreen.FileScreenClient
	FileScreenException() ifsrmfilescreenexception.FileScreenExceptionClient
	FileScreenManager() ifsrmfilescreenmanager.FileScreenManagerClient
	FileScreenTemplate() ifsrmfilescreentemplate.FileScreenTemplateClient
	FileScreenTemplateImported() ifsrmfilescreentemplateimported.FileScreenTemplateImportedClient
	FileScreenTemplateManager() ifsrmfilescreentemplatemanager.FileScreenTemplateManagerClient
	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) Client
}
type xxx_DefaultClient struct {
	cc dcerpc.Conn

	dcomClient dcom_client.Client

	object                     ifsrmobject.ObjectClient
	collection                 ifsrmcollection.CollectionClient
	mutableCollection          ifsrmmutablecollection.MutableCollectionClient
	committableCollection      ifsrmcommittablecollection.CommittableCollectionClient
	action                     ifsrmaction.ActionClient
	actionEmail                ifsrmactionemail.ActionEmailClient
	actionEmail2               ifsrmactionemail2.ActionEmail2Client
	actionReport               ifsrmactionreport.ActionReportClient
	actionEventLog             ifsrmactioneventlog.ActionEventLogClient
	actionCommand              ifsrmactioncommand.ActionCommandClient
	setting                    ifsrmsetting.SettingClient
	pathMapper                 ifsrmpathmapper.PathMapperClient
	derivedObjectsResult       ifsrmderivedobjectsresult.DerivedObjectsResultClient
	propertyDefinition         ifsrmpropertydefinition.PropertyDefinitionClient
	propertyDefinition2        ifsrmpropertydefinition2.PropertyDefinition2Client
	propertyDefinitionValue    ifsrmpropertydefinitionvalue.PropertyDefinitionValueClient
	property                   ifsrmproperty.PropertyClient
	rule                       ifsrmrule.RuleClient
	classificationRule         ifsrmclassificationrule.ClassificationRuleClient
	pipelineModuleDefinition   ifsrmpipelinemoduledefinition.PipelineModuleDefinitionClient
	classifierModuleDefinition ifsrmclassifiermoduledefinition.ClassifierModuleDefinitionClient
	classificationManager      ifsrmclassificationmanager.ClassificationManagerClient
	storageModuleDefinition    ifsrmstoragemoduledefinition.StorageModuleDefinitionClient
	quotaBase                  ifsrmquotabase.QuotaBaseClient
	quotaObject                ifsrmquotaobject.QuotaObjectClient
	quota                      ifsrmquota.QuotaClient
	autoApplyQuota             ifsrmautoapplyquota.AutoApplyQuotaClient
	quotaManager               ifsrmquotamanager.QuotaManagerClient
	quotaManagerEx             ifsrmquotamanagerex.QuotaManagerExClient
	quotaTemplate              ifsrmquotatemplate.QuotaTemplateClient
	quotaTemplateImported      ifsrmquotatemplateimported.QuotaTemplateImportedClient
	quotaTemplateManager       ifsrmquotatemplatemanager.QuotaTemplateManagerClient
	reportManager              ifsrmreportmanager.ReportManagerClient
	reportJob                  ifsrmreportjob.ReportJobClient
	report                     ifsrmreport.ReportClient
	reportScheduler            ifsrmreportscheduler.ReportSchedulerClient
	fileManagementJobManager   ifsrmfilemanagementjobmanager.FileManagementJobManagerClient
	fileManagementJob          ifsrmfilemanagementjob.FileManagementJobClient
	propertyCondition          ifsrmpropertycondition.PropertyConditionClient
	fileGroup                  ifsrmfilegroup.FileGroupClient
	fileGroupImported          ifsrmfilegroupimported.FileGroupImportedClient
	fileGroupManager           ifsrmfilegroupmanager.FileGroupManagerClient
	fileScreenBase             ifsrmfilescreenbase.FileScreenBaseClient
	fileScreen                 ifsrmfilescreen.FileScreenClient
	fileScreenException        ifsrmfilescreenexception.FileScreenExceptionClient
	fileScreenManager          ifsrmfilescreenmanager.FileScreenManagerClient
	fileScreenTemplate         ifsrmfilescreentemplate.FileScreenTemplateClient
	fileScreenTemplateImported ifsrmfilescreentemplateimported.FileScreenTemplateImportedClient
	fileScreenTemplateManager  ifsrmfilescreentemplatemanager.FileScreenTemplateManagerClient
}

func (o *xxx_DefaultClient) RemoteUnknown() iremunknown.RemoteUnknownClient {
	return o.dcomClient.RemoteUnknown()
}

func (o *xxx_DefaultClient) RemoteUnknown2() iremunknown2.RemoteUnknown2Client {
	return o.dcomClient.RemoteUnknown2()
}

func (o *xxx_DefaultClient) Object() ifsrmobject.ObjectClient {
	return o.object
}

func (o *xxx_DefaultClient) Collection() ifsrmcollection.CollectionClient {
	return o.collection
}

func (o *xxx_DefaultClient) MutableCollection() ifsrmmutablecollection.MutableCollectionClient {
	return o.mutableCollection
}

func (o *xxx_DefaultClient) CommittableCollection() ifsrmcommittablecollection.CommittableCollectionClient {
	return o.committableCollection
}

func (o *xxx_DefaultClient) Action() ifsrmaction.ActionClient {
	return o.action
}

func (o *xxx_DefaultClient) ActionEmail() ifsrmactionemail.ActionEmailClient {
	return o.actionEmail
}

func (o *xxx_DefaultClient) ActionEmail2() ifsrmactionemail2.ActionEmail2Client {
	return o.actionEmail2
}

func (o *xxx_DefaultClient) ActionReport() ifsrmactionreport.ActionReportClient {
	return o.actionReport
}

func (o *xxx_DefaultClient) ActionEventLog() ifsrmactioneventlog.ActionEventLogClient {
	return o.actionEventLog
}

func (o *xxx_DefaultClient) ActionCommand() ifsrmactioncommand.ActionCommandClient {
	return o.actionCommand
}

func (o *xxx_DefaultClient) Setting() ifsrmsetting.SettingClient {
	return o.setting
}

func (o *xxx_DefaultClient) PathMapper() ifsrmpathmapper.PathMapperClient {
	return o.pathMapper
}

func (o *xxx_DefaultClient) DerivedObjectsResult() ifsrmderivedobjectsresult.DerivedObjectsResultClient {
	return o.derivedObjectsResult
}

func (o *xxx_DefaultClient) PropertyDefinition() ifsrmpropertydefinition.PropertyDefinitionClient {
	return o.propertyDefinition
}

func (o *xxx_DefaultClient) PropertyDefinition2() ifsrmpropertydefinition2.PropertyDefinition2Client {
	return o.propertyDefinition2
}

func (o *xxx_DefaultClient) PropertyDefinitionValue() ifsrmpropertydefinitionvalue.PropertyDefinitionValueClient {
	return o.propertyDefinitionValue
}

func (o *xxx_DefaultClient) Property() ifsrmproperty.PropertyClient {
	return o.property
}

func (o *xxx_DefaultClient) Rule() ifsrmrule.RuleClient {
	return o.rule
}

func (o *xxx_DefaultClient) ClassificationRule() ifsrmclassificationrule.ClassificationRuleClient {
	return o.classificationRule
}

func (o *xxx_DefaultClient) PipelineModuleDefinition() ifsrmpipelinemoduledefinition.PipelineModuleDefinitionClient {
	return o.pipelineModuleDefinition
}

func (o *xxx_DefaultClient) ClassifierModuleDefinition() ifsrmclassifiermoduledefinition.ClassifierModuleDefinitionClient {
	return o.classifierModuleDefinition
}

func (o *xxx_DefaultClient) ClassificationManager() ifsrmclassificationmanager.ClassificationManagerClient {
	return o.classificationManager
}

func (o *xxx_DefaultClient) StorageModuleDefinition() ifsrmstoragemoduledefinition.StorageModuleDefinitionClient {
	return o.storageModuleDefinition
}

func (o *xxx_DefaultClient) QuotaBase() ifsrmquotabase.QuotaBaseClient {
	return o.quotaBase
}

func (o *xxx_DefaultClient) QuotaObject() ifsrmquotaobject.QuotaObjectClient {
	return o.quotaObject
}

func (o *xxx_DefaultClient) Quota() ifsrmquota.QuotaClient {
	return o.quota
}

func (o *xxx_DefaultClient) AutoApplyQuota() ifsrmautoapplyquota.AutoApplyQuotaClient {
	return o.autoApplyQuota
}

func (o *xxx_DefaultClient) QuotaManager() ifsrmquotamanager.QuotaManagerClient {
	return o.quotaManager
}

func (o *xxx_DefaultClient) QuotaManagerEx() ifsrmquotamanagerex.QuotaManagerExClient {
	return o.quotaManagerEx
}

func (o *xxx_DefaultClient) QuotaTemplate() ifsrmquotatemplate.QuotaTemplateClient {
	return o.quotaTemplate
}

func (o *xxx_DefaultClient) QuotaTemplateImported() ifsrmquotatemplateimported.QuotaTemplateImportedClient {
	return o.quotaTemplateImported
}

func (o *xxx_DefaultClient) QuotaTemplateManager() ifsrmquotatemplatemanager.QuotaTemplateManagerClient {
	return o.quotaTemplateManager
}

func (o *xxx_DefaultClient) ReportManager() ifsrmreportmanager.ReportManagerClient {
	return o.reportManager
}

func (o *xxx_DefaultClient) ReportJob() ifsrmreportjob.ReportJobClient {
	return o.reportJob
}

func (o *xxx_DefaultClient) Report() ifsrmreport.ReportClient {
	return o.report
}

func (o *xxx_DefaultClient) ReportScheduler() ifsrmreportscheduler.ReportSchedulerClient {
	return o.reportScheduler
}

func (o *xxx_DefaultClient) FileManagementJobManager() ifsrmfilemanagementjobmanager.FileManagementJobManagerClient {
	return o.fileManagementJobManager
}

func (o *xxx_DefaultClient) FileManagementJob() ifsrmfilemanagementjob.FileManagementJobClient {
	return o.fileManagementJob
}

func (o *xxx_DefaultClient) PropertyCondition() ifsrmpropertycondition.PropertyConditionClient {
	return o.propertyCondition
}

func (o *xxx_DefaultClient) FileGroup() ifsrmfilegroup.FileGroupClient {
	return o.fileGroup
}

func (o *xxx_DefaultClient) FileGroupImported() ifsrmfilegroupimported.FileGroupImportedClient {
	return o.fileGroupImported
}

func (o *xxx_DefaultClient) FileGroupManager() ifsrmfilegroupmanager.FileGroupManagerClient {
	return o.fileGroupManager
}

func (o *xxx_DefaultClient) FileScreenBase() ifsrmfilescreenbase.FileScreenBaseClient {
	return o.fileScreenBase
}

func (o *xxx_DefaultClient) FileScreen() ifsrmfilescreen.FileScreenClient {
	return o.fileScreen
}

func (o *xxx_DefaultClient) FileScreenException() ifsrmfilescreenexception.FileScreenExceptionClient {
	return o.fileScreenException
}

func (o *xxx_DefaultClient) FileScreenManager() ifsrmfilescreenmanager.FileScreenManagerClient {
	return o.fileScreenManager
}

func (o *xxx_DefaultClient) FileScreenTemplate() ifsrmfilescreentemplate.FileScreenTemplateClient {
	return o.fileScreenTemplate
}

func (o *xxx_DefaultClient) FileScreenTemplateImported() ifsrmfilescreentemplateimported.FileScreenTemplateImportedClient {
	return o.fileScreenTemplateImported
}

func (o *xxx_DefaultClient) FileScreenTemplateManager() ifsrmfilescreentemplatemanager.FileScreenTemplateManagerClient {
	return o.fileScreenTemplateManager
}

func NewClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (Client, error) {

	opts = append(opts,
		dcerpc.WithAbstractSyntax(ifsrmobject.ObjectSyntaxV0_0),
		dcerpc.WithAbstractSyntax(ifsrmcollection.CollectionSyntaxV0_0),
		dcerpc.WithAbstractSyntax(ifsrmmutablecollection.MutableCollectionSyntaxV0_0),
		dcerpc.WithAbstractSyntax(ifsrmcommittablecollection.CommittableCollectionSyntaxV0_0),
		dcerpc.WithAbstractSyntax(ifsrmaction.ActionSyntaxV0_0),
		dcerpc.WithAbstractSyntax(ifsrmactionemail.ActionEmailSyntaxV0_0),
		dcerpc.WithAbstractSyntax(ifsrmactionemail2.ActionEmail2SyntaxV0_0),
		dcerpc.WithAbstractSyntax(ifsrmactionreport.ActionReportSyntaxV0_0),
		dcerpc.WithAbstractSyntax(ifsrmactioneventlog.ActionEventLogSyntaxV0_0),
		dcerpc.WithAbstractSyntax(ifsrmactioncommand.ActionCommandSyntaxV0_0),
		dcerpc.WithAbstractSyntax(ifsrmsetting.SettingSyntaxV0_0),
		dcerpc.WithAbstractSyntax(ifsrmpathmapper.PathMapperSyntaxV0_0),
		dcerpc.WithAbstractSyntax(ifsrmderivedobjectsresult.DerivedObjectsResultSyntaxV0_0),
		dcerpc.WithAbstractSyntax(ifsrmpropertydefinition.PropertyDefinitionSyntaxV0_0),
		dcerpc.WithAbstractSyntax(ifsrmpropertydefinition2.PropertyDefinition2SyntaxV0_0),
		dcerpc.WithAbstractSyntax(ifsrmpropertydefinitionvalue.PropertyDefinitionValueSyntaxV0_0),
		dcerpc.WithAbstractSyntax(ifsrmproperty.PropertySyntaxV0_0),
		dcerpc.WithAbstractSyntax(ifsrmrule.RuleSyntaxV0_0),
		dcerpc.WithAbstractSyntax(ifsrmclassificationrule.ClassificationRuleSyntaxV0_0),
		dcerpc.WithAbstractSyntax(ifsrmpipelinemoduledefinition.PipelineModuleDefinitionSyntaxV0_0),
		dcerpc.WithAbstractSyntax(ifsrmclassifiermoduledefinition.ClassifierModuleDefinitionSyntaxV0_0),
		dcerpc.WithAbstractSyntax(ifsrmclassificationmanager.ClassificationManagerSyntaxV0_0),
		dcerpc.WithAbstractSyntax(ifsrmstoragemoduledefinition.StorageModuleDefinitionSyntaxV0_0),
		dcerpc.WithAbstractSyntax(ifsrmquotabase.QuotaBaseSyntaxV0_0),
		dcerpc.WithAbstractSyntax(ifsrmquotaobject.QuotaObjectSyntaxV0_0),
		dcerpc.WithAbstractSyntax(ifsrmquota.QuotaSyntaxV0_0),
		dcerpc.WithAbstractSyntax(ifsrmautoapplyquota.AutoApplyQuotaSyntaxV0_0),
		dcerpc.WithAbstractSyntax(ifsrmquotamanager.QuotaManagerSyntaxV0_0),
		dcerpc.WithAbstractSyntax(ifsrmquotamanagerex.QuotaManagerExSyntaxV0_0),
		dcerpc.WithAbstractSyntax(ifsrmquotatemplate.QuotaTemplateSyntaxV0_0),
		dcerpc.WithAbstractSyntax(ifsrmquotatemplateimported.QuotaTemplateImportedSyntaxV0_0),
		dcerpc.WithAbstractSyntax(ifsrmquotatemplatemanager.QuotaTemplateManagerSyntaxV0_0),
		dcerpc.WithAbstractSyntax(ifsrmreportmanager.ReportManagerSyntaxV1_0),
		dcerpc.WithAbstractSyntax(ifsrmreportjob.ReportJobSyntaxV1_0),
		dcerpc.WithAbstractSyntax(ifsrmreport.ReportSyntaxV1_0),
		dcerpc.WithAbstractSyntax(ifsrmreportscheduler.ReportSchedulerSyntaxV1_0),
		dcerpc.WithAbstractSyntax(ifsrmfilemanagementjobmanager.FileManagementJobManagerSyntaxV1_0),
		dcerpc.WithAbstractSyntax(ifsrmfilemanagementjob.FileManagementJobSyntaxV1_0),
		dcerpc.WithAbstractSyntax(ifsrmpropertycondition.PropertyConditionSyntaxV1_0),
		dcerpc.WithAbstractSyntax(ifsrmfilegroup.FileGroupSyntaxV0_0),
		dcerpc.WithAbstractSyntax(ifsrmfilegroupimported.FileGroupImportedSyntaxV0_0),
		dcerpc.WithAbstractSyntax(ifsrmfilegroupmanager.FileGroupManagerSyntaxV0_0),
		dcerpc.WithAbstractSyntax(ifsrmfilescreenbase.FileScreenBaseSyntaxV0_0),
		dcerpc.WithAbstractSyntax(ifsrmfilescreen.FileScreenSyntaxV0_0),
		dcerpc.WithAbstractSyntax(ifsrmfilescreenexception.FileScreenExceptionSyntaxV0_0),
		dcerpc.WithAbstractSyntax(ifsrmfilescreenmanager.FileScreenManagerSyntaxV0_0),
		dcerpc.WithAbstractSyntax(ifsrmfilescreentemplate.FileScreenTemplateSyntaxV0_0),
		dcerpc.WithAbstractSyntax(ifsrmfilescreentemplateimported.FileScreenTemplateImportedSyntaxV0_0),
		dcerpc.WithAbstractSyntax(ifsrmfilescreentemplatemanager.FileScreenTemplateManagerSyntaxV0_0),
		dcerpc.WithAbstractSyntax(iremunknown.RemoteUnknownSyntaxV0_0),
		dcerpc.WithAbstractSyntax(iremunknown2.RemoteUnknown2SyntaxV0_0),
	)

	cc, err := cc.Bind(ctx, opts...)
	if err != nil {
		return nil, err
	}

	o := &xxx_DefaultClient{cc: cc}

	dcomClient, err := dcom_client.NewClient(ctx, cc, append(opts, dcerpc.WithNoBind(cc))...)
	if err != nil {
		return nil, err
	}
	o.dcomClient = dcomClient

	sub, ok := cc.(dcerpc.SubConn)
	if !ok {
		return nil, fmt.Errorf("sub-conn is not supported")
	}

	objectSubConn, err := sub.SubConn(ctx, ifsrmobject.ObjectSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		objectSubConn = sub
	}

	o.object, err = ifsrmobject.NewObjectClient(ctx, objectSubConn, append(opts, dcerpc.WithNoBind(objectSubConn))...)

	collectionSubConn, err := sub.SubConn(ctx, ifsrmcollection.CollectionSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		collectionSubConn = sub
	}

	o.collection, err = ifsrmcollection.NewCollectionClient(ctx, collectionSubConn, append(opts, dcerpc.WithNoBind(collectionSubConn))...)

	mutableCollectionSubConn, err := sub.SubConn(ctx, ifsrmmutablecollection.MutableCollectionSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		mutableCollectionSubConn = sub
	}

	o.mutableCollection, err = ifsrmmutablecollection.NewMutableCollectionClient(ctx, mutableCollectionSubConn, append(opts, dcerpc.WithNoBind(mutableCollectionSubConn))...)

	committableCollectionSubConn, err := sub.SubConn(ctx, ifsrmcommittablecollection.CommittableCollectionSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		committableCollectionSubConn = sub
	}

	o.committableCollection, err = ifsrmcommittablecollection.NewCommittableCollectionClient(ctx, committableCollectionSubConn, append(opts, dcerpc.WithNoBind(committableCollectionSubConn))...)

	actionSubConn, err := sub.SubConn(ctx, ifsrmaction.ActionSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		actionSubConn = sub
	}

	o.action, err = ifsrmaction.NewActionClient(ctx, actionSubConn, append(opts, dcerpc.WithNoBind(actionSubConn))...)

	actionEmailSubConn, err := sub.SubConn(ctx, ifsrmactionemail.ActionEmailSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		actionEmailSubConn = sub
	}

	o.actionEmail, err = ifsrmactionemail.NewActionEmailClient(ctx, actionEmailSubConn, append(opts, dcerpc.WithNoBind(actionEmailSubConn))...)

	actionEmail2SubConn, err := sub.SubConn(ctx, ifsrmactionemail2.ActionEmail2SyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		actionEmail2SubConn = sub
	}

	o.actionEmail2, err = ifsrmactionemail2.NewActionEmail2Client(ctx, actionEmail2SubConn, append(opts, dcerpc.WithNoBind(actionEmail2SubConn))...)

	actionReportSubConn, err := sub.SubConn(ctx, ifsrmactionreport.ActionReportSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		actionReportSubConn = sub
	}

	o.actionReport, err = ifsrmactionreport.NewActionReportClient(ctx, actionReportSubConn, append(opts, dcerpc.WithNoBind(actionReportSubConn))...)

	actionEventLogSubConn, err := sub.SubConn(ctx, ifsrmactioneventlog.ActionEventLogSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		actionEventLogSubConn = sub
	}

	o.actionEventLog, err = ifsrmactioneventlog.NewActionEventLogClient(ctx, actionEventLogSubConn, append(opts, dcerpc.WithNoBind(actionEventLogSubConn))...)

	actionCommandSubConn, err := sub.SubConn(ctx, ifsrmactioncommand.ActionCommandSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		actionCommandSubConn = sub
	}

	o.actionCommand, err = ifsrmactioncommand.NewActionCommandClient(ctx, actionCommandSubConn, append(opts, dcerpc.WithNoBind(actionCommandSubConn))...)

	settingSubConn, err := sub.SubConn(ctx, ifsrmsetting.SettingSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		settingSubConn = sub
	}

	o.setting, err = ifsrmsetting.NewSettingClient(ctx, settingSubConn, append(opts, dcerpc.WithNoBind(settingSubConn))...)

	pathMapperSubConn, err := sub.SubConn(ctx, ifsrmpathmapper.PathMapperSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		pathMapperSubConn = sub
	}

	o.pathMapper, err = ifsrmpathmapper.NewPathMapperClient(ctx, pathMapperSubConn, append(opts, dcerpc.WithNoBind(pathMapperSubConn))...)

	derivedObjectsResultSubConn, err := sub.SubConn(ctx, ifsrmderivedobjectsresult.DerivedObjectsResultSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		derivedObjectsResultSubConn = sub
	}

	o.derivedObjectsResult, err = ifsrmderivedobjectsresult.NewDerivedObjectsResultClient(ctx, derivedObjectsResultSubConn, append(opts, dcerpc.WithNoBind(derivedObjectsResultSubConn))...)

	propertyDefinitionSubConn, err := sub.SubConn(ctx, ifsrmpropertydefinition.PropertyDefinitionSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		propertyDefinitionSubConn = sub
	}

	o.propertyDefinition, err = ifsrmpropertydefinition.NewPropertyDefinitionClient(ctx, propertyDefinitionSubConn, append(opts, dcerpc.WithNoBind(propertyDefinitionSubConn))...)

	propertyDefinition2SubConn, err := sub.SubConn(ctx, ifsrmpropertydefinition2.PropertyDefinition2SyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		propertyDefinition2SubConn = sub
	}

	o.propertyDefinition2, err = ifsrmpropertydefinition2.NewPropertyDefinition2Client(ctx, propertyDefinition2SubConn, append(opts, dcerpc.WithNoBind(propertyDefinition2SubConn))...)

	propertyDefinitionValueSubConn, err := sub.SubConn(ctx, ifsrmpropertydefinitionvalue.PropertyDefinitionValueSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		propertyDefinitionValueSubConn = sub
	}

	o.propertyDefinitionValue, err = ifsrmpropertydefinitionvalue.NewPropertyDefinitionValueClient(ctx, propertyDefinitionValueSubConn, append(opts, dcerpc.WithNoBind(propertyDefinitionValueSubConn))...)

	propertySubConn, err := sub.SubConn(ctx, ifsrmproperty.PropertySyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		propertySubConn = sub
	}

	o.property, err = ifsrmproperty.NewPropertyClient(ctx, propertySubConn, append(opts, dcerpc.WithNoBind(propertySubConn))...)

	ruleSubConn, err := sub.SubConn(ctx, ifsrmrule.RuleSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		ruleSubConn = sub
	}

	o.rule, err = ifsrmrule.NewRuleClient(ctx, ruleSubConn, append(opts, dcerpc.WithNoBind(ruleSubConn))...)

	classificationRuleSubConn, err := sub.SubConn(ctx, ifsrmclassificationrule.ClassificationRuleSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		classificationRuleSubConn = sub
	}

	o.classificationRule, err = ifsrmclassificationrule.NewClassificationRuleClient(ctx, classificationRuleSubConn, append(opts, dcerpc.WithNoBind(classificationRuleSubConn))...)

	pipelineModuleDefinitionSubConn, err := sub.SubConn(ctx, ifsrmpipelinemoduledefinition.PipelineModuleDefinitionSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		pipelineModuleDefinitionSubConn = sub
	}

	o.pipelineModuleDefinition, err = ifsrmpipelinemoduledefinition.NewPipelineModuleDefinitionClient(ctx, pipelineModuleDefinitionSubConn, append(opts, dcerpc.WithNoBind(pipelineModuleDefinitionSubConn))...)

	classifierModuleDefinitionSubConn, err := sub.SubConn(ctx, ifsrmclassifiermoduledefinition.ClassifierModuleDefinitionSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		classifierModuleDefinitionSubConn = sub
	}

	o.classifierModuleDefinition, err = ifsrmclassifiermoduledefinition.NewClassifierModuleDefinitionClient(ctx, classifierModuleDefinitionSubConn, append(opts, dcerpc.WithNoBind(classifierModuleDefinitionSubConn))...)

	classificationManagerSubConn, err := sub.SubConn(ctx, ifsrmclassificationmanager.ClassificationManagerSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		classificationManagerSubConn = sub
	}

	o.classificationManager, err = ifsrmclassificationmanager.NewClassificationManagerClient(ctx, classificationManagerSubConn, append(opts, dcerpc.WithNoBind(classificationManagerSubConn))...)

	storageModuleDefinitionSubConn, err := sub.SubConn(ctx, ifsrmstoragemoduledefinition.StorageModuleDefinitionSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		storageModuleDefinitionSubConn = sub
	}

	o.storageModuleDefinition, err = ifsrmstoragemoduledefinition.NewStorageModuleDefinitionClient(ctx, storageModuleDefinitionSubConn, append(opts, dcerpc.WithNoBind(storageModuleDefinitionSubConn))...)

	quotaBaseSubConn, err := sub.SubConn(ctx, ifsrmquotabase.QuotaBaseSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		quotaBaseSubConn = sub
	}

	o.quotaBase, err = ifsrmquotabase.NewQuotaBaseClient(ctx, quotaBaseSubConn, append(opts, dcerpc.WithNoBind(quotaBaseSubConn))...)

	quotaObjectSubConn, err := sub.SubConn(ctx, ifsrmquotaobject.QuotaObjectSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		quotaObjectSubConn = sub
	}

	o.quotaObject, err = ifsrmquotaobject.NewQuotaObjectClient(ctx, quotaObjectSubConn, append(opts, dcerpc.WithNoBind(quotaObjectSubConn))...)

	quotaSubConn, err := sub.SubConn(ctx, ifsrmquota.QuotaSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		quotaSubConn = sub
	}

	o.quota, err = ifsrmquota.NewQuotaClient(ctx, quotaSubConn, append(opts, dcerpc.WithNoBind(quotaSubConn))...)

	autoApplyQuotaSubConn, err := sub.SubConn(ctx, ifsrmautoapplyquota.AutoApplyQuotaSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		autoApplyQuotaSubConn = sub
	}

	o.autoApplyQuota, err = ifsrmautoapplyquota.NewAutoApplyQuotaClient(ctx, autoApplyQuotaSubConn, append(opts, dcerpc.WithNoBind(autoApplyQuotaSubConn))...)

	quotaManagerSubConn, err := sub.SubConn(ctx, ifsrmquotamanager.QuotaManagerSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		quotaManagerSubConn = sub
	}

	o.quotaManager, err = ifsrmquotamanager.NewQuotaManagerClient(ctx, quotaManagerSubConn, append(opts, dcerpc.WithNoBind(quotaManagerSubConn))...)

	quotaManagerExSubConn, err := sub.SubConn(ctx, ifsrmquotamanagerex.QuotaManagerExSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		quotaManagerExSubConn = sub
	}

	o.quotaManagerEx, err = ifsrmquotamanagerex.NewQuotaManagerExClient(ctx, quotaManagerExSubConn, append(opts, dcerpc.WithNoBind(quotaManagerExSubConn))...)

	quotaTemplateSubConn, err := sub.SubConn(ctx, ifsrmquotatemplate.QuotaTemplateSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		quotaTemplateSubConn = sub
	}

	o.quotaTemplate, err = ifsrmquotatemplate.NewQuotaTemplateClient(ctx, quotaTemplateSubConn, append(opts, dcerpc.WithNoBind(quotaTemplateSubConn))...)

	quotaTemplateImportedSubConn, err := sub.SubConn(ctx, ifsrmquotatemplateimported.QuotaTemplateImportedSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		quotaTemplateImportedSubConn = sub
	}

	o.quotaTemplateImported, err = ifsrmquotatemplateimported.NewQuotaTemplateImportedClient(ctx, quotaTemplateImportedSubConn, append(opts, dcerpc.WithNoBind(quotaTemplateImportedSubConn))...)

	quotaTemplateManagerSubConn, err := sub.SubConn(ctx, ifsrmquotatemplatemanager.QuotaTemplateManagerSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		quotaTemplateManagerSubConn = sub
	}

	o.quotaTemplateManager, err = ifsrmquotatemplatemanager.NewQuotaTemplateManagerClient(ctx, quotaTemplateManagerSubConn, append(opts, dcerpc.WithNoBind(quotaTemplateManagerSubConn))...)

	reportManagerSubConn, err := sub.SubConn(ctx, ifsrmreportmanager.ReportManagerSyntaxV1_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		reportManagerSubConn = sub
	}

	o.reportManager, err = ifsrmreportmanager.NewReportManagerClient(ctx, reportManagerSubConn, append(opts, dcerpc.WithNoBind(reportManagerSubConn))...)

	reportJobSubConn, err := sub.SubConn(ctx, ifsrmreportjob.ReportJobSyntaxV1_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		reportJobSubConn = sub
	}

	o.reportJob, err = ifsrmreportjob.NewReportJobClient(ctx, reportJobSubConn, append(opts, dcerpc.WithNoBind(reportJobSubConn))...)

	reportSubConn, err := sub.SubConn(ctx, ifsrmreport.ReportSyntaxV1_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		reportSubConn = sub
	}

	o.report, err = ifsrmreport.NewReportClient(ctx, reportSubConn, append(opts, dcerpc.WithNoBind(reportSubConn))...)

	reportSchedulerSubConn, err := sub.SubConn(ctx, ifsrmreportscheduler.ReportSchedulerSyntaxV1_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		reportSchedulerSubConn = sub
	}

	o.reportScheduler, err = ifsrmreportscheduler.NewReportSchedulerClient(ctx, reportSchedulerSubConn, append(opts, dcerpc.WithNoBind(reportSchedulerSubConn))...)

	fileManagementJobManagerSubConn, err := sub.SubConn(ctx, ifsrmfilemanagementjobmanager.FileManagementJobManagerSyntaxV1_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		fileManagementJobManagerSubConn = sub
	}

	o.fileManagementJobManager, err = ifsrmfilemanagementjobmanager.NewFileManagementJobManagerClient(ctx, fileManagementJobManagerSubConn, append(opts, dcerpc.WithNoBind(fileManagementJobManagerSubConn))...)

	fileManagementJobSubConn, err := sub.SubConn(ctx, ifsrmfilemanagementjob.FileManagementJobSyntaxV1_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		fileManagementJobSubConn = sub
	}

	o.fileManagementJob, err = ifsrmfilemanagementjob.NewFileManagementJobClient(ctx, fileManagementJobSubConn, append(opts, dcerpc.WithNoBind(fileManagementJobSubConn))...)

	propertyConditionSubConn, err := sub.SubConn(ctx, ifsrmpropertycondition.PropertyConditionSyntaxV1_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		propertyConditionSubConn = sub
	}

	o.propertyCondition, err = ifsrmpropertycondition.NewPropertyConditionClient(ctx, propertyConditionSubConn, append(opts, dcerpc.WithNoBind(propertyConditionSubConn))...)

	fileGroupSubConn, err := sub.SubConn(ctx, ifsrmfilegroup.FileGroupSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		fileGroupSubConn = sub
	}

	o.fileGroup, err = ifsrmfilegroup.NewFileGroupClient(ctx, fileGroupSubConn, append(opts, dcerpc.WithNoBind(fileGroupSubConn))...)

	fileGroupImportedSubConn, err := sub.SubConn(ctx, ifsrmfilegroupimported.FileGroupImportedSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		fileGroupImportedSubConn = sub
	}

	o.fileGroupImported, err = ifsrmfilegroupimported.NewFileGroupImportedClient(ctx, fileGroupImportedSubConn, append(opts, dcerpc.WithNoBind(fileGroupImportedSubConn))...)

	fileGroupManagerSubConn, err := sub.SubConn(ctx, ifsrmfilegroupmanager.FileGroupManagerSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		fileGroupManagerSubConn = sub
	}

	o.fileGroupManager, err = ifsrmfilegroupmanager.NewFileGroupManagerClient(ctx, fileGroupManagerSubConn, append(opts, dcerpc.WithNoBind(fileGroupManagerSubConn))...)

	fileScreenBaseSubConn, err := sub.SubConn(ctx, ifsrmfilescreenbase.FileScreenBaseSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		fileScreenBaseSubConn = sub
	}

	o.fileScreenBase, err = ifsrmfilescreenbase.NewFileScreenBaseClient(ctx, fileScreenBaseSubConn, append(opts, dcerpc.WithNoBind(fileScreenBaseSubConn))...)

	fileScreenSubConn, err := sub.SubConn(ctx, ifsrmfilescreen.FileScreenSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		fileScreenSubConn = sub
	}

	o.fileScreen, err = ifsrmfilescreen.NewFileScreenClient(ctx, fileScreenSubConn, append(opts, dcerpc.WithNoBind(fileScreenSubConn))...)

	fileScreenExceptionSubConn, err := sub.SubConn(ctx, ifsrmfilescreenexception.FileScreenExceptionSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		fileScreenExceptionSubConn = sub
	}

	o.fileScreenException, err = ifsrmfilescreenexception.NewFileScreenExceptionClient(ctx, fileScreenExceptionSubConn, append(opts, dcerpc.WithNoBind(fileScreenExceptionSubConn))...)

	fileScreenManagerSubConn, err := sub.SubConn(ctx, ifsrmfilescreenmanager.FileScreenManagerSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		fileScreenManagerSubConn = sub
	}

	o.fileScreenManager, err = ifsrmfilescreenmanager.NewFileScreenManagerClient(ctx, fileScreenManagerSubConn, append(opts, dcerpc.WithNoBind(fileScreenManagerSubConn))...)

	fileScreenTemplateSubConn, err := sub.SubConn(ctx, ifsrmfilescreentemplate.FileScreenTemplateSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		fileScreenTemplateSubConn = sub
	}

	o.fileScreenTemplate, err = ifsrmfilescreentemplate.NewFileScreenTemplateClient(ctx, fileScreenTemplateSubConn, append(opts, dcerpc.WithNoBind(fileScreenTemplateSubConn))...)

	fileScreenTemplateImportedSubConn, err := sub.SubConn(ctx, ifsrmfilescreentemplateimported.FileScreenTemplateImportedSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		fileScreenTemplateImportedSubConn = sub
	}

	o.fileScreenTemplateImported, err = ifsrmfilescreentemplateimported.NewFileScreenTemplateImportedClient(ctx, fileScreenTemplateImportedSubConn, append(opts, dcerpc.WithNoBind(fileScreenTemplateImportedSubConn))...)

	fileScreenTemplateManagerSubConn, err := sub.SubConn(ctx, ifsrmfilescreentemplatemanager.FileScreenTemplateManagerSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		fileScreenTemplateManagerSubConn = sub
	}

	o.fileScreenTemplateManager, err = ifsrmfilescreentemplatemanager.NewFileScreenTemplateManagerClient(ctx, fileScreenTemplateManagerSubConn, append(opts, dcerpc.WithNoBind(fileScreenTemplateManagerSubConn))...)
	return o, nil
}

func (o *xxx_DefaultClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultClient) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultClient) IPID(ctx context.Context, ipid *dcom.IPID) Client {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultClient{
		dcomClient:                 o.dcomClient.IPID(ctx, ipid),
		object:                     o.object.IPID(ctx, ipid),
		collection:                 o.collection.IPID(ctx, ipid),
		mutableCollection:          o.mutableCollection.IPID(ctx, ipid),
		committableCollection:      o.committableCollection.IPID(ctx, ipid),
		action:                     o.action.IPID(ctx, ipid),
		actionEmail:                o.actionEmail.IPID(ctx, ipid),
		actionEmail2:               o.actionEmail2.IPID(ctx, ipid),
		actionReport:               o.actionReport.IPID(ctx, ipid),
		actionEventLog:             o.actionEventLog.IPID(ctx, ipid),
		actionCommand:              o.actionCommand.IPID(ctx, ipid),
		setting:                    o.setting.IPID(ctx, ipid),
		pathMapper:                 o.pathMapper.IPID(ctx, ipid),
		derivedObjectsResult:       o.derivedObjectsResult.IPID(ctx, ipid),
		propertyDefinition:         o.propertyDefinition.IPID(ctx, ipid),
		propertyDefinition2:        o.propertyDefinition2.IPID(ctx, ipid),
		propertyDefinitionValue:    o.propertyDefinitionValue.IPID(ctx, ipid),
		property:                   o.property.IPID(ctx, ipid),
		rule:                       o.rule.IPID(ctx, ipid),
		classificationRule:         o.classificationRule.IPID(ctx, ipid),
		pipelineModuleDefinition:   o.pipelineModuleDefinition.IPID(ctx, ipid),
		classifierModuleDefinition: o.classifierModuleDefinition.IPID(ctx, ipid),
		classificationManager:      o.classificationManager.IPID(ctx, ipid),
		storageModuleDefinition:    o.storageModuleDefinition.IPID(ctx, ipid),
		quotaBase:                  o.quotaBase.IPID(ctx, ipid),
		quotaObject:                o.quotaObject.IPID(ctx, ipid),
		quota:                      o.quota.IPID(ctx, ipid),
		autoApplyQuota:             o.autoApplyQuota.IPID(ctx, ipid),
		quotaManager:               o.quotaManager.IPID(ctx, ipid),
		quotaManagerEx:             o.quotaManagerEx.IPID(ctx, ipid),
		quotaTemplate:              o.quotaTemplate.IPID(ctx, ipid),
		quotaTemplateImported:      o.quotaTemplateImported.IPID(ctx, ipid),
		quotaTemplateManager:       o.quotaTemplateManager.IPID(ctx, ipid),
		reportManager:              o.reportManager.IPID(ctx, ipid),
		reportJob:                  o.reportJob.IPID(ctx, ipid),
		report:                     o.report.IPID(ctx, ipid),
		reportScheduler:            o.reportScheduler.IPID(ctx, ipid),
		fileManagementJobManager:   o.fileManagementJobManager.IPID(ctx, ipid),
		fileManagementJob:          o.fileManagementJob.IPID(ctx, ipid),
		propertyCondition:          o.propertyCondition.IPID(ctx, ipid),
		fileGroup:                  o.fileGroup.IPID(ctx, ipid),
		fileGroupImported:          o.fileGroupImported.IPID(ctx, ipid),
		fileGroupManager:           o.fileGroupManager.IPID(ctx, ipid),
		fileScreenBase:             o.fileScreenBase.IPID(ctx, ipid),
		fileScreen:                 o.fileScreen.IPID(ctx, ipid),
		fileScreenException:        o.fileScreenException.IPID(ctx, ipid),
		fileScreenManager:          o.fileScreenManager.IPID(ctx, ipid),
		fileScreenTemplate:         o.fileScreenTemplate.IPID(ctx, ipid),
		fileScreenTemplateImported: o.fileScreenTemplateImported.IPID(ctx, ipid),
		fileScreenTemplateManager:  o.fileScreenTemplateManager.IPID(ctx, ipid),
		cc:                         o.cc,
	}
}
