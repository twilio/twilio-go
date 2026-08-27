# A2PBrandRegistrationsApi

All URIs are relative to *https://trusthub.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateA2pBrandRegistration**](A2PBrandRegistrationsApi.md#CreateA2pBrandRegistration) | **Post** /v1/A2PBrandRegistrations | Create a new A2P Brand Registration and initialize an embedded compliance session.



## CreateA2pBrandRegistration

> TrusthubV1A2pEmbeddedSession CreateA2pBrandRegistration(ctx, optional)

Create a new A2P Brand Registration and initialize an embedded compliance session.

Create a new A2P Brand Registration and initialize an embedded compliance session.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateA2pBrandRegistrationParams struct


Name | Type | Description
------------- | ------------- | -------------
**TrusthubV1A2pBrandRegistrationRequest** | [**TrusthubV1A2pBrandRegistrationRequest**](TrusthubV1A2pBrandRegistrationRequest.md) | 

### Return type

[**TrusthubV1A2pEmbeddedSession**](TrusthubV1A2pEmbeddedSession.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

