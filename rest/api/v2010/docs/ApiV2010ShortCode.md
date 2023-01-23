# ApiV2010ShortCode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | **string** | The SID of the Account that created this resource |[optional] 
**ApiVersion** | **string** | The API version used to start a new TwiML session |[optional] 
**DateCreated** | **string** | The RFC 2822 date and time in GMT that this resource was created |[optional] 
**DateUpdated** | **string** | The RFC 2822 date and time in GMT that this resource was last updated |[optional] 
**FriendlyName** | **string** | A string that you assigned to describe this resource |[optional] 
**ShortCode** | **string** | The short code. e.g., 894546. |[optional] 
**Sid** | **string** | The unique string that identifies this resource |[optional] 
**SmsFallbackMethod** | **string** | HTTP method we use to call the sms_fallback_url |[optional] 
**SmsFallbackUrl** | **string** | URL Twilio will request if an error occurs in executing TwiML |[optional] 
**SmsMethod** | **string** | HTTP method to use when requesting the sms url |[optional] 
**SmsUrl** | **string** | URL we call when receiving an incoming SMS message to this short code |[optional] 
**Uri** | **string** | The URI of this resource, relative to `https://api.twilio.com` |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


