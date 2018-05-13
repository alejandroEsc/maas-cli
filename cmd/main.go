package main

import (
	"github.com/alejandroEsc/maas-client-sample/pkg/util"
	m "github.com/alejandroEsc/maas-client-sample/pkg/maas"
	"github.com/juju/loggo"
	"github.com/spf13/viper"
	"github.com/juju/gomaasapi"
	"encoding/json"
)

const (
	apiKeyKey = "api_key"
	maasUrlKey = "url"
	maasAPIVersionKey = "api_version"
)

var (
	logger        loggo.Logger
)


// Init initializes the environment variables to be used by the app
func Init() {
	viper.AutomaticEnv()
	viper.SetEnvPrefix("maas_client")
	viper.BindEnv(maasUrlKey)
	viper.BindEnv(apiKeyKey)
	viper.BindEnv(maasAPIVersionKey)
}


func main() {
	Init()
	apiKey := viper.GetString(apiKeyKey)
	maasUrl := viper.GetString(maasUrlKey)
	apiVersion := viper.GetString(maasAPIVersionKey)


	logger = util.GetModuleLogger("cmd", loggo.INFO)
	logger.Infof("%s %s %s",apiKey, maasUrl, apiVersion)

	logger.Infof("Starting Sample-MAAS Client...")

	// Create API server endpoint.
	authClient, err := gomaasapi.NewAuthenticatedClient(gomaasapi.AddAPIVersionToURL(maasUrl, apiVersion), apiKey)
	checkError(err)
	maas := gomaasapi.NewMAAS(*authClient)

	maasCLI := m.NewMaasClient(maas)

	getMAASVersion(maasCLI)

	listMachines(maasCLI)

	listNodes(maasCLI)

	getMachineAddresses(maasCLI)

}

func getMAASVersion(maasCLI *m.MAASClient) {
	version, err := maasCLI.GetMAASVersion()
	checkError(err)
	jp, err := json.MarshalIndent(version, "", "\t")
	checkError(err)
	logger.Infof("\n%s", jp)
}


// ManipulateFiles exercises the /api/1.0/nodes/ API endpoint.  Most precisely,
// it lists the existing nodes, creates a new node, updates it and then
// deletes it.
func listMachines(maasCLI *m.MAASClient) {
	listObj, err:= maasCLI.GetMachines()
	checkError(err)

	machinesArray, err := listObj.GetArray()
	checkError(err)

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

		logger.Infof("|\t %d \t|\t %s \t|\t %s \t|\t %s-%s \t|", i, machineSystemID, machineName, os, hweKernel)
	}
}

func listNodes(maasCLI *m.MAASClient) {
	listObj, err:= maasCLI.GetNodes()
	checkError(err)

	nodesArray, err := listObj.GetArray()
	checkError(err)

	for i, nodeObj := range nodesArray {
		node, err := nodeObj.GetMAASObject()
		checkError(err)

		name, err := node.GetField("hostname")
		checkErrorMsg(err, "could not get hostname")

		systemID, err := node.GetField("system_id")
		checkErrorMsg(err, "could not get system_id")

		hweKernel, _ := node.GetField("hwe_kernel")
		os, _ := node.GetField("osystem")
		ips, _ := node.GetField("ip_addresses")

		logger.Infof("|\t %d \t|\t %s \t|\t %s \t|\t %s \t|\t %s-%s \t|", i, systemID, name, ips, os, hweKernel)
	}

}

func getMachineAddresses(maasCLI *m.MAASClient){
	listObj, err:= maasCLI.GetMachines()
	checkError(err)

	machinesArray, err := listObj.GetArray()
	checkError(err)

	logger.Infof("%d", len(machinesArray))

	for _, machineObj := range machinesArray {
		machine, err := machineObj.GetMAASObject()
		checkError(err)

		machineSystemID, err := machine.GetField("system_id")
		checkErrorMsg(err, "could not get system_id")
		if err == nil {
			interfaces, err := maasCLI.GetInterfaces(machineSystemID)
			checkError(err)
			jp, err := json.MarshalIndent(interfaces, "", "\t")
			checkError(err)
			logger.Infof("\n%s", jp)

		}
	}

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