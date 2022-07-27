# AccountsCallsFeedbackApi

All URIs are relative to *https://api.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchCallFeedback**](AccountsCallsFeedbackApi.md#FetchCallFeedback) | **Get** /2010-04-01/Accounts/{AccountSid}/Calls/{CallSid}/Feedback.json | 
[**UpdateCallFeedback**](AccountsCallsFeedbackApi.md#UpdateCallFeedback) | **Post** /2010-04-01/Accounts/{AccountSid}/Calls/{CallSid}/Feedback.json | 



## FetchCallFeedback

> ApiV2010CallFeedback FetchCallFeedback(ctx, CallSidoptional)



Fetch a Feedback resource from a call

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**CallSid** | **string** | The call sid that uniquely identifies the call

### Other Parameters

Other parameters are passed through a pointer to a FetchCallFeedbackParams struct


Name | Type | Description
------------- | ------------- | -------------
**PathAccountSid** | **string** | The unique id of the [Account](https://www.twilio.com/docs/iam/api/account) responsible for this resource.

### Return type

[**ApiV2010CallFeedback**](ApiV2010CallFeedback.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCallFeedback

> ApiV2010CallFeedback UpdateCallFeedback(ctx, CallSidoptional)



Update a Feedback resource for a call

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**CallSid** | **string** | The call sid that uniquely identifies the call

### Other Parameters

Other parameters are passed through a pointer to a UpdateCallFeedbackParams struct


Name | Type | Description
------------- | ------------- | -------------
**PathAccountSid** | **string** | The unique id of the [Account](https://www.twilio.com/docs/iam/api/account) responsible for this resource.
**QualityScore** | **int** | The call quality expressed as an integer from &#x60;1&#x60; to &#x60;5&#x60; where &#x60;1&#x60; represents very poor call quality and &#x60;5&#x60; represents a perfect call.
**Issue** | [**[]CallFeedbackEnumIssues**](CallFeedbackEnumIssues.md) | One or more issues experienced during the call. The issues can be: &#x60;imperfect-audio&#x60;, &#x60;dropped-call&#x60;, &#x60;incorrect-caller-id&#x60;, &#x60;post-dial-delay&#x60;, &#x60;digits-not-captured&#x60;, &#x60;audio-latency&#x60;, &#x60;unsolicited-call&#x60;, or &#x60;one-way-audio&#x60;.

### Return type

[**ApiV2010CallFeedback**](ApiV2010CallFeedback.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

