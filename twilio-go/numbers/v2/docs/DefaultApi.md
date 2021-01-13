# \DefaultApi

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
 **optional** | ***CreateBundleOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateBundleOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **email** | **optional.String**| The email address that will receive updates when the Bundle resource changes status. | 
 **endUserType** | **optional.String**| The type of End User of the Bundle resource. | 
 **friendlyName** | **optional.String**| The string that you assigned to describe the resource. | 
 **isoCountry** | **optional.String**| The ISO country code of the Bundle&#39;s phone number country ownership request. | 
 **numberType** | **optional.String**| The type of phone number of the Bundle&#39;s ownership request. | 
 **regulationSid** | **optional.String**| The unique string of a regulation that is associated to the Bundle resource. | 
 **statusCallback** | **optional.String**| The URL we call to inform your application of status changes. | 

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
 **optional** | ***CreateEndUserOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateEndUserOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **attributes** | [**optional.Interface of map[string]interface{}**](map[string]interface{}.md)| The set of parameters that are the attributes of the End User resource which are derived End User Types. | 
 **friendlyName** | **optional.String**| The string that you assigned to describe the resource. | 
 **type_** | **optional.String**| The type of end user of the Bundle resource - can be &#x60;individual&#x60; or &#x60;business&#x60;. | 

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

