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
**ShortCode** | **string** | Only show the ShortCode resources that match this pattern. You can specify partial numbers and use &#39;*&#39; as a wildcard for any digit.
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
**FriendlyName** | **string** | A descriptive string that you created to describe this resource. It can be up to 64 characters long. By default, the &#x60;FriendlyName&#x60; is the short code.
**ApiVersion** | **string** | The API version to use to start a new TwiML session. Can be: &#x60;2010-04-01&#x60; or &#x60;2008-08-01&#x60;.
**SmsUrl** | **string** | The URL we should call when receiving an incoming SMS message to this short code.
**SmsMethod** | **string** | The HTTP method we should use when calling the &#x60;sms_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60;.
**SmsFallbackUrl** | **string** | The URL that we should call if an error occurs while retrieving or executing the TwiML from &#x60;sms_url&#x60;.
**SmsFallbackMethod** | **string** | The HTTP method that we should use to call the &#x60;sms_fallback_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60;.

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

