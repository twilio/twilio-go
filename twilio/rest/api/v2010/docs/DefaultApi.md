# \DefaultApi

All URIs are relative to *https://api.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAccount**](DefaultApi.md#CreateAccount) | **Post** /2010-04-01/Accounts.json | 
[**CreateAddress**](DefaultApi.md#CreateAddress) | **Post** /2010-04-01/Accounts/{AccountSid}/Addresses.json | 
[**CreateApplication**](DefaultApi.md#CreateApplication) | **Post** /2010-04-01/Accounts/{AccountSid}/Applications.json | 
[**CreateCall**](DefaultApi.md#CreateCall) | **Post** /2010-04-01/Accounts/{AccountSid}/Calls.json | 
[**CreateCallFeedbackSummary**](DefaultApi.md#CreateCallFeedbackSummary) | **Post** /2010-04-01/Accounts/{AccountSid}/Calls/FeedbackSummary.json | 
[**CreateCallRecording**](DefaultApi.md#CreateCallRecording) | **Post** /2010-04-01/Accounts/{AccountSid}/Calls/{CallSid}/Recordings.json | 
[**CreateIncomingPhoneNumber**](DefaultApi.md#CreateIncomingPhoneNumber) | **Post** /2010-04-01/Accounts/{AccountSid}/IncomingPhoneNumbers.json | 
[**CreateIncomingPhoneNumberAssignedAddOn**](DefaultApi.md#CreateIncomingPhoneNumberAssignedAddOn) | **Post** /2010-04-01/Accounts/{AccountSid}/IncomingPhoneNumbers/{ResourceSid}/AssignedAddOns.json | 
[**CreateIncomingPhoneNumberLocal**](DefaultApi.md#CreateIncomingPhoneNumberLocal) | **Post** /2010-04-01/Accounts/{AccountSid}/IncomingPhoneNumbers/Local.json | 
[**CreateIncomingPhoneNumberMobile**](DefaultApi.md#CreateIncomingPhoneNumberMobile) | **Post** /2010-04-01/Accounts/{AccountSid}/IncomingPhoneNumbers/Mobile.json | 
[**CreateIncomingPhoneNumberTollFree**](DefaultApi.md#CreateIncomingPhoneNumberTollFree) | **Post** /2010-04-01/Accounts/{AccountSid}/IncomingPhoneNumbers/TollFree.json | 
[**CreateMessage**](DefaultApi.md#CreateMessage) | **Post** /2010-04-01/Accounts/{AccountSid}/Messages.json | 
[**CreateMessageFeedback**](DefaultApi.md#CreateMessageFeedback) | **Post** /2010-04-01/Accounts/{AccountSid}/Messages/{MessageSid}/Feedback.json | 
[**CreateNewKey**](DefaultApi.md#CreateNewKey) | **Post** /2010-04-01/Accounts/{AccountSid}/Keys.json | 
[**CreateNewSigningKey**](DefaultApi.md#CreateNewSigningKey) | **Post** /2010-04-01/Accounts/{AccountSid}/SigningKeys.json | 
[**CreateParticipant**](DefaultApi.md#CreateParticipant) | **Post** /2010-04-01/Accounts/{AccountSid}/Conferences/{ConferenceSid}/Participants.json | 
[**CreatePayments**](DefaultApi.md#CreatePayments) | **Post** /2010-04-01/Accounts/{AccountSid}/Calls/{CallSid}/Payments.json | 
[**CreateQueue**](DefaultApi.md#CreateQueue) | **Post** /2010-04-01/Accounts/{AccountSid}/Queues.json | 
[**CreateSipAuthCallsCredentialListMapping**](DefaultApi.md#CreateSipAuthCallsCredentialListMapping) | **Post** /2010-04-01/Accounts/{AccountSid}/SIP/Domains/{DomainSid}/Auth/Calls/CredentialListMappings.json | 
[**CreateSipAuthCallsIpAccessControlListMapping**](DefaultApi.md#CreateSipAuthCallsIpAccessControlListMapping) | **Post** /2010-04-01/Accounts/{AccountSid}/SIP/Domains/{DomainSid}/Auth/Calls/IpAccessControlListMappings.json | 
[**CreateSipAuthRegistrationsCredentialListMapping**](DefaultApi.md#CreateSipAuthRegistrationsCredentialListMapping) | **Post** /2010-04-01/Accounts/{AccountSid}/SIP/Domains/{DomainSid}/Auth/Registrations/CredentialListMappings.json | 
[**CreateSipCredential**](DefaultApi.md#CreateSipCredential) | **Post** /2010-04-01/Accounts/{AccountSid}/SIP/CredentialLists/{CredentialListSid}/Credentials.json | 
[**CreateSipCredentialList**](DefaultApi.md#CreateSipCredentialList) | **Post** /2010-04-01/Accounts/{AccountSid}/SIP/CredentialLists.json | 
[**CreateSipCredentialListMapping**](DefaultApi.md#CreateSipCredentialListMapping) | **Post** /2010-04-01/Accounts/{AccountSid}/SIP/Domains/{DomainSid}/CredentialListMappings.json | 
[**CreateSipDomain**](DefaultApi.md#CreateSipDomain) | **Post** /2010-04-01/Accounts/{AccountSid}/SIP/Domains.json | 
[**CreateSipIpAccessControlList**](DefaultApi.md#CreateSipIpAccessControlList) | **Post** /2010-04-01/Accounts/{AccountSid}/SIP/IpAccessControlLists.json | 
[**CreateSipIpAccessControlListMapping**](DefaultApi.md#CreateSipIpAccessControlListMapping) | **Post** /2010-04-01/Accounts/{AccountSid}/SIP/Domains/{DomainSid}/IpAccessControlListMappings.json | 
[**CreateSipIpAddress**](DefaultApi.md#CreateSipIpAddress) | **Post** /2010-04-01/Accounts/{AccountSid}/SIP/IpAccessControlLists/{IpAccessControlListSid}/IpAddresses.json | 
[**CreateToken**](DefaultApi.md#CreateToken) | **Post** /2010-04-01/Accounts/{AccountSid}/Tokens.json | 
[**CreateUsageTrigger**](DefaultApi.md#CreateUsageTrigger) | **Post** /2010-04-01/Accounts/{AccountSid}/Usage/Triggers.json | 
[**CreateValidationRequest**](DefaultApi.md#CreateValidationRequest) | **Post** /2010-04-01/Accounts/{AccountSid}/OutgoingCallerIds.json | 
[**DeleteAddress**](DefaultApi.md#DeleteAddress) | **Delete** /2010-04-01/Accounts/{AccountSid}/Addresses/{Sid}.json | 
[**DeleteApplication**](DefaultApi.md#DeleteApplication) | **Delete** /2010-04-01/Accounts/{AccountSid}/Applications/{Sid}.json | 
[**DeleteCall**](DefaultApi.md#DeleteCall) | **Delete** /2010-04-01/Accounts/{AccountSid}/Calls/{Sid}.json | 
[**DeleteCallFeedbackSummary**](DefaultApi.md#DeleteCallFeedbackSummary) | **Delete** /2010-04-01/Accounts/{AccountSid}/Calls/FeedbackSummary/{Sid}.json | 
[**DeleteCallRecording**](DefaultApi.md#DeleteCallRecording) | **Delete** /2010-04-01/Accounts/{AccountSid}/Calls/{CallSid}/Recordings/{Sid}.json | 
[**DeleteConferenceRecording**](DefaultApi.md#DeleteConferenceRecording) | **Delete** /2010-04-01/Accounts/{AccountSid}/Conferences/{ConferenceSid}/Recordings/{Sid}.json | 
[**DeleteConnectApp**](DefaultApi.md#DeleteConnectApp) | **Delete** /2010-04-01/Accounts/{AccountSid}/ConnectApps/{Sid}.json | 
[**DeleteIncomingPhoneNumber**](DefaultApi.md#DeleteIncomingPhoneNumber) | **Delete** /2010-04-01/Accounts/{AccountSid}/IncomingPhoneNumbers/{Sid}.json | 
[**DeleteIncomingPhoneNumberAssignedAddOn**](DefaultApi.md#DeleteIncomingPhoneNumberAssignedAddOn) | **Delete** /2010-04-01/Accounts/{AccountSid}/IncomingPhoneNumbers/{ResourceSid}/AssignedAddOns/{Sid}.json | 
[**DeleteKey**](DefaultApi.md#DeleteKey) | **Delete** /2010-04-01/Accounts/{AccountSid}/Keys/{Sid}.json | 
[**DeleteMedia**](DefaultApi.md#DeleteMedia) | **Delete** /2010-04-01/Accounts/{AccountSid}/Messages/{MessageSid}/Media/{Sid}.json | 
[**DeleteMessage**](DefaultApi.md#DeleteMessage) | **Delete** /2010-04-01/Accounts/{AccountSid}/Messages/{Sid}.json | 
[**DeleteOutgoingCallerId**](DefaultApi.md#DeleteOutgoingCallerId) | **Delete** /2010-04-01/Accounts/{AccountSid}/OutgoingCallerIds/{Sid}.json | 
[**DeleteParticipant**](DefaultApi.md#DeleteParticipant) | **Delete** /2010-04-01/Accounts/{AccountSid}/Conferences/{ConferenceSid}/Participants/{CallSid}.json | 
[**DeleteQueue**](DefaultApi.md#DeleteQueue) | **Delete** /2010-04-01/Accounts/{AccountSid}/Queues/{Sid}.json | 
[**DeleteRecording**](DefaultApi.md#DeleteRecording) | **Delete** /2010-04-01/Accounts/{AccountSid}/Recordings/{Sid}.json | 
[**DeleteRecordingAddOnResult**](DefaultApi.md#DeleteRecordingAddOnResult) | **Delete** /2010-04-01/Accounts/{AccountSid}/Recordings/{ReferenceSid}/AddOnResults/{Sid}.json | 
[**DeleteRecordingAddOnResultPayload**](DefaultApi.md#DeleteRecordingAddOnResultPayload) | **Delete** /2010-04-01/Accounts/{AccountSid}/Recordings/{ReferenceSid}/AddOnResults/{AddOnResultSid}/Payloads/{Sid}.json | 
[**DeleteRecordingTranscription**](DefaultApi.md#DeleteRecordingTranscription) | **Delete** /2010-04-01/Accounts/{AccountSid}/Recordings/{RecordingSid}/Transcriptions/{Sid}.json | 
[**DeleteSigningKey**](DefaultApi.md#DeleteSigningKey) | **Delete** /2010-04-01/Accounts/{AccountSid}/SigningKeys/{Sid}.json | 
[**DeleteSipAuthCallsCredentialListMapping**](DefaultApi.md#DeleteSipAuthCallsCredentialListMapping) | **Delete** /2010-04-01/Accounts/{AccountSid}/SIP/Domains/{DomainSid}/Auth/Calls/CredentialListMappings/{Sid}.json | 
[**DeleteSipAuthCallsIpAccessControlListMapping**](DefaultApi.md#DeleteSipAuthCallsIpAccessControlListMapping) | **Delete** /2010-04-01/Accounts/{AccountSid}/SIP/Domains/{DomainSid}/Auth/Calls/IpAccessControlListMappings/{Sid}.json | 
[**DeleteSipAuthRegistrationsCredentialListMapping**](DefaultApi.md#DeleteSipAuthRegistrationsCredentialListMapping) | **Delete** /2010-04-01/Accounts/{AccountSid}/SIP/Domains/{DomainSid}/Auth/Registrations/CredentialListMappings/{Sid}.json | 
[**DeleteSipCredential**](DefaultApi.md#DeleteSipCredential) | **Delete** /2010-04-01/Accounts/{AccountSid}/SIP/CredentialLists/{CredentialListSid}/Credentials/{Sid}.json | 
[**DeleteSipCredentialList**](DefaultApi.md#DeleteSipCredentialList) | **Delete** /2010-04-01/Accounts/{AccountSid}/SIP/CredentialLists/{Sid}.json | 
[**DeleteSipCredentialListMapping**](DefaultApi.md#DeleteSipCredentialListMapping) | **Delete** /2010-04-01/Accounts/{AccountSid}/SIP/Domains/{DomainSid}/CredentialListMappings/{Sid}.json | 
[**DeleteSipDomain**](DefaultApi.md#DeleteSipDomain) | **Delete** /2010-04-01/Accounts/{AccountSid}/SIP/Domains/{Sid}.json | 
[**DeleteSipIpAccessControlList**](DefaultApi.md#DeleteSipIpAccessControlList) | **Delete** /2010-04-01/Accounts/{AccountSid}/SIP/IpAccessControlLists/{Sid}.json | 
[**DeleteSipIpAccessControlListMapping**](DefaultApi.md#DeleteSipIpAccessControlListMapping) | **Delete** /2010-04-01/Accounts/{AccountSid}/SIP/Domains/{DomainSid}/IpAccessControlListMappings/{Sid}.json | 
[**DeleteSipIpAddress**](DefaultApi.md#DeleteSipIpAddress) | **Delete** /2010-04-01/Accounts/{AccountSid}/SIP/IpAccessControlLists/{IpAccessControlListSid}/IpAddresses/{Sid}.json | 
[**DeleteTranscription**](DefaultApi.md#DeleteTranscription) | **Delete** /2010-04-01/Accounts/{AccountSid}/Transcriptions/{Sid}.json | 
[**DeleteUsageTrigger**](DefaultApi.md#DeleteUsageTrigger) | **Delete** /2010-04-01/Accounts/{AccountSid}/Usage/Triggers/{Sid}.json | 
[**FetchAccount**](DefaultApi.md#FetchAccount) | **Get** /2010-04-01/Accounts/{Sid}.json | 
[**FetchAddress**](DefaultApi.md#FetchAddress) | **Get** /2010-04-01/Accounts/{AccountSid}/Addresses/{Sid}.json | 
[**FetchApplication**](DefaultApi.md#FetchApplication) | **Get** /2010-04-01/Accounts/{AccountSid}/Applications/{Sid}.json | 
[**FetchAuthorizedConnectApp**](DefaultApi.md#FetchAuthorizedConnectApp) | **Get** /2010-04-01/Accounts/{AccountSid}/AuthorizedConnectApps/{ConnectAppSid}.json | 
[**FetchAvailablePhoneNumberCountry**](DefaultApi.md#FetchAvailablePhoneNumberCountry) | **Get** /2010-04-01/Accounts/{AccountSid}/AvailablePhoneNumbers/{CountryCode}.json | 
[**FetchBalance**](DefaultApi.md#FetchBalance) | **Get** /2010-04-01/Accounts/{AccountSid}/Balance.json | 
[**FetchCall**](DefaultApi.md#FetchCall) | **Get** /2010-04-01/Accounts/{AccountSid}/Calls/{Sid}.json | 
[**FetchCallFeedback**](DefaultApi.md#FetchCallFeedback) | **Get** /2010-04-01/Accounts/{AccountSid}/Calls/{CallSid}/Feedback.json | 
[**FetchCallFeedbackSummary**](DefaultApi.md#FetchCallFeedbackSummary) | **Get** /2010-04-01/Accounts/{AccountSid}/Calls/FeedbackSummary/{Sid}.json | 
[**FetchCallNotification**](DefaultApi.md#FetchCallNotification) | **Get** /2010-04-01/Accounts/{AccountSid}/Calls/{CallSid}/Notifications/{Sid}.json | 
[**FetchCallRecording**](DefaultApi.md#FetchCallRecording) | **Get** /2010-04-01/Accounts/{AccountSid}/Calls/{CallSid}/Recordings/{Sid}.json | 
[**FetchConference**](DefaultApi.md#FetchConference) | **Get** /2010-04-01/Accounts/{AccountSid}/Conferences/{Sid}.json | 
[**FetchConferenceRecording**](DefaultApi.md#FetchConferenceRecording) | **Get** /2010-04-01/Accounts/{AccountSid}/Conferences/{ConferenceSid}/Recordings/{Sid}.json | 
[**FetchConnectApp**](DefaultApi.md#FetchConnectApp) | **Get** /2010-04-01/Accounts/{AccountSid}/ConnectApps/{Sid}.json | 
[**FetchIncomingPhoneNumber**](DefaultApi.md#FetchIncomingPhoneNumber) | **Get** /2010-04-01/Accounts/{AccountSid}/IncomingPhoneNumbers/{Sid}.json | 
[**FetchIncomingPhoneNumberAssignedAddOn**](DefaultApi.md#FetchIncomingPhoneNumberAssignedAddOn) | **Get** /2010-04-01/Accounts/{AccountSid}/IncomingPhoneNumbers/{ResourceSid}/AssignedAddOns/{Sid}.json | 
[**FetchIncomingPhoneNumberAssignedAddOnExtension**](DefaultApi.md#FetchIncomingPhoneNumberAssignedAddOnExtension) | **Get** /2010-04-01/Accounts/{AccountSid}/IncomingPhoneNumbers/{ResourceSid}/AssignedAddOns/{AssignedAddOnSid}/Extensions/{Sid}.json | 
[**FetchKey**](DefaultApi.md#FetchKey) | **Get** /2010-04-01/Accounts/{AccountSid}/Keys/{Sid}.json | 
[**FetchMedia**](DefaultApi.md#FetchMedia) | **Get** /2010-04-01/Accounts/{AccountSid}/Messages/{MessageSid}/Media/{Sid}.json | 
[**FetchMember**](DefaultApi.md#FetchMember) | **Get** /2010-04-01/Accounts/{AccountSid}/Queues/{QueueSid}/Members/{CallSid}.json | 
[**FetchMessage**](DefaultApi.md#FetchMessage) | **Get** /2010-04-01/Accounts/{AccountSid}/Messages/{Sid}.json | 
[**FetchNotification**](DefaultApi.md#FetchNotification) | **Get** /2010-04-01/Accounts/{AccountSid}/Notifications/{Sid}.json | 
[**FetchOutgoingCallerId**](DefaultApi.md#FetchOutgoingCallerId) | **Get** /2010-04-01/Accounts/{AccountSid}/OutgoingCallerIds/{Sid}.json | 
[**FetchParticipant**](DefaultApi.md#FetchParticipant) | **Get** /2010-04-01/Accounts/{AccountSid}/Conferences/{ConferenceSid}/Participants/{CallSid}.json | 
[**FetchQueue**](DefaultApi.md#FetchQueue) | **Get** /2010-04-01/Accounts/{AccountSid}/Queues/{Sid}.json | 
[**FetchRecording**](DefaultApi.md#FetchRecording) | **Get** /2010-04-01/Accounts/{AccountSid}/Recordings/{Sid}.json | 
[**FetchRecordingAddOnResult**](DefaultApi.md#FetchRecordingAddOnResult) | **Get** /2010-04-01/Accounts/{AccountSid}/Recordings/{ReferenceSid}/AddOnResults/{Sid}.json | 
[**FetchRecordingAddOnResultPayload**](DefaultApi.md#FetchRecordingAddOnResultPayload) | **Get** /2010-04-01/Accounts/{AccountSid}/Recordings/{ReferenceSid}/AddOnResults/{AddOnResultSid}/Payloads/{Sid}.json | 
[**FetchRecordingTranscription**](DefaultApi.md#FetchRecordingTranscription) | **Get** /2010-04-01/Accounts/{AccountSid}/Recordings/{RecordingSid}/Transcriptions/{Sid}.json | 
[**FetchShortCode**](DefaultApi.md#FetchShortCode) | **Get** /2010-04-01/Accounts/{AccountSid}/SMS/ShortCodes/{Sid}.json | 
[**FetchSigningKey**](DefaultApi.md#FetchSigningKey) | **Get** /2010-04-01/Accounts/{AccountSid}/SigningKeys/{Sid}.json | 
[**FetchSipAuthCallsCredentialListMapping**](DefaultApi.md#FetchSipAuthCallsCredentialListMapping) | **Get** /2010-04-01/Accounts/{AccountSid}/SIP/Domains/{DomainSid}/Auth/Calls/CredentialListMappings/{Sid}.json | 
[**FetchSipAuthCallsIpAccessControlListMapping**](DefaultApi.md#FetchSipAuthCallsIpAccessControlListMapping) | **Get** /2010-04-01/Accounts/{AccountSid}/SIP/Domains/{DomainSid}/Auth/Calls/IpAccessControlListMappings/{Sid}.json | 
[**FetchSipAuthRegistrationsCredentialListMapping**](DefaultApi.md#FetchSipAuthRegistrationsCredentialListMapping) | **Get** /2010-04-01/Accounts/{AccountSid}/SIP/Domains/{DomainSid}/Auth/Registrations/CredentialListMappings/{Sid}.json | 
[**FetchSipCredential**](DefaultApi.md#FetchSipCredential) | **Get** /2010-04-01/Accounts/{AccountSid}/SIP/CredentialLists/{CredentialListSid}/Credentials/{Sid}.json | 
[**FetchSipCredentialList**](DefaultApi.md#FetchSipCredentialList) | **Get** /2010-04-01/Accounts/{AccountSid}/SIP/CredentialLists/{Sid}.json | 
[**FetchSipCredentialListMapping**](DefaultApi.md#FetchSipCredentialListMapping) | **Get** /2010-04-01/Accounts/{AccountSid}/SIP/Domains/{DomainSid}/CredentialListMappings/{Sid}.json | 
[**FetchSipDomain**](DefaultApi.md#FetchSipDomain) | **Get** /2010-04-01/Accounts/{AccountSid}/SIP/Domains/{Sid}.json | 
[**FetchSipIpAccessControlList**](DefaultApi.md#FetchSipIpAccessControlList) | **Get** /2010-04-01/Accounts/{AccountSid}/SIP/IpAccessControlLists/{Sid}.json | 
[**FetchSipIpAccessControlListMapping**](DefaultApi.md#FetchSipIpAccessControlListMapping) | **Get** /2010-04-01/Accounts/{AccountSid}/SIP/Domains/{DomainSid}/IpAccessControlListMappings/{Sid}.json | 
[**FetchSipIpAddress**](DefaultApi.md#FetchSipIpAddress) | **Get** /2010-04-01/Accounts/{AccountSid}/SIP/IpAccessControlLists/{IpAccessControlListSid}/IpAddresses/{Sid}.json | 
[**FetchTranscription**](DefaultApi.md#FetchTranscription) | **Get** /2010-04-01/Accounts/{AccountSid}/Transcriptions/{Sid}.json | 
[**FetchUsageTrigger**](DefaultApi.md#FetchUsageTrigger) | **Get** /2010-04-01/Accounts/{AccountSid}/Usage/Triggers/{Sid}.json | 
[**ListAccount**](DefaultApi.md#ListAccount) | **Get** /2010-04-01/Accounts.json | 
[**ListAddress**](DefaultApi.md#ListAddress) | **Get** /2010-04-01/Accounts/{AccountSid}/Addresses.json | 
[**ListApplication**](DefaultApi.md#ListApplication) | **Get** /2010-04-01/Accounts/{AccountSid}/Applications.json | 
[**ListAuthorizedConnectApp**](DefaultApi.md#ListAuthorizedConnectApp) | **Get** /2010-04-01/Accounts/{AccountSid}/AuthorizedConnectApps.json | 
[**ListAvailablePhoneNumberCountry**](DefaultApi.md#ListAvailablePhoneNumberCountry) | **Get** /2010-04-01/Accounts/{AccountSid}/AvailablePhoneNumbers.json | 
[**ListAvailablePhoneNumberLocal**](DefaultApi.md#ListAvailablePhoneNumberLocal) | **Get** /2010-04-01/Accounts/{AccountSid}/AvailablePhoneNumbers/{CountryCode}/Local.json | 
[**ListAvailablePhoneNumberMachineToMachine**](DefaultApi.md#ListAvailablePhoneNumberMachineToMachine) | **Get** /2010-04-01/Accounts/{AccountSid}/AvailablePhoneNumbers/{CountryCode}/MachineToMachine.json | 
[**ListAvailablePhoneNumberMobile**](DefaultApi.md#ListAvailablePhoneNumberMobile) | **Get** /2010-04-01/Accounts/{AccountSid}/AvailablePhoneNumbers/{CountryCode}/Mobile.json | 
[**ListAvailablePhoneNumberNational**](DefaultApi.md#ListAvailablePhoneNumberNational) | **Get** /2010-04-01/Accounts/{AccountSid}/AvailablePhoneNumbers/{CountryCode}/National.json | 
[**ListAvailablePhoneNumberSharedCost**](DefaultApi.md#ListAvailablePhoneNumberSharedCost) | **Get** /2010-04-01/Accounts/{AccountSid}/AvailablePhoneNumbers/{CountryCode}/SharedCost.json | 
[**ListAvailablePhoneNumberTollFree**](DefaultApi.md#ListAvailablePhoneNumberTollFree) | **Get** /2010-04-01/Accounts/{AccountSid}/AvailablePhoneNumbers/{CountryCode}/TollFree.json | 
[**ListAvailablePhoneNumberVoip**](DefaultApi.md#ListAvailablePhoneNumberVoip) | **Get** /2010-04-01/Accounts/{AccountSid}/AvailablePhoneNumbers/{CountryCode}/Voip.json | 
[**ListCall**](DefaultApi.md#ListCall) | **Get** /2010-04-01/Accounts/{AccountSid}/Calls.json | 
[**ListCallEvent**](DefaultApi.md#ListCallEvent) | **Get** /2010-04-01/Accounts/{AccountSid}/Calls/{CallSid}/Events.json | 
[**ListCallNotification**](DefaultApi.md#ListCallNotification) | **Get** /2010-04-01/Accounts/{AccountSid}/Calls/{CallSid}/Notifications.json | 
[**ListCallRecording**](DefaultApi.md#ListCallRecording) | **Get** /2010-04-01/Accounts/{AccountSid}/Calls/{CallSid}/Recordings.json | 
[**ListConference**](DefaultApi.md#ListConference) | **Get** /2010-04-01/Accounts/{AccountSid}/Conferences.json | 
[**ListConferenceRecording**](DefaultApi.md#ListConferenceRecording) | **Get** /2010-04-01/Accounts/{AccountSid}/Conferences/{ConferenceSid}/Recordings.json | 
[**ListConnectApp**](DefaultApi.md#ListConnectApp) | **Get** /2010-04-01/Accounts/{AccountSid}/ConnectApps.json | 
[**ListDependentPhoneNumber**](DefaultApi.md#ListDependentPhoneNumber) | **Get** /2010-04-01/Accounts/{AccountSid}/Addresses/{AddressSid}/DependentPhoneNumbers.json | 
[**ListIncomingPhoneNumber**](DefaultApi.md#ListIncomingPhoneNumber) | **Get** /2010-04-01/Accounts/{AccountSid}/IncomingPhoneNumbers.json | 
[**ListIncomingPhoneNumberAssignedAddOn**](DefaultApi.md#ListIncomingPhoneNumberAssignedAddOn) | **Get** /2010-04-01/Accounts/{AccountSid}/IncomingPhoneNumbers/{ResourceSid}/AssignedAddOns.json | 
[**ListIncomingPhoneNumberAssignedAddOnExtension**](DefaultApi.md#ListIncomingPhoneNumberAssignedAddOnExtension) | **Get** /2010-04-01/Accounts/{AccountSid}/IncomingPhoneNumbers/{ResourceSid}/AssignedAddOns/{AssignedAddOnSid}/Extensions.json | 
[**ListIncomingPhoneNumberLocal**](DefaultApi.md#ListIncomingPhoneNumberLocal) | **Get** /2010-04-01/Accounts/{AccountSid}/IncomingPhoneNumbers/Local.json | 
[**ListIncomingPhoneNumberMobile**](DefaultApi.md#ListIncomingPhoneNumberMobile) | **Get** /2010-04-01/Accounts/{AccountSid}/IncomingPhoneNumbers/Mobile.json | 
[**ListIncomingPhoneNumberTollFree**](DefaultApi.md#ListIncomingPhoneNumberTollFree) | **Get** /2010-04-01/Accounts/{AccountSid}/IncomingPhoneNumbers/TollFree.json | 
[**ListKey**](DefaultApi.md#ListKey) | **Get** /2010-04-01/Accounts/{AccountSid}/Keys.json | 
[**ListMedia**](DefaultApi.md#ListMedia) | **Get** /2010-04-01/Accounts/{AccountSid}/Messages/{MessageSid}/Media.json | 
[**ListMember**](DefaultApi.md#ListMember) | **Get** /2010-04-01/Accounts/{AccountSid}/Queues/{QueueSid}/Members.json | 
[**ListMessage**](DefaultApi.md#ListMessage) | **Get** /2010-04-01/Accounts/{AccountSid}/Messages.json | 
[**ListNotification**](DefaultApi.md#ListNotification) | **Get** /2010-04-01/Accounts/{AccountSid}/Notifications.json | 
[**ListOutgoingCallerId**](DefaultApi.md#ListOutgoingCallerId) | **Get** /2010-04-01/Accounts/{AccountSid}/OutgoingCallerIds.json | 
[**ListParticipant**](DefaultApi.md#ListParticipant) | **Get** /2010-04-01/Accounts/{AccountSid}/Conferences/{ConferenceSid}/Participants.json | 
[**ListQueue**](DefaultApi.md#ListQueue) | **Get** /2010-04-01/Accounts/{AccountSid}/Queues.json | 
[**ListRecording**](DefaultApi.md#ListRecording) | **Get** /2010-04-01/Accounts/{AccountSid}/Recordings.json | 
[**ListRecordingAddOnResult**](DefaultApi.md#ListRecordingAddOnResult) | **Get** /2010-04-01/Accounts/{AccountSid}/Recordings/{ReferenceSid}/AddOnResults.json | 
[**ListRecordingAddOnResultPayload**](DefaultApi.md#ListRecordingAddOnResultPayload) | **Get** /2010-04-01/Accounts/{AccountSid}/Recordings/{ReferenceSid}/AddOnResults/{AddOnResultSid}/Payloads.json | 
[**ListRecordingTranscription**](DefaultApi.md#ListRecordingTranscription) | **Get** /2010-04-01/Accounts/{AccountSid}/Recordings/{RecordingSid}/Transcriptions.json | 
[**ListShortCode**](DefaultApi.md#ListShortCode) | **Get** /2010-04-01/Accounts/{AccountSid}/SMS/ShortCodes.json | 
[**ListSigningKey**](DefaultApi.md#ListSigningKey) | **Get** /2010-04-01/Accounts/{AccountSid}/SigningKeys.json | 
[**ListSipAuthCallsCredentialListMapping**](DefaultApi.md#ListSipAuthCallsCredentialListMapping) | **Get** /2010-04-01/Accounts/{AccountSid}/SIP/Domains/{DomainSid}/Auth/Calls/CredentialListMappings.json | 
[**ListSipAuthCallsIpAccessControlListMapping**](DefaultApi.md#ListSipAuthCallsIpAccessControlListMapping) | **Get** /2010-04-01/Accounts/{AccountSid}/SIP/Domains/{DomainSid}/Auth/Calls/IpAccessControlListMappings.json | 
[**ListSipAuthRegistrationsCredentialListMapping**](DefaultApi.md#ListSipAuthRegistrationsCredentialListMapping) | **Get** /2010-04-01/Accounts/{AccountSid}/SIP/Domains/{DomainSid}/Auth/Registrations/CredentialListMappings.json | 
[**ListSipCredential**](DefaultApi.md#ListSipCredential) | **Get** /2010-04-01/Accounts/{AccountSid}/SIP/CredentialLists/{CredentialListSid}/Credentials.json | 
[**ListSipCredentialList**](DefaultApi.md#ListSipCredentialList) | **Get** /2010-04-01/Accounts/{AccountSid}/SIP/CredentialLists.json | 
[**ListSipCredentialListMapping**](DefaultApi.md#ListSipCredentialListMapping) | **Get** /2010-04-01/Accounts/{AccountSid}/SIP/Domains/{DomainSid}/CredentialListMappings.json | 
[**ListSipDomain**](DefaultApi.md#ListSipDomain) | **Get** /2010-04-01/Accounts/{AccountSid}/SIP/Domains.json | 
[**ListSipIpAccessControlList**](DefaultApi.md#ListSipIpAccessControlList) | **Get** /2010-04-01/Accounts/{AccountSid}/SIP/IpAccessControlLists.json | 
[**ListSipIpAccessControlListMapping**](DefaultApi.md#ListSipIpAccessControlListMapping) | **Get** /2010-04-01/Accounts/{AccountSid}/SIP/Domains/{DomainSid}/IpAccessControlListMappings.json | 
[**ListSipIpAddress**](DefaultApi.md#ListSipIpAddress) | **Get** /2010-04-01/Accounts/{AccountSid}/SIP/IpAccessControlLists/{IpAccessControlListSid}/IpAddresses.json | 
[**ListTranscription**](DefaultApi.md#ListTranscription) | **Get** /2010-04-01/Accounts/{AccountSid}/Transcriptions.json | 
[**ListUsageRecord**](DefaultApi.md#ListUsageRecord) | **Get** /2010-04-01/Accounts/{AccountSid}/Usage/Records.json | 
[**ListUsageRecordAllTime**](DefaultApi.md#ListUsageRecordAllTime) | **Get** /2010-04-01/Accounts/{AccountSid}/Usage/Records/AllTime.json | 
[**ListUsageRecordDaily**](DefaultApi.md#ListUsageRecordDaily) | **Get** /2010-04-01/Accounts/{AccountSid}/Usage/Records/Daily.json | 
[**ListUsageRecordLastMonth**](DefaultApi.md#ListUsageRecordLastMonth) | **Get** /2010-04-01/Accounts/{AccountSid}/Usage/Records/LastMonth.json | 
[**ListUsageRecordMonthly**](DefaultApi.md#ListUsageRecordMonthly) | **Get** /2010-04-01/Accounts/{AccountSid}/Usage/Records/Monthly.json | 
[**ListUsageRecordThisMonth**](DefaultApi.md#ListUsageRecordThisMonth) | **Get** /2010-04-01/Accounts/{AccountSid}/Usage/Records/ThisMonth.json | 
[**ListUsageRecordToday**](DefaultApi.md#ListUsageRecordToday) | **Get** /2010-04-01/Accounts/{AccountSid}/Usage/Records/Today.json | 
[**ListUsageRecordYearly**](DefaultApi.md#ListUsageRecordYearly) | **Get** /2010-04-01/Accounts/{AccountSid}/Usage/Records/Yearly.json | 
[**ListUsageRecordYesterday**](DefaultApi.md#ListUsageRecordYesterday) | **Get** /2010-04-01/Accounts/{AccountSid}/Usage/Records/Yesterday.json | 
[**ListUsageTrigger**](DefaultApi.md#ListUsageTrigger) | **Get** /2010-04-01/Accounts/{AccountSid}/Usage/Triggers.json | 
[**UpdateAccount**](DefaultApi.md#UpdateAccount) | **Post** /2010-04-01/Accounts/{Sid}.json | 
[**UpdateAddress**](DefaultApi.md#UpdateAddress) | **Post** /2010-04-01/Accounts/{AccountSid}/Addresses/{Sid}.json | 
[**UpdateApplication**](DefaultApi.md#UpdateApplication) | **Post** /2010-04-01/Accounts/{AccountSid}/Applications/{Sid}.json | 
[**UpdateCall**](DefaultApi.md#UpdateCall) | **Post** /2010-04-01/Accounts/{AccountSid}/Calls/{Sid}.json | 
[**UpdateCallFeedback**](DefaultApi.md#UpdateCallFeedback) | **Post** /2010-04-01/Accounts/{AccountSid}/Calls/{CallSid}/Feedback.json | 
[**UpdateCallRecording**](DefaultApi.md#UpdateCallRecording) | **Post** /2010-04-01/Accounts/{AccountSid}/Calls/{CallSid}/Recordings/{Sid}.json | 
[**UpdateConference**](DefaultApi.md#UpdateConference) | **Post** /2010-04-01/Accounts/{AccountSid}/Conferences/{Sid}.json | 
[**UpdateConferenceRecording**](DefaultApi.md#UpdateConferenceRecording) | **Post** /2010-04-01/Accounts/{AccountSid}/Conferences/{ConferenceSid}/Recordings/{Sid}.json | 
[**UpdateConnectApp**](DefaultApi.md#UpdateConnectApp) | **Post** /2010-04-01/Accounts/{AccountSid}/ConnectApps/{Sid}.json | 
[**UpdateIncomingPhoneNumber**](DefaultApi.md#UpdateIncomingPhoneNumber) | **Post** /2010-04-01/Accounts/{AccountSid}/IncomingPhoneNumbers/{Sid}.json | 
[**UpdateKey**](DefaultApi.md#UpdateKey) | **Post** /2010-04-01/Accounts/{AccountSid}/Keys/{Sid}.json | 
[**UpdateMember**](DefaultApi.md#UpdateMember) | **Post** /2010-04-01/Accounts/{AccountSid}/Queues/{QueueSid}/Members/{CallSid}.json | 
[**UpdateMessage**](DefaultApi.md#UpdateMessage) | **Post** /2010-04-01/Accounts/{AccountSid}/Messages/{Sid}.json | 
[**UpdateOutgoingCallerId**](DefaultApi.md#UpdateOutgoingCallerId) | **Post** /2010-04-01/Accounts/{AccountSid}/OutgoingCallerIds/{Sid}.json | 
[**UpdateParticipant**](DefaultApi.md#UpdateParticipant) | **Post** /2010-04-01/Accounts/{AccountSid}/Conferences/{ConferenceSid}/Participants/{CallSid}.json | 
[**UpdatePayments**](DefaultApi.md#UpdatePayments) | **Post** /2010-04-01/Accounts/{AccountSid}/Calls/{CallSid}/Payments/{Sid}.json | 
[**UpdateQueue**](DefaultApi.md#UpdateQueue) | **Post** /2010-04-01/Accounts/{AccountSid}/Queues/{Sid}.json | 
[**UpdateShortCode**](DefaultApi.md#UpdateShortCode) | **Post** /2010-04-01/Accounts/{AccountSid}/SMS/ShortCodes/{Sid}.json | 
[**UpdateSigningKey**](DefaultApi.md#UpdateSigningKey) | **Post** /2010-04-01/Accounts/{AccountSid}/SigningKeys/{Sid}.json | 
[**UpdateSipCredential**](DefaultApi.md#UpdateSipCredential) | **Post** /2010-04-01/Accounts/{AccountSid}/SIP/CredentialLists/{CredentialListSid}/Credentials/{Sid}.json | 
[**UpdateSipCredentialList**](DefaultApi.md#UpdateSipCredentialList) | **Post** /2010-04-01/Accounts/{AccountSid}/SIP/CredentialLists/{Sid}.json | 
[**UpdateSipDomain**](DefaultApi.md#UpdateSipDomain) | **Post** /2010-04-01/Accounts/{AccountSid}/SIP/Domains/{Sid}.json | 
[**UpdateSipIpAccessControlList**](DefaultApi.md#UpdateSipIpAccessControlList) | **Post** /2010-04-01/Accounts/{AccountSid}/SIP/IpAccessControlLists/{Sid}.json | 
[**UpdateSipIpAddress**](DefaultApi.md#UpdateSipIpAddress) | **Post** /2010-04-01/Accounts/{AccountSid}/SIP/IpAccessControlLists/{IpAccessControlListSid}/IpAddresses/{Sid}.json | 
[**UpdateUsageTrigger**](DefaultApi.md#UpdateUsageTrigger) | **Post** /2010-04-01/Accounts/{AccountSid}/Usage/Triggers/{Sid}.json | 



## CreateAccount

> ApiV2010Account CreateAccount(ctx).FriendlyName(FriendlyName).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    FriendlyName := "FriendlyName_example" // string | A human readable description of the account to create, defaults to `SubAccount Created at {YYYY-MM-DD HH:MM meridian}` (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateAccount(context.Background()).FriendlyName(FriendlyName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAccount`: ApiV2010Account
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateAccount`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateAccountParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------
 **FriendlyName** | **string** | A human readable description of the account to create, defaults to &#x60;SubAccount Created at {YYYY-MM-DD HH:MM meridian}&#x60;

### Return type

[**ApiV2010Account**](ApiV2010Account.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAddress

> ApiV2010AccountAddress CreateAddress(ctx, AccountSid).AutoCorrectAddress(AutoCorrectAddress).City(City).CustomerName(CustomerName).EmergencyEnabled(EmergencyEnabled).FriendlyName(FriendlyName).IsoCountry(IsoCountry).PostalCode(PostalCode).Region(Region).Street(Street).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AccountSid := "AccountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that will be responsible for the new Address resource.
    AutoCorrectAddress := true // bool | Whether we should automatically correct the address. Can be: `true` or `false` and the default is `true`. If empty or `true`, we will correct the address you provide if necessary. If `false`, we won't alter the address you provide. (optional)
    City := "City_example" // string | The city of the new address. (optional)
    CustomerName := "CustomerName_example" // string | The name to associate with the new address. (optional)
    EmergencyEnabled := true // bool | Whether to enable emergency calling on the new address. Can be: `true` or `false`. (optional)
    FriendlyName := "FriendlyName_example" // string | A descriptive string that you create to describe the new address. It can be up to 64 characters long. (optional)
    IsoCountry := "IsoCountry_example" // string | The ISO country code of the new address. (optional)
    PostalCode := "PostalCode_example" // string | The postal code of the new address. (optional)
    Region := "Region_example" // string | The state or region of the new address. (optional)
    Street := "Street_example" // string | The number and street address of the new address. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateAddress(context.Background(), AccountSid).AutoCorrectAddress(AutoCorrectAddress).City(City).CustomerName(CustomerName).EmergencyEnabled(EmergencyEnabled).FriendlyName(FriendlyName).IsoCountry(IsoCountry).PostalCode(PostalCode).Region(Region).Street(Street).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateAddress``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAddress`: ApiV2010AccountAddress
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateAddress`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that will be responsible for the new Address resource.

### Other Parameters

Other parameters are passed through a pointer to a CreateAddressParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **AutoCorrectAddress** | **bool** | Whether we should automatically correct the address. Can be: &#x60;true&#x60; or &#x60;false&#x60; and the default is &#x60;true&#x60;. If empty or &#x60;true&#x60;, we will correct the address you provide if necessary. If &#x60;false&#x60;, we won&#39;t alter the address you provide.
 **City** | **string** | The city of the new address.
 **CustomerName** | **string** | The name to associate with the new address.
 **EmergencyEnabled** | **bool** | Whether to enable emergency calling on the new address. Can be: &#x60;true&#x60; or &#x60;false&#x60;.
 **FriendlyName** | **string** | A descriptive string that you create to describe the new address. It can be up to 64 characters long.
 **IsoCountry** | **string** | The ISO country code of the new address.
 **PostalCode** | **string** | The postal code of the new address.
 **Region** | **string** | The state or region of the new address.
 **Street** | **string** | The number and street address of the new address.

### Return type

[**ApiV2010AccountAddress**](ApiV2010AccountAddress.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateApplication

> ApiV2010AccountApplication CreateApplication(ctx, AccountSid).ApiVersion(ApiVersion).FriendlyName(FriendlyName).MessageStatusCallback(MessageStatusCallback).SmsFallbackMethod(SmsFallbackMethod).SmsFallbackUrl(SmsFallbackUrl).SmsMethod(SmsMethod).SmsStatusCallback(SmsStatusCallback).SmsUrl(SmsUrl).StatusCallback(StatusCallback).StatusCallbackMethod(StatusCallbackMethod).VoiceCallerIdLookup(VoiceCallerIdLookup).VoiceFallbackMethod(VoiceFallbackMethod).VoiceFallbackUrl(VoiceFallbackUrl).VoiceMethod(VoiceMethod).VoiceUrl(VoiceUrl).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AccountSid := "AccountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that will create the resource.
    ApiVersion := "ApiVersion_example" // string | The API version to use to start a new TwiML session. Can be: `2010-04-01` or `2008-08-01`. The default value is the account's default API version. (optional)
    FriendlyName := "FriendlyName_example" // string | A descriptive string that you create to describe the new application. It can be up to 64 characters long. (optional)
    MessageStatusCallback := "MessageStatusCallback_example" // string | The URL we should call using a POST method to send message status information to your application. (optional)
    SmsFallbackMethod := "SmsFallbackMethod_example" // string | The HTTP method we should use to call `sms_fallback_url`. Can be: `GET` or `POST`. (optional)
    SmsFallbackUrl := "SmsFallbackUrl_example" // string | The URL that we should call when an error occurs while retrieving or executing the TwiML from `sms_url`. (optional)
    SmsMethod := "SmsMethod_example" // string | The HTTP method we should use to call `sms_url`. Can be: `GET` or `POST`. (optional)
    SmsStatusCallback := "SmsStatusCallback_example" // string | The URL we should call using a POST method to send status information about SMS messages sent by the application. (optional)
    SmsUrl := "SmsUrl_example" // string | The URL we should call when the phone number receives an incoming SMS message. (optional)
    StatusCallback := "StatusCallback_example" // string | The URL we should call using the `status_callback_method` to send status information to your application. (optional)
    StatusCallbackMethod := "StatusCallbackMethod_example" // string | The HTTP method we should use to call `status_callback`. Can be: `GET` or `POST`. (optional)
    VoiceCallerIdLookup := true // bool | Whether we should look up the caller's caller-ID name from the CNAM database (additional charges apply). Can be: `true` or `false`. (optional)
    VoiceFallbackMethod := "VoiceFallbackMethod_example" // string | The HTTP method we should use to call `voice_fallback_url`. Can be: `GET` or `POST`. (optional)
    VoiceFallbackUrl := "VoiceFallbackUrl_example" // string | The URL that we should call when an error occurs retrieving or executing the TwiML requested by `url`. (optional)
    VoiceMethod := "VoiceMethod_example" // string | The HTTP method we should use to call `voice_url`. Can be: `GET` or `POST`. (optional)
    VoiceUrl := "VoiceUrl_example" // string | The URL we should call when the phone number assigned to this application receives a call. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateApplication(context.Background(), AccountSid).ApiVersion(ApiVersion).FriendlyName(FriendlyName).MessageStatusCallback(MessageStatusCallback).SmsFallbackMethod(SmsFallbackMethod).SmsFallbackUrl(SmsFallbackUrl).SmsMethod(SmsMethod).SmsStatusCallback(SmsStatusCallback).SmsUrl(SmsUrl).StatusCallback(StatusCallback).StatusCallbackMethod(StatusCallbackMethod).VoiceCallerIdLookup(VoiceCallerIdLookup).VoiceFallbackMethod(VoiceFallbackMethod).VoiceFallbackUrl(VoiceFallbackUrl).VoiceMethod(VoiceMethod).VoiceUrl(VoiceUrl).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateApplication``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateApplication`: ApiV2010AccountApplication
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateApplication`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that will create the resource.

### Other Parameters

Other parameters are passed through a pointer to a CreateApplicationParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **ApiVersion** | **string** | The API version to use to start a new TwiML session. Can be: &#x60;2010-04-01&#x60; or &#x60;2008-08-01&#x60;. The default value is the account&#39;s default API version.
 **FriendlyName** | **string** | A descriptive string that you create to describe the new application. It can be up to 64 characters long.
 **MessageStatusCallback** | **string** | The URL we should call using a POST method to send message status information to your application.
 **SmsFallbackMethod** | **string** | The HTTP method we should use to call &#x60;sms_fallback_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60;.
 **SmsFallbackUrl** | **string** | The URL that we should call when an error occurs while retrieving or executing the TwiML from &#x60;sms_url&#x60;.
 **SmsMethod** | **string** | The HTTP method we should use to call &#x60;sms_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60;.
 **SmsStatusCallback** | **string** | The URL we should call using a POST method to send status information about SMS messages sent by the application.
 **SmsUrl** | **string** | The URL we should call when the phone number receives an incoming SMS message.
 **StatusCallback** | **string** | The URL we should call using the &#x60;status_callback_method&#x60; to send status information to your application.
 **StatusCallbackMethod** | **string** | The HTTP method we should use to call &#x60;status_callback&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60;.
 **VoiceCallerIdLookup** | **bool** | Whether we should look up the caller&#39;s caller-ID name from the CNAM database (additional charges apply). Can be: &#x60;true&#x60; or &#x60;false&#x60;.
 **VoiceFallbackMethod** | **string** | The HTTP method we should use to call &#x60;voice_fallback_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60;.
 **VoiceFallbackUrl** | **string** | The URL that we should call when an error occurs retrieving or executing the TwiML requested by &#x60;url&#x60;.
 **VoiceMethod** | **string** | The HTTP method we should use to call &#x60;voice_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60;.
 **VoiceUrl** | **string** | The URL we should call when the phone number assigned to this application receives a call.

### Return type

[**ApiV2010AccountApplication**](ApiV2010AccountApplication.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCall

> ApiV2010AccountCall CreateCall(ctx, AccountSid).ApplicationSid(ApplicationSid).AsyncAmd(AsyncAmd).AsyncAmdStatusCallback(AsyncAmdStatusCallback).AsyncAmdStatusCallbackMethod(AsyncAmdStatusCallbackMethod).Byoc(Byoc).CallReason(CallReason).CallToken(CallToken).CallerId(CallerId).FallbackMethod(FallbackMethod).FallbackUrl(FallbackUrl).From(From).MachineDetection(MachineDetection).MachineDetectionSilenceTimeout(MachineDetectionSilenceTimeout).MachineDetectionSpeechEndThreshold(MachineDetectionSpeechEndThreshold).MachineDetectionSpeechThreshold(MachineDetectionSpeechThreshold).MachineDetectionTimeout(MachineDetectionTimeout).Method(Method).Record(Record).RecordingChannels(RecordingChannels).RecordingStatusCallback(RecordingStatusCallback).RecordingStatusCallbackEvent(RecordingStatusCallbackEvent).RecordingStatusCallbackMethod(RecordingStatusCallbackMethod).RecordingTrack(RecordingTrack).SendDigits(SendDigits).SipAuthPassword(SipAuthPassword).SipAuthUsername(SipAuthUsername).StatusCallback(StatusCallback).StatusCallbackEvent(StatusCallbackEvent).StatusCallbackMethod(StatusCallbackMethod).Timeout(Timeout).To(To).Trim(Trim).Twiml(Twiml).Url(Url).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AccountSid := "AccountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that will create the resource.
    ApplicationSid := "ApplicationSid_example" // string | The SID of the Application resource that will handle the call, if the call will be handled by an application. (optional)
    AsyncAmd := "AsyncAmd_example" // string | Select whether to perform answering machine detection in the background. Default, blocks the execution of the call until Answering Machine Detection is completed. Can be: `true` or `false`. (optional)
    AsyncAmdStatusCallback := "AsyncAmdStatusCallback_example" // string | The URL that we should call using the `async_amd_status_callback_method` to notify customer application whether the call was answered by human, machine or fax. (optional)
    AsyncAmdStatusCallbackMethod := "AsyncAmdStatusCallbackMethod_example" // string | The HTTP method we should use when calling the `async_amd_status_callback` URL. Can be: `GET` or `POST` and the default is `POST`. (optional)
    Byoc := "Byoc_example" // string | The SID of a BYOC (Bring Your Own Carrier) trunk to route this call with. Note that `byoc` is only meaningful when `to` is a phone number; it will otherwise be ignored. (Beta) (optional)
    CallReason := "CallReason_example" // string | The Reason for the outgoing call. Use it to specify the purpose of the call that is presented on the called party's phone. (Branded Calls Beta) (optional)
    CallToken := "CallToken_example" // string | A token string needed to invoke a forwarded call. A call_token is generated when an incoming call is received on a Twilio number. this field should be populated by the incoming call's call_token to make this outgoing call as a forwarded call of incoming call. A forwarded call should bear the same caller-id of incoming call. (optional)
    CallerId := "CallerId_example" // string | The phone number, SIP address, or Client identifier that made this call. Phone numbers are in [E.164 format](https://wwnw.twilio.com/docs/glossary/what-e164) (e.g., +16175551212). SIP addresses are formatted as `name@company.com`. (optional)
    FallbackMethod := "FallbackMethod_example" // string | The HTTP method that we should use to request the `fallback_url`. Can be: `GET` or `POST` and the default is `POST`. If an `application_sid` parameter is present, this parameter is ignored. (optional)
    FallbackUrl := "FallbackUrl_example" // string | The URL that we call using the `fallback_method` if an error occurs when requesting or executing the TwiML at `url`. If an `application_sid` parameter is present, this parameter is ignored. (optional)
    From := "From_example" // string | The phone number or client identifier to use as the caller id. If using a phone number, it must be a Twilio number or a Verified [outgoing caller id](https://www.twilio.com/docs/voice/api/outgoing-caller-ids) for your account. If the `to` parameter is a phone number, `From` must also be a phone number. (optional)
    MachineDetection := "MachineDetection_example" // string | Whether to detect if a human, answering machine, or fax has picked up the call. Can be: `Enable` or `DetectMessageEnd`. Use `Enable` if you would like us to return `AnsweredBy` as soon as the called party is identified. Use `DetectMessageEnd`, if you would like to leave a message on an answering machine. If `send_digits` is provided, this parameter is ignored. For more information, see [Answering Machine Detection](https://www.twilio.com/docs/voice/answering-machine-detection). (optional)
    MachineDetectionSilenceTimeout := int32(56) // int32 | The number of milliseconds of initial silence after which an `unknown` AnsweredBy result will be returned. Possible Values: 2000-10000. Default: 5000. (optional)
    MachineDetectionSpeechEndThreshold := int32(56) // int32 | The number of milliseconds of silence after speech activity at which point the speech activity is considered complete. Possible Values: 500-5000. Default: 1200. (optional)
    MachineDetectionSpeechThreshold := int32(56) // int32 | The number of milliseconds that is used as the measuring stick for the length of the speech activity, where durations lower than this value will be interpreted as a human and longer than this value as a machine. Possible Values: 1000-6000. Default: 2400. (optional)
    MachineDetectionTimeout := int32(56) // int32 | The number of seconds that we should attempt to detect an answering machine before timing out and sending a voice request with `AnsweredBy` of `unknown`. The default timeout is 30 seconds. (optional)
    Method := "Method_example" // string | The HTTP method we should use when calling the `url` parameter's value. Can be: `GET` or `POST` and the default is `POST`. If an `application_sid` parameter is present, this parameter is ignored. (optional)
    Record := true // bool | Whether to record the call. Can be `true` to record the phone call, or `false` to not. The default is `false`. The `recording_url` is sent to the `status_callback` URL. (optional)
    RecordingChannels := "RecordingChannels_example" // string | The number of channels in the final recording. Can be: `mono` or `dual`. The default is `mono`. `mono` records both legs of the call in a single channel of the recording file. `dual` records each leg to a separate channel of the recording file. The first channel of a dual-channel recording contains the parent call and the second channel contains the child call. (optional)
    RecordingStatusCallback := "RecordingStatusCallback_example" // string | The URL that we call when the recording is available to be accessed. (optional)
    RecordingStatusCallbackEvent := []string{"Inner_example"} // []string | The recording status events that will trigger calls to the URL specified in `recording_status_callback`. Can be: `in-progress`, `completed` and `absent`. Defaults to `completed`. Separate  multiple values with a space. (optional)
    RecordingStatusCallbackMethod := "RecordingStatusCallbackMethod_example" // string | The HTTP method we should use when calling the `recording_status_callback` URL. Can be: `GET` or `POST` and the default is `POST`. (optional)
    RecordingTrack := "RecordingTrack_example" // string | The audio track to record for the call. Can be: `inbound`, `outbound` or `both`. The default is `both`. `inbound` records the audio that is received by Twilio. `outbound` records the audio that is generated from Twilio. `both` records the audio that is received and generated by Twilio. (optional)
    SendDigits := "SendDigits_example" // string | A string of keys to dial after connecting to the number, maximum of 32 digits. Valid digits in the string include: any digit (`0`-`9`), '`#`', '`*`' and '`w`', to insert a half second pause. For example, if you connected to a company phone number and wanted to pause for one second, and then dial extension 1234 followed by the pound key, the value of this parameter would be `ww1234#`. Remember to URL-encode this string, since the '`#`' character has special meaning in a URL. If both `SendDigits` and `MachineDetection` parameters are provided, then `MachineDetection` will be ignored. (optional)
    SipAuthPassword := "SipAuthPassword_example" // string | The password required to authenticate the user account specified in `sip_auth_username`. (optional)
    SipAuthUsername := "SipAuthUsername_example" // string | The username used to authenticate the caller making a SIP call. (optional)
    StatusCallback := "StatusCallback_example" // string | The URL we should call using the `status_callback_method` to send status information to your application. If no `status_callback_event` is specified, we will send the `completed` status. If an `application_sid` parameter is present, this parameter is ignored. URLs must contain a valid hostname (underscores are not permitted). (optional)
    StatusCallbackEvent := []string{"Inner_example"} // []string | The call progress events that we will send to the `status_callback` URL. Can be: `initiated`, `ringing`, `answered`, and `completed`. If no event is specified, we send the `completed` status. If you want to receive multiple events, specify each one in a separate `status_callback_event` parameter. See the code sample for [monitoring call progress](https://www.twilio.com/docs/voice/api/call-resource?code-sample=code-create-a-call-resource-and-specify-a-statuscallbackevent&code-sdk-version=json). If an `application_sid` is present, this parameter is ignored. (optional)
    StatusCallbackMethod := "StatusCallbackMethod_example" // string | The HTTP method we should use when calling the `status_callback` URL. Can be: `GET` or `POST` and the default is `POST`. If an `application_sid` parameter is present, this parameter is ignored. (optional)
    Timeout := int32(56) // int32 | The integer number of seconds that we should allow the phone to ring before assuming there is no answer. The default is `60` seconds and the maximum is `600` seconds. For some call flows, we will add a 5-second buffer to the timeout value you provide. For this reason, a timeout value of 10 seconds could result in an actual timeout closer to 15 seconds. You can set this to a short time, such as `15` seconds, to hang up before reaching an answering machine or voicemail. (optional)
    To := "To_example" // string | The phone number, SIP address, or client identifier to call. (optional)
    Trim := "Trim_example" // string | Whether to trim any leading and trailing silence from the recording. Can be: `trim-silence` or `do-not-trim` and the default is `trim-silence`. (optional)
    Twiml := "Twiml_example" // string | TwiML instructions for the call Twilio will use without fetching Twiml from url parameter. If both `twiml` and `url` are provided then `twiml` parameter will be ignored. (optional)
    Url := "Url_example" // string | The absolute URL that returns the TwiML instructions for the call. We will call this URL using the `method` when the call connects. For more information, see the [Url Parameter](https://www.twilio.com/docs/voice/make-calls#specify-a-url-parameter) section in [Making Calls](https://www.twilio.com/docs/voice/make-calls). (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateCall(context.Background(), AccountSid).ApplicationSid(ApplicationSid).AsyncAmd(AsyncAmd).AsyncAmdStatusCallback(AsyncAmdStatusCallback).AsyncAmdStatusCallbackMethod(AsyncAmdStatusCallbackMethod).Byoc(Byoc).CallReason(CallReason).CallToken(CallToken).CallerId(CallerId).FallbackMethod(FallbackMethod).FallbackUrl(FallbackUrl).From(From).MachineDetection(MachineDetection).MachineDetectionSilenceTimeout(MachineDetectionSilenceTimeout).MachineDetectionSpeechEndThreshold(MachineDetectionSpeechEndThreshold).MachineDetectionSpeechThreshold(MachineDetectionSpeechThreshold).MachineDetectionTimeout(MachineDetectionTimeout).Method(Method).Record(Record).RecordingChannels(RecordingChannels).RecordingStatusCallback(RecordingStatusCallback).RecordingStatusCallbackEvent(RecordingStatusCallbackEvent).RecordingStatusCallbackMethod(RecordingStatusCallbackMethod).RecordingTrack(RecordingTrack).SendDigits(SendDigits).SipAuthPassword(SipAuthPassword).SipAuthUsername(SipAuthUsername).StatusCallback(StatusCallback).StatusCallbackEvent(StatusCallbackEvent).StatusCallbackMethod(StatusCallbackMethod).Timeout(Timeout).To(To).Trim(Trim).Twiml(Twiml).Url(Url).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateCall``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCall`: ApiV2010AccountCall
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateCall`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that will create the resource.

### Other Parameters

Other parameters are passed through a pointer to a CreateCallParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **ApplicationSid** | **string** | The SID of the Application resource that will handle the call, if the call will be handled by an application.
 **AsyncAmd** | **string** | Select whether to perform answering machine detection in the background. Default, blocks the execution of the call until Answering Machine Detection is completed. Can be: &#x60;true&#x60; or &#x60;false&#x60;.
 **AsyncAmdStatusCallback** | **string** | The URL that we should call using the &#x60;async_amd_status_callback_method&#x60; to notify customer application whether the call was answered by human, machine or fax.
 **AsyncAmdStatusCallbackMethod** | **string** | The HTTP method we should use when calling the &#x60;async_amd_status_callback&#x60; URL. Can be: &#x60;GET&#x60; or &#x60;POST&#x60; and the default is &#x60;POST&#x60;.
 **Byoc** | **string** | The SID of a BYOC (Bring Your Own Carrier) trunk to route this call with. Note that &#x60;byoc&#x60; is only meaningful when &#x60;to&#x60; is a phone number; it will otherwise be ignored. (Beta)
 **CallReason** | **string** | The Reason for the outgoing call. Use it to specify the purpose of the call that is presented on the called party&#39;s phone. (Branded Calls Beta)
 **CallToken** | **string** | A token string needed to invoke a forwarded call. A call_token is generated when an incoming call is received on a Twilio number. this field should be populated by the incoming call&#39;s call_token to make this outgoing call as a forwarded call of incoming call. A forwarded call should bear the same caller-id of incoming call.
 **CallerId** | **string** | The phone number, SIP address, or Client identifier that made this call. Phone numbers are in [E.164 format](https://wwnw.twilio.com/docs/glossary/what-e164) (e.g., +16175551212). SIP addresses are formatted as &#x60;name@company.com&#x60;.
 **FallbackMethod** | **string** | The HTTP method that we should use to request the &#x60;fallback_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60; and the default is &#x60;POST&#x60;. If an &#x60;application_sid&#x60; parameter is present, this parameter is ignored.
 **FallbackUrl** | **string** | The URL that we call using the &#x60;fallback_method&#x60; if an error occurs when requesting or executing the TwiML at &#x60;url&#x60;. If an &#x60;application_sid&#x60; parameter is present, this parameter is ignored.
 **From** | **string** | The phone number or client identifier to use as the caller id. If using a phone number, it must be a Twilio number or a Verified [outgoing caller id](https://www.twilio.com/docs/voice/api/outgoing-caller-ids) for your account. If the &#x60;to&#x60; parameter is a phone number, &#x60;From&#x60; must also be a phone number.
 **MachineDetection** | **string** | Whether to detect if a human, answering machine, or fax has picked up the call. Can be: &#x60;Enable&#x60; or &#x60;DetectMessageEnd&#x60;. Use &#x60;Enable&#x60; if you would like us to return &#x60;AnsweredBy&#x60; as soon as the called party is identified. Use &#x60;DetectMessageEnd&#x60;, if you would like to leave a message on an answering machine. If &#x60;send_digits&#x60; is provided, this parameter is ignored. For more information, see [Answering Machine Detection](https://www.twilio.com/docs/voice/answering-machine-detection).
 **MachineDetectionSilenceTimeout** | **int32** | The number of milliseconds of initial silence after which an &#x60;unknown&#x60; AnsweredBy result will be returned. Possible Values: 2000-10000. Default: 5000.
 **MachineDetectionSpeechEndThreshold** | **int32** | The number of milliseconds of silence after speech activity at which point the speech activity is considered complete. Possible Values: 500-5000. Default: 1200.
 **MachineDetectionSpeechThreshold** | **int32** | The number of milliseconds that is used as the measuring stick for the length of the speech activity, where durations lower than this value will be interpreted as a human and longer than this value as a machine. Possible Values: 1000-6000. Default: 2400.
 **MachineDetectionTimeout** | **int32** | The number of seconds that we should attempt to detect an answering machine before timing out and sending a voice request with &#x60;AnsweredBy&#x60; of &#x60;unknown&#x60;. The default timeout is 30 seconds.
 **Method** | **string** | The HTTP method we should use when calling the &#x60;url&#x60; parameter&#39;s value. Can be: &#x60;GET&#x60; or &#x60;POST&#x60; and the default is &#x60;POST&#x60;. If an &#x60;application_sid&#x60; parameter is present, this parameter is ignored.
 **Record** | **bool** | Whether to record the call. Can be &#x60;true&#x60; to record the phone call, or &#x60;false&#x60; to not. The default is &#x60;false&#x60;. The &#x60;recording_url&#x60; is sent to the &#x60;status_callback&#x60; URL.
 **RecordingChannels** | **string** | The number of channels in the final recording. Can be: &#x60;mono&#x60; or &#x60;dual&#x60;. The default is &#x60;mono&#x60;. &#x60;mono&#x60; records both legs of the call in a single channel of the recording file. &#x60;dual&#x60; records each leg to a separate channel of the recording file. The first channel of a dual-channel recording contains the parent call and the second channel contains the child call.
 **RecordingStatusCallback** | **string** | The URL that we call when the recording is available to be accessed.
 **RecordingStatusCallbackEvent** | **[]string** | The recording status events that will trigger calls to the URL specified in &#x60;recording_status_callback&#x60;. Can be: &#x60;in-progress&#x60;, &#x60;completed&#x60; and &#x60;absent&#x60;. Defaults to &#x60;completed&#x60;. Separate  multiple values with a space.
 **RecordingStatusCallbackMethod** | **string** | The HTTP method we should use when calling the &#x60;recording_status_callback&#x60; URL. Can be: &#x60;GET&#x60; or &#x60;POST&#x60; and the default is &#x60;POST&#x60;.
 **RecordingTrack** | **string** | The audio track to record for the call. Can be: &#x60;inbound&#x60;, &#x60;outbound&#x60; or &#x60;both&#x60;. The default is &#x60;both&#x60;. &#x60;inbound&#x60; records the audio that is received by Twilio. &#x60;outbound&#x60; records the audio that is generated from Twilio. &#x60;both&#x60; records the audio that is received and generated by Twilio.
 **SendDigits** | **string** | A string of keys to dial after connecting to the number, maximum of 32 digits. Valid digits in the string include: any digit (&#x60;0&#x60;-&#x60;9&#x60;), &#39;&#x60;#&#x60;&#39;, &#39;&#x60;*&#x60;&#39; and &#39;&#x60;w&#x60;&#39;, to insert a half second pause. For example, if you connected to a company phone number and wanted to pause for one second, and then dial extension 1234 followed by the pound key, the value of this parameter would be &#x60;ww1234#&#x60;. Remember to URL-encode this string, since the &#39;&#x60;#&#x60;&#39; character has special meaning in a URL. If both &#x60;SendDigits&#x60; and &#x60;MachineDetection&#x60; parameters are provided, then &#x60;MachineDetection&#x60; will be ignored.
 **SipAuthPassword** | **string** | The password required to authenticate the user account specified in &#x60;sip_auth_username&#x60;.
 **SipAuthUsername** | **string** | The username used to authenticate the caller making a SIP call.
 **StatusCallback** | **string** | The URL we should call using the &#x60;status_callback_method&#x60; to send status information to your application. If no &#x60;status_callback_event&#x60; is specified, we will send the &#x60;completed&#x60; status. If an &#x60;application_sid&#x60; parameter is present, this parameter is ignored. URLs must contain a valid hostname (underscores are not permitted).
 **StatusCallbackEvent** | **[]string** | The call progress events that we will send to the &#x60;status_callback&#x60; URL. Can be: &#x60;initiated&#x60;, &#x60;ringing&#x60;, &#x60;answered&#x60;, and &#x60;completed&#x60;. If no event is specified, we send the &#x60;completed&#x60; status. If you want to receive multiple events, specify each one in a separate &#x60;status_callback_event&#x60; parameter. See the code sample for [monitoring call progress](https://www.twilio.com/docs/voice/api/call-resource?code-sample&#x3D;code-create-a-call-resource-and-specify-a-statuscallbackevent&amp;code-sdk-version&#x3D;json). If an &#x60;application_sid&#x60; is present, this parameter is ignored.
 **StatusCallbackMethod** | **string** | The HTTP method we should use when calling the &#x60;status_callback&#x60; URL. Can be: &#x60;GET&#x60; or &#x60;POST&#x60; and the default is &#x60;POST&#x60;. If an &#x60;application_sid&#x60; parameter is present, this parameter is ignored.
 **Timeout** | **int32** | The integer number of seconds that we should allow the phone to ring before assuming there is no answer. The default is &#x60;60&#x60; seconds and the maximum is &#x60;600&#x60; seconds. For some call flows, we will add a 5-second buffer to the timeout value you provide. For this reason, a timeout value of 10 seconds could result in an actual timeout closer to 15 seconds. You can set this to a short time, such as &#x60;15&#x60; seconds, to hang up before reaching an answering machine or voicemail.
 **To** | **string** | The phone number, SIP address, or client identifier to call.
 **Trim** | **string** | Whether to trim any leading and trailing silence from the recording. Can be: &#x60;trim-silence&#x60; or &#x60;do-not-trim&#x60; and the default is &#x60;trim-silence&#x60;.
 **Twiml** | **string** | TwiML instructions for the call Twilio will use without fetching Twiml from url parameter. If both &#x60;twiml&#x60; and &#x60;url&#x60; are provided then &#x60;twiml&#x60; parameter will be ignored.
 **Url** | **string** | The absolute URL that returns the TwiML instructions for the call. We will call this URL using the &#x60;method&#x60; when the call connects. For more information, see the [Url Parameter](https://www.twilio.com/docs/voice/make-calls#specify-a-url-parameter) section in [Making Calls](https://www.twilio.com/docs/voice/make-calls).

### Return type

[**ApiV2010AccountCall**](ApiV2010AccountCall.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCallFeedbackSummary

> ApiV2010AccountCallCallFeedbackSummary CreateCallFeedbackSummary(ctx, AccountSid).EndDate(EndDate).IncludeSubaccounts(IncludeSubaccounts).StartDate(StartDate).StatusCallback(StatusCallback).StatusCallbackMethod(StatusCallbackMethod).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    AccountSid := "AccountSid_example" // string | The unique id of the [Account](https://www.twilio.com/docs/iam/api/account) responsible for this resource.
    EndDate := time.Now() // string | Only include feedback given on or before this date. Format is `YYYY-MM-DD` and specified in UTC. (optional)
    IncludeSubaccounts := true // bool | Whether to also include Feedback resources from all subaccounts. `true` includes feedback from all subaccounts and `false`, the default, includes feedback from only the specified account. (optional)
    StartDate := time.Now() // string | Only include feedback given on or after this date. Format is `YYYY-MM-DD` and specified in UTC. (optional)
    StatusCallback := "StatusCallback_example" // string | The URL that we will request when the feedback summary is complete. (optional)
    StatusCallbackMethod := "StatusCallbackMethod_example" // string | The HTTP method (`GET` or `POST`) we use to make the request to the `StatusCallback` URL. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateCallFeedbackSummary(context.Background(), AccountSid).EndDate(EndDate).IncludeSubaccounts(IncludeSubaccounts).StartDate(StartDate).StatusCallback(StatusCallback).StatusCallbackMethod(StatusCallbackMethod).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateCallFeedbackSummary``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCallFeedbackSummary`: ApiV2010AccountCallCallFeedbackSummary
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateCallFeedbackSummary`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string** | The unique id of the [Account](https://www.twilio.com/docs/iam/api/account) responsible for this resource.

### Other Parameters

Other parameters are passed through a pointer to a CreateCallFeedbackSummaryParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **EndDate** | **string** | Only include feedback given on or before this date. Format is &#x60;YYYY-MM-DD&#x60; and specified in UTC.
 **IncludeSubaccounts** | **bool** | Whether to also include Feedback resources from all subaccounts. &#x60;true&#x60; includes feedback from all subaccounts and &#x60;false&#x60;, the default, includes feedback from only the specified account.
 **StartDate** | **string** | Only include feedback given on or after this date. Format is &#x60;YYYY-MM-DD&#x60; and specified in UTC.
 **StatusCallback** | **string** | The URL that we will request when the feedback summary is complete.
 **StatusCallbackMethod** | **string** | The HTTP method (&#x60;GET&#x60; or &#x60;POST&#x60;) we use to make the request to the &#x60;StatusCallback&#x60; URL.

### Return type

[**ApiV2010AccountCallCallFeedbackSummary**](ApiV2010AccountCallCallFeedbackSummary.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCallRecording

> ApiV2010AccountCallCallRecording CreateCallRecording(ctx, AccountSid, CallSid).RecordingChannels(RecordingChannels).RecordingStatusCallback(RecordingStatusCallback).RecordingStatusCallbackEvent(RecordingStatusCallbackEvent).RecordingStatusCallbackMethod(RecordingStatusCallbackMethod).RecordingTrack(RecordingTrack).Trim(Trim).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AccountSid := "AccountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that will create the resource.
    CallSid := "CallSid_example" // string | The SID of the [Call](https://www.twilio.com/docs/voice/api/call-resource) to associate the resource with.
    RecordingChannels := "RecordingChannels_example" // string | The number of channels used in the recording. Can be: `mono` or `dual` and the default is `mono`. `mono` records all parties of the call into one channel. `dual` records each party of a 2-party call into separate channels. (optional)
    RecordingStatusCallback := "RecordingStatusCallback_example" // string | The URL we should call using the `recording_status_callback_method` on each recording event specified in  `recording_status_callback_event`. For more information, see [RecordingStatusCallback parameters](https://www.twilio.com/docs/voice/api/recording#recordingstatuscallback). (optional)
    RecordingStatusCallbackEvent := []string{"Inner_example"} // []string | The recording status events on which we should call the `recording_status_callback` URL. Can be: `in-progress`, `completed` and `absent` and the default is `completed`. Separate multiple event values with a space. (optional)
    RecordingStatusCallbackMethod := "RecordingStatusCallbackMethod_example" // string | The HTTP method we should use to call `recording_status_callback`. Can be: `GET` or `POST` and the default is `POST`. (optional)
    RecordingTrack := "RecordingTrack_example" // string | The audio track to record for the call. Can be: `inbound`, `outbound` or `both`. The default is `both`. `inbound` records the audio that is received by Twilio. `outbound` records the audio that is generated from Twilio. `both` records the audio that is received and generated by Twilio. (optional)
    Trim := "Trim_example" // string | Whether to trim any leading and trailing silence in the recording. Can be: `trim-silence` or `do-not-trim` and the default is `do-not-trim`. `trim-silence` trims the silence from the beginning and end of the recording and `do-not-trim` does not. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateCallRecording(context.Background(), AccountSid, CallSid).RecordingChannels(RecordingChannels).RecordingStatusCallback(RecordingStatusCallback).RecordingStatusCallbackEvent(RecordingStatusCallbackEvent).RecordingStatusCallbackMethod(RecordingStatusCallbackMethod).RecordingTrack(RecordingTrack).Trim(Trim).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateCallRecording``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCallRecording`: ApiV2010AccountCallCallRecording
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateCallRecording`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that will create the resource.
**CallSid** | **string** | The SID of the [Call](https://www.twilio.com/docs/voice/api/call-resource) to associate the resource with.

### Other Parameters

Other parameters are passed through a pointer to a CreateCallRecordingParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


 **RecordingChannels** | **string** | The number of channels used in the recording. Can be: &#x60;mono&#x60; or &#x60;dual&#x60; and the default is &#x60;mono&#x60;. &#x60;mono&#x60; records all parties of the call into one channel. &#x60;dual&#x60; records each party of a 2-party call into separate channels.
 **RecordingStatusCallback** | **string** | The URL we should call using the &#x60;recording_status_callback_method&#x60; on each recording event specified in  &#x60;recording_status_callback_event&#x60;. For more information, see [RecordingStatusCallback parameters](https://www.twilio.com/docs/voice/api/recording#recordingstatuscallback).
 **RecordingStatusCallbackEvent** | **[]string** | The recording status events on which we should call the &#x60;recording_status_callback&#x60; URL. Can be: &#x60;in-progress&#x60;, &#x60;completed&#x60; and &#x60;absent&#x60; and the default is &#x60;completed&#x60;. Separate multiple event values with a space.
 **RecordingStatusCallbackMethod** | **string** | The HTTP method we should use to call &#x60;recording_status_callback&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60; and the default is &#x60;POST&#x60;.
 **RecordingTrack** | **string** | The audio track to record for the call. Can be: &#x60;inbound&#x60;, &#x60;outbound&#x60; or &#x60;both&#x60;. The default is &#x60;both&#x60;. &#x60;inbound&#x60; records the audio that is received by Twilio. &#x60;outbound&#x60; records the audio that is generated from Twilio. &#x60;both&#x60; records the audio that is received and generated by Twilio.
 **Trim** | **string** | Whether to trim any leading and trailing silence in the recording. Can be: &#x60;trim-silence&#x60; or &#x60;do-not-trim&#x60; and the default is &#x60;do-not-trim&#x60;. &#x60;trim-silence&#x60; trims the silence from the beginning and end of the recording and &#x60;do-not-trim&#x60; does not.

### Return type

[**ApiV2010AccountCallCallRecording**](ApiV2010AccountCallCallRecording.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateIncomingPhoneNumber

> ApiV2010AccountIncomingPhoneNumber CreateIncomingPhoneNumber(ctx, AccountSid).AddressSid(AddressSid).ApiVersion(ApiVersion).AreaCode(AreaCode).BundleSid(BundleSid).EmergencyAddressSid(EmergencyAddressSid).EmergencyStatus(EmergencyStatus).FriendlyName(FriendlyName).IdentitySid(IdentitySid).PhoneNumber(PhoneNumber).SmsApplicationSid(SmsApplicationSid).SmsFallbackMethod(SmsFallbackMethod).SmsFallbackUrl(SmsFallbackUrl).SmsMethod(SmsMethod).SmsUrl(SmsUrl).StatusCallback(StatusCallback).StatusCallbackMethod(StatusCallbackMethod).TrunkSid(TrunkSid).VoiceApplicationSid(VoiceApplicationSid).VoiceCallerIdLookup(VoiceCallerIdLookup).VoiceFallbackMethod(VoiceFallbackMethod).VoiceFallbackUrl(VoiceFallbackUrl).VoiceMethod(VoiceMethod).VoiceReceiveMode(VoiceReceiveMode).VoiceUrl(VoiceUrl).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AccountSid := "AccountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that will create the resource.
    AddressSid := "AddressSid_example" // string | The SID of the Address resource we should associate with the new phone number. Some regions require addresses to meet local regulations. (optional)
    ApiVersion := "ApiVersion_example" // string | The API version to use for incoming calls made to the new phone number. The default is `2010-04-01`. (optional)
    AreaCode := "AreaCode_example" // string | The desired area code for your new incoming phone number. Can be any three-digit, US or Canada area code. We will provision an available phone number within this area code for you. **You must provide an `area_code` or a `phone_number`.** (US and Canada only). (optional)
    BundleSid := "BundleSid_example" // string | The SID of the Bundle resource that you associate with the phone number. Some regions require a Bundle to meet local Regulations. (optional)
    EmergencyAddressSid := "EmergencyAddressSid_example" // string | The SID of the emergency address configuration to use for emergency calling from the new phone number. (optional)
    EmergencyStatus := "EmergencyStatus_example" // string | The configuration status parameter that determines whether the new phone number is enabled for emergency calling. (optional)
    FriendlyName := "FriendlyName_example" // string | A descriptive string that you created to describe the new phone number. It can be up to 64 characters long. By default, this is a formatted version of the new phone number. (optional)
    IdentitySid := "IdentitySid_example" // string | The SID of the Identity resource that we should associate with the new phone number. Some regions require an identity to meet local regulations. (optional)
    PhoneNumber := "PhoneNumber_example" // string | The phone number to purchase specified in [E.164](https://www.twilio.com/docs/glossary/what-e164) format.  E.164 phone numbers consist of a + followed by the country code and subscriber number without punctuation characters. For example, +14155551234. (optional)
    SmsApplicationSid := "SmsApplicationSid_example" // string | The SID of the application that should handle SMS messages sent to the new phone number. If an `sms_application_sid` is present, we ignore all of the `sms_*_url` urls and use those set on the application. (optional)
    SmsFallbackMethod := "SmsFallbackMethod_example" // string | The HTTP method that we should use to call `sms_fallback_url`. Can be: `GET` or `POST` and defaults to `POST`. (optional)
    SmsFallbackUrl := "SmsFallbackUrl_example" // string | The URL that we should call when an error occurs while requesting or executing the TwiML defined by `sms_url`. (optional)
    SmsMethod := "SmsMethod_example" // string | The HTTP method that we should use to call `sms_url`. Can be: `GET` or `POST` and defaults to `POST`. (optional)
    SmsUrl := "SmsUrl_example" // string | The URL we should call when the new phone number receives an incoming SMS message. (optional)
    StatusCallback := "StatusCallback_example" // string | The URL we should call using the `status_callback_method` to send status information to your application. (optional)
    StatusCallbackMethod := "StatusCallbackMethod_example" // string | The HTTP method we should use to call `status_callback`. Can be: `GET` or `POST` and defaults to `POST`. (optional)
    TrunkSid := "TrunkSid_example" // string | The SID of the Trunk we should use to handle calls to the new phone number. If a `trunk_sid` is present, we ignore all of the voice urls and voice applications and use only those set on the Trunk. Setting a `trunk_sid` will automatically delete your `voice_application_sid` and vice versa. (optional)
    VoiceApplicationSid := "VoiceApplicationSid_example" // string | The SID of the application we should use to handle calls to the new phone number. If a `voice_application_sid` is present, we ignore all of the voice urls and use only those set on the application. Setting a `voice_application_sid` will automatically delete your `trunk_sid` and vice versa. (optional)
    VoiceCallerIdLookup := true // bool | Whether to lookup the caller's name from the CNAM database and post it to your app. Can be: `true` or `false` and defaults to `false`. (optional)
    VoiceFallbackMethod := "VoiceFallbackMethod_example" // string | The HTTP method that we should use to call `voice_fallback_url`. Can be: `GET` or `POST` and defaults to `POST`. (optional)
    VoiceFallbackUrl := "VoiceFallbackUrl_example" // string | The URL that we should call when an error occurs retrieving or executing the TwiML requested by `url`. (optional)
    VoiceMethod := "VoiceMethod_example" // string | The HTTP method that we should use to call `voice_url`. Can be: `GET` or `POST` and defaults to `POST`. (optional)
    VoiceReceiveMode := "VoiceReceiveMode_example" // string | The configuration parameter for the new phone number to receive incoming voice calls or faxes. Can be: `fax` or `voice` and defaults to `voice`. (optional)
    VoiceUrl := "VoiceUrl_example" // string | The URL that we should call to answer a call to the new phone number. The `voice_url` will not be called if a `voice_application_sid` or a `trunk_sid` is set. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateIncomingPhoneNumber(context.Background(), AccountSid).AddressSid(AddressSid).ApiVersion(ApiVersion).AreaCode(AreaCode).BundleSid(BundleSid).EmergencyAddressSid(EmergencyAddressSid).EmergencyStatus(EmergencyStatus).FriendlyName(FriendlyName).IdentitySid(IdentitySid).PhoneNumber(PhoneNumber).SmsApplicationSid(SmsApplicationSid).SmsFallbackMethod(SmsFallbackMethod).SmsFallbackUrl(SmsFallbackUrl).SmsMethod(SmsMethod).SmsUrl(SmsUrl).StatusCallback(StatusCallback).StatusCallbackMethod(StatusCallbackMethod).TrunkSid(TrunkSid).VoiceApplicationSid(VoiceApplicationSid).VoiceCallerIdLookup(VoiceCallerIdLookup).VoiceFallbackMethod(VoiceFallbackMethod).VoiceFallbackUrl(VoiceFallbackUrl).VoiceMethod(VoiceMethod).VoiceReceiveMode(VoiceReceiveMode).VoiceUrl(VoiceUrl).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateIncomingPhoneNumber``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateIncomingPhoneNumber`: ApiV2010AccountIncomingPhoneNumber
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateIncomingPhoneNumber`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that will create the resource.

### Other Parameters

Other parameters are passed through a pointer to a CreateIncomingPhoneNumberParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **AddressSid** | **string** | The SID of the Address resource we should associate with the new phone number. Some regions require addresses to meet local regulations.
 **ApiVersion** | **string** | The API version to use for incoming calls made to the new phone number. The default is &#x60;2010-04-01&#x60;.
 **AreaCode** | **string** | The desired area code for your new incoming phone number. Can be any three-digit, US or Canada area code. We will provision an available phone number within this area code for you. **You must provide an &#x60;area_code&#x60; or a &#x60;phone_number&#x60;.** (US and Canada only).
 **BundleSid** | **string** | The SID of the Bundle resource that you associate with the phone number. Some regions require a Bundle to meet local Regulations.
 **EmergencyAddressSid** | **string** | The SID of the emergency address configuration to use for emergency calling from the new phone number.
 **EmergencyStatus** | **string** | The configuration status parameter that determines whether the new phone number is enabled for emergency calling.
 **FriendlyName** | **string** | A descriptive string that you created to describe the new phone number. It can be up to 64 characters long. By default, this is a formatted version of the new phone number.
 **IdentitySid** | **string** | The SID of the Identity resource that we should associate with the new phone number. Some regions require an identity to meet local regulations.
 **PhoneNumber** | **string** | The phone number to purchase specified in [E.164](https://www.twilio.com/docs/glossary/what-e164) format.  E.164 phone numbers consist of a + followed by the country code and subscriber number without punctuation characters. For example, +14155551234.
 **SmsApplicationSid** | **string** | The SID of the application that should handle SMS messages sent to the new phone number. If an &#x60;sms_application_sid&#x60; is present, we ignore all of the &#x60;sms_*_url&#x60; urls and use those set on the application.
 **SmsFallbackMethod** | **string** | The HTTP method that we should use to call &#x60;sms_fallback_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60; and defaults to &#x60;POST&#x60;.
 **SmsFallbackUrl** | **string** | The URL that we should call when an error occurs while requesting or executing the TwiML defined by &#x60;sms_url&#x60;.
 **SmsMethod** | **string** | The HTTP method that we should use to call &#x60;sms_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60; and defaults to &#x60;POST&#x60;.
 **SmsUrl** | **string** | The URL we should call when the new phone number receives an incoming SMS message.
 **StatusCallback** | **string** | The URL we should call using the &#x60;status_callback_method&#x60; to send status information to your application.
 **StatusCallbackMethod** | **string** | The HTTP method we should use to call &#x60;status_callback&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60; and defaults to &#x60;POST&#x60;.
 **TrunkSid** | **string** | The SID of the Trunk we should use to handle calls to the new phone number. If a &#x60;trunk_sid&#x60; is present, we ignore all of the voice urls and voice applications and use only those set on the Trunk. Setting a &#x60;trunk_sid&#x60; will automatically delete your &#x60;voice_application_sid&#x60; and vice versa.
 **VoiceApplicationSid** | **string** | The SID of the application we should use to handle calls to the new phone number. If a &#x60;voice_application_sid&#x60; is present, we ignore all of the voice urls and use only those set on the application. Setting a &#x60;voice_application_sid&#x60; will automatically delete your &#x60;trunk_sid&#x60; and vice versa.
 **VoiceCallerIdLookup** | **bool** | Whether to lookup the caller&#39;s name from the CNAM database and post it to your app. Can be: &#x60;true&#x60; or &#x60;false&#x60; and defaults to &#x60;false&#x60;.
 **VoiceFallbackMethod** | **string** | The HTTP method that we should use to call &#x60;voice_fallback_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60; and defaults to &#x60;POST&#x60;.
 **VoiceFallbackUrl** | **string** | The URL that we should call when an error occurs retrieving or executing the TwiML requested by &#x60;url&#x60;.
 **VoiceMethod** | **string** | The HTTP method that we should use to call &#x60;voice_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60; and defaults to &#x60;POST&#x60;.
 **VoiceReceiveMode** | **string** | The configuration parameter for the new phone number to receive incoming voice calls or faxes. Can be: &#x60;fax&#x60; or &#x60;voice&#x60; and defaults to &#x60;voice&#x60;.
 **VoiceUrl** | **string** | The URL that we should call to answer a call to the new phone number. The &#x60;voice_url&#x60; will not be called if a &#x60;voice_application_sid&#x60; or a &#x60;trunk_sid&#x60; is set.

### Return type

[**ApiV2010AccountIncomingPhoneNumber**](ApiV2010AccountIncomingPhoneNumber.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateIncomingPhoneNumberAssignedAddOn

> ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberAssignedAddOn CreateIncomingPhoneNumberAssignedAddOn(ctx, AccountSid, ResourceSid).InstalledAddOnSid(InstalledAddOnSid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AccountSid := "AccountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that will create the resource.
    ResourceSid := "ResourceSid_example" // string | The SID of the Phone Number to assign the Add-on.
    InstalledAddOnSid := "InstalledAddOnSid_example" // string | The SID that identifies the Add-on installation. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateIncomingPhoneNumberAssignedAddOn(context.Background(), AccountSid, ResourceSid).InstalledAddOnSid(InstalledAddOnSid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateIncomingPhoneNumberAssignedAddOn``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateIncomingPhoneNumberAssignedAddOn`: ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberAssignedAddOn
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateIncomingPhoneNumberAssignedAddOn`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that will create the resource.
**ResourceSid** | **string** | The SID of the Phone Number to assign the Add-on.

### Other Parameters

Other parameters are passed through a pointer to a CreateIncomingPhoneNumberAssignedAddOnParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


 **InstalledAddOnSid** | **string** | The SID that identifies the Add-on installation.

### Return type

[**ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberAssignedAddOn**](ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberAssignedAddOn.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateIncomingPhoneNumberLocal

> ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberLocal CreateIncomingPhoneNumberLocal(ctx, AccountSid).AddressSid(AddressSid).ApiVersion(ApiVersion).BundleSid(BundleSid).EmergencyAddressSid(EmergencyAddressSid).EmergencyStatus(EmergencyStatus).FriendlyName(FriendlyName).IdentitySid(IdentitySid).PhoneNumber(PhoneNumber).SmsApplicationSid(SmsApplicationSid).SmsFallbackMethod(SmsFallbackMethod).SmsFallbackUrl(SmsFallbackUrl).SmsMethod(SmsMethod).SmsUrl(SmsUrl).StatusCallback(StatusCallback).StatusCallbackMethod(StatusCallbackMethod).TrunkSid(TrunkSid).VoiceApplicationSid(VoiceApplicationSid).VoiceCallerIdLookup(VoiceCallerIdLookup).VoiceFallbackMethod(VoiceFallbackMethod).VoiceFallbackUrl(VoiceFallbackUrl).VoiceMethod(VoiceMethod).VoiceReceiveMode(VoiceReceiveMode).VoiceUrl(VoiceUrl).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AccountSid := "AccountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that will create the resource.
    AddressSid := "AddressSid_example" // string | The SID of the Address resource we should associate with the new phone number. Some regions require addresses to meet local regulations. (optional)
    ApiVersion := "ApiVersion_example" // string | The API version to use for incoming calls made to the new phone number. The default is `2010-04-01`. (optional)
    BundleSid := "BundleSid_example" // string | The SID of the Bundle resource that you associate with the phone number. Some regions require a Bundle to meet local Regulations. (optional)
    EmergencyAddressSid := "EmergencyAddressSid_example" // string | The SID of the emergency address configuration to use for emergency calling from the new phone number. (optional)
    EmergencyStatus := "EmergencyStatus_example" // string | The configuration status parameter that determines whether the new phone number is enabled for emergency calling. (optional)
    FriendlyName := "FriendlyName_example" // string | A descriptive string that you created to describe the new phone number. It can be up to 64 characters long. By default, this is a formatted version of the phone number. (optional)
    IdentitySid := "IdentitySid_example" // string | The SID of the Identity resource that we should associate with the new phone number. Some regions require an identity to meet local regulations. (optional)
    PhoneNumber := "PhoneNumber_example" // string | The phone number to purchase specified in [E.164](https://www.twilio.com/docs/glossary/what-e164) format.  E.164 phone numbers consist of a + followed by the country code and subscriber number without punctuation characters. For example, +14155551234. (optional)
    SmsApplicationSid := "SmsApplicationSid_example" // string | The SID of the application that should handle SMS messages sent to the new phone number. If an `sms_application_sid` is present, we ignore all of the `sms_*_url` urls and use those set on the application. (optional)
    SmsFallbackMethod := "SmsFallbackMethod_example" // string | The HTTP method that we should use to call `sms_fallback_url`. Can be: `GET` or `POST` and defaults to `POST`. (optional)
    SmsFallbackUrl := "SmsFallbackUrl_example" // string | The URL that we should call when an error occurs while requesting or executing the TwiML defined by `sms_url`. (optional)
    SmsMethod := "SmsMethod_example" // string | The HTTP method that we should use to call `sms_url`. Can be: `GET` or `POST` and defaults to `POST`. (optional)
    SmsUrl := "SmsUrl_example" // string | The URL we should call when the new phone number receives an incoming SMS message. (optional)
    StatusCallback := "StatusCallback_example" // string | The URL we should call using the `status_callback_method` to send status information to your application. (optional)
    StatusCallbackMethod := "StatusCallbackMethod_example" // string | The HTTP method we should use to call `status_callback`. Can be: `GET` or `POST` and defaults to `POST`. (optional)
    TrunkSid := "TrunkSid_example" // string | The SID of the Trunk we should use to handle calls to the new phone number. If a `trunk_sid` is present, we ignore all of the voice urls and voice applications and use only those set on the Trunk. Setting a `trunk_sid` will automatically delete your `voice_application_sid` and vice versa. (optional)
    VoiceApplicationSid := "VoiceApplicationSid_example" // string | The SID of the application we should use to handle calls to the new phone number. If a `voice_application_sid` is present, we ignore all of the voice urls and use only those set on the application. Setting a `voice_application_sid` will automatically delete your `trunk_sid` and vice versa. (optional)
    VoiceCallerIdLookup := true // bool | Whether to lookup the caller's name from the CNAM database and post it to your app. Can be: `true` or `false` and defaults to `false`. (optional)
    VoiceFallbackMethod := "VoiceFallbackMethod_example" // string | The HTTP method that we should use to call `voice_fallback_url`. Can be: `GET` or `POST` and defaults to `POST`. (optional)
    VoiceFallbackUrl := "VoiceFallbackUrl_example" // string | The URL that we should call when an error occurs retrieving or executing the TwiML requested by `url`. (optional)
    VoiceMethod := "VoiceMethod_example" // string | The HTTP method that we should use to call `voice_url`. Can be: `GET` or `POST` and defaults to `POST`. (optional)
    VoiceReceiveMode := "VoiceReceiveMode_example" // string | The configuration parameter for the new phone number to receive incoming voice calls or faxes. Can be: `fax` or `voice` and defaults to `voice`. (optional)
    VoiceUrl := "VoiceUrl_example" // string | The URL that we should call to answer a call to the new phone number. The `voice_url` will not be called if a `voice_application_sid` or a `trunk_sid` is set. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateIncomingPhoneNumberLocal(context.Background(), AccountSid).AddressSid(AddressSid).ApiVersion(ApiVersion).BundleSid(BundleSid).EmergencyAddressSid(EmergencyAddressSid).EmergencyStatus(EmergencyStatus).FriendlyName(FriendlyName).IdentitySid(IdentitySid).PhoneNumber(PhoneNumber).SmsApplicationSid(SmsApplicationSid).SmsFallbackMethod(SmsFallbackMethod).SmsFallbackUrl(SmsFallbackUrl).SmsMethod(SmsMethod).SmsUrl(SmsUrl).StatusCallback(StatusCallback).StatusCallbackMethod(StatusCallbackMethod).TrunkSid(TrunkSid).VoiceApplicationSid(VoiceApplicationSid).VoiceCallerIdLookup(VoiceCallerIdLookup).VoiceFallbackMethod(VoiceFallbackMethod).VoiceFallbackUrl(VoiceFallbackUrl).VoiceMethod(VoiceMethod).VoiceReceiveMode(VoiceReceiveMode).VoiceUrl(VoiceUrl).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateIncomingPhoneNumberLocal``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateIncomingPhoneNumberLocal`: ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberLocal
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateIncomingPhoneNumberLocal`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that will create the resource.

### Other Parameters

Other parameters are passed through a pointer to a CreateIncomingPhoneNumberLocalParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **AddressSid** | **string** | The SID of the Address resource we should associate with the new phone number. Some regions require addresses to meet local regulations.
 **ApiVersion** | **string** | The API version to use for incoming calls made to the new phone number. The default is &#x60;2010-04-01&#x60;.
 **BundleSid** | **string** | The SID of the Bundle resource that you associate with the phone number. Some regions require a Bundle to meet local Regulations.
 **EmergencyAddressSid** | **string** | The SID of the emergency address configuration to use for emergency calling from the new phone number.
 **EmergencyStatus** | **string** | The configuration status parameter that determines whether the new phone number is enabled for emergency calling.
 **FriendlyName** | **string** | A descriptive string that you created to describe the new phone number. It can be up to 64 characters long. By default, this is a formatted version of the phone number.
 **IdentitySid** | **string** | The SID of the Identity resource that we should associate with the new phone number. Some regions require an identity to meet local regulations.
 **PhoneNumber** | **string** | The phone number to purchase specified in [E.164](https://www.twilio.com/docs/glossary/what-e164) format.  E.164 phone numbers consist of a + followed by the country code and subscriber number without punctuation characters. For example, +14155551234.
 **SmsApplicationSid** | **string** | The SID of the application that should handle SMS messages sent to the new phone number. If an &#x60;sms_application_sid&#x60; is present, we ignore all of the &#x60;sms_*_url&#x60; urls and use those set on the application.
 **SmsFallbackMethod** | **string** | The HTTP method that we should use to call &#x60;sms_fallback_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60; and defaults to &#x60;POST&#x60;.
 **SmsFallbackUrl** | **string** | The URL that we should call when an error occurs while requesting or executing the TwiML defined by &#x60;sms_url&#x60;.
 **SmsMethod** | **string** | The HTTP method that we should use to call &#x60;sms_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60; and defaults to &#x60;POST&#x60;.
 **SmsUrl** | **string** | The URL we should call when the new phone number receives an incoming SMS message.
 **StatusCallback** | **string** | The URL we should call using the &#x60;status_callback_method&#x60; to send status information to your application.
 **StatusCallbackMethod** | **string** | The HTTP method we should use to call &#x60;status_callback&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60; and defaults to &#x60;POST&#x60;.
 **TrunkSid** | **string** | The SID of the Trunk we should use to handle calls to the new phone number. If a &#x60;trunk_sid&#x60; is present, we ignore all of the voice urls and voice applications and use only those set on the Trunk. Setting a &#x60;trunk_sid&#x60; will automatically delete your &#x60;voice_application_sid&#x60; and vice versa.
 **VoiceApplicationSid** | **string** | The SID of the application we should use to handle calls to the new phone number. If a &#x60;voice_application_sid&#x60; is present, we ignore all of the voice urls and use only those set on the application. Setting a &#x60;voice_application_sid&#x60; will automatically delete your &#x60;trunk_sid&#x60; and vice versa.
 **VoiceCallerIdLookup** | **bool** | Whether to lookup the caller&#39;s name from the CNAM database and post it to your app. Can be: &#x60;true&#x60; or &#x60;false&#x60; and defaults to &#x60;false&#x60;.
 **VoiceFallbackMethod** | **string** | The HTTP method that we should use to call &#x60;voice_fallback_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60; and defaults to &#x60;POST&#x60;.
 **VoiceFallbackUrl** | **string** | The URL that we should call when an error occurs retrieving or executing the TwiML requested by &#x60;url&#x60;.
 **VoiceMethod** | **string** | The HTTP method that we should use to call &#x60;voice_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60; and defaults to &#x60;POST&#x60;.
 **VoiceReceiveMode** | **string** | The configuration parameter for the new phone number to receive incoming voice calls or faxes. Can be: &#x60;fax&#x60; or &#x60;voice&#x60; and defaults to &#x60;voice&#x60;.
 **VoiceUrl** | **string** | The URL that we should call to answer a call to the new phone number. The &#x60;voice_url&#x60; will not be called if a &#x60;voice_application_sid&#x60; or a &#x60;trunk_sid&#x60; is set.

### Return type

[**ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberLocal**](ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberLocal.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateIncomingPhoneNumberMobile

> ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile CreateIncomingPhoneNumberMobile(ctx, AccountSid).AddressSid(AddressSid).ApiVersion(ApiVersion).BundleSid(BundleSid).EmergencyAddressSid(EmergencyAddressSid).EmergencyStatus(EmergencyStatus).FriendlyName(FriendlyName).IdentitySid(IdentitySid).PhoneNumber(PhoneNumber).SmsApplicationSid(SmsApplicationSid).SmsFallbackMethod(SmsFallbackMethod).SmsFallbackUrl(SmsFallbackUrl).SmsMethod(SmsMethod).SmsUrl(SmsUrl).StatusCallback(StatusCallback).StatusCallbackMethod(StatusCallbackMethod).TrunkSid(TrunkSid).VoiceApplicationSid(VoiceApplicationSid).VoiceCallerIdLookup(VoiceCallerIdLookup).VoiceFallbackMethod(VoiceFallbackMethod).VoiceFallbackUrl(VoiceFallbackUrl).VoiceMethod(VoiceMethod).VoiceReceiveMode(VoiceReceiveMode).VoiceUrl(VoiceUrl).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AccountSid := "AccountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that will create the resource.
    AddressSid := "AddressSid_example" // string | The SID of the Address resource we should associate with the new phone number. Some regions require addresses to meet local regulations. (optional)
    ApiVersion := "ApiVersion_example" // string | The API version to use for incoming calls made to the new phone number. The default is `2010-04-01`. (optional)
    BundleSid := "BundleSid_example" // string | The SID of the Bundle resource that you associate with the phone number. Some regions require a Bundle to meet local Regulations. (optional)
    EmergencyAddressSid := "EmergencyAddressSid_example" // string | The SID of the emergency address configuration to use for emergency calling from the new phone number. (optional)
    EmergencyStatus := "EmergencyStatus_example" // string | The configuration status parameter that determines whether the new phone number is enabled for emergency calling. (optional)
    FriendlyName := "FriendlyName_example" // string | A descriptive string that you created to describe the new phone number. It can be up to 64 characters long. By default, the is a formatted version of the phone number. (optional)
    IdentitySid := "IdentitySid_example" // string | The SID of the Identity resource that we should associate with the new phone number. Some regions require an identity to meet local regulations. (optional)
    PhoneNumber := "PhoneNumber_example" // string | The phone number to purchase specified in [E.164](https://www.twilio.com/docs/glossary/what-e164) format.  E.164 phone numbers consist of a + followed by the country code and subscriber number without punctuation characters. For example, +14155551234. (optional)
    SmsApplicationSid := "SmsApplicationSid_example" // string | The SID of the application that should handle SMS messages sent to the new phone number. If an `sms_application_sid` is present, we ignore all of the `sms_*_url` urls and use those of the application. (optional)
    SmsFallbackMethod := "SmsFallbackMethod_example" // string | The HTTP method that we should use to call `sms_fallback_url`. Can be: `GET` or `POST` and defaults to `POST`. (optional)
    SmsFallbackUrl := "SmsFallbackUrl_example" // string | The URL that we should call when an error occurs while requesting or executing the TwiML defined by `sms_url`. (optional)
    SmsMethod := "SmsMethod_example" // string | The HTTP method that we should use to call `sms_url`. Can be: `GET` or `POST` and defaults to `POST`. (optional)
    SmsUrl := "SmsUrl_example" // string | The URL we should call when the new phone number receives an incoming SMS message. (optional)
    StatusCallback := "StatusCallback_example" // string | The URL we should call using the `status_callback_method` to send status information to your application. (optional)
    StatusCallbackMethod := "StatusCallbackMethod_example" // string | The HTTP method we should use to call `status_callback`. Can be: `GET` or `POST` and defaults to `POST`. (optional)
    TrunkSid := "TrunkSid_example" // string | The SID of the Trunk we should use to handle calls to the new phone number. If a `trunk_sid` is present, we ignore all of the voice urls and voice applications and use only those set on the Trunk. Setting a `trunk_sid` will automatically delete your `voice_application_sid` and vice versa. (optional)
    VoiceApplicationSid := "VoiceApplicationSid_example" // string | The SID of the application we should use to handle calls to the new phone number. If a `voice_application_sid` is present, we ignore all of the voice urls and use only those set on the application. Setting a `voice_application_sid` will automatically delete your `trunk_sid` and vice versa. (optional)
    VoiceCallerIdLookup := true // bool | Whether to lookup the caller's name from the CNAM database and post it to your app. Can be: `true` or `false` and defaults to `false`. (optional)
    VoiceFallbackMethod := "VoiceFallbackMethod_example" // string | The HTTP method that we should use to call `voice_fallback_url`. Can be: `GET` or `POST` and defaults to `POST`. (optional)
    VoiceFallbackUrl := "VoiceFallbackUrl_example" // string | The URL that we should call when an error occurs retrieving or executing the TwiML requested by `url`. (optional)
    VoiceMethod := "VoiceMethod_example" // string | The HTTP method that we should use to call `voice_url`. Can be: `GET` or `POST` and defaults to `POST`. (optional)
    VoiceReceiveMode := "VoiceReceiveMode_example" // string | The configuration parameter for the new phone number to receive incoming voice calls or faxes. Can be: `fax` or `voice` and defaults to `voice`. (optional)
    VoiceUrl := "VoiceUrl_example" // string | The URL that we should call to answer a call to the new phone number. The `voice_url` will not be called if a `voice_application_sid` or a `trunk_sid` is set. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateIncomingPhoneNumberMobile(context.Background(), AccountSid).AddressSid(AddressSid).ApiVersion(ApiVersion).BundleSid(BundleSid).EmergencyAddressSid(EmergencyAddressSid).EmergencyStatus(EmergencyStatus).FriendlyName(FriendlyName).IdentitySid(IdentitySid).PhoneNumber(PhoneNumber).SmsApplicationSid(SmsApplicationSid).SmsFallbackMethod(SmsFallbackMethod).SmsFallbackUrl(SmsFallbackUrl).SmsMethod(SmsMethod).SmsUrl(SmsUrl).StatusCallback(StatusCallback).StatusCallbackMethod(StatusCallbackMethod).TrunkSid(TrunkSid).VoiceApplicationSid(VoiceApplicationSid).VoiceCallerIdLookup(VoiceCallerIdLookup).VoiceFallbackMethod(VoiceFallbackMethod).VoiceFallbackUrl(VoiceFallbackUrl).VoiceMethod(VoiceMethod).VoiceReceiveMode(VoiceReceiveMode).VoiceUrl(VoiceUrl).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateIncomingPhoneNumberMobile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateIncomingPhoneNumberMobile`: ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateIncomingPhoneNumberMobile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that will create the resource.

### Other Parameters

Other parameters are passed through a pointer to a CreateIncomingPhoneNumberMobileParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **AddressSid** | **string** | The SID of the Address resource we should associate with the new phone number. Some regions require addresses to meet local regulations.
 **ApiVersion** | **string** | The API version to use for incoming calls made to the new phone number. The default is &#x60;2010-04-01&#x60;.
 **BundleSid** | **string** | The SID of the Bundle resource that you associate with the phone number. Some regions require a Bundle to meet local Regulations.
 **EmergencyAddressSid** | **string** | The SID of the emergency address configuration to use for emergency calling from the new phone number.
 **EmergencyStatus** | **string** | The configuration status parameter that determines whether the new phone number is enabled for emergency calling.
 **FriendlyName** | **string** | A descriptive string that you created to describe the new phone number. It can be up to 64 characters long. By default, the is a formatted version of the phone number.
 **IdentitySid** | **string** | The SID of the Identity resource that we should associate with the new phone number. Some regions require an identity to meet local regulations.
 **PhoneNumber** | **string** | The phone number to purchase specified in [E.164](https://www.twilio.com/docs/glossary/what-e164) format.  E.164 phone numbers consist of a + followed by the country code and subscriber number without punctuation characters. For example, +14155551234.
 **SmsApplicationSid** | **string** | The SID of the application that should handle SMS messages sent to the new phone number. If an &#x60;sms_application_sid&#x60; is present, we ignore all of the &#x60;sms_*_url&#x60; urls and use those of the application.
 **SmsFallbackMethod** | **string** | The HTTP method that we should use to call &#x60;sms_fallback_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60; and defaults to &#x60;POST&#x60;.
 **SmsFallbackUrl** | **string** | The URL that we should call when an error occurs while requesting or executing the TwiML defined by &#x60;sms_url&#x60;.
 **SmsMethod** | **string** | The HTTP method that we should use to call &#x60;sms_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60; and defaults to &#x60;POST&#x60;.
 **SmsUrl** | **string** | The URL we should call when the new phone number receives an incoming SMS message.
 **StatusCallback** | **string** | The URL we should call using the &#x60;status_callback_method&#x60; to send status information to your application.
 **StatusCallbackMethod** | **string** | The HTTP method we should use to call &#x60;status_callback&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60; and defaults to &#x60;POST&#x60;.
 **TrunkSid** | **string** | The SID of the Trunk we should use to handle calls to the new phone number. If a &#x60;trunk_sid&#x60; is present, we ignore all of the voice urls and voice applications and use only those set on the Trunk. Setting a &#x60;trunk_sid&#x60; will automatically delete your &#x60;voice_application_sid&#x60; and vice versa.
 **VoiceApplicationSid** | **string** | The SID of the application we should use to handle calls to the new phone number. If a &#x60;voice_application_sid&#x60; is present, we ignore all of the voice urls and use only those set on the application. Setting a &#x60;voice_application_sid&#x60; will automatically delete your &#x60;trunk_sid&#x60; and vice versa.
 **VoiceCallerIdLookup** | **bool** | Whether to lookup the caller&#39;s name from the CNAM database and post it to your app. Can be: &#x60;true&#x60; or &#x60;false&#x60; and defaults to &#x60;false&#x60;.
 **VoiceFallbackMethod** | **string** | The HTTP method that we should use to call &#x60;voice_fallback_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60; and defaults to &#x60;POST&#x60;.
 **VoiceFallbackUrl** | **string** | The URL that we should call when an error occurs retrieving or executing the TwiML requested by &#x60;url&#x60;.
 **VoiceMethod** | **string** | The HTTP method that we should use to call &#x60;voice_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60; and defaults to &#x60;POST&#x60;.
 **VoiceReceiveMode** | **string** | The configuration parameter for the new phone number to receive incoming voice calls or faxes. Can be: &#x60;fax&#x60; or &#x60;voice&#x60; and defaults to &#x60;voice&#x60;.
 **VoiceUrl** | **string** | The URL that we should call to answer a call to the new phone number. The &#x60;voice_url&#x60; will not be called if a &#x60;voice_application_sid&#x60; or a &#x60;trunk_sid&#x60; is set.

### Return type

[**ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile**](ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateIncomingPhoneNumberTollFree

> ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree CreateIncomingPhoneNumberTollFree(ctx, AccountSid).AddressSid(AddressSid).ApiVersion(ApiVersion).BundleSid(BundleSid).EmergencyAddressSid(EmergencyAddressSid).EmergencyStatus(EmergencyStatus).FriendlyName(FriendlyName).IdentitySid(IdentitySid).PhoneNumber(PhoneNumber).SmsApplicationSid(SmsApplicationSid).SmsFallbackMethod(SmsFallbackMethod).SmsFallbackUrl(SmsFallbackUrl).SmsMethod(SmsMethod).SmsUrl(SmsUrl).StatusCallback(StatusCallback).StatusCallbackMethod(StatusCallbackMethod).TrunkSid(TrunkSid).VoiceApplicationSid(VoiceApplicationSid).VoiceCallerIdLookup(VoiceCallerIdLookup).VoiceFallbackMethod(VoiceFallbackMethod).VoiceFallbackUrl(VoiceFallbackUrl).VoiceMethod(VoiceMethod).VoiceReceiveMode(VoiceReceiveMode).VoiceUrl(VoiceUrl).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AccountSid := "AccountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that will create the resource.
    AddressSid := "AddressSid_example" // string | The SID of the Address resource we should associate with the new phone number. Some regions require addresses to meet local regulations. (optional)
    ApiVersion := "ApiVersion_example" // string | The API version to use for incoming calls made to the new phone number. The default is `2010-04-01`. (optional)
    BundleSid := "BundleSid_example" // string | The SID of the Bundle resource that you associate with the phone number. Some regions require a Bundle to meet local Regulations. (optional)
    EmergencyAddressSid := "EmergencyAddressSid_example" // string | The SID of the emergency address configuration to use for emergency calling from the new phone number. (optional)
    EmergencyStatus := "EmergencyStatus_example" // string | The configuration status parameter that determines whether the new phone number is enabled for emergency calling. (optional)
    FriendlyName := "FriendlyName_example" // string | A descriptive string that you created to describe the new phone number. It can be up to 64 characters long. By default, this is a formatted version of the phone number. (optional)
    IdentitySid := "IdentitySid_example" // string | The SID of the Identity resource that we should associate with the new phone number. Some regions require an Identity to meet local regulations. (optional)
    PhoneNumber := "PhoneNumber_example" // string | The phone number to purchase specified in [E.164](https://www.twilio.com/docs/glossary/what-e164) format.  E.164 phone numbers consist of a + followed by the country code and subscriber number without punctuation characters. For example, +14155551234. (optional)
    SmsApplicationSid := "SmsApplicationSid_example" // string | The SID of the application that should handle SMS messages sent to the new phone number. If an `sms_application_sid` is present, we ignore all `sms_*_url` values and use those of the application. (optional)
    SmsFallbackMethod := "SmsFallbackMethod_example" // string | The HTTP method that we should use to call `sms_fallback_url`. Can be: `GET` or `POST` and defaults to `POST`. (optional)
    SmsFallbackUrl := "SmsFallbackUrl_example" // string | The URL that we should call when an error occurs while requesting or executing the TwiML defined by `sms_url`. (optional)
    SmsMethod := "SmsMethod_example" // string | The HTTP method that we should use to call `sms_url`. Can be: `GET` or `POST` and defaults to `POST`. (optional)
    SmsUrl := "SmsUrl_example" // string | The URL we should call when the new phone number receives an incoming SMS message. (optional)
    StatusCallback := "StatusCallback_example" // string | The URL we should call using the `status_callback_method` to send status information to your application. (optional)
    StatusCallbackMethod := "StatusCallbackMethod_example" // string | The HTTP method we should use to call `status_callback`. Can be: `GET` or `POST` and defaults to `POST`. (optional)
    TrunkSid := "TrunkSid_example" // string | The SID of the Trunk we should use to handle calls to the new phone number. If a `trunk_sid` is present, we ignore all of the voice urls and voice applications and use only those set on the Trunk. Setting a `trunk_sid` will automatically delete your `voice_application_sid` and vice versa. (optional)
    VoiceApplicationSid := "VoiceApplicationSid_example" // string | The SID of the application we should use to handle calls to the new phone number. If a `voice_application_sid` is present, we ignore all of the voice urls and use those set on the application. Setting a `voice_application_sid` will automatically delete your `trunk_sid` and vice versa. (optional)
    VoiceCallerIdLookup := true // bool | Whether to lookup the caller's name from the CNAM database and post it to your app. Can be: `true` or `false` and defaults to `false`. (optional)
    VoiceFallbackMethod := "VoiceFallbackMethod_example" // string | The HTTP method that we should use to call `voice_fallback_url`. Can be: `GET` or `POST` and defaults to `POST`. (optional)
    VoiceFallbackUrl := "VoiceFallbackUrl_example" // string | The URL that we should call when an error occurs retrieving or executing the TwiML requested by `url`. (optional)
    VoiceMethod := "VoiceMethod_example" // string | The HTTP method that we should use to call `voice_url`. Can be: `GET` or `POST` and defaults to `POST`. (optional)
    VoiceReceiveMode := "VoiceReceiveMode_example" // string | The configuration parameter for the new phone number to receive incoming voice calls or faxes. Can be: `fax` or `voice` and defaults to `voice`. (optional)
    VoiceUrl := "VoiceUrl_example" // string | The URL that we should call to answer a call to the new phone number. The `voice_url` will not be called if a `voice_application_sid` or a `trunk_sid` is set. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateIncomingPhoneNumberTollFree(context.Background(), AccountSid).AddressSid(AddressSid).ApiVersion(ApiVersion).BundleSid(BundleSid).EmergencyAddressSid(EmergencyAddressSid).EmergencyStatus(EmergencyStatus).FriendlyName(FriendlyName).IdentitySid(IdentitySid).PhoneNumber(PhoneNumber).SmsApplicationSid(SmsApplicationSid).SmsFallbackMethod(SmsFallbackMethod).SmsFallbackUrl(SmsFallbackUrl).SmsMethod(SmsMethod).SmsUrl(SmsUrl).StatusCallback(StatusCallback).StatusCallbackMethod(StatusCallbackMethod).TrunkSid(TrunkSid).VoiceApplicationSid(VoiceApplicationSid).VoiceCallerIdLookup(VoiceCallerIdLookup).VoiceFallbackMethod(VoiceFallbackMethod).VoiceFallbackUrl(VoiceFallbackUrl).VoiceMethod(VoiceMethod).VoiceReceiveMode(VoiceReceiveMode).VoiceUrl(VoiceUrl).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateIncomingPhoneNumberTollFree``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateIncomingPhoneNumberTollFree`: ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateIncomingPhoneNumberTollFree`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that will create the resource.

### Other Parameters

Other parameters are passed through a pointer to a CreateIncomingPhoneNumberTollFreeParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **AddressSid** | **string** | The SID of the Address resource we should associate with the new phone number. Some regions require addresses to meet local regulations.
 **ApiVersion** | **string** | The API version to use for incoming calls made to the new phone number. The default is &#x60;2010-04-01&#x60;.
 **BundleSid** | **string** | The SID of the Bundle resource that you associate with the phone number. Some regions require a Bundle to meet local Regulations.
 **EmergencyAddressSid** | **string** | The SID of the emergency address configuration to use for emergency calling from the new phone number.
 **EmergencyStatus** | **string** | The configuration status parameter that determines whether the new phone number is enabled for emergency calling.
 **FriendlyName** | **string** | A descriptive string that you created to describe the new phone number. It can be up to 64 characters long. By default, this is a formatted version of the phone number.
 **IdentitySid** | **string** | The SID of the Identity resource that we should associate with the new phone number. Some regions require an Identity to meet local regulations.
 **PhoneNumber** | **string** | The phone number to purchase specified in [E.164](https://www.twilio.com/docs/glossary/what-e164) format.  E.164 phone numbers consist of a + followed by the country code and subscriber number without punctuation characters. For example, +14155551234.
 **SmsApplicationSid** | **string** | The SID of the application that should handle SMS messages sent to the new phone number. If an &#x60;sms_application_sid&#x60; is present, we ignore all &#x60;sms_*_url&#x60; values and use those of the application.
 **SmsFallbackMethod** | **string** | The HTTP method that we should use to call &#x60;sms_fallback_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60; and defaults to &#x60;POST&#x60;.
 **SmsFallbackUrl** | **string** | The URL that we should call when an error occurs while requesting or executing the TwiML defined by &#x60;sms_url&#x60;.
 **SmsMethod** | **string** | The HTTP method that we should use to call &#x60;sms_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60; and defaults to &#x60;POST&#x60;.
 **SmsUrl** | **string** | The URL we should call when the new phone number receives an incoming SMS message.
 **StatusCallback** | **string** | The URL we should call using the &#x60;status_callback_method&#x60; to send status information to your application.
 **StatusCallbackMethod** | **string** | The HTTP method we should use to call &#x60;status_callback&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60; and defaults to &#x60;POST&#x60;.
 **TrunkSid** | **string** | The SID of the Trunk we should use to handle calls to the new phone number. If a &#x60;trunk_sid&#x60; is present, we ignore all of the voice urls and voice applications and use only those set on the Trunk. Setting a &#x60;trunk_sid&#x60; will automatically delete your &#x60;voice_application_sid&#x60; and vice versa.
 **VoiceApplicationSid** | **string** | The SID of the application we should use to handle calls to the new phone number. If a &#x60;voice_application_sid&#x60; is present, we ignore all of the voice urls and use those set on the application. Setting a &#x60;voice_application_sid&#x60; will automatically delete your &#x60;trunk_sid&#x60; and vice versa.
 **VoiceCallerIdLookup** | **bool** | Whether to lookup the caller&#39;s name from the CNAM database and post it to your app. Can be: &#x60;true&#x60; or &#x60;false&#x60; and defaults to &#x60;false&#x60;.
 **VoiceFallbackMethod** | **string** | The HTTP method that we should use to call &#x60;voice_fallback_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60; and defaults to &#x60;POST&#x60;.
 **VoiceFallbackUrl** | **string** | The URL that we should call when an error occurs retrieving or executing the TwiML requested by &#x60;url&#x60;.
 **VoiceMethod** | **string** | The HTTP method that we should use to call &#x60;voice_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60; and defaults to &#x60;POST&#x60;.
 **VoiceReceiveMode** | **string** | The configuration parameter for the new phone number to receive incoming voice calls or faxes. Can be: &#x60;fax&#x60; or &#x60;voice&#x60; and defaults to &#x60;voice&#x60;.
 **VoiceUrl** | **string** | The URL that we should call to answer a call to the new phone number. The &#x60;voice_url&#x60; will not be called if a &#x60;voice_application_sid&#x60; or a &#x60;trunk_sid&#x60; is set.

### Return type

[**ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree**](ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateMessage

> ApiV2010AccountMessage CreateMessage(ctx, AccountSid).AddressRetention(AddressRetention).ApplicationSid(ApplicationSid).Attempt(Attempt).Body(Body).ContentRetention(ContentRetention).ForceDelivery(ForceDelivery).From(From).MaxPrice(MaxPrice).MediaUrl(MediaUrl).MessagingServiceSid(MessagingServiceSid).PersistentAction(PersistentAction).ProvideFeedback(ProvideFeedback).SmartEncoded(SmartEncoded).StatusCallback(StatusCallback).To(To).ValidityPeriod(ValidityPeriod).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AccountSid := "AccountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that will create the resource.
    AddressRetention := "AddressRetention_example" // string | Determines if the address can be stored or obfuscated based on privacy settings (optional)
    ApplicationSid := "ApplicationSid_example" // string | The SID of the application that should receive message status. We POST a `message_sid` parameter and a `message_status` parameter with a value of `sent` or `failed` to the [application](https://www.twilio.com/docs/usage/api/applications)'s `message_status_callback`. If a `status_callback` parameter is also passed, it will be ignored and the application's `message_status_callback` parameter will be used. (optional)
    Attempt := int32(56) // int32 | Total number of attempts made ( including this ) to send out the message regardless of the provider used (optional)
    Body := "Body_example" // string | The text of the message you want to send. Can be up to 1,600 characters in length. (optional)
    ContentRetention := "ContentRetention_example" // string | Determines if the message content can be stored or redacted based on privacy settings (optional)
    ForceDelivery := true // bool | Reserved (optional)
    From := "From_example" // string | A Twilio phone number in [E.164](https://www.twilio.com/docs/glossary/what-e164) format, an [alphanumeric sender ID](https://www.twilio.com/docs/sms/send-messages#use-an-alphanumeric-sender-id), or a [Channel Endpoint address](https://www.twilio.com/docs/sms/channels#channel-addresses) that is enabled for the type of message you want to send. Phone numbers or [short codes](https://www.twilio.com/docs/sms/api/short-code) purchased from Twilio also work here. You cannot, for example, spoof messages from a private cell phone number. If you are using `messaging_service_sid`, this parameter must be empty. (optional)
    MaxPrice := float32(8.14) // float32 | The maximum total price in US dollars that you will pay for the message to be delivered. Can be a decimal value that has up to 4 decimal places. All messages are queued for delivery and the message cost is checked before the message is sent. If the cost exceeds `max_price`, the message will fail and a status of `Failed` is sent to the status callback. If `MaxPrice` is not set, the message cost is not checked. (optional)
    MediaUrl := []string{"Inner_example"} // []string | The URL of the media to send with the message. The media can be of type `gif`, `png`, and `jpeg` and will be formatted correctly on the recipient's device. The media size limit is 5MB for supported file types (JPEG, PNG, GIF) and 500KB for [other types](https://www.twilio.com/docs/sms/accepted-mime-types) of accepted media. To send more than one image in the message body, provide multiple `media_url` parameters in the POST request. You can include up to 10 `media_url` parameters per message. You can send images in an SMS message in only the US and Canada. (optional)
    MessagingServiceSid := "MessagingServiceSid_example" // string | The SID of the [Messaging Service](https://www.twilio.com/docs/sms/services#send-a-message-with-copilot) you want to associate with the Message. Set this parameter to use the [Messaging Service Settings and Copilot Features](https://www.twilio.com/console/sms/services) you have configured and leave the `from` parameter empty. When only this parameter is set, Twilio will use your enabled Copilot Features to select the `from` phone number for delivery. (optional)
    PersistentAction := []string{"Inner_example"} // []string | Rich actions for Channels Messages. (optional)
    ProvideFeedback := true // bool | Whether to confirm delivery of the message. Set this value to `true` if you are sending messages that have a trackable user action and you intend to confirm delivery of the message using the [Message Feedback API](https://www.twilio.com/docs/sms/api/message-feedback-resource). This parameter is `false` by default. (optional)
    SmartEncoded := true // bool | Whether to detect Unicode characters that have a similar GSM-7 character and replace them. Can be: `true` or `false`. (optional)
    StatusCallback := "StatusCallback_example" // string | The URL we should call using the `status_callback_method` to send status information to your application. If specified, we POST these message status changes to the URL: `queued`, `failed`, `sent`, `delivered`, or `undelivered`. Twilio will POST its [standard request parameters](https://www.twilio.com/docs/sms/twiml#request-parameters) as well as some additional parameters including `MessageSid`, `MessageStatus`, and `ErrorCode`. If you include this parameter with the `messaging_service_sid`, we use this URL instead of the Status Callback URL of the [Messaging Service](https://www.twilio.com/docs/sms/services/api). URLs must contain a valid hostname and underscores are not allowed. (optional)
    To := "To_example" // string | The destination phone number in [E.164](https://www.twilio.com/docs/glossary/what-e164) format for SMS/MMS or [Channel user address](https://www.twilio.com/docs/sms/channels#channel-addresses) for other 3rd-party channels. (optional)
    ValidityPeriod := int32(56) // int32 | How long in seconds the message can remain in our outgoing message queue. After this period elapses, the message fails and we call your status callback. Can be between 1 and the default value of 14,400 seconds. After a message has been accepted by a carrier, however, we cannot guarantee that the message will not be queued after this period. We recommend that this value be at least 5 seconds. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateMessage(context.Background(), AccountSid).AddressRetention(AddressRetention).ApplicationSid(ApplicationSid).Attempt(Attempt).Body(Body).ContentRetention(ContentRetention).ForceDelivery(ForceDelivery).From(From).MaxPrice(MaxPrice).MediaUrl(MediaUrl).MessagingServiceSid(MessagingServiceSid).PersistentAction(PersistentAction).ProvideFeedback(ProvideFeedback).SmartEncoded(SmartEncoded).StatusCallback(StatusCallback).To(To).ValidityPeriod(ValidityPeriod).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateMessage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateMessage`: ApiV2010AccountMessage
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateMessage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that will create the resource.

### Other Parameters

Other parameters are passed through a pointer to a CreateMessageParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **AddressRetention** | **string** | Determines if the address can be stored or obfuscated based on privacy settings
 **ApplicationSid** | **string** | The SID of the application that should receive message status. We POST a &#x60;message_sid&#x60; parameter and a &#x60;message_status&#x60; parameter with a value of &#x60;sent&#x60; or &#x60;failed&#x60; to the [application](https://www.twilio.com/docs/usage/api/applications)&#39;s &#x60;message_status_callback&#x60;. If a &#x60;status_callback&#x60; parameter is also passed, it will be ignored and the application&#39;s &#x60;message_status_callback&#x60; parameter will be used.
 **Attempt** | **int32** | Total number of attempts made ( including this ) to send out the message regardless of the provider used
 **Body** | **string** | The text of the message you want to send. Can be up to 1,600 characters in length.
 **ContentRetention** | **string** | Determines if the message content can be stored or redacted based on privacy settings
 **ForceDelivery** | **bool** | Reserved
 **From** | **string** | A Twilio phone number in [E.164](https://www.twilio.com/docs/glossary/what-e164) format, an [alphanumeric sender ID](https://www.twilio.com/docs/sms/send-messages#use-an-alphanumeric-sender-id), or a [Channel Endpoint address](https://www.twilio.com/docs/sms/channels#channel-addresses) that is enabled for the type of message you want to send. Phone numbers or [short codes](https://www.twilio.com/docs/sms/api/short-code) purchased from Twilio also work here. You cannot, for example, spoof messages from a private cell phone number. If you are using &#x60;messaging_service_sid&#x60;, this parameter must be empty.
 **MaxPrice** | **float32** | The maximum total price in US dollars that you will pay for the message to be delivered. Can be a decimal value that has up to 4 decimal places. All messages are queued for delivery and the message cost is checked before the message is sent. If the cost exceeds &#x60;max_price&#x60;, the message will fail and a status of &#x60;Failed&#x60; is sent to the status callback. If &#x60;MaxPrice&#x60; is not set, the message cost is not checked.
 **MediaUrl** | **[]string** | The URL of the media to send with the message. The media can be of type &#x60;gif&#x60;, &#x60;png&#x60;, and &#x60;jpeg&#x60; and will be formatted correctly on the recipient&#39;s device. The media size limit is 5MB for supported file types (JPEG, PNG, GIF) and 500KB for [other types](https://www.twilio.com/docs/sms/accepted-mime-types) of accepted media. To send more than one image in the message body, provide multiple &#x60;media_url&#x60; parameters in the POST request. You can include up to 10 &#x60;media_url&#x60; parameters per message. You can send images in an SMS message in only the US and Canada.
 **MessagingServiceSid** | **string** | The SID of the [Messaging Service](https://www.twilio.com/docs/sms/services#send-a-message-with-copilot) you want to associate with the Message. Set this parameter to use the [Messaging Service Settings and Copilot Features](https://www.twilio.com/console/sms/services) you have configured and leave the &#x60;from&#x60; parameter empty. When only this parameter is set, Twilio will use your enabled Copilot Features to select the &#x60;from&#x60; phone number for delivery.
 **PersistentAction** | **[]string** | Rich actions for Channels Messages.
 **ProvideFeedback** | **bool** | Whether to confirm delivery of the message. Set this value to &#x60;true&#x60; if you are sending messages that have a trackable user action and you intend to confirm delivery of the message using the [Message Feedback API](https://www.twilio.com/docs/sms/api/message-feedback-resource). This parameter is &#x60;false&#x60; by default.
 **SmartEncoded** | **bool** | Whether to detect Unicode characters that have a similar GSM-7 character and replace them. Can be: &#x60;true&#x60; or &#x60;false&#x60;.
 **StatusCallback** | **string** | The URL we should call using the &#x60;status_callback_method&#x60; to send status information to your application. If specified, we POST these message status changes to the URL: &#x60;queued&#x60;, &#x60;failed&#x60;, &#x60;sent&#x60;, &#x60;delivered&#x60;, or &#x60;undelivered&#x60;. Twilio will POST its [standard request parameters](https://www.twilio.com/docs/sms/twiml#request-parameters) as well as some additional parameters including &#x60;MessageSid&#x60;, &#x60;MessageStatus&#x60;, and &#x60;ErrorCode&#x60;. If you include this parameter with the &#x60;messaging_service_sid&#x60;, we use this URL instead of the Status Callback URL of the [Messaging Service](https://www.twilio.com/docs/sms/services/api). URLs must contain a valid hostname and underscores are not allowed.
 **To** | **string** | The destination phone number in [E.164](https://www.twilio.com/docs/glossary/what-e164) format for SMS/MMS or [Channel user address](https://www.twilio.com/docs/sms/channels#channel-addresses) for other 3rd-party channels.
 **ValidityPeriod** | **int32** | How long in seconds the message can remain in our outgoing message queue. After this period elapses, the message fails and we call your status callback. Can be between 1 and the default value of 14,400 seconds. After a message has been accepted by a carrier, however, we cannot guarantee that the message will not be queued after this period. We recommend that this value be at least 5 seconds.

### Return type

[**ApiV2010AccountMessage**](ApiV2010AccountMessage.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateMessageFeedback

> ApiV2010AccountMessageMessageFeedback CreateMessageFeedback(ctx, AccountSid, MessageSid).Outcome(Outcome).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AccountSid := "AccountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that will create the resource.
    MessageSid := "MessageSid_example" // string | The SID of the Message resource for which the feedback was provided.
    Outcome := "Outcome_example" // string | Whether the feedback has arrived. Can be: `unconfirmed` or `confirmed`. If `provide_feedback`=`true` in [the initial HTTP POST](https://www.twilio.com/docs/sms/api/message-resource#create-a-message-resource), the initial value of this property is `unconfirmed`. After the message arrives, update the value to `confirmed`. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateMessageFeedback(context.Background(), AccountSid, MessageSid).Outcome(Outcome).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateMessageFeedback``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateMessageFeedback`: ApiV2010AccountMessageMessageFeedback
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateMessageFeedback`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that will create the resource.
**MessageSid** | **string** | The SID of the Message resource for which the feedback was provided.

### Other Parameters

Other parameters are passed through a pointer to a CreateMessageFeedbackParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


 **Outcome** | **string** | Whether the feedback has arrived. Can be: &#x60;unconfirmed&#x60; or &#x60;confirmed&#x60;. If &#x60;provide_feedback&#x60;&#x3D;&#x60;true&#x60; in [the initial HTTP POST](https://www.twilio.com/docs/sms/api/message-resource#create-a-message-resource), the initial value of this property is &#x60;unconfirmed&#x60;. After the message arrives, update the value to &#x60;confirmed&#x60;.

### Return type

[**ApiV2010AccountMessageMessageFeedback**](ApiV2010AccountMessageMessageFeedback.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateNewKey

> ApiV2010AccountNewKey CreateNewKey(ctx, AccountSid).FriendlyName(FriendlyName).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AccountSid := "AccountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that will be responsible for the new Key resource.
    FriendlyName := "FriendlyName_example" // string | A descriptive string that you create to describe the resource. It can be up to 64 characters long. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateNewKey(context.Background(), AccountSid).FriendlyName(FriendlyName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateNewKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNewKey`: ApiV2010AccountNewKey
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateNewKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that will be responsible for the new Key resource.

### Other Parameters

Other parameters are passed through a pointer to a CreateNewKeyParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **FriendlyName** | **string** | A descriptive string that you create to describe the resource. It can be up to 64 characters long.

### Return type

[**ApiV2010AccountNewKey**](ApiV2010AccountNewKey.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateNewSigningKey

> ApiV2010AccountNewSigningKey CreateNewSigningKey(ctx, AccountSid).FriendlyName(FriendlyName).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AccountSid := "AccountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that will be responsible for the new Key resource.
    FriendlyName := "FriendlyName_example" // string | A descriptive string that you create to describe the resource. It can be up to 64 characters long. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateNewSigningKey(context.Background(), AccountSid).FriendlyName(FriendlyName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateNewSigningKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNewSigningKey`: ApiV2010AccountNewSigningKey
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateNewSigningKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that will be responsible for the new Key resource.

### Other Parameters

Other parameters are passed through a pointer to a CreateNewSigningKeyParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **FriendlyName** | **string** | A descriptive string that you create to describe the resource. It can be up to 64 characters long.

### Return type

[**ApiV2010AccountNewSigningKey**](ApiV2010AccountNewSigningKey.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateParticipant

> ApiV2010AccountConferenceParticipant CreateParticipant(ctx, AccountSid, ConferenceSid).Beep(Beep).Byoc(Byoc).CallReason(CallReason).CallSidToCoach(CallSidToCoach).CallerId(CallerId).Coaching(Coaching).ConferenceRecord(ConferenceRecord).ConferenceRecordingStatusCallback(ConferenceRecordingStatusCallback).ConferenceRecordingStatusCallbackEvent(ConferenceRecordingStatusCallbackEvent).ConferenceRecordingStatusCallbackMethod(ConferenceRecordingStatusCallbackMethod).ConferenceStatusCallback(ConferenceStatusCallback).ConferenceStatusCallbackEvent(ConferenceStatusCallbackEvent).ConferenceStatusCallbackMethod(ConferenceStatusCallbackMethod).ConferenceTrim(ConferenceTrim).EarlyMedia(EarlyMedia).EndConferenceOnExit(EndConferenceOnExit).From(From).JitterBufferSize(JitterBufferSize).Label(Label).MaxParticipants(MaxParticipants).Muted(Muted).Record(Record).RecordingChannels(RecordingChannels).RecordingStatusCallback(RecordingStatusCallback).RecordingStatusCallbackEvent(RecordingStatusCallbackEvent).RecordingStatusCallbackMethod(RecordingStatusCallbackMethod).RecordingTrack(RecordingTrack).Region(Region).SipAuthPassword(SipAuthPassword).SipAuthUsername(SipAuthUsername).StartConferenceOnEnter(StartConferenceOnEnter).StatusCallback(StatusCallback).StatusCallbackEvent(StatusCallbackEvent).StatusCallbackMethod(StatusCallbackMethod).Timeout(Timeout).To(To).WaitMethod(WaitMethod).WaitUrl(WaitUrl).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AccountSid := "AccountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that will create the resource.
    ConferenceSid := "ConferenceSid_example" // string | The SID of the participant's conference.
    Beep := "Beep_example" // string | Whether to play a notification beep to the conference when the participant joins. Can be: `true`, `false`, `onEnter`, or `onExit`. The default value is `true`. (optional)
    Byoc := "Byoc_example" // string | The SID of a BYOC (Bring Your Own Carrier) trunk to route this call with. Note that `byoc` is only meaningful when `to` is a phone number; it will otherwise be ignored. (Beta) (optional)
    CallReason := "CallReason_example" // string | The Reason for the outgoing call. Use it to specify the purpose of the call that is presented on the called party's phone. (Branded Calls Beta) (optional)
    CallSidToCoach := "CallSidToCoach_example" // string | The SID of the participant who is being `coached`. The participant being coached is the only participant who can hear the participant who is `coaching`. (optional)
    CallerId := "CallerId_example" // string | The phone number, Client identifier, or username portion of SIP address that made this call. Phone numbers are in [E.164](https://www.twilio.com/docs/glossary/what-e164) format (e.g., +16175551212). Client identifiers are formatted `client:name`. If using a phone number, it must be a Twilio number or a Verified [outgoing caller id](https://www.twilio.com/docs/voice/api/outgoing-caller-ids) for your account. If the `to` parameter is a phone number, `callerId` must also be a phone number. If `to` is sip address, this value of `callerId` should be a username portion to be used to populate the From header that is passed to the SIP endpoint. (optional)
    Coaching := true // bool | Whether the participant is coaching another call. Can be: `true` or `false`. If not present, defaults to `false` unless `call_sid_to_coach` is defined. If `true`, `call_sid_to_coach` must be defined. (optional)
    ConferenceRecord := "ConferenceRecord_example" // string | Whether to record the conference the participant is joining. Can be: `true`, `false`, `record-from-start`, and `do-not-record`. The default value is `false`. (optional)
    ConferenceRecordingStatusCallback := "ConferenceRecordingStatusCallback_example" // string | The URL we should call using the `conference_recording_status_callback_method` when the conference recording is available. (optional)
    ConferenceRecordingStatusCallbackEvent := []string{"Inner_example"} // []string | The conference recording state changes that generate a call to `conference_recording_status_callback`. Can be: `in-progress`, `completed`, and `failed`. Separate multiple values with a space. The default value is `in-progress completed failed`. (optional)
    ConferenceRecordingStatusCallbackMethod := "ConferenceRecordingStatusCallbackMethod_example" // string | The HTTP method we should use to call `conference_recording_status_callback`. Can be: `GET` or `POST` and defaults to `POST`. (optional)
    ConferenceStatusCallback := "ConferenceStatusCallback_example" // string | The URL we should call using the `conference_status_callback_method` when the conference events in `conference_status_callback_event` occur. Only the value set by the first participant to join the conference is used. Subsequent `conference_status_callback` values are ignored. (optional)
    ConferenceStatusCallbackEvent := []string{"Inner_example"} // []string | The conference state changes that should generate a call to `conference_status_callback`. Can be: `start`, `end`, `join`, `leave`, `mute`, `hold`, `speaker`, and `announcement`. Separate multiple values with a space. Defaults to `start end`. (optional)
    ConferenceStatusCallbackMethod := "ConferenceStatusCallbackMethod_example" // string | The HTTP method we should use to call `conference_status_callback`. Can be: `GET` or `POST` and defaults to `POST`. (optional)
    ConferenceTrim := "ConferenceTrim_example" // string | Whether to trim leading and trailing silence from your recorded conference audio files. Can be: `trim-silence` or `do-not-trim` and defaults to `trim-silence`. (optional)
    EarlyMedia := true // bool | Whether to allow an agent to hear the state of the outbound call, including ringing or disconnect messages. Can be: `true` or `false` and defaults to `true`. (optional)
    EndConferenceOnExit := true // bool | Whether to end the conference when the participant leaves. Can be: `true` or `false` and defaults to `false`. (optional)
    From := "From_example" // string | The phone number, Client identifier, or username portion of SIP address that made this call. Phone numbers are in [E.164](https://www.twilio.com/docs/glossary/what-e164) format (e.g., +16175551212). Client identifiers are formatted `client:name`. If using a phone number, it must be a Twilio number or a Verified [outgoing caller id](https://www.twilio.com/docs/voice/api/outgoing-caller-ids) for your account. If the `to` parameter is a phone number, `from` must also be a phone number. If `to` is sip address, this value of `from` should be a username portion to be used to populate the P-Asserted-Identity header that is passed to the SIP endpoint. (optional)
    JitterBufferSize := "JitterBufferSize_example" // string | Jitter buffer size for the connecting participant. Twilio will use this setting to apply Jitter Buffer before participant's audio is mixed into the conference. Can be: `off`, `small`, `medium`, and `large`. Default to `large`. (optional)
    Label := "Label_example" // string | A label for this participant. If one is supplied, it may subsequently be used to fetch, update or delete the participant. (optional)
    MaxParticipants := int32(56) // int32 | The maximum number of participants in the conference. Can be a positive integer from `2` to `250`. The default value is `250`. (optional)
    Muted := true // bool | Whether the agent is muted in the conference. Can be `true` or `false` and the default is `false`. (optional)
    Record := true // bool | Whether to record the participant and their conferences, including the time between conferences. Can be `true` or `false` and the default is `false`. (optional)
    RecordingChannels := "RecordingChannels_example" // string | The recording channels for the final recording. Can be: `mono` or `dual` and the default is `mono`. (optional)
    RecordingStatusCallback := "RecordingStatusCallback_example" // string | The URL that we should call using the `recording_status_callback_method` when the recording status changes. (optional)
    RecordingStatusCallbackEvent := []string{"Inner_example"} // []string | The recording state changes that should generate a call to `recording_status_callback`. Can be: `in-progress`, `completed`, and `failed`. Separate multiple values with a space. The default value is `in-progress completed failed`. (optional)
    RecordingStatusCallbackMethod := "RecordingStatusCallbackMethod_example" // string | The HTTP method we should use when we call `recording_status_callback`. Can be: `GET` or `POST` and defaults to `POST`. (optional)
    RecordingTrack := "RecordingTrack_example" // string | The audio track to record for the call. Can be: `inbound`, `outbound` or `both`. The default is `both`. `inbound` records the audio that is received by Twilio. `outbound` records the audio that is sent from Twilio. `both` records the audio that is received and sent by Twilio. (optional)
    Region := "Region_example" // string | The [region](https://support.twilio.com/hc/en-us/articles/223132167-How-global-low-latency-routing-and-region-selection-work-for-conferences-and-Client-calls) where we should mix the recorded audio. Can be:`us1`, `ie1`, `de1`, `sg1`, `br1`, `au1`, or `jp1`. (optional)
    SipAuthPassword := "SipAuthPassword_example" // string | The SIP password for authentication. (optional)
    SipAuthUsername := "SipAuthUsername_example" // string | The SIP username used for authentication. (optional)
    StartConferenceOnEnter := true // bool | Whether to start the conference when the participant joins, if it has not already started. Can be: `true` or `false` and the default is `true`. If `false` and the conference has not started, the participant is muted and hears background music until another participant starts the conference. (optional)
    StatusCallback := "StatusCallback_example" // string | The URL we should call using the `status_callback_method` to send status information to your application. (optional)
    StatusCallbackEvent := []string{"Inner_example"} // []string | The conference state changes that should generate a call to `status_callback`. Can be: `initiated`, `ringing`, `answered`, and `completed`. Separate multiple values with a space. The default value is `completed`. (optional)
    StatusCallbackMethod := "StatusCallbackMethod_example" // string | The HTTP method we should use to call `status_callback`. Can be: `GET` and `POST` and defaults to `POST`. (optional)
    Timeout := int32(56) // int32 | The number of seconds that we should allow the phone to ring before assuming there is no answer. Can be an integer between `5` and `600`, inclusive. The default value is `60`. We always add a 5-second timeout buffer to outgoing calls, so  value of 10 would result in an actual timeout that was closer to 15 seconds. (optional)
    To := "To_example" // string | The phone number, SIP address, or Client identifier that received this call. Phone numbers are in [E.164](https://www.twilio.com/docs/glossary/what-e164) format (e.g., +16175551212). SIP addresses are formatted as `sip:name@company.com`. Client identifiers are formatted `client:name`. [Custom parameters](https://www.twilio.com/docs/voice/api/conference-participant-resource#custom-parameters) may also be specified. (optional)
    WaitMethod := "WaitMethod_example" // string | The HTTP method we should use to call `wait_url`. Can be `GET` or `POST` and the default is `POST`. When using a static audio file, this should be `GET` so that we can cache the file. (optional)
    WaitUrl := "WaitUrl_example" // string | The URL we should call using the `wait_method` for the music to play while participants are waiting for the conference to start. The default value is the URL of our standard hold music. [Learn more about hold music](https://www.twilio.com/labs/twimlets/holdmusic). (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateParticipant(context.Background(), AccountSid, ConferenceSid).Beep(Beep).Byoc(Byoc).CallReason(CallReason).CallSidToCoach(CallSidToCoach).CallerId(CallerId).Coaching(Coaching).ConferenceRecord(ConferenceRecord).ConferenceRecordingStatusCallback(ConferenceRecordingStatusCallback).ConferenceRecordingStatusCallbackEvent(ConferenceRecordingStatusCallbackEvent).ConferenceRecordingStatusCallbackMethod(ConferenceRecordingStatusCallbackMethod).ConferenceStatusCallback(ConferenceStatusCallback).ConferenceStatusCallbackEvent(ConferenceStatusCallbackEvent).ConferenceStatusCallbackMethod(ConferenceStatusCallbackMethod).ConferenceTrim(ConferenceTrim).EarlyMedia(EarlyMedia).EndConferenceOnExit(EndConferenceOnExit).From(From).JitterBufferSize(JitterBufferSize).Label(Label).MaxParticipants(MaxParticipants).Muted(Muted).Record(Record).RecordingChannels(RecordingChannels).RecordingStatusCallback(RecordingStatusCallback).RecordingStatusCallbackEvent(RecordingStatusCallbackEvent).RecordingStatusCallbackMethod(RecordingStatusCallbackMethod).RecordingTrack(RecordingTrack).Region(Region).SipAuthPassword(SipAuthPassword).SipAuthUsername(SipAuthUsername).StartConferenceOnEnter(StartConferenceOnEnter).StatusCallback(StatusCallback).StatusCallbackEvent(StatusCallbackEvent).StatusCallbackMethod(StatusCallbackMethod).Timeout(Timeout).To(To).WaitMethod(WaitMethod).WaitUrl(WaitUrl).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateParticipant``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateParticipant`: ApiV2010AccountConferenceParticipant
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateParticipant`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that will create the resource.
**ConferenceSid** | **string** | The SID of the participant&#39;s conference.

### Other Parameters

Other parameters are passed through a pointer to a CreateParticipantParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


 **Beep** | **string** | Whether to play a notification beep to the conference when the participant joins. Can be: &#x60;true&#x60;, &#x60;false&#x60;, &#x60;onEnter&#x60;, or &#x60;onExit&#x60;. The default value is &#x60;true&#x60;.
 **Byoc** | **string** | The SID of a BYOC (Bring Your Own Carrier) trunk to route this call with. Note that &#x60;byoc&#x60; is only meaningful when &#x60;to&#x60; is a phone number; it will otherwise be ignored. (Beta)
 **CallReason** | **string** | The Reason for the outgoing call. Use it to specify the purpose of the call that is presented on the called party&#39;s phone. (Branded Calls Beta)
 **CallSidToCoach** | **string** | The SID of the participant who is being &#x60;coached&#x60;. The participant being coached is the only participant who can hear the participant who is &#x60;coaching&#x60;.
 **CallerId** | **string** | The phone number, Client identifier, or username portion of SIP address that made this call. Phone numbers are in [E.164](https://www.twilio.com/docs/glossary/what-e164) format (e.g., +16175551212). Client identifiers are formatted &#x60;client:name&#x60;. If using a phone number, it must be a Twilio number or a Verified [outgoing caller id](https://www.twilio.com/docs/voice/api/outgoing-caller-ids) for your account. If the &#x60;to&#x60; parameter is a phone number, &#x60;callerId&#x60; must also be a phone number. If &#x60;to&#x60; is sip address, this value of &#x60;callerId&#x60; should be a username portion to be used to populate the From header that is passed to the SIP endpoint.
 **Coaching** | **bool** | Whether the participant is coaching another call. Can be: &#x60;true&#x60; or &#x60;false&#x60;. If not present, defaults to &#x60;false&#x60; unless &#x60;call_sid_to_coach&#x60; is defined. If &#x60;true&#x60;, &#x60;call_sid_to_coach&#x60; must be defined.
 **ConferenceRecord** | **string** | Whether to record the conference the participant is joining. Can be: &#x60;true&#x60;, &#x60;false&#x60;, &#x60;record-from-start&#x60;, and &#x60;do-not-record&#x60;. The default value is &#x60;false&#x60;.
 **ConferenceRecordingStatusCallback** | **string** | The URL we should call using the &#x60;conference_recording_status_callback_method&#x60; when the conference recording is available.
 **ConferenceRecordingStatusCallbackEvent** | **[]string** | The conference recording state changes that generate a call to &#x60;conference_recording_status_callback&#x60;. Can be: &#x60;in-progress&#x60;, &#x60;completed&#x60;, and &#x60;failed&#x60;. Separate multiple values with a space. The default value is &#x60;in-progress completed failed&#x60;.
 **ConferenceRecordingStatusCallbackMethod** | **string** | The HTTP method we should use to call &#x60;conference_recording_status_callback&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60; and defaults to &#x60;POST&#x60;.
 **ConferenceStatusCallback** | **string** | The URL we should call using the &#x60;conference_status_callback_method&#x60; when the conference events in &#x60;conference_status_callback_event&#x60; occur. Only the value set by the first participant to join the conference is used. Subsequent &#x60;conference_status_callback&#x60; values are ignored.
 **ConferenceStatusCallbackEvent** | **[]string** | The conference state changes that should generate a call to &#x60;conference_status_callback&#x60;. Can be: &#x60;start&#x60;, &#x60;end&#x60;, &#x60;join&#x60;, &#x60;leave&#x60;, &#x60;mute&#x60;, &#x60;hold&#x60;, &#x60;speaker&#x60;, and &#x60;announcement&#x60;. Separate multiple values with a space. Defaults to &#x60;start end&#x60;.
 **ConferenceStatusCallbackMethod** | **string** | The HTTP method we should use to call &#x60;conference_status_callback&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60; and defaults to &#x60;POST&#x60;.
 **ConferenceTrim** | **string** | Whether to trim leading and trailing silence from your recorded conference audio files. Can be: &#x60;trim-silence&#x60; or &#x60;do-not-trim&#x60; and defaults to &#x60;trim-silence&#x60;.
 **EarlyMedia** | **bool** | Whether to allow an agent to hear the state of the outbound call, including ringing or disconnect messages. Can be: &#x60;true&#x60; or &#x60;false&#x60; and defaults to &#x60;true&#x60;.
 **EndConferenceOnExit** | **bool** | Whether to end the conference when the participant leaves. Can be: &#x60;true&#x60; or &#x60;false&#x60; and defaults to &#x60;false&#x60;.
 **From** | **string** | The phone number, Client identifier, or username portion of SIP address that made this call. Phone numbers are in [E.164](https://www.twilio.com/docs/glossary/what-e164) format (e.g., +16175551212). Client identifiers are formatted &#x60;client:name&#x60;. If using a phone number, it must be a Twilio number or a Verified [outgoing caller id](https://www.twilio.com/docs/voice/api/outgoing-caller-ids) for your account. If the &#x60;to&#x60; parameter is a phone number, &#x60;from&#x60; must also be a phone number. If &#x60;to&#x60; is sip address, this value of &#x60;from&#x60; should be a username portion to be used to populate the P-Asserted-Identity header that is passed to the SIP endpoint.
 **JitterBufferSize** | **string** | Jitter buffer size for the connecting participant. Twilio will use this setting to apply Jitter Buffer before participant&#39;s audio is mixed into the conference. Can be: &#x60;off&#x60;, &#x60;small&#x60;, &#x60;medium&#x60;, and &#x60;large&#x60;. Default to &#x60;large&#x60;.
 **Label** | **string** | A label for this participant. If one is supplied, it may subsequently be used to fetch, update or delete the participant.
 **MaxParticipants** | **int32** | The maximum number of participants in the conference. Can be a positive integer from &#x60;2&#x60; to &#x60;250&#x60;. The default value is &#x60;250&#x60;.
 **Muted** | **bool** | Whether the agent is muted in the conference. Can be &#x60;true&#x60; or &#x60;false&#x60; and the default is &#x60;false&#x60;.
 **Record** | **bool** | Whether to record the participant and their conferences, including the time between conferences. Can be &#x60;true&#x60; or &#x60;false&#x60; and the default is &#x60;false&#x60;.
 **RecordingChannels** | **string** | The recording channels for the final recording. Can be: &#x60;mono&#x60; or &#x60;dual&#x60; and the default is &#x60;mono&#x60;.
 **RecordingStatusCallback** | **string** | The URL that we should call using the &#x60;recording_status_callback_method&#x60; when the recording status changes.
 **RecordingStatusCallbackEvent** | **[]string** | The recording state changes that should generate a call to &#x60;recording_status_callback&#x60;. Can be: &#x60;in-progress&#x60;, &#x60;completed&#x60;, and &#x60;failed&#x60;. Separate multiple values with a space. The default value is &#x60;in-progress completed failed&#x60;.
 **RecordingStatusCallbackMethod** | **string** | The HTTP method we should use when we call &#x60;recording_status_callback&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60; and defaults to &#x60;POST&#x60;.
 **RecordingTrack** | **string** | The audio track to record for the call. Can be: &#x60;inbound&#x60;, &#x60;outbound&#x60; or &#x60;both&#x60;. The default is &#x60;both&#x60;. &#x60;inbound&#x60; records the audio that is received by Twilio. &#x60;outbound&#x60; records the audio that is sent from Twilio. &#x60;both&#x60; records the audio that is received and sent by Twilio.
 **Region** | **string** | The [region](https://support.twilio.com/hc/en-us/articles/223132167-How-global-low-latency-routing-and-region-selection-work-for-conferences-and-Client-calls) where we should mix the recorded audio. Can be:&#x60;us1&#x60;, &#x60;ie1&#x60;, &#x60;de1&#x60;, &#x60;sg1&#x60;, &#x60;br1&#x60;, &#x60;au1&#x60;, or &#x60;jp1&#x60;.
 **SipAuthPassword** | **string** | The SIP password for authentication.
 **SipAuthUsername** | **string** | The SIP username used for authentication.
 **StartConferenceOnEnter** | **bool** | Whether to start the conference when the participant joins, if it has not already started. Can be: &#x60;true&#x60; or &#x60;false&#x60; and the default is &#x60;true&#x60;. If &#x60;false&#x60; and the conference has not started, the participant is muted and hears background music until another participant starts the conference.
 **StatusCallback** | **string** | The URL we should call using the &#x60;status_callback_method&#x60; to send status information to your application.
 **StatusCallbackEvent** | **[]string** | The conference state changes that should generate a call to &#x60;status_callback&#x60;. Can be: &#x60;initiated&#x60;, &#x60;ringing&#x60;, &#x60;answered&#x60;, and &#x60;completed&#x60;. Separate multiple values with a space. The default value is &#x60;completed&#x60;.
 **StatusCallbackMethod** | **string** | The HTTP method we should use to call &#x60;status_callback&#x60;. Can be: &#x60;GET&#x60; and &#x60;POST&#x60; and defaults to &#x60;POST&#x60;.
 **Timeout** | **int32** | The number of seconds that we should allow the phone to ring before assuming there is no answer. Can be an integer between &#x60;5&#x60; and &#x60;600&#x60;, inclusive. The default value is &#x60;60&#x60;. We always add a 5-second timeout buffer to outgoing calls, so  value of 10 would result in an actual timeout that was closer to 15 seconds.
 **To** | **string** | The phone number, SIP address, or Client identifier that received this call. Phone numbers are in [E.164](https://www.twilio.com/docs/glossary/what-e164) format (e.g., +16175551212). SIP addresses are formatted as &#x60;sip:name@company.com&#x60;. Client identifiers are formatted &#x60;client:name&#x60;. [Custom parameters](https://www.twilio.com/docs/voice/api/conference-participant-resource#custom-parameters) may also be specified.
 **WaitMethod** | **string** | The HTTP method we should use to call &#x60;wait_url&#x60;. Can be &#x60;GET&#x60; or &#x60;POST&#x60; and the default is &#x60;POST&#x60;. When using a static audio file, this should be &#x60;GET&#x60; so that we can cache the file.
 **WaitUrl** | **string** | The URL we should call using the &#x60;wait_method&#x60; for the music to play while participants are waiting for the conference to start. The default value is the URL of our standard hold music. [Learn more about hold music](https://www.twilio.com/labs/twimlets/holdmusic).

### Return type

[**ApiV2010AccountConferenceParticipant**](ApiV2010AccountConferenceParticipant.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePayments

> ApiV2010AccountCallPayments CreatePayments(ctx, AccountSid, CallSid).BankAccountType(BankAccountType).ChargeAmount(ChargeAmount).Currency(Currency).Description(Description).IdempotencyKey(IdempotencyKey).Input(Input).MinPostalCodeLength(MinPostalCodeLength).Parameter(Parameter).PaymentConnector(PaymentConnector).PaymentMethod(PaymentMethod).PostalCode(PostalCode).SecurityCode(SecurityCode).StatusCallback(StatusCallback).Timeout(Timeout).TokenType(TokenType).ValidCardTypes(ValidCardTypes).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AccountSid := "AccountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that will create the resource.
    CallSid := "CallSid_example" // string | The SID of the call that will create the resource. Call leg associated with this sid is expected to provide payment information thru DTMF.
    BankAccountType := "BankAccountType_example" // string | Type of bank account if payment source is ACH. One of `consumer-checking`, `consumer-savings`, or `commercial-checking`. The default value is `consumer-checking`. (optional)
    ChargeAmount := float32(8.14) // float32 | A positive decimal value less than 1,000,000 to charge against the credit card or bank account. Default currency can be overwritten with `currency` field. Leave blank or set to 0 to tokenize. (optional)
    Currency := "Currency_example" // string | The currency of the `charge_amount`, formatted as [ISO 4127](http://www.iso.org/iso/home/standards/currency_codes.htm) format. The default value is `USD` and all values allowed from the <Pay> Connector are accepted. (optional)
    Description := "Description_example" // string | The description can be used to provide more details regarding the transaction. This information is submitted along with the payment details to the Payment Connector which are then posted on the transactions. (optional)
    IdempotencyKey := "IdempotencyKey_example" // string | A unique token that will be used to ensure that multiple API calls with the same information do not result in multiple transactions. This should be a unique string value per API call and can be a randomly generated. (optional)
    Input := "Input_example" // string | A list of inputs that should be accepted. Currently only `dtmf` is supported. All digits captured during a pay session are redacted from the logs. (optional)
    MinPostalCodeLength := int32(56) // int32 | A positive integer that is used to validate the length of the `PostalCode` inputted by the user. User must enter this many digits. (optional)
    Parameter := TODO // map[string]interface{} | A single level JSON string that is required when accepting certain information specific only to ACH payments. The information that has to be included here depends on the <Pay> Connector. [Read more](https://www.twilio.com/console/voice/pay-connectors). (optional)
    PaymentConnector := "PaymentConnector_example" // string | This is the unique name corresponding to the Payment Gateway Connector installed in the Twilio Add-ons. Learn more about [<Pay> Connectors](https://www.twilio.com/console/voice/pay-connectors). The default value is `Default`. (optional)
    PaymentMethod := "PaymentMethod_example" // string | Type of payment being captured. One of `credit-card` or `ach-debit`. The default value is `credit-card`. (optional)
    PostalCode := true // bool | Indicates whether the credit card postal code (zip code) is a required piece of payment information that must be provided by the caller. The default is `true`. (optional)
    SecurityCode := true // bool | Indicates whether the credit card security code is a required piece of payment information that must be provided by the caller. The default is `true`. (optional)
    StatusCallback := "StatusCallback_example" // string | Provide an absolute or relative URL to receive status updates regarding your Pay session. Read more about the [expected StatusCallback values](https://www.twilio.com/docs/voice/api/payment-resource#statuscallback) (optional)
    Timeout := int32(56) // int32 | The number of seconds that <Pay> should wait for the caller to press a digit between each subsequent digit, after the first one, before moving on to validate the digits captured. The default is `5`, maximum is `600`. (optional)
    TokenType := "TokenType_example" // string | Indicates whether the payment method should be tokenized as a `one-time` or `reusable` token. The default value is `reusable`. Do not enter a charge amount when tokenizing. If a charge amount is entered, the payment method will be charged and not tokenized. (optional)
    ValidCardTypes := "ValidCardTypes_example" // string | Credit card types separated by space that Pay should accept. The default value is `visa mastercard amex` (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreatePayments(context.Background(), AccountSid, CallSid).BankAccountType(BankAccountType).ChargeAmount(ChargeAmount).Currency(Currency).Description(Description).IdempotencyKey(IdempotencyKey).Input(Input).MinPostalCodeLength(MinPostalCodeLength).Parameter(Parameter).PaymentConnector(PaymentConnector).PaymentMethod(PaymentMethod).PostalCode(PostalCode).SecurityCode(SecurityCode).StatusCallback(StatusCallback).Timeout(Timeout).TokenType(TokenType).ValidCardTypes(ValidCardTypes).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreatePayments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePayments`: ApiV2010AccountCallPayments
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreatePayments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that will create the resource.
**CallSid** | **string** | The SID of the call that will create the resource. Call leg associated with this sid is expected to provide payment information thru DTMF.

### Other Parameters

Other parameters are passed through a pointer to a CreatePaymentsParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


 **BankAccountType** | **string** | Type of bank account if payment source is ACH. One of &#x60;consumer-checking&#x60;, &#x60;consumer-savings&#x60;, or &#x60;commercial-checking&#x60;. The default value is &#x60;consumer-checking&#x60;.
 **ChargeAmount** | **float32** | A positive decimal value less than 1,000,000 to charge against the credit card or bank account. Default currency can be overwritten with &#x60;currency&#x60; field. Leave blank or set to 0 to tokenize.
 **Currency** | **string** | The currency of the &#x60;charge_amount&#x60;, formatted as [ISO 4127](http://www.iso.org/iso/home/standards/currency_codes.htm) format. The default value is &#x60;USD&#x60; and all values allowed from the &lt;Pay&gt; Connector are accepted.
 **Description** | **string** | The description can be used to provide more details regarding the transaction. This information is submitted along with the payment details to the Payment Connector which are then posted on the transactions.
 **IdempotencyKey** | **string** | A unique token that will be used to ensure that multiple API calls with the same information do not result in multiple transactions. This should be a unique string value per API call and can be a randomly generated.
 **Input** | **string** | A list of inputs that should be accepted. Currently only &#x60;dtmf&#x60; is supported. All digits captured during a pay session are redacted from the logs.
 **MinPostalCodeLength** | **int32** | A positive integer that is used to validate the length of the &#x60;PostalCode&#x60; inputted by the user. User must enter this many digits.
 **Parameter** | [**map[string]interface{}**](map[string]interface{}.md) | A single level JSON string that is required when accepting certain information specific only to ACH payments. The information that has to be included here depends on the &lt;Pay&gt; Connector. [Read more](https://www.twilio.com/console/voice/pay-connectors).
 **PaymentConnector** | **string** | This is the unique name corresponding to the Payment Gateway Connector installed in the Twilio Add-ons. Learn more about [&lt;Pay&gt; Connectors](https://www.twilio.com/console/voice/pay-connectors). The default value is &#x60;Default&#x60;.
 **PaymentMethod** | **string** | Type of payment being captured. One of &#x60;credit-card&#x60; or &#x60;ach-debit&#x60;. The default value is &#x60;credit-card&#x60;.
 **PostalCode** | **bool** | Indicates whether the credit card postal code (zip code) is a required piece of payment information that must be provided by the caller. The default is &#x60;true&#x60;.
 **SecurityCode** | **bool** | Indicates whether the credit card security code is a required piece of payment information that must be provided by the caller. The default is &#x60;true&#x60;.
 **StatusCallback** | **string** | Provide an absolute or relative URL to receive status updates regarding your Pay session. Read more about the [expected StatusCallback values](https://www.twilio.com/docs/voice/api/payment-resource#statuscallback)
 **Timeout** | **int32** | The number of seconds that &lt;Pay&gt; should wait for the caller to press a digit between each subsequent digit, after the first one, before moving on to validate the digits captured. The default is &#x60;5&#x60;, maximum is &#x60;600&#x60;.
 **TokenType** | **string** | Indicates whether the payment method should be tokenized as a &#x60;one-time&#x60; or &#x60;reusable&#x60; token. The default value is &#x60;reusable&#x60;. Do not enter a charge amount when tokenizing. If a charge amount is entered, the payment method will be charged and not tokenized.
 **ValidCardTypes** | **string** | Credit card types separated by space that Pay should accept. The default value is &#x60;visa mastercard amex&#x60;

### Return type

[**ApiV2010AccountCallPayments**](ApiV2010AccountCallPayments.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateQueue

> ApiV2010AccountQueue CreateQueue(ctx, AccountSid).FriendlyName(FriendlyName).MaxSize(MaxSize).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AccountSid := "AccountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that will create the resource.
    FriendlyName := "FriendlyName_example" // string | A descriptive string that you created to describe this resource. It can be up to 64 characters long. (optional)
    MaxSize := int32(56) // int32 | The maximum number of calls allowed to be in the queue. The default is 100. The maximum is 5000. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateQueue(context.Background(), AccountSid).FriendlyName(FriendlyName).MaxSize(MaxSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateQueue``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateQueue`: ApiV2010AccountQueue
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateQueue`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that will create the resource.

### Other Parameters

Other parameters are passed through a pointer to a CreateQueueParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **FriendlyName** | **string** | A descriptive string that you created to describe this resource. It can be up to 64 characters long.
 **MaxSize** | **int32** | The maximum number of calls allowed to be in the queue. The default is 100. The maximum is 5000.

### Return type

[**ApiV2010AccountQueue**](ApiV2010AccountQueue.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSipAuthCallsCredentialListMapping

> ApiV2010AccountSipSipDomainSipAuthSipAuthCallsSipAuthCallsCredentialListMapping CreateSipAuthCallsCredentialListMapping(ctx, AccountSid, DomainSid).CredentialListSid(CredentialListSid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AccountSid := "AccountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that will create the resource.
    DomainSid := "DomainSid_example" // string | The SID of the SIP domain that will contain the new resource.
    CredentialListSid := "CredentialListSid_example" // string | The SID of the CredentialList resource to map to the SIP domain. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateSipAuthCallsCredentialListMapping(context.Background(), AccountSid, DomainSid).CredentialListSid(CredentialListSid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateSipAuthCallsCredentialListMapping``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSipAuthCallsCredentialListMapping`: ApiV2010AccountSipSipDomainSipAuthSipAuthCallsSipAuthCallsCredentialListMapping
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateSipAuthCallsCredentialListMapping`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that will create the resource.
**DomainSid** | **string** | The SID of the SIP domain that will contain the new resource.

### Other Parameters

Other parameters are passed through a pointer to a CreateSipAuthCallsCredentialListMappingParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


 **CredentialListSid** | **string** | The SID of the CredentialList resource to map to the SIP domain.

### Return type

[**ApiV2010AccountSipSipDomainSipAuthSipAuthCallsSipAuthCallsCredentialListMapping**](ApiV2010AccountSipSipDomainSipAuthSipAuthCallsSipAuthCallsCredentialListMapping.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSipAuthCallsIpAccessControlListMapping

> ApiV2010AccountSipSipDomainSipAuthSipAuthCallsSipAuthCallsIpAccessControlListMapping CreateSipAuthCallsIpAccessControlListMapping(ctx, AccountSid, DomainSid).IpAccessControlListSid(IpAccessControlListSid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AccountSid := "AccountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that will create the resource.
    DomainSid := "DomainSid_example" // string | The SID of the SIP domain that will contain the new resource.
    IpAccessControlListSid := "IpAccessControlListSid_example" // string | The SID of the IpAccessControlList resource to map to the SIP domain. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateSipAuthCallsIpAccessControlListMapping(context.Background(), AccountSid, DomainSid).IpAccessControlListSid(IpAccessControlListSid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateSipAuthCallsIpAccessControlListMapping``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSipAuthCallsIpAccessControlListMapping`: ApiV2010AccountSipSipDomainSipAuthSipAuthCallsSipAuthCallsIpAccessControlListMapping
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateSipAuthCallsIpAccessControlListMapping`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that will create the resource.
**DomainSid** | **string** | The SID of the SIP domain that will contain the new resource.

### Other Parameters

Other parameters are passed through a pointer to a CreateSipAuthCallsIpAccessControlListMappingParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


 **IpAccessControlListSid** | **string** | The SID of the IpAccessControlList resource to map to the SIP domain.

### Return type

[**ApiV2010AccountSipSipDomainSipAuthSipAuthCallsSipAuthCallsIpAccessControlListMapping**](ApiV2010AccountSipSipDomainSipAuthSipAuthCallsSipAuthCallsIpAccessControlListMapping.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSipAuthRegistrationsCredentialListMapping

> ApiV2010AccountSipSipDomainSipAuthSipAuthRegistrationsSipAuthRegistrationsCredentialListMapping CreateSipAuthRegistrationsCredentialListMapping(ctx, AccountSid, DomainSid).CredentialListSid(CredentialListSid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AccountSid := "AccountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that will create the resource.
    DomainSid := "DomainSid_example" // string | The SID of the SIP domain that will contain the new resource.
    CredentialListSid := "CredentialListSid_example" // string | The SID of the CredentialList resource to map to the SIP domain. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateSipAuthRegistrationsCredentialListMapping(context.Background(), AccountSid, DomainSid).CredentialListSid(CredentialListSid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateSipAuthRegistrationsCredentialListMapping``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSipAuthRegistrationsCredentialListMapping`: ApiV2010AccountSipSipDomainSipAuthSipAuthRegistrationsSipAuthRegistrationsCredentialListMapping
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateSipAuthRegistrationsCredentialListMapping`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that will create the resource.
**DomainSid** | **string** | The SID of the SIP domain that will contain the new resource.

### Other Parameters

Other parameters are passed through a pointer to a CreateSipAuthRegistrationsCredentialListMappingParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


 **CredentialListSid** | **string** | The SID of the CredentialList resource to map to the SIP domain.

### Return type

[**ApiV2010AccountSipSipDomainSipAuthSipAuthRegistrationsSipAuthRegistrationsCredentialListMapping**](ApiV2010AccountSipSipDomainSipAuthSipAuthRegistrationsSipAuthRegistrationsCredentialListMapping.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSipCredential

> ApiV2010AccountSipSipCredentialListSipCredential CreateSipCredential(ctx, AccountSid, CredentialListSid).Password(Password).Username(Username).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AccountSid := "AccountSid_example" // string | The unique id of the Account that is responsible for this resource.
    CredentialListSid := "CredentialListSid_example" // string | The unique id that identifies the credential list to include the created credential.
    Password := "Password_example" // string | The password that the username will use when authenticating SIP requests. The password must be a minimum of 12 characters, contain at least 1 digit, and have mixed case. (eg `IWasAtSignal2018`) (optional)
    Username := "Username_example" // string | The username that will be passed when authenticating SIP requests. The username should be sent in response to Twilio's challenge of the initial INVITE. It can be up to 32 characters long. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateSipCredential(context.Background(), AccountSid, CredentialListSid).Password(Password).Username(Username).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateSipCredential``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSipCredential`: ApiV2010AccountSipSipCredentialListSipCredential
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateSipCredential`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string** | The unique id of the Account that is responsible for this resource.
**CredentialListSid** | **string** | The unique id that identifies the credential list to include the created credential.

### Other Parameters

Other parameters are passed through a pointer to a CreateSipCredentialParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


 **Password** | **string** | The password that the username will use when authenticating SIP requests. The password must be a minimum of 12 characters, contain at least 1 digit, and have mixed case. (eg &#x60;IWasAtSignal2018&#x60;)
 **Username** | **string** | The username that will be passed when authenticating SIP requests. The username should be sent in response to Twilio&#39;s challenge of the initial INVITE. It can be up to 32 characters long.

### Return type

[**ApiV2010AccountSipSipCredentialListSipCredential**](ApiV2010AccountSipSipCredentialListSipCredential.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSipCredentialList

> ApiV2010AccountSipSipCredentialList CreateSipCredentialList(ctx, AccountSid).FriendlyName(FriendlyName).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AccountSid := "AccountSid_example" // string | The unique id of the Account that is responsible for this resource.
    FriendlyName := "FriendlyName_example" // string | A human readable descriptive text that describes the CredentialList, up to 64 characters long. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateSipCredentialList(context.Background(), AccountSid).FriendlyName(FriendlyName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateSipCredentialList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSipCredentialList`: ApiV2010AccountSipSipCredentialList
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateSipCredentialList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string** | The unique id of the Account that is responsible for this resource.

### Other Parameters

Other parameters are passed through a pointer to a CreateSipCredentialListParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **FriendlyName** | **string** | A human readable descriptive text that describes the CredentialList, up to 64 characters long.

### Return type

[**ApiV2010AccountSipSipCredentialList**](ApiV2010AccountSipSipCredentialList.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSipCredentialListMapping

> ApiV2010AccountSipSipDomainSipCredentialListMapping CreateSipCredentialListMapping(ctx, AccountSid, DomainSid).CredentialListSid(CredentialListSid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AccountSid := "AccountSid_example" // string | The unique id of the [Account](https://www.twilio.com/docs/iam/api/account) responsible for this resource.
    DomainSid := "DomainSid_example" // string | A 34 character string that uniquely identifies the SIP Domain for which the CredentialList resource will be mapped.
    CredentialListSid := "CredentialListSid_example" // string | A 34 character string that uniquely identifies the CredentialList resource to map to the SIP domain. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateSipCredentialListMapping(context.Background(), AccountSid, DomainSid).CredentialListSid(CredentialListSid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateSipCredentialListMapping``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSipCredentialListMapping`: ApiV2010AccountSipSipDomainSipCredentialListMapping
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateSipCredentialListMapping`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string** | The unique id of the [Account](https://www.twilio.com/docs/iam/api/account) responsible for this resource.
**DomainSid** | **string** | A 34 character string that uniquely identifies the SIP Domain for which the CredentialList resource will be mapped.

### Other Parameters

Other parameters are passed through a pointer to a CreateSipCredentialListMappingParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


 **CredentialListSid** | **string** | A 34 character string that uniquely identifies the CredentialList resource to map to the SIP domain.

### Return type

[**ApiV2010AccountSipSipDomainSipCredentialListMapping**](ApiV2010AccountSipSipDomainSipCredentialListMapping.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSipDomain

> ApiV2010AccountSipSipDomain CreateSipDomain(ctx, AccountSid).ByocTrunkSid(ByocTrunkSid).DomainName(DomainName).EmergencyCallerSid(EmergencyCallerSid).EmergencyCallingEnabled(EmergencyCallingEnabled).FriendlyName(FriendlyName).Secure(Secure).SipRegistration(SipRegistration).VoiceFallbackMethod(VoiceFallbackMethod).VoiceFallbackUrl(VoiceFallbackUrl).VoiceMethod(VoiceMethod).VoiceStatusCallbackMethod(VoiceStatusCallbackMethod).VoiceStatusCallbackUrl(VoiceStatusCallbackUrl).VoiceUrl(VoiceUrl).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AccountSid := "AccountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that will create the resource.
    ByocTrunkSid := "ByocTrunkSid_example" // string | The SID of the BYOC Trunk(Bring Your Own Carrier) resource that the Sip Domain will be associated with. (optional)
    DomainName := "DomainName_example" // string | The unique address you reserve on Twilio to which you route your SIP traffic. Domain names can contain letters, digits, and \\\"-\\\" and must end with `sip.twilio.com`. (optional)
    EmergencyCallerSid := "EmergencyCallerSid_example" // string | Whether an emergency caller sid is configured for the domain. If present, this phone number will be used as the callback for the emergency call. (optional)
    EmergencyCallingEnabled := true // bool | Whether emergency calling is enabled for the domain. If enabled, allows emergency calls on the domain from phone numbers with validated addresses. (optional)
    FriendlyName := "FriendlyName_example" // string | A descriptive string that you created to describe the resource. It can be up to 64 characters long. (optional)
    Secure := true // bool | Whether secure SIP is enabled for the domain. If enabled, TLS will be enforced and SRTP will be negotiated on all incoming calls to this sip domain. (optional)
    SipRegistration := true // bool | Whether to allow SIP Endpoints to register with the domain to receive calls. Can be `true` or `false`. `true` allows SIP Endpoints to register with the domain to receive calls, `false` does not. (optional)
    VoiceFallbackMethod := "VoiceFallbackMethod_example" // string | The HTTP method we should use to call `voice_fallback_url`. Can be: `GET` or `POST`. (optional)
    VoiceFallbackUrl := "VoiceFallbackUrl_example" // string | The URL that we should call when an error occurs while retrieving or executing the TwiML from `voice_url`. (optional)
    VoiceMethod := "VoiceMethod_example" // string | The HTTP method we should use to call `voice_url`. Can be: `GET` or `POST`. (optional)
    VoiceStatusCallbackMethod := "VoiceStatusCallbackMethod_example" // string | The HTTP method we should use to call `voice_status_callback_url`. Can be: `GET` or `POST`. (optional)
    VoiceStatusCallbackUrl := "VoiceStatusCallbackUrl_example" // string | The URL that we should call to pass status parameters (such as call ended) to your application. (optional)
    VoiceUrl := "VoiceUrl_example" // string | The URL we should when the domain receives a call. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateSipDomain(context.Background(), AccountSid).ByocTrunkSid(ByocTrunkSid).DomainName(DomainName).EmergencyCallerSid(EmergencyCallerSid).EmergencyCallingEnabled(EmergencyCallingEnabled).FriendlyName(FriendlyName).Secure(Secure).SipRegistration(SipRegistration).VoiceFallbackMethod(VoiceFallbackMethod).VoiceFallbackUrl(VoiceFallbackUrl).VoiceMethod(VoiceMethod).VoiceStatusCallbackMethod(VoiceStatusCallbackMethod).VoiceStatusCallbackUrl(VoiceStatusCallbackUrl).VoiceUrl(VoiceUrl).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateSipDomain``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSipDomain`: ApiV2010AccountSipSipDomain
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateSipDomain`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that will create the resource.

### Other Parameters

Other parameters are passed through a pointer to a CreateSipDomainParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **ByocTrunkSid** | **string** | The SID of the BYOC Trunk(Bring Your Own Carrier) resource that the Sip Domain will be associated with.
 **DomainName** | **string** | The unique address you reserve on Twilio to which you route your SIP traffic. Domain names can contain letters, digits, and \\\&quot;-\\\&quot; and must end with &#x60;sip.twilio.com&#x60;.
 **EmergencyCallerSid** | **string** | Whether an emergency caller sid is configured for the domain. If present, this phone number will be used as the callback for the emergency call.
 **EmergencyCallingEnabled** | **bool** | Whether emergency calling is enabled for the domain. If enabled, allows emergency calls on the domain from phone numbers with validated addresses.
 **FriendlyName** | **string** | A descriptive string that you created to describe the resource. It can be up to 64 characters long.
 **Secure** | **bool** | Whether secure SIP is enabled for the domain. If enabled, TLS will be enforced and SRTP will be negotiated on all incoming calls to this sip domain.
 **SipRegistration** | **bool** | Whether to allow SIP Endpoints to register with the domain to receive calls. Can be &#x60;true&#x60; or &#x60;false&#x60;. &#x60;true&#x60; allows SIP Endpoints to register with the domain to receive calls, &#x60;false&#x60; does not.
 **VoiceFallbackMethod** | **string** | The HTTP method we should use to call &#x60;voice_fallback_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60;.
 **VoiceFallbackUrl** | **string** | The URL that we should call when an error occurs while retrieving or executing the TwiML from &#x60;voice_url&#x60;.
 **VoiceMethod** | **string** | The HTTP method we should use to call &#x60;voice_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60;.
 **VoiceStatusCallbackMethod** | **string** | The HTTP method we should use to call &#x60;voice_status_callback_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60;.
 **VoiceStatusCallbackUrl** | **string** | The URL that we should call to pass status parameters (such as call ended) to your application.
 **VoiceUrl** | **string** | The URL we should when the domain receives a call.

### Return type

[**ApiV2010AccountSipSipDomain**](ApiV2010AccountSipSipDomain.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSipIpAccessControlList

> ApiV2010AccountSipSipIpAccessControlList CreateSipIpAccessControlList(ctx, AccountSid).FriendlyName(FriendlyName).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AccountSid := "AccountSid_example" // string | The unique id of the [Account](https://www.twilio.com/docs/iam/api/account) responsible for this resource.
    FriendlyName := "FriendlyName_example" // string | A human readable descriptive text that describes the IpAccessControlList, up to 64 characters long. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateSipIpAccessControlList(context.Background(), AccountSid).FriendlyName(FriendlyName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateSipIpAccessControlList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSipIpAccessControlList`: ApiV2010AccountSipSipIpAccessControlList
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateSipIpAccessControlList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string** | The unique id of the [Account](https://www.twilio.com/docs/iam/api/account) responsible for this resource.

### Other Parameters

Other parameters are passed through a pointer to a CreateSipIpAccessControlListParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **FriendlyName** | **string** | A human readable descriptive text that describes the IpAccessControlList, up to 64 characters long.

### Return type

[**ApiV2010AccountSipSipIpAccessControlList**](ApiV2010AccountSipSipIpAccessControlList.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSipIpAccessControlListMapping

> ApiV2010AccountSipSipDomainSipIpAccessControlListMapping CreateSipIpAccessControlListMapping(ctx, AccountSid, DomainSid).IpAccessControlListSid(IpAccessControlListSid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AccountSid := "AccountSid_example" // string | The unique id of the Account that is responsible for this resource.
    DomainSid := "DomainSid_example" // string | A 34 character string that uniquely identifies the SIP domain.
    IpAccessControlListSid := "IpAccessControlListSid_example" // string | The unique id of the IP access control list to map to the SIP domain. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateSipIpAccessControlListMapping(context.Background(), AccountSid, DomainSid).IpAccessControlListSid(IpAccessControlListSid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateSipIpAccessControlListMapping``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSipIpAccessControlListMapping`: ApiV2010AccountSipSipDomainSipIpAccessControlListMapping
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateSipIpAccessControlListMapping`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string** | The unique id of the Account that is responsible for this resource.
**DomainSid** | **string** | A 34 character string that uniquely identifies the SIP domain.

### Other Parameters

Other parameters are passed through a pointer to a CreateSipIpAccessControlListMappingParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


 **IpAccessControlListSid** | **string** | The unique id of the IP access control list to map to the SIP domain.

### Return type

[**ApiV2010AccountSipSipDomainSipIpAccessControlListMapping**](ApiV2010AccountSipSipDomainSipIpAccessControlListMapping.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSipIpAddress

> ApiV2010AccountSipSipIpAccessControlListSipIpAddress CreateSipIpAddress(ctx, AccountSid, IpAccessControlListSid).CidrPrefixLength(CidrPrefixLength).FriendlyName(FriendlyName).IpAddress(IpAddress).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AccountSid := "AccountSid_example" // string | The unique id of the [Account](https://www.twilio.com/docs/iam/api/account) responsible for this resource.
    IpAccessControlListSid := "IpAccessControlListSid_example" // string | The IpAccessControlList Sid with which to associate the created IpAddress resource.
    CidrPrefixLength := int32(56) // int32 | An integer representing the length of the CIDR prefix to use with this IP address when accepting traffic. By default the entire IP address is used. (optional)
    FriendlyName := "FriendlyName_example" // string | A human readable descriptive text for this resource, up to 64 characters long. (optional)
    IpAddress := "IpAddress_example" // string | An IP address in dotted decimal notation from which you want to accept traffic. Any SIP requests from this IP address will be allowed by Twilio. IPv4 only supported today. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateSipIpAddress(context.Background(), AccountSid, IpAccessControlListSid).CidrPrefixLength(CidrPrefixLength).FriendlyName(FriendlyName).IpAddress(IpAddress).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateSipIpAddress``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSipIpAddress`: ApiV2010AccountSipSipIpAccessControlListSipIpAddress
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateSipIpAddress`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string** | The unique id of the [Account](https://www.twilio.com/docs/iam/api/account) responsible for this resource.
**IpAccessControlListSid** | **string** | The IpAccessControlList Sid with which to associate the created IpAddress resource.

### Other Parameters

Other parameters are passed through a pointer to a CreateSipIpAddressParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


 **CidrPrefixLength** | **int32** | An integer representing the length of the CIDR prefix to use with this IP address when accepting traffic. By default the entire IP address is used.
 **FriendlyName** | **string** | A human readable descriptive text for this resource, up to 64 characters long.
 **IpAddress** | **string** | An IP address in dotted decimal notation from which you want to accept traffic. Any SIP requests from this IP address will be allowed by Twilio. IPv4 only supported today.

### Return type

[**ApiV2010AccountSipSipIpAccessControlListSipIpAddress**](ApiV2010AccountSipSipIpAccessControlListSipIpAddress.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateToken

> ApiV2010AccountToken CreateToken(ctx, AccountSid).Ttl(Ttl).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AccountSid := "AccountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that will create the resource.
    Ttl := int32(56) // int32 | The duration in seconds for which the generated credentials are valid. The default value is 86400 (24 hours). (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateToken(context.Background(), AccountSid).Ttl(Ttl).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateToken`: ApiV2010AccountToken
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that will create the resource.

### Other Parameters

Other parameters are passed through a pointer to a CreateTokenParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **Ttl** | **int32** | The duration in seconds for which the generated credentials are valid. The default value is 86400 (24 hours).

### Return type

[**ApiV2010AccountToken**](ApiV2010AccountToken.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateUsageTrigger

> ApiV2010AccountUsageUsageTrigger CreateUsageTrigger(ctx, AccountSid).CallbackMethod(CallbackMethod).CallbackUrl(CallbackUrl).FriendlyName(FriendlyName).Recurring(Recurring).TriggerBy(TriggerBy).TriggerValue(TriggerValue).UsageCategory(UsageCategory).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AccountSid := "AccountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that will create the resource.
    CallbackMethod := "CallbackMethod_example" // string | The HTTP method we should use to call `callback_url`. Can be: `GET` or `POST` and the default is `POST`. (optional)
    CallbackUrl := "CallbackUrl_example" // string | The URL we should call using `callback_method` when the trigger fires. (optional)
    FriendlyName := "FriendlyName_example" // string | A descriptive string that you create to describe the resource. It can be up to 64 characters long. (optional)
    Recurring := "Recurring_example" // string | The frequency of a recurring UsageTrigger.  Can be: `daily`, `monthly`, or `yearly` for recurring triggers or empty for non-recurring triggers. A trigger will only fire once during each period. Recurring times are in GMT. (optional)
    TriggerBy := "TriggerBy_example" // string | The field in the [UsageRecord](https://www.twilio.com/docs/usage/api/usage-record) resource that should fire the trigger.  Can be: `count`, `usage`, or `price` as described in the [UsageRecords documentation](https://www.twilio.com/docs/usage/api/usage-record#usage-count-price).  The default is `usage`. (optional)
    TriggerValue := "TriggerValue_example" // string | The usage value at which the trigger should fire.  For convenience, you can use an offset value such as `+30` to specify a trigger_value that is 30 units more than the current usage value. Be sure to urlencode a `+` as `%2B`. (optional)
    UsageCategory := "UsageCategory_example" // string | The usage category that the trigger should watch.  Use one of the supported [usage categories](https://www.twilio.com/docs/usage/api/usage-record#usage-categories) for this value. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateUsageTrigger(context.Background(), AccountSid).CallbackMethod(CallbackMethod).CallbackUrl(CallbackUrl).FriendlyName(FriendlyName).Recurring(Recurring).TriggerBy(TriggerBy).TriggerValue(TriggerValue).UsageCategory(UsageCategory).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateUsageTrigger``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateUsageTrigger`: ApiV2010AccountUsageUsageTrigger
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateUsageTrigger`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that will create the resource.

### Other Parameters

Other parameters are passed through a pointer to a CreateUsageTriggerParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **CallbackMethod** | **string** | The HTTP method we should use to call &#x60;callback_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60; and the default is &#x60;POST&#x60;.
 **CallbackUrl** | **string** | The URL we should call using &#x60;callback_method&#x60; when the trigger fires.
 **FriendlyName** | **string** | A descriptive string that you create to describe the resource. It can be up to 64 characters long.
 **Recurring** | **string** | The frequency of a recurring UsageTrigger.  Can be: &#x60;daily&#x60;, &#x60;monthly&#x60;, or &#x60;yearly&#x60; for recurring triggers or empty for non-recurring triggers. A trigger will only fire once during each period. Recurring times are in GMT.
 **TriggerBy** | **string** | The field in the [UsageRecord](https://www.twilio.com/docs/usage/api/usage-record) resource that should fire the trigger.  Can be: &#x60;count&#x60;, &#x60;usage&#x60;, or &#x60;price&#x60; as described in the [UsageRecords documentation](https://www.twilio.com/docs/usage/api/usage-record#usage-count-price).  The default is &#x60;usage&#x60;.
 **TriggerValue** | **string** | The usage value at which the trigger should fire.  For convenience, you can use an offset value such as &#x60;+30&#x60; to specify a trigger_value that is 30 units more than the current usage value. Be sure to urlencode a &#x60;+&#x60; as &#x60;%2B&#x60;.
 **UsageCategory** | **string** | The usage category that the trigger should watch.  Use one of the supported [usage categories](https://www.twilio.com/docs/usage/api/usage-record#usage-categories) for this value.

### Return type

[**ApiV2010AccountUsageUsageTrigger**](ApiV2010AccountUsageUsageTrigger.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateValidationRequest

> ApiV2010AccountValidationRequest CreateValidationRequest(ctx, AccountSid).CallDelay(CallDelay).Extension(Extension).FriendlyName(FriendlyName).PhoneNumber(PhoneNumber).StatusCallback(StatusCallback).StatusCallbackMethod(StatusCallbackMethod).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AccountSid := "AccountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) responsible for the new caller ID resource.
    CallDelay := int32(56) // int32 | The number of seconds to delay before initiating the verification call. Can be an integer between `0` and `60`, inclusive. The default is `0`. (optional)
    Extension := "Extension_example" // string | The digits to dial after connecting the verification call. (optional)
    FriendlyName := "FriendlyName_example" // string | A descriptive string that you create to describe the new caller ID resource. It can be up to 64 characters long. The default value is a formatted version of the phone number. (optional)
    PhoneNumber := "PhoneNumber_example" // string | The phone number to verify in [E.164](https://www.twilio.com/docs/glossary/what-e164) format, which consists of a + followed by the country code and subscriber number. (optional)
    StatusCallback := "StatusCallback_example" // string | The URL we should call using the `status_callback_method` to send status information about the verification process to your application. (optional)
    StatusCallbackMethod := "StatusCallbackMethod_example" // string | The HTTP method we should use to call `status_callback`. Can be: `GET` or `POST`, and the default is `POST`. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateValidationRequest(context.Background(), AccountSid).CallDelay(CallDelay).Extension(Extension).FriendlyName(FriendlyName).PhoneNumber(PhoneNumber).StatusCallback(StatusCallback).StatusCallbackMethod(StatusCallbackMethod).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateValidationRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateValidationRequest`: ApiV2010AccountValidationRequest
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateValidationRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) responsible for the new caller ID resource.

### Other Parameters

Other parameters are passed through a pointer to a CreateValidationRequestParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **CallDelay** | **int32** | The number of seconds to delay before initiating the verification call. Can be an integer between &#x60;0&#x60; and &#x60;60&#x60;, inclusive. The default is &#x60;0&#x60;.
 **Extension** | **string** | The digits to dial after connecting the verification call.
 **FriendlyName** | **string** | A descriptive string that you create to describe the new caller ID resource. It can be up to 64 characters long. The default value is a formatted version of the phone number.
 **PhoneNumber** | **string** | The phone number to verify in [E.164](https://www.twilio.com/docs/glossary/what-e164) format, which consists of a + followed by the country code and subscriber number.
 **StatusCallback** | **string** | The URL we should call using the &#x60;status_callback_method&#x60; to send status information about the verification process to your application.
 **StatusCallbackMethod** | **string** | The HTTP method we should use to call &#x60;status_callback&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60;, and the default is &#x60;POST&#x60;.

### Return type

[**ApiV2010AccountValidationRequest**](ApiV2010AccountValidationRequest.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAddress

> DeleteAddress(ctx, AccountSid, Sid).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AccountSid := "AccountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that is responsible for the Address resource to delete.
    Sid := "Sid_example" // string | The Twilio-provided string that uniquely identifies the Address resource to delete.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteAddress(context.Background(), AccountSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteAddress``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that is responsible for the Address resource to delete.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the Address resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeleteAddressParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteApplication

> DeleteApplication(ctx, AccountSid, Sid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AccountSid := "AccountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Application resources to delete.
    Sid := "Sid_example" // string | The Twilio-provided string that uniquely identifies the Application resource to delete.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteApplication(context.Background(), AccountSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteApplication``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Application resources to delete.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the Application resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeleteApplicationParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCall

> DeleteCall(ctx, AccountSid, Sid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AccountSid := "AccountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Call resource(s) to delete.
    Sid := "Sid_example" // string | The Twilio-provided Call SID that uniquely identifies the Call resource to delete

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteCall(context.Background(), AccountSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteCall``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Call resource(s) to delete.
**Sid** | **string** | The Twilio-provided Call SID that uniquely identifies the Call resource to delete

### Other Parameters

Other parameters are passed through a pointer to a DeleteCallParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCallFeedbackSummary

> DeleteCallFeedbackSummary(ctx, AccountSid, Sid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AccountSid := "AccountSid_example" // string | The unique id of the [Account](https://www.twilio.com/docs/iam/api/account) responsible for this resource.
    Sid := "Sid_example" // string | A 34 character string that uniquely identifies this resource.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteCallFeedbackSummary(context.Background(), AccountSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteCallFeedbackSummary``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string** | The unique id of the [Account](https://www.twilio.com/docs/iam/api/account) responsible for this resource.
**Sid** | **string** | A 34 character string that uniquely identifies this resource.

### Other Parameters

Other parameters are passed through a pointer to a DeleteCallFeedbackSummaryParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCallRecording

> DeleteCallRecording(ctx, AccountSid, CallSid, Sid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AccountSid := "AccountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Recording resources to delete.
    CallSid := "CallSid_example" // string | The [Call](https://www.twilio.com/docs/voice/api/call-resource) SID of the resources to delete.
    Sid := "Sid_example" // string | The Twilio-provided string that uniquely identifies the Recording resource to delete.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteCallRecording(context.Background(), AccountSid, CallSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteCallRecording``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Recording resources to delete.
**CallSid** | **string** | The [Call](https://www.twilio.com/docs/voice/api/call-resource) SID of the resources to delete.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the Recording resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeleteCallRecordingParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteConferenceRecording

> DeleteConferenceRecording(ctx, AccountSid, ConferenceSid, Sid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AccountSid := "AccountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Conference Recording resources to delete.
    ConferenceSid := "ConferenceSid_example" // string | The Conference SID that identifies the conference associated with the recording to delete.
    Sid := "Sid_example" // string | The Twilio-provided string that uniquely identifies the Conference Recording resource to delete.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteConferenceRecording(context.Background(), AccountSid, ConferenceSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteConferenceRecording``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Conference Recording resources to delete.
**ConferenceSid** | **string** | The Conference SID that identifies the conference associated with the recording to delete.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the Conference Recording resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeleteConferenceRecordingParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteConnectApp

> DeleteConnectApp(ctx, AccountSid, Sid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AccountSid := "AccountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the ConnectApp resource to fetch.
    Sid := "Sid_example" // string | The Twilio-provided string that uniquely identifies the ConnectApp resource to fetch.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteConnectApp(context.Background(), AccountSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteConnectApp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the ConnectApp resource to fetch.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the ConnectApp resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a DeleteConnectAppParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteIncomingPhoneNumber

> DeleteIncomingPhoneNumber(ctx, AccountSid, Sid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AccountSid := "AccountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the IncomingPhoneNumber resources to delete.
    Sid := "Sid_example" // string | The Twilio-provided string that uniquely identifies the IncomingPhoneNumber resource to delete.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteIncomingPhoneNumber(context.Background(), AccountSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteIncomingPhoneNumber``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the IncomingPhoneNumber resources to delete.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the IncomingPhoneNumber resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeleteIncomingPhoneNumberParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteIncomingPhoneNumberAssignedAddOn

> DeleteIncomingPhoneNumberAssignedAddOn(ctx, AccountSid, ResourceSid, Sid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AccountSid := "AccountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the resources to delete.
    ResourceSid := "ResourceSid_example" // string | The SID of the Phone Number to which the Add-on is assigned.
    Sid := "Sid_example" // string | The Twilio-provided string that uniquely identifies the resource to delete.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteIncomingPhoneNumberAssignedAddOn(context.Background(), AccountSid, ResourceSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteIncomingPhoneNumberAssignedAddOn``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the resources to delete.
**ResourceSid** | **string** | The SID of the Phone Number to which the Add-on is assigned.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeleteIncomingPhoneNumberAssignedAddOnParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteKey

> DeleteKey(ctx, AccountSid, Sid).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AccountSid := "AccountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Key resources to delete.
    Sid := "Sid_example" // string | The Twilio-provided string that uniquely identifies the Key resource to delete.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteKey(context.Background(), AccountSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Key resources to delete.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the Key resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeleteKeyParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteMedia

> DeleteMedia(ctx, AccountSid, MessageSid, Sid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AccountSid := "AccountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Media resource(s) to delete.
    MessageSid := "MessageSid_example" // string | The SID of the Message resource that this Media resource belongs to.
    Sid := "Sid_example" // string | The Twilio-provided string that uniquely identifies the Media resource to delete

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteMedia(context.Background(), AccountSid, MessageSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteMedia``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Media resource(s) to delete.
**MessageSid** | **string** | The SID of the Message resource that this Media resource belongs to.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the Media resource to delete

### Other Parameters

Other parameters are passed through a pointer to a DeleteMediaParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteMessage

> DeleteMessage(ctx, AccountSid, Sid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AccountSid := "AccountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Message resources to delete.
    Sid := "Sid_example" // string | The Twilio-provided string that uniquely identifies the Message resource to delete.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteMessage(context.Background(), AccountSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteMessage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Message resources to delete.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the Message resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeleteMessageParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOutgoingCallerId

> DeleteOutgoingCallerId(ctx, AccountSid, Sid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AccountSid := "AccountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the OutgoingCallerId resources to delete.
    Sid := "Sid_example" // string | The Twilio-provided string that uniquely identifies the OutgoingCallerId resource to delete.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteOutgoingCallerId(context.Background(), AccountSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteOutgoingCallerId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the OutgoingCallerId resources to delete.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the OutgoingCallerId resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeleteOutgoingCallerIdParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteParticipant

> DeleteParticipant(ctx, AccountSid, ConferenceSid, CallSid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AccountSid := "AccountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Participant resources to delete.
    ConferenceSid := "ConferenceSid_example" // string | The SID of the conference with the participants to delete.
    CallSid := "CallSid_example" // string | The [Call](https://www.twilio.com/docs/voice/api/call-resource) SID or label of the participant to delete. Non URL safe characters in a label must be percent encoded, for example, a space character is represented as %20.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteParticipant(context.Background(), AccountSid, ConferenceSid, CallSid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteParticipant``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Participant resources to delete.
**ConferenceSid** | **string** | The SID of the conference with the participants to delete.
**CallSid** | **string** | The [Call](https://www.twilio.com/docs/voice/api/call-resource) SID or label of the participant to delete. Non URL safe characters in a label must be percent encoded, for example, a space character is represented as %20.

### Other Parameters

Other parameters are passed through a pointer to a DeleteParticipantParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteQueue

> DeleteQueue(ctx, AccountSid, Sid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AccountSid := "AccountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Queue resource to delete.
    Sid := "Sid_example" // string | The Twilio-provided string that uniquely identifies the Queue resource to delete

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteQueue(context.Background(), AccountSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteQueue``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Queue resource to delete.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the Queue resource to delete

### Other Parameters

Other parameters are passed through a pointer to a DeleteQueueParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRecording

> DeleteRecording(ctx, AccountSid, Sid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AccountSid := "AccountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Recording resources to delete.
    Sid := "Sid_example" // string | The Twilio-provided string that uniquely identifies the Recording resource to delete.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteRecording(context.Background(), AccountSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteRecording``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Recording resources to delete.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the Recording resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeleteRecordingParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRecordingAddOnResult

> DeleteRecordingAddOnResult(ctx, AccountSid, ReferenceSid, Sid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AccountSid := "AccountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Recording AddOnResult resources to delete.
    ReferenceSid := "ReferenceSid_example" // string | The SID of the recording to which the result to delete belongs.
    Sid := "Sid_example" // string | The Twilio-provided string that uniquely identifies the Recording AddOnResult resource to delete.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteRecordingAddOnResult(context.Background(), AccountSid, ReferenceSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteRecordingAddOnResult``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Recording AddOnResult resources to delete.
**ReferenceSid** | **string** | The SID of the recording to which the result to delete belongs.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the Recording AddOnResult resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeleteRecordingAddOnResultParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRecordingAddOnResultPayload

> DeleteRecordingAddOnResultPayload(ctx, AccountSid, ReferenceSid, AddOnResultSid, Sid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AccountSid := "AccountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Recording AddOnResult Payload resources to delete.
    ReferenceSid := "ReferenceSid_example" // string | The SID of the recording to which the AddOnResult resource that contains the payloads to delete belongs.
    AddOnResultSid := "AddOnResultSid_example" // string | The SID of the AddOnResult to which the payloads to delete belongs.
    Sid := "Sid_example" // string | The Twilio-provided string that uniquely identifies the Recording AddOnResult Payload resource to delete.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteRecordingAddOnResultPayload(context.Background(), AccountSid, ReferenceSid, AddOnResultSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteRecordingAddOnResultPayload``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Recording AddOnResult Payload resources to delete.
**ReferenceSid** | **string** | The SID of the recording to which the AddOnResult resource that contains the payloads to delete belongs.
**AddOnResultSid** | **string** | The SID of the AddOnResult to which the payloads to delete belongs.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the Recording AddOnResult Payload resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeleteRecordingAddOnResultPayloadParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------





### Return type

 (empty response body)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRecordingTranscription

> DeleteRecordingTranscription(ctx, AccountSid, RecordingSid, Sid).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AccountSid := "AccountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Transcription resources to delete.
    RecordingSid := "RecordingSid_example" // string | The SID of the [Recording](https://www.twilio.com/docs/voice/api/recording) that created the transcription to delete.
    Sid := "Sid_example" // string | The Twilio-provided string that uniquely identifies the Transcription resource to delete.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteRecordingTranscription(context.Background(), AccountSid, RecordingSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteRecordingTranscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Transcription resources to delete.
**RecordingSid** | **string** | The SID of the [Recording](https://www.twilio.com/docs/voice/api/recording) that created the transcription to delete.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the Transcription resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeleteRecordingTranscriptionParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSigningKey

> DeleteSigningKey(ctx, AccountSid, Sid).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AccountSid := "AccountSid_example" // string | 
    Sid := "Sid_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteSigningKey(context.Background(), AccountSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteSigningKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string** | 
**Sid** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a DeleteSigningKeyParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSipAuthCallsCredentialListMapping

> DeleteSipAuthCallsCredentialListMapping(ctx, AccountSid, DomainSid, Sid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AccountSid := "AccountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the CredentialListMapping resources to delete.
    DomainSid := "DomainSid_example" // string | The SID of the SIP domain that contains the resource to delete.
    Sid := "Sid_example" // string | The Twilio-provided string that uniquely identifies the CredentialListMapping resource to delete.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteSipAuthCallsCredentialListMapping(context.Background(), AccountSid, DomainSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteSipAuthCallsCredentialListMapping``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the CredentialListMapping resources to delete.
**DomainSid** | **string** | The SID of the SIP domain that contains the resource to delete.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the CredentialListMapping resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeleteSipAuthCallsCredentialListMappingParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSipAuthCallsIpAccessControlListMapping

> DeleteSipAuthCallsIpAccessControlListMapping(ctx, AccountSid, DomainSid, Sid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AccountSid := "AccountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the IpAccessControlListMapping resources to delete.
    DomainSid := "DomainSid_example" // string | The SID of the SIP domain that contains the resources to delete.
    Sid := "Sid_example" // string | The Twilio-provided string that uniquely identifies the IpAccessControlListMapping resource to delete.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteSipAuthCallsIpAccessControlListMapping(context.Background(), AccountSid, DomainSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteSipAuthCallsIpAccessControlListMapping``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the IpAccessControlListMapping resources to delete.
**DomainSid** | **string** | The SID of the SIP domain that contains the resources to delete.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the IpAccessControlListMapping resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeleteSipAuthCallsIpAccessControlListMappingParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSipAuthRegistrationsCredentialListMapping

> DeleteSipAuthRegistrationsCredentialListMapping(ctx, AccountSid, DomainSid, Sid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AccountSid := "AccountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the CredentialListMapping resources to delete.
    DomainSid := "DomainSid_example" // string | The SID of the SIP domain that contains the resources to delete.
    Sid := "Sid_example" // string | The Twilio-provided string that uniquely identifies the CredentialListMapping resource to delete.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteSipAuthRegistrationsCredentialListMapping(context.Background(), AccountSid, DomainSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteSipAuthRegistrationsCredentialListMapping``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the CredentialListMapping resources to delete.
**DomainSid** | **string** | The SID of the SIP domain that contains the resources to delete.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the CredentialListMapping resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeleteSipAuthRegistrationsCredentialListMappingParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSipCredential

> DeleteSipCredential(ctx, AccountSid, CredentialListSid, Sid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AccountSid := "AccountSid_example" // string | The unique id of the Account that is responsible for this resource.
    CredentialListSid := "CredentialListSid_example" // string | The unique id that identifies the credential list that contains the desired credentials.
    Sid := "Sid_example" // string | The unique id that identifies the resource to delete.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteSipCredential(context.Background(), AccountSid, CredentialListSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteSipCredential``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string** | The unique id of the Account that is responsible for this resource.
**CredentialListSid** | **string** | The unique id that identifies the credential list that contains the desired credentials.
**Sid** | **string** | The unique id that identifies the resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeleteSipCredentialParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSipCredentialList

> DeleteSipCredentialList(ctx, AccountSid, Sid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AccountSid := "AccountSid_example" // string | The unique id of the Account that is responsible for this resource.
    Sid := "Sid_example" // string | The credential list Sid that uniquely identifies this resource

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteSipCredentialList(context.Background(), AccountSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteSipCredentialList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string** | The unique id of the Account that is responsible for this resource.
**Sid** | **string** | The credential list Sid that uniquely identifies this resource

### Other Parameters

Other parameters are passed through a pointer to a DeleteSipCredentialListParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSipCredentialListMapping

> DeleteSipCredentialListMapping(ctx, AccountSid, DomainSid, Sid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AccountSid := "AccountSid_example" // string | The unique id of the [Account](https://www.twilio.com/docs/iam/api/account) responsible for this resource.
    DomainSid := "DomainSid_example" // string | A 34 character string that uniquely identifies the SIP Domain that includes the resource to delete.
    Sid := "Sid_example" // string | A 34 character string that uniquely identifies the resource to delete.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteSipCredentialListMapping(context.Background(), AccountSid, DomainSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteSipCredentialListMapping``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string** | The unique id of the [Account](https://www.twilio.com/docs/iam/api/account) responsible for this resource.
**DomainSid** | **string** | A 34 character string that uniquely identifies the SIP Domain that includes the resource to delete.
**Sid** | **string** | A 34 character string that uniquely identifies the resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeleteSipCredentialListMappingParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSipDomain

> DeleteSipDomain(ctx, AccountSid, Sid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AccountSid := "AccountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the SipDomain resources to delete.
    Sid := "Sid_example" // string | The Twilio-provided string that uniquely identifies the SipDomain resource to delete.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteSipDomain(context.Background(), AccountSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteSipDomain``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the SipDomain resources to delete.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the SipDomain resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeleteSipDomainParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSipIpAccessControlList

> DeleteSipIpAccessControlList(ctx, AccountSid, Sid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AccountSid := "AccountSid_example" // string | The unique id of the [Account](https://www.twilio.com/docs/iam/api/account) responsible for this resource.
    Sid := "Sid_example" // string | A 34 character string that uniquely identifies the resource to delete.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteSipIpAccessControlList(context.Background(), AccountSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteSipIpAccessControlList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string** | The unique id of the [Account](https://www.twilio.com/docs/iam/api/account) responsible for this resource.
**Sid** | **string** | A 34 character string that uniquely identifies the resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeleteSipIpAccessControlListParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSipIpAccessControlListMapping

> DeleteSipIpAccessControlListMapping(ctx, AccountSid, DomainSid, Sid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AccountSid := "AccountSid_example" // string | The unique id of the Account that is responsible for this resource.
    DomainSid := "DomainSid_example" // string | A 34 character string that uniquely identifies the SIP domain.
    Sid := "Sid_example" // string | A 34 character string that uniquely identifies the resource to delete.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteSipIpAccessControlListMapping(context.Background(), AccountSid, DomainSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteSipIpAccessControlListMapping``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string** | The unique id of the Account that is responsible for this resource.
**DomainSid** | **string** | A 34 character string that uniquely identifies the SIP domain.
**Sid** | **string** | A 34 character string that uniquely identifies the resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeleteSipIpAccessControlListMappingParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSipIpAddress

> DeleteSipIpAddress(ctx, AccountSid, IpAccessControlListSid, Sid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AccountSid := "AccountSid_example" // string | The unique id of the [Account](https://www.twilio.com/docs/iam/api/account) responsible for this resource.
    IpAccessControlListSid := "IpAccessControlListSid_example" // string | The IpAccessControlList Sid that identifies the IpAddress resources to delete.
    Sid := "Sid_example" // string | A 34 character string that uniquely identifies the resource to delete.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteSipIpAddress(context.Background(), AccountSid, IpAccessControlListSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteSipIpAddress``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string** | The unique id of the [Account](https://www.twilio.com/docs/iam/api/account) responsible for this resource.
**IpAccessControlListSid** | **string** | The IpAccessControlList Sid that identifies the IpAddress resources to delete.
**Sid** | **string** | A 34 character string that uniquely identifies the resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeleteSipIpAddressParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTranscription

> DeleteTranscription(ctx, AccountSid, Sid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AccountSid := "AccountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Transcription resources to delete.
    Sid := "Sid_example" // string | The Twilio-provided string that uniquely identifies the Transcription resource to delete.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteTranscription(context.Background(), AccountSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteTranscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Transcription resources to delete.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the Transcription resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeleteTranscriptionParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteUsageTrigger

> DeleteUsageTrigger(ctx, AccountSid, Sid).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AccountSid := "AccountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the UsageTrigger resources to delete.
    Sid := "Sid_example" // string | The Twilio-provided string that uniquely identifies the UsageTrigger resource to delete.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteUsageTrigger(context.Background(), AccountSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteUsageTrigger``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the UsageTrigger resources to delete.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the UsageTrigger resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeleteUsageTriggerParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchAccount

> ApiV2010Account FetchAccount(ctx, Sid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    Sid := "Sid_example" // string | The Account Sid that uniquely identifies the account to fetch

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchAccount(context.Background(), Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchAccount`: ApiV2010Account
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The Account Sid that uniquely identifies the account to fetch

### Other Parameters

Other parameters are passed through a pointer to a FetchAccountParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


### Return type

[**ApiV2010Account**](ApiV2010Account.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchAddress

> ApiV2010AccountAddress FetchAddress(ctx, AccountSid, Sid).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AccountSid := "AccountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that is responsible for the Address resource to fetch.
    Sid := "Sid_example" // string | The Twilio-provided string that uniquely identifies the Address resource to fetch.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchAddress(context.Background(), AccountSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchAddress``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchAddress`: ApiV2010AccountAddress
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchAddress`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that is responsible for the Address resource to fetch.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the Address resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchAddressParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------



### Return type

[**ApiV2010AccountAddress**](ApiV2010AccountAddress.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchApplication

> ApiV2010AccountApplication FetchApplication(ctx, AccountSid, Sid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AccountSid := "AccountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Application resource to fetch.
    Sid := "Sid_example" // string | The Twilio-provided string that uniquely identifies the Application resource to fetch.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchApplication(context.Background(), AccountSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchApplication``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchApplication`: ApiV2010AccountApplication
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchApplication`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Application resource to fetch.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the Application resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchApplicationParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------



### Return type

[**ApiV2010AccountApplication**](ApiV2010AccountApplication.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchAuthorizedConnectApp

> ApiV2010AccountAuthorizedConnectApp FetchAuthorizedConnectApp(ctx, AccountSid, ConnectAppSid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AccountSid := "AccountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the AuthorizedConnectApp resource to fetch.
    ConnectAppSid := "ConnectAppSid_example" // string | The SID of the Connect App to fetch.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchAuthorizedConnectApp(context.Background(), AccountSid, ConnectAppSid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchAuthorizedConnectApp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchAuthorizedConnectApp`: ApiV2010AccountAuthorizedConnectApp
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchAuthorizedConnectApp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the AuthorizedConnectApp resource to fetch.
**ConnectAppSid** | **string** | The SID of the Connect App to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchAuthorizedConnectAppParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------



### Return type

[**ApiV2010AccountAuthorizedConnectApp**](ApiV2010AccountAuthorizedConnectApp.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchAvailablePhoneNumberCountry

> ApiV2010AccountAvailablePhoneNumberCountry FetchAvailablePhoneNumberCountry(ctx, AccountSid, CountryCode).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AccountSid := "AccountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) requesting the available phone number Country resource.
    CountryCode := "CountryCode_example" // string | The [ISO-3166-1](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country code of the country to fetch available phone number information about.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchAvailablePhoneNumberCountry(context.Background(), AccountSid, CountryCode).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchAvailablePhoneNumberCountry``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchAvailablePhoneNumberCountry`: ApiV2010AccountAvailablePhoneNumberCountry
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchAvailablePhoneNumberCountry`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) requesting the available phone number Country resource.
**CountryCode** | **string** | The [ISO-3166-1](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country code of the country to fetch available phone number information about.

### Other Parameters

Other parameters are passed through a pointer to a FetchAvailablePhoneNumberCountryParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------



### Return type

[**ApiV2010AccountAvailablePhoneNumberCountry**](ApiV2010AccountAvailablePhoneNumberCountry.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchBalance

> ApiV2010AccountBalance FetchBalance(ctx, AccountSid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AccountSid := "AccountSid_example" // string | The unique SID identifier of the Account.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchBalance(context.Background(), AccountSid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchBalance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchBalance`: ApiV2010AccountBalance
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchBalance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string** | The unique SID identifier of the Account.

### Other Parameters

Other parameters are passed through a pointer to a FetchBalanceParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


### Return type

[**ApiV2010AccountBalance**](ApiV2010AccountBalance.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchCall

> ApiV2010AccountCall FetchCall(ctx, AccountSid, Sid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AccountSid := "AccountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Call resource(s) to fetch.
    Sid := "Sid_example" // string | The SID of the Call resource to fetch.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchCall(context.Background(), AccountSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchCall``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchCall`: ApiV2010AccountCall
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchCall`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Call resource(s) to fetch.
**Sid** | **string** | The SID of the Call resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchCallParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------



### Return type

[**ApiV2010AccountCall**](ApiV2010AccountCall.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchCallFeedback

> ApiV2010AccountCallCallFeedback FetchCallFeedback(ctx, AccountSid, CallSid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AccountSid := "AccountSid_example" // string | The unique id of the [Account](https://www.twilio.com/docs/iam/api/account) responsible for this resource.
    CallSid := "CallSid_example" // string | The call sid that uniquely identifies the call

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchCallFeedback(context.Background(), AccountSid, CallSid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchCallFeedback``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchCallFeedback`: ApiV2010AccountCallCallFeedback
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchCallFeedback`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string** | The unique id of the [Account](https://www.twilio.com/docs/iam/api/account) responsible for this resource.
**CallSid** | **string** | The call sid that uniquely identifies the call

### Other Parameters

Other parameters are passed through a pointer to a FetchCallFeedbackParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------



### Return type

[**ApiV2010AccountCallCallFeedback**](ApiV2010AccountCallCallFeedback.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchCallFeedbackSummary

> ApiV2010AccountCallCallFeedbackSummary FetchCallFeedbackSummary(ctx, AccountSid, Sid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AccountSid := "AccountSid_example" // string | The unique id of the [Account](https://www.twilio.com/docs/iam/api/account) responsible for this resource.
    Sid := "Sid_example" // string | A 34 character string that uniquely identifies this resource.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchCallFeedbackSummary(context.Background(), AccountSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchCallFeedbackSummary``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchCallFeedbackSummary`: ApiV2010AccountCallCallFeedbackSummary
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchCallFeedbackSummary`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string** | The unique id of the [Account](https://www.twilio.com/docs/iam/api/account) responsible for this resource.
**Sid** | **string** | A 34 character string that uniquely identifies this resource.

### Other Parameters

Other parameters are passed through a pointer to a FetchCallFeedbackSummaryParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------



### Return type

[**ApiV2010AccountCallCallFeedbackSummary**](ApiV2010AccountCallCallFeedbackSummary.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchCallNotification

> ApiV2010AccountCallCallNotificationInstance FetchCallNotification(ctx, AccountSid, CallSid, Sid).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AccountSid := "AccountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Call Notification resource to fetch.
    CallSid := "CallSid_example" // string | The [Call](https://www.twilio.com/docs/voice/api/call-resource) SID of the Call Notification resource to fetch.
    Sid := "Sid_example" // string | The Twilio-provided string that uniquely identifies the Call Notification resource to fetch.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchCallNotification(context.Background(), AccountSid, CallSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchCallNotification``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchCallNotification`: ApiV2010AccountCallCallNotificationInstance
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchCallNotification`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Call Notification resource to fetch.
**CallSid** | **string** | The [Call](https://www.twilio.com/docs/voice/api/call-resource) SID of the Call Notification resource to fetch.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the Call Notification resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchCallNotificationParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------




### Return type

[**ApiV2010AccountCallCallNotificationInstance**](ApiV2010AccountCallCallNotificationInstance.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchCallRecording

> ApiV2010AccountCallCallRecording FetchCallRecording(ctx, AccountSid, CallSid, Sid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AccountSid := "AccountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Recording resource to fetch.
    CallSid := "CallSid_example" // string | The [Call](https://www.twilio.com/docs/voice/api/call-resource) SID of the resource to fetch.
    Sid := "Sid_example" // string | The Twilio-provided string that uniquely identifies the Recording resource to fetch.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchCallRecording(context.Background(), AccountSid, CallSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchCallRecording``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchCallRecording`: ApiV2010AccountCallCallRecording
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchCallRecording`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Recording resource to fetch.
**CallSid** | **string** | The [Call](https://www.twilio.com/docs/voice/api/call-resource) SID of the resource to fetch.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the Recording resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchCallRecordingParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------




### Return type

[**ApiV2010AccountCallCallRecording**](ApiV2010AccountCallCallRecording.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchConference

> ApiV2010AccountConference FetchConference(ctx, AccountSid, Sid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AccountSid := "AccountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Conference resource(s) to fetch.
    Sid := "Sid_example" // string | The Twilio-provided string that uniquely identifies the Conference resource to fetch

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchConference(context.Background(), AccountSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchConference``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchConference`: ApiV2010AccountConference
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchConference`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Conference resource(s) to fetch.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the Conference resource to fetch

### Other Parameters

Other parameters are passed through a pointer to a FetchConferenceParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------



### Return type

[**ApiV2010AccountConference**](ApiV2010AccountConference.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchConferenceRecording

> ApiV2010AccountConferenceConferenceRecording FetchConferenceRecording(ctx, AccountSid, ConferenceSid, Sid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AccountSid := "AccountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Conference Recording resource to fetch.
    ConferenceSid := "ConferenceSid_example" // string | The Conference SID that identifies the conference associated with the recording to fetch.
    Sid := "Sid_example" // string | The Twilio-provided string that uniquely identifies the Conference Recording resource to fetch.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchConferenceRecording(context.Background(), AccountSid, ConferenceSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchConferenceRecording``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchConferenceRecording`: ApiV2010AccountConferenceConferenceRecording
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchConferenceRecording`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Conference Recording resource to fetch.
**ConferenceSid** | **string** | The Conference SID that identifies the conference associated with the recording to fetch.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the Conference Recording resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchConferenceRecordingParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------




### Return type

[**ApiV2010AccountConferenceConferenceRecording**](ApiV2010AccountConferenceConferenceRecording.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchConnectApp

> ApiV2010AccountConnectApp FetchConnectApp(ctx, AccountSid, Sid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AccountSid := "AccountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the ConnectApp resource to fetch.
    Sid := "Sid_example" // string | The Twilio-provided string that uniquely identifies the ConnectApp resource to fetch.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchConnectApp(context.Background(), AccountSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchConnectApp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchConnectApp`: ApiV2010AccountConnectApp
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchConnectApp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the ConnectApp resource to fetch.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the ConnectApp resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchConnectAppParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------



### Return type

[**ApiV2010AccountConnectApp**](ApiV2010AccountConnectApp.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchIncomingPhoneNumber

> ApiV2010AccountIncomingPhoneNumber FetchIncomingPhoneNumber(ctx, AccountSid, Sid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AccountSid := "AccountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the IncomingPhoneNumber resource to fetch.
    Sid := "Sid_example" // string | The Twilio-provided string that uniquely identifies the IncomingPhoneNumber resource to fetch.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchIncomingPhoneNumber(context.Background(), AccountSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchIncomingPhoneNumber``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchIncomingPhoneNumber`: ApiV2010AccountIncomingPhoneNumber
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchIncomingPhoneNumber`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the IncomingPhoneNumber resource to fetch.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the IncomingPhoneNumber resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchIncomingPhoneNumberParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------



### Return type

[**ApiV2010AccountIncomingPhoneNumber**](ApiV2010AccountIncomingPhoneNumber.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchIncomingPhoneNumberAssignedAddOn

> ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberAssignedAddOn FetchIncomingPhoneNumberAssignedAddOn(ctx, AccountSid, ResourceSid, Sid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AccountSid := "AccountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the resource to fetch.
    ResourceSid := "ResourceSid_example" // string | The SID of the Phone Number to which the Add-on is assigned.
    Sid := "Sid_example" // string | The Twilio-provided string that uniquely identifies the resource to fetch.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchIncomingPhoneNumberAssignedAddOn(context.Background(), AccountSid, ResourceSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchIncomingPhoneNumberAssignedAddOn``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchIncomingPhoneNumberAssignedAddOn`: ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberAssignedAddOn
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchIncomingPhoneNumberAssignedAddOn`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the resource to fetch.
**ResourceSid** | **string** | The SID of the Phone Number to which the Add-on is assigned.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchIncomingPhoneNumberAssignedAddOnParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------




### Return type

[**ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberAssignedAddOn**](ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberAssignedAddOn.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchIncomingPhoneNumberAssignedAddOnExtension

> ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberAssignedAddOnIncomingPhoneNumberAssignedAddOnExtension FetchIncomingPhoneNumberAssignedAddOnExtension(ctx, AccountSid, ResourceSid, AssignedAddOnSid, Sid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AccountSid := "AccountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the resource to fetch.
    ResourceSid := "ResourceSid_example" // string | The SID of the Phone Number to which the Add-on is assigned.
    AssignedAddOnSid := "AssignedAddOnSid_example" // string | The SID that uniquely identifies the assigned Add-on installation.
    Sid := "Sid_example" // string | The Twilio-provided string that uniquely identifies the resource to fetch.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchIncomingPhoneNumberAssignedAddOnExtension(context.Background(), AccountSid, ResourceSid, AssignedAddOnSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchIncomingPhoneNumberAssignedAddOnExtension``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchIncomingPhoneNumberAssignedAddOnExtension`: ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberAssignedAddOnIncomingPhoneNumberAssignedAddOnExtension
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchIncomingPhoneNumberAssignedAddOnExtension`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the resource to fetch.
**ResourceSid** | **string** | The SID of the Phone Number to which the Add-on is assigned.
**AssignedAddOnSid** | **string** | The SID that uniquely identifies the assigned Add-on installation.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchIncomingPhoneNumberAssignedAddOnExtensionParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------





### Return type

[**ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberAssignedAddOnIncomingPhoneNumberAssignedAddOnExtension**](ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberAssignedAddOnIncomingPhoneNumberAssignedAddOnExtension.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchKey

> ApiV2010AccountKey FetchKey(ctx, AccountSid, Sid).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AccountSid := "AccountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Key resource to fetch.
    Sid := "Sid_example" // string | The Twilio-provided string that uniquely identifies the Key resource to fetch.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchKey(context.Background(), AccountSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchKey`: ApiV2010AccountKey
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Key resource to fetch.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the Key resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchKeyParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------



### Return type

[**ApiV2010AccountKey**](ApiV2010AccountKey.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchMedia

> ApiV2010AccountMessageMedia FetchMedia(ctx, AccountSid, MessageSid, Sid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AccountSid := "AccountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Media resource(s) to fetch.
    MessageSid := "MessageSid_example" // string | The SID of the Message resource that this Media resource belongs to.
    Sid := "Sid_example" // string | The Twilio-provided string that uniquely identifies the Media resource to fetch

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchMedia(context.Background(), AccountSid, MessageSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchMedia``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchMedia`: ApiV2010AccountMessageMedia
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchMedia`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Media resource(s) to fetch.
**MessageSid** | **string** | The SID of the Message resource that this Media resource belongs to.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the Media resource to fetch

### Other Parameters

Other parameters are passed through a pointer to a FetchMediaParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------




### Return type

[**ApiV2010AccountMessageMedia**](ApiV2010AccountMessageMedia.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchMember

> ApiV2010AccountQueueMember FetchMember(ctx, AccountSid, QueueSid, CallSid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AccountSid := "AccountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Member resource(s) to fetch.
    QueueSid := "QueueSid_example" // string | The SID of the Queue in which to find the members to fetch.
    CallSid := "CallSid_example" // string | The [Call](https://www.twilio.com/docs/voice/api/call-resource) SID of the resource(s) to fetch.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchMember(context.Background(), AccountSid, QueueSid, CallSid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchMember``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchMember`: ApiV2010AccountQueueMember
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchMember`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Member resource(s) to fetch.
**QueueSid** | **string** | The SID of the Queue in which to find the members to fetch.
**CallSid** | **string** | The [Call](https://www.twilio.com/docs/voice/api/call-resource) SID of the resource(s) to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchMemberParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------




### Return type

[**ApiV2010AccountQueueMember**](ApiV2010AccountQueueMember.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchMessage

> ApiV2010AccountMessage FetchMessage(ctx, AccountSid, Sid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AccountSid := "AccountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Message resource to fetch.
    Sid := "Sid_example" // string | The Twilio-provided string that uniquely identifies the Message resource to fetch.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchMessage(context.Background(), AccountSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchMessage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchMessage`: ApiV2010AccountMessage
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchMessage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Message resource to fetch.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the Message resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchMessageParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------



### Return type

[**ApiV2010AccountMessage**](ApiV2010AccountMessage.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchNotification

> ApiV2010AccountNotificationInstance FetchNotification(ctx, AccountSid, Sid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AccountSid := "AccountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Notification resource to fetch.
    Sid := "Sid_example" // string | The Twilio-provided string that uniquely identifies the Notification resource to fetch.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchNotification(context.Background(), AccountSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchNotification``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchNotification`: ApiV2010AccountNotificationInstance
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchNotification`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Notification resource to fetch.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the Notification resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchNotificationParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------



### Return type

[**ApiV2010AccountNotificationInstance**](ApiV2010AccountNotificationInstance.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchOutgoingCallerId

> ApiV2010AccountOutgoingCallerId FetchOutgoingCallerId(ctx, AccountSid, Sid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AccountSid := "AccountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the OutgoingCallerId resource to fetch.
    Sid := "Sid_example" // string | The Twilio-provided string that uniquely identifies the OutgoingCallerId resource to fetch.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchOutgoingCallerId(context.Background(), AccountSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchOutgoingCallerId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchOutgoingCallerId`: ApiV2010AccountOutgoingCallerId
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchOutgoingCallerId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the OutgoingCallerId resource to fetch.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the OutgoingCallerId resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchOutgoingCallerIdParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------



### Return type

[**ApiV2010AccountOutgoingCallerId**](ApiV2010AccountOutgoingCallerId.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchParticipant

> ApiV2010AccountConferenceParticipant FetchParticipant(ctx, AccountSid, ConferenceSid, CallSid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AccountSid := "AccountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Participant resource to fetch.
    ConferenceSid := "ConferenceSid_example" // string | The SID of the conference with the participant to fetch.
    CallSid := "CallSid_example" // string | The [Call](https://www.twilio.com/docs/voice/api/call-resource) SID or label of the participant to fetch. Non URL safe characters in a label must be percent encoded, for example, a space character is represented as %20.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchParticipant(context.Background(), AccountSid, ConferenceSid, CallSid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchParticipant``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchParticipant`: ApiV2010AccountConferenceParticipant
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchParticipant`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Participant resource to fetch.
**ConferenceSid** | **string** | The SID of the conference with the participant to fetch.
**CallSid** | **string** | The [Call](https://www.twilio.com/docs/voice/api/call-resource) SID or label of the participant to fetch. Non URL safe characters in a label must be percent encoded, for example, a space character is represented as %20.

### Other Parameters

Other parameters are passed through a pointer to a FetchParticipantParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------




### Return type

[**ApiV2010AccountConferenceParticipant**](ApiV2010AccountConferenceParticipant.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchQueue

> ApiV2010AccountQueue FetchQueue(ctx, AccountSid, Sid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AccountSid := "AccountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Queue resource to fetch.
    Sid := "Sid_example" // string | The Twilio-provided string that uniquely identifies the Queue resource to fetch

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchQueue(context.Background(), AccountSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchQueue``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchQueue`: ApiV2010AccountQueue
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchQueue`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Queue resource to fetch.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the Queue resource to fetch

### Other Parameters

Other parameters are passed through a pointer to a FetchQueueParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------



### Return type

[**ApiV2010AccountQueue**](ApiV2010AccountQueue.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchRecording

> ApiV2010AccountRecording FetchRecording(ctx, AccountSid, Sid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AccountSid := "AccountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Recording resource to fetch.
    Sid := "Sid_example" // string | The Twilio-provided string that uniquely identifies the Recording resource to fetch.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchRecording(context.Background(), AccountSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchRecording``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchRecording`: ApiV2010AccountRecording
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchRecording`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Recording resource to fetch.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the Recording resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchRecordingParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------



### Return type

[**ApiV2010AccountRecording**](ApiV2010AccountRecording.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchRecordingAddOnResult

> ApiV2010AccountRecordingRecordingAddOnResult FetchRecordingAddOnResult(ctx, AccountSid, ReferenceSid, Sid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AccountSid := "AccountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Recording AddOnResult resource to fetch.
    ReferenceSid := "ReferenceSid_example" // string | The SID of the recording to which the result to fetch belongs.
    Sid := "Sid_example" // string | The Twilio-provided string that uniquely identifies the Recording AddOnResult resource to fetch.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchRecordingAddOnResult(context.Background(), AccountSid, ReferenceSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchRecordingAddOnResult``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchRecordingAddOnResult`: ApiV2010AccountRecordingRecordingAddOnResult
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchRecordingAddOnResult`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Recording AddOnResult resource to fetch.
**ReferenceSid** | **string** | The SID of the recording to which the result to fetch belongs.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the Recording AddOnResult resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchRecordingAddOnResultParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------




### Return type

[**ApiV2010AccountRecordingRecordingAddOnResult**](ApiV2010AccountRecordingRecordingAddOnResult.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchRecordingAddOnResultPayload

> ApiV2010AccountRecordingRecordingAddOnResultRecordingAddOnResultPayload FetchRecordingAddOnResultPayload(ctx, AccountSid, ReferenceSid, AddOnResultSid, Sid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AccountSid := "AccountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Recording AddOnResult Payload resource to fetch.
    ReferenceSid := "ReferenceSid_example" // string | The SID of the recording to which the AddOnResult resource that contains the payload to fetch belongs.
    AddOnResultSid := "AddOnResultSid_example" // string | The SID of the AddOnResult to which the payload to fetch belongs.
    Sid := "Sid_example" // string | The Twilio-provided string that uniquely identifies the Recording AddOnResult Payload resource to fetch.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchRecordingAddOnResultPayload(context.Background(), AccountSid, ReferenceSid, AddOnResultSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchRecordingAddOnResultPayload``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchRecordingAddOnResultPayload`: ApiV2010AccountRecordingRecordingAddOnResultRecordingAddOnResultPayload
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchRecordingAddOnResultPayload`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Recording AddOnResult Payload resource to fetch.
**ReferenceSid** | **string** | The SID of the recording to which the AddOnResult resource that contains the payload to fetch belongs.
**AddOnResultSid** | **string** | The SID of the AddOnResult to which the payload to fetch belongs.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the Recording AddOnResult Payload resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchRecordingAddOnResultPayloadParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------





### Return type

[**ApiV2010AccountRecordingRecordingAddOnResultRecordingAddOnResultPayload**](ApiV2010AccountRecordingRecordingAddOnResultRecordingAddOnResultPayload.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchRecordingTranscription

> ApiV2010AccountRecordingRecordingTranscription FetchRecordingTranscription(ctx, AccountSid, RecordingSid, Sid).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AccountSid := "AccountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Transcription resource to fetch.
    RecordingSid := "RecordingSid_example" // string | The SID of the [Recording](https://www.twilio.com/docs/voice/api/recording) that created the transcription to fetch.
    Sid := "Sid_example" // string | The Twilio-provided string that uniquely identifies the Transcription resource to fetch.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchRecordingTranscription(context.Background(), AccountSid, RecordingSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchRecordingTranscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchRecordingTranscription`: ApiV2010AccountRecordingRecordingTranscription
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchRecordingTranscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Transcription resource to fetch.
**RecordingSid** | **string** | The SID of the [Recording](https://www.twilio.com/docs/voice/api/recording) that created the transcription to fetch.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the Transcription resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchRecordingTranscriptionParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------




### Return type

[**ApiV2010AccountRecordingRecordingTranscription**](ApiV2010AccountRecordingRecordingTranscription.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchShortCode

> ApiV2010AccountShortCode FetchShortCode(ctx, AccountSid, Sid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AccountSid := "AccountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the ShortCode resource(s) to fetch.
    Sid := "Sid_example" // string | The Twilio-provided string that uniquely identifies the ShortCode resource to fetch

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchShortCode(context.Background(), AccountSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchShortCode``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchShortCode`: ApiV2010AccountShortCode
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchShortCode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the ShortCode resource(s) to fetch.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the ShortCode resource to fetch

### Other Parameters

Other parameters are passed through a pointer to a FetchShortCodeParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------



### Return type

[**ApiV2010AccountShortCode**](ApiV2010AccountShortCode.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchSigningKey

> ApiV2010AccountSigningKey FetchSigningKey(ctx, AccountSid, Sid).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AccountSid := "AccountSid_example" // string | 
    Sid := "Sid_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchSigningKey(context.Background(), AccountSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchSigningKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchSigningKey`: ApiV2010AccountSigningKey
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchSigningKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string** | 
**Sid** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a FetchSigningKeyParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------



### Return type

[**ApiV2010AccountSigningKey**](ApiV2010AccountSigningKey.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchSipAuthCallsCredentialListMapping

> ApiV2010AccountSipSipDomainSipAuthSipAuthCallsSipAuthCallsCredentialListMapping FetchSipAuthCallsCredentialListMapping(ctx, AccountSid, DomainSid, Sid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AccountSid := "AccountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the CredentialListMapping resource to fetch.
    DomainSid := "DomainSid_example" // string | The SID of the SIP domain that contains the resource to fetch.
    Sid := "Sid_example" // string | The Twilio-provided string that uniquely identifies the CredentialListMapping resource to fetch.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchSipAuthCallsCredentialListMapping(context.Background(), AccountSid, DomainSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchSipAuthCallsCredentialListMapping``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchSipAuthCallsCredentialListMapping`: ApiV2010AccountSipSipDomainSipAuthSipAuthCallsSipAuthCallsCredentialListMapping
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchSipAuthCallsCredentialListMapping`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the CredentialListMapping resource to fetch.
**DomainSid** | **string** | The SID of the SIP domain that contains the resource to fetch.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the CredentialListMapping resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchSipAuthCallsCredentialListMappingParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------




### Return type

[**ApiV2010AccountSipSipDomainSipAuthSipAuthCallsSipAuthCallsCredentialListMapping**](ApiV2010AccountSipSipDomainSipAuthSipAuthCallsSipAuthCallsCredentialListMapping.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchSipAuthCallsIpAccessControlListMapping

> ApiV2010AccountSipSipDomainSipAuthSipAuthCallsSipAuthCallsIpAccessControlListMapping FetchSipAuthCallsIpAccessControlListMapping(ctx, AccountSid, DomainSid, Sid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AccountSid := "AccountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the IpAccessControlListMapping resource to fetch.
    DomainSid := "DomainSid_example" // string | The SID of the SIP domain that contains the resource to fetch.
    Sid := "Sid_example" // string | The Twilio-provided string that uniquely identifies the IpAccessControlListMapping resource to fetch.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchSipAuthCallsIpAccessControlListMapping(context.Background(), AccountSid, DomainSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchSipAuthCallsIpAccessControlListMapping``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchSipAuthCallsIpAccessControlListMapping`: ApiV2010AccountSipSipDomainSipAuthSipAuthCallsSipAuthCallsIpAccessControlListMapping
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchSipAuthCallsIpAccessControlListMapping`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the IpAccessControlListMapping resource to fetch.
**DomainSid** | **string** | The SID of the SIP domain that contains the resource to fetch.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the IpAccessControlListMapping resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchSipAuthCallsIpAccessControlListMappingParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------




### Return type

[**ApiV2010AccountSipSipDomainSipAuthSipAuthCallsSipAuthCallsIpAccessControlListMapping**](ApiV2010AccountSipSipDomainSipAuthSipAuthCallsSipAuthCallsIpAccessControlListMapping.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchSipAuthRegistrationsCredentialListMapping

> ApiV2010AccountSipSipDomainSipAuthSipAuthRegistrationsSipAuthRegistrationsCredentialListMapping FetchSipAuthRegistrationsCredentialListMapping(ctx, AccountSid, DomainSid, Sid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AccountSid := "AccountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the CredentialListMapping resource to fetch.
    DomainSid := "DomainSid_example" // string | The SID of the SIP domain that contains the resource to fetch.
    Sid := "Sid_example" // string | The Twilio-provided string that uniquely identifies the CredentialListMapping resource to fetch.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchSipAuthRegistrationsCredentialListMapping(context.Background(), AccountSid, DomainSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchSipAuthRegistrationsCredentialListMapping``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchSipAuthRegistrationsCredentialListMapping`: ApiV2010AccountSipSipDomainSipAuthSipAuthRegistrationsSipAuthRegistrationsCredentialListMapping
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchSipAuthRegistrationsCredentialListMapping`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the CredentialListMapping resource to fetch.
**DomainSid** | **string** | The SID of the SIP domain that contains the resource to fetch.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the CredentialListMapping resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchSipAuthRegistrationsCredentialListMappingParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------




### Return type

[**ApiV2010AccountSipSipDomainSipAuthSipAuthRegistrationsSipAuthRegistrationsCredentialListMapping**](ApiV2010AccountSipSipDomainSipAuthSipAuthRegistrationsSipAuthRegistrationsCredentialListMapping.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchSipCredential

> ApiV2010AccountSipSipCredentialListSipCredential FetchSipCredential(ctx, AccountSid, CredentialListSid, Sid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AccountSid := "AccountSid_example" // string | The unique id of the Account that is responsible for this resource.
    CredentialListSid := "CredentialListSid_example" // string | The unique id that identifies the credential list that contains the desired credential.
    Sid := "Sid_example" // string | The unique id that identifies the resource to fetch.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchSipCredential(context.Background(), AccountSid, CredentialListSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchSipCredential``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchSipCredential`: ApiV2010AccountSipSipCredentialListSipCredential
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchSipCredential`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string** | The unique id of the Account that is responsible for this resource.
**CredentialListSid** | **string** | The unique id that identifies the credential list that contains the desired credential.
**Sid** | **string** | The unique id that identifies the resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchSipCredentialParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------




### Return type

[**ApiV2010AccountSipSipCredentialListSipCredential**](ApiV2010AccountSipSipCredentialListSipCredential.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchSipCredentialList

> ApiV2010AccountSipSipCredentialList FetchSipCredentialList(ctx, AccountSid, Sid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AccountSid := "AccountSid_example" // string | The unique id of the Account that is responsible for this resource.
    Sid := "Sid_example" // string | The credential list Sid that uniquely identifies this resource

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchSipCredentialList(context.Background(), AccountSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchSipCredentialList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchSipCredentialList`: ApiV2010AccountSipSipCredentialList
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchSipCredentialList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string** | The unique id of the Account that is responsible for this resource.
**Sid** | **string** | The credential list Sid that uniquely identifies this resource

### Other Parameters

Other parameters are passed through a pointer to a FetchSipCredentialListParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------



### Return type

[**ApiV2010AccountSipSipCredentialList**](ApiV2010AccountSipSipCredentialList.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchSipCredentialListMapping

> ApiV2010AccountSipSipDomainSipCredentialListMapping FetchSipCredentialListMapping(ctx, AccountSid, DomainSid, Sid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AccountSid := "AccountSid_example" // string | The unique id of the [Account](https://www.twilio.com/docs/iam/api/account) responsible for this resource.
    DomainSid := "DomainSid_example" // string | A 34 character string that uniquely identifies the SIP Domain that includes the resource to fetch.
    Sid := "Sid_example" // string | A 34 character string that uniquely identifies the resource to fetch.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchSipCredentialListMapping(context.Background(), AccountSid, DomainSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchSipCredentialListMapping``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchSipCredentialListMapping`: ApiV2010AccountSipSipDomainSipCredentialListMapping
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchSipCredentialListMapping`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string** | The unique id of the [Account](https://www.twilio.com/docs/iam/api/account) responsible for this resource.
**DomainSid** | **string** | A 34 character string that uniquely identifies the SIP Domain that includes the resource to fetch.
**Sid** | **string** | A 34 character string that uniquely identifies the resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchSipCredentialListMappingParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------




### Return type

[**ApiV2010AccountSipSipDomainSipCredentialListMapping**](ApiV2010AccountSipSipDomainSipCredentialListMapping.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchSipDomain

> ApiV2010AccountSipSipDomain FetchSipDomain(ctx, AccountSid, Sid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AccountSid := "AccountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the SipDomain resource to fetch.
    Sid := "Sid_example" // string | The Twilio-provided string that uniquely identifies the SipDomain resource to fetch.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchSipDomain(context.Background(), AccountSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchSipDomain``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchSipDomain`: ApiV2010AccountSipSipDomain
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchSipDomain`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the SipDomain resource to fetch.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the SipDomain resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchSipDomainParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------



### Return type

[**ApiV2010AccountSipSipDomain**](ApiV2010AccountSipSipDomain.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchSipIpAccessControlList

> ApiV2010AccountSipSipIpAccessControlList FetchSipIpAccessControlList(ctx, AccountSid, Sid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AccountSid := "AccountSid_example" // string | The unique id of the [Account](https://www.twilio.com/docs/iam/api/account) responsible for this resource.
    Sid := "Sid_example" // string | A 34 character string that uniquely identifies the resource to fetch.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchSipIpAccessControlList(context.Background(), AccountSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchSipIpAccessControlList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchSipIpAccessControlList`: ApiV2010AccountSipSipIpAccessControlList
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchSipIpAccessControlList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string** | The unique id of the [Account](https://www.twilio.com/docs/iam/api/account) responsible for this resource.
**Sid** | **string** | A 34 character string that uniquely identifies the resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchSipIpAccessControlListParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------



### Return type

[**ApiV2010AccountSipSipIpAccessControlList**](ApiV2010AccountSipSipIpAccessControlList.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchSipIpAccessControlListMapping

> ApiV2010AccountSipSipDomainSipIpAccessControlListMapping FetchSipIpAccessControlListMapping(ctx, AccountSid, DomainSid, Sid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AccountSid := "AccountSid_example" // string | The unique id of the Account that is responsible for this resource.
    DomainSid := "DomainSid_example" // string | A 34 character string that uniquely identifies the SIP domain.
    Sid := "Sid_example" // string | A 34 character string that uniquely identifies the resource to fetch.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchSipIpAccessControlListMapping(context.Background(), AccountSid, DomainSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchSipIpAccessControlListMapping``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchSipIpAccessControlListMapping`: ApiV2010AccountSipSipDomainSipIpAccessControlListMapping
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchSipIpAccessControlListMapping`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string** | The unique id of the Account that is responsible for this resource.
**DomainSid** | **string** | A 34 character string that uniquely identifies the SIP domain.
**Sid** | **string** | A 34 character string that uniquely identifies the resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchSipIpAccessControlListMappingParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------




### Return type

[**ApiV2010AccountSipSipDomainSipIpAccessControlListMapping**](ApiV2010AccountSipSipDomainSipIpAccessControlListMapping.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchSipIpAddress

> ApiV2010AccountSipSipIpAccessControlListSipIpAddress FetchSipIpAddress(ctx, AccountSid, IpAccessControlListSid, Sid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AccountSid := "AccountSid_example" // string | The unique id of the [Account](https://www.twilio.com/docs/iam/api/account) responsible for this resource.
    IpAccessControlListSid := "IpAccessControlListSid_example" // string | The IpAccessControlList Sid that identifies the IpAddress resources to fetch.
    Sid := "Sid_example" // string | A 34 character string that uniquely identifies the IpAddress resource to fetch.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchSipIpAddress(context.Background(), AccountSid, IpAccessControlListSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchSipIpAddress``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchSipIpAddress`: ApiV2010AccountSipSipIpAccessControlListSipIpAddress
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchSipIpAddress`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string** | The unique id of the [Account](https://www.twilio.com/docs/iam/api/account) responsible for this resource.
**IpAccessControlListSid** | **string** | The IpAccessControlList Sid that identifies the IpAddress resources to fetch.
**Sid** | **string** | A 34 character string that uniquely identifies the IpAddress resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchSipIpAddressParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------




### Return type

[**ApiV2010AccountSipSipIpAccessControlListSipIpAddress**](ApiV2010AccountSipSipIpAccessControlListSipIpAddress.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchTranscription

> ApiV2010AccountTranscription FetchTranscription(ctx, AccountSid, Sid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AccountSid := "AccountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Transcription resource to fetch.
    Sid := "Sid_example" // string | The Twilio-provided string that uniquely identifies the Transcription resource to fetch.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchTranscription(context.Background(), AccountSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchTranscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchTranscription`: ApiV2010AccountTranscription
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchTranscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Transcription resource to fetch.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the Transcription resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchTranscriptionParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------



### Return type

[**ApiV2010AccountTranscription**](ApiV2010AccountTranscription.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchUsageTrigger

> ApiV2010AccountUsageUsageTrigger FetchUsageTrigger(ctx, AccountSid, Sid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AccountSid := "AccountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the UsageTrigger resource to fetch.
    Sid := "Sid_example" // string | The Twilio-provided string that uniquely identifies the UsageTrigger resource to fetch.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchUsageTrigger(context.Background(), AccountSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchUsageTrigger``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchUsageTrigger`: ApiV2010AccountUsageUsageTrigger
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchUsageTrigger`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the UsageTrigger resource to fetch.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the UsageTrigger resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchUsageTriggerParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------



### Return type

[**ApiV2010AccountUsageUsageTrigger**](ApiV2010AccountUsageUsageTrigger.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAccount

> ListAccountResponse ListAccount(ctx).FriendlyName(FriendlyName).Status(Status).PageSize(PageSize).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    FriendlyName := "FriendlyName_example" // string | Only return the Account resources with friendly names that exactly match this name. (optional)
    Status := "Status_example" // string | Only return Account resources with the given status. Can be `closed`, `suspended` or `active`. (optional)
    PageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListAccount(context.Background()).FriendlyName(FriendlyName).Status(Status).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAccount`: ListAccountResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListAccount`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListAccountParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------
 **FriendlyName** | **string** | Only return the Account resources with friendly names that exactly match this name.
 **Status** | **string** | Only return Account resources with the given status. Can be &#x60;closed&#x60;, &#x60;suspended&#x60; or &#x60;active&#x60;.
 **PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListAccountResponse**](ListAccountResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAddress

> ListAddressResponse ListAddress(ctx, AccountSid).CustomerName(CustomerName).FriendlyName(FriendlyName).IsoCountry(IsoCountry).PageSize(PageSize).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AccountSid := "AccountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that is responsible for the Address resource to read.
    CustomerName := "CustomerName_example" // string | The `customer_name` of the Address resources to read. (optional)
    FriendlyName := "FriendlyName_example" // string | The string that identifies the Address resources to read. (optional)
    IsoCountry := "IsoCountry_example" // string | The ISO country code of the Address resources to read. (optional)
    PageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListAddress(context.Background(), AccountSid).CustomerName(CustomerName).FriendlyName(FriendlyName).IsoCountry(IsoCountry).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListAddress``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAddress`: ListAddressResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListAddress`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that is responsible for the Address resource to read.

### Other Parameters

Other parameters are passed through a pointer to a ListAddressParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **CustomerName** | **string** | The &#x60;customer_name&#x60; of the Address resources to read.
 **FriendlyName** | **string** | The string that identifies the Address resources to read.
 **IsoCountry** | **string** | The ISO country code of the Address resources to read.
 **PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListAddressResponse**](ListAddressResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListApplication

> ListApplicationResponse ListApplication(ctx, AccountSid).FriendlyName(FriendlyName).PageSize(PageSize).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AccountSid := "AccountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Application resources to read.
    FriendlyName := "FriendlyName_example" // string | The string that identifies the Application resources to read. (optional)
    PageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListApplication(context.Background(), AccountSid).FriendlyName(FriendlyName).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListApplication``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListApplication`: ListApplicationResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListApplication`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Application resources to read.

### Other Parameters

Other parameters are passed through a pointer to a ListApplicationParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **FriendlyName** | **string** | The string that identifies the Application resources to read.
 **PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListApplicationResponse**](ListApplicationResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAuthorizedConnectApp

> ListAuthorizedConnectAppResponse ListAuthorizedConnectApp(ctx, AccountSid).PageSize(PageSize).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AccountSid := "AccountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the AuthorizedConnectApp resources to read.
    PageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListAuthorizedConnectApp(context.Background(), AccountSid).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListAuthorizedConnectApp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAuthorizedConnectApp`: ListAuthorizedConnectAppResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListAuthorizedConnectApp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the AuthorizedConnectApp resources to read.

### Other Parameters

Other parameters are passed through a pointer to a ListAuthorizedConnectAppParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListAuthorizedConnectAppResponse**](ListAuthorizedConnectAppResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAvailablePhoneNumberCountry

> ListAvailablePhoneNumberCountryResponse ListAvailablePhoneNumberCountry(ctx, AccountSid).PageSize(PageSize).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AccountSid := "AccountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) requesting the available phone number Country resources.
    PageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListAvailablePhoneNumberCountry(context.Background(), AccountSid).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListAvailablePhoneNumberCountry``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAvailablePhoneNumberCountry`: ListAvailablePhoneNumberCountryResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListAvailablePhoneNumberCountry`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) requesting the available phone number Country resources.

### Other Parameters

Other parameters are passed through a pointer to a ListAvailablePhoneNumberCountryParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListAvailablePhoneNumberCountryResponse**](ListAvailablePhoneNumberCountryResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAvailablePhoneNumberLocal

> ListAvailablePhoneNumberLocalResponse ListAvailablePhoneNumberLocal(ctx, AccountSid, CountryCode).AreaCode(AreaCode).Contains(Contains).SmsEnabled(SmsEnabled).MmsEnabled(MmsEnabled).VoiceEnabled(VoiceEnabled).ExcludeAllAddressRequired(ExcludeAllAddressRequired).ExcludeLocalAddressRequired(ExcludeLocalAddressRequired).ExcludeForeignAddressRequired(ExcludeForeignAddressRequired).Beta(Beta).NearNumber(NearNumber).NearLatLong(NearLatLong).Distance(Distance).InPostalCode(InPostalCode).InRegion(InRegion).InRateCenter(InRateCenter).InLata(InLata).InLocality(InLocality).FaxEnabled(FaxEnabled).PageSize(PageSize).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AccountSid := "AccountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) requesting the AvailablePhoneNumber resources.
    CountryCode := "CountryCode_example" // string | The [ISO-3166-1](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country code of the country from which to read phone numbers.
    AreaCode := int32(56) // int32 | The area code of the phone numbers to read. Applies to only phone numbers in the US and Canada. (optional)
    Contains := "Contains_example" // string | The pattern on which to match phone numbers. Valid characters are `*`, `0-9`, `a-z`, and `A-Z`. The `*` character matches any single digit. For examples, see [Example 2](https://www.twilio.com/docs/phone-numbers/api/availablephonenumberlocal-resource?code-sample=code-find-phone-numbers-by-number-pattern) and [Example 3](https://www.twilio.com/docs/phone-numbers/api/availablephonenumberlocal-resource?code-sample=code-find-phone-numbers-by-character-pattern). If specified, this value must have at least two characters. (optional)
    SmsEnabled := true // bool | Whether the phone numbers can receive text messages. Can be: `true` or `false`. (optional)
    MmsEnabled := true // bool | Whether the phone numbers can receive MMS messages. Can be: `true` or `false`. (optional)
    VoiceEnabled := true // bool | Whether the phone numbers can receive calls. Can be: `true` or `false`. (optional)
    ExcludeAllAddressRequired := true // bool | Whether to exclude phone numbers that require an [Address](https://www.twilio.com/docs/usage/api/address). Can be: `true` or `false` and the default is `false`. (optional)
    ExcludeLocalAddressRequired := true // bool | Whether to exclude phone numbers that require a local [Address](https://www.twilio.com/docs/usage/api/address). Can be: `true` or `false` and the default is `false`. (optional)
    ExcludeForeignAddressRequired := true // bool | Whether to exclude phone numbers that require a foreign [Address](https://www.twilio.com/docs/usage/api/address). Can be: `true` or `false` and the default is `false`. (optional)
    Beta := true // bool | Whether to read phone numbers that are new to the Twilio platform. Can be: `true` or `false` and the default is `true`. (optional)
    NearNumber := "NearNumber_example" // string | Given a phone number, find a geographically close number within `distance` miles. Distance defaults to 25 miles. Applies to only phone numbers in the US and Canada. (optional)
    NearLatLong := "NearLatLong_example" // string | Given a latitude/longitude pair `lat,long` find geographically close numbers within `distance` miles. Applies to only phone numbers in the US and Canada. (optional)
    Distance := int32(56) // int32 | The search radius, in miles, for a `near_` query.  Can be up to `500` and the default is `25`. Applies to only phone numbers in the US and Canada. (optional)
    InPostalCode := "InPostalCode_example" // string | Limit results to a particular postal code. Given a phone number, search within the same postal code as that number. Applies to only phone numbers in the US and Canada. (optional)
    InRegion := "InRegion_example" // string | Limit results to a particular region, state, or province. Given a phone number, search within the same region as that number. Applies to only phone numbers in the US and Canada. (optional)
    InRateCenter := "InRateCenter_example" // string | Limit results to a specific rate center, or given a phone number search within the same rate center as that number. Requires `in_lata` to be set as well. Applies to only phone numbers in the US and Canada. (optional)
    InLata := "InLata_example" // string | Limit results to a specific local access and transport area ([LATA](https://en.wikipedia.org/wiki/Local_access_and_transport_area)). Given a phone number, search within the same [LATA](https://en.wikipedia.org/wiki/Local_access_and_transport_area) as that number. Applies to only phone numbers in the US and Canada. (optional)
    InLocality := "InLocality_example" // string | Limit results to a particular locality or city. Given a phone number, search within the same Locality as that number. (optional)
    FaxEnabled := true // bool | Whether the phone numbers can receive faxes. Can be: `true` or `false`. (optional)
    PageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListAvailablePhoneNumberLocal(context.Background(), AccountSid, CountryCode).AreaCode(AreaCode).Contains(Contains).SmsEnabled(SmsEnabled).MmsEnabled(MmsEnabled).VoiceEnabled(VoiceEnabled).ExcludeAllAddressRequired(ExcludeAllAddressRequired).ExcludeLocalAddressRequired(ExcludeLocalAddressRequired).ExcludeForeignAddressRequired(ExcludeForeignAddressRequired).Beta(Beta).NearNumber(NearNumber).NearLatLong(NearLatLong).Distance(Distance).InPostalCode(InPostalCode).InRegion(InRegion).InRateCenter(InRateCenter).InLata(InLata).InLocality(InLocality).FaxEnabled(FaxEnabled).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListAvailablePhoneNumberLocal``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAvailablePhoneNumberLocal`: ListAvailablePhoneNumberLocalResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListAvailablePhoneNumberLocal`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) requesting the AvailablePhoneNumber resources.
**CountryCode** | **string** | The [ISO-3166-1](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country code of the country from which to read phone numbers.

### Other Parameters

Other parameters are passed through a pointer to a ListAvailablePhoneNumberLocalParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


 **AreaCode** | **int32** | The area code of the phone numbers to read. Applies to only phone numbers in the US and Canada.
 **Contains** | **string** | The pattern on which to match phone numbers. Valid characters are &#x60;*&#x60;, &#x60;0-9&#x60;, &#x60;a-z&#x60;, and &#x60;A-Z&#x60;. The &#x60;*&#x60; character matches any single digit. For examples, see [Example 2](https://www.twilio.com/docs/phone-numbers/api/availablephonenumberlocal-resource?code-sample&#x3D;code-find-phone-numbers-by-number-pattern) and [Example 3](https://www.twilio.com/docs/phone-numbers/api/availablephonenumberlocal-resource?code-sample&#x3D;code-find-phone-numbers-by-character-pattern). If specified, this value must have at least two characters.
 **SmsEnabled** | **bool** | Whether the phone numbers can receive text messages. Can be: &#x60;true&#x60; or &#x60;false&#x60;.
 **MmsEnabled** | **bool** | Whether the phone numbers can receive MMS messages. Can be: &#x60;true&#x60; or &#x60;false&#x60;.
 **VoiceEnabled** | **bool** | Whether the phone numbers can receive calls. Can be: &#x60;true&#x60; or &#x60;false&#x60;.
 **ExcludeAllAddressRequired** | **bool** | Whether to exclude phone numbers that require an [Address](https://www.twilio.com/docs/usage/api/address). Can be: &#x60;true&#x60; or &#x60;false&#x60; and the default is &#x60;false&#x60;.
 **ExcludeLocalAddressRequired** | **bool** | Whether to exclude phone numbers that require a local [Address](https://www.twilio.com/docs/usage/api/address). Can be: &#x60;true&#x60; or &#x60;false&#x60; and the default is &#x60;false&#x60;.
 **ExcludeForeignAddressRequired** | **bool** | Whether to exclude phone numbers that require a foreign [Address](https://www.twilio.com/docs/usage/api/address). Can be: &#x60;true&#x60; or &#x60;false&#x60; and the default is &#x60;false&#x60;.
 **Beta** | **bool** | Whether to read phone numbers that are new to the Twilio platform. Can be: &#x60;true&#x60; or &#x60;false&#x60; and the default is &#x60;true&#x60;.
 **NearNumber** | **string** | Given a phone number, find a geographically close number within &#x60;distance&#x60; miles. Distance defaults to 25 miles. Applies to only phone numbers in the US and Canada.
 **NearLatLong** | **string** | Given a latitude/longitude pair &#x60;lat,long&#x60; find geographically close numbers within &#x60;distance&#x60; miles. Applies to only phone numbers in the US and Canada.
 **Distance** | **int32** | The search radius, in miles, for a &#x60;near_&#x60; query.  Can be up to &#x60;500&#x60; and the default is &#x60;25&#x60;. Applies to only phone numbers in the US and Canada.
 **InPostalCode** | **string** | Limit results to a particular postal code. Given a phone number, search within the same postal code as that number. Applies to only phone numbers in the US and Canada.
 **InRegion** | **string** | Limit results to a particular region, state, or province. Given a phone number, search within the same region as that number. Applies to only phone numbers in the US and Canada.
 **InRateCenter** | **string** | Limit results to a specific rate center, or given a phone number search within the same rate center as that number. Requires &#x60;in_lata&#x60; to be set as well. Applies to only phone numbers in the US and Canada.
 **InLata** | **string** | Limit results to a specific local access and transport area ([LATA](https://en.wikipedia.org/wiki/Local_access_and_transport_area)). Given a phone number, search within the same [LATA](https://en.wikipedia.org/wiki/Local_access_and_transport_area) as that number. Applies to only phone numbers in the US and Canada.
 **InLocality** | **string** | Limit results to a particular locality or city. Given a phone number, search within the same Locality as that number.
 **FaxEnabled** | **bool** | Whether the phone numbers can receive faxes. Can be: &#x60;true&#x60; or &#x60;false&#x60;.
 **PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListAvailablePhoneNumberLocalResponse**](ListAvailablePhoneNumberLocalResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAvailablePhoneNumberMachineToMachine

> ListAvailablePhoneNumberMachineToMachineResponse ListAvailablePhoneNumberMachineToMachine(ctx, AccountSid, CountryCode).AreaCode(AreaCode).Contains(Contains).SmsEnabled(SmsEnabled).MmsEnabled(MmsEnabled).VoiceEnabled(VoiceEnabled).ExcludeAllAddressRequired(ExcludeAllAddressRequired).ExcludeLocalAddressRequired(ExcludeLocalAddressRequired).ExcludeForeignAddressRequired(ExcludeForeignAddressRequired).Beta(Beta).NearNumber(NearNumber).NearLatLong(NearLatLong).Distance(Distance).InPostalCode(InPostalCode).InRegion(InRegion).InRateCenter(InRateCenter).InLata(InLata).InLocality(InLocality).FaxEnabled(FaxEnabled).PageSize(PageSize).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AccountSid := "AccountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) requesting the AvailablePhoneNumber resources.
    CountryCode := "CountryCode_example" // string | The [ISO-3166-1](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country code of the country from which to read phone numbers.
    AreaCode := int32(56) // int32 | The area code of the phone numbers to read. Applies to only phone numbers in the US and Canada. (optional)
    Contains := "Contains_example" // string | The pattern on which to match phone numbers. Valid characters are `*`, `0-9`, `a-z`, and `A-Z`. The `*` character matches any single digit. For examples, see [Example 2](https://www.twilio.com/docs/phone-numbers/api/availablephonenumber-resource#local-get-basic-example-2) and [Example 3](https://www.twilio.com/docs/phone-numbers/api/availablephonenumber-resource#local-get-basic-example-3). If specified, this value must have at least two characters. (optional)
    SmsEnabled := true // bool | Whether the phone numbers can receive text messages. Can be: `true` or `false`. (optional)
    MmsEnabled := true // bool | Whether the phone numbers can receive MMS messages. Can be: `true` or `false`. (optional)
    VoiceEnabled := true // bool | Whether the phone numbers can receive calls. Can be: `true` or `false`. (optional)
    ExcludeAllAddressRequired := true // bool | Whether to exclude phone numbers that require an [Address](https://www.twilio.com/docs/usage/api/address). Can be: `true` or `false` and the default is `false`. (optional)
    ExcludeLocalAddressRequired := true // bool | Whether to exclude phone numbers that require a local [Address](https://www.twilio.com/docs/usage/api/address). Can be: `true` or `false` and the default is `false`. (optional)
    ExcludeForeignAddressRequired := true // bool | Whether to exclude phone numbers that require a foreign [Address](https://www.twilio.com/docs/usage/api/address). Can be: `true` or `false` and the default is `false`. (optional)
    Beta := true // bool | Whether to read phone numbers that are new to the Twilio platform. Can be: `true` or `false` and the default is `true`. (optional)
    NearNumber := "NearNumber_example" // string | Given a phone number, find a geographically close number within `distance` miles. Distance defaults to 25 miles. Applies to only phone numbers in the US and Canada. (optional)
    NearLatLong := "NearLatLong_example" // string | Given a latitude/longitude pair `lat,long` find geographically close numbers within `distance` miles. Applies to only phone numbers in the US and Canada. (optional)
    Distance := int32(56) // int32 | The search radius, in miles, for a `near_` query.  Can be up to `500` and the default is `25`. Applies to only phone numbers in the US and Canada. (optional)
    InPostalCode := "InPostalCode_example" // string | Limit results to a particular postal code. Given a phone number, search within the same postal code as that number. Applies to only phone numbers in the US and Canada. (optional)
    InRegion := "InRegion_example" // string | Limit results to a particular region, state, or province. Given a phone number, search within the same region as that number. Applies to only phone numbers in the US and Canada. (optional)
    InRateCenter := "InRateCenter_example" // string | Limit results to a specific rate center, or given a phone number search within the same rate center as that number. Requires `in_lata` to be set as well. Applies to only phone numbers in the US and Canada. (optional)
    InLata := "InLata_example" // string | Limit results to a specific local access and transport area ([LATA](https://en.wikipedia.org/wiki/Local_access_and_transport_area)). Given a phone number, search within the same [LATA](https://en.wikipedia.org/wiki/Local_access_and_transport_area) as that number. Applies to only phone numbers in the US and Canada. (optional)
    InLocality := "InLocality_example" // string | Limit results to a particular locality or city. Given a phone number, search within the same Locality as that number. (optional)
    FaxEnabled := true // bool | Whether the phone numbers can receive faxes. Can be: `true` or `false`. (optional)
    PageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListAvailablePhoneNumberMachineToMachine(context.Background(), AccountSid, CountryCode).AreaCode(AreaCode).Contains(Contains).SmsEnabled(SmsEnabled).MmsEnabled(MmsEnabled).VoiceEnabled(VoiceEnabled).ExcludeAllAddressRequired(ExcludeAllAddressRequired).ExcludeLocalAddressRequired(ExcludeLocalAddressRequired).ExcludeForeignAddressRequired(ExcludeForeignAddressRequired).Beta(Beta).NearNumber(NearNumber).NearLatLong(NearLatLong).Distance(Distance).InPostalCode(InPostalCode).InRegion(InRegion).InRateCenter(InRateCenter).InLata(InLata).InLocality(InLocality).FaxEnabled(FaxEnabled).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListAvailablePhoneNumberMachineToMachine``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAvailablePhoneNumberMachineToMachine`: ListAvailablePhoneNumberMachineToMachineResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListAvailablePhoneNumberMachineToMachine`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) requesting the AvailablePhoneNumber resources.
**CountryCode** | **string** | The [ISO-3166-1](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country code of the country from which to read phone numbers.

### Other Parameters

Other parameters are passed through a pointer to a ListAvailablePhoneNumberMachineToMachineParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


 **AreaCode** | **int32** | The area code of the phone numbers to read. Applies to only phone numbers in the US and Canada.
 **Contains** | **string** | The pattern on which to match phone numbers. Valid characters are &#x60;*&#x60;, &#x60;0-9&#x60;, &#x60;a-z&#x60;, and &#x60;A-Z&#x60;. The &#x60;*&#x60; character matches any single digit. For examples, see [Example 2](https://www.twilio.com/docs/phone-numbers/api/availablephonenumber-resource#local-get-basic-example-2) and [Example 3](https://www.twilio.com/docs/phone-numbers/api/availablephonenumber-resource#local-get-basic-example-3). If specified, this value must have at least two characters.
 **SmsEnabled** | **bool** | Whether the phone numbers can receive text messages. Can be: &#x60;true&#x60; or &#x60;false&#x60;.
 **MmsEnabled** | **bool** | Whether the phone numbers can receive MMS messages. Can be: &#x60;true&#x60; or &#x60;false&#x60;.
 **VoiceEnabled** | **bool** | Whether the phone numbers can receive calls. Can be: &#x60;true&#x60; or &#x60;false&#x60;.
 **ExcludeAllAddressRequired** | **bool** | Whether to exclude phone numbers that require an [Address](https://www.twilio.com/docs/usage/api/address). Can be: &#x60;true&#x60; or &#x60;false&#x60; and the default is &#x60;false&#x60;.
 **ExcludeLocalAddressRequired** | **bool** | Whether to exclude phone numbers that require a local [Address](https://www.twilio.com/docs/usage/api/address). Can be: &#x60;true&#x60; or &#x60;false&#x60; and the default is &#x60;false&#x60;.
 **ExcludeForeignAddressRequired** | **bool** | Whether to exclude phone numbers that require a foreign [Address](https://www.twilio.com/docs/usage/api/address). Can be: &#x60;true&#x60; or &#x60;false&#x60; and the default is &#x60;false&#x60;.
 **Beta** | **bool** | Whether to read phone numbers that are new to the Twilio platform. Can be: &#x60;true&#x60; or &#x60;false&#x60; and the default is &#x60;true&#x60;.
 **NearNumber** | **string** | Given a phone number, find a geographically close number within &#x60;distance&#x60; miles. Distance defaults to 25 miles. Applies to only phone numbers in the US and Canada.
 **NearLatLong** | **string** | Given a latitude/longitude pair &#x60;lat,long&#x60; find geographically close numbers within &#x60;distance&#x60; miles. Applies to only phone numbers in the US and Canada.
 **Distance** | **int32** | The search radius, in miles, for a &#x60;near_&#x60; query.  Can be up to &#x60;500&#x60; and the default is &#x60;25&#x60;. Applies to only phone numbers in the US and Canada.
 **InPostalCode** | **string** | Limit results to a particular postal code. Given a phone number, search within the same postal code as that number. Applies to only phone numbers in the US and Canada.
 **InRegion** | **string** | Limit results to a particular region, state, or province. Given a phone number, search within the same region as that number. Applies to only phone numbers in the US and Canada.
 **InRateCenter** | **string** | Limit results to a specific rate center, or given a phone number search within the same rate center as that number. Requires &#x60;in_lata&#x60; to be set as well. Applies to only phone numbers in the US and Canada.
 **InLata** | **string** | Limit results to a specific local access and transport area ([LATA](https://en.wikipedia.org/wiki/Local_access_and_transport_area)). Given a phone number, search within the same [LATA](https://en.wikipedia.org/wiki/Local_access_and_transport_area) as that number. Applies to only phone numbers in the US and Canada.
 **InLocality** | **string** | Limit results to a particular locality or city. Given a phone number, search within the same Locality as that number.
 **FaxEnabled** | **bool** | Whether the phone numbers can receive faxes. Can be: &#x60;true&#x60; or &#x60;false&#x60;.
 **PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListAvailablePhoneNumberMachineToMachineResponse**](ListAvailablePhoneNumberMachineToMachineResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAvailablePhoneNumberMobile

> ListAvailablePhoneNumberMobileResponse ListAvailablePhoneNumberMobile(ctx, AccountSid, CountryCode).AreaCode(AreaCode).Contains(Contains).SmsEnabled(SmsEnabled).MmsEnabled(MmsEnabled).VoiceEnabled(VoiceEnabled).ExcludeAllAddressRequired(ExcludeAllAddressRequired).ExcludeLocalAddressRequired(ExcludeLocalAddressRequired).ExcludeForeignAddressRequired(ExcludeForeignAddressRequired).Beta(Beta).NearNumber(NearNumber).NearLatLong(NearLatLong).Distance(Distance).InPostalCode(InPostalCode).InRegion(InRegion).InRateCenter(InRateCenter).InLata(InLata).InLocality(InLocality).FaxEnabled(FaxEnabled).PageSize(PageSize).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AccountSid := "AccountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) requesting the AvailablePhoneNumber resources.
    CountryCode := "CountryCode_example" // string | The [ISO-3166-1](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country code of the country from which to read phone numbers.
    AreaCode := int32(56) // int32 | The area code of the phone numbers to read. Applies to only phone numbers in the US and Canada. (optional)
    Contains := "Contains_example" // string | The pattern on which to match phone numbers. Valid characters are `*`, `0-9`, `a-z`, and `A-Z`. The `*` character matches any single digit. For examples, see [Example 2](https://www.twilio.com/docs/phone-numbers/api/availablephonenumber-resource#local-get-basic-example-2) and [Example 3](https://www.twilio.com/docs/phone-numbers/api/availablephonenumber-resource#local-get-basic-example-3). If specified, this value must have at least two characters. (optional)
    SmsEnabled := true // bool | Whether the phone numbers can receive text messages. Can be: `true` or `false`. (optional)
    MmsEnabled := true // bool | Whether the phone numbers can receive MMS messages. Can be: `true` or `false`. (optional)
    VoiceEnabled := true // bool | Whether the phone numbers can receive calls. Can be: `true` or `false`. (optional)
    ExcludeAllAddressRequired := true // bool | Whether to exclude phone numbers that require an [Address](https://www.twilio.com/docs/usage/api/address). Can be: `true` or `false` and the default is `false`. (optional)
    ExcludeLocalAddressRequired := true // bool | Whether to exclude phone numbers that require a local [Address](https://www.twilio.com/docs/usage/api/address). Can be: `true` or `false` and the default is `false`. (optional)
    ExcludeForeignAddressRequired := true // bool | Whether to exclude phone numbers that require a foreign [Address](https://www.twilio.com/docs/usage/api/address). Can be: `true` or `false` and the default is `false`. (optional)
    Beta := true // bool | Whether to read phone numbers that are new to the Twilio platform. Can be: `true` or `false` and the default is `true`. (optional)
    NearNumber := "NearNumber_example" // string | Given a phone number, find a geographically close number within `distance` miles. Distance defaults to 25 miles. Applies to only phone numbers in the US and Canada. (optional)
    NearLatLong := "NearLatLong_example" // string | Given a latitude/longitude pair `lat,long` find geographically close numbers within `distance` miles. Applies to only phone numbers in the US and Canada. (optional)
    Distance := int32(56) // int32 | The search radius, in miles, for a `near_` query.  Can be up to `500` and the default is `25`. Applies to only phone numbers in the US and Canada. (optional)
    InPostalCode := "InPostalCode_example" // string | Limit results to a particular postal code. Given a phone number, search within the same postal code as that number. Applies to only phone numbers in the US and Canada. (optional)
    InRegion := "InRegion_example" // string | Limit results to a particular region, state, or province. Given a phone number, search within the same region as that number. Applies to only phone numbers in the US and Canada. (optional)
    InRateCenter := "InRateCenter_example" // string | Limit results to a specific rate center, or given a phone number search within the same rate center as that number. Requires `in_lata` to be set as well. Applies to only phone numbers in the US and Canada. (optional)
    InLata := "InLata_example" // string | Limit results to a specific local access and transport area ([LATA](https://en.wikipedia.org/wiki/Local_access_and_transport_area)). Given a phone number, search within the same [LATA](https://en.wikipedia.org/wiki/Local_access_and_transport_area) as that number. Applies to only phone numbers in the US and Canada. (optional)
    InLocality := "InLocality_example" // string | Limit results to a particular locality or city. Given a phone number, search within the same Locality as that number. (optional)
    FaxEnabled := true // bool | Whether the phone numbers can receive faxes. Can be: `true` or `false`. (optional)
    PageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListAvailablePhoneNumberMobile(context.Background(), AccountSid, CountryCode).AreaCode(AreaCode).Contains(Contains).SmsEnabled(SmsEnabled).MmsEnabled(MmsEnabled).VoiceEnabled(VoiceEnabled).ExcludeAllAddressRequired(ExcludeAllAddressRequired).ExcludeLocalAddressRequired(ExcludeLocalAddressRequired).ExcludeForeignAddressRequired(ExcludeForeignAddressRequired).Beta(Beta).NearNumber(NearNumber).NearLatLong(NearLatLong).Distance(Distance).InPostalCode(InPostalCode).InRegion(InRegion).InRateCenter(InRateCenter).InLata(InLata).InLocality(InLocality).FaxEnabled(FaxEnabled).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListAvailablePhoneNumberMobile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAvailablePhoneNumberMobile`: ListAvailablePhoneNumberMobileResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListAvailablePhoneNumberMobile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) requesting the AvailablePhoneNumber resources.
**CountryCode** | **string** | The [ISO-3166-1](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country code of the country from which to read phone numbers.

### Other Parameters

Other parameters are passed through a pointer to a ListAvailablePhoneNumberMobileParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


 **AreaCode** | **int32** | The area code of the phone numbers to read. Applies to only phone numbers in the US and Canada.
 **Contains** | **string** | The pattern on which to match phone numbers. Valid characters are &#x60;*&#x60;, &#x60;0-9&#x60;, &#x60;a-z&#x60;, and &#x60;A-Z&#x60;. The &#x60;*&#x60; character matches any single digit. For examples, see [Example 2](https://www.twilio.com/docs/phone-numbers/api/availablephonenumber-resource#local-get-basic-example-2) and [Example 3](https://www.twilio.com/docs/phone-numbers/api/availablephonenumber-resource#local-get-basic-example-3). If specified, this value must have at least two characters.
 **SmsEnabled** | **bool** | Whether the phone numbers can receive text messages. Can be: &#x60;true&#x60; or &#x60;false&#x60;.
 **MmsEnabled** | **bool** | Whether the phone numbers can receive MMS messages. Can be: &#x60;true&#x60; or &#x60;false&#x60;.
 **VoiceEnabled** | **bool** | Whether the phone numbers can receive calls. Can be: &#x60;true&#x60; or &#x60;false&#x60;.
 **ExcludeAllAddressRequired** | **bool** | Whether to exclude phone numbers that require an [Address](https://www.twilio.com/docs/usage/api/address). Can be: &#x60;true&#x60; or &#x60;false&#x60; and the default is &#x60;false&#x60;.
 **ExcludeLocalAddressRequired** | **bool** | Whether to exclude phone numbers that require a local [Address](https://www.twilio.com/docs/usage/api/address). Can be: &#x60;true&#x60; or &#x60;false&#x60; and the default is &#x60;false&#x60;.
 **ExcludeForeignAddressRequired** | **bool** | Whether to exclude phone numbers that require a foreign [Address](https://www.twilio.com/docs/usage/api/address). Can be: &#x60;true&#x60; or &#x60;false&#x60; and the default is &#x60;false&#x60;.
 **Beta** | **bool** | Whether to read phone numbers that are new to the Twilio platform. Can be: &#x60;true&#x60; or &#x60;false&#x60; and the default is &#x60;true&#x60;.
 **NearNumber** | **string** | Given a phone number, find a geographically close number within &#x60;distance&#x60; miles. Distance defaults to 25 miles. Applies to only phone numbers in the US and Canada.
 **NearLatLong** | **string** | Given a latitude/longitude pair &#x60;lat,long&#x60; find geographically close numbers within &#x60;distance&#x60; miles. Applies to only phone numbers in the US and Canada.
 **Distance** | **int32** | The search radius, in miles, for a &#x60;near_&#x60; query.  Can be up to &#x60;500&#x60; and the default is &#x60;25&#x60;. Applies to only phone numbers in the US and Canada.
 **InPostalCode** | **string** | Limit results to a particular postal code. Given a phone number, search within the same postal code as that number. Applies to only phone numbers in the US and Canada.
 **InRegion** | **string** | Limit results to a particular region, state, or province. Given a phone number, search within the same region as that number. Applies to only phone numbers in the US and Canada.
 **InRateCenter** | **string** | Limit results to a specific rate center, or given a phone number search within the same rate center as that number. Requires &#x60;in_lata&#x60; to be set as well. Applies to only phone numbers in the US and Canada.
 **InLata** | **string** | Limit results to a specific local access and transport area ([LATA](https://en.wikipedia.org/wiki/Local_access_and_transport_area)). Given a phone number, search within the same [LATA](https://en.wikipedia.org/wiki/Local_access_and_transport_area) as that number. Applies to only phone numbers in the US and Canada.
 **InLocality** | **string** | Limit results to a particular locality or city. Given a phone number, search within the same Locality as that number.
 **FaxEnabled** | **bool** | Whether the phone numbers can receive faxes. Can be: &#x60;true&#x60; or &#x60;false&#x60;.
 **PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListAvailablePhoneNumberMobileResponse**](ListAvailablePhoneNumberMobileResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAvailablePhoneNumberNational

> ListAvailablePhoneNumberNationalResponse ListAvailablePhoneNumberNational(ctx, AccountSid, CountryCode).AreaCode(AreaCode).Contains(Contains).SmsEnabled(SmsEnabled).MmsEnabled(MmsEnabled).VoiceEnabled(VoiceEnabled).ExcludeAllAddressRequired(ExcludeAllAddressRequired).ExcludeLocalAddressRequired(ExcludeLocalAddressRequired).ExcludeForeignAddressRequired(ExcludeForeignAddressRequired).Beta(Beta).NearNumber(NearNumber).NearLatLong(NearLatLong).Distance(Distance).InPostalCode(InPostalCode).InRegion(InRegion).InRateCenter(InRateCenter).InLata(InLata).InLocality(InLocality).FaxEnabled(FaxEnabled).PageSize(PageSize).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AccountSid := "AccountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) requesting the AvailablePhoneNumber resources.
    CountryCode := "CountryCode_example" // string | The [ISO-3166-1](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country code of the country from which to read phone numbers.
    AreaCode := int32(56) // int32 | The area code of the phone numbers to read. Applies to only phone numbers in the US and Canada. (optional)
    Contains := "Contains_example" // string | The pattern on which to match phone numbers. Valid characters are `*`, `0-9`, `a-z`, and `A-Z`. The `*` character matches any single digit. For examples, see [Example 2](https://www.twilio.com/docs/phone-numbers/api/availablephonenumber-resource#local-get-basic-example-2) and [Example 3](https://www.twilio.com/docs/phone-numbers/api/availablephonenumber-resource#local-get-basic-example-3). If specified, this value must have at least two characters. (optional)
    SmsEnabled := true // bool | Whether the phone numbers can receive text messages. Can be: `true` or `false`. (optional)
    MmsEnabled := true // bool | Whether the phone numbers can receive MMS messages. Can be: `true` or `false`. (optional)
    VoiceEnabled := true // bool | Whether the phone numbers can receive calls. Can be: `true` or `false`. (optional)
    ExcludeAllAddressRequired := true // bool | Whether to exclude phone numbers that require an [Address](https://www.twilio.com/docs/usage/api/address). Can be: `true` or `false` and the default is `false`. (optional)
    ExcludeLocalAddressRequired := true // bool | Whether to exclude phone numbers that require a local [Address](https://www.twilio.com/docs/usage/api/address). Can be: `true` or `false` and the default is `false`. (optional)
    ExcludeForeignAddressRequired := true // bool | Whether to exclude phone numbers that require a foreign [Address](https://www.twilio.com/docs/usage/api/address). Can be: `true` or `false` and the default is `false`. (optional)
    Beta := true // bool | Whether to read phone numbers that are new to the Twilio platform. Can be: `true` or `false` and the default is `true`. (optional)
    NearNumber := "NearNumber_example" // string | Given a phone number, find a geographically close number within `distance` miles. Distance defaults to 25 miles. Applies to only phone numbers in the US and Canada. (optional)
    NearLatLong := "NearLatLong_example" // string | Given a latitude/longitude pair `lat,long` find geographically close numbers within `distance` miles. Applies to only phone numbers in the US and Canada. (optional)
    Distance := int32(56) // int32 | The search radius, in miles, for a `near_` query.  Can be up to `500` and the default is `25`. Applies to only phone numbers in the US and Canada. (optional)
    InPostalCode := "InPostalCode_example" // string | Limit results to a particular postal code. Given a phone number, search within the same postal code as that number. Applies to only phone numbers in the US and Canada. (optional)
    InRegion := "InRegion_example" // string | Limit results to a particular region, state, or province. Given a phone number, search within the same region as that number. Applies to only phone numbers in the US and Canada. (optional)
    InRateCenter := "InRateCenter_example" // string | Limit results to a specific rate center, or given a phone number search within the same rate center as that number. Requires `in_lata` to be set as well. Applies to only phone numbers in the US and Canada. (optional)
    InLata := "InLata_example" // string | Limit results to a specific local access and transport area ([LATA](https://en.wikipedia.org/wiki/Local_access_and_transport_area)). Given a phone number, search within the same [LATA](https://en.wikipedia.org/wiki/Local_access_and_transport_area) as that number. Applies to only phone numbers in the US and Canada. (optional)
    InLocality := "InLocality_example" // string | Limit results to a particular locality or city. Given a phone number, search within the same Locality as that number. (optional)
    FaxEnabled := true // bool | Whether the phone numbers can receive faxes. Can be: `true` or `false`. (optional)
    PageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListAvailablePhoneNumberNational(context.Background(), AccountSid, CountryCode).AreaCode(AreaCode).Contains(Contains).SmsEnabled(SmsEnabled).MmsEnabled(MmsEnabled).VoiceEnabled(VoiceEnabled).ExcludeAllAddressRequired(ExcludeAllAddressRequired).ExcludeLocalAddressRequired(ExcludeLocalAddressRequired).ExcludeForeignAddressRequired(ExcludeForeignAddressRequired).Beta(Beta).NearNumber(NearNumber).NearLatLong(NearLatLong).Distance(Distance).InPostalCode(InPostalCode).InRegion(InRegion).InRateCenter(InRateCenter).InLata(InLata).InLocality(InLocality).FaxEnabled(FaxEnabled).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListAvailablePhoneNumberNational``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAvailablePhoneNumberNational`: ListAvailablePhoneNumberNationalResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListAvailablePhoneNumberNational`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) requesting the AvailablePhoneNumber resources.
**CountryCode** | **string** | The [ISO-3166-1](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country code of the country from which to read phone numbers.

### Other Parameters

Other parameters are passed through a pointer to a ListAvailablePhoneNumberNationalParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


 **AreaCode** | **int32** | The area code of the phone numbers to read. Applies to only phone numbers in the US and Canada.
 **Contains** | **string** | The pattern on which to match phone numbers. Valid characters are &#x60;*&#x60;, &#x60;0-9&#x60;, &#x60;a-z&#x60;, and &#x60;A-Z&#x60;. The &#x60;*&#x60; character matches any single digit. For examples, see [Example 2](https://www.twilio.com/docs/phone-numbers/api/availablephonenumber-resource#local-get-basic-example-2) and [Example 3](https://www.twilio.com/docs/phone-numbers/api/availablephonenumber-resource#local-get-basic-example-3). If specified, this value must have at least two characters.
 **SmsEnabled** | **bool** | Whether the phone numbers can receive text messages. Can be: &#x60;true&#x60; or &#x60;false&#x60;.
 **MmsEnabled** | **bool** | Whether the phone numbers can receive MMS messages. Can be: &#x60;true&#x60; or &#x60;false&#x60;.
 **VoiceEnabled** | **bool** | Whether the phone numbers can receive calls. Can be: &#x60;true&#x60; or &#x60;false&#x60;.
 **ExcludeAllAddressRequired** | **bool** | Whether to exclude phone numbers that require an [Address](https://www.twilio.com/docs/usage/api/address). Can be: &#x60;true&#x60; or &#x60;false&#x60; and the default is &#x60;false&#x60;.
 **ExcludeLocalAddressRequired** | **bool** | Whether to exclude phone numbers that require a local [Address](https://www.twilio.com/docs/usage/api/address). Can be: &#x60;true&#x60; or &#x60;false&#x60; and the default is &#x60;false&#x60;.
 **ExcludeForeignAddressRequired** | **bool** | Whether to exclude phone numbers that require a foreign [Address](https://www.twilio.com/docs/usage/api/address). Can be: &#x60;true&#x60; or &#x60;false&#x60; and the default is &#x60;false&#x60;.
 **Beta** | **bool** | Whether to read phone numbers that are new to the Twilio platform. Can be: &#x60;true&#x60; or &#x60;false&#x60; and the default is &#x60;true&#x60;.
 **NearNumber** | **string** | Given a phone number, find a geographically close number within &#x60;distance&#x60; miles. Distance defaults to 25 miles. Applies to only phone numbers in the US and Canada.
 **NearLatLong** | **string** | Given a latitude/longitude pair &#x60;lat,long&#x60; find geographically close numbers within &#x60;distance&#x60; miles. Applies to only phone numbers in the US and Canada.
 **Distance** | **int32** | The search radius, in miles, for a &#x60;near_&#x60; query.  Can be up to &#x60;500&#x60; and the default is &#x60;25&#x60;. Applies to only phone numbers in the US and Canada.
 **InPostalCode** | **string** | Limit results to a particular postal code. Given a phone number, search within the same postal code as that number. Applies to only phone numbers in the US and Canada.
 **InRegion** | **string** | Limit results to a particular region, state, or province. Given a phone number, search within the same region as that number. Applies to only phone numbers in the US and Canada.
 **InRateCenter** | **string** | Limit results to a specific rate center, or given a phone number search within the same rate center as that number. Requires &#x60;in_lata&#x60; to be set as well. Applies to only phone numbers in the US and Canada.
 **InLata** | **string** | Limit results to a specific local access and transport area ([LATA](https://en.wikipedia.org/wiki/Local_access_and_transport_area)). Given a phone number, search within the same [LATA](https://en.wikipedia.org/wiki/Local_access_and_transport_area) as that number. Applies to only phone numbers in the US and Canada.
 **InLocality** | **string** | Limit results to a particular locality or city. Given a phone number, search within the same Locality as that number.
 **FaxEnabled** | **bool** | Whether the phone numbers can receive faxes. Can be: &#x60;true&#x60; or &#x60;false&#x60;.
 **PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListAvailablePhoneNumberNationalResponse**](ListAvailablePhoneNumberNationalResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAvailablePhoneNumberSharedCost

> ListAvailablePhoneNumberSharedCostResponse ListAvailablePhoneNumberSharedCost(ctx, AccountSid, CountryCode).AreaCode(AreaCode).Contains(Contains).SmsEnabled(SmsEnabled).MmsEnabled(MmsEnabled).VoiceEnabled(VoiceEnabled).ExcludeAllAddressRequired(ExcludeAllAddressRequired).ExcludeLocalAddressRequired(ExcludeLocalAddressRequired).ExcludeForeignAddressRequired(ExcludeForeignAddressRequired).Beta(Beta).NearNumber(NearNumber).NearLatLong(NearLatLong).Distance(Distance).InPostalCode(InPostalCode).InRegion(InRegion).InRateCenter(InRateCenter).InLata(InLata).InLocality(InLocality).FaxEnabled(FaxEnabled).PageSize(PageSize).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AccountSid := "AccountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) requesting the AvailablePhoneNumber resources.
    CountryCode := "CountryCode_example" // string | The [ISO-3166-1](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country code of the country from which to read phone numbers.
    AreaCode := int32(56) // int32 | The area code of the phone numbers to read. Applies to only phone numbers in the US and Canada. (optional)
    Contains := "Contains_example" // string | The pattern on which to match phone numbers. Valid characters are `*`, `0-9`, `a-z`, and `A-Z`. The `*` character matches any single digit. For examples, see [Example 2](https://www.twilio.com/docs/phone-numbers/api/availablephonenumber-resource#local-get-basic-example-2) and [Example 3](https://www.twilio.com/docs/phone-numbers/api/availablephonenumber-resource#local-get-basic-example-3). If specified, this value must have at least two characters. (optional)
    SmsEnabled := true // bool | Whether the phone numbers can receive text messages. Can be: `true` or `false`. (optional)
    MmsEnabled := true // bool | Whether the phone numbers can receive MMS messages. Can be: `true` or `false`. (optional)
    VoiceEnabled := true // bool | Whether the phone numbers can receive calls. Can be: `true` or `false`. (optional)
    ExcludeAllAddressRequired := true // bool | Whether to exclude phone numbers that require an [Address](https://www.twilio.com/docs/usage/api/address). Can be: `true` or `false` and the default is `false`. (optional)
    ExcludeLocalAddressRequired := true // bool | Whether to exclude phone numbers that require a local [Address](https://www.twilio.com/docs/usage/api/address). Can be: `true` or `false` and the default is `false`. (optional)
    ExcludeForeignAddressRequired := true // bool | Whether to exclude phone numbers that require a foreign [Address](https://www.twilio.com/docs/usage/api/address). Can be: `true` or `false` and the default is `false`. (optional)
    Beta := true // bool | Whether to read phone numbers that are new to the Twilio platform. Can be: `true` or `false` and the default is `true`. (optional)
    NearNumber := "NearNumber_example" // string | Given a phone number, find a geographically close number within `distance` miles. Distance defaults to 25 miles. Applies to only phone numbers in the US and Canada. (optional)
    NearLatLong := "NearLatLong_example" // string | Given a latitude/longitude pair `lat,long` find geographically close numbers within `distance` miles. Applies to only phone numbers in the US and Canada. (optional)
    Distance := int32(56) // int32 | The search radius, in miles, for a `near_` query.  Can be up to `500` and the default is `25`. Applies to only phone numbers in the US and Canada. (optional)
    InPostalCode := "InPostalCode_example" // string | Limit results to a particular postal code. Given a phone number, search within the same postal code as that number. Applies to only phone numbers in the US and Canada. (optional)
    InRegion := "InRegion_example" // string | Limit results to a particular region, state, or province. Given a phone number, search within the same region as that number. Applies to only phone numbers in the US and Canada. (optional)
    InRateCenter := "InRateCenter_example" // string | Limit results to a specific rate center, or given a phone number search within the same rate center as that number. Requires `in_lata` to be set as well. Applies to only phone numbers in the US and Canada. (optional)
    InLata := "InLata_example" // string | Limit results to a specific local access and transport area ([LATA](https://en.wikipedia.org/wiki/Local_access_and_transport_area)). Given a phone number, search within the same [LATA](https://en.wikipedia.org/wiki/Local_access_and_transport_area) as that number. Applies to only phone numbers in the US and Canada. (optional)
    InLocality := "InLocality_example" // string | Limit results to a particular locality or city. Given a phone number, search within the same Locality as that number. (optional)
    FaxEnabled := true // bool | Whether the phone numbers can receive faxes. Can be: `true` or `false`. (optional)
    PageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListAvailablePhoneNumberSharedCost(context.Background(), AccountSid, CountryCode).AreaCode(AreaCode).Contains(Contains).SmsEnabled(SmsEnabled).MmsEnabled(MmsEnabled).VoiceEnabled(VoiceEnabled).ExcludeAllAddressRequired(ExcludeAllAddressRequired).ExcludeLocalAddressRequired(ExcludeLocalAddressRequired).ExcludeForeignAddressRequired(ExcludeForeignAddressRequired).Beta(Beta).NearNumber(NearNumber).NearLatLong(NearLatLong).Distance(Distance).InPostalCode(InPostalCode).InRegion(InRegion).InRateCenter(InRateCenter).InLata(InLata).InLocality(InLocality).FaxEnabled(FaxEnabled).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListAvailablePhoneNumberSharedCost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAvailablePhoneNumberSharedCost`: ListAvailablePhoneNumberSharedCostResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListAvailablePhoneNumberSharedCost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) requesting the AvailablePhoneNumber resources.
**CountryCode** | **string** | The [ISO-3166-1](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country code of the country from which to read phone numbers.

### Other Parameters

Other parameters are passed through a pointer to a ListAvailablePhoneNumberSharedCostParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


 **AreaCode** | **int32** | The area code of the phone numbers to read. Applies to only phone numbers in the US and Canada.
 **Contains** | **string** | The pattern on which to match phone numbers. Valid characters are &#x60;*&#x60;, &#x60;0-9&#x60;, &#x60;a-z&#x60;, and &#x60;A-Z&#x60;. The &#x60;*&#x60; character matches any single digit. For examples, see [Example 2](https://www.twilio.com/docs/phone-numbers/api/availablephonenumber-resource#local-get-basic-example-2) and [Example 3](https://www.twilio.com/docs/phone-numbers/api/availablephonenumber-resource#local-get-basic-example-3). If specified, this value must have at least two characters.
 **SmsEnabled** | **bool** | Whether the phone numbers can receive text messages. Can be: &#x60;true&#x60; or &#x60;false&#x60;.
 **MmsEnabled** | **bool** | Whether the phone numbers can receive MMS messages. Can be: &#x60;true&#x60; or &#x60;false&#x60;.
 **VoiceEnabled** | **bool** | Whether the phone numbers can receive calls. Can be: &#x60;true&#x60; or &#x60;false&#x60;.
 **ExcludeAllAddressRequired** | **bool** | Whether to exclude phone numbers that require an [Address](https://www.twilio.com/docs/usage/api/address). Can be: &#x60;true&#x60; or &#x60;false&#x60; and the default is &#x60;false&#x60;.
 **ExcludeLocalAddressRequired** | **bool** | Whether to exclude phone numbers that require a local [Address](https://www.twilio.com/docs/usage/api/address). Can be: &#x60;true&#x60; or &#x60;false&#x60; and the default is &#x60;false&#x60;.
 **ExcludeForeignAddressRequired** | **bool** | Whether to exclude phone numbers that require a foreign [Address](https://www.twilio.com/docs/usage/api/address). Can be: &#x60;true&#x60; or &#x60;false&#x60; and the default is &#x60;false&#x60;.
 **Beta** | **bool** | Whether to read phone numbers that are new to the Twilio platform. Can be: &#x60;true&#x60; or &#x60;false&#x60; and the default is &#x60;true&#x60;.
 **NearNumber** | **string** | Given a phone number, find a geographically close number within &#x60;distance&#x60; miles. Distance defaults to 25 miles. Applies to only phone numbers in the US and Canada.
 **NearLatLong** | **string** | Given a latitude/longitude pair &#x60;lat,long&#x60; find geographically close numbers within &#x60;distance&#x60; miles. Applies to only phone numbers in the US and Canada.
 **Distance** | **int32** | The search radius, in miles, for a &#x60;near_&#x60; query.  Can be up to &#x60;500&#x60; and the default is &#x60;25&#x60;. Applies to only phone numbers in the US and Canada.
 **InPostalCode** | **string** | Limit results to a particular postal code. Given a phone number, search within the same postal code as that number. Applies to only phone numbers in the US and Canada.
 **InRegion** | **string** | Limit results to a particular region, state, or province. Given a phone number, search within the same region as that number. Applies to only phone numbers in the US and Canada.
 **InRateCenter** | **string** | Limit results to a specific rate center, or given a phone number search within the same rate center as that number. Requires &#x60;in_lata&#x60; to be set as well. Applies to only phone numbers in the US and Canada.
 **InLata** | **string** | Limit results to a specific local access and transport area ([LATA](https://en.wikipedia.org/wiki/Local_access_and_transport_area)). Given a phone number, search within the same [LATA](https://en.wikipedia.org/wiki/Local_access_and_transport_area) as that number. Applies to only phone numbers in the US and Canada.
 **InLocality** | **string** | Limit results to a particular locality or city. Given a phone number, search within the same Locality as that number.
 **FaxEnabled** | **bool** | Whether the phone numbers can receive faxes. Can be: &#x60;true&#x60; or &#x60;false&#x60;.
 **PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListAvailablePhoneNumberSharedCostResponse**](ListAvailablePhoneNumberSharedCostResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAvailablePhoneNumberTollFree

> ListAvailablePhoneNumberTollFreeResponse ListAvailablePhoneNumberTollFree(ctx, AccountSid, CountryCode).AreaCode(AreaCode).Contains(Contains).SmsEnabled(SmsEnabled).MmsEnabled(MmsEnabled).VoiceEnabled(VoiceEnabled).ExcludeAllAddressRequired(ExcludeAllAddressRequired).ExcludeLocalAddressRequired(ExcludeLocalAddressRequired).ExcludeForeignAddressRequired(ExcludeForeignAddressRequired).Beta(Beta).NearNumber(NearNumber).NearLatLong(NearLatLong).Distance(Distance).InPostalCode(InPostalCode).InRegion(InRegion).InRateCenter(InRateCenter).InLata(InLata).InLocality(InLocality).FaxEnabled(FaxEnabled).PageSize(PageSize).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AccountSid := "AccountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) requesting the AvailablePhoneNumber resources.
    CountryCode := "CountryCode_example" // string | The [ISO-3166-1](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country code of the country from which to read phone numbers.
    AreaCode := int32(56) // int32 | The area code of the phone numbers to read. Applies to only phone numbers in the US and Canada. (optional)
    Contains := "Contains_example" // string | The pattern on which to match phone numbers. Valid characters are `*`, `0-9`, `a-z`, and `A-Z`. The `*` character matches any single digit. For examples, see [Example 2](https://www.twilio.com/docs/phone-numbers/api/availablephonenumber-resource#local-get-basic-example-2) and [Example 3](https://www.twilio.com/docs/phone-numbers/api/availablephonenumber-resource#local-get-basic-example-3). If specified, this value must have at least two characters. (optional)
    SmsEnabled := true // bool | Whether the phone numbers can receive text messages. Can be: `true` or `false`. (optional)
    MmsEnabled := true // bool | Whether the phone numbers can receive MMS messages. Can be: `true` or `false`. (optional)
    VoiceEnabled := true // bool | Whether the phone numbers can receive calls. Can be: `true` or `false`. (optional)
    ExcludeAllAddressRequired := true // bool | Whether to exclude phone numbers that require an [Address](https://www.twilio.com/docs/usage/api/address). Can be: `true` or `false` and the default is `false`. (optional)
    ExcludeLocalAddressRequired := true // bool | Whether to exclude phone numbers that require a local [Address](https://www.twilio.com/docs/usage/api/address). Can be: `true` or `false` and the default is `false`. (optional)
    ExcludeForeignAddressRequired := true // bool | Whether to exclude phone numbers that require a foreign [Address](https://www.twilio.com/docs/usage/api/address). Can be: `true` or `false` and the default is `false`. (optional)
    Beta := true // bool | Whether to read phone numbers that are new to the Twilio platform. Can be: `true` or `false` and the default is `true`. (optional)
    NearNumber := "NearNumber_example" // string | Given a phone number, find a geographically close number within `distance` miles. Distance defaults to 25 miles. Applies to only phone numbers in the US and Canada. (optional)
    NearLatLong := "NearLatLong_example" // string | Given a latitude/longitude pair `lat,long` find geographically close numbers within `distance` miles. Applies to only phone numbers in the US and Canada. (optional)
    Distance := int32(56) // int32 | The search radius, in miles, for a `near_` query.  Can be up to `500` and the default is `25`. Applies to only phone numbers in the US and Canada. (optional)
    InPostalCode := "InPostalCode_example" // string | Limit results to a particular postal code. Given a phone number, search within the same postal code as that number. Applies to only phone numbers in the US and Canada. (optional)
    InRegion := "InRegion_example" // string | Limit results to a particular region, state, or province. Given a phone number, search within the same region as that number. Applies to only phone numbers in the US and Canada. (optional)
    InRateCenter := "InRateCenter_example" // string | Limit results to a specific rate center, or given a phone number search within the same rate center as that number. Requires `in_lata` to be set as well. Applies to only phone numbers in the US and Canada. (optional)
    InLata := "InLata_example" // string | Limit results to a specific local access and transport area ([LATA](https://en.wikipedia.org/wiki/Local_access_and_transport_area)). Given a phone number, search within the same [LATA](https://en.wikipedia.org/wiki/Local_access_and_transport_area) as that number. Applies to only phone numbers in the US and Canada. (optional)
    InLocality := "InLocality_example" // string | Limit results to a particular locality or city. Given a phone number, search within the same Locality as that number. (optional)
    FaxEnabled := true // bool | Whether the phone numbers can receive faxes. Can be: `true` or `false`. (optional)
    PageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListAvailablePhoneNumberTollFree(context.Background(), AccountSid, CountryCode).AreaCode(AreaCode).Contains(Contains).SmsEnabled(SmsEnabled).MmsEnabled(MmsEnabled).VoiceEnabled(VoiceEnabled).ExcludeAllAddressRequired(ExcludeAllAddressRequired).ExcludeLocalAddressRequired(ExcludeLocalAddressRequired).ExcludeForeignAddressRequired(ExcludeForeignAddressRequired).Beta(Beta).NearNumber(NearNumber).NearLatLong(NearLatLong).Distance(Distance).InPostalCode(InPostalCode).InRegion(InRegion).InRateCenter(InRateCenter).InLata(InLata).InLocality(InLocality).FaxEnabled(FaxEnabled).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListAvailablePhoneNumberTollFree``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAvailablePhoneNumberTollFree`: ListAvailablePhoneNumberTollFreeResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListAvailablePhoneNumberTollFree`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) requesting the AvailablePhoneNumber resources.
**CountryCode** | **string** | The [ISO-3166-1](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country code of the country from which to read phone numbers.

### Other Parameters

Other parameters are passed through a pointer to a ListAvailablePhoneNumberTollFreeParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


 **AreaCode** | **int32** | The area code of the phone numbers to read. Applies to only phone numbers in the US and Canada.
 **Contains** | **string** | The pattern on which to match phone numbers. Valid characters are &#x60;*&#x60;, &#x60;0-9&#x60;, &#x60;a-z&#x60;, and &#x60;A-Z&#x60;. The &#x60;*&#x60; character matches any single digit. For examples, see [Example 2](https://www.twilio.com/docs/phone-numbers/api/availablephonenumber-resource#local-get-basic-example-2) and [Example 3](https://www.twilio.com/docs/phone-numbers/api/availablephonenumber-resource#local-get-basic-example-3). If specified, this value must have at least two characters.
 **SmsEnabled** | **bool** | Whether the phone numbers can receive text messages. Can be: &#x60;true&#x60; or &#x60;false&#x60;.
 **MmsEnabled** | **bool** | Whether the phone numbers can receive MMS messages. Can be: &#x60;true&#x60; or &#x60;false&#x60;.
 **VoiceEnabled** | **bool** | Whether the phone numbers can receive calls. Can be: &#x60;true&#x60; or &#x60;false&#x60;.
 **ExcludeAllAddressRequired** | **bool** | Whether to exclude phone numbers that require an [Address](https://www.twilio.com/docs/usage/api/address). Can be: &#x60;true&#x60; or &#x60;false&#x60; and the default is &#x60;false&#x60;.
 **ExcludeLocalAddressRequired** | **bool** | Whether to exclude phone numbers that require a local [Address](https://www.twilio.com/docs/usage/api/address). Can be: &#x60;true&#x60; or &#x60;false&#x60; and the default is &#x60;false&#x60;.
 **ExcludeForeignAddressRequired** | **bool** | Whether to exclude phone numbers that require a foreign [Address](https://www.twilio.com/docs/usage/api/address). Can be: &#x60;true&#x60; or &#x60;false&#x60; and the default is &#x60;false&#x60;.
 **Beta** | **bool** | Whether to read phone numbers that are new to the Twilio platform. Can be: &#x60;true&#x60; or &#x60;false&#x60; and the default is &#x60;true&#x60;.
 **NearNumber** | **string** | Given a phone number, find a geographically close number within &#x60;distance&#x60; miles. Distance defaults to 25 miles. Applies to only phone numbers in the US and Canada.
 **NearLatLong** | **string** | Given a latitude/longitude pair &#x60;lat,long&#x60; find geographically close numbers within &#x60;distance&#x60; miles. Applies to only phone numbers in the US and Canada.
 **Distance** | **int32** | The search radius, in miles, for a &#x60;near_&#x60; query.  Can be up to &#x60;500&#x60; and the default is &#x60;25&#x60;. Applies to only phone numbers in the US and Canada.
 **InPostalCode** | **string** | Limit results to a particular postal code. Given a phone number, search within the same postal code as that number. Applies to only phone numbers in the US and Canada.
 **InRegion** | **string** | Limit results to a particular region, state, or province. Given a phone number, search within the same region as that number. Applies to only phone numbers in the US and Canada.
 **InRateCenter** | **string** | Limit results to a specific rate center, or given a phone number search within the same rate center as that number. Requires &#x60;in_lata&#x60; to be set as well. Applies to only phone numbers in the US and Canada.
 **InLata** | **string** | Limit results to a specific local access and transport area ([LATA](https://en.wikipedia.org/wiki/Local_access_and_transport_area)). Given a phone number, search within the same [LATA](https://en.wikipedia.org/wiki/Local_access_and_transport_area) as that number. Applies to only phone numbers in the US and Canada.
 **InLocality** | **string** | Limit results to a particular locality or city. Given a phone number, search within the same Locality as that number.
 **FaxEnabled** | **bool** | Whether the phone numbers can receive faxes. Can be: &#x60;true&#x60; or &#x60;false&#x60;.
 **PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListAvailablePhoneNumberTollFreeResponse**](ListAvailablePhoneNumberTollFreeResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAvailablePhoneNumberVoip

> ListAvailablePhoneNumberVoipResponse ListAvailablePhoneNumberVoip(ctx, AccountSid, CountryCode).AreaCode(AreaCode).Contains(Contains).SmsEnabled(SmsEnabled).MmsEnabled(MmsEnabled).VoiceEnabled(VoiceEnabled).ExcludeAllAddressRequired(ExcludeAllAddressRequired).ExcludeLocalAddressRequired(ExcludeLocalAddressRequired).ExcludeForeignAddressRequired(ExcludeForeignAddressRequired).Beta(Beta).NearNumber(NearNumber).NearLatLong(NearLatLong).Distance(Distance).InPostalCode(InPostalCode).InRegion(InRegion).InRateCenter(InRateCenter).InLata(InLata).InLocality(InLocality).FaxEnabled(FaxEnabled).PageSize(PageSize).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AccountSid := "AccountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) requesting the AvailablePhoneNumber resources.
    CountryCode := "CountryCode_example" // string | The [ISO-3166-1](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country code of the country from which to read phone numbers.
    AreaCode := int32(56) // int32 | The area code of the phone numbers to read. Applies to only phone numbers in the US and Canada. (optional)
    Contains := "Contains_example" // string | The pattern on which to match phone numbers. Valid characters are `*`, `0-9`, `a-z`, and `A-Z`. The `*` character matches any single digit. For examples, see [Example 2](https://www.twilio.com/docs/phone-numbers/api/availablephonenumber-resource#local-get-basic-example-2) and [Example 3](https://www.twilio.com/docs/phone-numbers/api/availablephonenumber-resource#local-get-basic-example-3). If specified, this value must have at least two characters. (optional)
    SmsEnabled := true // bool | Whether the phone numbers can receive text messages. Can be: `true` or `false`. (optional)
    MmsEnabled := true // bool | Whether the phone numbers can receive MMS messages. Can be: `true` or `false`. (optional)
    VoiceEnabled := true // bool | Whether the phone numbers can receive calls. Can be: `true` or `false`. (optional)
    ExcludeAllAddressRequired := true // bool | Whether to exclude phone numbers that require an [Address](https://www.twilio.com/docs/usage/api/address). Can be: `true` or `false` and the default is `false`. (optional)
    ExcludeLocalAddressRequired := true // bool | Whether to exclude phone numbers that require a local [Address](https://www.twilio.com/docs/usage/api/address). Can be: `true` or `false` and the default is `false`. (optional)
    ExcludeForeignAddressRequired := true // bool | Whether to exclude phone numbers that require a foreign [Address](https://www.twilio.com/docs/usage/api/address). Can be: `true` or `false` and the default is `false`. (optional)
    Beta := true // bool | Whether to read phone numbers that are new to the Twilio platform. Can be: `true` or `false` and the default is `true`. (optional)
    NearNumber := "NearNumber_example" // string | Given a phone number, find a geographically close number within `distance` miles. Distance defaults to 25 miles. Applies to only phone numbers in the US and Canada. (optional)
    NearLatLong := "NearLatLong_example" // string | Given a latitude/longitude pair `lat,long` find geographically close numbers within `distance` miles. Applies to only phone numbers in the US and Canada. (optional)
    Distance := int32(56) // int32 | The search radius, in miles, for a `near_` query.  Can be up to `500` and the default is `25`. Applies to only phone numbers in the US and Canada. (optional)
    InPostalCode := "InPostalCode_example" // string | Limit results to a particular postal code. Given a phone number, search within the same postal code as that number. Applies to only phone numbers in the US and Canada. (optional)
    InRegion := "InRegion_example" // string | Limit results to a particular region, state, or province. Given a phone number, search within the same region as that number. Applies to only phone numbers in the US and Canada. (optional)
    InRateCenter := "InRateCenter_example" // string | Limit results to a specific rate center, or given a phone number search within the same rate center as that number. Requires `in_lata` to be set as well. Applies to only phone numbers in the US and Canada. (optional)
    InLata := "InLata_example" // string | Limit results to a specific local access and transport area ([LATA](https://en.wikipedia.org/wiki/Local_access_and_transport_area)). Given a phone number, search within the same [LATA](https://en.wikipedia.org/wiki/Local_access_and_transport_area) as that number. Applies to only phone numbers in the US and Canada. (optional)
    InLocality := "InLocality_example" // string | Limit results to a particular locality or city. Given a phone number, search within the same Locality as that number. (optional)
    FaxEnabled := true // bool | Whether the phone numbers can receive faxes. Can be: `true` or `false`. (optional)
    PageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListAvailablePhoneNumberVoip(context.Background(), AccountSid, CountryCode).AreaCode(AreaCode).Contains(Contains).SmsEnabled(SmsEnabled).MmsEnabled(MmsEnabled).VoiceEnabled(VoiceEnabled).ExcludeAllAddressRequired(ExcludeAllAddressRequired).ExcludeLocalAddressRequired(ExcludeLocalAddressRequired).ExcludeForeignAddressRequired(ExcludeForeignAddressRequired).Beta(Beta).NearNumber(NearNumber).NearLatLong(NearLatLong).Distance(Distance).InPostalCode(InPostalCode).InRegion(InRegion).InRateCenter(InRateCenter).InLata(InLata).InLocality(InLocality).FaxEnabled(FaxEnabled).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListAvailablePhoneNumberVoip``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAvailablePhoneNumberVoip`: ListAvailablePhoneNumberVoipResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListAvailablePhoneNumberVoip`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) requesting the AvailablePhoneNumber resources.
**CountryCode** | **string** | The [ISO-3166-1](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country code of the country from which to read phone numbers.

### Other Parameters

Other parameters are passed through a pointer to a ListAvailablePhoneNumberVoipParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


 **AreaCode** | **int32** | The area code of the phone numbers to read. Applies to only phone numbers in the US and Canada.
 **Contains** | **string** | The pattern on which to match phone numbers. Valid characters are &#x60;*&#x60;, &#x60;0-9&#x60;, &#x60;a-z&#x60;, and &#x60;A-Z&#x60;. The &#x60;*&#x60; character matches any single digit. For examples, see [Example 2](https://www.twilio.com/docs/phone-numbers/api/availablephonenumber-resource#local-get-basic-example-2) and [Example 3](https://www.twilio.com/docs/phone-numbers/api/availablephonenumber-resource#local-get-basic-example-3). If specified, this value must have at least two characters.
 **SmsEnabled** | **bool** | Whether the phone numbers can receive text messages. Can be: &#x60;true&#x60; or &#x60;false&#x60;.
 **MmsEnabled** | **bool** | Whether the phone numbers can receive MMS messages. Can be: &#x60;true&#x60; or &#x60;false&#x60;.
 **VoiceEnabled** | **bool** | Whether the phone numbers can receive calls. Can be: &#x60;true&#x60; or &#x60;false&#x60;.
 **ExcludeAllAddressRequired** | **bool** | Whether to exclude phone numbers that require an [Address](https://www.twilio.com/docs/usage/api/address). Can be: &#x60;true&#x60; or &#x60;false&#x60; and the default is &#x60;false&#x60;.
 **ExcludeLocalAddressRequired** | **bool** | Whether to exclude phone numbers that require a local [Address](https://www.twilio.com/docs/usage/api/address). Can be: &#x60;true&#x60; or &#x60;false&#x60; and the default is &#x60;false&#x60;.
 **ExcludeForeignAddressRequired** | **bool** | Whether to exclude phone numbers that require a foreign [Address](https://www.twilio.com/docs/usage/api/address). Can be: &#x60;true&#x60; or &#x60;false&#x60; and the default is &#x60;false&#x60;.
 **Beta** | **bool** | Whether to read phone numbers that are new to the Twilio platform. Can be: &#x60;true&#x60; or &#x60;false&#x60; and the default is &#x60;true&#x60;.
 **NearNumber** | **string** | Given a phone number, find a geographically close number within &#x60;distance&#x60; miles. Distance defaults to 25 miles. Applies to only phone numbers in the US and Canada.
 **NearLatLong** | **string** | Given a latitude/longitude pair &#x60;lat,long&#x60; find geographically close numbers within &#x60;distance&#x60; miles. Applies to only phone numbers in the US and Canada.
 **Distance** | **int32** | The search radius, in miles, for a &#x60;near_&#x60; query.  Can be up to &#x60;500&#x60; and the default is &#x60;25&#x60;. Applies to only phone numbers in the US and Canada.
 **InPostalCode** | **string** | Limit results to a particular postal code. Given a phone number, search within the same postal code as that number. Applies to only phone numbers in the US and Canada.
 **InRegion** | **string** | Limit results to a particular region, state, or province. Given a phone number, search within the same region as that number. Applies to only phone numbers in the US and Canada.
 **InRateCenter** | **string** | Limit results to a specific rate center, or given a phone number search within the same rate center as that number. Requires &#x60;in_lata&#x60; to be set as well. Applies to only phone numbers in the US and Canada.
 **InLata** | **string** | Limit results to a specific local access and transport area ([LATA](https://en.wikipedia.org/wiki/Local_access_and_transport_area)). Given a phone number, search within the same [LATA](https://en.wikipedia.org/wiki/Local_access_and_transport_area) as that number. Applies to only phone numbers in the US and Canada.
 **InLocality** | **string** | Limit results to a particular locality or city. Given a phone number, search within the same Locality as that number.
 **FaxEnabled** | **bool** | Whether the phone numbers can receive faxes. Can be: &#x60;true&#x60; or &#x60;false&#x60;.
 **PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListAvailablePhoneNumberVoipResponse**](ListAvailablePhoneNumberVoipResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCall

> ListCallResponse ListCall(ctx, AccountSid).To(To).From(From).ParentCallSid(ParentCallSid).Status(Status).StartTime(StartTime).StartTimeBefore(StartTimeBefore).StartTimeAfter(StartTimeAfter).EndTime(EndTime).EndTimeBefore(EndTimeBefore).EndTimeAfter(EndTimeAfter).PageSize(PageSize).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    AccountSid := "AccountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Call resource(s) to read.
    To := "To_example" // string | Only show calls made to this phone number, SIP address, Client identifier or SIM SID. (optional)
    From := "From_example" // string | Only include calls from this phone number, SIP address, Client identifier or SIM SID. (optional)
    ParentCallSid := "ParentCallSid_example" // string | Only include calls spawned by calls with this SID. (optional)
    Status := "Status_example" // string | The status of the calls to include. Can be: `queued`, `ringing`, `in-progress`, `canceled`, `completed`, `failed`, `busy`, or `no-answer`. (optional)
    StartTime := time.Now() // time.Time | Only include calls that started on this date. Specify a date as `YYYY-MM-DD` in GMT, for example: `2009-07-06`, to read only calls that started on this date. You can also specify an inequality, such as `StartTime<=YYYY-MM-DD`, to read calls that started on or before midnight of this date, and `StartTime>=YYYY-MM-DD` to read calls that started on or after midnight of this date. (optional)
    StartTimeBefore := time.Now() // time.Time | Only include calls that started on this date. Specify a date as `YYYY-MM-DD` in GMT, for example: `2009-07-06`, to read only calls that started on this date. You can also specify an inequality, such as `StartTime<=YYYY-MM-DD`, to read calls that started on or before midnight of this date, and `StartTime>=YYYY-MM-DD` to read calls that started on or after midnight of this date. (optional)
    StartTimeAfter := time.Now() // time.Time | Only include calls that started on this date. Specify a date as `YYYY-MM-DD` in GMT, for example: `2009-07-06`, to read only calls that started on this date. You can also specify an inequality, such as `StartTime<=YYYY-MM-DD`, to read calls that started on or before midnight of this date, and `StartTime>=YYYY-MM-DD` to read calls that started on or after midnight of this date. (optional)
    EndTime := time.Now() // time.Time | Only include calls that ended on this date. Specify a date as `YYYY-MM-DD` in GMT, for example: `2009-07-06`, to read only calls that ended on this date. You can also specify an inequality, such as `EndTime<=YYYY-MM-DD`, to read calls that ended on or before midnight of this date, and `EndTime>=YYYY-MM-DD` to read calls that ended on or after midnight of this date. (optional)
    EndTimeBefore := time.Now() // time.Time | Only include calls that ended on this date. Specify a date as `YYYY-MM-DD` in GMT, for example: `2009-07-06`, to read only calls that ended on this date. You can also specify an inequality, such as `EndTime<=YYYY-MM-DD`, to read calls that ended on or before midnight of this date, and `EndTime>=YYYY-MM-DD` to read calls that ended on or after midnight of this date. (optional)
    EndTimeAfter := time.Now() // time.Time | Only include calls that ended on this date. Specify a date as `YYYY-MM-DD` in GMT, for example: `2009-07-06`, to read only calls that ended on this date. You can also specify an inequality, such as `EndTime<=YYYY-MM-DD`, to read calls that ended on or before midnight of this date, and `EndTime>=YYYY-MM-DD` to read calls that ended on or after midnight of this date. (optional)
    PageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListCall(context.Background(), AccountSid).To(To).From(From).ParentCallSid(ParentCallSid).Status(Status).StartTime(StartTime).StartTimeBefore(StartTimeBefore).StartTimeAfter(StartTimeAfter).EndTime(EndTime).EndTimeBefore(EndTimeBefore).EndTimeAfter(EndTimeAfter).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListCall``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCall`: ListCallResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListCall`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Call resource(s) to read.

### Other Parameters

Other parameters are passed through a pointer to a ListCallParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **To** | **string** | Only show calls made to this phone number, SIP address, Client identifier or SIM SID.
 **From** | **string** | Only include calls from this phone number, SIP address, Client identifier or SIM SID.
 **ParentCallSid** | **string** | Only include calls spawned by calls with this SID.
 **Status** | **string** | The status of the calls to include. Can be: &#x60;queued&#x60;, &#x60;ringing&#x60;, &#x60;in-progress&#x60;, &#x60;canceled&#x60;, &#x60;completed&#x60;, &#x60;failed&#x60;, &#x60;busy&#x60;, or &#x60;no-answer&#x60;.
 **StartTime** | **time.Time** | Only include calls that started on this date. Specify a date as &#x60;YYYY-MM-DD&#x60; in GMT, for example: &#x60;2009-07-06&#x60;, to read only calls that started on this date. You can also specify an inequality, such as &#x60;StartTime&lt;&#x3D;YYYY-MM-DD&#x60;, to read calls that started on or before midnight of this date, and &#x60;StartTime&gt;&#x3D;YYYY-MM-DD&#x60; to read calls that started on or after midnight of this date.
 **StartTimeBefore** | **time.Time** | Only include calls that started on this date. Specify a date as &#x60;YYYY-MM-DD&#x60; in GMT, for example: &#x60;2009-07-06&#x60;, to read only calls that started on this date. You can also specify an inequality, such as &#x60;StartTime&lt;&#x3D;YYYY-MM-DD&#x60;, to read calls that started on or before midnight of this date, and &#x60;StartTime&gt;&#x3D;YYYY-MM-DD&#x60; to read calls that started on or after midnight of this date.
 **StartTimeAfter** | **time.Time** | Only include calls that started on this date. Specify a date as &#x60;YYYY-MM-DD&#x60; in GMT, for example: &#x60;2009-07-06&#x60;, to read only calls that started on this date. You can also specify an inequality, such as &#x60;StartTime&lt;&#x3D;YYYY-MM-DD&#x60;, to read calls that started on or before midnight of this date, and &#x60;StartTime&gt;&#x3D;YYYY-MM-DD&#x60; to read calls that started on or after midnight of this date.
 **EndTime** | **time.Time** | Only include calls that ended on this date. Specify a date as &#x60;YYYY-MM-DD&#x60; in GMT, for example: &#x60;2009-07-06&#x60;, to read only calls that ended on this date. You can also specify an inequality, such as &#x60;EndTime&lt;&#x3D;YYYY-MM-DD&#x60;, to read calls that ended on or before midnight of this date, and &#x60;EndTime&gt;&#x3D;YYYY-MM-DD&#x60; to read calls that ended on or after midnight of this date.
 **EndTimeBefore** | **time.Time** | Only include calls that ended on this date. Specify a date as &#x60;YYYY-MM-DD&#x60; in GMT, for example: &#x60;2009-07-06&#x60;, to read only calls that ended on this date. You can also specify an inequality, such as &#x60;EndTime&lt;&#x3D;YYYY-MM-DD&#x60;, to read calls that ended on or before midnight of this date, and &#x60;EndTime&gt;&#x3D;YYYY-MM-DD&#x60; to read calls that ended on or after midnight of this date.
 **EndTimeAfter** | **time.Time** | Only include calls that ended on this date. Specify a date as &#x60;YYYY-MM-DD&#x60; in GMT, for example: &#x60;2009-07-06&#x60;, to read only calls that ended on this date. You can also specify an inequality, such as &#x60;EndTime&lt;&#x3D;YYYY-MM-DD&#x60;, to read calls that ended on or before midnight of this date, and &#x60;EndTime&gt;&#x3D;YYYY-MM-DD&#x60; to read calls that ended on or after midnight of this date.
 **PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListCallResponse**](ListCallResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCallEvent

> ListCallEventResponse ListCallEvent(ctx, AccountSid, CallSid).PageSize(PageSize).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AccountSid := "AccountSid_example" // string | The unique SID identifier of the Account.
    CallSid := "CallSid_example" // string | The unique SID identifier of the Call.
    PageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListCallEvent(context.Background(), AccountSid, CallSid).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListCallEvent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCallEvent`: ListCallEventResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListCallEvent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string** | The unique SID identifier of the Account.
**CallSid** | **string** | The unique SID identifier of the Call.

### Other Parameters

Other parameters are passed through a pointer to a ListCallEventParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


 **PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListCallEventResponse**](ListCallEventResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCallNotification

> ListCallNotificationResponse ListCallNotification(ctx, AccountSid, CallSid).Log(Log).MessageDate(MessageDate).MessageDateBefore(MessageDateBefore).MessageDateAfter(MessageDateAfter).PageSize(PageSize).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    AccountSid := "AccountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Call Notification resources to read.
    CallSid := "CallSid_example" // string | The [Call](https://www.twilio.com/docs/voice/api/call-resource) SID of the Call Notification resources to read.
    Log := int32(56) // int32 | Only read notifications of the specified log level. Can be:  `0` to read only ERROR notifications or `1` to read only WARNING notifications. By default, all notifications are read. (optional)
    MessageDate := time.Now() // string | Only show notifications for the specified date, formatted as `YYYY-MM-DD`. You can also specify an inequality, such as `<=YYYY-MM-DD` for messages logged at or before midnight on a date, or `>=YYYY-MM-DD` for messages logged at or after midnight on a date. (optional)
    MessageDateBefore := time.Now() // string | Only show notifications for the specified date, formatted as `YYYY-MM-DD`. You can also specify an inequality, such as `<=YYYY-MM-DD` for messages logged at or before midnight on a date, or `>=YYYY-MM-DD` for messages logged at or after midnight on a date. (optional)
    MessageDateAfter := time.Now() // string | Only show notifications for the specified date, formatted as `YYYY-MM-DD`. You can also specify an inequality, such as `<=YYYY-MM-DD` for messages logged at or before midnight on a date, or `>=YYYY-MM-DD` for messages logged at or after midnight on a date. (optional)
    PageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListCallNotification(context.Background(), AccountSid, CallSid).Log(Log).MessageDate(MessageDate).MessageDateBefore(MessageDateBefore).MessageDateAfter(MessageDateAfter).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListCallNotification``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCallNotification`: ListCallNotificationResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListCallNotification`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Call Notification resources to read.
**CallSid** | **string** | The [Call](https://www.twilio.com/docs/voice/api/call-resource) SID of the Call Notification resources to read.

### Other Parameters

Other parameters are passed through a pointer to a ListCallNotificationParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


 **Log** | **int32** | Only read notifications of the specified log level. Can be:  &#x60;0&#x60; to read only ERROR notifications or &#x60;1&#x60; to read only WARNING notifications. By default, all notifications are read.
 **MessageDate** | **string** | Only show notifications for the specified date, formatted as &#x60;YYYY-MM-DD&#x60;. You can also specify an inequality, such as &#x60;&lt;&#x3D;YYYY-MM-DD&#x60; for messages logged at or before midnight on a date, or &#x60;&gt;&#x3D;YYYY-MM-DD&#x60; for messages logged at or after midnight on a date.
 **MessageDateBefore** | **string** | Only show notifications for the specified date, formatted as &#x60;YYYY-MM-DD&#x60;. You can also specify an inequality, such as &#x60;&lt;&#x3D;YYYY-MM-DD&#x60; for messages logged at or before midnight on a date, or &#x60;&gt;&#x3D;YYYY-MM-DD&#x60; for messages logged at or after midnight on a date.
 **MessageDateAfter** | **string** | Only show notifications for the specified date, formatted as &#x60;YYYY-MM-DD&#x60;. You can also specify an inequality, such as &#x60;&lt;&#x3D;YYYY-MM-DD&#x60; for messages logged at or before midnight on a date, or &#x60;&gt;&#x3D;YYYY-MM-DD&#x60; for messages logged at or after midnight on a date.
 **PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListCallNotificationResponse**](ListCallNotificationResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCallRecording

> ListCallRecordingResponse ListCallRecording(ctx, AccountSid, CallSid).DateCreated(DateCreated).DateCreatedBefore(DateCreatedBefore).DateCreatedAfter(DateCreatedAfter).PageSize(PageSize).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    AccountSid := "AccountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Recording resources to read.
    CallSid := "CallSid_example" // string | The [Call](https://www.twilio.com/docs/voice/api/call-resource) SID of the resources to read.
    DateCreated := time.Now() // string | The `date_created` value, specified as `YYYY-MM-DD`, of the resources to read. You can also specify inequality: `DateCreated<=YYYY-MM-DD` will return recordings generated at or before midnight on a given date, and `DateCreated>=YYYY-MM-DD` returns recordings generated at or after midnight on a date. (optional)
    DateCreatedBefore := time.Now() // string | The `date_created` value, specified as `YYYY-MM-DD`, of the resources to read. You can also specify inequality: `DateCreated<=YYYY-MM-DD` will return recordings generated at or before midnight on a given date, and `DateCreated>=YYYY-MM-DD` returns recordings generated at or after midnight on a date. (optional)
    DateCreatedAfter := time.Now() // string | The `date_created` value, specified as `YYYY-MM-DD`, of the resources to read. You can also specify inequality: `DateCreated<=YYYY-MM-DD` will return recordings generated at or before midnight on a given date, and `DateCreated>=YYYY-MM-DD` returns recordings generated at or after midnight on a date. (optional)
    PageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListCallRecording(context.Background(), AccountSid, CallSid).DateCreated(DateCreated).DateCreatedBefore(DateCreatedBefore).DateCreatedAfter(DateCreatedAfter).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListCallRecording``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCallRecording`: ListCallRecordingResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListCallRecording`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Recording resources to read.
**CallSid** | **string** | The [Call](https://www.twilio.com/docs/voice/api/call-resource) SID of the resources to read.

### Other Parameters

Other parameters are passed through a pointer to a ListCallRecordingParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


 **DateCreated** | **string** | The &#x60;date_created&#x60; value, specified as &#x60;YYYY-MM-DD&#x60;, of the resources to read. You can also specify inequality: &#x60;DateCreated&lt;&#x3D;YYYY-MM-DD&#x60; will return recordings generated at or before midnight on a given date, and &#x60;DateCreated&gt;&#x3D;YYYY-MM-DD&#x60; returns recordings generated at or after midnight on a date.
 **DateCreatedBefore** | **string** | The &#x60;date_created&#x60; value, specified as &#x60;YYYY-MM-DD&#x60;, of the resources to read. You can also specify inequality: &#x60;DateCreated&lt;&#x3D;YYYY-MM-DD&#x60; will return recordings generated at or before midnight on a given date, and &#x60;DateCreated&gt;&#x3D;YYYY-MM-DD&#x60; returns recordings generated at or after midnight on a date.
 **DateCreatedAfter** | **string** | The &#x60;date_created&#x60; value, specified as &#x60;YYYY-MM-DD&#x60;, of the resources to read. You can also specify inequality: &#x60;DateCreated&lt;&#x3D;YYYY-MM-DD&#x60; will return recordings generated at or before midnight on a given date, and &#x60;DateCreated&gt;&#x3D;YYYY-MM-DD&#x60; returns recordings generated at or after midnight on a date.
 **PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListCallRecordingResponse**](ListCallRecordingResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListConference

> ListConferenceResponse ListConference(ctx, AccountSid).DateCreated(DateCreated).DateCreatedBefore(DateCreatedBefore).DateCreatedAfter(DateCreatedAfter).DateUpdated(DateUpdated).DateUpdatedBefore(DateUpdatedBefore).DateUpdatedAfter(DateUpdatedAfter).FriendlyName(FriendlyName).Status(Status).PageSize(PageSize).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    AccountSid := "AccountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Conference resource(s) to read.
    DateCreated := time.Now() // string | The `date_created` value, specified as `YYYY-MM-DD`, of the resources to read. To read conferences that started on or before midnight on a date, use `<=YYYY-MM-DD`, and to specify  conferences that started on or after midnight on a date, use `>=YYYY-MM-DD`. (optional)
    DateCreatedBefore := time.Now() // string | The `date_created` value, specified as `YYYY-MM-DD`, of the resources to read. To read conferences that started on or before midnight on a date, use `<=YYYY-MM-DD`, and to specify  conferences that started on or after midnight on a date, use `>=YYYY-MM-DD`. (optional)
    DateCreatedAfter := time.Now() // string | The `date_created` value, specified as `YYYY-MM-DD`, of the resources to read. To read conferences that started on or before midnight on a date, use `<=YYYY-MM-DD`, and to specify  conferences that started on or after midnight on a date, use `>=YYYY-MM-DD`. (optional)
    DateUpdated := time.Now() // string | The `date_updated` value, specified as `YYYY-MM-DD`, of the resources to read. To read conferences that were last updated on or before midnight on a date, use `<=YYYY-MM-DD`, and to specify conferences that were last updated on or after midnight on a given date, use  `>=YYYY-MM-DD`. (optional)
    DateUpdatedBefore := time.Now() // string | The `date_updated` value, specified as `YYYY-MM-DD`, of the resources to read. To read conferences that were last updated on or before midnight on a date, use `<=YYYY-MM-DD`, and to specify conferences that were last updated on or after midnight on a given date, use  `>=YYYY-MM-DD`. (optional)
    DateUpdatedAfter := time.Now() // string | The `date_updated` value, specified as `YYYY-MM-DD`, of the resources to read. To read conferences that were last updated on or before midnight on a date, use `<=YYYY-MM-DD`, and to specify conferences that were last updated on or after midnight on a given date, use  `>=YYYY-MM-DD`. (optional)
    FriendlyName := "FriendlyName_example" // string | The string that identifies the Conference resources to read. (optional)
    Status := "Status_example" // string | The status of the resources to read. Can be: `init`, `in-progress`, or `completed`. (optional)
    PageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListConference(context.Background(), AccountSid).DateCreated(DateCreated).DateCreatedBefore(DateCreatedBefore).DateCreatedAfter(DateCreatedAfter).DateUpdated(DateUpdated).DateUpdatedBefore(DateUpdatedBefore).DateUpdatedAfter(DateUpdatedAfter).FriendlyName(FriendlyName).Status(Status).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListConference``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListConference`: ListConferenceResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListConference`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Conference resource(s) to read.

### Other Parameters

Other parameters are passed through a pointer to a ListConferenceParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **DateCreated** | **string** | The &#x60;date_created&#x60; value, specified as &#x60;YYYY-MM-DD&#x60;, of the resources to read. To read conferences that started on or before midnight on a date, use &#x60;&lt;&#x3D;YYYY-MM-DD&#x60;, and to specify  conferences that started on or after midnight on a date, use &#x60;&gt;&#x3D;YYYY-MM-DD&#x60;.
 **DateCreatedBefore** | **string** | The &#x60;date_created&#x60; value, specified as &#x60;YYYY-MM-DD&#x60;, of the resources to read. To read conferences that started on or before midnight on a date, use &#x60;&lt;&#x3D;YYYY-MM-DD&#x60;, and to specify  conferences that started on or after midnight on a date, use &#x60;&gt;&#x3D;YYYY-MM-DD&#x60;.
 **DateCreatedAfter** | **string** | The &#x60;date_created&#x60; value, specified as &#x60;YYYY-MM-DD&#x60;, of the resources to read. To read conferences that started on or before midnight on a date, use &#x60;&lt;&#x3D;YYYY-MM-DD&#x60;, and to specify  conferences that started on or after midnight on a date, use &#x60;&gt;&#x3D;YYYY-MM-DD&#x60;.
 **DateUpdated** | **string** | The &#x60;date_updated&#x60; value, specified as &#x60;YYYY-MM-DD&#x60;, of the resources to read. To read conferences that were last updated on or before midnight on a date, use &#x60;&lt;&#x3D;YYYY-MM-DD&#x60;, and to specify conferences that were last updated on or after midnight on a given date, use  &#x60;&gt;&#x3D;YYYY-MM-DD&#x60;.
 **DateUpdatedBefore** | **string** | The &#x60;date_updated&#x60; value, specified as &#x60;YYYY-MM-DD&#x60;, of the resources to read. To read conferences that were last updated on or before midnight on a date, use &#x60;&lt;&#x3D;YYYY-MM-DD&#x60;, and to specify conferences that were last updated on or after midnight on a given date, use  &#x60;&gt;&#x3D;YYYY-MM-DD&#x60;.
 **DateUpdatedAfter** | **string** | The &#x60;date_updated&#x60; value, specified as &#x60;YYYY-MM-DD&#x60;, of the resources to read. To read conferences that were last updated on or before midnight on a date, use &#x60;&lt;&#x3D;YYYY-MM-DD&#x60;, and to specify conferences that were last updated on or after midnight on a given date, use  &#x60;&gt;&#x3D;YYYY-MM-DD&#x60;.
 **FriendlyName** | **string** | The string that identifies the Conference resources to read.
 **Status** | **string** | The status of the resources to read. Can be: &#x60;init&#x60;, &#x60;in-progress&#x60;, or &#x60;completed&#x60;.
 **PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListConferenceResponse**](ListConferenceResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListConferenceRecording

> ListConferenceRecordingResponse ListConferenceRecording(ctx, AccountSid, ConferenceSid).DateCreated(DateCreated).DateCreatedBefore(DateCreatedBefore).DateCreatedAfter(DateCreatedAfter).PageSize(PageSize).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    AccountSid := "AccountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Conference Recording resources to read.
    ConferenceSid := "ConferenceSid_example" // string | The Conference SID that identifies the conference associated with the recording to read.
    DateCreated := time.Now() // string | The `date_created` value, specified as `YYYY-MM-DD`, of the resources to read. You can also specify inequality: `DateCreated<=YYYY-MM-DD` will return recordings generated at or before midnight on a given date, and `DateCreated>=YYYY-MM-DD` returns recordings generated at or after midnight on a date. (optional)
    DateCreatedBefore := time.Now() // string | The `date_created` value, specified as `YYYY-MM-DD`, of the resources to read. You can also specify inequality: `DateCreated<=YYYY-MM-DD` will return recordings generated at or before midnight on a given date, and `DateCreated>=YYYY-MM-DD` returns recordings generated at or after midnight on a date. (optional)
    DateCreatedAfter := time.Now() // string | The `date_created` value, specified as `YYYY-MM-DD`, of the resources to read. You can also specify inequality: `DateCreated<=YYYY-MM-DD` will return recordings generated at or before midnight on a given date, and `DateCreated>=YYYY-MM-DD` returns recordings generated at or after midnight on a date. (optional)
    PageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListConferenceRecording(context.Background(), AccountSid, ConferenceSid).DateCreated(DateCreated).DateCreatedBefore(DateCreatedBefore).DateCreatedAfter(DateCreatedAfter).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListConferenceRecording``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListConferenceRecording`: ListConferenceRecordingResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListConferenceRecording`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Conference Recording resources to read.
**ConferenceSid** | **string** | The Conference SID that identifies the conference associated with the recording to read.

### Other Parameters

Other parameters are passed through a pointer to a ListConferenceRecordingParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


 **DateCreated** | **string** | The &#x60;date_created&#x60; value, specified as &#x60;YYYY-MM-DD&#x60;, of the resources to read. You can also specify inequality: &#x60;DateCreated&lt;&#x3D;YYYY-MM-DD&#x60; will return recordings generated at or before midnight on a given date, and &#x60;DateCreated&gt;&#x3D;YYYY-MM-DD&#x60; returns recordings generated at or after midnight on a date.
 **DateCreatedBefore** | **string** | The &#x60;date_created&#x60; value, specified as &#x60;YYYY-MM-DD&#x60;, of the resources to read. You can also specify inequality: &#x60;DateCreated&lt;&#x3D;YYYY-MM-DD&#x60; will return recordings generated at or before midnight on a given date, and &#x60;DateCreated&gt;&#x3D;YYYY-MM-DD&#x60; returns recordings generated at or after midnight on a date.
 **DateCreatedAfter** | **string** | The &#x60;date_created&#x60; value, specified as &#x60;YYYY-MM-DD&#x60;, of the resources to read. You can also specify inequality: &#x60;DateCreated&lt;&#x3D;YYYY-MM-DD&#x60; will return recordings generated at or before midnight on a given date, and &#x60;DateCreated&gt;&#x3D;YYYY-MM-DD&#x60; returns recordings generated at or after midnight on a date.
 **PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListConferenceRecordingResponse**](ListConferenceRecordingResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListConnectApp

> ListConnectAppResponse ListConnectApp(ctx, AccountSid).PageSize(PageSize).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AccountSid := "AccountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the ConnectApp resources to read.
    PageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListConnectApp(context.Background(), AccountSid).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListConnectApp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListConnectApp`: ListConnectAppResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListConnectApp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the ConnectApp resources to read.

### Other Parameters

Other parameters are passed through a pointer to a ListConnectAppParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListConnectAppResponse**](ListConnectAppResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDependentPhoneNumber

> ListDependentPhoneNumberResponse ListDependentPhoneNumber(ctx, AccountSid, AddressSid).PageSize(PageSize).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AccountSid := "AccountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the DependentPhoneNumber resources to read.
    AddressSid := "AddressSid_example" // string | The SID of the Address resource associated with the phone number.
    PageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListDependentPhoneNumber(context.Background(), AccountSid, AddressSid).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListDependentPhoneNumber``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListDependentPhoneNumber`: ListDependentPhoneNumberResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListDependentPhoneNumber`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the DependentPhoneNumber resources to read.
**AddressSid** | **string** | The SID of the Address resource associated with the phone number.

### Other Parameters

Other parameters are passed through a pointer to a ListDependentPhoneNumberParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


 **PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListDependentPhoneNumberResponse**](ListDependentPhoneNumberResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListIncomingPhoneNumber

> ListIncomingPhoneNumberResponse ListIncomingPhoneNumber(ctx, AccountSid).Beta(Beta).FriendlyName(FriendlyName).PhoneNumber(PhoneNumber).Origin(Origin).PageSize(PageSize).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AccountSid := "AccountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the IncomingPhoneNumber resources to read.
    Beta := true // bool | Whether to include phone numbers new to the Twilio platform. Can be: `true` or `false` and the default is `true`. (optional)
    FriendlyName := "FriendlyName_example" // string | A string that identifies the IncomingPhoneNumber resources to read. (optional)
    PhoneNumber := "PhoneNumber_example" // string | The phone numbers of the IncomingPhoneNumber resources to read. You can specify partial numbers and use '*' as a wildcard for any digit. (optional)
    Origin := "Origin_example" // string | Whether to include phone numbers based on their origin. Can be: `twilio` or `hosted`. By default, phone numbers of all origin are included. (optional)
    PageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListIncomingPhoneNumber(context.Background(), AccountSid).Beta(Beta).FriendlyName(FriendlyName).PhoneNumber(PhoneNumber).Origin(Origin).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListIncomingPhoneNumber``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListIncomingPhoneNumber`: ListIncomingPhoneNumberResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListIncomingPhoneNumber`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the IncomingPhoneNumber resources to read.

### Other Parameters

Other parameters are passed through a pointer to a ListIncomingPhoneNumberParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **Beta** | **bool** | Whether to include phone numbers new to the Twilio platform. Can be: &#x60;true&#x60; or &#x60;false&#x60; and the default is &#x60;true&#x60;.
 **FriendlyName** | **string** | A string that identifies the IncomingPhoneNumber resources to read.
 **PhoneNumber** | **string** | The phone numbers of the IncomingPhoneNumber resources to read. You can specify partial numbers and use &#39;*&#39; as a wildcard for any digit.
 **Origin** | **string** | Whether to include phone numbers based on their origin. Can be: &#x60;twilio&#x60; or &#x60;hosted&#x60;. By default, phone numbers of all origin are included.
 **PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListIncomingPhoneNumberResponse**](ListIncomingPhoneNumberResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListIncomingPhoneNumberAssignedAddOn

> ListIncomingPhoneNumberAssignedAddOnResponse ListIncomingPhoneNumberAssignedAddOn(ctx, AccountSid, ResourceSid).PageSize(PageSize).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AccountSid := "AccountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the resources to read.
    ResourceSid := "ResourceSid_example" // string | The SID of the Phone Number to which the Add-on is assigned.
    PageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListIncomingPhoneNumberAssignedAddOn(context.Background(), AccountSid, ResourceSid).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListIncomingPhoneNumberAssignedAddOn``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListIncomingPhoneNumberAssignedAddOn`: ListIncomingPhoneNumberAssignedAddOnResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListIncomingPhoneNumberAssignedAddOn`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the resources to read.
**ResourceSid** | **string** | The SID of the Phone Number to which the Add-on is assigned.

### Other Parameters

Other parameters are passed through a pointer to a ListIncomingPhoneNumberAssignedAddOnParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


 **PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListIncomingPhoneNumberAssignedAddOnResponse**](ListIncomingPhoneNumberAssignedAddOnResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListIncomingPhoneNumberAssignedAddOnExtension

> ListIncomingPhoneNumberAssignedAddOnExtensionResponse ListIncomingPhoneNumberAssignedAddOnExtension(ctx, AccountSid, ResourceSid, AssignedAddOnSid).PageSize(PageSize).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AccountSid := "AccountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the resources to read.
    ResourceSid := "ResourceSid_example" // string | The SID of the Phone Number to which the Add-on is assigned.
    AssignedAddOnSid := "AssignedAddOnSid_example" // string | The SID that uniquely identifies the assigned Add-on installation.
    PageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListIncomingPhoneNumberAssignedAddOnExtension(context.Background(), AccountSid, ResourceSid, AssignedAddOnSid).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListIncomingPhoneNumberAssignedAddOnExtension``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListIncomingPhoneNumberAssignedAddOnExtension`: ListIncomingPhoneNumberAssignedAddOnExtensionResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListIncomingPhoneNumberAssignedAddOnExtension`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the resources to read.
**ResourceSid** | **string** | The SID of the Phone Number to which the Add-on is assigned.
**AssignedAddOnSid** | **string** | The SID that uniquely identifies the assigned Add-on installation.

### Other Parameters

Other parameters are passed through a pointer to a ListIncomingPhoneNumberAssignedAddOnExtensionParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------



 **PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListIncomingPhoneNumberAssignedAddOnExtensionResponse**](ListIncomingPhoneNumberAssignedAddOnExtensionResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListIncomingPhoneNumberLocal

> ListIncomingPhoneNumberLocalResponse ListIncomingPhoneNumberLocal(ctx, AccountSid).Beta(Beta).FriendlyName(FriendlyName).PhoneNumber(PhoneNumber).Origin(Origin).PageSize(PageSize).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AccountSid := "AccountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the resources to read.
    Beta := true // bool | Whether to include phone numbers new to the Twilio platform. Can be: `true` or `false` and the default is `true`. (optional)
    FriendlyName := "FriendlyName_example" // string | A string that identifies the resources to read. (optional)
    PhoneNumber := "PhoneNumber_example" // string | The phone numbers of the IncomingPhoneNumber resources to read. You can specify partial numbers and use '*' as a wildcard for any digit. (optional)
    Origin := "Origin_example" // string | Whether to include phone numbers based on their origin. Can be: `twilio` or `hosted`. By default, phone numbers of all origin are included. (optional)
    PageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListIncomingPhoneNumberLocal(context.Background(), AccountSid).Beta(Beta).FriendlyName(FriendlyName).PhoneNumber(PhoneNumber).Origin(Origin).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListIncomingPhoneNumberLocal``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListIncomingPhoneNumberLocal`: ListIncomingPhoneNumberLocalResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListIncomingPhoneNumberLocal`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the resources to read.

### Other Parameters

Other parameters are passed through a pointer to a ListIncomingPhoneNumberLocalParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **Beta** | **bool** | Whether to include phone numbers new to the Twilio platform. Can be: &#x60;true&#x60; or &#x60;false&#x60; and the default is &#x60;true&#x60;.
 **FriendlyName** | **string** | A string that identifies the resources to read.
 **PhoneNumber** | **string** | The phone numbers of the IncomingPhoneNumber resources to read. You can specify partial numbers and use &#39;*&#39; as a wildcard for any digit.
 **Origin** | **string** | Whether to include phone numbers based on their origin. Can be: &#x60;twilio&#x60; or &#x60;hosted&#x60;. By default, phone numbers of all origin are included.
 **PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListIncomingPhoneNumberLocalResponse**](ListIncomingPhoneNumberLocalResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListIncomingPhoneNumberMobile

> ListIncomingPhoneNumberMobileResponse ListIncomingPhoneNumberMobile(ctx, AccountSid).Beta(Beta).FriendlyName(FriendlyName).PhoneNumber(PhoneNumber).Origin(Origin).PageSize(PageSize).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AccountSid := "AccountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the resources to read.
    Beta := true // bool | Whether to include phone numbers new to the Twilio platform. Can be: `true` or `false` and the default is `true`. (optional)
    FriendlyName := "FriendlyName_example" // string | A string that identifies the resources to read. (optional)
    PhoneNumber := "PhoneNumber_example" // string | The phone numbers of the IncomingPhoneNumber resources to read. You can specify partial numbers and use '*' as a wildcard for any digit. (optional)
    Origin := "Origin_example" // string | Whether to include phone numbers based on their origin. Can be: `twilio` or `hosted`. By default, phone numbers of all origin are included. (optional)
    PageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListIncomingPhoneNumberMobile(context.Background(), AccountSid).Beta(Beta).FriendlyName(FriendlyName).PhoneNumber(PhoneNumber).Origin(Origin).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListIncomingPhoneNumberMobile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListIncomingPhoneNumberMobile`: ListIncomingPhoneNumberMobileResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListIncomingPhoneNumberMobile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the resources to read.

### Other Parameters

Other parameters are passed through a pointer to a ListIncomingPhoneNumberMobileParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **Beta** | **bool** | Whether to include phone numbers new to the Twilio platform. Can be: &#x60;true&#x60; or &#x60;false&#x60; and the default is &#x60;true&#x60;.
 **FriendlyName** | **string** | A string that identifies the resources to read.
 **PhoneNumber** | **string** | The phone numbers of the IncomingPhoneNumber resources to read. You can specify partial numbers and use &#39;*&#39; as a wildcard for any digit.
 **Origin** | **string** | Whether to include phone numbers based on their origin. Can be: &#x60;twilio&#x60; or &#x60;hosted&#x60;. By default, phone numbers of all origin are included.
 **PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListIncomingPhoneNumberMobileResponse**](ListIncomingPhoneNumberMobileResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListIncomingPhoneNumberTollFree

> ListIncomingPhoneNumberTollFreeResponse ListIncomingPhoneNumberTollFree(ctx, AccountSid).Beta(Beta).FriendlyName(FriendlyName).PhoneNumber(PhoneNumber).Origin(Origin).PageSize(PageSize).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AccountSid := "AccountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the resources to read.
    Beta := true // bool | Whether to include phone numbers new to the Twilio platform. Can be: `true` or `false` and the default is `true`. (optional)
    FriendlyName := "FriendlyName_example" // string | A string that identifies the resources to read. (optional)
    PhoneNumber := "PhoneNumber_example" // string | The phone numbers of the IncomingPhoneNumber resources to read. You can specify partial numbers and use '*' as a wildcard for any digit. (optional)
    Origin := "Origin_example" // string | Whether to include phone numbers based on their origin. Can be: `twilio` or `hosted`. By default, phone numbers of all origin are included. (optional)
    PageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListIncomingPhoneNumberTollFree(context.Background(), AccountSid).Beta(Beta).FriendlyName(FriendlyName).PhoneNumber(PhoneNumber).Origin(Origin).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListIncomingPhoneNumberTollFree``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListIncomingPhoneNumberTollFree`: ListIncomingPhoneNumberTollFreeResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListIncomingPhoneNumberTollFree`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the resources to read.

### Other Parameters

Other parameters are passed through a pointer to a ListIncomingPhoneNumberTollFreeParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **Beta** | **bool** | Whether to include phone numbers new to the Twilio platform. Can be: &#x60;true&#x60; or &#x60;false&#x60; and the default is &#x60;true&#x60;.
 **FriendlyName** | **string** | A string that identifies the resources to read.
 **PhoneNumber** | **string** | The phone numbers of the IncomingPhoneNumber resources to read. You can specify partial numbers and use &#39;*&#39; as a wildcard for any digit.
 **Origin** | **string** | Whether to include phone numbers based on their origin. Can be: &#x60;twilio&#x60; or &#x60;hosted&#x60;. By default, phone numbers of all origin are included.
 **PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListIncomingPhoneNumberTollFreeResponse**](ListIncomingPhoneNumberTollFreeResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListKey

> ListKeyResponse ListKey(ctx, AccountSid).PageSize(PageSize).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AccountSid := "AccountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Key resources to read.
    PageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListKey(context.Background(), AccountSid).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListKey`: ListKeyResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Key resources to read.

### Other Parameters

Other parameters are passed through a pointer to a ListKeyParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListKeyResponse**](ListKeyResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMedia

> ListMediaResponse ListMedia(ctx, AccountSid, MessageSid).DateCreated(DateCreated).DateCreatedBefore(DateCreatedBefore).DateCreatedAfter(DateCreatedAfter).PageSize(PageSize).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    AccountSid := "AccountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Media resource(s) to read.
    MessageSid := "MessageSid_example" // string | The SID of the Message resource that this Media resource belongs to.
    DateCreated := time.Now() // time.Time | Only include media that was created on this date. Specify a date as `YYYY-MM-DD` in GMT, for example: `2009-07-06`, to read media that was created on this date. You can also specify an inequality, such as `StartTime<=YYYY-MM-DD`, to read media that was created on or before midnight of this date, and `StartTime>=YYYY-MM-DD` to read media that was created on or after midnight of this date. (optional)
    DateCreatedBefore := time.Now() // time.Time | Only include media that was created on this date. Specify a date as `YYYY-MM-DD` in GMT, for example: `2009-07-06`, to read media that was created on this date. You can also specify an inequality, such as `StartTime<=YYYY-MM-DD`, to read media that was created on or before midnight of this date, and `StartTime>=YYYY-MM-DD` to read media that was created on or after midnight of this date. (optional)
    DateCreatedAfter := time.Now() // time.Time | Only include media that was created on this date. Specify a date as `YYYY-MM-DD` in GMT, for example: `2009-07-06`, to read media that was created on this date. You can also specify an inequality, such as `StartTime<=YYYY-MM-DD`, to read media that was created on or before midnight of this date, and `StartTime>=YYYY-MM-DD` to read media that was created on or after midnight of this date. (optional)
    PageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListMedia(context.Background(), AccountSid, MessageSid).DateCreated(DateCreated).DateCreatedBefore(DateCreatedBefore).DateCreatedAfter(DateCreatedAfter).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListMedia``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListMedia`: ListMediaResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListMedia`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Media resource(s) to read.
**MessageSid** | **string** | The SID of the Message resource that this Media resource belongs to.

### Other Parameters

Other parameters are passed through a pointer to a ListMediaParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


 **DateCreated** | **time.Time** | Only include media that was created on this date. Specify a date as &#x60;YYYY-MM-DD&#x60; in GMT, for example: &#x60;2009-07-06&#x60;, to read media that was created on this date. You can also specify an inequality, such as &#x60;StartTime&lt;&#x3D;YYYY-MM-DD&#x60;, to read media that was created on or before midnight of this date, and &#x60;StartTime&gt;&#x3D;YYYY-MM-DD&#x60; to read media that was created on or after midnight of this date.
 **DateCreatedBefore** | **time.Time** | Only include media that was created on this date. Specify a date as &#x60;YYYY-MM-DD&#x60; in GMT, for example: &#x60;2009-07-06&#x60;, to read media that was created on this date. You can also specify an inequality, such as &#x60;StartTime&lt;&#x3D;YYYY-MM-DD&#x60;, to read media that was created on or before midnight of this date, and &#x60;StartTime&gt;&#x3D;YYYY-MM-DD&#x60; to read media that was created on or after midnight of this date.
 **DateCreatedAfter** | **time.Time** | Only include media that was created on this date. Specify a date as &#x60;YYYY-MM-DD&#x60; in GMT, for example: &#x60;2009-07-06&#x60;, to read media that was created on this date. You can also specify an inequality, such as &#x60;StartTime&lt;&#x3D;YYYY-MM-DD&#x60;, to read media that was created on or before midnight of this date, and &#x60;StartTime&gt;&#x3D;YYYY-MM-DD&#x60; to read media that was created on or after midnight of this date.
 **PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListMediaResponse**](ListMediaResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMember

> ListMemberResponse ListMember(ctx, AccountSid, QueueSid).PageSize(PageSize).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AccountSid := "AccountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Member resource(s) to read.
    QueueSid := "QueueSid_example" // string | The SID of the Queue in which to find the members
    PageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListMember(context.Background(), AccountSid, QueueSid).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListMember``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListMember`: ListMemberResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListMember`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Member resource(s) to read.
**QueueSid** | **string** | The SID of the Queue in which to find the members

### Other Parameters

Other parameters are passed through a pointer to a ListMemberParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


 **PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListMemberResponse**](ListMemberResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMessage

> ListMessageResponse ListMessage(ctx, AccountSid).To(To).From(From).DateSent(DateSent).DateSentBefore(DateSentBefore).DateSentAfter(DateSentAfter).PageSize(PageSize).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    AccountSid := "AccountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Message resources to read.
    To := "To_example" // string | Read messages sent to only this phone number. (optional)
    From := "From_example" // string | Read messages sent from only this phone number or alphanumeric sender ID. (optional)
    DateSent := time.Now() // time.Time | The date of the messages to show. Specify a date as `YYYY-MM-DD` in GMT to read only messages sent on this date. For example: `2009-07-06`. You can also specify an inequality, such as `DateSent<=YYYY-MM-DD`, to read messages sent on or before midnight on a date, and `DateSent>=YYYY-MM-DD` to read messages sent on or after midnight on a date. (optional)
    DateSentBefore := time.Now() // time.Time | The date of the messages to show. Specify a date as `YYYY-MM-DD` in GMT to read only messages sent on this date. For example: `2009-07-06`. You can also specify an inequality, such as `DateSent<=YYYY-MM-DD`, to read messages sent on or before midnight on a date, and `DateSent>=YYYY-MM-DD` to read messages sent on or after midnight on a date. (optional)
    DateSentAfter := time.Now() // time.Time | The date of the messages to show. Specify a date as `YYYY-MM-DD` in GMT to read only messages sent on this date. For example: `2009-07-06`. You can also specify an inequality, such as `DateSent<=YYYY-MM-DD`, to read messages sent on or before midnight on a date, and `DateSent>=YYYY-MM-DD` to read messages sent on or after midnight on a date. (optional)
    PageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListMessage(context.Background(), AccountSid).To(To).From(From).DateSent(DateSent).DateSentBefore(DateSentBefore).DateSentAfter(DateSentAfter).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListMessage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListMessage`: ListMessageResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListMessage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Message resources to read.

### Other Parameters

Other parameters are passed through a pointer to a ListMessageParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **To** | **string** | Read messages sent to only this phone number.
 **From** | **string** | Read messages sent from only this phone number or alphanumeric sender ID.
 **DateSent** | **time.Time** | The date of the messages to show. Specify a date as &#x60;YYYY-MM-DD&#x60; in GMT to read only messages sent on this date. For example: &#x60;2009-07-06&#x60;. You can also specify an inequality, such as &#x60;DateSent&lt;&#x3D;YYYY-MM-DD&#x60;, to read messages sent on or before midnight on a date, and &#x60;DateSent&gt;&#x3D;YYYY-MM-DD&#x60; to read messages sent on or after midnight on a date.
 **DateSentBefore** | **time.Time** | The date of the messages to show. Specify a date as &#x60;YYYY-MM-DD&#x60; in GMT to read only messages sent on this date. For example: &#x60;2009-07-06&#x60;. You can also specify an inequality, such as &#x60;DateSent&lt;&#x3D;YYYY-MM-DD&#x60;, to read messages sent on or before midnight on a date, and &#x60;DateSent&gt;&#x3D;YYYY-MM-DD&#x60; to read messages sent on or after midnight on a date.
 **DateSentAfter** | **time.Time** | The date of the messages to show. Specify a date as &#x60;YYYY-MM-DD&#x60; in GMT to read only messages sent on this date. For example: &#x60;2009-07-06&#x60;. You can also specify an inequality, such as &#x60;DateSent&lt;&#x3D;YYYY-MM-DD&#x60;, to read messages sent on or before midnight on a date, and &#x60;DateSent&gt;&#x3D;YYYY-MM-DD&#x60; to read messages sent on or after midnight on a date.
 **PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListMessageResponse**](ListMessageResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListNotification

> ListNotificationResponse ListNotification(ctx, AccountSid).Log(Log).MessageDate(MessageDate).MessageDateBefore(MessageDateBefore).MessageDateAfter(MessageDateAfter).PageSize(PageSize).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    AccountSid := "AccountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Notification resources to read.
    Log := int32(56) // int32 | Only read notifications of the specified log level. Can be:  `0` to read only ERROR notifications or `1` to read only WARNING notifications. By default, all notifications are read. (optional)
    MessageDate := time.Now() // string | Only show notifications for the specified date, formatted as `YYYY-MM-DD`. You can also specify an inequality, such as `<=YYYY-MM-DD` for messages logged at or before midnight on a date, or `>=YYYY-MM-DD` for messages logged at or after midnight on a date. (optional)
    MessageDateBefore := time.Now() // string | Only show notifications for the specified date, formatted as `YYYY-MM-DD`. You can also specify an inequality, such as `<=YYYY-MM-DD` for messages logged at or before midnight on a date, or `>=YYYY-MM-DD` for messages logged at or after midnight on a date. (optional)
    MessageDateAfter := time.Now() // string | Only show notifications for the specified date, formatted as `YYYY-MM-DD`. You can also specify an inequality, such as `<=YYYY-MM-DD` for messages logged at or before midnight on a date, or `>=YYYY-MM-DD` for messages logged at or after midnight on a date. (optional)
    PageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListNotification(context.Background(), AccountSid).Log(Log).MessageDate(MessageDate).MessageDateBefore(MessageDateBefore).MessageDateAfter(MessageDateAfter).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListNotification``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListNotification`: ListNotificationResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListNotification`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Notification resources to read.

### Other Parameters

Other parameters are passed through a pointer to a ListNotificationParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **Log** | **int32** | Only read notifications of the specified log level. Can be:  &#x60;0&#x60; to read only ERROR notifications or &#x60;1&#x60; to read only WARNING notifications. By default, all notifications are read.
 **MessageDate** | **string** | Only show notifications for the specified date, formatted as &#x60;YYYY-MM-DD&#x60;. You can also specify an inequality, such as &#x60;&lt;&#x3D;YYYY-MM-DD&#x60; for messages logged at or before midnight on a date, or &#x60;&gt;&#x3D;YYYY-MM-DD&#x60; for messages logged at or after midnight on a date.
 **MessageDateBefore** | **string** | Only show notifications for the specified date, formatted as &#x60;YYYY-MM-DD&#x60;. You can also specify an inequality, such as &#x60;&lt;&#x3D;YYYY-MM-DD&#x60; for messages logged at or before midnight on a date, or &#x60;&gt;&#x3D;YYYY-MM-DD&#x60; for messages logged at or after midnight on a date.
 **MessageDateAfter** | **string** | Only show notifications for the specified date, formatted as &#x60;YYYY-MM-DD&#x60;. You can also specify an inequality, such as &#x60;&lt;&#x3D;YYYY-MM-DD&#x60; for messages logged at or before midnight on a date, or &#x60;&gt;&#x3D;YYYY-MM-DD&#x60; for messages logged at or after midnight on a date.
 **PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListNotificationResponse**](ListNotificationResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOutgoingCallerId

> ListOutgoingCallerIdResponse ListOutgoingCallerId(ctx, AccountSid).PhoneNumber(PhoneNumber).FriendlyName(FriendlyName).PageSize(PageSize).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AccountSid := "AccountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the OutgoingCallerId resources to read.
    PhoneNumber := "PhoneNumber_example" // string | The phone number of the OutgoingCallerId resources to read. (optional)
    FriendlyName := "FriendlyName_example" // string | The string that identifies the OutgoingCallerId resources to read. (optional)
    PageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListOutgoingCallerId(context.Background(), AccountSid).PhoneNumber(PhoneNumber).FriendlyName(FriendlyName).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListOutgoingCallerId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListOutgoingCallerId`: ListOutgoingCallerIdResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListOutgoingCallerId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the OutgoingCallerId resources to read.

### Other Parameters

Other parameters are passed through a pointer to a ListOutgoingCallerIdParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **PhoneNumber** | **string** | The phone number of the OutgoingCallerId resources to read.
 **FriendlyName** | **string** | The string that identifies the OutgoingCallerId resources to read.
 **PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListOutgoingCallerIdResponse**](ListOutgoingCallerIdResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListParticipant

> ListParticipantResponse ListParticipant(ctx, AccountSid, ConferenceSid).Muted(Muted).Hold(Hold).Coaching(Coaching).PageSize(PageSize).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AccountSid := "AccountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Participant resources to read.
    ConferenceSid := "ConferenceSid_example" // string | The SID of the conference with the participants to read.
    Muted := true // bool | Whether to return only participants that are muted. Can be: `true` or `false`. (optional)
    Hold := true // bool | Whether to return only participants that are on hold. Can be: `true` or `false`. (optional)
    Coaching := true // bool | Whether to return only participants who are coaching another call. Can be: `true` or `false`. (optional)
    PageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListParticipant(context.Background(), AccountSid, ConferenceSid).Muted(Muted).Hold(Hold).Coaching(Coaching).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListParticipant``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListParticipant`: ListParticipantResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListParticipant`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Participant resources to read.
**ConferenceSid** | **string** | The SID of the conference with the participants to read.

### Other Parameters

Other parameters are passed through a pointer to a ListParticipantParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


 **Muted** | **bool** | Whether to return only participants that are muted. Can be: &#x60;true&#x60; or &#x60;false&#x60;.
 **Hold** | **bool** | Whether to return only participants that are on hold. Can be: &#x60;true&#x60; or &#x60;false&#x60;.
 **Coaching** | **bool** | Whether to return only participants who are coaching another call. Can be: &#x60;true&#x60; or &#x60;false&#x60;.
 **PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListParticipantResponse**](ListParticipantResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListQueue

> ListQueueResponse ListQueue(ctx, AccountSid).PageSize(PageSize).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AccountSid := "AccountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Queue resources to read.
    PageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListQueue(context.Background(), AccountSid).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListQueue``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListQueue`: ListQueueResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListQueue`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Queue resources to read.

### Other Parameters

Other parameters are passed through a pointer to a ListQueueParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListQueueResponse**](ListQueueResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRecording

> ListRecordingResponse ListRecording(ctx, AccountSid).DateCreated(DateCreated).DateCreatedBefore(DateCreatedBefore).DateCreatedAfter(DateCreatedAfter).CallSid(CallSid).ConferenceSid(ConferenceSid).PageSize(PageSize).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    AccountSid := "AccountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Recording resources to read.
    DateCreated := time.Now() // time.Time | Only include recordings that were created on this date. Specify a date as `YYYY-MM-DD` in GMT, for example: `2009-07-06`, to read recordings that were created on this date. You can also specify an inequality, such as `DateCreated<=YYYY-MM-DD`, to read recordings that were created on or before midnight of this date, and `DateCreated>=YYYY-MM-DD` to read recordings that were created on or after midnight of this date. (optional)
    DateCreatedBefore := time.Now() // time.Time | Only include recordings that were created on this date. Specify a date as `YYYY-MM-DD` in GMT, for example: `2009-07-06`, to read recordings that were created on this date. You can also specify an inequality, such as `DateCreated<=YYYY-MM-DD`, to read recordings that were created on or before midnight of this date, and `DateCreated>=YYYY-MM-DD` to read recordings that were created on or after midnight of this date. (optional)
    DateCreatedAfter := time.Now() // time.Time | Only include recordings that were created on this date. Specify a date as `YYYY-MM-DD` in GMT, for example: `2009-07-06`, to read recordings that were created on this date. You can also specify an inequality, such as `DateCreated<=YYYY-MM-DD`, to read recordings that were created on or before midnight of this date, and `DateCreated>=YYYY-MM-DD` to read recordings that were created on or after midnight of this date. (optional)
    CallSid := "CallSid_example" // string | The [Call](https://www.twilio.com/docs/voice/api/call-resource) SID of the resources to read. (optional)
    ConferenceSid := "ConferenceSid_example" // string | The Conference SID that identifies the conference associated with the recording to read. (optional)
    PageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListRecording(context.Background(), AccountSid).DateCreated(DateCreated).DateCreatedBefore(DateCreatedBefore).DateCreatedAfter(DateCreatedAfter).CallSid(CallSid).ConferenceSid(ConferenceSid).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListRecording``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListRecording`: ListRecordingResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListRecording`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Recording resources to read.

### Other Parameters

Other parameters are passed through a pointer to a ListRecordingParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **DateCreated** | **time.Time** | Only include recordings that were created on this date. Specify a date as &#x60;YYYY-MM-DD&#x60; in GMT, for example: &#x60;2009-07-06&#x60;, to read recordings that were created on this date. You can also specify an inequality, such as &#x60;DateCreated&lt;&#x3D;YYYY-MM-DD&#x60;, to read recordings that were created on or before midnight of this date, and &#x60;DateCreated&gt;&#x3D;YYYY-MM-DD&#x60; to read recordings that were created on or after midnight of this date.
 **DateCreatedBefore** | **time.Time** | Only include recordings that were created on this date. Specify a date as &#x60;YYYY-MM-DD&#x60; in GMT, for example: &#x60;2009-07-06&#x60;, to read recordings that were created on this date. You can also specify an inequality, such as &#x60;DateCreated&lt;&#x3D;YYYY-MM-DD&#x60;, to read recordings that were created on or before midnight of this date, and &#x60;DateCreated&gt;&#x3D;YYYY-MM-DD&#x60; to read recordings that were created on or after midnight of this date.
 **DateCreatedAfter** | **time.Time** | Only include recordings that were created on this date. Specify a date as &#x60;YYYY-MM-DD&#x60; in GMT, for example: &#x60;2009-07-06&#x60;, to read recordings that were created on this date. You can also specify an inequality, such as &#x60;DateCreated&lt;&#x3D;YYYY-MM-DD&#x60;, to read recordings that were created on or before midnight of this date, and &#x60;DateCreated&gt;&#x3D;YYYY-MM-DD&#x60; to read recordings that were created on or after midnight of this date.
 **CallSid** | **string** | The [Call](https://www.twilio.com/docs/voice/api/call-resource) SID of the resources to read.
 **ConferenceSid** | **string** | The Conference SID that identifies the conference associated with the recording to read.
 **PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListRecordingResponse**](ListRecordingResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRecordingAddOnResult

> ListRecordingAddOnResultResponse ListRecordingAddOnResult(ctx, AccountSid, ReferenceSid).PageSize(PageSize).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AccountSid := "AccountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Recording AddOnResult resources to read.
    ReferenceSid := "ReferenceSid_example" // string | The SID of the recording to which the result to read belongs.
    PageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListRecordingAddOnResult(context.Background(), AccountSid, ReferenceSid).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListRecordingAddOnResult``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListRecordingAddOnResult`: ListRecordingAddOnResultResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListRecordingAddOnResult`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Recording AddOnResult resources to read.
**ReferenceSid** | **string** | The SID of the recording to which the result to read belongs.

### Other Parameters

Other parameters are passed through a pointer to a ListRecordingAddOnResultParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


 **PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListRecordingAddOnResultResponse**](ListRecordingAddOnResultResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRecordingAddOnResultPayload

> ListRecordingAddOnResultPayloadResponse ListRecordingAddOnResultPayload(ctx, AccountSid, ReferenceSid, AddOnResultSid).PageSize(PageSize).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AccountSid := "AccountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Recording AddOnResult Payload resources to read.
    ReferenceSid := "ReferenceSid_example" // string | The SID of the recording to which the AddOnResult resource that contains the payloads to read belongs.
    AddOnResultSid := "AddOnResultSid_example" // string | The SID of the AddOnResult to which the payloads to read belongs.
    PageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListRecordingAddOnResultPayload(context.Background(), AccountSid, ReferenceSid, AddOnResultSid).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListRecordingAddOnResultPayload``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListRecordingAddOnResultPayload`: ListRecordingAddOnResultPayloadResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListRecordingAddOnResultPayload`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Recording AddOnResult Payload resources to read.
**ReferenceSid** | **string** | The SID of the recording to which the AddOnResult resource that contains the payloads to read belongs.
**AddOnResultSid** | **string** | The SID of the AddOnResult to which the payloads to read belongs.

### Other Parameters

Other parameters are passed through a pointer to a ListRecordingAddOnResultPayloadParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------



 **PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListRecordingAddOnResultPayloadResponse**](ListRecordingAddOnResultPayloadResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRecordingTranscription

> ListRecordingTranscriptionResponse ListRecordingTranscription(ctx, AccountSid, RecordingSid).PageSize(PageSize).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AccountSid := "AccountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Transcription resources to read.
    RecordingSid := "RecordingSid_example" // string | The SID of the [Recording](https://www.twilio.com/docs/voice/api/recording) that created the transcriptions to read.
    PageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListRecordingTranscription(context.Background(), AccountSid, RecordingSid).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListRecordingTranscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListRecordingTranscription`: ListRecordingTranscriptionResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListRecordingTranscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Transcription resources to read.
**RecordingSid** | **string** | The SID of the [Recording](https://www.twilio.com/docs/voice/api/recording) that created the transcriptions to read.

### Other Parameters

Other parameters are passed through a pointer to a ListRecordingTranscriptionParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


 **PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListRecordingTranscriptionResponse**](ListRecordingTranscriptionResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListShortCode

> ListShortCodeResponse ListShortCode(ctx, AccountSid).FriendlyName(FriendlyName).ShortCode(ShortCode).PageSize(PageSize).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AccountSid := "AccountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the ShortCode resource(s) to read.
    FriendlyName := "FriendlyName_example" // string | The string that identifies the ShortCode resources to read. (optional)
    ShortCode := "ShortCode_example" // string | Only show the ShortCode resources that match this pattern. You can specify partial numbers and use '*' as a wildcard for any digit. (optional)
    PageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListShortCode(context.Background(), AccountSid).FriendlyName(FriendlyName).ShortCode(ShortCode).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListShortCode``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListShortCode`: ListShortCodeResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListShortCode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the ShortCode resource(s) to read.

### Other Parameters

Other parameters are passed through a pointer to a ListShortCodeParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **FriendlyName** | **string** | The string that identifies the ShortCode resources to read.
 **ShortCode** | **string** | Only show the ShortCode resources that match this pattern. You can specify partial numbers and use &#39;*&#39; as a wildcard for any digit.
 **PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListShortCodeResponse**](ListShortCodeResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSigningKey

> ListSigningKeyResponse ListSigningKey(ctx, AccountSid).PageSize(PageSize).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AccountSid := "AccountSid_example" // string | 
    PageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListSigningKey(context.Background(), AccountSid).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListSigningKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSigningKey`: ListSigningKeyResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListSigningKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a ListSigningKeyParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListSigningKeyResponse**](ListSigningKeyResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSipAuthCallsCredentialListMapping

> ListSipAuthCallsCredentialListMappingResponse ListSipAuthCallsCredentialListMapping(ctx, AccountSid, DomainSid).PageSize(PageSize).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AccountSid := "AccountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the CredentialListMapping resources to read.
    DomainSid := "DomainSid_example" // string | The SID of the SIP domain that contains the resources to read.
    PageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListSipAuthCallsCredentialListMapping(context.Background(), AccountSid, DomainSid).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListSipAuthCallsCredentialListMapping``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSipAuthCallsCredentialListMapping`: ListSipAuthCallsCredentialListMappingResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListSipAuthCallsCredentialListMapping`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the CredentialListMapping resources to read.
**DomainSid** | **string** | The SID of the SIP domain that contains the resources to read.

### Other Parameters

Other parameters are passed through a pointer to a ListSipAuthCallsCredentialListMappingParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


 **PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListSipAuthCallsCredentialListMappingResponse**](ListSipAuthCallsCredentialListMappingResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSipAuthCallsIpAccessControlListMapping

> ListSipAuthCallsIpAccessControlListMappingResponse ListSipAuthCallsIpAccessControlListMapping(ctx, AccountSid, DomainSid).PageSize(PageSize).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AccountSid := "AccountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the IpAccessControlListMapping resources to read.
    DomainSid := "DomainSid_example" // string | The SID of the SIP domain that contains the resources to read.
    PageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListSipAuthCallsIpAccessControlListMapping(context.Background(), AccountSid, DomainSid).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListSipAuthCallsIpAccessControlListMapping``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSipAuthCallsIpAccessControlListMapping`: ListSipAuthCallsIpAccessControlListMappingResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListSipAuthCallsIpAccessControlListMapping`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the IpAccessControlListMapping resources to read.
**DomainSid** | **string** | The SID of the SIP domain that contains the resources to read.

### Other Parameters

Other parameters are passed through a pointer to a ListSipAuthCallsIpAccessControlListMappingParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


 **PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListSipAuthCallsIpAccessControlListMappingResponse**](ListSipAuthCallsIpAccessControlListMappingResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSipAuthRegistrationsCredentialListMapping

> ListSipAuthRegistrationsCredentialListMappingResponse ListSipAuthRegistrationsCredentialListMapping(ctx, AccountSid, DomainSid).PageSize(PageSize).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AccountSid := "AccountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the CredentialListMapping resources to read.
    DomainSid := "DomainSid_example" // string | The SID of the SIP domain that contains the resources to read.
    PageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListSipAuthRegistrationsCredentialListMapping(context.Background(), AccountSid, DomainSid).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListSipAuthRegistrationsCredentialListMapping``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSipAuthRegistrationsCredentialListMapping`: ListSipAuthRegistrationsCredentialListMappingResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListSipAuthRegistrationsCredentialListMapping`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the CredentialListMapping resources to read.
**DomainSid** | **string** | The SID of the SIP domain that contains the resources to read.

### Other Parameters

Other parameters are passed through a pointer to a ListSipAuthRegistrationsCredentialListMappingParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


 **PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListSipAuthRegistrationsCredentialListMappingResponse**](ListSipAuthRegistrationsCredentialListMappingResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSipCredential

> ListSipCredentialResponse ListSipCredential(ctx, AccountSid, CredentialListSid).PageSize(PageSize).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AccountSid := "AccountSid_example" // string | The unique id of the Account that is responsible for this resource.
    CredentialListSid := "CredentialListSid_example" // string | The unique id that identifies the credential list that contains the desired credentials.
    PageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListSipCredential(context.Background(), AccountSid, CredentialListSid).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListSipCredential``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSipCredential`: ListSipCredentialResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListSipCredential`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string** | The unique id of the Account that is responsible for this resource.
**CredentialListSid** | **string** | The unique id that identifies the credential list that contains the desired credentials.

### Other Parameters

Other parameters are passed through a pointer to a ListSipCredentialParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


 **PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListSipCredentialResponse**](ListSipCredentialResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSipCredentialList

> ListSipCredentialListResponse ListSipCredentialList(ctx, AccountSid).PageSize(PageSize).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AccountSid := "AccountSid_example" // string | The unique id of the Account that is responsible for this resource.
    PageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListSipCredentialList(context.Background(), AccountSid).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListSipCredentialList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSipCredentialList`: ListSipCredentialListResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListSipCredentialList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string** | The unique id of the Account that is responsible for this resource.

### Other Parameters

Other parameters are passed through a pointer to a ListSipCredentialListParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListSipCredentialListResponse**](ListSipCredentialListResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSipCredentialListMapping

> ListSipCredentialListMappingResponse ListSipCredentialListMapping(ctx, AccountSid, DomainSid).PageSize(PageSize).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AccountSid := "AccountSid_example" // string | The unique id of the [Account](https://www.twilio.com/docs/iam/api/account) responsible for this resource.
    DomainSid := "DomainSid_example" // string | A 34 character string that uniquely identifies the SIP Domain that includes the resource to read.
    PageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListSipCredentialListMapping(context.Background(), AccountSid, DomainSid).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListSipCredentialListMapping``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSipCredentialListMapping`: ListSipCredentialListMappingResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListSipCredentialListMapping`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string** | The unique id of the [Account](https://www.twilio.com/docs/iam/api/account) responsible for this resource.
**DomainSid** | **string** | A 34 character string that uniquely identifies the SIP Domain that includes the resource to read.

### Other Parameters

Other parameters are passed through a pointer to a ListSipCredentialListMappingParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


 **PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListSipCredentialListMappingResponse**](ListSipCredentialListMappingResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSipDomain

> ListSipDomainResponse ListSipDomain(ctx, AccountSid).PageSize(PageSize).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AccountSid := "AccountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the SipDomain resources to read.
    PageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListSipDomain(context.Background(), AccountSid).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListSipDomain``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSipDomain`: ListSipDomainResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListSipDomain`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the SipDomain resources to read.

### Other Parameters

Other parameters are passed through a pointer to a ListSipDomainParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListSipDomainResponse**](ListSipDomainResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSipIpAccessControlList

> ListSipIpAccessControlListResponse ListSipIpAccessControlList(ctx, AccountSid).PageSize(PageSize).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AccountSid := "AccountSid_example" // string | The unique id of the [Account](https://www.twilio.com/docs/iam/api/account) responsible for this resource.
    PageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListSipIpAccessControlList(context.Background(), AccountSid).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListSipIpAccessControlList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSipIpAccessControlList`: ListSipIpAccessControlListResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListSipIpAccessControlList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string** | The unique id of the [Account](https://www.twilio.com/docs/iam/api/account) responsible for this resource.

### Other Parameters

Other parameters are passed through a pointer to a ListSipIpAccessControlListParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListSipIpAccessControlListResponse**](ListSipIpAccessControlListResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSipIpAccessControlListMapping

> ListSipIpAccessControlListMappingResponse ListSipIpAccessControlListMapping(ctx, AccountSid, DomainSid).PageSize(PageSize).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AccountSid := "AccountSid_example" // string | The unique id of the Account that is responsible for this resource.
    DomainSid := "DomainSid_example" // string | A 34 character string that uniquely identifies the SIP domain.
    PageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListSipIpAccessControlListMapping(context.Background(), AccountSid, DomainSid).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListSipIpAccessControlListMapping``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSipIpAccessControlListMapping`: ListSipIpAccessControlListMappingResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListSipIpAccessControlListMapping`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string** | The unique id of the Account that is responsible for this resource.
**DomainSid** | **string** | A 34 character string that uniquely identifies the SIP domain.

### Other Parameters

Other parameters are passed through a pointer to a ListSipIpAccessControlListMappingParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


 **PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListSipIpAccessControlListMappingResponse**](ListSipIpAccessControlListMappingResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSipIpAddress

> ListSipIpAddressResponse ListSipIpAddress(ctx, AccountSid, IpAccessControlListSid).PageSize(PageSize).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AccountSid := "AccountSid_example" // string | The unique id of the [Account](https://www.twilio.com/docs/iam/api/account) responsible for this resource.
    IpAccessControlListSid := "IpAccessControlListSid_example" // string | The IpAccessControlList Sid that identifies the IpAddress resources to read.
    PageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListSipIpAddress(context.Background(), AccountSid, IpAccessControlListSid).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListSipIpAddress``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSipIpAddress`: ListSipIpAddressResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListSipIpAddress`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string** | The unique id of the [Account](https://www.twilio.com/docs/iam/api/account) responsible for this resource.
**IpAccessControlListSid** | **string** | The IpAccessControlList Sid that identifies the IpAddress resources to read.

### Other Parameters

Other parameters are passed through a pointer to a ListSipIpAddressParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


 **PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListSipIpAddressResponse**](ListSipIpAddressResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTranscription

> ListTranscriptionResponse ListTranscription(ctx, AccountSid).PageSize(PageSize).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AccountSid := "AccountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Transcription resources to read.
    PageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListTranscription(context.Background(), AccountSid).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListTranscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListTranscription`: ListTranscriptionResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListTranscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Transcription resources to read.

### Other Parameters

Other parameters are passed through a pointer to a ListTranscriptionParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListTranscriptionResponse**](ListTranscriptionResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUsageRecord

> ListUsageRecordResponse ListUsageRecord(ctx, AccountSid).Category(Category).StartDate(StartDate).EndDate(EndDate).IncludeSubaccounts(IncludeSubaccounts).PageSize(PageSize).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    AccountSid := "AccountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the UsageRecord resources to read.
    Category := "Category_example" // string | The [usage category](https://www.twilio.com/docs/usage/api/usage-record#usage-categories) of the UsageRecord resources to read. Only UsageRecord resources in the specified category are retrieved. (optional)
    StartDate := time.Now() // string | Only include usage that has occurred on or after this date. Specify the date in GMT and format as `YYYY-MM-DD`. You can also specify offsets from the current date, such as: `-30days`, which will set the start date to be 30 days before the current date. (optional)
    EndDate := time.Now() // string | Only include usage that occurred on or before this date. Specify the date in GMT and format as `YYYY-MM-DD`.  You can also specify offsets from the current date, such as: `+30days`, which will set the end date to 30 days from the current date. (optional)
    IncludeSubaccounts := true // bool | Whether to include usage from the master account and all its subaccounts. Can be: `true` (the default) to include usage from the master account and all subaccounts or `false` to retrieve usage from only the specified account. (optional)
    PageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListUsageRecord(context.Background(), AccountSid).Category(Category).StartDate(StartDate).EndDate(EndDate).IncludeSubaccounts(IncludeSubaccounts).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListUsageRecord``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListUsageRecord`: ListUsageRecordResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListUsageRecord`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the UsageRecord resources to read.

### Other Parameters

Other parameters are passed through a pointer to a ListUsageRecordParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **Category** | **string** | The [usage category](https://www.twilio.com/docs/usage/api/usage-record#usage-categories) of the UsageRecord resources to read. Only UsageRecord resources in the specified category are retrieved.
 **StartDate** | **string** | Only include usage that has occurred on or after this date. Specify the date in GMT and format as &#x60;YYYY-MM-DD&#x60;. You can also specify offsets from the current date, such as: &#x60;-30days&#x60;, which will set the start date to be 30 days before the current date.
 **EndDate** | **string** | Only include usage that occurred on or before this date. Specify the date in GMT and format as &#x60;YYYY-MM-DD&#x60;.  You can also specify offsets from the current date, such as: &#x60;+30days&#x60;, which will set the end date to 30 days from the current date.
 **IncludeSubaccounts** | **bool** | Whether to include usage from the master account and all its subaccounts. Can be: &#x60;true&#x60; (the default) to include usage from the master account and all subaccounts or &#x60;false&#x60; to retrieve usage from only the specified account.
 **PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListUsageRecordResponse**](ListUsageRecordResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUsageRecordAllTime

> ListUsageRecordAllTimeResponse ListUsageRecordAllTime(ctx, AccountSid).Category(Category).StartDate(StartDate).EndDate(EndDate).IncludeSubaccounts(IncludeSubaccounts).PageSize(PageSize).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    AccountSid := "AccountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the UsageRecord resources to read.
    Category := "Category_example" // string | The [usage category](https://www.twilio.com/docs/usage/api/usage-record#usage-categories) of the UsageRecord resources to read. Only UsageRecord resources in the specified category are retrieved. (optional)
    StartDate := time.Now() // string | Only include usage that has occurred on or after this date. Specify the date in GMT and format as `YYYY-MM-DD`. You can also specify offsets from the current date, such as: `-30days`, which will set the start date to be 30 days before the current date. (optional)
    EndDate := time.Now() // string | Only include usage that occurred on or before this date. Specify the date in GMT and format as `YYYY-MM-DD`.  You can also specify offsets from the current date, such as: `+30days`, which will set the end date to 30 days from the current date. (optional)
    IncludeSubaccounts := true // bool | Whether to include usage from the master account and all its subaccounts. Can be: `true` (the default) to include usage from the master account and all subaccounts or `false` to retrieve usage from only the specified account. (optional)
    PageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListUsageRecordAllTime(context.Background(), AccountSid).Category(Category).StartDate(StartDate).EndDate(EndDate).IncludeSubaccounts(IncludeSubaccounts).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListUsageRecordAllTime``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListUsageRecordAllTime`: ListUsageRecordAllTimeResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListUsageRecordAllTime`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the UsageRecord resources to read.

### Other Parameters

Other parameters are passed through a pointer to a ListUsageRecordAllTimeParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **Category** | **string** | The [usage category](https://www.twilio.com/docs/usage/api/usage-record#usage-categories) of the UsageRecord resources to read. Only UsageRecord resources in the specified category are retrieved.
 **StartDate** | **string** | Only include usage that has occurred on or after this date. Specify the date in GMT and format as &#x60;YYYY-MM-DD&#x60;. You can also specify offsets from the current date, such as: &#x60;-30days&#x60;, which will set the start date to be 30 days before the current date.
 **EndDate** | **string** | Only include usage that occurred on or before this date. Specify the date in GMT and format as &#x60;YYYY-MM-DD&#x60;.  You can also specify offsets from the current date, such as: &#x60;+30days&#x60;, which will set the end date to 30 days from the current date.
 **IncludeSubaccounts** | **bool** | Whether to include usage from the master account and all its subaccounts. Can be: &#x60;true&#x60; (the default) to include usage from the master account and all subaccounts or &#x60;false&#x60; to retrieve usage from only the specified account.
 **PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListUsageRecordAllTimeResponse**](ListUsageRecordAllTimeResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUsageRecordDaily

> ListUsageRecordDailyResponse ListUsageRecordDaily(ctx, AccountSid).Category(Category).StartDate(StartDate).EndDate(EndDate).IncludeSubaccounts(IncludeSubaccounts).PageSize(PageSize).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    AccountSid := "AccountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the UsageRecord resources to read.
    Category := "Category_example" // string | The [usage category](https://www.twilio.com/docs/usage/api/usage-record#usage-categories) of the UsageRecord resources to read. Only UsageRecord resources in the specified category are retrieved. (optional)
    StartDate := time.Now() // string | Only include usage that has occurred on or after this date. Specify the date in GMT and format as `YYYY-MM-DD`. You can also specify offsets from the current date, such as: `-30days`, which will set the start date to be 30 days before the current date. (optional)
    EndDate := time.Now() // string | Only include usage that occurred on or before this date. Specify the date in GMT and format as `YYYY-MM-DD`.  You can also specify offsets from the current date, such as: `+30days`, which will set the end date to 30 days from the current date. (optional)
    IncludeSubaccounts := true // bool | Whether to include usage from the master account and all its subaccounts. Can be: `true` (the default) to include usage from the master account and all subaccounts or `false` to retrieve usage from only the specified account. (optional)
    PageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListUsageRecordDaily(context.Background(), AccountSid).Category(Category).StartDate(StartDate).EndDate(EndDate).IncludeSubaccounts(IncludeSubaccounts).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListUsageRecordDaily``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListUsageRecordDaily`: ListUsageRecordDailyResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListUsageRecordDaily`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the UsageRecord resources to read.

### Other Parameters

Other parameters are passed through a pointer to a ListUsageRecordDailyParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **Category** | **string** | The [usage category](https://www.twilio.com/docs/usage/api/usage-record#usage-categories) of the UsageRecord resources to read. Only UsageRecord resources in the specified category are retrieved.
 **StartDate** | **string** | Only include usage that has occurred on or after this date. Specify the date in GMT and format as &#x60;YYYY-MM-DD&#x60;. You can also specify offsets from the current date, such as: &#x60;-30days&#x60;, which will set the start date to be 30 days before the current date.
 **EndDate** | **string** | Only include usage that occurred on or before this date. Specify the date in GMT and format as &#x60;YYYY-MM-DD&#x60;.  You can also specify offsets from the current date, such as: &#x60;+30days&#x60;, which will set the end date to 30 days from the current date.
 **IncludeSubaccounts** | **bool** | Whether to include usage from the master account and all its subaccounts. Can be: &#x60;true&#x60; (the default) to include usage from the master account and all subaccounts or &#x60;false&#x60; to retrieve usage from only the specified account.
 **PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListUsageRecordDailyResponse**](ListUsageRecordDailyResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUsageRecordLastMonth

> ListUsageRecordLastMonthResponse ListUsageRecordLastMonth(ctx, AccountSid).Category(Category).StartDate(StartDate).EndDate(EndDate).IncludeSubaccounts(IncludeSubaccounts).PageSize(PageSize).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    AccountSid := "AccountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the UsageRecord resources to read.
    Category := "Category_example" // string | The [usage category](https://www.twilio.com/docs/usage/api/usage-record#usage-categories) of the UsageRecord resources to read. Only UsageRecord resources in the specified category are retrieved. (optional)
    StartDate := time.Now() // string | Only include usage that has occurred on or after this date. Specify the date in GMT and format as `YYYY-MM-DD`. You can also specify offsets from the current date, such as: `-30days`, which will set the start date to be 30 days before the current date. (optional)
    EndDate := time.Now() // string | Only include usage that occurred on or before this date. Specify the date in GMT and format as `YYYY-MM-DD`.  You can also specify offsets from the current date, such as: `+30days`, which will set the end date to 30 days from the current date. (optional)
    IncludeSubaccounts := true // bool | Whether to include usage from the master account and all its subaccounts. Can be: `true` (the default) to include usage from the master account and all subaccounts or `false` to retrieve usage from only the specified account. (optional)
    PageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListUsageRecordLastMonth(context.Background(), AccountSid).Category(Category).StartDate(StartDate).EndDate(EndDate).IncludeSubaccounts(IncludeSubaccounts).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListUsageRecordLastMonth``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListUsageRecordLastMonth`: ListUsageRecordLastMonthResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListUsageRecordLastMonth`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the UsageRecord resources to read.

### Other Parameters

Other parameters are passed through a pointer to a ListUsageRecordLastMonthParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **Category** | **string** | The [usage category](https://www.twilio.com/docs/usage/api/usage-record#usage-categories) of the UsageRecord resources to read. Only UsageRecord resources in the specified category are retrieved.
 **StartDate** | **string** | Only include usage that has occurred on or after this date. Specify the date in GMT and format as &#x60;YYYY-MM-DD&#x60;. You can also specify offsets from the current date, such as: &#x60;-30days&#x60;, which will set the start date to be 30 days before the current date.
 **EndDate** | **string** | Only include usage that occurred on or before this date. Specify the date in GMT and format as &#x60;YYYY-MM-DD&#x60;.  You can also specify offsets from the current date, such as: &#x60;+30days&#x60;, which will set the end date to 30 days from the current date.
 **IncludeSubaccounts** | **bool** | Whether to include usage from the master account and all its subaccounts. Can be: &#x60;true&#x60; (the default) to include usage from the master account and all subaccounts or &#x60;false&#x60; to retrieve usage from only the specified account.
 **PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListUsageRecordLastMonthResponse**](ListUsageRecordLastMonthResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUsageRecordMonthly

> ListUsageRecordMonthlyResponse ListUsageRecordMonthly(ctx, AccountSid).Category(Category).StartDate(StartDate).EndDate(EndDate).IncludeSubaccounts(IncludeSubaccounts).PageSize(PageSize).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    AccountSid := "AccountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the UsageRecord resources to read.
    Category := "Category_example" // string | The [usage category](https://www.twilio.com/docs/usage/api/usage-record#usage-categories) of the UsageRecord resources to read. Only UsageRecord resources in the specified category are retrieved. (optional)
    StartDate := time.Now() // string | Only include usage that has occurred on or after this date. Specify the date in GMT and format as `YYYY-MM-DD`. You can also specify offsets from the current date, such as: `-30days`, which will set the start date to be 30 days before the current date. (optional)
    EndDate := time.Now() // string | Only include usage that occurred on or before this date. Specify the date in GMT and format as `YYYY-MM-DD`.  You can also specify offsets from the current date, such as: `+30days`, which will set the end date to 30 days from the current date. (optional)
    IncludeSubaccounts := true // bool | Whether to include usage from the master account and all its subaccounts. Can be: `true` (the default) to include usage from the master account and all subaccounts or `false` to retrieve usage from only the specified account. (optional)
    PageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListUsageRecordMonthly(context.Background(), AccountSid).Category(Category).StartDate(StartDate).EndDate(EndDate).IncludeSubaccounts(IncludeSubaccounts).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListUsageRecordMonthly``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListUsageRecordMonthly`: ListUsageRecordMonthlyResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListUsageRecordMonthly`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the UsageRecord resources to read.

### Other Parameters

Other parameters are passed through a pointer to a ListUsageRecordMonthlyParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **Category** | **string** | The [usage category](https://www.twilio.com/docs/usage/api/usage-record#usage-categories) of the UsageRecord resources to read. Only UsageRecord resources in the specified category are retrieved.
 **StartDate** | **string** | Only include usage that has occurred on or after this date. Specify the date in GMT and format as &#x60;YYYY-MM-DD&#x60;. You can also specify offsets from the current date, such as: &#x60;-30days&#x60;, which will set the start date to be 30 days before the current date.
 **EndDate** | **string** | Only include usage that occurred on or before this date. Specify the date in GMT and format as &#x60;YYYY-MM-DD&#x60;.  You can also specify offsets from the current date, such as: &#x60;+30days&#x60;, which will set the end date to 30 days from the current date.
 **IncludeSubaccounts** | **bool** | Whether to include usage from the master account and all its subaccounts. Can be: &#x60;true&#x60; (the default) to include usage from the master account and all subaccounts or &#x60;false&#x60; to retrieve usage from only the specified account.
 **PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListUsageRecordMonthlyResponse**](ListUsageRecordMonthlyResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUsageRecordThisMonth

> ListUsageRecordThisMonthResponse ListUsageRecordThisMonth(ctx, AccountSid).Category(Category).StartDate(StartDate).EndDate(EndDate).IncludeSubaccounts(IncludeSubaccounts).PageSize(PageSize).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    AccountSid := "AccountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the UsageRecord resources to read.
    Category := "Category_example" // string | The [usage category](https://www.twilio.com/docs/usage/api/usage-record#usage-categories) of the UsageRecord resources to read. Only UsageRecord resources in the specified category are retrieved. (optional)
    StartDate := time.Now() // string | Only include usage that has occurred on or after this date. Specify the date in GMT and format as `YYYY-MM-DD`. You can also specify offsets from the current date, such as: `-30days`, which will set the start date to be 30 days before the current date. (optional)
    EndDate := time.Now() // string | Only include usage that occurred on or before this date. Specify the date in GMT and format as `YYYY-MM-DD`.  You can also specify offsets from the current date, such as: `+30days`, which will set the end date to 30 days from the current date. (optional)
    IncludeSubaccounts := true // bool | Whether to include usage from the master account and all its subaccounts. Can be: `true` (the default) to include usage from the master account and all subaccounts or `false` to retrieve usage from only the specified account. (optional)
    PageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListUsageRecordThisMonth(context.Background(), AccountSid).Category(Category).StartDate(StartDate).EndDate(EndDate).IncludeSubaccounts(IncludeSubaccounts).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListUsageRecordThisMonth``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListUsageRecordThisMonth`: ListUsageRecordThisMonthResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListUsageRecordThisMonth`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the UsageRecord resources to read.

### Other Parameters

Other parameters are passed through a pointer to a ListUsageRecordThisMonthParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **Category** | **string** | The [usage category](https://www.twilio.com/docs/usage/api/usage-record#usage-categories) of the UsageRecord resources to read. Only UsageRecord resources in the specified category are retrieved.
 **StartDate** | **string** | Only include usage that has occurred on or after this date. Specify the date in GMT and format as &#x60;YYYY-MM-DD&#x60;. You can also specify offsets from the current date, such as: &#x60;-30days&#x60;, which will set the start date to be 30 days before the current date.
 **EndDate** | **string** | Only include usage that occurred on or before this date. Specify the date in GMT and format as &#x60;YYYY-MM-DD&#x60;.  You can also specify offsets from the current date, such as: &#x60;+30days&#x60;, which will set the end date to 30 days from the current date.
 **IncludeSubaccounts** | **bool** | Whether to include usage from the master account and all its subaccounts. Can be: &#x60;true&#x60; (the default) to include usage from the master account and all subaccounts or &#x60;false&#x60; to retrieve usage from only the specified account.
 **PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListUsageRecordThisMonthResponse**](ListUsageRecordThisMonthResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUsageRecordToday

> ListUsageRecordTodayResponse ListUsageRecordToday(ctx, AccountSid).Category(Category).StartDate(StartDate).EndDate(EndDate).IncludeSubaccounts(IncludeSubaccounts).PageSize(PageSize).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    AccountSid := "AccountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the UsageRecord resources to read.
    Category := "Category_example" // string | The [usage category](https://www.twilio.com/docs/usage/api/usage-record#usage-categories) of the UsageRecord resources to read. Only UsageRecord resources in the specified category are retrieved. (optional)
    StartDate := time.Now() // string | Only include usage that has occurred on or after this date. Specify the date in GMT and format as `YYYY-MM-DD`. You can also specify offsets from the current date, such as: `-30days`, which will set the start date to be 30 days before the current date. (optional)
    EndDate := time.Now() // string | Only include usage that occurred on or before this date. Specify the date in GMT and format as `YYYY-MM-DD`.  You can also specify offsets from the current date, such as: `+30days`, which will set the end date to 30 days from the current date. (optional)
    IncludeSubaccounts := true // bool | Whether to include usage from the master account and all its subaccounts. Can be: `true` (the default) to include usage from the master account and all subaccounts or `false` to retrieve usage from only the specified account. (optional)
    PageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListUsageRecordToday(context.Background(), AccountSid).Category(Category).StartDate(StartDate).EndDate(EndDate).IncludeSubaccounts(IncludeSubaccounts).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListUsageRecordToday``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListUsageRecordToday`: ListUsageRecordTodayResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListUsageRecordToday`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the UsageRecord resources to read.

### Other Parameters

Other parameters are passed through a pointer to a ListUsageRecordTodayParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **Category** | **string** | The [usage category](https://www.twilio.com/docs/usage/api/usage-record#usage-categories) of the UsageRecord resources to read. Only UsageRecord resources in the specified category are retrieved.
 **StartDate** | **string** | Only include usage that has occurred on or after this date. Specify the date in GMT and format as &#x60;YYYY-MM-DD&#x60;. You can also specify offsets from the current date, such as: &#x60;-30days&#x60;, which will set the start date to be 30 days before the current date.
 **EndDate** | **string** | Only include usage that occurred on or before this date. Specify the date in GMT and format as &#x60;YYYY-MM-DD&#x60;.  You can also specify offsets from the current date, such as: &#x60;+30days&#x60;, which will set the end date to 30 days from the current date.
 **IncludeSubaccounts** | **bool** | Whether to include usage from the master account and all its subaccounts. Can be: &#x60;true&#x60; (the default) to include usage from the master account and all subaccounts or &#x60;false&#x60; to retrieve usage from only the specified account.
 **PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListUsageRecordTodayResponse**](ListUsageRecordTodayResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUsageRecordYearly

> ListUsageRecordYearlyResponse ListUsageRecordYearly(ctx, AccountSid).Category(Category).StartDate(StartDate).EndDate(EndDate).IncludeSubaccounts(IncludeSubaccounts).PageSize(PageSize).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    AccountSid := "AccountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the UsageRecord resources to read.
    Category := "Category_example" // string | The [usage category](https://www.twilio.com/docs/usage/api/usage-record#usage-categories) of the UsageRecord resources to read. Only UsageRecord resources in the specified category are retrieved. (optional)
    StartDate := time.Now() // string | Only include usage that has occurred on or after this date. Specify the date in GMT and format as `YYYY-MM-DD`. You can also specify offsets from the current date, such as: `-30days`, which will set the start date to be 30 days before the current date. (optional)
    EndDate := time.Now() // string | Only include usage that occurred on or before this date. Specify the date in GMT and format as `YYYY-MM-DD`.  You can also specify offsets from the current date, such as: `+30days`, which will set the end date to 30 days from the current date. (optional)
    IncludeSubaccounts := true // bool | Whether to include usage from the master account and all its subaccounts. Can be: `true` (the default) to include usage from the master account and all subaccounts or `false` to retrieve usage from only the specified account. (optional)
    PageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListUsageRecordYearly(context.Background(), AccountSid).Category(Category).StartDate(StartDate).EndDate(EndDate).IncludeSubaccounts(IncludeSubaccounts).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListUsageRecordYearly``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListUsageRecordYearly`: ListUsageRecordYearlyResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListUsageRecordYearly`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the UsageRecord resources to read.

### Other Parameters

Other parameters are passed through a pointer to a ListUsageRecordYearlyParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **Category** | **string** | The [usage category](https://www.twilio.com/docs/usage/api/usage-record#usage-categories) of the UsageRecord resources to read. Only UsageRecord resources in the specified category are retrieved.
 **StartDate** | **string** | Only include usage that has occurred on or after this date. Specify the date in GMT and format as &#x60;YYYY-MM-DD&#x60;. You can also specify offsets from the current date, such as: &#x60;-30days&#x60;, which will set the start date to be 30 days before the current date.
 **EndDate** | **string** | Only include usage that occurred on or before this date. Specify the date in GMT and format as &#x60;YYYY-MM-DD&#x60;.  You can also specify offsets from the current date, such as: &#x60;+30days&#x60;, which will set the end date to 30 days from the current date.
 **IncludeSubaccounts** | **bool** | Whether to include usage from the master account and all its subaccounts. Can be: &#x60;true&#x60; (the default) to include usage from the master account and all subaccounts or &#x60;false&#x60; to retrieve usage from only the specified account.
 **PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListUsageRecordYearlyResponse**](ListUsageRecordYearlyResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUsageRecordYesterday

> ListUsageRecordYesterdayResponse ListUsageRecordYesterday(ctx, AccountSid).Category(Category).StartDate(StartDate).EndDate(EndDate).IncludeSubaccounts(IncludeSubaccounts).PageSize(PageSize).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    AccountSid := "AccountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the UsageRecord resources to read.
    Category := "Category_example" // string | The [usage category](https://www.twilio.com/docs/usage/api/usage-record#usage-categories) of the UsageRecord resources to read. Only UsageRecord resources in the specified category are retrieved. (optional)
    StartDate := time.Now() // string | Only include usage that has occurred on or after this date. Specify the date in GMT and format as `YYYY-MM-DD`. You can also specify offsets from the current date, such as: `-30days`, which will set the start date to be 30 days before the current date. (optional)
    EndDate := time.Now() // string | Only include usage that occurred on or before this date. Specify the date in GMT and format as `YYYY-MM-DD`.  You can also specify offsets from the current date, such as: `+30days`, which will set the end date to 30 days from the current date. (optional)
    IncludeSubaccounts := true // bool | Whether to include usage from the master account and all its subaccounts. Can be: `true` (the default) to include usage from the master account and all subaccounts or `false` to retrieve usage from only the specified account. (optional)
    PageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListUsageRecordYesterday(context.Background(), AccountSid).Category(Category).StartDate(StartDate).EndDate(EndDate).IncludeSubaccounts(IncludeSubaccounts).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListUsageRecordYesterday``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListUsageRecordYesterday`: ListUsageRecordYesterdayResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListUsageRecordYesterday`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the UsageRecord resources to read.

### Other Parameters

Other parameters are passed through a pointer to a ListUsageRecordYesterdayParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **Category** | **string** | The [usage category](https://www.twilio.com/docs/usage/api/usage-record#usage-categories) of the UsageRecord resources to read. Only UsageRecord resources in the specified category are retrieved.
 **StartDate** | **string** | Only include usage that has occurred on or after this date. Specify the date in GMT and format as &#x60;YYYY-MM-DD&#x60;. You can also specify offsets from the current date, such as: &#x60;-30days&#x60;, which will set the start date to be 30 days before the current date.
 **EndDate** | **string** | Only include usage that occurred on or before this date. Specify the date in GMT and format as &#x60;YYYY-MM-DD&#x60;.  You can also specify offsets from the current date, such as: &#x60;+30days&#x60;, which will set the end date to 30 days from the current date.
 **IncludeSubaccounts** | **bool** | Whether to include usage from the master account and all its subaccounts. Can be: &#x60;true&#x60; (the default) to include usage from the master account and all subaccounts or &#x60;false&#x60; to retrieve usage from only the specified account.
 **PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListUsageRecordYesterdayResponse**](ListUsageRecordYesterdayResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUsageTrigger

> ListUsageTriggerResponse ListUsageTrigger(ctx, AccountSid).Recurring(Recurring).TriggerBy(TriggerBy).UsageCategory(UsageCategory).PageSize(PageSize).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AccountSid := "AccountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the UsageTrigger resources to read.
    Recurring := "Recurring_example" // string | The frequency of recurring UsageTriggers to read. Can be: `daily`, `monthly`, or `yearly` to read recurring UsageTriggers. An empty value or a value of `alltime` reads non-recurring UsageTriggers. (optional)
    TriggerBy := "TriggerBy_example" // string | The trigger field of the UsageTriggers to read.  Can be: `count`, `usage`, or `price` as described in the [UsageRecords documentation](https://www.twilio.com/docs/usage/api/usage-record#usage-count-price). (optional)
    UsageCategory := "UsageCategory_example" // string | The usage category of the UsageTriggers to read. Must be a supported [usage categories](https://www.twilio.com/docs/usage/api/usage-record#usage-categories). (optional)
    PageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListUsageTrigger(context.Background(), AccountSid).Recurring(Recurring).TriggerBy(TriggerBy).UsageCategory(UsageCategory).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListUsageTrigger``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListUsageTrigger`: ListUsageTriggerResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListUsageTrigger`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the UsageTrigger resources to read.

### Other Parameters

Other parameters are passed through a pointer to a ListUsageTriggerParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **Recurring** | **string** | The frequency of recurring UsageTriggers to read. Can be: &#x60;daily&#x60;, &#x60;monthly&#x60;, or &#x60;yearly&#x60; to read recurring UsageTriggers. An empty value or a value of &#x60;alltime&#x60; reads non-recurring UsageTriggers.
 **TriggerBy** | **string** | The trigger field of the UsageTriggers to read.  Can be: &#x60;count&#x60;, &#x60;usage&#x60;, or &#x60;price&#x60; as described in the [UsageRecords documentation](https://www.twilio.com/docs/usage/api/usage-record#usage-count-price).
 **UsageCategory** | **string** | The usage category of the UsageTriggers to read. Must be a supported [usage categories](https://www.twilio.com/docs/usage/api/usage-record#usage-categories).
 **PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListUsageTriggerResponse**](ListUsageTriggerResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAccount

> ApiV2010Account UpdateAccount(ctx, Sid).FriendlyName(FriendlyName).Status(Status).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    Sid := "Sid_example" // string | The Account Sid that uniquely identifies the account to update
    FriendlyName := "FriendlyName_example" // string | Update the human-readable description of this Account (optional)
    Status := "Status_example" // string | Alter the status of this account: use `closed` to irreversibly close this account, `suspended` to temporarily suspend it, or `active` to reactivate it. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateAccount(context.Background(), Sid).FriendlyName(FriendlyName).Status(Status).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAccount`: ApiV2010Account
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The Account Sid that uniquely identifies the account to update

### Other Parameters

Other parameters are passed through a pointer to a UpdateAccountParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **FriendlyName** | **string** | Update the human-readable description of this Account
 **Status** | **string** | Alter the status of this account: use &#x60;closed&#x60; to irreversibly close this account, &#x60;suspended&#x60; to temporarily suspend it, or &#x60;active&#x60; to reactivate it.

### Return type

[**ApiV2010Account**](ApiV2010Account.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAddress

> ApiV2010AccountAddress UpdateAddress(ctx, AccountSid, Sid).AutoCorrectAddress(AutoCorrectAddress).City(City).CustomerName(CustomerName).EmergencyEnabled(EmergencyEnabled).FriendlyName(FriendlyName).PostalCode(PostalCode).Region(Region).Street(Street).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AccountSid := "AccountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that is responsible for the Address resource to update.
    Sid := "Sid_example" // string | The Twilio-provided string that uniquely identifies the Address resource to update.
    AutoCorrectAddress := true // bool | Whether we should automatically correct the address. Can be: `true` or `false` and the default is `true`. If empty or `true`, we will correct the address you provide if necessary. If `false`, we won't alter the address you provide. (optional)
    City := "City_example" // string | The city of the address. (optional)
    CustomerName := "CustomerName_example" // string | The name to associate with the address. (optional)
    EmergencyEnabled := true // bool | Whether to enable emergency calling on the address. Can be: `true` or `false`. (optional)
    FriendlyName := "FriendlyName_example" // string | A descriptive string that you create to describe the address. It can be up to 64 characters long. (optional)
    PostalCode := "PostalCode_example" // string | The postal code of the address. (optional)
    Region := "Region_example" // string | The state or region of the address. (optional)
    Street := "Street_example" // string | The number and street address of the address. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateAddress(context.Background(), AccountSid, Sid).AutoCorrectAddress(AutoCorrectAddress).City(City).CustomerName(CustomerName).EmergencyEnabled(EmergencyEnabled).FriendlyName(FriendlyName).PostalCode(PostalCode).Region(Region).Street(Street).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateAddress``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAddress`: ApiV2010AccountAddress
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateAddress`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that is responsible for the Address resource to update.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the Address resource to update.

### Other Parameters

Other parameters are passed through a pointer to a UpdateAddressParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


 **AutoCorrectAddress** | **bool** | Whether we should automatically correct the address. Can be: &#x60;true&#x60; or &#x60;false&#x60; and the default is &#x60;true&#x60;. If empty or &#x60;true&#x60;, we will correct the address you provide if necessary. If &#x60;false&#x60;, we won&#39;t alter the address you provide.
 **City** | **string** | The city of the address.
 **CustomerName** | **string** | The name to associate with the address.
 **EmergencyEnabled** | **bool** | Whether to enable emergency calling on the address. Can be: &#x60;true&#x60; or &#x60;false&#x60;.
 **FriendlyName** | **string** | A descriptive string that you create to describe the address. It can be up to 64 characters long.
 **PostalCode** | **string** | The postal code of the address.
 **Region** | **string** | The state or region of the address.
 **Street** | **string** | The number and street address of the address.

### Return type

[**ApiV2010AccountAddress**](ApiV2010AccountAddress.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateApplication

> ApiV2010AccountApplication UpdateApplication(ctx, AccountSid, Sid).ApiVersion(ApiVersion).FriendlyName(FriendlyName).MessageStatusCallback(MessageStatusCallback).SmsFallbackMethod(SmsFallbackMethod).SmsFallbackUrl(SmsFallbackUrl).SmsMethod(SmsMethod).SmsStatusCallback(SmsStatusCallback).SmsUrl(SmsUrl).StatusCallback(StatusCallback).StatusCallbackMethod(StatusCallbackMethod).VoiceCallerIdLookup(VoiceCallerIdLookup).VoiceFallbackMethod(VoiceFallbackMethod).VoiceFallbackUrl(VoiceFallbackUrl).VoiceMethod(VoiceMethod).VoiceUrl(VoiceUrl).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AccountSid := "AccountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Application resources to update.
    Sid := "Sid_example" // string | The Twilio-provided string that uniquely identifies the Application resource to update.
    ApiVersion := "ApiVersion_example" // string | The API version to use to start a new TwiML session. Can be: `2010-04-01` or `2008-08-01`. The default value is your account's default API version. (optional)
    FriendlyName := "FriendlyName_example" // string | A descriptive string that you create to describe the resource. It can be up to 64 characters long. (optional)
    MessageStatusCallback := "MessageStatusCallback_example" // string | The URL we should call using a POST method to send message status information to your application. (optional)
    SmsFallbackMethod := "SmsFallbackMethod_example" // string | The HTTP method we should use to call `sms_fallback_url`. Can be: `GET` or `POST`. (optional)
    SmsFallbackUrl := "SmsFallbackUrl_example" // string | The URL that we should call when an error occurs while retrieving or executing the TwiML from `sms_url`. (optional)
    SmsMethod := "SmsMethod_example" // string | The HTTP method we should use to call `sms_url`. Can be: `GET` or `POST`. (optional)
    SmsStatusCallback := "SmsStatusCallback_example" // string | Same as message_status_callback: The URL we should call using a POST method to send status information about SMS messages sent by the application. Deprecated, included for backwards compatibility. (optional)
    SmsUrl := "SmsUrl_example" // string | The URL we should call when the phone number receives an incoming SMS message. (optional)
    StatusCallback := "StatusCallback_example" // string | The URL we should call using the `status_callback_method` to send status information to your application. (optional)
    StatusCallbackMethod := "StatusCallbackMethod_example" // string | The HTTP method we should use to call `status_callback`. Can be: `GET` or `POST`. (optional)
    VoiceCallerIdLookup := true // bool | Whether we should look up the caller's caller-ID name from the CNAM database (additional charges apply). Can be: `true` or `false`. (optional)
    VoiceFallbackMethod := "VoiceFallbackMethod_example" // string | The HTTP method we should use to call `voice_fallback_url`. Can be: `GET` or `POST`. (optional)
    VoiceFallbackUrl := "VoiceFallbackUrl_example" // string | The URL that we should call when an error occurs retrieving or executing the TwiML requested by `url`. (optional)
    VoiceMethod := "VoiceMethod_example" // string | The HTTP method we should use to call `voice_url`. Can be: `GET` or `POST`. (optional)
    VoiceUrl := "VoiceUrl_example" // string | The URL we should call when the phone number assigned to this application receives a call. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateApplication(context.Background(), AccountSid, Sid).ApiVersion(ApiVersion).FriendlyName(FriendlyName).MessageStatusCallback(MessageStatusCallback).SmsFallbackMethod(SmsFallbackMethod).SmsFallbackUrl(SmsFallbackUrl).SmsMethod(SmsMethod).SmsStatusCallback(SmsStatusCallback).SmsUrl(SmsUrl).StatusCallback(StatusCallback).StatusCallbackMethod(StatusCallbackMethod).VoiceCallerIdLookup(VoiceCallerIdLookup).VoiceFallbackMethod(VoiceFallbackMethod).VoiceFallbackUrl(VoiceFallbackUrl).VoiceMethod(VoiceMethod).VoiceUrl(VoiceUrl).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateApplication``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateApplication`: ApiV2010AccountApplication
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateApplication`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Application resources to update.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the Application resource to update.

### Other Parameters

Other parameters are passed through a pointer to a UpdateApplicationParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


 **ApiVersion** | **string** | The API version to use to start a new TwiML session. Can be: &#x60;2010-04-01&#x60; or &#x60;2008-08-01&#x60;. The default value is your account&#39;s default API version.
 **FriendlyName** | **string** | A descriptive string that you create to describe the resource. It can be up to 64 characters long.
 **MessageStatusCallback** | **string** | The URL we should call using a POST method to send message status information to your application.
 **SmsFallbackMethod** | **string** | The HTTP method we should use to call &#x60;sms_fallback_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60;.
 **SmsFallbackUrl** | **string** | The URL that we should call when an error occurs while retrieving or executing the TwiML from &#x60;sms_url&#x60;.
 **SmsMethod** | **string** | The HTTP method we should use to call &#x60;sms_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60;.
 **SmsStatusCallback** | **string** | Same as message_status_callback: The URL we should call using a POST method to send status information about SMS messages sent by the application. Deprecated, included for backwards compatibility.
 **SmsUrl** | **string** | The URL we should call when the phone number receives an incoming SMS message.
 **StatusCallback** | **string** | The URL we should call using the &#x60;status_callback_method&#x60; to send status information to your application.
 **StatusCallbackMethod** | **string** | The HTTP method we should use to call &#x60;status_callback&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60;.
 **VoiceCallerIdLookup** | **bool** | Whether we should look up the caller&#39;s caller-ID name from the CNAM database (additional charges apply). Can be: &#x60;true&#x60; or &#x60;false&#x60;.
 **VoiceFallbackMethod** | **string** | The HTTP method we should use to call &#x60;voice_fallback_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60;.
 **VoiceFallbackUrl** | **string** | The URL that we should call when an error occurs retrieving or executing the TwiML requested by &#x60;url&#x60;.
 **VoiceMethod** | **string** | The HTTP method we should use to call &#x60;voice_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60;.
 **VoiceUrl** | **string** | The URL we should call when the phone number assigned to this application receives a call.

### Return type

[**ApiV2010AccountApplication**](ApiV2010AccountApplication.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCall

> ApiV2010AccountCall UpdateCall(ctx, AccountSid, Sid).FallbackMethod(FallbackMethod).FallbackUrl(FallbackUrl).Method(Method).Status(Status).StatusCallback(StatusCallback).StatusCallbackMethod(StatusCallbackMethod).Twiml(Twiml).Url(Url).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AccountSid := "AccountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Call resource(s) to update.
    Sid := "Sid_example" // string | The Twilio-provided string that uniquely identifies the Call resource to update
    FallbackMethod := "FallbackMethod_example" // string | The HTTP method that we should use to request the `fallback_url`. Can be: `GET` or `POST` and the default is `POST`. If an `application_sid` parameter is present, this parameter is ignored. (optional)
    FallbackUrl := "FallbackUrl_example" // string | The URL that we call using the `fallback_method` if an error occurs when requesting or executing the TwiML at `url`. If an `application_sid` parameter is present, this parameter is ignored. (optional)
    Method := "Method_example" // string | The HTTP method we should use when calling the `url`. Can be: `GET` or `POST` and the default is `POST`. If an `application_sid` parameter is present, this parameter is ignored. (optional)
    Status := "Status_example" // string | The new status of the resource. Can be: `canceled` or `completed`. Specifying `canceled` will attempt to hang up calls that are queued or ringing; however, it will not affect calls already in progress. Specifying `completed` will attempt to hang up a call even if it's already in progress. (optional)
    StatusCallback := "StatusCallback_example" // string | The URL we should call using the `status_callback_method` to send status information to your application. If no `status_callback_event` is specified, we will send the `completed` status. If an `application_sid` parameter is present, this parameter is ignored. URLs must contain a valid hostname (underscores are not permitted). (optional)
    StatusCallbackMethod := "StatusCallbackMethod_example" // string | The HTTP method we should use when requesting the `status_callback` URL. Can be: `GET` or `POST` and the default is `POST`. If an `application_sid` parameter is present, this parameter is ignored. (optional)
    Twiml := "Twiml_example" // string | TwiML instructions for the call Twilio will use without fetching Twiml from url. Twiml and url parameters are mutually exclusive (optional)
    Url := "Url_example" // string | The absolute URL that returns the TwiML instructions for the call. We will call this URL using the `method` when the call connects. For more information, see the [Url Parameter](https://www.twilio.com/docs/voice/make-calls#specify-a-url-parameter) section in [Making Calls](https://www.twilio.com/docs/voice/make-calls). (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateCall(context.Background(), AccountSid, Sid).FallbackMethod(FallbackMethod).FallbackUrl(FallbackUrl).Method(Method).Status(Status).StatusCallback(StatusCallback).StatusCallbackMethod(StatusCallbackMethod).Twiml(Twiml).Url(Url).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateCall``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCall`: ApiV2010AccountCall
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateCall`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Call resource(s) to update.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the Call resource to update

### Other Parameters

Other parameters are passed through a pointer to a UpdateCallParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


 **FallbackMethod** | **string** | The HTTP method that we should use to request the &#x60;fallback_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60; and the default is &#x60;POST&#x60;. If an &#x60;application_sid&#x60; parameter is present, this parameter is ignored.
 **FallbackUrl** | **string** | The URL that we call using the &#x60;fallback_method&#x60; if an error occurs when requesting or executing the TwiML at &#x60;url&#x60;. If an &#x60;application_sid&#x60; parameter is present, this parameter is ignored.
 **Method** | **string** | The HTTP method we should use when calling the &#x60;url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60; and the default is &#x60;POST&#x60;. If an &#x60;application_sid&#x60; parameter is present, this parameter is ignored.
 **Status** | **string** | The new status of the resource. Can be: &#x60;canceled&#x60; or &#x60;completed&#x60;. Specifying &#x60;canceled&#x60; will attempt to hang up calls that are queued or ringing; however, it will not affect calls already in progress. Specifying &#x60;completed&#x60; will attempt to hang up a call even if it&#39;s already in progress.
 **StatusCallback** | **string** | The URL we should call using the &#x60;status_callback_method&#x60; to send status information to your application. If no &#x60;status_callback_event&#x60; is specified, we will send the &#x60;completed&#x60; status. If an &#x60;application_sid&#x60; parameter is present, this parameter is ignored. URLs must contain a valid hostname (underscores are not permitted).
 **StatusCallbackMethod** | **string** | The HTTP method we should use when requesting the &#x60;status_callback&#x60; URL. Can be: &#x60;GET&#x60; or &#x60;POST&#x60; and the default is &#x60;POST&#x60;. If an &#x60;application_sid&#x60; parameter is present, this parameter is ignored.
 **Twiml** | **string** | TwiML instructions for the call Twilio will use without fetching Twiml from url. Twiml and url parameters are mutually exclusive
 **Url** | **string** | The absolute URL that returns the TwiML instructions for the call. We will call this URL using the &#x60;method&#x60; when the call connects. For more information, see the [Url Parameter](https://www.twilio.com/docs/voice/make-calls#specify-a-url-parameter) section in [Making Calls](https://www.twilio.com/docs/voice/make-calls).

### Return type

[**ApiV2010AccountCall**](ApiV2010AccountCall.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCallFeedback

> ApiV2010AccountCallCallFeedback UpdateCallFeedback(ctx, AccountSid, CallSid).Issue(Issue).QualityScore(QualityScore).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AccountSid := "AccountSid_example" // string | The unique id of the [Account](https://www.twilio.com/docs/iam/api/account) responsible for this resource.
    CallSid := "CallSid_example" // string | The call sid that uniquely identifies the call
    Issue := []string{"Inner_example"} // []string | One or more issues experienced during the call. The issues can be: `imperfect-audio`, `dropped-call`, `incorrect-caller-id`, `post-dial-delay`, `digits-not-captured`, `audio-latency`, `unsolicited-call`, or `one-way-audio`. (optional)
    QualityScore := int32(56) // int32 | The call quality expressed as an integer from `1` to `5` where `1` represents very poor call quality and `5` represents a perfect call. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateCallFeedback(context.Background(), AccountSid, CallSid).Issue(Issue).QualityScore(QualityScore).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateCallFeedback``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCallFeedback`: ApiV2010AccountCallCallFeedback
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateCallFeedback`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string** | The unique id of the [Account](https://www.twilio.com/docs/iam/api/account) responsible for this resource.
**CallSid** | **string** | The call sid that uniquely identifies the call

### Other Parameters

Other parameters are passed through a pointer to a UpdateCallFeedbackParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


 **Issue** | **[]string** | One or more issues experienced during the call. The issues can be: &#x60;imperfect-audio&#x60;, &#x60;dropped-call&#x60;, &#x60;incorrect-caller-id&#x60;, &#x60;post-dial-delay&#x60;, &#x60;digits-not-captured&#x60;, &#x60;audio-latency&#x60;, &#x60;unsolicited-call&#x60;, or &#x60;one-way-audio&#x60;.
 **QualityScore** | **int32** | The call quality expressed as an integer from &#x60;1&#x60; to &#x60;5&#x60; where &#x60;1&#x60; represents very poor call quality and &#x60;5&#x60; represents a perfect call.

### Return type

[**ApiV2010AccountCallCallFeedback**](ApiV2010AccountCallCallFeedback.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCallRecording

> ApiV2010AccountCallCallRecording UpdateCallRecording(ctx, AccountSid, CallSid, Sid).PauseBehavior(PauseBehavior).Status(Status).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AccountSid := "AccountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Recording resource to update.
    CallSid := "CallSid_example" // string | The [Call](https://www.twilio.com/docs/voice/api/call-resource) SID of the resource to update.
    Sid := "Sid_example" // string | The Twilio-provided string that uniquely identifies the Recording resource to update.
    PauseBehavior := "PauseBehavior_example" // string | Whether to record during a pause. Can be: `skip` or `silence` and the default is `silence`. `skip` does not record during the pause period, while `silence` will replace the actual audio of the call with silence during the pause period. This parameter only applies when setting `status` is set to `paused`. (optional)
    Status := "Status_example" // string | The new status of the recording. Can be: `stopped`, `paused`, `in-progress`. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateCallRecording(context.Background(), AccountSid, CallSid, Sid).PauseBehavior(PauseBehavior).Status(Status).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateCallRecording``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCallRecording`: ApiV2010AccountCallCallRecording
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateCallRecording`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Recording resource to update.
**CallSid** | **string** | The [Call](https://www.twilio.com/docs/voice/api/call-resource) SID of the resource to update.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the Recording resource to update.

### Other Parameters

Other parameters are passed through a pointer to a UpdateCallRecordingParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------



 **PauseBehavior** | **string** | Whether to record during a pause. Can be: &#x60;skip&#x60; or &#x60;silence&#x60; and the default is &#x60;silence&#x60;. &#x60;skip&#x60; does not record during the pause period, while &#x60;silence&#x60; will replace the actual audio of the call with silence during the pause period. This parameter only applies when setting &#x60;status&#x60; is set to &#x60;paused&#x60;.
 **Status** | **string** | The new status of the recording. Can be: &#x60;stopped&#x60;, &#x60;paused&#x60;, &#x60;in-progress&#x60;.

### Return type

[**ApiV2010AccountCallCallRecording**](ApiV2010AccountCallCallRecording.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateConference

> ApiV2010AccountConference UpdateConference(ctx, AccountSid, Sid).AnnounceMethod(AnnounceMethod).AnnounceUrl(AnnounceUrl).Status(Status).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AccountSid := "AccountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Conference resource(s) to update.
    Sid := "Sid_example" // string | The Twilio-provided string that uniquely identifies the Conference resource to update
    AnnounceMethod := "AnnounceMethod_example" // string | The HTTP method used to call `announce_url`. Can be: `GET` or `POST` and the default is `POST` (optional)
    AnnounceUrl := "AnnounceUrl_example" // string | The URL we should call to announce something into the conference. The URL can return an MP3, a WAV, or a TwiML document with `<Play>` or `<Say>`. (optional)
    Status := "Status_example" // string | The new status of the resource. Can be:  Can be: `init`, `in-progress`, or `completed`. Specifying `completed` will end the conference and hang up all participants (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateConference(context.Background(), AccountSid, Sid).AnnounceMethod(AnnounceMethod).AnnounceUrl(AnnounceUrl).Status(Status).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateConference``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateConference`: ApiV2010AccountConference
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateConference`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Conference resource(s) to update.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the Conference resource to update

### Other Parameters

Other parameters are passed through a pointer to a UpdateConferenceParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


 **AnnounceMethod** | **string** | The HTTP method used to call &#x60;announce_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60; and the default is &#x60;POST&#x60;
 **AnnounceUrl** | **string** | The URL we should call to announce something into the conference. The URL can return an MP3, a WAV, or a TwiML document with &#x60;&lt;Play&gt;&#x60; or &#x60;&lt;Say&gt;&#x60;.
 **Status** | **string** | The new status of the resource. Can be:  Can be: &#x60;init&#x60;, &#x60;in-progress&#x60;, or &#x60;completed&#x60;. Specifying &#x60;completed&#x60; will end the conference and hang up all participants

### Return type

[**ApiV2010AccountConference**](ApiV2010AccountConference.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateConferenceRecording

> ApiV2010AccountConferenceConferenceRecording UpdateConferenceRecording(ctx, AccountSid, ConferenceSid, Sid).PauseBehavior(PauseBehavior).Status(Status).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AccountSid := "AccountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Conference Recording resource to update.
    ConferenceSid := "ConferenceSid_example" // string | The Conference SID that identifies the conference associated with the recording to update.
    Sid := "Sid_example" // string | The Twilio-provided string that uniquely identifies the Conference Recording resource to update. Use `Twilio.CURRENT` to reference the current active recording.
    PauseBehavior := "PauseBehavior_example" // string | Whether to record during a pause. Can be: `skip` or `silence` and the default is `silence`. `skip` does not record during the pause period, while `silence` will replace the actual audio of the call with silence during the pause period. This parameter only applies when setting `status` is set to `paused`. (optional)
    Status := "Status_example" // string | The new status of the recording. Can be: `stopped`, `paused`, `in-progress`. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateConferenceRecording(context.Background(), AccountSid, ConferenceSid, Sid).PauseBehavior(PauseBehavior).Status(Status).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateConferenceRecording``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateConferenceRecording`: ApiV2010AccountConferenceConferenceRecording
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateConferenceRecording`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Conference Recording resource to update.
**ConferenceSid** | **string** | The Conference SID that identifies the conference associated with the recording to update.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the Conference Recording resource to update. Use &#x60;Twilio.CURRENT&#x60; to reference the current active recording.

### Other Parameters

Other parameters are passed through a pointer to a UpdateConferenceRecordingParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------



 **PauseBehavior** | **string** | Whether to record during a pause. Can be: &#x60;skip&#x60; or &#x60;silence&#x60; and the default is &#x60;silence&#x60;. &#x60;skip&#x60; does not record during the pause period, while &#x60;silence&#x60; will replace the actual audio of the call with silence during the pause period. This parameter only applies when setting &#x60;status&#x60; is set to &#x60;paused&#x60;.
 **Status** | **string** | The new status of the recording. Can be: &#x60;stopped&#x60;, &#x60;paused&#x60;, &#x60;in-progress&#x60;.

### Return type

[**ApiV2010AccountConferenceConferenceRecording**](ApiV2010AccountConferenceConferenceRecording.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateConnectApp

> ApiV2010AccountConnectApp UpdateConnectApp(ctx, AccountSid, Sid).AuthorizeRedirectUrl(AuthorizeRedirectUrl).CompanyName(CompanyName).DeauthorizeCallbackMethod(DeauthorizeCallbackMethod).DeauthorizeCallbackUrl(DeauthorizeCallbackUrl).Description(Description).FriendlyName(FriendlyName).HomepageUrl(HomepageUrl).Permissions(Permissions).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AccountSid := "AccountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the ConnectApp resources to update.
    Sid := "Sid_example" // string | The Twilio-provided string that uniquely identifies the ConnectApp resource to update.
    AuthorizeRedirectUrl := "AuthorizeRedirectUrl_example" // string | The URL to redirect the user to after we authenticate the user and obtain authorization to access the Connect App. (optional)
    CompanyName := "CompanyName_example" // string | The company name to set for the Connect App. (optional)
    DeauthorizeCallbackMethod := "DeauthorizeCallbackMethod_example" // string | The HTTP method to use when calling `deauthorize_callback_url`. (optional)
    DeauthorizeCallbackUrl := "DeauthorizeCallbackUrl_example" // string | The URL to call using the `deauthorize_callback_method` to de-authorize the Connect App. (optional)
    Description := "Description_example" // string | A description of the Connect App. (optional)
    FriendlyName := "FriendlyName_example" // string | A descriptive string that you create to describe the resource. It can be up to 64 characters long. (optional)
    HomepageUrl := "HomepageUrl_example" // string | A public URL where users can obtain more information about this Connect App. (optional)
    Permissions := []string{"Inner_example"} // []string | A comma-separated list of the permissions you will request from the users of this ConnectApp.  Can include: `get-all` and `post-all`. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateConnectApp(context.Background(), AccountSid, Sid).AuthorizeRedirectUrl(AuthorizeRedirectUrl).CompanyName(CompanyName).DeauthorizeCallbackMethod(DeauthorizeCallbackMethod).DeauthorizeCallbackUrl(DeauthorizeCallbackUrl).Description(Description).FriendlyName(FriendlyName).HomepageUrl(HomepageUrl).Permissions(Permissions).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateConnectApp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateConnectApp`: ApiV2010AccountConnectApp
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateConnectApp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the ConnectApp resources to update.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the ConnectApp resource to update.

### Other Parameters

Other parameters are passed through a pointer to a UpdateConnectAppParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


 **AuthorizeRedirectUrl** | **string** | The URL to redirect the user to after we authenticate the user and obtain authorization to access the Connect App.
 **CompanyName** | **string** | The company name to set for the Connect App.
 **DeauthorizeCallbackMethod** | **string** | The HTTP method to use when calling &#x60;deauthorize_callback_url&#x60;.
 **DeauthorizeCallbackUrl** | **string** | The URL to call using the &#x60;deauthorize_callback_method&#x60; to de-authorize the Connect App.
 **Description** | **string** | A description of the Connect App.
 **FriendlyName** | **string** | A descriptive string that you create to describe the resource. It can be up to 64 characters long.
 **HomepageUrl** | **string** | A public URL where users can obtain more information about this Connect App.
 **Permissions** | **[]string** | A comma-separated list of the permissions you will request from the users of this ConnectApp.  Can include: &#x60;get-all&#x60; and &#x60;post-all&#x60;.

### Return type

[**ApiV2010AccountConnectApp**](ApiV2010AccountConnectApp.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateIncomingPhoneNumber

> ApiV2010AccountIncomingPhoneNumber UpdateIncomingPhoneNumber(ctx, AccountSid, Sid).AccountSid2(AccountSid2).AddressSid(AddressSid).ApiVersion(ApiVersion).BundleSid(BundleSid).EmergencyAddressSid(EmergencyAddressSid).EmergencyStatus(EmergencyStatus).FriendlyName(FriendlyName).IdentitySid(IdentitySid).SmsApplicationSid(SmsApplicationSid).SmsFallbackMethod(SmsFallbackMethod).SmsFallbackUrl(SmsFallbackUrl).SmsMethod(SmsMethod).SmsUrl(SmsUrl).StatusCallback(StatusCallback).StatusCallbackMethod(StatusCallbackMethod).TrunkSid(TrunkSid).VoiceApplicationSid(VoiceApplicationSid).VoiceCallerIdLookup(VoiceCallerIdLookup).VoiceFallbackMethod(VoiceFallbackMethod).VoiceFallbackUrl(VoiceFallbackUrl).VoiceMethod(VoiceMethod).VoiceReceiveMode(VoiceReceiveMode).VoiceUrl(VoiceUrl).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AccountSid := "AccountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the IncomingPhoneNumber resource to update.  For more information, see [Exchanging Numbers Between Subaccounts](https://www.twilio.com/docs/iam/api/subaccounts#exchanging-numbers).
    Sid := "Sid_example" // string | The Twilio-provided string that uniquely identifies the IncomingPhoneNumber resource to update.
    AccountSid2 := "AccountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the IncomingPhoneNumber resource to update.  For more information, see [Exchanging Numbers Between Subaccounts](https://www.twilio.com/docs/iam/api/subaccounts#exchanging-numbers). (optional)
    AddressSid := "AddressSid_example" // string | The SID of the Address resource we should associate with the phone number. Some regions require addresses to meet local regulations. (optional)
    ApiVersion := "ApiVersion_example" // string | The API version to use for incoming calls made to the phone number. The default is `2010-04-01`. (optional)
    BundleSid := "BundleSid_example" // string | The SID of the Bundle resource that you associate with the phone number. Some regions require a Bundle to meet local Regulations. (optional)
    EmergencyAddressSid := "EmergencyAddressSid_example" // string | The SID of the emergency address configuration to use for emergency calling from this phone number. (optional)
    EmergencyStatus := "EmergencyStatus_example" // string | The configuration status parameter that determines whether the phone number is enabled for emergency calling. (optional)
    FriendlyName := "FriendlyName_example" // string | A descriptive string that you created to describe this phone number. It can be up to 64 characters long. By default, this is a formatted version of the phone number. (optional)
    IdentitySid := "IdentitySid_example" // string | The SID of the Identity resource that we should associate with the phone number. Some regions require an identity to meet local regulations. (optional)
    SmsApplicationSid := "SmsApplicationSid_example" // string | The SID of the application that should handle SMS messages sent to the number. If an `sms_application_sid` is present, we ignore all of the `sms_*_url` urls and use those set on the application. (optional)
    SmsFallbackMethod := "SmsFallbackMethod_example" // string | The HTTP method that we should use to call `sms_fallback_url`. Can be: `GET` or `POST` and defaults to `POST`. (optional)
    SmsFallbackUrl := "SmsFallbackUrl_example" // string | The URL that we should call when an error occurs while requesting or executing the TwiML defined by `sms_url`. (optional)
    SmsMethod := "SmsMethod_example" // string | The HTTP method that we should use to call `sms_url`. Can be: `GET` or `POST` and defaults to `POST`. (optional)
    SmsUrl := "SmsUrl_example" // string | The URL we should call when the phone number receives an incoming SMS message. (optional)
    StatusCallback := "StatusCallback_example" // string | The URL we should call using the `status_callback_method` to send status information to your application. (optional)
    StatusCallbackMethod := "StatusCallbackMethod_example" // string | The HTTP method we should use to call `status_callback`. Can be: `GET` or `POST` and defaults to `POST`. (optional)
    TrunkSid := "TrunkSid_example" // string | The SID of the Trunk we should use to handle phone calls to the phone number. If a `trunk_sid` is present, we ignore all of the voice urls and voice applications and use only those set on the Trunk. Setting a `trunk_sid` will automatically delete your `voice_application_sid` and vice versa. (optional)
    VoiceApplicationSid := "VoiceApplicationSid_example" // string | The SID of the application we should use to handle phone calls to the phone number. If a `voice_application_sid` is present, we ignore all of the voice urls and use only those set on the application. Setting a `voice_application_sid` will automatically delete your `trunk_sid` and vice versa. (optional)
    VoiceCallerIdLookup := true // bool | Whether to lookup the caller's name from the CNAM database and post it to your app. Can be: `true` or `false` and defaults to `false`. (optional)
    VoiceFallbackMethod := "VoiceFallbackMethod_example" // string | The HTTP method that we should use to call `voice_fallback_url`. Can be: `GET` or `POST` and defaults to `POST`. (optional)
    VoiceFallbackUrl := "VoiceFallbackUrl_example" // string | The URL that we should call when an error occurs retrieving or executing the TwiML requested by `url`. (optional)
    VoiceMethod := "VoiceMethod_example" // string | The HTTP method that we should use to call `voice_url`. Can be: `GET` or `POST` and defaults to `POST`. (optional)
    VoiceReceiveMode := "VoiceReceiveMode_example" // string | The configuration parameter for the phone number to receive incoming voice calls or faxes. Can be: `fax` or `voice` and defaults to `voice`. (optional)
    VoiceUrl := "VoiceUrl_example" // string | The URL that we should call to answer a call to the phone number. The `voice_url` will not be called if a `voice_application_sid` or a `trunk_sid` is set. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateIncomingPhoneNumber(context.Background(), AccountSid, Sid).AccountSid2(AccountSid2).AddressSid(AddressSid).ApiVersion(ApiVersion).BundleSid(BundleSid).EmergencyAddressSid(EmergencyAddressSid).EmergencyStatus(EmergencyStatus).FriendlyName(FriendlyName).IdentitySid(IdentitySid).SmsApplicationSid(SmsApplicationSid).SmsFallbackMethod(SmsFallbackMethod).SmsFallbackUrl(SmsFallbackUrl).SmsMethod(SmsMethod).SmsUrl(SmsUrl).StatusCallback(StatusCallback).StatusCallbackMethod(StatusCallbackMethod).TrunkSid(TrunkSid).VoiceApplicationSid(VoiceApplicationSid).VoiceCallerIdLookup(VoiceCallerIdLookup).VoiceFallbackMethod(VoiceFallbackMethod).VoiceFallbackUrl(VoiceFallbackUrl).VoiceMethod(VoiceMethod).VoiceReceiveMode(VoiceReceiveMode).VoiceUrl(VoiceUrl).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateIncomingPhoneNumber``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateIncomingPhoneNumber`: ApiV2010AccountIncomingPhoneNumber
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateIncomingPhoneNumber`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the IncomingPhoneNumber resource to update.  For more information, see [Exchanging Numbers Between Subaccounts](https://www.twilio.com/docs/iam/api/subaccounts#exchanging-numbers).
**Sid** | **string** | The Twilio-provided string that uniquely identifies the IncomingPhoneNumber resource to update.

### Other Parameters

Other parameters are passed through a pointer to a UpdateIncomingPhoneNumberParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


 **AccountSid2** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the IncomingPhoneNumber resource to update.  For more information, see [Exchanging Numbers Between Subaccounts](https://www.twilio.com/docs/iam/api/subaccounts#exchanging-numbers).
 **AddressSid** | **string** | The SID of the Address resource we should associate with the phone number. Some regions require addresses to meet local regulations.
 **ApiVersion** | **string** | The API version to use for incoming calls made to the phone number. The default is &#x60;2010-04-01&#x60;.
 **BundleSid** | **string** | The SID of the Bundle resource that you associate with the phone number. Some regions require a Bundle to meet local Regulations.
 **EmergencyAddressSid** | **string** | The SID of the emergency address configuration to use for emergency calling from this phone number.
 **EmergencyStatus** | **string** | The configuration status parameter that determines whether the phone number is enabled for emergency calling.
 **FriendlyName** | **string** | A descriptive string that you created to describe this phone number. It can be up to 64 characters long. By default, this is a formatted version of the phone number.
 **IdentitySid** | **string** | The SID of the Identity resource that we should associate with the phone number. Some regions require an identity to meet local regulations.
 **SmsApplicationSid** | **string** | The SID of the application that should handle SMS messages sent to the number. If an &#x60;sms_application_sid&#x60; is present, we ignore all of the &#x60;sms_*_url&#x60; urls and use those set on the application.
 **SmsFallbackMethod** | **string** | The HTTP method that we should use to call &#x60;sms_fallback_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60; and defaults to &#x60;POST&#x60;.
 **SmsFallbackUrl** | **string** | The URL that we should call when an error occurs while requesting or executing the TwiML defined by &#x60;sms_url&#x60;.
 **SmsMethod** | **string** | The HTTP method that we should use to call &#x60;sms_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60; and defaults to &#x60;POST&#x60;.
 **SmsUrl** | **string** | The URL we should call when the phone number receives an incoming SMS message.
 **StatusCallback** | **string** | The URL we should call using the &#x60;status_callback_method&#x60; to send status information to your application.
 **StatusCallbackMethod** | **string** | The HTTP method we should use to call &#x60;status_callback&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60; and defaults to &#x60;POST&#x60;.
 **TrunkSid** | **string** | The SID of the Trunk we should use to handle phone calls to the phone number. If a &#x60;trunk_sid&#x60; is present, we ignore all of the voice urls and voice applications and use only those set on the Trunk. Setting a &#x60;trunk_sid&#x60; will automatically delete your &#x60;voice_application_sid&#x60; and vice versa.
 **VoiceApplicationSid** | **string** | The SID of the application we should use to handle phone calls to the phone number. If a &#x60;voice_application_sid&#x60; is present, we ignore all of the voice urls and use only those set on the application. Setting a &#x60;voice_application_sid&#x60; will automatically delete your &#x60;trunk_sid&#x60; and vice versa.
 **VoiceCallerIdLookup** | **bool** | Whether to lookup the caller&#39;s name from the CNAM database and post it to your app. Can be: &#x60;true&#x60; or &#x60;false&#x60; and defaults to &#x60;false&#x60;.
 **VoiceFallbackMethod** | **string** | The HTTP method that we should use to call &#x60;voice_fallback_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60; and defaults to &#x60;POST&#x60;.
 **VoiceFallbackUrl** | **string** | The URL that we should call when an error occurs retrieving or executing the TwiML requested by &#x60;url&#x60;.
 **VoiceMethod** | **string** | The HTTP method that we should use to call &#x60;voice_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60; and defaults to &#x60;POST&#x60;.
 **VoiceReceiveMode** | **string** | The configuration parameter for the phone number to receive incoming voice calls or faxes. Can be: &#x60;fax&#x60; or &#x60;voice&#x60; and defaults to &#x60;voice&#x60;.
 **VoiceUrl** | **string** | The URL that we should call to answer a call to the phone number. The &#x60;voice_url&#x60; will not be called if a &#x60;voice_application_sid&#x60; or a &#x60;trunk_sid&#x60; is set.

### Return type

[**ApiV2010AccountIncomingPhoneNumber**](ApiV2010AccountIncomingPhoneNumber.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateKey

> ApiV2010AccountKey UpdateKey(ctx, AccountSid, Sid).FriendlyName(FriendlyName).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AccountSid := "AccountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Key resources to update.
    Sid := "Sid_example" // string | The Twilio-provided string that uniquely identifies the Key resource to update.
    FriendlyName := "FriendlyName_example" // string | A descriptive string that you create to describe the resource. It can be up to 64 characters long. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateKey(context.Background(), AccountSid, Sid).FriendlyName(FriendlyName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateKey`: ApiV2010AccountKey
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Key resources to update.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the Key resource to update.

### Other Parameters

Other parameters are passed through a pointer to a UpdateKeyParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


 **FriendlyName** | **string** | A descriptive string that you create to describe the resource. It can be up to 64 characters long.

### Return type

[**ApiV2010AccountKey**](ApiV2010AccountKey.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMember

> ApiV2010AccountQueueMember UpdateMember(ctx, AccountSid, QueueSid, CallSid).Method(Method).Url(Url).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AccountSid := "AccountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Member resource(s) to update.
    QueueSid := "QueueSid_example" // string | The SID of the Queue in which to find the members to update.
    CallSid := "CallSid_example" // string | The [Call](https://www.twilio.com/docs/voice/api/call-resource) SID of the resource(s) to update.
    Method := "Method_example" // string | How to pass the update request data. Can be `GET` or `POST` and the default is `POST`. `POST` sends the data as encoded form data and `GET` sends the data as query parameters. (optional)
    Url := "Url_example" // string | The absolute URL of the Queue resource. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateMember(context.Background(), AccountSid, QueueSid, CallSid).Method(Method).Url(Url).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateMember``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateMember`: ApiV2010AccountQueueMember
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateMember`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Member resource(s) to update.
**QueueSid** | **string** | The SID of the Queue in which to find the members to update.
**CallSid** | **string** | The [Call](https://www.twilio.com/docs/voice/api/call-resource) SID of the resource(s) to update.

### Other Parameters

Other parameters are passed through a pointer to a UpdateMemberParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------



 **Method** | **string** | How to pass the update request data. Can be &#x60;GET&#x60; or &#x60;POST&#x60; and the default is &#x60;POST&#x60;. &#x60;POST&#x60; sends the data as encoded form data and &#x60;GET&#x60; sends the data as query parameters.
 **Url** | **string** | The absolute URL of the Queue resource.

### Return type

[**ApiV2010AccountQueueMember**](ApiV2010AccountQueueMember.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMessage

> ApiV2010AccountMessage UpdateMessage(ctx, AccountSid, Sid).Body(Body).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AccountSid := "AccountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Message resources to update.
    Sid := "Sid_example" // string | The Twilio-provided string that uniquely identifies the Message resource to update.
    Body := "Body_example" // string | The text of the message you want to send. Can be up to 1,600 characters long. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateMessage(context.Background(), AccountSid, Sid).Body(Body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateMessage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateMessage`: ApiV2010AccountMessage
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateMessage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Message resources to update.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the Message resource to update.

### Other Parameters

Other parameters are passed through a pointer to a UpdateMessageParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


 **Body** | **string** | The text of the message you want to send. Can be up to 1,600 characters long.

### Return type

[**ApiV2010AccountMessage**](ApiV2010AccountMessage.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOutgoingCallerId

> ApiV2010AccountOutgoingCallerId UpdateOutgoingCallerId(ctx, AccountSid, Sid).FriendlyName(FriendlyName).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AccountSid := "AccountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the OutgoingCallerId resources to update.
    Sid := "Sid_example" // string | The Twilio-provided string that uniquely identifies the OutgoingCallerId resource to update.
    FriendlyName := "FriendlyName_example" // string | A descriptive string that you create to describe the resource. It can be up to 64 characters long. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateOutgoingCallerId(context.Background(), AccountSid, Sid).FriendlyName(FriendlyName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateOutgoingCallerId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOutgoingCallerId`: ApiV2010AccountOutgoingCallerId
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateOutgoingCallerId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the OutgoingCallerId resources to update.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the OutgoingCallerId resource to update.

### Other Parameters

Other parameters are passed through a pointer to a UpdateOutgoingCallerIdParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


 **FriendlyName** | **string** | A descriptive string that you create to describe the resource. It can be up to 64 characters long.

### Return type

[**ApiV2010AccountOutgoingCallerId**](ApiV2010AccountOutgoingCallerId.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateParticipant

> ApiV2010AccountConferenceParticipant UpdateParticipant(ctx, AccountSid, ConferenceSid, CallSid).AnnounceMethod(AnnounceMethod).AnnounceUrl(AnnounceUrl).BeepOnExit(BeepOnExit).CallSidToCoach(CallSidToCoach).Coaching(Coaching).EndConferenceOnExit(EndConferenceOnExit).Hold(Hold).HoldMethod(HoldMethod).HoldUrl(HoldUrl).Muted(Muted).WaitMethod(WaitMethod).WaitUrl(WaitUrl).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AccountSid := "AccountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Participant resources to update.
    ConferenceSid := "ConferenceSid_example" // string | The SID of the conference with the participant to update.
    CallSid := "CallSid_example" // string | The [Call](https://www.twilio.com/docs/voice/api/call-resource) SID or label of the participant to update. Non URL safe characters in a label must be percent encoded, for example, a space character is represented as %20.
    AnnounceMethod := "AnnounceMethod_example" // string | The HTTP method we should use to call `announce_url`. Can be: `GET` or `POST` and defaults to `POST`. (optional)
    AnnounceUrl := "AnnounceUrl_example" // string | The URL we call using the `announce_method` for an announcement to the participant. The URL must return an MP3 file, a WAV file, or a TwiML document that contains `<Play>` or `<Say>` commands. (optional)
    BeepOnExit := true // bool | Whether to play a notification beep to the conference when the participant exits. Can be: `true` or `false`. (optional)
    CallSidToCoach := "CallSidToCoach_example" // string | The SID of the participant who is being `coached`. The participant being coached is the only participant who can hear the participant who is `coaching`. (optional)
    Coaching := true // bool | Whether the participant is coaching another call. Can be: `true` or `false`. If not present, defaults to `false` unless `call_sid_to_coach` is defined. If `true`, `call_sid_to_coach` must be defined. (optional)
    EndConferenceOnExit := true // bool | Whether to end the conference when the participant leaves. Can be: `true` or `false` and defaults to `false`. (optional)
    Hold := true // bool | Whether the participant should be on hold. Can be: `true` or `false`. `true` puts the participant on hold, and `false` lets them rejoin the conference. (optional)
    HoldMethod := "HoldMethod_example" // string | The HTTP method we should use to call `hold_url`. Can be: `GET` or `POST` and the default is `GET`. (optional)
    HoldUrl := "HoldUrl_example" // string | The URL we call using the `hold_method` for  music that plays when the participant is on hold. The URL may return an MP3 file, a WAV file, or a TwiML document that contains the `<Play>`, `<Say>` or `<Redirect>` commands. (optional)
    Muted := true // bool | Whether the participant should be muted. Can be `true` or `false`. `true` will mute the participant, and `false` will un-mute them. Anything value other than `true` or `false` is interpreted as `false`. (optional)
    WaitMethod := "WaitMethod_example" // string | The HTTP method we should use to call `wait_url`. Can be `GET` or `POST` and the default is `POST`. When using a static audio file, this should be `GET` so that we can cache the file. (optional)
    WaitUrl := "WaitUrl_example" // string | The URL we should call using the `wait_method` for the music to play while participants are waiting for the conference to start. The default value is the URL of our standard hold music. [Learn more about hold music](https://www.twilio.com/labs/twimlets/holdmusic). (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateParticipant(context.Background(), AccountSid, ConferenceSid, CallSid).AnnounceMethod(AnnounceMethod).AnnounceUrl(AnnounceUrl).BeepOnExit(BeepOnExit).CallSidToCoach(CallSidToCoach).Coaching(Coaching).EndConferenceOnExit(EndConferenceOnExit).Hold(Hold).HoldMethod(HoldMethod).HoldUrl(HoldUrl).Muted(Muted).WaitMethod(WaitMethod).WaitUrl(WaitUrl).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateParticipant``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateParticipant`: ApiV2010AccountConferenceParticipant
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateParticipant`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Participant resources to update.
**ConferenceSid** | **string** | The SID of the conference with the participant to update.
**CallSid** | **string** | The [Call](https://www.twilio.com/docs/voice/api/call-resource) SID or label of the participant to update. Non URL safe characters in a label must be percent encoded, for example, a space character is represented as %20.

### Other Parameters

Other parameters are passed through a pointer to a UpdateParticipantParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------



 **AnnounceMethod** | **string** | The HTTP method we should use to call &#x60;announce_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60; and defaults to &#x60;POST&#x60;.
 **AnnounceUrl** | **string** | The URL we call using the &#x60;announce_method&#x60; for an announcement to the participant. The URL must return an MP3 file, a WAV file, or a TwiML document that contains &#x60;&lt;Play&gt;&#x60; or &#x60;&lt;Say&gt;&#x60; commands.
 **BeepOnExit** | **bool** | Whether to play a notification beep to the conference when the participant exits. Can be: &#x60;true&#x60; or &#x60;false&#x60;.
 **CallSidToCoach** | **string** | The SID of the participant who is being &#x60;coached&#x60;. The participant being coached is the only participant who can hear the participant who is &#x60;coaching&#x60;.
 **Coaching** | **bool** | Whether the participant is coaching another call. Can be: &#x60;true&#x60; or &#x60;false&#x60;. If not present, defaults to &#x60;false&#x60; unless &#x60;call_sid_to_coach&#x60; is defined. If &#x60;true&#x60;, &#x60;call_sid_to_coach&#x60; must be defined.
 **EndConferenceOnExit** | **bool** | Whether to end the conference when the participant leaves. Can be: &#x60;true&#x60; or &#x60;false&#x60; and defaults to &#x60;false&#x60;.
 **Hold** | **bool** | Whether the participant should be on hold. Can be: &#x60;true&#x60; or &#x60;false&#x60;. &#x60;true&#x60; puts the participant on hold, and &#x60;false&#x60; lets them rejoin the conference.
 **HoldMethod** | **string** | The HTTP method we should use to call &#x60;hold_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60; and the default is &#x60;GET&#x60;.
 **HoldUrl** | **string** | The URL we call using the &#x60;hold_method&#x60; for  music that plays when the participant is on hold. The URL may return an MP3 file, a WAV file, or a TwiML document that contains the &#x60;&lt;Play&gt;&#x60;, &#x60;&lt;Say&gt;&#x60; or &#x60;&lt;Redirect&gt;&#x60; commands.
 **Muted** | **bool** | Whether the participant should be muted. Can be &#x60;true&#x60; or &#x60;false&#x60;. &#x60;true&#x60; will mute the participant, and &#x60;false&#x60; will un-mute them. Anything value other than &#x60;true&#x60; or &#x60;false&#x60; is interpreted as &#x60;false&#x60;.
 **WaitMethod** | **string** | The HTTP method we should use to call &#x60;wait_url&#x60;. Can be &#x60;GET&#x60; or &#x60;POST&#x60; and the default is &#x60;POST&#x60;. When using a static audio file, this should be &#x60;GET&#x60; so that we can cache the file.
 **WaitUrl** | **string** | The URL we should call using the &#x60;wait_method&#x60; for the music to play while participants are waiting for the conference to start. The default value is the URL of our standard hold music. [Learn more about hold music](https://www.twilio.com/labs/twimlets/holdmusic).

### Return type

[**ApiV2010AccountConferenceParticipant**](ApiV2010AccountConferenceParticipant.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePayments

> ApiV2010AccountCallPayments UpdatePayments(ctx, AccountSid, CallSid, Sid).Capture(Capture).IdempotencyKey(IdempotencyKey).Status(Status).StatusCallback(StatusCallback).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AccountSid := "AccountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that will update the resource.
    CallSid := "CallSid_example" // string | The SID of the call that will update the resource. This should be the same call sid that was used to create payments resource.
    Sid := "Sid_example" // string | The SID of Payments session that needs to be updated.
    Capture := "Capture_example" // string | The piece of payment information that you wish the caller to enter. Must be one of `payment-card-number`, `expiration-date`, `security-code`, `postal-code`, `bank-routing-number`, or `bank-account-number`. (optional)
    IdempotencyKey := "IdempotencyKey_example" // string | A unique token that will be used to ensure that multiple API calls with the same information do not result in multiple transactions. This should be a unique string value per API call and can be a randomly generated. (optional)
    Status := "Status_example" // string | Indicates whether the current payment session should be cancelled or completed. When `cancel` the payment session is cancelled. When `complete`, Twilio sends the payment information to the selected <Pay> connector for processing. (optional)
    StatusCallback := "StatusCallback_example" // string | Provide an absolute or relative URL to receive status updates regarding your Pay session. Read more about the [Update](https://www.twilio.com/docs/voice/api/payment-resource#statuscallback-update) and [Complete/Cancel](https://www.twilio.com/docs/voice/api/payment-resource#statuscallback-cancelcomplete) POST requests. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdatePayments(context.Background(), AccountSid, CallSid, Sid).Capture(Capture).IdempotencyKey(IdempotencyKey).Status(Status).StatusCallback(StatusCallback).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdatePayments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdatePayments`: ApiV2010AccountCallPayments
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdatePayments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that will update the resource.
**CallSid** | **string** | The SID of the call that will update the resource. This should be the same call sid that was used to create payments resource.
**Sid** | **string** | The SID of Payments session that needs to be updated.

### Other Parameters

Other parameters are passed through a pointer to a UpdatePaymentsParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------



 **Capture** | **string** | The piece of payment information that you wish the caller to enter. Must be one of &#x60;payment-card-number&#x60;, &#x60;expiration-date&#x60;, &#x60;security-code&#x60;, &#x60;postal-code&#x60;, &#x60;bank-routing-number&#x60;, or &#x60;bank-account-number&#x60;.
 **IdempotencyKey** | **string** | A unique token that will be used to ensure that multiple API calls with the same information do not result in multiple transactions. This should be a unique string value per API call and can be a randomly generated.
 **Status** | **string** | Indicates whether the current payment session should be cancelled or completed. When &#x60;cancel&#x60; the payment session is cancelled. When &#x60;complete&#x60;, Twilio sends the payment information to the selected &lt;Pay&gt; connector for processing.
 **StatusCallback** | **string** | Provide an absolute or relative URL to receive status updates regarding your Pay session. Read more about the [Update](https://www.twilio.com/docs/voice/api/payment-resource#statuscallback-update) and [Complete/Cancel](https://www.twilio.com/docs/voice/api/payment-resource#statuscallback-cancelcomplete) POST requests.

### Return type

[**ApiV2010AccountCallPayments**](ApiV2010AccountCallPayments.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateQueue

> ApiV2010AccountQueue UpdateQueue(ctx, AccountSid, Sid).FriendlyName(FriendlyName).MaxSize(MaxSize).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AccountSid := "AccountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Queue resource to update.
    Sid := "Sid_example" // string | The Twilio-provided string that uniquely identifies the Queue resource to update
    FriendlyName := "FriendlyName_example" // string | A descriptive string that you created to describe this resource. It can be up to 64 characters long. (optional)
    MaxSize := int32(56) // int32 | The maximum number of calls allowed to be in the queue. The default is 100. The maximum is 5000. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateQueue(context.Background(), AccountSid, Sid).FriendlyName(FriendlyName).MaxSize(MaxSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateQueue``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateQueue`: ApiV2010AccountQueue
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateQueue`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Queue resource to update.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the Queue resource to update

### Other Parameters

Other parameters are passed through a pointer to a UpdateQueueParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


 **FriendlyName** | **string** | A descriptive string that you created to describe this resource. It can be up to 64 characters long.
 **MaxSize** | **int32** | The maximum number of calls allowed to be in the queue. The default is 100. The maximum is 5000.

### Return type

[**ApiV2010AccountQueue**](ApiV2010AccountQueue.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateShortCode

> ApiV2010AccountShortCode UpdateShortCode(ctx, AccountSid, Sid).ApiVersion(ApiVersion).FriendlyName(FriendlyName).SmsFallbackMethod(SmsFallbackMethod).SmsFallbackUrl(SmsFallbackUrl).SmsMethod(SmsMethod).SmsUrl(SmsUrl).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AccountSid := "AccountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the ShortCode resource(s) to update.
    Sid := "Sid_example" // string | The Twilio-provided string that uniquely identifies the ShortCode resource to update
    ApiVersion := "ApiVersion_example" // string | The API version to use to start a new TwiML session. Can be: `2010-04-01` or `2008-08-01`. (optional)
    FriendlyName := "FriendlyName_example" // string | A descriptive string that you created to describe this resource. It can be up to 64 characters long. By default, the `FriendlyName` is the short code. (optional)
    SmsFallbackMethod := "SmsFallbackMethod_example" // string | The HTTP method that we should use to call the `sms_fallback_url`. Can be: `GET` or `POST`. (optional)
    SmsFallbackUrl := "SmsFallbackUrl_example" // string | The URL that we should call if an error occurs while retrieving or executing the TwiML from `sms_url`. (optional)
    SmsMethod := "SmsMethod_example" // string | The HTTP method we should use when calling the `sms_url`. Can be: `GET` or `POST`. (optional)
    SmsUrl := "SmsUrl_example" // string | The URL we should call when receiving an incoming SMS message to this short code. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateShortCode(context.Background(), AccountSid, Sid).ApiVersion(ApiVersion).FriendlyName(FriendlyName).SmsFallbackMethod(SmsFallbackMethod).SmsFallbackUrl(SmsFallbackUrl).SmsMethod(SmsMethod).SmsUrl(SmsUrl).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateShortCode``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateShortCode`: ApiV2010AccountShortCode
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateShortCode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the ShortCode resource(s) to update.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the ShortCode resource to update

### Other Parameters

Other parameters are passed through a pointer to a UpdateShortCodeParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


 **ApiVersion** | **string** | The API version to use to start a new TwiML session. Can be: &#x60;2010-04-01&#x60; or &#x60;2008-08-01&#x60;.
 **FriendlyName** | **string** | A descriptive string that you created to describe this resource. It can be up to 64 characters long. By default, the &#x60;FriendlyName&#x60; is the short code.
 **SmsFallbackMethod** | **string** | The HTTP method that we should use to call the &#x60;sms_fallback_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60;.
 **SmsFallbackUrl** | **string** | The URL that we should call if an error occurs while retrieving or executing the TwiML from &#x60;sms_url&#x60;.
 **SmsMethod** | **string** | The HTTP method we should use when calling the &#x60;sms_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60;.
 **SmsUrl** | **string** | The URL we should call when receiving an incoming SMS message to this short code.

### Return type

[**ApiV2010AccountShortCode**](ApiV2010AccountShortCode.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSigningKey

> ApiV2010AccountSigningKey UpdateSigningKey(ctx, AccountSid, Sid).FriendlyName(FriendlyName).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AccountSid := "AccountSid_example" // string | 
    Sid := "Sid_example" // string | 
    FriendlyName := "FriendlyName_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateSigningKey(context.Background(), AccountSid, Sid).FriendlyName(FriendlyName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateSigningKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSigningKey`: ApiV2010AccountSigningKey
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateSigningKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string** | 
**Sid** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a UpdateSigningKeyParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


 **FriendlyName** | **string** | 

### Return type

[**ApiV2010AccountSigningKey**](ApiV2010AccountSigningKey.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSipCredential

> ApiV2010AccountSipSipCredentialListSipCredential UpdateSipCredential(ctx, AccountSid, CredentialListSid, Sid).Password(Password).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AccountSid := "AccountSid_example" // string | The unique id of the Account that is responsible for this resource.
    CredentialListSid := "CredentialListSid_example" // string | The unique id that identifies the credential list that includes this credential.
    Sid := "Sid_example" // string | The unique id that identifies the resource to update.
    Password := "Password_example" // string | The password that the username will use when authenticating SIP requests. The password must be a minimum of 12 characters, contain at least 1 digit, and have mixed case. (eg `IWasAtSignal2018`) (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateSipCredential(context.Background(), AccountSid, CredentialListSid, Sid).Password(Password).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateSipCredential``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSipCredential`: ApiV2010AccountSipSipCredentialListSipCredential
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateSipCredential`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string** | The unique id of the Account that is responsible for this resource.
**CredentialListSid** | **string** | The unique id that identifies the credential list that includes this credential.
**Sid** | **string** | The unique id that identifies the resource to update.

### Other Parameters

Other parameters are passed through a pointer to a UpdateSipCredentialParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------



 **Password** | **string** | The password that the username will use when authenticating SIP requests. The password must be a minimum of 12 characters, contain at least 1 digit, and have mixed case. (eg &#x60;IWasAtSignal2018&#x60;)

### Return type

[**ApiV2010AccountSipSipCredentialListSipCredential**](ApiV2010AccountSipSipCredentialListSipCredential.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSipCredentialList

> ApiV2010AccountSipSipCredentialList UpdateSipCredentialList(ctx, AccountSid, Sid).FriendlyName(FriendlyName).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AccountSid := "AccountSid_example" // string | The unique id of the Account that is responsible for this resource.
    Sid := "Sid_example" // string | The credential list Sid that uniquely identifies this resource
    FriendlyName := "FriendlyName_example" // string | A human readable descriptive text for a CredentialList, up to 64 characters long. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateSipCredentialList(context.Background(), AccountSid, Sid).FriendlyName(FriendlyName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateSipCredentialList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSipCredentialList`: ApiV2010AccountSipSipCredentialList
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateSipCredentialList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string** | The unique id of the Account that is responsible for this resource.
**Sid** | **string** | The credential list Sid that uniquely identifies this resource

### Other Parameters

Other parameters are passed through a pointer to a UpdateSipCredentialListParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


 **FriendlyName** | **string** | A human readable descriptive text for a CredentialList, up to 64 characters long.

### Return type

[**ApiV2010AccountSipSipCredentialList**](ApiV2010AccountSipSipCredentialList.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSipDomain

> ApiV2010AccountSipSipDomain UpdateSipDomain(ctx, AccountSid, Sid).ByocTrunkSid(ByocTrunkSid).DomainName(DomainName).EmergencyCallerSid(EmergencyCallerSid).EmergencyCallingEnabled(EmergencyCallingEnabled).FriendlyName(FriendlyName).Secure(Secure).SipRegistration(SipRegistration).VoiceFallbackMethod(VoiceFallbackMethod).VoiceFallbackUrl(VoiceFallbackUrl).VoiceMethod(VoiceMethod).VoiceStatusCallbackMethod(VoiceStatusCallbackMethod).VoiceStatusCallbackUrl(VoiceStatusCallbackUrl).VoiceUrl(VoiceUrl).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AccountSid := "AccountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the SipDomain resource to update.
    Sid := "Sid_example" // string | The Twilio-provided string that uniquely identifies the SipDomain resource to update.
    ByocTrunkSid := "ByocTrunkSid_example" // string | The SID of the BYOC Trunk(Bring Your Own Carrier) resource that the Sip Domain will be associated with. (optional)
    DomainName := "DomainName_example" // string | The unique address you reserve on Twilio to which you route your SIP traffic. Domain names can contain letters, digits, and \\\"-\\\" and must end with `sip.twilio.com`. (optional)
    EmergencyCallerSid := "EmergencyCallerSid_example" // string | Whether an emergency caller sid is configured for the domain. If present, this phone number will be used as the callback for the emergency call. (optional)
    EmergencyCallingEnabled := true // bool | Whether emergency calling is enabled for the domain. If enabled, allows emergency calls on the domain from phone numbers with validated addresses. (optional)
    FriendlyName := "FriendlyName_example" // string | A descriptive string that you created to describe the resource. It can be up to 64 characters long. (optional)
    Secure := true // bool | Whether secure SIP is enabled for the domain. If enabled, TLS will be enforced and SRTP will be negotiated on all incoming calls to this sip domain. (optional)
    SipRegistration := true // bool | Whether to allow SIP Endpoints to register with the domain to receive calls. Can be `true` or `false`. `true` allows SIP Endpoints to register with the domain to receive calls, `false` does not. (optional)
    VoiceFallbackMethod := "VoiceFallbackMethod_example" // string | The HTTP method we should use to call `voice_fallback_url`. Can be: `GET` or `POST`. (optional)
    VoiceFallbackUrl := "VoiceFallbackUrl_example" // string | The URL that we should call when an error occurs while retrieving or executing the TwiML requested by `voice_url`. (optional)
    VoiceMethod := "VoiceMethod_example" // string | The HTTP method we should use to call `voice_url` (optional)
    VoiceStatusCallbackMethod := "VoiceStatusCallbackMethod_example" // string | The HTTP method we should use to call `voice_status_callback_url`. Can be: `GET` or `POST`. (optional)
    VoiceStatusCallbackUrl := "VoiceStatusCallbackUrl_example" // string | The URL that we should call to pass status parameters (such as call ended) to your application. (optional)
    VoiceUrl := "VoiceUrl_example" // string | The URL we should call when the domain receives a call. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateSipDomain(context.Background(), AccountSid, Sid).ByocTrunkSid(ByocTrunkSid).DomainName(DomainName).EmergencyCallerSid(EmergencyCallerSid).EmergencyCallingEnabled(EmergencyCallingEnabled).FriendlyName(FriendlyName).Secure(Secure).SipRegistration(SipRegistration).VoiceFallbackMethod(VoiceFallbackMethod).VoiceFallbackUrl(VoiceFallbackUrl).VoiceMethod(VoiceMethod).VoiceStatusCallbackMethod(VoiceStatusCallbackMethod).VoiceStatusCallbackUrl(VoiceStatusCallbackUrl).VoiceUrl(VoiceUrl).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateSipDomain``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSipDomain`: ApiV2010AccountSipSipDomain
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateSipDomain`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the SipDomain resource to update.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the SipDomain resource to update.

### Other Parameters

Other parameters are passed through a pointer to a UpdateSipDomainParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


 **ByocTrunkSid** | **string** | The SID of the BYOC Trunk(Bring Your Own Carrier) resource that the Sip Domain will be associated with.
 **DomainName** | **string** | The unique address you reserve on Twilio to which you route your SIP traffic. Domain names can contain letters, digits, and \\\&quot;-\\\&quot; and must end with &#x60;sip.twilio.com&#x60;.
 **EmergencyCallerSid** | **string** | Whether an emergency caller sid is configured for the domain. If present, this phone number will be used as the callback for the emergency call.
 **EmergencyCallingEnabled** | **bool** | Whether emergency calling is enabled for the domain. If enabled, allows emergency calls on the domain from phone numbers with validated addresses.
 **FriendlyName** | **string** | A descriptive string that you created to describe the resource. It can be up to 64 characters long.
 **Secure** | **bool** | Whether secure SIP is enabled for the domain. If enabled, TLS will be enforced and SRTP will be negotiated on all incoming calls to this sip domain.
 **SipRegistration** | **bool** | Whether to allow SIP Endpoints to register with the domain to receive calls. Can be &#x60;true&#x60; or &#x60;false&#x60;. &#x60;true&#x60; allows SIP Endpoints to register with the domain to receive calls, &#x60;false&#x60; does not.
 **VoiceFallbackMethod** | **string** | The HTTP method we should use to call &#x60;voice_fallback_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60;.
 **VoiceFallbackUrl** | **string** | The URL that we should call when an error occurs while retrieving or executing the TwiML requested by &#x60;voice_url&#x60;.
 **VoiceMethod** | **string** | The HTTP method we should use to call &#x60;voice_url&#x60;
 **VoiceStatusCallbackMethod** | **string** | The HTTP method we should use to call &#x60;voice_status_callback_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60;.
 **VoiceStatusCallbackUrl** | **string** | The URL that we should call to pass status parameters (such as call ended) to your application.
 **VoiceUrl** | **string** | The URL we should call when the domain receives a call.

### Return type

[**ApiV2010AccountSipSipDomain**](ApiV2010AccountSipSipDomain.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSipIpAccessControlList

> ApiV2010AccountSipSipIpAccessControlList UpdateSipIpAccessControlList(ctx, AccountSid, Sid).FriendlyName(FriendlyName).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AccountSid := "AccountSid_example" // string | The unique id of the [Account](https://www.twilio.com/docs/iam/api/account) responsible for this resource.
    Sid := "Sid_example" // string | A 34 character string that uniquely identifies the resource to udpate.
    FriendlyName := "FriendlyName_example" // string | A human readable descriptive text, up to 64 characters long. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateSipIpAccessControlList(context.Background(), AccountSid, Sid).FriendlyName(FriendlyName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateSipIpAccessControlList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSipIpAccessControlList`: ApiV2010AccountSipSipIpAccessControlList
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateSipIpAccessControlList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string** | The unique id of the [Account](https://www.twilio.com/docs/iam/api/account) responsible for this resource.
**Sid** | **string** | A 34 character string that uniquely identifies the resource to udpate.

### Other Parameters

Other parameters are passed through a pointer to a UpdateSipIpAccessControlListParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


 **FriendlyName** | **string** | A human readable descriptive text, up to 64 characters long.

### Return type

[**ApiV2010AccountSipSipIpAccessControlList**](ApiV2010AccountSipSipIpAccessControlList.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSipIpAddress

> ApiV2010AccountSipSipIpAccessControlListSipIpAddress UpdateSipIpAddress(ctx, AccountSid, IpAccessControlListSid, Sid).CidrPrefixLength(CidrPrefixLength).FriendlyName(FriendlyName).IpAddress(IpAddress).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AccountSid := "AccountSid_example" // string | The unique id of the [Account](https://www.twilio.com/docs/iam/api/account) responsible for this resource.
    IpAccessControlListSid := "IpAccessControlListSid_example" // string | The IpAccessControlList Sid that identifies the IpAddress resources to update.
    Sid := "Sid_example" // string | A 34 character string that identifies the IpAddress resource to update.
    CidrPrefixLength := int32(56) // int32 | An integer representing the length of the CIDR prefix to use with this IP address when accepting traffic. By default the entire IP address is used. (optional)
    FriendlyName := "FriendlyName_example" // string | A human readable descriptive text for this resource, up to 64 characters long. (optional)
    IpAddress := "IpAddress_example" // string | An IP address in dotted decimal notation from which you want to accept traffic. Any SIP requests from this IP address will be allowed by Twilio. IPv4 only supported today. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateSipIpAddress(context.Background(), AccountSid, IpAccessControlListSid, Sid).CidrPrefixLength(CidrPrefixLength).FriendlyName(FriendlyName).IpAddress(IpAddress).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateSipIpAddress``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSipIpAddress`: ApiV2010AccountSipSipIpAccessControlListSipIpAddress
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateSipIpAddress`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string** | The unique id of the [Account](https://www.twilio.com/docs/iam/api/account) responsible for this resource.
**IpAccessControlListSid** | **string** | The IpAccessControlList Sid that identifies the IpAddress resources to update.
**Sid** | **string** | A 34 character string that identifies the IpAddress resource to update.

### Other Parameters

Other parameters are passed through a pointer to a UpdateSipIpAddressParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------



 **CidrPrefixLength** | **int32** | An integer representing the length of the CIDR prefix to use with this IP address when accepting traffic. By default the entire IP address is used.
 **FriendlyName** | **string** | A human readable descriptive text for this resource, up to 64 characters long.
 **IpAddress** | **string** | An IP address in dotted decimal notation from which you want to accept traffic. Any SIP requests from this IP address will be allowed by Twilio. IPv4 only supported today.

### Return type

[**ApiV2010AccountSipSipIpAccessControlListSipIpAddress**](ApiV2010AccountSipSipIpAccessControlListSipIpAddress.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUsageTrigger

> ApiV2010AccountUsageUsageTrigger UpdateUsageTrigger(ctx, AccountSid, Sid).CallbackMethod(CallbackMethod).CallbackUrl(CallbackUrl).FriendlyName(FriendlyName).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    AccountSid := "AccountSid_example" // string | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the UsageTrigger resources to update.
    Sid := "Sid_example" // string | The Twilio-provided string that uniquely identifies the UsageTrigger resource to update.
    CallbackMethod := "CallbackMethod_example" // string | The HTTP method we should use to call `callback_url`. Can be: `GET` or `POST` and the default is `POST`. (optional)
    CallbackUrl := "CallbackUrl_example" // string | The URL we should call using `callback_method` when the trigger fires. (optional)
    FriendlyName := "FriendlyName_example" // string | A descriptive string that you create to describe the resource. It can be up to 64 characters long. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateUsageTrigger(context.Background(), AccountSid, Sid).CallbackMethod(CallbackMethod).CallbackUrl(CallbackUrl).FriendlyName(FriendlyName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateUsageTrigger``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateUsageTrigger`: ApiV2010AccountUsageUsageTrigger
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateUsageTrigger`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the UsageTrigger resources to update.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the UsageTrigger resource to update.

### Other Parameters

Other parameters are passed through a pointer to a UpdateUsageTriggerParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


 **CallbackMethod** | **string** | The HTTP method we should use to call &#x60;callback_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60; and the default is &#x60;POST&#x60;.
 **CallbackUrl** | **string** | The URL we should call using &#x60;callback_method&#x60; when the trigger fires.
 **FriendlyName** | **string** | A descriptive string that you create to describe the resource. It can be up to 64 characters long.

### Return type

[**ApiV2010AccountUsageUsageTrigger**](ApiV2010AccountUsageUsageTrigger.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

