# RatePlansApi

All URIs are relative to *https://wireless.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRatePlan**](RatePlansApi.md#CreateRatePlan) | **Post** /v1/RatePlans | 
[**DeleteRatePlan**](RatePlansApi.md#DeleteRatePlan) | **Delete** /v1/RatePlans/{Sid} | 
[**FetchRatePlan**](RatePlansApi.md#FetchRatePlan) | **Get** /v1/RatePlans/{Sid} | 
[**ListRatePlan**](RatePlansApi.md#ListRatePlan) | **Get** /v1/RatePlans | 
[**UpdateRatePlan**](RatePlansApi.md#UpdateRatePlan) | **Post** /v1/RatePlans/{Sid} | 



## CreateRatePlan

> WirelessV1RatePlan CreateRatePlan(ctx, optional)





### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateRatePlanParams struct


Name | Type | Description
------------- | ------------- | -------------
**UniqueName** | **string** | An application-defined string that uniquely identifies the resource. It can be used in place of the resource's `sid` in the URL to address the resource.
**FriendlyName** | **string** | A descriptive string that you create to describe the resource. It does not have to be unique.
**DataEnabled** | **bool** | Whether SIMs can use GPRS/3G/4G/LTE data connectivity.
**DataLimit** | **int** | The total data usage (download and upload combined) in Megabytes that the Network allows during one month on the home network (T-Mobile USA). The metering period begins the day of activation and ends on the same day in the following month. Can be up to 2TB and the default value is `1000`.
**DataMetering** | **string** | The model used to meter data usage. Can be: `payg` and `quota-1`, `quota-10`, and `quota-50`. Learn more about the available [data metering models](https://www.twilio.com/docs/wireless/api/rateplan-resource#payg-vs-quota-data-plans).
**MessagingEnabled** | **bool** | Whether SIMs can make, send, and receive SMS using [Commands](https://www.twilio.com/docs/wireless/api/command-resource).
**VoiceEnabled** | **bool** | Deprecated.
**NationalRoamingEnabled** | **bool** | Whether SIMs can roam on networks other than the home network (T-Mobile USA) in the United States. See [national roaming](https://www.twilio.com/docs/wireless/api/rateplan-resource#national-roaming).
**InternationalRoaming** | **[]string** | The list of services that SIMs capable of using GPRS/3G/4G/LTE data connectivity can use outside of the United States. Can contain: `data` and `messaging`.
**NationalRoamingDataLimit** | **int** | The total data usage (download and upload combined) in Megabytes that the Network allows during one month on non-home networks in the United States. The metering period begins the day of activation and ends on the same day in the following month. Can be up to 2TB. See [national roaming](https://www.twilio.com/docs/wireless/api/rateplan-resource#national-roaming) for more info.
**InternationalRoamingDataLimit** | **int** | The total data usage (download and upload combined) in Megabytes that the Network allows during one month when roaming outside the United States. Can be up to 2TB.

### Return type

[**WirelessV1RatePlan**](WirelessV1RatePlan.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRatePlan

> DeleteRatePlan(ctx, Sid)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The SID of the RatePlan resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeleteRatePlanParams struct


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


## FetchRatePlan

> WirelessV1RatePlan FetchRatePlan(ctx, Sid)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The SID of the RatePlan resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchRatePlanParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**WirelessV1RatePlan**](WirelessV1RatePlan.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRatePlan

> []WirelessV1RatePlan ListRatePlan(ctx, optional)





### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListRatePlanParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]WirelessV1RatePlan**](WirelessV1RatePlan.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRatePlan

> WirelessV1RatePlan UpdateRatePlan(ctx, Sidoptional)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The SID of the RatePlan resource to update.

### Other Parameters

Other parameters are passed through a pointer to a UpdateRatePlanParams struct


Name | Type | Description
------------- | ------------- | -------------
**UniqueName** | **string** | An application-defined string that uniquely identifies the resource. It can be used in place of the resource's `sid` in the URL to address the resource.
**FriendlyName** | **string** | A descriptive string that you create to describe the resource. It does not have to be unique.

### Return type

[**WirelessV1RatePlan**](WirelessV1RatePlan.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

