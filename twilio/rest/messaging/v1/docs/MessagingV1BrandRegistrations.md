# MessagingV1BrandRegistrations

## Properties

Name | Type | Description
------------ | ------------- | -------------
**A2pProfileBundleSid** | Pointer to **NullableString** | A2P Messaging Profile Bundle BundleSid | [optional] 
**AccountSid** | Pointer to **NullableString** | The SID of the Account that created the resource | [optional] 
**CustomerProfileBundleSid** | Pointer to **NullableString** | A2P Messaging Profile Bundle BundleSid | [optional] 
**DateCreated** | Pointer to **NullableTime** | The ISO 8601 date and time in GMT when the resource was created | [optional] 
**DateUpdated** | Pointer to **NullableTime** | The ISO 8601 date and time in GMT when the resource was last updated | [optional] 
**FailureReason** | Pointer to **NullableString** | A reason why brand registration has failed | [optional] 
**Sid** | Pointer to **NullableString** | A2P BrandRegistration Sid | [optional] 
**Status** | Pointer to **NullableString** | Brand Registration status | [optional] 
**TcrId** | Pointer to **NullableString** | Campaign Registry (TCR) Brand ID | [optional] 
**Url** | Pointer to **NullableString** | The absolute URL of the Brand Registration | [optional] 

## Methods

### NewMessagingV1BrandRegistrations

`func NewMessagingV1BrandRegistrations() *MessagingV1BrandRegistrations`

NewMessagingV1BrandRegistrations instantiates a new MessagingV1BrandRegistrations object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMessagingV1BrandRegistrationsWithDefaults

`func NewMessagingV1BrandRegistrationsWithDefaults() *MessagingV1BrandRegistrations`

NewMessagingV1BrandRegistrationsWithDefaults instantiates a new MessagingV1BrandRegistrations object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetA2pProfileBundleSid

`func (o *MessagingV1BrandRegistrations) GetA2pProfileBundleSid() string`

GetA2pProfileBundleSid returns the A2pProfileBundleSid field if non-nil, zero value otherwise.

### GetA2pProfileBundleSidOk

`func (o *MessagingV1BrandRegistrations) GetA2pProfileBundleSidOk() (*string, bool)`

GetA2pProfileBundleSidOk returns a tuple with the A2pProfileBundleSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetA2pProfileBundleSid

`func (o *MessagingV1BrandRegistrations) SetA2pProfileBundleSid(v string)`

SetA2pProfileBundleSid sets A2pProfileBundleSid field to given value.

### HasA2pProfileBundleSid

`func (o *MessagingV1BrandRegistrations) HasA2pProfileBundleSid() bool`

HasA2pProfileBundleSid returns a boolean if a field has been set.

### SetA2pProfileBundleSidNil

`func (o *MessagingV1BrandRegistrations) SetA2pProfileBundleSidNil(b bool)`

 SetA2pProfileBundleSidNil sets the value for A2pProfileBundleSid to be an explicit nil

### UnsetA2pProfileBundleSid
`func (o *MessagingV1BrandRegistrations) UnsetA2pProfileBundleSid()`

UnsetA2pProfileBundleSid ensures that no value is present for A2pProfileBundleSid, not even an explicit nil
### GetAccountSid

`func (o *MessagingV1BrandRegistrations) GetAccountSid() string`

GetAccountSid returns the AccountSid field if non-nil, zero value otherwise.

### GetAccountSidOk

`func (o *MessagingV1BrandRegistrations) GetAccountSidOk() (*string, bool)`

GetAccountSidOk returns a tuple with the AccountSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSid

`func (o *MessagingV1BrandRegistrations) SetAccountSid(v string)`

SetAccountSid sets AccountSid field to given value.

### HasAccountSid

`func (o *MessagingV1BrandRegistrations) HasAccountSid() bool`

HasAccountSid returns a boolean if a field has been set.

### SetAccountSidNil

`func (o *MessagingV1BrandRegistrations) SetAccountSidNil(b bool)`

 SetAccountSidNil sets the value for AccountSid to be an explicit nil

### UnsetAccountSid
`func (o *MessagingV1BrandRegistrations) UnsetAccountSid()`

UnsetAccountSid ensures that no value is present for AccountSid, not even an explicit nil
### GetCustomerProfileBundleSid

`func (o *MessagingV1BrandRegistrations) GetCustomerProfileBundleSid() string`

GetCustomerProfileBundleSid returns the CustomerProfileBundleSid field if non-nil, zero value otherwise.

### GetCustomerProfileBundleSidOk

`func (o *MessagingV1BrandRegistrations) GetCustomerProfileBundleSidOk() (*string, bool)`

GetCustomerProfileBundleSidOk returns a tuple with the CustomerProfileBundleSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerProfileBundleSid

`func (o *MessagingV1BrandRegistrations) SetCustomerProfileBundleSid(v string)`

SetCustomerProfileBundleSid sets CustomerProfileBundleSid field to given value.

### HasCustomerProfileBundleSid

`func (o *MessagingV1BrandRegistrations) HasCustomerProfileBundleSid() bool`

HasCustomerProfileBundleSid returns a boolean if a field has been set.

### SetCustomerProfileBundleSidNil

`func (o *MessagingV1BrandRegistrations) SetCustomerProfileBundleSidNil(b bool)`

 SetCustomerProfileBundleSidNil sets the value for CustomerProfileBundleSid to be an explicit nil

