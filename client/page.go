package client

import (
	"encoding/json"
	"net/http"
	"reflect"
)

//Represents a page of records in a collection
type Page struct {
	Payload         map[string]interface{}
	NextPageUrl     string
	PreviousPageUrl string
	Records         []interface{}
	BaseUrl         string
}

func NewPage(baseUrl string, response *http.Response) *Page {
	respBody := processResponse(response)
	var page = Page{
		Payload: respBody,
		BaseUrl: baseUrl,
	}

	page.Records = page.loadPage()
	page.NextPageUrl = page.getNextPageUrl()
	page.PreviousPageUrl = page.getPreviousPageUrl()
	return &page
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
	var data [][]interface{}
	for _, v := range p.Payload {
		if v != nil {
			kind := reflect.TypeOf(v).Kind()
			switch kind {
			//look for non metadata info
			case reflect.Slice:
				if len(data) > 0 {
					//we expect this to be exactly 1
					return nil
				}
				data = append(data, v.([]interface{}))
			}
		}
	}

	if len(data) == 1 {
		return data[0]
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
