# NumbersV2RegulatoryComplianceBundleEvaluation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **NullableString** | The SID of the Account that created the resource | [optional] 
**BundleSid** | Pointer to **NullableString** | The unique string that identifies the resource | [optional] 
**DateCreated** | Pointer to **NullableTime** |  | [optional] 
**RegulationSid** | Pointer to **NullableString** | The unique string of a regulation | [optional] 
**Results** | Pointer to **[]map[string]interface{}** | The results of the Evaluation resource | [optional] 
**Sid** | Pointer to **NullableString** | The unique string that identifies the Evaluation resource | [optional] 
**Status** | Pointer to **NullableString** | The compliance status of the Evaluation resource | [optional] 
**Url** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewNumbersV2RegulatoryComplianceBundleEvaluation

`func NewNumbersV2RegulatoryComplianceBundleEvaluation() *NumbersV2RegulatoryComplianceBundleEvaluation`

NewNumbersV2RegulatoryComplianceBundleEvaluation instantiates a new NumbersV2RegulatoryComplianceBundleEvaluation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNumbersV2RegulatoryComplianceBundleEvaluationWithDefaults

`func NewNumbersV2RegulatoryComplianceBundleEvaluationWithDefaults() *NumbersV2RegulatoryComplianceBundleEvaluation`

NewNumbersV2RegulatoryComplianceBundleEvaluationWithDefaults instantiates a new NumbersV2RegulatoryComplianceBundleEvaluation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSid

`func (o *NumbersV2RegulatoryComplianceBundleEvaluation) GetAccountSid() string`

GetAccountSid returns the AccountSid field if non-nil, zero value otherwise.

### GetAccountSidOk

`func (o *NumbersV2RegulatoryComplianceBundleEvaluation) GetAccountSidOk() (*string, bool)`

GetAccountSidOk returns a tuple with the AccountSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSid

`func (o *NumbersV2RegulatoryComplianceBundleEvaluation) SetAccountSid(v string)`

SetAccountSid sets AccountSid field to given value.

### HasAccountSid

`func (o *NumbersV2RegulatoryComplianceBundleEvaluation) HasAccountSid() bool`

HasAccountSid returns a boolean if a field has been set.

### SetAccountSidNil

`func (o *NumbersV2RegulatoryComplianceBundleEvaluation) SetAccountSidNil(b bool)`

 SetAccountSidNil sets the value for AccountSid to be an explicit nil

### UnsetAccountSid
`func (o *NumbersV2RegulatoryComplianceBundleEvaluation) UnsetAccountSid()`

UnsetAccountSid ensures that no value is present for AccountSid, not even an explicit nil
### GetBundleSid

`func (o *NumbersV2RegulatoryComplianceBundleEvaluation) GetBundleSid() string`

GetBundleSid returns the BundleSid field if non-nil, zero value otherwise.

### GetBundleSidOk

`func (o *NumbersV2RegulatoryComplianceBundleEvaluation) GetBundleSidOk() (*string, bool)`

GetBundleSidOk returns a tuple with the BundleSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundleSid

`func (o *NumbersV2RegulatoryComplianceBundleEvaluation) SetBundleSid(v string)`

SetBundleSid sets BundleSid field to given value.

### HasBundleSid

`func (o *NumbersV2RegulatoryComplianceBundleEvaluation) HasBundleSid() bool`

HasBundleSid returns a boolean if a field has been set.

### SetBundleSidNil

`func (o *NumbersV2RegulatoryComplianceBundleEvaluation) SetBundleSidNil(b bool)`

 SetBundleSidNil sets the value for BundleSid to be an explicit nil

### UnsetBundleSid
`func (o *NumbersV2RegulatoryComplianceBundleEvaluation) UnsetBundleSid()`

UnsetBundleSid ensures that no value is present for BundleSid, not even an explicit nil
### GetDateCreated

