# ConversationsV1ServiceServiceConfiguration

## Properties

Name | Type | Description
------------ | ------------- | -------------
**ChatServiceSid** | Pointer to **NullableString** | The unique string that identifies the resource | [optional] 
**DefaultChatServiceRoleSid** | Pointer to **NullableString** | The service role assigned to users when they are added to the service | [optional] 
**DefaultConversationCreatorRoleSid** | Pointer to **NullableString** | The role assigned to a conversation creator user when they join a new conversation | [optional] 
**DefaultConversationRoleSid** | Pointer to **NullableString** | The role assigned to users when they are added to a conversation | [optional] 
**Links** | Pointer to **map[string]interface{}** | Absolute URL to access the push notifications configuration of this service. | [optional] 
**ReachabilityEnabled** | Pointer to **NullableBool** | Whether the Reachability Indicator feature is enabled for this Conversations Service | [optional] 
**Url** | Pointer to **NullableString** | An absolute URL for this service configuration. | [optional] 

## Methods

### NewConversationsV1ServiceServiceConfiguration

`func NewConversationsV1ServiceServiceConfiguration() *ConversationsV1ServiceServiceConfiguration`

NewConversationsV1ServiceServiceConfiguration instantiates a new ConversationsV1ServiceServiceConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConversationsV1ServiceServiceConfigurationWithDefaults

`func NewConversationsV1ServiceServiceConfigurationWithDefaults() *ConversationsV1ServiceServiceConfiguration`

NewConversationsV1ServiceServiceConfigurationWithDefaults instantiates a new ConversationsV1ServiceServiceConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChatServiceSid

`func (o *ConversationsV1ServiceServiceConfiguration) GetChatServiceSid() string`

GetChatServiceSid returns the ChatServiceSid field if non-nil, zero value otherwise.

### GetChatServiceSidOk

`func (o *ConversationsV1ServiceServiceConfiguration) GetChatServiceSidOk() (*string, bool)`

GetChatServiceSidOk returns a tuple with the ChatServiceSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatServiceSid

`func (o *ConversationsV1ServiceServiceConfiguration) SetChatServiceSid(v string)`

SetChatServiceSid sets ChatServiceSid field to given value.

### HasChatServiceSid

`func (o *ConversationsV1ServiceServiceConfiguration) HasChatServiceSid() bool`

HasChatServiceSid returns a boolean if a field has been set.

### SetChatServiceSidNil

`func (o *ConversationsV1ServiceServiceConfiguration) SetChatServiceSidNil(b bool)`

 SetChatServiceSidNil sets the value for ChatServiceSid to be an explicit nil

### UnsetChatServiceSid
`func (o *ConversationsV1ServiceServiceConfiguration) UnsetChatServiceSid()`

UnsetChatServiceSid ensures that no value is present for ChatServiceSid, not even an explicit nil
### GetDefaultChatServiceRoleSid

`func (o *ConversationsV1ServiceServiceConfiguration) GetDefaultChatServiceRoleSid() string`

GetDefaultChatServiceRoleSid returns the DefaultChatServiceRoleSid field if non-nil, zero value otherwise.

### GetDefaultChatServiceRoleSidOk

`func (o *ConversationsV1ServiceServiceConfiguration) GetDefaultChatServiceRoleSidOk() (*string, bool)`

GetDefaultChatServiceRoleSidOk returns a tuple with the DefaultChatServiceRoleSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultChatServiceRoleSid

`func (o *ConversationsV1ServiceServiceConfiguration) SetDefaultChatServiceRoleSid(v string)`

SetDefaultChatServiceRoleSid sets DefaultChatServiceRoleSid field to given value.

### HasDefaultChatServiceRoleSid

`func (o *ConversationsV1ServiceServiceConfiguration) HasDefaultChatServiceRoleSid() bool`

HasDefaultChatServiceRoleSid returns a boolean if a field has been set.

### SetDefaultChatServiceRoleSidNil

`func (o *ConversationsV1ServiceServiceConfiguration) SetDefaultChatServiceRoleSidNil(b bool)`

 SetDefaultChatServiceRoleSidNil sets the value for DefaultChatServiceRoleSid to be an explicit nil

### UnsetDefaultChatServiceRoleSid
`func (o *ConversationsV1ServiceServiceConfiguration) UnsetDefaultChatServiceRoleSid()`

UnsetDefaultChatServiceRoleSid ensures that no value is present for DefaultChatServiceRoleSid, not even an explicit nil
### GetDefaultConversationCreatorRoleSid

`func (o *ConversationsV1ServiceServiceConfiguration) GetDefaultConversationCreatorRoleSid() string`

GetDefaultConversationCreatorRoleSid returns the DefaultConversationCreatorRoleSid field if non-nil, zero value otherwise.

### GetDefaultConversationCreatorRoleSidOk

`func (o *ConversationsV1ServiceServiceConfiguration) GetDefaultConversationCreatorRoleSidOk() (*string, bool)`

GetDefaultConversationCreatorRoleSidOk returns a tuple with the DefaultConversationCreatorRoleSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultConversationCreatorRoleSid

