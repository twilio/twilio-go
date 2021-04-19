# ChatV2Service

## Properties

Name | Type | Description
------------ | ------------- | -------------
**AccountSid** | Pointer to **NullableString** | The SID of the Account that created the resource | [optional] 
**ConsumptionReportInterval** | Pointer to **NullableInt32** | DEPRECATED | [optional] 
**DateCreated** | Pointer to **NullableTime** | The RFC 2822 date and time in GMT when the resource was created | [optional] 
**DateUpdated** | Pointer to **NullableTime** | The RFC 2822 date and time in GMT when the resource was last updated | [optional] 
**DefaultChannelCreatorRoleSid** | Pointer to **NullableString** | The channel role assigned to a channel creator when they join a new channel | [optional] 
**DefaultChannelRoleSid** | Pointer to **NullableString** | The channel role assigned to users when they are added to a channel | [optional] 
**DefaultServiceRoleSid** | Pointer to **NullableString** | The service role assigned to users when they are added to the service | [optional] 
**FriendlyName** | Pointer to **NullableString** | The string that you assigned to describe the resource | [optional] 
**Limits** | Pointer to **map[string]interface{}** | An object that describes the limits of the service instance | [optional] 
**Links** | Pointer to **map[string]interface{}** | The absolute URLs of the Service&#39;s Channels, Roles, and Users | [optional] 
**Media** | Pointer to **map[string]interface{}** | The properties of the media that the service supports | [optional] 
**Notifications** | Pointer to **map[string]interface{}** | The notification configuration for the Service instance | [optional] 
**PostWebhookRetryCount** | Pointer to **NullableInt32** | The number of times calls to the &#x60;post_webhook_url&#x60; will be retried | [optional] 
**PostWebhookUrl** | Pointer to **NullableString** | The URL for post-event webhooks | [optional] 
**PreWebhookRetryCount** | Pointer to **NullableInt32** | Count of times webhook will be retried in case of timeout or 429/503/504 HTTP responses | [optional] 
**PreWebhookUrl** | Pointer to **NullableString** | The webhook URL for pre-event webhooks | [optional] 
**ReachabilityEnabled** | Pointer to **NullableBool** | Whether the Reachability Indicator feature is enabled for this Service instance | [optional] 
**ReadStatusEnabled** | Pointer to **NullableBool** | Whether the Message Consumption Horizon feature is enabled | [optional] 
**Sid** | Pointer to **NullableString** | The unique string that identifies the resource | [optional] 
**TypingIndicatorTimeout** | Pointer to **NullableInt32** | How long in seconds to wait before assuming the user is no longer typing | [optional] 
**Url** | Pointer to **NullableString** | The absolute URL of the Service resource | [optional] 
**WebhookFilters** | Pointer to **[]string** | The list of webhook events that are enabled for this Service instance | [optional] 
**WebhookMethod** | Pointer to **NullableString** | The HTTP method  to use for both PRE and POST webhooks | [optional] 

## Methods

### NewChatV2Service

`func NewChatV2Service() *ChatV2Service`

NewChatV2Service instantiates a new ChatV2Service object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChatV2ServiceWithDefaults

`func NewChatV2ServiceWithDefaults() *ChatV2Service`

NewChatV2ServiceWithDefaults instantiates a new ChatV2Service object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSid

`func (o *ChatV2Service) GetAccountSid() string`

GetAccountSid returns the AccountSid field if non-nil, zero value otherwise.

### GetAccountSidOk

`func (o *ChatV2Service) GetAccountSidOk() (*string, bool)`

GetAccountSidOk returns a tuple with the AccountSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSid

`func (o *ChatV2Service) SetAccountSid(v string)`

SetAccountSid sets AccountSid field to given value.

### HasAccountSid

`func (o *ChatV2Service) HasAccountSid() bool`

HasAccountSid returns a boolean if a field has been set.

### SetAccountSidNil

`func (o *ChatV2Service) SetAccountSidNil(b bool)`

 SetAccountSidNil sets the value for AccountSid to be an explicit nil

### UnsetAccountSid
`func (o *ChatV2Service) UnsetAccountSid()`

UnsetAccountSid ensures that no value is present for AccountSid, not even an explicit nil
### GetConsumptionReportInterval

`func (o *ChatV2Service) GetConsumptionReportInterval() int32`

GetConsumptionReportInterval returns the ConsumptionReportInterval field if non-nil, zero value otherwise.

### GetConsumptionReportIntervalOk

