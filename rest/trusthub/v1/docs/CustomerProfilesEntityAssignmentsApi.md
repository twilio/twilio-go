# CustomerProfilesEntityAssignmentsApi

All URIs are relative to *https://trusthub.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCustomerProfileEntityAssignment**](CustomerProfilesEntityAssignmentsApi.md#CreateCustomerProfileEntityAssignment) | **Post** /v1/CustomerProfiles/{CustomerProfileSid}/EntityAssignments | 
[**DeleteCustomerProfileEntityAssignment**](CustomerProfilesEntityAssignmentsApi.md#DeleteCustomerProfileEntityAssignment) | **Delete** /v1/CustomerProfiles/{CustomerProfileSid}/EntityAssignments/{Sid} | 
[**FetchCustomerProfileEntityAssignment**](CustomerProfilesEntityAssignmentsApi.md#FetchCustomerProfileEntityAssignment) | **Get** /v1/CustomerProfiles/{CustomerProfileSid}/EntityAssignments/{Sid} | 
[**ListCustomerProfileEntityAssignment**](CustomerProfilesEntityAssignmentsApi.md#ListCustomerProfileEntityAssignment) | **Get** /v1/CustomerProfiles/{CustomerProfileSid}/EntityAssignments | 



## CreateCustomerProfileEntityAssignment

> TrusthubV1CustomerProfileEntityAssignment CreateCustomerProfileEntityAssignment(ctx, CustomerProfileSidoptional)



Create a new Assigned Item.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**CustomerProfileSid** | **string** | The unique string that we created to identify the CustomerProfile resource.

### Other Parameters

Other parameters are passed through a pointer to a CreateCustomerProfileEntityAssignmentParams struct


Name | Type | Description
------------- | ------------- | -------------
**ObjectSid** | **string** | The SID of an object bag that holds information of the different items.

### Return type

[**TrusthubV1CustomerProfileEntityAssignment**](TrusthubV1CustomerProfileEntityAssignment.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCustomerProfileEntityAssignment

> DeleteCustomerProfileEntityAssignment(ctx, CustomerProfileSidSid)



Remove an Assignment Item Instance.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**CustomerProfileSid** | **string** | The unique string that we created to identify the CustomerProfile resource.
**Sid** | **string** | The unique string that we created to identify the Identity resource.

### Other Parameters

Other parameters are passed through a pointer to a DeleteCustomerProfileEntityAssignmentParams struct


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


## FetchCustomerProfileEntityAssignment

> TrusthubV1CustomerProfileEntityAssignment FetchCustomerProfileEntityAssignment(ctx, CustomerProfileSidSid)



Fetch specific Assigned Item Instance.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**CustomerProfileSid** | **string** | The unique string that we created to identify the CustomerProfile resource.
**Sid** | **string** | The unique string that we created to identify the Identity resource.

### Other Parameters

Other parameters are passed through a pointer to a FetchCustomerProfileEntityAssignmentParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**TrusthubV1CustomerProfileEntityAssignment**](TrusthubV1CustomerProfileEntityAssignment.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCustomerProfileEntityAssignment

> []TrusthubV1CustomerProfileEntityAssignment ListCustomerProfileEntityAssignment(ctx, CustomerProfileSidoptional)



Retrieve a list of all Assigned Items for an account.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**CustomerProfileSid** | **string** | The unique string that we created to identify the CustomerProfile resource.

### Other Parameters

Other parameters are passed through a pointer to a ListCustomerProfileEntityAssignmentParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]TrusthubV1CustomerProfileEntityAssignment**](TrusthubV1CustomerProfileEntityAssignment.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

