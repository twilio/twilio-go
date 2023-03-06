# RegulatoryComplianceProofOfAddressesApi

All URIs are relative to *https://numbers.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateProofOfAddress**](RegulatoryComplianceProofOfAddressesApi.md#CreateProofOfAddress) | **Post** /v1/RegulatoryCompliance/ProofOfAddresses | 
[**FetchProofOfAddress**](RegulatoryComplianceProofOfAddressesApi.md#FetchProofOfAddress) | **Get** /v1/RegulatoryCompliance/ProofOfAddresses/{Sid} | 
[**ListProofOfAddress**](RegulatoryComplianceProofOfAddressesApi.md#ListProofOfAddress) | **Get** /v1/RegulatoryCompliance/ProofOfAddresses | 



## CreateProofOfAddress

> NumbersV1ProofOfAddress CreateProofOfAddress(ctx, optional)



Create an ProofOfAddress for authorizing the hosting of phone number capabilities on our platform.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateProofOfAddressParams struct


Name | Type | Description
------------- | ------------- | -------------
**AddressSid** | **string** | The SID of the [Address](https://www.twilio.com/docs/usage/api/address) resource that is associated with the ProofOfAddress resource to update.
**IdentitySid** | **string** | The SID of the Identity resource that is associated with the ProofOfAddress resource.

### Return type

[**NumbersV1ProofOfAddress**](NumbersV1ProofOfAddress.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchProofOfAddress

> NumbersV1ProofOfAddress FetchProofOfAddress(ctx, Sid)



Fetch a specific ProofOfAddress.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The SID of the ProofOfAddress resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchProofOfAddressParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**NumbersV1ProofOfAddress**](NumbersV1ProofOfAddress.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListProofOfAddress

> []NumbersV1ProofOfAddress ListProofOfAddress(ctx, optional)



Retrieve a list of ProofOfAddresss belonging to the account initiating the request.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListProofOfAddressParams struct


Name | Type | Description
------------- | ------------- | -------------
**AddressSid** | **string** | The SID of the [Address](https://www.twilio.com/docs/usage/api/address) resource that is associated with the ProofOfAddress resources to read.
**IdentitySid** | **string** | The SID of the Identity resource that is associated with the ProofOfAddress resources to read.
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]NumbersV1ProofOfAddress**](NumbersV1ProofOfAddress.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

