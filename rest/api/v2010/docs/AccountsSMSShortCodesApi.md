# AccountsSMSShortCodesApi

All URIs are relative to *https://api.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchShortCode**](AccountsSMSShortCodesApi.md#FetchShortCode) | **Get** /2010-04-01/Accounts/{AccountSid}/SMS/ShortCodes/{Sid}.json | 
[**ListShortCode**](AccountsSMSShortCodesApi.md#ListShortCode) | **Get** /2010-04-01/Accounts/{AccountSid}/SMS/ShortCodes.json | 
[**UpdateShortCode**](AccountsSMSShortCodesApi.md#UpdateShortCode) | **Post** /2010-04-01/Accounts/{AccountSid}/SMS/ShortCodes/{Sid}.json | 



## FetchShortCode

> ApiV2010ShortCode FetchShortCode(ctx, Sidoptional)



Fetch an instance of a short code

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the ShortCode resource to fetch

### Other Parameters

Other parameters are passed through a pointer to a FetchShortCodeParams struct


Name | Type | Description
------------- | ------------- | -------------
**PathAccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the ShortCode resource(s) to fetch.

### Return type

[**ApiV2010ShortCode**](ApiV2010ShortCode.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListShortCode

> []ApiV2010ShortCode ListShortCode(ctx, optional)



Retrieve a list of short-codes belonging to the account used to make the request

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListShortCodeParams struct


Name | Type | Description
------------- | ------------- | -------------
**PathAccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the ShortCode resource(s) to read.
**FriendlyName** | **string** | The string that identifies the ShortCode resources to read.
**ShortCode** | **string** | Only show the ShortCode resources that match this pattern. You can specify partial numbers and use '*' as a wildcard for any digit.
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]ApiV2010ShortCode**](ApiV2010ShortCode.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateShortCode

> ApiV2010ShortCode UpdateShortCode(ctx, Sidoptional)



Update a short code with the following parameters

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the ShortCode resource to update

### Other Parameters

Other parameters are passed through a pointer to a UpdateShortCodeParams struct


Name | Type | Description
------------- | ------------- | -------------
**PathAccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the ShortCode resource(s) to update.
**FriendlyName** | **string** | A descriptive string that you created to describe this resource. It can be up to 64 characters long. By default, the `FriendlyName` is the short code.
**ApiVersion** | **string** | The API version to use to start a new TwiML session. Can be: `2010-04-01` or `2008-08-01`.
**SmsUrl** | **string** | The URL we should call when receiving an incoming SMS message to this short code.
**SmsMethod** | **string** | The HTTP method we should use when calling the `sms_url`. Can be: `GET` or `POST`.
**SmsFallbackUrl** | **string** | The URL that we should call if an error occurs while retrieving or executing the TwiML from `sms_url`.
**SmsFallbackMethod** | **string** | The HTTP method that we should use to call the `sms_fallback_url`. Can be: `GET` or `POST`.

### Return type

[**ApiV2010ShortCode**](ApiV2010ShortCode.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

