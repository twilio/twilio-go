# ApiV2010ShortCode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created this ShortCode resource. |
**ApiVersion** | Pointer to **string** | The API version used to start a new TwiML session when an SMS message is sent to this short code. |
**DateCreated** | Pointer to **string** | The date and time in GMT that this resource was created specified in [RFC 2822](https://www.ietf.org/rfc/rfc2822.txt) format. |
**DateUpdated** | Pointer to **string** | The date and time in GMT that this resource was last updated, specified in [RFC 2822](https://www.ietf.org/rfc/rfc2822.txt) format. |
**FriendlyName** | Pointer to **string** | A string that you assigned to describe this resource. By default, the `FriendlyName` is the short code. |
**ShortCode** | Pointer to **string** | The short code. e.g., 894546. |
**Sid** | Pointer to **string** | The unique string that that we created to identify this ShortCode resource. |
**SmsFallbackMethod** | Pointer to **string** | The HTTP method we use to call the `sms_fallback_url`. Can be: `GET` or `POST`. |
**SmsFallbackUrl** | Pointer to **string** | The URL that we call if an error occurs while retrieving or executing the TwiML from `sms_url`. |
**SmsMethod** | Pointer to **string** | The HTTP method we use to call the `sms_url`. Can be: `GET` or `POST`. |
**SmsUrl** | Pointer to **string** | The URL we call when receiving an incoming SMS message to this short code. |
**Uri** | Pointer to **string** | The URI of this resource, relative to `https://api.twilio.com`. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


