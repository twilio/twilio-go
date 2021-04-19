# MonitorV1AlertInstance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **NullableString** | The SID of the Account that created the resource | [optional] 
**AlertText** | Pointer to **NullableString** | The text of the alert | [optional] 
**ApiVersion** | Pointer to **NullableString** | The API version used when the alert was generated | [optional] 
**DateCreated** | Pointer to **NullableTime** | The ISO 8601 date and time in GMT when the resource was created | [optional] 
**DateGenerated** | Pointer to **NullableTime** | The date and time when the alert was generated specified in ISO 8601 format | [optional] 
**DateUpdated** | Pointer to **NullableTime** | The ISO 8601 date and time in GMT when the resource was last updated | [optional] 
**ErrorCode** | Pointer to **NullableString** | The error code for the condition that generated the alert | [optional] 
**LogLevel** | Pointer to **NullableString** | The log level | [optional] 
**MoreInfo** | Pointer to **NullableString** | The URL of the page in our Error Dictionary with more information about the error condition | [optional] 
**RequestHeaders** | Pointer to **NullableString** | The request headers of the request that generated the alert | [optional] 
**RequestMethod** | Pointer to **NullableString** | The method used by the request that generated the alert | [optional] 
**RequestUrl** | Pointer to **NullableString** | The URL of the request that generated the alert | [optional] 
**RequestVariables** | Pointer to **NullableString** | The variables passed in the request that generated the alert | [optional] 
**ResourceSid** | Pointer to **NullableString** | The SID of the resource for which the alert was generated | [optional] 
**ResponseBody** | Pointer to **NullableString** | The response body of the request that generated the alert | [optional] 
**ResponseHeaders** | Pointer to **NullableString** | The response headers of the request that generated the alert | [optional] 
**ServiceSid** | Pointer to **NullableString** | The SID of the service or resource that generated the alert | [optional] 
**Sid** | Pointer to **NullableString** | The unique string that identifies the resource | [optional] 
**Url** | Pointer to **NullableString** | The absolute URL of the Alert resource | [optional] 

## Methods

### NewMonitorV1AlertInstance

`func NewMonitorV1AlertInstance() *MonitorV1AlertInstance`

NewMonitorV1AlertInstance instantiates a new MonitorV1AlertInstance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitorV1AlertInstanceWithDefaults

`func NewMonitorV1AlertInstanceWithDefaults() *MonitorV1AlertInstance`

NewMonitorV1AlertInstanceWithDefaults instantiates a new MonitorV1AlertInstance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSid

`func (o *MonitorV1AlertInstance) GetAccountSid() string`

GetAccountSid returns the AccountSid field if non-nil, zero value otherwise.

### GetAccountSidOk

`func (o *MonitorV1AlertInstance) GetAccountSidOk() (*string, bool)`

GetAccountSidOk returns a tuple with the AccountSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSid

`func (o *MonitorV1AlertInstance) SetAccountSid(v string)`

SetAccountSid sets AccountSid field to given value.

### HasAccountSid

`func (o *MonitorV1AlertInstance) HasAccountSid() bool`

HasAccountSid returns a boolean if a field has been set.

### SetAccountSidNil

`func (o *MonitorV1AlertInstance) SetAccountSidNil(b bool)`

 SetAccountSidNil sets the value for AccountSid to be an explicit nil

### UnsetAccountSid
`func (o *MonitorV1AlertInstance) UnsetAccountSid()`

UnsetAccountSid ensures that no value is present for AccountSid, not even an explicit nil
### GetAlertText

`func (o *MonitorV1AlertInstance) GetAlertText() string`

GetAlertText returns the AlertText field if non-nil, zero value otherwise.

### GetAlertTextOk

`func (o *MonitorV1AlertInstance) GetAlertTextOk() (*string, bool)`

GetAlertTextOk returns a tuple with the AlertText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertText

`func (o *MonitorV1AlertInstance) SetAlertText(v string)`

SetAlertText sets AlertText field to given value.

### HasAlertText

