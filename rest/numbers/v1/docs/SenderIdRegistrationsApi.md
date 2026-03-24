# SenderIdRegistrationsApi

All URIs are relative to *https://numbers.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSenderIdRegistration**](SenderIdRegistrationsApi.md#CreateSenderIdRegistration) | **Post** /v1/SenderIdRegistrations | Create Sender ID Registration



## CreateSenderIdRegistration

> NumbersV1CreateEmbeddedRegistrationResponse CreateSenderIdRegistration(ctx, optional)

Create Sender ID Registration

Creates a new sender ID registration and initializes an embedded Persona inquiry session. Returns registration details and embedded session credentials for rendering the Compliance Embeddable UI.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateSenderIdRegistrationParams struct


Name | Type | Description
------------- | ------------- | -------------
**NumbersV1CreateEmbeddedRegistrationRequest** | [**NumbersV1CreateEmbeddedRegistrationRequest**](NumbersV1CreateEmbeddedRegistrationRequest.md) | 

### Return type

[**NumbersV1CreateEmbeddedRegistrationResponse**](NumbersV1CreateEmbeddedRegistrationResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

