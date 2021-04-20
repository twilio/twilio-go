# DefaultApi

All URIs are relative to *https://fax.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateFax**](DefaultApi.md#CreateFax) | **Post** /v1/Faxes | 
[**DeleteFax**](DefaultApi.md#DeleteFax) | **Delete** /v1/Faxes/{Sid} | 
[**DeleteFaxMedia**](DefaultApi.md#DeleteFaxMedia) | **Delete** /v1/Faxes/{FaxSid}/Media/{Sid} | 
[**FetchFax**](DefaultApi.md#FetchFax) | **Get** /v1/Faxes/{Sid} | 
[**FetchFaxMedia**](DefaultApi.md#FetchFaxMedia) | **Get** /v1/Faxes/{FaxSid}/Media/{Sid} | 
[**ListFax**](DefaultApi.md#ListFax) | **Get** /v1/Faxes | 
[**ListFaxMedia**](DefaultApi.md#ListFaxMedia) | **Get** /v1/Faxes/{FaxSid}/Media | 
[**UpdateFax**](DefaultApi.md#UpdateFax) | **Post** /v1/Faxes/{Sid} | 



## CreateFax

> FaxV1Fax CreateFax(ctx, optional)



Create a new fax to send to a phone number or SIP endpoint.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateFaxParams struct


Name | Type | Description
------------- | ------------- | -------------
**From** | **string** | The number the fax was sent from. Can be the phone number in [E.164](https://www.twilio.com/docs/glossary/what-e164) format or the SIP &#x60;from&#x60; value. The caller ID displayed to the recipient uses this value. If this is a phone number, it must be a Twilio number or a verified outgoing caller id from your account. If &#x60;to&#x60; is a SIP address, this can be any alphanumeric string (and also the characters &#x60;+&#x60;, &#x60;_&#x60;, &#x60;.&#x60;, and &#x60;-&#x60;), which will be used in the &#x60;from&#x60; header of the SIP request.
**MediaUrl** | **string** | The URL of the PDF that contains the fax. See our [security](https://www.twilio.com/docs/usage/security) page for information on how to ensure the request for your media comes from Twilio.
**Quality** | **string** | The [Fax Quality value](https://www.twilio.com/docs/fax/api/fax-resource#fax-quality-values) that describes the fax quality. Can be: &#x60;standard&#x60;, &#x60;fine&#x60;, or &#x60;superfine&#x60; and defaults to &#x60;fine&#x60;.
**SipAuthPassword** | **string** | The password to use with &#x60;sip_auth_username&#x60; to authenticate faxes sent to a SIP address.
**SipAuthUsername** | **string** | The username to use with the &#x60;sip_auth_password&#x60; to authenticate faxes sent to a SIP address.
**StatusCallback** | **string** | The URL we should call using the &#x60;POST&#x60; method to send [status information](https://www.twilio.com/docs/fax/api/fax-resource#fax-status-callback) to your application when the status of the fax changes.
**StoreMedia** | **bool** | Whether to store a copy of the sent media on our servers for later retrieval. Can be: &#x60;true&#x60; or &#x60;false&#x60; and the default is &#x60;true&#x60;.
**To** | **string** | The phone number to receive the fax in [E.164](https://www.twilio.com/docs/glossary/what-e164) format or the recipient&#39;s SIP URI.
**Ttl** | **int32** | How long in minutes from when the fax is initiated that we should try to send the fax.

### Return type

[**FaxV1Fax**](FaxV1Fax.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded, 
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteFax

> DeleteFax(ctx, Sid)



Delete a specific fax and its associated media.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the Fax resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeleteFaxParams struct


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


## DeleteFaxMedia

> DeleteFaxMedia(ctx, FaxSidSid)



Delete a specific fax media instance.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**FaxSid** | **string** | The SID of the fax with the FaxMedia resource to delete.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the FaxMedia resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeleteFaxMediaParams struct


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


## FetchFax

> FaxV1Fax FetchFax(ctx, Sid)



Fetch a specific fax.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the Fax resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchFaxParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**FaxV1Fax**](FaxV1Fax.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchFaxMedia

> FaxV1FaxFaxMedia FetchFaxMedia(ctx, FaxSidSid)



Fetch a specific fax media instance.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**FaxSid** | **string** | The SID of the fax with the FaxMedia resource to fetch.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the FaxMedia resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchFaxMediaParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**FaxV1FaxFaxMedia**](FaxV1FaxFaxMedia.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListFax

> ListFaxResponse ListFax(ctx, optional)



Retrieve a list of all faxes.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListFaxParams struct


Name | Type | Description
------------- | ------------- | -------------
**From** | **string** | Retrieve only those faxes sent from this phone number, specified in [E.164](https://www.twilio.com/docs/glossary/what-e164) format.
**To** | **string** | Retrieve only those faxes sent to this phone number, specified in [E.164](https://www.twilio.com/docs/glossary/what-e164) format.
**DateCreatedOnOrBefore** | **time.Time** | Retrieve only those faxes with a &#x60;date_created&#x60; that is before or equal to this value, specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format.
**DateCreatedAfter** | **time.Time** | Retrieve only those faxes with a &#x60;date_created&#x60; that is later than this value, specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format.
**PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListFaxResponse**](ListFaxResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListFaxMedia

> ListFaxMediaResponse ListFaxMedia(ctx, FaxSidoptional)



Retrieve a list of all fax media instances for the specified fax.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**FaxSid** | **string** | The SID of the fax with the FaxMedia resources to read.

### Other Parameters

Other parameters are passed through a pointer to a ListFaxMediaParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListFaxMediaResponse**](ListFaxMediaResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateFax

> FaxV1Fax UpdateFax(ctx, Sidoptional)



Update a specific fax.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the Fax resource to update.

### Other Parameters

Other parameters are passed through a pointer to a UpdateFaxParams struct


Name | Type | Description
------------- | ------------- | -------------
**Status** | **string** | The new [status](https://www.twilio.com/docs/fax/api/fax-resource#fax-status-values) of the resource. Can be only &#x60;canceled&#x60;. This may fail if transmission has already started.

### Return type

[**FaxV1Fax**](FaxV1Fax.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded, 
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

