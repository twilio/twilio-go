# AccountsMessagesFeedbackApi

All URIs are relative to *https://api.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateMessageFeedback**](AccountsMessagesFeedbackApi.md#CreateMessageFeedback) | **Post** /2010-04-01/Accounts/{AccountSid}/Messages/{MessageSid}/Feedback.json | 



## CreateMessageFeedback

> ApiV2010MessageFeedback CreateMessageFeedback(ctx, MessageSidoptional)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**MessageSid** | **string** | The SID of the Message resource for which the feedback was provided.

### Other Parameters

Other parameters are passed through a pointer to a CreateMessageFeedbackParams struct


Name | Type | Description
------------- | ------------- | -------------
**PathAccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that will create the resource.
**Outcome** | **string** | 

### Return type

[**ApiV2010MessageFeedback**](ApiV2010MessageFeedback.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

