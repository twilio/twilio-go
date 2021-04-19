# ApiV2010AccountCallCallFeedbackSummary

## Properties

Name | Type | Description
------------ | ------------- | -------------
**AccountSid** | Pointer to **NullableString** | The unique sid that identifies this account | [optional] 
**CallCount** | Pointer to **NullableInt32** | The total number of calls | [optional] 
**CallFeedbackCount** | Pointer to **NullableInt32** | The total number of calls with a feedback entry | [optional] 
**DateCreated** | Pointer to **NullableString** | The date this resource was created | [optional] 
**DateUpdated** | Pointer to **NullableString** | The date this resource was last updated | [optional] 
**EndDate** | Pointer to **NullableTime** | The latest feedback entry date in the summary | [optional] 
**IncludeSubaccounts** | Pointer to **NullableBool** | Whether the feedback summary includes subaccounts | [optional] 
**Issues** | Pointer to [**[]ApiV2010AccountCallCallFeedbackSummaryIssues**](ApiV2010AccountCallCallFeedbackSummaryIssues.md) | Issues experienced during the call | [optional] 
**QualityScoreAverage** | Pointer to **NullableFloat32** | The average QualityScore of the feedback entries | [optional] 
**QualityScoreMedian** | Pointer to **NullableFloat32** | The median QualityScore of the feedback entries | [optional] 
**QualityScoreStandardDeviation** | Pointer to **NullableFloat32** | The standard deviation of the quality scores | [optional] 
**Sid** | Pointer to **NullableString** | A string that uniquely identifies this feedback entry | [optional] 
**StartDate** | Pointer to **NullableTime** | The earliest feedback entry date in the summary | [optional] 
**Status** | Pointer to **NullableString** | The status of the feedback summary | [optional] 

## Methods

### NewApiV2010AccountCallCallFeedbackSummary

`func NewApiV2010AccountCallCallFeedbackSummary() *ApiV2010AccountCallCallFeedbackSummary`

NewApiV2010AccountCallCallFeedbackSummary instantiates a new ApiV2010AccountCallCallFeedbackSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiV2010AccountCallCallFeedbackSummaryWithDefaults

`func NewApiV2010AccountCallCallFeedbackSummaryWithDefaults() *ApiV2010AccountCallCallFeedbackSummary`

NewApiV2010AccountCallCallFeedbackSummaryWithDefaults instantiates a new ApiV2010AccountCallCallFeedbackSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSid

`func (o *ApiV2010AccountCallCallFeedbackSummary) GetAccountSid() string`

GetAccountSid returns the AccountSid field if non-nil, zero value otherwise.

### GetAccountSidOk

`func (o *ApiV2010AccountCallCallFeedbackSummary) GetAccountSidOk() (*string, bool)`

GetAccountSidOk returns a tuple with the AccountSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSid

`func (o *ApiV2010AccountCallCallFeedbackSummary) SetAccountSid(v string)`

SetAccountSid sets AccountSid field to given value.

### HasAccountSid

`func (o *ApiV2010AccountCallCallFeedbackSummary) HasAccountSid() bool`

HasAccountSid returns a boolean if a field has been set.

### SetAccountSidNil

`func (o *ApiV2010AccountCallCallFeedbackSummary) SetAccountSidNil(b bool)`

 SetAccountSidNil sets the value for AccountSid to be an explicit nil

### UnsetAccountSid
`func (o *ApiV2010AccountCallCallFeedbackSummary) UnsetAccountSid()`

UnsetAccountSid ensures that no value is present for AccountSid, not even an explicit nil
### GetCallCount

`func (o *ApiV2010AccountCallCallFeedbackSummary) GetCallCount() int32`

GetCallCount returns the CallCount field if non-nil, zero value otherwise.

### GetCallCountOk

`func (o *ApiV2010AccountCallCallFeedbackSummary) GetCallCountOk() (*int32, bool)`

GetCallCountOk returns a tuple with the CallCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallCount

`func (o *ApiV2010AccountCallCallFeedbackSummary) SetCallCount(v int32)`

SetCallCount sets CallCount field to given value.

### HasCallCount

`func (o *ApiV2010AccountCallCallFeedbackSummary) HasCallCount() bool`

HasCallCount returns a boolean if a field has been set.

### SetCallCountNil

`func (o *ApiV2010AccountCallCallFeedbackSummary) SetCallCountNil(b bool)`

 SetCallCountNil sets the value for CallCount to be an explicit nil

