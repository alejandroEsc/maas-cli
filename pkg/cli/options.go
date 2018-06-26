package cli

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
