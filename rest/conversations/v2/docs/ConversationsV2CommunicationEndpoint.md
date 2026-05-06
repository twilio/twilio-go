# ConversationsV2CommunicationEndpoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ParticipantId** | **string** | Participant ID to resolve address from. When provided, Conversations looks up the participant's registered addresses and selects based on channel.  |[optional] 
**Address** | **string** | Explicit address formatted according to channel type: - SMS/VOICE: E.164 phone number (such as \"+18005550100\") - WHATSAPP: Phone number with whatsapp prefix (such as \"whatsapp:+18005550100\") - RCS: Sender ID or phone number with rcs prefix (such as \"rcs:brand_acme_agent\" or \"rcs:+18005550100\") - CHAT: Customer-defined string identifier  |[optional] 
**Channel** | **string** | Channel type. Required when participantId has multiple addresses or when using explicit address. |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


