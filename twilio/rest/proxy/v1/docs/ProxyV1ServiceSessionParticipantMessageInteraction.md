# ProxyV1ServiceSessionParticipantMessageInteraction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **NullableString** | The SID of the Account that created the resource | [optional] 
**Data** | Pointer to **NullableString** | A JSON string that includes the message body sent to the participant | [optional] 
**DateCreated** | Pointer to **NullableTime** | The ISO 8601 date and time in GMT when the resource was created | [optional] 
**DateUpdated** | Pointer to **NullableTime** | The ISO 8601 date and time in GMT when the resource was last updated | [optional] 
**InboundParticipantSid** | Pointer to **NullableString** | Always empty for Message Interactions | [optional] 
**InboundResourceSid** | Pointer to **NullableString** | Always empty for Message Interactions | [optional] 
**InboundResourceStatus** | Pointer to **NullableString** | Always empty for Message Interactions | [optional] 
**InboundResourceType** | Pointer to **NullableString** | Always empty for Message Interactions | [optional] 
**InboundResourceUrl** | Pointer to **NullableString** | Always empty for Message Interactions | [optional] 
**OutboundParticipantSid** | Pointer to **NullableString** | The SID of the outbound Participant resource | [optional] 
**OutboundResourceSid** | Pointer to **NullableString** | The SID of the outbound Message resource | [optional] 
**OutboundResourceStatus** | Pointer to **NullableString** | The outbound resource status | [optional] 
**OutboundResourceType** | Pointer to **NullableString** | The outbound resource type | [optional] 
**OutboundResourceUrl** | Pointer to **NullableString** | The URL of the Twilio message resource | [optional] 
**ParticipantSid** | Pointer to **NullableString** | The SID of the Participant resource | [optional] 
**ServiceSid** | Pointer to **NullableString** | The SID of the resource&#39;s parent Service | [optional] 
**SessionSid** | Pointer to **NullableString** | The SID of the resource&#39;s parent Session | [optional] 
**Sid** | Pointer to **NullableString** | The unique string that identifies the resource | [optional] 
**Type** | Pointer to **NullableString** | The Type of Message Interaction | [optional] 
**Url** | Pointer to **NullableString** | The absolute URL of the MessageInteraction resource | [optional] 

## Methods

### NewProxyV1ServiceSessionParticipantMessageInteraction

`func NewProxyV1ServiceSessionParticipantMessageInteraction() *ProxyV1ServiceSessionParticipantMessageInteraction`

NewProxyV1ServiceSessionParticipantMessageInteraction instantiates a new ProxyV1ServiceSessionParticipantMessageInteraction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProxyV1ServiceSessionParticipantMessageInteractionWithDefaults

`func NewProxyV1ServiceSessionParticipantMessageInteractionWithDefaults() *ProxyV1ServiceSessionParticipantMessageInteraction`

NewProxyV1ServiceSessionParticipantMessageInteractionWithDefaults instantiates a new ProxyV1ServiceSessionParticipantMessageInteraction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSid

`func (o *ProxyV1ServiceSessionParticipantMessageInteraction) GetAccountSid() string`

GetAccountSid returns the AccountSid field if non-nil, zero value otherwise.

### GetAccountSidOk

`func (o *ProxyV1ServiceSessionParticipantMessageInteraction) GetAccountSidOk() (*string, bool)`

GetAccountSidOk returns a tuple with the AccountSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSid

`func (o *ProxyV1ServiceSessionParticipantMessageInteraction) SetAccountSid(v string)`

SetAccountSid sets AccountSid field to given value.

### HasAccountSid

`func (o *ProxyV1ServiceSessionParticipantMessageInteraction) HasAccountSid() bool`

HasAccountSid returns a boolean if a field has been set.

### SetAccountSidNil

`func (o *ProxyV1ServiceSessionParticipantMessageInteraction) SetAccountSidNil(b bool)`

 SetAccountSidNil sets the value for AccountSid to be an explicit nil

### UnsetAccountSid
`func (o *ProxyV1ServiceSessionParticipantMessageInteraction) UnsetAccountSid()`

UnsetAccountSid ensures that no value is present for AccountSid, not even an explicit nil
### GetData

`func (o *ProxyV1ServiceSessionParticipantMessageInteraction) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ProxyV1ServiceSessionParticipantMessageInteraction) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ProxyV1ServiceSessionParticipantMessageInteraction) SetData(v string)`

