# TrusthubV1CustomerProfile

## Properties

Name | Type | Description
------------ | ------------- | -------------
**AccountSid** | Pointer to **NullableString** | The SID of the Account that created the resource | [optional] 
**DateCreated** | Pointer to **NullableTime** | The ISO 8601 date and time in GMT when the resource was created | [optional] 
**DateUpdated** | Pointer to **NullableTime** | The ISO 8601 date and time in GMT when the resource was last updated | [optional] 
**Email** | Pointer to **NullableString** | The email address | [optional] 
**FriendlyName** | Pointer to **NullableString** | The string that you assigned to describe the resource | [optional] 
**Links** | Pointer to **map[string]interface{}** | The URLs of the Assigned Items of the Customer-Profile resource | [optional] 
**PolicySid** | Pointer to **NullableString** | The unique string of a policy. | [optional] 
**Sid** | Pointer to **NullableString** | The unique string that identifies the resource. | [optional] 
**Status** | Pointer to **NullableString** | The verification status of the Customer-Profile resource | [optional] 
**StatusCallback** | Pointer to **NullableString** | The URL we call to inform your application of status changes. | [optional] 
**Url** | Pointer to **NullableString** | The absolute URL of the Customer-Profile resource | [optional] 
**ValidUntil** | Pointer to **NullableTime** | The ISO 8601 date and time in GMT when the resource will be valid until. | [optional] 

## Methods

### NewTrusthubV1CustomerProfile

`func NewTrusthubV1CustomerProfile() *TrusthubV1CustomerProfile`

NewTrusthubV1CustomerProfile instantiates a new TrusthubV1CustomerProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrusthubV1CustomerProfileWithDefaults

`func NewTrusthubV1CustomerProfileWithDefaults() *TrusthubV1CustomerProfile`

NewTrusthubV1CustomerProfileWithDefaults instantiates a new TrusthubV1CustomerProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSid

`func (o *TrusthubV1CustomerProfile) GetAccountSid() string`

GetAccountSid returns the AccountSid field if non-nil, zero value otherwise.

### GetAccountSidOk

`func (o *TrusthubV1CustomerProfile) GetAccountSidOk() (*string, bool)`

GetAccountSidOk returns a tuple with the AccountSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSid

`func (o *TrusthubV1CustomerProfile) SetAccountSid(v string)`

SetAccountSid sets AccountSid field to given value.

### HasAccountSid

`func (o *TrusthubV1CustomerProfile) HasAccountSid() bool`

HasAccountSid returns a boolean if a field has been set.

### SetAccountSidNil

`func (o *TrusthubV1CustomerProfile) SetAccountSidNil(b bool)`

 SetAccountSidNil sets the value for AccountSid to be an explicit nil

### UnsetAccountSid
`func (o *TrusthubV1CustomerProfile) UnsetAccountSid()`

UnsetAccountSid ensures that no value is present for AccountSid, not even an explicit nil
### GetDateCreated

`func (o *TrusthubV1CustomerProfile) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *TrusthubV1CustomerProfile) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *TrusthubV1CustomerProfile) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *TrusthubV1CustomerProfile) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### SetDateCreatedNil

`func (o *TrusthubV1CustomerProfile) SetDateCreatedNil(b bool)`

 SetDateCreatedNil sets the value for DateCreated to be an explicit nil

### UnsetDateCreated
`func (o *TrusthubV1CustomerProfile) UnsetDateCreated()`

UnsetDateCreated ensures that no value is present for DateCreated, not even an explicit nil
### GetDateUpdated

`func (o *TrusthubV1CustomerProfile) GetDateUpdated() time.Time`

GetDateUpdated returns the DateUpdated field if non-nil, zero value otherwise.

### GetDateUpdatedOk

`func (o *TrusthubV1CustomerProfile) GetDateUpdatedOk() (*time.Time, bool)`

GetDateUpdatedOk returns a tuple with the DateUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateUpdated

`func (o *TrusthubV1CustomerProfile) SetDateUpdated(v time.Time)`

SetDateUpdated sets DateUpdated field to given value.

### HasDateUpdated

`func (o *TrusthubV1CustomerProfile) HasDateUpdated() bool`

HasDateUpdated returns a boolean if a field has been set.

### SetDateUpdatedNil

`func (o *TrusthubV1CustomerProfile) SetDateUpdatedNil(b bool)`

 SetDateUpdatedNil sets the value for DateUpdated to be an explicit nil

### UnsetDateUpdated
`func (o *TrusthubV1CustomerProfile) UnsetDateUpdated()`

UnsetDateUpdated ensures that no value is present for DateUpdated, not even an explicit nil
### GetEmail

`func (o *TrusthubV1CustomerProfile) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *TrusthubV1CustomerProfile) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *TrusthubV1CustomerProfile) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *TrusthubV1CustomerProfile) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *TrusthubV1CustomerProfile) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *TrusthubV1CustomerProfile) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetFriendlyName

`func (o *TrusthubV1CustomerProfile) GetFriendlyName() string`

GetFriendlyName returns the FriendlyName field if non-nil, zero value otherwise.

### GetFriendlyNameOk

`func (o *TrusthubV1CustomerProfile) GetFriendlyNameOk() (*string, bool)`

GetFriendlyNameOk returns a tuple with the FriendlyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFriendlyName

`func (o *TrusthubV1CustomerProfile) SetFriendlyName(v string)`

SetFriendlyName sets FriendlyName field to given value.

### HasFriendlyName

`func (o *TrusthubV1CustomerProfile) HasFriendlyName() bool`

