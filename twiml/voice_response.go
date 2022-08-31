package twiml

func Voice(verbs []Verb)(TwimlXml string){
    doc, response := CreateDocument()
    if verbs != nil {
        AddAllVerbs(response, verbs)
    }
    return ToXML(doc)
}

//VoiceReferVerb <Refer> TwiML Verb
type VoiceReferVerb struct {
    // action: Action URL
    // method: Action URL method
    // OptionalAttributes: additional attributes
    Action string
    Method string
    Nouns []Noun
    OptionalAttributes map[string]string
}

func (m VoiceReferVerb) GetName() string {
    return "Refer"
}

func (m VoiceReferVerb) GetText() string {
    return ""
}

func (m VoiceReferVerb) GetAttr() (map[string]string, map[string]string) {
    paramsAttr := map[string]string{
    	"Action": m.Action,
    	"Method": m.Method,
    }
    return m.OptionalAttributes, paramsAttr
}

func (m VoiceReferVerb) GetNouns() []Noun {
    return m.Nouns
}
//VoiceReferSipNoun <Sip> TwiML Noun used in <Refer>
type VoiceReferSipNoun struct {
    // sip_url: SIP URL
    // OptionalAttributes: additional attributes
    SipUrl string
    OptionalAttributes map[string]string
}

func (m VoiceReferSipNoun) GetName() string {
    return "Sip"
}

func (m VoiceReferSipNoun) GetText() string {
    return m.SipUrl
}

func (m VoiceReferSipNoun) GetAttr() (map[string]string, map[string]string) {
    return m.OptionalAttributes, nil
}
//VoiceStopVerb <Stop> TwiML Verb
type VoiceStopVerb struct {
    // OptionalAttributes: additional attributes
    Nouns []Noun
    OptionalAttributes map[string]string
}

func (m VoiceStopVerb) GetName() string {
    return "Stop"
}

func (m VoiceStopVerb) GetText() string {
    return ""
}

func (m VoiceStopVerb) GetAttr() (map[string]string, map[string]string) {
    return m.OptionalAttributes, nil
}

func (m VoiceStopVerb) GetNouns() []Noun {
    return m.Nouns
}
//VoiceSipRecNoun <Siprec> TwiML Noun
type VoiceSipRecNoun struct {
    // name: Friendly name given to SIPREC
    // connector_name: Unique name for Connector
    // track: Track to be streamed to remote service
    // OptionalAttributes: additional attributes
    Name string
    ConnectorName string
    Track string
    OptionalAttributes map[string]string
}

func (m VoiceSipRecNoun) GetName() string {
    return "SipRec"
}

func (m VoiceSipRecNoun) GetText() string {
    return ""
}

func (m VoiceSipRecNoun) GetAttr() (map[string]string, map[string]string) {
    paramsAttr := map[string]string{
    	"Name": m.Name,
    	"ConnectorName": m.ConnectorName,
    	"Track": m.Track,
    }
    return m.OptionalAttributes, paramsAttr
}
//VoiceParameterNoun <Parameter> TwiML Noun
type VoiceParameterNoun struct {
    // name: The name of the custom parameter
    // value: The value of the custom parameter
    // OptionalAttributes: additional attributes
    Name string
    Value string
    OptionalAttributes map[string]string
}

func (m VoiceParameterNoun) GetName() string {
    return "Parameter"
}

func (m VoiceParameterNoun) GetText() string {
    return ""
}

func (m VoiceParameterNoun) GetAttr() (map[string]string, map[string]string) {
    paramsAttr := map[string]string{
    	"Name": m.Name,
    	"Value": m.Value,
    }
    return m.OptionalAttributes, paramsAttr
}
//VoiceStreamNoun <Stream> TwiML Noun
type VoiceStreamNoun struct {
    // name: Friendly name given to the Stream
    // connector_name: Unique name for Stream Connector
    // url: URL of the remote service where the Stream is routed
    // track: Track to be streamed to remote service
    // status_callback: Status Callback URL
    // status_callback_method: Status Callback URL method
    // OptionalAttributes: additional attributes
    Name string
    ConnectorName string
    Url string
    Track string
    StatusCallback string
    StatusCallbackMethod string
    OptionalAttributes map[string]string
}

func (m VoiceStreamNoun) GetName() string {
    return "Stream"
}

func (m VoiceStreamNoun) GetText() string {
    return ""
}

func (m VoiceStreamNoun) GetAttr() (map[string]string, map[string]string) {
    paramsAttr := map[string]string{
    	"Name": m.Name,
    	"ConnectorName": m.ConnectorName,
    	"Url": m.Url,
    	"Track": m.Track,
    	"StatusCallback": m.StatusCallback,
    	"StatusCallbackMethod": m.StatusCallbackMethod,
    }
    return m.OptionalAttributes, paramsAttr
}
//VoiceStartVerb <Start> TwiML Verb
type VoiceStartVerb struct {
    // action: Action URL
    // method: Action URL method
    // OptionalAttributes: additional attributes
    Action string
    Method string
    Nouns []Noun
    OptionalAttributes map[string]string
}

func (m VoiceStartVerb) GetName() string {
    return "Start"
}

func (m VoiceStartVerb) GetText() string {
    return ""
}

func (m VoiceStartVerb) GetAttr() (map[string]string, map[string]string) {
    paramsAttr := map[string]string{
    	"Action": m.Action,
    	"Method": m.Method,
    }
    return m.OptionalAttributes, paramsAttr
}

func (m VoiceStartVerb) GetNouns() []Noun {
    return m.Nouns
}
//VoicePromptVerb <Prompt> Twiml Verb
type VoicePromptVerb struct {
    // for_: Name of the payment source data element
    // error_type: Type of error
    // card_type: Type of the credit card
    // attempt: Current attempt count
    // OptionalAttributes: additional attributes
    For_ string
    ErrorType string
    CardType string
    Attempt string
    Nouns []Noun
    OptionalAttributes map[string]string
}

