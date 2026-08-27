# A2PCampaignRegistrationsApi

All URIs are relative to *https://trusthub.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateA2pCampaignRegistration**](A2PCampaignRegistrationsApi.md#CreateA2pCampaignRegistration) | **Post** /v1/A2PCampaignRegistrations | Create a new A2P Campaign Registration and initialize an embedded compliance session.



## CreateA2pCampaignRegistration

> TrusthubV1A2pEmbeddedSession CreateA2pCampaignRegistration(ctx, optional)

Create a new A2P Campaign Registration and initialize an embedded compliance session.

Create a new A2P Campaign Registration and initialize an embedded compliance session.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateA2pCampaignRegistrationParams struct


Name | Type | Description
------------- | ------------- | -------------
**TrusthubV1A2pCampaignRegistrationRequest** | [**TrusthubV1A2pCampaignRegistrationRequest**](TrusthubV1A2pCampaignRegistrationRequest.md) | 

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

