# IpMessagingV1Service

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **NullableString** |  | [optional] 
**ConsumptionReportInterval** | Pointer to **NullableInt32** |  | [optional] 
**DateCreated** | Pointer to **NullableTime** |  | [optional] 
**DateUpdated** | Pointer to **NullableTime** |  | [optional] 
**DefaultChannelCreatorRoleSid** | Pointer to **NullableString** |  | [optional] 
**DefaultChannelRoleSid** | Pointer to **NullableString** |  | [optional] 
**DefaultServiceRoleSid** | Pointer to **NullableString** |  | [optional] 
**FriendlyName** | Pointer to **NullableString** |  | [optional] 
**Limits** | Pointer to **map[string]interface{}** |  | [optional] 
**Links** | Pointer to **map[string]interface{}** |  | [optional] 
**Notifications** | Pointer to **map[string]interface{}** |  | [optional] 
**PostWebhookUrl** | Pointer to **NullableString** |  | [optional] 
**PreWebhookUrl** | Pointer to **NullableString** |  | [optional] 
**ReachabilityEnabled** | Pointer to **NullableBool** |  | [optional] 
**ReadStatusEnabled** | Pointer to **NullableBool** |  | [optional] 
**Sid** | Pointer to **NullableString** |  | [optional] 
**TypingIndicatorTimeout** | Pointer to **NullableInt32** |  | [optional] 
**Url** | Pointer to **NullableString** |  | [optional] 
**WebhookFilters** | Pointer to **[]string** |  | [optional] 
**WebhookMethod** | Pointer to **NullableString** |  | [optional] 
**Webhooks** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewIpMessagingV1Service

`func NewIpMessagingV1Service() *IpMessagingV1Service`

NewIpMessagingV1Service instantiates a new IpMessagingV1Service object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpMessagingV1ServiceWithDefaults

`func NewIpMessagingV1ServiceWithDefaults() *IpMessagingV1Service`

NewIpMessagingV1ServiceWithDefaults instantiates a new IpMessagingV1Service object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSid

`func (o *IpMessagingV1Service) GetAccountSid() string`

GetAccountSid returns the AccountSid field if non-nil, zero value otherwise.

### GetAccountSidOk

`func (o *IpMessagingV1Service) GetAccountSidOk() (*string, bool)`

GetAccountSidOk returns a tuple with the AccountSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSid

`func (o *IpMessagingV1Service) SetAccountSid(v string)`

SetAccountSid sets AccountSid field to given value.

### HasAccountSid

`func (o *IpMessagingV1Service) HasAccountSid() bool`

HasAccountSid returns a boolean if a field has been set.

### SetAccountSidNil

`func (o *IpMessagingV1Service) SetAccountSidNil(b bool)`

 SetAccountSidNil sets the value for AccountSid to be an explicit nil

### UnsetAccountSid
`func (o *IpMessagingV1Service) UnsetAccountSid()`

UnsetAccountSid ensures that no value is present for AccountSid, not even an explicit nil
### GetConsumptionReportInterval

`func (o *IpMessagingV1Service) GetConsumptionReportInterval() int32`

GetConsumptionReportInterval returns the ConsumptionReportInterval field if non-nil, zero value otherwise.

### GetConsumptionReportIntervalOk

`func (o *IpMessagingV1Service) GetConsumptionReportIntervalOk() (*int32, bool)`

GetConsumptionReportIntervalOk returns a tuple with the ConsumptionReportInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumptionReportInterval

`func (o *IpMessagingV1Service) SetConsumptionReportInterval(v int32)`

SetConsumptionReportInterval sets ConsumptionReportInterval field to given value.

### HasConsumptionReportInterval

`func (o *IpMessagingV1Service) HasConsumptionReportInterval() bool`

HasConsumptionReportInterval returns a boolean if a field has been set.

### SetConsumptionReportIntervalNil

`func (o *IpMessagingV1Service) SetConsumptionReportIntervalNil(b bool)`

 SetConsumptionReportIntervalNil sets the value for ConsumptionReportInterval to be an explicit nil

