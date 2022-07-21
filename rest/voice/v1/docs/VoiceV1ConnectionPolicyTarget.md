# VoiceV1ConnectionPolicyTarget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** | The SID of the Account that created the resource |
**ConnectionPolicySid** | Pointer to **string** | The SID of the Connection Policy that owns the Target |
**Sid** | Pointer to **string** | The unique string that identifies the resource |
**FriendlyName** | Pointer to **string** | The string that you assigned to describe the resource |
**Target** | Pointer to **string** | The SIP address you want Twilio to route your calls to |
**Priority** | Pointer to **int** | The relative importance of the target |
**Weight** | Pointer to **int** | The value that determines the relative load the Target should receive compared to others with the same priority |
**Enabled** | Pointer to **bool** | Whether the target is enabled |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The RFC 2822 date and time in GMT when the resource was created |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The RFC 2822 date and time in GMT when the resource was last updated |
**Url** | Pointer to **string** | The absolute URL of the resource |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


