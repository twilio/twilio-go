# Twilio Content API Examples

This document provides comprehensive examples for using Twilio's Content API with different content types. The Content API allows you to create, manage, and send structured content templates with various interactive elements.

## Basic Setup

```go
package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/twilio/twilio-go"
	content "github.com/twilio/twilio-go/rest/content/v1"
)

func main() {
	accountSid := os.Getenv("TWILIO_ACCOUNT_SID")
	authToken := os.Getenv("TWILIO_AUTH_TOKEN")

	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: accountSid,
		Password: authToken,
	})

	// Examples of different content types
	createTextContent(client)
	createMediaContent(client)
	createCardContent(client)
	createListPickerContent(client)
	createCallToActionContent(client)
	createQuickReplyContent(client)
	createCarouselContent(client)
	createLocationContent(client)
	
	// Examples of other operations
	listContent(client)
	fetchContent(client, "your-content-sid") // Replace with actual SID
}
```

## 1. Text Content (twilio/text)

Simple text-based content for basic messaging.

```go
func createTextContent(client *twilio.RestClient) {
	fmt.Println("Creating Text Content...")
	
	twilioText := &content.TwilioText{
		Body: "Hello! Welcome to our service. We're excited to help you today.",
	}
	
	contentType := &content.Types{TwilioText: twilioText}
	createContentRequest := &content.ContentCreateRequest{
		FriendlyName: "Welcome Message",
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
		fmt.Println("Error creating text content: " + err.Error())
	} else {
		response, _ := json.Marshal(*resp)
		fmt.Println("Text Content Response: " + string(response))
	}
}
```

## 2. Media Content (twilio/media)

Content with image, video, or document attachments.

```go
func createMediaContent(client *twilio.RestClient) {
	fmt.Println("Creating Media Content...")
	
	twilioMedia := &content.TwilioMedia{
		Body: "Check out our latest product catalog!",
		Media: []string{
			"https://example.com/product-image.jpg",
			"https://example.com/product-video.mp4",
		},
	}
	
	contentType := &content.Types{TwilioMedia: twilioMedia}
	createContentRequest := &content.ContentCreateRequest{
		FriendlyName: "Product Catalog",
		Language:     "en", 
		Types:        *contentType,
	}
	
	params := &content.CreateContentParams{}
	params.SetContentCreateRequest(*createContentRequest)

	resp, err := client.ContentV1.CreateContent(params)
	if err != nil {
		fmt.Println("Error creating media content: " + err.Error())
	} else {
		response, _ := json.Marshal(*resp)
		fmt.Println("Media Content Response: " + string(response))
	}
}
```

## 3. Card Content (twilio/card)

Structured card with title, subtitle, media, and action buttons.

```go
func createCardContent(client *twilio.RestClient) {
	fmt.Println("Creating Card Content...")
	
	cardActions := []content.CardAction{
		{
			Type:  content.CARDACTIONTYPE_URL,
			Title: "View Details",
			Url:   "https://example.com/product/123",
		},
		{
			Type:  content.CARDACTIONTYPE_PHONE_NUMBER,
			Title: "Call Support",
			Phone: "+1234567890",
		},
	}
	
	twilioCard := &content.TwilioCard{
		Title:    "Premium Wireless Headphones",
		Subtitle: "High-quality sound with noise cancellation",
		Media: []string{
			"https://example.com/headphones.jpg",
		},
		Actions: cardActions,
	}
	
	contentType := &content.Types{TwilioCard: twilioCard}
	createContentRequest := &content.ContentCreateRequest{
		FriendlyName: "Product Card",
		Language:     "en",
		Types:        *contentType,
	}
	
	params := &content.CreateContentParams{}
	params.SetContentCreateRequest(*createContentRequest)

	resp, err := client.ContentV1.CreateContent(params)
	if err != nil {
		fmt.Println("Error creating card content: " + err.Error())
	} else {
		response, _ := json.Marshal(*resp)
		fmt.Println("Card Content Response: " + string(response))
	}
}
```

## 4. List Picker Content (twilio/list-picker)

Interactive menu with selectable options.

```go
func createListPickerContent(client *twilio.RestClient) {
	fmt.Println("Creating List Picker Content...")
	
	listItems := []content.ListItem{
		{
			Id:          "support",
			Item:        "Customer Support",
			Description: "Get help with your account",
		},
		{
			Id:          "billing",
			Item:        "Billing Information", 
			Description: "View and manage your billing",
		},
		{
			Id:          "products",
			Item:        "Our Products",
			Description: "Browse our product catalog",
		},
	}
	
	twilioListPicker := &content.TwilioListPicker{
		Body:   "How can we help you today? Please select an option:",
		Button: "Select Option",
		Items:  listItems,
	}
	
	contentType := &content.Types{TwilioListPicker: twilioListPicker}
	createContentRequest := &content.ContentCreateRequest{
		FriendlyName: "Help Menu",
		Language:     "en",
		Types:        *contentType,
	}
	
	params := &content.CreateContentParams{}
	params.SetContentCreateRequest(*createContentRequest)

	resp, err := client.ContentV1.CreateContent(params)
	if err != nil {
		fmt.Println("Error creating list picker content: " + err.Error())
	} else {
		response, _ := json.Marshal(*resp)
		fmt.Println("List Picker Content Response: " + string(response))
	}
}
```

