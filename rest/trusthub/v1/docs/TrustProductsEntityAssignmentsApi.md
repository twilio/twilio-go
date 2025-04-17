# TrustProductsEntityAssignmentsApi

All URIs are relative to *https://trusthub.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTrustProductEntityAssignment**](TrustProductsEntityAssignmentsApi.md#CreateTrustProductEntityAssignment) | **Post** /v1/TrustProducts/{TrustProductSid}/EntityAssignments | Create a new Assigned Item.
[**DeleteTrustProductEntityAssignment**](TrustProductsEntityAssignmentsApi.md#DeleteTrustProductEntityAssignment) | **Delete** /v1/TrustProducts/{TrustProductSid}/EntityAssignments/{Sid} | Remove an Assignment Item Instance.
[**FetchTrustProductEntityAssignment**](TrustProductsEntityAssignmentsApi.md#FetchTrustProductEntityAssignment) | **Get** /v1/TrustProducts/{TrustProductSid}/EntityAssignments/{Sid} | Fetch specific Assigned Item Instance.
[**ListTrustProductEntityAssignment**](TrustProductsEntityAssignmentsApi.md#ListTrustProductEntityAssignment) | **Get** /v1/TrustProducts/{TrustProductSid}/EntityAssignments | Retrieve a list of all Assigned Items for an account.



## CreateTrustProductEntityAssignment

> TrusthubV1TrustProductEntityAssignment CreateTrustProductEntityAssignment(ctx, TrustProductSidoptional)

Create a new Assigned Item.

Create a new Assigned Item.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**TrustProductSid** | **string** | The unique string that we created to identify the TrustProduct resource.

### Other Parameters

Other parameters are passed through a pointer to a CreateTrustProductEntityAssignmentParams struct


Name | Type | Description
------------- | ------------- | -------------
**ObjectSid** | **string** | The SID of an object bag that holds information of the different items.

### Return type

[**TrusthubV1TrustProductEntityAssignment**](TrusthubV1TrustProductEntityAssignment.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTrustProductEntityAssignment

> DeleteTrustProductEntityAssignment(ctx, TrustProductSidSid)

Remove an Assignment Item Instance.

Remove an Assignment Item Instance.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**TrustProductSid** | **string** | The unique string that we created to identify the TrustProduct resource.
**Sid** | **string** | The unique string that we created to identify the Identity resource.

### Other Parameters

Other parameters are passed through a pointer to a DeleteTrustProductEntityAssignmentParams struct


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


## FetchTrustProductEntityAssignment

> TrusthubV1TrustProductEntityAssignment FetchTrustProductEntityAssignment(ctx, TrustProductSidSid)

Fetch specific Assigned Item Instance.

Fetch specific Assigned Item Instance.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**TrustProductSid** | **string** | The unique string that we created to identify the TrustProduct resource.
**Sid** | **string** | The unique string that we created to identify the Identity resource.

### Other Parameters

Other parameters are passed through a pointer to a FetchTrustProductEntityAssignmentParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**TrusthubV1TrustProductEntityAssignment**](TrusthubV1TrustProductEntityAssignment.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTrustProductEntityAssignment

> []TrusthubV1TrustProductEntityAssignment ListTrustProductEntityAssignment(ctx, TrustProductSidoptional)

Retrieve a list of all Assigned Items for an account.

Retrieve a list of all Assigned Items for an account.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**TrustProductSid** | **string** | The unique string that we created to identify the TrustProduct resource.

### Other Parameters

Other parameters are passed through a pointer to a ListTrustProductEntityAssignmentParams struct


Name | Type | Description
------------- | ------------- | -------------
**ObjectType** | **string** | A string to filter the results by (EndUserType or SupportingDocumentType) machine-name. This is useful when you want to retrieve the entity-assignment of a specific end-user or supporting document.
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]TrusthubV1TrustProductEntityAssignment**](TrusthubV1TrustProductEntityAssignment.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

