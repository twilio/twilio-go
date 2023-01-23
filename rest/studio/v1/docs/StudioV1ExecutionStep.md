# StudioV1ExecutionStep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | **string** | The unique string that identifies the resource |[optional] 
**AccountSid** | **string** | The SID of the Account that created the resource |[optional] 
**FlowSid** | **string** | The SID of the Flow |[optional] 
**ExecutionSid** | **string** | The SID of the Execution |[optional] 
**Name** | **string** | The event that caused the Flow to transition to the Step |[optional] 
**Context** | Pointer to **interface{}** | The current state of the flow |
**TransitionedFrom** | **string** | The Widget that preceded the Widget for the Step |[optional] 
**TransitionedTo** | **string** | The Widget that will follow the Widget for the Step |[optional] 
**DateCreated** | [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was created |[optional] 
**DateUpdated** | [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was last updated |[optional] 
**Url** | **string** | The absolute URL of the resource |[optional] 
**Links** | **map[string]interface{}** | The URLs of related resources |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


