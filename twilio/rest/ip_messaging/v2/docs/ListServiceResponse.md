# ListServiceResponse

## Properties

Name | Type | Description
------------ | ------------- | -------------
**Meta** | Pointer to [**ListCredentialResponseMeta**](ListCredentialResponse_meta.md) |  | [optional] 
**Services** | Pointer to [**[]IpMessagingV2Service**](IpMessagingV2Service.md) |  | [optional] 

## Methods

### NewListServiceResponse

`func NewListServiceResponse() *ListServiceResponse`

NewListServiceResponse instantiates a new ListServiceResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListServiceResponseWithDefaults

`func NewListServiceResponseWithDefaults() *ListServiceResponse`

NewListServiceResponseWithDefaults instantiates a new ListServiceResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *ListServiceResponse) GetMeta() ListCredentialResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListServiceResponse) GetMetaOk() (*ListCredentialResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListServiceResponse) SetMeta(v ListCredentialResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListServiceResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetServices

`func (o *ListServiceResponse) GetServices() []IpMessagingV2Service`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *ListServiceResponse) GetServicesOk() (*[]IpMessagingV2Service, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *ListServiceResponse) SetServices(v []IpMessagingV2Service)`

SetServices sets Services field to given value.

### HasServices

`func (o *ListServiceResponse) HasServices() bool`

HasServices returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


