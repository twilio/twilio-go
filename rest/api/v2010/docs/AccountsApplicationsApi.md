# AccountsApplicationsApi

All URIs are relative to *https://api.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateApplication**](AccountsApplicationsApi.md#CreateApplication) | **Post** /2010-04-01/Accounts/{AccountSid}/Applications.json | 
[**DeleteApplication**](AccountsApplicationsApi.md#DeleteApplication) | **Delete** /2010-04-01/Accounts/{AccountSid}/Applications/{Sid}.json | 
[**FetchApplication**](AccountsApplicationsApi.md#FetchApplication) | **Get** /2010-04-01/Accounts/{AccountSid}/Applications/{Sid}.json | 
[**ListApplication**](AccountsApplicationsApi.md#ListApplication) | **Get** /2010-04-01/Accounts/{AccountSid}/Applications.json | 
[**UpdateApplication**](AccountsApplicationsApi.md#UpdateApplication) | **Post** /2010-04-01/Accounts/{AccountSid}/Applications/{Sid}.json | 



## CreateApplication

> ApiV2010Application CreateApplication(ctx, optional)



Create a new application within your account

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateApplicationParams struct


Name | Type | Description
------------- | ------------- | -------------
**PathAccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that will create the resource.
**ApiVersion** | **string** | The API version to use to start a new TwiML session. Can be: `2010-04-01` or `2008-08-01`. The default value is the account's default API version.
**VoiceUrl** | **string** | The URL we should call when the phone number assigned to this application receives a call.
**VoiceMethod** | **string** | The HTTP method we should use to call `voice_url`. Can be: `GET` or `POST`.
**VoiceFallbackUrl** | **string** | The URL that we should call when an error occurs retrieving or executing the TwiML requested by `url`.
**VoiceFallbackMethod** | **string** | The HTTP method we should use to call `voice_fallback_url`. Can be: `GET` or `POST`.
**StatusCallback** | **string** | The URL we should call using the `status_callback_method` to send status information to your application.
**StatusCallbackMethod** | **string** | The HTTP method we should use to call `status_callback`. Can be: `GET` or `POST`.
**VoiceCallerIdLookup** | **bool** | Whether we should look up the caller's caller-ID name from the CNAM database (additional charges apply). Can be: `true` or `false`.
**SmsUrl** | **string** | The URL we should call when the phone number receives an incoming SMS message.
**SmsMethod** | **string** | The HTTP method we should use to call `sms_url`. Can be: `GET` or `POST`.
**SmsFallbackUrl** | **string** | The URL that we should call when an error occurs while retrieving or executing the TwiML from `sms_url`.
**SmsFallbackMethod** | **string** | The HTTP method we should use to call `sms_fallback_url`. Can be: `GET` or `POST`.
**SmsStatusCallback** | **string** | The URL we should call using a POST method to send status information about SMS messages sent by the application.
**MessageStatusCallback** | **string** | The URL we should call using a POST method to send message status information to your application.
**FriendlyName** | **string** | A descriptive string that you create to describe the new application. It can be up to 64 characters long.
**PublicApplicationConnectEnabled** | **bool** | Whether to allow other Twilio accounts to dial this applicaton using Dial verb. Can be: `true` or `false`.

### Return type

[**ApiV2010Application**](ApiV2010Application.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteApplication

> DeleteApplication(ctx, Sidoptional)



Delete the application by the specified application sid

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the Application resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeleteApplicationParams struct


Name | Type | Description
------------- | ------------- | -------------
**PathAccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Application resources to delete.

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


## FetchApplication

> ApiV2010Application FetchApplication(ctx, Sidoptional)



Fetch the application specified by the provided sid

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the Application resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchApplicationParams struct


Name | Type | Description
------------- | ------------- | -------------
**PathAccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Application resource to fetch.

### Return type

[**ApiV2010Application**](ApiV2010Application.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListApplication

> []ApiV2010Application ListApplication(ctx, optional)



Retrieve a list of applications representing an application within the requesting account

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListApplicationParams struct


Name | Type | Description
------------- | ------------- | -------------
**PathAccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Application resources to read.
**FriendlyName** | **string** | The string that identifies the Application resources to read.
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]ApiV2010Application**](ApiV2010Application.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateApplication

> ApiV2010Application UpdateApplication(ctx, Sidoptional)



Updates the application's properties

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the Application resource to update.

### Other Parameters

Other parameters are passed through a pointer to a UpdateApplicationParams struct


Name | Type | Description
------------- | ------------- | -------------
**PathAccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Application resources to update.
**FriendlyName** | **string** | A descriptive string that you create to describe the resource. It can be up to 64 characters long.
**ApiVersion** | **string** | The API version to use to start a new TwiML session. Can be: `2010-04-01` or `2008-08-01`. The default value is your account's default API version.
**VoiceUrl** | **string** | The URL we should call when the phone number assigned to this application receives a call.
**VoiceMethod** | **string** | The HTTP method we should use to call `voice_url`. Can be: `GET` or `POST`.
**VoiceFallbackUrl** | **string** | The URL that we should call when an error occurs retrieving or executing the TwiML requested by `url`.
**VoiceFallbackMethod** | **string** | The HTTP method we should use to call `voice_fallback_url`. Can be: `GET` or `POST`.
**StatusCallback** | **string** | The URL we should call using the `status_callback_method` to send status information to your application.
**StatusCallbackMethod** | **string** | The HTTP method we should use to call `status_callback`. Can be: `GET` or `POST`.
**VoiceCallerIdLookup** | **bool** | Whether we should look up the caller's caller-ID name from the CNAM database (additional charges apply). Can be: `true` or `false`.
**SmsUrl** | **string** | The URL we should call when the phone number receives an incoming SMS message.
**SmsMethod** | **string** | The HTTP method we should use to call `sms_url`. Can be: `GET` or `POST`.
**SmsFallbackUrl** | **string** | The URL that we should call when an error occurs while retrieving or executing the TwiML from `sms_url`.
**SmsFallbackMethod** | **string** | The HTTP method we should use to call `sms_fallback_url`. Can be: `GET` or `POST`.
**SmsStatusCallback** | **string** | Same as message_status_callback: The URL we should call using a POST method to send status information about SMS messages sent by the application. Deprecated, included for backwards compatibility.
**MessageStatusCallback** | **string** | The URL we should call using a POST method to send message status information to your application.
**PublicApplicationConnectEnabled** | **bool** | Whether to allow other Twilio accounts to dial this applicaton using Dial verb. Can be: `true` or `false`.

### Return type

[**ApiV2010Application**](ApiV2010Application.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

