# FlowsEngagementsContextApi

All URIs are relative to *https://studio.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchEngagementContext**](FlowsEngagementsContextApi.md#FetchEngagementContext) | **Get** /v1/Flows/{FlowSid}/Engagements/{EngagementSid}/Context | 



## FetchEngagementContext

> StudioV1EngagementContext FetchEngagementContext(ctx, FlowSidEngagementSid)



Retrieve the most recent context for an Engagement.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**FlowSid** | **string** | The SID of the Flow.
**EngagementSid** | **string** | The SID of the Engagement.

### Other Parameters

Other parameters are passed through a pointer to a FetchEngagementContextParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**StudioV1EngagementContext**](StudioV1EngagementContext.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

