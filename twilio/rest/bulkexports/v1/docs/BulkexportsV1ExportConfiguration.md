# BulkexportsV1ExportConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **NullableBool** | Whether files are automatically generated | [optional] 
**ResourceType** | Pointer to **NullableString** | The type of communication – Messages, Calls, Conferences, and Participants | [optional] 
**Url** | Pointer to **NullableString** | The URL of this resource. | [optional] 
**WebhookMethod** | Pointer to **NullableString** | Whether to GET or POST to the webhook url | [optional] 
**WebhookUrl** | Pointer to **NullableString** | URL targeted at export | [optional] 

## Methods

### NewBulkexportsV1ExportConfiguration

`func NewBulkexportsV1ExportConfiguration() *BulkexportsV1ExportConfiguration`

NewBulkexportsV1ExportConfiguration instantiates a new BulkexportsV1ExportConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkexportsV1ExportConfigurationWithDefaults

`func NewBulkexportsV1ExportConfigurationWithDefaults() *BulkexportsV1ExportConfiguration`

NewBulkexportsV1ExportConfigurationWithDefaults instantiates a new BulkexportsV1ExportConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *BulkexportsV1ExportConfiguration) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *BulkexportsV1ExportConfiguration) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *BulkexportsV1ExportConfiguration) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *BulkexportsV1ExportConfiguration) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### SetEnabledNil

`func (o *BulkexportsV1ExportConfiguration) SetEnabledNil(b bool)`

 SetEnabledNil sets the value for Enabled to be an explicit nil

### UnsetEnabled
`func (o *BulkexportsV1ExportConfiguration) UnsetEnabled()`

UnsetEnabled ensures that no value is present for Enabled, not even an explicit nil
### GetResourceType

`func (o *BulkexportsV1ExportConfiguration) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *BulkexportsV1ExportConfiguration) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *BulkexportsV1ExportConfiguration) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *BulkexportsV1ExportConfiguration) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.

### SetResourceTypeNil

`func (o *BulkexportsV1ExportConfiguration) SetResourceTypeNil(b bool)`

 SetResourceTypeNil sets the value for ResourceType to be an explicit nil

### UnsetResourceType
`func (o *BulkexportsV1ExportConfiguration) UnsetResourceType()`

UnsetResourceType ensures that no value is present for ResourceType, not even an explicit nil
### GetUrl

`func (o *BulkexportsV1ExportConfiguration) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *BulkexportsV1ExportConfiguration) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *BulkexportsV1ExportConfiguration) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *BulkexportsV1ExportConfiguration) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *BulkexportsV1ExportConfiguration) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *BulkexportsV1ExportConfiguration) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetWebhookMethod

`func (o *BulkexportsV1ExportConfiguration) GetWebhookMethod() string`

GetWebhookMethod returns the WebhookMethod field if non-nil, zero value otherwise.

### GetWebhookMethodOk

`func (o *BulkexportsV1ExportConfiguration) GetWebhookMethodOk() (*string, bool)`

GetWebhookMethodOk returns a tuple with the WebhookMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookMethod

`func (o *BulkexportsV1ExportConfiguration) SetWebhookMethod(v string)`

SetWebhookMethod sets WebhookMethod field to given value.

### HasWebhookMethod

`func (o *BulkexportsV1ExportConfiguration) HasWebhookMethod() bool`

HasWebhookMethod returns a boolean if a field has been set.

### SetWebhookMethodNil

`func (o *BulkexportsV1ExportConfiguration) SetWebhookMethodNil(b bool)`

 SetWebhookMethodNil sets the value for WebhookMethod to be an explicit nil

### UnsetWebhookMethod
`func (o *BulkexportsV1ExportConfiguration) UnsetWebhookMethod()`

UnsetWebhookMethod ensures that no value is present for WebhookMethod, not even an explicit nil
### GetWebhookUrl

`func (o *BulkexportsV1ExportConfiguration) GetWebhookUrl() string`

GetWebhookUrl returns the WebhookUrl field if non-nil, zero value otherwise.

### GetWebhookUrlOk

`func (o *BulkexportsV1ExportConfiguration) GetWebhookUrlOk() (*string, bool)`

GetWebhookUrlOk returns a tuple with the WebhookUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookUrl

`func (o *BulkexportsV1ExportConfiguration) SetWebhookUrl(v string)`

SetWebhookUrl sets WebhookUrl field to given value.

### HasWebhookUrl

`func (o *BulkexportsV1ExportConfiguration) HasWebhookUrl() bool`

HasWebhookUrl returns a boolean if a field has been set.

### SetWebhookUrlNil

`func (o *BulkexportsV1ExportConfiguration) SetWebhookUrlNil(b bool)`

 SetWebhookUrlNil sets the value for WebhookUrl to be an explicit nil

### UnsetWebhookUrl
`func (o *BulkexportsV1ExportConfiguration) UnsetWebhookUrl()`

UnsetWebhookUrl ensures that no value is present for WebhookUrl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


