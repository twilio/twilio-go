package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/twilio/twilio-go"
	content "github.com/twilio/twilio-go/rest/content/v1"
)

// Simple example showing the most common Content API usage patterns
func main() {
	accountSid := os.Getenv("TWILIO_ACCOUNT_SID")
	authToken := os.Getenv("TWILIO_AUTH_TOKEN")

	if accountSid == "" || authToken == "" {
		fmt.Println("Please set TWILIO_ACCOUNT_SID and TWILIO_AUTH_TOKEN environment variables")
		fmt.Println("Example usage:")
		fmt.Println("  export TWILIO_ACCOUNT_SID=your_account_sid")
		fmt.Println("  export TWILIO_AUTH_TOKEN=your_auth_token")
		fmt.Println("  go run content_api_example.go")
		return
	}

	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: accountSid,
		Password: authToken,
	})

	fmt.Println("=== Twilio Content API Examples ===\n")

	// Example 1: Create simple text content
	fmt.Println("1. Creating text content...")
	createSimpleTextContent(client)

	// Example 2: Create media content  
	fmt.Println("\n2. Creating media content...")
	createSimpleMediaContent(client)

	// Example 3: Create interactive card
	fmt.Println("\n3. Creating card content...")
	createSimpleCardContent(client)

	// Example 4: List existing content
	fmt.Println("\n4. Listing content...")
	listExistingContent(client)

	fmt.Println("\n=== Examples completed ===")
	fmt.Println("For more advanced examples, see: advanced-examples/content.md")
}

func createSimpleTextContent(client *twilio.RestClient) {
	twilioText := &content.TwilioText{
		Body: "Hello {{1}}! Welcome to our service.",
	}

	contentType := &content.Types{TwilioText: twilioText}
	createContentRequest := &content.ContentCreateRequest{
		FriendlyName: "Simple Welcome Message",
		Language:     "en",
		Types:        *contentType,
		Variables: map[string]string{
			"1": "customer_name",
		},
	}

	params := &content.CreateContentParams{}
	params.SetContentCreateRequest(*createContentRequest)

	resp, err := client.ContentV1.CreateContent(params)
	if err != nil {
		fmt.Printf("   Error: %s\n", err.Error())
	} else {
		fmt.Printf("   ✓ Created content with SID: %s\n", *resp.Sid)
		fmt.Printf("   ✓ Friendly name: %s\n", *resp.FriendlyName)
	}
}

func createSimpleMediaContent(client *twilio.RestClient) {
	twilioMedia := &content.TwilioMedia{
		Body: "Check out this image!",
		Media: []string{
			"https://picsum.photos/800/600", // Random placeholder image
		},
	}

	contentType := &content.Types{TwilioMedia: twilioMedia}
	createContentRequest := &content.ContentCreateRequest{
		FriendlyName: "Simple Media Content",
		Language:     "en",
		Types:        *contentType,
	}

	params := &content.CreateContentParams{}
	params.SetContentCreateRequest(*createContentRequest)

	resp, err := client.ContentV1.CreateContent(params)
	if err != nil {
		fmt.Printf("   Error: %s\n", err.Error())
	} else {
		fmt.Printf("   ✓ Created media content with SID: %s\n", *resp.Sid)
		fmt.Printf("   ✓ Friendly name: %s\n", *resp.FriendlyName)
	}
}

func createSimpleCardContent(client *twilio.RestClient) {
	cardActions := []content.CardAction{
		{
			Type:  content.CARDACTIONTYPE_URL,
			Title: "Learn More",
			Url:   "https://www.twilio.com",
		},
	}

	twilioCard := &content.TwilioCard{
		Title:   "Twilio Communications",
		Subtitle: "APIs for SMS, Voice, Video & more",
		Actions: cardActions,
	}

	contentType := &content.Types{TwilioCard: twilioCard}
	createContentRequest := &content.ContentCreateRequest{
		FriendlyName: "Simple Card Content",
		Language:     "en",
		Types:        *contentType,
	}

	params := &content.CreateContentParams{}
	params.SetContentCreateRequest(*createContentRequest)

	resp, err := client.ContentV1.CreateContent(params)
	if err != nil {
		fmt.Printf("   Error: %s\n", err.Error())
	} else {
		fmt.Printf("   ✓ Created card content with SID: %s\n", *resp.Sid)
		fmt.Printf("   ✓ Friendly name: %s\n", *resp.FriendlyName)
	}
}

func listExistingContent(client *twilio.RestClient) {
	params := &content.ListContentParams{
		PageSize: &[]int{5}[0], // Get just 5 items for this example
	}

	resp, err := client.ContentV1.ListContent(params)
	if err != nil {
		fmt.Printf("   Error: %s\n", err.Error())
	} else {
		fmt.Printf("   ✓ Found %d content items\n", len(resp))
		for i, content := range resp {
			fmt.Printf("   %d. %s (SID: %s)\n", i+1, *content.FriendlyName, *content.Sid)
		}
		
		if len(resp) > 0 {
			// Show details of the first content item
			fmt.Printf("\n   Example content details:\n")
			firstContent := resp[0]
			contentJson, _ := json.MarshalIndent(firstContent, "   ", "  ")
			fmt.Printf("   %s\n", string(contentJson))
		}
	}
}