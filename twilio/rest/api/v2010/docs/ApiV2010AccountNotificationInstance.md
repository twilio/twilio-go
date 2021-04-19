# ApiV2010AccountNotificationInstance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **NullableString** | The SID of the Account that created the resource | [optional] 
**ApiVersion** | Pointer to **NullableString** | The API version used to generate the notification | [optional] 
**CallSid** | Pointer to **NullableString** | The SID of the Call the resource is associated with | [optional] 
**DateCreated** | Pointer to **NullableString** | The RFC 2822 date and time in GMT that the resource was created | [optional] 
**DateUpdated** | Pointer to **NullableString** | The RFC 2822 date and time in GMT that the resource was last updated | [optional] 
**ErrorCode** | Pointer to **NullableString** | A unique error code corresponding to the notification | [optional] 
**Log** | Pointer to **NullableString** | An integer log level | [optional] 
**MessageDate** | Pointer to **NullableString** | The date the notification was generated | [optional] 
**MessageText** | Pointer to **NullableString** | The text of the notification | [optional] 
**MoreInfo** | Pointer to **NullableString** | A URL for more information about the error code | [optional] 
**RequestMethod** | Pointer to **NullableString** | HTTP method used with the request url | [optional] 
**RequestUrl** | Pointer to **NullableString** | URL of the resource that generated the notification | [optional] 
**RequestVariables** | Pointer to **NullableString** | Twilio-generated HTTP variables sent to the server | [optional] 
**ResponseBody** | Pointer to **NullableString** | The HTTP body returned by your server | [optional] 
**ResponseHeaders** | Pointer to **NullableString** | The HTTP headers returned by your server | [optional] 
**Sid** | Pointer to **NullableString** | The unique string that identifies the resource | [optional] 
**Uri** | Pointer to **NullableString** | The URI of the resource, relative to &#x60;https://api.twilio.com&#x60; | [optional] 

## Methods

### NewApiV2010AccountNotificationInstance

`func NewApiV2010AccountNotificationInstance() *ApiV2010AccountNotificationInstance`

NewApiV2010AccountNotificationInstance instantiates a new ApiV2010AccountNotificationInstance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiV2010AccountNotificationInstanceWithDefaults

`func NewApiV2010AccountNotificationInstanceWithDefaults() *ApiV2010AccountNotificationInstance`

NewApiV2010AccountNotificationInstanceWithDefaults instantiates a new ApiV2010AccountNotificationInstance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSid

`func (o *ApiV2010AccountNotificationInstance) GetAccountSid() string`

GetAccountSid returns the AccountSid field if non-nil, zero value otherwise.

### GetAccountSidOk

`func (o *ApiV2010AccountNotificationInstance) GetAccountSidOk() (*string, bool)`

GetAccountSidOk returns a tuple with the AccountSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSid

`func (o *ApiV2010AccountNotificationInstance) SetAccountSid(v string)`

SetAccountSid sets AccountSid field to given value.

### HasAccountSid

`func (o *ApiV2010AccountNotificationInstance) HasAccountSid() bool`

HasAccountSid returns a boolean if a field has been set.

### SetAccountSidNil

`func (o *ApiV2010AccountNotificationInstance) SetAccountSidNil(b bool)`

 SetAccountSidNil sets the value for AccountSid to be an explicit nil

### UnsetAccountSid
`func (o *ApiV2010AccountNotificationInstance) UnsetAccountSid()`

UnsetAccountSid ensures that no value is present for AccountSid, not even an explicit nil
### GetApiVersion

