# ProxyV1Service

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | Pointer to **string** | The unique string that we created to identify the Service resource. |
**UniqueName** | Pointer to **string** | An application-defined string that uniquely identifies the resource. This value must be 191 characters or fewer in length and be unique. Supports UTF-8 characters. **This value should not have PII.** |
**AccountSid** | Pointer to **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Service resource. |
**ChatInstanceSid** | Pointer to **string** | The SID of the Chat Service Instance managed by Proxy Service. The Chat Service enables Proxy to forward SMS and channel messages to this chat instance. This is a one-to-one relationship. |
**CallbackUrl** | Pointer to **string** | The URL we call when the interaction status changes. |
**DefaultTtl** | Pointer to **int** | The default `ttl` value for Sessions created in the Service. The TTL (time to live) is measured in seconds after the Session's last create or last Interaction. The default value of `0` indicates an unlimited Session length. You can override a Session's default TTL value by setting its `ttl` value. |
**NumberSelectionBehavior** | Pointer to [**string**](ServiceEnumNumberSelectionBehavior.md) |  |
**GeoMatchLevel** | Pointer to [**string**](ServiceEnumGeoMatchLevel.md) |  |
**InterceptCallbackUrl** | Pointer to **string** | The URL we call on each interaction. If we receive a 403 status, we block the interaction; otherwise the interaction continues. |
**OutOfSessionCallbackUrl** | Pointer to **string** | The URL we call when an inbound call or SMS action occurs on a closed or non-existent Session. If your server (or a Twilio [function](https://www.twilio.com/functions)) responds with valid [TwiML](https://www.twilio.com/docs/voice/twiml), we will process it. This means it is possible, for example, to play a message for a call, send an automated text message response, or redirect a call to another Phone Number. See [Out-of-Session Callback Response Guide](https://www.twilio.com/docs/proxy/out-session-callback-response-guide) for more information. |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time in GMT when the resource was created. |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time in GMT when the resource was last updated. |
**Url** | Pointer to **string** | The absolute URL of the Service resource. |
**Links** | Pointer to **map[string]interface{}** | The URLs of resources related to the Service. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


