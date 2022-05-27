package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type config struct {
	AppName string `mapstructure:"APP_NAME"`
	AppUrl  string `mapstructure:"APP_URL"`
	AppPort string `mapstructure:"APP_PORT"`

	ActualSecret string `mapstructure:"ACTUAL_SECRET"`
	ActualUrl    string `mapstructure:"ACTUAL_URL"`

	PlaidCountryCodes string `mapstructure:"PLAID_COUNTRY_CODES"`
	PlaidClientName   string `mapstructure:"PLAID_CLIENT_NAME"`
	PlaidLanguage     string `mapstructure:"PLAID_LANGUAGE"`
	PlaidClientId     string `mapstructure:"PLAID_CLIENT_ID"`
	PlaidSecret       string `mapstructure:"PLAID_SECRET"`
	PlaidMode         string `mapstructure:"PLAID_MODE"`
}

var Config *config

func LoadConfig(path string) (err error) {
	viper.AddConfigPath(path)

	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig() // Find and read the config file
	if err != nil {            // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	err = viper.Unmarshal(&Config)

	return
}
