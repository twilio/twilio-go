# IpCommandsApi

All URIs are relative to *https://supersim.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateIpCommand**](IpCommandsApi.md#CreateIpCommand) | **Post** /v1/IpCommands | 
[**FetchIpCommand**](IpCommandsApi.md#FetchIpCommand) | **Get** /v1/IpCommands/{Sid} | 
[**ListIpCommand**](IpCommandsApi.md#ListIpCommand) | **Get** /v1/IpCommands | 



## CreateIpCommand

> SupersimV1IpCommand CreateIpCommand(ctx, optional)



Send an IP Command to a Super SIM.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateIpCommandParams struct


Name | Type | Description
------------- | ------------- | -------------
**Sim** | **string** | The &#x60;sid&#x60; or &#x60;unique_name&#x60; of the [Super SIM](https://www.twilio.com/docs/iot/supersim/api/sim-resource) to send the IP Command to.
**Payload** | **string** | The payload to be delivered to the device.
**DevicePort** | **int** | The device port to which the IP Command will be sent.
**PayloadType** | **string** | 
**CallbackUrl** | **string** | The URL we should call using the &#x60;callback_method&#x60; after we have sent the IP Command.
**CallbackMethod** | **string** | The HTTP method we should use to call &#x60;callback_url&#x60;. Can be &#x60;GET&#x60; or &#x60;POST&#x60;, and the default is &#x60;POST&#x60;.

### Return type

[**SupersimV1IpCommand**](SupersimV1IpCommand.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchIpCommand

> SupersimV1IpCommand FetchIpCommand(ctx, Sid)



Fetch IP Command instance from your account.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The SID of the IP Command resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchIpCommandParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**SupersimV1IpCommand**](SupersimV1IpCommand.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListIpCommand

> []SupersimV1IpCommand ListIpCommand(ctx, optional)



Retrieve a list of IP Commands from your account.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListIpCommandParams struct


Name | Type | Description
------------- | ------------- | -------------
**Sim** | **string** | The SID or unique name of the Sim resource that IP Command was sent to or from.
**SimIccid** | **string** | The ICCID of the Sim resource that IP Command was sent to or from.
**Status** | **string** | The status of the IP Command. Can be: &#x60;queued&#x60;, &#x60;sent&#x60;, &#x60;received&#x60; or &#x60;failed&#x60;. See the [IP Command Status Values](https://www.twilio.com/docs/wireless/api/ipcommand-resource#status-values) for a description of each.
**Direction** | **string** | The direction of the IP Command. Can be &#x60;to_sim&#x60; or &#x60;from_sim&#x60;. The value of &#x60;to_sim&#x60; is synonymous with the term &#x60;mobile terminated&#x60;, and &#x60;from_sim&#x60; is synonymous with the term &#x60;mobile originated&#x60;.
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]SupersimV1IpCommand**](SupersimV1IpCommand.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

