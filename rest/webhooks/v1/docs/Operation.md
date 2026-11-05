# Operation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique operation identifier (TypeID/TTID format) |
**StatusUrl** | **string** | URL to poll for operation status (same as Location header in 202 response) |
**Status** | **string** | Current state of the operation: - PENDING: Accepted, not yet started - RUNNING: Propagating to all regions (typically 30-60 seconds) - COMPLETED: Rule active globally - CANCELLED: Operation was cancelled before completion - FAILED: Operation could not complete; see error field for details  |
**Error** | Pointer to [**OperationError**](OperationError.md) | Present only when status is FAILED |
**ResourceId** | **string** | ID of the affected resource |[optional] 
**CreatedAt** | [**time.Time**](time.Time.md) |  |[readonly] 
**CompletedAt** | [**time.Time**](time.Time.md) |  |[optional] [readonly] 
**RetentionPeriod** | **string** | ISO-8601 duration for how long this Operation record is retained after creation (fixed at 7 days) |[readonly] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


