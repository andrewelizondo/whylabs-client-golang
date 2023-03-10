# Go API client for openapi

WhyLabs API that enables end-to-end AI observability

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 0.1
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen
For more information, please visit [https://whylabs.ai](https://whylabs.ai)

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import openapi "github.com/andrewelizondo/whylabs-client-golang"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), openapi.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), openapi.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), openapi.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), openapi.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *http://localhost*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AdminApi* | [**PostMonitorConfigValidationJob**](docs/AdminApi.md#postmonitorconfigvalidationjob) | **Post** /v0/admin/monitor-config/create-validation-job | Create a monitor config validation job for all configs
*AlertsApi* | [**GetAlertsPaths**](docs/AlertsApi.md#getalertspaths) | **Get** /v0/organizations/{org_id}/alerts/models/{model_id}/paths | Get the alerts for a given time period.
*ApiKeyApi* | [**CreateApiKey**](docs/ApiKeyApi.md#createapikey) | **Post** /v0/organizations/{org_id}/api-key | Generate an API key for a user.
*ApiKeyApi* | [**GetApiKey**](docs/ApiKeyApi.md#getapikey) | **Get** /v0/organizations/{org_id}/api-key/{key_id} | Get an api key by its id
*ApiKeyApi* | [**ListApiKeys**](docs/ApiKeyApi.md#listapikeys) | **Get** /v0/organizations/{org_id}/api-key | List API key metadata for a given organization and user
*ApiKeyApi* | [**RevokeApiKey**](docs/ApiKeyApi.md#revokeapikey) | **Delete** /v0/organizations/{org_id}/api-key | Revoke the given API Key, removing its ability to access WhyLabs systems
*DatabricksApi* | [**GetConnection**](docs/DatabricksApi.md#getconnection) | **Post** /v0/databricks/get-connection | Get the connection metadata for a given org
*DatabricksApi* | [**ListJobs**](docs/DatabricksApi.md#listjobs) | **Post** /v0/databricks/list-jobs | List all of the jobs in a workspace.
*DatabricksApi* | [**RefreshConnection**](docs/DatabricksApi.md#refreshconnection) | **Post** /v0/databricks/refresh-connection | Refresh metadata for a workspace connection.
*DatabricksApi* | [**RunJob**](docs/DatabricksApi.md#runjob) | **Post** /v0/databricks/run-job | Run an existing job in a given databricks workspace.
*DatabricksApi* | [**UpdateConnection**](docs/DatabricksApi.md#updateconnection) | **Post** /v0/databricks/update-connection | Update the connection metadata for a given org
*DatasetProfileApi* | [**DeleteAnalyzerResults**](docs/DatasetProfileApi.md#deleteanalyzerresults) | **Delete** /v0/organizations/{org_id}/dataset-profiles/models/{dataset_id}/analyzer-results | Deletes a set of analyzer results
*DatasetProfileApi* | [**DeleteDatasetProfiles**](docs/DatasetProfileApi.md#deletedatasetprofiles) | **Delete** /v0/organizations/{org_id}/dataset-profiles/models/{dataset_id} | Deletes a set of dataset profiles
*DatasetProfileApi* | [**DeleteReferenceProfile**](docs/DatasetProfileApi.md#deletereferenceprofile) | **Delete** /v0/organizations/{org_id}/dataset-profiles/models/{model_id}/reference-profiles/{reference_id} | Delete a single reference profile
*DatasetProfileApi* | [**GetReferenceProfile**](docs/DatasetProfileApi.md#getreferenceprofile) | **Get** /v0/organizations/{org_id}/dataset-profiles/models/{model_id}/reference-profiles/{reference_id} | Returns a single reference profile
*DatasetProfileApi* | [**ListReferenceProfiles**](docs/DatasetProfileApi.md#listreferenceprofiles) | **Get** /v0/organizations/{org_id}/dataset-profiles/models/{model_id}/reference-profiles | Returns a list for reference profiles
*EventsApi* | [**GetEventsData**](docs/EventsApi.md#geteventsdata) | **Get** /v0/organizations/{org_id}/events/models/{model_id}/data | Get the event data as multi-line JSON for a given time period.
*EventsApi* | [**GetEventsPaths**](docs/EventsApi.md#geteventspaths) | **Get** /v0/organizations/{org_id}/events/models/{model_id}/paths | Get the events for a given time period.
*FeatureFlagsApi* | [**GetFeatureFlags**](docs/FeatureFlagsApi.md#getfeatureflags) | **Get** /v0/feature-flags | Get feature flags for the specified user/org
*FeatureWeightsApi* | [**GetColumnWeights**](docs/FeatureWeightsApi.md#getcolumnweights) | **Get** /v0/organizations/{org_id}/dataset/{dataset_id}/weights | Get column weights for the specified dataset
*FeatureWeightsApi* | [**PutColumnWeights**](docs/FeatureWeightsApi.md#putcolumnweights) | **Put** /v0/organizations/{org_id}/dataset/{dataset_id}/weights | Put column weights for the specified dataset
*InternalApi* | [**CreateOrganization**](docs/InternalApi.md#createorganization) | **Post** /v0/organizations | Create an organization
*InternalApi* | [**CreateUser**](docs/InternalApi.md#createuser) | **Post** /v0/user | Create a user.
*InternalApi* | [**DeleteMonitorConfig**](docs/InternalApi.md#deletemonitorconfig) | **Delete** /v0/organizations/{org_id}/models/{model_id}/monitor-config/ | Delete a monitor config from a model
*InternalApi* | [**DeleteOrganization**](docs/InternalApi.md#deleteorganization) | **Delete** /v0/organizations/{org_id} | Delete an org
*InternalApi* | [**DeleteReferenceProfile**](docs/InternalApi.md#deletereferenceprofile) | **Delete** /v0/organizations/{org_id}/dataset-profiles/models/{model_id}/reference-profiles/{reference_id} | Delete a single reference profile
*InternalApi* | [**GetAWSMarketplaceMetadata**](docs/InternalApi.md#getawsmarketplacemetadata) | **Get** /v0/organizations/{org_id}/marketplace-metadata/ | Get marketplace metadata for an org if any exists.
*InternalApi* | [**GetConnection**](docs/InternalApi.md#getconnection) | **Post** /v0/databricks/get-connection | Get the connection metadata for a given org
*InternalApi* | [**GetEventsPaths**](docs/InternalApi.md#geteventspaths) | **Get** /v0/organizations/{org_id}/events/models/{model_id}/paths | Get the events for a given time period.
*InternalApi* | [**GetFeatureFlags**](docs/InternalApi.md#getfeatureflags) | **Get** /v0/feature-flags | Get feature flags for the specified user/org
*InternalApi* | [**GetMonitorConfig**](docs/InternalApi.md#getmonitorconfig) | **Get** /v0/organizations/{org_id}/models/{model_id}/monitor-config/ | Get the monitor config for a given model + segment
*InternalApi* | [**GetMonitorConfigSchema**](docs/InternalApi.md#getmonitorconfigschema) | **Get** /v0/organizations/{org_id}/schema/monitor-config | Get the current supported schema of the monitor configuration
*InternalApi* | [**GetOrganization**](docs/InternalApi.md#getorganization) | **Get** /v0/organizations/{org_id} | Get the metadata about an organization.
*InternalApi* | [**GetReferenceProfile**](docs/InternalApi.md#getreferenceprofile) | **Get** /v0/organizations/{org_id}/dataset-profiles/models/{model_id}/reference-profiles/{reference_id} | Returns a single reference profile
*InternalApi* | [**GetUser**](docs/InternalApi.md#getuser) | **Get** /v0/user/{user_id} | Get a user by their id.
*InternalApi* | [**GetUserByEmail**](docs/InternalApi.md#getuserbyemail) | **Get** /v0/user | Get a user by their email.
*InternalApi* | [**ListJobs**](docs/InternalApi.md#listjobs) | **Post** /v0/databricks/list-jobs | List all of the jobs in a workspace.
*InternalApi* | [**ListOrganizations**](docs/InternalApi.md#listorganizations) | **Get** /v0/organizations | Get a list of all of the organization ids.
*InternalApi* | [**ListReferenceProfiles**](docs/InternalApi.md#listreferenceprofiles) | **Get** /v0/organizations/{org_id}/dataset-profiles/models/{model_id}/reference-profiles | Returns a list for reference profiles
*InternalApi* | [**PartiallyUpdateOrg**](docs/InternalApi.md#partiallyupdateorg) | **Put** /v0/organizations/partial/ | Update some fields of an organization to non-null values
*InternalApi* | [**PartiallyUpdateOrganization**](docs/InternalApi.md#partiallyupdateorganization) | **Put** /v0/organizations/partial/{org_id} | Update some fields of an organization to non-null values
*InternalApi* | [**PostMonitorConfigValidationJob**](docs/InternalApi.md#postmonitorconfigvalidationjob) | **Post** /v0/admin/monitor-config/create-validation-job | Create a monitor config validation job for all configs
*InternalApi* | [**ProvisionAWSMarketplaceNewUser**](docs/InternalApi.md#provisionawsmarketplacenewuser) | **Post** /v0/provision/marketplace/aws/new-user | Create resources for a new user coming from AWS Marketplace
*InternalApi* | [**ProvisionDatabricksConnection**](docs/InternalApi.md#provisiondatabricksconnection) | **Post** /v0/provision/connect/databricks | Create resources for a new user coming from Databricks
*InternalApi* | [**ProvisionNewUser**](docs/InternalApi.md#provisionnewuser) | **Post** /v0/provision/new-user | Create the resources that a new user needs to use WhyLabs via the website.
*InternalApi* | [**PutMonitorConfig**](docs/InternalApi.md#putmonitorconfig) | **Put** /v0/organizations/{org_id}/models/{model_id}/monitor-config/ | Save get the monitor config for a given model + segment
*InternalApi* | [**PutProvidedConfig**](docs/InternalApi.md#putprovidedconfig) | **Put** /v0/organizations/{org_id}/models/provided-config/ | Save a Provided Config for an organization/model/segment
*InternalApi* | [**RefreshConnection**](docs/InternalApi.md#refreshconnection) | **Post** /v0/databricks/refresh-connection | Refresh metadata for a workspace connection.
*InternalApi* | [**RegisterDatabricksConnection**](docs/InternalApi.md#registerdatabricksconnection) | **Post** /v0/provision/connect/databricks/staged | Register databricks metadata, temporarily storing it against a UUID so that it can be used to provision a databricks connection after email authentication
*InternalApi* | [**RunJob**](docs/InternalApi.md#runjob) | **Post** /v0/databricks/run-job | Run an existing job in a given databricks workspace.
*InternalApi* | [**UpdateConnection**](docs/InternalApi.md#updateconnection) | **Post** /v0/databricks/update-connection | Update the connection metadata for a given org
*InternalApi* | [**UpdateOrg**](docs/InternalApi.md#updateorg) | **Put** /v0/organizations | Update an existing organization
*InternalApi* | [**UpdateOrganization**](docs/InternalApi.md#updateorganization) | **Put** /v0/organizations/{org_id} | Update an existing organization
*InternalApi* | [**UpdateUser**](docs/InternalApi.md#updateuser) | **Put** /v0/user | Update a user.
*LogApi* | [**Log**](docs/LogApi.md#log) | **Post** /v0/organizations/{org_id}/log | Log a dataset profile entry to the backend
*LogApi* | [**LogAsync**](docs/LogApi.md#logasync) | **Post** /v0/organizations/{org_id}/log/async/{dataset_id} | Like /log, except this api doesn&#39;t take the actual profile content. It returns an upload link that can be used to upload the profile to.
*LogApi* | [**LogReference**](docs/LogApi.md#logreference) | **Post** /v0/organizations/{org_id}/log/reference/{model_id} | Returns a presigned URL for uploading the reference profile to.
*MembershipApi* | [**CreateMembership**](docs/MembershipApi.md#createmembership) | **Post** /v0/membership | Create a membership for a user, making them apart of an organization. Uses the user&#39;s current email address.
*MembershipApi* | [**CreateOrganizationMembership**](docs/MembershipApi.md#createorganizationmembership) | **Post** /v0/organizations/{org_id}/membership | Create a membership for a user, making them apart of an organization. Uses the user&#39;s current email address.
*MembershipApi* | [**GetDefaultMembershipForEmail**](docs/MembershipApi.md#getdefaultmembershipforemail) | **Get** /v0/membership/default | Get the default membership for a user.
*MembershipApi* | [**GetMemberships**](docs/MembershipApi.md#getmemberships) | **Get** /v0/membership/user/{user_id} | Get memberships for a user.
*MembershipApi* | [**GetMembershipsByEmail**](docs/MembershipApi.md#getmembershipsbyemail) | **Get** /v0/membership/user | Get memberships for a user given that user&#39;s email address.
*MembershipApi* | [**GetMembershipsByOrg**](docs/MembershipApi.md#getmembershipsbyorg) | **Get** /v0/membership/org/{org_id} | Get memberships for an org.
*MembershipApi* | [**ListOrganizationMemberships**](docs/MembershipApi.md#listorganizationmemberships) | **Get** /v0/organizations/{org_id}/membership | List organization memberships
*MembershipApi* | [**RemoveMembershipByEmail**](docs/MembershipApi.md#removemembershipbyemail) | **Delete** /v0/membership | Removes membership in a given org from a user, using the user&#39;s email address.
*MembershipApi* | [**RemoveOrganizationMembership**](docs/MembershipApi.md#removeorganizationmembership) | **Delete** /v0/organizations/{org_id}/membership | Removes membership in a given org from a user, using the user&#39;s email address.
*MembershipApi* | [**SetDefaultMembership**](docs/MembershipApi.md#setdefaultmembership) | **Post** /v0/membership/default | Sets the organization that should be used when logging a user in
*MembershipApi* | [**UpdateMembershipByEmail**](docs/MembershipApi.md#updatemembershipbyemail) | **Put** /v0/membership | Updates the role in an membership
*MembershipApi* | [**UpdateOrganizationMembership**](docs/MembershipApi.md#updateorganizationmembership) | **Put** /v0/organizations/{org_id}/membership | Updates the role in an membership
*ModelsApi* | [**CreateModel**](docs/ModelsApi.md#createmodel) | **Post** /v0/organizations/{org_id}/models | Create a model with a given name and a time period
*ModelsApi* | [**DeactivateModel**](docs/ModelsApi.md#deactivatemodel) | **Delete** /v0/organizations/{org_id}/models/{model_id} | Mark a model as inactive
*ModelsApi* | [**DeleteAnalyzer**](docs/ModelsApi.md#deleteanalyzer) | **Delete** /v0/organizations/{org_id}/models/{dataset_id}/monitor-config/analyzer/{analyzer_id} | Delete the analyzer config for a given dataset.
*ModelsApi* | [**DeleteEntitySchemaColumn**](docs/ModelsApi.md#deleteentityschemacolumn) | **Delete** /v0/organizations/{org_id}/models/{dataset_id}/schema/column/{column_id} | Delete the entity schema of a single column for a given dataset.
*ModelsApi* | [**DeleteMonitor**](docs/ModelsApi.md#deletemonitor) | **Delete** /v0/organizations/{org_id}/models/{dataset_id}/monitor-config/monitor/{monitor_id} | Delete the monitor for a given dataset.
*ModelsApi* | [**DeleteMonitorConfig**](docs/ModelsApi.md#deletemonitorconfig) | **Delete** /v0/organizations/{org_id}/models/{model_id}/monitor-config/ | Delete a monitor config from a model
*ModelsApi* | [**GetAnalyzer**](docs/ModelsApi.md#getanalyzer) | **Get** /v0/organizations/{org_id}/models/{dataset_id}/monitor-config/analyzer/{analyzer_id} | Get the analyzer config for a given dataset.
*ModelsApi* | [**GetEntitySchema**](docs/ModelsApi.md#getentityschema) | **Get** /v0/organizations/{org_id}/models/{dataset_id}/schema | Get the entity schema config for a given dataset.
*ModelsApi* | [**GetEntitySchemaColumn**](docs/ModelsApi.md#getentityschemacolumn) | **Get** /v0/organizations/{org_id}/models/{dataset_id}/schema/column/{column_id} | Get the entity schema of a single column for a given dataset.
*ModelsApi* | [**GetModel**](docs/ModelsApi.md#getmodel) | **Get** /v0/organizations/{org_id}/models/{model_id} | Get a model metadata
*ModelsApi* | [**GetMonitor**](docs/ModelsApi.md#getmonitor) | **Get** /v0/organizations/{org_id}/models/{dataset_id}/monitor-config/monitor/{monitor_id} | Get the monitor config for a given dataset.
*ModelsApi* | [**GetMonitorConfig**](docs/ModelsApi.md#getmonitorconfig) | **Get** /v0/organizations/{org_id}/models/{model_id}/monitor-config/ | Get the monitor config for a given model + segment
*ModelsApi* | [**GetMonitorConfigV2**](docs/ModelsApi.md#getmonitorconfigv2) | **Get** /v0/organizations/{org_id}/models/{model_id}/monitor-config/v2 | Get the monitor config for a given model or segment. The return of this api will include default values that we apply over any config that omits portions of the monitor config schema.
*ModelsApi* | [**GetMonitorConfigV3**](docs/ModelsApi.md#getmonitorconfigv3) | **Get** /v0/organizations/{org_id}/models/{dataset_id}/monitor-config/v3 | Get the monitor config document for a given dataset.
*ModelsApi* | [**GetMonitorConfigV3Version**](docs/ModelsApi.md#getmonitorconfigv3version) | **Get** /v0/organizations/{org_id}/models/{dataset_id}/monitor-config/v3/versions/{version_id} | Get the monitor config document version for a given dataset.
*ModelsApi* | [**GetProvidedConfig**](docs/ModelsApi.md#getprovidedconfig) | **Get** /v0/organizations/{org_id}/models/provided-config/ | Get the monitor&#39;s provided config for a given organization/model/segment
*ModelsApi* | [**ListModels**](docs/ModelsApi.md#listmodels) | **Get** /v0/organizations/{org_id}/models | Get a list of all of the model ids for an organization.
*ModelsApi* | [**ListMonitorConfigV3Versions**](docs/ModelsApi.md#listmonitorconfigv3versions) | **Get** /v0/organizations/{org_id}/models/{dataset_id}/monitor-config/v3/versions | List the monitor config document versions for a given dataset.
*ModelsApi* | [**ListSegments**](docs/ModelsApi.md#listsegments) | **Get** /v0/organizations/{org_id}/models/{model_id}/segments | Get a model metadata
*ModelsApi* | [**PutAnalyzer**](docs/ModelsApi.md#putanalyzer) | **Put** /v0/organizations/{org_id}/models/{dataset_id}/monitor-config/analyzer/{analyzer_id} | Save the analyzer config for a given dataset.
*ModelsApi* | [**PutEntitySchema**](docs/ModelsApi.md#putentityschema) | **Put** /v0/organizations/{org_id}/models/{dataset_id}/schema | Save the entity schema config for a given dataset.
*ModelsApi* | [**PutEntitySchemaColumn**](docs/ModelsApi.md#putentityschemacolumn) | **Put** /v0/organizations/{org_id}/models/{dataset_id}/schema/column/{column_id} | Save the entity schema of a single column for a given dataset.
*ModelsApi* | [**PutMonitor**](docs/ModelsApi.md#putmonitor) | **Put** /v0/organizations/{org_id}/models/{dataset_id}/monitor-config/monitor/{monitor_id} | Save the monitor for a given dataset.
*ModelsApi* | [**PutMonitorConfig**](docs/ModelsApi.md#putmonitorconfig) | **Put** /v0/organizations/{org_id}/models/{model_id}/monitor-config/ | Save get the monitor config for a given model + segment
*ModelsApi* | [**PutMonitorConfigV2**](docs/ModelsApi.md#putmonitorconfigv2) | **Put** /v0/organizations/{org_id}/models/{model_id}/monitor-config/v2 | Save the monitor config for a given model or segment
*ModelsApi* | [**PutMonitorConfigV3**](docs/ModelsApi.md#putmonitorconfigv3) | **Put** /v0/organizations/{org_id}/models/{dataset_id}/monitor-config/v3 | Save the monitor config document for a given dataset.
*ModelsApi* | [**PutProvidedConfig**](docs/ModelsApi.md#putprovidedconfig) | **Put** /v0/organizations/{org_id}/models/provided-config/ | Save a Provided Config for an organization/model/segment
*ModelsApi* | [**PutRequestMonitorRunConfig**](docs/ModelsApi.md#putrequestmonitorrunconfig) | **Put** /v0/organizations/{org_id}/models/{dataset_id}/request-monitor-run | Put the RequestMonitorRun config into S3.
*ModelsApi* | [**PutSegments**](docs/ModelsApi.md#putsegments) | **Put** /v0/organizations/{org_id}/models/{model_id}/segments | Add a segment to the dataset
*ModelsApi* | [**UpdateModel**](docs/ModelsApi.md#updatemodel) | **Put** /v0/organizations/{org_id}/models/{model_id} | Update a model&#39;s metadata
*ModelsApi* | [**ValidateMonitorConfigV3**](docs/ModelsApi.md#validatemonitorconfigv3) | **Put** /v0/organizations/{org_id}/models/{dataset_id}/monitor-config/v3/validate | Validate the monitor config document for a given dataset.
*NotificationSettingsApi* | [**DeleteNotificationAction**](docs/NotificationSettingsApi.md#deletenotificationaction) | **Delete** /v0/notification-settings/{org_id}/actions/{action_id} | Delete notification action
*NotificationSettingsApi* | [**GetEmailNotificationActionPayload**](docs/NotificationSettingsApi.md#getemailnotificationactionpayload) | **Get** /v0/notification-settings/actions/email/payload | Get dummy notification action payload
*NotificationSettingsApi* | [**GetNotificationAction**](docs/NotificationSettingsApi.md#getnotificationaction) | **Get** /v0/notification-settings/{org_id}/actions/{action_id} | Get notification action for id
*NotificationSettingsApi* | [**GetNotificationSettings**](docs/NotificationSettingsApi.md#getnotificationsettings) | **Get** /v0/notification-settings/{org_id} | Get notification settings for an org
*NotificationSettingsApi* | [**GetPagerDutyNotificationActionPayload**](docs/NotificationSettingsApi.md#getpagerdutynotificationactionpayload) | **Get** /v0/notification-settings/actions/pagerduty/payload | Get dummy notification action payload
*NotificationSettingsApi* | [**GetSlackNotificationActionPayload**](docs/NotificationSettingsApi.md#getslacknotificationactionpayload) | **Get** /v0/notification-settings/actions/slack/payload | Get dummy notification action payload
*NotificationSettingsApi* | [**ListNotificationActions**](docs/NotificationSettingsApi.md#listnotificationactions) | **Get** /v0/notification-settings/{org_id}/actions | List notification actions for an org
*NotificationSettingsApi* | [**PutNotificationAction**](docs/NotificationSettingsApi.md#putnotificationaction) | **Put** /v0/notification-settings/{org_id}/actions/{type}/{action_id} | Add new notification action
*NotificationSettingsApi* | [**UpdateNotificationAction**](docs/NotificationSettingsApi.md#updatenotificationaction) | **Patch** /v0/notification-settings/{org_id}/actions/{type}/{action_id} | Update notification action
*NotificationSettingsApi* | [**UpdateNotificationSettings**](docs/NotificationSettingsApi.md#updatenotificationsettings) | **Post** /v0/notification-settings/{org_id} | Update notification settings for an org
*OrganizationsApi* | [**CreateOrganization**](docs/OrganizationsApi.md#createorganization) | **Post** /v0/organizations | Create an organization
*OrganizationsApi* | [**DeleteOrganization**](docs/OrganizationsApi.md#deleteorganization) | **Delete** /v0/organizations/{org_id} | Delete an org
*OrganizationsApi* | [**GetAWSMarketplaceMetadata**](docs/OrganizationsApi.md#getawsmarketplacemetadata) | **Get** /v0/organizations/{org_id}/marketplace-metadata/ | Get marketplace metadata for an org if any exists.
*OrganizationsApi* | [**GetOrganization**](docs/OrganizationsApi.md#getorganization) | **Get** /v0/organizations/{org_id} | Get the metadata about an organization.
*OrganizationsApi* | [**ListOrganizations**](docs/OrganizationsApi.md#listorganizations) | **Get** /v0/organizations | Get a list of all of the organization ids.
*OrganizationsApi* | [**PartiallyUpdateOrg**](docs/OrganizationsApi.md#partiallyupdateorg) | **Put** /v0/organizations/partial/ | Update some fields of an organization to non-null values
*OrganizationsApi* | [**PartiallyUpdateOrganization**](docs/OrganizationsApi.md#partiallyupdateorganization) | **Put** /v0/organizations/partial/{org_id} | Update some fields of an organization to non-null values
*OrganizationsApi* | [**UpdateOrg**](docs/OrganizationsApi.md#updateorg) | **Put** /v0/organizations | Update an existing organization
*OrganizationsApi* | [**UpdateOrganization**](docs/OrganizationsApi.md#updateorganization) | **Put** /v0/organizations/{org_id} | Update an existing organization
*ProvisionApi* | [**ProvisionAWSMarketplaceNewUser**](docs/ProvisionApi.md#provisionawsmarketplacenewuser) | **Post** /v0/provision/marketplace/aws/new-user | Create resources for a new user coming from AWS Marketplace
*ProvisionApi* | [**ProvisionDatabricksConnection**](docs/ProvisionApi.md#provisiondatabricksconnection) | **Post** /v0/provision/connect/databricks | Create resources for a new user coming from Databricks
*ProvisionApi* | [**ProvisionNewUser**](docs/ProvisionApi.md#provisionnewuser) | **Post** /v0/provision/new-user | Create the resources that a new user needs to use WhyLabs via the website.
*ProvisionApi* | [**RegisterDatabricksConnection**](docs/ProvisionApi.md#registerdatabricksconnection) | **Post** /v0/provision/connect/databricks/staged | Register databricks metadata, temporarily storing it against a UUID so that it can be used to provision a databricks connection after email authentication
*SchemaApi* | [**GetMonitorConfigSchema**](docs/SchemaApi.md#getmonitorconfigschema) | **Get** /v0/organizations/{org_id}/schema/monitor-config | Get the current supported schema of the monitor configuration
*SessionsApi* | [**CloseSession**](docs/SessionsApi.md#closesession) | **Post** /v0/sessions/{session_token}/close | naddeo Close a session, triggering its display in whylabs and making it no longer accept any additional data.
*SessionsApi* | [**CreateDatasetProfileUpload**](docs/SessionsApi.md#createdatasetprofileupload) | **Post** /v0/sessions/{session_token}/upload | Create an upload for a given session.
*SessionsApi* | [**CreateSession**](docs/SessionsApi.md#createsession) | **Post** /v0/sessions | Create a new session that can be used to upload dataset profiles from whylogs for display in whylabs.
*SessionsApi* | [**GetSession**](docs/SessionsApi.md#getsession) | **Get** /v0/sessions/{session_token} | Get information about a session.
*UserApi* | [**CreateUser**](docs/UserApi.md#createuser) | **Post** /v0/user | Create a user.
*UserApi* | [**GetUser**](docs/UserApi.md#getuser) | **Get** /v0/user/{user_id} | Get a user by their id.
*UserApi* | [**GetUserByEmail**](docs/UserApi.md#getuserbyemail) | **Get** /v0/user | Get a user by their email.
*UserApi* | [**UpdateUser**](docs/UserApi.md#updateuser) | **Put** /v0/user | Update a user.


## Documentation For Models

 - [AWSMarketplaceMetadata](docs/AWSMarketplaceMetadata.md)
 - [ActionType](docs/ActionType.md)
 - [AddMembershipRequest](docs/AddMembershipRequest.md)
 - [AlertsPath](docs/AlertsPath.md)
 - [AsyncLogResponse](docs/AsyncLogResponse.md)
 - [CloseSessionResponse](docs/CloseSessionResponse.md)
 - [ColumnSchema](docs/ColumnSchema.md)
 - [CreateSessionResponse](docs/CreateSessionResponse.md)
 - [CreateSessionUploadResponse](docs/CreateSessionUploadResponse.md)
 - [CreateUserRequest](docs/CreateUserRequest.md)
 - [DTOAutoScaleDTO](docs/DTOAutoScaleDTO.md)
 - [DTOAwsAttributesDTO](docs/DTOAwsAttributesDTO.md)
 - [DTOAwsAvailabilityDTO](docs/DTOAwsAvailabilityDTO.md)
 - [DTOClusterLogConfDTO](docs/DTOClusterLogConfDTO.md)
 - [DTOCronScheduleDTO](docs/DTOCronScheduleDTO.md)
 - [DTODbfsStorageInfoDTO](docs/DTODbfsStorageInfoDTO.md)
 - [DTOEbsVolumeTypeDTO](docs/DTOEbsVolumeTypeDTO.md)
 - [DTOJobDTO](docs/DTOJobDTO.md)
 - [DTOJobEmailNotificationsDTO](docs/DTOJobEmailNotificationsDTO.md)
 - [DTOJobSettingsDTO](docs/DTOJobSettingsDTO.md)
 - [DTONewClusterDTO](docs/DTONewClusterDTO.md)
 - [DTONotebookTaskDTO](docs/DTONotebookTaskDTO.md)
 - [DTOS3StorageInfoDTO](docs/DTOS3StorageInfoDTO.md)
 - [DTOSparkJarTaskDTO](docs/DTOSparkJarTaskDTO.md)
 - [DTOSparkPythonTaskDTO](docs/DTOSparkPythonTaskDTO.md)
 - [DTOSparkSubmitTaskDTO](docs/DTOSparkSubmitTaskDTO.md)
 - [DatabricksConnection](docs/DatabricksConnection.md)
 - [DatasetRequestMonitorConfig](docs/DatasetRequestMonitorConfig.md)
 - [DatatypeMonitorRequestConfig](docs/DatatypeMonitorRequestConfig.md)
 - [DeleteAnalyzerResultsResponse](docs/DeleteAnalyzerResultsResponse.md)
 - [DeleteDatasetProfilesResponse](docs/DeleteDatasetProfilesResponse.md)
 - [DistributionMonitorRequestConfig](docs/DistributionMonitorRequestConfig.md)
 - [EmailNotificationAction](docs/EmailNotificationAction.md)
 - [EntitySchema](docs/EntitySchema.md)
 - [EntityWeightRecord](docs/EntityWeightRecord.md)
 - [EntityWeightRecordMetadata](docs/EntityWeightRecordMetadata.md)
 - [EventsPath](docs/EventsPath.md)
 - [FeatureFlags](docs/FeatureFlags.md)
 - [GetAlertsPathsResponse](docs/GetAlertsPathsResponse.md)
 - [GetConnectionRequest](docs/GetConnectionRequest.md)
 - [GetConnectionResponse](docs/GetConnectionResponse.md)
 - [GetDefaultMembershipResponse](docs/GetDefaultMembershipResponse.md)
 - [GetEventsPathResponse](docs/GetEventsPathResponse.md)
 - [GetMarketplaceMetadataResponse](docs/GetMarketplaceMetadataResponse.md)
 - [GetMembershipsResponse](docs/GetMembershipsResponse.md)
 - [GetMonitorConfigV2Response](docs/GetMonitorConfigV2Response.md)
 - [GetNotificationSettingsResponse](docs/GetNotificationSettingsResponse.md)
 - [GetSessionResponse](docs/GetSessionResponse.md)
 - [ListJobsRequest](docs/ListJobsRequest.md)
 - [ListJobsResponse](docs/ListJobsResponse.md)
 - [ListModelsResponse](docs/ListModelsResponse.md)
 - [ListOrganizationMembershipsResponse](docs/ListOrganizationMembershipsResponse.md)
 - [ListOrganizationsResponse](docs/ListOrganizationsResponse.md)
 - [ListSegmentsResponse](docs/ListSegmentsResponse.md)
 - [ListUserApiKeys](docs/ListUserApiKeys.md)
 - [LogAsyncRequest](docs/LogAsyncRequest.md)
 - [LogReferenceRequest](docs/LogReferenceRequest.md)
 - [LogReferenceResponse](docs/LogReferenceResponse.md)
 - [LogResponse](docs/LogResponse.md)
 - [MarketplaceDimensions](docs/MarketplaceDimensions.md)
 - [Membership](docs/Membership.md)
 - [MembershipMetadata](docs/MembershipMetadata.md)
 - [MissingRecentDataRequestConfig](docs/MissingRecentDataRequestConfig.md)
 - [MissingRecentProfilesRequestConfig](docs/MissingRecentProfilesRequestConfig.md)
 - [MissingValuesMonitorRequestConfig](docs/MissingValuesMonitorRequestConfig.md)
 - [ModelMetadata](docs/ModelMetadata.md)
 - [ModelType](docs/ModelType.md)
 - [MonitorConfig](docs/MonitorConfig.md)
 - [MonitorConfigVersion](docs/MonitorConfigVersion.md)
 - [MonitorRequestReference](docs/MonitorRequestReference.md)
 - [MonitorRequestReferenceType](docs/MonitorRequestReferenceType.md)
 - [NotificationAction](docs/NotificationAction.md)
 - [NotificationSettings](docs/NotificationSettings.md)
 - [NotificationSettingsDay](docs/NotificationSettingsDay.md)
 - [NotificationSqsMessageCadence](docs/NotificationSqsMessageCadence.md)
 - [OrganizationMetadata](docs/OrganizationMetadata.md)
 - [OrganizationSummary](docs/OrganizationSummary.md)
 - [PagerDutyNotificationAction](docs/PagerDutyNotificationAction.md)
 - [ProvidedConfig](docs/ProvidedConfig.md)
 - [ProvisionDatabricksConnectionRequest](docs/ProvisionDatabricksConnectionRequest.md)
 - [ProvisionDatabricksConnectionResponse](docs/ProvisionDatabricksConnectionResponse.md)
 - [ProvisionNewAWSMarketplaceUserResponse](docs/ProvisionNewAWSMarketplaceUserResponse.md)
 - [ProvisionNewMarketplaceUserRequest](docs/ProvisionNewMarketplaceUserRequest.md)
 - [ProvisionNewUserRequest](docs/ProvisionNewUserRequest.md)
 - [ProvisionNewUserResponse](docs/ProvisionNewUserResponse.md)
 - [PutRequestMonitorRunConfigRequest](docs/PutRequestMonitorRunConfigRequest.md)
 - [ReferenceProfileItemResponse](docs/ReferenceProfileItemResponse.md)
 - [RefreshAccessTokenRequest](docs/RefreshAccessTokenRequest.md)
 - [RefreshConnectionByOrgIdResponse](docs/RefreshConnectionByOrgIdResponse.md)
 - [RefreshConnectionRequest](docs/RefreshConnectionRequest.md)
 - [RegisterDatabricksConnectionRequest](docs/RegisterDatabricksConnectionRequest.md)
 - [RegisterDatabricksConnectionResponse](docs/RegisterDatabricksConnectionResponse.md)
 - [RemoveMembershipRequest](docs/RemoveMembershipRequest.md)
 - [RequestFeatureMonitorConfig](docs/RequestFeatureMonitorConfig.md)
 - [RequestMonitorConfig](docs/RequestMonitorConfig.md)
 - [Role](docs/Role.md)
 - [RunJobRequest](docs/RunJobRequest.md)
 - [RunJobResponse](docs/RunJobResponse.md)
 - [SchemaMetadata](docs/SchemaMetadata.md)
 - [SeasonalARIMARequestConfig](docs/SeasonalARIMARequestConfig.md)
 - [Segment](docs/Segment.md)
 - [SegmentMetadata](docs/SegmentMetadata.md)
 - [SegmentSummary](docs/SegmentSummary.md)
 - [SegmentTag](docs/SegmentTag.md)
 - [SegmentWeight](docs/SegmentWeight.md)
 - [SessionMetadata](docs/SessionMetadata.md)
 - [SetDefaultMembershipRequest](docs/SetDefaultMembershipRequest.md)
 - [SlackNotificationAction](docs/SlackNotificationAction.md)
 - [SubscriptionTier](docs/SubscriptionTier.md)
 - [TimePeriod](docs/TimePeriod.md)
 - [UberNotificationSchedule](docs/UberNotificationSchedule.md)
 - [UniqueValuesMonitorRequestConfig](docs/UniqueValuesMonitorRequestConfig.md)
 - [UpdateConnectionChanges](docs/UpdateConnectionChanges.md)
 - [UpdateConnectionRequest](docs/UpdateConnectionRequest.md)
 - [UpdateMembershipRequest](docs/UpdateMembershipRequest.md)
 - [UpdateOrgRequest](docs/UpdateOrgRequest.md)
 - [User](docs/User.md)
 - [UserApiKey](docs/UserApiKey.md)
 - [UserApiKeyResponse](docs/UserApiKeyResponse.md)
 - [WhyLogsMetric](docs/WhyLogsMetric.md)


## Documentation For Authorization



### ApiKeyAuth

- **Type**: API key
- **API key parameter name**: X-API-Key
- **Location**: HTTP header

Note, each API key must be added to a map of `map[string]APIKey` where the key is: X-API-Key and passed in as the auth context for each request.


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author

support@whylabs.ai

