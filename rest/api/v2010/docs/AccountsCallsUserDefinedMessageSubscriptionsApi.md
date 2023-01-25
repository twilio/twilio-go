# AccountsCallsUserDefinedMessageSubscriptionsApi

All URIs are relative to *https://api.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateUserDefinedMessageSubscription**](AccountsCallsUserDefinedMessageSubscriptionsApi.md#CreateUserDefinedMessageSubscription) | **Post** /2010-04-01/Accounts/{AccountSid}/Calls/{CallSid}/UserDefinedMessageSubscriptions.json | 
[**DeleteUserDefinedMessageSubscription**](AccountsCallsUserDefinedMessageSubscriptionsApi.md#DeleteUserDefinedMessageSubscription) | **Delete** /2010-04-01/Accounts/{AccountSid}/Calls/{CallSid}/UserDefinedMessageSubscriptions/{Sid}.json | 



## CreateUserDefinedMessageSubscription

> ApiV2010UserDefinedMessageSubscription CreateUserDefinedMessageSubscription(ctx, CallSidoptional)



Subscribe to User Defined Messages for a given Call SID.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**CallSid** | **string** | The SID of the [Call](https://www.twilio.com/docs/voice/api/call-resource) the User Defined Messages subscription is associated with. This refers to the Call SID that is producing the user defined messages.

### Other Parameters

Other parameters are passed through a pointer to a CreateUserDefinedMessageSubscriptionParams struct


Name | Type | Description
------------- | ------------- | -------------
**PathAccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that subscribed to the User Defined Messages.
**Callback** | **string** | The URL we should call using the `method` to send user defined events to your application. URLs must contain a valid hostname (underscores are not permitted).
**IdempotencyKey** | **string** | A unique string value to identify API call. This should be a unique string value per API call and can be a randomly generated.
**Method** | **string** | The HTTP method Twilio will use when requesting the above `Url`. Either `GET` or `POST`. Default is `POST`.

### Return type

[**ApiV2010UserDefinedMessageSubscription**](ApiV2010UserDefinedMessageSubscription.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteUserDefinedMessageSubscription

> DeleteUserDefinedMessageSubscription(ctx, CallSidSidoptional)



Delete a specific User Defined Message Subscription.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**CallSid** | **string** | The SID of the [Call](https://www.twilio.com/docs/voice/api/call-resource) the User Defined Message Subscription is associated with. This refers to the Call SID that is producing the User Defined Messages.
**Sid** | **string** | The SID that uniquely identifies this User Defined Message Subscription.

### Other Parameters

Other parameters are passed through a pointer to a DeleteUserDefinedMessageSubscriptionParams struct


Name | Type | Description
------------- | ------------- | -------------
**PathAccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that subscribed to the User Defined Messages.

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

