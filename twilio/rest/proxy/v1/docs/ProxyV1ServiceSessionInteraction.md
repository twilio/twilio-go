# ProxyV1ServiceSessionInteraction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **NullableString** | The SID of the Account that created the resource | [optional] 
**Data** | Pointer to **NullableString** | A JSON string that includes the message body of message interactions | [optional] 
**DateCreated** | Pointer to **NullableTime** | The ISO 8601 date and time in GMT when the Interaction was created | [optional] 
**DateUpdated** | Pointer to **NullableTime** | The ISO 8601 date and time in GMT when the resource was last updated | [optional] 
**InboundParticipantSid** | Pointer to **NullableString** | The SID of the inbound Participant resource | [optional] 
**InboundResourceSid** | Pointer to **NullableString** | The SID of the inbound resource | [optional] 
**InboundResourceStatus** | Pointer to **NullableString** | The inbound resource status of the Interaction | [optional] 
**InboundResourceType** | Pointer to **NullableString** | The inbound resource type | [optional] 
**InboundResourceUrl** | Pointer to **NullableString** | The URL of the Twilio inbound resource | [optional] 
**OutboundParticipantSid** | Pointer to **NullableString** | The SID of the outbound Participant | [optional] 
**OutboundResourceSid** | Pointer to **NullableString** | The SID of the outbound resource | [optional] 
**OutboundResourceStatus** | Pointer to **NullableString** | The outbound resource status of the Interaction | [optional] 
**OutboundResourceType** | Pointer to **NullableString** | The outbound resource type | [optional] 
**OutboundResourceUrl** | Pointer to **NullableString** | The URL of the Twilio outbound resource | [optional] 
**ServiceSid** | Pointer to **NullableString** | The SID of the resource&#39;s parent Service | [optional] 
**SessionSid** | Pointer to **NullableString** | The SID of the resource&#39;s parent Session | [optional] 
**Sid** | Pointer to **NullableString** | The unique string that identifies the resource | [optional] 
**Type** | Pointer to **NullableString** | The Type of the Interaction | [optional] 
**Url** | Pointer to **NullableString** | The absolute URL of the Interaction resource | [optional] 

## Methods

### NewProxyV1ServiceSessionInteraction

`func NewProxyV1ServiceSessionInteraction() *ProxyV1ServiceSessionInteraction`

NewProxyV1ServiceSessionInteraction instantiates a new ProxyV1ServiceSessionInteraction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProxyV1ServiceSessionInteractionWithDefaults

`func NewProxyV1ServiceSessionInteractionWithDefaults() *ProxyV1ServiceSessionInteraction`

NewProxyV1ServiceSessionInteractionWithDefaults instantiates a new ProxyV1ServiceSessionInteraction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSid

`func (o *ProxyV1ServiceSessionInteraction) GetAccountSid() string`

GetAccountSid returns the AccountSid field if non-nil, zero value otherwise.

### GetAccountSidOk

`func (o *ProxyV1ServiceSessionInteraction) GetAccountSidOk() (*string, bool)`

GetAccountSidOk returns a tuple with the AccountSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSid

`func (o *ProxyV1ServiceSessionInteraction) SetAccountSid(v string)`

SetAccountSid sets AccountSid field to given value.

### HasAccountSid

`func (o *ProxyV1ServiceSessionInteraction) HasAccountSid() bool`

HasAccountSid returns a boolean if a field has been set.

### SetAccountSidNil

`func (o *ProxyV1ServiceSessionInteraction) SetAccountSidNil(b bool)`

 SetAccountSidNil sets the value for AccountSid to be an explicit nil

### UnsetAccountSid
`func (o *ProxyV1ServiceSessionInteraction) UnsetAccountSid()`

UnsetAccountSid ensures that no value is present for AccountSid, not even an explicit nil
### GetData

`func (o *ProxyV1ServiceSessionInteraction) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ProxyV1ServiceSessionInteraction) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ProxyV1ServiceSessionInteraction) SetData(v string)`

SetData sets Data field to given value.

### HasData

`func (o *ProxyV1ServiceSessionInteraction) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *ProxyV1ServiceSessionInteraction) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *ProxyV1ServiceSessionInteraction) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetDateCreated

`func (o *ProxyV1ServiceSessionInteraction) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *ProxyV1ServiceSessionInteraction) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *ProxyV1ServiceSessionInteraction) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *ProxyV1ServiceSessionInteraction) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### SetDateCreatedNil

`func (o *ProxyV1ServiceSessionInteraction) SetDateCreatedNil(b bool)`

 SetDateCreatedNil sets the value for DateCreated to be an explicit nil

