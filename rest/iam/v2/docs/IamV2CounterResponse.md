# IamV2CounterResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** |  |
**CounterKey** | Pointer to **string** |  |
**MaxQuota** | Pointer to **int** |  |
**DailyThroughputQuota** | Pointer to **int** |  |
**TotalUsed** | **int** | The total amount used against the quota |[optional] 
**DailyThroughputUsed** | Pointer to **int** |  |
**ExpiresAt** | [**time.Time**](time.Time.md) | The timestamp when the counter expires |[optional] 
**LastUpdatedAt** | [**time.Time**](time.Time.md) | The timestamp when the counter was last updated |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