### UnsetConsumptionReportInterval
`func (o *IpMessagingV1Service) UnsetConsumptionReportInterval()`

UnsetConsumptionReportInterval ensures that no value is present for ConsumptionReportInterval, not even an explicit nil
### GetDateCreated

`func (o *IpMessagingV1Service) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *IpMessagingV1Service) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *IpMessagingV1Service) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *IpMessagingV1Service) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### SetDateCreatedNil

`func (o *IpMessagingV1Service) SetDateCreatedNil(b bool)`

 SetDateCreatedNil sets the value for DateCreated to be an explicit nil

### UnsetDateCreated
`func (o *IpMessagingV1Service) UnsetDateCreated()`

UnsetDateCreated ensures that no value is present for DateCreated, not even an explicit nil
### GetDateUpdated

`func (o *IpMessagingV1Service) GetDateUpdated() time.Time`

GetDateUpdated returns the DateUpdated field if non-nil, zero value otherwise.

### GetDateUpdatedOk

`func (o *IpMessagingV1Service) GetDateUpdatedOk() (*time.Time, bool)`

GetDateUpdatedOk returns a tuple with the DateUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateUpdated

`func (o *IpMessagingV1Service) SetDateUpdated(v time.Time)`

SetDateUpdated sets DateUpdated field to given value.

### HasDateUpdated

`func (o *IpMessagingV1Service) HasDateUpdated() bool`

HasDateUpdated returns a boolean if a field has been set.

### SetDateUpdatedNil

`func (o *IpMessagingV1Service) SetDateUpdatedNil(b bool)`

 SetDateUpdatedNil sets the value for DateUpdated to be an explicit nil

### UnsetDateUpdated
`func (o *IpMessagingV1Service) UnsetDateUpdated()`

UnsetDateUpdated ensures that no value is present for DateUpdated, not even an explicit nil
### GetDefaultChannelCreatorRoleSid

`func (o *IpMessagingV1Service) GetDefaultChannelCreatorRoleSid() string`

GetDefaultChannelCreatorRoleSid returns the DefaultChannelCreatorRoleSid field if non-nil, zero value otherwise.

### GetDefaultChannelCreatorRoleSidOk

`func (o *IpMessagingV1Service) GetDefaultChannelCreatorRoleSidOk() (*string, bool)`

GetDefaultChannelCreatorRoleSidOk returns a tuple with the DefaultChannelCreatorRoleSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultChannelCreatorRoleSid

`func (o *IpMessagingV1Service) SetDefaultChannelCreatorRoleSid(v string)`

SetDefaultChannelCreatorRoleSid sets DefaultChannelCreatorRoleSid field to given value.

### HasDefaultChannelCreatorRoleSid

`func (o *IpMessagingV1Service) HasDefaultChannelCreatorRoleSid() bool`

HasDefaultChannelCreatorRoleSid returns a boolean if a field has been set.

### SetDefaultChannelCreatorRoleSidNil

`func (o *IpMessagingV1Service) SetDefaultChannelCreatorRoleSidNil(b bool)`

 SetDefaultChannelCreatorRoleSidNil sets the value for DefaultChannelCreatorRoleSid to be an explicit nil

### UnsetDefaultChannelCreatorRoleSid
`func (o *IpMessagingV1Service) UnsetDefaultChannelCreatorRoleSid()`

UnsetDefaultChannelCreatorRoleSid ensures that no value is present for DefaultChannelCreatorRoleSid, not even an explicit nil
### GetDefaultChannelRoleSid

`func (o *IpMessagingV1Service) GetDefaultChannelRoleSid() string`

GetDefaultChannelRoleSid returns the DefaultChannelRoleSid field if non-nil, zero value otherwise.

### GetDefaultChannelRoleSidOk

`func (o *IpMessagingV1Service) GetDefaultChannelRoleSidOk() (*string, bool)`

GetDefaultChannelRoleSidOk returns a tuple with the DefaultChannelRoleSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultChannelRoleSid

`func (o *IpMessagingV1Service) SetDefaultChannelRoleSid(v string)`

SetDefaultChannelRoleSid sets DefaultChannelRoleSid field to given value.

### HasDefaultChannelRoleSid

`func (o *IpMessagingV1Service) HasDefaultChannelRoleSid() bool`

HasDefaultChannelRoleSid returns a boolean if a field has been set.

### SetDefaultChannelRoleSidNil

`func (o *IpMessagingV1Service) SetDefaultChannelRoleSidNil(b bool)`

 SetDefaultChannelRoleSidNil sets the value for DefaultChannelRoleSid to be an explicit nil

### UnsetDefaultChannelRoleSid
`func (o *IpMessagingV1Service) UnsetDefaultChannelRoleSid()`

UnsetDefaultChannelRoleSid ensures that no value is present for DefaultChannelRoleSid, not even an explicit nil
### GetDefaultServiceRoleSid

`func (o *IpMessagingV1Service) GetDefaultServiceRoleSid() string`

GetDefaultServiceRoleSid returns the DefaultServiceRoleSid field if non-nil, zero value otherwise.

### GetDefaultServiceRoleSidOk

`func (o *IpMessagingV1Service) GetDefaultServiceRoleSidOk() (*string, bool)`

GetDefaultServiceRoleSidOk returns a tuple with the DefaultServiceRoleSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultServiceRoleSid

`func (o *IpMessagingV1Service) SetDefaultServiceRoleSid(v string)`

SetDefaultServiceRoleSid sets DefaultServiceRoleSid field to given value.

### HasDefaultServiceRoleSid

`func (o *IpMessagingV1Service) HasDefaultServiceRoleSid() bool`

HasDefaultServiceRoleSid returns a boolean if a field has been set.

### SetDefaultServiceRoleSidNil

`func (o *IpMessagingV1Service) SetDefaultServiceRoleSidNil(b bool)`

 SetDefaultServiceRoleSidNil sets the value for DefaultServiceRoleSid to be an explicit nil

### UnsetDefaultServiceRoleSid
`func (o *IpMessagingV1Service) UnsetDefaultServiceRoleSid()`

UnsetDefaultServiceRoleSid ensures that no value is present for DefaultServiceRoleSid, not even an explicit nil
### GetFriendlyName

`func (o *IpMessagingV1Service) GetFriendlyName() string`

GetFriendlyName returns the FriendlyName field if non-nil, zero value otherwise.

### GetFriendlyNameOk

`func (o *IpMessagingV1Service) GetFriendlyNameOk() (*string, bool)`

GetFriendlyNameOk returns a tuple with the FriendlyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFriendlyName

`func (o *IpMessagingV1Service) SetFriendlyName(v string)`

SetFriendlyName sets FriendlyName field to given value.

### HasFriendlyName

`func (o *IpMessagingV1Service) HasFriendlyName() bool`

HasFriendlyName returns a boolean if a field has been set.

### SetFriendlyNameNil

`func (o *IpMessagingV1Service) SetFriendlyNameNil(b bool)`

 SetFriendlyNameNil sets the value for FriendlyName to be an explicit nil

### UnsetFriendlyName
`func (o *IpMessagingV1Service) UnsetFriendlyName()`

UnsetFriendlyName ensures that no value is present for FriendlyName, not even an explicit nil
### GetLimits

`func (o *IpMessagingV1Service) GetLimits() map[string]interface{}`

GetLimits returns the Limits field if non-nil, zero value otherwise.

### GetLimitsOk

`func (o *IpMessagingV1Service) GetLimitsOk() (*map[string]interface{}, bool)`

GetLimitsOk returns a tuple with the Limits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimits

`func (o *IpMessagingV1Service) SetLimits(v map[string]interface{})`

SetLimits sets Limits field to given value.

### HasLimits

`func (o *IpMessagingV1Service) HasLimits() bool`

HasLimits returns a boolean if a field has been set.

### SetLimitsNil

`func (o *IpMessagingV1Service) SetLimitsNil(b bool)`

 SetLimitsNil sets the value for Limits to be an explicit nil

### UnsetLimits
`func (o *IpMessagingV1Service) UnsetLimits()`

UnsetLimits ensures that no value is present for Limits, not even an explicit nil
### GetLinks

`func (o *IpMessagingV1Service) GetLinks() map[string]interface{}`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *IpMessagingV1Service) GetLinksOk() (*map[string]interface{}, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *IpMessagingV1Service) SetLinks(v map[string]interface{})`