### UnsetDateCreated
`func (o *ProxyV1ServiceSessionInteraction) UnsetDateCreated()`

UnsetDateCreated ensures that no value is present for DateCreated, not even an explicit nil
### GetDateUpdated

`func (o *ProxyV1ServiceSessionInteraction) GetDateUpdated() time.Time`

GetDateUpdated returns the DateUpdated field if non-nil, zero value otherwise.

### GetDateUpdatedOk

`func (o *ProxyV1ServiceSessionInteraction) GetDateUpdatedOk() (*time.Time, bool)`

GetDateUpdatedOk returns a tuple with the DateUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateUpdated

`func (o *ProxyV1ServiceSessionInteraction) SetDateUpdated(v time.Time)`

SetDateUpdated sets DateUpdated field to given value.

### HasDateUpdated

`func (o *ProxyV1ServiceSessionInteraction) HasDateUpdated() bool`

HasDateUpdated returns a boolean if a field has been set.

### SetDateUpdatedNil

`func (o *ProxyV1ServiceSessionInteraction) SetDateUpdatedNil(b bool)`

 SetDateUpdatedNil sets the value for DateUpdated to be an explicit nil

### UnsetDateUpdated
`func (o *ProxyV1ServiceSessionInteraction) UnsetDateUpdated()`

UnsetDateUpdated ensures that no value is present for DateUpdated, not even an explicit nil
### GetInboundParticipantSid

`func (o *ProxyV1ServiceSessionInteraction) GetInboundParticipantSid() string`

GetInboundParticipantSid returns the InboundParticipantSid field if non-nil, zero value otherwise.

### GetInboundParticipantSidOk

`func (o *ProxyV1ServiceSessionInteraction) GetInboundParticipantSidOk() (*string, bool)`

GetInboundParticipantSidOk returns a tuple with the InboundParticipantSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInboundParticipantSid

`func (o *ProxyV1ServiceSessionInteraction) SetInboundParticipantSid(v string)`

SetInboundParticipantSid sets InboundParticipantSid field to given value.

### HasInboundParticipantSid

`func (o *ProxyV1ServiceSessionInteraction) HasInboundParticipantSid() bool`

HasInboundParticipantSid returns a boolean if a field has been set.

### SetInboundParticipantSidNil

`func (o *ProxyV1ServiceSessionInteraction) SetInboundParticipantSidNil(b bool)`

 SetInboundParticipantSidNil sets the value for InboundParticipantSid to be an explicit nil

### UnsetInboundParticipantSid
`func (o *ProxyV1ServiceSessionInteraction) UnsetInboundParticipantSid()`

UnsetInboundParticipantSid ensures that no value is present for InboundParticipantSid, not even an explicit nil
### GetInboundResourceSid

`func (o *ProxyV1ServiceSessionInteraction) GetInboundResourceSid() string`

GetInboundResourceSid returns the InboundResourceSid field if non-nil, zero value otherwise.

### GetInboundResourceSidOk

`func (o *ProxyV1ServiceSessionInteraction) GetInboundResourceSidOk() (*string, bool)`

GetInboundResourceSidOk returns a tuple with the InboundResourceSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInboundResourceSid

`func (o *ProxyV1ServiceSessionInteraction) SetInboundResourceSid(v string)`

SetInboundResourceSid sets InboundResourceSid field to given value.

### HasInboundResourceSid

`func (o *ProxyV1ServiceSessionInteraction) HasInboundResourceSid() bool`

HasInboundResourceSid returns a boolean if a field has been set.

### SetInboundResourceSidNil

`func (o *ProxyV1ServiceSessionInteraction) SetInboundResourceSidNil(b bool)`

 SetInboundResourceSidNil sets the value for InboundResourceSid to be an explicit nil

### UnsetInboundResourceSid
`func (o *ProxyV1ServiceSessionInteraction) UnsetInboundResourceSid()`

UnsetInboundResourceSid ensures that no value is present for InboundResourceSid, not even an explicit nil
### GetInboundResourceStatus

`func (o *ProxyV1ServiceSessionInteraction) GetInboundResourceStatus() string`

GetInboundResourceStatus returns the InboundResourceStatus field if non-nil, zero value otherwise.

### GetInboundResourceStatusOk

`func (o *ProxyV1ServiceSessionInteraction) GetInboundResourceStatusOk() (*string, bool)`

GetInboundResourceStatusOk returns a tuple with the InboundResourceStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInboundResourceStatus

