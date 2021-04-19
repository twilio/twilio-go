# ListChallengeResponse

## Properties

Name | Type | Description
------------ | ------------- | -------------
**Challenges** | Pointer to [**[]VerifyV2ServiceEntityChallenge**](VerifyV2ServiceEntityChallenge.md) |  | [optional] 
**Meta** | Pointer to [**ListVerificationAttemptResponseMeta**](ListVerificationAttemptResponse_meta.md) |  | [optional] 

## Methods

### NewListChallengeResponse

`func NewListChallengeResponse() *ListChallengeResponse`

NewListChallengeResponse instantiates a new ListChallengeResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListChallengeResponseWithDefaults

`func NewListChallengeResponseWithDefaults() *ListChallengeResponse`

NewListChallengeResponseWithDefaults instantiates a new ListChallengeResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChallenges

`func (o *ListChallengeResponse) GetChallenges() []VerifyV2ServiceEntityChallenge`

GetChallenges returns the Challenges field if non-nil, zero value otherwise.

### GetChallengesOk

`func (o *ListChallengeResponse) GetChallengesOk() (*[]VerifyV2ServiceEntityChallenge, bool)`

GetChallengesOk returns a tuple with the Challenges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChallenges

`func (o *ListChallengeResponse) SetChallenges(v []VerifyV2ServiceEntityChallenge)`

SetChallenges sets Challenges field to given value.

### HasChallenges

`func (o *ListChallengeResponse) HasChallenges() bool`

HasChallenges returns a boolean if a field has been set.

### GetMeta

`func (o *ListChallengeResponse) GetMeta() ListVerificationAttemptResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListChallengeResponse) GetMetaOk() (*ListVerificationAttemptResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListChallengeResponse) SetMeta(v ListVerificationAttemptResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListChallengeResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


