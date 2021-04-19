# ApiV2010AccountUsageUsageTrigger

## Properties

Name | Type | Description
------------ | ------------- | -------------
**AccountSid** | Pointer to **NullableString** | The SID of the Account that this trigger monitors | [optional] 
**ApiVersion** | Pointer to **NullableString** | The API version used to create the resource | [optional] 
**CallbackMethod** | Pointer to **NullableString** | The HTTP method we use to call callback_url | [optional] 
**CallbackUrl** | Pointer to **NullableString** | he URL we call when the trigger fires | [optional] 
**CurrentValue** | Pointer to **NullableString** | The current value of the field the trigger is watching | [optional] 
**DateCreated** | Pointer to **NullableString** | The RFC 2822 date and time in GMT that the resource was created | [optional] 
**DateFired** | Pointer to **NullableString** | The RFC 2822 date and time in GMT that the trigger was last fired | [optional] 
**DateUpdated** | Pointer to **NullableString** | The RFC 2822 date and time in GMT that the resource was last updated | [optional] 
**FriendlyName** | Pointer to **NullableString** | The string that you assigned to describe the trigger | [optional] 
**Recurring** | Pointer to **NullableString** | The frequency of a recurring UsageTrigger | [optional] 
**Sid** | Pointer to **NullableString** | The unique string that identifies the resource | [optional] 
**TriggerBy** | Pointer to **NullableString** | The field in the UsageRecord resource that fires the trigger | [optional] 
**TriggerValue** | Pointer to **NullableString** | The value at which the trigger will fire | [optional] 
**Uri** | Pointer to **NullableString** | The URI of the resource, relative to &#x60;https://api.twilio.com&#x60; | [optional] 
**UsageCategory** | Pointer to **NullableString** | The usage category the trigger watches | [optional] 
**UsageRecordUri** | Pointer to **NullableString** | The URI of the UsageRecord resource this trigger watches | [optional] 

## Methods

### NewApiV2010AccountUsageUsageTrigger

`func NewApiV2010AccountUsageUsageTrigger() *ApiV2010AccountUsageUsageTrigger`

NewApiV2010AccountUsageUsageTrigger instantiates a new ApiV2010AccountUsageUsageTrigger object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiV2010AccountUsageUsageTriggerWithDefaults

`func NewApiV2010AccountUsageUsageTriggerWithDefaults() *ApiV2010AccountUsageUsageTrigger`

NewApiV2010AccountUsageUsageTriggerWithDefaults instantiates a new ApiV2010AccountUsageUsageTrigger object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSid

`func (o *ApiV2010AccountUsageUsageTrigger) GetAccountSid() string`

GetAccountSid returns the AccountSid field if non-nil, zero value otherwise.

### GetAccountSidOk

`func (o *ApiV2010AccountUsageUsageTrigger) GetAccountSidOk() (*string, bool)`

GetAccountSidOk returns a tuple with the AccountSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSid

`func (o *ApiV2010AccountUsageUsageTrigger) SetAccountSid(v string)`

SetAccountSid sets AccountSid field to given value.

### HasAccountSid

`func (o *ApiV2010AccountUsageUsageTrigger) HasAccountSid() bool`

HasAccountSid returns a boolean if a field has been set.

### SetAccountSidNil

`func (o *ApiV2010AccountUsageUsageTrigger) SetAccountSidNil(b bool)`

 SetAccountSidNil sets the value for AccountSid to be an explicit nil

### UnsetAccountSid
`func (o *ApiV2010AccountUsageUsageTrigger) UnsetAccountSid()`

UnsetAccountSid ensures that no value is present for AccountSid, not even an explicit nil
### GetApiVersion

