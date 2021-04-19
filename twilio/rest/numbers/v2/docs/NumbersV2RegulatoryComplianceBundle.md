# NumbersV2RegulatoryComplianceBundle

## Properties

Name | Type | Description
------------ | ------------- | -------------
**AccountSid** | Pointer to **NullableString** | The SID of the Account that created the resource | [optional] 
**DateCreated** | Pointer to **NullableTime** | The ISO 8601 date and time in GMT when the resource was created | [optional] 
**DateUpdated** | Pointer to **NullableTime** | The ISO 8601 date and time in GMT when the resource was last updated | [optional] 
**Email** | Pointer to **NullableString** | The email address | [optional] 
**FriendlyName** | Pointer to **NullableString** | The string that you assigned to describe the resource | [optional] 
**Links** | Pointer to **map[string]interface{}** | The URLs of the Assigned Items of the Bundle resource | [optional] 
**RegulationSid** | Pointer to **NullableString** | The unique string of a regulation. | [optional] 
**Sid** | Pointer to **NullableString** | The unique string that identifies the resource. | [optional] 
**Status** | Pointer to **NullableString** | The verification status of the Bundle resource | [optional] 
**StatusCallback** | Pointer to **NullableString** | The URL we call to inform your application of status changes. | [optional] 
**Url** | Pointer to **NullableString** | The absolute URL of the Bundle resource | [optional] 
**ValidUntil** | Pointer to **NullableTime** | The ISO 8601 date and time in GMT when the resource will be valid until. | [optional] 

## Methods

### NewNumbersV2RegulatoryComplianceBundle

`func NewNumbersV2RegulatoryComplianceBundle() *NumbersV2RegulatoryComplianceBundle`

NewNumbersV2RegulatoryComplianceBundle instantiates a new NumbersV2RegulatoryComplianceBundle object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNumbersV2RegulatoryComplianceBundleWithDefaults

`func NewNumbersV2RegulatoryComplianceBundleWithDefaults() *NumbersV2RegulatoryComplianceBundle`

NewNumbersV2RegulatoryComplianceBundleWithDefaults instantiates a new NumbersV2RegulatoryComplianceBundle object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSid

`func (o *NumbersV2RegulatoryComplianceBundle) GetAccountSid() string`

GetAccountSid returns the AccountSid field if non-nil, zero value otherwise.

### GetAccountSidOk

`func (o *NumbersV2RegulatoryComplianceBundle) GetAccountSidOk() (*string, bool)`

GetAccountSidOk returns a tuple with the AccountSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSid

`func (o *NumbersV2RegulatoryComplianceBundle) SetAccountSid(v string)`

SetAccountSid sets AccountSid field to given value.

### HasAccountSid

`func (o *NumbersV2RegulatoryComplianceBundle) HasAccountSid() bool`

HasAccountSid returns a boolean if a field has been set.

### SetAccountSidNil

`func (o *NumbersV2RegulatoryComplianceBundle) SetAccountSidNil(b bool)`

 SetAccountSidNil sets the value for AccountSid to be an explicit nil

### UnsetAccountSid
`func (o *NumbersV2RegulatoryComplianceBundle) UnsetAccountSid()`

UnsetAccountSid ensures that no value is present for AccountSid, not even an explicit nil
### GetDateCreated

