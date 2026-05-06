# ConversationsV2Participant

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Participant ID. |
**ConversationId** | **string** | Conversation ID. |
**AccountId** | **string** | Account ID. |
**Name** | **string** | Participant display name. |
**Type** | **string** | Type of Participant in the Conversation. |[optional] 
**ProfileId** | **string** | Profile ID. Note: This field is only resolved for `CUSTOMER` participant types, not for `HUMAN_AGENT` or `AI_AGENT` participants. |[optional] 
**Addresses** | [**[]ConversationsV2Address**](ConversationsV2Address.md) | Communication addresses for this Participant. Address format varies by channel: - SMS/VOICE: E.164 phone number (such as \"+18005550100\") - EMAIL: Email address (such as \"user@example.com\") - WHATSAPP: Phone number with whatsapp prefix (such as \"whatsapp:+18005550100\") - RCS: Sender ID or phone number with rcs prefix (such as \"rcs:brand_acme_agent\" or \"rcs:+18005550100\")  |[optional] 
**CreatedAt** | [**time.Time**](time.Time.md) | Timestamp when this Participant was created. |[optional] 
**UpdatedAt** | [**time.Time**](time.Time.md) | Timestamp when this Participant was last updated. |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


