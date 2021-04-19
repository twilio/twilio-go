# ListBucketResponse

## Properties

Name | Type | Description
------------ | ------------- | -------------
**Buckets** | Pointer to [**[]VerifyV2ServiceRateLimitBucket**](VerifyV2ServiceRateLimitBucket.md) |  | [optional] 
**Meta** | Pointer to [**ListVerificationAttemptResponseMeta**](ListVerificationAttemptResponse_meta.md) |  | [optional] 

## Methods

### NewListBucketResponse

`func NewListBucketResponse() *ListBucketResponse`

NewListBucketResponse instantiates a new ListBucketResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListBucketResponseWithDefaults

`func NewListBucketResponseWithDefaults() *ListBucketResponse`

NewListBucketResponseWithDefaults instantiates a new ListBucketResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBuckets

`func (o *ListBucketResponse) GetBuckets() []VerifyV2ServiceRateLimitBucket`

GetBuckets returns the Buckets field if non-nil, zero value otherwise.

### GetBucketsOk

`func (o *ListBucketResponse) GetBucketsOk() (*[]VerifyV2ServiceRateLimitBucket, bool)`

GetBucketsOk returns a tuple with the Buckets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuckets

`func (o *ListBucketResponse) SetBuckets(v []VerifyV2ServiceRateLimitBucket)`

SetBuckets sets Buckets field to given value.

### HasBuckets

`func (o *ListBucketResponse) HasBuckets() bool`

HasBuckets returns a boolean if a field has been set.

### GetMeta

`func (o *ListBucketResponse) GetMeta() ListVerificationAttemptResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListBucketResponse) GetMetaOk() (*ListVerificationAttemptResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListBucketResponse) SetMeta(v ListVerificationAttemptResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListBucketResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


