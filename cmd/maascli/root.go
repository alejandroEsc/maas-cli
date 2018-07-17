package main

import (
	"github.com/spf13/cobra"

	"github.com/alejandroEsc/maas-cli/pkg/cli"
	"github.com/alejandroEsc/maas-cli/pkg/util"
	"github.com/juju/loggo"

	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/viper"
)

var (
	logLevel = "UNSPECIFIED"
	logger   = util.GetModuleLogger("cmd.maascli", loggo.UNSPECIFIED)
	options  = &cli.MAASOptions{}
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "maas-cli",
	Short: "MAAS CLI tool",
	Run: func(cmd *cobra.Command, args []string) {
		level, isValid := loggo.ParseLevel(logLevel)
		if isValid {
			logger.SetLogLevel(level)
		}

		cmd.Help()
	},
}

func init() {
	// init viper defaults
	initEnvDefaults()

	// root flags
	RootCmd.PersistentFlags().StringVar(&options.APIKey, "api-key", viper.GetString(keyAPIKey), "maas apikey")
	RootCmd.PersistentFlags().StringVar(&options.MAASURLKey, "maas-url", viper.GetString(keyMAASURL), "maas url")
	RootCmd.PersistentFlags().StringVar(&options.MAASAPIVersionKey, "api-version", viper.GetString(keyMAASAPIVersion), "maas api version")
	RootCmd.PersistentFlags().StringVarP(&logLevel, "verbose", "v", "UNSPECIFIED", "log level")

	// bind environment vars
	bindEnvVars()

	// add commands
	addCommands()
}

func addCommands() {
	RootCmd.AddCommand(machineCmd())
	RootCmd.AddCommand(listCmd())
	RootCmd.AddCommand(versionCmd())
}

// Execute performs root command task.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func fmtPrintJSON(o []byte) error {
	jp, err := json.MarshalIndent(string(o), "", "\t")
	if err != nil {
		return err
	}
	fmt.Printf("\n%s", jp)
	return nil
}

func logError(err error) {
	if err != nil {
		logger.Errorf(err.Error())
	}
}
