# NotifyV1ServiceBinding

## Properties

Name | Type | Description
------------ | ------------- | -------------
**AccountSid** | Pointer to **NullableString** | The SID of the Account that created the resource | [optional] 
**Address** | Pointer to **NullableString** | The channel-specific address | [optional] 
**BindingType** | Pointer to **NullableString** | The type of the Binding | [optional] 
**CredentialSid** | Pointer to **NullableString** | The SID of the Credential resource to be used to send notifications to this Binding | [optional] 
**DateCreated** | Pointer to **NullableTime** | The RFC 2822 date and time in GMT when the resource was created | [optional] 
**DateUpdated** | Pointer to **NullableTime** | The RFC 2822 date and time in GMT when the resource was last updated | [optional] 
**Endpoint** | Pointer to **NullableString** | Deprecated | [optional] 
**Identity** | Pointer to **NullableString** | The &#x60;identity&#x60; value that identifies the new resource&#39;s User | [optional] 
**Links** | Pointer to **map[string]interface{}** | The URLs of related resources | [optional] 
**NotificationProtocolVersion** | Pointer to **NullableString** | The protocol version to use to send the notification | [optional] 
**ServiceSid** | Pointer to **NullableString** | The SID of the Service that the resource is associated with | [optional] 
**Sid** | Pointer to **NullableString** | The unique string that identifies the resource | [optional] 
**Tags** | Pointer to **[]string** | The list of tags associated with this Binding | [optional] 
**Url** | Pointer to **NullableString** | The absolute URL of the Binding resource | [optional] 

## Methods

### NewNotifyV1ServiceBinding

`func NewNotifyV1ServiceBinding() *NotifyV1ServiceBinding`

NewNotifyV1ServiceBinding instantiates a new NotifyV1ServiceBinding object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotifyV1ServiceBindingWithDefaults

`func NewNotifyV1ServiceBindingWithDefaults() *NotifyV1ServiceBinding`

NewNotifyV1ServiceBindingWithDefaults instantiates a new NotifyV1ServiceBinding object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSid

`func (o *NotifyV1ServiceBinding) GetAccountSid() string`

GetAccountSid returns the AccountSid field if non-nil, zero value otherwise.

### GetAccountSidOk

`func (o *NotifyV1ServiceBinding) GetAccountSidOk() (*string, bool)`

GetAccountSidOk returns a tuple with the AccountSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSid

`func (o *NotifyV1ServiceBinding) SetAccountSid(v string)`

SetAccountSid sets AccountSid field to given value.

### HasAccountSid

`func (o *NotifyV1ServiceBinding) HasAccountSid() bool`

HasAccountSid returns a boolean if a field has been set.

### SetAccountSidNil

`func (o *NotifyV1ServiceBinding) SetAccountSidNil(b bool)`

 SetAccountSidNil sets the value for AccountSid to be an explicit nil

### UnsetAccountSid
`func (o *NotifyV1ServiceBinding) UnsetAccountSid()`

UnsetAccountSid ensures that no value is present for AccountSid, not even an explicit nil
### GetAddress

