# IdentityApi

All URIs are relative to *https://context.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchIdentity**](IdentityApi.md#FetchIdentity) | **Get** /v1/Identity | 



## FetchIdentity

> ContextV1Identity FetchIdentity(ctx, optional)



Retrieves a list of identities.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a FetchIdentityParams struct


Name | Type | Description
------------- | ------------- | -------------
**ChannelFieldName** | **string** | Name of the channel used to fetch identities
**ChannelFieldValue** | **string** | Value of the channel for which identities are fetched

### Return type

[**ContextV1Identity**](ContextV1Identity.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

