package main

import (
	m "github.com/alejandroEsc/maas-cli/pkg/maas"
	"github.com/spf13/cobra"

	"os"

	"fmt"

	"github.com/alejandroEsc/maas-cli/pkg/cli"
)

func machineReleaseCmd() *cobra.Command {
	mo := &cli.MachineOptions{}
	cmd := &cobra.Command{
		Use:   "release [machineID]",
		Short: "Release action against one or more machines.",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {

			if err := runMachineActionCmd(m.ReleaseMachine, mo, args); err != nil {
				fmt.Println(err.Error())
				os.Exit(1)
			}

		},
	}

	fs := cmd.Flags()
	bindCommonMAASFlags(&mo.MAASOptions, fs)

	fs.StringVar(&mo.Params, "params", "", "paramaters to pass to an action")

	return cmd
}
