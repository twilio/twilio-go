# AccountsUsageTriggersApi

All URIs are relative to *https://api.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateUsageTrigger**](AccountsUsageTriggersApi.md#CreateUsageTrigger) | **Post** /2010-04-01/Accounts/{AccountSid}/Usage/Triggers.json | 
[**DeleteUsageTrigger**](AccountsUsageTriggersApi.md#DeleteUsageTrigger) | **Delete** /2010-04-01/Accounts/{AccountSid}/Usage/Triggers/{Sid}.json | 
[**FetchUsageTrigger**](AccountsUsageTriggersApi.md#FetchUsageTrigger) | **Get** /2010-04-01/Accounts/{AccountSid}/Usage/Triggers/{Sid}.json | 
[**ListUsageTrigger**](AccountsUsageTriggersApi.md#ListUsageTrigger) | **Get** /2010-04-01/Accounts/{AccountSid}/Usage/Triggers.json | 
[**UpdateUsageTrigger**](AccountsUsageTriggersApi.md#UpdateUsageTrigger) | **Post** /2010-04-01/Accounts/{AccountSid}/Usage/Triggers/{Sid}.json | 



## CreateUsageTrigger

> ApiV2010UsageTrigger CreateUsageTrigger(ctx, optional)



Create a new UsageTrigger

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateUsageTriggerParams struct


Name | Type | Description
------------- | ------------- | -------------
**PathAccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that will create the resource.
**CallbackUrl** | **string** | The URL we should call using `callback_method` when the trigger fires.
**TriggerValue** | **string** | The usage value at which the trigger should fire.  For convenience, you can use an offset value such as `+30` to specify a trigger_value that is 30 units more than the current usage value. Be sure to urlencode a `+` as `%2B`.
**UsageCategory** | **string** | 
**CallbackMethod** | **string** | The HTTP method we should use to call `callback_url`. Can be: `GET` or `POST` and the default is `POST`.
**FriendlyName** | **string** | A descriptive string that you create to describe the resource. It can be up to 64 characters long.
**Recurring** | **string** | 
**TriggerBy** | **string** | 

### Return type

[**ApiV2010UsageTrigger**](ApiV2010UsageTrigger.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteUsageTrigger

> DeleteUsageTrigger(ctx, Sidoptional)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the UsageTrigger resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeleteUsageTriggerParams struct


Name | Type | Description
------------- | ------------- | -------------
**PathAccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the UsageTrigger resources to delete.

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


## FetchUsageTrigger

> ApiV2010UsageTrigger FetchUsageTrigger(ctx, Sidoptional)



Fetch and instance of a usage-trigger

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the UsageTrigger resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchUsageTriggerParams struct


Name | Type | Description
------------- | ------------- | -------------
**PathAccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the UsageTrigger resource to fetch.

### Return type

[**ApiV2010UsageTrigger**](ApiV2010UsageTrigger.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUsageTrigger

> []ApiV2010UsageTrigger ListUsageTrigger(ctx, optional)



Retrieve a list of usage-triggers belonging to the account used to make the request

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListUsageTriggerParams struct


Name | Type | Description
------------- | ------------- | -------------
**PathAccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the UsageTrigger resources to read.
**Recurring** | **string** | The frequency of recurring UsageTriggers to read. Can be: `daily`, `monthly`, or `yearly` to read recurring UsageTriggers. An empty value or a value of `alltime` reads non-recurring UsageTriggers.
**TriggerBy** | **string** | The trigger field of the UsageTriggers to read.  Can be: `count`, `usage`, or `price` as described in the [UsageRecords documentation](https://www.twilio.com/docs/usage/api/usage-record#usage-count-price).
**UsageCategory** | **string** | The usage category of the UsageTriggers to read. Must be a supported [usage categories](https://www.twilio.com/docs/usage/api/usage-record#usage-categories).
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]ApiV2010UsageTrigger**](ApiV2010UsageTrigger.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUsageTrigger

> ApiV2010UsageTrigger UpdateUsageTrigger(ctx, Sidoptional)



Update an instance of a usage trigger

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the UsageTrigger resource to update.

### Other Parameters

Other parameters are passed through a pointer to a UpdateUsageTriggerParams struct


Name | Type | Description
------------- | ------------- | -------------
**PathAccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the UsageTrigger resources to update.
**CallbackMethod** | **string** | The HTTP method we should use to call `callback_url`. Can be: `GET` or `POST` and the default is `POST`.
**CallbackUrl** | **string** | The URL we should call using `callback_method` when the trigger fires.
**FriendlyName** | **string** | A descriptive string that you create to describe the resource. It can be up to 64 characters long.

### Return type

[**ApiV2010UsageTrigger**](ApiV2010UsageTrigger.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

