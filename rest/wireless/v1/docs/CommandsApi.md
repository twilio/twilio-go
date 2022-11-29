# CommandsApi

All URIs are relative to *https://wireless.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCommand**](CommandsApi.md#CreateCommand) | **Post** /v1/Commands | 
[**DeleteCommand**](CommandsApi.md#DeleteCommand) | **Delete** /v1/Commands/{Sid} | 
[**FetchCommand**](CommandsApi.md#FetchCommand) | **Get** /v1/Commands/{Sid} | 
[**ListCommand**](CommandsApi.md#ListCommand) | **Get** /v1/Commands | 



## CreateCommand

> WirelessV1Command CreateCommand(ctx, optional)



Send a Command to a Sim.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateCommandParams struct


Name | Type | Description
------------- | ------------- | -------------
**Command** | **string** | The message body of the Command. Can be plain text in text mode or a Base64 encoded byte string in binary mode.
**Sim** | **string** | The `sid` or `unique_name` of the [SIM](https://www.twilio.com/docs/wireless/api/sim-resource) to send the Command to.
**CallbackMethod** | **string** | The HTTP method we use to call `callback_url`. Can be: `POST` or `GET`, and the default is `POST`.
**CallbackUrl** | **string** | The URL we call using the `callback_url` when the Command has finished sending, whether the command was delivered or it failed.
**CommandMode** | **string** | 
**IncludeSid** | **string** | Whether to include the SID of the command in the message body. Can be: `none`, `start`, or `end`, and the default behavior is `none`. When sending a Command to a SIM in text mode, we can automatically include the SID of the Command in the message body, which could be used to ensure that the device does not process the same Command more than once.  A value of `start` will prepend the message with the Command SID, and `end` will append it to the end, separating the Command SID from the message body with a space. The length of the Command SID is included in the 160 character limit so the SMS body must be 128 characters or less before the Command SID is included.
**DeliveryReceiptRequested** | **bool** | Whether to request delivery receipt from the recipient. For Commands that request delivery receipt, the Command state transitions to 'delivered' once the server has received a delivery receipt from the device. The default value is `true`.

### Return type

[**WirelessV1Command**](WirelessV1Command.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCommand

> DeleteCommand(ctx, Sid)



Delete a Command instance from your account.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The SID of the Command resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeleteCommandParams struct


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


## FetchCommand

> WirelessV1Command FetchCommand(ctx, Sid)



Fetch a Command instance from your account.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The SID of the Command resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchCommandParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**WirelessV1Command**](WirelessV1Command.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCommand

> []WirelessV1Command ListCommand(ctx, optional)



Retrieve a list of Commands from your account.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListCommandParams struct


Name | Type | Description
------------- | ------------- | -------------
**Sim** | **string** | The `sid` or `unique_name` of the [Sim resources](https://www.twilio.com/docs/wireless/api/sim-resource) to read.
**Status** | **string** | The status of the resources to read. Can be: `queued`, `sent`, `delivered`, `received`, or `failed`.
**Direction** | **string** | Only return Commands with this direction value.
**Transport** | **string** | Only return Commands with this transport value. Can be: `sms` or `ip`.
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]WirelessV1Command**](WirelessV1Command.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

