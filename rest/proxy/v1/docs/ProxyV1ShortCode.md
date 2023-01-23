# ProxyV1ShortCode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | **string** | The unique string that identifies the resource |[optional] 
**AccountSid** | **string** | The SID of the Account that created the resource |[optional] 
**ServiceSid** | **string** | The SID of the resource's parent Service |[optional] 
**DateCreated** | [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was created |[optional] 
**DateUpdated** | [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was last updated |[optional] 
**ShortCode** | **string** | The short code's number |[optional] 
**IsoCountry** | **string** | The ISO Country Code |[optional] 
**Capabilities** | [**ProxyV1ServiceShortCodeCapabilities**](ProxyV1ServiceShortCodeCapabilities.md) |  |[optional] 
**Url** | **string** | The absolute URL of the ShortCode resource |[optional] 
**IsReserved** | **bool** | Whether the short code should be reserved for manual assignment to participants only |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


