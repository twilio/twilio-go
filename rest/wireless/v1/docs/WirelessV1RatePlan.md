# WirelessV1RatePlan

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | Pointer to **string** | The unique string that identifies the resource |
**UniqueName** | Pointer to **string** | An application-defined string that uniquely identifies the resource |
**AccountSid** | Pointer to **string** | The SID of the Account that created the resource |
**FriendlyName** | Pointer to **string** | The string that you assigned to describe the resource |
**DataEnabled** | Pointer to **bool** | Whether SIMs can use GPRS/3G/4G/LTE data connectivity |
**DataMetering** | Pointer to **string** | The model used to meter data usage |
**DataLimit** | Pointer to **int** | The total data usage in Megabytes that the Network allows during one month on the home network |
**MessagingEnabled** | Pointer to **bool** | Whether SIMs can make, send, and receive SMS using Commands |
**VoiceEnabled** | Pointer to **bool** | Deprecated. Whether SIMs can make and receive voice calls |
**NationalRoamingEnabled** | Pointer to **bool** | Whether SIMs can roam on networks other than the home network in the United States |
**NationalRoamingDataLimit** | Pointer to **int** | The total data usage in Megabytes that the Network allows during one month on non-home networks in the United States |
**InternationalRoaming** | Pointer to **[]string** | The services that SIMs capable of using GPRS/3G/4G/LTE data connectivity can use outside of the United States |
**InternationalRoamingDataLimit** | Pointer to **int** | The total data usage (download and upload combined) in Megabytes that the Network allows during one month when roaming outside the United States |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The date when the resource was created, given as GMT in ISO 8601 format |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The date when the resource was last updated, given as GMT in ISO 8601 format |
**Url** | Pointer to **string** | The absolute URL of the resource |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


