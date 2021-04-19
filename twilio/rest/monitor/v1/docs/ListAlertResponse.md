# ListAlertResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alerts** | Pointer to [**[]MonitorV1Alert**](MonitorV1Alert.md) |  | [optional] 
**Meta** | Pointer to [**ListAlertResponseMeta**](ListAlertResponse_meta.md) |  | [optional] 

## Methods

### NewListAlertResponse

`func NewListAlertResponse() *ListAlertResponse`

NewListAlertResponse instantiates a new ListAlertResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListAlertResponseWithDefaults

`func NewListAlertResponseWithDefaults() *ListAlertResponse`

NewListAlertResponseWithDefaults instantiates a new ListAlertResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlerts

`func (o *ListAlertResponse) GetAlerts() []MonitorV1Alert`

GetAlerts returns the Alerts field if non-nil, zero value otherwise.

### GetAlertsOk

`func (o *ListAlertResponse) GetAlertsOk() (*[]MonitorV1Alert, bool)`

GetAlertsOk returns a tuple with the Alerts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlerts

`func (o *ListAlertResponse) SetAlerts(v []MonitorV1Alert)`

SetAlerts sets Alerts field to given value.

### HasAlerts

`func (o *ListAlertResponse) HasAlerts() bool`

HasAlerts returns a boolean if a field has been set.

### GetMeta

`func (o *ListAlertResponse) GetMeta() ListAlertResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListAlertResponse) GetMetaOk() (*ListAlertResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListAlertResponse) SetMeta(v ListAlertResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListAlertResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


