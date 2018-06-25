package maas

import (
	"github.com/juju/gomaasapi"
	"net/url"
)

type Machine struct {
	Name string
	SystemID string
	Kernel string
	OS string
}




func (m *MAASclient) CommisionMachine(systemID string) (gomaasapi.JSONObject, error) {
	logger.Infof("Commission Machine %s...", systemID)
	params := url.Values{
		"enable_ssh":{"1"},
		"skip_networking":{"0"},
		"skip_storage": {"0"},
	}


	machineSubObject := m.massAPIObj.GetSubObject("machines").GetSubObject(systemID)
	return machineSubObject.CallPost("commission", params)
}

func (m *MAASclient) ReleaseMachine(systemID string) (gomaasapi.JSONObject, error) {
	logger.Infof("Release Machine %s...", systemID)
	params := url.Values{
		"erase":{"1"},
		"secure_erase":{"0"},
		"quick_erase": {"1"},
	}


	machineSubObject := m.massAPIObj.GetSubObject("machines").GetSubObject(systemID)
	return machineSubObject.CallPost("release", params)
}

func (m *MAASclient) DeployMachine(systemID string) (gomaasapi.JSONObject, error) {
	logger.Infof("Release Machine %s...", systemID)
	params := url.Values{
		"distro_series":{"ubuntu"},
		"hwe_kernel":{"ga-16.04"},
		"comment": {"deployed by maas cli"},
	}


	machineSubObject := m.massAPIObj.GetSubObject("machines").GetSubObject(systemID)
	return machineSubObject.CallPost("deploy", params)
}