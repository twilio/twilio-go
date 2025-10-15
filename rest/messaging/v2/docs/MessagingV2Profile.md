# MessagingV2Profile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the sender. Required for WhatsApp senders and must follow [Meta's display name guidelines](https://www.facebook.com/business/help/757569725593362). |
**About** | Pointer to **string** | The profile about text for the sender. |
**Address** | Pointer to **string** | The address of the sender. |
**Description** | Pointer to **string** | The description of the sender. |
**LogoUrl** | Pointer to **string** | The logo URL of the sender. |
**BannerUrl** | Pointer to **string** | The banner URL of the sender. |
**PrivacyUrl** | Pointer to **string** | The privacy URL of the sender. Must be a publicly accessible HTTP or HTTPS URI associated with the sender.  |
**TermsOfServiceUrl** | Pointer to **string** | The terms of service URL of the sender. |
**AccentColor** | Pointer to **string** | The color theme of the sender. Must be in hex format and have at least a 4:5:1 contrast ratio against white. |
**Vertical** | Pointer to **string** | The vertical of the sender. Allowed values are: - `Automotive` - `Beauty, Spa and Salon` - `Clothing and Apparel` - `Education` - `Entertainment` - `Event Planning and Service` - `Finance and Banking` - `Food and Grocery` - `Public Service` - `Hotel and Lodging` - `Medical and Health` - `Non-profit` - `Professional Services` - `Shopping and Retail` - `Travel and Transportation` - `Restaurant` - `Other`  |
**Websites** | Pointer to **interface{}** | The websites of the sender. |
**Emails** | Pointer to **interface{}** | The emails of the sender. |
**PhoneNumbers** | Pointer to **interface{}** | The phone numbers of the sender. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


