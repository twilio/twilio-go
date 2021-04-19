# ProxyV1ServiceSession

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **NullableString** | The SID of the Account that created the resource | [optional] 
**ClosedReason** | Pointer to **NullableString** | The reason the Session ended | [optional] 
**DateCreated** | Pointer to **NullableTime** | The ISO 8601 date and time in GMT when the resource was created | [optional] 
**DateEnded** | Pointer to **NullableTime** | The ISO 8601 date when the Session ended | [optional] 
**DateExpiry** | Pointer to **NullableTime** | The ISO 8601 date when the Session should expire | [optional] 
**DateLastInteraction** | Pointer to **NullableTime** | The ISO 8601 date when the Session last had an interaction | [optional] 
**DateStarted** | Pointer to **NullableTime** | The ISO 8601 date when the Session started | [optional] 
**DateUpdated** | Pointer to **NullableTime** | The ISO 8601 date and time in GMT when the resource was last updated | [optional] 
**Links** | Pointer to **map[string]interface{}** | The URLs of resources related to the Session | [optional] 
**Mode** | Pointer to **NullableString** | The Mode of the Session | [optional] 
**ServiceSid** | Pointer to **NullableString** | The SID of the resource&#39;s parent Service | [optional] 
**Sid** | Pointer to **NullableString** | The unique string that identifies the resource | [optional] 
**Status** | Pointer to **NullableString** | The status of the Session | [optional] 
**Ttl** | Pointer to **NullableInt32** | When the session will expire | [optional] 
**UniqueName** | Pointer to **NullableString** | An application-defined string that uniquely identifies the resource | [optional] 
**Url** | Pointer to **NullableString** | The absolute URL of the Session resource | [optional] 

## Methods

### NewProxyV1ServiceSession

`func NewProxyV1ServiceSession() *ProxyV1ServiceSession`

NewProxyV1ServiceSession instantiates a new ProxyV1ServiceSession object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProxyV1ServiceSessionWithDefaults

`func NewProxyV1ServiceSessionWithDefaults() *ProxyV1ServiceSession`

NewProxyV1ServiceSessionWithDefaults instantiates a new ProxyV1ServiceSession object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSid

`func (o *ProxyV1ServiceSession) GetAccountSid() string`

GetAccountSid returns the AccountSid field if non-nil, zero value otherwise.

### GetAccountSidOk

`func (o *ProxyV1ServiceSession) GetAccountSidOk() (*string, bool)`

GetAccountSidOk returns a tuple with the AccountSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSid

`func (o *ProxyV1ServiceSession) SetAccountSid(v string)`

SetAccountSid sets AccountSid field to given value.

### HasAccountSid

`func (o *ProxyV1ServiceSession) HasAccountSid() bool`

HasAccountSid returns a boolean if a field has been set.

### SetAccountSidNil

`func (o *ProxyV1ServiceSession) SetAccountSidNil(b bool)`

 SetAccountSidNil sets the value for AccountSid to be an explicit nil

### UnsetAccountSid
`func (o *ProxyV1ServiceSession) UnsetAccountSid()`

UnsetAccountSid ensures that no value is present for AccountSid, not even an explicit nil
### GetClosedReason

`func (o *ProxyV1ServiceSession) GetClosedReason() string`

GetClosedReason returns the ClosedReason field if non-nil, zero value otherwise.

### GetClosedReasonOk

`func (o *ProxyV1ServiceSession) GetClosedReasonOk() (*string, bool)`

GetClosedReasonOk returns a tuple with the ClosedReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosedReason

`func (o *ProxyV1ServiceSession) SetClosedReason(v string)`

SetClosedReason sets ClosedReason field to given value.

### HasClosedReason

`func (o *ProxyV1ServiceSession) HasClosedReason() bool`

HasClosedReason returns a boolean if a field has been set.

### SetClosedReasonNil

`func (o *ProxyV1ServiceSession) SetClosedReasonNil(b bool)`

 SetClosedReasonNil sets the value for ClosedReason to be an explicit nil

### UnsetClosedReason
`func (o *ProxyV1ServiceSession) UnsetClosedReason()`

UnsetClosedReason ensures that no value is present for ClosedReason, not even an explicit nil
### GetDateCreated

