# ChatV1ServiceChannelMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **NullableString** | The SID of the Account that created the resource | [optional] 
**Attributes** | Pointer to **NullableString** | The JSON string that stores application-specific data | [optional] 
**Body** | Pointer to **NullableString** | The content of the message | [optional] 
**ChannelSid** | Pointer to **NullableString** | The unique ID of the Channel the Message resource belongs to | [optional] 
**DateCreated** | Pointer to **NullableTime** | The RFC 2822 date and time in GMT when the resource was created | [optional] 
**DateUpdated** | Pointer to **NullableTime** | The RFC 2822 date and time in GMT when the resource was last updated | [optional] 
**From** | Pointer to **NullableString** | The identity of the message&#39;s author | [optional] 
**Index** | Pointer to **NullableInt32** | The index of the message within the Channel | [optional] 
**ServiceSid** | Pointer to **NullableString** | The SID of the Service that the resource is associated with | [optional] 
**Sid** | Pointer to **NullableString** | The unique string that identifies the resource | [optional] 
**To** | Pointer to **NullableString** | The SID of the Channel that the message was sent to | [optional] 
**Url** | Pointer to **NullableString** | The absolute URL of the Message resource | [optional] 
**WasEdited** | Pointer to **NullableBool** | Whether the message has been edited since  it was created | [optional] 

## Methods

### NewChatV1ServiceChannelMessage

`func NewChatV1ServiceChannelMessage() *ChatV1ServiceChannelMessage`

NewChatV1ServiceChannelMessage instantiates a new ChatV1ServiceChannelMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChatV1ServiceChannelMessageWithDefaults

`func NewChatV1ServiceChannelMessageWithDefaults() *ChatV1ServiceChannelMessage`

NewChatV1ServiceChannelMessageWithDefaults instantiates a new ChatV1ServiceChannelMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSid

`func (o *ChatV1ServiceChannelMessage) GetAccountSid() string`

GetAccountSid returns the AccountSid field if non-nil, zero value otherwise.

### GetAccountSidOk

`func (o *ChatV1ServiceChannelMessage) GetAccountSidOk() (*string, bool)`

GetAccountSidOk returns a tuple with the AccountSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSid

`func (o *ChatV1ServiceChannelMessage) SetAccountSid(v string)`

SetAccountSid sets AccountSid field to given value.

### HasAccountSid

`func (o *ChatV1ServiceChannelMessage) HasAccountSid() bool`

HasAccountSid returns a boolean if a field has been set.

### SetAccountSidNil

`func (o *ChatV1ServiceChannelMessage) SetAccountSidNil(b bool)`

 SetAccountSidNil sets the value for AccountSid to be an explicit nil

### UnsetAccountSid
`func (o *ChatV1ServiceChannelMessage) UnsetAccountSid()`

UnsetAccountSid ensures that no value is present for AccountSid, not even an explicit nil
### GetAttributes

`func (o *ChatV1ServiceChannelMessage) GetAttributes() string`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *ChatV1ServiceChannelMessage) GetAttributesOk() (*string, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *ChatV1ServiceChannelMessage) SetAttributes(v string)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *ChatV1ServiceChannelMessage) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### SetAttributesNil

`func (o *ChatV1ServiceChannelMessage) SetAttributesNil(b bool)`

 SetAttributesNil sets the value for Attributes to be an explicit nil

### UnsetAttributes
`func (o *ChatV1ServiceChannelMessage) UnsetAttributes()`

UnsetAttributes ensures that no value is present for Attributes, not even an explicit nil
### GetBody

`func (o *ChatV1ServiceChannelMessage) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *ChatV1ServiceChannelMessage) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *ChatV1ServiceChannelMessage) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *ChatV1ServiceChannelMessage) HasBody() bool`

HasBody returns a boolean if a field has been set.

### SetBodyNil

`func (o *ChatV1ServiceChannelMessage) SetBodyNil(b bool)`

 SetBodyNil sets the value for Body to be an explicit nil

