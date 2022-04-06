# DeactivationsApi

All URIs are relative to *https://messaging.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchDeactivation**](DeactivationsApi.md#FetchDeactivation) | **Get** /v1/Deactivations | 



## FetchDeactivation

> MessagingV1Deactivation FetchDeactivation(ctx, optional)



Fetch a list of all United States numbers that have been deactivated on a specific date.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a FetchDeactivationParams struct


Name | Type | Description
------------- | ------------- | -------------
**Date** | **string** | The request will return a list of all United States Phone Numbers that were deactivated on the day specified by this parameter. This date should be specified in YYYY-MM-DD format.

### Return type

[**MessagingV1Deactivation**](MessagingV1Deactivation.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