`func (o *ChatV2Service) GetConsumptionReportIntervalOk() (*int32, bool)`

GetConsumptionReportIntervalOk returns a tuple with the ConsumptionReportInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumptionReportInterval

`func (o *ChatV2Service) SetConsumptionReportInterval(v int32)`

SetConsumptionReportInterval sets ConsumptionReportInterval field to given value.

### HasConsumptionReportInterval

`func (o *ChatV2Service) HasConsumptionReportInterval() bool`

HasConsumptionReportInterval returns a boolean if a field has been set.

### SetConsumptionReportIntervalNil

`func (o *ChatV2Service) SetConsumptionReportIntervalNil(b bool)`

 SetConsumptionReportIntervalNil sets the value for ConsumptionReportInterval to be an explicit nil

### UnsetConsumptionReportInterval
`func (o *ChatV2Service) UnsetConsumptionReportInterval()`

UnsetConsumptionReportInterval ensures that no value is present for ConsumptionReportInterval, not even an explicit nil
### GetDateCreated

`func (o *ChatV2Service) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *ChatV2Service) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *ChatV2Service) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *ChatV2Service) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### SetDateCreatedNil

`func (o *ChatV2Service) SetDateCreatedNil(b bool)`

 SetDateCreatedNil sets the value for DateCreated to be an explicit nil

### UnsetDateCreated
`func (o *ChatV2Service) UnsetDateCreated()`

UnsetDateCreated ensures that no value is present for DateCreated, not even an explicit nil
### GetDateUpdated

`func (o *ChatV2Service) GetDateUpdated() time.Time`

GetDateUpdated returns the DateUpdated field if non-nil, zero value otherwise.

### GetDateUpdatedOk

`func (o *ChatV2Service) GetDateUpdatedOk() (*time.Time, bool)`

GetDateUpdatedOk returns a tuple with the DateUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateUpdated

`func (o *ChatV2Service) SetDateUpdated(v time.Time)`

SetDateUpdated sets DateUpdated field to given value.

### HasDateUpdated

`func (o *ChatV2Service) HasDateUpdated() bool`

HasDateUpdated returns a boolean if a field has been set.

### SetDateUpdatedNil

`func (o *ChatV2Service) SetDateUpdatedNil(b bool)`

 SetDateUpdatedNil sets the value for DateUpdated to be an explicit nil

### UnsetDateUpdated
`func (o *ChatV2Service) UnsetDateUpdated()`

UnsetDateUpdated ensures that no value is present for DateUpdated, not even an explicit nil
### GetDefaultChannelCreatorRoleSid

`func (o *ChatV2Service) GetDefaultChannelCreatorRoleSid() string`

GetDefaultChannelCreatorRoleSid returns the DefaultChannelCreatorRoleSid field if non-nil, zero value otherwise.

### GetDefaultChannelCreatorRoleSidOk

`func (o *ChatV2Service) GetDefaultChannelCreatorRoleSidOk() (*string, bool)`

GetDefaultChannelCreatorRoleSidOk returns a tuple with the DefaultChannelCreatorRoleSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultChannelCreatorRoleSid

`func (o *ChatV2Service) SetDefaultChannelCreatorRoleSid(v string)`

SetDefaultChannelCreatorRoleSid sets DefaultChannelCreatorRoleSid field to given value.

### HasDefaultChannelCreatorRoleSid

`func (o *ChatV2Service) HasDefaultChannelCreatorRoleSid() bool`

HasDefaultChannelCreatorRoleSid returns a boolean if a field has been set.

### SetDefaultChannelCreatorRoleSidNil

`func (o *ChatV2Service) SetDefaultChannelCreatorRoleSidNil(b bool)`

 SetDefaultChannelCreatorRoleSidNil sets the value for DefaultChannelCreatorRoleSid to be an explicit nil

### UnsetDefaultChannelCreatorRoleSid
`func (o *ChatV2Service) UnsetDefaultChannelCreatorRoleSid()`

UnsetDefaultChannelCreatorRoleSid ensures that no value is present for DefaultChannelCreatorRoleSid, not even an explicit nil
### GetDefaultChannelRoleSid

`func (o *ChatV2Service) GetDefaultChannelRoleSid() string`

GetDefaultChannelRoleSid returns the DefaultChannelRoleSid field if non-nil, zero value otherwise.

### GetDefaultChannelRoleSidOk

`func (o *ChatV2Service) GetDefaultChannelRoleSidOk() (*string, bool)`

GetDefaultChannelRoleSidOk returns a tuple with the DefaultChannelRoleSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultChannelRoleSid

`func (o *ChatV2Service) SetDefaultChannelRoleSid(v string)`

SetDefaultChannelRoleSid sets DefaultChannelRoleSid field to given value.

### HasDefaultChannelRoleSid

`func (o *ChatV2Service) HasDefaultChannelRoleSid() bool`

HasDefaultChannelRoleSid returns a boolean if a field has been set.

### SetDefaultChannelRoleSidNil

`func (o *ChatV2Service) SetDefaultChannelRoleSidNil(b bool)`

 SetDefaultChannelRoleSidNil sets the value for DefaultChannelRoleSid to be an explicit nil

### UnsetDefaultChannelRoleSid
`func (o *ChatV2Service) UnsetDefaultChannelRoleSid()`

UnsetDefaultChannelRoleSid ensures that no value is present for DefaultChannelRoleSid, not even an explicit nil
### GetDefaultServiceRoleSid

`func (o *ChatV2Service) GetDefaultServiceRoleSid() string`

GetDefaultServiceRoleSid returns the DefaultServiceRoleSid field if non-nil, zero value otherwise.

### GetDefaultServiceRoleSidOk

`func (o *ChatV2Service) GetDefaultServiceRoleSidOk() (*string, bool)`

GetDefaultServiceRoleSidOk returns a tuple with the DefaultServiceRoleSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultServiceRoleSid

`func (o *ChatV2Service) SetDefaultServiceRoleSid(v string)`

SetDefaultServiceRoleSid sets DefaultServiceRoleSid field to given value.

### HasDefaultServiceRoleSid

`func (o *ChatV2Service) HasDefaultServiceRoleSid() bool`

HasDefaultServiceRoleSid returns a boolean if a field has been set.

### SetDefaultServiceRoleSidNil

`func (o *ChatV2Service) SetDefaultServiceRoleSidNil(b bool)`

 SetDefaultServiceRoleSidNil sets the value for DefaultServiceRoleSid to be an explicit nil

### UnsetDefaultServiceRoleSid
`func (o *ChatV2Service) UnsetDefaultServiceRoleSid()`

UnsetDefaultServiceRoleSid ensures that no value is present for DefaultServiceRoleSid, not even an explicit nil
### GetFriendlyName

`func (o *ChatV2Service) GetFriendlyName() string`

GetFriendlyName returns the FriendlyName field if non-nil, zero value otherwise.

### GetFriendlyNameOk

`func (o *ChatV2Service) GetFriendlyNameOk() (*string, bool)`

GetFriendlyNameOk returns a tuple with the FriendlyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFriendlyName

`func (o *ChatV2Service) SetFriendlyName(v string)`

SetFriendlyName sets FriendlyName field to given value.

### HasFriendlyName

`func (o *ChatV2Service) HasFriendlyName() bool`

HasFriendlyName returns a boolean if a field has been set.

### SetFriendlyNameNil

`func (o *ChatV2Service) SetFriendlyNameNil(b bool)`

 SetFriendlyNameNil sets the value for FriendlyName to be an explicit nil

### UnsetFriendlyName
`func (o *ChatV2Service) UnsetFriendlyName()`

UnsetFriendlyName ensures that no value is present for FriendlyName, not even an explicit nil
### GetLimits

`func (o *ChatV2Service) GetLimits() map[string]interface{}`

GetLimits returns the Limits field if non-nil, zero value otherwise.

### GetLimitsOk

`func (o *ChatV2Service) GetLimitsOk() (*map[string]interface{}, bool)`

GetLimitsOk returns a tuple with the Limits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimits

`func (o *ChatV2Service) SetLimits(v map[string]interface{})`

SetLimits sets Limits field to given value.

### HasLimits

`func (o *ChatV2Service) HasLimits() bool`

HasLimits returns a boolean if a field has been set.

### SetLimitsNil

`func (o *ChatV2Service) SetLimitsNil(b bool)`

 SetLimitsNil sets the value for Limits to be an explicit nil

### UnsetLimits
`func (o *ChatV2Service) UnsetLimits()`

UnsetLimits ensures that no value is present for Limits, not even an explicit nil
### GetLinks

`func (o *ChatV2Service) GetLinks() map[string]interface{}`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ChatV2Service) GetLinksOk() (*map[string]interface{}, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ChatV2Service) SetLinks(v map[string]interface{})`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ChatV2Service) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### SetLinksNil

