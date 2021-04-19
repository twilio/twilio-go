# FlexV1Configuration

## Properties

Name | Type | Description
------------ | ------------- | -------------
**AccountSid** | Pointer to **NullableString** | The SID of the Account that created the resource | [optional] 
**Attributes** | Pointer to **map[string]interface{}** | An object that contains application-specific data | [optional] 
**CallRecordingEnabled** | Pointer to **NullableBool** | Whether call recording is enabled | [optional] 
**CallRecordingWebhookUrl** | Pointer to **NullableString** | The call recording webhook URL | [optional] 
**ChatServiceInstanceSid** | Pointer to **NullableString** | The SID of the chat service this user belongs to | [optional] 
**CrmAttributes** | Pointer to **map[string]interface{}** | An object that contains the CRM attributes | [optional] 
**CrmCallbackUrl** | Pointer to **NullableString** | The CRM Callback URL | [optional] 
**CrmEnabled** | Pointer to **NullableBool** | Whether CRM is present for Flex | [optional] 
**CrmFallbackUrl** | Pointer to **NullableString** | The CRM Fallback URL | [optional] 
**CrmType** | Pointer to **NullableString** | The CRM Type | [optional] 
**DateCreated** | Pointer to **NullableTime** | The ISO 8601 date and time in GMT when the Configuration resource was created | [optional] 
**DateUpdated** | Pointer to **NullableTime** | The ISO 8601 date and time in GMT when the Configuration resource was last updated | [optional] 
**FlexServiceInstanceSid** | Pointer to **NullableString** | The SID of the Flex service instance | [optional] 
**Integrations** | Pointer to **[]map[string]interface{}** | A list of objects that contain the configurations for the Integrations supported in this configuration | [optional] 
**Markdown** | Pointer to **map[string]interface{}** | Configurable parameters for Markdown | [optional] 
**MessagingServiceInstanceSid** | Pointer to **NullableString** | The SID of the Messaging service instance | [optional] 
**Notifications** | Pointer to **map[string]interface{}** | Configurable parameters for Notifications | [optional] 
**OutboundCallFlows** | Pointer to **map[string]interface{}** | The list of outbound call flows | [optional] 
**PluginServiceAttributes** | Pointer to **map[string]interface{}** | The plugin service attributes | [optional] 
**PluginServiceEnabled** | Pointer to **NullableBool** | Whether the plugin service enabled | [optional] 
**PublicAttributes** | Pointer to **map[string]interface{}** | The list of public attributes | [optional] 
**QueueStatsConfiguration** | Pointer to **map[string]interface{}** | Configurable parameters for Queues Statistics | [optional] 
**RuntimeDomain** | Pointer to **NullableString** | The URL where the Flex instance is hosted | [optional] 
**ServerlessServiceSids** | Pointer to **[]string** | The list of serverless service SIDs | [optional] 
**ServiceVersion** | Pointer to **NullableString** | The Flex Service version | [optional] 
**Status** | Pointer to **NullableString** | The status of the Flex onboarding | [optional] 
**TaskrouterOfflineActivitySid** | Pointer to **NullableString** | The TaskRouter SID of the offline activity | [optional] 
**TaskrouterSkills** | Pointer to **[]map[string]interface{}** | The Skill description for TaskRouter workers | [optional] 
**TaskrouterTargetTaskqueueSid** | Pointer to **NullableString** | The SID of the TaskRouter Target TaskQueue | [optional] 
**TaskrouterTargetWorkflowSid** | Pointer to **NullableString** | The SID of the TaskRouter target Workflow | [optional] 
**TaskrouterTaskqueues** | Pointer to **[]map[string]interface{}** | The list of TaskRouter TaskQueues | [optional] 
**TaskrouterWorkerAttributes** | Pointer to **map[string]interface{}** | The TaskRouter Worker attributes | [optional] 
**TaskrouterWorkerChannels** | Pointer to **map[string]interface{}** | The TaskRouter default channel capacities and availability for workers | [optional] 
**TaskrouterWorkspaceSid** | Pointer to **NullableString** | The SID of the TaskRouter Workspace | [optional] 
**UiAttributes** | Pointer to **map[string]interface{}** | The object that describes Flex UI characteristics and settings | [optional] 
**UiDependencies** | Pointer to **map[string]interface{}** | The object that defines the NPM packages and versions to be used in Hosted Flex | [optional] 
**UiLanguage** | Pointer to **NullableString** | The primary language of the Flex UI | [optional] 
**UiVersion** | Pointer to **NullableString** | The Pinned UI version | [optional] 
**Url** | Pointer to **NullableString** | The absolute URL of the Configuration resource | [optional] 

## Methods

### NewFlexV1Configuration

`func NewFlexV1Configuration() *FlexV1Configuration`

NewFlexV1Configuration instantiates a new FlexV1Configuration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlexV1ConfigurationWithDefaults

`func NewFlexV1ConfigurationWithDefaults() *FlexV1Configuration`

NewFlexV1ConfigurationWithDefaults instantiates a new FlexV1Configuration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSid

`func (o *FlexV1Configuration) GetAccountSid() string`

GetAccountSid returns the AccountSid field if non-nil, zero value otherwise.

### GetAccountSidOk

`func (o *FlexV1Configuration) GetAccountSidOk() (*string, bool)`

GetAccountSidOk returns a tuple with the AccountSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSid

`func (o *FlexV1Configuration) SetAccountSid(v string)`

SetAccountSid sets AccountSid field to given value.

### HasAccountSid

`func (o *FlexV1Configuration) HasAccountSid() bool`

HasAccountSid returns a boolean if a field has been set.

### SetAccountSidNil

`func (o *FlexV1Configuration) SetAccountSidNil(b bool)`

 SetAccountSidNil sets the value for AccountSid to be an explicit nil

### UnsetAccountSid
`func (o *FlexV1Configuration) UnsetAccountSid()`

UnsetAccountSid ensures that no value is present for AccountSid, not even an explicit nil
### GetAttributes

