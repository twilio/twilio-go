# VideoV1RoomRoomParticipant

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** | The SID of the Account that created the resource |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was created |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was last updated |
**Duration** | Pointer to **int32** | Duration of time in seconds the participant was connected |
**EndTime** | Pointer to [**time.Time**](time.Time.md) | The time when the participant disconnected from the room in ISO 8601 format |
**Identity** | Pointer to **string** | The string that identifies the resource's User |
**Links** | Pointer to **map[string]interface{}** | The URLs of related resources |
**RoomSid** | Pointer to **string** | The SID of the participant's room |
**Sid** | Pointer to **string** | The unique string that identifies the resource |
**StartTime** | Pointer to [**time.Time**](time.Time.md) | The time of participant connected to the room in ISO 8601 format |
**Status** | Pointer to **string** | The status of the Participant |
**Url** | Pointer to **string** | The absolute URL of the resource |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


