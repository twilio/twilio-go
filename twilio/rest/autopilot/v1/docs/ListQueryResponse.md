# ListQueryResponse

## Properties

Name | Type | Description
------------ | ------------- | -------------
**Meta** | Pointer to [**ListAssistantResponseMeta**](ListAssistantResponse_meta.md) |  | [optional] 
**Queries** | Pointer to [**[]AutopilotV1AssistantQuery**](AutopilotV1AssistantQuery.md) |  | [optional] 

## Methods

### NewListQueryResponse

`func NewListQueryResponse() *ListQueryResponse`

NewListQueryResponse instantiates a new ListQueryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListQueryResponseWithDefaults

`func NewListQueryResponseWithDefaults() *ListQueryResponse`

NewListQueryResponseWithDefaults instantiates a new ListQueryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *ListQueryResponse) GetMeta() ListAssistantResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListQueryResponse) GetMetaOk() (*ListAssistantResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListQueryResponse) SetMeta(v ListAssistantResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListQueryResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetQueries

`func (o *ListQueryResponse) GetQueries() []AutopilotV1AssistantQuery`

GetQueries returns the Queries field if non-nil, zero value otherwise.

### GetQueriesOk

`func (o *ListQueryResponse) GetQueriesOk() (*[]AutopilotV1AssistantQuery, bool)`

GetQueriesOk returns a tuple with the Queries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueries

`func (o *ListQueryResponse) SetQueries(v []AutopilotV1AssistantQuery)`

SetQueries sets Queries field to given value.

### HasQueries

`func (o *ListQueryResponse) HasQueries() bool`

HasQueries returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