`func (o *FlexV1Configuration) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *FlexV1Configuration) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *FlexV1Configuration) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *FlexV1Configuration) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### SetAttributesNil

`func (o *FlexV1Configuration) SetAttributesNil(b bool)`

 SetAttributesNil sets the value for Attributes to be an explicit nil

### UnsetAttributes
`func (o *FlexV1Configuration) UnsetAttributes()`

UnsetAttributes ensures that no value is present for Attributes, not even an explicit nil
### GetCallRecordingEnabled

`func (o *FlexV1Configuration) GetCallRecordingEnabled() bool`

GetCallRecordingEnabled returns the CallRecordingEnabled field if non-nil, zero value otherwise.

### GetCallRecordingEnabledOk

`func (o *FlexV1Configuration) GetCallRecordingEnabledOk() (*bool, bool)`

GetCallRecordingEnabledOk returns a tuple with the CallRecordingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallRecordingEnabled

`func (o *FlexV1Configuration) SetCallRecordingEnabled(v bool)`

SetCallRecordingEnabled sets CallRecordingEnabled field to given value.

### HasCallRecordingEnabled

`func (o *FlexV1Configuration) HasCallRecordingEnabled() bool`

HasCallRecordingEnabled returns a boolean if a field has been set.

### SetCallRecordingEnabledNil

`func (o *FlexV1Configuration) SetCallRecordingEnabledNil(b bool)`

 SetCallRecordingEnabledNil sets the value for CallRecordingEnabled to be an explicit nil

### UnsetCallRecordingEnabled
`func (o *FlexV1Configuration) UnsetCallRecordingEnabled()`

UnsetCallRecordingEnabled ensures that no value is present for CallRecordingEnabled, not even an explicit nil
### GetCallRecordingWebhookUrl

`func (o *FlexV1Configuration) GetCallRecordingWebhookUrl() string`

GetCallRecordingWebhookUrl returns the CallRecordingWebhookUrl field if non-nil, zero value otherwise.

### GetCallRecordingWebhookUrlOk

`func (o *FlexV1Configuration) GetCallRecordingWebhookUrlOk() (*string, bool)`

GetCallRecordingWebhookUrlOk returns a tuple with the CallRecordingWebhookUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallRecordingWebhookUrl

`func (o *FlexV1Configuration) SetCallRecordingWebhookUrl(v string)`

SetCallRecordingWebhookUrl sets CallRecordingWebhookUrl field to given value.

### HasCallRecordingWebhookUrl

`func (o *FlexV1Configuration) HasCallRecordingWebhookUrl() bool`

HasCallRecordingWebhookUrl returns a boolean if a field has been set.

### SetCallRecordingWebhookUrlNil

`func (o *FlexV1Configuration) SetCallRecordingWebhookUrlNil(b bool)`

 SetCallRecordingWebhookUrlNil sets the value for CallRecordingWebhookUrl to be an explicit nil

### UnsetCallRecordingWebhookUrl
`func (o *FlexV1Configuration) UnsetCallRecordingWebhookUrl()`

UnsetCallRecordingWebhookUrl ensures that no value is present for CallRecordingWebhookUrl, not even an explicit nil
### GetChatServiceInstanceSid

`func (o *FlexV1Configuration) GetChatServiceInstanceSid() string`

GetChatServiceInstanceSid returns the ChatServiceInstanceSid field if non-nil, zero value otherwise.

### GetChatServiceInstanceSidOk

`func (o *FlexV1Configuration) GetChatServiceInstanceSidOk() (*string, bool)`

GetChatServiceInstanceSidOk returns a tuple with the ChatServiceInstanceSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatServiceInstanceSid

`func (o *FlexV1Configuration) SetChatServiceInstanceSid(v string)`

SetChatServiceInstanceSid sets ChatServiceInstanceSid field to given value.

### HasChatServiceInstanceSid

`func (o *FlexV1Configuration) HasChatServiceInstanceSid() bool`

HasChatServiceInstanceSid returns a boolean if a field has been set.

### SetChatServiceInstanceSidNil

`func (o *FlexV1Configuration) SetChatServiceInstanceSidNil(b bool)`

 SetChatServiceInstanceSidNil sets the value for ChatServiceInstanceSid to be an explicit nil

### UnsetChatServiceInstanceSid
`func (o *FlexV1Configuration) UnsetChatServiceInstanceSid()`

UnsetChatServiceInstanceSid ensures that no value is present for ChatServiceInstanceSid, not even an explicit nil
### GetCrmAttributes

`func (o *FlexV1Configuration) GetCrmAttributes() map[string]interface{}`

GetCrmAttributes returns the CrmAttributes field if non-nil, zero value otherwise.

### GetCrmAttributesOk

`func (o *FlexV1Configuration) GetCrmAttributesOk() (*map[string]interface{}, bool)`

GetCrmAttributesOk returns a tuple with the CrmAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrmAttributes

`func (o *FlexV1Configuration) SetCrmAttributes(v map[string]interface{})`

SetCrmAttributes sets CrmAttributes field to given value.

### HasCrmAttributes

`func (o *FlexV1Configuration) HasCrmAttributes() bool`

HasCrmAttributes returns a boolean if a field has been set.

### SetCrmAttributesNil

`func (o *FlexV1Configuration) SetCrmAttributesNil(b bool)`

 SetCrmAttributesNil sets the value for CrmAttributes to be an explicit nil

### UnsetCrmAttributes
`func (o *FlexV1Configuration) UnsetCrmAttributes()`

UnsetCrmAttributes ensures that no value is present for CrmAttributes, not even an explicit nil
### GetCrmCallbackUrl

`func (o *FlexV1Configuration) GetCrmCallbackUrl() string`

GetCrmCallbackUrl returns the CrmCallbackUrl field if non-nil, zero value otherwise.

### GetCrmCallbackUrlOk

`func (o *FlexV1Configuration) GetCrmCallbackUrlOk() (*string, bool)`

GetCrmCallbackUrlOk returns a tuple with the CrmCallbackUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrmCallbackUrl

`func (o *FlexV1Configuration) SetCrmCallbackUrl(v string)`

SetCrmCallbackUrl sets CrmCallbackUrl field to given value.

### HasCrmCallbackUrl

`func (o *FlexV1Configuration) HasCrmCallbackUrl() bool`

HasCrmCallbackUrl returns a boolean if a field has been set.

### SetCrmCallbackUrlNil

`func (o *FlexV1Configuration) SetCrmCallbackUrlNil(b bool)`

 SetCrmCallbackUrlNil sets the value for CrmCallbackUrl to be an explicit nil

### UnsetCrmCallbackUrl
`func (o *FlexV1Configuration) UnsetCrmCallbackUrl()`

UnsetCrmCallbackUrl ensures that no value is present for CrmCallbackUrl, not even an explicit nil
### GetCrmEnabled

`func (o *FlexV1Configuration) GetCrmEnabled() bool`

GetCrmEnabled returns the CrmEnabled field if non-nil, zero value otherwise.

### GetCrmEnabledOk

`func (o *FlexV1Configuration) GetCrmEnabledOk() (*bool, bool)`

GetCrmEnabledOk returns a tuple with the CrmEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrmEnabled

`func (o *FlexV1Configuration) SetCrmEnabled(v bool)`

SetCrmEnabled sets CrmEnabled field to given value.

### HasCrmEnabled

`func (o *FlexV1Configuration) HasCrmEnabled() bool`

HasCrmEnabled returns a boolean if a field has been set.

### SetCrmEnabledNil

`func (o *FlexV1Configuration) SetCrmEnabledNil(b bool)`

 SetCrmEnabledNil sets the value for CrmEnabled to be an explicit nil

### UnsetCrmEnabled
`func (o *FlexV1Configuration) UnsetCrmEnabled()`

UnsetCrmEnabled ensures that no value is present for CrmEnabled, not even an explicit nil
### GetCrmFallbackUrl

`func (o *FlexV1Configuration) GetCrmFallbackUrl() string`

GetCrmFallbackUrl returns the CrmFallbackUrl field if non-nil, zero value otherwise.

### GetCrmFallbackUrlOk

`func (o *FlexV1Configuration) GetCrmFallbackUrlOk() (*string, bool)`

GetCrmFallbackUrlOk returns a tuple with the CrmFallbackUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrmFallbackUrl

`func (o *FlexV1Configuration) SetCrmFallbackUrl(v string)`

SetCrmFallbackUrl sets CrmFallbackUrl field to given value.

### HasCrmFallbackUrl

`func (o *FlexV1Configuration) HasCrmFallbackUrl() bool`

HasCrmFallbackUrl returns a boolean if a field has been set.

### SetCrmFallbackUrlNil

`func (o *FlexV1Configuration) SetCrmFallbackUrlNil(b bool)`

 SetCrmFallbackUrlNil sets the value for CrmFallbackUrl to be an explicit nil

### UnsetCrmFallbackUrl
`func (o *FlexV1Configuration) UnsetCrmFallbackUrl()`

UnsetCrmFallbackUrl ensures that no value is present for CrmFallbackUrl, not even an explicit nil
### GetCrmType

`func (o *FlexV1Configuration) GetCrmType() string`

GetCrmType returns the CrmType field if non-nil, zero value otherwise.

### GetCrmTypeOk

`func (o *FlexV1Configuration) GetCrmTypeOk() (*string, bool)`

GetCrmTypeOk returns a tuple with the CrmType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrmType

`func (o *FlexV1Configuration) SetCrmType(v string)`

SetCrmType sets CrmType field to given value.

### HasCrmType

`func (o *FlexV1Configuration) HasCrmType() bool`

HasCrmType returns a boolean if a field has been set.

### SetCrmTypeNil

`func (o *FlexV1Configuration) SetCrmTypeNil(b bool)`

 SetCrmTypeNil sets the value for CrmType to be an explicit nil

### UnsetCrmType
`func (o *FlexV1Configuration) UnsetCrmType()`

UnsetCrmType ensures that no value is present for CrmType, not even an explicit nil
### GetDateCreated

`func (o *FlexV1Configuration) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *FlexV1Configuration) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *FlexV1Configuration) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *FlexV1Configuration) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### SetDateCreatedNil

