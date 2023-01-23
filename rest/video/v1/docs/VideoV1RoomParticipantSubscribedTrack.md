# VideoV1RoomParticipantSubscribedTrack

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | Pointer to **string** | The unique string that we created to identify the RoomParticipantSubscribedTrack resource. |
**ParticipantSid** | Pointer to **string** | The SID of the participant that subscribes to the track. |
**PublisherSid** | Pointer to **string** | The SID of the participant that publishes the track. |
**RoomSid** | Pointer to **string** | The SID of the room where the track is published. |
**Name** | Pointer to **string** | The track name. Must have no more than 128 characters and be unique among the participant's published tracks. |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The date and time in GMT when the resource was created specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The date and time in GMT when the resource was last updated specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. |
**Enabled** | Pointer to **bool** | Whether the track is enabled. |
**Kind** | Pointer to [**string**](RoomParticipantSubscribedTrackEnumKind.md) |  |
**Url** | Pointer to **string** | The absolute URL of the resource. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


