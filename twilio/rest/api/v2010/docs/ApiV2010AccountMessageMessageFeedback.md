# ApiV2010AccountMessageMessageFeedback

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **NullableString** | The SID of the Account that created the resource | [optional] 
**DateCreated** | Pointer to **NullableString** | The RFC 2822 date and time in GMT that the resource was created | [optional] 
**DateUpdated** | Pointer to **NullableString** | The RFC 2822 date and time in GMT that the resource was last updated | [optional] 
**MessageSid** | Pointer to **NullableString** | The SID of the Message resource for which the feedback was provided | [optional] 
**Outcome** | Pointer to **NullableString** | Whether the feedback has arrived | [optional] 
**Uri** | Pointer to **NullableString** | The URI of the resource, relative to &#x60;https://api.twilio.com&#x60; | [optional] 

## Methods

### NewApiV2010AccountMessageMessageFeedback

`func NewApiV2010AccountMessageMessageFeedback() *ApiV2010AccountMessageMessageFeedback`

NewApiV2010AccountMessageMessageFeedback instantiates a new ApiV2010AccountMessageMessageFeedback object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiV2010AccountMessageMessageFeedbackWithDefaults

`func NewApiV2010AccountMessageMessageFeedbackWithDefaults() *ApiV2010AccountMessageMessageFeedback`

NewApiV2010AccountMessageMessageFeedbackWithDefaults instantiates a new ApiV2010AccountMessageMessageFeedback object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSid

`func (o *ApiV2010AccountMessageMessageFeedback) GetAccountSid() string`

GetAccountSid returns the AccountSid field if non-nil, zero value otherwise.

### GetAccountSidOk

`func (o *ApiV2010AccountMessageMessageFeedback) GetAccountSidOk() (*string, bool)`

GetAccountSidOk returns a tuple with the AccountSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSid

`func (o *ApiV2010AccountMessageMessageFeedback) SetAccountSid(v string)`

SetAccountSid sets AccountSid field to given value.

### HasAccountSid

`func (o *ApiV2010AccountMessageMessageFeedback) HasAccountSid() bool`

HasAccountSid returns a boolean if a field has been set.

### SetAccountSidNil

`func (o *ApiV2010AccountMessageMessageFeedback) SetAccountSidNil(b bool)`

 SetAccountSidNil sets the value for AccountSid to be an explicit nil

### UnsetAccountSid
`func (o *ApiV2010AccountMessageMessageFeedback) UnsetAccountSid()`

UnsetAccountSid ensures that no value is present for AccountSid, not even an explicit nil
### GetDateCreated

`func (o *ApiV2010AccountMessageMessageFeedback) GetDateCreated() string`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *ApiV2010AccountMessageMessageFeedback) GetDateCreatedOk() (*string, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *ApiV2010AccountMessageMessageFeedback) SetDateCreated(v string)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *ApiV2010AccountMessageMessageFeedback) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### SetDateCreatedNil

`func (o *ApiV2010AccountMessageMessageFeedback) SetDateCreatedNil(b bool)`

 SetDateCreatedNil sets the value for DateCreated to be an explicit nil

### UnsetDateCreated
`func (o *ApiV2010AccountMessageMessageFeedback) UnsetDateCreated()`

UnsetDateCreated ensures that no value is present for DateCreated, not even an explicit nil
### GetDateUpdated

`func (o *ApiV2010AccountMessageMessageFeedback) GetDateUpdated() string`

GetDateUpdated returns the DateUpdated field if non-nil, zero value otherwise.

### GetDateUpdatedOk

`func (o *ApiV2010AccountMessageMessageFeedback) GetDateUpdatedOk() (*string, bool)`

GetDateUpdatedOk returns a tuple with the DateUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateUpdated

`func (o *ApiV2010AccountMessageMessageFeedback) SetDateUpdated(v string)`

SetDateUpdated sets DateUpdated field to given value.

### HasDateUpdated

`func (o *ApiV2010AccountMessageMessageFeedback) HasDateUpdated() bool`

HasDateUpdated returns a boolean if a field has been set.

### SetDateUpdatedNil

`func (o *ApiV2010AccountMessageMessageFeedback) SetDateUpdatedNil(b bool)`

 SetDateUpdatedNil sets the value for DateUpdated to be an explicit nil

### UnsetDateUpdated
`func (o *ApiV2010AccountMessageMessageFeedback) UnsetDateUpdated()`

UnsetDateUpdated ensures that no value is present for DateUpdated, not even an explicit nil
### GetMessageSid

`func (o *ApiV2010AccountMessageMessageFeedback) GetMessageSid() string`

GetMessageSid returns the MessageSid field if non-nil, zero value otherwise.

### GetMessageSidOk

`func (o *ApiV2010AccountMessageMessageFeedback) GetMessageSidOk() (*string, bool)`

GetMessageSidOk returns a tuple with the MessageSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageSid

`func (o *ApiV2010AccountMessageMessageFeedback) SetMessageSid(v string)`

SetMessageSid sets MessageSid field to given value.

### HasMessageSid

`func (o *ApiV2010AccountMessageMessageFeedback) HasMessageSid() bool`

HasMessageSid returns a boolean if a field has been set.

### SetMessageSidNil

`func (o *ApiV2010AccountMessageMessageFeedback) SetMessageSidNil(b bool)`

 SetMessageSidNil sets the value for MessageSid to be an explicit nil

### UnsetMessageSid
`func (o *ApiV2010AccountMessageMessageFeedback) UnsetMessageSid()`

UnsetMessageSid ensures that no value is present for MessageSid, not even an explicit nil
### GetOutcome

`func (o *ApiV2010AccountMessageMessageFeedback) GetOutcome() string`

GetOutcome returns the Outcome field if non-nil, zero value otherwise.

### GetOutcomeOk

`func (o *ApiV2010AccountMessageMessageFeedback) GetOutcomeOk() (*string, bool)`

GetOutcomeOk returns a tuple with the Outcome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutcome

`func (o *ApiV2010AccountMessageMessageFeedback) SetOutcome(v string)`

SetOutcome sets Outcome field to given value.

### HasOutcome

`func (o *ApiV2010AccountMessageMessageFeedback) HasOutcome() bool`

HasOutcome returns a boolean if a field has been set.

### SetOutcomeNil

`func (o *ApiV2010AccountMessageMessageFeedback) SetOutcomeNil(b bool)`

 SetOutcomeNil sets the value for Outcome to be an explicit nil

### UnsetOutcome
`func (o *ApiV2010AccountMessageMessageFeedback) UnsetOutcome()`

UnsetOutcome ensures that no value is present for Outcome, not even an explicit nil
### GetUri

`func (o *ApiV2010AccountMessageMessageFeedback) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *ApiV2010AccountMessageMessageFeedback) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *ApiV2010AccountMessageMessageFeedback) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *ApiV2010AccountMessageMessageFeedback) HasUri() bool`

HasUri returns a boolean if a field has been set.

### SetUriNil

`func (o *ApiV2010AccountMessageMessageFeedback) SetUriNil(b bool)`

 SetUriNil sets the value for Uri to be an explicit nil

### UnsetUri
`func (o *ApiV2010AccountMessageMessageFeedback) UnsetUri()`

UnsetUri ensures that no value is present for Uri, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


