# ListPhoneNumberResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**ListServiceResponseMeta**](ListServiceResponse_meta.md) |  | [optional] 
**PhoneNumbers** | Pointer to [**[]ProxyV1ServicePhoneNumber**](ProxyV1ServicePhoneNumber.md) |  | [optional] 

## Methods

### NewListPhoneNumberResponse

`func NewListPhoneNumberResponse() *ListPhoneNumberResponse`

NewListPhoneNumberResponse instantiates a new ListPhoneNumberResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListPhoneNumberResponseWithDefaults

`func NewListPhoneNumberResponseWithDefaults() *ListPhoneNumberResponse`

NewListPhoneNumberResponseWithDefaults instantiates a new ListPhoneNumberResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *ListPhoneNumberResponse) GetMeta() ListServiceResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListPhoneNumberResponse) GetMetaOk() (*ListServiceResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListPhoneNumberResponse) SetMeta(v ListServiceResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListPhoneNumberResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetPhoneNumbers

`func (o *ListPhoneNumberResponse) GetPhoneNumbers() []ProxyV1ServicePhoneNumber`

GetPhoneNumbers returns the PhoneNumbers field if non-nil, zero value otherwise.

### GetPhoneNumbersOk

`func (o *ListPhoneNumberResponse) GetPhoneNumbersOk() (*[]ProxyV1ServicePhoneNumber, bool)`

GetPhoneNumbersOk returns a tuple with the PhoneNumbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumbers

`func (o *ListPhoneNumberResponse) SetPhoneNumbers(v []ProxyV1ServicePhoneNumber)`

SetPhoneNumbers sets PhoneNumbers field to given value.

### HasPhoneNumbers

`func (o *ListPhoneNumberResponse) HasPhoneNumbers() bool`

HasPhoneNumbers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


