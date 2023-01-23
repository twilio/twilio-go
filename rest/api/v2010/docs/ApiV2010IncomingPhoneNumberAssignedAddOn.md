# ApiV2010IncomingPhoneNumberAssignedAddOn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | **string** | The unique string that identifies the resource |[optional] 
**AccountSid** | **string** | The SID of the Account that created the resource |[optional] 
**ResourceSid** | **string** | The SID of the Phone Number that installed this Add-on |[optional] 
**FriendlyName** | **string** | The string that you assigned to describe the resource |[optional] 
**Description** | **string** | A short description of the Add-on functionality |[optional] 
**Configuration** | Pointer to **interface{}** | A JSON string that represents the current configuration |
**UniqueName** | **string** | An application-defined string that uniquely identifies the resource |[optional] 
**DateCreated** | **string** | The RFC 2822 date and time in GMT that the resource was created |[optional] 
**DateUpdated** | **string** | The RFC 2822 date and time in GMT that the resource was last updated |[optional] 
**Uri** | **string** | The URI of the resource, relative to `https://api.twilio.com` |[optional] 
**SubresourceUris** | **map[string]interface{}** | A list of related resources identified by their relative URIs |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


