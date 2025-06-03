package ui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/yplog/memotty/internal/models"
)

func RenderMultipleChoiceQuestion(m models.Model) string {
	var s strings.Builder

	titleStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("205")).
		MarginBottom(1)

	questionStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("86")).
		MarginBottom(1)

	selectedStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("212")).
		Background(lipgloss.Color("57")).
		Padding(0, 1)

	normalStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("252")).
		Padding(0, 1)

	if !HasValidQuestions(m) {
		return RenderEmptyQuestionsError()
	}

	s.WriteString(titleStyle.Render(fmt.Sprintf("📋 MULTIPLE CHOICE MODE - Question %d/%d", m.CurrentQ+1, len(m.Questions))))
	s.WriteString("\n\n")

	currentQ := m.Questions[m.CurrentQ]
	s.WriteString(questionStyle.Render(currentQ.Question))
	s.WriteString("\n\n")

	for i, option := range currentQ.Options {
		cursor := "  "
		if m.Cursor == i {
			cursor = "👉"
		}

		if m.Cursor == i {
			s.WriteString(fmt.Sprintf("%s %s\n", cursor, selectedStyle.Render(fmt.Sprintf("%c) %s", 'A'+i, option))))
		} else {
			s.WriteString(fmt.Sprintf("%s %s\n", cursor, normalStyle.Render(fmt.Sprintf("%c) %s", 'A'+i, option))))
		}
	}

	s.WriteString("\n")
	s.WriteString(lipgloss.NewStyle().Faint(true).Render("↑/↓: Change option • Enter: Answer • q: Exit"))

	return s.String()
}