SetData sets Data field to given value.

### HasData

`func (o *ProxyV1ServiceSessionParticipantMessageInteraction) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *ProxyV1ServiceSessionParticipantMessageInteraction) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *ProxyV1ServiceSessionParticipantMessageInteraction) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetDateCreated

`func (o *ProxyV1ServiceSessionParticipantMessageInteraction) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *ProxyV1ServiceSessionParticipantMessageInteraction) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *ProxyV1ServiceSessionParticipantMessageInteraction) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *ProxyV1ServiceSessionParticipantMessageInteraction) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### SetDateCreatedNil

`func (o *ProxyV1ServiceSessionParticipantMessageInteraction) SetDateCreatedNil(b bool)`

 SetDateCreatedNil sets the value for DateCreated to be an explicit nil

### UnsetDateCreated
`func (o *ProxyV1ServiceSessionParticipantMessageInteraction) UnsetDateCreated()`

UnsetDateCreated ensures that no value is present for DateCreated, not even an explicit nil
### GetDateUpdated

`func (o *ProxyV1ServiceSessionParticipantMessageInteraction) GetDateUpdated() time.Time`

GetDateUpdated returns the DateUpdated field if non-nil, zero value otherwise.

### GetDateUpdatedOk

`func (o *ProxyV1ServiceSessionParticipantMessageInteraction) GetDateUpdatedOk() (*time.Time, bool)`

GetDateUpdatedOk returns a tuple with the DateUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateUpdated

`func (o *ProxyV1ServiceSessionParticipantMessageInteraction) SetDateUpdated(v time.Time)`

SetDateUpdated sets DateUpdated field to given value.

### HasDateUpdated

`func (o *ProxyV1ServiceSessionParticipantMessageInteraction) HasDateUpdated() bool`

HasDateUpdated returns a boolean if a field has been set.

### SetDateUpdatedNil

`func (o *ProxyV1ServiceSessionParticipantMessageInteraction) SetDateUpdatedNil(b bool)`

 SetDateUpdatedNil sets the value for DateUpdated to be an explicit nil

### UnsetDateUpdated
`func (o *ProxyV1ServiceSessionParticipantMessageInteraction) UnsetDateUpdated()`

UnsetDateUpdated ensures that no value is present for DateUpdated, not even an explicit nil
### GetInboundParticipantSid

`func (o *ProxyV1ServiceSessionParticipantMessageInteraction) GetInboundParticipantSid() string`

GetInboundParticipantSid returns the InboundParticipantSid field if non-nil, zero value otherwise.

### GetInboundParticipantSidOk

`func (o *ProxyV1ServiceSessionParticipantMessageInteraction) GetInboundParticipantSidOk() (*string, bool)`

GetInboundParticipantSidOk returns a tuple with the InboundParticipantSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInboundParticipantSid

`func (o *ProxyV1ServiceSessionParticipantMessageInteraction) SetInboundParticipantSid(v string)`

SetInboundParticipantSid sets InboundParticipantSid field to given value.

### HasInboundParticipantSid

`func (o *ProxyV1ServiceSessionParticipantMessageInteraction) HasInboundParticipantSid() bool`

HasInboundParticipantSid returns a boolean if a field has been set.

### SetInboundParticipantSidNil

`func (o *ProxyV1ServiceSessionParticipantMessageInteraction) SetInboundParticipantSidNil(b bool)`

 SetInboundParticipantSidNil sets the value for InboundParticipantSid to be an explicit nil

### UnsetInboundParticipantSid
`func (o *ProxyV1ServiceSessionParticipantMessageInteraction) UnsetInboundParticipantSid()`

UnsetInboundParticipantSid ensures that no value is present for InboundParticipantSid, not even an explicit nil
### GetInboundResourceSid

`func (o *ProxyV1ServiceSessionParticipantMessageInteraction) GetInboundResourceSid() string`

GetInboundResourceSid returns the InboundResourceSid field if non-nil, zero value otherwise.

### GetInboundResourceSidOk

`func (o *ProxyV1ServiceSessionParticipantMessageInteraction) GetInboundResourceSidOk() (*string, bool)`

GetInboundResourceSidOk returns a tuple with the InboundResourceSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInboundResourceSid

`func (o *ProxyV1ServiceSessionParticipantMessageInteraction) SetInboundResourceSid(v string)`

