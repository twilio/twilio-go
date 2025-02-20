# FlexV1InteractionTransfer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | Pointer to **string** | The unique string created by Twilio to identify an Interaction Transfer resource. |
**InstanceSid** | Pointer to **string** | The SID of the Instance associated with the Transfer. |
**AccountSid** | Pointer to **string** | The SID of the Account that created the Transfer. |
**InteractionSid** | Pointer to **string** | The Interaction Sid for this channel. |
**ChannelSid** | Pointer to **string** | The Channel Sid for this Transfer. |
**ExecutionSid** | Pointer to **string** | The Execution SID associated with the Transfer. |
**Type** | Pointer to [**string**](InteractionTransferEnumTransferType.md) |  |
**Status** | Pointer to [**string**](InteractionTransferEnumTransferStatus.md) |  |
**From** | Pointer to **string** | The SID of the Participant initiating the Transfer. |
**To** | Pointer to **string** | The SID of the Participant receiving the Transfer. |
**NoteSid** | Pointer to **string** | The SID of the Note associated with the Transfer. |
**SummarySid** | Pointer to **string** | The SID of the Summary associated with the Transfer. |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The date and time when the Transfer was created. |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The date and time when the Transfer was last updated. |
**Url** | Pointer to **string** |  |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


