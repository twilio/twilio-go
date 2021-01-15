# CreateServiceConversationParticipantRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attributes** | **string** | An optional string metadata field you can use to store any data you wish. The string value must contain structurally valid JSON if specified.  **Note** that if the attributes are not set \&quot;{}\&quot; will be returned. | [optional] 
**DateCreated** | [**time.Time**](time.Time.md) | The date that this resource was created. | [optional] 
**DateUpdated** | [**time.Time**](time.Time.md) | The date that this resource was last updated. | [optional] 
**Identity** | **string** | A unique string identifier for the conversation participant as [Conversation User](https://www.twilio.com/docs/conversations/api/user-resource). This parameter is non-null if (and only if) the participant is using the Conversation SDK to communicate. Limited to 256 characters. | [optional] 
**MessagingBindingAddress** | **string** | The address of the participant&#39;s device, e.g. a phone or WhatsApp number. Together with the Proxy address, this determines a participant uniquely. This field (with proxy_address) is only null when the participant is interacting from an SDK endpoint (see the &#39;identity&#39; field). | [optional] 
**MessagingBindingProjectedAddress** | **string** | The address of the Twilio phone number that is used in Group MMS. Communication mask for the Conversation participant with Identity. | [optional] 
**MessagingBindingProxyAddress** | **string** | The address of the Twilio phone number (or WhatsApp number) that the participant is in contact with. This field, together with participant address, is only null when the participant is interacting from an SDK endpoint (see the &#39;identity&#39; field). | [optional] 
**RoleSid** | **string** | The SID of a conversation-level [Role](https://www.twilio.com/docs/conversations/api/role-resource) to assign to the participant. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


