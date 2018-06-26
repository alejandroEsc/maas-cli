package main


import (
	"github.com/spf13/cobra"
	m "github.com/alejandroEsc/maas-client-sample/pkg/maas"

	"github.com/alejandroEsc/maas-client-sample/pkg/cli"
	"os"
	"fmt"
	"github.com/juju/gomaasapi"
	"github.com/spf13/viper"
)


func MachineStatusCmd() *cobra.Command {
	mo := &cli.MachineOptions{}
	machineStatusCmd := &cobra.Command{
		Use: "status",
		Short: "get machine status",
		Long: "Returns the MAAS concept of machine status",
		Run: func(cmd *cobra.Command, args []string) {

			if len(args) == 0 {
				logger.Criticalf("please include machine id")
			}

			if err := runMachineStatusCmd(mo, args); err != nil {
				logger.Criticalf(err.Error())
				os.Exit(1)
			}

		},
	}
	fs := machineStatusCmd.Flags()

	fs.StringVar(&mo.APIKey, "api-key", viper.GetString(keyAPIKey), "maas apikey")
	fs.StringVar(&mo.MAASURLKey, "maas-url", viper.GetString(keyMAASURL), "maas url")
	fs.StringVar(&mo.MAASAPIVersionKey, "api-version", viper.GetString(keyMAASAPIVersion), "maas api version")

	return machineStatusCmd
}


func runMachineStatusCmd(o *cli.MachineOptions, args []string) error {
	var err error

	// Create API server endpoint.
	authClient, err := gomaasapi.NewAuthenticatedClient(gomaasapi.AddAPIVersionToURL(o.MAASURLKey, o.MAASAPIVersionKey), o.APIKey)
	if err != nil {
		return err
	}

	maas := gomaasapi.NewMAAS(*authClient)
	maasCLI := m.NewMaasClient(maas)

	for _, id := range args {
		result, err := maasCLI.GetStatus(id)
		if err != nil {
			logger.Errorf(err.Error())
			continue
		}

		fmt.Println(result)
	}

	return nil
}