SetInboundResourceSid sets InboundResourceSid field to given value.

### HasInboundResourceSid

`func (o *ProxyV1ServiceSessionParticipantMessageInteraction) HasInboundResourceSid() bool`

HasInboundResourceSid returns a boolean if a field has been set.

### SetInboundResourceSidNil

`func (o *ProxyV1ServiceSessionParticipantMessageInteraction) SetInboundResourceSidNil(b bool)`

 SetInboundResourceSidNil sets the value for InboundResourceSid to be an explicit nil

### UnsetInboundResourceSid
`func (o *ProxyV1ServiceSessionParticipantMessageInteraction) UnsetInboundResourceSid()`

UnsetInboundResourceSid ensures that no value is present for InboundResourceSid, not even an explicit nil
### GetInboundResourceStatus

`func (o *ProxyV1ServiceSessionParticipantMessageInteraction) GetInboundResourceStatus() string`

GetInboundResourceStatus returns the InboundResourceStatus field if non-nil, zero value otherwise.

### GetInboundResourceStatusOk

`func (o *ProxyV1ServiceSessionParticipantMessageInteraction) GetInboundResourceStatusOk() (*string, bool)`

GetInboundResourceStatusOk returns a tuple with the InboundResourceStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInboundResourceStatus

`func (o *ProxyV1ServiceSessionParticipantMessageInteraction) SetInboundResourceStatus(v string)`

SetInboundResourceStatus sets InboundResourceStatus field to given value.

### HasInboundResourceStatus

`func (o *ProxyV1ServiceSessionParticipantMessageInteraction) HasInboundResourceStatus() bool`

HasInboundResourceStatus returns a boolean if a field has been set.

### SetInboundResourceStatusNil

`func (o *ProxyV1ServiceSessionParticipantMessageInteraction) SetInboundResourceStatusNil(b bool)`

 SetInboundResourceStatusNil sets the value for InboundResourceStatus to be an explicit nil

### UnsetInboundResourceStatus
`func (o *ProxyV1ServiceSessionParticipantMessageInteraction) UnsetInboundResourceStatus()`

UnsetInboundResourceStatus ensures that no value is present for InboundResourceStatus, not even an explicit nil
### GetInboundResourceType

`func (o *ProxyV1ServiceSessionParticipantMessageInteraction) GetInboundResourceType() string`

GetInboundResourceType returns the InboundResourceType field if non-nil, zero value otherwise.

### GetInboundResourceTypeOk

`func (o *ProxyV1ServiceSessionParticipantMessageInteraction) GetInboundResourceTypeOk() (*string, bool)`

GetInboundResourceTypeOk returns a tuple with the InboundResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInboundResourceType

`func (o *ProxyV1ServiceSessionParticipantMessageInteraction) SetInboundResourceType(v string)`

SetInboundResourceType sets InboundResourceType field to given value.

### HasInboundResourceType

`func (o *ProxyV1ServiceSessionParticipantMessageInteraction) HasInboundResourceType() bool`

HasInboundResourceType returns a boolean if a field has been set.

### SetInboundResourceTypeNil

`func (o *ProxyV1ServiceSessionParticipantMessageInteraction) SetInboundResourceTypeNil(b bool)`

 SetInboundResourceTypeNil sets the value for InboundResourceType to be an explicit nil

### UnsetInboundResourceType
`func (o *ProxyV1ServiceSessionParticipantMessageInteraction) UnsetInboundResourceType()`

UnsetInboundResourceType ensures that no value is present for InboundResourceType, not even an explicit nil
### GetInboundResourceUrl

`func (o *ProxyV1ServiceSessionParticipantMessageInteraction) GetInboundResourceUrl() string`

GetInboundResourceUrl returns the InboundResourceUrl field if non-nil, zero value otherwise.

### GetInboundResourceUrlOk

`func (o *ProxyV1ServiceSessionParticipantMessageInteraction) GetInboundResourceUrlOk() (*string, bool)`

GetInboundResourceUrlOk returns a tuple with the InboundResourceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInboundResourceUrl

`func (o *ProxyV1ServiceSessionParticipantMessageInteraction) SetInboundResourceUrl(v string)`

SetInboundResourceUrl sets InboundResourceUrl field to given value.

### HasInboundResourceUrl

`func (o *ProxyV1ServiceSessionParticipantMessageInteraction) HasInboundResourceUrl() bool`

HasInboundResourceUrl returns a boolean if a field has been set.

