package maas

import (
	"net/url"

	"github.com/juju/gomaasapi"
)

type Node struct {
	Hostname   string `json:"hostname,omitempty"`
	SystemID   string `json:"system_id,omitempty"`
	Kernel     string `json:"hwe_kernel,omitempty"`
	OS         string `json:"osystem,omitempty"`
	Status     string `json:"status_name,omitempty"`
	IPAddresses []string `json:"ip_addresses,omitempty"`
	MACAddress string `json:"mac_address,omitempty"`
}
// GetNodes get a list of nodes.
func (m *Maas) GetNodes() (gomaasapi.JSONObject, error) {
	logger.Infof("Fetch list of nodes...")
	machineListing := m.massAPIObj.GetSubObject("nodes")
	return machineListing.CallGet("", url.Values{})
}