### UnsetCallCount
`func (o *ApiV2010AccountCallCallFeedbackSummary) UnsetCallCount()`

UnsetCallCount ensures that no value is present for CallCount, not even an explicit nil
### GetCallFeedbackCount

`func (o *ApiV2010AccountCallCallFeedbackSummary) GetCallFeedbackCount() int32`

GetCallFeedbackCount returns the CallFeedbackCount field if non-nil, zero value otherwise.

### GetCallFeedbackCountOk

`func (o *ApiV2010AccountCallCallFeedbackSummary) GetCallFeedbackCountOk() (*int32, bool)`

GetCallFeedbackCountOk returns a tuple with the CallFeedbackCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallFeedbackCount

`func (o *ApiV2010AccountCallCallFeedbackSummary) SetCallFeedbackCount(v int32)`

SetCallFeedbackCount sets CallFeedbackCount field to given value.

### HasCallFeedbackCount

`func (o *ApiV2010AccountCallCallFeedbackSummary) HasCallFeedbackCount() bool`

HasCallFeedbackCount returns a boolean if a field has been set.

### SetCallFeedbackCountNil

`func (o *ApiV2010AccountCallCallFeedbackSummary) SetCallFeedbackCountNil(b bool)`

 SetCallFeedbackCountNil sets the value for CallFeedbackCount to be an explicit nil

### UnsetCallFeedbackCount
`func (o *ApiV2010AccountCallCallFeedbackSummary) UnsetCallFeedbackCount()`

UnsetCallFeedbackCount ensures that no value is present for CallFeedbackCount, not even an explicit nil
### GetDateCreated

`func (o *ApiV2010AccountCallCallFeedbackSummary) GetDateCreated() string`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *ApiV2010AccountCallCallFeedbackSummary) GetDateCreatedOk() (*string, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *ApiV2010AccountCallCallFeedbackSummary) SetDateCreated(v string)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *ApiV2010AccountCallCallFeedbackSummary) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### SetDateCreatedNil

`func (o *ApiV2010AccountCallCallFeedbackSummary) SetDateCreatedNil(b bool)`

 SetDateCreatedNil sets the value for DateCreated to be an explicit nil

### UnsetDateCreated
`func (o *ApiV2010AccountCallCallFeedbackSummary) UnsetDateCreated()`

UnsetDateCreated ensures that no value is present for DateCreated, not even an explicit nil
### GetDateUpdated

`func (o *ApiV2010AccountCallCallFeedbackSummary) GetDateUpdated() string`

GetDateUpdated returns the DateUpdated field if non-nil, zero value otherwise.

### GetDateUpdatedOk

`func (o *ApiV2010AccountCallCallFeedbackSummary) GetDateUpdatedOk() (*string, bool)`

GetDateUpdatedOk returns a tuple with the DateUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateUpdated

`func (o *ApiV2010AccountCallCallFeedbackSummary) SetDateUpdated(v string)`

SetDateUpdated sets DateUpdated field to given value.

### HasDateUpdated

`func (o *ApiV2010AccountCallCallFeedbackSummary) HasDateUpdated() bool`

HasDateUpdated returns a boolean if a field has been set.

### SetDateUpdatedNil

`func (o *ApiV2010AccountCallCallFeedbackSummary) SetDateUpdatedNil(b bool)`

 SetDateUpdatedNil sets the value for DateUpdated to be an explicit nil

### UnsetDateUpdated
`func (o *ApiV2010AccountCallCallFeedbackSummary) UnsetDateUpdated()`

UnsetDateUpdated ensures that no value is present for DateUpdated, not even an explicit nil
### GetEndDate

