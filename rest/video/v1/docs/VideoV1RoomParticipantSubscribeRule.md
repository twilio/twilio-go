# VideoV1RoomParticipantSubscribeRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ParticipantSid** | Pointer to **string** | The SID of the Participant resource for the Subscribe Rules. |
**RoomSid** | Pointer to **string** | The SID of the Room resource for the Subscribe Rules |
**Rules** | Pointer to [**[]VideoV1RoomRoomParticipantRoomParticipantSubscribeRuleRules**](VideoV1RoomRoomParticipantRoomParticipantSubscribeRuleRules.md) | A collection of Subscribe Rules that describe how to include or exclude matching tracks. See the [Specifying Subscribe Rules](https://www.twilio.com/docs/video/api/track-subscriptions#specifying-sr) section for further information. |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The date and time in GMT when the resource was created specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The date and time in GMT when the resource was last updated specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


