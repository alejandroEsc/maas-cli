package main

import (
	"github.com/spf13/viper"
	flag "github.com/spf13/pflag"
	"github.com/alejandroEsc/maas-client-sample/pkg/cli"
)

const (

	envVarAPIKey = "MAAS_CLI_API_KEY"
	envVarURL = "MAAS_CLI_URL"
	envVarAPIVersion = "MAAS_CLI_API_VERSION"

	keyAPIKey         = "api_key"
	keyMAASURL        = "url"
	keyMAASAPIVersion = "api_version"


	keyMachineID = "machineID"
	keyMachineAction ="action"

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