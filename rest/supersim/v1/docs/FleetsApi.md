# FleetsApi

All URIs are relative to *https://supersim.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateFleet**](FleetsApi.md#CreateFleet) | **Post** /v1/Fleets | 
[**FetchFleet**](FleetsApi.md#FetchFleet) | **Get** /v1/Fleets/{Sid} | 
[**ListFleet**](FleetsApi.md#ListFleet) | **Get** /v1/Fleets | 
[**UpdateFleet**](FleetsApi.md#UpdateFleet) | **Post** /v1/Fleets/{Sid} | 



## CreateFleet

> SupersimV1Fleet CreateFleet(ctx, optional)



Create a Fleet

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateFleetParams struct


Name | Type | Description
------------- | ------------- | -------------
**NetworkAccessProfile** | **string** | The SID or unique name of the Network Access Profile that will control which cellular networks the Fleet's SIMs can connect to.
**UniqueName** | **string** | An application-defined string that uniquely identifies the resource. It can be used in place of the resource's `sid` in the URL to address the resource.
**DataEnabled** | **bool** | Defines whether SIMs in the Fleet are capable of using 2G/3G/4G/LTE/CAT-M data connectivity. Defaults to `true`.
**DataLimit** | **int** | The total data usage (download and upload combined) in Megabytes that each Super SIM assigned to the Fleet can consume during a billing period (normally one month). Value must be between 1MB (1) and 2TB (2,000,000). Defaults to 1GB (1,000).
**IpCommandsUrl** | **string** | The URL that will receive a webhook when a Super SIM in the Fleet is used to send an IP Command from your device to a special IP address. Your server should respond with an HTTP status code in the 200 range; any response body will be ignored.
**IpCommandsMethod** | **string** | A string representing the HTTP method to use when making a request to `ip_commands_url`. Can be one of `POST` or `GET`. Defaults to `POST`.
**SmsCommandsEnabled** | **bool** | Defines whether SIMs in the Fleet are capable of sending and receiving machine-to-machine SMS via Commands. Defaults to `true`.
**SmsCommandsUrl** | **string** | The URL that will receive a webhook when a Super SIM in the Fleet is used to send an SMS from your device to the SMS Commands number. Your server should respond with an HTTP status code in the 200 range; any response body will be ignored.
**SmsCommandsMethod** | **string** | A string representing the HTTP method to use when making a request to `sms_commands_url`. Can be one of `POST` or `GET`. Defaults to `POST`.

### Return type

[**SupersimV1Fleet**](SupersimV1Fleet.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchFleet

> SupersimV1Fleet FetchFleet(ctx, Sid)



Fetch a Fleet instance from your account.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The SID of the Fleet resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchFleetParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**SupersimV1Fleet**](SupersimV1Fleet.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListFleet

> []SupersimV1Fleet ListFleet(ctx, optional)



Retrieve a list of Fleets from your account.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListFleetParams struct


Name | Type | Description
------------- | ------------- | -------------
**NetworkAccessProfile** | **string** | The SID or unique name of the Network Access Profile that controls which cellular networks the Fleet's SIMs can connect to.
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]SupersimV1Fleet**](SupersimV1Fleet.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateFleet

> SupersimV1Fleet UpdateFleet(ctx, Sidoptional)



Updates the given properties of a Super SIM Fleet instance from your account.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The SID of the Fleet resource to update.

### Other Parameters

Other parameters are passed through a pointer to a UpdateFleetParams struct


Name | Type | Description
------------- | ------------- | -------------
**UniqueName** | **string** | An application-defined string that uniquely identifies the resource. It can be used in place of the resource's `sid` in the URL to address the resource.
**NetworkAccessProfile** | **string** | The SID or unique name of the Network Access Profile that will control which cellular networks the Fleet's SIMs can connect to.
**IpCommandsUrl** | **string** | The URL that will receive a webhook when a Super SIM in the Fleet is used to send an IP Command from your device to a special IP address. Your server should respond with an HTTP status code in the 200 range; any response body will be ignored.
**IpCommandsMethod** | **string** | A string representing the HTTP method to use when making a request to `ip_commands_url`. Can be one of `POST` or `GET`. Defaults to `POST`.
**SmsCommandsUrl** | **string** | The URL that will receive a webhook when a Super SIM in the Fleet is used to send an SMS from your device to the SMS Commands number. Your server should respond with an HTTP status code in the 200 range; any response body will be ignored.
**SmsCommandsMethod** | **string** | A string representing the HTTP method to use when making a request to `sms_commands_url`. Can be one of `POST` or `GET`. Defaults to `POST`.
**DataLimit** | **int** | The total data usage (download and upload combined) in Megabytes that each Super SIM assigned to the Fleet can consume during a billing period (normally one month). Value must be between 1MB (1) and 2TB (2,000,000). Defaults to 1GB (1,000).

### Return type

[**SupersimV1Fleet**](SupersimV1Fleet.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

