# NotifyV1User

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | Pointer to **string** | The unique string that we created to identify the User resource. |
**AccountSid** | Pointer to **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the User resource. |
**ServiceSid** | Pointer to **string** | The SID of the [Service](https://www.twilio.com/docs/notify/api/service-resource) the resource is associated with. |
**Identity** | Pointer to **string** | The application-defined string that uniquely identifies the resource's User within the Service. See [access tokens](https://www.twilio.com/docs/chat/create-tokens) for more info. |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The date and time in GMT when the resource was created specified in [RFC 2822](https://www.ietf.org/rfc/rfc2822.txt) format. |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The date and time in GMT when the resource was last updated specified in [RFC 2822](https://www.ietf.org/rfc/rfc2822.txt) format. |
**Segments** | Pointer to **[]string** | Deprecated. |
**Url** | Pointer to **string** | The absolute URL of the User resource. |
**Links** | Pointer to **map[string]interface{}** | The absolute URLs of related resources. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


