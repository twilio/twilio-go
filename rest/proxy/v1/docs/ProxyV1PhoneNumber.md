# ProxyV1PhoneNumber

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | **string** | The unique string that identifies the resource |[optional] 
**AccountSid** | **string** | The SID of the Account that created the resource |[optional] 
**ServiceSid** | **string** | The SID of the PhoneNumber resource's parent Service resource |[optional] 
**DateCreated** | [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was created |[optional] 
**DateUpdated** | [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was last updated |[optional] 
**PhoneNumber** | **string** | The phone number in E.164 format |[optional] 
**FriendlyName** | **string** | The string that you assigned to describe the resource |[optional] 
**IsoCountry** | **string** | The ISO Country Code |[optional] 
**Capabilities** | [**ProxyV1ServicePhoneNumberCapabilities**](ProxyV1ServicePhoneNumberCapabilities.md) |  |[optional] 
**Url** | **string** | The absolute URL of the PhoneNumber resource |[optional] 
**IsReserved** | **bool** | Reserve the phone number for manual assignment to participants only |[optional] 
**InUse** | **int** | The number of open session assigned to the number. |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


