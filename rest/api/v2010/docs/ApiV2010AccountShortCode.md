# ApiV2010AccountShortCode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** | The SID of the Account that created this resource |
**ApiVersion** | Pointer to **string** | The API version used to start a new TwiML session |
**DateCreated** | Pointer to **string** | The RFC 2822 date and time in GMT that this resource was created |
**DateUpdated** | Pointer to **string** | The RFC 2822 date and time in GMT that this resource was last updated |
**FriendlyName** | Pointer to **string** | A string that you assigned to describe this resource |
**ShortCode** | Pointer to **string** | The short code. e.g., 894546. |
**Sid** | Pointer to **string** | The unique string that identifies this resource |
**SmsFallbackMethod** | Pointer to **string** | HTTP method we use to call the sms_fallback_url |
**SmsFallbackUrl** | Pointer to **string** | URL Twilio will request if an error occurs in executing TwiML |
**SmsMethod** | Pointer to **string** | HTTP method to use when requesting the sms url |
**SmsUrl** | Pointer to **string** | URL we call when receiving an incoming SMS message to this short code |
**Uri** | Pointer to **string** | The URI of this resource, relative to `https://api.twilio.com` |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


