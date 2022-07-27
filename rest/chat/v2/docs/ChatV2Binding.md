# ChatV2Binding

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | Pointer to **string** | The unique string that identifies the resource |
**AccountSid** | Pointer to **string** | The SID of the Account that created the resource |
**ServiceSid** | Pointer to **string** | The SID of the Service that the Binding resource is associated with |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was created |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was last updated |
**Endpoint** | Pointer to **string** | The unique endpoint identifier for the Binding |
**Identity** | Pointer to **string** | The string that identifies the resource's User |
**CredentialSid** | Pointer to **string** | The SID of the Credential for the binding |
**BindingType** | Pointer to [**string**](BindingEnumBindingType.md) |  |
**MessageTypes** | Pointer to **[]string** | The Programmable Chat message types the binding is subscribed to |
**Url** | Pointer to **string** | The absolute URL of the Binding resource |
**Links** | Pointer to **map[string]interface{}** | The absolute URLs of the Binding's User |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


