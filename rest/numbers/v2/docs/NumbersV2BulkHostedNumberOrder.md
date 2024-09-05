# NumbersV2BulkHostedNumberOrder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BulkHostingSid** | Pointer to **string** | A 34 character string that uniquely identifies this BulkHostedNumberOrder. |
**RequestStatus** | Pointer to [**string**](BulkHostedNumberOrderEnumRequestStatus.md) |  |
**FriendlyName** | Pointer to **string** | A 128 character string that is a human-readable text that describes this resource. |
**NotificationEmail** | Pointer to **string** | Email address used for send notifications about this Bulk hosted number request. |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The date this resource was created, given as [GMT RFC 2822](http://www.ietf.org/rfc/rfc2822.txt) format. |
**DateCompleted** | Pointer to [**time.Time**](time.Time.md) | The date that this resource was completed, given as [GMT RFC 2822](http://www.ietf.org/rfc/rfc2822.txt) format. |
**Url** | Pointer to **string** | The URL of this BulkHostedNumberOrder resource. |
**TotalCount** | **int** | The total count of phone numbers in this Bulk hosting request. |[optional] [default to 0]
**Results** | Pointer to **[]interface{}** | Contains a list of all the individual hosting orders and their information, for this Bulk request. Each result object is grouped by its order status. To see a complete list of order status, please check 'https://www.twilio.com/docs/phone-numbers/hosted-numbers/hosted-numbers-api/hosted-number-order-resource#status-values'. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


