# ConversationsV2SendAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** | The address value formatted according to channel type: - SMS: E.164 phone number (such as \"+18005550100\") - WHATSAPP: Phone number with whatsapp prefix (such as \"whatsapp:+18005550100\") - RCS: Sender ID or phone number with rcs prefix (such as \"rcs:brand_acme_agent\" or \"rcs:+18005550100\") - CHAT: Customer-defined string identifier  |
**Channel** | **string** | Channel type for sending communications. Note: VOICE is receive-only and not supported for send operations. |
**ParticipantId** | **string** | Optional Participant ID. If omitted, the system will resolve or create the participant based on address and channel. |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


