package twiml

func Messages(verbs []Verb)(TwimlXml string){
    doc, response := CreateDocument()
    if verbs != nil {
        AddAllVerbs(response, verbs)
    }
    return ToXML(doc)
}

//MessagingRedirectVerb <Redirect> TwiML Verb
type MessagingRedirectVerb struct {
    // url: Redirect URL
    // method: Redirect URL method
    // OptionalAttributes: additional attributes
    Url string
    Method string
    Nouns []Noun
    OptionalAttributes map[string]string
}

func (m MessagingRedirectVerb) GetName() string {
    return "Redirect"
}

func (m MessagingRedirectVerb) GetText() string {
    return m.Url
}

func (m MessagingRedirectVerb) GetAttr() (map[string]string, map[string]string) {
    paramsAttr := map[string]string{
    	"Method": m.Method,
    }
    return m.OptionalAttributes, paramsAttr
}

func (m MessagingRedirectVerb) GetNouns() []Noun {
    return m.Nouns
}
//MessagingMessageVerb <Message> TwiML Verb
type MessagingMessageVerb struct {
    // body: Message Body
    // to: Phone Number to send Message to
    // from: Phone Number to send Message from
    // action: Action URL
    // method: Action URL Method
    // status_callback: Status callback URL. Deprecated in favor of action.
    // OptionalAttributes: additional attributes
    Body string
    To string
    From string
    Action string
    Method string
    StatusCallback string
    Nouns []Noun
    OptionalAttributes map[string]string
}

func (m MessagingMessageVerb) GetName() string {
    return "Message"
}

func (m MessagingMessageVerb) GetText() string {
    return m.Body
}

func (m MessagingMessageVerb) GetAttr() (map[string]string, map[string]string) {
    paramsAttr := map[string]string{
    	"To": m.To,
    	"From": m.From,
    	"Action": m.Action,
    	"Method": m.Method,
    	"StatusCallback": m.StatusCallback,
    }
    return m.OptionalAttributes, paramsAttr
}

func (m MessagingMessageVerb) GetNouns() []Noun {
    return m.Nouns
}
//MessagingMediaNoun <Media> TwiML Noun
type MessagingMediaNoun struct {
    // url: Media URL
    // OptionalAttributes: additional attributes
    Url string
    OptionalAttributes map[string]string
}

func (m MessagingMediaNoun) GetName() string {
    return "Media"
}

func (m MessagingMediaNoun) GetText() string {
    return m.Url
}

func (m MessagingMediaNoun) GetAttr() (map[string]string, map[string]string) {
    return m.OptionalAttributes, nil
}
//MessagingBodyNoun <Body> TwiML Noun
type MessagingBodyNoun struct {
    // message: Message Body
    // OptionalAttributes: additional attributes
    Message string
    OptionalAttributes map[string]string
}

func (m MessagingBodyNoun) GetName() string {
    return "Body"
}

func (m MessagingBodyNoun) GetText() string {
    return m.Message
}

func (m MessagingBodyNoun) GetAttr() (map[string]string, map[string]string) {
    return m.OptionalAttributes, nil
}