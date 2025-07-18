# ContactsBulkApi

All URIs are relative to *https://accounts.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBulkContacts**](ContactsBulkApi.md#CreateBulkContacts) | **Post** /v1/Contacts/Bulk | 



## CreateBulkContacts

> AccountsV1BulkContacts CreateBulkContacts(ctx, optional)





### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateBulkContactsParams struct


Name | Type | Description
------------- | ------------- | -------------
**Items** | **[]interface{}** | A list of objects where each object represents a contact's details. Each object includes the following fields: `contact_id`, which must be a string representing phone number in [E.164 format](https://www.twilio.com/docs/glossary/what-e164); `correlation_id`, a unique 32-character UUID that maps the response to the original request; `country_iso_code`, a string representing the country using the ISO format (e.g., US for the United States); and `zip_code`, a string representing the postal code.

### Return type

[**AccountsV1BulkContacts**](AccountsV1BulkContacts.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

