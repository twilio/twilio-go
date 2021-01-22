# \DefaultApi

All URIs are relative to *http://localhost*

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



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateChannelOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateChannelOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ChatFriendlyName** | **optional.String**| The chat channel&#39;s friendly name. | 
 **ChatUniqueName** | **optional.String**| The chat channel&#39;s unique name. | 
 **ChatUserFriendlyName** | **optional.String**| The chat participant&#39;s friendly name. | 
 **FlexFlowSid** | **optional.String**| The SID of the Flex Flow. | 
 **Identity** | **optional.String**| The &#x60;identity&#x60; value that uniquely identifies the new resource&#39;s chat User. | 
 **LongLived** | **optional.Bool**| Whether to create the channel as long-lived. | 
 **PreEngagementData** | **optional.String**| The pre-engagement data. | 
 **Target** | **optional.String**| The Target Contact Identity, for example the phone number of an SMS. | 
 **TaskAttributes** | **optional.String**| The Task attributes to be added for the TaskRouter Task. | 
 **TaskSid** | **optional.String**| The SID of the TaskRouter Task. Only valid when integration type is &#x60;task&#x60;. &#x60;null&#x60; for integration types &#x60;studio&#x60; &amp; &#x60;external&#x60; | 

### Return type

[**FlexV1Channel**](flex.v1.channel.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateFlexFlow

> FlexV1FlexFlow CreateFlexFlow(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateFlexFlowOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateFlexFlowOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ChannelType** | **optional.String**| The channel type. Can be: &#x60;web&#x60;, &#x60;facebook&#x60;, &#x60;sms&#x60;, &#x60;whatsapp&#x60;, &#x60;line&#x60; or &#x60;custom&#x60;. | 
 **ChatServiceSid** | **optional.String**| The SID of the chat service. | 
 **ContactIdentity** | **optional.String**| The channel contact&#39;s Identity. | 
 **Enabled** | **optional.Bool**| Whether the new Flex Flow is enabled. | 
 **FriendlyName** | **optional.String**| A descriptive string that you create to describe the Flex Flow resource. | 
 **IntegrationChannel** | **optional.String**| The Task Channel for the TaskRouter Task that will be created. Applicable and required when integrationType is &#x60;task&#x60;. Set to &#x60;sms&#x60; for SMS, and to &#x60;chat&#x60; otherwise. The default value is &#x60;default&#x60; | 
 **IntegrationCreationOnMessage** | **optional.Bool**| In the context of outbound messaging, defines whether to create a Task immediately (and therefore reserve the conversation to current agent), or delay Task creation until the customer sends the first response. Set to false to create immediately, true to delay Task creation. This setting is only applicable for outbound messaging. | 
 **IntegrationFlowSid** | **optional.String**| The SID of the Studio Flow. Required when &#x60;integrationType&#x60; is &#x60;studio&#x60;. | 
 **IntegrationPriority** | **optional.Int32**| The Task priority of a new Task. The default priority is 0. Optional when &#x60;integrationType&#x60; is &#x60;task&#x60;, not applicable otherwise. | 
 **IntegrationRetryCount** | **optional.Int32**| The number of times to retry the webhook if the first attempt fails. Can be an integer between 0 and 3 (included), default is 0. Optional when integrationType is &#x60;external&#x60;, not applicable otherwise. | 
 **IntegrationTimeout** | **optional.Int32**| The Task timeout in seconds for a new Task. Default is 86,400 seconds (24 hours). Optional when &#x60;integrationType&#x60; is &#x60;task&#x60;, not applicable otherwise. | 
 **IntegrationUrl** | **optional.String**| The URL of the external webhook. Required when &#x60;integrationType&#x60; is &#x60;external&#x60;. | 
 **IntegrationWorkflowSid** | **optional.String**| The Workflow SID for a new Task. Required when &#x60;integrationType&#x60; is &#x60;task&#x60;. | 
 **IntegrationWorkspaceSid** | **optional.String**| The Workspace SID for a new Task. Required when &#x60;integrationType&#x60; is &#x60;task&#x60;. | 
 **IntegrationType** | **optional.String**| The integration type. Can be: &#x60;studio&#x60;, &#x60;external&#x60;, or &#x60;task&#x60;. | 
 **JanitorEnabled** | **optional.Bool**| When enabled, the Messaging Channel Janitor will remove active Proxy sessions if the associated Task is deleted outside of the Flex UI. Defaults to &#x60;false&#x60;. | 
 **LongLived** | **optional.Bool**| When enabled, Flex will keep the chat channel active so that it may be used for subsequent interactions with a contact identity. Defaults to &#x60;false&#x60;. | 

### Return type

[**FlexV1FlexFlow**](flex.v1.flex_flow.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateWebChannel

> FlexV1WebChannel CreateWebChannel(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateWebChannelOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateWebChannelOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ChatFriendlyName** | **optional.String**| The chat channel&#39;s friendly name. | 
 **ChatUniqueName** | **optional.String**| The chat channel&#39;s unique name. | 
 **CustomerFriendlyName** | **optional.String**| The chat participant&#39;s friendly name. | 
 **FlexFlowSid** | **optional.String**| The SID of the Flex Flow. | 
 **Identity** | **optional.String**| The chat identity. | 
 **PreEngagementData** | **optional.String**| The pre-engagement data. | 

### Return type

[**FlexV1WebChannel**](flex.v1.web_channel.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteChannel

> DeleteChannel(ctx, Sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string**| The SID of the Flex chat channel resource to delete. | 

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



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string**| The SID of the Flex Flow resource to delete. | 

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



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string**| The SID of the WebChannel resource to delete. | 

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



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string**| The SID of the Flex chat channel resource to fetch. | 

### Return type

[**FlexV1Channel**](flex.v1.channel.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchConfiguration

> FlexV1Configuration FetchConfiguration(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***FetchConfigurationOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a FetchConfigurationOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **UiVersion** | **optional.String**| The Pinned UI version of the Configuration resource to fetch. | 

### Return type

[**FlexV1Configuration**](flex.v1.configuration.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchFlexFlow

> FlexV1FlexFlow FetchFlexFlow(ctx, Sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string**| The SID of the Flex Flow resource to fetch. | 

### Return type

[**FlexV1FlexFlow**](flex.v1.flex_flow.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchWebChannel

> FlexV1WebChannel FetchWebChannel(ctx, Sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string**| The SID of the WebChannel resource to fetch. | 

### Return type

[**FlexV1WebChannel**](flex.v1.web_channel.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListChannel

> FlexV1ChannelReadResponse ListChannel(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListChannelOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListChannelOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **PageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**FlexV1ChannelReadResponse**](flex_v1_channelReadResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListFlexFlow

> FlexV1FlexFlowReadResponse ListFlexFlow(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListFlexFlowOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListFlexFlowOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **FriendlyName** | **optional.String**| The &#x60;friendly_name&#x60; of the Flex Flow resources to read. | 
 **PageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**FlexV1FlexFlowReadResponse**](flex_v1_flex_flowReadResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListWebChannel

> FlexV1WebChannelReadResponse ListWebChannel(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListWebChannelOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListWebChannelOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **PageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**FlexV1WebChannelReadResponse**](flex_v1_web_channelReadResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateConfiguration

> FlexV1Configuration UpdateConfiguration(ctx, )



### Required Parameters

This endpoint does not need any parameter.

### Return type

[**FlexV1Configuration**](flex.v1.configuration.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateFlexFlow

> FlexV1FlexFlow UpdateFlexFlow(ctx, Sid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string**| The SID of the Flex Flow resource to update. | 
 **optional** | ***UpdateFlexFlowOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateFlexFlowOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ChannelType** | **optional.String**| The channel type. Can be: &#x60;web&#x60;, &#x60;facebook&#x60;, &#x60;sms&#x60;, &#x60;whatsapp&#x60;, &#x60;line&#x60; or &#x60;custom&#x60;. | 
 **ChatServiceSid** | **optional.String**| The SID of the chat service. | 
 **ContactIdentity** | **optional.String**| The channel contact&#39;s Identity. | 
 **Enabled** | **optional.Bool**| Whether the new Flex Flow is enabled. | 
 **FriendlyName** | **optional.String**| A descriptive string that you create to describe the Flex Flow resource. | 
 **IntegrationChannel** | **optional.String**| The Task Channel for the TaskRouter Task that will be created. Applicable and required when integrationType is &#x60;task&#x60;. Set to &#x60;sms&#x60; for SMS, and to &#x60;chat&#x60; otherwise. The default value is &#x60;default&#x60; | 
 **IntegrationCreationOnMessage** | **optional.Bool**| In the context of outbound messaging, defines whether to create a Task immediately (and therefore reserve the conversation to current agent), or delay Task creation until the customer sends the first response. Set to false to create immediately, true to delay Task creation. This setting is only applicable for outbound messaging. | 
 **IntegrationFlowSid** | **optional.String**| The SID of the Studio Flow. Required when &#x60;integrationType&#x60; is &#x60;studio&#x60;. | 
 **IntegrationPriority** | **optional.Int32**| The Task priority of a new Task. The default priority is 0. Optional when &#x60;integrationType&#x60; is &#x60;task&#x60;, not applicable otherwise. | 
 **IntegrationRetryCount** | **optional.Int32**| The number of times to retry the webhook if the first attempt fails. Can be an integer between 0 and 3 (included), default is 0. Optional when integrationType is &#x60;external&#x60;, not applicable otherwise. | 
 **IntegrationTimeout** | **optional.Int32**| The Task timeout in seconds for a new Task. Default is 86,400 seconds (24 hours). Optional when &#x60;integrationType&#x60; is &#x60;task&#x60;, not applicable otherwise. | 
 **IntegrationUrl** | **optional.String**| The URL of the external webhook. Required when &#x60;integrationType&#x60; is &#x60;external&#x60;. | 
 **IntegrationWorkflowSid** | **optional.String**| The Workflow SID for a new Task. Required when &#x60;integrationType&#x60; is &#x60;task&#x60;. | 
 **IntegrationWorkspaceSid** | **optional.String**| The Workspace SID for a new Task. Required when &#x60;integrationType&#x60; is &#x60;task&#x60;. | 
 **IntegrationType** | **optional.String**| The integration type. Can be: &#x60;studio&#x60;, &#x60;external&#x60;, or &#x60;task&#x60;. | 
 **JanitorEnabled** | **optional.Bool**| When enabled, the Messaging Channel Janitor will remove active Proxy sessions if the associated Task is deleted outside of the Flex UI. Defaults to &#x60;false&#x60;. | 
 **LongLived** | **optional.Bool**| When enabled, Flex will keep the chat channel active so that it may be used for subsequent interactions with a contact identity. Defaults to &#x60;false&#x60;. | 

### Return type

[**FlexV1FlexFlow**](flex.v1.flex_flow.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateWebChannel

> FlexV1WebChannel UpdateWebChannel(ctx, Sid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string**| The SID of the WebChannel resource to update. | 
 **optional** | ***UpdateWebChannelOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateWebChannelOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ChatStatus** | **optional.String**| The chat status. Can only be &#x60;inactive&#x60;. | 
 **PostEngagementData** | **optional.String**| The post-engagement data. | 

### Return type

[**FlexV1WebChannel**](flex.v1.web_channel.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

