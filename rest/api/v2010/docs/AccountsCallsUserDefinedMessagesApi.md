# AccountsCallsUserDefinedMessagesApi

All URIs are relative to *https://api.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateUserDefinedMessage**](AccountsCallsUserDefinedMessagesApi.md#CreateUserDefinedMessage) | **Post** /2010-04-01/Accounts/{AccountSid}/Calls/{CallSid}/UserDefinedMessages.json | 



## CreateUserDefinedMessage

> ApiV2010UserDefinedMessage CreateUserDefinedMessage(ctx, CallSidoptional)



Create a new User Defined Message for the given Call SID.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**CallSid** | **string** | The SID of the [Call](https://www.twilio.com/docs/voice/api/call-resource) the User Defined Message is associated with.

### Other Parameters

Other parameters are passed through a pointer to a CreateUserDefinedMessageParams struct


Name | Type | Description
------------- | ------------- | -------------
**PathAccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created User Defined Message.
**Content** | **string** | The User Defined Message in the form of URL-encoded JSON string.
**IdempotencyKey** | **string** | A unique string value to identify API call. This should be a unique string value per API call and can be a randomly generated.

### Return type

[**ApiV2010UserDefinedMessage**](ApiV2010UserDefinedMessage.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