func (m VoicePromptVerb) GetName() string {
    return "Prompt"
}

func (m VoicePromptVerb) GetText() string {
    return ""
}

func (m VoicePromptVerb) GetAttr() (map[string]string, map[string]string) {
    paramsAttr := map[string]string{
    	"For_": m.For_,
    	"ErrorType": m.ErrorType,
    	"CardType": m.CardType,
    	"Attempt": m.Attempt,
    }
    return m.OptionalAttributes, paramsAttr
}

func (m VoicePromptVerb) GetNouns() []Noun {
    return m.Nouns
}
//VoicePauseVerb <Pause> TwiML Verb
type VoicePauseVerb struct {
    // length: Length in seconds to pause
    // OptionalAttributes: additional attributes
    Length string
    Nouns []Noun
    OptionalAttributes map[string]string
}

func (m VoicePauseVerb) GetName() string {
    return "Pause"
}

func (m VoicePauseVerb) GetText() string {
    return ""
}

func (m VoicePauseVerb) GetAttr() (map[string]string, map[string]string) {
    paramsAttr := map[string]string{
    	"Length": m.Length,
    }
    return m.OptionalAttributes, paramsAttr
}

func (m VoicePauseVerb) GetNouns() []Noun {
    return m.Nouns
}
//VoicePlayVerb <Play> TwiML Verb
type VoicePlayVerb struct {
    // url: Media URL
    // loop: Times to loop media
    // digits: Play DTMF tones for digits
    // OptionalAttributes: additional attributes
    Url string
    Loop string
    Digits string
    Nouns []Noun
    OptionalAttributes map[string]string
}

func (m VoicePlayVerb) GetName() string {
    return "Play"
}

func (m VoicePlayVerb) GetText() string {
    return m.Url
}

func (m VoicePlayVerb) GetAttr() (map[string]string, map[string]string) {
    paramsAttr := map[string]string{
    	"Loop": m.Loop,
    	"Digits": m.Digits,
    }
    return m.OptionalAttributes, paramsAttr
}

func (m VoicePlayVerb) GetNouns() []Noun {
    return m.Nouns
}
//VoiceSayVerb <Say> TwiML Verb
type VoiceSayVerb struct {
    // message: Message to say
    // voice: Voice to use
    // loop: Times to loop message
    // language: Message language
    // OptionalAttributes: additional attributes
    Message string
    Voice string
    Loop string
    Language string
    Nouns []Noun
    OptionalAttributes map[string]string
}

func (m VoiceSayVerb) GetName() string {
    return "Say"
}

func (m VoiceSayVerb) GetText() string {
    return m.Message
}

func (m VoiceSayVerb) GetAttr() (map[string]string, map[string]string) {
    paramsAttr := map[string]string{
    	"Voice": m.Voice,
    	"Loop": m.Loop,
    	"Language": m.Language,
    }
    return m.OptionalAttributes, paramsAttr
}

func (m VoiceSayVerb) GetNouns() []Noun {
    return m.Nouns
}
//VoiceWNoun Improving Pronunciation by Specifying Parts of Speech in <Say>
type VoiceWNoun struct {
    // words: Words to speak
    // role: Customize the pronunciation of words by specifying the word’s part of speech or alternate meaning
    // OptionalAttributes: additional attributes
    Words string
    Role string
    OptionalAttributes map[string]string
}

func (m VoiceWNoun) GetName() string {
    return "W"
}

func (m VoiceWNoun) GetText() string {
    return m.Words
}

func (m VoiceWNoun) GetAttr() (map[string]string, map[string]string) {
    paramsAttr := map[string]string{
    	"Role": m.Role,
    }
    return m.OptionalAttributes, paramsAttr
}
//VoiceSubNoun Pronouncing Acronyms and Abbreviations in <Say>
type VoiceSubNoun struct {
    // words: Words to be substituted
    // alias: Substitute a different word (or pronunciation) for selected text such as an acronym or abbreviation
    // OptionalAttributes: additional attributes
    Words string
    Alias string
    OptionalAttributes map[string]string
}

func (m VoiceSubNoun) GetName() string {
    return "Sub"
}

func (m VoiceSubNoun) GetText() string {
    return m.Words
}

func (m VoiceSubNoun) GetAttr() (map[string]string, map[string]string) {
    paramsAttr := map[string]string{
    	"Alias": m.Alias,
    }
    return m.OptionalAttributes, paramsAttr
}
//VoiceSayAsNoun Controlling How Special Types of Words Are Spoken in <Say>
type VoiceSayAsNoun struct {
    // words: Words to be interpreted
    // interpret-as: Specify the type of words are spoken
    // role: Specify the format of the date when interpret-as is set to date
    // OptionalAttributes: additional attributes
    Words string
    InterpretAs string
    Role string
    OptionalAttributes map[string]string
}

func (m VoiceSayAsNoun) GetName() string {
    return "SayAs"
}

func (m VoiceSayAsNoun) GetText() string {
    return m.Words
}

func (m VoiceSayAsNoun) GetAttr() (map[string]string, map[string]string) {
    paramsAttr := map[string]string{
    	"InterpretAs": m.InterpretAs,
    	"Role": m.Role,
    }
    return m.OptionalAttributes, paramsAttr
}
//VoiceProsodyNoun Controling Volume, Speaking Rate, and Pitch in <Say>
type VoiceProsodyNoun struct {
    // words: Words to speak
    // volume: Specify the volume, available values: default, silent, x-soft, soft, medium, loud, x-loud, +ndB, -ndB
    // rate: Specify the rate, available values: x-slow, slow, medium, fast, x-fast, n%
    // pitch: Specify the pitch, available values: default, x-low, low, medium, high, x-high, +n%, -n%
    // OptionalAttributes: additional attributes
    Words string
    Volume string
    Rate string
    Pitch string
    OptionalAttributes map[string]string
}

func (m VoiceProsodyNoun) GetName() string {
    return "Prosody"
}

