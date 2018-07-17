package main

import (
	"fmt"
	"os"

	"github.com/alejandroEsc/golang-maas-client/pkg/api"
	"github.com/alejandroEsc/golang-maas-client/pkg/api/v2"
	"github.com/spf13/cobra"

	"github.com/alejandroEsc/maas-cli/pkg/cli"
	"encoding/json"
)

func machineCommissionCmd() *cobra.Command {
	mo := &cli.CommissionMachineOpts{}
	cmd := &cobra.Command{
		Use:   "commission [machineID]",
		Short: "Commission action agains one more more machines.",
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

	fs.BoolVar(&mo.EnableSSH, "enable-ssh", true, "Whether to enable SSH for the commissioning environment using the user's SSH key(s).")
	fs.BoolVar(&mo.SkipBMCConfig, "skip-bmc-config", false, "Whether to skip re-configuration of the BMC for IPMI based machines.")
	fs.BoolVar(&mo.SkipNetworking, "skip-networking", false, "Whether to skip re-configuring the networking on the machine after the commissioning has completed.")
	fs.BoolVar(&mo.SkipStorage, "skip-storage", false, "Whether to skip re-configuring the storage on the machine after the commissioning has completed.")
	fs.StringVar(&mo.CommissioningScripts, "commissioning-scripts", "", "paramaters to pass to an action")
	fs.StringVar(&mo.TestingScript, "testing-scripts", "", "paramaters to pass to an action")

	return cmd
}

func runMachineCommissionCmd(o *cli.CommissionMachineOpts, args []string) error {
	params := v2.ComssionMachineParams(o.CommissionMachineArgs)
	maas, err := api.NewMASS(o.MAASURLKey, o.MAASAPIVersionKey, o.APIKey)
	if err != nil {
		return err
	}

	for i, id := range args {
		result, err := maas.Post("machines/"+id, string(v2.MachineComission), params.Values)
		if err != nil {
			return err
		}

		var m v2.Machine
		err = json.Unmarshal(result, &m)
		if err != nil {
			logger.Errorf(err.Error())
			continue
		}
		fmt.Printf(printMachineFmt, i,
			m.SystemID,
			m.Hostname,
			m.OperatingSystem,
			m.Kernel,
			m.PowerState,
			m.StatusName,
		)
	}

	return nil
}
