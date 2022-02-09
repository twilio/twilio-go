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
**CallbackMethod** | **string** | The HTTP method we should use to call &#x60;callback_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60; and the default is POST.
**CallbackUrl** | **string** | The URL we should call using the &#x60;callback_method&#x60; after we have sent the command.
**Payload** | **string** | The message body of the SMS Command.
**Sim** | **string** | The &#x60;sid&#x60; or &#x60;unique_name&#x60; of the [SIM](https://www.twilio.com/docs/iot/supersim/api/sim-resource) to send the SMS Command to.

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

> ListSmsCommandResponse ListSmsCommand(ctx, optional)



Retrieve a list of SMS Commands from your account.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListSmsCommandParams struct


Name | Type | Description
------------- | ------------- | -------------
**Sim** | **string** | The SID or unique name of the Sim resource that SMS Command was sent to or from.
**Status** | **string** | The status of the SMS Command. Can be: &#x60;queued&#x60;, &#x60;sent&#x60;, &#x60;delivered&#x60;, &#x60;received&#x60; or &#x60;failed&#x60;. See the [SMS Command Status Values](https://www.twilio.com/docs/wireless/api/smscommand-resource#status-values) for a description of each.
**Direction** | **string** | The direction of the SMS Command. Can be &#x60;to_sim&#x60; or &#x60;from_sim&#x60;. The value of &#x60;to_sim&#x60; is synonymous with the term &#x60;mobile terminated&#x60;, and &#x60;from_sim&#x60; is synonymous with the term &#x60;mobile originated&#x60;.
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**ListSmsCommandResponse**](ListSmsCommandResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

