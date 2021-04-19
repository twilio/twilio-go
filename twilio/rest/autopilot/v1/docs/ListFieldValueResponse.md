# ListFieldValueResponse

## Properties

Name | Type | Description
------------ | ------------- | -------------
**FieldValues** | Pointer to [**[]AutopilotV1AssistantFieldTypeFieldValue**](AutopilotV1AssistantFieldTypeFieldValue.md) |  | [optional] 
**Meta** | Pointer to [**ListAssistantResponseMeta**](ListAssistantResponse_meta.md) |  | [optional] 

## Methods

### NewListFieldValueResponse

`func NewListFieldValueResponse() *ListFieldValueResponse`

NewListFieldValueResponse instantiates a new ListFieldValueResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListFieldValueResponseWithDefaults

`func NewListFieldValueResponseWithDefaults() *ListFieldValueResponse`

NewListFieldValueResponseWithDefaults instantiates a new ListFieldValueResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFieldValues

`func (o *ListFieldValueResponse) GetFieldValues() []AutopilotV1AssistantFieldTypeFieldValue`

GetFieldValues returns the FieldValues field if non-nil, zero value otherwise.

### GetFieldValuesOk

`func (o *ListFieldValueResponse) GetFieldValuesOk() (*[]AutopilotV1AssistantFieldTypeFieldValue, bool)`

GetFieldValuesOk returns a tuple with the FieldValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldValues

`func (o *ListFieldValueResponse) SetFieldValues(v []AutopilotV1AssistantFieldTypeFieldValue)`

SetFieldValues sets FieldValues field to given value.

### HasFieldValues

`func (o *ListFieldValueResponse) HasFieldValues() bool`

HasFieldValues returns a boolean if a field has been set.

### GetMeta

`func (o *ListFieldValueResponse) GetMeta() ListAssistantResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListFieldValueResponse) GetMetaOk() (*ListAssistantResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListFieldValueResponse) SetMeta(v ListAssistantResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListFieldValueResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


