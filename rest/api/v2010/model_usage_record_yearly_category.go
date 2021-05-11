/*
 * Twilio - Api
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.0.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// UsageRecordYearlyCategory the model 'UsageRecordYearlyCategory'
type UsageRecordYearlyCategory string

// List of usage_record_yearly_category
const (
	USAGERECORDYEARLYCATEGORY_AGENT_CONFERENCE                                       UsageRecordYearlyCategory = "agent-conference"
	USAGERECORDYEARLYCATEGORY_ANSWERING_MACHINE_DETECTION                            UsageRecordYearlyCategory = "answering-machine-detection"
	USAGERECORDYEARLYCATEGORY_AUTHY_AUTHENTICATIONS                                  UsageRecordYearlyCategory = "authy-authentications"
	USAGERECORDYEARLYCATEGORY_AUTHY_CALLS_OUTBOUND                                   UsageRecordYearlyCategory = "authy-calls-outbound"
	USAGERECORDYEARLYCATEGORY_AUTHY_MONTHLY_FEES                                     UsageRecordYearlyCategory = "authy-monthly-fees"
	USAGERECORDYEARLYCATEGORY_AUTHY_PHONE_INTELLIGENCE                               UsageRecordYearlyCategory = "authy-phone-intelligence"
	USAGERECORDYEARLYCATEGORY_AUTHY_PHONE_VERIFICATIONS                              UsageRecordYearlyCategory = "authy-phone-verifications"
	USAGERECORDYEARLYCATEGORY_AUTHY_SMS_OUTBOUND                                     UsageRecordYearlyCategory = "authy-sms-outbound"
	USAGERECORDYEARLYCATEGORY_CALL_PROGESS_EVENTS                                    UsageRecordYearlyCategory = "call-progess-events"
	USAGERECORDYEARLYCATEGORY_CALLERIDLOOKUPS                                        UsageRecordYearlyCategory = "calleridlookups"
	USAGERECORDYEARLYCATEGORY_CALLS                                                  UsageRecordYearlyCategory = "calls"
	USAGERECORDYEARLYCATEGORY_CALLS_CLIENT                                           UsageRecordYearlyCategory = "calls-client"
	USAGERECORDYEARLYCATEGORY_CALLS_GLOBALCONFERENCE                                 UsageRecordYearlyCategory = "calls-globalconference"
	USAGERECORDYEARLYCATEGORY_CALLS_INBOUND                                          UsageRecordYearlyCategory = "calls-inbound"
	USAGERECORDYEARLYCATEGORY_CALLS_INBOUND_LOCAL                                    UsageRecordYearlyCategory = "calls-inbound-local"
	USAGERECORDYEARLYCATEGORY_CALLS_INBOUND_MOBILE                                   UsageRecordYearlyCategory = "calls-inbound-mobile"
	USAGERECORDYEARLYCATEGORY_CALLS_INBOUND_TOLLFREE                                 UsageRecordYearlyCategory = "calls-inbound-tollfree"
	USAGERECORDYEARLYCATEGORY_CALLS_OUTBOUND                                         UsageRecordYearlyCategory = "calls-outbound"
	USAGERECORDYEARLYCATEGORY_CALLS_PAY_VERB_TRANSACTIONS                            UsageRecordYearlyCategory = "calls-pay-verb-transactions"
	USAGERECORDYEARLYCATEGORY_CALLS_RECORDINGS                                       UsageRecordYearlyCategory = "calls-recordings"
	USAGERECORDYEARLYCATEGORY_CALLS_SIP                                              UsageRecordYearlyCategory = "calls-sip"
	USAGERECORDYEARLYCATEGORY_CALLS_SIP_INBOUND                                      UsageRecordYearlyCategory = "calls-sip-inbound"
	USAGERECORDYEARLYCATEGORY_CALLS_SIP_OUTBOUND                                     UsageRecordYearlyCategory = "calls-sip-outbound"
	USAGERECORDYEARLYCATEGORY_CARRIER_LOOKUPS                                        UsageRecordYearlyCategory = "carrier-lookups"
	USAGERECORDYEARLYCATEGORY_CONVERSATIONS                                          UsageRecordYearlyCategory = "conversations"
	USAGERECORDYEARLYCATEGORY_CONVERSATIONS_API_REQUESTS                             UsageRecordYearlyCategory = "conversations-api-requests"
	USAGERECORDYEARLYCATEGORY_CONVERSATIONS_CONVERSATION_EVENTS                      UsageRecordYearlyCategory = "conversations-conversation-events"
	USAGERECORDYEARLYCATEGORY_CONVERSATIONS_ENDPOINT_CONNECTIVITY                    UsageRecordYearlyCategory = "conversations-endpoint-connectivity"
	USAGERECORDYEARLYCATEGORY_CONVERSATIONS_EVENTS                                   UsageRecordYearlyCategory = "conversations-events"
	USAGERECORDYEARLYCATEGORY_CONVERSATIONS_PARTICIPANT_EVENTS                       UsageRecordYearlyCategory = "conversations-participant-events"
	USAGERECORDYEARLYCATEGORY_CONVERSATIONS_PARTICIPANTS                             UsageRecordYearlyCategory = "conversations-participants"
	USAGERECORDYEARLYCATEGORY_CPS                                                    UsageRecordYearlyCategory = "cps"
	USAGERECORDYEARLYCATEGORY_FRAUD_LOOKUPS                                          UsageRecordYearlyCategory = "fraud-lookups"
	USAGERECORDYEARLYCATEGORY_GROUP_ROOMS                                            UsageRecordYearlyCategory = "group-rooms"
	USAGERECORDYEARLYCATEGORY_GROUP_ROOMS_DATA_TRACK                                 UsageRecordYearlyCategory = "group-rooms-data-track"
	USAGERECORDYEARLYCATEGORY_GROUP_ROOMS_ENCRYPTED_MEDIA_RECORDED                   UsageRecordYearlyCategory = "group-rooms-encrypted-media-recorded"
	USAGERECORDYEARLYCATEGORY_GROUP_ROOMS_MEDIA_DOWNLOADED                           UsageRecordYearlyCategory = "group-rooms-media-downloaded"
	USAGERECORDYEARLYCATEGORY_GROUP_ROOMS_MEDIA_RECORDED                             UsageRecordYearlyCategory = "group-rooms-media-recorded"
	USAGERECORDYEARLYCATEGORY_GROUP_ROOMS_MEDIA_ROUTED                               UsageRecordYearlyCategory = "group-rooms-media-routed"
	USAGERECORDYEARLYCATEGORY_GROUP_ROOMS_MEDIA_STORED                               UsageRecordYearlyCategory = "group-rooms-media-stored"
	USAGERECORDYEARLYCATEGORY_GROUP_ROOMS_PARTICIPANT_MINUTES                        UsageRecordYearlyCategory = "group-rooms-participant-minutes"
	USAGERECORDYEARLYCATEGORY_GROUP_ROOMS_RECORDED_MINUTES                           UsageRecordYearlyCategory = "group-rooms-recorded-minutes"
	USAGERECORDYEARLYCATEGORY_IMP_V1_USAGE                                           UsageRecordYearlyCategory = "imp-v1-usage"
	USAGERECORDYEARLYCATEGORY_LOOKUPS                                                UsageRecordYearlyCategory = "lookups"
	USAGERECORDYEARLYCATEGORY_MARKETPLACE                                            UsageRecordYearlyCategory = "marketplace"
	USAGERECORDYEARLYCATEGORY_MARKETPLACE_ALGORITHMIA_NAMED_ENTITY_RECOGNITION       UsageRecordYearlyCategory = "marketplace-algorithmia-named-entity-recognition"
	USAGERECORDYEARLYCATEGORY_MARKETPLACE_CADENCE_TRANSCRIPTION                      UsageRecordYearlyCategory = "marketplace-cadence-transcription"
	USAGERECORDYEARLYCATEGORY_MARKETPLACE_CADENCE_TRANSLATION                        UsageRecordYearlyCategory = "marketplace-cadence-translation"
	USAGERECORDYEARLYCATEGORY_MARKETPLACE_CAPIO_SPEECH_TO_TEXT                       UsageRecordYearlyCategory = "marketplace-capio-speech-to-text"
	USAGERECORDYEARLYCATEGORY_MARKETPLACE_CONVRIZA_ABABA                             UsageRecordYearlyCategory = "marketplace-convriza-ababa"
	USAGERECORDYEARLYCATEGORY_MARKETPLACE_DEEPGRAM_PHRASE_DETECTOR                   UsageRecordYearlyCategory = "marketplace-deepgram-phrase-detector"
	USAGERECORDYEARLYCATEGORY_MARKETPLACE_DIGITAL_SEGMENT_BUSINESS_INFO              UsageRecordYearlyCategory = "marketplace-digital-segment-business-info"
	USAGERECORDYEARLYCATEGORY_MARKETPLACE_FACEBOOK_OFFLINE_CONVERSIONS               UsageRecordYearlyCategory = "marketplace-facebook-offline-conversions"
	USAGERECORDYEARLYCATEGORY_MARKETPLACE_GOOGLE_SPEECH_TO_TEXT                      UsageRecordYearlyCategory = "marketplace-google-speech-to-text"
	USAGERECORDYEARLYCATEGORY_MARKETPLACE_IBM_WATSON_MESSAGE_INSIGHTS                UsageRecordYearlyCategory = "marketplace-ibm-watson-message-insights"
	USAGERECORDYEARLYCATEGORY_MARKETPLACE_IBM_WATSON_MESSAGE_SENTIMENT               UsageRecordYearlyCategory = "marketplace-ibm-watson-message-sentiment"
	USAGERECORDYEARLYCATEGORY_MARKETPLACE_IBM_WATSON_RECORDING_ANALYSIS              UsageRecordYearlyCategory = "marketplace-ibm-watson-recording-analysis"
	USAGERECORDYEARLYCATEGORY_MARKETPLACE_IBM_WATSON_TONE_ANALYZER                   UsageRecordYearlyCategory = "marketplace-ibm-watson-tone-analyzer"
	USAGERECORDYEARLYCATEGORY_MARKETPLACE_ICEHOOK_SYSTEMS_SCOUT                      UsageRecordYearlyCategory = "marketplace-icehook-systems-scout"
	USAGERECORDYEARLYCATEGORY_MARKETPLACE_INFOGROUP_DATAAXLE_BIZINFO                 UsageRecordYearlyCategory = "marketplace-infogroup-dataaxle-bizinfo"
	USAGERECORDYEARLYCATEGORY_MARKETPLACE_KEEN_IO_CONTACT_CENTER_ANALYTICS           UsageRecordYearlyCategory = "marketplace-keen-io-contact-center-analytics"
	USAGERECORDYEARLYCATEGORY_MARKETPLACE_MARCHEX_CLEANCALL                          UsageRecordYearlyCategory = "marketplace-marchex-cleancall"
	USAGERECORDYEARLYCATEGORY_MARKETPLACE_MARCHEX_SENTIMENT_ANALYSIS_FOR_SMS         UsageRecordYearlyCategory = "marketplace-marchex-sentiment-analysis-for-sms"
	USAGERECORDYEARLYCATEGORY_MARKETPLACE_MARKETPLACE_NEXTCALLER_SOCIAL_ID           UsageRecordYearlyCategory = "marketplace-marketplace-nextcaller-social-id"
	USAGERECORDYEARLYCATEGORY_MARKETPLACE_MOBILE_COMMONS_OPT_OUT_CLASSIFIER          UsageRecordYearlyCategory = "marketplace-mobile-commons-opt-out-classifier"
	USAGERECORDYEARLYCATEGORY_MARKETPLACE_NEXIWAVE_VOICEMAIL_TO_TEXT                 UsageRecordYearlyCategory = "marketplace-nexiwave-voicemail-to-text"
	USAGERECORDYEARLYCATEGORY_MARKETPLACE_NEXTCALLER_ADVANCED_CALLER_IDENTIFICATION  UsageRecordYearlyCategory = "marketplace-nextcaller-advanced-caller-identification"
	USAGERECORDYEARLYCATEGORY_MARKETPLACE_NOMOROBO_SPAM_SCORE                        UsageRecordYearlyCategory = "marketplace-nomorobo-spam-score"
	USAGERECORDYEARLYCATEGORY_MARKETPLACE_PAYFONE_TCPA_COMPLIANCE                    UsageRecordYearlyCategory = "marketplace-payfone-tcpa-compliance"
	USAGERECORDYEARLYCATEGORY_MARKETPLACE_REMEETING_AUTOMATIC_SPEECH_RECOGNITION     UsageRecordYearlyCategory = "marketplace-remeeting-automatic-speech-recognition"
	USAGERECORDYEARLYCATEGORY_MARKETPLACE_TCPA_DEFENSE_SOLUTIONS_BLACKLIST_FEED      UsageRecordYearlyCategory = "marketplace-tcpa-defense-solutions-blacklist-feed"
	USAGERECORDYEARLYCATEGORY_MARKETPLACE_TELO_OPENCNAM                              UsageRecordYearlyCategory = "marketplace-telo-opencnam"
	USAGERECORDYEARLYCATEGORY_MARKETPLACE_TRUECNAM_TRUE_SPAM                         UsageRecordYearlyCategory = "marketplace-truecnam-true-spam"
	USAGERECORDYEARLYCATEGORY_MARKETPLACE_TWILIO_CALLER_NAME_LOOKUP_US               UsageRecordYearlyCategory = "marketplace-twilio-caller-name-lookup-us"
	USAGERECORDYEARLYCATEGORY_MARKETPLACE_TWILIO_CARRIER_INFORMATION_LOOKUP          UsageRecordYearlyCategory = "marketplace-twilio-carrier-information-lookup"
	USAGERECORDYEARLYCATEGORY_MARKETPLACE_VOICEBASE_PCI                              UsageRecordYearlyCategory = "marketplace-voicebase-pci"
	USAGERECORDYEARLYCATEGORY_MARKETPLACE_VOICEBASE_TRANSCRIPTION                    UsageRecordYearlyCategory = "marketplace-voicebase-transcription"
	USAGERECORDYEARLYCATEGORY_MARKETPLACE_VOICEBASE_TRANSCRIPTION_CUSTOM_VOCABULARY  UsageRecordYearlyCategory = "marketplace-voicebase-transcription-custom-vocabulary"
	USAGERECORDYEARLYCATEGORY_MARKETPLACE_WHITEPAGES_PRO_CALLER_IDENTIFICATION       UsageRecordYearlyCategory = "marketplace-whitepages-pro-caller-identification"
	USAGERECORDYEARLYCATEGORY_MARKETPLACE_WHITEPAGES_PRO_PHONE_INTELLIGENCE          UsageRecordYearlyCategory = "marketplace-whitepages-pro-phone-intelligence"
	USAGERECORDYEARLYCATEGORY_MARKETPLACE_WHITEPAGES_PRO_PHONE_REPUTATION            UsageRecordYearlyCategory = "marketplace-whitepages-pro-phone-reputation"
	USAGERECORDYEARLYCATEGORY_MARKETPLACE_WOLFARM_SPOKEN_RESULTS                     UsageRecordYearlyCategory = "marketplace-wolfarm-spoken-results"
	USAGERECORDYEARLYCATEGORY_MARKETPLACE_WOLFRAM_SHORT_ANSWER                       UsageRecordYearlyCategory = "marketplace-wolfram-short-answer"
	USAGERECORDYEARLYCATEGORY_MARKETPLACE_YTICA_CONTACT_CENTER_REPORTING_ANALYTICS   UsageRecordYearlyCategory = "marketplace-ytica-contact-center-reporting-analytics"
	USAGERECORDYEARLYCATEGORY_MEDIASTORAGE                                           UsageRecordYearlyCategory = "mediastorage"
	USAGERECORDYEARLYCATEGORY_MMS                                                    UsageRecordYearlyCategory = "mms"
	USAGERECORDYEARLYCATEGORY_MMS_INBOUND                                            UsageRecordYearlyCategory = "mms-inbound"
	USAGERECORDYEARLYCATEGORY_MMS_INBOUND_LONGCODE                                   UsageRecordYearlyCategory = "mms-inbound-longcode"
	USAGERECORDYEARLYCATEGORY_MMS_INBOUND_SHORTCODE                                  UsageRecordYearlyCategory = "mms-inbound-shortcode"
	USAGERECORDYEARLYCATEGORY_MMS_MESSAGES_CARRIERFEES                               UsageRecordYearlyCategory = "mms-messages-carrierfees"
	USAGERECORDYEARLYCATEGORY_MMS_OUTBOUND                                           UsageRecordYearlyCategory = "mms-outbound"
	USAGERECORDYEARLYCATEGORY_MMS_OUTBOUND_LONGCODE                                  UsageRecordYearlyCategory = "mms-outbound-longcode"
	USAGERECORDYEARLYCATEGORY_MMS_OUTBOUND_SHORTCODE                                 UsageRecordYearlyCategory = "mms-outbound-shortcode"
	USAGERECORDYEARLYCATEGORY_MONITOR_READS                                          UsageRecordYearlyCategory = "monitor-reads"
	USAGERECORDYEARLYCATEGORY_MONITOR_STORAGE                                        UsageRecordYearlyCategory = "monitor-storage"
	USAGERECORDYEARLYCATEGORY_MONITOR_WRITES                                         UsageRecordYearlyCategory = "monitor-writes"
	USAGERECORDYEARLYCATEGORY_NOTIFY                                                 UsageRecordYearlyCategory = "notify"
	USAGERECORDYEARLYCATEGORY_NOTIFY_ACTIONS_ATTEMPTS                                UsageRecordYearlyCategory = "notify-actions-attempts"
	USAGERECORDYEARLYCATEGORY_NOTIFY_CHANNELS                                        UsageRecordYearlyCategory = "notify-channels"
	USAGERECORDYEARLYCATEGORY_NUMBER_FORMAT_LOOKUPS                                  UsageRecordYearlyCategory = "number-format-lookups"
	USAGERECORDYEARLYCATEGORY_PCHAT                                                  UsageRecordYearlyCategory = "pchat"
	USAGERECORDYEARLYCATEGORY_PCHAT_USERS                                            UsageRecordYearlyCategory = "pchat-users"
	USAGERECORDYEARLYCATEGORY_PEER_TO_PEER_ROOMS_PARTICIPANT_MINUTES                 UsageRecordYearlyCategory = "peer-to-peer-rooms-participant-minutes"
	USAGERECORDYEARLYCATEGORY_PFAX                                                   UsageRecordYearlyCategory = "pfax"
	USAGERECORDYEARLYCATEGORY_PFAX_MINUTES                                           UsageRecordYearlyCategory = "pfax-minutes"
	USAGERECORDYEARLYCATEGORY_PFAX_MINUTES_INBOUND                                   UsageRecordYearlyCategory = "pfax-minutes-inbound"
	USAGERECORDYEARLYCATEGORY_PFAX_MINUTES_OUTBOUND                                  UsageRecordYearlyCategory = "pfax-minutes-outbound"
	USAGERECORDYEARLYCATEGORY_PFAX_PAGES                                             UsageRecordYearlyCategory = "pfax-pages"
	USAGERECORDYEARLYCATEGORY_PHONENUMBERS                                           UsageRecordYearlyCategory = "phonenumbers"
	USAGERECORDYEARLYCATEGORY_PHONENUMBERS_CPS                                       UsageRecordYearlyCategory = "phonenumbers-cps"
	USAGERECORDYEARLYCATEGORY_PHONENUMBERS_EMERGENCY                                 UsageRecordYearlyCategory = "phonenumbers-emergency"
	USAGERECORDYEARLYCATEGORY_PHONENUMBERS_LOCAL                                     UsageRecordYearlyCategory = "phonenumbers-local"
	USAGERECORDYEARLYCATEGORY_PHONENUMBERS_MOBILE                                    UsageRecordYearlyCategory = "phonenumbers-mobile"
	USAGERECORDYEARLYCATEGORY_PHONENUMBERS_SETUPS                                    UsageRecordYearlyCategory = "phonenumbers-setups"
	USAGERECORDYEARLYCATEGORY_PHONENUMBERS_TOLLFREE                                  UsageRecordYearlyCategory = "phonenumbers-tollfree"
	USAGERECORDYEARLYCATEGORY_PREMIUMSUPPORT                                         UsageRecordYearlyCategory = "premiumsupport"
	USAGERECORDYEARLYCATEGORY_PROXY                                                  UsageRecordYearlyCategory = "proxy"
	USAGERECORDYEARLYCATEGORY_PROXY_ACTIVE_SESSIONS                                  UsageRecordYearlyCategory = "proxy-active-sessions"
	USAGERECORDYEARLYCATEGORY_PSTNCONNECTIVITY                                       UsageRecordYearlyCategory = "pstnconnectivity"
	USAGERECORDYEARLYCATEGORY_PV                                                     UsageRecordYearlyCategory = "pv"
	USAGERECORDYEARLYCATEGORY_PV_COMPOSITION_MEDIA_DOWNLOADED                        UsageRecordYearlyCategory = "pv-composition-media-downloaded"
	USAGERECORDYEARLYCATEGORY_PV_COMPOSITION_MEDIA_ENCRYPTED                         UsageRecordYearlyCategory = "pv-composition-media-encrypted"
	USAGERECORDYEARLYCATEGORY_PV_COMPOSITION_MEDIA_STORED                            UsageRecordYearlyCategory = "pv-composition-media-stored"
	USAGERECORDYEARLYCATEGORY_PV_COMPOSITION_MINUTES                                 UsageRecordYearlyCategory = "pv-composition-minutes"
	USAGERECORDYEARLYCATEGORY_PV_RECORDING_COMPOSITIONS                              UsageRecordYearlyCategory = "pv-recording-compositions"
	USAGERECORDYEARLYCATEGORY_PV_ROOM_PARTICIPANTS                                   UsageRecordYearlyCategory = "pv-room-participants"
	USAGERECORDYEARLYCATEGORY_PV_ROOM_PARTICIPANTS_AU1                               UsageRecordYearlyCategory = "pv-room-participants-au1"
	USAGERECORDYEARLYCATEGORY_PV_ROOM_PARTICIPANTS_BR1                               UsageRecordYearlyCategory = "pv-room-participants-br1"
	USAGERECORDYEARLYCATEGORY_PV_ROOM_PARTICIPANTS_IE1                               UsageRecordYearlyCategory = "pv-room-participants-ie1"
	USAGERECORDYEARLYCATEGORY_PV_ROOM_PARTICIPANTS_JP1                               UsageRecordYearlyCategory = "pv-room-participants-jp1"
	USAGERECORDYEARLYCATEGORY_PV_ROOM_PARTICIPANTS_SG1                               UsageRecordYearlyCategory = "pv-room-participants-sg1"
	USAGERECORDYEARLYCATEGORY_PV_ROOM_PARTICIPANTS_US1                               UsageRecordYearlyCategory = "pv-room-participants-us1"
	USAGERECORDYEARLYCATEGORY_PV_ROOM_PARTICIPANTS_US2                               UsageRecordYearlyCategory = "pv-room-participants-us2"
	USAGERECORDYEARLYCATEGORY_PV_ROOMS                                               UsageRecordYearlyCategory = "pv-rooms"
	USAGERECORDYEARLYCATEGORY_PV_SIP_ENDPOINT_REGISTRATIONS                          UsageRecordYearlyCategory = "pv-sip-endpoint-registrations"
	USAGERECORDYEARLYCATEGORY_RECORDINGS                                             UsageRecordYearlyCategory = "recordings"
	USAGERECORDYEARLYCATEGORY_RECORDINGSTORAGE                                       UsageRecordYearlyCategory = "recordingstorage"
	USAGERECORDYEARLYCATEGORY_ROOMS_GROUP_BANDWIDTH                                  UsageRecordYearlyCategory = "rooms-group-bandwidth"
	USAGERECORDYEARLYCATEGORY_ROOMS_GROUP_MINUTES                                    UsageRecordYearlyCategory = "rooms-group-minutes"
	USAGERECORDYEARLYCATEGORY_ROOMS_PEER_TO_PEER_MINUTES                             UsageRecordYearlyCategory = "rooms-peer-to-peer-minutes"
	USAGERECORDYEARLYCATEGORY_SHORTCODES                                             UsageRecordYearlyCategory = "shortcodes"
	USAGERECORDYEARLYCATEGORY_SHORTCODES_CUSTOMEROWNED                               UsageRecordYearlyCategory = "shortcodes-customerowned"
	USAGERECORDYEARLYCATEGORY_SHORTCODES_MMS_ENABLEMENT                              UsageRecordYearlyCategory = "shortcodes-mms-enablement"
	USAGERECORDYEARLYCATEGORY_SHORTCODES_MPS                                         UsageRecordYearlyCategory = "shortcodes-mps"
	USAGERECORDYEARLYCATEGORY_SHORTCODES_RANDOM                                      UsageRecordYearlyCategory = "shortcodes-random"
	USAGERECORDYEARLYCATEGORY_SHORTCODES_UK                                          UsageRecordYearlyCategory = "shortcodes-uk"
	USAGERECORDYEARLYCATEGORY_SHORTCODES_VANITY                                      UsageRecordYearlyCategory = "shortcodes-vanity"
	USAGERECORDYEARLYCATEGORY_SMALL_GROUP_ROOMS                                      UsageRecordYearlyCategory = "small-group-rooms"
	USAGERECORDYEARLYCATEGORY_SMALL_GROUP_ROOMS_DATA_TRACK                           UsageRecordYearlyCategory = "small-group-rooms-data-track"
	USAGERECORDYEARLYCATEGORY_SMALL_GROUP_ROOMS_PARTICIPANT_MINUTES                  UsageRecordYearlyCategory = "small-group-rooms-participant-minutes"
	USAGERECORDYEARLYCATEGORY_SMS                                                    UsageRecordYearlyCategory = "sms"
	USAGERECORDYEARLYCATEGORY_SMS_INBOUND                                            UsageRecordYearlyCategory = "sms-inbound"
	USAGERECORDYEARLYCATEGORY_SMS_INBOUND_LONGCODE                                   UsageRecordYearlyCategory = "sms-inbound-longcode"
	USAGERECORDYEARLYCATEGORY_SMS_INBOUND_SHORTCODE                                  UsageRecordYearlyCategory = "sms-inbound-shortcode"
	USAGERECORDYEARLYCATEGORY_SMS_MESSAGES_CARRIERFEES                               UsageRecordYearlyCategory = "sms-messages-carrierfees"
	USAGERECORDYEARLYCATEGORY_SMS_MESSAGES_FEATURES                                  UsageRecordYearlyCategory = "sms-messages-features"
	USAGERECORDYEARLYCATEGORY_SMS_MESSAGES_FEATURES_SENDERID                         UsageRecordYearlyCategory = "sms-messages-features-senderid"
	USAGERECORDYEARLYCATEGORY_SMS_OUTBOUND                                           UsageRecordYearlyCategory = "sms-outbound"
	USAGERECORDYEARLYCATEGORY_SMS_OUTBOUND_CONTENT_INSPECTION                        UsageRecordYearlyCategory = "sms-outbound-content-inspection"
	USAGERECORDYEARLYCATEGORY_SMS_OUTBOUND_LONGCODE                                  UsageRecordYearlyCategory = "sms-outbound-longcode"
	USAGERECORDYEARLYCATEGORY_SMS_OUTBOUND_SHORTCODE                                 UsageRecordYearlyCategory = "sms-outbound-shortcode"
	USAGERECORDYEARLYCATEGORY_SPEECH_RECOGNITION                                     UsageRecordYearlyCategory = "speech-recognition"
	USAGERECORDYEARLYCATEGORY_STUDIO_ENGAGEMENTS                                     UsageRecordYearlyCategory = "studio-engagements"
	USAGERECORDYEARLYCATEGORY_SYNC                                                   UsageRecordYearlyCategory = "sync"
	USAGERECORDYEARLYCATEGORY_SYNC_ACTIONS                                           UsageRecordYearlyCategory = "sync-actions"
	USAGERECORDYEARLYCATEGORY_SYNC_ENDPOINT_HOURS                                    UsageRecordYearlyCategory = "sync-endpoint-hours"
	USAGERECORDYEARLYCATEGORY_SYNC_ENDPOINT_HOURS_ABOVE_DAILY_CAP                    UsageRecordYearlyCategory = "sync-endpoint-hours-above-daily-cap"
	USAGERECORDYEARLYCATEGORY_TASKROUTER_TASKS                                       UsageRecordYearlyCategory = "taskrouter-tasks"
	USAGERECORDYEARLYCATEGORY_TOTALPRICE                                             UsageRecordYearlyCategory = "totalprice"
	USAGERECORDYEARLYCATEGORY_TRANSCRIPTIONS                                         UsageRecordYearlyCategory = "transcriptions"
	USAGERECORDYEARLYCATEGORY_TRUNKING_CPS                                           UsageRecordYearlyCategory = "trunking-cps"
	USAGERECORDYEARLYCATEGORY_TRUNKING_EMERGENCY_CALLS                               UsageRecordYearlyCategory = "trunking-emergency-calls"
	USAGERECORDYEARLYCATEGORY_TRUNKING_ORIGINATION                                   UsageRecordYearlyCategory = "trunking-origination"
	USAGERECORDYEARLYCATEGORY_TRUNKING_ORIGINATION_LOCAL                             UsageRecordYearlyCategory = "trunking-origination-local"
	USAGERECORDYEARLYCATEGORY_TRUNKING_ORIGINATION_MOBILE                            UsageRecordYearlyCategory = "trunking-origination-mobile"
	USAGERECORDYEARLYCATEGORY_TRUNKING_ORIGINATION_TOLLFREE                          UsageRecordYearlyCategory = "trunking-origination-tollfree"
	USAGERECORDYEARLYCATEGORY_TRUNKING_RECORDINGS                                    UsageRecordYearlyCategory = "trunking-recordings"
	USAGERECORDYEARLYCATEGORY_TRUNKING_SECURE                                        UsageRecordYearlyCategory = "trunking-secure"
	USAGERECORDYEARLYCATEGORY_TRUNKING_TERMINATION                                   UsageRecordYearlyCategory = "trunking-termination"
	USAGERECORDYEARLYCATEGORY_TURNMEGABYTES                                          UsageRecordYearlyCategory = "turnmegabytes"
	USAGERECORDYEARLYCATEGORY_TURNMEGABYTES_AUSTRALIA                                UsageRecordYearlyCategory = "turnmegabytes-australia"
	USAGERECORDYEARLYCATEGORY_TURNMEGABYTES_BRASIL                                   UsageRecordYearlyCategory = "turnmegabytes-brasil"
	USAGERECORDYEARLYCATEGORY_TURNMEGABYTES_GERMANY                                  UsageRecordYearlyCategory = "turnmegabytes-germany"
	USAGERECORDYEARLYCATEGORY_TURNMEGABYTES_INDIA                                    UsageRecordYearlyCategory = "turnmegabytes-india"
	USAGERECORDYEARLYCATEGORY_TURNMEGABYTES_IRELAND                                  UsageRecordYearlyCategory = "turnmegabytes-ireland"
	USAGERECORDYEARLYCATEGORY_TURNMEGABYTES_JAPAN                                    UsageRecordYearlyCategory = "turnmegabytes-japan"
	USAGERECORDYEARLYCATEGORY_TURNMEGABYTES_SINGAPORE                                UsageRecordYearlyCategory = "turnmegabytes-singapore"
	USAGERECORDYEARLYCATEGORY_TURNMEGABYTES_USEAST                                   UsageRecordYearlyCategory = "turnmegabytes-useast"
	USAGERECORDYEARLYCATEGORY_TURNMEGABYTES_USWEST                                   UsageRecordYearlyCategory = "turnmegabytes-uswest"
	USAGERECORDYEARLYCATEGORY_TWILIO_INTERCONNECT                                    UsageRecordYearlyCategory = "twilio-interconnect"
	USAGERECORDYEARLYCATEGORY_VERIFY_PUSH                                            UsageRecordYearlyCategory = "verify-push"
	USAGERECORDYEARLYCATEGORY_VIDEO_RECORDINGS                                       UsageRecordYearlyCategory = "video-recordings"
	USAGERECORDYEARLYCATEGORY_VOICE_INSIGHTS                                         UsageRecordYearlyCategory = "voice-insights"
	USAGERECORDYEARLYCATEGORY_VOICE_INSIGHTS_CLIENT_INSIGHTS_ON_DEMAND_MINUTE        UsageRecordYearlyCategory = "voice-insights-client-insights-on-demand-minute"
	USAGERECORDYEARLYCATEGORY_VOICE_INSIGHTS_PTSN_INSIGHTS_ON_DEMAND_MINUTE          UsageRecordYearlyCategory = "voice-insights-ptsn-insights-on-demand-minute"
	USAGERECORDYEARLYCATEGORY_VOICE_INSIGHTS_SIP_INTERFACE_INSIGHTS_ON_DEMAND_MINUTE UsageRecordYearlyCategory = "voice-insights-sip-interface-insights-on-demand-minute"
	USAGERECORDYEARLYCATEGORY_VOICE_INSIGHTS_SIP_TRUNKING_INSIGHTS_ON_DEMAND_MINUTE  UsageRecordYearlyCategory = "voice-insights-sip-trunking-insights-on-demand-minute"
	USAGERECORDYEARLYCATEGORY_WIRELESS                                               UsageRecordYearlyCategory = "wireless"
	USAGERECORDYEARLYCATEGORY_WIRELESS_ORDERS                                        UsageRecordYearlyCategory = "wireless-orders"
	USAGERECORDYEARLYCATEGORY_WIRELESS_ORDERS_ARTWORK                                UsageRecordYearlyCategory = "wireless-orders-artwork"
	USAGERECORDYEARLYCATEGORY_WIRELESS_ORDERS_BULK                                   UsageRecordYearlyCategory = "wireless-orders-bulk"
	USAGERECORDYEARLYCATEGORY_WIRELESS_ORDERS_ESIM                                   UsageRecordYearlyCategory = "wireless-orders-esim"
	USAGERECORDYEARLYCATEGORY_WIRELESS_ORDERS_STARTER                                UsageRecordYearlyCategory = "wireless-orders-starter"
	USAGERECORDYEARLYCATEGORY_WIRELESS_USAGE                                         UsageRecordYearlyCategory = "wireless-usage"
	USAGERECORDYEARLYCATEGORY_WIRELESS_USAGE_COMMANDS                                UsageRecordYearlyCategory = "wireless-usage-commands"
	USAGERECORDYEARLYCATEGORY_WIRELESS_USAGE_COMMANDS_AFRICA                         UsageRecordYearlyCategory = "wireless-usage-commands-africa"
	USAGERECORDYEARLYCATEGORY_WIRELESS_USAGE_COMMANDS_ASIA                           UsageRecordYearlyCategory = "wireless-usage-commands-asia"
	USAGERECORDYEARLYCATEGORY_WIRELESS_USAGE_COMMANDS_CENTRALANDSOUTHAMERICA         UsageRecordYearlyCategory = "wireless-usage-commands-centralandsouthamerica"
	USAGERECORDYEARLYCATEGORY_WIRELESS_USAGE_COMMANDS_EUROPE                         UsageRecordYearlyCategory = "wireless-usage-commands-europe"
	USAGERECORDYEARLYCATEGORY_WIRELESS_USAGE_COMMANDS_HOME                           UsageRecordYearlyCategory = "wireless-usage-commands-home"
	USAGERECORDYEARLYCATEGORY_WIRELESS_USAGE_COMMANDS_NORTHAMERICA                   UsageRecordYearlyCategory = "wireless-usage-commands-northamerica"
	USAGERECORDYEARLYCATEGORY_WIRELESS_USAGE_COMMANDS_OCEANIA                        UsageRecordYearlyCategory = "wireless-usage-commands-oceania"
	USAGERECORDYEARLYCATEGORY_WIRELESS_USAGE_COMMANDS_ROAMING                        UsageRecordYearlyCategory = "wireless-usage-commands-roaming"
	USAGERECORDYEARLYCATEGORY_WIRELESS_USAGE_DATA                                    UsageRecordYearlyCategory = "wireless-usage-data"
	USAGERECORDYEARLYCATEGORY_WIRELESS_USAGE_DATA_AFRICA                             UsageRecordYearlyCategory = "wireless-usage-data-africa"
	USAGERECORDYEARLYCATEGORY_WIRELESS_USAGE_DATA_ASIA                               UsageRecordYearlyCategory = "wireless-usage-data-asia"
	USAGERECORDYEARLYCATEGORY_WIRELESS_USAGE_DATA_CENTRALANDSOUTHAMERICA             UsageRecordYearlyCategory = "wireless-usage-data-centralandsouthamerica"
	USAGERECORDYEARLYCATEGORY_WIRELESS_USAGE_DATA_CUSTOM_ADDITIONALMB                UsageRecordYearlyCategory = "wireless-usage-data-custom-additionalmb"
	USAGERECORDYEARLYCATEGORY_WIRELESS_USAGE_DATA_CUSTOM_FIRST5MB                    UsageRecordYearlyCategory = "wireless-usage-data-custom-first5mb"
	USAGERECORDYEARLYCATEGORY_WIRELESS_USAGE_DATA_DOMESTIC_ROAMING                   UsageRecordYearlyCategory = "wireless-usage-data-domestic-roaming"
	USAGERECORDYEARLYCATEGORY_WIRELESS_USAGE_DATA_EUROPE                             UsageRecordYearlyCategory = "wireless-usage-data-europe"
	USAGERECORDYEARLYCATEGORY_WIRELESS_USAGE_DATA_INDIVIDUAL_ADDITIONALGB            UsageRecordYearlyCategory = "wireless-usage-data-individual-additionalgb"
	USAGERECORDYEARLYCATEGORY_WIRELESS_USAGE_DATA_INDIVIDUAL_FIRSTGB                 UsageRecordYearlyCategory = "wireless-usage-data-individual-firstgb"
	USAGERECORDYEARLYCATEGORY_WIRELESS_USAGE_DATA_INTERNATIONAL_ROAMING_CANADA       UsageRecordYearlyCategory = "wireless-usage-data-international-roaming-canada"
	USAGERECORDYEARLYCATEGORY_WIRELESS_USAGE_DATA_INTERNATIONAL_ROAMING_INDIA        UsageRecordYearlyCategory = "wireless-usage-data-international-roaming-india"
	USAGERECORDYEARLYCATEGORY_WIRELESS_USAGE_DATA_INTERNATIONAL_ROAMING_MEXICO       UsageRecordYearlyCategory = "wireless-usage-data-international-roaming-mexico"
	USAGERECORDYEARLYCATEGORY_WIRELESS_USAGE_DATA_NORTHAMERICA                       UsageRecordYearlyCategory = "wireless-usage-data-northamerica"
	USAGERECORDYEARLYCATEGORY_WIRELESS_USAGE_DATA_OCEANIA                            UsageRecordYearlyCategory = "wireless-usage-data-oceania"
	USAGERECORDYEARLYCATEGORY_WIRELESS_USAGE_DATA_POOLED                             UsageRecordYearlyCategory = "wireless-usage-data-pooled"
	USAGERECORDYEARLYCATEGORY_WIRELESS_USAGE_DATA_POOLED_DOWNLINK                    UsageRecordYearlyCategory = "wireless-usage-data-pooled-downlink"
	USAGERECORDYEARLYCATEGORY_WIRELESS_USAGE_DATA_POOLED_UPLINK                      UsageRecordYearlyCategory = "wireless-usage-data-pooled-uplink"
	USAGERECORDYEARLYCATEGORY_WIRELESS_USAGE_MRC                                     UsageRecordYearlyCategory = "wireless-usage-mrc"
	USAGERECORDYEARLYCATEGORY_WIRELESS_USAGE_MRC_CUSTOM                              UsageRecordYearlyCategory = "wireless-usage-mrc-custom"
	USAGERECORDYEARLYCATEGORY_WIRELESS_USAGE_MRC_INDIVIDUAL                          UsageRecordYearlyCategory = "wireless-usage-mrc-individual"
	USAGERECORDYEARLYCATEGORY_WIRELESS_USAGE_MRC_POOLED                              UsageRecordYearlyCategory = "wireless-usage-mrc-pooled"
	USAGERECORDYEARLYCATEGORY_WIRELESS_USAGE_MRC_SUSPENDED                           UsageRecordYearlyCategory = "wireless-usage-mrc-suspended"
	USAGERECORDYEARLYCATEGORY_WIRELESS_USAGE_SMS                                     UsageRecordYearlyCategory = "wireless-usage-sms"
	USAGERECORDYEARLYCATEGORY_WIRELESS_USAGE_VOICE                                   UsageRecordYearlyCategory = "wireless-usage-voice"
)
