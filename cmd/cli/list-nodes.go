package main

import (
	"fmt"
	"os"

	gomaasclient "github.com/maas/gomaasclient/client"
	"github.com/spf13/cobra"

	"github.com/alejandroEsc/maas-cli/pkg/cli"
)

func listNodesCmd() *cobra.Command {
	no := &cli.ListNodeOptions{}
	cmd := &cobra.Command{
		Use:   "nodes ...",
		Short: "list node resources in a MAAS server.",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {

			if err := runListNodeCmd(no); err != nil {
				fmt.Println(err.Error())
				os.Exit(1)
			}

		},
	}

	fs := cmd.Flags()

	bindCommonMAASFlags(&no.MAASOptions, fs)

	fs.BoolVar(&no.Detailed, "detailed", false, "print all details")

	return cmd
}

func runListNodeCmd(o *cli.ListNodeOptions) error {
	// Create API server endpoint.

	maas, err := gomaasclient.GetClient(o.MAASURLKey, o.APIKey, o.MAASAPIVersionKey)
	if err != nil {
		return err
	}

	devices, err := maas.Devices.Get()
	if err != nil {
		return err
	}

	fmt.Printf("%+v\n", devices)

	return nil
}
