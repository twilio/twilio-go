# ApiV2010IncomingPhoneNumberAssignedAddOnExtension

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | Pointer to **string** | The unique string that that we created to identify the resource. |
**AccountSid** | Pointer to **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the resource. |
**ResourceSid** | Pointer to **string** | The SID of the Phone Number to which the Add-on is assigned. |
**AssignedAddOnSid** | Pointer to **string** | The SID that uniquely identifies the assigned Add-on installation. |
**FriendlyName** | Pointer to **string** | The string that you assigned to describe the resource. |
**ProductName** | Pointer to **string** | A string that you assigned to describe the Product this Extension is used within. |
**UniqueName** | Pointer to **string** | An application-defined string that uniquely identifies the resource. It can be used in place of the resource's `sid` in the URL to address the resource. |
**Uri** | Pointer to **string** | The URI of the resource, relative to `https://api.twilio.com`. |
**Enabled** | Pointer to **bool** | Whether the Extension will be invoked. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


