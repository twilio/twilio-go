# ListVersionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**ListVersionResponseMeta**](ListVersionResponse_meta.md) |  | [optional] 
**SchemaVersions** | Pointer to [**[]EventsV1SchemaVersion**](EventsV1SchemaVersion.md) |  | [optional] 

## Methods

### NewListVersionResponse

`func NewListVersionResponse() *ListVersionResponse`

NewListVersionResponse instantiates a new ListVersionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListVersionResponseWithDefaults

`func NewListVersionResponseWithDefaults() *ListVersionResponse`

NewListVersionResponseWithDefaults instantiates a new ListVersionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *ListVersionResponse) GetMeta() ListVersionResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListVersionResponse) GetMetaOk() (*ListVersionResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListVersionResponse) SetMeta(v ListVersionResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListVersionResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetSchemaVersions

`func (o *ListVersionResponse) GetSchemaVersions() []EventsV1SchemaVersion`

GetSchemaVersions returns the SchemaVersions field if non-nil, zero value otherwise.

### GetSchemaVersionsOk

`func (o *ListVersionResponse) GetSchemaVersionsOk() (*[]EventsV1SchemaVersion, bool)`

GetSchemaVersionsOk returns a tuple with the SchemaVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaVersions

`func (o *ListVersionResponse) SetSchemaVersions(v []EventsV1SchemaVersion)`

SetSchemaVersions sets SchemaVersions field to given value.

### HasSchemaVersions

`func (o *ListVersionResponse) HasSchemaVersions() bool`

HasSchemaVersions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


