# BroadcastsApi

All URIs are relative to *https://preview.messaging.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBroadcast**](BroadcastsApi.md#CreateBroadcast) | **Post** /v1/Broadcasts | 



## CreateBroadcast

> MessagingV1Broadcast CreateBroadcast(ctx, optional)



Create a new Broadcast

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateBroadcastParams struct


Name | Type | Description
------------- | ------------- | -------------
**XTwilioRequestKey** | **string** | Idempotency key provided by the client

### Return type

[**MessagingV1Broadcast**](MessagingV1Broadcast.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

