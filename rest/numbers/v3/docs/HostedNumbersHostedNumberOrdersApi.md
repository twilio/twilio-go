# HostedNumbersHostedNumberOrdersApi

All URIs are relative to *https://numbers.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateHostedNumbersHostedNumberOrder**](HostedNumbersHostedNumberOrdersApi.md#CreateHostedNumbersHostedNumberOrder) | **Post** /v3/HostedNumbers/HostedNumberOrders | 



## CreateHostedNumbersHostedNumberOrder

> NumbersV3HostedNumberOrder CreateHostedNumbersHostedNumberOrder(ctx, optional)



Host a phone number's capability on Twilio's platform.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateHostedNumbersHostedNumberOrderParams struct


Name | Type | Description
------------- | ------------- | -------------
**PhoneNumber** | **string** | The number to host in [+E.164](https://en.wikipedia.org/wiki/E.164) format
**SmsCapability** | **bool** | Used to specify that the SMS capability will be hosted on Twilio's platform.
**AccountSid** | **string** | This defaults to the AccountSid of the authorization the user is using. This can be provided to specify a subaccount to add the HostedNumberOrder to.
**FriendlyName** | **string** | A 64 character string that is a human readable text that describes this resource.
**UniqueName** | **string** | Optional. Provides a unique and addressable name to be assigned to this HostedNumberOrder, assigned by the developer, to be optionally used in addition to SID.
**CcEmails** | **[]string** | Optional. A list of emails that the LOA document for this HostedNumberOrder will be carbon copied to.
**SmsUrl** | **string** | The URL that Twilio should request when somebody sends an SMS to the phone number. This will be copied onto the IncomingPhoneNumber resource.
**SmsMethod** | **string** | The HTTP method that should be used to request the SmsUrl. Must be either `GET` or `POST`.  This will be copied onto the IncomingPhoneNumber resource.
**SmsFallbackUrl** | **string** | A URL that Twilio will request if an error occurs requesting or executing the TwiML defined by SmsUrl. This will be copied onto the IncomingPhoneNumber resource.
**SmsFallbackMethod** | **string** | The HTTP method that should be used to request the SmsFallbackUrl. Must be either `GET` or `POST`. This will be copied onto the IncomingPhoneNumber resource.
**StatusCallbackUrl** | **string** | Optional. The Status Callback URL attached to the IncomingPhoneNumber resource.
**StatusCallbackMethod** | **string** | Optional. The Status Callback Method attached to the IncomingPhoneNumber resource.
**SmsApplicationSid** | **string** | Optional. The 34 character sid of the application Twilio should use to handle SMS messages sent to this number. If a `SmsApplicationSid` is present, Twilio will ignore all of the SMS urls above and use those set on the application.
**AddressSid** | **string** | Optional. A 34 character string that uniquely identifies the Address resource that represents the address of the owner of this phone number.
**Email** | **string** | Optional. Email of the owner of this phone number that is being hosted.
**VerificationType** | **string** | 
**VerificationDocumentSid** | **string** | Optional. The unique sid identifier of the Identity Document that represents the document for verifying ownership of the number to be hosted. Required when VerificationType is phone-bill.

### Return type

[**NumbersV3HostedNumberOrder**](NumbersV3HostedNumberOrder.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

