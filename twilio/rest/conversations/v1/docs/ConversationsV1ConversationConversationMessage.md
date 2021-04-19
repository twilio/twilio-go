# ConversationsV1ConversationConversationMessage

## Properties

Name | Type | Description
------------ | ------------- | -------------
**AccountSid** | Pointer to **NullableString** | The unique ID of the Account responsible for this message. | [optional] 
**Attributes** | Pointer to **NullableString** | A string metadata field you can use to store any data you wish. | [optional] 
**Author** | Pointer to **NullableString** | The channel specific identifier of the message&#39;s author. | [optional] 
**Body** | Pointer to **NullableString** | The content of the message. | [optional] 
**ConversationSid** | Pointer to **NullableString** | The unique ID of the Conversation for this message. | [optional] 
**DateCreated** | Pointer to **NullableTime** | The date that this resource was created. | [optional] 
**DateUpdated** | Pointer to **NullableTime** | The date that this resource was last updated. | [optional] 
**Delivery** | Pointer to **map[string]interface{}** | An object that contains the summary of delivery statuses for the message to non-chat participants. | [optional] 
**Index** | Pointer to **NullableInt32** | The index of the message within the Conversation. | [optional] 
**Links** | Pointer to **map[string]interface{}** | Absolute URL to access the receipts of this message. | [optional] 
**Media** | Pointer to **[]map[string]interface{}** | An array of objects that describe the Message&#39;s media if attached, otherwise, null. | [optional] 
**ParticipantSid** | Pointer to **NullableString** | The unique ID of messages&#39;s author participant. | [optional] 
**Sid** | Pointer to **NullableString** | A 34 character string that uniquely identifies this resource. | [optional] 
**Url** | Pointer to **NullableString** | An absolute API URL for this message. | [optional] 

## Methods

### NewConversationsV1ConversationConversationMessage

`func NewConversationsV1ConversationConversationMessage() *ConversationsV1ConversationConversationMessage`

NewConversationsV1ConversationConversationMessage instantiates a new ConversationsV1ConversationConversationMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConversationsV1ConversationConversationMessageWithDefaults

`func NewConversationsV1ConversationConversationMessageWithDefaults() *ConversationsV1ConversationConversationMessage`

NewConversationsV1ConversationConversationMessageWithDefaults instantiates a new ConversationsV1ConversationConversationMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSid

`func (o *ConversationsV1ConversationConversationMessage) GetAccountSid() string`

GetAccountSid returns the AccountSid field if non-nil, zero value otherwise.

### GetAccountSidOk

`func (o *ConversationsV1ConversationConversationMessage) GetAccountSidOk() (*string, bool)`

GetAccountSidOk returns a tuple with the AccountSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSid

`func (o *ConversationsV1ConversationConversationMessage) SetAccountSid(v string)`

SetAccountSid sets AccountSid field to given value.

### HasAccountSid

`func (o *ConversationsV1ConversationConversationMessage) HasAccountSid() bool`

HasAccountSid returns a boolean if a field has been set.

### SetAccountSidNil

`func (o *ConversationsV1ConversationConversationMessage) SetAccountSidNil(b bool)`

 SetAccountSidNil sets the value for AccountSid to be an explicit nil

### UnsetAccountSid
`func (o *ConversationsV1ConversationConversationMessage) UnsetAccountSid()`

UnsetAccountSid ensures that no value is present for AccountSid, not even an explicit nil
### GetAttributes

`func (o *ConversationsV1ConversationConversationMessage) GetAttributes() string`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *ConversationsV1ConversationConversationMessage) GetAttributesOk() (*string, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *ConversationsV1ConversationConversationMessage) SetAttributes(v string)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *ConversationsV1ConversationConversationMessage) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### SetAttributesNil

`func (o *ConversationsV1ConversationConversationMessage) SetAttributesNil(b bool)`

 SetAttributesNil sets the value for Attributes to be an explicit nil

### UnsetAttributes
`func (o *ConversationsV1ConversationConversationMessage) UnsetAttributes()`

UnsetAttributes ensures that no value is present for Attributes, not even an explicit nil
### GetAuthor

