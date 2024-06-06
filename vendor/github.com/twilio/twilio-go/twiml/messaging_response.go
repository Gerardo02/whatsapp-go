//This code was generated by
//\ / _    _  _|   _  _
// | (_)\/(_)(_|\/| |(/_  v1.0.0
//      /       /

package twiml

func Messages(verbs []Element) (string, error) {
	doc, response := CreateDocument()
	if verbs != nil {
		AddAllVerbs(response, verbs)
	}
	return ToXML(doc)
}

//MessagingRedirect <Redirect> TwiML Verb
type MessagingRedirect struct {
	// url: Redirect URL
	// method: Redirect URL method
	// OptionalAttributes: additional attributes
	Url                string
	Method             string
	InnerElements      []Element
	OptionalAttributes map[string]string
}

func (m MessagingRedirect) GetName() string {
	return "Redirect"
}

func (m MessagingRedirect) GetText() string {
	return m.Url
}

func (m MessagingRedirect) GetAttr() (map[string]string, map[string]string) {
	paramsAttr := map[string]string{
		"Method": m.Method,
	}
	return m.OptionalAttributes, paramsAttr
}

func (m MessagingRedirect) GetInnerElements() []Element {
	return m.InnerElements
}

//MessagingMessage <Message> TwiML Verb
type MessagingMessage struct {
	// body: Message Body
	// to: Phone Number to send Message to
	// from: Phone Number to send Message from
	// action: A URL specifying where Twilio should send status callbacks for the created outbound message.
	// method: Action URL Method
	// status_callback: Status callback URL. Deprecated in favor of action.
	// OptionalAttributes: additional attributes
	Body               string
	To                 string
	From               string
	Action             string
	Method             string
	StatusCallback     string
	InnerElements      []Element
	OptionalAttributes map[string]string
}

func (m MessagingMessage) GetName() string {
	return "Message"
}

func (m MessagingMessage) GetText() string {
	return m.Body
}

func (m MessagingMessage) GetAttr() (map[string]string, map[string]string) {
	paramsAttr := map[string]string{
		"To":             m.To,
		"From":           m.From,
		"Action":         m.Action,
		"Method":         m.Method,
		"StatusCallback": m.StatusCallback,
	}
	return m.OptionalAttributes, paramsAttr
}

func (m MessagingMessage) GetInnerElements() []Element {
	return m.InnerElements
}

//MessagingMedia <Media> TwiML Noun
type MessagingMedia struct {
	// url: Media URL
	// OptionalAttributes: additional attributes
	Url                string
	InnerElements      []Element
	OptionalAttributes map[string]string
}

func (m MessagingMedia) GetName() string {
	return "Media"
}

func (m MessagingMedia) GetText() string {
	return m.Url
}

func (m MessagingMedia) GetAttr() (map[string]string, map[string]string) {
	return m.OptionalAttributes, nil
}

func (m MessagingMedia) GetInnerElements() []Element {
	return m.InnerElements
}

//MessagingBody <Body> TwiML Noun
type MessagingBody struct {
	// message: Message Body
	// OptionalAttributes: additional attributes
	Message            string
	InnerElements      []Element
	OptionalAttributes map[string]string
}

func (m MessagingBody) GetName() string {
	return "Body"
}

func (m MessagingBody) GetText() string {
	return m.Message
}

func (m MessagingBody) GetAttr() (map[string]string, map[string]string) {
	return m.OptionalAttributes, nil
}

func (m MessagingBody) GetInnerElements() []Element {
	return m.InnerElements
}
