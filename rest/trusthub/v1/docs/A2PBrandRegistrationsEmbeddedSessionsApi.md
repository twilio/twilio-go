# A2PBrandRegistrationsEmbeddedSessionsApi

All URIs are relative to *https://trusthub.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateA2pBrandRegistrationEmbeddedSession**](A2PBrandRegistrationsEmbeddedSessionsApi.md#CreateA2pBrandRegistrationEmbeddedSession) | **Post** /v1/A2PBrandRegistrations/{Id}/EmbeddedSessions | Resume an A2P Brand Registration embedded compliance session. Returns the existing active session if one is still valid, or creates a new session if the previous one has expired.



## CreateA2pBrandRegistrationEmbeddedSession

> TrusthubV1A2pEmbeddedSession CreateA2pBrandRegistrationEmbeddedSession(ctx, Id)

Resume an A2P Brand Registration embedded compliance session. Returns the existing active session if one is still valid, or creates a new session if the previous one has expired.

Resume an A2P Brand Registration embedded compliance session. Returns the existing active session if one is still valid, or creates a new session if the previous one has expired.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Id** | **string** | The unique identifier for the A2P Brand Registration.

### Other Parameters

Other parameters are passed through a pointer to a CreateA2pBrandRegistrationEmbeddedSessionParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**TrusthubV1A2pEmbeddedSession**](TrusthubV1A2pEmbeddedSession.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