`func (o *ConversationsV1ConversationConversationMessage) GetAuthor() string`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *ConversationsV1ConversationConversationMessage) GetAuthorOk() (*string, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *ConversationsV1ConversationConversationMessage) SetAuthor(v string)`

SetAuthor sets Author field to given value.

### HasAuthor

`func (o *ConversationsV1ConversationConversationMessage) HasAuthor() bool`

HasAuthor returns a boolean if a field has been set.

### SetAuthorNil

`func (o *ConversationsV1ConversationConversationMessage) SetAuthorNil(b bool)`

 SetAuthorNil sets the value for Author to be an explicit nil

### UnsetAuthor
`func (o *ConversationsV1ConversationConversationMessage) UnsetAuthor()`

UnsetAuthor ensures that no value is present for Author, not even an explicit nil
### GetBody

`func (o *ConversationsV1ConversationConversationMessage) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *ConversationsV1ConversationConversationMessage) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *ConversationsV1ConversationConversationMessage) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *ConversationsV1ConversationConversationMessage) HasBody() bool`

HasBody returns a boolean if a field has been set.

### SetBodyNil

`func (o *ConversationsV1ConversationConversationMessage) SetBodyNil(b bool)`

 SetBodyNil sets the value for Body to be an explicit nil

### UnsetBody
`func (o *ConversationsV1ConversationConversationMessage) UnsetBody()`

UnsetBody ensures that no value is present for Body, not even an explicit nil
### GetConversationSid

`func (o *ConversationsV1ConversationConversationMessage) GetConversationSid() string`

GetConversationSid returns the ConversationSid field if non-nil, zero value otherwise.

### GetConversationSidOk

`func (o *ConversationsV1ConversationConversationMessage) GetConversationSidOk() (*string, bool)`

GetConversationSidOk returns a tuple with the ConversationSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversationSid

`func (o *ConversationsV1ConversationConversationMessage) SetConversationSid(v string)`

SetConversationSid sets ConversationSid field to given value.

### HasConversationSid

`func (o *ConversationsV1ConversationConversationMessage) HasConversationSid() bool`

HasConversationSid returns a boolean if a field has been set.

### SetConversationSidNil

`func (o *ConversationsV1ConversationConversationMessage) SetConversationSidNil(b bool)`

 SetConversationSidNil sets the value for ConversationSid to be an explicit nil

### UnsetConversationSid
`func (o *ConversationsV1ConversationConversationMessage) UnsetConversationSid()`

UnsetConversationSid ensures that no value is present for ConversationSid, not even an explicit nil
### GetDateCreated

`func (o *ConversationsV1ConversationConversationMessage) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *ConversationsV1ConversationConversationMessage) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *ConversationsV1ConversationConversationMessage) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *ConversationsV1ConversationConversationMessage) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### SetDateCreatedNil

`func (o *ConversationsV1ConversationConversationMessage) SetDateCreatedNil(b bool)`

 SetDateCreatedNil sets the value for DateCreated to be an explicit nil

### UnsetDateCreated
`func (o *ConversationsV1ConversationConversationMessage) UnsetDateCreated()`

UnsetDateCreated ensures that no value is present for DateCreated, not even an explicit nil
### GetDateUpdated

`func (o *ConversationsV1ConversationConversationMessage) GetDateUpdated() time.Time`

GetDateUpdated returns the DateUpdated field if non-nil, zero value otherwise.

### GetDateUpdatedOk

`func (o *ConversationsV1ConversationConversationMessage) GetDateUpdatedOk() (*time.Time, bool)`

GetDateUpdatedOk returns a tuple with the DateUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateUpdated

`func (o *ConversationsV1ConversationConversationMessage) SetDateUpdated(v time.Time)`

SetDateUpdated sets DateUpdated field to given value.

### HasDateUpdated

`func (o *ConversationsV1ConversationConversationMessage) HasDateUpdated() bool`

HasDateUpdated returns a boolean if a field has been set.

### SetDateUpdatedNil

`func (o *ConversationsV1ConversationConversationMessage) SetDateUpdatedNil(b bool)`

 SetDateUpdatedNil sets the value for DateUpdated to be an explicit nil

### UnsetDateUpdated
`func (o *ConversationsV1ConversationConversationMessage) UnsetDateUpdated()`

UnsetDateUpdated ensures that no value is present for DateUpdated, not even an explicit nil
### GetDelivery

`func (o *ConversationsV1ConversationConversationMessage) GetDelivery() map[string]interface{}`

GetDelivery returns the Delivery field if non-nil, zero value otherwise.

### GetDeliveryOk

`func (o *ConversationsV1ConversationConversationMessage) GetDeliveryOk() (*map[string]interface{}, bool)`

GetDeliveryOk returns a tuple with the Delivery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelivery

`func (o *ConversationsV1ConversationConversationMessage) SetDelivery(v map[string]interface{})`

SetDelivery sets Delivery field to given value.

### HasDelivery

`func (o *ConversationsV1ConversationConversationMessage) HasDelivery() bool`

HasDelivery returns a boolean if a field has been set.

### SetDeliveryNil

`func (o *ConversationsV1ConversationConversationMessage) SetDeliveryNil(b bool)`

 SetDeliveryNil sets the value for Delivery to be an explicit nil

### UnsetDelivery
`func (o *ConversationsV1ConversationConversationMessage) UnsetDelivery()`

