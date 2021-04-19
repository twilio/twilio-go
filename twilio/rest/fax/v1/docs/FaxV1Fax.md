# FaxV1Fax

## Properties

Name | Type | Description
------------ | ------------- | -------------
**AccountSid** | Pointer to **NullableString** | The SID of the Account that created the resource | [optional] 
**ApiVersion** | Pointer to **NullableString** | The API version used to transmit the fax | [optional] 
**DateCreated** | Pointer to **NullableTime** | The ISO 8601 formatted date and time in GMT when the resource was created | [optional] 
**DateUpdated** | Pointer to **NullableTime** | The ISO 8601 formatted date and time in GMT when the resource was last updated | [optional] 
**Direction** | Pointer to **NullableString** | The direction of the fax | [optional] 
**Duration** | Pointer to **NullableInt32** | The time it took to transmit the fax | [optional] 
**From** | Pointer to **NullableString** | The number the fax was sent from | [optional] 
**Links** | Pointer to **map[string]interface{}** | The URLs of the fax&#39;s related resources | [optional] 
**MediaSid** | Pointer to **NullableString** | The SID of the FaxMedia resource that is associated with the Fax | [optional] 
**MediaUrl** | Pointer to **NullableString** | The Twilio-hosted URL that can be used to download fax media | [optional] 
**NumPages** | Pointer to **NullableInt32** | The number of pages contained in the fax document | [optional] 
**Price** | Pointer to **NullableFloat32** | The fax transmission price | [optional] 
**PriceUnit** | Pointer to **NullableString** | The ISO 4217 currency used for billing | [optional] 
**Quality** | Pointer to **NullableString** | The quality of the fax | [optional] 
**Sid** | Pointer to **NullableString** | The unique string that identifies the resource | [optional] 
**Status** | Pointer to **NullableString** | The status of the fax | [optional] 
**To** | Pointer to **NullableString** | The phone number that received the fax | [optional] 
**Url** | Pointer to **NullableString** | The absolute URL of the fax resource | [optional] 

## Methods

### NewFaxV1Fax

`func NewFaxV1Fax() *FaxV1Fax`

NewFaxV1Fax instantiates a new FaxV1Fax object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFaxV1FaxWithDefaults

`func NewFaxV1FaxWithDefaults() *FaxV1Fax`

NewFaxV1FaxWithDefaults instantiates a new FaxV1Fax object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSid

`func (o *FaxV1Fax) GetAccountSid() string`

GetAccountSid returns the AccountSid field if non-nil, zero value otherwise.

### GetAccountSidOk

`func (o *FaxV1Fax) GetAccountSidOk() (*string, bool)`

GetAccountSidOk returns a tuple with the AccountSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSid

`func (o *FaxV1Fax) SetAccountSid(v string)`

SetAccountSid sets AccountSid field to given value.

### HasAccountSid

`func (o *FaxV1Fax) HasAccountSid() bool`

HasAccountSid returns a boolean if a field has been set.

### SetAccountSidNil

`func (o *FaxV1Fax) SetAccountSidNil(b bool)`

 SetAccountSidNil sets the value for AccountSid to be an explicit nil

### UnsetAccountSid
`func (o *FaxV1Fax) UnsetAccountSid()`

UnsetAccountSid ensures that no value is present for AccountSid, not even an explicit nil
### GetApiVersion

`func (o *FaxV1Fax) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *FaxV1Fax) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *FaxV1Fax) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *FaxV1Fax) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### SetApiVersionNil

`func (o *FaxV1Fax) SetApiVersionNil(b bool)`

 SetApiVersionNil sets the value for ApiVersion to be an explicit nil

### UnsetApiVersion
`func (o *FaxV1Fax) UnsetApiVersion()`

UnsetApiVersion ensures that no value is present for ApiVersion, not even an explicit nil
### GetDateCreated

`func (o *FaxV1Fax) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *FaxV1Fax) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *FaxV1Fax) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *FaxV1Fax) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### SetDateCreatedNil

`func (o *FaxV1Fax) SetDateCreatedNil(b bool)`

 SetDateCreatedNil sets the value for DateCreated to be an explicit nil

### UnsetDateCreated
`func (o *FaxV1Fax) UnsetDateCreated()`

UnsetDateCreated ensures that no value is present for DateCreated, not even an explicit nil
### GetDateUpdated

`func (o *FaxV1Fax) GetDateUpdated() time.Time`

GetDateUpdated returns the DateUpdated field if non-nil, zero value otherwise.

### GetDateUpdatedOk

`func (o *FaxV1Fax) GetDateUpdatedOk() (*time.Time, bool)`

GetDateUpdatedOk returns a tuple with the DateUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateUpdated

