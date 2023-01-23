# MessagingV1ShortCode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | **string** | The unique string that identifies the resource |[optional] 
**AccountSid** | **string** | The SID of the Account that created the resource |[optional] 
**ServiceSid** | **string** | The SID of the Service that the resource is associated with |[optional] 
**DateCreated** | [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was created |[optional] 
**DateUpdated** | [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was last updated |[optional] 
**ShortCode** | **string** | The E.164 format of the short code |[optional] 
**CountryCode** | **string** | The 2-character ISO Country Code of the number |[optional] 
**Capabilities** | **[]string** | An array of values that describe whether the number can receive calls or messages |[optional] 
**Url** | **string** | The absolute URL of the ShortCode resource |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


