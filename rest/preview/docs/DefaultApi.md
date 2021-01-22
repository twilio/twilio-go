# \DefaultApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAssistant**](DefaultApi.md#CreateAssistant) | **Post** /understand/Assistants | 
[**CreateAuthorizationDocument**](DefaultApi.md#CreateAuthorizationDocument) | **Post** /HostedNumbers/AuthorizationDocuments | 
[**CreateCertificate**](DefaultApi.md#CreateCertificate) | **Post** /DeployedDevices/Fleets/{FleetSid}/Certificates | 
[**CreateChannel**](DefaultApi.md#CreateChannel) | **Post** /TrustedComms/BrandedChannels/{BrandedChannelSid}/Channels | 
[**CreateCommand**](DefaultApi.md#CreateCommand) | **Post** /wireless/Commands | 
[**CreateDeployment**](DefaultApi.md#CreateDeployment) | **Post** /DeployedDevices/Fleets/{FleetSid}/Deployments | 
[**CreateDevice**](DefaultApi.md#CreateDevice) | **Post** /DeployedDevices/Fleets/{FleetSid}/Devices | 
[**CreateDocument**](DefaultApi.md#CreateDocument) | **Post** /Sync/Services/{ServiceSid}/Documents | 
[**CreateExportCustomJob**](DefaultApi.md#CreateExportCustomJob) | **Post** /BulkExports/Exports/{ResourceType}/Jobs | 
[**CreateField**](DefaultApi.md#CreateField) | **Post** /understand/Assistants/{AssistantSid}/Tasks/{TaskSid}/Fields | 
[**CreateFieldType**](DefaultApi.md#CreateFieldType) | **Post** /understand/Assistants/{AssistantSid}/FieldTypes | 
[**CreateFieldValue**](DefaultApi.md#CreateFieldValue) | **Post** /understand/Assistants/{AssistantSid}/FieldTypes/{FieldTypeSid}/FieldValues | 
[**CreateFleet**](DefaultApi.md#CreateFleet) | **Post** /DeployedDevices/Fleets | 
[**CreateHostedNumberOrder**](DefaultApi.md#CreateHostedNumberOrder) | **Post** /HostedNumbers/HostedNumberOrders | 
[**CreateInstalledAddOn**](DefaultApi.md#CreateInstalledAddOn) | **Post** /marketplace/InstalledAddOns | 
[**CreateKey**](DefaultApi.md#CreateKey) | **Post** /DeployedDevices/Fleets/{FleetSid}/Keys | 
[**CreateModelBuild**](DefaultApi.md#CreateModelBuild) | **Post** /understand/Assistants/{AssistantSid}/ModelBuilds | 
[**CreateQuery**](DefaultApi.md#CreateQuery) | **Post** /understand/Assistants/{AssistantSid}/Queries | 
[**CreateRatePlan**](DefaultApi.md#CreateRatePlan) | **Post** /wireless/RatePlans | 
[**CreateSample**](DefaultApi.md#CreateSample) | **Post** /understand/Assistants/{AssistantSid}/Tasks/{TaskSid}/Samples | 
[**CreateService**](DefaultApi.md#CreateService) | **Post** /Sync/Services | 
[**CreateSyncList**](DefaultApi.md#CreateSyncList) | **Post** /Sync/Services/{ServiceSid}/Lists | 
[**CreateSyncListItem**](DefaultApi.md#CreateSyncListItem) | **Post** /Sync/Services/{ServiceSid}/Lists/{ListSid}/Items | 
[**CreateSyncMap**](DefaultApi.md#CreateSyncMap) | **Post** /Sync/Services/{ServiceSid}/Maps | 
[**CreateSyncMapItem**](DefaultApi.md#CreateSyncMapItem) | **Post** /Sync/Services/{ServiceSid}/Maps/{MapSid}/Items | 
[**CreateTask**](DefaultApi.md#CreateTask) | **Post** /understand/Assistants/{AssistantSid}/Tasks | 
[**DeleteAssistant**](DefaultApi.md#DeleteAssistant) | **Delete** /understand/Assistants/{Sid} | 
[**DeleteCertificate**](DefaultApi.md#DeleteCertificate) | **Delete** /DeployedDevices/Fleets/{FleetSid}/Certificates/{Sid} | 
[**DeleteDeployment**](DefaultApi.md#DeleteDeployment) | **Delete** /DeployedDevices/Fleets/{FleetSid}/Deployments/{Sid} | 
[**DeleteDevice**](DefaultApi.md#DeleteDevice) | **Delete** /DeployedDevices/Fleets/{FleetSid}/Devices/{Sid} | 
[**DeleteDocument**](DefaultApi.md#DeleteDocument) | **Delete** /Sync/Services/{ServiceSid}/Documents/{Sid} | 
[**DeleteDocumentPermission**](DefaultApi.md#DeleteDocumentPermission) | **Delete** /Sync/Services/{ServiceSid}/Documents/{DocumentSid}/Permissions/{Identity} | 
[**DeleteField**](DefaultApi.md#DeleteField) | **Delete** /understand/Assistants/{AssistantSid}/Tasks/{TaskSid}/Fields/{Sid} | 
[**DeleteFieldType**](DefaultApi.md#DeleteFieldType) | **Delete** /understand/Assistants/{AssistantSid}/FieldTypes/{Sid} | 
[**DeleteFieldValue**](DefaultApi.md#DeleteFieldValue) | **Delete** /understand/Assistants/{AssistantSid}/FieldTypes/{FieldTypeSid}/FieldValues/{Sid} | 
[**DeleteFleet**](DefaultApi.md#DeleteFleet) | **Delete** /DeployedDevices/Fleets/{Sid} | 
[**DeleteHostedNumberOrder**](DefaultApi.md#DeleteHostedNumberOrder) | **Delete** /HostedNumbers/HostedNumberOrders/{Sid} | 
[**DeleteInstalledAddOn**](DefaultApi.md#DeleteInstalledAddOn) | **Delete** /marketplace/InstalledAddOns/{Sid} | 
[**DeleteJob**](DefaultApi.md#DeleteJob) | **Delete** /BulkExports/Exports/Jobs/{JobSid} | 
[**DeleteKey**](DefaultApi.md#DeleteKey) | **Delete** /DeployedDevices/Fleets/{FleetSid}/Keys/{Sid} | 
[**DeleteModelBuild**](DefaultApi.md#DeleteModelBuild) | **Delete** /understand/Assistants/{AssistantSid}/ModelBuilds/{Sid} | 
[**DeleteQuery**](DefaultApi.md#DeleteQuery) | **Delete** /understand/Assistants/{AssistantSid}/Queries/{Sid} | 
[**DeleteRatePlan**](DefaultApi.md#DeleteRatePlan) | **Delete** /wireless/RatePlans/{Sid} | 
[**DeleteSample**](DefaultApi.md#DeleteSample) | **Delete** /understand/Assistants/{AssistantSid}/Tasks/{TaskSid}/Samples/{Sid} | 
[**DeleteService**](DefaultApi.md#DeleteService) | **Delete** /Sync/Services/{Sid} | 
[**DeleteSyncList**](DefaultApi.md#DeleteSyncList) | **Delete** /Sync/Services/{ServiceSid}/Lists/{Sid} | 
[**DeleteSyncListItem**](DefaultApi.md#DeleteSyncListItem) | **Delete** /Sync/Services/{ServiceSid}/Lists/{ListSid}/Items/{Index} | 
[**DeleteSyncListPermission**](DefaultApi.md#DeleteSyncListPermission) | **Delete** /Sync/Services/{ServiceSid}/Lists/{ListSid}/Permissions/{Identity} | 
[**DeleteSyncMap**](DefaultApi.md#DeleteSyncMap) | **Delete** /Sync/Services/{ServiceSid}/Maps/{Sid} | 
[**DeleteSyncMapItem**](DefaultApi.md#DeleteSyncMapItem) | **Delete** /Sync/Services/{ServiceSid}/Maps/{MapSid}/Items/{Key} | 
[**DeleteSyncMapPermission**](DefaultApi.md#DeleteSyncMapPermission) | **Delete** /Sync/Services/{ServiceSid}/Maps/{MapSid}/Permissions/{Identity} | 
[**DeleteTask**](DefaultApi.md#DeleteTask) | **Delete** /understand/Assistants/{AssistantSid}/Tasks/{Sid} | 
[**FetchAssistant**](DefaultApi.md#FetchAssistant) | **Get** /understand/Assistants/{Sid} | 
[**FetchAssistantFallbackActions**](DefaultApi.md#FetchAssistantFallbackActions) | **Get** /understand/Assistants/{AssistantSid}/FallbackActions | 
[**FetchAssistantInitiationActions**](DefaultApi.md#FetchAssistantInitiationActions) | **Get** /understand/Assistants/{AssistantSid}/InitiationActions | 
[**FetchAuthorizationDocument**](DefaultApi.md#FetchAuthorizationDocument) | **Get** /HostedNumbers/AuthorizationDocuments/{Sid} | 
[**FetchAvailableAddOn**](DefaultApi.md#FetchAvailableAddOn) | **Get** /marketplace/AvailableAddOns/{Sid} | 
[**FetchAvailableAddOnExtension**](DefaultApi.md#FetchAvailableAddOnExtension) | **Get** /marketplace/AvailableAddOns/{AvailableAddOnSid}/Extensions/{Sid} | 
[**FetchBrandedChannel**](DefaultApi.md#FetchBrandedChannel) | **Get** /TrustedComms/BrandedChannels/{Sid} | 
[**FetchBrandsInformation**](DefaultApi.md#FetchBrandsInformation) | **Get** /TrustedComms/BrandsInformation | 
[**FetchCertificate**](DefaultApi.md#FetchCertificate) | **Get** /DeployedDevices/Fleets/{FleetSid}/Certificates/{Sid} | 
[**FetchCommand**](DefaultApi.md#FetchCommand) | **Get** /wireless/Commands/{Sid} | 
[**FetchCps**](DefaultApi.md#FetchCps) | **Get** /TrustedComms/CPS | 
[**FetchCurrentCall**](DefaultApi.md#FetchCurrentCall) | **Get** /TrustedComms/CurrentCall | 
[**FetchDay**](DefaultApi.md#FetchDay) | **Get** /BulkExports/Exports/{ResourceType}/Days/{Day} | 
[**FetchDeployment**](DefaultApi.md#FetchDeployment) | **Get** /DeployedDevices/Fleets/{FleetSid}/Deployments/{Sid} | 
[**FetchDevice**](DefaultApi.md#FetchDevice) | **Get** /DeployedDevices/Fleets/{FleetSid}/Devices/{Sid} | 
[**FetchDialogue**](DefaultApi.md#FetchDialogue) | **Get** /understand/Assistants/{AssistantSid}/Dialogues/{Sid} | 
[**FetchDocument**](DefaultApi.md#FetchDocument) | **Get** /Sync/Services/{ServiceSid}/Documents/{Sid} | 
[**FetchDocumentPermission**](DefaultApi.md#FetchDocumentPermission) | **Get** /Sync/Services/{ServiceSid}/Documents/{DocumentSid}/Permissions/{Identity} | 
[**FetchExport**](DefaultApi.md#FetchExport) | **Get** /BulkExports/Exports/{ResourceType} | 
[**FetchExportConfiguration**](DefaultApi.md#FetchExportConfiguration) | **Get** /BulkExports/Exports/{ResourceType}/Configuration | 
[**FetchField**](DefaultApi.md#FetchField) | **Get** /understand/Assistants/{AssistantSid}/Tasks/{TaskSid}/Fields/{Sid} | 
[**FetchFieldType**](DefaultApi.md#FetchFieldType) | **Get** /understand/Assistants/{AssistantSid}/FieldTypes/{Sid} | 
[**FetchFieldValue**](DefaultApi.md#FetchFieldValue) | **Get** /understand/Assistants/{AssistantSid}/FieldTypes/{FieldTypeSid}/FieldValues/{Sid} | 
[**FetchFleet**](DefaultApi.md#FetchFleet) | **Get** /DeployedDevices/Fleets/{Sid} | 
[**FetchHostedNumberOrder**](DefaultApi.md#FetchHostedNumberOrder) | **Get** /HostedNumbers/HostedNumberOrders/{Sid} | 
[**FetchInstalledAddOn**](DefaultApi.md#FetchInstalledAddOn) | **Get** /marketplace/InstalledAddOns/{Sid} | 
[**FetchInstalledAddOnExtension**](DefaultApi.md#FetchInstalledAddOnExtension) | **Get** /marketplace/InstalledAddOns/{InstalledAddOnSid}/Extensions/{Sid} | 
[**FetchJob**](DefaultApi.md#FetchJob) | **Get** /BulkExports/Exports/Jobs/{JobSid} | 
[**FetchKey**](DefaultApi.md#FetchKey) | **Get** /DeployedDevices/Fleets/{FleetSid}/Keys/{Sid} | 
[**FetchModelBuild**](DefaultApi.md#FetchModelBuild) | **Get** /understand/Assistants/{AssistantSid}/ModelBuilds/{Sid} | 
[**FetchQuery**](DefaultApi.md#FetchQuery) | **Get** /understand/Assistants/{AssistantSid}/Queries/{Sid} | 
[**FetchRatePlan**](DefaultApi.md#FetchRatePlan) | **Get** /wireless/RatePlans/{Sid} | 
[**FetchSample**](DefaultApi.md#FetchSample) | **Get** /understand/Assistants/{AssistantSid}/Tasks/{TaskSid}/Samples/{Sid} | 
[**FetchService**](DefaultApi.md#FetchService) | **Get** /Sync/Services/{Sid} | 
[**FetchSim**](DefaultApi.md#FetchSim) | **Get** /wireless/Sims/{Sid} | 
[**FetchStyleSheet**](DefaultApi.md#FetchStyleSheet) | **Get** /understand/Assistants/{AssistantSid}/StyleSheet | 
[**FetchSyncList**](DefaultApi.md#FetchSyncList) | **Get** /Sync/Services/{ServiceSid}/Lists/{Sid} | 
[**FetchSyncListItem**](DefaultApi.md#FetchSyncListItem) | **Get** /Sync/Services/{ServiceSid}/Lists/{ListSid}/Items/{Index} | 
[**FetchSyncListPermission**](DefaultApi.md#FetchSyncListPermission) | **Get** /Sync/Services/{ServiceSid}/Lists/{ListSid}/Permissions/{Identity} | 
[**FetchSyncMap**](DefaultApi.md#FetchSyncMap) | **Get** /Sync/Services/{ServiceSid}/Maps/{Sid} | 
[**FetchSyncMapItem**](DefaultApi.md#FetchSyncMapItem) | **Get** /Sync/Services/{ServiceSid}/Maps/{MapSid}/Items/{Key} | 
[**FetchSyncMapPermission**](DefaultApi.md#FetchSyncMapPermission) | **Get** /Sync/Services/{ServiceSid}/Maps/{MapSid}/Permissions/{Identity} | 
[**FetchTask**](DefaultApi.md#FetchTask) | **Get** /understand/Assistants/{AssistantSid}/Tasks/{Sid} | 
[**FetchTaskActions**](DefaultApi.md#FetchTaskActions) | **Get** /understand/Assistants/{AssistantSid}/Tasks/{TaskSid}/Actions | 
[**FetchTaskStatistics**](DefaultApi.md#FetchTaskStatistics) | **Get** /understand/Assistants/{AssistantSid}/Tasks/{TaskSid}/Statistics | 
[**FetchUsage**](DefaultApi.md#FetchUsage) | **Get** /wireless/Sims/{SimSid}/Usage | 
[**ListAssistant**](DefaultApi.md#ListAssistant) | **Get** /understand/Assistants | 
[**ListAuthorizationDocument**](DefaultApi.md#ListAuthorizationDocument) | **Get** /HostedNumbers/AuthorizationDocuments | 
[**ListAvailableAddOn**](DefaultApi.md#ListAvailableAddOn) | **Get** /marketplace/AvailableAddOns | 
[**ListAvailableAddOnExtension**](DefaultApi.md#ListAvailableAddOnExtension) | **Get** /marketplace/AvailableAddOns/{AvailableAddOnSid}/Extensions | 
[**ListCertificate**](DefaultApi.md#ListCertificate) | **Get** /DeployedDevices/Fleets/{FleetSid}/Certificates | 
[**ListCommand**](DefaultApi.md#ListCommand) | **Get** /wireless/Commands | 
[**ListDay**](DefaultApi.md#ListDay) | **Get** /BulkExports/Exports/{ResourceType}/Days | 
[**ListDependentHostedNumberOrder**](DefaultApi.md#ListDependentHostedNumberOrder) | **Get** /HostedNumbers/AuthorizationDocuments/{SigningDocumentSid}/DependentHostedNumberOrders | 
[**ListDeployment**](DefaultApi.md#ListDeployment) | **Get** /DeployedDevices/Fleets/{FleetSid}/Deployments | 
[**ListDevice**](DefaultApi.md#ListDevice) | **Get** /DeployedDevices/Fleets/{FleetSid}/Devices | 
[**ListDocument**](DefaultApi.md#ListDocument) | **Get** /Sync/Services/{ServiceSid}/Documents | 
[**ListDocumentPermission**](DefaultApi.md#ListDocumentPermission) | **Get** /Sync/Services/{ServiceSid}/Documents/{DocumentSid}/Permissions | 
[**ListExportCustomJob**](DefaultApi.md#ListExportCustomJob) | **Get** /BulkExports/Exports/{ResourceType}/Jobs | 
[**ListField**](DefaultApi.md#ListField) | **Get** /understand/Assistants/{AssistantSid}/Tasks/{TaskSid}/Fields | 
[**ListFieldType**](DefaultApi.md#ListFieldType) | **Get** /understand/Assistants/{AssistantSid}/FieldTypes | 
[**ListFieldValue**](DefaultApi.md#ListFieldValue) | **Get** /understand/Assistants/{AssistantSid}/FieldTypes/{FieldTypeSid}/FieldValues | 
[**ListFleet**](DefaultApi.md#ListFleet) | **Get** /DeployedDevices/Fleets | 
[**ListHostedNumberOrder**](DefaultApi.md#ListHostedNumberOrder) | **Get** /HostedNumbers/HostedNumberOrders | 
[**ListInstalledAddOn**](DefaultApi.md#ListInstalledAddOn) | **Get** /marketplace/InstalledAddOns | 
[**ListInstalledAddOnExtension**](DefaultApi.md#ListInstalledAddOnExtension) | **Get** /marketplace/InstalledAddOns/{InstalledAddOnSid}/Extensions | 
[**ListKey**](DefaultApi.md#ListKey) | **Get** /DeployedDevices/Fleets/{FleetSid}/Keys | 
[**ListModelBuild**](DefaultApi.md#ListModelBuild) | **Get** /understand/Assistants/{AssistantSid}/ModelBuilds | 
[**ListQuery**](DefaultApi.md#ListQuery) | **Get** /understand/Assistants/{AssistantSid}/Queries | 
[**ListRatePlan**](DefaultApi.md#ListRatePlan) | **Get** /wireless/RatePlans | 
[**ListSample**](DefaultApi.md#ListSample) | **Get** /understand/Assistants/{AssistantSid}/Tasks/{TaskSid}/Samples | 
[**ListService**](DefaultApi.md#ListService) | **Get** /Sync/Services | 
[**ListSim**](DefaultApi.md#ListSim) | **Get** /wireless/Sims | 
[**ListSyncList**](DefaultApi.md#ListSyncList) | **Get** /Sync/Services/{ServiceSid}/Lists | 
[**ListSyncListItem**](DefaultApi.md#ListSyncListItem) | **Get** /Sync/Services/{ServiceSid}/Lists/{ListSid}/Items | 
[**ListSyncListPermission**](DefaultApi.md#ListSyncListPermission) | **Get** /Sync/Services/{ServiceSid}/Lists/{ListSid}/Permissions | 
[**ListSyncMap**](DefaultApi.md#ListSyncMap) | **Get** /Sync/Services/{ServiceSid}/Maps | 
[**ListSyncMapItem**](DefaultApi.md#ListSyncMapItem) | **Get** /Sync/Services/{ServiceSid}/Maps/{MapSid}/Items | 
[**ListSyncMapPermission**](DefaultApi.md#ListSyncMapPermission) | **Get** /Sync/Services/{ServiceSid}/Maps/{MapSid}/Permissions | 
[**ListTask**](DefaultApi.md#ListTask) | **Get** /understand/Assistants/{AssistantSid}/Tasks | 
[**UpdateAssistant**](DefaultApi.md#UpdateAssistant) | **Post** /understand/Assistants/{Sid} | 
[**UpdateAssistantFallbackActions**](DefaultApi.md#UpdateAssistantFallbackActions) | **Post** /understand/Assistants/{AssistantSid}/FallbackActions | 
[**UpdateAssistantInitiationActions**](DefaultApi.md#UpdateAssistantInitiationActions) | **Post** /understand/Assistants/{AssistantSid}/InitiationActions | 
[**UpdateAuthorizationDocument**](DefaultApi.md#UpdateAuthorizationDocument) | **Post** /HostedNumbers/AuthorizationDocuments/{Sid} | 
[**UpdateCertificate**](DefaultApi.md#UpdateCertificate) | **Post** /DeployedDevices/Fleets/{FleetSid}/Certificates/{Sid} | 
[**UpdateDeployment**](DefaultApi.md#UpdateDeployment) | **Post** /DeployedDevices/Fleets/{FleetSid}/Deployments/{Sid} | 
[**UpdateDevice**](DefaultApi.md#UpdateDevice) | **Post** /DeployedDevices/Fleets/{FleetSid}/Devices/{Sid} | 
[**UpdateDocument**](DefaultApi.md#UpdateDocument) | **Post** /Sync/Services/{ServiceSid}/Documents/{Sid} | 
[**UpdateDocumentPermission**](DefaultApi.md#UpdateDocumentPermission) | **Post** /Sync/Services/{ServiceSid}/Documents/{DocumentSid}/Permissions/{Identity} | 
[**UpdateExportConfiguration**](DefaultApi.md#UpdateExportConfiguration) | **Post** /BulkExports/Exports/{ResourceType}/Configuration | 
[**UpdateFieldType**](DefaultApi.md#UpdateFieldType) | **Post** /understand/Assistants/{AssistantSid}/FieldTypes/{Sid} | 
[**UpdateFleet**](DefaultApi.md#UpdateFleet) | **Post** /DeployedDevices/Fleets/{Sid} | 
[**UpdateHostedNumberOrder**](DefaultApi.md#UpdateHostedNumberOrder) | **Post** /HostedNumbers/HostedNumberOrders/{Sid} | 
[**UpdateInstalledAddOn**](DefaultApi.md#UpdateInstalledAddOn) | **Post** /marketplace/InstalledAddOns/{Sid} | 
[**UpdateInstalledAddOnExtension**](DefaultApi.md#UpdateInstalledAddOnExtension) | **Post** /marketplace/InstalledAddOns/{InstalledAddOnSid}/Extensions/{Sid} | 
[**UpdateKey**](DefaultApi.md#UpdateKey) | **Post** /DeployedDevices/Fleets/{FleetSid}/Keys/{Sid} | 
[**UpdateModelBuild**](DefaultApi.md#UpdateModelBuild) | **Post** /understand/Assistants/{AssistantSid}/ModelBuilds/{Sid} | 
[**UpdateQuery**](DefaultApi.md#UpdateQuery) | **Post** /understand/Assistants/{AssistantSid}/Queries/{Sid} | 
[**UpdateRatePlan**](DefaultApi.md#UpdateRatePlan) | **Post** /wireless/RatePlans/{Sid} | 
[**UpdateSample**](DefaultApi.md#UpdateSample) | **Post** /understand/Assistants/{AssistantSid}/Tasks/{TaskSid}/Samples/{Sid} | 
[**UpdateService**](DefaultApi.md#UpdateService) | **Post** /Sync/Services/{Sid} | 
[**UpdateSim**](DefaultApi.md#UpdateSim) | **Post** /wireless/Sims/{Sid} | 
[**UpdateStyleSheet**](DefaultApi.md#UpdateStyleSheet) | **Post** /understand/Assistants/{AssistantSid}/StyleSheet | 
[**UpdateSyncListItem**](DefaultApi.md#UpdateSyncListItem) | **Post** /Sync/Services/{ServiceSid}/Lists/{ListSid}/Items/{Index} | 
[**UpdateSyncListPermission**](DefaultApi.md#UpdateSyncListPermission) | **Post** /Sync/Services/{ServiceSid}/Lists/{ListSid}/Permissions/{Identity} | 
[**UpdateSyncMapItem**](DefaultApi.md#UpdateSyncMapItem) | **Post** /Sync/Services/{ServiceSid}/Maps/{MapSid}/Items/{Key} | 
[**UpdateSyncMapPermission**](DefaultApi.md#UpdateSyncMapPermission) | **Post** /Sync/Services/{ServiceSid}/Maps/{MapSid}/Permissions/{Identity} | 
[**UpdateTask**](DefaultApi.md#UpdateTask) | **Post** /understand/Assistants/{AssistantSid}/Tasks/{Sid} | 
[**UpdateTaskActions**](DefaultApi.md#UpdateTaskActions) | **Post** /understand/Assistants/{AssistantSid}/Tasks/{TaskSid}/Actions | 



## CreateAssistant

> PreviewUnderstandAssistant CreateAssistant(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateAssistantOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateAssistantOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **CallbackEvents** | **optional.String**| Space-separated list of callback events that will trigger callbacks. | 
 **CallbackUrl** | **optional.String**| A user-provided URL to send event callbacks to. | 
 **FallbackActions** | [**optional.Interface of map[string]interface{}**](map[string]interface{}.md)| The JSON actions to be executed when the user&#39;s input is not recognized as matching any Task. | 
 **FriendlyName** | **optional.String**| A text description for the Assistant. It is non-unique and can up to 255 characters long. | 
 **InitiationActions** | [**optional.Interface of map[string]interface{}**](map[string]interface{}.md)| The JSON actions to be executed on inbound phone calls when the Assistant has to say something first. | 
 **LogQueries** | **optional.Bool**| A boolean that specifies whether queries should be logged for 30 days further training. If false, no queries will be stored, if true, queries will be stored for 30 days and deleted thereafter. Defaults to true if no value is provided. | 
 **StyleSheet** | [**optional.Interface of map[string]interface{}**](map[string]interface{}.md)| The JSON object that holds the style sheet for the assistant | 
 **UniqueName** | **optional.String**| A user-provided string that uniquely identifies this resource as an alternative to the sid. Unique up to 64 characters long. | 

### Return type

[**PreviewUnderstandAssistant**](preview.understand.assistant.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAuthorizationDocument

> PreviewHostedNumbersAuthorizationDocument CreateAuthorizationDocument(ctx, optional)



Create an AuthorizationDocument for authorizing the hosting of phone number capabilities on Twilio's platform.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateAuthorizationDocumentOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateAuthorizationDocumentOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **AddressSid** | **optional.String**| A 34 character string that uniquely identifies the Address resource that is associated with this AuthorizationDocument. | 
 **CcEmails** | [**optional.Interface of []string**](string.md)| Email recipients who will be informed when an Authorization Document has been sent and signed. | 
 **ContactPhoneNumber** | **optional.String**| The contact phone number of the person authorized to sign the Authorization Document. | 
 **ContactTitle** | **optional.String**| The title of the person authorized to sign the Authorization Document for this phone number. | 
 **Email** | **optional.String**| Email that this AuthorizationDocument will be sent to for signing. | 
 **HostedNumberOrderSids** | [**optional.Interface of []string**](string.md)| A list of HostedNumberOrder sids that this AuthorizationDocument will authorize for hosting phone number capabilities on Twilio&#39;s platform. | 

### Return type

[**PreviewHostedNumbersAuthorizationDocument**](preview.hosted_numbers.authorization_document.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCertificate

> PreviewDeployedDevicesFleetCertificate CreateCertificate(ctx, FleetSid, optional)



Enroll a new Certificate credential to the Fleet, optionally giving it a friendly name and assigning to a Device.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**FleetSid** | **string**|  | 
 **optional** | ***CreateCertificateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateCertificateOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **CertificateData** | **optional.String**| Provides a URL encoded representation of the public certificate in PEM format. | 
 **DeviceSid** | **optional.String**| Provides the unique string identifier of an existing Device to become authenticated with this Certificate credential. | 
 **FriendlyName** | **optional.String**| Provides a human readable descriptive text for this Certificate credential, up to 256 characters long. | 

### Return type

[**PreviewDeployedDevicesFleetCertificate**](preview.deployed_devices.fleet.certificate.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateChannel

> PreviewTrustedCommsBrandedChannelChannel CreateChannel(ctx, BrandedChannelSid, optional)



Associate a channel to a branded channel

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**BrandedChannelSid** | **string**| The unique SID identifier of the Branded Channel. The given phone number is going to be assigned to this Branded Channel | 
 **optional** | ***CreateChannelOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateChannelOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **PhoneNumberSid** | **optional.String**| The unique SID identifier of the Phone Number of the Phone number to be assigned to the Branded Channel. | 

### Return type

[**PreviewTrustedCommsBrandedChannelChannel**](preview.trusted_comms.branded_channel.channel.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCommand

> PreviewWirelessCommand CreateCommand(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateCommandOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateCommandOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **CallbackMethod** | **optional.String**|  | 
 **CallbackUrl** | **optional.String**|  | 
 **Command** | **optional.String**|  | 
 **CommandMode** | **optional.String**|  | 
 **Device** | **optional.String**|  | 
 **IncludeSid** | **optional.String**|  | 
 **Sim** | **optional.String**|  | 

### Return type

[**PreviewWirelessCommand**](preview.wireless.command.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateDeployment

> PreviewDeployedDevicesFleetDeployment CreateDeployment(ctx, FleetSid, optional)



Create a new Deployment in the Fleet, optionally giving it a friendly name and linking to a specific Twilio Sync service instance.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**FleetSid** | **string**|  | 
 **optional** | ***CreateDeploymentOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateDeploymentOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **FriendlyName** | **optional.String**| Provides a human readable descriptive text for this Deployment, up to 256 characters long. | 
 **SyncServiceSid** | **optional.String**| Provides the unique string identifier of the Twilio Sync service instance that will be linked to and accessible by this Deployment. | 

### Return type

[**PreviewDeployedDevicesFleetDeployment**](preview.deployed_devices.fleet.deployment.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateDevice

> PreviewDeployedDevicesFleetDevice CreateDevice(ctx, FleetSid, optional)



Create a new Device in the Fleet, optionally giving it a unique name, friendly name, and assigning to a Deployment and/or human identity.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**FleetSid** | **string**|  | 
 **optional** | ***CreateDeviceOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateDeviceOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **DeploymentSid** | **optional.String**| Specifies the unique string identifier of the Deployment group that this Device is going to be associated with. | 
 **Enabled** | **optional.Bool**|  | 
 **FriendlyName** | **optional.String**| Provides a human readable descriptive text to be assigned to this Device, up to 256 characters long. | 
 **Identity** | **optional.String**| Provides an arbitrary string identifier representing a human user to be associated with this Device, up to 256 characters long. | 
 **UniqueName** | **optional.String**| Provides a unique and addressable name to be assigned to this Device, to be used in addition to SID, up to 128 characters long. | 

### Return type

[**PreviewDeployedDevicesFleetDevice**](preview.deployed_devices.fleet.device.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateDocument

> PreviewSyncServiceDocument CreateDocument(ctx, ServiceSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string**|  | 
 **optional** | ***CreateDocumentOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateDocumentOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **Data** | [**optional.Interface of map[string]interface{}**](map[string]interface{}.md)|  | 
 **UniqueName** | **optional.String**|  | 

### Return type

[**PreviewSyncServiceDocument**](preview.sync.service.document.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateExportCustomJob

> PreviewBulkExportsExportExportCustomJob CreateExportCustomJob(ctx, ResourceType, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ResourceType** | **string**| The type of communication – Messages or Calls | 
 **optional** | ***CreateExportCustomJobOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateExportCustomJobOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **Email** | **optional.String**| The optional email to send the completion notification to | 
 **EndDay** | **optional.String**| The end day for the custom export specified as a string in the format of yyyy-mm-dd. End day is inclusive and must be 2 days earlier than the current UTC day. | 
 **FriendlyName** | **optional.String**| The friendly name specified when creating the job | 
 **StartDay** | **optional.String**| The start day for the custom export specified as a string in the format of yyyy-mm-dd | 
 **WebhookMethod** | **optional.String**| This is the method used to call the webhook on completion of the job. If this is supplied, &#x60;WebhookUrl&#x60; must also be supplied. | 
 **WebhookUrl** | **optional.String**| The optional webhook url called on completion of the job. If this is supplied, &#x60;WebhookMethod&#x60; must also be supplied. | 

### Return type

[**PreviewBulkExportsExportExportCustomJob**](preview.bulk_exports.export.export_custom_job.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateField

> PreviewUnderstandAssistantTaskField CreateField(ctx, AssistantSid, TaskSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string**| The unique ID of the parent Assistant. | 
**TaskSid** | **string**| The unique ID of the Task associated with this Field. | 
 **optional** | ***CreateFieldOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateFieldOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **FieldType** | **optional.String**| The unique name or sid of the FieldType. It can be any [Built-in Field Type](https://www.twilio.com/docs/assistant/api/built-in-field-types) or the unique_name or the Field Type sid of a custom Field Type. | 
 **UniqueName** | **optional.String**| A user-provided string that uniquely identifies this resource as an alternative to the sid. Unique up to 64 characters long. | 

### Return type

[**PreviewUnderstandAssistantTaskField**](preview.understand.assistant.task.field.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateFieldType

> PreviewUnderstandAssistantFieldType CreateFieldType(ctx, AssistantSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string**|  | 
 **optional** | ***CreateFieldTypeOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateFieldTypeOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **FriendlyName** | **optional.String**| A user-provided string that identifies this resource. It is non-unique and can up to 255 characters long. | 
 **UniqueName** | **optional.String**| A user-provided string that uniquely identifies this resource as an alternative to the sid. Unique up to 64 characters long. | 

### Return type

[**PreviewUnderstandAssistantFieldType**](preview.understand.assistant.field_type.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateFieldValue

> PreviewUnderstandAssistantFieldTypeFieldValue CreateFieldValue(ctx, AssistantSid, FieldTypeSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string**|  | 
**FieldTypeSid** | **string**|  | 
 **optional** | ***CreateFieldValueOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateFieldValueOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **Language** | **optional.String**| An ISO language-country string of the value. | 
 **SynonymOf** | **optional.String**| A value that indicates this field value is a synonym of. Empty if the value is not a synonym. | 
 **Value** | **optional.String**| A user-provided string that uniquely identifies this resource as an alternative to the sid. Unique up to 64 characters long. | 

### Return type

[**PreviewUnderstandAssistantFieldTypeFieldValue**](preview.understand.assistant.field_type.field_value.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateFleet

> PreviewDeployedDevicesFleet CreateFleet(ctx, optional)



Create a new Fleet for scoping of deployed devices within your account.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateFleetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateFleetOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **FriendlyName** | **optional.String**| Provides a human readable descriptive text for this Fleet, up to 256 characters long. | 

### Return type

[**PreviewDeployedDevicesFleet**](preview.deployed_devices.fleet.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateHostedNumberOrder

> PreviewHostedNumbersHostedNumberOrder CreateHostedNumberOrder(ctx, optional)



Host a phone number's capability on Twilio's platform.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateHostedNumberOrderOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateHostedNumberOrderOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **AccountSid** | **optional.String**| This defaults to the AccountSid of the authorization the user is using. This can be provided to specify a subaccount to add the HostedNumberOrder to. | 
 **AddressSid** | **optional.String**| Optional. A 34 character string that uniquely identifies the Address resource that represents the address of the owner of this phone number. | 
 **CcEmails** | [**optional.Interface of []string**](string.md)| Optional. A list of emails that the LOA document for this HostedNumberOrder will be carbon copied to. | 
 **Email** | **optional.String**| Optional. Email of the owner of this phone number that is being hosted. | 
 **FriendlyName** | **optional.String**| A 64 character string that is a human readable text that describes this resource. | 
 **PhoneNumber** | **optional.String**| The number to host in [+E.164](https://en.wikipedia.org/wiki/E.164) format | 
 **SmsApplicationSid** | **optional.String**| Optional. The 34 character sid of the application Twilio should use to handle SMS messages sent to this number. If a &#x60;SmsApplicationSid&#x60; is present, Twilio will ignore all of the SMS urls above and use those set on the application. | 
 **SmsCapability** | **optional.Bool**| Used to specify that the SMS capability will be hosted on Twilio&#39;s platform. | 
 **SmsFallbackMethod** | **optional.String**| The HTTP method that should be used to request the SmsFallbackUrl. Must be either &#x60;GET&#x60; or &#x60;POST&#x60;. This will be copied onto the IncomingPhoneNumber resource. | 
 **SmsFallbackUrl** | **optional.String**| A URL that Twilio will request if an error occurs requesting or executing the TwiML defined by SmsUrl. This will be copied onto the IncomingPhoneNumber resource. | 
 **SmsMethod** | **optional.String**| The HTTP method that should be used to request the SmsUrl. Must be either &#x60;GET&#x60; or &#x60;POST&#x60;.  This will be copied onto the IncomingPhoneNumber resource. | 
 **SmsUrl** | **optional.String**| The URL that Twilio should request when somebody sends an SMS to the phone number. This will be copied onto the IncomingPhoneNumber resource. | 
 **StatusCallbackMethod** | **optional.String**| Optional. The Status Callback Method attached to the IncomingPhoneNumber resource. | 
 **StatusCallbackUrl** | **optional.String**| Optional. The Status Callback URL attached to the IncomingPhoneNumber resource. | 
 **UniqueName** | **optional.String**| Optional. Provides a unique and addressable name to be assigned to this HostedNumberOrder, assigned by the developer, to be optionally used in addition to SID. | 
 **VerificationDocumentSid** | **optional.String**| Optional. The unique sid identifier of the Identity Document that represents the document for verifying ownership of the number to be hosted. Required when VerificationType is phone-bill. | 
 **VerificationType** | **optional.String**| Optional. The method used for verifying ownership of the number to be hosted. One of phone-call (default) or phone-bill. | 

### Return type

[**PreviewHostedNumbersHostedNumberOrder**](preview.hosted_numbers.hosted_number_order.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateInstalledAddOn

> PreviewMarketplaceInstalledAddOn CreateInstalledAddOn(ctx, optional)



Install an Add-on for the Account specified.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateInstalledAddOnOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateInstalledAddOnOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **AcceptTermsOfService** | **optional.Bool**| Whether the Terms of Service were accepted. | 
 **AvailableAddOnSid** | **optional.String**| The SID of the AvaliableAddOn to install. | 
 **Configuration** | [**optional.Interface of map[string]interface{}**](map[string]interface{}.md)| The JSON object that represents the configuration of the new Add-on being installed. | 
 **UniqueName** | **optional.String**| An application-defined string that uniquely identifies the resource. This value must be unique within the Account. | 

### Return type

[**PreviewMarketplaceInstalledAddOn**](preview.marketplace.installed_add_on.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateKey

> PreviewDeployedDevicesFleetKey CreateKey(ctx, FleetSid, optional)



Create a new Key credential in the Fleet, optionally giving it a friendly name and assigning to a Device.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**FleetSid** | **string**|  | 
 **optional** | ***CreateKeyOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateKeyOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **DeviceSid** | **optional.String**| Provides the unique string identifier of an existing Device to become authenticated with this Key credential. | 
 **FriendlyName** | **optional.String**| Provides a human readable descriptive text for this Key credential, up to 256 characters long. | 

### Return type

[**PreviewDeployedDevicesFleetKey**](preview.deployed_devices.fleet.key.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateModelBuild

> PreviewUnderstandAssistantModelBuild CreateModelBuild(ctx, AssistantSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string**|  | 
 **optional** | ***CreateModelBuildOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateModelBuildOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **StatusCallback** | **optional.String**|  | 
 **UniqueName** | **optional.String**| A user-provided string that uniquely identifies this resource as an alternative to the sid. Unique up to 64 characters long. For example: v0.1 | 

### Return type

[**PreviewUnderstandAssistantModelBuild**](preview.understand.assistant.model_build.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateQuery

> PreviewUnderstandAssistantQuery CreateQuery(ctx, AssistantSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string**| The unique ID of the parent Assistant. | 
 **optional** | ***CreateQueryOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateQueryOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **Field** | **optional.String**| Constraints the query to a given Field with an task. Useful when you know the Field you are expecting. It accepts one field in the format *task-unique-name-1*:*field-unique-name* | 
 **Language** | **optional.String**| An ISO language-country string of the sample. | 
 **ModelBuild** | **optional.String**| The Model Build Sid or unique name of the Model Build to be queried. | 
 **Query** | **optional.String**| A user-provided string that uniquely identifies this resource as an alternative to the sid. It can be up to 2048 characters long. | 
 **Tasks** | **optional.String**| Constraints the query to a set of tasks. Useful when you need to constrain the paths the user can take. Tasks should be comma separated *task-unique-name-1*, *task-unique-name-2* | 

### Return type

[**PreviewUnderstandAssistantQuery**](preview.understand.assistant.query.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateRatePlan

> PreviewWirelessRatePlan CreateRatePlan(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateRatePlanOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateRatePlanOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **CommandsEnabled** | **optional.Bool**|  | 
 **DataEnabled** | **optional.Bool**|  | 
 **DataLimit** | **optional.Int32**|  | 
 **DataMetering** | **optional.String**|  | 
 **FriendlyName** | **optional.String**|  | 
 **InternationalRoaming** | [**optional.Interface of []string**](string.md)|  | 
 **MessagingEnabled** | **optional.Bool**|  | 
 **NationalRoamingEnabled** | **optional.Bool**|  | 
 **UniqueName** | **optional.String**|  | 
 **VoiceEnabled** | **optional.Bool**|  | 

### Return type

[**PreviewWirelessRatePlan**](preview.wireless.rate_plan.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSample

> PreviewUnderstandAssistantTaskSample CreateSample(ctx, AssistantSid, TaskSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string**| The unique ID of the Assistant. | 
**TaskSid** | **string**| The unique ID of the Task associated with this Sample. | 
 **optional** | ***CreateSampleOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateSampleOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **Language** | **optional.String**| An ISO language-country string of the sample. | 
 **SourceChannel** | **optional.String**| The communication channel the sample was captured. It can be: *voice*, *sms*, *chat*, *alexa*, *google-assistant*, or *slack*. If not included the value will be null | 
 **TaggedText** | **optional.String**| The text example of how end-users may express this task. The sample may contain Field tag blocks. | 

### Return type

[**PreviewUnderstandAssistantTaskSample**](preview.understand.assistant.task.sample.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateService

> PreviewSyncService CreateService(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateServiceOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateServiceOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **AclEnabled** | **optional.Bool**|  | 
 **FriendlyName** | **optional.String**|  | 
 **ReachabilityWebhooksEnabled** | **optional.Bool**|  | 
 **WebhookUrl** | **optional.String**|  | 

### Return type

[**PreviewSyncService**](preview.sync.service.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSyncList

> PreviewSyncServiceSyncList CreateSyncList(ctx, ServiceSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string**|  | 
 **optional** | ***CreateSyncListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateSyncListOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **UniqueName** | **optional.String**|  | 

### Return type

[**PreviewSyncServiceSyncList**](preview.sync.service.sync_list.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSyncListItem

> PreviewSyncServiceSyncListSyncListItem CreateSyncListItem(ctx, ServiceSid, ListSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string**|  | 
**ListSid** | **string**|  | 
 **optional** | ***CreateSyncListItemOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateSyncListItemOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **Data** | [**optional.Interface of map[string]interface{}**](map[string]interface{}.md)|  | 

### Return type

[**PreviewSyncServiceSyncListSyncListItem**](preview.sync.service.sync_list.sync_list_item.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSyncMap

> PreviewSyncServiceSyncMap CreateSyncMap(ctx, ServiceSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string**|  | 
 **optional** | ***CreateSyncMapOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateSyncMapOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **UniqueName** | **optional.String**|  | 

### Return type

[**PreviewSyncServiceSyncMap**](preview.sync.service.sync_map.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSyncMapItem

> PreviewSyncServiceSyncMapSyncMapItem CreateSyncMapItem(ctx, ServiceSid, MapSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string**|  | 
**MapSid** | **string**|  | 
 **optional** | ***CreateSyncMapItemOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateSyncMapItemOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **Data** | [**optional.Interface of map[string]interface{}**](map[string]interface{}.md)|  | 
 **Key** | **optional.String**|  | 

### Return type

[**PreviewSyncServiceSyncMapSyncMapItem**](preview.sync.service.sync_map.sync_map_item.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTask

> PreviewUnderstandAssistantTask CreateTask(ctx, AssistantSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string**| The unique ID of the Assistant. | 
 **optional** | ***CreateTaskOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateTaskOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **Actions** | [**optional.Interface of map[string]interface{}**](map[string]interface{}.md)| A user-provided JSON object encoded as a string to specify the actions for this task. It is optional and non-unique. | 
 **ActionsUrl** | **optional.String**| User-provided HTTP endpoint where from the assistant fetches actions | 
 **FriendlyName** | **optional.String**| A user-provided string that identifies this resource. It is non-unique and can up to 255 characters long. | 
 **UniqueName** | **optional.String**| A user-provided string that uniquely identifies this resource as an alternative to the sid. Unique up to 64 characters long. | 

### Return type

[**PreviewUnderstandAssistantTask**](preview.understand.assistant.task.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAssistant

> DeleteAssistant(ctx, Sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string**| A 34 character string that uniquely identifies this resource. | 

### Return type

 (empty response body)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCertificate

> DeleteCertificate(ctx, FleetSid, Sid)



Unregister a specific Certificate credential from the Fleet, effectively disallowing any inbound client connections that are presenting it.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**FleetSid** | **string**|  | 
**Sid** | **string**| Provides a 34 character string that uniquely identifies the requested Certificate credential resource. | 

### Return type

 (empty response body)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDeployment

> DeleteDeployment(ctx, FleetSid, Sid)



Delete a specific Deployment from the Fleet, leaving associated devices effectively undeployed.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**FleetSid** | **string**|  | 
**Sid** | **string**| Provides a 34 character string that uniquely identifies the requested Deployment resource. | 

### Return type

 (empty response body)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDevice

> DeleteDevice(ctx, FleetSid, Sid)



Delete a specific Device from the Fleet, also removing it from associated Deployments.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**FleetSid** | **string**|  | 
**Sid** | **string**| Provides a 34 character string that uniquely identifies the requested Device resource. | 

### Return type

 (empty response body)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDocument

> DeleteDocument(ctx, ServiceSid, Sid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string**|  | 
**Sid** | **string**|  | 
 **optional** | ***DeleteDocumentOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeleteDocumentOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **IfMatch** | **optional.String**| The If-Match HTTP request header | 

### Return type

 (empty response body)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDocumentPermission

> DeleteDocumentPermission(ctx, ServiceSid, DocumentSid, Identity)



Delete a specific Sync Document Permission.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string**|  | 
**DocumentSid** | **string**| Identifier of the Sync Document. Either a SID or a unique name. | 
**Identity** | **string**| Arbitrary string identifier representing a user associated with an FPA token, assigned by the developer. | 

### Return type

 (empty response body)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteField

> DeleteField(ctx, AssistantSid, TaskSid, Sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string**| The unique ID of the Assistant. | 
**TaskSid** | **string**| The unique ID of the Task associated with this Field. | 
**Sid** | **string**| A 34 character string that uniquely identifies this resource. | 

### Return type

 (empty response body)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteFieldType

> DeleteFieldType(ctx, AssistantSid, Sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string**|  | 
**Sid** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteFieldValue

> DeleteFieldValue(ctx, AssistantSid, FieldTypeSid, Sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string**|  | 
**FieldTypeSid** | **string**|  | 
**Sid** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteFleet

> DeleteFleet(ctx, Sid)



Delete a specific Fleet from your account, also destroys all nested resources: Devices, Deployments, Certificates, Keys.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string**| Provides a 34 character string that uniquely identifies the requested Fleet resource. | 

### Return type

 (empty response body)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteHostedNumberOrder

> DeleteHostedNumberOrder(ctx, Sid)



Cancel the HostedNumberOrder (only available when the status is in `received`).

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string**| A 34 character string that uniquely identifies this HostedNumberOrder. | 

### Return type

 (empty response body)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteInstalledAddOn

> DeleteInstalledAddOn(ctx, Sid)



Remove an Add-on installation from your account

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string**| The SID of the InstalledAddOn resource to delete. | 

### Return type

 (empty response body)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteJob

> DeleteJob(ctx, JobSid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**JobSid** | **string**| The unique string that that we created to identify the Bulk Export job | 

### Return type

 (empty response body)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteKey

> DeleteKey(ctx, FleetSid, Sid)



Delete a specific Key credential from the Fleet, effectively disallowing any inbound client connections that are presenting it.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**FleetSid** | **string**|  | 
**Sid** | **string**| Provides a 34 character string that uniquely identifies the requested Key credential resource. | 

### Return type

 (empty response body)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteModelBuild

> DeleteModelBuild(ctx, AssistantSid, Sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string**|  | 
**Sid** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteQuery

> DeleteQuery(ctx, AssistantSid, Sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string**| The unique ID of the Assistant. | 
**Sid** | **string**| A 34 character string that uniquely identifies this resource. | 

### Return type

 (empty response body)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRatePlan

> DeleteRatePlan(ctx, Sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSample

> DeleteSample(ctx, AssistantSid, TaskSid, Sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string**| The unique ID of the Assistant. | 
**TaskSid** | **string**| The unique ID of the Task associated with this Sample. | 
**Sid** | **string**| A 34 character string that uniquely identifies this resource. | 

### Return type

 (empty response body)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteService

> DeleteService(ctx, Sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSyncList

> DeleteSyncList(ctx, ServiceSid, Sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string**|  | 
**Sid** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSyncListItem

> DeleteSyncListItem(ctx, ServiceSid, ListSid, Index, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string**|  | 
**ListSid** | **string**|  | 
**Index** | **int32**|  | 
 **optional** | ***DeleteSyncListItemOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeleteSyncListItemOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **IfMatch** | **optional.String**| The If-Match HTTP request header | 

### Return type

 (empty response body)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSyncListPermission

> DeleteSyncListPermission(ctx, ServiceSid, ListSid, Identity)



Delete a specific Sync List Permission.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string**|  | 
**ListSid** | **string**| Identifier of the Sync List. Either a SID or a unique name. | 
**Identity** | **string**| Arbitrary string identifier representing a user associated with an FPA token, assigned by the developer. | 

### Return type

 (empty response body)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSyncMap

> DeleteSyncMap(ctx, ServiceSid, Sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string**|  | 
**Sid** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSyncMapItem

> DeleteSyncMapItem(ctx, ServiceSid, MapSid, Key, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string**|  | 
**MapSid** | **string**|  | 
**Key** | **string**|  | 
 **optional** | ***DeleteSyncMapItemOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeleteSyncMapItemOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **IfMatch** | **optional.String**| The If-Match HTTP request header | 

### Return type

 (empty response body)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSyncMapPermission

> DeleteSyncMapPermission(ctx, ServiceSid, MapSid, Identity)



Delete a specific Sync Map Permission.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string**|  | 
**MapSid** | **string**| Identifier of the Sync Map. Either a SID or a unique name. | 
**Identity** | **string**| Arbitrary string identifier representing a user associated with an FPA token, assigned by the developer. | 

### Return type

 (empty response body)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTask

> DeleteTask(ctx, AssistantSid, Sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string**| The unique ID of the Assistant. | 
**Sid** | **string**| A 34 character string that uniquely identifies this resource. | 

### Return type

 (empty response body)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchAssistant

> PreviewUnderstandAssistant FetchAssistant(ctx, Sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string**| A 34 character string that uniquely identifies this resource. | 

### Return type

[**PreviewUnderstandAssistant**](preview.understand.assistant.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchAssistantFallbackActions

> PreviewUnderstandAssistantAssistantFallbackActions FetchAssistantFallbackActions(ctx, AssistantSid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string**|  | 

### Return type

[**PreviewUnderstandAssistantAssistantFallbackActions**](preview.understand.assistant.assistant_fallback_actions.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchAssistantInitiationActions

> PreviewUnderstandAssistantAssistantInitiationActions FetchAssistantInitiationActions(ctx, AssistantSid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string**|  | 

### Return type

[**PreviewUnderstandAssistantAssistantInitiationActions**](preview.understand.assistant.assistant_initiation_actions.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchAuthorizationDocument

> PreviewHostedNumbersAuthorizationDocument FetchAuthorizationDocument(ctx, Sid)



Fetch a specific AuthorizationDocument.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string**| A 34 character string that uniquely identifies this AuthorizationDocument. | 

### Return type

[**PreviewHostedNumbersAuthorizationDocument**](preview.hosted_numbers.authorization_document.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchAvailableAddOn

> PreviewMarketplaceAvailableAddOn FetchAvailableAddOn(ctx, Sid)



Fetch an instance of an Add-on currently available to be installed.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string**| The SID of the AvailableAddOn resource to fetch. | 

### Return type

[**PreviewMarketplaceAvailableAddOn**](preview.marketplace.available_add_on.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchAvailableAddOnExtension

> PreviewMarketplaceAvailableAddOnAvailableAddOnExtension FetchAvailableAddOnExtension(ctx, AvailableAddOnSid, Sid)



Fetch an instance of an Extension for the Available Add-on.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AvailableAddOnSid** | **string**| The SID of the AvailableAddOn resource with the extension to fetch. | 
**Sid** | **string**| The SID of the AvailableAddOn Extension resource to fetch. | 

### Return type

[**PreviewMarketplaceAvailableAddOnAvailableAddOnExtension**](preview.marketplace.available_add_on.available_add_on_extension.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchBrandedChannel

> PreviewTrustedCommsBrandedChannel FetchBrandedChannel(ctx, Sid)



Fetch a specific Branded Channel.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string**| The unique SID identifier of the Branded Channel. | 

### Return type

[**PreviewTrustedCommsBrandedChannel**](preview.trusted_comms.branded_channel.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchBrandsInformation

> PreviewTrustedCommsBrandsInformation FetchBrandsInformation(ctx, optional)



Retrieve the newest available BrandInformation

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***FetchBrandsInformationOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a FetchBrandsInformationOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **IfNoneMatch** | **optional.String**| The If-None-Match HTTP request header | 

### Return type

[**PreviewTrustedCommsBrandsInformation**](preview.trusted_comms.brands_information.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchCertificate

> PreviewDeployedDevicesFleetCertificate FetchCertificate(ctx, FleetSid, Sid)



Fetch information about a specific Certificate credential in the Fleet.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**FleetSid** | **string**|  | 
**Sid** | **string**| Provides a 34 character string that uniquely identifies the requested Certificate credential resource. | 

### Return type

[**PreviewDeployedDevicesFleetCertificate**](preview.deployed_devices.fleet.certificate.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchCommand

> PreviewWirelessCommand FetchCommand(ctx, Sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string**|  | 

### Return type

[**PreviewWirelessCommand**](preview.wireless.command.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchCps

> PreviewTrustedCommsCps FetchCps(ctx, optional)



Fetch a specific Call Placement Service (CPS) given a phone number via `X-XCNAM-Sensitive-Phone-Number` header.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***FetchCpsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a FetchCpsOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **XXcnamSensitivePhoneNumber** | **optional.String**| The X-Xcnam-Sensitive-Phone-Number HTTP request header | 

### Return type

[**PreviewTrustedCommsCps**](preview.trusted_comms.cps.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchCurrentCall

> PreviewTrustedCommsCurrentCall FetchCurrentCall(ctx, optional)



Retrieve a current call given the originating and terminating number via `X-XCNAM-Sensitive-Phone-Number-From` and `X-XCNAM-Sensitive-Phone-Number-To` headers.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***FetchCurrentCallOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a FetchCurrentCallOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **XXcnamSensitivePhoneNumberFrom** | **optional.String**| The X-Xcnam-Sensitive-Phone-Number-From HTTP request header | 
 **XXcnamSensitivePhoneNumberTo** | **optional.String**| The X-Xcnam-Sensitive-Phone-Number-To HTTP request header | 

### Return type

[**PreviewTrustedCommsCurrentCall**](preview.trusted_comms.current_call.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchDay

> FetchDay(ctx, ResourceType, Day)



Fetch a specific Day.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ResourceType** | **string**| The type of communication – Messages, Calls | 
**Day** | **string**| The ISO 8601 format date of the resources in the file, for a UTC day | 

### Return type

 (empty response body)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchDeployment

> PreviewDeployedDevicesFleetDeployment FetchDeployment(ctx, FleetSid, Sid)



Fetch information about a specific Deployment in the Fleet.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**FleetSid** | **string**|  | 
**Sid** | **string**| Provides a 34 character string that uniquely identifies the requested Deployment resource. | 

### Return type

[**PreviewDeployedDevicesFleetDeployment**](preview.deployed_devices.fleet.deployment.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchDevice

> PreviewDeployedDevicesFleetDevice FetchDevice(ctx, FleetSid, Sid)



Fetch information about a specific Device in the Fleet.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**FleetSid** | **string**|  | 
**Sid** | **string**| Provides a 34 character string that uniquely identifies the requested Device resource. | 

### Return type

[**PreviewDeployedDevicesFleetDevice**](preview.deployed_devices.fleet.device.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchDialogue

> PreviewUnderstandAssistantDialogue FetchDialogue(ctx, AssistantSid, Sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string**|  | 
**Sid** | **string**|  | 

### Return type

[**PreviewUnderstandAssistantDialogue**](preview.understand.assistant.dialogue.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchDocument

> PreviewSyncServiceDocument FetchDocument(ctx, ServiceSid, Sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string**|  | 
**Sid** | **string**|  | 

### Return type

[**PreviewSyncServiceDocument**](preview.sync.service.document.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchDocumentPermission

> PreviewSyncServiceDocumentDocumentPermission FetchDocumentPermission(ctx, ServiceSid, DocumentSid, Identity)



Fetch a specific Sync Document Permission.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string**|  | 
**DocumentSid** | **string**| Identifier of the Sync Document. Either a SID or a unique name. | 
**Identity** | **string**| Arbitrary string identifier representing a user associated with an FPA token, assigned by the developer. | 

### Return type

[**PreviewSyncServiceDocumentDocumentPermission**](preview.sync.service.document.document_permission.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchExport

> PreviewBulkExportsExport FetchExport(ctx, ResourceType)



Fetch a specific Export.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ResourceType** | **string**| The type of communication – Messages, Calls | 

### Return type

[**PreviewBulkExportsExport**](preview.bulk_exports.export.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchExportConfiguration

> PreviewBulkExportsExportConfiguration FetchExportConfiguration(ctx, ResourceType)



Fetch a specific Export Configuration.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ResourceType** | **string**| The type of communication – Messages, Calls | 

### Return type

[**PreviewBulkExportsExportConfiguration**](preview.bulk_exports.export_configuration.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchField

> PreviewUnderstandAssistantTaskField FetchField(ctx, AssistantSid, TaskSid, Sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string**| The unique ID of the Assistant. | 
**TaskSid** | **string**| The unique ID of the Task associated with this Field. | 
**Sid** | **string**| A 34 character string that uniquely identifies this resource. | 

### Return type

[**PreviewUnderstandAssistantTaskField**](preview.understand.assistant.task.field.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchFieldType

> PreviewUnderstandAssistantFieldType FetchFieldType(ctx, AssistantSid, Sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string**|  | 
**Sid** | **string**|  | 

### Return type

[**PreviewUnderstandAssistantFieldType**](preview.understand.assistant.field_type.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchFieldValue

> PreviewUnderstandAssistantFieldTypeFieldValue FetchFieldValue(ctx, AssistantSid, FieldTypeSid, Sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string**|  | 
**FieldTypeSid** | **string**|  | 
**Sid** | **string**|  | 

### Return type

[**PreviewUnderstandAssistantFieldTypeFieldValue**](preview.understand.assistant.field_type.field_value.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchFleet

> PreviewDeployedDevicesFleet FetchFleet(ctx, Sid)



Fetch information about a specific Fleet in your account.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string**| Provides a 34 character string that uniquely identifies the requested Fleet resource. | 

### Return type

[**PreviewDeployedDevicesFleet**](preview.deployed_devices.fleet.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchHostedNumberOrder

> PreviewHostedNumbersHostedNumberOrder FetchHostedNumberOrder(ctx, Sid)



Fetch a specific HostedNumberOrder.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string**| A 34 character string that uniquely identifies this HostedNumberOrder. | 

### Return type

[**PreviewHostedNumbersHostedNumberOrder**](preview.hosted_numbers.hosted_number_order.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchInstalledAddOn

> PreviewMarketplaceInstalledAddOn FetchInstalledAddOn(ctx, Sid)



Fetch an instance of an Add-on currently installed on this Account.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string**| The SID of the InstalledAddOn resource to fetch. | 

### Return type

[**PreviewMarketplaceInstalledAddOn**](preview.marketplace.installed_add_on.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchInstalledAddOnExtension

> PreviewMarketplaceInstalledAddOnInstalledAddOnExtension FetchInstalledAddOnExtension(ctx, InstalledAddOnSid, Sid)



Fetch an instance of an Extension for the Installed Add-on.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**InstalledAddOnSid** | **string**| The SID of the InstalledAddOn resource with the extension to fetch. | 
**Sid** | **string**| The SID of the InstalledAddOn Extension resource to fetch. | 

### Return type

[**PreviewMarketplaceInstalledAddOnInstalledAddOnExtension**](preview.marketplace.installed_add_on.installed_add_on_extension.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchJob

> PreviewBulkExportsExportJob FetchJob(ctx, JobSid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**JobSid** | **string**|  | 

### Return type

[**PreviewBulkExportsExportJob**](preview.bulk_exports.export.job.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchKey

> PreviewDeployedDevicesFleetKey FetchKey(ctx, FleetSid, Sid)



Fetch information about a specific Key credential in the Fleet.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**FleetSid** | **string**|  | 
**Sid** | **string**| Provides a 34 character string that uniquely identifies the requested Key credential resource. | 

### Return type

[**PreviewDeployedDevicesFleetKey**](preview.deployed_devices.fleet.key.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchModelBuild

> PreviewUnderstandAssistantModelBuild FetchModelBuild(ctx, AssistantSid, Sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string**|  | 
**Sid** | **string**|  | 

### Return type

[**PreviewUnderstandAssistantModelBuild**](preview.understand.assistant.model_build.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchQuery

> PreviewUnderstandAssistantQuery FetchQuery(ctx, AssistantSid, Sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string**| The unique ID of the Assistant. | 
**Sid** | **string**| A 34 character string that uniquely identifies this resource. | 

### Return type

[**PreviewUnderstandAssistantQuery**](preview.understand.assistant.query.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchRatePlan

> PreviewWirelessRatePlan FetchRatePlan(ctx, Sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string**|  | 

### Return type

[**PreviewWirelessRatePlan**](preview.wireless.rate_plan.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchSample

> PreviewUnderstandAssistantTaskSample FetchSample(ctx, AssistantSid, TaskSid, Sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string**| The unique ID of the Assistant. | 
**TaskSid** | **string**| The unique ID of the Task associated with this Sample. | 
**Sid** | **string**| A 34 character string that uniquely identifies this resource. | 

### Return type

[**PreviewUnderstandAssistantTaskSample**](preview.understand.assistant.task.sample.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchService

> PreviewSyncService FetchService(ctx, Sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string**|  | 

### Return type

[**PreviewSyncService**](preview.sync.service.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchSim

> PreviewWirelessSim FetchSim(ctx, Sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string**|  | 

### Return type

[**PreviewWirelessSim**](preview.wireless.sim.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchStyleSheet

> PreviewUnderstandAssistantStyleSheet FetchStyleSheet(ctx, AssistantSid)



Returns Style sheet JSON object for this Assistant

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string**| The unique ID of the Assistant | 

### Return type

[**PreviewUnderstandAssistantStyleSheet**](preview.understand.assistant.style_sheet.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchSyncList

> PreviewSyncServiceSyncList FetchSyncList(ctx, ServiceSid, Sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string**|  | 
**Sid** | **string**|  | 

### Return type

[**PreviewSyncServiceSyncList**](preview.sync.service.sync_list.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchSyncListItem

> PreviewSyncServiceSyncListSyncListItem FetchSyncListItem(ctx, ServiceSid, ListSid, Index)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string**|  | 
**ListSid** | **string**|  | 
**Index** | **int32**|  | 

### Return type

[**PreviewSyncServiceSyncListSyncListItem**](preview.sync.service.sync_list.sync_list_item.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchSyncListPermission

> PreviewSyncServiceSyncListSyncListPermission FetchSyncListPermission(ctx, ServiceSid, ListSid, Identity)



Fetch a specific Sync List Permission.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string**|  | 
**ListSid** | **string**| Identifier of the Sync List. Either a SID or a unique name. | 
**Identity** | **string**| Arbitrary string identifier representing a user associated with an FPA token, assigned by the developer. | 

### Return type

[**PreviewSyncServiceSyncListSyncListPermission**](preview.sync.service.sync_list.sync_list_permission.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchSyncMap

> PreviewSyncServiceSyncMap FetchSyncMap(ctx, ServiceSid, Sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string**|  | 
**Sid** | **string**|  | 

### Return type

[**PreviewSyncServiceSyncMap**](preview.sync.service.sync_map.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchSyncMapItem

> PreviewSyncServiceSyncMapSyncMapItem FetchSyncMapItem(ctx, ServiceSid, MapSid, Key)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string**|  | 
**MapSid** | **string**|  | 
**Key** | **string**|  | 

### Return type

[**PreviewSyncServiceSyncMapSyncMapItem**](preview.sync.service.sync_map.sync_map_item.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchSyncMapPermission

> PreviewSyncServiceSyncMapSyncMapPermission FetchSyncMapPermission(ctx, ServiceSid, MapSid, Identity)



Fetch a specific Sync Map Permission.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string**|  | 
**MapSid** | **string**| Identifier of the Sync Map. Either a SID or a unique name. | 
**Identity** | **string**| Arbitrary string identifier representing a user associated with an FPA token, assigned by the developer. | 

### Return type

[**PreviewSyncServiceSyncMapSyncMapPermission**](preview.sync.service.sync_map.sync_map_permission.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchTask

> PreviewUnderstandAssistantTask FetchTask(ctx, AssistantSid, Sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string**| The unique ID of the Assistant. | 
**Sid** | **string**| A 34 character string that uniquely identifies this resource. | 

### Return type

[**PreviewUnderstandAssistantTask**](preview.understand.assistant.task.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchTaskActions

> PreviewUnderstandAssistantTaskTaskActions FetchTaskActions(ctx, AssistantSid, TaskSid)



Returns JSON actions for this Task.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string**| The unique ID of the parent Assistant. | 
**TaskSid** | **string**| The unique ID of the Task. | 

### Return type

[**PreviewUnderstandAssistantTaskTaskActions**](preview.understand.assistant.task.task_actions.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchTaskStatistics

> PreviewUnderstandAssistantTaskTaskStatistics FetchTaskStatistics(ctx, AssistantSid, TaskSid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string**| The unique ID of the parent Assistant. | 
**TaskSid** | **string**| The unique ID of the Task associated with this Field. | 

### Return type

[**PreviewUnderstandAssistantTaskTaskStatistics**](preview.understand.assistant.task.task_statistics.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchUsage

> PreviewWirelessSimUsage FetchUsage(ctx, SimSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**SimSid** | **string**|  | 
 **optional** | ***FetchUsageOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a FetchUsageOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **End** | **optional.String**|  | 
 **Start** | **optional.String**|  | 

### Return type

[**PreviewWirelessSimUsage**](preview.wireless.sim.usage.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAssistant

> PreviewUnderstandAssistantReadResponse ListAssistant(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListAssistantOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListAssistantOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **PageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**PreviewUnderstandAssistantReadResponse**](preview_understand_assistantReadResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAuthorizationDocument

> PreviewHostedNumbersAuthorizationDocumentReadResponse ListAuthorizationDocument(ctx, optional)



Retrieve a list of AuthorizationDocuments belonging to the account initiating the request.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListAuthorizationDocumentOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListAuthorizationDocumentOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **Email** | **optional.String**| Email that this AuthorizationDocument will be sent to for signing. | 
 **Status** | **optional.String**| Status of an instance resource. It can hold one of the values: 1. opened 2. signing, 3. signed LOA, 4. canceled, 5. failed. See the section entitled [Status Values](https://www.twilio.com/docs/api/phone-numbers/hosted-number-authorization-documents#status-values) for more information on each of these statuses. | 
 **PageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**PreviewHostedNumbersAuthorizationDocumentReadResponse**](preview_hosted_numbers_authorization_documentReadResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAvailableAddOn

> PreviewMarketplaceAvailableAddOnReadResponse ListAvailableAddOn(ctx, optional)



Retrieve a list of Add-ons currently available to be installed.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListAvailableAddOnOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListAvailableAddOnOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **PageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**PreviewMarketplaceAvailableAddOnReadResponse**](preview_marketplace_available_add_onReadResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAvailableAddOnExtension

> PreviewMarketplaceAvailableAddOnAvailableAddOnExtensionReadResponse ListAvailableAddOnExtension(ctx, AvailableAddOnSid, optional)



Retrieve a list of Extensions for the Available Add-on.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AvailableAddOnSid** | **string**| The SID of the AvailableAddOn resource with the extensions to read. | 
 **optional** | ***ListAvailableAddOnExtensionOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListAvailableAddOnExtensionOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **PageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**PreviewMarketplaceAvailableAddOnAvailableAddOnExtensionReadResponse**](preview_marketplace_available_add_on_available_add_on_extensionReadResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCertificate

> PreviewDeployedDevicesFleetCertificateReadResponse ListCertificate(ctx, FleetSid, optional)



Retrieve a list of all Certificate credentials belonging to the Fleet.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**FleetSid** | **string**|  | 
 **optional** | ***ListCertificateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListCertificateOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **DeviceSid** | **optional.String**| Filters the resulting list of Certificates by a unique string identifier of an authenticated Device. | 
 **PageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**PreviewDeployedDevicesFleetCertificateReadResponse**](preview_deployed_devices_fleet_certificateReadResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCommand

> PreviewWirelessCommandReadResponse ListCommand(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListCommandOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListCommandOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **Device** | **optional.String**|  | 
 **Sim** | **optional.String**|  | 
 **Status** | **optional.String**|  | 
 **Direction** | **optional.String**|  | 
 **PageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**PreviewWirelessCommandReadResponse**](preview_wireless_commandReadResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDay

> PreviewBulkExportsExportDayReadResponse ListDay(ctx, ResourceType, optional)



Retrieve a list of all Days for a resource.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ResourceType** | **string**| The type of communication – Messages, Calls | 
 **optional** | ***ListDayOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListDayOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **PageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**PreviewBulkExportsExportDayReadResponse**](preview_bulk_exports_export_dayReadResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDependentHostedNumberOrder

> PreviewHostedNumbersAuthorizationDocumentDependentHostedNumberOrderReadResponse ListDependentHostedNumberOrder(ctx, SigningDocumentSid, optional)



Retrieve a list of dependent HostedNumberOrders belonging to the AuthorizationDocument.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**SigningDocumentSid** | **string**|  | 
 **optional** | ***ListDependentHostedNumberOrderOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListDependentHostedNumberOrderOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **Status** | **optional.String**| Status of an instance resource. It can hold one of the values: 1. opened 2. signing, 3. signed LOA, 4. canceled, 5. failed. See the section entitled [Status Values](https://www.twilio.com/docs/api/phone-numbers/hosted-number-authorization-documents#status-values) for more information on each of these statuses. | 
 **PhoneNumber** | **optional.String**| An E164 formatted phone number hosted by this HostedNumberOrder. | 
 **IncomingPhoneNumberSid** | **optional.String**| A 34 character string that uniquely identifies the IncomingPhoneNumber resource created by this HostedNumberOrder. | 
 **FriendlyName** | **optional.String**| A human readable description of this resource, up to 64 characters. | 
 **UniqueName** | **optional.String**| Provides a unique and addressable name to be assigned to this HostedNumberOrder, assigned by the developer, to be optionally used in addition to SID. | 
 **PageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**PreviewHostedNumbersAuthorizationDocumentDependentHostedNumberOrderReadResponse**](preview_hosted_numbers_authorization_document_dependent_hosted_number_orderReadResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDeployment

> PreviewDeployedDevicesFleetDeploymentReadResponse ListDeployment(ctx, FleetSid, optional)



Retrieve a list of all Deployments belonging to the Fleet.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**FleetSid** | **string**|  | 
 **optional** | ***ListDeploymentOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListDeploymentOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **PageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**PreviewDeployedDevicesFleetDeploymentReadResponse**](preview_deployed_devices_fleet_deploymentReadResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDevice

> PreviewDeployedDevicesFleetDeviceReadResponse ListDevice(ctx, FleetSid, optional)



Retrieve a list of all Devices belonging to the Fleet.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**FleetSid** | **string**|  | 
 **optional** | ***ListDeviceOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListDeviceOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **DeploymentSid** | **optional.String**| Filters the resulting list of Devices by a unique string identifier of the Deployment they are associated with. | 
 **PageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**PreviewDeployedDevicesFleetDeviceReadResponse**](preview_deployed_devices_fleet_deviceReadResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDocument

> PreviewSyncServiceDocumentReadResponse ListDocument(ctx, ServiceSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string**|  | 
 **optional** | ***ListDocumentOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListDocumentOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **PageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**PreviewSyncServiceDocumentReadResponse**](preview_sync_service_documentReadResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDocumentPermission

> PreviewSyncServiceDocumentDocumentPermissionReadResponse ListDocumentPermission(ctx, ServiceSid, DocumentSid, optional)



Retrieve a list of all Permissions applying to a Sync Document.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string**|  | 
**DocumentSid** | **string**| Identifier of the Sync Document. Either a SID or a unique name. | 
 **optional** | ***ListDocumentPermissionOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListDocumentPermissionOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **PageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**PreviewSyncServiceDocumentDocumentPermissionReadResponse**](preview_sync_service_document_document_permissionReadResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListExportCustomJob

> PreviewBulkExportsExportExportCustomJobReadResponse ListExportCustomJob(ctx, ResourceType, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ResourceType** | **string**| The type of communication – Messages, Calls | 
 **optional** | ***ListExportCustomJobOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListExportCustomJobOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **PageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**PreviewBulkExportsExportExportCustomJobReadResponse**](preview_bulk_exports_export_export_custom_jobReadResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListField

> PreviewUnderstandAssistantTaskFieldReadResponse ListField(ctx, AssistantSid, TaskSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string**| The unique ID of the Assistant. | 
**TaskSid** | **string**| The unique ID of the Task associated with this Field. | 
 **optional** | ***ListFieldOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListFieldOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **PageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**PreviewUnderstandAssistantTaskFieldReadResponse**](preview_understand_assistant_task_fieldReadResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListFieldType

> PreviewUnderstandAssistantFieldTypeReadResponse ListFieldType(ctx, AssistantSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string**|  | 
 **optional** | ***ListFieldTypeOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListFieldTypeOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **PageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**PreviewUnderstandAssistantFieldTypeReadResponse**](preview_understand_assistant_field_typeReadResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListFieldValue

> PreviewUnderstandAssistantFieldTypeFieldValueReadResponse ListFieldValue(ctx, AssistantSid, FieldTypeSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string**|  | 
**FieldTypeSid** | **string**|  | 
 **optional** | ***ListFieldValueOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListFieldValueOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **Language** | **optional.String**| An ISO language-country string of the value. For example: *en-US* | 
 **PageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**PreviewUnderstandAssistantFieldTypeFieldValueReadResponse**](preview_understand_assistant_field_type_field_valueReadResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListFleet

> PreviewDeployedDevicesFleetReadResponse ListFleet(ctx, optional)



Retrieve a list of all Fleets belonging to your account.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListFleetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListFleetOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **PageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**PreviewDeployedDevicesFleetReadResponse**](preview_deployed_devices_fleetReadResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListHostedNumberOrder

> PreviewHostedNumbersHostedNumberOrderReadResponse ListHostedNumberOrder(ctx, optional)



Retrieve a list of HostedNumberOrders belonging to the account initiating the request.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListHostedNumberOrderOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListHostedNumberOrderOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **Status** | **optional.String**| The Status of this HostedNumberOrder. One of &#x60;received&#x60;, &#x60;pending-verification&#x60;, &#x60;verified&#x60;, &#x60;pending-loa&#x60;, &#x60;carrier-processing&#x60;, &#x60;testing&#x60;, &#x60;completed&#x60;, &#x60;failed&#x60;, or &#x60;action-required&#x60;. | 
 **PhoneNumber** | **optional.String**| An E164 formatted phone number hosted by this HostedNumberOrder. | 
 **IncomingPhoneNumberSid** | **optional.String**| A 34 character string that uniquely identifies the IncomingPhoneNumber resource created by this HostedNumberOrder. | 
 **FriendlyName** | **optional.String**| A human readable description of this resource, up to 64 characters. | 
 **UniqueName** | **optional.String**| Provides a unique and addressable name to be assigned to this HostedNumberOrder, assigned by the developer, to be optionally used in addition to SID. | 
 **PageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**PreviewHostedNumbersHostedNumberOrderReadResponse**](preview_hosted_numbers_hosted_number_orderReadResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListInstalledAddOn

> PreviewMarketplaceInstalledAddOnReadResponse ListInstalledAddOn(ctx, optional)



Retrieve a list of Add-ons currently installed on this Account.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListInstalledAddOnOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListInstalledAddOnOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **PageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**PreviewMarketplaceInstalledAddOnReadResponse**](preview_marketplace_installed_add_onReadResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListInstalledAddOnExtension

> PreviewMarketplaceInstalledAddOnInstalledAddOnExtensionReadResponse ListInstalledAddOnExtension(ctx, InstalledAddOnSid, optional)



Retrieve a list of Extensions for the Installed Add-on.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**InstalledAddOnSid** | **string**| The SID of the InstalledAddOn resource with the extensions to read. | 
 **optional** | ***ListInstalledAddOnExtensionOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListInstalledAddOnExtensionOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **PageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**PreviewMarketplaceInstalledAddOnInstalledAddOnExtensionReadResponse**](preview_marketplace_installed_add_on_installed_add_on_extensionReadResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListKey

> PreviewDeployedDevicesFleetKeyReadResponse ListKey(ctx, FleetSid, optional)



Retrieve a list of all Keys credentials belonging to the Fleet.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**FleetSid** | **string**|  | 
 **optional** | ***ListKeyOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListKeyOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **DeviceSid** | **optional.String**| Filters the resulting list of Keys by a unique string identifier of an authenticated Device. | 
 **PageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**PreviewDeployedDevicesFleetKeyReadResponse**](preview_deployed_devices_fleet_keyReadResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListModelBuild

> PreviewUnderstandAssistantModelBuildReadResponse ListModelBuild(ctx, AssistantSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string**|  | 
 **optional** | ***ListModelBuildOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListModelBuildOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **PageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**PreviewUnderstandAssistantModelBuildReadResponse**](preview_understand_assistant_model_buildReadResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListQuery

> PreviewUnderstandAssistantQueryReadResponse ListQuery(ctx, AssistantSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string**| The unique ID of the parent Assistant. | 
 **optional** | ***ListQueryOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListQueryOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **Language** | **optional.String**| An ISO language-country string of the sample. | 
 **ModelBuild** | **optional.String**| The Model Build Sid or unique name of the Model Build to be queried. | 
 **Status** | **optional.String**| A string that described the query status. The values can be: pending_review, reviewed, discarded | 
 **PageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**PreviewUnderstandAssistantQueryReadResponse**](preview_understand_assistant_queryReadResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRatePlan

> PreviewWirelessRatePlanReadResponse ListRatePlan(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListRatePlanOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListRatePlanOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **PageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**PreviewWirelessRatePlanReadResponse**](preview_wireless_rate_planReadResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSample

> PreviewUnderstandAssistantTaskSampleReadResponse ListSample(ctx, AssistantSid, TaskSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string**| The unique ID of the Assistant. | 
**TaskSid** | **string**| The unique ID of the Task associated with this Sample. | 
 **optional** | ***ListSampleOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListSampleOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **Language** | **optional.String**| An ISO language-country string of the sample. | 
 **PageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**PreviewUnderstandAssistantTaskSampleReadResponse**](preview_understand_assistant_task_sampleReadResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListService

> PreviewSyncServiceReadResponse ListService(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListServiceOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListServiceOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **PageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**PreviewSyncServiceReadResponse**](preview_sync_serviceReadResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSim

> PreviewWirelessSimReadResponse ListSim(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListSimOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListSimOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **Status** | **optional.String**|  | 
 **Iccid** | **optional.String**|  | 
 **RatePlan** | **optional.String**|  | 
 **EId** | **optional.String**|  | 
 **SimRegistrationCode** | **optional.String**|  | 
 **PageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**PreviewWirelessSimReadResponse**](preview_wireless_simReadResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSyncList

> PreviewSyncServiceSyncListReadResponse ListSyncList(ctx, ServiceSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string**|  | 
 **optional** | ***ListSyncListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListSyncListOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **PageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**PreviewSyncServiceSyncListReadResponse**](preview_sync_service_sync_listReadResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSyncListItem

> PreviewSyncServiceSyncListSyncListItemReadResponse ListSyncListItem(ctx, ServiceSid, ListSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string**|  | 
**ListSid** | **string**|  | 
 **optional** | ***ListSyncListItemOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListSyncListItemOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **Order** | **optional.String**|  | 
 **From** | **optional.String**|  | 
 **Bounds** | **optional.String**|  | 
 **PageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**PreviewSyncServiceSyncListSyncListItemReadResponse**](preview_sync_service_sync_list_sync_list_itemReadResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSyncListPermission

> PreviewSyncServiceSyncListSyncListPermissionReadResponse ListSyncListPermission(ctx, ServiceSid, ListSid, optional)



Retrieve a list of all Permissions applying to a Sync List.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string**|  | 
**ListSid** | **string**| Identifier of the Sync List. Either a SID or a unique name. | 
 **optional** | ***ListSyncListPermissionOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListSyncListPermissionOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **PageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**PreviewSyncServiceSyncListSyncListPermissionReadResponse**](preview_sync_service_sync_list_sync_list_permissionReadResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSyncMap

> PreviewSyncServiceSyncMapReadResponse ListSyncMap(ctx, ServiceSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string**|  | 
 **optional** | ***ListSyncMapOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListSyncMapOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **PageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**PreviewSyncServiceSyncMapReadResponse**](preview_sync_service_sync_mapReadResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSyncMapItem

> PreviewSyncServiceSyncMapSyncMapItemReadResponse ListSyncMapItem(ctx, ServiceSid, MapSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string**|  | 
**MapSid** | **string**|  | 
 **optional** | ***ListSyncMapItemOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListSyncMapItemOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **Order** | **optional.String**|  | 
 **From** | **optional.String**|  | 
 **Bounds** | **optional.String**|  | 
 **PageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**PreviewSyncServiceSyncMapSyncMapItemReadResponse**](preview_sync_service_sync_map_sync_map_itemReadResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSyncMapPermission

> PreviewSyncServiceSyncMapSyncMapPermissionReadResponse ListSyncMapPermission(ctx, ServiceSid, MapSid, optional)



Retrieve a list of all Permissions applying to a Sync Map.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string**|  | 
**MapSid** | **string**| Identifier of the Sync Map. Either a SID or a unique name. | 
 **optional** | ***ListSyncMapPermissionOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListSyncMapPermissionOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **PageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**PreviewSyncServiceSyncMapSyncMapPermissionReadResponse**](preview_sync_service_sync_map_sync_map_permissionReadResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTask

> PreviewUnderstandAssistantTaskReadResponse ListTask(ctx, AssistantSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string**| The unique ID of the Assistant. | 
 **optional** | ***ListTaskOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListTaskOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **PageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**PreviewUnderstandAssistantTaskReadResponse**](preview_understand_assistant_taskReadResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAssistant

> PreviewUnderstandAssistant UpdateAssistant(ctx, Sid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string**| A 34 character string that uniquely identifies this resource. | 
 **optional** | ***UpdateAssistantOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateAssistantOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **CallbackEvents** | **optional.String**| Space-separated list of callback events that will trigger callbacks. | 
 **CallbackUrl** | **optional.String**| A user-provided URL to send event callbacks to. | 
 **FallbackActions** | [**optional.Interface of map[string]interface{}**](map[string]interface{}.md)| The JSON actions to be executed when the user&#39;s input is not recognized as matching any Task. | 
 **FriendlyName** | **optional.String**| A text description for the Assistant. It is non-unique and can up to 255 characters long. | 
 **InitiationActions** | [**optional.Interface of map[string]interface{}**](map[string]interface{}.md)| The JSON actions to be executed on inbound phone calls when the Assistant has to say something first. | 
 **LogQueries** | **optional.Bool**| A boolean that specifies whether queries should be logged for 30 days further training. If false, no queries will be stored, if true, queries will be stored for 30 days and deleted thereafter. Defaults to true if no value is provided. | 
 **StyleSheet** | [**optional.Interface of map[string]interface{}**](map[string]interface{}.md)| The JSON object that holds the style sheet for the assistant | 
 **UniqueName** | **optional.String**| A user-provided string that uniquely identifies this resource as an alternative to the sid. Unique up to 64 characters long. | 

### Return type

[**PreviewUnderstandAssistant**](preview.understand.assistant.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAssistantFallbackActions

> PreviewUnderstandAssistantAssistantFallbackActions UpdateAssistantFallbackActions(ctx, AssistantSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string**|  | 
 **optional** | ***UpdateAssistantFallbackActionsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateAssistantFallbackActionsOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **FallbackActions** | [**optional.Interface of map[string]interface{}**](map[string]interface{}.md)|  | 

### Return type

[**PreviewUnderstandAssistantAssistantFallbackActions**](preview.understand.assistant.assistant_fallback_actions.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAssistantInitiationActions

> PreviewUnderstandAssistantAssistantInitiationActions UpdateAssistantInitiationActions(ctx, AssistantSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string**|  | 
 **optional** | ***UpdateAssistantInitiationActionsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateAssistantInitiationActionsOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **InitiationActions** | [**optional.Interface of map[string]interface{}**](map[string]interface{}.md)|  | 

### Return type

[**PreviewUnderstandAssistantAssistantInitiationActions**](preview.understand.assistant.assistant_initiation_actions.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAuthorizationDocument

> PreviewHostedNumbersAuthorizationDocument UpdateAuthorizationDocument(ctx, Sid, optional)



Updates a specific AuthorizationDocument.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string**|  | 
 **optional** | ***UpdateAuthorizationDocumentOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateAuthorizationDocumentOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **AddressSid** | **optional.String**| A 34 character string that uniquely identifies the Address resource that is associated with this AuthorizationDocument. | 
 **CcEmails** | [**optional.Interface of []string**](string.md)| Email recipients who will be informed when an Authorization Document has been sent and signed | 
 **ContactPhoneNumber** | **optional.String**| The contact phone number of the person authorized to sign the Authorization Document. | 
 **ContactTitle** | **optional.String**| The title of the person authorized to sign the Authorization Document for this phone number. | 
 **Email** | **optional.String**| Email that this AuthorizationDocument will be sent to for signing. | 
 **HostedNumberOrderSids** | [**optional.Interface of []string**](string.md)| A list of HostedNumberOrder sids that this AuthorizationDocument will authorize for hosting phone number capabilities on Twilio&#39;s platform. | 
 **Status** | **optional.String**| Status of an instance resource. It can hold one of the values: 1. opened 2. signing, 3. signed LOA, 4. canceled, 5. failed. See the section entitled [Status Values](https://www.twilio.com/docs/api/phone-numbers/hosted-number-authorization-documents#status-values) for more information on each of these statuses. | 

### Return type

[**PreviewHostedNumbersAuthorizationDocument**](preview.hosted_numbers.authorization_document.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCertificate

> PreviewDeployedDevicesFleetCertificate UpdateCertificate(ctx, FleetSid, Sid, optional)



Update the given properties of a specific Certificate credential in the Fleet, giving it a friendly name or assigning to a Device.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**FleetSid** | **string**|  | 
**Sid** | **string**| Provides a 34 character string that uniquely identifies the requested Certificate credential resource. | 
 **optional** | ***UpdateCertificateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateCertificateOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **DeviceSid** | **optional.String**| Provides the unique string identifier of an existing Device to become authenticated with this Certificate credential. | 
 **FriendlyName** | **optional.String**| Provides a human readable descriptive text for this Certificate credential, up to 256 characters long. | 

### Return type

[**PreviewDeployedDevicesFleetCertificate**](preview.deployed_devices.fleet.certificate.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDeployment

> PreviewDeployedDevicesFleetDeployment UpdateDeployment(ctx, FleetSid, Sid, optional)



Update the given properties of a specific Deployment credential in the Fleet, giving it a friendly name or linking to a specific Twilio Sync service instance.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**FleetSid** | **string**|  | 
**Sid** | **string**| Provides a 34 character string that uniquely identifies the requested Deployment resource. | 
 **optional** | ***UpdateDeploymentOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateDeploymentOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **FriendlyName** | **optional.String**| Provides a human readable descriptive text for this Deployment, up to 64 characters long | 
 **SyncServiceSid** | **optional.String**| Provides the unique string identifier of the Twilio Sync service instance that will be linked to and accessible by this Deployment. | 

### Return type

[**PreviewDeployedDevicesFleetDeployment**](preview.deployed_devices.fleet.deployment.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDevice

> PreviewDeployedDevicesFleetDevice UpdateDevice(ctx, FleetSid, Sid, optional)



Update the given properties of a specific Device in the Fleet, giving it a friendly name, assigning to a Deployment, or a human identity.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**FleetSid** | **string**|  | 
**Sid** | **string**| Provides a 34 character string that uniquely identifies the requested Device resource. | 
 **optional** | ***UpdateDeviceOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateDeviceOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **DeploymentSid** | **optional.String**| Specifies the unique string identifier of the Deployment group that this Device is going to be associated with. | 
 **Enabled** | **optional.Bool**|  | 
 **FriendlyName** | **optional.String**| Provides a human readable descriptive text to be assigned to this Device, up to 256 characters long. | 
 **Identity** | **optional.String**| Provides an arbitrary string identifier representing a human user to be associated with this Device, up to 256 characters long. | 

### Return type

[**PreviewDeployedDevicesFleetDevice**](preview.deployed_devices.fleet.device.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDocument

> PreviewSyncServiceDocument UpdateDocument(ctx, ServiceSid, Sid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string**|  | 
**Sid** | **string**|  | 
 **optional** | ***UpdateDocumentOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateDocumentOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **IfMatch** | **optional.String**| The If-Match HTTP request header | 
 **Data** | [**optional.Interface of map[string]interface{}**](map[string]interface{}.md)|  | 

### Return type

[**PreviewSyncServiceDocument**](preview.sync.service.document.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDocumentPermission

> PreviewSyncServiceDocumentDocumentPermission UpdateDocumentPermission(ctx, ServiceSid, DocumentSid, Identity, optional)



Update an identity's access to a specific Sync Document.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string**| The unique SID identifier of the Sync Service Instance. | 
**DocumentSid** | **string**| Identifier of the Sync Document. Either a SID or a unique name. | 
**Identity** | **string**| Arbitrary string identifier representing a human user associated with an FPA token, assigned by the developer. | 
 **optional** | ***UpdateDocumentPermissionOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateDocumentPermissionOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **Manage** | **optional.Bool**| Boolean flag specifying whether the identity can delete the Sync Document. | 
 **Read** | **optional.Bool**| Boolean flag specifying whether the identity can read the Sync Document. | 
 **Write** | **optional.Bool**| Boolean flag specifying whether the identity can update the Sync Document. | 

### Return type

[**PreviewSyncServiceDocumentDocumentPermission**](preview.sync.service.document.document_permission.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateExportConfiguration

> PreviewBulkExportsExportConfiguration UpdateExportConfiguration(ctx, ResourceType, optional)



Update a specific Export Configuration.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ResourceType** | **string**| The type of communication – Messages, Calls | 
 **optional** | ***UpdateExportConfigurationOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateExportConfigurationOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **Enabled** | **optional.Bool**| If true, Twilio will automatically generate every day&#39;s file when the day is over. | 
 **WebhookMethod** | **optional.String**| Sets whether Twilio should call a webhook URL when the automatic generation is complete, using GET or POST. The actual destination is set in the webhook_url | 
 **WebhookUrl** | **optional.String**| Stores the URL destination for the method specified in webhook_method. | 

### Return type

[**PreviewBulkExportsExportConfiguration**](preview.bulk_exports.export_configuration.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateFieldType

> PreviewUnderstandAssistantFieldType UpdateFieldType(ctx, AssistantSid, Sid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string**|  | 
**Sid** | **string**|  | 
 **optional** | ***UpdateFieldTypeOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateFieldTypeOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **FriendlyName** | **optional.String**| A user-provided string that identifies this resource. It is non-unique and can up to 255 characters long. | 
 **UniqueName** | **optional.String**| A user-provided string that uniquely identifies this resource as an alternative to the sid. Unique up to 64 characters long. | 

### Return type

[**PreviewUnderstandAssistantFieldType**](preview.understand.assistant.field_type.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateFleet

> PreviewDeployedDevicesFleet UpdateFleet(ctx, Sid, optional)



Update the friendly name property of a specific Fleet in your account.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string**| Provides a 34 character string that uniquely identifies the requested Fleet resource. | 
 **optional** | ***UpdateFleetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateFleetOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **DefaultDeploymentSid** | **optional.String**| Provides a string identifier of a Deployment that is going to be used as a default one for this Fleet. | 
 **FriendlyName** | **optional.String**| Provides a human readable descriptive text for this Fleet, up to 256 characters long. | 

### Return type

[**PreviewDeployedDevicesFleet**](preview.deployed_devices.fleet.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateHostedNumberOrder

> PreviewHostedNumbersHostedNumberOrder UpdateHostedNumberOrder(ctx, Sid, optional)



Updates a specific HostedNumberOrder.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string**|  | 
 **optional** | ***UpdateHostedNumberOrderOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateHostedNumberOrderOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **CallDelay** | **optional.Int32**| The number of seconds, between 0 and 60, to delay before initiating the verification call. Defaults to 0. | 
 **CcEmails** | [**optional.Interface of []string**](string.md)| Optional. A list of emails that LOA document for this HostedNumberOrder will be carbon copied to. | 
 **Email** | **optional.String**| Email of the owner of this phone number that is being hosted. | 
 **Extension** | **optional.String**| Digits to dial after connecting the verification call. | 
 **FriendlyName** | **optional.String**| A 64 character string that is a human readable text that describes this resource. | 
 **Status** | **optional.String**| User can only post to &#x60;pending-verification&#x60; status to transition the HostedNumberOrder to initiate a verification call or verification of ownership with a copy of a phone bill. | 
 **UniqueName** | **optional.String**| Provides a unique and addressable name to be assigned to this HostedNumberOrder, assigned by the developer, to be optionally used in addition to SID. | 
 **VerificationCode** | **optional.String**| A verification code that is given to the user via a phone call to the phone number that is being hosted. | 
 **VerificationDocumentSid** | **optional.String**| Optional. The unique sid identifier of the Identity Document that represents the document for verifying ownership of the number to be hosted. Required when VerificationType is phone-bill. | 
 **VerificationType** | **optional.String**| Optional. The method used for verifying ownership of the number to be hosted. One of phone-call (default) or phone-bill. | 

### Return type

[**PreviewHostedNumbersHostedNumberOrder**](preview.hosted_numbers.hosted_number_order.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateInstalledAddOn

> PreviewMarketplaceInstalledAddOn UpdateInstalledAddOn(ctx, Sid, optional)



Update an Add-on installation for the Account specified.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string**| The SID of the InstalledAddOn resource to update. | 
 **optional** | ***UpdateInstalledAddOnOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateInstalledAddOnOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **Configuration** | [**optional.Interface of map[string]interface{}**](map[string]interface{}.md)| Valid JSON object that conform to the configuration schema exposed by the associated AvailableAddOn resource. This is only required by Add-ons that need to be configured | 
 **UniqueName** | **optional.String**| An application-defined string that uniquely identifies the resource. This value must be unique within the Account. | 

### Return type

[**PreviewMarketplaceInstalledAddOn**](preview.marketplace.installed_add_on.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateInstalledAddOnExtension

> PreviewMarketplaceInstalledAddOnInstalledAddOnExtension UpdateInstalledAddOnExtension(ctx, InstalledAddOnSid, Sid, optional)



Update an Extension for an Add-on installation.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**InstalledAddOnSid** | **string**| The SID of the InstalledAddOn resource with the extension to update. | 
**Sid** | **string**| The SID of the InstalledAddOn Extension resource to update. | 
 **optional** | ***UpdateInstalledAddOnExtensionOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateInstalledAddOnExtensionOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **Enabled** | **optional.Bool**| Whether the Extension should be invoked. | 

### Return type

[**PreviewMarketplaceInstalledAddOnInstalledAddOnExtension**](preview.marketplace.installed_add_on.installed_add_on_extension.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateKey

> PreviewDeployedDevicesFleetKey UpdateKey(ctx, FleetSid, Sid, optional)



Update the given properties of a specific Key credential in the Fleet, giving it a friendly name or assigning to a Device.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**FleetSid** | **string**|  | 
**Sid** | **string**| Provides a 34 character string that uniquely identifies the requested Key credential resource. | 
 **optional** | ***UpdateKeyOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateKeyOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **DeviceSid** | **optional.String**| Provides the unique string identifier of an existing Device to become authenticated with this Key credential. | 
 **FriendlyName** | **optional.String**| Provides a human readable descriptive text for this Key credential, up to 256 characters long. | 

### Return type

[**PreviewDeployedDevicesFleetKey**](preview.deployed_devices.fleet.key.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateModelBuild

> PreviewUnderstandAssistantModelBuild UpdateModelBuild(ctx, AssistantSid, Sid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string**|  | 
**Sid** | **string**|  | 
 **optional** | ***UpdateModelBuildOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateModelBuildOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **UniqueName** | **optional.String**| A user-provided string that uniquely identifies this resource as an alternative to the sid. Unique up to 64 characters long. For example: v0.1 | 

### Return type

[**PreviewUnderstandAssistantModelBuild**](preview.understand.assistant.model_build.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateQuery

> PreviewUnderstandAssistantQuery UpdateQuery(ctx, AssistantSid, Sid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string**| The unique ID of the parent Assistant. | 
**Sid** | **string**| A 34 character string that uniquely identifies this resource. | 
 **optional** | ***UpdateQueryOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateQueryOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **SampleSid** | **optional.String**| An optional reference to the Sample created from this query. | 
 **Status** | **optional.String**| A string that described the query status. The values can be: pending_review, reviewed, discarded | 

### Return type

[**PreviewUnderstandAssistantQuery**](preview.understand.assistant.query.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRatePlan

> PreviewWirelessRatePlan UpdateRatePlan(ctx, Sid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string**|  | 
 **optional** | ***UpdateRatePlanOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateRatePlanOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **FriendlyName** | **optional.String**|  | 
 **UniqueName** | **optional.String**|  | 

### Return type

[**PreviewWirelessRatePlan**](preview.wireless.rate_plan.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSample

> PreviewUnderstandAssistantTaskSample UpdateSample(ctx, AssistantSid, TaskSid, Sid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string**| The unique ID of the Assistant. | 
**TaskSid** | **string**| The unique ID of the Task associated with this Sample. | 
**Sid** | **string**| A 34 character string that uniquely identifies this resource. | 
 **optional** | ***UpdateSampleOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateSampleOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **Language** | **optional.String**| An ISO language-country string of the sample. | 
 **SourceChannel** | **optional.String**| The communication channel the sample was captured. It can be: *voice*, *sms*, *chat*, *alexa*, *google-assistant*, or *slack*. If not included the value will be null | 
 **TaggedText** | **optional.String**| The text example of how end-users may express this task. The sample may contain Field tag blocks. | 

### Return type

[**PreviewUnderstandAssistantTaskSample**](preview.understand.assistant.task.sample.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateService

> PreviewSyncService UpdateService(ctx, Sid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string**|  | 
 **optional** | ***UpdateServiceOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateServiceOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **AclEnabled** | **optional.Bool**|  | 
 **FriendlyName** | **optional.String**|  | 
 **ReachabilityWebhooksEnabled** | **optional.Bool**|  | 
 **WebhookUrl** | **optional.String**|  | 

### Return type

[**PreviewSyncService**](preview.sync.service.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSim

> PreviewWirelessSim UpdateSim(ctx, Sid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string**|  | 
 **optional** | ***UpdateSimOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateSimOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **CallbackMethod** | **optional.String**|  | 
 **CallbackUrl** | **optional.String**|  | 
 **CommandsCallbackMethod** | **optional.String**|  | 
 **CommandsCallbackUrl** | **optional.String**|  | 
 **FriendlyName** | **optional.String**|  | 
 **RatePlan** | **optional.String**|  | 
 **SmsFallbackMethod** | **optional.String**|  | 
 **SmsFallbackUrl** | **optional.String**|  | 
 **SmsMethod** | **optional.String**|  | 
 **SmsUrl** | **optional.String**|  | 
 **Status** | **optional.String**|  | 
 **UniqueName** | **optional.String**|  | 
 **VoiceFallbackMethod** | **optional.String**|  | 
 **VoiceFallbackUrl** | **optional.String**|  | 
 **VoiceMethod** | **optional.String**|  | 
 **VoiceUrl** | **optional.String**|  | 

### Return type

[**PreviewWirelessSim**](preview.wireless.sim.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateStyleSheet

> PreviewUnderstandAssistantStyleSheet UpdateStyleSheet(ctx, AssistantSid, optional)



Updates the style sheet for an assistant identified by {AssistantSid} or {AssistantUniqueName}.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string**| The unique ID of the Assistant | 
 **optional** | ***UpdateStyleSheetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateStyleSheetOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **StyleSheet** | [**optional.Interface of map[string]interface{}**](map[string]interface{}.md)| The JSON Style sheet string | 

### Return type

[**PreviewUnderstandAssistantStyleSheet**](preview.understand.assistant.style_sheet.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSyncListItem

> PreviewSyncServiceSyncListSyncListItem UpdateSyncListItem(ctx, ServiceSid, ListSid, Index, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string**|  | 
**ListSid** | **string**|  | 
**Index** | **int32**|  | 
 **optional** | ***UpdateSyncListItemOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateSyncListItemOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **IfMatch** | **optional.String**| The If-Match HTTP request header | 
 **Data** | [**optional.Interface of map[string]interface{}**](map[string]interface{}.md)|  | 

### Return type

[**PreviewSyncServiceSyncListSyncListItem**](preview.sync.service.sync_list.sync_list_item.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSyncListPermission

> PreviewSyncServiceSyncListSyncListPermission UpdateSyncListPermission(ctx, ServiceSid, ListSid, Identity, optional)



Update an identity's access to a specific Sync List.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string**| The unique SID identifier of the Sync Service Instance. | 
**ListSid** | **string**| Identifier of the Sync List. Either a SID or a unique name. | 
**Identity** | **string**| Arbitrary string identifier representing a human user associated with an FPA token, assigned by the developer. | 
 **optional** | ***UpdateSyncListPermissionOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateSyncListPermissionOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **Manage** | **optional.Bool**| Boolean flag specifying whether the identity can delete the Sync List. | 
 **Read** | **optional.Bool**| Boolean flag specifying whether the identity can read the Sync List. | 
 **Write** | **optional.Bool**| Boolean flag specifying whether the identity can create, update and delete Items of the Sync List. | 

### Return type

[**PreviewSyncServiceSyncListSyncListPermission**](preview.sync.service.sync_list.sync_list_permission.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSyncMapItem

> PreviewSyncServiceSyncMapSyncMapItem UpdateSyncMapItem(ctx, ServiceSid, MapSid, Key, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string**|  | 
**MapSid** | **string**|  | 
**Key** | **string**|  | 
 **optional** | ***UpdateSyncMapItemOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateSyncMapItemOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **IfMatch** | **optional.String**| The If-Match HTTP request header | 
 **Data** | [**optional.Interface of map[string]interface{}**](map[string]interface{}.md)|  | 

### Return type

[**PreviewSyncServiceSyncMapSyncMapItem**](preview.sync.service.sync_map.sync_map_item.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSyncMapPermission

> PreviewSyncServiceSyncMapSyncMapPermission UpdateSyncMapPermission(ctx, ServiceSid, MapSid, Identity, optional)



Update an identity's access to a specific Sync Map.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string**| The unique SID identifier of the Sync Service Instance. | 
**MapSid** | **string**| Identifier of the Sync Map. Either a SID or a unique name. | 
**Identity** | **string**| Arbitrary string identifier representing a human user associated with an FPA token, assigned by the developer. | 
 **optional** | ***UpdateSyncMapPermissionOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateSyncMapPermissionOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **Manage** | **optional.Bool**| Boolean flag specifying whether the identity can delete the Sync Map. | 
 **Read** | **optional.Bool**| Boolean flag specifying whether the identity can read the Sync Map. | 
 **Write** | **optional.Bool**| Boolean flag specifying whether the identity can create, update and delete Items of the Sync Map. | 

### Return type

[**PreviewSyncServiceSyncMapSyncMapPermission**](preview.sync.service.sync_map.sync_map_permission.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTask

> PreviewUnderstandAssistantTask UpdateTask(ctx, AssistantSid, Sid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string**| The unique ID of the Assistant. | 
**Sid** | **string**| A 34 character string that uniquely identifies this resource. | 
 **optional** | ***UpdateTaskOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateTaskOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **Actions** | [**optional.Interface of map[string]interface{}**](map[string]interface{}.md)| A user-provided JSON object encoded as a string to specify the actions for this task. It is optional and non-unique. | 
 **ActionsUrl** | **optional.String**| User-provided HTTP endpoint where from the assistant fetches actions | 
 **FriendlyName** | **optional.String**| A user-provided string that identifies this resource. It is non-unique and can up to 255 characters long. | 
 **UniqueName** | **optional.String**| A user-provided string that uniquely identifies this resource as an alternative to the sid. Unique up to 64 characters long. | 

### Return type

[**PreviewUnderstandAssistantTask**](preview.understand.assistant.task.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTaskActions

> PreviewUnderstandAssistantTaskTaskActions UpdateTaskActions(ctx, AssistantSid, TaskSid, optional)



Updates the actions of an Task identified by {TaskSid} or {TaskUniqueName}.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string**| The unique ID of the parent Assistant. | 
**TaskSid** | **string**| The unique ID of the Task. | 
 **optional** | ***UpdateTaskActionsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateTaskActionsOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **Actions** | [**optional.Interface of map[string]interface{}**](map[string]interface{}.md)| The JSON actions that instruct the Assistant how to perform this task. | 

### Return type

[**PreviewUnderstandAssistantTaskTaskActions**](preview.understand.assistant.task.task_actions.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