func (m VoiceProsodyNoun) GetText() string {
    return m.Words
}

func (m VoiceProsodyNoun) GetAttr() (map[string]string, map[string]string) {
    paramsAttr := map[string]string{
    	"Volume": m.Volume,
    	"Rate": m.Rate,
    	"Pitch": m.Pitch,
    }
    return m.OptionalAttributes, paramsAttr
}
//VoiceSNoun Adding A Pause Between Sentences in <Say>
type VoiceSNoun struct {
    // words: Words to speak
    // OptionalAttributes: additional attributes
    Words string
    OptionalAttributes map[string]string
}

func (m VoiceSNoun) GetName() string {
    return "S"
}

func (m VoiceSNoun) GetText() string {
    return m.Words
}

func (m VoiceSNoun) GetAttr() (map[string]string, map[string]string) {
    return m.OptionalAttributes, nil
}
//VoicePhonemeNoun Using Phonetic Pronunciation in <Say>
type VoicePhonemeNoun struct {
    // words: Words to speak
    // alphabet: Specify the phonetic alphabet
    // ph: Specifiy the phonetic symbols for pronunciation
    // OptionalAttributes: additional attributes
    Words string
    Alphabet string
    Ph string
    OptionalAttributes map[string]string
}

func (m VoicePhonemeNoun) GetName() string {
    return "Phoneme"
}

func (m VoicePhonemeNoun) GetText() string {
    return m.Words
}

func (m VoicePhonemeNoun) GetAttr() (map[string]string, map[string]string) {
    paramsAttr := map[string]string{
    	"Alphabet": m.Alphabet,
    	"Ph": m.Ph,
    }
    return m.OptionalAttributes, paramsAttr
}
//VoiceLangNoun Specifying Another Language for Specific Words in <Say>
type VoiceLangNoun struct {
    // words: Words to speak
    // xml:lang: Specify the language
    // OptionalAttributes: additional attributes
    Words string
    XmlLang string
    OptionalAttributes map[string]string
}

func (m VoiceLangNoun) GetName() string {
    return "Lang"
}

func (m VoiceLangNoun) GetText() string {
    return m.Words
}

func (m VoiceLangNoun) GetAttr() (map[string]string, map[string]string) {
    paramsAttr := map[string]string{
    	"XmlLang": m.XmlLang,
    }
    return m.OptionalAttributes, paramsAttr
}
//VoicePNoun Adding a Pause Between Paragraphs in <Say>
type VoicePNoun struct {
    // words: Words to speak
    // OptionalAttributes: additional attributes
    Words string
    OptionalAttributes map[string]string
}

func (m VoicePNoun) GetName() string {
    return "P"
}

func (m VoicePNoun) GetText() string {
    return m.Words
}

func (m VoicePNoun) GetAttr() (map[string]string, map[string]string) {
    return m.OptionalAttributes, nil
}
//VoiceEmphasisNoun Emphasizing Words in <Say>
type VoiceEmphasisNoun struct {
    // words: Words to emphasize
    // level: Specify the degree of emphasis
    // OptionalAttributes: additional attributes
    Words string
    Level string
    OptionalAttributes map[string]string
}

func (m VoiceEmphasisNoun) GetName() string {
    return "Emphasis"
}

func (m VoiceEmphasisNoun) GetText() string {
    return m.Words
}

func (m VoiceEmphasisNoun) GetAttr() (map[string]string, map[string]string) {
    paramsAttr := map[string]string{
    	"Level": m.Level,
    }
    return m.OptionalAttributes, paramsAttr
}
//VoiceBreakNoun Adding a Pause in <Say>
type VoiceBreakNoun struct {
    // strength: Set a pause based on strength
    // time: Set a pause to a specific length of time in seconds or milliseconds, available values: [number]s, [number]ms
    // OptionalAttributes: additional attributes
    Strength string
    Time string
    OptionalAttributes map[string]string
}

func (m VoiceBreakNoun) GetName() string {
    return "Break"
}

func (m VoiceBreakNoun) GetText() string {
    return ""
}

func (m VoiceBreakNoun) GetAttr() (map[string]string, map[string]string) {
    paramsAttr := map[string]string{
    	"Strength": m.Strength,
    	"Time": m.Time,
    }
    return m.OptionalAttributes, paramsAttr
}
//VoicePayVerb <Pay> Twiml Verb
type VoicePayVerb struct {
    // input: Input type Twilio should accept
    // action: Action URL
    // bank_account_type: Bank account type for ach transactions. If set, payment method attribute must be provided and value should be set to ach-debit. defaults to consumer-checking
    // status_callback: Status callback URL
    // status_callback_method: Status callback method
    // timeout: Time to wait to gather input
    // max_attempts: Maximum number of allowed retries when gathering input
    // security_code: Prompt for security code
    // postal_code: Prompt for postal code and it should be true/false or default postal code
    // min_postal_code_length: Prompt for minimum postal code length
    // payment_connector: Unique name for payment connector
    // payment_method: Payment method to be used. defaults to credit-card
    // token_type: Type of token
    // charge_amount: Amount to process. If value is greater than 0 then make the payment else create a payment token
    // currency: Currency of the amount attribute
    // description: Details regarding the payment
    // valid_card_types: Comma separated accepted card types
    // language: Language to use
    // OptionalAttributes: additional attributes
    Input string
    Action string
    BankAccountType string
    StatusCallback string
    StatusCallbackMethod string
    Timeout string
    MaxAttempts string
    SecurityCode string
    PostalCode string
    MinPostalCodeLength string
    PaymentConnector string
    PaymentMethod string
    TokenType string
    ChargeAmount string
    Currency string
    Description string
    ValidCardTypes string
    Language string
    Nouns []Noun
    OptionalAttributes map[string]string
}

func (m VoicePayVerb) GetName() string {
    return "Pay"
}

func (m VoicePayVerb) GetText() string {
    return ""
}

