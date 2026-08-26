# QueryJobStatusResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OperationId** | **string** | The unique identifier for the asynchronous query operation, in TTID format. |
**Status** | [**LongRunningOperationStatus**](LongRunningOperationStatus.md) |  |
**StatusUrl** | **string** |  |[optional] 
**CreatedAt** | [**time.Time**](time.Time.md) |  |
**CompletedAt** | Pointer to [**time.Time**](time.Time.md) | The time when the operation reached a terminal state (CANCELLED, COMPLETED, or FAILED), in RFC 3339 UTC format. Null while the operation is still PENDING or RUNNING. |
**Error** | Pointer to [**OperationError**](OperationError.md) | The error details, present only when status is FAILED. |
**ResultUrl** | Pointer to **string** | The URL to retrieve query results, present only when status is COMPLETED. Supports pagination via the pageSize and pageToken query parameters.  |
**ResultId** | Pointer to **string** | The identifier of the results resource, present only when status is COMPLETED. Equal to operationId, since results are keyed by the operation ID.  |
**ResultRetentionPeriod** | Pointer to **string** | The duration for how long results are retained, in ISO 8601 format. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


