package ui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/yplog/memotty/internal/models"
)

func RenderFileSelection(m models.Model) string {
	var s strings.Builder

	titleStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("205")).
		MarginBottom(2)

	selectedStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("212")).
		Background(lipgloss.Color("57")).
		Padding(0, 2).
		MarginBottom(1)

	normalStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("252")).
		Padding(0, 2).
		MarginBottom(1)

	infoStyle := lipgloss.NewStyle().
		Italic(true).
		Foreground(lipgloss.Color("240"))

	s.WriteString(titleStyle.Render("📁 MEMOTTY - SELECT QUESTION FILE"))
	s.WriteString("\n\n")

	if len(m.CSVFiles) == 0 {
		s.WriteString("No CSV files found in ~/.memotty/\n\n")
		s.WriteString(infoStyle.Render("A sample file has been created for you."))
		s.WriteString("\n\n")
		s.WriteString(lipgloss.NewStyle().Faint(true).Render("Press q to exit"))
	} else {
		s.WriteString("Select a question file from ~/.memotty/:\n\n")

		for i, file := range m.CSVFiles {
			cursor := "  "
			if m.Cursor == i {
				cursor = "👉"
			}

			if m.Cursor == i {
				s.WriteString(fmt.Sprintf("%s %s\n", cursor, selectedStyle.Render(fmt.Sprintf("📄 %s", file))))
			} else {
				s.WriteString(fmt.Sprintf("%s %s\n", cursor, normalStyle.Render(fmt.Sprintf("📄 %s", file))))
			}
		}

		s.WriteString("\n")
		s.WriteString(infoStyle.Render("CSV format: question,answer (one per line)"))
		s.WriteString("\n\n")
		s.WriteString(lipgloss.NewStyle().Faint(true).Render("↑/↓: Select file • Enter: Continue • u: Update • q: Exit"))
		s.WriteString("\n\n")
		s.WriteString(lipgloss.NewStyle().Faint(true).Render(GetDetailedVersionInfo()))
	}

	return s.String()
}
