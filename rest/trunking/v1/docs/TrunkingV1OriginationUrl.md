# TrunkingV1OriginationUrl

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the OriginationUrl resource. |
**Sid** | Pointer to **string** | The unique string that we created to identify the OriginationUrl resource. |
**TrunkSid** | Pointer to **string** | The SID of the Trunk that owns the Origination URL. |
**Weight** | Pointer to **int** | The value that determines the relative share of the load the URI should receive compared to other URIs with the same priority. Can be an integer from 1 to 65535, inclusive, and the default is 10. URLs with higher values receive more load than those with lower ones with the same priority. |
**Enabled** | Pointer to **bool** | Whether the URL is enabled. The default is `true`. |
**SipUrl** | Pointer to **string** | The SIP address you want Twilio to route your Origination calls to. This must be a `sip:` schema. |
**FriendlyName** | Pointer to **string** | The string that you assigned to describe the resource. |
**Priority** | Pointer to **int** | The relative importance of the URI. Can be an integer from 0 to 65535, inclusive, and the default is 10. The lowest number represents the most important URI. |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The date and time in GMT when the resource was created specified in [RFC 2822](https://www.ietf.org/rfc/rfc2822.txt) format. |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The date and time in GMT when the resource was last updated specified in [RFC 2822](https://www.ietf.org/rfc/rfc2822.txt) format. |
**Url** | Pointer to **string** | The absolute URL of the resource. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