## 5. Call to Action Content (twilio/call-to-action)

Content with action buttons for URLs, phone calls, etc.

```go
func createCallToActionContent(client *twilio.RestClient) {
	fmt.Println("Creating Call to Action Content...")
	
	ctaActions := []content.CallToActionAction{
		{
			Type:  content.CALLTOACTIONACTIONTYPE_URL,
			Title: "Visit Website",
			Url:   "https://example.com",
		},
		{
			Type:  content.CALLTOACTIONACTIONTYPE_PHONE_NUMBER,
			Title: "Call Now",
			Phone: "+1234567890",
		},
		{
			Type:  content.CALLTOACTIONACTIONTYPE_COPY_CODE,
			Title: "Copy Discount Code",
			Code:  "SAVE20",
		},
	}
	
	twilioCallToAction := &content.TwilioCallToAction{
		Body:    "Special offer just for you! Get 20% off your next purchase.",
		Actions: ctaActions,
	}
	
	contentType := &content.Types{TwilioCallToAction: twilioCallToAction}
	createContentRequest := &content.ContentCreateRequest{
		FriendlyName: "Discount Offer",
		Language:     "en",
		Types:        *contentType,
	}
	
	params := &content.CreateContentParams{}
	params.SetContentCreateRequest(*createContentRequest)

	resp, err := client.ContentV1.CreateContent(params)
	if err != nil {
		fmt.Println("Error creating call to action content: " + err.Error())
	} else {
		response, _ := json.Marshal(*resp)
		fmt.Println("Call to Action Content Response: " + string(response))
	}
}
```

## 6. Quick Reply Content (twilio/quick-reply)

Content with quick reply buttons for easy user responses.

```go
func createQuickReplyContent(client *twilio.RestClient) {
	fmt.Println("Creating Quick Reply Content...")
	
	quickReplyActions := []content.QuickReplyAction{
		{
			Type:  content.QUICKREPLYACTIONTYPE_QUICK_REPLY,
			Title: "Yes",
			Id:    "yes",
		},
		{
			Type:  content.QUICKREPLYACTIONTYPE_QUICK_REPLY,
			Title: "No",
			Id:    "no",
		},
		{
			Type:  content.QUICKREPLYACTIONTYPE_QUICK_REPLY,
			Title: "Maybe",
			Id:    "maybe",
		},
	}
	
	twilioQuickReply := &content.TwilioQuickReply{
		Body:    "Are you interested in receiving updates about our new products?",
		Actions: quickReplyActions,
	}
	
	contentType := &content.Types{TwilioQuickReply: twilioQuickReply}
	createContentRequest := &content.ContentCreateRequest{
		FriendlyName: "Product Updates Survey",
		Language:     "en",
		Types:        *contentType,
	}
	
	params := &content.CreateContentParams{}
	params.SetContentCreateRequest(*createContentRequest)

	resp, err := client.ContentV1.CreateContent(params)
	if err != nil {
		fmt.Println("Error creating quick reply content: " + err.Error())
	} else {
		response, _ := json.Marshal(*resp)
		fmt.Println("Quick Reply Content Response: " + string(response))
	}
}
```

## 7. Carousel Content (twilio/carousel)

Multiple cards in a horizontally scrollable format.

