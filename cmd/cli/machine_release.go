package main

import (
	"fmt"
	"os"

	gomaasclient "github.com/maas/gomaasclient/client"
	"github.com/spf13/cobra"

	"github.com/alejandroEsc/maas-cli/pkg/cli"
)

func machineReleaseCmd() *cobra.Command {
	mo := &cli.ReleaseMachineOpts{}
	cmd := &cobra.Command{
		Use:   "release [machineID]",
		Short: "Release action against one or more machines.",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {

			if err := runMachineReleaseCmd(mo, args); err != nil {
				fmt.Println(err.Error())
				os.Exit(1)
			}

		},
	}

	fs := cmd.Flags()
	bindCommonMAASFlags(&mo.MAASOptions, fs)

	fs.StringVar(&mo.Comment, "comment", "", "Optional comment for the event log")
	fs.BoolVar(&mo.Erase, "erase", true, "Erase the disk when releasing")
	fs.BoolVar(&mo.SecureErase, "secure-erase", false, "Erase the disk when releasing")
	fs.BoolVar(&mo.QuickErase, "quick-erase", true, "Erase the disk when releasing")
	return cmd
}

func runMachineReleaseCmd(o *cli.ReleaseMachineOpts, args []string) error {
	maas, err := gomaasclient.GetClient(o.MAASURLKey, o.APIKey, o.MAASAPIVersionKey)
	if err != nil {
		return err
	}

	for _, id := range args {
		result, errRelease := maas.Machine.Release(id, &o.MachineReleaseParams)
		if errRelease != nil {
			return errRelease
		}

		fmt.Printf("%v", result)

	}

	return nil
}
