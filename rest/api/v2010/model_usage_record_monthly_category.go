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

// UsageRecordMonthlyCategory the model 'UsageRecordMonthlyCategory'
type UsageRecordMonthlyCategory string

// List of usage_record_monthly_category
const (
	USAGERECORDMONTHLYCATEGORY_AGENT_CONFERENCE                                       UsageRecordMonthlyCategory = "agent-conference"
	USAGERECORDMONTHLYCATEGORY_ANSWERING_MACHINE_DETECTION                            UsageRecordMonthlyCategory = "answering-machine-detection"
	USAGERECORDMONTHLYCATEGORY_AUTHY_AUTHENTICATIONS                                  UsageRecordMonthlyCategory = "authy-authentications"
	USAGERECORDMONTHLYCATEGORY_AUTHY_CALLS_OUTBOUND                                   UsageRecordMonthlyCategory = "authy-calls-outbound"
	USAGERECORDMONTHLYCATEGORY_AUTHY_MONTHLY_FEES                                     UsageRecordMonthlyCategory = "authy-monthly-fees"
	USAGERECORDMONTHLYCATEGORY_AUTHY_PHONE_INTELLIGENCE                               UsageRecordMonthlyCategory = "authy-phone-intelligence"
	USAGERECORDMONTHLYCATEGORY_AUTHY_PHONE_VERIFICATIONS                              UsageRecordMonthlyCategory = "authy-phone-verifications"
	USAGERECORDMONTHLYCATEGORY_AUTHY_SMS_OUTBOUND                                     UsageRecordMonthlyCategory = "authy-sms-outbound"
	USAGERECORDMONTHLYCATEGORY_CALL_PROGESS_EVENTS                                    UsageRecordMonthlyCategory = "call-progess-events"
	USAGERECORDMONTHLYCATEGORY_CALLERIDLOOKUPS                                        UsageRecordMonthlyCategory = "calleridlookups"
	USAGERECORDMONTHLYCATEGORY_CALLS                                                  UsageRecordMonthlyCategory = "calls"
	USAGERECORDMONTHLYCATEGORY_CALLS_CLIENT                                           UsageRecordMonthlyCategory = "calls-client"
	USAGERECORDMONTHLYCATEGORY_CALLS_GLOBALCONFERENCE                                 UsageRecordMonthlyCategory = "calls-globalconference"
	USAGERECORDMONTHLYCATEGORY_CALLS_INBOUND                                          UsageRecordMonthlyCategory = "calls-inbound"
	USAGERECORDMONTHLYCATEGORY_CALLS_INBOUND_LOCAL                                    UsageRecordMonthlyCategory = "calls-inbound-local"
	USAGERECORDMONTHLYCATEGORY_CALLS_INBOUND_MOBILE                                   UsageRecordMonthlyCategory = "calls-inbound-mobile"
	USAGERECORDMONTHLYCATEGORY_CALLS_INBOUND_TOLLFREE                                 UsageRecordMonthlyCategory = "calls-inbound-tollfree"
	USAGERECORDMONTHLYCATEGORY_CALLS_OUTBOUND                                         UsageRecordMonthlyCategory = "calls-outbound"
	USAGERECORDMONTHLYCATEGORY_CALLS_PAY_VERB_TRANSACTIONS                            UsageRecordMonthlyCategory = "calls-pay-verb-transactions"
	USAGERECORDMONTHLYCATEGORY_CALLS_RECORDINGS                                       UsageRecordMonthlyCategory = "calls-recordings"
	USAGERECORDMONTHLYCATEGORY_CALLS_SIP                                              UsageRecordMonthlyCategory = "calls-sip"
	USAGERECORDMONTHLYCATEGORY_CALLS_SIP_INBOUND                                      UsageRecordMonthlyCategory = "calls-sip-inbound"
	USAGERECORDMONTHLYCATEGORY_CALLS_SIP_OUTBOUND                                     UsageRecordMonthlyCategory = "calls-sip-outbound"
	USAGERECORDMONTHLYCATEGORY_CARRIER_LOOKUPS                                        UsageRecordMonthlyCategory = "carrier-lookups"
	USAGERECORDMONTHLYCATEGORY_CONVERSATIONS                                          UsageRecordMonthlyCategory = "conversations"
	USAGERECORDMONTHLYCATEGORY_CONVERSATIONS_API_REQUESTS                             UsageRecordMonthlyCategory = "conversations-api-requests"
	USAGERECORDMONTHLYCATEGORY_CONVERSATIONS_CONVERSATION_EVENTS                      UsageRecordMonthlyCategory = "conversations-conversation-events"
	USAGERECORDMONTHLYCATEGORY_CONVERSATIONS_ENDPOINT_CONNECTIVITY                    UsageRecordMonthlyCategory = "conversations-endpoint-connectivity"
	USAGERECORDMONTHLYCATEGORY_CONVERSATIONS_EVENTS                                   UsageRecordMonthlyCategory = "conversations-events"
	USAGERECORDMONTHLYCATEGORY_CONVERSATIONS_PARTICIPANT_EVENTS                       UsageRecordMonthlyCategory = "conversations-participant-events"
	USAGERECORDMONTHLYCATEGORY_CONVERSATIONS_PARTICIPANTS                             UsageRecordMonthlyCategory = "conversations-participants"
	USAGERECORDMONTHLYCATEGORY_CPS                                                    UsageRecordMonthlyCategory = "cps"
	USAGERECORDMONTHLYCATEGORY_FRAUD_LOOKUPS                                          UsageRecordMonthlyCategory = "fraud-lookups"
	USAGERECORDMONTHLYCATEGORY_GROUP_ROOMS                                            UsageRecordMonthlyCategory = "group-rooms"
	USAGERECORDMONTHLYCATEGORY_GROUP_ROOMS_DATA_TRACK                                 UsageRecordMonthlyCategory = "group-rooms-data-track"
	USAGERECORDMONTHLYCATEGORY_GROUP_ROOMS_ENCRYPTED_MEDIA_RECORDED                   UsageRecordMonthlyCategory = "group-rooms-encrypted-media-recorded"
	USAGERECORDMONTHLYCATEGORY_GROUP_ROOMS_MEDIA_DOWNLOADED                           UsageRecordMonthlyCategory = "group-rooms-media-downloaded"
	USAGERECORDMONTHLYCATEGORY_GROUP_ROOMS_MEDIA_RECORDED                             UsageRecordMonthlyCategory = "group-rooms-media-recorded"
	USAGERECORDMONTHLYCATEGORY_GROUP_ROOMS_MEDIA_ROUTED                               UsageRecordMonthlyCategory = "group-rooms-media-routed"
	USAGERECORDMONTHLYCATEGORY_GROUP_ROOMS_MEDIA_STORED                               UsageRecordMonthlyCategory = "group-rooms-media-stored"
	USAGERECORDMONTHLYCATEGORY_GROUP_ROOMS_PARTICIPANT_MINUTES                        UsageRecordMonthlyCategory = "group-rooms-participant-minutes"
	USAGERECORDMONTHLYCATEGORY_GROUP_ROOMS_RECORDED_MINUTES                           UsageRecordMonthlyCategory = "group-rooms-recorded-minutes"
	USAGERECORDMONTHLYCATEGORY_IMP_V1_USAGE                                           UsageRecordMonthlyCategory = "imp-v1-usage"
	USAGERECORDMONTHLYCATEGORY_LOOKUPS                                                UsageRecordMonthlyCategory = "lookups"
	USAGERECORDMONTHLYCATEGORY_MARKETPLACE                                            UsageRecordMonthlyCategory = "marketplace"
	USAGERECORDMONTHLYCATEGORY_MARKETPLACE_ALGORITHMIA_NAMED_ENTITY_RECOGNITION       UsageRecordMonthlyCategory = "marketplace-algorithmia-named-entity-recognition"
	USAGERECORDMONTHLYCATEGORY_MARKETPLACE_CADENCE_TRANSCRIPTION                      UsageRecordMonthlyCategory = "marketplace-cadence-transcription"
	USAGERECORDMONTHLYCATEGORY_MARKETPLACE_CADENCE_TRANSLATION                        UsageRecordMonthlyCategory = "marketplace-cadence-translation"
	USAGERECORDMONTHLYCATEGORY_MARKETPLACE_CAPIO_SPEECH_TO_TEXT                       UsageRecordMonthlyCategory = "marketplace-capio-speech-to-text"
	USAGERECORDMONTHLYCATEGORY_MARKETPLACE_CONVRIZA_ABABA                             UsageRecordMonthlyCategory = "marketplace-convriza-ababa"
	USAGERECORDMONTHLYCATEGORY_MARKETPLACE_DEEPGRAM_PHRASE_DETECTOR                   UsageRecordMonthlyCategory = "marketplace-deepgram-phrase-detector"
	USAGERECORDMONTHLYCATEGORY_MARKETPLACE_DIGITAL_SEGMENT_BUSINESS_INFO              UsageRecordMonthlyCategory = "marketplace-digital-segment-business-info"
	USAGERECORDMONTHLYCATEGORY_MARKETPLACE_FACEBOOK_OFFLINE_CONVERSIONS               UsageRecordMonthlyCategory = "marketplace-facebook-offline-conversions"
	USAGERECORDMONTHLYCATEGORY_MARKETPLACE_GOOGLE_SPEECH_TO_TEXT                      UsageRecordMonthlyCategory = "marketplace-google-speech-to-text"
	USAGERECORDMONTHLYCATEGORY_MARKETPLACE_IBM_WATSON_MESSAGE_INSIGHTS                UsageRecordMonthlyCategory = "marketplace-ibm-watson-message-insights"
	USAGERECORDMONTHLYCATEGORY_MARKETPLACE_IBM_WATSON_MESSAGE_SENTIMENT               UsageRecordMonthlyCategory = "marketplace-ibm-watson-message-sentiment"
	USAGERECORDMONTHLYCATEGORY_MARKETPLACE_IBM_WATSON_RECORDING_ANALYSIS              UsageRecordMonthlyCategory = "marketplace-ibm-watson-recording-analysis"
	USAGERECORDMONTHLYCATEGORY_MARKETPLACE_IBM_WATSON_TONE_ANALYZER                   UsageRecordMonthlyCategory = "marketplace-ibm-watson-tone-analyzer"
	USAGERECORDMONTHLYCATEGORY_MARKETPLACE_ICEHOOK_SYSTEMS_SCOUT                      UsageRecordMonthlyCategory = "marketplace-icehook-systems-scout"
	USAGERECORDMONTHLYCATEGORY_MARKETPLACE_INFOGROUP_DATAAXLE_BIZINFO                 UsageRecordMonthlyCategory = "marketplace-infogroup-dataaxle-bizinfo"
	USAGERECORDMONTHLYCATEGORY_MARKETPLACE_KEEN_IO_CONTACT_CENTER_ANALYTICS           UsageRecordMonthlyCategory = "marketplace-keen-io-contact-center-analytics"
	USAGERECORDMONTHLYCATEGORY_MARKETPLACE_MARCHEX_CLEANCALL                          UsageRecordMonthlyCategory = "marketplace-marchex-cleancall"
	USAGERECORDMONTHLYCATEGORY_MARKETPLACE_MARCHEX_SENTIMENT_ANALYSIS_FOR_SMS         UsageRecordMonthlyCategory = "marketplace-marchex-sentiment-analysis-for-sms"
	USAGERECORDMONTHLYCATEGORY_MARKETPLACE_MARKETPLACE_NEXTCALLER_SOCIAL_ID           UsageRecordMonthlyCategory = "marketplace-marketplace-nextcaller-social-id"
	USAGERECORDMONTHLYCATEGORY_MARKETPLACE_MOBILE_COMMONS_OPT_OUT_CLASSIFIER          UsageRecordMonthlyCategory = "marketplace-mobile-commons-opt-out-classifier"
	USAGERECORDMONTHLYCATEGORY_MARKETPLACE_NEXIWAVE_VOICEMAIL_TO_TEXT                 UsageRecordMonthlyCategory = "marketplace-nexiwave-voicemail-to-text"
	USAGERECORDMONTHLYCATEGORY_MARKETPLACE_NEXTCALLER_ADVANCED_CALLER_IDENTIFICATION  UsageRecordMonthlyCategory = "marketplace-nextcaller-advanced-caller-identification"
	USAGERECORDMONTHLYCATEGORY_MARKETPLACE_NOMOROBO_SPAM_SCORE                        UsageRecordMonthlyCategory = "marketplace-nomorobo-spam-score"
	USAGERECORDMONTHLYCATEGORY_MARKETPLACE_PAYFONE_TCPA_COMPLIANCE                    UsageRecordMonthlyCategory = "marketplace-payfone-tcpa-compliance"
	USAGERECORDMONTHLYCATEGORY_MARKETPLACE_REMEETING_AUTOMATIC_SPEECH_RECOGNITION     UsageRecordMonthlyCategory = "marketplace-remeeting-automatic-speech-recognition"
	USAGERECORDMONTHLYCATEGORY_MARKETPLACE_TCPA_DEFENSE_SOLUTIONS_BLACKLIST_FEED      UsageRecordMonthlyCategory = "marketplace-tcpa-defense-solutions-blacklist-feed"
	USAGERECORDMONTHLYCATEGORY_MARKETPLACE_TELO_OPENCNAM                              UsageRecordMonthlyCategory = "marketplace-telo-opencnam"
	USAGERECORDMONTHLYCATEGORY_MARKETPLACE_TRUECNAM_TRUE_SPAM                         UsageRecordMonthlyCategory = "marketplace-truecnam-true-spam"
	USAGERECORDMONTHLYCATEGORY_MARKETPLACE_TWILIO_CALLER_NAME_LOOKUP_US               UsageRecordMonthlyCategory = "marketplace-twilio-caller-name-lookup-us"
	USAGERECORDMONTHLYCATEGORY_MARKETPLACE_TWILIO_CARRIER_INFORMATION_LOOKUP          UsageRecordMonthlyCategory = "marketplace-twilio-carrier-information-lookup"
	USAGERECORDMONTHLYCATEGORY_MARKETPLACE_VOICEBASE_PCI                              UsageRecordMonthlyCategory = "marketplace-voicebase-pci"
	USAGERECORDMONTHLYCATEGORY_MARKETPLACE_VOICEBASE_TRANSCRIPTION                    UsageRecordMonthlyCategory = "marketplace-voicebase-transcription"
	USAGERECORDMONTHLYCATEGORY_MARKETPLACE_VOICEBASE_TRANSCRIPTION_CUSTOM_VOCABULARY  UsageRecordMonthlyCategory = "marketplace-voicebase-transcription-custom-vocabulary"
	USAGERECORDMONTHLYCATEGORY_MARKETPLACE_WHITEPAGES_PRO_CALLER_IDENTIFICATION       UsageRecordMonthlyCategory = "marketplace-whitepages-pro-caller-identification"
	USAGERECORDMONTHLYCATEGORY_MARKETPLACE_WHITEPAGES_PRO_PHONE_INTELLIGENCE          UsageRecordMonthlyCategory = "marketplace-whitepages-pro-phone-intelligence"
	USAGERECORDMONTHLYCATEGORY_MARKETPLACE_WHITEPAGES_PRO_PHONE_REPUTATION            UsageRecordMonthlyCategory = "marketplace-whitepages-pro-phone-reputation"
	USAGERECORDMONTHLYCATEGORY_MARKETPLACE_WOLFARM_SPOKEN_RESULTS                     UsageRecordMonthlyCategory = "marketplace-wolfarm-spoken-results"
	USAGERECORDMONTHLYCATEGORY_MARKETPLACE_WOLFRAM_SHORT_ANSWER                       UsageRecordMonthlyCategory = "marketplace-wolfram-short-answer"
	USAGERECORDMONTHLYCATEGORY_MARKETPLACE_YTICA_CONTACT_CENTER_REPORTING_ANALYTICS   UsageRecordMonthlyCategory = "marketplace-ytica-contact-center-reporting-analytics"
	USAGERECORDMONTHLYCATEGORY_MEDIASTORAGE                                           UsageRecordMonthlyCategory = "mediastorage"
	USAGERECORDMONTHLYCATEGORY_MMS                                                    UsageRecordMonthlyCategory = "mms"
	USAGERECORDMONTHLYCATEGORY_MMS_INBOUND                                            UsageRecordMonthlyCategory = "mms-inbound"
	USAGERECORDMONTHLYCATEGORY_MMS_INBOUND_LONGCODE                                   UsageRecordMonthlyCategory = "mms-inbound-longcode"
	USAGERECORDMONTHLYCATEGORY_MMS_INBOUND_SHORTCODE                                  UsageRecordMonthlyCategory = "mms-inbound-shortcode"
	USAGERECORDMONTHLYCATEGORY_MMS_MESSAGES_CARRIERFEES                               UsageRecordMonthlyCategory = "mms-messages-carrierfees"
	USAGERECORDMONTHLYCATEGORY_MMS_OUTBOUND                                           UsageRecordMonthlyCategory = "mms-outbound"
	USAGERECORDMONTHLYCATEGORY_MMS_OUTBOUND_LONGCODE                                  UsageRecordMonthlyCategory = "mms-outbound-longcode"
	USAGERECORDMONTHLYCATEGORY_MMS_OUTBOUND_SHORTCODE                                 UsageRecordMonthlyCategory = "mms-outbound-shortcode"
	USAGERECORDMONTHLYCATEGORY_MONITOR_READS                                          UsageRecordMonthlyCategory = "monitor-reads"
	USAGERECORDMONTHLYCATEGORY_MONITOR_STORAGE                                        UsageRecordMonthlyCategory = "monitor-storage"
	USAGERECORDMONTHLYCATEGORY_MONITOR_WRITES                                         UsageRecordMonthlyCategory = "monitor-writes"
	USAGERECORDMONTHLYCATEGORY_NOTIFY                                                 UsageRecordMonthlyCategory = "notify"
	USAGERECORDMONTHLYCATEGORY_NOTIFY_ACTIONS_ATTEMPTS                                UsageRecordMonthlyCategory = "notify-actions-attempts"
	USAGERECORDMONTHLYCATEGORY_NOTIFY_CHANNELS                                        UsageRecordMonthlyCategory = "notify-channels"
	USAGERECORDMONTHLYCATEGORY_NUMBER_FORMAT_LOOKUPS                                  UsageRecordMonthlyCategory = "number-format-lookups"
	USAGERECORDMONTHLYCATEGORY_PCHAT                                                  UsageRecordMonthlyCategory = "pchat"
	USAGERECORDMONTHLYCATEGORY_PCHAT_USERS                                            UsageRecordMonthlyCategory = "pchat-users"
	USAGERECORDMONTHLYCATEGORY_PEER_TO_PEER_ROOMS_PARTICIPANT_MINUTES                 UsageRecordMonthlyCategory = "peer-to-peer-rooms-participant-minutes"
	USAGERECORDMONTHLYCATEGORY_PFAX                                                   UsageRecordMonthlyCategory = "pfax"
	USAGERECORDMONTHLYCATEGORY_PFAX_MINUTES                                           UsageRecordMonthlyCategory = "pfax-minutes"
	USAGERECORDMONTHLYCATEGORY_PFAX_MINUTES_INBOUND                                   UsageRecordMonthlyCategory = "pfax-minutes-inbound"
	USAGERECORDMONTHLYCATEGORY_PFAX_MINUTES_OUTBOUND                                  UsageRecordMonthlyCategory = "pfax-minutes-outbound"
	USAGERECORDMONTHLYCATEGORY_PFAX_PAGES                                             UsageRecordMonthlyCategory = "pfax-pages"
	USAGERECORDMONTHLYCATEGORY_PHONENUMBERS                                           UsageRecordMonthlyCategory = "phonenumbers"
	USAGERECORDMONTHLYCATEGORY_PHONENUMBERS_CPS                                       UsageRecordMonthlyCategory = "phonenumbers-cps"
	USAGERECORDMONTHLYCATEGORY_PHONENUMBERS_EMERGENCY                                 UsageRecordMonthlyCategory = "phonenumbers-emergency"
	USAGERECORDMONTHLYCATEGORY_PHONENUMBERS_LOCAL                                     UsageRecordMonthlyCategory = "phonenumbers-local"
	USAGERECORDMONTHLYCATEGORY_PHONENUMBERS_MOBILE                                    UsageRecordMonthlyCategory = "phonenumbers-mobile"
	USAGERECORDMONTHLYCATEGORY_PHONENUMBERS_SETUPS                                    UsageRecordMonthlyCategory = "phonenumbers-setups"
	USAGERECORDMONTHLYCATEGORY_PHONENUMBERS_TOLLFREE                                  UsageRecordMonthlyCategory = "phonenumbers-tollfree"
	USAGERECORDMONTHLYCATEGORY_PREMIUMSUPPORT                                         UsageRecordMonthlyCategory = "premiumsupport"
	USAGERECORDMONTHLYCATEGORY_PROXY                                                  UsageRecordMonthlyCategory = "proxy"
	USAGERECORDMONTHLYCATEGORY_PROXY_ACTIVE_SESSIONS                                  UsageRecordMonthlyCategory = "proxy-active-sessions"
	USAGERECORDMONTHLYCATEGORY_PSTNCONNECTIVITY                                       UsageRecordMonthlyCategory = "pstnconnectivity"
	USAGERECORDMONTHLYCATEGORY_PV                                                     UsageRecordMonthlyCategory = "pv"
	USAGERECORDMONTHLYCATEGORY_PV_COMPOSITION_MEDIA_DOWNLOADED                        UsageRecordMonthlyCategory = "pv-composition-media-downloaded"
	USAGERECORDMONTHLYCATEGORY_PV_COMPOSITION_MEDIA_ENCRYPTED                         UsageRecordMonthlyCategory = "pv-composition-media-encrypted"
	USAGERECORDMONTHLYCATEGORY_PV_COMPOSITION_MEDIA_STORED                            UsageRecordMonthlyCategory = "pv-composition-media-stored"
	USAGERECORDMONTHLYCATEGORY_PV_COMPOSITION_MINUTES                                 UsageRecordMonthlyCategory = "pv-composition-minutes"
	USAGERECORDMONTHLYCATEGORY_PV_RECORDING_COMPOSITIONS                              UsageRecordMonthlyCategory = "pv-recording-compositions"
	USAGERECORDMONTHLYCATEGORY_PV_ROOM_PARTICIPANTS                                   UsageRecordMonthlyCategory = "pv-room-participants"
	USAGERECORDMONTHLYCATEGORY_PV_ROOM_PARTICIPANTS_AU1                               UsageRecordMonthlyCategory = "pv-room-participants-au1"
	USAGERECORDMONTHLYCATEGORY_PV_ROOM_PARTICIPANTS_BR1                               UsageRecordMonthlyCategory = "pv-room-participants-br1"
	USAGERECORDMONTHLYCATEGORY_PV_ROOM_PARTICIPANTS_IE1                               UsageRecordMonthlyCategory = "pv-room-participants-ie1"
	USAGERECORDMONTHLYCATEGORY_PV_ROOM_PARTICIPANTS_JP1                               UsageRecordMonthlyCategory = "pv-room-participants-jp1"
	USAGERECORDMONTHLYCATEGORY_PV_ROOM_PARTICIPANTS_SG1                               UsageRecordMonthlyCategory = "pv-room-participants-sg1"
	USAGERECORDMONTHLYCATEGORY_PV_ROOM_PARTICIPANTS_US1                               UsageRecordMonthlyCategory = "pv-room-participants-us1"
	USAGERECORDMONTHLYCATEGORY_PV_ROOM_PARTICIPANTS_US2                               UsageRecordMonthlyCategory = "pv-room-participants-us2"
	USAGERECORDMONTHLYCATEGORY_PV_ROOMS                                               UsageRecordMonthlyCategory = "pv-rooms"
	USAGERECORDMONTHLYCATEGORY_PV_SIP_ENDPOINT_REGISTRATIONS                          UsageRecordMonthlyCategory = "pv-sip-endpoint-registrations"
	USAGERECORDMONTHLYCATEGORY_RECORDINGS                                             UsageRecordMonthlyCategory = "recordings"
	USAGERECORDMONTHLYCATEGORY_RECORDINGSTORAGE                                       UsageRecordMonthlyCategory = "recordingstorage"
	USAGERECORDMONTHLYCATEGORY_ROOMS_GROUP_BANDWIDTH                                  UsageRecordMonthlyCategory = "rooms-group-bandwidth"
	USAGERECORDMONTHLYCATEGORY_ROOMS_GROUP_MINUTES                                    UsageRecordMonthlyCategory = "rooms-group-minutes"
	USAGERECORDMONTHLYCATEGORY_ROOMS_PEER_TO_PEER_MINUTES                             UsageRecordMonthlyCategory = "rooms-peer-to-peer-minutes"
	USAGERECORDMONTHLYCATEGORY_SHORTCODES                                             UsageRecordMonthlyCategory = "shortcodes"
	USAGERECORDMONTHLYCATEGORY_SHORTCODES_CUSTOMEROWNED                               UsageRecordMonthlyCategory = "shortcodes-customerowned"
	USAGERECORDMONTHLYCATEGORY_SHORTCODES_MMS_ENABLEMENT                              UsageRecordMonthlyCategory = "shortcodes-mms-enablement"
	USAGERECORDMONTHLYCATEGORY_SHORTCODES_MPS                                         UsageRecordMonthlyCategory = "shortcodes-mps"
	USAGERECORDMONTHLYCATEGORY_SHORTCODES_RANDOM                                      UsageRecordMonthlyCategory = "shortcodes-random"
	USAGERECORDMONTHLYCATEGORY_SHORTCODES_UK                                          UsageRecordMonthlyCategory = "shortcodes-uk"
	USAGERECORDMONTHLYCATEGORY_SHORTCODES_VANITY                                      UsageRecordMonthlyCategory = "shortcodes-vanity"
	USAGERECORDMONTHLYCATEGORY_SMALL_GROUP_ROOMS                                      UsageRecordMonthlyCategory = "small-group-rooms"
	USAGERECORDMONTHLYCATEGORY_SMALL_GROUP_ROOMS_DATA_TRACK                           UsageRecordMonthlyCategory = "small-group-rooms-data-track"
	USAGERECORDMONTHLYCATEGORY_SMALL_GROUP_ROOMS_PARTICIPANT_MINUTES                  UsageRecordMonthlyCategory = "small-group-rooms-participant-minutes"
	USAGERECORDMONTHLYCATEGORY_SMS                                                    UsageRecordMonthlyCategory = "sms"
	USAGERECORDMONTHLYCATEGORY_SMS_INBOUND                                            UsageRecordMonthlyCategory = "sms-inbound"
	USAGERECORDMONTHLYCATEGORY_SMS_INBOUND_LONGCODE                                   UsageRecordMonthlyCategory = "sms-inbound-longcode"
	USAGERECORDMONTHLYCATEGORY_SMS_INBOUND_SHORTCODE                                  UsageRecordMonthlyCategory = "sms-inbound-shortcode"
	USAGERECORDMONTHLYCATEGORY_SMS_MESSAGES_CARRIERFEES                               UsageRecordMonthlyCategory = "sms-messages-carrierfees"
	USAGERECORDMONTHLYCATEGORY_SMS_MESSAGES_FEATURES                                  UsageRecordMonthlyCategory = "sms-messages-features"
	USAGERECORDMONTHLYCATEGORY_SMS_MESSAGES_FEATURES_SENDERID                         UsageRecordMonthlyCategory = "sms-messages-features-senderid"
	USAGERECORDMONTHLYCATEGORY_SMS_OUTBOUND                                           UsageRecordMonthlyCategory = "sms-outbound"
	USAGERECORDMONTHLYCATEGORY_SMS_OUTBOUND_CONTENT_INSPECTION                        UsageRecordMonthlyCategory = "sms-outbound-content-inspection"
	USAGERECORDMONTHLYCATEGORY_SMS_OUTBOUND_LONGCODE                                  UsageRecordMonthlyCategory = "sms-outbound-longcode"
	USAGERECORDMONTHLYCATEGORY_SMS_OUTBOUND_SHORTCODE                                 UsageRecordMonthlyCategory = "sms-outbound-shortcode"
	USAGERECORDMONTHLYCATEGORY_SPEECH_RECOGNITION                                     UsageRecordMonthlyCategory = "speech-recognition"
	USAGERECORDMONTHLYCATEGORY_STUDIO_ENGAGEMENTS                                     UsageRecordMonthlyCategory = "studio-engagements"
	USAGERECORDMONTHLYCATEGORY_SYNC                                                   UsageRecordMonthlyCategory = "sync"
	USAGERECORDMONTHLYCATEGORY_SYNC_ACTIONS                                           UsageRecordMonthlyCategory = "sync-actions"
	USAGERECORDMONTHLYCATEGORY_SYNC_ENDPOINT_HOURS                                    UsageRecordMonthlyCategory = "sync-endpoint-hours"
	USAGERECORDMONTHLYCATEGORY_SYNC_ENDPOINT_HOURS_ABOVE_DAILY_CAP                    UsageRecordMonthlyCategory = "sync-endpoint-hours-above-daily-cap"
	USAGERECORDMONTHLYCATEGORY_TASKROUTER_TASKS                                       UsageRecordMonthlyCategory = "taskrouter-tasks"
	USAGERECORDMONTHLYCATEGORY_TOTALPRICE                                             UsageRecordMonthlyCategory = "totalprice"
	USAGERECORDMONTHLYCATEGORY_TRANSCRIPTIONS                                         UsageRecordMonthlyCategory = "transcriptions"
	USAGERECORDMONTHLYCATEGORY_TRUNKING_CPS                                           UsageRecordMonthlyCategory = "trunking-cps"
	USAGERECORDMONTHLYCATEGORY_TRUNKING_EMERGENCY_CALLS                               UsageRecordMonthlyCategory = "trunking-emergency-calls"
	USAGERECORDMONTHLYCATEGORY_TRUNKING_ORIGINATION                                   UsageRecordMonthlyCategory = "trunking-origination"
	USAGERECORDMONTHLYCATEGORY_TRUNKING_ORIGINATION_LOCAL                             UsageRecordMonthlyCategory = "trunking-origination-local"
	USAGERECORDMONTHLYCATEGORY_TRUNKING_ORIGINATION_MOBILE                            UsageRecordMonthlyCategory = "trunking-origination-mobile"
	USAGERECORDMONTHLYCATEGORY_TRUNKING_ORIGINATION_TOLLFREE                          UsageRecordMonthlyCategory = "trunking-origination-tollfree"
	USAGERECORDMONTHLYCATEGORY_TRUNKING_RECORDINGS                                    UsageRecordMonthlyCategory = "trunking-recordings"
	USAGERECORDMONTHLYCATEGORY_TRUNKING_SECURE                                        UsageRecordMonthlyCategory = "trunking-secure"
	USAGERECORDMONTHLYCATEGORY_TRUNKING_TERMINATION                                   UsageRecordMonthlyCategory = "trunking-termination"
	USAGERECORDMONTHLYCATEGORY_TURNMEGABYTES                                          UsageRecordMonthlyCategory = "turnmegabytes"
	USAGERECORDMONTHLYCATEGORY_TURNMEGABYTES_AUSTRALIA                                UsageRecordMonthlyCategory = "turnmegabytes-australia"
	USAGERECORDMONTHLYCATEGORY_TURNMEGABYTES_BRASIL                                   UsageRecordMonthlyCategory = "turnmegabytes-brasil"
	USAGERECORDMONTHLYCATEGORY_TURNMEGABYTES_GERMANY                                  UsageRecordMonthlyCategory = "turnmegabytes-germany"
	USAGERECORDMONTHLYCATEGORY_TURNMEGABYTES_INDIA                                    UsageRecordMonthlyCategory = "turnmegabytes-india"
	USAGERECORDMONTHLYCATEGORY_TURNMEGABYTES_IRELAND                                  UsageRecordMonthlyCategory = "turnmegabytes-ireland"
	USAGERECORDMONTHLYCATEGORY_TURNMEGABYTES_JAPAN                                    UsageRecordMonthlyCategory = "turnmegabytes-japan"
	USAGERECORDMONTHLYCATEGORY_TURNMEGABYTES_SINGAPORE                                UsageRecordMonthlyCategory = "turnmegabytes-singapore"
	USAGERECORDMONTHLYCATEGORY_TURNMEGABYTES_USEAST                                   UsageRecordMonthlyCategory = "turnmegabytes-useast"
	USAGERECORDMONTHLYCATEGORY_TURNMEGABYTES_USWEST                                   UsageRecordMonthlyCategory = "turnmegabytes-uswest"
	USAGERECORDMONTHLYCATEGORY_TWILIO_INTERCONNECT                                    UsageRecordMonthlyCategory = "twilio-interconnect"
	USAGERECORDMONTHLYCATEGORY_VERIFY_PUSH                                            UsageRecordMonthlyCategory = "verify-push"
	USAGERECORDMONTHLYCATEGORY_VIDEO_RECORDINGS                                       UsageRecordMonthlyCategory = "video-recordings"
	USAGERECORDMONTHLYCATEGORY_VOICE_INSIGHTS                                         UsageRecordMonthlyCategory = "voice-insights"
	USAGERECORDMONTHLYCATEGORY_VOICE_INSIGHTS_CLIENT_INSIGHTS_ON_DEMAND_MINUTE        UsageRecordMonthlyCategory = "voice-insights-client-insights-on-demand-minute"
	USAGERECORDMONTHLYCATEGORY_VOICE_INSIGHTS_PTSN_INSIGHTS_ON_DEMAND_MINUTE          UsageRecordMonthlyCategory = "voice-insights-ptsn-insights-on-demand-minute"
	USAGERECORDMONTHLYCATEGORY_VOICE_INSIGHTS_SIP_INTERFACE_INSIGHTS_ON_DEMAND_MINUTE UsageRecordMonthlyCategory = "voice-insights-sip-interface-insights-on-demand-minute"
	USAGERECORDMONTHLYCATEGORY_VOICE_INSIGHTS_SIP_TRUNKING_INSIGHTS_ON_DEMAND_MINUTE  UsageRecordMonthlyCategory = "voice-insights-sip-trunking-insights-on-demand-minute"
	USAGERECORDMONTHLYCATEGORY_WIRELESS                                               UsageRecordMonthlyCategory = "wireless"
	USAGERECORDMONTHLYCATEGORY_WIRELESS_ORDERS                                        UsageRecordMonthlyCategory = "wireless-orders"
	USAGERECORDMONTHLYCATEGORY_WIRELESS_ORDERS_ARTWORK                                UsageRecordMonthlyCategory = "wireless-orders-artwork"
	USAGERECORDMONTHLYCATEGORY_WIRELESS_ORDERS_BULK                                   UsageRecordMonthlyCategory = "wireless-orders-bulk"
	USAGERECORDMONTHLYCATEGORY_WIRELESS_ORDERS_ESIM                                   UsageRecordMonthlyCategory = "wireless-orders-esim"
	USAGERECORDMONTHLYCATEGORY_WIRELESS_ORDERS_STARTER                                UsageRecordMonthlyCategory = "wireless-orders-starter"
	USAGERECORDMONTHLYCATEGORY_WIRELESS_USAGE                                         UsageRecordMonthlyCategory = "wireless-usage"
	USAGERECORDMONTHLYCATEGORY_WIRELESS_USAGE_COMMANDS                                UsageRecordMonthlyCategory = "wireless-usage-commands"
	USAGERECORDMONTHLYCATEGORY_WIRELESS_USAGE_COMMANDS_AFRICA                         UsageRecordMonthlyCategory = "wireless-usage-commands-africa"
	USAGERECORDMONTHLYCATEGORY_WIRELESS_USAGE_COMMANDS_ASIA                           UsageRecordMonthlyCategory = "wireless-usage-commands-asia"
	USAGERECORDMONTHLYCATEGORY_WIRELESS_USAGE_COMMANDS_CENTRALANDSOUTHAMERICA         UsageRecordMonthlyCategory = "wireless-usage-commands-centralandsouthamerica"
	USAGERECORDMONTHLYCATEGORY_WIRELESS_USAGE_COMMANDS_EUROPE                         UsageRecordMonthlyCategory = "wireless-usage-commands-europe"
	USAGERECORDMONTHLYCATEGORY_WIRELESS_USAGE_COMMANDS_HOME                           UsageRecordMonthlyCategory = "wireless-usage-commands-home"
	USAGERECORDMONTHLYCATEGORY_WIRELESS_USAGE_COMMANDS_NORTHAMERICA                   UsageRecordMonthlyCategory = "wireless-usage-commands-northamerica"
	USAGERECORDMONTHLYCATEGORY_WIRELESS_USAGE_COMMANDS_OCEANIA                        UsageRecordMonthlyCategory = "wireless-usage-commands-oceania"
	USAGERECORDMONTHLYCATEGORY_WIRELESS_USAGE_COMMANDS_ROAMING                        UsageRecordMonthlyCategory = "wireless-usage-commands-roaming"
	USAGERECORDMONTHLYCATEGORY_WIRELESS_USAGE_DATA                                    UsageRecordMonthlyCategory = "wireless-usage-data"
	USAGERECORDMONTHLYCATEGORY_WIRELESS_USAGE_DATA_AFRICA                             UsageRecordMonthlyCategory = "wireless-usage-data-africa"
	USAGERECORDMONTHLYCATEGORY_WIRELESS_USAGE_DATA_ASIA                               UsageRecordMonthlyCategory = "wireless-usage-data-asia"
	USAGERECORDMONTHLYCATEGORY_WIRELESS_USAGE_DATA_CENTRALANDSOUTHAMERICA             UsageRecordMonthlyCategory = "wireless-usage-data-centralandsouthamerica"
	USAGERECORDMONTHLYCATEGORY_WIRELESS_USAGE_DATA_CUSTOM_ADDITIONALMB                UsageRecordMonthlyCategory = "wireless-usage-data-custom-additionalmb"
	USAGERECORDMONTHLYCATEGORY_WIRELESS_USAGE_DATA_CUSTOM_FIRST5MB                    UsageRecordMonthlyCategory = "wireless-usage-data-custom-first5mb"
	USAGERECORDMONTHLYCATEGORY_WIRELESS_USAGE_DATA_DOMESTIC_ROAMING                   UsageRecordMonthlyCategory = "wireless-usage-data-domestic-roaming"
	USAGERECORDMONTHLYCATEGORY_WIRELESS_USAGE_DATA_EUROPE                             UsageRecordMonthlyCategory = "wireless-usage-data-europe"
	USAGERECORDMONTHLYCATEGORY_WIRELESS_USAGE_DATA_INDIVIDUAL_ADDITIONALGB            UsageRecordMonthlyCategory = "wireless-usage-data-individual-additionalgb"
	USAGERECORDMONTHLYCATEGORY_WIRELESS_USAGE_DATA_INDIVIDUAL_FIRSTGB                 UsageRecordMonthlyCategory = "wireless-usage-data-individual-firstgb"
	USAGERECORDMONTHLYCATEGORY_WIRELESS_USAGE_DATA_INTERNATIONAL_ROAMING_CANADA       UsageRecordMonthlyCategory = "wireless-usage-data-international-roaming-canada"
	USAGERECORDMONTHLYCATEGORY_WIRELESS_USAGE_DATA_INTERNATIONAL_ROAMING_INDIA        UsageRecordMonthlyCategory = "wireless-usage-data-international-roaming-india"
	USAGERECORDMONTHLYCATEGORY_WIRELESS_USAGE_DATA_INTERNATIONAL_ROAMING_MEXICO       UsageRecordMonthlyCategory = "wireless-usage-data-international-roaming-mexico"
	USAGERECORDMONTHLYCATEGORY_WIRELESS_USAGE_DATA_NORTHAMERICA                       UsageRecordMonthlyCategory = "wireless-usage-data-northamerica"
	USAGERECORDMONTHLYCATEGORY_WIRELESS_USAGE_DATA_OCEANIA                            UsageRecordMonthlyCategory = "wireless-usage-data-oceania"
	USAGERECORDMONTHLYCATEGORY_WIRELESS_USAGE_DATA_POOLED                             UsageRecordMonthlyCategory = "wireless-usage-data-pooled"
	USAGERECORDMONTHLYCATEGORY_WIRELESS_USAGE_DATA_POOLED_DOWNLINK                    UsageRecordMonthlyCategory = "wireless-usage-data-pooled-downlink"
	USAGERECORDMONTHLYCATEGORY_WIRELESS_USAGE_DATA_POOLED_UPLINK                      UsageRecordMonthlyCategory = "wireless-usage-data-pooled-uplink"
	USAGERECORDMONTHLYCATEGORY_WIRELESS_USAGE_MRC                                     UsageRecordMonthlyCategory = "wireless-usage-mrc"
	USAGERECORDMONTHLYCATEGORY_WIRELESS_USAGE_MRC_CUSTOM                              UsageRecordMonthlyCategory = "wireless-usage-mrc-custom"
	USAGERECORDMONTHLYCATEGORY_WIRELESS_USAGE_MRC_INDIVIDUAL                          UsageRecordMonthlyCategory = "wireless-usage-mrc-individual"
	USAGERECORDMONTHLYCATEGORY_WIRELESS_USAGE_MRC_POOLED                              UsageRecordMonthlyCategory = "wireless-usage-mrc-pooled"
	USAGERECORDMONTHLYCATEGORY_WIRELESS_USAGE_MRC_SUSPENDED                           UsageRecordMonthlyCategory = "wireless-usage-mrc-suspended"
	USAGERECORDMONTHLYCATEGORY_WIRELESS_USAGE_SMS                                     UsageRecordMonthlyCategory = "wireless-usage-sms"
	USAGERECORDMONTHLYCATEGORY_WIRELESS_USAGE_VOICE                                   UsageRecordMonthlyCategory = "wireless-usage-voice"
)
