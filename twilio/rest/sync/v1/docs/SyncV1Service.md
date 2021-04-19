# SyncV1Service

## Properties

Name | Type | Description
------------ | ------------- | -------------
**AccountSid** | Pointer to **NullableString** | The SID of the Account that created the resource | [optional] 
**AclEnabled** | Pointer to **NullableBool** | Whether token identities in the Service must be granted access to Sync objects by using the Permissions resource | [optional] 
**DateCreated** | Pointer to **NullableTime** | The ISO 8601 date and time in GMT when the resource was created | [optional] 
**DateUpdated** | Pointer to **NullableTime** | The ISO 8601 date and time in GMT when the resource was last updated | [optional] 
**FriendlyName** | Pointer to **NullableString** | The string that you assigned to describe the resource | [optional] 
**Links** | Pointer to **map[string]interface{}** | The URLs of related resources | [optional] 
**ReachabilityDebouncingEnabled** | Pointer to **NullableBool** | Whether every endpoint_disconnected event occurs after a configurable delay | [optional] 
**ReachabilityDebouncingWindow** | Pointer to **NullableInt32** | The reachability event delay in milliseconds | [optional] 
**ReachabilityWebhooksEnabled** | Pointer to **NullableBool** | Whether the service instance calls webhook_url when client endpoints connect to Sync | [optional] 
**Sid** | Pointer to **NullableString** | The unique string that identifies the resource | [optional] 
**UniqueName** | Pointer to **NullableString** | An application-defined string that uniquely identifies the resource | [optional] 
**Url** | Pointer to **NullableString** | The absolute URL of the Service resource | [optional] 
**WebhookUrl** | Pointer to **NullableString** | The URL we call when Sync objects are manipulated | [optional] 
**WebhooksFromRestEnabled** | Pointer to **NullableBool** | Whether the Service instance should call webhook_url when the REST API is used to update Sync objects | [optional] 

## Methods

### NewSyncV1Service

`func NewSyncV1Service() *SyncV1Service`

NewSyncV1Service instantiates a new SyncV1Service object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyncV1ServiceWithDefaults

`func NewSyncV1ServiceWithDefaults() *SyncV1Service`

NewSyncV1ServiceWithDefaults instantiates a new SyncV1Service object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSid

`func (o *SyncV1Service) GetAccountSid() string`

GetAccountSid returns the AccountSid field if non-nil, zero value otherwise.

### GetAccountSidOk

`func (o *SyncV1Service) GetAccountSidOk() (*string, bool)`

GetAccountSidOk returns a tuple with the AccountSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSid

`func (o *SyncV1Service) SetAccountSid(v string)`

SetAccountSid sets AccountSid field to given value.

### HasAccountSid

`func (o *SyncV1Service) HasAccountSid() bool`

HasAccountSid returns a boolean if a field has been set.

### SetAccountSidNil

`func (o *SyncV1Service) SetAccountSidNil(b bool)`

 SetAccountSidNil sets the value for AccountSid to be an explicit nil

### UnsetAccountSid
`func (o *SyncV1Service) UnsetAccountSid()`

UnsetAccountSid ensures that no value is present for AccountSid, not even an explicit nil
### GetAclEnabled

`func (o *SyncV1Service) GetAclEnabled() bool`

GetAclEnabled returns the AclEnabled field if non-nil, zero value otherwise.

### GetAclEnabledOk

`func (o *SyncV1Service) GetAclEnabledOk() (*bool, bool)`

GetAclEnabledOk returns a tuple with the AclEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAclEnabled

`func (o *SyncV1Service) SetAclEnabled(v bool)`

SetAclEnabled sets AclEnabled field to given value.

### HasAclEnabled

`func (o *SyncV1Service) HasAclEnabled() bool`

HasAclEnabled returns a boolean if a field has been set.

### SetAclEnabledNil

`func (o *SyncV1Service) SetAclEnabledNil(b bool)`

 SetAclEnabledNil sets the value for AclEnabled to be an explicit nil

