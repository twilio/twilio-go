# ListLogResponse

## Properties

Name | Type | Description
------------ | ------------- | -------------
**Logs** | Pointer to [**[]ServerlessV1ServiceEnvironmentLog**](ServerlessV1ServiceEnvironmentLog.md) |  | [optional] 
**Meta** | Pointer to [**ListServiceResponseMeta**](ListServiceResponse_meta.md) |  | [optional] 

## Methods

### NewListLogResponse

`func NewListLogResponse() *ListLogResponse`

NewListLogResponse instantiates a new ListLogResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListLogResponseWithDefaults

`func NewListLogResponseWithDefaults() *ListLogResponse`

NewListLogResponseWithDefaults instantiates a new ListLogResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLogs

`func (o *ListLogResponse) GetLogs() []ServerlessV1ServiceEnvironmentLog`

GetLogs returns the Logs field if non-nil, zero value otherwise.

### GetLogsOk

`func (o *ListLogResponse) GetLogsOk() (*[]ServerlessV1ServiceEnvironmentLog, bool)`

GetLogsOk returns a tuple with the Logs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogs

`func (o *ListLogResponse) SetLogs(v []ServerlessV1ServiceEnvironmentLog)`

SetLogs sets Logs field to given value.

### HasLogs

`func (o *ListLogResponse) HasLogs() bool`

HasLogs returns a boolean if a field has been set.

### GetMeta

`func (o *ListLogResponse) GetMeta() ListServiceResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListLogResponse) GetMetaOk() (*ListServiceResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListLogResponse) SetMeta(v ListServiceResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListLogResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


