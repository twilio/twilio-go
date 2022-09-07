package twiml_test

import (
	"strings"
	"testing"

	"github.com/beevik/etree"
	"github.com/stretchr/testify/assert"
	"github.com/twilio/twilio-go/twiml"
)

type MockElement struct {
	name          string
	text          string
	optAttrs      map[string]string
	paramAttrs    map[string]string
	innerElements []twiml.Element
}

func (e MockElement) GetName() string {
	return e.name
}

func (e MockElement) GetText() string {
	return e.text
}

func (e MockElement) GetAttr() (map[string]string, map[string]string) {
	return e.optAttrs, e.paramAttrs
}

func (e MockElement) GetInnerElements() []twiml.Element {
	return e.innerElements
}

// Tests
func TestCreateDocument_ReturnsResponseElement(t *testing.T) {
	_, e := twiml.CreateDocument()
	assert.Equal(t, "Response", e.Tag)
}

func TestAddAllVerbs_EmptyVerbList(t *testing.T) {
	verbs := make([]twiml.Element, 0)
	e := etree.Element{}
	twiml.AddAllVerbs(&e, verbs)
	assert.Empty(t, e.ChildElements())
}

func TestAddAllVerbs_VerbListOfOne(t *testing.T) {
	v := MockElement{
		"Gather",
		"Gather things",
		map[string]string{"Input": "Speech"},
		map[string]string{"Action": "https://demo.twilio.com"},
		[]twiml.Element{},
	}
	verbs := []twiml.Element{v}
	resEl := etree.Element{}

	twiml.AddAllVerbs(&resEl, verbs)

	assertDefinesTwiMLXMLElements(t, &resEl, verbs)
}

func TestAddAllVerbs_ManyVerbList(t *testing.T) {
	v1 := MockElement{
		"Pay",
		"Enter card numbers",
		map[string]string{"Timeout": "10"},
		map[string]string{"PaymentMethod": "ach-debit", "Description": "A description"},
		[]twiml.Element{},
	}
	v2 := MockElement{
		"Say",
		"Thank you!",
		map[string]string{"Voice": "alice", "Language": "en"},
		map[string]string{"Loop": "1"},
		[]twiml.Element{}}
	verbs := []twiml.Element{v1, v2}
	resEl := etree.Element{}

	twiml.AddAllVerbs(&resEl, verbs)

	assertDefinesTwiMLXMLElements(t, &resEl, verbs)
}

func TestAddAllVerbs_VerbWithManyNouns(t *testing.T) {
	n1 := MockElement{
		"Autopilot",
		"Jarvis",
		map[string]string{"Foo": "Bar"},
		map[string]string{"Action": "https://demo.twilio.com", "Method": "GET"},
		make([]twiml.Element, 0),
	}
	n2 := MockElement{
		"Room",
		"Conference room",
		map[string]string{"participantIdentity": "Alice"},
		map[string]string{},
		make([]twiml.Element, 0),
	}
	v := MockElement{
		"Connect",
		"Connect things",
		map[string]string{},
		map[string]string{},
		[]twiml.Element{n1, n2},
	}
	verbs := []twiml.Element{v}
	resEl := etree.Element{}

	twiml.AddAllVerbs(&resEl, verbs)

	assertDefinesTwiMLXMLElements(t, &resEl, verbs)
}

func TestAddAllVerbs_VerbWithManyAttributes(t *testing.T) {
	v := MockElement{
		"Gather",
		"Gather things",
		map[string]string{
			"Input":                 "Speech",
			"PartialResultCallback": "https://demo.twilio.com",
		},
		map[string]string{
			"Action": "https://demo.twilio.com",
			"Method": "POST",
		},
		[]twiml.Element{},
	}
	verbs := []twiml.Element{v}
	resEl := etree.Element{}

	twiml.AddAllVerbs(&resEl, verbs)

	assertDefinesTwiMLXMLElements(t, &resEl, verbs)
}

func TestToXML(t *testing.T) {
	doc := etree.NewDocument()
	doc.CreateProcInst("xml", `version="1.0" encoding="UTF-8"`)
	resEl := doc.CreateElement("Response")
	childEl := etree.NewElement("Say")
	childEl.SetText("Hello World!")
	childEl.CreateAttr("Voice", "alice")
	resEl.AddChild(childEl)
	expected := `<?xml version="1.0" encoding="UTF-8"?><Response><Say Voice="alice">Hello World!</Say></Response>`

	xml, _ := twiml.ToXML(doc)

	assert.Equal(t,
		expected,
		xml,
		"Output XML string should match expected value.")
}

// Utilities
/*
Asserts that a given XML element contains the list of TwiML Elements as
child XML elements.
*/
func assertDefinesTwiMLXMLElements(t *testing.T, el *etree.Element, elements []twiml.Element) {
	assert.Equal(t,
		len(elements),
		len(el.ChildElements()),
		"Parent element should have same number of child elements as child TwiML Elements.")

	for _, element := range elements {
		xmlEl := el.FindElement(element.GetName())
		assert.NotNil(t,
			xmlEl,
			"Parent element should contain a child corresponding to the twiML element.")
		assert.Equal(t, element.GetName(), xmlEl.Tag, "Child element tag should match twiML element name.")
		assert.Equal(t, element.GetText(), xmlEl.Text(), "Child element text should match twiML element text.")

		// Attributes
		optAttrs, paramAttrs := element.GetAttr()
		assertHasVerbNounAttributes(t, xmlEl, optAttrs)
		assertHasVerbNounAttributes(t, xmlEl, paramAttrs)

		// Nouns
		assertDefinesTwiMLXMLElements(t, xmlEl, element.GetInnerElements())
	}
}

/*
Asserts that the given XML element defines attributes and their values
matching the attributes in a map of TwiML element attribute name, value
pairs.
*/
func assertHasVerbNounAttributes(t *testing.T, el *etree.Element, attrs map[string]string) {
	for k, v := range attrs {
		name := toFirstCharacterLowerCase(k) // TwiML XML attribute names are lower camelcase
		attr := el.SelectAttr(name)
		assert.NotNil(t, attr, "Element should define the verb attribute.")
		assert.Equal(t, v, attr.Value, "Element attribute value should match verb attribute value.")
	}
}

/*
Makes the first character of the given string lower case and returns the value
as a new string.
*/
func toFirstCharacterLowerCase(s string) string {
	return strings.ToLower(string(s[0])) + s[1:]
}
