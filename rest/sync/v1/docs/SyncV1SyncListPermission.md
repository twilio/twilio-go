# SyncV1SyncListPermission

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Sync List Permission resource. |
**ServiceSid** | Pointer to **string** | The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) the resource is associated with. |
**ListSid** | Pointer to **string** | The SID of the Sync List to which the Permission applies. |
**Identity** | Pointer to **string** | The application-defined string that uniquely identifies the resource's User within the Service to an FPA token. |
**Read** | Pointer to **bool** | Whether the identity can read the Sync List and its Items. |
**Write** | Pointer to **bool** | Whether the identity can create, update, and delete Items in the Sync List. |
**Manage** | Pointer to **bool** | Whether the identity can delete the Sync List. |
**Url** | Pointer to **string** | The absolute URL of the Sync List Permission resource. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


