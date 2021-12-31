package model

type Database struct {
	Name string `json:"name"`
}

type Configuration struct {
	Port                int      `json:"port"`
	Database            Database `json:"database"`
	ApiKeyCurrencyLayer string   `json:"api_key_currency_layer"`
}
