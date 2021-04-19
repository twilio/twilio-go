# ApiV2010AccountMessage

## Properties

Name | Type | Description
------------ | ------------- | -------------
**AccountSid** | Pointer to **NullableString** | The SID of the Account that created the resource | [optional] 
**ApiVersion** | Pointer to **NullableString** | The API version used to process the message | [optional] 
**Body** | Pointer to **NullableString** | The message text | [optional] 
**DateCreated** | Pointer to **NullableString** | The RFC 2822 date and time in GMT that the resource was created | [optional] 
**DateSent** | Pointer to **NullableString** | The RFC 2822 date and time in GMT when the message was sent | [optional] 
**DateUpdated** | Pointer to **NullableString** | The RFC 2822 date and time in GMT that the resource was last updated | [optional] 
**Direction** | Pointer to **NullableString** | The direction of the message | [optional] 
**ErrorCode** | Pointer to **NullableInt32** | The error code associated with the message | [optional] 
**ErrorMessage** | Pointer to **NullableString** | The description of the error_code | [optional] 
**From** | Pointer to **NullableString** | The phone number that initiated the message | [optional] 
**MessagingServiceSid** | Pointer to **NullableString** | The SID of the Messaging Service used with the message. | [optional] 
**NumMedia** | Pointer to **NullableString** | The number of media files associated with the message | [optional] 
**NumSegments** | Pointer to **NullableString** | The number of messages used to deliver the message body | [optional] 
**Price** | Pointer to **NullableString** | The amount billed for the message | [optional] 
**PriceUnit** | Pointer to **NullableString** | The currency in which price is measured | [optional] 
**Sid** | Pointer to **NullableString** | The unique string that identifies the resource | [optional] 
**Status** | Pointer to **NullableString** | The status of the message | [optional] 
**SubresourceUris** | Pointer to **map[string]interface{}** | A list of related resources identified by their relative URIs | [optional] 
**To** | Pointer to **NullableString** | The phone number that received the message | [optional] 
**Uri** | Pointer to **NullableString** | The URI of the resource, relative to &#x60;https://api.twilio.com&#x60; | [optional] 

## Methods

### NewApiV2010AccountMessage

`func NewApiV2010AccountMessage() *ApiV2010AccountMessage`

NewApiV2010AccountMessage instantiates a new ApiV2010AccountMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiV2010AccountMessageWithDefaults

`func NewApiV2010AccountMessageWithDefaults() *ApiV2010AccountMessage`

NewApiV2010AccountMessageWithDefaults instantiates a new ApiV2010AccountMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSid

`func (o *ApiV2010AccountMessage) GetAccountSid() string`

GetAccountSid returns the AccountSid field if non-nil, zero value otherwise.

### GetAccountSidOk

`func (o *ApiV2010AccountMessage) GetAccountSidOk() (*string, bool)`

GetAccountSidOk returns a tuple with the AccountSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSid

`func (o *ApiV2010AccountMessage) SetAccountSid(v string)`

SetAccountSid sets AccountSid field to given value.

### HasAccountSid

`func (o *ApiV2010AccountMessage) HasAccountSid() bool`

HasAccountSid returns a boolean if a field has been set.

### SetAccountSidNil

`func (o *ApiV2010AccountMessage) SetAccountSidNil(b bool)`

 SetAccountSidNil sets the value for AccountSid to be an explicit nil

### UnsetAccountSid
`func (o *ApiV2010AccountMessage) UnsetAccountSid()`

UnsetAccountSid ensures that no value is present for AccountSid, not even an explicit nil
### GetApiVersion

