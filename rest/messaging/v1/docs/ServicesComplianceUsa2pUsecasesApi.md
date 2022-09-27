# ServicesComplianceUsa2pUsecasesApi

All URIs are relative to *https://messaging.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchUsAppToPersonUsecase**](ServicesComplianceUsa2pUsecasesApi.md#FetchUsAppToPersonUsecase) | **Get** /v1/Services/{MessagingServiceSid}/Compliance/Usa2p/Usecases | 



## FetchUsAppToPersonUsecase

> MessagingV1UsAppToPersonUsecase FetchUsAppToPersonUsecase(ctx, MessagingServiceSidoptional)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**MessagingServiceSid** | **string** | The SID of the [Messaging Service](https://www.twilio.com/docs/messaging/services/api) to fetch the resource from.

### Other Parameters

Other parameters are passed through a pointer to a FetchUsAppToPersonUsecaseParams struct


Name | Type | Description
------------- | ------------- | -------------
**BrandRegistrationSid** | **string** | The unique string to identify the A2P brand.

### Return type

[**MessagingV1UsAppToPersonUsecase**](MessagingV1UsAppToPersonUsecase.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

