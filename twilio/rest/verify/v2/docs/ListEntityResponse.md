# ListEntityResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entities** | Pointer to [**[]VerifyV2ServiceEntity**](VerifyV2ServiceEntity.md) |  | [optional] 
**Meta** | Pointer to [**ListVerificationAttemptResponseMeta**](ListVerificationAttemptResponse_meta.md) |  | [optional] 

## Methods

### NewListEntityResponse

`func NewListEntityResponse() *ListEntityResponse`

NewListEntityResponse instantiates a new ListEntityResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListEntityResponseWithDefaults

`func NewListEntityResponseWithDefaults() *ListEntityResponse`

NewListEntityResponseWithDefaults instantiates a new ListEntityResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntities

`func (o *ListEntityResponse) GetEntities() []VerifyV2ServiceEntity`

GetEntities returns the Entities field if non-nil, zero value otherwise.

### GetEntitiesOk

`func (o *ListEntityResponse) GetEntitiesOk() (*[]VerifyV2ServiceEntity, bool)`

GetEntitiesOk returns a tuple with the Entities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntities

`func (o *ListEntityResponse) SetEntities(v []VerifyV2ServiceEntity)`

SetEntities sets Entities field to given value.

### HasEntities

`func (o *ListEntityResponse) HasEntities() bool`

HasEntities returns a boolean if a field has been set.

### GetMeta

`func (o *ListEntityResponse) GetMeta() ListVerificationAttemptResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListEntityResponse) GetMetaOk() (*ListVerificationAttemptResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListEntityResponse) SetMeta(v ListVerificationAttemptResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListEntityResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


