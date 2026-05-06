# ListOperatorVersionsResponseItems

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | **int** | The numeric version number. |
**Status** | [**OperatorVersionStatus**](OperatorVersionStatus.md) |  |
**DateCreated** | [**time.Time**](time.Time.md) | Timestamp of when this version was created. |
**DateDeprecated** | [**time.Time**](time.Time.md) | Timestamp of when this version was deprecated. Present when status is `DEPRECATED` or `RETIRED`. |[optional] 
**RetirementDate** | [**time.Time**](time.Time.md) | Scheduled retirement date for this version. Present when status is `DEPRECATED`. |[optional] 
**DateRetired** | [**time.Time**](time.Time.md) | Timestamp of when this version was retired. Present when status is `RETIRED`. |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


