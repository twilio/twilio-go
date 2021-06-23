# DefaultApi

All URIs are relative to *https://numbers.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBundle**](DefaultApi.md#CreateBundle) | **Post** /v2/RegulatoryCompliance/Bundles | 
[**CreateEndUser**](DefaultApi.md#CreateEndUser) | **Post** /v2/RegulatoryCompliance/EndUsers | 
[**CreateEvaluation**](DefaultApi.md#CreateEvaluation) | **Post** /v2/RegulatoryCompliance/Bundles/{BundleSid}/Evaluations | 
[**CreateItemAssignment**](DefaultApi.md#CreateItemAssignment) | **Post** /v2/RegulatoryCompliance/Bundles/{BundleSid}/ItemAssignments | 
[**CreateSupportingDocument**](DefaultApi.md#CreateSupportingDocument) | **Post** /v2/RegulatoryCompliance/SupportingDocuments | 
[**DeleteBundle**](DefaultApi.md#DeleteBundle) | **Delete** /v2/RegulatoryCompliance/Bundles/{Sid} | 
[**DeleteEndUser**](DefaultApi.md#DeleteEndUser) | **Delete** /v2/RegulatoryCompliance/EndUsers/{Sid} | 
[**DeleteItemAssignment**](DefaultApi.md#DeleteItemAssignment) | **Delete** /v2/RegulatoryCompliance/Bundles/{BundleSid}/ItemAssignments/{Sid} | 
[**DeleteSupportingDocument**](DefaultApi.md#DeleteSupportingDocument) | **Delete** /v2/RegulatoryCompliance/SupportingDocuments/{Sid} | 
[**FetchBundle**](DefaultApi.md#FetchBundle) | **Get** /v2/RegulatoryCompliance/Bundles/{Sid} | 
[**FetchEndUser**](DefaultApi.md#FetchEndUser) | **Get** /v2/RegulatoryCompliance/EndUsers/{Sid} | 
[**FetchEndUserType**](DefaultApi.md#FetchEndUserType) | **Get** /v2/RegulatoryCompliance/EndUserTypes/{Sid} | 
[**FetchEvaluation**](DefaultApi.md#FetchEvaluation) | **Get** /v2/RegulatoryCompliance/Bundles/{BundleSid}/Evaluations/{Sid} | 
[**FetchItemAssignment**](DefaultApi.md#FetchItemAssignment) | **Get** /v2/RegulatoryCompliance/Bundles/{BundleSid}/ItemAssignments/{Sid} | 
[**FetchRegulation**](DefaultApi.md#FetchRegulation) | **Get** /v2/RegulatoryCompliance/Regulations/{Sid} | 
[**FetchSupportingDocument**](DefaultApi.md#FetchSupportingDocument) | **Get** /v2/RegulatoryCompliance/SupportingDocuments/{Sid} | 
[**FetchSupportingDocumentType**](DefaultApi.md#FetchSupportingDocumentType) | **Get** /v2/RegulatoryCompliance/SupportingDocumentTypes/{Sid} | 
[**ListBundle**](DefaultApi.md#ListBundle) | **Get** /v2/RegulatoryCompliance/Bundles | 
[**ListEndUser**](DefaultApi.md#ListEndUser) | **Get** /v2/RegulatoryCompliance/EndUsers | 
[**ListEndUserType**](DefaultApi.md#ListEndUserType) | **Get** /v2/RegulatoryCompliance/EndUserTypes | 
[**ListEvaluation**](DefaultApi.md#ListEvaluation) | **Get** /v2/RegulatoryCompliance/Bundles/{BundleSid}/Evaluations | 
[**ListItemAssignment**](DefaultApi.md#ListItemAssignment) | **Get** /v2/RegulatoryCompliance/Bundles/{BundleSid}/ItemAssignments | 
[**ListRegulation**](DefaultApi.md#ListRegulation) | **Get** /v2/RegulatoryCompliance/Regulations | 
[**ListSupportingDocument**](DefaultApi.md#ListSupportingDocument) | **Get** /v2/RegulatoryCompliance/SupportingDocuments | 
[**ListSupportingDocumentType**](DefaultApi.md#ListSupportingDocumentType) | **Get** /v2/RegulatoryCompliance/SupportingDocumentTypes | 
[**UpdateBundle**](DefaultApi.md#UpdateBundle) | **Post** /v2/RegulatoryCompliance/Bundles/{Sid} | 
[**UpdateEndUser**](DefaultApi.md#UpdateEndUser) | **Post** /v2/RegulatoryCompliance/EndUsers/{Sid} | 
[**UpdateSupportingDocument**](DefaultApi.md#UpdateSupportingDocument) | **Post** /v2/RegulatoryCompliance/SupportingDocuments/{Sid} | 



## CreateBundle

> NumbersV2RegulatoryComplianceBundle CreateBundle(ctx, optional)



Create a new Bundle.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateBundleParams struct


Name | Type | Description
------------- | ------------- | -------------
**Email** | **string** | The email address that will receive updates when the Bundle resource changes status.
**EndUserType** | **string** | The type of End User of the Bundle resource.
**FriendlyName** | **string** | The string that you assigned to describe the resource.
**IsoCountry** | **string** | The ISO country code of the Bundle&#39;s phone number country ownership request.
**NumberType** | **string** | The type of phone number of the Bundle&#39;s ownership request.
**RegulationSid** | **string** | The unique string of a regulation that is associated to the Bundle resource.
**StatusCallback** | **string** | The URL we call to inform your application of status changes.

### Return type

[**NumbersV2RegulatoryComplianceBundle**](NumbersV2RegulatoryComplianceBundle.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateEndUser

> NumbersV2RegulatoryComplianceEndUser CreateEndUser(ctx, optional)



Create a new End User.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateEndUserParams struct


Name | Type | Description
------------- | ------------- | -------------
**Attributes** | [**map[string]interface{}**](map[string]interface{}.md) | The set of parameters that are the attributes of the End User resource which are derived End User Types.
**FriendlyName** | **string** | The string that you assigned to describe the resource.
**Type** | **string** | The type of end user of the Bundle resource - can be &#x60;individual&#x60; or &#x60;business&#x60;.

### Return type

[**NumbersV2RegulatoryComplianceEndUser**](NumbersV2RegulatoryComplianceEndUser.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateEvaluation

> NumbersV2RegulatoryComplianceBundleEvaluation CreateEvaluation(ctx, BundleSid)



Creates an evaluation for a bundle

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**BundleSid** | **string** | The unique string that identifies the Bundle resource.

### Other Parameters

Other parameters are passed through a pointer to a CreateEvaluationParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**NumbersV2RegulatoryComplianceBundleEvaluation**](NumbersV2RegulatoryComplianceBundleEvaluation.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateItemAssignment

> NumbersV2RegulatoryComplianceBundleItemAssignment CreateItemAssignment(ctx, BundleSidoptional)



Create a new Assigned Item.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**BundleSid** | **string** | The unique string that we created to identify the Bundle resource.

### Other Parameters

Other parameters are passed through a pointer to a CreateItemAssignmentParams struct


Name | Type | Description
------------- | ------------- | -------------
**ObjectSid** | **string** | The SID of an object bag that holds information of the different items.

### Return type

[**NumbersV2RegulatoryComplianceBundleItemAssignment**](NumbersV2RegulatoryComplianceBundleItemAssignment.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSupportingDocument

> NumbersV2RegulatoryComplianceSupportingDocument CreateSupportingDocument(ctx, optional)



Create a new Supporting Document.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateSupportingDocumentParams struct


Name | Type | Description
------------- | ------------- | -------------
**Attributes** | [**map[string]interface{}**](map[string]interface{}.md) | The set of parameters that are the attributes of the Supporting Documents resource which are derived Supporting Document Types.
**FriendlyName** | **string** | The string that you assigned to describe the resource.
**Type** | **string** | The type of the Supporting Document.

### Return type

[**NumbersV2RegulatoryComplianceSupportingDocument**](NumbersV2RegulatoryComplianceSupportingDocument.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteBundle

> DeleteBundle(ctx, Sid)



Delete a specific Bundle.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The unique string that we created to identify the Bundle resource.

### Other Parameters

Other parameters are passed through a pointer to a DeleteBundleParams struct


Name | Type | Description
------------- | ------------- | -------------

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


## DeleteEndUser

> DeleteEndUser(ctx, Sid)



Delete a specific End User.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The unique string created by Twilio to identify the End User resource.

### Other Parameters

Other parameters are passed through a pointer to a DeleteEndUserParams struct


Name | Type | Description
------------- | ------------- | -------------

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


## DeleteItemAssignment

> DeleteItemAssignment(ctx, BundleSidSid)



Remove an Assignment Item Instance.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**BundleSid** | **string** | The unique string that we created to identify the Bundle resource.
**Sid** | **string** | The unique string that we created to identify the Identity resource.

### Other Parameters

Other parameters are passed through a pointer to a DeleteItemAssignmentParams struct


Name | Type | Description
------------- | ------------- | -------------

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


## DeleteSupportingDocument

> DeleteSupportingDocument(ctx, Sid)



Delete a specific Supporting Document.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The unique string created by Twilio to identify the Supporting Document resource.

### Other Parameters

Other parameters are passed through a pointer to a DeleteSupportingDocumentParams struct


Name | Type | Description
------------- | ------------- | -------------

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


## FetchBundle

> NumbersV2RegulatoryComplianceBundle FetchBundle(ctx, Sid)



Fetch a specific Bundle instance.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The unique string that we created to identify the Bundle resource.

### Other Parameters

Other parameters are passed through a pointer to a FetchBundleParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**NumbersV2RegulatoryComplianceBundle**](NumbersV2RegulatoryComplianceBundle.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchEndUser

> NumbersV2RegulatoryComplianceEndUser FetchEndUser(ctx, Sid)



Fetch specific End User Instance.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The unique string created by Twilio to identify the End User resource.

### Other Parameters

Other parameters are passed through a pointer to a FetchEndUserParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**NumbersV2RegulatoryComplianceEndUser**](NumbersV2RegulatoryComplianceEndUser.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchEndUserType

> NumbersV2RegulatoryComplianceEndUserType FetchEndUserType(ctx, Sid)



Fetch a specific End-User Type Instance.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The unique string that identifies the End-User Type resource.

### Other Parameters

Other parameters are passed through a pointer to a FetchEndUserTypeParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**NumbersV2RegulatoryComplianceEndUserType**](NumbersV2RegulatoryComplianceEndUserType.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchEvaluation

> NumbersV2RegulatoryComplianceBundleEvaluation FetchEvaluation(ctx, BundleSidSid)



Fetch specific Evaluation Instance.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**BundleSid** | **string** | The unique string that we created to identify the Bundle resource.
**Sid** | **string** | The unique string that identifies the Evaluation resource.

### Other Parameters

Other parameters are passed through a pointer to a FetchEvaluationParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**NumbersV2RegulatoryComplianceBundleEvaluation**](NumbersV2RegulatoryComplianceBundleEvaluation.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchItemAssignment

> NumbersV2RegulatoryComplianceBundleItemAssignment FetchItemAssignment(ctx, BundleSidSid)



Fetch specific Assigned Item Instance.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**BundleSid** | **string** | The unique string that we created to identify the Bundle resource.
**Sid** | **string** | The unique string that we created to identify the Identity resource.

### Other Parameters

Other parameters are passed through a pointer to a FetchItemAssignmentParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**NumbersV2RegulatoryComplianceBundleItemAssignment**](NumbersV2RegulatoryComplianceBundleItemAssignment.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchRegulation

> NumbersV2RegulatoryComplianceRegulation FetchRegulation(ctx, Sid)



Fetch specific Regulation Instance.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The unique string that identifies the Regulation resource.

### Other Parameters

Other parameters are passed through a pointer to a FetchRegulationParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**NumbersV2RegulatoryComplianceRegulation**](NumbersV2RegulatoryComplianceRegulation.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchSupportingDocument

> NumbersV2RegulatoryComplianceSupportingDocument FetchSupportingDocument(ctx, Sid)



Fetch specific Supporting Document Instance.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The unique string created by Twilio to identify the Supporting Document resource.

### Other Parameters

Other parameters are passed through a pointer to a FetchSupportingDocumentParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**NumbersV2RegulatoryComplianceSupportingDocument**](NumbersV2RegulatoryComplianceSupportingDocument.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchSupportingDocumentType

> NumbersV2RegulatoryComplianceSupportingDocumentType FetchSupportingDocumentType(ctx, Sid)



Fetch a specific Supporting Document Type Instance.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The unique string that identifies the Supporting Document Type resource.

### Other Parameters

Other parameters are passed through a pointer to a FetchSupportingDocumentTypeParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**NumbersV2RegulatoryComplianceSupportingDocumentType**](NumbersV2RegulatoryComplianceSupportingDocumentType.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListBundle

> ListBundleResponse ListBundle(ctx, optional)



Retrieve a list of all Bundles for an account.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListBundleParams struct


Name | Type | Description
------------- | ------------- | -------------
**Status** | **string** | The verification status of the Bundle resource.
**FriendlyName** | **string** | The string that you assigned to describe the resource.
**RegulationSid** | **string** | The unique string of a regulation that is associated to the Bundle resource.
**IsoCountry** | **string** | The ISO country code of the Bundle&#39;s phone number country ownership request.
**NumberType** | **string** | The type of phone number of the Bundle&#39;s ownership request.
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListBundleResponse**](ListBundleResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListEndUser

> ListEndUserResponse ListEndUser(ctx, optional)



Retrieve a list of all End User for an account.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListEndUserParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListEndUserResponse**](ListEndUserResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListEndUserType

> ListEndUserTypeResponse ListEndUserType(ctx, optional)



Retrieve a list of all End-User Types.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListEndUserTypeParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListEndUserTypeResponse**](ListEndUserTypeResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListEvaluation

> ListEvaluationResponse ListEvaluation(ctx, BundleSidoptional)



Retrieve a list of Evaluations associated to the Bundle resource.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**BundleSid** | **string** | The unique string that identifies the Bundle resource.

### Other Parameters

Other parameters are passed through a pointer to a ListEvaluationParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListEvaluationResponse**](ListEvaluationResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListItemAssignment

> ListItemAssignmentResponse ListItemAssignment(ctx, BundleSidoptional)



Retrieve a list of all Assigned Items for an account.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**BundleSid** | **string** | The unique string that we created to identify the Bundle resource.

### Other Parameters

Other parameters are passed through a pointer to a ListItemAssignmentParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListItemAssignmentResponse**](ListItemAssignmentResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRegulation

> ListRegulationResponse ListRegulation(ctx, optional)



Retrieve a list of all Regulations.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListRegulationParams struct


Name | Type | Description
------------- | ------------- | -------------
**EndUserType** | **string** | The type of End User the regulation requires - can be &#x60;individual&#x60; or &#x60;business&#x60;.
**IsoCountry** | **string** | The ISO country code of the phone number&#39;s country.
**NumberType** | **string** | The type of phone number that the regulatory requiremnt is restricting.
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListRegulationResponse**](ListRegulationResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSupportingDocument

> ListSupportingDocumentResponse ListSupportingDocument(ctx, optional)



Retrieve a list of all Supporting Document for an account.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListSupportingDocumentParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListSupportingDocumentResponse**](ListSupportingDocumentResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSupportingDocumentType

> ListSupportingDocumentTypeResponse ListSupportingDocumentType(ctx, optional)



Retrieve a list of all Supporting Document Types.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListSupportingDocumentTypeParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListSupportingDocumentTypeResponse**](ListSupportingDocumentTypeResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateBundle

> NumbersV2RegulatoryComplianceBundle UpdateBundle(ctx, Sidoptional)



Updates a Bundle in an account.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The unique string that we created to identify the Bundle resource.

### Other Parameters

Other parameters are passed through a pointer to a UpdateBundleParams struct


Name | Type | Description
------------- | ------------- | -------------
**Email** | **string** | The email address that will receive updates when the Bundle resource changes status.
**FriendlyName** | **string** | The string that you assigned to describe the resource.
**Status** | **string** | The verification status of the Bundle resource.
**StatusCallback** | **string** | The URL we call to inform your application of status changes.

### Return type

[**NumbersV2RegulatoryComplianceBundle**](NumbersV2RegulatoryComplianceBundle.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateEndUser

> NumbersV2RegulatoryComplianceEndUser UpdateEndUser(ctx, Sidoptional)



Update an existing End User.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The unique string created by Twilio to identify the End User resource.

### Other Parameters

Other parameters are passed through a pointer to a UpdateEndUserParams struct


Name | Type | Description
------------- | ------------- | -------------
**Attributes** | [**map[string]interface{}**](map[string]interface{}.md) | The set of parameters that are the attributes of the End User resource which are derived End User Types.
**FriendlyName** | **string** | The string that you assigned to describe the resource.

### Return type

[**NumbersV2RegulatoryComplianceEndUser**](NumbersV2RegulatoryComplianceEndUser.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSupportingDocument

> NumbersV2RegulatoryComplianceSupportingDocument UpdateSupportingDocument(ctx, Sidoptional)



Update an existing Supporting Document.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The unique string created by Twilio to identify the Supporting Document resource.

### Other Parameters

Other parameters are passed through a pointer to a UpdateSupportingDocumentParams struct


Name | Type | Description
------------- | ------------- | -------------
**Attributes** | [**map[string]interface{}**](map[string]interface{}.md) | The set of parameters that are the attributes of the Supporting Document resource which are derived Supporting Document Types.
**FriendlyName** | **string** | The string that you assigned to describe the resource.

### Return type

[**NumbersV2RegulatoryComplianceSupportingDocument**](NumbersV2RegulatoryComplianceSupportingDocument.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

