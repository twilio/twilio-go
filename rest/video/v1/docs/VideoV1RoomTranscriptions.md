# VideoV1RoomTranscriptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ttid** | Pointer to **string** | The unique string that we created to identify the transcriptions resource. |
**AccountSid** | Pointer to **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Room resource. |
**RoomSid** | Pointer to **string** | The SID of the transcriptions's room. |
**Status** | Pointer to [**string**](RoomTranscriptionsEnumStatus.md) |  |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The date and time in GMT when the resource was created specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The date and time in GMT when the resource was last updated specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. |
**StartTime** | Pointer to [**time.Time**](time.Time.md) | The time of transcriptions connected to the room in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601#UTC) format. |
**EndTime** | Pointer to [**time.Time**](time.Time.md) | The time when the transcriptions disconnected from the room in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601#UTC) format. |
**Duration** | Pointer to **int** | The duration in seconds that the transcriptions were `connected`. Populated only after the transcriptions is `stopped`. |
**Url** | Pointer to **string** | The absolute URL of the resource. |
**Configuration** | Pointer to **map[string]interface{}** | An JSON object that describes the video layout of the composition in terms of regions. See [Specifying Video Layouts](https://www.twilio.com/docs/video/api/compositions-resource#specifying-video-layouts) for more info. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


