# MessagingV1ServiceUsAppToPerson

## Properties

Name | Type | Description
------------ | ------------- | -------------
**AccountSid** | Pointer to **NullableString** | The SID of the Account that created the resource | [optional] 
**BrandRegistrationSid** | Pointer to **NullableString** | A2P Brand Registration SID | [optional] 
**CampaignId** | Pointer to **NullableString** | The Campaign Registry (TCR) Campaign ID. | [optional] 
**DateCreated** | Pointer to **NullableTime** | The ISO 8601 date and time in GMT when the resource was created | [optional] 
**DateUpdated** | Pointer to **NullableTime** | The ISO 8601 date and time in GMT when the resource was last updated | [optional] 
**Description** | Pointer to **NullableString** | A short description of what this SMS campaign does | [optional] 
**HasEmbeddedLinks** | Pointer to **NullableBool** | Indicate that this SMS campaign will send messages that contain links | [optional] 
**HasEmbeddedPhone** | Pointer to **NullableBool** | Indicates that this SMS campaign will send messages that contain phone numbers | [optional] 
**IsExternallyRegistered** | Pointer to **NullableBool** | Indicates whether the campaign was registered externally or not | [optional] 
**MessageSamples** | Pointer to **[]string** | Message samples | [optional] 
**MessagingServiceSid** | Pointer to **NullableString** | The SID of the Messaging Service the resource is associated with | [optional] 
**RateLimits** | Pointer to **map[string]interface{}** | Rate limit and/or classification set by each carrier | [optional] 
**Status** | Pointer to **NullableString** | Campaign status | [optional] 
**Url** | Pointer to **NullableString** | The absolute URL of the US App to Person resource | [optional] 
**UsAppToPersonUsecase** | Pointer to **NullableString** | A2P Campaign Use Case. | [optional] 

## Methods

### NewMessagingV1ServiceUsAppToPerson

`func NewMessagingV1ServiceUsAppToPerson() *MessagingV1ServiceUsAppToPerson`

NewMessagingV1ServiceUsAppToPerson instantiates a new MessagingV1ServiceUsAppToPerson object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMessagingV1ServiceUsAppToPersonWithDefaults

`func NewMessagingV1ServiceUsAppToPersonWithDefaults() *MessagingV1ServiceUsAppToPerson`

NewMessagingV1ServiceUsAppToPersonWithDefaults instantiates a new MessagingV1ServiceUsAppToPerson object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSid

`func (o *MessagingV1ServiceUsAppToPerson) GetAccountSid() string`

GetAccountSid returns the AccountSid field if non-nil, zero value otherwise.

### GetAccountSidOk

`func (o *MessagingV1ServiceUsAppToPerson) GetAccountSidOk() (*string, bool)`

GetAccountSidOk returns a tuple with the AccountSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSid

`func (o *MessagingV1ServiceUsAppToPerson) SetAccountSid(v string)`

SetAccountSid sets AccountSid field to given value.

### HasAccountSid

`func (o *MessagingV1ServiceUsAppToPerson) HasAccountSid() bool`

HasAccountSid returns a boolean if a field has been set.

### SetAccountSidNil

`func (o *MessagingV1ServiceUsAppToPerson) SetAccountSidNil(b bool)`

 SetAccountSidNil sets the value for AccountSid to be an explicit nil

### UnsetAccountSid
`func (o *MessagingV1ServiceUsAppToPerson) UnsetAccountSid()`

UnsetAccountSid ensures that no value is present for AccountSid, not even an explicit nil
### GetBrandRegistrationSid

`func (o *MessagingV1ServiceUsAppToPerson) GetBrandRegistrationSid() string`

GetBrandRegistrationSid returns the BrandRegistrationSid field if non-nil, zero value otherwise.

### GetBrandRegistrationSidOk

`func (o *MessagingV1ServiceUsAppToPerson) GetBrandRegistrationSidOk() (*string, bool)`

GetBrandRegistrationSidOk returns a tuple with the BrandRegistrationSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrandRegistrationSid

`func (o *MessagingV1ServiceUsAppToPerson) SetBrandRegistrationSid(v string)`

SetBrandRegistrationSid sets BrandRegistrationSid field to given value.

### HasBrandRegistrationSid

`func (o *MessagingV1ServiceUsAppToPerson) HasBrandRegistrationSid() bool`

HasBrandRegistrationSid returns a boolean if a field has been set.

### SetBrandRegistrationSidNil

`func (o *MessagingV1ServiceUsAppToPerson) SetBrandRegistrationSidNil(b bool)`

 SetBrandRegistrationSidNil sets the value for BrandRegistrationSid to be an explicit nil

