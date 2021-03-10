# DefaultApi

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



### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateAssistantParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**CallbackEvents** | **string** | Space-separated list of callback events that will trigger callbacks. | 
**CallbackUrl** | **string** | A user-provided URL to send event callbacks to. | 
**FallbackActions** | [**map[string]interface{}**](map[string]interface{}.md) | The JSON actions to be executed when the user&#39;s input is not recognized as matching any Task. | 
**FriendlyName** | **string** | A text description for the Assistant. It is non-unique and can up to 255 characters long. | 
**InitiationActions** | [**map[string]interface{}**](map[string]interface{}.md) | The JSON actions to be executed on inbound phone calls when the Assistant has to say something first. | 
**LogQueries** | **bool** | A boolean that specifies whether queries should be logged for 30 days further training. If false, no queries will be stored, if true, queries will be stored for 30 days and deleted thereafter. Defaults to true if no value is provided. | 
**StyleSheet** | [**map[string]interface{}**](map[string]interface{}.md) | The JSON object that holds the style sheet for the assistant | 
**UniqueName** | **string** | A user-provided string that uniquely identifies this resource as an alternative to the sid. Unique up to 64 characters long. | 

### Return type

[**PreviewUnderstandAssistant**](PreviewUnderstandAssistant.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded, 
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAuthorizationDocument

> PreviewHostedNumbersAuthorizationDocument CreateAuthorizationDocument(ctx, optional)



Create an AuthorizationDocument for authorizing the hosting of phone number capabilities on Twilio's platform.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateAuthorizationDocumentParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**AddressSid** | **string** | A 34 character string that uniquely identifies the Address resource that is associated with this AuthorizationDocument. | 
**CcEmails** | **[]string** | Email recipients who will be informed when an Authorization Document has been sent and signed. | 
**ContactPhoneNumber** | **string** | The contact phone number of the person authorized to sign the Authorization Document. | 
**ContactTitle** | **string** | The title of the person authorized to sign the Authorization Document for this phone number. | 
**Email** | **string** | Email that this AuthorizationDocument will be sent to for signing. | 
**HostedNumberOrderSids** | **[]string** | A list of HostedNumberOrder sids that this AuthorizationDocument will authorize for hosting phone number capabilities on Twilio&#39;s platform. | 

### Return type

[**PreviewHostedNumbersAuthorizationDocument**](PreviewHostedNumbersAuthorizationDocument.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded, 
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCertificate

> PreviewDeployedDevicesFleetCertificate CreateCertificate(ctx, FleetSidoptional)



Enroll a new Certificate credential to the Fleet, optionally giving it a friendly name and assigning to a Device.

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**FleetSid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a CreateCertificateParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**CertificateData** | **string** | Provides a URL encoded representation of the public certificate in PEM format. | 
**DeviceSid** | **string** | Provides the unique string identifier of an existing Device to become authenticated with this Certificate credential. | 
**FriendlyName** | **string** | Provides a human readable descriptive text for this Certificate credential, up to 256 characters long. | 

### Return type

[**PreviewDeployedDevicesFleetCertificate**](PreviewDeployedDevicesFleetCertificate.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded, 
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateChannel

> PreviewTrustedCommsBrandedChannelChannel CreateChannel(ctx, BrandedChannelSidoptional)



Associate a channel to a branded channel

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**BrandedChannelSid** | **string** | The unique SID identifier of the Branded Channel. The given phone number is going to be assigned to this Branded Channel | 

### Other Parameters

Other parameters are passed through a pointer to a CreateChannelParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**PhoneNumberSid** | **string** | The unique SID identifier of the Phone Number of the Phone number to be assigned to the Branded Channel. | 

### Return type

[**PreviewTrustedCommsBrandedChannelChannel**](PreviewTrustedCommsBrandedChannelChannel.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded, 
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCommand

> PreviewWirelessCommand CreateCommand(ctx, optional)



### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateCommandParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**CallbackMethod** | **string** |  | 
**CallbackUrl** | **string** |  | 
**Command** | **string** |  | 
**CommandMode** | **string** |  | 
**Device** | **string** |  | 
**IncludeSid** | **string** |  | 
**Sim** | **string** |  | 

### Return type

[**PreviewWirelessCommand**](PreviewWirelessCommand.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded, 
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateDeployment

> PreviewDeployedDevicesFleetDeployment CreateDeployment(ctx, FleetSidoptional)



Create a new Deployment in the Fleet, optionally giving it a friendly name and linking to a specific Twilio Sync service instance.

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**FleetSid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a CreateDeploymentParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**FriendlyName** | **string** | Provides a human readable descriptive text for this Deployment, up to 256 characters long. | 
**SyncServiceSid** | **string** | Provides the unique string identifier of the Twilio Sync service instance that will be linked to and accessible by this Deployment. | 

### Return type

[**PreviewDeployedDevicesFleetDeployment**](PreviewDeployedDevicesFleetDeployment.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded, 
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateDevice

> PreviewDeployedDevicesFleetDevice CreateDevice(ctx, FleetSidoptional)



Create a new Device in the Fleet, optionally giving it a unique name, friendly name, and assigning to a Deployment and/or human identity.

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**FleetSid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a CreateDeviceParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**DeploymentSid** | **string** | Specifies the unique string identifier of the Deployment group that this Device is going to be associated with. | 
**Enabled** | **bool** |  | 
**FriendlyName** | **string** | Provides a human readable descriptive text to be assigned to this Device, up to 256 characters long. | 
**Identity** | **string** | Provides an arbitrary string identifier representing a human user to be associated with this Device, up to 256 characters long. | 
**UniqueName** | **string** | Provides a unique and addressable name to be assigned to this Device, to be used in addition to SID, up to 128 characters long. | 

### Return type

[**PreviewDeployedDevicesFleetDevice**](PreviewDeployedDevicesFleetDevice.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded, 
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateDocument

> PreviewSyncServiceDocument CreateDocument(ctx, ServiceSidoptional)



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a CreateDocumentParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**Data** | [**map[string]interface{}**](map[string]interface{}.md) |  | 
**UniqueName** | **string** |  | 

### Return type

[**PreviewSyncServiceDocument**](PreviewSyncServiceDocument.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded, 
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateExportCustomJob

> PreviewBulkExportsExportExportCustomJob CreateExportCustomJob(ctx, ResourceTypeoptional)



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ResourceType** | **string** | The type of communication – Messages or Calls, Conferences, and Participants | 

### Other Parameters

Other parameters are passed through a pointer to a CreateExportCustomJobParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**Email** | **string** | The optional email to send the completion notification to | 
**EndDay** | **string** | The end day for the custom export specified as a string in the format of yyyy-mm-dd. End day is inclusive and must be 2 days earlier than the current UTC day. | 
**FriendlyName** | **string** | The friendly name specified when creating the job | 
**StartDay** | **string** | The start day for the custom export specified as a string in the format of yyyy-mm-dd | 
**WebhookMethod** | **string** | This is the method used to call the webhook on completion of the job. If this is supplied, &#x60;WebhookUrl&#x60; must also be supplied. | 
**WebhookUrl** | **string** | The optional webhook url called on completion of the job. If this is supplied, &#x60;WebhookMethod&#x60; must also be supplied. | 

### Return type

[**PreviewBulkExportsExportExportCustomJob**](PreviewBulkExportsExportExportCustomJob.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded, 
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateField

> PreviewUnderstandAssistantTaskField CreateField(ctx, AssistantSidTaskSidoptional)



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string** | The unique ID of the parent Assistant. | 
**TaskSid** | **string** | The unique ID of the Task associated with this Field. | 

### Other Parameters

Other parameters are passed through a pointer to a CreateFieldParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**FieldType** | **string** | The unique name or sid of the FieldType. It can be any [Built-in Field Type](https://www.twilio.com/docs/assistant/api/built-in-field-types) or the unique_name or the Field Type sid of a custom Field Type. | 
**UniqueName** | **string** | A user-provided string that uniquely identifies this resource as an alternative to the sid. Unique up to 64 characters long. | 

### Return type

[**PreviewUnderstandAssistantTaskField**](PreviewUnderstandAssistantTaskField.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded, 
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateFieldType

> PreviewUnderstandAssistantFieldType CreateFieldType(ctx, AssistantSidoptional)



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a CreateFieldTypeParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**FriendlyName** | **string** | A user-provided string that identifies this resource. It is non-unique and can up to 255 characters long. | 
**UniqueName** | **string** | A user-provided string that uniquely identifies this resource as an alternative to the sid. Unique up to 64 characters long. | 

### Return type

[**PreviewUnderstandAssistantFieldType**](PreviewUnderstandAssistantFieldType.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded, 
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateFieldValue

> PreviewUnderstandAssistantFieldTypeFieldValue CreateFieldValue(ctx, AssistantSidFieldTypeSidoptional)



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string** |  | 
**FieldTypeSid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a CreateFieldValueParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**Language** | **string** | An ISO language-country string of the value. | 
**SynonymOf** | **string** | A value that indicates this field value is a synonym of. Empty if the value is not a synonym. | 
**Value** | **string** | A user-provided string that uniquely identifies this resource as an alternative to the sid. Unique up to 64 characters long. | 

### Return type

[**PreviewUnderstandAssistantFieldTypeFieldValue**](PreviewUnderstandAssistantFieldTypeFieldValue.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded, 
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateFleet

> PreviewDeployedDevicesFleet CreateFleet(ctx, optional)



Create a new Fleet for scoping of deployed devices within your account.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateFleetParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**FriendlyName** | **string** | Provides a human readable descriptive text for this Fleet, up to 256 characters long. | 

### Return type

[**PreviewDeployedDevicesFleet**](PreviewDeployedDevicesFleet.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded, 
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateHostedNumberOrder

> PreviewHostedNumbersHostedNumberOrder CreateHostedNumberOrder(ctx, optional)



Host a phone number's capability on Twilio's platform.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateHostedNumberOrderParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**AccountSid** | **string** | This defaults to the AccountSid of the authorization the user is using. This can be provided to specify a subaccount to add the HostedNumberOrder to. | 
**AddressSid** | **string** | Optional. A 34 character string that uniquely identifies the Address resource that represents the address of the owner of this phone number. | 
**CcEmails** | **[]string** | Optional. A list of emails that the LOA document for this HostedNumberOrder will be carbon copied to. | 
**Email** | **string** | Optional. Email of the owner of this phone number that is being hosted. | 
**FriendlyName** | **string** | A 64 character string that is a human readable text that describes this resource. | 
**PhoneNumber** | **string** | The number to host in [+E.164](https://en.wikipedia.org/wiki/E.164) format | 
**SmsApplicationSid** | **string** | Optional. The 34 character sid of the application Twilio should use to handle SMS messages sent to this number. If a &#x60;SmsApplicationSid&#x60; is present, Twilio will ignore all of the SMS urls above and use those set on the application. | 
**SmsCapability** | **bool** | Used to specify that the SMS capability will be hosted on Twilio&#39;s platform. | 
**SmsFallbackMethod** | **string** | The HTTP method that should be used to request the SmsFallbackUrl. Must be either &#x60;GET&#x60; or &#x60;POST&#x60;. This will be copied onto the IncomingPhoneNumber resource. | 
**SmsFallbackUrl** | **string** | A URL that Twilio will request if an error occurs requesting or executing the TwiML defined by SmsUrl. This will be copied onto the IncomingPhoneNumber resource. | 
**SmsMethod** | **string** | The HTTP method that should be used to request the SmsUrl. Must be either &#x60;GET&#x60; or &#x60;POST&#x60;.  This will be copied onto the IncomingPhoneNumber resource. | 
**SmsUrl** | **string** | The URL that Twilio should request when somebody sends an SMS to the phone number. This will be copied onto the IncomingPhoneNumber resource. | 
**StatusCallbackMethod** | **string** | Optional. The Status Callback Method attached to the IncomingPhoneNumber resource. | 
**StatusCallbackUrl** | **string** | Optional. The Status Callback URL attached to the IncomingPhoneNumber resource. | 
**UniqueName** | **string** | Optional. Provides a unique and addressable name to be assigned to this HostedNumberOrder, assigned by the developer, to be optionally used in addition to SID. | 
**VerificationDocumentSid** | **string** | Optional. The unique sid identifier of the Identity Document that represents the document for verifying ownership of the number to be hosted. Required when VerificationType is phone-bill. | 
**VerificationType** | **string** | Optional. The method used for verifying ownership of the number to be hosted. One of phone-call (default) or phone-bill. | 

### Return type

[**PreviewHostedNumbersHostedNumberOrder**](PreviewHostedNumbersHostedNumberOrder.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded, 
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateInstalledAddOn

> PreviewMarketplaceInstalledAddOn CreateInstalledAddOn(ctx, optional)



Install an Add-on for the Account specified.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateInstalledAddOnParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**AcceptTermsOfService** | **bool** | Whether the Terms of Service were accepted. | 
**AvailableAddOnSid** | **string** | The SID of the AvaliableAddOn to install. | 
**Configuration** | [**map[string]interface{}**](map[string]interface{}.md) | The JSON object that represents the configuration of the new Add-on being installed. | 
**UniqueName** | **string** | An application-defined string that uniquely identifies the resource. This value must be unique within the Account. | 

### Return type

[**PreviewMarketplaceInstalledAddOn**](PreviewMarketplaceInstalledAddOn.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded, 
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateKey

> PreviewDeployedDevicesFleetKey CreateKey(ctx, FleetSidoptional)



Create a new Key credential in the Fleet, optionally giving it a friendly name and assigning to a Device.

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**FleetSid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a CreateKeyParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**DeviceSid** | **string** | Provides the unique string identifier of an existing Device to become authenticated with this Key credential. | 
**FriendlyName** | **string** | Provides a human readable descriptive text for this Key credential, up to 256 characters long. | 

### Return type

[**PreviewDeployedDevicesFleetKey**](PreviewDeployedDevicesFleetKey.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded, 
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateModelBuild

> PreviewUnderstandAssistantModelBuild CreateModelBuild(ctx, AssistantSidoptional)



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a CreateModelBuildParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**StatusCallback** | **string** |  | 
**UniqueName** | **string** | A user-provided string that uniquely identifies this resource as an alternative to the sid. Unique up to 64 characters long. For example: v0.1 | 

### Return type

[**PreviewUnderstandAssistantModelBuild**](PreviewUnderstandAssistantModelBuild.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded, 
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateQuery

> PreviewUnderstandAssistantQuery CreateQuery(ctx, AssistantSidoptional)



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string** | The unique ID of the parent Assistant. | 

### Other Parameters

Other parameters are passed through a pointer to a CreateQueryParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**Field** | **string** | Constraints the query to a given Field with an task. Useful when you know the Field you are expecting. It accepts one field in the format *task-unique-name-1*:*field-unique-name* | 
**Language** | **string** | An ISO language-country string of the sample. | 
**ModelBuild** | **string** | The Model Build Sid or unique name of the Model Build to be queried. | 
**Query** | **string** | A user-provided string that uniquely identifies this resource as an alternative to the sid. It can be up to 2048 characters long. | 
**Tasks** | **string** | Constraints the query to a set of tasks. Useful when you need to constrain the paths the user can take. Tasks should be comma separated *task-unique-name-1*, *task-unique-name-2* | 

### Return type

[**PreviewUnderstandAssistantQuery**](PreviewUnderstandAssistantQuery.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded, 
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateRatePlan

> PreviewWirelessRatePlan CreateRatePlan(ctx, optional)



### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateRatePlanParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**CommandsEnabled** | **bool** |  | 
**DataEnabled** | **bool** |  | 
**DataLimit** | **int32** |  | 
**DataMetering** | **string** |  | 
**FriendlyName** | **string** |  | 
**InternationalRoaming** | **[]string** |  | 
**MessagingEnabled** | **bool** |  | 
**NationalRoamingEnabled** | **bool** |  | 
**UniqueName** | **string** |  | 
**VoiceEnabled** | **bool** |  | 

### Return type

[**PreviewWirelessRatePlan**](PreviewWirelessRatePlan.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded, 
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSample

> PreviewUnderstandAssistantTaskSample CreateSample(ctx, AssistantSidTaskSidoptional)



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string** | The unique ID of the Assistant. | 
**TaskSid** | **string** | The unique ID of the Task associated with this Sample. | 

### Other Parameters

Other parameters are passed through a pointer to a CreateSampleParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**Language** | **string** | An ISO language-country string of the sample. | 
**SourceChannel** | **string** | The communication channel the sample was captured. It can be: *voice*, *sms*, *chat*, *alexa*, *google-assistant*, or *slack*. If not included the value will be null | 
**TaggedText** | **string** | The text example of how end-users may express this task. The sample may contain Field tag blocks. | 

### Return type

[**PreviewUnderstandAssistantTaskSample**](PreviewUnderstandAssistantTaskSample.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded, 
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateService

> PreviewSyncService CreateService(ctx, optional)



### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateServiceParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**AclEnabled** | **bool** |  | 
**FriendlyName** | **string** |  | 
**ReachabilityWebhooksEnabled** | **bool** |  | 
**WebhookUrl** | **string** |  | 

### Return type

[**PreviewSyncService**](PreviewSyncService.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded, 
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSyncList

> PreviewSyncServiceSyncList CreateSyncList(ctx, ServiceSidoptional)



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a CreateSyncListParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**UniqueName** | **string** |  | 

### Return type

[**PreviewSyncServiceSyncList**](PreviewSyncServiceSyncList.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded, 
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSyncListItem

> PreviewSyncServiceSyncListSyncListItem CreateSyncListItem(ctx, ServiceSidListSidoptional)



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** |  | 
**ListSid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a CreateSyncListItemParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**Data** | [**map[string]interface{}**](map[string]interface{}.md) |  | 

### Return type

[**PreviewSyncServiceSyncListSyncListItem**](PreviewSyncServiceSyncListSyncListItem.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded, 
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSyncMap

> PreviewSyncServiceSyncMap CreateSyncMap(ctx, ServiceSidoptional)



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a CreateSyncMapParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**UniqueName** | **string** |  | 

### Return type

[**PreviewSyncServiceSyncMap**](PreviewSyncServiceSyncMap.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded, 
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSyncMapItem

> PreviewSyncServiceSyncMapSyncMapItem CreateSyncMapItem(ctx, ServiceSidMapSidoptional)



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** |  | 
**MapSid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a CreateSyncMapItemParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**Data** | [**map[string]interface{}**](map[string]interface{}.md) |  | 
**Key** | **string** |  | 

### Return type

[**PreviewSyncServiceSyncMapSyncMapItem**](PreviewSyncServiceSyncMapSyncMapItem.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded, 
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTask

> PreviewUnderstandAssistantTask CreateTask(ctx, AssistantSidoptional)



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string** | The unique ID of the Assistant. | 

### Other Parameters

Other parameters are passed through a pointer to a CreateTaskParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**Actions** | [**map[string]interface{}**](map[string]interface{}.md) | A user-provided JSON object encoded as a string to specify the actions for this task. It is optional and non-unique. | 
**ActionsUrl** | **string** | User-provided HTTP endpoint where from the assistant fetches actions | 
**FriendlyName** | **string** | A user-provided string that identifies this resource. It is non-unique and can up to 255 characters long. | 
**UniqueName** | **string** | A user-provided string that uniquely identifies this resource as an alternative to the sid. Unique up to 64 characters long. | 

### Return type

[**PreviewUnderstandAssistantTask**](PreviewUnderstandAssistantTask.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded, 
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAssistant

> DeleteAssistant(ctx, Sid)



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | A 34 character string that uniquely identifies this resource. | 

### Other Parameters

Other parameters are passed through a pointer to a DeleteAssistantParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

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

> DeleteCertificate(ctx, FleetSidSid)



Unregister a specific Certificate credential from the Fleet, effectively disallowing any inbound client connections that are presenting it.

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**FleetSid** | **string** |  | 
**Sid** | **string** | Provides a 34 character string that uniquely identifies the requested Certificate credential resource. | 

### Other Parameters

Other parameters are passed through a pointer to a DeleteCertificateParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

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

> DeleteDeployment(ctx, FleetSidSid)



Delete a specific Deployment from the Fleet, leaving associated devices effectively undeployed.

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**FleetSid** | **string** |  | 
**Sid** | **string** | Provides a 34 character string that uniquely identifies the requested Deployment resource. | 

### Other Parameters

Other parameters are passed through a pointer to a DeleteDeploymentParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

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

> DeleteDevice(ctx, FleetSidSid)



Delete a specific Device from the Fleet, also removing it from associated Deployments.

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**FleetSid** | **string** |  | 
**Sid** | **string** | Provides a 34 character string that uniquely identifies the requested Device resource. | 

### Other Parameters

Other parameters are passed through a pointer to a DeleteDeviceParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

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

> DeleteDocument(ctx, ServiceSidSid)



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** |  | 
**Sid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a DeleteDocumentParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

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

> DeleteDocumentPermission(ctx, ServiceSidDocumentSidIdentity)



Delete a specific Sync Document Permission.

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** |  | 
**DocumentSid** | **string** | Identifier of the Sync Document. Either a SID or a unique name. | 
**Identity** | **string** | Arbitrary string identifier representing a user associated with an FPA token, assigned by the developer. | 

### Other Parameters

Other parameters are passed through a pointer to a DeleteDocumentPermissionParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

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

> DeleteField(ctx, AssistantSidTaskSidSid)



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string** | The unique ID of the Assistant. | 
**TaskSid** | **string** | The unique ID of the Task associated with this Field. | 
**Sid** | **string** | A 34 character string that uniquely identifies this resource. | 

### Other Parameters

Other parameters are passed through a pointer to a DeleteFieldParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

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

> DeleteFieldType(ctx, AssistantSidSid)



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string** |  | 
**Sid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a DeleteFieldTypeParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

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

> DeleteFieldValue(ctx, AssistantSidFieldTypeSidSid)



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string** |  | 
**FieldTypeSid** | **string** |  | 
**Sid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a DeleteFieldValueParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

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

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | Provides a 34 character string that uniquely identifies the requested Fleet resource. | 

### Other Parameters

Other parameters are passed through a pointer to a DeleteFleetParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

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

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | A 34 character string that uniquely identifies this HostedNumberOrder. | 

### Other Parameters

Other parameters are passed through a pointer to a DeleteHostedNumberOrderParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

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

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The SID of the InstalledAddOn resource to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a DeleteInstalledAddOnParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

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



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**JobSid** | **string** | The unique string that that we created to identify the Bulk Export job | 

### Other Parameters

Other parameters are passed through a pointer to a DeleteJobParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

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

> DeleteKey(ctx, FleetSidSid)



Delete a specific Key credential from the Fleet, effectively disallowing any inbound client connections that are presenting it.

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**FleetSid** | **string** |  | 
**Sid** | **string** | Provides a 34 character string that uniquely identifies the requested Key credential resource. | 

### Other Parameters

Other parameters are passed through a pointer to a DeleteKeyParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

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

> DeleteModelBuild(ctx, AssistantSidSid)



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string** |  | 
**Sid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a DeleteModelBuildParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

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

> DeleteQuery(ctx, AssistantSidSid)



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string** | The unique ID of the Assistant. | 
**Sid** | **string** | A 34 character string that uniquely identifies this resource. | 

### Other Parameters

Other parameters are passed through a pointer to a DeleteQueryParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

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



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a DeleteRatePlanParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

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

> DeleteSample(ctx, AssistantSidTaskSidSid)



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string** | The unique ID of the Assistant. | 
**TaskSid** | **string** | The unique ID of the Task associated with this Sample. | 
**Sid** | **string** | A 34 character string that uniquely identifies this resource. | 

### Other Parameters

Other parameters are passed through a pointer to a DeleteSampleParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

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



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a DeleteServiceParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

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

> DeleteSyncList(ctx, ServiceSidSid)



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** |  | 
**Sid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a DeleteSyncListParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

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

> DeleteSyncListItem(ctx, ServiceSidListSidIndexoptional)



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** |  | 
**ListSid** | **string** |  | 
**Index** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a DeleteSyncListItemParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**IfMatch** | **string** | The If-Match HTTP request header | 

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

> DeleteSyncListPermission(ctx, ServiceSidListSidIdentity)



Delete a specific Sync List Permission.

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** |  | 
**ListSid** | **string** | Identifier of the Sync List. Either a SID or a unique name. | 
**Identity** | **string** | Arbitrary string identifier representing a user associated with an FPA token, assigned by the developer. | 

### Other Parameters

Other parameters are passed through a pointer to a DeleteSyncListPermissionParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

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

> DeleteSyncMap(ctx, ServiceSidSid)



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** |  | 
**Sid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a DeleteSyncMapParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

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

> DeleteSyncMapItem(ctx, ServiceSidMapSidKeyoptional)



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** |  | 
**MapSid** | **string** |  | 
**Key** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a DeleteSyncMapItemParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**IfMatch** | **string** | The If-Match HTTP request header | 

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

> DeleteSyncMapPermission(ctx, ServiceSidMapSidIdentity)



Delete a specific Sync Map Permission.

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** |  | 
**MapSid** | **string** | Identifier of the Sync Map. Either a SID or a unique name. | 
**Identity** | **string** | Arbitrary string identifier representing a user associated with an FPA token, assigned by the developer. | 

### Other Parameters

Other parameters are passed through a pointer to a DeleteSyncMapPermissionParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

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

> DeleteTask(ctx, AssistantSidSid)



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string** | The unique ID of the Assistant. | 
**Sid** | **string** | A 34 character string that uniquely identifies this resource. | 

### Other Parameters

Other parameters are passed through a pointer to a DeleteTaskParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

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



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | A 34 character string that uniquely identifies this resource. | 

### Other Parameters

Other parameters are passed through a pointer to a FetchAssistantParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**PreviewUnderstandAssistant**](PreviewUnderstandAssistant.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchAssistantFallbackActions

> PreviewUnderstandAssistantAssistantFallbackActions FetchAssistantFallbackActions(ctx, AssistantSid)



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a FetchAssistantFallbackActionsParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**PreviewUnderstandAssistantAssistantFallbackActions**](PreviewUnderstandAssistantAssistantFallbackActions.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchAssistantInitiationActions

> PreviewUnderstandAssistantAssistantInitiationActions FetchAssistantInitiationActions(ctx, AssistantSid)



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a FetchAssistantInitiationActionsParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**PreviewUnderstandAssistantAssistantInitiationActions**](PreviewUnderstandAssistantAssistantInitiationActions.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchAuthorizationDocument

> PreviewHostedNumbersAuthorizationDocument FetchAuthorizationDocument(ctx, Sid)



Fetch a specific AuthorizationDocument.

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | A 34 character string that uniquely identifies this AuthorizationDocument. | 

### Other Parameters

Other parameters are passed through a pointer to a FetchAuthorizationDocumentParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**PreviewHostedNumbersAuthorizationDocument**](PreviewHostedNumbersAuthorizationDocument.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchAvailableAddOn

> PreviewMarketplaceAvailableAddOn FetchAvailableAddOn(ctx, Sid)



Fetch an instance of an Add-on currently available to be installed.

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The SID of the AvailableAddOn resource to fetch. | 

### Other Parameters

Other parameters are passed through a pointer to a FetchAvailableAddOnParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**PreviewMarketplaceAvailableAddOn**](PreviewMarketplaceAvailableAddOn.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchAvailableAddOnExtension

> PreviewMarketplaceAvailableAddOnAvailableAddOnExtension FetchAvailableAddOnExtension(ctx, AvailableAddOnSidSid)



Fetch an instance of an Extension for the Available Add-on.

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AvailableAddOnSid** | **string** | The SID of the AvailableAddOn resource with the extension to fetch. | 
**Sid** | **string** | The SID of the AvailableAddOn Extension resource to fetch. | 

### Other Parameters

Other parameters are passed through a pointer to a FetchAvailableAddOnExtensionParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**PreviewMarketplaceAvailableAddOnAvailableAddOnExtension**](PreviewMarketplaceAvailableAddOnAvailableAddOnExtension.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchBrandedChannel

> PreviewTrustedCommsBrandedChannel FetchBrandedChannel(ctx, Sid)



Fetch a specific Branded Channel.

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The unique SID identifier of the Branded Channel. | 

### Other Parameters

Other parameters are passed through a pointer to a FetchBrandedChannelParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**PreviewTrustedCommsBrandedChannel**](PreviewTrustedCommsBrandedChannel.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchBrandsInformation

> PreviewTrustedCommsBrandsInformation FetchBrandsInformation(ctx, optional)



Retrieve the newest available BrandInformation

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a FetchBrandsInformationParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**IfNoneMatch** | **string** | Standard &#x60;If-None-Match&#x60; HTTP header. For more information visit: https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/If-None-Match. | 

### Return type

[**PreviewTrustedCommsBrandsInformation**](PreviewTrustedCommsBrandsInformation.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchCertificate

> PreviewDeployedDevicesFleetCertificate FetchCertificate(ctx, FleetSidSid)



Fetch information about a specific Certificate credential in the Fleet.

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**FleetSid** | **string** |  | 
**Sid** | **string** | Provides a 34 character string that uniquely identifies the requested Certificate credential resource. | 

### Other Parameters

Other parameters are passed through a pointer to a FetchCertificateParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**PreviewDeployedDevicesFleetCertificate**](PreviewDeployedDevicesFleetCertificate.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchCommand

> PreviewWirelessCommand FetchCommand(ctx, Sid)



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a FetchCommandParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**PreviewWirelessCommand**](PreviewWirelessCommand.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchCps

> PreviewTrustedCommsCps FetchCps(ctx, optional)



Fetch a specific Call Placement Service (CPS) given a phone number via `X-XCNAM-Sensitive-Phone-Number` header.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a FetchCpsParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**XXcnamSensitivePhoneNumber** | **string** | Phone number used to retrieve its corresponding CPS. | 

### Return type

[**PreviewTrustedCommsCps**](PreviewTrustedCommsCps.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchCurrentCall

> PreviewTrustedCommsCurrentCall FetchCurrentCall(ctx, optional)



Retrieve a current call given the originating and terminating number via `X-XCNAM-Sensitive-Phone-Number-From` and `X-XCNAM-Sensitive-Phone-Number-To` headers.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a FetchCurrentCallParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**XXcnamSensitivePhoneNumberFrom** | **string** | The originating Phone Number, given in [E.164 format](https://www.twilio.com/docs/glossary/what-e164). This phone number should be a Twilio number, otherwise it will return an error with HTTP Status Code 400. | 
**XXcnamSensitivePhoneNumberTo** | **string** | The terminating Phone Number, given in [E.164 format](https://www.twilio.com/docs/glossary/what-e164). | 

### Return type

[**PreviewTrustedCommsCurrentCall**](PreviewTrustedCommsCurrentCall.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchDay

> FetchDay(ctx, ResourceTypeDay)



Fetch a specific Day.

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ResourceType** | **string** | The type of communication – Messages, Calls, Conferences, and Participants | 
**Day** | **string** | The ISO 8601 format date of the resources in the file, for a UTC day | 

### Other Parameters

Other parameters are passed through a pointer to a FetchDayParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

 (empty response body)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchDeployment

> PreviewDeployedDevicesFleetDeployment FetchDeployment(ctx, FleetSidSid)



Fetch information about a specific Deployment in the Fleet.

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**FleetSid** | **string** |  | 
**Sid** | **string** | Provides a 34 character string that uniquely identifies the requested Deployment resource. | 

### Other Parameters

Other parameters are passed through a pointer to a FetchDeploymentParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**PreviewDeployedDevicesFleetDeployment**](PreviewDeployedDevicesFleetDeployment.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchDevice

> PreviewDeployedDevicesFleetDevice FetchDevice(ctx, FleetSidSid)



Fetch information about a specific Device in the Fleet.

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**FleetSid** | **string** |  | 
**Sid** | **string** | Provides a 34 character string that uniquely identifies the requested Device resource. | 

### Other Parameters

Other parameters are passed through a pointer to a FetchDeviceParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**PreviewDeployedDevicesFleetDevice**](PreviewDeployedDevicesFleetDevice.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchDialogue

> PreviewUnderstandAssistantDialogue FetchDialogue(ctx, AssistantSidSid)



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string** |  | 
**Sid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a FetchDialogueParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**PreviewUnderstandAssistantDialogue**](PreviewUnderstandAssistantDialogue.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchDocument

> PreviewSyncServiceDocument FetchDocument(ctx, ServiceSidSid)



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** |  | 
**Sid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a FetchDocumentParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**PreviewSyncServiceDocument**](PreviewSyncServiceDocument.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchDocumentPermission

> PreviewSyncServiceDocumentDocumentPermission FetchDocumentPermission(ctx, ServiceSidDocumentSidIdentity)



Fetch a specific Sync Document Permission.

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** |  | 
**DocumentSid** | **string** | Identifier of the Sync Document. Either a SID or a unique name. | 
**Identity** | **string** | Arbitrary string identifier representing a user associated with an FPA token, assigned by the developer. | 

### Other Parameters

Other parameters are passed through a pointer to a FetchDocumentPermissionParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**PreviewSyncServiceDocumentDocumentPermission**](PreviewSyncServiceDocumentDocumentPermission.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchExport

> PreviewBulkExportsExport FetchExport(ctx, ResourceType)



Fetch a specific Export.

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ResourceType** | **string** | The type of communication – Messages, Calls, Conferences, and Participants | 

### Other Parameters

Other parameters are passed through a pointer to a FetchExportParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**PreviewBulkExportsExport**](PreviewBulkExportsExport.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchExportConfiguration

> PreviewBulkExportsExportConfiguration FetchExportConfiguration(ctx, ResourceType)



Fetch a specific Export Configuration.

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ResourceType** | **string** | The type of communication – Messages, Calls, Conferences, and Participants | 

### Other Parameters

Other parameters are passed through a pointer to a FetchExportConfigurationParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**PreviewBulkExportsExportConfiguration**](PreviewBulkExportsExportConfiguration.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchField

> PreviewUnderstandAssistantTaskField FetchField(ctx, AssistantSidTaskSidSid)



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string** | The unique ID of the Assistant. | 
**TaskSid** | **string** | The unique ID of the Task associated with this Field. | 
**Sid** | **string** | A 34 character string that uniquely identifies this resource. | 

### Other Parameters

Other parameters are passed through a pointer to a FetchFieldParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**PreviewUnderstandAssistantTaskField**](PreviewUnderstandAssistantTaskField.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchFieldType

> PreviewUnderstandAssistantFieldType FetchFieldType(ctx, AssistantSidSid)



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string** |  | 
**Sid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a FetchFieldTypeParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**PreviewUnderstandAssistantFieldType**](PreviewUnderstandAssistantFieldType.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchFieldValue

> PreviewUnderstandAssistantFieldTypeFieldValue FetchFieldValue(ctx, AssistantSidFieldTypeSidSid)



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string** |  | 
**FieldTypeSid** | **string** |  | 
**Sid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a FetchFieldValueParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**PreviewUnderstandAssistantFieldTypeFieldValue**](PreviewUnderstandAssistantFieldTypeFieldValue.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchFleet

> PreviewDeployedDevicesFleet FetchFleet(ctx, Sid)



Fetch information about a specific Fleet in your account.

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | Provides a 34 character string that uniquely identifies the requested Fleet resource. | 

### Other Parameters

Other parameters are passed through a pointer to a FetchFleetParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**PreviewDeployedDevicesFleet**](PreviewDeployedDevicesFleet.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchHostedNumberOrder

> PreviewHostedNumbersHostedNumberOrder FetchHostedNumberOrder(ctx, Sid)



Fetch a specific HostedNumberOrder.

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | A 34 character string that uniquely identifies this HostedNumberOrder. | 

### Other Parameters

Other parameters are passed through a pointer to a FetchHostedNumberOrderParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**PreviewHostedNumbersHostedNumberOrder**](PreviewHostedNumbersHostedNumberOrder.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchInstalledAddOn

> PreviewMarketplaceInstalledAddOn FetchInstalledAddOn(ctx, Sid)



Fetch an instance of an Add-on currently installed on this Account.

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The SID of the InstalledAddOn resource to fetch. | 

### Other Parameters

Other parameters are passed through a pointer to a FetchInstalledAddOnParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**PreviewMarketplaceInstalledAddOn**](PreviewMarketplaceInstalledAddOn.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchInstalledAddOnExtension

> PreviewMarketplaceInstalledAddOnInstalledAddOnExtension FetchInstalledAddOnExtension(ctx, InstalledAddOnSidSid)



Fetch an instance of an Extension for the Installed Add-on.

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**InstalledAddOnSid** | **string** | The SID of the InstalledAddOn resource with the extension to fetch. | 
**Sid** | **string** | The SID of the InstalledAddOn Extension resource to fetch. | 

### Other Parameters

Other parameters are passed through a pointer to a FetchInstalledAddOnExtensionParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**PreviewMarketplaceInstalledAddOnInstalledAddOnExtension**](PreviewMarketplaceInstalledAddOnInstalledAddOnExtension.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchJob

> PreviewBulkExportsExportJob FetchJob(ctx, JobSid)



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**JobSid** | **string** | The unique string that that we created to identify the Bulk Export job | 

### Other Parameters

Other parameters are passed through a pointer to a FetchJobParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**PreviewBulkExportsExportJob**](PreviewBulkExportsExportJob.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchKey

> PreviewDeployedDevicesFleetKey FetchKey(ctx, FleetSidSid)



Fetch information about a specific Key credential in the Fleet.

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**FleetSid** | **string** |  | 
**Sid** | **string** | Provides a 34 character string that uniquely identifies the requested Key credential resource. | 

### Other Parameters

Other parameters are passed through a pointer to a FetchKeyParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**PreviewDeployedDevicesFleetKey**](PreviewDeployedDevicesFleetKey.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchModelBuild

> PreviewUnderstandAssistantModelBuild FetchModelBuild(ctx, AssistantSidSid)



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string** |  | 
**Sid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a FetchModelBuildParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**PreviewUnderstandAssistantModelBuild**](PreviewUnderstandAssistantModelBuild.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchQuery

> PreviewUnderstandAssistantQuery FetchQuery(ctx, AssistantSidSid)



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string** | The unique ID of the Assistant. | 
**Sid** | **string** | A 34 character string that uniquely identifies this resource. | 

### Other Parameters

Other parameters are passed through a pointer to a FetchQueryParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**PreviewUnderstandAssistantQuery**](PreviewUnderstandAssistantQuery.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchRatePlan

> PreviewWirelessRatePlan FetchRatePlan(ctx, Sid)



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a FetchRatePlanParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**PreviewWirelessRatePlan**](PreviewWirelessRatePlan.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchSample

> PreviewUnderstandAssistantTaskSample FetchSample(ctx, AssistantSidTaskSidSid)



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string** | The unique ID of the Assistant. | 
**TaskSid** | **string** | The unique ID of the Task associated with this Sample. | 
**Sid** | **string** | A 34 character string that uniquely identifies this resource. | 

### Other Parameters

Other parameters are passed through a pointer to a FetchSampleParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**PreviewUnderstandAssistantTaskSample**](PreviewUnderstandAssistantTaskSample.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchService

> PreviewSyncService FetchService(ctx, Sid)



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a FetchServiceParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**PreviewSyncService**](PreviewSyncService.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchSim

> PreviewWirelessSim FetchSim(ctx, Sid)



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a FetchSimParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**PreviewWirelessSim**](PreviewWirelessSim.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchStyleSheet

> PreviewUnderstandAssistantStyleSheet FetchStyleSheet(ctx, AssistantSid)



Returns Style sheet JSON object for this Assistant

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string** | The unique ID of the Assistant | 

### Other Parameters

Other parameters are passed through a pointer to a FetchStyleSheetParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**PreviewUnderstandAssistantStyleSheet**](PreviewUnderstandAssistantStyleSheet.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchSyncList

> PreviewSyncServiceSyncList FetchSyncList(ctx, ServiceSidSid)



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** |  | 
**Sid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a FetchSyncListParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**PreviewSyncServiceSyncList**](PreviewSyncServiceSyncList.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchSyncListItem

> PreviewSyncServiceSyncListSyncListItem FetchSyncListItem(ctx, ServiceSidListSidIndex)



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** |  | 
**ListSid** | **string** |  | 
**Index** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a FetchSyncListItemParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**PreviewSyncServiceSyncListSyncListItem**](PreviewSyncServiceSyncListSyncListItem.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchSyncListPermission

> PreviewSyncServiceSyncListSyncListPermission FetchSyncListPermission(ctx, ServiceSidListSidIdentity)



Fetch a specific Sync List Permission.

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** |  | 
**ListSid** | **string** | Identifier of the Sync List. Either a SID or a unique name. | 
**Identity** | **string** | Arbitrary string identifier representing a user associated with an FPA token, assigned by the developer. | 

### Other Parameters

Other parameters are passed through a pointer to a FetchSyncListPermissionParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**PreviewSyncServiceSyncListSyncListPermission**](PreviewSyncServiceSyncListSyncListPermission.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchSyncMap

> PreviewSyncServiceSyncMap FetchSyncMap(ctx, ServiceSidSid)



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** |  | 
**Sid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a FetchSyncMapParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**PreviewSyncServiceSyncMap**](PreviewSyncServiceSyncMap.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchSyncMapItem

> PreviewSyncServiceSyncMapSyncMapItem FetchSyncMapItem(ctx, ServiceSidMapSidKey)



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** |  | 
**MapSid** | **string** |  | 
**Key** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a FetchSyncMapItemParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**PreviewSyncServiceSyncMapSyncMapItem**](PreviewSyncServiceSyncMapSyncMapItem.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchSyncMapPermission

> PreviewSyncServiceSyncMapSyncMapPermission FetchSyncMapPermission(ctx, ServiceSidMapSidIdentity)



Fetch a specific Sync Map Permission.

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** |  | 
**MapSid** | **string** | Identifier of the Sync Map. Either a SID or a unique name. | 
**Identity** | **string** | Arbitrary string identifier representing a user associated with an FPA token, assigned by the developer. | 

### Other Parameters

Other parameters are passed through a pointer to a FetchSyncMapPermissionParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**PreviewSyncServiceSyncMapSyncMapPermission**](PreviewSyncServiceSyncMapSyncMapPermission.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchTask

> PreviewUnderstandAssistantTask FetchTask(ctx, AssistantSidSid)



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string** | The unique ID of the Assistant. | 
**Sid** | **string** | A 34 character string that uniquely identifies this resource. | 

### Other Parameters

Other parameters are passed through a pointer to a FetchTaskParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**PreviewUnderstandAssistantTask**](PreviewUnderstandAssistantTask.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchTaskActions

> PreviewUnderstandAssistantTaskTaskActions FetchTaskActions(ctx, AssistantSidTaskSid)



Returns JSON actions for this Task.

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string** | The unique ID of the parent Assistant. | 
**TaskSid** | **string** | The unique ID of the Task. | 

### Other Parameters

Other parameters are passed through a pointer to a FetchTaskActionsParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**PreviewUnderstandAssistantTaskTaskActions**](PreviewUnderstandAssistantTaskTaskActions.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchTaskStatistics

> PreviewUnderstandAssistantTaskTaskStatistics FetchTaskStatistics(ctx, AssistantSidTaskSid)



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string** | The unique ID of the parent Assistant. | 
**TaskSid** | **string** | The unique ID of the Task associated with this Field. | 

### Other Parameters

Other parameters are passed through a pointer to a FetchTaskStatisticsParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**PreviewUnderstandAssistantTaskTaskStatistics**](PreviewUnderstandAssistantTaskTaskStatistics.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchUsage

> PreviewWirelessSimUsage FetchUsage(ctx, SimSidoptional)



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**SimSid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a FetchUsageParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**End** | **string** |  | 
**Start** | **string** |  | 

### Return type

[**PreviewWirelessSimUsage**](PreviewWirelessSimUsage.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAssistant

> ListAssistantResponse ListAssistant(ctx, optional)



### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListAssistantParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**ListAssistantResponse**](ListAssistantResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAuthorizationDocument

> ListAuthorizationDocumentResponse ListAuthorizationDocument(ctx, optional)



Retrieve a list of AuthorizationDocuments belonging to the account initiating the request.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListAuthorizationDocumentParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**Email** | **string** | Email that this AuthorizationDocument will be sent to for signing. | 
**Status** | **string** | Status of an instance resource. It can hold one of the values: 1. opened 2. signing, 3. signed LOA, 4. canceled, 5. failed. See the section entitled [Status Values](https://www.twilio.com/docs/api/phone-numbers/hosted-number-authorization-documents#status-values) for more information on each of these statuses. | 
**PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**ListAuthorizationDocumentResponse**](ListAuthorizationDocumentResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAvailableAddOn

> ListAvailableAddOnResponse ListAvailableAddOn(ctx, optional)



Retrieve a list of Add-ons currently available to be installed.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListAvailableAddOnParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**ListAvailableAddOnResponse**](ListAvailableAddOnResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAvailableAddOnExtension

> ListAvailableAddOnExtensionResponse ListAvailableAddOnExtension(ctx, AvailableAddOnSidoptional)



Retrieve a list of Extensions for the Available Add-on.

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AvailableAddOnSid** | **string** | The SID of the AvailableAddOn resource with the extensions to read. | 

### Other Parameters

Other parameters are passed through a pointer to a ListAvailableAddOnExtensionParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**ListAvailableAddOnExtensionResponse**](ListAvailableAddOnExtensionResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCertificate

> ListCertificateResponse ListCertificate(ctx, FleetSidoptional)



Retrieve a list of all Certificate credentials belonging to the Fleet.

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**FleetSid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a ListCertificateParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**DeviceSid** | **string** | Filters the resulting list of Certificates by a unique string identifier of an authenticated Device. | 
**PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**ListCertificateResponse**](ListCertificateResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCommand

> ListCommandResponse ListCommand(ctx, optional)



### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListCommandParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**Device** | **string** |  | 
**Sim** | **string** |  | 
**Status** | **string** |  | 
**Direction** | **string** |  | 
**PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**ListCommandResponse**](ListCommandResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDay

> ListDayResponse ListDay(ctx, ResourceTypeoptional)



Retrieve a list of all Days for a resource.

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ResourceType** | **string** | The type of communication – Messages, Calls, Conferences, and Participants | 

### Other Parameters

Other parameters are passed through a pointer to a ListDayParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**ListDayResponse**](ListDayResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDependentHostedNumberOrder

> ListDependentHostedNumberOrderResponse ListDependentHostedNumberOrder(ctx, SigningDocumentSidoptional)



Retrieve a list of dependent HostedNumberOrders belonging to the AuthorizationDocument.

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**SigningDocumentSid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a ListDependentHostedNumberOrderParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**Status** | **string** | Status of an instance resource. It can hold one of the values: 1. opened 2. signing, 3. signed LOA, 4. canceled, 5. failed. See the section entitled [Status Values](https://www.twilio.com/docs/api/phone-numbers/hosted-number-authorization-documents#status-values) for more information on each of these statuses. | 
**PhoneNumber** | **string** | An E164 formatted phone number hosted by this HostedNumberOrder. | 
**IncomingPhoneNumberSid** | **string** | A 34 character string that uniquely identifies the IncomingPhoneNumber resource created by this HostedNumberOrder. | 
**FriendlyName** | **string** | A human readable description of this resource, up to 64 characters. | 
**UniqueName** | **string** | Provides a unique and addressable name to be assigned to this HostedNumberOrder, assigned by the developer, to be optionally used in addition to SID. | 
**PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**ListDependentHostedNumberOrderResponse**](ListDependentHostedNumberOrderResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDeployment

> ListDeploymentResponse ListDeployment(ctx, FleetSidoptional)



Retrieve a list of all Deployments belonging to the Fleet.

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**FleetSid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a ListDeploymentParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**ListDeploymentResponse**](ListDeploymentResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDevice

> ListDeviceResponse ListDevice(ctx, FleetSidoptional)



Retrieve a list of all Devices belonging to the Fleet.

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**FleetSid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a ListDeviceParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**DeploymentSid** | **string** | Filters the resulting list of Devices by a unique string identifier of the Deployment they are associated with. | 
**PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**ListDeviceResponse**](ListDeviceResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDocument

> ListDocumentResponse ListDocument(ctx, ServiceSidoptional)



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a ListDocumentParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**ListDocumentResponse**](ListDocumentResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDocumentPermission

> ListDocumentPermissionResponse ListDocumentPermission(ctx, ServiceSidDocumentSidoptional)



Retrieve a list of all Permissions applying to a Sync Document.

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** |  | 
**DocumentSid** | **string** | Identifier of the Sync Document. Either a SID or a unique name. | 

### Other Parameters

Other parameters are passed through a pointer to a ListDocumentPermissionParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**ListDocumentPermissionResponse**](ListDocumentPermissionResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListExportCustomJob

> ListExportCustomJobResponse ListExportCustomJob(ctx, ResourceTypeoptional)



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ResourceType** | **string** | The type of communication – Messages, Calls, Conferences, and Participants | 

### Other Parameters

Other parameters are passed through a pointer to a ListExportCustomJobParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**ListExportCustomJobResponse**](ListExportCustomJobResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListField

> ListFieldResponse ListField(ctx, AssistantSidTaskSidoptional)



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string** | The unique ID of the Assistant. | 
**TaskSid** | **string** | The unique ID of the Task associated with this Field. | 

### Other Parameters

Other parameters are passed through a pointer to a ListFieldParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**ListFieldResponse**](ListFieldResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListFieldType

> ListFieldTypeResponse ListFieldType(ctx, AssistantSidoptional)



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a ListFieldTypeParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**ListFieldTypeResponse**](ListFieldTypeResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListFieldValue

> ListFieldValueResponse ListFieldValue(ctx, AssistantSidFieldTypeSidoptional)



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string** |  | 
**FieldTypeSid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a ListFieldValueParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**Language** | **string** | An ISO language-country string of the value. For example: *en-US* | 
**PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**ListFieldValueResponse**](ListFieldValueResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListFleet

> ListFleetResponse ListFleet(ctx, optional)



Retrieve a list of all Fleets belonging to your account.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListFleetParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**ListFleetResponse**](ListFleetResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListHostedNumberOrder

> ListHostedNumberOrderResponse ListHostedNumberOrder(ctx, optional)



Retrieve a list of HostedNumberOrders belonging to the account initiating the request.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListHostedNumberOrderParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**Status** | **string** | The Status of this HostedNumberOrder. One of &#x60;received&#x60;, &#x60;pending-verification&#x60;, &#x60;verified&#x60;, &#x60;pending-loa&#x60;, &#x60;carrier-processing&#x60;, &#x60;testing&#x60;, &#x60;completed&#x60;, &#x60;failed&#x60;, or &#x60;action-required&#x60;. | 
**PhoneNumber** | **string** | An E164 formatted phone number hosted by this HostedNumberOrder. | 
**IncomingPhoneNumberSid** | **string** | A 34 character string that uniquely identifies the IncomingPhoneNumber resource created by this HostedNumberOrder. | 
**FriendlyName** | **string** | A human readable description of this resource, up to 64 characters. | 
**UniqueName** | **string** | Provides a unique and addressable name to be assigned to this HostedNumberOrder, assigned by the developer, to be optionally used in addition to SID. | 
**PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**ListHostedNumberOrderResponse**](ListHostedNumberOrderResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListInstalledAddOn

> ListInstalledAddOnResponse ListInstalledAddOn(ctx, optional)



Retrieve a list of Add-ons currently installed on this Account.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListInstalledAddOnParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**ListInstalledAddOnResponse**](ListInstalledAddOnResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListInstalledAddOnExtension

> ListInstalledAddOnExtensionResponse ListInstalledAddOnExtension(ctx, InstalledAddOnSidoptional)



Retrieve a list of Extensions for the Installed Add-on.

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**InstalledAddOnSid** | **string** | The SID of the InstalledAddOn resource with the extensions to read. | 

### Other Parameters

Other parameters are passed through a pointer to a ListInstalledAddOnExtensionParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**ListInstalledAddOnExtensionResponse**](ListInstalledAddOnExtensionResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListKey

> ListKeyResponse ListKey(ctx, FleetSidoptional)



Retrieve a list of all Keys credentials belonging to the Fleet.

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**FleetSid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a ListKeyParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**DeviceSid** | **string** | Filters the resulting list of Keys by a unique string identifier of an authenticated Device. | 
**PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**ListKeyResponse**](ListKeyResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListModelBuild

> ListModelBuildResponse ListModelBuild(ctx, AssistantSidoptional)



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a ListModelBuildParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**ListModelBuildResponse**](ListModelBuildResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListQuery

> ListQueryResponse ListQuery(ctx, AssistantSidoptional)



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string** | The unique ID of the parent Assistant. | 

### Other Parameters

Other parameters are passed through a pointer to a ListQueryParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**Language** | **string** | An ISO language-country string of the sample. | 
**ModelBuild** | **string** | The Model Build Sid or unique name of the Model Build to be queried. | 
**Status** | **string** | A string that described the query status. The values can be: pending_review, reviewed, discarded | 
**PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**ListQueryResponse**](ListQueryResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRatePlan

> ListRatePlanResponse ListRatePlan(ctx, optional)



### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListRatePlanParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**ListRatePlanResponse**](ListRatePlanResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSample

> ListSampleResponse ListSample(ctx, AssistantSidTaskSidoptional)



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string** | The unique ID of the Assistant. | 
**TaskSid** | **string** | The unique ID of the Task associated with this Sample. | 

### Other Parameters

Other parameters are passed through a pointer to a ListSampleParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**Language** | **string** | An ISO language-country string of the sample. | 
**PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**ListSampleResponse**](ListSampleResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListService

> ListServiceResponse ListService(ctx, optional)



### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListServiceParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**ListServiceResponse**](ListServiceResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSim

> ListSimResponse ListSim(ctx, optional)



### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListSimParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**Status** | **string** |  | 
**Iccid** | **string** |  | 
**RatePlan** | **string** |  | 
**EId** | **string** |  | 
**SimRegistrationCode** | **string** |  | 
**PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**ListSimResponse**](ListSimResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSyncList

> ListSyncListResponse ListSyncList(ctx, ServiceSidoptional)



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a ListSyncListParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**ListSyncListResponse**](ListSyncListResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSyncListItem

> ListSyncListItemResponse ListSyncListItem(ctx, ServiceSidListSidoptional)



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** |  | 
**ListSid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a ListSyncListItemParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**Order** | **string** |  | 
**From** | **string** |  | 
**Bounds** | **string** |  | 
**PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**ListSyncListItemResponse**](ListSyncListItemResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSyncListPermission

> ListSyncListPermissionResponse ListSyncListPermission(ctx, ServiceSidListSidoptional)



Retrieve a list of all Permissions applying to a Sync List.

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** |  | 
**ListSid** | **string** | Identifier of the Sync List. Either a SID or a unique name. | 

### Other Parameters

Other parameters are passed through a pointer to a ListSyncListPermissionParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**ListSyncListPermissionResponse**](ListSyncListPermissionResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSyncMap

> ListSyncMapResponse ListSyncMap(ctx, ServiceSidoptional)



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a ListSyncMapParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**ListSyncMapResponse**](ListSyncMapResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSyncMapItem

> ListSyncMapItemResponse ListSyncMapItem(ctx, ServiceSidMapSidoptional)



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** |  | 
**MapSid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a ListSyncMapItemParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**Order** | **string** |  | 
**From** | **string** |  | 
**Bounds** | **string** |  | 
**PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**ListSyncMapItemResponse**](ListSyncMapItemResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSyncMapPermission

> ListSyncMapPermissionResponse ListSyncMapPermission(ctx, ServiceSidMapSidoptional)



Retrieve a list of all Permissions applying to a Sync Map.

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** |  | 
**MapSid** | **string** | Identifier of the Sync Map. Either a SID or a unique name. | 

### Other Parameters

Other parameters are passed through a pointer to a ListSyncMapPermissionParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**ListSyncMapPermissionResponse**](ListSyncMapPermissionResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTask

> ListTaskResponse ListTask(ctx, AssistantSidoptional)



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string** | The unique ID of the Assistant. | 

### Other Parameters

Other parameters are passed through a pointer to a ListTaskParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**ListTaskResponse**](ListTaskResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAssistant

> PreviewUnderstandAssistant UpdateAssistant(ctx, Sidoptional)



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | A 34 character string that uniquely identifies this resource. | 

### Other Parameters

Other parameters are passed through a pointer to a UpdateAssistantParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**CallbackEvents** | **string** | Space-separated list of callback events that will trigger callbacks. | 
**CallbackUrl** | **string** | A user-provided URL to send event callbacks to. | 
**FallbackActions** | [**map[string]interface{}**](map[string]interface{}.md) | The JSON actions to be executed when the user&#39;s input is not recognized as matching any Task. | 
**FriendlyName** | **string** | A text description for the Assistant. It is non-unique and can up to 255 characters long. | 
**InitiationActions** | [**map[string]interface{}**](map[string]interface{}.md) | The JSON actions to be executed on inbound phone calls when the Assistant has to say something first. | 
**LogQueries** | **bool** | A boolean that specifies whether queries should be logged for 30 days further training. If false, no queries will be stored, if true, queries will be stored for 30 days and deleted thereafter. Defaults to true if no value is provided. | 
**StyleSheet** | [**map[string]interface{}**](map[string]interface{}.md) | The JSON object that holds the style sheet for the assistant | 
**UniqueName** | **string** | A user-provided string that uniquely identifies this resource as an alternative to the sid. Unique up to 64 characters long. | 

### Return type

[**PreviewUnderstandAssistant**](PreviewUnderstandAssistant.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded, 
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAssistantFallbackActions

> PreviewUnderstandAssistantAssistantFallbackActions UpdateAssistantFallbackActions(ctx, AssistantSidoptional)



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a UpdateAssistantFallbackActionsParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**FallbackActions** | [**map[string]interface{}**](map[string]interface{}.md) |  | 

### Return type

[**PreviewUnderstandAssistantAssistantFallbackActions**](PreviewUnderstandAssistantAssistantFallbackActions.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded, 
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAssistantInitiationActions

> PreviewUnderstandAssistantAssistantInitiationActions UpdateAssistantInitiationActions(ctx, AssistantSidoptional)



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a UpdateAssistantInitiationActionsParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**InitiationActions** | [**map[string]interface{}**](map[string]interface{}.md) |  | 

### Return type

[**PreviewUnderstandAssistantAssistantInitiationActions**](PreviewUnderstandAssistantAssistantInitiationActions.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded, 
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAuthorizationDocument

> PreviewHostedNumbersAuthorizationDocument UpdateAuthorizationDocument(ctx, Sidoptional)



Updates a specific AuthorizationDocument.

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a UpdateAuthorizationDocumentParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**AddressSid** | **string** | A 34 character string that uniquely identifies the Address resource that is associated with this AuthorizationDocument. | 
**CcEmails** | **[]string** | Email recipients who will be informed when an Authorization Document has been sent and signed | 
**ContactPhoneNumber** | **string** | The contact phone number of the person authorized to sign the Authorization Document. | 
**ContactTitle** | **string** | The title of the person authorized to sign the Authorization Document for this phone number. | 
**Email** | **string** | Email that this AuthorizationDocument will be sent to for signing. | 
**HostedNumberOrderSids** | **[]string** | A list of HostedNumberOrder sids that this AuthorizationDocument will authorize for hosting phone number capabilities on Twilio&#39;s platform. | 
**Status** | **string** | Status of an instance resource. It can hold one of the values: 1. opened 2. signing, 3. signed LOA, 4. canceled, 5. failed. See the section entitled [Status Values](https://www.twilio.com/docs/api/phone-numbers/hosted-number-authorization-documents#status-values) for more information on each of these statuses. | 

### Return type

[**PreviewHostedNumbersAuthorizationDocument**](PreviewHostedNumbersAuthorizationDocument.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded, 
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCertificate

> PreviewDeployedDevicesFleetCertificate UpdateCertificate(ctx, FleetSidSidoptional)



Update the given properties of a specific Certificate credential in the Fleet, giving it a friendly name or assigning to a Device.

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**FleetSid** | **string** |  | 
**Sid** | **string** | Provides a 34 character string that uniquely identifies the requested Certificate credential resource. | 

### Other Parameters

Other parameters are passed through a pointer to a UpdateCertificateParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**DeviceSid** | **string** | Provides the unique string identifier of an existing Device to become authenticated with this Certificate credential. | 
**FriendlyName** | **string** | Provides a human readable descriptive text for this Certificate credential, up to 256 characters long. | 

### Return type

[**PreviewDeployedDevicesFleetCertificate**](PreviewDeployedDevicesFleetCertificate.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded, 
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDeployment

> PreviewDeployedDevicesFleetDeployment UpdateDeployment(ctx, FleetSidSidoptional)



Update the given properties of a specific Deployment credential in the Fleet, giving it a friendly name or linking to a specific Twilio Sync service instance.

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**FleetSid** | **string** |  | 
**Sid** | **string** | Provides a 34 character string that uniquely identifies the requested Deployment resource. | 

### Other Parameters

Other parameters are passed through a pointer to a UpdateDeploymentParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**FriendlyName** | **string** | Provides a human readable descriptive text for this Deployment, up to 64 characters long | 
**SyncServiceSid** | **string** | Provides the unique string identifier of the Twilio Sync service instance that will be linked to and accessible by this Deployment. | 

### Return type

[**PreviewDeployedDevicesFleetDeployment**](PreviewDeployedDevicesFleetDeployment.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded, 
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDevice

> PreviewDeployedDevicesFleetDevice UpdateDevice(ctx, FleetSidSidoptional)



Update the given properties of a specific Device in the Fleet, giving it a friendly name, assigning to a Deployment, or a human identity.

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**FleetSid** | **string** |  | 
**Sid** | **string** | Provides a 34 character string that uniquely identifies the requested Device resource. | 

### Other Parameters

Other parameters are passed through a pointer to a UpdateDeviceParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**DeploymentSid** | **string** | Specifies the unique string identifier of the Deployment group that this Device is going to be associated with. | 
**Enabled** | **bool** |  | 
**FriendlyName** | **string** | Provides a human readable descriptive text to be assigned to this Device, up to 256 characters long. | 
**Identity** | **string** | Provides an arbitrary string identifier representing a human user to be associated with this Device, up to 256 characters long. | 

### Return type

[**PreviewDeployedDevicesFleetDevice**](PreviewDeployedDevicesFleetDevice.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded, 
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDocument

> PreviewSyncServiceDocument UpdateDocument(ctx, ServiceSidSidoptional)



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** |  | 
**Sid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a UpdateDocumentParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**IfMatch** | **string** | The If-Match HTTP request header | 
**Data** | [**map[string]interface{}**](map[string]interface{}.md) |  | 

### Return type

[**PreviewSyncServiceDocument**](PreviewSyncServiceDocument.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded, 
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDocumentPermission

> PreviewSyncServiceDocumentDocumentPermission UpdateDocumentPermission(ctx, ServiceSidDocumentSidIdentityoptional)



Update an identity's access to a specific Sync Document.

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The unique SID identifier of the Sync Service Instance. | 
**DocumentSid** | **string** | Identifier of the Sync Document. Either a SID or a unique name. | 
**Identity** | **string** | Arbitrary string identifier representing a human user associated with an FPA token, assigned by the developer. | 

### Other Parameters

Other parameters are passed through a pointer to a UpdateDocumentPermissionParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**Manage** | **bool** | Boolean flag specifying whether the identity can delete the Sync Document. | 
**Read** | **bool** | Boolean flag specifying whether the identity can read the Sync Document. | 
**Write** | **bool** | Boolean flag specifying whether the identity can update the Sync Document. | 

### Return type

[**PreviewSyncServiceDocumentDocumentPermission**](PreviewSyncServiceDocumentDocumentPermission.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded, 
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateExportConfiguration

> PreviewBulkExportsExportConfiguration UpdateExportConfiguration(ctx, ResourceTypeoptional)



Update a specific Export Configuration.

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ResourceType** | **string** | The type of communication – Messages, Calls, Conferences, and Participants | 

### Other Parameters

Other parameters are passed through a pointer to a UpdateExportConfigurationParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**Enabled** | **bool** | If true, Twilio will automatically generate every day&#39;s file when the day is over. | 
**WebhookMethod** | **string** | Sets whether Twilio should call a webhook URL when the automatic generation is complete, using GET or POST. The actual destination is set in the webhook_url | 
**WebhookUrl** | **string** | Stores the URL destination for the method specified in webhook_method. | 

### Return type

[**PreviewBulkExportsExportConfiguration**](PreviewBulkExportsExportConfiguration.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded, 
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateFieldType

> PreviewUnderstandAssistantFieldType UpdateFieldType(ctx, AssistantSidSidoptional)



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string** |  | 
**Sid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a UpdateFieldTypeParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**FriendlyName** | **string** | A user-provided string that identifies this resource. It is non-unique and can up to 255 characters long. | 
**UniqueName** | **string** | A user-provided string that uniquely identifies this resource as an alternative to the sid. Unique up to 64 characters long. | 

### Return type

[**PreviewUnderstandAssistantFieldType**](PreviewUnderstandAssistantFieldType.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded, 
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateFleet

> PreviewDeployedDevicesFleet UpdateFleet(ctx, Sidoptional)



Update the friendly name property of a specific Fleet in your account.

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | Provides a 34 character string that uniquely identifies the requested Fleet resource. | 

### Other Parameters

Other parameters are passed through a pointer to a UpdateFleetParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**DefaultDeploymentSid** | **string** | Provides a string identifier of a Deployment that is going to be used as a default one for this Fleet. | 
**FriendlyName** | **string** | Provides a human readable descriptive text for this Fleet, up to 256 characters long. | 

### Return type

[**PreviewDeployedDevicesFleet**](PreviewDeployedDevicesFleet.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded, 
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateHostedNumberOrder

> PreviewHostedNumbersHostedNumberOrder UpdateHostedNumberOrder(ctx, Sidoptional)



Updates a specific HostedNumberOrder.

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a UpdateHostedNumberOrderParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**CallDelay** | **int32** | The number of seconds, between 0 and 60, to delay before initiating the verification call. Defaults to 0. | 
**CcEmails** | **[]string** | Optional. A list of emails that LOA document for this HostedNumberOrder will be carbon copied to. | 
**Email** | **string** | Email of the owner of this phone number that is being hosted. | 
**Extension** | **string** | Digits to dial after connecting the verification call. | 
**FriendlyName** | **string** | A 64 character string that is a human readable text that describes this resource. | 
**Status** | **string** | User can only post to &#x60;pending-verification&#x60; status to transition the HostedNumberOrder to initiate a verification call or verification of ownership with a copy of a phone bill. | 
**UniqueName** | **string** | Provides a unique and addressable name to be assigned to this HostedNumberOrder, assigned by the developer, to be optionally used in addition to SID. | 
**VerificationCode** | **string** | A verification code that is given to the user via a phone call to the phone number that is being hosted. | 
**VerificationDocumentSid** | **string** | Optional. The unique sid identifier of the Identity Document that represents the document for verifying ownership of the number to be hosted. Required when VerificationType is phone-bill. | 
**VerificationType** | **string** | Optional. The method used for verifying ownership of the number to be hosted. One of phone-call (default) or phone-bill. | 

### Return type

[**PreviewHostedNumbersHostedNumberOrder**](PreviewHostedNumbersHostedNumberOrder.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded, 
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateInstalledAddOn

> PreviewMarketplaceInstalledAddOn UpdateInstalledAddOn(ctx, Sidoptional)



Update an Add-on installation for the Account specified.

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The SID of the InstalledAddOn resource to update. | 

### Other Parameters

Other parameters are passed through a pointer to a UpdateInstalledAddOnParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**Configuration** | [**map[string]interface{}**](map[string]interface{}.md) | Valid JSON object that conform to the configuration schema exposed by the associated AvailableAddOn resource. This is only required by Add-ons that need to be configured | 
**UniqueName** | **string** | An application-defined string that uniquely identifies the resource. This value must be unique within the Account. | 

### Return type

[**PreviewMarketplaceInstalledAddOn**](PreviewMarketplaceInstalledAddOn.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded, 
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateInstalledAddOnExtension

> PreviewMarketplaceInstalledAddOnInstalledAddOnExtension UpdateInstalledAddOnExtension(ctx, InstalledAddOnSidSidoptional)



Update an Extension for an Add-on installation.

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**InstalledAddOnSid** | **string** | The SID of the InstalledAddOn resource with the extension to update. | 
**Sid** | **string** | The SID of the InstalledAddOn Extension resource to update. | 

### Other Parameters

Other parameters are passed through a pointer to a UpdateInstalledAddOnExtensionParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**Enabled** | **bool** | Whether the Extension should be invoked. | 

### Return type

[**PreviewMarketplaceInstalledAddOnInstalledAddOnExtension**](PreviewMarketplaceInstalledAddOnInstalledAddOnExtension.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded, 
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateKey

> PreviewDeployedDevicesFleetKey UpdateKey(ctx, FleetSidSidoptional)



Update the given properties of a specific Key credential in the Fleet, giving it a friendly name or assigning to a Device.

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**FleetSid** | **string** |  | 
**Sid** | **string** | Provides a 34 character string that uniquely identifies the requested Key credential resource. | 

### Other Parameters

Other parameters are passed through a pointer to a UpdateKeyParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**DeviceSid** | **string** | Provides the unique string identifier of an existing Device to become authenticated with this Key credential. | 
**FriendlyName** | **string** | Provides a human readable descriptive text for this Key credential, up to 256 characters long. | 

### Return type

[**PreviewDeployedDevicesFleetKey**](PreviewDeployedDevicesFleetKey.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded, 
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateModelBuild

> PreviewUnderstandAssistantModelBuild UpdateModelBuild(ctx, AssistantSidSidoptional)



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string** |  | 
**Sid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a UpdateModelBuildParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**UniqueName** | **string** | A user-provided string that uniquely identifies this resource as an alternative to the sid. Unique up to 64 characters long. For example: v0.1 | 

### Return type

[**PreviewUnderstandAssistantModelBuild**](PreviewUnderstandAssistantModelBuild.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded, 
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateQuery

> PreviewUnderstandAssistantQuery UpdateQuery(ctx, AssistantSidSidoptional)



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string** | The unique ID of the parent Assistant. | 
**Sid** | **string** | A 34 character string that uniquely identifies this resource. | 

### Other Parameters

Other parameters are passed through a pointer to a UpdateQueryParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**SampleSid** | **string** | An optional reference to the Sample created from this query. | 
**Status** | **string** | A string that described the query status. The values can be: pending_review, reviewed, discarded | 

### Return type

[**PreviewUnderstandAssistantQuery**](PreviewUnderstandAssistantQuery.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded, 
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRatePlan

> PreviewWirelessRatePlan UpdateRatePlan(ctx, Sidoptional)



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a UpdateRatePlanParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**FriendlyName** | **string** |  | 
**UniqueName** | **string** |  | 

### Return type

[**PreviewWirelessRatePlan**](PreviewWirelessRatePlan.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded, 
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSample

> PreviewUnderstandAssistantTaskSample UpdateSample(ctx, AssistantSidTaskSidSidoptional)



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string** | The unique ID of the Assistant. | 
**TaskSid** | **string** | The unique ID of the Task associated with this Sample. | 
**Sid** | **string** | A 34 character string that uniquely identifies this resource. | 

### Other Parameters

Other parameters are passed through a pointer to a UpdateSampleParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**Language** | **string** | An ISO language-country string of the sample. | 
**SourceChannel** | **string** | The communication channel the sample was captured. It can be: *voice*, *sms*, *chat*, *alexa*, *google-assistant*, or *slack*. If not included the value will be null | 
**TaggedText** | **string** | The text example of how end-users may express this task. The sample may contain Field tag blocks. | 

### Return type

[**PreviewUnderstandAssistantTaskSample**](PreviewUnderstandAssistantTaskSample.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded, 
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateService

> PreviewSyncService UpdateService(ctx, Sidoptional)



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a UpdateServiceParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**AclEnabled** | **bool** |  | 
**FriendlyName** | **string** |  | 
**ReachabilityWebhooksEnabled** | **bool** |  | 
**WebhookUrl** | **string** |  | 

### Return type

[**PreviewSyncService**](PreviewSyncService.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded, 
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSim

> PreviewWirelessSim UpdateSim(ctx, Sidoptional)



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a UpdateSimParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**CallbackMethod** | **string** |  | 
**CallbackUrl** | **string** |  | 
**CommandsCallbackMethod** | **string** |  | 
**CommandsCallbackUrl** | **string** |  | 
**FriendlyName** | **string** |  | 
**RatePlan** | **string** |  | 
**SmsFallbackMethod** | **string** |  | 
**SmsFallbackUrl** | **string** |  | 
**SmsMethod** | **string** |  | 
**SmsUrl** | **string** |  | 
**Status** | **string** |  | 
**UniqueName** | **string** |  | 
**VoiceFallbackMethod** | **string** |  | 
**VoiceFallbackUrl** | **string** |  | 
**VoiceMethod** | **string** |  | 
**VoiceUrl** | **string** |  | 

### Return type

[**PreviewWirelessSim**](PreviewWirelessSim.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded, 
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateStyleSheet

> PreviewUnderstandAssistantStyleSheet UpdateStyleSheet(ctx, AssistantSidoptional)



Updates the style sheet for an assistant identified by {AssistantSid} or {AssistantUniqueName}.

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string** | The unique ID of the Assistant | 

### Other Parameters

Other parameters are passed through a pointer to a UpdateStyleSheetParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**StyleSheet** | [**map[string]interface{}**](map[string]interface{}.md) | The JSON Style sheet string | 

### Return type

[**PreviewUnderstandAssistantStyleSheet**](PreviewUnderstandAssistantStyleSheet.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded, 
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSyncListItem

> PreviewSyncServiceSyncListSyncListItem UpdateSyncListItem(ctx, ServiceSidListSidIndexoptional)



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** |  | 
**ListSid** | **string** |  | 
**Index** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a UpdateSyncListItemParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**IfMatch** | **string** | The If-Match HTTP request header | 
**Data** | [**map[string]interface{}**](map[string]interface{}.md) |  | 

### Return type

[**PreviewSyncServiceSyncListSyncListItem**](PreviewSyncServiceSyncListSyncListItem.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded, 
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSyncListPermission

> PreviewSyncServiceSyncListSyncListPermission UpdateSyncListPermission(ctx, ServiceSidListSidIdentityoptional)



Update an identity's access to a specific Sync List.

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The unique SID identifier of the Sync Service Instance. | 
**ListSid** | **string** | Identifier of the Sync List. Either a SID or a unique name. | 
**Identity** | **string** | Arbitrary string identifier representing a human user associated with an FPA token, assigned by the developer. | 

### Other Parameters

Other parameters are passed through a pointer to a UpdateSyncListPermissionParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**Manage** | **bool** | Boolean flag specifying whether the identity can delete the Sync List. | 
**Read** | **bool** | Boolean flag specifying whether the identity can read the Sync List. | 
**Write** | **bool** | Boolean flag specifying whether the identity can create, update and delete Items of the Sync List. | 

### Return type

[**PreviewSyncServiceSyncListSyncListPermission**](PreviewSyncServiceSyncListSyncListPermission.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded, 
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSyncMapItem

> PreviewSyncServiceSyncMapSyncMapItem UpdateSyncMapItem(ctx, ServiceSidMapSidKeyoptional)



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** |  | 
**MapSid** | **string** |  | 
**Key** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a UpdateSyncMapItemParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**IfMatch** | **string** | The If-Match HTTP request header | 
**Data** | [**map[string]interface{}**](map[string]interface{}.md) |  | 

### Return type

[**PreviewSyncServiceSyncMapSyncMapItem**](PreviewSyncServiceSyncMapSyncMapItem.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded, 
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSyncMapPermission

> PreviewSyncServiceSyncMapSyncMapPermission UpdateSyncMapPermission(ctx, ServiceSidMapSidIdentityoptional)



Update an identity's access to a specific Sync Map.

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The unique SID identifier of the Sync Service Instance. | 
**MapSid** | **string** | Identifier of the Sync Map. Either a SID or a unique name. | 
**Identity** | **string** | Arbitrary string identifier representing a human user associated with an FPA token, assigned by the developer. | 

### Other Parameters

Other parameters are passed through a pointer to a UpdateSyncMapPermissionParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**Manage** | **bool** | Boolean flag specifying whether the identity can delete the Sync Map. | 
**Read** | **bool** | Boolean flag specifying whether the identity can read the Sync Map. | 
**Write** | **bool** | Boolean flag specifying whether the identity can create, update and delete Items of the Sync Map. | 

### Return type

[**PreviewSyncServiceSyncMapSyncMapPermission**](PreviewSyncServiceSyncMapSyncMapPermission.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded, 
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTask

> PreviewUnderstandAssistantTask UpdateTask(ctx, AssistantSidSidoptional)



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string** | The unique ID of the Assistant. | 
**Sid** | **string** | A 34 character string that uniquely identifies this resource. | 

### Other Parameters

Other parameters are passed through a pointer to a UpdateTaskParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**Actions** | [**map[string]interface{}**](map[string]interface{}.md) | A user-provided JSON object encoded as a string to specify the actions for this task. It is optional and non-unique. | 
**ActionsUrl** | **string** | User-provided HTTP endpoint where from the assistant fetches actions | 
**FriendlyName** | **string** | A user-provided string that identifies this resource. It is non-unique and can up to 255 characters long. | 
**UniqueName** | **string** | A user-provided string that uniquely identifies this resource as an alternative to the sid. Unique up to 64 characters long. | 

### Return type

[**PreviewUnderstandAssistantTask**](PreviewUnderstandAssistantTask.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded, 
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTaskActions

> PreviewUnderstandAssistantTaskTaskActions UpdateTaskActions(ctx, AssistantSidTaskSidoptional)



Updates the actions of an Task identified by {TaskSid} or {TaskUniqueName}.

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AssistantSid** | **string** | The unique ID of the parent Assistant. | 
**TaskSid** | **string** | The unique ID of the Task. | 

### Other Parameters

Other parameters are passed through a pointer to a UpdateTaskActionsParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**Actions** | [**map[string]interface{}**](map[string]interface{}.md) | The JSON actions that instruct the Assistant how to perform this task. | 

### Return type

[**PreviewUnderstandAssistantTaskTaskActions**](PreviewUnderstandAssistantTaskTaskActions.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded, 
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

