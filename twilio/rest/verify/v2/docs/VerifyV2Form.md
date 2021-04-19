# VerifyV2Form

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FormMeta** | Pointer to **map[string]interface{}** | Additional information for the available forms for this type. | [optional] 
**FormType** | Pointer to **NullableString** | The Type of this Form | [optional] 
**Forms** | Pointer to **map[string]interface{}** | Object that contains the available forms for this type. | [optional] 
**Url** | Pointer to **NullableString** | The URL to access the forms for this type. | [optional] 

## Methods

### NewVerifyV2Form

`func NewVerifyV2Form() *VerifyV2Form`

NewVerifyV2Form instantiates a new VerifyV2Form object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVerifyV2FormWithDefaults

`func NewVerifyV2FormWithDefaults() *VerifyV2Form`

NewVerifyV2FormWithDefaults instantiates a new VerifyV2Form object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFormMeta

`func (o *VerifyV2Form) GetFormMeta() map[string]interface{}`

GetFormMeta returns the FormMeta field if non-nil, zero value otherwise.

### GetFormMetaOk

`func (o *VerifyV2Form) GetFormMetaOk() (*map[string]interface{}, bool)`

GetFormMetaOk returns a tuple with the FormMeta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormMeta

`func (o *VerifyV2Form) SetFormMeta(v map[string]interface{})`

SetFormMeta sets FormMeta field to given value.

### HasFormMeta

`func (o *VerifyV2Form) HasFormMeta() bool`

HasFormMeta returns a boolean if a field has been set.

### SetFormMetaNil

`func (o *VerifyV2Form) SetFormMetaNil(b bool)`

 SetFormMetaNil sets the value for FormMeta to be an explicit nil

### UnsetFormMeta
`func (o *VerifyV2Form) UnsetFormMeta()`

UnsetFormMeta ensures that no value is present for FormMeta, not even an explicit nil
### GetFormType

`func (o *VerifyV2Form) GetFormType() string`

GetFormType returns the FormType field if non-nil, zero value otherwise.

### GetFormTypeOk

`func (o *VerifyV2Form) GetFormTypeOk() (*string, bool)`

GetFormTypeOk returns a tuple with the FormType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormType

`func (o *VerifyV2Form) SetFormType(v string)`

SetFormType sets FormType field to given value.

### HasFormType

`func (o *VerifyV2Form) HasFormType() bool`

HasFormType returns a boolean if a field has been set.

### SetFormTypeNil

`func (o *VerifyV2Form) SetFormTypeNil(b bool)`

 SetFormTypeNil sets the value for FormType to be an explicit nil

### UnsetFormType
`func (o *VerifyV2Form) UnsetFormType()`

UnsetFormType ensures that no value is present for FormType, not even an explicit nil
### GetForms

`func (o *VerifyV2Form) GetForms() map[string]interface{}`

GetForms returns the Forms field if non-nil, zero value otherwise.

### GetFormsOk

`func (o *VerifyV2Form) GetFormsOk() (*map[string]interface{}, bool)`

GetFormsOk returns a tuple with the Forms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForms

`func (o *VerifyV2Form) SetForms(v map[string]interface{})`

SetForms sets Forms field to given value.

### HasForms

`func (o *VerifyV2Form) HasForms() bool`

HasForms returns a boolean if a field has been set.

### SetFormsNil

`func (o *VerifyV2Form) SetFormsNil(b bool)`

 SetFormsNil sets the value for Forms to be an explicit nil

### UnsetForms
`func (o *VerifyV2Form) UnsetForms()`

UnsetForms ensures that no value is present for Forms, not even an explicit nil
### GetUrl

`func (o *VerifyV2Form) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *VerifyV2Form) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *VerifyV2Form) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *VerifyV2Form) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *VerifyV2Form) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *VerifyV2Form) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


