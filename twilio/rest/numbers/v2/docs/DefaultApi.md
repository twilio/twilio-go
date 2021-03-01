# DefaultApi

All URIs are relative to *http://localhost*

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

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateBundleRequest** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateBundleRequest struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**Email** | **String**| The email address that will receive updates when the Bundle resource changes status. | 
**EndUserType** | **String**| The type of End User of the Bundle resource. | 
**FriendlyName** | **String**| The string that you assigned to describe the resource. | 
**IsoCountry** | **String**| The ISO country code of the Bundle&#39;s phone number country ownership request. | 
**NumberType** | **String**| The type of phone number of the Bundle&#39;s ownership request. | 
**RegulationSid** | **String**| The unique string of a regulation that is associated to the Bundle resource. | 
**StatusCallback** | **String**| The URL we call to inform your application of status changes. | 

### Return type

[**NumbersV2RegulatoryComplianceBundle**](numbers.v2.regulatory_compliance.bundle.md)

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

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateEndUserRequest** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateEndUserRequest struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**Attributes** | [**map[string]interface{}**](map[string]interface{}.md)| The set of parameters that are the attributes of the End User resource which are derived End User Types. | 
**FriendlyName** | **String**| The string that you assigned to describe the resource. | 
**Type** | **String**| The type of end user of the Bundle resource - can be &#x60;individual&#x60; or &#x60;business&#x60;. | 

### Return type

[**NumbersV2RegulatoryComplianceEndUser**](numbers.v2.regulatory_compliance.end_user.md)

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



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**BundleSid** | **string**|  | 

### Return type

