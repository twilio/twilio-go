# ApiV2010AccountQueue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **NullableString** | The SID of the Account that created this resource | [optional] 
**AverageWaitTime** | Pointer to **NullableInt32** | Average wait time of members in the queue | [optional] 
**CurrentSize** | Pointer to **NullableInt32** | The number of calls currently in the queue. | [optional] 
**DateCreated** | Pointer to **NullableString** | The RFC 2822 date and time in GMT that this resource was created | [optional] 
**DateUpdated** | Pointer to **NullableString** | The RFC 2822 date and time in GMT that this resource was last updated | [optional] 
**FriendlyName** | Pointer to **NullableString** | A string that you assigned to describe this resource | [optional] 
**MaxSize** | Pointer to **NullableInt32** | The max number of calls allowed in the queue | [optional] 
**Sid** | Pointer to **NullableString** | The unique string that identifies this resource | [optional] 
**Uri** | Pointer to **NullableString** | The URI of this resource, relative to &#x60;https://api.twilio.com&#x60; | [optional] 

## Methods

### NewApiV2010AccountQueue

`func NewApiV2010AccountQueue() *ApiV2010AccountQueue`

NewApiV2010AccountQueue instantiates a new ApiV2010AccountQueue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiV2010AccountQueueWithDefaults

`func NewApiV2010AccountQueueWithDefaults() *ApiV2010AccountQueue`

NewApiV2010AccountQueueWithDefaults instantiates a new ApiV2010AccountQueue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSid

`func (o *ApiV2010AccountQueue) GetAccountSid() string`

GetAccountSid returns the AccountSid field if non-nil, zero value otherwise.

### GetAccountSidOk

`func (o *ApiV2010AccountQueue) GetAccountSidOk() (*string, bool)`

GetAccountSidOk returns a tuple with the AccountSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSid

`func (o *ApiV2010AccountQueue) SetAccountSid(v string)`

SetAccountSid sets AccountSid field to given value.

### HasAccountSid

`func (o *ApiV2010AccountQueue) HasAccountSid() bool`

HasAccountSid returns a boolean if a field has been set.

### SetAccountSidNil

`func (o *ApiV2010AccountQueue) SetAccountSidNil(b bool)`

 SetAccountSidNil sets the value for AccountSid to be an explicit nil

### UnsetAccountSid
`func (o *ApiV2010AccountQueue) UnsetAccountSid()`

UnsetAccountSid ensures that no value is present for AccountSid, not even an explicit nil
### GetAverageWaitTime

`func (o *ApiV2010AccountQueue) GetAverageWaitTime() int32`

GetAverageWaitTime returns the AverageWaitTime field if non-nil, zero value otherwise.

### GetAverageWaitTimeOk

`func (o *ApiV2010AccountQueue) GetAverageWaitTimeOk() (*int32, bool)`

GetAverageWaitTimeOk returns a tuple with the AverageWaitTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageWaitTime

`func (o *ApiV2010AccountQueue) SetAverageWaitTime(v int32)`

SetAverageWaitTime sets AverageWaitTime field to given value.

### HasAverageWaitTime

`func (o *ApiV2010AccountQueue) HasAverageWaitTime() bool`

HasAverageWaitTime returns a boolean if a field has been set.

### SetAverageWaitTimeNil

`func (o *ApiV2010AccountQueue) SetAverageWaitTimeNil(b bool)`

 SetAverageWaitTimeNil sets the value for AverageWaitTime to be an explicit nil

### UnsetAverageWaitTime
`func (o *ApiV2010AccountQueue) UnsetAverageWaitTime()`

UnsetAverageWaitTime ensures that no value is present for AverageWaitTime, not even an explicit nil
### GetCurrentSize

`func (o *ApiV2010AccountQueue) GetCurrentSize() int32`

GetCurrentSize returns the CurrentSize field if non-nil, zero value otherwise.

### GetCurrentSizeOk

`func (o *ApiV2010AccountQueue) GetCurrentSizeOk() (*int32, bool)`

GetCurrentSizeOk returns a tuple with the CurrentSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentSize

`func (o *ApiV2010AccountQueue) SetCurrentSize(v int32)`

SetCurrentSize sets CurrentSize field to given value.

### HasCurrentSize

`func (o *ApiV2010AccountQueue) HasCurrentSize() bool`

HasCurrentSize returns a boolean if a field has been set.

### SetCurrentSizeNil

`func (o *ApiV2010AccountQueue) SetCurrentSizeNil(b bool)`

 SetCurrentSizeNil sets the value for CurrentSize to be an explicit nil

### UnsetCurrentSize
`func (o *ApiV2010AccountQueue) UnsetCurrentSize()`

UnsetCurrentSize ensures that no value is present for CurrentSize, not even an explicit nil
### GetDateCreated

