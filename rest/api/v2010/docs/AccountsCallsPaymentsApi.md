# AccountsCallsPaymentsApi

All URIs are relative to *https://api.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePayments**](AccountsCallsPaymentsApi.md#CreatePayments) | **Post** /2010-04-01/Accounts/{AccountSid}/Calls/{CallSid}/Payments.json | 
[**UpdatePayments**](AccountsCallsPaymentsApi.md#UpdatePayments) | **Post** /2010-04-01/Accounts/{AccountSid}/Calls/{CallSid}/Payments/{Sid}.json | 



## CreatePayments

> ApiV2010Payments CreatePayments(ctx, CallSidoptional)



create an instance of payments. This will start a new payments session

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**CallSid** | **string** | The SID of the call that will create the resource. Call leg associated with this sid is expected to provide payment information thru DTMF.

### Other Parameters

Other parameters are passed through a pointer to a CreatePaymentsParams struct


Name | Type | Description
------------- | ------------- | -------------
**PathAccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that will create the resource.
**BankAccountType** | **string** | Type of bank account if payment source is ACH. One of &#x60;consumer-checking&#x60;, &#x60;consumer-savings&#x60;, or &#x60;commercial-checking&#x60;. The default value is &#x60;consumer-checking&#x60;.
**ChargeAmount** | **float32** | A positive decimal value less than 1,000,000 to charge against the credit card or bank account. Default currency can be overwritten with &#x60;currency&#x60; field. Leave blank or set to 0 to tokenize.
**Currency** | **string** | The currency of the &#x60;charge_amount&#x60;, formatted as [ISO 4127](http://www.iso.org/iso/home/standards/currency_codes.htm) format. The default value is &#x60;USD&#x60; and all values allowed from the &lt;Pay&gt; Connector are accepted.
**Description** | **string** | The description can be used to provide more details regarding the transaction. This information is submitted along with the payment details to the Payment Connector which are then posted on the transactions.
**IdempotencyKey** | **string** | A unique token that will be used to ensure that multiple API calls with the same information do not result in multiple transactions. This should be a unique string value per API call and can be a randomly generated.
**Input** | **string** | A list of inputs that should be accepted. Currently only &#x60;dtmf&#x60; is supported. All digits captured during a pay session are redacted from the logs.
**MinPostalCodeLength** | **int** | A positive integer that is used to validate the length of the &#x60;PostalCode&#x60; inputted by the user. User must enter this many digits.
**Parameter** | [**map[string]interface{}**](map[string]interface{}.md) | A single-level JSON object used to pass custom parameters to payment processors. (Required for ACH payments). The information that has to be included here depends on the &lt;Pay&gt; Connector. [Read more](https://www.twilio.com/console/voice/pay-connectors).
**PaymentConnector** | **string** | This is the unique name corresponding to the Payment Gateway Connector installed in the Twilio Add-ons. Learn more about [&lt;Pay&gt; Connectors](https://www.twilio.com/console/voice/pay-connectors). The default value is &#x60;Default&#x60;.
**PaymentMethod** | **string** | Type of payment being captured. One of &#x60;credit-card&#x60; or &#x60;ach-debit&#x60;. The default value is &#x60;credit-card&#x60;.
**PostalCode** | **bool** | Indicates whether the credit card postal code (zip code) is a required piece of payment information that must be provided by the caller. The default is &#x60;true&#x60;.
**SecurityCode** | **bool** | Indicates whether the credit card security code is a required piece of payment information that must be provided by the caller. The default is &#x60;true&#x60;.
**StatusCallback** | **string** | Provide an absolute or relative URL to receive status updates regarding your Pay session. Read more about the [expected StatusCallback values](https://www.twilio.com/docs/voice/api/payment-resource#statuscallback)
**Timeout** | **int** | The number of seconds that &lt;Pay&gt; should wait for the caller to press a digit between each subsequent digit, after the first one, before moving on to validate the digits captured. The default is &#x60;5&#x60;, maximum is &#x60;600&#x60;.
**TokenType** | **string** | Indicates whether the payment method should be tokenized as a &#x60;one-time&#x60; or &#x60;reusable&#x60; token. The default value is &#x60;reusable&#x60;. Do not enter a charge amount when tokenizing. If a charge amount is entered, the payment method will be charged and not tokenized.
**ValidCardTypes** | **string** | Credit card types separated by space that Pay should accept. The default value is &#x60;visa mastercard amex&#x60;

### Return type

[**ApiV2010Payments**](ApiV2010Payments.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePayments

> ApiV2010Payments UpdatePayments(ctx, CallSidSidoptional)



update an instance of payments with different phases of payment flows.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**CallSid** | **string** | The SID of the call that will update the resource. This should be the same call sid that was used to create payments resource.
**Sid** | **string** | The SID of Payments session that needs to be updated.

### Other Parameters

Other parameters are passed through a pointer to a UpdatePaymentsParams struct


Name | Type | Description
------------- | ------------- | -------------
**PathAccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that will update the resource.
**Capture** | **string** | The piece of payment information that you wish the caller to enter. Must be one of &#x60;payment-card-number&#x60;, &#x60;expiration-date&#x60;, &#x60;security-code&#x60;, &#x60;postal-code&#x60;, &#x60;bank-routing-number&#x60;, or &#x60;bank-account-number&#x60;.
**IdempotencyKey** | **string** | A unique token that will be used to ensure that multiple API calls with the same information do not result in multiple transactions. This should be a unique string value per API call and can be a randomly generated.
**Status** | **string** | Indicates whether the current payment session should be cancelled or completed. When &#x60;cancel&#x60; the payment session is cancelled. When &#x60;complete&#x60;, Twilio sends the payment information to the selected &lt;Pay&gt; connector for processing.
**StatusCallback** | **string** | Provide an absolute or relative URL to receive status updates regarding your Pay session. Read more about the [Update](https://www.twilio.com/docs/voice/api/payment-resource#statuscallback-update) and [Complete/Cancel](https://www.twilio.com/docs/voice/api/payment-resource#statuscallback-cancelcomplete) POST requests.

### Return type

[**ApiV2010Payments**](ApiV2010Payments.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

