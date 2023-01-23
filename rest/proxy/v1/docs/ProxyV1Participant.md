# ProxyV1Participant

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | Pointer to **string** | The unique string that we created to identify the Participant resource. |
**SessionSid** | Pointer to **string** | The SID of the parent [Session](https://www.twilio.com/docs/proxy/api/session) resource. |
**ServiceSid** | Pointer to **string** | The SID of the resource's parent [Service](https://www.twilio.com/docs/proxy/api/service) resource. |
**AccountSid** | Pointer to **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Participant resource. |
**FriendlyName** | Pointer to **string** | The string that you assigned to describe the participant. This value must be 255 characters or fewer. Supports UTF-8 characters. **This value should not have PII.** |
**Identifier** | Pointer to **string** | The phone number or channel identifier of the Participant. This value must be 191 characters or fewer. Supports UTF-8 characters. |
**ProxyIdentifier** | Pointer to **string** | The phone number or short code (masked number) of the participant's partner. The participant will call or message the partner participant at this number. |
**ProxyIdentifierSid** | Pointer to **string** | The SID of the Proxy Identifier assigned to the Participant. |
**DateDeleted** | Pointer to [**time.Time**](time.Time.md) | The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date when the Participant was removed from the session. |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time in GMT when the resource was created. |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time in GMT when the resource was last updated. |
**Url** | Pointer to **string** | The absolute URL of the Participant resource. |
**Links** | Pointer to **map[string]interface{}** | The URLs to resources related the participant. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


