package cli

import (
	"github.com/alejandroEsc/golang-maas-client/pkg/api/v2"
)

// MAASOptions contains the options to allow communication with MAAS
type MAASOptions struct {
	APIKey            string
	MAASURLKey        string
	MAASAPIVersionKey string
}

// MachineOptions options for the machines command
type MachineOptions struct {
	MAASOptions
	Params string
}

type CommissionMachineOpts struct {
	MAASOptions
	v2.CommissionMachineArgs
}

type DeployMachineOpts struct {
	MAASOptions
	v2.DeployMachineArgs
}

type ReleaseMachineOpts struct {
	MAASOptions
	v2.ReleaseMachinesArgs
}

// ListOptions options listing machines
type ListOptions struct {
	MAASOptions
	Detailed bool
}

// ListMachineOptions options listing machines
type ListMachineOptions struct {
	MAASOptions
	Detailed bool
}

// ListNodeOptions options listing machines
type ListNodeOptions struct {
	MAASOptions
	Detailed bool
}

// VersionOptions version info
type VersionOptions struct {
	MAASOptions
}
