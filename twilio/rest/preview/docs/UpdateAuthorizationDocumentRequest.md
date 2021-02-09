# UpdateAuthorizationDocumentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddressSid** | **string** | A 34 character string that uniquely identifies the Address resource that is associated with this AuthorizationDocument. | [optional] 
**CcEmails** | **[]string** | Email recipients who will be informed when an Authorization Document has been sent and signed | [optional] 
**ContactPhoneNumber** | **string** | The contact phone number of the person authorized to sign the Authorization Document. | [optional] 
**ContactTitle** | **string** | The title of the person authorized to sign the Authorization Document for this phone number. | [optional] 
**Email** | **string** | Email that this AuthorizationDocument will be sent to for signing. | [optional] 
**HostedNumberOrderSids** | **[]string** | A list of HostedNumberOrder sids that this AuthorizationDocument will authorize for hosting phone number capabilities on Twilio&#39;s platform. | [optional] 
**Status** | **string** | Status of an instance resource. It can hold one of the values: 1. opened 2. signing, 3. signed LOA, 4. canceled, 5. failed. See the section entitled [Status Values](https://www.twilio.com/docs/api/phone-numbers/hosted-number-authorization-documents#status-values) for more information on each of these statuses. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


