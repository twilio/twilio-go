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
 **callbackEvents** | **optional.String**| Space-separated list of callback events that will trigger callbacks. | 
 **callbackUrl** | **optional.String**| A user-provided URL to send event callbacks to. | 
 **fallbackActions** | [**optional.Interface of map[string]interface{}**](map[string]interface{}.md)| The JSON actions to be executed when the user&#39;s input is not recognized as matching any Task. | 
 **friendlyName** | **optional.String**| A text description for the Assistant. It is non-unique and can up to 255 characters long. | 
 **initiationActions** | [**optional.Interface of map[string]interface{}**](map[string]interface{}.md)| The JSON actions to be executed on inbound phone calls when the Assistant has to say something first. | 
 **logQueries** | **optional.Bool**| A boolean that specifies whether queries should be logged for 30 days further training. If false, no queries will be stored, if true, queries will be stored for 30 days and deleted thereafter. Defaults to true if no value is provided. | 
 **styleSheet** | [**optional.Interface of map[string]interface{}**](map[string]interface{}.md)| The JSON object that holds the style sheet for the assistant | 
 **uniqueName** | **optional.String**| A user-provided string that uniquely identifies this resource as an alternative to the sid. Unique up to 64 characters long. | 

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
 **addressSid** | **optional.String**| A 34 character string that uniquely identifies the Address resource that is associated with this AuthorizationDocument. | 
 **ccEmails** | [**optional.Interface of []string**](string.md)| Email recipients who will be informed when an Authorization Document has been sent and signed. | 
 **contactPhoneNumber** | **optional.String**| The contact phone number of the person authorized to sign the Authorization Document. | 
 **contactTitle** | **optional.String**| The title of the person authorized to sign the Authorization Document for this phone number. | 
 **email** | **optional.String**| Email that this AuthorizationDocument will be sent to for signing. | 
 **hostedNumberOrderSids** | [**optional.Interface of []string**](string.md)| A list of HostedNumberOrder sids that this AuthorizationDocument will authorize for hosting phone number capabilities on Twilio&#39;s platform. | 

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

> PreviewDeployedDevicesFleetCertificate CreateCertificate(ctx, fleetSid, optional)



Enroll a new Certificate credential to the Fleet, optionally giving it a friendly name and assigning to a Device.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fleetSid** | **string**|  | 
 **optional** | ***CreateCertificateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateCertificateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **certificateData** | **optional.String**| Provides a URL encoded representation of the public certificate in PEM format. | 
 **deviceSid** | **optional.String**| Provides the unique string identifier of an existing Device to become authenticated with this Certificate credential. | 
 **friendlyName** | **optional.String**| Provides a human readable descriptive text for this Certificate credential, up to 256 characters long. | 

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

> PreviewTrustedCommsBrandedChannelChannel CreateChannel(ctx, brandedChannelSid, optional)



Associate a channel to a branded channel

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**brandedChannelSid** | **string**| The unique SID identifier of the Branded Channel. The given phone number is going to be assigned to this Branded Channel | 
 **optional** | ***CreateChannelOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateChannelOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **phoneNumberSid** | **optional.String**| The unique SID identifier of the Phone Number of the Phone number to be assigned to the Branded Channel. | 

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
 **callbackMethod** | **optional.String**|  | 
 **callbackUrl** | **optional.String**|  | 
 **command** | **optional.String**|  | 
 **commandMode** | **optional.String**|  | 
 **device** | **optional.String**|  | 
 **includeSid** | **optional.String**|  | 
 **sim** | **optional.String**|  | 

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

> PreviewDeployedDevicesFleetDeployment CreateDeployment(ctx, fleetSid, optional)



Create a new Deployment in the Fleet, optionally giving it a friendly name and linking to a specific Twilio Sync service instance.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fleetSid** | **string**|  | 
 **optional** | ***CreateDeploymentOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateDeploymentOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **friendlyName** | **optional.String**| Provides a human readable descriptive text for this Deployment, up to 256 characters long. | 
 **syncServiceSid** | **optional.String**| Provides the unique string identifier of the Twilio Sync service instance that will be linked to and accessible by this Deployment. | 

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

> PreviewDeployedDevicesFleetDevice CreateDevice(ctx, fleetSid, optional)



Create a new Device in the Fleet, optionally giving it a unique name, friendly name, and assigning to a Deployment and/or human identity.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fleetSid** | **string**|  | 
 **optional** | ***CreateDeviceOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateDeviceOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deploymentSid** | **optional.String**| Specifies the unique string identifier of the Deployment group that this Device is going to be associated with. | 
 **enabled** | **optional.Bool**|  | 
 **friendlyName** | **optional.String**| Provides a human readable descriptive text to be assigned to this Device, up to 256 characters long. | 
 **identity** | **optional.String**| Provides an arbitrary string identifier representing a human user to be associated with this Device, up to 256 characters long. | 
 **uniqueName** | **optional.String**| Provides a unique and addressable name to be assigned to this Device, to be used in addition to SID, up to 128 characters long. | 

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

