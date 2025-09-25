# NumbersV1PortingLosingCarrierInformation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerName** | **string** | Customer name as it is registered with the losing carrier. This can be an individual or a business name depending on the customer type selected. |
**AccountNumber** | **string** | The account number of the customer for the losing carrier. Only require for mobile phone numbers. |[optional] 
**AccountTelephoneNumber** | **string** | The account phone number of the customer for the losing carrier. |[optional] 
**AddressSid** | **string** | If you already have an Address SID that represents the address needed for the LOA, you can provide an Address SID instead of providing the address object in the request body. This will copy the address into the port in request. If changes are made to the Address SID after port in request creation, those changes will not be reflected in the port in request. |[optional] 
**Address** | [**NumbersV1PortingAddress**](NumbersV1PortingAddress.md) |  |[optional] 
**AuthorizedRepresentative** | **string** | The first and last name of the person listed with the losing carrier who is authorized to make changes on the account. |
**AuthorizedRepresentativeEmail** | **string** | Email address of the person (owner of the number) who will sign the letter of authorization for the port in request. This email address should belong to the person named in as the authorized representative. |
**CustomerType** | **string** | The type of customer account in the losing carrier. This should either be: 'Individual' or 'Business'. |[optional] 
**AuthorizedRepresentativeKatakana** | **string** |  |[optional] 
**SubMunicipality** | **string** |  |[optional] 
**Building** | **string** |  |[optional] 
**KatakanaName** | **string** |  |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


