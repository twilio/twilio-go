# ApiV2010AccountAvailablePhoneNumberCountry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Beta** | Pointer to **NullableBool** | Whether all phone numbers available in the country are new to the Twilio platform. | [optional] 
**Country** | Pointer to **NullableString** | The name of the country | [optional] 
**CountryCode** | Pointer to **NullableString** | The ISO-3166-1 country code of the country. | [optional] 
**SubresourceUris** | Pointer to **map[string]interface{}** | A list of related resources identified by their relative URIs | [optional] 
**Uri** | Pointer to **NullableString** | The URI of the Country resource, relative to &#x60;https://api.twilio.com&#x60; | [optional] 

## Methods

### NewApiV2010AccountAvailablePhoneNumberCountry

`func NewApiV2010AccountAvailablePhoneNumberCountry() *ApiV2010AccountAvailablePhoneNumberCountry`

NewApiV2010AccountAvailablePhoneNumberCountry instantiates a new ApiV2010AccountAvailablePhoneNumberCountry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiV2010AccountAvailablePhoneNumberCountryWithDefaults

`func NewApiV2010AccountAvailablePhoneNumberCountryWithDefaults() *ApiV2010AccountAvailablePhoneNumberCountry`

NewApiV2010AccountAvailablePhoneNumberCountryWithDefaults instantiates a new ApiV2010AccountAvailablePhoneNumberCountry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBeta

`func (o *ApiV2010AccountAvailablePhoneNumberCountry) GetBeta() bool`

GetBeta returns the Beta field if non-nil, zero value otherwise.

### GetBetaOk

`func (o *ApiV2010AccountAvailablePhoneNumberCountry) GetBetaOk() (*bool, bool)`

GetBetaOk returns a tuple with the Beta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeta

`func (o *ApiV2010AccountAvailablePhoneNumberCountry) SetBeta(v bool)`

SetBeta sets Beta field to given value.

### HasBeta

`func (o *ApiV2010AccountAvailablePhoneNumberCountry) HasBeta() bool`

HasBeta returns a boolean if a field has been set.

### SetBetaNil

`func (o *ApiV2010AccountAvailablePhoneNumberCountry) SetBetaNil(b bool)`

 SetBetaNil sets the value for Beta to be an explicit nil

### UnsetBeta
`func (o *ApiV2010AccountAvailablePhoneNumberCountry) UnsetBeta()`

UnsetBeta ensures that no value is present for Beta, not even an explicit nil
### GetCountry

`func (o *ApiV2010AccountAvailablePhoneNumberCountry) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *ApiV2010AccountAvailablePhoneNumberCountry) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *ApiV2010AccountAvailablePhoneNumberCountry) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *ApiV2010AccountAvailablePhoneNumberCountry) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### SetCountryNil

`func (o *ApiV2010AccountAvailablePhoneNumberCountry) SetCountryNil(b bool)`

 SetCountryNil sets the value for Country to be an explicit nil

### UnsetCountry
`func (o *ApiV2010AccountAvailablePhoneNumberCountry) UnsetCountry()`

UnsetCountry ensures that no value is present for Country, not even an explicit nil
### GetCountryCode

`func (o *ApiV2010AccountAvailablePhoneNumberCountry) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *ApiV2010AccountAvailablePhoneNumberCountry) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *ApiV2010AccountAvailablePhoneNumberCountry) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *ApiV2010AccountAvailablePhoneNumberCountry) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### SetCountryCodeNil

`func (o *ApiV2010AccountAvailablePhoneNumberCountry) SetCountryCodeNil(b bool)`

 SetCountryCodeNil sets the value for CountryCode to be an explicit nil

### UnsetCountryCode
`func (o *ApiV2010AccountAvailablePhoneNumberCountry) UnsetCountryCode()`

UnsetCountryCode ensures that no value is present for CountryCode, not even an explicit nil
### GetSubresourceUris

`func (o *ApiV2010AccountAvailablePhoneNumberCountry) GetSubresourceUris() map[string]interface{}`

GetSubresourceUris returns the SubresourceUris field if non-nil, zero value otherwise.

### GetSubresourceUrisOk

`func (o *ApiV2010AccountAvailablePhoneNumberCountry) GetSubresourceUrisOk() (*map[string]interface{}, bool)`

GetSubresourceUrisOk returns a tuple with the SubresourceUris field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubresourceUris

`func (o *ApiV2010AccountAvailablePhoneNumberCountry) SetSubresourceUris(v map[string]interface{})`

SetSubresourceUris sets SubresourceUris field to given value.

### HasSubresourceUris

`func (o *ApiV2010AccountAvailablePhoneNumberCountry) HasSubresourceUris() bool`

HasSubresourceUris returns a boolean if a field has been set.

### SetSubresourceUrisNil

`func (o *ApiV2010AccountAvailablePhoneNumberCountry) SetSubresourceUrisNil(b bool)`

 SetSubresourceUrisNil sets the value for SubresourceUris to be an explicit nil

### UnsetSubresourceUris
`func (o *ApiV2010AccountAvailablePhoneNumberCountry) UnsetSubresourceUris()`

UnsetSubresourceUris ensures that no value is present for SubresourceUris, not even an explicit nil
### GetUri

`func (o *ApiV2010AccountAvailablePhoneNumberCountry) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *ApiV2010AccountAvailablePhoneNumberCountry) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *ApiV2010AccountAvailablePhoneNumberCountry) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *ApiV2010AccountAvailablePhoneNumberCountry) HasUri() bool`

HasUri returns a boolean if a field has been set.

### SetUriNil

`func (o *ApiV2010AccountAvailablePhoneNumberCountry) SetUriNil(b bool)`

 SetUriNil sets the value for Uri to be an explicit nil

### UnsetUri
`func (o *ApiV2010AccountAvailablePhoneNumberCountry) UnsetUri()`

UnsetUri ensures that no value is present for Uri, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


