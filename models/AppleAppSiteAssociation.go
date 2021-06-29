package models

type AppleAppSiteAssociation struct {
	WebCredentials WebCredentials `json:"webcredentials"`
}

type WebCredentials struct {
	Apps []string `json:"apps"`
}
