# SinksTestApi

All URIs are relative to *https://events.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSinkTest**](SinksTestApi.md#CreateSinkTest) | **Post** /v1/Sinks/{Sid}/Test | 



## CreateSinkTest

> EventsV1SinkTest CreateSinkTest(ctx, Sid)



Create a new Sink Test Event for the given Sink.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | A 34 character string that uniquely identifies the Sink to be Tested.

### Other Parameters

Other parameters are passed through a pointer to a CreateSinkTestParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**EventsV1SinkTest**](EventsV1SinkTest.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