### UnsetBrandRegistrationSid
`func (o *MessagingV1ServiceUsAppToPerson) UnsetBrandRegistrationSid()`

UnsetBrandRegistrationSid ensures that no value is present for BrandRegistrationSid, not even an explicit nil
### GetCampaignId

`func (o *MessagingV1ServiceUsAppToPerson) GetCampaignId() string`

GetCampaignId returns the CampaignId field if non-nil, zero value otherwise.

### GetCampaignIdOk

`func (o *MessagingV1ServiceUsAppToPerson) GetCampaignIdOk() (*string, bool)`

GetCampaignIdOk returns a tuple with the CampaignId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignId

`func (o *MessagingV1ServiceUsAppToPerson) SetCampaignId(v string)`

SetCampaignId sets CampaignId field to given value.

### HasCampaignId

`func (o *MessagingV1ServiceUsAppToPerson) HasCampaignId() bool`

HasCampaignId returns a boolean if a field has been set.

### SetCampaignIdNil

`func (o *MessagingV1ServiceUsAppToPerson) SetCampaignIdNil(b bool)`

 SetCampaignIdNil sets the value for CampaignId to be an explicit nil

### UnsetCampaignId
`func (o *MessagingV1ServiceUsAppToPerson) UnsetCampaignId()`

UnsetCampaignId ensures that no value is present for CampaignId, not even an explicit nil
### GetDateCreated

`func (o *MessagingV1ServiceUsAppToPerson) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *MessagingV1ServiceUsAppToPerson) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *MessagingV1ServiceUsAppToPerson) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *MessagingV1ServiceUsAppToPerson) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### SetDateCreatedNil

`func (o *MessagingV1ServiceUsAppToPerson) SetDateCreatedNil(b bool)`

 SetDateCreatedNil sets the value for DateCreated to be an explicit nil

### UnsetDateCreated
`func (o *MessagingV1ServiceUsAppToPerson) UnsetDateCreated()`

UnsetDateCreated ensures that no value is present for DateCreated, not even an explicit nil
### GetDateUpdated

`func (o *MessagingV1ServiceUsAppToPerson) GetDateUpdated() time.Time`

GetDateUpdated returns the DateUpdated field if non-nil, zero value otherwise.

### GetDateUpdatedOk

`func (o *MessagingV1ServiceUsAppToPerson) GetDateUpdatedOk() (*time.Time, bool)`

GetDateUpdatedOk returns a tuple with the DateUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateUpdated

`func (o *MessagingV1ServiceUsAppToPerson) SetDateUpdated(v time.Time)`

SetDateUpdated sets DateUpdated field to given value.

### HasDateUpdated

`func (o *MessagingV1ServiceUsAppToPerson) HasDateUpdated() bool`

HasDateUpdated returns a boolean if a field has been set.

### SetDateUpdatedNil

`func (o *MessagingV1ServiceUsAppToPerson) SetDateUpdatedNil(b bool)`

 SetDateUpdatedNil sets the value for DateUpdated to be an explicit nil

### UnsetDateUpdated
`func (o *MessagingV1ServiceUsAppToPerson) UnsetDateUpdated()`

UnsetDateUpdated ensures that no value is present for DateUpdated, not even an explicit nil
### GetDescription

`func (o *MessagingV1ServiceUsAppToPerson) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MessagingV1ServiceUsAppToPerson) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MessagingV1ServiceUsAppToPerson) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MessagingV1ServiceUsAppToPerson) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *MessagingV1ServiceUsAppToPerson) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *MessagingV1ServiceUsAppToPerson) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetHasEmbeddedLinks

`func (o *MessagingV1ServiceUsAppToPerson) GetHasEmbeddedLinks() bool`

GetHasEmbeddedLinks returns the HasEmbeddedLinks field if non-nil, zero value otherwise.

### GetHasEmbeddedLinksOk

`func (o *MessagingV1ServiceUsAppToPerson) GetHasEmbeddedLinksOk() (*bool, bool)`

GetHasEmbeddedLinksOk returns a tuple with the HasEmbeddedLinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasEmbeddedLinks

`func (o *MessagingV1ServiceUsAppToPerson) SetHasEmbeddedLinks(v bool)`

SetHasEmbeddedLinks sets HasEmbeddedLinks field to given value.

### HasHasEmbeddedLinks

`func (o *MessagingV1ServiceUsAppToPerson) HasHasEmbeddedLinks() bool`

HasHasEmbeddedLinks returns a boolean if a field has been set.

### SetHasEmbeddedLinksNil

