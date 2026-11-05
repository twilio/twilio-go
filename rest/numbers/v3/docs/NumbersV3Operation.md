# NumbersV3Operation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OperationId** | **string** | The operation ID. Treat it as opaque. |
**Status** | **string** |  |
**CreatedAt** | [**time.Time**](time.Time.md) | When the request was accepted. |
**StatusUrl** | **string** | The operation's URL. Follow it as returned. |[optional] 
**CompletedAt** | Pointer to [**time.Time**](time.Time.md) | When the operation reached a final status. |
**Payload** | Pointer to [**NumbersV3NumbersSendersOperationPayload**](NumbersV3NumbersSendersOperationPayload.md) |  |
**Result** | Pointer to [**NumbersV3NumbersSendersOperationResult**](NumbersV3NumbersSendersOperationResult.md) |  |
**ResultId** | Pointer to **string** | The Compliance Registration ID of a `COMPLETED` association. `null` otherwise. |
**ResultUrl** | Pointer to **string** | The resource's ComplianceRegistrations URL, set with `resultId`. `null` otherwise. |
**ResultRetentionPeriod** | Pointer to **string** | Bulk only. ISO 8601 duration, from `createdAt`, that the results stay available. |
**Error** | Pointer to [**NumbersV3NumbersSendersOperationError**](NumbersV3NumbersSendersOperationError.md) |  |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


