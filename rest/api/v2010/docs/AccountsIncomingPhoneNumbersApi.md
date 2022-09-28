# AccountsIncomingPhoneNumbersApi

All URIs are relative to *https://api.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateIncomingPhoneNumber**](AccountsIncomingPhoneNumbersApi.md#CreateIncomingPhoneNumber) | **Post** /2010-04-01/Accounts/{AccountSid}/IncomingPhoneNumbers.json | 
[**DeleteIncomingPhoneNumber**](AccountsIncomingPhoneNumbersApi.md#DeleteIncomingPhoneNumber) | **Delete** /2010-04-01/Accounts/{AccountSid}/IncomingPhoneNumbers/{Sid}.json | 
[**FetchIncomingPhoneNumber**](AccountsIncomingPhoneNumbersApi.md#FetchIncomingPhoneNumber) | **Get** /2010-04-01/Accounts/{AccountSid}/IncomingPhoneNumbers/{Sid}.json | 
[**ListIncomingPhoneNumber**](AccountsIncomingPhoneNumbersApi.md#ListIncomingPhoneNumber) | **Get** /2010-04-01/Accounts/{AccountSid}/IncomingPhoneNumbers.json | 
[**UpdateIncomingPhoneNumber**](AccountsIncomingPhoneNumbersApi.md#UpdateIncomingPhoneNumber) | **Post** /2010-04-01/Accounts/{AccountSid}/IncomingPhoneNumbers/{Sid}.json | 



## CreateIncomingPhoneNumber

> ApiV2010IncomingPhoneNumber CreateIncomingPhoneNumber(ctx, optional)



Purchase a phone-number for the account.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateIncomingPhoneNumberParams struct


Name | Type | Description
------------- | ------------- | -------------
**PathAccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that will create the resource.
**ApiVersion** | **string** | The API version to use for incoming calls made to the new phone number. The default is &#x60;2010-04-01&#x60;.
**FriendlyName** | **string** | A descriptive string that you created to describe the new phone number. It can be up to 64 characters long. By default, this is a formatted version of the new phone number.
**SmsApplicationSid** | **string** | The SID of the application that should handle SMS messages sent to the new phone number. If an &#x60;sms_application_sid&#x60; is present, we ignore all of the &#x60;sms_*_url&#x60; urls and use those set on the application.
**SmsFallbackMethod** | **string** | The HTTP method that we should use to call &#x60;sms_fallback_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60; and defaults to &#x60;POST&#x60;.
**SmsFallbackUrl** | **string** | The URL that we should call when an error occurs while requesting or executing the TwiML defined by &#x60;sms_url&#x60;.
**SmsMethod** | **string** | The HTTP method that we should use to call &#x60;sms_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60; and defaults to &#x60;POST&#x60;.
**SmsUrl** | **string** | The URL we should call when the new phone number receives an incoming SMS message.
**StatusCallback** | **string** | The URL we should call using the &#x60;status_callback_method&#x60; to send status information to your application.
**StatusCallbackMethod** | **string** | The HTTP method we should use to call &#x60;status_callback&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60; and defaults to &#x60;POST&#x60;.
**VoiceApplicationSid** | **string** | The SID of the application we should use to handle calls to the new phone number. If a &#x60;voice_application_sid&#x60; is present, we ignore all of the voice urls and use only those set on the application. Setting a &#x60;voice_application_sid&#x60; will automatically delete your &#x60;trunk_sid&#x60; and vice versa.
**VoiceCallerIdLookup** | **bool** | Whether to lookup the caller&#39;s name from the CNAM database and post it to your app. Can be: &#x60;true&#x60; or &#x60;false&#x60; and defaults to &#x60;false&#x60;.
**VoiceFallbackMethod** | **string** | The HTTP method that we should use to call &#x60;voice_fallback_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60; and defaults to &#x60;POST&#x60;.
**VoiceFallbackUrl** | **string** | The URL that we should call when an error occurs retrieving or executing the TwiML requested by &#x60;url&#x60;.
**VoiceMethod** | **string** | The HTTP method that we should use to call &#x60;voice_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60; and defaults to &#x60;POST&#x60;.
**VoiceUrl** | **string** | The URL that we should call to answer a call to the new phone number. The &#x60;voice_url&#x60; will not be called if a &#x60;voice_application_sid&#x60; or a &#x60;trunk_sid&#x60; is set.
**EmergencyStatus** | **string** | 
**EmergencyAddressSid** | **string** | The SID of the emergency address configuration to use for emergency calling from the new phone number.
**TrunkSid** | **string** | The SID of the Trunk we should use to handle calls to the new phone number. If a &#x60;trunk_sid&#x60; is present, we ignore all of the voice urls and voice applications and use only those set on the Trunk. Setting a &#x60;trunk_sid&#x60; will automatically delete your &#x60;voice_application_sid&#x60; and vice versa.
**IdentitySid** | **string** | The SID of the Identity resource that we should associate with the new phone number. Some regions require an identity to meet local regulations.
**AddressSid** | **string** | The SID of the Address resource we should associate with the new phone number. Some regions require addresses to meet local regulations.
**VoiceReceiveMode** | **string** | 
**BundleSid** | **string** | The SID of the Bundle resource that you associate with the phone number. Some regions require a Bundle to meet local Regulations.
**PhoneNumber** | **string** | The phone number to purchase specified in [E.164](https://www.twilio.com/docs/glossary/what-e164) format.  E.164 phone numbers consist of a + followed by the country code and subscriber number without punctuation characters. For example, +14155551234.
**AreaCode** | **string** | The desired area code for your new incoming phone number. Can be any three-digit, US or Canada area code. We will provision an available phone number within this area code for you. **You must provide an &#x60;area_code&#x60; or a &#x60;phone_number&#x60;.** (US and Canada only).

### Return type

[**ApiV2010IncomingPhoneNumber**](ApiV2010IncomingPhoneNumber.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteIncomingPhoneNumber

> DeleteIncomingPhoneNumber(ctx, Sidoptional)



Delete a phone-numbers belonging to the account used to make the request.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the IncomingPhoneNumber resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeleteIncomingPhoneNumberParams struct


Name | Type | Description
------------- | ------------- | -------------
**PathAccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the IncomingPhoneNumber resources to delete.

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


## FetchIncomingPhoneNumber

> ApiV2010IncomingPhoneNumber FetchIncomingPhoneNumber(ctx, Sidoptional)



Fetch an incoming-phone-number belonging to the account used to make the request.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the IncomingPhoneNumber resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchIncomingPhoneNumberParams struct


Name | Type | Description
------------- | ------------- | -------------
**PathAccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the IncomingPhoneNumber resource to fetch.

### Return type

[**ApiV2010IncomingPhoneNumber**](ApiV2010IncomingPhoneNumber.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListIncomingPhoneNumber

> []ApiV2010IncomingPhoneNumber ListIncomingPhoneNumber(ctx, optional)



Retrieve a list of incoming-phone-numbers belonging to the account used to make the request.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListIncomingPhoneNumberParams struct


Name | Type | Description
------------- | ------------- | -------------
**PathAccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the IncomingPhoneNumber resources to read.
**Beta** | **bool** | Whether to include phone numbers new to the Twilio platform. Can be: &#x60;true&#x60; or &#x60;false&#x60; and the default is &#x60;true&#x60;.
**FriendlyName** | **string** | A string that identifies the IncomingPhoneNumber resources to read.
**PhoneNumber** | **string** | The phone numbers of the IncomingPhoneNumber resources to read. You can specify partial numbers and use &#39;*&#39; as a wildcard for any digit.
**Origin** | **string** | Whether to include phone numbers based on their origin. Can be: &#x60;twilio&#x60; or &#x60;hosted&#x60;. By default, phone numbers of all origin are included.
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]ApiV2010IncomingPhoneNumber**](ApiV2010IncomingPhoneNumber.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateIncomingPhoneNumber

> ApiV2010IncomingPhoneNumber UpdateIncomingPhoneNumber(ctx, Sidoptional)



Update an incoming-phone-number instance.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the IncomingPhoneNumber resource to update.

### Other Parameters

Other parameters are passed through a pointer to a UpdateIncomingPhoneNumberParams struct


Name | Type | Description
------------- | ------------- | -------------
**PathAccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the IncomingPhoneNumber resource to update.  For more information, see [Exchanging Numbers Between Subaccounts](https://www.twilio.com/docs/iam/api/subaccounts#exchanging-numbers).
**AccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the IncomingPhoneNumber resource to update.  For more information, see [Exchanging Numbers Between Subaccounts](https://www.twilio.com/docs/iam/api/subaccounts#exchanging-numbers).
**ApiVersion** | **string** | The API version to use for incoming calls made to the phone number. The default is &#x60;2010-04-01&#x60;.
**FriendlyName** | **string** | A descriptive string that you created to describe this phone number. It can be up to 64 characters long. By default, this is a formatted version of the phone number.
**SmsApplicationSid** | **string** | The SID of the application that should handle SMS messages sent to the number. If an &#x60;sms_application_sid&#x60; is present, we ignore all of the &#x60;sms_*_url&#x60; urls and use those set on the application.
**SmsFallbackMethod** | **string** | The HTTP method that we should use to call &#x60;sms_fallback_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60; and defaults to &#x60;POST&#x60;.
**SmsFallbackUrl** | **string** | The URL that we should call when an error occurs while requesting or executing the TwiML defined by &#x60;sms_url&#x60;.
**SmsMethod** | **string** | The HTTP method that we should use to call &#x60;sms_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60; and defaults to &#x60;POST&#x60;.
**SmsUrl** | **string** | The URL we should call when the phone number receives an incoming SMS message.
**StatusCallback** | **string** | The URL we should call using the &#x60;status_callback_method&#x60; to send status information to your application.
**StatusCallbackMethod** | **string** | The HTTP method we should use to call &#x60;status_callback&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60; and defaults to &#x60;POST&#x60;.
**VoiceApplicationSid** | **string** | The SID of the application we should use to handle phone calls to the phone number. If a &#x60;voice_application_sid&#x60; is present, we ignore all of the voice urls and use only those set on the application. Setting a &#x60;voice_application_sid&#x60; will automatically delete your &#x60;trunk_sid&#x60; and vice versa.
**VoiceCallerIdLookup** | **bool** | Whether to lookup the caller&#39;s name from the CNAM database and post it to your app. Can be: &#x60;true&#x60; or &#x60;false&#x60; and defaults to &#x60;false&#x60;.
**VoiceFallbackMethod** | **string** | The HTTP method that we should use to call &#x60;voice_fallback_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60; and defaults to &#x60;POST&#x60;.
**VoiceFallbackUrl** | **string** | The URL that we should call when an error occurs retrieving or executing the TwiML requested by &#x60;url&#x60;.
**VoiceMethod** | **string** | The HTTP method that we should use to call &#x60;voice_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60; and defaults to &#x60;POST&#x60;.
**VoiceUrl** | **string** | The URL that we should call to answer a call to the phone number. The &#x60;voice_url&#x60; will not be called if a &#x60;voice_application_sid&#x60; or a &#x60;trunk_sid&#x60; is set.
**EmergencyStatus** | **string** | 
**EmergencyAddressSid** | **string** | The SID of the emergency address configuration to use for emergency calling from this phone number.
**TrunkSid** | **string** | The SID of the Trunk we should use to handle phone calls to the phone number. If a &#x60;trunk_sid&#x60; is present, we ignore all of the voice urls and voice applications and use only those set on the Trunk. Setting a &#x60;trunk_sid&#x60; will automatically delete your &#x60;voice_application_sid&#x60; and vice versa.
**VoiceReceiveMode** | **string** | 
**IdentitySid** | **string** | The SID of the Identity resource that we should associate with the phone number. Some regions require an identity to meet local regulations.
**AddressSid** | **string** | The SID of the Address resource we should associate with the phone number. Some regions require addresses to meet local regulations.
**BundleSid** | **string** | The SID of the Bundle resource that you associate with the phone number. Some regions require a Bundle to meet local Regulations.

### Return type

[**ApiV2010IncomingPhoneNumber**](ApiV2010IncomingPhoneNumber.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