func (m VoicePayVerb) GetAttr() (map[string]string, map[string]string) {
    paramsAttr := map[string]string{
    	"Input": m.Input,
    	"Action": m.Action,
    	"BankAccountType": m.BankAccountType,
    	"StatusCallback": m.StatusCallback,
    	"StatusCallbackMethod": m.StatusCallbackMethod,
    	"Timeout": m.Timeout,
    	"MaxAttempts": m.MaxAttempts,
    	"SecurityCode": m.SecurityCode,
    	"PostalCode": m.PostalCode,
    	"MinPostalCodeLength": m.MinPostalCodeLength,
    	"PaymentConnector": m.PaymentConnector,
    	"PaymentMethod": m.PaymentMethod,
    	"TokenType": m.TokenType,
    	"ChargeAmount": m.ChargeAmount,
    	"Currency": m.Currency,
    	"Description": m.Description,
    	"ValidCardTypes": m.ValidCardTypes,
    	"Language": m.Language,
    }
    return m.OptionalAttributes, paramsAttr
}

func (m VoicePayVerb) GetNouns() []Noun {
    return m.Nouns
}
//VoiceSmsNoun <Sms> TwiML Noun
type VoiceSmsNoun struct {
    // message: Message body
    // to: Number to send message to
    // from: Number to send message from
    // action: Action URL
    // method: Action URL method
    // status_callback: Status callback URL
    // OptionalAttributes: additional attributes
    Message string
    To string
    From string
    Action string
    Method string
    StatusCallback string
    OptionalAttributes map[string]string
}

func (m VoiceSmsNoun) GetName() string {
    return "Sms"
}

func (m VoiceSmsNoun) GetText() string {
    return m.Message
}

func (m VoiceSmsNoun) GetAttr() (map[string]string, map[string]string) {
    paramsAttr := map[string]string{
    	"To": m.To,
    	"From": m.From,
    	"Action": m.Action,
    	"Method": m.Method,
    	"StatusCallback": m.StatusCallback,
    }
    return m.OptionalAttributes, paramsAttr
}
//VoiceRejectVerb <Reject> TwiML Verb
type VoiceRejectVerb struct {
    // reason: Rejection reason
    // OptionalAttributes: additional attributes
    Reason string
    Nouns []Noun
    OptionalAttributes map[string]string
}

func (m VoiceRejectVerb) GetName() string {
    return "Reject"
}

func (m VoiceRejectVerb) GetText() string {
    return ""
}

func (m VoiceRejectVerb) GetAttr() (map[string]string, map[string]string) {
    paramsAttr := map[string]string{
    	"Reason": m.Reason,
    }
    return m.OptionalAttributes, paramsAttr
}

func (m VoiceRejectVerb) GetNouns() []Noun {
    return m.Nouns
}
//VoiceRedirectVerb <Redirect> TwiML Verb
type VoiceRedirectVerb struct {
    // url: Redirect URL
    // method: Redirect URL method
    // OptionalAttributes: additional attributes
    Url string
    Method string
    Nouns []Noun
    OptionalAttributes map[string]string
}

func (m VoiceRedirectVerb) GetName() string {
    return "Redirect"
}

func (m VoiceRedirectVerb) GetText() string {
    return m.Url
}

func (m VoiceRedirectVerb) GetAttr() (map[string]string, map[string]string) {
    paramsAttr := map[string]string{
    	"Method": m.Method,
    }
    return m.OptionalAttributes, paramsAttr
}

func (m VoiceRedirectVerb) GetNouns() []Noun {
    return m.Nouns
}
//VoiceRecordVerb <Record> TwiML Verb
type VoiceRecordVerb struct {
    // action: Action URL
    // method: Action URL method
    // timeout: Timeout to begin recording
    // finish_on_key: Finish recording on key
    // max_length: Max time to record in seconds
    // play_beep: Play beep
    // trim: Trim the recording
    // recording_status_callback: Status callback URL
    // recording_status_callback_method: Status callback URL method
    // recording_status_callback_event: Recording status callback events
    // transcribe: Transcribe the recording
    // transcribe_callback: Transcribe callback URL
    // OptionalAttributes: additional attributes
    Action string
    Method string
    Timeout string
    FinishOnKey string
    MaxLength string
    PlayBeep string
    Trim string
    RecordingStatusCallback string
    RecordingStatusCallbackMethod string
    RecordingStatusCallbackEvent string
    Transcribe string
    TranscribeCallback string
    Nouns []Noun
    OptionalAttributes map[string]string
}

func (m VoiceRecordVerb) GetName() string {
    return "Record"
}

func (m VoiceRecordVerb) GetText() string {
    return ""
}

func (m VoiceRecordVerb) GetAttr() (map[string]string, map[string]string) {
    paramsAttr := map[string]string{
    	"Action": m.Action,
    	"Method": m.Method,
    	"Timeout": m.Timeout,
    	"FinishOnKey": m.FinishOnKey,
    	"MaxLength": m.MaxLength,
    	"PlayBeep": m.PlayBeep,
    	"Trim": m.Trim,
    	"RecordingStatusCallback": m.RecordingStatusCallback,
    	"RecordingStatusCallbackMethod": m.RecordingStatusCallbackMethod,
    	"RecordingStatusCallbackEvent": m.RecordingStatusCallbackEvent,
    	"Transcribe": m.Transcribe,
    	"TranscribeCallback": m.TranscribeCallback,
    }
    return m.OptionalAttributes, paramsAttr
}

func (m VoiceRecordVerb) GetNouns() []Noun {
    return m.Nouns
}
//VoiceQueueNoun <Queue> TwiML Noun
type VoiceQueueNoun struct {
    // name: Queue name
    // url: Action URL
    // method: Action URL method
    // reservation_sid: TaskRouter Reservation SID
    // post_work_activity_sid: TaskRouter Activity SID
    // OptionalAttributes: additional attributes
    Name string
    Url string
    Method string
    ReservationSid string
    PostWorkActivitySid string
    OptionalAttributes map[string]string
}

