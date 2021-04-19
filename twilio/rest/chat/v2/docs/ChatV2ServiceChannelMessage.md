# ChatV2ServiceChannelMessage

## Properties

Name | Type | Description
------------ | ------------- | -------------
**AccountSid** | Pointer to **NullableString** | The SID of the Account that created the resource | [optional] 
**Attributes** | Pointer to **NullableString** | The JSON string that stores application-specific data | [optional] 
**Body** | Pointer to **NullableString** | The content of the message | [optional] 
**ChannelSid** | Pointer to **NullableString** | The SID of the Channel the Message resource belongs to | [optional] 
**DateCreated** | Pointer to **NullableTime** | The RFC 2822 date and time in GMT when the resource was created | [optional] 
**DateUpdated** | Pointer to **NullableTime** | The RFC 2822 date and time in GMT when the resource was last updated | [optional] 
**From** | Pointer to **NullableString** | The Identity of the message&#39;s author | [optional] 
**Index** | Pointer to **NullableInt32** | The index of the message within the Channel | [optional] 
**LastUpdatedBy** | Pointer to **NullableString** | The Identity of the User who last updated the Message | [optional] 
**Media** | Pointer to **map[string]interface{}** | A Media object that describes the Message&#39;s media if attached; otherwise, null | [optional] 
**ServiceSid** | Pointer to **NullableString** | The SID of the Service that the resource is associated with | [optional] 
**Sid** | Pointer to **NullableString** | The unique string that identifies the resource | [optional] 
**To** | Pointer to **NullableString** | The SID of the Channel that the message was sent to | [optional] 
**Type** | Pointer to **NullableString** | The Message type | [optional] 
**Url** | Pointer to **NullableString** | The absolute URL of the Message resource | [optional] 
**WasEdited** | Pointer to **NullableBool** | Whether the message has been edited since  it was created | [optional] 

## Methods

### NewChatV2ServiceChannelMessage

`func NewChatV2ServiceChannelMessage() *ChatV2ServiceChannelMessage`

NewChatV2ServiceChannelMessage instantiates a new ChatV2ServiceChannelMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChatV2ServiceChannelMessageWithDefaults

`func NewChatV2ServiceChannelMessageWithDefaults() *ChatV2ServiceChannelMessage`

NewChatV2ServiceChannelMessageWithDefaults instantiates a new ChatV2ServiceChannelMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSid

`func (o *ChatV2ServiceChannelMessage) GetAccountSid() string`

GetAccountSid returns the AccountSid field if non-nil, zero value otherwise.

### GetAccountSidOk

`func (o *ChatV2ServiceChannelMessage) GetAccountSidOk() (*string, bool)`

GetAccountSidOk returns a tuple with the AccountSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSid

`func (o *ChatV2ServiceChannelMessage) SetAccountSid(v string)`

SetAccountSid sets AccountSid field to given value.

### HasAccountSid

`func (o *ChatV2ServiceChannelMessage) HasAccountSid() bool`

HasAccountSid returns a boolean if a field has been set.

### SetAccountSidNil

`func (o *ChatV2ServiceChannelMessage) SetAccountSidNil(b bool)`

 SetAccountSidNil sets the value for AccountSid to be an explicit nil

### UnsetAccountSid
`func (o *ChatV2ServiceChannelMessage) UnsetAccountSid()`

UnsetAccountSid ensures that no value is present for AccountSid, not even an explicit nil
### GetAttributes

`func (o *ChatV2ServiceChannelMessage) GetAttributes() string`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *ChatV2ServiceChannelMessage) GetAttributesOk() (*string, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *ChatV2ServiceChannelMessage) SetAttributes(v string)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *ChatV2ServiceChannelMessage) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### SetAttributesNil

`func (o *ChatV2ServiceChannelMessage) SetAttributesNil(b bool)`

 SetAttributesNil sets the value for Attributes to be an explicit nil

### UnsetAttributes
`func (o *ChatV2ServiceChannelMessage) UnsetAttributes()`

UnsetAttributes ensures that no value is present for Attributes, not even an explicit nil
### GetBody

`func (o *ChatV2ServiceChannelMessage) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *ChatV2ServiceChannelMessage) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *ChatV2ServiceChannelMessage) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *ChatV2ServiceChannelMessage) HasBody() bool`

HasBody returns a boolean if a field has been set.

### SetBodyNil

`func (o *ChatV2ServiceChannelMessage) SetBodyNil(b bool)`

 SetBodyNil sets the value for Body to be an explicit nil

### UnsetBody
`func (o *ChatV2ServiceChannelMessage) UnsetBody()`

UnsetBody ensures that no value is present for Body, not even an explicit nil
### GetChannelSid

`func (o *ChatV2ServiceChannelMessage) GetChannelSid() string`

GetChannelSid returns the ChannelSid field if non-nil, zero value otherwise.

### GetChannelSidOk

`func (o *ChatV2ServiceChannelMessage) GetChannelSidOk() (*string, bool)`

