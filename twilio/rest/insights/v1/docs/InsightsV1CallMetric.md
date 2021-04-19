# InsightsV1CallMetric

## Properties

Name | Type | Description
------------ | ------------- | -------------
**AccountSid** | Pointer to **NullableString** |  | [optional] 
**CallSid** | Pointer to **NullableString** |  | [optional] 
**CarrierEdge** | Pointer to **map[string]interface{}** |  | [optional] 
**ClientEdge** | Pointer to **map[string]interface{}** |  | [optional] 
**Direction** | Pointer to **NullableString** |  | [optional] 
**Edge** | Pointer to **NullableString** |  | [optional] 
**SdkEdge** | Pointer to **map[string]interface{}** |  | [optional] 
**SipEdge** | Pointer to **map[string]interface{}** |  | [optional] 
**Timestamp** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewInsightsV1CallMetric

`func NewInsightsV1CallMetric() *InsightsV1CallMetric`

NewInsightsV1CallMetric instantiates a new InsightsV1CallMetric object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInsightsV1CallMetricWithDefaults

`func NewInsightsV1CallMetricWithDefaults() *InsightsV1CallMetric`

NewInsightsV1CallMetricWithDefaults instantiates a new InsightsV1CallMetric object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSid

`func (o *InsightsV1CallMetric) GetAccountSid() string`

GetAccountSid returns the AccountSid field if non-nil, zero value otherwise.

### GetAccountSidOk

`func (o *InsightsV1CallMetric) GetAccountSidOk() (*string, bool)`

GetAccountSidOk returns a tuple with the AccountSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSid

`func (o *InsightsV1CallMetric) SetAccountSid(v string)`

SetAccountSid sets AccountSid field to given value.

### HasAccountSid

`func (o *InsightsV1CallMetric) HasAccountSid() bool`

HasAccountSid returns a boolean if a field has been set.

### SetAccountSidNil

`func (o *InsightsV1CallMetric) SetAccountSidNil(b bool)`

 SetAccountSidNil sets the value for AccountSid to be an explicit nil

### UnsetAccountSid
`func (o *InsightsV1CallMetric) UnsetAccountSid()`

UnsetAccountSid ensures that no value is present for AccountSid, not even an explicit nil
### GetCallSid

`func (o *InsightsV1CallMetric) GetCallSid() string`

GetCallSid returns the CallSid field if non-nil, zero value otherwise.

### GetCallSidOk

`func (o *InsightsV1CallMetric) GetCallSidOk() (*string, bool)`

GetCallSidOk returns a tuple with the CallSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallSid

`func (o *InsightsV1CallMetric) SetCallSid(v string)`

SetCallSid sets CallSid field to given value.

### HasCallSid

`func (o *InsightsV1CallMetric) HasCallSid() bool`

HasCallSid returns a boolean if a field has been set.

### SetCallSidNil

`func (o *InsightsV1CallMetric) SetCallSidNil(b bool)`

 SetCallSidNil sets the value for CallSid to be an explicit nil

### UnsetCallSid
`func (o *InsightsV1CallMetric) UnsetCallSid()`

UnsetCallSid ensures that no value is present for CallSid, not even an explicit nil
### GetCarrierEdge

`func (o *InsightsV1CallMetric) GetCarrierEdge() map[string]interface{}`

GetCarrierEdge returns the CarrierEdge field if non-nil, zero value otherwise.

### GetCarrierEdgeOk

`func (o *InsightsV1CallMetric) GetCarrierEdgeOk() (*map[string]interface{}, bool)`

GetCarrierEdgeOk returns a tuple with the CarrierEdge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrierEdge

`func (o *InsightsV1CallMetric) SetCarrierEdge(v map[string]interface{})`

SetCarrierEdge sets CarrierEdge field to given value.

### HasCarrierEdge

`func (o *InsightsV1CallMetric) HasCarrierEdge() bool`

HasCarrierEdge returns a boolean if a field has been set.

### SetCarrierEdgeNil

`func (o *InsightsV1CallMetric) SetCarrierEdgeNil(b bool)`

 SetCarrierEdgeNil sets the value for CarrierEdge to be an explicit nil

### UnsetCarrierEdge
`func (o *InsightsV1CallMetric) UnsetCarrierEdge()`

UnsetCarrierEdge ensures that no value is present for CarrierEdge, not even an explicit nil
### GetClientEdge

`func (o *InsightsV1CallMetric) GetClientEdge() map[string]interface{}`

GetClientEdge returns the ClientEdge field if non-nil, zero value otherwise.

### GetClientEdgeOk

`func (o *InsightsV1CallMetric) GetClientEdgeOk() (*map[string]interface{}, bool)`

