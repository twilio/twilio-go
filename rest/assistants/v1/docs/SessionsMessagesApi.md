# SessionsMessagesApi

All URIs are relative to *https://assistants.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListMessages**](SessionsMessagesApi.md#ListMessages) | **Get** /v1/Sessions/{sessionId}/Messages | List messages



## ListMessages

> []AssistantsV1Message ListMessages(ctx, SessionIdoptional)

List messages

List messages

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**SessionId** | **string** | Session id or name

### Other Parameters

Other parameters are passed through a pointer to a ListMessagesParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]AssistantsV1Message**](AssistantsV1Message.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

