# VerifyV2ServiceRateLimitBucket

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **NullableString** | The SID of the Account that created the resource | [optional] 
**DateCreated** | Pointer to **NullableTime** | The RFC 2822 date and time in GMT when the resource was created | [optional] 
**DateUpdated** | Pointer to **NullableTime** | The RFC 2822 date and time in GMT when the resource was last updated | [optional] 
**Interval** | Pointer to **NullableInt32** | Number of seconds that the rate limit will be enforced over. | [optional] 
**Max** | Pointer to **NullableInt32** | Max number of requests. | [optional] 
**RateLimitSid** | Pointer to **NullableString** | Rate Limit Sid. | [optional] 
**ServiceSid** | Pointer to **NullableString** | The SID of the Service that the resource is associated with | [optional] 
**Sid** | Pointer to **NullableString** | A string that uniquely identifies this Bucket. | [optional] 
**Url** | Pointer to **NullableString** | The URL of this resource. | [optional] 

## Methods

### NewVerifyV2ServiceRateLimitBucket

`func NewVerifyV2ServiceRateLimitBucket() *VerifyV2ServiceRateLimitBucket`

NewVerifyV2ServiceRateLimitBucket instantiates a new VerifyV2ServiceRateLimitBucket object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVerifyV2ServiceRateLimitBucketWithDefaults

`func NewVerifyV2ServiceRateLimitBucketWithDefaults() *VerifyV2ServiceRateLimitBucket`

NewVerifyV2ServiceRateLimitBucketWithDefaults instantiates a new VerifyV2ServiceRateLimitBucket object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSid

`func (o *VerifyV2ServiceRateLimitBucket) GetAccountSid() string`

GetAccountSid returns the AccountSid field if non-nil, zero value otherwise.

### GetAccountSidOk

`func (o *VerifyV2ServiceRateLimitBucket) GetAccountSidOk() (*string, bool)`

GetAccountSidOk returns a tuple with the AccountSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSid

`func (o *VerifyV2ServiceRateLimitBucket) SetAccountSid(v string)`

SetAccountSid sets AccountSid field to given value.

### HasAccountSid

`func (o *VerifyV2ServiceRateLimitBucket) HasAccountSid() bool`

HasAccountSid returns a boolean if a field has been set.

### SetAccountSidNil

`func (o *VerifyV2ServiceRateLimitBucket) SetAccountSidNil(b bool)`

 SetAccountSidNil sets the value for AccountSid to be an explicit nil

### UnsetAccountSid
`func (o *VerifyV2ServiceRateLimitBucket) UnsetAccountSid()`

UnsetAccountSid ensures that no value is present for AccountSid, not even an explicit nil
### GetDateCreated

