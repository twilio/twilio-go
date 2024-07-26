# NumbersV1SigningRequestConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LogoSid** | Pointer to **string** | The SID of the document  that includes the logo that will appear in the LOA. To upload documents follow the following guide: https://www.twilio.com/docs/phone-numbers/regulatory/getting-started/create-new-bundle-public-rest-apis#supporting-document-create  |
**FriendlyName** | Pointer to **string** | This is the string that you assigned as a friendly name for describing the creation of the configuration. |
**Product** | Pointer to **string** | The product or service for which is requesting the signature. |
**Country** | Pointer to **string** | The country ISO code to apply the configuration. |
**EmailSubject** | Pointer to **string** | Subject of the email that the end client will receive ex: “Twilio Hosting Request”, maximum length of 255 characters. |
**EmailMessage** | Pointer to **string** | Content of the email that the end client will receive ex: “This is a Hosting request from Twilio, please check the document and sign it”, maximum length of 5,000 characters. |
**UrlRedirection** | Pointer to **string** | Url the end client will be redirected after signing a document. |
**Url** | Pointer to **string** |  |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


