# ListMetricResponse

## Properties

Name | Type | Description
------------ | ------------- | -------------
**Meta** | Pointer to [**ListVideoRoomSummaryResponseMeta**](ListVideoRoomSummaryResponse_meta.md) |  | [optional] 
**Metrics** | Pointer to [**[]InsightsV1CallMetric**](InsightsV1CallMetric.md) |  | [optional] 

## Methods

### NewListMetricResponse

`func NewListMetricResponse() *ListMetricResponse`

NewListMetricResponse instantiates a new ListMetricResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListMetricResponseWithDefaults

`func NewListMetricResponseWithDefaults() *ListMetricResponse`

NewListMetricResponseWithDefaults instantiates a new ListMetricResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *ListMetricResponse) GetMeta() ListVideoRoomSummaryResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListMetricResponse) GetMetaOk() (*ListVideoRoomSummaryResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListMetricResponse) SetMeta(v ListVideoRoomSummaryResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListMetricResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetMetrics

`func (o *ListMetricResponse) GetMetrics() []InsightsV1CallMetric`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *ListMetricResponse) GetMetricsOk() (*[]InsightsV1CallMetric, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *ListMetricResponse) SetMetrics(v []InsightsV1CallMetric)`

SetMetrics sets Metrics field to given value.

### HasMetrics

`func (o *ListMetricResponse) HasMetrics() bool`

HasMetrics returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


