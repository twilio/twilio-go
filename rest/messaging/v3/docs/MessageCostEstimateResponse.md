# MessageCostEstimateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Currency** | **string** | The WhatsApp Business Account's currency (e.g. `USD`). `bidAmount` and cost bounds are in this currency's smallest unit. |[optional] 
**Country** | **string** | The ISO-3166-1 alpha-2 country code the estimate was computed for. |[optional] 
**LookbackWindow** | **string** | The historical lookback window used to compute the estimate. |[optional] 
**Estimates** | [**[]MessageCostEstimateBidLevel**](MessageCostEstimateBidLevel.md) | One entry per bid level. |[optional] 
**Disclaimer** | **string** | Estimates are historical and informational only; actual delivery and cost may vary. |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


