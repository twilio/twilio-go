# StudioV1FlowEngagementStep

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** | The SID of the Account that created the resource |
**Context** | Pointer to **map[string]interface{}** | The current state of the flow |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was created |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was last updated |
**EngagementSid** | Pointer to **string** | The SID of the Engagement |
**FlowSid** | Pointer to **string** | The SID of the Flow |
**Links** | Pointer to **map[string]interface{}** | The URLs of related resources |
**Name** | Pointer to **string** | The event that caused the Flow to transition to the Step |
**Sid** | Pointer to **string** | The unique string that identifies the resource |
**TransitionedFrom** | Pointer to **string** | The Widget that preceded the Widget for the Step |
**TransitionedTo** | Pointer to **string** | The Widget that will follow the Widget for the Step |
**Url** | Pointer to **string** | The absolute URL of the resource |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