### SetInboundResourceUrlNil

`func (o *ProxyV1ServiceSessionParticipantMessageInteraction) SetInboundResourceUrlNil(b bool)`

 SetInboundResourceUrlNil sets the value for InboundResourceUrl to be an explicit nil

### UnsetInboundResourceUrl
`func (o *ProxyV1ServiceSessionParticipantMessageInteraction) UnsetInboundResourceUrl()`

UnsetInboundResourceUrl ensures that no value is present for InboundResourceUrl, not even an explicit nil
### GetOutboundParticipantSid

`func (o *ProxyV1ServiceSessionParticipantMessageInteraction) GetOutboundParticipantSid() string`

GetOutboundParticipantSid returns the OutboundParticipantSid field if non-nil, zero value otherwise.

### GetOutboundParticipantSidOk

`func (o *ProxyV1ServiceSessionParticipantMessageInteraction) GetOutboundParticipantSidOk() (*string, bool)`

GetOutboundParticipantSidOk returns a tuple with the OutboundParticipantSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutboundParticipantSid

`func (o *ProxyV1ServiceSessionParticipantMessageInteraction) SetOutboundParticipantSid(v string)`

SetOutboundParticipantSid sets OutboundParticipantSid field to given value.

### HasOutboundParticipantSid

`func (o *ProxyV1ServiceSessionParticipantMessageInteraction) HasOutboundParticipantSid() bool`

HasOutboundParticipantSid returns a boolean if a field has been set.

### SetOutboundParticipantSidNil

`func (o *ProxyV1ServiceSessionParticipantMessageInteraction) SetOutboundParticipantSidNil(b bool)`

 SetOutboundParticipantSidNil sets the value for OutboundParticipantSid to be an explicit nil

### UnsetOutboundParticipantSid
`func (o *ProxyV1ServiceSessionParticipantMessageInteraction) UnsetOutboundParticipantSid()`

UnsetOutboundParticipantSid ensures that no value is present for OutboundParticipantSid, not even an explicit nil
### GetOutboundResourceSid

`func (o *ProxyV1ServiceSessionParticipantMessageInteraction) GetOutboundResourceSid() string`

GetOutboundResourceSid returns the OutboundResourceSid field if non-nil, zero value otherwise.

### GetOutboundResourceSidOk

`func (o *ProxyV1ServiceSessionParticipantMessageInteraction) GetOutboundResourceSidOk() (*string, bool)`

GetOutboundResourceSidOk returns a tuple with the OutboundResourceSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutboundResourceSid

`func (o *ProxyV1ServiceSessionParticipantMessageInteraction) SetOutboundResourceSid(v string)`

SetOutboundResourceSid sets OutboundResourceSid field to given value.

### HasOutboundResourceSid

`func (o *ProxyV1ServiceSessionParticipantMessageInteraction) HasOutboundResourceSid() bool`

HasOutboundResourceSid returns a boolean if a field has been set.

### SetOutboundResourceSidNil

`func (o *ProxyV1ServiceSessionParticipantMessageInteraction) SetOutboundResourceSidNil(b bool)`

 SetOutboundResourceSidNil sets the value for OutboundResourceSid to be an explicit nil

### UnsetOutboundResourceSid
`func (o *ProxyV1ServiceSessionParticipantMessageInteraction) UnsetOutboundResourceSid()`

UnsetOutboundResourceSid ensures that no value is present for OutboundResourceSid, not even an explicit nil
### GetOutboundResourceStatus

`func (o *ProxyV1ServiceSessionParticipantMessageInteraction) GetOutboundResourceStatus() string`

GetOutboundResourceStatus returns the OutboundResourceStatus field if non-nil, zero value otherwise.

### GetOutboundResourceStatusOk

`func (o *ProxyV1ServiceSessionParticipantMessageInteraction) GetOutboundResourceStatusOk() (*string, bool)`

GetOutboundResourceStatusOk returns a tuple with the OutboundResourceStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutboundResourceStatus

`func (o *ProxyV1ServiceSessionParticipantMessageInteraction) SetOutboundResourceStatus(v string)`

SetOutboundResourceStatus sets OutboundResourceStatus field to given value.

### HasOutboundResourceStatus

`func (o *ProxyV1ServiceSessionParticipantMessageInteraction) HasOutboundResourceStatus() bool`

HasOutboundResourceStatus returns a boolean if a field has been set.