`func (o *ProxyV1ServiceSessionInteraction) SetInboundResourceStatus(v string)`

SetInboundResourceStatus sets InboundResourceStatus field to given value.

### HasInboundResourceStatus

`func (o *ProxyV1ServiceSessionInteraction) HasInboundResourceStatus() bool`

HasInboundResourceStatus returns a boolean if a field has been set.

### SetInboundResourceStatusNil

`func (o *ProxyV1ServiceSessionInteraction) SetInboundResourceStatusNil(b bool)`

 SetInboundResourceStatusNil sets the value for InboundResourceStatus to be an explicit nil

### UnsetInboundResourceStatus
`func (o *ProxyV1ServiceSessionInteraction) UnsetInboundResourceStatus()`

UnsetInboundResourceStatus ensures that no value is present for InboundResourceStatus, not even an explicit nil
### GetInboundResourceType

`func (o *ProxyV1ServiceSessionInteraction) GetInboundResourceType() string`

GetInboundResourceType returns the InboundResourceType field if non-nil, zero value otherwise.

### GetInboundResourceTypeOk

`func (o *ProxyV1ServiceSessionInteraction) GetInboundResourceTypeOk() (*string, bool)`

GetInboundResourceTypeOk returns a tuple with the InboundResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInboundResourceType

`func (o *ProxyV1ServiceSessionInteraction) SetInboundResourceType(v string)`

SetInboundResourceType sets InboundResourceType field to given value.

### HasInboundResourceType

`func (o *ProxyV1ServiceSessionInteraction) HasInboundResourceType() bool`

HasInboundResourceType returns a boolean if a field has been set.

### SetInboundResourceTypeNil

`func (o *ProxyV1ServiceSessionInteraction) SetInboundResourceTypeNil(b bool)`

 SetInboundResourceTypeNil sets the value for InboundResourceType to be an explicit nil

### UnsetInboundResourceType
`func (o *ProxyV1ServiceSessionInteraction) UnsetInboundResourceType()`

UnsetInboundResourceType ensures that no value is present for InboundResourceType, not even an explicit nil
### GetInboundResourceUrl

`func (o *ProxyV1ServiceSessionInteraction) GetInboundResourceUrl() string`

GetInboundResourceUrl returns the InboundResourceUrl field if non-nil, zero value otherwise.

### GetInboundResourceUrlOk

`func (o *ProxyV1ServiceSessionInteraction) GetInboundResourceUrlOk() (*string, bool)`

GetInboundResourceUrlOk returns a tuple with the InboundResourceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInboundResourceUrl

`func (o *ProxyV1ServiceSessionInteraction) SetInboundResourceUrl(v string)`

SetInboundResourceUrl sets InboundResourceUrl field to given value.

### HasInboundResourceUrl

`func (o *ProxyV1ServiceSessionInteraction) HasInboundResourceUrl() bool`

HasInboundResourceUrl returns a boolean if a field has been set.

### SetInboundResourceUrlNil

`func (o *ProxyV1ServiceSessionInteraction) SetInboundResourceUrlNil(b bool)`

 SetInboundResourceUrlNil sets the value for InboundResourceUrl to be an explicit nil

### UnsetInboundResourceUrl
`func (o *ProxyV1ServiceSessionInteraction) UnsetInboundResourceUrl()`

UnsetInboundResourceUrl ensures that no value is present for InboundResourceUrl, not even an explicit nil
### GetOutboundParticipantSid

`func (o *ProxyV1ServiceSessionInteraction) GetOutboundParticipantSid() string`

GetOutboundParticipantSid returns the OutboundParticipantSid field if non-nil, zero value otherwise.

### GetOutboundParticipantSidOk

`func (o *ProxyV1ServiceSessionInteraction) GetOutboundParticipantSidOk() (*string, bool)`

GetOutboundParticipantSidOk returns a tuple with the OutboundParticipantSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutboundParticipantSid

`func (o *ProxyV1ServiceSessionInteraction) SetOutboundParticipantSid(v string)`

SetOutboundParticipantSid sets OutboundParticipantSid field to given value.

### HasOutboundParticipantSid

`func (o *ProxyV1ServiceSessionInteraction) HasOutboundParticipantSid() bool`

HasOutboundParticipantSid returns a boolean if a field has been set.

### SetOutboundParticipantSidNil

`func (o *ProxyV1ServiceSessionInteraction) SetOutboundParticipantSidNil(b bool)`

 SetOutboundParticipantSidNil sets the value for OutboundParticipantSid to be an explicit nil

