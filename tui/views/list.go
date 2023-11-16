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

package views

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/0x6flab/dtop/tui/styles"
	"github.com/charmbracelet/bubbles/table"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/dustin/go-humanize"
)

func ListContainers(ctx context.Context, client *client.Client) (table.Model, error) {
	containers, err := client.ContainerList(ctx, types.ContainerListOptions{})
	if err != nil {
		return table.Model{}, err
	}

	columns := []table.Column{
		{Title: "Container ID", Width: 15},
		{Title: "Name", Width: 30},
		{Title: "Image", Width: 40},
		{Title: "Command", Width: 30},
		{Title: "Created", Width: 15},
		{Title: "State", Width: 10},
		{Title: "Ports", Width: 70},
	}

	rows := []table.Row{}
	for i := range containers {
		if len(containers[i].Names) == 0 {
			continue
		}

		rows = append(rows, table.Row{
			containers[i].ID[:12],
			containers[i].Names[0][1:],
			containers[i].Image,
			containers[i].Command,
			humanize.Time(time.Unix(containers[i].Created, 0)),
			containers[i].State,
			formatPorts(containers[i].Ports),
		})
	}

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithStyles(styles.TableStyle()),
	)

	return t, nil
}

func ListImages(ctx context.Context, client *client.Client) (table.Model, error) {
	images, err := client.ImageList(ctx, types.ImageListOptions{})
	if err != nil {
		return table.Model{}, err
	}
	columns := []table.Column{
		{Title: "Image ID", Width: 15},
		{Title: "Repository", Width: 50},
		{Title: "Tag", Width: 30},
		{Title: "Created", Width: 15},
		{Title: "Size", Width: 10},
	}

	rows := []table.Row{}
	for i := range images {
		if len(images[i].RepoTags) == 0 {
			continue
		}

		if len(images[i].RepoTags) > 1 {
			for _, tag := range images[i].RepoTags {
				rows = append(rows, table.Row{
					strings.SplitAfter(images[i].ID, ":")[1][:12],
					strings.Split(tag, ":")[0],
					strings.SplitAfter(tag, ":")[1],
					humanize.Time(time.Unix(images[i].Created, 0)),
					humanize.Bytes(uint64(images[i].Size)),
				})
			}

			continue
		}

		rows = append(rows, table.Row{
			strings.SplitAfter(images[i].ID, ":")[1][:12],
			strings.Split(images[i].RepoTags[0], ":")[0],
			strings.SplitAfter(images[i].RepoTags[0], ":")[1],
			humanize.Time(time.Unix(images[i].Created, 0)),
			humanize.Bytes(uint64(images[i].Size)),
		})
	}

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithStyles(styles.TableStyle()),
	)

	return t, nil
}

func formatPorts(ports []types.Port) string {
	var portsStr string
	for _, port := range ports {
		portsStr += fmt.Sprintf("%d/%s ", port.PrivatePort, port.Type)
	}

	return portsStr
}
