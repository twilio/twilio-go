# ContextV1Record

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | Pointer to **string** | A 34 character string that uniquely identifies the record for an object. |
**AccountSid** | Pointer to **string** | The unique id of the [Account](https://www.twilio.com/docs/api/rest/account) that this Account belongs to. |
**SourceSchemaId** | Pointer to **string** | A 34 character string that uniquely identifies the source schema name of the record. |
**SourceType** | Pointer to [**string**](RecordEnumSourceType.md) |  |
**SourceName** | Pointer to **string** | Human-readable name of the source of the record. |
**RecordData** | Pointer to **interface{}** | A key-value pair that contains the description of the given record instance. |
**CreatedAt** | Pointer to [**time.Time**](time.Time.md) | The date that this resource was created, given as GMT in [ISO 8601](http://www.iso.org/iso/home/standards/iso8601.htm) format. |
**UpdatedAt** | Pointer to [**time.Time**](time.Time.md) | The date that this resource was last updated, given as GMT in [ISO 8601](http://www.iso.org/iso/home/standards/iso8601.htm) format. |
**Url** | Pointer to **string** | The absolute URL of the Record resource. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


