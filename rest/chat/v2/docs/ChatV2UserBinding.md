# ChatV2UserBinding

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | Pointer to **string** | The unique string that identifies the resource |
**AccountSid** | Pointer to **string** | The SID of the Account that created the resource |
**ServiceSid** | Pointer to **string** | The SID of the Service that the resource is associated with |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was created |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was last updated |
**Endpoint** | Pointer to **string** | The unique endpoint identifier for the User Binding |
**Identity** | Pointer to **string** | The string that identifies the resource's User |
**UserSid** | Pointer to **string** | The SID of the User with the binding |
**CredentialSid** | Pointer to **string** | The SID of the Credential for the binding |
**BindingType** | Pointer to [**string**](UserBindingEnumBindingType.md) |  |
**MessageTypes** | Pointer to **[]string** | The Programmable Chat message types the binding is subscribed to |
**Url** | Pointer to **string** | The absolute URL of the User Binding resource |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


