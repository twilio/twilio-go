# VoiceV1DialingPermissionsDialingPermissionsCountry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Continent** | Pointer to **NullableString** | The name of the continent in which the country is located | [optional] 
**CountryCodes** | Pointer to **[]string** | The E.164 assigned country codes(s) | [optional] 
**HighRiskSpecialNumbersEnabled** | Pointer to **NullableBool** | Whether dialing to high-risk special services numbers is enabled | [optional] 
**HighRiskTollfraudNumbersEnabled** | Pointer to **NullableBool** | Whether dialing to high-risk toll fraud numbers is enabled, else &#x60;false&#x60; | [optional] 
**IsoCode** | Pointer to **NullableString** | The ISO country code | [optional] 
**Links** | Pointer to **map[string]interface{}** | A list of URLs related to this resource | [optional] 
**LowRiskNumbersEnabled** | Pointer to **NullableBool** | Whether dialing to low-risk numbers is enabled | [optional] 
**Name** | Pointer to **NullableString** | The name of the country | [optional] 
**Url** | Pointer to **NullableString** | The absolute URL of this resource | [optional] 

## Methods

### NewVoiceV1DialingPermissionsDialingPermissionsCountry

`func NewVoiceV1DialingPermissionsDialingPermissionsCountry() *VoiceV1DialingPermissionsDialingPermissionsCountry`

NewVoiceV1DialingPermissionsDialingPermissionsCountry instantiates a new VoiceV1DialingPermissionsDialingPermissionsCountry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVoiceV1DialingPermissionsDialingPermissionsCountryWithDefaults

`func NewVoiceV1DialingPermissionsDialingPermissionsCountryWithDefaults() *VoiceV1DialingPermissionsDialingPermissionsCountry`

NewVoiceV1DialingPermissionsDialingPermissionsCountryWithDefaults instantiates a new VoiceV1DialingPermissionsDialingPermissionsCountry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContinent

`func (o *VoiceV1DialingPermissionsDialingPermissionsCountry) GetContinent() string`

GetContinent returns the Continent field if non-nil, zero value otherwise.

### GetContinentOk

`func (o *VoiceV1DialingPermissionsDialingPermissionsCountry) GetContinentOk() (*string, bool)`

GetContinentOk returns a tuple with the Continent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinent

`func (o *VoiceV1DialingPermissionsDialingPermissionsCountry) SetContinent(v string)`

SetContinent sets Continent field to given value.

### HasContinent

`func (o *VoiceV1DialingPermissionsDialingPermissionsCountry) HasContinent() bool`

HasContinent returns a boolean if a field has been set.

### SetContinentNil

`func (o *VoiceV1DialingPermissionsDialingPermissionsCountry) SetContinentNil(b bool)`

 SetContinentNil sets the value for Continent to be an explicit nil

### UnsetContinent
`func (o *VoiceV1DialingPermissionsDialingPermissionsCountry) UnsetContinent()`

UnsetContinent ensures that no value is present for Continent, not even an explicit nil
### GetCountryCodes

`func (o *VoiceV1DialingPermissionsDialingPermissionsCountry) GetCountryCodes() []string`

GetCountryCodes returns the CountryCodes field if non-nil, zero value otherwise.

### GetCountryCodesOk

`func (o *VoiceV1DialingPermissionsDialingPermissionsCountry) GetCountryCodesOk() (*[]string, bool)`

GetCountryCodesOk returns a tuple with the CountryCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCodes

`func (o *VoiceV1DialingPermissionsDialingPermissionsCountry) SetCountryCodes(v []string)`

SetCountryCodes sets CountryCodes field to given value.

### HasCountryCodes

`func (o *VoiceV1DialingPermissionsDialingPermissionsCountry) HasCountryCodes() bool`

HasCountryCodes returns a boolean if a field has been set.

### SetCountryCodesNil

`func (o *VoiceV1DialingPermissionsDialingPermissionsCountry) SetCountryCodesNil(b bool)`

 SetCountryCodesNil sets the value for CountryCodes to be an explicit nil

### UnsetCountryCodes
`func (o *VoiceV1DialingPermissionsDialingPermissionsCountry) UnsetCountryCodes()`

UnsetCountryCodes ensures that no value is present for CountryCodes, not even an explicit nil
### GetHighRiskSpecialNumbersEnabled

