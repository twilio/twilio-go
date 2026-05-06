# ExecutionDetailsTrigger

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**On** | **string** | The conversational lifecycle event that activated execution of the Rule. Available values are: - `COMMUNICATION`: Trigger the Rule on each communication within the Conversation. - `CONVERSATION_END`: Trigger the Rule when the Conversation moves to the `closed` state. - `CONVERSATION_INACTIVE`: Trigger the Rule when the Conversation moves to the `inactive` state.  |
**Timestamp** | [**time.Time**](time.Time.md) | Timestamp of when the trigger was activated. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


