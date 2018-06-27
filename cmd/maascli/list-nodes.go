package main

import (
	m "github.com/alejandroEsc/maas-cli/pkg/maas"
	"github.com/spf13/cobra"

	"fmt"
	"os"

	"github.com/alejandroEsc/maas-cli/pkg/cli"
	"github.com/juju/gomaasapi"
)

const (
	printNodeFmt = "|\t %d \t|\t %s \t|\t %s \t|\t %s \t|\t %s-%s \t|\n"
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
	authClient, err := gomaasapi.NewAuthenticatedClient(gomaasapi.AddAPIVersionToURL(o.MAASURLKey, o.MAASAPIVersionKey), o.APIKey)
	if err != nil {
		return err
	}

	maas := gomaasapi.NewMAAS(*authClient)
	maasCLI := m.NewMaas(maas)

	listObj, err := maasCLI.GetNodes()
	if err != nil {
		return err
	}

	nodesArray, err := listObj.GetArray()
	if err != nil {
		return err
	}

	if o.Detailed {
		return printNodesDetailed(nodesArray)
	}

	printNodesSummary(nodesArray)

	return nil
}

func printNodesSummary(nodeArray []gomaasapi.JSONObject) {

	for i, nodeObj := range nodeArray {
		node, err := nodeObj.GetMAASObject()
		logError(err)

		name, err := node.GetField("hostname")
		logError(err)

		systemID, err := node.GetField("system_id")
		logError(err)

		hweKernel, _ := node.GetField("hwe_kernel")
		os, _ := node.GetField("osystem")
		ips, _ := node.GetField("ip_addresses")

		fmt.Printf(printNodeFmt, i, systemID, name, ips, os, hweKernel)
	}

}

func printNodesDetailed(nodesArray []gomaasapi.JSONObject) error {
	for i, nodeObj := range nodesArray {
		node, err := nodeObj.GetMAASObject()
		j, err := node.MarshalJSON()
		if err != nil {
			return err
		}
		fmt.Printf("\n --- node: %d ---\n", i)
		fmt.Printf("%s", j)
	}
	return nil
}