`func (o *ChatV2Service) SetLinksNil(b bool)`

 SetLinksNil sets the value for Links to be an explicit nil

### UnsetLinks
`func (o *ChatV2Service) UnsetLinks()`

UnsetLinks ensures that no value is present for Links, not even an explicit nil
### GetMedia

`func (o *ChatV2Service) GetMedia() map[string]interface{}`

GetMedia returns the Media field if non-nil, zero value otherwise.

### GetMediaOk

`func (o *ChatV2Service) GetMediaOk() (*map[string]interface{}, bool)`

GetMediaOk returns a tuple with the Media field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMedia

`func (o *ChatV2Service) SetMedia(v map[string]interface{})`

SetMedia sets Media field to given value.

### HasMedia

`func (o *ChatV2Service) HasMedia() bool`

HasMedia returns a boolean if a field has been set.

### SetMediaNil

`func (o *ChatV2Service) SetMediaNil(b bool)`

 SetMediaNil sets the value for Media to be an explicit nil

### UnsetMedia
`func (o *ChatV2Service) UnsetMedia()`

UnsetMedia ensures that no value is present for Media, not even an explicit nil
### GetNotifications

`func (o *ChatV2Service) GetNotifications() map[string]interface{}`

GetNotifications returns the Notifications field if non-nil, zero value otherwise.

