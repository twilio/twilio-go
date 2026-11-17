# MessageCostEstimateBidLevel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BidAmount** | **int64** | Max price per 1,000 message deliveries, in the smallest unit of `currency`. |[optional] 
**TargetedUsers** | **int64** | The normalized user base the estimate is scaled to (fixed at 1,000 during beta). |[optional] 
**EstimatedDeliveries** | Pointer to [**MessageCostEstimateRange**](MessageCostEstimateRange.md) |  |
**EstimatedCost** | Pointer to [**MessageCostEstimateRange**](MessageCostEstimateRange.md) |  |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


