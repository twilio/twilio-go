# ServicesUsersApi

All URIs are relative to *https://notify.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateUser**](ServicesUsersApi.md#CreateUser) | **Post** /v1/Services/{ServiceSid}/Users | 
[**DeleteUser**](ServicesUsersApi.md#DeleteUser) | **Delete** /v1/Services/{ServiceSid}/Users/{Identity} | 
[**FetchUser**](ServicesUsersApi.md#FetchUser) | **Get** /v1/Services/{ServiceSid}/Users/{Identity} | 
[**ListUser**](ServicesUsersApi.md#ListUser) | **Get** /v1/Services/{ServiceSid}/Users | 



## CreateUser

> NotifyV1User CreateUser(ctx, ServiceSidoptional)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Service](https://www.twilio.com/docs/notify/api/service-resource) to create the resource under.

### Other Parameters

Other parameters are passed through a pointer to a CreateUserParams struct


Name | Type | Description
------------- | ------------- | -------------
**Identity** | **string** | The application-defined string that uniquely identifies the resource's User within the Service. See [access tokens](https://www.twilio.com/docs/chat/create-tokens) for more info.
**Segment** | **[]string** | The list of segments this User belongs to. Segments can be used to select recipients of a notification. Maximum 20 Segments per User allowed.

### Return type

[**NotifyV1User**](NotifyV1User.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteUser

> DeleteUser(ctx, ServiceSidIdentity)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Service](https://www.twilio.com/docs/notify/api/service-resource) to delete the resource from.
**Identity** | **string** | The application-defined string that uniquely identifies the resource to delete. See [access tokens](https://www.twilio.com/docs/chat/create-tokens) for more info.

### Other Parameters

Other parameters are passed through a pointer to a DeleteUserParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

 (empty response body)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchUser

> NotifyV1User FetchUser(ctx, ServiceSidIdentity)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Service](https://www.twilio.com/docs/notify/api/service-resource) to fetch the resource from.
**Identity** | **string** | The application-defined string that uniquely identifies the resource to fetch. See [access tokens](https://www.twilio.com/docs/chat/create-tokens) for more info.

### Other Parameters

Other parameters are passed through a pointer to a FetchUserParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**NotifyV1User**](NotifyV1User.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUser

> []NotifyV1User ListUser(ctx, ServiceSidoptional)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Service](https://www.twilio.com/docs/notify/api/service-resource) to read the resource from.

### Other Parameters

Other parameters are passed through a pointer to a ListUserParams struct


Name | Type | Description
------------- | ------------- | -------------
**Identity** | **[]string** | The application-defined string that uniquely identifies the resource to read. See [access tokens](https://www.twilio.com/docs/chat/create-tokens) for more info.
**Segment** | **string** | The list of segments this User belongs to. Segments can be used to select recipients of a notification. Maximum 20 Segments per User allowed.
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]NotifyV1User**](NotifyV1User.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

