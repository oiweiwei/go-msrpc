package iisa

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
	iapphostadminmanager "github.com/oiweiwei/go-msrpc/msrpc/dcom/iisa/iapphostadminmanager/v0"
	iapphostchangehandler "github.com/oiweiwei/go-msrpc/msrpc/dcom/iisa/iapphostchangehandler/v0"
	iapphostchildelementcollection "github.com/oiweiwei/go-msrpc/msrpc/dcom/iisa/iapphostchildelementcollection/v0"
	iapphostcollectionschema "github.com/oiweiwei/go-msrpc/msrpc/dcom/iisa/iapphostcollectionschema/v0"
	iapphostconfigexception "github.com/oiweiwei/go-msrpc/msrpc/dcom/iisa/iapphostconfigexception/v0"
	iapphostconfigfile "github.com/oiweiwei/go-msrpc/msrpc/dcom/iisa/iapphostconfigfile/v0"
	iapphostconfiglocation "github.com/oiweiwei/go-msrpc/msrpc/dcom/iisa/iapphostconfiglocation/v0"
	iapphostconfiglocationcollection "github.com/oiweiwei/go-msrpc/msrpc/dcom/iisa/iapphostconfiglocationcollection/v0"
	iapphostconfigmanager "github.com/oiweiwei/go-msrpc/msrpc/dcom/iisa/iapphostconfigmanager/v0"
	iapphostconstantvalue "github.com/oiweiwei/go-msrpc/msrpc/dcom/iisa/iapphostconstantvalue/v0"
	iapphostconstantvaluecollection "github.com/oiweiwei/go-msrpc/msrpc/dcom/iisa/iapphostconstantvaluecollection/v0"
	iapphostelement "github.com/oiweiwei/go-msrpc/msrpc/dcom/iisa/iapphostelement/v0"
	iapphostelementcollection "github.com/oiweiwei/go-msrpc/msrpc/dcom/iisa/iapphostelementcollection/v0"
	iapphostelementschema "github.com/oiweiwei/go-msrpc/msrpc/dcom/iisa/iapphostelementschema/v0"
	iapphostelementschemacollection "github.com/oiweiwei/go-msrpc/msrpc/dcom/iisa/iapphostelementschemacollection/v0"
	iapphostmappingextension "github.com/oiweiwei/go-msrpc/msrpc/dcom/iisa/iapphostmappingextension/v0"
	iapphostmethod "github.com/oiweiwei/go-msrpc/msrpc/dcom/iisa/iapphostmethod/v0"
	iapphostmethodcollection "github.com/oiweiwei/go-msrpc/msrpc/dcom/iisa/iapphostmethodcollection/v0"
	iapphostmethodinstance "github.com/oiweiwei/go-msrpc/msrpc/dcom/iisa/iapphostmethodinstance/v0"
	iapphostmethodschema "github.com/oiweiwei/go-msrpc/msrpc/dcom/iisa/iapphostmethodschema/v0"
	iapphostpathmapper "github.com/oiweiwei/go-msrpc/msrpc/dcom/iisa/iapphostpathmapper/v0"
	iapphostproperty "github.com/oiweiwei/go-msrpc/msrpc/dcom/iisa/iapphostproperty/v0"
	iapphostpropertycollection "github.com/oiweiwei/go-msrpc/msrpc/dcom/iisa/iapphostpropertycollection/v0"
	iapphostpropertyexception "github.com/oiweiwei/go-msrpc/msrpc/dcom/iisa/iapphostpropertyexception/v0"
	iapphostpropertyschema "github.com/oiweiwei/go-msrpc/msrpc/dcom/iisa/iapphostpropertyschema/v0"
	iapphostpropertyschemacollection "github.com/oiweiwei/go-msrpc/msrpc/dcom/iisa/iapphostpropertyschemacollection/v0"
	iapphostsectiondefinition "github.com/oiweiwei/go-msrpc/msrpc/dcom/iisa/iapphostsectiondefinition/v0"
	iapphostsectiondefinitioncollection "github.com/oiweiwei/go-msrpc/msrpc/dcom/iisa/iapphostsectiondefinitioncollection/v0"
	iapphostsectiongroup "github.com/oiweiwei/go-msrpc/msrpc/dcom/iisa/iapphostsectiongroup/v0"
	iapphostwritableadminmanager "github.com/oiweiwei/go-msrpc/msrpc/dcom/iisa/iapphostwritableadminmanager/v0"
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
	_ = iapphostmappingextension.GoPackage
	_ = iapphostchildelementcollection.GoPackage
	_ = iapphostpropertycollection.GoPackage
	_ = iapphostconfiglocationcollection.GoPackage
	_ = iapphostmethodcollection.GoPackage
	_ = iapphostelementschemacollection.GoPackage
	_ = iapphostpropertyschemacollection.GoPackage
	_ = iapphostconstantvaluecollection.GoPackage
	_ = iapphostconstantvalue.GoPackage
	_ = iapphostpropertyschema.GoPackage
	_ = iapphostcollectionschema.GoPackage
	_ = iapphostelementschema.GoPackage
	_ = iapphostmethodschema.GoPackage
	_ = iapphostmethodinstance.GoPackage
	_ = iapphostmethod.GoPackage
	_ = iapphostconfigexception.GoPackage
	_ = iapphostpropertyexception.GoPackage
	_ = iapphostelementcollection.GoPackage
	_ = iapphostelement.GoPackage
	_ = iapphostproperty.GoPackage
	_ = iapphostconfiglocation.GoPackage
	_ = iapphostsectiondefinition.GoPackage
	_ = iapphostsectiondefinitioncollection.GoPackage
	_ = iapphostsectiongroup.GoPackage
	_ = iapphostconfigfile.GoPackage
	_ = iapphostpathmapper.GoPackage
	_ = iapphostchangehandler.GoPackage
	_ = iapphostadminmanager.GoPackage
	_ = iapphostwritableadminmanager.GoPackage
	_ = iapphostconfigmanager.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/iisa"
)

