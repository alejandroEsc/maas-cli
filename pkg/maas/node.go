package maas

import (
	"github.com/juju/gomaasapi"
	"net/url"
)

// GetNodes get a list of nodes.
func (m *Maas) GetNodes() (gomaasapi.JSONObject, error) {
	logger.Infof("Fetch list of nodes...")
	machineListing := m.massAPIObj.GetSubObject("nodes")
	return machineListing.CallGet("", url.Values{})
}