`func (o *FlexV1Configuration) SetDateCreatedNil(b bool)`

 SetDateCreatedNil sets the value for DateCreated to be an explicit nil

### UnsetDateCreated
`func (o *FlexV1Configuration) UnsetDateCreated()`

UnsetDateCreated ensures that no value is present for DateCreated, not even an explicit nil
### GetDateUpdated

`func (o *FlexV1Configuration) GetDateUpdated() time.Time`

GetDateUpdated returns the DateUpdated field if non-nil, zero value otherwise.

### GetDateUpdatedOk

`func (o *FlexV1Configuration) GetDateUpdatedOk() (*time.Time, bool)`

GetDateUpdatedOk returns a tuple with the DateUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateUpdated

`func (o *FlexV1Configuration) SetDateUpdated(v time.Time)`

SetDateUpdated sets DateUpdated field to given value.

### HasDateUpdated

`func (o *FlexV1Configuration) HasDateUpdated() bool`

HasDateUpdated returns a boolean if a field has been set.

### SetDateUpdatedNil

`func (o *FlexV1Configuration) SetDateUpdatedNil(b bool)`

 SetDateUpdatedNil sets the value for DateUpdated to be an explicit nil

### UnsetDateUpdated
`func (o *FlexV1Configuration) UnsetDateUpdated()`

UnsetDateUpdated ensures that no value is present for DateUpdated, not even an explicit nil
### GetFlexServiceInstanceSid

`func (o *FlexV1Configuration) GetFlexServiceInstanceSid() string`

GetFlexServiceInstanceSid returns the FlexServiceInstanceSid field if non-nil, zero value otherwise.

### GetFlexServiceInstanceSidOk

`func (o *FlexV1Configuration) GetFlexServiceInstanceSidOk() (*string, bool)`

GetFlexServiceInstanceSidOk returns a tuple with the FlexServiceInstanceSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlexServiceInstanceSid

`func (o *FlexV1Configuration) SetFlexServiceInstanceSid(v string)`

SetFlexServiceInstanceSid sets FlexServiceInstanceSid field to given value.

### HasFlexServiceInstanceSid

`func (o *FlexV1Configuration) HasFlexServiceInstanceSid() bool`

HasFlexServiceInstanceSid returns a boolean if a field has been set.

### SetFlexServiceInstanceSidNil

