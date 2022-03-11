# RegulatoryComplianceBundlesItemAssignmentsApi

All URIs are relative to *https://numbers.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateItemAssignment**](RegulatoryComplianceBundlesItemAssignmentsApi.md#CreateItemAssignment) | **Post** /v2/RegulatoryCompliance/Bundles/{BundleSid}/ItemAssignments | 
[**DeleteItemAssignment**](RegulatoryComplianceBundlesItemAssignmentsApi.md#DeleteItemAssignment) | **Delete** /v2/RegulatoryCompliance/Bundles/{BundleSid}/ItemAssignments/{Sid} | 
[**FetchItemAssignment**](RegulatoryComplianceBundlesItemAssignmentsApi.md#FetchItemAssignment) | **Get** /v2/RegulatoryCompliance/Bundles/{BundleSid}/ItemAssignments/{Sid} | 
[**ListItemAssignment**](RegulatoryComplianceBundlesItemAssignmentsApi.md#ListItemAssignment) | **Get** /v2/RegulatoryCompliance/Bundles/{BundleSid}/ItemAssignments | 



## CreateItemAssignment

> NumbersV2ItemAssignment CreateItemAssignment(ctx, BundleSidoptional)



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

[**NumbersV2ItemAssignment**](NumbersV2ItemAssignment.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

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


## FetchItemAssignment

> NumbersV2ItemAssignment FetchItemAssignment(ctx, BundleSidSid)



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

[**NumbersV2ItemAssignment**](NumbersV2ItemAssignment.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListItemAssignment

> []NumbersV2ItemAssignment ListItemAssignment(ctx, BundleSidoptional)



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
**Limit** | **int** | Max number of records to return.

### Return type

[**[]NumbersV2ItemAssignment**](NumbersV2ItemAssignment.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