`func (o *ApiV2010AccountNotificationInstance) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *ApiV2010AccountNotificationInstance) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *ApiV2010AccountNotificationInstance) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *ApiV2010AccountNotificationInstance) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### SetApiVersionNil

`func (o *ApiV2010AccountNotificationInstance) SetApiVersionNil(b bool)`

 SetApiVersionNil sets the value for ApiVersion to be an explicit nil

### UnsetApiVersion
`func (o *ApiV2010AccountNotificationInstance) UnsetApiVersion()`

UnsetApiVersion ensures that no value is present for ApiVersion, not even an explicit nil
### GetCallSid

`func (o *ApiV2010AccountNotificationInstance) GetCallSid() string`

GetCallSid returns the CallSid field if non-nil, zero value otherwise.

### GetCallSidOk

`func (o *ApiV2010AccountNotificationInstance) GetCallSidOk() (*string, bool)`

GetCallSidOk returns a tuple with the CallSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallSid

`func (o *ApiV2010AccountNotificationInstance) SetCallSid(v string)`

SetCallSid sets CallSid field to given value.

### HasCallSid

`func (o *ApiV2010AccountNotificationInstance) HasCallSid() bool`

HasCallSid returns a boolean if a field has been set.

### SetCallSidNil

`func (o *ApiV2010AccountNotificationInstance) SetCallSidNil(b bool)`

 SetCallSidNil sets the value for CallSid to be an explicit nil

### UnsetCallSid
`func (o *ApiV2010AccountNotificationInstance) UnsetCallSid()`

UnsetCallSid ensures that no value is present for CallSid, not even an explicit nil
### GetDateCreated

`func (o *ApiV2010AccountNotificationInstance) GetDateCreated() string`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *ApiV2010AccountNotificationInstance) GetDateCreatedOk() (*string, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *ApiV2010AccountNotificationInstance) SetDateCreated(v string)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *ApiV2010AccountNotificationInstance) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### SetDateCreatedNil

`func (o *ApiV2010AccountNotificationInstance) SetDateCreatedNil(b bool)`

 SetDateCreatedNil sets the value for DateCreated to be an explicit nil

### UnsetDateCreated
`func (o *ApiV2010AccountNotificationInstance) UnsetDateCreated()`

UnsetDateCreated ensures that no value is present for DateCreated, not even an explicit nil
### GetDateUpdated

`func (o *ApiV2010AccountNotificationInstance) GetDateUpdated() string`

GetDateUpdated returns the DateUpdated field if non-nil, zero value otherwise.

### GetDateUpdatedOk

`func (o *ApiV2010AccountNotificationInstance) GetDateUpdatedOk() (*string, bool)`

GetDateUpdatedOk returns a tuple with the DateUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateUpdated

`func (o *ApiV2010AccountNotificationInstance) SetDateUpdated(v string)`

SetDateUpdated sets DateUpdated field to given value.

### HasDateUpdated

`func (o *ApiV2010AccountNotificationInstance) HasDateUpdated() bool`

HasDateUpdated returns a boolean if a field has been set.

### SetDateUpdatedNil

`func (o *ApiV2010AccountNotificationInstance) SetDateUpdatedNil(b bool)`

 SetDateUpdatedNil sets the value for DateUpdated to be an explicit nil

### UnsetDateUpdated
`func (o *ApiV2010AccountNotificationInstance) UnsetDateUpdated()`

UnsetDateUpdated ensures that no value is present for DateUpdated, not even an explicit nil
### GetErrorCode

`func (o *ApiV2010AccountNotificationInstance) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *ApiV2010AccountNotificationInstance) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *ApiV2010AccountNotificationInstance) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *ApiV2010AccountNotificationInstance) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### SetErrorCodeNil

`func (o *ApiV2010AccountNotificationInstance) SetErrorCodeNil(b bool)`

 SetErrorCodeNil sets the value for ErrorCode to be an explicit nil

### UnsetErrorCode
`func (o *ApiV2010AccountNotificationInstance) UnsetErrorCode()`

UnsetErrorCode ensures that no value is present for ErrorCode, not even an explicit nil
### GetLog

`func (o *ApiV2010AccountNotificationInstance) GetLog() string`

GetLog returns the Log field if non-nil, zero value otherwise.

### GetLogOk

`func (o *ApiV2010AccountNotificationInstance) GetLogOk() (*string, bool)`

GetLogOk returns a tuple with the Log field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLog

`func (o *ApiV2010AccountNotificationInstance) SetLog(v string)`

SetLog sets Log field to given value.

### HasLog

`func (o *ApiV2010AccountNotificationInstance) HasLog() bool`

HasLog returns a boolean if a field has been set.

### SetLogNil

`func (o *ApiV2010AccountNotificationInstance) SetLogNil(b bool)`

 SetLogNil sets the value for Log to be an explicit nil

### UnsetLog
`func (o *ApiV2010AccountNotificationInstance) UnsetLog()`

UnsetLog ensures that no value is present for Log, not even an explicit nil
### GetMessageDate

`func (o *ApiV2010AccountNotificationInstance) GetMessageDate() string`

GetMessageDate returns the MessageDate field if non-nil, zero value otherwise.

### GetMessageDateOk

`func (o *ApiV2010AccountNotificationInstance) GetMessageDateOk() (*string, bool)`

GetMessageDateOk returns a tuple with the MessageDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageDate

`func (o *ApiV2010AccountNotificationInstance) SetMessageDate(v string)`

SetMessageDate sets MessageDate field to given value.

### HasMessageDate

`func (o *ApiV2010AccountNotificationInstance) HasMessageDate() bool`

HasMessageDate returns a boolean if a field has been set.

### SetMessageDateNil

`func (o *ApiV2010AccountNotificationInstance) SetMessageDateNil(b bool)`

 SetMessageDateNil sets the value for MessageDate to be an explicit nil

