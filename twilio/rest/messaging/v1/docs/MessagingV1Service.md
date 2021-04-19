# MessagingV1Service

## Properties

Name | Type | Description
------------ | ------------- | -------------
**AccountSid** | Pointer to **NullableString** | The SID of the Account that created the resource | [optional] 
**AreaCodeGeomatch** | Pointer to **NullableBool** | Whether to enable Area Code Geomatch on the Service Instance | [optional] 
**DateCreated** | Pointer to **NullableTime** | The ISO 8601 date and time in GMT when the resource was created | [optional] 
**DateUpdated** | Pointer to **NullableTime** | The ISO 8601 date and time in GMT when the resource was last updated | [optional] 
**FallbackMethod** | Pointer to **NullableString** | The HTTP method we use to call fallback_url | [optional] 
**FallbackToLongCode** | Pointer to **NullableBool** | Whether to enable Fallback to Long Code for messages sent through the Service instance | [optional] 
**FallbackUrl** | Pointer to **NullableString** | The URL that we call using fallback_method if an error occurs while retrieving or executing the TwiML from the Inbound Request URL | [optional] 
**FriendlyName** | Pointer to **NullableString** | The string that you assigned to describe the resource | [optional] 
**InboundMethod** | Pointer to **NullableString** | The HTTP method we use to call inbound_request_url | [optional] 
**InboundRequestUrl** | Pointer to **NullableString** | The URL we call using inbound_method when a message is received by any phone number or short code in the Service | [optional] 
**Links** | Pointer to **map[string]interface{}** | The absolute URLs of related resources | [optional] 
**MmsConverter** | Pointer to **NullableBool** | Whether to enable the MMS Converter for messages sent through the Service instance | [optional] 
**ScanMessageContent** | Pointer to **NullableString** | Reserved | [optional] 
**Sid** | Pointer to **NullableString** | The unique string that identifies the resource | [optional] 
**SmartEncoding** | Pointer to **NullableBool** | Whether to enable Encoding for messages sent through the Service instance | [optional] 
**StatusCallback** | Pointer to **NullableString** | The URL we call to pass status updates about message delivery | [optional] 
**StickySender** | Pointer to **NullableBool** | Whether to enable Sticky Sender on the Service instance | [optional] 
**SynchronousValidation** | Pointer to **NullableBool** | Reserved | [optional] 
**Url** | Pointer to **NullableString** | The absolute URL of the Service resource | [optional] 
**ValidityPeriod** | Pointer to **NullableInt32** | How long, in seconds, messages sent from the Service are valid | [optional] 

## Methods

### NewMessagingV1Service

`func NewMessagingV1Service() *MessagingV1Service`

NewMessagingV1Service instantiates a new MessagingV1Service object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMessagingV1ServiceWithDefaults

`func NewMessagingV1ServiceWithDefaults() *MessagingV1Service`

NewMessagingV1ServiceWithDefaults instantiates a new MessagingV1Service object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSid

`func (o *MessagingV1Service) GetAccountSid() string`

GetAccountSid returns the AccountSid field if non-nil, zero value otherwise.

### GetAccountSidOk

`func (o *MessagingV1Service) GetAccountSidOk() (*string, bool)`

GetAccountSidOk returns a tuple with the AccountSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSid

`func (o *MessagingV1Service) SetAccountSid(v string)`

SetAccountSid sets AccountSid field to given value.

### HasAccountSid

`func (o *MessagingV1Service) HasAccountSid() bool`

HasAccountSid returns a boolean if a field has been set.

### SetAccountSidNil

`func (o *MessagingV1Service) SetAccountSidNil(b bool)`

 SetAccountSidNil sets the value for AccountSid to be an explicit nil

### UnsetAccountSid
`func (o *MessagingV1Service) UnsetAccountSid()`

UnsetAccountSid ensures that no value is present for AccountSid, not even an explicit nil
### GetAreaCodeGeomatch

`func (o *MessagingV1Service) GetAreaCodeGeomatch() bool`

GetAreaCodeGeomatch returns the AreaCodeGeomatch field if non-nil, zero value otherwise.

### GetAreaCodeGeomatchOk

`func (o *MessagingV1Service) GetAreaCodeGeomatchOk() (*bool, bool)`

GetAreaCodeGeomatchOk returns a tuple with the AreaCodeGeomatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAreaCodeGeomatch

`func (o *MessagingV1Service) SetAreaCodeGeomatch(v bool)`

