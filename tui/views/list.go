package views

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/table"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/dustin/go-humanize"
	"github.com/rodneyosodo/dtop/tui/styles"
)

func ListContainers(ctx context.Context, client *client.Client) (table.Model, error) {
	containers, err := client.ContainerList(ctx, types.ContainerListOptions{})
	if err != nil {
		return table.Model{}, err
	}

	columns := []table.Column{
		{Title: "Conatiner ID", Width: 15},
		{Title: "Name", Width: 30},
		{Title: "Image", Width: 40},
		{Title: "Command", Width: 30},
		{Title: "Created", Width: 15},
		{Title: "State", Width: 10},
		{Title: "Ports", Width: 70},
	}

	var rows = []table.Row{}
	for _, container := range containers {
		if len(container.Names) == 0 {
			continue
		}

		rows = append(rows, table.Row{
			container.ID[:12],
			container.Names[0][1:],
			container.Image,
			container.Command,
			humanize.Time(time.Unix(container.Created, 0)),
			container.State,
			formatPorts(container.Ports),
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

	var rows = []table.Row{}
	for _, image := range images {
		if len(image.RepoTags) == 0 {
			continue
		}

		if len(image.RepoTags) > 1 {
			for _, tag := range image.RepoTags {
				rows = append(rows, table.Row{
					strings.SplitAfter(image.ID, ":")[1][:12],
					strings.Split(tag, ":")[0],
					strings.SplitAfter(tag, ":")[1],
					humanize.Time(time.Unix(image.Created, 0)),
					humanize.Bytes(uint64(image.Size)),
				})
			}
			continue
		}

		rows = append(rows, table.Row{
			strings.SplitAfter(image.ID, ":")[1][:12],
			strings.Split(image.RepoTags[0], ":")[0],
			strings.SplitAfter(image.RepoTags[0], ":")[1],
			humanize.Time(time.Unix(image.Created, 0)),
			humanize.Bytes(uint64(image.Size)),
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
