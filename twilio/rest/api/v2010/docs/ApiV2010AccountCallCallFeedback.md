# ApiV2010AccountCallCallFeedback

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **NullableString** | The unique sid that identifies this account | [optional] 
**DateCreated** | Pointer to **NullableString** | The date this resource was created | [optional] 
**DateUpdated** | Pointer to **NullableString** | The date this resource was last updated | [optional] 
**Issues** | Pointer to **[]string** | Issues experienced during the call | [optional] 
**QualityScore** | Pointer to **NullableInt32** | 1 to 5 quality score | [optional] 
**Sid** | Pointer to **NullableString** | A string that uniquely identifies this feedback resource | [optional] 

## Methods

### NewApiV2010AccountCallCallFeedback

`func NewApiV2010AccountCallCallFeedback() *ApiV2010AccountCallCallFeedback`

NewApiV2010AccountCallCallFeedback instantiates a new ApiV2010AccountCallCallFeedback object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiV2010AccountCallCallFeedbackWithDefaults

`func NewApiV2010AccountCallCallFeedbackWithDefaults() *ApiV2010AccountCallCallFeedback`

NewApiV2010AccountCallCallFeedbackWithDefaults instantiates a new ApiV2010AccountCallCallFeedback object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSid

`func (o *ApiV2010AccountCallCallFeedback) GetAccountSid() string`

GetAccountSid returns the AccountSid field if non-nil, zero value otherwise.

### GetAccountSidOk

`func (o *ApiV2010AccountCallCallFeedback) GetAccountSidOk() (*string, bool)`

GetAccountSidOk returns a tuple with the AccountSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSid

`func (o *ApiV2010AccountCallCallFeedback) SetAccountSid(v string)`

SetAccountSid sets AccountSid field to given value.

### HasAccountSid

`func (o *ApiV2010AccountCallCallFeedback) HasAccountSid() bool`

HasAccountSid returns a boolean if a field has been set.

### SetAccountSidNil

`func (o *ApiV2010AccountCallCallFeedback) SetAccountSidNil(b bool)`

 SetAccountSidNil sets the value for AccountSid to be an explicit nil

### UnsetAccountSid
`func (o *ApiV2010AccountCallCallFeedback) UnsetAccountSid()`

UnsetAccountSid ensures that no value is present for AccountSid, not even an explicit nil
### GetDateCreated

`func (o *ApiV2010AccountCallCallFeedback) GetDateCreated() string`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *ApiV2010AccountCallCallFeedback) GetDateCreatedOk() (*string, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *ApiV2010AccountCallCallFeedback) SetDateCreated(v string)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *ApiV2010AccountCallCallFeedback) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### SetDateCreatedNil

`func (o *ApiV2010AccountCallCallFeedback) SetDateCreatedNil(b bool)`

 SetDateCreatedNil sets the value for DateCreated to be an explicit nil

### UnsetDateCreated
`func (o *ApiV2010AccountCallCallFeedback) UnsetDateCreated()`

UnsetDateCreated ensures that no value is present for DateCreated, not even an explicit nil
### GetDateUpdated

`func (o *ApiV2010AccountCallCallFeedback) GetDateUpdated() string`

GetDateUpdated returns the DateUpdated field if non-nil, zero value otherwise.

### GetDateUpdatedOk

`func (o *ApiV2010AccountCallCallFeedback) GetDateUpdatedOk() (*string, bool)`

GetDateUpdatedOk returns a tuple with the DateUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateUpdated

`func (o *ApiV2010AccountCallCallFeedback) SetDateUpdated(v string)`

SetDateUpdated sets DateUpdated field to given value.

### HasDateUpdated

`func (o *ApiV2010AccountCallCallFeedback) HasDateUpdated() bool`

HasDateUpdated returns a boolean if a field has been set.

### SetDateUpdatedNil

`func (o *ApiV2010AccountCallCallFeedback) SetDateUpdatedNil(b bool)`

 SetDateUpdatedNil sets the value for DateUpdated to be an explicit nil

### UnsetDateUpdated
`func (o *ApiV2010AccountCallCallFeedback) UnsetDateUpdated()`

UnsetDateUpdated ensures that no value is present for DateUpdated, not even an explicit nil
### GetIssues

`func (o *ApiV2010AccountCallCallFeedback) GetIssues() []string`

GetIssues returns the Issues field if non-nil, zero value otherwise.

### GetIssuesOk

`func (o *ApiV2010AccountCallCallFeedback) GetIssuesOk() (*[]string, bool)`

GetIssuesOk returns a tuple with the Issues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssues

`func (o *ApiV2010AccountCallCallFeedback) SetIssues(v []string)`

SetIssues sets Issues field to given value.

### HasIssues

`func (o *ApiV2010AccountCallCallFeedback) HasIssues() bool`

HasIssues returns a boolean if a field has been set.

### SetIssuesNil

`func (o *ApiV2010AccountCallCallFeedback) SetIssuesNil(b bool)`

 SetIssuesNil sets the value for Issues to be an explicit nil

### UnsetIssues
`func (o *ApiV2010AccountCallCallFeedback) UnsetIssues()`

UnsetIssues ensures that no value is present for Issues, not even an explicit nil
### GetQualityScore

`func (o *ApiV2010AccountCallCallFeedback) GetQualityScore() int32`

GetQualityScore returns the QualityScore field if non-nil, zero value otherwise.

### GetQualityScoreOk

`func (o *ApiV2010AccountCallCallFeedback) GetQualityScoreOk() (*int32, bool)`

GetQualityScoreOk returns a tuple with the QualityScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQualityScore

`func (o *ApiV2010AccountCallCallFeedback) SetQualityScore(v int32)`

SetQualityScore sets QualityScore field to given value.

### HasQualityScore

`func (o *ApiV2010AccountCallCallFeedback) HasQualityScore() bool`

HasQualityScore returns a boolean if a field has been set.

### SetQualityScoreNil

`func (o *ApiV2010AccountCallCallFeedback) SetQualityScoreNil(b bool)`

 SetQualityScoreNil sets the value for QualityScore to be an explicit nil

### UnsetQualityScore
`func (o *ApiV2010AccountCallCallFeedback) UnsetQualityScore()`

UnsetQualityScore ensures that no value is present for QualityScore, not even an explicit nil
### GetSid

`func (o *ApiV2010AccountCallCallFeedback) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *ApiV2010AccountCallCallFeedback) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *ApiV2010AccountCallCallFeedback) SetSid(v string)`

SetSid sets Sid field to given value.

### HasSid

`func (o *ApiV2010AccountCallCallFeedback) HasSid() bool`

HasSid returns a boolean if a field has been set.

### SetSidNil

`func (o *ApiV2010AccountCallCallFeedback) SetSidNil(b bool)`

 SetSidNil sets the value for Sid to be an explicit nil

### UnsetSid
`func (o *ApiV2010AccountCallCallFeedback) UnsetSid()`

UnsetSid ensures that no value is present for Sid, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


