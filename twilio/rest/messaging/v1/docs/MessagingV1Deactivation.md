# MessagingV1Deactivation

## Properties

Name | Type | Description
------------ | ------------- | -------------
**RedirectTo** | Pointer to **NullableString** | Redirect url to the list of deactivated numbers. | [optional] 

## Methods

### NewMessagingV1Deactivation

`func NewMessagingV1Deactivation() *MessagingV1Deactivation`

NewMessagingV1Deactivation instantiates a new MessagingV1Deactivation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMessagingV1DeactivationWithDefaults

`func NewMessagingV1DeactivationWithDefaults() *MessagingV1Deactivation`

NewMessagingV1DeactivationWithDefaults instantiates a new MessagingV1Deactivation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRedirectTo

`func (o *MessagingV1Deactivation) GetRedirectTo() string`

GetRedirectTo returns the RedirectTo field if non-nil, zero value otherwise.

### GetRedirectToOk

`func (o *MessagingV1Deactivation) GetRedirectToOk() (*string, bool)`

GetRedirectToOk returns a tuple with the RedirectTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectTo

`func (o *MessagingV1Deactivation) SetRedirectTo(v string)`

SetRedirectTo sets RedirectTo field to given value.

### HasRedirectTo

`func (o *MessagingV1Deactivation) HasRedirectTo() bool`

HasRedirectTo returns a boolean if a field has been set.

### SetRedirectToNil

`func (o *MessagingV1Deactivation) SetRedirectToNil(b bool)`

 SetRedirectToNil sets the value for RedirectTo to be an explicit nil

### UnsetRedirectTo
`func (o *MessagingV1Deactivation) UnsetRedirectTo()`

UnsetRedirectTo ensures that no value is present for RedirectTo, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


