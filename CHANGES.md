twilio-go changelog
====================
[2021-06-30] Version 0.11.0
---------------------------
**Library - Fix**
- [PR #91](https://github.com/twilio/twilio-go/pull/91): revert client submodule creation. Thanks to [@childish-sambino](https://github.com/childish-sambino)!

**Library - Feature**
- [PR #90](https://github.com/twilio/twilio-go/pull/90): moving client to submodule. Thanks to [@childish-sambino](https://github.com/childish-sambino)!

**Library - Chore**
- [PR #88](https://github.com/twilio/twilio-go/pull/88): split resources into separate files. Thanks to [@childish-sambino](https://github.com/childish-sambino)!
- [PR #87](https://github.com/twilio/twilio-go/pull/87): upgrade openapi-generator to version 5.1.1. Thanks to [@childish-sambino](https://github.com/childish-sambino)!
- [PR #86](https://github.com/twilio/twilio-go/pull/86): use 'int' for integer types. Thanks to [@shwetha-manvinkurke](https://github.com/shwetha-manvinkurke)!

**Library - Docs**
- [PR #89](https://github.com/twilio/twilio-go/pull/89): particulate -> participate, and minor formatting. Thanks to [@stern-shawn](https://github.com/stern-shawn)!
- [PR #85](https://github.com/twilio/twilio-go/pull/85): remove path params docs for operations without them. Thanks to [@childish-sambino](https://github.com/childish-sambino)!

**Conversations**
- Read-only Conversation Email Binding property `binding`

**Supersim**
- Add Billing Period resource for the Super Sim Pilot
- Add List endpoint to Billing Period resource for Super Sim Pilot
- Add Fetch endpoint to Billing Period resource for Super Sim Pilot

**Taskrouter**
- Update `transcribe` & `transcription_configuration` form params in Reservation update endpoint to have private visibility **(breaking change)**
- Add `transcribe` & `transcription_configuration` form params to Reservation update endpoint


[2021-06-16] Version 0.10.0
---------------------------
**Library - Fix**
- [PR #82](https://github.com/twilio/twilio-go/pull/82): array type template. Thanks to [@shamigor](https://github.com/shamigor)!

**Api**
- Update `status` enum for Messages to include 'canceled'
- Update `update_status` enum for Messages to include 'canceled'

**Trusthub**
- Corrected the sid for policy sid in customer_profile_evaluation.json and trust_product_evaluation.json **(breaking change)**


[2021-06-02] Version 0.9.0
--------------------------
**Library - Docs**
- [PR #81](https://github.com/twilio/twilio-go/pull/81): add standalone usage example. Thanks to [@eshanholtz](https://github.com/eshanholtz)!

**Library - Feature**
- [PR #80](https://github.com/twilio/twilio-go/pull/80): add RequestHandler for custom client support. Thanks to [@eshanholtz](https://github.com/eshanholtz)!

**Events**
- join Sinks and Subscriptions service

**Verify**
- Improved the documentation of `challenge` adding the maximum and minimum expected lengths of some fields.
- Improve documentation regarding `notification` by updating the documentation of the field `ttl`.


[2021-05-19] Version 0.8.0
--------------------------
**Library - Chore**
- [PR #79](https://github.com/twilio/twilio-go/pull/79): rename Ip_MessagingVx to IpMessagingVx. Thanks to [@shwetha-manvinkurke](https://github.com/shwetha-manvinkurke)! **(breaking change)**
- [PR #77](https://github.com/twilio/twilio-go/pull/77): remove the unused client pointers. Thanks to [@childish-sambino](https://github.com/childish-sambino)!
- [PR #76](https://github.com/twilio/twilio-go/pull/76): sync the v2010 API with latest generator changes. Thanks to [@childish-sambino](https://github.com/childish-sambino)!
- [PR #75](https://github.com/twilio/twilio-go/pull/75): install pre-commit hooks on install. Thanks to [@childish-sambino](https://github.com/childish-sambino)!
- [PR #73](https://github.com/twilio/twilio-go/pull/73): move contents of 'twilio' folder to the top-level. Thanks to [@childish-sambino](https://github.com/childish-sambino)! **(breaking change)**

**Library - Docs**
- [PR #78](https://github.com/twilio/twilio-go/pull/78): update formatting for godocs. Thanks to [@thinkingserious](https://github.com/thinkingserious)!

**Library - Feature**
- [PR #74](https://github.com/twilio/twilio-go/pull/74): add param setters. Thanks to [@childish-sambino](https://github.com/childish-sambino)!
- [PR #71](https://github.com/twilio/twilio-go/pull/71): add subaccount support for v2010 apis. Thanks to [@shwetha-manvinkurke](https://github.com/shwetha-manvinkurke)! **(breaking change)**

**Events**
- add query param to return types filtered by Schema Id
- Add query param to return sinks filtered by status
- Add query param to return sinks used/not used by a subscription

**Messaging**
- Add fetch and delete instance endpoints to us_app_to_person api **(breaking change)**
- Remove delete list endpoint from us_app_to_person api **(breaking change)**
- Update read list endpoint to return a list of us_app_to_person compliance objects **(breaking change)**
- Add `sid` field to Preregistered US App To Person response

**Supersim**
- Mark `unique_name` in Sim, Fleet, NAP resources as not PII

**Video**
- [Composer] GA maturity level


[2021-05-05] Version 0.7.0
--------------------------
**Library - Chore**
- [PR #70](https://github.com/twilio/twilio-go/pull/70): add check in BuildHost function. Thanks to [@JenniferMah](https://github.com/JenniferMah)!
- [PR #69](https://github.com/twilio/twilio-go/pull/69): update docs links. Thanks to [@thinkingserious](https://github.com/thinkingserious)!
- [PR #66](https://github.com/twilio/twilio-go/pull/66): remove redundant env var tests. Thanks to [@eshanholtz](https://github.com/eshanholtz)!

**Library - Docs**
- [PR #68](https://github.com/twilio/twilio-go/pull/68): update readme to include non v2010 examples. Thanks to [@shwetha-manvinkurke](https://github.com/shwetha-manvinkurke)!

**Api**
- Corrected the data types for feedback summary fields **(breaking change)**
- Update the conference participant create `from` and `to` param to be endpoint type for supporting client identifier and sip address

**Bulkexports**
- promoting API maturity to GA

**Events**
- Add endpoint to update description in sink
- Remove beta-feature account flag

**Messaging**
- Update `status` field in us_app_to_person api to `campaign_status` **(breaking change)**

**Verify**
- Improve documentation regarding `push` factor and include extra information about `totp` factor.


[2021-04-21] Version 0.6.0
--------------------------
**Library - Feature**
- [PR #62](https://github.com/twilio/twilio-go/pull/62): add support for region and edge values in url. Thanks to [@eshanholtz](https://github.com/eshanholtz)!

**Library - Chore**
- [PR #65](https://github.com/twilio/twilio-go/pull/65): update github action linter. Thanks to [@eshanholtz](https://github.com/eshanholtz)!
- [PR #60](https://github.com/twilio/twilio-go/pull/60): add user-agent header on all requests. Thanks to [@shwetha-manvinkurke](https://github.com/shwetha-manvinkurke)!

**Library - Fix**
- [PR #63](https://github.com/twilio/twilio-go/pull/63): parameter names in the request. Thanks to [@shwetha-manvinkurke](https://github.com/shwetha-manvinkurke)!
- [PR #59](https://github.com/twilio/twilio-go/pull/59): custom headers, cleanup docs, regenerate with the latest specs. Thanks to [@shwetha-manvinkurke](https://github.com/shwetha-manvinkurke)!

**Library - Docs**
- [PR #61](https://github.com/twilio/twilio-go/pull/61): Update readme for launch. Thanks to [@garethpaul](https://github.com/garethpaul)!

**Api**
- Revert Update the conference participant create `from` and `to` param to be endpoint type for supporting client identifier and sip address
- Update the conference participant create `from` and `to` param to be endpoint type for supporting client identifier and sip address

**Bulkexports**
- moving enum to doc root for auto generating documentation
- adding status enum and default output properties

**Events**
- Change schema_versions prop and key to versions **(breaking change)**

**Messaging**
- Add `use_inbound_webhook_on_number` field in Service API for fetch, create, update, read

**Taskrouter**
- Add `If-Match` Header based on ETag for Task Delete

**Verify**
- Add `AuthPayload` parameter to support verifying a `Challenge` upon creation. This is only supported for `totp` factors.
- Add support to resend the notifications of a `Challenge`. This is only supported for `push` factors.


[2021-04-07] Version 0.5.0
--------------------------
**Library - Docs**
- [PR #58](https://github.com/twilio/twilio-go/pull/58): update readme and documentation link. Thanks to [@JenniferMah](https://github.com/JenniferMah)!

**Library - Chore**
- [PR #56](https://github.com/twilio/twilio-go/pull/56): refactor client and regenerate with updated mustache. Thanks to [@shwetha-manvinkurke](https://github.com/shwetha-manvinkurke)!
- [PR #55](https://github.com/twilio/twilio-go/pull/55): regenerate after splitting model into separate components. Thanks to [@JenniferMah](https://github.com/JenniferMah)!

**Api**
- Added `announcement` event to conference status callback events
- Removed optional property `time_limit` in the call create request. **(breaking change)**

**Messaging**
- Add rate_limits field to Messaging Services US App To Person API
- Add usecase field in Service API for fetch, create, update, read
- Add us app to person api and us app to person usecase api as dependents in service
- Add us_app_to_person_registered field in service api for fetch, read, create, update
- Add us app to person api
- Add us app to person usecase api
- Add A2P external campaign api
- Add Usecases API

**Supersim**
- Add Create endpoint to Sims resource

**Verify**
- The `Binding` field is now returned when creating a `Factor`. This value won't be returned for other endpoints.

**Video**
- [Rooms] max_concurrent_published_tracks has got GA maturity


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

