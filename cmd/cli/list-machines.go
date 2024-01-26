package main

import (
	"fmt"
	"os"

	gomaasclient "github.com/maas/gomaasclient/client"
	"github.com/spf13/cobra"

	"github.com/alejandroEsc/maas-cli/pkg/cli"
)

func listMachinesCmd() *cobra.Command {
	mo := &cli.ListMachineOptions{}
	cmd := &cobra.Command{
		Use:   "machines ...",
		Short: "list machines resources in a MAAS server",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {

			if err := runListMachineCmd(mo); err != nil {
				fmt.Println(err.Error())
				os.Exit(1)
			}

		},
	}

	fs := cmd.Flags()
	bindCommonMAASFlags(&mo.MAASOptions, fs)

	fs.BoolVar(&mo.Detailed, "detailed", false, "print all details")

	return cmd
}

func runListMachineCmd(o *cli.ListMachineOptions) error {

	maas, err := gomaasclient.GetClient(o.MAASURLKey, o.APIKey, o.MAASAPIVersionKey)
	if err != nil {
		return err
	}

	machines, err := maas.Machines.Get()
	if err != nil {
		return err
	}

	fmt.Println(machines)

	return nil
}
