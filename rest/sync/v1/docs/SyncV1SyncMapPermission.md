# SyncV1SyncMapPermission

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Sync Map Permission resource. |
**ServiceSid** | Pointer to **string** | The SID of the [Sync Service](https://www.twilio.com/docs/sync/api/service) the resource is associated with. |
**MapSid** | Pointer to **string** | The SID of the Sync Map to which the Permission applies. |
**Identity** | Pointer to **string** | The application-defined string that uniquely identifies the resource's User within the Service to an FPA token. |
**Read** | Pointer to **bool** | Whether the identity can read the Sync Map and its Items. |
**Write** | Pointer to **bool** | Whether the identity can create, update, and delete Items in the Sync Map. |
**Manage** | Pointer to **bool** | Whether the identity can delete the Sync Map. |
**Url** | Pointer to **string** | The absolute URL of the Sync Map Permission resource. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


