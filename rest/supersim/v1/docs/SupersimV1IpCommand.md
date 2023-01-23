# SupersimV1IpCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | Pointer to **string** | The unique string that we created to identify the IP Command resource. |
**AccountSid** | Pointer to **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the IP Command resource. |
**SimSid** | Pointer to **string** | The SID of the [Super SIM](https://www.twilio.com/docs/iot/supersim/api/sim-resource) that this IP Command was sent to or from. |
**SimIccid** | Pointer to **string** | The [ICCID](https://en.wikipedia.org/wiki/Subscriber_identity_module#ICCID) of the [Super SIM](https://www.twilio.com/docs/iot/supersim/api/sim-resource) that this IP Command was sent to or from. |
**Status** | Pointer to [**string**](IpCommandEnumStatus.md) |  |
**Direction** | Pointer to [**string**](IpCommandEnumDirection.md) |  |
**DeviceIp** | Pointer to **string** | The IP address of the device that the IP Command was sent to or received from. For an IP Command sent to a Super SIM, `device_ip` starts out as `null`, and once the IP Command is “sent”, the `device_ip` will be filled out. An IP Command sent from a Super SIM have its `device_ip` always set. |
**DevicePort** | Pointer to **int** | For an IP Command sent to a Super SIM, it would be the destination port of the IP message. For an IP Command sent from a Super SIM, it would be the source port of the IP message. |
**PayloadType** | Pointer to [**string**](IpCommandEnumPayloadType.md) |  |
**Payload** | Pointer to **string** | The payload that is carried in the IP/UDP message. The payload can be encoded in either text or binary format. For text payload, UTF-8 encoding must be used.  For an IP Command sent to a Super SIM, the payload is appended to the IP/UDP message “as is”. The payload should not exceed 1300 bytes.  For an IP Command sent from a Super SIM, the payload from the received IP/UDP message is extracted and sent in binary encoding. For an IP Command sent from a Super SIM, the payload should not exceed 1300 bytes. If it is larger than 1300 bytes, there might be fragmentation on the upstream and the message may appear truncated. |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The date and time in GMT when the resource was created specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The date and time in GMT when the resource was last updated specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. |
**Url** | Pointer to **string** | The absolute URL of the IP Command resource. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


