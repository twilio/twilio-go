# ListPhoneNumberCountryResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Countries** | Pointer to [**[]PricingV1PhoneNumberPhoneNumberCountry**](PricingV1PhoneNumberPhoneNumberCountry.md) |  | [optional] 
**Meta** | Pointer to [**ListMessagingCountryResponseMeta**](ListMessagingCountryResponse_meta.md) |  | [optional] 

## Methods

### NewListPhoneNumberCountryResponse

`func NewListPhoneNumberCountryResponse() *ListPhoneNumberCountryResponse`

NewListPhoneNumberCountryResponse instantiates a new ListPhoneNumberCountryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListPhoneNumberCountryResponseWithDefaults

`func NewListPhoneNumberCountryResponseWithDefaults() *ListPhoneNumberCountryResponse`

NewListPhoneNumberCountryResponseWithDefaults instantiates a new ListPhoneNumberCountryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCountries

`func (o *ListPhoneNumberCountryResponse) GetCountries() []PricingV1PhoneNumberPhoneNumberCountry`

GetCountries returns the Countries field if non-nil, zero value otherwise.

### GetCountriesOk

`func (o *ListPhoneNumberCountryResponse) GetCountriesOk() (*[]PricingV1PhoneNumberPhoneNumberCountry, bool)`

GetCountriesOk returns a tuple with the Countries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountries

`func (o *ListPhoneNumberCountryResponse) SetCountries(v []PricingV1PhoneNumberPhoneNumberCountry)`

SetCountries sets Countries field to given value.

### HasCountries

`func (o *ListPhoneNumberCountryResponse) HasCountries() bool`

HasCountries returns a boolean if a field has been set.

### GetMeta

`func (o *ListPhoneNumberCountryResponse) GetMeta() ListMessagingCountryResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListPhoneNumberCountryResponse) GetMetaOk() (*ListMessagingCountryResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListPhoneNumberCountryResponse) SetMeta(v ListMessagingCountryResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListPhoneNumberCountryResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