func (m VoiceQueueNoun) GetName() string {
    return "Queue"
}

func (m VoiceQueueNoun) GetText() string {
    return m.Name
}

func (m VoiceQueueNoun) GetAttr() (map[string]string, map[string]string) {
    paramsAttr := map[string]string{
    	"Url": m.Url,
    	"Method": m.Method,
    	"ReservationSid": m.ReservationSid,
    	"PostWorkActivitySid": m.PostWorkActivitySid,
    }
    return m.OptionalAttributes, paramsAttr
}
//VoiceLeaveVerb <Leave> TwiML Verb
type VoiceLeaveVerb struct {
    // OptionalAttributes: additional attributes
    Nouns []Noun
    OptionalAttributes map[string]string
}

func (m VoiceLeaveVerb) GetName() string {
    return "Leave"
}

func (m VoiceLeaveVerb) GetText() string {
    return ""
}

func (m VoiceLeaveVerb) GetAttr() (map[string]string, map[string]string) {
    return m.OptionalAttributes, nil
}

func (m VoiceLeaveVerb) GetNouns() []Noun {
    return m.Nouns
}
//VoiceHangupVerb <Hangup> TwiML Verb
type VoiceHangupVerb struct {
    // OptionalAttributes: additional attributes
    Nouns []Noun
    OptionalAttributes map[string]string
}

func (m VoiceHangupVerb) GetName() string {
    return "Hangup"
}

func (m VoiceHangupVerb) GetText() string {
    return ""
}

func (m VoiceHangupVerb) GetAttr() (map[string]string, map[string]string) {
    return m.OptionalAttributes, nil
}

func (m VoiceHangupVerb) GetNouns() []Noun {
    return m.Nouns
}
//VoiceGatherVerb <Gather> TwiML Verb
type VoiceGatherVerb struct {
    // input: Input type Twilio should accept
    // action: Action URL
    // method: Action URL method
    // timeout: Time to wait to gather input
    // speech_timeout: Time to wait to gather speech input and it should be either auto or a positive integer.
    // max_speech_time: Max allowed time for speech input
    // profanity_filter: Profanity Filter on speech
    // finish_on_key: Finish gather on key
    // num_digits: Number of digits to collect
    // partial_result_callback: Partial result callback URL
    // partial_result_callback_method: Partial result callback URL method
    // language: Language to use
    // hints: Speech recognition hints
    // barge_in: Stop playing media upon speech
    // debug: Allow debug for gather
    // action_on_empty_result: Force webhook to the action URL event if there is no input
    // speech_model: Specify the model that is best suited for your use case
    // enhanced: Use enhanced speech model
    // OptionalAttributes: additional attributes
    Input string
    Action string
    Method string
    Timeout string
    SpeechTimeout string
    MaxSpeechTime string
    ProfanityFilter string
    FinishOnKey string
    NumDigits string
    PartialResultCallback string
    PartialResultCallbackMethod string
    Language string
    Hints string
    BargeIn string
    Debug string
    ActionOnEmptyResult string
    SpeechModel string
    Enhanced string
    Nouns []Noun
    OptionalAttributes map[string]string
}

func (m VoiceGatherVerb) GetName() string {
    return "Gather"
}

func (m VoiceGatherVerb) GetText() string {
    return ""
}

func (m VoiceGatherVerb) GetAttr() (map[string]string, map[string]string) {
    paramsAttr := map[string]string{
    	"Input": m.Input,
    	"Action": m.Action,
    	"Method": m.Method,
    	"Timeout": m.Timeout,
    	"SpeechTimeout": m.SpeechTimeout,
    	"MaxSpeechTime": m.MaxSpeechTime,
    	"ProfanityFilter": m.ProfanityFilter,
    	"FinishOnKey": m.FinishOnKey,
    	"NumDigits": m.NumDigits,
    	"PartialResultCallback": m.PartialResultCallback,
    	"PartialResultCallbackMethod": m.PartialResultCallbackMethod,
    	"Language": m.Language,
    	"Hints": m.Hints,
    	"BargeIn": m.BargeIn,
    	"Debug": m.Debug,
    	"ActionOnEmptyResult": m.ActionOnEmptyResult,
    	"SpeechModel": m.SpeechModel,
    	"Enhanced": m.Enhanced,
    }
    return m.OptionalAttributes, paramsAttr
}

func (m VoiceGatherVerb) GetNouns() []Noun {
    return m.Nouns
}
//VoiceEnqueueNoun <Enqueue> TwiML Noun
type VoiceEnqueueNoun struct {
    // name: Friendly name
    // action: Action URL
    // max_queue_size: Maximum size of queue
    // method: Action URL method
    // wait_url: Wait URL
    // wait_url_method: Wait URL method
    // workflow_sid: TaskRouter Workflow SID
    // OptionalAttributes: additional attributes
    Name string
    Action string
    MaxQueueSize string
    Method string
    WaitUrl string
    WaitUrlMethod string
    WorkflowSid string
    OptionalAttributes map[string]string
}

func (m VoiceEnqueueNoun) GetName() string {
    return "Enqueue"
}

func (m VoiceEnqueueNoun) GetText() string {
    return m.Name
}

func (m VoiceEnqueueNoun) GetAttr() (map[string]string, map[string]string) {
    paramsAttr := map[string]string{
    	"Action": m.Action,
    	"MaxQueueSize": m.MaxQueueSize,
    	"Method": m.Method,
    	"WaitUrl": m.WaitUrl,
    	"WaitUrlMethod": m.WaitUrlMethod,
    	"WorkflowSid": m.WorkflowSid,
    }
    return m.OptionalAttributes, paramsAttr
}
//VoiceTaskNoun <Task> TwiML Noun
type VoiceTaskNoun struct {
    // body: TaskRouter task attributes
    // priority: Task priority
    // timeout: Timeout associated with task
    // OptionalAttributes: additional attributes
    Body string
    Priority string
    Timeout string
    OptionalAttributes map[string]string
}

