# ConversationsV1ServiceConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChatServiceSid** | Pointer to **string** | The unique string that we created to identify the Service configuration resource. |
**DefaultConversationCreatorRoleSid** | Pointer to **string** | The conversation-level role assigned to a conversation creator user when they join a new conversation. See the [Conversation Role](https://www.twilio.com/docs/conversations/api/role-resource) for more info about roles. |
**DefaultConversationRoleSid** | Pointer to **string** | The conversation-level role assigned to users when they are added to a conversation. See the [Conversation Role](https://www.twilio.com/docs/conversations/api/role-resource) for more info about roles. |
**DefaultChatServiceRoleSid** | Pointer to **string** | The service-level role assigned to users when they are added to the service. See the [Conversation Role](https://www.twilio.com/docs/conversations/api/role-resource) for more info about roles. |
**Url** | Pointer to **string** | An absolute API resource URL for this service configuration. |
**Links** | Pointer to **map[string]interface{}** | Contains an absolute API resource URL to access the push notifications configuration of this service. |
**ReachabilityEnabled** | Pointer to **bool** | Whether the [Reachability Indicator](https://www.twilio.com/docs/chat/reachability-indicator) is enabled for this Conversations Service. The default is `false`. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


