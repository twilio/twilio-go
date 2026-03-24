# SenderIdRegistrationsEmbeddedSessionsApi

All URIs are relative to *https://numbers.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSenderIdRegistrationEmbeddedSession**](SenderIdRegistrationsEmbeddedSessionsApi.md#CreateSenderIdRegistrationEmbeddedSession) | **Post** /v1/SenderIdRegistrations/{BundleSid}/EmbeddedSessions | Create Embedded Session



## CreateSenderIdRegistrationEmbeddedSession

> NumbersV1CreateEmbeddedSessionResponse CreateSenderIdRegistrationEmbeddedSession(ctx, BundleSidoptional)

Create Embedded Session

Creates a new embedded Persona inquiry session for an existing registration in DRAFT or TWILIO_REJECTED status. Use this to resume an incomplete registration or resubmit a rejected one.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**BundleSid** | **string** | The unique identifier of the registration (BU-prefixed).

### Other Parameters

Other parameters are passed through a pointer to a CreateSenderIdRegistrationEmbeddedSessionParams struct


Name | Type | Description
------------- | ------------- | -------------
**NumbersV1CreateEmbeddedSessionRequest** | [**NumbersV1CreateEmbeddedSessionRequest**](NumbersV1CreateEmbeddedSessionRequest.md) | 

### Return type

[**NumbersV1CreateEmbeddedSessionResponse**](NumbersV1CreateEmbeddedSessionResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