`func (o *FlexV1Configuration) SetFlexServiceInstanceSidNil(b bool)`

 SetFlexServiceInstanceSidNil sets the value for FlexServiceInstanceSid to be an explicit nil

### UnsetFlexServiceInstanceSid
`func (o *FlexV1Configuration) UnsetFlexServiceInstanceSid()`

UnsetFlexServiceInstanceSid ensures that no value is present for FlexServiceInstanceSid, not even an explicit nil
### GetIntegrations

`func (o *FlexV1Configuration) GetIntegrations() []map[string]interface{}`

GetIntegrations returns the Integrations field if non-nil, zero value otherwise.

### GetIntegrationsOk

`func (o *FlexV1Configuration) GetIntegrationsOk() (*[]map[string]interface{}, bool)`

GetIntegrationsOk returns a tuple with the Integrations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrations

`func (o *FlexV1Configuration) SetIntegrations(v []map[string]interface{})`

SetIntegrations sets Integrations field to given value.

### HasIntegrations

`func (o *FlexV1Configuration) HasIntegrations() bool`

HasIntegrations returns a boolean if a field has been set.

### SetIntegrationsNil

`func (o *FlexV1Configuration) SetIntegrationsNil(b bool)`

 SetIntegrationsNil sets the value for Integrations to be an explicit nil

### UnsetIntegrations
`func (o *FlexV1Configuration) UnsetIntegrations()`

UnsetIntegrations ensures that no value is present for Integrations, not even an explicit nil
### GetMarkdown

`func (o *FlexV1Configuration) GetMarkdown() map[string]interface{}`

GetMarkdown returns the Markdown field if non-nil, zero value otherwise.

### GetMarkdownOk

`func (o *FlexV1Configuration) GetMarkdownOk() (*map[string]interface{}, bool)`

GetMarkdownOk returns a tuple with the Markdown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkdown

`func (o *FlexV1Configuration) SetMarkdown(v map[string]interface{})`

SetMarkdown sets Markdown field to given value.

### HasMarkdown

`func (o *FlexV1Configuration) HasMarkdown() bool`

HasMarkdown returns a boolean if a field has been set.

### SetMarkdownNil

`func (o *FlexV1Configuration) SetMarkdownNil(b bool)`

 SetMarkdownNil sets the value for Markdown to be an explicit nil

### UnsetMarkdown
`func (o *FlexV1Configuration) UnsetMarkdown()`

UnsetMarkdown ensures that no value is present for Markdown, not even an explicit nil
### GetMessagingServiceInstanceSid

`func (o *FlexV1Configuration) GetMessagingServiceInstanceSid() string`

GetMessagingServiceInstanceSid returns the MessagingServiceInstanceSid field if non-nil, zero value otherwise.

### GetMessagingServiceInstanceSidOk

`func (o *FlexV1Configuration) GetMessagingServiceInstanceSidOk() (*string, bool)`

GetMessagingServiceInstanceSidOk returns a tuple with the MessagingServiceInstanceSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessagingServiceInstanceSid

`func (o *FlexV1Configuration) SetMessagingServiceInstanceSid(v string)`

SetMessagingServiceInstanceSid sets MessagingServiceInstanceSid field to given value.

### HasMessagingServiceInstanceSid

`func (o *FlexV1Configuration) HasMessagingServiceInstanceSid() bool`

HasMessagingServiceInstanceSid returns a boolean if a field has been set.

### SetMessagingServiceInstanceSidNil

`func (o *FlexV1Configuration) SetMessagingServiceInstanceSidNil(b bool)`

 SetMessagingServiceInstanceSidNil sets the value for MessagingServiceInstanceSid to be an explicit nil

### UnsetMessagingServiceInstanceSid
`func (o *FlexV1Configuration) UnsetMessagingServiceInstanceSid()`

UnsetMessagingServiceInstanceSid ensures that no value is present for MessagingServiceInstanceSid, not even an explicit nil
### GetNotifications

`func (o *FlexV1Configuration) GetNotifications() map[string]interface{}`

GetNotifications returns the Notifications field if non-nil, zero value otherwise.

### GetNotificationsOk

`func (o *FlexV1Configuration) GetNotificationsOk() (*map[string]interface{}, bool)`

GetNotificationsOk returns a tuple with the Notifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifications

`func (o *FlexV1Configuration) SetNotifications(v map[string]interface{})`

SetNotifications sets Notifications field to given value.

### HasNotifications

`func (o *FlexV1Configuration) HasNotifications() bool`

HasNotifications returns a boolean if a field has been set.

### SetNotificationsNil

`func (o *FlexV1Configuration) SetNotificationsNil(b bool)`

 SetNotificationsNil sets the value for Notifications to be an explicit nil

### UnsetNotifications
`func (o *FlexV1Configuration) UnsetNotifications()`

UnsetNotifications ensures that no value is present for Notifications, not even an explicit nil
### GetOutboundCallFlows

`func (o *FlexV1Configuration) GetOutboundCallFlows() map[string]interface{}`

GetOutboundCallFlows returns the OutboundCallFlows field if non-nil, zero value otherwise.

### GetOutboundCallFlowsOk

`func (o *FlexV1Configuration) GetOutboundCallFlowsOk() (*map[string]interface{}, bool)`

GetOutboundCallFlowsOk returns a tuple with the OutboundCallFlows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutboundCallFlows

`func (o *FlexV1Configuration) SetOutboundCallFlows(v map[string]interface{})`

SetOutboundCallFlows sets OutboundCallFlows field to given value.

### HasOutboundCallFlows

`func (o *FlexV1Configuration) HasOutboundCallFlows() bool`

HasOutboundCallFlows returns a boolean if a field has been set.

### SetOutboundCallFlowsNil

`func (o *FlexV1Configuration) SetOutboundCallFlowsNil(b bool)`

 SetOutboundCallFlowsNil sets the value for OutboundCallFlows to be an explicit nil

### UnsetOutboundCallFlows
`func (o *FlexV1Configuration) UnsetOutboundCallFlows()`

UnsetOutboundCallFlows ensures that no value is present for OutboundCallFlows, not even an explicit nil
### GetPluginServiceAttributes

`func (o *FlexV1Configuration) GetPluginServiceAttributes() map[string]interface{}`

