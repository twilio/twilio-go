# ProfileImportCreatedResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ImportId** | **string** | ID of the import task. |
**Url** | **string** | Pre-signed URL to upload the CSV via a single PUT request. |
**OperationId** | **string** | The ID of an asynchronous profile domain operation. |[optional] 
**CreatedAt** | [**time.Time**](time.Time.md) | ISO 8601 timestamp when the operation was created. |[optional] 
**StatusUrl** | **string** | URL to query the operation status. |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


