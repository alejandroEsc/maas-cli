package maas

import (
	"fmt"
	"net/url"

	"github.com/juju/gomaasapi"
)

// Machine is a convenient internal representation of a machine
type Machine struct {
	Hostname    string   `json:"hostname,omitempty"`
	SystemID    string   `json:"system_id,omitempty"`
	Kernel      string   `json:"hwe_kernel,omitempty"`
	OS          string   `json:"osystem,omitempty"`
	PowerState  string   `json:"power_state,omitempty"`
	Status      string   `json:"status_name,omitempty"`
	IPAddresses []string `json:"ip_addresses,omitempty"`
	MACAddress  string   `json:"mac_address,omitempty"`
}

// MachineAction represents actions that can be taken against a machine
type MachineAction string

const (
	// CommissionMachine is the action of commissioning
	CommissionMachine MachineAction = "commission"
	// ReleaseMachine is the action of releasing
	ReleaseMachine MachineAction = "release"
	// DeployMachine is the action of deploying
	DeployMachine MachineAction = "deploy"
)

// DefaultParams returns, depending on a particular action, a set of query parameters
func DefaultParams(action MachineAction) url.Values {
	switch action {
	case CommissionMachine:
		return url.Values{
			"enable_ssh":      {"1"},
			"skip_networking": {"0"},
			"skip_storage":    {"0"},
		}
	case ReleaseMachine:
		return url.Values{
			"erase":        {"1"},
			"secure_erase": {"0"},
			"quick_erase":  {"1"},
		}
	case DeployMachine:
		return url.Values{
			"distro_series": {"ubuntu"},
			"hwe_kernel":    {"ga-16.04"},
			"comment":       {"deployed by Maas cli"},
		}
	default:
		return url.Values{}

	}
}

// GetMachines returns a gomassapi json object from a client request
func (m *Maas) GetMachines() (gomaasapi.JSONObject, error) {
	logger.Infof("Fetch list of machines...")
	machineListing := m.massAPIObj.GetSubObject("machines")
	return machineListing.CallGet("", url.Values{})
}

// PerformMachineAction performs a commission
func (m *Maas) PerformMachineAction(action MachineAction, systemID string, params url.Values) (gomaasapi.JSONObject, error) {
	logger.Infof("%s Machine %s...", action, systemID)
	machineSubObject := m.massAPIObj.GetSubObject("machines").GetSubObject(systemID)
	if params == nil {
		params = DefaultParams(action)
	}

	return machineSubObject.CallPost(fmt.Sprintf("%s", action), params)
}

// GetStatus Returns the status of a machine given a systemID
func (m *Maas) GetStatus(systemID string) (string, error) {
	machineSubObject := m.massAPIObj.GetSubObject("machines").GetSubObject(systemID)

	obj, err := machineSubObject.CallGet("", url.Values{})
	if err != nil {
		return "", err
	}

	machine, err := obj.GetMAASObject()
	if err != nil {
		return "", err
	}

	// we dont care about individual error states here.
	power, err := machine.GetField("power_state")
	if err != nil {
		logger.Errorf(err.Error())
	}

	status, err := machine.GetField("status_name")
	if err != nil {
		logger.Errorf(err.Error())
	}

	return fmt.Sprintf("|\t %s \t|\t %s \t|\t %s \t|", systemID, power, status), nil
}
