# MediaV1PlayerStreamerPlaybackGrant

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | Pointer to **string** | The unique string generated to identify the PlayerStreamer resource that this PlaybackGrant authorizes views for. |
**Url** | Pointer to **string** | The absolute URL of the resource. |
**AccountSid** | Pointer to **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created this resource. |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The date and time in GMT when the resource was created specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. |
**Grant** | Pointer to **interface{}** | The grant that authorizes the player sdk to connect to the livestream |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


