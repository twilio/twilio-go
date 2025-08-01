/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Iam
 * This is the public Twilio REST API.
 *
 * NOTE: This class is auto generated by OpenAPI Generator.
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

package openapi

import (
	"time"
)

// IamV1Vendoroauthapp struct for IamV1Vendoroauthapp
type IamV1Vendoroauthapp struct {
	Type             string                    `json:"type,omitempty"`
	Sid              string                    `json:"sid,omitempty"`
	FriendlyName     string                    `json:"friendly_name,omitempty"`
	Description      string                    `json:"description,omitempty"`
	DateCreated      time.Time                 `json:"date_created,omitempty"`
	Status           string                    `json:"status,omitempty"`
	Policy           IamV1VendoroauthappPolicy `json:"policy,omitempty"`
	CreatedBy        string                    `json:"created_by,omitempty"`
	CompanyName      string                    `json:"company_name,omitempty"`
	HomepageUrl      string                    `json:"homepage_url,omitempty"`
	TosUrl           string                    `json:"tos_url,omitempty"`
	RedirectUrls     []string                  `json:"redirect_urls,omitempty"`
	ImageUrl         string                    `json:"image_url,omitempty"`
	AuthorizationUrl string                    `json:"authorization_url,omitempty"`
	AccessTokenTtl   int                       `json:"access_token_ttl,omitempty"`
}
