# InteractionsChannelsTransfersApi

All URIs are relative to *https://flex-api.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateInteractionTransfer**](InteractionsChannelsTransfersApi.md#CreateInteractionTransfer) | **Post** /v1/Interactions/{InteractionSid}/Channels/{ChannelSid}/Transfers | Create a new Transfer.
[**FetchInteractionTransfer**](InteractionsChannelsTransfersApi.md#FetchInteractionTransfer) | **Get** /v1/Interactions/{InteractionSid}/Channels/{ChannelSid}/Transfers/{Sid} | Fetch a specific Transfer by SID.
[**UpdateInteractionTransfer**](InteractionsChannelsTransfersApi.md#UpdateInteractionTransfer) | **Post** /v1/Interactions/{InteractionSid}/Channels/{ChannelSid}/Transfers/{Sid} | Update an existing Transfer.



## CreateInteractionTransfer

> FlexV1InteractionTransfer CreateInteractionTransfer(ctx, InteractionSidChannelSidoptional)

Create a new Transfer.

Create a new Transfer.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**InteractionSid** | **string** | The Interaction Sid for the Interaction
**ChannelSid** | **string** | The Channel Sid for the Channel.

### Other Parameters

Other parameters are passed through a pointer to a CreateInteractionTransferParams struct


Name | Type | Description
------------- | ------------- | -------------
**Body** | **map[string]interface{}** | 

### Return type

[**FlexV1InteractionTransfer**](FlexV1InteractionTransfer.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchInteractionTransfer

> FlexV1InteractionTransfer FetchInteractionTransfer(ctx, InteractionSidChannelSidSid)

Fetch a specific Transfer by SID.

Fetch a specific Transfer by SID.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**InteractionSid** | **string** | The Interaction Sid for this channel.
**ChannelSid** | **string** | The Channel Sid for this Transfer.
**Sid** | **string** | The unique string created by Twilio to identify a Transfer resource.

### Other Parameters

Other parameters are passed through a pointer to a FetchInteractionTransferParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**FlexV1InteractionTransfer**](FlexV1InteractionTransfer.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateInteractionTransfer

> FlexV1InteractionTransfer UpdateInteractionTransfer(ctx, InteractionSidChannelSidSidoptional)

Update an existing Transfer.

Update an existing Transfer.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**InteractionSid** | **string** | The Interaction Sid for this channel.
**ChannelSid** | **string** | The Channel Sid for this Transfer.
**Sid** | **string** | The unique string created by Twilio to identify a Transfer resource.

### Other Parameters

Other parameters are passed through a pointer to a UpdateInteractionTransferParams struct


Name | Type | Description
------------- | ------------- | -------------
**Body** | **map[string]interface{}** | 

### Return type

[**FlexV1InteractionTransfer**](FlexV1InteractionTransfer.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

