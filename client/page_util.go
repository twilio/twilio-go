package client

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Takes a limit on the max number of records to read and a max pageSize and calculates the max number of pages to read.
func ReadLimits(pageSize *int, limit *int) int {
	//don't care about pageSize
	if pageSize == nil {
		if limit == nil {
			//don't care about the limit either
			return 50 //default
		}
		//return the most efficient pageSize
		return min(*limit, 1000)
	} else {
		if limit == nil {
			//we care about the pageSize but not the limit
			return *pageSize
		}
		return min(*pageSize, *limit)
	}
}

func GetNext(baseUrl string, response interface{}, getNextPage func(nextPageUri string) (interface{}, error)) (interface{}, error) {
	nextPageUrl, err := getNextPageUrl(baseUrl, response)
	if err != nil {
		return nil, err
	}

	return getNextPage(nextPageUrl)
}

// GetPrevious fetches the previous page using the provided callback function.
// Supports URL-based, URI-based, and token-based pagination.
func GetPrevious(baseUrl string, response interface{}, getPreviousPage func(previousPageUrl string) (interface{}, error)) (interface{}, error) {
	previousPageUrl, err := getPreviousPageUrl(baseUrl, response)
	if err != nil {
		return nil, err
	}

	return getPreviousPage(previousPageUrl)
}

func toMap(s interface{}) (map[string]interface{}, error) {
	var payload map[string]interface{}
	data, err := json.Marshal(s)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &payload)
	if err != nil {
		return nil, err
	}

	return payload, err
}

func getNextPageUrl(baseUrl string, response interface{}) (string, error) {
	payload, err := toMap(response)
	if err != nil {
		return "", err
	}

	// Try URL-based pagination first (meta.next_page_url)
	if nextPageUrl, err := getStringFromMap(payload, "meta", "next_page_url"); err == nil && nextPageUrl != "" {
		return nextPageUrl, nil
	}

	// Try URI-based pagination (next_page_uri)
	if nextPageUri, err := getStringFromMap(payload, "next_page_uri"); err == nil && nextPageUri != "" {
		// remove any leading and trailing '/'
		return fmt.Sprintf("%s/%s", strings.Trim(baseUrl, "/"), strings.Trim(nextPageUri, "/")), nil
	}

	// Try token-based pagination (meta.next_token)
	tokenMeta, err := extractTokenMeta(payload)
	if err != nil {
		return "", err
	}
	if tokenMeta != nil && tokenMeta.NextToken != "" {
		return buildTokenUrl(baseUrl, tokenMeta.NextToken, tokenMeta.PageSize), nil
	}

	return "", nil
}

// getPreviousPageUrl extracts the previous page URL from a response.
// Supports URL-based, URI-based, and token-based pagination.
func getPreviousPageUrl(baseUrl string, response interface{}) (string, error) {
	payload, err := toMap(response)
	if err != nil {
		return "", err
	}

	// Try URL-based pagination first (meta.previous_page_url)
	if prevPageUrl, err := getStringFromMap(payload, "meta", "previous_page_url"); err == nil && prevPageUrl != "" {
		return prevPageUrl, nil
	}

	// Try URI-based pagination (previous_page_uri)
	if prevPageUri, err := getStringFromMap(payload, "previous_page_uri"); err == nil && prevPageUri != "" {
		// remove any leading and trailing '/'
		return fmt.Sprintf("%s/%s", strings.Trim(baseUrl, "/"), strings.Trim(prevPageUri, "/")), nil
	}

	// Try token-based pagination (meta.previous_token)
	tokenMeta, err := extractTokenMeta(payload)
	if err != nil {
		return "", err
	}
	if tokenMeta != nil && tokenMeta.PreviousToken != "" {
		return buildTokenUrl(baseUrl, tokenMeta.PreviousToken, tokenMeta.PageSize), nil
	}

	return "", nil
}

// getStringFromMap safely extracts a string value from nested maps.
// The path parameter specifies the nested key path (e.g., "meta", "next_token").
// Returns empty string and error if the path doesn't exist or type is wrong.
func getStringFromMap(data map[string]interface{}, path ...string) (string, error) {
	if len(path) == 0 {
		return "", fmt.Errorf("path cannot be empty")
	}

	var current interface{} = data
	for i, key := range path {
		m, ok := current.(map[string]interface{})
		if !ok {
			return "", fmt.Errorf("expected map at path index %d, got %T", i, current)
		}

		value, exists := m[key]
		if !exists {
			return "", nil // Return empty string if key doesn't exist
		}

		if i == len(path)-1 {
			// Last key in path - extract string value
			str, ok := value.(string)
			if !ok {
				return "", fmt.Errorf("expected string at path %v, got %T", path, value)
			}
			return str, nil
		}

		current = value
	}

	return "", fmt.Errorf("unexpected end of path")
}