> PreviewSyncServiceDocument CreateDocument(ctx, serviceSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**|  | 
 **optional** | ***CreateDocumentOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateDocumentOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **data** | [**optional.Interface of map[string]interface{}**](map[string]interface{}.md)|  | 
 **uniqueName** | **optional.String**|  | 

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

> PreviewBulkExportsExportExportCustomJob CreateExportCustomJob(ctx, resourceType, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourceType** | **string**| The type of communication – Messages or Calls | 
 **optional** | ***CreateExportCustomJobOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateExportCustomJobOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **email** | **optional.String**| The optional email to send the completion notification to | 
 **endDay** | **optional.String**| The end day for the custom export specified as a string in the format of yyyy-mm-dd. End day is inclusive and must be 2 days earlier than the current UTC day. | 
 **friendlyName** | **optional.String**| The friendly name specified when creating the job | 
 **startDay** | **optional.String**| The start day for the custom export specified as a string in the format of yyyy-mm-dd | 
 **webhookMethod** | **optional.String**| This is the method used to call the webhook on completion of the job. If this is supplied, &#x60;WebhookUrl&#x60; must also be supplied. | 
 **webhookUrl** | **optional.String**| The optional webhook url called on completion of the job. If this is supplied, &#x60;WebhookMethod&#x60; must also be supplied. | 

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

> PreviewUnderstandAssistantTaskField CreateField(ctx, assistantSid, taskSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assistantSid** | **string**| The unique ID of the parent Assistant. | 
**taskSid** | **string**| The unique ID of the Task associated with this Field. | 
 **optional** | ***CreateFieldOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateFieldOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fieldType** | **optional.String**| The unique name or sid of the FieldType. It can be any [Built-in Field Type](https://www.twilio.com/docs/assistant/api/built-in-field-types) or the unique_name or the Field Type sid of a custom Field Type. | 
 **uniqueName** | **optional.String**| A user-provided string that uniquely identifies this resource as an alternative to the sid. Unique up to 64 characters long. | 

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

> PreviewUnderstandAssistantFieldType CreateFieldType(ctx, assistantSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assistantSid** | **string**|  | 
 **optional** | ***CreateFieldTypeOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateFieldTypeOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **friendlyName** | **optional.String**| A user-provided string that identifies this resource. It is non-unique and can up to 255 characters long. | 
 **uniqueName** | **optional.String**| A user-provided string that uniquely identifies this resource as an alternative to the sid. Unique up to 64 characters long. | 

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

> PreviewUnderstandAssistantFieldTypeFieldValue CreateFieldValue(ctx, assistantSid, fieldTypeSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assistantSid** | **string**|  | 
**fieldTypeSid** | **string**|  | 
 **optional** | ***CreateFieldValueOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateFieldValueOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **language** | **optional.String**| An ISO language-country string of the value. | 
 **synonymOf** | **optional.String**| A value that indicates this field value is a synonym of. Empty if the value is not a synonym. | 
 **value** | **optional.String**| A user-provided string that uniquely identifies this resource as an alternative to the sid. Unique up to 64 characters long. | 

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
 **friendlyName** | **optional.String**| Provides a human readable descriptive text for this Fleet, up to 256 characters long. | 

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
 **accountSid** | **optional.String**| This defaults to the AccountSid of the authorization the user is using. This can be provided to specify a subaccount to add the HostedNumberOrder to. | 
 **addressSid** | **optional.String**| Optional. A 34 character string that uniquely identifies the Address resource that represents the address of the owner of this phone number. | 
 **ccEmails** | [**optional.Interface of []string**](string.md)| Optional. A list of emails that the LOA document for this HostedNumberOrder will be carbon copied to. | 
 **email** | **optional.String**| Optional. Email of the owner of this phone number that is being hosted. | 
 **friendlyName** | **optional.String**| A 64 character string that is a human readable text that describes this resource. | 
 **phoneNumber** | **optional.String**| The number to host in [+E.164](https://en.wikipedia.org/wiki/E.164) format | 
 **smsApplicationSid** | **optional.String**| Optional. The 34 character sid of the application Twilio should use to handle SMS messages sent to this number. If a &#x60;SmsApplicationSid&#x60; is present, Twilio will ignore all of the SMS urls above and use those set on the application. | 
 **smsCapability** | **optional.Bool**| Used to specify that the SMS capability will be hosted on Twilio&#39;s platform. | 
 **smsFallbackMethod** | **optional.String**| The HTTP method that should be used to request the SmsFallbackUrl. Must be either &#x60;GET&#x60; or &#x60;POST&#x60;. This will be copied onto the IncomingPhoneNumber resource. | 
 **smsFallbackUrl** | **optional.String**| A URL that Twilio will request if an error occurs requesting or executing the TwiML defined by SmsUrl. This will be copied onto the IncomingPhoneNumber resource. | 
 **smsMethod** | **optional.String**| The HTTP method that should be used to request the SmsUrl. Must be either &#x60;GET&#x60; or &#x60;POST&#x60;.  This will be copied onto the IncomingPhoneNumber resource. | 
 **smsUrl** | **optional.String**| The URL that Twilio should request when somebody sends an SMS to the phone number. This will be copied onto the IncomingPhoneNumber resource. | 
 **statusCallbackMethod** | **optional.String**| Optional. The Status Callback Method attached to the IncomingPhoneNumber resource. | 
 **statusCallbackUrl** | **optional.String**| Optional. The Status Callback URL attached to the IncomingPhoneNumber resource. | 
 **uniqueName** | **optional.String**| Optional. Provides a unique and addressable name to be assigned to this HostedNumberOrder, assigned by the developer, to be optionally used in addition to SID. | 
 **verificationDocumentSid** | **optional.String**| Optional. The unique sid identifier of the Identity Document that represents the document for verifying ownership of the number to be hosted. Required when VerificationType is phone-bill. | 
 **verificationType** | **optional.String**| Optional. The method used for verifying ownership of the number to be hosted. One of phone-call (default) or phone-bill. | 

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
 **acceptTermsOfService** | **optional.Bool**| Whether the Terms of Service were accepted. | 
 **availableAddOnSid** | **optional.String**| The SID of the AvaliableAddOn to install. | 
 **configuration** | [**optional.Interface of map[string]interface{}**](map[string]interface{}.md)| The JSON object that represents the configuration of the new Add-on being installed. | 
 **uniqueName** | **optional.String**| An application-defined string that uniquely identifies the resource. This value must be unique within the Account. | 

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

> PreviewDeployedDevicesFleetKey CreateKey(ctx, fleetSid, optional)



Create a new Key credential in the Fleet, optionally giving it a friendly name and assigning to a Device.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fleetSid** | **string**|  | 
 **optional** | ***CreateKeyOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateKeyOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deviceSid** | **optional.String**| Provides the unique string identifier of an existing Device to become authenticated with this Key credential. | 
 **friendlyName** | **optional.String**| Provides a human readable descriptive text for this Key credential, up to 256 characters long. | 

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

> PreviewUnderstandAssistantModelBuild CreateModelBuild(ctx, assistantSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assistantSid** | **string**|  | 
 **optional** | ***CreateModelBuildOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateModelBuildOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **statusCallback** | **optional.String**|  | 
 **uniqueName** | **optional.String**| A user-provided string that uniquely identifies this resource as an alternative to the sid. Unique up to 64 characters long. For example: v0.1 | 

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

> PreviewUnderstandAssistantQuery CreateQuery(ctx, assistantSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assistantSid** | **string**| The unique ID of the parent Assistant. | 
 **optional** | ***CreateQueryOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateQueryOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **field** | **optional.String**| Constraints the query to a given Field with an task. Useful when you know the Field you are expecting. It accepts one field in the format *task-unique-name-1*:*field-unique-name* | 
 **language** | **optional.String**| An ISO language-country string of the sample. | 
 **modelBuild** | **optional.String**| The Model Build Sid or unique name of the Model Build to be queried. | 
 **query** | **optional.String**| A user-provided string that uniquely identifies this resource as an alternative to the sid. It can be up to 2048 characters long. | 
 **tasks** | **optional.String**| Constraints the query to a set of tasks. Useful when you need to constrain the paths the user can take. Tasks should be comma separated *task-unique-name-1*, *task-unique-name-2* | 

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
 **commandsEnabled** | **optional.Bool**|  | 
 **dataEnabled** | **optional.Bool**|  | 
 **dataLimit** | **optional.Int32**|  | 
 **dataMetering** | **optional.String**|  | 
 **friendlyName** | **optional.String**|  | 
 **internationalRoaming** | [**optional.Interface of []string**](string.md)|  | 
 **messagingEnabled** | **optional.Bool**|  | 
 **nationalRoamingEnabled** | **optional.Bool**|  | 
 **uniqueName** | **optional.String**|  | 
 **voiceEnabled** | **optional.Bool**|  | 

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

> PreviewUnderstandAssistantTaskSample CreateSample(ctx, assistantSid, taskSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assistantSid** | **string**| The unique ID of the Assistant. | 
**taskSid** | **string**| The unique ID of the Task associated with this Sample. | 
 **optional** | ***CreateSampleOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateSampleOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **language** | **optional.String**| An ISO language-country string of the sample. | 
 **sourceChannel** | **optional.String**| The communication channel the sample was captured. It can be: *voice*, *sms*, *chat*, *alexa*, *google-assistant*, or *slack*. If not included the value will be null | 
 **taggedText** | **optional.String**| The text example of how end-users may express this task. The sample may contain Field tag blocks. | 

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
 **aclEnabled** | **optional.Bool**|  | 
 **friendlyName** | **optional.String**|  | 
 **reachabilityWebhooksEnabled** | **optional.Bool**|  | 
 **webhookUrl** | **optional.String**|  | 

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

> PreviewSyncServiceSyncList CreateSyncList(ctx, serviceSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**|  | 
 **optional** | ***CreateSyncListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateSyncListOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **uniqueName** | **optional.String**|  | 

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

> PreviewSyncServiceSyncListSyncListItem CreateSyncListItem(ctx, serviceSid, listSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**|  | 
**listSid** | **string**|  | 
 **optional** | ***CreateSyncListItemOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateSyncListItemOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **data** | [**optional.Interface of map[string]interface{}**](map[string]interface{}.md)|  | 

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

> PreviewSyncServiceSyncMap CreateSyncMap(ctx, serviceSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**|  | 
 **optional** | ***CreateSyncMapOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateSyncMapOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **uniqueName** | **optional.String**|  | 

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

> PreviewSyncServiceSyncMapSyncMapItem CreateSyncMapItem(ctx, serviceSid, mapSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**|  | 
**mapSid** | **string**|  | 
 **optional** | ***CreateSyncMapItemOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateSyncMapItemOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **data** | [**optional.Interface of map[string]interface{}**](map[string]interface{}.md)|  | 
 **key** | **optional.String**|  | 

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

> PreviewUnderstandAssistantTask CreateTask(ctx, assistantSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assistantSid** | **string**| The unique ID of the Assistant. | 
 **optional** | ***CreateTaskOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateTaskOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **actions** | [**optional.Interface of map[string]interface{}**](map[string]interface{}.md)| A user-provided JSON object encoded as a string to specify the actions for this task. It is optional and non-unique. | 
 **actionsUrl** | **optional.String**| User-provided HTTP endpoint where from the assistant fetches actions | 
 **friendlyName** | **optional.String**| A user-provided string that identifies this resource. It is non-unique and can up to 255 characters long. | 
 **uniqueName** | **optional.String**| A user-provided string that uniquely identifies this resource as an alternative to the sid. Unique up to 64 characters long. | 

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

> DeleteAssistant(ctx, sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sid** | **string**| A 34 character string that uniquely identifies this resource. | 

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

> DeleteCertificate(ctx, fleetSid, sid)



Unregister a specific Certificate credential from the Fleet, effectively disallowing any inbound client connections that are presenting it.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fleetSid** | **string**|  | 
**sid** | **string**| Provides a 34 character string that uniquely identifies the requested Certificate credential resource. | 

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

> DeleteDeployment(ctx, fleetSid, sid)



Delete a specific Deployment from the Fleet, leaving associated devices effectively undeployed.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fleetSid** | **string**|  | 
**sid** | **string**| Provides a 34 character string that uniquely identifies the requested Deployment resource. | 

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

> DeleteDevice(ctx, fleetSid, sid)



Delete a specific Device from the Fleet, also removing it from associated Deployments.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fleetSid** | **string**|  | 
**sid** | **string**| Provides a 34 character string that uniquely identifies the requested Device resource. | 

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

> DeleteDocument(ctx, serviceSid, sid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**|  | 
**sid** | **string**|  | 
 **optional** | ***DeleteDocumentOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeleteDocumentOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **ifMatch** | **optional.String**| The If-Match HTTP request header | 

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

> DeleteDocumentPermission(ctx, serviceSid, documentSid, identity)



Delete a specific Sync Document Permission.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**|  | 
**documentSid** | **string**| Identifier of the Sync Document. Either a SID or a unique name. | 
**identity** | **string**| Arbitrary string identifier representing a user associated with an FPA token, assigned by the developer. | 

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

> DeleteField(ctx, assistantSid, taskSid, sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assistantSid** | **string**| The unique ID of the Assistant. | 
**taskSid** | **string**| The unique ID of the Task associated with this Field. | 
**sid** | **string**| A 34 character string that uniquely identifies this resource. | 

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

> DeleteFieldType(ctx, assistantSid, sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assistantSid** | **string**|  | 
**sid** | **string**|  | 

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

> DeleteFieldValue(ctx, assistantSid, fieldTypeSid, sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assistantSid** | **string**|  | 
**fieldTypeSid** | **string**|  | 
**sid** | **string**|  | 

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

> DeleteFleet(ctx, sid)



Delete a specific Fleet from your account, also destroys all nested resources: Devices, Deployments, Certificates, Keys.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sid** | **string**| Provides a 34 character string that uniquely identifies the requested Fleet resource. | 

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

> DeleteHostedNumberOrder(ctx, sid)



Cancel the HostedNumberOrder (only available when the status is in `received`).

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sid** | **string**| A 34 character string that uniquely identifies this HostedNumberOrder. | 

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

> DeleteInstalledAddOn(ctx, sid)



Remove an Add-on installation from your account

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sid** | **string**| The SID of the InstalledAddOn resource to delete. | 

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

> DeleteJob(ctx, jobSid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobSid** | **string**| The unique string that that we created to identify the Bulk Export job | 

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

> DeleteKey(ctx, fleetSid, sid)



Delete a specific Key credential from the Fleet, effectively disallowing any inbound client connections that are presenting it.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fleetSid** | **string**|  | 
**sid** | **string**| Provides a 34 character string that uniquely identifies the requested Key credential resource. | 

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

> DeleteModelBuild(ctx, assistantSid, sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assistantSid** | **string**|  | 
**sid** | **string**|  | 

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

> DeleteQuery(ctx, assistantSid, sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assistantSid** | **string**| The unique ID of the Assistant. | 
**sid** | **string**| A 34 character string that uniquely identifies this resource. | 

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

> DeleteRatePlan(ctx, sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sid** | **string**|  | 

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

> DeleteSample(ctx, assistantSid, taskSid, sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assistantSid** | **string**| The unique ID of the Assistant. | 
**taskSid** | **string**| The unique ID of the Task associated with this Sample. | 
**sid** | **string**| A 34 character string that uniquely identifies this resource. | 

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

> DeleteService(ctx, sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sid** | **string**|  | 

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

> DeleteSyncList(ctx, serviceSid, sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**|  | 
**sid** | **string**|  | 

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

> DeleteSyncListItem(ctx, serviceSid, listSid, index, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**|  | 
**listSid** | **string**|  | 
**index** | **int32**|  | 
 **optional** | ***DeleteSyncListItemOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeleteSyncListItemOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **ifMatch** | **optional.String**| The If-Match HTTP request header | 

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

> DeleteSyncListPermission(ctx, serviceSid, listSid, identity)



Delete a specific Sync List Permission.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**|  | 
**listSid** | **string**| Identifier of the Sync List. Either a SID or a unique name. | 
**identity** | **string**| Arbitrary string identifier representing a user associated with an FPA token, assigned by the developer. | 

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

> DeleteSyncMap(ctx, serviceSid, sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**|  | 
**sid** | **string**|  | 

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

> DeleteSyncMapItem(ctx, serviceSid, mapSid, key, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**|  | 
**mapSid** | **string**|  | 
**key** | **string**|  | 
 **optional** | ***DeleteSyncMapItemOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeleteSyncMapItemOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **ifMatch** | **optional.String**| The If-Match HTTP request header | 

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

> DeleteSyncMapPermission(ctx, serviceSid, mapSid, identity)



Delete a specific Sync Map Permission.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**|  | 
**mapSid** | **string**| Identifier of the Sync Map. Either a SID or a unique name. | 
**identity** | **string**| Arbitrary string identifier representing a user associated with an FPA token, assigned by the developer. | 

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

> DeleteTask(ctx, assistantSid, sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assistantSid** | **string**| The unique ID of the Assistant. | 
**sid** | **string**| A 34 character string that uniquely identifies this resource. | 

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

> PreviewUnderstandAssistant FetchAssistant(ctx, sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sid** | **string**| A 34 character string that uniquely identifies this resource. | 

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

> PreviewUnderstandAssistantAssistantFallbackActions FetchAssistantFallbackActions(ctx, assistantSid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assistantSid** | **string**|  | 

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

> PreviewUnderstandAssistantAssistantInitiationActions FetchAssistantInitiationActions(ctx, assistantSid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assistantSid** | **string**|  | 

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

> PreviewHostedNumbersAuthorizationDocument FetchAuthorizationDocument(ctx, sid)



Fetch a specific AuthorizationDocument.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sid** | **string**| A 34 character string that uniquely identifies this AuthorizationDocument. | 

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

> PreviewMarketplaceAvailableAddOn FetchAvailableAddOn(ctx, sid)



Fetch an instance of an Add-on currently available to be installed.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sid** | **string**| The SID of the AvailableAddOn resource to fetch. | 

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

> PreviewMarketplaceAvailableAddOnAvailableAddOnExtension FetchAvailableAddOnExtension(ctx, availableAddOnSid, sid)



Fetch an instance of an Extension for the Available Add-on.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**availableAddOnSid** | **string**| The SID of the AvailableAddOn resource with the extension to fetch. | 
**sid** | **string**| The SID of the AvailableAddOn Extension resource to fetch. | 

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

> PreviewTrustedCommsBrandedChannel FetchBrandedChannel(ctx, sid)



Fetch a specific Branded Channel.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sid** | **string**| The unique SID identifier of the Branded Channel. | 

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
 **ifNoneMatch** | **optional.String**| The If-None-Match HTTP request header | 

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

> PreviewDeployedDevicesFleetCertificate FetchCertificate(ctx, fleetSid, sid)



Fetch information about a specific Certificate credential in the Fleet.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fleetSid** | **string**|  | 
**sid** | **string**| Provides a 34 character string that uniquely identifies the requested Certificate credential resource. | 

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

> PreviewWirelessCommand FetchCommand(ctx, sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sid** | **string**|  | 

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
 **xXcnamSensitivePhoneNumber** | **optional.String**| The X-Xcnam-Sensitive-Phone-Number HTTP request header | 

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
 **xXcnamSensitivePhoneNumberFrom** | **optional.String**| The X-Xcnam-Sensitive-Phone-Number-From HTTP request header | 
 **xXcnamSensitivePhoneNumberTo** | **optional.String**| The X-Xcnam-Sensitive-Phone-Number-To HTTP request header | 

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

> FetchDay(ctx, resourceType, day)



Fetch a specific Day.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourceType** | **string**| The type of communication – Messages, Calls | 
**day** | **string**| The ISO 8601 format date of the resources in the file, for a UTC day | 

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

> PreviewDeployedDevicesFleetDeployment FetchDeployment(ctx, fleetSid, sid)



Fetch information about a specific Deployment in the Fleet.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fleetSid** | **string**|  | 
**sid** | **string**| Provides a 34 character string that uniquely identifies the requested Deployment resource. | 

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

> PreviewDeployedDevicesFleetDevice FetchDevice(ctx, fleetSid, sid)



Fetch information about a specific Device in the Fleet.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fleetSid** | **string**|  | 
**sid** | **string**| Provides a 34 character string that uniquely identifies the requested Device resource. | 

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

> PreviewUnderstandAssistantDialogue FetchDialogue(ctx, assistantSid, sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assistantSid** | **string**|  | 
**sid** | **string**|  | 

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

> PreviewSyncServiceDocument FetchDocument(ctx, serviceSid, sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**|  | 
**sid** | **string**|  | 

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

> PreviewSyncServiceDocumentDocumentPermission FetchDocumentPermission(ctx, serviceSid, documentSid, identity)



Fetch a specific Sync Document Permission.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**|  | 
**documentSid** | **string**| Identifier of the Sync Document. Either a SID or a unique name. | 
**identity** | **string**| Arbitrary string identifier representing a user associated with an FPA token, assigned by the developer. | 

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

> PreviewBulkExportsExport FetchExport(ctx, resourceType)



Fetch a specific Export.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourceType** | **string**| The type of communication – Messages, Calls | 

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

> PreviewBulkExportsExportConfiguration FetchExportConfiguration(ctx, resourceType)



Fetch a specific Export Configuration.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourceType** | **string**| The type of communication – Messages, Calls | 

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

> PreviewUnderstandAssistantTaskField FetchField(ctx, assistantSid, taskSid, sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assistantSid** | **string**| The unique ID of the Assistant. | 
**taskSid** | **string**| The unique ID of the Task associated with this Field. | 
**sid** | **string**| A 34 character string that uniquely identifies this resource. | 

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

> PreviewUnderstandAssistantFieldType FetchFieldType(ctx, assistantSid, sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assistantSid** | **string**|  | 
**sid** | **string**|  | 

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

> PreviewUnderstandAssistantFieldTypeFieldValue FetchFieldValue(ctx, assistantSid, fieldTypeSid, sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assistantSid** | **string**|  | 
**fieldTypeSid** | **string**|  | 
**sid** | **string**|  | 

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

> PreviewDeployedDevicesFleet FetchFleet(ctx, sid)



Fetch information about a specific Fleet in your account.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sid** | **string**| Provides a 34 character string that uniquely identifies the requested Fleet resource. | 

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

> PreviewHostedNumbersHostedNumberOrder FetchHostedNumberOrder(ctx, sid)



Fetch a specific HostedNumberOrder.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sid** | **string**| A 34 character string that uniquely identifies this HostedNumberOrder. | 

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

> PreviewMarketplaceInstalledAddOn FetchInstalledAddOn(ctx, sid)



Fetch an instance of an Add-on currently installed on this Account.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sid** | **string**| The SID of the InstalledAddOn resource to fetch. | 

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

> PreviewMarketplaceInstalledAddOnInstalledAddOnExtension FetchInstalledAddOnExtension(ctx, installedAddOnSid, sid)



Fetch an instance of an Extension for the Installed Add-on.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**installedAddOnSid** | **string**| The SID of the InstalledAddOn resource with the extension to fetch. | 
**sid** | **string**| The SID of the InstalledAddOn Extension resource to fetch. | 

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

> PreviewBulkExportsExportJob FetchJob(ctx, jobSid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobSid** | **string**|  | 

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

> PreviewDeployedDevicesFleetKey FetchKey(ctx, fleetSid, sid)



Fetch information about a specific Key credential in the Fleet.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fleetSid** | **string**|  | 
**sid** | **string**| Provides a 34 character string that uniquely identifies the requested Key credential resource. | 

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

> PreviewUnderstandAssistantModelBuild FetchModelBuild(ctx, assistantSid, sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assistantSid** | **string**|  | 
**sid** | **string**|  | 

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

> PreviewUnderstandAssistantQuery FetchQuery(ctx, assistantSid, sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assistantSid** | **string**| The unique ID of the Assistant. | 
**sid** | **string**| A 34 character string that uniquely identifies this resource. | 

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

> PreviewWirelessRatePlan FetchRatePlan(ctx, sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sid** | **string**|  | 

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

> PreviewUnderstandAssistantTaskSample FetchSample(ctx, assistantSid, taskSid, sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assistantSid** | **string**| The unique ID of the Assistant. | 
**taskSid** | **string**| The unique ID of the Task associated with this Sample. | 
**sid** | **string**| A 34 character string that uniquely identifies this resource. | 

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

> PreviewSyncService FetchService(ctx, sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sid** | **string**|  | 

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

> PreviewWirelessSim FetchSim(ctx, sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sid** | **string**|  | 

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

> PreviewUnderstandAssistantStyleSheet FetchStyleSheet(ctx, assistantSid)



Returns Style sheet JSON object for this Assistant

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assistantSid** | **string**| The unique ID of the Assistant | 

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

> PreviewSyncServiceSyncList FetchSyncList(ctx, serviceSid, sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**|  | 
**sid** | **string**|  | 

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

> PreviewSyncServiceSyncListSyncListItem FetchSyncListItem(ctx, serviceSid, listSid, index)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**|  | 
**listSid** | **string**|  | 
**index** | **int32**|  | 

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

> PreviewSyncServiceSyncListSyncListPermission FetchSyncListPermission(ctx, serviceSid, listSid, identity)



Fetch a specific Sync List Permission.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**|  | 
**listSid** | **string**| Identifier of the Sync List. Either a SID or a unique name. | 
**identity** | **string**| Arbitrary string identifier representing a user associated with an FPA token, assigned by the developer. | 

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

> PreviewSyncServiceSyncMap FetchSyncMap(ctx, serviceSid, sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**|  | 
**sid** | **string**|  | 

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

> PreviewSyncServiceSyncMapSyncMapItem FetchSyncMapItem(ctx, serviceSid, mapSid, key)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**|  | 
**mapSid** | **string**|  | 
**key** | **string**|  | 

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

> PreviewSyncServiceSyncMapSyncMapPermission FetchSyncMapPermission(ctx, serviceSid, mapSid, identity)



Fetch a specific Sync Map Permission.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**|  | 
**mapSid** | **string**| Identifier of the Sync Map. Either a SID or a unique name. | 
**identity** | **string**| Arbitrary string identifier representing a user associated with an FPA token, assigned by the developer. | 

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

> PreviewUnderstandAssistantTask FetchTask(ctx, assistantSid, sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assistantSid** | **string**| The unique ID of the Assistant. | 
**sid** | **string**| A 34 character string that uniquely identifies this resource. | 

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

> PreviewUnderstandAssistantTaskTaskActions FetchTaskActions(ctx, assistantSid, taskSid)



Returns JSON actions for this Task.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assistantSid** | **string**| The unique ID of the parent Assistant. | 
**taskSid** | **string**| The unique ID of the Task. | 

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

> PreviewUnderstandAssistantTaskTaskStatistics FetchTaskStatistics(ctx, assistantSid, taskSid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assistantSid** | **string**| The unique ID of the parent Assistant. | 
**taskSid** | **string**| The unique ID of the Task associated with this Field. | 

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

> PreviewWirelessSimUsage FetchUsage(ctx, simSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**simSid** | **string**|  | 
 **optional** | ***FetchUsageOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a FetchUsageOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **end** | **optional.String**|  | 
 **start** | **optional.String**|  | 

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
 **pageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

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
 **email** | **optional.String**| Email that this AuthorizationDocument will be sent to for signing. | 
 **status** | **optional.String**| Status of an instance resource. It can hold one of the values: 1. opened 2. signing, 3. signed LOA, 4. canceled, 5. failed. See the section entitled [Status Values](https://www.twilio.com/docs/api/phone-numbers/hosted-number-authorization-documents#status-values) for more information on each of these statuses. | 
 **pageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

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
 **pageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

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

> PreviewMarketplaceAvailableAddOnAvailableAddOnExtensionReadResponse ListAvailableAddOnExtension(ctx, availableAddOnSid, optional)



Retrieve a list of Extensions for the Available Add-on.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**availableAddOnSid** | **string**| The SID of the AvailableAddOn resource with the extensions to read. | 
 **optional** | ***ListAvailableAddOnExtensionOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListAvailableAddOnExtensionOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

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

> PreviewDeployedDevicesFleetCertificateReadResponse ListCertificate(ctx, fleetSid, optional)



Retrieve a list of all Certificate credentials belonging to the Fleet.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fleetSid** | **string**|  | 
 **optional** | ***ListCertificateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListCertificateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deviceSid** | **optional.String**| Filters the resulting list of Certificates by a unique string identifier of an authenticated Device. | 
 **pageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

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
 **device** | **optional.String**|  | 
 **sim** | **optional.String**|  | 
 **status** | **optional.String**|  | 
 **direction** | **optional.String**|  | 
 **pageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

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

> PreviewBulkExportsExportDayReadResponse ListDay(ctx, resourceType, optional)



Retrieve a list of all Days for a resource.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourceType** | **string**| The type of communication – Messages, Calls | 
 **optional** | ***ListDayOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListDayOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

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

> PreviewHostedNumbersAuthorizationDocumentDependentHostedNumberOrderReadResponse ListDependentHostedNumberOrder(ctx, signingDocumentSid, optional)



Retrieve a list of dependent HostedNumberOrders belonging to the AuthorizationDocument.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**signingDocumentSid** | **string**|  | 
 **optional** | ***ListDependentHostedNumberOrderOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListDependentHostedNumberOrderOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **status** | **optional.String**| Status of an instance resource. It can hold one of the values: 1. opened 2. signing, 3. signed LOA, 4. canceled, 5. failed. See the section entitled [Status Values](https://www.twilio.com/docs/api/phone-numbers/hosted-number-authorization-documents#status-values) for more information on each of these statuses. | 
 **phoneNumber** | **optional.String**| An E164 formatted phone number hosted by this HostedNumberOrder. | 
 **incomingPhoneNumberSid** | **optional.String**| A 34 character string that uniquely identifies the IncomingPhoneNumber resource created by this HostedNumberOrder. | 
 **friendlyName** | **optional.String**| A human readable description of this resource, up to 64 characters. | 
 **uniqueName** | **optional.String**| Provides a unique and addressable name to be assigned to this HostedNumberOrder, assigned by the developer, to be optionally used in addition to SID. | 
 **pageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

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

> PreviewDeployedDevicesFleetDeploymentReadResponse ListDeployment(ctx, fleetSid, optional)



Retrieve a list of all Deployments belonging to the Fleet.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fleetSid** | **string**|  | 
 **optional** | ***ListDeploymentOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListDeploymentOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

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

> PreviewDeployedDevicesFleetDeviceReadResponse ListDevice(ctx, fleetSid, optional)



Retrieve a list of all Devices belonging to the Fleet.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fleetSid** | **string**|  | 
 **optional** | ***ListDeviceOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListDeviceOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deploymentSid** | **optional.String**| Filters the resulting list of Devices by a unique string identifier of the Deployment they are associated with. | 
 **pageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

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

> PreviewSyncServiceDocumentReadResponse ListDocument(ctx, serviceSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**|  | 
 **optional** | ***ListDocumentOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListDocumentOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

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

> PreviewSyncServiceDocumentDocumentPermissionReadResponse ListDocumentPermission(ctx, serviceSid, documentSid, optional)



Retrieve a list of all Permissions applying to a Sync Document.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**|  | 
**documentSid** | **string**| Identifier of the Sync Document. Either a SID or a unique name. | 
 **optional** | ***ListDocumentPermissionOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListDocumentPermissionOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

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

> PreviewBulkExportsExportExportCustomJobReadResponse ListExportCustomJob(ctx, resourceType, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourceType** | **string**| The type of communication – Messages, Calls | 
 **optional** | ***ListExportCustomJobOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListExportCustomJobOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

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

> PreviewUnderstandAssistantTaskFieldReadResponse ListField(ctx, assistantSid, taskSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assistantSid** | **string**| The unique ID of the Assistant. | 
**taskSid** | **string**| The unique ID of the Task associated with this Field. | 
 **optional** | ***ListFieldOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListFieldOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

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

> PreviewUnderstandAssistantFieldTypeReadResponse ListFieldType(ctx, assistantSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assistantSid** | **string**|  | 
 **optional** | ***ListFieldTypeOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListFieldTypeOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

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

> PreviewUnderstandAssistantFieldTypeFieldValueReadResponse ListFieldValue(ctx, assistantSid, fieldTypeSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assistantSid** | **string**|  | 
**fieldTypeSid** | **string**|  | 
 **optional** | ***ListFieldValueOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListFieldValueOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **language** | **optional.String**| An ISO language-country string of the value. For example: *en-US* | 
 **pageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

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
 **pageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

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
 **status** | **optional.String**| The Status of this HostedNumberOrder. One of &#x60;received&#x60;, &#x60;pending-verification&#x60;, &#x60;verified&#x60;, &#x60;pending-loa&#x60;, &#x60;carrier-processing&#x60;, &#x60;testing&#x60;, &#x60;completed&#x60;, &#x60;failed&#x60;, or &#x60;action-required&#x60;. | 
 **phoneNumber** | **optional.String**| An E164 formatted phone number hosted by this HostedNumberOrder. | 
 **incomingPhoneNumberSid** | **optional.String**| A 34 character string that uniquely identifies the IncomingPhoneNumber resource created by this HostedNumberOrder. | 
 **friendlyName** | **optional.String**| A human readable description of this resource, up to 64 characters. | 
 **uniqueName** | **optional.String**| Provides a unique and addressable name to be assigned to this HostedNumberOrder, assigned by the developer, to be optionally used in addition to SID. | 
 **pageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

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
 **pageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

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

> PreviewMarketplaceInstalledAddOnInstalledAddOnExtensionReadResponse ListInstalledAddOnExtension(ctx, installedAddOnSid, optional)



Retrieve a list of Extensions for the Installed Add-on.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**installedAddOnSid** | **string**| The SID of the InstalledAddOn resource with the extensions to read. | 
 **optional** | ***ListInstalledAddOnExtensionOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListInstalledAddOnExtensionOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

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

> PreviewDeployedDevicesFleetKeyReadResponse ListKey(ctx, fleetSid, optional)



Retrieve a list of all Keys credentials belonging to the Fleet.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fleetSid** | **string**|  | 
 **optional** | ***ListKeyOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListKeyOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deviceSid** | **optional.String**| Filters the resulting list of Keys by a unique string identifier of an authenticated Device. | 
 **pageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

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

> PreviewUnderstandAssistantModelBuildReadResponse ListModelBuild(ctx, assistantSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assistantSid** | **string**|  | 
 **optional** | ***ListModelBuildOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListModelBuildOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

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

> PreviewUnderstandAssistantQueryReadResponse ListQuery(ctx, assistantSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assistantSid** | **string**| The unique ID of the parent Assistant. | 
 **optional** | ***ListQueryOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListQueryOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **language** | **optional.String**| An ISO language-country string of the sample. | 
 **modelBuild** | **optional.String**| The Model Build Sid or unique name of the Model Build to be queried. | 
 **status** | **optional.String**| A string that described the query status. The values can be: pending_review, reviewed, discarded | 
 **pageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

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
 **pageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

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

> PreviewUnderstandAssistantTaskSampleReadResponse ListSample(ctx, assistantSid, taskSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assistantSid** | **string**| The unique ID of the Assistant. | 
**taskSid** | **string**| The unique ID of the Task associated with this Sample. | 
 **optional** | ***ListSampleOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListSampleOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **language** | **optional.String**| An ISO language-country string of the sample. | 
 **pageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

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
 **pageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

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
 **status** | **optional.String**|  | 
 **iccid** | **optional.String**|  | 
 **ratePlan** | **optional.String**|  | 
 **eId** | **optional.String**|  | 
 **simRegistrationCode** | **optional.String**|  | 
 **pageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

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

> PreviewSyncServiceSyncListReadResponse ListSyncList(ctx, serviceSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**|  | 
 **optional** | ***ListSyncListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListSyncListOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

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

> PreviewSyncServiceSyncListSyncListItemReadResponse ListSyncListItem(ctx, serviceSid, listSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**|  | 
**listSid** | **string**|  | 
 **optional** | ***ListSyncListItemOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListSyncListItemOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **order** | **optional.String**|  | 
 **from** | **optional.String**|  | 
 **bounds** | **optional.String**|  | 
 **pageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

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

> PreviewSyncServiceSyncListSyncListPermissionReadResponse ListSyncListPermission(ctx, serviceSid, listSid, optional)



Retrieve a list of all Permissions applying to a Sync List.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**|  | 
**listSid** | **string**| Identifier of the Sync List. Either a SID or a unique name. | 
 **optional** | ***ListSyncListPermissionOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListSyncListPermissionOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

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

> PreviewSyncServiceSyncMapReadResponse ListSyncMap(ctx, serviceSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**|  | 
 **optional** | ***ListSyncMapOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListSyncMapOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

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

> PreviewSyncServiceSyncMapSyncMapItemReadResponse ListSyncMapItem(ctx, serviceSid, mapSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**|  | 
**mapSid** | **string**|  | 
 **optional** | ***ListSyncMapItemOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListSyncMapItemOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **order** | **optional.String**|  | 
 **from** | **optional.String**|  | 
 **bounds** | **optional.String**|  | 
 **pageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

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

> PreviewSyncServiceSyncMapSyncMapPermissionReadResponse ListSyncMapPermission(ctx, serviceSid, mapSid, optional)



Retrieve a list of all Permissions applying to a Sync Map.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**|  | 
**mapSid** | **string**| Identifier of the Sync Map. Either a SID or a unique name. | 
 **optional** | ***ListSyncMapPermissionOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListSyncMapPermissionOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

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

> PreviewUnderstandAssistantTaskReadResponse ListTask(ctx, assistantSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assistantSid** | **string**| The unique ID of the Assistant. | 
 **optional** | ***ListTaskOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListTaskOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

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

> PreviewUnderstandAssistant UpdateAssistant(ctx, sid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sid** | **string**| A 34 character string that uniquely identifies this resource. | 
 **optional** | ***UpdateAssistantOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateAssistantOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **callbackEvents** | **optional.String**| Space-separated list of callback events that will trigger callbacks. | 
 **callbackUrl** | **optional.String**| A user-provided URL to send event callbacks to. | 
 **fallbackActions** | [**optional.Interface of map[string]interface{}**](map[string]interface{}.md)| The JSON actions to be executed when the user&#39;s input is not recognized as matching any Task. | 
 **friendlyName** | **optional.String**| A text description for the Assistant. It is non-unique and can up to 255 characters long. | 
 **initiationActions** | [**optional.Interface of map[string]interface{}**](map[string]interface{}.md)| The JSON actions to be executed on inbound phone calls when the Assistant has to say something first. | 
 **logQueries** | **optional.Bool**| A boolean that specifies whether queries should be logged for 30 days further training. If false, no queries will be stored, if true, queries will be stored for 30 days and deleted thereafter. Defaults to true if no value is provided. | 
 **styleSheet** | [**optional.Interface of map[string]interface{}**](map[string]interface{}.md)| The JSON object that holds the style sheet for the assistant | 
 **uniqueName** | **optional.String**| A user-provided string that uniquely identifies this resource as an alternative to the sid. Unique up to 64 characters long. | 

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

> PreviewUnderstandAssistantAssistantFallbackActions UpdateAssistantFallbackActions(ctx, assistantSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assistantSid** | **string**|  | 
 **optional** | ***UpdateAssistantFallbackActionsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateAssistantFallbackActionsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fallbackActions** | [**optional.Interface of map[string]interface{}**](map[string]interface{}.md)|  | 

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

> PreviewUnderstandAssistantAssistantInitiationActions UpdateAssistantInitiationActions(ctx, assistantSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assistantSid** | **string**|  | 
 **optional** | ***UpdateAssistantInitiationActionsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateAssistantInitiationActionsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **initiationActions** | [**optional.Interface of map[string]interface{}**](map[string]interface{}.md)|  | 

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

> PreviewHostedNumbersAuthorizationDocument UpdateAuthorizationDocument(ctx, sid, optional)



Updates a specific AuthorizationDocument.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sid** | **string**|  | 
 **optional** | ***UpdateAuthorizationDocumentOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateAuthorizationDocumentOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **addressSid** | **optional.String**| A 34 character string that uniquely identifies the Address resource that is associated with this AuthorizationDocument. | 
 **ccEmails** | [**optional.Interface of []string**](string.md)| Email recipients who will be informed when an Authorization Document has been sent and signed | 
 **contactPhoneNumber** | **optional.String**| The contact phone number of the person authorized to sign the Authorization Document. | 
 **contactTitle** | **optional.String**| The title of the person authorized to sign the Authorization Document for this phone number. | 
 **email** | **optional.String**| Email that this AuthorizationDocument will be sent to for signing. | 
 **hostedNumberOrderSids** | [**optional.Interface of []string**](string.md)| A list of HostedNumberOrder sids that this AuthorizationDocument will authorize for hosting phone number capabilities on Twilio&#39;s platform. | 
 **status** | **optional.String**| Status of an instance resource. It can hold one of the values: 1. opened 2. signing, 3. signed LOA, 4. canceled, 5. failed. See the section entitled [Status Values](https://www.twilio.com/docs/api/phone-numbers/hosted-number-authorization-documents#status-values) for more information on each of these statuses. | 

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

> PreviewDeployedDevicesFleetCertificate UpdateCertificate(ctx, fleetSid, sid, optional)



Update the given properties of a specific Certificate credential in the Fleet, giving it a friendly name or assigning to a Device.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fleetSid** | **string**|  | 
**sid** | **string**| Provides a 34 character string that uniquely identifies the requested Certificate credential resource. | 
 **optional** | ***UpdateCertificateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateCertificateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **deviceSid** | **optional.String**| Provides the unique string identifier of an existing Device to become authenticated with this Certificate credential. | 
 **friendlyName** | **optional.String**| Provides a human readable descriptive text for this Certificate credential, up to 256 characters long. | 

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

> PreviewDeployedDevicesFleetDeployment UpdateDeployment(ctx, fleetSid, sid, optional)



Update the given properties of a specific Deployment credential in the Fleet, giving it a friendly name or linking to a specific Twilio Sync service instance.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fleetSid** | **string**|  | 
**sid** | **string**| Provides a 34 character string that uniquely identifies the requested Deployment resource. | 
 **optional** | ***UpdateDeploymentOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateDeploymentOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **friendlyName** | **optional.String**| Provides a human readable descriptive text for this Deployment, up to 64 characters long | 
 **syncServiceSid** | **optional.String**| Provides the unique string identifier of the Twilio Sync service instance that will be linked to and accessible by this Deployment. | 

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

> PreviewDeployedDevicesFleetDevice UpdateDevice(ctx, fleetSid, sid, optional)



Update the given properties of a specific Device in the Fleet, giving it a friendly name, assigning to a Deployment, or a human identity.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fleetSid** | **string**|  | 
**sid** | **string**| Provides a 34 character string that uniquely identifies the requested Device resource. | 
 **optional** | ***UpdateDeviceOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateDeviceOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **deploymentSid** | **optional.String**| Specifies the unique string identifier of the Deployment group that this Device is going to be associated with. | 
 **enabled** | **optional.Bool**|  | 
 **friendlyName** | **optional.String**| Provides a human readable descriptive text to be assigned to this Device, up to 256 characters long. | 
 **identity** | **optional.String**| Provides an arbitrary string identifier representing a human user to be associated with this Device, up to 256 characters long. | 

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

> PreviewSyncServiceDocument UpdateDocument(ctx, serviceSid, sid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**|  | 
**sid** | **string**|  | 
 **optional** | ***UpdateDocumentOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateDocumentOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **ifMatch** | **optional.String**| The If-Match HTTP request header | 
 **data** | [**optional.Interface of map[string]interface{}**](map[string]interface{}.md)|  | 

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

> PreviewSyncServiceDocumentDocumentPermission UpdateDocumentPermission(ctx, serviceSid, documentSid, identity, optional)



Update an identity's access to a specific Sync Document.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**| The unique SID identifier of the Sync Service Instance. | 
**documentSid** | **string**| Identifier of the Sync Document. Either a SID or a unique name. | 
**identity** | **string**| Arbitrary string identifier representing a human user associated with an FPA token, assigned by the developer. | 
 **optional** | ***UpdateDocumentPermissionOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateDocumentPermissionOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **manage** | **optional.Bool**| Boolean flag specifying whether the identity can delete the Sync Document. | 
 **read** | **optional.Bool**| Boolean flag specifying whether the identity can read the Sync Document. | 
 **write** | **optional.Bool**| Boolean flag specifying whether the identity can update the Sync Document. | 

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

> PreviewBulkExportsExportConfiguration UpdateExportConfiguration(ctx, resourceType, optional)



Update a specific Export Configuration.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourceType** | **string**| The type of communication – Messages, Calls | 
 **optional** | ***UpdateExportConfigurationOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateExportConfigurationOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **enabled** | **optional.Bool**| If true, Twilio will automatically generate every day&#39;s file when the day is over. | 
 **webhookMethod** | **optional.String**| Sets whether Twilio should call a webhook URL when the automatic generation is complete, using GET or POST. The actual destination is set in the webhook_url | 
 **webhookUrl** | **optional.String**| Stores the URL destination for the method specified in webhook_method. | 

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

> PreviewUnderstandAssistantFieldType UpdateFieldType(ctx, assistantSid, sid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assistantSid** | **string**|  | 
**sid** | **string**|  | 
 **optional** | ***UpdateFieldTypeOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateFieldTypeOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **friendlyName** | **optional.String**| A user-provided string that identifies this resource. It is non-unique and can up to 255 characters long. | 
 **uniqueName** | **optional.String**| A user-provided string that uniquely identifies this resource as an alternative to the sid. Unique up to 64 characters long. | 

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

> PreviewDeployedDevicesFleet UpdateFleet(ctx, sid, optional)



Update the friendly name property of a specific Fleet in your account.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sid** | **string**| Provides a 34 character string that uniquely identifies the requested Fleet resource. | 
 **optional** | ***UpdateFleetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateFleetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **defaultDeploymentSid** | **optional.String**| Provides a string identifier of a Deployment that is going to be used as a default one for this Fleet. | 
 **friendlyName** | **optional.String**| Provides a human readable descriptive text for this Fleet, up to 256 characters long. | 

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

> PreviewHostedNumbersHostedNumberOrder UpdateHostedNumberOrder(ctx, sid, optional)



Updates a specific HostedNumberOrder.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sid** | **string**|  | 
 **optional** | ***UpdateHostedNumberOrderOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateHostedNumberOrderOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **callDelay** | **optional.Int32**| The number of seconds, between 0 and 60, to delay before initiating the verification call. Defaults to 0. | 
 **ccEmails** | [**optional.Interface of []string**](string.md)| Optional. A list of emails that LOA document for this HostedNumberOrder will be carbon copied to. | 
 **email** | **optional.String**| Email of the owner of this phone number that is being hosted. | 
 **extension** | **optional.String**| Digits to dial after connecting the verification call. | 
 **friendlyName** | **optional.String**| A 64 character string that is a human readable text that describes this resource. | 
 **status** | **optional.String**| User can only post to &#x60;pending-verification&#x60; status to transition the HostedNumberOrder to initiate a verification call or verification of ownership with a copy of a phone bill. | 
 **uniqueName** | **optional.String**| Provides a unique and addressable name to be assigned to this HostedNumberOrder, assigned by the developer, to be optionally used in addition to SID. | 
 **verificationCode** | **optional.String**| A verification code that is given to the user via a phone call to the phone number that is being hosted. | 
 **verificationDocumentSid** | **optional.String**| Optional. The unique sid identifier of the Identity Document that represents the document for verifying ownership of the number to be hosted. Required when VerificationType is phone-bill. | 
 **verificationType** | **optional.String**| Optional. The method used for verifying ownership of the number to be hosted. One of phone-call (default) or phone-bill. | 

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

> PreviewMarketplaceInstalledAddOn UpdateInstalledAddOn(ctx, sid, optional)



Update an Add-on installation for the Account specified.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sid** | **string**| The SID of the InstalledAddOn resource to update. | 
 **optional** | ***UpdateInstalledAddOnOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateInstalledAddOnOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **configuration** | [**optional.Interface of map[string]interface{}**](map[string]interface{}.md)| Valid JSON object that conform to the configuration schema exposed by the associated AvailableAddOn resource. This is only required by Add-ons that need to be configured | 
 **uniqueName** | **optional.String**| An application-defined string that uniquely identifies the resource. This value must be unique within the Account. | 

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

> PreviewMarketplaceInstalledAddOnInstalledAddOnExtension UpdateInstalledAddOnExtension(ctx, installedAddOnSid, sid, optional)



Update an Extension for an Add-on installation.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**installedAddOnSid** | **string**| The SID of the InstalledAddOn resource with the extension to update. | 
**sid** | **string**| The SID of the InstalledAddOn Extension resource to update. | 
 **optional** | ***UpdateInstalledAddOnExtensionOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateInstalledAddOnExtensionOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **enabled** | **optional.Bool**| Whether the Extension should be invoked. | 

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

> PreviewDeployedDevicesFleetKey UpdateKey(ctx, fleetSid, sid, optional)



Update the given properties of a specific Key credential in the Fleet, giving it a friendly name or assigning to a Device.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fleetSid** | **string**|  | 
**sid** | **string**| Provides a 34 character string that uniquely identifies the requested Key credential resource. | 
 **optional** | ***UpdateKeyOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateKeyOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **deviceSid** | **optional.String**| Provides the unique string identifier of an existing Device to become authenticated with this Key credential. | 
 **friendlyName** | **optional.String**| Provides a human readable descriptive text for this Key credential, up to 256 characters long. | 

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

> PreviewUnderstandAssistantModelBuild UpdateModelBuild(ctx, assistantSid, sid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assistantSid** | **string**|  | 
**sid** | **string**|  | 
 **optional** | ***UpdateModelBuildOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateModelBuildOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **uniqueName** | **optional.String**| A user-provided string that uniquely identifies this resource as an alternative to the sid. Unique up to 64 characters long. For example: v0.1 | 

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

> PreviewUnderstandAssistantQuery UpdateQuery(ctx, assistantSid, sid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assistantSid** | **string**| The unique ID of the parent Assistant. | 
**sid** | **string**| A 34 character string that uniquely identifies this resource. | 
 **optional** | ***UpdateQueryOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateQueryOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **sampleSid** | **optional.String**| An optional reference to the Sample created from this query. | 
 **status** | **optional.String**| A string that described the query status. The values can be: pending_review, reviewed, discarded | 

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

> PreviewWirelessRatePlan UpdateRatePlan(ctx, sid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sid** | **string**|  | 
 **optional** | ***UpdateRatePlanOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateRatePlanOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **friendlyName** | **optional.String**|  | 
 **uniqueName** | **optional.String**|  | 

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

> PreviewUnderstandAssistantTaskSample UpdateSample(ctx, assistantSid, taskSid, sid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assistantSid** | **string**| The unique ID of the Assistant. | 
**taskSid** | **string**| The unique ID of the Task associated with this Sample. | 
**sid** | **string**| A 34 character string that uniquely identifies this resource. | 
 **optional** | ***UpdateSampleOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateSampleOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **language** | **optional.String**| An ISO language-country string of the sample. | 
 **sourceChannel** | **optional.String**| The communication channel the sample was captured. It can be: *voice*, *sms*, *chat*, *alexa*, *google-assistant*, or *slack*. If not included the value will be null | 
 **taggedText** | **optional.String**| The text example of how end-users may express this task. The sample may contain Field tag blocks. | 

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

> PreviewSyncService UpdateService(ctx, sid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sid** | **string**|  | 
 **optional** | ***UpdateServiceOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateServiceOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **aclEnabled** | **optional.Bool**|  | 
 **friendlyName** | **optional.String**|  | 
 **reachabilityWebhooksEnabled** | **optional.Bool**|  | 
 **webhookUrl** | **optional.String**|  | 

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

> PreviewWirelessSim UpdateSim(ctx, sid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sid** | **string**|  | 
 **optional** | ***UpdateSimOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateSimOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **callbackMethod** | **optional.String**|  | 
 **callbackUrl** | **optional.String**|  | 
 **commandsCallbackMethod** | **optional.String**|  | 
 **commandsCallbackUrl** | **optional.String**|  | 
 **friendlyName** | **optional.String**|  | 
 **ratePlan** | **optional.String**|  | 
 **smsFallbackMethod** | **optional.String**|  | 
 **smsFallbackUrl** | **optional.String**|  | 
 **smsMethod** | **optional.String**|  | 
 **smsUrl** | **optional.String**|  | 
 **status** | **optional.String**|  | 
 **uniqueName** | **optional.String**|  | 
 **voiceFallbackMethod** | **optional.String**|  | 
 **voiceFallbackUrl** | **optional.String**|  | 
 **voiceMethod** | **optional.String**|  | 
 **voiceUrl** | **optional.String**|  | 

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

> PreviewUnderstandAssistantStyleSheet UpdateStyleSheet(ctx, assistantSid, optional)



Updates the style sheet for an assistant identified by {AssistantSid} or {AssistantUniqueName}.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assistantSid** | **string**| The unique ID of the Assistant | 
 **optional** | ***UpdateStyleSheetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateStyleSheetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **styleSheet** | [**optional.Interface of map[string]interface{}**](map[string]interface{}.md)| The JSON Style sheet string | 

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

> PreviewSyncServiceSyncListSyncListItem UpdateSyncListItem(ctx, serviceSid, listSid, index, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**|  | 
**listSid** | **string**|  | 
**index** | **int32**|  | 
 **optional** | ***UpdateSyncListItemOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateSyncListItemOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **ifMatch** | **optional.String**| The If-Match HTTP request header | 
 **data** | [**optional.Interface of map[string]interface{}**](map[string]interface{}.md)|  | 

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

> PreviewSyncServiceSyncListSyncListPermission UpdateSyncListPermission(ctx, serviceSid, listSid, identity, optional)



Update an identity's access to a specific Sync List.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**| The unique SID identifier of the Sync Service Instance. | 
**listSid** | **string**| Identifier of the Sync List. Either a SID or a unique name. | 
**identity** | **string**| Arbitrary string identifier representing a human user associated with an FPA token, assigned by the developer. | 
 **optional** | ***UpdateSyncListPermissionOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateSyncListPermissionOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **manage** | **optional.Bool**| Boolean flag specifying whether the identity can delete the Sync List. | 
 **read** | **optional.Bool**| Boolean flag specifying whether the identity can read the Sync List. | 
 **write** | **optional.Bool**| Boolean flag specifying whether the identity can create, update and delete Items of the Sync List. | 

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

> PreviewSyncServiceSyncMapSyncMapItem UpdateSyncMapItem(ctx, serviceSid, mapSid, key, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**|  | 
**mapSid** | **string**|  | 
**key** | **string**|  | 
 **optional** | ***UpdateSyncMapItemOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateSyncMapItemOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **ifMatch** | **optional.String**| The If-Match HTTP request header | 
 **data** | [**optional.Interface of map[string]interface{}**](map[string]interface{}.md)|  | 

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

> PreviewSyncServiceSyncMapSyncMapPermission UpdateSyncMapPermission(ctx, serviceSid, mapSid, identity, optional)



Update an identity's access to a specific Sync Map.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**| The unique SID identifier of the Sync Service Instance. | 
**mapSid** | **string**| Identifier of the Sync Map. Either a SID or a unique name. | 
**identity** | **string**| Arbitrary string identifier representing a human user associated with an FPA token, assigned by the developer. | 
 **optional** | ***UpdateSyncMapPermissionOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateSyncMapPermissionOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **manage** | **optional.Bool**| Boolean flag specifying whether the identity can delete the Sync Map. | 
 **read** | **optional.Bool**| Boolean flag specifying whether the identity can read the Sync Map. | 
 **write** | **optional.Bool**| Boolean flag specifying whether the identity can create, update and delete Items of the Sync Map. | 

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

> PreviewUnderstandAssistantTask UpdateTask(ctx, assistantSid, sid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assistantSid** | **string**| The unique ID of the Assistant. | 
**sid** | **string**| A 34 character string that uniquely identifies this resource. | 
 **optional** | ***UpdateTaskOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateTaskOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **actions** | [**optional.Interface of map[string]interface{}**](map[string]interface{}.md)| A user-provided JSON object encoded as a string to specify the actions for this task. It is optional and non-unique. | 
 **actionsUrl** | **optional.String**| User-provided HTTP endpoint where from the assistant fetches actions | 
 **friendlyName** | **optional.String**| A user-provided string that identifies this resource. It is non-unique and can up to 255 characters long. | 
 **uniqueName** | **optional.String**| A user-provided string that uniquely identifies this resource as an alternative to the sid. Unique up to 64 characters long. | 

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

> PreviewUnderstandAssistantTaskTaskActions UpdateTaskActions(ctx, assistantSid, taskSid, optional)



Updates the actions of an Task identified by {TaskSid} or {TaskUniqueName}.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assistantSid** | **string**| The unique ID of the parent Assistant. | 
**taskSid** | **string**| The unique ID of the Task. | 
 **optional** | ***UpdateTaskActionsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateTaskActionsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **actions** | [**optional.Interface of map[string]interface{}**](map[string]interface{}.md)| The JSON actions that instruct the Assistant how to perform this task. | 

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

