package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"reflect"
	"regexp"
)

//Takes a limit on the max number of records to read and a max pageSize and calculates the max number of pages to read.
func ReadLimits(pageSize *int, limit int) int {
	//don't care about pageSize
	if pageSize == nil {
		if limit == 0 {
			//don't care about the limit either
			return 50 //default
		}
		//return the most efficient pageSize
		return min(limit, 1000)
	} else {
		if limit == 0 {
			//we care about the pageSize but not the limit
			return *pageSize
		}
		return min(*pageSize, limit)
	}
}

func GetNext(response interface{}, curRecord *int, limit int, getNextPage func(nextPageUri string) (interface{}, error)) (interface{}, error) {
	nextPageUri, err := getNextPageUri(response, curRecord, limit)
	if err != nil {
		return nil, err
	}

	return getNextPage(nextPageUri)
}

func GetPayload(response interface{}) ([]interface{}, string, error) {
	payload := toMap(response)
	var data [][]interface{}
	for _, v := range payload {
		if v != nil {
			kind := reflect.TypeOf(v).Kind()
			switch kind {
			//look for non metadata info
			case reflect.Slice:
				if len(data) > 0 {
					//we expect this to be exactly 1
					return nil, "", errors.New("payload contains more than 1 record of type array")
				}
				data = append(data, v.([]interface{}))
			}
		}
	}

	if len(data) == 1 {
		return data[0], getNextPageUrl(payload), nil
	}
	return nil, "", errors.New("could not retrieve payload from response")
}

func toMap(s interface{}) map[string]interface{} {
	var payload map[string]interface{}
	test, err := json.Marshal(s)
	if err != nil {
		log.Print("Map creation error: ", err)
		return nil
	}
	_ = json.Unmarshal(test, &payload)
	return payload
}

func getNextPageUri(response interface{}, curRecord *int, limit int) (string, error) {
	//get just the non metadata info and the next page uri
	payload, nextPageUri, err := GetPayload(response)
	if err != nil {
		return "", err
	}

	*curRecord += len(payload)

	if limit != 0 {
		//we have reached the desired limit
		if limit <= *curRecord {
			return "", nil
		}

		remaining := limit - *curRecord
		if remaining > 0 {
			pageSize := min(len(payload), remaining)
			re := regexp.MustCompile(`PageSize=\d+`)
			nextPageUri = re.ReplaceAllString(nextPageUri, fmt.Sprintf("PageSize=%d", pageSize))
		}
	}

	return nextPageUri, err
}

func getNextPageUrl(payload map[string]interface{}) string {
	if payload != nil && payload["meta"] != nil && payload["meta"].(map[string]interface{})["next_page_url"] != nil {
		return payload["meta"].(map[string]interface{})["next_page_url"].(string)
	}

	if payload != nil && payload["next_page_uri"] != nil {
		return payload["next_page_uri"].(string)
	}

	return ""
}

func min(a int, b int) int {
	if a > b {
		return b
	}
	return a
}
