# OperationStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OperationId** | **string** | The unique identifier for this operation. |
**Status** | **string** | The current status of the operation. |
**CreatedAt** | [**time.Time**](time.Time.md) | When the operation was created. |
**StatusUrl** | **string** | URI to check operation status. |[optional] 
**CompletedAt** | [**time.Time**](time.Time.md) | When the operation completed or failed. |[optional] 
**Result** | [**OperationResultResourceId**](OperationResultResourceId.md) |  |[optional] 
**Error** | [**OperationStatusError**](OperationStatusError.md) |  |[optional] 
**ResultUrl** | **string** | URL to fetch the resulting resource. |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


