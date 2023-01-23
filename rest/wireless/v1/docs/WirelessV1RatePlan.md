# WirelessV1RatePlan

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | **string** | The unique string that identifies the resource |[optional] 
**UniqueName** | **string** | An application-defined string that uniquely identifies the resource |[optional] 
**AccountSid** | **string** | The SID of the Account that created the resource |[optional] 
**FriendlyName** | **string** | The string that you assigned to describe the resource |[optional] 
**DataEnabled** | **bool** | Whether SIMs can use GPRS/3G/4G/LTE data connectivity |[optional] 
**DataMetering** | **string** | The model used to meter data usage |[optional] 
**DataLimit** | **int** | The total data usage in Megabytes that the Network allows during one month on the home network |[optional] 
**MessagingEnabled** | **bool** | Whether SIMs can make, send, and receive SMS using Commands |[optional] 
**VoiceEnabled** | **bool** | Deprecated. Whether SIMs can make and receive voice calls |[optional] 
**NationalRoamingEnabled** | **bool** | Whether SIMs can roam on networks other than the home network in the United States |[optional] 
**NationalRoamingDataLimit** | **int** | The total data usage in Megabytes that the Network allows during one month on non-home networks in the United States |[optional] 
**InternationalRoaming** | **[]string** | The services that SIMs capable of using GPRS/3G/4G/LTE data connectivity can use outside of the United States |[optional] 
**InternationalRoamingDataLimit** | **int** | The total data usage (download and upload combined) in Megabytes that the Network allows during one month when roaming outside the United States |[optional] 
**DateCreated** | [**time.Time**](time.Time.md) | The date when the resource was created, given as GMT in ISO 8601 format |[optional] 
**DateUpdated** | [**time.Time**](time.Time.md) | The date when the resource was last updated, given as GMT in ISO 8601 format |[optional] 
**Url** | **string** | The absolute URL of the resource |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


