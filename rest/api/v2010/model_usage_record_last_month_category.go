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

// UsageRecordLastMonthCategory the model 'UsageRecordLastMonthCategory'
type UsageRecordLastMonthCategory string

// List of usage_record_last_month_category
const (
	USAGERECORDLASTMONTHCATEGORY_AGENT_CONFERENCE                                       UsageRecordLastMonthCategory = "agent-conference"
	USAGERECORDLASTMONTHCATEGORY_ANSWERING_MACHINE_DETECTION                            UsageRecordLastMonthCategory = "answering-machine-detection"
	USAGERECORDLASTMONTHCATEGORY_AUTHY_AUTHENTICATIONS                                  UsageRecordLastMonthCategory = "authy-authentications"
	USAGERECORDLASTMONTHCATEGORY_AUTHY_CALLS_OUTBOUND                                   UsageRecordLastMonthCategory = "authy-calls-outbound"
	USAGERECORDLASTMONTHCATEGORY_AUTHY_MONTHLY_FEES                                     UsageRecordLastMonthCategory = "authy-monthly-fees"
	USAGERECORDLASTMONTHCATEGORY_AUTHY_PHONE_INTELLIGENCE                               UsageRecordLastMonthCategory = "authy-phone-intelligence"
	USAGERECORDLASTMONTHCATEGORY_AUTHY_PHONE_VERIFICATIONS                              UsageRecordLastMonthCategory = "authy-phone-verifications"
	USAGERECORDLASTMONTHCATEGORY_AUTHY_SMS_OUTBOUND                                     UsageRecordLastMonthCategory = "authy-sms-outbound"
	USAGERECORDLASTMONTHCATEGORY_CALL_PROGESS_EVENTS                                    UsageRecordLastMonthCategory = "call-progess-events"
	USAGERECORDLASTMONTHCATEGORY_CALLERIDLOOKUPS                                        UsageRecordLastMonthCategory = "calleridlookups"
	USAGERECORDLASTMONTHCATEGORY_CALLS                                                  UsageRecordLastMonthCategory = "calls"
	USAGERECORDLASTMONTHCATEGORY_CALLS_CLIENT                                           UsageRecordLastMonthCategory = "calls-client"
	USAGERECORDLASTMONTHCATEGORY_CALLS_GLOBALCONFERENCE                                 UsageRecordLastMonthCategory = "calls-globalconference"
	USAGERECORDLASTMONTHCATEGORY_CALLS_INBOUND                                          UsageRecordLastMonthCategory = "calls-inbound"
	USAGERECORDLASTMONTHCATEGORY_CALLS_INBOUND_LOCAL                                    UsageRecordLastMonthCategory = "calls-inbound-local"
	USAGERECORDLASTMONTHCATEGORY_CALLS_INBOUND_MOBILE                                   UsageRecordLastMonthCategory = "calls-inbound-mobile"
	USAGERECORDLASTMONTHCATEGORY_CALLS_INBOUND_TOLLFREE                                 UsageRecordLastMonthCategory = "calls-inbound-tollfree"
	USAGERECORDLASTMONTHCATEGORY_CALLS_OUTBOUND                                         UsageRecordLastMonthCategory = "calls-outbound"
	USAGERECORDLASTMONTHCATEGORY_CALLS_PAY_VERB_TRANSACTIONS                            UsageRecordLastMonthCategory = "calls-pay-verb-transactions"
	USAGERECORDLASTMONTHCATEGORY_CALLS_RECORDINGS                                       UsageRecordLastMonthCategory = "calls-recordings"
	USAGERECORDLASTMONTHCATEGORY_CALLS_SIP                                              UsageRecordLastMonthCategory = "calls-sip"
	USAGERECORDLASTMONTHCATEGORY_CALLS_SIP_INBOUND                                      UsageRecordLastMonthCategory = "calls-sip-inbound"
	USAGERECORDLASTMONTHCATEGORY_CALLS_SIP_OUTBOUND                                     UsageRecordLastMonthCategory = "calls-sip-outbound"
	USAGERECORDLASTMONTHCATEGORY_CARRIER_LOOKUPS                                        UsageRecordLastMonthCategory = "carrier-lookups"
	USAGERECORDLASTMONTHCATEGORY_CONVERSATIONS                                          UsageRecordLastMonthCategory = "conversations"
	USAGERECORDLASTMONTHCATEGORY_CONVERSATIONS_API_REQUESTS                             UsageRecordLastMonthCategory = "conversations-api-requests"
	USAGERECORDLASTMONTHCATEGORY_CONVERSATIONS_CONVERSATION_EVENTS                      UsageRecordLastMonthCategory = "conversations-conversation-events"
	USAGERECORDLASTMONTHCATEGORY_CONVERSATIONS_ENDPOINT_CONNECTIVITY                    UsageRecordLastMonthCategory = "conversations-endpoint-connectivity"
	USAGERECORDLASTMONTHCATEGORY_CONVERSATIONS_EVENTS                                   UsageRecordLastMonthCategory = "conversations-events"
	USAGERECORDLASTMONTHCATEGORY_CONVERSATIONS_PARTICIPANT_EVENTS                       UsageRecordLastMonthCategory = "conversations-participant-events"
	USAGERECORDLASTMONTHCATEGORY_CONVERSATIONS_PARTICIPANTS                             UsageRecordLastMonthCategory = "conversations-participants"
	USAGERECORDLASTMONTHCATEGORY_CPS                                                    UsageRecordLastMonthCategory = "cps"
	USAGERECORDLASTMONTHCATEGORY_FRAUD_LOOKUPS                                          UsageRecordLastMonthCategory = "fraud-lookups"
	USAGERECORDLASTMONTHCATEGORY_GROUP_ROOMS                                            UsageRecordLastMonthCategory = "group-rooms"
	USAGERECORDLASTMONTHCATEGORY_GROUP_ROOMS_DATA_TRACK                                 UsageRecordLastMonthCategory = "group-rooms-data-track"
	USAGERECORDLASTMONTHCATEGORY_GROUP_ROOMS_ENCRYPTED_MEDIA_RECORDED                   UsageRecordLastMonthCategory = "group-rooms-encrypted-media-recorded"
	USAGERECORDLASTMONTHCATEGORY_GROUP_ROOMS_MEDIA_DOWNLOADED                           UsageRecordLastMonthCategory = "group-rooms-media-downloaded"
	USAGERECORDLASTMONTHCATEGORY_GROUP_ROOMS_MEDIA_RECORDED                             UsageRecordLastMonthCategory = "group-rooms-media-recorded"
	USAGERECORDLASTMONTHCATEGORY_GROUP_ROOMS_MEDIA_ROUTED                               UsageRecordLastMonthCategory = "group-rooms-media-routed"
	USAGERECORDLASTMONTHCATEGORY_GROUP_ROOMS_MEDIA_STORED                               UsageRecordLastMonthCategory = "group-rooms-media-stored"
	USAGERECORDLASTMONTHCATEGORY_GROUP_ROOMS_PARTICIPANT_MINUTES                        UsageRecordLastMonthCategory = "group-rooms-participant-minutes"
	USAGERECORDLASTMONTHCATEGORY_GROUP_ROOMS_RECORDED_MINUTES                           UsageRecordLastMonthCategory = "group-rooms-recorded-minutes"
	USAGERECORDLASTMONTHCATEGORY_IMP_V1_USAGE                                           UsageRecordLastMonthCategory = "imp-v1-usage"
	USAGERECORDLASTMONTHCATEGORY_LOOKUPS                                                UsageRecordLastMonthCategory = "lookups"
	USAGERECORDLASTMONTHCATEGORY_MARKETPLACE                                            UsageRecordLastMonthCategory = "marketplace"
	USAGERECORDLASTMONTHCATEGORY_MARKETPLACE_ALGORITHMIA_NAMED_ENTITY_RECOGNITION       UsageRecordLastMonthCategory = "marketplace-algorithmia-named-entity-recognition"
	USAGERECORDLASTMONTHCATEGORY_MARKETPLACE_CADENCE_TRANSCRIPTION                      UsageRecordLastMonthCategory = "marketplace-cadence-transcription"
	USAGERECORDLASTMONTHCATEGORY_MARKETPLACE_CADENCE_TRANSLATION                        UsageRecordLastMonthCategory = "marketplace-cadence-translation"
	USAGERECORDLASTMONTHCATEGORY_MARKETPLACE_CAPIO_SPEECH_TO_TEXT                       UsageRecordLastMonthCategory = "marketplace-capio-speech-to-text"
	USAGERECORDLASTMONTHCATEGORY_MARKETPLACE_CONVRIZA_ABABA                             UsageRecordLastMonthCategory = "marketplace-convriza-ababa"
	USAGERECORDLASTMONTHCATEGORY_MARKETPLACE_DEEPGRAM_PHRASE_DETECTOR                   UsageRecordLastMonthCategory = "marketplace-deepgram-phrase-detector"
	USAGERECORDLASTMONTHCATEGORY_MARKETPLACE_DIGITAL_SEGMENT_BUSINESS_INFO              UsageRecordLastMonthCategory = "marketplace-digital-segment-business-info"
	USAGERECORDLASTMONTHCATEGORY_MARKETPLACE_FACEBOOK_OFFLINE_CONVERSIONS               UsageRecordLastMonthCategory = "marketplace-facebook-offline-conversions"
	USAGERECORDLASTMONTHCATEGORY_MARKETPLACE_GOOGLE_SPEECH_TO_TEXT                      UsageRecordLastMonthCategory = "marketplace-google-speech-to-text"
	USAGERECORDLASTMONTHCATEGORY_MARKETPLACE_IBM_WATSON_MESSAGE_INSIGHTS                UsageRecordLastMonthCategory = "marketplace-ibm-watson-message-insights"
	USAGERECORDLASTMONTHCATEGORY_MARKETPLACE_IBM_WATSON_MESSAGE_SENTIMENT               UsageRecordLastMonthCategory = "marketplace-ibm-watson-message-sentiment"
	USAGERECORDLASTMONTHCATEGORY_MARKETPLACE_IBM_WATSON_RECORDING_ANALYSIS              UsageRecordLastMonthCategory = "marketplace-ibm-watson-recording-analysis"
	USAGERECORDLASTMONTHCATEGORY_MARKETPLACE_IBM_WATSON_TONE_ANALYZER                   UsageRecordLastMonthCategory = "marketplace-ibm-watson-tone-analyzer"
	USAGERECORDLASTMONTHCATEGORY_MARKETPLACE_ICEHOOK_SYSTEMS_SCOUT                      UsageRecordLastMonthCategory = "marketplace-icehook-systems-scout"
	USAGERECORDLASTMONTHCATEGORY_MARKETPLACE_INFOGROUP_DATAAXLE_BIZINFO                 UsageRecordLastMonthCategory = "marketplace-infogroup-dataaxle-bizinfo"
	USAGERECORDLASTMONTHCATEGORY_MARKETPLACE_KEEN_IO_CONTACT_CENTER_ANALYTICS           UsageRecordLastMonthCategory = "marketplace-keen-io-contact-center-analytics"
	USAGERECORDLASTMONTHCATEGORY_MARKETPLACE_MARCHEX_CLEANCALL                          UsageRecordLastMonthCategory = "marketplace-marchex-cleancall"
	USAGERECORDLASTMONTHCATEGORY_MARKETPLACE_MARCHEX_SENTIMENT_ANALYSIS_FOR_SMS         UsageRecordLastMonthCategory = "marketplace-marchex-sentiment-analysis-for-sms"
	USAGERECORDLASTMONTHCATEGORY_MARKETPLACE_MARKETPLACE_NEXTCALLER_SOCIAL_ID           UsageRecordLastMonthCategory = "marketplace-marketplace-nextcaller-social-id"
	USAGERECORDLASTMONTHCATEGORY_MARKETPLACE_MOBILE_COMMONS_OPT_OUT_CLASSIFIER          UsageRecordLastMonthCategory = "marketplace-mobile-commons-opt-out-classifier"
	USAGERECORDLASTMONTHCATEGORY_MARKETPLACE_NEXIWAVE_VOICEMAIL_TO_TEXT                 UsageRecordLastMonthCategory = "marketplace-nexiwave-voicemail-to-text"
	USAGERECORDLASTMONTHCATEGORY_MARKETPLACE_NEXTCALLER_ADVANCED_CALLER_IDENTIFICATION  UsageRecordLastMonthCategory = "marketplace-nextcaller-advanced-caller-identification"
	USAGERECORDLASTMONTHCATEGORY_MARKETPLACE_NOMOROBO_SPAM_SCORE                        UsageRecordLastMonthCategory = "marketplace-nomorobo-spam-score"
	USAGERECORDLASTMONTHCATEGORY_MARKETPLACE_PAYFONE_TCPA_COMPLIANCE                    UsageRecordLastMonthCategory = "marketplace-payfone-tcpa-compliance"
	USAGERECORDLASTMONTHCATEGORY_MARKETPLACE_REMEETING_AUTOMATIC_SPEECH_RECOGNITION     UsageRecordLastMonthCategory = "marketplace-remeeting-automatic-speech-recognition"
	USAGERECORDLASTMONTHCATEGORY_MARKETPLACE_TCPA_DEFENSE_SOLUTIONS_BLACKLIST_FEED      UsageRecordLastMonthCategory = "marketplace-tcpa-defense-solutions-blacklist-feed"
	USAGERECORDLASTMONTHCATEGORY_MARKETPLACE_TELO_OPENCNAM                              UsageRecordLastMonthCategory = "marketplace-telo-opencnam"
	USAGERECORDLASTMONTHCATEGORY_MARKETPLACE_TRUECNAM_TRUE_SPAM                         UsageRecordLastMonthCategory = "marketplace-truecnam-true-spam"
	USAGERECORDLASTMONTHCATEGORY_MARKETPLACE_TWILIO_CALLER_NAME_LOOKUP_US               UsageRecordLastMonthCategory = "marketplace-twilio-caller-name-lookup-us"
	USAGERECORDLASTMONTHCATEGORY_MARKETPLACE_TWILIO_CARRIER_INFORMATION_LOOKUP          UsageRecordLastMonthCategory = "marketplace-twilio-carrier-information-lookup"
	USAGERECORDLASTMONTHCATEGORY_MARKETPLACE_VOICEBASE_PCI                              UsageRecordLastMonthCategory = "marketplace-voicebase-pci"
	USAGERECORDLASTMONTHCATEGORY_MARKETPLACE_VOICEBASE_TRANSCRIPTION                    UsageRecordLastMonthCategory = "marketplace-voicebase-transcription"
	USAGERECORDLASTMONTHCATEGORY_MARKETPLACE_VOICEBASE_TRANSCRIPTION_CUSTOM_VOCABULARY  UsageRecordLastMonthCategory = "marketplace-voicebase-transcription-custom-vocabulary"
	USAGERECORDLASTMONTHCATEGORY_MARKETPLACE_WHITEPAGES_PRO_CALLER_IDENTIFICATION       UsageRecordLastMonthCategory = "marketplace-whitepages-pro-caller-identification"
	USAGERECORDLASTMONTHCATEGORY_MARKETPLACE_WHITEPAGES_PRO_PHONE_INTELLIGENCE          UsageRecordLastMonthCategory = "marketplace-whitepages-pro-phone-intelligence"
	USAGERECORDLASTMONTHCATEGORY_MARKETPLACE_WHITEPAGES_PRO_PHONE_REPUTATION            UsageRecordLastMonthCategory = "marketplace-whitepages-pro-phone-reputation"
	USAGERECORDLASTMONTHCATEGORY_MARKETPLACE_WOLFARM_SPOKEN_RESULTS                     UsageRecordLastMonthCategory = "marketplace-wolfarm-spoken-results"
	USAGERECORDLASTMONTHCATEGORY_MARKETPLACE_WOLFRAM_SHORT_ANSWER                       UsageRecordLastMonthCategory = "marketplace-wolfram-short-answer"
	USAGERECORDLASTMONTHCATEGORY_MARKETPLACE_YTICA_CONTACT_CENTER_REPORTING_ANALYTICS   UsageRecordLastMonthCategory = "marketplace-ytica-contact-center-reporting-analytics"
	USAGERECORDLASTMONTHCATEGORY_MEDIASTORAGE                                           UsageRecordLastMonthCategory = "mediastorage"
	USAGERECORDLASTMONTHCATEGORY_MMS                                                    UsageRecordLastMonthCategory = "mms"
	USAGERECORDLASTMONTHCATEGORY_MMS_INBOUND                                            UsageRecordLastMonthCategory = "mms-inbound"
	USAGERECORDLASTMONTHCATEGORY_MMS_INBOUND_LONGCODE                                   UsageRecordLastMonthCategory = "mms-inbound-longcode"
	USAGERECORDLASTMONTHCATEGORY_MMS_INBOUND_SHORTCODE                                  UsageRecordLastMonthCategory = "mms-inbound-shortcode"
	USAGERECORDLASTMONTHCATEGORY_MMS_MESSAGES_CARRIERFEES                               UsageRecordLastMonthCategory = "mms-messages-carrierfees"
	USAGERECORDLASTMONTHCATEGORY_MMS_OUTBOUND                                           UsageRecordLastMonthCategory = "mms-outbound"
	USAGERECORDLASTMONTHCATEGORY_MMS_OUTBOUND_LONGCODE                                  UsageRecordLastMonthCategory = "mms-outbound-longcode"
	USAGERECORDLASTMONTHCATEGORY_MMS_OUTBOUND_SHORTCODE                                 UsageRecordLastMonthCategory = "mms-outbound-shortcode"
	USAGERECORDLASTMONTHCATEGORY_MONITOR_READS                                          UsageRecordLastMonthCategory = "monitor-reads"
	USAGERECORDLASTMONTHCATEGORY_MONITOR_STORAGE                                        UsageRecordLastMonthCategory = "monitor-storage"
	USAGERECORDLASTMONTHCATEGORY_MONITOR_WRITES                                         UsageRecordLastMonthCategory = "monitor-writes"
	USAGERECORDLASTMONTHCATEGORY_NOTIFY                                                 UsageRecordLastMonthCategory = "notify"
	USAGERECORDLASTMONTHCATEGORY_NOTIFY_ACTIONS_ATTEMPTS                                UsageRecordLastMonthCategory = "notify-actions-attempts"
	USAGERECORDLASTMONTHCATEGORY_NOTIFY_CHANNELS                                        UsageRecordLastMonthCategory = "notify-channels"
	USAGERECORDLASTMONTHCATEGORY_NUMBER_FORMAT_LOOKUPS                                  UsageRecordLastMonthCategory = "number-format-lookups"
	USAGERECORDLASTMONTHCATEGORY_PCHAT                                                  UsageRecordLastMonthCategory = "pchat"
	USAGERECORDLASTMONTHCATEGORY_PCHAT_USERS                                            UsageRecordLastMonthCategory = "pchat-users"
	USAGERECORDLASTMONTHCATEGORY_PEER_TO_PEER_ROOMS_PARTICIPANT_MINUTES                 UsageRecordLastMonthCategory = "peer-to-peer-rooms-participant-minutes"
	USAGERECORDLASTMONTHCATEGORY_PFAX                                                   UsageRecordLastMonthCategory = "pfax"
	USAGERECORDLASTMONTHCATEGORY_PFAX_MINUTES                                           UsageRecordLastMonthCategory = "pfax-minutes"
	USAGERECORDLASTMONTHCATEGORY_PFAX_MINUTES_INBOUND                                   UsageRecordLastMonthCategory = "pfax-minutes-inbound"
	USAGERECORDLASTMONTHCATEGORY_PFAX_MINUTES_OUTBOUND                                  UsageRecordLastMonthCategory = "pfax-minutes-outbound"
	USAGERECORDLASTMONTHCATEGORY_PFAX_PAGES                                             UsageRecordLastMonthCategory = "pfax-pages"
	USAGERECORDLASTMONTHCATEGORY_PHONENUMBERS                                           UsageRecordLastMonthCategory = "phonenumbers"
	USAGERECORDLASTMONTHCATEGORY_PHONENUMBERS_CPS                                       UsageRecordLastMonthCategory = "phonenumbers-cps"
	USAGERECORDLASTMONTHCATEGORY_PHONENUMBERS_EMERGENCY                                 UsageRecordLastMonthCategory = "phonenumbers-emergency"
	USAGERECORDLASTMONTHCATEGORY_PHONENUMBERS_LOCAL                                     UsageRecordLastMonthCategory = "phonenumbers-local"
	USAGERECORDLASTMONTHCATEGORY_PHONENUMBERS_MOBILE                                    UsageRecordLastMonthCategory = "phonenumbers-mobile"
	USAGERECORDLASTMONTHCATEGORY_PHONENUMBERS_SETUPS                                    UsageRecordLastMonthCategory = "phonenumbers-setups"
	USAGERECORDLASTMONTHCATEGORY_PHONENUMBERS_TOLLFREE                                  UsageRecordLastMonthCategory = "phonenumbers-tollfree"
	USAGERECORDLASTMONTHCATEGORY_PREMIUMSUPPORT                                         UsageRecordLastMonthCategory = "premiumsupport"
	USAGERECORDLASTMONTHCATEGORY_PROXY                                                  UsageRecordLastMonthCategory = "proxy"
	USAGERECORDLASTMONTHCATEGORY_PROXY_ACTIVE_SESSIONS                                  UsageRecordLastMonthCategory = "proxy-active-sessions"
	USAGERECORDLASTMONTHCATEGORY_PSTNCONNECTIVITY                                       UsageRecordLastMonthCategory = "pstnconnectivity"
	USAGERECORDLASTMONTHCATEGORY_PV                                                     UsageRecordLastMonthCategory = "pv"
	USAGERECORDLASTMONTHCATEGORY_PV_COMPOSITION_MEDIA_DOWNLOADED                        UsageRecordLastMonthCategory = "pv-composition-media-downloaded"
	USAGERECORDLASTMONTHCATEGORY_PV_COMPOSITION_MEDIA_ENCRYPTED                         UsageRecordLastMonthCategory = "pv-composition-media-encrypted"
	USAGERECORDLASTMONTHCATEGORY_PV_COMPOSITION_MEDIA_STORED                            UsageRecordLastMonthCategory = "pv-composition-media-stored"
	USAGERECORDLASTMONTHCATEGORY_PV_COMPOSITION_MINUTES                                 UsageRecordLastMonthCategory = "pv-composition-minutes"
	USAGERECORDLASTMONTHCATEGORY_PV_RECORDING_COMPOSITIONS                              UsageRecordLastMonthCategory = "pv-recording-compositions"
	USAGERECORDLASTMONTHCATEGORY_PV_ROOM_PARTICIPANTS                                   UsageRecordLastMonthCategory = "pv-room-participants"
	USAGERECORDLASTMONTHCATEGORY_PV_ROOM_PARTICIPANTS_AU1                               UsageRecordLastMonthCategory = "pv-room-participants-au1"
	USAGERECORDLASTMONTHCATEGORY_PV_ROOM_PARTICIPANTS_BR1                               UsageRecordLastMonthCategory = "pv-room-participants-br1"
	USAGERECORDLASTMONTHCATEGORY_PV_ROOM_PARTICIPANTS_IE1                               UsageRecordLastMonthCategory = "pv-room-participants-ie1"
	USAGERECORDLASTMONTHCATEGORY_PV_ROOM_PARTICIPANTS_JP1                               UsageRecordLastMonthCategory = "pv-room-participants-jp1"
	USAGERECORDLASTMONTHCATEGORY_PV_ROOM_PARTICIPANTS_SG1                               UsageRecordLastMonthCategory = "pv-room-participants-sg1"
	USAGERECORDLASTMONTHCATEGORY_PV_ROOM_PARTICIPANTS_US1                               UsageRecordLastMonthCategory = "pv-room-participants-us1"
	USAGERECORDLASTMONTHCATEGORY_PV_ROOM_PARTICIPANTS_US2                               UsageRecordLastMonthCategory = "pv-room-participants-us2"
	USAGERECORDLASTMONTHCATEGORY_PV_ROOMS                                               UsageRecordLastMonthCategory = "pv-rooms"
	USAGERECORDLASTMONTHCATEGORY_PV_SIP_ENDPOINT_REGISTRATIONS                          UsageRecordLastMonthCategory = "pv-sip-endpoint-registrations"
	USAGERECORDLASTMONTHCATEGORY_RECORDINGS                                             UsageRecordLastMonthCategory = "recordings"
	USAGERECORDLASTMONTHCATEGORY_RECORDINGSTORAGE                                       UsageRecordLastMonthCategory = "recordingstorage"
	USAGERECORDLASTMONTHCATEGORY_ROOMS_GROUP_BANDWIDTH                                  UsageRecordLastMonthCategory = "rooms-group-bandwidth"
	USAGERECORDLASTMONTHCATEGORY_ROOMS_GROUP_MINUTES                                    UsageRecordLastMonthCategory = "rooms-group-minutes"
	USAGERECORDLASTMONTHCATEGORY_ROOMS_PEER_TO_PEER_MINUTES                             UsageRecordLastMonthCategory = "rooms-peer-to-peer-minutes"
	USAGERECORDLASTMONTHCATEGORY_SHORTCODES                                             UsageRecordLastMonthCategory = "shortcodes"
	USAGERECORDLASTMONTHCATEGORY_SHORTCODES_CUSTOMEROWNED                               UsageRecordLastMonthCategory = "shortcodes-customerowned"
	USAGERECORDLASTMONTHCATEGORY_SHORTCODES_MMS_ENABLEMENT                              UsageRecordLastMonthCategory = "shortcodes-mms-enablement"
	USAGERECORDLASTMONTHCATEGORY_SHORTCODES_MPS                                         UsageRecordLastMonthCategory = "shortcodes-mps"
	USAGERECORDLASTMONTHCATEGORY_SHORTCODES_RANDOM                                      UsageRecordLastMonthCategory = "shortcodes-random"
	USAGERECORDLASTMONTHCATEGORY_SHORTCODES_UK                                          UsageRecordLastMonthCategory = "shortcodes-uk"
	USAGERECORDLASTMONTHCATEGORY_SHORTCODES_VANITY                                      UsageRecordLastMonthCategory = "shortcodes-vanity"
	USAGERECORDLASTMONTHCATEGORY_SMALL_GROUP_ROOMS                                      UsageRecordLastMonthCategory = "small-group-rooms"
	USAGERECORDLASTMONTHCATEGORY_SMALL_GROUP_ROOMS_DATA_TRACK                           UsageRecordLastMonthCategory = "small-group-rooms-data-track"
	USAGERECORDLASTMONTHCATEGORY_SMALL_GROUP_ROOMS_PARTICIPANT_MINUTES                  UsageRecordLastMonthCategory = "small-group-rooms-participant-minutes"
	USAGERECORDLASTMONTHCATEGORY_SMS                                                    UsageRecordLastMonthCategory = "sms"
	USAGERECORDLASTMONTHCATEGORY_SMS_INBOUND                                            UsageRecordLastMonthCategory = "sms-inbound"
	USAGERECORDLASTMONTHCATEGORY_SMS_INBOUND_LONGCODE                                   UsageRecordLastMonthCategory = "sms-inbound-longcode"
	USAGERECORDLASTMONTHCATEGORY_SMS_INBOUND_SHORTCODE                                  UsageRecordLastMonthCategory = "sms-inbound-shortcode"
	USAGERECORDLASTMONTHCATEGORY_SMS_MESSAGES_CARRIERFEES                               UsageRecordLastMonthCategory = "sms-messages-carrierfees"
	USAGERECORDLASTMONTHCATEGORY_SMS_MESSAGES_FEATURES                                  UsageRecordLastMonthCategory = "sms-messages-features"
	USAGERECORDLASTMONTHCATEGORY_SMS_MESSAGES_FEATURES_SENDERID                         UsageRecordLastMonthCategory = "sms-messages-features-senderid"
	USAGERECORDLASTMONTHCATEGORY_SMS_OUTBOUND                                           UsageRecordLastMonthCategory = "sms-outbound"
	USAGERECORDLASTMONTHCATEGORY_SMS_OUTBOUND_CONTENT_INSPECTION                        UsageRecordLastMonthCategory = "sms-outbound-content-inspection"
	USAGERECORDLASTMONTHCATEGORY_SMS_OUTBOUND_LONGCODE                                  UsageRecordLastMonthCategory = "sms-outbound-longcode"
	USAGERECORDLASTMONTHCATEGORY_SMS_OUTBOUND_SHORTCODE                                 UsageRecordLastMonthCategory = "sms-outbound-shortcode"
	USAGERECORDLASTMONTHCATEGORY_SPEECH_RECOGNITION                                     UsageRecordLastMonthCategory = "speech-recognition"
	USAGERECORDLASTMONTHCATEGORY_STUDIO_ENGAGEMENTS                                     UsageRecordLastMonthCategory = "studio-engagements"
	USAGERECORDLASTMONTHCATEGORY_SYNC                                                   UsageRecordLastMonthCategory = "sync"
	USAGERECORDLASTMONTHCATEGORY_SYNC_ACTIONS                                           UsageRecordLastMonthCategory = "sync-actions"
	USAGERECORDLASTMONTHCATEGORY_SYNC_ENDPOINT_HOURS                                    UsageRecordLastMonthCategory = "sync-endpoint-hours"
	USAGERECORDLASTMONTHCATEGORY_SYNC_ENDPOINT_HOURS_ABOVE_DAILY_CAP                    UsageRecordLastMonthCategory = "sync-endpoint-hours-above-daily-cap"
	USAGERECORDLASTMONTHCATEGORY_TASKROUTER_TASKS                                       UsageRecordLastMonthCategory = "taskrouter-tasks"
	USAGERECORDLASTMONTHCATEGORY_TOTALPRICE                                             UsageRecordLastMonthCategory = "totalprice"
	USAGERECORDLASTMONTHCATEGORY_TRANSCRIPTIONS                                         UsageRecordLastMonthCategory = "transcriptions"
	USAGERECORDLASTMONTHCATEGORY_TRUNKING_CPS                                           UsageRecordLastMonthCategory = "trunking-cps"
	USAGERECORDLASTMONTHCATEGORY_TRUNKING_EMERGENCY_CALLS                               UsageRecordLastMonthCategory = "trunking-emergency-calls"
	USAGERECORDLASTMONTHCATEGORY_TRUNKING_ORIGINATION                                   UsageRecordLastMonthCategory = "trunking-origination"
	USAGERECORDLASTMONTHCATEGORY_TRUNKING_ORIGINATION_LOCAL                             UsageRecordLastMonthCategory = "trunking-origination-local"
	USAGERECORDLASTMONTHCATEGORY_TRUNKING_ORIGINATION_MOBILE                            UsageRecordLastMonthCategory = "trunking-origination-mobile"
	USAGERECORDLASTMONTHCATEGORY_TRUNKING_ORIGINATION_TOLLFREE                          UsageRecordLastMonthCategory = "trunking-origination-tollfree"
	USAGERECORDLASTMONTHCATEGORY_TRUNKING_RECORDINGS                                    UsageRecordLastMonthCategory = "trunking-recordings"
	USAGERECORDLASTMONTHCATEGORY_TRUNKING_SECURE                                        UsageRecordLastMonthCategory = "trunking-secure"
	USAGERECORDLASTMONTHCATEGORY_TRUNKING_TERMINATION                                   UsageRecordLastMonthCategory = "trunking-termination"
	USAGERECORDLASTMONTHCATEGORY_TURNMEGABYTES                                          UsageRecordLastMonthCategory = "turnmegabytes"
	USAGERECORDLASTMONTHCATEGORY_TURNMEGABYTES_AUSTRALIA                                UsageRecordLastMonthCategory = "turnmegabytes-australia"
	USAGERECORDLASTMONTHCATEGORY_TURNMEGABYTES_BRASIL                                   UsageRecordLastMonthCategory = "turnmegabytes-brasil"
	USAGERECORDLASTMONTHCATEGORY_TURNMEGABYTES_GERMANY                                  UsageRecordLastMonthCategory = "turnmegabytes-germany"
	USAGERECORDLASTMONTHCATEGORY_TURNMEGABYTES_INDIA                                    UsageRecordLastMonthCategory = "turnmegabytes-india"
	USAGERECORDLASTMONTHCATEGORY_TURNMEGABYTES_IRELAND                                  UsageRecordLastMonthCategory = "turnmegabytes-ireland"
	USAGERECORDLASTMONTHCATEGORY_TURNMEGABYTES_JAPAN                                    UsageRecordLastMonthCategory = "turnmegabytes-japan"
	USAGERECORDLASTMONTHCATEGORY_TURNMEGABYTES_SINGAPORE                                UsageRecordLastMonthCategory = "turnmegabytes-singapore"
	USAGERECORDLASTMONTHCATEGORY_TURNMEGABYTES_USEAST                                   UsageRecordLastMonthCategory = "turnmegabytes-useast"
	USAGERECORDLASTMONTHCATEGORY_TURNMEGABYTES_USWEST                                   UsageRecordLastMonthCategory = "turnmegabytes-uswest"
	USAGERECORDLASTMONTHCATEGORY_TWILIO_INTERCONNECT                                    UsageRecordLastMonthCategory = "twilio-interconnect"
	USAGERECORDLASTMONTHCATEGORY_VERIFY_PUSH                                            UsageRecordLastMonthCategory = "verify-push"
	USAGERECORDLASTMONTHCATEGORY_VIDEO_RECORDINGS                                       UsageRecordLastMonthCategory = "video-recordings"
	USAGERECORDLASTMONTHCATEGORY_VOICE_INSIGHTS                                         UsageRecordLastMonthCategory = "voice-insights"
	USAGERECORDLASTMONTHCATEGORY_VOICE_INSIGHTS_CLIENT_INSIGHTS_ON_DEMAND_MINUTE        UsageRecordLastMonthCategory = "voice-insights-client-insights-on-demand-minute"
	USAGERECORDLASTMONTHCATEGORY_VOICE_INSIGHTS_PTSN_INSIGHTS_ON_DEMAND_MINUTE          UsageRecordLastMonthCategory = "voice-insights-ptsn-insights-on-demand-minute"
	USAGERECORDLASTMONTHCATEGORY_VOICE_INSIGHTS_SIP_INTERFACE_INSIGHTS_ON_DEMAND_MINUTE UsageRecordLastMonthCategory = "voice-insights-sip-interface-insights-on-demand-minute"
	USAGERECORDLASTMONTHCATEGORY_VOICE_INSIGHTS_SIP_TRUNKING_INSIGHTS_ON_DEMAND_MINUTE  UsageRecordLastMonthCategory = "voice-insights-sip-trunking-insights-on-demand-minute"
	USAGERECORDLASTMONTHCATEGORY_WIRELESS                                               UsageRecordLastMonthCategory = "wireless"
	USAGERECORDLASTMONTHCATEGORY_WIRELESS_ORDERS                                        UsageRecordLastMonthCategory = "wireless-orders"
	USAGERECORDLASTMONTHCATEGORY_WIRELESS_ORDERS_ARTWORK                                UsageRecordLastMonthCategory = "wireless-orders-artwork"
	USAGERECORDLASTMONTHCATEGORY_WIRELESS_ORDERS_BULK                                   UsageRecordLastMonthCategory = "wireless-orders-bulk"
	USAGERECORDLASTMONTHCATEGORY_WIRELESS_ORDERS_ESIM                                   UsageRecordLastMonthCategory = "wireless-orders-esim"
	USAGERECORDLASTMONTHCATEGORY_WIRELESS_ORDERS_STARTER                                UsageRecordLastMonthCategory = "wireless-orders-starter"
	USAGERECORDLASTMONTHCATEGORY_WIRELESS_USAGE                                         UsageRecordLastMonthCategory = "wireless-usage"
	USAGERECORDLASTMONTHCATEGORY_WIRELESS_USAGE_COMMANDS                                UsageRecordLastMonthCategory = "wireless-usage-commands"
	USAGERECORDLASTMONTHCATEGORY_WIRELESS_USAGE_COMMANDS_AFRICA                         UsageRecordLastMonthCategory = "wireless-usage-commands-africa"
	USAGERECORDLASTMONTHCATEGORY_WIRELESS_USAGE_COMMANDS_ASIA                           UsageRecordLastMonthCategory = "wireless-usage-commands-asia"
	USAGERECORDLASTMONTHCATEGORY_WIRELESS_USAGE_COMMANDS_CENTRALANDSOUTHAMERICA         UsageRecordLastMonthCategory = "wireless-usage-commands-centralandsouthamerica"
	USAGERECORDLASTMONTHCATEGORY_WIRELESS_USAGE_COMMANDS_EUROPE                         UsageRecordLastMonthCategory = "wireless-usage-commands-europe"
	USAGERECORDLASTMONTHCATEGORY_WIRELESS_USAGE_COMMANDS_HOME                           UsageRecordLastMonthCategory = "wireless-usage-commands-home"
	USAGERECORDLASTMONTHCATEGORY_WIRELESS_USAGE_COMMANDS_NORTHAMERICA                   UsageRecordLastMonthCategory = "wireless-usage-commands-northamerica"
	USAGERECORDLASTMONTHCATEGORY_WIRELESS_USAGE_COMMANDS_OCEANIA                        UsageRecordLastMonthCategory = "wireless-usage-commands-oceania"
	USAGERECORDLASTMONTHCATEGORY_WIRELESS_USAGE_COMMANDS_ROAMING                        UsageRecordLastMonthCategory = "wireless-usage-commands-roaming"
	USAGERECORDLASTMONTHCATEGORY_WIRELESS_USAGE_DATA                                    UsageRecordLastMonthCategory = "wireless-usage-data"
	USAGERECORDLASTMONTHCATEGORY_WIRELESS_USAGE_DATA_AFRICA                             UsageRecordLastMonthCategory = "wireless-usage-data-africa"
	USAGERECORDLASTMONTHCATEGORY_WIRELESS_USAGE_DATA_ASIA                               UsageRecordLastMonthCategory = "wireless-usage-data-asia"
	USAGERECORDLASTMONTHCATEGORY_WIRELESS_USAGE_DATA_CENTRALANDSOUTHAMERICA             UsageRecordLastMonthCategory = "wireless-usage-data-centralandsouthamerica"
	USAGERECORDLASTMONTHCATEGORY_WIRELESS_USAGE_DATA_CUSTOM_ADDITIONALMB                UsageRecordLastMonthCategory = "wireless-usage-data-custom-additionalmb"
	USAGERECORDLASTMONTHCATEGORY_WIRELESS_USAGE_DATA_CUSTOM_FIRST5MB                    UsageRecordLastMonthCategory = "wireless-usage-data-custom-first5mb"
	USAGERECORDLASTMONTHCATEGORY_WIRELESS_USAGE_DATA_DOMESTIC_ROAMING                   UsageRecordLastMonthCategory = "wireless-usage-data-domestic-roaming"
	USAGERECORDLASTMONTHCATEGORY_WIRELESS_USAGE_DATA_EUROPE                             UsageRecordLastMonthCategory = "wireless-usage-data-europe"
	USAGERECORDLASTMONTHCATEGORY_WIRELESS_USAGE_DATA_INDIVIDUAL_ADDITIONALGB            UsageRecordLastMonthCategory = "wireless-usage-data-individual-additionalgb"
	USAGERECORDLASTMONTHCATEGORY_WIRELESS_USAGE_DATA_INDIVIDUAL_FIRSTGB                 UsageRecordLastMonthCategory = "wireless-usage-data-individual-firstgb"
	USAGERECORDLASTMONTHCATEGORY_WIRELESS_USAGE_DATA_INTERNATIONAL_ROAMING_CANADA       UsageRecordLastMonthCategory = "wireless-usage-data-international-roaming-canada"
	USAGERECORDLASTMONTHCATEGORY_WIRELESS_USAGE_DATA_INTERNATIONAL_ROAMING_INDIA        UsageRecordLastMonthCategory = "wireless-usage-data-international-roaming-india"
	USAGERECORDLASTMONTHCATEGORY_WIRELESS_USAGE_DATA_INTERNATIONAL_ROAMING_MEXICO       UsageRecordLastMonthCategory = "wireless-usage-data-international-roaming-mexico"
	USAGERECORDLASTMONTHCATEGORY_WIRELESS_USAGE_DATA_NORTHAMERICA                       UsageRecordLastMonthCategory = "wireless-usage-data-northamerica"
	USAGERECORDLASTMONTHCATEGORY_WIRELESS_USAGE_DATA_OCEANIA                            UsageRecordLastMonthCategory = "wireless-usage-data-oceania"
	USAGERECORDLASTMONTHCATEGORY_WIRELESS_USAGE_DATA_POOLED                             UsageRecordLastMonthCategory = "wireless-usage-data-pooled"
	USAGERECORDLASTMONTHCATEGORY_WIRELESS_USAGE_DATA_POOLED_DOWNLINK                    UsageRecordLastMonthCategory = "wireless-usage-data-pooled-downlink"
	USAGERECORDLASTMONTHCATEGORY_WIRELESS_USAGE_DATA_POOLED_UPLINK                      UsageRecordLastMonthCategory = "wireless-usage-data-pooled-uplink"
	USAGERECORDLASTMONTHCATEGORY_WIRELESS_USAGE_MRC                                     UsageRecordLastMonthCategory = "wireless-usage-mrc"
	USAGERECORDLASTMONTHCATEGORY_WIRELESS_USAGE_MRC_CUSTOM                              UsageRecordLastMonthCategory = "wireless-usage-mrc-custom"
	USAGERECORDLASTMONTHCATEGORY_WIRELESS_USAGE_MRC_INDIVIDUAL                          UsageRecordLastMonthCategory = "wireless-usage-mrc-individual"
	USAGERECORDLASTMONTHCATEGORY_WIRELESS_USAGE_MRC_POOLED                              UsageRecordLastMonthCategory = "wireless-usage-mrc-pooled"
	USAGERECORDLASTMONTHCATEGORY_WIRELESS_USAGE_MRC_SUSPENDED                           UsageRecordLastMonthCategory = "wireless-usage-mrc-suspended"
	USAGERECORDLASTMONTHCATEGORY_WIRELESS_USAGE_SMS                                     UsageRecordLastMonthCategory = "wireless-usage-sms"
	USAGERECORDLASTMONTHCATEGORY_WIRELESS_USAGE_VOICE                                   UsageRecordLastMonthCategory = "wireless-usage-voice"
)