func (m VoiceTaskNoun) GetName() string {
    return "Task"
}

func (m VoiceTaskNoun) GetText() string {
    return m.Body
}

func (m VoiceTaskNoun) GetAttr() (map[string]string, map[string]string) {
    paramsAttr := map[string]string{
    	"Priority": m.Priority,
    	"Timeout": m.Timeout,
    }
    return m.OptionalAttributes, paramsAttr
}
//VoiceEchoVerb <Echo> TwiML Verb
type VoiceEchoVerb struct {
    // OptionalAttributes: additional attributes
    Nouns []Noun
    OptionalAttributes map[string]string
}

func (m VoiceEchoVerb) GetName() string {
    return "Echo"
}

func (m VoiceEchoVerb) GetText() string {
    return ""
}

func (m VoiceEchoVerb) GetAttr() (map[string]string, map[string]string) {
    return m.OptionalAttributes, nil
}

func (m VoiceEchoVerb) GetNouns() []Noun {
    return m.Nouns
}
//VoiceDialVerb <Dial> TwiML Verb
type VoiceDialVerb struct {
    // number: Phone number to dial
    // action: Action URL
    // method: Action URL method
    // timeout: Time to wait for answer
    // hangup_on_star: Hangup call on star press
    // time_limit: Max time length
    // caller_id: Caller ID to display
    // record: Record the call
    // trim: Trim the recording
    // recording_status_callback: Recording status callback URL
    // recording_status_callback_method: Recording status callback URL method
    // recording_status_callback_event: Recording status callback events
    // answer_on_bridge: Preserve the ringing behavior of the inbound call until the Dialed call picks up
    // ring_tone: Ringtone allows you to override the ringback tone that Twilio will play back to the caller while executing the Dial
    // recording_track: To indicate which audio track should be recorded
    // sequential: Used to determine if child TwiML nouns should be dialed in order, one after the other (sequential) or dial all at once (parallel). Default is false, parallel
    // refer_url: Webhook that will receive future SIP REFER requests
    // refer_method: The HTTP method to use for the refer Webhook
    // OptionalAttributes: additional attributes
    Number string
    Action string
    Method string
    Timeout string
    HangupOnStar string
    TimeLimit string
    CallerId string
    Record string
    Trim string
    RecordingStatusCallback string
    RecordingStatusCallbackMethod string
    RecordingStatusCallbackEvent string
    AnswerOnBridge string
    RingTone string
    RecordingTrack string
    Sequential string
    ReferUrl string
    ReferMethod string
    Nouns []Noun
    OptionalAttributes map[string]string
}

func (m VoiceDialVerb) GetName() string {
    return "Dial"
}

func (m VoiceDialVerb) GetText() string {
    return m.Number
}

func (m VoiceDialVerb) GetAttr() (map[string]string, map[string]string) {
    paramsAttr := map[string]string{
    	"Action": m.Action,
    	"Method": m.Method,
    	"Timeout": m.Timeout,
    	"HangupOnStar": m.HangupOnStar,
    	"TimeLimit": m.TimeLimit,
    	"CallerId": m.CallerId,
    	"Record": m.Record,
    	"Trim": m.Trim,
    	"RecordingStatusCallback": m.RecordingStatusCallback,
    	"RecordingStatusCallbackMethod": m.RecordingStatusCallbackMethod,
    	"RecordingStatusCallbackEvent": m.RecordingStatusCallbackEvent,
    	"AnswerOnBridge": m.AnswerOnBridge,
    	"RingTone": m.RingTone,
    	"RecordingTrack": m.RecordingTrack,
    	"Sequential": m.Sequential,
    	"ReferUrl": m.ReferUrl,
    	"ReferMethod": m.ReferMethod,
    }
    return m.OptionalAttributes, paramsAttr
}

func (m VoiceDialVerb) GetNouns() []Noun {
    return m.Nouns
}
//VoiceSipNoun <Sip> TwiML Noun
type VoiceSipNoun struct {
    // sip_url: SIP URL
    // username: SIP Username
    // password: SIP Password
    // url: Action URL
    // method: Action URL method
    // status_callback_event: Status callback events
    // status_callback: Status callback URL
    // status_callback_method: Status callback URL method
    // OptionalAttributes: additional attributes
    SipUrl string
    Username string
    Password string
    Url string
    Method string
    StatusCallbackEvent string
    StatusCallback string
    StatusCallbackMethod string
    OptionalAttributes map[string]string
}

func (m VoiceSipNoun) GetName() string {
    return "Sip"
}

func (m VoiceSipNoun) GetText() string {
    return m.SipUrl
}

func (m VoiceSipNoun) GetAttr() (map[string]string, map[string]string) {
    paramsAttr := map[string]string{
    	"Username": m.Username,
    	"Password": m.Password,
    	"Url": m.Url,
    	"Method": m.Method,
    	"StatusCallbackEvent": m.StatusCallbackEvent,
    	"StatusCallback": m.StatusCallback,
    	"StatusCallbackMethod": m.StatusCallbackMethod,
    }
    return m.OptionalAttributes, paramsAttr
}
//VoiceSimNoun <Sim> TwiML Noun
type VoiceSimNoun struct {
    // sim_sid: SIM SID
    // OptionalAttributes: additional attributes
    SimSid string
    OptionalAttributes map[string]string
}

func (m VoiceSimNoun) GetName() string {
    return "Sim"
}

func (m VoiceSimNoun) GetText() string {
    return m.SimSid
}

func (m VoiceSimNoun) GetAttr() (map[string]string, map[string]string) {
    return m.OptionalAttributes, nil
}
//VoiceNumberNoun <Number> TwiML Noun
type VoiceNumberNoun struct {
    // phone_number: Phone Number to dial
    // send_digits: DTMF tones to play when the call is answered
    // url: TwiML URL
    // method: TwiML URL method
    // status_callback_event: Events to call status callback
    // status_callback: Status callback URL
    // status_callback_method: Status callback URL method
    // byoc: BYOC trunk SID (Beta)
    // OptionalAttributes: additional attributes
    PhoneNumber string
    SendDigits string
    Url string
    Method string
    StatusCallbackEvent string
    StatusCallback string
    StatusCallbackMethod string
    Byoc string
    OptionalAttributes map[string]string
}

