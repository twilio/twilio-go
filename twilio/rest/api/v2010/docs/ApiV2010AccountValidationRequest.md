# ApiV2010AccountValidationRequest

## Properties

Name | Type | Description
------------ | ------------- | -------------
**AccountSid** | Pointer to **NullableString** | The SID of the Account that created the resource | [optional] 
**CallSid** | Pointer to **NullableString** | The SID of the Call the resource is associated with | [optional] 
**FriendlyName** | Pointer to **NullableString** | The string that you assigned to describe the resource | [optional] 
**PhoneNumber** | Pointer to **NullableString** | The phone number to verify in E.164 format | [optional] 
**ValidationCode** | Pointer to **NullableString** | The 6 digit validation code that someone must enter to validate the Caller ID  when &#x60;phone_number&#x60; is called | [optional] 

## Methods

### NewApiV2010AccountValidationRequest

`func NewApiV2010AccountValidationRequest() *ApiV2010AccountValidationRequest`

NewApiV2010AccountValidationRequest instantiates a new ApiV2010AccountValidationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiV2010AccountValidationRequestWithDefaults

`func NewApiV2010AccountValidationRequestWithDefaults() *ApiV2010AccountValidationRequest`

NewApiV2010AccountValidationRequestWithDefaults instantiates a new ApiV2010AccountValidationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSid

`func (o *ApiV2010AccountValidationRequest) GetAccountSid() string`

GetAccountSid returns the AccountSid field if non-nil, zero value otherwise.

### GetAccountSidOk

`func (o *ApiV2010AccountValidationRequest) GetAccountSidOk() (*string, bool)`

GetAccountSidOk returns a tuple with the AccountSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSid

`func (o *ApiV2010AccountValidationRequest) SetAccountSid(v string)`

SetAccountSid sets AccountSid field to given value.

### HasAccountSid

`func (o *ApiV2010AccountValidationRequest) HasAccountSid() bool`

HasAccountSid returns a boolean if a field has been set.

### SetAccountSidNil

`func (o *ApiV2010AccountValidationRequest) SetAccountSidNil(b bool)`

 SetAccountSidNil sets the value for AccountSid to be an explicit nil

### UnsetAccountSid
`func (o *ApiV2010AccountValidationRequest) UnsetAccountSid()`

UnsetAccountSid ensures that no value is present for AccountSid, not even an explicit nil
### GetCallSid

`func (o *ApiV2010AccountValidationRequest) GetCallSid() string`

GetCallSid returns the CallSid field if non-nil, zero value otherwise.

### GetCallSidOk

`func (o *ApiV2010AccountValidationRequest) GetCallSidOk() (*string, bool)`

GetCallSidOk returns a tuple with the CallSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallSid

`func (o *ApiV2010AccountValidationRequest) SetCallSid(v string)`

SetCallSid sets CallSid field to given value.

### HasCallSid

`func (o *ApiV2010AccountValidationRequest) HasCallSid() bool`

HasCallSid returns a boolean if a field has been set.

### SetCallSidNil

`func (o *ApiV2010AccountValidationRequest) SetCallSidNil(b bool)`

 SetCallSidNil sets the value for CallSid to be an explicit nil

### UnsetCallSid
`func (o *ApiV2010AccountValidationRequest) UnsetCallSid()`

UnsetCallSid ensures that no value is present for CallSid, not even an explicit nil
### GetFriendlyName

`func (o *ApiV2010AccountValidationRequest) GetFriendlyName() string`

GetFriendlyName returns the FriendlyName field if non-nil, zero value otherwise.

### GetFriendlyNameOk

`func (o *ApiV2010AccountValidationRequest) GetFriendlyNameOk() (*string, bool)`

GetFriendlyNameOk returns a tuple with the FriendlyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFriendlyName

`func (o *ApiV2010AccountValidationRequest) SetFriendlyName(v string)`

SetFriendlyName sets FriendlyName field to given value.

### HasFriendlyName

`func (o *ApiV2010AccountValidationRequest) HasFriendlyName() bool`

HasFriendlyName returns a boolean if a field has been set.

### SetFriendlyNameNil

`func (o *ApiV2010AccountValidationRequest) SetFriendlyNameNil(b bool)`

 SetFriendlyNameNil sets the value for FriendlyName to be an explicit nil

### UnsetFriendlyName
`func (o *ApiV2010AccountValidationRequest) UnsetFriendlyName()`

UnsetFriendlyName ensures that no value is present for FriendlyName, not even an explicit nil
### GetPhoneNumber

`func (o *ApiV2010AccountValidationRequest) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *ApiV2010AccountValidationRequest) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *ApiV2010AccountValidationRequest) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *ApiV2010AccountValidationRequest) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### SetPhoneNumberNil

`func (o *ApiV2010AccountValidationRequest) SetPhoneNumberNil(b bool)`

 SetPhoneNumberNil sets the value for PhoneNumber to be an explicit nil

### UnsetPhoneNumber
`func (o *ApiV2010AccountValidationRequest) UnsetPhoneNumber()`

UnsetPhoneNumber ensures that no value is present for PhoneNumber, not even an explicit nil
### GetValidationCode

`func (o *ApiV2010AccountValidationRequest) GetValidationCode() string`

GetValidationCode returns the ValidationCode field if non-nil, zero value otherwise.

### GetValidationCodeOk

`func (o *ApiV2010AccountValidationRequest) GetValidationCodeOk() (*string, bool)`

GetValidationCodeOk returns a tuple with the ValidationCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationCode

`func (o *ApiV2010AccountValidationRequest) SetValidationCode(v string)`

SetValidationCode sets ValidationCode field to given value.

### HasValidationCode

`func (o *ApiV2010AccountValidationRequest) HasValidationCode() bool`

HasValidationCode returns a boolean if a field has been set.

### SetValidationCodeNil

`func (o *ApiV2010AccountValidationRequest) SetValidationCodeNil(b bool)`

 SetValidationCodeNil sets the value for ValidationCode to be an explicit nil

### UnsetValidationCode
`func (o *ApiV2010AccountValidationRequest) UnsetValidationCode()`

UnsetValidationCode ensures that no value is present for ValidationCode, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


