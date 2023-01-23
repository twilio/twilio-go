# MediaV1PlayerStreamer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the PlayerStreamer resource. |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The date and time in GMT when the resource was created specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The date and time in GMT when the resource was last updated specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. |
**Video** | Pointer to **bool** | Specifies whether the PlayerStreamer is configured to stream video. Defaults to `true`. |
**Links** | Pointer to **map[string]interface{}** | The URLs of related resources. |
**Sid** | Pointer to **string** | The unique string generated to identify the PlayerStreamer resource. |
**Status** | Pointer to [**string**](PlayerStreamerEnumStatus.md) |  |
**Url** | Pointer to **string** | The absolute URL of the resource. |
**StatusCallback** | Pointer to **string** | The URL to which Twilio will send asynchronous webhook requests for every PlayerStreamer event. See [Status Callbacks](/docs/live/status-callbacks) for more details. |
**StatusCallbackMethod** | Pointer to **string** | The HTTP method Twilio should use to call the `status_callback` URL. Can be `POST` or `GET` and the default is `POST`. |
**EndedReason** | Pointer to [**string**](PlayerStreamerEnumEndedReason.md) |  |
**MaxDuration** | Pointer to **int** | The maximum time, in seconds, that the PlayerStreamer is active (`created` or `started`) before automatically ends. The default value is 300 seconds, and the maximum value is 90000 seconds. Once this maximum duration is reached, Twilio will end the PlayerStreamer, regardless of whether media is still streaming. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


