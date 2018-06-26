package main


import (
	"github.com/spf13/cobra"
	m "github.com/alejandroEsc/maas-client-sample/pkg/maas"

	"github.com/alejandroEsc/maas-client-sample/pkg/cli"
	"os"
	"github.com/spf13/viper"
)


func MachineReleaseCmd() *cobra.Command {
	mo := &cli.MachineOptions{}
	machineReleaseCmd := &cobra.Command{
		Use: "machine",
		Short: "Run a few simple machine commands",
		Long: "",
		Run: func(cmd *cobra.Command, args []string) {

			if err := runMachineActionCmd(m.ReleaseMachine, mo, args); err != nil {
				logger.Criticalf(err.Error())
				os.Exit(1)
			}

		},
	}
	fs := machineReleaseCmd.Flags()
	fs.StringVar(&mo.APIKey, "api-key", viper.GetString(keyAPIKey), "maas apikey")
	fs.StringVar(&mo.MAASURLKey, "maas-url", viper.GetString(keyMAASURL), "maas url")
	fs.StringVar(&mo.MAASAPIVersionKey, "api-version", viper.GetString(keyMAASAPIVersion), "maas api version")
	fs.StringVar(&mo.Params, "params","","paramaters to pass to an action")

	return machineReleaseCmd
}





