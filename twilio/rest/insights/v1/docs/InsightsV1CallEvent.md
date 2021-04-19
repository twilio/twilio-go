# InsightsV1CallEvent

## Properties

Name | Type | Description
------------ | ------------- | -------------
**AccountSid** | Pointer to **NullableString** |  | [optional] 
**CallSid** | Pointer to **NullableString** |  | [optional] 
**CarrierEdge** | Pointer to **map[string]interface{}** |  | [optional] 
**ClientEdge** | Pointer to **map[string]interface{}** |  | [optional] 
**Edge** | Pointer to **NullableString** |  | [optional] 
**Group** | Pointer to **NullableString** |  | [optional] 
**Level** | Pointer to **NullableString** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**SdkEdge** | Pointer to **map[string]interface{}** |  | [optional] 
**SipEdge** | Pointer to **map[string]interface{}** |  | [optional] 
**Timestamp** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewInsightsV1CallEvent

`func NewInsightsV1CallEvent() *InsightsV1CallEvent`

NewInsightsV1CallEvent instantiates a new InsightsV1CallEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInsightsV1CallEventWithDefaults

`func NewInsightsV1CallEventWithDefaults() *InsightsV1CallEvent`

NewInsightsV1CallEventWithDefaults instantiates a new InsightsV1CallEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSid

`func (o *InsightsV1CallEvent) GetAccountSid() string`

GetAccountSid returns the AccountSid field if non-nil, zero value otherwise.

### GetAccountSidOk

`func (o *InsightsV1CallEvent) GetAccountSidOk() (*string, bool)`

GetAccountSidOk returns a tuple with the AccountSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSid

`func (o *InsightsV1CallEvent) SetAccountSid(v string)`

SetAccountSid sets AccountSid field to given value.

### HasAccountSid

`func (o *InsightsV1CallEvent) HasAccountSid() bool`

HasAccountSid returns a boolean if a field has been set.

### SetAccountSidNil

`func (o *InsightsV1CallEvent) SetAccountSidNil(b bool)`

 SetAccountSidNil sets the value for AccountSid to be an explicit nil

### UnsetAccountSid
`func (o *InsightsV1CallEvent) UnsetAccountSid()`

UnsetAccountSid ensures that no value is present for AccountSid, not even an explicit nil
### GetCallSid

`func (o *InsightsV1CallEvent) GetCallSid() string`

GetCallSid returns the CallSid field if non-nil, zero value otherwise.

### GetCallSidOk

`func (o *InsightsV1CallEvent) GetCallSidOk() (*string, bool)`

GetCallSidOk returns a tuple with the CallSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallSid

`func (o *InsightsV1CallEvent) SetCallSid(v string)`

SetCallSid sets CallSid field to given value.

### HasCallSid

`func (o *InsightsV1CallEvent) HasCallSid() bool`

HasCallSid returns a boolean if a field has been set.

### SetCallSidNil

`func (o *InsightsV1CallEvent) SetCallSidNil(b bool)`

 SetCallSidNil sets the value for CallSid to be an explicit nil

### UnsetCallSid
`func (o *InsightsV1CallEvent) UnsetCallSid()`

UnsetCallSid ensures that no value is present for CallSid, not even an explicit nil
### GetCarrierEdge

`func (o *InsightsV1CallEvent) GetCarrierEdge() map[string]interface{}`

GetCarrierEdge returns the CarrierEdge field if non-nil, zero value otherwise.

### GetCarrierEdgeOk

`func (o *InsightsV1CallEvent) GetCarrierEdgeOk() (*map[string]interface{}, bool)`

GetCarrierEdgeOk returns a tuple with the CarrierEdge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrierEdge

`func (o *InsightsV1CallEvent) SetCarrierEdge(v map[string]interface{})`

SetCarrierEdge sets CarrierEdge field to given value.

### HasCarrierEdge

`func (o *InsightsV1CallEvent) HasCarrierEdge() bool`

HasCarrierEdge returns a boolean if a field has been set.

### SetCarrierEdgeNil

`func (o *InsightsV1CallEvent) SetCarrierEdgeNil(b bool)`

 SetCarrierEdgeNil sets the value for CarrierEdge to be an explicit nil

### UnsetCarrierEdge
`func (o *InsightsV1CallEvent) UnsetCarrierEdge()`

UnsetCarrierEdge ensures that no value is present for CarrierEdge, not even an explicit nil
### GetClientEdge

`func (o *InsightsV1CallEvent) GetClientEdge() map[string]interface{}`

GetClientEdge returns the ClientEdge field if non-nil, zero value otherwise.

### GetClientEdgeOk

`func (o *InsightsV1CallEvent) GetClientEdgeOk() (*map[string]interface{}, bool)`

GetClientEdgeOk returns a tuple with the ClientEdge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientEdge

`func (o *InsightsV1CallEvent) SetClientEdge(v map[string]interface{})`

SetClientEdge sets ClientEdge field to given value.

### HasClientEdge

`func (o *InsightsV1CallEvent) HasClientEdge() bool`

HasClientEdge returns a boolean if a field has been set.

### SetClientEdgeNil

