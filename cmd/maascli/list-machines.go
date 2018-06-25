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

	return listMachinesCmd
}


func runListMachineCmd(o *cli.ListMachineOptions) error {
	logger.Infof("%s %s %s", o.APIKey, o.MAASURLKey, o.MAASAPIVersionKey)

	// Create API server endpoint.
	authClient, err := gomaasapi.NewAuthenticatedClient(gomaasapi.AddAPIVersionToURL(o.MAASURLKey, o.MAASAPIVersionKey), o.APIKey)
	CheckError(err)
	maas := gomaasapi.NewMAAS(*authClient)

	maasCLI := m.NewMaasClient(maas)

	listObj, err := maasCLI.GetMachines()
	CheckError(err)

	machinesArray, err := listObj.GetArray()
	CheckError(err)


	ms := make([]m.Machine, len(machinesArray))

	for i, machineObj := range machinesArray {
		machine, err := machineObj.GetMAASObject()
		CheckError(err)


		machineName, err := machine.GetField("hostname")
		CheckErrorMsg(err, "could not get hostname")

		machineSystemID, err := machine.GetField("system_id")
		CheckErrorMsg(err, "could not get system_id")

		hweKernel, err := machine.GetField("hwe_kernel")
		CheckErrorMsg(err, "could not get hwe_kernel")

		os, err := machine.GetField("osystem")
		CheckErrorMsg(err, "could not get osystem")

		m := m.Machine{Name: machineName, SystemID: machineSystemID, Kernel: hweKernel, OS: os}
		ms[i] = m
	}

	printList(ms)
	return nil
}


func printList(ms []m.Machine) {

	if len(ms) <=0 {

	} else {
		fmt.Printf("|\t %s \t|\t %s \t|\t %s \t\t|\t %s-%s \t| \n", "item", "system-ID", "name", "OS", "kernel")
		for i, machine := range ms {
			fmt.Printf("|\t %d \t|\t %s \t|\t %s \t|\t %s-%s \t| \n", i, machine.SystemID, machine.Name, machine.OS, machine.Kernel)

		}
	}


}