### UnsetMessageDate
`func (o *ApiV2010AccountNotificationInstance) UnsetMessageDate()`

UnsetMessageDate ensures that no value is present for MessageDate, not even an explicit nil
### GetMessageText

`func (o *ApiV2010AccountNotificationInstance) GetMessageText() string`

GetMessageText returns the MessageText field if non-nil, zero value otherwise.

### GetMessageTextOk

`func (o *ApiV2010AccountNotificationInstance) GetMessageTextOk() (*string, bool)`

GetMessageTextOk returns a tuple with the MessageText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageText

`func (o *ApiV2010AccountNotificationInstance) SetMessageText(v string)`

SetMessageText sets MessageText field to given value.

### HasMessageText

`func (o *ApiV2010AccountNotificationInstance) HasMessageText() bool`

HasMessageText returns a boolean if a field has been set.

### SetMessageTextNil

`func (o *ApiV2010AccountNotificationInstance) SetMessageTextNil(b bool)`

 SetMessageTextNil sets the value for MessageText to be an explicit nil

### UnsetMessageText
`func (o *ApiV2010AccountNotificationInstance) UnsetMessageText()`

UnsetMessageText ensures that no value is present for MessageText, not even an explicit nil
### GetMoreInfo

`func (o *ApiV2010AccountNotificationInstance) GetMoreInfo() string`

GetMoreInfo returns the MoreInfo field if non-nil, zero value otherwise.

### GetMoreInfoOk

`func (o *ApiV2010AccountNotificationInstance) GetMoreInfoOk() (*string, bool)`

GetMoreInfoOk returns a tuple with the MoreInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoreInfo

`func (o *ApiV2010AccountNotificationInstance) SetMoreInfo(v string)`

SetMoreInfo sets MoreInfo field to given value.

### HasMoreInfo

`func (o *ApiV2010AccountNotificationInstance) HasMoreInfo() bool`

HasMoreInfo returns a boolean if a field has been set.

### SetMoreInfoNil

`func (o *ApiV2010AccountNotificationInstance) SetMoreInfoNil(b bool)`

 SetMoreInfoNil sets the value for MoreInfo to be an explicit nil

### UnsetMoreInfo
`func (o *ApiV2010AccountNotificationInstance) UnsetMoreInfo()`

UnsetMoreInfo ensures that no value is present for MoreInfo, not even an explicit nil
### GetRequestMethod

`func (o *ApiV2010AccountNotificationInstance) GetRequestMethod() string`

GetRequestMethod returns the RequestMethod field if non-nil, zero value otherwise.

### GetRequestMethodOk

`func (o *ApiV2010AccountNotificationInstance) GetRequestMethodOk() (*string, bool)`

GetRequestMethodOk returns a tuple with the RequestMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestMethod

`func (o *ApiV2010AccountNotificationInstance) SetRequestMethod(v string)`

SetRequestMethod sets RequestMethod field to given value.

### HasRequestMethod

`func (o *ApiV2010AccountNotificationInstance) HasRequestMethod() bool`

HasRequestMethod returns a boolean if a field has been set.

### SetRequestMethodNil

`func (o *ApiV2010AccountNotificationInstance) SetRequestMethodNil(b bool)`

 SetRequestMethodNil sets the value for RequestMethod to be an explicit nil

### UnsetRequestMethod
`func (o *ApiV2010AccountNotificationInstance) UnsetRequestMethod()`

UnsetRequestMethod ensures that no value is present for RequestMethod, not even an explicit nil
### GetRequestUrl

`func (o *ApiV2010AccountNotificationInstance) GetRequestUrl() string`

GetRequestUrl returns the RequestUrl field if non-nil, zero value otherwise.

### GetRequestUrlOk

`func (o *ApiV2010AccountNotificationInstance) GetRequestUrlOk() (*string, bool)`

GetRequestUrlOk returns a tuple with the RequestUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestUrl

`func (o *ApiV2010AccountNotificationInstance) SetRequestUrl(v string)`

SetRequestUrl sets RequestUrl field to given value.

### HasRequestUrl

`func (o *ApiV2010AccountNotificationInstance) HasRequestUrl() bool`

HasRequestUrl returns a boolean if a field has been set.

### SetRequestUrlNil

`func (o *ApiV2010AccountNotificationInstance) SetRequestUrlNil(b bool)`

 SetRequestUrlNil sets the value for RequestUrl to be an explicit nil

### UnsetRequestUrl
`func (o *ApiV2010AccountNotificationInstance) UnsetRequestUrl()`

UnsetRequestUrl ensures that no value is present for RequestUrl, not even an explicit nil
### GetRequestVariables