// getIntFromMap safely extracts an int value from nested maps.
// The path parameter specifies the nested key path (e.g., "meta", "page_size").
// Handles both int and float64 (from JSON unmarshaling).
// Returns 0 and error if the path doesn't exist or type is wrong.
func getIntFromMap(data map[string]interface{}, path ...string) (int, error) {
	if len(path) == 0 {
		return 0, fmt.Errorf("path cannot be empty")
	}

	var current interface{} = data
	for i, key := range path {
		m, ok := current.(map[string]interface{})
		if !ok {
			return 0, fmt.Errorf("expected map at path index %d, got %T", i, current)
		}

		value, exists := m[key]
		if !exists {
			return 0, nil // Return 0 if key doesn't exist
		}

		if i == len(path)-1 {
			// Last key in path - extract int value
			switch v := value.(type) {
			case int:
				return v, nil
			case float64:
				return int(v), nil
			case int64:
				return int(v), nil
			default:
				return 0, fmt.Errorf("expected int or float64 at path %v, got %T", path, value)
			}
		}

		current = value
	}

	return 0, fmt.Errorf("unexpected end of path")
}

// TokenMetadata holds token-based pagination information extracted from API responses.
// This supports the new Twilio API Standards V1 pagination pattern.
type TokenMetadata struct {
	NextToken     string
	PreviousToken string
	PageSize      int
	Key           string
}

// extractTokenMeta extracts token pagination metadata from a response payload.
// Supports both snake_case (next_token) and camelCase (nextToken) field names
// to accommodate different API naming conventions.
// Returns nil if no meta field exists or if token fields are not present.
func extractTokenMeta(payload map[string]interface{}) (*TokenMetadata, error) {
	if payload == nil {
		return nil, nil
	}

	// Check if meta field exists
	metaInterface, exists := payload["meta"]
	if !exists {
		return nil, nil
	}

	_, ok := metaInterface.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("meta field is not a map")
	}

	metadata := &TokenMetadata{}

	// Try extracting next_token (snake_case) first, then nextToken (camelCase)
	if nextToken, err := getStringFromMap(payload, "meta", "next_token"); err == nil && nextToken != "" {
		metadata.NextToken = nextToken
	} else if nextToken, err := getStringFromMap(payload, "meta", "nextToken"); err == nil && nextToken != "" {
		metadata.NextToken = nextToken
	}

	// Try extracting previous_token (snake_case) first, then previousToken (camelCase)
	if prevToken, err := getStringFromMap(payload, "meta", "previous_token"); err == nil && prevToken != "" {
		metadata.PreviousToken = prevToken
	} else if prevToken, err := getStringFromMap(payload, "meta", "previousToken"); err == nil && prevToken != "" {
		metadata.PreviousToken = prevToken
	}

	// Try extracting page_size (snake_case) first, then pageSize (camelCase)
	if pageSize, err := getIntFromMap(payload, "meta", "page_size"); err == nil && pageSize > 0 {
		metadata.PageSize = pageSize
	} else if pageSize, err := getIntFromMap(payload, "meta", "pageSize"); err == nil && pageSize > 0 {
		metadata.PageSize = pageSize
	}

	// Try extracting key field
	if key, err := getStringFromMap(payload, "meta", "key"); err == nil && key != "" {
		metadata.Key = key
	}

	// Return nil if no token fields were found
	if metadata.NextToken == "" && metadata.PreviousToken == "" {
		return nil, nil
	}

	return metadata, nil
}

// buildTokenUrl constructs a URL with pageToken and pageSize query parameters.
// Returns baseUrl?PageSize=X&PageToken=Y format.
// Preserves existing query parameters if present in baseUrl.
// Returns empty string if pageToken is empty.
func buildTokenUrl(baseUrl string, pageToken string, pageSize int) string {
	if pageToken == "" {
		return ""
	}

	params := []string{}
	if pageSize > 0 {
		params = append(params, fmt.Sprintf("PageSize=%d", pageSize))
	}
	if pageToken != "" {
		params = append(params, fmt.Sprintf("PageToken=%s", pageToken))
	}

	if len(params) == 0 {
		return ""
	}

	// Check if baseUrl already has query parameters
	separator := "?"
	if strings.Contains(baseUrl, "?") {
		separator = "&"
	}

	return fmt.Sprintf("%s%s%s", baseUrl, separator, strings.Join(params, "&"))
}

func min(a int, b int) int {
	if a > b {
		return b
	}
	return a
}
