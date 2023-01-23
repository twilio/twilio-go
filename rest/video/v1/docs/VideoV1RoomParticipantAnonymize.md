# VideoV1RoomParticipantAnonymize

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | **string** | The unique string that identifies the resource |[optional] 
**RoomSid** | **string** | The SID of the participant's room |[optional] 
**AccountSid** | **string** | The SID of the Account that created the resource |[optional] 
**Status** | Pointer to [**string**](RoomParticipantAnonymizeEnumStatus.md) |  |
**Identity** | **string** | The SID of the participant |[optional] 
**DateCreated** | [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was created |[optional] 
**DateUpdated** | [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was last updated |[optional] 
**StartTime** | [**time.Time**](time.Time.md) | The time of participant connected to the room in ISO 8601 format |[optional] 
**EndTime** | [**time.Time**](time.Time.md) | The time when the participant disconnected from the room in ISO 8601 format |[optional] 
**Duration** | Pointer to **int** | Duration of time in seconds the participant was connected |
**Url** | **string** | The absolute URL of the resource |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


