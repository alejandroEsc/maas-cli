package maas

import (
	"fmt"
	"net/url"

	"github.com/juju/gomaasapi"
)

// Machine is a convenient internal representation of a machine
type Machine struct {
	Name       string
	SystemID   string
	Kernel     string
	OS         string
	PowerState string
	Status     string
}

// MachineAction represents actions that can be taken against a machine
type MachineAction string

const (
	CommissionMachine MachineAction = "commission"
	ReleaseMachine    MachineAction = "release"
	DeployMachine     MachineAction = "deploy"
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
			"comment":       {"deployed by maas cli"},
		}
	default:
		return url.Values{}

	}
}

// PerformMachineAction performs a commission
func (m *MAASclient) PerformMachineAction(action MachineAction, systemID string, params url.Values) (gomaasapi.JSONObject, error) {
	logger.Infof("%s Machine %s...", action, systemID)
	machineSubObject := m.massAPIObj.GetSubObject("machines").GetSubObject(systemID)
	if params == nil {
		params = DefaultParams(action)
	}

	return machineSubObject.CallPost(fmt.Sprintf("%s", action), params)
}

// GetStatus Returns the status of a machine given a systemID
func (m *MAASclient) GetStatus(systemID string) (string, error) {
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
