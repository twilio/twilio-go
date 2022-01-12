twilio-go changelog
====================
[2022-01-12] Version 0.20.0
---------------------------
**Library - Chore**
- [PR #135](https://github.com/twilio/twilio-go/pull/135): add sonarcloud coverage analysis. Thanks to [@shwetha-manvinkurke](https://github.com/shwetha-manvinkurke)!
- [PR #132](https://github.com/twilio/twilio-go/pull/132): remove githook dependency from make install. Thanks to [@JenniferMah](https://github.com/JenniferMah)!

**Library - Feature**
- [PR #134](https://github.com/twilio/twilio-go/pull/134): add GitHub release step during deploy. Thanks to [@childish-sambino](https://github.com/childish-sambino)!

**Api**
- Make fixed time scheduling parameters public **(breaking change)**

**Messaging**
- Add update brand registration API

**Numbers**
- Add API endpoint for List Bundle Copies resource

**Video**
- Enable external storage for all customers


[2021-12-15] Version 0.19.0
---------------------------
**Library - Feature**
- [PR #129](https://github.com/twilio/twilio-go/pull/129): run tests before deploying. Thanks to [@childish-sambino](https://github.com/childish-sambino)!

**Library - Docs**
- [PR #131](https://github.com/twilio/twilio-go/pull/131): fix whitespace issue in handling exceptions example. Thanks to [@EmFord](https://github.com/EmFord)!

**Api**
- Add optional boolean send_as_mms parameter to the create action of Message resource **(breaking change)**
- Change team ownership for `call` delete

**Conversations**
- Change wording for `Service Webhook Configuration` resource fields

**Insights**
- Added new APIs for updating and getting voice insights flags by accountSid.

**Media**
- Add max_duration param to MediaProcessor

**Video**
- Add `EmptyRoomTimeout` and `UnusedRoomTimeout` properties to a room; add corresponding parameters to room creation

**Voice**
- Add endpoint to delete archived Calls


[2021-12-01] Version 0.18.2
---------------------------
**Conversations**
- Add `Service Webhook Configuration` resource

**Flex**
- Adding `flex_insights_drilldown` and `flex_url` objects to Flex Configuration

**Messaging**
- Update us_app_to_person endpoints to remove beta feature flag based access

**Supersim**
- Add IP Commands resource

**Verify**
- Add optional `factor_friendly_name` parameter to the create access token endpoint.

**Video**
- Add maxParticipantDuration param to Rooms


[2021-11-17] Version 0.18.1
---------------------------
**Library - Chore**
- [PR #127](https://github.com/twilio/twilio-go/pull/127): consolidate workflows. Thanks to [@eshanholtz](https://github.com/eshanholtz)!
- [PR #126](https://github.com/twilio/twilio-go/pull/126): add event type cluster test. Thanks to [@JenniferMah](https://github.com/JenniferMah)!

**Library - Fix**
- [PR #124](https://github.com/twilio/twilio-go/pull/124): git log retrieval issues. Thanks to [@shwetha-manvinkurke](https://github.com/shwetha-manvinkurke)!

**Frontline**
- Added `is_available` to User's resource

**Messaging**
- Added GET vetting API

**Verify**
- Add `WHATSAPP` to the attempts API.
- Allow to update `config.notification_platform` from `none` to `apn` or `fcm` and viceversa for Verify Push
- Add `none` as a valid `config.notification_platform` value for Verify Push


[2021-11-03] Version 0.18.0
---------------------------
**Library - Fix**
- [PR #125](https://github.com/twilio/twilio-go/pull/125): remove cluster test. Thanks to [@JenniferMah](https://github.com/JenniferMah)!

**Library - Chore**
- [PR #123](https://github.com/twilio/twilio-go/pull/123): migrate to github actions from travis ci. Thanks to [@shwetha-manvinkurke](https://github.com/shwetha-manvinkurke)!

**Api**
- Updated `media_url` property to be treated as PII

**Messaging**
- Added a new enum for brand registration status named DELETED **(breaking change)**
- Add a new K12_EDUCATION use case in us_app_to_person_usecase api transaction
- Added a new enum for brand registration status named IN_REVIEW

**Serverless**
- Add node14 as a valid Build runtime

**Verify**
- Fix typos in Verify Push Factor documentation for the `config.notification_token` parameter.
- Added `TemplateCustomSubstitutions` on verification creation
- Make `TemplateSid` parameter public for Verification resource and `DefaultTemplateSid` parameter public for Service resource. **(breaking change)**


[2021-10-18] Version 0.17.0
---------------------------
**Library - Feature**
- [PR #122](https://github.com/twilio/twilio-go/pull/122): Add PlaybackGrant. Thanks to [@miguelgrinberg](https://github.com/miguelgrinberg)!

**Api**
- Corrected enum values for `emergency_address_status` values in `/IncomingPhoneNumbers` response. **(breaking change)**
- Clarify `emergency_address_status` values in `/IncomingPhoneNumbers` response.

**Messaging**
- Add PUT and List brand vettings api
- Removes beta feature flag based visibility for us_app_to_person_registered and usecase field.Updates test cases to add POLITICAL usecase. **(breaking change)**
- Add brand_feedback as optional field to BrandRegistrations

**Video**
- Add `AudioOnly` to create room


[2021-10-06] Version 0.16.0
---------------------------
**Library - Fix**
- [PR #121](https://github.com/twilio/twilio-go/pull/121): parameter casing with numbers. Thanks to [@shwetha-manvinkurke](https://github.com/shwetha-manvinkurke)!

**Api**
- Add `emergency_address_status` attribute to `/IncomingPhoneNumbers` response.
- Add `siprec` resource

**Conversations**
- Added attachment parameters in configuration for `NewMessage` type of push notifications

**Flex**
- Adding `flex_insights_hr` object to Flex Configuration

**Numbers**
- Add API endpoint for Bundle ReplaceItems resource
- Add API endpoint for Bundle Copies resource

**Serverless**
- Add domain_base field to Service response

**Taskrouter**
- Add `If-Match` Header based on ETag for Worker Delete **(breaking change)**
- Add `If-Match` Header based on Etag for Reservation Update
- Add `If-Match` Header based on ETag for Worker Update
- Add `If-Match` Header based on ETag for Worker Delete
- Add `ETag` as Response Header to Worker

**Trunking**
- Added `transfer_caller_id` property on Trunks.

**Verify**
- Document new pilot `whatsapp` channel.


[2021-09-22] Version 0.15.0
---------------------------
**Library - Feature**
- [PR #116](https://github.com/twilio/twilio-go/pull/116): add jwt token signing and verification logic. Thanks to [@shwetha-manvinkurke](https://github.com/shwetha-manvinkurke)!

**Library - Test**
- [PR #117](https://github.com/twilio/twilio-go/pull/117): increase Client test coverage. Thanks to [@eshanholtz](https://github.com/eshanholtz)!

**Events**
- Add segment sink

**Messaging**
- Add post_approval_required attribute in GET us_app_to_person_usecase api response
- Add Identity Status, Russell 3000, Tax Exempt Status and Should Skip SecVet fields for Brand Registrations
- Add Should Skip Secondary Vetting optional flag parameter to create Brand API


[2021-09-08] Version 0.14.2
---------------------------
**Library - Fix**
- [PR #115](https://github.com/twilio/twilio-go/pull/115): typo in page function. Thanks to [@shwetha-manvinkurke](https://github.com/shwetha-manvinkurke)!

**Api**
- Revert adding `siprec` resource
- Add `siprec` resource

**Messaging**
- Add 'mock' as an optional field to brand_registration api
- Add 'mock' as an optional field to us_app_to_person api
- Adds more Use Cases in us_app_to_person_usecase api transaction and updates us_app_to_person_usecase docs

**Verify**
- Verify List Templates API endpoint added.


[2021-08-25] Version 0.14.1
---------------------------
**Api**
- Add Programmabled Voice SIP Refer call transfers (`calls-transfers`) to usage records
- Add Flex Voice Usage category (`flex-usage`) to usage records

**Conversations**
- Add `Order` query parameter to Message resource read operation

**Insights**
- Added `partial` to enum processing_state_request
- Added abnormal session filter in Call Summaries

**Messaging**
- Add brand_registration_sid as an optional query param for us_app_to_person_usecase api

**Pricing**
- add trunking_numbers resource (v2)
- add trunking_country resource (v2)

**Verify**
- Changed to private beta the `TemplateSid` optional parameter on Verification creation.
- Added the optional parameter `Order` to the list Challenges endpoint to define the list order.


[2021-08-11] Version 0.14.0
---------------------------
**Library - Fix**
- [PR #110](https://github.com/twilio/twilio-go/pull/110): pagination with next_page_url. Thanks to [@shwetha-manvinkurke](https://github.com/shwetha-manvinkurke)!

**Library - Chore**
- [PR #108](https://github.com/twilio/twilio-go/pull/108): shorten generated model names. Thanks to [@shwetha-manvinkurke](https://github.com/shwetha-manvinkurke)!
- [PR #106](https://github.com/twilio/twilio-go/pull/106): integrate with sonarcloud. Thanks to [@shwetha-manvinkurke](https://github.com/shwetha-manvinkurke)!

**Api**
- Corrected the `price`, `call_sid_to_coach`, and `uri` data types for Conference, Participant, and Recording **(breaking change)**
- Made documentation for property `time_limit` in the call api public. **(breaking change)**
- Added `domain_sid` in sip_credential_list_mapping and sip_ip_access_control_list_mapping APIs **(breaking change)**

**Insights**
- Added new endpoint to fetch Call Summaries

**Messaging**
- Add brand_type field to a2p brand_registration api
- Revert brand registration api update to add brand_type field
- Add brand_type field to a2p brand_registration api

**Taskrouter**
- Add `X-Rate-Limit-Limit`, `X-Rate-Limit-Remaining`, and `X-Rate-Limit-Config` as Response Headers to all TaskRouter endpoints

**Verify**
- Add `TemplateSid` optional parameter on Verification creation.
- Include `whatsapp` as a channel type in the verifications API.


[2021-07-28] Version 0.13.0
---------------------------
**Library - Feature**
- [PR #105](https://github.com/twilio/twilio-go/pull/105): publish go docker image. Thanks to [@shwetha-manvinkurke](https://github.com/shwetha-manvinkurke)!
- [PR #92](https://github.com/twilio/twilio-go/pull/92): add pagination. Thanks to [@shwetha-manvinkurke](https://github.com/shwetha-manvinkurke)! **(breaking change)**
- [PR #100](https://github.com/twilio/twilio-go/pull/100): add cluster testing. Thanks to [@shwetha-manvinkurke](https://github.com/shwetha-manvinkurke)!
- [PR #96](https://github.com/twilio/twilio-go/pull/96): Twilio Signature Validation. Thanks to [@pmcanseco](https://github.com/pmcanseco)!
- [PR #95](https://github.com/twilio/twilio-go/pull/95): support env var loading of API credentials. Thanks to [@eshanholtz](https://github.com/eshanholtz)!

**Library - Test**
- [PR #103](https://github.com/twilio/twilio-go/pull/103): add edge case tests. Thanks to [@eshanholtz](https://github.com/eshanholtz)!

**Library - Fix**
- [PR #102](https://github.com/twilio/twilio-go/pull/102): list of stringified json marshaling. Thanks to [@shwetha-manvinkurke](https://github.com/shwetha-manvinkurke)!
- [PR #99](https://github.com/twilio/twilio-go/pull/99): add nil check. Thanks to [@shwetha-manvinkurke](https://github.com/shwetha-manvinkurke)!

**Library - Chore**
- [PR #101](https://github.com/twilio/twilio-go/pull/101): refactor list params to include 'limit'. Thanks to [@shwetha-manvinkurke](https://github.com/shwetha-manvinkurke)!

**Library - Docs**
- [PR #97](https://github.com/twilio/twilio-go/pull/97): add pagination docs. Thanks to [@eshanholtz](https://github.com/eshanholtz)!

**Conversations**
- Expose ParticipantConversations resource

**Taskrouter**
- Adding `links` to the activity resource

**Verify**
- Added a `Version` to Verify Factors `Webhooks` to add new fields without breaking old Webhooks.


[2021-07-14] Version 0.12.0
---------------------------
**Library - Fix**
- [PR #93](https://github.com/twilio/twilio-go/pull/93): list of stringified json marshaling. Thanks to [@eshanholtz](https://github.com/eshanholtz)!

**Conversations**
- Changed `last_read_message_index` and `unread_messages_count` type in User Conversation's resource **(breaking change)**
- Expose UserConversations resource

**Messaging**
- Add brand_score field to brand registration responses


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

