# NumbersV3NumbersSendersOperationPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestType** | **string** | `CREATE` for an association, `DELETE` for a removal. |
**ResourceType** | **string** | `ComplianceRegistration`. More may be added. |
**TotalItems** | Pointer to **int** | Bulk only. Null until known. |
**ProcessedItems** | Pointer to **int** | Bulk only. |
**FailedItems** | Pointer to **int** | Bulk only. Null when not counted. |
**UpdatedAt** | Pointer to [**time.Time**](time.Time.md) | Bulk only. When progress last changed. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