### SetOutboundResourceStatusNil

`func (o *ProxyV1ServiceSessionParticipantMessageInteraction) SetOutboundResourceStatusNil(b bool)`

 SetOutboundResourceStatusNil sets the value for OutboundResourceStatus to be an explicit nil

### UnsetOutboundResourceStatus
`func (o *ProxyV1ServiceSessionParticipantMessageInteraction) UnsetOutboundResourceStatus()`

UnsetOutboundResourceStatus ensures that no value is present for OutboundResourceStatus, not even an explicit nil
### GetOutboundResourceType

`func (o *ProxyV1ServiceSessionParticipantMessageInteraction) GetOutboundResourceType() string`

GetOutboundResourceType returns the OutboundResourceType field if non-nil, zero value otherwise.

### GetOutboundResourceTypeOk

`func (o *ProxyV1ServiceSessionParticipantMessageInteraction) GetOutboundResourceTypeOk() (*string, bool)`

GetOutboundResourceTypeOk returns a tuple with the OutboundResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutboundResourceType

`func (o *ProxyV1ServiceSessionParticipantMessageInteraction) SetOutboundResourceType(v string)`

SetOutboundResourceType sets OutboundResourceType field to given value.

### HasOutboundResourceType

`func (o *ProxyV1ServiceSessionParticipantMessageInteraction) HasOutboundResourceType() bool`

HasOutboundResourceType returns a boolean if a field has been set.

### SetOutboundResourceTypeNil

`func (o *ProxyV1ServiceSessionParticipantMessageInteraction) SetOutboundResourceTypeNil(b bool)`

 SetOutboundResourceTypeNil sets the value for OutboundResourceType to be an explicit nil

### UnsetOutboundResourceType
`func (o *ProxyV1ServiceSessionParticipantMessageInteraction) UnsetOutboundResourceType()`

UnsetOutboundResourceType ensures that no value is present for OutboundResourceType, not even an explicit nil
### GetOutboundResourceUrl

`func (o *ProxyV1ServiceSessionParticipantMessageInteraction) GetOutboundResourceUrl() string`

GetOutboundResourceUrl returns the OutboundResourceUrl field if non-nil, zero value otherwise.

### GetOutboundResourceUrlOk

`func (o *ProxyV1ServiceSessionParticipantMessageInteraction) GetOutboundResourceUrlOk() (*string, bool)`

GetOutboundResourceUrlOk returns a tuple with the OutboundResourceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutboundResourceUrl

`func (o *ProxyV1ServiceSessionParticipantMessageInteraction) SetOutboundResourceUrl(v string)`

SetOutboundResourceUrl sets OutboundResourceUrl field to given value.

### HasOutboundResourceUrl

`func (o *ProxyV1ServiceSessionParticipantMessageInteraction) HasOutboundResourceUrl() bool`

HasOutboundResourceUrl returns a boolean if a field has been set.

### SetOutboundResourceUrlNil

`func (o *ProxyV1ServiceSessionParticipantMessageInteraction) SetOutboundResourceUrlNil(b bool)`

 SetOutboundResourceUrlNil sets the value for OutboundResourceUrl to be an explicit nil

### UnsetOutboundResourceUrl
`func (o *ProxyV1ServiceSessionParticipantMessageInteraction) UnsetOutboundResourceUrl()`

UnsetOutboundResourceUrl ensures that no value is present for OutboundResourceUrl, not even an explicit nil
### GetParticipantSid

`func (o *ProxyV1ServiceSessionParticipantMessageInteraction) GetParticipantSid() string`

GetParticipantSid returns the ParticipantSid field if non-nil, zero value otherwise.

### GetParticipantSidOk

`func (o *ProxyV1ServiceSessionParticipantMessageInteraction) GetParticipantSidOk() (*string, bool)`

GetParticipantSidOk returns a tuple with the ParticipantSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParticipantSid

`func (o *ProxyV1ServiceSessionParticipantMessageInteraction) SetParticipantSid(v string)`

SetParticipantSid sets ParticipantSid field to given value.

### HasParticipantSid

`func (o *ProxyV1ServiceSessionParticipantMessageInteraction) HasParticipantSid() bool`

HasParticipantSid returns a boolean if a field has been set.

### SetParticipantSidNil

`func (o *ProxyV1ServiceSessionParticipantMessageInteraction) SetParticipantSidNil(b bool)`

 SetParticipantSidNil sets the value for ParticipantSid to be an explicit nil

