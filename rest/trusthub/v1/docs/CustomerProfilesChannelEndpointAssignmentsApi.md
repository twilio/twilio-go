# CustomerProfilesChannelEndpointAssignmentsApi

All URIs are relative to *https://trusthub.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCustomerProfileChannelEndpointAssignment**](CustomerProfilesChannelEndpointAssignmentsApi.md#CreateCustomerProfileChannelEndpointAssignment) | **Post** /v1/CustomerProfiles/{CustomerProfileSid}/ChannelEndpointAssignments | 
[**DeleteCustomerProfileChannelEndpointAssignment**](CustomerProfilesChannelEndpointAssignmentsApi.md#DeleteCustomerProfileChannelEndpointAssignment) | **Delete** /v1/CustomerProfiles/{CustomerProfileSid}/ChannelEndpointAssignments/{Sid} | 
[**FetchCustomerProfileChannelEndpointAssignment**](CustomerProfilesChannelEndpointAssignmentsApi.md#FetchCustomerProfileChannelEndpointAssignment) | **Get** /v1/CustomerProfiles/{CustomerProfileSid}/ChannelEndpointAssignments/{Sid} | 
[**ListCustomerProfileChannelEndpointAssignment**](CustomerProfilesChannelEndpointAssignmentsApi.md#ListCustomerProfileChannelEndpointAssignment) | **Get** /v1/CustomerProfiles/{CustomerProfileSid}/ChannelEndpointAssignments | 



## CreateCustomerProfileChannelEndpointAssignment

> TrusthubV1CustomerProfileChannelEndpointAssignment CreateCustomerProfileChannelEndpointAssignment(ctx, CustomerProfileSidoptional)



Create a new Assigned Item.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**CustomerProfileSid** | **string** | The unique string that we created to identify the CustomerProfile resource.

### Other Parameters

Other parameters are passed through a pointer to a CreateCustomerProfileChannelEndpointAssignmentParams struct


Name | Type | Description
------------- | ------------- | -------------
**ChannelEndpointType** | **string** | The type of channel endpoint. eg: phone-number
**ChannelEndpointSid** | **string** | The SID of an channel endpoint

### Return type

[**TrusthubV1CustomerProfileChannelEndpointAssignment**](TrusthubV1CustomerProfileChannelEndpointAssignment.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCustomerProfileChannelEndpointAssignment

> DeleteCustomerProfileChannelEndpointAssignment(ctx, CustomerProfileSidSid)



Remove an Assignment Item Instance.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**CustomerProfileSid** | **string** | The unique string that we created to identify the CustomerProfile resource.
**Sid** | **string** | The unique string that we created to identify the resource.

### Other Parameters

Other parameters are passed through a pointer to a DeleteCustomerProfileChannelEndpointAssignmentParams struct


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


## FetchCustomerProfileChannelEndpointAssignment

> TrusthubV1CustomerProfileChannelEndpointAssignment FetchCustomerProfileChannelEndpointAssignment(ctx, CustomerProfileSidSid)



Fetch specific Assigned Item Instance.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**CustomerProfileSid** | **string** | The unique string that we created to identify the CustomerProfile resource.
**Sid** | **string** | The unique string that we created to identify the resource.

### Other Parameters

Other parameters are passed through a pointer to a FetchCustomerProfileChannelEndpointAssignmentParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**TrusthubV1CustomerProfileChannelEndpointAssignment**](TrusthubV1CustomerProfileChannelEndpointAssignment.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCustomerProfileChannelEndpointAssignment

> []TrusthubV1CustomerProfileChannelEndpointAssignment ListCustomerProfileChannelEndpointAssignment(ctx, CustomerProfileSidoptional)



Retrieve a list of all Assigned Items for an account.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**CustomerProfileSid** | **string** | The unique string that we created to identify the CustomerProfile resource.

### Other Parameters

Other parameters are passed through a pointer to a ListCustomerProfileChannelEndpointAssignmentParams struct


Name | Type | Description
------------- | ------------- | -------------
**ChannelEndpointSid** | **string** | The SID of an channel endpoint
**ChannelEndpointSids** | **string** | comma separated list of channel endpoint sids
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]TrusthubV1CustomerProfileChannelEndpointAssignment**](TrusthubV1CustomerProfileChannelEndpointAssignment.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

