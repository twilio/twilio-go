# ProxyV1Session

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | **string** | The unique string that identifies the resource |[optional] 
**ServiceSid** | **string** | The SID of the resource's parent Service |[optional] 
**AccountSid** | **string** | The SID of the Account that created the resource |[optional] 
**DateStarted** | [**time.Time**](time.Time.md) | The ISO 8601 date when the Session started |[optional] 
**DateEnded** | [**time.Time**](time.Time.md) | The ISO 8601 date when the Session ended |[optional] 
**DateLastInteraction** | [**time.Time**](time.Time.md) | The ISO 8601 date when the Session last had an interaction |[optional] 
**DateExpiry** | [**time.Time**](time.Time.md) | The ISO 8601 date when the Session should expire |[optional] 
**UniqueName** | **string** | An application-defined string that uniquely identifies the resource |[optional] 
**Status** | Pointer to [**string**](SessionEnumStatus.md) |  |
**ClosedReason** | **string** | The reason the Session ended |[optional] 
**Ttl** | **int** | When the session will expire |[optional] 
**Mode** | Pointer to [**string**](SessionEnumMode.md) |  |
**DateCreated** | [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was created |[optional] 
**DateUpdated** | [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was last updated |[optional] 
**Url** | **string** | The absolute URL of the Session resource |[optional] 
**Links** | **map[string]interface{}** | The URLs of resources related to the Session |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