UnsetDelivery ensures that no value is present for Delivery, not even an explicit nil
### GetIndex

`func (o *ConversationsV1ConversationConversationMessage) GetIndex() int32`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *ConversationsV1ConversationConversationMessage) GetIndexOk() (*int32, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *ConversationsV1ConversationConversationMessage) SetIndex(v int32)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *ConversationsV1ConversationConversationMessage) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### SetIndexNil

`func (o *ConversationsV1ConversationConversationMessage) SetIndexNil(b bool)`

 SetIndexNil sets the value for Index to be an explicit nil

### UnsetIndex
`func (o *ConversationsV1ConversationConversationMessage) UnsetIndex()`

UnsetIndex ensures that no value is present for Index, not even an explicit nil
### GetLinks

`func (o *ConversationsV1ConversationConversationMessage) GetLinks() map[string]interface{}`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ConversationsV1ConversationConversationMessage) GetLinksOk() (*map[string]interface{}, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ConversationsV1ConversationConversationMessage) SetLinks(v map[string]interface{})`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ConversationsV1ConversationConversationMessage) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### SetLinksNil

`func (o *ConversationsV1ConversationConversationMessage) SetLinksNil(b bool)`

 SetLinksNil sets the value for Links to be an explicit nil

### UnsetLinks
`func (o *ConversationsV1ConversationConversationMessage) UnsetLinks()`

UnsetLinks ensures that no value is present for Links, not even an explicit nil
### GetMedia

`func (o *ConversationsV1ConversationConversationMessage) GetMedia() []map[string]interface{}`

GetMedia returns the Media field if non-nil, zero value otherwise.

### GetMediaOk

`func (o *ConversationsV1ConversationConversationMessage) GetMediaOk() (*[]map[string]interface{}, bool)`

GetMediaOk returns a tuple with the Media field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMedia

`func (o *ConversationsV1ConversationConversationMessage) SetMedia(v []map[string]interface{})`

SetMedia sets Media field to given value.

### HasMedia

`func (o *ConversationsV1ConversationConversationMessage) HasMedia() bool`

HasMedia returns a boolean if a field has been set.

### SetMediaNil

`func (o *ConversationsV1ConversationConversationMessage) SetMediaNil(b bool)`

 SetMediaNil sets the value for Media to be an explicit nil

### UnsetMedia
`func (o *ConversationsV1ConversationConversationMessage) UnsetMedia()`

UnsetMedia ensures that no value is present for Media, not even an explicit nil
### GetParticipantSid

`func (o *ConversationsV1ConversationConversationMessage) GetParticipantSid() string`

GetParticipantSid returns the ParticipantSid field if non-nil, zero value otherwise.

### GetParticipantSidOk

`func (o *ConversationsV1ConversationConversationMessage) GetParticipantSidOk() (*string, bool)`

GetParticipantSidOk returns a tuple with the ParticipantSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParticipantSid

`func (o *ConversationsV1ConversationConversationMessage) SetParticipantSid(v string)`

SetParticipantSid sets ParticipantSid field to given value.

### HasParticipantSid

`func (o *ConversationsV1ConversationConversationMessage) HasParticipantSid() bool`

HasParticipantSid returns a boolean if a field has been set.

### SetParticipantSidNil

`func (o *ConversationsV1ConversationConversationMessage) SetParticipantSidNil(b bool)`

 SetParticipantSidNil sets the value for ParticipantSid to be an explicit nil

### UnsetParticipantSid
`func (o *ConversationsV1ConversationConversationMessage) UnsetParticipantSid()`

UnsetParticipantSid ensures that no value is present for ParticipantSid, not even an explicit nil
### GetSid

`func (o *ConversationsV1ConversationConversationMessage) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *ConversationsV1ConversationConversationMessage) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *ConversationsV1ConversationConversationMessage) SetSid(v string)`

SetSid sets Sid field to given value.

### HasSid

`func (o *ConversationsV1ConversationConversationMessage) HasSid() bool`

HasSid returns a boolean if a field has been set.

### SetSidNil

`func (o *ConversationsV1ConversationConversationMessage) SetSidNil(b bool)`

 SetSidNil sets the value for Sid to be an explicit nil

### UnsetSid
`func (o *ConversationsV1ConversationConversationMessage) UnsetSid()`

UnsetSid ensures that no value is present for Sid, not even an explicit nil
### GetUrl

`func (o *ConversationsV1ConversationConversationMessage) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ConversationsV1ConversationConversationMessage) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ConversationsV1ConversationConversationMessage) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *ConversationsV1ConversationConversationMessage) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *ConversationsV1ConversationConversationMessage) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *ConversationsV1ConversationConversationMessage) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


