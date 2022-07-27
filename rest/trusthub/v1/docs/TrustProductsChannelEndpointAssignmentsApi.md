# TrustProductsChannelEndpointAssignmentsApi

All URIs are relative to *https://trusthub.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTrustProductChannelEndpointAssignment**](TrustProductsChannelEndpointAssignmentsApi.md#CreateTrustProductChannelEndpointAssignment) | **Post** /v1/TrustProducts/{TrustProductSid}/ChannelEndpointAssignments | 
[**DeleteTrustProductChannelEndpointAssignment**](TrustProductsChannelEndpointAssignmentsApi.md#DeleteTrustProductChannelEndpointAssignment) | **Delete** /v1/TrustProducts/{TrustProductSid}/ChannelEndpointAssignments/{Sid} | 
[**FetchTrustProductChannelEndpointAssignment**](TrustProductsChannelEndpointAssignmentsApi.md#FetchTrustProductChannelEndpointAssignment) | **Get** /v1/TrustProducts/{TrustProductSid}/ChannelEndpointAssignments/{Sid} | 
[**ListTrustProductChannelEndpointAssignment**](TrustProductsChannelEndpointAssignmentsApi.md#ListTrustProductChannelEndpointAssignment) | **Get** /v1/TrustProducts/{TrustProductSid}/ChannelEndpointAssignments | 



## CreateTrustProductChannelEndpointAssignment

> TrusthubV1TrustProductChannelEndpointAssignment CreateTrustProductChannelEndpointAssignment(ctx, TrustProductSidoptional)



Create a new Assigned Item.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**TrustProductSid** | **string** | The unique string that we created to identify the CustomerProfile resource.

### Other Parameters

Other parameters are passed through a pointer to a CreateTrustProductChannelEndpointAssignmentParams struct


Name | Type | Description
------------- | ------------- | -------------
**ChannelEndpointType** | **string** | The type of channel endpoint. eg: phone-number
**ChannelEndpointSid** | **string** | The SID of an channel endpoint

### Return type

[**TrusthubV1TrustProductChannelEndpointAssignment**](TrusthubV1TrustProductChannelEndpointAssignment.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTrustProductChannelEndpointAssignment

> DeleteTrustProductChannelEndpointAssignment(ctx, TrustProductSidSid)



Remove an Assignment Item Instance.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**TrustProductSid** | **string** | The unique string that we created to identify the CustomerProfile resource.
**Sid** | **string** | The unique string that we created to identify the resource.

### Other Parameters

Other parameters are passed through a pointer to a DeleteTrustProductChannelEndpointAssignmentParams struct


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


## FetchTrustProductChannelEndpointAssignment

> TrusthubV1TrustProductChannelEndpointAssignment FetchTrustProductChannelEndpointAssignment(ctx, TrustProductSidSid)



Fetch specific Assigned Item Instance.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**TrustProductSid** | **string** | The unique string that we created to identify the CustomerProfile resource.
**Sid** | **string** | The unique string that we created to identify the resource.

### Other Parameters

Other parameters are passed through a pointer to a FetchTrustProductChannelEndpointAssignmentParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**TrusthubV1TrustProductChannelEndpointAssignment**](TrusthubV1TrustProductChannelEndpointAssignment.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTrustProductChannelEndpointAssignment

> []TrusthubV1TrustProductChannelEndpointAssignment ListTrustProductChannelEndpointAssignment(ctx, TrustProductSidoptional)



Retrieve a list of all Assigned Items for an account.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**TrustProductSid** | **string** | The unique string that we created to identify the CustomerProfile resource.

### Other Parameters

Other parameters are passed through a pointer to a ListTrustProductChannelEndpointAssignmentParams struct


Name | Type | Description
------------- | ------------- | -------------
**ChannelEndpointSid** | **string** | The SID of an channel endpoint
**ChannelEndpointSids** | **string** | comma separated list of channel endpoint sids
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]TrusthubV1TrustProductChannelEndpointAssignment**](TrusthubV1TrustProductChannelEndpointAssignment.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

