# InteractionsApi

All URIs are relative to *https://flex-api.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateInteraction**](InteractionsApi.md#CreateInteraction) | **Post** /v1/Interactions | Create a new Interaction.
[**FetchInteraction**](InteractionsApi.md#FetchInteraction) | **Get** /v1/Interactions/{Sid} | 
[**UpdateInteraction**](InteractionsApi.md#UpdateInteraction) | **Post** /v1/Interactions/{Sid} | Updates an interaction.



## CreateInteraction

> FlexV1Interaction CreateInteraction(ctx, optional)

Create a new Interaction.

Create a new Interaction.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateInteractionParams struct


Name | Type | Description
------------- | ------------- | -------------
**Channel** | [**interface{}**](interface{}.md) | The Interaction's channel.
**Routing** | [**interface{}**](interface{}.md) | The Interaction's routing logic.
**InteractionContextSid** | **string** | The Interaction context sid is used for adding a context lookup sid
**WebhookTtid** | **string** | The unique identifier for Interaction level webhook

### Return type

[**FlexV1Interaction**](FlexV1Interaction.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchInteraction

> FlexV1Interaction FetchInteraction(ctx, Sid)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The SID of the Interaction resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchInteractionParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**FlexV1Interaction**](FlexV1Interaction.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateInteraction

> FlexV1Interaction UpdateInteraction(ctx, Sidoptional)

Updates an interaction.

Updates an interaction.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The SID of the Interaction resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a UpdateInteractionParams struct


Name | Type | Description
------------- | ------------- | -------------
**WebhookTtid** | **string** | The unique identifier for Interaction level webhook

### Return type

[**FlexV1Interaction**](FlexV1Interaction.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

