package client

import (
	"encoding/json"
	"net/http"
)

//Represents a page of records in a collection
type Page struct {
	Payload         map[string]interface{}
	NextPageUrl     string
	PreviousPageUrl string
	Records         []interface{}
	MetaKeys        map[string]bool
	BaseUrl         string
}

func NewPage(baseUrl string, response *http.Response) *Page {
	respBody := processResponse(response)
	var page = Page{
		Payload:  respBody,
		MetaKeys: loadMetaKeys(),
		BaseUrl:  baseUrl,
	}

	page.Records = page.loadPage()
	page.NextPageUrl = page.getNextPageUrl()
	page.PreviousPageUrl = page.getPreviousPageUrl()
	return &page
}

func loadMetaKeys() map[string]bool {
	return map[string]bool{
		"end":               true,
		"first_page_uri":    true,
		"last_page_uri":     true,
		"next_page_uri":     true,
		"page":              true,
		"page_size":         true,
		"previous_page_uri": true,
		"total":             true,
		"num_pages":         true,
		"start":             true,
		"uri":               true,
	}
}

func processResponse(response *http.Response) map[string]interface{} {
	var data map[string]interface{}
	if response != nil {
		err := json.NewDecoder(response.Body).Decode(&data)

		if err != nil {
			return nil
		}
	}

	return data
}

//Parses the collection of records out of a map payload and returns the records
func (p *Page) loadPage() []interface{} {
	if p.Payload["meta"] != nil && p.Payload["meta"].(map[string]interface{})["key"] != nil {
		key := p.Payload["meta"].(map[string]interface{})["key"].(string)
		return p.Payload[key].([]interface{})
	} else {
		var difference []string
		for pk := range p.Payload {
			if _, found := p.MetaKeys[pk]; found {
				continue
			}

			difference = append(difference, pk)
		}

		if len(difference) == 1 {
			return p.Payload[difference[0]].([]interface{})
		}
	}

	return nil
}

func (p *Page) getNextPageUrl() string {
	if p.Payload["meta"] != nil && p.Payload["meta"].(map[string]interface{})["next_page_url"] != nil {
		return p.Payload["meta"].(map[string]interface{})["next_page_url"].(string)
	}

	if p.Payload["next_page_uri"] != nil {
		return p.BaseUrl + p.Payload["next_page_uri"].(string)
	}

	return ""
}

func (p *Page) getPreviousPageUrl() string {
	if p.Payload["meta"] != nil && p.Payload["meta"].(map[string]interface{})["previous_page_url"] != nil {
		return p.Payload["meta"].(map[string]interface{})["previous_page_url"].(string)
	}

	if p.Payload["previous_page_uri"] != nil {
		return p.BaseUrl + p.Payload["previous_page_uri"].(string)
	}

	return ""
}

func (p *Page) getNextPage(c *RequestHandler) *Page {
	if p.NextPageUrl != "" {
		resp, _ := c.Get(p.NextPageUrl, nil, nil)
		nextPage := NewPage(p.BaseUrl, resp)
		return nextPage
	}

	return nil
}

func (p *Page) getPreviousPage(c *RequestHandler) *Page { //nolint
	if p.PreviousPageUrl != "" {
		resp, _ := c.Get(p.PreviousPageUrl, nil, nil)
		prevPage := NewPage(p.BaseUrl, resp)
		return prevPage
	}

	return nil
}
