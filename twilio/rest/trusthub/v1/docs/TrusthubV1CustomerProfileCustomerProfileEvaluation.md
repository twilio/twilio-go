# TrusthubV1CustomerProfileCustomerProfileEvaluation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **NullableString** | The SID of the Account that created the resource | [optional] 
**CustomerProfileSid** | Pointer to **NullableString** | The unique string that identifies the resource | [optional] 
**DateCreated** | Pointer to **NullableTime** |  | [optional] 
**PolicySid** | Pointer to **NullableString** | The unique string of a policy | [optional] 
**Results** | Pointer to **[]map[string]interface{}** | The results of the Evaluation resource | [optional] 
**Sid** | Pointer to **NullableString** | The unique string that identifies the Evaluation resource | [optional] 
**Status** | Pointer to **NullableString** | The compliance status of the Evaluation resource | [optional] 
**Url** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewTrusthubV1CustomerProfileCustomerProfileEvaluation

`func NewTrusthubV1CustomerProfileCustomerProfileEvaluation() *TrusthubV1CustomerProfileCustomerProfileEvaluation`

NewTrusthubV1CustomerProfileCustomerProfileEvaluation instantiates a new TrusthubV1CustomerProfileCustomerProfileEvaluation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrusthubV1CustomerProfileCustomerProfileEvaluationWithDefaults

`func NewTrusthubV1CustomerProfileCustomerProfileEvaluationWithDefaults() *TrusthubV1CustomerProfileCustomerProfileEvaluation`

NewTrusthubV1CustomerProfileCustomerProfileEvaluationWithDefaults instantiates a new TrusthubV1CustomerProfileCustomerProfileEvaluation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSid

`func (o *TrusthubV1CustomerProfileCustomerProfileEvaluation) GetAccountSid() string`

GetAccountSid returns the AccountSid field if non-nil, zero value otherwise.

### GetAccountSidOk

`func (o *TrusthubV1CustomerProfileCustomerProfileEvaluation) GetAccountSidOk() (*string, bool)`

GetAccountSidOk returns a tuple with the AccountSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSid

`func (o *TrusthubV1CustomerProfileCustomerProfileEvaluation) SetAccountSid(v string)`

SetAccountSid sets AccountSid field to given value.

### HasAccountSid

`func (o *TrusthubV1CustomerProfileCustomerProfileEvaluation) HasAccountSid() bool`

HasAccountSid returns a boolean if a field has been set.

### SetAccountSidNil

`func (o *TrusthubV1CustomerProfileCustomerProfileEvaluation) SetAccountSidNil(b bool)`

 SetAccountSidNil sets the value for AccountSid to be an explicit nil

### UnsetAccountSid
`func (o *TrusthubV1CustomerProfileCustomerProfileEvaluation) UnsetAccountSid()`

UnsetAccountSid ensures that no value is present for AccountSid, not even an explicit nil
### GetCustomerProfileSid

`func (o *TrusthubV1CustomerProfileCustomerProfileEvaluation) GetCustomerProfileSid() string`

GetCustomerProfileSid returns the CustomerProfileSid field if non-nil, zero value otherwise.

### GetCustomerProfileSidOk

`func (o *TrusthubV1CustomerProfileCustomerProfileEvaluation) GetCustomerProfileSidOk() (*string, bool)`

GetCustomerProfileSidOk returns a tuple with the CustomerProfileSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerProfileSid

`func (o *TrusthubV1CustomerProfileCustomerProfileEvaluation) SetCustomerProfileSid(v string)`

SetCustomerProfileSid sets CustomerProfileSid field to given value.

### HasCustomerProfileSid

`func (o *TrusthubV1CustomerProfileCustomerProfileEvaluation) HasCustomerProfileSid() bool`

HasCustomerProfileSid returns a boolean if a field has been set.

### SetCustomerProfileSidNil

`func (o *TrusthubV1CustomerProfileCustomerProfileEvaluation) SetCustomerProfileSidNil(b bool)`

 SetCustomerProfileSidNil sets the value for CustomerProfileSid to be an explicit nil

### UnsetCustomerProfileSid
`func (o *TrusthubV1CustomerProfileCustomerProfileEvaluation) UnsetCustomerProfileSid()`

UnsetCustomerProfileSid ensures that no value is present for CustomerProfileSid, not even an explicit nil
### GetDateCreated