`func (o *InsightsV1CallEvent) SetClientEdgeNil(b bool)`

 SetClientEdgeNil sets the value for ClientEdge to be an explicit nil

### UnsetClientEdge
`func (o *InsightsV1CallEvent) UnsetClientEdge()`

UnsetClientEdge ensures that no value is present for ClientEdge, not even an explicit nil
### GetEdge

`func (o *InsightsV1CallEvent) GetEdge() string`

GetEdge returns the Edge field if non-nil, zero value otherwise.

### GetEdgeOk

`func (o *InsightsV1CallEvent) GetEdgeOk() (*string, bool)`

GetEdgeOk returns a tuple with the Edge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdge

`func (o *InsightsV1CallEvent) SetEdge(v string)`

SetEdge sets Edge field to given value.

### HasEdge

`func (o *InsightsV1CallEvent) HasEdge() bool`

HasEdge returns a boolean if a field has been set.

### SetEdgeNil

`func (o *InsightsV1CallEvent) SetEdgeNil(b bool)`

 SetEdgeNil sets the value for Edge to be an explicit nil

### UnsetEdge
`func (o *InsightsV1CallEvent) UnsetEdge()`

UnsetEdge ensures that no value is present for Edge, not even an explicit nil
### GetGroup

`func (o *InsightsV1CallEvent) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *InsightsV1CallEvent) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *InsightsV1CallEvent) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *InsightsV1CallEvent) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### SetGroupNil

`func (o *InsightsV1CallEvent) SetGroupNil(b bool)`

 SetGroupNil sets the value for Group to be an explicit nil

### UnsetGroup
`func (o *InsightsV1CallEvent) UnsetGroup()`

UnsetGroup ensures that no value is present for Group, not even an explicit nil
### GetLevel

`func (o *InsightsV1CallEvent) GetLevel() string`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *InsightsV1CallEvent) GetLevelOk() (*string, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *InsightsV1CallEvent) SetLevel(v string)`

SetLevel sets Level field to given value.

### HasLevel

`func (o *InsightsV1CallEvent) HasLevel() bool`

HasLevel returns a boolean if a field has been set.

### SetLevelNil

`func (o *InsightsV1CallEvent) SetLevelNil(b bool)`

 SetLevelNil sets the value for Level to be an explicit nil

### UnsetLevel
`func (o *InsightsV1CallEvent) UnsetLevel()`

UnsetLevel ensures that no value is present for Level, not even an explicit nil
### GetName

`func (o *InsightsV1CallEvent) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InsightsV1CallEvent) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InsightsV1CallEvent) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InsightsV1CallEvent) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *InsightsV1CallEvent) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *InsightsV1CallEvent) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetSdkEdge

`func (o *InsightsV1CallEvent) GetSdkEdge() map[string]interface{}`

GetSdkEdge returns the SdkEdge field if non-nil, zero value otherwise.

### GetSdkEdgeOk

`func (o *InsightsV1CallEvent) GetSdkEdgeOk() (*map[string]interface{}, bool)`

GetSdkEdgeOk returns a tuple with the SdkEdge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSdkEdge

`func (o *InsightsV1CallEvent) SetSdkEdge(v map[string]interface{})`

SetSdkEdge sets SdkEdge field to given value.

### HasSdkEdge

`func (o *InsightsV1CallEvent) HasSdkEdge() bool`

HasSdkEdge returns a boolean if a field has been set.

### SetSdkEdgeNil

`func (o *InsightsV1CallEvent) SetSdkEdgeNil(b bool)`

 SetSdkEdgeNil sets the value for SdkEdge to be an explicit nil

### UnsetSdkEdge
`func (o *InsightsV1CallEvent) UnsetSdkEdge()`

UnsetSdkEdge ensures that no value is present for SdkEdge, not even an explicit nil
### GetSipEdge

`func (o *InsightsV1CallEvent) GetSipEdge() map[string]interface{}`

GetSipEdge returns the SipEdge field if non-nil, zero value otherwise.

### GetSipEdgeOk

`func (o *InsightsV1CallEvent) GetSipEdgeOk() (*map[string]interface{}, bool)`

GetSipEdgeOk returns a tuple with the SipEdge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSipEdge

`func (o *InsightsV1CallEvent) SetSipEdge(v map[string]interface{})`

SetSipEdge sets SipEdge field to given value.

### HasSipEdge

`func (o *InsightsV1CallEvent) HasSipEdge() bool`

HasSipEdge returns a boolean if a field has been set.

### SetSipEdgeNil

`func (o *InsightsV1CallEvent) SetSipEdgeNil(b bool)`

 SetSipEdgeNil sets the value for SipEdge to be an explicit nil

### UnsetSipEdge
`func (o *InsightsV1CallEvent) UnsetSipEdge()`

UnsetSipEdge ensures that no value is present for SipEdge, not even an explicit nil
### GetTimestamp

`func (o *InsightsV1CallEvent) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *InsightsV1CallEvent) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *InsightsV1CallEvent) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *InsightsV1CallEvent) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *InsightsV1CallEvent) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *InsightsV1CallEvent) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