### UnsetCustomerProfileBundleSid
`func (o *MessagingV1BrandRegistrations) UnsetCustomerProfileBundleSid()`

UnsetCustomerProfileBundleSid ensures that no value is present for CustomerProfileBundleSid, not even an explicit nil
### GetDateCreated

`func (o *MessagingV1BrandRegistrations) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *MessagingV1BrandRegistrations) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *MessagingV1BrandRegistrations) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *MessagingV1BrandRegistrations) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### SetDateCreatedNil

`func (o *MessagingV1BrandRegistrations) SetDateCreatedNil(b bool)`

 SetDateCreatedNil sets the value for DateCreated to be an explicit nil

### UnsetDateCreated
`func (o *MessagingV1BrandRegistrations) UnsetDateCreated()`

UnsetDateCreated ensures that no value is present for DateCreated, not even an explicit nil
### GetDateUpdated

`func (o *MessagingV1BrandRegistrations) GetDateUpdated() time.Time`

GetDateUpdated returns the DateUpdated field if non-nil, zero value otherwise.

### GetDateUpdatedOk

`func (o *MessagingV1BrandRegistrations) GetDateUpdatedOk() (*time.Time, bool)`

GetDateUpdatedOk returns a tuple with the DateUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateUpdated

`func (o *MessagingV1BrandRegistrations) SetDateUpdated(v time.Time)`

SetDateUpdated sets DateUpdated field to given value.

### HasDateUpdated

`func (o *MessagingV1BrandRegistrations) HasDateUpdated() bool`

HasDateUpdated returns a boolean if a field has been set.

### SetDateUpdatedNil

`func (o *MessagingV1BrandRegistrations) SetDateUpdatedNil(b bool)`

 SetDateUpdatedNil sets the value for DateUpdated to be an explicit nil

### UnsetDateUpdated
`func (o *MessagingV1BrandRegistrations) UnsetDateUpdated()`

UnsetDateUpdated ensures that no value is present for DateUpdated, not even an explicit nil
### GetFailureReason

`func (o *MessagingV1BrandRegistrations) GetFailureReason() string`

GetFailureReason returns the FailureReason field if non-nil, zero value otherwise.

### GetFailureReasonOk

`func (o *MessagingV1BrandRegistrations) GetFailureReasonOk() (*string, bool)`

GetFailureReasonOk returns a tuple with the FailureReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureReason

`func (o *MessagingV1BrandRegistrations) SetFailureReason(v string)`

SetFailureReason sets FailureReason field to given value.

### HasFailureReason

`func (o *MessagingV1BrandRegistrations) HasFailureReason() bool`

HasFailureReason returns a boolean if a field has been set.

### SetFailureReasonNil

`func (o *MessagingV1BrandRegistrations) SetFailureReasonNil(b bool)`

 SetFailureReasonNil sets the value for FailureReason to be an explicit nil

### UnsetFailureReason
`func (o *MessagingV1BrandRegistrations) UnsetFailureReason()`

UnsetFailureReason ensures that no value is present for FailureReason, not even an explicit nil
### GetSid

`func (o *MessagingV1BrandRegistrations) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *MessagingV1BrandRegistrations) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *MessagingV1BrandRegistrations) SetSid(v string)`

SetSid sets Sid field to given value.

### HasSid

`func (o *MessagingV1BrandRegistrations) HasSid() bool`

HasSid returns a boolean if a field has been set.

### SetSidNil

`func (o *MessagingV1BrandRegistrations) SetSidNil(b bool)`

 SetSidNil sets the value for Sid to be an explicit nil

### UnsetSid
`func (o *MessagingV1BrandRegistrations) UnsetSid()`

UnsetSid ensures that no value is present for Sid, not even an explicit nil
### GetStatus

`func (o *MessagingV1BrandRegistrations) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MessagingV1BrandRegistrations) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MessagingV1BrandRegistrations) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MessagingV1BrandRegistrations) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *MessagingV1BrandRegistrations) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *MessagingV1BrandRegistrations) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetTcrId

`func (o *MessagingV1BrandRegistrations) GetTcrId() string`

GetTcrId returns the TcrId field if non-nil, zero value otherwise.

### GetTcrIdOk

`func (o *MessagingV1BrandRegistrations) GetTcrIdOk() (*string, bool)`

GetTcrIdOk returns a tuple with the TcrId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcrId

`func (o *MessagingV1BrandRegistrations) SetTcrId(v string)`

SetTcrId sets TcrId field to given value.

### HasTcrId

`func (o *MessagingV1BrandRegistrations) HasTcrId() bool`

HasTcrId returns a boolean if a field has been set.

### SetTcrIdNil

`func (o *MessagingV1BrandRegistrations) SetTcrIdNil(b bool)`

 SetTcrIdNil sets the value for TcrId to be an explicit nil

### UnsetTcrId
`func (o *MessagingV1BrandRegistrations) UnsetTcrId()`

UnsetTcrId ensures that no value is present for TcrId, not even an explicit nil
### GetUrl

`func (o *MessagingV1BrandRegistrations) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *MessagingV1BrandRegistrations) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *MessagingV1BrandRegistrations) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *MessagingV1BrandRegistrations) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *MessagingV1BrandRegistrations) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *MessagingV1BrandRegistrations) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


