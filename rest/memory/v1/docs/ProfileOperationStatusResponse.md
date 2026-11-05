# ProfileOperationStatusResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OperationId** | **string** | The ID of an asynchronous profile domain operation. |
**Status** | **string** | Current operation status: - PENDING: Operation accepted but not yet started - RUNNING: Operation is actively processing - CANCELLED: Operation was cancelled by user - COMPLETED: Operation finished successfully - FAILED: Operation failed and cannot complete  |
**CreatedAt** | [**time.Time**](time.Time.md) | ISO 8601 timestamp when the operation was created. |
**CompletedAt** | Pointer to [**time.Time**](time.Time.md) | ISO 8601 timestamp when the operation reached a terminal state. Null if still active. |
**Error** | Pointer to [**ProfileOperationStatusResponseError**](ProfileOperationStatusResponseError.md) |  |
**Payload** | Pointer to [**BaseOperationPayload**](BaseOperationPayload.md) |  |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


