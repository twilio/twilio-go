# RoutesV3Rcs

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AgentId** | Pointer to **string** | The RCS agent ID, in the format `agent_name_12312321`. |
**Url** | Pointer to **string** | The absolute URL of the resource. |
**AccountSid** | Pointer to **string** | The unique SID identifier of the Account. |
**FriendlyName** | Pointer to **string** | A human readable description of the Inbound Processing Region assignment for this RCS agent, up to 64 characters. |
**MessagingRegion** | Pointer to **string** | The Inbound Processing Region used for this RCS agent for messaging. |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The date that this RCS agent was assigned an Inbound Processing Region, given in ISO 8601 format. |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The date that the Inbound Processing Region was updated for this RCS agent, given in ISO 8601 format. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


