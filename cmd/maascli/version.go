package main

import (
	"github.com/spf13/cobra"
	"github.com/juju/gomaasapi"
	m "github.com/alejandroEsc/maas-client-sample/pkg/maas"

	"os"
	"github.com/spf13/viper"
	"github.com/alejandroEsc/maas-client-sample/pkg/cli"
)
func VersionCmd() *cobra.Command {
	vo := &cli.VersionOptions{}
	versionCmd := &cobra.Command{
		Use: "version",
		Short: "Get Version info",
		Long: "",
		Run: func(cmd *cobra.Command, args []string) {

			if err := runVersionCmd(vo); err != nil {
				logger.Criticalf(err.Error())
				os.Exit(1)
			}

		},
	}

	fs := versionCmd.Flags()
	fs.StringVar(&vo.APIKey, "api-key", viper.GetString(keyAPIKey), "maas apikey")
	fs.StringVar(&vo.MAASURLKey, "maas-url", viper.GetString(keyMAASURL), "maas url")
	fs.StringVar(&vo.MAASAPIVersionKey, "api-version", viper.GetString(keyMAASAPIVersion), "maas api version")


	return versionCmd
}



func runVersionCmd(o *cli.VersionOptions) error {
	var err error

	// Create API server endpoint.
	authClient, err := gomaasapi.NewAuthenticatedClient(gomaasapi.AddAPIVersionToURL(o.MAASURLKey, o.MAASAPIVersionKey), o.APIKey)
	if err != nil {
		return err
	}
	maas := gomaasapi.NewMAAS(*authClient)

	maasCLI := m.NewMaasClient(maas)

	version, err := maasCLI.GetMAASVersion()
	if err != nil {
		return err
	}
	fmtPrintJson(version)

	return nil
}