# VoiceV1SourceIpMapping

## Properties

Name | Type | Description
------------ | ------------- | -------------
**DateCreated** | Pointer to **NullableTime** | The RFC 2822 date and time in GMT that the resource was created | [optional] 
**DateUpdated** | Pointer to **NullableTime** | The RFC 2822 date and time in GMT that the resource was last updated | [optional] 
**IpRecordSid** | Pointer to **NullableString** | The unique string that identifies an IP Record | [optional] 
**Sid** | Pointer to **NullableString** | The unique string that identifies the resource | [optional] 
**SipDomainSid** | Pointer to **NullableString** | The unique string that identifies a SIP Domain | [optional] 
**Url** | Pointer to **NullableString** | The absolute URL of the resource | [optional] 

## Methods

### NewVoiceV1SourceIpMapping

`func NewVoiceV1SourceIpMapping() *VoiceV1SourceIpMapping`

NewVoiceV1SourceIpMapping instantiates a new VoiceV1SourceIpMapping object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVoiceV1SourceIpMappingWithDefaults

`func NewVoiceV1SourceIpMappingWithDefaults() *VoiceV1SourceIpMapping`

NewVoiceV1SourceIpMappingWithDefaults instantiates a new VoiceV1SourceIpMapping object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDateCreated

`func (o *VoiceV1SourceIpMapping) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *VoiceV1SourceIpMapping) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *VoiceV1SourceIpMapping) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *VoiceV1SourceIpMapping) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### SetDateCreatedNil

`func (o *VoiceV1SourceIpMapping) SetDateCreatedNil(b bool)`

 SetDateCreatedNil sets the value for DateCreated to be an explicit nil

### UnsetDateCreated
`func (o *VoiceV1SourceIpMapping) UnsetDateCreated()`

UnsetDateCreated ensures that no value is present for DateCreated, not even an explicit nil
### GetDateUpdated

`func (o *VoiceV1SourceIpMapping) GetDateUpdated() time.Time`

GetDateUpdated returns the DateUpdated field if non-nil, zero value otherwise.

### GetDateUpdatedOk

`func (o *VoiceV1SourceIpMapping) GetDateUpdatedOk() (*time.Time, bool)`

GetDateUpdatedOk returns a tuple with the DateUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateUpdated

`func (o *VoiceV1SourceIpMapping) SetDateUpdated(v time.Time)`

SetDateUpdated sets DateUpdated field to given value.

### HasDateUpdated

`func (o *VoiceV1SourceIpMapping) HasDateUpdated() bool`

HasDateUpdated returns a boolean if a field has been set.

### SetDateUpdatedNil

`func (o *VoiceV1SourceIpMapping) SetDateUpdatedNil(b bool)`

 SetDateUpdatedNil sets the value for DateUpdated to be an explicit nil

### UnsetDateUpdated
`func (o *VoiceV1SourceIpMapping) UnsetDateUpdated()`

UnsetDateUpdated ensures that no value is present for DateUpdated, not even an explicit nil
### GetIpRecordSid

`func (o *VoiceV1SourceIpMapping) GetIpRecordSid() string`

GetIpRecordSid returns the IpRecordSid field if non-nil, zero value otherwise.

### GetIpRecordSidOk

`func (o *VoiceV1SourceIpMapping) GetIpRecordSidOk() (*string, bool)`

GetIpRecordSidOk returns a tuple with the IpRecordSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpRecordSid

`func (o *VoiceV1SourceIpMapping) SetIpRecordSid(v string)`

SetIpRecordSid sets IpRecordSid field to given value.

### HasIpRecordSid

`func (o *VoiceV1SourceIpMapping) HasIpRecordSid() bool`

HasIpRecordSid returns a boolean if a field has been set.

### SetIpRecordSidNil

`func (o *VoiceV1SourceIpMapping) SetIpRecordSidNil(b bool)`

 SetIpRecordSidNil sets the value for IpRecordSid to be an explicit nil

### UnsetIpRecordSid
`func (o *VoiceV1SourceIpMapping) UnsetIpRecordSid()`

UnsetIpRecordSid ensures that no value is present for IpRecordSid, not even an explicit nil
### GetSid

`func (o *VoiceV1SourceIpMapping) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *VoiceV1SourceIpMapping) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *VoiceV1SourceIpMapping) SetSid(v string)`

SetSid sets Sid field to given value.

### HasSid

`func (o *VoiceV1SourceIpMapping) HasSid() bool`

HasSid returns a boolean if a field has been set.

### SetSidNil

`func (o *VoiceV1SourceIpMapping) SetSidNil(b bool)`

 SetSidNil sets the value for Sid to be an explicit nil

### UnsetSid
`func (o *VoiceV1SourceIpMapping) UnsetSid()`

UnsetSid ensures that no value is present for Sid, not even an explicit nil
### GetSipDomainSid

`func (o *VoiceV1SourceIpMapping) GetSipDomainSid() string`

GetSipDomainSid returns the SipDomainSid field if non-nil, zero value otherwise.

### GetSipDomainSidOk

`func (o *VoiceV1SourceIpMapping) GetSipDomainSidOk() (*string, bool)`

GetSipDomainSidOk returns a tuple with the SipDomainSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSipDomainSid

`func (o *VoiceV1SourceIpMapping) SetSipDomainSid(v string)`

SetSipDomainSid sets SipDomainSid field to given value.

### HasSipDomainSid

`func (o *VoiceV1SourceIpMapping) HasSipDomainSid() bool`

HasSipDomainSid returns a boolean if a field has been set.

### SetSipDomainSidNil

`func (o *VoiceV1SourceIpMapping) SetSipDomainSidNil(b bool)`

 SetSipDomainSidNil sets the value for SipDomainSid to be an explicit nil

### UnsetSipDomainSid
`func (o *VoiceV1SourceIpMapping) UnsetSipDomainSid()`

UnsetSipDomainSid ensures that no value is present for SipDomainSid, not even an explicit nil
### GetUrl

`func (o *VoiceV1SourceIpMapping) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *VoiceV1SourceIpMapping) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *VoiceV1SourceIpMapping) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *VoiceV1SourceIpMapping) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *VoiceV1SourceIpMapping) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *VoiceV1SourceIpMapping) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