`func (o *MessagingV1ServiceUsAppToPerson) SetHasEmbeddedLinksNil(b bool)`

 SetHasEmbeddedLinksNil sets the value for HasEmbeddedLinks to be an explicit nil

### UnsetHasEmbeddedLinks
`func (o *MessagingV1ServiceUsAppToPerson) UnsetHasEmbeddedLinks()`

UnsetHasEmbeddedLinks ensures that no value is present for HasEmbeddedLinks, not even an explicit nil
### GetHasEmbeddedPhone

`func (o *MessagingV1ServiceUsAppToPerson) GetHasEmbeddedPhone() bool`

GetHasEmbeddedPhone returns the HasEmbeddedPhone field if non-nil, zero value otherwise.

### GetHasEmbeddedPhoneOk

`func (o *MessagingV1ServiceUsAppToPerson) GetHasEmbeddedPhoneOk() (*bool, bool)`

GetHasEmbeddedPhoneOk returns a tuple with the HasEmbeddedPhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasEmbeddedPhone

`func (o *MessagingV1ServiceUsAppToPerson) SetHasEmbeddedPhone(v bool)`

SetHasEmbeddedPhone sets HasEmbeddedPhone field to given value.

### HasHasEmbeddedPhone

`func (o *MessagingV1ServiceUsAppToPerson) HasHasEmbeddedPhone() bool`

HasHasEmbeddedPhone returns a boolean if a field has been set.

### SetHasEmbeddedPhoneNil

`func (o *MessagingV1ServiceUsAppToPerson) SetHasEmbeddedPhoneNil(b bool)`

 SetHasEmbeddedPhoneNil sets the value for HasEmbeddedPhone to be an explicit nil

### UnsetHasEmbeddedPhone
`func (o *MessagingV1ServiceUsAppToPerson) UnsetHasEmbeddedPhone()`

UnsetHasEmbeddedPhone ensures that no value is present for HasEmbeddedPhone, not even an explicit nil
### GetIsExternallyRegistered

`func (o *MessagingV1ServiceUsAppToPerson) GetIsExternallyRegistered() bool`

GetIsExternallyRegistered returns the IsExternallyRegistered field if non-nil, zero value otherwise.

### GetIsExternallyRegisteredOk

`func (o *MessagingV1ServiceUsAppToPerson) GetIsExternallyRegisteredOk() (*bool, bool)`

GetIsExternallyRegisteredOk returns a tuple with the IsExternallyRegistered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsExternallyRegistered

`func (o *MessagingV1ServiceUsAppToPerson) SetIsExternallyRegistered(v bool)`

SetIsExternallyRegistered sets IsExternallyRegistered field to given value.

### HasIsExternallyRegistered

`func (o *MessagingV1ServiceUsAppToPerson) HasIsExternallyRegistered() bool`

HasIsExternallyRegistered returns a boolean if a field has been set.

### SetIsExternallyRegisteredNil

`func (o *MessagingV1ServiceUsAppToPerson) SetIsExternallyRegisteredNil(b bool)`

 SetIsExternallyRegisteredNil sets the value for IsExternallyRegistered to be an explicit nil

### UnsetIsExternallyRegistered
`func (o *MessagingV1ServiceUsAppToPerson) UnsetIsExternallyRegistered()`

UnsetIsExternallyRegistered ensures that no value is present for IsExternallyRegistered, not even an explicit nil
### GetMessageSamples

`func (o *MessagingV1ServiceUsAppToPerson) GetMessageSamples() []string`

GetMessageSamples returns the MessageSamples field if non-nil, zero value otherwise.

### GetMessageSamplesOk

`func (o *MessagingV1ServiceUsAppToPerson) GetMessageSamplesOk() (*[]string, bool)`

GetMessageSamplesOk returns a tuple with the MessageSamples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageSamples

`func (o *MessagingV1ServiceUsAppToPerson) SetMessageSamples(v []string)`

SetMessageSamples sets MessageSamples field to given value.

### HasMessageSamples

`func (o *MessagingV1ServiceUsAppToPerson) HasMessageSamples() bool`

HasMessageSamples returns a boolean if a field has been set.

### SetMessageSamplesNil

`func (o *MessagingV1ServiceUsAppToPerson) SetMessageSamplesNil(b bool)`

 SetMessageSamplesNil sets the value for MessageSamples to be an explicit nil

### UnsetMessageSamples
`func (o *MessagingV1ServiceUsAppToPerson) UnsetMessageSamples()`

UnsetMessageSamples ensures that no value is present for MessageSamples, not even an explicit nil
### GetMessagingServiceSid

`func (o *MessagingV1ServiceUsAppToPerson) GetMessagingServiceSid() string`