[**NumbersV2RegulatoryComplianceBundleEvaluation**](numbers.v2.regulatory_compliance.bundle.evaluation.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateItemAssignment

> NumbersV2RegulatoryComplianceBundleItemAssignment CreateItemAssignment(ctx, BundleSid, optional)



Create a new Assigned Item.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**BundleSid** | **string**| The unique string that we created to identify the Bundle resource. | 
 **optional** | ***CreateItemAssignmentRequest** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateItemAssignmentRequest struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ObjectSid** | **String**| The SID of an object bag that holds information of the different items. | 

### Return type

[**NumbersV2RegulatoryComplianceBundleItemAssignment**](numbers.v2.regulatory_compliance.bundle.item_assignment.md)

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

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateSupportingDocumentRequest** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateSupportingDocumentRequest struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**Attributes** | [**map[string]interface{}**](map[string]interface{}.md)| The set of parameters that are the attributes of the Supporting Documents resource which are derived Supporting Document Types. | 
**FriendlyName** | **String**| The string that you assigned to describe the resource. | 
**Type** | **String**| The type of the Supporting Document. | 

### Return type

[**NumbersV2RegulatoryComplianceSupportingDocument**](numbers.v2.regulatory_compliance.supporting_document.md)

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

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string**| The unique string that we created to identify the Bundle resource. | 

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

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string**| The unique string created by Twilio to identify the End User resource. | 

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

> DeleteItemAssignment(ctx, BundleSid, Sid)



Remove an Assignment Item Instance.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**BundleSid** | **string**| The unique string that we created to identify the Bundle resource. | 
**Sid** | **string**| The unique string that we created to identify the Identity resource. | 

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

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string**| The unique string created by Twilio to identify the Supporting Document resource. | 

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

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string**| The unique string that we created to identify the Bundle resource. | 

### Return type

[**NumbersV2RegulatoryComplianceBundle**](numbers.v2.regulatory_compliance.bundle.md)

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

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string**| The unique string created by Twilio to identify the End User resource. | 

### Return type

[**NumbersV2RegulatoryComplianceEndUser**](numbers.v2.regulatory_compliance.end_user.md)

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

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string**| The unique string that identifies the End-User Type resource. | 

### Return type

[**NumbersV2RegulatoryComplianceEndUserType**](numbers.v2.regulatory_compliance.end_user_type.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchEvaluation

> NumbersV2RegulatoryComplianceBundleEvaluation FetchEvaluation(ctx, BundleSid, Sid)



Fetch specific Evaluation Instance.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**BundleSid** | **string**| The unique string that we created to identify the Bundle resource. | 
**Sid** | **string**| The unique string that identifies the Evaluation resource. | 

### Return type

[**NumbersV2RegulatoryComplianceBundleEvaluation**](numbers.v2.regulatory_compliance.bundle.evaluation.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchItemAssignment

> NumbersV2RegulatoryComplianceBundleItemAssignment FetchItemAssignment(ctx, BundleSid, Sid)



Fetch specific Assigned Item Instance.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**BundleSid** | **string**| The unique string that we created to identify the Bundle resource. | 
**Sid** | **string**| The unique string that we created to identify the Identity resource. | 

### Return type

[**NumbersV2RegulatoryComplianceBundleItemAssignment**](numbers.v2.regulatory_compliance.bundle.item_assignment.md)

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

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string**| The unique string that identifies the Regulation resource. | 

### Return type

[**NumbersV2RegulatoryComplianceRegulation**](numbers.v2.regulatory_compliance.regulation.md)

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

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string**| The unique string created by Twilio to identify the Supporting Document resource. | 

### Return type

[**NumbersV2RegulatoryComplianceSupportingDocument**](numbers.v2.regulatory_compliance.supporting_document.md)

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

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string**| The unique string that identifies the Supporting Document Type resource. | 

### Return type

[**NumbersV2RegulatoryComplianceSupportingDocumentType**](numbers.v2.regulatory_compliance.supporting_document_type.md)

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

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListBundleRequest** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListBundleRequest struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**Status** | **String**| The verification status of the Bundle resource. | 
**FriendlyName** | **String**| The string that you assigned to describe the resource. | 
**RegulationSid** | **String**| The unique string of a regulation that is associated to the Bundle resource. | 
**IsoCountry** | **String**| The ISO country code of the Bundle&#39;s phone number country ownership request. | 
**NumberType** | **String**| The type of phone number of the Bundle&#39;s ownership request. | 
**PageSize** | **Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

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

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListEndUserRequest** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListEndUserRequest struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**PageSize** | **Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

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

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListEndUserTypeRequest** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListEndUserTypeRequest struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**PageSize** | **Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

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

> ListEvaluationResponse ListEvaluation(ctx, BundleSid, optional)



Retrieve a list of Evaluations associated to the Bundle resource.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**BundleSid** | **string**|  | 
 **optional** | ***ListEvaluationRequest** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListEvaluationRequest struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**PageSize** | **Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

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

> ListItemAssignmentResponse ListItemAssignment(ctx, BundleSid, optional)



Retrieve a list of all Assigned Items for an account.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**BundleSid** | **string**| The unique string that we created to identify the Bundle resource. | 
 **optional** | ***ListItemAssignmentRequest** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListItemAssignmentRequest struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**PageSize** | **Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

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

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListRegulationRequest** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListRegulationRequest struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**EndUserType** | **String**| The type of End User the regulation requires - can be &#x60;individual&#x60; or &#x60;business&#x60;. | 
**IsoCountry** | **String**| The ISO country code of the phone number&#39;s country. | 
**NumberType** | **String**| The type of phone number that the regulatory requiremnt is restricting. | 
**PageSize** | **Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

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

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListSupportingDocumentRequest** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListSupportingDocumentRequest struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**PageSize** | **Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

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

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListSupportingDocumentTypeRequest** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListSupportingDocumentTypeRequest struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**PageSize** | **Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

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

> NumbersV2RegulatoryComplianceBundle UpdateBundle(ctx, Sid, optional)



Updates a Bundle in an account.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string**| The unique string that we created to identify the Bundle resource. | 
 **optional** | ***UpdateBundleRequest** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateBundleRequest struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**Email** | **String**| The email address that will receive updates when the Bundle resource changes status. | 
**FriendlyName** | **String**| The string that you assigned to describe the resource. | 
**Status** | **String**| The verification status of the Bundle resource. | 
**StatusCallback** | **String**| The URL we call to inform your application of status changes. | 

### Return type

[**NumbersV2RegulatoryComplianceBundle**](numbers.v2.regulatory_compliance.bundle.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateEndUser

> NumbersV2RegulatoryComplianceEndUser UpdateEndUser(ctx, Sid, optional)



Update an existing End User.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string**| The unique string created by Twilio to identify the End User resource. | 
 **optional** | ***UpdateEndUserRequest** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateEndUserRequest struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**Attributes** | [**map[string]interface{}**](map[string]interface{}.md)| The set of parameters that are the attributes of the End User resource which are derived End User Types. | 
**FriendlyName** | **String**| The string that you assigned to describe the resource. | 

### Return type

[**NumbersV2RegulatoryComplianceEndUser**](numbers.v2.regulatory_compliance.end_user.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSupportingDocument

> NumbersV2RegulatoryComplianceSupportingDocument UpdateSupportingDocument(ctx, Sid, optional)



Update an existing Supporting Document.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string**| The unique string created by Twilio to identify the Supporting Document resource. | 
 **optional** | ***UpdateSupportingDocumentRequest** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateSupportingDocumentRequest struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**Attributes** | [**map[string]interface{}**](map[string]interface{}.md)| The set of parameters that are the attributes of the Supporting Document resource which are derived Supporting Document Types. | 
**FriendlyName** | **String**| The string that you assigned to describe the resource. | 

### Return type

[**NumbersV2RegulatoryComplianceSupportingDocument**](numbers.v2.regulatory_compliance.supporting_document.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

