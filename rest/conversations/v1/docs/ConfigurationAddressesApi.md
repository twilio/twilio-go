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
**Address** | **string** | The unique address to be configured. The address can be a whatsapp address or phone number
**AutoCreationConversationServiceSid** | **string** | Conversation Service for the auto-created conversation. If not set, the conversation is created in the default service.
**AutoCreationEnabled** | **bool** | Enable/Disable auto-creating conversations for messages to this address
**AutoCreationStudioFlowSid** | **string** | For type &#x60;studio&#x60;, the studio flow SID where the webhook should be sent to.
**AutoCreationStudioRetryCount** | **int** | For type &#x60;studio&#x60;, number of times to retry the webhook request
**AutoCreationType** | **string** | Type of Auto Creation. Value can be one of &#x60;webhook&#x60;, &#x60;studio&#x60; or &#x60;default&#x60;.
**AutoCreationWebhookFilters** | **[]string** | The list of events, firing webhook event for this Conversation. Values can be any of the following: &#x60;onMessageAdded&#x60;, &#x60;onMessageUpdated&#x60;, &#x60;onMessageRemoved&#x60;, &#x60;onConversationUpdated&#x60;, &#x60;onConversationStateUpdated&#x60;, &#x60;onConversationRemoved&#x60;, &#x60;onParticipantAdded&#x60;, &#x60;onParticipantUpdated&#x60;, &#x60;onParticipantRemoved&#x60;, &#x60;onDeliveryUpdated&#x60;
**AutoCreationWebhookMethod** | **string** | For type &#x60;webhook&#x60;, the HTTP method to be used when sending a webhook request.
**AutoCreationWebhookUrl** | **string** | For type &#x60;webhook&#x60;, the url for the webhook request.
**FriendlyName** | **string** | The human-readable name of this configuration, limited to 256 characters. Optional.
**Type** | **string** | Type of Address. Value can be &#x60;whatsapp&#x60; or &#x60;sms&#x60;.

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
**Sid** | **string** | The SID of the Address Configuration resource. This value can be either the &#x60;sid&#x60; or the &#x60;address&#x60; of the configuration

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
**Sid** | **string** | The SID of the Address Configuration resource. This value can be either the &#x60;sid&#x60; or the &#x60;address&#x60; of the configuration

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
**Sid** | **string** | The SID of the Address Configuration resource. This value can be either the &#x60;sid&#x60; or the &#x60;address&#x60; of the configuration

### Other Parameters

Other parameters are passed through a pointer to a UpdateConfigurationAddressParams struct


Name | Type | Description
------------- | ------------- | -------------
**AutoCreationConversationServiceSid** | **string** | Conversation Service for the auto-created conversation. If not set, the conversation is created in the default service.
**AutoCreationEnabled** | **bool** | Enable/Disable auto-creating conversations for messages to this address
**AutoCreationStudioFlowSid** | **string** | For type &#x60;studio&#x60;, the studio flow SID where the webhook should be sent to.
**AutoCreationStudioRetryCount** | **int** | For type &#x60;studio&#x60;, number of times to retry the webhook request
**AutoCreationType** | **string** | Type of Auto Creation. Value can be one of &#x60;webhook&#x60;, &#x60;studio&#x60; or &#x60;default&#x60;.
**AutoCreationWebhookFilters** | **[]string** | The list of events, firing webhook event for this Conversation. Values can be any of the following: &#x60;onMessageAdded&#x60;, &#x60;onMessageUpdated&#x60;, &#x60;onMessageRemoved&#x60;, &#x60;onConversationUpdated&#x60;, &#x60;onConversationStateUpdated&#x60;, &#x60;onConversationRemoved&#x60;, &#x60;onParticipantAdded&#x60;, &#x60;onParticipantUpdated&#x60;, &#x60;onParticipantRemoved&#x60;, &#x60;onDeliveryUpdated&#x60;
**AutoCreationWebhookMethod** | **string** | For type &#x60;webhook&#x60;, the HTTP method to be used when sending a webhook request.
**AutoCreationWebhookUrl** | **string** | For type &#x60;webhook&#x60;, the url for the webhook request.
**FriendlyName** | **string** | The human-readable name of this configuration, limited to 256 characters. Optional.

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

