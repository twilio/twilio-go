# DefaultApi

All URIs are relative to *https://flex-api.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateChannel**](DefaultApi.md#CreateChannel) | **Post** /v1/Channels | 
[**CreateFlexFlow**](DefaultApi.md#CreateFlexFlow) | **Post** /v1/FlexFlows | 
[**CreateWebChannel**](DefaultApi.md#CreateWebChannel) | **Post** /v1/WebChannels | 
[**DeleteChannel**](DefaultApi.md#DeleteChannel) | **Delete** /v1/Channels/{Sid} | 
[**DeleteFlexFlow**](DefaultApi.md#DeleteFlexFlow) | **Delete** /v1/FlexFlows/{Sid} | 
[**DeleteWebChannel**](DefaultApi.md#DeleteWebChannel) | **Delete** /v1/WebChannels/{Sid} | 
[**FetchChannel**](DefaultApi.md#FetchChannel) | **Get** /v1/Channels/{Sid} | 
[**FetchConfiguration**](DefaultApi.md#FetchConfiguration) | **Get** /v1/Configuration | 
[**FetchFlexFlow**](DefaultApi.md#FetchFlexFlow) | **Get** /v1/FlexFlows/{Sid} | 
[**FetchWebChannel**](DefaultApi.md#FetchWebChannel) | **Get** /v1/WebChannels/{Sid} | 
[**ListChannel**](DefaultApi.md#ListChannel) | **Get** /v1/Channels | 
[**ListFlexFlow**](DefaultApi.md#ListFlexFlow) | **Get** /v1/FlexFlows | 
[**ListWebChannel**](DefaultApi.md#ListWebChannel) | **Get** /v1/WebChannels | 
[**UpdateConfiguration**](DefaultApi.md#UpdateConfiguration) | **Post** /v1/Configuration | 
[**UpdateFlexFlow**](DefaultApi.md#UpdateFlexFlow) | **Post** /v1/FlexFlows/{Sid} | 
[**UpdateWebChannel**](DefaultApi.md#UpdateWebChannel) | **Post** /v1/WebChannels/{Sid} | 



## CreateChannel

> FlexV1Channel CreateChannel(ctx, optional)



### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateChannelParams struct


Name | Type | Description
------------- | ------------- | -------------
**ChatFriendlyName** | **string** | The chat channel&#39;s friendly name.
**ChatUniqueName** | **string** | The chat channel&#39;s unique name.
**ChatUserFriendlyName** | **string** | The chat participant&#39;s friendly name.
**FlexFlowSid** | **string** | The SID of the Flex Flow.
**Identity** | **string** | The &#x60;identity&#x60; value that uniquely identifies the new resource&#39;s chat User.
**LongLived** | **bool** | Whether to create the channel as long-lived.
**PreEngagementData** | **string** | The pre-engagement data.
**Target** | **string** | The Target Contact Identity, for example the phone number of an SMS.
**TaskAttributes** | **string** | The Task attributes to be added for the TaskRouter Task.
**TaskSid** | **string** | The SID of the TaskRouter Task. Only valid when integration type is &#x60;task&#x60;. &#x60;null&#x60; for integration types &#x60;studio&#x60; &amp; &#x60;external&#x60;

### Return type

[**FlexV1Channel**](FlexV1Channel.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded, 
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateFlexFlow

> FlexV1FlexFlow CreateFlexFlow(ctx, optional)



### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateFlexFlowParams struct


Name | Type | Description
------------- | ------------- | -------------
**ChannelType** | **string** | The channel type. Can be: &#x60;web&#x60;, &#x60;facebook&#x60;, &#x60;sms&#x60;, &#x60;whatsapp&#x60;, &#x60;line&#x60; or &#x60;custom&#x60;.
**ChatServiceSid** | **string** | The SID of the chat service.
**ContactIdentity** | **string** | The channel contact&#39;s Identity.
**Enabled** | **bool** | Whether the new Flex Flow is enabled.
**FriendlyName** | **string** | A descriptive string that you create to describe the Flex Flow resource.
**IntegrationChannel** | **string** | The Task Channel for the TaskRouter Task that will be created. Applicable and required when integrationType is &#x60;task&#x60;. Set to &#x60;sms&#x60; for SMS, and to &#x60;chat&#x60; otherwise. The default value is &#x60;default&#x60;
**IntegrationCreationOnMessage** | **bool** | In the context of outbound messaging, defines whether to create a Task immediately (and therefore reserve the conversation to current agent), or delay Task creation until the customer sends the first response. Set to false to create immediately, true to delay Task creation. This setting is only applicable for outbound messaging.
**IntegrationFlowSid** | **string** | The SID of the Studio Flow. Required when &#x60;integrationType&#x60; is &#x60;studio&#x60;.
**IntegrationPriority** | **int32** | The Task priority of a new Task. The default priority is 0. Optional when &#x60;integrationType&#x60; is &#x60;task&#x60;, not applicable otherwise.
**IntegrationRetryCount** | **int32** | The number of times to retry the webhook if the first attempt fails. Can be an integer between 0 and 3 (included), default is 0. Optional when integrationType is &#x60;external&#x60;, not applicable otherwise.
**IntegrationTimeout** | **int32** | The Task timeout in seconds for a new Task. Default is 86,400 seconds (24 hours). Optional when &#x60;integrationType&#x60; is &#x60;task&#x60;, not applicable otherwise.
**IntegrationUrl** | **string** | The URL of the external webhook. Required when &#x60;integrationType&#x60; is &#x60;external&#x60;.
**IntegrationWorkflowSid** | **string** | The Workflow SID for a new Task. Required when &#x60;integrationType&#x60; is &#x60;task&#x60;.
**IntegrationWorkspaceSid** | **string** | The Workspace SID for a new Task. Required when &#x60;integrationType&#x60; is &#x60;task&#x60;.
**IntegrationType** | **string** | The integration type. Can be: &#x60;studio&#x60;, &#x60;external&#x60;, or &#x60;task&#x60;.
**JanitorEnabled** | **bool** | When enabled, the Messaging Channel Janitor will remove active Proxy sessions if the associated Task is deleted outside of the Flex UI. Defaults to &#x60;false&#x60;.
**LongLived** | **bool** | When enabled, Flex will keep the chat channel active so that it may be used for subsequent interactions with a contact identity. Defaults to &#x60;false&#x60;.

### Return type

[**FlexV1FlexFlow**](FlexV1FlexFlow.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded, 
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateWebChannel

> FlexV1WebChannel CreateWebChannel(ctx, optional)



### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateWebChannelParams struct


Name | Type | Description
------------- | ------------- | -------------
**ChatFriendlyName** | **string** | The chat channel&#39;s friendly name.
**ChatUniqueName** | **string** | The chat channel&#39;s unique name.
**CustomerFriendlyName** | **string** | The chat participant&#39;s friendly name.
**FlexFlowSid** | **string** | The SID of the Flex Flow.
**Identity** | **string** | The chat identity.
**PreEngagementData** | **string** | The pre-engagement data.

### Return type

[**FlexV1WebChannel**](FlexV1WebChannel.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded, 
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteChannel

> DeleteChannel(ctx, Sid)



### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The SID of the Flex chat channel resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeleteChannelParams struct


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


## DeleteFlexFlow

> DeleteFlexFlow(ctx, Sid)



### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The SID of the Flex Flow resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeleteFlexFlowParams struct


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


## DeleteWebChannel

> DeleteWebChannel(ctx, Sid)



### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The SID of the WebChannel resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeleteWebChannelParams struct


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


## FetchChannel

> FlexV1Channel FetchChannel(ctx, Sid)



### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The SID of the Flex chat channel resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchChannelParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**FlexV1Channel**](FlexV1Channel.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchConfiguration

> FlexV1Configuration FetchConfiguration(ctx, optional)



### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a FetchConfigurationParams struct


Name | Type | Description
------------- | ------------- | -------------
**UiVersion** | **string** | The Pinned UI version of the Configuration resource to fetch.

### Return type

[**FlexV1Configuration**](FlexV1Configuration.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchFlexFlow

> FlexV1FlexFlow FetchFlexFlow(ctx, Sid)



### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The SID of the Flex Flow resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchFlexFlowParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**FlexV1FlexFlow**](FlexV1FlexFlow.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchWebChannel

> FlexV1WebChannel FetchWebChannel(ctx, Sid)



### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The SID of the WebChannel resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchWebChannelParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**FlexV1WebChannel**](FlexV1WebChannel.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListChannel

> ListChannelResponse ListChannel(ctx, optional)



### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListChannelParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListChannelResponse**](ListChannelResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListFlexFlow

> ListFlexFlowResponse ListFlexFlow(ctx, optional)



### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListFlexFlowParams struct


Name | Type | Description
------------- | ------------- | -------------
**FriendlyName** | **string** | The &#x60;friendly_name&#x60; of the Flex Flow resources to read.
**PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListFlexFlowResponse**](ListFlexFlowResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListWebChannel

> ListWebChannelResponse ListWebChannel(ctx, optional)



### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListWebChannelParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListWebChannelResponse**](ListWebChannelResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateConfiguration

> FlexV1Configuration UpdateConfiguration(ctx, )



### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a UpdateConfigurationParams struct


### Return type

[**FlexV1Configuration**](FlexV1Configuration.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateFlexFlow

> FlexV1FlexFlow UpdateFlexFlow(ctx, Sidoptional)



### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The SID of the Flex Flow resource to update.

### Other Parameters

Other parameters are passed through a pointer to a UpdateFlexFlowParams struct


Name | Type | Description
------------- | ------------- | -------------
**ChannelType** | **string** | The channel type. Can be: &#x60;web&#x60;, &#x60;facebook&#x60;, &#x60;sms&#x60;, &#x60;whatsapp&#x60;, &#x60;line&#x60; or &#x60;custom&#x60;.
**ChatServiceSid** | **string** | The SID of the chat service.
**ContactIdentity** | **string** | The channel contact&#39;s Identity.
**Enabled** | **bool** | Whether the new Flex Flow is enabled.
**FriendlyName** | **string** | A descriptive string that you create to describe the Flex Flow resource.
**IntegrationChannel** | **string** | The Task Channel for the TaskRouter Task that will be created. Applicable and required when integrationType is &#x60;task&#x60;. Set to &#x60;sms&#x60; for SMS, and to &#x60;chat&#x60; otherwise. The default value is &#x60;default&#x60;
**IntegrationCreationOnMessage** | **bool** | In the context of outbound messaging, defines whether to create a Task immediately (and therefore reserve the conversation to current agent), or delay Task creation until the customer sends the first response. Set to false to create immediately, true to delay Task creation. This setting is only applicable for outbound messaging.
**IntegrationFlowSid** | **string** | The SID of the Studio Flow. Required when &#x60;integrationType&#x60; is &#x60;studio&#x60;.
**IntegrationPriority** | **int32** | The Task priority of a new Task. The default priority is 0. Optional when &#x60;integrationType&#x60; is &#x60;task&#x60;, not applicable otherwise.
**IntegrationRetryCount** | **int32** | The number of times to retry the webhook if the first attempt fails. Can be an integer between 0 and 3 (included), default is 0. Optional when integrationType is &#x60;external&#x60;, not applicable otherwise.
**IntegrationTimeout** | **int32** | The Task timeout in seconds for a new Task. Default is 86,400 seconds (24 hours). Optional when &#x60;integrationType&#x60; is &#x60;task&#x60;, not applicable otherwise.
**IntegrationUrl** | **string** | The URL of the external webhook. Required when &#x60;integrationType&#x60; is &#x60;external&#x60;.
**IntegrationWorkflowSid** | **string** | The Workflow SID for a new Task. Required when &#x60;integrationType&#x60; is &#x60;task&#x60;.
**IntegrationWorkspaceSid** | **string** | The Workspace SID for a new Task. Required when &#x60;integrationType&#x60; is &#x60;task&#x60;.
**IntegrationType** | **string** | The integration type. Can be: &#x60;studio&#x60;, &#x60;external&#x60;, or &#x60;task&#x60;.
**JanitorEnabled** | **bool** | When enabled, the Messaging Channel Janitor will remove active Proxy sessions if the associated Task is deleted outside of the Flex UI. Defaults to &#x60;false&#x60;.
**LongLived** | **bool** | When enabled, Flex will keep the chat channel active so that it may be used for subsequent interactions with a contact identity. Defaults to &#x60;false&#x60;.

### Return type

[**FlexV1FlexFlow**](FlexV1FlexFlow.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded, 
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateWebChannel

> FlexV1WebChannel UpdateWebChannel(ctx, Sidoptional)



### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The SID of the WebChannel resource to update.

### Other Parameters

Other parameters are passed through a pointer to a UpdateWebChannelParams struct


Name | Type | Description
------------- | ------------- | -------------
**ChatStatus** | **string** | The chat status. Can only be &#x60;inactive&#x60;.
**PostEngagementData** | **string** | The post-engagement data.

### Return type

[**FlexV1WebChannel**](FlexV1WebChannel.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded, 
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

