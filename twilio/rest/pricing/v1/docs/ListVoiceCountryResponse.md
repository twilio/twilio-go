# ListVoiceCountryResponse

## Properties

Name | Type | Description
------------ | ------------- | -------------
**Countries** | Pointer to [**[]PricingV1VoiceVoiceCountry**](PricingV1VoiceVoiceCountry.md) |  | [optional] 
**Meta** | Pointer to [**ListMessagingCountryResponseMeta**](ListMessagingCountryResponse_meta.md) |  | [optional] 

## Methods

### NewListVoiceCountryResponse

`func NewListVoiceCountryResponse() *ListVoiceCountryResponse`

NewListVoiceCountryResponse instantiates a new ListVoiceCountryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListVoiceCountryResponseWithDefaults

`func NewListVoiceCountryResponseWithDefaults() *ListVoiceCountryResponse`

NewListVoiceCountryResponseWithDefaults instantiates a new ListVoiceCountryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCountries

`func (o *ListVoiceCountryResponse) GetCountries() []PricingV1VoiceVoiceCountry`

GetCountries returns the Countries field if non-nil, zero value otherwise.

### GetCountriesOk

`func (o *ListVoiceCountryResponse) GetCountriesOk() (*[]PricingV1VoiceVoiceCountry, bool)`

GetCountriesOk returns a tuple with the Countries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountries

`func (o *ListVoiceCountryResponse) SetCountries(v []PricingV1VoiceVoiceCountry)`

SetCountries sets Countries field to given value.

### HasCountries

`func (o *ListVoiceCountryResponse) HasCountries() bool`

HasCountries returns a boolean if a field has been set.

### GetMeta

`func (o *ListVoiceCountryResponse) GetMeta() ListMessagingCountryResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListVoiceCountryResponse) GetMetaOk() (*ListMessagingCountryResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListVoiceCountryResponse) SetMeta(v ListMessagingCountryResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListVoiceCountryResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


