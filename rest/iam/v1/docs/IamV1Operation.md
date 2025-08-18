# IamV1Operation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | ID of the operation |[optional] 
**ParentId** | **string** | ID of the operation's parent operation, if applicable |[optional] 
**Type** | **string** | The type of the operation. |[optional] 
**State** | **string** |  |[optional] 
**DateCreated** | [**time.Time**](time.Time.md) | The date that this Operation was created, given in RFC 2822 format. |[optional] 
**DateCompleted** | [**time.Time**](time.Time.md) | The date that the Operation was completed, if applicable. |[optional] 
**Meta** | [**IamV1OperationMeta**](IamV1OperationMeta.md) |  |[optional] 
**Input** | **map[string]interface{}** | A polymorphic input object. Fields vary depending on the operation. |[optional] 
**Result** | **map[string]interface{}** | A polymorphic result object. Fields vary depending on the operation. Common fields include:   - `resourceId`: string  |[optional] 
**Error** | [**IamV1OperationError**](IamV1OperationError.md) |  |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


