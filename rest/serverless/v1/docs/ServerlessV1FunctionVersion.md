# ServerlessV1FunctionVersion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | **string** | The unique string that identifies the Function Version resource |[optional] 
**AccountSid** | **string** | The SID of the Account that created the Function Version resource |[optional] 
**ServiceSid** | **string** | The SID of the Service that the Function Version resource is associated with |[optional] 
**FunctionSid** | **string** | The SID of the Function resource that is the parent of the Function Version resource |[optional] 
**Path** | **string** | The URL-friendly string by which the Function Version resource can be referenced |[optional] 
**Visibility** | Pointer to [**string**](FunctionVersionEnumVisibility.md) |  |
**DateCreated** | [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the Function Version resource was created |[optional] 
**Url** | **string** | The absolute URL of the Function Version resource |[optional] 
**Links** | **map[string]interface{}** |  |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


