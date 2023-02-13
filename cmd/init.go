package cmd

import (
	"github.com/spf13/cobra"

	log "github.com/sirupsen/logrus"
)

var (
	initCmd = &cobra.Command{
		Use:     "init [path]",
		Aliases: []string{"initialize", "initialise", "create"},
		Short:   "Initialize a IaC Application",
		Long: `Initialize will create a new application, with a license
and the appropriate structure for a iac application.
`,

		Run: func(_ *cobra.Command, args []string) {
			log.Info("TODO:")
		},
	}
)
