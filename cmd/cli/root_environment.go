package main

import (
	"github.com/alejandroEsc/maas-cli/pkg/cli"
	flag "github.com/spf13/pflag"
	"github.com/spf13/viper"
)

const (
	envVarAPIKey     = "MAAS_APIKEY"
	envVarURL        = "MAAS_URL"
	envVarAPIVersion = "MAAS_VERSION"

	keyAPIKey         = "api_key"
	keyMAASURL        = "url"
	keyMAASAPIVersion = "api_version"
)

func initEnvDefaults() {
	viper.SetDefault(keyAPIKey, "")
	viper.SetDefault(keyMAASURL, "")
	viper.SetDefault(keyMAASAPIVersion, "2.0")

}

func bindEnvVars() {
	viper.AutomaticEnv()
	viper.SetEnvPrefix("maas_cli")

	viper.BindEnv(keyMAASURL, envVarURL)
	viper.BindEnv(keyAPIKey, envVarAPIKey)
	viper.BindEnv(keyMAASAPIVersion, envVarAPIVersion)
}

func bindCommonMAASFlags(o *cli.MAASOptions, fs *flag.FlagSet) {
	fs.StringVar(&o.APIKey, "api-key", viper.GetString(keyAPIKey), "maas apikey")
	fs.StringVar(&o.MAASURLKey, "maas-url", viper.GetString(keyMAASURL), "maas url")
	fs.StringVar(&o.MAASAPIVersionKey, "api-version", viper.GetString(keyMAASAPIVersion), "maas api version")
}
