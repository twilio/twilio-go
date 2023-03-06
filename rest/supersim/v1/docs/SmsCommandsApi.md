# SmsCommandsApi

All URIs are relative to *https://supersim.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSmsCommand**](SmsCommandsApi.md#CreateSmsCommand) | **Post** /v1/SmsCommands | 
[**FetchSmsCommand**](SmsCommandsApi.md#FetchSmsCommand) | **Get** /v1/SmsCommands/{Sid} | 
[**ListSmsCommand**](SmsCommandsApi.md#ListSmsCommand) | **Get** /v1/SmsCommands | 



## CreateSmsCommand

> SupersimV1SmsCommand CreateSmsCommand(ctx, optional)



Send SMS Command to a Sim.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateSmsCommandParams struct


Name | Type | Description
------------- | ------------- | -------------
**Sim** | **string** | The `sid` or `unique_name` of the [SIM](https://www.twilio.com/docs/iot/supersim/api/sim-resource) to send the SMS Command to.
**Payload** | **string** | The message body of the SMS Command.
**CallbackMethod** | **string** | The HTTP method we should use to call `callback_url`. Can be: `GET` or `POST` and the default is POST.
**CallbackUrl** | **string** | The URL we should call using the `callback_method` after we have sent the command.

### Return type

[**SupersimV1SmsCommand**](SupersimV1SmsCommand.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchSmsCommand

> SupersimV1SmsCommand FetchSmsCommand(ctx, Sid)



Fetch SMS Command instance from your account.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The SID of the SMS Command resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchSmsCommandParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**SupersimV1SmsCommand**](SupersimV1SmsCommand.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSmsCommand

> []SupersimV1SmsCommand ListSmsCommand(ctx, optional)



Retrieve a list of SMS Commands from your account.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListSmsCommandParams struct


Name | Type | Description
------------- | ------------- | -------------
**Sim** | **string** | The SID or unique name of the Sim resource that SMS Command was sent to or from.
**Status** | **string** | The status of the SMS Command. Can be: `queued`, `sent`, `delivered`, `received` or `failed`. See the [SMS Command Status Values](https://www.twilio.com/docs/iot/supersim/api/smscommand-resource#status-values) for a description of each.
**Direction** | **string** | The direction of the SMS Command. Can be `to_sim` or `from_sim`. The value of `to_sim` is synonymous with the term `mobile terminated`, and `from_sim` is synonymous with the term `mobile originated`.
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]SupersimV1SmsCommand**](SupersimV1SmsCommand.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