SetAreaCodeGeomatch sets AreaCodeGeomatch field to given value.

### HasAreaCodeGeomatch

`func (o *MessagingV1Service) HasAreaCodeGeomatch() bool`

HasAreaCodeGeomatch returns a boolean if a field has been set.

### SetAreaCodeGeomatchNil

`func (o *MessagingV1Service) SetAreaCodeGeomatchNil(b bool)`

 SetAreaCodeGeomatchNil sets the value for AreaCodeGeomatch to be an explicit nil

### UnsetAreaCodeGeomatch
`func (o *MessagingV1Service) UnsetAreaCodeGeomatch()`

UnsetAreaCodeGeomatch ensures that no value is present for AreaCodeGeomatch, not even an explicit nil
### GetDateCreated

`func (o *MessagingV1Service) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *MessagingV1Service) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *MessagingV1Service) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *MessagingV1Service) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### SetDateCreatedNil

`func (o *MessagingV1Service) SetDateCreatedNil(b bool)`

 SetDateCreatedNil sets the value for DateCreated to be an explicit nil

### UnsetDateCreated
`func (o *MessagingV1Service) UnsetDateCreated()`

UnsetDateCreated ensures that no value is present for DateCreated, not even an explicit nil
### GetDateUpdated

`func (o *MessagingV1Service) GetDateUpdated() time.Time`

GetDateUpdated returns the DateUpdated field if non-nil, zero value otherwise.

### GetDateUpdatedOk

`func (o *MessagingV1Service) GetDateUpdatedOk() (*time.Time, bool)`

GetDateUpdatedOk returns a tuple with the DateUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateUpdated

`func (o *MessagingV1Service) SetDateUpdated(v time.Time)`

SetDateUpdated sets DateUpdated field to given value.

### HasDateUpdated

`func (o *MessagingV1Service) HasDateUpdated() bool`

HasDateUpdated returns a boolean if a field has been set.

### SetDateUpdatedNil

`func (o *MessagingV1Service) SetDateUpdatedNil(b bool)`

 SetDateUpdatedNil sets the value for DateUpdated to be an explicit nil

### UnsetDateUpdated
`func (o *MessagingV1Service) UnsetDateUpdated()`

UnsetDateUpdated ensures that no value is present for DateUpdated, not even an explicit nil
### GetFallbackMethod

`func (o *MessagingV1Service) GetFallbackMethod() string`

GetFallbackMethod returns the FallbackMethod field if non-nil, zero value otherwise.

### GetFallbackMethodOk

`func (o *MessagingV1Service) GetFallbackMethodOk() (*string, bool)`

GetFallbackMethodOk returns a tuple with the FallbackMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFallbackMethod

`func (o *MessagingV1Service) SetFallbackMethod(v string)`

SetFallbackMethod sets FallbackMethod field to given value.

### HasFallbackMethod

`func (o *MessagingV1Service) HasFallbackMethod() bool`

HasFallbackMethod returns a boolean if a field has been set.

### SetFallbackMethodNil

`func (o *MessagingV1Service) SetFallbackMethodNil(b bool)`

 SetFallbackMethodNil sets the value for FallbackMethod to be an explicit nil

### UnsetFallbackMethod
`func (o *MessagingV1Service) UnsetFallbackMethod()`

UnsetFallbackMethod ensures that no value is present for FallbackMethod, not even an explicit nil
### GetFallbackToLongCode

`func (o *MessagingV1Service) GetFallbackToLongCode() bool`

GetFallbackToLongCode returns the FallbackToLongCode field if non-nil, zero value otherwise.

### GetFallbackToLongCodeOk

`func (o *MessagingV1Service) GetFallbackToLongCodeOk() (*bool, bool)`

GetFallbackToLongCodeOk returns a tuple with the FallbackToLongCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFallbackToLongCode

`func (o *MessagingV1Service) SetFallbackToLongCode(v bool)`

SetFallbackToLongCode sets FallbackToLongCode field to given value.

### HasFallbackToLongCode

`func (o *MessagingV1Service) HasFallbackToLongCode() bool`

HasFallbackToLongCode returns a boolean if a field has been set.

### SetFallbackToLongCodeNil

`func (o *MessagingV1Service) SetFallbackToLongCodeNil(b bool)`

 SetFallbackToLongCodeNil sets the value for FallbackToLongCode to be an explicit nil

### UnsetFallbackToLongCode
`func (o *MessagingV1Service) UnsetFallbackToLongCode()`

UnsetFallbackToLongCode ensures that no value is present for FallbackToLongCode, not even an explicit nil
### GetFallbackUrl

`func (o *MessagingV1Service) GetFallbackUrl() string`

GetFallbackUrl returns the FallbackUrl field if non-nil, zero value otherwise.

### GetFallbackUrlOk

`func (o *MessagingV1Service) GetFallbackUrlOk() (*string, bool)`

GetFallbackUrlOk returns a tuple with the FallbackUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFallbackUrl

`func (o *MessagingV1Service) SetFallbackUrl(v string)`

SetFallbackUrl sets FallbackUrl field to given value.

### HasFallbackUrl

`func (o *MessagingV1Service) HasFallbackUrl() bool`

HasFallbackUrl returns a boolean if a field has been set.

### SetFallbackUrlNil

`func (o *MessagingV1Service) SetFallbackUrlNil(b bool)`

 SetFallbackUrlNil sets the value for FallbackUrl to be an explicit nil

### UnsetFallbackUrl
`func (o *MessagingV1Service) UnsetFallbackUrl()`

UnsetFallbackUrl ensures that no value is present for FallbackUrl, not even an explicit nil
### GetFriendlyName

`func (o *MessagingV1Service) GetFriendlyName() string`

GetFriendlyName returns the FriendlyName field if non-nil, zero value otherwise.

### GetFriendlyNameOk

`func (o *MessagingV1Service) GetFriendlyNameOk() (*string, bool)`

GetFriendlyNameOk returns a tuple with the FriendlyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFriendlyName

`func (o *MessagingV1Service) SetFriendlyName(v string)`

SetFriendlyName sets FriendlyName field to given value.

### HasFriendlyName

`func (o *MessagingV1Service) HasFriendlyName() bool`

HasFriendlyName returns a boolean if a field has been set.

### SetFriendlyNameNil

`func (o *MessagingV1Service) SetFriendlyNameNil(b bool)`

 SetFriendlyNameNil sets the value for FriendlyName to be an explicit nil

### UnsetFriendlyName
`func (o *MessagingV1Service) UnsetFriendlyName()`

UnsetFriendlyName ensures that no value is present for FriendlyName, not even an explicit nil
### GetInboundMethod

`func (o *MessagingV1Service) GetInboundMethod() string`

GetInboundMethod returns the InboundMethod field if non-nil, zero value otherwise.

### GetInboundMethodOk

`func (o *MessagingV1Service) GetInboundMethodOk() (*string, bool)`

GetInboundMethodOk returns a tuple with the InboundMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInboundMethod

`func (o *MessagingV1Service) SetInboundMethod(v string)`

SetInboundMethod sets InboundMethod field to given value.

### HasInboundMethod

`func (o *MessagingV1Service) HasInboundMethod() bool`

HasInboundMethod returns a boolean if a field has been set.

### SetInboundMethodNil

`func (o *MessagingV1Service) SetInboundMethodNil(b bool)`

 SetInboundMethodNil sets the value for InboundMethod to be an explicit nil

### UnsetInboundMethod
`func (o *MessagingV1Service) UnsetInboundMethod()`

UnsetInboundMethod ensures that no value is present for InboundMethod, not even an explicit nil
### GetInboundRequestUrl

`func (o *MessagingV1Service) GetInboundRequestUrl() string`

GetInboundRequestUrl returns the InboundRequestUrl field if non-nil, zero value otherwise.

### GetInboundRequestUrlOk

`func (o *MessagingV1Service) GetInboundRequestUrlOk() (*string, bool)`

GetInboundRequestUrlOk returns a tuple with the InboundRequestUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInboundRequestUrl

`func (o *MessagingV1Service) SetInboundRequestUrl(v string)`

SetInboundRequestUrl sets InboundRequestUrl field to given value.

### HasInboundRequestUrl

`func (o *MessagingV1Service) HasInboundRequestUrl() bool`

HasInboundRequestUrl returns a boolean if a field has been set.

### SetInboundRequestUrlNil

`func (o *MessagingV1Service) SetInboundRequestUrlNil(b bool)`

 SetInboundRequestUrlNil sets the value for InboundRequestUrl to be an explicit nil

### UnsetInboundRequestUrl
`func (o *MessagingV1Service) UnsetInboundRequestUrl()`

UnsetInboundRequestUrl ensures that no value is present for InboundRequestUrl, not even an explicit nil
### GetLinks

`func (o *MessagingV1Service) GetLinks() map[string]interface{}`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *MessagingV1Service) GetLinksOk() (*map[string]interface{}, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *MessagingV1Service) SetLinks(v map[string]interface{})`

