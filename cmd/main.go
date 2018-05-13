package main

import (
	"encoding/json"

	m "github.com/alejandroEsc/maas-client-sample/pkg/maas"
	"github.com/alejandroEsc/maas-client-sample/pkg/util"
	"github.com/juju/gomaasapi"
	"github.com/juju/loggo"
	"github.com/spf13/viper"
)

const (
	apiKeyKey         = "api_key"
	maasURLKey        = "url"
	maasAPIVersionKey = "api_version"
)

var (
	logger loggo.Logger
)

// Init initializes the environment variables to be used by the app
func Init() {
	viper.AutomaticEnv()
	viper.SetEnvPrefix("maas_client")
	viper.BindEnv(maasURLKey)
	viper.BindEnv(apiKeyKey)
	viper.BindEnv(maasAPIVersionKey)
}

func main() {
	Init()
	apiKey := viper.GetString(apiKeyKey)
	maasURL := viper.GetString(maasURLKey)
	apiVersion := viper.GetString(maasAPIVersionKey)

	logger = util.GetModuleLogger("cmd", loggo.INFO)
	logger.Infof("%s %s %s", apiKey, maasURL, apiVersion)

	logger.Infof("Starting Sample-MAAS Client...")

	// Create API server endpoint.
	authClient, err := gomaasapi.NewAuthenticatedClient(gomaasapi.AddAPIVersionToURL(maasURL, apiVersion), apiKey)
	checkError(err)
	maas := gomaasapi.NewMAAS(*authClient)

	maasCLI := m.NewMaasClient(maas)

	getMAASVersion(maasCLI)

	listMachines(maasCLI)

	listNodes(maasCLI)

	getMachineAddresses(maasCLI)

}

func getMAASVersion(maasCLI *m.MAASclient) {
	version, err := maasCLI.GetMAASVersion()
	checkError(err)
	jp, err := json.MarshalIndent(version, "", "\t")
	checkError(err)
	logger.Infof("\n%s", jp)
}

func listMachines(maasCLI *m.MAASclient) {
	listObj, err := maasCLI.GetMachines()
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

func listNodes(maasCLI *m.MAASclient) {
	listObj, err := maasCLI.GetNodes()
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

func getMachineAddresses(maasCLI *m.MAASclient) {
	listObj, err := maasCLI.GetMachines()
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