SetLinks sets Links field to given value.

### HasLinks

`func (o *IpMessagingV1Service) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### SetLinksNil

`func (o *IpMessagingV1Service) SetLinksNil(b bool)`

 SetLinksNil sets the value for Links to be an explicit nil

### UnsetLinks
`func (o *IpMessagingV1Service) UnsetLinks()`

UnsetLinks ensures that no value is present for Links, not even an explicit nil
### GetNotifications

`func (o *IpMessagingV1Service) GetNotifications() map[string]interface{}`

GetNotifications returns the Notifications field if non-nil, zero value otherwise.

### GetNotificationsOk

`func (o *IpMessagingV1Service) GetNotificationsOk() (*map[string]interface{}, bool)`

GetNotificationsOk returns a tuple with the Notifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifications

`func (o *IpMessagingV1Service) SetNotifications(v map[string]interface{})`

SetNotifications sets Notifications field to given value.

### HasNotifications

`func (o *IpMessagingV1Service) HasNotifications() bool`

HasNotifications returns a boolean if a field has been set.

### SetNotificationsNil

`func (o *IpMessagingV1Service) SetNotificationsNil(b bool)`

 SetNotificationsNil sets the value for Notifications to be an explicit nil

### UnsetNotifications
`func (o *IpMessagingV1Service) UnsetNotifications()`

UnsetNotifications ensures that no value is present for Notifications, not even an explicit nil
### GetPostWebhookUrl

`func (o *IpMessagingV1Service) GetPostWebhookUrl() string`

GetPostWebhookUrl returns the PostWebhookUrl field if non-nil, zero value otherwise.

### GetPostWebhookUrlOk

`func (o *IpMessagingV1Service) GetPostWebhookUrlOk() (*string, bool)`

GetPostWebhookUrlOk returns a tuple with the PostWebhookUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostWebhookUrl

`func (o *IpMessagingV1Service) SetPostWebhookUrl(v string)`

SetPostWebhookUrl sets PostWebhookUrl field to given value.

### HasPostWebhookUrl

`func (o *IpMessagingV1Service) HasPostWebhookUrl() bool`

HasPostWebhookUrl returns a boolean if a field has been set.

### SetPostWebhookUrlNil

`func (o *IpMessagingV1Service) SetPostWebhookUrlNil(b bool)`

 SetPostWebhookUrlNil sets the value for PostWebhookUrl to be an explicit nil

### UnsetPostWebhookUrl
`func (o *IpMessagingV1Service) UnsetPostWebhookUrl()`

UnsetPostWebhookUrl ensures that no value is present for PostWebhookUrl, not even an explicit nil
### GetPreWebhookUrl

`func (o *IpMessagingV1Service) GetPreWebhookUrl() string`

GetPreWebhookUrl returns the PreWebhookUrl field if non-nil, zero value otherwise.

### GetPreWebhookUrlOk

`func (o *IpMessagingV1Service) GetPreWebhookUrlOk() (*string, bool)`

GetPreWebhookUrlOk returns a tuple with the PreWebhookUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreWebhookUrl

`func (o *IpMessagingV1Service) SetPreWebhookUrl(v string)`

SetPreWebhookUrl sets PreWebhookUrl field to given value.

### HasPreWebhookUrl

`func (o *IpMessagingV1Service) HasPreWebhookUrl() bool`

HasPreWebhookUrl returns a boolean if a field has been set.

### SetPreWebhookUrlNil

`func (o *IpMessagingV1Service) SetPreWebhookUrlNil(b bool)`

 SetPreWebhookUrlNil sets the value for PreWebhookUrl to be an explicit nil

### UnsetPreWebhookUrl
`func (o *IpMessagingV1Service) UnsetPreWebhookUrl()`

UnsetPreWebhookUrl ensures that no value is present for PreWebhookUrl, not even an explicit nil
### GetReachabilityEnabled

`func (o *IpMessagingV1Service) GetReachabilityEnabled() bool`

GetReachabilityEnabled returns the ReachabilityEnabled field if non-nil, zero value otherwise.

### GetReachabilityEnabledOk

`func (o *IpMessagingV1Service) GetReachabilityEnabledOk() (*bool, bool)`

GetReachabilityEnabledOk returns a tuple with the ReachabilityEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReachabilityEnabled

`func (o *IpMessagingV1Service) SetReachabilityEnabled(v bool)`

SetReachabilityEnabled sets ReachabilityEnabled field to given value.

### HasReachabilityEnabled

`func (o *IpMessagingV1Service) HasReachabilityEnabled() bool`

HasReachabilityEnabled returns a boolean if a field has been set.

### SetReachabilityEnabledNil

`func (o *IpMessagingV1Service) SetReachabilityEnabledNil(b bool)`

 SetReachabilityEnabledNil sets the value for ReachabilityEnabled to be an explicit nil

### UnsetReachabilityEnabled
`func (o *IpMessagingV1Service) UnsetReachabilityEnabled()`

UnsetReachabilityEnabled ensures that no value is present for ReachabilityEnabled, not even an explicit nil
### GetReadStatusEnabled

`func (o *IpMessagingV1Service) GetReadStatusEnabled() bool`

GetReadStatusEnabled returns the ReadStatusEnabled field if non-nil, zero value otherwise.

### GetReadStatusEnabledOk

`func (o *IpMessagingV1Service) GetReadStatusEnabledOk() (*bool, bool)`

GetReadStatusEnabledOk returns a tuple with the ReadStatusEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadStatusEnabled

`func (o *IpMessagingV1Service) SetReadStatusEnabled(v bool)`

SetReadStatusEnabled sets ReadStatusEnabled field to given value.

### HasReadStatusEnabled

`func (o *IpMessagingV1Service) HasReadStatusEnabled() bool`

HasReadStatusEnabled returns a boolean if a field has been set.

### SetReadStatusEnabledNil

`func (o *IpMessagingV1Service) SetReadStatusEnabledNil(b bool)`

 SetReadStatusEnabledNil sets the value for ReadStatusEnabled to be an explicit nil

### UnsetReadStatusEnabled
`func (o *IpMessagingV1Service) UnsetReadStatusEnabled()`

UnsetReadStatusEnabled ensures that no value is present for ReadStatusEnabled, not even an explicit nil
### GetSid

`func (o *IpMessagingV1Service) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *IpMessagingV1Service) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *IpMessagingV1Service) SetSid(v string)`

SetSid sets Sid field to given value.

### HasSid

`func (o *IpMessagingV1Service) HasSid() bool`

HasSid returns a boolean if a field has been set.

### SetSidNil

`func (o *IpMessagingV1Service) SetSidNil(b bool)`

 SetSidNil sets the value for Sid to be an explicit nil

### UnsetSid
`func (o *IpMessagingV1Service) UnsetSid()`

UnsetSid ensures that no value is present for Sid, not even an explicit nil
### GetTypingIndicatorTimeout

`func (o *IpMessagingV1Service) GetTypingIndicatorTimeout() int32`

GetTypingIndicatorTimeout returns the TypingIndicatorTimeout field if non-nil, zero value otherwise.

### GetTypingIndicatorTimeoutOk

`func (o *IpMessagingV1Service) GetTypingIndicatorTimeoutOk() (*int32, bool)`

GetTypingIndicatorTimeoutOk returns a tuple with the TypingIndicatorTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypingIndicatorTimeout

`func (o *IpMessagingV1Service) SetTypingIndicatorTimeout(v int32)`

SetTypingIndicatorTimeout sets TypingIndicatorTimeout field to given value.

### HasTypingIndicatorTimeout

`func (o *IpMessagingV1Service) HasTypingIndicatorTimeout() bool`

HasTypingIndicatorTimeout returns a boolean if a field has been set.

### SetTypingIndicatorTimeoutNil

`func (o *IpMessagingV1Service) SetTypingIndicatorTimeoutNil(b bool)`

 SetTypingIndicatorTimeoutNil sets the value for TypingIndicatorTimeout to be an explicit nil

### UnsetTypingIndicatorTimeout
`func (o *IpMessagingV1Service) UnsetTypingIndicatorTimeout()`

UnsetTypingIndicatorTimeout ensures that no value is present for TypingIndicatorTimeout, not even an explicit nil
### GetUrl

`func (o *IpMessagingV1Service) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *IpMessagingV1Service) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *IpMessagingV1Service) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *IpMessagingV1Service) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *IpMessagingV1Service) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *IpMessagingV1Service) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetWebhookFilters

