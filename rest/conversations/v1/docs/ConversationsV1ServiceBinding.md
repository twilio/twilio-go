# ConversationsV1ServiceBinding

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | **string** | A 34 character string that uniquely identifies this resource. |[optional] 
**AccountSid** | **string** | The unique ID of the Account responsible for this binding. |[optional] 
**ChatServiceSid** | **string** | The SID of the Conversation Service that the resource is associated with. |[optional] 
**CredentialSid** | **string** | The SID of the Credential for the binding. |[optional] 
**DateCreated** | [**time.Time**](time.Time.md) | The date that this resource was created. |[optional] 
**DateUpdated** | [**time.Time**](time.Time.md) | The date that this resource was last updated. |[optional] 
**Endpoint** | **string** | The unique endpoint identifier for the Binding. |[optional] 
**Identity** | **string** | The identity of Conversation User associated with this binding. |[optional] 
**BindingType** | Pointer to [**string**](ServiceBindingEnumBindingType.md) |  |
**MessageTypes** | **[]string** | The Conversation message types the binding is subscribed to. |[optional] 
**Url** | **string** | An absolute URL for this binding. |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


