package main


import (
	"github.com/spf13/cobra"
	m "github.com/alejandroEsc/maas-client-sample/pkg/maas"

	"github.com/alejandroEsc/maas-client-sample/pkg/cli"
	"github.com/spf13/viper"
	"github.com/juju/gomaasapi"
	"net/url"
)


func MachineCmd() *cobra.Command {
	mo := &cli.MachineOptions{}
	machinesCmd := &cobra.Command{
		Use: "machine",
		Short: "Run a few simple machine commands",
		Long: "",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Usage()
		},
	}
	fs := machinesCmd.Flags()
	fs.StringVar(&mo.APIKey, "api-key", viper.GetString(keyAPIKey), "maas apikey")
	fs.StringVar(&mo.MAASURLKey, "maas-url", viper.GetString(keyMAASURL), "maas url")
	fs.StringVar(&mo.MAASAPIVersionKey, "api-version", viper.GetString(keyMAASAPIVersion), "maas api version")

	machinesCmd.AddCommand(MachineStatusCmd())
	machinesCmd.AddCommand(MachineReleaseCmd())
	machinesCmd.AddCommand(MachineDeployCmd())


	return machinesCmd
}


func runMachineActionCmd(action m.MachineAction, o *cli.MachineOptions, args []string) error {
	var err error
	var params url.Values = nil

	if o.Params != "" {
		params, err = url.ParseQuery(o.Params)
		if err != nil {
			return err
		}
	}

	// Create API server endpoint.
	authClient, err := gomaasapi.NewAuthenticatedClient(gomaasapi.AddAPIVersionToURL(o.MAASURLKey, o.MAASAPIVersionKey), o.APIKey)
	if err != nil {
		return err
	}

	maas := gomaasapi.NewMAAS(*authClient)
	maasCLI := m.NewMaasClient(maas)

	for _, id := range args {
		result, err := maasCLI.PerformMachineAction(action, id, params)
		if err != nil {
			return err
		}

		err = fmtPrintJson(result)
		if err != nil {
			return err
		}
	}

	return nil
}
