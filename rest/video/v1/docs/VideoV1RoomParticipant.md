# VideoV1RoomParticipant

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | Pointer to **string** | The unique string that identifies the resource |
**RoomSid** | Pointer to **string** | The SID of the participant's room |
**AccountSid** | Pointer to **string** | The SID of the Account that created the resource |
**Status** | Pointer to [**string**](RoomParticipantEnumStatus.md) |  |
**Identity** | Pointer to **string** | The string that identifies the resource's User |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was created |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was last updated |
**StartTime** | Pointer to [**time.Time**](time.Time.md) | The time of participant connected to the room in ISO 8601 format |
**EndTime** | Pointer to [**time.Time**](time.Time.md) | The time when the participant disconnected from the room in ISO 8601 format |
**Duration** | Pointer to **int** | Duration of time in seconds the participant was connected |
**Url** | Pointer to **string** | The absolute URL of the resource |
**Links** | Pointer to **map[string]interface{}** | The URLs of related resources |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