`func (o *NumbersV2RegulatoryComplianceBundle) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *NumbersV2RegulatoryComplianceBundle) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *NumbersV2RegulatoryComplianceBundle) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *NumbersV2RegulatoryComplianceBundle) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### SetDateCreatedNil

`func (o *NumbersV2RegulatoryComplianceBundle) SetDateCreatedNil(b bool)`

 SetDateCreatedNil sets the value for DateCreated to be an explicit nil

### UnsetDateCreated
`func (o *NumbersV2RegulatoryComplianceBundle) UnsetDateCreated()`

UnsetDateCreated ensures that no value is present for DateCreated, not even an explicit nil
### GetDateUpdated

`func (o *NumbersV2RegulatoryComplianceBundle) GetDateUpdated() time.Time`

GetDateUpdated returns the DateUpdated field if non-nil, zero value otherwise.

### GetDateUpdatedOk

`func (o *NumbersV2RegulatoryComplianceBundle) GetDateUpdatedOk() (*time.Time, bool)`

GetDateUpdatedOk returns a tuple with the DateUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateUpdated

`func (o *NumbersV2RegulatoryComplianceBundle) SetDateUpdated(v time.Time)`

SetDateUpdated sets DateUpdated field to given value.

### HasDateUpdated

`func (o *NumbersV2RegulatoryComplianceBundle) HasDateUpdated() bool`

HasDateUpdated returns a boolean if a field has been set.

### SetDateUpdatedNil

`func (o *NumbersV2RegulatoryComplianceBundle) SetDateUpdatedNil(b bool)`

 SetDateUpdatedNil sets the value for DateUpdated to be an explicit nil

### UnsetDateUpdated
`func (o *NumbersV2RegulatoryComplianceBundle) UnsetDateUpdated()`

UnsetDateUpdated ensures that no value is present for DateUpdated, not even an explicit nil
### GetEmail

`func (o *NumbersV2RegulatoryComplianceBundle) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *NumbersV2RegulatoryComplianceBundle) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *NumbersV2RegulatoryComplianceBundle) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *NumbersV2RegulatoryComplianceBundle) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *NumbersV2RegulatoryComplianceBundle) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *NumbersV2RegulatoryComplianceBundle) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetFriendlyName

`func (o *NumbersV2RegulatoryComplianceBundle) GetFriendlyName() string`

GetFriendlyName returns the FriendlyName field if non-nil, zero value otherwise.

### GetFriendlyNameOk

`func (o *NumbersV2RegulatoryComplianceBundle) GetFriendlyNameOk() (*string, bool)`

GetFriendlyNameOk returns a tuple with the FriendlyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFriendlyName

`func (o *NumbersV2RegulatoryComplianceBundle) SetFriendlyName(v string)`

SetFriendlyName sets FriendlyName field to given value.

### HasFriendlyName

`func (o *NumbersV2RegulatoryComplianceBundle) HasFriendlyName() bool`

HasFriendlyName returns a boolean if a field has been set.

### SetFriendlyNameNil

`func (o *NumbersV2RegulatoryComplianceBundle) SetFriendlyNameNil(b bool)`

 SetFriendlyNameNil sets the value for FriendlyName to be an explicit nil

### UnsetFriendlyName
`func (o *NumbersV2RegulatoryComplianceBundle) UnsetFriendlyName()`

UnsetFriendlyName ensures that no value is present for FriendlyName, not even an explicit nil
### GetLinks

`func (o *NumbersV2RegulatoryComplianceBundle) GetLinks() map[string]interface{}`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *NumbersV2RegulatoryComplianceBundle) GetLinksOk() (*map[string]interface{}, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *NumbersV2RegulatoryComplianceBundle) SetLinks(v map[string]interface{})`

SetLinks sets Links field to given value.

### HasLinks

`func (o *NumbersV2RegulatoryComplianceBundle) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### SetLinksNil

`func (o *NumbersV2RegulatoryComplianceBundle) SetLinksNil(b bool)`

 SetLinksNil sets the value for Links to be an explicit nil

### UnsetLinks
`func (o *NumbersV2RegulatoryComplianceBundle) UnsetLinks()`

UnsetLinks ensures that no value is present for Links, not even an explicit nil
### GetRegulationSid

`func (o *NumbersV2RegulatoryComplianceBundle) GetRegulationSid() string`

GetRegulationSid returns the RegulationSid field if non-nil, zero value otherwise.

### GetRegulationSidOk

`func (o *NumbersV2RegulatoryComplianceBundle) GetRegulationSidOk() (*string, bool)`

GetRegulationSidOk returns a tuple with the RegulationSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegulationSid

`func (o *NumbersV2RegulatoryComplianceBundle) SetRegulationSid(v string)`

SetRegulationSid sets RegulationSid field to given value.

### HasRegulationSid

`func (o *NumbersV2RegulatoryComplianceBundle) HasRegulationSid() bool`

HasRegulationSid returns a boolean if a field has been set.

### SetRegulationSidNil

