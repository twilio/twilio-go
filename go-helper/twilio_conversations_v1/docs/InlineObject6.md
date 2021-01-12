# InlineObject6

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attributes** | **string** | An optional string metadata field you can use to store any data you wish. The string value must contain structurally valid JSON if specified.  **Note** that if the attributes are not set \&quot;{}\&quot; will be returned. | [optional] 
**DateCreated** | [**time.Time**](time.Time.md) | The date that this resource was created. | [optional] 
**DateUpdated** | [**time.Time**](time.Time.md) | The date that this resource was last updated. | [optional] 
**Identity** | **string** | A unique string identifier for the conversation participant as [Conversation User](https://www.twilio.com/docs/conversations/api/user-resource). This parameter is non-null if (and only if) the participant is using the Conversations SDK to communicate. Limited to 256 characters. | [optional] 
**LastReadMessageIndex** | Pointer to **int32** | Index of last “read” message in the [Conversation](https://www.twilio.com/docs/conversations/api/conversation-resource) for the Participant. | [optional] 
**LastReadTimestamp** | **string** | Timestamp of last “read” message in the [Conversation](https://www.twilio.com/docs/conversations/api/conversation-resource) for the Participant. | [optional] 
**MessagingBindingProjectedAddress** | **string** | The address of the Twilio phone number that is used in Group MMS. &#39;null&#39; value will remove it. | [optional] 
**MessagingBindingProxyAddress** | **string** | The address of the Twilio phone number that the participant is in contact with. &#39;null&#39; value will remove it. | [optional] 
**RoleSid** | **string** | The SID of a conversation-level [Role](https://www.twilio.com/docs/conversations/api/role-resource) to assign to the participant. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


