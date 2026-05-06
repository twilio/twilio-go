# ConversationsV2OperationStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OperationId** | **string** | Unique identifier for the long-running operation. |
**Status** | **string** | Current status of the operation. |
**CreatedAt** | [**time.Time**](time.Time.md) | Timestamp when the operation was created. |
**CompletedAt** | Pointer to [**time.Time**](time.Time.md) | Timestamp when the operation completed. Only present for completed or failed operations. |
**StatusUrl** | **string** | URL to poll for operation status. |
**Error** | Pointer to [**FetchOperationStatusResponseError**](FetchOperationStatusResponseError.md) |  |
**Related** | Pointer to **map[string]string** | Named resource identifiers associated with this operation. Keys depend on the operation type: - config-create, config-update, config-delete: configurationId - conversation-delete: conversationId  |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


