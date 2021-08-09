# ConversationsV1ServiceBinding

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** | The unique ID of the Account responsible for this binding. |
**BindingType** | Pointer to **string** | The push technology to use for the binding. |
**ChatServiceSid** | Pointer to **string** | The SID of the Conversation Service that the resource is associated with. |
**CredentialSid** | Pointer to **string** | The SID of the Credential for the binding. |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The date that this resource was created. |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The date that this resource was last updated. |
**Endpoint** | Pointer to **string** | The unique endpoint identifier for the Binding. |
**Identity** | Pointer to **string** | The identity of Conversation User associated with this binding. |
**MessageTypes** | Pointer to **[]string** | The Conversation message types the binding is subscribed to. |
**Sid** | Pointer to **string** | A 34 character string that uniquely identifies this resource. |
**Url** | Pointer to **string** | An absolute URL for this binding. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


