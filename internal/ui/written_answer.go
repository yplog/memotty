package ui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/yplog/memotty/internal/models"
)

func RenderWrittenQuestion(m models.Model) string {
	var s strings.Builder

	titleStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("205")).
		MarginBottom(1)

	questionStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("86")).
		MarginBottom(1)

	inputStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("212")).
		Background(lipgloss.Color("240")).
		Padding(0, 2).
		MarginTop(1).
		MarginBottom(1)

	if !HasValidQuestions(m) {
		return RenderEmptyQuestionsError()
	}

	s.WriteString(titleStyle.Render(fmt.Sprintf("✏️ WRITTEN ANSWER MODE - Question %d/%d", m.CurrentQ+1, len(m.Questions))))
	s.WriteString("\n\n")

	currentQ := m.Questions[m.CurrentQ]
	s.WriteString(questionStyle.Render(currentQ.Question))
	s.WriteString("\n\n")

	s.WriteString("Type your answer:\n")
	inputText := m.InputText
	if inputText == "" {
		inputText = "..."
	}
	s.WriteString(inputStyle.Render(fmt.Sprintf("📝 %s", inputText)))
	s.WriteString("\n\n")

	s.WriteString(lipgloss.NewStyle().Faint(true).Render("Type and press Enter • Backspace: Delete • q: Exit"))

	return s.String()
}
