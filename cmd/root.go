package cmd

import (
	"os"
	"time"

	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
)

const (
	appName      = "dtop"
	shortAppDesc = "A graphical CLI for your Docker instance."
	longAppDesc  = "dtop is a CLI to view and manage your Docker instance."
)

var (
	rootCmd = &cobra.Command{
		Use:   appName,
		Short: shortAppDesc,
		Long:  longAppDesc,
		RunE:  run,
	}
	logger = log.NewWithOptions(
		os.Stderr,
		log.Options{
			ReportTimestamp: true,
			TimeFormat:      time.Kitchen,
			Prefix:          "dtop 🖥",
		},
	)
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		logger.Fatal(err)
	}
}

func run(cmd *cobra.Command, args []string) error {
	logger.Infof("Starting %s", appName)

	return nil
}
