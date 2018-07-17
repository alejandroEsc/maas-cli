package main

import (
	"github.com/spf13/cobra"

	"fmt"
	"os"

	"encoding/json"
	"net/url"

	"github.com/alejandroEsc/golang-maas-client/pkg/api"
	"github.com/alejandroEsc/golang-maas-client/pkg/api/v2"
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
	var err error

	maas, err := api.NewMASS(o.MAASURLKey, o.MAASAPIVersionKey, o.APIKey)
	if err != nil {
		return err
	}

	for _, id := range args {
		result, err := maas.Get("machines/"+id, "", url.Values{})

		var machine v2.Machine
		err = json.Unmarshal(result, &machine)
		if err != nil {
			logger.Errorf(err.Error())
			continue
		}

		fmt.Printf("\t %s \t %s \t %s \t\n", machine.SystemID, machine.PowerState, machine.StatusName)
	}

	return nil
}
