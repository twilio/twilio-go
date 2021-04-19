# TrusthubV1TrustProduct

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

### NewTrusthubV1TrustProduct

`func NewTrusthubV1TrustProduct() *TrusthubV1TrustProduct`

NewTrusthubV1TrustProduct instantiates a new TrusthubV1TrustProduct object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrusthubV1TrustProductWithDefaults

`func NewTrusthubV1TrustProductWithDefaults() *TrusthubV1TrustProduct`

NewTrusthubV1TrustProductWithDefaults instantiates a new TrusthubV1TrustProduct object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSid

`func (o *TrusthubV1TrustProduct) GetAccountSid() string`

GetAccountSid returns the AccountSid field if non-nil, zero value otherwise.

### GetAccountSidOk

`func (o *TrusthubV1TrustProduct) GetAccountSidOk() (*string, bool)`

GetAccountSidOk returns a tuple with the AccountSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSid

`func (o *TrusthubV1TrustProduct) SetAccountSid(v string)`

SetAccountSid sets AccountSid field to given value.

### HasAccountSid

`func (o *TrusthubV1TrustProduct) HasAccountSid() bool`

HasAccountSid returns a boolean if a field has been set.

### SetAccountSidNil

`func (o *TrusthubV1TrustProduct) SetAccountSidNil(b bool)`

 SetAccountSidNil sets the value for AccountSid to be an explicit nil

### UnsetAccountSid
`func (o *TrusthubV1TrustProduct) UnsetAccountSid()`

UnsetAccountSid ensures that no value is present for AccountSid, not even an explicit nil
### GetDateCreated

`func (o *TrusthubV1TrustProduct) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *TrusthubV1TrustProduct) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *TrusthubV1TrustProduct) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *TrusthubV1TrustProduct) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### SetDateCreatedNil

`func (o *TrusthubV1TrustProduct) SetDateCreatedNil(b bool)`

 SetDateCreatedNil sets the value for DateCreated to be an explicit nil

### UnsetDateCreated
`func (o *TrusthubV1TrustProduct) UnsetDateCreated()`

UnsetDateCreated ensures that no value is present for DateCreated, not even an explicit nil
### GetDateUpdated

`func (o *TrusthubV1TrustProduct) GetDateUpdated() time.Time`

GetDateUpdated returns the DateUpdated field if non-nil, zero value otherwise.

### GetDateUpdatedOk

`func (o *TrusthubV1TrustProduct) GetDateUpdatedOk() (*time.Time, bool)`

GetDateUpdatedOk returns a tuple with the DateUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateUpdated

`func (o *TrusthubV1TrustProduct) SetDateUpdated(v time.Time)`

SetDateUpdated sets DateUpdated field to given value.

### HasDateUpdated

`func (o *TrusthubV1TrustProduct) HasDateUpdated() bool`

HasDateUpdated returns a boolean if a field has been set.

### SetDateUpdatedNil

`func (o *TrusthubV1TrustProduct) SetDateUpdatedNil(b bool)`

 SetDateUpdatedNil sets the value for DateUpdated to be an explicit nil

### UnsetDateUpdated
`func (o *TrusthubV1TrustProduct) UnsetDateUpdated()`

UnsetDateUpdated ensures that no value is present for DateUpdated, not even an explicit nil
### GetEmail

`func (o *TrusthubV1TrustProduct) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *TrusthubV1TrustProduct) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *TrusthubV1TrustProduct) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *TrusthubV1TrustProduct) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *TrusthubV1TrustProduct) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *TrusthubV1TrustProduct) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetFriendlyName

`func (o *TrusthubV1TrustProduct) GetFriendlyName() string`

GetFriendlyName returns the FriendlyName field if non-nil, zero value otherwise.

### GetFriendlyNameOk

`func (o *TrusthubV1TrustProduct) GetFriendlyNameOk() (*string, bool)`

GetFriendlyNameOk returns a tuple with the FriendlyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFriendlyName

`func (o *TrusthubV1TrustProduct) SetFriendlyName(v string)`

SetFriendlyName sets FriendlyName field to given value.

### HasFriendlyName

`func (o *TrusthubV1TrustProduct) HasFriendlyName() bool`

HasFriendlyName returns a boolean if a field has been set.

### SetFriendlyNameNil

`func (o *TrusthubV1TrustProduct) SetFriendlyNameNil(b bool)`

 SetFriendlyNameNil sets the value for FriendlyName to be an explicit nil

### UnsetFriendlyName
`func (o *TrusthubV1TrustProduct) UnsetFriendlyName()`

UnsetFriendlyName ensures that no value is present for FriendlyName, not even an explicit nil
### GetLinks