HasFriendlyName returns a boolean if a field has been set.

### SetFriendlyNameNil

`func (o *TrusthubV1CustomerProfile) SetFriendlyNameNil(b bool)`

 SetFriendlyNameNil sets the value for FriendlyName to be an explicit nil

### UnsetFriendlyName
`func (o *TrusthubV1CustomerProfile) UnsetFriendlyName()`

UnsetFriendlyName ensures that no value is present for FriendlyName, not even an explicit nil
### GetLinks

`func (o *TrusthubV1CustomerProfile) GetLinks() map[string]interface{}`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *TrusthubV1CustomerProfile) GetLinksOk() (*map[string]interface{}, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *TrusthubV1CustomerProfile) SetLinks(v map[string]interface{})`

SetLinks sets Links field to given value.

### HasLinks

`func (o *TrusthubV1CustomerProfile) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### SetLinksNil

`func (o *TrusthubV1CustomerProfile) SetLinksNil(b bool)`

 SetLinksNil sets the value for Links to be an explicit nil

### UnsetLinks
`func (o *TrusthubV1CustomerProfile) UnsetLinks()`

UnsetLinks ensures that no value is present for Links, not even an explicit nil
### GetPolicySid

`func (o *TrusthubV1CustomerProfile) GetPolicySid() string`

GetPolicySid returns the PolicySid field if non-nil, zero value otherwise.

### GetPolicySidOk

`func (o *TrusthubV1CustomerProfile) GetPolicySidOk() (*string, bool)`

GetPolicySidOk returns a tuple with the PolicySid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicySid

`func (o *TrusthubV1CustomerProfile) SetPolicySid(v string)`

SetPolicySid sets PolicySid field to given value.

### HasPolicySid

`func (o *TrusthubV1CustomerProfile) HasPolicySid() bool`

HasPolicySid returns a boolean if a field has been set.

### SetPolicySidNil

`func (o *TrusthubV1CustomerProfile) SetPolicySidNil(b bool)`

 SetPolicySidNil sets the value for PolicySid to be an explicit nil

### UnsetPolicySid
`func (o *TrusthubV1CustomerProfile) UnsetPolicySid()`

UnsetPolicySid ensures that no value is present for PolicySid, not even an explicit nil
### GetSid

`func (o *TrusthubV1CustomerProfile) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *TrusthubV1CustomerProfile) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *TrusthubV1CustomerProfile) SetSid(v string)`

SetSid sets Sid field to given value.

### HasSid

`func (o *TrusthubV1CustomerProfile) HasSid() bool`

HasSid returns a boolean if a field has been set.

### SetSidNil

`func (o *TrusthubV1CustomerProfile) SetSidNil(b bool)`

 SetSidNil sets the value for Sid to be an explicit nil

### UnsetSid
`func (o *TrusthubV1CustomerProfile) UnsetSid()`

UnsetSid ensures that no value is present for Sid, not even an explicit nil
### GetStatus

`func (o *TrusthubV1CustomerProfile) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TrusthubV1CustomerProfile) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TrusthubV1CustomerProfile) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TrusthubV1CustomerProfile) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *TrusthubV1CustomerProfile) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *TrusthubV1CustomerProfile) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetStatusCallback

`func (o *TrusthubV1CustomerProfile) GetStatusCallback() string`

GetStatusCallback returns the StatusCallback field if non-nil, zero value otherwise.

### GetStatusCallbackOk

`func (o *TrusthubV1CustomerProfile) GetStatusCallbackOk() (*string, bool)`

GetStatusCallbackOk returns a tuple with the StatusCallback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCallback

`func (o *TrusthubV1CustomerProfile) SetStatusCallback(v string)`

SetStatusCallback sets StatusCallback field to given value.

### HasStatusCallback

`func (o *TrusthubV1CustomerProfile) HasStatusCallback() bool`

HasStatusCallback returns a boolean if a field has been set.

### SetStatusCallbackNil

`func (o *TrusthubV1CustomerProfile) SetStatusCallbackNil(b bool)`

 SetStatusCallbackNil sets the value for StatusCallback to be an explicit nil

### UnsetStatusCallback
`func (o *TrusthubV1CustomerProfile) UnsetStatusCallback()`

UnsetStatusCallback ensures that no value is present for StatusCallback, not even an explicit nil
### GetUrl

`func (o *TrusthubV1CustomerProfile) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *TrusthubV1CustomerProfile) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *TrusthubV1CustomerProfile) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *TrusthubV1CustomerProfile) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *TrusthubV1CustomerProfile) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *TrusthubV1CustomerProfile) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetValidUntil

`func (o *TrusthubV1CustomerProfile) GetValidUntil() time.Time`

GetValidUntil returns the ValidUntil field if non-nil, zero value otherwise.

### GetValidUntilOk

`func (o *TrusthubV1CustomerProfile) GetValidUntilOk() (*time.Time, bool)`

GetValidUntilOk returns a tuple with the ValidUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidUntil

`func (o *TrusthubV1CustomerProfile) SetValidUntil(v time.Time)`

SetValidUntil sets ValidUntil field to given value.

### HasValidUntil

`func (o *TrusthubV1CustomerProfile) HasValidUntil() bool`

HasValidUntil returns a boolean if a field has been set.

### SetValidUntilNil

`func (o *TrusthubV1CustomerProfile) SetValidUntilNil(b bool)`

 SetValidUntilNil sets the value for ValidUntil to be an explicit nil

### UnsetValidUntil
`func (o *TrusthubV1CustomerProfile) UnsetValidUntil()`

UnsetValidUntil ensures that no value is present for ValidUntil, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