`func (o *ApiV2010AccountUsageUsageTrigger) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *ApiV2010AccountUsageUsageTrigger) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *ApiV2010AccountUsageUsageTrigger) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *ApiV2010AccountUsageUsageTrigger) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### SetApiVersionNil

`func (o *ApiV2010AccountUsageUsageTrigger) SetApiVersionNil(b bool)`

 SetApiVersionNil sets the value for ApiVersion to be an explicit nil

### UnsetApiVersion
`func (o *ApiV2010AccountUsageUsageTrigger) UnsetApiVersion()`

UnsetApiVersion ensures that no value is present for ApiVersion, not even an explicit nil
### GetCallbackMethod

`func (o *ApiV2010AccountUsageUsageTrigger) GetCallbackMethod() string`

GetCallbackMethod returns the CallbackMethod field if non-nil, zero value otherwise.

### GetCallbackMethodOk

`func (o *ApiV2010AccountUsageUsageTrigger) GetCallbackMethodOk() (*string, bool)`

GetCallbackMethodOk returns a tuple with the CallbackMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackMethod

`func (o *ApiV2010AccountUsageUsageTrigger) SetCallbackMethod(v string)`

SetCallbackMethod sets CallbackMethod field to given value.

### HasCallbackMethod

`func (o *ApiV2010AccountUsageUsageTrigger) HasCallbackMethod() bool`

HasCallbackMethod returns a boolean if a field has been set.

### SetCallbackMethodNil

`func (o *ApiV2010AccountUsageUsageTrigger) SetCallbackMethodNil(b bool)`

 SetCallbackMethodNil sets the value for CallbackMethod to be an explicit nil

### UnsetCallbackMethod
`func (o *ApiV2010AccountUsageUsageTrigger) UnsetCallbackMethod()`

UnsetCallbackMethod ensures that no value is present for CallbackMethod, not even an explicit nil
### GetCallbackUrl

`func (o *ApiV2010AccountUsageUsageTrigger) GetCallbackUrl() string`

GetCallbackUrl returns the CallbackUrl field if non-nil, zero value otherwise.

### GetCallbackUrlOk

`func (o *ApiV2010AccountUsageUsageTrigger) GetCallbackUrlOk() (*string, bool)`

GetCallbackUrlOk returns a tuple with the CallbackUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackUrl

`func (o *ApiV2010AccountUsageUsageTrigger) SetCallbackUrl(v string)`

SetCallbackUrl sets CallbackUrl field to given value.

### HasCallbackUrl

`func (o *ApiV2010AccountUsageUsageTrigger) HasCallbackUrl() bool`

HasCallbackUrl returns a boolean if a field has been set.

### SetCallbackUrlNil

`func (o *ApiV2010AccountUsageUsageTrigger) SetCallbackUrlNil(b bool)`

 SetCallbackUrlNil sets the value for CallbackUrl to be an explicit nil

### UnsetCallbackUrl
`func (o *ApiV2010AccountUsageUsageTrigger) UnsetCallbackUrl()`

UnsetCallbackUrl ensures that no value is present for CallbackUrl, not even an explicit nil
### GetCurrentValue

`func (o *ApiV2010AccountUsageUsageTrigger) GetCurrentValue() string`

GetCurrentValue returns the CurrentValue field if non-nil, zero value otherwise.

### GetCurrentValueOk

`func (o *ApiV2010AccountUsageUsageTrigger) GetCurrentValueOk() (*string, bool)`

GetCurrentValueOk returns a tuple with the CurrentValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentValue

`func (o *ApiV2010AccountUsageUsageTrigger) SetCurrentValue(v string)`

SetCurrentValue sets CurrentValue field to given value.

### HasCurrentValue

`func (o *ApiV2010AccountUsageUsageTrigger) HasCurrentValue() bool`

HasCurrentValue returns a boolean if a field has been set.

### SetCurrentValueNil

`func (o *ApiV2010AccountUsageUsageTrigger) SetCurrentValueNil(b bool)`

 SetCurrentValueNil sets the value for CurrentValue to be an explicit nil

### UnsetCurrentValue
`func (o *ApiV2010AccountUsageUsageTrigger) UnsetCurrentValue()`

UnsetCurrentValue ensures that no value is present for CurrentValue, not even an explicit nil
### GetDateCreated

`func (o *ApiV2010AccountUsageUsageTrigger) GetDateCreated() string`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *ApiV2010AccountUsageUsageTrigger) GetDateCreatedOk() (*string, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *ApiV2010AccountUsageUsageTrigger) SetDateCreated(v string)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *ApiV2010AccountUsageUsageTrigger) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### SetDateCreatedNil

`func (o *ApiV2010AccountUsageUsageTrigger) SetDateCreatedNil(b bool)`

 SetDateCreatedNil sets the value for DateCreated to be an explicit nil

### UnsetDateCreated
`func (o *ApiV2010AccountUsageUsageTrigger) UnsetDateCreated()`

UnsetDateCreated ensures that no value is present for DateCreated, not even an explicit nil
### GetDateFired

`func (o *ApiV2010AccountUsageUsageTrigger) GetDateFired() string`

GetDateFired returns the DateFired field if non-nil, zero value otherwise.

### GetDateFiredOk

`func (o *ApiV2010AccountUsageUsageTrigger) GetDateFiredOk() (*string, bool)`

GetDateFiredOk returns a tuple with the DateFired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateFired

`func (o *ApiV2010AccountUsageUsageTrigger) SetDateFired(v string)`

SetDateFired sets DateFired field to given value.

### HasDateFired

`func (o *ApiV2010AccountUsageUsageTrigger) HasDateFired() bool`

HasDateFired returns a boolean if a field has been set.

### SetDateFiredNil

`func (o *ApiV2010AccountUsageUsageTrigger) SetDateFiredNil(b bool)`

 SetDateFiredNil sets the value for DateFired to be an explicit nil

### UnsetDateFired
`func (o *ApiV2010AccountUsageUsageTrigger) UnsetDateFired()`

UnsetDateFired ensures that no value is present for DateFired, not even an explicit nil
### GetDateUpdated

`func (o *ApiV2010AccountUsageUsageTrigger) GetDateUpdated() string`

GetDateUpdated returns the DateUpdated field if non-nil, zero value otherwise.

### GetDateUpdatedOk

`func (o *ApiV2010AccountUsageUsageTrigger) GetDateUpdatedOk() (*string, bool)`

GetDateUpdatedOk returns a tuple with the DateUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateUpdated

`func (o *ApiV2010AccountUsageUsageTrigger) SetDateUpdated(v string)`

SetDateUpdated sets DateUpdated field to given value.

### HasDateUpdated

`func (o *ApiV2010AccountUsageUsageTrigger) HasDateUpdated() bool`

HasDateUpdated returns a boolean if a field has been set.

### SetDateUpdatedNil

`func (o *ApiV2010AccountUsageUsageTrigger) SetDateUpdatedNil(b bool)`

 SetDateUpdatedNil sets the value for DateUpdated to be an explicit nil

### UnsetDateUpdated
`func (o *ApiV2010AccountUsageUsageTrigger) UnsetDateUpdated()`

UnsetDateUpdated ensures that no value is present for DateUpdated, not even an explicit nil
### GetFriendlyName

`func (o *ApiV2010AccountUsageUsageTrigger) GetFriendlyName() string`

GetFriendlyName returns the FriendlyName field if non-nil, zero value otherwise.

### GetFriendlyNameOk

`func (o *ApiV2010AccountUsageUsageTrigger) GetFriendlyNameOk() (*string, bool)`

GetFriendlyNameOk returns a tuple with the FriendlyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFriendlyName

`func (o *ApiV2010AccountUsageUsageTrigger) SetFriendlyName(v string)`

SetFriendlyName sets FriendlyName field to given value.

### HasFriendlyName

`func (o *ApiV2010AccountUsageUsageTrigger) HasFriendlyName() bool`

HasFriendlyName returns a boolean if a field has been set.

### SetFriendlyNameNil

`func (o *ApiV2010AccountUsageUsageTrigger) SetFriendlyNameNil(b bool)`

 SetFriendlyNameNil sets the value for FriendlyName to be an explicit nil

### UnsetFriendlyName
`func (o *ApiV2010AccountUsageUsageTrigger) UnsetFriendlyName()`

UnsetFriendlyName ensures that no value is present for FriendlyName, not even an explicit nil
### GetRecurring

`func (o *ApiV2010AccountUsageUsageTrigger) GetRecurring() string`

GetRecurring returns the Recurring field if non-nil, zero value otherwise.

### GetRecurringOk

`func (o *ApiV2010AccountUsageUsageTrigger) GetRecurringOk() (*string, bool)`

GetRecurringOk returns a tuple with the Recurring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurring

`func (o *ApiV2010AccountUsageUsageTrigger) SetRecurring(v string)`

SetRecurring sets Recurring field to given value.

### HasRecurring

`func (o *ApiV2010AccountUsageUsageTrigger) HasRecurring() bool`

HasRecurring returns a boolean if a field has been set.

### SetRecurringNil

`func (o *ApiV2010AccountUsageUsageTrigger) SetRecurringNil(b bool)`

 SetRecurringNil sets the value for Recurring to be an explicit nil

### UnsetRecurring
`func (o *ApiV2010AccountUsageUsageTrigger) UnsetRecurring()`

UnsetRecurring ensures that no value is present for Recurring, not even an explicit nil
### GetSid

`func (o *ApiV2010AccountUsageUsageTrigger) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *ApiV2010AccountUsageUsageTrigger) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *ApiV2010AccountUsageUsageTrigger) SetSid(v string)`

SetSid sets Sid field to given value.

### HasSid

`func (o *ApiV2010AccountUsageUsageTrigger) HasSid() bool`

HasSid returns a boolean if a field has been set.

### SetSidNil

`func (o *ApiV2010AccountUsageUsageTrigger) SetSidNil(b bool)`

 SetSidNil sets the value for Sid to be an explicit nil

### UnsetSid
`func (o *ApiV2010AccountUsageUsageTrigger) UnsetSid()`

UnsetSid ensures that no value is present for Sid, not even an explicit nil
### GetTriggerBy

`func (o *ApiV2010AccountUsageUsageTrigger) GetTriggerBy() string`

GetTriggerBy returns the TriggerBy field if non-nil, zero value otherwise.

### GetTriggerByOk

`func (o *ApiV2010AccountUsageUsageTrigger) GetTriggerByOk() (*string, bool)`

GetTriggerByOk returns a tuple with the TriggerBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerBy

`func (o *ApiV2010AccountUsageUsageTrigger) SetTriggerBy(v string)`

SetTriggerBy sets TriggerBy field to given value.

### HasTriggerBy

`func (o *ApiV2010AccountUsageUsageTrigger) HasTriggerBy() bool`

HasTriggerBy returns a boolean if a field has been set.

### SetTriggerByNil

`func (o *ApiV2010AccountUsageUsageTrigger) SetTriggerByNil(b bool)`

 SetTriggerByNil sets the value for TriggerBy to be an explicit nil

### UnsetTriggerBy
`func (o *ApiV2010AccountUsageUsageTrigger) UnsetTriggerBy()`

UnsetTriggerBy ensures that no value is present for TriggerBy, not even an explicit nil
### GetTriggerValue

`func (o *ApiV2010AccountUsageUsageTrigger) GetTriggerValue() string`

GetTriggerValue returns the TriggerValue field if non-nil, zero value otherwise.

### GetTriggerValueOk

`func (o *ApiV2010AccountUsageUsageTrigger) GetTriggerValueOk() (*string, bool)`

GetTriggerValueOk returns a tuple with the TriggerValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerValue

`func (o *ApiV2010AccountUsageUsageTrigger) SetTriggerValue(v string)`

SetTriggerValue sets TriggerValue field to given value.

### HasTriggerValue

`func (o *ApiV2010AccountUsageUsageTrigger) HasTriggerValue() bool`

HasTriggerValue returns a boolean if a field has been set.

### SetTriggerValueNil

`func (o *ApiV2010AccountUsageUsageTrigger) SetTriggerValueNil(b bool)`

 SetTriggerValueNil sets the value for TriggerValue to be an explicit nil

### UnsetTriggerValue
`func (o *ApiV2010AccountUsageUsageTrigger) UnsetTriggerValue()`

UnsetTriggerValue ensures that no value is present for TriggerValue, not even an explicit nil
### GetUri

`func (o *ApiV2010AccountUsageUsageTrigger) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *ApiV2010AccountUsageUsageTrigger) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *ApiV2010AccountUsageUsageTrigger) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *ApiV2010AccountUsageUsageTrigger) HasUri() bool`

HasUri returns a boolean if a field has been set.

### SetUriNil

`func (o *ApiV2010AccountUsageUsageTrigger) SetUriNil(b bool)`

 SetUriNil sets the value for Uri to be an explicit nil

### UnsetUri
`func (o *ApiV2010AccountUsageUsageTrigger) UnsetUri()`

UnsetUri ensures that no value is present for Uri, not even an explicit nil
### GetUsageCategory

`func (o *ApiV2010AccountUsageUsageTrigger) GetUsageCategory() string`

GetUsageCategory returns the UsageCategory field if non-nil, zero value otherwise.

### GetUsageCategoryOk

`func (o *ApiV2010AccountUsageUsageTrigger) GetUsageCategoryOk() (*string, bool)`

GetUsageCategoryOk returns a tuple with the UsageCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageCategory

`func (o *ApiV2010AccountUsageUsageTrigger) SetUsageCategory(v string)`

SetUsageCategory sets UsageCategory field to given value.

### HasUsageCategory

`func (o *ApiV2010AccountUsageUsageTrigger) HasUsageCategory() bool`

HasUsageCategory returns a boolean if a field has been set.

### SetUsageCategoryNil

`func (o *ApiV2010AccountUsageUsageTrigger) SetUsageCategoryNil(b bool)`

 SetUsageCategoryNil sets the value for UsageCategory to be an explicit nil

### UnsetUsageCategory
`func (o *ApiV2010AccountUsageUsageTrigger) UnsetUsageCategory()`

UnsetUsageCategory ensures that no value is present for UsageCategory, not even an explicit nil
### GetUsageRecordUri

`func (o *ApiV2010AccountUsageUsageTrigger) GetUsageRecordUri() string`

GetUsageRecordUri returns the UsageRecordUri field if non-nil, zero value otherwise.

### GetUsageRecordUriOk

`func (o *ApiV2010AccountUsageUsageTrigger) GetUsageRecordUriOk() (*string, bool)`

GetUsageRecordUriOk returns a tuple with the UsageRecordUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageRecordUri

`func (o *ApiV2010AccountUsageUsageTrigger) SetUsageRecordUri(v string)`

SetUsageRecordUri sets UsageRecordUri field to given value.

### HasUsageRecordUri

`func (o *ApiV2010AccountUsageUsageTrigger) HasUsageRecordUri() bool`

HasUsageRecordUri returns a boolean if a field has been set.

### SetUsageRecordUriNil

`func (o *ApiV2010AccountUsageUsageTrigger) SetUsageRecordUriNil(b bool)`

 SetUsageRecordUriNil sets the value for UsageRecordUri to be an explicit nil

### UnsetUsageRecordUri
`func (o *ApiV2010AccountUsageUsageTrigger) UnsetUsageRecordUri()`

UnsetUsageRecordUri ensures that no value is present for UsageRecordUri, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