`func (o *ApiV2010AccountNotificationInstance) GetRequestVariables() string`

GetRequestVariables returns the RequestVariables field if non-nil, zero value otherwise.

### GetRequestVariablesOk

`func (o *ApiV2010AccountNotificationInstance) GetRequestVariablesOk() (*string, bool)`

GetRequestVariablesOk returns a tuple with the RequestVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestVariables

`func (o *ApiV2010AccountNotificationInstance) SetRequestVariables(v string)`

SetRequestVariables sets RequestVariables field to given value.

### HasRequestVariables

`func (o *ApiV2010AccountNotificationInstance) HasRequestVariables() bool`

HasRequestVariables returns a boolean if a field has been set.

### SetRequestVariablesNil

`func (o *ApiV2010AccountNotificationInstance) SetRequestVariablesNil(b bool)`

 SetRequestVariablesNil sets the value for RequestVariables to be an explicit nil

### UnsetRequestVariables
`func (o *ApiV2010AccountNotificationInstance) UnsetRequestVariables()`

UnsetRequestVariables ensures that no value is present for RequestVariables, not even an explicit nil
### GetResponseBody

`func (o *ApiV2010AccountNotificationInstance) GetResponseBody() string`

GetResponseBody returns the ResponseBody field if non-nil, zero value otherwise.

### GetResponseBodyOk

`func (o *ApiV2010AccountNotificationInstance) GetResponseBodyOk() (*string, bool)`

GetResponseBodyOk returns a tuple with the ResponseBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseBody

`func (o *ApiV2010AccountNotificationInstance) SetResponseBody(v string)`

SetResponseBody sets ResponseBody field to given value.

### HasResponseBody

`func (o *ApiV2010AccountNotificationInstance) HasResponseBody() bool`

HasResponseBody returns a boolean if a field has been set.

### SetResponseBodyNil

`func (o *ApiV2010AccountNotificationInstance) SetResponseBodyNil(b bool)`

 SetResponseBodyNil sets the value for ResponseBody to be an explicit nil

### UnsetResponseBody
`func (o *ApiV2010AccountNotificationInstance) UnsetResponseBody()`

UnsetResponseBody ensures that no value is present for ResponseBody, not even an explicit nil
### GetResponseHeaders

`func (o *ApiV2010AccountNotificationInstance) GetResponseHeaders() string`

GetResponseHeaders returns the ResponseHeaders field if non-nil, zero value otherwise.

### GetResponseHeadersOk

`func (o *ApiV2010AccountNotificationInstance) GetResponseHeadersOk() (*string, bool)`

GetResponseHeadersOk returns a tuple with the ResponseHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseHeaders

`func (o *ApiV2010AccountNotificationInstance) SetResponseHeaders(v string)`

SetResponseHeaders sets ResponseHeaders field to given value.

### HasResponseHeaders

`func (o *ApiV2010AccountNotificationInstance) HasResponseHeaders() bool`

HasResponseHeaders returns a boolean if a field has been set.

### SetResponseHeadersNil

`func (o *ApiV2010AccountNotificationInstance) SetResponseHeadersNil(b bool)`

 SetResponseHeadersNil sets the value for ResponseHeaders to be an explicit nil

### UnsetResponseHeaders
`func (o *ApiV2010AccountNotificationInstance) UnsetResponseHeaders()`

UnsetResponseHeaders ensures that no value is present for ResponseHeaders, not even an explicit nil
### GetSid

`func (o *ApiV2010AccountNotificationInstance) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *ApiV2010AccountNotificationInstance) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *ApiV2010AccountNotificationInstance) SetSid(v string)`

SetSid sets Sid field to given value.

### HasSid

`func (o *ApiV2010AccountNotificationInstance) HasSid() bool`

HasSid returns a boolean if a field has been set.

### SetSidNil

`func (o *ApiV2010AccountNotificationInstance) SetSidNil(b bool)`

 SetSidNil sets the value for Sid to be an explicit nil

### UnsetSid
`func (o *ApiV2010AccountNotificationInstance) UnsetSid()`

UnsetSid ensures that no value is present for Sid, not even an explicit nil
### GetUri

`func (o *ApiV2010AccountNotificationInstance) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *ApiV2010AccountNotificationInstance) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *ApiV2010AccountNotificationInstance) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *ApiV2010AccountNotificationInstance) HasUri() bool`

HasUri returns a boolean if a field has been set.

### SetUriNil

`func (o *ApiV2010AccountNotificationInstance) SetUriNil(b bool)`

 SetUriNil sets the value for Uri to be an explicit nil

### UnsetUri
`func (o *ApiV2010AccountNotificationInstance) UnsetUri()`

UnsetUri ensures that no value is present for Uri, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


