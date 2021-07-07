package client

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"reflect"
	"regexp"
)

//Represents a page of records in a collection
type Page struct {
	Payload         map[string]interface{}
	NextPageUrl     string
	PreviousPageUrl string
	Records         []interface{}
	BaseUrl         string
}

func NewPage(baseUrl string, response interface{}) *Page {
	var payload map[string]interface{}
	if reflect.TypeOf(response).Kind() != reflect.Map {
		test, err := json.Marshal(response)
		if err != nil {
			log.Print("Page creation error: ", err)
			return nil
		}
		_ = json.Unmarshal(test, &payload)
	} else {
		payload = response.(map[string]interface{})
	}

	var page = Page{
		Payload: payload,
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
	if p.Payload != nil && p.Payload["meta"] != nil && p.Payload["meta"].(map[string]interface{})["next_page_url"] != nil {
		return p.Payload["meta"].(map[string]interface{})["next_page_url"].(string)
	}

	if p.Payload != nil && p.Payload["next_page_uri"] != nil {
		return p.BaseUrl + p.Payload["next_page_uri"].(string)
	}

	return ""
}

func (p *Page) getPreviousPageUrl() string {
	if p.Payload != nil && p.Payload["meta"] != nil && p.Payload["meta"].(map[string]interface{})["previous_page_url"] != nil {
		return p.Payload["meta"].(map[string]interface{})["previous_page_url"].(string)
	}

	if p.Payload != nil && p.Payload["previous_page_uri"] != nil {
		return p.BaseUrl + p.Payload["previous_page_uri"].(string)
	}

	return ""
}

func (p *Page) getNextPage(c *RequestHandler, pageSize *int) *Page {
	if p.NextPageUrl != "" {
		//limit has been reached
		if *pageSize != len(p.Records) {
			//set the PageSize to the expected number of records
			re := regexp.MustCompile(`PageSize=\d+`)
			p.NextPageUrl = re.ReplaceAllString(p.NextPageUrl, fmt.Sprintf("PageSize=%d", *pageSize))

		}
		resp, _ := c.Get(p.NextPageUrl, nil, nil)
		respBody := processResponse(resp)
		nextPage := NewPage(p.BaseUrl, respBody)
		return nextPage
	}

	return nil
}

func (p *Page) getPreviousPage(c *RequestHandler) *Page { //nolint
	if p.PreviousPageUrl != "" {
		resp, _ := c.Get(p.PreviousPageUrl, nil, nil)
		respBody := processResponse(resp)
		prevPage := NewPage(p.BaseUrl, respBody)
		return prevPage
	}

	return nil
}
