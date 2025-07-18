# ConsentsBulkApi

All URIs are relative to *https://accounts.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBulkConsents**](ConsentsBulkApi.md#CreateBulkConsents) | **Post** /v1/Consents/Bulk | 



## CreateBulkConsents

> AccountsV1BulkConsents CreateBulkConsents(ctx, optional)





### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateBulkConsentsParams struct


Name | Type | Description
------------- | ------------- | -------------
**Items** | **[]interface{}** | This is a list of objects that describes a contact's opt-in status. Each object contains the following fields: `contact_id`, which must be a string representing phone number in [E.164 format](https://www.twilio.com/docs/glossary/what-e164); `correlation_id`, a unique 32-character UUID used to uniquely map the request item with the response item; `sender_id`, which can be either a valid messaging service SID or a from phone number; `status`, a string representing the consent status. Can be one of [`opt-in`, `opt-out`]; `source`, a string indicating the medium through which the consent was collected. Can be one of [`website`, `offline`, `opt-in-message`, `opt-out-message`, `others`]; `date_of_consent`, an optional datetime string field in ISO-8601 format that captures the exact date and time when the user gave or revoked consent. If not provided, it will be empty.

### Return type

[**AccountsV1BulkConsents**](AccountsV1BulkConsents.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

