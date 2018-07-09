package main

import (
	m "github.com/alejandroEsc/maas-cli/pkg/maas"
	"github.com/spf13/cobra"

	"fmt"
	"os"

	"encoding/json"

	"github.com/alejandroEsc/maas-cli/pkg/cli"
	"github.com/juju/gomaasapi"
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
	// Create API server endpoint.
	authClient, err := gomaasapi.NewAuthenticatedClient(gomaasapi.AddAPIVersionToURL(o.MAASURLKey, o.MAASAPIVersionKey), o.APIKey)
	if err != nil {
		return err
	}
	maas := gomaasapi.NewMAAS(*authClient)
	maasCLI := m.NewMaas(maas)

	listObj, err := maasCLI.GetMachines()
	if err != nil {
		return err
	}

	machinesArray, err := listObj.GetArray()
	if err != nil {
		return err
	}

	if len(machinesArray) == 0 {
		return nil
	}

	if o.Detailed {
		return printMachinesDetailed(machinesArray)
	}

	printMachinesSummary(machinesArray)

	return nil
}

func printMachinesSummary(machinesArray []gomaasapi.JSONObject) {

	for i, machineObj := range machinesArray {
		var m m.Machine
		machine, err := machineObj.GetMAASObject()
		logError(err)

		j, err := machine.MarshalJSON()
		logError(err)

		err = json.Unmarshal(j, &m)
		logError(err)

		fmt.Printf(printMachineFmt, i,
			m.SystemID,
			m.Hostname,
			m.OS,
			m.Kernel,
			m.PowerState,
			m.Status,
		)
	}

}

func printMachinesDetailed(machinesArray []gomaasapi.JSONObject) error {
	for i, machineObj := range machinesArray {
		machine, err := machineObj.GetMAASObject()
		j, err := machine.MarshalJSON()
		if err != nil {
			return err
		}
		fmt.Printf("\n --- machine: %d ---\n", i)
		fmt.Printf("%s", j)
	}
	return nil
}
