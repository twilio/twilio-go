# SimsApi

All URIs are relative to *https://wireless.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteSim**](SimsApi.md#DeleteSim) | **Delete** /v1/Sims/{Sid} | 
[**FetchSim**](SimsApi.md#FetchSim) | **Get** /v1/Sims/{Sid} | 
[**ListSim**](SimsApi.md#ListSim) | **Get** /v1/Sims | 
[**UpdateSim**](SimsApi.md#UpdateSim) | **Post** /v1/Sims/{Sid} | 



## DeleteSim

> DeleteSim(ctx, Sid)



Delete a Sim resource on your Account.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The SID or the `unique_name` of the Sim resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeleteSimParams struct


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


## FetchSim

> WirelessV1Sim FetchSim(ctx, Sid)



Fetch a Sim resource on your Account.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The SID or the `unique_name` of the Sim resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchSimParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**WirelessV1Sim**](WirelessV1Sim.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSim

> []WirelessV1Sim ListSim(ctx, optional)



Retrieve a list of Sim resources on your Account.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListSimParams struct


Name | Type | Description
------------- | ------------- | -------------
**Status** | **string** | Only return Sim resources with this status.
**Iccid** | **string** | Only return Sim resources with this ICCID. This will return a list with a maximum size of 1.
**RatePlan** | **string** | The SID or unique name of a [RatePlan resource](https://www.twilio.com/docs/wireless/api/rateplan-resource). Only return Sim resources assigned to this RatePlan resource.
**EId** | **string** | Deprecated.
**SimRegistrationCode** | **string** | Only return Sim resources with this registration code. This will return a list with a maximum size of 1.
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]WirelessV1Sim**](WirelessV1Sim.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSim

> WirelessV1Sim UpdateSim(ctx, Sidoptional)



Updates the given properties of a Sim resource on your Account.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The SID or the `unique_name` of the Sim resource to update.

### Other Parameters

Other parameters are passed through a pointer to a UpdateSimParams struct


Name | Type | Description
------------- | ------------- | -------------
**UniqueName** | **string** | An application-defined string that uniquely identifies the resource. It can be used in place of the `sid` in the URL path to address the resource.
**CallbackMethod** | **string** | The HTTP method we should use to call `callback_url`. Can be: `POST` or `GET`. The default is `POST`.
**CallbackUrl** | **string** | The URL we should call using the `callback_url` when the SIM has finished updating. When the SIM transitions from `new` to `ready` or from any status to `deactivated`, we call this URL when the status changes to an intermediate status (`ready` or `deactivated`) and again when the status changes to its final status (`active` or `canceled`).
**FriendlyName** | **string** | A descriptive string that you create to describe the Sim resource. It does not need to be unique.
**RatePlan** | **string** | The SID or unique name of the [RatePlan resource](https://www.twilio.com/docs/wireless/api/rateplan-resource) to which the Sim resource should be assigned.
**Status** | **string** | 
**CommandsCallbackMethod** | **string** | The HTTP method we should use to call `commands_callback_url`. Can be: `POST` or `GET`. The default is `POST`.
**CommandsCallbackUrl** | **string** | The URL we should call using the `commands_callback_method` when the SIM sends a [Command](https://www.twilio.com/docs/wireless/api/command-resource). Your server should respond with an HTTP status code in the 200 range; any response body is ignored.
**SmsFallbackMethod** | **string** | The HTTP method we should use to call `sms_fallback_url`. Can be: `GET` or `POST`. Default is `POST`.
**SmsFallbackUrl** | **string** | The URL we should call using the `sms_fallback_method` when an error occurs while retrieving or executing the TwiML requested from `sms_url`.
**SmsMethod** | **string** | The HTTP method we should use to call `sms_url`. Can be: `GET` or `POST`. Default is `POST`.
**SmsUrl** | **string** | The URL we should call using the `sms_method` when the SIM-connected device sends an SMS message that is not a [Command](https://www.twilio.com/docs/wireless/api/command-resource).
**VoiceFallbackMethod** | **string** | Deprecated.
**VoiceFallbackUrl** | **string** | Deprecated.
**VoiceMethod** | **string** | Deprecated.
**VoiceUrl** | **string** | Deprecated.
**ResetStatus** | **string** | 
**AccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) to which the Sim resource should belong. The Account SID can only be that of the requesting Account or that of a [Subaccount](https://www.twilio.com/docs/iam/api/subaccounts) of the requesting Account. Only valid when the Sim resource's status is `new`. For more information, see the [Move SIMs between Subaccounts documentation](https://www.twilio.com/docs/wireless/api/sim-resource#move-sims-between-subaccounts).

### Return type

[**WirelessV1Sim**](WirelessV1Sim.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