### UnsetBody
`func (o *ChatV1ServiceChannelMessage) UnsetBody()`

UnsetBody ensures that no value is present for Body, not even an explicit nil
### GetChannelSid

`func (o *ChatV1ServiceChannelMessage) GetChannelSid() string`

GetChannelSid returns the ChannelSid field if non-nil, zero value otherwise.

### GetChannelSidOk

`func (o *ChatV1ServiceChannelMessage) GetChannelSidOk() (*string, bool)`

GetChannelSidOk returns a tuple with the ChannelSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelSid

`func (o *ChatV1ServiceChannelMessage) SetChannelSid(v string)`

SetChannelSid sets ChannelSid field to given value.

### HasChannelSid

`func (o *ChatV1ServiceChannelMessage) HasChannelSid() bool`

HasChannelSid returns a boolean if a field has been set.

### SetChannelSidNil

`func (o *ChatV1ServiceChannelMessage) SetChannelSidNil(b bool)`

 SetChannelSidNil sets the value for ChannelSid to be an explicit nil

### UnsetChannelSid
`func (o *ChatV1ServiceChannelMessage) UnsetChannelSid()`

UnsetChannelSid ensures that no value is present for ChannelSid, not even an explicit nil
### GetDateCreated

`func (o *ChatV1ServiceChannelMessage) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *ChatV1ServiceChannelMessage) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *ChatV1ServiceChannelMessage) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *ChatV1ServiceChannelMessage) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### SetDateCreatedNil

`func (o *ChatV1ServiceChannelMessage) SetDateCreatedNil(b bool)`

 SetDateCreatedNil sets the value for DateCreated to be an explicit nil

### UnsetDateCreated
`func (o *ChatV1ServiceChannelMessage) UnsetDateCreated()`

UnsetDateCreated ensures that no value is present for DateCreated, not even an explicit nil
### GetDateUpdated

`func (o *ChatV1ServiceChannelMessage) GetDateUpdated() time.Time`

GetDateUpdated returns the DateUpdated field if non-nil, zero value otherwise.

### GetDateUpdatedOk

`func (o *ChatV1ServiceChannelMessage) GetDateUpdatedOk() (*time.Time, bool)`

GetDateUpdatedOk returns a tuple with the DateUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateUpdated

`func (o *ChatV1ServiceChannelMessage) SetDateUpdated(v time.Time)`

SetDateUpdated sets DateUpdated field to given value.

### HasDateUpdated

`func (o *ChatV1ServiceChannelMessage) HasDateUpdated() bool`

HasDateUpdated returns a boolean if a field has been set.

### SetDateUpdatedNil

`func (o *ChatV1ServiceChannelMessage) SetDateUpdatedNil(b bool)`

 SetDateUpdatedNil sets the value for DateUpdated to be an explicit nil

### UnsetDateUpdated
`func (o *ChatV1ServiceChannelMessage) UnsetDateUpdated()`

UnsetDateUpdated ensures that no value is present for DateUpdated, not even an explicit nil
### GetFrom

`func (o *ChatV1ServiceChannelMessage) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *ChatV1ServiceChannelMessage) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *ChatV1ServiceChannelMessage) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *ChatV1ServiceChannelMessage) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### SetFromNil

`func (o *ChatV1ServiceChannelMessage) SetFromNil(b bool)`

 SetFromNil sets the value for From to be an explicit nil

### UnsetFrom
`func (o *ChatV1ServiceChannelMessage) UnsetFrom()`

UnsetFrom ensures that no value is present for From, not even an explicit nil
### GetIndex

`func (o *ChatV1ServiceChannelMessage) GetIndex() int32`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *ChatV1ServiceChannelMessage) GetIndexOk() (*int32, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *ChatV1ServiceChannelMessage) SetIndex(v int32)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *ChatV1ServiceChannelMessage) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### SetIndexNil

`func (o *ChatV1ServiceChannelMessage) SetIndexNil(b bool)`

 SetIndexNil sets the value for Index to be an explicit nil

