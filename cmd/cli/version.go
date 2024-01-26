package main

import (
	"fmt"
	"os"

	"github.com/alejandroEsc/maas-cli/pkg/cli"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	gomaasclient "github.com/maas/gomaasclient/client"
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

	maas, err := gomaasclient.GetClient(o.MAASURLKey, o.APIKey, o.MAASAPIVersionKey)
	if err != nil {
		return err
	}

	versionBytes, err := maas.MAASServer.Get("version")
	if err != nil {
		return err
	}

	fmt.Printf("%+v\n", string(versionBytes))

	return nil
}
