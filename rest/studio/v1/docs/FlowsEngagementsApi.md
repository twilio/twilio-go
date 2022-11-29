# FlowsEngagementsApi

All URIs are relative to *https://studio.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateEngagement**](FlowsEngagementsApi.md#CreateEngagement) | **Post** /v1/Flows/{FlowSid}/Engagements | 
[**DeleteEngagement**](FlowsEngagementsApi.md#DeleteEngagement) | **Delete** /v1/Flows/{FlowSid}/Engagements/{Sid} | 
[**FetchEngagement**](FlowsEngagementsApi.md#FetchEngagement) | **Get** /v1/Flows/{FlowSid}/Engagements/{Sid} | 
[**ListEngagement**](FlowsEngagementsApi.md#ListEngagement) | **Get** /v1/Flows/{FlowSid}/Engagements | 



## CreateEngagement

> StudioV1Engagement CreateEngagement(ctx, FlowSidoptional)



Triggers a new Engagement for the Flow

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**FlowSid** | **string** | The SID of the Flow.

### Other Parameters

Other parameters are passed through a pointer to a CreateEngagementParams struct


Name | Type | Description
------------- | ------------- | -------------
**To** | **string** | The Contact phone number to start a Studio Flow Engagement, available as variable `{{contact.channel.address}}`.
**From** | **string** | The Twilio phone number to send messages or initiate calls from during the Flow Engagement. Available as variable `{{flow.channel.address}}`
**Parameters** | [**interface{}**](interface{}.md) | A JSON string we will add to your flow's context and that you can access as variables inside your flow. For example, if you pass in `Parameters={'name':'Zeke'}` then inside a widget you can reference the variable `{{flow.data.name}}` which will return the string 'Zeke'. Note: the JSON value must explicitly be passed as a string, not as a hash object. Depending on your particular HTTP library, you may need to add quotes or URL encode your JSON string.

### Return type

[**StudioV1Engagement**](StudioV1Engagement.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteEngagement

> DeleteEngagement(ctx, FlowSidSid)



Delete this Engagement and all Steps relating to it.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**FlowSid** | **string** | The SID of the Flow to delete Engagements from.
**Sid** | **string** | The SID of the Engagement resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeleteEngagementParams struct


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


## FetchEngagement

> StudioV1Engagement FetchEngagement(ctx, FlowSidSid)



Retrieve an Engagement

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**FlowSid** | **string** | The SID of the Flow.
**Sid** | **string** | The SID of the Engagement resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchEngagementParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**StudioV1Engagement**](StudioV1Engagement.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListEngagement

> []StudioV1Engagement ListEngagement(ctx, FlowSidoptional)



Retrieve a list of all Engagements for the Flow.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**FlowSid** | **string** | The SID of the Flow to read Engagements from.

### Other Parameters

Other parameters are passed through a pointer to a ListEngagementParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]StudioV1Engagement**](StudioV1Engagement.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

