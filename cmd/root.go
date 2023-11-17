// Copyright (c) 2023 0x6flab
//
// SPDX-License-Identifier: GPL-3.0-only
//
// This program is free software: you can redistribute it and/or modify it under
// the terms of the GNU General Public License as published by the Free Software
// Foundation, version 3.
//
// This program is distributed in the hope that it will be useful, but WITHOUT ANY
// WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A
// PARTICULAR PURPOSE. See the GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License along with
// this program. If not, see <https://www.gnu.org/licenses/>.

package cmd

import (
	"os"
	"time"

	"github.com/0x6flab/dtop/tui"
	tea "github.com/charmbracelet/bubbletea"
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

func run(cmd *cobra.Command, _ []string) error {
	m, err := tui.NewModel(cmd.Context())
	if err != nil {
		return err
	}

	if _, err := tea.NewProgram(m, tea.WithAltScreen(), tea.WithContext(cmd.Context())).Run(); err != nil {
		logger.Fatal(err)
	}

	return nil
}
