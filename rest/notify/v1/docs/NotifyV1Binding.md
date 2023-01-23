# NotifyV1Binding

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | **string** | The unique string that identifies the resource |[optional] 
**AccountSid** | **string** | The SID of the Account that created the resource |[optional] 
**ServiceSid** | **string** | The SID of the Service that the resource is associated with |[optional] 
**CredentialSid** | **string** | The SID of the Credential resource to be used to send notifications to this Binding |[optional] 
**DateCreated** | [**time.Time**](time.Time.md) | The RFC 2822 date and time in GMT when the resource was created |[optional] 
**DateUpdated** | [**time.Time**](time.Time.md) | The RFC 2822 date and time in GMT when the resource was last updated |[optional] 
**NotificationProtocolVersion** | **string** | The protocol version to use to send the notification |[optional] 
**Endpoint** | **string** | Deprecated |[optional] 
**Identity** | **string** | The `identity` value that identifies the new resource's User |[optional] 
**BindingType** | **string** | The type of the Binding |[optional] 
**Address** | **string** | The channel-specific address |[optional] 
**Tags** | **[]string** | The list of tags associated with this Binding |[optional] 
**Url** | **string** | The absolute URL of the Binding resource |[optional] 
**Links** | **map[string]interface{}** | The URLs of related resources |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