GetMessagingServiceSid returns the MessagingServiceSid field if non-nil, zero value otherwise.

### GetMessagingServiceSidOk

`func (o *MessagingV1ServiceUsAppToPerson) GetMessagingServiceSidOk() (*string, bool)`

GetMessagingServiceSidOk returns a tuple with the MessagingServiceSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessagingServiceSid

`func (o *MessagingV1ServiceUsAppToPerson) SetMessagingServiceSid(v string)`

SetMessagingServiceSid sets MessagingServiceSid field to given value.

### HasMessagingServiceSid

`func (o *MessagingV1ServiceUsAppToPerson) HasMessagingServiceSid() bool`

HasMessagingServiceSid returns a boolean if a field has been set.

### SetMessagingServiceSidNil

`func (o *MessagingV1ServiceUsAppToPerson) SetMessagingServiceSidNil(b bool)`

 SetMessagingServiceSidNil sets the value for MessagingServiceSid to be an explicit nil

### UnsetMessagingServiceSid
`func (o *MessagingV1ServiceUsAppToPerson) UnsetMessagingServiceSid()`

UnsetMessagingServiceSid ensures that no value is present for MessagingServiceSid, not even an explicit nil
### GetRateLimits

`func (o *MessagingV1ServiceUsAppToPerson) GetRateLimits() map[string]interface{}`

GetRateLimits returns the RateLimits field if non-nil, zero value otherwise.

### GetRateLimitsOk

`func (o *MessagingV1ServiceUsAppToPerson) GetRateLimitsOk() (*map[string]interface{}, bool)`

GetRateLimitsOk returns a tuple with the RateLimits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateLimits

`func (o *MessagingV1ServiceUsAppToPerson) SetRateLimits(v map[string]interface{})`

SetRateLimits sets RateLimits field to given value.

### HasRateLimits

`func (o *MessagingV1ServiceUsAppToPerson) HasRateLimits() bool`

HasRateLimits returns a boolean if a field has been set.

### SetRateLimitsNil

`func (o *MessagingV1ServiceUsAppToPerson) SetRateLimitsNil(b bool)`

 SetRateLimitsNil sets the value for RateLimits to be an explicit nil

### UnsetRateLimits
`func (o *MessagingV1ServiceUsAppToPerson) UnsetRateLimits()`

UnsetRateLimits ensures that no value is present for RateLimits, not even an explicit nil
### GetStatus

`func (o *MessagingV1ServiceUsAppToPerson) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MessagingV1ServiceUsAppToPerson) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MessagingV1ServiceUsAppToPerson) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MessagingV1ServiceUsAppToPerson) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *MessagingV1ServiceUsAppToPerson) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *MessagingV1ServiceUsAppToPerson) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetUrl

`func (o *MessagingV1ServiceUsAppToPerson) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *MessagingV1ServiceUsAppToPerson) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *MessagingV1ServiceUsAppToPerson) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *MessagingV1ServiceUsAppToPerson) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *MessagingV1ServiceUsAppToPerson) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *MessagingV1ServiceUsAppToPerson) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetUsAppToPersonUsecase

`func (o *MessagingV1ServiceUsAppToPerson) GetUsAppToPersonUsecase() string`

GetUsAppToPersonUsecase returns the UsAppToPersonUsecase field if non-nil, zero value otherwise.

### GetUsAppToPersonUsecaseOk

`func (o *MessagingV1ServiceUsAppToPerson) GetUsAppToPersonUsecaseOk() (*string, bool)`

GetUsAppToPersonUsecaseOk returns a tuple with the UsAppToPersonUsecase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsAppToPersonUsecase

`func (o *MessagingV1ServiceUsAppToPerson) SetUsAppToPersonUsecase(v string)`

SetUsAppToPersonUsecase sets UsAppToPersonUsecase field to given value.

### HasUsAppToPersonUsecase

`func (o *MessagingV1ServiceUsAppToPerson) HasUsAppToPersonUsecase() bool`

HasUsAppToPersonUsecase returns a boolean if a field has been set.

### SetUsAppToPersonUsecaseNil

`func (o *MessagingV1ServiceUsAppToPerson) SetUsAppToPersonUsecaseNil(b bool)`

 SetUsAppToPersonUsecaseNil sets the value for UsAppToPersonUsecase to be an explicit nil

### UnsetUsAppToPersonUsecase
`func (o *MessagingV1ServiceUsAppToPerson) UnsetUsAppToPersonUsecase()`

UnsetUsAppToPersonUsecase ensures that no value is present for UsAppToPersonUsecase, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