`func (o *ProxyV1ServiceSession) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *ProxyV1ServiceSession) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *ProxyV1ServiceSession) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *ProxyV1ServiceSession) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### SetDateCreatedNil

`func (o *ProxyV1ServiceSession) SetDateCreatedNil(b bool)`

 SetDateCreatedNil sets the value for DateCreated to be an explicit nil

### UnsetDateCreated
`func (o *ProxyV1ServiceSession) UnsetDateCreated()`

UnsetDateCreated ensures that no value is present for DateCreated, not even an explicit nil
### GetDateEnded

`func (o *ProxyV1ServiceSession) GetDateEnded() time.Time`

GetDateEnded returns the DateEnded field if non-nil, zero value otherwise.

### GetDateEndedOk

`func (o *ProxyV1ServiceSession) GetDateEndedOk() (*time.Time, bool)`

GetDateEndedOk returns a tuple with the DateEnded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateEnded

`func (o *ProxyV1ServiceSession) SetDateEnded(v time.Time)`

SetDateEnded sets DateEnded field to given value.

### HasDateEnded

`func (o *ProxyV1ServiceSession) HasDateEnded() bool`

HasDateEnded returns a boolean if a field has been set.

### SetDateEndedNil

`func (o *ProxyV1ServiceSession) SetDateEndedNil(b bool)`

 SetDateEndedNil sets the value for DateEnded to be an explicit nil

### UnsetDateEnded
`func (o *ProxyV1ServiceSession) UnsetDateEnded()`

UnsetDateEnded ensures that no value is present for DateEnded, not even an explicit nil
### GetDateExpiry

`func (o *ProxyV1ServiceSession) GetDateExpiry() time.Time`

GetDateExpiry returns the DateExpiry field if non-nil, zero value otherwise.

### GetDateExpiryOk

`func (o *ProxyV1ServiceSession) GetDateExpiryOk() (*time.Time, bool)`

GetDateExpiryOk returns a tuple with the DateExpiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateExpiry

`func (o *ProxyV1ServiceSession) SetDateExpiry(v time.Time)`

SetDateExpiry sets DateExpiry field to given value.

### HasDateExpiry

`func (o *ProxyV1ServiceSession) HasDateExpiry() bool`

HasDateExpiry returns a boolean if a field has been set.

### SetDateExpiryNil

`func (o *ProxyV1ServiceSession) SetDateExpiryNil(b bool)`

 SetDateExpiryNil sets the value for DateExpiry to be an explicit nil

### UnsetDateExpiry
`func (o *ProxyV1ServiceSession) UnsetDateExpiry()`

UnsetDateExpiry ensures that no value is present for DateExpiry, not even an explicit nil
### GetDateLastInteraction

`func (o *ProxyV1ServiceSession) GetDateLastInteraction() time.Time`

GetDateLastInteraction returns the DateLastInteraction field if non-nil, zero value otherwise.

### GetDateLastInteractionOk

`func (o *ProxyV1ServiceSession) GetDateLastInteractionOk() (*time.Time, bool)`

GetDateLastInteractionOk returns a tuple with the DateLastInteraction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateLastInteraction

`func (o *ProxyV1ServiceSession) SetDateLastInteraction(v time.Time)`

SetDateLastInteraction sets DateLastInteraction field to given value.

### HasDateLastInteraction

`func (o *ProxyV1ServiceSession) HasDateLastInteraction() bool`

HasDateLastInteraction returns a boolean if a field has been set.

### SetDateLastInteractionNil

`func (o *ProxyV1ServiceSession) SetDateLastInteractionNil(b bool)`

 SetDateLastInteractionNil sets the value for DateLastInteraction to be an explicit nil

### UnsetDateLastInteraction
`func (o *ProxyV1ServiceSession) UnsetDateLastInteraction()`

UnsetDateLastInteraction ensures that no value is present for DateLastInteraction, not even an explicit nil
### GetDateStarted

`func (o *ProxyV1ServiceSession) GetDateStarted() time.Time`

GetDateStarted returns the DateStarted field if non-nil, zero value otherwise.

### GetDateStartedOk

`func (o *ProxyV1ServiceSession) GetDateStartedOk() (*time.Time, bool)`

GetDateStartedOk returns a tuple with the DateStarted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateStarted

`func (o *ProxyV1ServiceSession) SetDateStarted(v time.Time)`

SetDateStarted sets DateStarted field to given value.

### HasDateStarted

`func (o *ProxyV1ServiceSession) HasDateStarted() bool`

HasDateStarted returns a boolean if a field has been set.

### SetDateStartedNil

`func (o *ProxyV1ServiceSession) SetDateStartedNil(b bool)`

 SetDateStartedNil sets the value for DateStarted to be an explicit nil

### UnsetDateStarted
`func (o *ProxyV1ServiceSession) UnsetDateStarted()`

UnsetDateStarted ensures that no value is present for DateStarted, not even an explicit nil
### GetDateUpdated

`func (o *ProxyV1ServiceSession) GetDateUpdated() time.Time`

GetDateUpdated returns the DateUpdated field if non-nil, zero value otherwise.

### GetDateUpdatedOk

`func (o *ProxyV1ServiceSession) GetDateUpdatedOk() (*time.Time, bool)`

GetDateUpdatedOk returns a tuple with the DateUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateUpdated

`func (o *ProxyV1ServiceSession) SetDateUpdated(v time.Time)`

SetDateUpdated sets DateUpdated field to given value.

### HasDateUpdated

`func (o *ProxyV1ServiceSession) HasDateUpdated() bool`

HasDateUpdated returns a boolean if a field has been set.

### SetDateUpdatedNil

`func (o *ProxyV1ServiceSession) SetDateUpdatedNil(b bool)`

 SetDateUpdatedNil sets the value for DateUpdated to be an explicit nil

### UnsetDateUpdated
`func (o *ProxyV1ServiceSession) UnsetDateUpdated()`

UnsetDateUpdated ensures that no value is present for DateUpdated, not even an explicit nil
### GetLinks

`func (o *ProxyV1ServiceSession) GetLinks() map[string]interface{}`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ProxyV1ServiceSession) GetLinksOk() (*map[string]interface{}, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ProxyV1ServiceSession) SetLinks(v map[string]interface{})`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ProxyV1ServiceSession) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### SetLinksNil