// dcom/iisa client set.
type Client interface {

	// DCOM common interfaces
	RemoteUnknown() iremunknown.RemoteUnknownClient
	RemoteUnknown2() iremunknown2.RemoteUnknown2Client

	// Package specific interfaces
	AppHostMappingExtension() iapphostmappingextension.AppHostMappingExtensionClient
	AppHostChildElementCollection() iapphostchildelementcollection.AppHostChildElementCollectionClient
	AppHostPropertyCollection() iapphostpropertycollection.AppHostPropertyCollectionClient
	AppHostConfigLocationCollection() iapphostconfiglocationcollection.AppHostConfigLocationCollectionClient
	AppHostMethodCollection() iapphostmethodcollection.AppHostMethodCollectionClient
	AppHostElementSchemaCollection() iapphostelementschemacollection.AppHostElementSchemaCollectionClient
	AppHostPropertySchemaCollection() iapphostpropertyschemacollection.AppHostPropertySchemaCollectionClient
	AppHostConstantValueCollection() iapphostconstantvaluecollection.AppHostConstantValueCollectionClient
	AppHostConstantValue() iapphostconstantvalue.AppHostConstantValueClient
	AppHostPropertySchema() iapphostpropertyschema.AppHostPropertySchemaClient
	AppHostCollectionSchema() iapphostcollectionschema.AppHostCollectionSchemaClient
	AppHostElementSchema() iapphostelementschema.AppHostElementSchemaClient
	AppHostMethodSchema() iapphostmethodschema.AppHostMethodSchemaClient
	AppHostMethodInstance() iapphostmethodinstance.AppHostMethodInstanceClient
	AppHostMethod() iapphostmethod.AppHostMethodClient
	AppHostConfigException() iapphostconfigexception.AppHostConfigExceptionClient
	AppHostPropertyException() iapphostpropertyexception.AppHostPropertyExceptionClient
	AppHostElementCollection() iapphostelementcollection.AppHostElementCollectionClient
	AppHostElement() iapphostelement.AppHostElementClient
	AppHostProperty() iapphostproperty.AppHostPropertyClient
	AppHostConfigLocation() iapphostconfiglocation.AppHostConfigLocationClient
	AppHostSectionDefinition() iapphostsectiondefinition.AppHostSectionDefinitionClient
	AppHostSectionDefinitionCollection() iapphostsectiondefinitioncollection.AppHostSectionDefinitionCollectionClient
	AppHostSectionGroup() iapphostsectiongroup.AppHostSectionGroupClient
	AppHostConfigFile() iapphostconfigfile.AppHostConfigFileClient
	AppHostPathMapper() iapphostpathmapper.AppHostPathMapperClient
	AppHostChangeHandler() iapphostchangehandler.AppHostChangeHandlerClient
	AppHostAdminManager() iapphostadminmanager.AppHostAdminManagerClient
	AppHostWritableAdminManager() iapphostwritableadminmanager.AppHostWritableAdminManagerClient
	AppHostConfigManager() iapphostconfigmanager.AppHostConfigManagerClient
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

	appHostMappingExtension            iapphostmappingextension.AppHostMappingExtensionClient
	appHostChildElementCollection      iapphostchildelementcollection.AppHostChildElementCollectionClient
	appHostPropertyCollection          iapphostpropertycollection.AppHostPropertyCollectionClient
	appHostConfigLocationCollection    iapphostconfiglocationcollection.AppHostConfigLocationCollectionClient
	appHostMethodCollection            iapphostmethodcollection.AppHostMethodCollectionClient
	appHostElementSchemaCollection     iapphostelementschemacollection.AppHostElementSchemaCollectionClient
	appHostPropertySchemaCollection    iapphostpropertyschemacollection.AppHostPropertySchemaCollectionClient
	appHostConstantValueCollection     iapphostconstantvaluecollection.AppHostConstantValueCollectionClient
	appHostConstantValue               iapphostconstantvalue.AppHostConstantValueClient
	appHostPropertySchema              iapphostpropertyschema.AppHostPropertySchemaClient
	appHostCollectionSchema            iapphostcollectionschema.AppHostCollectionSchemaClient
	appHostElementSchema               iapphostelementschema.AppHostElementSchemaClient
	appHostMethodSchema                iapphostmethodschema.AppHostMethodSchemaClient
	appHostMethodInstance              iapphostmethodinstance.AppHostMethodInstanceClient
	appHostMethod                      iapphostmethod.AppHostMethodClient
	appHostConfigException             iapphostconfigexception.AppHostConfigExceptionClient
	appHostPropertyException           iapphostpropertyexception.AppHostPropertyExceptionClient
	appHostElementCollection           iapphostelementcollection.AppHostElementCollectionClient
	appHostElement                     iapphostelement.AppHostElementClient
	appHostProperty                    iapphostproperty.AppHostPropertyClient
	appHostConfigLocation              iapphostconfiglocation.AppHostConfigLocationClient
	appHostSectionDefinition           iapphostsectiondefinition.AppHostSectionDefinitionClient
	appHostSectionDefinitionCollection iapphostsectiondefinitioncollection.AppHostSectionDefinitionCollectionClient
	appHostSectionGroup                iapphostsectiongroup.AppHostSectionGroupClient
	appHostConfigFile                  iapphostconfigfile.AppHostConfigFileClient
	appHostPathMapper                  iapphostpathmapper.AppHostPathMapperClient
	appHostChangeHandler               iapphostchangehandler.AppHostChangeHandlerClient
	appHostAdminManager                iapphostadminmanager.AppHostAdminManagerClient
	appHostWritableAdminManager        iapphostwritableadminmanager.AppHostWritableAdminManagerClient
	appHostConfigManager               iapphostconfigmanager.AppHostConfigManagerClient
}

