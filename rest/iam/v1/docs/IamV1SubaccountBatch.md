# IamV1SubaccountBatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BatchId** | **string** | The unique identifier for this batch operation, also used as the public Job ID. |
**AccountId** | **string** | The ID of the parent account that owns this batch. |
**WorkflowType** | [**IamV1SubaccountBatchWorkflowType**](IamV1SubaccountBatchWorkflowType.md) |  |
**Status** | [**IamV1SubaccountBatchStatus**](IamV1SubaccountBatchStatus.md) |  |
**ActorUserId** | **string** | The ID of the user who submitted the batch. |
**ActorUserEmail** | Pointer to **string** | The email of the user who submitted the batch. |
**TotalCount** | Pointer to **int** | The total number of subaccounts in the batch. |
**SuccessCount** | Pointer to **int** | The number of subaccounts processed successfully. |
**FailureCount** | Pointer to **int** | The number of subaccounts that failed processing. |
**DateCreated** | [**time.Time**](time.Time.md) | The date the batch was submitted, given in RFC 3339 format. |
**DateUpdated** | [**time.Time**](time.Time.md) | The date the batch was last updated, given in RFC 3339 format. |
**DateCompleted** | Pointer to [**time.Time**](time.Time.md) | The date the batch reached a terminal state, given in RFC 3339 format. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


