# ConversationsV1ConfigurationConfigurationWebhook

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **NullableString** | The unique ID of the Account responsible for this conversation. | [optional] 
**Filters** | Pointer to **[]string** | The list of webhook event triggers that are enabled for this Service. | [optional] 
**Method** | Pointer to **NullableString** | The HTTP method to be used when sending a webhook request. | [optional] 
**PostWebhookUrl** | Pointer to **NullableString** | The absolute url the post-event webhook request should be sent to. | [optional] 
**PreWebhookUrl** | Pointer to **NullableString** | The absolute url the pre-event webhook request should be sent to. | [optional] 
**Target** | Pointer to **NullableString** | The routing target of the webhook. | [optional] 
**Url** | Pointer to **NullableString** | An absolute URL for this webhook. | [optional] 

## Methods

### NewConversationsV1ConfigurationConfigurationWebhook

`func NewConversationsV1ConfigurationConfigurationWebhook() *ConversationsV1ConfigurationConfigurationWebhook`

NewConversationsV1ConfigurationConfigurationWebhook instantiates a new ConversationsV1ConfigurationConfigurationWebhook object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConversationsV1ConfigurationConfigurationWebhookWithDefaults

`func NewConversationsV1ConfigurationConfigurationWebhookWithDefaults() *ConversationsV1ConfigurationConfigurationWebhook`

NewConversationsV1ConfigurationConfigurationWebhookWithDefaults instantiates a new ConversationsV1ConfigurationConfigurationWebhook object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSid

`func (o *ConversationsV1ConfigurationConfigurationWebhook) GetAccountSid() string`

GetAccountSid returns the AccountSid field if non-nil, zero value otherwise.

### GetAccountSidOk

`func (o *ConversationsV1ConfigurationConfigurationWebhook) GetAccountSidOk() (*string, bool)`

GetAccountSidOk returns a tuple with the AccountSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSid

`func (o *ConversationsV1ConfigurationConfigurationWebhook) SetAccountSid(v string)`

SetAccountSid sets AccountSid field to given value.

### HasAccountSid

`func (o *ConversationsV1ConfigurationConfigurationWebhook) HasAccountSid() bool`

HasAccountSid returns a boolean if a field has been set.

### SetAccountSidNil

`func (o *ConversationsV1ConfigurationConfigurationWebhook) SetAccountSidNil(b bool)`

 SetAccountSidNil sets the value for AccountSid to be an explicit nil

### UnsetAccountSid
`func (o *ConversationsV1ConfigurationConfigurationWebhook) UnsetAccountSid()`

UnsetAccountSid ensures that no value is present for AccountSid, not even an explicit nil
### GetFilters

`func (o *ConversationsV1ConfigurationConfigurationWebhook) GetFilters() []string`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *ConversationsV1ConfigurationConfigurationWebhook) GetFiltersOk() (*[]string, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *ConversationsV1ConfigurationConfigurationWebhook) SetFilters(v []string)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *ConversationsV1ConfigurationConfigurationWebhook) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### SetFiltersNil

`func (o *ConversationsV1ConfigurationConfigurationWebhook) SetFiltersNil(b bool)`

 SetFiltersNil sets the value for Filters to be an explicit nil

### UnsetFilters
`func (o *ConversationsV1ConfigurationConfigurationWebhook) UnsetFilters()`

UnsetFilters ensures that no value is present for Filters, not even an explicit nil
### GetMethod

