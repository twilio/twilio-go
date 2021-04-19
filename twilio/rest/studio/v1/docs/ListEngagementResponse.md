# ListEngagementResponse

## Properties

Name | Type | Description
------------ | ------------- | -------------
**Engagements** | Pointer to [**[]StudioV1FlowEngagement**](StudioV1FlowEngagement.md) |  | [optional] 
**Meta** | Pointer to [**ListFlowResponseMeta**](ListFlowResponse_meta.md) |  | [optional] 

## Methods

### NewListEngagementResponse

`func NewListEngagementResponse() *ListEngagementResponse`

NewListEngagementResponse instantiates a new ListEngagementResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListEngagementResponseWithDefaults

`func NewListEngagementResponseWithDefaults() *ListEngagementResponse`

NewListEngagementResponseWithDefaults instantiates a new ListEngagementResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEngagements

`func (o *ListEngagementResponse) GetEngagements() []StudioV1FlowEngagement`

GetEngagements returns the Engagements field if non-nil, zero value otherwise.

### GetEngagementsOk

`func (o *ListEngagementResponse) GetEngagementsOk() (*[]StudioV1FlowEngagement, bool)`

GetEngagementsOk returns a tuple with the Engagements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngagements

`func (o *ListEngagementResponse) SetEngagements(v []StudioV1FlowEngagement)`

SetEngagements sets Engagements field to given value.

### HasEngagements

`func (o *ListEngagementResponse) HasEngagements() bool`

HasEngagements returns a boolean if a field has been set.

### GetMeta

`func (o *ListEngagementResponse) GetMeta() ListFlowResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListEngagementResponse) GetMetaOk() (*ListFlowResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListEngagementResponse) SetMeta(v ListFlowResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListEngagementResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


