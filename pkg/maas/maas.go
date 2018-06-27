package maas

import (
	"net/url"

	"github.com/alejandroEsc/maas-cli/pkg/util"
	"github.com/juju/gomaasapi"
	"github.com/juju/loggo"
)

var (
	logger = util.GetModuleLogger("pkg.Maas", loggo.INFO)
)

// Maas encapsulates calls to Maas via library calls
type Maas struct {
	massAPIObj *gomaasapi.MAASObject
}

// NewMaas returns wrapper struct for common api calls
func NewMaas(m *gomaasapi.MAASObject) *Maas {
	return &Maas{m}
}

// GetIPAddresses get list of available ip addresses
func (m *Maas) GetIPAddresses() (gomaasapi.JSONObject, error) {
	logger.Infof("Fetch list of ip addresses...")
	list := m.massAPIObj.GetSubObject("ipaddresses")
	return list.CallGet("", url.Values{})
}

// GetMAASVersion returns the version of Maas being used
func (m *Maas) GetMAASVersion() (gomaasapi.JSONObject, error) {
	logger.Infof("Fetch MAAS Version...")
	version := m.massAPIObj.GetSubObject("version")
	return version.CallGet("", url.Values{})
}

// GetInterfaces returns the list of interfaces for a machine givent a system_id
func (m *Maas) GetInterfaces(systemID string) (gomaasapi.JSONObject, error) {
	logger.Infof("Fetch interfaces for %s...", systemID)
	interfaces := m.massAPIObj.GetSubObject("nodes").GetSubObject(systemID).GetSubObject("interfaces")
	return interfaces.CallGet("", url.Values{})
}
