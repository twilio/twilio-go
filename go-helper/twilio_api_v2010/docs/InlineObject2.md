# InlineObject2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Capture** | **string** | The piece of payment information that you wish the caller to enter. Must be one of &#x60;payment-card-number&#x60;, &#x60;expiration-date&#x60;, &#x60;security-code&#x60;, &#x60;postal-code&#x60;, &#x60;bank-routing-number&#x60;, or &#x60;bank-account-number&#x60;. | [optional] 
**IdempotencyKey** | **string** | A unique token that will be used to ensure that multiple API calls with the same information do not result in multiple transactions. This should be a unique string value per API call and can be a randomly generated. | 
**Status** | **string** | Indicates whether the current payment session should be cancelled or completed. When &#x60;cancel&#x60; the payment session is cancelled. When &#x60;complete&#x60;, Twilio sends the payment information to the selected &lt;Pay&gt; connector for processing. | [optional] 
**StatusCallback** | **string** | Provide an absolute or relative URL to receive status updates regarding your Pay session. Read more about the [Update](https://www.twilio.com/docs/voice/api/payment-resource#statuscallback-update) and [Complete/Cancel](https://www.twilio.com/docs/voice/api/payment-resource#statuscallback-cancelcomplete) POST requests. | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