`func (o *ApiV2010AccountQueue) GetDateCreated() string`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *ApiV2010AccountQueue) GetDateCreatedOk() (*string, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *ApiV2010AccountQueue) SetDateCreated(v string)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *ApiV2010AccountQueue) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### SetDateCreatedNil

`func (o *ApiV2010AccountQueue) SetDateCreatedNil(b bool)`

 SetDateCreatedNil sets the value for DateCreated to be an explicit nil

### UnsetDateCreated
`func (o *ApiV2010AccountQueue) UnsetDateCreated()`

UnsetDateCreated ensures that no value is present for DateCreated, not even an explicit nil
### GetDateUpdated

`func (o *ApiV2010AccountQueue) GetDateUpdated() string`

GetDateUpdated returns the DateUpdated field if non-nil, zero value otherwise.

### GetDateUpdatedOk

`func (o *ApiV2010AccountQueue) GetDateUpdatedOk() (*string, bool)`

GetDateUpdatedOk returns a tuple with the DateUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateUpdated

`func (o *ApiV2010AccountQueue) SetDateUpdated(v string)`

SetDateUpdated sets DateUpdated field to given value.

### HasDateUpdated

`func (o *ApiV2010AccountQueue) HasDateUpdated() bool`

HasDateUpdated returns a boolean if a field has been set.

### SetDateUpdatedNil

`func (o *ApiV2010AccountQueue) SetDateUpdatedNil(b bool)`

 SetDateUpdatedNil sets the value for DateUpdated to be an explicit nil

### UnsetDateUpdated
`func (o *ApiV2010AccountQueue) UnsetDateUpdated()`

UnsetDateUpdated ensures that no value is present for DateUpdated, not even an explicit nil
### GetFriendlyName

`func (o *ApiV2010AccountQueue) GetFriendlyName() string`

GetFriendlyName returns the FriendlyName field if non-nil, zero value otherwise.

### GetFriendlyNameOk

`func (o *ApiV2010AccountQueue) GetFriendlyNameOk() (*string, bool)`

GetFriendlyNameOk returns a tuple with the FriendlyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFriendlyName

`func (o *ApiV2010AccountQueue) SetFriendlyName(v string)`

SetFriendlyName sets FriendlyName field to given value.

### HasFriendlyName

`func (o *ApiV2010AccountQueue) HasFriendlyName() bool`

HasFriendlyName returns a boolean if a field has been set.

### SetFriendlyNameNil

`func (o *ApiV2010AccountQueue) SetFriendlyNameNil(b bool)`

 SetFriendlyNameNil sets the value for FriendlyName to be an explicit nil

### UnsetFriendlyName
`func (o *ApiV2010AccountQueue) UnsetFriendlyName()`

UnsetFriendlyName ensures that no value is present for FriendlyName, not even an explicit nil
### GetMaxSize

`func (o *ApiV2010AccountQueue) GetMaxSize() int32`

GetMaxSize returns the MaxSize field if non-nil, zero value otherwise.

### GetMaxSizeOk

`func (o *ApiV2010AccountQueue) GetMaxSizeOk() (*int32, bool)`

GetMaxSizeOk returns a tuple with the MaxSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSize

`func (o *ApiV2010AccountQueue) SetMaxSize(v int32)`

SetMaxSize sets MaxSize field to given value.

### HasMaxSize

`func (o *ApiV2010AccountQueue) HasMaxSize() bool`

HasMaxSize returns a boolean if a field has been set.

### SetMaxSizeNil

`func (o *ApiV2010AccountQueue) SetMaxSizeNil(b bool)`

 SetMaxSizeNil sets the value for MaxSize to be an explicit nil

### UnsetMaxSize
`func (o *ApiV2010AccountQueue) UnsetMaxSize()`

UnsetMaxSize ensures that no value is present for MaxSize, not even an explicit nil
### GetSid

`func (o *ApiV2010AccountQueue) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *ApiV2010AccountQueue) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *ApiV2010AccountQueue) SetSid(v string)`

SetSid sets Sid field to given value.

### HasSid

`func (o *ApiV2010AccountQueue) HasSid() bool`

HasSid returns a boolean if a field has been set.

### SetSidNil

`func (o *ApiV2010AccountQueue) SetSidNil(b bool)`

 SetSidNil sets the value for Sid to be an explicit nil

### UnsetSid
`func (o *ApiV2010AccountQueue) UnsetSid()`

UnsetSid ensures that no value is present for Sid, not even an explicit nil
### GetUri

`func (o *ApiV2010AccountQueue) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *ApiV2010AccountQueue) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *ApiV2010AccountQueue) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *ApiV2010AccountQueue) HasUri() bool`

HasUri returns a boolean if a field has been set.

### SetUriNil

`func (o *ApiV2010AccountQueue) SetUriNil(b bool)`

 SetUriNil sets the value for Uri to be an explicit nil

### UnsetUri
`func (o *ApiV2010AccountQueue) UnsetUri()`

UnsetUri ensures that no value is present for Uri, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


