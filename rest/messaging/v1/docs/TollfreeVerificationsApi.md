# TollfreeVerificationsApi

All URIs are relative to *https://messaging.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTollfreeVerification**](TollfreeVerificationsApi.md#CreateTollfreeVerification) | **Post** /v1/Tollfree/Verifications | 
[**DeleteTollfreeVerification**](TollfreeVerificationsApi.md#DeleteTollfreeVerification) | **Delete** /v1/Tollfree/Verifications/{Sid} | 
[**FetchTollfreeVerification**](TollfreeVerificationsApi.md#FetchTollfreeVerification) | **Get** /v1/Tollfree/Verifications/{Sid} | 
[**ListTollfreeVerification**](TollfreeVerificationsApi.md#ListTollfreeVerification) | **Get** /v1/Tollfree/Verifications | 
[**UpdateTollfreeVerification**](TollfreeVerificationsApi.md#UpdateTollfreeVerification) | **Post** /v1/Tollfree/Verifications/{Sid} | 



## CreateTollfreeVerification

> MessagingV1TollfreeVerification CreateTollfreeVerification(ctx, optional)





### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateTollfreeVerificationParams struct


Name | Type | Description
------------- | ------------- | -------------
**BusinessName** | **string** | The name of the business or organization using the Tollfree number.
**BusinessWebsite** | **string** | The website of the business or organization using the Tollfree number.
**NotificationEmail** | **string** | The email address to receive the notification about the verification result. .
**UseCaseCategories** | **[]string** | The category of the use case for the Tollfree Number. List as many are applicable..
**UseCaseSummary** | **string** | Use this to further explain how messaging is used by the business or organization.
**ProductionMessageSample** | **string** | An example of message content, i.e. a sample message.
**OptInImageUrls** | **[]string** | Link to an image that shows the opt-in workflow. Multiple images allowed and must be a publicly hosted URL.
**OptInType** | **string** | 
**MessageVolume** | **string** | Estimate monthly volume of messages from the Tollfree Number.
**TollfreePhoneNumberSid** | **string** | The SID of the Phone Number associated with the Tollfree Verification.
**CustomerProfileSid** | **string** | Customer's Profile Bundle BundleSid.
**BusinessStreetAddress** | **string** | The address of the business or organization using the Tollfree number.
**BusinessStreetAddress2** | **string** | The address of the business or organization using the Tollfree number.
**BusinessCity** | **string** | The city of the business or organization using the Tollfree number.
**BusinessStateProvinceRegion** | **string** | The state/province/region of the business or organization using the Tollfree number.
**BusinessPostalCode** | **string** | The postal code of the business or organization using the Tollfree number.
**BusinessCountry** | **string** | The country of the business or organization using the Tollfree number.
**AdditionalInformation** | **string** | Additional information to be provided for verification.
**BusinessContactFirstName** | **string** | The first name of the contact for the business or organization using the Tollfree number.
**BusinessContactLastName** | **string** | The last name of the contact for the business or organization using the Tollfree number.
**BusinessContactEmail** | **string** | The email address of the contact for the business or organization using the Tollfree number.
**BusinessContactPhone** | **string** | The E.164 formatted phone number of the contact for the business or organization using the Tollfree number.
**ExternalReferenceId** | **string** | An optional external reference ID supplied by customer and echoed back on status retrieval.

### Return type

[**MessagingV1TollfreeVerification**](MessagingV1TollfreeVerification.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTollfreeVerification

> DeleteTollfreeVerification(ctx, Sid)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The unique string to identify Tollfree Verification.

### Other Parameters

Other parameters are passed through a pointer to a DeleteTollfreeVerificationParams struct


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


## FetchTollfreeVerification

> MessagingV1TollfreeVerification FetchTollfreeVerification(ctx, Sid)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | A unique string identifying a Tollfree Verification.

### Other Parameters

Other parameters are passed through a pointer to a FetchTollfreeVerificationParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**MessagingV1TollfreeVerification**](MessagingV1TollfreeVerification.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTollfreeVerification

> []MessagingV1TollfreeVerification ListTollfreeVerification(ctx, optional)





### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListTollfreeVerificationParams struct


Name | Type | Description
------------- | ------------- | -------------
**TollfreePhoneNumberSid** | **string** | The SID of the Phone Number associated with the Tollfree Verification.
**Status** | **string** | The compliance status of the Tollfree Verification record.
**ExternalReferenceId** | **string** | Customer supplied reference id for the Tollfree Verification record.
**IncludeSubAccounts** | **bool** | Whether to include Tollfree Verifications from sub accounts in list response.
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]MessagingV1TollfreeVerification**](MessagingV1TollfreeVerification.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTollfreeVerification

> MessagingV1TollfreeVerification UpdateTollfreeVerification(ctx, Sidoptional)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The unique string to identify Tollfree Verification.

### Other Parameters

Other parameters are passed through a pointer to a UpdateTollfreeVerificationParams struct


Name | Type | Description
------------- | ------------- | -------------
**BusinessName** | **string** | The name of the business or organization using the Tollfree number.
**BusinessWebsite** | **string** | The website of the business or organization using the Tollfree number.
**NotificationEmail** | **string** | The email address to receive the notification about the verification result. .
**UseCaseCategories** | **[]string** | The category of the use case for the Tollfree Number. List as many are applicable..
**UseCaseSummary** | **string** | Use this to further explain how messaging is used by the business or organization.
**ProductionMessageSample** | **string** | An example of message content, i.e. a sample message.
**OptInImageUrls** | **[]string** | Link to an image that shows the opt-in workflow. Multiple images allowed and must be a publicly hosted URL.
**OptInType** | **string** | 
**MessageVolume** | **string** | Estimate monthly volume of messages from the Tollfree Number.
**BusinessStreetAddress** | **string** | The address of the business or organization using the Tollfree number.
**BusinessStreetAddress2** | **string** | The address of the business or organization using the Tollfree number.
**BusinessCity** | **string** | The city of the business or organization using the Tollfree number.
**BusinessStateProvinceRegion** | **string** | The state/province/region of the business or organization using the Tollfree number.
**BusinessPostalCode** | **string** | The postal code of the business or organization using the Tollfree number.
**BusinessCountry** | **string** | The country of the business or organization using the Tollfree number.
**AdditionalInformation** | **string** | Additional information to be provided for verification.
**BusinessContactFirstName** | **string** | The first name of the contact for the business or organization using the Tollfree number.
**BusinessContactLastName** | **string** | The last name of the contact for the business or organization using the Tollfree number.
**BusinessContactEmail** | **string** | The email address of the contact for the business or organization using the Tollfree number.
**BusinessContactPhone** | **string** | The E.164 formatted phone number of the contact for the business or organization using the Tollfree number.
**EditReason** | **string** | Describe why the verification is being edited. If the verification was rejected because of a technical issue, such as the website being down, and the issue has been resolved this parameter should be set to something similar to 'Website fixed'.

### Return type

[**MessagingV1TollfreeVerification**](MessagingV1TollfreeVerification.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

