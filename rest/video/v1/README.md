# Go API client for openapi

This is the public Twilio REST API.

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project from the OpenAPI specs located at [twilio/twilio-oai](https://github.com/twilio/twilio-oai/tree/main/spec).  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 1.0.0
- Package version: 1.0.0
- Build package: com.twilio.oai.TwilioGoGenerator
For more information, please visit [https://support.twilio.com](https://support.twilio.com)

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import "./openapi"
```

## Documentation for API Endpoints

All URIs are relative to *https://video.twilio.com*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*CompositionHooksApi* | [**CreateCompositionHook**](docs/CompositionHooksApi.md#createcompositionhook) | **Post** /v1/CompositionHooks | 
*CompositionHooksApi* | [**DeleteCompositionHook**](docs/CompositionHooksApi.md#deletecompositionhook) | **Delete** /v1/CompositionHooks/{Sid} | Delete a Recording CompositionHook resource identified by a &#x60;CompositionHook SID&#x60;.
*CompositionHooksApi* | [**FetchCompositionHook**](docs/CompositionHooksApi.md#fetchcompositionhook) | **Get** /v1/CompositionHooks/{Sid} | Returns a single CompositionHook resource identified by a CompositionHook SID.
*CompositionHooksApi* | [**ListCompositionHook**](docs/CompositionHooksApi.md#listcompositionhook) | **Get** /v1/CompositionHooks | List of all Recording CompositionHook resources.
*CompositionHooksApi* | [**UpdateCompositionHook**](docs/CompositionHooksApi.md#updatecompositionhook) | **Post** /v1/CompositionHooks/{Sid} | 
*CompositionSettingsDefaultApi* | [**CreateCompositionSettings**](docs/CompositionSettingsDefaultApi.md#createcompositionsettings) | **Post** /v1/CompositionSettings/Default | 
*CompositionSettingsDefaultApi* | [**FetchCompositionSettings**](docs/CompositionSettingsDefaultApi.md#fetchcompositionsettings) | **Get** /v1/CompositionSettings/Default | 
*CompositionsApi* | [**CreateComposition**](docs/CompositionsApi.md#createcomposition) | **Post** /v1/Compositions | 
*CompositionsApi* | [**DeleteComposition**](docs/CompositionsApi.md#deletecomposition) | **Delete** /v1/Compositions/{Sid} | Delete a Recording Composition resource identified by a Composition SID.
*CompositionsApi* | [**FetchComposition**](docs/CompositionsApi.md#fetchcomposition) | **Get** /v1/Compositions/{Sid} | Returns a single Composition resource identified by a Composition SID.
*CompositionsApi* | [**ListComposition**](docs/CompositionsApi.md#listcomposition) | **Get** /v1/Compositions | List of all Recording compositions.
*RecordingSettingsDefaultApi* | [**CreateRecordingSettings**](docs/RecordingSettingsDefaultApi.md#createrecordingsettings) | **Post** /v1/RecordingSettings/Default | 
*RecordingSettingsDefaultApi* | [**FetchRecordingSettings**](docs/RecordingSettingsDefaultApi.md#fetchrecordingsettings) | **Get** /v1/RecordingSettings/Default | 
*RecordingsApi* | [**DeleteRecording**](docs/RecordingsApi.md#deleterecording) | **Delete** /v1/Recordings/{Sid} | Delete a Recording resource identified by a Recording SID.
*RecordingsApi* | [**FetchRecording**](docs/RecordingsApi.md#fetchrecording) | **Get** /v1/Recordings/{Sid} | Returns a single Recording resource identified by a Recording SID.
*RecordingsApi* | [**ListRecording**](docs/RecordingsApi.md#listrecording) | **Get** /v1/Recordings | List of all Track recordings.
*RoomsApi* | [**CreateRoom**](docs/RoomsApi.md#createroom) | **Post** /v1/Rooms | 
*RoomsApi* | [**FetchRoom**](docs/RoomsApi.md#fetchroom) | **Get** /v1/Rooms/{Sid} | 
*RoomsApi* | [**ListRoom**](docs/RoomsApi.md#listroom) | **Get** /v1/Rooms | 
*RoomsApi* | [**UpdateRoom**](docs/RoomsApi.md#updateroom) | **Post** /v1/Rooms/{Sid} | 
*RoomsParticipantsApi* | [**FetchRoomParticipant**](docs/RoomsParticipantsApi.md#fetchroomparticipant) | **Get** /v1/Rooms/{RoomSid}/Participants/{Sid} | 
*RoomsParticipantsApi* | [**ListRoomParticipant**](docs/RoomsParticipantsApi.md#listroomparticipant) | **Get** /v1/Rooms/{RoomSid}/Participants | 
*RoomsParticipantsApi* | [**UpdateRoomParticipant**](docs/RoomsParticipantsApi.md#updateroomparticipant) | **Post** /v1/Rooms/{RoomSid}/Participants/{Sid} | 
*RoomsParticipantsAnonymizeApi* | [**UpdateRoomParticipantAnonymize**](docs/RoomsParticipantsAnonymizeApi.md#updateroomparticipantanonymize) | **Post** /v1/Rooms/{RoomSid}/Participants/{Sid}/Anonymize | 
*RoomsParticipantsPublishedTracksApi* | [**FetchRoomParticipantPublishedTrack**](docs/RoomsParticipantsPublishedTracksApi.md#fetchroomparticipantpublishedtrack) | **Get** /v1/Rooms/{RoomSid}/Participants/{ParticipantSid}/PublishedTracks/{Sid} | Returns a single Track resource represented by TrackName or SID.
*RoomsParticipantsPublishedTracksApi* | [**ListRoomParticipantPublishedTrack**](docs/RoomsParticipantsPublishedTracksApi.md#listroomparticipantpublishedtrack) | **Get** /v1/Rooms/{RoomSid}/Participants/{ParticipantSid}/PublishedTracks | Returns a list of tracks associated with a given Participant. Only &#x60;currently&#x60; Published Tracks are in the list resource.
*RoomsParticipantsSubscribeRulesApi* | [**FetchRoomParticipantSubscribeRule**](docs/RoomsParticipantsSubscribeRulesApi.md#fetchroomparticipantsubscriberule) | **Get** /v1/Rooms/{RoomSid}/Participants/{ParticipantSid}/SubscribeRules | Returns a list of Subscribe Rules for the Participant.
*RoomsParticipantsSubscribeRulesApi* | [**UpdateRoomParticipantSubscribeRule**](docs/RoomsParticipantsSubscribeRulesApi.md#updateroomparticipantsubscriberule) | **Post** /v1/Rooms/{RoomSid}/Participants/{ParticipantSid}/SubscribeRules | Update the Subscribe Rules for the Participant
*RoomsParticipantsSubscribedTracksApi* | [**FetchRoomParticipantSubscribedTrack**](docs/RoomsParticipantsSubscribedTracksApi.md#fetchroomparticipantsubscribedtrack) | **Get** /v1/Rooms/{RoomSid}/Participants/{ParticipantSid}/SubscribedTracks/{Sid} | Returns a single Track resource represented by &#x60;track_sid&#x60;.  Note: This is one resource with the Video API that requires a SID, be Track Name on the subscriber side is not guaranteed to be unique.
*RoomsParticipantsSubscribedTracksApi* | [**ListRoomParticipantSubscribedTrack**](docs/RoomsParticipantsSubscribedTracksApi.md#listroomparticipantsubscribedtrack) | **Get** /v1/Rooms/{RoomSid}/Participants/{ParticipantSid}/SubscribedTracks | Returns a list of tracks that are subscribed for the participant.
*RoomsRecordingRulesApi* | [**FetchRoomRecordingRule**](docs/RoomsRecordingRulesApi.md#fetchroomrecordingrule) | **Get** /v1/Rooms/{RoomSid}/RecordingRules | Returns a list of Recording Rules for the Room.
*RoomsRecordingRulesApi* | [**UpdateRoomRecordingRule**](docs/RoomsRecordingRulesApi.md#updateroomrecordingrule) | **Post** /v1/Rooms/{RoomSid}/RecordingRules | Update the Recording Rules for the Room
*RoomsRecordingsApi* | [**DeleteRoomRecording**](docs/RoomsRecordingsApi.md#deleteroomrecording) | **Delete** /v1/Rooms/{RoomSid}/Recordings/{Sid} | 
*RoomsRecordingsApi* | [**FetchRoomRecording**](docs/RoomsRecordingsApi.md#fetchroomrecording) | **Get** /v1/Rooms/{RoomSid}/Recordings/{Sid} | 
*RoomsRecordingsApi* | [**ListRoomRecording**](docs/RoomsRecordingsApi.md#listroomrecording) | **Get** /v1/Rooms/{RoomSid}/Recordings | 
*RoomsTranscriptionsApi* | [**CreateRoomTranscriptions**](docs/RoomsTranscriptionsApi.md#createroomtranscriptions) | **Post** /v1/Rooms/{RoomSid}/Transcriptions | 
*RoomsTranscriptionsApi* | [**FetchRoomTranscriptions**](docs/RoomsTranscriptionsApi.md#fetchroomtranscriptions) | **Get** /v1/Rooms/{RoomSid}/Transcriptions/{Ttid} | 
*RoomsTranscriptionsApi* | [**ListRoomTranscriptions**](docs/RoomsTranscriptionsApi.md#listroomtranscriptions) | **Get** /v1/Rooms/{RoomSid}/Transcriptions | 
*RoomsTranscriptionsApi* | [**UpdateRoomTranscriptions**](docs/RoomsTranscriptionsApi.md#updateroomtranscriptions) | **Post** /v1/Rooms/{RoomSid}/Transcriptions/{Ttid} | 


## Documentation For Models

 - [VideoV1RoomRecording](docs/VideoV1RoomRecording.md)
 - [VideoV1RoomParticipant](docs/VideoV1RoomParticipant.md)
 - [VideoV1RoomParticipantPublishedTrack](docs/VideoV1RoomParticipantPublishedTrack.md)
 - [VideoV1RoomTranscriptions](docs/VideoV1RoomTranscriptions.md)
 - [VideoV1CompositionSettings](docs/VideoV1CompositionSettings.md)
 - [VideoV1RoomParticipantSubscribeRule](docs/VideoV1RoomParticipantSubscribeRule.md)
 - [VideoV1CompositionHook](docs/VideoV1CompositionHook.md)
 - [ListCompositionResponse](docs/ListCompositionResponse.md)
 - [VideoV1Room](docs/VideoV1Room.md)
 - [ListRoomTranscriptionsResponse](docs/ListRoomTranscriptionsResponse.md)
 - [ListRoomParticipantSubscribedTrackResponse](docs/ListRoomParticipantSubscribedTrackResponse.md)
 - [ListRoomParticipantResponse](docs/ListRoomParticipantResponse.md)
 - [ListRoomParticipantPublishedTrackResponse](docs/ListRoomParticipantPublishedTrackResponse.md)
 - [VideoV1Composition](docs/VideoV1Composition.md)
 - [ListRoomRecordingResponse](docs/ListRoomRecordingResponse.md)
 - [VideoV1RoomRecordingRule](docs/VideoV1RoomRecordingRule.md)
 - [VideoV1RoomRoomRecordingRuleRules](docs/VideoV1RoomRoomRecordingRuleRules.md)
 - [ListRecordingResponse](docs/ListRecordingResponse.md)
 - [VideoV1Recording](docs/VideoV1Recording.md)
 - [VideoV1RoomParticipantSubscribedTrack](docs/VideoV1RoomParticipantSubscribedTrack.md)
 - [ListRoomResponse](docs/ListRoomResponse.md)
 - [ListCompositionHookResponse](docs/ListCompositionHookResponse.md)
 - [VideoV1RoomRoomParticipantRoomParticipantSubscribeRuleRules](docs/VideoV1RoomRoomParticipantRoomParticipantSubscribeRuleRules.md)
 - [ListCompositionResponseMeta](docs/ListCompositionResponseMeta.md)
 - [VideoV1RoomParticipantAnonymize](docs/VideoV1RoomParticipantAnonymize.md)
 - [VideoV1RecordingSettings](docs/VideoV1RecordingSettings.md)


## Documentation For Authorization



## accountSid_authToken

- **Type**: HTTP basic authentication

Example

```golang
auth := context.WithValue(context.Background(), sw.ContextBasicAuth, sw.BasicAuth{
    UserName: "username",
    Password: "password",
})
r, err := client.Service.Operation(auth, args)
```

