# Content API Examples

This directory contains comprehensive examples for using Twilio's Content API with the twilio-go SDK.

## Quick Start

### 1. Simple Runnable Example

For a quick introduction, run the standalone example:

```bash
export TWILIO_ACCOUNT_SID=your_account_sid
export TWILIO_AUTH_TOKEN=your_auth_token
go run content_api_example.go
```

This example shows the basics:
- Creating text, media, and card content
- Listing existing content
- Handling responses and errors

### 2. Comprehensive Examples

See [content.md](content.md) for detailed examples covering all content types:

- **twilio/text** - Simple text messages
- **twilio/media** - Messages with images/videos/documents  
- **twilio/card** - Rich cards with action buttons
- **twilio/list-picker** - Interactive selection menus
- **twilio/call-to-action** - Action buttons (URLs, phone calls, etc.)
- **twilio/quick-reply** - Quick response buttons
- **twilio/carousel** - Multiple cards in scrollable format
- **twilio/location** - Location sharing

## Content Types Overview

| Content Type | Use Case | Key Features |
|-------------|----------|--------------|
| `twilio/text` | Simple messaging | Plain text with variable substitution |
| `twilio/media` | Rich media | Images, videos, documents + optional text |
| `twilio/card` | Product showcases | Title, subtitle, media, action buttons |
| `twilio/list-picker` | Menu selection | Interactive list with descriptions |
| `twilio/call-to-action` | Action prompts | Multiple action buttons |
| `twilio/quick-reply` | Quick responses | Fast reply options |
| `twilio/carousel` | Product catalogs | Multiple cards in horizontal scroll |
| `twilio/location` | Location sharing | Coordinates with address/label |

## Key Concepts

### Variables and Personalization
```go
Variables: map[string]string{
    "1": "customer_name",
    "2": "order_number",
}
```
Use `{{1}}`, `{{2}}` etc. in your content text to substitute variables.

### Action Types
Content can include interactive elements:
- **URL** - Open web links
- **PHONE_NUMBER** - Initiate phone calls  
- **QUICK_REPLY** - Send predefined responses
- **COPY_CODE** - Copy codes to clipboard

### Error Handling
Always check for errors when making API calls:
```go
resp, err := client.ContentV1.CreateContent(params)
if err != nil {
    // Handle authentication, validation, or network errors
    fmt.Printf("Error: %s\n", err.Error())
    return
}
```

## Authentication

Set your Twilio credentials as environment variables:
```bash
export TWILIO_ACCOUNT_SID=your_account_sid
export TWILIO_AUTH_TOKEN=your_auth_token
```

Or pass them directly:
```go
client := twilio.NewRestClientWithParams(twilio.ClientParams{
    Username: "your_account_sid",
    Password: "your_auth_token",
})
```

## Additional Resources

- [Twilio Content API Documentation](https://www.twilio.com/docs/content)
- [Content Types Overview](https://www.twilio.com/docs/content/content-types-overview)
- [twilio-go SDK Documentation](https://github.com/twilio/twilio-go)