`func (o *FaxV1Fax) SetDateUpdated(v time.Time)`

SetDateUpdated sets DateUpdated field to given value.

### HasDateUpdated

`func (o *FaxV1Fax) HasDateUpdated() bool`

HasDateUpdated returns a boolean if a field has been set.

### SetDateUpdatedNil

`func (o *FaxV1Fax) SetDateUpdatedNil(b bool)`

 SetDateUpdatedNil sets the value for DateUpdated to be an explicit nil

### UnsetDateUpdated
`func (o *FaxV1Fax) UnsetDateUpdated()`

UnsetDateUpdated ensures that no value is present for DateUpdated, not even an explicit nil
### GetDirection

`func (o *FaxV1Fax) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *FaxV1Fax) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *FaxV1Fax) SetDirection(v string)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *FaxV1Fax) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### SetDirectionNil

`func (o *FaxV1Fax) SetDirectionNil(b bool)`

 SetDirectionNil sets the value for Direction to be an explicit nil

### UnsetDirection
`func (o *FaxV1Fax) UnsetDirection()`

UnsetDirection ensures that no value is present for Direction, not even an explicit nil
### GetDuration

`func (o *FaxV1Fax) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *FaxV1Fax) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *FaxV1Fax) SetDuration(v int32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *FaxV1Fax) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### SetDurationNil

`func (o *FaxV1Fax) SetDurationNil(b bool)`

 SetDurationNil sets the value for Duration to be an explicit nil

### UnsetDuration
`func (o *FaxV1Fax) UnsetDuration()`

UnsetDuration ensures that no value is present for Duration, not even an explicit nil
### GetFrom

`func (o *FaxV1Fax) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *FaxV1Fax) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *FaxV1Fax) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *FaxV1Fax) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### SetFromNil

`func (o *FaxV1Fax) SetFromNil(b bool)`

 SetFromNil sets the value for From to be an explicit nil

### UnsetFrom
`func (o *FaxV1Fax) UnsetFrom()`

UnsetFrom ensures that no value is present for From, not even an explicit nil
### GetLinks

`func (o *FaxV1Fax) GetLinks() map[string]interface{}`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *FaxV1Fax) GetLinksOk() (*map[string]interface{}, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *FaxV1Fax) SetLinks(v map[string]interface{})`

SetLinks sets Links field to given value.

### HasLinks

`func (o *FaxV1Fax) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### SetLinksNil

`func (o *FaxV1Fax) SetLinksNil(b bool)`

 SetLinksNil sets the value for Links to be an explicit nil

### UnsetLinks
`func (o *FaxV1Fax) UnsetLinks()`

UnsetLinks ensures that no value is present for Links, not even an explicit nil
### GetMediaSid

`func (o *FaxV1Fax) GetMediaSid() string`

GetMediaSid returns the MediaSid field if non-nil, zero value otherwise.

### GetMediaSidOk

`func (o *FaxV1Fax) GetMediaSidOk() (*string, bool)`

GetMediaSidOk returns a tuple with the MediaSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaSid

`func (o *FaxV1Fax) SetMediaSid(v string)`

SetMediaSid sets MediaSid field to given value.

### HasMediaSid

`func (o *FaxV1Fax) HasMediaSid() bool`

HasMediaSid returns a boolean if a field has been set.

### SetMediaSidNil

`func (o *FaxV1Fax) SetMediaSidNil(b bool)`

 SetMediaSidNil sets the value for MediaSid to be an explicit nil

### UnsetMediaSid
`func (o *FaxV1Fax) UnsetMediaSid()`

UnsetMediaSid ensures that no value is present for MediaSid, not even an explicit nil
### GetMediaUrl

`func (o *FaxV1Fax) GetMediaUrl() string`

GetMediaUrl returns the MediaUrl field if non-nil, zero value otherwise.

### GetMediaUrlOk

`func (o *FaxV1Fax) GetMediaUrlOk() (*string, bool)`

GetMediaUrlOk returns a tuple with the MediaUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaUrl

`func (o *FaxV1Fax) SetMediaUrl(v string)`

SetMediaUrl sets MediaUrl field to given value.

### HasMediaUrl

`func (o *FaxV1Fax) HasMediaUrl() bool`

HasMediaUrl returns a boolean if a field has been set.

### SetMediaUrlNil

`func (o *FaxV1Fax) SetMediaUrlNil(b bool)`

 SetMediaUrlNil sets the value for MediaUrl to be an explicit nil

### UnsetMediaUrl
`func (o *FaxV1Fax) UnsetMediaUrl()`

UnsetMediaUrl ensures that no value is present for MediaUrl, not even an explicit nil
### GetNumPages

`func (o *FaxV1Fax) GetNumPages() int32`

GetNumPages returns the NumPages field if non-nil, zero value otherwise.

### GetNumPagesOk

`func (o *FaxV1Fax) GetNumPagesOk() (*int32, bool)`

GetNumPagesOk returns a tuple with the NumPages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumPages

`func (o *FaxV1Fax) SetNumPages(v int32)`

SetNumPages sets NumPages field to given value.

### HasNumPages

`func (o *FaxV1Fax) HasNumPages() bool`

HasNumPages returns a boolean if a field has been set.

### SetNumPagesNil

`func (o *FaxV1Fax) SetNumPagesNil(b bool)`

 SetNumPagesNil sets the value for NumPages to be an explicit nil

### UnsetNumPages
`func (o *FaxV1Fax) UnsetNumPages()`

UnsetNumPages ensures that no value is present for NumPages, not even an explicit nil
### GetPrice

`func (o *FaxV1Fax) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *FaxV1Fax) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *FaxV1Fax) SetPrice(v float32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *FaxV1Fax) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### SetPriceNil

