# ProxyV1Participant

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | **string** | The unique string that identifies the resource |[optional] 
**SessionSid** | **string** | The SID of the resource's parent Session |[optional] 
**ServiceSid** | **string** | The SID of the resource's parent Service |[optional] 
**AccountSid** | **string** | The SID of the Account that created the resource |[optional] 
**FriendlyName** | **string** | The string that you assigned to describe the participant |[optional] 
**Identifier** | **string** | The phone number or channel identifier of the Participant |[optional] 
**ProxyIdentifier** | **string** | The phone number or short code of the participant's partner |[optional] 
**ProxyIdentifierSid** | **string** | The SID of the Proxy Identifier assigned to the Participant |[optional] 
**DateDeleted** | [**time.Time**](time.Time.md) | The ISO 8601 date the Participant was removed |[optional] 
**DateCreated** | [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was created |[optional] 
**DateUpdated** | [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was last updated |[optional] 
**Url** | **string** | The absolute URL of the Participant resource |[optional] 
**Links** | **map[string]interface{}** | The URLs to resources related the participant |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