GetChannelSidOk returns a tuple with the ChannelSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelSid

`func (o *ChatV2ServiceChannelMessage) SetChannelSid(v string)`

SetChannelSid sets ChannelSid field to given value.

### HasChannelSid

`func (o *ChatV2ServiceChannelMessage) HasChannelSid() bool`

HasChannelSid returns a boolean if a field has been set.

### SetChannelSidNil

`func (o *ChatV2ServiceChannelMessage) SetChannelSidNil(b bool)`

 SetChannelSidNil sets the value for ChannelSid to be an explicit nil

### UnsetChannelSid
`func (o *ChatV2ServiceChannelMessage) UnsetChannelSid()`

UnsetChannelSid ensures that no value is present for ChannelSid, not even an explicit nil
### GetDateCreated

`func (o *ChatV2ServiceChannelMessage) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *ChatV2ServiceChannelMessage) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *ChatV2ServiceChannelMessage) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *ChatV2ServiceChannelMessage) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### SetDateCreatedNil

`func (o *ChatV2ServiceChannelMessage) SetDateCreatedNil(b bool)`

 SetDateCreatedNil sets the value for DateCreated to be an explicit nil

### UnsetDateCreated
`func (o *ChatV2ServiceChannelMessage) UnsetDateCreated()`

UnsetDateCreated ensures that no value is present for DateCreated, not even an explicit nil
### GetDateUpdated

`func (o *ChatV2ServiceChannelMessage) GetDateUpdated() time.Time`

GetDateUpdated returns the DateUpdated field if non-nil, zero value otherwise.

### GetDateUpdatedOk

`func (o *ChatV2ServiceChannelMessage) GetDateUpdatedOk() (*time.Time, bool)`

GetDateUpdatedOk returns a tuple with the DateUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateUpdated

`func (o *ChatV2ServiceChannelMessage) SetDateUpdated(v time.Time)`

SetDateUpdated sets DateUpdated field to given value.

### HasDateUpdated

`func (o *ChatV2ServiceChannelMessage) HasDateUpdated() bool`

HasDateUpdated returns a boolean if a field has been set.

### SetDateUpdatedNil

`func (o *ChatV2ServiceChannelMessage) SetDateUpdatedNil(b bool)`

 SetDateUpdatedNil sets the value for DateUpdated to be an explicit nil

### UnsetDateUpdated
`func (o *ChatV2ServiceChannelMessage) UnsetDateUpdated()`

UnsetDateUpdated ensures that no value is present for DateUpdated, not even an explicit nil
### GetFrom

`func (o *ChatV2ServiceChannelMessage) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *ChatV2ServiceChannelMessage) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *ChatV2ServiceChannelMessage) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *ChatV2ServiceChannelMessage) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### SetFromNil

`func (o *ChatV2ServiceChannelMessage) SetFromNil(b bool)`

 SetFromNil sets the value for From to be an explicit nil

### UnsetFrom
`func (o *ChatV2ServiceChannelMessage) UnsetFrom()`

UnsetFrom ensures that no value is present for From, not even an explicit nil
### GetIndex

`func (o *ChatV2ServiceChannelMessage) GetIndex() int32`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *ChatV2ServiceChannelMessage) GetIndexOk() (*int32, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *ChatV2ServiceChannelMessage) SetIndex(v int32)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *ChatV2ServiceChannelMessage) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### SetIndexNil

`func (o *ChatV2ServiceChannelMessage) SetIndexNil(b bool)`

 SetIndexNil sets the value for Index to be an explicit nil

### UnsetIndex
`func (o *ChatV2ServiceChannelMessage) UnsetIndex()`

UnsetIndex ensures that no value is present for Index, not even an explicit nil
### GetLastUpdatedBy

`func (o *ChatV2ServiceChannelMessage) GetLastUpdatedBy() string`

GetLastUpdatedBy returns the LastUpdatedBy field if non-nil, zero value otherwise.

### GetLastUpdatedByOk

`func (o *ChatV2ServiceChannelMessage) GetLastUpdatedByOk() (*string, bool)`

GetLastUpdatedByOk returns a tuple with the LastUpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedBy

`func (o *ChatV2ServiceChannelMessage) SetLastUpdatedBy(v string)`

SetLastUpdatedBy sets LastUpdatedBy field to given value.

### HasLastUpdatedBy

`func (o *ChatV2ServiceChannelMessage) HasLastUpdatedBy() bool`

HasLastUpdatedBy returns a boolean if a field has been set.

### SetLastUpdatedByNil

`func (o *ChatV2ServiceChannelMessage) SetLastUpdatedByNil(b bool)`

 SetLastUpdatedByNil sets the value for LastUpdatedBy to be an explicit nil

### UnsetLastUpdatedBy
`func (o *ChatV2ServiceChannelMessage) UnsetLastUpdatedBy()`

UnsetLastUpdatedBy ensures that no value is present for LastUpdatedBy, not even an explicit nil
### GetMedia

