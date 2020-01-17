package twilio

type AvailablePhoneNumberLocal struct {
	FriendlyName        string          `json:"friendly_name,omitempty"`
	PhoneNumber         string          `json:"phone_number,omitempty"`
	Lata                string          `json:"lata,omitempty"`
	Locality            string          `json:"locality,omitempty"`
	RateCenter          string          `json:"rate_center,omitempty"`
	Latitude            float32         `json:"latitude,omitempty"`
	Longitude           float32         `json:"longitude,omitempty"`
	Region              string          `json:"region,omitempty"`
	PostalCode          string          `json:"postal_code,omitempty"`
	IsoCountry          int             `json:"iso_country,omitempty"`
	AddressRequirements string          `json:"address_requirements,omitempty"`
	Beta                bool            `json:"beta,omitempty"`
	Capabilities        map[string]bool `json:"capabilities,omitempty"`
}

type AvailablePhoneNumberLocalReadParams struct {
	// Read on AvailablePhoneNumberLocal requires account_sid in URI
	// Read on AvailablePhoneNumberLocal requires country_code in URI
	AreaCode                      int    `url:"AreaCode,omitempty"`
	Contains                      string `url:"Contains,omitempty"`
	SmsEnabled                    bool   `url:"SmsEnabled,omitempty"`
	MmsEnabled                    bool   `url:"MmsEnabled,omitempty"`
	VoiceEnabled                  bool   `url:"VoiceEnabled,omitempty"`
	ExcludeAllAddressRequired     bool   `url:"ExcludeAllAddressRequired,omitempty"`
	ExcludeLocalAddressRequired   bool   `url:"ExcludeLocalAddressRequired,omitempty"`
	ExcludeForeignAddressRequired bool   `url:"ExcludeForeignAddressRequired,omitempty"`
	Beta                          bool   `url:"Beta,omitempty"`
	NearNumber                    string `url:"NearNumber,omitempty"`
	NearLatLong                   string `url:"NearLatLong,omitempty"`
	Distance                      int    `url:"Distance,omitempty"`
	InPostalCode                  string `url:"InPostalCode,omitempty"`
	InRegion                      string `url:"InRegion,omitempty"`
	InRateCenter                  string `url:"InRateCenter,omitempty"`
	InLata                        string `url:"InLata,omitempty"`
	InLocality                    string `url:"InLocality,omitempty"`
	FaxEnabled                    bool   `url:"FaxEnabled,omitempty"`
}