`func (o *NumbersV2RegulatoryComplianceBundle) SetRegulationSidNil(b bool)`

 SetRegulationSidNil sets the value for RegulationSid to be an explicit nil

### UnsetRegulationSid
`func (o *NumbersV2RegulatoryComplianceBundle) UnsetRegulationSid()`

UnsetRegulationSid ensures that no value is present for RegulationSid, not even an explicit nil
### GetSid

`func (o *NumbersV2RegulatoryComplianceBundle) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *NumbersV2RegulatoryComplianceBundle) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *NumbersV2RegulatoryComplianceBundle) SetSid(v string)`

SetSid sets Sid field to given value.

### HasSid

`func (o *NumbersV2RegulatoryComplianceBundle) HasSid() bool`

HasSid returns a boolean if a field has been set.

### SetSidNil

`func (o *NumbersV2RegulatoryComplianceBundle) SetSidNil(b bool)`

 SetSidNil sets the value for Sid to be an explicit nil

### UnsetSid
`func (o *NumbersV2RegulatoryComplianceBundle) UnsetSid()`

UnsetSid ensures that no value is present for Sid, not even an explicit nil
### GetStatus

`func (o *NumbersV2RegulatoryComplianceBundle) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *NumbersV2RegulatoryComplianceBundle) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *NumbersV2RegulatoryComplianceBundle) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *NumbersV2RegulatoryComplianceBundle) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *NumbersV2RegulatoryComplianceBundle) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *NumbersV2RegulatoryComplianceBundle) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetStatusCallback

`func (o *NumbersV2RegulatoryComplianceBundle) GetStatusCallback() string`

GetStatusCallback returns the StatusCallback field if non-nil, zero value otherwise.

### GetStatusCallbackOk

`func (o *NumbersV2RegulatoryComplianceBundle) GetStatusCallbackOk() (*string, bool)`

GetStatusCallbackOk returns a tuple with the StatusCallback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCallback

`func (o *NumbersV2RegulatoryComplianceBundle) SetStatusCallback(v string)`

SetStatusCallback sets StatusCallback field to given value.

### HasStatusCallback

`func (o *NumbersV2RegulatoryComplianceBundle) HasStatusCallback() bool`

HasStatusCallback returns a boolean if a field has been set.

### SetStatusCallbackNil

`func (o *NumbersV2RegulatoryComplianceBundle) SetStatusCallbackNil(b bool)`

 SetStatusCallbackNil sets the value for StatusCallback to be an explicit nil

### UnsetStatusCallback
`func (o *NumbersV2RegulatoryComplianceBundle) UnsetStatusCallback()`

UnsetStatusCallback ensures that no value is present for StatusCallback, not even an explicit nil
### GetUrl

`func (o *NumbersV2RegulatoryComplianceBundle) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *NumbersV2RegulatoryComplianceBundle) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *NumbersV2RegulatoryComplianceBundle) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *NumbersV2RegulatoryComplianceBundle) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *NumbersV2RegulatoryComplianceBundle) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *NumbersV2RegulatoryComplianceBundle) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetValidUntil

`func (o *NumbersV2RegulatoryComplianceBundle) GetValidUntil() time.Time`

GetValidUntil returns the ValidUntil field if non-nil, zero value otherwise.

### GetValidUntilOk

`func (o *NumbersV2RegulatoryComplianceBundle) GetValidUntilOk() (*time.Time, bool)`

GetValidUntilOk returns a tuple with the ValidUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidUntil

`func (o *NumbersV2RegulatoryComplianceBundle) SetValidUntil(v time.Time)`

SetValidUntil sets ValidUntil field to given value.

### HasValidUntil

`func (o *NumbersV2RegulatoryComplianceBundle) HasValidUntil() bool`

HasValidUntil returns a boolean if a field has been set.

### SetValidUntilNil

`func (o *NumbersV2RegulatoryComplianceBundle) SetValidUntilNil(b bool)`

 SetValidUntilNil sets the value for ValidUntil to be an explicit nil

### UnsetValidUntil
`func (o *NumbersV2RegulatoryComplianceBundle) UnsetValidUntil()`

UnsetValidUntil ensures that no value is present for ValidUntil, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


