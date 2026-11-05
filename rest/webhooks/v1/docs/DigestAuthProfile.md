# DigestAuthProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique AuthProfile identifier (TypeID/TTID format) |[readonly] 
**Type** | **string** | Authentication type |
**FriendlyName** | **string** | Human-readable name for identification |
**CreatedAt** | [**time.Time**](time.Time.md) |  |[readonly] 
**UpdatedAt** | [**time.Time**](time.Time.md) |  |[readonly] 
**Username** | **string** | Digest auth username |
**Password** | **string** | Digest auth password (never returned, stored securely as secret) |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


