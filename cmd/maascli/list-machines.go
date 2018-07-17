package main

import (
	"github.com/alejandroEsc/golang-maas-client/pkg/api"
	"github.com/alejandroEsc/golang-maas-client/pkg/api/v2"
	"github.com/spf13/cobra"

	"fmt"
	"os"

	"encoding/json"

	"github.com/alejandroEsc/maas-cli/pkg/cli"
)

const (
	printMachineFmt = "\t %d \t %s \t %s \t %s \t %s \t %s \t %s \t\n"
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

	maas, err := api.NewMASS(o.MAASURLKey, o.MAASAPIVersionKey, o.APIKey)
	if err != nil {
		return err
	}

	params := v2.MachinesParams(v2.MachinesArgs{})
	rawMachines, err := maas.Get("machines", "", params.Values)

	var machines []v2.Machine
	err = json.Unmarshal(rawMachines, &machines)
	if err != nil {
		return err
	}

	if len(machines) == 0 {
		return nil
	}

	if o.Detailed {
		return printMachinesDetailed(machines)
	}

	printMachinesSummary(machines)

	return nil
}

func printMachinesSummary(machinesArray []v2.Machine) {
	for i, m := range machinesArray {
		fmt.Printf(printMachineFmt, i,
			m.SystemID,
			m.Hostname,
			m.OperatingSystem,
			m.Kernel,
			m.PowerState,
			m.StatusName,
		)
	}
}

func printMachinesDetailed(machinesArray []v2.Machine) error {
	for i, m := range machinesArray {
		fmt.Printf("\n --- machine: %d ---\n", i)
		fmt.Printf("%+v\n", m)
	}
	return nil
}
