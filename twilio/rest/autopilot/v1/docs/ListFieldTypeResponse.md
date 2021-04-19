# ListFieldTypeResponse

## Properties

Name | Type | Description
------------ | ------------- | -------------
**FieldTypes** | Pointer to [**[]AutopilotV1AssistantFieldType**](AutopilotV1AssistantFieldType.md) |  | [optional] 
**Meta** | Pointer to [**ListAssistantResponseMeta**](ListAssistantResponse_meta.md) |  | [optional] 

## Methods

### NewListFieldTypeResponse

`func NewListFieldTypeResponse() *ListFieldTypeResponse`

NewListFieldTypeResponse instantiates a new ListFieldTypeResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListFieldTypeResponseWithDefaults

`func NewListFieldTypeResponseWithDefaults() *ListFieldTypeResponse`

NewListFieldTypeResponseWithDefaults instantiates a new ListFieldTypeResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFieldTypes

`func (o *ListFieldTypeResponse) GetFieldTypes() []AutopilotV1AssistantFieldType`

GetFieldTypes returns the FieldTypes field if non-nil, zero value otherwise.

### GetFieldTypesOk

`func (o *ListFieldTypeResponse) GetFieldTypesOk() (*[]AutopilotV1AssistantFieldType, bool)`

GetFieldTypesOk returns a tuple with the FieldTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldTypes

`func (o *ListFieldTypeResponse) SetFieldTypes(v []AutopilotV1AssistantFieldType)`

SetFieldTypes sets FieldTypes field to given value.

### HasFieldTypes

`func (o *ListFieldTypeResponse) HasFieldTypes() bool`

HasFieldTypes returns a boolean if a field has been set.

### GetMeta

`func (o *ListFieldTypeResponse) GetMeta() ListAssistantResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListFieldTypeResponse) GetMetaOk() (*ListAssistantResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListFieldTypeResponse) SetMeta(v ListAssistantResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListFieldTypeResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