`func (o *ConversationsV1ConfigurationConfigurationWebhook) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *ConversationsV1ConfigurationConfigurationWebhook) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *ConversationsV1ConfigurationConfigurationWebhook) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *ConversationsV1ConfigurationConfigurationWebhook) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### SetMethodNil

`func (o *ConversationsV1ConfigurationConfigurationWebhook) SetMethodNil(b bool)`

 SetMethodNil sets the value for Method to be an explicit nil

### UnsetMethod
`func (o *ConversationsV1ConfigurationConfigurationWebhook) UnsetMethod()`

UnsetMethod ensures that no value is present for Method, not even an explicit nil
### GetPostWebhookUrl

`func (o *ConversationsV1ConfigurationConfigurationWebhook) GetPostWebhookUrl() string`

GetPostWebhookUrl returns the PostWebhookUrl field if non-nil, zero value otherwise.

### GetPostWebhookUrlOk

`func (o *ConversationsV1ConfigurationConfigurationWebhook) GetPostWebhookUrlOk() (*string, bool)`

GetPostWebhookUrlOk returns a tuple with the PostWebhookUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostWebhookUrl

`func (o *ConversationsV1ConfigurationConfigurationWebhook) SetPostWebhookUrl(v string)`

SetPostWebhookUrl sets PostWebhookUrl field to given value.

### HasPostWebhookUrl

`func (o *ConversationsV1ConfigurationConfigurationWebhook) HasPostWebhookUrl() bool`

HasPostWebhookUrl returns a boolean if a field has been set.

### SetPostWebhookUrlNil

`func (o *ConversationsV1ConfigurationConfigurationWebhook) SetPostWebhookUrlNil(b bool)`

 SetPostWebhookUrlNil sets the value for PostWebhookUrl to be an explicit nil

### UnsetPostWebhookUrl
`func (o *ConversationsV1ConfigurationConfigurationWebhook) UnsetPostWebhookUrl()`

UnsetPostWebhookUrl ensures that no value is present for PostWebhookUrl, not even an explicit nil
### GetPreWebhookUrl

`func (o *ConversationsV1ConfigurationConfigurationWebhook) GetPreWebhookUrl() string`

GetPreWebhookUrl returns the PreWebhookUrl field if non-nil, zero value otherwise.

### GetPreWebhookUrlOk

`func (o *ConversationsV1ConfigurationConfigurationWebhook) GetPreWebhookUrlOk() (*string, bool)`

GetPreWebhookUrlOk returns a tuple with the PreWebhookUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreWebhookUrl

`func (o *ConversationsV1ConfigurationConfigurationWebhook) SetPreWebhookUrl(v string)`

SetPreWebhookUrl sets PreWebhookUrl field to given value.

### HasPreWebhookUrl

`func (o *ConversationsV1ConfigurationConfigurationWebhook) HasPreWebhookUrl() bool`

HasPreWebhookUrl returns a boolean if a field has been set.

### SetPreWebhookUrlNil

`func (o *ConversationsV1ConfigurationConfigurationWebhook) SetPreWebhookUrlNil(b bool)`

 SetPreWebhookUrlNil sets the value for PreWebhookUrl to be an explicit nil

### UnsetPreWebhookUrl
`func (o *ConversationsV1ConfigurationConfigurationWebhook) UnsetPreWebhookUrl()`

UnsetPreWebhookUrl ensures that no value is present for PreWebhookUrl, not even an explicit nil
### GetTarget

`func (o *ConversationsV1ConfigurationConfigurationWebhook) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *ConversationsV1ConfigurationConfigurationWebhook) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *ConversationsV1ConfigurationConfigurationWebhook) SetTarget(v string)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *ConversationsV1ConfigurationConfigurationWebhook) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### SetTargetNil

`func (o *ConversationsV1ConfigurationConfigurationWebhook) SetTargetNil(b bool)`

 SetTargetNil sets the value for Target to be an explicit nil

### UnsetTarget
`func (o *ConversationsV1ConfigurationConfigurationWebhook) UnsetTarget()`

UnsetTarget ensures that no value is present for Target, not even an explicit nil
### GetUrl

`func (o *ConversationsV1ConfigurationConfigurationWebhook) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ConversationsV1ConfigurationConfigurationWebhook) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ConversationsV1ConfigurationConfigurationWebhook) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *ConversationsV1ConfigurationConfigurationWebhook) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *ConversationsV1ConfigurationConfigurationWebhook) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *ConversationsV1ConfigurationConfigurationWebhook) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


