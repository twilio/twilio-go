# InteractionsChannelsApi

All URIs are relative to *https://flex-api.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchInteractionChannel**](InteractionsChannelsApi.md#FetchInteractionChannel) | **Get** /v1/Interactions/{InteractionSid}/Channels/{Sid} | 
[**ListInteractionChannel**](InteractionsChannelsApi.md#ListInteractionChannel) | **Get** /v1/Interactions/{InteractionSid}/Channels | 
[**UpdateInteractionChannel**](InteractionsChannelsApi.md#UpdateInteractionChannel) | **Post** /v1/Interactions/{InteractionSid}/Channels/{Sid} | 



## FetchInteractionChannel

> FlexV1InteractionChannel FetchInteractionChannel(ctx, InteractionSidSid)



Fetch a Channel for an Interaction.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**InteractionSid** | **string** | The Interaction Sid for this channel.
**Sid** | **string** | The Channel Sid for this Participant.

### Other Parameters

Other parameters are passed through a pointer to a FetchInteractionChannelParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**FlexV1InteractionChannel**](FlexV1InteractionChannel.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListInteractionChannel

> []FlexV1InteractionChannel ListInteractionChannel(ctx, InteractionSidoptional)



List all Channels for an Interaction.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**InteractionSid** | **string** | The Interaction Sid for this channel.

### Other Parameters

Other parameters are passed through a pointer to a ListInteractionChannelParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]FlexV1InteractionChannel**](FlexV1InteractionChannel.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateInteractionChannel

> FlexV1InteractionChannel UpdateInteractionChannel(ctx, InteractionSidSidoptional)



Update an existing Interaction.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**InteractionSid** | **string** | The Interaction Sid for this channel.
**Sid** | **string** | The unique string created by Twilio to identify an Interaction Channel resource.

### Other Parameters

Other parameters are passed through a pointer to a UpdateInteractionChannelParams struct


Name | Type | Description
------------- | ------------- | -------------
**Routing** | [**interface{}**](interface{}.md) | The Interaction Channel&#39;s routing parameters.  Optional, may contain &#x60;status&#x60; to set the agent&#39;s Reservation state to either &#x60;closed&#x60; or &#x60;wrapup&#x60;.  Default is &#x60;wrapup&#x60; if unspecified.
**Status** | **string** | The Interaction Channel&#39;s status. Can be: &#x60;closed&#x60; or &#x60;wrapup&#x60;.

### Return type

[**FlexV1InteractionChannel**](FlexV1InteractionChannel.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

