# ChatV2Invite

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | Pointer to **string** | The unique string that we created to identify the Invite resource. |
**AccountSid** | Pointer to **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Invite resource. |
**ChannelSid** | Pointer to **string** | The SID of the [Channel](https://www.twilio.com/docs/chat/channels) the Invite resource belongs to. |
**ServiceSid** | Pointer to **string** | The SID of the [Service](https://www.twilio.com/docs/chat/rest/service-resource) the Invite resource is associated with. |
**Identity** | Pointer to **string** | The application-defined string that uniquely identifies the resource's [User](https://www.twilio.com/docs/chat/rest/user-resource) within the [Service](https://www.twilio.com/docs/chat/rest/service-resource). See [access tokens](https://www.twilio.com/docs/chat/create-tokens) for more info. |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The date and time in GMT when the resource was created specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The date and time in GMT when the resource was last updated specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. |
**RoleSid** | Pointer to **string** | The SID of the [Role](https://www.twilio.com/docs/chat/rest/role-resource) assigned to the resource. |
**CreatedBy** | Pointer to **string** | The `identity` of the User that created the invite. |
**Url** | Pointer to **string** | The absolute URL of the Invite resource. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


