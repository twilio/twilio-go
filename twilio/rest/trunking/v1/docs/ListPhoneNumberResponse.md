# ListPhoneNumberResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**ListTrunkResponseMeta**](ListTrunkResponse_meta.md) |  | [optional] 
**PhoneNumbers** | Pointer to [**[]TrunkingV1TrunkPhoneNumber**](TrunkingV1TrunkPhoneNumber.md) |  | [optional] 

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

`func (o *ListPhoneNumberResponse) GetMeta() ListTrunkResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListPhoneNumberResponse) GetMetaOk() (*ListTrunkResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListPhoneNumberResponse) SetMeta(v ListTrunkResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListPhoneNumberResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetPhoneNumbers

`func (o *ListPhoneNumberResponse) GetPhoneNumbers() []TrunkingV1TrunkPhoneNumber`

GetPhoneNumbers returns the PhoneNumbers field if non-nil, zero value otherwise.

### GetPhoneNumbersOk

`func (o *ListPhoneNumberResponse) GetPhoneNumbersOk() (*[]TrunkingV1TrunkPhoneNumber, bool)`

GetPhoneNumbersOk returns a tuple with the PhoneNumbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumbers

`func (o *ListPhoneNumberResponse) SetPhoneNumbers(v []TrunkingV1TrunkPhoneNumber)`

SetPhoneNumbers sets PhoneNumbers field to given value.

### HasPhoneNumbers

`func (o *ListPhoneNumberResponse) HasPhoneNumbers() bool`

HasPhoneNumbers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