GetClientEdgeOk returns a tuple with the ClientEdge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientEdge

`func (o *InsightsV1CallMetric) SetClientEdge(v map[string]interface{})`

SetClientEdge sets ClientEdge field to given value.

### HasClientEdge

`func (o *InsightsV1CallMetric) HasClientEdge() bool`

HasClientEdge returns a boolean if a field has been set.

### SetClientEdgeNil

`func (o *InsightsV1CallMetric) SetClientEdgeNil(b bool)`

 SetClientEdgeNil sets the value for ClientEdge to be an explicit nil

### UnsetClientEdge
`func (o *InsightsV1CallMetric) UnsetClientEdge()`

UnsetClientEdge ensures that no value is present for ClientEdge, not even an explicit nil
### GetDirection

`func (o *InsightsV1CallMetric) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *InsightsV1CallMetric) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *InsightsV1CallMetric) SetDirection(v string)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *InsightsV1CallMetric) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### SetDirectionNil

`func (o *InsightsV1CallMetric) SetDirectionNil(b bool)`

 SetDirectionNil sets the value for Direction to be an explicit nil

### UnsetDirection
`func (o *InsightsV1CallMetric) UnsetDirection()`

UnsetDirection ensures that no value is present for Direction, not even an explicit nil
### GetEdge

`func (o *InsightsV1CallMetric) GetEdge() string`

GetEdge returns the Edge field if non-nil, zero value otherwise.

### GetEdgeOk

`func (o *InsightsV1CallMetric) GetEdgeOk() (*string, bool)`

GetEdgeOk returns a tuple with the Edge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdge

`func (o *InsightsV1CallMetric) SetEdge(v string)`

SetEdge sets Edge field to given value.

### HasEdge

`func (o *InsightsV1CallMetric) HasEdge() bool`

HasEdge returns a boolean if a field has been set.

### SetEdgeNil

`func (o *InsightsV1CallMetric) SetEdgeNil(b bool)`

 SetEdgeNil sets the value for Edge to be an explicit nil

### UnsetEdge
`func (o *InsightsV1CallMetric) UnsetEdge()`

UnsetEdge ensures that no value is present for Edge, not even an explicit nil
### GetSdkEdge

`func (o *InsightsV1CallMetric) GetSdkEdge() map[string]interface{}`

GetSdkEdge returns the SdkEdge field if non-nil, zero value otherwise.

### GetSdkEdgeOk

`func (o *InsightsV1CallMetric) GetSdkEdgeOk() (*map[string]interface{}, bool)`

GetSdkEdgeOk returns a tuple with the SdkEdge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSdkEdge

`func (o *InsightsV1CallMetric) SetSdkEdge(v map[string]interface{})`

SetSdkEdge sets SdkEdge field to given value.

### HasSdkEdge

`func (o *InsightsV1CallMetric) HasSdkEdge() bool`

HasSdkEdge returns a boolean if a field has been set.

### SetSdkEdgeNil

`func (o *InsightsV1CallMetric) SetSdkEdgeNil(b bool)`

 SetSdkEdgeNil sets the value for SdkEdge to be an explicit nil

### UnsetSdkEdge
`func (o *InsightsV1CallMetric) UnsetSdkEdge()`

UnsetSdkEdge ensures that no value is present for SdkEdge, not even an explicit nil
### GetSipEdge

`func (o *InsightsV1CallMetric) GetSipEdge() map[string]interface{}`

GetSipEdge returns the SipEdge field if non-nil, zero value otherwise.

### GetSipEdgeOk

`func (o *InsightsV1CallMetric) GetSipEdgeOk() (*map[string]interface{}, bool)`

GetSipEdgeOk returns a tuple with the SipEdge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSipEdge

`func (o *InsightsV1CallMetric) SetSipEdge(v map[string]interface{})`

SetSipEdge sets SipEdge field to given value.

### HasSipEdge

`func (o *InsightsV1CallMetric) HasSipEdge() bool`

HasSipEdge returns a boolean if a field has been set.

### SetSipEdgeNil

`func (o *InsightsV1CallMetric) SetSipEdgeNil(b bool)`

 SetSipEdgeNil sets the value for SipEdge to be an explicit nil

### UnsetSipEdge
`func (o *InsightsV1CallMetric) UnsetSipEdge()`

UnsetSipEdge ensures that no value is present for SipEdge, not even an explicit nil
### GetTimestamp

`func (o *InsightsV1CallMetric) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *InsightsV1CallMetric) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *InsightsV1CallMetric) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *InsightsV1CallMetric) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *InsightsV1CallMetric) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *InsightsV1CallMetric) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