`func (o *ApiV2010AccountMessage) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *ApiV2010AccountMessage) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *ApiV2010AccountMessage) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *ApiV2010AccountMessage) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### SetApiVersionNil

`func (o *ApiV2010AccountMessage) SetApiVersionNil(b bool)`

 SetApiVersionNil sets the value for ApiVersion to be an explicit nil

### UnsetApiVersion
`func (o *ApiV2010AccountMessage) UnsetApiVersion()`

UnsetApiVersion ensures that no value is present for ApiVersion, not even an explicit nil
### GetBody

`func (o *ApiV2010AccountMessage) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *ApiV2010AccountMessage) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *ApiV2010AccountMessage) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *ApiV2010AccountMessage) HasBody() bool`

HasBody returns a boolean if a field has been set.

### SetBodyNil

`func (o *ApiV2010AccountMessage) SetBodyNil(b bool)`

 SetBodyNil sets the value for Body to be an explicit nil

### UnsetBody
`func (o *ApiV2010AccountMessage) UnsetBody()`

UnsetBody ensures that no value is present for Body, not even an explicit nil
### GetDateCreated

`func (o *ApiV2010AccountMessage) GetDateCreated() string`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *ApiV2010AccountMessage) GetDateCreatedOk() (*string, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *ApiV2010AccountMessage) SetDateCreated(v string)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *ApiV2010AccountMessage) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### SetDateCreatedNil

`func (o *ApiV2010AccountMessage) SetDateCreatedNil(b bool)`

 SetDateCreatedNil sets the value for DateCreated to be an explicit nil

### UnsetDateCreated
`func (o *ApiV2010AccountMessage) UnsetDateCreated()`

UnsetDateCreated ensures that no value is present for DateCreated, not even an explicit nil
### GetDateSent

`func (o *ApiV2010AccountMessage) GetDateSent() string`

GetDateSent returns the DateSent field if non-nil, zero value otherwise.

### GetDateSentOk

`func (o *ApiV2010AccountMessage) GetDateSentOk() (*string, bool)`

GetDateSentOk returns a tuple with the DateSent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateSent

`func (o *ApiV2010AccountMessage) SetDateSent(v string)`

SetDateSent sets DateSent field to given value.

### HasDateSent

`func (o *ApiV2010AccountMessage) HasDateSent() bool`

HasDateSent returns a boolean if a field has been set.

### SetDateSentNil

`func (o *ApiV2010AccountMessage) SetDateSentNil(b bool)`

 SetDateSentNil sets the value for DateSent to be an explicit nil

### UnsetDateSent
`func (o *ApiV2010AccountMessage) UnsetDateSent()`

UnsetDateSent ensures that no value is present for DateSent, not even an explicit nil
### GetDateUpdated

`func (o *ApiV2010AccountMessage) GetDateUpdated() string`

GetDateUpdated returns the DateUpdated field if non-nil, zero value otherwise.

### GetDateUpdatedOk

`func (o *ApiV2010AccountMessage) GetDateUpdatedOk() (*string, bool)`

GetDateUpdatedOk returns a tuple with the DateUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateUpdated

`func (o *ApiV2010AccountMessage) SetDateUpdated(v string)`

SetDateUpdated sets DateUpdated field to given value.

### HasDateUpdated

`func (o *ApiV2010AccountMessage) HasDateUpdated() bool`

HasDateUpdated returns a boolean if a field has been set.

### SetDateUpdatedNil

`func (o *ApiV2010AccountMessage) SetDateUpdatedNil(b bool)`

 SetDateUpdatedNil sets the value for DateUpdated to be an explicit nil

### UnsetDateUpdated
`func (o *ApiV2010AccountMessage) UnsetDateUpdated()`

UnsetDateUpdated ensures that no value is present for DateUpdated, not even an explicit nil
### GetDirection

`func (o *ApiV2010AccountMessage) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *ApiV2010AccountMessage) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *ApiV2010AccountMessage) SetDirection(v string)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *ApiV2010AccountMessage) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### SetDirectionNil

`func (o *ApiV2010AccountMessage) SetDirectionNil(b bool)`

 SetDirectionNil sets the value for Direction to be an explicit nil

### UnsetDirection
`func (o *ApiV2010AccountMessage) UnsetDirection()`

UnsetDirection ensures that no value is present for Direction, not even an explicit nil
### GetErrorCode

`func (o *ApiV2010AccountMessage) GetErrorCode() int32`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *ApiV2010AccountMessage) GetErrorCodeOk() (*int32, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *ApiV2010AccountMessage) SetErrorCode(v int32)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *ApiV2010AccountMessage) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### SetErrorCodeNil

`func (o *ApiV2010AccountMessage) SetErrorCodeNil(b bool)`

 SetErrorCodeNil sets the value for ErrorCode to be an explicit nil

### UnsetErrorCode
`func (o *ApiV2010AccountMessage) UnsetErrorCode()`

UnsetErrorCode ensures that no value is present for ErrorCode, not even an explicit nil
### GetErrorMessage