`func (o *TrusthubV1TrustProduct) GetLinks() map[string]interface{}`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *TrusthubV1TrustProduct) GetLinksOk() (*map[string]interface{}, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *TrusthubV1TrustProduct) SetLinks(v map[string]interface{})`

SetLinks sets Links field to given value.

### HasLinks

`func (o *TrusthubV1TrustProduct) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### SetLinksNil

`func (o *TrusthubV1TrustProduct) SetLinksNil(b bool)`

 SetLinksNil sets the value for Links to be an explicit nil

### UnsetLinks
`func (o *TrusthubV1TrustProduct) UnsetLinks()`

UnsetLinks ensures that no value is present for Links, not even an explicit nil
### GetPolicySid

`func (o *TrusthubV1TrustProduct) GetPolicySid() string`

GetPolicySid returns the PolicySid field if non-nil, zero value otherwise.

### GetPolicySidOk

`func (o *TrusthubV1TrustProduct) GetPolicySidOk() (*string, bool)`

GetPolicySidOk returns a tuple with the PolicySid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicySid

`func (o *TrusthubV1TrustProduct) SetPolicySid(v string)`

SetPolicySid sets PolicySid field to given value.

### HasPolicySid

`func (o *TrusthubV1TrustProduct) HasPolicySid() bool`

HasPolicySid returns a boolean if a field has been set.

### SetPolicySidNil

`func (o *TrusthubV1TrustProduct) SetPolicySidNil(b bool)`

 SetPolicySidNil sets the value for PolicySid to be an explicit nil

### UnsetPolicySid
`func (o *TrusthubV1TrustProduct) UnsetPolicySid()`

UnsetPolicySid ensures that no value is present for PolicySid, not even an explicit nil
### GetSid

`func (o *TrusthubV1TrustProduct) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *TrusthubV1TrustProduct) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *TrusthubV1TrustProduct) SetSid(v string)`

SetSid sets Sid field to given value.

### HasSid

`func (o *TrusthubV1TrustProduct) HasSid() bool`

HasSid returns a boolean if a field has been set.

### SetSidNil

`func (o *TrusthubV1TrustProduct) SetSidNil(b bool)`

 SetSidNil sets the value for Sid to be an explicit nil

### UnsetSid
`func (o *TrusthubV1TrustProduct) UnsetSid()`

UnsetSid ensures that no value is present for Sid, not even an explicit nil
### GetStatus

`func (o *TrusthubV1TrustProduct) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TrusthubV1TrustProduct) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TrusthubV1TrustProduct) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TrusthubV1TrustProduct) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *TrusthubV1TrustProduct) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *TrusthubV1TrustProduct) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetStatusCallback

`func (o *TrusthubV1TrustProduct) GetStatusCallback() string`

GetStatusCallback returns the StatusCallback field if non-nil, zero value otherwise.

### GetStatusCallbackOk

`func (o *TrusthubV1TrustProduct) GetStatusCallbackOk() (*string, bool)`

GetStatusCallbackOk returns a tuple with the StatusCallback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCallback

`func (o *TrusthubV1TrustProduct) SetStatusCallback(v string)`

SetStatusCallback sets StatusCallback field to given value.

### HasStatusCallback

`func (o *TrusthubV1TrustProduct) HasStatusCallback() bool`

HasStatusCallback returns a boolean if a field has been set.

### SetStatusCallbackNil

`func (o *TrusthubV1TrustProduct) SetStatusCallbackNil(b bool)`

 SetStatusCallbackNil sets the value for StatusCallback to be an explicit nil

### UnsetStatusCallback
`func (o *TrusthubV1TrustProduct) UnsetStatusCallback()`

UnsetStatusCallback ensures that no value is present for StatusCallback, not even an explicit nil
### GetUrl

`func (o *TrusthubV1TrustProduct) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *TrusthubV1TrustProduct) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *TrusthubV1TrustProduct) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *TrusthubV1TrustProduct) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *TrusthubV1TrustProduct) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *TrusthubV1TrustProduct) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetValidUntil

`func (o *TrusthubV1TrustProduct) GetValidUntil() time.Time`

GetValidUntil returns the ValidUntil field if non-nil, zero value otherwise.

### GetValidUntilOk

`func (o *TrusthubV1TrustProduct) GetValidUntilOk() (*time.Time, bool)`

GetValidUntilOk returns a tuple with the ValidUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidUntil

`func (o *TrusthubV1TrustProduct) SetValidUntil(v time.Time)`

SetValidUntil sets ValidUntil field to given value.

### HasValidUntil

`func (o *TrusthubV1TrustProduct) HasValidUntil() bool`

HasValidUntil returns a boolean if a field has been set.

### SetValidUntilNil

`func (o *TrusthubV1TrustProduct) SetValidUntilNil(b bool)`

 SetValidUntilNil sets the value for ValidUntil to be an explicit nil

### UnsetValidUntil
`func (o *TrusthubV1TrustProduct) UnsetValidUntil()`

UnsetValidUntil ensures that no value is present for ValidUntil, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


