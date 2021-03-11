# NotifyV1ServiceBinding

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** | The SID of the Account that created the resource |
**Address** | Pointer to **string** | The channel-specific address |
**BindingType** | Pointer to **string** | The type of the Binding |
**CredentialSid** | Pointer to **string** | The SID of the Credential resource to be used to send notifications to this Binding |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The RFC 2822 date and time in GMT when the resource was created |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The RFC 2822 date and time in GMT when the resource was last updated |
**Endpoint** | Pointer to **string** | Deprecated |
**Identity** | Pointer to **string** | The `identity` value that identifies the new resource's User |
**Links** | Pointer to **map[string]interface{}** | The URLs of related resources |
**NotificationProtocolVersion** | Pointer to **string** | The protocol version to use to send the notification |
**ServiceSid** | Pointer to **string** | The SID of the Service that the resource is associated with |
**Sid** | Pointer to **string** | The unique string that identifies the resource |
**Tags** | Pointer to **[]string** | The list of tags associated with this Binding |
**Url** | Pointer to **string** | The absolute URL of the Binding resource |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


