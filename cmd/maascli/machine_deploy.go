package main

import (
	"github.com/spf13/cobra"

	"fmt"
	"os"

	"github.com/alejandroEsc/golang-maas-client/pkg/api"
	"github.com/alejandroEsc/golang-maas-client/pkg/api/v2"

	"encoding/json"

	"github.com/alejandroEsc/maas-cli/pkg/cli"
)

func machineDeployCmd() *cobra.Command {
	mo := &cli.DeployMachineOpts{}
	cmd := &cobra.Command{
		Use:   "deploy [machineID]",
		Short: "Deploy action against one ore multiple machines.",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {

			if err := runMachineDeployCmd(mo, args); err != nil {
				fmt.Println(err.Error())
				os.Exit(1)
			}

		},
	}
	fs := cmd.Flags()
	bindCommonMAASFlags(&mo.MAASOptions, fs)

	fs.StringVar(&mo.UserData, "user-data", "", " If present, this blob of user-data to be made available to the machines through the metadata service")
	fs.StringVar(&mo.DistroSeries, "distro-series", "", "If present, this parameter specifies the OS release the machine will use")
	fs.StringVar(&mo.Kernel, "hwe-kernel", "", "If present, this parameter specified the kernel to be used on the machine")
	fs.StringVar(&mo.AgentName, "agent-name", "", "An optional agent name to attach to the acquired machine")
	fs.BoolVar(&mo.BridgeAll, "bridge-all", false, "Optionally create a bridge interface for every configured interface on the machine. The created bridges will be removed once the machine is released")
	fs.BoolVar(&mo.BridgeSTP, "bridge-stp", false, "Optionally turn spanning tree protocol on or off for the bridges created on every configured interface")
	fs.IntVar(&mo.BridgeFD, "bridge-fd", 15, "Optionally adjust the forward delay to time seconds")
	fs.StringVar(&mo.Comment, "comment", "", " Optional comment for the event log")
	fs.BoolVar(&mo.InstallRackd, "install-rackd", false, "If True, the Rack Controller will be installed on this machine")

	return cmd
}

func runMachineDeployCmd(o *cli.DeployMachineOpts, args []string) error {
	params := v2.DeploytMachineParams(o.DeployMachineArgs)
	maas, err := api.NewMASS(o.MAASURLKey, o.MAASAPIVersionKey, o.APIKey)
	if err != nil {
		return err
	}

	for i, id := range args {
		result, err := maas.Post("machines/"+id, string(v2.MachineDeploy), params.Values)
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
