/*
 * Twilio - Flex
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.8.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
import (
	"time"
)
// FlexV1Configuration struct for FlexV1Configuration
type FlexV1Configuration struct {
	AccountSid string `json:"AccountSid,omitempty"`
	Attributes map[string]interface{} `json:"Attributes,omitempty"`
	CallRecordingEnabled bool `json:"CallRecordingEnabled,omitempty"`
	CallRecordingWebhookUrl string `json:"CallRecordingWebhookUrl,omitempty"`
	ChatServiceInstanceSid string `json:"ChatServiceInstanceSid,omitempty"`
	CrmAttributes map[string]interface{} `json:"CrmAttributes,omitempty"`
	CrmCallbackUrl string `json:"CrmCallbackUrl,omitempty"`
	CrmEnabled bool `json:"CrmEnabled,omitempty"`
	CrmFallbackUrl string `json:"CrmFallbackUrl,omitempty"`
	CrmType string `json:"CrmType,omitempty"`
	DateCreated time.Time `json:"DateCreated,omitempty"`
	DateUpdated time.Time `json:"DateUpdated,omitempty"`
	FlexServiceInstanceSid string `json:"FlexServiceInstanceSid,omitempty"`
	Integrations []map[string]interface{} `json:"Integrations,omitempty"`
	MessagingServiceInstanceSid string `json:"MessagingServiceInstanceSid,omitempty"`
	OutboundCallFlows map[string]interface{} `json:"OutboundCallFlows,omitempty"`
	PluginServiceAttributes map[string]interface{} `json:"PluginServiceAttributes,omitempty"`
	PluginServiceEnabled bool `json:"PluginServiceEnabled,omitempty"`
	PublicAttributes map[string]interface{} `json:"PublicAttributes,omitempty"`
	QueueStatsConfiguration map[string]interface{} `json:"QueueStatsConfiguration,omitempty"`
	RuntimeDomain string `json:"RuntimeDomain,omitempty"`
	ServerlessServiceSids []string `json:"ServerlessServiceSids,omitempty"`
	ServiceVersion string `json:"ServiceVersion,omitempty"`
	Status Status `json:"Status,omitempty"`
	TaskrouterOfflineActivitySid string `json:"TaskrouterOfflineActivitySid,omitempty"`
	TaskrouterSkills []map[string]interface{} `json:"TaskrouterSkills,omitempty"`
	TaskrouterTargetTaskqueueSid string `json:"TaskrouterTargetTaskqueueSid,omitempty"`
	TaskrouterTargetWorkflowSid string `json:"TaskrouterTargetWorkflowSid,omitempty"`
	TaskrouterTaskqueues []map[string]interface{} `json:"TaskrouterTaskqueues,omitempty"`
	TaskrouterWorkerAttributes map[string]interface{} `json:"TaskrouterWorkerAttributes,omitempty"`
	TaskrouterWorkerChannels map[string]interface{} `json:"TaskrouterWorkerChannels,omitempty"`
	TaskrouterWorkspaceSid string `json:"TaskrouterWorkspaceSid,omitempty"`
	UiAttributes map[string]interface{} `json:"UiAttributes,omitempty"`
	UiDependencies map[string]interface{} `json:"UiDependencies,omitempty"`
	UiLanguage string `json:"UiLanguage,omitempty"`
	UiVersion string `json:"UiVersion,omitempty"`
	Url string `json:"Url,omitempty"`
}
