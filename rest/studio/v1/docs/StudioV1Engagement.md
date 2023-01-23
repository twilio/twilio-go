# StudioV1Engagement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | **string** | The unique string that identifies the resource |[optional] 
**AccountSid** | **string** | The SID of the Account that created the resource |[optional] 
**FlowSid** | **string** | The SID of the Flow |[optional] 
**ContactSid** | **string** | The SID of the Contact |[optional] 
**ContactChannelAddress** | **string** | The phone number, SIP address or Client identifier that triggered this Engagement |[optional] 
**Context** | Pointer to **interface{}** | The current state of the execution flow |
**Status** | Pointer to [**string**](EngagementEnumStatus.md) |  |
**DateCreated** | [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the Engagement was created |[optional] 
**DateUpdated** | [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the Engagement was last updated |[optional] 
**Url** | **string** | The absolute URL of the resource |[optional] 
**Links** | **map[string]interface{}** | The URLs of the Engagement's nested resources |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