### UnsetOutboundParticipantSid
`func (o *ProxyV1ServiceSessionInteraction) UnsetOutboundParticipantSid()`

UnsetOutboundParticipantSid ensures that no value is present for OutboundParticipantSid, not even an explicit nil
### GetOutboundResourceSid

`func (o *ProxyV1ServiceSessionInteraction) GetOutboundResourceSid() string`

GetOutboundResourceSid returns the OutboundResourceSid field if non-nil, zero value otherwise.

### GetOutboundResourceSidOk

`func (o *ProxyV1ServiceSessionInteraction) GetOutboundResourceSidOk() (*string, bool)`

GetOutboundResourceSidOk returns a tuple with the OutboundResourceSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutboundResourceSid

`func (o *ProxyV1ServiceSessionInteraction) SetOutboundResourceSid(v string)`

SetOutboundResourceSid sets OutboundResourceSid field to given value.

### HasOutboundResourceSid

`func (o *ProxyV1ServiceSessionInteraction) HasOutboundResourceSid() bool`

HasOutboundResourceSid returns a boolean if a field has been set.

### SetOutboundResourceSidNil

`func (o *ProxyV1ServiceSessionInteraction) SetOutboundResourceSidNil(b bool)`

 SetOutboundResourceSidNil sets the value for OutboundResourceSid to be an explicit nil

### UnsetOutboundResourceSid
`func (o *ProxyV1ServiceSessionInteraction) UnsetOutboundResourceSid()`

UnsetOutboundResourceSid ensures that no value is present for OutboundResourceSid, not even an explicit nil
### GetOutboundResourceStatus

`func (o *ProxyV1ServiceSessionInteraction) GetOutboundResourceStatus() string`

GetOutboundResourceStatus returns the OutboundResourceStatus field if non-nil, zero value otherwise.

### GetOutboundResourceStatusOk

`func (o *ProxyV1ServiceSessionInteraction) GetOutboundResourceStatusOk() (*string, bool)`

GetOutboundResourceStatusOk returns a tuple with the OutboundResourceStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutboundResourceStatus

`func (o *ProxyV1ServiceSessionInteraction) SetOutboundResourceStatus(v string)`

SetOutboundResourceStatus sets OutboundResourceStatus field to given value.

### HasOutboundResourceStatus

`func (o *ProxyV1ServiceSessionInteraction) HasOutboundResourceStatus() bool`

HasOutboundResourceStatus returns a boolean if a field has been set.

### SetOutboundResourceStatusNil

`func (o *ProxyV1ServiceSessionInteraction) SetOutboundResourceStatusNil(b bool)`

 SetOutboundResourceStatusNil sets the value for OutboundResourceStatus to be an explicit nil

### UnsetOutboundResourceStatus
`func (o *ProxyV1ServiceSessionInteraction) UnsetOutboundResourceStatus()`

UnsetOutboundResourceStatus ensures that no value is present for OutboundResourceStatus, not even an explicit nil
### GetOutboundResourceType

`func (o *ProxyV1ServiceSessionInteraction) GetOutboundResourceType() string`

GetOutboundResourceType returns the OutboundResourceType field if non-nil, zero value otherwise.

### GetOutboundResourceTypeOk

`func (o *ProxyV1ServiceSessionInteraction) GetOutboundResourceTypeOk() (*string, bool)`

GetOutboundResourceTypeOk returns a tuple with the OutboundResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutboundResourceType

`func (o *ProxyV1ServiceSessionInteraction) SetOutboundResourceType(v string)`

SetOutboundResourceType sets OutboundResourceType field to given value.

### HasOutboundResourceType

`func (o *ProxyV1ServiceSessionInteraction) HasOutboundResourceType() bool`

HasOutboundResourceType returns a boolean if a field has been set.

### SetOutboundResourceTypeNil

`func (o *ProxyV1ServiceSessionInteraction) SetOutboundResourceTypeNil(b bool)`

 SetOutboundResourceTypeNil sets the value for OutboundResourceType to be an explicit nil

### UnsetOutboundResourceType
`func (o *ProxyV1ServiceSessionInteraction) UnsetOutboundResourceType()`

UnsetOutboundResourceType ensures that no value is present for OutboundResourceType, not even an explicit nil
### GetOutboundResourceUrl

`func (o *ProxyV1ServiceSessionInteraction) GetOutboundResourceUrl() string`

GetOutboundResourceUrl returns the OutboundResourceUrl field if non-nil, zero value otherwise.

### GetOutboundResourceUrlOk

`func (o *ProxyV1ServiceSessionInteraction) GetOutboundResourceUrlOk() (*string, bool)`

