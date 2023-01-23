# FlexV1Configuration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | **string** | The SID of the Account that created the resource |[optional] 
**DateCreated** | [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the Configuration resource was created |[optional] 
**DateUpdated** | [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the Configuration resource was last updated |[optional] 
**Attributes** | Pointer to **interface{}** | An object that contains application-specific data |
**Status** | Pointer to [**string**](ConfigurationEnumStatus.md) |  |
**TaskrouterWorkspaceSid** | **string** | The SID of the TaskRouter Workspace |[optional] 
**TaskrouterTargetWorkflowSid** | **string** | The SID of the TaskRouter target Workflow |[optional] 
**TaskrouterTargetTaskqueueSid** | **string** | The SID of the TaskRouter Target TaskQueue |[optional] 
**TaskrouterTaskqueues** | **[]interface{}** | The list of TaskRouter TaskQueues |[optional] 
**TaskrouterSkills** | **[]interface{}** | The Skill description for TaskRouter workers |[optional] 
**TaskrouterWorkerChannels** | Pointer to **interface{}** | The TaskRouter default channel capacities and availability for workers |
**TaskrouterWorkerAttributes** | Pointer to **interface{}** | The TaskRouter Worker attributes |
**TaskrouterOfflineActivitySid** | **string** | The TaskRouter SID of the offline activity |[optional] 
**RuntimeDomain** | **string** | The URL where the Flex instance is hosted |[optional] 
**MessagingServiceInstanceSid** | **string** | The SID of the Messaging service instance |[optional] 
**ChatServiceInstanceSid** | **string** | The SID of the chat service this user belongs to |[optional] 
**FlexServiceInstanceSid** | **string** | The SID of the Flex service instance |[optional] 
**UiLanguage** | **string** | The primary language of the Flex UI |[optional] 
**UiAttributes** | Pointer to **interface{}** | The object that describes Flex UI characteristics and settings |
**UiDependencies** | Pointer to **interface{}** | The object that defines the NPM packages and versions to be used in Hosted Flex |
**UiVersion** | **string** | The Pinned UI version |[optional] 
**ServiceVersion** | **string** | The Flex Service version |[optional] 
**CallRecordingEnabled** | **bool** | Whether call recording is enabled |[optional] 
**CallRecordingWebhookUrl** | **string** | The call recording webhook URL |[optional] 
**CrmEnabled** | **bool** | Whether CRM is present for Flex |[optional] 
**CrmType** | **string** | The CRM Type |[optional] 
**CrmCallbackUrl** | **string** | The CRM Callback URL |[optional] 
**CrmFallbackUrl** | **string** | The CRM Fallback URL |[optional] 
**CrmAttributes** | Pointer to **interface{}** | An object that contains the CRM attributes |
**PublicAttributes** | Pointer to **interface{}** | The list of public attributes |
**PluginServiceEnabled** | **bool** | Whether the plugin service enabled |[optional] 
**PluginServiceAttributes** | Pointer to **interface{}** | The plugin service attributes |
**Integrations** | **[]interface{}** | A list of objects that contain the configurations for the Integrations supported in this configuration |[optional] 
**OutboundCallFlows** | Pointer to **interface{}** | The list of outbound call flows |
**ServerlessServiceSids** | **[]string** | The list of serverless service SIDs |[optional] 
**QueueStatsConfiguration** | Pointer to **interface{}** | Configurable parameters for Queues Statistics |
**Notifications** | Pointer to **interface{}** | Configurable parameters for Notifications |
**Markdown** | Pointer to **interface{}** | Configurable parameters for Markdown |
**Url** | **string** | The absolute URL of the Configuration resource |[optional] 
**FlexInsightsHr** | Pointer to **interface{}** | Object that controls workspace reporting |
**FlexInsightsDrilldown** | **bool** | Setting to enable Flex UI redirection |[optional] 
**FlexUrl** | **string** | URL to redirect to in case drilldown is enabled. |[optional] 
**ChannelConfigs** | **[]interface{}** | Flex Conversations channels' attachments configurations |[optional] 
**DebuggerIntegration** | Pointer to **interface{}** | Configurable parameters for Debugger Integration |
**FlexUiStatusReport** | Pointer to **interface{}** | Configurable parameters for Flex UI Status report |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