func (o *xxx_DefaultClient) RemoteUnknown() iremunknown.RemoteUnknownClient {
	return o.dcomClient.RemoteUnknown()
}

func (o *xxx_DefaultClient) RemoteUnknown2() iremunknown2.RemoteUnknown2Client {
	return o.dcomClient.RemoteUnknown2()
}

func (o *xxx_DefaultClient) AppHostMappingExtension() iapphostmappingextension.AppHostMappingExtensionClient {
	return o.appHostMappingExtension
}

func (o *xxx_DefaultClient) AppHostChildElementCollection() iapphostchildelementcollection.AppHostChildElementCollectionClient {
	return o.appHostChildElementCollection
}

func (o *xxx_DefaultClient) AppHostPropertyCollection() iapphostpropertycollection.AppHostPropertyCollectionClient {
	return o.appHostPropertyCollection
}

func (o *xxx_DefaultClient) AppHostConfigLocationCollection() iapphostconfiglocationcollection.AppHostConfigLocationCollectionClient {
	return o.appHostConfigLocationCollection
}

func (o *xxx_DefaultClient) AppHostMethodCollection() iapphostmethodcollection.AppHostMethodCollectionClient {
	return o.appHostMethodCollection
}

func (o *xxx_DefaultClient) AppHostElementSchemaCollection() iapphostelementschemacollection.AppHostElementSchemaCollectionClient {
	return o.appHostElementSchemaCollection
}

func (o *xxx_DefaultClient) AppHostPropertySchemaCollection() iapphostpropertyschemacollection.AppHostPropertySchemaCollectionClient {
	return o.appHostPropertySchemaCollection
}

func (o *xxx_DefaultClient) AppHostConstantValueCollection() iapphostconstantvaluecollection.AppHostConstantValueCollectionClient {
	return o.appHostConstantValueCollection
}

func (o *xxx_DefaultClient) AppHostConstantValue() iapphostconstantvalue.AppHostConstantValueClient {
	return o.appHostConstantValue
}

func (o *xxx_DefaultClient) AppHostPropertySchema() iapphostpropertyschema.AppHostPropertySchemaClient {
	return o.appHostPropertySchema
}

