# A2PCampaignRegistrationsEmbeddedSessionsApi

All URIs are relative to *https://trusthub.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateA2pCampaignRegistrationEmbeddedSession**](A2PCampaignRegistrationsEmbeddedSessionsApi.md#CreateA2pCampaignRegistrationEmbeddedSession) | **Post** /v1/A2PCampaignRegistrations/{Id}/EmbeddedSessions | Resume an A2P Campaign Registration embedded compliance session. Returns the existing active session if one is still valid, or creates a new session if the previous one has expired.



## CreateA2pCampaignRegistrationEmbeddedSession

> TrusthubV1A2pEmbeddedSession CreateA2pCampaignRegistrationEmbeddedSession(ctx, Id)

Resume an A2P Campaign Registration embedded compliance session. Returns the existing active session if one is still valid, or creates a new session if the previous one has expired.

Resume an A2P Campaign Registration embedded compliance session. Returns the existing active session if one is still valid, or creates a new session if the previous one has expired.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Id** | **string** | The unique identifier for the A2P Campaign Registration. This can be either an A2P Campaign Registration ID (formatted `tri1.us1.account.AC.../registration.BU...`) returned from the Initialize Campaign endpoint, or a Messaging Service SID (`MG...`) for campaigns created outside the embeddable.

### Other Parameters

Other parameters are passed through a pointer to a CreateA2pCampaignRegistrationEmbeddedSessionParams struct


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

