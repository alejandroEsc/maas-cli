package main

import (
	"github.com/spf13/cobra"

	"github.com/alejandroEsc/maas-cli/pkg/cli"
)

func machineCmd() *cobra.Command {
	mo := &cli.MachineOptions{}
	cmd := &cobra.Command{
		Use:   "machine",
		Short: "Run a few simple machine commands",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Usage()
		},
	}
	fs := cmd.Flags()
	bindCommonMAASFlags(&mo.MAASOptions, fs)

	cmd.AddCommand(machineStatusCmd())
	cmd.AddCommand(machineReleaseCmd())
	cmd.AddCommand(machineDeployCmd())
	cmd.AddCommand(machineCommissionCmd())

	return cmd
}