GetOutboundResourceUrlOk returns a tuple with the OutboundResourceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutboundResourceUrl

`func (o *ProxyV1ServiceSessionInteraction) SetOutboundResourceUrl(v string)`

SetOutboundResourceUrl sets OutboundResourceUrl field to given value.

### HasOutboundResourceUrl

`func (o *ProxyV1ServiceSessionInteraction) HasOutboundResourceUrl() bool`

HasOutboundResourceUrl returns a boolean if a field has been set.

### SetOutboundResourceUrlNil

`func (o *ProxyV1ServiceSessionInteraction) SetOutboundResourceUrlNil(b bool)`

 SetOutboundResourceUrlNil sets the value for OutboundResourceUrl to be an explicit nil

### UnsetOutboundResourceUrl
`func (o *ProxyV1ServiceSessionInteraction) UnsetOutboundResourceUrl()`

UnsetOutboundResourceUrl ensures that no value is present for OutboundResourceUrl, not even an explicit nil
### GetServiceSid

`func (o *ProxyV1ServiceSessionInteraction) GetServiceSid() string`

GetServiceSid returns the ServiceSid field if non-nil, zero value otherwise.

### GetServiceSidOk

`func (o *ProxyV1ServiceSessionInteraction) GetServiceSidOk() (*string, bool)`

GetServiceSidOk returns a tuple with the ServiceSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceSid

`func (o *ProxyV1ServiceSessionInteraction) SetServiceSid(v string)`

SetServiceSid sets ServiceSid field to given value.

### HasServiceSid

`func (o *ProxyV1ServiceSessionInteraction) HasServiceSid() bool`

HasServiceSid returns a boolean if a field has been set.

### SetServiceSidNil

`func (o *ProxyV1ServiceSessionInteraction) SetServiceSidNil(b bool)`

 SetServiceSidNil sets the value for ServiceSid to be an explicit nil

### UnsetServiceSid
`func (o *ProxyV1ServiceSessionInteraction) UnsetServiceSid()`

UnsetServiceSid ensures that no value is present for ServiceSid, not even an explicit nil
### GetSessionSid

`func (o *ProxyV1ServiceSessionInteraction) GetSessionSid() string`

GetSessionSid returns the SessionSid field if non-nil, zero value otherwise.

### GetSessionSidOk

`func (o *ProxyV1ServiceSessionInteraction) GetSessionSidOk() (*string, bool)`

GetSessionSidOk returns a tuple with the SessionSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionSid

`func (o *ProxyV1ServiceSessionInteraction) SetSessionSid(v string)`

SetSessionSid sets SessionSid field to given value.

### HasSessionSid

`func (o *ProxyV1ServiceSessionInteraction) HasSessionSid() bool`

HasSessionSid returns a boolean if a field has been set.

### SetSessionSidNil

`func (o *ProxyV1ServiceSessionInteraction) SetSessionSidNil(b bool)`

 SetSessionSidNil sets the value for SessionSid to be an explicit nil

### UnsetSessionSid
`func (o *ProxyV1ServiceSessionInteraction) UnsetSessionSid()`

UnsetSessionSid ensures that no value is present for SessionSid, not even an explicit nil
### GetSid

`func (o *ProxyV1ServiceSessionInteraction) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *ProxyV1ServiceSessionInteraction) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *ProxyV1ServiceSessionInteraction) SetSid(v string)`

SetSid sets Sid field to given value.

### HasSid

`func (o *ProxyV1ServiceSessionInteraction) HasSid() bool`

HasSid returns a boolean if a field has been set.

### SetSidNil

`func (o *ProxyV1ServiceSessionInteraction) SetSidNil(b bool)`

 SetSidNil sets the value for Sid to be an explicit nil

### UnsetSid
`func (o *ProxyV1ServiceSessionInteraction) UnsetSid()`

UnsetSid ensures that no value is present for Sid, not even an explicit nil
### GetType

`func (o *ProxyV1ServiceSessionInteraction) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ProxyV1ServiceSessionInteraction) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ProxyV1ServiceSessionInteraction) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ProxyV1ServiceSessionInteraction) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *ProxyV1ServiceSessionInteraction) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *ProxyV1ServiceSessionInteraction) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetUrl

`func (o *ProxyV1ServiceSessionInteraction) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ProxyV1ServiceSessionInteraction) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ProxyV1ServiceSessionInteraction) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *ProxyV1ServiceSessionInteraction) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *ProxyV1ServiceSessionInteraction) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *ProxyV1ServiceSessionInteraction) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