`func (o *MonitorV1AlertInstance) HasAlertText() bool`

HasAlertText returns a boolean if a field has been set.

### SetAlertTextNil

`func (o *MonitorV1AlertInstance) SetAlertTextNil(b bool)`

 SetAlertTextNil sets the value for AlertText to be an explicit nil

### UnsetAlertText
`func (o *MonitorV1AlertInstance) UnsetAlertText()`

UnsetAlertText ensures that no value is present for AlertText, not even an explicit nil
### GetApiVersion

`func (o *MonitorV1AlertInstance) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *MonitorV1AlertInstance) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *MonitorV1AlertInstance) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *MonitorV1AlertInstance) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### SetApiVersionNil

`func (o *MonitorV1AlertInstance) SetApiVersionNil(b bool)`

 SetApiVersionNil sets the value for ApiVersion to be an explicit nil

### UnsetApiVersion
`func (o *MonitorV1AlertInstance) UnsetApiVersion()`

UnsetApiVersion ensures that no value is present for ApiVersion, not even an explicit nil
### GetDateCreated

`func (o *MonitorV1AlertInstance) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *MonitorV1AlertInstance) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *MonitorV1AlertInstance) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *MonitorV1AlertInstance) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### SetDateCreatedNil

`func (o *MonitorV1AlertInstance) SetDateCreatedNil(b bool)`

 SetDateCreatedNil sets the value for DateCreated to be an explicit nil

### UnsetDateCreated
`func (o *MonitorV1AlertInstance) UnsetDateCreated()`

UnsetDateCreated ensures that no value is present for DateCreated, not even an explicit nil
### GetDateGenerated

`func (o *MonitorV1AlertInstance) GetDateGenerated() time.Time`

GetDateGenerated returns the DateGenerated field if non-nil, zero value otherwise.

### GetDateGeneratedOk

`func (o *MonitorV1AlertInstance) GetDateGeneratedOk() (*time.Time, bool)`

GetDateGeneratedOk returns a tuple with the DateGenerated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateGenerated

`func (o *MonitorV1AlertInstance) SetDateGenerated(v time.Time)`

SetDateGenerated sets DateGenerated field to given value.

### HasDateGenerated

`func (o *MonitorV1AlertInstance) HasDateGenerated() bool`

HasDateGenerated returns a boolean if a field has been set.

### SetDateGeneratedNil

`func (o *MonitorV1AlertInstance) SetDateGeneratedNil(b bool)`

 SetDateGeneratedNil sets the value for DateGenerated to be an explicit nil

### UnsetDateGenerated
`func (o *MonitorV1AlertInstance) UnsetDateGenerated()`

UnsetDateGenerated ensures that no value is present for DateGenerated, not even an explicit nil
### GetDateUpdated

`func (o *MonitorV1AlertInstance) GetDateUpdated() time.Time`

GetDateUpdated returns the DateUpdated field if non-nil, zero value otherwise.

### GetDateUpdatedOk

`func (o *MonitorV1AlertInstance) GetDateUpdatedOk() (*time.Time, bool)`

GetDateUpdatedOk returns a tuple with the DateUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateUpdated

`func (o *MonitorV1AlertInstance) SetDateUpdated(v time.Time)`

SetDateUpdated sets DateUpdated field to given value.

### HasDateUpdated

`func (o *MonitorV1AlertInstance) HasDateUpdated() bool`

HasDateUpdated returns a boolean if a field has been set.

### SetDateUpdatedNil

`func (o *MonitorV1AlertInstance) SetDateUpdatedNil(b bool)`

 SetDateUpdatedNil sets the value for DateUpdated to be an explicit nil

### UnsetDateUpdated
`func (o *MonitorV1AlertInstance) UnsetDateUpdated()`

UnsetDateUpdated ensures that no value is present for DateUpdated, not even an explicit nil
### GetErrorCode

