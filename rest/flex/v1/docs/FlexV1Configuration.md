# FlexV1Configuration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** | The SID of the Account that created the resource |
**Attributes** | Pointer to **map[string]interface{}** | An object that contains application-specific data |
**CallRecordingEnabled** | Pointer to **bool** | Whether call recording is enabled |
**CallRecordingWebhookUrl** | Pointer to **string** | The call recording webhook URL |
**ChatServiceInstanceSid** | Pointer to **string** | The SID of the chat service this user belongs to |
**CrmAttributes** | Pointer to **map[string]interface{}** | An object that contains the CRM attributes |
**CrmCallbackUrl** | Pointer to **string** | The CRM Callback URL |
**CrmEnabled** | Pointer to **bool** | Whether CRM is present for Flex |
**CrmFallbackUrl** | Pointer to **string** | The CRM Fallback URL |
**CrmType** | Pointer to **string** | The CRM Type |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the Configuration resource was created |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the Configuration resource was last updated |
**FlexServiceInstanceSid** | Pointer to **string** | The SID of the Flex service instance |
**Integrations** | Pointer to **[]map[string]interface{}** | A list of objects that contain the configurations for the Integrations supported in this configuration |
**Markdown** | Pointer to **map[string]interface{}** | Configurable parameters for Markdown |
**MessagingServiceInstanceSid** | Pointer to **string** | The SID of the Messaging service instance |
**Notifications** | Pointer to **map[string]interface{}** | Configurable parameters for Notifications |
**OutboundCallFlows** | Pointer to **map[string]interface{}** | The list of outbound call flows |
**PluginServiceAttributes** | Pointer to **map[string]interface{}** | The plugin service attributes |
**PluginServiceEnabled** | Pointer to **bool** | Whether the plugin service enabled |
**PublicAttributes** | Pointer to **map[string]interface{}** | The list of public attributes |
**QueueStatsConfiguration** | Pointer to **map[string]interface{}** | Configurable parameters for Queues Statistics |
**RuntimeDomain** | Pointer to **string** | The URL where the Flex instance is hosted |
**ServerlessServiceSids** | Pointer to **[]string** | The list of serverless service SIDs |
**ServiceVersion** | Pointer to **string** | The Flex Service version |
**Status** | Pointer to **string** | The status of the Flex onboarding |
**TaskrouterOfflineActivitySid** | Pointer to **string** | The TaskRouter SID of the offline activity |
**TaskrouterSkills** | Pointer to **[]map[string]interface{}** | The Skill description for TaskRouter workers |
**TaskrouterTargetTaskqueueSid** | Pointer to **string** | The SID of the TaskRouter Target TaskQueue |
**TaskrouterTargetWorkflowSid** | Pointer to **string** | The SID of the TaskRouter target Workflow |
**TaskrouterTaskqueues** | Pointer to **[]map[string]interface{}** | The list of TaskRouter TaskQueues |
**TaskrouterWorkerAttributes** | Pointer to **map[string]interface{}** | The TaskRouter Worker attributes |
**TaskrouterWorkerChannels** | Pointer to **map[string]interface{}** | The TaskRouter default channel capacities and availability for workers |
**TaskrouterWorkspaceSid** | Pointer to **string** | The SID of the TaskRouter Workspace |
**UiAttributes** | Pointer to **map[string]interface{}** | The object that describes Flex UI characteristics and settings |
**UiDependencies** | Pointer to **map[string]interface{}** | The object that defines the NPM packages and versions to be used in Hosted Flex |
**UiLanguage** | Pointer to **string** | The primary language of the Flex UI |
**UiVersion** | Pointer to **string** | The Pinned UI version |
**Url** | Pointer to **string** | The absolute URL of the Configuration resource |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


