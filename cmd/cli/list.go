package main

import (
	"github.com/spf13/cobra"

	"github.com/alejandroEsc/maas-cli/pkg/cli"
)

func listCmd() *cobra.Command {
	mo := &cli.ListOptions{}
	cmd := &cobra.Command{
		Use:   "list",
		Short: "list MAAS resources.",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Usage()
		},
	}

	fs := cmd.Flags()
	bindCommonMAASFlags(&mo.MAASOptions, fs)

	cmd.AddCommand(listMachinesCmd())
	cmd.AddCommand(listNodesCmd())

	return cmd
}
