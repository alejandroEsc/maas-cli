package main

import (
	"fmt"
	"os"

	gomaasclient "github.com/maas/gomaasclient/client"
	"github.com/spf13/cobra"

	"github.com/alejandroEsc/maas-cli/pkg/cli"
)

func machineCommissionCmd() *cobra.Command {
	mo := &cli.CommissionMachineOpts{}
	cmd := &cobra.Command{
		Use:   "commission [machineID]",
		Short: "Commission action against one more more machines.",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {

			if err := runMachineCommissionCmd(mo, args); err != nil {
				fmt.Println(err.Error())
				os.Exit(1)
			}

		},
	}
	fs := cmd.Flags()
	bindCommonMAASFlags(&mo.MAASOptions, fs)

	fs.IntVar(&mo.EnableSSH, "enable-ssh", 1, "Whether to enable SSH for the commissioning environment using the user's SSH key(s).")
	fs.IntVar(&mo.SkipBMCConfig, "skip-bmc-config", 0, "Whether to skip re-configuration of the BMC for IPMI based machines.")
	fs.IntVar(&mo.SkipNetworking, "skip-networking", 0, "Whether to skip re-configuring the networking on the machine after the commissioning has completed.")
	fs.IntVar(&mo.SkipStorage, "skip-storage", 0, "Whether to skip re-configuring the storage on the machine after the commissioning has completed.")
	fs.StringVar(&mo.CommissioningScripts, "commissioning-scripts", "", "paramaters to pass to an action")
	fs.StringVar(&mo.TestingScripts, "testing-scripts", "", "paramaters to pass to an action")

	return cmd
}

func runMachineCommissionCmd(o *cli.CommissionMachineOpts, args []string) error {
	maas, err := gomaasclient.GetClient(o.MAASURLKey, o.APIKey, o.MAASAPIVersionKey)
	if err != nil {
		return err
	}

	for _, id := range args {
		result, err := maas.Machine.Commission(id, &o.MachineCommissionParams)
		if err != nil {
			return err
		}

		fmt.Printf("%v\n", result)

	}

	return nil
}
