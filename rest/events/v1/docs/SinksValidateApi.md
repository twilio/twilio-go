# SinksValidateApi

All URIs are relative to *https://events.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSinkValidate**](SinksValidateApi.md#CreateSinkValidate) | **Post** /v1/Sinks/{Sid}/Validate | 



## CreateSinkValidate

> EventsV1SinkValidate CreateSinkValidate(ctx, Sidoptional)



Validate that a test event for a Sink was received.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | A 34 character string that uniquely identifies the Sink being validated.

### Other Parameters

Other parameters are passed through a pointer to a CreateSinkValidateParams struct


Name | Type | Description
------------- | ------------- | -------------
**TestId** | **string** | A 34 character string that uniquely identifies the test event for a Sink being validated.

### Return type

[**EventsV1SinkValidate**](EventsV1SinkValidate.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

