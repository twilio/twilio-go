# VideoV1RoomParticipantPublishedTrack

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | Pointer to **string** | The unique string that we created to identify the RoomParticipantPublishedTrack resource. |
**ParticipantSid** | Pointer to **string** | The SID of the Participant resource with the published track. |
**RoomSid** | Pointer to **string** | The SID of the Room resource where the track is published. |
**Name** | Pointer to **string** | The track name. Must be no more than 128 characters, and be unique among the participant's published tracks. |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The date and time in GMT when the resource was created specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The date and time in GMT when the resource was last updated specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. |
**Enabled** | Pointer to **bool** | Whether the track is enabled. |
**Kind** | Pointer to [**string**](RoomParticipantPublishedTrackEnumKind.md) |  |
**Url** | Pointer to **string** | The absolute URL of the resource. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