`func (o *IpMessagingV1Service) GetWebhookFilters() []string`

GetWebhookFilters returns the WebhookFilters field if non-nil, zero value otherwise.

### GetWebhookFiltersOk

`func (o *IpMessagingV1Service) GetWebhookFiltersOk() (*[]string, bool)`

GetWebhookFiltersOk returns a tuple with the WebhookFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookFilters

`func (o *IpMessagingV1Service) SetWebhookFilters(v []string)`

SetWebhookFilters sets WebhookFilters field to given value.

### HasWebhookFilters

`func (o *IpMessagingV1Service) HasWebhookFilters() bool`

HasWebhookFilters returns a boolean if a field has been set.

### SetWebhookFiltersNil

`func (o *IpMessagingV1Service) SetWebhookFiltersNil(b bool)`

 SetWebhookFiltersNil sets the value for WebhookFilters to be an explicit nil

### UnsetWebhookFilters
`func (o *IpMessagingV1Service) UnsetWebhookFilters()`

UnsetWebhookFilters ensures that no value is present for WebhookFilters, not even an explicit nil
### GetWebhookMethod

`func (o *IpMessagingV1Service) GetWebhookMethod() string`

GetWebhookMethod returns the WebhookMethod field if non-nil, zero value otherwise.

### GetWebhookMethodOk