GetPluginServiceAttributes returns the PluginServiceAttributes field if non-nil, zero value otherwise.

### GetPluginServiceAttributesOk

`func (o *FlexV1Configuration) GetPluginServiceAttributesOk() (*map[string]interface{}, bool)`

GetPluginServiceAttributesOk returns a tuple with the PluginServiceAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginServiceAttributes

`func (o *FlexV1Configuration) SetPluginServiceAttributes(v map[string]interface{})`

SetPluginServiceAttributes sets PluginServiceAttributes field to given value.

### HasPluginServiceAttributes

`func (o *FlexV1Configuration) HasPluginServiceAttributes() bool`

HasPluginServiceAttributes returns a boolean if a field has been set.

### SetPluginServiceAttributesNil

`func (o *FlexV1Configuration) SetPluginServiceAttributesNil(b bool)`

 SetPluginServiceAttributesNil sets the value for PluginServiceAttributes to be an explicit nil

### UnsetPluginServiceAttributes
`func (o *FlexV1Configuration) UnsetPluginServiceAttributes()`

UnsetPluginServiceAttributes ensures that no value is present for PluginServiceAttributes, not even an explicit nil
### GetPluginServiceEnabled

`func (o *FlexV1Configuration) GetPluginServiceEnabled() bool`

GetPluginServiceEnabled returns the PluginServiceEnabled field if non-nil, zero value otherwise.

### GetPluginServiceEnabledOk

`func (o *FlexV1Configuration) GetPluginServiceEnabledOk() (*bool, bool)`

GetPluginServiceEnabledOk returns a tuple with the PluginServiceEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginServiceEnabled

`func (o *FlexV1Configuration) SetPluginServiceEnabled(v bool)`

SetPluginServiceEnabled sets PluginServiceEnabled field to given value.

### HasPluginServiceEnabled

`func (o *FlexV1Configuration) HasPluginServiceEnabled() bool`

HasPluginServiceEnabled returns a boolean if a field has been set.

### SetPluginServiceEnabledNil

`func (o *FlexV1Configuration) SetPluginServiceEnabledNil(b bool)`

 SetPluginServiceEnabledNil sets the value for PluginServiceEnabled to be an explicit nil

### UnsetPluginServiceEnabled
`func (o *FlexV1Configuration) UnsetPluginServiceEnabled()`

UnsetPluginServiceEnabled ensures that no value is present for PluginServiceEnabled, not even an explicit nil
### GetPublicAttributes

`func (o *FlexV1Configuration) GetPublicAttributes() map[string]interface{}`

GetPublicAttributes returns the PublicAttributes field if non-nil, zero value otherwise.

### GetPublicAttributesOk

`func (o *FlexV1Configuration) GetPublicAttributesOk() (*map[string]interface{}, bool)`

GetPublicAttributesOk returns a tuple with the PublicAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicAttributes

`func (o *FlexV1Configuration) SetPublicAttributes(v map[string]interface{})`

SetPublicAttributes sets PublicAttributes field to given value.

### HasPublicAttributes

`func (o *FlexV1Configuration) HasPublicAttributes() bool`

HasPublicAttributes returns a boolean if a field has been set.

### SetPublicAttributesNil

`func (o *FlexV1Configuration) SetPublicAttributesNil(b bool)`

 SetPublicAttributesNil sets the value for PublicAttributes to be an explicit nil

### UnsetPublicAttributes
`func (o *FlexV1Configuration) UnsetPublicAttributes()`

UnsetPublicAttributes ensures that no value is present for PublicAttributes, not even an explicit nil
### GetQueueStatsConfiguration

`func (o *FlexV1Configuration) GetQueueStatsConfiguration() map[string]interface{}`

GetQueueStatsConfiguration returns the QueueStatsConfiguration field if non-nil, zero value otherwise.

### GetQueueStatsConfigurationOk

`func (o *FlexV1Configuration) GetQueueStatsConfigurationOk() (*map[string]interface{}, bool)`

GetQueueStatsConfigurationOk returns a tuple with the QueueStatsConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueueStatsConfiguration

`func (o *FlexV1Configuration) SetQueueStatsConfiguration(v map[string]interface{})`

SetQueueStatsConfiguration sets QueueStatsConfiguration field to given value.

### HasQueueStatsConfiguration

`func (o *FlexV1Configuration) HasQueueStatsConfiguration() bool`

HasQueueStatsConfiguration returns a boolean if a field has been set.

### SetQueueStatsConfigurationNil

`func (o *FlexV1Configuration) SetQueueStatsConfigurationNil(b bool)`

 SetQueueStatsConfigurationNil sets the value for QueueStatsConfiguration to be an explicit nil

### UnsetQueueStatsConfiguration
`func (o *FlexV1Configuration) UnsetQueueStatsConfiguration()`

UnsetQueueStatsConfiguration ensures that no value is present for QueueStatsConfiguration, not even an explicit nil
### GetRuntimeDomain

`func (o *FlexV1Configuration) GetRuntimeDomain() string`

GetRuntimeDomain returns the RuntimeDomain field if non-nil, zero value otherwise.

### GetRuntimeDomainOk

`func (o *FlexV1Configuration) GetRuntimeDomainOk() (*string, bool)`

GetRuntimeDomainOk returns a tuple with the RuntimeDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuntimeDomain

`func (o *FlexV1Configuration) SetRuntimeDomain(v string)`

SetRuntimeDomain sets RuntimeDomain field to given value.

### HasRuntimeDomain

`func (o *FlexV1Configuration) HasRuntimeDomain() bool`

HasRuntimeDomain returns a boolean if a field has been set.

### SetRuntimeDomainNil

`func (o *FlexV1Configuration) SetRuntimeDomainNil(b bool)`

 SetRuntimeDomainNil sets the value for RuntimeDomain to be an explicit nil

### UnsetRuntimeDomain
`func (o *FlexV1Configuration) UnsetRuntimeDomain()`

UnsetRuntimeDomain ensures that no value is present for RuntimeDomain, not even an explicit nil
### GetServerlessServiceSids

`func (o *FlexV1Configuration) GetServerlessServiceSids() []string`

GetServerlessServiceSids returns the ServerlessServiceSids field if non-nil, zero value otherwise.

