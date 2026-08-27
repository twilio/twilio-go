# CustomerProfilesProvisionalCopyApi

All URIs are relative to *https://trusthub.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCustomerProfileProvisionalCopy**](CustomerProfilesProvisionalCopyApi.md#CreateCustomerProfileProvisionalCopy) | **Post** /v1/CustomerProfiles/{CustomerProfileSid}/ProvisionalCopy | Create the provisional copy for a given customer profile with status: TWILIO_APPROVED, this is useful for making updates to an existing immutable customer profile instance.
[**FetchCustomerProfileProvisionalCopy**](CustomerProfilesProvisionalCopyApi.md#FetchCustomerProfileProvisionalCopy) | **Get** /v1/CustomerProfiles/{CustomerProfileSid}/ProvisionalCopy | Fetch the provisional copy of a given customer profile.



## CreateCustomerProfileProvisionalCopy

> TrusthubV1CustomerProfileProvisionalCopy CreateCustomerProfileProvisionalCopy(ctx, CustomerProfileSid)

Create the provisional copy for a given customer profile with status: TWILIO_APPROVED, this is useful for making updates to an existing immutable customer profile instance.

Create the provisional copy for a given customer profile with status: TWILIO_APPROVED, this is useful for making updates to an existing immutable customer profile instance.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**CustomerProfileSid** | **string** | The unique SID identifier of the Customer Profile.

### Other Parameters

Other parameters are passed through a pointer to a CreateCustomerProfileProvisionalCopyParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**TrusthubV1CustomerProfileProvisionalCopy**](TrusthubV1CustomerProfileProvisionalCopy.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchCustomerProfileProvisionalCopy

> TrusthubV1CustomerProfileProvisionalCopy FetchCustomerProfileProvisionalCopy(ctx, CustomerProfileSid)

Fetch the provisional copy of a given customer profile.

Fetch the provisional copy of a given customer profile.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**CustomerProfileSid** | **string** | The unique SID identifier of the Customer Profile.

### Other Parameters

Other parameters are passed through a pointer to a FetchCustomerProfileProvisionalCopyParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**TrusthubV1CustomerProfileProvisionalCopy**](TrusthubV1CustomerProfileProvisionalCopy.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

