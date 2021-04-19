# InsightsV1CallSummary

## Properties

Name | Type | Description
------------ | ------------- | -------------
**AccountSid** | Pointer to **NullableString** |  | [optional] 
**Attributes** | Pointer to **map[string]interface{}** |  | [optional] 
**CallSid** | Pointer to **NullableString** |  | [optional] 
**CallState** | Pointer to **NullableString** |  | [optional] 
**CallType** | Pointer to **NullableString** |  | [optional] 
**CarrierEdge** | Pointer to **map[string]interface{}** |  | [optional] 
**ClientEdge** | Pointer to **map[string]interface{}** |  | [optional] 
**ConnectDuration** | Pointer to **NullableInt32** |  | [optional] 
**CreatedTime** | Pointer to **NullableTime** |  | [optional] 
**Duration** | Pointer to **NullableInt32** |  | [optional] 
**EndTime** | Pointer to **NullableTime** |  | [optional] 
**From** | Pointer to **map[string]interface{}** |  | [optional] 
**ProcessingState** | Pointer to **NullableString** |  | [optional] 
**Properties** | Pointer to **map[string]interface{}** |  | [optional] 
**SdkEdge** | Pointer to **map[string]interface{}** |  | [optional] 
**SipEdge** | Pointer to **map[string]interface{}** |  | [optional] 
**StartTime** | Pointer to **NullableTime** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**To** | Pointer to **map[string]interface{}** |  | [optional] 
**Trust** | Pointer to **map[string]interface{}** |  | [optional] 
**Url** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewInsightsV1CallSummary

`func NewInsightsV1CallSummary() *InsightsV1CallSummary`

NewInsightsV1CallSummary instantiates a new InsightsV1CallSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInsightsV1CallSummaryWithDefaults

`func NewInsightsV1CallSummaryWithDefaults() *InsightsV1CallSummary`

NewInsightsV1CallSummaryWithDefaults instantiates a new InsightsV1CallSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSid

`func (o *InsightsV1CallSummary) GetAccountSid() string`

GetAccountSid returns the AccountSid field if non-nil, zero value otherwise.

### GetAccountSidOk

`func (o *InsightsV1CallSummary) GetAccountSidOk() (*string, bool)`

GetAccountSidOk returns a tuple with the AccountSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSid

`func (o *InsightsV1CallSummary) SetAccountSid(v string)`

SetAccountSid sets AccountSid field to given value.

### HasAccountSid

`func (o *InsightsV1CallSummary) HasAccountSid() bool`

HasAccountSid returns a boolean if a field has been set.

### SetAccountSidNil

`func (o *InsightsV1CallSummary) SetAccountSidNil(b bool)`

 SetAccountSidNil sets the value for AccountSid to be an explicit nil

### UnsetAccountSid
`func (o *InsightsV1CallSummary) UnsetAccountSid()`

UnsetAccountSid ensures that no value is present for AccountSid, not even an explicit nil
### GetAttributes