`func (o *ChatV2ServiceChannelMessage) GetMedia() map[string]interface{}`

GetMedia returns the Media field if non-nil, zero value otherwise.

### GetMediaOk

`func (o *ChatV2ServiceChannelMessage) GetMediaOk() (*map[string]interface{}, bool)`

GetMediaOk returns a tuple with the Media field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMedia

`func (o *ChatV2ServiceChannelMessage) SetMedia(v map[string]interface{})`

SetMedia sets Media field to given value.

### HasMedia

`func (o *ChatV2ServiceChannelMessage) HasMedia() bool`

HasMedia returns a boolean if a field has been set.

### SetMediaNil

`func (o *ChatV2ServiceChannelMessage) SetMediaNil(b bool)`

 SetMediaNil sets the value for Media to be an explicit nil

### UnsetMedia
`func (o *ChatV2ServiceChannelMessage) UnsetMedia()`

UnsetMedia ensures that no value is present for Media, not even an explicit nil
### GetServiceSid

`func (o *ChatV2ServiceChannelMessage) GetServiceSid() string`

GetServiceSid returns the ServiceSid field if non-nil, zero value otherwise.

### GetServiceSidOk

`func (o *ChatV2ServiceChannelMessage) GetServiceSidOk() (*string, bool)`

GetServiceSidOk returns a tuple with the ServiceSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceSid

`func (o *ChatV2ServiceChannelMessage) SetServiceSid(v string)`

SetServiceSid sets ServiceSid field to given value.

### HasServiceSid

`func (o *ChatV2ServiceChannelMessage) HasServiceSid() bool`

HasServiceSid returns a boolean if a field has been set.

### SetServiceSidNil

`func (o *ChatV2ServiceChannelMessage) SetServiceSidNil(b bool)`

 SetServiceSidNil sets the value for ServiceSid to be an explicit nil

### UnsetServiceSid
`func (o *ChatV2ServiceChannelMessage) UnsetServiceSid()`

UnsetServiceSid ensures that no value is present for ServiceSid, not even an explicit nil
### GetSid

`func (o *ChatV2ServiceChannelMessage) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *ChatV2ServiceChannelMessage) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *ChatV2ServiceChannelMessage) SetSid(v string)`

SetSid sets Sid field to given value.

### HasSid

`func (o *ChatV2ServiceChannelMessage) HasSid() bool`

HasSid returns a boolean if a field has been set.

### SetSidNil

`func (o *ChatV2ServiceChannelMessage) SetSidNil(b bool)`

 SetSidNil sets the value for Sid to be an explicit nil

### UnsetSid
`func (o *ChatV2ServiceChannelMessage) UnsetSid()`

UnsetSid ensures that no value is present for Sid, not even an explicit nil
### GetTo

`func (o *ChatV2ServiceChannelMessage) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *ChatV2ServiceChannelMessage) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *ChatV2ServiceChannelMessage) SetTo(v string)`

SetTo sets To field to given value.

### HasTo

`func (o *ChatV2ServiceChannelMessage) HasTo() bool`

HasTo returns a boolean if a field has been set.

### SetToNil

`func (o *ChatV2ServiceChannelMessage) SetToNil(b bool)`

 SetToNil sets the value for To to be an explicit nil

### UnsetTo
`func (o *ChatV2ServiceChannelMessage) UnsetTo()`

UnsetTo ensures that no value is present for To, not even an explicit nil
### GetType

`func (o *ChatV2ServiceChannelMessage) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ChatV2ServiceChannelMessage) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ChatV2ServiceChannelMessage) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ChatV2ServiceChannelMessage) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *ChatV2ServiceChannelMessage) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *ChatV2ServiceChannelMessage) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetUrl

`func (o *ChatV2ServiceChannelMessage) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ChatV2ServiceChannelMessage) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ChatV2ServiceChannelMessage) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *ChatV2ServiceChannelMessage) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *ChatV2ServiceChannelMessage) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *ChatV2ServiceChannelMessage) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetWasEdited

`func (o *ChatV2ServiceChannelMessage) GetWasEdited() bool`

GetWasEdited returns the WasEdited field if non-nil, zero value otherwise.

### GetWasEditedOk

`func (o *ChatV2ServiceChannelMessage) GetWasEditedOk() (*bool, bool)`

GetWasEditedOk returns a tuple with the WasEdited field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWasEdited

`func (o *ChatV2ServiceChannelMessage) SetWasEdited(v bool)`

SetWasEdited sets WasEdited field to given value.

### HasWasEdited

`func (o *ChatV2ServiceChannelMessage) HasWasEdited() bool`

HasWasEdited returns a boolean if a field has been set.

### SetWasEditedNil

`func (o *ChatV2ServiceChannelMessage) SetWasEditedNil(b bool)`

 SetWasEditedNil sets the value for WasEdited to be an explicit nil

### UnsetWasEdited
`func (o *ChatV2ServiceChannelMessage) UnsetWasEdited()`

UnsetWasEdited ensures that no value is present for WasEdited, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