`func (o *VerifyV2ServiceRateLimitBucket) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *VerifyV2ServiceRateLimitBucket) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *VerifyV2ServiceRateLimitBucket) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *VerifyV2ServiceRateLimitBucket) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### SetDateCreatedNil

`func (o *VerifyV2ServiceRateLimitBucket) SetDateCreatedNil(b bool)`

 SetDateCreatedNil sets the value for DateCreated to be an explicit nil

### UnsetDateCreated
`func (o *VerifyV2ServiceRateLimitBucket) UnsetDateCreated()`

UnsetDateCreated ensures that no value is present for DateCreated, not even an explicit nil
### GetDateUpdated

`func (o *VerifyV2ServiceRateLimitBucket) GetDateUpdated() time.Time`

GetDateUpdated returns the DateUpdated field if non-nil, zero value otherwise.

### GetDateUpdatedOk

`func (o *VerifyV2ServiceRateLimitBucket) GetDateUpdatedOk() (*time.Time, bool)`

GetDateUpdatedOk returns a tuple with the DateUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateUpdated

`func (o *VerifyV2ServiceRateLimitBucket) SetDateUpdated(v time.Time)`

SetDateUpdated sets DateUpdated field to given value.

### HasDateUpdated

`func (o *VerifyV2ServiceRateLimitBucket) HasDateUpdated() bool`

HasDateUpdated returns a boolean if a field has been set.

### SetDateUpdatedNil

`func (o *VerifyV2ServiceRateLimitBucket) SetDateUpdatedNil(b bool)`

 SetDateUpdatedNil sets the value for DateUpdated to be an explicit nil

### UnsetDateUpdated
`func (o *VerifyV2ServiceRateLimitBucket) UnsetDateUpdated()`

UnsetDateUpdated ensures that no value is present for DateUpdated, not even an explicit nil
### GetInterval

`func (o *VerifyV2ServiceRateLimitBucket) GetInterval() int32`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *VerifyV2ServiceRateLimitBucket) GetIntervalOk() (*int32, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *VerifyV2ServiceRateLimitBucket) SetInterval(v int32)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *VerifyV2ServiceRateLimitBucket) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### SetIntervalNil

`func (o *VerifyV2ServiceRateLimitBucket) SetIntervalNil(b bool)`

 SetIntervalNil sets the value for Interval to be an explicit nil

### UnsetInterval
`func (o *VerifyV2ServiceRateLimitBucket) UnsetInterval()`

UnsetInterval ensures that no value is present for Interval, not even an explicit nil
### GetMax

`func (o *VerifyV2ServiceRateLimitBucket) GetMax() int32`

GetMax returns the Max field if non-nil, zero value otherwise.

### GetMaxOk

`func (o *VerifyV2ServiceRateLimitBucket) GetMaxOk() (*int32, bool)`

GetMaxOk returns a tuple with the Max field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMax

`func (o *VerifyV2ServiceRateLimitBucket) SetMax(v int32)`

SetMax sets Max field to given value.

### HasMax

`func (o *VerifyV2ServiceRateLimitBucket) HasMax() bool`

HasMax returns a boolean if a field has been set.

### SetMaxNil

`func (o *VerifyV2ServiceRateLimitBucket) SetMaxNil(b bool)`

 SetMaxNil sets the value for Max to be an explicit nil

### UnsetMax
`func (o *VerifyV2ServiceRateLimitBucket) UnsetMax()`

UnsetMax ensures that no value is present for Max, not even an explicit nil
### GetRateLimitSid

`func (o *VerifyV2ServiceRateLimitBucket) GetRateLimitSid() string`

GetRateLimitSid returns the RateLimitSid field if non-nil, zero value otherwise.

### GetRateLimitSidOk

`func (o *VerifyV2ServiceRateLimitBucket) GetRateLimitSidOk() (*string, bool)`

GetRateLimitSidOk returns a tuple with the RateLimitSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateLimitSid

`func (o *VerifyV2ServiceRateLimitBucket) SetRateLimitSid(v string)`

SetRateLimitSid sets RateLimitSid field to given value.

### HasRateLimitSid

`func (o *VerifyV2ServiceRateLimitBucket) HasRateLimitSid() bool`

HasRateLimitSid returns a boolean if a field has been set.

### SetRateLimitSidNil

`func (o *VerifyV2ServiceRateLimitBucket) SetRateLimitSidNil(b bool)`

 SetRateLimitSidNil sets the value for RateLimitSid to be an explicit nil

### UnsetRateLimitSid
`func (o *VerifyV2ServiceRateLimitBucket) UnsetRateLimitSid()`

UnsetRateLimitSid ensures that no value is present for RateLimitSid, not even an explicit nil
### GetServiceSid

`func (o *VerifyV2ServiceRateLimitBucket) GetServiceSid() string`

GetServiceSid returns the ServiceSid field if non-nil, zero value otherwise.

### GetServiceSidOk

`func (o *VerifyV2ServiceRateLimitBucket) GetServiceSidOk() (*string, bool)`

GetServiceSidOk returns a tuple with the ServiceSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceSid

`func (o *VerifyV2ServiceRateLimitBucket) SetServiceSid(v string)`

SetServiceSid sets ServiceSid field to given value.

### HasServiceSid

`func (o *VerifyV2ServiceRateLimitBucket) HasServiceSid() bool`

HasServiceSid returns a boolean if a field has been set.

### SetServiceSidNil

`func (o *VerifyV2ServiceRateLimitBucket) SetServiceSidNil(b bool)`

 SetServiceSidNil sets the value for ServiceSid to be an explicit nil

### UnsetServiceSid
`func (o *VerifyV2ServiceRateLimitBucket) UnsetServiceSid()`

UnsetServiceSid ensures that no value is present for ServiceSid, not even an explicit nil
### GetSid

`func (o *VerifyV2ServiceRateLimitBucket) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *VerifyV2ServiceRateLimitBucket) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *VerifyV2ServiceRateLimitBucket) SetSid(v string)`

SetSid sets Sid field to given value.

### HasSid

`func (o *VerifyV2ServiceRateLimitBucket) HasSid() bool`

HasSid returns a boolean if a field has been set.

### SetSidNil

`func (o *VerifyV2ServiceRateLimitBucket) SetSidNil(b bool)`

 SetSidNil sets the value for Sid to be an explicit nil

### UnsetSid
`func (o *VerifyV2ServiceRateLimitBucket) UnsetSid()`

UnsetSid ensures that no value is present for Sid, not even an explicit nil
### GetUrl

`func (o *VerifyV2ServiceRateLimitBucket) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *VerifyV2ServiceRateLimitBucket) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *VerifyV2ServiceRateLimitBucket) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *VerifyV2ServiceRateLimitBucket) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *VerifyV2ServiceRateLimitBucket) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *VerifyV2ServiceRateLimitBucket) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


