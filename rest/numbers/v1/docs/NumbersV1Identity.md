# NumbersV1Identity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | Pointer to **string** | The unique string that we created to identify the Identity resource. |
**AccountSid** | Pointer to **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Identity resource. |
**FriendlyName** | Pointer to **string** | The string that you assigned to describe the resource. |
**Type** | Pointer to **string** | The type of Identity resource. |
**Email** | Pointer to **string** | The email address. |
**Status** | Pointer to **string** | The verification status of the Identity resource. |
**StatusCallbackUrl** | Pointer to **string** | The URL we call using the `status_callback_method` to inform your application of status changes. |
**StatusCallbackMethod** | Pointer to **string** | The HTTP method we use when calling `status_callback_url`. |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The date and time in GMT when the resource was created specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The date and time in GMT when the resource was last updated specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. |
**Url** | Pointer to **string** | The absolute URL of the Identity resource. |
**PurchaseIntentNumberType** | Pointer to **string** | The type of phone number. |
**PurchaseIntentIsoCountry** | Pointer to **string** | The ISO country code of the country. |
**Links** | Pointer to **map[string]interface{}** | The URLs of the documents associated with the Identity resource. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


