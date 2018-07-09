package main

import (
	m "github.com/alejandroEsc/maas-cli/pkg/maas"
	"github.com/spf13/cobra"

	"fmt"
	"os"

	"encoding/json"

	"sort"

	"github.com/alejandroEsc/maas-cli/pkg/cli"
	"github.com/juju/gomaasapi"
)

const (
	printNodeFmt = "\t %d \t %s \t %s \t %s \t %s \t %s \t\n"
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
	nodeSlice := make([]m.Node, 0)

	for _, nodeObj := range nodeArray {
		var n m.Node
		node, err := nodeObj.GetMAASObject()
		logError(err)

		j, err := node.MarshalJSON()
		logError(err)

		err = json.Unmarshal(j, &n)
		logError(err)

		nodeSlice = append(nodeSlice, n)
	}

	sort.Slice(nodeSlice, func(i, j int) bool {
		return nodeSlice[i].Status < nodeSlice[j].Status
	})

	for i, n := range nodeSlice {
		fmt.Printf(printNodeFmt, i, n.SystemID, n.Hostname, n.IPAddresses, n.OS, n.Kernel)
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
