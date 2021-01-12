# InlineObject14

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | **string** | This defaults to the AccountSid of the authorization the user is using. This can be provided to specify a subaccount to add the HostedNumberOrder to. | [optional] 
**AddressSid** | **string** | Optional. A 34 character string that uniquely identifies the Address resource that represents the address of the owner of this phone number. | [optional] 
**CcEmails** | **[]string** | Optional. A list of emails that the LOA document for this HostedNumberOrder will be carbon copied to. | [optional] 
**Email** | **string** | Optional. Email of the owner of this phone number that is being hosted. | [optional] 
**FriendlyName** | **string** | A 64 character string that is a human readable text that describes this resource. | [optional] 
**PhoneNumber** | **string** | The number to host in [+E.164](https://en.wikipedia.org/wiki/E.164) format | 
**SmsApplicationSid** | **string** | Optional. The 34 character sid of the application Twilio should use to handle SMS messages sent to this number. If a &#x60;SmsApplicationSid&#x60; is present, Twilio will ignore all of the SMS urls above and use those set on the application. | [optional] 
**SmsCapability** | **bool** | Used to specify that the SMS capability will be hosted on Twilio&#39;s platform. | 
**SmsFallbackMethod** | **string** | The HTTP method that should be used to request the SmsFallbackUrl. Must be either &#x60;GET&#x60; or &#x60;POST&#x60;. This will be copied onto the IncomingPhoneNumber resource. | [optional] 
**SmsFallbackUrl** | **string** | A URL that Twilio will request if an error occurs requesting or executing the TwiML defined by SmsUrl. This will be copied onto the IncomingPhoneNumber resource. | [optional] 
**SmsMethod** | **string** | The HTTP method that should be used to request the SmsUrl. Must be either &#x60;GET&#x60; or &#x60;POST&#x60;.  This will be copied onto the IncomingPhoneNumber resource. | [optional] 
**SmsUrl** | **string** | The URL that Twilio should request when somebody sends an SMS to the phone number. This will be copied onto the IncomingPhoneNumber resource. | [optional] 
**StatusCallbackMethod** | **string** | Optional. The Status Callback Method attached to the IncomingPhoneNumber resource. | [optional] 
**StatusCallbackUrl** | **string** | Optional. The Status Callback URL attached to the IncomingPhoneNumber resource. | [optional] 
**UniqueName** | **string** | Optional. Provides a unique and addressable name to be assigned to this HostedNumberOrder, assigned by the developer, to be optionally used in addition to SID. | [optional] 
**VerificationDocumentSid** | **string** | Optional. The unique sid identifier of the Identity Document that represents the document for verifying ownership of the number to be hosted. Required when VerificationType is phone-bill. | [optional] 
**VerificationType** | **string** | Optional. The method used for verifying ownership of the number to be hosted. One of phone-call (default) or phone-bill. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