SetLinks sets Links field to given value.

### HasLinks

`func (o *MessagingV1Service) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### SetLinksNil

`func (o *MessagingV1Service) SetLinksNil(b bool)`

 SetLinksNil sets the value for Links to be an explicit nil

### UnsetLinks
`func (o *MessagingV1Service) UnsetLinks()`

UnsetLinks ensures that no value is present for Links, not even an explicit nil
### GetMmsConverter

`func (o *MessagingV1Service) GetMmsConverter() bool`

GetMmsConverter returns the MmsConverter field if non-nil, zero value otherwise.

### GetMmsConverterOk

`func (o *MessagingV1Service) GetMmsConverterOk() (*bool, bool)`

GetMmsConverterOk returns a tuple with the MmsConverter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMmsConverter

`func (o *MessagingV1Service) SetMmsConverter(v bool)`

SetMmsConverter sets MmsConverter field to given value.

### HasMmsConverter

`func (o *MessagingV1Service) HasMmsConverter() bool`

HasMmsConverter returns a boolean if a field has been set.

### SetMmsConverterNil

`func (o *MessagingV1Service) SetMmsConverterNil(b bool)`

 SetMmsConverterNil sets the value for MmsConverter to be an explicit nil

### UnsetMmsConverter
`func (o *MessagingV1Service) UnsetMmsConverter()`

UnsetMmsConverter ensures that no value is present for MmsConverter, not even an explicit nil
### GetScanMessageContent

`func (o *MessagingV1Service) GetScanMessageContent() string`

GetScanMessageContent returns the ScanMessageContent field if non-nil, zero value otherwise.

### GetScanMessageContentOk

`func (o *MessagingV1Service) GetScanMessageContentOk() (*string, bool)`

GetScanMessageContentOk returns a tuple with the ScanMessageContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScanMessageContent

`func (o *MessagingV1Service) SetScanMessageContent(v string)`

SetScanMessageContent sets ScanMessageContent field to given value.

### HasScanMessageContent

`func (o *MessagingV1Service) HasScanMessageContent() bool`

HasScanMessageContent returns a boolean if a field has been set.

### SetScanMessageContentNil

`func (o *MessagingV1Service) SetScanMessageContentNil(b bool)`

 SetScanMessageContentNil sets the value for ScanMessageContent to be an explicit nil

### UnsetScanMessageContent
`func (o *MessagingV1Service) UnsetScanMessageContent()`

UnsetScanMessageContent ensures that no value is present for ScanMessageContent, not even an explicit nil
### GetSid

`func (o *MessagingV1Service) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *MessagingV1Service) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *MessagingV1Service) SetSid(v string)`

SetSid sets Sid field to given value.

### HasSid

`func (o *MessagingV1Service) HasSid() bool`

HasSid returns a boolean if a field has been set.

### SetSidNil

`func (o *MessagingV1Service) SetSidNil(b bool)`

 SetSidNil sets the value for Sid to be an explicit nil

### UnsetSid
`func (o *MessagingV1Service) UnsetSid()`

UnsetSid ensures that no value is present for Sid, not even an explicit nil
### GetSmartEncoding

`func (o *MessagingV1Service) GetSmartEncoding() bool`

GetSmartEncoding returns the SmartEncoding field if non-nil, zero value otherwise.

### GetSmartEncodingOk

`func (o *MessagingV1Service) GetSmartEncodingOk() (*bool, bool)`

GetSmartEncodingOk returns a tuple with the SmartEncoding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmartEncoding

`func (o *MessagingV1Service) SetSmartEncoding(v bool)`

SetSmartEncoding sets SmartEncoding field to given value.

### HasSmartEncoding

`func (o *MessagingV1Service) HasSmartEncoding() bool`

HasSmartEncoding returns a boolean if a field has been set.

### SetSmartEncodingNil

`func (o *MessagingV1Service) SetSmartEncodingNil(b bool)`

 SetSmartEncodingNil sets the value for SmartEncoding to be an explicit nil

### UnsetSmartEncoding
`func (o *MessagingV1Service) UnsetSmartEncoding()`

UnsetSmartEncoding ensures that no value is present for SmartEncoding, not even an explicit nil
### GetStatusCallback

`func (o *MessagingV1Service) GetStatusCallback() string`

GetStatusCallback returns the StatusCallback field if non-nil, zero value otherwise.

### GetStatusCallbackOk

`func (o *MessagingV1Service) GetStatusCallbackOk() (*string, bool)`

GetStatusCallbackOk returns a tuple with the StatusCallback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCallback

`func (o *MessagingV1Service) SetStatusCallback(v string)`

SetStatusCallback sets StatusCallback field to given value.

### HasStatusCallback

`func (o *MessagingV1Service) HasStatusCallback() bool`

HasStatusCallback returns a boolean if a field has been set.

### SetStatusCallbackNil

`func (o *MessagingV1Service) SetStatusCallbackNil(b bool)`

 SetStatusCallbackNil sets the value for StatusCallback to be an explicit nil

### UnsetStatusCallback
`func (o *MessagingV1Service) UnsetStatusCallback()`

UnsetStatusCallback ensures that no value is present for StatusCallback, not even an explicit nil
### GetStickySender

`func (o *MessagingV1Service) GetStickySender() bool`

GetStickySender returns the StickySender field if non-nil, zero value otherwise.

### GetStickySenderOk

`func (o *MessagingV1Service) GetStickySenderOk() (*bool, bool)`

GetStickySenderOk returns a tuple with the StickySender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStickySender

`func (o *MessagingV1Service) SetStickySender(v bool)`

SetStickySender sets StickySender field to given value.

### HasStickySender

`func (o *MessagingV1Service) HasStickySender() bool`

HasStickySender returns a boolean if a field has been set.

### SetStickySenderNil

`func (o *MessagingV1Service) SetStickySenderNil(b bool)`

 SetStickySenderNil sets the value for StickySender to be an explicit nil

### UnsetStickySender
`func (o *MessagingV1Service) UnsetStickySender()`

UnsetStickySender ensures that no value is present for StickySender, not even an explicit nil
### GetSynchronousValidation

`func (o *MessagingV1Service) GetSynchronousValidation() bool`

GetSynchronousValidation returns the SynchronousValidation field if non-nil, zero value otherwise.

### GetSynchronousValidationOk

`func (o *MessagingV1Service) GetSynchronousValidationOk() (*bool, bool)`

GetSynchronousValidationOk returns a tuple with the SynchronousValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSynchronousValidation

`func (o *MessagingV1Service) SetSynchronousValidation(v bool)`

SetSynchronousValidation sets SynchronousValidation field to given value.

### HasSynchronousValidation

`func (o *MessagingV1Service) HasSynchronousValidation() bool`

HasSynchronousValidation returns a boolean if a field has been set.

### SetSynchronousValidationNil

`func (o *MessagingV1Service) SetSynchronousValidationNil(b bool)`

 SetSynchronousValidationNil sets the value for SynchronousValidation to be an explicit nil

### UnsetSynchronousValidation
`func (o *MessagingV1Service) UnsetSynchronousValidation()`

UnsetSynchronousValidation ensures that no value is present for SynchronousValidation, not even an explicit nil
### GetUrl

`func (o *MessagingV1Service) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *MessagingV1Service) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *MessagingV1Service) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *MessagingV1Service) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *MessagingV1Service) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *MessagingV1Service) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetValidityPeriod

`func (o *MessagingV1Service) GetValidityPeriod() int32`

GetValidityPeriod returns the ValidityPeriod field if non-nil, zero value otherwise.

### GetValidityPeriodOk

`func (o *MessagingV1Service) GetValidityPeriodOk() (*int32, bool)`

GetValidityPeriodOk returns a tuple with the ValidityPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidityPeriod

`func (o *MessagingV1Service) SetValidityPeriod(v int32)`

SetValidityPeriod sets ValidityPeriod field to given value.

### HasValidityPeriod

`func (o *MessagingV1Service) HasValidityPeriod() bool`

HasValidityPeriod returns a boolean if a field has been set.

### SetValidityPeriodNil

`func (o *MessagingV1Service) SetValidityPeriodNil(b bool)`

 SetValidityPeriodNil sets the value for ValidityPeriod to be an explicit nil

### UnsetValidityPeriod
`func (o *MessagingV1Service) UnsetValidityPeriod()`

UnsetValidityPeriod ensures that no value is present for ValidityPeriod, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


