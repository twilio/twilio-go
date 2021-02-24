/*
 * Twilio - Api
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.10.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// UsageRecordCategory the model 'UsageRecordCategory'
type UsageRecordCategory string

// List of usage_record_category
const (
	USAGERECORDCATEGORY_AGENT_CONFERENCE                                       UsageRecordCategory = "agent-conference"
	USAGERECORDCATEGORY_ANSWERING_MACHINE_DETECTION                            UsageRecordCategory = "answering-machine-detection"
	USAGERECORDCATEGORY_AUTHY_AUTHENTICATIONS                                  UsageRecordCategory = "authy-authentications"
	USAGERECORDCATEGORY_AUTHY_CALLS_OUTBOUND                                   UsageRecordCategory = "authy-calls-outbound"
	USAGERECORDCATEGORY_AUTHY_MONTHLY_FEES                                     UsageRecordCategory = "authy-monthly-fees"
	USAGERECORDCATEGORY_AUTHY_PHONE_INTELLIGENCE                               UsageRecordCategory = "authy-phone-intelligence"
	USAGERECORDCATEGORY_AUTHY_PHONE_VERIFICATIONS                              UsageRecordCategory = "authy-phone-verifications"
	USAGERECORDCATEGORY_AUTHY_SMS_OUTBOUND                                     UsageRecordCategory = "authy-sms-outbound"
	USAGERECORDCATEGORY_CALL_PROGESS_EVENTS                                    UsageRecordCategory = "call-progess-events"
	USAGERECORDCATEGORY_CALLERIDLOOKUPS                                        UsageRecordCategory = "calleridlookups"
	USAGERECORDCATEGORY_CALLS                                                  UsageRecordCategory = "calls"
	USAGERECORDCATEGORY_CALLS_CLIENT                                           UsageRecordCategory = "calls-client"
	USAGERECORDCATEGORY_CALLS_GLOBALCONFERENCE                                 UsageRecordCategory = "calls-globalconference"
	USAGERECORDCATEGORY_CALLS_INBOUND                                          UsageRecordCategory = "calls-inbound"
	USAGERECORDCATEGORY_CALLS_INBOUND_LOCAL                                    UsageRecordCategory = "calls-inbound-local"
	USAGERECORDCATEGORY_CALLS_INBOUND_MOBILE                                   UsageRecordCategory = "calls-inbound-mobile"
	USAGERECORDCATEGORY_CALLS_INBOUND_TOLLFREE                                 UsageRecordCategory = "calls-inbound-tollfree"
	USAGERECORDCATEGORY_CALLS_OUTBOUND                                         UsageRecordCategory = "calls-outbound"
	USAGERECORDCATEGORY_CALLS_PAY_VERB_TRANSACTIONS                            UsageRecordCategory = "calls-pay-verb-transactions"
	USAGERECORDCATEGORY_CALLS_RECORDINGS                                       UsageRecordCategory = "calls-recordings"
	USAGERECORDCATEGORY_CALLS_SIP                                              UsageRecordCategory = "calls-sip"
	USAGERECORDCATEGORY_CALLS_SIP_INBOUND                                      UsageRecordCategory = "calls-sip-inbound"
	USAGERECORDCATEGORY_CALLS_SIP_OUTBOUND                                     UsageRecordCategory = "calls-sip-outbound"
	USAGERECORDCATEGORY_CARRIER_LOOKUPS                                        UsageRecordCategory = "carrier-lookups"
	USAGERECORDCATEGORY_CONVERSATIONS                                          UsageRecordCategory = "conversations"
	USAGERECORDCATEGORY_CONVERSATIONS_API_REQUESTS                             UsageRecordCategory = "conversations-api-requests"
	USAGERECORDCATEGORY_CONVERSATIONS_CONVERSATION_EVENTS                      UsageRecordCategory = "conversations-conversation-events"
	USAGERECORDCATEGORY_CONVERSATIONS_ENDPOINT_CONNECTIVITY                    UsageRecordCategory = "conversations-endpoint-connectivity"
	USAGERECORDCATEGORY_CONVERSATIONS_EVENTS                                   UsageRecordCategory = "conversations-events"
	USAGERECORDCATEGORY_CONVERSATIONS_PARTICIPANT_EVENTS                       UsageRecordCategory = "conversations-participant-events"
	USAGERECORDCATEGORY_CONVERSATIONS_PARTICIPANTS                             UsageRecordCategory = "conversations-participants"
	USAGERECORDCATEGORY_CPS                                                    UsageRecordCategory = "cps"
	USAGERECORDCATEGORY_FRAUD_LOOKUPS                                          UsageRecordCategory = "fraud-lookups"
	USAGERECORDCATEGORY_GROUP_ROOMS                                            UsageRecordCategory = "group-rooms"
	USAGERECORDCATEGORY_GROUP_ROOMS_DATA_TRACK                                 UsageRecordCategory = "group-rooms-data-track"
	USAGERECORDCATEGORY_GROUP_ROOMS_ENCRYPTED_MEDIA_RECORDED                   UsageRecordCategory = "group-rooms-encrypted-media-recorded"
	USAGERECORDCATEGORY_GROUP_ROOMS_MEDIA_DOWNLOADED                           UsageRecordCategory = "group-rooms-media-downloaded"
	USAGERECORDCATEGORY_GROUP_ROOMS_MEDIA_RECORDED                             UsageRecordCategory = "group-rooms-media-recorded"
	USAGERECORDCATEGORY_GROUP_ROOMS_MEDIA_ROUTED                               UsageRecordCategory = "group-rooms-media-routed"
	USAGERECORDCATEGORY_GROUP_ROOMS_MEDIA_STORED                               UsageRecordCategory = "group-rooms-media-stored"
	USAGERECORDCATEGORY_GROUP_ROOMS_PARTICIPANT_MINUTES                        UsageRecordCategory = "group-rooms-participant-minutes"
	USAGERECORDCATEGORY_GROUP_ROOMS_RECORDED_MINUTES                           UsageRecordCategory = "group-rooms-recorded-minutes"
	USAGERECORDCATEGORY_IMP_V1_USAGE                                           UsageRecordCategory = "imp-v1-usage"
	USAGERECORDCATEGORY_LOOKUPS                                                UsageRecordCategory = "lookups"
	USAGERECORDCATEGORY_MARKETPLACE                                            UsageRecordCategory = "marketplace"
	USAGERECORDCATEGORY_MARKETPLACE_ALGORITHMIA_NAMED_ENTITY_RECOGNITION       UsageRecordCategory = "marketplace-algorithmia-named-entity-recognition"
	USAGERECORDCATEGORY_MARKETPLACE_CADENCE_TRANSCRIPTION                      UsageRecordCategory = "marketplace-cadence-transcription"
	USAGERECORDCATEGORY_MARKETPLACE_CADENCE_TRANSLATION                        UsageRecordCategory = "marketplace-cadence-translation"
	USAGERECORDCATEGORY_MARKETPLACE_CAPIO_SPEECH_TO_TEXT                       UsageRecordCategory = "marketplace-capio-speech-to-text"
	USAGERECORDCATEGORY_MARKETPLACE_CONVRIZA_ABABA                             UsageRecordCategory = "marketplace-convriza-ababa"
	USAGERECORDCATEGORY_MARKETPLACE_DEEPGRAM_PHRASE_DETECTOR                   UsageRecordCategory = "marketplace-deepgram-phrase-detector"
	USAGERECORDCATEGORY_MARKETPLACE_DIGITAL_SEGMENT_BUSINESS_INFO              UsageRecordCategory = "marketplace-digital-segment-business-info"
	USAGERECORDCATEGORY_MARKETPLACE_FACEBOOK_OFFLINE_CONVERSIONS               UsageRecordCategory = "marketplace-facebook-offline-conversions"
	USAGERECORDCATEGORY_MARKETPLACE_GOOGLE_SPEECH_TO_TEXT                      UsageRecordCategory = "marketplace-google-speech-to-text"
	USAGERECORDCATEGORY_MARKETPLACE_IBM_WATSON_MESSAGE_INSIGHTS                UsageRecordCategory = "marketplace-ibm-watson-message-insights"
	USAGERECORDCATEGORY_MARKETPLACE_IBM_WATSON_MESSAGE_SENTIMENT               UsageRecordCategory = "marketplace-ibm-watson-message-sentiment"
	USAGERECORDCATEGORY_MARKETPLACE_IBM_WATSON_RECORDING_ANALYSIS              UsageRecordCategory = "marketplace-ibm-watson-recording-analysis"
	USAGERECORDCATEGORY_MARKETPLACE_IBM_WATSON_TONE_ANALYZER                   UsageRecordCategory = "marketplace-ibm-watson-tone-analyzer"
	USAGERECORDCATEGORY_MARKETPLACE_ICEHOOK_SYSTEMS_SCOUT                      UsageRecordCategory = "marketplace-icehook-systems-scout"
	USAGERECORDCATEGORY_MARKETPLACE_INFOGROUP_DATAAXLE_BIZINFO                 UsageRecordCategory = "marketplace-infogroup-dataaxle-bizinfo"
	USAGERECORDCATEGORY_MARKETPLACE_KEEN_IO_CONTACT_CENTER_ANALYTICS           UsageRecordCategory = "marketplace-keen-io-contact-center-analytics"
	USAGERECORDCATEGORY_MARKETPLACE_MARCHEX_CLEANCALL                          UsageRecordCategory = "marketplace-marchex-cleancall"
	USAGERECORDCATEGORY_MARKETPLACE_MARCHEX_SENTIMENT_ANALYSIS_FOR_SMS         UsageRecordCategory = "marketplace-marchex-sentiment-analysis-for-sms"
	USAGERECORDCATEGORY_MARKETPLACE_MARKETPLACE_NEXTCALLER_SOCIAL_ID           UsageRecordCategory = "marketplace-marketplace-nextcaller-social-id"
	USAGERECORDCATEGORY_MARKETPLACE_MOBILE_COMMONS_OPT_OUT_CLASSIFIER          UsageRecordCategory = "marketplace-mobile-commons-opt-out-classifier"
	USAGERECORDCATEGORY_MARKETPLACE_NEXIWAVE_VOICEMAIL_TO_TEXT                 UsageRecordCategory = "marketplace-nexiwave-voicemail-to-text"
	USAGERECORDCATEGORY_MARKETPLACE_NEXTCALLER_ADVANCED_CALLER_IDENTIFICATION  UsageRecordCategory = "marketplace-nextcaller-advanced-caller-identification"
	USAGERECORDCATEGORY_MARKETPLACE_NOMOROBO_SPAM_SCORE                        UsageRecordCategory = "marketplace-nomorobo-spam-score"
	USAGERECORDCATEGORY_MARKETPLACE_PAYFONE_TCPA_COMPLIANCE                    UsageRecordCategory = "marketplace-payfone-tcpa-compliance"
	USAGERECORDCATEGORY_MARKETPLACE_REMEETING_AUTOMATIC_SPEECH_RECOGNITION     UsageRecordCategory = "marketplace-remeeting-automatic-speech-recognition"
	USAGERECORDCATEGORY_MARKETPLACE_TCPA_DEFENSE_SOLUTIONS_BLACKLIST_FEED      UsageRecordCategory = "marketplace-tcpa-defense-solutions-blacklist-feed"
	USAGERECORDCATEGORY_MARKETPLACE_TELO_OPENCNAM                              UsageRecordCategory = "marketplace-telo-opencnam"
	USAGERECORDCATEGORY_MARKETPLACE_TRUECNAM_TRUE_SPAM                         UsageRecordCategory = "marketplace-truecnam-true-spam"
	USAGERECORDCATEGORY_MARKETPLACE_TWILIO_CALLER_NAME_LOOKUP_US               UsageRecordCategory = "marketplace-twilio-caller-name-lookup-us"
	USAGERECORDCATEGORY_MARKETPLACE_TWILIO_CARRIER_INFORMATION_LOOKUP          UsageRecordCategory = "marketplace-twilio-carrier-information-lookup"
	USAGERECORDCATEGORY_MARKETPLACE_VOICEBASE_PCI                              UsageRecordCategory = "marketplace-voicebase-pci"
	USAGERECORDCATEGORY_MARKETPLACE_VOICEBASE_TRANSCRIPTION                    UsageRecordCategory = "marketplace-voicebase-transcription"
	USAGERECORDCATEGORY_MARKETPLACE_VOICEBASE_TRANSCRIPTION_CUSTOM_VOCABULARY  UsageRecordCategory = "marketplace-voicebase-transcription-custom-vocabulary"
	USAGERECORDCATEGORY_MARKETPLACE_WHITEPAGES_PRO_CALLER_IDENTIFICATION       UsageRecordCategory = "marketplace-whitepages-pro-caller-identification"
	USAGERECORDCATEGORY_MARKETPLACE_WHITEPAGES_PRO_PHONE_INTELLIGENCE          UsageRecordCategory = "marketplace-whitepages-pro-phone-intelligence"
	USAGERECORDCATEGORY_MARKETPLACE_WHITEPAGES_PRO_PHONE_REPUTATION            UsageRecordCategory = "marketplace-whitepages-pro-phone-reputation"
	USAGERECORDCATEGORY_MARKETPLACE_WOLFARM_SPOKEN_RESULTS                     UsageRecordCategory = "marketplace-wolfarm-spoken-results"
	USAGERECORDCATEGORY_MARKETPLACE_WOLFRAM_SHORT_ANSWER                       UsageRecordCategory = "marketplace-wolfram-short-answer"
	USAGERECORDCATEGORY_MARKETPLACE_YTICA_CONTACT_CENTER_REPORTING_ANALYTICS   UsageRecordCategory = "marketplace-ytica-contact-center-reporting-analytics"
	USAGERECORDCATEGORY_MEDIASTORAGE                                           UsageRecordCategory = "mediastorage"
	USAGERECORDCATEGORY_MMS                                                    UsageRecordCategory = "mms"
	USAGERECORDCATEGORY_MMS_INBOUND                                            UsageRecordCategory = "mms-inbound"
	USAGERECORDCATEGORY_MMS_INBOUND_LONGCODE                                   UsageRecordCategory = "mms-inbound-longcode"
	USAGERECORDCATEGORY_MMS_INBOUND_SHORTCODE                                  UsageRecordCategory = "mms-inbound-shortcode"
	USAGERECORDCATEGORY_MMS_MESSAGES_CARRIERFEES                               UsageRecordCategory = "mms-messages-carrierfees"
	USAGERECORDCATEGORY_MMS_OUTBOUND                                           UsageRecordCategory = "mms-outbound"
	USAGERECORDCATEGORY_MMS_OUTBOUND_LONGCODE                                  UsageRecordCategory = "mms-outbound-longcode"
	USAGERECORDCATEGORY_MMS_OUTBOUND_SHORTCODE                                 UsageRecordCategory = "mms-outbound-shortcode"
	USAGERECORDCATEGORY_MONITOR_READS                                          UsageRecordCategory = "monitor-reads"
	USAGERECORDCATEGORY_MONITOR_STORAGE                                        UsageRecordCategory = "monitor-storage"
	USAGERECORDCATEGORY_MONITOR_WRITES                                         UsageRecordCategory = "monitor-writes"
	USAGERECORDCATEGORY_NOTIFY                                                 UsageRecordCategory = "notify"
	USAGERECORDCATEGORY_NOTIFY_ACTIONS_ATTEMPTS                                UsageRecordCategory = "notify-actions-attempts"
	USAGERECORDCATEGORY_NOTIFY_CHANNELS                                        UsageRecordCategory = "notify-channels"
	USAGERECORDCATEGORY_NUMBER_FORMAT_LOOKUPS                                  UsageRecordCategory = "number-format-lookups"
	USAGERECORDCATEGORY_PCHAT                                                  UsageRecordCategory = "pchat"
	USAGERECORDCATEGORY_PCHAT_USERS                                            UsageRecordCategory = "pchat-users"
	USAGERECORDCATEGORY_PEER_TO_PEER_ROOMS_PARTICIPANT_MINUTES                 UsageRecordCategory = "peer-to-peer-rooms-participant-minutes"
	USAGERECORDCATEGORY_PFAX                                                   UsageRecordCategory = "pfax"
	USAGERECORDCATEGORY_PFAX_MINUTES                                           UsageRecordCategory = "pfax-minutes"
	USAGERECORDCATEGORY_PFAX_MINUTES_INBOUND                                   UsageRecordCategory = "pfax-minutes-inbound"
	USAGERECORDCATEGORY_PFAX_MINUTES_OUTBOUND                                  UsageRecordCategory = "pfax-minutes-outbound"
	USAGERECORDCATEGORY_PFAX_PAGES                                             UsageRecordCategory = "pfax-pages"
	USAGERECORDCATEGORY_PHONENUMBERS                                           UsageRecordCategory = "phonenumbers"
	USAGERECORDCATEGORY_PHONENUMBERS_CPS                                       UsageRecordCategory = "phonenumbers-cps"
	USAGERECORDCATEGORY_PHONENUMBERS_EMERGENCY                                 UsageRecordCategory = "phonenumbers-emergency"
	USAGERECORDCATEGORY_PHONENUMBERS_LOCAL                                     UsageRecordCategory = "phonenumbers-local"
	USAGERECORDCATEGORY_PHONENUMBERS_MOBILE                                    UsageRecordCategory = "phonenumbers-mobile"
	USAGERECORDCATEGORY_PHONENUMBERS_SETUPS                                    UsageRecordCategory = "phonenumbers-setups"
	USAGERECORDCATEGORY_PHONENUMBERS_TOLLFREE                                  UsageRecordCategory = "phonenumbers-tollfree"
	USAGERECORDCATEGORY_PREMIUMSUPPORT                                         UsageRecordCategory = "premiumsupport"
	USAGERECORDCATEGORY_PROXY                                                  UsageRecordCategory = "proxy"
	USAGERECORDCATEGORY_PROXY_ACTIVE_SESSIONS                                  UsageRecordCategory = "proxy-active-sessions"
	USAGERECORDCATEGORY_PSTNCONNECTIVITY                                       UsageRecordCategory = "pstnconnectivity"
	USAGERECORDCATEGORY_PV                                                     UsageRecordCategory = "pv"
	USAGERECORDCATEGORY_PV_COMPOSITION_MEDIA_DOWNLOADED                        UsageRecordCategory = "pv-composition-media-downloaded"
	USAGERECORDCATEGORY_PV_COMPOSITION_MEDIA_ENCRYPTED                         UsageRecordCategory = "pv-composition-media-encrypted"
	USAGERECORDCATEGORY_PV_COMPOSITION_MEDIA_STORED                            UsageRecordCategory = "pv-composition-media-stored"
	USAGERECORDCATEGORY_PV_COMPOSITION_MINUTES                                 UsageRecordCategory = "pv-composition-minutes"
	USAGERECORDCATEGORY_PV_RECORDING_COMPOSITIONS                              UsageRecordCategory = "pv-recording-compositions"
	USAGERECORDCATEGORY_PV_ROOM_PARTICIPANTS                                   UsageRecordCategory = "pv-room-participants"
	USAGERECORDCATEGORY_PV_ROOM_PARTICIPANTS_AU1                               UsageRecordCategory = "pv-room-participants-au1"
	USAGERECORDCATEGORY_PV_ROOM_PARTICIPANTS_BR1                               UsageRecordCategory = "pv-room-participants-br1"
	USAGERECORDCATEGORY_PV_ROOM_PARTICIPANTS_IE1                               UsageRecordCategory = "pv-room-participants-ie1"
	USAGERECORDCATEGORY_PV_ROOM_PARTICIPANTS_JP1                               UsageRecordCategory = "pv-room-participants-jp1"
	USAGERECORDCATEGORY_PV_ROOM_PARTICIPANTS_SG1                               UsageRecordCategory = "pv-room-participants-sg1"
	USAGERECORDCATEGORY_PV_ROOM_PARTICIPANTS_US1                               UsageRecordCategory = "pv-room-participants-us1"
	USAGERECORDCATEGORY_PV_ROOM_PARTICIPANTS_US2                               UsageRecordCategory = "pv-room-participants-us2"
	USAGERECORDCATEGORY_PV_ROOMS                                               UsageRecordCategory = "pv-rooms"
	USAGERECORDCATEGORY_PV_SIP_ENDPOINT_REGISTRATIONS                          UsageRecordCategory = "pv-sip-endpoint-registrations"
	USAGERECORDCATEGORY_RECORDINGS                                             UsageRecordCategory = "recordings"
	USAGERECORDCATEGORY_RECORDINGSTORAGE                                       UsageRecordCategory = "recordingstorage"
	USAGERECORDCATEGORY_ROOMS_GROUP_BANDWIDTH                                  UsageRecordCategory = "rooms-group-bandwidth"
	USAGERECORDCATEGORY_ROOMS_GROUP_MINUTES                                    UsageRecordCategory = "rooms-group-minutes"
	USAGERECORDCATEGORY_ROOMS_PEER_TO_PEER_MINUTES                             UsageRecordCategory = "rooms-peer-to-peer-minutes"
	USAGERECORDCATEGORY_SHORTCODES                                             UsageRecordCategory = "shortcodes"
	USAGERECORDCATEGORY_SHORTCODES_CUSTOMEROWNED                               UsageRecordCategory = "shortcodes-customerowned"
	USAGERECORDCATEGORY_SHORTCODES_MMS_ENABLEMENT                              UsageRecordCategory = "shortcodes-mms-enablement"
	USAGERECORDCATEGORY_SHORTCODES_MPS                                         UsageRecordCategory = "shortcodes-mps"
	USAGERECORDCATEGORY_SHORTCODES_RANDOM                                      UsageRecordCategory = "shortcodes-random"
	USAGERECORDCATEGORY_SHORTCODES_UK                                          UsageRecordCategory = "shortcodes-uk"
	USAGERECORDCATEGORY_SHORTCODES_VANITY                                      UsageRecordCategory = "shortcodes-vanity"
	USAGERECORDCATEGORY_SMALL_GROUP_ROOMS                                      UsageRecordCategory = "small-group-rooms"
	USAGERECORDCATEGORY_SMALL_GROUP_ROOMS_DATA_TRACK                           UsageRecordCategory = "small-group-rooms-data-track"
	USAGERECORDCATEGORY_SMALL_GROUP_ROOMS_PARTICIPANT_MINUTES                  UsageRecordCategory = "small-group-rooms-participant-minutes"
	USAGERECORDCATEGORY_SMS                                                    UsageRecordCategory = "sms"
	USAGERECORDCATEGORY_SMS_INBOUND                                            UsageRecordCategory = "sms-inbound"
	USAGERECORDCATEGORY_SMS_INBOUND_LONGCODE                                   UsageRecordCategory = "sms-inbound-longcode"
	USAGERECORDCATEGORY_SMS_INBOUND_SHORTCODE                                  UsageRecordCategory = "sms-inbound-shortcode"
	USAGERECORDCATEGORY_SMS_MESSAGES_CARRIERFEES                               UsageRecordCategory = "sms-messages-carrierfees"
	USAGERECORDCATEGORY_SMS_MESSAGES_FEATURES                                  UsageRecordCategory = "sms-messages-features"
	USAGERECORDCATEGORY_SMS_MESSAGES_FEATURES_SENDERID                         UsageRecordCategory = "sms-messages-features-senderid"
	USAGERECORDCATEGORY_SMS_OUTBOUND                                           UsageRecordCategory = "sms-outbound"
	USAGERECORDCATEGORY_SMS_OUTBOUND_CONTENT_INSPECTION                        UsageRecordCategory = "sms-outbound-content-inspection"
	USAGERECORDCATEGORY_SMS_OUTBOUND_LONGCODE                                  UsageRecordCategory = "sms-outbound-longcode"
	USAGERECORDCATEGORY_SMS_OUTBOUND_SHORTCODE                                 UsageRecordCategory = "sms-outbound-shortcode"
	USAGERECORDCATEGORY_SPEECH_RECOGNITION                                     UsageRecordCategory = "speech-recognition"
	USAGERECORDCATEGORY_STUDIO_ENGAGEMENTS                                     UsageRecordCategory = "studio-engagements"
	USAGERECORDCATEGORY_SYNC                                                   UsageRecordCategory = "sync"
	USAGERECORDCATEGORY_SYNC_ACTIONS                                           UsageRecordCategory = "sync-actions"
	USAGERECORDCATEGORY_SYNC_ENDPOINT_HOURS                                    UsageRecordCategory = "sync-endpoint-hours"
	USAGERECORDCATEGORY_SYNC_ENDPOINT_HOURS_ABOVE_DAILY_CAP                    UsageRecordCategory = "sync-endpoint-hours-above-daily-cap"
	USAGERECORDCATEGORY_TASKROUTER_TASKS                                       UsageRecordCategory = "taskrouter-tasks"
	USAGERECORDCATEGORY_TOTALPRICE                                             UsageRecordCategory = "totalprice"
	USAGERECORDCATEGORY_TRANSCRIPTIONS                                         UsageRecordCategory = "transcriptions"
	USAGERECORDCATEGORY_TRUNKING_CPS                                           UsageRecordCategory = "trunking-cps"
	USAGERECORDCATEGORY_TRUNKING_EMERGENCY_CALLS                               UsageRecordCategory = "trunking-emergency-calls"
	USAGERECORDCATEGORY_TRUNKING_ORIGINATION                                   UsageRecordCategory = "trunking-origination"
	USAGERECORDCATEGORY_TRUNKING_ORIGINATION_LOCAL                             UsageRecordCategory = "trunking-origination-local"
	USAGERECORDCATEGORY_TRUNKING_ORIGINATION_MOBILE                            UsageRecordCategory = "trunking-origination-mobile"
	USAGERECORDCATEGORY_TRUNKING_ORIGINATION_TOLLFREE                          UsageRecordCategory = "trunking-origination-tollfree"
	USAGERECORDCATEGORY_TRUNKING_RECORDINGS                                    UsageRecordCategory = "trunking-recordings"
	USAGERECORDCATEGORY_TRUNKING_SECURE                                        UsageRecordCategory = "trunking-secure"
	USAGERECORDCATEGORY_TRUNKING_TERMINATION                                   UsageRecordCategory = "trunking-termination"
	USAGERECORDCATEGORY_TURNMEGABYTES                                          UsageRecordCategory = "turnmegabytes"
	USAGERECORDCATEGORY_TURNMEGABYTES_AUSTRALIA                                UsageRecordCategory = "turnmegabytes-australia"
	USAGERECORDCATEGORY_TURNMEGABYTES_BRASIL                                   UsageRecordCategory = "turnmegabytes-brasil"
	USAGERECORDCATEGORY_TURNMEGABYTES_GERMANY                                  UsageRecordCategory = "turnmegabytes-germany"
	USAGERECORDCATEGORY_TURNMEGABYTES_INDIA                                    UsageRecordCategory = "turnmegabytes-india"
	USAGERECORDCATEGORY_TURNMEGABYTES_IRELAND                                  UsageRecordCategory = "turnmegabytes-ireland"
	USAGERECORDCATEGORY_TURNMEGABYTES_JAPAN                                    UsageRecordCategory = "turnmegabytes-japan"
	USAGERECORDCATEGORY_TURNMEGABYTES_SINGAPORE                                UsageRecordCategory = "turnmegabytes-singapore"
	USAGERECORDCATEGORY_TURNMEGABYTES_USEAST                                   UsageRecordCategory = "turnmegabytes-useast"
	USAGERECORDCATEGORY_TURNMEGABYTES_USWEST                                   UsageRecordCategory = "turnmegabytes-uswest"
	USAGERECORDCATEGORY_TWILIO_INTERCONNECT                                    UsageRecordCategory = "twilio-interconnect"
	USAGERECORDCATEGORY_VERIFY_PUSH                                            UsageRecordCategory = "verify-push"
	USAGERECORDCATEGORY_VIDEO_RECORDINGS                                       UsageRecordCategory = "video-recordings"
	USAGERECORDCATEGORY_VOICE_INSIGHTS                                         UsageRecordCategory = "voice-insights"
	USAGERECORDCATEGORY_VOICE_INSIGHTS_CLIENT_INSIGHTS_ON_DEMAND_MINUTE        UsageRecordCategory = "voice-insights-client-insights-on-demand-minute"
	USAGERECORDCATEGORY_VOICE_INSIGHTS_PTSN_INSIGHTS_ON_DEMAND_MINUTE          UsageRecordCategory = "voice-insights-ptsn-insights-on-demand-minute"
	USAGERECORDCATEGORY_VOICE_INSIGHTS_SIP_INTERFACE_INSIGHTS_ON_DEMAND_MINUTE UsageRecordCategory = "voice-insights-sip-interface-insights-on-demand-minute"
	USAGERECORDCATEGORY_VOICE_INSIGHTS_SIP_TRUNKING_INSIGHTS_ON_DEMAND_MINUTE  UsageRecordCategory = "voice-insights-sip-trunking-insights-on-demand-minute"
	USAGERECORDCATEGORY_WIRELESS                                               UsageRecordCategory = "wireless"
	USAGERECORDCATEGORY_WIRELESS_ORDERS                                        UsageRecordCategory = "wireless-orders"
	USAGERECORDCATEGORY_WIRELESS_ORDERS_ARTWORK                                UsageRecordCategory = "wireless-orders-artwork"
	USAGERECORDCATEGORY_WIRELESS_ORDERS_BULK                                   UsageRecordCategory = "wireless-orders-bulk"
	USAGERECORDCATEGORY_WIRELESS_ORDERS_ESIM                                   UsageRecordCategory = "wireless-orders-esim"
	USAGERECORDCATEGORY_WIRELESS_ORDERS_STARTER                                UsageRecordCategory = "wireless-orders-starter"
	USAGERECORDCATEGORY_WIRELESS_USAGE                                         UsageRecordCategory = "wireless-usage"
	USAGERECORDCATEGORY_WIRELESS_USAGE_COMMANDS                                UsageRecordCategory = "wireless-usage-commands"
	USAGERECORDCATEGORY_WIRELESS_USAGE_COMMANDS_AFRICA                         UsageRecordCategory = "wireless-usage-commands-africa"
	USAGERECORDCATEGORY_WIRELESS_USAGE_COMMANDS_ASIA                           UsageRecordCategory = "wireless-usage-commands-asia"
	USAGERECORDCATEGORY_WIRELESS_USAGE_COMMANDS_CENTRALANDSOUTHAMERICA         UsageRecordCategory = "wireless-usage-commands-centralandsouthamerica"
	USAGERECORDCATEGORY_WIRELESS_USAGE_COMMANDS_EUROPE                         UsageRecordCategory = "wireless-usage-commands-europe"
	USAGERECORDCATEGORY_WIRELESS_USAGE_COMMANDS_HOME                           UsageRecordCategory = "wireless-usage-commands-home"
	USAGERECORDCATEGORY_WIRELESS_USAGE_COMMANDS_NORTHAMERICA                   UsageRecordCategory = "wireless-usage-commands-northamerica"
	USAGERECORDCATEGORY_WIRELESS_USAGE_COMMANDS_OCEANIA                        UsageRecordCategory = "wireless-usage-commands-oceania"
	USAGERECORDCATEGORY_WIRELESS_USAGE_COMMANDS_ROAMING                        UsageRecordCategory = "wireless-usage-commands-roaming"
	USAGERECORDCATEGORY_WIRELESS_USAGE_DATA                                    UsageRecordCategory = "wireless-usage-data"
	USAGERECORDCATEGORY_WIRELESS_USAGE_DATA_AFRICA                             UsageRecordCategory = "wireless-usage-data-africa"
	USAGERECORDCATEGORY_WIRELESS_USAGE_DATA_ASIA                               UsageRecordCategory = "wireless-usage-data-asia"
	USAGERECORDCATEGORY_WIRELESS_USAGE_DATA_CENTRALANDSOUTHAMERICA             UsageRecordCategory = "wireless-usage-data-centralandsouthamerica"
	USAGERECORDCATEGORY_WIRELESS_USAGE_DATA_CUSTOM_ADDITIONALMB                UsageRecordCategory = "wireless-usage-data-custom-additionalmb"
	USAGERECORDCATEGORY_WIRELESS_USAGE_DATA_CUSTOM_FIRST5MB                    UsageRecordCategory = "wireless-usage-data-custom-first5mb"
	USAGERECORDCATEGORY_WIRELESS_USAGE_DATA_DOMESTIC_ROAMING                   UsageRecordCategory = "wireless-usage-data-domestic-roaming"
	USAGERECORDCATEGORY_WIRELESS_USAGE_DATA_EUROPE                             UsageRecordCategory = "wireless-usage-data-europe"
	USAGERECORDCATEGORY_WIRELESS_USAGE_DATA_INDIVIDUAL_ADDITIONALGB            UsageRecordCategory = "wireless-usage-data-individual-additionalgb"
	USAGERECORDCATEGORY_WIRELESS_USAGE_DATA_INDIVIDUAL_FIRSTGB                 UsageRecordCategory = "wireless-usage-data-individual-firstgb"
	USAGERECORDCATEGORY_WIRELESS_USAGE_DATA_INTERNATIONAL_ROAMING_CANADA       UsageRecordCategory = "wireless-usage-data-international-roaming-canada"
	USAGERECORDCATEGORY_WIRELESS_USAGE_DATA_INTERNATIONAL_ROAMING_INDIA        UsageRecordCategory = "wireless-usage-data-international-roaming-india"
	USAGERECORDCATEGORY_WIRELESS_USAGE_DATA_INTERNATIONAL_ROAMING_MEXICO       UsageRecordCategory = "wireless-usage-data-international-roaming-mexico"
	USAGERECORDCATEGORY_WIRELESS_USAGE_DATA_NORTHAMERICA                       UsageRecordCategory = "wireless-usage-data-northamerica"
	USAGERECORDCATEGORY_WIRELESS_USAGE_DATA_OCEANIA                            UsageRecordCategory = "wireless-usage-data-oceania"
	USAGERECORDCATEGORY_WIRELESS_USAGE_DATA_POOLED                             UsageRecordCategory = "wireless-usage-data-pooled"
	USAGERECORDCATEGORY_WIRELESS_USAGE_DATA_POOLED_DOWNLINK                    UsageRecordCategory = "wireless-usage-data-pooled-downlink"
	USAGERECORDCATEGORY_WIRELESS_USAGE_DATA_POOLED_UPLINK                      UsageRecordCategory = "wireless-usage-data-pooled-uplink"
	USAGERECORDCATEGORY_WIRELESS_USAGE_MRC                                     UsageRecordCategory = "wireless-usage-mrc"
	USAGERECORDCATEGORY_WIRELESS_USAGE_MRC_CUSTOM                              UsageRecordCategory = "wireless-usage-mrc-custom"
	USAGERECORDCATEGORY_WIRELESS_USAGE_MRC_INDIVIDUAL                          UsageRecordCategory = "wireless-usage-mrc-individual"
	USAGERECORDCATEGORY_WIRELESS_USAGE_MRC_POOLED                              UsageRecordCategory = "wireless-usage-mrc-pooled"
	USAGERECORDCATEGORY_WIRELESS_USAGE_MRC_SUSPENDED                           UsageRecordCategory = "wireless-usage-mrc-suspended"
	USAGERECORDCATEGORY_WIRELESS_USAGE_SMS                                     UsageRecordCategory = "wireless-usage-sms"
	USAGERECORDCATEGORY_WIRELESS_USAGE_VOICE                                   UsageRecordCategory = "wireless-usage-voice"
)
