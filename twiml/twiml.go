package twiml

import (
	"github.com/beevik/etree"
	"strings"
)

type Element interface {
	GetName() string
	GetText() string
	GetAttr() (map[string]string, map[string]string)
	GetNouns() []Element
}

func AddAllVerbs(response *etree.Element, verbs []Element) {
	for _, verb := range verbs {
		verbEl := createElement(verb)
		response.AddChild(verbEl)
	}
}

func createElement(tag Element) *etree.Element {
	el := etree.NewElement(tag.GetName())
	optAttr, paramAttr := tag.GetAttr()
	addPropertyToElement(el, tag.GetText(), optAttr, paramAttr)
	//Loop through all Nouns
	if len(tag.GetNouns()) != 0 {
		for _, noun := range tag.GetNouns() {
			child := createElement(noun)
			el.AddChild(child)
		}
	}
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

	for _, attr := range []map[string]string{paramAttr, optAttr} {
		for k, v := range attr {
			if v != "" {
				treeElement.CreateAttr(formatAttrKey(k), v)
			}
		}
	}
}

func formatAttrKey(s string) string {
	return strings.ToLower(string(s[0])) + s[1:]
}