`func (o *ApiV2010AccountMessage) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *ApiV2010AccountMessage) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *ApiV2010AccountMessage) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *ApiV2010AccountMessage) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### SetErrorMessageNil

`func (o *ApiV2010AccountMessage) SetErrorMessageNil(b bool)`

 SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil

### UnsetErrorMessage
`func (o *ApiV2010AccountMessage) UnsetErrorMessage()`

UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil
### GetFrom

`func (o *ApiV2010AccountMessage) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *ApiV2010AccountMessage) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *ApiV2010AccountMessage) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *ApiV2010AccountMessage) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### SetFromNil

`func (o *ApiV2010AccountMessage) SetFromNil(b bool)`

 SetFromNil sets the value for From to be an explicit nil

### UnsetFrom
`func (o *ApiV2010AccountMessage) UnsetFrom()`

UnsetFrom ensures that no value is present for From, not even an explicit nil
### GetMessagingServiceSid

`func (o *ApiV2010AccountMessage) GetMessagingServiceSid() string`

GetMessagingServiceSid returns the MessagingServiceSid field if non-nil, zero value otherwise.

### GetMessagingServiceSidOk

`func (o *ApiV2010AccountMessage) GetMessagingServiceSidOk() (*string, bool)`

GetMessagingServiceSidOk returns a tuple with the MessagingServiceSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessagingServiceSid

`func (o *ApiV2010AccountMessage) SetMessagingServiceSid(v string)`

SetMessagingServiceSid sets MessagingServiceSid field to given value.

### HasMessagingServiceSid

`func (o *ApiV2010AccountMessage) HasMessagingServiceSid() bool`

HasMessagingServiceSid returns a boolean if a field has been set.

### SetMessagingServiceSidNil

`func (o *ApiV2010AccountMessage) SetMessagingServiceSidNil(b bool)`

 SetMessagingServiceSidNil sets the value for MessagingServiceSid to be an explicit nil

### UnsetMessagingServiceSid
`func (o *ApiV2010AccountMessage) UnsetMessagingServiceSid()`

UnsetMessagingServiceSid ensures that no value is present for MessagingServiceSid, not even an explicit nil
### GetNumMedia

`func (o *ApiV2010AccountMessage) GetNumMedia() string`

GetNumMedia returns the NumMedia field if non-nil, zero value otherwise.

### GetNumMediaOk

`func (o *ApiV2010AccountMessage) GetNumMediaOk() (*string, bool)`

GetNumMediaOk returns a tuple with the NumMedia field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumMedia

`func (o *ApiV2010AccountMessage) SetNumMedia(v string)`

SetNumMedia sets NumMedia field to given value.

### HasNumMedia

`func (o *ApiV2010AccountMessage) HasNumMedia() bool`

HasNumMedia returns a boolean if a field has been set.

### SetNumMediaNil

`func (o *ApiV2010AccountMessage) SetNumMediaNil(b bool)`

 SetNumMediaNil sets the value for NumMedia to be an explicit nil

### UnsetNumMedia
`func (o *ApiV2010AccountMessage) UnsetNumMedia()`

UnsetNumMedia ensures that no value is present for NumMedia, not even an explicit nil
### GetNumSegments

`func (o *ApiV2010AccountMessage) GetNumSegments() string`

GetNumSegments returns the NumSegments field if non-nil, zero value otherwise.

### GetNumSegmentsOk

`func (o *ApiV2010AccountMessage) GetNumSegmentsOk() (*string, bool)`

GetNumSegmentsOk returns a tuple with the NumSegments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumSegments

`func (o *ApiV2010AccountMessage) SetNumSegments(v string)`

SetNumSegments sets NumSegments field to given value.

### HasNumSegments

`func (o *ApiV2010AccountMessage) HasNumSegments() bool`

HasNumSegments returns a boolean if a field has been set.

### SetNumSegmentsNil

`func (o *ApiV2010AccountMessage) SetNumSegmentsNil(b bool)`

 SetNumSegmentsNil sets the value for NumSegments to be an explicit nil

### UnsetNumSegments
`func (o *ApiV2010AccountMessage) UnsetNumSegments()`

UnsetNumSegments ensures that no value is present for NumSegments, not even an explicit nil
### GetPrice

`func (o *ApiV2010AccountMessage) GetPrice() string`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *ApiV2010AccountMessage) GetPriceOk() (*string, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *ApiV2010AccountMessage) SetPrice(v string)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *ApiV2010AccountMessage) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### SetPriceNil

`func (o *ApiV2010AccountMessage) SetPriceNil(b bool)`

 SetPriceNil sets the value for Price to be an explicit nil

### UnsetPrice
`func (o *ApiV2010AccountMessage) UnsetPrice()`

