# EdgeZones

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Zones** | **[]string** | Priority-ordered list of zone IDs. Twilio routes to the first available zone in order. Individual PoP identifiers (e.g. `IAD-1`) are not valid — use zone IDs only (see `GET /EdgeZones` for available values).  |
**AllowZoneFallback** | **bool** | Controls delivery when none of the configured zones have available PoPs. - `false` (default): delivery fails rather than routing through an unconfigured zone. - `true`: Twilio falls back to any available zone outside the configured list.  |[optional] [default to false]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


