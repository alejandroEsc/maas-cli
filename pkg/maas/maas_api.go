package maas

import (
	"github.com/juju/loggo"
	"net/url"
	"github.com/alejandroEsc/maas-client-sample/pkg/util"
	"github.com/juju/gomaasapi"
)

var (
	logger = util.GetModuleLogger("pkg.maas", loggo.INFO)
)

// MAASCLient encapsulates calls to maas via library calls
type MAASClient struct {
	massAPIObj *gomaasapi.MAASObject
}

// NewMaasClient returns wrapper struct for common api calls
func NewMaasClient(m *gomaasapi.MAASObject) *MAASClient {
	return &MAASClient{m}
}


// GetMachines returns a gomassapi json object from a client request
func (m *MAASClient) GetMachines() (gomaasapi.JSONObject, error) {
	logger.Infof("Fetch list of machines...")
	machineListing:= m.massAPIObj.GetSubObject("machines")
	return machineListing.CallGet("", url.Values{})
}

// GetNodes get a list of nodes.
func (m *MAASClient) GetNodes() (gomaasapi.JSONObject, error) {
	logger.Infof("Fetch list of nodes...")
	machineListing:= m.massAPIObj.GetSubObject("nodes")
	return machineListing.CallGet("", url.Values{})
}

// GetIPAddresses get list of available ip addresses
func (m *MAASClient) GetIPAddresses() (gomaasapi.JSONObject, error) {
	logger.Infof("Fetch list of ip addresses...")
	list := m.massAPIObj.GetSubObject("ipaddresses")
	return list.CallGet("", url.Values{})
}

// GetMAASVersion returns the version of maas being used
func (m *MAASClient) GetMAASVersion() (gomaasapi.JSONObject, error) {
	logger.Infof("Fetch MAAS Version...")
	version := m.massAPIObj.GetSubObject("version")
	return version.CallGet("", url.Values{})
}

// GetInterfaces returns the list of interfaces for a machine givent a system_id
func (m *MAASClient) GetInterfaces(systemID string) (gomaasapi.JSONObject, error) {
	logger.Infof("Fetch interfaces for %s...", systemID)
	interfaces := m.massAPIObj.GetSubObject("nodes").GetSubObject(systemID).GetSubObject("interfaces")
	return interfaces.CallGet("", url.Values{})
}