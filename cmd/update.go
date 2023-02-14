package cmd

import (
	"github.com/spf13/cobra"

	log "github.com/sirupsen/logrus"
)

var (
	updateCmd = &cobra.Command{
		Use:     "update [path]",
		Aliases: []string{"update"},
		Short:   "update a IaC Application",
		Long: `Update will update a new application.`,

		Run: func(_ *cobra.Command, args []string) {
			log.Info("TODO:")
		},
	}
)
