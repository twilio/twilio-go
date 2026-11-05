# OperationPayloadBase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Identifies which fields, if any, are present beyond the standard ones. `BASE` means only the standard fields are present. |
**TotalCount** | **int64** | Total number of items to be processed. |[optional] 
**SucceededCount** | **int64** | Number of items successfully processed. |[optional] 
**FailedCount** | **int64** | Number of items that failed processing. |[optional] 
**FailedEventOrdinals** | **[]string** | List of ordinal identifiers (row indices or event coordinates) that failed processing. Provided for diagnostic purposes. May be truncated for large failure sets.  |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