### GetServerlessServiceSidsOk

`func (o *FlexV1Configuration) GetServerlessServiceSidsOk() (*[]string, bool)`

GetServerlessServiceSidsOk returns a tuple with the ServerlessServiceSids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerlessServiceSids

`func (o *FlexV1Configuration) SetServerlessServiceSids(v []string)`

SetServerlessServiceSids sets ServerlessServiceSids field to given value.

### HasServerlessServiceSids

`func (o *FlexV1Configuration) HasServerlessServiceSids() bool`

HasServerlessServiceSids returns a boolean if a field has been set.

### SetServerlessServiceSidsNil

`func (o *FlexV1Configuration) SetServerlessServiceSidsNil(b bool)`

 SetServerlessServiceSidsNil sets the value for ServerlessServiceSids to be an explicit nil

### UnsetServerlessServiceSids
`func (o *FlexV1Configuration) UnsetServerlessServiceSids()`

UnsetServerlessServiceSids ensures that no value is present for ServerlessServiceSids, not even an explicit nil
### GetServiceVersion

`func (o *FlexV1Configuration) GetServiceVersion() string`

GetServiceVersion returns the ServiceVersion field if non-nil, zero value otherwise.

### GetServiceVersionOk

`func (o *FlexV1Configuration) GetServiceVersionOk() (*string, bool)`

GetServiceVersionOk returns a tuple with the ServiceVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceVersion

`func (o *FlexV1Configuration) SetServiceVersion(v string)`

SetServiceVersion sets ServiceVersion field to given value.

### HasServiceVersion

`func (o *FlexV1Configuration) HasServiceVersion() bool`

HasServiceVersion returns a boolean if a field has been set.

### SetServiceVersionNil

`func (o *FlexV1Configuration) SetServiceVersionNil(b bool)`

 SetServiceVersionNil sets the value for ServiceVersion to be an explicit nil

### UnsetServiceVersion
`func (o *FlexV1Configuration) UnsetServiceVersion()`

UnsetServiceVersion ensures that no value is present for ServiceVersion, not even an explicit nil
### GetStatus

