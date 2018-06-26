package main

import (
	m "github.com/alejandroEsc/maas-cli/pkg/maas"
	"github.com/spf13/cobra"

	"os"

	"github.com/alejandroEsc/maas-cli/pkg/cli"
	"fmt"
)

func machineCommissionCmd() *cobra.Command {
	mo := &cli.MachineOptions{}
	cmd := &cobra.Command{
		Use:   "commission [machineID]",
		Short: "Commission action agains one more more machines.",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {

			if err := runMachineActionCmd(m.CommissionMachine, mo, args); err != nil {
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
