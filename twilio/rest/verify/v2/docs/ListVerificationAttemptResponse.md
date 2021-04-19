# ListVerificationAttemptResponse

## Properties

Name | Type | Description
------------ | ------------- | -------------
**Attempts** | Pointer to [**[]VerifyV2VerificationAttempt**](VerifyV2VerificationAttempt.md) |  | [optional] 
**Meta** | Pointer to [**ListVerificationAttemptResponseMeta**](ListVerificationAttemptResponse_meta.md) |  | [optional] 

## Methods

### NewListVerificationAttemptResponse

`func NewListVerificationAttemptResponse() *ListVerificationAttemptResponse`

NewListVerificationAttemptResponse instantiates a new ListVerificationAttemptResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListVerificationAttemptResponseWithDefaults

`func NewListVerificationAttemptResponseWithDefaults() *ListVerificationAttemptResponse`

NewListVerificationAttemptResponseWithDefaults instantiates a new ListVerificationAttemptResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttempts

`func (o *ListVerificationAttemptResponse) GetAttempts() []VerifyV2VerificationAttempt`

GetAttempts returns the Attempts field if non-nil, zero value otherwise.

### GetAttemptsOk

`func (o *ListVerificationAttemptResponse) GetAttemptsOk() (*[]VerifyV2VerificationAttempt, bool)`

GetAttemptsOk returns a tuple with the Attempts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttempts

`func (o *ListVerificationAttemptResponse) SetAttempts(v []VerifyV2VerificationAttempt)`

SetAttempts sets Attempts field to given value.

### HasAttempts

`func (o *ListVerificationAttemptResponse) HasAttempts() bool`

HasAttempts returns a boolean if a field has been set.

### GetMeta

`func (o *ListVerificationAttemptResponse) GetMeta() ListVerificationAttemptResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListVerificationAttemptResponse) GetMetaOk() (*ListVerificationAttemptResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListVerificationAttemptResponse) SetMeta(v ListVerificationAttemptResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListVerificationAttemptResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


