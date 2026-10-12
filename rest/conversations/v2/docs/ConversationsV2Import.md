# ConversationsV2Import

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  |
**Status** | [**ConversationsV2ImportStatus**](ConversationsV2ImportStatus.md) |  |
**Sources** | [**[]ConversationsV2ImportSource**](ConversationsV2ImportSource.md) |  |
**CreatedAt** | [**time.Time**](time.Time.md) |  |
**UpdatedAt** | [**time.Time**](time.Time.md) |  |
**CompletedAt** | [**time.Time**](time.Time.md) | Only present once the import has finished. |[optional] 
**FailureReason** | **string** | Why the import failed. Only present when status is FAILED. |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


