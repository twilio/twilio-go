twilio-go changelog
====================
[2021-03-25] Version 0.4.0
--------------------------
**Api**
- Added optional parameter `CallToken` for create calls api
- Add optional property `time_limit` in the call create request.

**Bulkexports**
- adding two new fields with job api queue_position and estimated_completion_time

**Events**
- Add new endpoints to manage subscribed_events in subscriptions

**Numbers**
- Remove feature flags for RegulatoryCompliance endpoints

**Supersim**
- Add SmsCommands resource
- Add fields `SmsCommandsUrl`, `SmsCommandsMethod` and `SmsCommandsEnabled` to a Fleet resource

**Taskrouter**
- Add `If-Match` Header based on ETag for Task Update
- Add `ETag` as Response Headers to Tasks and Reservations

**Video**
- Recording rule beta flag **(breaking change)**
- [Rooms] Add RecordingRules param to Rooms


[2021-03-16] Version 0.3.0
--------------------------
**Library - Docs**
- [PR #54](https://github.com/twilio/twilio-go/pull/54): add property descriptions to docs. Thanks to [@JenniferMah](https://github.com/JenniferMah)!
- [PR #51](https://github.com/twilio/twilio-go/pull/51): fix model reference in docs. Thanks to [@shwetha-manvinkurke](https://github.com/shwetha-manvinkurke)!
- [PR #48](https://github.com/twilio/twilio-go/pull/48): adding docs for enums and getting rid of 'description' column in model docs. Thanks to [@shwetha-manvinkurke](https://github.com/shwetha-manvinkurke)!
- [PR #50](https://github.com/twilio/twilio-go/pull/50): remove optional note for nullable properties. Thanks to [@eshanholtz](https://github.com/eshanholtz)!

**Library - Feature**
- [PR #53](https://github.com/twilio/twilio-go/pull/53): regenerating with openapi-generator 5.0.1. Thanks to [@shwetha-manvinkurke](https://github.com/shwetha-manvinkurke)!

**Library - Chore**
- [PR #52](https://github.com/twilio/twilio-go/pull/52): add go linting as a pre commit hook. Thanks to [@shwetha-manvinkurke](https://github.com/shwetha-manvinkurke)!

**Events**
- Set maturity to beta

**Messaging**
- Adjust A2P brand registration status enum **(breaking change)**

**Studio**
- Remove internal safeguards for Studio V2 API usage now that it's GA

**Verify**
- Add support for creating and verifying totp factors. Support for totp factors is behind the `api.verify.totp` beta feature.


[2021-02-25] Version 0.2.0
--------------------------
**Library - Fix**
- [PR #49](https://github.com/twilio/twilio-go/pull/49): re-add enum prefixes. Thanks to [@eshanholtz](https://github.com/eshanholtz)!

**Events**
- Update description of types in the create sink resource

**Messaging**
- Add WA template header and footer
- Remove A2P campaign and use cases API **(breaking change)**
- Add number_registration_status field to read and fetch campaign responses

**Trusthub**
- Make all resources public

**Verify**
- Verify List Attempts API endpoints added.


[2021-02-17] Version 0.1.1
--------------------------


[2021-02-11] Version 0.1.0
--------------------------

