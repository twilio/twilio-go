# ChatV2UserBinding

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | **string** | The unique string that identifies the resource |[optional] 
**AccountSid** | **string** | The SID of the Account that created the resource |[optional] 
**ServiceSid** | **string** | The SID of the Service that the resource is associated with |[optional] 
**DateCreated** | [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was created |[optional] 
**DateUpdated** | [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was last updated |[optional] 
**Endpoint** | **string** | The unique endpoint identifier for the User Binding |[optional] 
**Identity** | **string** | The string that identifies the resource's User |[optional] 
**UserSid** | **string** | The SID of the User with the binding |[optional] 
**CredentialSid** | **string** | The SID of the Credential for the binding |[optional] 
**BindingType** | Pointer to [**string**](UserBindingEnumBindingType.md) |  |
**MessageTypes** | **[]string** | The Programmable Chat message types the binding is subscribed to |[optional] 
**Url** | **string** | The absolute URL of the User Binding resource |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


