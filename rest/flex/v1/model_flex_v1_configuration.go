/*
 * Twilio - Flex
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.0.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"time"
)

// FlexV1Configuration struct for FlexV1Configuration
type FlexV1Configuration struct {
	AccountSid                   *string                   `json:"account_sid,omitempty"`
	Attributes                   *map[string]interface{}   `json:"attributes,omitempty"`
	CallRecordingEnabled         *bool                     `json:"call_recording_enabled,omitempty"`
	CallRecordingWebhookUrl      *string                   `json:"call_recording_webhook_url,omitempty"`
	ChatServiceInstanceSid       *string                   `json:"chat_service_instance_sid,omitempty"`
	CrmAttributes                *map[string]interface{}   `json:"crm_attributes,omitempty"`
	CrmCallbackUrl               *string                   `json:"crm_callback_url,omitempty"`
	CrmEnabled                   *bool                     `json:"crm_enabled,omitempty"`
	CrmFallbackUrl               *string                   `json:"crm_fallback_url,omitempty"`
	CrmType                      *string                   `json:"crm_type,omitempty"`
	DateCreated                  *time.Time                `json:"date_created,omitempty"`
	DateUpdated                  *time.Time                `json:"date_updated,omitempty"`
	FlexServiceInstanceSid       *string                   `json:"flex_service_instance_sid,omitempty"`
	Integrations                 *[]map[string]interface{} `json:"integrations,omitempty"`
	Markdown                     *map[string]interface{}   `json:"markdown,omitempty"`
	MessagingServiceInstanceSid  *string                   `json:"messaging_service_instance_sid,omitempty"`
	Notifications                *map[string]interface{}   `json:"notifications,omitempty"`
	OutboundCallFlows            *map[string]interface{}   `json:"outbound_call_flows,omitempty"`
	PluginServiceAttributes      *map[string]interface{}   `json:"plugin_service_attributes,omitempty"`
	PluginServiceEnabled         *bool                     `json:"plugin_service_enabled,omitempty"`
	PublicAttributes             *map[string]interface{}   `json:"public_attributes,omitempty"`
	QueueStatsConfiguration      *map[string]interface{}   `json:"queue_stats_configuration,omitempty"`
	RuntimeDomain                *string                   `json:"runtime_domain,omitempty"`
	ServerlessServiceSids        *[]string                 `json:"serverless_service_sids,omitempty"`
	ServiceVersion               *string                   `json:"service_version,omitempty"`
	Status                       *ConfigurationStatus      `json:"status,omitempty"`
	TaskrouterOfflineActivitySid *string                   `json:"taskrouter_offline_activity_sid,omitempty"`
	TaskrouterSkills             *[]map[string]interface{} `json:"taskrouter_skills,omitempty"`
	TaskrouterTargetTaskqueueSid *string                   `json:"taskrouter_target_taskqueue_sid,omitempty"`
	TaskrouterTargetWorkflowSid  *string                   `json:"taskrouter_target_workflow_sid,omitempty"`
	TaskrouterTaskqueues         *[]map[string]interface{} `json:"taskrouter_taskqueues,omitempty"`
	TaskrouterWorkerAttributes   *map[string]interface{}   `json:"taskrouter_worker_attributes,omitempty"`
	TaskrouterWorkerChannels     *map[string]interface{}   `json:"taskrouter_worker_channels,omitempty"`
	TaskrouterWorkspaceSid       *string                   `json:"taskrouter_workspace_sid,omitempty"`
	UiAttributes                 *map[string]interface{}   `json:"ui_attributes,omitempty"`
	UiDependencies               *map[string]interface{}   `json:"ui_dependencies,omitempty"`
	UiLanguage                   *string                   `json:"ui_language,omitempty"`
	UiVersion                    *string                   `json:"ui_version,omitempty"`
	Url                          *string                   `json:"url,omitempty"`
}
