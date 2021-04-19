# ListMessagingConfigurationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MessagingConfigurations** | Pointer to [**[]VerifyV2ServiceMessagingConfiguration**](VerifyV2ServiceMessagingConfiguration.md) |  | [optional] 
**Meta** | Pointer to [**ListVerificationAttemptResponseMeta**](ListVerificationAttemptResponse_meta.md) |  | [optional] 

## Methods

### NewListMessagingConfigurationResponse

`func NewListMessagingConfigurationResponse() *ListMessagingConfigurationResponse`

NewListMessagingConfigurationResponse instantiates a new ListMessagingConfigurationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListMessagingConfigurationResponseWithDefaults

`func NewListMessagingConfigurationResponseWithDefaults() *ListMessagingConfigurationResponse`

NewListMessagingConfigurationResponseWithDefaults instantiates a new ListMessagingConfigurationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessagingConfigurations

`func (o *ListMessagingConfigurationResponse) GetMessagingConfigurations() []VerifyV2ServiceMessagingConfiguration`

GetMessagingConfigurations returns the MessagingConfigurations field if non-nil, zero value otherwise.

### GetMessagingConfigurationsOk

`func (o *ListMessagingConfigurationResponse) GetMessagingConfigurationsOk() (*[]VerifyV2ServiceMessagingConfiguration, bool)`

GetMessagingConfigurationsOk returns a tuple with the MessagingConfigurations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessagingConfigurations

`func (o *ListMessagingConfigurationResponse) SetMessagingConfigurations(v []VerifyV2ServiceMessagingConfiguration)`

SetMessagingConfigurations sets MessagingConfigurations field to given value.

### HasMessagingConfigurations

`func (o *ListMessagingConfigurationResponse) HasMessagingConfigurations() bool`

HasMessagingConfigurations returns a boolean if a field has been set.

### GetMeta

`func (o *ListMessagingConfigurationResponse) GetMeta() ListVerificationAttemptResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListMessagingConfigurationResponse) GetMetaOk() (*ListVerificationAttemptResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListMessagingConfigurationResponse) SetMeta(v ListVerificationAttemptResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListMessagingConfigurationResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


