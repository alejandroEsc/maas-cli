package main

import (
	m "github.com/alejandroEsc/maas-client-sample/pkg/maas"
	"github.com/juju/gomaasapi"
	"github.com/spf13/cobra"

	"os"

	"github.com/alejandroEsc/maas-client-sample/pkg/cli"
	"github.com/spf13/viper"
)

func versionCmd() *cobra.Command {
	vo := &cli.VersionOptions{}
	cmd := &cobra.Command{
		Use:   "version",
		Short: "Get Version info",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {

			if err := runVersionCmd(vo); err != nil {
				logger.Criticalf(err.Error())
				os.Exit(1)
			}

		},
	}

	fs := cmd.Flags()
	fs.StringVar(&vo.APIKey, "api-key", viper.GetString(keyAPIKey), "maas apikey")
	fs.StringVar(&vo.MAASURLKey, "maas-url", viper.GetString(keyMAASURL), "maas url")
	fs.StringVar(&vo.MAASAPIVersionKey, "api-version", viper.GetString(keyMAASAPIVersion), "maas api version")

	return cmd
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
	fmtPrintJSON(version)

	return nil
}
