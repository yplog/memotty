package ui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/yplog/memotty/internal/models"
)

func RenderMenu(m models.Model) string {
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

	s.WriteString(titleStyle.Render("🎯 MEMOTTY - MODE SELECTION"))
	s.WriteString("\n\n")

	if m.SelectedFile != "" {
		s.WriteString(infoStyle.Render(fmt.Sprintf("📄 Selected file: %s", m.SelectedFile)))
		s.WriteString("\n\n")
	}

	s.WriteString("Please select a quiz mode:\n\n")

	modes := []string{
		"📋 Multiple Choice Mode (A, B, C, D options)",
		"✏️ Written Answer Mode (Type your own answer)",
	}

	for i, mode := range modes {
		cursor := "  "
		if m.Cursor == i {
			cursor = "👉"
		}

		if m.Cursor == i {
			s.WriteString(fmt.Sprintf("%s %s\n", cursor, selectedStyle.Render(mode)))
		} else {
			s.WriteString(fmt.Sprintf("%s %s\n", cursor, normalStyle.Render(mode)))
		}
	}

	s.WriteString("\n")
	s.WriteString(lipgloss.NewStyle().Faint(true).Render("↑/↓: Select mode • Enter: Start • b: Back to files • q: Exit"))

	return s.String()
}
