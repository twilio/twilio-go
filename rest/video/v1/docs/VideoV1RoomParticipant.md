# VideoV1RoomParticipant

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | Pointer to **string** | The unique string that we created to identify the RoomParticipant resource. |
**RoomSid** | Pointer to **string** | The SID of the participant's room. |
**AccountSid** | Pointer to **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the RoomParticipant resource. |
**Status** | Pointer to [**string**](RoomParticipantEnumStatus.md) |  |
**Identity** | Pointer to **string** | The application-defined string that uniquely identifies the resource's User within a Room. If a client joins with an existing Identity, the existing client is disconnected. See [access tokens](https://www.twilio.com/docs/video/tutorials/user-identity-access-tokens) and [limits](https://www.twilio.com/docs/video/programmable-video-limits) for more info.  |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The date and time in GMT when the resource was created specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The date and time in GMT when the resource was last updated specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. |
**StartTime** | Pointer to [**time.Time**](time.Time.md) | The time of participant connected to the room in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601#UTC) format. |
**EndTime** | Pointer to [**time.Time**](time.Time.md) | The time when the participant disconnected from the room in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601#UTC) format. |
**Duration** | Pointer to **int** | The duration in seconds that the participant was `connected`. Populated only after the participant is `disconnected`. |
**Url** | Pointer to **string** | The absolute URL of the resource. |
**Links** | Pointer to **map[string]interface{}** | The URLs of related resources. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


