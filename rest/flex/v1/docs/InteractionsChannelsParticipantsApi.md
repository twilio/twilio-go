# InteractionsChannelsParticipantsApi

All URIs are relative to *https://flex-api.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateInteractionChannelParticipant**](InteractionsChannelsParticipantsApi.md#CreateInteractionChannelParticipant) | **Post** /v1/Interactions/{InteractionSid}/Channels/{ChannelSid}/Participants | 
[**ListInteractionChannelParticipant**](InteractionsChannelsParticipantsApi.md#ListInteractionChannelParticipant) | **Get** /v1/Interactions/{InteractionSid}/Channels/{ChannelSid}/Participants | 
[**UpdateInteractionChannelParticipant**](InteractionsChannelsParticipantsApi.md#UpdateInteractionChannelParticipant) | **Post** /v1/Interactions/{InteractionSid}/Channels/{ChannelSid}/Participants/{Sid} | 



## CreateInteractionChannelParticipant

> FlexV1InteractionChannelParticipant CreateInteractionChannelParticipant(ctx, InteractionSidChannelSidoptional)



Add a Participant to a Channel.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**InteractionSid** | **string** | The Interaction Sid for the new Channel Participant.
**ChannelSid** | **string** | The Channel Sid for the new Channel Participant.

### Other Parameters

Other parameters are passed through a pointer to a CreateInteractionChannelParticipantParams struct


Name | Type | Description
------------- | ------------- | -------------
**MediaProperties** | [**map[string]interface{}**](map[string]interface{}.md) | JSON representing the Media Properties for the new Participant.
**Type** | **string** | Participant type.  Can be: &#x60;agent&#x60;, &#x60;customer&#x60;, &#x60;supervisor&#x60;, &#x60;external&#x60; or &#x60;unknown&#x60;.

### Return type

[**FlexV1InteractionChannelParticipant**](FlexV1InteractionChannelParticipant.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListInteractionChannelParticipant

> []FlexV1InteractionChannelParticipant ListInteractionChannelParticipant(ctx, InteractionSidChannelSidoptional)



List all Participants for a Channel.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**InteractionSid** | **string** | The Interaction Sid for this channel.
**ChannelSid** | **string** | The Channel Sid for this Participant.

### Other Parameters

Other parameters are passed through a pointer to a ListInteractionChannelParticipantParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]FlexV1InteractionChannelParticipant**](FlexV1InteractionChannelParticipant.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateInteractionChannelParticipant

> FlexV1InteractionChannelParticipant UpdateInteractionChannelParticipant(ctx, InteractionSidChannelSidSidoptional)



Update an existing Channel Participant.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**InteractionSid** | **string** | The Interaction Sid for this channel.
**ChannelSid** | **string** | The Channel Sid for this Participant.
**Sid** | **string** | The unique string created by Twilio to identify an Interaction Channel resource.

### Other Parameters

Other parameters are passed through a pointer to a UpdateInteractionChannelParticipantParams struct


Name | Type | Description
------------- | ------------- | -------------
**Status** | **string** | The Participant&#39;s status. Can be: &#x60;closed&#x60; or &#x60;wrapup&#x60;.  Participant must be an agent.

### Return type

[**FlexV1InteractionChannelParticipant**](FlexV1InteractionChannelParticipant.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