```go
func createCarouselContent(client *twilio.RestClient) {
	fmt.Println("Creating Carousel Content...")
	
	carouselCards := []content.CarouselCard{
		{
			Title: "Premium Headphones",
			Body:  "High-quality audio experience",
			Media: "https://example.com/headphones.jpg",
			Actions: []content.CarouselAction{
				{
					Type:  content.CAROUSELACTIONTYPE_URL,
					Title: "View Details",
					Url:   "https://example.com/product/headphones",
				},
			},
		},
		{
			Title: "Wireless Speaker",
			Body:  "Portable sound for everywhere",
			Media: "https://example.com/speaker.jpg",
			Actions: []content.CarouselAction{
				{
					Type:  content.CAROUSELACTIONTYPE_URL,
					Title: "Learn More",
					Url:   "https://example.com/product/speaker",
				},
			},
		},
		{
			Title: "Smart Watch",
			Body:  "Stay connected on the go",
			Media: "https://example.com/watch.jpg",
			Actions: []content.CarouselAction{
				{
					Type:  content.CAROUSELACTIONTYPE_QUICK_REPLY,
					Title: "Add to Cart",
					Id:    "add_watch",
				},
			},
		},
	}
	
	twilioCarousel := &content.TwilioCarousel{
		Body:  "Check out our featured products:",
		Cards: carouselCards,
	}
	
	contentType := &content.Types{TwilioCarousel: twilioCarousel}
	createContentRequest := &content.ContentCreateRequest{
		FriendlyName: "Featured Products Carousel",
		Language:     "en",
		Types:        *contentType,
	}
	
	params := &content.CreateContentParams{}
	params.SetContentCreateRequest(*createContentRequest)

	resp, err := client.ContentV1.CreateContent(params)
	if err != nil {
		fmt.Println("Error creating carousel content: " + err.Error())
	} else {
		response, _ := json.Marshal(*resp)
		fmt.Println("Carousel Content Response: " + string(response))
	}
}
```

## 8. Location Content (twilio/location)

Share location information with recipients.

```go
func createLocationContent(client *twilio.RestClient) {
	fmt.Println("Creating Location Content...")
	
	twilioLocation := &content.TwilioLocation{
		Latitude:  37.7749,  // San Francisco latitude
		Longitude: -122.4194, // San Francisco longitude
		Label:     "Twilio Headquarters",
		Address:   "375 Beale St, San Francisco, CA 94105",
		Id:        "twilio-hq",
	}
	
	contentType := &content.Types{TwilioLocation: twilioLocation}
	createContentRequest := &content.ContentCreateRequest{
		FriendlyName: "Office Location",
		Language:     "en",
		Types:        *contentType,
	}
	
	params := &content.CreateContentParams{}
	params.SetContentCreateRequest(*createContentRequest)

	resp, err := client.ContentV1.CreateContent(params)
	if err != nil {
		fmt.Println("Error creating location content: " + err.Error())
	} else {
		response, _ := json.Marshal(*resp)
		fmt.Println("Location Content Response: " + string(response))
	}
}
```

## 9. List Content

Retrieve all content templates for your account.

```go
func listContent(client *twilio.RestClient) {
	fmt.Println("Listing Content...")
	
	params := &content.ListContentParams{
		PageSize: &[]int{20}[0], // Get 20 items per page
		Limit:    &[]int{100}[0], // Maximum 100 items total
	}

	resp, err := client.ContentV1.ListContent(params)
	if err != nil {
		fmt.Println("Error listing content: " + err.Error())
	} else {
		fmt.Printf("Found %d content items\n", len(resp))
		for i, content := range resp {
			fmt.Printf("Content %d: SID=%s, Name=%s\n", 
				i+1, 
				*content.Sid, 
				*content.FriendlyName)
		}
	}
}
```

## 10. Fetch Content

Retrieve a specific content template by its SID.

```go
func fetchContent(client *twilio.RestClient, contentSid string) {
	fmt.Printf("Fetching Content with SID: %s\n", contentSid)
	
	resp, err := client.ContentV1.FetchContent(contentSid)
	if err != nil {
		fmt.Println("Error fetching content: " + err.Error())
	} else {
		response, _ := json.Marshal(*resp)
		fmt.Println("Fetched Content: " + string(response))
	}
}
```

## Content Types Overview

The Content API supports various content types, each with specific use cases:

- **twilio/text**: Plain text messages
- **twilio/media**: Messages with media attachments (images, videos, documents)
- **twilio/card**: Rich cards with titles, subtitles, media, and action buttons
- **twilio/list-picker**: Interactive menus with selectable options
- **twilio/call-to-action**: Messages with action buttons (URLs, phone calls, etc.)
- **twilio/quick-reply**: Messages with quick response buttons
- **twilio/catalog**: Product catalog displays
- **twilio/carousel**: Multiple cards in a scrollable format
- **twilio/location**: Location sharing
- **whatsapp/card**: WhatsApp-specific card format
- **whatsapp/authentication**: WhatsApp authentication templates

## Variables and Personalization

You can use variables in your content to personalize messages:

```go
Variables: map[string]string{
    "1": "customer_name",
    "2": "order_number",
    "3": "delivery_date",
}
```

These variables can be replaced with actual values when sending the content.

## Error Handling

Always implement proper error handling when working with the Content API:

```go
resp, err := client.ContentV1.CreateContent(params)
if err != nil {
    // Handle specific error types
    if strings.Contains(err.Error(), "authentication") {
        fmt.Println("Authentication failed")
    } else if strings.Contains(err.Error(), "validation") {
        fmt.Println("Content validation failed")
    } else {
        fmt.Println("Unexpected error: " + err.Error())
    }
    return
}
```
