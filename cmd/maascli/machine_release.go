package main

import (
	"github.com/spf13/cobra"

	"os"

	"fmt"

	"github.com/alejandroEsc/golang-maas-client/pkg/api"
	"github.com/alejandroEsc/golang-maas-client/pkg/api/v2"

	"encoding/json"

	"github.com/alejandroEsc/maas-cli/pkg/cli"
)

func machineReleaseCmd() *cobra.Command {
	mo := &cli.ReleaseMachineOpts{}
	cmd := &cobra.Command{
		Use:   "release [machineID]",
		Short: "Release action against one or more machines.",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {

			if err := runMachineReleaseCmd(mo, args); err != nil {
				fmt.Println(err.Error())
				os.Exit(1)
			}

		},
	}

	fs := cmd.Flags()
	bindCommonMAASFlags(&mo.MAASOptions, fs)

	fs.StringVar(&mo.Comment, "comment", "", " Optional comment for the event log")
	fs.BoolVar(&mo.Erase, "erase", true, " Erase the disk when releasing.")
	fs.BoolVar(&mo.SecureErase, "secure-erase", false, " Erase the disk when releasing.")
	fs.BoolVar(&mo.QuickErase, "quick-erase", true, " Erase the disk when releasing.")
	return cmd
}

func runMachineReleaseCmd(o *cli.ReleaseMachineOpts, args []string) error {
	params := v2.ReleaseMachinesParams(o.ReleaseMachinesArgs)
	maas, err := api.NewMASS(o.MAASURLKey, o.MAASAPIVersionKey, o.APIKey)
	if err != nil {
		return err
	}

	for i, id := range args {
		result, err := maas.Post("machines/"+id, string(v2.MachineRelease), params.Values)
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
