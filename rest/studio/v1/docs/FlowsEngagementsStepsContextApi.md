# FlowsEngagementsStepsContextApi

All URIs are relative to *https://studio.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchStepContext**](FlowsEngagementsStepsContextApi.md#FetchStepContext) | **Get** /v1/Flows/{FlowSid}/Engagements/{EngagementSid}/Steps/{StepSid}/Context | 



## FetchStepContext

> StudioV1StepContext FetchStepContext(ctx, FlowSidEngagementSidStepSid)



Retrieve the context for an Engagement Step.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**FlowSid** | **string** | The SID of the Flow with the Step to fetch.
**EngagementSid** | **string** | The SID of the Engagement with the Step to fetch.
**StepSid** | **string** | The SID of the Step to fetch

### Other Parameters

Other parameters are passed through a pointer to a FetchStepContextParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**StudioV1StepContext**](StudioV1StepContext.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

