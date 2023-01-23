# ProxyV1Session

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | Pointer to **string** | The unique string that we created to identify the Session resource. |
**ServiceSid** | Pointer to **string** | The SID of the [Service](https://www.twilio.com/docs/proxy/api/service) the session is associated with. |
**AccountSid** | Pointer to **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Session resource. |
**DateStarted** | Pointer to [**time.Time**](time.Time.md) | The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date when the Session started. |
**DateEnded** | Pointer to [**time.Time**](time.Time.md) | The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date when the Session ended. |
**DateLastInteraction** | Pointer to [**time.Time**](time.Time.md) | The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date when the Session last had an interaction. |
**DateExpiry** | Pointer to [**time.Time**](time.Time.md) | The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date when the Session should expire. If this is value is present, it overrides the `ttl` value. |
**UniqueName** | Pointer to **string** | An application-defined string that uniquely identifies the resource. This value must be 191 characters or fewer in length and be unique. Supports UTF-8 characters. **This value should not have PII.** |
**Status** | Pointer to [**string**](SessionEnumStatus.md) |  |
**ClosedReason** | Pointer to **string** | The reason the Session ended. |
**Ttl** | Pointer to **int** | The time, in seconds, when the session will expire. The time is measured from the last Session create or the Session's last Interaction. |
**Mode** | Pointer to [**string**](SessionEnumMode.md) |  |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time in GMT when the resource was created. |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time in GMT when the resource was last updated. |
**Url** | Pointer to **string** | The absolute URL of the Session resource. |
**Links** | Pointer to **map[string]interface{}** | The URLs of resources related to the Session. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


