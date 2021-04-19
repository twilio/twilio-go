# ListMessagingCountryResponse

## Properties

Name | Type | Description
------------ | ------------- | -------------
**Countries** | Pointer to [**[]PricingV1MessagingMessagingCountry**](PricingV1MessagingMessagingCountry.md) |  | [optional] 
**Meta** | Pointer to [**ListMessagingCountryResponseMeta**](ListMessagingCountryResponse_meta.md) |  | [optional] 

## Methods

### NewListMessagingCountryResponse

`func NewListMessagingCountryResponse() *ListMessagingCountryResponse`

NewListMessagingCountryResponse instantiates a new ListMessagingCountryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListMessagingCountryResponseWithDefaults

`func NewListMessagingCountryResponseWithDefaults() *ListMessagingCountryResponse`

NewListMessagingCountryResponseWithDefaults instantiates a new ListMessagingCountryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCountries

`func (o *ListMessagingCountryResponse) GetCountries() []PricingV1MessagingMessagingCountry`

GetCountries returns the Countries field if non-nil, zero value otherwise.

### GetCountriesOk

`func (o *ListMessagingCountryResponse) GetCountriesOk() (*[]PricingV1MessagingMessagingCountry, bool)`

GetCountriesOk returns a tuple with the Countries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountries

`func (o *ListMessagingCountryResponse) SetCountries(v []PricingV1MessagingMessagingCountry)`

SetCountries sets Countries field to given value.

### HasCountries

`func (o *ListMessagingCountryResponse) HasCountries() bool`

HasCountries returns a boolean if a field has been set.

### GetMeta

`func (o *ListMessagingCountryResponse) GetMeta() ListMessagingCountryResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListMessagingCountryResponse) GetMetaOk() (*ListMessagingCountryResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListMessagingCountryResponse) SetMeta(v ListMessagingCountryResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListMessagingCountryResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


