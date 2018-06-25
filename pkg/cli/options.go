package cli


// MAASOptions contains the options to allow communication with MAAS
type MAASOptions struct {
	APIKey string
	MAASURLKey string
	MAASAPIVersionKey string
}



// MachinesOptions options for the machines command
type MachineOptions struct {
	MAASOptions
	MachineID string
	MachineAction string
}

// ListMachineOptions options listing machines
type ListMachineOptions struct {
	MAASOptions
}
