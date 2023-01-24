# StudioV2ExecutionStep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | Pointer to **string** | The unique string that we created to identify the ExecutionStep resource. |
**AccountSid** | Pointer to **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the ExecutionStep resource. |
**FlowSid** | Pointer to **string** | The SID of the Flow. |
**ExecutionSid** | Pointer to **string** | The SID of the Step's Execution resource. |
**Name** | Pointer to **string** | The event that caused the Flow to transition to the Step. |
**Context** | Pointer to **interface{}** | The current state of the Flow's Execution. As a flow executes, we save its state in this context. We save data that your widgets can access as variables in configuration fields or in text areas as variable substitution. |
**TransitionedFrom** | Pointer to **string** | The Widget that preceded the Widget for the Step. |
**TransitionedTo** | Pointer to **string** | The Widget that will follow the Widget for the Step. |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The date and time in GMT when the resource was created specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The date and time in GMT when the resource was last updated specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. |
**Url** | Pointer to **string** | The absolute URL of the resource. |
**Links** | Pointer to **map[string]interface{}** | The URLs of related resources. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


