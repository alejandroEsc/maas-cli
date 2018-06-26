package main


import (
	"github.com/spf13/cobra"
	m "github.com/alejandroEsc/maas-client-sample/pkg/maas"

	"github.com/alejandroEsc/maas-client-sample/pkg/cli"
	"os"
	"github.com/juju/gomaasapi"
	"github.com/spf13/viper"
	"fmt"
)


func ListMachinesCmd() *cobra.Command {
	mo := &cli.ListMachineOptions{}
	listMachinesCmd := &cobra.Command{
		Use: "list-machines ...",
		Short: "list machines in maas",
		Long: "",
		Run: func(cmd *cobra.Command, args []string) {

			if err := runListMachineCmd(mo); err != nil {
				logger.Criticalf(err.Error())
				os.Exit(1)
			}

		},
	}

	fs := listMachinesCmd.Flags()

	//bindCommonMAASFlags(&mo.MAASOptions, fs)
	fs.StringVar(&mo.APIKey, "api-key", viper.GetString(keyAPIKey), "maas apikey")
	fs.StringVar(&mo.MAASURLKey, "maas-url", viper.GetString(keyMAASURL), "maas url")
	fs.StringVar(&mo.MAASAPIVersionKey, "api-version", viper.GetString(keyMAASAPIVersion), "maas api version")

	fs.BoolVar(&mo.Detailed, "detailed", false, "print all details")

	return listMachinesCmd
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

	if o.Detailed {
		return printLong(machinesArray)
	} else {
		printShort(machinesArray)
	}
	return nil
}

func printShort(machinesArray []gomaasapi.JSONObject) {
	ms := make([]m.Machine, len(machinesArray))

	for i, machineObj := range machinesArray {
		machine, err := machineObj.GetMAASObject()
		checkError(err)

		machineName, err := machine.GetField("hostname")
		checkErrorMsg(err, "could not get hostname")

		machineSystemID, err := machine.GetField("system_id")
		checkErrorMsg(err, "could not get system_id")

		hweKernel, err := machine.GetField("hwe_kernel")
		checkErrorMsg(err, "could not get hwe_kernel")

		os, err := machine.GetField("osystem")
		checkErrorMsg(err, "could not get osystem")

		m := m.Machine{Name: machineName, SystemID: machineSystemID, Kernel: hweKernel, OS: os}
		ms[i] = m
	}

	printList(ms)
}

func printLong(machinesArray []gomaasapi.JSONObject) error {
	for i, machineObj := range machinesArray {
		machine, err := machineObj.GetMAASObject()
		j, err := machine.MarshalJSON()
		if err != nil {
			return err
		}
		fmt.Printf("\n --- machine: %d ---\n",i)
		fmt.Printf("%s",j)
	}
	return nil
}

func checkError(err error) {
	if err != nil {
		logger.Errorf(err.Error())
	}
}

func checkErrorMsg(err error, msg string) {
	if err != nil {
		logger.Errorf("%s, %s", msg, err.Error())
	}
}


func printList(ms []m.Machine) {
	if len(ms) <=0 {

	} else {
		for i, machine := range ms {
			fmt.Printf("|\t %d \t|\t %s \t|\t %s \t|\t %s-%s \t| \n", i, machine.SystemID, machine.Name, machine.OS, machine.Kernel)
		}
	}


}