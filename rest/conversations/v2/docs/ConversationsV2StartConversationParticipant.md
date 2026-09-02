# ConversationsV2StartConversationParticipant

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Display name for this Participant, and the handle the action's `from` and `to` select it by. Must be unique within the request. |
**Type** | [**ConversationsV2ParticipantType**](ConversationsV2ParticipantType.md) |  |[optional] 
**ProfileId** | **string** | Profile to associate with this Participant, instead of resolving one from the addresses. |[optional] 
**AgentConnectConnectionId** | **string** | Agent Connect connection backing an AI_AGENT Participant. Becomes a path segment of the relay endpoint, so it is constrained to an opaque handle. |[optional] 
**Addresses** | [**[]ConversationsV2StartConversationAddress**](ConversationsV2StartConversationAddress.md) | Channel addresses this Participant can be reached on. |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