> NumbersV2RegulatoryComplianceBundleEvaluation CreateEvaluation(ctx, bundleSid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bundleSid** | **string**|  | 

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

> NumbersV2RegulatoryComplianceBundleItemAssignment CreateItemAssignment(ctx, bundleSid, optional)



Create a new Assigned Item.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bundleSid** | **string**| The unique string that we created to identify the Bundle resource. | 
 **optional** | ***CreateItemAssignmentOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateItemAssignmentOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **objectSid** | **optional.String**| The SID of an object bag that holds information of the different items. | 

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
 **optional** | ***CreateSupportingDocumentOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateSupportingDocumentOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **attributes** | [**optional.Interface of map[string]interface{}**](map[string]interface{}.md)| The set of parameters that are the attributes of the Supporting Documents resource which are derived Supporting Document Types. | 
 **friendlyName** | **optional.String**| The string that you assigned to describe the resource. | 
 **type_** | **optional.String**| The type of the Supporting Document. | 

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

> DeleteBundle(ctx, sid)



Delete a specific Bundle.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sid** | **string**| The unique string that we created to identify the Bundle resource. | 

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

> DeleteEndUser(ctx, sid)



Delete a specific End User.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sid** | **string**| The unique string created by Twilio to identify the End User resource. | 

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

> DeleteItemAssignment(ctx, bundleSid, sid)



Remove an Assignment Item Instance.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bundleSid** | **string**| The unique string that we created to identify the Bundle resource. | 
**sid** | **string**| The unique string that we created to identify the Identity resource. | 

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

> DeleteSupportingDocument(ctx, sid)



Delete a specific Supporting Document.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sid** | **string**| The unique string created by Twilio to identify the Supporting Document resource. | 

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

> NumbersV2RegulatoryComplianceBundle FetchBundle(ctx, sid)



Fetch a specific Bundle instance.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sid** | **string**| The unique string that we created to identify the Bundle resource. | 

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

> NumbersV2RegulatoryComplianceEndUser FetchEndUser(ctx, sid)



Fetch specific End User Instance.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sid** | **string**| The unique string created by Twilio to identify the End User resource. | 

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

> NumbersV2RegulatoryComplianceEndUserType FetchEndUserType(ctx, sid)



Fetch a specific End-User Type Instance.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sid** | **string**| The unique string that identifies the End-User Type resource. | 

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

> NumbersV2RegulatoryComplianceBundleEvaluation FetchEvaluation(ctx, bundleSid, sid)



Fetch specific Evaluation Instance.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bundleSid** | **string**| The unique string that we created to identify the Bundle resource. | 
**sid** | **string**| The unique string that identifies the Evaluation resource. | 

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

> NumbersV2RegulatoryComplianceBundleItemAssignment FetchItemAssignment(ctx, bundleSid, sid)



Fetch specific Assigned Item Instance.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bundleSid** | **string**| The unique string that we created to identify the Bundle resource. | 
**sid** | **string**| The unique string that we created to identify the Identity resource. | 

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

> NumbersV2RegulatoryComplianceRegulation FetchRegulation(ctx, sid)



Fetch specific Regulation Instance.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sid** | **string**| The unique string that identifies the Regulation resource. | 

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

> NumbersV2RegulatoryComplianceSupportingDocument FetchSupportingDocument(ctx, sid)



Fetch specific Supporting Document Instance.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sid** | **string**| The unique string created by Twilio to identify the Supporting Document resource. | 

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

> NumbersV2RegulatoryComplianceSupportingDocumentType FetchSupportingDocumentType(ctx, sid)



Fetch a specific Supporting Document Type Instance.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sid** | **string**| The unique string that identifies the Supporting Document Type resource. | 

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

> NumbersV2RegulatoryComplianceBundleReadResponse ListBundle(ctx, optional)



Retrieve a list of all Bundles for an account.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListBundleOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListBundleOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **status** | **optional.String**| The verification status of the Bundle resource. | 
 **friendlyName** | **optional.String**| The string that you assigned to describe the resource. | 
 **regulationSid** | **optional.String**| The unique string of a regulation that is associated to the Bundle resource. | 
 **isoCountry** | **optional.String**| The ISO country code of the Bundle&#39;s phone number country ownership request. | 
 **numberType** | **optional.String**| The type of phone number of the Bundle&#39;s ownership request. | 
 **pageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**NumbersV2RegulatoryComplianceBundleReadResponse**](numbers_v2_regulatory_compliance_bundleReadResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListEndUser

> NumbersV2RegulatoryComplianceEndUserReadResponse ListEndUser(ctx, optional)



Retrieve a list of all End User for an account.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListEndUserOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListEndUserOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**NumbersV2RegulatoryComplianceEndUserReadResponse**](numbers_v2_regulatory_compliance_end_userReadResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListEndUserType

> NumbersV2RegulatoryComplianceEndUserTypeReadResponse ListEndUserType(ctx, optional)



Retrieve a list of all End-User Types.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListEndUserTypeOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListEndUserTypeOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**NumbersV2RegulatoryComplianceEndUserTypeReadResponse**](numbers_v2_regulatory_compliance_end_user_typeReadResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListEvaluation

> NumbersV2RegulatoryComplianceBundleEvaluationReadResponse ListEvaluation(ctx, bundleSid, optional)



Retrieve a list of Evaluations associated to the Bundle resource.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bundleSid** | **string**|  | 
 **optional** | ***ListEvaluationOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListEvaluationOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**NumbersV2RegulatoryComplianceBundleEvaluationReadResponse**](numbers_v2_regulatory_compliance_bundle_evaluationReadResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListItemAssignment

> NumbersV2RegulatoryComplianceBundleItemAssignmentReadResponse ListItemAssignment(ctx, bundleSid, optional)



Retrieve a list of all Assigned Items for an account.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bundleSid** | **string**| The unique string that we created to identify the Bundle resource. | 
 **optional** | ***ListItemAssignmentOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListItemAssignmentOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**NumbersV2RegulatoryComplianceBundleItemAssignmentReadResponse**](numbers_v2_regulatory_compliance_bundle_item_assignmentReadResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRegulation

> NumbersV2RegulatoryComplianceRegulationReadResponse ListRegulation(ctx, optional)



Retrieve a list of all Regulations.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListRegulationOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListRegulationOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **endUserType** | **optional.String**| The type of End User the regulation requires - can be &#x60;individual&#x60; or &#x60;business&#x60;. | 
 **isoCountry** | **optional.String**| The ISO country code of the phone number&#39;s country. | 
 **numberType** | **optional.String**| The type of phone number that the regulatory requiremnt is restricting. | 
 **pageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**NumbersV2RegulatoryComplianceRegulationReadResponse**](numbers_v2_regulatory_compliance_regulationReadResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSupportingDocument

> NumbersV2RegulatoryComplianceSupportingDocumentReadResponse ListSupportingDocument(ctx, optional)



Retrieve a list of all Supporting Document for an account.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListSupportingDocumentOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListSupportingDocumentOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**NumbersV2RegulatoryComplianceSupportingDocumentReadResponse**](numbers_v2_regulatory_compliance_supporting_documentReadResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSupportingDocumentType

> NumbersV2RegulatoryComplianceSupportingDocumentTypeReadResponse ListSupportingDocumentType(ctx, optional)



Retrieve a list of all Supporting Document Types.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListSupportingDocumentTypeOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListSupportingDocumentTypeOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**NumbersV2RegulatoryComplianceSupportingDocumentTypeReadResponse**](numbers_v2_regulatory_compliance_supporting_document_typeReadResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateBundle

> NumbersV2RegulatoryComplianceBundle UpdateBundle(ctx, sid, optional)



Updates a Bundle in an account.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sid** | **string**| The unique string that we created to identify the Bundle resource. | 
 **optional** | ***UpdateBundleOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateBundleOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **email** | **optional.String**| The email address that will receive updates when the Bundle resource changes status. | 
 **friendlyName** | **optional.String**| The string that you assigned to describe the resource. | 
 **status** | **optional.String**| The verification status of the Bundle resource. | 
 **statusCallback** | **optional.String**| The URL we call to inform your application of status changes. | 

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

> NumbersV2RegulatoryComplianceEndUser UpdateEndUser(ctx, sid, optional)



Update an existing End User.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sid** | **string**| The unique string created by Twilio to identify the End User resource. | 
 **optional** | ***UpdateEndUserOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateEndUserOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **attributes** | [**optional.Interface of map[string]interface{}**](map[string]interface{}.md)| The set of parameters that are the attributes of the End User resource which are derived End User Types. | 
 **friendlyName** | **optional.String**| The string that you assigned to describe the resource. | 

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

> NumbersV2RegulatoryComplianceSupportingDocument UpdateSupportingDocument(ctx, sid, optional)



Update an existing Supporting Document.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sid** | **string**| The unique string created by Twilio to identify the Supporting Document resource. | 
 **optional** | ***UpdateSupportingDocumentOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateSupportingDocumentOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **attributes** | [**optional.Interface of map[string]interface{}**](map[string]interface{}.md)| The set of parameters that are the attributes of the Supporting Document resource which are derived Supporting Document Types. | 
 **friendlyName** | **optional.String**| The string that you assigned to describe the resource. | 

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