### UnsetAclEnabled
`func (o *SyncV1Service) UnsetAclEnabled()`

UnsetAclEnabled ensures that no value is present for AclEnabled, not even an explicit nil
### GetDateCreated

`func (o *SyncV1Service) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *SyncV1Service) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *SyncV1Service) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *SyncV1Service) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### SetDateCreatedNil

`func (o *SyncV1Service) SetDateCreatedNil(b bool)`

 SetDateCreatedNil sets the value for DateCreated to be an explicit nil

### UnsetDateCreated
`func (o *SyncV1Service) UnsetDateCreated()`

UnsetDateCreated ensures that no value is present for DateCreated, not even an explicit nil
### GetDateUpdated

`func (o *SyncV1Service) GetDateUpdated() time.Time`

GetDateUpdated returns the DateUpdated field if non-nil, zero value otherwise.

### GetDateUpdatedOk

`func (o *SyncV1Service) GetDateUpdatedOk() (*time.Time, bool)`

GetDateUpdatedOk returns a tuple with the DateUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateUpdated

`func (o *SyncV1Service) SetDateUpdated(v time.Time)`

SetDateUpdated sets DateUpdated field to given value.

### HasDateUpdated

`func (o *SyncV1Service) HasDateUpdated() bool`

HasDateUpdated returns a boolean if a field has been set.

### SetDateUpdatedNil

`func (o *SyncV1Service) SetDateUpdatedNil(b bool)`

 SetDateUpdatedNil sets the value for DateUpdated to be an explicit nil

### UnsetDateUpdated
`func (o *SyncV1Service) UnsetDateUpdated()`

UnsetDateUpdated ensures that no value is present for DateUpdated, not even an explicit nil
### GetFriendlyName

`func (o *SyncV1Service) GetFriendlyName() string`

GetFriendlyName returns the FriendlyName field if non-nil, zero value otherwise.

### GetFriendlyNameOk

`func (o *SyncV1Service) GetFriendlyNameOk() (*string, bool)`

GetFriendlyNameOk returns a tuple with the FriendlyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFriendlyName

`func (o *SyncV1Service) SetFriendlyName(v string)`

SetFriendlyName sets FriendlyName field to given value.

### HasFriendlyName

`func (o *SyncV1Service) HasFriendlyName() bool`

HasFriendlyName returns a boolean if a field has been set.

### SetFriendlyNameNil

`func (o *SyncV1Service) SetFriendlyNameNil(b bool)`

 SetFriendlyNameNil sets the value for FriendlyName to be an explicit nil

### UnsetFriendlyName
`func (o *SyncV1Service) UnsetFriendlyName()`

UnsetFriendlyName ensures that no value is present for FriendlyName, not even an explicit nil
### GetLinks

`func (o *SyncV1Service) GetLinks() map[string]interface{}`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *SyncV1Service) GetLinksOk() (*map[string]interface{}, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *SyncV1Service) SetLinks(v map[string]interface{})`

SetLinks sets Links field to given value.

### HasLinks

