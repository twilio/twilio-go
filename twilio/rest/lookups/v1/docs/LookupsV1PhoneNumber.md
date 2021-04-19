# LookupsV1PhoneNumber

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddOns** | Pointer to **map[string]interface{}** | A JSON string with the results of the Add-ons you specified | [optional] 
**CallerName** | Pointer to **map[string]interface{}** | The name of the phone number&#39;s owner | [optional] 
**Carrier** | Pointer to **map[string]interface{}** | The telecom company that provides the phone number | [optional] 
**CountryCode** | Pointer to **NullableString** | The ISO country code for the phone number | [optional] 
**NationalFormat** | Pointer to **NullableString** | The phone number, in national format | [optional] 
**PhoneNumber** | Pointer to **NullableString** | The phone number in E.164 format | [optional] 
**Url** | Pointer to **NullableString** | The absolute URL of the resource | [optional] 

## Methods

### NewLookupsV1PhoneNumber

`func NewLookupsV1PhoneNumber() *LookupsV1PhoneNumber`

NewLookupsV1PhoneNumber instantiates a new LookupsV1PhoneNumber object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLookupsV1PhoneNumberWithDefaults

`func NewLookupsV1PhoneNumberWithDefaults() *LookupsV1PhoneNumber`

NewLookupsV1PhoneNumberWithDefaults instantiates a new LookupsV1PhoneNumber object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddOns

`func (o *LookupsV1PhoneNumber) GetAddOns() map[string]interface{}`

GetAddOns returns the AddOns field if non-nil, zero value otherwise.

### GetAddOnsOk

`func (o *LookupsV1PhoneNumber) GetAddOnsOk() (*map[string]interface{}, bool)`

GetAddOnsOk returns a tuple with the AddOns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddOns

`func (o *LookupsV1PhoneNumber) SetAddOns(v map[string]interface{})`

SetAddOns sets AddOns field to given value.

### HasAddOns

`func (o *LookupsV1PhoneNumber) HasAddOns() bool`

HasAddOns returns a boolean if a field has been set.

### SetAddOnsNil

`func (o *LookupsV1PhoneNumber) SetAddOnsNil(b bool)`

 SetAddOnsNil sets the value for AddOns to be an explicit nil

### UnsetAddOns
`func (o *LookupsV1PhoneNumber) UnsetAddOns()`

UnsetAddOns ensures that no value is present for AddOns, not even an explicit nil
### GetCallerName

`func (o *LookupsV1PhoneNumber) GetCallerName() map[string]interface{}`

GetCallerName returns the CallerName field if non-nil, zero value otherwise.

### GetCallerNameOk

`func (o *LookupsV1PhoneNumber) GetCallerNameOk() (*map[string]interface{}, bool)`

GetCallerNameOk returns a tuple with the CallerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallerName

`func (o *LookupsV1PhoneNumber) SetCallerName(v map[string]interface{})`

SetCallerName sets CallerName field to given value.

### HasCallerName

`func (o *LookupsV1PhoneNumber) HasCallerName() bool`

HasCallerName returns a boolean if a field has been set.

### SetCallerNameNil

`func (o *LookupsV1PhoneNumber) SetCallerNameNil(b bool)`

 SetCallerNameNil sets the value for CallerName to be an explicit nil

### UnsetCallerName
`func (o *LookupsV1PhoneNumber) UnsetCallerName()`

UnsetCallerName ensures that no value is present for CallerName, not even an explicit nil
### GetCarrier

`func (o *LookupsV1PhoneNumber) GetCarrier() map[string]interface{}`

GetCarrier returns the Carrier field if non-nil, zero value otherwise.

### GetCarrierOk

`func (o *LookupsV1PhoneNumber) GetCarrierOk() (*map[string]interface{}, bool)`

GetCarrierOk returns a tuple with the Carrier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrier

`func (o *LookupsV1PhoneNumber) SetCarrier(v map[string]interface{})`

SetCarrier sets Carrier field to given value.

### HasCarrier

`func (o *LookupsV1PhoneNumber) HasCarrier() bool`

HasCarrier returns a boolean if a field has been set.

### SetCarrierNil

`func (o *LookupsV1PhoneNumber) SetCarrierNil(b bool)`

 SetCarrierNil sets the value for Carrier to be an explicit nil

### UnsetCarrier
`func (o *LookupsV1PhoneNumber) UnsetCarrier()`

UnsetCarrier ensures that no value is present for Carrier, not even an explicit nil
### GetCountryCode

`func (o *LookupsV1PhoneNumber) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *LookupsV1PhoneNumber) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *LookupsV1PhoneNumber) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *LookupsV1PhoneNumber) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### SetCountryCodeNil

`func (o *LookupsV1PhoneNumber) SetCountryCodeNil(b bool)`

 SetCountryCodeNil sets the value for CountryCode to be an explicit nil

### UnsetCountryCode
`func (o *LookupsV1PhoneNumber) UnsetCountryCode()`

UnsetCountryCode ensures that no value is present for CountryCode, not even an explicit nil
### GetNationalFormat

`func (o *LookupsV1PhoneNumber) GetNationalFormat() string`

GetNationalFormat returns the NationalFormat field if non-nil, zero value otherwise.

### GetNationalFormatOk

`func (o *LookupsV1PhoneNumber) GetNationalFormatOk() (*string, bool)`

GetNationalFormatOk returns a tuple with the NationalFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNationalFormat

`func (o *LookupsV1PhoneNumber) SetNationalFormat(v string)`

SetNationalFormat sets NationalFormat field to given value.

### HasNationalFormat

`func (o *LookupsV1PhoneNumber) HasNationalFormat() bool`

HasNationalFormat returns a boolean if a field has been set.

### SetNationalFormatNil

`func (o *LookupsV1PhoneNumber) SetNationalFormatNil(b bool)`

 SetNationalFormatNil sets the value for NationalFormat to be an explicit nil

### UnsetNationalFormat
`func (o *LookupsV1PhoneNumber) UnsetNationalFormat()`

UnsetNationalFormat ensures that no value is present for NationalFormat, not even an explicit nil
### GetPhoneNumber

`func (o *LookupsV1PhoneNumber) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *LookupsV1PhoneNumber) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *LookupsV1PhoneNumber) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *LookupsV1PhoneNumber) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### SetPhoneNumberNil

`func (o *LookupsV1PhoneNumber) SetPhoneNumberNil(b bool)`

 SetPhoneNumberNil sets the value for PhoneNumber to be an explicit nil

### UnsetPhoneNumber
`func (o *LookupsV1PhoneNumber) UnsetPhoneNumber()`

UnsetPhoneNumber ensures that no value is present for PhoneNumber, not even an explicit nil
### GetUrl

`func (o *LookupsV1PhoneNumber) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *LookupsV1PhoneNumber) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *LookupsV1PhoneNumber) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *LookupsV1PhoneNumber) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *LookupsV1PhoneNumber) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *LookupsV1PhoneNumber) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