`func (o *ConversationsV1ServiceServiceConfiguration) SetDefaultConversationCreatorRoleSid(v string)`

SetDefaultConversationCreatorRoleSid sets DefaultConversationCreatorRoleSid field to given value.

### HasDefaultConversationCreatorRoleSid

`func (o *ConversationsV1ServiceServiceConfiguration) HasDefaultConversationCreatorRoleSid() bool`

HasDefaultConversationCreatorRoleSid returns a boolean if a field has been set.

### SetDefaultConversationCreatorRoleSidNil

`func (o *ConversationsV1ServiceServiceConfiguration) SetDefaultConversationCreatorRoleSidNil(b bool)`

 SetDefaultConversationCreatorRoleSidNil sets the value for DefaultConversationCreatorRoleSid to be an explicit nil

### UnsetDefaultConversationCreatorRoleSid
`func (o *ConversationsV1ServiceServiceConfiguration) UnsetDefaultConversationCreatorRoleSid()`

UnsetDefaultConversationCreatorRoleSid ensures that no value is present for DefaultConversationCreatorRoleSid, not even an explicit nil
### GetDefaultConversationRoleSid

`func (o *ConversationsV1ServiceServiceConfiguration) GetDefaultConversationRoleSid() string`

GetDefaultConversationRoleSid returns the DefaultConversationRoleSid field if non-nil, zero value otherwise.

### GetDefaultConversationRoleSidOk

`func (o *ConversationsV1ServiceServiceConfiguration) GetDefaultConversationRoleSidOk() (*string, bool)`

GetDefaultConversationRoleSidOk returns a tuple with the DefaultConversationRoleSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultConversationRoleSid

`func (o *ConversationsV1ServiceServiceConfiguration) SetDefaultConversationRoleSid(v string)`

SetDefaultConversationRoleSid sets DefaultConversationRoleSid field to given value.

### HasDefaultConversationRoleSid

`func (o *ConversationsV1ServiceServiceConfiguration) HasDefaultConversationRoleSid() bool`

HasDefaultConversationRoleSid returns a boolean if a field has been set.

### SetDefaultConversationRoleSidNil

`func (o *ConversationsV1ServiceServiceConfiguration) SetDefaultConversationRoleSidNil(b bool)`

 SetDefaultConversationRoleSidNil sets the value for DefaultConversationRoleSid to be an explicit nil

### UnsetDefaultConversationRoleSid
`func (o *ConversationsV1ServiceServiceConfiguration) UnsetDefaultConversationRoleSid()`

UnsetDefaultConversationRoleSid ensures that no value is present for DefaultConversationRoleSid, not even an explicit nil
### GetLinks

`func (o *ConversationsV1ServiceServiceConfiguration) GetLinks() map[string]interface{}`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ConversationsV1ServiceServiceConfiguration) GetLinksOk() (*map[string]interface{}, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ConversationsV1ServiceServiceConfiguration) SetLinks(v map[string]interface{})`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ConversationsV1ServiceServiceConfiguration) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### SetLinksNil

`func (o *ConversationsV1ServiceServiceConfiguration) SetLinksNil(b bool)`

 SetLinksNil sets the value for Links to be an explicit nil

### UnsetLinks
`func (o *ConversationsV1ServiceServiceConfiguration) UnsetLinks()`

UnsetLinks ensures that no value is present for Links, not even an explicit nil
### GetReachabilityEnabled

`func (o *ConversationsV1ServiceServiceConfiguration) GetReachabilityEnabled() bool`

GetReachabilityEnabled returns the ReachabilityEnabled field if non-nil, zero value otherwise.

### GetReachabilityEnabledOk

`func (o *ConversationsV1ServiceServiceConfiguration) GetReachabilityEnabledOk() (*bool, bool)`

GetReachabilityEnabledOk returns a tuple with the ReachabilityEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReachabilityEnabled

`func (o *ConversationsV1ServiceServiceConfiguration) SetReachabilityEnabled(v bool)`

SetReachabilityEnabled sets ReachabilityEnabled field to given value.

### HasReachabilityEnabled

`func (o *ConversationsV1ServiceServiceConfiguration) HasReachabilityEnabled() bool`

HasReachabilityEnabled returns a boolean if a field has been set.

### SetReachabilityEnabledNil

`func (o *ConversationsV1ServiceServiceConfiguration) SetReachabilityEnabledNil(b bool)`

 SetReachabilityEnabledNil sets the value for ReachabilityEnabled to be an explicit nil

### UnsetReachabilityEnabled
`func (o *ConversationsV1ServiceServiceConfiguration) UnsetReachabilityEnabled()`

UnsetReachabilityEnabled ensures that no value is present for ReachabilityEnabled, not even an explicit nil
### GetUrl

`func (o *ConversationsV1ServiceServiceConfiguration) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ConversationsV1ServiceServiceConfiguration) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ConversationsV1ServiceServiceConfiguration) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *ConversationsV1ServiceServiceConfiguration) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *ConversationsV1ServiceServiceConfiguration) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *ConversationsV1ServiceServiceConfiguration) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


