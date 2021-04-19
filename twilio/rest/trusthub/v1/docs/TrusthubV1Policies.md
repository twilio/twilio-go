# TrusthubV1Policies

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FriendlyName** | Pointer to **NullableString** | A human-readable description of the Policy resource | [optional] 
**Requirements** | Pointer to **map[string]interface{}** | The sid of a Policy object that dictates requirements | [optional] 
**Sid** | Pointer to **NullableString** | The unique string that identifies the Policy resource | [optional] 
**Url** | Pointer to **NullableString** | The absolute URL of the Policy resource | [optional] 

## Methods

### NewTrusthubV1Policies

`func NewTrusthubV1Policies() *TrusthubV1Policies`

NewTrusthubV1Policies instantiates a new TrusthubV1Policies object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrusthubV1PoliciesWithDefaults

`func NewTrusthubV1PoliciesWithDefaults() *TrusthubV1Policies`

NewTrusthubV1PoliciesWithDefaults instantiates a new TrusthubV1Policies object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFriendlyName

`func (o *TrusthubV1Policies) GetFriendlyName() string`

GetFriendlyName returns the FriendlyName field if non-nil, zero value otherwise.

### GetFriendlyNameOk

`func (o *TrusthubV1Policies) GetFriendlyNameOk() (*string, bool)`

GetFriendlyNameOk returns a tuple with the FriendlyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFriendlyName

`func (o *TrusthubV1Policies) SetFriendlyName(v string)`

SetFriendlyName sets FriendlyName field to given value.

### HasFriendlyName

`func (o *TrusthubV1Policies) HasFriendlyName() bool`

HasFriendlyName returns a boolean if a field has been set.

### SetFriendlyNameNil

`func (o *TrusthubV1Policies) SetFriendlyNameNil(b bool)`

 SetFriendlyNameNil sets the value for FriendlyName to be an explicit nil

### UnsetFriendlyName
`func (o *TrusthubV1Policies) UnsetFriendlyName()`

UnsetFriendlyName ensures that no value is present for FriendlyName, not even an explicit nil
### GetRequirements

`func (o *TrusthubV1Policies) GetRequirements() map[string]interface{}`

GetRequirements returns the Requirements field if non-nil, zero value otherwise.

### GetRequirementsOk

`func (o *TrusthubV1Policies) GetRequirementsOk() (*map[string]interface{}, bool)`

GetRequirementsOk returns a tuple with the Requirements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequirements

`func (o *TrusthubV1Policies) SetRequirements(v map[string]interface{})`

SetRequirements sets Requirements field to given value.

### HasRequirements

`func (o *TrusthubV1Policies) HasRequirements() bool`

HasRequirements returns a boolean if a field has been set.

### SetRequirementsNil

`func (o *TrusthubV1Policies) SetRequirementsNil(b bool)`

 SetRequirementsNil sets the value for Requirements to be an explicit nil

### UnsetRequirements
`func (o *TrusthubV1Policies) UnsetRequirements()`

UnsetRequirements ensures that no value is present for Requirements, not even an explicit nil
### GetSid

`func (o *TrusthubV1Policies) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *TrusthubV1Policies) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *TrusthubV1Policies) SetSid(v string)`

SetSid sets Sid field to given value.

### HasSid

`func (o *TrusthubV1Policies) HasSid() bool`

HasSid returns a boolean if a field has been set.

### SetSidNil

`func (o *TrusthubV1Policies) SetSidNil(b bool)`

 SetSidNil sets the value for Sid to be an explicit nil

### UnsetSid
`func (o *TrusthubV1Policies) UnsetSid()`

UnsetSid ensures that no value is present for Sid, not even an explicit nil
### GetUrl

`func (o *TrusthubV1Policies) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *TrusthubV1Policies) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *TrusthubV1Policies) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *TrusthubV1Policies) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *TrusthubV1Policies) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *TrusthubV1Policies) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


