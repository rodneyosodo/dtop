package styles

import (
	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/lipgloss"
)

func TableStyle() table.Styles {
	style := table.DefaultStyles()
	style.Header = style.Header.BorderStyle(lipgloss.NormalBorder()).BorderForeground(lipgloss.Color("240")).BorderBottom(true).Bold(false)
	style.Selected = style.Selected.Foreground(lipgloss.Color("229")).Background(lipgloss.Color("57")).Bold(false)
	style.Cell = style.Cell.PaddingLeft(1).PaddingRight(1).PaddingTop(0).PaddingBottom(0)

	return style
}