func (m VoiceNumberNoun) GetName() string {
    return "Number"
}

func (m VoiceNumberNoun) GetText() string {
    return m.PhoneNumber
}

func (m VoiceNumberNoun) GetAttr() (map[string]string, map[string]string) {
    paramsAttr := map[string]string{
    	"SendDigits": m.SendDigits,
    	"Url": m.Url,
    	"Method": m.Method,
    	"StatusCallbackEvent": m.StatusCallbackEvent,
    	"StatusCallback": m.StatusCallback,
    	"StatusCallbackMethod": m.StatusCallbackMethod,
    	"Byoc": m.Byoc,
    }
    return m.OptionalAttributes, paramsAttr
}
//VoiceConferenceNoun <Conference> TwiML Noun
type VoiceConferenceNoun struct {
    // name: Conference name
    // muted: Join the conference muted
    // beep: Play beep when joining
    // start_conference_on_enter: Start the conference on enter
    // end_conference_on_exit: End the conferenceon exit
    // wait_url: Wait URL
    // wait_method: Wait URL method
    // max_participants: Maximum number of participants
    // record: Record the conference
    // region: Conference region
    // coach: Call coach
    // trim: Trim the conference recording
    // status_callback_event: Events to call status callback URL
    // status_callback: Status callback URL
    // status_callback_method: Status callback URL method
    // recording_status_callback: Recording status callback URL
    // recording_status_callback_method: Recording status callback URL method
    // recording_status_callback_event: Recording status callback events
    // event_callback_url: Event callback URL
    // jitter_buffer_size: Size of jitter buffer for participant
    // participant_label: A label for participant
    // OptionalAttributes: additional attributes
    Name string
    Muted string
    Beep string
    StartConferenceOnEnter string
    EndConferenceOnExit string
    WaitUrl string
    WaitMethod string
    MaxParticipants string
    Record string
    Region string
    Coach string
    Trim string
    StatusCallbackEvent string
    StatusCallback string
    StatusCallbackMethod string
    RecordingStatusCallback string
    RecordingStatusCallbackMethod string
    RecordingStatusCallbackEvent string
    EventCallbackUrl string
    JitterBufferSize string
    ParticipantLabel string
    OptionalAttributes map[string]string
}

func (m VoiceConferenceNoun) GetName() string {
    return "Conference"
}

func (m VoiceConferenceNoun) GetText() string {
    return m.Name
}

func (m VoiceConferenceNoun) GetAttr() (map[string]string, map[string]string) {
    paramsAttr := map[string]string{
    	"Muted": m.Muted,
    	"Beep": m.Beep,
    	"StartConferenceOnEnter": m.StartConferenceOnEnter,
    	"EndConferenceOnExit": m.EndConferenceOnExit,
    	"WaitUrl": m.WaitUrl,
    	"WaitMethod": m.WaitMethod,
    	"MaxParticipants": m.MaxParticipants,
    	"Record": m.Record,
    	"Region": m.Region,
    	"Coach": m.Coach,
    	"Trim": m.Trim,
    	"StatusCallbackEvent": m.StatusCallbackEvent,
    	"StatusCallback": m.StatusCallback,
    	"StatusCallbackMethod": m.StatusCallbackMethod,
    	"RecordingStatusCallback": m.RecordingStatusCallback,
    	"RecordingStatusCallbackMethod": m.RecordingStatusCallbackMethod,
    	"RecordingStatusCallbackEvent": m.RecordingStatusCallbackEvent,
    	"EventCallbackUrl": m.EventCallbackUrl,
    	"JitterBufferSize": m.JitterBufferSize,
    	"ParticipantLabel": m.ParticipantLabel,
    }
    return m.OptionalAttributes, paramsAttr
}
//VoiceClientNoun <Client> TwiML Noun
type VoiceClientNoun struct {
    // identity: Client identity
    // url: Client URL
    // method: Client URL Method
    // status_callback_event: Events to trigger status callback
    // status_callback: Status Callback URL
    // status_callback_method: Status Callback URL Method
    // OptionalAttributes: additional attributes
    Identity string
    Url string
    Method string
    StatusCallbackEvent string
    StatusCallback string
    StatusCallbackMethod string
    OptionalAttributes map[string]string
}

func (m VoiceClientNoun) GetName() string {
    return "Client"
}

func (m VoiceClientNoun) GetText() string {
    return m.Identity
}

func (m VoiceClientNoun) GetAttr() (map[string]string, map[string]string) {
    paramsAttr := map[string]string{
    	"Url": m.Url,
    	"Method": m.Method,
    	"StatusCallbackEvent": m.StatusCallbackEvent,
    	"StatusCallback": m.StatusCallback,
    	"StatusCallbackMethod": m.StatusCallbackMethod,
    }
    return m.OptionalAttributes, paramsAttr
}
//VoiceIdentityNoun <Identity> TwiML Noun
type VoiceIdentityNoun struct {
    // client_identity: Identity of the client to dial
    // OptionalAttributes: additional attributes
    ClientIdentity string
    OptionalAttributes map[string]string
}

func (m VoiceIdentityNoun) GetName() string {
    return "Identity"
}

func (m VoiceIdentityNoun) GetText() string {
    return m.ClientIdentity
}

func (m VoiceIdentityNoun) GetAttr() (map[string]string, map[string]string) {
    return m.OptionalAttributes, nil
}
//VoiceConnectVerb <Connect> TwiML Verb
type VoiceConnectVerb struct {
    // action: Action URL
    // method: Action URL method
    // OptionalAttributes: additional attributes
    Action string
    Method string
    Nouns []Noun
    OptionalAttributes map[string]string
}

