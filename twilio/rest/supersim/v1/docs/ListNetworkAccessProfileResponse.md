# ListNetworkAccessProfileResponse

## Properties

Name | Type | Description
------------ | ------------- | -------------
**Meta** | Pointer to [**ListCommandResponseMeta**](ListCommandResponse_meta.md) |  | [optional] 
**NetworkAccessProfiles** | Pointer to [**[]SupersimV1NetworkAccessProfile**](SupersimV1NetworkAccessProfile.md) |  | [optional] 

## Methods

### NewListNetworkAccessProfileResponse

`func NewListNetworkAccessProfileResponse() *ListNetworkAccessProfileResponse`

NewListNetworkAccessProfileResponse instantiates a new ListNetworkAccessProfileResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListNetworkAccessProfileResponseWithDefaults

`func NewListNetworkAccessProfileResponseWithDefaults() *ListNetworkAccessProfileResponse`

NewListNetworkAccessProfileResponseWithDefaults instantiates a new ListNetworkAccessProfileResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *ListNetworkAccessProfileResponse) GetMeta() ListCommandResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListNetworkAccessProfileResponse) GetMetaOk() (*ListCommandResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListNetworkAccessProfileResponse) SetMeta(v ListCommandResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListNetworkAccessProfileResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetNetworkAccessProfiles

`func (o *ListNetworkAccessProfileResponse) GetNetworkAccessProfiles() []SupersimV1NetworkAccessProfile`

GetNetworkAccessProfiles returns the NetworkAccessProfiles field if non-nil, zero value otherwise.

### GetNetworkAccessProfilesOk

`func (o *ListNetworkAccessProfileResponse) GetNetworkAccessProfilesOk() (*[]SupersimV1NetworkAccessProfile, bool)`

GetNetworkAccessProfilesOk returns a tuple with the NetworkAccessProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkAccessProfiles

`func (o *ListNetworkAccessProfileResponse) SetNetworkAccessProfiles(v []SupersimV1NetworkAccessProfile)`

SetNetworkAccessProfiles sets NetworkAccessProfiles field to given value.

### HasNetworkAccessProfiles

`func (o *ListNetworkAccessProfileResponse) HasNetworkAccessProfiles() bool`

HasNetworkAccessProfiles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