### GetNotificationsOk

`func (o *ChatV2Service) GetNotificationsOk() (*map[string]interface{}, bool)`

GetNotificationsOk returns a tuple with the Notifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifications

`func (o *ChatV2Service) SetNotifications(v map[string]interface{})`

SetNotifications sets Notifications field to given value.

### HasNotifications

`func (o *ChatV2Service) HasNotifications() bool`

HasNotifications returns a boolean if a field has been set.

### SetNotificationsNil

`func (o *ChatV2Service) SetNotificationsNil(b bool)`

 SetNotificationsNil sets the value for Notifications to be an explicit nil

### UnsetNotifications
`func (o *ChatV2Service) UnsetNotifications()`

UnsetNotifications ensures that no value is present for Notifications, not even an explicit nil
### GetPostWebhookRetryCount

`func (o *ChatV2Service) GetPostWebhookRetryCount() int32`

GetPostWebhookRetryCount returns the PostWebhookRetryCount field if non-nil, zero value otherwise.

### GetPostWebhookRetryCountOk

`func (o *ChatV2Service) GetPostWebhookRetryCountOk() (*int32, bool)`

GetPostWebhookRetryCountOk returns a tuple with the PostWebhookRetryCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostWebhookRetryCount

`func (o *ChatV2Service) SetPostWebhookRetryCount(v int32)`

SetPostWebhookRetryCount sets PostWebhookRetryCount field to given value.

### HasPostWebhookRetryCount

`func (o *ChatV2Service) HasPostWebhookRetryCount() bool`

HasPostWebhookRetryCount returns a boolean if a field has been set.

### SetPostWebhookRetryCountNil

`func (o *ChatV2Service) SetPostWebhookRetryCountNil(b bool)`

 SetPostWebhookRetryCountNil sets the value for PostWebhookRetryCount to be an explicit nil

### UnsetPostWebhookRetryCount
`func (o *ChatV2Service) UnsetPostWebhookRetryCount()`

UnsetPostWebhookRetryCount ensures that no value is present for PostWebhookRetryCount, not even an explicit nil
### GetPostWebhookUrl

`func (o *ChatV2Service) GetPostWebhookUrl() string`

GetPostWebhookUrl returns the PostWebhookUrl field if non-nil, zero value otherwise.

### GetPostWebhookUrlOk

`func (o *ChatV2Service) GetPostWebhookUrlOk() (*string, bool)`

GetPostWebhookUrlOk returns a tuple with the PostWebhookUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostWebhookUrl

`func (o *ChatV2Service) SetPostWebhookUrl(v string)`

SetPostWebhookUrl sets PostWebhookUrl field to given value.

### HasPostWebhookUrl

`func (o *ChatV2Service) HasPostWebhookUrl() bool`

HasPostWebhookUrl returns a boolean if a field has been set.

### SetPostWebhookUrlNil

`func (o *ChatV2Service) SetPostWebhookUrlNil(b bool)`

 SetPostWebhookUrlNil sets the value for PostWebhookUrl to be an explicit nil

### UnsetPostWebhookUrl
`func (o *ChatV2Service) UnsetPostWebhookUrl()`

UnsetPostWebhookUrl ensures that no value is present for PostWebhookUrl, not even an explicit nil
### GetPreWebhookRetryCount