`func (o *IpMessagingV1Service) GetWebhookMethodOk() (*string, bool)`

GetWebhookMethodOk returns a tuple with the WebhookMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookMethod

`func (o *IpMessagingV1Service) SetWebhookMethod(v string)`

SetWebhookMethod sets WebhookMethod field to given value.

### HasWebhookMethod

`func (o *IpMessagingV1Service) HasWebhookMethod() bool`

HasWebhookMethod returns a boolean if a field has been set.

### SetWebhookMethodNil

`func (o *IpMessagingV1Service) SetWebhookMethodNil(b bool)`

 SetWebhookMethodNil sets the value for WebhookMethod to be an explicit nil

### UnsetWebhookMethod
`func (o *IpMessagingV1Service) UnsetWebhookMethod()`

UnsetWebhookMethod ensures that no value is present for WebhookMethod, not even an explicit nil
### GetWebhooks

`func (o *IpMessagingV1Service) GetWebhooks() map[string]interface{}`

GetWebhooks returns the Webhooks field if non-nil, zero value otherwise.

### GetWebhooksOk

`func (o *IpMessagingV1Service) GetWebhooksOk() (*map[string]interface{}, bool)`

GetWebhooksOk returns a tuple with the Webhooks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhooks

`func (o *IpMessagingV1Service) SetWebhooks(v map[string]interface{})`

SetWebhooks sets Webhooks field to given value.

### HasWebhooks

`func (o *IpMessagingV1Service) HasWebhooks() bool`

HasWebhooks returns a boolean if a field has been set.

### SetWebhooksNil

`func (o *IpMessagingV1Service) SetWebhooksNil(b bool)`

 SetWebhooksNil sets the value for Webhooks to be an explicit nil

### UnsetWebhooks
`func (o *IpMessagingV1Service) UnsetWebhooks()`

UnsetWebhooks ensures that no value is present for Webhooks, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


