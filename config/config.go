package config

// Auth0Config is ...
type Auth0Config struct {
	Domain       string `toml:"domain"`
	ClientID     string `toml:"clientID"`
	Secret       string `toml:"secret"`
	Callback     string `toml:"callback"`
	CallbackForm string `toml:"callback_form"`
	Login        string `toml:"login"`
}
