package main

import (
	"github.com/spf13/cobra"

	"github.com/juju/loggo"
	"github.com/alejandroEsc/maas-client-sample/pkg/util"
	"github.com/alejandroEsc/maas-client-sample/pkg/cli"

	"fmt"
	"os"
	"github.com/spf13/viper"
	"github.com/juju/gomaasapi"
	"encoding/json"
)

var (
	logger = util.GetModuleLogger("cmd.maascli", loggo.INFO)
	MAASOptions = &cli.MAASOptions{}
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "maascli",
	Short: "MAAS CLI tool",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}


func init(){
	// init viper defaults
	initEnvDefaults()

	// root flags
	RootCmd.PersistentFlags().StringVar(&MAASOptions.APIKey, "api-key", viper.GetString(keyAPIKey), "maas apikey")
	RootCmd.PersistentFlags().StringVar(&MAASOptions.MAASURLKey, "maas-url", viper.GetString(keyMAASURL), "maas url")
	RootCmd.PersistentFlags().StringVar(&MAASOptions.MAASAPIVersionKey, "api-version", viper.GetString(keyMAASAPIVersion), "maas api version")

	// bind environment vars
	bindEnvVars()

	// add commands
	addCommands()
}


func addCommands() {
	RootCmd.AddCommand(MachineCmd())
	RootCmd.AddCommand(ListMachinesCmd())
	RootCmd.AddCommand(VersionCmd())
}



// Execute performs root command task.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func fmtPrintJson(o gomaasapi.JSONObject) error {
	jp, err := json.MarshalIndent(o, "", "\t")
	if err != nil {
		return err
	}
	fmt.Printf("\n%s",jp)
	return nil
}
