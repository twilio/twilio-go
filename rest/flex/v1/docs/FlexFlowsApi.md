# FlexFlowsApi

All URIs are relative to *https://flex-api.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateFlexFlow**](FlexFlowsApi.md#CreateFlexFlow) | **Post** /v1/FlexFlows | 
[**DeleteFlexFlow**](FlexFlowsApi.md#DeleteFlexFlow) | **Delete** /v1/FlexFlows/{Sid} | 
[**FetchFlexFlow**](FlexFlowsApi.md#FetchFlexFlow) | **Get** /v1/FlexFlows/{Sid} | 
[**ListFlexFlow**](FlexFlowsApi.md#ListFlexFlow) | **Get** /v1/FlexFlows | 
[**UpdateFlexFlow**](FlexFlowsApi.md#UpdateFlexFlow) | **Post** /v1/FlexFlows/{Sid} | 



## CreateFlexFlow

> FlexV1FlexFlow CreateFlexFlow(ctx, optional)





### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateFlexFlowParams struct


Name | Type | Description
------------- | ------------- | -------------
**FriendlyName** | **string** | A descriptive string that you create to describe the Flex Flow resource.
**ChatServiceSid** | **string** | The SID of the chat service.
**ChannelType** | **string** | 
**ContactIdentity** | **string** | The channel contact's Identity.
**Enabled** | **bool** | Whether the new Flex Flow is enabled.
**IntegrationType** | **string** | 
**IntegrationFlowSid** | **string** | The SID of the Studio Flow. Required when `integrationType` is `studio`.
**IntegrationUrl** | **string** | The URL of the external webhook. Required when `integrationType` is `external`.
**IntegrationWorkspaceSid** | **string** | The Workspace SID for a new Task. Required when `integrationType` is `task`.
**IntegrationWorkflowSid** | **string** | The Workflow SID for a new Task. Required when `integrationType` is `task`.
**IntegrationChannel** | **string** | The Task Channel SID (TCXXXX) or unique name (e.g., `sms`) to use for the Task that will be created. Applicable and required when `integrationType` is `task`. The default value is `default`.
**IntegrationTimeout** | **int** | The Task timeout in seconds for a new Task. Default is 86,400 seconds (24 hours). Optional when `integrationType` is `task`, not applicable otherwise.
**IntegrationPriority** | **int** | The Task priority of a new Task. The default priority is 0. Optional when `integrationType` is `task`, not applicable otherwise.
**IntegrationCreationOnMessage** | **bool** | In the context of outbound messaging, defines whether to create a Task immediately (and therefore reserve the conversation to current agent), or delay Task creation until the customer sends the first response. Set to false to create immediately, true to delay Task creation. This setting is only applicable for outbound messaging.
**LongLived** | **bool** | When enabled, Flex will keep the chat channel active so that it may be used for subsequent interactions with a contact identity. Defaults to `false`.
**JanitorEnabled** | **bool** | When enabled, the Messaging Channel Janitor will remove active Proxy sessions if the associated Task is deleted outside of the Flex UI. Defaults to `false`.
**IntegrationRetryCount** | **int** | The number of times to retry the Studio Flow or webhook in case of failure. Takes integer values from 0 to 3 with the default being 3. Optional when `integrationType` is `studio` or `external`, not applicable otherwise.

### Return type

[**FlexV1FlexFlow**](FlexV1FlexFlow.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

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
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListFlexFlow

> []FlexV1FlexFlow ListFlexFlow(ctx, optional)





### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListFlexFlowParams struct


Name | Type | Description
------------- | ------------- | -------------
**FriendlyName** | **string** | The `friendly_name` of the Flex Flow resources to read.
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]FlexV1FlexFlow**](FlexV1FlexFlow.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

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
**FriendlyName** | **string** | A descriptive string that you create to describe the Flex Flow resource.
**ChatServiceSid** | **string** | The SID of the chat service.
**ChannelType** | **string** | 
**ContactIdentity** | **string** | The channel contact's Identity.
**Enabled** | **bool** | Whether the new Flex Flow is enabled.
**IntegrationType** | **string** | 
**IntegrationFlowSid** | **string** | The SID of the Studio Flow. Required when `integrationType` is `studio`.
**IntegrationUrl** | **string** | The URL of the external webhook. Required when `integrationType` is `external`.
**IntegrationWorkspaceSid** | **string** | The Workspace SID for a new Task. Required when `integrationType` is `task`.
**IntegrationWorkflowSid** | **string** | The Workflow SID for a new Task. Required when `integrationType` is `task`.
**IntegrationChannel** | **string** | The Task Channel SID (TCXXXX) or unique name (e.g., `sms`) to use for the Task that will be created. Applicable and required when `integrationType` is `task`. The default value is `default`.
**IntegrationTimeout** | **int** | The Task timeout in seconds for a new Task. Default is 86,400 seconds (24 hours). Optional when `integrationType` is `task`, not applicable otherwise.
**IntegrationPriority** | **int** | The Task priority of a new Task. The default priority is 0. Optional when `integrationType` is `task`, not applicable otherwise.
**IntegrationCreationOnMessage** | **bool** | In the context of outbound messaging, defines whether to create a Task immediately (and therefore reserve the conversation to current agent), or delay Task creation until the customer sends the first response. Set to false to create immediately, true to delay Task creation. This setting is only applicable for outbound messaging.
**LongLived** | **bool** | When enabled, Flex will keep the chat channel active so that it may be used for subsequent interactions with a contact identity. Defaults to `false`.
**JanitorEnabled** | **bool** | When enabled, the Messaging Channel Janitor will remove active Proxy sessions if the associated Task is deleted outside of the Flex UI. Defaults to `false`.
**IntegrationRetryCount** | **int** | The number of times to retry the Studio Flow or webhook in case of failure. Takes integer values from 0 to 3 with the default being 3. Optional when `integrationType` is `studio` or `external`, not applicable otherwise.

### Return type

[**FlexV1FlexFlow**](FlexV1FlexFlow.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

