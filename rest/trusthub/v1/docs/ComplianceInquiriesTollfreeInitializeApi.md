# ComplianceInquiriesTollfreeInitializeApi

All URIs are relative to *https://trusthub.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateComplianceTollfreeInquiry**](ComplianceInquiriesTollfreeInitializeApi.md#CreateComplianceTollfreeInquiry) | **Post** /v1/ComplianceInquiries/Tollfree/Initialize | 



## CreateComplianceTollfreeInquiry

> TrusthubV1ComplianceTollfreeInquiry CreateComplianceTollfreeInquiry(ctx, optional)



Create a new Compliance Tollfree Verification Inquiry for the authenticated account. This is necessary to start a new embedded session.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateComplianceTollfreeInquiryParams struct


Name | Type | Description
------------- | ------------- | -------------
**TollfreePhoneNumber** | **string** | The Tollfree phone number to be verified
**NotificationEmail** | **string** | The email address to receive the notification about the verification result.
**BusinessName** | **string** | The name of the business or organization using the Tollfree number.
**BusinessWebsite** | **string** | The website of the business or organization using the Tollfree number.
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
**BusinessContactPhone** | **string** | The phone number of the contact for the business or organization using the Tollfree number.
**ThemeSetId** | **string** | Theme id for styling the inquiry form.
**SkipMessagingUseCase** | **bool** | Skip the messaging use case screen of the inquiry form.

### Return type

[**TrusthubV1ComplianceTollfreeInquiry**](TrusthubV1ComplianceTollfreeInquiry.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

