package main

import (
	m "github.com/alejandroEsc/maas-client-sample/pkg/maas"
	"github.com/spf13/cobra"

	"fmt"
	"os"

	"github.com/alejandroEsc/maas-client-sample/pkg/cli"
	"github.com/juju/gomaasapi"
	"github.com/spf13/viper"
)

const (
	printFmt = "|\t %d \t|\t %s \t|\t %s \t|\t %s:%s \t|\t %s \t|\t %s \t| \n"
)

func listMachinesCmd() *cobra.Command {
	mo := &cli.ListMachineOptions{}
	cmd := &cobra.Command{
		Use:   "list-machines ...",
		Short: "list machines in maas",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {

			if err := runListMachineCmd(mo); err != nil {
				logger.Criticalf(err.Error())
				os.Exit(1)
			}

		},
	}

	fs := cmd.Flags()

	//bindCommonMAASFlags(&mo.options, fs)
	fs.StringVar(&mo.APIKey, "api-key", viper.GetString(keyAPIKey), "maas apikey")
	fs.StringVar(&mo.MAASURLKey, "maas-url", viper.GetString(keyMAASURL), "maas url")
	fs.StringVar(&mo.MAASAPIVersionKey, "api-version", viper.GetString(keyMAASAPIVersion), "maas api version")

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

	maasCLI := m.NewMaasClient(maas)

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
		return printLong(machinesArray)
	}

	printShort(machinesArray)

	return nil
}

func printShort(machinesArray []gomaasapi.JSONObject) {
	mON := make([]m.Machine, 0)
	mOFF := make([]m.Machine, 0)
	mUnknown := make([]m.Machine, 0)

	for _, machineObj := range machinesArray {
		machine, err := machineObj.GetMAASObject()
		logError(err)

		machineName, err := machine.GetField("hostname")
		logError(err)

		machineSystemID, err := machine.GetField("system_id")
		logError(err)

		hweKernel, err := machine.GetField("hwe_kernel")
		logError(err)

		os, err := machine.GetField("osystem")
		logError(err)

		power, err := machine.GetField("power_state")
		logError(err)

		status, err := machine.GetField("status_name")
		logError(err)

		m := m.Machine{
			Name:       machineName,
			SystemID:   machineSystemID,
			Kernel:     hweKernel,
			OS:         os,
			PowerState: power,
			Status:     status,
		}

		switch power {
		case "on":
			mON = append(mON, m)
		case "off":
			mOFF = append(mOFF, m)
		default:
			mUnknown = append(mUnknown, m)
		}
	}

	// print machines that are on

	if len(mON) != 0 {
		fmt.Println("--- ON ---")
		printMachines(mON)
	}

	if len(mOFF) != 0 {
		fmt.Println("--- OFF ---")
		printMachines(mOFF)
	}

	if len(mUnknown) != 0 {
		fmt.Println("--- UNKONWN ---")
		printMachines(mUnknown)
	}

}

func printMachines(ms []m.Machine) {
	for i, mn := range ms {
		fmt.Printf(
			printFmt,
			i,
			mn.SystemID,
			mn.Name,
			mn.OS,
			mn.Kernel,
			mn.PowerState,
			mn.Status)
	}
}

func printLong(machinesArray []gomaasapi.JSONObject) error {
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

func logError(err error) {
	if err != nil {
		logger.Errorf(err.Error())
	}
}
