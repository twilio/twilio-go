# SupersimV1EsimProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | Pointer to **string** | The unique string that we created to identify the eSIM Profile resource. |
**AccountSid** | Pointer to **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) to which the eSIM Profile resource belongs. |
**Iccid** | Pointer to **string** | The [ICCID](https://en.wikipedia.org/wiki/Subscriber_identity_module#ICCID) associated with the Sim resource. |
**SimSid** | Pointer to **string** | The SID of the [Sim](https://www.twilio.com/docs/wireless/api/sim-resource) resource that this eSIM Profile controls. |
**Status** | Pointer to [**string**](EsimProfileEnumStatus.md) |  |
**Eid** | Pointer to **string** | Identifier of the eUICC that can claim the eSIM Profile. |
**SmdpPlusAddress** | Pointer to **string** | Address of the SM-DP+ server from which the Profile will be downloaded. The URL will appear once the eSIM Profile reaches the status `available`. |
**ErrorCode** | Pointer to **string** | Code indicating the failure if the download of the SIM Profile failed and the eSIM Profile is in `failed` state. |
**ErrorMessage** | Pointer to **string** | Error message describing the failure if the download of the SIM Profile failed and the eSIM Profile is in `failed` state. |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The date and time in GMT when the resource was created specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The date and time in GMT when the resource was last updated specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. |
**Url** | Pointer to **string** | The absolute URL of the eSIM Profile resource. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


