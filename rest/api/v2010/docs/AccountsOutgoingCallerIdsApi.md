# AccountsOutgoingCallerIdsApi

All URIs are relative to *https://api.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateValidationRequest**](AccountsOutgoingCallerIdsApi.md#CreateValidationRequest) | **Post** /2010-04-01/Accounts/{AccountSid}/OutgoingCallerIds.json | 
[**DeleteOutgoingCallerId**](AccountsOutgoingCallerIdsApi.md#DeleteOutgoingCallerId) | **Delete** /2010-04-01/Accounts/{AccountSid}/OutgoingCallerIds/{Sid}.json | 
[**FetchOutgoingCallerId**](AccountsOutgoingCallerIdsApi.md#FetchOutgoingCallerId) | **Get** /2010-04-01/Accounts/{AccountSid}/OutgoingCallerIds/{Sid}.json | 
[**ListOutgoingCallerId**](AccountsOutgoingCallerIdsApi.md#ListOutgoingCallerId) | **Get** /2010-04-01/Accounts/{AccountSid}/OutgoingCallerIds.json | 
[**UpdateOutgoingCallerId**](AccountsOutgoingCallerIdsApi.md#UpdateOutgoingCallerId) | **Post** /2010-04-01/Accounts/{AccountSid}/OutgoingCallerIds/{Sid}.json | 



## CreateValidationRequest

> ApiV2010ValidationRequest CreateValidationRequest(ctx, optional)





### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateValidationRequestParams struct


Name | Type | Description
------------- | ------------- | -------------
**PathAccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) responsible for the new caller ID resource.
**PhoneNumber** | **string** | The phone number to verify in [E.164](https://www.twilio.com/docs/glossary/what-e164) format, which consists of a + followed by the country code and subscriber number.
**FriendlyName** | **string** | A descriptive string that you create to describe the new caller ID resource. It can be up to 64 characters long. The default value is a formatted version of the phone number.
**CallDelay** | **int** | The number of seconds to delay before initiating the verification call. Can be an integer between `0` and `60`, inclusive. The default is `0`.
**Extension** | **string** | The digits to dial after connecting the verification call.
**StatusCallback** | **string** | The URL we should call using the `status_callback_method` to send status information about the verification process to your application.
**StatusCallbackMethod** | **string** | The HTTP method we should use to call `status_callback`. Can be: `GET` or `POST`, and the default is `POST`.

### Return type

[**ApiV2010ValidationRequest**](ApiV2010ValidationRequest.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOutgoingCallerId

> DeleteOutgoingCallerId(ctx, Sidoptional)



Delete the caller-id specified from the account

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the OutgoingCallerId resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeleteOutgoingCallerIdParams struct


Name | Type | Description
------------- | ------------- | -------------
**PathAccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the OutgoingCallerId resources to delete.

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


## FetchOutgoingCallerId

> ApiV2010OutgoingCallerId FetchOutgoingCallerId(ctx, Sidoptional)



Fetch an outgoing-caller-id belonging to the account used to make the request

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the OutgoingCallerId resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchOutgoingCallerIdParams struct


Name | Type | Description
------------- | ------------- | -------------
**PathAccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the OutgoingCallerId resource to fetch.

### Return type

[**ApiV2010OutgoingCallerId**](ApiV2010OutgoingCallerId.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOutgoingCallerId

> []ApiV2010OutgoingCallerId ListOutgoingCallerId(ctx, optional)



Retrieve a list of outgoing-caller-ids belonging to the account used to make the request

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListOutgoingCallerIdParams struct


Name | Type | Description
------------- | ------------- | -------------
**PathAccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the OutgoingCallerId resources to read.
**PhoneNumber** | **string** | The phone number of the OutgoingCallerId resources to read.
**FriendlyName** | **string** | The string that identifies the OutgoingCallerId resources to read.
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]ApiV2010OutgoingCallerId**](ApiV2010OutgoingCallerId.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOutgoingCallerId

> ApiV2010OutgoingCallerId UpdateOutgoingCallerId(ctx, Sidoptional)



Updates the caller-id

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the OutgoingCallerId resource to update.

### Other Parameters

Other parameters are passed through a pointer to a UpdateOutgoingCallerIdParams struct


Name | Type | Description
------------- | ------------- | -------------
**PathAccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the OutgoingCallerId resources to update.
**FriendlyName** | **string** | A descriptive string that you create to describe the resource. It can be up to 64 characters long.

### Return type

[**ApiV2010OutgoingCallerId**](ApiV2010OutgoingCallerId.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

