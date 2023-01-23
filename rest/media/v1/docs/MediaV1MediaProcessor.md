# MediaV1MediaProcessor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the MediaProcessor resource. |
**Sid** | Pointer to **string** | The unique string generated to identify the MediaProcessor resource. |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The date and time in GMT when the resource was created specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The date and time in GMT when the resource was last updated specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. |
**Extension** | Pointer to **string** | The [Media Extension](/docs/live/api/media-extensions-overview) name or URL. Ex: `video-composer-v2` |
**ExtensionContext** | Pointer to **string** | The context of the Media Extension, represented as a JSON dictionary. See the documentation for the specific [Media Extension](/docs/live/api/media-extensions-overview) you are using for more information about the context to send. |
**Status** | Pointer to [**string**](MediaProcessorEnumStatus.md) |  |
**Url** | Pointer to **string** | The absolute URL of the resource. |
**EndedReason** | Pointer to **string** | The reason why a MediaProcessor ended. When a MediaProcessor is in progress, will be `null`. When a MediaProcessor is completed, can be `ended-via-api`, `max-duration-exceeded`, `error-loading-extension`, `error-streaming-media` or `internal-service-error`. See [ended reasons](/docs/live/api/mediaprocessors#mediaprocessor-ended-reason-values) for more details. |
**StatusCallback** | Pointer to **string** | The URL to which Twilio will send asynchronous webhook requests for every MediaProcessor event. See [Status Callbacks](/docs/live/status-callbacks) for details. |
**StatusCallbackMethod** | Pointer to **string** | The HTTP method Twilio should use to call the `status_callback` URL. Can be `POST` or `GET` and the default is `POST`. |
**MaxDuration** | Pointer to **int** | The maximum time, in seconds, that the MediaProcessor can run before automatically ends. The default value is 300 seconds, and the maximum value is 90000 seconds. Once this maximum duration is reached, Twilio will end the MediaProcessor, regardless of whether media is still streaming. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