`func (o *NumbersV2RegulatoryComplianceBundleEvaluation) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *NumbersV2RegulatoryComplianceBundleEvaluation) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *NumbersV2RegulatoryComplianceBundleEvaluation) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *NumbersV2RegulatoryComplianceBundleEvaluation) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### SetDateCreatedNil

`func (o *NumbersV2RegulatoryComplianceBundleEvaluation) SetDateCreatedNil(b bool)`

 SetDateCreatedNil sets the value for DateCreated to be an explicit nil

### UnsetDateCreated
`func (o *NumbersV2RegulatoryComplianceBundleEvaluation) UnsetDateCreated()`

UnsetDateCreated ensures that no value is present for DateCreated, not even an explicit nil
### GetRegulationSid

`func (o *NumbersV2RegulatoryComplianceBundleEvaluation) GetRegulationSid() string`

GetRegulationSid returns the RegulationSid field if non-nil, zero value otherwise.

### GetRegulationSidOk

`func (o *NumbersV2RegulatoryComplianceBundleEvaluation) GetRegulationSidOk() (*string, bool)`

GetRegulationSidOk returns a tuple with the RegulationSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegulationSid

`func (o *NumbersV2RegulatoryComplianceBundleEvaluation) SetRegulationSid(v string)`

SetRegulationSid sets RegulationSid field to given value.

### HasRegulationSid

`func (o *NumbersV2RegulatoryComplianceBundleEvaluation) HasRegulationSid() bool`

HasRegulationSid returns a boolean if a field has been set.

### SetRegulationSidNil

`func (o *NumbersV2RegulatoryComplianceBundleEvaluation) SetRegulationSidNil(b bool)`

 SetRegulationSidNil sets the value for RegulationSid to be an explicit nil

### UnsetRegulationSid
`func (o *NumbersV2RegulatoryComplianceBundleEvaluation) UnsetRegulationSid()`

UnsetRegulationSid ensures that no value is present for RegulationSid, not even an explicit nil
### GetResults

`func (o *NumbersV2RegulatoryComplianceBundleEvaluation) GetResults() []map[string]interface{}`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *NumbersV2RegulatoryComplianceBundleEvaluation) GetResultsOk() (*[]map[string]interface{}, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *NumbersV2RegulatoryComplianceBundleEvaluation) SetResults(v []map[string]interface{})`

SetResults sets Results field to given value.

### HasResults

`func (o *NumbersV2RegulatoryComplianceBundleEvaluation) HasResults() bool`

HasResults returns a boolean if a field has been set.

### SetResultsNil

`func (o *NumbersV2RegulatoryComplianceBundleEvaluation) SetResultsNil(b bool)`

 SetResultsNil sets the value for Results to be an explicit nil

### UnsetResults
`func (o *NumbersV2RegulatoryComplianceBundleEvaluation) UnsetResults()`

UnsetResults ensures that no value is present for Results, not even an explicit nil
### GetSid

`func (o *NumbersV2RegulatoryComplianceBundleEvaluation) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *NumbersV2RegulatoryComplianceBundleEvaluation) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *NumbersV2RegulatoryComplianceBundleEvaluation) SetSid(v string)`

SetSid sets Sid field to given value.

### HasSid

`func (o *NumbersV2RegulatoryComplianceBundleEvaluation) HasSid() bool`

HasSid returns a boolean if a field has been set.

### SetSidNil

`func (o *NumbersV2RegulatoryComplianceBundleEvaluation) SetSidNil(b bool)`

 SetSidNil sets the value for Sid to be an explicit nil

### UnsetSid
`func (o *NumbersV2RegulatoryComplianceBundleEvaluation) UnsetSid()`

UnsetSid ensures that no value is present for Sid, not even an explicit nil
### GetStatus

`func (o *NumbersV2RegulatoryComplianceBundleEvaluation) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *NumbersV2RegulatoryComplianceBundleEvaluation) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *NumbersV2RegulatoryComplianceBundleEvaluation) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *NumbersV2RegulatoryComplianceBundleEvaluation) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *NumbersV2RegulatoryComplianceBundleEvaluation) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *NumbersV2RegulatoryComplianceBundleEvaluation) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetUrl

`func (o *NumbersV2RegulatoryComplianceBundleEvaluation) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *NumbersV2RegulatoryComplianceBundleEvaluation) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *NumbersV2RegulatoryComplianceBundleEvaluation) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *NumbersV2RegulatoryComplianceBundleEvaluation) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *NumbersV2RegulatoryComplianceBundleEvaluation) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *NumbersV2RegulatoryComplianceBundleEvaluation) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


