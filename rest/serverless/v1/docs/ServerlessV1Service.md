# ServerlessV1Service

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | **string** | The unique string that identifies the Service resource |[optional] 
**AccountSid** | **string** | The SID of the Account that created the Service resource |[optional] 
**FriendlyName** | **string** | The string that you assigned to describe the Service resource |[optional] 
**UniqueName** | **string** | A user-defined string that uniquely identifies the Service resource |[optional] 
**IncludeCredentials** | **bool** | Whether to inject Account credentials into a function invocation context |[optional] 
**UiEditable** | **bool** | Whether the Service resource's properties and subresources can be edited via the UI |[optional] 
**DomainBase** | **string** | The base domain name for this Service, which is a combination of the unique name and a randomly generated string |[optional] 
**DateCreated** | [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the Service resource was created |[optional] 
**DateUpdated** | [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the Service resource was last updated |[optional] 
**Url** | **string** | The absolute URL of the Service resource |[optional] 
**Links** | **map[string]interface{}** | The URLs of the Service's nested resources |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