`func (o *NotifyV1ServiceBinding) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *NotifyV1ServiceBinding) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *NotifyV1ServiceBinding) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *NotifyV1ServiceBinding) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### SetAddressNil

`func (o *NotifyV1ServiceBinding) SetAddressNil(b bool)`

 SetAddressNil sets the value for Address to be an explicit nil

### UnsetAddress
`func (o *NotifyV1ServiceBinding) UnsetAddress()`

UnsetAddress ensures that no value is present for Address, not even an explicit nil
### GetBindingType

`func (o *NotifyV1ServiceBinding) GetBindingType() string`

GetBindingType returns the BindingType field if non-nil, zero value otherwise.

### GetBindingTypeOk

`func (o *NotifyV1ServiceBinding) GetBindingTypeOk() (*string, bool)`

GetBindingTypeOk returns a tuple with the BindingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindingType

`func (o *NotifyV1ServiceBinding) SetBindingType(v string)`

SetBindingType sets BindingType field to given value.

### HasBindingType

`func (o *NotifyV1ServiceBinding) HasBindingType() bool`

HasBindingType returns a boolean if a field has been set.

### SetBindingTypeNil

`func (o *NotifyV1ServiceBinding) SetBindingTypeNil(b bool)`

 SetBindingTypeNil sets the value for BindingType to be an explicit nil

### UnsetBindingType
`func (o *NotifyV1ServiceBinding) UnsetBindingType()`

UnsetBindingType ensures that no value is present for BindingType, not even an explicit nil
### GetCredentialSid

`func (o *NotifyV1ServiceBinding) GetCredentialSid() string`

GetCredentialSid returns the CredentialSid field if non-nil, zero value otherwise.

### GetCredentialSidOk

`func (o *NotifyV1ServiceBinding) GetCredentialSidOk() (*string, bool)`

GetCredentialSidOk returns a tuple with the CredentialSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialSid

`func (o *NotifyV1ServiceBinding) SetCredentialSid(v string)`

SetCredentialSid sets CredentialSid field to given value.

### HasCredentialSid

`func (o *NotifyV1ServiceBinding) HasCredentialSid() bool`

HasCredentialSid returns a boolean if a field has been set.

### SetCredentialSidNil

`func (o *NotifyV1ServiceBinding) SetCredentialSidNil(b bool)`

 SetCredentialSidNil sets the value for CredentialSid to be an explicit nil

### UnsetCredentialSid
`func (o *NotifyV1ServiceBinding) UnsetCredentialSid()`

UnsetCredentialSid ensures that no value is present for CredentialSid, not even an explicit nil
### GetDateCreated

`func (o *NotifyV1ServiceBinding) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *NotifyV1ServiceBinding) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *NotifyV1ServiceBinding) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *NotifyV1ServiceBinding) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### SetDateCreatedNil

`func (o *NotifyV1ServiceBinding) SetDateCreatedNil(b bool)`

 SetDateCreatedNil sets the value for DateCreated to be an explicit nil

### UnsetDateCreated
`func (o *NotifyV1ServiceBinding) UnsetDateCreated()`

UnsetDateCreated ensures that no value is present for DateCreated, not even an explicit nil
### GetDateUpdated

`func (o *NotifyV1ServiceBinding) GetDateUpdated() time.Time`

GetDateUpdated returns the DateUpdated field if non-nil, zero value otherwise.

### GetDateUpdatedOk

`func (o *NotifyV1ServiceBinding) GetDateUpdatedOk() (*time.Time, bool)`

GetDateUpdatedOk returns a tuple with the DateUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateUpdated

`func (o *NotifyV1ServiceBinding) SetDateUpdated(v time.Time)`

SetDateUpdated sets DateUpdated field to given value.

### HasDateUpdated

`func (o *NotifyV1ServiceBinding) HasDateUpdated() bool`

HasDateUpdated returns a boolean if a field has been set.

### SetDateUpdatedNil

`func (o *NotifyV1ServiceBinding) SetDateUpdatedNil(b bool)`

 SetDateUpdatedNil sets the value for DateUpdated to be an explicit nil

### UnsetDateUpdated
`func (o *NotifyV1ServiceBinding) UnsetDateUpdated()`

UnsetDateUpdated ensures that no value is present for DateUpdated, not even an explicit nil
### GetEndpoint

`func (o *NotifyV1ServiceBinding) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *NotifyV1ServiceBinding) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *NotifyV1ServiceBinding) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.

### HasEndpoint

`func (o *NotifyV1ServiceBinding) HasEndpoint() bool`

HasEndpoint returns a boolean if a field has been set.

### SetEndpointNil

`func (o *NotifyV1ServiceBinding) SetEndpointNil(b bool)`

 SetEndpointNil sets the value for Endpoint to be an explicit nil

### UnsetEndpoint
`func (o *NotifyV1ServiceBinding) UnsetEndpoint()`

UnsetEndpoint ensures that no value is present for Endpoint, not even an explicit nil
### GetIdentity

`func (o *NotifyV1ServiceBinding) GetIdentity() string`

GetIdentity returns the Identity field if non-nil, zero value otherwise.

### GetIdentityOk

`func (o *NotifyV1ServiceBinding) GetIdentityOk() (*string, bool)`

GetIdentityOk returns a tuple with the Identity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentity

`func (o *NotifyV1ServiceBinding) SetIdentity(v string)`

SetIdentity sets Identity field to given value.

### HasIdentity

`func (o *NotifyV1ServiceBinding) HasIdentity() bool`

HasIdentity returns a boolean if a field has been set.

### SetIdentityNil

`func (o *NotifyV1ServiceBinding) SetIdentityNil(b bool)`

 SetIdentityNil sets the value for Identity to be an explicit nil

### UnsetIdentity
`func (o *NotifyV1ServiceBinding) UnsetIdentity()`

UnsetIdentity ensures that no value is present for Identity, not even an explicit nil
### GetLinks

`func (o *NotifyV1ServiceBinding) GetLinks() map[string]interface{}`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *NotifyV1ServiceBinding) GetLinksOk() (*map[string]interface{}, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *NotifyV1ServiceBinding) SetLinks(v map[string]interface{})`

