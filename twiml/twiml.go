package twiml

import (
	"strings"

	"github.com/beevik/etree"
)

type Element interface {
	GetName() string
	GetText() string
	GetAttr() (map[string]string, map[string]string)
	GetInnerElements() []Element
}

func AddAllVerbs(response *etree.Element, verbs []Element) {
	for _, verb := range verbs {
		verbEl := createElement(verb)
		response.AddChild(verbEl)
	}
}

func createElement(element Element) *etree.Element {
	el := etree.NewElement(element.GetName())
	optAttr, paramAttr := element.GetAttr()
	addPropertyToElement(el, element.GetText(), optAttr, paramAttr)
	//Loop through all Inner Elements
	if len(element.GetInnerElements()) != 0 {
		for _, innerElement := range element.GetInnerElements() {
			child := createElement(innerElement)
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

func ToXML(document *etree.Document) (string, error) {
	return document.WriteToString()
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
