# ListRatePlanResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**ListCommandResponseMeta**](ListCommandResponse_meta.md) |  | [optional] 
**RatePlans** | Pointer to [**[]WirelessV1RatePlan**](WirelessV1RatePlan.md) |  | [optional] 

## Methods

### NewListRatePlanResponse

`func NewListRatePlanResponse() *ListRatePlanResponse`

NewListRatePlanResponse instantiates a new ListRatePlanResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListRatePlanResponseWithDefaults

`func NewListRatePlanResponseWithDefaults() *ListRatePlanResponse`

NewListRatePlanResponseWithDefaults instantiates a new ListRatePlanResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *ListRatePlanResponse) GetMeta() ListCommandResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListRatePlanResponse) GetMetaOk() (*ListCommandResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListRatePlanResponse) SetMeta(v ListCommandResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListRatePlanResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetRatePlans

`func (o *ListRatePlanResponse) GetRatePlans() []WirelessV1RatePlan`

GetRatePlans returns the RatePlans field if non-nil, zero value otherwise.

### GetRatePlansOk

`func (o *ListRatePlanResponse) GetRatePlansOk() (*[]WirelessV1RatePlan, bool)`

GetRatePlansOk returns a tuple with the RatePlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatePlans

`func (o *ListRatePlanResponse) SetRatePlans(v []WirelessV1RatePlan)`

SetRatePlans sets RatePlans field to given value.

### HasRatePlans

`func (o *ListRatePlanResponse) HasRatePlans() bool`

HasRatePlans returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


