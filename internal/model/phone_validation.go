package model

type PhoneValidationResult struct {
	IsValid       bool
	CountryCode   string
	CountryPrefix string
	Location      string
	Carrier       string
}