`func (o *ProxyV1ServiceSession) SetLinksNil(b bool)`

 SetLinksNil sets the value for Links to be an explicit nil

### UnsetLinks
`func (o *ProxyV1ServiceSession) UnsetLinks()`

UnsetLinks ensures that no value is present for Links, not even an explicit nil
### GetMode

`func (o *ProxyV1ServiceSession) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *ProxyV1ServiceSession) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *ProxyV1ServiceSession) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *ProxyV1ServiceSession) HasMode() bool`

HasMode returns a boolean if a field has been set.

### SetModeNil

`func (o *ProxyV1ServiceSession) SetModeNil(b bool)`

 SetModeNil sets the value for Mode to be an explicit nil

### UnsetMode
`func (o *ProxyV1ServiceSession) UnsetMode()`

UnsetMode ensures that no value is present for Mode, not even an explicit nil
### GetServiceSid

`func (o *ProxyV1ServiceSession) GetServiceSid() string`

GetServiceSid returns the ServiceSid field if non-nil, zero value otherwise.

### GetServiceSidOk

`func (o *ProxyV1ServiceSession) GetServiceSidOk() (*string, bool)`

GetServiceSidOk returns a tuple with the ServiceSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceSid

`func (o *ProxyV1ServiceSession) SetServiceSid(v string)`

SetServiceSid sets ServiceSid field to given value.

### HasServiceSid

`func (o *ProxyV1ServiceSession) HasServiceSid() bool`

HasServiceSid returns a boolean if a field has been set.

### SetServiceSidNil

`func (o *ProxyV1ServiceSession) SetServiceSidNil(b bool)`

 SetServiceSidNil sets the value for ServiceSid to be an explicit nil

### UnsetServiceSid
`func (o *ProxyV1ServiceSession) UnsetServiceSid()`

UnsetServiceSid ensures that no value is present for ServiceSid, not even an explicit nil
### GetSid

`func (o *ProxyV1ServiceSession) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *ProxyV1ServiceSession) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *ProxyV1ServiceSession) SetSid(v string)`

SetSid sets Sid field to given value.

### HasSid

`func (o *ProxyV1ServiceSession) HasSid() bool`

HasSid returns a boolean if a field has been set.

### SetSidNil

`func (o *ProxyV1ServiceSession) SetSidNil(b bool)`

 SetSidNil sets the value for Sid to be an explicit nil

### UnsetSid
`func (o *ProxyV1ServiceSession) UnsetSid()`

UnsetSid ensures that no value is present for Sid, not even an explicit nil
### GetStatus

`func (o *ProxyV1ServiceSession) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ProxyV1ServiceSession) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ProxyV1ServiceSession) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ProxyV1ServiceSession) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *ProxyV1ServiceSession) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *ProxyV1ServiceSession) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetTtl

`func (o *ProxyV1ServiceSession) GetTtl() int32`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *ProxyV1ServiceSession) GetTtlOk() (*int32, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *ProxyV1ServiceSession) SetTtl(v int32)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *ProxyV1ServiceSession) HasTtl() bool`

HasTtl returns a boolean if a field has been set.

### SetTtlNil

`func (o *ProxyV1ServiceSession) SetTtlNil(b bool)`

 SetTtlNil sets the value for Ttl to be an explicit nil

### UnsetTtl
`func (o *ProxyV1ServiceSession) UnsetTtl()`

UnsetTtl ensures that no value is present for Ttl, not even an explicit nil
### GetUniqueName

`func (o *ProxyV1ServiceSession) GetUniqueName() string`

GetUniqueName returns the UniqueName field if non-nil, zero value otherwise.

### GetUniqueNameOk

`func (o *ProxyV1ServiceSession) GetUniqueNameOk() (*string, bool)`

GetUniqueNameOk returns a tuple with the UniqueName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueName

`func (o *ProxyV1ServiceSession) SetUniqueName(v string)`

SetUniqueName sets UniqueName field to given value.

### HasUniqueName

`func (o *ProxyV1ServiceSession) HasUniqueName() bool`

HasUniqueName returns a boolean if a field has been set.

### SetUniqueNameNil

`func (o *ProxyV1ServiceSession) SetUniqueNameNil(b bool)`

 SetUniqueNameNil sets the value for UniqueName to be an explicit nil

### UnsetUniqueName
`func (o *ProxyV1ServiceSession) UnsetUniqueName()`

UnsetUniqueName ensures that no value is present for UniqueName, not even an explicit nil
### GetUrl

`func (o *ProxyV1ServiceSession) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ProxyV1ServiceSession) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ProxyV1ServiceSession) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *ProxyV1ServiceSession) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *ProxyV1ServiceSession) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *ProxyV1ServiceSession) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


