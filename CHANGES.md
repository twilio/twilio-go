twilio-go changelog
====================
[2022-06-29] Version 1.0.0-rc.2
-------------------------------
**Library - Test**
- [PR #169](https://github.com/twilio/twilio-go/pull/169): edit regex for rc naming for useragent tests. Thanks to [@JenniferMah](https://github.com/JenniferMah)!

**Api**
- Added `amazon-polly` to `usage_record` API.

**Insights**
- Added `annotation` field in call summary
- Added new endpoint to fetch/create/update Call Annotations

**Verify**
- Remove `api.verify.totp` beta flag and set maturity to `beta` for Verify TOTP properties and parameters. **(breaking change)**
- Changed summary param `verify_service_sid` to `service_sid` to be consistent with list attempts API **(breaking change)**


[2022-06-15] Version 1.0.0-rc.1
-------------------------------
**Lookups**
- Adding support for Lookup V2 API

**Studio**
- Corrected PII labels to be 30 days and added context to be PII


[2022-06-15] Version 1.0.0-rc.0
---------------------------
- Release Candidate prep


[2022-05-18] Version 0.26.0
---------------------------
**Note:** This release contains breaking changes, check our [upgrade guide](./UPGRADE.md#2022-05-18-025x-to-026x) for detailed migration notes.

**Library - Chore**
- [PR #164](https://github.com/twilio/twilio-go/pull/164): rename ApiV2010 to Api. Thanks to [@JenniferMah](https://github.com/JenniferMah)! **(breaking change)**

**Api**
- Add property `media_url` to the recording resources

**Verify**
- Include `silent` as a channel type in the verifications API.


[2022-05-04] Version 0.25.0
---------------------------
**Library - Feature**
- [PR #162](https://github.com/twilio/twilio-go/pull/162): twilio-go stream functions also return error channel. Thanks to [@childish-sambino](https://github.com/childish-sambino)!
- [PR #158](https://github.com/twilio/twilio-go/pull/158): add unmarshaling helper for float32. Thanks to [@childish-sambino](https://github.com/childish-sambino)!

**Library - Chore**
- [PR #159](https://github.com/twilio/twilio-go/pull/159): [DI-1566] modify ua string. Thanks to [@claudiachua](https://github.com/claudiachua)!

**Conversations**
- Expose query parameter `type` in list operation on Address Configurations resource

**Supersim**
- Add `data_total_billed` and `billed_units` fields to Super SIM UsageRecords API response.
- Change ESimProfiles `Eid` parameter to optional to enable Activation Code download method support **(breaking change)**

**Verify**
- Deprecate `push.include_date` parameter in create and update service.


[2022-04-20] Version 0.24.1
---------------------------
**Library - Test**
- [PR #156](https://github.com/twilio/twilio-go/pull/156): fix type on broken cluster test. Thanks to [@childish-sambino](https://github.com/childish-sambino)!

**Library - Fix**
- [PR #155](https://github.com/twilio/twilio-go/pull/155): switch api-def object types to open-api any types. Thanks to [@childish-sambino](https://github.com/childish-sambino)!
- [PR #154](https://github.com/twilio/twilio-go/pull/154): support operations with no 2XX responses. Thanks to [@childish-sambino](https://github.com/childish-sambino)!


[2022-04-06] Version 0.24.0
---------------------------
**Library - Fix**
- [PR #153](https://github.com/twilio/twilio-go/pull/153): use go install instead of go get. Thanks to [@beebzz](https://github.com/beebzz)!

**Library - Feature**
- [PR #152](https://github.com/twilio/twilio-go/pull/152): update cluster tests authentication. Thanks to [@JenniferMah](https://github.com/JenniferMah)!
- [PR #151](https://github.com/twilio/twilio-go/pull/151): add support for Go 1.18. Thanks to [@beebzz](https://github.com/beebzz)!

**Api**
- Updated `provider_sid` visibility to private

**Verify**
- Verify List Attempts API summary endpoint added.
- Update PII documentation for `AccessTokens` `factor_friendly_name` property.

**Voice**
- make annotation parameter from /Calls API private


[2022-03-23] Version 0.23.0
---------------------------
**Library - Docs**
- [PR #150](https://github.com/twilio/twilio-go/pull/150): Renaming RestClientParams to ClientParams and updating docs examples. Thanks to [@rakatyal](https://github.com/rakatyal)!
- [PR #147](https://github.com/twilio/twilio-go/pull/147): fix the 'List*' operation return type documentation. Thanks to [@childish-sambino](https://github.com/childish-sambino)!

**Library - Fix**
- [PR #143](https://github.com/twilio/twilio-go/pull/143): handle buildUrl errors gracefully. Thanks to [@kpritam](https://github.com/kpritam)!
- [PR #148](https://github.com/twilio/twilio-go/pull/148): refactor limit enforcement and fetching next page. Thanks to [@childish-sambino](https://github.com/childish-sambino)!

**Api**
- Change `stream` url parameter to non optional
- Add `verify-totp` and `verify-whatsapp-conversations-business-initiated` categories to `usage_record` API

**Chat**
- Added v3 Channel update endpoint to support Public to Private channel migration

**Flex**
- Private Beta release of the Interactions API to support the upcoming release of Flex Conversations at the end of Q1 2022.
- Adding `channel_configs` object to Flex Configuration

**Media**
- Add max_duration param to PlayerStreamer

**Supersim**
- Remove Commands resource, use SmsCommands resource instead **(breaking change)**

**Taskrouter**
- Add limits to `split_by_wait_time` for Cumulative Statistics Endpoint

**Video**
- Change recording `status_callback_method` type from `enum` to `http_method` **(breaking change)**
- Add `status_callback` and `status_callback_method` to composition
- Add `status_callback` and `status_callback_method` to recording


[2022-03-09] Version 0.22.2
---------------------------
**Library - Chore**
- [PR #142](https://github.com/twilio/twilio-go/pull/142): push Datadog Release Metric upon deploy success. Thanks to [@eshanholtz](https://github.com/eshanholtz)!

**Api**
- Add optional boolean include_soft_deleted parameter to retrieve soft deleted recordings

**Chat**
- Add `X-Twilio-Wehook-Enabled` header to `delete` method in UserChannel resource

**Numbers**
- Expose `failure_reason` in the Supporting Documents resources

**Verify**
- Add optional `metadata` parameter to "verify challenge" endpoint, so the SDK/App can attach relevant information from the device when responding to challenges.
- remove beta feature flag to list atempt api operations.
- Add `ttl` and `date_created` properties to `AccessTokens`.


[2022-02-23] Version 0.22.1
---------------------------
**Api**
- Add `uri` to `stream` resource
- Add A2P Registration Fee category (`a2p-registration-fee`) to usage records

**Verify**
- Remove outdated documentation commentary to contact sales. Product is already in public beta.


[2022-02-17] Version 0.22.0
---------------------------
**Api**
- Detected a bug and removed optional boolean include_soft_deleted parameter to retrieve soft deleted recordings. **(breaking change)**
- Add optional boolean include_soft_deleted parameter to retrieve soft deleted recordings.

**Numbers**
- Unrevert valid_until and sort filter params added to List Bundles resource
- Revert valid_until and sort filter params added to List Bundles resource
- Update sorting params added to List Bundles resource in the previous release

**Preview**
- Moved `web_channels` from preview to beta under `flex-api` **(breaking change)**

**Taskrouter**
- Add `ETag` as Response Header to List of Task, Reservation & Worker

**Verify**
- Add optional `metadata` to factors.


[2022-02-09] Version 0.21.0
---------------------------
**Library - Chore**
- [PR #138](https://github.com/twilio/twilio-go/pull/138): upgrade supported language versions. Thanks to [@childish-sambino](https://github.com/childish-sambino)!

**Api**
- Add `stream` resource

**Conversations**
- Fixed DELETE request to accept "sid_like" params in Address Configuration resources **(breaking change)**
- Expose Address Configuration resource for `sms` and `whatsapp`

**Fax**
- Removed deprecated Programmable Fax Create and Update methods **(breaking change)**

**Insights**
- Rename `call_state` to `call_status` and remove `whisper` in conference participant summary **(breaking change)**

**Numbers**
- Expose valid_until filters as part of provisionally-approved compliance feature on the List Bundles resource

**Supersim**
- Fix typo in Fleet resource docs
- Updated documentation for the Fleet resource indicating that fields related to commands have been deprecated and to use sms_command fields instead.
- Add support for setting and reading `ip_commands_url` and `ip_commands_method` on Fleets resource for helper libraries
- Changed `sim` property in requests to create an SMS Command made to the /SmsCommands to accept SIM UniqueNames in addition to SIDs

**Verify**
- Update list attempts API to include new filters and response fields.


[2022-01-26] Version 0.20.1
---------------------------
**Insights**
- Added new endpoint to fetch Conference Participant Summary
- Added new endpoint to fetch Conference Summary

**Messaging**
- Add government_entity parameter to brand apis

**Verify**
- Add Access Token fetch endpoint to retrieve a previously created token.
- Add Access Token payload to the Access Token creation endpoint, including a unique Sid, so it's addressable while it's TTL is valid.


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