`func (o *FlexV1Configuration) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *FlexV1Configuration) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *FlexV1Configuration) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *FlexV1Configuration) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *FlexV1Configuration) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *FlexV1Configuration) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetTaskrouterOfflineActivitySid

`func (o *FlexV1Configuration) GetTaskrouterOfflineActivitySid() string`

GetTaskrouterOfflineActivitySid returns the TaskrouterOfflineActivitySid field if non-nil, zero value otherwise.

### GetTaskrouterOfflineActivitySidOk

`func (o *FlexV1Configuration) GetTaskrouterOfflineActivitySidOk() (*string, bool)`

GetTaskrouterOfflineActivitySidOk returns a tuple with the TaskrouterOfflineActivitySid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskrouterOfflineActivitySid

`func (o *FlexV1Configuration) SetTaskrouterOfflineActivitySid(v string)`

SetTaskrouterOfflineActivitySid sets TaskrouterOfflineActivitySid field to given value.

### HasTaskrouterOfflineActivitySid

`func (o *FlexV1Configuration) HasTaskrouterOfflineActivitySid() bool`

HasTaskrouterOfflineActivitySid returns a boolean if a field has been set.

### SetTaskrouterOfflineActivitySidNil

`func (o *FlexV1Configuration) SetTaskrouterOfflineActivitySidNil(b bool)`

 SetTaskrouterOfflineActivitySidNil sets the value for TaskrouterOfflineActivitySid to be an explicit nil

### UnsetTaskrouterOfflineActivitySid
`func (o *FlexV1Configuration) UnsetTaskrouterOfflineActivitySid()`

UnsetTaskrouterOfflineActivitySid ensures that no value is present for TaskrouterOfflineActivitySid, not even an explicit nil
### GetTaskrouterSkills

`func (o *FlexV1Configuration) GetTaskrouterSkills() []map[string]interface{}`

GetTaskrouterSkills returns the TaskrouterSkills field if non-nil, zero value otherwise.

### GetTaskrouterSkillsOk

`func (o *FlexV1Configuration) GetTaskrouterSkillsOk() (*[]map[string]interface{}, bool)`

GetTaskrouterSkillsOk returns a tuple with the TaskrouterSkills field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskrouterSkills

`func (o *FlexV1Configuration) SetTaskrouterSkills(v []map[string]interface{})`

SetTaskrouterSkills sets TaskrouterSkills field to given value.

### HasTaskrouterSkills

`func (o *FlexV1Configuration) HasTaskrouterSkills() bool`

HasTaskrouterSkills returns a boolean if a field has been set.

### SetTaskrouterSkillsNil

`func (o *FlexV1Configuration) SetTaskrouterSkillsNil(b bool)`

 SetTaskrouterSkillsNil sets the value for TaskrouterSkills to be an explicit nil

### UnsetTaskrouterSkills
`func (o *FlexV1Configuration) UnsetTaskrouterSkills()`

UnsetTaskrouterSkills ensures that no value is present for TaskrouterSkills, not even an explicit nil
### GetTaskrouterTargetTaskqueueSid

`func (o *FlexV1Configuration) GetTaskrouterTargetTaskqueueSid() string`

GetTaskrouterTargetTaskqueueSid returns the TaskrouterTargetTaskqueueSid field if non-nil, zero value otherwise.

### GetTaskrouterTargetTaskqueueSidOk

`func (o *FlexV1Configuration) GetTaskrouterTargetTaskqueueSidOk() (*string, bool)`

GetTaskrouterTargetTaskqueueSidOk returns a tuple with the TaskrouterTargetTaskqueueSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskrouterTargetTaskqueueSid

`func (o *FlexV1Configuration) SetTaskrouterTargetTaskqueueSid(v string)`

SetTaskrouterTargetTaskqueueSid sets TaskrouterTargetTaskqueueSid field to given value.

### HasTaskrouterTargetTaskqueueSid

`func (o *FlexV1Configuration) HasTaskrouterTargetTaskqueueSid() bool`

HasTaskrouterTargetTaskqueueSid returns a boolean if a field has been set.

### SetTaskrouterTargetTaskqueueSidNil

`func (o *FlexV1Configuration) SetTaskrouterTargetTaskqueueSidNil(b bool)`

 SetTaskrouterTargetTaskqueueSidNil sets the value for TaskrouterTargetTaskqueueSid to be an explicit nil

### UnsetTaskrouterTargetTaskqueueSid
`func (o *FlexV1Configuration) UnsetTaskrouterTargetTaskqueueSid()`

UnsetTaskrouterTargetTaskqueueSid ensures that no value is present for TaskrouterTargetTaskqueueSid, not even an explicit nil
### GetTaskrouterTargetWorkflowSid

`func (o *FlexV1Configuration) GetTaskrouterTargetWorkflowSid() string`

GetTaskrouterTargetWorkflowSid returns the TaskrouterTargetWorkflowSid field if non-nil, zero value otherwise.

### GetTaskrouterTargetWorkflowSidOk

`func (o *FlexV1Configuration) GetTaskrouterTargetWorkflowSidOk() (*string, bool)`

GetTaskrouterTargetWorkflowSidOk returns a tuple with the TaskrouterTargetWorkflowSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskrouterTargetWorkflowSid

`func (o *FlexV1Configuration) SetTaskrouterTargetWorkflowSid(v string)`

SetTaskrouterTargetWorkflowSid sets TaskrouterTargetWorkflowSid field to given value.

### HasTaskrouterTargetWorkflowSid

`func (o *FlexV1Configuration) HasTaskrouterTargetWorkflowSid() bool`

HasTaskrouterTargetWorkflowSid returns a boolean if a field has been set.

### SetTaskrouterTargetWorkflowSidNil

`func (o *FlexV1Configuration) SetTaskrouterTargetWorkflowSidNil(b bool)`

 SetTaskrouterTargetWorkflowSidNil sets the value for TaskrouterTargetWorkflowSid to be an explicit nil

### UnsetTaskrouterTargetWorkflowSid
`func (o *FlexV1Configuration) UnsetTaskrouterTargetWorkflowSid()`

UnsetTaskrouterTargetWorkflowSid ensures that no value is present for TaskrouterTargetWorkflowSid, not even an explicit nil
### GetTaskrouterTaskqueues

`func (o *FlexV1Configuration) GetTaskrouterTaskqueues() []map[string]interface{}`

GetTaskrouterTaskqueues returns the TaskrouterTaskqueues field if non-nil, zero value otherwise.

### GetTaskrouterTaskqueuesOk

`func (o *FlexV1Configuration) GetTaskrouterTaskqueuesOk() (*[]map[string]interface{}, bool)`

GetTaskrouterTaskqueuesOk returns a tuple with the TaskrouterTaskqueues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskrouterTaskqueues

`func (o *FlexV1Configuration) SetTaskrouterTaskqueues(v []map[string]interface{})`

SetTaskrouterTaskqueues sets TaskrouterTaskqueues field to given value.

### HasTaskrouterTaskqueues

`func (o *FlexV1Configuration) HasTaskrouterTaskqueues() bool`

HasTaskrouterTaskqueues returns a boolean if a field has been set.

### SetTaskrouterTaskqueuesNil

`func (o *FlexV1Configuration) SetTaskrouterTaskqueuesNil(b bool)`

 SetTaskrouterTaskqueuesNil sets the value for TaskrouterTaskqueues to be an explicit nil

### UnsetTaskrouterTaskqueues
`func (o *FlexV1Configuration) UnsetTaskrouterTaskqueues()`

UnsetTaskrouterTaskqueues ensures that no value is present for TaskrouterTaskqueues, not even an explicit nil
### GetTaskrouterWorkerAttributes

`func (o *FlexV1Configuration) GetTaskrouterWorkerAttributes() map[string]interface{}`

GetTaskrouterWorkerAttributes returns the TaskrouterWorkerAttributes field if non-nil, zero value otherwise.

### GetTaskrouterWorkerAttributesOk

`func (o *FlexV1Configuration) GetTaskrouterWorkerAttributesOk() (*map[string]interface{}, bool)`

GetTaskrouterWorkerAttributesOk returns a tuple with the TaskrouterWorkerAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskrouterWorkerAttributes

`func (o *FlexV1Configuration) SetTaskrouterWorkerAttributes(v map[string]interface{})`

SetTaskrouterWorkerAttributes sets TaskrouterWorkerAttributes field to given value.

### HasTaskrouterWorkerAttributes

`func (o *FlexV1Configuration) HasTaskrouterWorkerAttributes() bool`

HasTaskrouterWorkerAttributes returns a boolean if a field has been set.

### SetTaskrouterWorkerAttributesNil

`func (o *FlexV1Configuration) SetTaskrouterWorkerAttributesNil(b bool)`

 SetTaskrouterWorkerAttributesNil sets the value for TaskrouterWorkerAttributes to be an explicit nil

### UnsetTaskrouterWorkerAttributes
`func (o *FlexV1Configuration) UnsetTaskrouterWorkerAttributes()`

UnsetTaskrouterWorkerAttributes ensures that no value is present for TaskrouterWorkerAttributes, not even an explicit nil
### GetTaskrouterWorkerChannels

`func (o *FlexV1Configuration) GetTaskrouterWorkerChannels() map[string]interface{}`

GetTaskrouterWorkerChannels returns the TaskrouterWorkerChannels field if non-nil, zero value otherwise.

### GetTaskrouterWorkerChannelsOk

`func (o *FlexV1Configuration) GetTaskrouterWorkerChannelsOk() (*map[string]interface{}, bool)`

GetTaskrouterWorkerChannelsOk returns a tuple with the TaskrouterWorkerChannels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskrouterWorkerChannels

`func (o *FlexV1Configuration) SetTaskrouterWorkerChannels(v map[string]interface{})`

SetTaskrouterWorkerChannels sets TaskrouterWorkerChannels field to given value.

### HasTaskrouterWorkerChannels

`func (o *FlexV1Configuration) HasTaskrouterWorkerChannels() bool`

HasTaskrouterWorkerChannels returns a boolean if a field has been set.

### SetTaskrouterWorkerChannelsNil

`func (o *FlexV1Configuration) SetTaskrouterWorkerChannelsNil(b bool)`

 SetTaskrouterWorkerChannelsNil sets the value for TaskrouterWorkerChannels to be an explicit nil

### UnsetTaskrouterWorkerChannels
`func (o *FlexV1Configuration) UnsetTaskrouterWorkerChannels()`

UnsetTaskrouterWorkerChannels ensures that no value is present for TaskrouterWorkerChannels, not even an explicit nil
### GetTaskrouterWorkspaceSid

`func (o *FlexV1Configuration) GetTaskrouterWorkspaceSid() string`

GetTaskrouterWorkspaceSid returns the TaskrouterWorkspaceSid field if non-nil, zero value otherwise.

### GetTaskrouterWorkspaceSidOk

`func (o *FlexV1Configuration) GetTaskrouterWorkspaceSidOk() (*string, bool)`

GetTaskrouterWorkspaceSidOk returns a tuple with the TaskrouterWorkspaceSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskrouterWorkspaceSid

`func (o *FlexV1Configuration) SetTaskrouterWorkspaceSid(v string)`

SetTaskrouterWorkspaceSid sets TaskrouterWorkspaceSid field to given value.

### HasTaskrouterWorkspaceSid

`func (o *FlexV1Configuration) HasTaskrouterWorkspaceSid() bool`

HasTaskrouterWorkspaceSid returns a boolean if a field has been set.

### SetTaskrouterWorkspaceSidNil

`func (o *FlexV1Configuration) SetTaskrouterWorkspaceSidNil(b bool)`

 SetTaskrouterWorkspaceSidNil sets the value for TaskrouterWorkspaceSid to be an explicit nil

### UnsetTaskrouterWorkspaceSid
`func (o *FlexV1Configuration) UnsetTaskrouterWorkspaceSid()`

UnsetTaskrouterWorkspaceSid ensures that no value is present for TaskrouterWorkspaceSid, not even an explicit nil
### GetUiAttributes

`func (o *FlexV1Configuration) GetUiAttributes() map[string]interface{}`

GetUiAttributes returns the UiAttributes field if non-nil, zero value otherwise.

### GetUiAttributesOk

`func (o *FlexV1Configuration) GetUiAttributesOk() (*map[string]interface{}, bool)`

GetUiAttributesOk returns a tuple with the UiAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUiAttributes

`func (o *FlexV1Configuration) SetUiAttributes(v map[string]interface{})`

SetUiAttributes sets UiAttributes field to given value.

### HasUiAttributes

`func (o *FlexV1Configuration) HasUiAttributes() bool`

HasUiAttributes returns a boolean if a field has been set.

### SetUiAttributesNil

`func (o *FlexV1Configuration) SetUiAttributesNil(b bool)`

 SetUiAttributesNil sets the value for UiAttributes to be an explicit nil

### UnsetUiAttributes
`func (o *FlexV1Configuration) UnsetUiAttributes()`

UnsetUiAttributes ensures that no value is present for UiAttributes, not even an explicit nil
### GetUiDependencies

`func (o *FlexV1Configuration) GetUiDependencies() map[string]interface{}`

GetUiDependencies returns the UiDependencies field if non-nil, zero value otherwise.

### GetUiDependenciesOk

`func (o *FlexV1Configuration) GetUiDependenciesOk() (*map[string]interface{}, bool)`

GetUiDependenciesOk returns a tuple with the UiDependencies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUiDependencies

`func (o *FlexV1Configuration) SetUiDependencies(v map[string]interface{})`

SetUiDependencies sets UiDependencies field to given value.

### HasUiDependencies

`func (o *FlexV1Configuration) HasUiDependencies() bool`

HasUiDependencies returns a boolean if a field has been set.

### SetUiDependenciesNil

`func (o *FlexV1Configuration) SetUiDependenciesNil(b bool)`

 SetUiDependenciesNil sets the value for UiDependencies to be an explicit nil

### UnsetUiDependencies
`func (o *FlexV1Configuration) UnsetUiDependencies()`

UnsetUiDependencies ensures that no value is present for UiDependencies, not even an explicit nil
### GetUiLanguage

`func (o *FlexV1Configuration) GetUiLanguage() string`

GetUiLanguage returns the UiLanguage field if non-nil, zero value otherwise.

### GetUiLanguageOk

`func (o *FlexV1Configuration) GetUiLanguageOk() (*string, bool)`

GetUiLanguageOk returns a tuple with the UiLanguage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUiLanguage

`func (o *FlexV1Configuration) SetUiLanguage(v string)`

SetUiLanguage sets UiLanguage field to given value.

### HasUiLanguage

`func (o *FlexV1Configuration) HasUiLanguage() bool`

HasUiLanguage returns a boolean if a field has been set.

### SetUiLanguageNil

`func (o *FlexV1Configuration) SetUiLanguageNil(b bool)`

 SetUiLanguageNil sets the value for UiLanguage to be an explicit nil

### UnsetUiLanguage
`func (o *FlexV1Configuration) UnsetUiLanguage()`

UnsetUiLanguage ensures that no value is present for UiLanguage, not even an explicit nil
### GetUiVersion

`func (o *FlexV1Configuration) GetUiVersion() string`

GetUiVersion returns the UiVersion field if non-nil, zero value otherwise.

### GetUiVersionOk

`func (o *FlexV1Configuration) GetUiVersionOk() (*string, bool)`

GetUiVersionOk returns a tuple with the UiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUiVersion

`func (o *FlexV1Configuration) SetUiVersion(v string)`

SetUiVersion sets UiVersion field to given value.

### HasUiVersion

`func (o *FlexV1Configuration) HasUiVersion() bool`

HasUiVersion returns a boolean if a field has been set.

### SetUiVersionNil

`func (o *FlexV1Configuration) SetUiVersionNil(b bool)`

 SetUiVersionNil sets the value for UiVersion to be an explicit nil

### UnsetUiVersion
`func (o *FlexV1Configuration) UnsetUiVersion()`

UnsetUiVersion ensures that no value is present for UiVersion, not even an explicit nil
### GetUrl

`func (o *FlexV1Configuration) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *FlexV1Configuration) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *FlexV1Configuration) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *FlexV1Configuration) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *FlexV1Configuration) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *FlexV1Configuration) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


