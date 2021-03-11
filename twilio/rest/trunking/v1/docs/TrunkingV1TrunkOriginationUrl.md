# TrunkingV1TrunkOriginationUrl

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** | The SID of the Account that created the resource |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The RFC 2822 date and time in GMT when the resource was created |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The RFC 2822 date and time in GMT when the resource was last updated |
**Enabled** | Pointer to **bool** | Whether the URL is enabled |
**FriendlyName** | Pointer to **string** | The string that you assigned to describe the resource |
**Priority** | Pointer to **int32** | The relative importance of the URI |
**Sid** | Pointer to **string** | The unique string that identifies the resource |
**SipUrl** | Pointer to **string** | The SIP address you want Twilio to route your Origination calls to |
**TrunkSid** | Pointer to **string** | The SID of the Trunk that owns the Origination URL |
**Url** | Pointer to **string** | The absolute URL of the resource |
**Weight** | Pointer to **int32** | The value that determines the relative load the URI should receive compared to others with the same priority |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


