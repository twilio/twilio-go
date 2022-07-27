# NotifyV1Binding

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | Pointer to **string** | The unique string that identifies the resource |
**AccountSid** | Pointer to **string** | The SID of the Account that created the resource |
**ServiceSid** | Pointer to **string** | The SID of the Service that the resource is associated with |
**CredentialSid** | Pointer to **string** | The SID of the Credential resource to be used to send notifications to this Binding |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The RFC 2822 date and time in GMT when the resource was created |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The RFC 2822 date and time in GMT when the resource was last updated |
**NotificationProtocolVersion** | Pointer to **string** | The protocol version to use to send the notification |
**Endpoint** | Pointer to **string** | Deprecated |
**Identity** | Pointer to **string** | The `identity` value that identifies the new resource's User |
**BindingType** | Pointer to **string** | The type of the Binding |
**Address** | Pointer to **string** | The channel-specific address |
**Tags** | Pointer to **[]string** | The list of tags associated with this Binding |
**Url** | Pointer to **string** | The absolute URL of the Binding resource |
**Links** | Pointer to **map[string]interface{}** | The URLs of related resources |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