func (m VoiceConnectVerb) GetName() string {
    return "Connect"
}

func (m VoiceConnectVerb) GetText() string {
    return ""
}

func (m VoiceConnectVerb) GetAttr() (map[string]string, map[string]string) {
    paramsAttr := map[string]string{
    	"Action": m.Action,
    	"Method": m.Method,
    }
    return m.OptionalAttributes, paramsAttr
}

func (m VoiceConnectVerb) GetNouns() []Noun {
    return m.Nouns
}
//VoiceConversationNoun <Conversation> TwiML Noun
type VoiceConversationNoun struct {
    // service_instance_sid: Service instance Sid
    // inbound_autocreation: Inbound autocreation
    // routing_assignment_timeout: Routing assignment timeout
    // inbound_timeout: Inbound timeout
    // record: Record
    // trim: Trim
    // recording_status_callback: Recording status callback URL
    // recording_status_callback_method: Recording status callback URL method
    // recording_status_callback_event: Recording status callback events
    // status_callback: Status callback URL
    // status_callback_method: Status callback URL method
    // status_callback_event: Events to call status callback URL
    // OptionalAttributes: additional attributes
    ServiceInstanceSid string
    InboundAutocreation string
    RoutingAssignmentTimeout string
    InboundTimeout string
    Record string
    Trim string
    RecordingStatusCallback string
    RecordingStatusCallbackMethod string
    RecordingStatusCallbackEvent string
    StatusCallback string
    StatusCallbackMethod string
    StatusCallbackEvent string
    OptionalAttributes map[string]string
}

func (m VoiceConversationNoun) GetName() string {
    return "Conversation"
}

func (m VoiceConversationNoun) GetText() string {
    return ""
}

func (m VoiceConversationNoun) GetAttr() (map[string]string, map[string]string) {
    paramsAttr := map[string]string{
    	"ServiceInstanceSid": m.ServiceInstanceSid,
    	"InboundAutocreation": m.InboundAutocreation,
    	"RoutingAssignmentTimeout": m.RoutingAssignmentTimeout,
    	"InboundTimeout": m.InboundTimeout,
    	"Record": m.Record,
    	"Trim": m.Trim,
    	"RecordingStatusCallback": m.RecordingStatusCallback,
    	"RecordingStatusCallbackMethod": m.RecordingStatusCallbackMethod,
    	"RecordingStatusCallbackEvent": m.RecordingStatusCallbackEvent,
    	"StatusCallback": m.StatusCallback,
    	"StatusCallbackMethod": m.StatusCallbackMethod,
    	"StatusCallbackEvent": m.StatusCallbackEvent,
    }
    return m.OptionalAttributes, paramsAttr
}
//VoiceVirtualAgentNoun <VirtualAgent> TwiML Noun
type VoiceVirtualAgentNoun struct {
    // connector_name: Defines the conversation profile Dialogflow needs to use
    // language: Language to be used by Dialogflow to transcribe speech
    // sentiment_analysis: Whether sentiment analysis needs to be enabled or not
    // status_callback: URL to post status callbacks from Twilio
    // status_callback_method: HTTP method to use when requesting the status callback URL
    // OptionalAttributes: additional attributes
    ConnectorName string
    Language string
    SentimentAnalysis string
    StatusCallback string
    StatusCallbackMethod string
    OptionalAttributes map[string]string
}

func (m VoiceVirtualAgentNoun) GetName() string {
    return "VirtualAgent"
}

func (m VoiceVirtualAgentNoun) GetText() string {
    return ""
}

func (m VoiceVirtualAgentNoun) GetAttr() (map[string]string, map[string]string) {
    paramsAttr := map[string]string{
    	"ConnectorName": m.ConnectorName,
    	"Language": m.Language,
    	"SentimentAnalysis": m.SentimentAnalysis,
    	"StatusCallback": m.StatusCallback,
    	"StatusCallbackMethod": m.StatusCallbackMethod,
    }
    return m.OptionalAttributes, paramsAttr
}
//VoiceConfigNoun <Config> TwiML Noun
type VoiceConfigNoun struct {
    // name: The name of the custom config
    // value: The value of the custom config
    // OptionalAttributes: additional attributes
    Name string
    Value string
    OptionalAttributes map[string]string
}

func (m VoiceConfigNoun) GetName() string {
    return "Config"
}

func (m VoiceConfigNoun) GetText() string {
    return ""
}

func (m VoiceConfigNoun) GetAttr() (map[string]string, map[string]string) {
    paramsAttr := map[string]string{
    	"Name": m.Name,
    	"Value": m.Value,
    }
    return m.OptionalAttributes, paramsAttr
}
//VoiceAutopilotNoun <Autopilot> TwiML Noun
type VoiceAutopilotNoun struct {
    // name: Autopilot assistant sid or unique name
    // OptionalAttributes: additional attributes
    Name string
    OptionalAttributes map[string]string
}

func (m VoiceAutopilotNoun) GetName() string {
    return "Autopilot"
}

func (m VoiceAutopilotNoun) GetText() string {
    return m.Name
}

func (m VoiceAutopilotNoun) GetAttr() (map[string]string, map[string]string) {
    return m.OptionalAttributes, nil
}
//VoiceRoomNoun <Room> TwiML Noun
type VoiceRoomNoun struct {
    // name: Room name
    // participant_identity: Participant identity when connecting to the Room
    // OptionalAttributes: additional attributes
    Name string
    ParticipantIdentity string
    OptionalAttributes map[string]string
}

func (m VoiceRoomNoun) GetName() string {
    return "Room"
}

func (m VoiceRoomNoun) GetText() string {
    return m.Name
}

func (m VoiceRoomNoun) GetAttr() (map[string]string, map[string]string) {
    paramsAttr := map[string]string{
    	"ParticipantIdentity": m.ParticipantIdentity,
    }
    return m.OptionalAttributes, paramsAttr
}