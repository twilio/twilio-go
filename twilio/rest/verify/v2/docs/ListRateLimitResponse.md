# ListRateLimitResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**ListVerificationAttemptResponseMeta**](ListVerificationAttemptResponse_meta.md) |  | [optional] 
**RateLimits** | Pointer to [**[]VerifyV2ServiceRateLimit**](VerifyV2ServiceRateLimit.md) |  | [optional] 

## Methods

### NewListRateLimitResponse

`func NewListRateLimitResponse() *ListRateLimitResponse`

NewListRateLimitResponse instantiates a new ListRateLimitResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListRateLimitResponseWithDefaults

`func NewListRateLimitResponseWithDefaults() *ListRateLimitResponse`

NewListRateLimitResponseWithDefaults instantiates a new ListRateLimitResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *ListRateLimitResponse) GetMeta() ListVerificationAttemptResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListRateLimitResponse) GetMetaOk() (*ListVerificationAttemptResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListRateLimitResponse) SetMeta(v ListVerificationAttemptResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListRateLimitResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetRateLimits

`func (o *ListRateLimitResponse) GetRateLimits() []VerifyV2ServiceRateLimit`

GetRateLimits returns the RateLimits field if non-nil, zero value otherwise.

### GetRateLimitsOk

`func (o *ListRateLimitResponse) GetRateLimitsOk() (*[]VerifyV2ServiceRateLimit, bool)`

GetRateLimitsOk returns a tuple with the RateLimits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateLimits

`func (o *ListRateLimitResponse) SetRateLimits(v []VerifyV2ServiceRateLimit)`

SetRateLimits sets RateLimits field to given value.

### HasRateLimits

`func (o *ListRateLimitResponse) HasRateLimits() bool`

HasRateLimits returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


