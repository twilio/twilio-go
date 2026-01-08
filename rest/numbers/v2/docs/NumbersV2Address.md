# NumbersV2Address

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | **string** | A 34 character string that uniquely identifies this Address resource. |[optional] 
**AccountSid** | **string** | The unique SID identifier of the Account. |[optional] 
**CustomerName** | Pointer to **string** | The name of the customer associated with this address. |
**DateCreated** | [**time.Time**](time.Time.md) | The date that this resource was created, given in RFC 2822 format. |[optional] 
**DateUpdated** | [**time.Time**](time.Time.md) | The date that this resource was last updated, given in RFC 2822 format. |[optional] 
**EmergencyEnabled** | **bool** | Whether this address is enabled for emergency services. |[optional] 
**FriendlyName** | Pointer to **string** | A human-readable description of this resource, up to 64 characters. |
**IsoCountry** | **string** | The ISO country code of this address. |[optional] 
**Locality** | Pointer to **string** | The locality or city of this address. |
**PostalCode** | Pointer to **string** | The postal code of this address. |
**Region** | Pointer to **string** | The state or region of this address. |
**Source** | **string** | The source system that created this address. |[optional] 
**Status** | **string** | The status of this address. |[optional] 
**Street** | Pointer to **string** | The street address. |
**StreetSecondary** | Pointer to **string** | The additional street address information. |
**SubresourceUris** | **map[string]interface{}** | A list of related resources identified by their URIs. |[optional] 
**Validated** | **bool** | Whether this address has been validated. |[optional] 
**Verified** | **bool** | Whether this address has been verified. |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


