# ListSessionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**ListServiceResponseMeta**](ListServiceResponse_meta.md) |  | [optional] 
**Sessions** | Pointer to [**[]ProxyV1ServiceSession**](ProxyV1ServiceSession.md) |  | [optional] 

## Methods

### NewListSessionResponse

`func NewListSessionResponse() *ListSessionResponse`

NewListSessionResponse instantiates a new ListSessionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListSessionResponseWithDefaults

`func NewListSessionResponseWithDefaults() *ListSessionResponse`

NewListSessionResponseWithDefaults instantiates a new ListSessionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *ListSessionResponse) GetMeta() ListServiceResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListSessionResponse) GetMetaOk() (*ListServiceResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListSessionResponse) SetMeta(v ListServiceResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListSessionResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetSessions

`func (o *ListSessionResponse) GetSessions() []ProxyV1ServiceSession`

GetSessions returns the Sessions field if non-nil, zero value otherwise.

### GetSessionsOk

`func (o *ListSessionResponse) GetSessionsOk() (*[]ProxyV1ServiceSession, bool)`

GetSessionsOk returns a tuple with the Sessions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessions

`func (o *ListSessionResponse) SetSessions(v []ProxyV1ServiceSession)`

SetSessions sets Sessions field to given value.

### HasSessions

`func (o *ListSessionResponse) HasSessions() bool`

HasSessions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