`func (o *ChatV2Service) GetPreWebhookRetryCount() int32`

GetPreWebhookRetryCount returns the PreWebhookRetryCount field if non-nil, zero value otherwise.

### GetPreWebhookRetryCountOk

`func (o *ChatV2Service) GetPreWebhookRetryCountOk() (*int32, bool)`

GetPreWebhookRetryCountOk returns a tuple with the PreWebhookRetryCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreWebhookRetryCount

`func (o *ChatV2Service) SetPreWebhookRetryCount(v int32)`

SetPreWebhookRetryCount sets PreWebhookRetryCount field to given value.

### HasPreWebhookRetryCount

`func (o *ChatV2Service) HasPreWebhookRetryCount() bool`

HasPreWebhookRetryCount returns a boolean if a field has been set.

### SetPreWebhookRetryCountNil

`func (o *ChatV2Service) SetPreWebhookRetryCountNil(b bool)`

 SetPreWebhookRetryCountNil sets the value for PreWebhookRetryCount to be an explicit nil

### UnsetPreWebhookRetryCount
`func (o *ChatV2Service) UnsetPreWebhookRetryCount()`

UnsetPreWebhookRetryCount ensures that no value is present for PreWebhookRetryCount, not even an explicit nil
### GetPreWebhookUrl

`func (o *ChatV2Service) GetPreWebhookUrl() string`

GetPreWebhookUrl returns the PreWebhookUrl field if non-nil, zero value otherwise.

### GetPreWebhookUrlOk

`func (o *ChatV2Service) GetPreWebhookUrlOk() (*string, bool)`

GetPreWebhookUrlOk returns a tuple with the PreWebhookUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreWebhookUrl

`func (o *ChatV2Service) SetPreWebhookUrl(v string)`

SetPreWebhookUrl sets PreWebhookUrl field to given value.

### HasPreWebhookUrl

`func (o *ChatV2Service) HasPreWebhookUrl() bool`

HasPreWebhookUrl returns a boolean if a field has been set.

### SetPreWebhookUrlNil

`func (o *ChatV2Service) SetPreWebhookUrlNil(b bool)`

 SetPreWebhookUrlNil sets the value for PreWebhookUrl to be an explicit nil

### UnsetPreWebhookUrl
`func (o *ChatV2Service) UnsetPreWebhookUrl()`

UnsetPreWebhookUrl ensures that no value is present for PreWebhookUrl, not even an explicit nil
### GetReachabilityEnabled

`func (o *ChatV2Service) GetReachabilityEnabled() bool`

GetReachabilityEnabled returns the ReachabilityEnabled field if non-nil, zero value otherwise.

### GetReachabilityEnabledOk

`func (o *ChatV2Service) GetReachabilityEnabledOk() (*bool, bool)`

GetReachabilityEnabledOk returns a tuple with the ReachabilityEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReachabilityEnabled

`func (o *ChatV2Service) SetReachabilityEnabled(v bool)`

SetReachabilityEnabled sets ReachabilityEnabled field to given value.

### HasReachabilityEnabled

`func (o *ChatV2Service) HasReachabilityEnabled() bool`

HasReachabilityEnabled returns a boolean if a field has been set.

### SetReachabilityEnabledNil

`func (o *ChatV2Service) SetReachabilityEnabledNil(b bool)`

 SetReachabilityEnabledNil sets the value for ReachabilityEnabled to be an explicit nil

### UnsetReachabilityEnabled
`func (o *ChatV2Service) UnsetReachabilityEnabled()`

UnsetReachabilityEnabled ensures that no value is present for ReachabilityEnabled, not even an explicit nil
### GetReadStatusEnabled

`func (o *ChatV2Service) GetReadStatusEnabled() bool`

GetReadStatusEnabled returns the ReadStatusEnabled field if non-nil, zero value otherwise.

### GetReadStatusEnabledOk

`func (o *ChatV2Service) GetReadStatusEnabledOk() (*bool, bool)`

GetReadStatusEnabledOk returns a tuple with the ReadStatusEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadStatusEnabled

`func (o *ChatV2Service) SetReadStatusEnabled(v bool)`

SetReadStatusEnabled sets ReadStatusEnabled field to given value.

### HasReadStatusEnabled

`func (o *ChatV2Service) HasReadStatusEnabled() bool`

HasReadStatusEnabled returns a boolean if a field has been set.

### SetReadStatusEnabledNil

`func (o *ChatV2Service) SetReadStatusEnabledNil(b bool)`

 SetReadStatusEnabledNil sets the value for ReadStatusEnabled to be an explicit nil

