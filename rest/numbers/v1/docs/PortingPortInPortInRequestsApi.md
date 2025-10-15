# PortingPortInPortInRequestsApi

All URIs are relative to *https://numbers.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListPortInRequests**](PortingPortInPortInRequestsApi.md#ListPortInRequests) | **Get** /v1/Porting/PortIn/PortInRequests | Fetch all PortInRequests for a user



## ListPortInRequests

> []NumbersV1PortInRequestList ListPortInRequests(ctx, optional)

Fetch all PortInRequests for a user

Retrieve a list of all PortInRequests for a user

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListPortInRequestsParams struct


Name | Type | Description
------------- | ------------- | -------------
**Token** | **string** | Page start token, if null then it will start from the beginning
**Size** | **int** | Number of items per page
**PortInRequestSid** | **string** | Filter by Port in request SID, supports multiple values separated by comma
**PortInRequestStatus** | **string** | Filter by Port In request status
**CreatedBefore** | **time.Time** | Find all created before a certain date
**CreatedAfter** | **time.Time** | Find all created after a certain date
**Limit** | **int** | Max number of records to return.
**PageSize** | **int** | Max number of records to return in a page

### Return type

[**[]NumbersV1PortInRequestList**](NumbersV1PortInRequestList.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

