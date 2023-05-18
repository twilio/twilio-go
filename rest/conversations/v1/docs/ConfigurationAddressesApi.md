# ConfigurationAddressesApi

All URIs are relative to *https://conversations.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateConfigurationAddress**](ConfigurationAddressesApi.md#CreateConfigurationAddress) | **Post** /v1/Configuration/Addresses | 
[**DeleteConfigurationAddress**](ConfigurationAddressesApi.md#DeleteConfigurationAddress) | **Delete** /v1/Configuration/Addresses/{Sid} | 
[**FetchConfigurationAddress**](ConfigurationAddressesApi.md#FetchConfigurationAddress) | **Get** /v1/Configuration/Addresses/{Sid} | 
[**ListConfigurationAddress**](ConfigurationAddressesApi.md#ListConfigurationAddress) | **Get** /v1/Configuration/Addresses | 
[**UpdateConfigurationAddress**](ConfigurationAddressesApi.md#UpdateConfigurationAddress) | **Post** /v1/Configuration/Addresses/{Sid} | 



## CreateConfigurationAddress

> ConversationsV1ConfigurationAddress CreateConfigurationAddress(ctx, optional)



Create a new address configuration

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateConfigurationAddressParams struct


Name | Type | Description
------------- | ------------- | -------------
**Type** | **string** | 
**Address** | **string** | The unique address to be configured. The address can be a whatsapp address or phone number
**FriendlyName** | **string** | The human-readable name of this configuration, limited to 256 characters. Optional.
**AutoCreationEnabled** | **bool** | Enable/Disable auto-creating conversations for messages to this address
**AutoCreationType** | **string** | 
**AutoCreationConversationServiceSid** | **string** | Conversation Service for the auto-created conversation. If not set, the conversation is created in the default service.
**AutoCreationWebhookUrl** | **string** | For type `webhook`, the url for the webhook request.
**AutoCreationWebhookMethod** | **string** | 
**AutoCreationWebhookFilters** | **[]string** | The list of events, firing webhook event for this Conversation. Values can be any of the following: `onMessageAdded`, `onMessageUpdated`, `onMessageRemoved`, `onConversationUpdated`, `onConversationStateUpdated`, `onConversationRemoved`, `onParticipantAdded`, `onParticipantUpdated`, `onParticipantRemoved`, `onDeliveryUpdated`
**AutoCreationStudioFlowSid** | **string** | For type `studio`, the studio flow SID where the webhook should be sent to.
**AutoCreationStudioRetryCount** | **int** | For type `studio`, number of times to retry the webhook request
**AddressCountry** | **string** | An ISO 3166-1 alpha-2n country code which the address belongs to. This is currently only applicable to short code addresses.

### Return type

[**ConversationsV1ConfigurationAddress**](ConversationsV1ConfigurationAddress.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteConfigurationAddress

> DeleteConfigurationAddress(ctx, Sid)



Remove an existing address configuration

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The SID of the Address Configuration resource. This value can be either the `sid` or the `address` of the configuration

### Other Parameters

Other parameters are passed through a pointer to a DeleteConfigurationAddressParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

 (empty response body)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchConfigurationAddress

> ConversationsV1ConfigurationAddress FetchConfigurationAddress(ctx, Sid)



Fetch an address configuration 

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The SID of the Address Configuration resource. This value can be either the `sid` or the `address` of the configuration

### Other Parameters

Other parameters are passed through a pointer to a FetchConfigurationAddressParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**ConversationsV1ConfigurationAddress**](ConversationsV1ConfigurationAddress.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListConfigurationAddress

> []ConversationsV1ConfigurationAddress ListConfigurationAddress(ctx, optional)



Retrieve a list of address configurations for an account

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListConfigurationAddressParams struct


Name | Type | Description
------------- | ------------- | -------------
**Type** | **string** | Filter the address configurations by its type. This value can be one of: `whatsapp`, `sms`.
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]ConversationsV1ConfigurationAddress**](ConversationsV1ConfigurationAddress.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateConfigurationAddress

> ConversationsV1ConfigurationAddress UpdateConfigurationAddress(ctx, Sidoptional)



Update an existing address configuration

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The SID of the Address Configuration resource. This value can be either the `sid` or the `address` of the configuration

### Other Parameters

Other parameters are passed through a pointer to a UpdateConfigurationAddressParams struct


Name | Type | Description
------------- | ------------- | -------------
**FriendlyName** | **string** | The human-readable name of this configuration, limited to 256 characters. Optional.
**AutoCreationEnabled** | **bool** | Enable/Disable auto-creating conversations for messages to this address
**AutoCreationType** | **string** | 
**AutoCreationConversationServiceSid** | **string** | Conversation Service for the auto-created conversation. If not set, the conversation is created in the default service.
**AutoCreationWebhookUrl** | **string** | For type `webhook`, the url for the webhook request.
**AutoCreationWebhookMethod** | **string** | 
**AutoCreationWebhookFilters** | **[]string** | The list of events, firing webhook event for this Conversation. Values can be any of the following: `onMessageAdded`, `onMessageUpdated`, `onMessageRemoved`, `onConversationUpdated`, `onConversationStateUpdated`, `onConversationRemoved`, `onParticipantAdded`, `onParticipantUpdated`, `onParticipantRemoved`, `onDeliveryUpdated`
**AutoCreationStudioFlowSid** | **string** | For type `studio`, the studio flow SID where the webhook should be sent to.
**AutoCreationStudioRetryCount** | **int** | For type `studio`, number of times to retry the webhook request

### Return type

[**ConversationsV1ConfigurationAddress**](ConversationsV1ConfigurationAddress.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

