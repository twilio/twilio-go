# VideoV1RoomTranscriptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ttid** | Pointer to **string** | The unique string that we created to identify the transcriptions resource. |
**AccountSid** | Pointer to **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Room resource. |
**RoomSid** | Pointer to **string** | The SID of the transcriptions's room. |
**Status** | Pointer to [**string**](RoomTranscriptionsEnumStatus.md) |  |
**Identity** | Pointer to **string** | The application-defined string that uniquely identifies the resource's User within a Room. If a client joins with an existing Identity, the existing client is disconnected. See [access tokens](https://www.twilio.com/docs/video/tutorials/user-identity-access-tokens) and [limits](https://www.twilio.com/docs/video/programmable-video-limits) for more info.  |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The date and time in GMT when the resource was created specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The date and time in GMT when the resource was last updated specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. |
**StartTime** | Pointer to [**time.Time**](time.Time.md) | The time of transcriptions connected to the room in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601#UTC) format. |
**EndTime** | Pointer to [**time.Time**](time.Time.md) | The time when the transcriptions disconnected from the room in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601#UTC) format. |
**Duration** | Pointer to **int** | The duration in seconds that the transcriptions were `connected`. Populated only after the transcriptions is `stopped`. |
**Url** | Pointer to **string** | The absolute URL of the resource. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


