# ListSourceIpMappingResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**ListByocTrunkResponseMeta**](ListByocTrunkResponse_meta.md) |  | [optional] 
**SourceIpMappings** | Pointer to [**[]VoiceV1SourceIpMapping**](VoiceV1SourceIpMapping.md) |  | [optional] 

## Methods

### NewListSourceIpMappingResponse

`func NewListSourceIpMappingResponse() *ListSourceIpMappingResponse`

NewListSourceIpMappingResponse instantiates a new ListSourceIpMappingResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListSourceIpMappingResponseWithDefaults

`func NewListSourceIpMappingResponseWithDefaults() *ListSourceIpMappingResponse`

NewListSourceIpMappingResponseWithDefaults instantiates a new ListSourceIpMappingResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *ListSourceIpMappingResponse) GetMeta() ListByocTrunkResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListSourceIpMappingResponse) GetMetaOk() (*ListByocTrunkResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListSourceIpMappingResponse) SetMeta(v ListByocTrunkResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListSourceIpMappingResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetSourceIpMappings

`func (o *ListSourceIpMappingResponse) GetSourceIpMappings() []VoiceV1SourceIpMapping`

GetSourceIpMappings returns the SourceIpMappings field if non-nil, zero value otherwise.

### GetSourceIpMappingsOk

`func (o *ListSourceIpMappingResponse) GetSourceIpMappingsOk() (*[]VoiceV1SourceIpMapping, bool)`

GetSourceIpMappingsOk returns a tuple with the SourceIpMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceIpMappings

`func (o *ListSourceIpMappingResponse) SetSourceIpMappings(v []VoiceV1SourceIpMapping)`

SetSourceIpMappings sets SourceIpMappings field to given value.

### HasSourceIpMappings

`func (o *ListSourceIpMappingResponse) HasSourceIpMappings() bool`

HasSourceIpMappings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