### UnsetParticipantSid
`func (o *ProxyV1ServiceSessionParticipantMessageInteraction) UnsetParticipantSid()`

UnsetParticipantSid ensures that no value is present for ParticipantSid, not even an explicit nil
### GetServiceSid

`func (o *ProxyV1ServiceSessionParticipantMessageInteraction) GetServiceSid() string`

GetServiceSid returns the ServiceSid field if non-nil, zero value otherwise.

### GetServiceSidOk

`func (o *ProxyV1ServiceSessionParticipantMessageInteraction) GetServiceSidOk() (*string, bool)`

GetServiceSidOk returns a tuple with the ServiceSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceSid

`func (o *ProxyV1ServiceSessionParticipantMessageInteraction) SetServiceSid(v string)`

SetServiceSid sets ServiceSid field to given value.

### HasServiceSid

`func (o *ProxyV1ServiceSessionParticipantMessageInteraction) HasServiceSid() bool`

HasServiceSid returns a boolean if a field has been set.

### SetServiceSidNil

`func (o *ProxyV1ServiceSessionParticipantMessageInteraction) SetServiceSidNil(b bool)`

 SetServiceSidNil sets the value for ServiceSid to be an explicit nil

### UnsetServiceSid
`func (o *ProxyV1ServiceSessionParticipantMessageInteraction) UnsetServiceSid()`

UnsetServiceSid ensures that no value is present for ServiceSid, not even an explicit nil
### GetSessionSid

`func (o *ProxyV1ServiceSessionParticipantMessageInteraction) GetSessionSid() string`

GetSessionSid returns the SessionSid field if non-nil, zero value otherwise.

### GetSessionSidOk

`func (o *ProxyV1ServiceSessionParticipantMessageInteraction) GetSessionSidOk() (*string, bool)`

GetSessionSidOk returns a tuple with the SessionSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionSid

`func (o *ProxyV1ServiceSessionParticipantMessageInteraction) SetSessionSid(v string)`

SetSessionSid sets SessionSid field to given value.

### HasSessionSid

`func (o *ProxyV1ServiceSessionParticipantMessageInteraction) HasSessionSid() bool`

HasSessionSid returns a boolean if a field has been set.

### SetSessionSidNil

`func (o *ProxyV1ServiceSessionParticipantMessageInteraction) SetSessionSidNil(b bool)`

 SetSessionSidNil sets the value for SessionSid to be an explicit nil

### UnsetSessionSid
`func (o *ProxyV1ServiceSessionParticipantMessageInteraction) UnsetSessionSid()`

UnsetSessionSid ensures that no value is present for SessionSid, not even an explicit nil
### GetSid

`func (o *ProxyV1ServiceSessionParticipantMessageInteraction) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *ProxyV1ServiceSessionParticipantMessageInteraction) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *ProxyV1ServiceSessionParticipantMessageInteraction) SetSid(v string)`

SetSid sets Sid field to given value.

### HasSid

`func (o *ProxyV1ServiceSessionParticipantMessageInteraction) HasSid() bool`

HasSid returns a boolean if a field has been set.

### SetSidNil

`func (o *ProxyV1ServiceSessionParticipantMessageInteraction) SetSidNil(b bool)`

 SetSidNil sets the value for Sid to be an explicit nil

### UnsetSid
`func (o *ProxyV1ServiceSessionParticipantMessageInteraction) UnsetSid()`

UnsetSid ensures that no value is present for Sid, not even an explicit nil
### GetType

`func (o *ProxyV1ServiceSessionParticipantMessageInteraction) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ProxyV1ServiceSessionParticipantMessageInteraction) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ProxyV1ServiceSessionParticipantMessageInteraction) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ProxyV1ServiceSessionParticipantMessageInteraction) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *ProxyV1ServiceSessionParticipantMessageInteraction) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *ProxyV1ServiceSessionParticipantMessageInteraction) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetUrl

`func (o *ProxyV1ServiceSessionParticipantMessageInteraction) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ProxyV1ServiceSessionParticipantMessageInteraction) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ProxyV1ServiceSessionParticipantMessageInteraction) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *ProxyV1ServiceSessionParticipantMessageInteraction) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *ProxyV1ServiceSessionParticipantMessageInteraction) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *ProxyV1ServiceSessionParticipantMessageInteraction) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


