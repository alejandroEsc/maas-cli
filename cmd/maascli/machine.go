package main


import (
	"github.com/spf13/cobra"
	m "github.com/alejandroEsc/maas-client-sample/pkg/maas"

	"github.com/alejandroEsc/maas-client-sample/pkg/cli"
	"os"
	"github.com/spf13/viper"
	"fmt"
	"strings"
	"github.com/juju/gomaasapi"
)


func MachineCmd() *cobra.Command {
	mo := &cli.MachineOptions{}
	machinesCmd := &cobra.Command{
		Use: "machine ...",
		Short: "Run a few simple machine commands",
		Long: "",
		Run: func(cmd *cobra.Command, args []string) {

			if err := runMachineCmd(mo); err != nil {
				logger.Criticalf(err.Error())
				os.Exit(1)
			}

		},
	}
	fs := machinesCmd.Flags()
	fs.StringVar(&mo.MachineID, keyMachineID, viper.GetString(keyMachineID), "id of machine provisioned in maas")
	fs.StringVar(&mo.MachineAction, keyMachineAction, viper.GetString(keyMachineAction), "action to perform against a machine, e.g., Commission, Deploy, etc..")


	fs.StringVar(&mo.APIKey, "api-key", viper.GetString(keyAPIKey), "maas apikey")
	fs.StringVar(&mo.MAASURLKey, "maas-url", viper.GetString(keyMAASURL), "maas url")
	fs.StringVar(&mo.MAASAPIVersionKey, "api-version", viper.GetString(keyMAASAPIVersion), "maas api version")


	return machinesCmd
}


func runMachineCmd(o *cli.MachineOptions) error {
	var err error
	logger.Infof("%s %s %s", o.APIKey, o.MAASURLKey, o.MAASAPIVersionKey)

	// Create API server endpoint.
	authClient, err := gomaasapi.NewAuthenticatedClient(gomaasapi.AddAPIVersionToURL(o.MAASURLKey, o.MAASAPIVersionKey), o.APIKey)
	CheckError(err)
	maas := gomaasapi.NewMAAS(*authClient)

	maasCLI := m.NewMaasClient(maas)


	if o.MachineID == "" {
		fmt.Println("No machine id defined, cannot proceed.")
	}


	if o.MachineAction == "" {
		fmt.Println("No action defined.")
	}

	switch strings.ToLower(o.MachineAction) {
	case "commission":
		result, err := maasCLI.CommisionMachine(o.MachineID)
		if err != nil {
			return err
		}
		fmt.Printf("result: %v", result)
	case "release":
		result, err := maasCLI.ReleaseMachine(o.MachineID)
		if err != nil {
			return err
		}
		fmt.Printf("result: %v", result)
	case "deploy":
		result, err := maasCLI.DeployMachine(o.MachineID)
		if err != nil {
			return err
		}
		fmt.Printf("result: %v", result)
	default:
		fmt.Printf("Action %s is not supported.\n",o.MachineAction)
	}


	return nil
}