SetLinks sets Links field to given value.

### HasLinks

`func (o *NotifyV1ServiceBinding) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### SetLinksNil

`func (o *NotifyV1ServiceBinding) SetLinksNil(b bool)`

 SetLinksNil sets the value for Links to be an explicit nil

### UnsetLinks
`func (o *NotifyV1ServiceBinding) UnsetLinks()`

UnsetLinks ensures that no value is present for Links, not even an explicit nil
### GetNotificationProtocolVersion

`func (o *NotifyV1ServiceBinding) GetNotificationProtocolVersion() string`

GetNotificationProtocolVersion returns the NotificationProtocolVersion field if non-nil, zero value otherwise.

### GetNotificationProtocolVersionOk

`func (o *NotifyV1ServiceBinding) GetNotificationProtocolVersionOk() (*string, bool)`

GetNotificationProtocolVersionOk returns a tuple with the NotificationProtocolVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationProtocolVersion

`func (o *NotifyV1ServiceBinding) SetNotificationProtocolVersion(v string)`

SetNotificationProtocolVersion sets NotificationProtocolVersion field to given value.

### HasNotificationProtocolVersion

`func (o *NotifyV1ServiceBinding) HasNotificationProtocolVersion() bool`

HasNotificationProtocolVersion returns a boolean if a field has been set.

### SetNotificationProtocolVersionNil

`func (o *NotifyV1ServiceBinding) SetNotificationProtocolVersionNil(b bool)`

 SetNotificationProtocolVersionNil sets the value for NotificationProtocolVersion to be an explicit nil

### UnsetNotificationProtocolVersion
`func (o *NotifyV1ServiceBinding) UnsetNotificationProtocolVersion()`

UnsetNotificationProtocolVersion ensures that no value is present for NotificationProtocolVersion, not even an explicit nil
### GetServiceSid

`func (o *NotifyV1ServiceBinding) GetServiceSid() string`

GetServiceSid returns the ServiceSid field if non-nil, zero value otherwise.

### GetServiceSidOk

`func (o *NotifyV1ServiceBinding) GetServiceSidOk() (*string, bool)`

GetServiceSidOk returns a tuple with the ServiceSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceSid

`func (o *NotifyV1ServiceBinding) SetServiceSid(v string)`

SetServiceSid sets ServiceSid field to given value.

### HasServiceSid

`func (o *NotifyV1ServiceBinding) HasServiceSid() bool`

HasServiceSid returns a boolean if a field has been set.

### SetServiceSidNil

`func (o *NotifyV1ServiceBinding) SetServiceSidNil(b bool)`

 SetServiceSidNil sets the value for ServiceSid to be an explicit nil

### UnsetServiceSid
`func (o *NotifyV1ServiceBinding) UnsetServiceSid()`

UnsetServiceSid ensures that no value is present for ServiceSid, not even an explicit nil
### GetSid

`func (o *NotifyV1ServiceBinding) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *NotifyV1ServiceBinding) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *NotifyV1ServiceBinding) SetSid(v string)`

SetSid sets Sid field to given value.

### HasSid

`func (o *NotifyV1ServiceBinding) HasSid() bool`

HasSid returns a boolean if a field has been set.

### SetSidNil

`func (o *NotifyV1ServiceBinding) SetSidNil(b bool)`

 SetSidNil sets the value for Sid to be an explicit nil

### UnsetSid
`func (o *NotifyV1ServiceBinding) UnsetSid()`

UnsetSid ensures that no value is present for Sid, not even an explicit nil
### GetTags

`func (o *NotifyV1ServiceBinding) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *NotifyV1ServiceBinding) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *NotifyV1ServiceBinding) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *NotifyV1ServiceBinding) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *NotifyV1ServiceBinding) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *NotifyV1ServiceBinding) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetUrl

`func (o *NotifyV1ServiceBinding) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *NotifyV1ServiceBinding) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *NotifyV1ServiceBinding) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *NotifyV1ServiceBinding) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *NotifyV1ServiceBinding) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *NotifyV1ServiceBinding) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


