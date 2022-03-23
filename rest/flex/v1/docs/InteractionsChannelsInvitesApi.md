# InteractionsChannelsInvitesApi

All URIs are relative to *https://flex-api.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateInteractionChannelInvite**](InteractionsChannelsInvitesApi.md#CreateInteractionChannelInvite) | **Post** /v1/Interactions/{InteractionSid}/Channels/{ChannelSid}/Invites | 
[**ListInteractionChannelInvite**](InteractionsChannelsInvitesApi.md#ListInteractionChannelInvite) | **Get** /v1/Interactions/{InteractionSid}/Channels/{ChannelSid}/Invites | 



## CreateInteractionChannelInvite

> FlexV1InteractionChannelInvite CreateInteractionChannelInvite(ctx, InteractionSidChannelSidoptional)



### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**InteractionSid** | **string** | 
**ChannelSid** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a CreateInteractionChannelInviteParams struct


Name | Type | Description
------------- | ------------- | -------------
**Routing** | [**map[string]interface{}**](map[string]interface{}.md) | 

### Return type

[**FlexV1InteractionChannelInvite**](FlexV1InteractionChannelInvite.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListInteractionChannelInvite

> []FlexV1InteractionChannelInvite ListInteractionChannelInvite(ctx, InteractionSidChannelSidoptional)



### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**InteractionSid** | **string** | 
**ChannelSid** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a ListInteractionChannelInviteParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]FlexV1InteractionChannelInvite**](FlexV1InteractionChannelInvite.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