func (o *xxx_DefaultClient) AppHostCollectionSchema() iapphostcollectionschema.AppHostCollectionSchemaClient {
	return o.appHostCollectionSchema
}

func (o *xxx_DefaultClient) AppHostElementSchema() iapphostelementschema.AppHostElementSchemaClient {
	return o.appHostElementSchema
}

func (o *xxx_DefaultClient) AppHostMethodSchema() iapphostmethodschema.AppHostMethodSchemaClient {
	return o.appHostMethodSchema
}

func (o *xxx_DefaultClient) AppHostMethodInstance() iapphostmethodinstance.AppHostMethodInstanceClient {
	return o.appHostMethodInstance
}

func (o *xxx_DefaultClient) AppHostMethod() iapphostmethod.AppHostMethodClient {
	return o.appHostMethod
}

func (o *xxx_DefaultClient) AppHostConfigException() iapphostconfigexception.AppHostConfigExceptionClient {
	return o.appHostConfigException
}

func (o *xxx_DefaultClient) AppHostPropertyException() iapphostpropertyexception.AppHostPropertyExceptionClient {
	return o.appHostPropertyException
}

func (o *xxx_DefaultClient) AppHostElementCollection() iapphostelementcollection.AppHostElementCollectionClient {
	return o.appHostElementCollection
}

func (o *xxx_DefaultClient) AppHostElement() iapphostelement.AppHostElementClient {
	return o.appHostElement
}

func (o *xxx_DefaultClient) AppHostProperty() iapphostproperty.AppHostPropertyClient {
	return o.appHostProperty
}

func (o *xxx_DefaultClient) AppHostConfigLocation() iapphostconfiglocation.AppHostConfigLocationClient {
	return o.appHostConfigLocation
}

func (o *xxx_DefaultClient) AppHostSectionDefinition() iapphostsectiondefinition.AppHostSectionDefinitionClient {
	return o.appHostSectionDefinition
}

func (o *xxx_DefaultClient) AppHostSectionDefinitionCollection() iapphostsectiondefinitioncollection.AppHostSectionDefinitionCollectionClient {
	return o.appHostSectionDefinitionCollection
}

func (o *xxx_DefaultClient) AppHostSectionGroup() iapphostsectiongroup.AppHostSectionGroupClient {
	return o.appHostSectionGroup
}

func (o *xxx_DefaultClient) AppHostConfigFile() iapphostconfigfile.AppHostConfigFileClient {
	return o.appHostConfigFile
}

func (o *xxx_DefaultClient) AppHostPathMapper() iapphostpathmapper.AppHostPathMapperClient {
	return o.appHostPathMapper
}

func (o *xxx_DefaultClient) AppHostChangeHandler() iapphostchangehandler.AppHostChangeHandlerClient {
	return o.appHostChangeHandler
}

func (o *xxx_DefaultClient) AppHostAdminManager() iapphostadminmanager.AppHostAdminManagerClient {
	return o.appHostAdminManager
}

func (o *xxx_DefaultClient) AppHostWritableAdminManager() iapphostwritableadminmanager.AppHostWritableAdminManagerClient {
	return o.appHostWritableAdminManager
}

func (o *xxx_DefaultClient) AppHostConfigManager() iapphostconfigmanager.AppHostConfigManagerClient {
	return o.appHostConfigManager
}

func NewClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (Client, error) {

	opts = append(opts,
		dcerpc.WithAbstractSyntax(iapphostmappingextension.AppHostMappingExtensionSyntaxV0_0),
		dcerpc.WithAbstractSyntax(iapphostchildelementcollection.AppHostChildElementCollectionSyntaxV0_0),
		dcerpc.WithAbstractSyntax(iapphostpropertycollection.AppHostPropertyCollectionSyntaxV0_0),
		dcerpc.WithAbstractSyntax(iapphostconfiglocationcollection.AppHostConfigLocationCollectionSyntaxV0_0),
		dcerpc.WithAbstractSyntax(iapphostmethodcollection.AppHostMethodCollectionSyntaxV0_0),
		dcerpc.WithAbstractSyntax(iapphostelementschemacollection.AppHostElementSchemaCollectionSyntaxV0_0),
		dcerpc.WithAbstractSyntax(iapphostpropertyschemacollection.AppHostPropertySchemaCollectionSyntaxV0_0),
		dcerpc.WithAbstractSyntax(iapphostconstantvaluecollection.AppHostConstantValueCollectionSyntaxV0_0),
		dcerpc.WithAbstractSyntax(iapphostconstantvalue.AppHostConstantValueSyntaxV0_0),
		dcerpc.WithAbstractSyntax(iapphostpropertyschema.AppHostPropertySchemaSyntaxV0_0),
		dcerpc.WithAbstractSyntax(iapphostcollectionschema.AppHostCollectionSchemaSyntaxV0_0),
		dcerpc.WithAbstractSyntax(iapphostelementschema.AppHostElementSchemaSyntaxV0_0),
		dcerpc.WithAbstractSyntax(iapphostmethodschema.AppHostMethodSchemaSyntaxV0_0),
		dcerpc.WithAbstractSyntax(iapphostmethodinstance.AppHostMethodInstanceSyntaxV0_0),
		dcerpc.WithAbstractSyntax(iapphostmethod.AppHostMethodSyntaxV0_0),
		dcerpc.WithAbstractSyntax(iapphostconfigexception.AppHostConfigExceptionSyntaxV0_0),
		dcerpc.WithAbstractSyntax(iapphostpropertyexception.AppHostPropertyExceptionSyntaxV0_0),
		dcerpc.WithAbstractSyntax(iapphostelementcollection.AppHostElementCollectionSyntaxV0_0),
		dcerpc.WithAbstractSyntax(iapphostelement.AppHostElementSyntaxV0_0),
		dcerpc.WithAbstractSyntax(iapphostproperty.AppHostPropertySyntaxV0_0),
		dcerpc.WithAbstractSyntax(iapphostconfiglocation.AppHostConfigLocationSyntaxV0_0),
		dcerpc.WithAbstractSyntax(iapphostsectiondefinition.AppHostSectionDefinitionSyntaxV0_0),
		dcerpc.WithAbstractSyntax(iapphostsectiondefinitioncollection.AppHostSectionDefinitionCollectionSyntaxV0_0),
		dcerpc.WithAbstractSyntax(iapphostsectiongroup.AppHostSectionGroupSyntaxV0_0),
		dcerpc.WithAbstractSyntax(iapphostconfigfile.AppHostConfigFileSyntaxV0_0),
		dcerpc.WithAbstractSyntax(iapphostpathmapper.AppHostPathMapperSyntaxV0_0),
		dcerpc.WithAbstractSyntax(iapphostchangehandler.AppHostChangeHandlerSyntaxV0_0),
		dcerpc.WithAbstractSyntax(iapphostadminmanager.AppHostAdminManagerSyntaxV0_0),
		dcerpc.WithAbstractSyntax(iapphostwritableadminmanager.AppHostWritableAdminManagerSyntaxV0_0),
		dcerpc.WithAbstractSyntax(iapphostconfigmanager.AppHostConfigManagerSyntaxV0_0),
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

	appHostMappingExtensionSubConn, err := sub.SubConn(ctx, iapphostmappingextension.AppHostMappingExtensionSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		appHostMappingExtensionSubConn = sub
	}

	o.appHostMappingExtension, err = iapphostmappingextension.NewAppHostMappingExtensionClient(ctx, appHostMappingExtensionSubConn, append(opts, dcerpc.WithNoBind(appHostMappingExtensionSubConn))...)

	appHostChildElementCollectionSubConn, err := sub.SubConn(ctx, iapphostchildelementcollection.AppHostChildElementCollectionSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		appHostChildElementCollectionSubConn = sub
	}

	o.appHostChildElementCollection, err = iapphostchildelementcollection.NewAppHostChildElementCollectionClient(ctx, appHostChildElementCollectionSubConn, append(opts, dcerpc.WithNoBind(appHostChildElementCollectionSubConn))...)

	appHostPropertyCollectionSubConn, err := sub.SubConn(ctx, iapphostpropertycollection.AppHostPropertyCollectionSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		appHostPropertyCollectionSubConn = sub
	}

	o.appHostPropertyCollection, err = iapphostpropertycollection.NewAppHostPropertyCollectionClient(ctx, appHostPropertyCollectionSubConn, append(opts, dcerpc.WithNoBind(appHostPropertyCollectionSubConn))...)

	appHostConfigLocationCollectionSubConn, err := sub.SubConn(ctx, iapphostconfiglocationcollection.AppHostConfigLocationCollectionSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		appHostConfigLocationCollectionSubConn = sub
	}

	o.appHostConfigLocationCollection, err = iapphostconfiglocationcollection.NewAppHostConfigLocationCollectionClient(ctx, appHostConfigLocationCollectionSubConn, append(opts, dcerpc.WithNoBind(appHostConfigLocationCollectionSubConn))...)

	appHostMethodCollectionSubConn, err := sub.SubConn(ctx, iapphostmethodcollection.AppHostMethodCollectionSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		appHostMethodCollectionSubConn = sub
	}

	o.appHostMethodCollection, err = iapphostmethodcollection.NewAppHostMethodCollectionClient(ctx, appHostMethodCollectionSubConn, append(opts, dcerpc.WithNoBind(appHostMethodCollectionSubConn))...)

	appHostElementSchemaCollectionSubConn, err := sub.SubConn(ctx, iapphostelementschemacollection.AppHostElementSchemaCollectionSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		appHostElementSchemaCollectionSubConn = sub
	}

	o.appHostElementSchemaCollection, err = iapphostelementschemacollection.NewAppHostElementSchemaCollectionClient(ctx, appHostElementSchemaCollectionSubConn, append(opts, dcerpc.WithNoBind(appHostElementSchemaCollectionSubConn))...)

	appHostPropertySchemaCollectionSubConn, err := sub.SubConn(ctx, iapphostpropertyschemacollection.AppHostPropertySchemaCollectionSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		appHostPropertySchemaCollectionSubConn = sub
	}

	o.appHostPropertySchemaCollection, err = iapphostpropertyschemacollection.NewAppHostPropertySchemaCollectionClient(ctx, appHostPropertySchemaCollectionSubConn, append(opts, dcerpc.WithNoBind(appHostPropertySchemaCollectionSubConn))...)

	appHostConstantValueCollectionSubConn, err := sub.SubConn(ctx, iapphostconstantvaluecollection.AppHostConstantValueCollectionSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		appHostConstantValueCollectionSubConn = sub
	}

	o.appHostConstantValueCollection, err = iapphostconstantvaluecollection.NewAppHostConstantValueCollectionClient(ctx, appHostConstantValueCollectionSubConn, append(opts, dcerpc.WithNoBind(appHostConstantValueCollectionSubConn))...)

	appHostConstantValueSubConn, err := sub.SubConn(ctx, iapphostconstantvalue.AppHostConstantValueSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		appHostConstantValueSubConn = sub
	}

	o.appHostConstantValue, err = iapphostconstantvalue.NewAppHostConstantValueClient(ctx, appHostConstantValueSubConn, append(opts, dcerpc.WithNoBind(appHostConstantValueSubConn))...)

	appHostPropertySchemaSubConn, err := sub.SubConn(ctx, iapphostpropertyschema.AppHostPropertySchemaSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		appHostPropertySchemaSubConn = sub
	}

	o.appHostPropertySchema, err = iapphostpropertyschema.NewAppHostPropertySchemaClient(ctx, appHostPropertySchemaSubConn, append(opts, dcerpc.WithNoBind(appHostPropertySchemaSubConn))...)

	appHostCollectionSchemaSubConn, err := sub.SubConn(ctx, iapphostcollectionschema.AppHostCollectionSchemaSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		appHostCollectionSchemaSubConn = sub
	}

	o.appHostCollectionSchema, err = iapphostcollectionschema.NewAppHostCollectionSchemaClient(ctx, appHostCollectionSchemaSubConn, append(opts, dcerpc.WithNoBind(appHostCollectionSchemaSubConn))...)

	appHostElementSchemaSubConn, err := sub.SubConn(ctx, iapphostelementschema.AppHostElementSchemaSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		appHostElementSchemaSubConn = sub
	}

	o.appHostElementSchema, err = iapphostelementschema.NewAppHostElementSchemaClient(ctx, appHostElementSchemaSubConn, append(opts, dcerpc.WithNoBind(appHostElementSchemaSubConn))...)

	appHostMethodSchemaSubConn, err := sub.SubConn(ctx, iapphostmethodschema.AppHostMethodSchemaSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		appHostMethodSchemaSubConn = sub
	}

	o.appHostMethodSchema, err = iapphostmethodschema.NewAppHostMethodSchemaClient(ctx, appHostMethodSchemaSubConn, append(opts, dcerpc.WithNoBind(appHostMethodSchemaSubConn))...)

	appHostMethodInstanceSubConn, err := sub.SubConn(ctx, iapphostmethodinstance.AppHostMethodInstanceSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		appHostMethodInstanceSubConn = sub
	}

	o.appHostMethodInstance, err = iapphostmethodinstance.NewAppHostMethodInstanceClient(ctx, appHostMethodInstanceSubConn, append(opts, dcerpc.WithNoBind(appHostMethodInstanceSubConn))...)

	appHostMethodSubConn, err := sub.SubConn(ctx, iapphostmethod.AppHostMethodSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		appHostMethodSubConn = sub
	}

	o.appHostMethod, err = iapphostmethod.NewAppHostMethodClient(ctx, appHostMethodSubConn, append(opts, dcerpc.WithNoBind(appHostMethodSubConn))...)

	appHostConfigExceptionSubConn, err := sub.SubConn(ctx, iapphostconfigexception.AppHostConfigExceptionSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		appHostConfigExceptionSubConn = sub
	}

	o.appHostConfigException, err = iapphostconfigexception.NewAppHostConfigExceptionClient(ctx, appHostConfigExceptionSubConn, append(opts, dcerpc.WithNoBind(appHostConfigExceptionSubConn))...)

	appHostPropertyExceptionSubConn, err := sub.SubConn(ctx, iapphostpropertyexception.AppHostPropertyExceptionSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		appHostPropertyExceptionSubConn = sub
	}

	o.appHostPropertyException, err = iapphostpropertyexception.NewAppHostPropertyExceptionClient(ctx, appHostPropertyExceptionSubConn, append(opts, dcerpc.WithNoBind(appHostPropertyExceptionSubConn))...)

	appHostElementCollectionSubConn, err := sub.SubConn(ctx, iapphostelementcollection.AppHostElementCollectionSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		appHostElementCollectionSubConn = sub
	}

	o.appHostElementCollection, err = iapphostelementcollection.NewAppHostElementCollectionClient(ctx, appHostElementCollectionSubConn, append(opts, dcerpc.WithNoBind(appHostElementCollectionSubConn))...)

	appHostElementSubConn, err := sub.SubConn(ctx, iapphostelement.AppHostElementSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		appHostElementSubConn = sub
	}

	o.appHostElement, err = iapphostelement.NewAppHostElementClient(ctx, appHostElementSubConn, append(opts, dcerpc.WithNoBind(appHostElementSubConn))...)

	appHostPropertySubConn, err := sub.SubConn(ctx, iapphostproperty.AppHostPropertySyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		appHostPropertySubConn = sub
	}

	o.appHostProperty, err = iapphostproperty.NewAppHostPropertyClient(ctx, appHostPropertySubConn, append(opts, dcerpc.WithNoBind(appHostPropertySubConn))...)

	appHostConfigLocationSubConn, err := sub.SubConn(ctx, iapphostconfiglocation.AppHostConfigLocationSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		appHostConfigLocationSubConn = sub
	}

	o.appHostConfigLocation, err = iapphostconfiglocation.NewAppHostConfigLocationClient(ctx, appHostConfigLocationSubConn, append(opts, dcerpc.WithNoBind(appHostConfigLocationSubConn))...)

	appHostSectionDefinitionSubConn, err := sub.SubConn(ctx, iapphostsectiondefinition.AppHostSectionDefinitionSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		appHostSectionDefinitionSubConn = sub
	}

	o.appHostSectionDefinition, err = iapphostsectiondefinition.NewAppHostSectionDefinitionClient(ctx, appHostSectionDefinitionSubConn, append(opts, dcerpc.WithNoBind(appHostSectionDefinitionSubConn))...)

	appHostSectionDefinitionCollectionSubConn, err := sub.SubConn(ctx, iapphostsectiondefinitioncollection.AppHostSectionDefinitionCollectionSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		appHostSectionDefinitionCollectionSubConn = sub
	}

	o.appHostSectionDefinitionCollection, err = iapphostsectiondefinitioncollection.NewAppHostSectionDefinitionCollectionClient(ctx, appHostSectionDefinitionCollectionSubConn, append(opts, dcerpc.WithNoBind(appHostSectionDefinitionCollectionSubConn))...)

	appHostSectionGroupSubConn, err := sub.SubConn(ctx, iapphostsectiongroup.AppHostSectionGroupSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		appHostSectionGroupSubConn = sub
	}

	o.appHostSectionGroup, err = iapphostsectiongroup.NewAppHostSectionGroupClient(ctx, appHostSectionGroupSubConn, append(opts, dcerpc.WithNoBind(appHostSectionGroupSubConn))...)

	appHostConfigFileSubConn, err := sub.SubConn(ctx, iapphostconfigfile.AppHostConfigFileSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		appHostConfigFileSubConn = sub
	}

	o.appHostConfigFile, err = iapphostconfigfile.NewAppHostConfigFileClient(ctx, appHostConfigFileSubConn, append(opts, dcerpc.WithNoBind(appHostConfigFileSubConn))...)

	appHostPathMapperSubConn, err := sub.SubConn(ctx, iapphostpathmapper.AppHostPathMapperSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		appHostPathMapperSubConn = sub
	}

	o.appHostPathMapper, err = iapphostpathmapper.NewAppHostPathMapperClient(ctx, appHostPathMapperSubConn, append(opts, dcerpc.WithNoBind(appHostPathMapperSubConn))...)

	appHostChangeHandlerSubConn, err := sub.SubConn(ctx, iapphostchangehandler.AppHostChangeHandlerSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		appHostChangeHandlerSubConn = sub
	}

	o.appHostChangeHandler, err = iapphostchangehandler.NewAppHostChangeHandlerClient(ctx, appHostChangeHandlerSubConn, append(opts, dcerpc.WithNoBind(appHostChangeHandlerSubConn))...)

	appHostAdminManagerSubConn, err := sub.SubConn(ctx, iapphostadminmanager.AppHostAdminManagerSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		appHostAdminManagerSubConn = sub
	}

	o.appHostAdminManager, err = iapphostadminmanager.NewAppHostAdminManagerClient(ctx, appHostAdminManagerSubConn, append(opts, dcerpc.WithNoBind(appHostAdminManagerSubConn))...)

	appHostWritableAdminManagerSubConn, err := sub.SubConn(ctx, iapphostwritableadminmanager.AppHostWritableAdminManagerSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		appHostWritableAdminManagerSubConn = sub
	}

	o.appHostWritableAdminManager, err = iapphostwritableadminmanager.NewAppHostWritableAdminManagerClient(ctx, appHostWritableAdminManagerSubConn, append(opts, dcerpc.WithNoBind(appHostWritableAdminManagerSubConn))...)

	appHostConfigManagerSubConn, err := sub.SubConn(ctx, iapphostconfigmanager.AppHostConfigManagerSyntaxV0_0)
	if err != nil {
		// XXX: use main subconnection as a last resort
		// it was noticed that we can reuse the main connection for dcom interfaces
		appHostConfigManagerSubConn = sub
	}

	o.appHostConfigManager, err = iapphostconfigmanager.NewAppHostConfigManagerClient(ctx, appHostConfigManagerSubConn, append(opts, dcerpc.WithNoBind(appHostConfigManagerSubConn))...)
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
		dcomClient:                         o.dcomClient.IPID(ctx, ipid),
		appHostMappingExtension:            o.appHostMappingExtension.IPID(ctx, ipid),
		appHostChildElementCollection:      o.appHostChildElementCollection.IPID(ctx, ipid),
		appHostPropertyCollection:          o.appHostPropertyCollection.IPID(ctx, ipid),
		appHostConfigLocationCollection:    o.appHostConfigLocationCollection.IPID(ctx, ipid),
		appHostMethodCollection:            o.appHostMethodCollection.IPID(ctx, ipid),
		appHostElementSchemaCollection:     o.appHostElementSchemaCollection.IPID(ctx, ipid),
		appHostPropertySchemaCollection:    o.appHostPropertySchemaCollection.IPID(ctx, ipid),
		appHostConstantValueCollection:     o.appHostConstantValueCollection.IPID(ctx, ipid),
		appHostConstantValue:               o.appHostConstantValue.IPID(ctx, ipid),
		appHostPropertySchema:              o.appHostPropertySchema.IPID(ctx, ipid),
		appHostCollectionSchema:            o.appHostCollectionSchema.IPID(ctx, ipid),
		appHostElementSchema:               o.appHostElementSchema.IPID(ctx, ipid),
		appHostMethodSchema:                o.appHostMethodSchema.IPID(ctx, ipid),
		appHostMethodInstance:              o.appHostMethodInstance.IPID(ctx, ipid),
		appHostMethod:                      o.appHostMethod.IPID(ctx, ipid),
		appHostConfigException:             o.appHostConfigException.IPID(ctx, ipid),
		appHostPropertyException:           o.appHostPropertyException.IPID(ctx, ipid),
		appHostElementCollection:           o.appHostElementCollection.IPID(ctx, ipid),
		appHostElement:                     o.appHostElement.IPID(ctx, ipid),
		appHostProperty:                    o.appHostProperty.IPID(ctx, ipid),
		appHostConfigLocation:              o.appHostConfigLocation.IPID(ctx, ipid),
		appHostSectionDefinition:           o.appHostSectionDefinition.IPID(ctx, ipid),
		appHostSectionDefinitionCollection: o.appHostSectionDefinitionCollection.IPID(ctx, ipid),
		appHostSectionGroup:                o.appHostSectionGroup.IPID(ctx, ipid),
		appHostConfigFile:                  o.appHostConfigFile.IPID(ctx, ipid),
		appHostPathMapper:                  o.appHostPathMapper.IPID(ctx, ipid),
		appHostChangeHandler:               o.appHostChangeHandler.IPID(ctx, ipid),
		appHostAdminManager:                o.appHostAdminManager.IPID(ctx, ipid),
		appHostWritableAdminManager:        o.appHostWritableAdminManager.IPID(ctx, ipid),
		appHostConfigManager:               o.appHostConfigManager.IPID(ctx, ipid),
		cc:                                 o.cc,
	}
}