### UnsetReadStatusEnabled
`func (o *ChatV2Service) UnsetReadStatusEnabled()`

UnsetReadStatusEnabled ensures that no value is present for ReadStatusEnabled, not even an explicit nil
### GetSid

`func (o *ChatV2Service) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *ChatV2Service) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *ChatV2Service) SetSid(v string)`

SetSid sets Sid field to given value.

### HasSid

`func (o *ChatV2Service) HasSid() bool`

HasSid returns a boolean if a field has been set.

### SetSidNil

`func (o *ChatV2Service) SetSidNil(b bool)`

 SetSidNil sets the value for Sid to be an explicit nil

### UnsetSid
`func (o *ChatV2Service) UnsetSid()`

UnsetSid ensures that no value is present for Sid, not even an explicit nil
### GetTypingIndicatorTimeout

`func (o *ChatV2Service) GetTypingIndicatorTimeout() int32`

GetTypingIndicatorTimeout returns the TypingIndicatorTimeout field if non-nil, zero value otherwise.

### GetTypingIndicatorTimeoutOk

`func (o *ChatV2Service) GetTypingIndicatorTimeoutOk() (*int32, bool)`

GetTypingIndicatorTimeoutOk returns a tuple with the TypingIndicatorTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypingIndicatorTimeout

`func (o *ChatV2Service) SetTypingIndicatorTimeout(v int32)`

SetTypingIndicatorTimeout sets TypingIndicatorTimeout field to given value.

### HasTypingIndicatorTimeout

`func (o *ChatV2Service) HasTypingIndicatorTimeout() bool`

HasTypingIndicatorTimeout returns a boolean if a field has been set.

### SetTypingIndicatorTimeoutNil

`func (o *ChatV2Service) SetTypingIndicatorTimeoutNil(b bool)`

 SetTypingIndicatorTimeoutNil sets the value for TypingIndicatorTimeout to be an explicit nil

### UnsetTypingIndicatorTimeout
`func (o *ChatV2Service) UnsetTypingIndicatorTimeout()`

UnsetTypingIndicatorTimeout ensures that no value is present for TypingIndicatorTimeout, not even an explicit nil
### GetUrl

`func (o *ChatV2Service) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ChatV2Service) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ChatV2Service) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *ChatV2Service) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *ChatV2Service) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *ChatV2Service) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetWebhookFilters

`func (o *ChatV2Service) GetWebhookFilters() []string`

GetWebhookFilters returns the WebhookFilters field if non-nil, zero value otherwise.

### GetWebhookFiltersOk

`func (o *ChatV2Service) GetWebhookFiltersOk() (*[]string, bool)`

GetWebhookFiltersOk returns a tuple with the WebhookFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookFilters

`func (o *ChatV2Service) SetWebhookFilters(v []string)`

SetWebhookFilters sets WebhookFilters field to given value.

### HasWebhookFilters

`func (o *ChatV2Service) HasWebhookFilters() bool`

HasWebhookFilters returns a boolean if a field has been set.

### SetWebhookFiltersNil

`func (o *ChatV2Service) SetWebhookFiltersNil(b bool)`

 SetWebhookFiltersNil sets the value for WebhookFilters to be an explicit nil

### UnsetWebhookFilters
`func (o *ChatV2Service) UnsetWebhookFilters()`

UnsetWebhookFilters ensures that no value is present for WebhookFilters, not even an explicit nil
### GetWebhookMethod

`func (o *ChatV2Service) GetWebhookMethod() string`

GetWebhookMethod returns the WebhookMethod field if non-nil, zero value otherwise.

### GetWebhookMethodOk

`func (o *ChatV2Service) GetWebhookMethodOk() (*string, bool)`

GetWebhookMethodOk returns a tuple with the WebhookMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookMethod

`func (o *ChatV2Service) SetWebhookMethod(v string)`

SetWebhookMethod sets WebhookMethod field to given value.

### HasWebhookMethod

`func (o *ChatV2Service) HasWebhookMethod() bool`

HasWebhookMethod returns a boolean if a field has been set.

### SetWebhookMethodNil

`func (o *ChatV2Service) SetWebhookMethodNil(b bool)`

 SetWebhookMethodNil sets the value for WebhookMethod to be an explicit nil

### UnsetWebhookMethod
`func (o *ChatV2Service) UnsetWebhookMethod()`

UnsetWebhookMethod ensures that no value is present for WebhookMethod, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


