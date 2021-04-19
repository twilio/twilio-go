# ListAssistantResponse

## Properties

Name | Type | Description
------------ | ------------- | -------------
**Assistants** | Pointer to [**[]AutopilotV1Assistant**](AutopilotV1Assistant.md) |  | [optional] 
**Meta** | Pointer to [**ListAssistantResponseMeta**](ListAssistantResponse_meta.md) |  | [optional] 

## Methods

### NewListAssistantResponse

`func NewListAssistantResponse() *ListAssistantResponse`

NewListAssistantResponse instantiates a new ListAssistantResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListAssistantResponseWithDefaults

`func NewListAssistantResponseWithDefaults() *ListAssistantResponse`

NewListAssistantResponseWithDefaults instantiates a new ListAssistantResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssistants

`func (o *ListAssistantResponse) GetAssistants() []AutopilotV1Assistant`

GetAssistants returns the Assistants field if non-nil, zero value otherwise.

### GetAssistantsOk

`func (o *ListAssistantResponse) GetAssistantsOk() (*[]AutopilotV1Assistant, bool)`

GetAssistantsOk returns a tuple with the Assistants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssistants

`func (o *ListAssistantResponse) SetAssistants(v []AutopilotV1Assistant)`

SetAssistants sets Assistants field to given value.

### HasAssistants

`func (o *ListAssistantResponse) HasAssistants() bool`

HasAssistants returns a boolean if a field has been set.

### GetMeta

`func (o *ListAssistantResponse) GetMeta() ListAssistantResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListAssistantResponse) GetMetaOk() (*ListAssistantResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListAssistantResponse) SetMeta(v ListAssistantResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListAssistantResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


