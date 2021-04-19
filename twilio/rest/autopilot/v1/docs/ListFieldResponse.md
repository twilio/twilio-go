# ListFieldResponse

## Properties

Name | Type | Description
------------ | ------------- | -------------
**Fields** | Pointer to [**[]AutopilotV1AssistantTaskField**](AutopilotV1AssistantTaskField.md) |  | [optional] 
**Meta** | Pointer to [**ListAssistantResponseMeta**](ListAssistantResponse_meta.md) |  | [optional] 

## Methods

### NewListFieldResponse

`func NewListFieldResponse() *ListFieldResponse`

NewListFieldResponse instantiates a new ListFieldResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListFieldResponseWithDefaults

`func NewListFieldResponseWithDefaults() *ListFieldResponse`

NewListFieldResponseWithDefaults instantiates a new ListFieldResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFields

`func (o *ListFieldResponse) GetFields() []AutopilotV1AssistantTaskField`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *ListFieldResponse) GetFieldsOk() (*[]AutopilotV1AssistantTaskField, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *ListFieldResponse) SetFields(v []AutopilotV1AssistantTaskField)`

SetFields sets Fields field to given value.

### HasFields

`func (o *ListFieldResponse) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetMeta

`func (o *ListFieldResponse) GetMeta() ListAssistantResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListFieldResponse) GetMetaOk() (*ListAssistantResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListFieldResponse) SetMeta(v ListAssistantResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListFieldResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