`func (o *VoiceV1DialingPermissionsDialingPermissionsCountry) GetHighRiskSpecialNumbersEnabled() bool`

GetHighRiskSpecialNumbersEnabled returns the HighRiskSpecialNumbersEnabled field if non-nil, zero value otherwise.

### GetHighRiskSpecialNumbersEnabledOk

`func (o *VoiceV1DialingPermissionsDialingPermissionsCountry) GetHighRiskSpecialNumbersEnabledOk() (*bool, bool)`

GetHighRiskSpecialNumbersEnabledOk returns a tuple with the HighRiskSpecialNumbersEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighRiskSpecialNumbersEnabled

`func (o *VoiceV1DialingPermissionsDialingPermissionsCountry) SetHighRiskSpecialNumbersEnabled(v bool)`

SetHighRiskSpecialNumbersEnabled sets HighRiskSpecialNumbersEnabled field to given value.

### HasHighRiskSpecialNumbersEnabled

`func (o *VoiceV1DialingPermissionsDialingPermissionsCountry) HasHighRiskSpecialNumbersEnabled() bool`

HasHighRiskSpecialNumbersEnabled returns a boolean if a field has been set.

### SetHighRiskSpecialNumbersEnabledNil

`func (o *VoiceV1DialingPermissionsDialingPermissionsCountry) SetHighRiskSpecialNumbersEnabledNil(b bool)`

 SetHighRiskSpecialNumbersEnabledNil sets the value for HighRiskSpecialNumbersEnabled to be an explicit nil

### UnsetHighRiskSpecialNumbersEnabled
`func (o *VoiceV1DialingPermissionsDialingPermissionsCountry) UnsetHighRiskSpecialNumbersEnabled()`

UnsetHighRiskSpecialNumbersEnabled ensures that no value is present for HighRiskSpecialNumbersEnabled, not even an explicit nil
### GetHighRiskTollfraudNumbersEnabled

`func (o *VoiceV1DialingPermissionsDialingPermissionsCountry) GetHighRiskTollfraudNumbersEnabled() bool`

GetHighRiskTollfraudNumbersEnabled returns the HighRiskTollfraudNumbersEnabled field if non-nil, zero value otherwise.

### GetHighRiskTollfraudNumbersEnabledOk

`func (o *VoiceV1DialingPermissionsDialingPermissionsCountry) GetHighRiskTollfraudNumbersEnabledOk() (*bool, bool)`

GetHighRiskTollfraudNumbersEnabledOk returns a tuple with the HighRiskTollfraudNumbersEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighRiskTollfraudNumbersEnabled

`func (o *VoiceV1DialingPermissionsDialingPermissionsCountry) SetHighRiskTollfraudNumbersEnabled(v bool)`

SetHighRiskTollfraudNumbersEnabled sets HighRiskTollfraudNumbersEnabled field to given value.

### HasHighRiskTollfraudNumbersEnabled

`func (o *VoiceV1DialingPermissionsDialingPermissionsCountry) HasHighRiskTollfraudNumbersEnabled() bool`

HasHighRiskTollfraudNumbersEnabled returns a boolean if a field has been set.

### SetHighRiskTollfraudNumbersEnabledNil

`func (o *VoiceV1DialingPermissionsDialingPermissionsCountry) SetHighRiskTollfraudNumbersEnabledNil(b bool)`

 SetHighRiskTollfraudNumbersEnabledNil sets the value for HighRiskTollfraudNumbersEnabled to be an explicit nil

### UnsetHighRiskTollfraudNumbersEnabled
`func (o *VoiceV1DialingPermissionsDialingPermissionsCountry) UnsetHighRiskTollfraudNumbersEnabled()`

UnsetHighRiskTollfraudNumbersEnabled ensures that no value is present for HighRiskTollfraudNumbersEnabled, not even an explicit nil
### GetIsoCode

`func (o *VoiceV1DialingPermissionsDialingPermissionsCountry) GetIsoCode() string`

GetIsoCode returns the IsoCode field if non-nil, zero value otherwise.

### GetIsoCodeOk

`func (o *VoiceV1DialingPermissionsDialingPermissionsCountry) GetIsoCodeOk() (*string, bool)`

GetIsoCodeOk returns a tuple with the IsoCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsoCode

`func (o *VoiceV1DialingPermissionsDialingPermissionsCountry) SetIsoCode(v string)`

SetIsoCode sets IsoCode field to given value.

### HasIsoCode

`func (o *VoiceV1DialingPermissionsDialingPermissionsCountry) HasIsoCode() bool`

HasIsoCode returns a boolean if a field has been set.

### SetIsoCodeNil

`func (o *VoiceV1DialingPermissionsDialingPermissionsCountry) SetIsoCodeNil(b bool)`

 SetIsoCodeNil sets the value for IsoCode to be an explicit nil

### UnsetIsoCode
`func (o *VoiceV1DialingPermissionsDialingPermissionsCountry) UnsetIsoCode()`

UnsetIsoCode ensures that no value is present for IsoCode, not even an explicit nil
### GetLinks

`func (o *VoiceV1DialingPermissionsDialingPermissionsCountry) GetLinks() map[string]interface{}`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *VoiceV1DialingPermissionsDialingPermissionsCountry) GetLinksOk() (*map[string]interface{}, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *VoiceV1DialingPermissionsDialingPermissionsCountry) SetLinks(v map[string]interface{})`

SetLinks sets Links field to given value.

### HasLinks

`func (o *VoiceV1DialingPermissionsDialingPermissionsCountry) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### SetLinksNil

`func (o *VoiceV1DialingPermissionsDialingPermissionsCountry) SetLinksNil(b bool)`

 SetLinksNil sets the value for Links to be an explicit nil

### UnsetLinks
`func (o *VoiceV1DialingPermissionsDialingPermissionsCountry) UnsetLinks()`

UnsetLinks ensures that no value is present for Links, not even an explicit nil
### GetLowRiskNumbersEnabled

`func (o *VoiceV1DialingPermissionsDialingPermissionsCountry) GetLowRiskNumbersEnabled() bool`

GetLowRiskNumbersEnabled returns the LowRiskNumbersEnabled field if non-nil, zero value otherwise.

### GetLowRiskNumbersEnabledOk

`func (o *VoiceV1DialingPermissionsDialingPermissionsCountry) GetLowRiskNumbersEnabledOk() (*bool, bool)`

GetLowRiskNumbersEnabledOk returns a tuple with the LowRiskNumbersEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowRiskNumbersEnabled

`func (o *VoiceV1DialingPermissionsDialingPermissionsCountry) SetLowRiskNumbersEnabled(v bool)`

SetLowRiskNumbersEnabled sets LowRiskNumbersEnabled field to given value.

### HasLowRiskNumbersEnabled

`func (o *VoiceV1DialingPermissionsDialingPermissionsCountry) HasLowRiskNumbersEnabled() bool`

HasLowRiskNumbersEnabled returns a boolean if a field has been set.

### SetLowRiskNumbersEnabledNil

`func (o *VoiceV1DialingPermissionsDialingPermissionsCountry) SetLowRiskNumbersEnabledNil(b bool)`

 SetLowRiskNumbersEnabledNil sets the value for LowRiskNumbersEnabled to be an explicit nil

### UnsetLowRiskNumbersEnabled
`func (o *VoiceV1DialingPermissionsDialingPermissionsCountry) UnsetLowRiskNumbersEnabled()`

UnsetLowRiskNumbersEnabled ensures that no value is present for LowRiskNumbersEnabled, not even an explicit nil
### GetName

`func (o *VoiceV1DialingPermissionsDialingPermissionsCountry) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VoiceV1DialingPermissionsDialingPermissionsCountry) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VoiceV1DialingPermissionsDialingPermissionsCountry) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VoiceV1DialingPermissionsDialingPermissionsCountry) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *VoiceV1DialingPermissionsDialingPermissionsCountry) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *VoiceV1DialingPermissionsDialingPermissionsCountry) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetUrl

`func (o *VoiceV1DialingPermissionsDialingPermissionsCountry) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *VoiceV1DialingPermissionsDialingPermissionsCountry) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *VoiceV1DialingPermissionsDialingPermissionsCountry) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *VoiceV1DialingPermissionsDialingPermissionsCountry) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *VoiceV1DialingPermissionsDialingPermissionsCountry) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *VoiceV1DialingPermissionsDialingPermissionsCountry) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