`func (o *InsightsV1CallSummary) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *InsightsV1CallSummary) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *InsightsV1CallSummary) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *InsightsV1CallSummary) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### SetAttributesNil

`func (o *InsightsV1CallSummary) SetAttributesNil(b bool)`

 SetAttributesNil sets the value for Attributes to be an explicit nil

### UnsetAttributes
`func (o *InsightsV1CallSummary) UnsetAttributes()`

UnsetAttributes ensures that no value is present for Attributes, not even an explicit nil
### GetCallSid

`func (o *InsightsV1CallSummary) GetCallSid() string`

GetCallSid returns the CallSid field if non-nil, zero value otherwise.

### GetCallSidOk

`func (o *InsightsV1CallSummary) GetCallSidOk() (*string, bool)`

GetCallSidOk returns a tuple with the CallSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallSid

`func (o *InsightsV1CallSummary) SetCallSid(v string)`

SetCallSid sets CallSid field to given value.

### HasCallSid

`func (o *InsightsV1CallSummary) HasCallSid() bool`

HasCallSid returns a boolean if a field has been set.

### SetCallSidNil

`func (o *InsightsV1CallSummary) SetCallSidNil(b bool)`

 SetCallSidNil sets the value for CallSid to be an explicit nil

### UnsetCallSid
`func (o *InsightsV1CallSummary) UnsetCallSid()`

UnsetCallSid ensures that no value is present for CallSid, not even an explicit nil
### GetCallState

`func (o *InsightsV1CallSummary) GetCallState() string`

GetCallState returns the CallState field if non-nil, zero value otherwise.

### GetCallStateOk

`func (o *InsightsV1CallSummary) GetCallStateOk() (*string, bool)`

GetCallStateOk returns a tuple with the CallState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallState

`func (o *InsightsV1CallSummary) SetCallState(v string)`

SetCallState sets CallState field to given value.

### HasCallState

`func (o *InsightsV1CallSummary) HasCallState() bool`

HasCallState returns a boolean if a field has been set.

### SetCallStateNil

`func (o *InsightsV1CallSummary) SetCallStateNil(b bool)`

 SetCallStateNil sets the value for CallState to be an explicit nil

### UnsetCallState
`func (o *InsightsV1CallSummary) UnsetCallState()`

UnsetCallState ensures that no value is present for CallState, not even an explicit nil
### GetCallType

`func (o *InsightsV1CallSummary) GetCallType() string`

GetCallType returns the CallType field if non-nil, zero value otherwise.

### GetCallTypeOk

`func (o *InsightsV1CallSummary) GetCallTypeOk() (*string, bool)`

GetCallTypeOk returns a tuple with the CallType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallType

`func (o *InsightsV1CallSummary) SetCallType(v string)`

SetCallType sets CallType field to given value.

### HasCallType

`func (o *InsightsV1CallSummary) HasCallType() bool`

HasCallType returns a boolean if a field has been set.

### SetCallTypeNil

`func (o *InsightsV1CallSummary) SetCallTypeNil(b bool)`

 SetCallTypeNil sets the value for CallType to be an explicit nil

### UnsetCallType
`func (o *InsightsV1CallSummary) UnsetCallType()`

UnsetCallType ensures that no value is present for CallType, not even an explicit nil
### GetCarrierEdge

`func (o *InsightsV1CallSummary) GetCarrierEdge() map[string]interface{}`

GetCarrierEdge returns the CarrierEdge field if non-nil, zero value otherwise.

### GetCarrierEdgeOk

`func (o *InsightsV1CallSummary) GetCarrierEdgeOk() (*map[string]interface{}, bool)`

GetCarrierEdgeOk returns a tuple with the CarrierEdge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrierEdge

`func (o *InsightsV1CallSummary) SetCarrierEdge(v map[string]interface{})`

SetCarrierEdge sets CarrierEdge field to given value.

### HasCarrierEdge

`func (o *InsightsV1CallSummary) HasCarrierEdge() bool`

HasCarrierEdge returns a boolean if a field has been set.

### SetCarrierEdgeNil

`func (o *InsightsV1CallSummary) SetCarrierEdgeNil(b bool)`

 SetCarrierEdgeNil sets the value for CarrierEdge to be an explicit nil

### UnsetCarrierEdge
`func (o *InsightsV1CallSummary) UnsetCarrierEdge()`

UnsetCarrierEdge ensures that no value is present for CarrierEdge, not even an explicit nil
### GetClientEdge

`func (o *InsightsV1CallSummary) GetClientEdge() map[string]interface{}`

GetClientEdge returns the ClientEdge field if non-nil, zero value otherwise.

### GetClientEdgeOk

`func (o *InsightsV1CallSummary) GetClientEdgeOk() (*map[string]interface{}, bool)`

GetClientEdgeOk returns a tuple with the ClientEdge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientEdge

`func (o *InsightsV1CallSummary) SetClientEdge(v map[string]interface{})`

SetClientEdge sets ClientEdge field to given value.

### HasClientEdge

`func (o *InsightsV1CallSummary) HasClientEdge() bool`

HasClientEdge returns a boolean if a field has been set.

### SetClientEdgeNil

`func (o *InsightsV1CallSummary) SetClientEdgeNil(b bool)`

 SetClientEdgeNil sets the value for ClientEdge to be an explicit nil

### UnsetClientEdge
`func (o *InsightsV1CallSummary) UnsetClientEdge()`

UnsetClientEdge ensures that no value is present for ClientEdge, not even an explicit nil
### GetConnectDuration

`func (o *InsightsV1CallSummary) GetConnectDuration() int32`

GetConnectDuration returns the ConnectDuration field if non-nil, zero value otherwise.

### GetConnectDurationOk

`func (o *InsightsV1CallSummary) GetConnectDurationOk() (*int32, bool)`

GetConnectDurationOk returns a tuple with the ConnectDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectDuration

`func (o *InsightsV1CallSummary) SetConnectDuration(v int32)`

SetConnectDuration sets ConnectDuration field to given value.

### HasConnectDuration

`func (o *InsightsV1CallSummary) HasConnectDuration() bool`

HasConnectDuration returns a boolean if a field has been set.

### SetConnectDurationNil

`func (o *InsightsV1CallSummary) SetConnectDurationNil(b bool)`

 SetConnectDurationNil sets the value for ConnectDuration to be an explicit nil

### UnsetConnectDuration
`func (o *InsightsV1CallSummary) UnsetConnectDuration()`

UnsetConnectDuration ensures that no value is present for ConnectDuration, not even an explicit nil
### GetCreatedTime

`func (o *InsightsV1CallSummary) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *InsightsV1CallSummary) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *InsightsV1CallSummary) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *InsightsV1CallSummary) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### SetCreatedTimeNil