`func (o *FaxV1Fax) SetPriceNil(b bool)`

 SetPriceNil sets the value for Price to be an explicit nil

### UnsetPrice
`func (o *FaxV1Fax) UnsetPrice()`

UnsetPrice ensures that no value is present for Price, not even an explicit nil
### GetPriceUnit

`func (o *FaxV1Fax) GetPriceUnit() string`

GetPriceUnit returns the PriceUnit field if non-nil, zero value otherwise.

### GetPriceUnitOk

`func (o *FaxV1Fax) GetPriceUnitOk() (*string, bool)`

GetPriceUnitOk returns a tuple with the PriceUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceUnit

`func (o *FaxV1Fax) SetPriceUnit(v string)`

SetPriceUnit sets PriceUnit field to given value.

### HasPriceUnit

`func (o *FaxV1Fax) HasPriceUnit() bool`

HasPriceUnit returns a boolean if a field has been set.

### SetPriceUnitNil

`func (o *FaxV1Fax) SetPriceUnitNil(b bool)`

 SetPriceUnitNil sets the value for PriceUnit to be an explicit nil

### UnsetPriceUnit
`func (o *FaxV1Fax) UnsetPriceUnit()`

UnsetPriceUnit ensures that no value is present for PriceUnit, not even an explicit nil
### GetQuality

`func (o *FaxV1Fax) GetQuality() string`

GetQuality returns the Quality field if non-nil, zero value otherwise.

### GetQualityOk

`func (o *FaxV1Fax) GetQualityOk() (*string, bool)`

GetQualityOk returns a tuple with the Quality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuality

`func (o *FaxV1Fax) SetQuality(v string)`

SetQuality sets Quality field to given value.

### HasQuality

`func (o *FaxV1Fax) HasQuality() bool`

HasQuality returns a boolean if a field has been set.

### SetQualityNil

`func (o *FaxV1Fax) SetQualityNil(b bool)`

 SetQualityNil sets the value for Quality to be an explicit nil

### UnsetQuality
`func (o *FaxV1Fax) UnsetQuality()`

UnsetQuality ensures that no value is present for Quality, not even an explicit nil
### GetSid

`func (o *FaxV1Fax) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *FaxV1Fax) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *FaxV1Fax) SetSid(v string)`

SetSid sets Sid field to given value.

### HasSid

`func (o *FaxV1Fax) HasSid() bool`

HasSid returns a boolean if a field has been set.

### SetSidNil

`func (o *FaxV1Fax) SetSidNil(b bool)`

 SetSidNil sets the value for Sid to be an explicit nil

### UnsetSid
`func (o *FaxV1Fax) UnsetSid()`

UnsetSid ensures that no value is present for Sid, not even an explicit nil
### GetStatus

`func (o *FaxV1Fax) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *FaxV1Fax) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *FaxV1Fax) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *FaxV1Fax) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *FaxV1Fax) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *FaxV1Fax) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetTo

`func (o *FaxV1Fax) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *FaxV1Fax) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *FaxV1Fax) SetTo(v string)`

SetTo sets To field to given value.

### HasTo

`func (o *FaxV1Fax) HasTo() bool`

HasTo returns a boolean if a field has been set.

### SetToNil

`func (o *FaxV1Fax) SetToNil(b bool)`

 SetToNil sets the value for To to be an explicit nil

### UnsetTo
`func (o *FaxV1Fax) UnsetTo()`

UnsetTo ensures that no value is present for To, not even an explicit nil
### GetUrl

`func (o *FaxV1Fax) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *FaxV1Fax) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *FaxV1Fax) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *FaxV1Fax) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *FaxV1Fax) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *FaxV1Fax) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


