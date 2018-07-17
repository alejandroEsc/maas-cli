package main

import (
	"github.com/spf13/cobra"

	"fmt"
	"os"

	"encoding/json"

	"github.com/alejandroEsc/golang-maas-client/pkg/api"
	"github.com/alejandroEsc/golang-maas-client/pkg/api/v2"
	"github.com/alejandroEsc/maas-cli/pkg/cli"
)

const (
	printNodeFmt = "\t %d \t %s \t %s \t %s\n"
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
	maas, err := api.NewMASS(o.MAASURLKey, o.MAASAPIVersionKey, o.APIKey)
	if err != nil {
		return err
	}

	params := v2.NodesParams(v2.NodesArgs{})

	rawNodes, err := maas.Get("nodes", "", params.Values)

	var nodes []v2.Node
	err = json.Unmarshal(rawNodes, &nodes)
	if err != nil {
		return err
	}

	if o.Detailed {
		return printNodesDetailed(nodes)
	}

	printNodesSummary(nodes)

	return nil
}

func printNodesSummary(nodes []v2.Node) {
	for i, n := range nodes {
		fmt.Printf(printNodeFmt, i, n.SystemID, n.Hostname, n.IPAddresses)

	}
}

func printNodesDetailed(nodes []v2.Node) error {
	for i, n := range nodes {
		fmt.Printf("\n --- node: %d ---\n", i)
		fmt.Printf("%+v\n", n)

	}
	return nil
}
