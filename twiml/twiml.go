package twiml

import (
	"github.com/beevik/etree"
	"strings"
)

type Noun interface {
	GetName() string
	GetText() string
	GetAttr() (map[string]string, map[string]string)
}

type Verb interface {
	GetName() string
	GetText() string
	GetAttr() (map[string]string, map[string]string)
	GetNouns() []Noun
}

func AddAllVerbs(response *etree.Element, verbs []Verb) {
	for _, verb := range verbs {
		verbEl := createVerbElement(verb)
		response.AddChild(verbEl)
	}
}

func createVerbElement(verb Verb) *etree.Element {
	el := etree.NewElement(verb.GetName())
	optAttr, paramAttr := verb.GetAttr()
	addPropertyToElement(el, verb.GetText(), optAttr, paramAttr)
	//Loop through all Nouns
	if len(verb.GetNouns()) != 0 {
		for _, noun := range verb.GetNouns() {
			child := createNounElement(noun)
			el.AddChild(child)
		}
	}
	return el
}

func createNounElement(noun Noun) *etree.Element {
	el := etree.NewElement(noun.GetName())
	optAttr, paramAttr := noun.GetAttr()
	addPropertyToElement(el, noun.GetText(), optAttr, paramAttr)
	return el
}

func CreateDocument() (*etree.Document, *etree.Element) {
	doc := etree.NewDocument()
	doc.CreateProcInst("xml", `version="1.0" encoding="UTF-8"`)
	element := doc.CreateElement("Response")
	return doc, element
}

func ToXML(document *etree.Document) string {
	xml, err := document.WriteToString()
	if err == nil {
		return xml
	}
	return err.Error()
}

func addPropertyToElement(treeElement *etree.Element, text string, optAttr map[string]string, paramAttr map[string]string) {
	if text != "" {
		treeElement.SetText(text)
	}
	for k, v := range paramAttr {
		if v != "" {
			treeElement.CreateAttr(formatAttrKey(k), v)
		}
	}
	for k, v := range optAttr {
		treeElement.CreateAttr(formatAttrKey(k), v)
	}
}

func formatAttrKey(s string) string {
	return strings.ToLower(string(s[0])) + s[1:]
}
