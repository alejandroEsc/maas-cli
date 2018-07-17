package main

import (
	"github.com/spf13/cobra"

	"os"

	"github.com/alejandroEsc/maas-cli/pkg/cli"
	"github.com/spf13/viper"

	"github.com/alejandroEsc/golang-maas-client/pkg/api"
	"github.com/alejandroEsc/golang-maas-client/pkg/api/v2"

	"encoding/json"
	"fmt"
	"net/url"
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

	maas, err := api.NewMASS(o.MAASURLKey, o.MAASAPIVersionKey, o.APIKey)
	if err != nil {
		return err
	}

	versionBytes, err := maas.Get("version", "", url.Values{})
	if err != nil {
		return err
	}

	var version v2.Version
	err = json.Unmarshal(versionBytes, &version)
	if err != nil {
		return err
	}

	fmt.Printf("Version: %s\nSubVersion %s\n", version.Version, version.SubVersion)

	return nil
}
