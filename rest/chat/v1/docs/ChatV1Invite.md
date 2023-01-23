# ChatV1Invite

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | Pointer to **string** | The unique string that we created to identify the Invite resource. |
**AccountSid** | Pointer to **string** | The SID of the [Account](https://www.twilio.com/docs/api/rest/account) that created the Invite resource. |
**ChannelSid** | Pointer to **string** | The SID of the [Channel](https://www.twilio.com/docs/api/chat/rest/channels) the resource belongs to. |
**ServiceSid** | Pointer to **string** | The SID of the [Service](https://www.twilio.com/docs/api/chat/rest/services) the resource is associated with. |
**Identity** | Pointer to **string** | The application-defined string that uniquely identifies the resource's [User](https://www.twilio.com/docs/api/chat/rest/users) within the [Service](https://www.twilio.com/docs/api/chat/rest/services). See [access tokens](https://www.twilio.com/docs/api/chat/guides/create-tokens) for more info. |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The date and time in GMT when the resource was created specified in [RFC 2822](http://www.ietf.org/rfc/rfc2822.txt) format. |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The date and time in GMT when the resource was last updated specified in [RFC 2822](http://www.ietf.org/rfc/rfc2822.txt) format. |
**RoleSid** | Pointer to **string** | The SID of the [Role](https://www.twilio.com/docs/api/chat/rest/roles) assigned to the resource. |
**CreatedBy** | Pointer to **string** | The `identity` of the User that created the invite. |
**Url** | Pointer to **string** | The absolute URL of the Invite resource. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


