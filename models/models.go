package models

// Profile the user profile that you want to apply
type Profile struct {
	Name string
	EnabledProviders []string
	DisabledProviders []string
}