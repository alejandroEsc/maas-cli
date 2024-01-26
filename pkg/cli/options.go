package cli

import "github.com/maas/gomaasclient/entity"

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

// CommissionMachineOpts represents commission machine options
type CommissionMachineOpts struct {
	MAASOptions

	entity.MachineCommissionParams
}

// DeployMachineOpts represents deploy machine options
type DeployMachineOpts struct {
	MAASOptions

	entity.MachineDeployParams
}

// ReleaseMachineOpts represents release machine options
type ReleaseMachineOpts struct {
	MAASOptions

	entity.MachineReleaseParams
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

	entity.MachineParams
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
