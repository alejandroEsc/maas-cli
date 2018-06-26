package main

import (
	m "github.com/alejandroEsc/maas-cli/pkg/maas"
	"github.com/spf13/cobra"

	"fmt"
	"os"

	"github.com/alejandroEsc/maas-cli/pkg/cli"
	"github.com/juju/gomaasapi"
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

	// Create API server endpoint.
	authClient, err := gomaasapi.NewAuthenticatedClient(gomaasapi.AddAPIVersionToURL(o.MAASURLKey, o.MAASAPIVersionKey), o.APIKey)
	if err != nil {
		return err
	}

	maas := gomaasapi.NewMAAS(*authClient)
	maasCLI := m.NewMaas(maas)

	for _, id := range args {
		result, err := maasCLI.GetStatus(id)
		if err != nil {
			logger.Errorf(err.Error())
			continue
		}

		fmt.Println(result)
	}

	return nil
}
