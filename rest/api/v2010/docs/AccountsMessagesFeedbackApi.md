# AccountsMessagesFeedbackApi

All URIs are relative to *https://api.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateMessageFeedback**](AccountsMessagesFeedbackApi.md#CreateMessageFeedback) | **Post** /2010-04-01/Accounts/{AccountSid}/Messages/{MessageSid}/Feedback.json | Create Message Feedback to confirm a tracked user action was performed by the recipient of the associated Message



## CreateMessageFeedback

> ApiV2010MessageFeedback CreateMessageFeedback(ctx, MessageSidoptional)

Create Message Feedback to confirm a tracked user action was performed by the recipient of the associated Message

Create Message Feedback to confirm a tracked user action was performed by the recipient of the associated Message

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**MessageSid** | **string** | The SID of the Message resource for which to create MessageFeedback.

### Other Parameters

Other parameters are passed through a pointer to a CreateMessageFeedbackParams struct


Name | Type | Description
------------- | ------------- | -------------
**PathAccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) associated with the Message resource for which to create MessageFeedback.
**Outcome** | [**string**](string.md) | 

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

