# MessagingV1CreateMessagesResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalMessageCount** | **int** | The number of Messages processed in the request, equal to the sum of success_count and error_count. |[optional] 
**SuccessCount** | **int** | The number of Messages successfully created. |[optional] 
**ErrorCount** | **int** | The number of Messages unsuccessfully processed in the request. |[optional] 
**MessageReceipts** | [**[]MessagingV1MessageReceipt**](MessagingV1MessageReceipt.md) |  |[optional] 
**FailedMessageReceipts** | [**[]MessagingV1FailedMessageReceipt**](MessagingV1FailedMessageReceipt.md) |  |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


