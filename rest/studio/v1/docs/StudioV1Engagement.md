# StudioV1Engagement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | Pointer to **string** | The unique string that identifies the resource |
**AccountSid** | Pointer to **string** | The SID of the Account that created the resource |
**FlowSid** | Pointer to **string** | The SID of the Flow |
**ContactSid** | Pointer to **string** | The SID of the Contact |
**ContactChannelAddress** | Pointer to **string** | The phone number, SIP address or Client identifier that triggered this Engagement |
**Context** | Pointer to **interface{}** | The current state of the execution flow |
**Status** | Pointer to [**string**](EngagementEnumStatus.md) |  |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the Engagement was created |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the Engagement was last updated |
**Url** | Pointer to **string** | The absolute URL of the resource |
**Links** | Pointer to **map[string]interface{}** | The URLs of the Engagement's nested resources |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