### UnsetIndex
`func (o *ChatV1ServiceChannelMessage) UnsetIndex()`

UnsetIndex ensures that no value is present for Index, not even an explicit nil
### GetServiceSid

`func (o *ChatV1ServiceChannelMessage) GetServiceSid() string`

GetServiceSid returns the ServiceSid field if non-nil, zero value otherwise.

### GetServiceSidOk

`func (o *ChatV1ServiceChannelMessage) GetServiceSidOk() (*string, bool)`

GetServiceSidOk returns a tuple with the ServiceSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceSid

`func (o *ChatV1ServiceChannelMessage) SetServiceSid(v string)`

SetServiceSid sets ServiceSid field to given value.

### HasServiceSid

`func (o *ChatV1ServiceChannelMessage) HasServiceSid() bool`

HasServiceSid returns a boolean if a field has been set.

### SetServiceSidNil

`func (o *ChatV1ServiceChannelMessage) SetServiceSidNil(b bool)`

 SetServiceSidNil sets the value for ServiceSid to be an explicit nil

### UnsetServiceSid
`func (o *ChatV1ServiceChannelMessage) UnsetServiceSid()`

UnsetServiceSid ensures that no value is present for ServiceSid, not even an explicit nil
### GetSid

`func (o *ChatV1ServiceChannelMessage) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *ChatV1ServiceChannelMessage) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *ChatV1ServiceChannelMessage) SetSid(v string)`

SetSid sets Sid field to given value.

### HasSid

`func (o *ChatV1ServiceChannelMessage) HasSid() bool`

HasSid returns a boolean if a field has been set.

### SetSidNil

`func (o *ChatV1ServiceChannelMessage) SetSidNil(b bool)`

 SetSidNil sets the value for Sid to be an explicit nil

### UnsetSid
`func (o *ChatV1ServiceChannelMessage) UnsetSid()`

UnsetSid ensures that no value is present for Sid, not even an explicit nil
### GetTo

`func (o *ChatV1ServiceChannelMessage) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *ChatV1ServiceChannelMessage) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *ChatV1ServiceChannelMessage) SetTo(v string)`

SetTo sets To field to given value.

### HasTo

`func (o *ChatV1ServiceChannelMessage) HasTo() bool`

HasTo returns a boolean if a field has been set.

### SetToNil

`func (o *ChatV1ServiceChannelMessage) SetToNil(b bool)`

 SetToNil sets the value for To to be an explicit nil

### UnsetTo
`func (o *ChatV1ServiceChannelMessage) UnsetTo()`

UnsetTo ensures that no value is present for To, not even an explicit nil
### GetUrl

`func (o *ChatV1ServiceChannelMessage) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ChatV1ServiceChannelMessage) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ChatV1ServiceChannelMessage) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *ChatV1ServiceChannelMessage) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *ChatV1ServiceChannelMessage) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *ChatV1ServiceChannelMessage) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetWasEdited

`func (o *ChatV1ServiceChannelMessage) GetWasEdited() bool`

GetWasEdited returns the WasEdited field if non-nil, zero value otherwise.

### GetWasEditedOk

`func (o *ChatV1ServiceChannelMessage) GetWasEditedOk() (*bool, bool)`

GetWasEditedOk returns a tuple with the WasEdited field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWasEdited

`func (o *ChatV1ServiceChannelMessage) SetWasEdited(v bool)`

SetWasEdited sets WasEdited field to given value.

### HasWasEdited

`func (o *ChatV1ServiceChannelMessage) HasWasEdited() bool`

HasWasEdited returns a boolean if a field has been set.

### SetWasEditedNil

`func (o *ChatV1ServiceChannelMessage) SetWasEditedNil(b bool)`

 SetWasEditedNil sets the value for WasEdited to be an explicit nil

### UnsetWasEdited
`func (o *ChatV1ServiceChannelMessage) UnsetWasEdited()`

UnsetWasEdited ensures that no value is present for WasEdited, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


