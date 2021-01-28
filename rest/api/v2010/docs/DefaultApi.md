# \DefaultApi

All URIs are relative to *http://localhost*

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

> ApiV2010Account CreateAccount(ctx, optional)



Create a new Twilio Subaccount from the account making the request

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateAccountOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateAccountOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **FriendlyName** | **optional.String**| A human readable description of the account to create, defaults to &#x60;SubAccount Created at {YYYY-MM-DD HH:MM meridian}&#x60; | 

### Return type

[**ApiV2010Account**](api.v2010.account.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAddress

> ApiV2010AccountAddress CreateAddress(ctx, AccountSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string**| The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that will be responsible for the new Address resource. | 
 **optional** | ***CreateAddressOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateAddressOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **AutoCorrectAddress** | **optional.Bool**| Whether we should automatically correct the address. Can be: &#x60;true&#x60; or &#x60;false&#x60; and the default is &#x60;true&#x60;. If empty or &#x60;true&#x60;, we will correct the address you provide if necessary. If &#x60;false&#x60;, we won&#39;t alter the address you provide. | 
 **City** | **optional.String**| The city of the new address. | 
 **CustomerName** | **optional.String**| The name to associate with the new address. | 
 **EmergencyEnabled** | **optional.Bool**| Whether to enable emergency calling on the new address. Can be: &#x60;true&#x60; or &#x60;false&#x60;. | 
 **FriendlyName** | **optional.String**| A descriptive string that you create to describe the new address. It can be up to 64 characters long. | 
 **IsoCountry** | **optional.String**| The ISO country code of the new address. | 
 **PostalCode** | **optional.String**| The postal code of the new address. | 
 **Region** | **optional.String**| The state or region of the new address. | 
 **Street** | **optional.String**| The number and street address of the new address. | 

### Return type

[**ApiV2010AccountAddress**](api.v2010.account.address.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateApplication

> ApiV2010AccountApplication CreateApplication(ctx, AccountSid, optional)



Create a new application within your account

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string**| The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that will create the resource. | 
 **optional** | ***CreateApplicationOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateApplicationOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ApiVersion** | **optional.String**| The API version to use to start a new TwiML session. Can be: &#x60;2010-04-01&#x60; or &#x60;2008-08-01&#x60;. The default value is the account&#39;s default API version. | 
 **FriendlyName** | **optional.String**| A descriptive string that you create to describe the new application. It can be up to 64 characters long. | 
 **MessageStatusCallback** | **optional.String**| The URL we should call using a POST method to send message status information to your application. | 
 **SmsFallbackMethod** | **optional.String**| The HTTP method we should use to call &#x60;sms_fallback_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60;. | 
 **SmsFallbackUrl** | **optional.String**| The URL that we should call when an error occurs while retrieving or executing the TwiML from &#x60;sms_url&#x60;. | 
 **SmsMethod** | **optional.String**| The HTTP method we should use to call &#x60;sms_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60;. | 
 **SmsStatusCallback** | **optional.String**| The URL we should call using a POST method to send status information about SMS messages sent by the application. | 
 **SmsUrl** | **optional.String**| The URL we should call when the phone number receives an incoming SMS message. | 
 **StatusCallback** | **optional.String**| The URL we should call using the &#x60;status_callback_method&#x60; to send status information to your application. | 
 **StatusCallbackMethod** | **optional.String**| The HTTP method we should use to call &#x60;status_callback&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60;. | 
 **VoiceCallerIdLookup** | **optional.Bool**| Whether we should look up the caller&#39;s caller-ID name from the CNAM database (additional charges apply). Can be: &#x60;true&#x60; or &#x60;false&#x60;. | 
 **VoiceFallbackMethod** | **optional.String**| The HTTP method we should use to call &#x60;voice_fallback_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60;. | 
 **VoiceFallbackUrl** | **optional.String**| The URL that we should call when an error occurs retrieving or executing the TwiML requested by &#x60;url&#x60;. | 
 **VoiceMethod** | **optional.String**| The HTTP method we should use to call &#x60;voice_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60;. | 
 **VoiceUrl** | **optional.String**| The URL we should call when the phone number assigned to this application receives a call. | 

### Return type

[**ApiV2010AccountApplication**](api.v2010.account.application.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCall

> ApiV2010AccountCall CreateCall(ctx, AccountSid, optional)



Create a new outgoing call to phones, SIP-enabled endpoints or Twilio Client connections

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string**| The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that will create the resource. | 
 **optional** | ***CreateCallOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateCallOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ApplicationSid** | **optional.String**| The SID of the Application resource that will handle the call, if the call will be handled by an application. | 
 **AsyncAmd** | **optional.String**| Select whether to perform answering machine detection in the background. Default, blocks the execution of the call until Answering Machine Detection is completed. Can be: &#x60;true&#x60; or &#x60;false&#x60;. | 
 **AsyncAmdStatusCallback** | **optional.String**| The URL that we should call using the &#x60;async_amd_status_callback_method&#x60; to notify customer application whether the call was answered by human, machine or fax. | 
 **AsyncAmdStatusCallbackMethod** | **optional.String**| The HTTP method we should use when calling the &#x60;async_amd_status_callback&#x60; URL. Can be: &#x60;GET&#x60; or &#x60;POST&#x60; and the default is &#x60;POST&#x60;. | 
 **Byoc** | **optional.String**| The SID of a BYOC (Bring Your Own Carrier) trunk to route this call with. Note that &#x60;byoc&#x60; is only meaningful when &#x60;to&#x60; is a phone number; it will otherwise be ignored. (Beta) | 
 **CallReason** | **optional.String**| The Reason for the outgoing call. Use it to specify the purpose of the call that is presented on the called party&#39;s phone. (Branded Calls Beta) | 
 **CallerId** | **optional.String**| The phone number, SIP address, or Client identifier that made this call. Phone numbers are in [E.164 format](https://wwnw.twilio.com/docs/glossary/what-e164) (e.g., +16175551212). SIP addresses are formatted as &#x60;name@company.com&#x60;. | 
 **FallbackMethod** | **optional.String**| The HTTP method that we should use to request the &#x60;fallback_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60; and the default is &#x60;POST&#x60;. If an &#x60;application_sid&#x60; parameter is present, this parameter is ignored. | 
 **FallbackUrl** | **optional.String**| The URL that we call using the &#x60;fallback_method&#x60; if an error occurs when requesting or executing the TwiML at &#x60;url&#x60;. If an &#x60;application_sid&#x60; parameter is present, this parameter is ignored. | 
 **From** | **optional.String**| The phone number or client identifier to use as the caller id. If using a phone number, it must be a Twilio number or a Verified [outgoing caller id](https://www.twilio.com/docs/voice/api/outgoing-caller-ids) for your account. If the &#x60;to&#x60; parameter is a phone number, &#x60;From&#x60; must also be a phone number. | 
 **MachineDetection** | **optional.String**| Whether to detect if a human, answering machine, or fax has picked up the call. Can be: &#x60;Enable&#x60; or &#x60;DetectMessageEnd&#x60;. Use &#x60;Enable&#x60; if you would like us to return &#x60;AnsweredBy&#x60; as soon as the called party is identified. Use &#x60;DetectMessageEnd&#x60;, if you would like to leave a message on an answering machine. If &#x60;send_digits&#x60; is provided, this parameter is ignored. For more information, see [Answering Machine Detection](https://www.twilio.com/docs/voice/answering-machine-detection). | 
 **MachineDetectionSilenceTimeout** | **optional.Int32**| The number of milliseconds of initial silence after which an &#x60;unknown&#x60; AnsweredBy result will be returned. Possible Values: 2000-10000. Default: 5000. | 
 **MachineDetectionSpeechEndThreshold** | **optional.Int32**| The number of milliseconds of silence after speech activity at which point the speech activity is considered complete. Possible Values: 500-5000. Default: 1200. | 
 **MachineDetectionSpeechThreshold** | **optional.Int32**| The number of milliseconds that is used as the measuring stick for the length of the speech activity, where durations lower than this value will be interpreted as a human and longer than this value as a machine. Possible Values: 1000-6000. Default: 2400. | 
 **MachineDetectionTimeout** | **optional.Int32**| The number of seconds that we should attempt to detect an answering machine before timing out and sending a voice request with &#x60;AnsweredBy&#x60; of &#x60;unknown&#x60;. The default timeout is 30 seconds. | 
 **Method** | **optional.String**| The HTTP method we should use when calling the &#x60;url&#x60; parameter&#39;s value. Can be: &#x60;GET&#x60; or &#x60;POST&#x60; and the default is &#x60;POST&#x60;. If an &#x60;application_sid&#x60; parameter is present, this parameter is ignored. | 
 **Record** | **optional.Bool**| Whether to record the call. Can be &#x60;true&#x60; to record the phone call, or &#x60;false&#x60; to not. The default is &#x60;false&#x60;. The &#x60;recording_url&#x60; is sent to the &#x60;status_callback&#x60; URL. | 
 **RecordingChannels** | **optional.String**| The number of channels in the final recording. Can be: &#x60;mono&#x60; or &#x60;dual&#x60;. The default is &#x60;mono&#x60;. &#x60;mono&#x60; records both legs of the call in a single channel of the recording file. &#x60;dual&#x60; records each leg to a separate channel of the recording file. The first channel of a dual-channel recording contains the parent call and the second channel contains the child call. | 
 **RecordingStatusCallback** | **optional.String**| The URL that we call when the recording is available to be accessed. | 
 **RecordingStatusCallbackEvent** | [**optional.Interface of []string**](string.md)| The recording status events that will trigger calls to the URL specified in &#x60;recording_status_callback&#x60;. Can be: &#x60;in-progress&#x60;, &#x60;completed&#x60; and &#x60;absent&#x60;. Defaults to &#x60;completed&#x60;. Separate  multiple values with a space. | 
 **RecordingStatusCallbackMethod** | **optional.String**| The HTTP method we should use when calling the &#x60;recording_status_callback&#x60; URL. Can be: &#x60;GET&#x60; or &#x60;POST&#x60; and the default is &#x60;POST&#x60;. | 
 **RecordingTrack** | **optional.String**| The audio track to record for the call. Can be: &#x60;inbound&#x60;, &#x60;outbound&#x60; or &#x60;both&#x60;. The default is &#x60;both&#x60;. &#x60;inbound&#x60; records the audio that is received by Twilio. &#x60;outbound&#x60; records the audio that is generated from Twilio. &#x60;both&#x60; records the audio that is received and generated by Twilio. | 
 **SendDigits** | **optional.String**| A string of keys to dial after connecting to the number, maximum of 32 digits. Valid digits in the string include: any digit (&#x60;0&#x60;-&#x60;9&#x60;), &#39;&#x60;#&#x60;&#39;, &#39;&#x60;*&#x60;&#39; and &#39;&#x60;w&#x60;&#39;, to insert a half second pause. For example, if you connected to a company phone number and wanted to pause for one second, and then dial extension 1234 followed by the pound key, the value of this parameter would be &#x60;ww1234#&#x60;. Remember to URL-encode this string, since the &#39;&#x60;#&#x60;&#39; character has special meaning in a URL. If both &#x60;SendDigits&#x60; and &#x60;MachineDetection&#x60; parameters are provided, then &#x60;MachineDetection&#x60; will be ignored. | 
 **SipAuthPassword** | **optional.String**| The password required to authenticate the user account specified in &#x60;sip_auth_username&#x60;. | 
 **SipAuthUsername** | **optional.String**| The username used to authenticate the caller making a SIP call. | 
 **StatusCallback** | **optional.String**| The URL we should call using the &#x60;status_callback_method&#x60; to send status information to your application. If no &#x60;status_callback_event&#x60; is specified, we will send the &#x60;completed&#x60; status. If an &#x60;application_sid&#x60; parameter is present, this parameter is ignored. URLs must contain a valid hostname (underscores are not permitted). | 
 **StatusCallbackEvent** | [**optional.Interface of []string**](string.md)| The call progress events that we will send to the &#x60;status_callback&#x60; URL. Can be: &#x60;initiated&#x60;, &#x60;ringing&#x60;, &#x60;answered&#x60;, and &#x60;completed&#x60;. If no event is specified, we send the &#x60;completed&#x60; status. If you want to receive multiple events, specify each one in a separate &#x60;status_callback_event&#x60; parameter. See the code sample for [monitoring call progress](https://www.twilio.com/docs/voice/api/call-resource?code-sample&#x3D;code-create-a-call-resource-and-specify-a-statuscallbackevent&amp;code-sdk-version&#x3D;json). If an &#x60;application_sid&#x60; is present, this parameter is ignored. | 
 **StatusCallbackMethod** | **optional.String**| The HTTP method we should use when calling the &#x60;status_callback&#x60; URL. Can be: &#x60;GET&#x60; or &#x60;POST&#x60; and the default is &#x60;POST&#x60;. If an &#x60;application_sid&#x60; parameter is present, this parameter is ignored. | 
 **Timeout** | **optional.Int32**| The integer number of seconds that we should allow the phone to ring before assuming there is no answer. The default is &#x60;60&#x60; seconds and the maximum is &#x60;600&#x60; seconds. For some call flows, we will add a 5-second buffer to the timeout value you provide. For this reason, a timeout value of 10 seconds could result in an actual timeout closer to 15 seconds. You can set this to a short time, such as &#x60;15&#x60; seconds, to hang up before reaching an answering machine or voicemail. | 
 **To** | **optional.String**| The phone number, SIP address, or client identifier to call. | 
 **Trim** | **optional.String**| Whether to trim any leading and trailing silence from the recording. Can be: &#x60;trim-silence&#x60; or &#x60;do-not-trim&#x60; and the default is &#x60;trim-silence&#x60;. | 
 **Twiml** | **optional.String**| TwiML instructions for the call Twilio will use without fetching Twiml from url parameter. If both &#x60;twiml&#x60; and &#x60;url&#x60; are provided then &#x60;twiml&#x60; parameter will be ignored. | 
 **Url** | **optional.String**| The absolute URL that returns the TwiML instructions for the call. We will call this URL using the &#x60;method&#x60; when the call connects. For more information, see the [Url Parameter](https://www.twilio.com/docs/voice/make-calls#specify-a-url-parameter) section in [Making Calls](https://www.twilio.com/docs/voice/make-calls). | 

### Return type

[**ApiV2010AccountCall**](api.v2010.account.call.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCallFeedbackSummary

> ApiV2010AccountCallCallFeedbackSummary CreateCallFeedbackSummary(ctx, AccountSid, optional)



Create a FeedbackSummary resource for a call

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string**| The unique id of the [Account](https://www.twilio.com/docs/iam/api/account) responsible for this resource. | 
 **optional** | ***CreateCallFeedbackSummaryOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateCallFeedbackSummaryOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **EndDate** | **optional.Time**| Only include feedback given on or before this date. Format is &#x60;YYYY-MM-DD&#x60; and specified in UTC. | 
 **IncludeSubaccounts** | **optional.Bool**| Whether to also include Feedback resources from all subaccounts. &#x60;true&#x60; includes feedback from all subaccounts and &#x60;false&#x60;, the default, includes feedback from only the specified account. | 
 **StartDate** | **optional.Time**| Only include feedback given on or after this date. Format is &#x60;YYYY-MM-DD&#x60; and specified in UTC. | 
 **StatusCallback** | **optional.String**| The URL that we will request when the feedback summary is complete. | 
 **StatusCallbackMethod** | **optional.String**| The HTTP method (&#x60;GET&#x60; or &#x60;POST&#x60;) we use to make the request to the &#x60;StatusCallback&#x60; URL. | 

### Return type

[**ApiV2010AccountCallCallFeedbackSummary**](api.v2010.account.call.call_feedback_summary.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCallRecording

> ApiV2010AccountCallCallRecording CreateCallRecording(ctx, AccountSid, CallSid, optional)



Create a recording for the call

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string**| The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that will create the resource. | 
**CallSid** | **string**| The SID of the [Call](https://www.twilio.com/docs/voice/api/call-resource) to associate the resource with. | 
 **optional** | ***CreateCallRecordingOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateCallRecordingOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **RecordingChannels** | **optional.String**| The number of channels used in the recording. Can be: &#x60;mono&#x60; or &#x60;dual&#x60; and the default is &#x60;mono&#x60;. &#x60;mono&#x60; records all parties of the call into one channel. &#x60;dual&#x60; records each party of a 2-party call into separate channels. | 
 **RecordingStatusCallback** | **optional.String**| The URL we should call using the &#x60;recording_status_callback_method&#x60; on each recording event specified in  &#x60;recording_status_callback_event&#x60;. For more information, see [RecordingStatusCallback parameters](https://www.twilio.com/docs/voice/api/recording#recordingstatuscallback). | 
 **RecordingStatusCallbackEvent** | [**optional.Interface of []string**](string.md)| The recording status events on which we should call the &#x60;recording_status_callback&#x60; URL. Can be: &#x60;in-progress&#x60;, &#x60;completed&#x60; and &#x60;absent&#x60; and the default is &#x60;completed&#x60;. Separate multiple event values with a space. | 
 **RecordingStatusCallbackMethod** | **optional.String**| The HTTP method we should use to call &#x60;recording_status_callback&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60; and the default is &#x60;POST&#x60;. | 
 **RecordingTrack** | **optional.String**| The audio track to record for the call. Can be: &#x60;inbound&#x60;, &#x60;outbound&#x60; or &#x60;both&#x60;. The default is &#x60;both&#x60;. &#x60;inbound&#x60; records the audio that is received by Twilio. &#x60;outbound&#x60; records the audio that is generated from Twilio. &#x60;both&#x60; records the audio that is received and generated by Twilio. | 
 **Trim** | **optional.String**| Whether to trim any leading and trailing silence in the recording. Can be: &#x60;trim-silence&#x60; or &#x60;do-not-trim&#x60; and the default is &#x60;do-not-trim&#x60;. &#x60;trim-silence&#x60; trims the silence from the beginning and end of the recording and &#x60;do-not-trim&#x60; does not. | 

### Return type

[**ApiV2010AccountCallCallRecording**](api.v2010.account.call.call_recording.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateIncomingPhoneNumber

> ApiV2010AccountIncomingPhoneNumber CreateIncomingPhoneNumber(ctx, AccountSid, optional)



Purchase a phone-number for the account.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string**| The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that will create the resource. | 
 **optional** | ***CreateIncomingPhoneNumberOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateIncomingPhoneNumberOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **AddressSid** | **optional.String**| The SID of the Address resource we should associate with the new phone number. Some regions require addresses to meet local regulations. | 
 **ApiVersion** | **optional.String**| The API version to use for incoming calls made to the new phone number. The default is &#x60;2010-04-01&#x60;. | 
 **AreaCode** | **optional.String**| The desired area code for your new incoming phone number. Can be any three-digit, US or Canada area code. We will provision an available phone number within this area code for you. **You must provide an &#x60;area_code&#x60; or a &#x60;phone_number&#x60;.** (US and Canada only). | 
 **BundleSid** | **optional.String**| The SID of the Bundle resource that you associate with the phone number. Some regions require a Bundle to meet local Regulations. | 
 **EmergencyAddressSid** | **optional.String**| The SID of the emergency address configuration to use for emergency calling from the new phone number. | 
 **EmergencyStatus** | **optional.String**| The configuration status parameter that determines whether the new phone number is enabled for emergency calling. | 
 **FriendlyName** | **optional.String**| A descriptive string that you created to describe the new phone number. It can be up to 64 characters long. By default, this is a formatted version of the new phone number. | 
 **IdentitySid** | **optional.String**| The SID of the Identity resource that we should associate with the new phone number. Some regions require an identity to meet local regulations. | 
 **PhoneNumber** | **optional.String**| The phone number to purchase specified in [E.164](https://www.twilio.com/docs/glossary/what-e164) format.  E.164 phone numbers consist of a + followed by the country code and subscriber number without punctuation characters. For example, +14155551234. | 
 **SmsApplicationSid** | **optional.String**| The SID of the application that should handle SMS messages sent to the new phone number. If an &#x60;sms_application_sid&#x60; is present, we ignore all of the &#x60;sms_*_url&#x60; urls and use those set on the application. | 
 **SmsFallbackMethod** | **optional.String**| The HTTP method that we should use to call &#x60;sms_fallback_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60; and defaults to &#x60;POST&#x60;. | 
 **SmsFallbackUrl** | **optional.String**| The URL that we should call when an error occurs while requesting or executing the TwiML defined by &#x60;sms_url&#x60;. | 
 **SmsMethod** | **optional.String**| The HTTP method that we should use to call &#x60;sms_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60; and defaults to &#x60;POST&#x60;. | 
 **SmsUrl** | **optional.String**| The URL we should call when the new phone number receives an incoming SMS message. | 
 **StatusCallback** | **optional.String**| The URL we should call using the &#x60;status_callback_method&#x60; to send status information to your application. | 
 **StatusCallbackMethod** | **optional.String**| The HTTP method we should use to call &#x60;status_callback&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60; and defaults to &#x60;POST&#x60;. | 
 **TrunkSid** | **optional.String**| The SID of the Trunk we should use to handle calls to the new phone number. If a &#x60;trunk_sid&#x60; is present, we ignore all of the voice urls and voice applications and use only those set on the Trunk. Setting a &#x60;trunk_sid&#x60; will automatically delete your &#x60;voice_application_sid&#x60; and vice versa. | 
 **VoiceApplicationSid** | **optional.String**| The SID of the application we should use to handle calls to the new phone number. If a &#x60;voice_application_sid&#x60; is present, we ignore all of the voice urls and use only those set on the application. Setting a &#x60;voice_application_sid&#x60; will automatically delete your &#x60;trunk_sid&#x60; and vice versa. | 
 **VoiceCallerIdLookup** | **optional.Bool**| Whether to lookup the caller&#39;s name from the CNAM database and post it to your app. Can be: &#x60;true&#x60; or &#x60;false&#x60; and defaults to &#x60;false&#x60;. | 
 **VoiceFallbackMethod** | **optional.String**| The HTTP method that we should use to call &#x60;voice_fallback_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60; and defaults to &#x60;POST&#x60;. | 
 **VoiceFallbackUrl** | **optional.String**| The URL that we should call when an error occurs retrieving or executing the TwiML requested by &#x60;url&#x60;. | 
 **VoiceMethod** | **optional.String**| The HTTP method that we should use to call &#x60;voice_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60; and defaults to &#x60;POST&#x60;. | 
 **VoiceReceiveMode** | **optional.String**| The configuration parameter for the new phone number to receive incoming voice calls or faxes. Can be: &#x60;fax&#x60; or &#x60;voice&#x60; and defaults to &#x60;voice&#x60;. | 
 **VoiceUrl** | **optional.String**| The URL that we should call to answer a call to the new phone number. The &#x60;voice_url&#x60; will not be called if a &#x60;voice_application_sid&#x60; or a &#x60;trunk_sid&#x60; is set. | 

### Return type

[**ApiV2010AccountIncomingPhoneNumber**](api.v2010.account.incoming_phone_number.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateIncomingPhoneNumberAssignedAddOn

> ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberAssignedAddOn CreateIncomingPhoneNumberAssignedAddOn(ctx, AccountSid, ResourceSid, optional)



Assign an Add-on installation to the Number specified.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string**| The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that will create the resource. | 
**ResourceSid** | **string**| The SID of the Phone Number to assign the Add-on. | 
 **optional** | ***CreateIncomingPhoneNumberAssignedAddOnOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateIncomingPhoneNumberAssignedAddOnOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **InstalledAddOnSid** | **optional.String**| The SID that identifies the Add-on installation. | 

### Return type

[**ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberAssignedAddOn**](api.v2010.account.incoming_phone_number.incoming_phone_number_assigned_add_on.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateIncomingPhoneNumberLocal

> ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberLocal CreateIncomingPhoneNumberLocal(ctx, AccountSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string**| The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that will create the resource. | 
 **optional** | ***CreateIncomingPhoneNumberLocalOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateIncomingPhoneNumberLocalOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **AddressSid** | **optional.String**| The SID of the Address resource we should associate with the new phone number. Some regions require addresses to meet local regulations. | 
 **ApiVersion** | **optional.String**| The API version to use for incoming calls made to the new phone number. The default is &#x60;2010-04-01&#x60;. | 
 **BundleSid** | **optional.String**| The SID of the Bundle resource that you associate with the phone number. Some regions require a Bundle to meet local Regulations. | 
 **EmergencyAddressSid** | **optional.String**| The SID of the emergency address configuration to use for emergency calling from the new phone number. | 
 **EmergencyStatus** | **optional.String**| The configuration status parameter that determines whether the new phone number is enabled for emergency calling. | 
 **FriendlyName** | **optional.String**| A descriptive string that you created to describe the new phone number. It can be up to 64 characters long. By default, this is a formatted version of the phone number. | 
 **IdentitySid** | **optional.String**| The SID of the Identity resource that we should associate with the new phone number. Some regions require an identity to meet local regulations. | 
 **PhoneNumber** | **optional.String**| The phone number to purchase specified in [E.164](https://www.twilio.com/docs/glossary/what-e164) format.  E.164 phone numbers consist of a + followed by the country code and subscriber number without punctuation characters. For example, +14155551234. | 
 **SmsApplicationSid** | **optional.String**| The SID of the application that should handle SMS messages sent to the new phone number. If an &#x60;sms_application_sid&#x60; is present, we ignore all of the &#x60;sms_*_url&#x60; urls and use those set on the application. | 
 **SmsFallbackMethod** | **optional.String**| The HTTP method that we should use to call &#x60;sms_fallback_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60; and defaults to &#x60;POST&#x60;. | 
 **SmsFallbackUrl** | **optional.String**| The URL that we should call when an error occurs while requesting or executing the TwiML defined by &#x60;sms_url&#x60;. | 
 **SmsMethod** | **optional.String**| The HTTP method that we should use to call &#x60;sms_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60; and defaults to &#x60;POST&#x60;. | 
 **SmsUrl** | **optional.String**| The URL we should call when the new phone number receives an incoming SMS message. | 
 **StatusCallback** | **optional.String**| The URL we should call using the &#x60;status_callback_method&#x60; to send status information to your application. | 
 **StatusCallbackMethod** | **optional.String**| The HTTP method we should use to call &#x60;status_callback&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60; and defaults to &#x60;POST&#x60;. | 
 **TrunkSid** | **optional.String**| The SID of the Trunk we should use to handle calls to the new phone number. If a &#x60;trunk_sid&#x60; is present, we ignore all of the voice urls and voice applications and use only those set on the Trunk. Setting a &#x60;trunk_sid&#x60; will automatically delete your &#x60;voice_application_sid&#x60; and vice versa. | 
 **VoiceApplicationSid** | **optional.String**| The SID of the application we should use to handle calls to the new phone number. If a &#x60;voice_application_sid&#x60; is present, we ignore all of the voice urls and use only those set on the application. Setting a &#x60;voice_application_sid&#x60; will automatically delete your &#x60;trunk_sid&#x60; and vice versa. | 
 **VoiceCallerIdLookup** | **optional.Bool**| Whether to lookup the caller&#39;s name from the CNAM database and post it to your app. Can be: &#x60;true&#x60; or &#x60;false&#x60; and defaults to &#x60;false&#x60;. | 
 **VoiceFallbackMethod** | **optional.String**| The HTTP method that we should use to call &#x60;voice_fallback_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60; and defaults to &#x60;POST&#x60;. | 
 **VoiceFallbackUrl** | **optional.String**| The URL that we should call when an error occurs retrieving or executing the TwiML requested by &#x60;url&#x60;. | 
 **VoiceMethod** | **optional.String**| The HTTP method that we should use to call &#x60;voice_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60; and defaults to &#x60;POST&#x60;. | 
 **VoiceReceiveMode** | **optional.String**| The configuration parameter for the new phone number to receive incoming voice calls or faxes. Can be: &#x60;fax&#x60; or &#x60;voice&#x60; and defaults to &#x60;voice&#x60;. | 
 **VoiceUrl** | **optional.String**| The URL that we should call to answer a call to the new phone number. The &#x60;voice_url&#x60; will not be called if a &#x60;voice_application_sid&#x60; or a &#x60;trunk_sid&#x60; is set. | 

### Return type

[**ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberLocal**](api.v2010.account.incoming_phone_number.incoming_phone_number_local.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateIncomingPhoneNumberMobile

> ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile CreateIncomingPhoneNumberMobile(ctx, AccountSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string**| The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that will create the resource. | 
 **optional** | ***CreateIncomingPhoneNumberMobileOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateIncomingPhoneNumberMobileOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **AddressSid** | **optional.String**| The SID of the Address resource we should associate with the new phone number. Some regions require addresses to meet local regulations. | 
 **ApiVersion** | **optional.String**| The API version to use for incoming calls made to the new phone number. The default is &#x60;2010-04-01&#x60;. | 
 **BundleSid** | **optional.String**| The SID of the Bundle resource that you associate with the phone number. Some regions require a Bundle to meet local Regulations. | 
 **EmergencyAddressSid** | **optional.String**| The SID of the emergency address configuration to use for emergency calling from the new phone number. | 
 **EmergencyStatus** | **optional.String**| The configuration status parameter that determines whether the new phone number is enabled for emergency calling. | 
 **FriendlyName** | **optional.String**| A descriptive string that you created to describe the new phone number. It can be up to 64 characters long. By default, the is a formatted version of the phone number. | 
 **IdentitySid** | **optional.String**| The SID of the Identity resource that we should associate with the new phone number. Some regions require an identity to meet local regulations. | 
 **PhoneNumber** | **optional.String**| The phone number to purchase specified in [E.164](https://www.twilio.com/docs/glossary/what-e164) format.  E.164 phone numbers consist of a + followed by the country code and subscriber number without punctuation characters. For example, +14155551234. | 
 **SmsApplicationSid** | **optional.String**| The SID of the application that should handle SMS messages sent to the new phone number. If an &#x60;sms_application_sid&#x60; is present, we ignore all of the &#x60;sms_*_url&#x60; urls and use those of the application. | 
 **SmsFallbackMethod** | **optional.String**| The HTTP method that we should use to call &#x60;sms_fallback_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60; and defaults to &#x60;POST&#x60;. | 
 **SmsFallbackUrl** | **optional.String**| The URL that we should call when an error occurs while requesting or executing the TwiML defined by &#x60;sms_url&#x60;. | 
 **SmsMethod** | **optional.String**| The HTTP method that we should use to call &#x60;sms_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60; and defaults to &#x60;POST&#x60;. | 
 **SmsUrl** | **optional.String**| The URL we should call when the new phone number receives an incoming SMS message. | 
 **StatusCallback** | **optional.String**| The URL we should call using the &#x60;status_callback_method&#x60; to send status information to your application. | 
 **StatusCallbackMethod** | **optional.String**| The HTTP method we should use to call &#x60;status_callback&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60; and defaults to &#x60;POST&#x60;. | 
 **TrunkSid** | **optional.String**| The SID of the Trunk we should use to handle calls to the new phone number. If a &#x60;trunk_sid&#x60; is present, we ignore all of the voice urls and voice applications and use only those set on the Trunk. Setting a &#x60;trunk_sid&#x60; will automatically delete your &#x60;voice_application_sid&#x60; and vice versa. | 
 **VoiceApplicationSid** | **optional.String**| The SID of the application we should use to handle calls to the new phone number. If a &#x60;voice_application_sid&#x60; is present, we ignore all of the voice urls and use only those set on the application. Setting a &#x60;voice_application_sid&#x60; will automatically delete your &#x60;trunk_sid&#x60; and vice versa. | 
 **VoiceCallerIdLookup** | **optional.Bool**| Whether to lookup the caller&#39;s name from the CNAM database and post it to your app. Can be: &#x60;true&#x60; or &#x60;false&#x60; and defaults to &#x60;false&#x60;. | 
 **VoiceFallbackMethod** | **optional.String**| The HTTP method that we should use to call &#x60;voice_fallback_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60; and defaults to &#x60;POST&#x60;. | 
 **VoiceFallbackUrl** | **optional.String**| The URL that we should call when an error occurs retrieving or executing the TwiML requested by &#x60;url&#x60;. | 
 **VoiceMethod** | **optional.String**| The HTTP method that we should use to call &#x60;voice_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60; and defaults to &#x60;POST&#x60;. | 
 **VoiceReceiveMode** | **optional.String**| The configuration parameter for the new phone number to receive incoming voice calls or faxes. Can be: &#x60;fax&#x60; or &#x60;voice&#x60; and defaults to &#x60;voice&#x60;. | 
 **VoiceUrl** | **optional.String**| The URL that we should call to answer a call to the new phone number. The &#x60;voice_url&#x60; will not be called if a &#x60;voice_application_sid&#x60; or a &#x60;trunk_sid&#x60; is set. | 

### Return type

[**ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile**](api.v2010.account.incoming_phone_number.incoming_phone_number_mobile.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateIncomingPhoneNumberTollFree

> ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree CreateIncomingPhoneNumberTollFree(ctx, AccountSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string**| The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that will create the resource. | 
 **optional** | ***CreateIncomingPhoneNumberTollFreeOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateIncomingPhoneNumberTollFreeOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **AddressSid** | **optional.String**| The SID of the Address resource we should associate with the new phone number. Some regions require addresses to meet local regulations. | 
 **ApiVersion** | **optional.String**| The API version to use for incoming calls made to the new phone number. The default is &#x60;2010-04-01&#x60;. | 
 **BundleSid** | **optional.String**| The SID of the Bundle resource that you associate with the phone number. Some regions require a Bundle to meet local Regulations. | 
 **EmergencyAddressSid** | **optional.String**| The SID of the emergency address configuration to use for emergency calling from the new phone number. | 
 **EmergencyStatus** | **optional.String**| The configuration status parameter that determines whether the new phone number is enabled for emergency calling. | 
 **FriendlyName** | **optional.String**| A descriptive string that you created to describe the new phone number. It can be up to 64 characters long. By default, this is a formatted version of the phone number. | 
 **IdentitySid** | **optional.String**| The SID of the Identity resource that we should associate with the new phone number. Some regions require an Identity to meet local regulations. | 
 **PhoneNumber** | **optional.String**| The phone number to purchase specified in [E.164](https://www.twilio.com/docs/glossary/what-e164) format.  E.164 phone numbers consist of a + followed by the country code and subscriber number without punctuation characters. For example, +14155551234. | 
 **SmsApplicationSid** | **optional.String**| The SID of the application that should handle SMS messages sent to the new phone number. If an &#x60;sms_application_sid&#x60; is present, we ignore all &#x60;sms_*_url&#x60; values and use those of the application. | 
 **SmsFallbackMethod** | **optional.String**| The HTTP method that we should use to call &#x60;sms_fallback_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60; and defaults to &#x60;POST&#x60;. | 
 **SmsFallbackUrl** | **optional.String**| The URL that we should call when an error occurs while requesting or executing the TwiML defined by &#x60;sms_url&#x60;. | 
 **SmsMethod** | **optional.String**| The HTTP method that we should use to call &#x60;sms_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60; and defaults to &#x60;POST&#x60;. | 
 **SmsUrl** | **optional.String**| The URL we should call when the new phone number receives an incoming SMS message. | 
 **StatusCallback** | **optional.String**| The URL we should call using the &#x60;status_callback_method&#x60; to send status information to your application. | 
 **StatusCallbackMethod** | **optional.String**| The HTTP method we should use to call &#x60;status_callback&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60; and defaults to &#x60;POST&#x60;. | 
 **TrunkSid** | **optional.String**| The SID of the Trunk we should use to handle calls to the new phone number. If a &#x60;trunk_sid&#x60; is present, we ignore all of the voice urls and voice applications and use only those set on the Trunk. Setting a &#x60;trunk_sid&#x60; will automatically delete your &#x60;voice_application_sid&#x60; and vice versa. | 
 **VoiceApplicationSid** | **optional.String**| The SID of the application we should use to handle calls to the new phone number. If a &#x60;voice_application_sid&#x60; is present, we ignore all of the voice urls and use those set on the application. Setting a &#x60;voice_application_sid&#x60; will automatically delete your &#x60;trunk_sid&#x60; and vice versa. | 
 **VoiceCallerIdLookup** | **optional.Bool**| Whether to lookup the caller&#39;s name from the CNAM database and post it to your app. Can be: &#x60;true&#x60; or &#x60;false&#x60; and defaults to &#x60;false&#x60;. | 
 **VoiceFallbackMethod** | **optional.String**| The HTTP method that we should use to call &#x60;voice_fallback_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60; and defaults to &#x60;POST&#x60;. | 
 **VoiceFallbackUrl** | **optional.String**| The URL that we should call when an error occurs retrieving or executing the TwiML requested by &#x60;url&#x60;. | 
 **VoiceMethod** | **optional.String**| The HTTP method that we should use to call &#x60;voice_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60; and defaults to &#x60;POST&#x60;. | 
 **VoiceReceiveMode** | **optional.String**| The configuration parameter for the new phone number to receive incoming voice calls or faxes. Can be: &#x60;fax&#x60; or &#x60;voice&#x60; and defaults to &#x60;voice&#x60;. | 
 **VoiceUrl** | **optional.String**| The URL that we should call to answer a call to the new phone number. The &#x60;voice_url&#x60; will not be called if a &#x60;voice_application_sid&#x60; or a &#x60;trunk_sid&#x60; is set. | 

### Return type

[**ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree**](api.v2010.account.incoming_phone_number.incoming_phone_number_toll_free.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateMessage

> ApiV2010AccountMessage CreateMessage(ctx, AccountSid, optional)



Send a message from the account used to make the request

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string**| The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that will create the resource. | 
 **optional** | ***CreateMessageOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateMessageOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **AddressRetention** | **optional.String**| Determines if the address can be stored or obfuscated based on privacy settings | 
 **ApplicationSid** | **optional.String**| The SID of the application that should receive message status. We POST a &#x60;message_sid&#x60; parameter and a &#x60;message_status&#x60; parameter with a value of &#x60;sent&#x60; or &#x60;failed&#x60; to the [application](https://www.twilio.com/docs/usage/api/applications)&#39;s &#x60;message_status_callback&#x60;. If a &#x60;status_callback&#x60; parameter is also passed, it will be ignored and the application&#39;s &#x60;message_status_callback&#x60; parameter will be used. | 
 **Attempt** | **optional.Int32**| Total number of attempts made ( including this ) to send out the message regardless of the provider used | 
 **Body** | **optional.String**| The text of the message you want to send. Can be up to 1,600 characters in length. | 
 **ContentRetention** | **optional.String**| Determines if the message content can be stored or redacted based on privacy settings | 
 **ForceDelivery** | **optional.Bool**| Reserved | 
 **From** | **optional.String**| A Twilio phone number in [E.164](https://www.twilio.com/docs/glossary/what-e164) format, an [alphanumeric sender ID](https://www.twilio.com/docs/sms/send-messages#use-an-alphanumeric-sender-id), or a [Channel Endpoint address](https://www.twilio.com/docs/sms/channels#channel-addresses) that is enabled for the type of message you want to send. Phone numbers or [short codes](https://www.twilio.com/docs/sms/api/short-code) purchased from Twilio also work here. You cannot, for example, spoof messages from a private cell phone number. If you are using &#x60;messaging_service_sid&#x60;, this parameter must be empty. | 
 **MaxPrice** | **optional.Float32**| The maximum total price in US dollars that you will pay for the message to be delivered. Can be a decimal value that has up to 4 decimal places. All messages are queued for delivery and the message cost is checked before the message is sent. If the cost exceeds &#x60;max_price&#x60;, the message will fail and a status of &#x60;Failed&#x60; is sent to the status callback. If &#x60;MaxPrice&#x60; is not set, the message cost is not checked. | 
 **MediaUrl** | [**optional.Interface of []string**](string.md)| The URL of the media to send with the message. The media can be of type &#x60;gif&#x60;, &#x60;png&#x60;, and &#x60;jpeg&#x60; and will be formatted correctly on the recipient&#39;s device. The media size limit is 5MB for supported file types (JPEG, PNG, GIF) and 500KB for [other types](https://www.twilio.com/docs/sms/accepted-mime-types) of accepted media. To send more than one image in the message body, provide multiple &#x60;media_url&#x60; parameters in the POST request. You can include up to 10 &#x60;media_url&#x60; parameters per message. You can send images in an SMS message in only the US and Canada. | 
 **MessagingServiceSid** | **optional.String**| The SID of the [Messaging Service](https://www.twilio.com/docs/sms/services#send-a-message-with-copilot) you want to associate with the Message. Set this parameter to use the [Messaging Service Settings and Copilot Features](https://www.twilio.com/console/sms/services) you have configured and leave the &#x60;from&#x60; parameter empty. When only this parameter is set, Twilio will use your enabled Copilot Features to select the &#x60;from&#x60; phone number for delivery. | 
 **PersistentAction** | [**optional.Interface of []string**](string.md)| Rich actions for Channels Messages. | 
 **ProvideFeedback** | **optional.Bool**| Whether to confirm delivery of the message. Set this value to &#x60;true&#x60; if you are sending messages that have a trackable user action and you intend to confirm delivery of the message using the [Message Feedback API](https://www.twilio.com/docs/sms/api/message-feedback-resource). This parameter is &#x60;false&#x60; by default. | 
 **SmartEncoded** | **optional.Bool**| Whether to detect Unicode characters that have a similar GSM-7 character and replace them. Can be: &#x60;true&#x60; or &#x60;false&#x60;. | 
 **StatusCallback** | **optional.String**| The URL we should call using the &#x60;status_callback_method&#x60; to send status information to your application. If specified, we POST these message status changes to the URL: &#x60;queued&#x60;, &#x60;failed&#x60;, &#x60;sent&#x60;, &#x60;delivered&#x60;, or &#x60;undelivered&#x60;. Twilio will POST its [standard request parameters](https://www.twilio.com/docs/sms/twiml#request-parameters) as well as some additional parameters including &#x60;MessageSid&#x60;, &#x60;MessageStatus&#x60;, and &#x60;ErrorCode&#x60;. If you include this parameter with the &#x60;messaging_service_sid&#x60;, we use this URL instead of the Status Callback URL of the [Messaging Service](https://www.twilio.com/docs/sms/services/api). URLs must contain a valid hostname and underscores are not allowed. | 
 **To** | **optional.String**| The destination phone number in [E.164](https://www.twilio.com/docs/glossary/what-e164) format for SMS/MMS or [Channel user address](https://www.twilio.com/docs/sms/channels#channel-addresses) for other 3rd-party channels. | 
 **ValidityPeriod** | **optional.Int32**| How long in seconds the message can remain in our outgoing message queue. After this period elapses, the message fails and we call your status callback. Can be between 1 and the default value of 14,400 seconds. After a message has been accepted by a carrier, however, we cannot guarantee that the message will not be queued after this period. We recommend that this value be at least 5 seconds. | 

### Return type

[**ApiV2010AccountMessage**](api.v2010.account.message.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateMessageFeedback

> ApiV2010AccountMessageMessageFeedback CreateMessageFeedback(ctx, AccountSid, MessageSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string**| The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that will create the resource. | 
**MessageSid** | **string**| The SID of the Message resource for which the feedback was provided. | 
 **optional** | ***CreateMessageFeedbackOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateMessageFeedbackOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **Outcome** | **optional.String**| Whether the feedback has arrived. Can be: &#x60;unconfirmed&#x60; or &#x60;confirmed&#x60;. If &#x60;provide_feedback&#x60;&#x3D;&#x60;true&#x60; in [the initial HTTP POST](https://www.twilio.com/docs/sms/api/message-resource#create-a-message-resource), the initial value of this property is &#x60;unconfirmed&#x60;. After the message arrives, update the value to &#x60;confirmed&#x60;. | 

### Return type

[**ApiV2010AccountMessageMessageFeedback**](api.v2010.account.message.message_feedback.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateNewKey

> ApiV2010AccountNewKey CreateNewKey(ctx, AccountSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string**| The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that will be responsible for the new Key resource. | 
 **optional** | ***CreateNewKeyOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateNewKeyOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **FriendlyName** | **optional.String**| A descriptive string that you create to describe the resource. It can be up to 64 characters long. | 

### Return type

[**ApiV2010AccountNewKey**](api.v2010.account.new_key.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateNewSigningKey

> ApiV2010AccountNewSigningKey CreateNewSigningKey(ctx, AccountSid, optional)



Create a new Signing Key for the account making the request.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string**| The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that will be responsible for the new Key resource. | 
 **optional** | ***CreateNewSigningKeyOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateNewSigningKeyOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **FriendlyName** | **optional.String**| A descriptive string that you create to describe the resource. It can be up to 64 characters long. | 

### Return type

[**ApiV2010AccountNewSigningKey**](api.v2010.account.new_signing_key.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateParticipant

> ApiV2010AccountConferenceParticipant CreateParticipant(ctx, AccountSid, ConferenceSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string**| The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that will create the resource. | 
**ConferenceSid** | **string**| The SID of the participant&#39;s conference. | 
 **optional** | ***CreateParticipantOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateParticipantOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **Beep** | **optional.String**| Whether to play a notification beep to the conference when the participant joins. Can be: &#x60;true&#x60;, &#x60;false&#x60;, &#x60;onEnter&#x60;, or &#x60;onExit&#x60;. The default value is &#x60;true&#x60;. | 
 **Byoc** | **optional.String**| The SID of a BYOC (Bring Your Own Carrier) trunk to route this call with. Note that &#x60;byoc&#x60; is only meaningful when &#x60;to&#x60; is a phone number; it will otherwise be ignored. (Beta) | 
 **CallReason** | **optional.String**| The Reason for the outgoing call. Use it to specify the purpose of the call that is presented on the called party&#39;s phone. (Branded Calls Beta) | 
 **CallSidToCoach** | **optional.String**| The SID of the participant who is being &#x60;coached&#x60;. The participant being coached is the only participant who can hear the participant who is &#x60;coaching&#x60;. | 
 **CallerId** | **optional.String**| The phone number, Client identifier, or username portion of SIP address that made this call. Phone numbers are in [E.164](https://www.twilio.com/docs/glossary/what-e164) format (e.g., +16175551212). Client identifiers are formatted &#x60;client:name&#x60;. If using a phone number, it must be a Twilio number or a Verified [outgoing caller id](https://www.twilio.com/docs/voice/api/outgoing-caller-ids) for your account. If the &#x60;to&#x60; parameter is a phone number, &#x60;callerId&#x60; must also be a phone number. If &#x60;to&#x60; is sip address, this value of &#x60;callerId&#x60; should be a username portion to be used to populate the From header that is passed to the SIP endpoint. | 
 **Coaching** | **optional.Bool**| Whether the participant is coaching another call. Can be: &#x60;true&#x60; or &#x60;false&#x60;. If not present, defaults to &#x60;false&#x60; unless &#x60;call_sid_to_coach&#x60; is defined. If &#x60;true&#x60;, &#x60;call_sid_to_coach&#x60; must be defined. | 
 **ConferenceRecord** | **optional.String**| Whether to record the conference the participant is joining. Can be: &#x60;true&#x60;, &#x60;false&#x60;, &#x60;record-from-start&#x60;, and &#x60;do-not-record&#x60;. The default value is &#x60;false&#x60;. | 
 **ConferenceRecordingStatusCallback** | **optional.String**| The URL we should call using the &#x60;conference_recording_status_callback_method&#x60; when the conference recording is available. | 
 **ConferenceRecordingStatusCallbackEvent** | [**optional.Interface of []string**](string.md)| The conference recording state changes that generate a call to &#x60;conference_recording_status_callback&#x60;. Can be: &#x60;in-progress&#x60;, &#x60;completed&#x60;, and &#x60;failed&#x60;. Separate multiple values with a space. The default value is &#x60;in-progress completed failed&#x60;. | 
 **ConferenceRecordingStatusCallbackMethod** | **optional.String**| The HTTP method we should use to call &#x60;conference_recording_status_callback&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60; and defaults to &#x60;POST&#x60;. | 
 **ConferenceStatusCallback** | **optional.String**| The URL we should call using the &#x60;conference_status_callback_method&#x60; when the conference events in &#x60;conference_status_callback_event&#x60; occur. Only the value set by the first participant to join the conference is used. Subsequent &#x60;conference_status_callback&#x60; values are ignored. | 
 **ConferenceStatusCallbackEvent** | [**optional.Interface of []string**](string.md)| The conference state changes that should generate a call to &#x60;conference_status_callback&#x60;. Can be: &#x60;start&#x60;, &#x60;end&#x60;, &#x60;join&#x60;, &#x60;leave&#x60;, &#x60;mute&#x60;, &#x60;hold&#x60;, and &#x60;speaker&#x60;. Separate multiple values with a space. Defaults to &#x60;start end&#x60;. | 
 **ConferenceStatusCallbackMethod** | **optional.String**| The HTTP method we should use to call &#x60;conference_status_callback&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60; and defaults to &#x60;POST&#x60;. | 
 **ConferenceTrim** | **optional.String**| Whether to trim leading and trailing silence from your recorded conference audio files. Can be: &#x60;trim-silence&#x60; or &#x60;do-not-trim&#x60; and defaults to &#x60;trim-silence&#x60;. | 
 **EarlyMedia** | **optional.Bool**| Whether to allow an agent to hear the state of the outbound call, including ringing or disconnect messages. Can be: &#x60;true&#x60; or &#x60;false&#x60; and defaults to &#x60;true&#x60;. | 
 **EndConferenceOnExit** | **optional.Bool**| Whether to end the conference when the participant leaves. Can be: &#x60;true&#x60; or &#x60;false&#x60; and defaults to &#x60;false&#x60;. | 
 **From** | **optional.String**| The phone number, Client identifier, or username portion of SIP address that made this call. Phone numbers are in [E.164](https://www.twilio.com/docs/glossary/what-e164) format (e.g., +16175551212). Client identifiers are formatted &#x60;client:name&#x60;. If using a phone number, it must be a Twilio number or a Verified [outgoing caller id](https://www.twilio.com/docs/voice/api/outgoing-caller-ids) for your account. If the &#x60;to&#x60; parameter is a phone number, &#x60;from&#x60; must also be a phone number. If &#x60;to&#x60; is sip address, this value of &#x60;from&#x60; should be a username portion to be used to populate the P-Asserted-Identity header that is passed to the SIP endpoint. | 
 **JitterBufferSize** | **optional.String**| Jitter buffer size for the connecting participant. Twilio will use this setting to apply Jitter Buffer before participant&#39;s audio is mixed into the conference. Can be: &#x60;off&#x60;, &#x60;small&#x60;, &#x60;medium&#x60;, and &#x60;large&#x60;. Default to &#x60;large&#x60;. | 
 **Label** | **optional.String**| A label for this participant. If one is supplied, it may subsequently be used to fetch, update or delete the participant. | 
 **MaxParticipants** | **optional.Int32**| The maximum number of participants in the conference. Can be a positive integer from &#x60;2&#x60; to &#x60;250&#x60;. The default value is &#x60;250&#x60;. | 
 **Muted** | **optional.Bool**| Whether the agent is muted in the conference. Can be &#x60;true&#x60; or &#x60;false&#x60; and the default is &#x60;false&#x60;. | 
 **Record** | **optional.Bool**| Whether to record the participant and their conferences, including the time between conferences. Can be &#x60;true&#x60; or &#x60;false&#x60; and the default is &#x60;false&#x60;. | 
 **RecordingChannels** | **optional.String**| The recording channels for the final recording. Can be: &#x60;mono&#x60; or &#x60;dual&#x60; and the default is &#x60;mono&#x60;. | 
 **RecordingStatusCallback** | **optional.String**| The URL that we should call using the &#x60;recording_status_callback_method&#x60; when the recording status changes. | 
 **RecordingStatusCallbackEvent** | [**optional.Interface of []string**](string.md)| The recording state changes that should generate a call to &#x60;recording_status_callback&#x60;. Can be: &#x60;in-progress&#x60;, &#x60;completed&#x60;, and &#x60;failed&#x60;. Separate multiple values with a space. The default value is &#x60;in-progress completed failed&#x60;. | 
 **RecordingStatusCallbackMethod** | **optional.String**| The HTTP method we should use when we call &#x60;recording_status_callback&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60; and defaults to &#x60;POST&#x60;. | 
 **RecordingTrack** | **optional.String**| The audio track to record for the call. Can be: &#x60;inbound&#x60;, &#x60;outbound&#x60; or &#x60;both&#x60;. The default is &#x60;both&#x60;. &#x60;inbound&#x60; records the audio that is received by Twilio. &#x60;outbound&#x60; records the audio that is sent from Twilio. &#x60;both&#x60; records the audio that is received and sent by Twilio. | 
 **Region** | **optional.String**| The [region](https://support.twilio.com/hc/en-us/articles/223132167-How-global-low-latency-routing-and-region-selection-work-for-conferences-and-Client-calls) where we should mix the recorded audio. Can be:&#x60;us1&#x60;, &#x60;ie1&#x60;, &#x60;de1&#x60;, &#x60;sg1&#x60;, &#x60;br1&#x60;, &#x60;au1&#x60;, or &#x60;jp1&#x60;. | 
 **SipAuthPassword** | **optional.String**| The SIP password for authentication. | 
 **SipAuthUsername** | **optional.String**| The SIP username used for authentication. | 
 **StartConferenceOnEnter** | **optional.Bool**| Whether to start the conference when the participant joins, if it has not already started. Can be: &#x60;true&#x60; or &#x60;false&#x60; and the default is &#x60;true&#x60;. If &#x60;false&#x60; and the conference has not started, the participant is muted and hears background music until another participant starts the conference. | 
 **StatusCallback** | **optional.String**| The URL we should call using the &#x60;status_callback_method&#x60; to send status information to your application. | 
 **StatusCallbackEvent** | [**optional.Interface of []string**](string.md)| The conference state changes that should generate a call to &#x60;status_callback&#x60;. Can be: &#x60;initiated&#x60;, &#x60;ringing&#x60;, &#x60;answered&#x60;, and &#x60;completed&#x60;. Separate multiple values with a space. The default value is &#x60;completed&#x60;. | 
 **StatusCallbackMethod** | **optional.String**| The HTTP method we should use to call &#x60;status_callback&#x60;. Can be: &#x60;GET&#x60; and &#x60;POST&#x60; and defaults to &#x60;POST&#x60;. | 
 **Timeout** | **optional.Int32**| The number of seconds that we should allow the phone to ring before assuming there is no answer. Can be an integer between &#x60;5&#x60; and &#x60;600&#x60;, inclusive. The default value is &#x60;60&#x60;. We always add a 5-second timeout buffer to outgoing calls, so  value of 10 would result in an actual timeout that was closer to 15 seconds. | 
 **To** | **optional.String**| The phone number, SIP address, or Client identifier that received this call. Phone numbers are in [E.164](https://www.twilio.com/docs/glossary/what-e164) format (e.g., +16175551212). SIP addresses are formatted as &#x60;sip:name@company.com&#x60;. Client identifiers are formatted &#x60;client:name&#x60;. [Custom parameters](https://www.twilio.com/docs/voice/api/conference-participant-resource#custom-parameters) may also be specified. | 
 **WaitMethod** | **optional.String**| The HTTP method we should use to call &#x60;wait_url&#x60;. Can be &#x60;GET&#x60; or &#x60;POST&#x60; and the default is &#x60;POST&#x60;. When using a static audio file, this should be &#x60;GET&#x60; so that we can cache the file. | 
 **WaitUrl** | **optional.String**| The URL we should call using the &#x60;wait_method&#x60; for the music to play while participants are waiting for the conference to start. The default value is the URL of our standard hold music. [Learn more about hold music](https://www.twilio.com/labs/twimlets/holdmusic). | 

### Return type

[**ApiV2010AccountConferenceParticipant**](api.v2010.account.conference.participant.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePayments

> ApiV2010AccountCallPayments CreatePayments(ctx, AccountSid, CallSid, optional)



create an instance of payments. This will start a new payments session

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string**| The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that will create the resource. | 
**CallSid** | **string**| The SID of the call that will create the resource. Call leg associated with this sid is expected to provide payment information thru DTMF. | 
 **optional** | ***CreatePaymentsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreatePaymentsOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **BankAccountType** | **optional.String**| Type of bank account if payment source is ACH. One of &#x60;consumer-checking&#x60;, &#x60;consumer-savings&#x60;, or &#x60;commercial-checking&#x60;. The default value is &#x60;consumer-checking&#x60;. | 
 **ChargeAmount** | **optional.Float32**| A positive decimal value less than 1,000,000 to charge against the credit card or bank account. Default currency can be overwritten with &#x60;currency&#x60; field. Leave blank or set to 0 to tokenize. | 
 **Currency** | **optional.String**| The currency of the &#x60;charge_amount&#x60;, formatted as [ISO 4127](http://www.iso.org/iso/home/standards/currency_codes.htm) format. The default value is &#x60;USD&#x60; and all values allowed from the &lt;Pay&gt; Connector are accepted. | 
 **Description** | **optional.String**| The description can be used to provide more details regarding the transaction. This information is submitted along with the payment details to the Payment Connector which are then posted on the transactions. | 
 **IdempotencyKey** | **optional.String**| A unique token that will be used to ensure that multiple API calls with the same information do not result in multiple transactions. This should be a unique string value per API call and can be a randomly generated. | 
 **Input** | **optional.String**| A list of inputs that should be accepted. Currently only &#x60;dtmf&#x60; is supported. All digits captured during a pay session are redacted from the logs. | 
 **MinPostalCodeLength** | **optional.Int32**| A positive integer that is used to validate the length of the &#x60;PostalCode&#x60; inputted by the user. User must enter this many digits. | 
 **Parameter** | [**optional.Interface of map[string]interface{}**](map[string]interface{}.md)| A single level JSON string that is required when accepting certain information specific only to ACH payments. The information that has to be included here depends on the &lt;Pay&gt; Connector. [Read more](https://www.twilio.com/console/voice/pay-connectors). | 
 **PaymentConnector** | **optional.String**| This is the unique name corresponding to the Payment Gateway Connector installed in the Twilio Add-ons. Learn more about [&lt;Pay&gt; Connectors](https://www.twilio.com/console/voice/pay-connectors). The default value is &#x60;Default&#x60;. | 
 **PaymentMethod** | **optional.String**| Type of payment being captured. One of &#x60;credit-card&#x60; or &#x60;ach-debit&#x60;. The default value is &#x60;credit-card&#x60;. | 
 **PostalCode** | **optional.Bool**| Indicates whether the credit card postal code (zip code) is a required piece of payment information that must be provided by the caller. The default is &#x60;true&#x60;. | 
 **SecurityCode** | **optional.Bool**| Indicates whether the credit card security code is a required piece of payment information that must be provided by the caller. The default is &#x60;true&#x60;. | 
 **StatusCallback** | **optional.String**| Provide an absolute or relative URL to receive status updates regarding your Pay session. Read more about the [expected StatusCallback values](https://www.twilio.com/docs/voice/api/payment-resource#statuscallback) | 
 **Timeout** | **optional.Int32**| The number of seconds that &lt;Pay&gt; should wait for the caller to press a digit between each subsequent digit, after the first one, before moving on to validate the digits captured. The default is &#x60;5&#x60;, maximum is &#x60;600&#x60;. | 
 **TokenType** | **optional.String**| Indicates whether the payment method should be tokenized as a &#x60;one-time&#x60; or &#x60;reusable&#x60; token. The default value is &#x60;reusable&#x60;. Do not enter a charge amount when tokenizing. If a charge amount is entered, the payment method will be charged and not tokenized. | 
 **ValidCardTypes** | **optional.String**| Credit card types separated by space that Pay should accept. The default value is &#x60;visa mastercard amex&#x60; | 

### Return type

[**ApiV2010AccountCallPayments**](api.v2010.account.call.payments.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateQueue

> ApiV2010AccountQueue CreateQueue(ctx, AccountSid, optional)



Create a queue

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string**| The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that will create the resource. | 
 **optional** | ***CreateQueueOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateQueueOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **FriendlyName** | **optional.String**| A descriptive string that you created to describe this resource. It can be up to 64 characters long. | 
 **MaxSize** | **optional.Int32**| The maximum number of calls allowed to be in the queue. The default is 100. The maximum is 5000. | 

### Return type

[**ApiV2010AccountQueue**](api.v2010.account.queue.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSipAuthCallsCredentialListMapping

> ApiV2010AccountSipSipDomainSipAuthSipAuthCallsSipAuthCallsCredentialListMapping CreateSipAuthCallsCredentialListMapping(ctx, AccountSid, DomainSid, optional)



Create a new credential list mapping resource

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string**| The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that will create the resource. | 
**DomainSid** | **string**| The SID of the SIP domain that will contain the new resource. | 
 **optional** | ***CreateSipAuthCallsCredentialListMappingOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateSipAuthCallsCredentialListMappingOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **CredentialListSid** | **optional.String**| The SID of the CredentialList resource to map to the SIP domain. | 

### Return type

[**ApiV2010AccountSipSipDomainSipAuthSipAuthCallsSipAuthCallsCredentialListMapping**](api.v2010.account.sip.sip_domain.sip_auth.sip_auth_calls.sip_auth_calls_credential_list_mapping.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSipAuthCallsIpAccessControlListMapping

> ApiV2010AccountSipSipDomainSipAuthSipAuthCallsSipAuthCallsIpAccessControlListMapping CreateSipAuthCallsIpAccessControlListMapping(ctx, AccountSid, DomainSid, optional)



Create a new IP Access Control List mapping

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string**| The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that will create the resource. | 
**DomainSid** | **string**| The SID of the SIP domain that will contain the new resource. | 
 **optional** | ***CreateSipAuthCallsIpAccessControlListMappingOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateSipAuthCallsIpAccessControlListMappingOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **IpAccessControlListSid** | **optional.String**| The SID of the IpAccessControlList resource to map to the SIP domain. | 

### Return type

[**ApiV2010AccountSipSipDomainSipAuthSipAuthCallsSipAuthCallsIpAccessControlListMapping**](api.v2010.account.sip.sip_domain.sip_auth.sip_auth_calls.sip_auth_calls_ip_access_control_list_mapping.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSipAuthRegistrationsCredentialListMapping

> ApiV2010AccountSipSipDomainSipAuthSipAuthRegistrationsSipAuthRegistrationsCredentialListMapping CreateSipAuthRegistrationsCredentialListMapping(ctx, AccountSid, DomainSid, optional)



Create a new credential list mapping resource

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string**| The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that will create the resource. | 
**DomainSid** | **string**| The SID of the SIP domain that will contain the new resource. | 
 **optional** | ***CreateSipAuthRegistrationsCredentialListMappingOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateSipAuthRegistrationsCredentialListMappingOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **CredentialListSid** | **optional.String**| The SID of the CredentialList resource to map to the SIP domain. | 

### Return type

[**ApiV2010AccountSipSipDomainSipAuthSipAuthRegistrationsSipAuthRegistrationsCredentialListMapping**](api.v2010.account.sip.sip_domain.sip_auth.sip_auth_registrations.sip_auth_registrations_credential_list_mapping.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSipCredential

> ApiV2010AccountSipSipCredentialListSipCredential CreateSipCredential(ctx, AccountSid, CredentialListSid, optional)



Create a new credential resource.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string**| The unique id of the Account that is responsible for this resource. | 
**CredentialListSid** | **string**| The unique id that identifies the credential list to include the created credential. | 
 **optional** | ***CreateSipCredentialOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateSipCredentialOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **Password** | **optional.String**| The password that the username will use when authenticating SIP requests. The password must be a minimum of 12 characters, contain at least 1 digit, and have mixed case. (eg &#x60;IWasAtSignal2018&#x60;) | 
 **Username** | **optional.String**| The username that will be passed when authenticating SIP requests. The username should be sent in response to Twilio&#39;s challenge of the initial INVITE. It can be up to 32 characters long. | 

### Return type

[**ApiV2010AccountSipSipCredentialListSipCredential**](api.v2010.account.sip.sip_credential_list.sip_credential.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSipCredentialList

> ApiV2010AccountSipSipCredentialList CreateSipCredentialList(ctx, AccountSid, optional)



Create a Credential List

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string**| The unique id of the Account that is responsible for this resource. | 
 **optional** | ***CreateSipCredentialListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateSipCredentialListOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **FriendlyName** | **optional.String**| A human readable descriptive text that describes the CredentialList, up to 64 characters long. | 

### Return type

[**ApiV2010AccountSipSipCredentialList**](api.v2010.account.sip.sip_credential_list.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSipCredentialListMapping

> ApiV2010AccountSipSipDomainSipCredentialListMapping CreateSipCredentialListMapping(ctx, AccountSid, DomainSid, optional)



Create a CredentialListMapping resource for an account.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string**| The unique id of the [Account](https://www.twilio.com/docs/iam/api/account) responsible for this resource. | 
**DomainSid** | **string**| A 34 character string that uniquely identifies the SIP Domain for which the CredentialList resource will be mapped. | 
 **optional** | ***CreateSipCredentialListMappingOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateSipCredentialListMappingOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **CredentialListSid** | **optional.String**| A 34 character string that uniquely identifies the CredentialList resource to map to the SIP domain. | 

### Return type

[**ApiV2010AccountSipSipDomainSipCredentialListMapping**](api.v2010.account.sip.sip_domain.sip_credential_list_mapping.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSipDomain

> ApiV2010AccountSipSipDomain CreateSipDomain(ctx, AccountSid, optional)



Create a new Domain

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string**| The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that will create the resource. | 
 **optional** | ***CreateSipDomainOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateSipDomainOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ByocTrunkSid** | **optional.String**| The SID of the BYOC Trunk(Bring Your Own Carrier) resource that the Sip Domain will be associated with. | 
 **DomainName** | **optional.String**| The unique address you reserve on Twilio to which you route your SIP traffic. Domain names can contain letters, digits, and \\\&quot;-\\\&quot; and must end with &#x60;sip.twilio.com&#x60;. | 
 **EmergencyCallerSid** | **optional.String**| Whether an emergency caller sid is configured for the domain. If present, this phone number will be used as the callback for the emergency call. | 
 **EmergencyCallingEnabled** | **optional.Bool**| Whether emergency calling is enabled for the domain. If enabled, allows emergency calls on the domain from phone numbers with validated addresses. | 
 **FriendlyName** | **optional.String**| A descriptive string that you created to describe the resource. It can be up to 64 characters long. | 
 **Secure** | **optional.Bool**| Whether secure SIP is enabled for the domain. If enabled, TLS will be enforced and SRTP will be negotiated on all incoming calls to this sip domain. | 
 **SipRegistration** | **optional.Bool**| Whether to allow SIP Endpoints to register with the domain to receive calls. Can be &#x60;true&#x60; or &#x60;false&#x60;. &#x60;true&#x60; allows SIP Endpoints to register with the domain to receive calls, &#x60;false&#x60; does not. | 
 **VoiceFallbackMethod** | **optional.String**| The HTTP method we should use to call &#x60;voice_fallback_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60;. | 
 **VoiceFallbackUrl** | **optional.String**| The URL that we should call when an error occurs while retrieving or executing the TwiML from &#x60;voice_url&#x60;. | 
 **VoiceMethod** | **optional.String**| The HTTP method we should use to call &#x60;voice_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60;. | 
 **VoiceStatusCallbackMethod** | **optional.String**| The HTTP method we should use to call &#x60;voice_status_callback_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60;. | 
 **VoiceStatusCallbackUrl** | **optional.String**| The URL that we should call to pass status parameters (such as call ended) to your application. | 
 **VoiceUrl** | **optional.String**| The URL we should when the domain receives a call. | 

### Return type

[**ApiV2010AccountSipSipDomain**](api.v2010.account.sip.sip_domain.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSipIpAccessControlList

> ApiV2010AccountSipSipIpAccessControlList CreateSipIpAccessControlList(ctx, AccountSid, optional)



Create a new IpAccessControlList resource

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string**| The unique id of the [Account](https://www.twilio.com/docs/iam/api/account) responsible for this resource. | 
 **optional** | ***CreateSipIpAccessControlListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateSipIpAccessControlListOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **FriendlyName** | **optional.String**| A human readable descriptive text that describes the IpAccessControlList, up to 64 characters long. | 

### Return type

[**ApiV2010AccountSipSipIpAccessControlList**](api.v2010.account.sip.sip_ip_access_control_list.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSipIpAccessControlListMapping

> ApiV2010AccountSipSipDomainSipIpAccessControlListMapping CreateSipIpAccessControlListMapping(ctx, AccountSid, DomainSid, optional)



Create a new IpAccessControlListMapping resource.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string**| The unique id of the Account that is responsible for this resource. | 
**DomainSid** | **string**| A 34 character string that uniquely identifies the SIP domain. | 
 **optional** | ***CreateSipIpAccessControlListMappingOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateSipIpAccessControlListMappingOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **IpAccessControlListSid** | **optional.String**| The unique id of the IP access control list to map to the SIP domain. | 

### Return type

[**ApiV2010AccountSipSipDomainSipIpAccessControlListMapping**](api.v2010.account.sip.sip_domain.sip_ip_access_control_list_mapping.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSipIpAddress

> ApiV2010AccountSipSipIpAccessControlListSipIpAddress CreateSipIpAddress(ctx, AccountSid, IpAccessControlListSid, optional)



Create a new IpAddress resource.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string**| The unique id of the [Account](https://www.twilio.com/docs/iam/api/account) responsible for this resource. | 
**IpAccessControlListSid** | **string**| The IpAccessControlList Sid with which to associate the created IpAddress resource. | 
 **optional** | ***CreateSipIpAddressOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateSipIpAddressOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **CidrPrefixLength** | **optional.Int32**| An integer representing the length of the CIDR prefix to use with this IP address when accepting traffic. By default the entire IP address is used. | 
 **FriendlyName** | **optional.String**| A human readable descriptive text for this resource, up to 64 characters long. | 
 **IpAddress** | **optional.String**| An IP address in dotted decimal notation from which you want to accept traffic. Any SIP requests from this IP address will be allowed by Twilio. IPv4 only supported today. | 

### Return type

[**ApiV2010AccountSipSipIpAccessControlListSipIpAddress**](api.v2010.account.sip.sip_ip_access_control_list.sip_ip_address.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateToken

> ApiV2010AccountToken CreateToken(ctx, AccountSid, optional)



Create a new token for ICE servers

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string**| The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that will create the resource. | 
 **optional** | ***CreateTokenOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateTokenOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **Ttl** | **optional.Int32**| The duration in seconds for which the generated credentials are valid. The default value is 86400 (24 hours). | 

### Return type

[**ApiV2010AccountToken**](api.v2010.account.token.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateUsageTrigger

> ApiV2010AccountUsageUsageTrigger CreateUsageTrigger(ctx, AccountSid, optional)



Create a new UsageTrigger

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string**| The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that will create the resource. | 
 **optional** | ***CreateUsageTriggerOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateUsageTriggerOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **CallbackMethod** | **optional.String**| The HTTP method we should use to call &#x60;callback_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60; and the default is &#x60;POST&#x60;. | 
 **CallbackUrl** | **optional.String**| The URL we should call using &#x60;callback_method&#x60; when the trigger fires. | 
 **FriendlyName** | **optional.String**| A descriptive string that you create to describe the resource. It can be up to 64 characters long. | 
 **Recurring** | **optional.String**| The frequency of a recurring UsageTrigger.  Can be: &#x60;daily&#x60;, &#x60;monthly&#x60;, or &#x60;yearly&#x60; for recurring triggers or empty for non-recurring triggers. A trigger will only fire once during each period. Recurring times are in GMT. | 
 **TriggerBy** | **optional.String**| The field in the [UsageRecord](https://www.twilio.com/docs/usage/api/usage-record) resource that should fire the trigger.  Can be: &#x60;count&#x60;, &#x60;usage&#x60;, or &#x60;price&#x60; as described in the [UsageRecords documentation](https://www.twilio.com/docs/usage/api/usage-record#usage-count-price).  The default is &#x60;usage&#x60;. | 
 **TriggerValue** | **optional.String**| The usage value at which the trigger should fire.  For convenience, you can use an offset value such as &#x60;+30&#x60; to specify a trigger_value that is 30 units more than the current usage value. Be sure to urlencode a &#x60;+&#x60; as &#x60;%2B&#x60;. | 
 **UsageCategory** | **optional.String**| The usage category that the trigger should watch.  Use one of the supported [usage categories](https://www.twilio.com/docs/usage/api/usage-record#usage-categories) for this value. | 

### Return type

[**ApiV2010AccountUsageUsageTrigger**](api.v2010.account.usage.usage_trigger.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateValidationRequest

> ApiV2010AccountValidationRequest CreateValidationRequest(ctx, AccountSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string**| The SID of the [Account](https://www.twilio.com/docs/iam/api/account) responsible for the new caller ID resource. | 
 **optional** | ***CreateValidationRequestOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateValidationRequestOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **CallDelay** | **optional.Int32**| The number of seconds to delay before initiating the verification call. Can be an integer between &#x60;0&#x60; and &#x60;60&#x60;, inclusive. The default is &#x60;0&#x60;. | 
 **Extension** | **optional.String**| The digits to dial after connecting the verification call. | 
 **FriendlyName** | **optional.String**| A descriptive string that you create to describe the new caller ID resource. It can be up to 64 characters long. The default value is a formatted version of the phone number. | 
 **PhoneNumber** | **optional.String**| The phone number to verify in [E.164](https://www.twilio.com/docs/glossary/what-e164) format, which consists of a + followed by the country code and subscriber number. | 
 **StatusCallback** | **optional.String**| The URL we should call using the &#x60;status_callback_method&#x60; to send status information about the verification process to your application. | 
 **StatusCallbackMethod** | **optional.String**| The HTTP method we should use to call &#x60;status_callback&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60;, and the default is &#x60;POST&#x60;. | 

### Return type

[**ApiV2010AccountValidationRequest**](api.v2010.account.validation_request.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAddress

> DeleteAddress(ctx, AccountSid, Sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string**| The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that is responsible for the Address resource to delete. | 
**Sid** | **string**| The Twilio-provided string that uniquely identifies the Address resource to delete. | 

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

> DeleteApplication(ctx, AccountSid, Sid)



Delete the application by the specified application sid

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string**| The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Application resources to delete. | 
**Sid** | **string**| The Twilio-provided string that uniquely identifies the Application resource to delete. | 

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

> DeleteCall(ctx, AccountSid, Sid)



Delete a Call record from your account. Once the record is deleted, it will no longer appear in the API and Account Portal logs.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string**| The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Call resource(s) to delete. | 
**Sid** | **string**| The Twilio-provided Call SID that uniquely identifies the Call resource to delete | 

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

> DeleteCallFeedbackSummary(ctx, AccountSid, Sid)



Delete a FeedbackSummary resource from a call

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string**| The unique id of the [Account](https://www.twilio.com/docs/iam/api/account) responsible for this resource. | 
**Sid** | **string**| A 34 character string that uniquely identifies this resource. | 

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

> DeleteCallRecording(ctx, AccountSid, CallSid, Sid)



Delete a recording from your account

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string**| The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Recording resources to delete. | 
**CallSid** | **string**| The [Call](https://www.twilio.com/docs/voice/api/call-resource) SID of the resources to delete. | 
**Sid** | **string**| The Twilio-provided string that uniquely identifies the Recording resource to delete. | 

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

> DeleteConferenceRecording(ctx, AccountSid, ConferenceSid, Sid)



Delete a recording from your account

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string**| The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Conference Recording resources to delete. | 
**ConferenceSid** | **string**| The Conference SID that identifies the conference associated with the recording to delete. | 
**Sid** | **string**| The Twilio-provided string that uniquely identifies the Conference Recording resource to delete. | 

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

> DeleteConnectApp(ctx, AccountSid, Sid)



Delete an instance of a connect-app

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string**| The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the ConnectApp resource to fetch. | 
**Sid** | **string**| The Twilio-provided string that uniquely identifies the ConnectApp resource to fetch. | 

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

> DeleteIncomingPhoneNumber(ctx, AccountSid, Sid)



Delete a phone-numbers belonging to the account used to make the request.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string**| The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the IncomingPhoneNumber resources to delete. | 
**Sid** | **string**| The Twilio-provided string that uniquely identifies the IncomingPhoneNumber resource to delete. | 

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

> DeleteIncomingPhoneNumberAssignedAddOn(ctx, AccountSid, ResourceSid, Sid)



Remove the assignment of an Add-on installation from the Number specified.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string**| The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the resources to delete. | 
**ResourceSid** | **string**| The SID of the Phone Number to which the Add-on is assigned. | 
**Sid** | **string**| The Twilio-provided string that uniquely identifies the resource to delete. | 

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

> DeleteKey(ctx, AccountSid, Sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string**| The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Key resources to delete. | 
**Sid** | **string**| The Twilio-provided string that uniquely identifies the Key resource to delete. | 

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

> DeleteMedia(ctx, AccountSid, MessageSid, Sid)



Delete media from your account. Once delete, you will no longer be billed

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string**| The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Media resource(s) to delete. | 
**MessageSid** | **string**| The SID of the Message resource that this Media resource belongs to. | 
**Sid** | **string**| The Twilio-provided string that uniquely identifies the Media resource to delete | 

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

> DeleteMessage(ctx, AccountSid, Sid)



Deletes a message record from your account

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string**| The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Message resources to delete. | 
**Sid** | **string**| The Twilio-provided string that uniquely identifies the Message resource to delete. | 

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

> DeleteOutgoingCallerId(ctx, AccountSid, Sid)



Delete the caller-id specified from the account

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string**| The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the OutgoingCallerId resources to delete. | 
**Sid** | **string**| The Twilio-provided string that uniquely identifies the OutgoingCallerId resource to delete. | 

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

> DeleteParticipant(ctx, AccountSid, ConferenceSid, CallSid)



Kick a participant from a given conference

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string**| The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Participant resources to delete. | 
**ConferenceSid** | **string**| The SID of the conference with the participants to delete. | 
**CallSid** | **string**| The [Call](https://www.twilio.com/docs/voice/api/call-resource) SID or label of the participant to delete. Non URL safe characters in a label must be percent encoded, for example, a space character is represented as %20. | 

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

> DeleteQueue(ctx, AccountSid, Sid)



Remove an empty queue

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string**| The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Queue resource to delete. | 
**Sid** | **string**| The Twilio-provided string that uniquely identifies the Queue resource to delete | 

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

> DeleteRecording(ctx, AccountSid, Sid)



Delete a recording from your account

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string**| The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Recording resources to delete. | 
**Sid** | **string**| The Twilio-provided string that uniquely identifies the Recording resource to delete. | 

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

> DeleteRecordingAddOnResult(ctx, AccountSid, ReferenceSid, Sid)



Delete a result and purge all associated Payloads

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string**| The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Recording AddOnResult resources to delete. | 
**ReferenceSid** | **string**| The SID of the recording to which the result to delete belongs. | 
**Sid** | **string**| The Twilio-provided string that uniquely identifies the Recording AddOnResult resource to delete. | 

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

> DeleteRecordingAddOnResultPayload(ctx, AccountSid, ReferenceSid, AddOnResultSid, Sid)



Delete a payload from the result along with all associated Data

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string**| The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Recording AddOnResult Payload resources to delete. | 
**ReferenceSid** | **string**| The SID of the recording to which the AddOnResult resource that contains the payloads to delete belongs. | 
**AddOnResultSid** | **string**| The SID of the AddOnResult to which the payloads to delete belongs. | 
**Sid** | **string**| The Twilio-provided string that uniquely identifies the Recording AddOnResult Payload resource to delete. | 

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

> DeleteRecordingTranscription(ctx, AccountSid, RecordingSid, Sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string**| The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Transcription resources to delete. | 
**RecordingSid** | **string**| The SID of the [Recording](https://www.twilio.com/docs/voice/api/recording) that created the transcription to delete. | 
**Sid** | **string**| The Twilio-provided string that uniquely identifies the Transcription resource to delete. | 

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

> DeleteSigningKey(ctx, AccountSid, Sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string**|  | 
**Sid** | **string**|  | 

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

> DeleteSipAuthCallsCredentialListMapping(ctx, AccountSid, DomainSid, Sid)



Delete a credential list mapping from the requested domain

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string**| The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the CredentialListMapping resources to delete. | 
**DomainSid** | **string**| The SID of the SIP domain that contains the resource to delete. | 
**Sid** | **string**| The Twilio-provided string that uniquely identifies the CredentialListMapping resource to delete. | 

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

> DeleteSipAuthCallsIpAccessControlListMapping(ctx, AccountSid, DomainSid, Sid)



Delete an IP Access Control List mapping from the requested domain

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string**| The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the IpAccessControlListMapping resources to delete. | 
**DomainSid** | **string**| The SID of the SIP domain that contains the resources to delete. | 
**Sid** | **string**| The Twilio-provided string that uniquely identifies the IpAccessControlListMapping resource to delete. | 

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

> DeleteSipAuthRegistrationsCredentialListMapping(ctx, AccountSid, DomainSid, Sid)



Delete a credential list mapping from the requested domain

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string**| The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the CredentialListMapping resources to delete. | 
**DomainSid** | **string**| The SID of the SIP domain that contains the resources to delete. | 
**Sid** | **string**| The Twilio-provided string that uniquely identifies the CredentialListMapping resource to delete. | 

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

> DeleteSipCredential(ctx, AccountSid, CredentialListSid, Sid)



Delete a credential resource.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string**| The unique id of the Account that is responsible for this resource. | 
**CredentialListSid** | **string**| The unique id that identifies the credential list that contains the desired credentials. | 
**Sid** | **string**| The unique id that identifies the resource to delete. | 

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

> DeleteSipCredentialList(ctx, AccountSid, Sid)



Delete a Credential List

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string**| The unique id of the Account that is responsible for this resource. | 
**Sid** | **string**| The credential list Sid that uniquely identifies this resource | 

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

> DeleteSipCredentialListMapping(ctx, AccountSid, DomainSid, Sid)



Delete a CredentialListMapping resource from an account.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string**| The unique id of the [Account](https://www.twilio.com/docs/iam/api/account) responsible for this resource. | 
**DomainSid** | **string**| A 34 character string that uniquely identifies the SIP Domain that includes the resource to delete. | 
**Sid** | **string**| A 34 character string that uniquely identifies the resource to delete. | 

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

> DeleteSipDomain(ctx, AccountSid, Sid)



Delete an instance of a Domain

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string**| The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the SipDomain resources to delete. | 
**Sid** | **string**| The Twilio-provided string that uniquely identifies the SipDomain resource to delete. | 

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

> DeleteSipIpAccessControlList(ctx, AccountSid, Sid)



Delete an IpAccessControlList from the requested account

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string**| The unique id of the [Account](https://www.twilio.com/docs/iam/api/account) responsible for this resource. | 
**Sid** | **string**| A 34 character string that uniquely identifies the resource to delete. | 

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

> DeleteSipIpAccessControlListMapping(ctx, AccountSid, DomainSid, Sid)



Delete an IpAccessControlListMapping resource.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string**| The unique id of the Account that is responsible for this resource. | 
**DomainSid** | **string**| A 34 character string that uniquely identifies the SIP domain. | 
**Sid** | **string**| A 34 character string that uniquely identifies the resource to delete. | 

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

> DeleteSipIpAddress(ctx, AccountSid, IpAccessControlListSid, Sid)



Delete an IpAddress resource.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string**| The unique id of the [Account](https://www.twilio.com/docs/iam/api/account) responsible for this resource. | 
**IpAccessControlListSid** | **string**| The IpAccessControlList Sid that identifies the IpAddress resources to delete. | 
**Sid** | **string**| A 34 character string that uniquely identifies the resource to delete. | 

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

> DeleteTranscription(ctx, AccountSid, Sid)



Delete a transcription from the account used to make the request

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string**| The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Transcription resources to delete. | 
**Sid** | **string**| The Twilio-provided string that uniquely identifies the Transcription resource to delete. | 

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

> DeleteUsageTrigger(ctx, AccountSid, Sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string**| The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the UsageTrigger resources to delete. | 
**Sid** | **string**| The Twilio-provided string that uniquely identifies the UsageTrigger resource to delete. | 

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

> ApiV2010Account FetchAccount(ctx, Sid)



Fetch the account specified by the provided Account Sid

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string**| The Account Sid that uniquely identifies the account to fetch | 

### Return type

[**ApiV2010Account**](api.v2010.account.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchAddress

> ApiV2010AccountAddress FetchAddress(ctx, AccountSid, Sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string**| The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that is responsible for the Address resource to fetch. | 
**Sid** | **string**| The Twilio-provided string that uniquely identifies the Address resource to fetch. | 

### Return type

[**ApiV2010AccountAddress**](api.v2010.account.address.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchApplication

> ApiV2010AccountApplication FetchApplication(ctx, AccountSid, Sid)



Fetch the application specified by the provided sid

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string**| The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Application resource to fetch. | 
**Sid** | **string**| The Twilio-provided string that uniquely identifies the Application resource to fetch. | 

### Return type

[**ApiV2010AccountApplication**](api.v2010.account.application.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchAuthorizedConnectApp

> ApiV2010AccountAuthorizedConnectApp FetchAuthorizedConnectApp(ctx, AccountSid, ConnectAppSid)



Fetch an instance of an authorized-connect-app

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string**| The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the AuthorizedConnectApp resource to fetch. | 
**ConnectAppSid** | **string**| The SID of the Connect App to fetch. | 

### Return type

[**ApiV2010AccountAuthorizedConnectApp**](api.v2010.account.authorized_connect_app.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchAvailablePhoneNumberCountry

> ApiV2010AccountAvailablePhoneNumberCountry FetchAvailablePhoneNumberCountry(ctx, AccountSid, CountryCode)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string**| The SID of the [Account](https://www.twilio.com/docs/iam/api/account) requesting the available phone number Country resource. | 
**CountryCode** | **string**| The [ISO-3166-1](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country code of the country to fetch available phone number information about. | 

### Return type

[**ApiV2010AccountAvailablePhoneNumberCountry**](api.v2010.account.available_phone_number_country.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchBalance

> ApiV2010AccountBalance FetchBalance(ctx, AccountSid)



Fetch the balance for an Account based on Account Sid. Balance changes may not be reflected immediately. Child accounts do not contain balance information

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string**| The unique SID identifier of the Account. | 

### Return type

[**ApiV2010AccountBalance**](api.v2010.account.balance.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchCall

> ApiV2010AccountCall FetchCall(ctx, AccountSid, Sid)



Fetch the call specified by the provided Call SID

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string**| The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Call resource(s) to fetch. | 
**Sid** | **string**| The SID of the Call resource to fetch. | 

### Return type

[**ApiV2010AccountCall**](api.v2010.account.call.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchCallFeedback

> ApiV2010AccountCallCallFeedback FetchCallFeedback(ctx, AccountSid, CallSid)



Fetch a Feedback resource from a call

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string**| The unique id of the [Account](https://www.twilio.com/docs/iam/api/account) responsible for this resource. | 
**CallSid** | **string**| The call sid that uniquely identifies the call | 

### Return type

[**ApiV2010AccountCallCallFeedback**](api.v2010.account.call.call_feedback.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchCallFeedbackSummary

> ApiV2010AccountCallCallFeedbackSummary FetchCallFeedbackSummary(ctx, AccountSid, Sid)



Fetch a FeedbackSummary resource from a call

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string**| The unique id of the [Account](https://www.twilio.com/docs/iam/api/account) responsible for this resource. | 
**Sid** | **string**| A 34 character string that uniquely identifies this resource. | 

### Return type

[**ApiV2010AccountCallCallFeedbackSummary**](api.v2010.account.call.call_feedback_summary.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchCallNotification

> ApiV2010AccountCallCallNotificationInstance FetchCallNotification(ctx, AccountSid, CallSid, Sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string**| The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Call Notification resource to fetch. | 
**CallSid** | **string**| The [Call](https://www.twilio.com/docs/voice/api/call-resource) SID of the Call Notification resource to fetch. | 
**Sid** | **string**| The Twilio-provided string that uniquely identifies the Call Notification resource to fetch. | 

### Return type

[**ApiV2010AccountCallCallNotificationInstance**](api.v2010.account.call.call_notification-instance.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchCallRecording

> ApiV2010AccountCallCallRecording FetchCallRecording(ctx, AccountSid, CallSid, Sid)



Fetch an instance of a recording for a call

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string**| The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Recording resource to fetch. | 
**CallSid** | **string**| The [Call](https://www.twilio.com/docs/voice/api/call-resource) SID of the resource to fetch. | 
**Sid** | **string**| The Twilio-provided string that uniquely identifies the Recording resource to fetch. | 

### Return type

[**ApiV2010AccountCallCallRecording**](api.v2010.account.call.call_recording.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchConference

> ApiV2010AccountConference FetchConference(ctx, AccountSid, Sid)



Fetch an instance of a conference

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string**| The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Conference resource(s) to fetch. | 
**Sid** | **string**| The Twilio-provided string that uniquely identifies the Conference resource to fetch | 

### Return type

[**ApiV2010AccountConference**](api.v2010.account.conference.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchConferenceRecording

> ApiV2010AccountConferenceConferenceRecording FetchConferenceRecording(ctx, AccountSid, ConferenceSid, Sid)



Fetch an instance of a recording for a call

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string**| The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Conference Recording resource to fetch. | 
**ConferenceSid** | **string**| The Conference SID that identifies the conference associated with the recording to fetch. | 
**Sid** | **string**| The Twilio-provided string that uniquely identifies the Conference Recording resource to fetch. | 

### Return type

[**ApiV2010AccountConferenceConferenceRecording**](api.v2010.account.conference.conference_recording.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchConnectApp

> ApiV2010AccountConnectApp FetchConnectApp(ctx, AccountSid, Sid)



Fetch an instance of a connect-app

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string**| The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the ConnectApp resource to fetch. | 
**Sid** | **string**| The Twilio-provided string that uniquely identifies the ConnectApp resource to fetch. | 

### Return type

[**ApiV2010AccountConnectApp**](api.v2010.account.connect_app.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchIncomingPhoneNumber

> ApiV2010AccountIncomingPhoneNumber FetchIncomingPhoneNumber(ctx, AccountSid, Sid)



Fetch an incoming-phone-number belonging to the account used to make the request.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string**| The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the IncomingPhoneNumber resource to fetch. | 
**Sid** | **string**| The Twilio-provided string that uniquely identifies the IncomingPhoneNumber resource to fetch. | 

### Return type

[**ApiV2010AccountIncomingPhoneNumber**](api.v2010.account.incoming_phone_number.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchIncomingPhoneNumberAssignedAddOn

> ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberAssignedAddOn FetchIncomingPhoneNumberAssignedAddOn(ctx, AccountSid, ResourceSid, Sid)



Fetch an instance of an Add-on installation currently assigned to this Number.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string**| The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the resource to fetch. | 
**ResourceSid** | **string**| The SID of the Phone Number to which the Add-on is assigned. | 
**Sid** | **string**| The Twilio-provided string that uniquely identifies the resource to fetch. | 

### Return type

[**ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberAssignedAddOn**](api.v2010.account.incoming_phone_number.incoming_phone_number_assigned_add_on.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchIncomingPhoneNumberAssignedAddOnExtension

> ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberAssignedAddOnIncomingPhoneNumberAssignedAddOnExtension FetchIncomingPhoneNumberAssignedAddOnExtension(ctx, AccountSid, ResourceSid, AssignedAddOnSid, Sid)



Fetch an instance of an Extension for the Assigned Add-on.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string**| The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the resource to fetch. | 
**ResourceSid** | **string**| The SID of the Phone Number to which the Add-on is assigned. | 
**AssignedAddOnSid** | **string**| The SID that uniquely identifies the assigned Add-on installation. | 
**Sid** | **string**| The Twilio-provided string that uniquely identifies the resource to fetch. | 

### Return type

[**ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberAssignedAddOnIncomingPhoneNumberAssignedAddOnExtension**](api.v2010.account.incoming_phone_number.incoming_phone_number_assigned_add_on.incoming_phone_number_assigned_add_on_extension.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchKey

> ApiV2010AccountKey FetchKey(ctx, AccountSid, Sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string**| The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Key resource to fetch. | 
**Sid** | **string**| The Twilio-provided string that uniquely identifies the Key resource to fetch. | 

### Return type

[**ApiV2010AccountKey**](api.v2010.account.key.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchMedia

> ApiV2010AccountMessageMedia FetchMedia(ctx, AccountSid, MessageSid, Sid)



Fetch a single media instance belonging to the account used to make the request

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string**| The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Media resource(s) to fetch. | 
**MessageSid** | **string**| The SID of the Message resource that this Media resource belongs to. | 
**Sid** | **string**| The Twilio-provided string that uniquely identifies the Media resource to fetch | 

### Return type

[**ApiV2010AccountMessageMedia**](api.v2010.account.message.media.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchMember

> ApiV2010AccountQueueMember FetchMember(ctx, AccountSid, QueueSid, CallSid)



Fetch a specific member from the queue

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string**| The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Member resource(s) to fetch. | 
**QueueSid** | **string**| The SID of the Queue in which to find the members to fetch. | 
**CallSid** | **string**| The [Call](https://www.twilio.com/docs/voice/api/call-resource) SID of the resource(s) to fetch. | 

### Return type

[**ApiV2010AccountQueueMember**](api.v2010.account.queue.member.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchMessage

> ApiV2010AccountMessage FetchMessage(ctx, AccountSid, Sid)



Fetch a message belonging to the account used to make the request

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string**| The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Message resource to fetch. | 
**Sid** | **string**| The Twilio-provided string that uniquely identifies the Message resource to fetch. | 

### Return type

[**ApiV2010AccountMessage**](api.v2010.account.message.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchNotification

> ApiV2010AccountNotificationInstance FetchNotification(ctx, AccountSid, Sid)



Fetch a notification belonging to the account used to make the request

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string**| The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Notification resource to fetch. | 
**Sid** | **string**| The Twilio-provided string that uniquely identifies the Notification resource to fetch. | 

### Return type

[**ApiV2010AccountNotificationInstance**](api.v2010.account.notification-instance.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchOutgoingCallerId

> ApiV2010AccountOutgoingCallerId FetchOutgoingCallerId(ctx, AccountSid, Sid)



Fetch an outgoing-caller-id belonging to the account used to make the request

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string**| The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the OutgoingCallerId resource to fetch. | 
**Sid** | **string**| The Twilio-provided string that uniquely identifies the OutgoingCallerId resource to fetch. | 

### Return type

[**ApiV2010AccountOutgoingCallerId**](api.v2010.account.outgoing_caller_id.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchParticipant

> ApiV2010AccountConferenceParticipant FetchParticipant(ctx, AccountSid, ConferenceSid, CallSid)



Fetch an instance of a participant

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string**| The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Participant resource to fetch. | 
**ConferenceSid** | **string**| The SID of the conference with the participant to fetch. | 
**CallSid** | **string**| The [Call](https://www.twilio.com/docs/voice/api/call-resource) SID or label of the participant to fetch. Non URL safe characters in a label must be percent encoded, for example, a space character is represented as %20. | 

### Return type

[**ApiV2010AccountConferenceParticipant**](api.v2010.account.conference.participant.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchQueue

> ApiV2010AccountQueue FetchQueue(ctx, AccountSid, Sid)



Fetch an instance of a queue identified by the QueueSid

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string**| The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Queue resource to fetch. | 
**Sid** | **string**| The Twilio-provided string that uniquely identifies the Queue resource to fetch | 

### Return type

[**ApiV2010AccountQueue**](api.v2010.account.queue.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchRecording

> ApiV2010AccountRecording FetchRecording(ctx, AccountSid, Sid)



Fetch an instance of a recording

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string**| The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Recording resource to fetch. | 
**Sid** | **string**| The Twilio-provided string that uniquely identifies the Recording resource to fetch. | 

### Return type

[**ApiV2010AccountRecording**](api.v2010.account.recording.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchRecordingAddOnResult

> ApiV2010AccountRecordingRecordingAddOnResult FetchRecordingAddOnResult(ctx, AccountSid, ReferenceSid, Sid)



Fetch an instance of an AddOnResult

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string**| The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Recording AddOnResult resource to fetch. | 
**ReferenceSid** | **string**| The SID of the recording to which the result to fetch belongs. | 
**Sid** | **string**| The Twilio-provided string that uniquely identifies the Recording AddOnResult resource to fetch. | 

### Return type

[**ApiV2010AccountRecordingRecordingAddOnResult**](api.v2010.account.recording.recording_add_on_result.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchRecordingAddOnResultPayload

> ApiV2010AccountRecordingRecordingAddOnResultRecordingAddOnResultPayload FetchRecordingAddOnResultPayload(ctx, AccountSid, ReferenceSid, AddOnResultSid, Sid)



Fetch an instance of a result payload

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string**| The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Recording AddOnResult Payload resource to fetch. | 
**ReferenceSid** | **string**| The SID of the recording to which the AddOnResult resource that contains the payload to fetch belongs. | 
**AddOnResultSid** | **string**| The SID of the AddOnResult to which the payload to fetch belongs. | 
**Sid** | **string**| The Twilio-provided string that uniquely identifies the Recording AddOnResult Payload resource to fetch. | 

### Return type

[**ApiV2010AccountRecordingRecordingAddOnResultRecordingAddOnResultPayload**](api.v2010.account.recording.recording_add_on_result.recording_add_on_result_payload.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchRecordingTranscription

> ApiV2010AccountRecordingRecordingTranscription FetchRecordingTranscription(ctx, AccountSid, RecordingSid, Sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string**| The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Transcription resource to fetch. | 
**RecordingSid** | **string**| The SID of the [Recording](https://www.twilio.com/docs/voice/api/recording) that created the transcription to fetch. | 
**Sid** | **string**| The Twilio-provided string that uniquely identifies the Transcription resource to fetch. | 

### Return type

[**ApiV2010AccountRecordingRecordingTranscription**](api.v2010.account.recording.recording_transcription.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchShortCode

> ApiV2010AccountShortCode FetchShortCode(ctx, AccountSid, Sid)



Fetch an instance of a short code

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string**| The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the ShortCode resource(s) to fetch. | 
**Sid** | **string**| The Twilio-provided string that uniquely identifies the ShortCode resource to fetch | 

### Return type

[**ApiV2010AccountShortCode**](api.v2010.account.short_code.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchSigningKey

> ApiV2010AccountSigningKey FetchSigningKey(ctx, AccountSid, Sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string**|  | 
**Sid** | **string**|  | 

### Return type

[**ApiV2010AccountSigningKey**](api.v2010.account.signing_key.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchSipAuthCallsCredentialListMapping

> ApiV2010AccountSipSipDomainSipAuthSipAuthCallsSipAuthCallsCredentialListMapping FetchSipAuthCallsCredentialListMapping(ctx, AccountSid, DomainSid, Sid)



Fetch a specific instance of a credential list mapping

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string**| The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the CredentialListMapping resource to fetch. | 
**DomainSid** | **string**| The SID of the SIP domain that contains the resource to fetch. | 
**Sid** | **string**| The Twilio-provided string that uniquely identifies the CredentialListMapping resource to fetch. | 

### Return type

[**ApiV2010AccountSipSipDomainSipAuthSipAuthCallsSipAuthCallsCredentialListMapping**](api.v2010.account.sip.sip_domain.sip_auth.sip_auth_calls.sip_auth_calls_credential_list_mapping.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchSipAuthCallsIpAccessControlListMapping

> ApiV2010AccountSipSipDomainSipAuthSipAuthCallsSipAuthCallsIpAccessControlListMapping FetchSipAuthCallsIpAccessControlListMapping(ctx, AccountSid, DomainSid, Sid)



Fetch a specific instance of an IP Access Control List mapping

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string**| The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the IpAccessControlListMapping resource to fetch. | 
**DomainSid** | **string**| The SID of the SIP domain that contains the resource to fetch. | 
**Sid** | **string**| The Twilio-provided string that uniquely identifies the IpAccessControlListMapping resource to fetch. | 

### Return type

[**ApiV2010AccountSipSipDomainSipAuthSipAuthCallsSipAuthCallsIpAccessControlListMapping**](api.v2010.account.sip.sip_domain.sip_auth.sip_auth_calls.sip_auth_calls_ip_access_control_list_mapping.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchSipAuthRegistrationsCredentialListMapping

> ApiV2010AccountSipSipDomainSipAuthSipAuthRegistrationsSipAuthRegistrationsCredentialListMapping FetchSipAuthRegistrationsCredentialListMapping(ctx, AccountSid, DomainSid, Sid)



Fetch a specific instance of a credential list mapping

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string**| The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the CredentialListMapping resource to fetch. | 
**DomainSid** | **string**| The SID of the SIP domain that contains the resource to fetch. | 
**Sid** | **string**| The Twilio-provided string that uniquely identifies the CredentialListMapping resource to fetch. | 

### Return type

[**ApiV2010AccountSipSipDomainSipAuthSipAuthRegistrationsSipAuthRegistrationsCredentialListMapping**](api.v2010.account.sip.sip_domain.sip_auth.sip_auth_registrations.sip_auth_registrations_credential_list_mapping.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchSipCredential

> ApiV2010AccountSipSipCredentialListSipCredential FetchSipCredential(ctx, AccountSid, CredentialListSid, Sid)



Fetch a single credential.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string**| The unique id of the Account that is responsible for this resource. | 
**CredentialListSid** | **string**| The unique id that identifies the credential list that contains the desired credential. | 
**Sid** | **string**| The unique id that identifies the resource to fetch. | 

### Return type

[**ApiV2010AccountSipSipCredentialListSipCredential**](api.v2010.account.sip.sip_credential_list.sip_credential.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchSipCredentialList

> ApiV2010AccountSipSipCredentialList FetchSipCredentialList(ctx, AccountSid, Sid)



Get a Credential List

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string**| The unique id of the Account that is responsible for this resource. | 
**Sid** | **string**| The credential list Sid that uniquely identifies this resource | 

### Return type

[**ApiV2010AccountSipSipCredentialList**](api.v2010.account.sip.sip_credential_list.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchSipCredentialListMapping

> ApiV2010AccountSipSipDomainSipCredentialListMapping FetchSipCredentialListMapping(ctx, AccountSid, DomainSid, Sid)



Fetch a single CredentialListMapping resource from an account.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string**| The unique id of the [Account](https://www.twilio.com/docs/iam/api/account) responsible for this resource. | 
**DomainSid** | **string**| A 34 character string that uniquely identifies the SIP Domain that includes the resource to fetch. | 
**Sid** | **string**| A 34 character string that uniquely identifies the resource to fetch. | 

### Return type

[**ApiV2010AccountSipSipDomainSipCredentialListMapping**](api.v2010.account.sip.sip_domain.sip_credential_list_mapping.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchSipDomain

> ApiV2010AccountSipSipDomain FetchSipDomain(ctx, AccountSid, Sid)



Fetch an instance of a Domain

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string**| The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the SipDomain resource to fetch. | 
**Sid** | **string**| The Twilio-provided string that uniquely identifies the SipDomain resource to fetch. | 

### Return type

[**ApiV2010AccountSipSipDomain**](api.v2010.account.sip.sip_domain.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchSipIpAccessControlList

> ApiV2010AccountSipSipIpAccessControlList FetchSipIpAccessControlList(ctx, AccountSid, Sid)



Fetch a specific instance of an IpAccessControlList

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string**| The unique id of the [Account](https://www.twilio.com/docs/iam/api/account) responsible for this resource. | 
**Sid** | **string**| A 34 character string that uniquely identifies the resource to fetch. | 

### Return type

[**ApiV2010AccountSipSipIpAccessControlList**](api.v2010.account.sip.sip_ip_access_control_list.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchSipIpAccessControlListMapping

> ApiV2010AccountSipSipDomainSipIpAccessControlListMapping FetchSipIpAccessControlListMapping(ctx, AccountSid, DomainSid, Sid)



Fetch an IpAccessControlListMapping resource.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string**| The unique id of the Account that is responsible for this resource. | 
**DomainSid** | **string**| A 34 character string that uniquely identifies the SIP domain. | 
**Sid** | **string**| A 34 character string that uniquely identifies the resource to fetch. | 

### Return type

[**ApiV2010AccountSipSipDomainSipIpAccessControlListMapping**](api.v2010.account.sip.sip_domain.sip_ip_access_control_list_mapping.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchSipIpAddress

> ApiV2010AccountSipSipIpAccessControlListSipIpAddress FetchSipIpAddress(ctx, AccountSid, IpAccessControlListSid, Sid)



Read one IpAddress resource.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string**| The unique id of the [Account](https://www.twilio.com/docs/iam/api/account) responsible for this resource. | 
**IpAccessControlListSid** | **string**| The IpAccessControlList Sid that identifies the IpAddress resources to fetch. | 
**Sid** | **string**| A 34 character string that uniquely identifies the IpAddress resource to fetch. | 

### Return type

[**ApiV2010AccountSipSipIpAccessControlListSipIpAddress**](api.v2010.account.sip.sip_ip_access_control_list.sip_ip_address.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchTranscription

> ApiV2010AccountTranscription FetchTranscription(ctx, AccountSid, Sid)



Fetch an instance of a Transcription

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string**| The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Transcription resource to fetch. | 
**Sid** | **string**| The Twilio-provided string that uniquely identifies the Transcription resource to fetch. | 

### Return type

[**ApiV2010AccountTranscription**](api.v2010.account.transcription.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchUsageTrigger

> ApiV2010AccountUsageUsageTrigger FetchUsageTrigger(ctx, AccountSid, Sid)



Fetch and instance of a usage-trigger

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string**| The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the UsageTrigger resource to fetch. | 
**Sid** | **string**| The Twilio-provided string that uniquely identifies the UsageTrigger resource to fetch. | 

### Return type

[**ApiV2010AccountUsageUsageTrigger**](api.v2010.account.usage.usage_trigger.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAccount

> ListAccountResponse ListAccount(ctx, optional)



Retrieves a collection of Accounts belonging to the account used to make the request

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListAccountOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListAccountOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **FriendlyName** | **optional.String**| Only return the Account resources with friendly names that exactly match this name. | 
 **Status** | **optional.String**| Only return Account resources with the given status. Can be &#x60;closed&#x60;, &#x60;suspended&#x60; or &#x60;active&#x60;. | 
 **PageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

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

> ListAddressResponse ListAddress(ctx, AccountSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string**| The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that is responsible for the Address resource to read. | 
 **optional** | ***ListAddressOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListAddressOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **CustomerName** | **optional.String**| The &#x60;customer_name&#x60; of the Address resources to read. | 
 **FriendlyName** | **optional.String**| The string that identifies the Address resources to read. | 
 **IsoCountry** | **optional.String**| The ISO country code of the Address resources to read. | 
 **PageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

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

> ListApplicationResponse ListApplication(ctx, AccountSid, optional)



Retrieve a list of applications representing an application within the requesting account

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string**| The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Application resources to read. | 
 **optional** | ***ListApplicationOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListApplicationOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **FriendlyName** | **optional.String**| The string that identifies the Application resources to read. | 
 **PageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

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

> ListAuthorizedConnectAppResponse ListAuthorizedConnectApp(ctx, AccountSid, optional)



Retrieve a list of authorized-connect-apps belonging to the account used to make the request

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string**| The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the AuthorizedConnectApp resources to read. | 
 **optional** | ***ListAuthorizedConnectAppOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListAuthorizedConnectAppOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **PageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

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

> ListAvailablePhoneNumberCountryResponse ListAvailablePhoneNumberCountry(ctx, AccountSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string**| The SID of the [Account](https://www.twilio.com/docs/iam/api/account) requesting the available phone number Country resources. | 
 **optional** | ***ListAvailablePhoneNumberCountryOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListAvailablePhoneNumberCountryOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **PageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

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

> ListAvailablePhoneNumberLocalResponse ListAvailablePhoneNumberLocal(ctx, AccountSid, CountryCode, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string**| The SID of the [Account](https://www.twilio.com/docs/iam/api/account) requesting the AvailablePhoneNumber resources. | 
**CountryCode** | **string**| The [ISO-3166-1](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country code of the country from which to read phone numbers. | 
 **optional** | ***ListAvailablePhoneNumberLocalOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListAvailablePhoneNumberLocalOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **AreaCode** | **optional.Int32**| The area code of the phone numbers to read. Applies to only phone numbers in the US and Canada. | 
 **Contains** | **optional.String**| The pattern on which to match phone numbers. Valid characters are &#x60;*&#x60;, &#x60;0-9&#x60;, &#x60;a-z&#x60;, and &#x60;A-Z&#x60;. The &#x60;*&#x60; character matches any single digit. For examples, see [Example 2](https://www.twilio.com/docs/phone-numbers/api/availablephonenumberlocal-resource?code-sample&#x3D;code-find-phone-numbers-by-number-pattern) and [Example 3](https://www.twilio.com/docs/phone-numbers/api/availablephonenumberlocal-resource?code-sample&#x3D;code-find-phone-numbers-by-character-pattern). If specified, this value must have at least two characters. | 
 **SmsEnabled** | **optional.Bool**| Whether the phone numbers can receive text messages. Can be: &#x60;true&#x60; or &#x60;false&#x60;. | 
 **MmsEnabled** | **optional.Bool**| Whether the phone numbers can receive MMS messages. Can be: &#x60;true&#x60; or &#x60;false&#x60;. | 
 **VoiceEnabled** | **optional.Bool**| Whether the phone numbers can receive calls. Can be: &#x60;true&#x60; or &#x60;false&#x60;. | 
 **ExcludeAllAddressRequired** | **optional.Bool**| Whether to exclude phone numbers that require an [Address](https://www.twilio.com/docs/usage/api/address). Can be: &#x60;true&#x60; or &#x60;false&#x60; and the default is &#x60;false&#x60;. | 
 **ExcludeLocalAddressRequired** | **optional.Bool**| Whether to exclude phone numbers that require a local [Address](https://www.twilio.com/docs/usage/api/address). Can be: &#x60;true&#x60; or &#x60;false&#x60; and the default is &#x60;false&#x60;. | 
 **ExcludeForeignAddressRequired** | **optional.Bool**| Whether to exclude phone numbers that require a foreign [Address](https://www.twilio.com/docs/usage/api/address). Can be: &#x60;true&#x60; or &#x60;false&#x60; and the default is &#x60;false&#x60;. | 
 **Beta** | **optional.Bool**| Whether to read phone numbers that are new to the Twilio platform. Can be: &#x60;true&#x60; or &#x60;false&#x60; and the default is &#x60;true&#x60;. | 
 **NearNumber** | **optional.String**| Given a phone number, find a geographically close number within &#x60;distance&#x60; miles. Distance defaults to 25 miles. Applies to only phone numbers in the US and Canada. | 
 **NearLatLong** | **optional.String**| Given a latitude/longitude pair &#x60;lat,long&#x60; find geographically close numbers within &#x60;distance&#x60; miles. Applies to only phone numbers in the US and Canada. | 
 **Distance** | **optional.Int32**| The search radius, in miles, for a &#x60;near_&#x60; query.  Can be up to &#x60;500&#x60; and the default is &#x60;25&#x60;. Applies to only phone numbers in the US and Canada. | 
 **InPostalCode** | **optional.String**| Limit results to a particular postal code. Given a phone number, search within the same postal code as that number. Applies to only phone numbers in the US and Canada. | 
 **InRegion** | **optional.String**| Limit results to a particular region, state, or province. Given a phone number, search within the same region as that number. Applies to only phone numbers in the US and Canada. | 
 **InRateCenter** | **optional.String**| Limit results to a specific rate center, or given a phone number search within the same rate center as that number. Requires &#x60;in_lata&#x60; to be set as well. Applies to only phone numbers in the US and Canada. | 
 **InLata** | **optional.String**| Limit results to a specific local access and transport area ([LATA](https://en.wikipedia.org/wiki/Local_access_and_transport_area)). Given a phone number, search within the same [LATA](https://en.wikipedia.org/wiki/Local_access_and_transport_area) as that number. Applies to only phone numbers in the US and Canada. | 
 **InLocality** | **optional.String**| Limit results to a particular locality or city. Given a phone number, search within the same Locality as that number. | 
 **FaxEnabled** | **optional.Bool**| Whether the phone numbers can receive faxes. Can be: &#x60;true&#x60; or &#x60;false&#x60;. | 
 **PageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

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

> ListAvailablePhoneNumberMachineToMachineResponse ListAvailablePhoneNumberMachineToMachine(ctx, AccountSid, CountryCode, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string**| The SID of the [Account](https://www.twilio.com/docs/iam/api/account) requesting the AvailablePhoneNumber resources. | 
**CountryCode** | **string**| The [ISO-3166-1](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country code of the country from which to read phone numbers. | 
 **optional** | ***ListAvailablePhoneNumberMachineToMachineOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListAvailablePhoneNumberMachineToMachineOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **AreaCode** | **optional.Int32**| The area code of the phone numbers to read. Applies to only phone numbers in the US and Canada. | 
 **Contains** | **optional.String**| The pattern on which to match phone numbers. Valid characters are &#x60;*&#x60;, &#x60;0-9&#x60;, &#x60;a-z&#x60;, and &#x60;A-Z&#x60;. The &#x60;*&#x60; character matches any single digit. For examples, see [Example 2](https://www.twilio.com/docs/phone-numbers/api/availablephonenumber-resource#local-get-basic-example-2) and [Example 3](https://www.twilio.com/docs/phone-numbers/api/availablephonenumber-resource#local-get-basic-example-3). If specified, this value must have at least two characters. | 
 **SmsEnabled** | **optional.Bool**| Whether the phone numbers can receive text messages. Can be: &#x60;true&#x60; or &#x60;false&#x60;. | 
 **MmsEnabled** | **optional.Bool**| Whether the phone numbers can receive MMS messages. Can be: &#x60;true&#x60; or &#x60;false&#x60;. | 
 **VoiceEnabled** | **optional.Bool**| Whether the phone numbers can receive calls. Can be: &#x60;true&#x60; or &#x60;false&#x60;. | 
 **ExcludeAllAddressRequired** | **optional.Bool**| Whether to exclude phone numbers that require an [Address](https://www.twilio.com/docs/usage/api/address). Can be: &#x60;true&#x60; or &#x60;false&#x60; and the default is &#x60;false&#x60;. | 
 **ExcludeLocalAddressRequired** | **optional.Bool**| Whether to exclude phone numbers that require a local [Address](https://www.twilio.com/docs/usage/api/address). Can be: &#x60;true&#x60; or &#x60;false&#x60; and the default is &#x60;false&#x60;. | 
 **ExcludeForeignAddressRequired** | **optional.Bool**| Whether to exclude phone numbers that require a foreign [Address](https://www.twilio.com/docs/usage/api/address). Can be: &#x60;true&#x60; or &#x60;false&#x60; and the default is &#x60;false&#x60;. | 
 **Beta** | **optional.Bool**| Whether to read phone numbers that are new to the Twilio platform. Can be: &#x60;true&#x60; or &#x60;false&#x60; and the default is &#x60;true&#x60;. | 
 **NearNumber** | **optional.String**| Given a phone number, find a geographically close number within &#x60;distance&#x60; miles. Distance defaults to 25 miles. Applies to only phone numbers in the US and Canada. | 
 **NearLatLong** | **optional.String**| Given a latitude/longitude pair &#x60;lat,long&#x60; find geographically close numbers within &#x60;distance&#x60; miles. Applies to only phone numbers in the US and Canada. | 
 **Distance** | **optional.Int32**| The search radius, in miles, for a &#x60;near_&#x60; query.  Can be up to &#x60;500&#x60; and the default is &#x60;25&#x60;. Applies to only phone numbers in the US and Canada. | 
 **InPostalCode** | **optional.String**| Limit results to a particular postal code. Given a phone number, search within the same postal code as that number. Applies to only phone numbers in the US and Canada. | 
 **InRegion** | **optional.String**| Limit results to a particular region, state, or province. Given a phone number, search within the same region as that number. Applies to only phone numbers in the US and Canada. | 
 **InRateCenter** | **optional.String**| Limit results to a specific rate center, or given a phone number search within the same rate center as that number. Requires &#x60;in_lata&#x60; to be set as well. Applies to only phone numbers in the US and Canada. | 
 **InLata** | **optional.String**| Limit results to a specific local access and transport area ([LATA](https://en.wikipedia.org/wiki/Local_access_and_transport_area)). Given a phone number, search within the same [LATA](https://en.wikipedia.org/wiki/Local_access_and_transport_area) as that number. Applies to only phone numbers in the US and Canada. | 
 **InLocality** | **optional.String**| Limit results to a particular locality or city. Given a phone number, search within the same Locality as that number. | 
 **FaxEnabled** | **optional.Bool**| Whether the phone numbers can receive faxes. Can be: &#x60;true&#x60; or &#x60;false&#x60;. | 
 **PageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

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

> ListAvailablePhoneNumberMobileResponse ListAvailablePhoneNumberMobile(ctx, AccountSid, CountryCode, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string**| The SID of the [Account](https://www.twilio.com/docs/iam/api/account) requesting the AvailablePhoneNumber resources. | 
**CountryCode** | **string**| The [ISO-3166-1](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country code of the country from which to read phone numbers. | 
 **optional** | ***ListAvailablePhoneNumberMobileOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListAvailablePhoneNumberMobileOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **AreaCode** | **optional.Int32**| The area code of the phone numbers to read. Applies to only phone numbers in the US and Canada. | 
 **Contains** | **optional.String**| The pattern on which to match phone numbers. Valid characters are &#x60;*&#x60;, &#x60;0-9&#x60;, &#x60;a-z&#x60;, and &#x60;A-Z&#x60;. The &#x60;*&#x60; character matches any single digit. For examples, see [Example 2](https://www.twilio.com/docs/phone-numbers/api/availablephonenumber-resource#local-get-basic-example-2) and [Example 3](https://www.twilio.com/docs/phone-numbers/api/availablephonenumber-resource#local-get-basic-example-3). If specified, this value must have at least two characters. | 
 **SmsEnabled** | **optional.Bool**| Whether the phone numbers can receive text messages. Can be: &#x60;true&#x60; or &#x60;false&#x60;. | 
 **MmsEnabled** | **optional.Bool**| Whether the phone numbers can receive MMS messages. Can be: &#x60;true&#x60; or &#x60;false&#x60;. | 
 **VoiceEnabled** | **optional.Bool**| Whether the phone numbers can receive calls. Can be: &#x60;true&#x60; or &#x60;false&#x60;. | 
 **ExcludeAllAddressRequired** | **optional.Bool**| Whether to exclude phone numbers that require an [Address](https://www.twilio.com/docs/usage/api/address). Can be: &#x60;true&#x60; or &#x60;false&#x60; and the default is &#x60;false&#x60;. | 
 **ExcludeLocalAddressRequired** | **optional.Bool**| Whether to exclude phone numbers that require a local [Address](https://www.twilio.com/docs/usage/api/address). Can be: &#x60;true&#x60; or &#x60;false&#x60; and the default is &#x60;false&#x60;. | 
 **ExcludeForeignAddressRequired** | **optional.Bool**| Whether to exclude phone numbers that require a foreign [Address](https://www.twilio.com/docs/usage/api/address). Can be: &#x60;true&#x60; or &#x60;false&#x60; and the default is &#x60;false&#x60;. | 
 **Beta** | **optional.Bool**| Whether to read phone numbers that are new to the Twilio platform. Can be: &#x60;true&#x60; or &#x60;false&#x60; and the default is &#x60;true&#x60;. | 
 **NearNumber** | **optional.String**| Given a phone number, find a geographically close number within &#x60;distance&#x60; miles. Distance defaults to 25 miles. Applies to only phone numbers in the US and Canada. | 
 **NearLatLong** | **optional.String**| Given a latitude/longitude pair &#x60;lat,long&#x60; find geographically close numbers within &#x60;distance&#x60; miles. Applies to only phone numbers in the US and Canada. | 
 **Distance** | **optional.Int32**| The search radius, in miles, for a &#x60;near_&#x60; query.  Can be up to &#x60;500&#x60; and the default is &#x60;25&#x60;. Applies to only phone numbers in the US and Canada. | 
 **InPostalCode** | **optional.String**| Limit results to a particular postal code. Given a phone number, search within the same postal code as that number. Applies to only phone numbers in the US and Canada. | 
 **InRegion** | **optional.String**| Limit results to a particular region, state, or province. Given a phone number, search within the same region as that number. Applies to only phone numbers in the US and Canada. | 
 **InRateCenter** | **optional.String**| Limit results to a specific rate center, or given a phone number search within the same rate center as that number. Requires &#x60;in_lata&#x60; to be set as well. Applies to only phone numbers in the US and Canada. | 
 **InLata** | **optional.String**| Limit results to a specific local access and transport area ([LATA](https://en.wikipedia.org/wiki/Local_access_and_transport_area)). Given a phone number, search within the same [LATA](https://en.wikipedia.org/wiki/Local_access_and_transport_area) as that number. Applies to only phone numbers in the US and Canada. | 
 **InLocality** | **optional.String**| Limit results to a particular locality or city. Given a phone number, search within the same Locality as that number. | 
 **FaxEnabled** | **optional.Bool**| Whether the phone numbers can receive faxes. Can be: &#x60;true&#x60; or &#x60;false&#x60;. | 
 **PageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

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

> ListAvailablePhoneNumberNationalResponse ListAvailablePhoneNumberNational(ctx, AccountSid, CountryCode, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string**| The SID of the [Account](https://www.twilio.com/docs/iam/api/account) requesting the AvailablePhoneNumber resources. | 
**CountryCode** | **string**| The [ISO-3166-1](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country code of the country from which to read phone numbers. | 
 **optional** | ***ListAvailablePhoneNumberNationalOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListAvailablePhoneNumberNationalOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **AreaCode** | **optional.Int32**| The area code of the phone numbers to read. Applies to only phone numbers in the US and Canada. | 
 **Contains** | **optional.String**| The pattern on which to match phone numbers. Valid characters are &#x60;*&#x60;, &#x60;0-9&#x60;, &#x60;a-z&#x60;, and &#x60;A-Z&#x60;. The &#x60;*&#x60; character matches any single digit. For examples, see [Example 2](https://www.twilio.com/docs/phone-numbers/api/availablephonenumber-resource#local-get-basic-example-2) and [Example 3](https://www.twilio.com/docs/phone-numbers/api/availablephonenumber-resource#local-get-basic-example-3). If specified, this value must have at least two characters. | 
 **SmsEnabled** | **optional.Bool**| Whether the phone numbers can receive text messages. Can be: &#x60;true&#x60; or &#x60;false&#x60;. | 
 **MmsEnabled** | **optional.Bool**| Whether the phone numbers can receive MMS messages. Can be: &#x60;true&#x60; or &#x60;false&#x60;. | 
 **VoiceEnabled** | **optional.Bool**| Whether the phone numbers can receive calls. Can be: &#x60;true&#x60; or &#x60;false&#x60;. | 
 **ExcludeAllAddressRequired** | **optional.Bool**| Whether to exclude phone numbers that require an [Address](https://www.twilio.com/docs/usage/api/address). Can be: &#x60;true&#x60; or &#x60;false&#x60; and the default is &#x60;false&#x60;. | 
 **ExcludeLocalAddressRequired** | **optional.Bool**| Whether to exclude phone numbers that require a local [Address](https://www.twilio.com/docs/usage/api/address). Can be: &#x60;true&#x60; or &#x60;false&#x60; and the default is &#x60;false&#x60;. | 
 **ExcludeForeignAddressRequired** | **optional.Bool**| Whether to exclude phone numbers that require a foreign [Address](https://www.twilio.com/docs/usage/api/address). Can be: &#x60;true&#x60; or &#x60;false&#x60; and the default is &#x60;false&#x60;. | 
 **Beta** | **optional.Bool**| Whether to read phone numbers that are new to the Twilio platform. Can be: &#x60;true&#x60; or &#x60;false&#x60; and the default is &#x60;true&#x60;. | 
 **NearNumber** | **optional.String**| Given a phone number, find a geographically close number within &#x60;distance&#x60; miles. Distance defaults to 25 miles. Applies to only phone numbers in the US and Canada. | 
 **NearLatLong** | **optional.String**| Given a latitude/longitude pair &#x60;lat,long&#x60; find geographically close numbers within &#x60;distance&#x60; miles. Applies to only phone numbers in the US and Canada. | 
 **Distance** | **optional.Int32**| The search radius, in miles, for a &#x60;near_&#x60; query.  Can be up to &#x60;500&#x60; and the default is &#x60;25&#x60;. Applies to only phone numbers in the US and Canada. | 
 **InPostalCode** | **optional.String**| Limit results to a particular postal code. Given a phone number, search within the same postal code as that number. Applies to only phone numbers in the US and Canada. | 
 **InRegion** | **optional.String**| Limit results to a particular region, state, or province. Given a phone number, search within the same region as that number. Applies to only phone numbers in the US and Canada. | 
 **InRateCenter** | **optional.String**| Limit results to a specific rate center, or given a phone number search within the same rate center as that number. Requires &#x60;in_lata&#x60; to be set as well. Applies to only phone numbers in the US and Canada. | 
 **InLata** | **optional.String**| Limit results to a specific local access and transport area ([LATA](https://en.wikipedia.org/wiki/Local_access_and_transport_area)). Given a phone number, search within the same [LATA](https://en.wikipedia.org/wiki/Local_access_and_transport_area) as that number. Applies to only phone numbers in the US and Canada. | 
 **InLocality** | **optional.String**| Limit results to a particular locality or city. Given a phone number, search within the same Locality as that number. | 
 **FaxEnabled** | **optional.Bool**| Whether the phone numbers can receive faxes. Can be: &#x60;true&#x60; or &#x60;false&#x60;. | 
 **PageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

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

> ListAvailablePhoneNumberSharedCostResponse ListAvailablePhoneNumberSharedCost(ctx, AccountSid, CountryCode, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string**| The SID of the [Account](https://www.twilio.com/docs/iam/api/account) requesting the AvailablePhoneNumber resources. | 
**CountryCode** | **string**| The [ISO-3166-1](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country code of the country from which to read phone numbers. | 
 **optional** | ***ListAvailablePhoneNumberSharedCostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListAvailablePhoneNumberSharedCostOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **AreaCode** | **optional.Int32**| The area code of the phone numbers to read. Applies to only phone numbers in the US and Canada. | 
 **Contains** | **optional.String**| The pattern on which to match phone numbers. Valid characters are &#x60;*&#x60;, &#x60;0-9&#x60;, &#x60;a-z&#x60;, and &#x60;A-Z&#x60;. The &#x60;*&#x60; character matches any single digit. For examples, see [Example 2](https://www.twilio.com/docs/phone-numbers/api/availablephonenumber-resource#local-get-basic-example-2) and [Example 3](https://www.twilio.com/docs/phone-numbers/api/availablephonenumber-resource#local-get-basic-example-3). If specified, this value must have at least two characters. | 
 **SmsEnabled** | **optional.Bool**| Whether the phone numbers can receive text messages. Can be: &#x60;true&#x60; or &#x60;false&#x60;. | 
 **MmsEnabled** | **optional.Bool**| Whether the phone numbers can receive MMS messages. Can be: &#x60;true&#x60; or &#x60;false&#x60;. | 
 **VoiceEnabled** | **optional.Bool**| Whether the phone numbers can receive calls. Can be: &#x60;true&#x60; or &#x60;false&#x60;. | 
 **ExcludeAllAddressRequired** | **optional.Bool**| Whether to exclude phone numbers that require an [Address](https://www.twilio.com/docs/usage/api/address). Can be: &#x60;true&#x60; or &#x60;false&#x60; and the default is &#x60;false&#x60;. | 
 **ExcludeLocalAddressRequired** | **optional.Bool**| Whether to exclude phone numbers that require a local [Address](https://www.twilio.com/docs/usage/api/address). Can be: &#x60;true&#x60; or &#x60;false&#x60; and the default is &#x60;false&#x60;. | 
 **ExcludeForeignAddressRequired** | **optional.Bool**| Whether to exclude phone numbers that require a foreign [Address](https://www.twilio.com/docs/usage/api/address). Can be: &#x60;true&#x60; or &#x60;false&#x60; and the default is &#x60;false&#x60;. | 
 **Beta** | **optional.Bool**| Whether to read phone numbers that are new to the Twilio platform. Can be: &#x60;true&#x60; or &#x60;false&#x60; and the default is &#x60;true&#x60;. | 
 **NearNumber** | **optional.String**| Given a phone number, find a geographically close number within &#x60;distance&#x60; miles. Distance defaults to 25 miles. Applies to only phone numbers in the US and Canada. | 
 **NearLatLong** | **optional.String**| Given a latitude/longitude pair &#x60;lat,long&#x60; find geographically close numbers within &#x60;distance&#x60; miles. Applies to only phone numbers in the US and Canada. | 
 **Distance** | **optional.Int32**| The search radius, in miles, for a &#x60;near_&#x60; query.  Can be up to &#x60;500&#x60; and the default is &#x60;25&#x60;. Applies to only phone numbers in the US and Canada. | 
 **InPostalCode** | **optional.String**| Limit results to a particular postal code. Given a phone number, search within the same postal code as that number. Applies to only phone numbers in the US and Canada. | 
 **InRegion** | **optional.String**| Limit results to a particular region, state, or province. Given a phone number, search within the same region as that number. Applies to only phone numbers in the US and Canada. | 
 **InRateCenter** | **optional.String**| Limit results to a specific rate center, or given a phone number search within the same rate center as that number. Requires &#x60;in_lata&#x60; to be set as well. Applies to only phone numbers in the US and Canada. | 
 **InLata** | **optional.String**| Limit results to a specific local access and transport area ([LATA](https://en.wikipedia.org/wiki/Local_access_and_transport_area)). Given a phone number, search within the same [LATA](https://en.wikipedia.org/wiki/Local_access_and_transport_area) as that number. Applies to only phone numbers in the US and Canada. | 
 **InLocality** | **optional.String**| Limit results to a particular locality or city. Given a phone number, search within the same Locality as that number. | 
 **FaxEnabled** | **optional.Bool**| Whether the phone numbers can receive faxes. Can be: &#x60;true&#x60; or &#x60;false&#x60;. | 
 **PageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

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

> ListAvailablePhoneNumberTollFreeResponse ListAvailablePhoneNumberTollFree(ctx, AccountSid, CountryCode, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string**| The SID of the [Account](https://www.twilio.com/docs/iam/api/account) requesting the AvailablePhoneNumber resources. | 
**CountryCode** | **string**| The [ISO-3166-1](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country code of the country from which to read phone numbers. | 
 **optional** | ***ListAvailablePhoneNumberTollFreeOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListAvailablePhoneNumberTollFreeOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **AreaCode** | **optional.Int32**| The area code of the phone numbers to read. Applies to only phone numbers in the US and Canada. | 
 **Contains** | **optional.String**| The pattern on which to match phone numbers. Valid characters are &#x60;*&#x60;, &#x60;0-9&#x60;, &#x60;a-z&#x60;, and &#x60;A-Z&#x60;. The &#x60;*&#x60; character matches any single digit. For examples, see [Example 2](https://www.twilio.com/docs/phone-numbers/api/availablephonenumber-resource#local-get-basic-example-2) and [Example 3](https://www.twilio.com/docs/phone-numbers/api/availablephonenumber-resource#local-get-basic-example-3). If specified, this value must have at least two characters. | 
 **SmsEnabled** | **optional.Bool**| Whether the phone numbers can receive text messages. Can be: &#x60;true&#x60; or &#x60;false&#x60;. | 
 **MmsEnabled** | **optional.Bool**| Whether the phone numbers can receive MMS messages. Can be: &#x60;true&#x60; or &#x60;false&#x60;. | 
 **VoiceEnabled** | **optional.Bool**| Whether the phone numbers can receive calls. Can be: &#x60;true&#x60; or &#x60;false&#x60;. | 
 **ExcludeAllAddressRequired** | **optional.Bool**| Whether to exclude phone numbers that require an [Address](https://www.twilio.com/docs/usage/api/address). Can be: &#x60;true&#x60; or &#x60;false&#x60; and the default is &#x60;false&#x60;. | 
 **ExcludeLocalAddressRequired** | **optional.Bool**| Whether to exclude phone numbers that require a local [Address](https://www.twilio.com/docs/usage/api/address). Can be: &#x60;true&#x60; or &#x60;false&#x60; and the default is &#x60;false&#x60;. | 
 **ExcludeForeignAddressRequired** | **optional.Bool**| Whether to exclude phone numbers that require a foreign [Address](https://www.twilio.com/docs/usage/api/address). Can be: &#x60;true&#x60; or &#x60;false&#x60; and the default is &#x60;false&#x60;. | 
 **Beta** | **optional.Bool**| Whether to read phone numbers that are new to the Twilio platform. Can be: &#x60;true&#x60; or &#x60;false&#x60; and the default is &#x60;true&#x60;. | 
 **NearNumber** | **optional.String**| Given a phone number, find a geographically close number within &#x60;distance&#x60; miles. Distance defaults to 25 miles. Applies to only phone numbers in the US and Canada. | 
 **NearLatLong** | **optional.String**| Given a latitude/longitude pair &#x60;lat,long&#x60; find geographically close numbers within &#x60;distance&#x60; miles. Applies to only phone numbers in the US and Canada. | 
 **Distance** | **optional.Int32**| The search radius, in miles, for a &#x60;near_&#x60; query.  Can be up to &#x60;500&#x60; and the default is &#x60;25&#x60;. Applies to only phone numbers in the US and Canada. | 
 **InPostalCode** | **optional.String**| Limit results to a particular postal code. Given a phone number, search within the same postal code as that number. Applies to only phone numbers in the US and Canada. | 
 **InRegion** | **optional.String**| Limit results to a particular region, state, or province. Given a phone number, search within the same region as that number. Applies to only phone numbers in the US and Canada. | 
 **InRateCenter** | **optional.String**| Limit results to a specific rate center, or given a phone number search within the same rate center as that number. Requires &#x60;in_lata&#x60; to be set as well. Applies to only phone numbers in the US and Canada. | 
 **InLata** | **optional.String**| Limit results to a specific local access and transport area ([LATA](https://en.wikipedia.org/wiki/Local_access_and_transport_area)). Given a phone number, search within the same [LATA](https://en.wikipedia.org/wiki/Local_access_and_transport_area) as that number. Applies to only phone numbers in the US and Canada. | 
 **InLocality** | **optional.String**| Limit results to a particular locality or city. Given a phone number, search within the same Locality as that number. | 
 **FaxEnabled** | **optional.Bool**| Whether the phone numbers can receive faxes. Can be: &#x60;true&#x60; or &#x60;false&#x60;. | 
 **PageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

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

> ListAvailablePhoneNumberVoipResponse ListAvailablePhoneNumberVoip(ctx, AccountSid, CountryCode, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string**| The SID of the [Account](https://www.twilio.com/docs/iam/api/account) requesting the AvailablePhoneNumber resources. | 
**CountryCode** | **string**| The [ISO-3166-1](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country code of the country from which to read phone numbers. | 
 **optional** | ***ListAvailablePhoneNumberVoipOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListAvailablePhoneNumberVoipOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **AreaCode** | **optional.Int32**| The area code of the phone numbers to read. Applies to only phone numbers in the US and Canada. | 
 **Contains** | **optional.String**| The pattern on which to match phone numbers. Valid characters are &#x60;*&#x60;, &#x60;0-9&#x60;, &#x60;a-z&#x60;, and &#x60;A-Z&#x60;. The &#x60;*&#x60; character matches any single digit. For examples, see [Example 2](https://www.twilio.com/docs/phone-numbers/api/availablephonenumber-resource#local-get-basic-example-2) and [Example 3](https://www.twilio.com/docs/phone-numbers/api/availablephonenumber-resource#local-get-basic-example-3). If specified, this value must have at least two characters. | 
 **SmsEnabled** | **optional.Bool**| Whether the phone numbers can receive text messages. Can be: &#x60;true&#x60; or &#x60;false&#x60;. | 
 **MmsEnabled** | **optional.Bool**| Whether the phone numbers can receive MMS messages. Can be: &#x60;true&#x60; or &#x60;false&#x60;. | 
 **VoiceEnabled** | **optional.Bool**| Whether the phone numbers can receive calls. Can be: &#x60;true&#x60; or &#x60;false&#x60;. | 
 **ExcludeAllAddressRequired** | **optional.Bool**| Whether to exclude phone numbers that require an [Address](https://www.twilio.com/docs/usage/api/address). Can be: &#x60;true&#x60; or &#x60;false&#x60; and the default is &#x60;false&#x60;. | 
 **ExcludeLocalAddressRequired** | **optional.Bool**| Whether to exclude phone numbers that require a local [Address](https://www.twilio.com/docs/usage/api/address). Can be: &#x60;true&#x60; or &#x60;false&#x60; and the default is &#x60;false&#x60;. | 
 **ExcludeForeignAddressRequired** | **optional.Bool**| Whether to exclude phone numbers that require a foreign [Address](https://www.twilio.com/docs/usage/api/address). Can be: &#x60;true&#x60; or &#x60;false&#x60; and the default is &#x60;false&#x60;. | 
 **Beta** | **optional.Bool**| Whether to read phone numbers that are new to the Twilio platform. Can be: &#x60;true&#x60; or &#x60;false&#x60; and the default is &#x60;true&#x60;. | 
 **NearNumber** | **optional.String**| Given a phone number, find a geographically close number within &#x60;distance&#x60; miles. Distance defaults to 25 miles. Applies to only phone numbers in the US and Canada. | 
 **NearLatLong** | **optional.String**| Given a latitude/longitude pair &#x60;lat,long&#x60; find geographically close numbers within &#x60;distance&#x60; miles. Applies to only phone numbers in the US and Canada. | 
 **Distance** | **optional.Int32**| The search radius, in miles, for a &#x60;near_&#x60; query.  Can be up to &#x60;500&#x60; and the default is &#x60;25&#x60;. Applies to only phone numbers in the US and Canada. | 
 **InPostalCode** | **optional.String**| Limit results to a particular postal code. Given a phone number, search within the same postal code as that number. Applies to only phone numbers in the US and Canada. | 
 **InRegion** | **optional.String**| Limit results to a particular region, state, or province. Given a phone number, search within the same region as that number. Applies to only phone numbers in the US and Canada. | 
 **InRateCenter** | **optional.String**| Limit results to a specific rate center, or given a phone number search within the same rate center as that number. Requires &#x60;in_lata&#x60; to be set as well. Applies to only phone numbers in the US and Canada. | 
 **InLata** | **optional.String**| Limit results to a specific local access and transport area ([LATA](https://en.wikipedia.org/wiki/Local_access_and_transport_area)). Given a phone number, search within the same [LATA](https://en.wikipedia.org/wiki/Local_access_and_transport_area) as that number. Applies to only phone numbers in the US and Canada. | 
 **InLocality** | **optional.String**| Limit results to a particular locality or city. Given a phone number, search within the same Locality as that number. | 
 **FaxEnabled** | **optional.Bool**| Whether the phone numbers can receive faxes. Can be: &#x60;true&#x60; or &#x60;false&#x60;. | 
 **PageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

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

> ListCallResponse ListCall(ctx, AccountSid, optional)



Retrieves a collection of calls made to and from your account

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string**| The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Call resource(s) to read. | 
 **optional** | ***ListCallOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListCallOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **To** | **optional.String**| Only show calls made to this phone number, SIP address, Client identifier or SIM SID. | 
 **From** | **optional.String**| Only include calls from this phone number, SIP address, Client identifier or SIM SID. | 
 **ParentCallSid** | **optional.String**| Only include calls spawned by calls with this SID. | 
 **Status** | **optional.String**| The status of the calls to include. Can be: &#x60;queued&#x60;, &#x60;ringing&#x60;, &#x60;in-progress&#x60;, &#x60;canceled&#x60;, &#x60;completed&#x60;, &#x60;failed&#x60;, &#x60;busy&#x60;, or &#x60;no-answer&#x60;. | 
 **StartTime** | **optional.Time**| Only include calls that started on this date. Specify a date as &#x60;YYYY-MM-DD&#x60; in GMT, for example: &#x60;2009-07-06&#x60;, to read only calls that started on this date. You can also specify an inequality, such as &#x60;StartTime&lt;&#x3D;YYYY-MM-DD&#x60;, to read calls that started on or before midnight of this date, and &#x60;StartTime&gt;&#x3D;YYYY-MM-DD&#x60; to read calls that started on or after midnight of this date. | 
 **StartTimelesserThan** | **optional.Time**| Only include calls that started on this date. Specify a date as &#x60;YYYY-MM-DD&#x60; in GMT, for example: &#x60;2009-07-06&#x60;, to read only calls that started on this date. You can also specify an inequality, such as &#x60;StartTime&lt;&#x3D;YYYY-MM-DD&#x60;, to read calls that started on or before midnight of this date, and &#x60;StartTime&gt;&#x3D;YYYY-MM-DD&#x60; to read calls that started on or after midnight of this date. | 
 **StartTimegreaterThan** | **optional.Time**| Only include calls that started on this date. Specify a date as &#x60;YYYY-MM-DD&#x60; in GMT, for example: &#x60;2009-07-06&#x60;, to read only calls that started on this date. You can also specify an inequality, such as &#x60;StartTime&lt;&#x3D;YYYY-MM-DD&#x60;, to read calls that started on or before midnight of this date, and &#x60;StartTime&gt;&#x3D;YYYY-MM-DD&#x60; to read calls that started on or after midnight of this date. | 
 **EndTime** | **optional.Time**| Only include calls that ended on this date. Specify a date as &#x60;YYYY-MM-DD&#x60; in GMT, for example: &#x60;2009-07-06&#x60;, to read only calls that ended on this date. You can also specify an inequality, such as &#x60;EndTime&lt;&#x3D;YYYY-MM-DD&#x60;, to read calls that ended on or before midnight of this date, and &#x60;EndTime&gt;&#x3D;YYYY-MM-DD&#x60; to read calls that ended on or after midnight of this date. | 
 **EndTimelesserThan** | **optional.Time**| Only include calls that ended on this date. Specify a date as &#x60;YYYY-MM-DD&#x60; in GMT, for example: &#x60;2009-07-06&#x60;, to read only calls that ended on this date. You can also specify an inequality, such as &#x60;EndTime&lt;&#x3D;YYYY-MM-DD&#x60;, to read calls that ended on or before midnight of this date, and &#x60;EndTime&gt;&#x3D;YYYY-MM-DD&#x60; to read calls that ended on or after midnight of this date. | 
 **EndTimegreaterThan** | **optional.Time**| Only include calls that ended on this date. Specify a date as &#x60;YYYY-MM-DD&#x60; in GMT, for example: &#x60;2009-07-06&#x60;, to read only calls that ended on this date. You can also specify an inequality, such as &#x60;EndTime&lt;&#x3D;YYYY-MM-DD&#x60;, to read calls that ended on or before midnight of this date, and &#x60;EndTime&gt;&#x3D;YYYY-MM-DD&#x60; to read calls that ended on or after midnight of this date. | 
 **PageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

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

> ListCallEventResponse ListCallEvent(ctx, AccountSid, CallSid, optional)



Retrieve a list of all events for a call.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string**| The unique SID identifier of the Account. | 
**CallSid** | **string**| The unique SID identifier of the Call. | 
 **optional** | ***ListCallEventOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListCallEventOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **PageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

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

> ListCallNotificationResponse ListCallNotification(ctx, AccountSid, CallSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string**| The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Call Notification resources to read. | 
**CallSid** | **string**| The [Call](https://www.twilio.com/docs/voice/api/call-resource) SID of the Call Notification resources to read. | 
 **optional** | ***ListCallNotificationOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListCallNotificationOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **Log** | **optional.Int32**| Only read notifications of the specified log level. Can be:  &#x60;0&#x60; to read only ERROR notifications or &#x60;1&#x60; to read only WARNING notifications. By default, all notifications are read. | 
 **MessageDate** | **optional.Time**| Only show notifications for the specified date, formatted as &#x60;YYYY-MM-DD&#x60;. You can also specify an inequality, such as &#x60;&lt;&#x3D;YYYY-MM-DD&#x60; for messages logged at or before midnight on a date, or &#x60;&gt;&#x3D;YYYY-MM-DD&#x60; for messages logged at or after midnight on a date. | 
 **MessageDatelesserThan** | **optional.Time**| Only show notifications for the specified date, formatted as &#x60;YYYY-MM-DD&#x60;. You can also specify an inequality, such as &#x60;&lt;&#x3D;YYYY-MM-DD&#x60; for messages logged at or before midnight on a date, or &#x60;&gt;&#x3D;YYYY-MM-DD&#x60; for messages logged at or after midnight on a date. | 
 **MessageDategreaterThan** | **optional.Time**| Only show notifications for the specified date, formatted as &#x60;YYYY-MM-DD&#x60;. You can also specify an inequality, such as &#x60;&lt;&#x3D;YYYY-MM-DD&#x60; for messages logged at or before midnight on a date, or &#x60;&gt;&#x3D;YYYY-MM-DD&#x60; for messages logged at or after midnight on a date. | 
 **PageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

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

> ListCallRecordingResponse ListCallRecording(ctx, AccountSid, CallSid, optional)



Retrieve a list of recordings belonging to the call used to make the request

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string**| The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Recording resources to read. | 
**CallSid** | **string**| The [Call](https://www.twilio.com/docs/voice/api/call-resource) SID of the resources to read. | 
 **optional** | ***ListCallRecordingOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListCallRecordingOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **DateCreated** | **optional.Time**| The &#x60;date_created&#x60; value, specified as &#x60;YYYY-MM-DD&#x60;, of the resources to read. You can also specify inequality: &#x60;DateCreated&lt;&#x3D;YYYY-MM-DD&#x60; will return recordings generated at or before midnight on a given date, and &#x60;DateCreated&gt;&#x3D;YYYY-MM-DD&#x60; returns recordings generated at or after midnight on a date. | 
 **DateCreatedlesserThan** | **optional.Time**| The &#x60;date_created&#x60; value, specified as &#x60;YYYY-MM-DD&#x60;, of the resources to read. You can also specify inequality: &#x60;DateCreated&lt;&#x3D;YYYY-MM-DD&#x60; will return recordings generated at or before midnight on a given date, and &#x60;DateCreated&gt;&#x3D;YYYY-MM-DD&#x60; returns recordings generated at or after midnight on a date. | 
 **DateCreatedgreaterThan** | **optional.Time**| The &#x60;date_created&#x60; value, specified as &#x60;YYYY-MM-DD&#x60;, of the resources to read. You can also specify inequality: &#x60;DateCreated&lt;&#x3D;YYYY-MM-DD&#x60; will return recordings generated at or before midnight on a given date, and &#x60;DateCreated&gt;&#x3D;YYYY-MM-DD&#x60; returns recordings generated at or after midnight on a date. | 
 **PageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

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

> ListConferenceResponse ListConference(ctx, AccountSid, optional)



Retrieve a list of conferences belonging to the account used to make the request

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string**| The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Conference resource(s) to read. | 
 **optional** | ***ListConferenceOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListConferenceOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **DateCreated** | **optional.Time**| The &#x60;date_created&#x60; value, specified as &#x60;YYYY-MM-DD&#x60;, of the resources to read. To read conferences that started on or before midnight on a date, use &#x60;&lt;&#x3D;YYYY-MM-DD&#x60;, and to specify  conferences that started on or after midnight on a date, use &#x60;&gt;&#x3D;YYYY-MM-DD&#x60;. | 
 **DateCreatedlesserThan** | **optional.Time**| The &#x60;date_created&#x60; value, specified as &#x60;YYYY-MM-DD&#x60;, of the resources to read. To read conferences that started on or before midnight on a date, use &#x60;&lt;&#x3D;YYYY-MM-DD&#x60;, and to specify  conferences that started on or after midnight on a date, use &#x60;&gt;&#x3D;YYYY-MM-DD&#x60;. | 
 **DateCreatedgreaterThan** | **optional.Time**| The &#x60;date_created&#x60; value, specified as &#x60;YYYY-MM-DD&#x60;, of the resources to read. To read conferences that started on or before midnight on a date, use &#x60;&lt;&#x3D;YYYY-MM-DD&#x60;, and to specify  conferences that started on or after midnight on a date, use &#x60;&gt;&#x3D;YYYY-MM-DD&#x60;. | 
 **DateUpdated** | **optional.Time**| The &#x60;date_updated&#x60; value, specified as &#x60;YYYY-MM-DD&#x60;, of the resources to read. To read conferences that were last updated on or before midnight on a date, use &#x60;&lt;&#x3D;YYYY-MM-DD&#x60;, and to specify conferences that were last updated on or after midnight on a given date, use  &#x60;&gt;&#x3D;YYYY-MM-DD&#x60;. | 
 **DateUpdatedlesserThan** | **optional.Time**| The &#x60;date_updated&#x60; value, specified as &#x60;YYYY-MM-DD&#x60;, of the resources to read. To read conferences that were last updated on or before midnight on a date, use &#x60;&lt;&#x3D;YYYY-MM-DD&#x60;, and to specify conferences that were last updated on or after midnight on a given date, use  &#x60;&gt;&#x3D;YYYY-MM-DD&#x60;. | 
 **DateUpdatedgreaterThan** | **optional.Time**| The &#x60;date_updated&#x60; value, specified as &#x60;YYYY-MM-DD&#x60;, of the resources to read. To read conferences that were last updated on or before midnight on a date, use &#x60;&lt;&#x3D;YYYY-MM-DD&#x60;, and to specify conferences that were last updated on or after midnight on a given date, use  &#x60;&gt;&#x3D;YYYY-MM-DD&#x60;. | 
 **FriendlyName** | **optional.String**| The string that identifies the Conference resources to read. | 
 **Status** | **optional.String**| The status of the resources to read. Can be: &#x60;init&#x60;, &#x60;in-progress&#x60;, or &#x60;completed&#x60;. | 
 **PageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

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

> ListConferenceRecordingResponse ListConferenceRecording(ctx, AccountSid, ConferenceSid, optional)



Retrieve a list of recordings belonging to the call used to make the request

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string**| The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Conference Recording resources to read. | 
**ConferenceSid** | **string**| The Conference SID that identifies the conference associated with the recording to read. | 
 **optional** | ***ListConferenceRecordingOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListConferenceRecordingOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **DateCreated** | **optional.Time**| The &#x60;date_created&#x60; value, specified as &#x60;YYYY-MM-DD&#x60;, of the resources to read. You can also specify inequality: &#x60;DateCreated&lt;&#x3D;YYYY-MM-DD&#x60; will return recordings generated at or before midnight on a given date, and &#x60;DateCreated&gt;&#x3D;YYYY-MM-DD&#x60; returns recordings generated at or after midnight on a date. | 
 **DateCreatedlesserThan** | **optional.Time**| The &#x60;date_created&#x60; value, specified as &#x60;YYYY-MM-DD&#x60;, of the resources to read. You can also specify inequality: &#x60;DateCreated&lt;&#x3D;YYYY-MM-DD&#x60; will return recordings generated at or before midnight on a given date, and &#x60;DateCreated&gt;&#x3D;YYYY-MM-DD&#x60; returns recordings generated at or after midnight on a date. | 
 **DateCreatedgreaterThan** | **optional.Time**| The &#x60;date_created&#x60; value, specified as &#x60;YYYY-MM-DD&#x60;, of the resources to read. You can also specify inequality: &#x60;DateCreated&lt;&#x3D;YYYY-MM-DD&#x60; will return recordings generated at or before midnight on a given date, and &#x60;DateCreated&gt;&#x3D;YYYY-MM-DD&#x60; returns recordings generated at or after midnight on a date. | 
 **PageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

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

> ListConnectAppResponse ListConnectApp(ctx, AccountSid, optional)



Retrieve a list of connect-apps belonging to the account used to make the request

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string**| The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the ConnectApp resources to read. | 
 **optional** | ***ListConnectAppOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListConnectAppOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **PageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

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

> ListDependentPhoneNumberResponse ListDependentPhoneNumber(ctx, AccountSid, AddressSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string**| The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the DependentPhoneNumber resources to read. | 
**AddressSid** | **string**| The SID of the Address resource associated with the phone number. | 
 **optional** | ***ListDependentPhoneNumberOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListDependentPhoneNumberOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **PageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

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

> ListIncomingPhoneNumberResponse ListIncomingPhoneNumber(ctx, AccountSid, optional)



Retrieve a list of incoming-phone-numbers belonging to the account used to make the request.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string**| The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the IncomingPhoneNumber resources to read. | 
 **optional** | ***ListIncomingPhoneNumberOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListIncomingPhoneNumberOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **Beta** | **optional.Bool**| Whether to include phone numbers new to the Twilio platform. Can be: &#x60;true&#x60; or &#x60;false&#x60; and the default is &#x60;true&#x60;. | 
 **FriendlyName** | **optional.String**| A string that identifies the IncomingPhoneNumber resources to read. | 
 **PhoneNumber** | **optional.String**| The phone numbers of the IncomingPhoneNumber resources to read. You can specify partial numbers and use &#39;*&#39; as a wildcard for any digit. | 
 **Origin** | **optional.String**| Whether to include phone numbers based on their origin. Can be: &#x60;twilio&#x60; or &#x60;hosted&#x60;. By default, phone numbers of all origin are included. | 
 **PageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

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

> ListIncomingPhoneNumberAssignedAddOnResponse ListIncomingPhoneNumberAssignedAddOn(ctx, AccountSid, ResourceSid, optional)



Retrieve a list of Add-on installations currently assigned to this Number.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string**| The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the resources to read. | 
**ResourceSid** | **string**| The SID of the Phone Number to which the Add-on is assigned. | 
 **optional** | ***ListIncomingPhoneNumberAssignedAddOnOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListIncomingPhoneNumberAssignedAddOnOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **PageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

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

> ListIncomingPhoneNumberAssignedAddOnExtensionResponse ListIncomingPhoneNumberAssignedAddOnExtension(ctx, AccountSid, ResourceSid, AssignedAddOnSid, optional)



Retrieve a list of Extensions for the Assigned Add-on.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string**| The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the resources to read. | 
**ResourceSid** | **string**| The SID of the Phone Number to which the Add-on is assigned. | 
**AssignedAddOnSid** | **string**| The SID that uniquely identifies the assigned Add-on installation. | 
 **optional** | ***ListIncomingPhoneNumberAssignedAddOnExtensionOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListIncomingPhoneNumberAssignedAddOnExtensionOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **PageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

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

> ListIncomingPhoneNumberLocalResponse ListIncomingPhoneNumberLocal(ctx, AccountSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string**| The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the resources to read. | 
 **optional** | ***ListIncomingPhoneNumberLocalOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListIncomingPhoneNumberLocalOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **Beta** | **optional.Bool**| Whether to include phone numbers new to the Twilio platform. Can be: &#x60;true&#x60; or &#x60;false&#x60; and the default is &#x60;true&#x60;. | 
 **FriendlyName** | **optional.String**| A string that identifies the resources to read. | 
 **PhoneNumber** | **optional.String**| The phone numbers of the IncomingPhoneNumber resources to read. You can specify partial numbers and use &#39;*&#39; as a wildcard for any digit. | 
 **Origin** | **optional.String**| Whether to include phone numbers based on their origin. Can be: &#x60;twilio&#x60; or &#x60;hosted&#x60;. By default, phone numbers of all origin are included. | 
 **PageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

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

> ListIncomingPhoneNumberMobileResponse ListIncomingPhoneNumberMobile(ctx, AccountSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string**| The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the resources to read. | 
 **optional** | ***ListIncomingPhoneNumberMobileOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListIncomingPhoneNumberMobileOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **Beta** | **optional.Bool**| Whether to include phone numbers new to the Twilio platform. Can be: &#x60;true&#x60; or &#x60;false&#x60; and the default is &#x60;true&#x60;. | 
 **FriendlyName** | **optional.String**| A string that identifies the resources to read. | 
 **PhoneNumber** | **optional.String**| The phone numbers of the IncomingPhoneNumber resources to read. You can specify partial numbers and use &#39;*&#39; as a wildcard for any digit. | 
 **Origin** | **optional.String**| Whether to include phone numbers based on their origin. Can be: &#x60;twilio&#x60; or &#x60;hosted&#x60;. By default, phone numbers of all origin are included. | 
 **PageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

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

> ListIncomingPhoneNumberTollFreeResponse ListIncomingPhoneNumberTollFree(ctx, AccountSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string**| The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the resources to read. | 
 **optional** | ***ListIncomingPhoneNumberTollFreeOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListIncomingPhoneNumberTollFreeOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **Beta** | **optional.Bool**| Whether to include phone numbers new to the Twilio platform. Can be: &#x60;true&#x60; or &#x60;false&#x60; and the default is &#x60;true&#x60;. | 
 **FriendlyName** | **optional.String**| A string that identifies the resources to read. | 
 **PhoneNumber** | **optional.String**| The phone numbers of the IncomingPhoneNumber resources to read. You can specify partial numbers and use &#39;*&#39; as a wildcard for any digit. | 
 **Origin** | **optional.String**| Whether to include phone numbers based on their origin. Can be: &#x60;twilio&#x60; or &#x60;hosted&#x60;. By default, phone numbers of all origin are included. | 
 **PageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

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

> ListKeyResponse ListKey(ctx, AccountSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string**| The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Key resources to read. | 
 **optional** | ***ListKeyOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListKeyOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **PageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

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

> ListMediaResponse ListMedia(ctx, AccountSid, MessageSid, optional)



Retrieve a list of Media resources belonging to the account used to make the request

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string**| The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Media resource(s) to read. | 
**MessageSid** | **string**| The SID of the Message resource that this Media resource belongs to. | 
 **optional** | ***ListMediaOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListMediaOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **DateCreated** | **optional.Time**| Only include media that was created on this date. Specify a date as &#x60;YYYY-MM-DD&#x60; in GMT, for example: &#x60;2009-07-06&#x60;, to read media that was created on this date. You can also specify an inequality, such as &#x60;StartTime&lt;&#x3D;YYYY-MM-DD&#x60;, to read media that was created on or before midnight of this date, and &#x60;StartTime&gt;&#x3D;YYYY-MM-DD&#x60; to read media that was created on or after midnight of this date. | 
 **DateCreatedlesserThan** | **optional.Time**| Only include media that was created on this date. Specify a date as &#x60;YYYY-MM-DD&#x60; in GMT, for example: &#x60;2009-07-06&#x60;, to read media that was created on this date. You can also specify an inequality, such as &#x60;StartTime&lt;&#x3D;YYYY-MM-DD&#x60;, to read media that was created on or before midnight of this date, and &#x60;StartTime&gt;&#x3D;YYYY-MM-DD&#x60; to read media that was created on or after midnight of this date. | 
 **DateCreatedgreaterThan** | **optional.Time**| Only include media that was created on this date. Specify a date as &#x60;YYYY-MM-DD&#x60; in GMT, for example: &#x60;2009-07-06&#x60;, to read media that was created on this date. You can also specify an inequality, such as &#x60;StartTime&lt;&#x3D;YYYY-MM-DD&#x60;, to read media that was created on or before midnight of this date, and &#x60;StartTime&gt;&#x3D;YYYY-MM-DD&#x60; to read media that was created on or after midnight of this date. | 
 **PageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

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

> ListMemberResponse ListMember(ctx, AccountSid, QueueSid, optional)



Retrieve the members of the queue

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string**| The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Member resource(s) to read. | 
**QueueSid** | **string**| The SID of the Queue in which to find the members | 
 **optional** | ***ListMemberOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListMemberOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **PageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

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

> ListMessageResponse ListMessage(ctx, AccountSid, optional)



Retrieve a list of messages belonging to the account used to make the request

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string**| The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Message resources to read. | 
 **optional** | ***ListMessageOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListMessageOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **To** | **optional.String**| Read messages sent to only this phone number. | 
 **From** | **optional.String**| Read messages sent from only this phone number or alphanumeric sender ID. | 
 **DateSent** | **optional.Time**| The date of the messages to show. Specify a date as &#x60;YYYY-MM-DD&#x60; in GMT to read only messages sent on this date. For example: &#x60;2009-07-06&#x60;. You can also specify an inequality, such as &#x60;DateSent&lt;&#x3D;YYYY-MM-DD&#x60;, to read messages sent on or before midnight on a date, and &#x60;DateSent&gt;&#x3D;YYYY-MM-DD&#x60; to read messages sent on or after midnight on a date. | 
 **DateSentlesserThan** | **optional.Time**| The date of the messages to show. Specify a date as &#x60;YYYY-MM-DD&#x60; in GMT to read only messages sent on this date. For example: &#x60;2009-07-06&#x60;. You can also specify an inequality, such as &#x60;DateSent&lt;&#x3D;YYYY-MM-DD&#x60;, to read messages sent on or before midnight on a date, and &#x60;DateSent&gt;&#x3D;YYYY-MM-DD&#x60; to read messages sent on or after midnight on a date. | 
 **DateSentgreaterThan** | **optional.Time**| The date of the messages to show. Specify a date as &#x60;YYYY-MM-DD&#x60; in GMT to read only messages sent on this date. For example: &#x60;2009-07-06&#x60;. You can also specify an inequality, such as &#x60;DateSent&lt;&#x3D;YYYY-MM-DD&#x60;, to read messages sent on or before midnight on a date, and &#x60;DateSent&gt;&#x3D;YYYY-MM-DD&#x60; to read messages sent on or after midnight on a date. | 
 **PageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

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

> ListNotificationResponse ListNotification(ctx, AccountSid, optional)



Retrieve a list of notifications belonging to the account used to make the request

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string**| The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Notification resources to read. | 
 **optional** | ***ListNotificationOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListNotificationOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **Log** | **optional.Int32**| Only read notifications of the specified log level. Can be:  &#x60;0&#x60; to read only ERROR notifications or &#x60;1&#x60; to read only WARNING notifications. By default, all notifications are read. | 
 **MessageDate** | **optional.Time**| Only show notifications for the specified date, formatted as &#x60;YYYY-MM-DD&#x60;. You can also specify an inequality, such as &#x60;&lt;&#x3D;YYYY-MM-DD&#x60; for messages logged at or before midnight on a date, or &#x60;&gt;&#x3D;YYYY-MM-DD&#x60; for messages logged at or after midnight on a date. | 
 **MessageDatelesserThan** | **optional.Time**| Only show notifications for the specified date, formatted as &#x60;YYYY-MM-DD&#x60;. You can also specify an inequality, such as &#x60;&lt;&#x3D;YYYY-MM-DD&#x60; for messages logged at or before midnight on a date, or &#x60;&gt;&#x3D;YYYY-MM-DD&#x60; for messages logged at or after midnight on a date. | 
 **MessageDategreaterThan** | **optional.Time**| Only show notifications for the specified date, formatted as &#x60;YYYY-MM-DD&#x60;. You can also specify an inequality, such as &#x60;&lt;&#x3D;YYYY-MM-DD&#x60; for messages logged at or before midnight on a date, or &#x60;&gt;&#x3D;YYYY-MM-DD&#x60; for messages logged at or after midnight on a date. | 
 **PageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

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

> ListOutgoingCallerIdResponse ListOutgoingCallerId(ctx, AccountSid, optional)



Retrieve a list of outgoing-caller-ids belonging to the account used to make the request

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string**| The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the OutgoingCallerId resources to read. | 
 **optional** | ***ListOutgoingCallerIdOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListOutgoingCallerIdOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **PhoneNumber** | **optional.String**| The phone number of the OutgoingCallerId resources to read. | 
 **FriendlyName** | **optional.String**| The string that identifies the OutgoingCallerId resources to read. | 
 **PageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

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

> ListParticipantResponse ListParticipant(ctx, AccountSid, ConferenceSid, optional)



Retrieve a list of participants belonging to the account used to make the request

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string**| The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Participant resources to read. | 
**ConferenceSid** | **string**| The SID of the conference with the participants to read. | 
 **optional** | ***ListParticipantOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListParticipantOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **Muted** | **optional.Bool**| Whether to return only participants that are muted. Can be: &#x60;true&#x60; or &#x60;false&#x60;. | 
 **Hold** | **optional.Bool**| Whether to return only participants that are on hold. Can be: &#x60;true&#x60; or &#x60;false&#x60;. | 
 **Coaching** | **optional.Bool**| Whether to return only participants who are coaching another call. Can be: &#x60;true&#x60; or &#x60;false&#x60;. | 
 **PageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

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

> ListQueueResponse ListQueue(ctx, AccountSid, optional)



Retrieve a list of queues belonging to the account used to make the request

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string**| The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Queue resources to read. | 
 **optional** | ***ListQueueOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListQueueOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **PageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

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

> ListRecordingResponse ListRecording(ctx, AccountSid, optional)



Retrieve a list of recordings belonging to the account used to make the request

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string**| The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Recording resources to read. | 
 **optional** | ***ListRecordingOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListRecordingOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **DateCreated** | **optional.Time**| Only include recordings that were created on this date. Specify a date as &#x60;YYYY-MM-DD&#x60; in GMT, for example: &#x60;2009-07-06&#x60;, to read recordings that were created on this date. You can also specify an inequality, such as &#x60;DateCreated&lt;&#x3D;YYYY-MM-DD&#x60;, to read recordings that were created on or before midnight of this date, and &#x60;DateCreated&gt;&#x3D;YYYY-MM-DD&#x60; to read recordings that were created on or after midnight of this date. | 
 **DateCreatedlesserThan** | **optional.Time**| Only include recordings that were created on this date. Specify a date as &#x60;YYYY-MM-DD&#x60; in GMT, for example: &#x60;2009-07-06&#x60;, to read recordings that were created on this date. You can also specify an inequality, such as &#x60;DateCreated&lt;&#x3D;YYYY-MM-DD&#x60;, to read recordings that were created on or before midnight of this date, and &#x60;DateCreated&gt;&#x3D;YYYY-MM-DD&#x60; to read recordings that were created on or after midnight of this date. | 
 **DateCreatedgreaterThan** | **optional.Time**| Only include recordings that were created on this date. Specify a date as &#x60;YYYY-MM-DD&#x60; in GMT, for example: &#x60;2009-07-06&#x60;, to read recordings that were created on this date. You can also specify an inequality, such as &#x60;DateCreated&lt;&#x3D;YYYY-MM-DD&#x60;, to read recordings that were created on or before midnight of this date, and &#x60;DateCreated&gt;&#x3D;YYYY-MM-DD&#x60; to read recordings that were created on or after midnight of this date. | 
 **CallSid** | **optional.String**| The [Call](https://www.twilio.com/docs/voice/api/call-resource) SID of the resources to read. | 
 **ConferenceSid** | **optional.String**| The Conference SID that identifies the conference associated with the recording to read. | 
 **PageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

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

> ListRecordingAddOnResultResponse ListRecordingAddOnResult(ctx, AccountSid, ReferenceSid, optional)



Retrieve a list of results belonging to the recording

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string**| The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Recording AddOnResult resources to read. | 
**ReferenceSid** | **string**| The SID of the recording to which the result to read belongs. | 
 **optional** | ***ListRecordingAddOnResultOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListRecordingAddOnResultOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **PageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

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

> ListRecordingAddOnResultPayloadResponse ListRecordingAddOnResultPayload(ctx, AccountSid, ReferenceSid, AddOnResultSid, optional)



Retrieve a list of payloads belonging to the AddOnResult

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string**| The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Recording AddOnResult Payload resources to read. | 
**ReferenceSid** | **string**| The SID of the recording to which the AddOnResult resource that contains the payloads to read belongs. | 
**AddOnResultSid** | **string**| The SID of the AddOnResult to which the payloads to read belongs. | 
 **optional** | ***ListRecordingAddOnResultPayloadOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListRecordingAddOnResultPayloadOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **PageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

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

> ListRecordingTranscriptionResponse ListRecordingTranscription(ctx, AccountSid, RecordingSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string**| The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Transcription resources to read. | 
**RecordingSid** | **string**| The SID of the [Recording](https://www.twilio.com/docs/voice/api/recording) that created the transcriptions to read. | 
 **optional** | ***ListRecordingTranscriptionOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListRecordingTranscriptionOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **PageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

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

> ListShortCodeResponse ListShortCode(ctx, AccountSid, optional)



Retrieve a list of short-codes belonging to the account used to make the request

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string**| The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the ShortCode resource(s) to read. | 
 **optional** | ***ListShortCodeOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListShortCodeOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **FriendlyName** | **optional.String**| The string that identifies the ShortCode resources to read. | 
 **ShortCode** | **optional.String**| Only show the ShortCode resources that match this pattern. You can specify partial numbers and use &#39;*&#39; as a wildcard for any digit. | 
 **PageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

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

> ListSigningKeyResponse ListSigningKey(ctx, AccountSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string**|  | 
 **optional** | ***ListSigningKeyOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListSigningKeyOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **PageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

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

> ListSipAuthCallsCredentialListMappingResponse ListSipAuthCallsCredentialListMapping(ctx, AccountSid, DomainSid, optional)



Retrieve a list of credential list mappings belonging to the domain used in the request

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string**| The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the CredentialListMapping resources to read. | 
**DomainSid** | **string**| The SID of the SIP domain that contains the resources to read. | 
 **optional** | ***ListSipAuthCallsCredentialListMappingOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListSipAuthCallsCredentialListMappingOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **PageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

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

> ListSipAuthCallsIpAccessControlListMappingResponse ListSipAuthCallsIpAccessControlListMapping(ctx, AccountSid, DomainSid, optional)



Retrieve a list of IP Access Control List mappings belonging to the domain used in the request

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string**| The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the IpAccessControlListMapping resources to read. | 
**DomainSid** | **string**| The SID of the SIP domain that contains the resources to read. | 
 **optional** | ***ListSipAuthCallsIpAccessControlListMappingOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListSipAuthCallsIpAccessControlListMappingOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **PageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

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

> ListSipAuthRegistrationsCredentialListMappingResponse ListSipAuthRegistrationsCredentialListMapping(ctx, AccountSid, DomainSid, optional)



Retrieve a list of credential list mappings belonging to the domain used in the request

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string**| The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the CredentialListMapping resources to read. | 
**DomainSid** | **string**| The SID of the SIP domain that contains the resources to read. | 
 **optional** | ***ListSipAuthRegistrationsCredentialListMappingOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListSipAuthRegistrationsCredentialListMappingOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **PageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

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

> ListSipCredentialResponse ListSipCredential(ctx, AccountSid, CredentialListSid, optional)



Retrieve a list of credentials.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string**| The unique id of the Account that is responsible for this resource. | 
**CredentialListSid** | **string**| The unique id that identifies the credential list that contains the desired credentials. | 
 **optional** | ***ListSipCredentialOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListSipCredentialOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **PageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

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

> ListSipCredentialListResponse ListSipCredentialList(ctx, AccountSid, optional)



Get All Credential Lists

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string**| The unique id of the Account that is responsible for this resource. | 
 **optional** | ***ListSipCredentialListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListSipCredentialListOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **PageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

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

> ListSipCredentialListMappingResponse ListSipCredentialListMapping(ctx, AccountSid, DomainSid, optional)



Read multiple CredentialListMapping resources from an account.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string**| The unique id of the [Account](https://www.twilio.com/docs/iam/api/account) responsible for this resource. | 
**DomainSid** | **string**| A 34 character string that uniquely identifies the SIP Domain that includes the resource to read. | 
 **optional** | ***ListSipCredentialListMappingOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListSipCredentialListMappingOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **PageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

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

> ListSipDomainResponse ListSipDomain(ctx, AccountSid, optional)



Retrieve a list of domains belonging to the account used to make the request

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string**| The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the SipDomain resources to read. | 
 **optional** | ***ListSipDomainOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListSipDomainOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **PageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

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

> ListSipIpAccessControlListResponse ListSipIpAccessControlList(ctx, AccountSid, optional)



Retrieve a list of IpAccessControlLists that belong to the account used to make the request

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string**| The unique id of the [Account](https://www.twilio.com/docs/iam/api/account) responsible for this resource. | 
 **optional** | ***ListSipIpAccessControlListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListSipIpAccessControlListOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **PageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

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

> ListSipIpAccessControlListMappingResponse ListSipIpAccessControlListMapping(ctx, AccountSid, DomainSid, optional)



Retrieve a list of IpAccessControlListMapping resources.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string**| The unique id of the Account that is responsible for this resource. | 
**DomainSid** | **string**| A 34 character string that uniquely identifies the SIP domain. | 
 **optional** | ***ListSipIpAccessControlListMappingOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListSipIpAccessControlListMappingOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **PageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

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

> ListSipIpAddressResponse ListSipIpAddress(ctx, AccountSid, IpAccessControlListSid, optional)



Read multiple IpAddress resources.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string**| The unique id of the [Account](https://www.twilio.com/docs/iam/api/account) responsible for this resource. | 
**IpAccessControlListSid** | **string**| The IpAccessControlList Sid that identifies the IpAddress resources to read. | 
 **optional** | ***ListSipIpAddressOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListSipIpAddressOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **PageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

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

> ListTranscriptionResponse ListTranscription(ctx, AccountSid, optional)



Retrieve a list of transcriptions belonging to the account used to make the request

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string**| The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Transcription resources to read. | 
 **optional** | ***ListTranscriptionOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListTranscriptionOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **PageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

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

> ListUsageRecordResponse ListUsageRecord(ctx, AccountSid, optional)



Retrieve a list of usage-records belonging to the account used to make the request

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string**| The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the UsageRecord resources to read. | 
 **optional** | ***ListUsageRecordOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListUsageRecordOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **Category** | **optional.String**| The [usage category](https://www.twilio.com/docs/usage/api/usage-record#usage-categories) of the UsageRecord resources to read. Only UsageRecord resources in the specified category are retrieved. | 
 **StartDate** | **optional.Time**| Only include usage that has occurred on or after this date. Specify the date in GMT and format as &#x60;YYYY-MM-DD&#x60;. You can also specify offsets from the current date, such as: &#x60;-30days&#x60;, which will set the start date to be 30 days before the current date. | 
 **EndDate** | **optional.Time**| Only include usage that occurred on or before this date. Specify the date in GMT and format as &#x60;YYYY-MM-DD&#x60;.  You can also specify offsets from the current date, such as: &#x60;+30days&#x60;, which will set the end date to 30 days from the current date. | 
 **IncludeSubaccounts** | **optional.Bool**| Whether to include usage from the master account and all its subaccounts. Can be: &#x60;true&#x60; (the default) to include usage from the master account and all subaccounts or &#x60;false&#x60; to retrieve usage from only the specified account. | 
 **PageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

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

> ListUsageRecordAllTimeResponse ListUsageRecordAllTime(ctx, AccountSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string**| The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the UsageRecord resources to read. | 
 **optional** | ***ListUsageRecordAllTimeOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListUsageRecordAllTimeOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **Category** | **optional.String**| The [usage category](https://www.twilio.com/docs/usage/api/usage-record#usage-categories) of the UsageRecord resources to read. Only UsageRecord resources in the specified category are retrieved. | 
 **StartDate** | **optional.Time**| Only include usage that has occurred on or after this date. Specify the date in GMT and format as &#x60;YYYY-MM-DD&#x60;. You can also specify offsets from the current date, such as: &#x60;-30days&#x60;, which will set the start date to be 30 days before the current date. | 
 **EndDate** | **optional.Time**| Only include usage that occurred on or before this date. Specify the date in GMT and format as &#x60;YYYY-MM-DD&#x60;.  You can also specify offsets from the current date, such as: &#x60;+30days&#x60;, which will set the end date to 30 days from the current date. | 
 **IncludeSubaccounts** | **optional.Bool**| Whether to include usage from the master account and all its subaccounts. Can be: &#x60;true&#x60; (the default) to include usage from the master account and all subaccounts or &#x60;false&#x60; to retrieve usage from only the specified account. | 
 **PageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

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

> ListUsageRecordDailyResponse ListUsageRecordDaily(ctx, AccountSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string**| The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the UsageRecord resources to read. | 
 **optional** | ***ListUsageRecordDailyOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListUsageRecordDailyOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **Category** | **optional.String**| The [usage category](https://www.twilio.com/docs/usage/api/usage-record#usage-categories) of the UsageRecord resources to read. Only UsageRecord resources in the specified category are retrieved. | 
 **StartDate** | **optional.Time**| Only include usage that has occurred on or after this date. Specify the date in GMT and format as &#x60;YYYY-MM-DD&#x60;. You can also specify offsets from the current date, such as: &#x60;-30days&#x60;, which will set the start date to be 30 days before the current date. | 
 **EndDate** | **optional.Time**| Only include usage that occurred on or before this date. Specify the date in GMT and format as &#x60;YYYY-MM-DD&#x60;.  You can also specify offsets from the current date, such as: &#x60;+30days&#x60;, which will set the end date to 30 days from the current date. | 
 **IncludeSubaccounts** | **optional.Bool**| Whether to include usage from the master account and all its subaccounts. Can be: &#x60;true&#x60; (the default) to include usage from the master account and all subaccounts or &#x60;false&#x60; to retrieve usage from only the specified account. | 
 **PageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

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

> ListUsageRecordLastMonthResponse ListUsageRecordLastMonth(ctx, AccountSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string**| The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the UsageRecord resources to read. | 
 **optional** | ***ListUsageRecordLastMonthOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListUsageRecordLastMonthOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **Category** | **optional.String**| The [usage category](https://www.twilio.com/docs/usage/api/usage-record#usage-categories) of the UsageRecord resources to read. Only UsageRecord resources in the specified category are retrieved. | 
 **StartDate** | **optional.Time**| Only include usage that has occurred on or after this date. Specify the date in GMT and format as &#x60;YYYY-MM-DD&#x60;. You can also specify offsets from the current date, such as: &#x60;-30days&#x60;, which will set the start date to be 30 days before the current date. | 
 **EndDate** | **optional.Time**| Only include usage that occurred on or before this date. Specify the date in GMT and format as &#x60;YYYY-MM-DD&#x60;.  You can also specify offsets from the current date, such as: &#x60;+30days&#x60;, which will set the end date to 30 days from the current date. | 
 **IncludeSubaccounts** | **optional.Bool**| Whether to include usage from the master account and all its subaccounts. Can be: &#x60;true&#x60; (the default) to include usage from the master account and all subaccounts or &#x60;false&#x60; to retrieve usage from only the specified account. | 
 **PageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

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

> ListUsageRecordMonthlyResponse ListUsageRecordMonthly(ctx, AccountSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string**| The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the UsageRecord resources to read. | 
 **optional** | ***ListUsageRecordMonthlyOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListUsageRecordMonthlyOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **Category** | **optional.String**| The [usage category](https://www.twilio.com/docs/usage/api/usage-record#usage-categories) of the UsageRecord resources to read. Only UsageRecord resources in the specified category are retrieved. | 
 **StartDate** | **optional.Time**| Only include usage that has occurred on or after this date. Specify the date in GMT and format as &#x60;YYYY-MM-DD&#x60;. You can also specify offsets from the current date, such as: &#x60;-30days&#x60;, which will set the start date to be 30 days before the current date. | 
 **EndDate** | **optional.Time**| Only include usage that occurred on or before this date. Specify the date in GMT and format as &#x60;YYYY-MM-DD&#x60;.  You can also specify offsets from the current date, such as: &#x60;+30days&#x60;, which will set the end date to 30 days from the current date. | 
 **IncludeSubaccounts** | **optional.Bool**| Whether to include usage from the master account and all its subaccounts. Can be: &#x60;true&#x60; (the default) to include usage from the master account and all subaccounts or &#x60;false&#x60; to retrieve usage from only the specified account. | 
 **PageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

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

> ListUsageRecordThisMonthResponse ListUsageRecordThisMonth(ctx, AccountSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string**| The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the UsageRecord resources to read. | 
 **optional** | ***ListUsageRecordThisMonthOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListUsageRecordThisMonthOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **Category** | **optional.String**| The [usage category](https://www.twilio.com/docs/usage/api/usage-record#usage-categories) of the UsageRecord resources to read. Only UsageRecord resources in the specified category are retrieved. | 
 **StartDate** | **optional.Time**| Only include usage that has occurred on or after this date. Specify the date in GMT and format as &#x60;YYYY-MM-DD&#x60;. You can also specify offsets from the current date, such as: &#x60;-30days&#x60;, which will set the start date to be 30 days before the current date. | 
 **EndDate** | **optional.Time**| Only include usage that occurred on or before this date. Specify the date in GMT and format as &#x60;YYYY-MM-DD&#x60;.  You can also specify offsets from the current date, such as: &#x60;+30days&#x60;, which will set the end date to 30 days from the current date. | 
 **IncludeSubaccounts** | **optional.Bool**| Whether to include usage from the master account and all its subaccounts. Can be: &#x60;true&#x60; (the default) to include usage from the master account and all subaccounts or &#x60;false&#x60; to retrieve usage from only the specified account. | 
 **PageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

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

> ListUsageRecordTodayResponse ListUsageRecordToday(ctx, AccountSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string**| The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the UsageRecord resources to read. | 
 **optional** | ***ListUsageRecordTodayOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListUsageRecordTodayOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **Category** | **optional.String**| The [usage category](https://www.twilio.com/docs/usage/api/usage-record#usage-categories) of the UsageRecord resources to read. Only UsageRecord resources in the specified category are retrieved. | 
 **StartDate** | **optional.Time**| Only include usage that has occurred on or after this date. Specify the date in GMT and format as &#x60;YYYY-MM-DD&#x60;. You can also specify offsets from the current date, such as: &#x60;-30days&#x60;, which will set the start date to be 30 days before the current date. | 
 **EndDate** | **optional.Time**| Only include usage that occurred on or before this date. Specify the date in GMT and format as &#x60;YYYY-MM-DD&#x60;.  You can also specify offsets from the current date, such as: &#x60;+30days&#x60;, which will set the end date to 30 days from the current date. | 
 **IncludeSubaccounts** | **optional.Bool**| Whether to include usage from the master account and all its subaccounts. Can be: &#x60;true&#x60; (the default) to include usage from the master account and all subaccounts or &#x60;false&#x60; to retrieve usage from only the specified account. | 
 **PageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

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

> ListUsageRecordYearlyResponse ListUsageRecordYearly(ctx, AccountSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string**| The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the UsageRecord resources to read. | 
 **optional** | ***ListUsageRecordYearlyOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListUsageRecordYearlyOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **Category** | **optional.String**| The [usage category](https://www.twilio.com/docs/usage/api/usage-record#usage-categories) of the UsageRecord resources to read. Only UsageRecord resources in the specified category are retrieved. | 
 **StartDate** | **optional.Time**| Only include usage that has occurred on or after this date. Specify the date in GMT and format as &#x60;YYYY-MM-DD&#x60;. You can also specify offsets from the current date, such as: &#x60;-30days&#x60;, which will set the start date to be 30 days before the current date. | 
 **EndDate** | **optional.Time**| Only include usage that occurred on or before this date. Specify the date in GMT and format as &#x60;YYYY-MM-DD&#x60;.  You can also specify offsets from the current date, such as: &#x60;+30days&#x60;, which will set the end date to 30 days from the current date. | 
 **IncludeSubaccounts** | **optional.Bool**| Whether to include usage from the master account and all its subaccounts. Can be: &#x60;true&#x60; (the default) to include usage from the master account and all subaccounts or &#x60;false&#x60; to retrieve usage from only the specified account. | 
 **PageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

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

> ListUsageRecordYesterdayResponse ListUsageRecordYesterday(ctx, AccountSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string**| The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the UsageRecord resources to read. | 
 **optional** | ***ListUsageRecordYesterdayOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListUsageRecordYesterdayOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **Category** | **optional.String**| The [usage category](https://www.twilio.com/docs/usage/api/usage-record#usage-categories) of the UsageRecord resources to read. Only UsageRecord resources in the specified category are retrieved. | 
 **StartDate** | **optional.Time**| Only include usage that has occurred on or after this date. Specify the date in GMT and format as &#x60;YYYY-MM-DD&#x60;. You can also specify offsets from the current date, such as: &#x60;-30days&#x60;, which will set the start date to be 30 days before the current date. | 
 **EndDate** | **optional.Time**| Only include usage that occurred on or before this date. Specify the date in GMT and format as &#x60;YYYY-MM-DD&#x60;.  You can also specify offsets from the current date, such as: &#x60;+30days&#x60;, which will set the end date to 30 days from the current date. | 
 **IncludeSubaccounts** | **optional.Bool**| Whether to include usage from the master account and all its subaccounts. Can be: &#x60;true&#x60; (the default) to include usage from the master account and all subaccounts or &#x60;false&#x60; to retrieve usage from only the specified account. | 
 **PageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

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

> ListUsageTriggerResponse ListUsageTrigger(ctx, AccountSid, optional)



Retrieve a list of usage-triggers belonging to the account used to make the request

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string**| The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the UsageTrigger resources to read. | 
 **optional** | ***ListUsageTriggerOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListUsageTriggerOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **Recurring** | **optional.String**| The frequency of recurring UsageTriggers to read. Can be: &#x60;daily&#x60;, &#x60;monthly&#x60;, or &#x60;yearly&#x60; to read recurring UsageTriggers. An empty value or a value of &#x60;alltime&#x60; reads non-recurring UsageTriggers. | 
 **TriggerBy** | **optional.String**| The trigger field of the UsageTriggers to read.  Can be: &#x60;count&#x60;, &#x60;usage&#x60;, or &#x60;price&#x60; as described in the [UsageRecords documentation](https://www.twilio.com/docs/usage/api/usage-record#usage-count-price). | 
 **UsageCategory** | **optional.String**| The usage category of the UsageTriggers to read. Must be a supported [usage categories](https://www.twilio.com/docs/usage/api/usage-record#usage-categories). | 
 **PageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

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

> ApiV2010Account UpdateAccount(ctx, Sid, optional)



Modify the properties of a given Account

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string**| The Account Sid that uniquely identifies the account to update | 
 **optional** | ***UpdateAccountOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateAccountOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **FriendlyName** | **optional.String**| Update the human-readable description of this Account | 
 **Status** | **optional.String**| Alter the status of this account: use &#x60;closed&#x60; to irreversibly close this account, &#x60;suspended&#x60; to temporarily suspend it, or &#x60;active&#x60; to reactivate it. | 

### Return type

[**ApiV2010Account**](api.v2010.account.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAddress

> ApiV2010AccountAddress UpdateAddress(ctx, AccountSid, Sid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string**| The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that is responsible for the Address resource to update. | 
**Sid** | **string**| The Twilio-provided string that uniquely identifies the Address resource to update. | 
 **optional** | ***UpdateAddressOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateAddressOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **AutoCorrectAddress** | **optional.Bool**| Whether we should automatically correct the address. Can be: &#x60;true&#x60; or &#x60;false&#x60; and the default is &#x60;true&#x60;. If empty or &#x60;true&#x60;, we will correct the address you provide if necessary. If &#x60;false&#x60;, we won&#39;t alter the address you provide. | 
 **City** | **optional.String**| The city of the address. | 
 **CustomerName** | **optional.String**| The name to associate with the address. | 
 **EmergencyEnabled** | **optional.Bool**| Whether to enable emergency calling on the address. Can be: &#x60;true&#x60; or &#x60;false&#x60;. | 
 **FriendlyName** | **optional.String**| A descriptive string that you create to describe the address. It can be up to 64 characters long. | 
 **PostalCode** | **optional.String**| The postal code of the address. | 
 **Region** | **optional.String**| The state or region of the address. | 
 **Street** | **optional.String**| The number and street address of the address. | 

### Return type

[**ApiV2010AccountAddress**](api.v2010.account.address.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateApplication

> ApiV2010AccountApplication UpdateApplication(ctx, AccountSid, Sid, optional)



Updates the application's properties

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string**| The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Application resources to update. | 
**Sid** | **string**| The Twilio-provided string that uniquely identifies the Application resource to update. | 
 **optional** | ***UpdateApplicationOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateApplicationOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **ApiVersion** | **optional.String**| The API version to use to start a new TwiML session. Can be: &#x60;2010-04-01&#x60; or &#x60;2008-08-01&#x60;. The default value is your account&#39;s default API version. | 
 **FriendlyName** | **optional.String**| A descriptive string that you create to describe the resource. It can be up to 64 characters long. | 
 **MessageStatusCallback** | **optional.String**| The URL we should call using a POST method to send message status information to your application. | 
 **SmsFallbackMethod** | **optional.String**| The HTTP method we should use to call &#x60;sms_fallback_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60;. | 
 **SmsFallbackUrl** | **optional.String**| The URL that we should call when an error occurs while retrieving or executing the TwiML from &#x60;sms_url&#x60;. | 
 **SmsMethod** | **optional.String**| The HTTP method we should use to call &#x60;sms_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60;. | 
 **SmsStatusCallback** | **optional.String**| Same as message_status_callback: The URL we should call using a POST method to send status information about SMS messages sent by the application. Deprecated, included for backwards compatibility. | 
 **SmsUrl** | **optional.String**| The URL we should call when the phone number receives an incoming SMS message. | 
 **StatusCallback** | **optional.String**| The URL we should call using the &#x60;status_callback_method&#x60; to send status information to your application. | 
 **StatusCallbackMethod** | **optional.String**| The HTTP method we should use to call &#x60;status_callback&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60;. | 
 **VoiceCallerIdLookup** | **optional.Bool**| Whether we should look up the caller&#39;s caller-ID name from the CNAM database (additional charges apply). Can be: &#x60;true&#x60; or &#x60;false&#x60;. | 
 **VoiceFallbackMethod** | **optional.String**| The HTTP method we should use to call &#x60;voice_fallback_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60;. | 
 **VoiceFallbackUrl** | **optional.String**| The URL that we should call when an error occurs retrieving or executing the TwiML requested by &#x60;url&#x60;. | 
 **VoiceMethod** | **optional.String**| The HTTP method we should use to call &#x60;voice_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60;. | 
 **VoiceUrl** | **optional.String**| The URL we should call when the phone number assigned to this application receives a call. | 

### Return type

[**ApiV2010AccountApplication**](api.v2010.account.application.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCall

> ApiV2010AccountCall UpdateCall(ctx, AccountSid, Sid, optional)



Initiates a call redirect or terminates a call

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string**| The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Call resource(s) to update. | 
**Sid** | **string**| The Twilio-provided string that uniquely identifies the Call resource to update | 
 **optional** | ***UpdateCallOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateCallOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **FallbackMethod** | **optional.String**| The HTTP method that we should use to request the &#x60;fallback_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60; and the default is &#x60;POST&#x60;. If an &#x60;application_sid&#x60; parameter is present, this parameter is ignored. | 
 **FallbackUrl** | **optional.String**| The URL that we call using the &#x60;fallback_method&#x60; if an error occurs when requesting or executing the TwiML at &#x60;url&#x60;. If an &#x60;application_sid&#x60; parameter is present, this parameter is ignored. | 
 **Method** | **optional.String**| The HTTP method we should use when calling the &#x60;url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60; and the default is &#x60;POST&#x60;. If an &#x60;application_sid&#x60; parameter is present, this parameter is ignored. | 
 **Status** | **optional.String**| The new status of the resource. Can be: &#x60;canceled&#x60; or &#x60;completed&#x60;. Specifying &#x60;canceled&#x60; will attempt to hang up calls that are queued or ringing; however, it will not affect calls already in progress. Specifying &#x60;completed&#x60; will attempt to hang up a call even if it&#39;s already in progress. | 
 **StatusCallback** | **optional.String**| The URL we should call using the &#x60;status_callback_method&#x60; to send status information to your application. If no &#x60;status_callback_event&#x60; is specified, we will send the &#x60;completed&#x60; status. If an &#x60;application_sid&#x60; parameter is present, this parameter is ignored. URLs must contain a valid hostname (underscores are not permitted). | 
 **StatusCallbackMethod** | **optional.String**| The HTTP method we should use when requesting the &#x60;status_callback&#x60; URL. Can be: &#x60;GET&#x60; or &#x60;POST&#x60; and the default is &#x60;POST&#x60;. If an &#x60;application_sid&#x60; parameter is present, this parameter is ignored. | 
 **Twiml** | **optional.String**| TwiML instructions for the call Twilio will use without fetching Twiml from url. Twiml and url parameters are mutually exclusive | 
 **Url** | **optional.String**| The absolute URL that returns the TwiML instructions for the call. We will call this URL using the &#x60;method&#x60; when the call connects. For more information, see the [Url Parameter](https://www.twilio.com/docs/voice/make-calls#specify-a-url-parameter) section in [Making Calls](https://www.twilio.com/docs/voice/make-calls). | 

### Return type

[**ApiV2010AccountCall**](api.v2010.account.call.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCallFeedback

> ApiV2010AccountCallCallFeedback UpdateCallFeedback(ctx, AccountSid, CallSid, optional)



Update a Feedback resource for a call

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string**| The unique id of the [Account](https://www.twilio.com/docs/iam/api/account) responsible for this resource. | 
**CallSid** | **string**| The call sid that uniquely identifies the call | 
 **optional** | ***UpdateCallFeedbackOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateCallFeedbackOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **Issue** | [**optional.Interface of []string**](string.md)| One or more issues experienced during the call. The issues can be: &#x60;imperfect-audio&#x60;, &#x60;dropped-call&#x60;, &#x60;incorrect-caller-id&#x60;, &#x60;post-dial-delay&#x60;, &#x60;digits-not-captured&#x60;, &#x60;audio-latency&#x60;, &#x60;unsolicited-call&#x60;, or &#x60;one-way-audio&#x60;. | 
 **QualityScore** | **optional.Int32**| The call quality expressed as an integer from &#x60;1&#x60; to &#x60;5&#x60; where &#x60;1&#x60; represents very poor call quality and &#x60;5&#x60; represents a perfect call. | 

### Return type

[**ApiV2010AccountCallCallFeedback**](api.v2010.account.call.call_feedback.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCallRecording

> ApiV2010AccountCallCallRecording UpdateCallRecording(ctx, AccountSid, CallSid, Sid, optional)



Changes the status of the recording to paused, stopped, or in-progress. Note: Pass `Twilio.CURRENT` instead of recording sid to reference current active recording.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string**| The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Recording resource to update. | 
**CallSid** | **string**| The [Call](https://www.twilio.com/docs/voice/api/call-resource) SID of the resource to update. | 
**Sid** | **string**| The Twilio-provided string that uniquely identifies the Recording resource to update. | 
 **optional** | ***UpdateCallRecordingOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateCallRecordingOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **PauseBehavior** | **optional.String**| Whether to record during a pause. Can be: &#x60;skip&#x60; or &#x60;silence&#x60; and the default is &#x60;silence&#x60;. &#x60;skip&#x60; does not record during the pause period, while &#x60;silence&#x60; will replace the actual audio of the call with silence during the pause period. This parameter only applies when setting &#x60;status&#x60; is set to &#x60;paused&#x60;. | 
 **Status** | **optional.String**| The new status of the recording. Can be: &#x60;stopped&#x60;, &#x60;paused&#x60;, &#x60;in-progress&#x60;. | 

### Return type

[**ApiV2010AccountCallCallRecording**](api.v2010.account.call.call_recording.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateConference

> ApiV2010AccountConference UpdateConference(ctx, AccountSid, Sid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string**| The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Conference resource(s) to update. | 
**Sid** | **string**| The Twilio-provided string that uniquely identifies the Conference resource to update | 
 **optional** | ***UpdateConferenceOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateConferenceOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **AnnounceMethod** | **optional.String**| The HTTP method used to call &#x60;announce_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60; and the default is &#x60;POST&#x60; | 
 **AnnounceUrl** | **optional.String**| The URL we should call to announce something into the conference. The URL can return an MP3, a WAV, or a TwiML document with &#x60;&lt;Play&gt;&#x60; or &#x60;&lt;Say&gt;&#x60;. | 
 **Status** | **optional.String**| The new status of the resource. Can be:  Can be: &#x60;init&#x60;, &#x60;in-progress&#x60;, or &#x60;completed&#x60;. Specifying &#x60;completed&#x60; will end the conference and hang up all participants | 

### Return type

[**ApiV2010AccountConference**](api.v2010.account.conference.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateConferenceRecording

> ApiV2010AccountConferenceConferenceRecording UpdateConferenceRecording(ctx, AccountSid, ConferenceSid, Sid, optional)



Changes the status of the recording to paused, stopped, or in-progress. Note: To use `Twilio.CURRENT`, pass it as recording sid.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string**| The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Conference Recording resource to update. | 
**ConferenceSid** | **string**| The Conference SID that identifies the conference associated with the recording to update. | 
**Sid** | **string**| The Twilio-provided string that uniquely identifies the Conference Recording resource to update. Use &#x60;Twilio.CURRENT&#x60; to reference the current active recording. | 
 **optional** | ***UpdateConferenceRecordingOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateConferenceRecordingOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **PauseBehavior** | **optional.String**| Whether to record during a pause. Can be: &#x60;skip&#x60; or &#x60;silence&#x60; and the default is &#x60;silence&#x60;. &#x60;skip&#x60; does not record during the pause period, while &#x60;silence&#x60; will replace the actual audio of the call with silence during the pause period. This parameter only applies when setting &#x60;status&#x60; is set to &#x60;paused&#x60;. | 
 **Status** | **optional.String**| The new status of the recording. Can be: &#x60;stopped&#x60;, &#x60;paused&#x60;, &#x60;in-progress&#x60;. | 

### Return type

[**ApiV2010AccountConferenceConferenceRecording**](api.v2010.account.conference.conference_recording.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateConnectApp

> ApiV2010AccountConnectApp UpdateConnectApp(ctx, AccountSid, Sid, optional)



Update a connect-app with the specified parameters

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string**| The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the ConnectApp resources to update. | 
**Sid** | **string**| The Twilio-provided string that uniquely identifies the ConnectApp resource to update. | 
 **optional** | ***UpdateConnectAppOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateConnectAppOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **AuthorizeRedirectUrl** | **optional.String**| The URL to redirect the user to after we authenticate the user and obtain authorization to access the Connect App. | 
 **CompanyName** | **optional.String**| The company name to set for the Connect App. | 
 **DeauthorizeCallbackMethod** | **optional.String**| The HTTP method to use when calling &#x60;deauthorize_callback_url&#x60;. | 
 **DeauthorizeCallbackUrl** | **optional.String**| The URL to call using the &#x60;deauthorize_callback_method&#x60; to de-authorize the Connect App. | 
 **Description** | **optional.String**| A description of the Connect App. | 
 **FriendlyName** | **optional.String**| A descriptive string that you create to describe the resource. It can be up to 64 characters long. | 
 **HomepageUrl** | **optional.String**| A public URL where users can obtain more information about this Connect App. | 
 **Permissions** | [**optional.Interface of []string**](string.md)| A comma-separated list of the permissions you will request from the users of this ConnectApp.  Can include: &#x60;get-all&#x60; and &#x60;post-all&#x60;. | 

### Return type

[**ApiV2010AccountConnectApp**](api.v2010.account.connect_app.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateIncomingPhoneNumber

> ApiV2010AccountIncomingPhoneNumber UpdateIncomingPhoneNumber(ctx, AccountSid, Sid, optional)



Update an incoming-phone-number instance.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string**| The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the IncomingPhoneNumber resource to update.  For more information, see [Exchanging Numbers Between Subaccounts](https://www.twilio.com/docs/iam/api/subaccounts#exchanging-numbers). | 
**Sid** | **string**| The Twilio-provided string that uniquely identifies the IncomingPhoneNumber resource to update. | 
 **optional** | ***UpdateIncomingPhoneNumberOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateIncomingPhoneNumberOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **AccountSid** | **optional.String**| The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the IncomingPhoneNumber resource to update.  For more information, see [Exchanging Numbers Between Subaccounts](https://www.twilio.com/docs/iam/api/subaccounts#exchanging-numbers). | 
 **AddressSid** | **optional.String**| The SID of the Address resource we should associate with the phone number. Some regions require addresses to meet local regulations. | 
 **ApiVersion** | **optional.String**| The API version to use for incoming calls made to the phone number. The default is &#x60;2010-04-01&#x60;. | 
 **BundleSid** | **optional.String**| The SID of the Bundle resource that you associate with the phone number. Some regions require a Bundle to meet local Regulations. | 
 **EmergencyAddressSid** | **optional.String**| The SID of the emergency address configuration to use for emergency calling from this phone number. | 
 **EmergencyStatus** | **optional.String**| The configuration status parameter that determines whether the phone number is enabled for emergency calling. | 
 **FriendlyName** | **optional.String**| A descriptive string that you created to describe this phone number. It can be up to 64 characters long. By default, this is a formatted version of the phone number. | 
 **IdentitySid** | **optional.String**| The SID of the Identity resource that we should associate with the phone number. Some regions require an identity to meet local regulations. | 
 **SmsApplicationSid** | **optional.String**| The SID of the application that should handle SMS messages sent to the number. If an &#x60;sms_application_sid&#x60; is present, we ignore all of the &#x60;sms_*_url&#x60; urls and use those set on the application. | 
 **SmsFallbackMethod** | **optional.String**| The HTTP method that we should use to call &#x60;sms_fallback_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60; and defaults to &#x60;POST&#x60;. | 
 **SmsFallbackUrl** | **optional.String**| The URL that we should call when an error occurs while requesting or executing the TwiML defined by &#x60;sms_url&#x60;. | 
 **SmsMethod** | **optional.String**| The HTTP method that we should use to call &#x60;sms_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60; and defaults to &#x60;POST&#x60;. | 
 **SmsUrl** | **optional.String**| The URL we should call when the phone number receives an incoming SMS message. | 
 **StatusCallback** | **optional.String**| The URL we should call using the &#x60;status_callback_method&#x60; to send status information to your application. | 
 **StatusCallbackMethod** | **optional.String**| The HTTP method we should use to call &#x60;status_callback&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60; and defaults to &#x60;POST&#x60;. | 
 **TrunkSid** | **optional.String**| The SID of the Trunk we should use to handle phone calls to the phone number. If a &#x60;trunk_sid&#x60; is present, we ignore all of the voice urls and voice applications and use only those set on the Trunk. Setting a &#x60;trunk_sid&#x60; will automatically delete your &#x60;voice_application_sid&#x60; and vice versa. | 
 **VoiceApplicationSid** | **optional.String**| The SID of the application we should use to handle phone calls to the phone number. If a &#x60;voice_application_sid&#x60; is present, we ignore all of the voice urls and use only those set on the application. Setting a &#x60;voice_application_sid&#x60; will automatically delete your &#x60;trunk_sid&#x60; and vice versa. | 
 **VoiceCallerIdLookup** | **optional.Bool**| Whether to lookup the caller&#39;s name from the CNAM database and post it to your app. Can be: &#x60;true&#x60; or &#x60;false&#x60; and defaults to &#x60;false&#x60;. | 
 **VoiceFallbackMethod** | **optional.String**| The HTTP method that we should use to call &#x60;voice_fallback_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60; and defaults to &#x60;POST&#x60;. | 
 **VoiceFallbackUrl** | **optional.String**| The URL that we should call when an error occurs retrieving or executing the TwiML requested by &#x60;url&#x60;. | 
 **VoiceMethod** | **optional.String**| The HTTP method that we should use to call &#x60;voice_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60; and defaults to &#x60;POST&#x60;. | 
 **VoiceReceiveMode** | **optional.String**| The configuration parameter for the phone number to receive incoming voice calls or faxes. Can be: &#x60;fax&#x60; or &#x60;voice&#x60; and defaults to &#x60;voice&#x60;. | 
 **VoiceUrl** | **optional.String**| The URL that we should call to answer a call to the phone number. The &#x60;voice_url&#x60; will not be called if a &#x60;voice_application_sid&#x60; or a &#x60;trunk_sid&#x60; is set. | 

### Return type

[**ApiV2010AccountIncomingPhoneNumber**](api.v2010.account.incoming_phone_number.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateKey

> ApiV2010AccountKey UpdateKey(ctx, AccountSid, Sid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string**| The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Key resources to update. | 
**Sid** | **string**| The Twilio-provided string that uniquely identifies the Key resource to update. | 
 **optional** | ***UpdateKeyOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateKeyOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **FriendlyName** | **optional.String**| A descriptive string that you create to describe the resource. It can be up to 64 characters long. | 

### Return type

[**ApiV2010AccountKey**](api.v2010.account.key.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMember

> ApiV2010AccountQueueMember UpdateMember(ctx, AccountSid, QueueSid, CallSid, optional)



Dequeue a member from a queue and have the member's call begin executing the TwiML document at that URL

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string**| The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Member resource(s) to update. | 
**QueueSid** | **string**| The SID of the Queue in which to find the members to update. | 
**CallSid** | **string**| The [Call](https://www.twilio.com/docs/voice/api/call-resource) SID of the resource(s) to update. | 
 **optional** | ***UpdateMemberOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateMemberOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **Method** | **optional.String**| How to pass the update request data. Can be &#x60;GET&#x60; or &#x60;POST&#x60; and the default is &#x60;POST&#x60;. &#x60;POST&#x60; sends the data as encoded form data and &#x60;GET&#x60; sends the data as query parameters. | 
 **Url** | **optional.String**| The absolute URL of the Queue resource. | 

### Return type

[**ApiV2010AccountQueueMember**](api.v2010.account.queue.member.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMessage

> ApiV2010AccountMessage UpdateMessage(ctx, AccountSid, Sid, optional)



To redact a message-body from a post-flight message record, post to the message instance resource with an empty body

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string**| The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Message resources to update. | 
**Sid** | **string**| The Twilio-provided string that uniquely identifies the Message resource to update. | 
 **optional** | ***UpdateMessageOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateMessageOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **Body** | **optional.String**| The text of the message you want to send. Can be up to 1,600 characters long. | 

### Return type

[**ApiV2010AccountMessage**](api.v2010.account.message.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOutgoingCallerId

> ApiV2010AccountOutgoingCallerId UpdateOutgoingCallerId(ctx, AccountSid, Sid, optional)



Updates the caller-id

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string**| The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the OutgoingCallerId resources to update. | 
**Sid** | **string**| The Twilio-provided string that uniquely identifies the OutgoingCallerId resource to update. | 
 **optional** | ***UpdateOutgoingCallerIdOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateOutgoingCallerIdOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **FriendlyName** | **optional.String**| A descriptive string that you create to describe the resource. It can be up to 64 characters long. | 

### Return type

[**ApiV2010AccountOutgoingCallerId**](api.v2010.account.outgoing_caller_id.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateParticipant

> ApiV2010AccountConferenceParticipant UpdateParticipant(ctx, AccountSid, ConferenceSid, CallSid, optional)



Update the properties of the participant

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string**| The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Participant resources to update. | 
**ConferenceSid** | **string**| The SID of the conference with the participant to update. | 
**CallSid** | **string**| The [Call](https://www.twilio.com/docs/voice/api/call-resource) SID or label of the participant to update. Non URL safe characters in a label must be percent encoded, for example, a space character is represented as %20. | 
 **optional** | ***UpdateParticipantOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateParticipantOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **AnnounceMethod** | **optional.String**| The HTTP method we should use to call &#x60;announce_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60; and defaults to &#x60;POST&#x60;. | 
 **AnnounceUrl** | **optional.String**| The URL we call using the &#x60;announce_method&#x60; for an announcement to the participant. The URL must return an MP3 file, a WAV file, or a TwiML document that contains &#x60;&lt;Play&gt;&#x60; or &#x60;&lt;Say&gt;&#x60; commands. | 
 **BeepOnExit** | **optional.Bool**| Whether to play a notification beep to the conference when the participant exits. Can be: &#x60;true&#x60; or &#x60;false&#x60;. | 
 **CallSidToCoach** | **optional.String**| The SID of the participant who is being &#x60;coached&#x60;. The participant being coached is the only participant who can hear the participant who is &#x60;coaching&#x60;. | 
 **Coaching** | **optional.Bool**| Whether the participant is coaching another call. Can be: &#x60;true&#x60; or &#x60;false&#x60;. If not present, defaults to &#x60;false&#x60; unless &#x60;call_sid_to_coach&#x60; is defined. If &#x60;true&#x60;, &#x60;call_sid_to_coach&#x60; must be defined. | 
 **EndConferenceOnExit** | **optional.Bool**| Whether to end the conference when the participant leaves. Can be: &#x60;true&#x60; or &#x60;false&#x60; and defaults to &#x60;false&#x60;. | 
 **Hold** | **optional.Bool**| Whether the participant should be on hold. Can be: &#x60;true&#x60; or &#x60;false&#x60;. &#x60;true&#x60; puts the participant on hold, and &#x60;false&#x60; lets them rejoin the conference. | 
 **HoldMethod** | **optional.String**| The HTTP method we should use to call &#x60;hold_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60; and the default is &#x60;GET&#x60;. | 
 **HoldUrl** | **optional.String**| The URL we call using the &#x60;hold_method&#x60; for  music that plays when the participant is on hold. The URL may return an MP3 file, a WAV file, or a TwiML document that contains the &#x60;&lt;Play&gt;&#x60;, &#x60;&lt;Say&gt;&#x60; or &#x60;&lt;Redirect&gt;&#x60; commands. | 
 **Muted** | **optional.Bool**| Whether the participant should be muted. Can be &#x60;true&#x60; or &#x60;false&#x60;. &#x60;true&#x60; will mute the participant, and &#x60;false&#x60; will un-mute them. Anything value other than &#x60;true&#x60; or &#x60;false&#x60; is interpreted as &#x60;false&#x60;. | 
 **WaitMethod** | **optional.String**| The HTTP method we should use to call &#x60;wait_url&#x60;. Can be &#x60;GET&#x60; or &#x60;POST&#x60; and the default is &#x60;POST&#x60;. When using a static audio file, this should be &#x60;GET&#x60; so that we can cache the file. | 
 **WaitUrl** | **optional.String**| The URL we should call using the &#x60;wait_method&#x60; for the music to play while participants are waiting for the conference to start. The default value is the URL of our standard hold music. [Learn more about hold music](https://www.twilio.com/labs/twimlets/holdmusic). | 

### Return type

[**ApiV2010AccountConferenceParticipant**](api.v2010.account.conference.participant.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePayments

> ApiV2010AccountCallPayments UpdatePayments(ctx, AccountSid, CallSid, Sid, optional)



update an instance of payments with different phases of payment flows.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string**| The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that will update the resource. | 
**CallSid** | **string**| The SID of the call that will update the resource. This should be the same call sid that was used to create payments resource. | 
**Sid** | **string**| The SID of Payments session that needs to be updated. | 
 **optional** | ***UpdatePaymentsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdatePaymentsOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **Capture** | **optional.String**| The piece of payment information that you wish the caller to enter. Must be one of &#x60;payment-card-number&#x60;, &#x60;expiration-date&#x60;, &#x60;security-code&#x60;, &#x60;postal-code&#x60;, &#x60;bank-routing-number&#x60;, or &#x60;bank-account-number&#x60;. | 
 **IdempotencyKey** | **optional.String**| A unique token that will be used to ensure that multiple API calls with the same information do not result in multiple transactions. This should be a unique string value per API call and can be a randomly generated. | 
 **Status** | **optional.String**| Indicates whether the current payment session should be cancelled or completed. When &#x60;cancel&#x60; the payment session is cancelled. When &#x60;complete&#x60;, Twilio sends the payment information to the selected &lt;Pay&gt; connector for processing. | 
 **StatusCallback** | **optional.String**| Provide an absolute or relative URL to receive status updates regarding your Pay session. Read more about the [Update](https://www.twilio.com/docs/voice/api/payment-resource#statuscallback-update) and [Complete/Cancel](https://www.twilio.com/docs/voice/api/payment-resource#statuscallback-cancelcomplete) POST requests. | 

### Return type

[**ApiV2010AccountCallPayments**](api.v2010.account.call.payments.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateQueue

> ApiV2010AccountQueue UpdateQueue(ctx, AccountSid, Sid, optional)



Update the queue with the new parameters

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string**| The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Queue resource to update. | 
**Sid** | **string**| The Twilio-provided string that uniquely identifies the Queue resource to update | 
 **optional** | ***UpdateQueueOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateQueueOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **FriendlyName** | **optional.String**| A descriptive string that you created to describe this resource. It can be up to 64 characters long. | 
 **MaxSize** | **optional.Int32**| The maximum number of calls allowed to be in the queue. The default is 100. The maximum is 5000. | 

### Return type

[**ApiV2010AccountQueue**](api.v2010.account.queue.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateShortCode

> ApiV2010AccountShortCode UpdateShortCode(ctx, AccountSid, Sid, optional)



Update a short code with the following parameters

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string**| The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the ShortCode resource(s) to update. | 
**Sid** | **string**| The Twilio-provided string that uniquely identifies the ShortCode resource to update | 
 **optional** | ***UpdateShortCodeOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateShortCodeOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **ApiVersion** | **optional.String**| The API version to use to start a new TwiML session. Can be: &#x60;2010-04-01&#x60; or &#x60;2008-08-01&#x60;. | 
 **FriendlyName** | **optional.String**| A descriptive string that you created to describe this resource. It can be up to 64 characters long. By default, the &#x60;FriendlyName&#x60; is the short code. | 
 **SmsFallbackMethod** | **optional.String**| The HTTP method that we should use to call the &#x60;sms_fallback_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60;. | 
 **SmsFallbackUrl** | **optional.String**| The URL that we should call if an error occurs while retrieving or executing the TwiML from &#x60;sms_url&#x60;. | 
 **SmsMethod** | **optional.String**| The HTTP method we should use when calling the &#x60;sms_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60;. | 
 **SmsUrl** | **optional.String**| The URL we should call when receiving an incoming SMS message to this short code. | 

### Return type

[**ApiV2010AccountShortCode**](api.v2010.account.short_code.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSigningKey

> ApiV2010AccountSigningKey UpdateSigningKey(ctx, AccountSid, Sid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string**|  | 
**Sid** | **string**|  | 
 **optional** | ***UpdateSigningKeyOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateSigningKeyOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **FriendlyName** | **optional.String**|  | 

### Return type

[**ApiV2010AccountSigningKey**](api.v2010.account.signing_key.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSipCredential

> ApiV2010AccountSipSipCredentialListSipCredential UpdateSipCredential(ctx, AccountSid, CredentialListSid, Sid, optional)



Update a credential resource.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string**| The unique id of the Account that is responsible for this resource. | 
**CredentialListSid** | **string**| The unique id that identifies the credential list that includes this credential. | 
**Sid** | **string**| The unique id that identifies the resource to update. | 
 **optional** | ***UpdateSipCredentialOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateSipCredentialOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **Password** | **optional.String**| The password that the username will use when authenticating SIP requests. The password must be a minimum of 12 characters, contain at least 1 digit, and have mixed case. (eg &#x60;IWasAtSignal2018&#x60;) | 

### Return type

[**ApiV2010AccountSipSipCredentialListSipCredential**](api.v2010.account.sip.sip_credential_list.sip_credential.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSipCredentialList

> ApiV2010AccountSipSipCredentialList UpdateSipCredentialList(ctx, AccountSid, Sid, optional)



Update a Credential List

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string**| The unique id of the Account that is responsible for this resource. | 
**Sid** | **string**| The credential list Sid that uniquely identifies this resource | 
 **optional** | ***UpdateSipCredentialListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateSipCredentialListOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **FriendlyName** | **optional.String**| A human readable descriptive text for a CredentialList, up to 64 characters long. | 

### Return type

[**ApiV2010AccountSipSipCredentialList**](api.v2010.account.sip.sip_credential_list.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSipDomain

> ApiV2010AccountSipSipDomain UpdateSipDomain(ctx, AccountSid, Sid, optional)



Update the attributes of a domain

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string**| The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the SipDomain resource to update. | 
**Sid** | **string**| The Twilio-provided string that uniquely identifies the SipDomain resource to update. | 
 **optional** | ***UpdateSipDomainOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateSipDomainOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **ByocTrunkSid** | **optional.String**| The SID of the BYOC Trunk(Bring Your Own Carrier) resource that the Sip Domain will be associated with. | 
 **DomainName** | **optional.String**| The unique address you reserve on Twilio to which you route your SIP traffic. Domain names can contain letters, digits, and \\\&quot;-\\\&quot; and must end with &#x60;sip.twilio.com&#x60;. | 
 **EmergencyCallerSid** | **optional.String**| Whether an emergency caller sid is configured for the domain. If present, this phone number will be used as the callback for the emergency call. | 
 **EmergencyCallingEnabled** | **optional.Bool**| Whether emergency calling is enabled for the domain. If enabled, allows emergency calls on the domain from phone numbers with validated addresses. | 
 **FriendlyName** | **optional.String**| A descriptive string that you created to describe the resource. It can be up to 64 characters long. | 
 **Secure** | **optional.Bool**| Whether secure SIP is enabled for the domain. If enabled, TLS will be enforced and SRTP will be negotiated on all incoming calls to this sip domain. | 
 **SipRegistration** | **optional.Bool**| Whether to allow SIP Endpoints to register with the domain to receive calls. Can be &#x60;true&#x60; or &#x60;false&#x60;. &#x60;true&#x60; allows SIP Endpoints to register with the domain to receive calls, &#x60;false&#x60; does not. | 
 **VoiceFallbackMethod** | **optional.String**| The HTTP method we should use to call &#x60;voice_fallback_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60;. | 
 **VoiceFallbackUrl** | **optional.String**| The URL that we should call when an error occurs while retrieving or executing the TwiML requested by &#x60;voice_url&#x60;. | 
 **VoiceMethod** | **optional.String**| The HTTP method we should use to call &#x60;voice_url&#x60; | 
 **VoiceStatusCallbackMethod** | **optional.String**| The HTTP method we should use to call &#x60;voice_status_callback_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60;. | 
 **VoiceStatusCallbackUrl** | **optional.String**| The URL that we should call to pass status parameters (such as call ended) to your application. | 
 **VoiceUrl** | **optional.String**| The URL we should call when the domain receives a call. | 

### Return type

[**ApiV2010AccountSipSipDomain**](api.v2010.account.sip.sip_domain.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSipIpAccessControlList

> ApiV2010AccountSipSipIpAccessControlList UpdateSipIpAccessControlList(ctx, AccountSid, Sid, optional)



Rename an IpAccessControlList

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string**| The unique id of the [Account](https://www.twilio.com/docs/iam/api/account) responsible for this resource. | 
**Sid** | **string**| A 34 character string that uniquely identifies the resource to udpate. | 
 **optional** | ***UpdateSipIpAccessControlListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateSipIpAccessControlListOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **FriendlyName** | **optional.String**| A human readable descriptive text, up to 64 characters long. | 

### Return type

[**ApiV2010AccountSipSipIpAccessControlList**](api.v2010.account.sip.sip_ip_access_control_list.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSipIpAddress

> ApiV2010AccountSipSipIpAccessControlListSipIpAddress UpdateSipIpAddress(ctx, AccountSid, IpAccessControlListSid, Sid, optional)



Update an IpAddress resource.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string**| The unique id of the [Account](https://www.twilio.com/docs/iam/api/account) responsible for this resource. | 
**IpAccessControlListSid** | **string**| The IpAccessControlList Sid that identifies the IpAddress resources to update. | 
**Sid** | **string**| A 34 character string that identifies the IpAddress resource to update. | 
 **optional** | ***UpdateSipIpAddressOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateSipIpAddressOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **CidrPrefixLength** | **optional.Int32**| An integer representing the length of the CIDR prefix to use with this IP address when accepting traffic. By default the entire IP address is used. | 
 **FriendlyName** | **optional.String**| A human readable descriptive text for this resource, up to 64 characters long. | 
 **IpAddress** | **optional.String**| An IP address in dotted decimal notation from which you want to accept traffic. Any SIP requests from this IP address will be allowed by Twilio. IPv4 only supported today. | 

### Return type

[**ApiV2010AccountSipSipIpAccessControlListSipIpAddress**](api.v2010.account.sip.sip_ip_access_control_list.sip_ip_address.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUsageTrigger

> ApiV2010AccountUsageUsageTrigger UpdateUsageTrigger(ctx, AccountSid, Sid, optional)



Update an instance of a usage trigger

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AccountSid** | **string**| The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the UsageTrigger resources to update. | 
**Sid** | **string**| The Twilio-provided string that uniquely identifies the UsageTrigger resource to update. | 
 **optional** | ***UpdateUsageTriggerOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateUsageTriggerOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **CallbackMethod** | **optional.String**| The HTTP method we should use to call &#x60;callback_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60; and the default is &#x60;POST&#x60;. | 
 **CallbackUrl** | **optional.String**| The URL we should call using &#x60;callback_method&#x60; when the trigger fires. | 
 **FriendlyName** | **optional.String**| A descriptive string that you create to describe the resource. It can be up to 64 characters long. | 

### Return type

[**ApiV2010AccountUsageUsageTrigger**](api.v2010.account.usage.usage_trigger.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