UnsetPrice ensures that no value is present for Price, not even an explicit nil
### GetPriceUnit

`func (o *ApiV2010AccountMessage) GetPriceUnit() string`

GetPriceUnit returns the PriceUnit field if non-nil, zero value otherwise.

### GetPriceUnitOk

`func (o *ApiV2010AccountMessage) GetPriceUnitOk() (*string, bool)`

GetPriceUnitOk returns a tuple with the PriceUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceUnit

`func (o *ApiV2010AccountMessage) SetPriceUnit(v string)`

SetPriceUnit sets PriceUnit field to given value.

### HasPriceUnit

`func (o *ApiV2010AccountMessage) HasPriceUnit() bool`

HasPriceUnit returns a boolean if a field has been set.

### SetPriceUnitNil

`func (o *ApiV2010AccountMessage) SetPriceUnitNil(b bool)`

 SetPriceUnitNil sets the value for PriceUnit to be an explicit nil

### UnsetPriceUnit
`func (o *ApiV2010AccountMessage) UnsetPriceUnit()`

UnsetPriceUnit ensures that no value is present for PriceUnit, not even an explicit nil
### GetSid

`func (o *ApiV2010AccountMessage) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *ApiV2010AccountMessage) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *ApiV2010AccountMessage) SetSid(v string)`

SetSid sets Sid field to given value.

### HasSid

`func (o *ApiV2010AccountMessage) HasSid() bool`

HasSid returns a boolean if a field has been set.

### SetSidNil

`func (o *ApiV2010AccountMessage) SetSidNil(b bool)`

 SetSidNil sets the value for Sid to be an explicit nil

### UnsetSid
`func (o *ApiV2010AccountMessage) UnsetSid()`

UnsetSid ensures that no value is present for Sid, not even an explicit nil
### GetStatus

`func (o *ApiV2010AccountMessage) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ApiV2010AccountMessage) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ApiV2010AccountMessage) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ApiV2010AccountMessage) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *ApiV2010AccountMessage) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *ApiV2010AccountMessage) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetSubresourceUris

`func (o *ApiV2010AccountMessage) GetSubresourceUris() map[string]interface{}`

GetSubresourceUris returns the SubresourceUris field if non-nil, zero value otherwise.

### GetSubresourceUrisOk

`func (o *ApiV2010AccountMessage) GetSubresourceUrisOk() (*map[string]interface{}, bool)`

GetSubresourceUrisOk returns a tuple with the SubresourceUris field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubresourceUris

`func (o *ApiV2010AccountMessage) SetSubresourceUris(v map[string]interface{})`

SetSubresourceUris sets SubresourceUris field to given value.

### HasSubresourceUris

`func (o *ApiV2010AccountMessage) HasSubresourceUris() bool`

HasSubresourceUris returns a boolean if a field has been set.

### SetSubresourceUrisNil

`func (o *ApiV2010AccountMessage) SetSubresourceUrisNil(b bool)`

 SetSubresourceUrisNil sets the value for SubresourceUris to be an explicit nil

### UnsetSubresourceUris
`func (o *ApiV2010AccountMessage) UnsetSubresourceUris()`

UnsetSubresourceUris ensures that no value is present for SubresourceUris, not even an explicit nil
### GetTo

`func (o *ApiV2010AccountMessage) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *ApiV2010AccountMessage) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *ApiV2010AccountMessage) SetTo(v string)`

SetTo sets To field to given value.

### HasTo

`func (o *ApiV2010AccountMessage) HasTo() bool`

HasTo returns a boolean if a field has been set.

### SetToNil

`func (o *ApiV2010AccountMessage) SetToNil(b bool)`

 SetToNil sets the value for To to be an explicit nil

### UnsetTo
`func (o *ApiV2010AccountMessage) UnsetTo()`

UnsetTo ensures that no value is present for To, not even an explicit nil
### GetUri

`func (o *ApiV2010AccountMessage) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *ApiV2010AccountMessage) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *ApiV2010AccountMessage) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *ApiV2010AccountMessage) HasUri() bool`

HasUri returns a boolean if a field has been set.

### SetUriNil

`func (o *ApiV2010AccountMessage) SetUriNil(b bool)`

 SetUriNil sets the value for Uri to be an explicit nil

### UnsetUri
`func (o *ApiV2010AccountMessage) UnsetUri()`

UnsetUri ensures that no value is present for Uri, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


