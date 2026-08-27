# TypingIndicatorRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Channel** | **string** | The messaging channel. Must be \"RCS\". |
**MessageId** | **string** | The SID of a recent inbound message from the recipient. Must be an SM or MM SID format.  |
**From** | **string** | The RCS agent identifier of the sender (business). |
**To** | **string** | The RCS recipient identifier in E.164 format prefixed with \"rcs:\". |
**Event** | **string** | The type of typing event. Currently only \"START\" is supported for RCS, indicating the agent began typing. Defaults to \"START\".  |[optional] [default to "START"]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


