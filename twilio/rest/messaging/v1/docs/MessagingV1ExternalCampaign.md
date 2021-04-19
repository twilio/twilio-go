# MessagingV1ExternalCampaign

## Properties

Name | Type | Description
------------ | ------------- | -------------
**AccountSid** | Pointer to **NullableString** | The SID of the Account that created the resource | [optional] 
**CampaignId** | Pointer to **NullableString** | ID of the preregistered campaign. | [optional] 
**DateCreated** | Pointer to **NullableTime** | The ISO 8601 date and time in GMT when the resource was created | [optional] 
**MessagingServiceSid** | Pointer to **NullableString** | The SID of the Messaging Service the resource is associated with | [optional] 

## Methods

### NewMessagingV1ExternalCampaign

`func NewMessagingV1ExternalCampaign() *MessagingV1ExternalCampaign`

NewMessagingV1ExternalCampaign instantiates a new MessagingV1ExternalCampaign object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMessagingV1ExternalCampaignWithDefaults

`func NewMessagingV1ExternalCampaignWithDefaults() *MessagingV1ExternalCampaign`

NewMessagingV1ExternalCampaignWithDefaults instantiates a new MessagingV1ExternalCampaign object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSid

`func (o *MessagingV1ExternalCampaign) GetAccountSid() string`

GetAccountSid returns the AccountSid field if non-nil, zero value otherwise.

### GetAccountSidOk

`func (o *MessagingV1ExternalCampaign) GetAccountSidOk() (*string, bool)`

GetAccountSidOk returns a tuple with the AccountSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSid

`func (o *MessagingV1ExternalCampaign) SetAccountSid(v string)`

SetAccountSid sets AccountSid field to given value.

### HasAccountSid

`func (o *MessagingV1ExternalCampaign) HasAccountSid() bool`

HasAccountSid returns a boolean if a field has been set.

### SetAccountSidNil

`func (o *MessagingV1ExternalCampaign) SetAccountSidNil(b bool)`

 SetAccountSidNil sets the value for AccountSid to be an explicit nil

### UnsetAccountSid
`func (o *MessagingV1ExternalCampaign) UnsetAccountSid()`

UnsetAccountSid ensures that no value is present for AccountSid, not even an explicit nil
### GetCampaignId

`func (o *MessagingV1ExternalCampaign) GetCampaignId() string`

GetCampaignId returns the CampaignId field if non-nil, zero value otherwise.

### GetCampaignIdOk

`func (o *MessagingV1ExternalCampaign) GetCampaignIdOk() (*string, bool)`

GetCampaignIdOk returns a tuple with the CampaignId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignId

`func (o *MessagingV1ExternalCampaign) SetCampaignId(v string)`

SetCampaignId sets CampaignId field to given value.

### HasCampaignId

`func (o *MessagingV1ExternalCampaign) HasCampaignId() bool`

HasCampaignId returns a boolean if a field has been set.

### SetCampaignIdNil

`func (o *MessagingV1ExternalCampaign) SetCampaignIdNil(b bool)`

 SetCampaignIdNil sets the value for CampaignId to be an explicit nil

### UnsetCampaignId
`func (o *MessagingV1ExternalCampaign) UnsetCampaignId()`

UnsetCampaignId ensures that no value is present for CampaignId, not even an explicit nil
### GetDateCreated

`func (o *MessagingV1ExternalCampaign) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *MessagingV1ExternalCampaign) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *MessagingV1ExternalCampaign) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *MessagingV1ExternalCampaign) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### SetDateCreatedNil

`func (o *MessagingV1ExternalCampaign) SetDateCreatedNil(b bool)`

 SetDateCreatedNil sets the value for DateCreated to be an explicit nil

### UnsetDateCreated
`func (o *MessagingV1ExternalCampaign) UnsetDateCreated()`

UnsetDateCreated ensures that no value is present for DateCreated, not even an explicit nil
### GetMessagingServiceSid

`func (o *MessagingV1ExternalCampaign) GetMessagingServiceSid() string`

GetMessagingServiceSid returns the MessagingServiceSid field if non-nil, zero value otherwise.

### GetMessagingServiceSidOk

`func (o *MessagingV1ExternalCampaign) GetMessagingServiceSidOk() (*string, bool)`

GetMessagingServiceSidOk returns a tuple with the MessagingServiceSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessagingServiceSid

`func (o *MessagingV1ExternalCampaign) SetMessagingServiceSid(v string)`

SetMessagingServiceSid sets MessagingServiceSid field to given value.

### HasMessagingServiceSid

`func (o *MessagingV1ExternalCampaign) HasMessagingServiceSid() bool`

HasMessagingServiceSid returns a boolean if a field has been set.

### SetMessagingServiceSidNil

`func (o *MessagingV1ExternalCampaign) SetMessagingServiceSidNil(b bool)`

 SetMessagingServiceSidNil sets the value for MessagingServiceSid to be an explicit nil

### UnsetMessagingServiceSid
`func (o *MessagingV1ExternalCampaign) UnsetMessagingServiceSid()`

UnsetMessagingServiceSid ensures that no value is present for MessagingServiceSid, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