`func (o *InsightsV1CallSummary) SetCreatedTimeNil(b bool)`

 SetCreatedTimeNil sets the value for CreatedTime to be an explicit nil

### UnsetCreatedTime
`func (o *InsightsV1CallSummary) UnsetCreatedTime()`

UnsetCreatedTime ensures that no value is present for CreatedTime, not even an explicit nil
### GetDuration

`func (o *InsightsV1CallSummary) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *InsightsV1CallSummary) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *InsightsV1CallSummary) SetDuration(v int32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *InsightsV1CallSummary) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### SetDurationNil

`func (o *InsightsV1CallSummary) SetDurationNil(b bool)`

 SetDurationNil sets the value for Duration to be an explicit nil

### UnsetDuration
`func (o *InsightsV1CallSummary) UnsetDuration()`

UnsetDuration ensures that no value is present for Duration, not even an explicit nil
### GetEndTime

`func (o *InsightsV1CallSummary) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *InsightsV1CallSummary) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *InsightsV1CallSummary) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *InsightsV1CallSummary) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### SetEndTimeNil

`func (o *InsightsV1CallSummary) SetEndTimeNil(b bool)`

 SetEndTimeNil sets the value for EndTime to be an explicit nil

### UnsetEndTime
`func (o *InsightsV1CallSummary) UnsetEndTime()`

UnsetEndTime ensures that no value is present for EndTime, not even an explicit nil
### GetFrom

`func (o *InsightsV1CallSummary) GetFrom() map[string]interface{}`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *InsightsV1CallSummary) GetFromOk() (*map[string]interface{}, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *InsightsV1CallSummary) SetFrom(v map[string]interface{})`

SetFrom sets From field to given value.

### HasFrom

`func (o *InsightsV1CallSummary) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### SetFromNil

`func (o *InsightsV1CallSummary) SetFromNil(b bool)`

 SetFromNil sets the value for From to be an explicit nil

### UnsetFrom
`func (o *InsightsV1CallSummary) UnsetFrom()`

UnsetFrom ensures that no value is present for From, not even an explicit nil
### GetProcessingState

`func (o *InsightsV1CallSummary) GetProcessingState() string`

GetProcessingState returns the ProcessingState field if non-nil, zero value otherwise.

### GetProcessingStateOk

`func (o *InsightsV1CallSummary) GetProcessingStateOk() (*string, bool)`

GetProcessingStateOk returns a tuple with the ProcessingState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessingState

`func (o *InsightsV1CallSummary) SetProcessingState(v string)`

SetProcessingState sets ProcessingState field to given value.

### HasProcessingState

`func (o *InsightsV1CallSummary) HasProcessingState() bool`

HasProcessingState returns a boolean if a field has been set.

### SetProcessingStateNil

`func (o *InsightsV1CallSummary) SetProcessingStateNil(b bool)`

 SetProcessingStateNil sets the value for ProcessingState to be an explicit nil

### UnsetProcessingState
`func (o *InsightsV1CallSummary) UnsetProcessingState()`

UnsetProcessingState ensures that no value is present for ProcessingState, not even an explicit nil
### GetProperties