`func (o *TrusthubV1CustomerProfileCustomerProfileEvaluation) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *TrusthubV1CustomerProfileCustomerProfileEvaluation) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *TrusthubV1CustomerProfileCustomerProfileEvaluation) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *TrusthubV1CustomerProfileCustomerProfileEvaluation) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### SetDateCreatedNil

`func (o *TrusthubV1CustomerProfileCustomerProfileEvaluation) SetDateCreatedNil(b bool)`

 SetDateCreatedNil sets the value for DateCreated to be an explicit nil

### UnsetDateCreated
`func (o *TrusthubV1CustomerProfileCustomerProfileEvaluation) UnsetDateCreated()`

UnsetDateCreated ensures that no value is present for DateCreated, not even an explicit nil
### GetPolicySid

`func (o *TrusthubV1CustomerProfileCustomerProfileEvaluation) GetPolicySid() string`

GetPolicySid returns the PolicySid field if non-nil, zero value otherwise.

### GetPolicySidOk

`func (o *TrusthubV1CustomerProfileCustomerProfileEvaluation) GetPolicySidOk() (*string, bool)`

GetPolicySidOk returns a tuple with the PolicySid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicySid

`func (o *TrusthubV1CustomerProfileCustomerProfileEvaluation) SetPolicySid(v string)`

SetPolicySid sets PolicySid field to given value.

### HasPolicySid

`func (o *TrusthubV1CustomerProfileCustomerProfileEvaluation) HasPolicySid() bool`

HasPolicySid returns a boolean if a field has been set.

### SetPolicySidNil

`func (o *TrusthubV1CustomerProfileCustomerProfileEvaluation) SetPolicySidNil(b bool)`

 SetPolicySidNil sets the value for PolicySid to be an explicit nil

### UnsetPolicySid
`func (o *TrusthubV1CustomerProfileCustomerProfileEvaluation) UnsetPolicySid()`

UnsetPolicySid ensures that no value is present for PolicySid, not even an explicit nil
### GetResults

`func (o *TrusthubV1CustomerProfileCustomerProfileEvaluation) GetResults() []map[string]interface{}`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *TrusthubV1CustomerProfileCustomerProfileEvaluation) GetResultsOk() (*[]map[string]interface{}, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *TrusthubV1CustomerProfileCustomerProfileEvaluation) SetResults(v []map[string]interface{})`

SetResults sets Results field to given value.

### HasResults

`func (o *TrusthubV1CustomerProfileCustomerProfileEvaluation) HasResults() bool`

HasResults returns a boolean if a field has been set.

### SetResultsNil

`func (o *TrusthubV1CustomerProfileCustomerProfileEvaluation) SetResultsNil(b bool)`

 SetResultsNil sets the value for Results to be an explicit nil

### UnsetResults
`func (o *TrusthubV1CustomerProfileCustomerProfileEvaluation) UnsetResults()`

UnsetResults ensures that no value is present for Results, not even an explicit nil
### GetSid

`func (o *TrusthubV1CustomerProfileCustomerProfileEvaluation) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *TrusthubV1CustomerProfileCustomerProfileEvaluation) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *TrusthubV1CustomerProfileCustomerProfileEvaluation) SetSid(v string)`

SetSid sets Sid field to given value.

### HasSid

`func (o *TrusthubV1CustomerProfileCustomerProfileEvaluation) HasSid() bool`

HasSid returns a boolean if a field has been set.

### SetSidNil

`func (o *TrusthubV1CustomerProfileCustomerProfileEvaluation) SetSidNil(b bool)`

 SetSidNil sets the value for Sid to be an explicit nil

### UnsetSid
`func (o *TrusthubV1CustomerProfileCustomerProfileEvaluation) UnsetSid()`

UnsetSid ensures that no value is present for Sid, not even an explicit nil
### GetStatus

`func (o *TrusthubV1CustomerProfileCustomerProfileEvaluation) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TrusthubV1CustomerProfileCustomerProfileEvaluation) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TrusthubV1CustomerProfileCustomerProfileEvaluation) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TrusthubV1CustomerProfileCustomerProfileEvaluation) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *TrusthubV1CustomerProfileCustomerProfileEvaluation) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *TrusthubV1CustomerProfileCustomerProfileEvaluation) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetUrl

`func (o *TrusthubV1CustomerProfileCustomerProfileEvaluation) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *TrusthubV1CustomerProfileCustomerProfileEvaluation) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *TrusthubV1CustomerProfileCustomerProfileEvaluation) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *TrusthubV1CustomerProfileCustomerProfileEvaluation) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *TrusthubV1CustomerProfileCustomerProfileEvaluation) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *TrusthubV1CustomerProfileCustomerProfileEvaluation) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