`func (o *SyncV1Service) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### SetLinksNil

`func (o *SyncV1Service) SetLinksNil(b bool)`

 SetLinksNil sets the value for Links to be an explicit nil

### UnsetLinks
`func (o *SyncV1Service) UnsetLinks()`

UnsetLinks ensures that no value is present for Links, not even an explicit nil
### GetReachabilityDebouncingEnabled

`func (o *SyncV1Service) GetReachabilityDebouncingEnabled() bool`

GetReachabilityDebouncingEnabled returns the ReachabilityDebouncingEnabled field if non-nil, zero value otherwise.

### GetReachabilityDebouncingEnabledOk

`func (o *SyncV1Service) GetReachabilityDebouncingEnabledOk() (*bool, bool)`

GetReachabilityDebouncingEnabledOk returns a tuple with the ReachabilityDebouncingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReachabilityDebouncingEnabled

`func (o *SyncV1Service) SetReachabilityDebouncingEnabled(v bool)`

SetReachabilityDebouncingEnabled sets ReachabilityDebouncingEnabled field to given value.

### HasReachabilityDebouncingEnabled

`func (o *SyncV1Service) HasReachabilityDebouncingEnabled() bool`

HasReachabilityDebouncingEnabled returns a boolean if a field has been set.

### SetReachabilityDebouncingEnabledNil

`func (o *SyncV1Service) SetReachabilityDebouncingEnabledNil(b bool)`

 SetReachabilityDebouncingEnabledNil sets the value for ReachabilityDebouncingEnabled to be an explicit nil

### UnsetReachabilityDebouncingEnabled
`func (o *SyncV1Service) UnsetReachabilityDebouncingEnabled()`

UnsetReachabilityDebouncingEnabled ensures that no value is present for ReachabilityDebouncingEnabled, not even an explicit nil
### GetReachabilityDebouncingWindow

`func (o *SyncV1Service) GetReachabilityDebouncingWindow() int32`

GetReachabilityDebouncingWindow returns the ReachabilityDebouncingWindow field if non-nil, zero value otherwise.

### GetReachabilityDebouncingWindowOk

`func (o *SyncV1Service) GetReachabilityDebouncingWindowOk() (*int32, bool)`

GetReachabilityDebouncingWindowOk returns a tuple with the ReachabilityDebouncingWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReachabilityDebouncingWindow

`func (o *SyncV1Service) SetReachabilityDebouncingWindow(v int32)`

SetReachabilityDebouncingWindow sets ReachabilityDebouncingWindow field to given value.

### HasReachabilityDebouncingWindow

`func (o *SyncV1Service) HasReachabilityDebouncingWindow() bool`

HasReachabilityDebouncingWindow returns a boolean if a field has been set.

### SetReachabilityDebouncingWindowNil

`func (o *SyncV1Service) SetReachabilityDebouncingWindowNil(b bool)`

 SetReachabilityDebouncingWindowNil sets the value for ReachabilityDebouncingWindow to be an explicit nil

### UnsetReachabilityDebouncingWindow
`func (o *SyncV1Service) UnsetReachabilityDebouncingWindow()`

UnsetReachabilityDebouncingWindow ensures that no value is present for ReachabilityDebouncingWindow, not even an explicit nil
### GetReachabilityWebhooksEnabled

`func (o *SyncV1Service) GetReachabilityWebhooksEnabled() bool`

GetReachabilityWebhooksEnabled returns the ReachabilityWebhooksEnabled field if non-nil, zero value otherwise.

### GetReachabilityWebhooksEnabledOk

`func (o *SyncV1Service) GetReachabilityWebhooksEnabledOk() (*bool, bool)`

GetReachabilityWebhooksEnabledOk returns a tuple with the ReachabilityWebhooksEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReachabilityWebhooksEnabled

`func (o *SyncV1Service) SetReachabilityWebhooksEnabled(v bool)`

SetReachabilityWebhooksEnabled sets ReachabilityWebhooksEnabled field to given value.

### HasReachabilityWebhooksEnabled

`func (o *SyncV1Service) HasReachabilityWebhooksEnabled() bool`

HasReachabilityWebhooksEnabled returns a boolean if a field has been set.

### SetReachabilityWebhooksEnabledNil

`func (o *SyncV1Service) SetReachabilityWebhooksEnabledNil(b bool)`

 SetReachabilityWebhooksEnabledNil sets the value for ReachabilityWebhooksEnabled to be an explicit nil

### UnsetReachabilityWebhooksEnabled
`func (o *SyncV1Service) UnsetReachabilityWebhooksEnabled()`

UnsetReachabilityWebhooksEnabled ensures that no value is present for ReachabilityWebhooksEnabled, not even an explicit nil
### GetSid

`func (o *SyncV1Service) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *SyncV1Service) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *SyncV1Service) SetSid(v string)`

SetSid sets Sid field to given value.

### HasSid

`func (o *SyncV1Service) HasSid() bool`

HasSid returns a boolean if a field has been set.

### SetSidNil

`func (o *SyncV1Service) SetSidNil(b bool)`

 SetSidNil sets the value for Sid to be an explicit nil

### UnsetSid
`func (o *SyncV1Service) UnsetSid()`

UnsetSid ensures that no value is present for Sid, not even an explicit nil
### GetUniqueName

`func (o *SyncV1Service) GetUniqueName() string`

GetUniqueName returns the UniqueName field if non-nil, zero value otherwise.

### GetUniqueNameOk

`func (o *SyncV1Service) GetUniqueNameOk() (*string, bool)`

GetUniqueNameOk returns a tuple with the UniqueName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueName

`func (o *SyncV1Service) SetUniqueName(v string)`

SetUniqueName sets UniqueName field to given value.

### HasUniqueName

`func (o *SyncV1Service) HasUniqueName() bool`

HasUniqueName returns a boolean if a field has been set.

### SetUniqueNameNil

`func (o *SyncV1Service) SetUniqueNameNil(b bool)`

 SetUniqueNameNil sets the value for UniqueName to be an explicit nil

### UnsetUniqueName
`func (o *SyncV1Service) UnsetUniqueName()`

UnsetUniqueName ensures that no value is present for UniqueName, not even an explicit nil
### GetUrl

`func (o *SyncV1Service) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *SyncV1Service) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *SyncV1Service) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *SyncV1Service) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *SyncV1Service) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *SyncV1Service) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetWebhookUrl

`func (o *SyncV1Service) GetWebhookUrl() string`

GetWebhookUrl returns the WebhookUrl field if non-nil, zero value otherwise.

### GetWebhookUrlOk

`func (o *SyncV1Service) GetWebhookUrlOk() (*string, bool)`

GetWebhookUrlOk returns a tuple with the WebhookUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookUrl

`func (o *SyncV1Service) SetWebhookUrl(v string)`

SetWebhookUrl sets WebhookUrl field to given value.

### HasWebhookUrl

`func (o *SyncV1Service) HasWebhookUrl() bool`

HasWebhookUrl returns a boolean if a field has been set.

### SetWebhookUrlNil

`func (o *SyncV1Service) SetWebhookUrlNil(b bool)`

 SetWebhookUrlNil sets the value for WebhookUrl to be an explicit nil

### UnsetWebhookUrl
`func (o *SyncV1Service) UnsetWebhookUrl()`

UnsetWebhookUrl ensures that no value is present for WebhookUrl, not even an explicit nil
### GetWebhooksFromRestEnabled

`func (o *SyncV1Service) GetWebhooksFromRestEnabled() bool`

GetWebhooksFromRestEnabled returns the WebhooksFromRestEnabled field if non-nil, zero value otherwise.

### GetWebhooksFromRestEnabledOk

`func (o *SyncV1Service) GetWebhooksFromRestEnabledOk() (*bool, bool)`

GetWebhooksFromRestEnabledOk returns a tuple with the WebhooksFromRestEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhooksFromRestEnabled

`func (o *SyncV1Service) SetWebhooksFromRestEnabled(v bool)`

SetWebhooksFromRestEnabled sets WebhooksFromRestEnabled field to given value.

### HasWebhooksFromRestEnabled

`func (o *SyncV1Service) HasWebhooksFromRestEnabled() bool`

HasWebhooksFromRestEnabled returns a boolean if a field has been set.

### SetWebhooksFromRestEnabledNil

`func (o *SyncV1Service) SetWebhooksFromRestEnabledNil(b bool)`

 SetWebhooksFromRestEnabledNil sets the value for WebhooksFromRestEnabled to be an explicit nil

### UnsetWebhooksFromRestEnabled
`func (o *SyncV1Service) UnsetWebhooksFromRestEnabled()`

UnsetWebhooksFromRestEnabled ensures that no value is present for WebhooksFromRestEnabled, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


