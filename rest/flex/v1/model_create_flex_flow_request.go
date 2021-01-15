/*
 * Twilio - Flex
 *
 * This is the public Twilio REST API.
 *
 * API version: 5.15.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// CreateFlexFlowRequest struct for CreateFlexFlowRequest
type CreateFlexFlowRequest struct {
	// The channel type. Can be: `web`, `facebook`, `sms`, `whatsapp`, `line` or `custom`.
	ChannelType string `json:"ChannelType"`
	// The SID of the chat service.
	ChatServiceSid string `json:"ChatServiceSid"`
	// The channel contact's Identity.
	ContactIdentity string `json:"ContactIdentity,omitempty"`
	// Whether the new Flex Flow is enabled.
	Enabled bool `json:"Enabled,omitempty"`
	// A descriptive string that you create to describe the Flex Flow resource.
	FriendlyName string `json:"FriendlyName"`
	// The Task Channel for the TaskRouter Task that will be created. Applicable and required when integrationType is `task`. Set to `sms` for SMS, and to `chat` otherwise. The default value is `default`
	IntegrationChannel string `json:"Integration.Channel,omitempty"`
	// In the context of outbound messaging, defines whether to create a Task immediately (and therefore reserve the conversation to current agent), or delay Task creation until the customer sends the first response. Set to false to create immediately, true to delay Task creation. This setting is only applicable for outbound messaging.
	IntegrationCreationOnMessage bool `json:"Integration.CreationOnMessage,omitempty"`
	// The SID of the Studio Flow. Required when `integrationType` is `studio`.
	IntegrationFlowSid string `json:"Integration.FlowSid,omitempty"`
	// The Task priority of a new Task. The default priority is 0. Optional when `integrationType` is `task`, not applicable otherwise.
	IntegrationPriority int32 `json:"Integration.Priority,omitempty"`
	// The number of times to retry the webhook if the first attempt fails. Can be an integer between 0 and 3 (included), default is 0. Optional when integrationType is `external`, not applicable otherwise.
	IntegrationRetryCount int32 `json:"Integration.RetryCount,omitempty"`
	// The Task timeout in seconds for a new Task. Default is 86,400 seconds (24 hours). Optional when `integrationType` is `task`, not applicable otherwise.
	IntegrationTimeout int32 `json:"Integration.Timeout,omitempty"`
	// The URL of the external webhook. Required when `integrationType` is `external`.
	IntegrationUrl string `json:"Integration.Url,omitempty"`
	// The Workflow SID for a new Task. Required when `integrationType` is `task`.
	IntegrationWorkflowSid string `json:"Integration.WorkflowSid,omitempty"`
	// The Workspace SID for a new Task. Required when `integrationType` is `task`.
	IntegrationWorkspaceSid string `json:"Integration.WorkspaceSid,omitempty"`
	// The integration type. Can be: `studio`, `external`, or `task`.
	IntegrationType string `json:"IntegrationType,omitempty"`
	// When enabled, the Messaging Channel Janitor will remove active Proxy sessions if the associated Task is deleted outside of the Flex UI. Defaults to `false`.
	JanitorEnabled bool `json:"JanitorEnabled,omitempty"`
	// When enabled, Flex will keep the chat channel active so that it may be used for subsequent interactions with a contact identity. Defaults to `false`.
	LongLived bool `json:"LongLived,omitempty"`
}