`func (o *InsightsV1CallSummary) GetProperties() map[string]interface{}`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *InsightsV1CallSummary) GetPropertiesOk() (*map[string]interface{}, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *InsightsV1CallSummary) SetProperties(v map[string]interface{})`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *InsightsV1CallSummary) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### SetPropertiesNil

`func (o *InsightsV1CallSummary) SetPropertiesNil(b bool)`

 SetPropertiesNil sets the value for Properties to be an explicit nil

### UnsetProperties
`func (o *InsightsV1CallSummary) UnsetProperties()`

UnsetProperties ensures that no value is present for Properties, not even an explicit nil
### GetSdkEdge

`func (o *InsightsV1CallSummary) GetSdkEdge() map[string]interface{}`

GetSdkEdge returns the SdkEdge field if non-nil, zero value otherwise.

### GetSdkEdgeOk

`func (o *InsightsV1CallSummary) GetSdkEdgeOk() (*map[string]interface{}, bool)`

GetSdkEdgeOk returns a tuple with the SdkEdge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSdkEdge

`func (o *InsightsV1CallSummary) SetSdkEdge(v map[string]interface{})`

SetSdkEdge sets SdkEdge field to given value.

### HasSdkEdge

`func (o *InsightsV1CallSummary) HasSdkEdge() bool`

HasSdkEdge returns a boolean if a field has been set.

### SetSdkEdgeNil

`func (o *InsightsV1CallSummary) SetSdkEdgeNil(b bool)`

 SetSdkEdgeNil sets the value for SdkEdge to be an explicit nil

### UnsetSdkEdge
`func (o *InsightsV1CallSummary) UnsetSdkEdge()`

UnsetSdkEdge ensures that no value is present for SdkEdge, not even an explicit nil
### GetSipEdge

`func (o *InsightsV1CallSummary) GetSipEdge() map[string]interface{}`

GetSipEdge returns the SipEdge field if non-nil, zero value otherwise.

### GetSipEdgeOk

`func (o *InsightsV1CallSummary) GetSipEdgeOk() (*map[string]interface{}, bool)`

GetSipEdgeOk returns a tuple with the SipEdge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSipEdge

`func (o *InsightsV1CallSummary) SetSipEdge(v map[string]interface{})`

SetSipEdge sets SipEdge field to given value.

### HasSipEdge

`func (o *InsightsV1CallSummary) HasSipEdge() bool`

HasSipEdge returns a boolean if a field has been set.

### SetSipEdgeNil

`func (o *InsightsV1CallSummary) SetSipEdgeNil(b bool)`

 SetSipEdgeNil sets the value for SipEdge to be an explicit nil

### UnsetSipEdge
`func (o *InsightsV1CallSummary) UnsetSipEdge()`

UnsetSipEdge ensures that no value is present for SipEdge, not even an explicit nil
### GetStartTime

`func (o *InsightsV1CallSummary) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *InsightsV1CallSummary) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *InsightsV1CallSummary) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *InsightsV1CallSummary) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### SetStartTimeNil

`func (o *InsightsV1CallSummary) SetStartTimeNil(b bool)`

 SetStartTimeNil sets the value for StartTime to be an explicit nil

### UnsetStartTime
`func (o *InsightsV1CallSummary) UnsetStartTime()`

UnsetStartTime ensures that no value is present for StartTime, not even an explicit nil
### GetTags

`func (o *InsightsV1CallSummary) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *InsightsV1CallSummary) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *InsightsV1CallSummary) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *InsightsV1CallSummary) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *InsightsV1CallSummary) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *InsightsV1CallSummary) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetTo

`func (o *InsightsV1CallSummary) GetTo() map[string]interface{}`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *InsightsV1CallSummary) GetToOk() (*map[string]interface{}, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *InsightsV1CallSummary) SetTo(v map[string]interface{})`

SetTo sets To field to given value.

### HasTo

`func (o *InsightsV1CallSummary) HasTo() bool`

HasTo returns a boolean if a field has been set.

### SetToNil

`func (o *InsightsV1CallSummary) SetToNil(b bool)`

 SetToNil sets the value for To to be an explicit nil

### UnsetTo
`func (o *InsightsV1CallSummary) UnsetTo()`

UnsetTo ensures that no value is present for To, not even an explicit nil
### GetTrust

`func (o *InsightsV1CallSummary) GetTrust() map[string]interface{}`

GetTrust returns the Trust field if non-nil, zero value otherwise.

### GetTrustOk

`func (o *InsightsV1CallSummary) GetTrustOk() (*map[string]interface{}, bool)`

GetTrustOk returns a tuple with the Trust field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrust

`func (o *InsightsV1CallSummary) SetTrust(v map[string]interface{})`

SetTrust sets Trust field to given value.

### HasTrust

`func (o *InsightsV1CallSummary) HasTrust() bool`

HasTrust returns a boolean if a field has been set.

### SetTrustNil

`func (o *InsightsV1CallSummary) SetTrustNil(b bool)`

 SetTrustNil sets the value for Trust to be an explicit nil

### UnsetTrust
`func (o *InsightsV1CallSummary) UnsetTrust()`

UnsetTrust ensures that no value is present for Trust, not even an explicit nil
### GetUrl

`func (o *InsightsV1CallSummary) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *InsightsV1CallSummary) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *InsightsV1CallSummary) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *InsightsV1CallSummary) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *InsightsV1CallSummary) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *InsightsV1CallSummary) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