`func (o *MonitorV1AlertInstance) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *MonitorV1AlertInstance) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *MonitorV1AlertInstance) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *MonitorV1AlertInstance) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### SetErrorCodeNil

`func (o *MonitorV1AlertInstance) SetErrorCodeNil(b bool)`

 SetErrorCodeNil sets the value for ErrorCode to be an explicit nil

### UnsetErrorCode
`func (o *MonitorV1AlertInstance) UnsetErrorCode()`

UnsetErrorCode ensures that no value is present for ErrorCode, not even an explicit nil
### GetLogLevel

`func (o *MonitorV1AlertInstance) GetLogLevel() string`

GetLogLevel returns the LogLevel field if non-nil, zero value otherwise.

### GetLogLevelOk

`func (o *MonitorV1AlertInstance) GetLogLevelOk() (*string, bool)`

GetLogLevelOk returns a tuple with the LogLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogLevel

`func (o *MonitorV1AlertInstance) SetLogLevel(v string)`

SetLogLevel sets LogLevel field to given value.

### HasLogLevel

`func (o *MonitorV1AlertInstance) HasLogLevel() bool`

HasLogLevel returns a boolean if a field has been set.

### SetLogLevelNil

`func (o *MonitorV1AlertInstance) SetLogLevelNil(b bool)`

 SetLogLevelNil sets the value for LogLevel to be an explicit nil

### UnsetLogLevel
`func (o *MonitorV1AlertInstance) UnsetLogLevel()`

UnsetLogLevel ensures that no value is present for LogLevel, not even an explicit nil
### GetMoreInfo

`func (o *MonitorV1AlertInstance) GetMoreInfo() string`

GetMoreInfo returns the MoreInfo field if non-nil, zero value otherwise.

### GetMoreInfoOk

`func (o *MonitorV1AlertInstance) GetMoreInfoOk() (*string, bool)`

GetMoreInfoOk returns a tuple with the MoreInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoreInfo

`func (o *MonitorV1AlertInstance) SetMoreInfo(v string)`

SetMoreInfo sets MoreInfo field to given value.

### HasMoreInfo

`func (o *MonitorV1AlertInstance) HasMoreInfo() bool`

HasMoreInfo returns a boolean if a field has been set.

### SetMoreInfoNil

`func (o *MonitorV1AlertInstance) SetMoreInfoNil(b bool)`

 SetMoreInfoNil sets the value for MoreInfo to be an explicit nil

### UnsetMoreInfo
`func (o *MonitorV1AlertInstance) UnsetMoreInfo()`

UnsetMoreInfo ensures that no value is present for MoreInfo, not even an explicit nil
### GetRequestHeaders

`func (o *MonitorV1AlertInstance) GetRequestHeaders() string`

GetRequestHeaders returns the RequestHeaders field if non-nil, zero value otherwise.

### GetRequestHeadersOk

`func (o *MonitorV1AlertInstance) GetRequestHeadersOk() (*string, bool)`

GetRequestHeadersOk returns a tuple with the RequestHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestHeaders

`func (o *MonitorV1AlertInstance) SetRequestHeaders(v string)`

SetRequestHeaders sets RequestHeaders field to given value.

### HasRequestHeaders

`func (o *MonitorV1AlertInstance) HasRequestHeaders() bool`

HasRequestHeaders returns a boolean if a field has been set.

### SetRequestHeadersNil

`func (o *MonitorV1AlertInstance) SetRequestHeadersNil(b bool)`

 SetRequestHeadersNil sets the value for RequestHeaders to be an explicit nil

### UnsetRequestHeaders
`func (o *MonitorV1AlertInstance) UnsetRequestHeaders()`

UnsetRequestHeaders ensures that no value is present for RequestHeaders, not even an explicit nil
### GetRequestMethod

`func (o *MonitorV1AlertInstance) GetRequestMethod() string`

GetRequestMethod returns the RequestMethod field if non-nil, zero value otherwise.

### GetRequestMethodOk

`func (o *MonitorV1AlertInstance) GetRequestMethodOk() (*string, bool)`

GetRequestMethodOk returns a tuple with the RequestMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestMethod

`func (o *MonitorV1AlertInstance) SetRequestMethod(v string)`

SetRequestMethod sets RequestMethod field to given value.

### HasRequestMethod

`func (o *MonitorV1AlertInstance) HasRequestMethod() bool`

HasRequestMethod returns a boolean if a field has been set.

### SetRequestMethodNil

`func (o *MonitorV1AlertInstance) SetRequestMethodNil(b bool)`

 SetRequestMethodNil sets the value for RequestMethod to be an explicit nil

### UnsetRequestMethod
`func (o *MonitorV1AlertInstance) UnsetRequestMethod()`

UnsetRequestMethod ensures that no value is present for RequestMethod, not even an explicit nil
### GetRequestUrl

`func (o *MonitorV1AlertInstance) GetRequestUrl() string`

GetRequestUrl returns the RequestUrl field if non-nil, zero value otherwise.

### GetRequestUrlOk

`func (o *MonitorV1AlertInstance) GetRequestUrlOk() (*string, bool)`

GetRequestUrlOk returns a tuple with the RequestUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestUrl

`func (o *MonitorV1AlertInstance) SetRequestUrl(v string)`

SetRequestUrl sets RequestUrl field to given value.

### HasRequestUrl

`func (o *MonitorV1AlertInstance) HasRequestUrl() bool`

HasRequestUrl returns a boolean if a field has been set.

### SetRequestUrlNil

`func (o *MonitorV1AlertInstance) SetRequestUrlNil(b bool)`

 SetRequestUrlNil sets the value for RequestUrl to be an explicit nil

### UnsetRequestUrl
`func (o *MonitorV1AlertInstance) UnsetRequestUrl()`

UnsetRequestUrl ensures that no value is present for RequestUrl, not even an explicit nil
### GetRequestVariables

`func (o *MonitorV1AlertInstance) GetRequestVariables() string`

GetRequestVariables returns the RequestVariables field if non-nil, zero value otherwise.

### GetRequestVariablesOk

`func (o *MonitorV1AlertInstance) GetRequestVariablesOk() (*string, bool)`

GetRequestVariablesOk returns a tuple with the RequestVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestVariables

`func (o *MonitorV1AlertInstance) SetRequestVariables(v string)`

SetRequestVariables sets RequestVariables field to given value.

### HasRequestVariables

`func (o *MonitorV1AlertInstance) HasRequestVariables() bool`

HasRequestVariables returns a boolean if a field has been set.

### SetRequestVariablesNil

`func (o *MonitorV1AlertInstance) SetRequestVariablesNil(b bool)`

 SetRequestVariablesNil sets the value for RequestVariables to be an explicit nil

### UnsetRequestVariables
`func (o *MonitorV1AlertInstance) UnsetRequestVariables()`

UnsetRequestVariables ensures that no value is present for RequestVariables, not even an explicit nil
### GetResourceSid

`func (o *MonitorV1AlertInstance) GetResourceSid() string`

GetResourceSid returns the ResourceSid field if non-nil, zero value otherwise.

### GetResourceSidOk

`func (o *MonitorV1AlertInstance) GetResourceSidOk() (*string, bool)`

GetResourceSidOk returns a tuple with the ResourceSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceSid

`func (o *MonitorV1AlertInstance) SetResourceSid(v string)`

SetResourceSid sets ResourceSid field to given value.

### HasResourceSid

`func (o *MonitorV1AlertInstance) HasResourceSid() bool`

HasResourceSid returns a boolean if a field has been set.

### SetResourceSidNil

`func (o *MonitorV1AlertInstance) SetResourceSidNil(b bool)`

 SetResourceSidNil sets the value for ResourceSid to be an explicit nil

### UnsetResourceSid
`func (o *MonitorV1AlertInstance) UnsetResourceSid()`

UnsetResourceSid ensures that no value is present for ResourceSid, not even an explicit nil
### GetResponseBody

`func (o *MonitorV1AlertInstance) GetResponseBody() string`

GetResponseBody returns the ResponseBody field if non-nil, zero value otherwise.

### GetResponseBodyOk

`func (o *MonitorV1AlertInstance) GetResponseBodyOk() (*string, bool)`

GetResponseBodyOk returns a tuple with the ResponseBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseBody

`func (o *MonitorV1AlertInstance) SetResponseBody(v string)`

SetResponseBody sets ResponseBody field to given value.

### HasResponseBody

`func (o *MonitorV1AlertInstance) HasResponseBody() bool`

HasResponseBody returns a boolean if a field has been set.

### SetResponseBodyNil

`func (o *MonitorV1AlertInstance) SetResponseBodyNil(b bool)`

 SetResponseBodyNil sets the value for ResponseBody to be an explicit nil

### UnsetResponseBody
`func (o *MonitorV1AlertInstance) UnsetResponseBody()`

UnsetResponseBody ensures that no value is present for ResponseBody, not even an explicit nil
### GetResponseHeaders

`func (o *MonitorV1AlertInstance) GetResponseHeaders() string`

GetResponseHeaders returns the ResponseHeaders field if non-nil, zero value otherwise.

### GetResponseHeadersOk

`func (o *MonitorV1AlertInstance) GetResponseHeadersOk() (*string, bool)`

GetResponseHeadersOk returns a tuple with the ResponseHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseHeaders

`func (o *MonitorV1AlertInstance) SetResponseHeaders(v string)`

SetResponseHeaders sets ResponseHeaders field to given value.

### HasResponseHeaders

`func (o *MonitorV1AlertInstance) HasResponseHeaders() bool`

HasResponseHeaders returns a boolean if a field has been set.

### SetResponseHeadersNil

`func (o *MonitorV1AlertInstance) SetResponseHeadersNil(b bool)`

 SetResponseHeadersNil sets the value for ResponseHeaders to be an explicit nil

### UnsetResponseHeaders
`func (o *MonitorV1AlertInstance) UnsetResponseHeaders()`

UnsetResponseHeaders ensures that no value is present for ResponseHeaders, not even an explicit nil
### GetServiceSid

`func (o *MonitorV1AlertInstance) GetServiceSid() string`

GetServiceSid returns the ServiceSid field if non-nil, zero value otherwise.

### GetServiceSidOk

`func (o *MonitorV1AlertInstance) GetServiceSidOk() (*string, bool)`

GetServiceSidOk returns a tuple with the ServiceSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceSid

`func (o *MonitorV1AlertInstance) SetServiceSid(v string)`

SetServiceSid sets ServiceSid field to given value.

### HasServiceSid

`func (o *MonitorV1AlertInstance) HasServiceSid() bool`

HasServiceSid returns a boolean if a field has been set.

### SetServiceSidNil

`func (o *MonitorV1AlertInstance) SetServiceSidNil(b bool)`

 SetServiceSidNil sets the value for ServiceSid to be an explicit nil

### UnsetServiceSid
`func (o *MonitorV1AlertInstance) UnsetServiceSid()`

UnsetServiceSid ensures that no value is present for ServiceSid, not even an explicit nil
### GetSid

`func (o *MonitorV1AlertInstance) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *MonitorV1AlertInstance) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *MonitorV1AlertInstance) SetSid(v string)`

SetSid sets Sid field to given value.

### HasSid

`func (o *MonitorV1AlertInstance) HasSid() bool`

HasSid returns a boolean if a field has been set.

### SetSidNil

`func (o *MonitorV1AlertInstance) SetSidNil(b bool)`

 SetSidNil sets the value for Sid to be an explicit nil

### UnsetSid
`func (o *MonitorV1AlertInstance) UnsetSid()`

UnsetSid ensures that no value is present for Sid, not even an explicit nil
### GetUrl

`func (o *MonitorV1AlertInstance) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *MonitorV1AlertInstance) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *MonitorV1AlertInstance) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *MonitorV1AlertInstance) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *MonitorV1AlertInstance) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *MonitorV1AlertInstance) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


