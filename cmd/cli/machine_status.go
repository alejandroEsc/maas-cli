package main

import (
	"fmt"
	"os"

	gomaasclient "github.com/maas/gomaasclient/client"
	"github.com/spf13/cobra"

	"github.com/alejandroEsc/maas-cli/pkg/cli"
)

func machineStatusCmd() *cobra.Command {
	mo := &cli.MachineOptions{}
	cmd := &cobra.Command{
		Use:   "status",
		Short: "get machine status",
		Long:  "Returns the MAAS concept of machine status",
		Run: func(cmd *cobra.Command, args []string) {

			if len(args) == 0 {
				fmt.Printf("Please include machine id\n")
			}

			if err := runMachineStatusCmd(mo, args); err != nil {
				fmt.Println(err.Error())
				os.Exit(1)
			}

		},
	}
	fs := cmd.Flags()
	bindCommonMAASFlags(&mo.MAASOptions, fs)

	return cmd
}

func runMachineStatusCmd(o *cli.MachineOptions, args []string) error {
	maas, err := gomaasclient.GetClient(o.MAASURLKey, o.APIKey, o.MAASAPIVersionKey)
	if err != nil {
		return err
	}

	for _, id := range args {
		result, errRelease := maas.Machine.GetPowerState(id)
		if errRelease != nil {
			return errRelease
		}

		fmt.Printf("%v", result)
	}

	return nil
}
