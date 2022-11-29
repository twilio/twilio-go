# AccountsIncomingPhoneNumbersMobileApi

All URIs are relative to *https://api.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateIncomingPhoneNumberMobile**](AccountsIncomingPhoneNumbersMobileApi.md#CreateIncomingPhoneNumberMobile) | **Post** /2010-04-01/Accounts/{AccountSid}/IncomingPhoneNumbers/Mobile.json | 
[**ListIncomingPhoneNumberMobile**](AccountsIncomingPhoneNumbersMobileApi.md#ListIncomingPhoneNumberMobile) | **Get** /2010-04-01/Accounts/{AccountSid}/IncomingPhoneNumbers/Mobile.json | 



## CreateIncomingPhoneNumberMobile

> ApiV2010IncomingPhoneNumberMobile CreateIncomingPhoneNumberMobile(ctx, optional)





### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateIncomingPhoneNumberMobileParams struct


Name | Type | Description
------------- | ------------- | -------------
**PathAccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that will create the resource.
**PhoneNumber** | **string** | The phone number to purchase specified in [E.164](https://www.twilio.com/docs/glossary/what-e164) format.  E.164 phone numbers consist of a + followed by the country code and subscriber number without punctuation characters. For example, +14155551234.
**ApiVersion** | **string** | The API version to use for incoming calls made to the new phone number. The default is `2010-04-01`.
**FriendlyName** | **string** | A descriptive string that you created to describe the new phone number. It can be up to 64 characters long. By default, the is a formatted version of the phone number.
**SmsApplicationSid** | **string** | The SID of the application that should handle SMS messages sent to the new phone number. If an `sms_application_sid` is present, we ignore all of the `sms_*_url` urls and use those of the application.
**SmsFallbackMethod** | **string** | The HTTP method that we should use to call `sms_fallback_url`. Can be: `GET` or `POST` and defaults to `POST`.
**SmsFallbackUrl** | **string** | The URL that we should call when an error occurs while requesting or executing the TwiML defined by `sms_url`.
**SmsMethod** | **string** | The HTTP method that we should use to call `sms_url`. Can be: `GET` or `POST` and defaults to `POST`.
**SmsUrl** | **string** | The URL we should call when the new phone number receives an incoming SMS message.
**StatusCallback** | **string** | The URL we should call using the `status_callback_method` to send status information to your application.
**StatusCallbackMethod** | **string** | The HTTP method we should use to call `status_callback`. Can be: `GET` or `POST` and defaults to `POST`.
**VoiceApplicationSid** | **string** | The SID of the application we should use to handle calls to the new phone number. If a `voice_application_sid` is present, we ignore all of the voice urls and use only those set on the application. Setting a `voice_application_sid` will automatically delete your `trunk_sid` and vice versa.
**VoiceCallerIdLookup** | **bool** | Whether to lookup the caller's name from the CNAM database and post it to your app. Can be: `true` or `false` and defaults to `false`.
**VoiceFallbackMethod** | **string** | The HTTP method that we should use to call `voice_fallback_url`. Can be: `GET` or `POST` and defaults to `POST`.
**VoiceFallbackUrl** | **string** | The URL that we should call when an error occurs retrieving or executing the TwiML requested by `url`.
**VoiceMethod** | **string** | The HTTP method that we should use to call `voice_url`. Can be: `GET` or `POST` and defaults to `POST`.
**VoiceUrl** | **string** | The URL that we should call to answer a call to the new phone number. The `voice_url` will not be called if a `voice_application_sid` or a `trunk_sid` is set.
**IdentitySid** | **string** | The SID of the Identity resource that we should associate with the new phone number. Some regions require an identity to meet local regulations.
**AddressSid** | **string** | The SID of the Address resource we should associate with the new phone number. Some regions require addresses to meet local regulations.
**EmergencyStatus** | **string** | 
**EmergencyAddressSid** | **string** | The SID of the emergency address configuration to use for emergency calling from the new phone number.
**TrunkSid** | **string** | The SID of the Trunk we should use to handle calls to the new phone number. If a `trunk_sid` is present, we ignore all of the voice urls and voice applications and use only those set on the Trunk. Setting a `trunk_sid` will automatically delete your `voice_application_sid` and vice versa.
**VoiceReceiveMode** | **string** | 
**BundleSid** | **string** | The SID of the Bundle resource that you associate with the phone number. Some regions require a Bundle to meet local Regulations.

### Return type

[**ApiV2010IncomingPhoneNumberMobile**](ApiV2010IncomingPhoneNumberMobile.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListIncomingPhoneNumberMobile

> []ApiV2010IncomingPhoneNumberMobile ListIncomingPhoneNumberMobile(ctx, optional)





### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListIncomingPhoneNumberMobileParams struct


Name | Type | Description
------------- | ------------- | -------------
**PathAccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the resources to read.
**Beta** | **bool** | Whether to include phone numbers new to the Twilio platform. Can be: `true` or `false` and the default is `true`.
**FriendlyName** | **string** | A string that identifies the resources to read.
**PhoneNumber** | **string** | The phone numbers of the IncomingPhoneNumber resources to read. You can specify partial numbers and use '*' as a wildcard for any digit.
**Origin** | **string** | Whether to include phone numbers based on their origin. Can be: `twilio` or `hosted`. By default, phone numbers of all origin are included.
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]ApiV2010IncomingPhoneNumberMobile**](ApiV2010IncomingPhoneNumberMobile.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

