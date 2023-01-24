# StudioV1ExecutionStepContext

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the ExecutionStepContext resource. |
**Context** | Pointer to **interface{}** | The current state of the Flow's Execution. As a flow executes, we save its state in this context. We save data that your widgets can access as variables in configuration fields or in text areas as variable substitution. |
**ExecutionSid** | Pointer to **string** | The SID of the context's Execution resource. |
**FlowSid** | Pointer to **string** | The SID of the Flow. |
**StepSid** | Pointer to **string** | The SID of the Step that the context is associated with. |
**Url** | Pointer to **string** | The absolute URL of the resource. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