`func (o *ApiV2010AccountCallCallFeedbackSummary) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *ApiV2010AccountCallCallFeedbackSummary) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *ApiV2010AccountCallCallFeedbackSummary) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *ApiV2010AccountCallCallFeedbackSummary) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### SetEndDateNil

`func (o *ApiV2010AccountCallCallFeedbackSummary) SetEndDateNil(b bool)`

 SetEndDateNil sets the value for EndDate to be an explicit nil

### UnsetEndDate
`func (o *ApiV2010AccountCallCallFeedbackSummary) UnsetEndDate()`

UnsetEndDate ensures that no value is present for EndDate, not even an explicit nil
### GetIncludeSubaccounts

`func (o *ApiV2010AccountCallCallFeedbackSummary) GetIncludeSubaccounts() bool`

GetIncludeSubaccounts returns the IncludeSubaccounts field if non-nil, zero value otherwise.

### GetIncludeSubaccountsOk

`func (o *ApiV2010AccountCallCallFeedbackSummary) GetIncludeSubaccountsOk() (*bool, bool)`

GetIncludeSubaccountsOk returns a tuple with the IncludeSubaccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeSubaccounts

`func (o *ApiV2010AccountCallCallFeedbackSummary) SetIncludeSubaccounts(v bool)`

SetIncludeSubaccounts sets IncludeSubaccounts field to given value.

### HasIncludeSubaccounts

`func (o *ApiV2010AccountCallCallFeedbackSummary) HasIncludeSubaccounts() bool`

HasIncludeSubaccounts returns a boolean if a field has been set.

### SetIncludeSubaccountsNil

`func (o *ApiV2010AccountCallCallFeedbackSummary) SetIncludeSubaccountsNil(b bool)`

 SetIncludeSubaccountsNil sets the value for IncludeSubaccounts to be an explicit nil

### UnsetIncludeSubaccounts
`func (o *ApiV2010AccountCallCallFeedbackSummary) UnsetIncludeSubaccounts()`

UnsetIncludeSubaccounts ensures that no value is present for IncludeSubaccounts, not even an explicit nil
### GetIssues

`func (o *ApiV2010AccountCallCallFeedbackSummary) GetIssues() []ApiV2010AccountCallCallFeedbackSummaryIssues`

GetIssues returns the Issues field if non-nil, zero value otherwise.

### GetIssuesOk

`func (o *ApiV2010AccountCallCallFeedbackSummary) GetIssuesOk() (*[]ApiV2010AccountCallCallFeedbackSummaryIssues, bool)`

GetIssuesOk returns a tuple with the Issues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssues

`func (o *ApiV2010AccountCallCallFeedbackSummary) SetIssues(v []ApiV2010AccountCallCallFeedbackSummaryIssues)`

SetIssues sets Issues field to given value.

### HasIssues

`func (o *ApiV2010AccountCallCallFeedbackSummary) HasIssues() bool`

HasIssues returns a boolean if a field has been set.

### SetIssuesNil

`func (o *ApiV2010AccountCallCallFeedbackSummary) SetIssuesNil(b bool)`

 SetIssuesNil sets the value for Issues to be an explicit nil

### UnsetIssues
`func (o *ApiV2010AccountCallCallFeedbackSummary) UnsetIssues()`

UnsetIssues ensures that no value is present for Issues, not even an explicit nil
### GetQualityScoreAverage

`func (o *ApiV2010AccountCallCallFeedbackSummary) GetQualityScoreAverage() float32`

GetQualityScoreAverage returns the QualityScoreAverage field if non-nil, zero value otherwise.

### GetQualityScoreAverageOk

`func (o *ApiV2010AccountCallCallFeedbackSummary) GetQualityScoreAverageOk() (*float32, bool)`

GetQualityScoreAverageOk returns a tuple with the QualityScoreAverage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQualityScoreAverage

`func (o *ApiV2010AccountCallCallFeedbackSummary) SetQualityScoreAverage(v float32)`

SetQualityScoreAverage sets QualityScoreAverage field to given value.

### HasQualityScoreAverage

`func (o *ApiV2010AccountCallCallFeedbackSummary) HasQualityScoreAverage() bool`

HasQualityScoreAverage returns a boolean if a field has been set.

### SetQualityScoreAverageNil

`func (o *ApiV2010AccountCallCallFeedbackSummary) SetQualityScoreAverageNil(b bool)`

 SetQualityScoreAverageNil sets the value for QualityScoreAverage to be an explicit nil

### UnsetQualityScoreAverage
`func (o *ApiV2010AccountCallCallFeedbackSummary) UnsetQualityScoreAverage()`

UnsetQualityScoreAverage ensures that no value is present for QualityScoreAverage, not even an explicit nil
### GetQualityScoreMedian

`func (o *ApiV2010AccountCallCallFeedbackSummary) GetQualityScoreMedian() float32`

GetQualityScoreMedian returns the QualityScoreMedian field if non-nil, zero value otherwise.

### GetQualityScoreMedianOk

`func (o *ApiV2010AccountCallCallFeedbackSummary) GetQualityScoreMedianOk() (*float32, bool)`

GetQualityScoreMedianOk returns a tuple with the QualityScoreMedian field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQualityScoreMedian

`func (o *ApiV2010AccountCallCallFeedbackSummary) SetQualityScoreMedian(v float32)`

SetQualityScoreMedian sets QualityScoreMedian field to given value.

### HasQualityScoreMedian

`func (o *ApiV2010AccountCallCallFeedbackSummary) HasQualityScoreMedian() bool`

HasQualityScoreMedian returns a boolean if a field has been set.

### SetQualityScoreMedianNil

`func (o *ApiV2010AccountCallCallFeedbackSummary) SetQualityScoreMedianNil(b bool)`

 SetQualityScoreMedianNil sets the value for QualityScoreMedian to be an explicit nil

### UnsetQualityScoreMedian
`func (o *ApiV2010AccountCallCallFeedbackSummary) UnsetQualityScoreMedian()`

UnsetQualityScoreMedian ensures that no value is present for QualityScoreMedian, not even an explicit nil
### GetQualityScoreStandardDeviation

`func (o *ApiV2010AccountCallCallFeedbackSummary) GetQualityScoreStandardDeviation() float32`

GetQualityScoreStandardDeviation returns the QualityScoreStandardDeviation field if non-nil, zero value otherwise.

### GetQualityScoreStandardDeviationOk

`func (o *ApiV2010AccountCallCallFeedbackSummary) GetQualityScoreStandardDeviationOk() (*float32, bool)`

GetQualityScoreStandardDeviationOk returns a tuple with the QualityScoreStandardDeviation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQualityScoreStandardDeviation

`func (o *ApiV2010AccountCallCallFeedbackSummary) SetQualityScoreStandardDeviation(v float32)`

SetQualityScoreStandardDeviation sets QualityScoreStandardDeviation field to given value.

### HasQualityScoreStandardDeviation

`func (o *ApiV2010AccountCallCallFeedbackSummary) HasQualityScoreStandardDeviation() bool`

HasQualityScoreStandardDeviation returns a boolean if a field has been set.

### SetQualityScoreStandardDeviationNil

`func (o *ApiV2010AccountCallCallFeedbackSummary) SetQualityScoreStandardDeviationNil(b bool)`

 SetQualityScoreStandardDeviationNil sets the value for QualityScoreStandardDeviation to be an explicit nil

### UnsetQualityScoreStandardDeviation
`func (o *ApiV2010AccountCallCallFeedbackSummary) UnsetQualityScoreStandardDeviation()`

UnsetQualityScoreStandardDeviation ensures that no value is present for QualityScoreStandardDeviation, not even an explicit nil
### GetSid

`func (o *ApiV2010AccountCallCallFeedbackSummary) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *ApiV2010AccountCallCallFeedbackSummary) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *ApiV2010AccountCallCallFeedbackSummary) SetSid(v string)`

SetSid sets Sid field to given value.

### HasSid

`func (o *ApiV2010AccountCallCallFeedbackSummary) HasSid() bool`

HasSid returns a boolean if a field has been set.

### SetSidNil

`func (o *ApiV2010AccountCallCallFeedbackSummary) SetSidNil(b bool)`

 SetSidNil sets the value for Sid to be an explicit nil

### UnsetSid
`func (o *ApiV2010AccountCallCallFeedbackSummary) UnsetSid()`

UnsetSid ensures that no value is present for Sid, not even an explicit nil
### GetStartDate

`func (o *ApiV2010AccountCallCallFeedbackSummary) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *ApiV2010AccountCallCallFeedbackSummary) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *ApiV2010AccountCallCallFeedbackSummary) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *ApiV2010AccountCallCallFeedbackSummary) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### SetStartDateNil

`func (o *ApiV2010AccountCallCallFeedbackSummary) SetStartDateNil(b bool)`

 SetStartDateNil sets the value for StartDate to be an explicit nil

### UnsetStartDate
`func (o *ApiV2010AccountCallCallFeedbackSummary) UnsetStartDate()`

UnsetStartDate ensures that no value is present for StartDate, not even an explicit nil
### GetStatus

`func (o *ApiV2010AccountCallCallFeedbackSummary) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ApiV2010AccountCallCallFeedbackSummary) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ApiV2010AccountCallCallFeedbackSummary) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ApiV2010AccountCallCallFeedbackSummary) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *ApiV2010AccountCallCallFeedbackSummary) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *ApiV2010AccountCallCallFeedbackSummary) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


